/*
Copyright 2022 The Kubernetes Authors.

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
	v1 "sigs.k8s.io/gateway-api/apis/v1"
	informers "sigs.k8s.io/gateway-api/pkg/client/informers/externalversions"
	informers_v1 "sigs.k8s.io/gateway-api/pkg/client/informers/externalversions/apis/v1"
)

// NewGatewayGRPCRouteSource creates a new Gateway GRPCRoute source with the given config.
func NewGatewayGRPCRouteSource(clients ClientGenerator, config *Config) (Source, error) {
	return newGatewayRouteSource(clients, config, "GRPCRoute", func(factory informers.SharedInformerFactory) gatewayRouteInformer {
		return &gatewayGRPCRouteInformer{factory.Gateway().V1().GRPCRoutes()}
	})
}

type gatewayGRPCRoute struct{ route v1.GRPCRoute } // NOTE: Must update TypeMeta in List when changing the APIVersion.

func (rt *gatewayGRPCRoute) Object() kubeObject           { return &rt.route }
func (rt *gatewayGRPCRoute) Metadata() *metav1.ObjectMeta { return &rt.route.ObjectMeta }
func (rt *gatewayGRPCRoute) Hostnames() []v1.Hostname     { return rt.route.Spec.Hostnames }
func (rt *gatewayGRPCRoute) Protocol() v1.ProtocolType    { return v1.HTTPSProtocolType }
func (rt *gatewayGRPCRoute) RouteStatus() v1.RouteStatus  { return rt.route.Status.RouteStatus }

type gatewayGRPCRouteInformer struct {
	informers_v1.GRPCRouteInformer
}

func (inf gatewayGRPCRouteInformer) List(namespace string, selector labels.Selector) ([]gatewayRoute, error) {
	list, err := inf.GRPCRouteInformer.Lister().GRPCRoutes(namespace).List(selector)
	if err != nil {
		return nil, err
	}
	routes := make([]gatewayRoute, len(list))
	for i, rt := range list {
		// List results are supposed to be treated as read-only.
		// We make a shallow copy since we're only interested in setting the TypeMeta.
		clone := *rt
		clone.TypeMeta = metav1.TypeMeta{
			APIVersion: v1.GroupVersion.String(),
			Kind:       "GRPCRoute",
		}
		routes[i] = &gatewayGRPCRoute{clone}
	}
	return routes, nil
}
