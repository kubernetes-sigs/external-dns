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
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
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
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))

	cfg := externaldns.NewConfig()
	if err := cfg.ParseFlags(os.Args[1:]); err != nil {
		level.Error(logger).Log("msg", "failed to parse flags", "error", err)
		os.Exit(1)
	}
	level.Info(logger).Log("msg", "successfully parsed flags", "config", fmt.Sprintf("%+v", cfg))

	if err := validation.ValidateConfig(cfg); err != nil {
		level.Error(logger).Log("msg", "failed to validate config", "error", err)
		os.Exit(1)
	}

	if !cfg.Debug {
		logger = level.NewFilter(logger, level.AllowInfo())
	}
	logger = log.With(logger, "time", log.DefaultTimestampUTC, "caller", log.DefaultCaller)

	// if cfg.LogFormat == "json" {
	// 	logger.SetFormatter(&logger.JSONFormatter{})
	// }

	if cfg.DryRun {
		level.Info(logger).Log("msg", "running in dry-run mode. No changes to DNS records will be made.")
	}

	stopChan := make(chan struct{}, 1)

	go serveMetrics(cfg.MetricsAddress, logger)
	go handleSigterm(stopChan, logger)

	client, err := newClient(cfg, logger)
	if err != nil {
		level.Error(logger).Log("msg", "failed to initialize Kubernetes client", "error", err)
		os.Exit(1)
	}

	serviceSource, err := source.NewServiceSource(client, cfg.Namespace, cfg.FqdnTemplate, cfg.Compatibility)
	if err != nil {
		level.Error(logger).Log("msg", "failed to initialize source", "source", "Service", "error", err)
		os.Exit(1)
	}
	source.Register("service", serviceSource)

	ingressSource, err := source.NewIngressSource(client, cfg.Namespace, cfg.FqdnTemplate)
	if err != nil {
		level.Error(logger).Log("msg", "failed to initialize source", "source", "Ingress", "error", err)
		os.Exit(1)
	}
	source.Register("ingress", ingressSource)

	sources, err := source.LookupMultiple(cfg.Sources)
	if err != nil {
		level.Error(logger).Log("msg", "failed to initialize source", "source", "Multi", "error", err)
		os.Exit(1)
	}

	endpointsSource := source.NewDedupSource(source.NewMultiSource(sources), logger)

	var p provider.Provider
	switch cfg.Provider {
	case "google":
		p, err = provider.NewGoogleProvider(cfg.GoogleProject, cfg.DomainFilter, cfg.DryRun)
	case "aws":
		p, err = provider.NewAWSProvider(cfg.DomainFilter, cfg.DryRun)
	default:
		level.Error(logger).Log("msg", "unknown provider", "provider", cfg.Provider)
		os.Exit(1)
	}
	if err != nil {
		level.Error(logger).Log("msg", "failed to initialize provider", "provider", cfg.Provider, "error", err)
		os.Exit(1)
	}

	var r registry.Registry
	switch cfg.Registry {
	case "noop":
		r, err = registry.NewNoopRegistry(p)
	case "txt":
		r, err = registry.NewTXTRegistry(p, cfg.TXTPrefix, cfg.TXTOwnerID)
	default:
		level.Error(logger).Log("msg", "unknown registry", "registry", cfg.Registry)
		os.Exit(1)
	}
	if err != nil {
		level.Error(logger).Log("msg", "failed to initialize registry", "registry", cfg.Registry, "error", err)
		os.Exit(1)
	}

	policy, exists := plan.Policies[cfg.Policy]
	if !exists {
		level.Error(logger).Log("msg", "unknown policy", "policy", cfg.Policy)
		os.Exit(1)
	}

	ctrl := controller.Controller{
		Source:   endpointsSource,
		Registry: r,
		Policy:   policy,
		Interval: cfg.Interval,
	}

	if cfg.Once {
		if err := ctrl.RunOnce(); err != nil {
			level.Error(logger).Log("msg", "failed to run control loop", "error", err)
			os.Exit(1)
		}

		os.Exit(0)
	}

	ctrl.Run(stopChan)
	for {
		level.Info(logger).Log("msg", "Pod waiting to be deleted")
		time.Sleep(time.Second * 30)
	}
}

func handleSigterm(stopChan chan struct{}, logger log.Logger) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM)
	<-signals
	level.Info(logger).Log("msg", "Received SIGTERM. Terminating...")
	close(stopChan)
}

func newClient(cfg *externaldns.Config, logger log.Logger) (*kubernetes.Clientset, error) {
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

	level.Info(logger).Log("msg", "Connected to Kubernetes cluster", "host", config.Host)

	return client, nil
}

func serveMetrics(address string, logger log.Logger) {
	http.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	http.Handle("/metrics", promhttp.Handler())

	err := http.ListenAndServe(address, nil)
	if err != nil {
		level.Error(logger).Log("msg", "failed to serve metrics", "error", err)
		os.Exit(1)
	}
}
