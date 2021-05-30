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
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

type domainFilterTest struct {
	domainFilter []string
	exclusions   []string
	domains      []string
	expected     bool
}

type regexDomainFilterTest struct {
	regex          *regexp.Regexp
	regexExclusion *regexp.Regexp
	domains        []string
	expected       bool
}

var domainFilterTests = []domainFilterTest{
	{
		[]string{"google.com.", "exaring.de", "inovex.de"},
		[]string{},
		[]string{"google.com", "exaring.de", "inovex.de"},
		true,
	},
	{
		[]string{"google.com.", "exaring.de", "inovex.de"},
		[]string{},
		[]string{"google.com", "exaring.de", "inovex.de"},
		true,
	},
	{
		[]string{"google.com.", "exaring.de.", "inovex.de"},
		[]string{},
		[]string{"google.com", "exaring.de", "inovex.de"},
		true,
	},
	{
		[]string{"foo.org.      "},
		[]string{},
		[]string{"foo.org"},
		true,
	},
	{
		[]string{"   foo.org"},
		[]string{},
		[]string{"foo.org"},
		true,
	},
	{
		[]string{"foo.org."},
		[]string{},
		[]string{"foo.org"},
		true,
	},
	{
		[]string{"foo.org."},
		[]string{},
		[]string{"baz.org"},
		false,
	},
	{
		[]string{"baz.foo.org."},
		[]string{},
		[]string{"foo.org"},
		false,
	},
	{
		[]string{"", "foo.org."},
		[]string{},
		[]string{"foo.org"},
		true,
	},
	{
		[]string{"", "foo.org."},
		[]string{},
		[]string{},
		true,
	},
	{
		[]string{""},
		[]string{},
		[]string{"foo.org"},
		true,
	},
	{
		[]string{""},
		[]string{},
		[]string{},
		true,
	},
	{
		[]string{" "},
		[]string{},
		[]string{},
		true,
	},
	{
		[]string{"bar.sub.example.org"},
		[]string{},
		[]string{"foo.bar.sub.example.org"},
		true,
	},
	{
		[]string{"example.org"},
		[]string{},
		[]string{"anexample.org", "test.anexample.org"},
		false,
	},
	{
		[]string{".example.org"},
		[]string{},
		[]string{"anexample.org", "test.anexample.org"},
		false,
	},
	{
		[]string{".example.org"},
		[]string{},
		[]string{"example.org"},
		false,
	},
	{
		[]string{".example.org"},
		[]string{},
		[]string{"test.example.org"},
		true,
	},
	{
		[]string{"anexample.org"},
		[]string{},
		[]string{"example.org", "test.example.org"},
		false,
	},
	{
		[]string{".org"},
		[]string{},
		[]string{"example.org", "test.example.org", "foo.test.example.org"},
		true,
	},
	{
		[]string{"example.org"},
		[]string{"api.example.org"},
		[]string{"example.org", "test.example.org", "foo.test.example.org"},
		true,
	},
	{
		[]string{"example.org"},
		[]string{"api.example.org"},
		[]string{"foo.api.example.org", "api.example.org"},
		false,
	},
	{
		[]string{"   example.org. "},
		[]string{"   .api.example.org    "},
		[]string{"foo.api.example.org", "bar.baz.api.example.org."},
		false,
	},
	{
		[]string{"example.org."},
		[]string{"api.example.org"},
		[]string{"dev-api.example.org", "qa-api.example.org"},
		true,
	},
	{
		[]string{"example.org."},
		[]string{"api.example.org"},
		[]string{"dev.api.example.org", "qa.api.example.org"},
		false,
	},
	{
		[]string{"example.org", "api.example.org"},
		[]string{"internal.api.example.org"},
		[]string{"foo.api.example.org"},
		true,
	},
	{
		[]string{"example.org", "api.example.org"},
		[]string{"internal.api.example.org"},
		[]string{"foo.internal.api.example.org"},
		false,
	},
	{
		[]string{"eXaMPle.ORG", "API.example.ORG"},
		[]string{"Foo-Bar.Example.Org"},
		[]string{"FoOoo.Api.Example.Org"},
		true,
	},
	{
		[]string{"eXaMPle.ORG", "API.example.ORG"},
		[]string{"api.example.org"},
		[]string{"foobar.Example.Org"},
		true,
	},
	{
		[]string{"eXaMPle.ORG", "API.example.ORG"},
		[]string{"api.example.org"},
		[]string{"foobar.API.Example.Org"},
		false,
	},
}

var regexDomainFilterTests = []regexDomainFilterTest{
	{
		regexp.MustCompile("\\.org$"),
		regexp.MustCompile(""),
		[]string{"foo.org", "bar.org", "foo.bar.org"},
		true,
	},
	{
		regexp.MustCompile("\\.bar\\.org$"),
		regexp.MustCompile(""),
		[]string{"foo.org", "bar.org", "example.com"},
		false,
	},
	{
		regexp.MustCompile("(?:foo|bar)\\.org$"),
		regexp.MustCompile(""),
		[]string{"foo.org", "bar.org", "example.foo.org", "example.bar.org", "a.example.foo.org", "a.example.bar.org"},
		true,
	},
	{
		regexp.MustCompile("(?:foo|bar)\\.org$"),
		regexp.MustCompile("^example\\.(?:foo|bar)\\.org$"),
		[]string{"foo.org", "bar.org", "a.example.foo.org", "a.example.bar.org"},
		true,
	},
	{
		regexp.MustCompile("(?:foo|bar)\\.org$"),
		regexp.MustCompile("^example\\.(?:foo|bar)\\.org$"),
		[]string{"example.foo.org", "example.bar.org"},
		false,
	},
	{
		regexp.MustCompile("(?:foo|bar)\\.org$"),
		regexp.MustCompile("^example\\.(?:foo|bar)\\.org$"),
		[]string{"foo.org", "bar.org", "a.example.foo.org", "a.example.bar.org"},
		true,
	},
}

