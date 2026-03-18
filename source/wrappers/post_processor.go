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
	"context"
	"maps"
	"net"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/source"
	"sigs.k8s.io/external-dns/source/annotations"
)

type postProcessor struct {
	source source.Source
	cfg    PostProcessorConfig
}

type PostProcessorConfig struct {
	ttl                         int64
	provider                    string
	preferAlias                 bool
	resolveLoadBalancerHostname bool
	lookupIP                    func(string) ([]net.IP, error)
	isConfigured                bool
}

type PostProcessorOption func(*PostProcessorConfig)

func WithTTL(ttl time.Duration) PostProcessorOption {
	return func(cfg *PostProcessorConfig) {
		if int64(ttl.Seconds()) > 0 {
			cfg.isConfigured = true
			cfg.ttl = int64(ttl.Seconds())
		}
	}
}

// WithPostProcessorProvider sets the provider used to retain provider-specific
// properties on endpoints. Empty or whitespace-only values are ignored.
func WithPostProcessorProvider(input string) PostProcessorOption {
	return func(cfg *PostProcessorConfig) {
		if p := strings.TrimSpace(input); p != "" {
			cfg.isConfigured = true
			cfg.provider = p
		}
	}
}

// WithPostProcessorPreferAlias enables setting alias=true on CNAME endpoints.
// This signals to providers that support ALIAS records (like PowerDNS, AWS)
// to create ALIAS records instead of CNAMEs.
func WithPostProcessorPreferAlias(enabled bool) PostProcessorOption {
	return func(cfg *PostProcessorConfig) {
		cfg.preferAlias = enabled
		if enabled {
			cfg.isConfigured = true
		}
	}
}

// WithPostProcessorResolveLoadBalancerHostname enables resolving CNAME targets that are
// hostnames to their A/AAAA IP addresses via DNS, replacing the CNAME endpoint.
func WithPostProcessorResolveLoadBalancerHostname(enabled bool) PostProcessorOption {
	return func(cfg *PostProcessorConfig) {
		cfg.resolveLoadBalancerHostname = enabled
		if enabled {
			cfg.isConfigured = true
		}
	}
}

func NewPostProcessor(source source.Source, opts ...PostProcessorOption) source.Source {
	cfg := PostProcessorConfig{}
	for _, opt := range opts {
		opt(&cfg)
	}
	if cfg.lookupIP == nil {
		cfg.lookupIP = net.LookupIP
	}
	return &postProcessor{source: source, cfg: cfg}
}

func (pp *postProcessor) Endpoints(ctx context.Context) ([]*endpoint.Endpoint, error) {
	endpoints, err := pp.source.Endpoints(ctx)
	if err != nil {
		return nil, err
	}

	var result []*endpoint.Endpoint
	for _, ep := range endpoints {
		if ep == nil {
			result = append(result, nil)
			continue
		}
		ep.WithMinTTL(pp.cfg.ttl)
		ep.RetainProviderProperties(pp.cfg.provider)
		// Set alias annotation for CNAME records when preferAlias is enabled
		// Only set if not already explicitly configured at the source level
		if pp.cfg.preferAlias && ep.RecordType == endpoint.RecordTypeCNAME {
			if _, exists := ep.GetProviderSpecificProperty(annotations.AliasKey); !exists {
				ep.WithProviderSpecific("alias", "true")
			}
		}
		// Per-endpoint annotation overrides the global flag.
		shouldResolve := pp.cfg.resolveLoadBalancerHostname
		if v, ok := ep.GetProviderSpecificProperty("resolve-target"); ok {
			shouldResolve = v == "true"
			ep.DeleteProviderSpecificProperty("resolve-target")
		}
		if shouldResolve && ep.RecordType == endpoint.RecordTypeCNAME {
			var ipTargets endpoint.Targets
			for _, target := range ep.Targets {
				ips, err := pp.cfg.lookupIP(target)
				if err != nil {
					log.Errorf("Unable to resolve %q: %v", target, err)
					continue
				}
				for _, ip := range ips {
					ipTargets = append(ipTargets, ip.String())
				}
			}
			if len(ipTargets) == 0 {
				// All resolutions failed; skip this endpoint entirely.
				continue
			}
			resolved := endpoint.EndpointsForHostname(ep.DNSName, ipTargets, ep.RecordTTL, ep.ProviderSpecific, ep.SetIdentifier, "")
			for _, r := range resolved {
				maps.Copy(r.Labels, ep.Labels)
			}
			result = append(result, resolved...)
			continue
		}
		result = append(result, ep)
	}

	return result, nil
}

func (pp *postProcessor) AddEventHandler(ctx context.Context, handler func()) {
	log.Debug("postProcessor: adding event handler")
	pp.source.AddEventHandler(ctx, handler)
}
