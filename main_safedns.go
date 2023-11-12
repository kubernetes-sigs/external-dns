//go:build all || safedns
// +build all safedns

package main

import (
	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/provider/safedns"
)

func init() {
	if cfg.Provider == "safedns" {
		p, err := safedns.NewSafeDNSProvider(domainFilter, cfg.DryRun)
		if err != nil {
			log.Fatal(err)
		}
		providerMap[cfg.Provider] = p
	}
}
