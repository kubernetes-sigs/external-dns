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
	"net/netip"
	"reflect"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/external-dns/endpoint"
)

func TestExampleSameEndpoints(t *testing.T) {
	eps := []*endpoint.Endpoint{
		{
			DNSName: "example.org",
			Targets: endpoint.Targets{"load-balancer.org"},
		},
		{
			DNSName:    "example.org",
			Targets:    endpoint.Targets{"load-balancer.org"},
			RecordType: endpoint.RecordTypeTXT,
		},
		{
			DNSName:    "abc.com",
			Targets:    endpoint.Targets{"something"},
			RecordType: endpoint.RecordTypeTXT,
		},
		{
			DNSName:       "abc.com",
			Targets:       endpoint.Targets{"1.2.3.4"},
			RecordType:    endpoint.RecordTypeA,
			SetIdentifier: "test-set-1",
		},
		{
			DNSName:    "bbc.com",
			Targets:    endpoint.Targets{"foo.com"},
			RecordType: endpoint.RecordTypeCNAME,
		},
		{
			DNSName:    "cbc.com",
			Targets:    endpoint.Targets{"foo.com"},
			RecordType: "CNAME",
			RecordTTL:  endpoint.TTL(60),
		},
		{
			DNSName: "example.org",
			Targets: endpoint.Targets{"load-balancer.org"},
			ProviderSpecific: endpoint.ProviderSpecific{
				endpoint.ProviderSpecificProperty{Name: "foo", Value: "bar"},
			},
		},
	}
	sort.Sort(byAllFields(eps))
	for _, ep := range eps {
		fmt.Println(ep)
	}
	// Output:
	// abc.com 0 IN A test-set-1 1.2.3.4 []
	// abc.com 0 IN TXT  something []
	// bbc.com 0 IN CNAME  foo.com []
	// cbc.com 60 IN CNAME  foo.com []
	// example.org 0 IN   load-balancer.org []
	// example.org 0 IN   load-balancer.org [{foo bar}]
	// example.org 0 IN TXT  load-balancer.org []
}

func makeEndpoint(DNSName string) *endpoint.Endpoint {
	return &endpoint.Endpoint{
		DNSName:       DNSName,
		Targets:       endpoint.Targets{"target.com"},
		RecordType:    "A",
		SetIdentifier: "set1",
		RecordTTL:     300,
		Labels: map[string]string{
			endpoint.OwnerLabelKey:       "owner",
			endpoint.ResourceLabelKey:    "resource",
			endpoint.OwnedRecordLabelKey: "owned",
		},
		ProviderSpecific: endpoint.ProviderSpecific{
			{Name: "key", Value: "val"},
		},
	}
}

