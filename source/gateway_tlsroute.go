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
	informers "sigs.k8s.io/gateway-api/pkg/client/informers/gateway/externalversions"
	informers_v1a2 "sigs.k8s.io/gateway-api/pkg/client/informers/gateway/externalversions/apis/v1alpha2"
)

// NewGatewayTLSRouteSource creates a new Gateway TLSRoute source with the given config.
func NewGatewayTLSRouteSource(clients ClientGenerator, config *Config) (Source, error) {
	return newGatewayRouteSource(clients, config, "TLSRoute", func(factory informers.SharedInformerFactory) gatewayRouteInformer {
		return &gatewayTLSRouteInformer{factory.Gateway().V1alpha2().TLSRoutes()}
	})
}

type gatewayTLSRoute struct{ route *v1alpha2.TLSRoute }

func (rt *gatewayTLSRoute) Object() kubeObject                { return rt.route }
func (rt *gatewayTLSRoute) Metadata() *metav1.ObjectMeta      { return &rt.route.ObjectMeta }
func (rt *gatewayTLSRoute) Hostnames() []v1alpha2.Hostname    { return rt.route.Spec.Hostnames }
func (rt *gatewayTLSRoute) Protocol() v1alpha2.ProtocolType   { return v1alpha2.TLSProtocolType }
func (rt *gatewayTLSRoute) RouteStatus() v1alpha2.RouteStatus { return rt.route.Status.RouteStatus }

type gatewayTLSRouteInformer struct {
	informers_v1a2.TLSRouteInformer
}

func (inf gatewayTLSRouteInformer) List(namespace string, selector labels.Selector) ([]gatewayRoute, error) {
	list, err := inf.TLSRouteInformer.Lister().TLSRoutes(namespace).List(selector)
	if err != nil {
		return nil, err
	}
	routes := make([]gatewayRoute, len(list))
	for i, rt := range list {
		routes[i] = &gatewayTLSRoute{rt}
	}
	return routes, nil
}
