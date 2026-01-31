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

package testutils

import (
	"fmt"
	"maps"
	"math/rand"
	"net/netip"
	"reflect"
	"slices"
	"sort"
	"strings"

	"sigs.k8s.io/external-dns/endpoint"
)

/** test utility functions for endpoints verifications */

type byNames endpoint.ProviderSpecific

func (p byNames) Len() int           { return len(p) }
func (p byNames) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p byNames) Less(i, j int) bool { return p[i].Name < p[j].Name }

type byAllFields []*endpoint.Endpoint

func (b byAllFields) Len() int      { return len(b) }
func (b byAllFields) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b byAllFields) Less(i, j int) bool {
	if b[i].DNSName < b[j].DNSName {
		return true
	}
	if b[i].DNSName == b[j].DNSName {
		// This rather bad, we need a more complex comparison for Targets, which considers all elements
		if b[i].Targets.Same(b[j].Targets) {
			if b[i].RecordType == (b[j].RecordType) {
				sa := b[i].ProviderSpecific
				sb := b[j].ProviderSpecific
				sort.Sort(byNames(sa))
				sort.Sort(byNames(sb))
				return reflect.DeepEqual(sa, sb)
			}
			return b[i].RecordType <= b[j].RecordType
		}
		return b[i].Targets.String() <= b[j].Targets.String()
	}
	return false
}

// SameEndpoint returns true if two endpoints are same
// considers example.org. and example.org DNSName/Target as different endpoints
func SameEndpoint(a, b *endpoint.Endpoint) bool {
	if a == nil || b == nil {
		return a == b
	}
	return a.DNSName == b.DNSName && a.Targets.Same(b.Targets) && a.RecordType == b.RecordType && a.SetIdentifier == b.SetIdentifier &&
		a.Labels[endpoint.OwnerLabelKey] == b.Labels[endpoint.OwnerLabelKey] && a.RecordTTL == b.RecordTTL &&
		a.Labels[endpoint.ResourceLabelKey] == b.Labels[endpoint.ResourceLabelKey] &&
		a.Labels[endpoint.OwnedRecordLabelKey] == b.Labels[endpoint.OwnedRecordLabelKey] &&
		SameProviderSpecific(a.ProviderSpecific, b.ProviderSpecific)
}

// SameEndpoints compares two slices of endpoints regardless of order
// [x,y,z] == [z,x,y]
// [x,x,z] == [x,z,x]
// [x,y,y] != [x,x,y]
// [x,x,x] != [x,x,z]
func SameEndpoints(a, b []*endpoint.Endpoint) bool {
	if len(a) != len(b) {
		return false
	}

	sa := a
	sb := b
	sort.Sort(byAllFields(sa))
	sort.Sort(byAllFields(sb))

	for i := range sa {
		if !SameEndpoint(sa[i], sb[i]) {
			return false
		}
	}
	return true
}

// SameEndpointLabels verifies that labels of the two slices of endpoints are the same
func SameEndpointLabels(a, b []*endpoint.Endpoint) bool {
	if len(a) != len(b) {
		return false
	}

	sa := a
	sb := b
	sort.Sort(byAllFields(sa))
	sort.Sort(byAllFields(sb))

	for i := range sa {
		if !reflect.DeepEqual(sa[i].Labels, sb[i].Labels) {
			return false
		}
	}
	return true
}

// SamePlanChanges verifies that two set of changes are the same
func SamePlanChanges(a, b map[string][]*endpoint.Endpoint) bool {
	return SameEndpoints(a["Create"], b["Create"]) && SameEndpoints(a["Delete"], b["Delete"]) &&
		SameEndpoints(a["UpdateOld"], b["UpdateOld"]) && SameEndpoints(a["UpdateNew"], b["UpdateNew"])
}

// SameProviderSpecific verifies that two maps contain the same string/string key/value pairs
func SameProviderSpecific(a, b endpoint.ProviderSpecific) bool {
	sa := a
	sb := b
	sort.Sort(byNames(sa))
	sort.Sort(byNames(sb))
	return reflect.DeepEqual(sa, sb)
}

