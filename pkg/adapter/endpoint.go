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
package adapter

import (
	apiv1alpha1 "sigs.k8s.io/external-dns/apis/v1alpha1"
	"sigs.k8s.io/external-dns/endpoint"
)

func ToInternalEndpoint(crdEp *apiv1alpha1.Endpoint) *endpoint.Endpoint {
	if crdEp == nil {
		return nil
	}
	p := make(endpoint.ProviderSpecific)
	for _, ps := range crdEp.ProviderSpecific {
		p.Set(ps.Name, ps.Value)
	}
	ep := &endpoint.Endpoint{
		DNSName:          crdEp.DNSName,
		Targets:          crdEp.Targets,
		RecordType:       crdEp.RecordType,
		SetIdentifier:    crdEp.SetIdentifier,
		RecordTTL:        endpoint.TTL(crdEp.RecordTTL),
		Labels:           crdEp.Labels,
		ProviderSpecific: p,
	}
	return ep
}

func ToInternalEndpoints(crdEps []*apiv1alpha1.Endpoint) []*endpoint.Endpoint {
	if len(crdEps) == 0 {
		return nil
	}
	eps := make([]*endpoint.Endpoint, 0, len(crdEps))
	for _, crdEp := range crdEps {
		eps = append(eps, ToInternalEndpoint(crdEp))
	}
	return eps
}

func ToAPIEndpoint(ep *endpoint.Endpoint) *apiv1alpha1.Endpoint {
	if ep == nil {
		return nil
	}
	ps := make([]apiv1alpha1.ProviderSpecificProperty, 0, len(ep.ProviderSpecific))
	for k, v := range ep.ProviderSpecific {
		ps = append(ps, apiv1alpha1.ProviderSpecificProperty{
			Name:  k,
			Value: v,
		})
	}
	crdEp := &apiv1alpha1.Endpoint{
		DNSName:          ep.DNSName,
		Targets:          ep.Targets,
		RecordType:       ep.RecordType,
		SetIdentifier:    ep.SetIdentifier,
		RecordTTL:        int64(ep.RecordTTL),
		Labels:           ep.Labels,
		ProviderSpecific: ps,
	}
	return crdEp
}

func ToAPIEndpoints(eps []*endpoint.Endpoint) []*apiv1alpha1.Endpoint {
	if len(eps) == 0 {
		return nil
	}
	crdEps := make([]*apiv1alpha1.Endpoint, 0, len(eps))
	for _, ep := range eps {
		crdEps = append(crdEps, ToAPIEndpoint(ep))
	}
	return crdEps
}
