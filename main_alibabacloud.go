//go:build all || alibabacloud
// +build all alibabacloud

package main

import (
	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/provider/alibabacloud"
)

func init() {
	if cfg.Provider == "alibabacloud" {
		zoneIDFilter := provider.NewZoneIDFilter(cfg.ZoneIDFilter)

		p, err := alibabacloud.NewAlibabaCloudProvider(cfg.AlibabaCloudConfigFile, domainFilter, zoneIDFilter, cfg.AlibabaCloudZoneType, cfg.DryRun)
		if err != nil {
			log.Fatal(err)
		}
		providerMap[cfg.Provider] = p
	}
}
