// Copyright 2018 The Prometheus Authors
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
	"runtime"
	"runtime/debug"
<<<<<<< HEAD
<<<<<<< HEAD
	"time"
)

func goRuntimeMemStats() memStatsMetrics {
	return memStatsMetrics{
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
				"Number of bytes obtained from system.",
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
		}, {
			desc: NewDesc(
				memstatNamespace("heap_inuse_bytes"),
				"Number of heap bytes that are in use.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.HeapInuse) },
			valType: GaugeValue,
		}, {
			desc: NewDesc(
				memstatNamespace("heap_released_bytes"),
				"Number of heap bytes released to OS.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.HeapReleased) },
			valType: GaugeValue,
		}, {
			desc: NewDesc(
				memstatNamespace("heap_objects"),
				"Number of allocated objects.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.HeapObjects) },
			valType: GaugeValue,
		}, {
			desc: NewDesc(
				memstatNamespace("stack_inuse_bytes"),
				"Number of bytes in use by the stack allocator.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.StackInuse) },
			valType: GaugeValue,
		}, {
			desc: NewDesc(
				memstatNamespace("stack_sys_bytes"),
				"Number of bytes obtained from system for stack allocator.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.StackSys) },
			valType: GaugeValue,
		}, {
			desc: NewDesc(
				memstatNamespace("mspan_inuse_bytes"),
				"Number of bytes in use by mspan structures.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.MSpanInuse) },
			valType: GaugeValue,
		}, {
			desc: NewDesc(
				memstatNamespace("mspan_sys_bytes"),
				"Number of bytes used for mspan structures obtained from system.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.MSpanSys) },
			valType: GaugeValue,
		}, {
			desc: NewDesc(
				memstatNamespace("mcache_inuse_bytes"),
				"Number of bytes in use by mcache structures.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.MCacheInuse) },
			valType: GaugeValue,
		}, {
			desc: NewDesc(
				memstatNamespace("mcache_sys_bytes"),
				"Number of bytes used for mcache structures obtained from system.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.MCacheSys) },
			valType: GaugeValue,
		}, {
			desc: NewDesc(
				memstatNamespace("buck_hash_sys_bytes"),
				"Number of bytes used by the profiling bucket hash table.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.BuckHashSys) },
			valType: GaugeValue,
		}, {
			desc: NewDesc(
				memstatNamespace("gc_sys_bytes"),
				"Number of bytes used for garbage collection system metadata.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.GCSys) },
			valType: GaugeValue,
		}, {
			desc: NewDesc(
				memstatNamespace("other_sys_bytes"),
				"Number of bytes used for other system allocations.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.OtherSys) },
			valType: GaugeValue,
		}, {
			desc: NewDesc(
				memstatNamespace("next_gc_bytes"),
				"Number of heap bytes when next garbage collection will take place.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.NextGC) },
			valType: GaugeValue,
		},
	}
}

type baseGoCollector struct {
	goroutinesDesc *Desc
	threadsDesc    *Desc
	gcDesc         *Desc
	gcLastTimeDesc *Desc
	goInfoDesc     *Desc
}

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
// NewGoCollector is the obsolete version of collectors.NewGoCollector.
// See there for documentation.
//
// Deprecated: Use collectors.NewGoCollector instead.
func NewGoCollector() Collector {
	return &goCollector{
||||||| parent of e1cd8261c (UPSTREAM: <carry>: update vendored files v0.13.1)
// NewGoCollector is the obsolete version of collectors.NewGoCollector.
// See there for documentation.
//
// Deprecated: Use collectors.NewGoCollector instead.
func NewGoCollector() Collector {
	return &goCollector{
=======
func newBaseGoCollector() baseGoCollector {
	return baseGoCollector{
>>>>>>> e1cd8261c (UPSTREAM: <carry>: update vendored files v0.13.1)
		goroutinesDesc: NewDesc(
			"go_goroutines",
			"Number of goroutines that currently exist.",
			nil, nil),
		threadsDesc: NewDesc(
			"go_threads",
			"Number of OS threads created.",
			nil, nil),
		gcDesc: NewDesc(
			"go_gc_duration_seconds",
			"A summary of the pause duration of garbage collection cycles.",
			nil, nil),
		gcLastTimeDesc: NewDesc(
			memstatNamespace("last_gc_time_seconds"),
			"Number of seconds since 1970 of last garbage collection.",
			nil, nil),
		goInfoDesc: NewDesc(
			"go_info",
			"Information about the Go environment.",
			nil, Labels{"version": runtime.Version()}),
	}
}

// Describe returns all descriptions of the collector.
func (c *baseGoCollector) Describe(ch chan<- *Desc) {
	ch <- c.goroutinesDesc
	ch <- c.threadsDesc
	ch <- c.gcDesc
	ch <- c.gcLastTimeDesc
	ch <- c.goInfoDesc
}

// Collect returns the current state of all metrics of the collector.
func (c *baseGoCollector) Collect(ch chan<- Metric) {
	ch <- MustNewConstMetric(c.goroutinesDesc, GaugeValue, float64(runtime.NumGoroutine()))
	n, _ := runtime.ThreadCreateProfile(nil)
	ch <- MustNewConstMetric(c.threadsDesc, GaugeValue, float64(n))

	var stats debug.GCStats
	stats.PauseQuantiles = make([]time.Duration, 5)
	debug.ReadGCStats(&stats)

	quantiles := make(map[float64]float64)
	for idx, pq := range stats.PauseQuantiles[1:] {
		quantiles[float64(idx+1)/float64(len(stats.PauseQuantiles)-1)] = pq.Seconds()
	}
	quantiles[0.0] = stats.PauseQuantiles[0].Seconds()
	ch <- MustNewConstSummary(c.gcDesc, uint64(stats.NumGC), stats.PauseTotal.Seconds(), quantiles)
	ch <- MustNewConstMetric(c.gcLastTimeDesc, GaugeValue, float64(stats.LastGC.UnixNano())/1e9)
	ch <- MustNewConstMetric(c.goInfoDesc, GaugeValue, 1)
}

func memstatNamespace(s string) string {
	return "go_memstats_" + s
}

// memStatsMetrics provide description, evaluator, runtime/metrics name, and
// value type for memstat metrics.
// TODO(bwplotka): Remove with end Go 1.16 EOL and replace with runtime/metrics.Description
type memStatsMetrics []struct {
	desc    *Desc
	eval    func(*runtime.MemStats) float64
	valType ValueType
}
<<<<<<< HEAD

// NewBuildInfoCollector is the obsolete version of collectors.NewBuildInfoCollector.
// See there for documentation.
//
// Deprecated: Use collectors.NewBuildInfoCollector instead.
func NewBuildInfoCollector() Collector {
	path, version, sum := "unknown", "unknown", "unknown"
	if bi, ok := debug.ReadBuildInfo(); ok {
		path = bi.Main.Path
		version = bi.Main.Version
		sum = bi.Main.Sum
	}
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// NewGoCollector returns a collector that exports metrics about the current Go
// process. This includes memory stats. To collect those, runtime.ReadMemStats
// is called. This requires to “stop the world”, which usually only happens for
// garbage collection (GC). Take the following implications into account when
// deciding whether to use the Go collector:
||||||| parent of 5ce8c7613 (update vendored files)
// NewGoCollector returns a collector that exports metrics about the current Go
// process. This includes memory stats. To collect those, runtime.ReadMemStats
// is called. This requires to “stop the world”, which usually only happens for
// garbage collection (GC). Take the following implications into account when
// deciding whether to use the Go collector:
=======
// NewGoCollector is the obsolete version of collectors.NewGoCollector.
// See there for documentation.
>>>>>>> 5ce8c7613 (update vendored files)
//
// Deprecated: Use collectors.NewGoCollector instead.
func NewGoCollector() Collector {
	return &goCollector{
		goroutinesDesc: NewDesc(
			"go_goroutines",
			"Number of goroutines that currently exist.",
			nil, nil),
		threadsDesc: NewDesc(
			"go_threads",
			"Number of OS threads created.",
			nil, nil),
		gcDesc: NewDesc(
			"go_gc_duration_seconds",
			"A summary of the pause duration of garbage collection cycles.",
			nil, nil),
		goInfoDesc: NewDesc(
			"go_info",
			"Information about the Go environment.",
			nil, Labels{"version": runtime.Version()}),
		msLast:    &runtime.MemStats{},
		msRead:    runtime.ReadMemStats,
		msMaxWait: time.Second,
		msMaxAge:  5 * time.Minute,
		msMetrics: memStatsMetrics{
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
					"Number of bytes obtained from system.",
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
			}, {
				desc: NewDesc(
					memstatNamespace("heap_inuse_bytes"),
					"Number of heap bytes that are in use.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return float64(ms.HeapInuse) },
				valType: GaugeValue,
			}, {
				desc: NewDesc(
					memstatNamespace("heap_released_bytes"),
					"Number of heap bytes released to OS.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return float64(ms.HeapReleased) },
				valType: GaugeValue,
			}, {
				desc: NewDesc(
					memstatNamespace("heap_objects"),
					"Number of allocated objects.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return float64(ms.HeapObjects) },
				valType: GaugeValue,
			}, {
				desc: NewDesc(
					memstatNamespace("stack_inuse_bytes"),
					"Number of bytes in use by the stack allocator.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return float64(ms.StackInuse) },
				valType: GaugeValue,
			}, {
				desc: NewDesc(
					memstatNamespace("stack_sys_bytes"),
					"Number of bytes obtained from system for stack allocator.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return float64(ms.StackSys) },
				valType: GaugeValue,
			}, {
				desc: NewDesc(
					memstatNamespace("mspan_inuse_bytes"),
					"Number of bytes in use by mspan structures.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return float64(ms.MSpanInuse) },
				valType: GaugeValue,
			}, {
				desc: NewDesc(
					memstatNamespace("mspan_sys_bytes"),
					"Number of bytes used for mspan structures obtained from system.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return float64(ms.MSpanSys) },
				valType: GaugeValue,
			}, {
				desc: NewDesc(
					memstatNamespace("mcache_inuse_bytes"),
					"Number of bytes in use by mcache structures.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return float64(ms.MCacheInuse) },
				valType: GaugeValue,
			}, {
				desc: NewDesc(
					memstatNamespace("mcache_sys_bytes"),
					"Number of bytes used for mcache structures obtained from system.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return float64(ms.MCacheSys) },
				valType: GaugeValue,
			}, {
				desc: NewDesc(
					memstatNamespace("buck_hash_sys_bytes"),
					"Number of bytes used by the profiling bucket hash table.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return float64(ms.BuckHashSys) },
				valType: GaugeValue,
			}, {
				desc: NewDesc(
					memstatNamespace("gc_sys_bytes"),
					"Number of bytes used for garbage collection system metadata.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return float64(ms.GCSys) },
				valType: GaugeValue,
			}, {
				desc: NewDesc(
					memstatNamespace("other_sys_bytes"),
					"Number of bytes used for other system allocations.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return float64(ms.OtherSys) },
				valType: GaugeValue,
			}, {
				desc: NewDesc(
					memstatNamespace("next_gc_bytes"),
					"Number of heap bytes when next garbage collection will take place.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return float64(ms.NextGC) },
				valType: GaugeValue,
			}, {
				desc: NewDesc(
					memstatNamespace("last_gc_time_seconds"),
					"Number of seconds since 1970 of last garbage collection.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return float64(ms.LastGC) / 1e9 },
				valType: GaugeValue,
			}, {
				desc: NewDesc(
					memstatNamespace("gc_cpu_fraction"),
					"The fraction of this program's available CPU time used by the GC since the program started.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return ms.GCCPUFraction },
				valType: GaugeValue,
			},
		},
	}
}

func memstatNamespace(s string) string {
	return "go_memstats_" + s
}

// Describe returns all descriptions of the collector.
func (c *goCollector) Describe(ch chan<- *Desc) {
	ch <- c.goroutinesDesc
	ch <- c.threadsDesc
	ch <- c.gcDesc
	ch <- c.goInfoDesc
	for _, i := range c.msMetrics {
		ch <- i.desc
	}
}

// Collect returns the current state of all metrics of the collector.
func (c *goCollector) Collect(ch chan<- Metric) {
	var (
		ms   = &runtime.MemStats{}
		done = make(chan struct{})
	)
	// Start reading memstats first as it might take a while.
	go func() {
		c.msRead(ms)
		c.msMtx.Lock()
		c.msLast = ms
		c.msLastTimestamp = time.Now()
		c.msMtx.Unlock()
		close(done)
	}()

	ch <- MustNewConstMetric(c.goroutinesDesc, GaugeValue, float64(runtime.NumGoroutine()))
	n, _ := runtime.ThreadCreateProfile(nil)
	ch <- MustNewConstMetric(c.threadsDesc, GaugeValue, float64(n))

	var stats debug.GCStats
	stats.PauseQuantiles = make([]time.Duration, 5)
	debug.ReadGCStats(&stats)

	quantiles := make(map[float64]float64)
	for idx, pq := range stats.PauseQuantiles[1:] {
		quantiles[float64(idx+1)/float64(len(stats.PauseQuantiles)-1)] = pq.Seconds()
	}
	quantiles[0.0] = stats.PauseQuantiles[0].Seconds()
	ch <- MustNewConstSummary(c.gcDesc, uint64(stats.NumGC), stats.PauseTotal.Seconds(), quantiles)

	ch <- MustNewConstMetric(c.goInfoDesc, GaugeValue, 1)

	timer := time.NewTimer(c.msMaxWait)
	select {
	case <-done: // Our own ReadMemStats succeeded in time. Use it.
		timer.Stop() // Important for high collection frequencies to not pile up timers.
		c.msCollect(ch, ms)
		return
	case <-timer.C: // Time out, use last memstats if possible. Continue below.
	}
	c.msMtx.Lock()
	if time.Since(c.msLastTimestamp) < c.msMaxAge {
		// Last memstats are recent enough. Collect from them under the lock.
		c.msCollect(ch, c.msLast)
		c.msMtx.Unlock()
		return
	}
	// If we are here, the last memstats are too old or don't exist. We have
	// to wait until our own ReadMemStats finally completes. For that to
	// happen, we have to release the lock.
	c.msMtx.Unlock()
	<-done
	c.msCollect(ch, ms)
}

func (c *goCollector) msCollect(ch chan<- Metric, ms *runtime.MemStats) {
	for _, i := range c.msMetrics {
		ch <- MustNewConstMetric(i.desc, i.valType, i.eval(ms))
	}
}

// memStatsMetrics provide description, value, and value type for memstat metrics.
type memStatsMetrics []struct {
	desc    *Desc
	eval    func(*runtime.MemStats) float64
	valType ValueType
}

// NewBuildInfoCollector is the obsolete version of collectors.NewBuildInfoCollector.
// See there for documentation.
//
// Deprecated: Use collectors.NewBuildInfoCollector instead.
func NewBuildInfoCollector() Collector {
<<<<<<< HEAD
	path, version, sum := readBuildInfo()
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
	path, version, sum := readBuildInfo()
=======
	path, version, sum := "unknown", "unknown", "unknown"
	if bi, ok := debug.ReadBuildInfo(); ok {
		path = bi.Main.Path
		version = bi.Main.Version
		sum = bi.Main.Sum
	}
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// NewGoCollector returns a collector that exports metrics about the current Go
// process. This includes memory stats. To collect those, runtime.ReadMemStats
// is called. This requires to “stop the world”, which usually only happens for
// garbage collection (GC). Take the following implications into account when
// deciding whether to use the Go collector:
||||||| parent of 6b7ce455e (update vendored files)
// NewGoCollector returns a collector that exports metrics about the current Go
// process. This includes memory stats. To collect those, runtime.ReadMemStats
// is called. This requires to “stop the world”, which usually only happens for
// garbage collection (GC). Take the following implications into account when
// deciding whether to use the Go collector:
=======
// NewGoCollector is the obsolete version of collectors.NewGoCollector.
// See there for documentation.
>>>>>>> 6b7ce455e (update vendored files)
//
// Deprecated: Use collectors.NewGoCollector instead.
func NewGoCollector() Collector {
	return &goCollector{
		goroutinesDesc: NewDesc(
			"go_goroutines",
			"Number of goroutines that currently exist.",
			nil, nil),
		threadsDesc: NewDesc(
			"go_threads",
			"Number of OS threads created.",
			nil, nil),
		gcDesc: NewDesc(
			"go_gc_duration_seconds",
			"A summary of the pause duration of garbage collection cycles.",
			nil, nil),
		goInfoDesc: NewDesc(
			"go_info",
			"Information about the Go environment.",
			nil, Labels{"version": runtime.Version()}),
		msLast:    &runtime.MemStats{},
		msRead:    runtime.ReadMemStats,
		msMaxWait: time.Second,
		msMaxAge:  5 * time.Minute,
		msMetrics: memStatsMetrics{
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
					"Number of bytes obtained from system.",
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
			}, {
				desc: NewDesc(
					memstatNamespace("heap_inuse_bytes"),
					"Number of heap bytes that are in use.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return float64(ms.HeapInuse) },
				valType: GaugeValue,
			}, {
				desc: NewDesc(
					memstatNamespace("heap_released_bytes"),
					"Number of heap bytes released to OS.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return float64(ms.HeapReleased) },
				valType: GaugeValue,
			}, {
				desc: NewDesc(
					memstatNamespace("heap_objects"),
					"Number of allocated objects.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return float64(ms.HeapObjects) },
				valType: GaugeValue,
			}, {
				desc: NewDesc(
					memstatNamespace("stack_inuse_bytes"),
					"Number of bytes in use by the stack allocator.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return float64(ms.StackInuse) },
				valType: GaugeValue,
			}, {
				desc: NewDesc(
					memstatNamespace("stack_sys_bytes"),
					"Number of bytes obtained from system for stack allocator.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return float64(ms.StackSys) },
				valType: GaugeValue,
			}, {
				desc: NewDesc(
					memstatNamespace("mspan_inuse_bytes"),
					"Number of bytes in use by mspan structures.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return float64(ms.MSpanInuse) },
				valType: GaugeValue,
			}, {
				desc: NewDesc(
					memstatNamespace("mspan_sys_bytes"),
					"Number of bytes used for mspan structures obtained from system.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return float64(ms.MSpanSys) },
				valType: GaugeValue,
			}, {
				desc: NewDesc(
					memstatNamespace("mcache_inuse_bytes"),
					"Number of bytes in use by mcache structures.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return float64(ms.MCacheInuse) },
				valType: GaugeValue,
			}, {
				desc: NewDesc(
					memstatNamespace("mcache_sys_bytes"),
					"Number of bytes used for mcache structures obtained from system.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return float64(ms.MCacheSys) },
				valType: GaugeValue,
			}, {
				desc: NewDesc(
					memstatNamespace("buck_hash_sys_bytes"),
					"Number of bytes used by the profiling bucket hash table.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return float64(ms.BuckHashSys) },
				valType: GaugeValue,
			}, {
				desc: NewDesc(
					memstatNamespace("gc_sys_bytes"),
					"Number of bytes used for garbage collection system metadata.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return float64(ms.GCSys) },
				valType: GaugeValue,
			}, {
				desc: NewDesc(
					memstatNamespace("other_sys_bytes"),
					"Number of bytes used for other system allocations.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return float64(ms.OtherSys) },
				valType: GaugeValue,
			}, {
				desc: NewDesc(
					memstatNamespace("next_gc_bytes"),
					"Number of heap bytes when next garbage collection will take place.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return float64(ms.NextGC) },
				valType: GaugeValue,
			}, {
				desc: NewDesc(
					memstatNamespace("last_gc_time_seconds"),
					"Number of seconds since 1970 of last garbage collection.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return float64(ms.LastGC) / 1e9 },
				valType: GaugeValue,
			}, {
				desc: NewDesc(
					memstatNamespace("gc_cpu_fraction"),
					"The fraction of this program's available CPU time used by the GC since the program started.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return ms.GCCPUFraction },
				valType: GaugeValue,
			},
		},
	}
}

func memstatNamespace(s string) string {
	return "go_memstats_" + s
}

// Describe returns all descriptions of the collector.
func (c *goCollector) Describe(ch chan<- *Desc) {
	ch <- c.goroutinesDesc
	ch <- c.threadsDesc
	ch <- c.gcDesc
	ch <- c.goInfoDesc
	for _, i := range c.msMetrics {
		ch <- i.desc
	}
}

// Collect returns the current state of all metrics of the collector.
func (c *goCollector) Collect(ch chan<- Metric) {
	var (
		ms   = &runtime.MemStats{}
		done = make(chan struct{})
	)
	// Start reading memstats first as it might take a while.
	go func() {
		c.msRead(ms)
		c.msMtx.Lock()
		c.msLast = ms
		c.msLastTimestamp = time.Now()
		c.msMtx.Unlock()
		close(done)
	}()

	ch <- MustNewConstMetric(c.goroutinesDesc, GaugeValue, float64(runtime.NumGoroutine()))
	n, _ := runtime.ThreadCreateProfile(nil)
	ch <- MustNewConstMetric(c.threadsDesc, GaugeValue, float64(n))

	var stats debug.GCStats
	stats.PauseQuantiles = make([]time.Duration, 5)
	debug.ReadGCStats(&stats)

	quantiles := make(map[float64]float64)
	for idx, pq := range stats.PauseQuantiles[1:] {
		quantiles[float64(idx+1)/float64(len(stats.PauseQuantiles)-1)] = pq.Seconds()
	}
	quantiles[0.0] = stats.PauseQuantiles[0].Seconds()
	ch <- MustNewConstSummary(c.gcDesc, uint64(stats.NumGC), stats.PauseTotal.Seconds(), quantiles)

	ch <- MustNewConstMetric(c.goInfoDesc, GaugeValue, 1)

	timer := time.NewTimer(c.msMaxWait)
	select {
	case <-done: // Our own ReadMemStats succeeded in time. Use it.
		timer.Stop() // Important for high collection frequencies to not pile up timers.
		c.msCollect(ch, ms)
		return
	case <-timer.C: // Time out, use last memstats if possible. Continue below.
	}
	c.msMtx.Lock()
	if time.Since(c.msLastTimestamp) < c.msMaxAge {
		// Last memstats are recent enough. Collect from them under the lock.
		c.msCollect(ch, c.msLast)
		c.msMtx.Unlock()
		return
	}
	// If we are here, the last memstats are too old or don't exist. We have
	// to wait until our own ReadMemStats finally completes. For that to
	// happen, we have to release the lock.
	c.msMtx.Unlock()
	<-done
	c.msCollect(ch, ms)
}

func (c *goCollector) msCollect(ch chan<- Metric, ms *runtime.MemStats) {
	for _, i := range c.msMetrics {
		ch <- MustNewConstMetric(i.desc, i.valType, i.eval(ms))
	}
}

// memStatsMetrics provide description, value, and value type for memstat metrics.
type memStatsMetrics []struct {
	desc    *Desc
	eval    func(*runtime.MemStats) float64
	valType ValueType
}

// NewBuildInfoCollector is the obsolete version of collectors.NewBuildInfoCollector.
// See there for documentation.
//
// Deprecated: Use collectors.NewBuildInfoCollector instead.
func NewBuildInfoCollector() Collector {
<<<<<<< HEAD
	path, version, sum := readBuildInfo()
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
	path, version, sum := readBuildInfo()
=======
	path, version, sum := "unknown", "unknown", "unknown"
	if bi, ok := debug.ReadBuildInfo(); ok {
		path = bi.Main.Path
		version = bi.Main.Version
		sum = bi.Main.Sum
	}
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// NewGoCollector returns a collector that exports metrics about the current Go
// process. This includes memory stats. To collect those, runtime.ReadMemStats
// is called. This requires to “stop the world”, which usually only happens for
// garbage collection (GC). Take the following implications into account when
// deciding whether to use the Go collector:
||||||| parent of 4d7e5ad26 (update vendored files)
// NewGoCollector returns a collector that exports metrics about the current Go
// process. This includes memory stats. To collect those, runtime.ReadMemStats
// is called. This requires to “stop the world”, which usually only happens for
// garbage collection (GC). Take the following implications into account when
// deciding whether to use the Go collector:
=======
// NewGoCollector is the obsolete version of collectors.NewGoCollector.
// See there for documentation.
>>>>>>> 4d7e5ad26 (update vendored files)
//
// Deprecated: Use collectors.NewGoCollector instead.
func NewGoCollector() Collector {
	return &goCollector{
		goroutinesDesc: NewDesc(
			"go_goroutines",
			"Number of goroutines that currently exist.",
			nil, nil),
		threadsDesc: NewDesc(
			"go_threads",
			"Number of OS threads created.",
			nil, nil),
		gcDesc: NewDesc(
			"go_gc_duration_seconds",
			"A summary of the pause duration of garbage collection cycles.",
			nil, nil),
		goInfoDesc: NewDesc(
			"go_info",
			"Information about the Go environment.",
			nil, Labels{"version": runtime.Version()}),
		msLast:    &runtime.MemStats{},
		msRead:    runtime.ReadMemStats,
		msMaxWait: time.Second,
		msMaxAge:  5 * time.Minute,
		msMetrics: memStatsMetrics{
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
					"Number of bytes obtained from system.",
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
			}, {
				desc: NewDesc(
					memstatNamespace("heap_inuse_bytes"),
					"Number of heap bytes that are in use.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return float64(ms.HeapInuse) },
				valType: GaugeValue,
			}, {
				desc: NewDesc(
					memstatNamespace("heap_released_bytes"),
					"Number of heap bytes released to OS.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return float64(ms.HeapReleased) },
				valType: GaugeValue,
			}, {
				desc: NewDesc(
					memstatNamespace("heap_objects"),
					"Number of allocated objects.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return float64(ms.HeapObjects) },
				valType: GaugeValue,
			}, {
				desc: NewDesc(
					memstatNamespace("stack_inuse_bytes"),
					"Number of bytes in use by the stack allocator.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return float64(ms.StackInuse) },
				valType: GaugeValue,
			}, {
				desc: NewDesc(
					memstatNamespace("stack_sys_bytes"),
					"Number of bytes obtained from system for stack allocator.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return float64(ms.StackSys) },
				valType: GaugeValue,
			}, {
				desc: NewDesc(
					memstatNamespace("mspan_inuse_bytes"),
					"Number of bytes in use by mspan structures.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return float64(ms.MSpanInuse) },
				valType: GaugeValue,
			}, {
				desc: NewDesc(
					memstatNamespace("mspan_sys_bytes"),
					"Number of bytes used for mspan structures obtained from system.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return float64(ms.MSpanSys) },
				valType: GaugeValue,
			}, {
				desc: NewDesc(
					memstatNamespace("mcache_inuse_bytes"),
					"Number of bytes in use by mcache structures.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return float64(ms.MCacheInuse) },
				valType: GaugeValue,
			}, {
				desc: NewDesc(
					memstatNamespace("mcache_sys_bytes"),
					"Number of bytes used for mcache structures obtained from system.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return float64(ms.MCacheSys) },
				valType: GaugeValue,
			}, {
				desc: NewDesc(
					memstatNamespace("buck_hash_sys_bytes"),
					"Number of bytes used by the profiling bucket hash table.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return float64(ms.BuckHashSys) },
				valType: GaugeValue,
			}, {
				desc: NewDesc(
					memstatNamespace("gc_sys_bytes"),
					"Number of bytes used for garbage collection system metadata.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return float64(ms.GCSys) },
				valType: GaugeValue,
			}, {
				desc: NewDesc(
					memstatNamespace("other_sys_bytes"),
					"Number of bytes used for other system allocations.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return float64(ms.OtherSys) },
				valType: GaugeValue,
			}, {
				desc: NewDesc(
					memstatNamespace("next_gc_bytes"),
					"Number of heap bytes when next garbage collection will take place.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return float64(ms.NextGC) },
				valType: GaugeValue,
			}, {
				desc: NewDesc(
					memstatNamespace("last_gc_time_seconds"),
					"Number of seconds since 1970 of last garbage collection.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return float64(ms.LastGC) / 1e9 },
				valType: GaugeValue,
			}, {
				desc: NewDesc(
					memstatNamespace("gc_cpu_fraction"),
					"The fraction of this program's available CPU time used by the GC since the program started.",
					nil, nil,
				),
				eval:    func(ms *runtime.MemStats) float64 { return ms.GCCPUFraction },
				valType: GaugeValue,
			},
		},
	}
}

func memstatNamespace(s string) string {
	return "go_memstats_" + s
}

// Describe returns all descriptions of the collector.
func (c *goCollector) Describe(ch chan<- *Desc) {
	ch <- c.goroutinesDesc
	ch <- c.threadsDesc
	ch <- c.gcDesc
	ch <- c.goInfoDesc
	for _, i := range c.msMetrics {
		ch <- i.desc
	}
}

// Collect returns the current state of all metrics of the collector.
func (c *goCollector) Collect(ch chan<- Metric) {
	var (
		ms   = &runtime.MemStats{}
		done = make(chan struct{})
	)
	// Start reading memstats first as it might take a while.
	go func() {
		c.msRead(ms)
		c.msMtx.Lock()
		c.msLast = ms
		c.msLastTimestamp = time.Now()
		c.msMtx.Unlock()
		close(done)
	}()

	ch <- MustNewConstMetric(c.goroutinesDesc, GaugeValue, float64(runtime.NumGoroutine()))
	n, _ := runtime.ThreadCreateProfile(nil)
	ch <- MustNewConstMetric(c.threadsDesc, GaugeValue, float64(n))

	var stats debug.GCStats
	stats.PauseQuantiles = make([]time.Duration, 5)
	debug.ReadGCStats(&stats)

	quantiles := make(map[float64]float64)
	for idx, pq := range stats.PauseQuantiles[1:] {
		quantiles[float64(idx+1)/float64(len(stats.PauseQuantiles)-1)] = pq.Seconds()
	}
	quantiles[0.0] = stats.PauseQuantiles[0].Seconds()
	ch <- MustNewConstSummary(c.gcDesc, uint64(stats.NumGC), stats.PauseTotal.Seconds(), quantiles)

	ch <- MustNewConstMetric(c.goInfoDesc, GaugeValue, 1)

	timer := time.NewTimer(c.msMaxWait)
	select {
	case <-done: // Our own ReadMemStats succeeded in time. Use it.
		timer.Stop() // Important for high collection frequencies to not pile up timers.
		c.msCollect(ch, ms)
		return
	case <-timer.C: // Time out, use last memstats if possible. Continue below.
	}
	c.msMtx.Lock()
	if time.Since(c.msLastTimestamp) < c.msMaxAge {
		// Last memstats are recent enough. Collect from them under the lock.
		c.msCollect(ch, c.msLast)
		c.msMtx.Unlock()
		return
	}
	// If we are here, the last memstats are too old or don't exist. We have
	// to wait until our own ReadMemStats finally completes. For that to
	// happen, we have to release the lock.
	c.msMtx.Unlock()
	<-done
	c.msCollect(ch, ms)
}

func (c *goCollector) msCollect(ch chan<- Metric, ms *runtime.MemStats) {
	for _, i := range c.msMetrics {
		ch <- MustNewConstMetric(i.desc, i.valType, i.eval(ms))
	}
}

// memStatsMetrics provide description, value, and value type for memstat metrics.
type memStatsMetrics []struct {
	desc    *Desc
	eval    func(*runtime.MemStats) float64
	valType ValueType
}

// NewBuildInfoCollector is the obsolete version of collectors.NewBuildInfoCollector.
// See there for documentation.
//
// Deprecated: Use collectors.NewBuildInfoCollector instead.
func NewBuildInfoCollector() Collector {
<<<<<<< HEAD
	path, version, sum := readBuildInfo()
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
	path, version, sum := readBuildInfo()
=======
	path, version, sum := "unknown", "unknown", "unknown"
	if bi, ok := debug.ReadBuildInfo(); ok {
		path = bi.Main.Path
		version = bi.Main.Version
		sum = bi.Main.Sum
	}
>>>>>>> 4d7e5ad26 (update vendored files)
	c := &selfCollector{MustNewConstMetric(
		NewDesc(
			"go_build_info",
			"Build information about the main Go module.",
			nil, Labels{"path": path, "version": version, "checksum": sum},
		),
		GaugeValue, 1)}
	c.init(c.self)
	return c
}
||||||| parent of e1cd8261c (UPSTREAM: <carry>: update vendored files v0.13.1)

// NewBuildInfoCollector is the obsolete version of collectors.NewBuildInfoCollector.
// See there for documentation.
//
// Deprecated: Use collectors.NewBuildInfoCollector instead.
func NewBuildInfoCollector() Collector {
	path, version, sum := "unknown", "unknown", "unknown"
	if bi, ok := debug.ReadBuildInfo(); ok {
		path = bi.Main.Path
		version = bi.Main.Version
		sum = bi.Main.Sum
	}
	c := &selfCollector{MustNewConstMetric(
		NewDesc(
			"go_build_info",
			"Build information about the main Go module.",
			nil, Labels{"path": path, "version": version, "checksum": sum},
		),
		GaugeValue, 1)}
	c.init(c.self)
	return c
}
=======
>>>>>>> e1cd8261c (UPSTREAM: <carry>: update vendored files v0.13.1)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	"sync"
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	"sync"
=======
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	"time"
)

// goRuntimeMemStats provides the metrics initially provided by runtime.ReadMemStats.
// From Go 1.17 those similar (and better) statistics are provided by runtime/metrics, so
// while eval closure works on runtime.MemStats, the struct from Go 1.17+ is
// populated using runtime/metrics. Those are the defaults we can't alter.
func goRuntimeMemStats() memStatsMetrics {
	return memStatsMetrics{
		{
			desc: NewDesc(
				memstatNamespace("alloc_bytes"),
				"Number of bytes allocated in heap and currently in use. Equals to /memory/classes/heap/objects:bytes.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.Alloc) },
			valType: GaugeValue,
		}, {
			desc: NewDesc(
				memstatNamespace("alloc_bytes_total"),
				"Total number of bytes allocated in heap until now, even if released already. Equals to /gc/heap/allocs:bytes.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.TotalAlloc) },
			valType: CounterValue,
		}, {
			desc: NewDesc(
				memstatNamespace("sys_bytes"),
				"Number of bytes obtained from system. Equals to /memory/classes/total:byte.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.Sys) },
			valType: GaugeValue,
		}, {
			desc: NewDesc(
				memstatNamespace("mallocs_total"),
				// TODO(bwplotka): We could add go_memstats_heap_objects, probably useful for discovery. Let's gather more feedback, kind of a waste of bytes for everybody for compatibility reasons to keep both, and we can't really rename/remove useful metric.
				"Total number of heap objects allocated, both live and gc-ed. Semantically a counter version for go_memstats_heap_objects gauge. Equals to /gc/heap/allocs:objects + /gc/heap/tiny/allocs:objects.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.Mallocs) },
			valType: CounterValue,
		}, {
			desc: NewDesc(
				memstatNamespace("frees_total"),
				"Total number of heap objects frees. Equals to /gc/heap/frees:objects + /gc/heap/tiny/allocs:objects.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.Frees) },
			valType: CounterValue,
		}, {
			desc: NewDesc(
				memstatNamespace("heap_alloc_bytes"),
				"Number of heap bytes allocated and currently in use, same as go_memstats_alloc_bytes. Equals to /memory/classes/heap/objects:bytes.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.HeapAlloc) },
			valType: GaugeValue,
		}, {
			desc: NewDesc(
				memstatNamespace("heap_sys_bytes"),
				"Number of heap bytes obtained from system. Equals to /memory/classes/heap/objects:bytes + /memory/classes/heap/unused:bytes + /memory/classes/heap/released:bytes + /memory/classes/heap/free:bytes.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.HeapSys) },
			valType: GaugeValue,
		}, {
			desc: NewDesc(
				memstatNamespace("heap_idle_bytes"),
				"Number of heap bytes waiting to be used. Equals to /memory/classes/heap/released:bytes + /memory/classes/heap/free:bytes.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.HeapIdle) },
			valType: GaugeValue,
		}, {
			desc: NewDesc(
				memstatNamespace("heap_inuse_bytes"),
				"Number of heap bytes that are in use. Equals to /memory/classes/heap/objects:bytes + /memory/classes/heap/unused:bytes",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.HeapInuse) },
			valType: GaugeValue,
		}, {
			desc: NewDesc(
				memstatNamespace("heap_released_bytes"),
				"Number of heap bytes released to OS. Equals to /memory/classes/heap/released:bytes.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.HeapReleased) },
			valType: GaugeValue,
		}, {
			desc: NewDesc(
				memstatNamespace("heap_objects"),
				"Number of currently allocated objects. Equals to /gc/heap/objects:objects.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.HeapObjects) },
			valType: GaugeValue,
		}, {
			desc: NewDesc(
				memstatNamespace("stack_inuse_bytes"),
				"Number of bytes obtained from system for stack allocator in non-CGO environments. Equals to /memory/classes/heap/stacks:bytes.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.StackInuse) },
			valType: GaugeValue,
		}, {
			desc: NewDesc(
				memstatNamespace("stack_sys_bytes"),
				"Number of bytes obtained from system for stack allocator. Equals to /memory/classes/heap/stacks:bytes + /memory/classes/os-stacks:bytes.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.StackSys) },
			valType: GaugeValue,
		}, {
			desc: NewDesc(
				memstatNamespace("mspan_inuse_bytes"),
				"Number of bytes in use by mspan structures. Equals to /memory/classes/metadata/mspan/inuse:bytes.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.MSpanInuse) },
			valType: GaugeValue,
		}, {
			desc: NewDesc(
				memstatNamespace("mspan_sys_bytes"),
				"Number of bytes used for mspan structures obtained from system. Equals to /memory/classes/metadata/mspan/inuse:bytes + /memory/classes/metadata/mspan/free:bytes.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.MSpanSys) },
			valType: GaugeValue,
		}, {
			desc: NewDesc(
				memstatNamespace("mcache_inuse_bytes"),
				"Number of bytes in use by mcache structures. Equals to /memory/classes/metadata/mcache/inuse:bytes.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.MCacheInuse) },
			valType: GaugeValue,
		}, {
			desc: NewDesc(
				memstatNamespace("mcache_sys_bytes"),
				"Number of bytes used for mcache structures obtained from system. Equals to /memory/classes/metadata/mcache/inuse:bytes + /memory/classes/metadata/mcache/free:bytes.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.MCacheSys) },
			valType: GaugeValue,
		}, {
			desc: NewDesc(
				memstatNamespace("buck_hash_sys_bytes"),
				"Number of bytes used by the profiling bucket hash table. Equals to /memory/classes/profiling/buckets:bytes.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.BuckHashSys) },
			valType: GaugeValue,
		}, {
			desc: NewDesc(
				memstatNamespace("gc_sys_bytes"),
				"Number of bytes used for garbage collection system metadata. Equals to /memory/classes/metadata/other:bytes.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.GCSys) },
			valType: GaugeValue,
		}, {
			desc: NewDesc(
				memstatNamespace("other_sys_bytes"),
				"Number of bytes used for other system allocations. Equals to /memory/classes/other:bytes.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.OtherSys) },
			valType: GaugeValue,
		}, {
			desc: NewDesc(
				memstatNamespace("next_gc_bytes"),
				"Number of heap bytes when next garbage collection will take place. Equals to /gc/heap/goal:bytes.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.NextGC) },
			valType: GaugeValue,
		},
	}
}

type baseGoCollector struct {
	goroutinesDesc *Desc
	threadsDesc    *Desc
	gcDesc         *Desc
	gcLastTimeDesc *Desc
	goInfoDesc     *Desc
}

func newBaseGoCollector() baseGoCollector {
	return baseGoCollector{
		goroutinesDesc: NewDesc(
			"go_goroutines",
			"Number of goroutines that currently exist.",
			nil, nil),
		threadsDesc: NewDesc(
			"go_threads",
			"Number of OS threads created.",
			nil, nil),
		gcDesc: NewDesc(
			"go_gc_duration_seconds",
			"A summary of the wall-time pause (stop-the-world) duration in garbage collection cycles.",
			nil, nil),
		gcLastTimeDesc: NewDesc(
			"go_memstats_last_gc_time_seconds",
			"Number of seconds since 1970 of last garbage collection.",
			nil, nil),
		goInfoDesc: NewDesc(
			"go_info",
			"Information about the Go environment.",
			nil, Labels{"version": runtime.Version()}),
	}
}

// Describe returns all descriptions of the collector.
func (c *baseGoCollector) Describe(ch chan<- *Desc) {
	ch <- c.goroutinesDesc
	ch <- c.threadsDesc
	ch <- c.gcDesc
	ch <- c.gcLastTimeDesc
	ch <- c.goInfoDesc
}

// Collect returns the current state of all metrics of the collector.
func (c *baseGoCollector) Collect(ch chan<- Metric) {
	ch <- MustNewConstMetric(c.goroutinesDesc, GaugeValue, float64(runtime.NumGoroutine()))

	n := getRuntimeNumThreads()
	ch <- MustNewConstMetric(c.threadsDesc, GaugeValue, n)

	var stats debug.GCStats
	stats.PauseQuantiles = make([]time.Duration, 5)
	debug.ReadGCStats(&stats)

	quantiles := make(map[float64]float64)
	for idx, pq := range stats.PauseQuantiles[1:] {
		quantiles[float64(idx+1)/float64(len(stats.PauseQuantiles)-1)] = pq.Seconds()
	}
	quantiles[0.0] = stats.PauseQuantiles[0].Seconds()
	ch <- MustNewConstSummary(c.gcDesc, uint64(stats.NumGC), stats.PauseTotal.Seconds(), quantiles)
	ch <- MustNewConstMetric(c.gcLastTimeDesc, GaugeValue, float64(stats.LastGC.UnixNano())/1e9)
	ch <- MustNewConstMetric(c.goInfoDesc, GaugeValue, 1)
}

func memstatNamespace(s string) string {
	return "go_memstats_" + s
}

// memStatsMetrics provide description, evaluator, runtime/metrics name, and
// value type for memstat metrics.
type memStatsMetrics []struct {
	desc    *Desc
	eval    func(*runtime.MemStats) float64
	valType ValueType
}
<<<<<<< HEAD

// NewBuildInfoCollector returns a collector collecting a single metric
// "go_build_info" with the constant value 1 and three labels "path", "version",
// and "checksum". Their label values contain the main module path, version, and
// checksum, respectively. The labels will only have meaningful values if the
// binary is built with Go module support and from source code retrieved from
// the source repository (rather than the local file system). This is usually
// accomplished by building from outside of GOPATH, specifying the full address
// of the main package, e.g. "GO111MODULE=on go run
// github.com/prometheus/client_golang/examples/random". If built without Go
// module support, all label values will be "unknown". If built with Go module
// support but using the source code from the local file system, the "path" will
// be set appropriately, but "checksum" will be empty and "version" will be
// "(devel)".
//
// This collector uses only the build information for the main module. See
// https://github.com/povilasv/prommod for an example of a collector for the
// module dependencies.
func NewBuildInfoCollector() Collector {
	path, version, sum := readBuildInfo()
	c := &selfCollector{MustNewConstMetric(
		NewDesc(
			"go_build_info",
			"Build information about the main Go module.",
			nil, Labels{"path": path, "version": version, "checksum": sum},
		),
		GaugeValue, 1)}
	c.init(c.self)
	return c
}
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)

// NewBuildInfoCollector returns a collector collecting a single metric
// "go_build_info" with the constant value 1 and three labels "path", "version",
// and "checksum". Their label values contain the main module path, version, and
// checksum, respectively. The labels will only have meaningful values if the
// binary is built with Go module support and from source code retrieved from
// the source repository (rather than the local file system). This is usually
// accomplished by building from outside of GOPATH, specifying the full address
// of the main package, e.g. "GO111MODULE=on go run
// github.com/prometheus/client_golang/examples/random". If built without Go
// module support, all label values will be "unknown". If built with Go module
// support but using the source code from the local file system, the "path" will
// be set appropriately, but "checksum" will be empty and "version" will be
// "(devel)".
//
// This collector uses only the build information for the main module. See
// https://github.com/povilasv/prommod for an example of a collector for the
// module dependencies.
func NewBuildInfoCollector() Collector {
	path, version, sum := readBuildInfo()
	c := &selfCollector{MustNewConstMetric(
		NewDesc(
			"go_build_info",
			"Build information about the main Go module.",
			nil, Labels{"path": path, "version": version, "checksum": sum},
		),
		GaugeValue, 1)}
	c.init(c.self)
	return c
}
=======
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
