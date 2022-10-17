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

// NewGatewayTCPRouteSource creates a new Gateway TCPRoute source with the given config.
func NewGatewayTCPRouteSource(clients ClientGenerator, config *Config) (Source, error) {
	return newGatewayRouteSource(clients, config, "TCPRoute", func(factory informers.SharedInformerFactory) gatewayRouteInformer {
		return &gatewayTCPRouteInformer{factory.Gateway().V1alpha2().TCPRoutes()}
	})
}

type gatewayTCPRoute struct{ route *v1alpha2.TCPRoute }

func (rt *gatewayTCPRoute) Object() kubeObject             { return rt.route }
func (rt *gatewayTCPRoute) Metadata() *metav1.ObjectMeta   { return &rt.route.ObjectMeta }
func (rt *gatewayTCPRoute) Hostnames() []v1beta1.Hostname  { return nil }
func (rt *gatewayTCPRoute) Protocol() v1beta1.ProtocolType { return v1beta1.TCPProtocolType }
func (rt *gatewayTCPRoute) RouteStatus() v1beta1.RouteStatus {
	return v1b1RouteStatus(rt.route.Status.RouteStatus)
}

type gatewayTCPRouteInformer struct {
	informers_v1a2.TCPRouteInformer
}

func (inf gatewayTCPRouteInformer) List(namespace string, selector labels.Selector) ([]gatewayRoute, error) {
	list, err := inf.TCPRouteInformer.Lister().TCPRoutes(namespace).List(selector)
	if err != nil {
		return nil, err
	}
	routes := make([]gatewayRoute, len(list))
	for i, rt := range list {
		routes[i] = &gatewayTCPRoute{rt}
	}
	return routes, nil
}

func v1b1Hostnames(hostnames []v1alpha2.Hostname) []v1beta1.Hostname {
	if len(hostnames) == 0 {
		return nil
	}
	list := make([]v1beta1.Hostname, len(hostnames))
	for i, s := range hostnames {
		list[i] = v1beta1.Hostname(s)
	}
	return list
}

func v1b1RouteStatus(s v1alpha2.RouteStatus) v1beta1.RouteStatus {
	return v1beta1.RouteStatus{
		Parents: v1b1RouteParentStatuses(s.Parents),
	}
}

func v1b1RouteParentStatuses(statuses []v1alpha2.RouteParentStatus) []v1beta1.RouteParentStatus {
	if len(statuses) == 0 {
		return nil
	}
	list := make([]v1beta1.RouteParentStatus, len(statuses))
	for i, s := range statuses {
		list[i] = v1b1RouteParentStatus(s)
	}
	return list
}

func v1b1RouteParentStatus(s v1alpha2.RouteParentStatus) v1beta1.RouteParentStatus {
	return v1beta1.RouteParentStatus{
		ParentRef:      v1b1ParentReference(s.ParentRef),
		ControllerName: v1beta1.GatewayController(s.ControllerName),
		Conditions:     s.Conditions,
	}
}

func v1b1ParentReference(s v1alpha2.ParentReference) v1beta1.ParentReference {
	return v1beta1.ParentReference{
		Group:       (*v1beta1.Group)(s.Group),
		Kind:        (*v1beta1.Kind)(s.Kind),
		Namespace:   (*v1beta1.Namespace)(s.Namespace),
		Name:        v1beta1.ObjectName(s.Name),
		SectionName: (*v1beta1.SectionName)(s.SectionName),
		Port:        (*v1beta1.PortNumber)(s.Port),
	}
}
