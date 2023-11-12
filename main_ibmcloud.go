//go:build all || ibmcloud
// +build all ibmcloud

package main

import (
	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/provider/ibmcloud"
	"sigs.k8s.io/external-dns/source"
)

func init() {
	zoneIDFilter := provider.NewZoneIDFilter(cfg.ZoneIDFilter)
	// Combine multiple sources into a single, deduplicated source.
	endpointsSource := source.NewDedupSource(source.NewMultiSource(sources, sourceCfg.DefaultTargets))
	endpointsSource = source.NewTargetFilterSource(endpointsSource, targetFilter)
	p, err := ibmcloud.NewIBMCloudProvider(cfg.IBMCloudConfigFile, domainFilter, zoneIDFilter, endpointsSource, cfg.IBMCloudProxied, cfg.DryRun)
	if err != nil {
		log.Fatal(err)
	}
	providerMap["ibmcloud"] = p
}
