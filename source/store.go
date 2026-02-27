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

package source

import (
	"context"
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	openshift "github.com/openshift/client-go/route/clientset/versioned"
	log "github.com/sirupsen/logrus"
	istioclient "istio.io/client-go/pkg/clientset/versioned"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	gateway "sigs.k8s.io/gateway-api/pkg/client/clientset/versioned"

	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
	kubeclient "sigs.k8s.io/external-dns/pkg/client"
	"sigs.k8s.io/external-dns/source/types"
)

// ErrSourceNotFound is returned when a requested source doesn't exist.
var ErrSourceNotFound = errors.New("source not found")

// Config holds shared configuration options for all Sources.
// This struct centralizes all source-related configuration to avoid parameter proliferation
// in individual source constructors. It follows the configuration pattern where a single
// config object is passed rather than individual parameters.
//
// Common Configuration Fields:
// - Namespace: Target namespace for source operations
// - AnnotationFilter: Filter sources by annotation patterns
// - LabelFilter: Filter sources by label selectors
// - FQDNTemplate: Template for generating fully qualified domain names
// - CombineFQDNAndAnnotation: Whether to combine FQDN template with annotations
// - IgnoreHostnameAnnotation: Whether to ignore hostname annotations
//
// The config is created from externaldns.Config via NewSourceConfig() which handles
// type conversions and validation.
type Config struct {
	Namespace                      string
	AnnotationFilter               string
	LabelFilter                    labels.Selector
	IngressClassNames              []string
	FQDNTemplate                   string
	CombineFQDNAndAnnotation       bool
	IgnoreHostnameAnnotation       bool
	IgnoreNonHostNetworkPods       bool
	IgnoreIngressTLSSpec           bool
	IgnoreIngressRulesSpec         bool
	ListenEndpointEvents           bool
	GatewayName                    string
	GatewayNamespace               string
	GatewayLabelFilter             string
	Compatibility                  string
	PodSourceDomain                string
	PublishInternal                bool
	PublishHostIP                  bool
	AlwaysPublishNotReadyAddresses bool
	ConnectorServer                string
	CRDSourceAPIVersion            string
	CRDSourceKind                  string
	KubeConfig                     string
	APIServerURL                   string
	ServiceTypeFilter              []string
	GlooNamespaces                 []string
	SkipperRouteGroupVersion       string
	RequestTimeout                 time.Duration
	DefaultTargets                 []string
	ForceDefaultTargets            bool
	OCPRouterName                  string
	UpdateEvents                   bool
	ResolveLoadBalancerHostname    bool
	TraefikEnableLegacy            bool
	TraefikDisableNew              bool
	ExcludeUnschedulable           bool
	ExposeInternalIPv6             bool
	ExcludeTargetNets              []string
	TargetNetFilter                []string
	NAT64Networks                  []string
	MinTTL                         time.Duration
	PreferAlias                    bool

	sources []string

	// clientGen is lazily initialized on first access for efficiency
	clientGen     *SingletonClientGenerator
	clientGenOnce sync.Once
}

