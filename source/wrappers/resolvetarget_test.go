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
	"errors"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	"sigs.k8s.io/external-dns/pkg/events"
	"sigs.k8s.io/external-dns/plan"
)

func TestResolveTargetEndpoints(t *testing.T) {
	tests := []struct {
		title     string
		lookupIP  func(string) ([]net.IP, error)
		endpoints []*endpoint.Endpoint
		expected  []*endpoint.Endpoint
	}{
		{
			title:    "A/AAAA endpoints pass through unchanged",
			lookupIP: func(string) ([]net.IP, error) { return nil, errors.New("should not be called") },
			endpoints: []*endpoint.Endpoint{
				endpoint.NewEndpoint("foo.example.com", endpoint.RecordTypeA, "1.2.3.4"),
				endpoint.NewEndpoint("foo.example.com", endpoint.RecordTypeAAAA, "2001:db8::1"),
			},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("foo.example.com", endpoint.RecordTypeA, "1.2.3.4"),
				endpoint.NewEndpoint("foo.example.com", endpoint.RecordTypeAAAA, "2001:db8::1"),
			},
		},
		{
			title:    "nil endpoints are skipped",
			lookupIP: func(string) ([]net.IP, error) { return nil, errors.New("should not be called") },
			endpoints: []*endpoint.Endpoint{
				nil,
				endpoint.NewEndpoint("foo.example.com", endpoint.RecordTypeA, "1.2.3.4"),
				nil,
			},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("foo.example.com", endpoint.RecordTypeA, "1.2.3.4"),
			},
		},
		{
			title: "CNAME with resolve-target:true is replaced by A and AAAA records",
			lookupIP: func(string) ([]net.IP, error) {
				return []net.IP{net.ParseIP("1.2.3.4"), net.ParseIP("2001:db8::1")}, nil
			},
			endpoints: []*endpoint.Endpoint{
				func() *endpoint.Endpoint {
					ep := endpoint.NewEndpoint("test.example.internal", endpoint.RecordTypeCNAME, "lb.example.com")
					ep.WithProviderSpecific(resolveTargetPropertyName, "true")
					return ep
				}(),
			},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("test.example.internal", endpoint.RecordTypeA, "1.2.3.4"),
				endpoint.NewEndpoint("test.example.internal", endpoint.RecordTypeAAAA, "2001:db8::1"),
			},
		},
		{
			title:    "CNAME with resolve-target:true and unresolvable hostname is skipped",
			lookupIP: func(string) ([]net.IP, error) { return nil, errors.New("no such host") },
			endpoints: []*endpoint.Endpoint{
				func() *endpoint.Endpoint {
					ep := endpoint.NewEndpoint("test.example.internal", endpoint.RecordTypeCNAME, "lb.example.com")
					ep.WithProviderSpecific(resolveTargetPropertyName, "true")
					return ep
				}(),
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			title:    "CNAME without resolve-target annotation passes through unchanged",
			lookupIP: func(string) ([]net.IP, error) { return nil, errors.New("should not be called") },
			endpoints: []*endpoint.Endpoint{
				endpoint.NewEndpoint("test.example.internal", endpoint.RecordTypeCNAME, "lb.example.com"),
			},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("test.example.internal", endpoint.RecordTypeCNAME, "lb.example.com"),
			},
		},
		{
			title: "TTL is preserved on resolved endpoints",
			lookupIP: func(string) ([]net.IP, error) {
				return []net.IP{net.ParseIP("1.2.3.4")}, nil
			},
			endpoints: []*endpoint.Endpoint{
				func() *endpoint.Endpoint {
					ep := endpoint.NewEndpointWithTTL("test.example.internal", endpoint.RecordTypeCNAME, endpoint.TTL(300), "lb.example.com")
					ep.WithProviderSpecific(resolveTargetPropertyName, "true")
					return ep
				}(),
			},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpointWithTTL("test.example.internal", endpoint.RecordTypeA, endpoint.TTL(300), "1.2.3.4"),
			},
		},
		{
			title: "labels are preserved on resolved endpoints",
			lookupIP: func(string) ([]net.IP, error) {
				return []net.IP{net.ParseIP("1.2.3.4")}, nil
			},
			endpoints: []*endpoint.Endpoint{
				func() *endpoint.Endpoint {
					ep := endpoint.NewEndpoint("test.example.internal", endpoint.RecordTypeCNAME, "lb.example.com")
					ep.Labels = endpoint.Labels{"resource": "gateway/default/test"}
					ep.WithProviderSpecific(resolveTargetPropertyName, "true")
					return ep
				}(),
			},
			expected: []*endpoint.Endpoint{
				func() *endpoint.Endpoint {
					ep := endpoint.NewEndpoint("test.example.internal", endpoint.RecordTypeA, "1.2.3.4")
					ep.Labels = endpoint.Labels{"resource": "gateway/default/test"}
					return ep
				}(),
			},
		},
		{
			title: "SetIdentifier is preserved on resolved endpoints",
			lookupIP: func(string) ([]net.IP, error) {
				return []net.IP{net.ParseIP("1.2.3.4")}, nil
			},
			endpoints: []*endpoint.Endpoint{
				func() *endpoint.Endpoint {
					ep := endpoint.NewEndpoint("test.example.internal", endpoint.RecordTypeCNAME, "lb.example.com")
					ep.SetIdentifier = "weighted-routing-1"
					ep.WithProviderSpecific(resolveTargetPropertyName, "true")
					return ep
				}(),
			},
			expected: []*endpoint.Endpoint{
				func() *endpoint.Endpoint {
					ep := endpoint.NewEndpoint("test.example.internal", endpoint.RecordTypeA, "1.2.3.4")
					ep.SetIdentifier = "weighted-routing-1"
					return ep
				}(),
			},
		},
		{
			title: "ProviderSpecific properties other than resolve-target are preserved",
			lookupIP: func(string) ([]net.IP, error) {
				return []net.IP{net.ParseIP("1.2.3.4")}, nil
			},
			endpoints: []*endpoint.Endpoint{
				func() *endpoint.Endpoint {
					ep := endpoint.NewEndpoint("test.example.internal", endpoint.RecordTypeCNAME, "lb.example.com")
					ep.WithProviderSpecific("aws/weight", "100")
					ep.WithProviderSpecific("cloudflare-proxied", "true")
					ep.WithProviderSpecific(resolveTargetPropertyName, "true")
					return ep
				}(),
			},
			expected: []*endpoint.Endpoint{
				func() *endpoint.Endpoint {
					ep := endpoint.NewEndpoint("test.example.internal", endpoint.RecordTypeA, "1.2.3.4")
					ep.WithProviderSpecific("aws/weight", "100")
					ep.WithProviderSpecific("cloudflare-proxied", "true")
					return ep
				}(),
			},
		},
		{
			title: "multiple targets with partial resolution failure keeps successful ones",
			lookupIP: func(host string) ([]net.IP, error) {
				if host == "lb1.example.com" {
					return []net.IP{net.ParseIP("1.2.3.4")}, nil
				}
				return nil, errors.New("no such host")
			},
			endpoints: []*endpoint.Endpoint{
				func() *endpoint.Endpoint {
					ep := endpoint.NewEndpoint("test.example.internal", endpoint.RecordTypeCNAME, "lb1.example.com", "lb2.example.com")
					ep.WithProviderSpecific(resolveTargetPropertyName, "true")
					return ep
				}(),
			},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("test.example.internal", endpoint.RecordTypeA, "1.2.3.4"),
			},
		},
		{
			title: "multiple targets all resolved produces merged A and AAAA records",
			lookupIP: func(host string) ([]net.IP, error) {
				if host == "lb1.example.com" {
					return []net.IP{net.ParseIP("1.2.3.4"), net.ParseIP("2001:db8::1")}, nil
				}
				if host == "lb2.example.com" {
					return []net.IP{net.ParseIP("5.6.7.8"), net.ParseIP("2001:db8::2")}, nil
				}
				return nil, errors.New("no such host")
			},
			endpoints: []*endpoint.Endpoint{
				func() *endpoint.Endpoint {
					ep := endpoint.NewEndpoint("test.example.internal", endpoint.RecordTypeCNAME, "lb1.example.com", "lb2.example.com")
					ep.WithProviderSpecific(resolveTargetPropertyName, "true")
					return ep
				}(),
			},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("test.example.internal", endpoint.RecordTypeA, "1.2.3.4", "5.6.7.8"),
				endpoint.NewEndpoint("test.example.internal", endpoint.RecordTypeAAAA, "2001:db8::1", "2001:db8::2"),
			},
		},
		{
			title: "resolved targets are sorted before grouping for deterministic output",
			// lookupIP returns IPs in reverse/unsorted order to verify sorting is applied.
			lookupIP: func(string) ([]net.IP, error) {
				return []net.IP{
					net.ParseIP("2001:db8::2"),
					net.ParseIP("9.8.7.6"),
					net.ParseIP("1.2.3.4"),
					net.ParseIP("2001:db8::1"),
				}, nil
			},
			endpoints: []*endpoint.Endpoint{
				func() *endpoint.Endpoint {
					ep := endpoint.NewEndpoint("test.example.internal", endpoint.RecordTypeCNAME, "lb.example.com")
					ep.WithProviderSpecific(resolveTargetPropertyName, "true")
					return ep
				}(),
			},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("test.example.internal", endpoint.RecordTypeA, "1.2.3.4", "9.8.7.6"),
				endpoint.NewEndpoint("test.example.internal", endpoint.RecordTypeAAAA, "2001:db8::1", "2001:db8::2"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			ms := new(testutils.MockSource)
			ms.On("Endpoints").Return(tt.endpoints, nil)
			src := NewResolveTarget(ms, WithResolveTargetLookupIP(tt.lookupIP))

			got, err := src.Endpoints(t.Context())
			require.NoError(t, err)
			require.Len(t, got, len(tt.expected))
			for i, ep := range got {
				require.Equal(t, tt.expected[i].DNSName, ep.DNSName, "DNSName mismatch at index %d", i)
				require.Equal(t, tt.expected[i].RecordType, ep.RecordType, "RecordType mismatch at index %d", i)
				require.Equal(t, tt.expected[i].Targets, ep.Targets, "Targets mismatch at index %d", i)
				require.Equal(t, tt.expected[i].RecordTTL, ep.RecordTTL, "RecordTTL mismatch at index %d", i)
				require.Equal(t, tt.expected[i].Labels, ep.Labels, "Labels mismatch at index %d", i)
				require.Equal(t, tt.expected[i].SetIdentifier, ep.SetIdentifier, "SetIdentifier mismatch at index %d", i)
				require.ElementsMatch(t, tt.expected[i].ProviderSpecific, ep.ProviderSpecific, "ProviderSpecific mismatch at index %d", i)
			}
		})
	}
}

