/*
Copyright 2025 The Kubernetes Authors.

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

package controller

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-logr/logr"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"k8s.io/klog/v2"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
	"sigs.k8s.io/external-dns/pkg/apis/externaldns/validation"
	"sigs.k8s.io/external-dns/pkg/events"
	"sigs.k8s.io/external-dns/pkg/metrics"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
	webhookapi "sigs.k8s.io/external-dns/provider/webhook/api"
	"sigs.k8s.io/external-dns/registry"
	"sigs.k8s.io/external-dns/source"
	"sigs.k8s.io/external-dns/source/annotations"
	"sigs.k8s.io/external-dns/source/wrappers"
)

func Execute() {
	cfg := externaldns.NewConfig()
	if err := cfg.ParseFlags(os.Args[1:]); err != nil {
		log.Fatalf("flag parsing error: %v", err)
	}
	log.Infof("config: %s", cfg)
	if err := validation.ValidateConfig(cfg); err != nil {
		log.Fatalf("config validation failed: %v", err)
	}

	// Set annotation prefix (required since init() was removed)
	annotations.SetAnnotationPrefix(cfg.AnnotationPrefix)
	if cfg.AnnotationPrefix != annotations.DefaultAnnotationPrefix {
		log.Infof("Using custom annotation prefix: %s", cfg.AnnotationPrefix)
	}

	configureLogger(cfg)

	if cfg.DryRun {
		log.Info("running in dry-run mode. No changes to DNS records will be made.")
	}

	if log.GetLevel() < log.DebugLevel {
		// Klog V2 is used by k8s.io/apimachinery/pkg/labels and can throw (a lot) of irrelevant logs
		// See https://github.com/kubernetes-sigs/external-dns/issues/2348
		defer klog.ClearLogger()
		klog.SetLogger(logr.Discard())
	}

	log.Info(externaldns.Banner())

	ctx, cancel := context.WithCancel(context.Background())

	go serveMetrics(cfg.MetricsAddress)
	go handleSigterm(cancel)

	sCfg := source.NewSourceConfig(cfg)
	// TODO: Move source construction to the source package (blocked by cyclic dependency with wrappers)
	endpointsSource, err := buildSource(ctx, sCfg)
	if err != nil {
		log.Fatal(err) // nolint: gocritic // exitAfterDefer
	}

	domainFilter := endpoint.NewDomainFilterWithOptions(
		endpoint.WithDomainFilter(cfg.DomainFilter),
		endpoint.WithDomainExclude(cfg.DomainExclude),
		endpoint.WithRegexDomainFilter(cfg.RegexDomainFilter),
		endpoint.WithRegexDomainExclude(cfg.RegexDomainExclude),
	)

	// TODO: Move provider construction to the provider package
	prvdr, err := buildProvider(ctx, cfg, domainFilter)
	if err != nil {
		log.Fatal(err)
	}

	if cfg.WebhookServer {
		webhookapi.StartHTTPApi(prvdr, nil, cfg.WebhookProviderReadTimeout, cfg.WebhookProviderWriteTimeout, "127.0.0.1:8888")
		os.Exit(0)
	}

	ctrl, err := buildController(ctx, cfg, endpointsSource, prvdr, domainFilter)
	if err != nil {
		log.Fatal(err)
	}

	if cfg.Once {
		err := ctrl.RunOnce(ctx)
		if err != nil {
			log.Fatal(err)
		}

		os.Exit(0)
	}

	if cfg.UpdateEvents {
		// Add RunOnce as the handler function that will be called when ingress/service sources have changed.
		// Note that k8s Informers will perform an initial list operation, which results in the handler
		// function initially being called for every Service/Ingress that exists
		ctrl.Source.AddEventHandler(ctx, func() { ctrl.ScheduleRunOnce(time.Now()) })
	}

	ctrl.ScheduleRunOnce(time.Now())
	ctrl.Run(ctx)
}

func buildProvider(
	ctx context.Context,
	cfg *externaldns.Config,
	domainFilter *endpoint.DomainFilter,
) (provider.Provider, error) {
	factory, ok := providerFactories[cfg.Provider]
	if !ok {
		return nil, fmt.Errorf("unknown dns provider: %s", cfg.Provider)
	}
	p, err := factory(ctx, cfg, domainFilter)
	if err != nil {
		return nil, err
	}

	if p != nil && cfg.ProviderCacheTime > 0 {
		p = provider.NewCachedProvider(
			p,
			cfg.ProviderCacheTime,
		)
	}
	return p, nil
}

func buildController(
	ctx context.Context,
	cfg *externaldns.Config,
	src source.Source,
	p provider.Provider,
	filter *endpoint.DomainFilter,
) (*Controller, error) {
	policy, ok := plan.Policies[cfg.Policy]
	if !ok {
		return nil, fmt.Errorf("unknown policy: %s", cfg.Policy)
	}
	reg, err := registry.SelectRegistry(cfg, p)
	if err != nil {
		return nil, err
	}
	eventsCfg := events.NewConfig(
		events.WithKubeConfig(cfg.KubeConfig, cfg.APIServerURL, cfg.RequestTimeout),
		events.WithEmitEvents(cfg.EmitEvents),
		events.WithDryRun(cfg.DryRun))
	var eventEmitter events.EventEmitter
	if eventsCfg.IsEnabled() {
		eventCtrl, err := events.NewEventController(eventsCfg)
		if err != nil {
			log.Fatal(err)
		}
		eventCtrl.Run(ctx)
		eventEmitter = eventCtrl
	}

	return &Controller{
		Source:               src,
		Registry:             reg,
		Policy:               policy,
		Interval:             cfg.Interval,
		DomainFilter:         filter,
		ManagedRecordTypes:   cfg.ManagedDNSRecordTypes,
		ExcludeRecordTypes:   cfg.ExcludeDNSRecordTypes,
		MinEventSyncInterval: cfg.MinEventSyncInterval,
		TXTOwnerOld:          cfg.TXTOwnerOld,
		EventEmitter:         eventEmitter,
	}, nil
}

// This function configures the logger format and level based on the provided configuration.
func configureLogger(cfg *externaldns.Config) {
	if cfg.LogFormat == "json" {
		log.SetFormatter(&log.JSONFormatter{})
	}
	ll, err := log.ParseLevel(cfg.LogLevel)
	if err != nil {
		log.Fatalf("failed to parse log level: %v", err)
	}
	log.SetLevel(ll)
}

// buildSource creates and configures the source(s) for endpoint discovery based on the provided configuration.
// It initializes the source configuration, generates the required sources, and combines them into a single,
// deduplicated source. Returns the combined source or an error if source creation fails.
func buildSource(ctx context.Context, cfg *source.Config) (source.Source, error) {
	sources, err := source.ByNames(ctx, cfg, cfg.ClientGenerator())
	if err != nil {
		return nil, err
	}
	opts := wrappers.NewConfig(
		wrappers.WithDefaultTargets(cfg.DefaultTargets),
		wrappers.WithForceDefaultTargets(cfg.ForceDefaultTargets),
		wrappers.WithNAT64Networks(cfg.NAT64Networks),
		wrappers.WithTargetNetFilter(cfg.TargetNetFilter),
		wrappers.WithExcludeTargetNets(cfg.ExcludeTargetNets),
		wrappers.WithMinTTL(cfg.MinTTL))
	return wrappers.WrapSources(sources, opts)
}

// handleSigterm listens for a SIGTERM signal and triggers the provided cancel function
// to gracefully terminate the application. It logs a message when the signal is received.
func handleSigterm(cancel func()) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM)
	<-signals
	log.Info("Received SIGTERM. Terminating...")
	cancel()
}

// serveMetrics starts an HTTP server that serves health and metrics endpoints.
// The /healthz endpoint returns a 200 OK status to indicate the service is healthy.
// The /metrics endpoint serves Prometheus metrics.
// The server listens on the specified address and logs debug information about the endpoints.
func serveMetrics(address string) {
	http.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	})

	log.Debugf("serving 'healthz' on '%s/healthz'", address)
	log.Debugf("serving 'metrics' on '%s/metrics'", address)
	log.Debugf("registered '%d' metrics", len(metrics.RegisterMetric.Metrics))

	http.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(address, nil))
}
