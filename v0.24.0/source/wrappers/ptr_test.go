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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	"sigs.k8s.io/external-dns/source"
)

var _ source.Source = &ptrSource{}

func TestPTRSource(t *testing.T) {
	tests := []struct {
		name           string
		defaultEnabled bool
		endpoints      []*endpoint.Endpoint
		expected       []*endpoint.Endpoint
	}{
		{
			name:           "A record produces PTR",
			defaultEnabled: true,
			endpoints: []*endpoint.Endpoint{
				{DNSName: "web.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.49.2"}},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "web.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.49.2"}},
				{DNSName: "2.49.168.192.in-addr.arpa", RecordType: endpoint.RecordTypePTR, Targets: endpoint.Targets{"web.example.com"}},
			},
		},
		{
			name:           "disabled by default, no PTR",
			defaultEnabled: false,
			endpoints: []*endpoint.Endpoint{
				{DNSName: "web.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.49.2"}},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "web.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.49.2"}},
			},
		},
		{
			name:           "CNAME skipped",
			defaultEnabled: true,
			endpoints: []*endpoint.Endpoint{
				{DNSName: "alias.example.com", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"web.example.com"}},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "alias.example.com", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"web.example.com"}},
			},
		},
		{
			name:           "wildcard skipped",
			defaultEnabled: true,
			endpoints: []*endpoint.Endpoint{
				{DNSName: "*.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.49.2"}},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "*.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.49.2"}},
			},
		},
		{
			name:           "same IP merges into single PTR",
			defaultEnabled: true,
			endpoints: []*endpoint.Endpoint{
				{DNSName: "a.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.49.2"}},
				{DNSName: "b.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.49.2"}},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "a.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.49.2"}},
				{DNSName: "b.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.49.2"}},
				{DNSName: "2.49.168.192.in-addr.arpa", RecordType: endpoint.RecordTypePTR, Targets: endpoint.Targets{"a.example.com", "b.example.com"}},
			},
		},
		{
			name:           "TTL preserved",
			defaultEnabled: true,
			endpoints: []*endpoint.Endpoint{
				endpoint.NewEndpointWithTTL("web.example.com", endpoint.RecordTypeA, 300, "10.0.0.1"),
			},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpointWithTTL("web.example.com", endpoint.RecordTypeA, 300, "10.0.0.1"),
				endpoint.NewEndpointWithTTL("1.0.0.10.in-addr.arpa", endpoint.RecordTypePTR, 300, "web.example.com"),
			},
		},
		{
			name:           "conflicting TTLs use minimum",
			defaultEnabled: true,
			endpoints: []*endpoint.Endpoint{
				endpoint.NewEndpointWithTTL("a.example.com", endpoint.RecordTypeA, 300, "10.0.0.1"),
				endpoint.NewEndpointWithTTL("b.example.com", endpoint.RecordTypeA, 60, "10.0.0.1"),
			},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpointWithTTL("a.example.com", endpoint.RecordTypeA, 300, "10.0.0.1"),
				endpoint.NewEndpointWithTTL("b.example.com", endpoint.RecordTypeA, 60, "10.0.0.1"),
				endpoint.NewEndpointWithTTL("1.0.0.10.in-addr.arpa", endpoint.RecordTypePTR, 60, "a.example.com", "b.example.com"),
			},
		},
		{
			name:           "conflicting TTLs use minimum reversed order",
			defaultEnabled: true,
			endpoints: []*endpoint.Endpoint{
				endpoint.NewEndpointWithTTL("a.example.com", endpoint.RecordTypeA, 60, "10.0.0.1"),
				endpoint.NewEndpointWithTTL("b.example.com", endpoint.RecordTypeA, 300, "10.0.0.1"),
			},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpointWithTTL("a.example.com", endpoint.RecordTypeA, 60, "10.0.0.1"),
				endpoint.NewEndpointWithTTL("b.example.com", endpoint.RecordTypeA, 300, "10.0.0.1"),
				endpoint.NewEndpointWithTTL("1.0.0.10.in-addr.arpa", endpoint.RecordTypePTR, 60, "a.example.com", "b.example.com"),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockSource := new(testutils.MockSource)
			mockSource.On("Endpoints").Return(tc.endpoints, nil)

			src := NewPTRSource(mockSource, tc.defaultEnabled)
			result, err := src.Endpoints(t.Context())
			require.NoError(t, err)
			assert.Len(t, result, len(tc.expected))
			for i, ep := range result {
				assert.Equal(t, tc.expected[i].DNSName, ep.DNSName)
				assert.Equal(t, tc.expected[i].RecordType, ep.RecordType)
				assert.Equal(t, tc.expected[i].RecordTTL, ep.RecordTTL)
				assert.ElementsMatch(t, tc.expected[i].Targets, ep.Targets)
			}
		})
	}
}

