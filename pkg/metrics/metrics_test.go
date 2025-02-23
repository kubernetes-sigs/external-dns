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
	"reflect"
	"testing"
	"unsafe"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/external-dns/internal/testutils"
)

type MockMetric struct {
	FQDN string
}

func (m *MockMetric) Get() *Metric {
	return &Metric{FQDN: m.FQDN}
}

func TestMustRegister(t *testing.T) {
	tests := []struct {
		name     string
		metrics  []IMetric
		expected int
	}{
		{
			name: "single metric",
			metrics: []IMetric{
				NewCounterWithOpts(prometheus.CounterOpts{Name: "test_counter"}),
			},
			expected: 1,
		},
		{
			name: "two metrics",
			metrics: []IMetric{
				NewGaugeWithOpts(prometheus.GaugeOpts{Name: "test_gauge"}),
				NewCounterWithOpts(prometheus.CounterOpts{Name: "test_counter"}),
			},
			expected: 2,
		},
		{
			name: "mix of metrics",
			metrics: []IMetric{
				NewGaugeWithOpts(prometheus.GaugeOpts{Name: "test_gauge"}),
				NewCounterWithOpts(prometheus.CounterOpts{Name: "test_counter"}),
				NewCounterVecWithOpts(prometheus.CounterOpts{Name: "test_counter_vec"}, []string{"label"}),
			},
			expected: 3,
		},
		{
			name: "unsupported metric",
			metrics: []IMetric{
				&MockMetric{FQDN: "unsupported_metric"},
			},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			registry := NewMetricsRegister()
			for _, m := range tt.metrics {
				registry.MustRegister(m)
			}

			assert.Len(t, registry.Metrics, tt.expected)
		})
	}
}

func TestUnsupportedMetricWarning(t *testing.T) {
	buf := testutils.LogsToBuffer(log.WarnLevel, t)
	registry := NewMetricsRegister()
	mockUnsupported := &MockMetric{FQDN: "unsupported_metric"}
	registry.MustRegister(mockUnsupported)
	assert.NotContains(t, registry.mName, "unsupported_metric")

	assert.Contains(t, buf.String(), "Unsupported metric type: *metrics.MockMetric")
}

func TestNewMetricsRegister(t *testing.T) {
	registry := NewMetricsRegister()

	assert.NotNil(t, registry)
	assert.NotNil(t, registry.Registerer)
	assert.Empty(t, registry.Metrics)
	assert.Empty(t, registry.mName)

	labels := prometheus.DefaultRegisterer

	// hacks to get the runtime metrics from prometheus library
	values := reflect.ValueOf(labels).Elem().FieldByName("dimHashesByName")
	values = reflect.NewAt(values.Type(), unsafe.Pointer(values.UnsafeAddr())).Elem()

	switch v := values.Interface().(type) {
	case map[string]uint64:
		for k := range v {
			fmt.Println(k)
		}
	}
	// for k, v := range dimHashesByName {
	// 	fmt.Println(k, v)
	// }

	// assert.Equal(t, cfg.Version, labels["version"])
	// assert.Equal(t, runtime.GOARCH, labels["go-arch"])
	// assert.Equal(t, runtime.Version(), labels["go-version"])
}
