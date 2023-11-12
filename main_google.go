//go:build all || google
// +build all google

package main

import (
	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/provider/google"
)

func init() {
	zoneIDFilter := provider.NewZoneIDFilter(cfg.ZoneIDFilter)
	p, err := google.NewGoogleProvider(ctx, cfg.GoogleProject, domainFilter, zoneIDFilter, cfg.GoogleBatchChangeSize, cfg.GoogleBatchChangeInterval, cfg.GoogleZoneVisibility, cfg.DryRun)
	if err != nil {
		log.Fatal(err)
	}
	providerMap["google"] = p
}
