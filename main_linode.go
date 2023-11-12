//go:build all || linode
// +build all linode

package main

import (
	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
	"sigs.k8s.io/external-dns/provider/linode"
)

func init() {
	if cfg.Provider == "linode" {
		p, err := linode.NewLinodeProvider(domainFilter, cfg.DryRun, externaldns.Version)
		if err != nil {
			log.Fatal(err)
		}
		providerMap[cfg.Provider] = p
	}
}