func NewSourceConfig(cfg *externaldns.Config) *Config {
	// error is explicitly ignored because the filter is already validated in validation.ValidateConfig
	labelSelector, _ := labels.Parse(cfg.LabelFilter)
	return &Config{
		Namespace:                      cfg.Namespace,
		AnnotationFilter:               cfg.AnnotationFilter,
		LabelFilter:                    labelSelector,
		IngressClassNames:              cfg.IngressClassNames,
		FQDNTemplate:                   cfg.FQDNTemplate,
		CombineFQDNAndAnnotation:       cfg.CombineFQDNAndAnnotation,
		IgnoreHostnameAnnotation:       cfg.IgnoreHostnameAnnotation,
		IgnoreNonHostNetworkPods:       cfg.IgnoreNonHostNetworkPods,
		IgnoreIngressTLSSpec:           cfg.IgnoreIngressTLSSpec,
		IgnoreIngressRulesSpec:         cfg.IgnoreIngressRulesSpec,
		ListenEndpointEvents:           cfg.ListenEndpointEvents,
		GatewayName:                    cfg.GatewayName,
		GatewayNamespace:               cfg.GatewayNamespace,
		GatewayLabelFilter:             cfg.GatewayLabelFilter,
		Compatibility:                  cfg.Compatibility,
		PodSourceDomain:                cfg.PodSourceDomain,
		PublishInternal:                cfg.PublishInternal,
		PublishHostIP:                  cfg.PublishHostIP,
		AlwaysPublishNotReadyAddresses: cfg.AlwaysPublishNotReadyAddresses,
		ConnectorServer:                cfg.ConnectorSourceServer,
		CRDSourceAPIVersion:            cfg.CRDSourceAPIVersion,
		CRDSourceKind:                  cfg.CRDSourceKind,
		KubeConfig:                     cfg.KubeConfig,
		APIServerURL:                   cfg.APIServerURL,
		ServiceTypeFilter:              cfg.ServiceTypeFilter,
		GlooNamespaces:                 cfg.GlooNamespaces,
		SkipperRouteGroupVersion:       cfg.SkipperRouteGroupVersion,
		RequestTimeout:                 cfg.RequestTimeout,
		DefaultTargets:                 cfg.DefaultTargets,
		ForceDefaultTargets:            cfg.ForceDefaultTargets,
		OCPRouterName:                  cfg.OCPRouterName,
		UpdateEvents:                   cfg.UpdateEvents,
		ResolveLoadBalancerHostname:    cfg.ResolveServiceLoadBalancerHostname,
		TraefikEnableLegacy:            cfg.TraefikEnableLegacy,
		TraefikDisableNew:              cfg.TraefikDisableNew,
		ExcludeUnschedulable:           cfg.ExcludeUnschedulable,
		ExposeInternalIPv6:             cfg.ExposeInternalIPV6,
		ExcludeTargetNets:              cfg.ExcludeTargetNets,
		TargetNetFilter:                cfg.TargetNetFilter,
		NAT64Networks:                  cfg.NAT64Networks,
		MinTTL:                         cfg.MinTTL,
		PreferAlias:                    cfg.PreferAlias,
		sources:                        cfg.Sources,
	}
}

// ClientGenerator returns a SingletonClientGenerator from this Config's connection settings.
// The generator is created once and cached for subsequent calls.
// This ensures consistent Kubernetes client creation across all sources using this configuration.
//
// The timeout behavior is special-cased: when UpdateEvents is true, the timeout is set to 0
// (no timeout) to allow long-running watch operations for event-driven source updates.
func (cfg *Config) ClientGenerator() *SingletonClientGenerator {
	cfg.clientGenOnce.Do(func() {
		cfg.clientGen = &SingletonClientGenerator{
			KubeConfig:   cfg.KubeConfig,
			APIServerURL: cfg.APIServerURL,
			RequestTimeout: func() time.Duration {
				if cfg.UpdateEvents {
					return 0
				}
				return cfg.RequestTimeout
			}(),
		}
	})
	return cfg.clientGen
}

// ClientGenerator provides clients for various Kubernetes APIs and external services.
// This interface abstracts client creation and enables dependency injection for testing.
// It uses the singleton pattern to ensure only one instance of each client is created
// and reused across multiple source instances.
//
// Supported Client Types:
// - KubeClient: Standard Kubernetes API client
// - GatewayClient: Gateway API client for Gateway resources
// - IstioClient: Istio service mesh client
// - DynamicKubernetesClient: Dynamic client for custom resources
// - OpenShiftClient: OpenShift-specific client for Route resources
// - RESTConfig: Instrumented REST config for creating custom clients
//
// The singleton behavior is implemented in SingletonClientGenerator which uses
// sync.Once to guarantee single initialization of each client type.
type ClientGenerator interface {
	KubeClient() (kubernetes.Interface, error)
	GatewayClient() (gateway.Interface, error)
	IstioClient() (istioclient.Interface, error)
	DynamicKubernetesClient() (dynamic.Interface, error)
	OpenShiftClient() (openshift.Interface, error)
	RESTConfig() (*rest.Config, error)
}

