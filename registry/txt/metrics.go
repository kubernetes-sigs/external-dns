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

// overflow_in label values.
const (
	overflowInSource    = "source"
	overflowInTXTPrefix = "txt_prefix"
)

// labelOverflowPerSync tracks records dropped because a DNS label exceeded RFC
// 1035's 63-char limit. The "domain" label uses the naked/apex domain (e.g.,
// "example.com") rather than full FQDNs to prevent cardinality explosion.
// "overflow_in" distinguishes whether the overflow was caught at NewEndpoint
// (source name) or in the projected TXT registry name after a record-type/user
// prefix is applied.
var labelOverflowPerSync = metrics.NewGaugedVectorOpts(
	prometheus.GaugeOpts{
		Subsystem: "registry",
		Name:      "skipped_records_label_overflow_per_sync",
		Help:      "Number of records skipped per sync because a DNS label exceeds RFC 1035's 63-char limit, by record type, apex domain, and whether the overflow is in the source name or the projected TXT registry name.",
	},
	[]string{"record_type", "domain", "overflow_in"},
)

func init() {
	metrics.RegisterMetric.MustRegister(labelOverflowPerSync)
	// Wire the endpoint package's source-side reporter so NewEndpoint's nil
	// path increments the same gauge. Without this, source overflows would go
	// untracked. endpoint cannot import this package directly (registry/txt
	// depends on endpoint), so a func var is set here at init time.
	endpoint.SourceLabelOverflowReporter = func(recordType, domain string) {
		labelOverflowPerSync.AddWithLabels(1.0, recordType, domain, overflowInSource)
	}
}

// recordTXTPrefixOverflow increments the per-sync gauge for an owned record
// dropped because the projected TXT name overflows RFC 1035's 63-char limit.
func recordTXTPrefixOverflow(skipped *endpoint.Endpoint) {
	labelOverflowPerSync.AddWithLabels(1.0, skipped.RecordType, skipped.GetNakedDomain(), overflowInTXTPrefix)
}

// resetLabelOverflow clears the per-sync gauge. Called at the start of each
// reconciliation loop, before sources or the registry can emit.
func resetLabelOverflow() {
	labelOverflowPerSync.Gauge.Reset()
}
