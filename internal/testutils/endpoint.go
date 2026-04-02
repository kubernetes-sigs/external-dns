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
	"cmp"
	"fmt"
	"maps"
	"math/rand"
	"net/netip"
	"reflect"
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	ctrlclient "sigs.k8s.io/controller-runtime/pkg/client"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/events"
)

// ValidateEndpoints asserts that endpoints and expected have the same length and equal fields.
func ValidateEndpoints(t *testing.T, endpoints, expected []*endpoint.Endpoint) {
	t.Helper()

	if len(endpoints) != len(expected) {
		t.Fatalf("expected %d endpoints, got %d", len(expected), len(endpoints))
	}

	// Make sure endpoints are sorted - validateEndpoint() depends on it.
	sortEndpoints(endpoints)
	sortEndpoints(expected)

	for i := range endpoints {
		validateEndpoint(t, endpoints[i], expected[i])
	}
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
		sameProviderSpecific(a.ProviderSpecific, b.ProviderSpecific)
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
	sa, sb := cloneAndSort(a), cloneAndSort(b)
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
	sa, sb := cloneAndSort(a), cloneAndSort(b)
	for i := range sa {
		if !maps.Equal(sa[i].Labels, sb[i].Labels) {
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
//	endpoints will contain 2 A records and 1 CNAME record with unique DNS names and targets.
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

// NewEndpointWithRef builds an endpoint attached to a Kubernetes object reference.
// The record type is inferred from target: A for IPv4, AAAA for IPv6, CNAME otherwise.
// Kind and APIVersion are resolved from the client-go scheme, so TypeMeta need not be set on obj.
func NewEndpointWithRef(dns, target string, obj ctrlclient.Object, source string) *endpoint.Endpoint {
	return endpoint.NewEndpoint(dns, endpoint.SuitableType(target), target).
		WithRefObject(events.NewObjectReference(obj, source))
}

// AssertEndpointsHaveRefObject asserts that endpoints have the expected count
// and each endpoint has a non-nil RefObject with the expected source type.
func AssertEndpointsHaveRefObject(
	t *testing.T,
	endpoints []*endpoint.Endpoint,
	expectedSource string,
	expectedCount int) {
	t.Helper()
	assert.Len(t, endpoints, expectedCount)
	for _, ep := range endpoints {
		assert.NotNil(t, ep.RefObject())
		assert.NotEmpty(t, ep.RefObject().UID)
		assert.Equal(t, expectedSource, ep.RefObject().Source)
	}
}

// sortEndpoints sorts a slice of endpoints in-place by DNSName, RecordType, and Targets.
// It also deduplicates and sorts each endpoint's Targets so the ordering is stable.
func sortEndpoints(endpoints []*endpoint.Endpoint) {
	for _, ep := range endpoints {
		if ep != nil {
			ep.Targets = endpoint.NewTargets(ep.Targets...)
		}
	}
	slices.SortFunc(endpoints, compareEndpoints)
}

// validateEndpoint asserts that two endpoints have equal fields.
func validateEndpoint(t *testing.T, ep, expected *endpoint.Endpoint) {
	t.Helper()

	if ep == nil || expected == nil {
		if ep != nil || expected != nil {
			t.Errorf("one endpoint is nil: got %v, expected %v", ep, expected)
		}
		return
	}

	if ep.DNSName != expected.DNSName {
		t.Errorf("DNSName expected %q, got %q", expected.DNSName, ep.DNSName)
	}

	if !ep.Targets.Same(expected.Targets) {
		t.Errorf("Targets expected %q, got %q", expected.Targets, ep.Targets)
	}

	if ep.RecordTTL != expected.RecordTTL {
		t.Errorf("RecordTTL expected %v, got %v", expected.RecordTTL, ep.RecordTTL)
	}

	if ep.RecordType != expected.RecordType {
		t.Errorf("RecordType expected %q, got %q", expected.RecordType, ep.RecordType)
	}

	if expected.Labels != nil && !reflect.DeepEqual(ep.Labels, expected.Labels) {
		t.Errorf("Labels expected %s, got %s", expected.Labels, ep.Labels)
	}

	if (len(expected.ProviderSpecific) != 0 || len(ep.ProviderSpecific) != 0) &&
		!reflect.DeepEqual(ep.ProviderSpecific, expected.ProviderSpecific) {
		t.Errorf("ProviderSpecific expected %s, got %s", expected.ProviderSpecific, ep.ProviderSpecific)
	}

	if ep.SetIdentifier != expected.SetIdentifier {
		t.Errorf("SetIdentifier expected %q, got %q", expected.SetIdentifier, ep.SetIdentifier)
	}
}

// cloneAndSort returns a sorted copy of endpoints; the original slice is not modified.
func cloneAndSort(eps []*endpoint.Endpoint) []*endpoint.Endpoint {
	s := slices.Clone(eps)
	slices.SortFunc(s, compareEndpoints)
	return s
}

// compareEndpoints returns a cmp-style ordering value for two endpoints.
// Nil endpoints sort before non-nil.
func compareEndpoints(a, b *endpoint.Endpoint) int {
	if a == nil || b == nil {
		if a == nil {
			return -1
		}
		return 1
	}
	if n := cmp.Compare(a.DNSName, b.DNSName); n != 0 {
		return n
	}
	if n := cmp.Compare(a.RecordType, b.RecordType); n != 0 {
		return n
	}
	return cmp.Compare(a.Targets.String(), b.Targets.String())
}

// sameProviderSpecific verifies that two slices contain the same key/value pairs.
func sameProviderSpecific(a, b endpoint.ProviderSpecific) bool {
	if len(a) != len(b) {
		return false
	}
	sa := slices.Clone(a)
	sb := slices.Clone(b)
	cmpProp := func(x, y endpoint.ProviderSpecificProperty) int { return cmp.Compare(x.Name, y.Name) }
	slices.SortFunc(sa, cmpProp)
	slices.SortFunc(sb, cmpProp)
	for i := range sa {
		if sa[i] != sb[i] {
			return false
		}
	}
	return true
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
