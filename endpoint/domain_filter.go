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

	log "github.com/sirupsen/logrus"
)

// DomainFilter holds a lists of valid domain names
type DomainFilter struct {
	// Filters define what domains to match
	Filters []string
	// exclude define what domains not to match
	exclude []string
	// regex defines a regular expression to match the domains
	regex string
	// regexExclusion defines a regular expression to exclude the domains matched
	regexExclusion string
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
	return DomainFilter{prepareFilters(domainFilters), prepareFilters(excludeDomains), "", ""}
}

// NewDomainFilter returns a new DomainFilter given a comma separated list of domains
func NewDomainFilter(domainFilters []string) DomainFilter {
	return DomainFilter{prepareFilters(domainFilters), []string{}, "", ""}
}

// NewRegexDomainFilter returns a new DomainFilter given a regular expression
func NewRegexDomainFilter(regexDomainFilter string, regexDomainExclusion string) DomainFilter {
	return DomainFilter{[]string{}, []string{}, regexDomainFilter, regexDomainExclusion}
}

// Match checks whether a domain can be found in the DomainFilter.
// RegexFilter takes precedence over Filters
func (df DomainFilter) Match(domain string) bool {
	if df.regex != "" {
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

// matchRegex determines if a domain matches the configured regular expressions in the DomainFilter.
// The negativeRegex, if set, takes precedence over regex. Therefore,
// matchRegex returns true when only regex regular expression matches the domain.
// Otherwise, if either negativeRegex matches or regex does not match the domain, it will return false.
func matchRegex(regex string, negativeRegex string, domain string) bool {
	strippedDomain := strings.ToLower(strings.TrimSuffix(domain, "."))

	if negativeRegex != "" {
		match, err := regexp.MatchString(negativeRegex, strippedDomain)
		if err != nil {
			log.Errorf("Failed to filter domain %s with the regex-exclusion filter: %v", domain, err)
		}
		if match {
			return false
		}
	}
	match, err := regexp.MatchString(regex, strippedDomain)
	if err != nil {
		log.Errorf("Failed to filter domain %s with the regex filter: %v", domain, err)
	}
	return match
}

// IsConfigured returns true if DomainFilter is configured, false otherwise
func (df DomainFilter) IsConfigured() bool {
	if df.regex != "" {
		return true
	} else if len(df.Filters) == 1 {
		return df.Filters[0] != ""
	}
	return len(df.Filters) > 0
}