// NewTargetsFromAddr convert an array of netip.Addr to Targets (array of string)
func NewTargetsFromAddr(targets []netip.Addr) endpoint.Targets {
	t := make(endpoint.Targets, len(targets))
	for i, target := range targets {
		t[i] = target.String()
	}
	return t
}

// GenerateTestEndpointsByType generates a shuffled slice of test Endpoints for each record type and count specified in typeCounts.
// Usage example:
//
//	endpoints := GenerateTestEndpointsByType(map[string]int{"A": 2, "CNAME": 1})
//	// endpoints will contain 2 A records and 1 CNAME record with unique DNS names and targets.
func GenerateTestEndpointsByType(typeCounts map[string]int) []*endpoint.Endpoint {
	return GenerateTestEndpointsWithDistribution(typeCounts, map[string]int{"example.com": 1}, nil)
}

// GenerateTestEndpointsWithDistribution generates test endpoints with specified distributions
// of record types, domains, and owners.
// - typeCounts: maps record type (e.g., "A", "CNAME") to how many endpoints of that type to create
// - domainWeights: maps domain suffix to weight; domains are distributed proportionally
// - ownerWeights: maps owner ID to weight; owners are distributed proportionally
//
// The total number of endpoints equals the sum of typeCounts values.
// Weights represent ratios: {"example.com": 2, "test.org": 1} means ~66% example.com, ~33% test.org
//
// Example:
//
//	endpoints := GenerateTestEndpointsWithDistribution(
//	    map[string]int{"A": 6, "CNAME": 4},       // 10 endpoints total
//	    map[string]int{"example.com": 2, "test.org": 1},  // ~66% example.com, ~33% test.org
//	    map[string]int{"owner1": 3, "owner2": 1},         // ~75% owner1, ~25% owner2
//	)
func GenerateTestEndpointsWithDistribution(
	typeCounts map[string]int,
	domainWeights map[string]int,
	ownerWeights map[string]int,
) []*endpoint.Endpoint {
	// Calculate total endpoints
	totalEndpoints := 0
	for _, count := range typeCounts {
		totalEndpoints += count
	}

	// Build domain distribution (sorted keys for determinism)
	domainKeys := slices.Sorted(maps.Keys(domainWeights))
	domains := distributeByWeight(domainKeys, domainWeights, totalEndpoints)

	// Build owner distribution (sorted keys for determinism)
	ownerKeys := slices.Sorted(maps.Keys(ownerWeights))
	owners := distributeByWeight(ownerKeys, ownerWeights, totalEndpoints)

	// Sort record types for deterministic iteration
	typeKeys := slices.Sorted(maps.Keys(typeCounts))

	var result []*endpoint.Endpoint
	idx := 0
	for _, rt := range typeKeys {
		count := typeCounts[rt]
		for range count {
			// Determine domain from distribution or use default
			domain := "example.com"
			if idx < len(domains) {
				domain = domains[idx]
			}

			// Create endpoint with labels
			ep := &endpoint.Endpoint{
				DNSName:    fmt.Sprintf("%s-%d.%s", strings.ToLower(rt), idx, domain),
				Targets:    endpoint.Targets{fmt.Sprintf("192.0.2.%d", idx)},
				RecordType: rt,
				RecordTTL:  300,
				Labels:     endpoint.Labels{},
			}

			// Assign owner from distribution
			if idx < len(owners) {
				ep.Labels[endpoint.OwnerLabelKey] = owners[idx]
			}

			result = append(result, ep)
			idx++
		}
	}

	rand.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})
	return result
}

