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

package endpoint

import (
	"testing"

	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type mockObjectMetaAccessor struct {
	namespace string
	name      string
}

func (m *mockObjectMetaAccessor) GetObjectMeta() metav1.Object {
	return &metav1.ObjectMeta{
		Namespace: m.namespace,
		Name:      m.name,
	}
}

func TestHasEmptyEndpoints(t *testing.T) {
	tests := []struct {
		name      string
		endpoints []*Endpoint
		rType     string
		entity    metav1.ObjectMetaAccessor
		expected  bool
	}{
		{
			name:      "nil endpoints returns true",
			endpoints: nil,
			rType:     "Service",
			entity:    &mockObjectMetaAccessor{namespace: "default", name: "my-service"},
			expected:  true,
		},
		{
			name:      "empty slice returns true",
			endpoints: []*Endpoint{},
			rType:     "Ingress",
			entity:    &mockObjectMetaAccessor{namespace: "kube-system", name: "my-ingress"},
			expected:  true,
		},
		{
			name: "single endpoint returns false",
			endpoints: []*Endpoint{
				NewEndpoint("example.org", "A", "1.2.3.4"),
			},
			rType:    "Service",
			entity:   &mockObjectMetaAccessor{namespace: "default", name: "my-service"},
			expected: false,
		},
		{
			name: "multiple endpoints returns false",
			endpoints: []*Endpoint{
				NewEndpoint("example.org", "A", "1.2.3.4"),
				NewEndpoint("test.example.org", "CNAME", "example.org"),
			},
			rType:    "Ingress",
			entity:   &mockObjectMetaAccessor{namespace: "production", name: "frontend"},
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := HasNoEmptyEndpoints(tc.endpoints, tc.rType, tc.entity)
			assert.Equal(t, tc.expected, result)
			// TODO: Add log capture and verification
		})
	}
}

