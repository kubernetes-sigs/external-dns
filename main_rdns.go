//go:build all || rdns
// +build all rdns

package main

import (
	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/provider/rdns"
)

func init() {
	p, err := rdns.NewRDNSProvider(
		rdns.RDNSConfig{
			DomainFilter: domainFilter,
			DryRun:       cfg.DryRun,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	providerMap["rdns"] = p
}
