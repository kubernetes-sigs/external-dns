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

	"sigs.k8s.io/external-dns/pkg/events"
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

func TestSuitableType(t *testing.T) {
	tests := []struct {
		target   string
		expected string
	}{
		// IPv4
		{"192.168.1.1", RecordTypeA},
		{"255.255.255.255", RecordTypeA},
		{"0.0.0.0", RecordTypeA},
		// IPv6
		{"2001:0db8:85a3:0000:0000:8a2e:0370:7334", RecordTypeAAAA},
		{"2001:db8:85a3::8a2e:370:7334", RecordTypeAAAA},
		{"::ffff:192.168.20.3", RecordTypeAAAA}, // IPv4-mapped IPv6
		{"::1", RecordTypeAAAA},
		{"::", RecordTypeAAAA},
		// CNAME (hostname or invalid)
		{"example.com", RecordTypeCNAME},
		{"", RecordTypeCNAME},
		{"256.256.256.256", RecordTypeCNAME},
		{"192.168.0.1/22", RecordTypeCNAME},
		{"192.168.1", RecordTypeCNAME},
		{"abc.def.ghi.jkl", RecordTypeCNAME},
	}

	for _, tt := range tests {
		t.Run(tt.target, func(t *testing.T) {
			assert.Equal(t, tt.expected, SuitableType(tt.target))
		})
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

func TestEndpointsForHostname(t *testing.T) {
	tests := []struct {
		name             string
		hostname         string
		targets          Targets
		ttl              TTL
		providerSpecific ProviderSpecific
		setIdentifier    string
		resource         string
		expected         []*Endpoint
	}{
		{
			name:     "A record targets",
			hostname: "example.com",
			targets:  Targets{"192.0.2.1", "192.0.2.2"},
			ttl:      TTL(300),
			providerSpecific: ProviderSpecific{
				{Name: "provider", Value: "value"},
			},
			setIdentifier: "identifier",
			resource:      "resource",
			expected: []*Endpoint{
				{
					DNSName:          "example.com",
					Targets:          Targets{"192.0.2.1", "192.0.2.2"},
					RecordType:       RecordTypeA,
					RecordTTL:        TTL(300),
					ProviderSpecific: ProviderSpecific{{Name: "provider", Value: "value"}},
					SetIdentifier:    "identifier",
					Labels:           map[string]string{ResourceLabelKey: "resource"},
				},
			},
		},
		{
			name:     "AAAA record targets",
			hostname: "example.com",
			targets:  Targets{"2001:db8::1", "2001:db8::2"},
			ttl:      TTL(300),
			providerSpecific: ProviderSpecific{
				{Name: "provider", Value: "value"},
			},
			setIdentifier: "identifier",
			resource:      "resource",
			expected: []*Endpoint{
				{
					DNSName:          "example.com",
					Targets:          Targets{"2001:db8::1", "2001:db8::2"},
					RecordType:       RecordTypeAAAA,
					RecordTTL:        TTL(300),
					ProviderSpecific: ProviderSpecific{{Name: "provider", Value: "value"}},
					SetIdentifier:    "identifier",
					Labels:           map[string]string{ResourceLabelKey: "resource"},
				},
			},
		},
		{
			name:     "CNAME record targets",
			hostname: "example.com",
			targets:  Targets{"cname.example.com"},
			ttl:      TTL(300),
			providerSpecific: ProviderSpecific{
				{Name: "provider", Value: "value"},
			},
			setIdentifier: "identifier",
			resource:      "resource",
			expected: []*Endpoint{
				{
					DNSName:          "example.com",
					Targets:          Targets{"cname.example.com"},
					RecordType:       RecordTypeCNAME,
					RecordTTL:        TTL(300),
					ProviderSpecific: ProviderSpecific{{Name: "provider", Value: "value"}},
					SetIdentifier:    "identifier",
					Labels:           map[string]string{ResourceLabelKey: "resource"},
				},
			},
		},
		{
			name:             "No targets",
			hostname:         "example.com",
			targets:          Targets{},
			ttl:              TTL(300),
			providerSpecific: ProviderSpecific{},
			setIdentifier:    "",
			resource:         "",
			expected:         []*Endpoint(nil),
		},
		{
			// NewEndpointWithTTL returns nil for labels > 63 chars; the nil is skipped.
			name:     "hostname with label exceeding 63 characters returns no endpoints",
			hostname: "this-label-is-way-too-long-and-exceeds-the-sixty-three-character-limit.example.com",
			targets:  Targets{"192.0.2.1"},
			expected: []*Endpoint(nil),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EndpointsForHostname(tt.hostname, tt.targets, tt.ttl, tt.providerSpecific, tt.setIdentifier, tt.resource)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestAttachRefObject(t *testing.T) {
	ref := &events.ObjectReference{Kind: "Service", Namespace: "default", Name: "svc"}
	eps := []*Endpoint{
		NewEndpoint("a.example.com", RecordTypeA, "1.2.3.4"),
		NewEndpoint("b.example.com", RecordTypeA, "5.6.7.8"),
	}

	AttachRefObject(eps, ref)

	for _, ep := range eps {
		assert.Same(t, ref, ep.RefObject())
	}
}
