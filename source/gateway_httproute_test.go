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
	"testing"

	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	kubefake "k8s.io/client-go/kubernetes/fake"
	"sigs.k8s.io/external-dns/endpoint"
	v1 "sigs.k8s.io/gateway-api/apis/v1"
	gatewayfake "sigs.k8s.io/gateway-api/pkg/client/clientset/versioned/fake"
)

func mustGetLabelSelector(s string) labels.Selector {
	v, err := getLabelSelector(s)
	if err != nil {
		panic(err)
	}
	return v
}

func gatewayStatus(ips ...string) v1.GatewayStatus {
	typ := v1.IPAddressType
	addrs := make([]v1.GatewayStatusAddress, len(ips))
	for i, ip := range ips {
		addrs[i] = v1.GatewayStatusAddress{Type: &typ, Value: ip}
	}
	return v1.GatewayStatus{Addresses: addrs}
}

func httpRouteStatus(refs ...v1.ParentReference) v1.HTTPRouteStatus {
	return v1.HTTPRouteStatus{RouteStatus: gwRouteStatus(refs...)}
}

func gwRouteStatus(refs ...v1.ParentReference) v1.RouteStatus {
	var v v1.RouteStatus
	for _, ref := range refs {
		v.Parents = append(v.Parents, v1.RouteParentStatus{
			ParentRef: ref,
			Conditions: []metav1.Condition{
				{
					Type:   string(v1.RouteConditionAccepted),
					Status: metav1.ConditionTrue,
				},
			},
		})
	}
	return v
}

func gwParentRef(namespace, name string, options ...gwParentRefOption) v1.ParentReference {
	group := v1.Group("gateway.networking.k8s.io")
	kind := v1.Kind("Gateway")
	ref := v1.ParentReference{
		Group:     &group,
		Kind:      &kind,
		Name:      v1.ObjectName(name),
		Namespace: (*v1.Namespace)(&namespace),
	}
	for _, opt := range options {
		opt(&ref)
	}
	return ref
}

type gwParentRefOption func(*v1.ParentReference)

func withSectionName(name v1.SectionName) gwParentRefOption {
	return func(ref *v1.ParentReference) { ref.SectionName = &name }
}

func withPortNumber(port v1.PortNumber) gwParentRefOption {
	return func(ref *v1.ParentReference) { ref.Port = &port }
}

func newTestEndpoint(dnsName, recordType string, targets ...string) *endpoint.Endpoint {
	return newTestEndpointWithTTL(dnsName, recordType, 0, targets...)
}

func newTestEndpointWithTTL(dnsName, recordType string, ttl int64, targets ...string) *endpoint.Endpoint {
	return &endpoint.Endpoint{
		DNSName:    dnsName,
		Targets:    append([]string(nil), targets...), // clone targets
		RecordType: recordType,
		RecordTTL:  endpoint.TTL(ttl),
	}
}

