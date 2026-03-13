/*
Copyright 2025 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package metrics

import (
	"fmt"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

type MetricRegistry struct {
	Registerer prometheus.Registerer
	Metrics    []*Metric
	mName      map[string]bool
}

type Metric struct {
	Type      string
	Namespace string
	Subsystem string
	Name      string
	Help      string
	FQDN      string
}

type IMetric interface {
	Get() *Metric
}

type GaugeMetric struct {
	Metric
	Gauge prometheus.Gauge
}

func (g GaugeMetric) Get() *Metric {
	return &g.Metric
}

type CounterMetric struct {
	Metric
	Counter prometheus.Counter
}

func (g CounterMetric) Get() *Metric {
	return &g.Metric
}

type CounterVecMetric struct {
	Metric
	CounterVec *prometheus.CounterVec
}

func (g CounterVecMetric) Get() *Metric {
	return &g.Metric
}

type GaugeVecMetric struct {
	Metric
	Gauge prometheus.GaugeVec
}

func (g GaugeVecMetric) Get() *Metric {
	return &g.Metric
}

// SetWithLabels sets the value of the Gauge metric for the specified label values.
// All label values are converted to lowercase before being applied.
func (g GaugeVecMetric) SetWithLabels(value float64, lvs ...string) {
	g.Gauge.WithLabelValues(toLower(lvs)...).Set(value)
}

// AddWithLabels adds the value to the Gauge metric for the specified label values.
// All label values are converted to lowercase before being applied.
//
// Without Reset(), values accumulate and reset only on process restart.
// Use Reset() + AddWithLabels() pattern for per-cycle counts.
func (g GaugeVecMetric) AddWithLabels(value float64, lvs ...string) {
	g.Gauge.WithLabelValues(toLower(lvs)...).Add(value)
}

func NewGaugeWithOpts(opts prometheus.GaugeOpts) GaugeMetric {
	opts.Namespace = Namespace
	return GaugeMetric{
		Metric: Metric{
			Type:      "gauge",
			Name:      opts.Name,
			FQDN:      fmt.Sprintf("%s_%s", opts.Subsystem, opts.Name),
			Namespace: opts.Namespace,
			Subsystem: opts.Subsystem,
			Help:      opts.Help,
		},
		Gauge: prometheus.NewGauge(opts),
	}
}

// NewGaugedVectorOpts creates a new GaugeVec based on the provided GaugeOpts and
// partitioned by the given label names.
func NewGaugedVectorOpts(opts prometheus.GaugeOpts, labelNames []string) GaugeVecMetric {
	opts.Namespace = Namespace
	return GaugeVecMetric{
		Metric: Metric{
			Type:      "gauge",
			Name:      opts.Name,
			FQDN:      fmt.Sprintf("%s_%s", opts.Subsystem, opts.Name),
			Namespace: opts.Namespace,
			Subsystem: opts.Subsystem,
			Help:      opts.Help,
		},
		Gauge: *prometheus.NewGaugeVec(opts, labelNames),
	}
}

func NewCounterWithOpts(opts prometheus.CounterOpts) CounterMetric {
	opts.Namespace = Namespace
	return CounterMetric{
		Metric: Metric{
			Type:      "counter",
			Name:      opts.Name,
			FQDN:      fmt.Sprintf("%s_%s", opts.Subsystem, opts.Name),
			Namespace: opts.Namespace,
			Subsystem: opts.Subsystem,
			Help:      opts.Help,
		},
		Counter: prometheus.NewCounter(opts),
	}
}

func NewCounterVecWithOpts(opts prometheus.CounterOpts, labelNames []string) CounterVecMetric {
	opts.Namespace = Namespace
	return CounterVecMetric{
		Metric: Metric{
			Type:      "counter",
			Name:      opts.Name,
			FQDN:      fmt.Sprintf("%s_%s", opts.Subsystem, opts.Name),
			Namespace: opts.Namespace,
			Subsystem: opts.Subsystem,
			Help:      opts.Help,
		},
		CounterVec: prometheus.NewCounterVec(opts, labelNames),
	}
}

type GaugeFuncMetric struct {
	Metric
	GaugeFunc prometheus.GaugeFunc
}

func (g GaugeFuncMetric) Get() *Metric {
	return &g.Metric
}

func NewGaugeFuncMetric(opts prometheus.GaugeOpts) GaugeFuncMetric {
	return GaugeFuncMetric{
		Metric: Metric{
			Type: "gauge",
			Name: opts.Name,
			FQDN: func() string {
				if opts.Subsystem != "" {
					return fmt.Sprintf("%s_%s", opts.Subsystem, opts.Name)
				}
				return opts.Name
			}(),
			Namespace: opts.Namespace,
			Subsystem: opts.Subsystem,
			Help:      opts.Help,
		},
		GaugeFunc: prometheus.NewGaugeFunc(opts, func() float64 { return 1 }),
	}
}

type SummaryVecMetric struct {
	Metric
	SummaryVec prometheus.SummaryVec
}

func (s SummaryVecMetric) Get() *Metric {
	return &s.Metric
}

func (s SummaryVecMetric) SetWithLabels(value float64, labels prometheus.Labels) {
	s.SummaryVec.With(labels).Observe(value)
}

func NewSummaryVecWithOpts(opts prometheus.SummaryOpts, labels []string) SummaryVecMetric {
	opts.Namespace = Namespace
	return SummaryVecMetric{
		Metric: Metric{
			Type:      "summaryVec",
			Name:      opts.Name,
			FQDN:      fmt.Sprintf("%s_%s", opts.Subsystem, opts.Name),
			Namespace: opts.Namespace,
			Subsystem: opts.Subsystem,
			Help:      opts.Help,
		},
		SummaryVec: *prometheus.NewSummaryVec(opts, labels),
	}
}

func PathProcessor(path string) string {
	parts := strings.Split(path, "/")
	return parts[len(parts)-1]
}

// toLower converts all label values to lowercase.
// The Prometheus maintainers have intentionally avoided magic transformations to keep label handling explicit and predictable.
// We expect consistent casing, normalizing at ingestion is the standard practice.
func toLower(lvs []string) []string {
	for i := range lvs {
		lvs[i] = strings.ToLower(lvs[i])
	}
	return lvs
}