// SingletonClientGenerator stores provider clients and guarantees that only one instance of each client
// will be generated throughout the application lifecycle.
//
// Thread Safety: Uses sync.Once for each client type to ensure thread-safe initialization.
// This is important because external-dns may create multiple sources concurrently.
//
// Memory Efficiency: Prevents creating multiple instances of expensive client objects
// that maintain their own connection pools and caches.
//
// Configuration: Clients are configured using KubeConfig, APIServerURL, and RequestTimeout
// which are set during SingletonClientGenerator initialization.
//
// TODO: Fix error handling pattern in client methods. Current implementation has a bug where
// errors are only returned on the first call due to sync.Once behavior. If initialization fails
// on the first call, subsequent calls return (nil, nil) instead of (nil, originalError), which
// can lead to nil pointer dereferences. Solution: Store error in a field alongside the client,
// similar to how the client itself is stored. Example:
//
//	type SingletonClientGenerator struct {
//	    restConfig    *rest.Config
//	    restConfigErr error        // Store error persistently
//	    restConfigOnce sync.Once
//	}
//
//	func (p *SingletonClientGenerator) RESTConfig() (*rest.Config, error) {
//	    p.restConfigOnce.Do(func() {
//	        p.restConfig, p.restConfigErr = kubeclient.InstrumentedRESTConfig(...)
//	    })
//	    return p.restConfig, p.restConfigErr  // Return stored error
//	}
//
// This pattern should be applied to all client methods: KubeClient, GatewayClient,
// DynamicKubernetesClient, OpenShiftClient, and RESTConfig.
type SingletonClientGenerator struct {
	KubeConfig      string
	APIServerURL    string
	RequestTimeout  time.Duration
	restConfig      *rest.Config
	kubeClient      kubernetes.Interface
	gatewayClient   gateway.Interface
	istioClient     *istioclient.Clientset
	dynKubeClient   dynamic.Interface
	openshiftClient openshift.Interface
	restConfigOnce  sync.Once
	kubeOnce        sync.Once
	gatewayOnce     sync.Once
	istioOnce       sync.Once
	dynCliOnce      sync.Once
	openshiftOnce   sync.Once
}

// KubeClient generates a kube client if it was not created before
func (p *SingletonClientGenerator) KubeClient() (kubernetes.Interface, error) {
	var err error
	p.kubeOnce.Do(func() {
		p.kubeClient, err = kubeclient.NewKubeClient(p.KubeConfig, p.APIServerURL, p.RequestTimeout)
	})
	return p.kubeClient, err
}

// RESTConfig generates an instrumented REST config if it was not created before.
// The config includes request timeout handling and metrics instrumentation.
// This is useful for sources that need to create custom clients (e.g., controller-runtime clients).
func (p *SingletonClientGenerator) RESTConfig() (*rest.Config, error) {
	var err error
	p.restConfigOnce.Do(func() {
		p.restConfig, err = kubeclient.InstrumentedRESTConfig(p.KubeConfig, p.APIServerURL, p.RequestTimeout)
	})
	return p.restConfig, err
}

// GatewayClient generates a gateway client if it was not created before
func (p *SingletonClientGenerator) GatewayClient() (gateway.Interface, error) {
	var err error
	p.gatewayOnce.Do(func() {
		var config *rest.Config
		config, err = p.RESTConfig()
		if err != nil {
			return
		}
		p.gatewayClient, err = gateway.NewForConfig(config)
		if err != nil {
			return
		}
		log.Infof("Created GatewayAPI client %s", config.Host)
	})
	return p.gatewayClient, err
}

// IstioClient generates an istio go client if it was not created before
func (p *SingletonClientGenerator) IstioClient() (istioclient.Interface, error) {
	var err error
	p.istioOnce.Do(func() {
		p.istioClient, err = NewIstioClient(p.KubeConfig, p.APIServerURL)
	})
	return p.istioClient, err
}

// DynamicKubernetesClient generates a dynamic client if it was not created before
func (p *SingletonClientGenerator) DynamicKubernetesClient() (dynamic.Interface, error) {
	var err error
	p.dynCliOnce.Do(func() {
		var config *rest.Config
		config, err = p.RESTConfig()
		if err != nil {
			return
		}
		p.dynKubeClient, err = dynamic.NewForConfig(config)
		if err != nil {
			return
		}
		log.Infof("Created Dynamic Kubernetes client %s", config.Host)
	})
	return p.dynKubeClient, err
}

// OpenShiftClient generates an openshift client if it was not created before
func (p *SingletonClientGenerator) OpenShiftClient() (openshift.Interface, error) {
	var err error
	p.openshiftOnce.Do(func() {
		var config *rest.Config
		config, err = p.RESTConfig()
		if err != nil {
			return
		}
		p.openshiftClient, err = openshift.NewForConfig(config)
		if err != nil {
			return
		}
		log.Infof("Created OpenShift client %s", config.Host)
	})
	return p.openshiftClient, err
}

