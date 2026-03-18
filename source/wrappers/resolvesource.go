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

	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/source"
)

// resolveSource is a Source wrapper that resolves CNAME targets (load balancer hostnames)
// to their underlying A/AAAA IP addresses via DNS lookup.
//
// The global flag enables resolution for all CNAME endpoints. Individual endpoints can
// override this via the "resolve-target" provider-specific property:
//
//	"true"  – resolve even when the global flag is off
//	"false" – keep as CNAME even when the global flag is on
//
// The property is always consumed (deleted) by this wrapper so it does not leak
// to downstream components.
type resolveSource struct {
	source     source.Source
	globalFlag bool
	lookupIP   func(string) ([]net.IP, error)
}

// resolveSourceOption is a functional option for resolveSource.
type resolveSourceOption func(*resolveSource)

// NewResolveSource creates a new resolveSource wrapping src.
func NewResolveSource(src source.Source, globalFlag bool, opts ...resolveSourceOption) source.Source {
	rs := &resolveSource{
		source:     src,
		globalFlag: globalFlag,
		lookupIP:   net.LookupIP,
	}
	for _, opt := range opts {
		opt(rs)
	}
	return rs
}

func (rs *resolveSource) Endpoints(ctx context.Context) ([]*endpoint.Endpoint, error) {
	endpoints, err := rs.source.Endpoints(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]*endpoint.Endpoint, 0, len(endpoints))
	for _, ep := range endpoints {
		if ep == nil {
			result = append(result, nil)
			continue
		}

		shouldResolve := rs.globalFlag
		if v, ok := ep.GetProviderSpecificProperty("resolve-target"); ok {
			shouldResolve = v == "true"
			ep.DeleteProviderSpecificProperty("resolve-target")
		}

		if !shouldResolve || ep.RecordType != endpoint.RecordTypeCNAME {
			result = append(result, ep)
			continue
		}

		var ipTargets endpoint.Targets
		for _, target := range ep.Targets {
			ips, err := rs.lookupIP(target)
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
	}

	return result, nil
}

func (rs *resolveSource) AddEventHandler(ctx context.Context, handler func()) {
	log.Debug("resolveSource: adding event handler")
	rs.source.AddEventHandler(ctx, handler)
}
