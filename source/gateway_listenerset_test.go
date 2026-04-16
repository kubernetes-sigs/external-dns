/*
Copyright 2026 The Kubernetes Authors.

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
	v1 "sigs.k8s.io/gateway-api/apis/v1"
	gatewayfake "sigs.k8s.io/gateway-api/pkg/client/clientset/versioned/fake"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
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

func allowAllListenerSets() *v1.AllowedListeners {
	fromAll := v1.NamespacesFromAll
	return &v1.AllowedListeners{
		Namespaces: &v1.ListenerNamespaces{From: &fromAll},
	}
}

func allowListenerSetsFromSelector(selector *metav1.LabelSelector) *v1.AllowedListeners {
	fromSelector := v1.NamespacesFromSelector
	return &v1.AllowedListeners{
		Namespaces: &v1.ListenerNamespaces{From: &fromSelector, Selector: selector},
	}
}

func conditionWithStatus(conditionType string, status metav1.ConditionStatus) metav1.Condition {
	return metav1.Condition{
		Type:   conditionType,
		Status: status,
	}
}

func acceptedCondition() metav1.Condition {
	return conditionWithStatus("Accepted", metav1.ConditionTrue)
}

func rejectedCondition() metav1.Condition {
	return conditionWithStatus("Accepted", metav1.ConditionFalse)
}

func listenerSetAcceptedStatus(names ...v1.SectionName) v1.ListenerSetStatus {
	listeners := make([]v1.ListenerEntryStatus, 0, len(names))
	for _, name := range names {
		listeners = append(listeners, v1.ListenerEntryStatus{
			Name:       name,
			Conditions: []metav1.Condition{acceptedCondition()},
		})
	}
	return v1.ListenerSetStatus{
		Conditions: []metav1.Condition{acceptedCondition()},
		Listeners:  listeners,
	}
}

func TestGatewayHTTPRouteWithListenerSetParentRef(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(t.Context(), 10*time.Second)
	defer cancel()

	gwClient := gatewayfake.NewSimpleClientset()
	kubeClient := kubefake.NewClientset()
	clients := new(testutils.MockClientGenerator)
	clients.On("GatewayClient").Return(gwClient, nil)
	clients.On("KubeClient").Return(kubeClient, nil)

	ns := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "default"}}
	_, err := kubeClient.CoreV1().Namespaces().Create(ctx, ns, metav1.CreateOptions{})
	require.NoError(t, err)

	ips := []string{"10.64.0.1", "10.64.0.2"}
	gw := &v1.Gateway{
		ObjectMeta: metav1.ObjectMeta{Name: "my-gateway", Namespace: "default"},
		Spec: v1.GatewaySpec{
			AllowedListeners: allowAllListenerSets(),
			Listeners: []v1.Listener{{
				Name:     "base",
				Protocol: v1.HTTPProtocolType,
				Port:     80,
			}},
		},
		Status: gatewayStatus(ips...),
	}
	_, err = gwClient.GatewayV1().Gateways(gw.Namespace).Create(ctx, gw, metav1.CreateOptions{})
	require.NoError(t, err)

	hostname := v1.Hostname("*.example.com")
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
		Status: listenerSetAcceptedStatus("app"),
	}
	_, err = gwClient.GatewayV1().ListenerSets(ls.Namespace).Create(ctx, ls, metav1.CreateOptions{})
	require.NoError(t, err)

	rt := &v1.HTTPRoute{
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
	_, err = gwClient.GatewayV1().HTTPRoutes(rt.Namespace).Create(ctx, rt, metav1.CreateOptions{})
	require.NoError(t, err)

	src, err := NewGatewayHTTPRouteSource(ctx, clients, &Config{GatewayListenerSets: true})
	require.NoError(t, err)

	endpoints, err := src.Endpoints(ctx)
	require.NoError(t, err)
	testutils.ValidateEndpoints(t, endpoints, []*endpoint.Endpoint{
		newTestEndpoint("app.example.com", ips...),
	})
}

func TestGatewayHTTPRouteWithListenerSetWildcardHostnameIntersection(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(t.Context(), 10*time.Second)
	defer cancel()

	gwClient := gatewayfake.NewSimpleClientset()
	kubeClient := kubefake.NewClientset()
	clients := new(testutils.MockClientGenerator)
	clients.On("GatewayClient").Return(gwClient, nil)
	clients.On("KubeClient").Return(kubeClient, nil)

	ns := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "default"}}
	_, err := kubeClient.CoreV1().Namespaces().Create(ctx, ns, metav1.CreateOptions{})
	require.NoError(t, err)

	ips := []string{"10.64.0.1", "10.64.0.2"}
	gw := &v1.Gateway{
		ObjectMeta: metav1.ObjectMeta{Name: "my-gateway", Namespace: "default"},
		Spec: v1.GatewaySpec{
			AllowedListeners: allowAllListenerSets(),
			Listeners: []v1.Listener{{
				Name:     "base",
				Protocol: v1.HTTPProtocolType,
				Port:     80,
			}},
		},
		Status: gatewayStatus(ips...),
	}
	_, err = gwClient.GatewayV1().Gateways(gw.Namespace).Create(ctx, gw, metav1.CreateOptions{})
	require.NoError(t, err)

	hostname := v1.Hostname("*.example.com")
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
		Status: listenerSetAcceptedStatus("app"),
	}
	_, err = gwClient.GatewayV1().ListenerSets(ls.Namespace).Create(ctx, ls, metav1.CreateOptions{})
	require.NoError(t, err)

	rt := &v1.HTTPRoute{
		ObjectMeta: metav1.ObjectMeta{Name: "my-route", Namespace: "default"},
		Spec: v1.HTTPRouteSpec{
			Hostnames: []v1.Hostname{"sub.domain.example.com"},
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
	_, err = gwClient.GatewayV1().HTTPRoutes(rt.Namespace).Create(ctx, rt, metav1.CreateOptions{})
	require.NoError(t, err)

	src, err := NewGatewayHTTPRouteSource(ctx, clients, &Config{GatewayListenerSets: true})
	require.NoError(t, err)

	endpoints, err := src.Endpoints(ctx)
	require.NoError(t, err)
	testutils.ValidateEndpoints(t, endpoints, []*endpoint.Endpoint{
		newTestEndpoint("sub.domain.example.com", ips...),
	})
}

func TestGatewayHTTPRouteWithListenerSetDisabled(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(t.Context(), 10*time.Second)
	defer cancel()

	gwClient := gatewayfake.NewSimpleClientset()
	kubeClient := kubefake.NewClientset()
	clients := new(testutils.MockClientGenerator)
	clients.On("GatewayClient").Return(gwClient, nil)
	clients.On("KubeClient").Return(kubeClient, nil)

	ns := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "default"}}
	_, err := kubeClient.CoreV1().Namespaces().Create(ctx, ns, metav1.CreateOptions{})
	require.NoError(t, err)

	gw := &v1.Gateway{
		ObjectMeta: metav1.ObjectMeta{Name: "gw", Namespace: "default"},
		Spec:       v1.GatewaySpec{AllowedListeners: allowAllListenerSets(), Listeners: []v1.Listener{{Protocol: v1.HTTPProtocolType, Port: 80}}},
		Status:     gatewayStatus("10.0.0.1"),
	}
	_, err = gwClient.GatewayV1().Gateways(gw.Namespace).Create(ctx, gw, metav1.CreateOptions{})
	require.NoError(t, err)

	hostname := v1.Hostname("*.example.com")
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
		Status: listenerSetAcceptedStatus("app"),
	}
	_, err = gwClient.GatewayV1().ListenerSets(ls.Namespace).Create(ctx, ls, metav1.CreateOptions{})
	require.NoError(t, err)

	// Route with ListenerSet parentRef, but feature is disabled.
	rt := &v1.HTTPRoute{
		ObjectMeta: metav1.ObjectMeta{Name: "rt", Namespace: "default"},
		Spec: v1.HTTPRouteSpec{
			Hostnames:       []v1.Hostname{"app.example.com"},
			CommonRouteSpec: v1.CommonRouteSpec{ParentRefs: []v1.ParentReference{lsParentRef("default", "ls")}},
		},
		Status: v1.HTTPRouteStatus{RouteStatus: gwRouteStatus(lsParentRef("default", "ls"))},
	}
	_, err = gwClient.GatewayV1().HTTPRoutes(rt.Namespace).Create(ctx, rt, metav1.CreateOptions{})
	require.NoError(t, err)

	src, err := NewGatewayHTTPRouteSource(ctx, clients, &Config{GatewayListenerSets: false})
	require.NoError(t, err)

	endpoints, err := src.Endpoints(ctx)
	require.NoError(t, err)
	require.Empty(t, endpoints, "expected no endpoints when ListenerSet support is disabled")
}

func TestGatewayHTTPRouteWithListenerSetGatewayLabelFilter(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(t.Context(), 10*time.Second)
	defer cancel()

	gwClient := gatewayfake.NewSimpleClientset()
	kubeClient := kubefake.NewClientset()
	clients := new(testutils.MockClientGenerator)
	clients.On("GatewayClient").Return(gwClient, nil)
	clients.On("KubeClient").Return(kubeClient, nil)

	ns := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "default"}}
	_, err := kubeClient.CoreV1().Namespaces().Create(ctx, ns, metav1.CreateOptions{})
	require.NoError(t, err)

	for _, gw := range []*v1.Gateway{
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "labels-match",
				Namespace: "default",
				Labels:    map[string]string{"foo": "bar"},
			},
			Spec:   v1.GatewaySpec{AllowedListeners: allowAllListenerSets(), Listeners: []v1.Listener{{Protocol: v1.HTTPProtocolType, Port: 80}}},
			Status: gatewayStatus("10.0.0.1"),
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "labels-dont-match",
				Namespace: "default",
				Labels:    map[string]string{"foo": "qux"},
			},
			Spec:   v1.GatewaySpec{AllowedListeners: allowAllListenerSets(), Listeners: []v1.Listener{{Protocol: v1.HTTPProtocolType, Port: 80}}},
			Status: gatewayStatus("10.0.0.2"),
		},
	} {
		_, err = gwClient.GatewayV1().Gateways(gw.Namespace).Create(ctx, gw, metav1.CreateOptions{})
		require.NoError(t, err)
	}

	hostname := v1.Hostname("app.example.com")
	fromAll := v1.NamespacesFromAll
	for _, ls := range []*v1.ListenerSet{
		{
			ObjectMeta: metav1.ObjectMeta{Name: "ls-match", Namespace: "default"},
			Spec: v1.ListenerSetSpec{
				ParentRef: v1.ParentGatewayReference{Name: "labels-match"},
				Listeners: []v1.ListenerEntry{{
					Name: "app", Hostname: &hostname, Port: 8080, Protocol: v1.HTTPProtocolType,
					AllowedRoutes: &v1.AllowedRoutes{Namespaces: &v1.RouteNamespaces{From: &fromAll}},
				}},
			},
			Status: listenerSetAcceptedStatus("app"),
		},
		{
			ObjectMeta: metav1.ObjectMeta{Name: "ls-dont-match", Namespace: "default"},
			Spec: v1.ListenerSetSpec{
				ParentRef: v1.ParentGatewayReference{Name: "labels-dont-match"},
				Listeners: []v1.ListenerEntry{{
					Name: "app", Hostname: &hostname, Port: 8080, Protocol: v1.HTTPProtocolType,
					AllowedRoutes: &v1.AllowedRoutes{Namespaces: &v1.RouteNamespaces{From: &fromAll}},
				}},
			},
			Status: listenerSetAcceptedStatus("app"),
		},
	} {
		_, err = gwClient.GatewayV1().ListenerSets(ls.Namespace).Create(ctx, ls, metav1.CreateOptions{})
		require.NoError(t, err)
	}

	rt := &v1.HTTPRoute{
		ObjectMeta: metav1.ObjectMeta{Name: "rt", Namespace: "default"},
		Spec: v1.HTTPRouteSpec{
			Hostnames: []v1.Hostname{"app.example.com"},
			CommonRouteSpec: v1.CommonRouteSpec{
				ParentRefs: []v1.ParentReference{
					lsParentRef("default", "ls-match"),
					lsParentRef("default", "ls-dont-match"),
				},
			},
		},
		Status: v1.HTTPRouteStatus{
			RouteStatus: gwRouteStatus(
				lsParentRef("default", "ls-match"),
				lsParentRef("default", "ls-dont-match"),
			),
		},
	}
	_, err = gwClient.GatewayV1().HTTPRoutes(rt.Namespace).Create(ctx, rt, metav1.CreateOptions{})
	require.NoError(t, err)

	src, err := NewGatewayHTTPRouteSource(ctx, clients, &Config{
		GatewayLabelFilter:  "foo=bar",
		GatewayListenerSets: true,
	})
	require.NoError(t, err)

	endpoints, err := src.Endpoints(ctx)
	require.NoError(t, err)
	testutils.ValidateEndpoints(t, endpoints, []*endpoint.Endpoint{
		newTestEndpoint("app.example.com", "10.0.0.1"),
	})
}

func TestGatewayHTTPRouteWithListenerSetRouteLabelFilter(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(t.Context(), 10*time.Second)
	defer cancel()

	gwClient := gatewayfake.NewSimpleClientset()
	kubeClient := kubefake.NewClientset()
	clients := new(testutils.MockClientGenerator)
	clients.On("GatewayClient").Return(gwClient, nil)
	clients.On("KubeClient").Return(kubeClient, nil)

	ns := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "default"}}
	_, err := kubeClient.CoreV1().Namespaces().Create(ctx, ns, metav1.CreateOptions{})
	require.NoError(t, err)

	gw := &v1.Gateway{
		ObjectMeta: metav1.ObjectMeta{Name: "gw", Namespace: "default"},
		Spec:       v1.GatewaySpec{AllowedListeners: allowAllListenerSets(), Listeners: []v1.Listener{{Protocol: v1.HTTPProtocolType, Port: 80}}},
		Status:     gatewayStatus("10.0.0.1"),
	}
	_, err = gwClient.GatewayV1().Gateways(gw.Namespace).Create(ctx, gw, metav1.CreateOptions{})
	require.NoError(t, err)

	hostname := v1.Hostname("*.example.com")
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
		Status: listenerSetAcceptedStatus("app"),
	}
	_, err = gwClient.GatewayV1().ListenerSets(ls.Namespace).Create(ctx, ls, metav1.CreateOptions{})
	require.NoError(t, err)

	for _, rt := range []*v1.HTTPRoute{
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "labels-match",
				Namespace: "default",
				Labels:    map[string]string{"foo": "bar"},
			},
			Spec: v1.HTTPRouteSpec{
				Hostnames:       []v1.Hostname{"labels-match.example.com"},
				CommonRouteSpec: v1.CommonRouteSpec{ParentRefs: []v1.ParentReference{lsParentRef("default", "ls")}},
			},
			Status: v1.HTTPRouteStatus{RouteStatus: gwRouteStatus(lsParentRef("default", "ls"))},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "labels-dont-match",
				Namespace: "default",
				Labels:    map[string]string{"foo": "qux"},
			},
			Spec: v1.HTTPRouteSpec{
				Hostnames:       []v1.Hostname{"labels-dont-match.example.com"},
				CommonRouteSpec: v1.CommonRouteSpec{ParentRefs: []v1.ParentReference{lsParentRef("default", "ls")}},
			},
			Status: v1.HTTPRouteStatus{RouteStatus: gwRouteStatus(lsParentRef("default", "ls"))},
		},
	} {
		_, err = gwClient.GatewayV1().HTTPRoutes(rt.Namespace).Create(ctx, rt, metav1.CreateOptions{})
		require.NoError(t, err)
	}

	src, err := NewGatewayHTTPRouteSource(ctx, clients, &Config{
		LabelFilter:         mustGetLabelSelector("foo=bar"),
		GatewayListenerSets: true,
	})
	require.NoError(t, err)

	endpoints, err := src.Endpoints(ctx)
	require.NoError(t, err)
	testutils.ValidateEndpoints(t, endpoints, []*endpoint.Endpoint{
		newTestEndpoint("labels-match.example.com", "10.0.0.1"),
	})
}

func TestGatewayHTTPRouteWithListenerSetNotAccepted(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(t.Context(), 10*time.Second)
	defer cancel()

	gwClient := gatewayfake.NewSimpleClientset()
	kubeClient := kubefake.NewClientset()
	clients := new(testutils.MockClientGenerator)
	clients.On("GatewayClient").Return(gwClient, nil)
	clients.On("KubeClient").Return(kubeClient, nil)

	ns := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "default"}}
	_, err := kubeClient.CoreV1().Namespaces().Create(ctx, ns, metav1.CreateOptions{})
	require.NoError(t, err)

	gw := &v1.Gateway{
		ObjectMeta: metav1.ObjectMeta{Name: "gw", Namespace: "default"},
		Spec:       v1.GatewaySpec{AllowedListeners: allowAllListenerSets(), Listeners: []v1.Listener{{Protocol: v1.HTTPProtocolType, Port: 80}}},
		Status:     gatewayStatus("10.0.0.1"),
	}
	_, err = gwClient.GatewayV1().Gateways(gw.Namespace).Create(ctx, gw, metav1.CreateOptions{})
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
		Status: listenerSetAcceptedStatus("app"),
	}
	_, err = gwClient.GatewayV1().ListenerSets(ls.Namespace).Create(ctx, ls, metav1.CreateOptions{})
	require.NoError(t, err)

	// Route with ListenerSet parentRef, but NOT accepted (no Accepted condition).
	rt := &v1.HTTPRoute{
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
	_, err = gwClient.GatewayV1().HTTPRoutes(rt.Namespace).Create(ctx, rt, metav1.CreateOptions{})
	require.NoError(t, err)

	src, err := NewGatewayHTTPRouteSource(ctx, clients, &Config{GatewayListenerSets: true})
	require.NoError(t, err)

	endpoints, err := src.Endpoints(ctx)
	require.NoError(t, err)
	require.Empty(t, endpoints, "expected no endpoints when route is not accepted by ListenerSet")
}

func TestGatewayHTTPRouteWithListenerSetNotAllowedByGateway(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(t.Context(), 10*time.Second)
	defer cancel()

	gwClient := gatewayfake.NewSimpleClientset()
	kubeClient := kubefake.NewClientset()
	clients := new(testutils.MockClientGenerator)
	clients.On("GatewayClient").Return(gwClient, nil)
	clients.On("KubeClient").Return(kubeClient, nil)

	ns := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "default"}}
	_, err := kubeClient.CoreV1().Namespaces().Create(ctx, ns, metav1.CreateOptions{})
	require.NoError(t, err)

	gw := &v1.Gateway{
		ObjectMeta: metav1.ObjectMeta{Name: "gw", Namespace: "default"},
		Spec:       v1.GatewaySpec{Listeners: []v1.Listener{{Protocol: v1.HTTPProtocolType, Port: 80}}},
		Status:     gatewayStatus("10.0.0.1"),
	}
	_, err = gwClient.GatewayV1().Gateways(gw.Namespace).Create(ctx, gw, metav1.CreateOptions{})
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
		Status: listenerSetAcceptedStatus("app"),
	}
	_, err = gwClient.GatewayV1().ListenerSets(ls.Namespace).Create(ctx, ls, metav1.CreateOptions{})
	require.NoError(t, err)

	rt := &v1.HTTPRoute{
		ObjectMeta: metav1.ObjectMeta{Name: "rt", Namespace: "default"},
		Spec: v1.HTTPRouteSpec{
			Hostnames:       []v1.Hostname{"app.example.com"},
			CommonRouteSpec: v1.CommonRouteSpec{ParentRefs: []v1.ParentReference{lsParentRef("default", "ls")}},
		},
		Status: v1.HTTPRouteStatus{RouteStatus: gwRouteStatus(lsParentRef("default", "ls"))},
	}
	_, err = gwClient.GatewayV1().HTTPRoutes(rt.Namespace).Create(ctx, rt, metav1.CreateOptions{})
	require.NoError(t, err)

	src, err := NewGatewayHTTPRouteSource(ctx, clients, &Config{GatewayListenerSets: true})
	require.NoError(t, err)

	endpoints, err := src.Endpoints(ctx)
	require.NoError(t, err)
	require.Empty(t, endpoints, "expected no endpoints when the Gateway does not allow ListenerSets")
}

func TestGatewayHTTPRouteWithListenerSetStatusNotAccepted(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(t.Context(), 10*time.Second)
	defer cancel()

	gwClient := gatewayfake.NewSimpleClientset()
	kubeClient := kubefake.NewClientset()
	clients := new(testutils.MockClientGenerator)
	clients.On("GatewayClient").Return(gwClient, nil)
	clients.On("KubeClient").Return(kubeClient, nil)

	ns := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "default"}}
	_, err := kubeClient.CoreV1().Namespaces().Create(ctx, ns, metav1.CreateOptions{})
	require.NoError(t, err)

	gw := &v1.Gateway{
		ObjectMeta: metav1.ObjectMeta{Name: "gw", Namespace: "default"},
		Spec:       v1.GatewaySpec{AllowedListeners: allowAllListenerSets(), Listeners: []v1.Listener{{Protocol: v1.HTTPProtocolType, Port: 80}}},
		Status:     gatewayStatus("10.0.0.1"),
	}
	_, err = gwClient.GatewayV1().Gateways(gw.Namespace).Create(ctx, gw, metav1.CreateOptions{})
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

	rt := &v1.HTTPRoute{
		ObjectMeta: metav1.ObjectMeta{Name: "rt", Namespace: "default"},
		Spec: v1.HTTPRouteSpec{
			Hostnames:       []v1.Hostname{"app.example.com"},
			CommonRouteSpec: v1.CommonRouteSpec{ParentRefs: []v1.ParentReference{lsParentRef("default", "ls")}},
		},
		Status: v1.HTTPRouteStatus{RouteStatus: gwRouteStatus(lsParentRef("default", "ls"))},
	}
	_, err = gwClient.GatewayV1().HTTPRoutes(rt.Namespace).Create(ctx, rt, metav1.CreateOptions{})
	require.NoError(t, err)

	src, err := NewGatewayHTTPRouteSource(ctx, clients, &Config{GatewayListenerSets: true})
	require.NoError(t, err)

	endpoints, err := src.Endpoints(ctx)
	require.NoError(t, err)
	require.Empty(t, endpoints, "expected no endpoints when the ListenerSet itself is not accepted")
}

func TestGatewayHTTPRouteWithListenerSetListenerStatusRequired(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(t.Context(), 10*time.Second)
	defer cancel()

	gwClient := gatewayfake.NewSimpleClientset()
	kubeClient := kubefake.NewClientset()
	clients := new(testutils.MockClientGenerator)
	clients.On("GatewayClient").Return(gwClient, nil)
	clients.On("KubeClient").Return(kubeClient, nil)

	ns := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "default"}}
	_, err := kubeClient.CoreV1().Namespaces().Create(ctx, ns, metav1.CreateOptions{})
	require.NoError(t, err)

	gw := &v1.Gateway{
		ObjectMeta: metav1.ObjectMeta{Name: "gw", Namespace: "default"},
		Spec:       v1.GatewaySpec{AllowedListeners: allowAllListenerSets(), Listeners: []v1.Listener{{Protocol: v1.HTTPProtocolType, Port: 80}}},
		Status:     gatewayStatus("10.0.0.1"),
	}
	_, err = gwClient.GatewayV1().Gateways(gw.Namespace).Create(ctx, gw, metav1.CreateOptions{})
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
		Status: v1.ListenerSetStatus{
			Conditions: []metav1.Condition{acceptedCondition()},
		},
	}
	_, err = gwClient.GatewayV1().ListenerSets(ls.Namespace).Create(ctx, ls, metav1.CreateOptions{})
	require.NoError(t, err)

	rtRef := lsParentRef("default", "ls", withSectionName("app"))
	rt := &v1.HTTPRoute{
		ObjectMeta: metav1.ObjectMeta{Name: "rt", Namespace: "default"},
		Spec: v1.HTTPRouteSpec{
			Hostnames:       []v1.Hostname{"app.example.com"},
			CommonRouteSpec: v1.CommonRouteSpec{ParentRefs: []v1.ParentReference{rtRef}},
		},
		Status: v1.HTTPRouteStatus{RouteStatus: gwRouteStatus(rtRef)},
	}
	_, err = gwClient.GatewayV1().HTTPRoutes(rt.Namespace).Create(ctx, rt, metav1.CreateOptions{})
	require.NoError(t, err)

	src, err := NewGatewayHTTPRouteSource(ctx, clients, &Config{GatewayListenerSets: true})
	require.NoError(t, err)

	endpoints, err := src.Endpoints(ctx)
	require.NoError(t, err)
	require.Empty(t, endpoints, "expected no endpoints when the ListenerSet listener has no Accepted status")
}

func TestGatewayHTTPRouteWithListenerSetListenerStatusNotAccepted(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(t.Context(), 10*time.Second)
	defer cancel()

	gwClient := gatewayfake.NewSimpleClientset()
	kubeClient := kubefake.NewClientset()
	clients := new(testutils.MockClientGenerator)
	clients.On("GatewayClient").Return(gwClient, nil)
	clients.On("KubeClient").Return(kubeClient, nil)

	ns := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "default"}}
	_, err := kubeClient.CoreV1().Namespaces().Create(ctx, ns, metav1.CreateOptions{})
	require.NoError(t, err)

	gw := &v1.Gateway{
		ObjectMeta: metav1.ObjectMeta{Name: "gw", Namespace: "default"},
		Spec:       v1.GatewaySpec{AllowedListeners: allowAllListenerSets(), Listeners: []v1.Listener{{Protocol: v1.HTTPProtocolType, Port: 80}}},
		Status:     gatewayStatus("10.0.0.1"),
	}
	_, err = gwClient.GatewayV1().Gateways(gw.Namespace).Create(ctx, gw, metav1.CreateOptions{})
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
		Status: v1.ListenerSetStatus{
			Conditions: []metav1.Condition{acceptedCondition()},
			Listeners: []v1.ListenerEntryStatus{{
				Name:       "app",
				Conditions: []metav1.Condition{rejectedCondition()},
			}},
		},
	}
	_, err = gwClient.GatewayV1().ListenerSets(ls.Namespace).Create(ctx, ls, metav1.CreateOptions{})
	require.NoError(t, err)

	rtRef := lsParentRef("default", "ls", withSectionName("app"))
	rt := &v1.HTTPRoute{
		ObjectMeta: metav1.ObjectMeta{Name: "rt", Namespace: "default"},
		Spec: v1.HTTPRouteSpec{
			Hostnames:       []v1.Hostname{"app.example.com"},
			CommonRouteSpec: v1.CommonRouteSpec{ParentRefs: []v1.ParentReference{rtRef}},
		},
		Status: v1.HTTPRouteStatus{RouteStatus: gwRouteStatus(rtRef)},
	}
	_, err = gwClient.GatewayV1().HTTPRoutes(rt.Namespace).Create(ctx, rt, metav1.CreateOptions{})
	require.NoError(t, err)

	src, err := NewGatewayHTTPRouteSource(ctx, clients, &Config{GatewayListenerSets: true})
	require.NoError(t, err)

	endpoints, err := src.Endpoints(ctx)
	require.NoError(t, err)
	require.Empty(t, endpoints, "expected no endpoints when the ListenerSet listener is explicitly not accepted")
}

func TestGatewayHTTPRouteWithListenerSetTargetAnnotation(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(t.Context(), 10*time.Second)
	defer cancel()

	gwClient := gatewayfake.NewSimpleClientset()
	kubeClient := kubefake.NewClientset()
	clients := new(testutils.MockClientGenerator)
	clients.On("GatewayClient").Return(gwClient, nil)
	clients.On("KubeClient").Return(kubeClient, nil)

	ns := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "default"}}
	_, err := kubeClient.CoreV1().Namespaces().Create(ctx, ns, metav1.CreateOptions{})
	require.NoError(t, err)

	// Gateway with target annotation override.
	gw := &v1.Gateway{
		ObjectMeta: metav1.ObjectMeta{
			Name: "gw", Namespace: "default",
			Annotations: map[string]string{"external-dns.alpha.kubernetes.io/target": "override.example.com"},
		},
		Spec:   v1.GatewaySpec{AllowedListeners: allowAllListenerSets(), Listeners: []v1.Listener{{Protocol: v1.HTTPProtocolType, Port: 80}}},
		Status: gatewayStatus("10.0.0.1"),
	}
	_, err = gwClient.GatewayV1().Gateways(gw.Namespace).Create(ctx, gw, metav1.CreateOptions{})
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
		Status: listenerSetAcceptedStatus("app"),
	}
	_, err = gwClient.GatewayV1().ListenerSets(ls.Namespace).Create(ctx, ls, metav1.CreateOptions{})
	require.NoError(t, err)

	rt := &v1.HTTPRoute{
		ObjectMeta: metav1.ObjectMeta{Name: "rt", Namespace: "default"},
		Spec: v1.HTTPRouteSpec{
			Hostnames:       []v1.Hostname{"app.example.com"},
			CommonRouteSpec: v1.CommonRouteSpec{ParentRefs: []v1.ParentReference{lsParentRef("default", "ls")}},
		},
		Status: v1.HTTPRouteStatus{RouteStatus: gwRouteStatus(lsParentRef("default", "ls"))},
	}
	_, err = gwClient.GatewayV1().HTTPRoutes(rt.Namespace).Create(ctx, rt, metav1.CreateOptions{})
	require.NoError(t, err)

	src, err := NewGatewayHTTPRouteSource(ctx, clients, &Config{GatewayListenerSets: true})
	require.NoError(t, err)

	endpoints, err := src.Endpoints(ctx)
	require.NoError(t, err)
	testutils.ValidateEndpoints(t, endpoints, []*endpoint.Endpoint{
		newTestEndpointWithTTL("app.example.com", endpoint.RecordTypeCNAME, 0, "override.example.com"),
	})
}

func TestGatewayHTTPRouteWithListenerSetAllowedRoutesSame(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(t.Context(), 10*time.Second)
	defer cancel()

	gwClient := gatewayfake.NewSimpleClientset()
	kubeClient := kubefake.NewClientset()
	clients := new(testutils.MockClientGenerator)
	clients.On("GatewayClient").Return(gwClient, nil)
	clients.On("KubeClient").Return(kubeClient, nil)

	for _, name := range []string{"default", "other"} {
		ns := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: name}}
		_, err := kubeClient.CoreV1().Namespaces().Create(ctx, ns, metav1.CreateOptions{})
		require.NoError(t, err)
	}

	gw := &v1.Gateway{
		ObjectMeta: metav1.ObjectMeta{Name: "gw", Namespace: "default"},
		Spec:       v1.GatewaySpec{AllowedListeners: allowAllListenerSets(), Listeners: []v1.Listener{{Protocol: v1.HTTPProtocolType, Port: 80}}},
		Status:     gatewayStatus("10.0.0.1"),
	}
	_, err := gwClient.GatewayV1().Gateways(gw.Namespace).Create(ctx, gw, metav1.CreateOptions{})
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
		Status: listenerSetAcceptedStatus("app"),
	}
	_, err = gwClient.GatewayV1().ListenerSets(ls.Namespace).Create(ctx, ls, metav1.CreateOptions{})
	require.NoError(t, err)

	// Route in "other" namespace → should NOT match (Same means ListenerSet namespace).
	rt := &v1.HTTPRoute{
		ObjectMeta: metav1.ObjectMeta{Name: "rt", Namespace: "other"},
		Spec: v1.HTTPRouteSpec{
			Hostnames:       []v1.Hostname{"app.example.com"},
			CommonRouteSpec: v1.CommonRouteSpec{ParentRefs: []v1.ParentReference{lsParentRef("default", "ls")}},
		},
		Status: v1.HTTPRouteStatus{RouteStatus: gwRouteStatus(lsParentRef("default", "ls"))},
	}
	_, err = gwClient.GatewayV1().HTTPRoutes(rt.Namespace).Create(ctx, rt, metav1.CreateOptions{})
	require.NoError(t, err)

	src, err := NewGatewayHTTPRouteSource(ctx, clients, &Config{GatewayListenerSets: true})
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
	clients := new(testutils.MockClientGenerator)
	clients.On("GatewayClient").Return(gwClient, nil)
	clients.On("KubeClient").Return(kubeClient, nil)

	ns := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "default"}}
	_, err := kubeClient.CoreV1().Namespaces().Create(ctx, ns, metav1.CreateOptions{})
	require.NoError(t, err)

	gw := &v1.Gateway{
		ObjectMeta: metav1.ObjectMeta{Name: "gw", Namespace: "default"},
		Spec:       v1.GatewaySpec{AllowedListeners: allowAllListenerSets(), Listeners: []v1.Listener{{Protocol: v1.HTTPProtocolType, Port: 80}}},
		Status:     gatewayStatus("10.0.0.1"),
	}
	_, err = gwClient.GatewayV1().Gateways(gw.Namespace).Create(ctx, gw, metav1.CreateOptions{})
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
		Status: listenerSetAcceptedStatus("app", "api"),
	}
	_, err = gwClient.GatewayV1().ListenerSets(ls.Namespace).Create(ctx, ls, metav1.CreateOptions{})
	require.NoError(t, err)

	// Route targeting only the "api" section of the ListenerSet.
	rt := &v1.HTTPRoute{
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
	_, err = gwClient.GatewayV1().HTTPRoutes(rt.Namespace).Create(ctx, rt, metav1.CreateOptions{})
	require.NoError(t, err)

	src, err := NewGatewayHTTPRouteSource(ctx, clients, &Config{GatewayListenerSets: true})
	require.NoError(t, err)

	endpoints, err := src.Endpoints(ctx)
	require.NoError(t, err)
	testutils.ValidateEndpoints(t, endpoints, []*endpoint.Endpoint{
		newTestEndpoint("api.example.com", "10.0.0.1"),
	})
}

func TestGatewayHTTPRouteWithListenerSetCrossNamespaceRoute(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(t.Context(), 10*time.Second)
	defer cancel()

	gwClient := gatewayfake.NewSimpleClientset()
	kubeClient := kubefake.NewClientset()
	clients := new(testutils.MockClientGenerator)
	clients.On("GatewayClient").Return(gwClient, nil)
	clients.On("KubeClient").Return(kubeClient, nil)

	for _, name := range []string{"infra", "apps"} {
		ns := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: name}}
		_, err := kubeClient.CoreV1().Namespaces().Create(ctx, ns, metav1.CreateOptions{})
		require.NoError(t, err)
	}

	// Gateway and ListenerSet in "infra" namespace.
	gw := &v1.Gateway{
		ObjectMeta: metav1.ObjectMeta{Name: "shared-gw", Namespace: "infra"},
		Spec:       v1.GatewaySpec{AllowedListeners: allowAllListenerSets(), Listeners: []v1.Listener{{Protocol: v1.HTTPProtocolType, Port: 80}}},
		Status:     gatewayStatus("10.0.0.1"),
	}
	_, err := gwClient.GatewayV1().Gateways(gw.Namespace).Create(ctx, gw, metav1.CreateOptions{})
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
		Status: listenerSetAcceptedStatus("app"),
	}
	_, err = gwClient.GatewayV1().ListenerSets(ls.Namespace).Create(ctx, ls, metav1.CreateOptions{})
	require.NoError(t, err)

	// Route in "apps" namespace referencing the ListenerSet in "infra".
	rt := &v1.HTTPRoute{
		ObjectMeta: metav1.ObjectMeta{Name: "rt", Namespace: "apps"},
		Spec: v1.HTTPRouteSpec{
			Hostnames:       []v1.Hostname{"app.example.com"},
			CommonRouteSpec: v1.CommonRouteSpec{ParentRefs: []v1.ParentReference{lsParentRef("infra", "ls")}},
		},
		Status: v1.HTTPRouteStatus{RouteStatus: gwRouteStatus(lsParentRef("infra", "ls"))},
	}
	_, err = gwClient.GatewayV1().HTTPRoutes(rt.Namespace).Create(ctx, rt, metav1.CreateOptions{})
	require.NoError(t, err)

	src, err := NewGatewayHTTPRouteSource(ctx, clients, &Config{
		GatewayNamespace:    "infra",
		GatewayListenerSets: true,
	})
	require.NoError(t, err)

	endpoints, err := src.Endpoints(ctx)
	require.NoError(t, err)
	testutils.ValidateEndpoints(t, endpoints, []*endpoint.Endpoint{
		newTestEndpoint("app.example.com", "10.0.0.1"),
	})
}

func TestGatewayHTTPRouteWithCrossNamespaceListenerSetSelectedByGateway(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(t.Context(), 10*time.Second)
	defer cancel()

	gwClient := gatewayfake.NewSimpleClientset()
	kubeClient := kubefake.NewClientset()
	clients := new(testutils.MockClientGenerator)
	clients.On("GatewayClient").Return(gwClient, nil)
	clients.On("KubeClient").Return(kubeClient, nil)

	for _, ns := range []*corev1.Namespace{
		{ObjectMeta: metav1.ObjectMeta{Name: "infra"}},
		{ObjectMeta: metav1.ObjectMeta{Name: "edge", Labels: map[string]string{"listener-set": "enabled"}}},
		{ObjectMeta: metav1.ObjectMeta{Name: "apps"}},
	} {
		_, err := kubeClient.CoreV1().Namespaces().Create(ctx, ns, metav1.CreateOptions{})
		require.NoError(t, err)
	}

	selector := &metav1.LabelSelector{MatchLabels: map[string]string{"listener-set": "enabled"}}
	gw := &v1.Gateway{
		ObjectMeta: metav1.ObjectMeta{Name: "shared-gw", Namespace: "infra"},
		Spec: v1.GatewaySpec{
			AllowedListeners: allowListenerSetsFromSelector(selector),
			Listeners:        []v1.Listener{{Protocol: v1.HTTPProtocolType, Port: 80}},
		},
		Status: gatewayStatus("10.0.0.1"),
	}
	_, err := gwClient.GatewayV1().Gateways(gw.Namespace).Create(ctx, gw, metav1.CreateOptions{})
	require.NoError(t, err)

	hostname := v1.Hostname("app.example.com")
	fromAll := v1.NamespacesFromAll
	parentNamespace := v1.Namespace("infra")
	ls := &v1.ListenerSet{
		ObjectMeta: metav1.ObjectMeta{Name: "edge-ls", Namespace: "edge"},
		Spec: v1.ListenerSetSpec{
			ParentRef: v1.ParentGatewayReference{Name: "shared-gw", Namespace: &parentNamespace},
			Listeners: []v1.ListenerEntry{{
				Name: "app", Hostname: &hostname, Port: 8080, Protocol: v1.HTTPProtocolType,
				AllowedRoutes: &v1.AllowedRoutes{Namespaces: &v1.RouteNamespaces{From: &fromAll}},
			}},
		},
		Status: listenerSetAcceptedStatus("app"),
	}
	_, err = gwClient.GatewayV1().ListenerSets(ls.Namespace).Create(ctx, ls, metav1.CreateOptions{})
	require.NoError(t, err)

	rt := &v1.HTTPRoute{
		ObjectMeta: metav1.ObjectMeta{Name: "rt", Namespace: "apps"},
		Spec: v1.HTTPRouteSpec{
			Hostnames:       []v1.Hostname{"app.example.com"},
			CommonRouteSpec: v1.CommonRouteSpec{ParentRefs: []v1.ParentReference{lsParentRef("edge", "edge-ls")}},
		},
		Status: v1.HTTPRouteStatus{RouteStatus: gwRouteStatus(lsParentRef("edge", "edge-ls"))},
	}
	_, err = gwClient.GatewayV1().HTTPRoutes(rt.Namespace).Create(ctx, rt, metav1.CreateOptions{})
	require.NoError(t, err)

	src, err := NewGatewayHTTPRouteSource(ctx, clients, &Config{
		GatewayListenerSets: true,
	})
	require.NoError(t, err)

	endpoints, err := src.Endpoints(ctx)
	require.NoError(t, err)
	testutils.ValidateEndpoints(t, endpoints, []*endpoint.Endpoint{
		newTestEndpoint("app.example.com", "10.0.0.1"),
	})
}

func TestGatewayHTTPRouteWithCrossNamespaceListenerSetAllowedWhenGatewayNamespaceSet(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(t.Context(), 10*time.Second)
	defer cancel()

	gwClient := gatewayfake.NewSimpleClientset()
	kubeClient := kubefake.NewClientset()
	clients := new(testutils.MockClientGenerator)
	clients.On("GatewayClient").Return(gwClient, nil)
	clients.On("KubeClient").Return(kubeClient, nil)

	for _, ns := range []*corev1.Namespace{
		{ObjectMeta: metav1.ObjectMeta{Name: "infra"}},
		{ObjectMeta: metav1.ObjectMeta{Name: "edge", Labels: map[string]string{"listener-set": "enabled"}}},
		{ObjectMeta: metav1.ObjectMeta{Name: "apps"}},
	} {
		_, err := kubeClient.CoreV1().Namespaces().Create(ctx, ns, metav1.CreateOptions{})
		require.NoError(t, err)
	}

	selector := &metav1.LabelSelector{MatchLabels: map[string]string{"listener-set": "enabled"}}
	gw := &v1.Gateway{
		ObjectMeta: metav1.ObjectMeta{Name: "shared-gw", Namespace: "infra"},
		Spec: v1.GatewaySpec{
			AllowedListeners: allowListenerSetsFromSelector(selector),
			Listeners:        []v1.Listener{{Protocol: v1.HTTPProtocolType, Port: 80}},
		},
		Status: gatewayStatus("10.0.0.1"),
	}
	_, err := gwClient.GatewayV1().Gateways(gw.Namespace).Create(ctx, gw, metav1.CreateOptions{})
	require.NoError(t, err)

	hostname := v1.Hostname("app.example.com")
	fromAll := v1.NamespacesFromAll
	parentNamespace := v1.Namespace("infra")
	ls := &v1.ListenerSet{
		ObjectMeta: metav1.ObjectMeta{Name: "edge-ls", Namespace: "edge"},
		Spec: v1.ListenerSetSpec{
			ParentRef: v1.ParentGatewayReference{Name: "shared-gw", Namespace: &parentNamespace},
			Listeners: []v1.ListenerEntry{{
				Name: "app", Hostname: &hostname, Port: 8080, Protocol: v1.HTTPProtocolType,
				AllowedRoutes: &v1.AllowedRoutes{Namespaces: &v1.RouteNamespaces{From: &fromAll}},
			}},
		},
		Status: listenerSetAcceptedStatus("app"),
	}
	_, err = gwClient.GatewayV1().ListenerSets(ls.Namespace).Create(ctx, ls, metav1.CreateOptions{})
	require.NoError(t, err)

	rt := &v1.HTTPRoute{
		ObjectMeta: metav1.ObjectMeta{Name: "rt", Namespace: "apps"},
		Spec: v1.HTTPRouteSpec{
			Hostnames:       []v1.Hostname{"app.example.com"},
			CommonRouteSpec: v1.CommonRouteSpec{ParentRefs: []v1.ParentReference{lsParentRef("edge", "edge-ls")}},
		},
		Status: v1.HTTPRouteStatus{RouteStatus: gwRouteStatus(lsParentRef("edge", "edge-ls"))},
	}
	_, err = gwClient.GatewayV1().HTTPRoutes(rt.Namespace).Create(ctx, rt, metav1.CreateOptions{})
	require.NoError(t, err)

	src, err := NewGatewayHTTPRouteSource(ctx, clients, &Config{
		GatewayNamespace:    "infra",
		GatewayListenerSets: true,
	})
	require.NoError(t, err)

	endpoints, err := src.Endpoints(ctx)
	require.NoError(t, err)
	testutils.ValidateEndpoints(t, endpoints, []*endpoint.Endpoint{
		newTestEndpoint("app.example.com", "10.0.0.1"),
	})
}

func TestGatewayHTTPRouteWithCrossNamespaceListenerSetGatewayLabelFilter(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(t.Context(), 10*time.Second)
	defer cancel()

	gwClient := gatewayfake.NewSimpleClientset()
	kubeClient := kubefake.NewClientset()
	clients := new(testutils.MockClientGenerator)
	clients.On("GatewayClient").Return(gwClient, nil)
	clients.On("KubeClient").Return(kubeClient, nil)

	for _, ns := range []*corev1.Namespace{
		{ObjectMeta: metav1.ObjectMeta{Name: "infra"}},
		{ObjectMeta: metav1.ObjectMeta{Name: "edge"}},
		{ObjectMeta: metav1.ObjectMeta{Name: "apps"}},
	} {
		_, err := kubeClient.CoreV1().Namespaces().Create(ctx, ns, metav1.CreateOptions{})
		require.NoError(t, err)
	}

	for _, gw := range []*v1.Gateway{
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "shared-gw-match",
				Namespace: "infra",
				Labels:    map[string]string{"foo": "bar"},
			},
			Spec:   v1.GatewaySpec{AllowedListeners: allowAllListenerSets(), Listeners: []v1.Listener{{Protocol: v1.HTTPProtocolType, Port: 80}}},
			Status: gatewayStatus("10.0.0.1"),
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "shared-gw-dont-match",
				Namespace: "infra",
				Labels:    map[string]string{"foo": "qux"},
			},
			Spec:   v1.GatewaySpec{AllowedListeners: allowAllListenerSets(), Listeners: []v1.Listener{{Protocol: v1.HTTPProtocolType, Port: 80}}},
			Status: gatewayStatus("10.0.0.2"),
		},
	} {
		_, err := gwClient.GatewayV1().Gateways(gw.Namespace).Create(ctx, gw, metav1.CreateOptions{})
		require.NoError(t, err)
	}

	hostname := v1.Hostname("app.example.com")
	fromAll := v1.NamespacesFromAll
	parentNamespace := v1.Namespace("infra")
	for _, ls := range []*v1.ListenerSet{
		{
			ObjectMeta: metav1.ObjectMeta{Name: "edge-ls-match", Namespace: "edge"},
			Spec: v1.ListenerSetSpec{
				ParentRef: v1.ParentGatewayReference{Name: "shared-gw-match", Namespace: &parentNamespace},
				Listeners: []v1.ListenerEntry{{
					Name: "app", Hostname: &hostname, Port: 8080, Protocol: v1.HTTPProtocolType,
					AllowedRoutes: &v1.AllowedRoutes{Namespaces: &v1.RouteNamespaces{From: &fromAll}},
				}},
			},
			Status: listenerSetAcceptedStatus("app"),
		},
		{
			ObjectMeta: metav1.ObjectMeta{Name: "edge-ls-dont-match", Namespace: "edge"},
			Spec: v1.ListenerSetSpec{
				ParentRef: v1.ParentGatewayReference{Name: "shared-gw-dont-match", Namespace: &parentNamespace},
				Listeners: []v1.ListenerEntry{{
					Name: "app", Hostname: &hostname, Port: 8080, Protocol: v1.HTTPProtocolType,
					AllowedRoutes: &v1.AllowedRoutes{Namespaces: &v1.RouteNamespaces{From: &fromAll}},
				}},
			},
			Status: listenerSetAcceptedStatus("app"),
		},
	} {
		_, err := gwClient.GatewayV1().ListenerSets(ls.Namespace).Create(ctx, ls, metav1.CreateOptions{})
		require.NoError(t, err)
	}

	rt := &v1.HTTPRoute{
		ObjectMeta: metav1.ObjectMeta{Name: "rt", Namespace: "apps"},
		Spec: v1.HTTPRouteSpec{
			Hostnames: []v1.Hostname{"app.example.com"},
			CommonRouteSpec: v1.CommonRouteSpec{
				ParentRefs: []v1.ParentReference{
					lsParentRef("edge", "edge-ls-match"),
					lsParentRef("edge", "edge-ls-dont-match"),
				},
			},
		},
		Status: v1.HTTPRouteStatus{
			RouteStatus: gwRouteStatus(
				lsParentRef("edge", "edge-ls-match"),
				lsParentRef("edge", "edge-ls-dont-match"),
			),
		},
	}
	_, err := gwClient.GatewayV1().HTTPRoutes(rt.Namespace).Create(ctx, rt, metav1.CreateOptions{})
	require.NoError(t, err)

	src, err := NewGatewayHTTPRouteSource(ctx, clients, &Config{
		GatewayLabelFilter:  "foo=bar",
		GatewayListenerSets: true,
	})
	require.NoError(t, err)

	endpoints, err := src.Endpoints(ctx)
	require.NoError(t, err)
	testutils.ValidateEndpoints(t, endpoints, []*endpoint.Endpoint{
		newTestEndpoint("app.example.com", "10.0.0.1"),
	})
}

func TestGatewayHTTPRouteWithListenerSetOwnTargetAnnotation(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(t.Context(), 10*time.Second)
	defer cancel()

	gwClient := gatewayfake.NewSimpleClientset()
	kubeClient := kubefake.NewClientset()
	clients := new(testutils.MockClientGenerator)
	clients.On("GatewayClient").Return(gwClient, nil)
	clients.On("KubeClient").Return(kubeClient, nil)

	ns := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "default"}}
	_, err := kubeClient.CoreV1().Namespaces().Create(ctx, ns, metav1.CreateOptions{})
	require.NoError(t, err)

	// Gateway without target annotation.
	gw := &v1.Gateway{
		ObjectMeta: metav1.ObjectMeta{Name: "gw", Namespace: "default"},
		Spec:       v1.GatewaySpec{AllowedListeners: allowAllListenerSets(), Listeners: []v1.Listener{{Protocol: v1.HTTPProtocolType, Port: 80}}},
		Status:     gatewayStatus("10.0.0.1"),
	}
	_, err = gwClient.GatewayV1().Gateways(gw.Namespace).Create(ctx, gw, metav1.CreateOptions{})
	require.NoError(t, err)

	// ListenerSet with its own target annotation.
	hostname := v1.Hostname("app.example.com")
	fromAll := v1.NamespacesFromAll
	ls := &v1.ListenerSet{
		ObjectMeta: metav1.ObjectMeta{
			Name: "ls", Namespace: "default",
			Annotations: map[string]string{"external-dns.alpha.kubernetes.io/target": "ls-override.example.com"},
		},
		Spec: v1.ListenerSetSpec{
			ParentRef: v1.ParentGatewayReference{Name: "gw"},
			Listeners: []v1.ListenerEntry{{
				Name: "app", Hostname: &hostname, Port: 8080, Protocol: v1.HTTPProtocolType,
				AllowedRoutes: &v1.AllowedRoutes{Namespaces: &v1.RouteNamespaces{From: &fromAll}},
			}},
		},
		Status: listenerSetAcceptedStatus("app"),
	}
	_, err = gwClient.GatewayV1().ListenerSets(ls.Namespace).Create(ctx, ls, metav1.CreateOptions{})
	require.NoError(t, err)

	rt := &v1.HTTPRoute{
		ObjectMeta: metav1.ObjectMeta{Name: "rt", Namespace: "default"},
		Spec: v1.HTTPRouteSpec{
			Hostnames:       []v1.Hostname{"app.example.com"},
			CommonRouteSpec: v1.CommonRouteSpec{ParentRefs: []v1.ParentReference{lsParentRef("default", "ls")}},
		},
		Status: v1.HTTPRouteStatus{RouteStatus: gwRouteStatus(lsParentRef("default", "ls"))},
	}
	_, err = gwClient.GatewayV1().HTTPRoutes(rt.Namespace).Create(ctx, rt, metav1.CreateOptions{})
	require.NoError(t, err)

	src, err := NewGatewayHTTPRouteSource(ctx, clients, &Config{GatewayListenerSets: true})
	require.NoError(t, err)

	endpoints, err := src.Endpoints(ctx)
	require.NoError(t, err)
	testutils.ValidateEndpoints(t, endpoints, []*endpoint.Endpoint{
		newTestEndpointWithTTL("app.example.com", endpoint.RecordTypeCNAME, 0, "ls-override.example.com"),
	})
}

func TestGatewayHTTPRouteWithListenerSetTargetAnnotationPrecedence(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(t.Context(), 10*time.Second)
	defer cancel()

	gwClient := gatewayfake.NewSimpleClientset()
	kubeClient := kubefake.NewClientset()
	clients := new(testutils.MockClientGenerator)
	clients.On("GatewayClient").Return(gwClient, nil)
	clients.On("KubeClient").Return(kubeClient, nil)

	ns := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "default"}}
	_, err := kubeClient.CoreV1().Namespaces().Create(ctx, ns, metav1.CreateOptions{})
	require.NoError(t, err)

	// Gateway with its own target annotation.
	gw := &v1.Gateway{
		ObjectMeta: metav1.ObjectMeta{
			Name: "gw", Namespace: "default",
			Annotations: map[string]string{"external-dns.alpha.kubernetes.io/target": "gw.example.com"},
		},
		Spec:   v1.GatewaySpec{AllowedListeners: allowAllListenerSets(), Listeners: []v1.Listener{{Protocol: v1.HTTPProtocolType, Port: 80}}},
		Status: gatewayStatus("10.0.0.1"),
	}
	_, err = gwClient.GatewayV1().Gateways(gw.Namespace).Create(ctx, gw, metav1.CreateOptions{})
	require.NoError(t, err)

	// ListenerSet with its own target annotation — should take precedence.
	hostname := v1.Hostname("app.example.com")
	fromAll := v1.NamespacesFromAll
	ls := &v1.ListenerSet{
		ObjectMeta: metav1.ObjectMeta{
			Name: "ls", Namespace: "default",
			Annotations: map[string]string{"external-dns.alpha.kubernetes.io/target": "ls.example.com"},
		},
		Spec: v1.ListenerSetSpec{
			ParentRef: v1.ParentGatewayReference{Name: "gw"},
			Listeners: []v1.ListenerEntry{{
				Name: "app", Hostname: &hostname, Port: 8080, Protocol: v1.HTTPProtocolType,
				AllowedRoutes: &v1.AllowedRoutes{Namespaces: &v1.RouteNamespaces{From: &fromAll}},
			}},
		},
		Status: listenerSetAcceptedStatus("app"),
	}
	_, err = gwClient.GatewayV1().ListenerSets(ls.Namespace).Create(ctx, ls, metav1.CreateOptions{})
	require.NoError(t, err)

	rt := &v1.HTTPRoute{
		ObjectMeta: metav1.ObjectMeta{Name: "rt", Namespace: "default"},
		Spec: v1.HTTPRouteSpec{
			Hostnames:       []v1.Hostname{"app.example.com"},
			CommonRouteSpec: v1.CommonRouteSpec{ParentRefs: []v1.ParentReference{lsParentRef("default", "ls")}},
		},
		Status: v1.HTTPRouteStatus{RouteStatus: gwRouteStatus(lsParentRef("default", "ls"))},
	}
	_, err = gwClient.GatewayV1().HTTPRoutes(rt.Namespace).Create(ctx, rt, metav1.CreateOptions{})
	require.NoError(t, err)

	src, err := NewGatewayHTTPRouteSource(ctx, clients, &Config{GatewayListenerSets: true})
	require.NoError(t, err)

	endpoints, err := src.Endpoints(ctx)
	require.NoError(t, err)
	// ListenerSet annotation should win over Gateway annotation.
	testutils.ValidateEndpoints(t, endpoints, []*endpoint.Endpoint{
		newTestEndpointWithTTL("app.example.com", endpoint.RecordTypeCNAME, 0, "ls.example.com"),
	})
}

func TestGatewayHTTPRouteWithListenerSetGatewayNotFound(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(t.Context(), 10*time.Second)
	defer cancel()

	gwClient := gatewayfake.NewSimpleClientset()
	kubeClient := kubefake.NewClientset()
	clients := new(testutils.MockClientGenerator)
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
		Status: listenerSetAcceptedStatus("app"),
	}
	_, err = gwClient.GatewayV1().ListenerSets(ls.Namespace).Create(ctx, ls, metav1.CreateOptions{})
	require.NoError(t, err)

	rt := &v1.HTTPRoute{
		ObjectMeta: metav1.ObjectMeta{Name: "rt", Namespace: "default"},
		Spec: v1.HTTPRouteSpec{
			Hostnames:       []v1.Hostname{"app.example.com"},
			CommonRouteSpec: v1.CommonRouteSpec{ParentRefs: []v1.ParentReference{lsParentRef("default", "ls")}},
		},
		Status: v1.HTTPRouteStatus{RouteStatus: gwRouteStatus(lsParentRef("default", "ls"))},
	}
	_, err = gwClient.GatewayV1().HTTPRoutes(rt.Namespace).Create(ctx, rt, metav1.CreateOptions{})
	require.NoError(t, err)

	src, err := NewGatewayHTTPRouteSource(ctx, clients, &Config{GatewayListenerSets: true})
	require.NoError(t, err)

	endpoints, err := src.Endpoints(ctx)
	require.NoError(t, err)
	require.Empty(t, endpoints, "expected no endpoints when ListenerSet's parent Gateway doesn't exist")
}
