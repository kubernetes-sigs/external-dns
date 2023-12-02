//go:build all || designate
// +build all designate

package main

import (
	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/provider/designate"
)

func init() {
	if cfg.Provider == "designate" {
		p, err := designate.NewDesignateProvider(domainFilter, cfg.DryRun)
		if err != nil {
			log.Fatal(err)
		}
		providerMap[cfg.Provider] = p
	}
}