func TestResolveTargetEndpointsPerAnnotation(t *testing.T) {
	stubLookup := func(string) ([]net.IP, error) {
		return []net.IP{net.ParseIP("1.2.3.4")}, nil
	}
	tests := []struct {
		title           string
		annotationValue string // empty means annotation absent
		expectResolved  bool
	}{
		{
			title:           "annotation true → resolves",
			annotationValue: "true",
			expectResolved:  true,
		},
		{
			title:           "annotation false → keeps CNAME",
			annotationValue: "false",
			expectResolved:  false,
		},
		{
			title:          "no annotation → keeps CNAME",
			expectResolved: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			ep := endpoint.NewEndpoint("test.example.com", endpoint.RecordTypeCNAME, "lb.example.com")
			if tt.annotationValue != "" {
				ep.WithProviderSpecific(resolveTargetPropertyName, tt.annotationValue)
			}

			ms := new(testutils.MockSource)
			ms.On("Endpoints").Return([]*endpoint.Endpoint{ep}, nil)
			src := NewResolveTarget(ms, WithResolveTargetLookupIP(stubLookup))

			got, err := src.Endpoints(t.Context())
			require.NoError(t, err)

			if tt.expectResolved {
				require.Len(t, got, 1)
				require.Equal(t, endpoint.RecordTypeA, got[0].RecordType)
			} else {
				require.Len(t, got, 1)
				require.Equal(t, endpoint.RecordTypeCNAME, got[0].RecordType)
			}
			// resolve-target property must always be consumed
			_, ok := got[0].GetProviderSpecificProperty(resolveTargetPropertyName)
			require.False(t, ok, "resolve-target property should have been removed")
		})
	}
}

