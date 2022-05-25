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

// NewGatewayHTTPRouteSource creates a new Gateway HTTPRoute source with the given config.
func NewGatewayHTTPRouteSource(clients ClientGenerator, config *Config) (Source, error) {
	return newGatewayRouteSource(clients, config, "HTTPRoute", func(factory informers.SharedInformerFactory) gatewayRouteInformer {
		return &gatewayHTTPRouteInformer{factory.Gateway().V1alpha2().HTTPRoutes()}
	})
}

type gatewayHTTPRoute struct{ route *v1alpha2.HTTPRoute }

func (rt *gatewayHTTPRoute) Object() kubeObject                { return rt.route }
func (rt *gatewayHTTPRoute) Metadata() *metav1.ObjectMeta      { return &rt.route.ObjectMeta }
func (rt *gatewayHTTPRoute) Hostnames() []v1alpha2.Hostname    { return rt.route.Spec.Hostnames }
func (rt *gatewayHTTPRoute) Protocol() v1alpha2.ProtocolType   { return v1alpha2.HTTPProtocolType }
func (rt *gatewayHTTPRoute) RouteStatus() v1alpha2.RouteStatus { return rt.route.Status.RouteStatus }

type gatewayHTTPRouteInformer struct {
	informers_v1a2.HTTPRouteInformer
}

func (inf gatewayHTTPRouteInformer) List(namespace string, selector labels.Selector) ([]gatewayRoute, error) {
	list, err := inf.HTTPRouteInformer.Lister().HTTPRoutes(namespace).List(selector)
	if err != nil {
		return nil, err
	}
	routes := make([]gatewayRoute, len(list))
	for i, rt := range list {
		routes[i] = &gatewayHTTPRoute{rt}
	}
	return routes, nil
}
