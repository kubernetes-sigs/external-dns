//go:build all || azureprivatedns
// +build all azureprivatedns

package main

import (
	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/provider/azure"
)

func init() {
	if cfg.Provider == "azure-private-dns" {
		zoneIDFilter := provider.NewZoneIDFilter(cfg.ZoneIDFilter)
		p, err := azure.NewAzurePrivateDNSProvider(cfg.AzureConfigFile, domainFilter, zoneIDFilter, cfg.AzureResourceGroup, cfg.AzureUserAssignedIdentityClientID, cfg.DryRun)
		if err != nil {
			log.Fatal(err)
		}
		providerMap[cfg.Provider] = p
	}
}
