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
	"testing"

	"github.com/stretchr/testify/require"
	istionetworking "istio.io/api/networking/v1beta1"
	networkingv1 "istio.io/client-go/pkg/apis/networking/v1"
	istiofake "istio.io/client-go/pkg/clientset/versioned/fake"
	networkv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/dynamic"
	fakedynamic "k8s.io/client-go/dynamic/fake"
	"k8s.io/client-go/kubernetes/fake"
	clienttesting "k8s.io/client-go/testing"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	"sigs.k8s.io/external-dns/source/annotations"
	"sigs.k8s.io/external-dns/source/targetfrom"
)

const (
	sharedLBRef     = "services/lb-ns/shared-lb#{{ range .status.loadBalancer.ingress }}{{ .ip }},{{ end }}"
	missingLBRef    = "services/lb-ns/missing#{{ range .status.loadBalancer.ingress }}{{ .ip }},{{ end }}"
	sharedLBAddress = "203.0.113.10"
)

// newTargetFromConfig returns a Config whose target-from resolver reads the
// shared-lb Service in lb-ns from a fake dynamic client.
func newTargetFromConfig(t *testing.T) *Config {
	t.Helper()
	gvr := schema.GroupVersionResource{Version: "v1", Resource: "services"}
	lb := &unstructured.Unstructured{Object: map[string]any{
		"apiVersion": "v1",
		"kind":       "Service",
		"metadata":   map[string]any{"namespace": "lb-ns", "name": "shared-lb"},
		"status": map[string]any{"loadBalancer": map[string]any{
			"ingress": []any{map[string]any{"ip": sharedLBAddress}},
		}},
	}}
	dyn := fakedynamic.NewSimpleDynamicClientWithCustomListKinds(runtime.NewScheme(),
		map[schema.GroupVersionResource]string{gvr: "ServiceList"}, lb)
	disc := &fakediscovery.FakeDiscovery{Fake: &clienttesting.Fake{Resources: []*metav1.APIResourceList{
		{GroupVersion: "v1", APIResources: []metav1.APIResource{{Name: "services", Kind: "Service", Namespaced: true}}},
	}}}
	return &Config{
		LabelFilter: labels.Everything(),
		targetFrom: targetfrom.NewResolver(t.Context(), func() (dynamic.Interface, discovery.DiscoveryInterface, error) {
			return dyn, disc, nil
		}),
	}
}

// newIstioClient creates objs through the typed client: the fake tracker
// guesses resources from kinds, and turns Gateway into "gatewaies".
func newIstioClient(t *testing.T, gw *networkingv1.Gateway, vs ...*networkingv1.VirtualService) *istiofake.Clientset {
	t.Helper()
	client := istiofake.NewSimpleClientset()
	_, err := client.NetworkingV1().Gateways(gw.Namespace).Create(t.Context(), gw, metav1.CreateOptions{})
	require.NoError(t, err)
	for _, v := range vs {
		_, err := client.NetworkingV1().VirtualServices(v.Namespace).Create(t.Context(), v, metav1.CreateOptions{})
		require.NoError(t, err)
	}
	return client
}

func TestIngressTargetFrom(t *testing.T) {
	for _, tt := range []struct {
		name        string
		annotations map[string]string
		expected    []*endpoint.Endpoint
	}{
		{
			name:        "resolved reference replaces the load balancer status",
			annotations: map[string]string{annotations.TargetFromKey: sharedLBRef},
			expected:    []*endpoint.Endpoint{endpoint.NewEndpoint("app.example.org", endpoint.RecordTypeA, sharedLBAddress).WithLabel(endpoint.ResourceLabelKey, "ingress/apps/app")},
		},
		{
			name: "target annotation wins",
			annotations: map[string]string{
				annotations.TargetKey:     "198.51.100.1",
				annotations.TargetFromKey: sharedLBRef,
			},
			expected: []*endpoint.Endpoint{endpoint.NewEndpoint("app.example.org", endpoint.RecordTypeA, "198.51.100.1").WithLabel(endpoint.ResourceLabelKey, "ingress/apps/app")},
		},
		{
			name:        "unresolved reference does not fall back to the load balancer status",
			annotations: map[string]string{annotations.TargetFromKey: missingLBRef},
			expected:    []*endpoint.Endpoint{},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ing := &networkv1.Ingress{
				Namespace: "apps", Name: "app", Annotations: tt.annotations,
				Spec: networkv1.IngressSpec{Rules: []networkv1.IngressRule{{Host: "app.example.org"}}},
				Status: networkv1.IngressStatus{LoadBalancer: networkv1.IngressLoadBalancerStatus{
					Ingress: []networkv1.IngressLoadBalancerIngress{{IP: "192.0.2.1"}},
				}},
			}
			src, err := NewIngressSource(t.Context(), fake.NewClientset(ing), newTargetFromConfig(t))
			require.NoError(t, err)

			endpoints, err := src.Endpoints(t.Context())
			require.NoError(t, err)
			testutils.ValidateEndpoints(t, endpoints, tt.expected)
		})
	}
}

