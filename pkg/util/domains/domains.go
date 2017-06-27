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
