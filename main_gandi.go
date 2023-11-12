//go:build all || gandi
// +build all gandi

package main

import (
	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/provider/gandi"
)

func init() {
	p, err := gandi.NewGandiProvider(ctx, domainFilter, cfg.DryRun)
	if err != nil {
		log.Fatal(err)
	}
	providerMap["gandi"] = p
}