func TestIstioTargetFrom(t *testing.T) {
	// Matches the gateway selector, so native discovery would publish 192.0.2.1.
	svc := (fakeIngressGatewayService{
		ips:       []string{"192.0.2.1"},
		namespace: "apps",
		name:      "istio-ingressgateway",
		selector:  map[string]string{"istio": "ingressgateway"},
	}).Service()

	for _, tt := range []struct {
		name     string
		ref      string
		resolved bool
	}{
		{name: "resolved reference replaces service discovery", ref: sharedLBRef, resolved: true},
		{name: "unresolved reference does not fall back to service discovery", ref: missingLBRef},
	} {
		expected := func(resource string) []*endpoint.Endpoint {
			if !tt.resolved {
				return []*endpoint.Endpoint{}
			}
			return []*endpoint.Endpoint{
				endpoint.NewEndpoint("gw.example.org", endpoint.RecordTypeA, sharedLBAddress).WithLabel(endpoint.ResourceLabelKey, resource),
			}
		}

		t.Run("gateway/"+tt.name, func(t *testing.T) {
			gw := &networkingv1.Gateway{
				Namespace: "apps", Name: "gw", Annotations: map[string]string{annotations.TargetFromKey: tt.ref},
				Spec: istionetworking.Gateway{
					Servers:  []*istionetworking.Server{{Hosts: []string{"gw.example.org"}}},
					Selector: map[string]string{"istio": "ingressgateway"},
				},
			}
			src, err := NewIstioGatewaySource(t.Context(), fake.NewClientset(svc), newIstioClient(t, gw), newTargetFromConfig(t))
			require.NoError(t, err)

			endpoints, err := src.Endpoints(t.Context())
			require.NoError(t, err)
			testutils.ValidateEndpoints(t, endpoints, expected("gateway/apps/gw"))
		})

		t.Run("virtualservice/"+tt.name, func(t *testing.T) {
			gw := &networkingv1.Gateway{
				Namespace: "apps", Name: "gw",
				Spec: istionetworking.Gateway{
					Servers:  []*istionetworking.Server{{Hosts: []string{"*"}}},
					Selector: map[string]string{"istio": "ingressgateway"},
				},
			}
			vs := &networkingv1.VirtualService{
				Namespace: "apps", Name: "vs", Annotations: map[string]string{annotations.TargetFromKey: tt.ref},
				Spec: istionetworking.VirtualService{
					Hosts:    []string{"gw.example.org"},
					Gateways: []string{"gw"},
				},
			}
			src, err := NewIstioVirtualServiceSource(t.Context(), fake.NewClientset(svc), newIstioClient(t, gw, vs), newTargetFromConfig(t))
			require.NoError(t, err)

			endpoints, err := src.Endpoints(t.Context())
			require.NoError(t, err)
			testutils.ValidateEndpoints(t, endpoints, expected("virtualservice/apps/vs"))
		})
	}
}

// A gateway's target-from also applies to the virtual services bound to it.
func TestIstioVirtualServiceGatewayTargetFrom(t *testing.T) {
	gw := &networkingv1.Gateway{
		Namespace: "apps", Name: "gw", Annotations: map[string]string{annotations.TargetFromKey: sharedLBRef},
		Spec: istionetworking.Gateway{Servers: []*istionetworking.Server{{Hosts: []string{"*"}}}},
	}
	vs := &networkingv1.VirtualService{
		Namespace: "apps", Name: "vs",
		Spec: istionetworking.VirtualService{Hosts: []string{"vs.example.org"}, Gateways: []string{"gw"}},
	}
	src, err := NewIstioVirtualServiceSource(t.Context(), fake.NewClientset(), newIstioClient(t, gw, vs), newTargetFromConfig(t))
	require.NoError(t, err)

	endpoints, err := src.Endpoints(t.Context())
	require.NoError(t, err)
	testutils.ValidateEndpoints(t, endpoints, []*endpoint.Endpoint{
		endpoint.NewEndpoint("vs.example.org", endpoint.RecordTypeA, sharedLBAddress).WithLabel(endpoint.ResourceLabelKey, "virtualservice/apps/vs"),
	})
}
