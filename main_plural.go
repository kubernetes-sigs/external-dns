//go:build all || plural
// +build all plural

package main

import (
	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/provider/plural"
)

func init() {
	p, err := plural.NewPluralProvider(cfg.PluralCluster, cfg.PluralProvider)
	if err != nil {
		log.Fatal(err)
	}
	providerMap["plural"] = p
}
