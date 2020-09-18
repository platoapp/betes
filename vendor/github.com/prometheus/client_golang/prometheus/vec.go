
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
	"fmt"
	"sync"

	"github.com/prometheus/common/model"
)

// MetricVec is a Collector to bundle metrics of the same name that
// differ in their label values. MetricVec is usually not used directly but as a
// building block for implementations of vectors of a given metric
// type. GaugeVec, CounterVec, SummaryVec, and UntypedVec are examples already
// provided in this package.
type MetricVec struct {
	mtx      sync.RWMutex // Protects the children.
	children map[uint64][]metricWithLabelValues
	desc     *Desc

	newMetric   func(labelValues ...string) Metric
	hashAdd     func(h uint64, s string) uint64 // replace hash function for testing collision handling
	hashAddByte func(h uint64, b byte) uint64
}

// newMetricVec returns an initialized MetricVec. The concrete value is
// returned for embedding into another struct.
func newMetricVec(desc *Desc, newMetric func(lvs ...string) Metric) *MetricVec {
	return &MetricVec{
		children:    map[uint64][]metricWithLabelValues{},
		desc:        desc,
		newMetric:   newMetric,
		hashAdd:     hashAdd,
		hashAddByte: hashAddByte,
	}
}

// metricWithLabelValues provides the metric and its label values for
// disambiguation on hash collision.
type metricWithLabelValues struct {
	values []string
	metric Metric
}

// Describe implements Collector. The length of the returned slice
// is always one.
func (m *MetricVec) Describe(ch chan<- *Desc) {
	ch <- m.desc
}

// Collect implements Collector.
func (m *MetricVec) Collect(ch chan<- Metric) {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	for _, metrics := range m.children {
		for _, metric := range metrics {
			ch <- metric.metric
		}
	}
}

// GetMetricWithLabelValues returns the Metric for the given slice of label
// values (same order as the VariableLabels in Desc). If that combination of
// label values is accessed for the first time, a new Metric is created.
//
// It is possible to call this method without using the returned Metric to only
// create the new Metric but leave it at its start value (e.g. a Summary or
// Histogram without any observations). See also the SummaryVec example.
//
// Keeping the Metric for later use is possible (and should be considered if
// performance is critical), but keep in mind that Reset, DeleteLabelValues and
// Delete can be used to delete the Metric from the MetricVec. In that case, the
// Metric will still exist, but it will not be exported anymore, even if a
// Metric with the same label values is created later. See also the CounterVec
// example.
//
// An error is returned if the number of label values is not the same as the
// number of VariableLabels in Desc.
//
// Note that for more than one label value, this method is prone to mistakes
// caused by an incorrect order of arguments. Consider GetMetricWith(Labels) as
// an alternative to avoid that type of mistake. For higher label numbers, the
// latter has a much more readable (albeit more verbose) syntax, but it comes
// with a performance overhead (for creating and processing the Labels map).
// See also the GaugeVec example.
func (m *MetricVec) GetMetricWithLabelValues(lvs ...string) (Metric, error) {
	h, err := m.hashLabelValues(lvs)
	if err != nil {
		return nil, err
	}

	return m.getOrCreateMetricWithLabelValues(h, lvs), nil
}

// GetMetricWith returns the Metric for the given Labels map (the label names
// must match those of the VariableLabels in Desc). If that label map is
// accessed for the first time, a new Metric is created. Implications of
// creating a Metric without using it and keeping the Metric for later use are
// the same as for GetMetricWithLabelValues.