//go:build all || dnsimple
// +build all dnsimple

package main

import (
	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/provider/dnsimple"
)

func init() {
	if cfg.Provider == "dnsimple" {
		zoneIDFilter := provider.NewZoneIDFilter(cfg.ZoneIDFilter)
		p, err := dnsimple.NewDnsimpleProvider(domainFilter, zoneIDFilter, cfg.DryRun)
		if err != nil {
			log.Fatal(err)
		}
		providerMap[cfg.Provider] = p
	}
}
