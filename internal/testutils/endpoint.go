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
	"github.com/kubernetes-incubator/external-dns/endpoint"
)

/** test utility functions for endpoints verifications */

// SameEndpoint returns true if two endpoint are same
// considers example.org. and example.org DNSName/Target as different endpoints
// TODO:might need reconsideration regarding trailing dot
func SameEndpoint(a, b endpoint.Endpoint) bool {
	return a.DNSName == b.DNSName && a.Target == b.Target
}

// SameSharedEndpoint returns true if two endpoint have same owner and endpoint fields
// considers example.org. and example.org DNSName/Target as different endpoints
// TODO:might need reconsideration regarding trailing dot
func SameSharedEndpoint(a, b endpoint.SharedEndpoint) bool {
	return a.Owner == b.Owner && SameEndpoint(a.Endpoint, b.Endpoint)
}

// SameEndpoints compares two slices of endpoints regardless of order
// [x,y,z] == [z,x,y]
// [x,x,z] == [x,z,x]
// [x,y,y] != [x,x,y]
// [x,x,x] != [x,x,z]
func SameEndpoints(a, b []endpoint.Endpoint) bool {
	if len(a) != len(b) {
		return false
	}

	calculator := map[string]map[string]uint8{} //testutils is not meant for large data sets
	for _, recordA := range a {
		if _, exists := calculator[recordA.DNSName]; !exists {
			calculator[recordA.DNSName] = map[string]uint8{}
		}
		if _, exists := calculator[recordA.DNSName][recordA.Target]; !exists {
			calculator[recordA.DNSName][recordA.Target] = 0
		}
		calculator[recordA.DNSName][recordA.Target]++
	}
	for _, recordB := range b {
		if _, exists := calculator[recordB.DNSName]; !exists {
			return false
		}
		if _, exists := calculator[recordB.DNSName][recordB.Target]; !exists {
			return false
		}
		calculator[recordB.DNSName][recordB.Target]--
	}

	for _, byDNSName := range calculator {
		for _, byCounter := range byDNSName {
			if byCounter != 0 {
				return false
			}
		}
	}

	return true
}

// SameSharedEndpoints compares two slices of endpoints regardless of order
// [x,y,z] == [z,x,y]
// [x,x,z] == [x,z,x]
// [x,y,y] != [x,x,y]
// [x,x,x] != [x,x,z]
func SameSharedEndpoints(a, b []*endpoint.SharedEndpoint) bool {
	if len(a) != len(b) {
		return false
	}
	calculator := map[string]map[string]map[string]uint8{} //testutils is not meant for large data sets
	for _, recordA := range a {
		if _, exists := calculator[recordA.Owner]; !exists {
			calculator[recordA.Owner] = map[string]map[string]uint8{}
		}
		if _, exists := calculator[recordA.Owner][recordA.DNSName]; !exists {
			calculator[recordA.Owner][recordA.DNSName] = map[string]uint8{}
		}
		if _, exists := calculator[recordA.Owner][recordA.DNSName][recordA.Target]; !exists {
			calculator[recordA.Owner][recordA.DNSName][recordA.Target] = 0
		}
		calculator[recordA.Owner][recordA.DNSName][recordA.Target]++
	}
	for _, recordB := range b {
		if _, exists := calculator[recordB.Owner]; !exists {
			return false
		}
		if _, exists := calculator[recordB.Owner][recordB.DNSName]; !exists {
			return false
		}
		if _, exists := calculator[recordB.Owner][recordB.DNSName][recordB.Target]; !exists {
			return false
		}
		calculator[recordB.Owner][recordB.DNSName][recordB.Target]--
	}
	for _, byOwner := range calculator {
		for _, byDNSName := range byOwner {
			for _, byCounter := range byDNSName {
				if byCounter != 0 {
					return false
				}
			}
		}
	}
	return true
}
