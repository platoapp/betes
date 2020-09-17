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
	"errors"
	"fmt"
	"math"
	"sort"
	"sync/atomic"

	dto "github.com/prometheus/client_model/go"

	"github.com/golang/protobuf/proto"
)

// ValueType is an enumeration of metric types that represent a simple value.
type ValueType int

// Possible values for the ValueType enum.
const (
	_ ValueType = iota
	CounterValue
	GaugeValue
	UntypedValue
)

var errInconsistentCardinality = errors.New("inconsistent label cardinality")

// value is a generic metric for simple values. It implements Metric, Collector,
// Counter, Gauge, and Untyped. Its effective type is determined by
// ValueType. This is a low-level building block used by the library to back the
// implementations of Counter, Gauge, and Untyped.
type value struct {
	// valBits containst the bits of the represented float64 value. It has
	// to go first in the struct to guarantee alignment for atomic
	// operations.  http://golang.org/pkg/sync/atomic/#pkg-note-BUG
	valBits uint64

	selfCollector

	desc       *Desc
	valType    ValueType
	labelPairs []*dto.LabelPair
}

// newValue returns a newly allocated value with the given Desc, ValueType,
// sample value and label values. It panics if the number of label
// values is different from the number of variable labels in Desc.
func newValue(desc *Desc, valueType ValueType, val float64, labelValues ...string) *value {
	if len(labelValues) != len(desc.variableLabels) {
		panic(errInconsistentCardinality)
	}
	result := &value{
		desc:       desc,
		valType:    valueType,
		valBits:    math.Float64bits(val),
		labelPairs: makeLabelPairs(desc, labelValues),
	}
	result.init(result)
	return result
}

func (v *value) Desc() *Desc {
	return v.desc
}

func (v *value) Set(val float64) {
	atomic.StoreUint64(&v.valBits, math.Float64bits(val))
}

func (v *value) Inc() {
	v.Add(1)
}

func (v *value) Dec() {
	v.Add(-1)
}

func (v *value) Add(val float64) {
	for {
		oldBits := atomic.LoadUint64(&v.valBits)
		newBits := math.Float64bits(math.Float64frombits(oldBits) + val)
		if atomic.CompareAndSwapUint64(&v.valBits, oldBits, newBits) {
			return
		}
	}
}

func (v *value) Sub(val float64) {
	v.Add(val * -1)
}

func (v *value) Write(out *dto.Metric) error {
	val := math.Float64frombits(atomic.LoadUint64(&v.valBits))
	return populateMetric(v.valType, val, v.labelPairs, out)
}

// valueFunc is a generic metric for simple values retrieved on collect time
// from a function. It implements Metric and Collector. Its effective type is
// determined by ValueType. This is a low-level building block used by the
// library to back the implementations of CounterFunc, GaugeFunc, and
// UntypedFunc.
type valueFunc struct {
	selfCollector

	desc       *Desc
	valType    ValueType
	function   func() float64
	labelPairs []*dto.LabelPair