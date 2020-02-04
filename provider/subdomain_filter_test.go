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

type subdomainFilterTest struct {
	domainFilter []string
	domains      []string
	expected     bool
}

var subdomainFilterTests = []subdomainFilterTest{
	{
		[]string{"foobar.google.com.", "foo.exaring.de", "foo.inovex.de"},
		[]string{"dev.foobar.google.com", "dev.foo.exaring.de", "dev.foo.inovex.de"},
		true,
	},
	{
		[]string{"foobar.google.com.", "foo.exaring.de", "foo.inovex.de"},
		[]string{"dev.google.com", "dev.exaring.de", "dev.inovex.de"},
		false,
	},
}

func TestSubdomainFilterMatch(t *testing.T) {
	for i, tt := range subdomainFilterTests {
		subdomainFilter := NewSubdomainFilter(tt.domainFilter)
		for _, domain := range tt.domains {
			assert.Equal(t, tt.expected, subdomainFilter.Match(domain), "should not fail: %v in test-case #%v", domain, i)
			assert.Equal(t, tt.expected, subdomainFilter.Match(domain+"."), "should not fail: %v in test-case #%v", domain+".", i)
		}
	}
}

func TestSubdomainFilterMatchWithEmptyFilter(t *testing.T) {
	for _, tt := range subdomainFilterTests {
		subdomainFilter := SubdomainFilter{}
		for i, domain := range tt.domains {
			assert.True(t, subdomainFilter.Match(domain), "should not fail: %v in test-case #%v", domain, i)
			assert.True(t, subdomainFilter.Match(domain+"."), "should not fail: %v in test-case #%v", domain+".", i)
		}
	}
}