func TestSameEndpoint(t *testing.T) {
	tests := []struct {
		name           string
		a              *endpoint.Endpoint
		b              *endpoint.Endpoint
		isSameEndpoint bool
	}{
		{
			name:           "DNSName is not equal",
			a:              &endpoint.Endpoint{DNSName: "example.org"},
			b:              &endpoint.Endpoint{DNSName: "example.com"},
			isSameEndpoint: false,
		},
		{
			name: "All fields are equal",
			a: &endpoint.Endpoint{
				DNSName:       "example.org",
				Targets:       endpoint.Targets{"lb.example.com"},
				RecordType:    "A",
				SetIdentifier: "set-1",
				RecordTTL:     300,
				Labels: map[string]string{
					endpoint.OwnerLabelKey:       "owner-1",
					endpoint.ResourceLabelKey:    "resource-1",
					endpoint.OwnedRecordLabelKey: "owned-true",
				},
				ProviderSpecific: endpoint.ProviderSpecific{
					{Name: "key1", Value: "val1"},
				},
			},
			b: &endpoint.Endpoint{
				DNSName:       "example.org",
				Targets:       endpoint.Targets{"lb.example.com"},
				RecordType:    "A",
				SetIdentifier: "set-1",
				RecordTTL:     300,
				Labels: map[string]string{
					endpoint.OwnerLabelKey:       "owner-1",
					endpoint.ResourceLabelKey:    "resource-1",
					endpoint.OwnedRecordLabelKey: "owned-true",
				},
				ProviderSpecific: endpoint.ProviderSpecific{
					{Name: "key1", Value: "val1"},
				},
			},
			isSameEndpoint: true,
		},
		{
			name:           "Different Targets",
			a:              &endpoint.Endpoint{DNSName: "example.org", Targets: endpoint.Targets{"a.com"}},
			b:              &endpoint.Endpoint{DNSName: "example.org", Targets: endpoint.Targets{"b.com"}},
			isSameEndpoint: false,
		},
		{
			name:           "Different RecordType",
			a:              &endpoint.Endpoint{DNSName: "example.org", RecordType: "A"},
			b:              &endpoint.Endpoint{DNSName: "example.org", RecordType: "CNAME"},
			isSameEndpoint: false,
		},
		{
			name:           "Different SetIdentifier",
			a:              &endpoint.Endpoint{DNSName: "example.org", SetIdentifier: "id1"},
			b:              &endpoint.Endpoint{DNSName: "example.org", SetIdentifier: "id2"},
			isSameEndpoint: false,
		},
		{
			name: "Different OwnerLabelKey",
			a: &endpoint.Endpoint{
				DNSName: "example.org",
				Labels: map[string]string{
					endpoint.OwnerLabelKey: "owner1",
				},
			},
			b: &endpoint.Endpoint{
				DNSName: "example.org",
				Labels: map[string]string{
					endpoint.OwnerLabelKey: "owner2",
				},
			},
			isSameEndpoint: false,
		},
		{
			name:           "Different RecordTTL",
			a:              &endpoint.Endpoint{DNSName: "example.org", RecordTTL: 300},
			b:              &endpoint.Endpoint{DNSName: "example.org", RecordTTL: 400},
			isSameEndpoint: false,
		},
		{
			name: "Different ProviderSpecific",
			a: &endpoint.Endpoint{
				DNSName: "example.org",
				ProviderSpecific: endpoint.ProviderSpecific{
					{Name: "key1", Value: "val1"},
				},
			},
			b: &endpoint.Endpoint{
				DNSName: "example.org",
				ProviderSpecific: endpoint.ProviderSpecific{
					{Name: "key1", Value: "val2"},
				},
			},
			isSameEndpoint: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isSameEndpoint := SameEndpoint(tt.a, tt.b)
			assert.Equal(t, tt.isSameEndpoint, isSameEndpoint)
		})
	}
}
func TestSameEndpoints(t *testing.T) {
	tests := []struct {
		name string
		a, b []*endpoint.Endpoint
		want bool
	}{
		{
			name: "Both slices nil",
			a:    nil,
			b:    nil,
			want: true,
		},
		{
			name: "One nil, one empty",
			a:    []*endpoint.Endpoint{},
			b:    nil,
			want: true,
		},
		{
			name: "Different lengths",
			a:    []*endpoint.Endpoint{makeEndpoint("a.com")},
			b:    []*endpoint.Endpoint{},
			want: false,
		},
		{
			name: "Same endpoints in same order",
			a:    []*endpoint.Endpoint{makeEndpoint("a.com"), makeEndpoint("b.com")},
			b:    []*endpoint.Endpoint{makeEndpoint("a.com"), makeEndpoint("b.com")},
			want: true,
		},
		{
			name: "Same endpoints in different order",
			a:    []*endpoint.Endpoint{makeEndpoint("b.com"), makeEndpoint("a.com")},
			b:    []*endpoint.Endpoint{makeEndpoint("a.com"), makeEndpoint("b.com")},
			want: true,
		},
		{
			name: "One endpoint differs",
			a:    []*endpoint.Endpoint{makeEndpoint("a.com"), makeEndpoint("b.com")},
			b:    []*endpoint.Endpoint{makeEndpoint("a.com"), makeEndpoint("c.com")},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isSameEndpoints := SameEndpoints(tt.a, tt.b)
			assert.Equal(t, tt.want, isSameEndpoints)
		})
	}
}

