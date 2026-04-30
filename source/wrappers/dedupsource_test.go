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

package wrappers

import (
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	v1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	logtest "sigs.k8s.io/external-dns/internal/testutils/log"
	"sigs.k8s.io/external-dns/source"
	"sigs.k8s.io/external-dns/source/types"
)

// Validates that dedupSource is a Source
var _ source.Source = &dedupSource{}

// TestDedupEndpoints tests that duplicates from the wrapped source are removed.
func TestDedupEndpoints(t *testing.T) {
	for _, tc := range []struct {
		title     string
		endpoints []*endpoint.Endpoint
		expected  []*endpoint.Endpoint
	}{
		{
			"one endpoint returns one endpoint",
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", Targets: endpoint.Targets{"1.2.3.4"}},
			},
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			"two different endpoints return two endpoints",
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", Targets: endpoint.Targets{"1.2.3.4"}},
				{DNSName: "bar.example.org", Targets: endpoint.Targets{"4.5.6.7"}},
			},
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", Targets: endpoint.Targets{"1.2.3.4"}},
				{DNSName: "bar.example.org", Targets: endpoint.Targets{"4.5.6.7"}},
			},
		},
		{
			"two endpoints with same dnsname and different targets return two endpoints",
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", Targets: endpoint.Targets{"1.2.3.4"}},
				{DNSName: "foo.example.org", Targets: endpoint.Targets{"4.5.6.7"}},
			},
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", Targets: endpoint.Targets{"1.2.3.4"}},
				{DNSName: "foo.example.org", Targets: endpoint.Targets{"4.5.6.7"}},
			},
		},
		{
			"two endpoints with different dnsname and same target return two endpoints",
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", Targets: endpoint.Targets{"1.2.3.4"}},
				{DNSName: "bar.example.org", Targets: endpoint.Targets{"1.2.3.4"}},
			},
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", Targets: endpoint.Targets{"1.2.3.4"}},
				{DNSName: "bar.example.org", Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			"two endpoints with same dnsname and same target return one endpoint",
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", Targets: endpoint.Targets{"1.2.3.4"}},
				{DNSName: "foo.example.org", Targets: endpoint.Targets{"1.2.3.4"}},
			},
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			"two endpoints with same dnsname, same type, and same target return one endpoint",
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
			},
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			"two endpoints with same dnsname, different record type, and same target return two endpoints",
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:db8::1"}},
			},
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:db8::1"}},
			},
		},
		{
			"two endpoints with same dnsname, one with record type, one without, and same target return two endpoints",
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
				{DNSName: "foo.example.org", Targets: endpoint.Targets{"1.2.3.4"}},
			},
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
				{DNSName: "foo.example.org", Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			"two endpoints with same dnsname, same type, same target but different SetIdentifier return two endpoints",
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}, SetIdentifier: "us-east-1"},
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}, SetIdentifier: "eu-west-1"},
			},
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}, SetIdentifier: "us-east-1"},
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}, SetIdentifier: "eu-west-1"},
			},
		},
		{
			"two endpoints with same dnsname, same type, same target and same SetIdentifier return one endpoint",
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}, SetIdentifier: "us-east-1"},
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}, SetIdentifier: "us-east-1"},
			},
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}, SetIdentifier: "us-east-1"},
			},
		},
		{
			"no endpoints returns empty endpoints",
			[]*endpoint.Endpoint{},
			[]*endpoint.Endpoint{},
		},
		{
			"one endpoint with multiple targets returns one endpoint and targets without duplicates",
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4", "34.66.66.77", "34.66.66.77"}},
			},
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4", "34.66.66.77"}},
			},
		},
	} {
		t.Run(tc.title, func(t *testing.T) {
			mockSource := new(testutils.MockSource)
			mockSource.On("Endpoints").Return(tc.endpoints, nil)

			// Create our object under test and get the endpoints.
			source := NewDedupSource(mockSource)

			endpoints, err := source.Endpoints(t.Context())
			if err != nil {
				t.Fatal(err)
			}

			// Validate returned endpoints against desired endpoints.
			testutils.ValidateEndpoints(t, endpoints, tc.expected)

			// Validate that the mock source was called.
			mockSource.AssertExpectations(t)
		})
	}

	t.Run("wrapped source error is propagated", func(t *testing.T) {
		mockSource := new(testutils.MockSource)
		mockSource.On("Endpoints").Return([]*endpoint.Endpoint(nil), assert.AnError)
		src := NewDedupSource(mockSource)
		_, err := src.Endpoints(t.Context())
		require.ErrorIs(t, err, assert.AnError)
	})

	t.Run("nil endpoint is skipped", func(t *testing.T) {
		mockSource := new(testutils.MockSource)
		mockSource.On("Endpoints").Return([]*endpoint.Endpoint{
			nil,
			{DNSName: "foo.example.org", Targets: endpoint.Targets{"1.2.3.4"}},
		}, nil)
		src := NewDedupSource(mockSource)
		endpoints, err := src.Endpoints(t.Context())
		require.NoError(t, err)
		testutils.ValidateEndpoints(t, endpoints, []*endpoint.Endpoint{
			{DNSName: "foo.example.org", Targets: endpoint.Targets{"1.2.3.4"}},
		})
	})
}

