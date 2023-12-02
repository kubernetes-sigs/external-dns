//go:build all || vinyldns
// +build all vinyldns

package main

import (
	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/provider/vinyldns"
)

func init() {
	if cfg.Provider == "vinyldns" {
		zoneIDFilter := provider.NewZoneIDFilter(cfg.ZoneIDFilter)
		p, err := vinyldns.NewVinylDNSProvider(domainFilter, zoneIDFilter, cfg.DryRun)
		if err != nil {
			log.Fatal(err)
		}
		providerMap[cfg.Provider] = p
	}
}
