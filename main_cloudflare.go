//go:build all || cloudflare
// +build all cloudflare

package main

import (
	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/provider/cloudflare"
)

func init() {
	zoneIDFilter := provider.NewZoneIDFilter(cfg.ZoneIDFilter)
	p, err := cloudflare.NewCloudFlareProvider(domainFilter, zoneIDFilter, cfg.CloudflareProxied, cfg.DryRun, cfg.CloudflareDNSRecordsPerPage)
	if err != nil {
		log.Fatal(err)
	}
	providerMap["cloudflare"] = p
}
