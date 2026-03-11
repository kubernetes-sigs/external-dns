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
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubefake "k8s.io/client-go/kubernetes/fake"
	"sigs.k8s.io/external-dns/endpoint"
	v1 "sigs.k8s.io/gateway-api/apis/v1"
	v1beta1 "sigs.k8s.io/gateway-api/apis/v1beta1"
	gatewayfake "sigs.k8s.io/gateway-api/pkg/client/clientset/versioned/fake"
)

func lsParentRef(namespace, name string, options ...gwParentRefOption) v1.ParentReference {
	group := v1.Group(gatewayGroup)
	kind := v1.Kind("ListenerSet")
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

func TestGatewayHTTPRouteWithListenerSetParentRef(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(t.Context(), 10*time.Second)
	defer cancel()

	gwClient := gatewayfake.NewSimpleClientset()
	kubeClient := kubefake.NewClientset()
	clients := new(MockClientGenerator)
	clients.On("GatewayClient").Return(gwClient, nil)
	clients.On("KubeClient").Return(kubeClient, nil)

	ns := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "default"}}
	_, err := kubeClient.CoreV1().Namespaces().Create(ctx, ns, metav1.CreateOptions{})
	require.NoError(t, err)

	ips := []string{"10.64.0.1", "10.64.0.2"}
	gw := &v1beta1.Gateway{
		ObjectMeta: metav1.ObjectMeta{Name: "my-gateway", Namespace: "default"},
		Spec: v1.GatewaySpec{
			Listeners: []v1.Listener{{
				Name:     "base",
				Protocol: v1.HTTPProtocolType,
				Port:     80,
			}},
		},
		Status: gatewayStatus(ips...),
	}
	_, err = gwClient.GatewayV1beta1().Gateways(gw.Namespace).Create(ctx, gw, metav1.CreateOptions{})
	require.NoError(t, err)

	hostname := v1.Hostname("app.example.com")
	fromAll := v1.NamespacesFromAll
	ls := &v1.ListenerSet{
		ObjectMeta: metav1.ObjectMeta{Name: "my-listenerset", Namespace: "default"},
		Spec: v1.ListenerSetSpec{
			ParentRef: v1.ParentGatewayReference{Name: "my-gateway"},
			Listeners: []v1.ListenerEntry{{
				Name:     "app",
				Hostname: &hostname,
				Port:     8080,
				Protocol: v1.HTTPProtocolType,
				AllowedRoutes: &v1.AllowedRoutes{
					Namespaces: &v1.RouteNamespaces{From: &fromAll},
				},
			}},
		},
	}
	_, err = gwClient.GatewayV1().ListenerSets(ls.Namespace).Create(ctx, ls, metav1.CreateOptions{})
	require.NoError(t, err)

	rt := &v1beta1.HTTPRoute{
		ObjectMeta: metav1.ObjectMeta{Name: "my-route", Namespace: "default"},
		Spec: v1.HTTPRouteSpec{
			Hostnames: []v1.Hostname{"app.example.com"},
			CommonRouteSpec: v1.CommonRouteSpec{
				ParentRefs: []v1.ParentReference{
					lsParentRef("default", "my-listenerset"),
				},
			},
		},
		Status: v1.HTTPRouteStatus{
			RouteStatus: gwRouteStatus(lsParentRef("default", "my-listenerset")),
		},
	}
	_, err = gwClient.GatewayV1beta1().HTTPRoutes(rt.Namespace).Create(ctx, rt, metav1.CreateOptions{})
	require.NoError(t, err)

	src, err := NewGatewayHTTPRouteSource(ctx, clients, &Config{})
	require.NoError(t, err)

	endpoints, err := src.Endpoints(ctx)
	require.NoError(t, err)
	validateEndpoints(t, endpoints, []*endpoint.Endpoint{
		newTestEndpoint("app.example.com", ips...),
	})
}

