//go:build all || ultradns
// +build all ultradns

package main

import (
	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/provider/ultradns"
)

func init() {
	if cfg.Provider == "ultradns" {
		p, err := ultradns.NewUltraDNSProvider(domainFilter, cfg.DryRun)
		if err != nil {
			log.Fatal(err)
		}
		providerMap[cfg.Provider] = p
	}
}
