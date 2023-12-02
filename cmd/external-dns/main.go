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

	awsSDK "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
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
	"sigs.k8s.io/external-dns/provider/aws"
	"sigs.k8s.io/external-dns/provider/awssd"
	"sigs.k8s.io/external-dns/provider/webhook"
	"sigs.k8s.io/external-dns/registry"
	"sigs.k8s.io/external-dns/source"
)

var providerMap = make(map[string]provider.Provider)
var cfg *externaldns.Config
var domainFilter endpoint.DomainFilter

var ctx context.Context
var cancel context.CancelFunc
var sources []source.Source
var sourceCfg *source.Config
var targetFilter endpoint.TargetNetFilter

var awsSession *session.Session

func init() {
	cfg = externaldns.NewConfig()
	if err := cfg.ParseFlags(os.Args[1:]); err != nil {
		log.Fatalf("flag parsing error: %v", err)
	}
	if cfg.LogFormat == "json" {
		log.SetFormatter(&log.JSONFormatter{})
	}
}

func main() {
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

	ctx, cancel = context.WithCancel(context.Background())

	go serveMetrics(cfg.MetricsAddress)
	go handleSigterm(cancel)

	// error is explicitly ignored because the filter is already validated in validation.ValidateConfig
	labelSelector, _ := labels.Parse(cfg.LabelFilter)

	// Create a source.Config from the flags passed by the user.
	sourceCfg = &source.Config{
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
		GlooNamespaces:                 cfg.GlooNamespaces,
		SkipperRouteGroupVersion:       cfg.SkipperRouteGroupVersion,
		RequestTimeout:                 cfg.RequestTimeout,
		DefaultTargets:                 cfg.DefaultTargets,
		OCPRouterName:                  cfg.OCPRouterName,
		UpdateEvents:                   cfg.UpdateEvents,
		ResolveLoadBalancerHostname:    cfg.ResolveServiceLoadBalancerHostname,
	}

	// Lookup all the selected sources by names and pass them the desired configuration.
	sources, err = source.ByNames(ctx, &source.SingletonClientGenerator{
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
	targetFilter = endpoint.NewTargetNetFilterWithExclusions(cfg.TargetNetFilter, cfg.ExcludeTargetNets)

	// Combine multiple sources into a single, deduplicated source.
	endpointsSource := source.NewDedupSource(source.NewMultiSource(sources, sourceCfg.DefaultTargets))
	endpointsSource = source.NewTargetFilterSource(endpointsSource, targetFilter)

	// RegexDomainFilter overrides DomainFilter

	if cfg.RegexDomainFilter.String() != "" {
		domainFilter = endpoint.NewRegexDomainFilter(cfg.RegexDomainFilter, cfg.RegexDomainExclusion)
	} else {
		domainFilter = endpoint.NewDomainFilterWithExclusions(cfg.DomainFilter, cfg.ExcludeDomains)
	}

	p := providerMap[cfg.Provider]
	if p == nil {
		log.Fatalf("unknown dns provider: %s", cfg.Provider)
	}
	log.Infof("using provider %s", cfg.Provider)

	if cfg.WebhookServer {
		webhook.StartHTTPApi(p, nil, cfg.WebhookProviderReadTimeout, cfg.WebhookProviderWriteTimeout, "127.0.0.1:8888")
		os.Exit(0)
	}

	var r registry.Registry
	switch cfg.Registry {
	case "dynamodb":
		if awsSession == nil {
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
		config := awsSDK.NewConfig()
		if cfg.AWSDynamoDBRegion != "" {
			config = config.WithRegion(cfg.AWSDynamoDBRegion)
		}
		r, err = registry.NewDynamoDBRegistry(p, cfg.TXTOwnerID, dynamodb.New(awsSession, config), cfg.AWSDynamoDBTable, cfg.TXTPrefix, cfg.TXTSuffix, cfg.TXTWildcardReplacement, cfg.ManagedDNSRecordTypes, cfg.ExcludeDNSRecordTypes, []byte(cfg.TXTEncryptAESKey), cfg.TXTCacheInterval)
	case "noop":
		r, err = registry.NewNoopRegistry(p)
	case "txt":
		r, err = registry.NewTXTRegistry(p, cfg.TXTPrefix, cfg.TXTSuffix, cfg.TXTOwnerID, cfg.TXTCacheInterval, cfg.TXTWildcardReplacement, cfg.ManagedDNSRecordTypes, cfg.ExcludeDNSRecordTypes, cfg.TXTEncryptEnabled, []byte(cfg.TXTEncryptAESKey))
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
		ExcludeRecordTypes:   cfg.ExcludeDNSRecordTypes,
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
