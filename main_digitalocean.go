//go:build all || digitalocean
// +build all digitalocean

package main

import (
	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/provider/digitalocean"
)

func init() {
	p, err := digitalocean.NewDigitalOceanProvider(ctx, domainFilter, cfg.DryRun, cfg.DigitalOceanAPIPageSize)
	if err != nil {
		log.Fatal(err)
	}
	providerMap["digitalocean"] = p
}
