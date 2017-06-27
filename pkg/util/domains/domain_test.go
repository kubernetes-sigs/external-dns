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
