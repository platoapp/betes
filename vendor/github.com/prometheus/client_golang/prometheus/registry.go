// Copyright 2014 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package prometheus

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"sort"
	"sync"

	"github.com/golang/protobuf/proto"

	dto "github.com/prometheus/client_model/go"
)

const (
	// Capacity for the channel to collect metrics and descriptors.
	capMetricChan = 1000
	capDescChan   = 10
)

// DefaultRegisterer and DefaultGatherer are the implementations of the
// Registerer and Gatherer interface a number of convenience functions in this
// package act on. Initially, both variables point to the same Registry, which
// has a process collector (see NewProcessCollector) and a Go collector (see
// NewGoCollector) already registered. This approach to keep default instances
// as global state mirrors the approach of other packages in the Go standard
// library. Note that there are caveats. Change the variables with caution and
// only if you understand the consequences. Users who want to avoid global state
// altogether should not use the convenience function and act on custom
// instances instead.
var (
	defaultRegistry              = NewRegistry()
	DefaultRegisterer Registerer = defaultRegistry
	DefaultGatherer   Gatherer   = defaultRegistry
)

func init() {
	MustRegister(NewProcessCollector(os.Getpid(), ""))
	MustRegister(NewGoCollector())
}

// NewRegistry creates a new vanilla Registry without any Collectors
// pre-registered.
func NewRegistry() *Registry {
	return &Registry{
		collectorsByID:  map[uint64]Collector{},
		descIDs:         map[uint64]struct{}{},
		dimHashesByName: map[string]uint64{},
	}
}

// NewPedanticRegistry returns a registry that checks during collection if each
// collected Metric is consistent with its reported Desc, and if the Desc has
// actually been registered with the registry.
//
// Usually, a Registry will be happy as long as the union of all collected
// Metrics is consistent and valid even if some metrics are not consistent with
// their own Desc or a Desc provided by their registered Collector. Well-behaved
// Collectors and Metrics will only provide consistent Descs. This Registry is
// useful to test the implementation of Collectors and Metrics.
func NewPedanticRegistry() *Registry {
	r := NewRegistry()
	r.pedanticChecksEnabled = true
	return r
}

// Registerer is the interface for the part of a registry in charge of
// registering and unregistering. Users of custom registries should use
// Registerer as type for registration purposes (rather then the Registry type
// directly). In that way, they are free to use custom Registerer implementation
// (e.g. for testing purposes).
type Registerer interface {
	// Register registers a new Collector to be included in metrics
	// collection. It returns an error if the descriptors provided by the
	// Collector are invalid or if they — in combination with descriptors of
	// already registered Collectors — do not fulfill the consistency and
	// uniqueness criteria described in the documentation of metric.Desc.
	//
	// If the provided Collector is equal to a Collector already registered
	// (which includes the case of re-registering the same Collector), the
	// returned error is an instance of AlreadyRegisteredError, which
	// contains the previously registered Collector.
	//
	// It is in general not safe to register the same Collector multiple
	// times concurrently.
	Register(Collector) error
	// MustRegister works like Register but registers any number of
	// Collectors and panics upon the first registration that causes an
	// error.
	MustRegister(...Collector)
	// Unregister unregisters the Collector that equals the Collector passed
	// in as an argument.  (Two Collectors are considered equal if their
	// Describe method yields the same set of descriptors.) The function
	// returns whether a Collector was unregistered.
	//
	// Note that even after unregistering, it will not be possible to
	// register a new Collector that is inconsistent with the unregistered
	// Collector, e.g. a Collector collecting metrics with the same name but
	// a different help string. The rationale here is that the same registry
	// instance must only collect consistent metrics throughout its
	// lifetime.
	Unregister(Collector) bool
}

// Gatherer is the interface for the part of a registry in charge of gathering
// the collected metrics into a number of MetricFamilies. The Gatherer interface
// comes with the same general implication as described for the Registerer
// interface.
type Gatherer interface {
	// Gather calls the Collect method of the registered Collectors and then
	// gathers the collected metrics into a lexicographically sorted slice
	// of MetricFamily protobufs. Even if an error occurs, Gather attempts
	// to gather as many metrics as possible. Hence, if a non-nil error is
	// returned, the returned MetricFamily slice could be nil (in case of a
	// fatal error that prevented any meaningful metric collection) or
	// contain a number of MetricFamily protobufs, some of which might be
	// incomplete, and some migh