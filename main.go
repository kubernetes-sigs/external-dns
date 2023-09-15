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
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	awsSDK "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/route53"
	sd "github.com/aws/aws-sdk-go/service/servicediscovery"
	"github.com/go-logr/logr"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/labels"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/klog/v2"

	"sigs.k8s.io/external-dns/controller"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
	"sigs.k8s.io/external-dns/pkg/apis/externaldns/validation"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/provider/akamai"
	"sigs.k8s.io/external-dns/provider/alibabacloud"
	"sigs.k8s.io/external-dns/provider/aws"
	"sigs.k8s.io/external-dns/provider/awssd"
	"sigs.k8s.io/external-dns/provider/azure"
	"sigs.k8s.io/external-dns/provider/bluecat"
	"sigs.k8s.io/external-dns/provider/civo"
	"sigs.k8s.io/external-dns/provider/cloudflare"
	"sigs.k8s.io/external-dns/provider/coredns"
	"sigs.k8s.io/external-dns/provider/designate"
	"sigs.k8s.io/external-dns/provider/digitalocean"
	"sigs.k8s.io/external-dns/provider/dnsimple"
	"sigs.k8s.io/external-dns/provider/dyn"
	"sigs.k8s.io/external-dns/provider/exoscale"
	"sigs.k8s.io/external-dns/provider/gandi"
	"sigs.k8s.io/external-dns/provider/godaddy"
	"sigs.k8s.io/external-dns/provider/google"
	"sigs.k8s.io/external-dns/provider/ibmcloud"
	"sigs.k8s.io/external-dns/provider/infoblox"
	"sigs.k8s.io/external-dns/provider/inmemory"
	"sigs.k8s.io/external-dns/provider/linode"
	"sigs.k8s.io/external-dns/provider/ns1"
	"sigs.k8s.io/external-dns/provider/oci"
	"sigs.k8s.io/external-dns/provider/ovh"
	"sigs.k8s.io/external-dns/provider/pdns"
	"sigs.k8s.io/external-dns/provider/pihole"
	"sigs.k8s.io/external-dns/provider/plural"
	"sigs.k8s.io/external-dns/provider/rcode0"
	"sigs.k8s.io/external-dns/provider/rdns"
	"sigs.k8s.io/external-dns/provider/rfc2136"
	"sigs.k8s.io/external-dns/provider/safedns"
	"sigs.k8s.io/external-dns/provider/scaleway"
	"sigs.k8s.io/external-dns/provider/tencentcloud"
	"sigs.k8s.io/external-dns/provider/transip"
	"sigs.k8s.io/external-dns/provider/ultradns"
	"sigs.k8s.io/external-dns/provider/vinyldns"
	"sigs.k8s.io/external-dns/provider/vultr"
	"sigs.k8s.io/external-dns/provider/webhook"
	"sigs.k8s.io/external-dns/registry"
	"sigs.k8s.io/external-dns/source"
)

