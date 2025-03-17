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

package utils

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/labels"
	coreinformers "k8s.io/client-go/informers/core/v1"

	"sigs.k8s.io/external-dns/endpoint"
)

func EndpointsForHostname(hostname string, targets endpoint.Targets, ttl endpoint.TTL, providerSpecific endpoint.ProviderSpecific, setIdentifier string, resource string) []*endpoint.Endpoint {
	var endpoints []*endpoint.Endpoint

	var aTargets endpoint.Targets
	var aaaaTargets endpoint.Targets
	var cnameTargets endpoint.Targets

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

func ProviderSpecificAnnotations(annotations map[string]string) (endpoint.ProviderSpecific, string) {
	providerSpecificAnnotations := endpoint.ProviderSpecific{}

	if v, exists := annotations[CloudflareProxiedKey]; exists {
		providerSpecificAnnotations = append(providerSpecificAnnotations, endpoint.ProviderSpecificProperty{
			Name:  CloudflareProxiedKey,
			Value: v,
		})
	}
	if v, exists := annotations[CloudflareCustomHostnameKey]; exists {
		providerSpecificAnnotations = append(providerSpecificAnnotations, endpoint.ProviderSpecificProperty{
			Name:  CloudflareCustomHostnameKey,
			Value: v,
		})
	}
	if getAliasFromAnnotations(annotations) {
		providerSpecificAnnotations = append(providerSpecificAnnotations, endpoint.ProviderSpecificProperty{
			Name:  "alias",
			Value: "true",
		})
	}
	setIdentifier := ""
	for k, v := range annotations {
		if k == SetIdentifierKey {
			setIdentifier = v
		} else if strings.HasPrefix(k, "external-dns.alpha.kubernetes.io/aws-") {
			attr := strings.TrimPrefix(k, "external-dns.alpha.kubernetes.io/aws-")
			providerSpecificAnnotations = append(providerSpecificAnnotations, endpoint.ProviderSpecificProperty{
				Name:  fmt.Sprintf("aws/%s", attr),
				Value: v,
			})
		} else if strings.HasPrefix(k, "external-dns.alpha.kubernetes.io/scw-") {
			attr := strings.TrimPrefix(k, "external-dns.alpha.kubernetes.io/scw-")
			providerSpecificAnnotations = append(providerSpecificAnnotations, endpoint.ProviderSpecificProperty{
				Name:  fmt.Sprintf("scw/%s", attr),
				Value: v,
			})
		} else if strings.HasPrefix(k, "external-dns.alpha.kubernetes.io/ibmcloud-") {
			attr := strings.TrimPrefix(k, "external-dns.alpha.kubernetes.io/ibmcloud-")
			providerSpecificAnnotations = append(providerSpecificAnnotations, endpoint.ProviderSpecificProperty{
				Name:  fmt.Sprintf("ibmcloud-%s", attr),
				Value: v,
			})
		} else if strings.HasPrefix(k, "external-dns.alpha.kubernetes.io/webhook-") {
			// Support for wildcard annotations for webhook providers
			attr := strings.TrimPrefix(k, "external-dns.alpha.kubernetes.io/webhook-")
			providerSpecificAnnotations = append(providerSpecificAnnotations, endpoint.ProviderSpecificProperty{
				Name:  fmt.Sprintf("webhook/%s", attr),
				Value: v,
			})
		}
	}
	return providerSpecificAnnotations, setIdentifier
}

func EndpointTargetsFromServices(svcInformer coreinformers.ServiceInformer, namespace string, selector map[string]string) (endpoint.Targets, error) {
	targets := endpoint.Targets{}

	services, err := svcInformer.Lister().Services(namespace).List(labels.Everything())
	if err != nil {
		log.Errorf("not able to list labels for services in namespace %s. %v", namespace, err)
		return nil, err
	}

	for _, service := range services {
		if !SelectorMatchesServiceSelector(selector, service.Spec.Selector) {
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
