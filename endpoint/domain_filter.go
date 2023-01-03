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
	"strings"
)

// DomainFilterInterface defines the interface to select matching domains for a specific provider or runtime
type DomainFilterInterface interface {
	Match(domain string) bool
	IsConfigured() bool
}

type MatchAllDomainFilters []DomainFilterInterface

func (f MatchAllDomainFilters) Match(domain string) bool {
	if !f.IsConfigured() {
		return true
	}
	for _, filter := range f {
		if filter == nil {
			continue
		}
		if filter.IsConfigured() && !filter.Match(domain) {
			return false
		}
	}
	return true
}

func (f MatchAllDomainFilters) IsConfigured() bool {
	if f == nil {
		return false
	}
	for _, filter := range f {
		if filter == nil {
			continue
		}
		if filter.IsConfigured() {
			return true
		}
	}
	return len(f) > 0
}

// DomainFilter holds a lists of valid domain names
type DomainFilter struct {
	// Filters define what domains to match
	Filters []string
	// exclude define what domains not to match
	exclude []string
	// regex defines a regular expression to match the domains
	regex *regexp.Regexp
	// regexExclusion defines a regular expression to exclude the domains matched
	regexExclusion *regexp.Regexp
}

// prepareFilters provides consistent trimming for filters/exclude params
func prepareFilters(filters []string) []string {
	var fs []string
	for _, filter := range filters {
		if domain := strings.ToLower(strings.TrimSuffix(strings.TrimSpace(filter), ".")); domain != "" {
			fs = append(fs, domain)
		}
	}
	return fs
}

// NewDomainFilterWithExclusions returns a new DomainFilter, given a list of matches and exclusions
func NewDomainFilterWithExclusions(domainFilters []string, excludeDomains []string) DomainFilter {
	return DomainFilter{Filters: prepareFilters(domainFilters), exclude: prepareFilters(excludeDomains)}
}

// NewDomainFilter returns a new DomainFilter given a comma separated list of domains
func NewDomainFilter(domainFilters []string) DomainFilter {
	return DomainFilter{Filters: prepareFilters(domainFilters)}
}

// NewRegexDomainFilter returns a new DomainFilter given a regular expression
func NewRegexDomainFilter(regexDomainFilter *regexp.Regexp, regexDomainExclusion *regexp.Regexp) DomainFilter {
	return DomainFilter{regex: regexDomainFilter, regexExclusion: regexDomainExclusion}
}

// Match checks whether a domain can be found in the DomainFilter.
// RegexFilter takes precedence over Filters
func (df DomainFilter) Match(domain string) bool {
	if df.regex != nil && df.regex.String() != "" || df.regexExclusion != nil && df.regexExclusion.String() != "" {
		return matchRegex(df.regex, df.regexExclusion, domain)
	}

	return matchFilter(df.Filters, domain, true) && !matchFilter(df.exclude, domain, false)
}

// matchFilter determines if any `filters` match `domain`.
// If no `filters` are provided, behavior depends on `emptyval`
// (empty `df.filters` matches everything, while empty `df.exclude` excludes nothing)
func matchFilter(filters []string, domain string, emptyval bool) bool {
	if len(filters) == 0 {
		return emptyval
	}

	strippedDomain := strings.ToLower(strings.TrimSuffix(domain, "."))
	for _, filter := range filters {
		if filter == "" {
			continue
		}

		if strings.HasPrefix(filter, ".") && strings.HasSuffix(strippedDomain, filter) {
			return true
		} else if strings.Count(strippedDomain, ".") == strings.Count(filter, ".") {
			if strippedDomain == filter {
				return true
			}
		} else if strings.HasSuffix(strippedDomain, "."+filter) {
			return true
		}
	}
	return false
}

// matchRegex determines if a domain matches the configured regular expressions in DomainFilter.
// negativeRegex, if set, takes precedence over regex.  Therefore, matchRegex returns true when
// only regex regular expression matches the domain
// Otherwise, if either negativeRegex matches or regex does not match the domain, it returns false
func matchRegex(regex *regexp.Regexp, negativeRegex *regexp.Regexp, domain string) bool {
	strippedDomain := strings.ToLower(strings.TrimSuffix(domain, "."))

	if negativeRegex != nil && negativeRegex.String() != "" {
		return !negativeRegex.MatchString(strippedDomain)
	}
	return regex.MatchString(strippedDomain)
}

// MatchParent checks wether DomainFilter matches a given parent domain.
func (df DomainFilter) MatchParent(domain string) bool {
	if matchFilter(df.exclude, domain, false) {
		return false
	}
	if len(df.Filters) == 0 {
		return true
	}

	strippedDomain := strings.ToLower(strings.TrimSuffix(domain, "."))
	for _, filter := range df.Filters {
		if filter == "" || strings.HasPrefix(filter, ".") {
			// We don't check parents if the filter is prefixed with "."
			continue
		}
		if strings.HasSuffix(filter, "."+strippedDomain) {
			return true
		}
	}
	return false
}

// IsConfigured returns true if any inclusion or exclusion rules have been specified.
func (df DomainFilter) IsConfigured() bool {
	if df.regex != nil && df.regex.String() != "" {
		return true
	} else if df.regexExclusion != nil && df.regexExclusion.String() != "" {
		return true
	}
	return len(df.Filters) > 0 || len(df.exclude) > 0
}
