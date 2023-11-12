//go:build all || tencentcloud
// +build all tencentcloud

package main

import (
	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/provider/tencentcloud"
)

func init() {
	if cfg.Provider == "tencentcloud" {
		zoneIDFilter := provider.NewZoneIDFilter(cfg.ZoneIDFilter)
		p, err := tencentcloud.NewTencentCloudProvider(domainFilter, zoneIDFilter, cfg.TencentCloudConfigFile, cfg.TencentCloudZoneType, cfg.DryRun)
		if err != nil {
			log.Fatal(err)
		}
		providerMap[cfg.Provider] = p
	}
}
