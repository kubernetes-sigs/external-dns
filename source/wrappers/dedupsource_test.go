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
	"context"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	"sigs.k8s.io/external-dns/source"
)

// Validates that dedupSource is a Source
var _ source.Source = &dedupSource{}

func TestDedup(t *testing.T) {
	t.Run("Endpoints", testDedupEndpoints)
}

// testDedupEndpoints tests that duplicates from the wrapped source are removed.
func testDedupEndpoints(t *testing.T) {
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
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"1.2.3.4"}},
			},
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"1.2.3.4"}},
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

			endpoints, err := source.Endpoints(context.Background())
			if err != nil {
				t.Fatal(err)
			}

			// Validate returned endpoints against desired endpoints.
			validateEndpoints(t, endpoints, tc.expected)

			// Validate that the mock source was called.
			mockSource.AssertExpectations(t)
		})
	}
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
				{DNSName: "example.org", RecordType: endpoint.RecordTypeMX, Targets: endpoint.Targets{"10 mail.example.org"}, ProviderSpecific: endpoint.ProviderSpecific{{Name: "alias", Value: "true"}}},
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			name: "A record with alias=true is kept",
			endpoints: []*endpoint.Endpoint{
				{DNSName: "example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.1.1"}, ProviderSpecific: endpoint.ProviderSpecific{{Name: "alias", Value: "true"}}},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.1.1"}, ProviderSpecific: endpoint.ProviderSpecific{{Name: "alias", Value: "true"}}},
			},
		},
		{
			name: "SRV record with alias=true is filtered out",
			endpoints: []*endpoint.Endpoint{
				{DNSName: "_sip._tcp.example.org", RecordType: endpoint.RecordTypeSRV, Targets: endpoint.Targets{"10 5 5060 sip.example.org."}, ProviderSpecific: endpoint.ProviderSpecific{{Name: "alias", Value: "true"}}},
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			name: "mixed valid and invalid TXT, A, AAAA records",
			endpoints: []*endpoint.Endpoint{
				{DNSName: "example.org", RecordType: endpoint.RecordTypeTXT, Targets: endpoint.Targets{"v=spf1 include:example.com ~all"}}, // valid
				{DNSName: "example.org", RecordType: endpoint.RecordTypeTXT, Targets: endpoint.Targets{""}},                                // invalid
				{DNSName: "example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.1.1"}},                       // valid
				{DNSName: "example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"not-an-ip"}},                         // invalid
				{DNSName: "example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:db8::1"}},                    // valid
				{DNSName: "example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"invalid-ipv6"}},                   // invalid
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "example.org", RecordType: endpoint.RecordTypeTXT, Targets: endpoint.Targets{"v=spf1 include:example.com ~all"}},
				{DNSName: "example.org", RecordType: endpoint.RecordTypeTXT, Targets: endpoint.Targets{""}},
				{DNSName: "example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.1.1"}},
				{DNSName: "example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"not-an-ip"}},
				{DNSName: "example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:db8::1"}},
				{DNSName: "example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"invalid-ipv6"}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSource := new(testutils.MockSource)
			mockSource.On("Endpoints").Return(tt.endpoints, nil)

			sr := NewDedupSource(mockSource)
			endpoints, err := sr.Endpoints(context.Background())
			require.NoError(t, err)

			validateEndpoints(t, endpoints, tt.expected)
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
				ProviderSpecific: endpoint.ProviderSpecific{{Name: "alias", Value: "true"}},
			},
			wantLogMsg: "Endpoint example.org of type MX does not support alias records",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hook := testutils.LogsUnderTestWithLogLevel(log.WarnLevel, t)

			mockSource := new(testutils.MockSource)
			mockSource.On("Endpoints").Return([]*endpoint.Endpoint{tt.endpoint}, nil)

			src := NewDedupSource(mockSource)
			_, err := src.Endpoints(context.Background())
			require.NoError(t, err)

			testutils.TestHelperLogContains(tt.wantLogMsg, hook, t)
		})
	}
}
