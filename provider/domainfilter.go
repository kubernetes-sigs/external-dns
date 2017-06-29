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
	"strings"
)

// DomainFilter holds a lists of valid domain names
type DomainFilter struct {
	filters []string
}

// NewDomainFilter returns a new DomainFilter given a comma separated list of domains
func NewDomainFilter(domainFilters []string) DomainFilter {
	filters := make([]string, len(domainFilters))

	// user can define filter domains either with trailing dot or without, we remove all trailing periods from
	// the internal representation
	for i, domain := range domainFilters {
		filters[i] = strings.TrimSuffix(strings.TrimSpace(domain), ".")
	}

	return DomainFilter{filters}
}

// Match checks whether a domain can be found in the DomainFilter.
func (df DomainFilter) Match(domain string) bool {
	// return always true, if not filter is specified
	if len(df.filters) == 0 {
		return true
	}

	for _, filter := range df.filters {

		if strings.HasSuffix(strings.TrimSuffix(domain, "."), filter) {
			return true
		}
	}

	return false
}
