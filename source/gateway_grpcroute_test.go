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
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubefake "k8s.io/client-go/kubernetes/fake"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/gateway-api/apis/v1alpha2"
	"sigs.k8s.io/gateway-api/apis/v1beta1"
	gatewayfake "sigs.k8s.io/gateway-api/pkg/client/clientset/versioned/fake"
)

func TestGatewayGRPCRouteSourceEndpoints(t *testing.T) {
	t.Parallel()

	gwClient := gatewayfake.NewSimpleClientset()
	kubeClient := kubefake.NewSimpleClientset()
	clients := new(MockClientGenerator)
	clients.On("GatewayClient").Return(gwClient, nil)
	clients.On("KubeClient").Return(kubeClient, nil)

	ctx := context.Background()
	ns := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "default",
		},
	}
	_, err := kubeClient.CoreV1().Namespaces().Create(ctx, ns, metav1.CreateOptions{})
	require.NoError(t, err, "failed to create Namespace")

	ips := []string{"10.64.0.1", "10.64.0.2"}
	gw := &v1beta1.Gateway{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "internal",
			Namespace: "default",
		},
		Spec: v1beta1.GatewaySpec{
			Listeners: []v1beta1.Listener{{
				Protocol: v1beta1.HTTPSProtocolType,
			}},
		},
		Status: gatewayStatus(ips...),
	}
	_, err = gwClient.GatewayV1beta1().Gateways(gw.Namespace).Create(ctx, gw, metav1.CreateOptions{})
	require.NoError(t, err, "failed to create Gateway")

	rt := &v1alpha2.GRPCRoute{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "api",
			Namespace: "default",
			Annotations: map[string]string{
				hostnameAnnotationKey: "api-annotation.foobar.internal",
			},
		},
		Spec: v1alpha2.GRPCRouteSpec{
			Hostnames: []v1alpha2.Hostname{"api-hostnames.foobar.internal"},
		},
		Status: v1alpha2.GRPCRouteStatus{
			RouteStatus: v1a2RouteStatus(v1a2ParentRef("default", "internal")),
		},
	}
	_, err = gwClient.GatewayV1alpha2().GRPCRoutes(rt.Namespace).Create(ctx, rt, metav1.CreateOptions{})
	require.NoError(t, err, "failed to create GRPCRoute")

	src, err := NewGatewayGRPCRouteSource(clients, &Config{
		FQDNTemplate:             "{{.Name}}-template.foobar.internal",
		CombineFQDNAndAnnotation: true,
	})
	require.NoError(t, err, "failed to create Gateway GRPCRoute Source")

	endpoints, err := src.Endpoints(ctx)
	require.NoError(t, err, "failed to get Endpoints")
	validateEndpoints(t, endpoints, []*endpoint.Endpoint{
		newTestEndpoint("api-annotation.foobar.internal", "A", ips...),
		newTestEndpoint("api-hostnames.foobar.internal", "A", ips...),
		newTestEndpoint("api-template.foobar.internal", "A", ips...),
	})
}