func TestGatewayHTTPRouteWithListenerSetNotAccepted(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(t.Context(), 10*time.Second)
	defer cancel()

	gwClient := gatewayfake.NewSimpleClientset()
	kubeClient := kubefake.NewClientset()
	clients := new(MockClientGenerator)
	clients.On("GatewayClient").Return(gwClient, nil)
	clients.On("KubeClient").Return(kubeClient, nil)

	ns := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "default"}}
	_, err := kubeClient.CoreV1().Namespaces().Create(ctx, ns, metav1.CreateOptions{})
	require.NoError(t, err)

	gw := &v1beta1.Gateway{
		ObjectMeta: metav1.ObjectMeta{Name: "gw", Namespace: "default"},
		Spec:       v1.GatewaySpec{Listeners: []v1.Listener{{Protocol: v1.HTTPProtocolType, Port: 80}}},
		Status:     gatewayStatus("10.0.0.1"),
	}
	_, err = gwClient.GatewayV1beta1().Gateways(gw.Namespace).Create(ctx, gw, metav1.CreateOptions{})
	require.NoError(t, err)

	hostname := v1.Hostname("app.example.com")
	fromAll := v1.NamespacesFromAll
	ls := &v1.ListenerSet{
		ObjectMeta: metav1.ObjectMeta{Name: "ls", Namespace: "default"},
		Spec: v1.ListenerSetSpec{
			ParentRef: v1.ParentGatewayReference{Name: "gw"},
			Listeners: []v1.ListenerEntry{{
				Name: "app", Hostname: &hostname, Port: 8080, Protocol: v1.HTTPProtocolType,
				AllowedRoutes: &v1.AllowedRoutes{Namespaces: &v1.RouteNamespaces{From: &fromAll}},
			}},
		},
	}
	_, err = gwClient.GatewayV1().ListenerSets(ls.Namespace).Create(ctx, ls, metav1.CreateOptions{})
	require.NoError(t, err)

	// Route with ListenerSet parentRef, but NOT accepted (no Accepted condition).
	rt := &v1beta1.HTTPRoute{
		ObjectMeta: metav1.ObjectMeta{Name: "rt", Namespace: "default"},
		Spec: v1.HTTPRouteSpec{
			Hostnames:       []v1.Hostname{"app.example.com"},
			CommonRouteSpec: v1.CommonRouteSpec{ParentRefs: []v1.ParentReference{lsParentRef("default", "ls")}},
		},
		Status: v1.HTTPRouteStatus{
			RouteStatus: v1.RouteStatus{
				Parents: []v1.RouteParentStatus{{
					ParentRef: lsParentRef("default", "ls"),
					Conditions: []metav1.Condition{{
						Type:   string(v1.RouteConditionAccepted),
						Status: metav1.ConditionFalse,
					}},
				}},
			},
		},
	}
	_, err = gwClient.GatewayV1beta1().HTTPRoutes(rt.Namespace).Create(ctx, rt, metav1.CreateOptions{})
	require.NoError(t, err)

	src, err := NewGatewayHTTPRouteSource(ctx, clients, &Config{})
	require.NoError(t, err)

	endpoints, err := src.Endpoints(ctx)
	require.NoError(t, err)
	require.Empty(t, endpoints, "expected no endpoints when route is not accepted by ListenerSet")
}

func TestGatewayHTTPRouteWithListenerSetTargetAnnotation(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(t.Context(), 10*time.Second)
	defer cancel()

	gwClient := gatewayfake.NewSimpleClientset()
	kubeClient := kubefake.NewClientset()
	clients := new(MockClientGenerator)
	clients.On("GatewayClient").Return(gwClient, nil)
	clients.On("KubeClient").Return(kubeClient, nil)

	ns := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "default"}}
	_, err := kubeClient.CoreV1().Namespaces().Create(ctx, ns, metav1.CreateOptions{})
	require.NoError(t, err)

	// Gateway with target annotation override.
	gw := &v1beta1.Gateway{
		ObjectMeta: metav1.ObjectMeta{
			Name: "gw", Namespace: "default",
			Annotations: map[string]string{"external-dns.alpha.kubernetes.io/target": "override.example.com"},
		},
		Spec:   v1.GatewaySpec{Listeners: []v1.Listener{{Protocol: v1.HTTPProtocolType, Port: 80}}},
		Status: gatewayStatus("10.0.0.1"),
	}
	_, err = gwClient.GatewayV1beta1().Gateways(gw.Namespace).Create(ctx, gw, metav1.CreateOptions{})
	require.NoError(t, err)

	hostname := v1.Hostname("app.example.com")
	fromAll := v1.NamespacesFromAll
	ls := &v1.ListenerSet{
		ObjectMeta: metav1.ObjectMeta{Name: "ls", Namespace: "default"},
		Spec: v1.ListenerSetSpec{
			ParentRef: v1.ParentGatewayReference{Name: "gw"},
			Listeners: []v1.ListenerEntry{{
				Name: "app", Hostname: &hostname, Port: 8080, Protocol: v1.HTTPProtocolType,
				AllowedRoutes: &v1.AllowedRoutes{Namespaces: &v1.RouteNamespaces{From: &fromAll}},
			}},
		},
	}
	_, err = gwClient.GatewayV1().ListenerSets(ls.Namespace).Create(ctx, ls, metav1.CreateOptions{})
	require.NoError(t, err)

	rt := &v1beta1.HTTPRoute{
		ObjectMeta: metav1.ObjectMeta{Name: "rt", Namespace: "default"},
		Spec: v1.HTTPRouteSpec{
			Hostnames:       []v1.Hostname{"app.example.com"},
			CommonRouteSpec: v1.CommonRouteSpec{ParentRefs: []v1.ParentReference{lsParentRef("default", "ls")}},
		},
		Status: v1.HTTPRouteStatus{RouteStatus: gwRouteStatus(lsParentRef("default", "ls"))},
	}
	_, err = gwClient.GatewayV1beta1().HTTPRoutes(rt.Namespace).Create(ctx, rt, metav1.CreateOptions{})
	require.NoError(t, err)

	src, err := NewGatewayHTTPRouteSource(ctx, clients, &Config{})
	require.NoError(t, err)

	endpoints, err := src.Endpoints(ctx)
	require.NoError(t, err)
	validateEndpoints(t, endpoints, []*endpoint.Endpoint{
		newTestEndpointWithTTL("app.example.com", endpoint.RecordTypeCNAME, 0, "override.example.com"),
	})
}