// ByNames returns multiple Sources given multiple names.
func ByNames(ctx context.Context, cfg *Config, p ClientGenerator) ([]Source, error) {
	sources := make([]Source, 0, len(cfg.sources))
	for _, name := range cfg.sources {
		source, err := BuildWithConfig(ctx, name, p, cfg)
		if err != nil {
			return nil, err
		}
		sources = append(sources, source)
	}

	return sources, nil
}

// BuildWithConfig creates a Source implementation using the factory pattern.
// This function serves as the central registry for all available source types.
//
// Source Selection: Uses a string identifier to determine which source type to create.
// This allows for runtime configuration and easy extension with new source types.
//
// Error Handling: Returns ErrSourceNotFound for unsupported source types,
// allowing callers to handle unknown sources gracefully.
//
// Supported Source Types:
// - "node": Kubernetes nodes
// - "service": Kubernetes services
// - "ingress": Kubernetes ingresses
// - "pod": Kubernetes pods
// - "gateway-*": Gateway API resources (httproute, grpcroute, tlsroute, tcproute, udproute)
// - "istio-*": Istio resources (gateway, virtualservice)
// - "ambassador-host": Ambassador Host resources
// - "contour-httpproxy": Contour HTTPProxy resources
// - "gloo-proxy": Gloo proxy resources
// - "traefik-proxy": Traefik proxy resources
// - "openshift-route": OpenShift Route resources
// - "crd": Custom Resource Definitions
// - "skipper-routegroup": Skipper RouteGroup resources
// - "kong-tcpingress": Kong TCP Ingress resources
// - "f5-*": F5 resources (virtualserver, transportserver)
// - "fake": Fake source for testing
// - "connector": Connector source for external systems
//
// Design Note: Gateway API sources use a different pattern (direct constructor calls)
// because they have simpler initialization requirements.
func BuildWithConfig(ctx context.Context, source string, p ClientGenerator, cfg *Config) (Source, error) {
	switch source {
	case types.Node:
		return buildNodeSource(ctx, p, cfg)
	case types.Service:
		return buildServiceSource(ctx, p, cfg)
	case types.Ingress:
		return buildIngressSource(ctx, p, cfg)
	case types.Pod:
		return buildPodSource(ctx, p, cfg)
	case types.GatewayHttpRoute:
		return NewGatewayHTTPRouteSource(ctx, p, cfg)
	case types.GatewayGrpcRoute:
		return NewGatewayGRPCRouteSource(ctx, p, cfg)
	case types.GatewayTlsRoute:
		return NewGatewayTLSRouteSource(ctx, p, cfg)
	case types.GatewayTcpRoute:
		return NewGatewayTCPRouteSource(ctx, p, cfg)
	case types.GatewayUdpRoute:
		return NewGatewayUDPRouteSource(ctx, p, cfg)
	case types.IstioGateway:
		return buildIstioGatewaySource(ctx, p, cfg)
	case types.IstioVirtualService:
		return buildIstioVirtualServiceSource(ctx, p, cfg)
	case types.AmbassadorHost:
		return buildAmbassadorHostSource(ctx, p, cfg)
	case types.ContourHTTPProxy:
		return buildContourHTTPProxySource(ctx, p, cfg)
	case types.GlooProxy:
		return buildGlooProxySource(ctx, p, cfg)
	case types.TraefikProxy:
		return buildTraefikProxySource(ctx, p, cfg)
	case types.OpenShiftRoute:
		return buildOpenShiftRouteSource(ctx, p, cfg)
	case types.Fake:
		return NewFakeSource(cfg.FQDNTemplate)
	case types.Connector:
		return NewConnectorSource(cfg.ConnectorServer)
	case types.CRD:
		return buildCRDSource(ctx, p, cfg)
	case types.SkipperRouteGroup:
		return buildSkipperRouteGroupSource(ctx, cfg)
	case types.KongTCPIngress:
		return buildKongTCPIngressSource(ctx, p, cfg)
	case types.F5VirtualServer:
		return buildF5VirtualServerSource(ctx, p, cfg)
	case types.F5TransportServer:
		return buildF5TransportServerSource(ctx, p, cfg)
	}
	return nil, ErrSourceNotFound
}

