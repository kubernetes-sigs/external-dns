//go:build all || pihole
// +build all pihole

package main

import (
	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/provider/pihole"
)

func init() {
	p, err := pihole.NewPiholeProvider(
		pihole.PiholeConfig{
			Server:                cfg.PiholeServer,
			Password:              cfg.PiholePassword,
			TLSInsecureSkipVerify: cfg.PiholeTLSInsecureSkipVerify,
			DomainFilter:          domainFilter,
			DryRun:                cfg.DryRun,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	providerMap["pihole"] = p
}