func TestGatewayHTTPRouteSourceEndpoints(t *testing.T) {
	t.Parallel()

	fromAll := v1.NamespacesFromAll
	fromSame := v1.NamespacesFromSame
	fromSelector := v1.NamespacesFromSelector
	allowAllNamespaces := &v1.AllowedRoutes{
		Namespaces: &v1.RouteNamespaces{
			From: &fromAll,
		},
	}
	objectMeta := func(namespace, name string) metav1.ObjectMeta {
		return metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		}
	}
	namespaces := func(names ...string) []*corev1.Namespace {
		v := make([]*corev1.Namespace, len(names))
		for i, name := range names {
			v[i] = &corev1.Namespace{ObjectMeta: objectMeta("", name)}
		}
		return v
	}
	hostnames := func(names ...v1.Hostname) []v1.Hostname { return names }

	tests := []struct {
		title      string
		config     Config
		namespaces []*corev1.Namespace
		gateways   []*v1.Gateway
		routes     []*v1.HTTPRoute
		endpoints  []*endpoint.Endpoint
	}{
		{
			title: "GatewayNamespace",
			config: Config{
				GatewayNamespace: "gateway-namespace",
			},
			namespaces: namespaces("gateway-namespace", "not-gateway-namespace", "route-namespace"),
			gateways: []*v1.Gateway{
				{
					ObjectMeta: objectMeta("gateway-namespace", "test"),
					Spec: v1.GatewaySpec{
						Listeners: []v1.Listener{{
							Protocol:      v1.HTTPProtocolType,
							AllowedRoutes: allowAllNamespaces,
						}},
					},
					Status: gatewayStatus("1.2.3.4"),
				},
				{
					ObjectMeta: objectMeta("not-gateway-namespace", "test"),
					Spec: v1.GatewaySpec{
						Listeners: []v1.Listener{{Protocol: v1.HTTPProtocolType}},
					},
					Status: gatewayStatus("2.3.4.5"),
				},
			},
			routes: []*v1.HTTPRoute{{
				ObjectMeta: objectMeta("route-namespace", "test"),
				Spec: v1.HTTPRouteSpec{
					Hostnames: hostnames("test.example.internal"),
				},
				Status: httpRouteStatus( // The route is attached to both gateways.
					gwParentRef("gateway-namespace", "test"),
					gwParentRef("not-gateway-namespace", "test"),
				),
			}},
			endpoints: []*endpoint.Endpoint{
				newTestEndpoint("test.example.internal", "A", "1.2.3.4"),
			},
		},
		{
			title: "RouteNamespace",
			config: Config{
				Namespace: "route-namespace",
			},
			namespaces: namespaces("gateway-namespace", "route-namespace", "not-route-namespace"),
			gateways: []*v1.Gateway{{
				ObjectMeta: objectMeta("gateway-namespace", "test"),
				Spec: v1.GatewaySpec{
					Listeners: []v1.Listener{{
						Protocol:      v1.HTTPProtocolType,
						AllowedRoutes: allowAllNamespaces,
					}},
				},
				Status: gatewayStatus("1.2.3.4"),
			}},
			routes: []*v1.HTTPRoute{
				{
					ObjectMeta: objectMeta("route-namespace", "test"),
					Spec: v1.HTTPRouteSpec{
						Hostnames: hostnames("route-namespace.example.internal"),
					},
					Status: httpRouteStatus(gwParentRef("gateway-namespace", "test")),
				},
				{
					ObjectMeta: objectMeta("not-route-namespace", "test"),
					Spec: v1.HTTPRouteSpec{
						Hostnames: hostnames("not-route-namespace.example.internal"),
					},
					Status: httpRouteStatus(gwParentRef("gateway-namespace", "test")),
				},
			},
			endpoints: []*endpoint.Endpoint{
				newTestEndpoint("route-namespace.example.internal", "A", "1.2.3.4"),
			},
		},
		{
			title: "GatewayLabelFilter",
			config: Config{
				GatewayLabelFilter: "foo=bar",
			},
			namespaces: namespaces("default"),
			gateways: []*v1.Gateway{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "labels-match",
						Namespace: "default",
						Labels:    map[string]string{"foo": "bar"},
					},
					Spec: v1.GatewaySpec{
						Listeners: []v1.Listener{{Protocol: v1.HTTPProtocolType}},
					},
					Status: gatewayStatus("1.2.3.4"),
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "labels-dont-match",
						Namespace: "default",
						Labels:    map[string]string{"foo": "qux"},
					},
					Spec: v1.GatewaySpec{
						Listeners: []v1.Listener{{Protocol: v1.HTTPProtocolType}},
					},
					Status: gatewayStatus("2.3.4.5"),
				},
			},
			routes: []*v1.HTTPRoute{{
				ObjectMeta: objectMeta("default", "test"),
				Spec: v1.HTTPRouteSpec{
					Hostnames: hostnames("test.example.internal"),
				},
				Status: httpRouteStatus( // The route is attached to both gateways.
					gwParentRef("default", "labels-match"),
					gwParentRef("default", "labels-dont-match"),
				),
			}},
			endpoints: []*endpoint.Endpoint{
				newTestEndpoint("test.example.internal", "A", "1.2.3.4"),
			},
		},
		{
			title: "RouteLabelFilter",
			config: Config{
				LabelFilter: mustGetLabelSelector("foo=bar"),
			},
			namespaces: namespaces("default"),
			gateways: []*v1.Gateway{{
				ObjectMeta: objectMeta("default", "test"),
				Spec: v1.GatewaySpec{
					Listeners: []v1.Listener{{Protocol: v1.HTTPProtocolType}},
				},
				Status: gatewayStatus("1.2.3.4"),
			}},
			routes: []*v1.HTTPRoute{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "labels-match",
						Namespace: "default",
						Labels:    map[string]string{"foo": "bar"},
					},
					Spec: v1.HTTPRouteSpec{
						Hostnames: hostnames("labels-match.example.internal"),
					},
					Status: httpRouteStatus(gwParentRef("default", "test")),
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "labels-dont-match",
						Namespace: "default",
						Labels:    map[string]string{"foo": "qux"},
					},
					Spec: v1.HTTPRouteSpec{
						Hostnames: hostnames("labels-dont-match.example.internal"),
					},
					Status: httpRouteStatus(gwParentRef("default", "test")),
				},
			},
			endpoints: []*endpoint.Endpoint{
				newTestEndpoint("labels-match.example.internal", "A", "1.2.3.4"),
			},
		},
		{
			title: "RouteAnnotationFilter",
			config: Config{
				AnnotationFilter: "foo=bar",
			},
			namespaces: namespaces("default"),
			gateways: []*v1.Gateway{{
				ObjectMeta: objectMeta("default", "test"),
				Spec: v1.GatewaySpec{
					Listeners: []v1.Listener{{Protocol: v1.HTTPProtocolType}},
				},
				Status: gatewayStatus("1.2.3.4"),
			}},
			routes: []*v1.HTTPRoute{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:        "annotations-match",
						Namespace:   "default",
						Annotations: map[string]string{"foo": "bar"},
					},
					Spec: v1.HTTPRouteSpec{
						Hostnames: hostnames("annotations-match.example.internal"),
					},
					Status: httpRouteStatus(gwParentRef("default", "test")),
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:        "annotations-dont-match",
						Namespace:   "default",
						Annotations: map[string]string{"foo": "qux"},
					},
					Spec: v1.HTTPRouteSpec{
						Hostnames: hostnames("annotations-dont-match.example.internal"),
					},
					Status: httpRouteStatus(gwParentRef("default", "test")),
				},
			},
			endpoints: []*endpoint.Endpoint{
				newTestEndpoint("annotations-match.example.internal", "A", "1.2.3.4"),
			},
		},
		{
			title:      "SkipControllerAnnotation",
			config:     Config{},
			namespaces: namespaces("default"),
			gateways: []*v1.Gateway{{
				ObjectMeta: objectMeta("default", "test"),
				Spec: v1.GatewaySpec{
					Listeners: []v1.Listener{{Protocol: v1.HTTPProtocolType}},
				},
				Status: gatewayStatus("1.2.3.4"),
			}},
			routes: []*v1.HTTPRoute{{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "api",
					Namespace: "default",
					Annotations: map[string]string{
						controllerAnnotationKey: "something-else",
					},
				},
				Spec: v1.HTTPRouteSpec{
					Hostnames: hostnames("api.example.internal"),
				},
				Status: httpRouteStatus(gwParentRef("default", "test")),
			}},
			endpoints: nil,
		},
		{
			title:      "MultipleGateways",
			config:     Config{},
			namespaces: namespaces("default"),
			gateways: []*v1.Gateway{
				{
					ObjectMeta: objectMeta("default", "one"),
					Spec: v1.GatewaySpec{
						Listeners: []v1.Listener{{Protocol: v1.HTTPProtocolType}},
					},
					Status: gatewayStatus("1.2.3.4"),
				},
				{
					ObjectMeta: objectMeta("default", "two"),
					Spec: v1.GatewaySpec{
						Listeners: []v1.Listener{{Protocol: v1.HTTPProtocolType}},
					},
					Status: gatewayStatus("2.3.4.5"),
				},
			},
			routes: []*v1.HTTPRoute{{
				ObjectMeta: objectMeta("default", "test"),
				Spec: v1.HTTPRouteSpec{
					Hostnames: hostnames("test.example.internal"),
				},
				Status: httpRouteStatus(
					gwParentRef("default", "one"),
					gwParentRef("default", "two"),
				),
			}},
			endpoints: []*endpoint.Endpoint{
				newTestEndpoint("test.example.internal", "A", "1.2.3.4", "2.3.4.5"),
			},
		},
		{
			title:      "MultipleListeners",
			config:     Config{},
			namespaces: namespaces("default"),
			gateways: []*v1.Gateway{{
				ObjectMeta: objectMeta("default", "one"),
				Spec: v1.GatewaySpec{
					Listeners: []v1.Listener{
						{
							Name:     "foo",
							Protocol: v1.HTTPProtocolType,
							Hostname: hostnamePtr("foo.example.internal"),
						},
						{
							Name:     "bar",
							Protocol: v1.HTTPProtocolType,
							Hostname: hostnamePtr("bar.example.internal"),
						},
					},
				},
				Status: gatewayStatus("1.2.3.4"),
			}},
			routes: []*v1.HTTPRoute{{
				ObjectMeta: objectMeta("default", "test"),
				Spec: v1.HTTPRouteSpec{
					Hostnames: hostnames("*.example.internal"),
				},
				Status: httpRouteStatus(
					gwParentRef("default", "one"),
				),
			}},
			endpoints: []*endpoint.Endpoint{
				newTestEndpoint("foo.example.internal", "A", "1.2.3.4"),
				newTestEndpoint("bar.example.internal", "A", "1.2.3.4"),
			},
		},
		{
			title:      "SectionNameMatch",
			config:     Config{},
			namespaces: namespaces("default"),
			gateways: []*v1.Gateway{{
				ObjectMeta: objectMeta("default", "test"),
				Spec: v1.GatewaySpec{
					Listeners: []v1.Listener{
						{
							Name:     "foo",
							Protocol: v1.HTTPProtocolType,
							Hostname: hostnamePtr("foo.example.internal"),
						},
						{
							Name:     "bar",
							Protocol: v1.HTTPProtocolType,
							Hostname: hostnamePtr("bar.example.internal"),
						},
					},
				},
				Status: gatewayStatus("1.2.3.4"),
			}},
			routes: []*v1.HTTPRoute{{
				ObjectMeta: objectMeta("default", "test"),
				Spec: v1.HTTPRouteSpec{
					Hostnames: hostnames("*.example.internal"),
				},
				Status: httpRouteStatus(
					gwParentRef("default", "test", withSectionName("foo")),
				),
			}},
			endpoints: []*endpoint.Endpoint{
				newTestEndpoint("foo.example.internal", "A", "1.2.3.4"),
			},
		},
		{
			// EXPERIMENTAL: https://gateway-api.sigs.k8s.io/geps/gep-957/
			title:      "PortNumberMatch",
			config:     Config{},
			namespaces: namespaces("default"),
			gateways: []*v1.Gateway{{
				ObjectMeta: objectMeta("default", "test"),
				Spec: v1.GatewaySpec{
					Listeners: []v1.Listener{
						{
							Name:     "foo",
							Protocol: v1.HTTPProtocolType,
							Hostname: hostnamePtr("foo.example.internal"),
							Port:     80,
						},
						{
							Name:     "bar",
							Protocol: v1.HTTPProtocolType,
							Hostname: hostnamePtr("bar.example.internal"),
							Port:     80,
						},
						{
							Name:     "qux",
							Protocol: v1.HTTPProtocolType,
							Hostname: hostnamePtr("qux.example.internal"),
							Port:     8080,
						},
					},
				},
				Status: gatewayStatus("1.2.3.4"),
			}},
			routes: []*v1.HTTPRoute{{
				ObjectMeta: objectMeta("default", "test"),
				Spec: v1.HTTPRouteSpec{
					Hostnames: hostnames("*.example.internal"),
				},
				Status: httpRouteStatus(
					gwParentRef("default", "test", withPortNumber(80)),
				),
			}},
			endpoints: []*endpoint.Endpoint{
				newTestEndpoint("foo.example.internal", "A", "1.2.3.4"),
				newTestEndpoint("bar.example.internal", "A", "1.2.3.4"),
			},
		},
		{
			title:      "WildcardInGateway",
			config:     Config{},
			namespaces: namespaces("default"),
			gateways: []*v1.Gateway{{
				ObjectMeta: objectMeta("default", "test"),
				Spec: v1.GatewaySpec{
					Listeners: []v1.Listener{{
						Protocol: v1.HTTPProtocolType,
						Hostname: hostnamePtr("*.example.internal"),
					}},
				},
				Status: gatewayStatus("1.2.3.4"),
			}},
			routes: []*v1.HTTPRoute{{
				ObjectMeta: objectMeta("default", "no-hostname"),
				Spec: v1.HTTPRouteSpec{
					Hostnames: []v1.Hostname{
						"foo.example.internal",
					},
				},
				Status: httpRouteStatus(gwParentRef("default", "test")),
			}},
			endpoints: []*endpoint.Endpoint{
				newTestEndpoint("foo.example.internal", "A", "1.2.3.4"),
			},
		},
		{
			title:      "WildcardInRoute",
			config:     Config{},
			namespaces: namespaces("default"),
			gateways: []*v1.Gateway{{
				ObjectMeta: objectMeta("default", "test"),
				Spec: v1.GatewaySpec{
					Listeners: []v1.Listener{{
						Protocol: v1.HTTPProtocolType,
						Hostname: hostnamePtr("foo.example.internal"),
					}},
				},
				Status: gatewayStatus("1.2.3.4"),
			}},
			routes: []*v1.HTTPRoute{{
				ObjectMeta: objectMeta("default", "no-hostname"),
				Spec: v1.HTTPRouteSpec{
					Hostnames: []v1.Hostname{
						"*.example.internal",
					},
				},
				Status: httpRouteStatus(gwParentRef("default", "test")),
			}},
			endpoints: []*endpoint.Endpoint{
				newTestEndpoint("foo.example.internal", "A", "1.2.3.4"),
			},
		},
		{
			title:      "WildcardInRouteAndGateway",
			config:     Config{},
			namespaces: namespaces("default"),
			gateways: []*v1.Gateway{{
				ObjectMeta: objectMeta("default", "test"),
				Spec: v1.GatewaySpec{
					Listeners: []v1.Listener{{
						Protocol: v1.HTTPProtocolType,
						Hostname: hostnamePtr("*.example.internal"),
					}},
				},
				Status: gatewayStatus("1.2.3.4"),
			}},
			routes: []*v1.HTTPRoute{{
				ObjectMeta: objectMeta("default", "no-hostname"),
				Spec: v1.HTTPRouteSpec{
					Hostnames: []v1.Hostname{
						"*.example.internal",
					},
				},
				Status: httpRouteStatus(gwParentRef("default", "test")),
			}},
			endpoints: []*endpoint.Endpoint{
				newTestEndpoint("*.example.internal", "A", "1.2.3.4"),
			},
		},
		{
			title:      "NoRouteHostname",
			config:     Config{},
			namespaces: namespaces("default"),
			gateways: []*v1.Gateway{{
				ObjectMeta: objectMeta("default", "test"),
				Spec: v1.GatewaySpec{
					Listeners: []v1.Listener{{
						Protocol: v1.HTTPProtocolType,
						Hostname: hostnamePtr("foo.example.internal"),
					}},
				},
				Status: gatewayStatus("1.2.3.4"),
			}},
			routes: []*v1.HTTPRoute{{
				ObjectMeta: objectMeta("default", "no-hostname"),
				Spec: v1.HTTPRouteSpec{
					Hostnames: nil,
				},
				Status: httpRouteStatus(gwParentRef("default", "test")),
			}},
			endpoints: []*endpoint.Endpoint{
				newTestEndpoint("foo.example.internal", "A", "1.2.3.4"),
			},
		},
		{
			title:      "NoGateways",
			config:     Config{},
			namespaces: namespaces("default"),
			gateways:   nil,
			routes: []*v1.HTTPRoute{{
				ObjectMeta: objectMeta("default", "test"),
				Spec: v1.HTTPRouteSpec{
					Hostnames: hostnames("example.internal"),
				},
				Status: httpRouteStatus(),
			}},
			endpoints: nil,
		},
		{
			title:      "NoHostnames",
			config:     Config{},
			namespaces: namespaces("default"),
			gateways: []*v1.Gateway{{
				ObjectMeta: objectMeta("default", "test"),
				Spec: v1.GatewaySpec{
					Listeners: []v1.Listener{{Protocol: v1.HTTPProtocolType}},
				},
				Status: gatewayStatus("1.2.3.4"),
			}},
			routes: []*v1.HTTPRoute{{
				ObjectMeta: objectMeta("default", "no-hostname"),
				Spec: v1.HTTPRouteSpec{
					Hostnames: nil,
				},
				Status: httpRouteStatus(gwParentRef("default", "test")),
			}},
			endpoints: nil,
		},
		{
			title:      "HostnameAnnotation",
			config:     Config{},
			namespaces: namespaces("default"),
			gateways: []*v1.Gateway{{
				ObjectMeta: objectMeta("default", "test"),
				Spec: v1.GatewaySpec{
					Listeners: []v1.Listener{{Protocol: v1.HTTPProtocolType}},
				},
				Status: gatewayStatus("1.2.3.4"),
			}},
			routes: []*v1.HTTPRoute{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "without-hostame",
						Namespace: "default",
						Annotations: map[string]string{
							hostnameAnnotationKey: "annotation.without-hostname.internal",
						},
					},
					Spec: v1.HTTPRouteSpec{
						Hostnames: nil,
					},
					Status: httpRouteStatus(gwParentRef("default", "test")),
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "with-hostame",
						Namespace: "default",
						Annotations: map[string]string{
							hostnameAnnotationKey: "annotation.with-hostname.internal",
						},
					},
					Spec: v1.HTTPRouteSpec{
						Hostnames: hostnames("with-hostname.internal"),
					},
					Status: httpRouteStatus(gwParentRef("default", "test")),
				},
			},
			endpoints: []*endpoint.Endpoint{
				newTestEndpoint("annotation.without-hostname.internal", "A", "1.2.3.4"),
				newTestEndpoint("annotation.with-hostname.internal", "A", "1.2.3.4"),
				newTestEndpoint("with-hostname.internal", "A", "1.2.3.4"),
			},
		},
		{
			title: "IgnoreHostnameAnnotation",
			config: Config{
				IgnoreHostnameAnnotation: true,
			},
			namespaces: namespaces("default"),
			gateways: []*v1.Gateway{{
				ObjectMeta: objectMeta("default", "test"),
				Spec: v1.GatewaySpec{
					Listeners: []v1.Listener{{Protocol: v1.HTTPProtocolType}},
				},
				Status: gatewayStatus("1.2.3.4"),
			}},
			routes: []*v1.HTTPRoute{{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "with-hostame",
					Namespace: "default",
					Annotations: map[string]string{
						hostnameAnnotationKey: "annotation.with-hostname.internal",
					},
				},
				Spec: v1.HTTPRouteSpec{
					Hostnames: hostnames("with-hostname.internal"),
				},
				Status: httpRouteStatus(gwParentRef("default", "test")),
			}},
			endpoints: []*endpoint.Endpoint{
				newTestEndpoint("with-hostname.internal", "A", "1.2.3.4"),
			},
		},
		{
			title: "FQDNTemplate",
			config: Config{
				FQDNTemplate: "{{.Name}}.zero.internal, {{.Name}}.one.internal. ,  {{.Name}}.two.internal  ",
			},
			namespaces: namespaces("default"),
			gateways: []*v1.Gateway{{
				ObjectMeta: objectMeta("default", "test"),
				Spec: v1.GatewaySpec{
					Listeners: []v1.Listener{{Protocol: v1.HTTPProtocolType}},
				},
				Status: gatewayStatus("1.2.3.4"),
			}},
			routes: []*v1.HTTPRoute{
				{
					ObjectMeta: objectMeta("default", "fqdn-with-hostnames"),
					Spec: v1.HTTPRouteSpec{
						Hostnames: hostnames("fqdn-with-hostnames.internal"),
					},
					Status: httpRouteStatus(gwParentRef("default", "test")),
				},
				{
					ObjectMeta: objectMeta("default", "fqdn-without-hostnames"),
					Spec: v1.HTTPRouteSpec{
						Hostnames: nil,
					},
					Status: httpRouteStatus(gwParentRef("default", "test")),
				},
			},
			endpoints: []*endpoint.Endpoint{
				newTestEndpoint("fqdn-without-hostnames.zero.internal", "A", "1.2.3.4"),
				newTestEndpoint("fqdn-without-hostnames.one.internal", "A", "1.2.3.4"),
				newTestEndpoint("fqdn-without-hostnames.two.internal", "A", "1.2.3.4"),
				newTestEndpoint("fqdn-with-hostnames.internal", "A", "1.2.3.4"),
			},
		},
		{
			title: "CombineFQDN",
			config: Config{
				FQDNTemplate:             "combine-{{.Name}}.internal",
				CombineFQDNAndAnnotation: true,
			},
			namespaces: namespaces("default"),
			gateways: []*v1.Gateway{{
				ObjectMeta: objectMeta("default", "test"),
				Spec: v1.GatewaySpec{
					Listeners: []v1.Listener{{Protocol: v1.HTTPProtocolType}},
				},
				Status: gatewayStatus("1.2.3.4"),
			}},
			routes: []*v1.HTTPRoute{{
				ObjectMeta: objectMeta("default", "fqdn-with-hostnames"),
				Spec: v1.HTTPRouteSpec{
					Hostnames: hostnames("fqdn-with-hostnames.internal"),
				},
				Status: httpRouteStatus(gwParentRef("default", "test")),
			}},
			endpoints: []*endpoint.Endpoint{
				newTestEndpoint("fqdn-with-hostnames.internal", "A", "1.2.3.4"),
				newTestEndpoint("combine-fqdn-with-hostnames.internal", "A", "1.2.3.4"),
			},
		},
		{
			title:      "TTL",
			config:     Config{},
			namespaces: namespaces("default"),
			gateways: []*v1.Gateway{{
				ObjectMeta: objectMeta("default", "test"),
				Spec: v1.GatewaySpec{
					Listeners: []v1.Listener{{Protocol: v1.HTTPProtocolType}},
				},
				Status: gatewayStatus("1.2.3.4"),
			}},
			routes: []*v1.HTTPRoute{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:        "valid-ttl",
						Namespace:   "default",
						Annotations: map[string]string{ttlAnnotationKey: "15s"},
					},
					Spec: v1.HTTPRouteSpec{
						Hostnames: hostnames("valid-ttl.internal"),
					},
					Status: httpRouteStatus(gwParentRef("default", "test")),
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:        "invalid-ttl",
						Namespace:   "default",
						Annotations: map[string]string{ttlAnnotationKey: "abc"},
					},
					Spec: v1.HTTPRouteSpec{
						Hostnames: hostnames("invalid-ttl.internal"),
					},
					Status: httpRouteStatus(gwParentRef("default", "test")),
				},
			},
			endpoints: []*endpoint.Endpoint{
				newTestEndpoint("invalid-ttl.internal", "A", "1.2.3.4"),
				newTestEndpointWithTTL("valid-ttl.internal", "A", 15, "1.2.3.4"),
			},
		},
		{
			title:      "ProviderAnnotations",
			config:     Config{},
			namespaces: namespaces("default"),
			gateways: []*v1.Gateway{{
				ObjectMeta: objectMeta("default", "test"),
				Spec: v1.GatewaySpec{
					Listeners: []v1.Listener{{Protocol: v1.HTTPProtocolType}},
				},
				Status: gatewayStatus("1.2.3.4"),
			}},
			routes: []*v1.HTTPRoute{{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "provider-annotations",
					Namespace: "default",
					Annotations: map[string]string{
						SetIdentifierKey:   "test-set-identifier",
						aliasAnnotationKey: "true",
					},
				},
				Spec: v1.HTTPRouteSpec{
					Hostnames: hostnames("provider-annotations.com"),
				},
				Status: httpRouteStatus(gwParentRef("default", "test")),
			}},
			endpoints: []*endpoint.Endpoint{
				newTestEndpoint("provider-annotations.com", "A", "1.2.3.4").
					WithProviderSpecific("alias", "true").
					WithSetIdentifier("test-set-identifier"),
			},
		},
		{
			title:      "DifferentHostnameDifferentGateway",
			config:     Config{},
			namespaces: namespaces("default"),
			gateways: []*v1.Gateway{
				{
					ObjectMeta: objectMeta("default", "one"),
					Spec: v1.GatewaySpec{
						Listeners: []v1.Listener{{
							Hostname: hostnamePtr("*.one.internal"),
							Protocol: v1.HTTPProtocolType,
						}},
					},
					Status: gatewayStatus("1.2.3.4"),
				},
				{
					ObjectMeta: objectMeta("default", "two"),
					Spec: v1.GatewaySpec{
						Listeners: []v1.Listener{{
							Hostname: hostnamePtr("*.two.internal"),
							Protocol: v1.HTTPProtocolType,
						}},
					},
					Status: gatewayStatus("2.3.4.5"),
				},
			},
			routes: []*v1.HTTPRoute{{
				ObjectMeta: objectMeta("default", "test"),
				Spec: v1.HTTPRouteSpec{
					Hostnames: hostnames("test.one.internal", "test.two.internal"),
				},
				Status: httpRouteStatus(
					gwParentRef("default", "one"),
					gwParentRef("default", "two"),
				),
			}},
			endpoints: []*endpoint.Endpoint{
				newTestEndpoint("test.one.internal", "A", "1.2.3.4"),
				newTestEndpoint("test.two.internal", "A", "2.3.4.5"),
			},
		},
		{
			title:      "AllowedRoutesSameNamespace",
			config:     Config{},
			namespaces: namespaces("same-namespace", "other-namespace"),
			gateways: []*v1.Gateway{{
				ObjectMeta: objectMeta("same-namespace", "test"),
				Spec: v1.GatewaySpec{
					Listeners: []v1.Listener{{
						Protocol: v1.HTTPProtocolType,
						AllowedRoutes: &v1.AllowedRoutes{
							Namespaces: &v1.RouteNamespaces{
								From: &fromSame,
							},
						},
					}},
				},
				Status: gatewayStatus("1.2.3.4"),
			}},
			routes: []*v1.HTTPRoute{
				{
					ObjectMeta: objectMeta("same-namespace", "test"),
					Spec: v1.HTTPRouteSpec{
						Hostnames: hostnames("same-namespace.example.internal"),
					},
					Status: httpRouteStatus(gwParentRef("same-namespace", "test")),
				},
				{
					ObjectMeta: objectMeta("other-namespace", "test"),
					Spec: v1.HTTPRouteSpec{
						Hostnames: hostnames("other-namespace.example.internal"),
					},
					Status: httpRouteStatus(gwParentRef("same-namespace", "test")),
				},
			},
			endpoints: []*endpoint.Endpoint{
				newTestEndpoint("same-namespace.example.internal", "A", "1.2.3.4"),
			},
		},
		{
			title:  "AllowedRoutesNamespaceSelector",
			config: Config{},
			namespaces: []*corev1.Namespace{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "default",
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:   "foo",
						Labels: map[string]string{"team": "foo"},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:   "bar",
						Labels: map[string]string{"team": "bar"},
					},
				},
			},
			gateways: []*v1.Gateway{{
				ObjectMeta: objectMeta("default", "test"),
				Spec: v1.GatewaySpec{
					Listeners: []v1.Listener{{
						Protocol: v1.HTTPProtocolType,
						AllowedRoutes: &v1.AllowedRoutes{
							Namespaces: &v1.RouteNamespaces{
								From: &fromSelector,
								Selector: &metav1.LabelSelector{
									MatchLabels: map[string]string{"team": "foo"},
								},
							},
						},
					}},
				},
				Status: gatewayStatus("1.2.3.4"),
			}},
			routes: []*v1.HTTPRoute{
				{
					ObjectMeta: objectMeta("foo", "test"),
					Spec: v1.HTTPRouteSpec{
						Hostnames: hostnames("foo.example.internal"),
					},
					Status: httpRouteStatus(gwParentRef("default", "test")),
				},
				{
					ObjectMeta: objectMeta("bar", "test"),
					Spec: v1.HTTPRouteSpec{
						Hostnames: hostnames("bar.example.internal"),
					},
					Status: httpRouteStatus(gwParentRef("default", "test")),
				},
			},
			endpoints: []*endpoint.Endpoint{
				newTestEndpoint("foo.example.internal", "A", "1.2.3.4"),
			},
		},
		{
			title:      "MissingNamespace",
			config:     Config{},
			namespaces: nil,
			gateways: []*v1.Gateway{{
				ObjectMeta: objectMeta("default", "test"),
				Spec: v1.GatewaySpec{
					Listeners: []v1.Listener{{
						Protocol: v1.HTTPProtocolType,
						AllowedRoutes: &v1.AllowedRoutes{
							Namespaces: &v1.RouteNamespaces{
								// Namespace selector triggers namespace lookup.
								From: &fromSelector,
								Selector: &metav1.LabelSelector{
									MatchLabels: map[string]string{"foo": "bar"},
								},
							},
						},
					}},
				},
				Status: gatewayStatus("1.2.3.4"),
			}},
			routes: []*v1.HTTPRoute{{
				ObjectMeta: objectMeta("default", "test"),
				Spec: v1.HTTPRouteSpec{
					Hostnames: hostnames("example.internal"),
				},
				Status: httpRouteStatus(gwParentRef("default", "test")),
			}},
			endpoints: nil,
		},
		{
			title: "AnnotationOverride",
			config: Config{
				GatewayNamespace: "gateway-namespace",
			},
			namespaces: namespaces("gateway-namespace", "route-namespace"),
			gateways: []*v1.Gateway{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "overriden-gateway",
						Namespace: "gateway-namespace",
						Annotations: map[string]string{
							targetAnnotationKey: "4.3.2.1",
						},
					},
					Spec: v1.GatewaySpec{
						Listeners: []v1.Listener{{
							Protocol:      v1.HTTPProtocolType,
							AllowedRoutes: allowAllNamespaces,
						}},
					},
					Status: gatewayStatus("1.2.3.4"),
				},
			},
			routes: []*v1.HTTPRoute{{
				ObjectMeta: objectMeta("route-namespace", "test"),
				Spec: v1.HTTPRouteSpec{
					Hostnames: hostnames("test.example.internal"),
				},
				Status: httpRouteStatus( // The route is attached to both gateways.
					gwParentRef("gateway-namespace", "overriden-gateway"),
				),
			}},
			endpoints: []*endpoint.Endpoint{
				newTestEndpoint("test.example.internal", "A", "4.3.2.1"),
			},
		},
		{
			title: "MutlipleGatewaysOneAnnotationOverride",
			config: Config{
				GatewayNamespace: "gateway-namespace",
			},
			namespaces: namespaces("gateway-namespace", "route-namespace"),
			gateways: []*v1.Gateway{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "overriden-gateway",
						Namespace: "gateway-namespace",
						Annotations: map[string]string{
							targetAnnotationKey: "4.3.2.1",
						},
					},
					Spec: v1.GatewaySpec{
						Listeners: []v1.Listener{{
							Protocol:      v1.HTTPProtocolType,
							AllowedRoutes: allowAllNamespaces,
						}},
					},
					Status: gatewayStatus("1.2.3.4"),
				},
				{
					ObjectMeta: objectMeta("gateway-namespace", "test"),
					Spec: v1.GatewaySpec{
						Listeners: []v1.Listener{{
							Protocol:      v1.HTTPProtocolType,
							AllowedRoutes: allowAllNamespaces,
						}},
					},
					Status: gatewayStatus("2.3.4.5"),
				},
			},
			routes: []*v1.HTTPRoute{{
				ObjectMeta: objectMeta("route-namespace", "test"),
				Spec: v1.HTTPRouteSpec{
					Hostnames: hostnames("test.example.internal"),
				},
				Status: httpRouteStatus( // The route is attached to both gateways.
					gwParentRef("gateway-namespace", "overriden-gateway"),
					gwParentRef("gateway-namespace", "test"),
				),
			}},
			endpoints: []*endpoint.Endpoint{
				newTestEndpoint("test.example.internal", "A", "4.3.2.1", "2.3.4.5"),
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.title, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			gwClient := gatewayfake.NewSimpleClientset()
			for _, gw := range tt.gateways {
				_, err := gwClient.GatewayV1().Gateways(gw.Namespace).Create(ctx, gw, metav1.CreateOptions{})
				require.NoError(t, err, "failed to create Gateway")

			}
			for _, rt := range tt.routes {
				_, err := gwClient.GatewayV1().HTTPRoutes(rt.Namespace).Create(ctx, rt, metav1.CreateOptions{})
				require.NoError(t, err, "failed to create HTTPRoute")
			}
			kubeClient := kubefake.NewSimpleClientset()
			for _, ns := range tt.namespaces {
				_, err := kubeClient.CoreV1().Namespaces().Create(ctx, ns, metav1.CreateOptions{})
				require.NoError(t, err, "failed to create Namespace")
			}

			clients := new(MockClientGenerator)
			clients.On("GatewayClient").Return(gwClient, nil)
			clients.On("KubeClient").Return(kubeClient, nil)

			src, err := NewGatewayHTTPRouteSource(clients, &tt.config)
			require.NoError(t, err, "failed to create Gateway HTTPRoute Source")

			endpoints, err := src.Endpoints(ctx)
			require.NoError(t, err, "failed to get Endpoints")
			validateEndpoints(t, endpoints, tt.endpoints)
		})
	}
}

func hostnamePtr(val v1.Hostname) *v1.Hostname { return &val }
