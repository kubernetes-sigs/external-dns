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
	"maps"
	"testing"
	"time"

	"sigs.k8s.io/external-dns/internal/testutils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubefake "k8s.io/client-go/kubernetes/fake"
	v1 "sigs.k8s.io/gateway-api/apis/v1"
	"sigs.k8s.io/gateway-api/apis/v1alpha2"
	gatewayfake "sigs.k8s.io/gateway-api/pkg/client/clientset/versioned/fake"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/source/annotations"
	templatetest "sigs.k8s.io/external-dns/source/template/testutil"
)

func TestGatewayTCPRouteSourceEndpoints(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(t.Context(), 10*time.Second)
	defer cancel()

	gwClient := gatewayfake.NewSimpleClientset()
	kubeClient := kubefake.NewClientset()
	clients := new(testutils.MockClientGenerator)
	clients.On("GatewayClient").Return(gwClient, nil)
	clients.On("KubeClient").Return(kubeClient, nil)

	ns := &corev1.Namespace{
		Name: "default",
	}
	_, err := kubeClient.CoreV1().Namespaces().Create(ctx, ns, metav1.CreateOptions{})
	require.NoError(t, err, "failed to create Namespace")

	ips := []string{"10.64.0.1", "10.64.0.2"}
	gw := &v1.Gateway{
		Name:      "internal",
		Namespace: "default",
		Spec: v1.GatewaySpec{
			Listeners: []v1.Listener{{
				Protocol: v1.TCPProtocolType,
			}},
		},
		Status: gatewayStatus(ips...),
	}
	_, err = gwClient.GatewayV1().Gateways(gw.Namespace).Create(ctx, gw, metav1.CreateOptions{})
	require.NoError(t, err, "failed to create Gateway")

	rt := &v1alpha2.TCPRoute{
		Name:      "api",
		Namespace: "default",
		Annotations: map[string]string{
			annotations.HostnameKey: "api-annotation.foobar.internal",
		},
		Spec: v1alpha2.TCPRouteSpec{
			CommonRouteSpec: v1.CommonRouteSpec{
				ParentRefs: []v1.ParentReference{
					gwParentRef("default", "internal"),
				},
			},
		},
		Status: v1alpha2.TCPRouteStatus{
			RouteStatus: gwRouteStatus(gwParentRef("default", "internal")),
		},
	}
	_, err = gwClient.GatewayV1alpha2().TCPRoutes(rt.Namespace).Create(ctx, rt, metav1.CreateOptions{})
	require.NoError(t, err, "failed to create TCPRoute")

	src, err := NewGatewayTCPRouteSource(ctx, clients, &Config{
		TemplateEngine: templatetest.MustEngine(t, "{{.Name}}-template.foobar.internal", "", "", true),
	})
	require.NoError(t, err, "failed to create Gateway TCPRoute Source")

	endpoints, err := src.Endpoints(ctx)
	require.NoError(t, err, "failed to get Endpoints")
	testutils.ValidateEndpoints(t, endpoints, []*endpoint.Endpoint{
		newTestEndpoint("api-annotation.foobar.internal", ips...),
		newTestEndpoint("api-template.foobar.internal", ips...),
	})
}

func TestGatewayTCPRouteSource_InformerTransform(t *testing.T) {
	t.Parallel()

	gwClient := gatewayfake.NewSimpleClientset()
	kubeClient := kubefake.NewClientset()

	rt := &v1alpha2.TCPRoute{ObjectMeta: informerTransformObjectMeta()}
	require.Contains(t, rt.GetAnnotations(), corev1.LastAppliedConfigAnnotation)
	require.NotEmpty(t, rt.GetManagedFields())

	_, err := gwClient.GatewayV1alpha2().TCPRoutes(rt.GetNamespace()).Create(t.Context(), rt, metav1.CreateOptions{})
	require.NoError(t, err)

	clients := new(testutils.MockClientGenerator)
	clients.On("GatewayClient").Return(gwClient, nil)
	clients.On("KubeClient").Return(kubeClient, nil)

	source, err := NewGatewayTCPRouteSource(t.Context(), clients, &Config{})
	require.NoError(t, err)
	require.IsType(t, &gatewayRouteSource{}, source)

	testInformerTransformHelper(t,
		source.(*gatewayRouteSource).rtInformer.Informer(),
		rt,
		withRemovedLastAppliedConfigAnnotation(),
		withRemovedManagedFields(),
	)
}

