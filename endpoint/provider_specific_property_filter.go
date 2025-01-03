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
	"strings"

	log "github.com/sirupsen/logrus"
)

type ProviderSpecificPropertyFilter struct {
	// Names of the provider specific properties to match
	Names []string
	// Prefixes of the provider specific properties to match
	Prefixes []string
}

// Match checks whether a ProviderSpecificProperty.Name can be found in the
// ProviderSpecificPropertyFilter.
func (pf ProviderSpecificPropertyFilter) Match(name string) bool {
	for _, pfName := range pf.Names {
		if pfName == name {
			return true
		}
	}
	for _, prefix := range pf.Prefixes {
		if strings.HasPrefix(name, prefix) {
			return true
		}
	}
	return false
}

// Filter removes all ProviderSpecificProperty's that don't match from every endpoint.
func (pf ProviderSpecificPropertyFilter) Filter(endpoints []*Endpoint) {
	for _, ep := range endpoints {
		for _, providerSpecific := range ep.ProviderSpecific {
			if !pf.Match(providerSpecific.Name) {
				log.WithFields(log.Fields{
					"dnsName": ep.DNSName,
					"targets": ep.Targets,
				}).Debugf("Provider specific property ignored by provider: %s", providerSpecific.Name)
				ep.DeleteProviderSpecificProperty(providerSpecific.Name)
			}
		}
	}
}
