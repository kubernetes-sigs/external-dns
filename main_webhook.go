//go:build all || webhook
// +build all webhook

package main

import (
	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/provider/webhook"
)

func init() {
	if cfg.Provider == "webhook" {
		p, err := webhook.NewWebhookProvider(cfg.WebhookProviderURL)
		if err != nil {
			log.Fatal(err)
		}
		providerMap[cfg.Provider] = p
	}
}
