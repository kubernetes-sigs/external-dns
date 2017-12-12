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

	"github.com/stretchr/testify/assert"
)

func ExampleSameEndpoints() {
	eps := []*endpoint.Endpoint{
		{
			DNSName: "example.org",
			Targets: []string{"load-balancer.org"},
		},
		{
			DNSName: "example2.org",
			Targets: []string{"load-balancer.org", "load-balancer.com"},
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
		{
			DNSName:    "cbc.com",
			Targets:    []string{"foo.com"},
			RecordType: "CNAME",
			RecordTTL:  endpoint.TTL(60),
		},
	}
	sort.Slice(eps, func(i, j int) bool { return eps[i].String() < eps[j].String() })
	for _, ep := range eps {
		fmt.Println(ep)
	}
	// Output:
	// abc.com 0 IN A [1.2.3.4]
	// abc.com 0 IN TXT [something]
	// bbc.com 0 IN CNAME [foo.com]
	// cbc.com 60 IN CNAME [foo.com]
	// example.org 0 IN  [load-balancer.org]
	// example.org 0 IN TXT [load-balancer.org]
	// example2.org 0 IN  [load-balancer.org load-balancer.com]
}

func TestSameEndpointDifferentHostnames(t *testing.T) {
	// a blank endpoint
	blank := &endpoint.Endpoint{}
	// a foo endpoint
	foo := endpoint.NewEndpoint("foo", []string{"8.8.8.8"}, "A")
	// a bar endpoint
	bar := endpoint.NewEndpoint("bar", []string{"8.8.8.8"}, "A")

	for _, tc := range []struct {
		endpoint  *endpoint.Endpoint
		candidate *endpoint.Endpoint
		same      bool
	}{
		// Two blank endpoints are equal
		{blank, blank, true},
		// Two endpoints with the same hostname are equal
		{foo, foo, true},
		// A blank endpoint differs from an endpoint with content
		{blank, foo, false},
		// Two endpoints with a different hostname are different
		{foo, bar, false},
	} {
		assert.Equal(t, tc.same, SameEndpoint(tc.endpoint, tc.candidate))
	}
}

func TestSameEndpointDifferentTypes(t *testing.T) {
	// a blank endpoint
	blank := &endpoint.Endpoint{}
	// a foo endpoint
	foo := endpoint.NewEndpoint("foo", []string{"8.8.8.8"}, "A")
	// a bar endpoint
	bar := endpoint.NewEndpoint("foo", []string{"8.8.8.8"}, "CNAME")

	for _, tc := range []struct {
		endpoint  *endpoint.Endpoint
		candidate *endpoint.Endpoint
		same      bool
	}{
		// Two blank endpoints are equal
		{blank, blank, true},
		// Two endpoints with the same types are equal
		{foo, foo, true},
		// A blank endpoint differs from an endpoint with a type
		{blank, foo, false},
		// Two endpoints with different types are different
		{foo, bar, false},
	} {
		assert.Equal(t, tc.same, SameEndpoint(tc.endpoint, tc.candidate))
	}
}

func TestSameEndpointDifferentTargets(t *testing.T) {
	// a blank endpoint
	blank := &endpoint.Endpoint{}
	// a single-target endpoint
	single := endpoint.NewEndpoint("foo", []string{"8.8.4.4"}, "A")
	// another single-target endpoint
	singleV2 := endpoint.NewEndpoint("foo", []string{"8.8.8.8"}, "A")
	// a multiple-target endpoint
	multiple := endpoint.NewEndpoint("foo", []string{"8.8.4.4", "8.8.8.8"}, "A")
	// a multiple-target endpoint but in reverse order
	multipleReverse := endpoint.NewEndpoint("foo", []string{"8.8.8.8", "8.8.4.4"}, "A")
	// another multiple-target endpoint
	multipleV2 := endpoint.NewEndpoint("foo", []string{"8.8.4.4", "1.2.3.4"}, "A")

	for _, tc := range []struct {
		endpoint  *endpoint.Endpoint
		candidate *endpoint.Endpoint
		same      bool
	}{
		// Two blank endpoints are equal
		{blank, blank, true},
		// The same single-target endpoint is equal
		{single, single, true},
		// The same multiple-target endpoint is equal
		{multiple, multiple, true},
		// The same multiple-target endpoint with different order is equal
		{multiple, multipleReverse, true},
		// A blank endpoint differs from an endpoint with a target
		{blank, single, false},
		// Two endpoints differ if their single target differs
		{single, singleV2, false},
		// Two endpoints differ if one of their targets differ
		{multiple, multipleV2, false},
		// Two endpoints always differ if their number of targets differ
		{single, multiple, false},
		// The order of comparison shouldn't lead to a panic
		{multiple, single, false},
	} {
		assert.Equal(t, tc.same, SameEndpoint(tc.endpoint, tc.candidate))
	}
}

func TestSameEndpointDifferentLabels(t *testing.T) {
	// a blank endpoint
	blank := &endpoint.Endpoint{}
	// a single-label endpoint
	single := endpoint.NewEndpointWithLabels("foo", []string{"8.8.8.8"}, "A", map[string]string{"foo": "bar"})
	// another single-target endpoint
	singleV2 := endpoint.NewEndpointWithLabels("foo", []string{"8.8.8.8"}, "A", map[string]string{"foo": "baz"})
	// a multiple-target endpoint
	multiple := endpoint.NewEndpointWithLabels("foo", []string{"8.8.8.8"}, "A", map[string]string{"foo": "bar", "qux": "wambo"})
	// another multiple-target endpoint with a different label value
	multipleV2 := endpoint.NewEndpointWithLabels("foo", []string{"8.8.8.8"}, "A", map[string]string{"foo": "bar", "qux": "rambo"})
	// another multiple-target endpoint with a different label key
	multipleV3 := endpoint.NewEndpointWithLabels("foo", []string{"8.8.8.8"}, "A", map[string]string{"foo": "bar", "quux": "wambo"})

	for _, tc := range []struct {
		endpoint  *endpoint.Endpoint
		candidate *endpoint.Endpoint
		same      bool
	}{
		// Two blank endpoints are equal
		{blank, blank, true},
		// The same single-label endpoint is equal
		{single, single, true},
		// The same multiple-label endpoint is equal
		{multiple, multiple, true},
		// A blank endpoint differs from an endpoint with a label
		{blank, single, false},
		// Two endpoints differ if their single label differs
		{single, singleV2, false},
		// Two endpoints differ if one of their labels differs
		{multiple, multipleV2, false},
		// Two endpoints differ if one of their labels differs
		{multiple, multipleV3, false},
		// Two endpoints always differ if their number of targets differ
		{single, multiple, false},
		// The order of comparison shouldn't lead to a panic
		{multiple, single, false},
	} {
		assert.Equal(t, tc.same, SameEndpoint(tc.endpoint, tc.candidate))
	}
}

func TestSameEndpoints(t *testing.T) {
	// two simple endpoints
	foo := endpoint.NewEndpoint("foo", []string{"8.8.8.8"}, "A")
	bar := endpoint.NewEndpoint("bar", []string{"8.8.4.4"}, "A")
	baz := endpoint.NewEndpoint("baz", []string{"1.2.3.4"}, "A")

	// an empty list of endpoints
	empty := []*endpoint.Endpoint{}
	// a list containing a single endpoint
	single := []*endpoint.Endpoint{foo}
	// a list containing another single endpoint
	singleV2 := []*endpoint.Endpoint{bar}
	// a list containing multiple endpoints
	multiple := []*endpoint.Endpoint{foo, bar}
	// a list containing multiple endpoints in reverse order
	multipleReverse := []*endpoint.Endpoint{bar, foo}
	// a list containing other multiple endpoints
	multipleV2 := []*endpoint.Endpoint{foo, baz}

	for _, tc := range []struct {
		endpoints  []*endpoint.Endpoint
		candidates []*endpoint.Endpoint
		same       bool
	}{
		// Two empty lists of endpoints are equal
		{empty, empty, true},
		// Two single-item list with the same item are equal
		{single, single, true},
		// Two multiple-item lists with the same items are equal
		{multiple, multiple, true},
		// The same multiple-item lists with different order are equal
		{multiple, multipleReverse, true},
		// An empty list of endpoints is different to a non-empty list
		{empty, single, false},
		// Two single-item list with different items are different
		{single, singleV2, false},
		// Two multiple-item lists with different items are different
		{multiple, multipleV2, false},
		// Two list always differ if their number of elements differ
		{single, multiple, false},
		// The order of comparison shouldn't lead to a panic
		{multiple, single, false},
	} {
		assert.Equal(t, tc.same, SameEndpoints(tc.endpoints, tc.candidates))
	}
}

// We need this?
// eps := []*endpoint.Endpoint{
// 	{
// 		DNSName: "example.org",
// 		Targets: []string{"load-balancer.org"},
// 	},
// 	{
// 		DNSName:    "example.org",
// 		Targets:    []string{"load-balancer.org"},
// 		RecordType: endpoint.RecordTypeTXT,
// 	},
// 	{
// 		DNSName:    "abc.com",
// 		Targets:    []string{"something"},
// 		RecordType: endpoint.RecordTypeTXT,
// 	},
// 	{
// 		DNSName:    "abc.com",
// 		Targets:    []string{"1.2.3.4"},
// 		RecordType: endpoint.RecordTypeA,
// 	},
// 	{
// 		DNSName:    "bbc.com",
// 		Targets:    []string{"foo.com"},
// 		RecordType: endpoint.RecordTypeCNAME,
// 	},
// 	{
// 		DNSName:    "cbc.com",
// 		Targets:    []string{"foo.com"},
// 		RecordType: "CNAME",
// 		RecordTTL:  endpoint.TTL(60),
// 	},
//

// func ExampleSameEndpoints() {
// 	eps := []*endpoint.Endpoint{
// 		{
// 			DNSName: "example.org",
// 			Targets: []string{"load-balancer.org", "load-balancer-2.org"},
// 		},
// 		{
// 			DNSName:    "example.org",
// 			Targets:    []string{"load-balancer.org", "load-balancer-2.org"},
// 			RecordType: "TXT",
// 		},
// 		{
// 			DNSName:    "abc.com",
// 			Targets:    []string{"something", "else"},
// 			RecordType: "TXT",
// 		},
// 		{
// 			DNSName:    "abc.com",
// 			Targets:    []string{"1.2.3.4", "8.8.8.8"},
// 			RecordType: "A",
// 		},
// 		{
// 			DNSName:    "bbc.com",
// 			Targets:    []string{"foo.com", "bar.com"},
// 			RecordType: "CNAME",
// 		},
// 	}
// 	sort.Sort(byAllFields(eps))
// 	for _, ep := range eps {
// 		fmt.Println(ep)
// 	}
// 	// Output:
// 	// abc.com -> [something else] (type "TXT")
// 	// abc.com -> [1.2.3.4 8.8.8.8] (type "A")
// 	// bbc.com -> [foo.com bar.com] (type "CNAME")
// 	// example.org -> [load-balancer.org load-balancer-2.org] (type "")
// 	// example.org -> [load-balancer.org load-balancer-2.org] (type "TXT")
// }
