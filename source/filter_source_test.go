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
	"net"
	"testing"

	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/internal/testutils"
)

// Validates that filterSource is a Source
var _ Source = &filterSource{}

func TestFilter(t *testing.T) {
	t.Run("Endpoints", testFilterEndpoints)
}

// testFilterEndpoints tests that filtered IPs from the wrapped source are removed.
func testFilterEndpoints(t *testing.T) {
	for _, tc := range []struct {
		title     string
		endpoints []*endpoint.Endpoint
		expected  []*endpoint.Endpoint
	}{
		{
			"one endpoint returns one endpoint",
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", Targets: endpoint.NewTargets("1.2.3.4"), RecordType: endpoint.RecordTypeA},
			},
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", Targets: endpoint.NewTargets("1.2.3.4"), RecordType: endpoint.RecordTypeA},
			},
		},
		{
			"two different endpoints return two endpoints",
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", Targets: endpoint.NewTargets("1.2.3.4"), RecordType: endpoint.RecordTypeA},
				{DNSName: "bar.example.org", Targets: endpoint.NewTargets("4.5.6.7"), RecordType: endpoint.RecordTypeA},
			},
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", Targets: endpoint.NewTargets("1.2.3.4"), RecordType: endpoint.RecordTypeA},
				{DNSName: "bar.example.org", Targets: endpoint.NewTargets("4.5.6.7"), RecordType: endpoint.RecordTypeA},
			},
		},
		{
			"non A-records ignores",
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", Targets: endpoint.NewTargets("1.2.3.4"), RecordType: endpoint.RecordTypeA},
				{DNSName: "foo.example.org", Targets: endpoint.NewTargets("192.168.100.10"), RecordType: endpoint.RecordTypeCNAME},
			},
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", Targets: endpoint.NewTargets("1.2.3.4"), RecordType: endpoint.RecordTypeA},
				{DNSName: "foo.example.org", Targets: endpoint.NewTargets("192.168.100.10"), RecordType: endpoint.RecordTypeCNAME},
			},
		},
		{
			"two endpoints with same dnsname and same target return one endpoint",
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", Targets: endpoint.NewTargets("1.2.3.4"), RecordType: endpoint.RecordTypeA},
				{DNSName: "foo.example.org", Targets: endpoint.NewTargets("192.168.100.10"), RecordType: endpoint.RecordTypeA},
			},
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", Targets: endpoint.NewTargets("1.2.3.4"), RecordType: endpoint.RecordTypeA},
			},
		},
	} {
		t.Run(tc.title, func(t *testing.T) {
			mockSource := new(testutils.MockSource)
			mockSource.On("Endpoints").Return(tc.endpoints, nil)

			// Create our object under test and get the endpoints.
			_, cidr, _ := net.ParseCIDR("192.168.100.0/24")
			source := NewFilterSource([]*net.IPNet{cidr}, mockSource)

			endpoints, err := source.Endpoints()
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
