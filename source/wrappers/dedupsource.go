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

package wrappers

import (
	"context"
	"maps"
	"net"
	"strings"

	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/source"
)

// dedupSource is a Source that removes duplicate endpoints from its wrapped source.
// It resolves CNAME targets (load balancer hostnames) to their underlying A/AAAA IP addresses
// via DNS lookup before deduplication, when the endpoint has a "resolve-target" provider-specific
// property set to "true".
type dedupSource struct {
	source   source.Source
	lookupIP func(string) ([]net.IP, error)
}

// DedupSourceOption is a functional option for dedupSource.
type DedupSourceOption func(*dedupSource)

// NewDedupSource creates a new dedupSource wrapping the provided Source.
func NewDedupSource(source source.Source, opts ...DedupSourceOption) source.Source {
	ds := &dedupSource{
		source:   source,
		lookupIP: net.LookupIP,
	}
	for _, opt := range opts {
		opt(ds)
	}
	return ds
}

// Endpoints collects endpoints from its wrapped source, optionally resolves CNAME targets,
// and returns them without duplicates.
func (ds *dedupSource) Endpoints(ctx context.Context) ([]*endpoint.Endpoint, error) {
	log.Debug("dedupSource: collecting endpoints and removing duplicates")

	raw, err := ds.source.Endpoints(ctx)
	if err != nil {
		return nil, err
	}

	// Resolve CNAME targets to A/AAAA records before deduplication.
	resolved := make([]*endpoint.Endpoint, 0, len(raw))
	for _, ep := range raw {
		if ep == nil {
			continue
		}

		shouldResolve := false
		if v, ok := ep.GetProviderSpecificProperty("resolve-target"); ok {
			shouldResolve = v == "true"
			ep.DeleteProviderSpecificProperty("resolve-target")
		}

		if shouldResolve && ep.RecordType == endpoint.RecordTypeCNAME {
			var ipTargets endpoint.Targets
			for _, target := range ep.Targets {
				ips, err := ds.lookupIP(target)
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
			eps := endpoint.EndpointsForHostname(ep.DNSName, ipTargets, ep.RecordTTL, ep.ProviderSpecific, ep.SetIdentifier, "")
			for _, r := range eps {
				maps.Copy(r.Labels, ep.Labels)
			}
			resolved = append(resolved, eps...)
			continue
		}

		resolved = append(resolved, ep)
	}

	// Deduplicate.
	result := make([]*endpoint.Endpoint, 0, len(resolved))
	collected := make(map[string]struct{})
	for _, ep := range resolved {
		// validate endpoint before normalization
		if ok := ep.CheckEndpoint(); !ok {
			log.Warnf("Skipping endpoint [%s:%s] due to invalid configuration [%s:%s]", ep.SetIdentifier, ep.DNSName, ep.RecordType, strings.Join(ep.Targets, ","))
			continue
		}

		if len(ep.Targets) > 1 {
			ep.Targets = endpoint.NewTargets(ep.Targets...)
		}

		identifier := strings.Join([]string{ep.RecordType, ep.DNSName, ep.SetIdentifier, ep.Targets.String()}, "/")

		if _, ok := collected[identifier]; ok {
			log.Debugf("Removing duplicate endpoint %s", ep)
			continue
		}

		collected[identifier] = struct{}{}
		result = append(result, ep)
	}

	return result, nil
}

func (ms *dedupSource) AddEventHandler(ctx context.Context, handler func()) {
	log.Debug("dedupSource: adding event handler")
	ms.source.AddEventHandler(ctx, handler)
}
