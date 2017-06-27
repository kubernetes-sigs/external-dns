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

package domains

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type domainFilterTest struct {
	domainFilter string
	domains      []string
	expected     bool
}

var domainFilterTests = []domainFilterTest{
	{
		"google.com.,exaring.de,inovex.de",
		[]string{"google.com", "exaring.de", "inovex.de"},
		true,
	},
	{
		"google.com.,exaring.de, inovex.de",
		[]string{"google.com", "exaring.de", "inovex.de"},
		true,
	},
	{
		"google.com.,exaring.de.,    inovex.de",
		[]string{"google.com", "exaring.de", "inovex.de"},
		true,
	},
	{
		"foo.org.      ",
		[]string{"foo.org"},
		true,
	},
	{
		"   foo.org",
		[]string{"foo.org"},
		true,
	},
	{
		"foo.org.",
		[]string{"foo.org"},
		true,
	},
	{
		"foo.org.",
		[]string{"baz.org"},
		false,
	},
	{
		"baz.foo.org.",
		[]string{"foo.org"},
		false,
	},
	{
		",foo.org.",
		[]string{"foo.org"},
		true,
	},
	{
		",foo.org.",
		[]string{},
		true,
	},
	{
		"",
		[]string{"foo.org"},
		true,
	},
	{
		"",
		[]string{},
		true,
	},
	{
		" ",
		[]string{},
		true,
	},
}

func TestDomainFilter_Match(t *testing.T) {
	for i, tt := range domainFilterTests {
		domainFilter := NewDomainFilter(tt.domainFilter)
		for _, domain := range tt.domains {
			require.Equal(t, tt.expected, domainFilter.Match(domain), "should not fail: %v in test-case #%v", domain, i)
		}
	}
}

func TestDomainFilter_Match_default_Filter_always_matches(t *testing.T) {
	for _, tt := range domainFilterTests {
		domainFilter := DomainFilter{}
		for i, domain := range tt.domains {
			require.True(t, domainFilter.Match(domain), "should not fail: %v in test-case #%v", domain, i)
		}
	}
}
