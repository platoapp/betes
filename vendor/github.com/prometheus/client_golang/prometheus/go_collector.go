package prometheus

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"time"
)

type goCollector struct {
	goroutines Gauge
	gcDesc     *Desc

	// metrics to describe and collect
	metrics memStatsMetrics
}

// NewGoCollector returns a collector which exports metrics about the current
// go process.
func NewGoCollector() Collector {
	return &goCollector{
		goroutines: NewGauge(GaugeOpts{
			Namespace: "go",
			Name:      "goroutines",
			Help:      "Number of goroutines that currently exist.",
		}),
		gcDesc: NewDesc(
			"go_gc_duration_seconds",
			"A summary of the GC invocation durations.",
			nil, nil),
		metrics: memStatsMetrics{
			{
				desc: NewDesc(
					memstatNamespace("alloc_bytes"),
					"Number of bytes allocated and still in use.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return float64(ms.Alloc) },
				valType: GaugeValue,
			}, {
				desc: NewDesc(
					memstatNamespace("alloc_bytes_total"),
					"Total number of bytes allocated, even if freed.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return float64(ms.TotalAlloc) },
				valType: CounterValue,
			}, {
				desc: NewDesc(
					memstatNamespace("sys_bytes"),
					"Number of bytes obtained by system. Sum of all system allocations.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return float64(ms.Sys) },
				valType: GaugeValue,
			}, {
				desc: NewDesc(
					memstatNamespace("lookups_total"),
					"Total number of pointer lookups.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return float64(ms.Lookups) },
				valType: CounterValue,
			}, {
				desc: NewDesc(
					memstatNamespace("mallocs_total"),
					"Total number of mallocs.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return float64(ms.Mallocs) },
				valType: CounterValue,
			}, {
				desc: NewDesc(
					memstatNamespace("frees_total"),
					"Total number of frees.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return float64(ms.Frees) },
				valType: CounterValue,
			}, {
				desc: NewDesc(
					memstatNamespace("heap_alloc_bytes"),
					"Number of heap bytes allocated and still in use.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return float64(ms.HeapAlloc) },
				valType: GaugeValue,
			}, {
				desc: NewDesc(
					memstatNamespace("heap_sys_bytes"),
					"Number of heap bytes obtained from system.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return float64(ms.HeapSys) },
				valType: GaugeValue,
			}, {
				desc: NewDesc(
					memstatNamespace("heap_idle_bytes"),
					"Number of heap bytes waiting to be used.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return float64(ms.HeapIdle) },
				valType: GaugeValue,
	