func TestSameEndpointLabel(t *testing.T) {
	tests := []struct {
		name string
		a    []*endpoint.Endpoint
		b    []*endpoint.Endpoint
		want bool
	}{
		{
			name: "length of a and b are not same",
			a:    []*endpoint.Endpoint{makeEndpoint("a.com")},
			b:    []*endpoint.Endpoint{makeEndpoint("b.com"), makeEndpoint("c.com")},
			want: false,
		},
		{
			name: "endpoint's labels are same in a and b",
			a:    []*endpoint.Endpoint{makeEndpoint("a.com"), makeEndpoint("c.com")},
			b:    []*endpoint.Endpoint{makeEndpoint("b.com"), makeEndpoint("c.com")},
			want: true,
		},
		{
			name: "endpoint's labels are not same in a and b",
			a: []*endpoint.Endpoint{
				{
					DNSName: "a.com",
					Labels: endpoint.Labels{
						endpoint.OwnerLabelKey:    "owner1",
						endpoint.ResourceLabelKey: "resource1",
					},
				},
				{
					DNSName: "b.com",
					Labels: endpoint.Labels{
						endpoint.OwnerLabelKey:    "owner2",
						endpoint.ResourceLabelKey: "resource2",
					},
				},
			},
			b: []*endpoint.Endpoint{
				{
					DNSName: "a.com",
					Labels: endpoint.Labels{
						endpoint.OwnerLabelKey:    "owner",
						endpoint.ResourceLabelKey: "resource",
					},
				},
				{
					DNSName: "b.com",
					Labels: endpoint.Labels{
						endpoint.OwnerLabelKey:    "owner1",
						endpoint.ResourceLabelKey: "resource1",
					},
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isSameEndpointLabels := SameEndpointLabels(tt.a, tt.b)
			assert.Equal(t, tt.want, isSameEndpointLabels)
		})
	}
}

func TestSamePlanChanges(t *testing.T) {
	tests := []struct {
		name string
		a    map[string][]*endpoint.Endpoint
		b    map[string][]*endpoint.Endpoint
		want bool
	}{
		{
			name: "endpoints with all operations in a and b are same",
			a: map[string][]*endpoint.Endpoint{
				"Create":    {makeEndpoint("a.com")},
				"Delete":    {makeEndpoint("b.com")},
				"UpdateOld": {makeEndpoint("a.com")},
				"UpdateNew": {makeEndpoint("c.com")},
			},
			b: map[string][]*endpoint.Endpoint{
				"Create":    {makeEndpoint("a.com")},
				"Delete":    {makeEndpoint("b.com")},
				"UpdateOld": {makeEndpoint("a.com")},
				"UpdateNew": {makeEndpoint("c.com")},
			},
			want: true,
		},
		{
			name: "endpoints for create operations in a and b are not same",
			a: map[string][]*endpoint.Endpoint{
				"Create":    {makeEndpoint("a.com")},
				"Delete":    {makeEndpoint("b.com")},
				"UpdateOld": {makeEndpoint("a.com")},
				"UpdateNew": {makeEndpoint("c.com")},
			},
			b: map[string][]*endpoint.Endpoint{
				"Create":    {makeEndpoint("x.com")},
				"Delete":    {makeEndpoint("b.com")},
				"UpdateOld": {makeEndpoint("a.com")},
				"UpdateNew": {makeEndpoint("c.com")},
			},
			want: false,
		},
		{
			name: "endpoints for delete operations in a and b are not same",
			a: map[string][]*endpoint.Endpoint{
				"Create":    {makeEndpoint("a.com")},
				"Delete":    {makeEndpoint("b.com")},
				"UpdateOld": {makeEndpoint("a.com")},
				"UpdateNew": {makeEndpoint("c.com")},
			},
			b: map[string][]*endpoint.Endpoint{
				"Create":    {makeEndpoint("a.com")},
				"Delete":    {makeEndpoint("g.com")},
				"UpdateOld": {makeEndpoint("a.com")},
				"UpdateNew": {makeEndpoint("c.com")},
			},
			want: false,
		},
		{
			name: "endpoints for updateOld operations in a and b are not same",
			a: map[string][]*endpoint.Endpoint{
				"Create":    {makeEndpoint("a.com")},
				"Delete":    {makeEndpoint("b.com")},
				"UpdateOld": {makeEndpoint("b.com")},
				"UpdateNew": {makeEndpoint("c.com")},
			},
			b: map[string][]*endpoint.Endpoint{
				"Create":    {makeEndpoint("a.com")},
				"Delete":    {makeEndpoint("b.com")},
				"UpdateOld": {makeEndpoint("c.com")},
				"UpdateNew": {makeEndpoint("c.com")},
			},
			want: false,
		},
		{
			name: "endpoints for updateNew operations in a and b are same",
			a: map[string][]*endpoint.Endpoint{
				"Create":    {makeEndpoint("a.com")},
				"Delete":    {makeEndpoint("b.com")},
				"UpdateOld": {makeEndpoint("a.com")},
				"UpdateNew": {makeEndpoint("d.com")},
			},
			b: map[string][]*endpoint.Endpoint{
				"Create":    {makeEndpoint("a.com")},
				"Delete":    {makeEndpoint("b.com")},
				"UpdateOld": {makeEndpoint("a.com")},
				"UpdateNew": {makeEndpoint("c.com")},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkPlanChanges := SamePlanChanges(tt.a, tt.b)
			assert.Equal(t, tt.want, checkPlanChanges)
		})
	}
}
func TestNewTargetsFromAddr(t *testing.T) {
	tests := []struct {
		name     string
		input    []netip.Addr
		expected endpoint.Targets
	}{
		{
			name:     "empty slice",
			input:    []netip.Addr{},
			expected: endpoint.Targets{},
		},
		{
			name: "single IPv4 address",
			input: []netip.Addr{
				netip.MustParseAddr("192.0.2.1"),
			},
			expected: endpoint.Targets{"192.0.2.1"},
		},
		{
			name: "multiple IP addresses",
			input: []netip.Addr{
				netip.MustParseAddr("192.0.2.1"),
				netip.MustParseAddr("2001:db8::1"),
			},
			expected: endpoint.Targets{"192.0.2.1", "2001:db8::1"},
		},
		{
			name: "IPv6 address only",
			input: []netip.Addr{
				netip.MustParseAddr("::1"),
			},
			expected: endpoint.Targets{"::1"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewTargetsFromAddr(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("NewTargetsFromAddr() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestWithLabel(t *testing.T) {
	e := &endpoint.Endpoint{}
	// should initialize Labels and set the key
	returned := e.WithLabel("foo", "bar")
	assert.Equal(t, e, returned)
	assert.NotNil(t, e.Labels)
	assert.Equal(t, "bar", e.Labels["foo"])

	// overriding an existing key
	e2 := e.WithLabel("foo", "baz")
	assert.Equal(t, e, e2)
	assert.Equal(t, "baz", e.Labels["foo"])

	// adding a new key without wiping others
	e.Labels["existing"] = "orig"
	e.WithLabel("new", "val")
	assert.Equal(t, "orig", e.Labels["existing"])
	assert.Equal(t, "val", e.Labels["new"])
}
