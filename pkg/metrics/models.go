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

func NewGaugeWithOpts(opts prometheus.GaugeOpts) GaugeMetric {
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

func NewCounterWithOpts(opts prometheus.CounterOpts) CounterMetric {
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