func TestGatewayHTTPRouteWithListenerSetAllowedRoutesSame(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(t.Context(), 10*time.Second)
	defer cancel()

	gwClient := gatewayfake.NewSimpleClientset()
	kubeClient := kubefake.NewClientset()
	clients := new(MockClientGenerator)
	clients.On("GatewayClient").Return(gwClient, nil)
	clients.On("KubeClient").Return(kubeClient, nil)

	for _, name := range []string{"default", "other"} {
		ns := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: name}}
		_, err := kubeClient.CoreV1().Namespaces().Create(ctx, ns, metav1.CreateOptions{})
		require.NoError(t, err)
	}

	gw := &v1beta1.Gateway{
		ObjectMeta: metav1.ObjectMeta{Name: "gw", Namespace: "default"},
		Spec:       v1.GatewaySpec{Listeners: []v1.Listener{{Protocol: v1.HTTPProtocolType, Port: 80}}},
		Status:     gatewayStatus("10.0.0.1"),
	}
	_, err := gwClient.GatewayV1beta1().Gateways(gw.Namespace).Create(ctx, gw, metav1.CreateOptions{})
	require.NoError(t, err)

	// ListenerSet in "default" namespace with AllowedRoutes: Same (default).
	hostname := v1.Hostname("app.example.com")
	ls := &v1.ListenerSet{
		ObjectMeta: metav1.ObjectMeta{Name: "ls", Namespace: "default"},
		Spec: v1.ListenerSetSpec{
			ParentRef: v1.ParentGatewayReference{Name: "gw"},
			Listeners: []v1.ListenerEntry{{
				Name: "app", Hostname: &hostname, Port: 8080, Protocol: v1.HTTPProtocolType,
				// No AllowedRoutes → defaults to Same namespace
			}},
		},
	}
	_, err = gwClient.GatewayV1().ListenerSets(ls.Namespace).Create(ctx, ls, metav1.CreateOptions{})
	require.NoError(t, err)

	// Route in "other" namespace → should NOT match (Same means ListenerSet namespace).
	rt := &v1beta1.HTTPRoute{
		ObjectMeta: metav1.ObjectMeta{Name: "rt", Namespace: "other"},
		Spec: v1.HTTPRouteSpec{
			Hostnames:       []v1.Hostname{"app.example.com"},
			CommonRouteSpec: v1.CommonRouteSpec{ParentRefs: []v1.ParentReference{lsParentRef("default", "ls")}},
		},
		Status: v1.HTTPRouteStatus{RouteStatus: gwRouteStatus(lsParentRef("default", "ls"))},
	}
	_, err = gwClient.GatewayV1beta1().HTTPRoutes(rt.Namespace).Create(ctx, rt, metav1.CreateOptions{})
	require.NoError(t, err)

	src, err := NewGatewayHTTPRouteSource(ctx, clients, &Config{})
	require.NoError(t, err)

	endpoints, err := src.Endpoints(ctx)
	require.NoError(t, err)
	require.Empty(t, endpoints, "expected no endpoints when route is in different namespace than ListenerSet with Same allowed")
}

