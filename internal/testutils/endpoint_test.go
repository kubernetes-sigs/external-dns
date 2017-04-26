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

	"github.com/kubernetes-incubator/external-dns/endpoint"
)

func ExampleSameEndpoints() {
	eps := []*endpoint.Endpoint{
		{
			DNSName: "example.org",
			Target:  "load-balancer.org",
		},
		{
			DNSName:    "example.org",
			Target:     "load-balancer.org",
			RecordType: "TXT",
		},
		{
			DNSName:    "abc.com",
			Target:     "something",
			RecordType: "TXT",
		},
		{
			DNSName:    "abc.com",
			Target:     "1.2.3.4",
			RecordType: "A",
		},
		{
			DNSName:    "bbc.com",
			Target:     "foo.com",
			RecordType: "CNAME",
		},
	}
	sort.Sort(byAllFields(eps))
	for _, ep := range eps {
		fmt.Println(ep)
	}
	// Output:
	// abc.com -> 1.2.3.4 (type "A")
	// abc.com -> something (type "TXT")
	// bbc.com -> foo.com (type "CNAME")
	// example.org -> load-balancer.org (type "")
	// example.org -> load-balancer.org (type "TXT")
}