func TestDomainFilterMatch(t *testing.T) {
	for i, tt := range domainFilterTests {
		if len(tt.exclusions) > 0 {
			t.Logf("NewDomainFilter() doesn't support exclusions - skipping test %+v", tt)
			continue
		}
		domainFilter := NewDomainFilter(tt.domainFilter)
		for _, domain := range tt.domains {
			assert.Equal(t, tt.expected, domainFilter.Match(domain), "should not fail: %v in test-case #%v", domain, i)
			assert.Equal(t, tt.expected, domainFilter.Match(domain+"."), "should not fail: %v in test-case #%v", domain+".", i)
		}
	}
}

func TestDomainFilterWithExclusions(t *testing.T) {
	for i, tt := range domainFilterTests {
		if len(tt.exclusions) == 0 {
			tt.exclusions = append(tt.exclusions, "")
		}
		domainFilter := NewDomainFilterWithExclusions(tt.domainFilter, tt.exclusions)
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

func TestDomainFilterMatchParent(t *testing.T) {
	parentMatchTests := []domainFilterTest{
		{
			[]string{"a.example.com."},
			[]string{},
			[]string{"example.com"},
			true,
		},
		{
			[]string{" a.example.com "},
			[]string{},
			[]string{"example.com"},
			true,
		},
		{
			[]string{""},
			[]string{},
			[]string{"example.com"},
			true,
		},
		{
			[]string{".a.example.com."},
			[]string{},
			[]string{"example.com"},
			false,
		},
		{
			[]string{"a.example.com.", "b.example.com"},
			[]string{},
			[]string{"example.com"},
			true,
		},
		{
			[]string{"a.example.com"},
			[]string{},
			[]string{"b.example.com"},
			false,
		},
		{
			[]string{"example.com"},
			[]string{},
			[]string{"example.com"},
			false,
		},
		{
			[]string{"example.com"},
			[]string{},
			[]string{"anexample.com"},
			false,
		},
		{
			[]string{""},
			[]string{},
			[]string{""},
			true,
		},
	}
	for i, tt := range parentMatchTests {
		domainFilter := NewDomainFilterWithExclusions(tt.domainFilter, tt.exclusions)
		for _, domain := range tt.domains {
			assert.Equal(t, tt.expected, domainFilter.MatchParent(domain), "should not fail: %v in test-case #%v", domain, i)
			assert.Equal(t, tt.expected, domainFilter.MatchParent(domain+"."), "should not fail: %v in test-case #%v", domain+".", i)
		}
	}
}

func TestRegexDomainFilter(t *testing.T) {
	for i, tt := range regexDomainFilterTests {
		domainFilter := NewRegexDomainFilter(tt.regex, tt.regexExclusion)
		for _, domain := range tt.domains {
			assert.Equal(t, tt.expected, domainFilter.Match(domain), "should not fail: %v in test-case #%v", domain, i)
			assert.Equal(t, tt.expected, domainFilter.Match(domain+"."), "should not fail: %v in test-case #%v", domain+".", i)
		}
	}
}

func TestPrepareFiltersStripsWhitespaceAndDotSuffix(t *testing.T) {
	for _, tt := range []struct {
		input  []string
		output []string
	}{
		{
			[]string{},
			[]string{},
		},
		{
			[]string{""},
			[]string{""},
		},
		{
			[]string{" ", "   ", ""},
			[]string{"", "", ""},
		},
		{
			[]string{"  foo   ", "  bar. ", "baz."},
			[]string{"foo", "bar", "baz"},
		},
		{
			[]string{"foo.bar", "  foo.bar.  ", " foo.bar.baz ", " foo.bar.baz.  "},
			[]string{"foo.bar", "foo.bar", "foo.bar.baz", "foo.bar.baz"},
		},
	} {
		t.Run("test string", func(t *testing.T) {
			assert.Equal(t, tt.output, prepareFilters(tt.input))
		})
	}
}

func TestMatchFilterReturnsProperEmptyVal(t *testing.T) {
	emptyFilters := []string{}
	assert.Equal(t, true, matchFilter(emptyFilters, "somedomain.com", true))
	assert.Equal(t, false, matchFilter(emptyFilters, "somedomain.com", false))
}

func TestDomainFilterIsConfigured(t *testing.T) {
	for _, tt := range []struct {
		filters  []string
		exclude  []string
		expected bool
	}{
		{
			[]string{""},
			[]string{""},
			false,
		},
		{
			[]string{"    "},
			[]string{"    "},
			false,
		},
		{
			[]string{"", ""},
			[]string{""},
			true,
		},
		{
			[]string{" . "},
			[]string{" . "},
			false,
		},
		{
			[]string{" notempty.com "},
			[]string{"  "},
			true,
		},
		{
			[]string{" notempty.com "},
			[]string{"  thisdoesntmatter.com "},
			true,
		},
	} {
		t.Run("test IsConfigured", func(t *testing.T) {
			df := NewDomainFilterWithExclusions(tt.filters, tt.exclude)
			assert.Equal(t, tt.expected, df.IsConfigured())
		})
	}
}
