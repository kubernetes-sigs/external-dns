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

package wrappers

import (
	"context"
	"net"
	"sort"

	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/source"
)

const resolveTargetPropertyName = "resolve-target"

// resolveTarget is a Source wrapper that resolves CNAME targets (load balancer hostnames)
// to their underlying A/AAAA IP addresses via DNS lookup.
//
// Resolution is controlled per-endpoint via the "resolve-target" provider-specific property:
//
//	"true"  – resolve this endpoint's hostname targets to A/AAAA records
//	"false" – keep this endpoint's hostname targets as CNAME records
//
// The property is always consumed (deleted) by this wrapper so it does not leak
// to downstream components.
type resolveTarget struct {
	source   source.Source
	lookupIP func(string) ([]net.IP, error)
}

// resolveTargetOption is a functional option for resolveTarget.
type resolveTargetOption func(*resolveTarget)

// WithResolveTargetLookupIP is a helper to create a resolveTargetOption that sets a custom lookupIP function for testing.
// If fn is nil, the default net.LookupIP is preserved.
func WithResolveTargetLookupIP(fn func(string) ([]net.IP, error)) resolveTargetOption {
	return func(rs *resolveTarget) {
		if fn != nil {
			rs.lookupIP = fn
		}
	}
}

// NewResolveTarget creates a new resolveTarget wrapping src.
func NewResolveTarget(src source.Source, opts ...resolveTargetOption) source.Source {
	rs := &resolveTarget{
		source:   src,
		lookupIP: net.LookupIP,
	}
	for _, opt := range opts {
		opt(rs)
	}
	return rs
}

func (rs *resolveTarget) Endpoints(ctx context.Context) ([]*endpoint.Endpoint, error) {
	endpoints, err := rs.source.Endpoints(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]*endpoint.Endpoint, 0, len(endpoints))
	for _, ep := range endpoints {
		if ep == nil {
			continue
		}

		// Always consume the "resolve-target" property so it never leaks downstream,
		// regardless of record type or value otherwise it won't converge with a non-empty UpdateNew in the plan.
		resolve, _ := ep.GetProviderSpecificProperty(resolveTargetPropertyName)
		ep.DeleteProviderSpecificProperty(resolveTargetPropertyName)

		// Only CNAME endpoints opted in via "true" are resolved; everything else passes through.
		if ep.RecordType != endpoint.RecordTypeCNAME || resolve != "true" {
			result = append(result, ep)
			continue
		}

		var ipTargets endpoint.Targets
		for _, target := range ep.Targets {
			ips, err := rs.lookupIP(target)
			if err != nil {
				log.Debugf("Unable to resolve %q, skipping target: %v", target, err)
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

		// Sort targets before grouping for deterministic output.
		sort.Strings(ipTargets)

		// Group targets by record type (A vs AAAA)
		byType := map[string]endpoint.Targets{}
		for _, t := range ipTargets {
			rt := endpoint.SuitableType(t)
			byType[rt] = append(byType[rt], t)
		}

		// Emit one endpoint per record type with the same DNSName and provider-specific properties.
		for _, rt := range []string{endpoint.RecordTypeA, endpoint.RecordTypeAAAA} {
			if targets := byType[rt]; len(targets) > 0 {
				result = append(result, ep.WithTargets(targets))
			}
		}

		log.Debugf("resolveTarget: resolved endpoint %q into %d IP target(s)", ep.DNSName, len(ipTargets))
	}

	return result, nil
}

func (rs *resolveTarget) AddEventHandler(ctx context.Context, handler func()) {
	log.Debug("resolveTarget: adding event handler")
	rs.source.AddEventHandler(ctx, handler)
}
