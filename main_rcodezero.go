//go:build all || rcodezero
// +build all rcodezero

package main

import (
	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/provider/rcode0"
)

func init() {
	p, err := rcode0.NewRcodeZeroProvider(domainFilter, cfg.DryRun, cfg.RcodezeroTXTEncrypt)
	if err != nil {
		log.Fatal(err)
	}
	providerMap["rcodezero"] = p
}