func TestGatewayHTTPRouteWithListenerSetSectionName(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(t.Context(), 10*time.Second)
	defer cancel()

	gwClient := gatewayfake.NewSimpleClientset()
	kubeClient := kubefake.NewClientset()
	clients := new(MockClientGenerator)
	clients.On("GatewayClient").Return(gwClient, nil)
	clients.On("KubeClient").Return(kubeClient, nil)

	ns := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "default"}}
	_, err := kubeClient.CoreV1().Namespaces().Create(ctx, ns, metav1.CreateOptions{})
	require.NoError(t, err)

	gw := &v1beta1.Gateway{
		ObjectMeta: metav1.ObjectMeta{Name: "gw", Namespace: "default"},
		Spec:       v1.GatewaySpec{Listeners: []v1.Listener{{Protocol: v1.HTTPProtocolType, Port: 80}}},
		Status:     gatewayStatus("10.0.0.1"),
	}
	_, err = gwClient.GatewayV1beta1().Gateways(gw.Namespace).Create(ctx, gw, metav1.CreateOptions{})
	require.NoError(t, err)

	hostname1 := v1.Hostname("app.example.com")
	hostname2 := v1.Hostname("api.example.com")
	fromAll := v1.NamespacesFromAll
	ls := &v1.ListenerSet{
		ObjectMeta: metav1.ObjectMeta{Name: "ls", Namespace: "default"},
		Spec: v1.ListenerSetSpec{
			ParentRef: v1.ParentGatewayReference{Name: "gw"},
			Listeners: []v1.ListenerEntry{
				{Name: "app", Hostname: &hostname1, Port: 8080, Protocol: v1.HTTPProtocolType,
					AllowedRoutes: &v1.AllowedRoutes{Namespaces: &v1.RouteNamespaces{From: &fromAll}}},
				{Name: "api", Hostname: &hostname2, Port: 8081, Protocol: v1.HTTPProtocolType,
					AllowedRoutes: &v1.AllowedRoutes{Namespaces: &v1.RouteNamespaces{From: &fromAll}}},
			},
		},
	}
	_, err = gwClient.GatewayV1().ListenerSets(ls.Namespace).Create(ctx, ls, metav1.CreateOptions{})
	require.NoError(t, err)

	// Route targeting only the "api" section of the ListenerSet.
	rt := &v1beta1.HTTPRoute{
		ObjectMeta: metav1.ObjectMeta{Name: "rt", Namespace: "default"},
		Spec: v1.HTTPRouteSpec{
			Hostnames: []v1.Hostname{"api.example.com"},
			CommonRouteSpec: v1.CommonRouteSpec{
				ParentRefs: []v1.ParentReference{
					lsParentRef("default", "ls", withSectionName("api")),
				},
			},
		},
		Status: v1.HTTPRouteStatus{
			RouteStatus: gwRouteStatus(lsParentRef("default", "ls", withSectionName("api"))),
		},
	}
	_, err = gwClient.GatewayV1beta1().HTTPRoutes(rt.Namespace).Create(ctx, rt, metav1.CreateOptions{})
	require.NoError(t, err)

	src, err := NewGatewayHTTPRouteSource(ctx, clients, &Config{})
	require.NoError(t, err)

	endpoints, err := src.Endpoints(ctx)
	require.NoError(t, err)
	validateEndpoints(t, endpoints, []*endpoint.Endpoint{
		newTestEndpoint("api.example.com", "10.0.0.1"),
	})
}

