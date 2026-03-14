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
	"strings"

	"github.com/miekg/dns"
	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/source"
	"sigs.k8s.io/external-dns/source/annotations"
)

// ptrSource wraps a Source to append PTR endpoints for every A/AAAA endpoint.
// The defaultEnabled flag corresponds to --create-ptr. Per-endpoint, the "record-type"
// provider-specific property (from the resource annotation) overrides this default.
type ptrSource struct {
	source         source.Source
	defaultEnabled bool
}

// NewPTRSource creates a source wrapper that generates PTR records.
func NewPTRSource(source source.Source, defaultEnabled bool) source.Source {
	return &ptrSource{source: source, defaultEnabled: defaultEnabled}
}

func (s *ptrSource) Endpoints(ctx context.Context) ([]*endpoint.Endpoint, error) {
	endpoints, err := s.source.Endpoints(ctx)
	if err != nil {
		return nil, err
	}

	ptrEndpoints := generatePTREndpoints(endpoints, s.defaultEnabled)
	if len(ptrEndpoints) > 0 {
		log.Debugf("PTR: generated %d PTR endpoints from A/AAAA records", len(ptrEndpoints))
		endpoints = append(endpoints, ptrEndpoints...)
	}
	return endpoints, nil
}

func (s *ptrSource) AddEventHandler(ctx context.Context, handler func()) {
	s.source.AddEventHandler(ctx, handler)
}

type ptrInfo struct {
	targets []string
	ttl     endpoint.TTL
}

// generatePTREndpoints creates PTR endpoints for A/AAAA endpoints.
// When multiple records share an IP, a single PTR with all hostnames is created.
func generatePTREndpoints(endpoints []*endpoint.Endpoint, defaultEnabled bool) []*endpoint.Endpoint {
	ptrMap := make(map[string]*ptrInfo)
	var order []string

	for _, ep := range endpoints {
		if !ep.SupportsPTR() {
			continue
		}

		enabled := defaultEnabled
		if val, ok := ep.GetProviderSpecificProperty(annotations.RecordTypeProviderSpecificProperty); ok {
			enabled = strings.EqualFold(val, "ptr")
			ep.DeleteProviderSpecificProperty(annotations.RecordTypeProviderSpecificProperty)
		}
		if !enabled {
			continue
		}

		for _, target := range ep.Targets {
			revAddr, err := dns.ReverseAddr(target)
			if err != nil {
				log.Warnf("PTR: failed to compute reverse address for %s: %v", target, err)
				continue
			}
			// Strip trailing dot from reverse address (external-dns convention)
			ptrName := strings.TrimSuffix(revAddr, ".")

			if info, ok := ptrMap[ptrName]; ok {
				info.targets = append(info.targets, ep.DNSName)
				if ep.RecordTTL != info.ttl {
					log.Warnf("PTR: conflicting TTLs for %s (from %s TTL=%d vs existing TTL=%d), using minimum", ptrName, ep.DNSName, ep.RecordTTL, info.ttl)
					if ep.RecordTTL < info.ttl {
						info.ttl = ep.RecordTTL
					}
				}
			} else {
				ptrMap[ptrName] = &ptrInfo{targets: []string{ep.DNSName}, ttl: ep.RecordTTL}
				order = append(order, ptrName)
			}
		}
	}

	result := make([]*endpoint.Endpoint, 0, len(ptrMap))
	for _, name := range order {
		info := ptrMap[name]
		result = append(result, endpoint.NewEndpointWithTTL(name, endpoint.RecordTypePTR, info.ttl, info.targets...))
	}
	return result
}