func TestPTRSource_AnnotationOverride(t *testing.T) {
	t.Run("annotation opts in when flag is off", func(t *testing.T) {
		eps := []*endpoint.Endpoint{
			endpoint.NewEndpoint("web.example.com", endpoint.RecordTypeA, "192.168.49.2").
				WithProviderSpecific(endpoint.ProviderSpecificRecordType, "ptr"),
		}
		mockSource := testutils.NewMockSource(eps...)
		src := NewPTRSource(mockSource, false)
		result, err := src.Endpoints(t.Context())
		require.NoError(t, err)
		assert.Len(t, result, 2)
		assert.Equal(t, endpoint.RecordTypePTR, result[1].RecordType)
		// provider-specific property should be removed after processing
		_, ok := result[0].GetProviderSpecificProperty(endpoint.ProviderSpecificRecordType)
		assert.False(t, ok, "record-type property should be removed from original endpoint")
	})

	t.Run("annotation opts out when flag is on", func(t *testing.T) {
		eps := []*endpoint.Endpoint{
			endpoint.NewEndpoint("web.example.com", endpoint.RecordTypeA, "192.168.49.2").
				WithProviderSpecific(endpoint.ProviderSpecificRecordType, ""),
		}
		mockSource := testutils.NewMockSource(eps...)
		src := NewPTRSource(mockSource, true)
		result, err := src.Endpoints(t.Context())
		require.NoError(t, err)
		assert.Len(t, result, 1) // only the original A record
		// provider-specific property should be removed after processing
		_, ok := result[0].GetProviderSpecificProperty(endpoint.ProviderSpecificRecordType)
		assert.False(t, ok, "record-type property should be removed from original endpoint")
	})

	t.Run("no annotation uses flag default true", func(t *testing.T) {
		eps := []*endpoint.Endpoint{
			endpoint.NewEndpoint("web.example.com", endpoint.RecordTypeA, "192.168.49.2"),
		}
		mockSource := testutils.NewMockSource(eps...)
		src := NewPTRSource(mockSource, true)
		result, err := src.Endpoints(t.Context())
		require.NoError(t, err)
		assert.Len(t, result, 2)
	})

	t.Run("no annotation uses flag default false", func(t *testing.T) {
		eps := []*endpoint.Endpoint{
			endpoint.NewEndpoint("web.example.com", endpoint.RecordTypeA, "192.168.49.2"),
		}
		mockSource := testutils.NewMockSource(eps...)
		src := NewPTRSource(mockSource, false)
		result, err := src.Endpoints(t.Context())
		require.NoError(t, err)
		assert.Len(t, result, 1)
	})
}

func TestPTRSource_IPv6(t *testing.T) {
	eps := []*endpoint.Endpoint{
		{DNSName: "v6.example.com", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:db8::1"}},
	}
	mockSource := testutils.NewMockSource(eps...)
	src := NewPTRSource(mockSource, true)
	result, err := src.Endpoints(t.Context())
	require.NoError(t, err)
	require.Len(t, result, 2)
	assert.Equal(t, "1.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.8.b.d.0.1.0.0.2.ip6.arpa", result[1].DNSName)
	assert.Equal(t, endpoint.RecordTypePTR, result[1].RecordType)
}

func TestPTRSource_AddEventHandler(t *testing.T) {
	mockSource := testutils.NewMockSource()

	src := NewPTRSource(mockSource, true)
	src.AddEventHandler(t.Context(), func() {})

	mockSource.AssertNumberOfCalls(t, "AddEventHandler", 1)
}
