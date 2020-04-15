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
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"

	_ "k8s.io/client-go/plugin/pkg/client/auth"

	"sigs.k8s.io/external-dns/controller"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
	"sigs.k8s.io/external-dns/pkg/apis/externaldns/validation"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/registry"
	"sigs.k8s.io/external-dns/source"
)

func main() {
	cfg := externaldns.NewConfig()
	if err := cfg.ParseFlags(os.Args[1:]); err != nil {
		log.Fatalf("flag parsing error: %v", err)
	}
	log.Infof("config: %s", cfg)

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

	ctx := context.Background()

	stopChan := make(chan struct{}, 1)

	go serveMetrics(cfg.MetricsAddress)
	go handleSigterm(stopChan)

	// Create a source.Config from the flags passed by the user.
	sourceCfg := &source.Config{
		Namespace:                      cfg.Namespace,
		AnnotationFilter:               cfg.AnnotationFilter,
		FQDNTemplate:                   cfg.FQDNTemplate,
		CombineFQDNAndAnnotation:       cfg.CombineFQDNAndAnnotation,
		IgnoreHostnameAnnotation:       cfg.IgnoreHostnameAnnotation,
		Compatibility:                  cfg.Compatibility,
		PublishInternal:                cfg.PublishInternal,
		PublishHostIP:                  cfg.PublishHostIP,
		AlwaysPublishNotReadyAddresses: cfg.AlwaysPublishNotReadyAddresses,
		ConnectorServer:                cfg.ConnectorSourceServer,
		CRDSourceAPIVersion:            cfg.CRDSourceAPIVersion,
		CRDSourceKind:                  cfg.CRDSourceKind,
		KubeConfig:                     cfg.KubeConfig,
		KubeMaster:                     cfg.Master,
		ServiceTypeFilter:              cfg.ServiceTypeFilter,
		IstioIngressGatewayServices:    cfg.IstioIngressGatewayServices,
		CFAPIEndpoint:                  cfg.CFAPIEndpoint,
		CFUsername:                     cfg.CFUsername,
		CFPassword:                     cfg.CFPassword,
		ContourLoadBalancerService:     cfg.ContourLoadBalancerService,
		SkipperRouteGroupVersion:       cfg.SkipperRouteGroupVersion,
		RequestTimeout:                 cfg.RequestTimeout,
	}

	// Lookup all the selected sources by names and pass them the desired configuration.
	sources, err := source.ByNames(&source.SingletonClientGenerator{
		KubeConfig: cfg.KubeConfig,
		KubeMaster: cfg.Master,
		// If update events are enabled, disable timeout.
		RequestTimeout: func() time.Duration {
			if cfg.UpdateEvents {
				return 0
			}
			return cfg.RequestTimeout
		}(),
	}, cfg.Sources, sourceCfg)
	if err != nil {
		log.Fatal(err)
	}

	// Combine multiple sources into a single, deduplicated source.
	endpointsSource := source.NewDedupSource(source.NewMultiSource(sources))

	domainFilter := endpoint.NewDomainFilterWithExclusions(cfg.DomainFilter, cfg.ExcludeDomains)
	zoneIDFilter := provider.NewZoneIDFilter(cfg.ZoneIDFilter)
	zoneTypeFilter := provider.NewZoneTypeFilter(cfg.AWSZoneType)
	zoneTagFilter := provider.NewZoneTagFilter(cfg.AWSZoneTagFilter)

	var p provider.Provider
	switch cfg.Provider {
	case "akamai":
		p = provider.NewAkamaiProvider(
			provider.AkamaiConfig{
				DomainFilter:          domainFilter,
				ZoneIDFilter:          zoneIDFilter,
				ServiceConsumerDomain: cfg.AkamaiServiceConsumerDomain,
				ClientToken:           cfg.AkamaiClientToken,
				ClientSecret:          cfg.AkamaiClientSecret,
				AccessToken:           cfg.AkamaiAccessToken,
				DryRun:                cfg.DryRun,
			},
		)
	case "alibabacloud":
		p, err = provider.NewAlibabaCloudProvider(cfg.AlibabaCloudConfigFile, domainFilter, zoneIDFilter, cfg.AlibabaCloudZoneType, cfg.DryRun)
	case "aws":
		p, err = provider.NewAWSProvider(
			provider.AWSConfig{
				DomainFilter:         domainFilter,
				ZoneIDFilter:         zoneIDFilter,
				ZoneTypeFilter:       zoneTypeFilter,
				ZoneTagFilter:        zoneTagFilter,
				BatchChangeSize:      cfg.AWSBatchChangeSize,
				BatchChangeInterval:  cfg.AWSBatchChangeInterval,
				EvaluateTargetHealth: cfg.AWSEvaluateTargetHealth,
				AssumeRole:           cfg.AWSAssumeRole,
				APIRetries:           cfg.AWSAPIRetries,
				PreferCNAME:          cfg.AWSPreferCNAME,
				DryRun:               cfg.DryRun,
			},
		)
	case "aws-sd":
		// Check that only compatible Registry is used with AWS-SD
		if cfg.Registry != "noop" && cfg.Registry != "aws-sd" {
			log.Infof("Registry \"%s\" cannot be used with AWS Cloud Map. Switching to \"aws-sd\".", cfg.Registry)
			cfg.Registry = "aws-sd"
		}
		p, err = provider.NewAWSSDProvider(domainFilter, cfg.AWSZoneType, cfg.AWSAssumeRole, cfg.DryRun)
	case "azure-dns", "azure":
		p, err = provider.NewAzureProvider(cfg.AzureConfigFile, domainFilter, zoneIDFilter, cfg.AzureResourceGroup, cfg.AzureUserAssignedIdentityClientID, cfg.DryRun)
	case "azure-private-dns":
		p, err = provider.NewAzurePrivateDNSProvider(domainFilter, zoneIDFilter, cfg.AzureResourceGroup, cfg.AzureSubscriptionID, cfg.DryRun)
	case "vinyldns":
		p, err = provider.NewVinylDNSProvider(domainFilter, zoneIDFilter, cfg.DryRun)
	case "cloudflare":
		p, err = provider.NewCloudFlareProvider(domainFilter, zoneIDFilter, cfg.CloudflareZonesPerPage, cfg.CloudflareProxied, cfg.DryRun)
	case "rcodezero":
		p, err = provider.NewRcodeZeroProvider(domainFilter, cfg.DryRun, cfg.RcodezeroTXTEncrypt)
	case "google":
		p, err = provider.NewGoogleProvider(ctx, cfg.GoogleProject, domainFilter, zoneIDFilter, cfg.GoogleBatchChangeSize, cfg.GoogleBatchChangeInterval, cfg.DryRun)
	case "digitalocean":
		p, err = provider.NewDigitalOceanProvider(ctx, domainFilter, cfg.DryRun)
	case "ovh":
		p, err = provider.NewOVHProvider(ctx, domainFilter, cfg.OVHEndpoint, cfg.DryRun)
	case "linode":
		p, err = provider.NewLinodeProvider(domainFilter, cfg.DryRun, externaldns.Version)
	case "dnsimple":
		p, err = provider.NewDnsimpleProvider(domainFilter, zoneIDFilter, cfg.DryRun)
	case "infoblox":
		p, err = provider.NewInfobloxProvider(
			provider.InfobloxConfig{
				DomainFilter: domainFilter,
				ZoneIDFilter: zoneIDFilter,
				Host:         cfg.InfobloxGridHost,
				Port:         cfg.InfobloxWapiPort,
				Username:     cfg.InfobloxWapiUsername,
				Password:     cfg.InfobloxWapiPassword,
				Version:      cfg.InfobloxWapiVersion,
				SSLVerify:    cfg.InfobloxSSLVerify,
				View:         cfg.InfobloxView,
				MaxResults:   cfg.InfobloxMaxResults,
				DryRun:       cfg.DryRun,
			},
		)
	case "dyn":
		p, err = provider.NewDynProvider(
			provider.DynConfig{
				DomainFilter:  domainFilter,
				ZoneIDFilter:  zoneIDFilter,
				DryRun:        cfg.DryRun,
				CustomerName:  cfg.DynCustomerName,
				Username:      cfg.DynUsername,
				Password:      cfg.DynPassword,
				MinTTLSeconds: cfg.DynMinTTLSeconds,
				AppVersion:    externaldns.Version,
			},
		)
	case "coredns", "skydns":
		p, err = provider.NewCoreDNSProvider(domainFilter, cfg.CoreDNSPrefix, cfg.DryRun)
	case "rdns":
		p, err = provider.NewRDNSProvider(
			provider.RDNSConfig{
				DomainFilter: domainFilter,
				DryRun:       cfg.DryRun,
			},
		)
	case "exoscale":
		p, err = provider.NewExoscaleProvider(cfg.ExoscaleEndpoint, cfg.ExoscaleAPIKey, cfg.ExoscaleAPISecret, cfg.DryRun, provider.ExoscaleWithDomain(domainFilter), provider.ExoscaleWithLogging()), nil
	case "inmemory":
		p, err = provider.NewInMemoryProvider(provider.InMemoryInitZones(cfg.InMemoryZones), provider.InMemoryWithDomain(domainFilter), provider.InMemoryWithLogging()), nil
	case "designate":
		p, err = provider.NewDesignateProvider(domainFilter, cfg.DryRun)
	case "pdns":
		p, err = provider.NewPDNSProvider(
			ctx,
			provider.PDNSConfig{
				DomainFilter: domainFilter,
				DryRun:       cfg.DryRun,
				Server:       cfg.PDNSServer,
				APIKey:       cfg.PDNSAPIKey,
				TLSConfig: provider.TLSConfig{
					TLSEnabled:            cfg.PDNSTLSEnabled,
					CAFilePath:            cfg.TLSCA,
					ClientCertFilePath:    cfg.TLSClientCert,
					ClientCertKeyFilePath: cfg.TLSClientCertKey,
				},
			},
		)
	case "oci":
		var config *provider.OCIConfig
		config, err = provider.LoadOCIConfig(cfg.OCIConfigFile)
		if err == nil {
			p, err = provider.NewOCIProvider(*config, domainFilter, zoneIDFilter, cfg.DryRun)
		}
	case "rfc2136":
		p, err = provider.NewRfc2136Provider(cfg.RFC2136Host, cfg.RFC2136Port, cfg.RFC2136Zone, cfg.RFC2136Insecure, cfg.RFC2136TSIGKeyName, cfg.RFC2136TSIGSecret, cfg.RFC2136TSIGSecretAlg, cfg.RFC2136TAXFR, domainFilter, cfg.DryRun, cfg.RFC2136MinTTL, nil)
	case "ns1":
		p, err = provider.NewNS1Provider(
			provider.NS1Config{
				DomainFilter: domainFilter,
				ZoneIDFilter: zoneIDFilter,
				NS1Endpoint:  cfg.NS1Endpoint,
				NS1IgnoreSSL: cfg.NS1IgnoreSSL,
				DryRun:       cfg.DryRun,
			},
		)
	case "transip":
		p, err = provider.NewTransIPProvider(cfg.TransIPAccountName, cfg.TransIPPrivateKeyFile, domainFilter, cfg.DryRun)
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
		r, err = registry.NewTXTRegistry(p, cfg.TXTPrefix, cfg.TXTOwnerID, cfg.TXTCacheInterval)
	case "aws-sd":
		r, err = registry.NewAWSSDRegistry(p.(*provider.AWSSDProvider), cfg.TXTOwnerID)
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
		Source:       endpointsSource,
		Registry:     r,
		Policy:       policy,
		Interval:     cfg.Interval,
		DomainFilter: domainFilter,
	}

	if cfg.UpdateEvents {
		// Add RunOnce as the handler function that will be called when ingress/service sources have changed.
		// Note that k8s Informers will perform an initial list operation, which results in the handler
		// function initially being called for every Service/Ingress that exists limted by minInterval.
		ctrl.Source.AddEventHandler(func() error { return ctrl.RunOnce(ctx) }, stopChan, 1*time.Minute)
	}

	if cfg.Once {
		err := ctrl.RunOnce(ctx)
		if err != nil {
			log.Fatal(err)
		}

		os.Exit(0)
	}
	ctrl.Run(ctx, stopChan)
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
		_, _ = w.Write([]byte("OK"))
	})

	http.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(address, nil))
}
