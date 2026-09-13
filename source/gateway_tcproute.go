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
	"k8s.io/client-go/kubernetes"
	v1 "sigs.k8s.io/gateway-api/apis/v1"
	"sigs.k8s.io/gateway-api/apis/v1alpha2"
	gwinformers "sigs.k8s.io/gateway-api/pkg/client/informers/externalversions"
	informers_v1 "sigs.k8s.io/gateway-api/pkg/client/informers/externalversions/apis/v1"
	informers_v1a2 "sigs.k8s.io/gateway-api/pkg/client/informers/externalversions/apis/v1alpha2"

	extInformers "sigs.k8s.io/external-dns/source/informers"
)

func apiVersionResourceAvailable(
	client kubernetes.Interface,
	apiVersion string,
	resourceName string,
) bool {
	resources, err := client.Discovery().ServerResourcesForGroupVersion(apiVersion)
	if err != nil {
		return false
	}

	for _, resource := range resources.APIResources {
		if resource.Name == resourceName {
			return true
		}
	}

	return false
}

func NewGatewayTCPRouteSource(ctx context.Context, clients ClientGenerator, config *Config) (Source, error) {
	kubeClient, err := clients.KubeClient()
	if err != nil {
		return nil, err
	}

	switch {
	case apiVersionResourceAvailable(
		kubeClient,
		v1.GroupVersion.String(),
		"tcproutes",
	):
		return newGatewayRouteSource(ctx, clients, config, "TCPRoute",
			func(factory gwinformers.SharedInformerFactory) gatewayRouteInformer {
				return &gatewayTCPRouteV1Informer{
					factory.Gateway().V1().TCPRoutes(),
				}
			})

	case apiVersionResourceAvailable(
		kubeClient,
		v1alpha2.GroupVersion.String(),
		"tcproutes",
	):
		return newGatewayRouteSource(ctx, clients, config, "TCPRoute",
			func(factory gwinformers.SharedInformerFactory) gatewayRouteInformer {
				return &gatewayTCPRouteV1alpha2Informer{
					factory.Gateway().V1alpha2().TCPRoutes(),
				}
			})
	}

	return nil, errors.New("no supported Gateway API TCPRoute version available")
}

type gatewayTCPRouteV1 struct {
	route v1.TCPRoute
}

func (rt *gatewayTCPRouteV1) Object() kubeObject               { return &rt.route }
func (rt *gatewayTCPRouteV1) Metadata() *metav1.ObjectMeta     { return &rt.route.ObjectMeta }
func (rt *gatewayTCPRouteV1) Hostnames() []v1.Hostname         { return nil }
func (rt *gatewayTCPRouteV1) ParentRefs() []v1.ParentReference { return rt.route.Spec.ParentRefs }
func (rt *gatewayTCPRouteV1) Protocol() v1.ProtocolType        { return v1.TCPProtocolType }
func (rt *gatewayTCPRouteV1) RouteStatus() v1.RouteStatus      { return rt.route.Status.RouteStatus }

type gatewayTCPRouteV1Informer struct {
	informers_v1.TCPRouteInformer
}

func (inf gatewayTCPRouteV1Informer) List() []gatewayRoute {
	list := extInformers.ListIndexed[*v1.TCPRoute](inf.TCPRouteInformer.Informer().GetIndexer())

	routes := make([]gatewayRoute, len(list))

	for i, rt := range list {
		// List results are supposed to be treated as read-only.
		// We make a shallow copy since we're only interested in setting the TypeMeta.
		clone := *rt
		clone.TypeMeta = metav1.TypeMeta{
			APIVersion: v1.GroupVersion.String(),
			Kind:       "TCPRoute",
		}

		routes[i] = &gatewayTCPRouteV1{clone}
	}

	return routes
}

type gatewayTCPRouteV1alpha2 struct {
	route v1alpha2.TCPRoute
}

func (rt *gatewayTCPRouteV1alpha2) Object() kubeObject               { return &rt.route }
func (rt *gatewayTCPRouteV1alpha2) Metadata() *metav1.ObjectMeta     { return &rt.route.ObjectMeta }
func (rt *gatewayTCPRouteV1alpha2) Hostnames() []v1.Hostname         { return nil }
func (rt *gatewayTCPRouteV1alpha2) ParentRefs() []v1.ParentReference { return rt.route.Spec.ParentRefs }
func (rt *gatewayTCPRouteV1alpha2) Protocol() v1.ProtocolType        { return v1.TCPProtocolType }
func (rt *gatewayTCPRouteV1alpha2) RouteStatus() v1.RouteStatus      { return rt.route.Status.RouteStatus }

type gatewayTCPRouteV1alpha2Informer struct {
	informers_v1a2.TCPRouteInformer
}

func (inf gatewayTCPRouteV1alpha2Informer) List() []gatewayRoute {
	list := extInformers.ListIndexed[*v1alpha2.TCPRoute](inf.TCPRouteInformer.Informer().GetIndexer())

	routes := make([]gatewayRoute, len(list))

	for i, rt := range list {
		// We make a shallow copy since we're only interested in setting the TypeMeta.
		clone := *rt
		clone.TypeMeta = metav1.TypeMeta{
			APIVersion: v1alpha2.GroupVersion.String(),
			Kind:       "TCPRoute",
		}

		routes[i] = &gatewayTCPRouteV1alpha2{clone}
	}
	return routes
}
