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

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	coreinformers "k8s.io/client-go/informers/core/v1"

	"sigs.k8s.io/external-dns/endpoint"
)

// EndpointTargetsFromServices retrieves endpoint targets from services in a given namespace
// that match the specified selector.
//
// Target selection is per service, first non-empty wins: external IPs, load balancer
// ingress addresses, then cluster IPs. The cluster IP fallback covers ClusterIP-type
// services and NodePort-type services (which are also allocated a cluster IP).
// Headless services (cluster IP "None") yield no targets, and NodePort node addresses
// stay out of scope: resolving them needs a cluster-scoped node informer while these
// sources are namespace-scoped.
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

		lbTargets := endpoint.Targets{}
		for _, lb := range service.Status.LoadBalancer.Ingress {
			if lb.IP != "" {
				lbTargets = append(lbTargets, lb.IP)
			} else if lb.Hostname != "" {
				lbTargets = append(lbTargets, lb.Hostname)
			}
		}
		if len(lbTargets) > 0 {
			targets = append(targets, lbTargets...)
			continue
		}

		for _, clusterIP := range service.Spec.ClusterIPs {
			if clusterIP != "" && clusterIP != corev1.ClusterIPNone {
				targets = append(targets, clusterIP)
			}
		}
	}
	return endpoint.NewTargets(targets...), nil
}
