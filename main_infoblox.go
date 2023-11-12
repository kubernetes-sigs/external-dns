//go:build all || infoblox
// +build all infoblox

package main

import (
	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/provider/infoblox"
)

func init() {
	if cfg.Provider == "infoblox" {
		zoneIDFilter := provider.NewZoneIDFilter(cfg.ZoneIDFilter)
		p, err := infoblox.NewInfobloxProvider(
			infoblox.StartupConfig{
				DomainFilter:  domainFilter,
				ZoneIDFilter:  zoneIDFilter,
				Host:          cfg.InfobloxGridHost,
				Port:          cfg.InfobloxWapiPort,
				Username:      cfg.InfobloxWapiUsername,
				Password:      cfg.InfobloxWapiPassword,
				Version:       cfg.InfobloxWapiVersion,
				SSLVerify:     cfg.InfobloxSSLVerify,
				View:          cfg.InfobloxView,
				MaxResults:    cfg.InfobloxMaxResults,
				DryRun:        cfg.DryRun,
				FQDNRegEx:     cfg.InfobloxFQDNRegEx,
				NameRegEx:     cfg.InfobloxNameRegEx,
				CreatePTR:     cfg.InfobloxCreatePTR,
				CacheDuration: cfg.InfobloxCacheDuration,
			},
		)
		if err != nil {
			log.Fatal(err)
		}
		providerMap[cfg.Provider] = p
	}
}
