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

	"github.com/kubernetes-incubator/external-dns/endpoint"
)

const (
	// The annotation used for figuring out which controller is responsible
	controllerAnnotationKey = "external-dns.kubernetes.io/controller"
	// The annotation used for defining the desired hostname
	hostnameAnnotationKey = "external-dns.kubernetes.io/hostname"
	// The value of the controller annotation so that we feel resposible
	controllerAnnotationValue = "dns-controller"
)

// ServiceSource is an implementation of Source for Kubernetes service objects.
// It will find all services that are under our jurisdiction, i.e. annotated
// desired hostname and matching or no controller annotation. For each of the
// matched services' external entrypoints it will return a corresponding
// Endpoint object.
type ServiceSource struct {
	Client    kubernetes.Interface
	Namespace string
}

// Endpoints returns endpoint objects for each service that should be processed.
func (sc *ServiceSource) Endpoints() ([]endpoint.Endpoint, error) {
	services, err := sc.Client.CoreV1().Services(sc.Namespace).List(v1.ListOptions{})
	if err != nil {
		return nil, err
	}

	endpoints := []endpoint.Endpoint{}

	for _, svc := range services.Items {
		svcEndpoints := endpointsFromService(&svc)
		if len(svcEndpoints) != 0 {
			endpoints = append(endpoints, svcEndpoints...)
		}
	}

	return endpoints, nil
}

// endpointsFromService extracts the endpoints from a service object
func endpointsFromService(svc *v1.Service) []endpoint.Endpoint {
	var endpoints []endpoint.Endpoint

	// Check controller annotation to see if we are responsible.
	controller, exists := svc.Annotations[controllerAnnotationKey]
	if exists && controller != controllerAnnotationValue {
		return endpoints
	}

	// Get the desired hostname of the service from the annotation.
	hostname, exists := svc.Annotations[hostnameAnnotationKey]
	if !exists {
		return endpoints
	}

	// Create a corresponding endpoint for each configured external entrypoint.
	for _, lb := range svc.Status.LoadBalancer.Ingress {
		if lb.IP != "" {
			endpoint := endpoint.Endpoint{
				DNSName: hostname,
				Target:  lb.IP,
			}
			endpoints = append(endpoints, endpoint)
		}
		if lb.Hostname != "" {
			endpoint := endpoint.Endpoint{
				DNSName: hostname,
				Target:  lb.Hostname,
			}
			endpoints = append(endpoints, endpoint)
		}
	}

	return endpoints
}
