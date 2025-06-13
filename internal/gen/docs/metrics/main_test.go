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

package main

import (
	"fmt"
	"io/fs"
	"os"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"sigs.k8s.io/external-dns/pkg/metrics"
)

const pathToDocs = "%s/../../../../docs/monitoring"

func TestComputeMetrics(t *testing.T) {
	reg := metrics.RegisterMetric

	if len(reg.Metrics) == 0 {
		t.Errorf("Expected not empty metrics registry, got %d", len(reg.Metrics))
	}

	assert.Len(t, reg.Metrics, 19)
}

func TestGenerateMarkdownTableRenderer(t *testing.T) {
	reg := metrics.NewMetricsRegister()

	got, err := generateMarkdownTable(reg, false)
	assert.NoError(t, err)

	assert.Contains(t, got, "# Available Metrics\n\n<!-- THIS FILE MUST NOT BE EDITED BY HAND -->\n")
	assert.Contains(t, got, "| Metric Type | Subsystem   |  Help")
}

func TestGenerateMarkdownTableWithSingleMetric(t *testing.T) {
	reg := metrics.NewMetricsRegister()

	reg.MustRegister(metrics.NewGaugeWithOpts(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "controller_0",
			Name:      "verified_aaaa_records",
			Help:      "This is just a test.",
		},
	))

	got, err := generateMarkdownTable(reg, false)
	require.NoError(t, err)

	assert.Contains(t, got, "verified_aaaa_records")
	assert.Contains(t, got, "This is just a test.")
}

func TestMetricsMdUpToDate(t *testing.T) {
	testPath, _ := os.Getwd()
	fsys := os.DirFS(fmt.Sprintf(pathToDocs, testPath))
	fileName := "metrics.md"
	expected, err := fs.ReadFile(fsys, fileName)
	assert.NoError(t, err, "expected file %s to exist", fileName)

	reg := metrics.RegisterMetric
	actual, err := generateMarkdownTable(reg, false)
	assert.NoError(t, err)
	assert.Contains(t, string(expected), actual)
}

func TestMetricsMdExtraMetricAdded(t *testing.T) {
	testPath, _ := os.Getwd()
	fsys := os.DirFS(fmt.Sprintf(pathToDocs, testPath))
	fileName := "metrics.md"
	expected, err := fs.ReadFile(fsys, fileName)
	assert.NoError(t, err, "expected file %s to exist", fileName)

	reg := metrics.RegisterMetric

	reg.MustRegister(metrics.NewGaugeWithOpts(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "controller_1",
			Name:      "verified_aaaa_records",
			Help:      "This is just a test.",
		},
	))

	actual, err := generateMarkdownTable(reg, false)
	assert.NoError(t, err)
	assert.NotEqual(t, string(expected), actual)
}

func TestGetRuntimeMetricsForNewRegistry(t *testing.T) {
	reg := prometheus.NewRegistry()
	// Register some runtime metrics
	reg.MustRegister(prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "go_goroutines",
		Help: "Number of goroutines that currently exist.",
	}))
	reg.MustRegister(prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "go_memstats_alloc_bytes",
		Help: "Number of bytes allocated and still in use.",
	}))
	runtimeMetrics := getRuntimeMetrics(reg)

	// Check that the runtime metrics are correctly retrieved
	expectedMetrics := []string{"go_goroutines", "go_memstats_alloc_bytes"}
	assert.ElementsMatch(t, expectedMetrics, runtimeMetrics)
	assert.Len(t, runtimeMetrics, 2)
}

func TestGetRuntimeMetricsForDefaultRegistry(t *testing.T) {
	reg := prometheus.DefaultRegisterer
	runtimeMetrics := getRuntimeMetrics(reg)
	if len(runtimeMetrics) == 0 {
		t.Errorf("Expected not empty runtime metrics, got %d", len(runtimeMetrics))
	}
}
