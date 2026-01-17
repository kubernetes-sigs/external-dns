package plan

import (
	"github.com/prometheus/client_golang/prometheus"

	"sigs.k8s.io/external-dns/pkg/metrics"
)

var (
	registryOwnerMismatchTotal = metrics.NewGaugedVectorOpts(
		prometheus.GaugeOpts{
			Subsystem: "registry",
			Name:      "skipped_records_owner_mismatch_total",
			Help:      "Number of records skipped with owner mismatch for each record type, owner mismatch ID and domain (vector).",
		},
		[]string{"record_type", "foreign_owner", "domain"},
	)
)

func init() {
	metrics.RegisterMetric.MustRegister(registryOwnerMismatchTotal)
}
