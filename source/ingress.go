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

package source

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pkg/api/v1"
	"k8s.io/client-go/pkg/apis/extensions/v1beta1"

	"github.com/kubernetes-incubator/external-dns/endpoint"
)

// IngressSource is an implementation of Source for Kubernetes ingress objects.
// Ingress implementation will use the spec.rules.host value for the hostname
// Ingress annotations are ignored
type IngressSource struct {
	Client kubernetes.Interface
}

// Endpoints returns endpoint objects for each host-target combination that should be processed.
// Retrieves all ingress resources on all namespaces
func (sc *IngressSource) Endpoints() ([]endpoint.Endpoint, error) {
	ingresses, err := sc.Client.Extensions().Ingresses(v1.NamespaceAll).List(v1.ListOptions{})
	if err != nil {
		return nil, err
	}

	endpoints := []endpoint.Endpoint{}

	for _, ing := range ingresses.Items {
		ingEndpoints := endpointsFromIngress(&ing)
		endpoints = append(endpoints, ingEndpoints...)
	}

	return endpoints, nil
}

// endpointsFromIngress extracts the endpoints from ingress object
func endpointsFromIngress(ing *v1beta1.Ingress) []endpoint.Endpoint {
	var endpoints []endpoint.Endpoint

	for _, rule := range ing.Spec.Rules {
		if rule.Host == "" {
			continue
		}
		for _, lb := range ing.Status.LoadBalancer.Ingress {
			endpoint := endpoint.Endpoint{
				DNSName: rule.Host,
			}
			if lb.IP != "" {
				endpoint.Target = lb.IP
				endpoints = append(endpoints, endpoint)
			}
			if lb.Hostname != "" {
				endpoint.Target = lb.Hostname
				endpoints = append(endpoints, endpoint)
			}
		}
	}

	return endpoints
}
