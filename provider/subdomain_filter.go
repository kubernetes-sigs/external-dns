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

// SubdomainFilter holds a lists of valid domain names
// A wrapper around DomainFilter
type SubdomainFilter struct {
	filters []string
	exclude []string
}

// NewSubdomainFilter returns a new DomainFilter given a comma separated list of domains
func NewSubdomainFilter(domainFilters []string) SubdomainFilter {
	newSubDomainFilter := NewDomainFilter(domainFilters)
	return SubdomainFilter(newSubDomainFilter)
}

// Match checks whether a domain can be found in the DomainFilter.
func (sdf SubdomainFilter) Match(domain string) bool {
	return matchFilter(sdf.filters, domain, true) && !matchFilter(sdf.exclude, domain, false)
}
