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
	"testing"

	"github.com/kubernetes-incubator/external-dns/endpoint"
)

// test helper functions

func validateEndpoints(t *testing.T, endpoints, expected []*endpoint.Endpoint) {
	if len(endpoints) != len(expected) {
		t.Fatalf("expected %d endpoints, got %d", len(expected), len(endpoints))
	}

	for i := range endpoints {
		validateEndpoint(t, endpoints[i], expected[i])
	}
}

func validateEndpoint(t *testing.T, ep, expected *endpoint.Endpoint) {
	if ep.DNSName != expected.DNSName {
		t.Errorf("expected %s, got %s", expected.DNSName, ep.DNSName)
	}

	if !endpoint.TargetSliceEquals(ep.Targets, expected.Targets) {
		t.Errorf("expected %v, got %v", expected.Targets, ep.Targets)
	}

	// if non-empty record type is expected, check that it matches.
	if expected.RecordType != "" && ep.RecordType != expected.RecordType {
		t.Errorf("expected %s, got %s", expected.RecordType, ep.RecordType)
	}
}
