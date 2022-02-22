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
	"strings"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"

	"sigs.k8s.io/external-dns/endpoint"
)

const (
	mateAnnotationKey     = "zalando.org/dnsname"
	moleculeAnnotationKey = "domainName"
	// kopsDNSControllerHostnameAnnotationKey is the annotation used for defining the desired hostname when kOps DNS controller compatibility mode
	kopsDNSControllerHostnameAnnotationKey = "dns.alpha.kubernetes.io/external"
	// kopsDNSControllerInternalHostnameAnnotationKey is the annotation used for defining the desired hostname when kOps DNS controller compatibility mode
	kopsDNSControllerInternalHostnameAnnotationKey = "dns.alpha.kubernetes.io/internal"
)

// legacyEndpointsFromService tries to retrieve Endpoints from Services
// annotated with legacy annotations.
func legacyEndpointsFromService(svc *v1.Service, sc *serviceSource) ([]*endpoint.Endpoint, error) {
	switch sc.compatibility {
	case "mate":
		return legacyEndpointsFromMateService(svc), nil
	case "molecule":
		return legacyEndpointsFromMoleculeService(svc), nil
	case "kops-dns-controller":
		return legacyEndpointsFromDNSControllerService(svc, sc)
	}

	return []*endpoint.Endpoint{}, nil
}

// legacyEndpointsFromMateService tries to retrieve Endpoints from Services
// annotated with Mate's annotation semantics.
func legacyEndpointsFromMateService(svc *v1.Service) []*endpoint.Endpoint {
	var endpoints []*endpoint.Endpoint

	// Get the desired hostname of the service from the annotation.
	hostname, exists := svc.Annotations[mateAnnotationKey]
	if !exists {
		return nil
	}

	// Create a corresponding endpoint for each configured external entrypoint.
	for _, lb := range svc.Status.LoadBalancer.Ingress {
		if lb.IP != "" {
			endpoints = append(endpoints, endpoint.NewEndpoint(hostname, endpoint.RecordTypeA, lb.IP))
		}
		if lb.Hostname != "" {
			endpoints = append(endpoints, endpoint.NewEndpoint(hostname, endpoint.RecordTypeCNAME, lb.Hostname))
		}
	}

	return endpoints
}

// legacyEndpointsFromMoleculeService tries to retrieve Endpoints from Services
// annotated with Molecule Software's annotation semantics.
func legacyEndpointsFromMoleculeService(svc *v1.Service) []*endpoint.Endpoint {
	var endpoints []*endpoint.Endpoint

	// Check that the Service opted-in to being processed.
	if svc.Labels["dns"] != "route53" {
		return nil
	}

	// Get the desired hostname of the service from the annotation.
	hostnameAnnotation, exists := svc.Annotations[moleculeAnnotationKey]
	if !exists {
		return nil
	}

	hostnameList := strings.Split(strings.Replace(hostnameAnnotation, " ", "", -1), ",")

	for _, hostname := range hostnameList {
		// Create a corresponding endpoint for each configured external entrypoint.
		for _, lb := range svc.Status.LoadBalancer.Ingress {
			if lb.IP != "" {
				endpoints = append(endpoints, endpoint.NewEndpoint(hostname, endpoint.RecordTypeA, lb.IP))
			}
			if lb.Hostname != "" {
				endpoints = append(endpoints, endpoint.NewEndpoint(hostname, endpoint.RecordTypeCNAME, lb.Hostname))
			}
		}
	}

	return endpoints
}

// legacyEndpointsFromDNSControllerService tries to retrieve Endpoints from Services
// annotated with DNS Controller's annotation semantics*.
func legacyEndpointsFromDNSControllerService(svc *v1.Service, sc *serviceSource) ([]*endpoint.Endpoint, error) {
	switch svc.Spec.Type {
	case v1.ServiceTypeNodePort:
		return legacyEndpointsFromDNSControllerNodePortService(svc, sc)
	case v1.ServiceTypeLoadBalancer:
		return legacyEndpointsFromDNSControllerLoadBalancerService(svc), nil
	}

	return []*endpoint.Endpoint{}, nil
}

// legacyEndpointsFromDNSControllerNodePortService implements DNS controller's semantics for NodePort services.
// It will use node role label to check if the node has the "node" role. This means control plane nodes and other
// roles will not be used as targets.
func legacyEndpointsFromDNSControllerNodePortService(svc *v1.Service, sc *serviceSource) ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint

	// Get the desired hostname of the service from the annotations.
	hostnameAnnotation, isExternal := svc.Annotations[kopsDNSControllerHostnameAnnotationKey]
	internalHostnameAnnotation, isInternal := svc.Annotations[kopsDNSControllerInternalHostnameAnnotationKey]

	if !isExternal && !isInternal {
		return nil, nil
	}

	// if both annotations are set, we just return empty, mimicking what dns-controller does
	if isInternal && isExternal {
		return nil, nil
	}

	for _, informer := range sc.informers {
		nodes, err := informer.nodeInformer.Lister().List(labels.Everything())
		if err != nil {
			return nil, err
		}

		var hostnameList []string
		if isExternal {
			hostnameList = strings.Split(strings.Replace(hostnameAnnotation, " ", "", -1), ",")
		} else {
			hostnameList = strings.Split(strings.Replace(internalHostnameAnnotation, " ", "", -1), ",")
		}

		for _, hostname := range hostnameList {
			for _, node := range nodes {
				_, isNode := node.Labels["node-role.kubernetes.io/node"]
				if !isNode {
					continue
				}
				for _, address := range node.Status.Addresses {
					if address.Type == v1.NodeExternalIP && isExternal {
						endpoints = append(endpoints, endpoint.NewEndpoint(hostname, endpoint.RecordTypeA, address.Address))
					}
					if address.Type == v1.NodeInternalIP && isInternal {
						endpoints = append(endpoints, endpoint.NewEndpoint(hostname, endpoint.RecordTypeA, address.Address))
					}
				}
			}
		}
	}
	return endpoints, nil
}

// legacyEndpointsFromDNSControllerLoadBalancerService will respect both annotations, but
// will not care if the load balancer actually is internal or not.
func legacyEndpointsFromDNSControllerLoadBalancerService(svc *v1.Service) []*endpoint.Endpoint {
	var endpoints []*endpoint.Endpoint

	// Get the desired hostname of the service from the annotations.
	hostnameAnnotation, hasExternal := svc.Annotations[kopsDNSControllerHostnameAnnotationKey]
	internalHostnameAnnotation, hasInternal := svc.Annotations[kopsDNSControllerInternalHostnameAnnotationKey]

	if !hasExternal && !hasInternal {
		return nil
	}

	var hostnameList []string
	if hasExternal {
		hostnameList = append(hostnameList, strings.Split(strings.Replace(hostnameAnnotation, " ", "", -1), ",")...)
	}
	if hasInternal {
		hostnameList = append(hostnameList, strings.Split(strings.Replace(internalHostnameAnnotation, " ", "", -1), ",")...)
	}

	for _, hostname := range hostnameList {
		// Create a corresponding endpoint for each configured external entrypoint.
		for _, lb := range svc.Status.LoadBalancer.Ingress {
			if lb.IP != "" {
				endpoints = append(endpoints, endpoint.NewEndpoint(hostname, endpoint.RecordTypeA, lb.IP))
			}
			if lb.Hostname != "" {
				endpoints = append(endpoints, endpoint.NewEndpoint(hostname, endpoint.RecordTypeCNAME, lb.Hostname))
			}
		}
	}

	return endpoints
}
