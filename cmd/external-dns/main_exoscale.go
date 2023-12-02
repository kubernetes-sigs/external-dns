//go:build all || exoscale
// +build all exoscale

package main

import (
	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/provider/exoscale"
)

func init() {
	if cfg.Provider == "exoscale" {
		p, err := exoscale.NewExoscaleProvider(
			cfg.ExoscaleAPIEnvironment,
			cfg.ExoscaleAPIZone,
			cfg.ExoscaleAPIKey,
			cfg.ExoscaleAPISecret,
			cfg.DryRun,
			exoscale.ExoscaleWithDomain(domainFilter),
			exoscale.ExoscaleWithLogging(),
		)
		if err != nil {
			log.Fatal(err)
		}
		providerMap[cfg.Provider] = p
	}
}
