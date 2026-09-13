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
	"context"
	"errors"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "sigs.k8s.io/gateway-api/apis/v1"
	"sigs.k8s.io/gateway-api/apis/v1alpha2"
	gwinformers "sigs.k8s.io/gateway-api/pkg/client/informers/externalversions"
	informers_v1 "sigs.k8s.io/gateway-api/pkg/client/informers/externalversions/apis/v1"
	informers_v1a2 "sigs.k8s.io/gateway-api/pkg/client/informers/externalversions/apis/v1alpha2"

	extInformers "sigs.k8s.io/external-dns/source/informers"
)

// NewGatewayUDPRouteSource creates a new Gateway UDPRoute source with the given config.
func NewGatewayUDPRouteSource(ctx context.Context, clients ClientGenerator, config *Config) (Source, error) {
	kubeClient, err := clients.KubeClient()
	if err != nil {
		return nil, err
	}

	switch {
	case apiVersionResourceAvailable(
		kubeClient,
		v1.GroupVersion.String(),
		"udproutes",
	):
		return newGatewayRouteSource(ctx, clients, config, "UDPRoute",
			func(factory gwinformers.SharedInformerFactory) gatewayRouteInformer {
				return &gatewayUDPRouteV1Informer{
					factory.Gateway().V1().UDPRoutes(),
				}
			})

	case apiVersionResourceAvailable(
		kubeClient,
		v1alpha2.GroupVersion.String(),
		"udproutes",
	):
		return newGatewayRouteSource(ctx, clients, config, "UDPRoute",
			func(factory gwinformers.SharedInformerFactory) gatewayRouteInformer {
				return &gatewayUDPRouteV1alpha2Informer{
					factory.Gateway().V1alpha2().UDPRoutes(),
				}
			})
	}

	return nil, errors.New("no supported Gateway API UDPRoute version available")
}

type gatewayUDPRouteV1 struct {
	route v1.UDPRoute
}

func (rt *gatewayUDPRouteV1) Object() kubeObject               { return &rt.route }
func (rt *gatewayUDPRouteV1) Metadata() *metav1.ObjectMeta     { return &rt.route.ObjectMeta }
func (rt *gatewayUDPRouteV1) Hostnames() []v1.Hostname         { return nil }
func (rt *gatewayUDPRouteV1) ParentRefs() []v1.ParentReference { return rt.route.Spec.ParentRefs }
func (rt *gatewayUDPRouteV1) Protocol() v1.ProtocolType        { return v1.UDPProtocolType }
func (rt *gatewayUDPRouteV1) RouteStatus() v1.RouteStatus      { return rt.route.Status.RouteStatus }

type gatewayUDPRouteV1Informer struct {
	informers_v1.UDPRouteInformer
}

func (inf gatewayUDPRouteV1Informer) List() []gatewayRoute {
	list := extInformers.ListIndexed[*v1.UDPRoute](
		inf.UDPRouteInformer.Informer().GetIndexer(),
	)

	routes := make([]gatewayRoute, len(list))
	for i, rt := range list {
		// List results are supposed to be treated as read-only.
		// We make a shallow copy since we're only interested in setting the TypeMeta.
		clone := *rt
		clone.TypeMeta = metav1.TypeMeta{
			APIVersion: v1.GroupVersion.String(),
			Kind:       "UDPRoute",
		}
		routes[i] = &gatewayUDPRouteV1{clone}
	}

	return routes
}

type gatewayUDPRouteV1alpha2 struct {
	route v1alpha2.UDPRoute
}

func (rt *gatewayUDPRouteV1alpha2) Object() kubeObject               { return &rt.route }
func (rt *gatewayUDPRouteV1alpha2) Metadata() *metav1.ObjectMeta     { return &rt.route.ObjectMeta }
func (rt *gatewayUDPRouteV1alpha2) Hostnames() []v1.Hostname         { return nil }
func (rt *gatewayUDPRouteV1alpha2) ParentRefs() []v1.ParentReference { return rt.route.Spec.ParentRefs }
func (rt *gatewayUDPRouteV1alpha2) Protocol() v1.ProtocolType        { return v1.UDPProtocolType }
func (rt *gatewayUDPRouteV1alpha2) RouteStatus() v1.RouteStatus      { return rt.route.Status.RouteStatus }

type gatewayUDPRouteV1alpha2Informer struct {
	informers_v1a2.UDPRouteInformer
}

func (inf gatewayUDPRouteV1alpha2Informer) List() []gatewayRoute {
	list := extInformers.ListIndexed[*v1alpha2.UDPRoute](
		inf.UDPRouteInformer.Informer().GetIndexer(),
	)

	routes := make([]gatewayRoute, len(list))
	for i, rt := range list {
		// We make a shallow copy since we're only interested in setting the TypeMeta.
		clone := *rt
		clone.TypeMeta = metav1.TypeMeta{
			APIVersion: v1alpha2.GroupVersion.String(),
			Kind:       "UDPRoute",
		}
		routes[i] = &gatewayUDPRouteV1alpha2{clone}
	}

	return routes
}
