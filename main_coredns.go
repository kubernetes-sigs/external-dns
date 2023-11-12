//go:build all || coredns
// +build all coredns

package main

import (
	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/provider/coredns"
)

func init() {
	p, err := coredns.NewCoreDNSProvider(domainFilter, cfg.CoreDNSPrefix, cfg.DryRun)
	if err != nil {
		log.Fatal(err)
	}
	providerMap["coredns"] = p
	providerMap["skydns"] = p
}
