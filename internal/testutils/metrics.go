package testutils

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
	"github.com/stretchr/testify/assert"
)

// TestHelperVerifyMetricsGaugeVectorWithLabels verifies that a prometheus.GaugeVec metric with specific labels has the expected value.
//
// Example usage:
//
//	labels := map[string]string{"method": "GET", "status": "200"}
//	TestHelperVerifyMetricsGaugeVectorWithLabels(t, 42.0, myGaugeVec, labels)
func TestHelperVerifyMetricsGaugeVectorWithLabels(t *testing.T, expected float64, metric prometheus.GaugeVec, labels map[string]string) {
	TestHelperVerifyMetricsGaugeVectorWithLabelsFunc(t, expected, assert.Equal, metric, labels)
}

// TestHelperVerifyMetricsGaugeVectorWithLabelsFunc is a helper function that verifies a prometheus.GaugeVec metric with specific labels using a custom assertion function.
func TestHelperVerifyMetricsGaugeVectorWithLabelsFunc(t *testing.T, expected float64, aFunc assert.ComparisonAssertionFunc, metric prometheus.GaugeVec, labels map[string]string) {
	t.Helper()

	g, err := metric.MetricVec.GetMetricWith(labels)
	assert.NoError(t, err)

	var m dto.Metric
	err = g.Write(&m)
	assert.NoError(t, err)

	assert.NotNil(t, m.Gauge)

	aFunc(t, expected, *m.Gauge.Value, "Expected gauge value does not match the actual value", labels)
}