// Source Builder Functions
//
// The following functions follow a standardized pattern for creating source instances.
// This standardization improves code consistency, maintainability, and readability.
//
// Standardized Function Signature Pattern:
//
//	func buildXXXSource(ctx context.Context, p ClientGenerator, cfg *Config) (Source, error)
//
// Standardized Constructor Parameter Pattern (where applicable):
//  1. ctx (context.Context) - Always first when supported by the source constructor
//  2. client(s) (kubernetes.Interface, dynamic.Interface, etc.) - Kubernetes clients
//  3. namespace (string) - Target namespace for the source
//  4. annotationFilter (string) - Filter for annotations
//  5. labelFilter (labels.Selector) - Filter for labels (when applicable)
//  6. fqdnTemplate (string) - FQDN template for DNS record generation
//  7. combineFQDNAndAnnotation (bool) - Whether to combine FQDN template with annotations
//  8. ...other parameters - Source-specific parameters in logical order
//
// Design Principles:
// - Each source type has its own specific requirements and dependencies
// - Separating build functions allows for clearer code organization and easier maintenance
// - Individual functions enable straightforward error handling and independent testing
// - Modularity makes it easier to add new source types or modify existing ones
// - Consistent parameter ordering reduces cognitive load when working with multiple sources
//
// Note: Some sources may deviate from the standard pattern due to their unique requirements
// (e.g., RouteGroupSource doesn't use ClientGenerator, GlooSource doesn't accept context)
// buildNodeSource creates a Node source for exposing node information as DNS records.
// Follows standard pattern: ctx, client, annotationFilter, fqdnTemplate, labelFilter, ...other
func buildNodeSource(ctx context.Context, p ClientGenerator, cfg *Config) (Source, error) {
	client, err := p.KubeClient()
	if err != nil {
		return nil, err
	}
	return NewNodeSource(ctx, client, cfg.AnnotationFilter, cfg.FQDNTemplate, cfg.LabelFilter, cfg.ExposeInternalIPv6, cfg.ExcludeUnschedulable, cfg.CombineFQDNAndAnnotation)
}

// buildServiceSource creates a Service source for exposing Kubernetes services as DNS records.
// Follows standard pattern: ctx, client, namespace, annotationFilter, fqdnTemplate, ...other
func buildServiceSource(ctx context.Context, p ClientGenerator, cfg *Config) (Source, error) {
	client, err := p.KubeClient()
	if err != nil {
		return nil, err
	}
	return NewServiceSource(ctx, client, cfg.Namespace,
		cfg.AnnotationFilter, cfg.FQDNTemplate, cfg.CombineFQDNAndAnnotation,
		cfg.Compatibility, cfg.PublishInternal, cfg.PublishHostIP,
		cfg.AlwaysPublishNotReadyAddresses, cfg.ServiceTypeFilter, cfg.IgnoreHostnameAnnotation,
		cfg.LabelFilter, cfg.ResolveLoadBalancerHostname, cfg.ListenEndpointEvents,
		cfg.ExposeInternalIPv6, cfg.ExcludeUnschedulable)
}

// buildIngressSource creates an Ingress source for exposing Kubernetes ingresses as DNS records.
// Follows standard pattern: ctx, client, namespace, annotationFilter, fqdnTemplate, ...other
func buildIngressSource(ctx context.Context, p ClientGenerator, cfg *Config) (Source, error) {
	client, err := p.KubeClient()
	if err != nil {
		return nil, err
	}
	return NewIngressSource(ctx, client, cfg.Namespace, cfg.AnnotationFilter, cfg.FQDNTemplate, cfg.CombineFQDNAndAnnotation, cfg.IgnoreHostnameAnnotation, cfg.IgnoreIngressTLSSpec, cfg.IgnoreIngressRulesSpec, cfg.LabelFilter, cfg.IngressClassNames)
}

// buildPodSource creates a Pod source for exposing Kubernetes pods as DNS records.
// Follows standard pattern: ctx, client, namespace, ...other (no annotation/label filters)
func buildPodSource(ctx context.Context, p ClientGenerator, cfg *Config) (Source, error) {
	client, err := p.KubeClient()
	if err != nil {
		return nil, err
	}
	return NewPodSource(ctx, client, cfg.Namespace, cfg.Compatibility, cfg.IgnoreNonHostNetworkPods, cfg.PodSourceDomain, cfg.FQDNTemplate, cfg.CombineFQDNAndAnnotation, cfg.AnnotationFilter, cfg.LabelFilter)
}

