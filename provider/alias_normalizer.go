/*
Copyright 2026 The Kubernetes Authors.

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
	"sigs.k8s.io/external-dns/endpoint"
)

// AliasNormalizingMiddleware wraps a Provider and normalizes alias ProviderSpecific
// values after AdjustEndpoints. Providers may convert CNAME endpoints to A/AAAA
// alias records but leave the alias property as "A" or "AAAA", while Records()
// always returns "true" for alias records. This mismatch causes the plan to
// generate a spurious update on every reconciliation loop.
type AliasNormalizingMiddleware struct {
	Provider
}

func NewAliasNormalizingMiddleware(p Provider) *AliasNormalizingMiddleware {
	return &AliasNormalizingMiddleware{Provider: p}
}

// AdjustEndpoints delegates to the inner provider then normalizes alias values
// so they match what Records() returns.
func (p *AliasNormalizingMiddleware) AdjustEndpoints(endpoints []*endpoint.Endpoint) ([]*endpoint.Endpoint, error) {
	eps, err := p.Provider.AdjustEndpoints(endpoints)
	if err != nil {
		return nil, err
	}
	for _, ep := range eps {
		alias := ep.GetAliasProperty()
		if (ep.RecordType == endpoint.RecordTypeA && alias == endpoint.AliasA) ||
			(ep.RecordType == endpoint.RecordTypeAAAA && alias == endpoint.AliasAAAA) {
			ep.SetProviderSpecificProperty(endpoint.ProviderSpecificAlias, string(endpoint.AliasTrue))
		}
	}
	return eps, nil
}
