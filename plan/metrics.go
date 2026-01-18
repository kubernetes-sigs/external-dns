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

package plan

import (
	"github.com/prometheus/client_golang/prometheus"

	"sigs.k8s.io/external-dns/pkg/metrics"
)

var (
	// registryOwnerMismatchPerSync tracks records skipped due to owner mismatch.
	// The "domain" label uses the naked/apex domain (e.g., "example.com") rather than
	// full FQDNs to prevent cardinality explosion. With thousands of subdomains under
	// one apex domain, using full FQDNs would create excessive metric series.
	registryOwnerMismatchPerSync = metrics.NewGaugedVectorOpts(
		prometheus.GaugeOpts{
			Subsystem: "registry",
			Name:      "skipped_records_owner_mismatch_per_sync",
			Help:      "Number of records skipped with owner mismatch for each record type, owner mismatch ID and domain (vector).",
		},
		[]string{"record_type", "owner", "foreign_owner", "domain"},
	)
)

type mismatchKey struct {
	recordType   string
	owner        string
	foreignOwner string
	domain       string
}

func init() {
	metrics.RegisterMetric.MustRegister(registryOwnerMismatchPerSync)
}