func TestGatewayTCPRouteIndexer(t *testing.T) {
	t.Parallel()

	fromAll := v1.NamespacesFromAll

	makeRoute := func(namespace, name string, ann, lbls map[string]string) *v1alpha2.TCPRoute {
		allAnn := map[string]string{annotations.HostnameKey: name + ".example.com"}
		maps.Copy(allAnn, ann)
		return &v1alpha2.TCPRoute{
			Namespace:   namespace,
			Name:        name,
			Annotations: allAnn,
			Labels:      lbls,
			Spec: v1alpha2.TCPRouteSpec{
				CommonRouteSpec: v1.CommonRouteSpec{
					ParentRefs: []v1.ParentReference{gwParentRef("default", "gw")},
				},
			},
			Status: v1alpha2.TCPRouteStatus{
				RouteStatus: gwRouteStatus(gwParentRef("default", "gw")),
			},
		}
	}

	for _, tc := range []struct {
		name             string
		annotationFilter string
		labelFilter      string
		routes           []*v1alpha2.TCPRoute
		wantCount        int
	}{
		{
			name: "no filters — all namespaces included",
			routes: []*v1alpha2.TCPRoute{
				makeRoute("default", "r1", nil, nil),
				makeRoute("staging", "r2", nil, nil),
				makeRoute("production", "r3", nil, nil),
			},
			wantCount: 3,
		},
		{
			name:             "annotation filter matches",
			annotationFilter: "external-dns.kubernetes.io/managed=true",
			routes: []*v1alpha2.TCPRoute{
				makeRoute("default", "r1", map[string]string{"external-dns.kubernetes.io/managed": "true"}, nil),
				makeRoute("default", "r2", nil, nil),
			},
			wantCount: 1,
		},
		{
			name:        "label filter matches",
			labelFilter: "tier=external",
			routes: []*v1alpha2.TCPRoute{
				makeRoute("default", "r1", nil, map[string]string{"tier": "external"}),
				makeRoute("default", "r2", nil, map[string]string{"tier": "internal"}),
			},
			wantCount: 1,
		},
		{
			name:             "annotation and label filter combined",
			annotationFilter: "external-dns.kubernetes.io/managed=true",
			labelFilter:      "tier=external",
			routes: []*v1alpha2.TCPRoute{
				makeRoute("default", "r1",
					map[string]string{"external-dns.kubernetes.io/managed": "true"},
					map[string]string{"tier": "external"}),
				makeRoute("default", "r2",
					map[string]string{"external-dns.kubernetes.io/managed": "true"},
					map[string]string{"tier": "internal"}),
			},
			wantCount: 1,
		},
		{
			name:             "no-match annotation filter",
			annotationFilter: "external-dns.kubernetes.io/managed=true",
			routes: []*v1alpha2.TCPRoute{
				makeRoute("default", "r1", nil, nil),
				makeRoute("default", "r2", nil, nil),
			},
			wantCount: 0,
		},
		{
			name: "controller mismatch is excluded",
			routes: []*v1alpha2.TCPRoute{
				makeRoute("default", "r1",
					map[string]string{annotations.ControllerKey: "other-controller"},
					nil),
			},
			wantCount: 0,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctx, cancel := context.WithTimeout(t.Context(), 10*time.Second)
			defer cancel()

			gwClient := gatewayfake.NewSimpleClientset()
			kubeClient := kubefake.NewClientset()

			gw := &v1.Gateway{
				Namespace: "default", Name: "gw",
				Spec: v1.GatewaySpec{
					Listeners: []v1.Listener{{
						Protocol: v1.TCPProtocolType,
						AllowedRoutes: &v1.AllowedRoutes{
							Namespaces: &v1.RouteNamespaces{From: &fromAll},
						},
					}},
				},
				Status: gatewayStatus("1.2.3.4"),
			}
			_, err := gwClient.GatewayV1().Gateways("default").Create(ctx, gw, metav1.CreateOptions{})
			require.NoError(t, err)

			for _, rt := range tc.routes {
				_, err := gwClient.GatewayV1alpha2().TCPRoutes(rt.Namespace).Create(ctx, rt, metav1.CreateOptions{})
				require.NoError(t, err)
			}

			clients := new(testutils.MockClientGenerator)
			clients.On("GatewayClient").Return(gwClient, nil)
			clients.On("KubeClient").Return(kubeClient, nil)

			src, err := NewGatewayTCPRouteSource(ctx, clients, &Config{
				AnnotationFilter: parseLabelSelectorOrEverything(t, tc.annotationFilter),
				LabelFilter:      parseLabelSelectorOrEverything(t, tc.labelFilter),
			})
			require.NoError(t, err)

			endpoints, err := src.Endpoints(ctx)
			require.NoError(t, err)
			assert.Len(t, endpoints, tc.wantCount)
		})
	}
}
