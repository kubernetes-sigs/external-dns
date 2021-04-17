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
	fs := make([]string, len(filters))
	for i, domain := range filters {
		fs[i] = strings.ToLower(strings.TrimSuffix(strings.TrimSpace(domain), "."))
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
	if df.regex != nil && df.regex.String() != "" {
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

	for _, filter := range filters {
		strippedDomain := strings.ToLower(strings.TrimSuffix(domain, "."))

		if filter == "" {
			return emptyval
		} else if strings.HasPrefix(filter, ".") && strings.HasSuffix(strippedDomain, filter) {
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

// IsConfigured returns true if DomainFilter is configured, false otherwise
func (df DomainFilter) IsConfigured() bool {
	if df.regex != nil && df.regex.String() != "" {
		return true
	} else if len(df.Filters) == 1 {
		return df.Filters[0] != ""
	}
	return len(df.Filters) > 0
}
