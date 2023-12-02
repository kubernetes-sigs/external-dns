//go:build all || ovh
// +build all ovh

package main

import (
	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/provider/ovh"
)

func init() {
	if cfg.Provider == "ovh" {
		p, err := ovh.NewOVHProvider(ctx, domainFilter, cfg.OVHEndpoint, cfg.OVHApiRateLimit, cfg.DryRun)
		if err != nil {
			log.Fatal(err)
		}
		providerMap[cfg.Provider] = p
	}
}
