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

	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/source"
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
	log.Debug("ptrSource: adding event handler")
	s.source.AddEventHandler(ctx, handler)
}

// supportsPTR returns true if the endpoint is eligible for PTR record generation.
// Only non-wildcard A and AAAA records can have meaningful reverse DNS mappings.
func supportsPTR(ep *endpoint.Endpoint) bool {
	return (ep.RecordType == endpoint.RecordTypeA || ep.RecordType == endpoint.RecordTypeAAAA) &&
		!strings.HasPrefix(ep.DNSName, "*.")
}

// generatePTREndpoints creates PTR endpoints for A/AAAA endpoints.
// When multiple records share an IP, their hostnames are grouped into a single
// endpoint with multiple targets, which translates to multiple PTR RRs in DNS.
func generatePTREndpoints(endpoints []*endpoint.Endpoint, defaultEnabled bool) []*endpoint.Endpoint {
	type ptrInfo struct {
		targets []string
		ttl     endpoint.TTL
	}
	ptrMap := make(map[string]*ptrInfo)
	var order []string

	for _, ep := range endpoints {
		if !supportsPTR(ep) {
			if _, ok := ep.RequestedRecordType(); ok {
				log.Warnf("PTR: ignoring record-type annotation on unsupported %s record %s (only A/AAAA supported)", ep.RecordType, ep.DNSName)
				ep.DeleteProviderSpecificProperty(endpoint.ProviderSpecificRecordType)
			}
			continue
		}

		enabled := defaultEnabled
		if val, ok := ep.RequestedRecordType(); ok {
			enabled = strings.EqualFold(val, endpoint.RecordTypePTR)
			ep.DeleteProviderSpecificProperty(endpoint.ProviderSpecificRecordType)
		}
		if !enabled {
			continue
		}

		for _, target := range ep.Targets {
			if info, ok := ptrMap[target]; ok {
				if ep.RecordTTL < info.ttl {
					log.Warnf("PTR: conflicting TTLs for %s (from %s TTL=%d vs existing TTL=%d), using minimum", target, ep.DNSName, ep.RecordTTL, info.ttl)
					info.ttl = ep.RecordTTL
				}
				info.targets = append(info.targets, ep.DNSName)
			} else {
				ptrMap[target] = &ptrInfo{targets: []string{ep.DNSName}, ttl: ep.RecordTTL}
				order = append(order, target)
			}
		}
	}

	result := make([]*endpoint.Endpoint, 0, len(ptrMap))
	for _, ip := range order {
		info := ptrMap[ip]
		ep, err := endpoint.NewPTREndpoint(ip, info.ttl, info.targets...)
		if err != nil {
			log.Warnf("PTR: %v", err)
			continue
		}
		result = append(result, ep)
	}
	return result
}
