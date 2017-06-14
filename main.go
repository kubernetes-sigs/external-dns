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
	"strings"
	"syscall"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/linki/instrumented_http"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

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
	if cfg.Debug {
		log.SetLevel(log.DebugLevel)
	}

	stopChan := make(chan struct{}, 1)

	go serveMetrics(cfg.MetricsAddress)
	go handleSigterm(stopChan)

	var client *kubernetes.Clientset

	// create only those services we explicitly ask for in cfg.Sources
	for _, sourceType := range cfg.Sources {
		// we only need a k8s client if we're creating a non-fake source, and
		// have not already instantiated a k8s client
		if sourceType != "fake" && client == nil {
			var err error
			client, err = newClient(cfg)
			if err != nil {
				log.Fatal(err)
			}
		}

		var src source.Source
		var err error
		switch sourceType {
		case "fake":
			src, err = source.NewFakeSource(cfg.FqdnTemplate)
		case "service":
			src, err = source.NewServiceSource(client, cfg.Namespace, cfg.FqdnTemplate, cfg.Compatibility)
		case "ingress":
			src, err = source.NewIngressSource(client, cfg.Namespace, cfg.FqdnTemplate)
		default:
			log.Fatalf("Don't know how to handle sourceType '%s'", sourceType)
		}

		if err != nil {
			log.Fatal(err)
		}

		source.Register(sourceType, src)
	}

	sources, err := source.LookupMultiple(cfg.Sources)
	if err != nil {
		log.Fatal(err)
	}

	endpointsSource := source.NewDedupSource(source.NewMultiSource(sources))

	var p provider.Provider
	switch cfg.Provider {
	case "google":
		p, err = provider.NewGoogleProvider(cfg.GoogleProject, cfg.DomainFilter, cfg.DryRun)
	case "aws":
		p, err = provider.NewAWSProvider(cfg.DomainFilter, cfg.DryRun)
	case "digitalocean":
		p, err = provider.NewDigitalOceanProvider(cfg.DomainFilter, cfg.DryRun)
	case "inmemory":
		p, err = provider.NewInMemoryProviderWithDomainAndLogging("example.com"), nil
	case "azure":
		p, err = provider.NewAzureProvider(cfg.AzureConfigFile, cfg.DomainFilter, cfg.AzureResourceGroup, cfg.DryRun)
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

func newClient(cfg *externaldns.Config) (*kubernetes.Clientset, error) {
	if cfg.KubeConfig == "" {
		if _, err := os.Stat(clientcmd.RecommendedHomeFile); err == nil {
			cfg.KubeConfig = clientcmd.RecommendedHomeFile
		}
	}

	config, err := clientcmd.BuildConfigFromFlags(cfg.Master, cfg.KubeConfig)
	if err != nil {
		return nil, err
	}

	config.WrapTransport = func(rt http.RoundTripper) http.RoundTripper {
		return instrumented_http.NewTransport(rt, &instrumented_http.Callbacks{
			PathProcessor: func(path string) string {
				parts := strings.Split(path, "/")
				return parts[len(parts)-1]
			},
		})
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	log.Infof("Connected to cluster at %s", config.Host)

	return client, nil
}

func serveMetrics(address string) {
	http.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	http.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(address, nil))
}
