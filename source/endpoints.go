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
	"strconv"

	"k8s.io/apimachinery/pkg/labels"
	coreinformers "k8s.io/client-go/informers/core/v1"

	"sigs.k8s.io/external-dns/endpoint"
)

// EndpointTargetsFromServices retrieves endpoint targets from services in a given namespace
// that match the specified selector. It returns external IPs, internal service addresses,
// and load balancer addresses that can be used for DNS records.
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

		targets = append(targets, service.Spec.ExternalIPs...)

		if service.Spec.ClusterIP != "" && service.Spec.ClusterIP != "None" {
			targets = append(targets, service.Spec.ClusterIP)
		}
		for _, clusterIP := range service.Spec.ClusterIPs {
			if clusterIP != "" && clusterIP != "None" {
				targets = append(targets, clusterIP)
			}
		}
		for _, port := range service.Spec.Ports {
			if port.NodePort != 0 {
				targets = append(targets, strconv.Itoa(int(port.NodePort)))
			}
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
