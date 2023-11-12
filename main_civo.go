//go:build all || civo
// +build all civo

package main

import (
	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/provider/civo"
)

func init() {
	p, err := civo.NewCivoProvider(domainFilter, cfg.DryRun)
	if err != nil {
		log.Fatal(err)
	}
	providerMap["civo"] = p
}
