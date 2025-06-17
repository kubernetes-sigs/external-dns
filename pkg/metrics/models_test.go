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
	"testing"

	dto "github.com/prometheus/client_model/go"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
)

func TestNewGaugeWithOpts(t *testing.T) {
	opts := prometheus.GaugeOpts{
		Name:      "test_gauge",
		Namespace: "test_namespace",
		Subsystem: "test_subsystem",
		Help:      "This is a test gauge",
	}

	gaugeMetric := NewGaugeWithOpts(opts)

	assert.Equal(t, "gauge", gaugeMetric.Type)
	assert.Equal(t, "test_gauge", gaugeMetric.Name)
	assert.Equal(t, "test_namespace", gaugeMetric.Namespace)
	assert.Equal(t, "test_subsystem", gaugeMetric.Subsystem)
	assert.Equal(t, "This is a test gauge", gaugeMetric.Help)
	assert.Equal(t, "test_subsystem_test_gauge", gaugeMetric.FQDN)
	assert.NotNil(t, gaugeMetric.Gauge)
}

func TestNewCounterWithOpts(t *testing.T) {
	opts := prometheus.CounterOpts{
		Name:      "test_counter",
		Namespace: "test_namespace",
		Subsystem: "test_subsystem",
		Help:      "This is a test counter",
	}

	counterMetric := NewCounterWithOpts(opts)

	assert.Equal(t, "counter", counterMetric.Type)
	assert.Equal(t, "test_counter", counterMetric.Name)
	assert.Equal(t, "test_namespace", counterMetric.Namespace)
	assert.Equal(t, "test_subsystem", counterMetric.Subsystem)
	assert.Equal(t, "This is a test counter", counterMetric.Help)
	assert.Equal(t, "test_subsystem_test_counter", counterMetric.FQDN)
	assert.NotNil(t, counterMetric.Counter)
}

func TestNewCounterVecWithOpts(t *testing.T) {
	opts := prometheus.CounterOpts{
		Name:      "test_counter_vec",
		Namespace: "test_namespace",
		Subsystem: "test_subsystem",
		Help:      "This is a test counter vector",
	}

	labelNames := []string{"label1", "label2"}

	counterVecMetric := NewCounterVecWithOpts(opts, labelNames)

	assert.Equal(t, "counter", counterVecMetric.Type)
	assert.Equal(t, "test_counter_vec", counterVecMetric.Name)
	assert.Equal(t, "test_namespace", counterVecMetric.Namespace)
	assert.Equal(t, "test_subsystem", counterVecMetric.Subsystem)
	assert.Equal(t, "This is a test counter vector", counterVecMetric.Help)
	assert.Equal(t, "test_subsystem_test_counter_vec", counterVecMetric.FQDN)
	assert.NotNil(t, counterVecMetric.CounterVec)
}

func TestGaugeV_SetWithLabels(t *testing.T) {
	opts := prometheus.GaugeOpts{
		Name:      "test_gauge",
		Namespace: "test_ns",
		Subsystem: "test_sub",
		Help:      "help text",
	}
	gv := NewGaugedVectorOpts(opts, []string{"label1", "label2"})

	gv.SetWithLabels(1.23, "Alpha", "BETA")

	g, err := gv.Gauge.GetMetricWithLabelValues("alpha", "beta")
	assert.NoError(t, err)

	var m dto.Metric
	err = g.Write(&m)
	assert.NoError(t, err)
	assert.NotNil(t, m.Gauge)
	assert.InDelta(t, 1.23, *m.Gauge.Value, 0.01)

	// Override the value
	gv.SetWithLabels(4.56, "ALPHA", "beta")
	// reuse g (same label combination)
	err = g.Write(&m)
	assert.NoError(t, err)
	assert.InDelta(t, 4.56, *m.Gauge.Value, 0.01)

	assert.Len(t, m.Label, 2)
}
