//go:build all || akamai
// +build all akamai

package main

import (
	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/provider/akamai"
)

func init() {
	if cfg.Provider == "akamai" {
		zoneIDFilter := provider.NewZoneIDFilter(cfg.ZoneIDFilter)

		p, err := akamai.NewAkamaiProvider(
			akamai.AkamaiConfig{
				DomainFilter:          domainFilter,
				ZoneIDFilter:          zoneIDFilter,
				ServiceConsumerDomain: cfg.AkamaiServiceConsumerDomain,
				ClientToken:           cfg.AkamaiClientToken,
				ClientSecret:          cfg.AkamaiClientSecret,
				AccessToken:           cfg.AkamaiAccessToken,
				EdgercPath:            cfg.AkamaiEdgercPath,
				EdgercSection:         cfg.AkamaiEdgercSection,
				DryRun:                cfg.DryRun,
			}, nil)
		if err != nil {
			log.Fatal(err)
		}

		providerMap[cfg.Provider] = p
	}
}
