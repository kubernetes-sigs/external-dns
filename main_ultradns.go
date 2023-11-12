//go:build all || ultradns
// +build all ultradns

package main

import (
	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/provider/ultradns"
)

func init() {
	p, err := ultradns.NewUltraDNSProvider(domainFilter, cfg.DryRun)
	if err != nil {
		log.Fatal(err)
	}
	providerMap["ultradns"] = p
}
