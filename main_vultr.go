//go:build all || vultr
// +build all vultr

package main

import (
	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/provider/vultr"
)

func init() {
	if cfg.Provider == "vultr" {
		p, err := vultr.NewVultrProvider(ctx, domainFilter, cfg.DryRun)
		if err != nil {
			log.Fatal(err)
		}
		providerMap[cfg.Provider] = p
	}
}
