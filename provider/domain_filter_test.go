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

package provider

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type domainFilterTest struct {
	domainFilter []string
	domains      []string
	expected     bool
}

var domainFilterTests = []domainFilterTest{
	{
		[]string{"google.com.", "exaring.de", "inovex.de"},
		[]string{"google.com", "exaring.de", "inovex.de"},
		true,
	},
	{
		[]string{"google.com.", "exaring.de", "inovex.de"},
		[]string{"google.com", "exaring.de", "inovex.de"},
		true,
	},
	{
		[]string{"google.com.", "exaring.de.", "inovex.de"},
		[]string{"google.com", "exaring.de", "inovex.de"},
		true,
	},
	{
		[]string{"foo.org.      "},
		[]string{"foo.org"},
		true,
	},
	{
		[]string{"   foo.org"},
		[]string{"foo.org"},
		true,
	},
	{
		[]string{"foo.org."},
		[]string{"foo.org"},
		true,
	},
	{
		[]string{"foo.org."},
		[]string{"baz.org"},
		false,
	},
	{
		[]string{"baz.foo.org."},
		[]string{"foo.org"},
		false,
	},
	{
		[]string{"", "foo.org."},
		[]string{"foo.org"},
		true,
	},
	{
		[]string{"", "foo.org."},
		[]string{},
		true,
	},
	{
		[]string{""},
		[]string{"foo.org"},
		true,
	},
	{
		[]string{""},
		[]string{},
		true,
	},
	{
		[]string{" "},
		[]string{},
		true,
	},
	{
		[]string{"bar.sub.example.org"},
		[]string{"foo.bar.sub.example.org"},
		true,
	},
	{
		[]string{"example.org"},
		[]string{"anexample.org", "test.anexample.org"},
		false,
	},
	{
		[]string{".example.org"},
		[]string{"anexample.org", "test.anexample.org"},
		false,
	},
	{
		[]string{".example.org"},
		[]string{"example.org"},
		false,
	},
	{
		[]string{".example.org"},
		[]string{"test.example.org"},
		true,
	},
	{
		[]string{"anexample.org"},
		[]string{"example.org", "test.example.org"},
		false,
	},
	{
		[]string{".org"},
		[]string{"example.org", "test.example.org", "foo.test.example.org"},
		true,
	},
}

func TestDomainFilterMatch(t *testing.T) {
	for i, tt := range domainFilterTests {
		domainFilter := NewDomainFilter(tt.domainFilter)
		for _, domain := range tt.domains {
			assert.Equal(t, tt.expected, domainFilter.Match(domain), "should not fail: %v in test-case #%v", domain, i)
			assert.Equal(t, tt.expected, domainFilter.Match(domain+"."), "should not fail: %v in test-case #%v", domain+".", i)
		}
	}
}

func TestDomainFilterMatchWithEmptyFilter(t *testing.T) {
	for _, tt := range domainFilterTests {
		domainFilter := DomainFilter{}
		for i, domain := range tt.domains {
			assert.True(t, domainFilter.Match(domain), "should not fail: %v in test-case #%v", domain, i)
			assert.True(t, domainFilter.Match(domain+"."), "should not fail: %v in test-case #%v", domain+".", i)
		}
	}
}
