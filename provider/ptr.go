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

package provider

import (
	"fmt"
	"net"
	"strings"

	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
)

const hexDigit = "0123456789abcdef"

// ReverseAddr returns the in-addr.arpa. or ip6.arpa. name for the given IP address,
// suitable for rDNS (PTR) record lookup.
// For example, "1.2.3.4" becomes "4.3.2.1.in-addr.arpa."
// and "2001:db8::1" becomes "1.0.0.0...8.b.d.0.1.0.0.2.ip6.arpa."
// The returned address always includes a trailing dot.
func ReverseAddr(addr string) (string, error) {
	ip := net.ParseIP(addr)
	if ip == nil {
		return "", fmt.Errorf("invalid IP address: %s", addr)
	}
	if ip.To4() != nil {
		return fmt.Sprintf("%d.%d.%d.%d.in-addr.arpa.", ip[15], ip[14], ip[13], ip[12]), nil
	}
	// Must be IPv6
	buf := make([]byte, 0, len(ip)*4+len("ip6.arpa."))
	for i := len(ip) - 1; i >= 0; i-- {
		v := ip[i]
		buf = append(buf, hexDigit[v&0xF], '.', hexDigit[v>>4], '.')
	}
	buf = append(buf, "ip6.arpa."...)
	return string(buf), nil
}

// GeneratePTREndpoints generates PTR endpoints for the given A/AAAA endpoints.
// For each A/AAAA endpoint, a corresponding PTR endpoint is created with:
//   - DNSName: the reverse address (e.g. "4.3.2.1.in-addr.arpa")
//   - RecordType: "PTR"
//   - Targets: the original DNSName (with trailing dot)
//   - RecordTTL: preserved from the original endpoint
//
// Non-A/AAAA endpoints are silently skipped.
//
// When multiple A/AAAA records point to the same IP, a single PTR endpoint
// is created with all hostnames as targets, since external-dns tracks
// ownership per (name, type) pair.
//
// If domainFilter is non-nil, only A/AAAA endpoints whose DNSName matches
// the filter will produce PTR records. This prevents PTR creation for
// records that fall outside the managed domain set.
//
// The defaultEnabled parameter is the value of the --create-ptr flag.
// Per-endpoint, the "ptr" provider-specific property (sourced from
// the resource annotation) overrides this default:
//   - annotation "true"  → include the endpoint (even if defaultEnabled is false)
//   - annotation "false" → exclude the endpoint (even if defaultEnabled is true)
//   - annotation absent  → use defaultEnabled
func GeneratePTREndpoints(endpoints []*endpoint.Endpoint, domainFilter endpoint.DomainFilterInterface, defaultEnabled bool) []*endpoint.Endpoint {
	// Collect targets per reverse address, preserving the first TTL seen.
	type ptrInfo struct {
		targets []string
		ttl     endpoint.TTL
	}
	ptrMap := make(map[string]*ptrInfo)
	var order []string // preserve insertion order

	for _, ep := range endpoints {
		if ep.RecordType != endpoint.RecordTypeA && ep.RecordType != endpoint.RecordTypeAAAA {
			continue
		}
		// Skip wildcard records — a PTR pointing to "*.example.com" is not
		// a meaningful reverse mapping.
		if strings.HasPrefix(ep.DNSName, "*.") {
			log.Debugf("PTR: skipping wildcard record %s", ep.DNSName)
			continue
		}
		// Skip A/AAAA records whose hostname does not match the domain filter.
		// This prevents PTR creation for records that will be filtered out of
		// the desired state (e.g. hostname changed to a domain not managed by
		// this external-dns instance).
		if domainFilter != nil && !domainFilter.Match(ep.DNSName) {
			log.Debugf("PTR: skipping %s %s — does not match domain filter", ep.DNSName, ep.RecordType)
			continue
		}
		// Determine whether this endpoint should produce a PTR record.
		// Annotation overrides the CLI flag (configuration precedence).
		enabled := defaultEnabled
		if val, ok := ep.GetProviderSpecificProperty("ptr"); ok {
			enabled = val == "true"
		}
		if !enabled {
			log.Debugf("PTR: skipping %s %s — PTR creation not enabled", ep.DNSName, ep.RecordType)
			continue
		}
		for _, target := range ep.Targets {
			revAddr, err := ReverseAddr(target)
			if err != nil {
				log.Warnf("Failed to compute reverse address for %s: %v", target, err)
				continue
			}
			// Strip trailing dot for the PTR record name (external-dns convention)
			ptrName := revAddr[:len(revAddr)-1]
			hostname := EnsureTrailingDot(ep.DNSName)

			if info, ok := ptrMap[ptrName]; ok {
				info.targets = append(info.targets, hostname)
			} else {
				ptrMap[ptrName] = &ptrInfo{
					targets: []string{hostname},
					ttl:     ep.RecordTTL,
				}
				order = append(order, ptrName)
			}
		}
	}

	ptrEndpoints := make([]*endpoint.Endpoint, 0, len(ptrMap))
	for _, ptrName := range order {
		info := ptrMap[ptrName]
		ptrEp := endpoint.NewEndpointWithTTL(
			ptrName,
			endpoint.RecordTypePTR,
			info.ttl,
			info.targets...,
		)
		ptrEndpoints = append(ptrEndpoints, ptrEp)
	}
	return ptrEndpoints
}

// PTRProvider is a decorator that wraps a Provider to automatically manage
// PTR records for A and AAAA records. It injects PTR endpoints into the
// desired endpoint set via AdjustEndpoints, so the plan can detect missing
// or stale PTR records and generate the appropriate changes.
//
// This allows any provider that manages reverse DNS zones to automatically
// create and delete PTR records without provider-specific PTR logic.
//
// Note: The underlying provider must have authority over the relevant reverse
// zones (e.g. in-addr.arpa. zones) for PTR records to be created successfully.
// Users should include reverse zones in their --domain-filter configuration.
// PTR must also be included in ManagedDNSRecordTypes (added automatically
// when --create-ptr is enabled).
type PTRProvider struct {
	Provider
	defaultEnabled bool
}

// NewPTRProvider creates a new PTRProvider wrapping the given provider.
// The defaultEnabled parameter is the value of the --create-ptr flag.
// Per-resource annotations can override this default.
//
// The underlying provider's GetDomainFilter() is used to skip A/AAAA endpoints
// whose hostname does not match the configured domain filter.
func NewPTRProvider(p Provider, defaultEnabled bool) *PTRProvider {
	return &PTRProvider{Provider: p, defaultEnabled: defaultEnabled}
}

// AdjustEndpoints augments the desired endpoint set with PTR records
// for every A/AAAA endpoint, then delegates to the underlying provider's
// AdjustEndpoints for any provider-specific canonicalization.
//
// This ensures PTR records appear in the plan's desired state so the
// planner can detect missing PTRs and generate Create changes, or
// detect stale PTRs and generate Delete changes.
func (p *PTRProvider) AdjustEndpoints(endpoints []*endpoint.Endpoint) ([]*endpoint.Endpoint, error) {
	ptrEndpoints := GeneratePTREndpoints(endpoints, p.GetDomainFilter(), p.defaultEnabled)
	if len(ptrEndpoints) > 0 {
		log.Debugf("PTR: generated %d PTR endpoints from desired A/AAAA records", len(ptrEndpoints))
		endpoints = append(endpoints, ptrEndpoints...)
	}
	return p.Provider.AdjustEndpoints(endpoints)
}
