//go:build all || pdns
// +build all pdns

package main

import (
	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/provider/pdns"
)

func init() {
	if cfg.Provider == "pdns" {
		p, err := pdns.NewPDNSProvider(
			ctx,
			pdns.PDNSConfig{
				DomainFilter: domainFilter,
				DryRun:       cfg.DryRun,
				Server:       cfg.PDNSServer,
				APIKey:       cfg.PDNSAPIKey,
				TLSConfig: pdns.TLSConfig{
					SkipTLSVerify:         cfg.PDNSSkipTLSVerify,
					CAFilePath:            cfg.TLSCA,
					ClientCertFilePath:    cfg.TLSClientCert,
					ClientCertKeyFilePath: cfg.TLSClientCertKey,
				},
			},
		)
		if err != nil {
			log.Fatal(err)
		}
		providerMap[cfg.Provider] = p
	}
}
