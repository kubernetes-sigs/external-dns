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

package testutils

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
	"github.com/stretchr/testify/assert"
)

// TestHelperVerifyMetricsGaugeVectorWithLabels verifies that a prometheus.GaugeVec metric with specific labels has the expected value.
// Supports partial label matching - if only some labels are provided, it sums all metrics matching those labels.
//
// Example usage:
//
//	Exact match (all labels)
//	labels := map[string]string{"method": "GET", "status": "200"}
//	TestHelperVerifyMetricsGaugeVectorWithLabels(t, 42.0, myGaugeVec, labels)
//
//	Partial match (sum all metrics with method=GET)
//	labels := map[string]string{"method": "GET"}
//	TestHelperVerifyMetricsGaugeVectorWithLabels(t, 100.0, myGaugeVec, labels)
func TestHelperVerifyMetricsGaugeVectorWithLabels(t *testing.T, expected float64, metric prometheus.GaugeVec, labels map[string]string) {
	TestHelperVerifyMetricsGaugeVectorWithLabelsFunc(t, expected, assert.Equal, metric, labels)
}

// TestHelperVerifyMetricsGaugeVectorWithLabelsFunc is a helper function that verifies a prometheus.GaugeVec metric with specific labels using a custom assertion function.
// Supports partial label matching - if only some labels are provided, it sums all metrics matching those labels.
func TestHelperVerifyMetricsGaugeVectorWithLabelsFunc(t *testing.T, expected float64, aFunc assert.ComparisonAssertionFunc, metric prometheus.GaugeVec, labels map[string]string) {
	t.Helper()

	// Collect all metrics and find matching ones
	actual := sumMetricsWithLabels(&metric, labels)

	if !aFunc(t, expected, actual, "Expected gauge value does not match the actual value", labels) {
		t.Logf("Available metrics:\n%s", collectGaugeVecMetrics(&metric))
	}
}

// SummaryVecSampleCount returns the total observation count from a SummaryVec
// across all label combinations that match every key/value pair in match.
// Unspecified labels (e.g. status) are ignored, so this supports partial matching.
func SummaryVecSampleCount(t *testing.T, sv *prometheus.SummaryVec, match prometheus.Labels) uint64 {
	t.Helper()
	var total uint64
	for _, dm := range collectAll(sv) {
		if labelsMatch(dm, match) {
			total += dm.GetSummary().GetSampleCount()
		}
	}
	return total
}

// labelsMatch reports whether dm contains all key/value pairs in match (case-insensitive, partial).
func labelsMatch(dm *dto.Metric, match map[string]string) bool {
	var matchCount int
	for _, lp := range dm.GetLabel() {
		v, found := match[lp.GetName()]
		if found {
			if !strings.EqualFold(v, lp.GetValue()) {
				return false
			}
			matchCount++
		}
	}
	return matchCount == len(match)
}

// collectAll drains all current observations from a Collector into a slice.
func collectAll(collector prometheus.Collector) []*dto.Metric {
	ch := make(chan prometheus.Metric, 1024)
	go func() {
		collector.Collect(ch)
		close(ch)
	}()
	var result []*dto.Metric
	for m := range ch {
		var dm dto.Metric
		if err := m.Write(&dm); err != nil {
			continue
		}
		result = append(result, &dm)
	}
	return result
}

// sumMetricsWithLabels sums all metric values that match the provided labels (partial match supported).
func sumMetricsWithLabels(metric *prometheus.GaugeVec, matchLabels map[string]string) float64 {
	var sum float64
	for _, dm := range collectAll(metric) {
		if labelsMatch(dm, matchLabels) && dm.Gauge != nil {
			sum += dm.Gauge.GetValue()
		}
	}
	return sum
}

// collectGaugeVecMetrics collects all metrics from a GaugeVec and returns a formatted string.
// Shows both per-label aggregates and detailed metrics.
func collectGaugeVecMetrics(metric *prometheus.GaugeVec) string {
	// Collect all metrics and aggregate by label
	type metricEntry struct {
		labels map[string]string
		value  float64
	}
	var entries []metricEntry
	aggregates := make(map[string]map[string]float64) // labelName -> labelValue -> sum

	for _, dm := range collectAll(metric) {
		entry := metricEntry{labels: make(map[string]string)}
		for _, lp := range dm.Label {
			name, value := lp.GetName(), lp.GetValue()
			entry.labels[name] = value

			if aggregates[name] == nil {
				aggregates[name] = make(map[string]float64)
			}
			if dm.Gauge != nil {
				aggregates[name][value] += dm.Gauge.GetValue()
			}
		}
		if dm.Gauge != nil {
			entry.value = dm.Gauge.GetValue()
		}
		entries = append(entries, entry)
	}

	if len(entries) == 0 {
		return "  (no metrics collected)"
	}

	var sb strings.Builder

	// Output aggregates by label (sorted)
	sb.WriteString("Totals by label:\n")
	var labelNames []string
	for name := range aggregates {
		labelNames = append(labelNames, name)
	}
	sort.Strings(labelNames)

	for _, name := range labelNames {
		values := aggregates[name]
		var pairs []string
		for v, sum := range values {
			pairs = append(pairs, fmt.Sprintf("%s=%.0f", v, sum))
		}
		sort.Strings(pairs)
		fmt.Fprintf(&sb, "  %s: %s\n", name, strings.Join(pairs, ", "))
	}

	// Output detailed metrics
	sb.WriteString("\nAll metrics:\n")
	for _, e := range entries {
		var labels []string
		for k, v := range e.labels {
			labels = append(labels, fmt.Sprintf("%s=%q", k, v))
		}
		sort.Strings(labels)
		fmt.Fprintf(&sb, "  {%s} = %.2f\n", strings.Join(labels, ", "), e.value)
	}

	return sb.String()
}
