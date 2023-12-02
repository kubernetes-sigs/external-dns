//go:build all || inmemory
// +build all inmemory

package main

import (
	"sigs.k8s.io/external-dns/provider/inmemory"
)

func init() {
	if cfg.Provider == "inmemory" {
		providerMap[cfg.Provider] = inmemory.NewInMemoryProvider(inmemory.InMemoryInitZones(cfg.InMemoryZones), inmemory.InMemoryWithDomain(domainFilter), inmemory.InMemoryWithLogging())
	}
}