// distributeByWeight distributes n items according to weights.
// Returns a slice of length n with items distributed proportionally.
func distributeByWeight(keys []string, weights map[string]int, n int) []string {
	if len(keys) == 0 || n == 0 {
		return nil
	}

	totalWeight := 0
	for _, key := range keys {
		totalWeight += weights[key]
	}
	if totalWeight == 0 {
		return nil
	}

	result := make([]string, 0, n)
	for _, key := range keys {
		count := (weights[key] * n) / totalWeight
		for range count {
			result = append(result, key)
		}
	}

	// Fill any remaining slots due to rounding with the last key
	for len(result) < n {
		result = append(result, keys[len(keys)-1])
	}

	return result
}

// EndpointGeneratorConfig holds configuration for generating test endpoints.
type EndpointGeneratorConfig struct {
	// TypeCounts maps record type (e.g., "A", "CNAME") to how many endpoints of that type to create.
	TypeCounts map[string]int
	// DomainWeights maps domain suffix to weight; domains are distributed proportionally.
	DomainWeights map[string]int
	// OwnerWeights maps owner ID to weight; owners are distributed proportionally.
	OwnerWeights map[string]int
	// ProviderSpecificCount specifies how many provider-specific properties each endpoint should have.
	ProviderSpecificCount int
}

// GenerateTestEndpoints generates test endpoints based on the provided configuration.
// This is a flexible generator that supports all endpoint configuration options.
//
// Example:
//
//	endpoints := GenerateTestEndpoints(EndpointGeneratorConfig{
//	    TypeCounts:            map[string]int{"A": 100},
//	    DomainWeights:         map[string]int{"example.com": 1},
//	    ProviderSpecificCount: 5,
//	})
func GenerateTestEndpoints(cfg EndpointGeneratorConfig) []*endpoint.Endpoint {
	// Calculate total endpoints
	totalEndpoints := 0
	for _, count := range cfg.TypeCounts {
		totalEndpoints += count
	}

	// Set default domain if not specified
	if len(cfg.DomainWeights) == 0 {
		cfg.DomainWeights = map[string]int{"example.com": 1}
	}

	// Build domain distribution (sorted keys for determinism)
	domainKeys := slices.Sorted(maps.Keys(cfg.DomainWeights))
	domains := distributeByWeight(domainKeys, cfg.DomainWeights, totalEndpoints)

	// Build owner distribution (sorted keys for determinism)
	ownerKeys := slices.Sorted(maps.Keys(cfg.OwnerWeights))
	owners := distributeByWeight(ownerKeys, cfg.OwnerWeights, totalEndpoints)

	// Sort record types for deterministic iteration
	typeKeys := slices.Sorted(maps.Keys(cfg.TypeCounts))

	var result []*endpoint.Endpoint
	idx := 0
	for _, rt := range typeKeys {
		count := cfg.TypeCounts[rt]
		for range count {
			// Determine domain from distribution or use default
			domain := "example.com"
			if idx < len(domains) {
				domain = domains[idx]
			}

			// Create endpoint with labels
			ep := &endpoint.Endpoint{
				DNSName:          fmt.Sprintf("%s-%d.%s", strings.ToLower(rt), idx, domain),
				Targets:          endpoint.Targets{fmt.Sprintf("192.0.2.%d", idx%256)},
				RecordType:       rt,
				RecordTTL:        300,
				Labels:           endpoint.Labels{},
				ProviderSpecific: generateProviderSpecificProperties(cfg.ProviderSpecificCount),
			}

			// Assign owner from distribution
			if idx < len(owners) {
				ep.Labels[endpoint.OwnerLabelKey] = owners[idx]
			}

			result = append(result, ep)
			idx++
		}
	}

	rand.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})
	return result
}

// generateProviderSpecificProperties creates a slice of ProviderSpecificProperty with the given count.
func generateProviderSpecificProperties(count int) endpoint.ProviderSpecific {
	if count <= 0 {
		return nil
	}
	props := make(endpoint.ProviderSpecific, count)
	for i := range count {
		props[i] = endpoint.ProviderSpecificProperty{
			Name:  fmt.Sprintf("property-%d", i),
			Value: fmt.Sprintf("value-%d", i),
		}
	}
	return props
}
