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

package source

import (
	"cmp"
	"maps"
	"reflect"
	"slices"
	"testing"

	"sigs.k8s.io/external-dns/endpoint"
)

func sortEndpoints(endpoints []*endpoint.Endpoint) {
	slices.SortStableFunc(endpoints, func(a, b *endpoint.Endpoint) int {
		if c := cmp.Compare(a.DNSName, b.DNSName); c != 0 {
			return c
		}
		if c := cmp.Compare(a.RecordType, b.RecordType); c != 0 {
			return c
		}
		if c := slices.Compare(slices.Sorted(slices.Values(a.Targets)), slices.Sorted(slices.Values(b.Targets))); c != 0 {
			return c
		}
		if c := slices.Compare(slices.Sorted(maps.Keys(a.Labels)), slices.Sorted(maps.Keys(b.Labels))); c != 0 {
			return c
		}
		for key, value := range a.Labels {
			if c := cmp.Compare(value, b.Labels[key]); c != 0 {
				return c
			}
		}
		return 0
	})
}

func validateEndpoints(t *testing.T, endpoints, expected []*endpoint.Endpoint) {
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

func validateEndpoint(t *testing.T, endpoint, expected *endpoint.Endpoint) {
	t.Helper()

	if endpoint.DNSName != expected.DNSName {
		t.Errorf("DNSName expected %q, got %q", expected.DNSName, endpoint.DNSName)
	}

	if !endpoint.Targets.Same(expected.Targets) {
		t.Errorf("Targets expected %q, got %q", expected.Targets, endpoint.Targets)
	}

	if endpoint.RecordTTL != expected.RecordTTL {
		t.Errorf("RecordTTL expected %v, got %v", expected.RecordTTL, endpoint.RecordTTL)
	}

	// if non-empty record type is expected, check that it matches.
	if endpoint.RecordType != expected.RecordType {
		t.Errorf("RecordType expected %q, got %q", expected.RecordType, endpoint.RecordType)
	}

	// if non-empty labels are expected, check that they matches.
	if expected.Labels != nil && !reflect.DeepEqual(endpoint.Labels, expected.Labels) {
		t.Errorf("Labels expected %s, got %s", expected.Labels, endpoint.Labels)
	}

	if (len(expected.ProviderSpecific) != 0 || len(endpoint.ProviderSpecific) != 0) &&
		!reflect.DeepEqual(endpoint.ProviderSpecific, expected.ProviderSpecific) {
		t.Errorf("ProviderSpecific expected %s, got %s", expected.ProviderSpecific, endpoint.ProviderSpecific)
	}

	if endpoint.SetIdentifier != expected.SetIdentifier {
		t.Errorf("SetIdentifier expected %q, got %q", expected.SetIdentifier, endpoint.SetIdentifier)
	}
}
