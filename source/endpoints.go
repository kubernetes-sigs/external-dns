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

	log "github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/labels"
	coreinformers "k8s.io/client-go/informers/core/v1"
	gwinformers_v1 "sigs.k8s.io/gateway-api/pkg/client/informers/externalversions/apis/v1"

	"sigs.k8s.io/external-dns/endpoint"
)

// EndpointTargetsFromServices retrieves endpoint targets from services in a given namespace
// that match the specified selector. It returns external IPs or load balancer addresses.
//
// TODO: add support for service.Spec.Ports (type NodePort) and service.Spec.ClusterIPs (type ClusterIP)
func EndpointTargetsFromServices(svcInformer coreinformers.ServiceInformer, namespace string, selector map[string]string) (endpoint.Targets, error) {
	targets := endpoint.Targets{}

	services, err := svcInformer.Lister().Services(namespace).List(labels.Everything())

	if err != nil {
		return nil, fmt.Errorf("failed to list labels for services in namespace %q: %w", namespace, err)
	}

	labelsSelector := labels.SelectorFromSet(selector)
	for _, service := range services {
		if !labelsSelector.Matches(labels.Set(service.Spec.Selector)) {
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
	return endpoint.NewTargets(targets...), nil
}

// EndpointTargetsFromK8sGateway resolves endpoint targets from a Kubernetes
// Gateway API Gateway object identified by a "namespace/name" or "name" reference.
// defaultNamespace is used when the reference omits a namespace.
func EndpointTargetsFromK8sGateway(gwInformer gwinformers_v1.GatewayInformer, ref string, defaultNamespace string) (endpoint.Targets, error) {
	if gwInformer == nil {
		return nil, fmt.Errorf("Gateway API client not configured but %q annotation is set", K8sGatewaySource)
	}

	namespace, name, err := ParseNamespacedName(ref)
	if err != nil {
		return nil, err
	}
	if namespace == "" {
		namespace = defaultNamespace
	}

	gw, err := gwInformer.Lister().Gateways(namespace).Get(name)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	targets := make(endpoint.Targets, 0, len(gw.Status.Addresses))
	for _, addr := range gw.Status.Addresses {
		if addr.Value != "" {
			targets = append(targets, addr.Value)
		}
	}
	return targets, nil
}