func main() {
	cfg := externaldns.NewConfig()
	if err := cfg.ParseFlags(os.Args[1:]); err != nil {
		log.Fatalf("flag parsing error: %v", err)
	}
	if cfg.LogFormat == "json" {
		log.SetFormatter(&log.JSONFormatter{})
	}
	log.Infof("config: %s", cfg)

	if err := validation.ValidateConfig(cfg); err != nil {
		log.Fatalf("config validation failed: %v", err)
	}

	if cfg.DryRun {
		log.Info("running in dry-run mode. No changes to DNS records will be made.")
	}

	ll, err := log.ParseLevel(cfg.LogLevel)
	if err != nil {
		log.Fatalf("failed to parse log level: %v", err)
	}
	log.SetLevel(ll)

	// Klog V2 is used by k8s.io/apimachinery/pkg/labels and can throw (a lot) of irrelevant logs
	// See https://github.com/kubernetes-sigs/external-dns/issues/2348
	defer klog.ClearLogger()
	klog.SetLogger(logr.Discard())

	ctx, cancel := context.WithCancel(context.Background())

	go serveMetrics(cfg.MetricsAddress)
	go handleSigterm(cancel)

	// error is explicitly ignored because the filter is already validated in validation.ValidateConfig
	labelSelector, _ := labels.Parse(cfg.LabelFilter)

	// Create a source.Config from the flags passed by the user.
	sourceCfg := &source.Config{
		Namespace:                      cfg.Namespace,
		AnnotationFilter:               cfg.AnnotationFilter,
		LabelFilter:                    labelSelector,
		IngressClassNames:              cfg.IngressClassNames,
		FQDNTemplate:                   cfg.FQDNTemplate,
		CombineFQDNAndAnnotation:       cfg.CombineFQDNAndAnnotation,
		IgnoreHostnameAnnotation:       cfg.IgnoreHostnameAnnotation,
		IgnoreIngressTLSSpec:           cfg.IgnoreIngressTLSSpec,
		IgnoreIngressRulesSpec:         cfg.IgnoreIngressRulesSpec,
		GatewayNamespace:               cfg.GatewayNamespace,
		GatewayLabelFilter:             cfg.GatewayLabelFilter,
		Compatibility:                  cfg.Compatibility,
		PublishInternal:                cfg.PublishInternal,
		PublishHostIP:                  cfg.PublishHostIP,
		AlwaysPublishNotReadyAddresses: cfg.AlwaysPublishNotReadyAddresses,
		ConnectorServer:                cfg.ConnectorSourceServer,
		CRDSourceAPIVersion:            cfg.CRDSourceAPIVersion,
		CRDSourceKind:                  cfg.CRDSourceKind,
		KubeConfig:                     cfg.KubeConfig,
		APIServerURL:                   cfg.APIServerURL,
		ServiceTypeFilter:              cfg.ServiceTypeFilter,
		CFAPIEndpoint:                  cfg.CFAPIEndpoint,
		CFUsername:                     cfg.CFUsername,
		CFPassword:                     cfg.CFPassword,
		GlooNamespace:                  cfg.GlooNamespace,
		SkipperRouteGroupVersion:       cfg.SkipperRouteGroupVersion,
		RequestTimeout:                 cfg.RequestTimeout,
		DefaultTargets:                 cfg.DefaultTargets,
		OCPRouterName:                  cfg.OCPRouterName,
		UpdateEvents:                   cfg.UpdateEvents,
		ResolveLoadBalancerHostname:    cfg.ResolveServiceLoadBalancerHostname,
	}

	// Lookup all the selected sources by names and pass them the desired configuration.
	sources, err := source.ByNames(ctx, &source.SingletonClientGenerator{
		KubeConfig:   cfg.KubeConfig,
		APIServerURL: cfg.APIServerURL,
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

	// Filter targets
	targetFilter := endpoint.NewTargetNetFilterWithExclusions(cfg.TargetNetFilter, cfg.ExcludeTargetNets)

	// Combine multiple sources into a single, deduplicated source.
	endpointsSource := source.NewDedupSource(source.NewMultiSource(sources, sourceCfg.DefaultTargets))
	endpointsSource = source.NewTargetFilterSource(endpointsSource, targetFilter)

	// RegexDomainFilter overrides DomainFilter
	var domainFilter endpoint.DomainFilter
	if cfg.RegexDomainFilter.String() != "" {
		domainFilter = endpoint.NewRegexDomainFilter(cfg.RegexDomainFilter, cfg.RegexDomainExclusion)
	} else {
		domainFilter = endpoint.NewDomainFilterWithExclusions(cfg.DomainFilter, cfg.ExcludeDomains)
	}
	zoneNameFilter := endpoint.NewDomainFilter(cfg.ZoneNameFilter)
	zoneIDFilter := provider.NewZoneIDFilter(cfg.ZoneIDFilter)
	zoneTypeFilter := provider.NewZoneTypeFilter(cfg.AWSZoneType)
	zoneTagFilter := provider.NewZoneTagFilter(cfg.AWSZoneTagFilter)

	var awsSession *session.Session
	if cfg.Provider == "aws" || cfg.Provider == "aws-sd" || cfg.Registry == "dynamodb" || cfg.RunAWSProviderAsWebhook {
		awsSession, err = aws.NewSession(
			aws.AWSSessionConfig{
				AssumeRole:           cfg.AWSAssumeRole,
				AssumeRoleExternalID: cfg.AWSAssumeRoleExternalID,
				APIRetries:           cfg.AWSAPIRetries,
			},
		)
		if err != nil {
			log.Fatal(err)
		}
	}

	var p provider.Provider
	switch cfg.Provider {
	case "akamai":
		p, err = akamai.NewAkamaiProvider(
			akamai.AkamaiConfig{
				DomainFilter:          domainFilter,
				ZoneIDFilter:          zoneIDFilter,
				ServiceConsumerDomain: cfg.AkamaiServiceConsumerDomain,
				ClientToken:           cfg.AkamaiClientToken,
				ClientSecret:          cfg.AkamaiClientSecret,
				AccessToken:           cfg.AkamaiAccessToken,
				EdgercPath:            cfg.AkamaiEdgercPath,
				EdgercSection:         cfg.AkamaiEdgercSection,
				DryRun:                cfg.DryRun,
			}, nil)
	case "alibabacloud":
		p, err = alibabacloud.NewAlibabaCloudProvider(cfg.AlibabaCloudConfigFile, domainFilter, zoneIDFilter, cfg.AlibabaCloudZoneType, cfg.DryRun)
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
				PreferCNAME:          cfg.AWSPreferCNAME,
				DryRun:               cfg.DryRun,
				ZoneCacheDuration:    cfg.AWSZoneCacheDuration,
			},
			route53.New(awsSession),
		)
	case "aws-sd":
		// Check that only compatible Registry is used with AWS-SD
		if cfg.Registry != "noop" && cfg.Registry != "aws-sd" {
			log.Infof("Registry \"%s\" cannot be used with AWS Cloud Map. Switching to \"aws-sd\".", cfg.Registry)
			cfg.Registry = "aws-sd"
		}
		p, err = awssd.NewAWSSDProvider(domainFilter, cfg.AWSZoneType, cfg.DryRun, cfg.AWSSDServiceCleanup, cfg.TXTOwnerID, sd.New(awsSession))
	case "azure-dns", "azure":
		p, err = azure.NewAzureProvider(cfg.AzureConfigFile, domainFilter, zoneNameFilter, zoneIDFilter, cfg.AzureResourceGroup, cfg.AzureUserAssignedIdentityClientID, cfg.DryRun)
	case "azure-private-dns":
		p, err = azure.NewAzurePrivateDNSProvider(cfg.AzureConfigFile, domainFilter, zoneIDFilter, cfg.AzureResourceGroup, cfg.AzureUserAssignedIdentityClientID, cfg.DryRun)
	case "bluecat":
		p, err = bluecat.NewBluecatProvider(cfg.BluecatConfigFile, cfg.BluecatDNSConfiguration, cfg.BluecatDNSServerName, cfg.BluecatDNSDeployType, cfg.BluecatDNSView, cfg.BluecatGatewayHost, cfg.BluecatRootZone, cfg.TXTPrefix, cfg.TXTSuffix, domainFilter, zoneIDFilter, cfg.DryRun, cfg.BluecatSkipTLSVerify)
	case "vinyldns":
		p, err = vinyldns.NewVinylDNSProvider(domainFilter, zoneIDFilter, cfg.DryRun)
	case "vultr":
		p, err = vultr.NewVultrProvider(ctx, domainFilter, cfg.DryRun)
	case "ultradns":
		p, err = ultradns.NewUltraDNSProvider(domainFilter, cfg.DryRun)
	case "civo":
		p, err = civo.NewCivoProvider(domainFilter, cfg.DryRun)
	case "cloudflare":
		p, err = cloudflare.NewCloudFlareProvider(domainFilter, zoneIDFilter, cfg.CloudflareProxied, cfg.DryRun, cfg.CloudflareDNSRecordsPerPage)
	case "rcodezero":
		p, err = rcode0.NewRcodeZeroProvider(domainFilter, cfg.DryRun, cfg.RcodezeroTXTEncrypt)
	case "google":
		p, err = google.NewGoogleProvider(ctx, cfg.GoogleProject, domainFilter, zoneIDFilter, cfg.GoogleBatchChangeSize, cfg.GoogleBatchChangeInterval, cfg.GoogleZoneVisibility, cfg.DryRun)
	case "digitalocean":
		p, err = digitalocean.NewDigitalOceanProvider(ctx, domainFilter, cfg.DryRun, cfg.DigitalOceanAPIPageSize)
	case "ovh":
		p, err = ovh.NewOVHProvider(ctx, domainFilter, cfg.OVHEndpoint, cfg.OVHApiRateLimit, cfg.DryRun)
	case "linode":
		p, err = linode.NewLinodeProvider(domainFilter, cfg.DryRun, externaldns.Version)
	case "dnsimple":
		p, err = dnsimple.NewDnsimpleProvider(domainFilter, zoneIDFilter, cfg.DryRun)
	case "infoblox":
		p, err = infoblox.NewInfobloxProvider(
			infoblox.StartupConfig{
				DomainFilter:  domainFilter,
				ZoneIDFilter:  zoneIDFilter,
				Host:          cfg.InfobloxGridHost,
				Port:          cfg.InfobloxWapiPort,
				Username:      cfg.InfobloxWapiUsername,
				Password:      cfg.InfobloxWapiPassword,
				Version:       cfg.InfobloxWapiVersion,
				SSLVerify:     cfg.InfobloxSSLVerify,
				View:          cfg.InfobloxView,
				MaxResults:    cfg.InfobloxMaxResults,
				DryRun:        cfg.DryRun,
				FQDNRegEx:     cfg.InfobloxFQDNRegEx,
				NameRegEx:     cfg.InfobloxNameRegEx,
				CreatePTR:     cfg.InfobloxCreatePTR,
				CacheDuration: cfg.InfobloxCacheDuration,
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
		p, err = coredns.NewCoreDNSProvider(domainFilter, cfg.CoreDNSPrefix, cfg.DryRun)
	case "rdns":
		p, err = rdns.NewRDNSProvider(
			rdns.RDNSConfig{
				DomainFilter: domainFilter,
				DryRun:       cfg.DryRun,
			},
		)
	case "exoscale":
		p, err = exoscale.NewExoscaleProvider(
			cfg.ExoscaleAPIEnvironment,
			cfg.ExoscaleAPIZone,
			cfg.ExoscaleAPIKey,
			cfg.ExoscaleAPISecret,
			cfg.DryRun,
			exoscale.ExoscaleWithDomain(domainFilter),
			exoscale.ExoscaleWithLogging(),
		)
	case "inmemory":
		p, err = inmemory.NewInMemoryProvider(inmemory.InMemoryInitZones(cfg.InMemoryZones), inmemory.InMemoryWithDomain(domainFilter), inmemory.InMemoryWithLogging()), nil
	case "designate":
		p, err = designate.NewDesignateProvider(domainFilter, cfg.DryRun)
	case "pdns":
		p, err = pdns.NewPDNSProvider(
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
	case "oci":
		var config *oci.OCIConfig
		// if the instance-principals flag was set, and a compartment OCID was provided, then ignore the
		// OCI config file, and provide a config that uses instance principal authentication.
		if cfg.OCIAuthInstancePrincipal {
			if len(cfg.OCICompartmentOCID) == 0 {
				err = fmt.Errorf("instance principal authentication requested, but no compartment OCID provided")
			} else {
				authConfig := oci.OCIAuthConfig{UseInstancePrincipal: true}
				config = &oci.OCIConfig{Auth: authConfig, CompartmentID: cfg.OCICompartmentOCID}
			}
		} else {
			config, err = oci.LoadOCIConfig(cfg.OCIConfigFile)
		}

		if err == nil {
			p, err = oci.NewOCIProvider(*config, domainFilter, zoneIDFilter, cfg.DryRun)
		}
	case "rfc2136":
		p, err = rfc2136.NewRfc2136Provider(cfg.RFC2136Host, cfg.RFC2136Port, cfg.RFC2136Zone, cfg.RFC2136Insecure, cfg.RFC2136TSIGKeyName, cfg.RFC2136TSIGSecret, cfg.RFC2136TSIGSecretAlg, cfg.RFC2136TAXFR, domainFilter, cfg.DryRun, cfg.RFC2136MinTTL, cfg.RFC2136GSSTSIG, cfg.RFC2136KerberosUsername, cfg.RFC2136KerberosPassword, cfg.RFC2136KerberosRealm, cfg.RFC2136BatchChangeSize, nil)
	case "ns1":
		p, err = ns1.NewNS1Provider(
			ns1.NS1Config{
				DomainFilter:  domainFilter,
				ZoneIDFilter:  zoneIDFilter,
				NS1Endpoint:   cfg.NS1Endpoint,
				NS1IgnoreSSL:  cfg.NS1IgnoreSSL,
				DryRun:        cfg.DryRun,
				MinTTLSeconds: cfg.NS1MinTTLSeconds,
			},
		)
	case "transip":
		p, err = transip.NewTransIPProvider(cfg.TransIPAccountName, cfg.TransIPPrivateKeyFile, domainFilter, cfg.DryRun)
	case "scaleway":
		p, err = scaleway.NewScalewayProvider(ctx, domainFilter, cfg.DryRun)
	case "godaddy":
		p, err = godaddy.NewGoDaddyProvider(ctx, domainFilter, cfg.GoDaddyTTL, cfg.GoDaddyAPIKey, cfg.GoDaddySecretKey, cfg.GoDaddyOTE, cfg.DryRun)
	case "gandi":
		p, err = gandi.NewGandiProvider(ctx, domainFilter, cfg.DryRun)
	case "pihole":
		p, err = pihole.NewPiholeProvider(
			pihole.PiholeConfig{
				Server:                cfg.PiholeServer,
				Password:              cfg.PiholePassword,
				TLSInsecureSkipVerify: cfg.PiholeTLSInsecureSkipVerify,
				DomainFilter:          domainFilter,
				DryRun:                cfg.DryRun,
			},
		)
	case "ibmcloud":
		p, err = ibmcloud.NewIBMCloudProvider(cfg.IBMCloudConfigFile, domainFilter, zoneIDFilter, endpointsSource, cfg.IBMCloudProxied, cfg.DryRun)
	case "safedns":
		p, err = safedns.NewSafeDNSProvider(domainFilter, cfg.DryRun)
	case "plural":
		p, err = plural.NewPluralProvider(cfg.PluralCluster, cfg.PluralProvider)
	case "tencentcloud":
		p, err = tencentcloud.NewTencentCloudProvider(domainFilter, zoneIDFilter, cfg.TencentCloudConfigFile, cfg.TencentCloudZoneType, cfg.DryRun)
	case "webhook":
		startedChan := make(chan struct{})
		if cfg.RunAWSProviderAsWebhook {
			awsProvider, awsErr := aws.NewAWSProvider(aws.AWSConfig{
				DomainFilter:         domainFilter,
				ZoneIDFilter:         zoneIDFilter,
				ZoneTypeFilter:       zoneTypeFilter,
				ZoneTagFilter:        zoneTagFilter,
				BatchChangeSize:      cfg.AWSBatchChangeSize,
				BatchChangeInterval:  cfg.AWSBatchChangeInterval,
				EvaluateTargetHealth: cfg.AWSEvaluateTargetHealth,
				PreferCNAME:          cfg.AWSPreferCNAME,
				DryRun:               cfg.DryRun,
				ZoneCacheDuration:    cfg.AWSZoneCacheDuration,
			}, route53.New(awsSession))
			if awsErr != nil {
				log.Fatal(awsErr)
			}
			go webhook.StartHTTPApi(awsProvider, startedChan, cfg.WebhookProviderReadTimeout, cfg.WebhookProviderWriteTimeout, "127.0.0.1:8888")
			<-startedChan
		}
		p, err = webhook.NewWebhookProvider(cfg.WebhookProviderURL)
	default:
		log.Fatalf("unknown dns provider: %s", cfg.Provider)
	}
	if err != nil {
		log.Fatal(err)
	}

	var r registry.Registry
	switch cfg.Registry {
	case "dynamodb":
		config := awsSDK.NewConfig()
		if cfg.AWSDynamoDBRegion != "" {
			config = config.WithRegion(cfg.AWSDynamoDBRegion)
		}
		r, err = registry.NewDynamoDBRegistry(p, cfg.TXTOwnerID, dynamodb.New(awsSession, config), cfg.AWSDynamoDBTable, cfg.TXTPrefix, cfg.TXTSuffix, cfg.TXTWildcardReplacement, cfg.ManagedDNSRecordTypes, []byte(cfg.TXTEncryptAESKey), cfg.TXTCacheInterval)
	case "noop":
		r, err = registry.NewNoopRegistry(p)
	case "txt":
		r, err = registry.NewTXTRegistry(p, cfg.TXTPrefix, cfg.TXTSuffix, cfg.TXTOwnerID, cfg.TXTCacheInterval, cfg.TXTWildcardReplacement, cfg.ManagedDNSRecordTypes, cfg.TXTEncryptEnabled, []byte(cfg.TXTEncryptAESKey))
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
		Source:               endpointsSource,
		Registry:             r,
		Policy:               policy,
		Interval:             cfg.Interval,
		DomainFilter:         domainFilter,
		ManagedRecordTypes:   cfg.ManagedDNSRecordTypes,
		MinEventSyncInterval: cfg.MinEventSyncInterval,
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

func handleSigterm(cancel func()) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM)
	<-signals
	log.Info("Received SIGTERM. Terminating...")
	cancel()
}

func serveMetrics(address string) {
	http.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	http.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(address, nil))
}
