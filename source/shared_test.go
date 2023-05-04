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
	"reflect"
	"sort"
	"testing"

	"sigs.k8s.io/external-dns/endpoint"
)

func sortEndpoints(endpoints []*endpoint.Endpoint) {
	for _, ep := range endpoints {
		sort.Strings([]string(ep.Targets))
	}
	sort.Slice(endpoints, func(i, k int) bool {
		// Sort by DNSName and Targets
		ei, ek := endpoints[i], endpoints[k]
		if ei.DNSName != ek.DNSName {
			return ei.DNSName < ek.DNSName
		}
		// Targets are sorted ahead of time.
		for j, ti := range ei.Targets {
			if j >= len(ek.Targets) {
				return true
			}
			if tk := ek.Targets[j]; ti != tk {
				return ti < tk
			}
		}
		return false
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
