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
	"reflect"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
	"github.com/stretchr/testify/assert"
)

func TestNewGaugeWithOpts(t *testing.T) {
	opts := prometheus.GaugeOpts{
		Name:      "test_gauge",
		Subsystem: "test_subsystem",
		Help:      "This is a test gauge",
	}

	gaugeMetric := NewGaugeWithOpts(opts)

	assert.Equal(t, "gauge", gaugeMetric.Type)
	assert.Equal(t, "test_gauge", gaugeMetric.Name)
	assert.Equal(t, Namespace, gaugeMetric.Namespace)
	assert.Equal(t, "test_subsystem", gaugeMetric.Subsystem)
	assert.Equal(t, "This is a test gauge", gaugeMetric.Help)
	assert.Equal(t, "test_subsystem_test_gauge", gaugeMetric.FQDN)
	assert.NotNil(t, gaugeMetric.Gauge)
}

func TestNewCounterWithOpts(t *testing.T) {
	opts := prometheus.CounterOpts{
		Name:      "test_counter",
		Subsystem: "test_subsystem",
		Help:      "This is a test counter",
	}

	counterMetric := NewCounterWithOpts(opts)

	assert.Equal(t, "counter", counterMetric.Type)
	assert.Equal(t, "test_counter", counterMetric.Name)
	assert.Equal(t, Namespace, counterMetric.Namespace)
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
	assert.Equal(t, Namespace, counterVecMetric.Namespace)
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

func TestNewBuildInfoCollector(t *testing.T) {
	metric := NewGaugeFuncMetric(prometheus.GaugeOpts{
		Namespace: Namespace,
		Name:      "build_info",
		ConstLabels: prometheus.Labels{
			"version":   "0.0.1",
			"goversion": "1.24",
			"arch":      "arm64",
		},
	})

	desc := metric.GaugeFunc.Desc()

	assert.Equal(t, "external_dns_build_info", reflect.ValueOf(desc).Elem().FieldByName("fqName").String())
	assert.Contains(t, desc.String(), "version=\"0.0.1\"")
}
