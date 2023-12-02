//go:build all || transip
// +build all transip

package main

import (
	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/provider/transip"
)

func init() {
	if cfg.Provider == "transip" {
		p, err := transip.NewTransIPProvider(cfg.TransIPAccountName, cfg.TransIPPrivateKeyFile, domainFilter, cfg.DryRun)
		if err != nil {
			log.Fatal(err)
		}
		providerMap[cfg.Provider] = p
	}
}
