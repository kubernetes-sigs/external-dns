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
	"sigs.k8s.io/external-dns/source"
)

type Config struct {
	defaultTargets      []string
	forceDefaultTargets bool
	nat64Networks       []string
	targetNetFilter     []string
	excludeTargetNets   []string
	minTTL              time.Duration
	preferAlias         bool
	sourceWrappers      map[string]bool // map of source wrappers, e.g. "targetfilter", "nat64"
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

func WithPreferAlias(enabled bool) Option {
	return func(o *Config) {
		o.preferAlias = enabled
	}
}

// addSourceWrapper registers a source wrapper by name in the Config.
// It initializes the sourceWrappers map if it is nil.
func (o *Config) addSourceWrapper(name string) {
	if o.sourceWrappers == nil {
		o.sourceWrappers = make(map[string]bool)
	}
	o.sourceWrappers[name] = true
}

// isSourceWrapperInstrumented returns whether a source wrapper is enabled or not.
func (o *Config) isSourceWrapperInstrumented(name string) bool {
	if o.sourceWrappers == nil {
		return false
	}
	_, ok := o.sourceWrappers[name]
	return ok
}

// WrapSources combines multiple sources into a single source,
// applies optional NAT64 and target network filtering wrappers, and sets a minimum TTL.
// It registers each applied wrapper in the Config for instrumentation.
func WrapSources(
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
	combinedSource = NewPostProcessor(combinedSource, WithTTL(opts.minTTL), WithPostProcessorPreferAlias(opts.preferAlias))
	opts.addSourceWrapper("post-processor")
	return combinedSource, nil
}