// buildIstioGatewaySource creates an Istio Gateway source for exposing Istio gateways as DNS records.
// Requires both Kubernetes and Istio clients. Follows standard parameter pattern.
func buildIstioGatewaySource(ctx context.Context, p ClientGenerator, cfg *Config) (Source, error) {
	kubernetesClient, err := p.KubeClient()
	if err != nil {
		return nil, err
	}
	istioClient, err := p.IstioClient()
	if err != nil {
		return nil, err
	}
	return NewIstioGatewaySource(ctx, kubernetesClient, istioClient, cfg.Namespace, cfg.AnnotationFilter, cfg.FQDNTemplate, cfg.CombineFQDNAndAnnotation, cfg.IgnoreHostnameAnnotation)
}

// buildIstioVirtualServiceSource creates an Istio VirtualService source for exposing virtual services as DNS records.
// Requires both Kubernetes and Istio clients. Follows standard parameter pattern.
func buildIstioVirtualServiceSource(ctx context.Context, p ClientGenerator, cfg *Config) (Source, error) {
	kubernetesClient, err := p.KubeClient()
	if err != nil {
		return nil, err
	}
	istioClient, err := p.IstioClient()
	if err != nil {
		return nil, err
	}
	return NewIstioVirtualServiceSource(ctx, kubernetesClient, istioClient, cfg.Namespace, cfg.AnnotationFilter, cfg.FQDNTemplate, cfg.CombineFQDNAndAnnotation, cfg.IgnoreHostnameAnnotation)
}

func buildAmbassadorHostSource(ctx context.Context, p ClientGenerator, cfg *Config) (Source, error) {
	kubernetesClient, err := p.KubeClient()
	if err != nil {
		return nil, err
	}
	dynamicClient, err := p.DynamicKubernetesClient()
	if err != nil {
		return nil, err
	}
	return NewAmbassadorHostSource(ctx, dynamicClient, kubernetesClient, cfg.Namespace, cfg.AnnotationFilter, cfg.LabelFilter)
}

func buildContourHTTPProxySource(ctx context.Context, p ClientGenerator, cfg *Config) (Source, error) {
	dynamicClient, err := p.DynamicKubernetesClient()
	if err != nil {
		return nil, err
	}
	return NewContourHTTPProxySource(ctx, dynamicClient, cfg.Namespace, cfg.AnnotationFilter, cfg.FQDNTemplate, cfg.CombineFQDNAndAnnotation, cfg.IgnoreHostnameAnnotation)
}

// buildGlooProxySource creates a Gloo source for exposing Gloo proxies as DNS records.
// Requires both dynamic and standard Kubernetes clients.
// Note: Does not accept context parameter in constructor (legacy design).
func buildGlooProxySource(ctx context.Context, p ClientGenerator, cfg *Config) (Source, error) {
	kubernetesClient, err := p.KubeClient()
	if err != nil {
		return nil, err
	}
	dynamicClient, err := p.DynamicKubernetesClient()
	if err != nil {
		return nil, err
	}
	return NewGlooSource(ctx, dynamicClient, kubernetesClient, cfg.GlooNamespaces)
}

func buildTraefikProxySource(ctx context.Context, p ClientGenerator, cfg *Config) (Source, error) {
	kubernetesClient, err := p.KubeClient()
	if err != nil {
		return nil, err
	}
	dynamicClient, err := p.DynamicKubernetesClient()
	if err != nil {
		return nil, err
	}
	return NewTraefikSource(ctx, dynamicClient, kubernetesClient, cfg.Namespace, cfg.AnnotationFilter, cfg.IgnoreHostnameAnnotation, cfg.TraefikEnableLegacy, cfg.TraefikDisableNew)
}

func buildOpenShiftRouteSource(ctx context.Context, p ClientGenerator, cfg *Config) (Source, error) {
	ocpClient, err := p.OpenShiftClient()
	if err != nil {
		return nil, err
	}
	return NewOcpRouteSource(ctx, ocpClient, cfg.Namespace, cfg.AnnotationFilter, cfg.FQDNTemplate, cfg.CombineFQDNAndAnnotation, cfg.IgnoreHostnameAnnotation, cfg.LabelFilter, cfg.OCPRouterName)
}

