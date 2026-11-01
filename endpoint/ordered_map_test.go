/*
Copyright 2026 The Kubernetes Authors.

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

package endpoint

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// resourceEP builds an endpoint with the given ResourceLabelKey label (unset when empty).
func resourceEP(name, resource string, targets ...string) *Endpoint {
	ep := NewEndpoint(name, RecordTypeA, targets...)
	if resource != "" {
		ep.Labels[ResourceLabelKey] = resource
	}
	return ep
}

func resourceEPNames(eps []*Endpoint) []string {
	names := make([]string, 0, len(eps))
	for _, ep := range eps {
		names = append(names, ep.DNSName)
	}
	return names
}

func TestOrderedMap_KeepsFirstSeenKeyOrder(t *testing.T) {
	om := NewOrderedMap[string](3)
	b1, a1, b2 := resourceEP("b1.example.com", "", "1.1.1.1"),
		resourceEP("a1.example.com", "", "1.1.1.1"),
		resourceEP("b2.example.com", "", "1.1.1.1")
	om.Append("b", b1)
	om.Append("a", a1)
	om.Append("b", b2)

	var keys []string
	groups := map[string][]*Endpoint{}
	for k, v := range om.All() {
		keys = append(keys, k)
		groups[k] = v
	}

	assert.Equal(t, []string{"b", "a"}, keys)
	assert.Equal(t, []*Endpoint{b1, b2}, groups["b"])
	assert.Equal(t, []*Endpoint{a1}, groups["a"])
}

func TestSortByResource(t *testing.T) {
	eps := []*Endpoint{
		resourceEP("x1.example.com", "b", "1.1.1.1"),
		resourceEP("x2.example.com", "a", "1.1.1.1"),
		resourceEP("x3.example.com", "b", "1.1.1.1"),
		resourceEP("x4.example.com", "", "1.1.1.1"),
	}

	SortByResource(eps)

	// unlabeled first, then by label; equal labels keep input order
	assert.Equal(t, []string{"x4.example.com", "x2.example.com", "x1.example.com", "x3.example.com"}, resourceEPNames(eps))
}

func TestMergeEndpoints_SortedByResource(t *testing.T) {
	tests := []struct {
		name     string
		input    []*Endpoint
		expected []string
	}{
		{
			name: "records are ordered by resource label",
			input: []*Endpoint{
				resourceEP("z.example.com", "svc/z", "1.1.1.1"),
				resourceEP("a.example.com", "svc/a", "1.1.1.1"),
				resourceEP("m.example.com", "svc/m", "1.1.1.1"),
			},
			expected: []string{"a.example.com", "m.example.com", "z.example.com"},
		},
		{
			name: "unlabeled records sort first and keep first-seen order",
			input: []*Endpoint{
				resourceEP("c.example.com", "", "1.1.1.1"),
				resourceEP("b.example.com", "svc/a", "1.1.1.1"),
				resourceEP("a.example.com", "", "1.1.1.1"),
			},
			expected: []string{"c.example.com", "a.example.com", "b.example.com"},
		},
		{
			name: "merged record is positioned by its lowest resource",
			input: []*Endpoint{
				resourceEP("x.example.com", "svc/m", "1.1.1.1"),
				resourceEP("y.example.com", "svc/c", "1.1.1.1"),
				resourceEP("x.example.com", "svc/b", "2.2.2.2"),
				resourceEP("w.example.com", "svc/a", "1.1.1.1"),
			},
			expected: []string{"w.example.com", "x.example.com", "y.example.com"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, resourceEPNames(MergeEndpoints(tt.input)))
		})
	}
}

func TestMergeEndpoints_BaseIsLowestResource(t *testing.T) {
	// The base keeps its metadata no matter which order the inputs arrive in.
	for _, order := range [][]string{{"svc/b", "svc/a"}, {"svc/a", "svc/b"}} {
		targets := map[string]string{"svc/a": "1.1.1.1", "svc/b": "2.2.2.2"}
		in := make([]*Endpoint, 0, len(order))
		for _, res := range order {
			in = append(in, resourceEP("a.example.com", res, targets[res]).WithProviderSpecific("marker", res))
		}

		result := MergeEndpoints(in)

		require.Len(t, result, 1, "input order %v", order)
		marker, _ := result[0].GetProviderSpecificProperty("marker")
		assert.Equal(t, "svc/a", marker, "input order %v", order)
		assert.Equal(t, "svc/a", result[0].Labels[ResourceLabelKey], "input order %v", order)
		assert.Equal(t, Targets{"1.1.1.1", "2.2.2.2"}, result[0].Targets, "input order %v", order)
	}
}
