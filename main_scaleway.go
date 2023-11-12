//go:build all || scaleway
// +build all scaleway

package main

import (
	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/provider/scaleway"
)

func init() {
	p, err := scaleway.NewScalewayProvider(ctx, domainFilter, cfg.DryRun)
	if err != nil {
		log.Fatal(err)
	}
	providerMap["scaleway"] = p
}
