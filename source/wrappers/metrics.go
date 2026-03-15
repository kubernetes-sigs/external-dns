/*
Copyright 2026 The Kubernetes Authors.

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

package wrappers

import (
	"github.com/prometheus/client_golang/prometheus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/metrics"
)

var (
	invalidEndpoints = metrics.NewGaugedVectorOpts(
		prometheus.GaugeOpts{
			Subsystem: "source",
			Name:      "invalid_endpoints",
			Help:      "Number of endpoints currently rejected due to invalid configuration, partitioned by record type and source.",
		},
		[]string{"record_type", "source_type"},
	)

	deduplicatedEndpoints = metrics.NewGaugedVectorOpts(
		prometheus.GaugeOpts{
			Subsystem: "source",
			Name:      "deduplicated_endpoints",
			Help:      "Number of endpoints currently removed as duplicates, partitioned by record type and source.",
		},
		[]string{"record_type", "source_type"},
	)
)

// endpointSource returns the source type from the endpoint's object reference,
// or "unknown" if the reference is not set. Sources that set RefObject will
// populate this with their name (e.g. "ingress", "service").
func endpointSource(ep *endpoint.Endpoint) string {
	if ref := ep.RefObject(); ref != nil && ref.Source != "" {
		return ref.Source
	}
	return "unknown"
}

// resetMetrics clears all gauge label combinations so that stale
// record-type/source-type pairs from a previous cycle do not linger.
// Must be called at the start of each Endpoints() cycle in the
// first wrapper in the chain (dedupSource).
func resetMetrics() {
	invalidEndpoints.Reset()
	deduplicatedEndpoints.Reset()
}

func init() {
	metrics.RegisterMetric.MustRegister(invalidEndpoints)
	metrics.RegisterMetric.MustRegister(deduplicatedEndpoints)
}
