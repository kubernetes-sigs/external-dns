/*
Copyright 2021 The Kubernetes Authors.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"sigs.k8s.io/gateway-api/apis/v1alpha2"
	"sigs.k8s.io/gateway-api/apis/v1beta1"
	informers "sigs.k8s.io/gateway-api/pkg/client/informers/externalversions"
	informers_v1a2 "sigs.k8s.io/gateway-api/pkg/client/informers/externalversions/apis/v1alpha2"
)

// NewGatewayUDPRouteSource creates a new Gateway UDPRoute source with the given config.
func NewGatewayUDPRouteSource(clients ClientGenerator, config *Config) (Source, error) {
	return newGatewayRouteSource(clients, config, "UDPRoute", func(factory informers.SharedInformerFactory) gatewayRouteInformer {
		return &gatewayUDPRouteInformer{factory.Gateway().V1alpha2().UDPRoutes()}
	})
}

type gatewayUDPRoute struct{ route *v1alpha2.UDPRoute }

func (rt *gatewayUDPRoute) Object() kubeObject             { return rt.route }
func (rt *gatewayUDPRoute) Metadata() *metav1.ObjectMeta   { return &rt.route.ObjectMeta }
func (rt *gatewayUDPRoute) Hostnames() []v1beta1.Hostname  { return nil }
func (rt *gatewayUDPRoute) Protocol() v1beta1.ProtocolType { return v1beta1.UDPProtocolType }
func (rt *gatewayUDPRoute) RouteStatus() v1beta1.RouteStatus {
	return v1b1RouteStatus(rt.route.Status.RouteStatus)
}

type gatewayUDPRouteInformer struct {
	informers_v1a2.UDPRouteInformer
}

func (inf gatewayUDPRouteInformer) List(namespace string, selector labels.Selector) ([]gatewayRoute, error) {
	list, err := inf.UDPRouteInformer.Lister().UDPRoutes(namespace).List(selector)
	if err != nil {
		return nil, err
	}
	routes := make([]gatewayRoute, len(list))
	for i, rt := range list {
		routes[i] = &gatewayUDPRoute{rt}
	}
	return routes, nil
}