func TestGatewayHTTPRouteWithListenerSetCrossNamespaceRoute(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(t.Context(), 10*time.Second)
	defer cancel()

	gwClient := gatewayfake.NewSimpleClientset()
	kubeClient := kubefake.NewClientset()
	clients := new(MockClientGenerator)
	clients.On("GatewayClient").Return(gwClient, nil)
	clients.On("KubeClient").Return(kubeClient, nil)

	for _, name := range []string{"infra", "apps"} {
		ns := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: name}}
		_, err := kubeClient.CoreV1().Namespaces().Create(ctx, ns, metav1.CreateOptions{})
		require.NoError(t, err)
	}

	// Gateway and ListenerSet in "infra" namespace.
	gw := &v1beta1.Gateway{
		ObjectMeta: metav1.ObjectMeta{Name: "shared-gw", Namespace: "infra"},
		Spec:       v1.GatewaySpec{Listeners: []v1.Listener{{Protocol: v1.HTTPProtocolType, Port: 80}}},
		Status:     gatewayStatus("10.0.0.1"),
	}
	_, err := gwClient.GatewayV1beta1().Gateways(gw.Namespace).Create(ctx, gw, metav1.CreateOptions{})
	require.NoError(t, err)

	hostname := v1.Hostname("app.example.com")
	fromAll := v1.NamespacesFromAll
	ls := &v1.ListenerSet{
		ObjectMeta: metav1.ObjectMeta{Name: "ls", Namespace: "infra"},
		Spec: v1.ListenerSetSpec{
			ParentRef: v1.ParentGatewayReference{Name: "shared-gw"},
			Listeners: []v1.ListenerEntry{{
				Name: "app", Hostname: &hostname, Port: 8080, Protocol: v1.HTTPProtocolType,
				AllowedRoutes: &v1.AllowedRoutes{Namespaces: &v1.RouteNamespaces{From: &fromAll}},
			}},
		},
	}
	_, err = gwClient.GatewayV1().ListenerSets(ls.Namespace).Create(ctx, ls, metav1.CreateOptions{})
	require.NoError(t, err)

	// Route in "apps" namespace referencing the ListenerSet in "infra".
	rt := &v1beta1.HTTPRoute{
		ObjectMeta: metav1.ObjectMeta{Name: "rt", Namespace: "apps"},
		Spec: v1.HTTPRouteSpec{
			Hostnames:       []v1.Hostname{"app.example.com"},
			CommonRouteSpec: v1.CommonRouteSpec{ParentRefs: []v1.ParentReference{lsParentRef("infra", "ls")}},
		},
		Status: v1.HTTPRouteStatus{RouteStatus: gwRouteStatus(lsParentRef("infra", "ls"))},
	}
	_, err = gwClient.GatewayV1beta1().HTTPRoutes(rt.Namespace).Create(ctx, rt, metav1.CreateOptions{})
	require.NoError(t, err)

	src, err := NewGatewayHTTPRouteSource(ctx, clients, &Config{
		GatewayNamespace: "infra",
	})
	require.NoError(t, err)

	endpoints, err := src.Endpoints(ctx)
	require.NoError(t, err)
	validateEndpoints(t, endpoints, []*endpoint.Endpoint{
		newTestEndpoint("app.example.com", "10.0.0.1"),
	})
}

func TestGatewayHTTPRouteWithListenerSetGatewayNotFound(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(t.Context(), 10*time.Second)
	defer cancel()

	gwClient := gatewayfake.NewSimpleClientset()
	kubeClient := kubefake.NewClientset()
	clients := new(MockClientGenerator)
	clients.On("GatewayClient").Return(gwClient, nil)
	clients.On("KubeClient").Return(kubeClient, nil)

	ns := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "default"}}
	_, err := kubeClient.CoreV1().Namespaces().Create(ctx, ns, metav1.CreateOptions{})
	require.NoError(t, err)

	// ListenerSet references a Gateway that doesn't exist.
	hostname := v1.Hostname("app.example.com")
	fromAll := v1.NamespacesFromAll
	ls := &v1.ListenerSet{
		ObjectMeta: metav1.ObjectMeta{Name: "ls", Namespace: "default"},
		Spec: v1.ListenerSetSpec{
			ParentRef: v1.ParentGatewayReference{Name: "nonexistent"},
			Listeners: []v1.ListenerEntry{{
				Name: "app", Hostname: &hostname, Port: 8080, Protocol: v1.HTTPProtocolType,
				AllowedRoutes: &v1.AllowedRoutes{Namespaces: &v1.RouteNamespaces{From: &fromAll}},
			}},
		},
	}
	_, err = gwClient.GatewayV1().ListenerSets(ls.Namespace).Create(ctx, ls, metav1.CreateOptions{})
	require.NoError(t, err)

	rt := &v1beta1.HTTPRoute{
		ObjectMeta: metav1.ObjectMeta{Name: "rt", Namespace: "default"},
		Spec: v1.HTTPRouteSpec{
			Hostnames:       []v1.Hostname{"app.example.com"},
			CommonRouteSpec: v1.CommonRouteSpec{ParentRefs: []v1.ParentReference{lsParentRef("default", "ls")}},
		},
		Status: v1.HTTPRouteStatus{RouteStatus: gwRouteStatus(lsParentRef("default", "ls"))},
	}
	_, err = gwClient.GatewayV1beta1().HTTPRoutes(rt.Namespace).Create(ctx, rt, metav1.CreateOptions{})
	require.NoError(t, err)

	src, err := NewGatewayHTTPRouteSource(ctx, clients, &Config{})
	require.NoError(t, err)

	endpoints, err := src.Endpoints(ctx)
	require.NoError(t, err)
	require.Empty(t, endpoints, "expected no endpoints when ListenerSet's parent Gateway doesn't exist")
}