func TestSuitableType(t *testing.T) {
	tests := []struct {
		name     string
		target   string
		expected string
	}{
		{
			name:     "IPv4 address returns A",
			target:   "192.168.1.1",
			expected: RecordTypeA,
		},
		{
			name:     "another IPv4 address returns A",
			target:   "10.0.0.1",
			expected: RecordTypeA,
		},
		{
			name:     "IPv6 address returns AAAA",
			target:   "2001:db8::1",
			expected: RecordTypeAAAA,
		},
		{
			name:     "full IPv6 address returns AAAA",
			target:   "2001:0db8:0000:0000:0000:0000:0000:0001",
			expected: RecordTypeAAAA,
		},
		{
			name:     "hostname returns CNAME",
			target:   "example.com",
			expected: RecordTypeCNAME,
		},
		{
			name:     "hostname with subdomain returns CNAME",
			target:   "api.example.com",
			expected: RecordTypeCNAME,
		},
		{
			name:     "empty string returns CNAME",
			target:   "",
			expected: RecordTypeCNAME,
		},
		{
			name:     "invalid IP returns CNAME",
			target:   "999.999.999.999",
			expected: RecordTypeCNAME,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := SuitableType(tc.target)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestEndpointsForHostsAndTargets(t *testing.T) {
	tests := []struct {
		name      string
		hostnames []string
		targets   []string
		expected  []*Endpoint
	}{
		{
			name:      "empty hostnames returns nil",
			hostnames: []string{},
			targets:   []string{"192.168.1.1"},
			expected:  nil,
		},
		{
			name:      "nil hostnames returns nil",
			hostnames: nil,
			targets:   []string{"192.168.1.1"},
			expected:  nil,
		},
		{
			name:      "empty targets returns nil",
			hostnames: []string{"example.com"},
			targets:   []string{},
			expected:  nil,
		},
		{
			name:      "nil targets returns nil",
			hostnames: []string{"example.com"},
			targets:   nil,
			expected:  nil,
		},
		{
			name:      "single hostname and IPv4 target",
			hostnames: []string{"example.com"},
			targets:   []string{"192.168.1.1"},
			expected: []*Endpoint{
				NewEndpoint("example.com", RecordTypeA, "192.168.1.1"),
			},
		},
		{
			name:      "single hostname and CNAME target",
			hostnames: []string{"example.com"},
			targets:   []string{"other.example.com"},
			expected: []*Endpoint{
				NewEndpoint("example.com", RecordTypeCNAME, "other.example.com"),
			},
		},
		{
			name:      "single hostname with mixed targets (A and CNAME)",
			hostnames: []string{"example.com"},
			targets:   []string{"192.168.1.1", "other.example.com"},
			expected: []*Endpoint{
				NewEndpoint("example.com", RecordTypeA, "192.168.1.1"),
				NewEndpoint("example.com", RecordTypeCNAME, "other.example.com"),
			},
		},
		{
			name:      "single hostname with IPv4 and IPv6 targets",
			hostnames: []string{"example.com"},
			targets:   []string{"192.168.1.1", "2001:db8::1"},
			expected: []*Endpoint{
				NewEndpoint("example.com", RecordTypeA, "192.168.1.1"),
				NewEndpoint("example.com", RecordTypeAAAA, "2001:db8::1"),
			},
		},
		{
			name:      "multiple hostnames with single target",
			hostnames: []string{"example.com", "www.example.com"},
			targets:   []string{"192.168.1.1"},
			expected: []*Endpoint{
				NewEndpoint("example.com", RecordTypeA, "192.168.1.1"),
				NewEndpoint("www.example.com", RecordTypeA, "192.168.1.1"),
			},
		},
		{
			name:      "multiple hostnames with multiple IPv4 targets",
			hostnames: []string{"example.com", "www.example.com"},
			targets:   []string{"192.168.1.1", "192.168.1.2"},
			expected: []*Endpoint{
				NewEndpoint("example.com", RecordTypeA, "192.168.1.1", "192.168.1.2"),
				NewEndpoint("www.example.com", RecordTypeA, "192.168.1.1", "192.168.1.2"),
			},
		},
		{
			name:      "all record types (A, AAAA, CNAME) deterministic order",
			hostnames: []string{"example.com"},
			targets:   []string{"192.168.1.1", "2001:db8::1", "cname.example.com"},
			expected: []*Endpoint{
				NewEndpoint("example.com", RecordTypeA, "192.168.1.1"),
				NewEndpoint("example.com", RecordTypeAAAA, "2001:db8::1"),
				NewEndpoint("example.com", RecordTypeCNAME, "cname.example.com"),
			},
		},
		{
			name:      "multiple of each type maintains grouping",
			hostnames: []string{"example.com"},
			targets:   []string{"192.168.1.1", "192.168.1.2", "2001:db8::1", "2001:db8::2", "a.example.com", "b.example.com"},
			expected: []*Endpoint{
				NewEndpoint("example.com", RecordTypeA, "192.168.1.1", "192.168.1.2"),
				NewEndpoint("example.com", RecordTypeAAAA, "2001:db8::1", "2001:db8::2"),
				NewEndpoint("example.com", RecordTypeCNAME, "a.example.com", "b.example.com"),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := EndpointsForHostsAndTargets(tc.hostnames, tc.targets)
			if tc.expected == nil {
				assert.Nil(t, result)
				return
			}
			assert.Len(t, result, len(tc.expected))
			for i, ep := range result {
				assert.Equal(t, tc.expected[i].DNSName, ep.DNSName, "DNSName mismatch at index %d", i)
				assert.Equal(t, tc.expected[i].RecordType, ep.RecordType, "RecordType mismatch at index %d", i)
				assert.ElementsMatch(t, tc.expected[i].Targets, ep.Targets, "Targets mismatch at index %d", i)
			}
		})
	}
}

func TestEndpointsForHostsAndTargets_DeterministicOrder(t *testing.T) {
	// Run multiple times to verify deterministic ordering
	hostnames := []string{"example.com"}
	targets := []string{"cname.example.com", "192.168.1.1", "2001:db8::1"}

	for range 10 {
		result := EndpointsForHostsAndTargets(hostnames, targets)
		assert.Len(t, result, 3)
		// Order should always be A, AAAA, CNAME (alphabetically sorted)
		assert.Equal(t, RecordTypeA, result[0].RecordType)
		assert.Equal(t, RecordTypeAAAA, result[1].RecordType)
		assert.Equal(t, RecordTypeCNAME, result[2].RecordType)
	}
}
