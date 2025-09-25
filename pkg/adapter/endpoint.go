/*
Copyright 2025 The Kubernetes Authors.

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
package adapter

import (
	apiv1alpha1 "sigs.k8s.io/external-dns/apis/v1alpha1"
	"sigs.k8s.io/external-dns/endpoint"
)

// FromAPIEndpoint converts an API endpoint (v1alpha1.Endpoint) to an internal endpoint (endpoint.Endpoint).
func FromAPIEndpoint(apiEndpoint *apiv1alpha1.Endpoint) *endpoint.Endpoint {
	if apiEndpoint == nil {
		return nil
	}
	p := make(endpoint.ProviderSpecific)
	for _, ps := range apiEndpoint.ProviderSpecific {
		p.Set(ps.Name, ps.Value)
	}
	return &endpoint.Endpoint{
		DNSName:          apiEndpoint.DNSName,
		Targets:          apiEndpoint.Targets,
		RecordType:       apiEndpoint.RecordType,
		SetIdentifier:    apiEndpoint.SetIdentifier,
		RecordTTL:        endpoint.TTL(apiEndpoint.RecordTTL),
		Labels:           apiEndpoint.Labels,
		ProviderSpecific: p,
	}
}

// FromAPIEndpoints converts a slice of API endpoints to internal endpoints.
func FromAPIEndpoints(apiEndpoints []*apiv1alpha1.Endpoint) []*endpoint.Endpoint {
	if len(apiEndpoints) == 0 {
		return nil
	}
	eps := make([]*endpoint.Endpoint, 0, len(apiEndpoints))
	for _, apiEndpoint := range apiEndpoints {
		eps = append(eps, FromAPIEndpoint(apiEndpoint))
	}
	return eps
}

// ToAPIEndpoint converts an internal endpoint (endpoint.Endpoint) to an API endpoint (v1alpha1.Endpoint).
func ToAPIEndpoint(internalEndpoint *endpoint.Endpoint) *apiv1alpha1.Endpoint {
	if internalEndpoint == nil {
		return nil
	}
	ps := make([]apiv1alpha1.ProviderSpecificProperty, 0, len(internalEndpoint.ProviderSpecific))
	for k, v := range internalEndpoint.ProviderSpecific {
		ps = append(ps, apiv1alpha1.ProviderSpecificProperty{
			Name:  k,
			Value: v,
		})
	}
	return &apiv1alpha1.Endpoint{
		DNSName:          internalEndpoint.DNSName,
		Targets:          internalEndpoint.Targets,
		RecordType:       internalEndpoint.RecordType,
		SetIdentifier:    internalEndpoint.SetIdentifier,
		RecordTTL:        int64(internalEndpoint.RecordTTL),
		Labels:           internalEndpoint.Labels,
		ProviderSpecific: ps,
	}
}

// ToAPIEndpoints converts a slice of internal endpoints to API endpoints.
func ToAPIEndpoints(internalEndpoints []*endpoint.Endpoint) []*apiv1alpha1.Endpoint {
	if len(internalEndpoints) == 0 {
		return nil
	}
	apiEndpoints := make([]*apiv1alpha1.Endpoint, 0, len(internalEndpoints))
	for _, internalEndpoint := range internalEndpoints {
		apiEndpoints = append(apiEndpoints, ToAPIEndpoint(internalEndpoint))
	}
	return apiEndpoints
}
