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

package txt

import (
	"github.com/prometheus/client_golang/prometheus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/metrics"
)

// registrySkippedLabelTooLongPerSync tracks records dropped because the projected
// TXT name has a label exceeding the 63-char RFC 1035 limit.
var registrySkippedLabelTooLongPerSync = metrics.NewGaugedVectorOpts(
	prometheus.GaugeOpts{
		Subsystem: "registry",
		Name:      "skipped_records_label_too_long_per_sync",
		Help:      "Number of records skipped because the projected TXT registry name has a DNS label exceeding RFC 1035's 63-char limit, for each record type and domain (vector).",
	},
	[]string{"record_type", "domain"},
)

func init() {
	metrics.RegisterMetric.MustRegister(registrySkippedLabelTooLongPerSync)
}

// recordSkippedLabelTooLong increments the per-sync gauge for a record dropped
// because its projected TXT name overflows RFC 1035's 63-char label limit.
func recordSkippedLabelTooLong(skipped *endpoint.Endpoint) {
	registrySkippedLabelTooLongPerSync.AddWithLabels(1.0, skipped.RecordType, skipped.GetNakedDomain())
}
