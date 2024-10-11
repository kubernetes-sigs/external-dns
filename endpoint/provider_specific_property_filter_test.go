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

package endpoint

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type providerSpecificPropertyFilterTest struct {
	names             []string
	prefixes          []string
	endpoints         []*Endpoint
	expectedEndpoints []*Endpoint
}

var providerSpecificPropertyFilterTests = []providerSpecificPropertyFilterTest{
	// 0. No Names or Prefixes
	{
		[]string{},
		[]string{},
		[]*Endpoint{
			{
				DNSName: "a.example.com",
				ProviderSpecific: []ProviderSpecificProperty{
					{Name: "prefix1/name1"},
					{Name: "name1"},
				},
			},
			{
				DNSName: "b.example.com",
				ProviderSpecific: []ProviderSpecificProperty{
					{Name: "prefix2/name1"},
					{Name: "name2"},
				},
			},
		},
		[]*Endpoint{
			{
				DNSName:          "a.example.com",
				ProviderSpecific: []ProviderSpecificProperty{},
			},
			{
				DNSName:          "b.example.com",
				ProviderSpecific: []ProviderSpecificProperty{},
			},
		},
	},
	// 1. Only Names
	{
		[]string{"name1", "name2"},
		[]string{},
		[]*Endpoint{
			{
				DNSName: "a.example.com",
				ProviderSpecific: []ProviderSpecificProperty{
					{Name: "name1"},
					{Name: "name2"},
					{Name: "name3"},
					{Name: "prefix1/name1"},
				},
			},
		},
		[]*Endpoint{
			{
				DNSName: "a.example.com",
				ProviderSpecific: []ProviderSpecificProperty{
					{Name: "name1"},
					{Name: "name2"},
				},
			},
		},
	},
	// 2. Only Prefixes
	{
		[]string{},
		[]string{"prefix1/", "prefix2"},
		[]*Endpoint{
			{
				DNSName: "a.example.com",
				ProviderSpecific: []ProviderSpecificProperty{
					{Name: "prefix1/foo"},
					{Name: "prefix1-bar"},
					{Name: "prefix2"},
					{Name: "prefix3"},
				},
			},
		},
		[]*Endpoint{
			{
				DNSName: "a.example.com",
				ProviderSpecific: []ProviderSpecificProperty{
					{Name: "prefix1/foo"},
					{Name: "prefix2"},
				},
			},
		},
	},
	// 3. No Endpoints
	{
		[]string{"name1"},
		[]string{"prefix1/"},
		[]*Endpoint{},
		[]*Endpoint{},
	},
	// 4. Both Names and Prefixes
	{
		[]string{"name1"},
		[]string{"prefix1/"},
		[]*Endpoint{
			{
				DNSName: "a.example.com",
				ProviderSpecific: []ProviderSpecificProperty{
					{Name: "prefix1/property"},
					{Name: "prefix2/property"},
					{Name: "name10"},
				},
			},
			{
				DNSName: "b.example.com",
				ProviderSpecific: []ProviderSpecificProperty{
					{Name: "prefix1/property"},
					{Name: "name1"},
					{Name: "name2"},
				},
			},
		},
		[]*Endpoint{
			{
				DNSName: "a.example.com",
				ProviderSpecific: []ProviderSpecificProperty{
					{Name: "prefix1/property"},
				},
			},
			{
				DNSName: "b.example.com",
				ProviderSpecific: []ProviderSpecificProperty{
					{Name: "prefix1/property"},
					{Name: "name1"},
				},
			},
		},
	},
}

func TestPropertyFilter(t *testing.T) {
	for i, tt := range providerSpecificPropertyFilterTests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			propertyFilter := ProviderSpecificPropertyFilter{
				Names:    tt.names,
				Prefixes: tt.prefixes,
			}

			propertyFilter.Filter(tt.endpoints)

			assert.Equal(t, tt.endpoints, tt.expectedEndpoints)
		})
	}
}
