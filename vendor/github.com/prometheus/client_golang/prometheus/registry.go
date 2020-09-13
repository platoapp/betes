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
	// incomplete, and some might be missing altogether. The returned error
	// (which might be a MultiError) explains the details. In scenarios
	// where complete collection is critical, the returned MetricFamily
	// protobufs should be disregarded if the returned error is non-nil.
	Gather() ([]*dto.MetricFamily, error)
}

// Register registers the provided Collector with the DefaultRegisterer.
//
// Register is a shortcut for DefaultRegisterer.Register(c). See there for more
// details.
func Register(c Collector) error {
	return DefaultRegisterer.Register(c)
}

// MustRegister registers the provided Collectors with the DefaultRegisterer and
// panics if any error occurs.
//
// MustRegister is a shortcut for DefaultRegisterer.MustRegister(cs...). See
// there for more details.
func MustRegister(cs ...Collector) {
	DefaultRegisterer.MustRegister(cs...)
}

// RegisterOrGet registers the provided Collector with the DefaultRegisterer and
// returns the Collector, unless an equal Collector was registered before, in
// which case that Collector is returned.
//
// Deprecated: RegisterOrGet is merely a convenience function for the
// implementation as described in the documentation for
// AlreadyRegisteredError. As the use case is relatively rare, this function
// will be removed in a future version of this package to clean up the
// namespace.
func RegisterOrGet(c Collector) (Collector, error) {
	if err := Register(c); err != nil {
		if are, ok := err.(AlreadyRegisteredError); ok {
			return are.ExistingCollector, nil
		}
		return nil, err
	}
	return c, nil
}

// MustRegisterOrGet behaves like RegisterOrGet but panics instead of returning
// an error.
//
// Deprecated: This is deprecated for the same reason RegisterOrGet is. See
// there for details.
func MustRegisterOrGet(c Collector) Collector {
	c, err := RegisterOrGet(c)
	if err != nil {
		panic(err)
	}
	return c
}

// Unregister removes the registration of the provided Collector from the
// DefaultRegisterer.
//
// Unregister is a shortcut for DefaultRegisterer.Unregister(c). See there for
// more details.
func Unregister(c Collector) bool {
	return DefaultRegisterer.Unregister(c)
}

// GathererFunc turns a function into a Gatherer.
type GathererFunc func() ([]*dto.MetricFamily, error)

// Gather implements Gatherer.
func (gf GathererFunc) Gather() ([]*dto.MetricFamily, error) {
	return gf()
}

// SetMetricFamilyInjectionHook replaces the DefaultGatherer with one that
// gathers from the previous DefaultGatherers but then merges the MetricFamily
// protobufs returned from the provided hook function with the MetricFamily
// protobufs returned from the original DefaultGatherer.
//
// Deprecated: This function manipulates the DefaultGatherer variable. Consider
// the implications, i.e. don't do this concurrently with any uses of the
// DefaultGatherer. In the rare cases where you need to inject MetricFamily
// protobufs directly, it is recommended to use a custom Registry and combine it
// with a custom Gatherer using the Gatherers type (see
// there). SetMetricFamilyInjectionHook only exists for compatibility reasons
// with previous versions of this package.
func SetMetricFamilyInjectionHook(hook func() []*dto.MetricFamily) {
	DefaultGatherer = Gatherers{
		DefaultGatherer,
		GathererFunc(func() ([]*dto.MetricFamily, error) { return hook(), nil }),
	}
}

// AlreadyRegisteredError is returned by the Register method if the Collector to
// be registered has already been registered before, or a different Collector
// that collects the same metrics has been registered before. Registration fails
// in that case, but you can detect from the kind of error what has
// happened. The error contains fields for the existing Collector and the
// (rejected) new Collector that equals the existing one. This can be used to
// find out if an equal Collector has been registered before and switch over to
// using the old one, as demonstrated in the example.
type AlreadyRegisteredError struct {
	ExistingCollector, NewCollector Collector
}

func (err AlreadyRegisteredError) Error() string {
	return "duplicate metrics collector registration attempted"
}

// MultiError is a slice of errors implementing the error interface. It is used
// by a Gatherer to report multiple errors during MetricFamily gathering.
type MultiError []error

func (errs MultiError) Error() string {
	if len(errs) == 0 {
		return ""
	}
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "%d error(s) occurred:", len(errs))
	for _, err := range errs {
		fmt.Fprintf(buf, "\n* %s", err)
	}
	return buf.String()
}

// MaybeUnwrap returns nil if len(errs) is 0. It returns the first and only
// contained error as error if len(errs is 1). In all other cases, it returns
// the MultiError directly. This is helpful for returning a MultiError in a way
// that only uses the MultiError if needed.
func (errs MultiError) MaybeUnwrap() error {
	switch len(errs) {
	case 0:
		return nil
	case 1:
		return errs[0]
	default:
		return errs
	}
}

// Registry registers Prometheus collectors, collects their metrics, and gathers
// them into MetricFamilies for exposition. It implements both Registerer and
// Gatherer. The zero value is not usable. Create instances with NewRegistry or
// NewPedanticRegistry.
type Registry struct {
	mtx                   sync.RWMutex
	collectorsByID        map[uint64]Collector // ID is a hash of the descIDs.
	descIDs               map[uint64]struct{}
	dimHashesByName       map[string]uint64
	pedanticChecksEnabled bool
}

