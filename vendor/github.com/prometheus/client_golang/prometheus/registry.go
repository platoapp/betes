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
	