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

package wrappers

import (
	"fmt"
	"time"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/sets"
	"sigs.k8s.io/external-dns/source"
)

type Config struct {
	defaultTargets      []string
	forceDefaultTargets bool
	provider            string
	nat64Networks       []string
	targetNetFilter     []string
	excludeTargetNets   []string
	minTTL              time.Duration
	preferAlias         bool
	ptrSupported        bool             // PTR is in --managed-record-types
	createPTR           bool             // --create-ptr default for all A/AAAA records
	sourceWrappers      sets.Set[string] // set of source wrappers, e.g. "targetfilter", "nat64"

	// acmeDelegation* correspond to the --acme-cname-delegation-* flags.
	acmeDelegationTargetTemplate string
	acmeDelegationDomainFilter   []string
	acmeDelegationTTL            time.Duration
}

func NewConfig(opts ...Option) *Config {
	o := &Config{}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

type Option func(config *Config)

func WithDefaultTargets(input []string) Option {
	return func(o *Config) {
		o.defaultTargets = input
	}
}

func WithForceDefaultTargets(input bool) Option {
	return func(o *Config) {
		o.forceDefaultTargets = input
	}
}

func WithNAT64Networks(input []string) Option {
	return func(o *Config) {
		o.nat64Networks = input
	}
}

func WithTargetNetFilter(input []string) Option {
	return func(o *Config) {
		o.targetNetFilter = input
	}
}

func WithExcludeTargetNets(input []string) Option {
	return func(o *Config) {
		o.excludeTargetNets = input
	}
}

func WithMinTTL(ttl time.Duration) Option {
	return func(o *Config) {
		o.minTTL = ttl
	}
}

// WithProvider sets the DNS provider name, used to filter provider-specific
// endpoint properties to only those belonging to the configured provider.
func WithProvider(input string) Option {
	return func(o *Config) {
		o.provider = input
	}
}

func WithPreferAlias(enabled bool) Option {
	return func(o *Config) {
		o.preferAlias = enabled
	}
}

// WithPTRSupported indicates whether PTR is included in --managed-record-types.
// When false the PTR source wrapper is not installed at all, so no reverse
// records are generated regardless of the --create-ptr flag.
func WithPTRSupported(supported bool) Option {
	return func(o *Config) {
		o.ptrSupported = supported
	}
}

// WithCreatePTR sets the global default for automatic PTR record creation
// (the --create-ptr flag). When true, every A/AAAA endpoint gets a PTR record
// unless the resource opts out via annotation. When false, only resources that
// explicitly request PTR via annotation produce reverse records.
func WithCreatePTR(enabled bool) Option {
	return func(o *Config) {
		o.createPTR = enabled
	}
}

// WithACMEDelegationTargetTemplate sets the target template for ACME DNS-01
// delegation CNAME generation (the --acme-cname-delegation-target-template flag).
// A non-empty template enables the ACME delegation source wrapper.
func WithACMEDelegationTargetTemplate(input string) Option {
	return func(o *Config) {
		o.acmeDelegationTargetTemplate = input
	}
}

// WithACMEDelegationDomainFilter limits ACME delegation CNAME generation to
// hostnames matching the given domain suffixes (the --acme-cname-delegation-domain-filter
// flag). An empty filter matches all hostnames.
func WithACMEDelegationDomainFilter(input []string) Option {
	return func(o *Config) {
		o.acmeDelegationDomainFilter = input
	}
}

// WithACMEDelegationTTL sets the TTL for generated ACME delegation CNAME records
// (the --acme-cname-delegation-ttl flag). Zero leaves the TTL unconfigured.
func WithACMEDelegationTTL(ttl time.Duration) Option {
	return func(o *Config) {
		o.acmeDelegationTTL = ttl
	}
}

// addSourceWrapper registers a source wrapper by name in the Config.
// It initializes the sourceWrappers map if it is nil.
func (o *Config) addSourceWrapper(name string) {
	if o.sourceWrappers == nil {
		o.sourceWrappers = sets.New[string]()
	}
	o.sourceWrappers.Insert(name)
}

// isSourceWrapperInstrumented returns whether a source wrapper is enabled or not.
func (o *Config) isSourceWrapperInstrumented(name string) bool {
	if len(o.sourceWrappers) == 0 {
		return false
	}
	return o.sourceWrappers.Has(name)
}

// wrapSources combines multiple sources into a single source,
// applies optional NAT64 and target network filtering wrappers, and sets a minimum TTL.
// It registers each applied wrapper in the Config for instrumentation.
func wrapSources(
	sources []source.Source,
	opts *Config,
) (source.Source, error) {
	combinedSource := NewDedupSource(NewMultiSource(sources, opts.defaultTargets, opts.forceDefaultTargets))
	opts.addSourceWrapper("dedup")
	if len(opts.nat64Networks) > 0 {
		var err error
		combinedSource, err = NewNAT64Source(combinedSource, opts.nat64Networks)
		if err != nil {
			return nil, fmt.Errorf("failed to create NAT64 source wrapper: %w", err)
		}
		opts.addSourceWrapper("nat64")
	}
	targetFilter := endpoint.NewTargetNetFilterWithExclusions(opts.targetNetFilter, opts.excludeTargetNets)
	if targetFilter.IsEnabled() {
		combinedSource = NewTargetFilterSource(combinedSource, targetFilter)
		opts.addSourceWrapper("target-filter")
	}
	if opts.ptrSupported {
		combinedSource = NewPTRSource(combinedSource, opts.createPTR)
		opts.addSourceWrapper("ptr")
	}
	if opts.acmeDelegationTargetTemplate != "" {
		var err error
		combinedSource, err = NewACMEDelegationSource(combinedSource, opts.acmeDelegationTargetTemplate, opts.acmeDelegationDomainFilter, opts.acmeDelegationTTL)
		if err != nil {
			return nil, fmt.Errorf("failed to create ACME delegation source wrapper: %w", err)
		}
		opts.addSourceWrapper("acme-delegation")
	}
	combinedSource = NewPostProcessor(combinedSource, WithTTL(opts.minTTL), WithPostProcessorPreferAlias(opts.preferAlias),
		WithPostProcessorProvider(opts.provider))
	opts.addSourceWrapper("post-processor")
	return combinedSource, nil
}
