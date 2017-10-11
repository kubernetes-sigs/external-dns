/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"

	_ "k8s.io/client-go/plugin/pkg/client/auth"

	"github.com/kubernetes-incubator/external-dns/controller"
	"github.com/kubernetes-incubator/external-dns/pkg/apis/externaldns"
	"github.com/kubernetes-incubator/external-dns/pkg/apis/externaldns/validation"
	"github.com/kubernetes-incubator/external-dns/plan"
	"github.com/kubernetes-incubator/external-dns/provider"
	"github.com/kubernetes-incubator/external-dns/registry"
	"github.com/kubernetes-incubator/external-dns/source"
)

func main() {
	cfg := externaldns.NewConfig()
	if err := cfg.ParseFlags(os.Args[1:]); err != nil {
		log.Fatalf("flag parsing error: %v", err)
	}
	log.Infof("config: %+v", cfg)

	if err := validation.ValidateConfig(cfg); err != nil {
		log.Fatalf("config validation failed: %v", err)
	}

	if cfg.LogFormat == "json" {
		log.SetFormatter(&log.JSONFormatter{})
	}
	if cfg.DryRun {
		log.Info("running in dry-run mode. No changes to DNS records will be made.")
	}

	ll, err := log.ParseLevel(cfg.LogLevel)
	if err != nil {
		log.Fatalf("failed to parse log level: %v", err)
	}
	log.SetLevel(ll)

	stopChan := make(chan struct{}, 1)

	go serveMetrics(cfg.MetricsAddress)
	go handleSigterm(stopChan)

	// Create a source.Config from the flags passed by the user.
	sourceCfg := &source.Config{
		Namespace:       cfg.Namespace,
		FQDNTemplate:    cfg.FQDNTemplate,
		Compatibility:   cfg.Compatibility,
		PublishInternal: cfg.PublishInternal,
	}

	// Lookup all the selected sources by names and pass them the desired configuration.
	sources, err := source.ByNames(&source.SingletonClientGenerator{
		KubeConfig: cfg.KubeConfig,
		KubeMaster: cfg.Master,
	}, cfg.Sources, sourceCfg)
	if err != nil {
		log.Fatal(err)
	}

	// Combine multiple sources into a single, deduplicated source.
	endpointsSource := source.NewDedupSource(source.NewMultiSource(sources))

	domainFilter := provider.NewDomainFilter(cfg.DomainFilter)
	zoneTypeFilter := provider.NewZoneTypeFilter(cfg.AWSZoneType)

	var p provider.Provider
	switch cfg.Provider {
	case "aws":
		p, err = provider.NewAWSProvider(domainFilter, zoneTypeFilter, cfg.DryRun)
	case "azure":
		p, err = provider.NewAzureProvider(cfg.AzureConfigFile, domainFilter, cfg.AzureResourceGroup, cfg.DryRun)
	case "cloudflare":
		p, err = provider.NewCloudFlareProvider(domainFilter, cfg.CloudflareProxied, cfg.DryRun)
	case "google":
		p, err = provider.NewGoogleProvider(cfg.GoogleProject, domainFilter, cfg.DryRun)
	case "digitalocean":
		p, err = provider.NewDigitalOceanProvider(domainFilter, cfg.DryRun)
	case "dnsimple":
		p, err = provider.NewDnsimpleProvider(domainFilter, cfg.DryRun)
	case "inmemory":
		p, err = provider.NewInMemoryProvider(provider.InMemoryWithDomain(domainFilter), provider.InMemoryWithLogging()), nil
	default:
		log.Fatalf("unknown dns provider: %s", cfg.Provider)
	}
	if err != nil {
		log.Fatal(err)
	}

	var r registry.Registry
	switch cfg.Registry {
	case "noop":
		r, err = registry.NewNoopRegistry(p)
	case "txt":
		r, err = registry.NewTXTRegistry(p, cfg.TXTPrefix, cfg.TXTOwnerID)
	default:
		log.Fatalf("unknown registry: %s", cfg.Registry)
	}

	if err != nil {
		log.Fatal(err)
	}

	policy, exists := plan.Policies[cfg.Policy]
	if !exists {
		log.Fatalf("unknown policy: %s", cfg.Policy)
	}

	ctrl := controller.Controller{
		Source:   endpointsSource,
		Registry: r,
		Policy:   policy,
		Interval: cfg.Interval,
	}

	if cfg.Once {
		err := ctrl.RunOnce()
		if err != nil {
			log.Fatal(err)
		}

		os.Exit(0)
	}

	ctrl.Run(stopChan)
	for {
		log.Info("Pod waiting to be deleted")
		time.Sleep(time.Second * 30)
	}
}

func handleSigterm(stopChan chan struct{}) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM)
	<-signals
	log.Info("Received SIGTERM. Terminating...")
	close(stopChan)
}

func serveMetrics(address string) {
	http.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	http.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(address, nil))
}
