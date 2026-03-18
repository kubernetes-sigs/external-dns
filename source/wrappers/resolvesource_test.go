/*
Copyright 2025 The Kubernetes Authors.

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

	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
)

// withLookupIP is a test-only resolveSourceOption that replaces the DNS lookup function.
func withLookupIP(fn func(string) ([]net.IP, error)) resolveSourceOption {
	return func(rs *resolveSource) {
		rs.lookupIP = fn
	}
}

func TestResolveSourceEndpointsGlobalFlag(t *testing.T) {
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
			title: "CNAME with resolvable hostname is replaced by A and AAAA records",
			lookupIP: func(string) ([]net.IP, error) {
				return []net.IP{net.ParseIP("1.2.3.4"), net.ParseIP("2001:db8::1")}, nil
			},
			endpoints: []*endpoint.Endpoint{
				endpoint.NewEndpoint("test.example.internal", endpoint.RecordTypeCNAME, "lb.example.com"),
			},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("test.example.internal", endpoint.RecordTypeA, "1.2.3.4"),
				endpoint.NewEndpoint("test.example.internal", endpoint.RecordTypeAAAA, "2001:db8::1"),
			},
		},
		{
			title:    "CNAME with unresolvable hostname is skipped",
			lookupIP: func(string) ([]net.IP, error) { return nil, errors.New("no such host") },
			endpoints: []*endpoint.Endpoint{
				endpoint.NewEndpoint("test.example.internal", endpoint.RecordTypeCNAME, "lb.example.com"),
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			title: "mixed IP and CNAME endpoints: CNAMEs are resolved, IP endpoints pass through",
			lookupIP: func(host string) ([]net.IP, error) {
				if host == "lb.example.com" {
					return []net.IP{net.ParseIP("10.0.0.1")}, nil
				}
				return nil, errors.New("no such host")
			},
			endpoints: []*endpoint.Endpoint{
				endpoint.NewEndpoint("a.example.com", endpoint.RecordTypeA, "1.2.3.4"),
				endpoint.NewEndpoint("aaaa.example.com", endpoint.RecordTypeAAAA, "2001:db8::1"),
				endpoint.NewEndpoint("cname.example.com", endpoint.RecordTypeCNAME, "lb.example.com"),
			},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("a.example.com", endpoint.RecordTypeA, "1.2.3.4"),
				endpoint.NewEndpoint("aaaa.example.com", endpoint.RecordTypeAAAA, "2001:db8::1"),
				endpoint.NewEndpoint("cname.example.com", endpoint.RecordTypeA, "10.0.0.1"),
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
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			ms := new(testutils.MockSource)
			ms.On("Endpoints").Return(tt.endpoints, nil)
			src := NewResolveSource(ms, true, withLookupIP(tt.lookupIP))

			got, err := src.Endpoints(t.Context())
			require.NoError(t, err)
			validateEndpoints(t, got, tt.expected)
		})
	}
}

func TestResolveSourceEndpointsPerAnnotation(t *testing.T) {
	stubLookup := func(string) ([]net.IP, error) {
		return []net.IP{net.ParseIP("1.2.3.4")}, nil
	}
	tests := []struct {
		title           string
		globalFlag      bool
		annotationValue string // empty means annotation absent
		expectResolved  bool
	}{
		{
			title:           "annotation true, global flag off → resolves",
			globalFlag:      false,
			annotationValue: "true",
			expectResolved:  true,
		},
		{
			title:           "annotation false, global flag on → keeps CNAME",
			globalFlag:      true,
			annotationValue: "false",
			expectResolved:  false,
		},
		{
			title:          "no annotation, global flag on → resolves",
			globalFlag:     true,
			expectResolved: true,
		},
		{
			title:          "no annotation, global flag off → keeps CNAME",
			globalFlag:     false,
			expectResolved: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			ep := endpoint.NewEndpoint("test.example.com", endpoint.RecordTypeCNAME, "lb.example.com")
			if tt.annotationValue != "" {
				ep.WithProviderSpecific("resolve-target", tt.annotationValue)
			}

			ms := new(testutils.MockSource)
			ms.On("Endpoints").Return([]*endpoint.Endpoint{ep}, nil)
			src := NewResolveSource(ms, tt.globalFlag, withLookupIP(stubLookup))

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
			_, ok := got[0].GetProviderSpecificProperty("resolve-target")
			require.False(t, ok, "resolve-target property should have been removed")
		})
	}
}
