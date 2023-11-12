//go:build all || bluecat
// +build all bluecat

package main

import (
	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/provider/bluecat"
)

func init() {
	if cfg.Provider == "bluecat" {
		zoneIDFilter := provider.NewZoneIDFilter(cfg.ZoneIDFilter)
		p, err := bluecat.NewBluecatProvider(cfg.BluecatConfigFile, cfg.BluecatDNSConfiguration, cfg.BluecatDNSServerName, cfg.BluecatDNSDeployType, cfg.BluecatDNSView, cfg.BluecatGatewayHost, cfg.BluecatRootZone, cfg.TXTPrefix, cfg.TXTSuffix, domainFilter, zoneIDFilter, cfg.DryRun, cfg.BluecatSkipTLSVerify)
		if err != nil {
			log.Fatal(err)
		}
		providerMap[cfg.Provider] = p
	}
}
