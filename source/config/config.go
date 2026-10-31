/*
Copyright 2026 The Kubernetes Authors.

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

// Package config holds the configuration passed to source constructors.
package config

import (
	"fmt"
	"sync"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"

	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
	"sigs.k8s.io/external-dns/pkg/crd"
	"sigs.k8s.io/external-dns/source/annotations"
	"sigs.k8s.io/external-dns/source/annotations/schema"
	"sigs.k8s.io/external-dns/source/template"
)

// Config holds shared configuration options for all Sources.
// This struct centralizes all source-related configuration to avoid parameter proliferation
// in individual source constructors. It follows the configuration pattern where a single
// config object is passed rather than individual parameters.
//
// Common Configuration Fields:
// - Namespaces: Target namespaces for source operations, empty for all of them
// - AnnotationFilter: Filter sources by annotation selector
// - LabelFilter: Filter sources by label selectors
// - FQDNTemplate: Template for generating fully qualified domain names
// - CombineFQDNAndAnnotation: Whether to combine FQDN template with annotations
// - IgnoreHostnameAnnotation: Whether to ignore hostname annotations
//
// The config is created from externaldns.Config via NewSourceConfig() which handles
// type conversions and validation.
type Config struct {
	Namespaces                     []string
	AnnotationFilter               labels.Selector
	Schema                         *schema.Schema
	LabelFilter                    labels.Selector
	IngressClassNames              []string
	TemplateEngine                 template.Engine
	IgnoreHostnameAnnotation       bool
	IgnoreNonHostNetworkPods       bool
	IgnoreIngressTLSSpec           bool
	IgnoreIngressRulesSpec         bool
	ListenEndpointEvents           bool
	GatewayName                    string
	GatewayNamespace               string
	GatewayLabelFilter             string
	GatewayListenerSets            bool
	Compatibility                  string
	Provider                       string
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
	KubeAPIRequestTimeout          time.Duration
	KubeAPIQPS                     int
	KubeAPIBurst                   int
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
	UnstructuredResources          []string
	PreferAlias                    bool
	PTRSupported                   bool
	CreatePTR                      bool

	Sources []string

	// clientGen holds a source.ClientGenerator.
	// It is any to keep unnecessary clientsets out of this package.
	clientGen     any
	clientGenOnce sync.Once

	// crdClients lets callers outside this package reuse the crd source's client
	// instead of building a second one.
	crdClients *crd.CRDClients
}

// CRDClients returns the reader/writer built for the crd source, or nil if the
// crd source was not built.
func (cfg *Config) CRDClients() *crd.CRDClients {
	return cfg.crdClients
}

// SetCRDClients records the crd source's clients for reuse by CRDClients.
func (cfg *Config) SetCRDClients(cc *crd.CRDClients) {
	cfg.crdClients = cc
}

// SetClientGen sets the Kubernetes client factory. gen must be a source.ClientGenerator.
func (cfg *Config) SetClientGen(gen any) {
	cfg.clientGen = gen
}

// EnsureClientGen returns the client factory, building it with init on the
// first call if none was set. It panics if the stored value is not a T.
func (cfg *Config) EnsureClientGen[T any](init func() T) T {
	cfg.clientGenOnce.Do(func() {
		if cfg.clientGen == nil {
			cfg.clientGen = init()
		}
	})
	return cfg.clientGen.(T)
}

// OverrideConfigOption configures a Config.
type OverrideConfigOption func(*Config)

// WithCRDClients overrides the crd source's clients. Intended for testing.
func WithCRDClients(cc *crd.CRDClients) OverrideConfigOption {
	return func(cfg *Config) {
		cfg.crdClients = cc
	}
}

// NewSourceConfig builds a Config from the parsed CLI configuration. It
// validates label and annotation filters and compiles the template engine so
// invalid user input fails before any source is instantiated.
func NewSourceConfig(cfg *externaldns.Config, opts ...OverrideConfigOption) (*Config, error) {
	labelSelector, err := labels.Parse(cfg.LabelFilter)
	if err != nil {
		return nil, fmt.Errorf("label filter: %w", err)
	}
	annotationSelector, err := annotations.ParseFilter(cfg.AnnotationFilter)
	if err != nil {
		return nil, fmt.Errorf("annotation filter: %w", err)
	}
	tmpls, err := template.NewEngine(cfg.FQDNTemplate, cfg.TargetTemplate, cfg.FQDNTargetTemplate, cfg.CombineFQDNAndAnnotation)
	if err != nil {
		return nil, err
	}
	c := &Config{
		Namespaces:       externaldns.NormalizeNamespaces(cfg.Namespaces),
		AnnotationFilter: annotationSelector,
		Schema: schema.New(schema.Options{
			Mode:                     schema.Mode(cfg.AnnotationValidationMode),
			IgnoreHostnameAnnotation: cfg.IgnoreHostnameAnnotation,
		}),
		LabelFilter:                    labelSelector,
		IngressClassNames:              cfg.IngressClassNames,
		IgnoreHostnameAnnotation:       cfg.IgnoreHostnameAnnotation,
		IgnoreNonHostNetworkPods:       cfg.IgnoreNonHostNetworkPods,
		IgnoreIngressTLSSpec:           cfg.IgnoreIngressTLSSpec,
		IgnoreIngressRulesSpec:         cfg.IgnoreIngressRulesSpec,
		ListenEndpointEvents:           cfg.ListenEndpointEvents,
		GatewayName:                    cfg.GatewayName,
		GatewayNamespace:               cfg.GatewayNamespace,
		GatewayLabelFilter:             cfg.GatewayLabelFilter,
		GatewayListenerSets:            cfg.GatewayListenerSets,
		Compatibility:                  cfg.Compatibility,
		PodSourceDomain:                cfg.PodSourceDomain,
		PublishInternal:                cfg.PublishInternal,
		PublishHostIP:                  cfg.PublishHostIP,
		Provider:                       cfg.Provider,
		AlwaysPublishNotReadyAddresses: cfg.AlwaysPublishNotReadyAddresses,
		ConnectorServer:                cfg.ConnectorSourceServer,
		CRDSourceAPIVersion:            cfg.CRDSourceAPIVersion,
		CRDSourceKind:                  cfg.CRDSourceKind,
		KubeConfig:                     cfg.KubeConfig,
		APIServerURL:                   cfg.APIServerURL,
		ServiceTypeFilter:              cfg.ServiceTypeFilter,
		GlooNamespaces:                 cfg.GlooNamespaces,
		KubeAPIRequestTimeout:          cfg.KubeAPIRequestTimeout,
		KubeAPIQPS:                     cfg.KubeAPIQPS,
		KubeAPIBurst:                   cfg.KubeAPIBurst,
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
		UnstructuredResources:          cfg.UnstructuredResources,
		TemplateEngine:                 tmpls,
		PreferAlias:                    cfg.PreferAlias,
		PTRSupported:                   cfg.IsPTRSupported(),
		CreatePTR:                      cfg.CreatePTR,
		Sources:                        cfg.Sources,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c, nil
}

// Namespace returns the namespace watched by sources supporting a single one,
// which ValidateConfig guarantees. Empty means all namespaces.
func (cfg *Config) Namespace() string {
	if len(cfg.Namespaces) == 0 {
		return v1.NamespaceAll
	}
	return cfg.Namespaces[0]
}
