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

package controller

import (
	"github.com/prometheus/client_golang/prometheus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/metrics"
)

var (
	registryErrorsTotal = metrics.NewCounterWithOpts(
		prometheus.CounterOpts{
			Subsystem: "registry",
			Name:      "errors_total",
			Help:      "Number of Registry errors.",
		},
	)
	sourceErrorsTotal = metrics.NewCounterWithOpts(
		prometheus.CounterOpts{
			Subsystem: "source",
			Name:      "errors_total",
			Help:      "Number of Source errors.",
		},
	)
	sourceEndpointsTotal = metrics.NewGaugeWithOpts(
		prometheus.GaugeOpts{
			Subsystem: "source",
			Name:      "endpoints_total",
			Help:      "Number of Endpoints in all sources",
		},
	)
	registryEndpointsTotal = metrics.NewGaugeWithOpts(
		prometheus.GaugeOpts{
			Subsystem: "registry",
			Name:      "endpoints_total",
			Help:      "Number of Endpoints in the registry",
		},
	)
	lastSyncTimestamp = metrics.NewGaugeWithOpts(
		prometheus.GaugeOpts{
			Subsystem: "controller",
			Name:      "last_sync_timestamp_seconds",
			Help:      "Timestamp of last successful sync with the DNS provider",
		},
	)
	lastReconcileTimestamp = metrics.NewGaugeWithOpts(
		prometheus.GaugeOpts{
			Subsystem: "controller",
			Name:      "last_reconcile_timestamp_seconds",
			Help:      "Timestamp of last attempted sync with the DNS provider",
		},
	)
	controllerNoChangesTotal = metrics.NewCounterWithOpts(
		prometheus.CounterOpts{
			Subsystem: "controller",
			Name:      "no_op_runs_total",
			Help:      "Number of reconcile loops ending up with no changes on the DNS provider side.",
		},
	)
	deprecatedRegistryErrors = metrics.NewCounterWithOpts(
		prometheus.CounterOpts{
			Subsystem: "registry",
			Name:      "errors_total",
			Help:      "Number of Registry errors.",
		},
	)
	deprecatedSourceErrors = metrics.NewCounterWithOpts(
		prometheus.CounterOpts{
			Subsystem: "source",
			Name:      "errors_total",
			Help:      "Number of Source errors.",
		},
	)

	registryRecords = metrics.NewGaugedVectorOpts(
		prometheus.GaugeOpts{
			Subsystem: "registry",
			Name:      "records",
			Help:      "Number of registry records partitioned by label name (vector).",
		},
		[]string{"record_type"},
	)

	sourceRecords = metrics.NewGaugedVectorOpts(
		prometheus.GaugeOpts{
			Subsystem: "source",
			Name:      "records",
			Help:      "Number of source records partitioned by label name (vector).",
		},
		[]string{"record_type"},
	)

	verifiedRecords = metrics.NewGaugedVectorOpts(
		prometheus.GaugeOpts{
			Subsystem: "controller",
			Name:      "verified_records",
			Help:      "Number of DNS records that exists both in source and registry (vector).",
		},
		[]string{"record_type"},
	)

	consecutiveSoftErrors = metrics.NewGaugeWithOpts(
		prometheus.GaugeOpts{
			Subsystem: "controller",
			Name:      "consecutive_soft_errors",
			Help:      "Number of consecutive soft errors in reconciliation loop.",
		},
	)
)

func init() {
	metrics.RegisterMetric.MustRegister(registryErrorsTotal)
	metrics.RegisterMetric.MustRegister(sourceErrorsTotal)
	metrics.RegisterMetric.MustRegister(sourceEndpointsTotal)
	metrics.RegisterMetric.MustRegister(registryEndpointsTotal)
	metrics.RegisterMetric.MustRegister(lastSyncTimestamp)
	metrics.RegisterMetric.MustRegister(lastReconcileTimestamp)
	metrics.RegisterMetric.MustRegister(deprecatedRegistryErrors)
	metrics.RegisterMetric.MustRegister(deprecatedSourceErrors)
	metrics.RegisterMetric.MustRegister(controllerNoChangesTotal)

	metrics.RegisterMetric.MustRegister(registryRecords)
	metrics.RegisterMetric.MustRegister(sourceRecords)
	metrics.RegisterMetric.MustRegister(verifiedRecords)

	metrics.RegisterMetric.MustRegister(consecutiveSoftErrors)
}

type dnsKey struct {
	name       string
	recordType string
}

// countMatchingAddressRecords counts records that exist in both endpoints and registry.
func countMatchingAddressRecords(endpoints []*endpoint.Endpoint, registryRecords []*endpoint.Endpoint, metric metrics.GaugeVecMetric) {
	metric.Gauge.Reset()

	registry := make(map[dnsKey]struct{}, len(registryRecords))
	for _, r := range registryRecords {
		registry[dnsKey{r.DNSName, r.RecordType}] = struct{}{}
	}

	counts := make(map[string]float64)
	for _, ep := range endpoints {
		if _, found := registry[dnsKey{ep.DNSName, ep.RecordType}]; found {
			counts[ep.RecordType]++
		}
	}

	for recordType, count := range counts {
		metric.AddWithLabels(count, recordType)
	}
}

// countAddressRecords counts each record type in the provided endpoints slice.
func countAddressRecords(endpoints []*endpoint.Endpoint, metric metrics.GaugeVecMetric) {
	metric.Gauge.Reset()

	counts := make(map[string]float64)
	for _, ep := range endpoints {
		counts[ep.RecordType]++
	}

	for recordType, count := range counts {
		metric.AddWithLabels(count, recordType)
	}
}
