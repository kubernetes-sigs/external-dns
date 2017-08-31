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

package testutils

import (
	"fmt"
	"sort"
	"testing"

	"github.com/kubernetes-incubator/external-dns/endpoint"
)

func ExampleSameEndpoints() {
	eps := []*endpoint.Endpoint{
		{
			DNSName:    "example.org",
			Targets:    []string{"load-balancer.org"},
			RecordType: endpoint.RecordTypeCNAME,
		},
		{
			DNSName:    "example.org",
			Targets:    []string{"load-balancer.org", "load-balancer.gov"},
			RecordType: endpoint.RecordTypeCNAME,
		},
		{
			DNSName:    "example.org",
			Targets:    []string{"load-balancer.org", "load-balancer.gov, load-balancer.net"},
			RecordType: endpoint.RecordTypeCNAME,
		},
		{
			DNSName:    "example.org",
			Targets:    []string{"load-balancer.org"},
			RecordType: endpoint.RecordTypeTXT,
		},
		{
			DNSName:    "abc.com",
			Targets:    []string{"something"},
			RecordType: endpoint.RecordTypeTXT,
		},
		{
			DNSName:    "abc.com",
			Targets:    []string{"1.2.3.4"},
			RecordType: endpoint.RecordTypeA,
		},
		{
			DNSName:    "bbc.com",
			Targets:    []string{"foo.com"},
			RecordType: endpoint.RecordTypeCNAME,
		},
	}
	sort.Sort(byAllFields(eps))
	for _, ep := range eps {
		fmt.Println(ep)
	}
	// Output:
	// abc.com -> [1.2.3.4] (type "A")
	// abc.com -> [something] (type "TXT")
	// bbc.com -> [foo.com] (type "CNAME")
	// example.org -> [load-balancer.org] (type "CNAME")
	// example.org -> [load-balancer.org] (type "TXT")
	// example.org -> [load-balancer.gov load-balancer.org] (type "CNAME")
	// example.org -> [load-balancer.gov, load-balancer.net load-balancer.org] (type "CNAME")

}

func TestSameEndpoints(t *testing.T) {
	eps1 := []*endpoint.Endpoint{
		{
			DNSName:    "example.gov",
			Targets:    []string{"load-balancer.net", "load-balancer.org"},
			RecordType: endpoint.RecordTypeCNAME,
		},
		{
			DNSName:    "example.org",
			Targets:    []string{"load-balancer.org", "load-balancer.gov", "load-balancer.god"},
			RecordType: endpoint.RecordTypeCNAME,
		},
		{
			DNSName:    "example.gov",
			Targets:    []string{"load-balancer.net"},
			RecordType: endpoint.RecordTypeCNAME,
		},
		{
			DNSName:    "zyx.org",
			Targets:    []string{"load-balancer.org", "load-balancer.gov", "load-balancer.net"},
			RecordType: endpoint.RecordTypeCNAME,
		},
	}
	eps2 := []*endpoint.Endpoint{
		{
			DNSName:    "example.gov",
			Targets:    []string{"load-balancer.org", "load-balancer.net"},
			RecordType: endpoint.RecordTypeCNAME,
		},
		{
			DNSName:    "example.gov",
			Targets:    []string{"load-balancer.net"},
			RecordType: endpoint.RecordTypeCNAME,
		},
		{
			DNSName:    "example.org",
			Targets:    []string{"load-balancer.org", "load-balancer.gov", "load-balancer.god"},
			RecordType: endpoint.RecordTypeCNAME,
		},
		{
			DNSName:    "zyx.org",
			Targets:    []string{"load-balancer.net", "load-balancer.org", "load-balancer.gov"},
			RecordType: endpoint.RecordTypeCNAME,
		},
	}

	if !SameEndpoints(eps1, eps2) {
		t.Errorf("endpoint (%v) does not match endpoint (%v)", eps1, eps2)
	}
}