func TestDedupSource_AddEventHandler(t *testing.T) {
	tests := []struct {
		title string
		input []string
		times int
	}{
		{
			title: "should add event handler",
			times: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			mockSource := testutils.NewMockSource()

			src := NewDedupSource(mockSource)
			src.AddEventHandler(t.Context(), func() {})

			mockSource.AssertNumberOfCalls(t, "AddEventHandler", tt.times)
		})
	}
}

func TestDedupEndpointsValidation(t *testing.T) {
	tests := []struct {
		name      string
		endpoints []*endpoint.Endpoint
		expected  []*endpoint.Endpoint
	}{
		{
			name: "mix of SRV records",
			endpoints: []*endpoint.Endpoint{
				{DNSName: "_service._tcp.example.org", RecordType: endpoint.RecordTypeSRV, Targets: endpoint.Targets{"10 5 443 target.example.org."}}, // valid
				{DNSName: "_service._tcp.example.org", RecordType: endpoint.RecordTypeSRV, Targets: endpoint.Targets{"11 5 target.example.org"}},      // invalid
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "_service._tcp.example.org", RecordType: endpoint.RecordTypeSRV, Targets: endpoint.Targets{"10 5 443 target.example.org."}},
			},
		},
		{
			name: "invalid SRV record - missing priority",
			endpoints: []*endpoint.Endpoint{
				{DNSName: "_service._tcp.example.org", RecordType: endpoint.RecordTypeSRV, Targets: endpoint.Targets{"5 443 target.example.org"}},
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			name: "valid MX record",
			endpoints: []*endpoint.Endpoint{
				{DNSName: "example.org", RecordType: endpoint.RecordTypeMX, Targets: endpoint.Targets{"10 mail.example.org"}},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "example.org", RecordType: endpoint.RecordTypeMX, Targets: endpoint.Targets{"10 mail.example.org"}},
			},
		},
		{
			name: "invalid MX record - missing priority",
			endpoints: []*endpoint.Endpoint{
				{DNSName: "example.org", RecordType: endpoint.RecordTypeMX, Targets: endpoint.Targets{"mail.example.org"}},
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			name: "valid NAPTR record",
			endpoints: []*endpoint.Endpoint{
				{DNSName: "example.org", RecordType: endpoint.RecordTypeNAPTR, Targets: endpoint.Targets{"100 10 \"u\" \"E2U+sip\" \"!^.*$!sip:info@example.org!\" ."}},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "example.org", RecordType: endpoint.RecordTypeNAPTR, Targets: endpoint.Targets{"100 10 \"u\" \"E2U+sip\" \"!^.*$!sip:info@example.org!\" ."}},
			},
		},
		{
			name: "invalid NAPTR record - incomplete format",
			endpoints: []*endpoint.Endpoint{
				{DNSName: "example.org", RecordType: endpoint.RecordTypeNAPTR, Targets: endpoint.Targets{"100 10 \"u\""}}, // invalid
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "example.org", RecordType: endpoint.RecordTypeNAPTR, Targets: endpoint.Targets{"100 10 \"u\""}},
			},
		},
		{
			name: "mixed valid and invalid records",
			endpoints: []*endpoint.Endpoint{
				{DNSName: "_service._tcp.example.org", RecordType: endpoint.RecordTypeSRV, Targets: endpoint.Targets{"10 5 443"}}, // invalid
				{DNSName: "example.org", RecordType: endpoint.RecordTypeMX, Targets: endpoint.Targets{"mail.example.org"}},        // invalid
				{DNSName: "example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			name: "MX record with alias=true is filtered out",
			endpoints: []*endpoint.Endpoint{
				{DNSName: "example.org", RecordType: endpoint.RecordTypeMX, Targets: endpoint.Targets{"10 mail.example.org"}, ProviderSpecific: endpoint.ProviderSpecific{{Name: endpoint.ProviderSpecificAlias, Value: "true"}}},
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			name: "A record with alias=true is kept",
			endpoints: []*endpoint.Endpoint{
				{DNSName: "example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.1.1"}, ProviderSpecific: endpoint.ProviderSpecific{{Name: endpoint.ProviderSpecificAlias, Value: "true"}}},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.1.1"}, ProviderSpecific: endpoint.ProviderSpecific{{Name: endpoint.ProviderSpecificAlias, Value: "true"}}},
			},
		},
		{
			name: "SRV record with alias=true is filtered out",
			endpoints: []*endpoint.Endpoint{
				{DNSName: "_sip._tcp.example.org", RecordType: endpoint.RecordTypeSRV, Targets: endpoint.Targets{"10 5 5060 sip.example.org."}, ProviderSpecific: endpoint.ProviderSpecific{{Name: endpoint.ProviderSpecificAlias, Value: "true"}}},
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			name: "mixed valid and invalid TXT, A, AAAA records",
			endpoints: []*endpoint.Endpoint{
				{DNSName: "example.org", RecordType: endpoint.RecordTypeTXT, Targets: endpoint.Targets{"v=spf1 include:example.com ~all"}}, // valid
				{DNSName: "example.org", RecordType: endpoint.RecordTypeTXT, Targets: endpoint.Targets{""}},                                // valid (TXT allows empty)
				{DNSName: "example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.1.1"}},                       // valid
				{DNSName: "example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"not-an-ip"}},                         // invalid
				{DNSName: "example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:db8::1"}},                    // valid
				{DNSName: "example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"invalid-ipv6"}},                   // invalid
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "example.org", RecordType: endpoint.RecordTypeTXT, Targets: endpoint.Targets{"v=spf1 include:example.com ~all"}},
				{DNSName: "example.org", RecordType: endpoint.RecordTypeTXT, Targets: endpoint.Targets{""}},
				{DNSName: "example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.1.1"}},
				{DNSName: "example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:db8::1"}},
			},
		},
		{
			name: "valid PTR record with reverse DNS name",
			endpoints: []*endpoint.Endpoint{
				{DNSName: "2.49.168.192.in-addr.arpa", RecordType: endpoint.RecordTypePTR, Targets: endpoint.Targets{"web.example.com"}},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "2.49.168.192.in-addr.arpa", RecordType: endpoint.RecordTypePTR, Targets: endpoint.Targets{"web.example.com"}},
			},
		},
		{
			name: "invalid PTR record - non-reverse DNS name",
			endpoints: []*endpoint.Endpoint{
				{DNSName: "web.example.com", RecordType: endpoint.RecordTypePTR, Targets: endpoint.Targets{"other.example.com"}},
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			name: "invalid PTR record - target is an IP",
			endpoints: []*endpoint.Endpoint{
				{DNSName: "1.0.0.10.in-addr.arpa", RecordType: endpoint.RecordTypePTR, Targets: endpoint.Targets{"10.0.0.1"}},
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			name: "A record with record-type annotation passes through",
			endpoints: []*endpoint.Endpoint{
				{DNSName: "example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}, ProviderSpecific: endpoint.ProviderSpecific{{Name: endpoint.ProviderSpecificRecordType, Value: "PTR"}}},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}, ProviderSpecific: endpoint.ProviderSpecific{{Name: endpoint.ProviderSpecificRecordType, Value: "PTR"}}},
			},
		},
		{
			name: "duplicate A records with same record-type annotation are deduped",
			endpoints: []*endpoint.Endpoint{
				{DNSName: "example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}, ProviderSpecific: endpoint.ProviderSpecific{{Name: endpoint.ProviderSpecificRecordType, Value: "PTR"}}},
				{DNSName: "example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}, ProviderSpecific: endpoint.ProviderSpecific{{Name: endpoint.ProviderSpecificRecordType, Value: "PTR"}}},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}, ProviderSpecific: endpoint.ProviderSpecific{{Name: endpoint.ProviderSpecificRecordType, Value: "PTR"}}},
			},
		},
		{
			name: "A records with and without record-type annotation are deduped by identity key",
			endpoints: []*endpoint.Endpoint{
				{DNSName: "example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}, ProviderSpecific: endpoint.ProviderSpecific{{Name: endpoint.ProviderSpecificRecordType, Value: "PTR"}}},
				{DNSName: "example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}, ProviderSpecific: endpoint.ProviderSpecific{{Name: endpoint.ProviderSpecificRecordType, Value: "PTR"}}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSource := new(testutils.MockSource)
			mockSource.On("Endpoints").Return(tt.endpoints, nil)

			sr := NewDedupSource(mockSource)
			endpoints, err := sr.Endpoints(t.Context())
			require.NoError(t, err)

			testutils.ValidateEndpoints(t, endpoints, tt.expected)
			mockSource.AssertExpectations(t)
		})
	}
}

func TestDedupSource_WarnsOnInvalidEndpoint(t *testing.T) {
	tests := []struct {
		name       string
		endpoint   *endpoint.Endpoint
		wantLogMsg string
	}{
		{
			name: "invalid SRV record",
			endpoint: &endpoint.Endpoint{
				DNSName:       "example.org",
				RecordType:    endpoint.RecordTypeSRV,
				SetIdentifier: "default/svc/my-service",
				Targets:       endpoint.Targets{"10 mail.example.org"},
			},
			wantLogMsg: "Skipping endpoint [default/svc/my-service:example.org] due to invalid configuration [SRV:10 mail.example.org]",
		},
		{
			name: "unsupported alias on MX record",
			endpoint: &endpoint.Endpoint{
				DNSName:          "example.org",
				RecordType:       endpoint.RecordTypeMX,
				Targets:          endpoint.Targets{"10 mail.example.org"},
				ProviderSpecific: endpoint.ProviderSpecific{{Name: endpoint.ProviderSpecificAlias, Value: "true"}},
			},
			wantLogMsg: "Endpoint example.org of type MX does not support alias records",
		},
		{
			name: "invalid PTR record with non-reverse DNS name",
			endpoint: &endpoint.Endpoint{
				DNSName:    "web.example.org",
				RecordType: endpoint.RecordTypePTR,
				Targets:    endpoint.Targets{"other.example.org"},
			},
			wantLogMsg: "Skipping endpoint [:web.example.org] due to invalid configuration [PTR:other.example.org]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hook := logtest.LogsUnderTestWithLogLevel(log.WarnLevel, t)

			mockSource := new(testutils.MockSource)
			mockSource.On("Endpoints").Return([]*endpoint.Endpoint{tt.endpoint}, nil)

			src := NewDedupSource(mockSource)
			_, err := src.Endpoints(t.Context())
			require.NoError(t, err)

			logtest.TestHelperLogContains(tt.wantLogMsg, hook, t)
		})
	}
}

func TestDedupSource_RefObjects(t *testing.T) {
	tests := []struct {
		name     string
		input    func() []*endpoint.Endpoint
		expected func(*testing.T, []*endpoint.Endpoint)
	}{
		{
			name:  "empty input",
			input: func() []*endpoint.Endpoint { return []*endpoint.Endpoint{} },
			expected: func(t *testing.T, ep []*endpoint.Endpoint) {
				require.Empty(t, ep)
			},
		},
		{
			name: "single endpoint with RefObject preserved",
			input: func() []*endpoint.Endpoint {
				return []*endpoint.Endpoint{
					testutils.NewEndpointWithRef("example.com", "1.2.3.4", &v1.Service{
						ObjectMeta: metav1.ObjectMeta{Name: "foo", Namespace: "default", UID: "123"},
					}, types.Service),
				}
			},
			expected: func(t *testing.T, ep []*endpoint.Endpoint) {
				require.Len(t, ep, 1)
				require.NotNil(t, ep[0].RefObject())
				require.Equal(t, types.Service, ep[0].RefObject().Source)
				require.Equal(t, "foo", ep[0].RefObject().Name)
				require.Equal(t, "123", string(ep[0].RefObject().UID))
			},
		},
		{
			name: "duplicate endpoints with same source type - first RefObject preserved",
			input: func() []*endpoint.Endpoint {
				return []*endpoint.Endpoint{
					testutils.NewEndpointWithRef("example.com", "1.2.3.4", &v1.Service{
						ObjectMeta: metav1.ObjectMeta{Name: "first-svc", Namespace: "default", UID: "uid-first"},
					}, types.Service),
					testutils.NewEndpointWithRef("example.com", "1.2.3.4", &v1.Service{
						ObjectMeta: metav1.ObjectMeta{Name: "second-svc", Namespace: "other", UID: "uid-second"},
					}, types.Service),
				}
			},
			expected: func(t *testing.T, ep []*endpoint.Endpoint) {
				require.Len(t, ep, 1)
				require.NotNil(t, ep[0].RefObject())
				require.Equal(t, types.Service, ep[0].RefObject().Source)
				require.Equal(t, "first-svc", ep[0].RefObject().Name)
				require.Equal(t, "uid-first", string(ep[0].RefObject().UID))
			},
		},
		{
			name: "duplicate endpoints with different source types - first RefObject preserved",
			input: func() []*endpoint.Endpoint {
				return []*endpoint.Endpoint{
					testutils.NewEndpointWithRef("example.com", "1.2.3.4", &v1.Service{
						ObjectMeta: metav1.ObjectMeta{Name: "my-service", Namespace: "default", UID: "svc-uid"},
					}, types.Service),
					testutils.NewEndpointWithRef("example.com", "1.2.3.4", &networkingv1.Ingress{
						ObjectMeta: metav1.ObjectMeta{Name: "my-ingress", Namespace: "default", UID: "ing-uid"},
					}, types.Ingress),
				}
			},
			expected: func(t *testing.T, ep []*endpoint.Endpoint) {
				require.Len(t, ep, 1)
				require.NotNil(t, ep[0].RefObject())
				// First endpoint (Service) wins, Ingress is discarded
				require.Equal(t, types.Service, ep[0].RefObject().Source)
				require.Equal(t, "my-service", ep[0].RefObject().Name)
				require.Equal(t, "svc-uid", string(ep[0].RefObject().UID))
			},
		},
		{
			name: "duplicate endpoints - Ingress first, Service second - Ingress RefObject preserved",
			input: func() []*endpoint.Endpoint {
				return []*endpoint.Endpoint{
					testutils.NewEndpointWithRef("example.com", "1.2.3.4", &networkingv1.Ingress{
						ObjectMeta: metav1.ObjectMeta{Name: "my-ingress", Namespace: "default", UID: "ing-uid"},
					}, types.Ingress),
					testutils.NewEndpointWithRef("example.com", "1.2.3.4", &v1.Service{
						ObjectMeta: metav1.ObjectMeta{Name: "my-service", Namespace: "default", UID: "svc-uid"},
					}, types.Service),
				}
			},
			expected: func(t *testing.T, ep []*endpoint.Endpoint) {
				require.Len(t, ep, 1)
				require.NotNil(t, ep[0].RefObject())
				// First endpoint (Ingress) wins, Service is discarded
				require.Equal(t, types.Ingress, ep[0].RefObject().Source)
				require.Equal(t, "my-ingress", ep[0].RefObject().Name)
				require.Equal(t, "ing-uid", string(ep[0].RefObject().UID))
			},
		},
		{
			name: "non-duplicate endpoints with different source types - both RefObjects preserved",
			input: func() []*endpoint.Endpoint {
				return []*endpoint.Endpoint{
					testutils.NewEndpointWithRef("a.example.com", "1.1.1.1", &v1.Service{
						ObjectMeta: metav1.ObjectMeta{Name: "my-service", Namespace: "default", UID: "123"},
					}, types.Service),
					testutils.NewEndpointWithRef("b.example.com", "2.2.2.2", &networkingv1.Ingress{
						ObjectMeta: metav1.ObjectMeta{Name: "my-ingress", Namespace: "default", UID: "234"},
					}, types.Ingress),
				}
			},
			expected: func(t *testing.T, ep []*endpoint.Endpoint) {
				require.Len(t, ep, 2)

				// Find endpoints by DNS name since order may vary
				var svcEndpoint, ingEndpoint *endpoint.Endpoint
				for _, e := range ep {
					switch e.DNSName {
					case "a.example.com":
						svcEndpoint = e
					case "b.example.com":
						ingEndpoint = e
					}
				}

				require.NotNil(t, svcEndpoint)
				require.NotNil(t, svcEndpoint.RefObject())
				require.Equal(t, types.Service, svcEndpoint.RefObject().Source)
				require.Equal(t, "my-service", svcEndpoint.RefObject().Name)

				require.NotNil(t, ingEndpoint)
				require.NotNil(t, ingEndpoint.RefObject())
				require.Equal(t, types.Ingress, ingEndpoint.RefObject().Source)
				require.Equal(t, "my-ingress", ingEndpoint.RefObject().Name)
			},
		},
		{
			name: "three duplicate endpoints from different sources - first RefObject preserved",
			input: func() []*endpoint.Endpoint {
				return []*endpoint.Endpoint{
					testutils.NewEndpointWithRef("example.com", "1.2.3.4", &v1.Service{
						ObjectMeta: metav1.ObjectMeta{Name: "my-service", Namespace: "default", UID: "123"},
					}, types.Service),
					testutils.NewEndpointWithRef("example.com", "1.2.3.4", &networkingv1.Ingress{
						ObjectMeta: metav1.ObjectMeta{Name: "my-ingress", Namespace: "default", UID: "345"},
					}, types.Ingress),
					testutils.NewEndpointWithRef("example.com", "1.2.3.4", &v1.Pod{
						ObjectMeta: metav1.ObjectMeta{Name: "my-pod", Namespace: "default", UID: "456"},
					}, types.Pod),
				}
			},
			expected: func(t *testing.T, ep []*endpoint.Endpoint) {
				require.Len(t, ep, 1)
				require.NotNil(t, ep[0].RefObject())
				// First endpoint (Service) wins
				require.Equal(t, types.Service, ep[0].RefObject().Source)
				require.Equal(t, "my-service", ep[0].RefObject().Name)
				require.Equal(t, "123", string(ep[0].RefObject().UID))
			},
		},
		{
			name: "duplicate endpoints with one having nil RefObject - first RefObject preserved",
			input: func() []*endpoint.Endpoint {
				return []*endpoint.Endpoint{
					testutils.NewEndpointWithRef("example.com", "1.2.3.4", &v1.Service{
						ObjectMeta: metav1.ObjectMeta{Name: "my-service", Namespace: "default", UID: "123"},
					}, types.Service),
					endpoint.NewEndpoint("example.com", endpoint.RecordTypeA, "1.2.3.4"),
				}
			},
			expected: func(t *testing.T, ep []*endpoint.Endpoint) {
				require.Len(t, ep, 1)
				require.NotNil(t, ep[0].RefObject())
				require.Equal(t, types.Service, ep[0].RefObject().Source)
				require.Equal(t, "123", string(ep[0].RefObject().UID))
			},
		},
		{
			name: "duplicate endpoints with first having nil RefObject - nil preserved",
			input: func() []*endpoint.Endpoint {
				return []*endpoint.Endpoint{
					endpoint.NewEndpoint("example.com", endpoint.RecordTypeA, "1.2.3.4"),
					testutils.NewEndpointWithRef("example.com", "1.2.3.4", &v1.Service{
						ObjectMeta: metav1.ObjectMeta{Name: "my-service", Namespace: "default", UID: "345"},
					}, types.Service),
				}
			},
			expected: func(t *testing.T, ep []*endpoint.Endpoint) {
				require.Len(t, ep, 1)
				// First endpoint (without RefObject) wins
				require.Nil(t, ep[0].RefObject())
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSource := new(testutils.MockSource)
			mockSource.On("Endpoints").Return(tt.input(), nil)

			src := NewDedupSource(mockSource)
			endpoints, err := src.Endpoints(t.Context())
			require.NoError(t, err)

			tt.expected(t, endpoints)
			mockSource.AssertExpectations(t)
		})
	}
}

func TestDedupSource_DeduplicatedEndpointsMetric(t *testing.T) {
	deduplicatedEndpoints.Reset()

	eps := []*endpoint.Endpoint{
		endpoint.NewEndpoint("web.example.com", endpoint.RecordTypeA, "192.168.1.1"),
		endpoint.NewEndpoint("web.example.com", endpoint.RecordTypeA, "192.168.1.1"), // duplicate
		endpoint.NewEndpoint("api.example.com", endpoint.RecordTypeAAAA, "2001:db8::1"),
		endpoint.NewEndpoint("api.example.com", endpoint.RecordTypeAAAA, "2001:db8::1"), // duplicate
		endpoint.NewEndpoint("api.example.com", endpoint.RecordTypeAAAA, "2001:db8::1"), // another duplicate
	}

	mockSource := testutils.NewMockSource(eps...)
	src := NewDedupSource(mockSource)
	result, err := src.Endpoints(t.Context())
	require.NoError(t, err)
	require.Len(t, result, 2) // only unique endpoints

	testutils.TestHelperVerifyMetricsGaugeVectorWithLabels(
		t, 1.0, deduplicatedEndpoints.Gauge,
		map[string]string{"record_type": "a", "source_type": "unknown"},
	)
	testutils.TestHelperVerifyMetricsGaugeVectorWithLabels(
		t, 2.0, deduplicatedEndpoints.Gauge,
		map[string]string{"record_type": "aaaa", "source_type": "unknown"},
	)
}

func TestDedupSource_InvalidEndpointsMetric(t *testing.T) {
	invalidEndpoints.Reset()

	eps := []*endpoint.Endpoint{
		// valid A record
		endpoint.NewEndpoint("web.example.com", endpoint.RecordTypeA, "192.168.1.1"),
		// invalid SRV record (missing port and target host)
		{DNSName: "_svc._tcp.example.org", RecordType: endpoint.RecordTypeSRV, Targets: endpoint.Targets{"10 mail.example.org"}},
	}

	mockSource := testutils.NewMockSource(eps...)
	src := NewDedupSource(mockSource)
	result, err := src.Endpoints(t.Context())
	require.NoError(t, err)
	require.Len(t, result, 1) // only the valid A record

	testutils.TestHelperVerifyMetricsGaugeVectorWithLabels(
		t, 1.0, invalidEndpoints.Gauge,
		map[string]string{"record_type": "srv", "source_type": "unknown"},
	)
}
