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
	"runtime"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestNewGaugeFuncMetric(t *testing.T) {
	tests := []struct {
		name                    string
		metricName              string
		subSystem               string
		constLabels             prometheus.Labels
		expectedFqName          string
		expectedDescString      string
		expectedGaugeFuncReturn float64
	}{
		{
			name:       "NewGaugeFuncMetric returns build_info",
			metricName: "build_info",
			subSystem:  "",
			constLabels: prometheus.Labels{
				"version":   "0.0.1",
				"goversion": runtime.Version(),
				"arch":      "arm64",
			},
			expectedFqName:          "external_dns_build_info",
			expectedDescString:      "version=\"0.0.1\"",
			expectedGaugeFuncReturn: 1,
		},
		{
			name:                    "NewGaugeFuncMetric subsystem alters name",
			metricName:              "metricName",
			subSystem:               "subSystem",
			constLabels:             prometheus.Labels{},
			expectedFqName:          "external_dns_subSystem_metricName",
			expectedDescString:      "",
			expectedGaugeFuncReturn: 1,
		},
		{
			name:                    "NewGaugeFuncMetric GaugeFunc returns 1",
			metricName:              "metricName",
			subSystem:               "",
			constLabels:             prometheus.Labels{},
			expectedFqName:          "external_dns_metricName",
			expectedDescString:      "",
			expectedGaugeFuncReturn: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			metric := NewGaugeFuncMetric(prometheus.GaugeOpts{
				Namespace:   Namespace,
				Name:        tt.metricName,
				Subsystem:   tt.subSystem,
				ConstLabels: tt.constLabels,
			})

			desc := metric.GaugeFunc.Desc()

			assert.Equal(t, tt.expectedFqName, reflect.ValueOf(desc).Elem().FieldByName("fqName").String())
			assert.Contains(t, desc.String(), tt.expectedDescString)

			testRegistry := prometheus.NewRegistry()
			err := testRegistry.Register(metric.GaugeFunc)
			require.NoError(t, err)

			metricFamily, err := testRegistry.Gather()
			require.NoError(t, err)
			require.Len(t, metricFamily, 1)

			require.NotNil(t, metricFamily[0].Metric[0].Gauge)
			assert.InDelta(t, tt.expectedGaugeFuncReturn, metricFamily[0].Metric[0].GetGauge().GetValue(), 0.0001)
		})
	}
}

func TestSummaryV_SetWithLabels(t *testing.T) {
	opts := prometheus.SummaryOpts{
		Name:      "test_summaryVec",
		Namespace: "test_ns",
		Subsystem: "test_sub",
		Help:      "help text",
	}

	labels := Labels{}
	sv := NewSummaryVecWithOpts(opts, []string{"label1", "label2"})

	labels["label1"] = "alpha"
	labels["label2"] = "beta"

	sv.SetWithLabels(5.01, labels)

	reg := prometheus.NewRegistry()
	reg.MustRegister(sv.SummaryVec)

	metricsFamilies, err := reg.Gather()
	assert.NoError(t, err)
	assert.Len(t, metricsFamilies, 1)

	s, err := sv.SummaryVec.GetMetricWithLabelValues("alpha", "beta")
	assert.NoError(t, err)
	metricsFamilies, err = reg.Gather()

	s.Observe(5.21)
	metricsFamilies, err = reg.Gather()
	assert.NoError(t, err)

	assert.InDelta(t, 10.22, *metricsFamilies[0].Metric[0].Summary.SampleSum, 0.01)
	assert.Len(t, metricsFamilies[0].Metric[0].Label, 2)
}

func TestPathProcessor(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"/foo/bar", "bar"},
		{"/foo/", ""},
		{"/", ""},
		{"", ""},
		{"/foo/bar/baz", "baz"},
		{"foo/bar", "bar"},
		{"foo", "foo"},
		{"foo/", ""},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			require.Equal(t, tt.expected, PathProcessor(tt.input))
		})
	}
}

func TestGaugeV_AddWithLabels(t *testing.T) {
	opts := prometheus.GaugeOpts{
		Name:      "test_gauge_add",
		Namespace: "test_ns",
		Subsystem: "test_sub",
		Help:      "help text",
	}
	gv := NewGaugedVectorOpts(opts, []string{"label1", "label2"})

	// Add with mixed case labels - should be lowercased
	gv.AddWithLabels(1.0, "Alpha", "BETA")

	g, err := gv.Gauge.GetMetricWithLabelValues("alpha", "beta")
	assert.NoError(t, err)

	var m dto.Metric
	err = g.Write(&m)
	assert.NoError(t, err)
	assert.NotNil(t, m.Gauge)
	assert.InDelta(t, 1.0, *m.Gauge.Value, 0.01)

	// Add again - should increment, not override
	gv.AddWithLabels(2.0, "ALPHA", "beta")
	err = g.Write(&m)
	assert.NoError(t, err)
	assert.InDelta(t, 3.0, *m.Gauge.Value, 0.01) // 1.0 + 2.0 = 3.0

	// Add one more time
	gv.AddWithLabels(0.5, "alpha", "Beta")
	err = g.Write(&m)
	assert.NoError(t, err)
	assert.InDelta(t, 3.5, *m.Gauge.Value, 0.01) // 3.0 + 0.5 = 3.5

	assert.Len(t, m.Label, 2)
}
