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
	"sort"

	"github.com/kubernetes-incubator/external-dns/endpoint"
)

// SameEndpoint returns true if two endpoints are same
// considers example.org. and example.org DNSName/Target as different endpoints
func SameEndpoint(a, b *endpoint.Endpoint) bool {
	if a.DNSName != b.DNSName {
		return false
	}

	if a.RecordType != b.RecordType {
		return false
	}

	if a.RecordTTL != b.RecordTTL {
		return false
	}

	if len(a.Targets) != len(b.Targets) {
		return false
	}

	sort.Strings(a.Targets)
	sort.Strings(b.Targets)

	for i := range a.Targets {
		if a.Targets[i] != b.Targets[i] {
			return false
		}
	}

	if len(a.Labels) != len(b.Labels) {
		return false
	}

	for k := range a.Labels {
		if a.Labels[k] != b.Labels[k] {
			return false
		}
	}

	return true
}

// SameEndpoints compares two slices of endpoints regardless of order
// [x,y,z] == [z,x,y]
// [x,x,z] == [x,z,x]
// [x,y,y] != [x,x,y]
// [x,x,x] != [x,x,z]
func SameEndpoints(a, b []*endpoint.Endpoint) bool {
	if len(a) != len(b) {
		return false
	}

	sort.Slice(a, func(i, j int) bool { return a[i].String() < a[j].String() })
	sort.Slice(b, func(i, j int) bool { return b[i].String() < b[j].String() })

	for i := range a {
		if !SameEndpoint(a[i], b[i]) {
			return false
		}
	}

	return true
}

// SamePlanChanges verifies that two set of changes are the same
func SamePlanChanges(a, b map[string][]*endpoint.Endpoint) bool {
	return SameEndpoints(a["Create"], b["Create"]) && SameEndpoints(a["Delete"], b["Delete"]) &&
		SameEndpoints(a["UpdateOld"], b["UpdateOld"]) && SameEndpoints(a["UpdateNew"], b["UpdateNew"])
}
