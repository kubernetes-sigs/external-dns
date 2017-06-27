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
	"strings"
)

type DomainFilter struct {
	filters []string
}

func NewDomainFilter(domainFilter string) DomainFilter {

	if strings.TrimSpace(domainFilter) == "" {
		return DomainFilter{[]string{""}}
	}

	filters := strings.Split(domainFilter, ",")

	for i := 0; i < len(filters); i++ {
		filters[i] = strings.TrimSuffix(strings.TrimSpace(filters[i]), ".") + "."
		if filters[i] == "." {
			filters = append(filters[:i], filters[i+1:]...)
			i--
		}
	}

	return DomainFilter{filters}
}

func (df DomainFilter) Match(domain string) bool {

	// return always true, if not filter is specified
	if len(df.filters) == 0 {
		return true
	}

	for _, filter := range df.filters {

		// user can define domains either with trailing dot or without.
		if strings.HasSuffix(domain, filter) || strings.HasSuffix(domain+".", filter) {
			return true
		}
	}

	return false
}