// Register implements Registerer.
func (r *Registry) Register(c Collector) error {
	var (
		descChan           = make(chan *Desc, capDescChan)
		newDescIDs         = map[uint64]struct{}{}
		newDimHashesByName = map[string]uint64{}
		collectorID        uint64 // Just a sum of all desc IDs.
		duplicateDescErr   error
	)
	go func() {
		c.Describe(descChan)
		close(descChan)
	}()
	r.mtx.Lock()
	defer r.mtx.Unlock()
	// Coduct various tests...
	for desc := range descChan {

		// Is the descriptor valid at all?
		if desc.err != nil {
			return fmt.Errorf("descriptor %s is invalid: %s", desc, desc.err)
		}

		// Is the descID unique?
		// (In other words: Is the fqName + constLabel combination unique?)
		if _, exists := r.descIDs[desc.id]; exists {
			duplicateDescErr = fmt.Errorf("descriptor %s already exists with the same fully-qualified name and const label values", desc)
		}
		// If it is not a duplicate desc in this collector, add it to
		// the collectorID.  (We allow duplicate descs within the same
		// collector, but their existence must be a no-op.)
		if _, exists := newDescIDs[desc.id]; !exists {
			newDescIDs[desc.id] = struct{}{}
			collectorID += desc.id
		}

		// Are all the label names and the help string consistent with
		// previous descriptors of the same name?
		// First check existing descriptors...
		if dimHash, exists := r.dimHashesByName[desc.fqName]; exists {
			if dimHash != desc.dimHash {
				return fmt.Errorf("a previously registered descriptor with the same fully-qualified name as %s has different label names or a different help string", desc)
			}
		} else {
			// ...then check the new descriptors already seen.
			if dimHash, exists := newDimHashesByName[desc.fqName]; exists {
				if dimHash != desc.dimHash {
					return fmt.Errorf("descriptors reported by collector have inconsistent label names or help strings for the same fully-qualified name, offender is %s", desc)
				}
			} else {
				newDimHashesByName[desc.fqName] = desc.dimHash
			}
		}
	}
	// Did anything happen at all?
	if len(newDescIDs) == 0 {
		return errors.New("collector has no descriptors")
	}
	if existing, exists := r.collectorsByID[collectorID]; exists {
		return AlreadyRegisteredError{
			ExistingCollector: existing,
			NewCollector:      c,
		}
	}
	// If the collectorID is new, but at least one of the descs existed
	// before, we are in trouble.
	if duplicateDescErr != nil {
		return duplicateDescErr
	}

	// Only after all tests have passed, actually register.
	r.collectorsByID[collectorID] = c
	for hash := range newDescIDs {
		r.descIDs[hash] = struct{}{}
	}
	for name, dimHash := range newDimHashesByName {
		r.dimHashesByName[name] = dimHash
	}
	return nil
}

// Unregister implements Registerer.
func (r *Registry) Unregister(c Collector) bool {
	var (
		descChan    = make(chan *Desc, capDescChan)
		descIDs     = map[uint64]struct{}{}
		collectorID uint64 // Just a sum of the desc IDs.
	)
	go func() {
		c.Describe(descChan)
		close(descChan)
	}()
	for desc := range descChan {
		if _, exists := descIDs[desc.id]; !exists {
			collectorID += desc.id
			descIDs[desc.id] = struct{}{}
		}
	}

	r.mtx.RLock()
	if _, exists := r.collectorsByID[collectorID]; !exists {
		r.mtx.RUnlock()
		return false
	}
	r.mtx.RUnlock()

	r.mtx.Lock()
	defer r.mtx.Unlock()

	delete(r.collectorsByID, collectorID)
	for id := range descIDs {
		delete(r.descIDs, id)
	}
	// dimHashesByName is left untouched as those must be consistent
	// throughout the lifetime of a program.
	return true
}

// MustRegister implements Registerer.
func (r *Registry) MustRegister(cs ...Collector) {
	for _, c := range cs {
		if err := r.Register(c); err != nil {
			panic(err)
		}
	}
}

// Gather implements Gatherer.
func (r *Registry) Gather() ([]*dto.MetricFamily, error) {
	var (
		metricChan        = make(chan Metric, capMetricChan)
		metricHashes      = map[uint64]struct{}{}
		dimHashes         = map[string]uint64{}
		wg                sync.WaitGroup
		errs              MultiError          // The collected errors to return in the end.
		registeredDescIDs map[uint64]struct{} // Only used for pedantic checks
	)

	r.mtx.RLock()
	metricFamiliesByName := make(map[string]*dto.MetricFamily, len(r.dimHashesByName))

	// Scatter.
	// (Collectors could be complex and slow, so we call them all at once.)
	wg.Add(len(r.collectorsByID))
	go func() {
		wg.Wait()
		close(metricChan)
	}()
	for _, collector := range r.collectorsByID {
		go func(collector Collector) {
			defer wg.Done()
			collector.Collect(metricChan)
		}(collector)
	}

	// In case pedantic checks are enabled, we have to copy the map before
	// giving up the RLock.
	if r.pedanticChecksEnabled {
		registeredDescIDs = make(map[uint64]struct{}, len(r.descIDs))
		for id := range r.descIDs {
			registeredDescIDs[id] = struct{}{}
		}