// buildCRDSource creates a CRD source for exposing custom resources as DNS records.
// Uses a specialized CRD client created via NewCRDClientForAPIVersionKind.
// Parameter order: crdClient, namespace, kind, annotationFilter, labelFilter, scheme, updateEvents
func buildCRDSource(_ context.Context, p ClientGenerator, cfg *Config) (Source, error) {
	client, err := p.KubeClient()
	if err != nil {
		return nil, err
	}
	crdClient, scheme, err := NewCRDClientForAPIVersionKind(client, cfg.KubeConfig, cfg.APIServerURL, cfg.CRDSourceAPIVersion, cfg.CRDSourceKind)
	if err != nil {
		return nil, err
	}
	return NewCRDSource(crdClient, cfg.Namespace, cfg.CRDSourceKind, cfg.AnnotationFilter, cfg.LabelFilter, scheme, cfg.UpdateEvents)
}

// buildSkipperRouteGroupSource creates a Skipper RouteGroup source for exposing route groups as DNS records.
// Special case: Does not use ClientGenerator pattern, instead manages its own authentication.
// Retrieves bearer token from REST config for API server authentication.
func buildSkipperRouteGroupSource(_ context.Context, cfg *Config) (Source, error) {
	apiServerURL := cfg.APIServerURL
	tokenPath := ""
	token := ""
	restConfig, err := kubeclient.GetRestConfig(cfg.KubeConfig, cfg.APIServerURL)
	if err == nil {
		apiServerURL = restConfig.Host
		tokenPath = restConfig.BearerTokenFile
		token = restConfig.BearerToken
	}
	return NewRouteGroupSource(cfg.RequestTimeout, token, tokenPath, apiServerURL, cfg.Namespace, cfg.AnnotationFilter, cfg.FQDNTemplate, cfg.SkipperRouteGroupVersion, cfg.CombineFQDNAndAnnotation, cfg.IgnoreHostnameAnnotation)
}

func buildKongTCPIngressSource(ctx context.Context, p ClientGenerator, cfg *Config) (Source, error) {
	kubernetesClient, err := p.KubeClient()
	if err != nil {
		return nil, err
	}
	dynamicClient, err := p.DynamicKubernetesClient()
	if err != nil {
		return nil, err
	}
	return NewKongTCPIngressSource(ctx, dynamicClient, kubernetesClient, cfg.Namespace, cfg.AnnotationFilter, cfg.IgnoreHostnameAnnotation)
}

func buildF5VirtualServerSource(ctx context.Context, p ClientGenerator, cfg *Config) (Source, error) {
	kubernetesClient, err := p.KubeClient()
	if err != nil {
		return nil, err
	}
	dynamicClient, err := p.DynamicKubernetesClient()
	if err != nil {
		return nil, err
	}
	return NewF5VirtualServerSource(ctx, dynamicClient, kubernetesClient, cfg.Namespace, cfg.AnnotationFilter)
}

func buildF5TransportServerSource(ctx context.Context, p ClientGenerator, cfg *Config) (Source, error) {
	kubernetesClient, err := p.KubeClient()
	if err != nil {
		return nil, err
	}
	dynamicClient, err := p.DynamicKubernetesClient()
	if err != nil {
		return nil, err
	}
	return NewF5TransportServerSource(ctx, dynamicClient, kubernetesClient, cfg.Namespace, cfg.AnnotationFilter)
}

// NewIstioClient returns a new Istio client object. It uses the configured
// KubeConfig attribute to connect to the cluster. If KubeConfig isn't provided
// it defaults to using the recommended default.
// NB: Istio controls the creation of the underlying Kubernetes client, so we
// have no ability to tack on transport wrappers (e.g., Prometheus request
// wrappers) to the client's config at this level. Furthermore, the Istio client
// constructor does not expose the ability to override the Kubernetes API server endpoint,
// so the apiServerURL config attribute has no effect.
func NewIstioClient(kubeConfig string, apiServerURL string) (*istioclient.Clientset, error) {
	if kubeConfig == "" {
		if _, err := os.Stat(clientcmd.RecommendedHomeFile); err == nil {
			kubeConfig = clientcmd.RecommendedHomeFile
		}
	}

	restCfg, err := clientcmd.BuildConfigFromFlags(apiServerURL, kubeConfig)
	if err != nil {
		return nil, err
	}

	ic, err := istioclient.NewForConfig(restCfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create istio client: %w", err)
	}

	return ic, nil
}
