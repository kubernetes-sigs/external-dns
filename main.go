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

	"github.com/kubernetes-incubator/external-dns/controller"
	"github.com/kubernetes-incubator/external-dns/pkg/apis/externaldns"
	"github.com/kubernetes-incubator/external-dns/pkg/apis/externaldns/validation"
	"github.com/kubernetes-incubator/external-dns/plan"
	"github.com/kubernetes-incubator/external-dns/provider"
	"github.com/kubernetes-incubator/external-dns/provider/alibaba"
	"github.com/kubernetes-incubator/external-dns/provider/aws"
	"github.com/kubernetes-incubator/external-dns/provider/awssd"
	"github.com/kubernetes-incubator/external-dns/provider/azure"
	"github.com/kubernetes-incubator/external-dns/provider/cloudflare"
	"github.com/kubernetes-incubator/external-dns/provider/coredns"
	"github.com/kubernetes-incubator/external-dns/provider/designate"
	"github.com/kubernetes-incubator/external-dns/provider/dnsimple"
	"github.com/kubernetes-incubator/external-dns/provider/do"
	"github.com/kubernetes-incubator/external-dns/provider/dyn"
	"github.com/kubernetes-incubator/external-dns/provider/exoscale"
	"github.com/kubernetes-incubator/external-dns/provider/google"
	"github.com/kubernetes-incubator/external-dns/provider/infoblox"
	"github.com/kubernetes-incubator/external-dns/provider/inmemory"
	"github.com/kubernetes-incubator/external-dns/provider/linode"
	"github.com/kubernetes-incubator/external-dns/provider/ns1"
	"github.com/kubernetes-incubator/external-dns/provider/oci"
	"github.com/kubernetes-incubator/external-dns/provider/pdns"
	"github.com/kubernetes-incubator/external-dns/provider/rancher"
	"github.com/kubernetes-incubator/external-dns/provider/rcode0"
	"github.com/kubernetes-incubator/external-dns/provider/rfc2136"
	"github.com/kubernetes-incubator/external-dns/provider/transip"
	"github.com/kubernetes-incubator/external-dns/provider/vinyldns"
	"github.com/kubernetes-incubator/external-dns/registry"
	"github.com/kubernetes-incubator/external-dns/source"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
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

	stopChan := make(chan struct{}, 1)

	go serveMetrics(cfg.MetricsAddress)
	go handleSigterm(stopChan)

	// Create a source.Config from the flags passed by the user.
	sourceCfg := &source.Config{
		Namespace:                   cfg.Namespace,
		AnnotationFilter:            cfg.AnnotationFilter,
		FQDNTemplate:                cfg.FQDNTemplate,
		CombineFQDNAndAnnotation:    cfg.CombineFQDNAndAnnotation,
		IgnoreHostnameAnnotation:    cfg.IgnoreHostnameAnnotation,
		Compatibility:               cfg.Compatibility,
		PublishInternal:             cfg.PublishInternal,
		PublishHostIP:               cfg.PublishHostIP,
		ConnectorServer:             cfg.ConnectorSourceServer,
		CRDSourceAPIVersion:         cfg.CRDSourceAPIVersion,
		CRDSourceKind:               cfg.CRDSourceKind,
		KubeConfig:                  cfg.KubeConfig,
		KubeMaster:                  cfg.Master,
		ServiceTypeFilter:           cfg.ServiceTypeFilter,
		IstioIngressGatewayServices: cfg.IstioIngressGatewayServices,
		CFAPIEndpoint:               cfg.CFAPIEndpoint,
		CFUsername:                  cfg.CFUsername,
		CFPassword:                  cfg.CFPassword,
	}

	// Lookup all the selected sources by names and pass them the desired configuration.
	sources, err := source.ByNames(&source.SingletonClientGenerator{
		KubeConfig:     cfg.KubeConfig,
		KubeMaster:     cfg.Master,
		RequestTimeout: cfg.RequestTimeout,
	}, cfg.Sources, sourceCfg)
	if err != nil {
		log.Fatal(err)
	}

	// Combine multiple sources into a single, deduplicated source.
	endpointsSource := source.NewDedupSource(source.NewMultiSource(sources))

	domainFilter := provider.NewDomainFilterWithExclusions(cfg.DomainFilter, cfg.ExcludeDomains)
	zoneIDFilter := provider.NewZoneIDFilter(cfg.ZoneIDFilter)
	zoneTypeFilter := provider.NewZoneTypeFilter(cfg.AWSZoneType)
	zoneTagFilter := provider.NewZoneTagFilter(cfg.AWSZoneTagFilter)

	var p provider.Provider
	switch cfg.Provider {
	case "alibabacloud":
		p, err = alibaba.NewAlibabaCloudProvider(cfg.AlibabaCloudConfigFile, domainFilter, zoneIDFilter, cfg.AlibabaCloudZoneType, cfg.DryRun)
	case "aws":
		p, err = aws.NewAWSProvider(
			aws.AWSConfig{
				DomainFilter:         domainFilter,
				ZoneIDFilter:         zoneIDFilter,
				ZoneTypeFilter:       zoneTypeFilter,
				ZoneTagFilter:        zoneTagFilter,
				BatchChangeSize:      cfg.AWSBatchChangeSize,
				BatchChangeInterval:  cfg.AWSBatchChangeInterval,
				EvaluateTargetHealth: cfg.AWSEvaluateTargetHealth,
				AssumeRole:           cfg.AWSAssumeRole,
				DryRun:               cfg.DryRun,
			},
		)
	case "aws-sd":
		// Check that only compatible Registry is used with AWS-SD
		if cfg.Registry != "noop" && cfg.Registry != "aws-sd" {
			log.Infof("Registry \"%s\" cannot be used with AWS ServiceDiscovery. Switching to \"aws-sd\".", cfg.Registry)
			cfg.Registry = "aws-sd"
		}
		p, err = awssd.NewAWSSDProvider(domainFilter, cfg.AWSZoneType, cfg.AWSAssumeRole, cfg.DryRun)
	case "azure":
		p, err = azure.NewAzureProvider(cfg.AzureConfigFile, domainFilter, zoneIDFilter, cfg.AzureResourceGroup, cfg.DryRun)
	case "vinyldns":
		p, err = vinyldns.NewVinylDNSProvider(domainFilter, zoneIDFilter, cfg.DryRun)
	case "cloudflare":
		p, err = cloudflare.NewCloudFlareProvider(domainFilter, zoneIDFilter, cfg.CloudflareZonesPerPage, cfg.CloudflareProxied, cfg.DryRun)
	case "rcodezero":
		p, err = rcode0.NewRcodeZeroProvider(domainFilter, cfg.DryRun, cfg.RcodezeroTXTEncrypt)
	case "google":
		p, err = google.NewGoogleProvider(cfg.GoogleProject, domainFilter, zoneIDFilter, cfg.DryRun)
	case "digitalocean":
		p, err = do.NewDigitalOceanProvider(domainFilter, cfg.DryRun)
	case "linode":
		p, err = linode.NewLinodeProvider(domainFilter, cfg.DryRun, externaldns.Version)
	case "dnsimple":
		p, err = dnsimple.NewDnsimpleProvider(domainFilter, zoneIDFilter, cfg.DryRun)
	case "infoblox":
		p, err = infoblox.NewInfobloxProvider(
			infoblox.InfobloxConfig{
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
		p, err = dyn.NewDynProvider(
			dyn.DynConfig{
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
		p, err = coredns.NewCoreDNSProvider(domainFilter, cfg.DryRun)
	case "rdns":
		p, err = rancher.NewRDNSProvider(
			rancher.RDNSConfig{
				DomainFilter: domainFilter,
				DryRun:       cfg.DryRun,
			},
		)
	case "exoscale":
		p, err = exoscale.NewExoscaleProvider(cfg.ExoscaleEndpoint, cfg.ExoscaleAPIKey, cfg.ExoscaleAPISecret, cfg.DryRun, exoscale.ExoscaleWithDomain(domainFilter), exoscale.ExoscaleWithLogging()), nil
	case "inmemory":
		p, err = inmemory.NewInMemoryProvider(inmemory.InMemoryInitZones(cfg.InMemoryZones), inmemory.InMemoryWithDomain(domainFilter), inmemory.InMemoryWithLogging()), nil
	case "designate":
		p, err = designate.NewDesignateProvider(domainFilter, cfg.DryRun)
	case "pdns":
		p, err = pdns.NewPDNSProvider(
			pdns.PDNSConfig{
				DomainFilter: domainFilter,
				DryRun:       cfg.DryRun,
				Server:       cfg.PDNSServer,
				APIKey:       cfg.PDNSAPIKey,
				TLSConfig: pdns.TLSConfig{
					TLSEnabled:            cfg.PDNSTLSEnabled,
					CAFilePath:            cfg.TLSCA,
					ClientCertFilePath:    cfg.TLSClientCert,
					ClientCertKeyFilePath: cfg.TLSClientCertKey,
				},
			},
		)
	case "oci":
		var config *oci.OCIConfig
		config, err = oci.LoadOCIConfig(cfg.OCIConfigFile)
		if err == nil {
			p, err = oci.NewOCIProvider(*config, domainFilter, zoneIDFilter, cfg.DryRun)
		}
	case "rfc2136":
		p, err = rfc2136.NewRfc2136Provider(cfg.RFC2136Host, cfg.RFC2136Port, cfg.RFC2136Zone, cfg.RFC2136Insecure, cfg.RFC2136TSIGKeyName, cfg.RFC2136TSIGSecret, cfg.RFC2136TSIGSecretAlg, cfg.RFC2136TAXFR, domainFilter, cfg.DryRun, nil)
	case "ns1":
		p, err = ns1.NewNS1Provider(
			ns1.NS1Config{
				DomainFilter: domainFilter,
				ZoneIDFilter: zoneIDFilter,
				NS1Endpoint:  cfg.NS1Endpoint,
				NS1IgnoreSSL: cfg.NS1IgnoreSSL,
				DryRun:       cfg.DryRun,
			},
		)
	case "transip":
		p, err = transip.NewTransIPProvider(cfg.TransIPAccountName, cfg.TransIPPrivateKeyFile, domainFilter, cfg.DryRun)
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
		r, err = registry.NewAWSSDRegistry(p.(*awssd.AWSSDProvider), cfg.TXTOwnerID)
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
