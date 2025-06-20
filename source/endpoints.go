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

package source

import (
	"fmt"

	"k8s.io/apimachinery/pkg/labels"
	coreinformers "k8s.io/client-go/informers/core/v1"

	"sigs.k8s.io/external-dns/endpoint"
)

// endpointsForHostname returns the endpoint objects for each host-target combination.
func endpointsForHostname(hostname string, targets endpoint.Targets, ttl endpoint.TTL, providerSpecific endpoint.ProviderSpecific, setIdentifier string, resource string) []*endpoint.Endpoint {
	var (
		endpoints    []*endpoint.Endpoint
		aTargets     endpoint.Targets
		aaaaTargets  endpoint.Targets
		cnameTargets endpoint.Targets
	)

	for _, t := range targets {
		switch suitableType(t) {
		case endpoint.RecordTypeA:
			aTargets = append(aTargets, t)
		case endpoint.RecordTypeAAAA:
			aaaaTargets = append(aaaaTargets, t)
		default:
			cnameTargets = append(cnameTargets, t)
		}
	}

	if len(aTargets) > 0 {
		epA := endpoint.NewEndpointWithTTL(hostname, endpoint.RecordTypeA, ttl, aTargets...)
		if epA != nil {
			epA.ProviderSpecific = providerSpecific
			epA.SetIdentifier = setIdentifier
			if resource != "" {
				epA.Labels[endpoint.ResourceLabelKey] = resource
			}
			endpoints = append(endpoints, epA)
		}
	}

	if len(aaaaTargets) > 0 {
		epAAAA := endpoint.NewEndpointWithTTL(hostname, endpoint.RecordTypeAAAA, ttl, aaaaTargets...)
		if epAAAA != nil {
			epAAAA.ProviderSpecific = providerSpecific
			epAAAA.SetIdentifier = setIdentifier
			if resource != "" {
				epAAAA.Labels[endpoint.ResourceLabelKey] = resource
			}
			endpoints = append(endpoints, epAAAA)
		}
	}

	if len(cnameTargets) > 0 {
		epCNAME := endpoint.NewEndpointWithTTL(hostname, endpoint.RecordTypeCNAME, ttl, cnameTargets...)
		if epCNAME != nil {
			epCNAME.ProviderSpecific = providerSpecific
			epCNAME.SetIdentifier = setIdentifier
			if resource != "" {
				epCNAME.Labels[endpoint.ResourceLabelKey] = resource
			}
			endpoints = append(endpoints, epCNAME)
		}
	}

	return endpoints
}

func EndpointTargetsFromServices(svcInformer coreinformers.ServiceInformer, namespace string, selector map[string]string) (endpoint.Targets, error) {
	targets := endpoint.Targets{}

	services, err := svcInformer.Lister().Services(namespace).List(labels.Everything())

	if err != nil {
		return nil, fmt.Errorf("failed to list labels for services in namespace %q: %w", namespace, err)
	}

	for _, service := range services {
		if !MatchesServiceSelector(selector, service.Spec.Selector) {
			continue
		}

		if len(service.Spec.ExternalIPs) > 0 {
			targets = append(targets, service.Spec.ExternalIPs...)
			continue
		}

		for _, lb := range service.Status.LoadBalancer.Ingress {
			if lb.IP != "" {
				targets = append(targets, lb.IP)
			} else if lb.Hostname != "" {
				targets = append(targets, lb.Hostname)
			}
		}
	}
	return targets, nil
}
