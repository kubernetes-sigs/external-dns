//go:build all || dyn
// +build all dyn

package main

import (
	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/provider/dyn"
)

func init() {
	if cfg.Provider == "dyn" {
		zoneIDFilter := provider.NewZoneIDFilter(cfg.ZoneIDFilter)
		p, err := dyn.NewDynProvider(
			dyn.DynConfig{
				DomainFilter:  domainFilter,
				ZoneIDFilter:  zoneIDFilter,
				DryRun:        cfg.DryRun,
				CustomerName:  cfg.DynCustomerName,
				Username:      cfg.DynUsername,
				Password:      cfg.DynPassword,
				MinTTLSeconds: cfg.DynMinTTLSeconds,
				AppVersion:    externaldns.Version,
			},
		)
		if err != nil {
			log.Fatal(err)
		}
		providerMap[cfg.Provider] = p
	}
}
