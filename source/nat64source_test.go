/*
Copyright 2024 The Kubernetes Authors.

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
	"context"
	"testing"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
)

// Validates that dedupSource is a Source
var _ Source = &nat64Source{}

func TestNAT64Source(t *testing.T) {
	t.Run("Endpoints", testNat64Source)
}

// testDedupEndpoints tests that duplicates from the wrapped source are removed.
func testNat64Source(t *testing.T) {
	for _, tc := range []struct {
		title     string
		endpoints []*endpoint.Endpoint
		expected  []*endpoint.Endpoint
	}{
		{
			"single non-nat64 ipv6 endpoint returns one ipv6 endpoint",
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:db8:1::1"}},
			},
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:db8:1::1"}},
			},
		},
		{
			"single nat64 ipv6 endpoint returns one ipv4 endpoint and one ipv6 endpoint",
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:db8::192.0.2.42"}},
			},
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:db8::192.0.2.42"}},
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.0.2.42"}},
			},
		},
		{
			"single nat64 ipv6 endpoint returns one ipv4 endpoint and one ipv6 endpoint",
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:db8::c000:22a"}},
			},
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:db8::c000:22a"}},
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.0.2.42"}},
			},
		},
	} {
		t.Run(tc.title, func(t *testing.T) {
			mockSource := new(testutils.MockSource)
			mockSource.On("Endpoints").Return(tc.endpoints, nil)

			// Create our object under test and get the endpoints.
			source := NewNAT64Source(mockSource, []string{"2001:DB8::/96"})

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