func TestResolveTarget_AddEventHandler(t *testing.T) {
	t.Run("should delegate to underlying source", func(t *testing.T) {
		mockSource := testutils.NewMockSource()

		src := NewResolveTarget(mockSource)
		src.AddEventHandler(t.Context(), func() {})

		mockSource.AssertNumberOfCalls(t, "AddEventHandler", 1)
	})
}

func TestResolveTarget_SourceError(t *testing.T) {
	t.Run("should propagate error from underlying source", func(t *testing.T) {
		expectedErr := errors.New("source error")
		ms := new(testutils.MockSource)
		ms.On("Endpoints").Return(nil, expectedErr)

		src := NewResolveTarget(ms)
		got, err := src.Endpoints(t.Context())

		require.Nil(t, got)
		require.ErrorIs(t, err, expectedErr)
	})
}

func TestResolveTarget_RefObjectIsPreserved(t *testing.T) {
	stubLookup := func(string) ([]net.IP, error) {
		return []net.IP{net.ParseIP("1.2.3.4")}, nil
	}

	ref := events.NewObjectReferenceFromParts("Ingress", "networking.k8s.io/v1", "default", "my-ingress", "", "")

	ep := endpoint.NewEndpoint("test.example.internal", endpoint.RecordTypeCNAME, "lb.example.com")
	ep.WithRefObject(ref)
	ep.WithProviderSpecific(resolveTargetPropertyName, "true")

	ms := new(testutils.MockSource)
	ms.On("Endpoints").Return([]*endpoint.Endpoint{ep}, nil)
	src := NewResolveTarget(ms, WithResolveTargetLookupIP(stubLookup))

	got, err := src.Endpoints(t.Context())
	require.NoError(t, err)
	require.Len(t, got, 1)
	require.Equal(t, endpoint.RecordTypeA, got[0].RecordType)
	require.Equal(t, ref, got[0].RefObject(), "RefObject should be preserved after resolution")
}

func TestLeakedPropertyShouldNotUpdate(t *testing.T) {
	current := endpoint.NewEndpoint("foo.example.com", endpoint.RecordTypeA, "1.2.3.4")

	desired := endpoint.NewEndpoint("foo.example.com", endpoint.RecordTypeA, "1.2.3.4")
	desired.WithProviderSpecific("resolve-target", "true")

	ms := new(testutils.MockSource)
	ms.On("Endpoints").Return([]*endpoint.Endpoint{desired}, nil)
	wrapped := NewResolveTarget(ms)

	desiredEndpoints, err := wrapped.Endpoints(t.Context())
	require.NoError(t, err)

	changes := (&plan.Plan{
		Current:        []*endpoint.Endpoint{current},
		Desired:        desiredEndpoints,
		ManagedRecords: []string{endpoint.RecordTypeA},
	}).Calculate().Changes

	assert.Empty(t, changes.Create, "no create expected")
	assert.Empty(t, changes.Delete, "no delete expected")
	// Correct: unchanged record must NOT be updated. Fails today due to the leak.
	assert.Empty(t, changes.UpdateNew, "unchanged record should not be updated")
}
