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
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	fakeKube "k8s.io/client-go/kubernetes/fake"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	templatetest "sigs.k8s.io/external-dns/source/template/testutil"
)

func TestGlooProxyFQDNTemplate(t *testing.T) {
	t.Parallel()

	const ns = defaultGlooNamespace

	// A proxy with one virtual host whose domain is "static.example.com".
	// The backing LoadBalancer service has no IP (simulating IPAM not yet populated).
	makeProxy := func(name string) proxy {
		return proxy{
			TypeMeta: metav1.TypeMeta{
				APIVersion: proxyGVR.GroupVersion().String(),
				Kind:       "Proxy",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      name,
				Namespace: ns,
			},
			Spec: proxySpec{
				Listeners: []proxySpecListener{
					{
						HTTPListener: proxySpecHTTPListener{
							VirtualHosts: []proxyVirtualHost{
								{Domains: []string{"static.example.com"}},
							},
						},
					},
				},
			},
		}
	}

	// Service with no LoadBalancer IP — simulates IPAM not yet populated.
	makeEmptySvc := func(name string) corev1.Service {
		return corev1.Service{
			ObjectMeta: metav1.ObjectMeta{Name: name, Namespace: ns},
			Spec:       corev1.ServiceSpec{Type: corev1.ServiceTypeLoadBalancer},
		}
	}

	// Service with a real LoadBalancer IP.
	makeIPSvc := func(name, ip string) corev1.Service {
		return corev1.Service{
			ObjectMeta: metav1.ObjectMeta{Name: name, Namespace: ns},
			Spec:       corev1.ServiceSpec{Type: corev1.ServiceTypeLoadBalancer},
			Status: corev1.ServiceStatus{
				LoadBalancer: corev1.LoadBalancerStatus{
					Ingress: []corev1.LoadBalancerIngress{{IP: ip}},
				},
			},
		}
	}

	tests := []struct {
		title              string
		proxyName          string
		customProxy        *proxy // overrides makeProxy when set
		svc                corev1.Service
		fqdnTemplate       string
		targetTemplate     string
		fqdnTargetTemplate string
		combine            bool
		expected           []*endpoint.Endpoint
	}{
		{
			title:          "fqdn-template + target-template generate endpoint when no service IP",
			proxyName:      "my-proxy",
			svc:            makeEmptySvc("my-proxy"),
			fqdnTemplate:   "{{.Name}}.example.com",
			targetTemplate: "lb.example.com",
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("my-proxy.example.com", endpoint.RecordTypeCNAME, "lb.example.com"),
			},
		},
		{
			title:              "fqdn-target-template generates endpoint when no service IP",
			proxyName:          "my-proxy",
			svc:                makeEmptySvc("my-proxy"),
			fqdnTargetTemplate: "{{.Name}}.example.com:1.2.3.4",
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("my-proxy.example.com", endpoint.RecordTypeA, "1.2.3.4"),
			},
		},
		{
			title:          "fqdn-template with combine adds template endpoint alongside virtual-host endpoint",
			proxyName:      "my-proxy",
			svc:            makeIPSvc("my-proxy", "10.0.0.1"),
			fqdnTemplate:   "{{.Name}}.example.com",
			targetTemplate: "lb.example.com",
			combine:        true,
			expected: []*endpoint.Endpoint{
				{
					DNSName:          "static.example.com",
					Targets:          endpoint.Targets{"10.0.0.1"},
					RecordType:       endpoint.RecordTypeA,
					Labels:           endpoint.Labels{},
					ProviderSpecific: endpoint.ProviderSpecific{},
				},
				{
					DNSName:          "my-proxy.example.com",
					Targets:          endpoint.Targets{"lb.example.com"},
					RecordType:       endpoint.RecordTypeCNAME,
					Labels:           endpoint.Labels{},
					ProviderSpecific: endpoint.ProviderSpecific{},
				},
			},
		},
		{
			title:              "fqdn-target-template with combine adds template endpoint alongside virtual-host endpoint",
			proxyName:          "my-proxy",
			svc:                makeIPSvc("my-proxy", "10.0.0.1"),
			fqdnTargetTemplate: "{{.Name}}.example.com:lb.example.com",
			combine:            true,
			expected: []*endpoint.Endpoint{
				{
					DNSName:          "static.example.com",
					Targets:          endpoint.Targets{"10.0.0.1"},
					RecordType:       endpoint.RecordTypeA,
					Labels:           endpoint.Labels{},
					ProviderSpecific: endpoint.ProviderSpecific{},
				},
				{
					DNSName:          "my-proxy.example.com",
					Targets:          endpoint.Targets{"lb.example.com"},
					RecordType:       endpoint.RecordTypeCNAME,
					Labels:           endpoint.Labels{},
					ProviderSpecific: endpoint.ProviderSpecific{},
				},
			},
		},
		{
			title:          "fqdn-template without combine is ignored when virtual-host endpoints exist",
			proxyName:      "my-proxy",
			svc:            makeIPSvc("my-proxy", "10.0.0.1"),
			fqdnTemplate:   "{{.Name}}.example.com",
			targetTemplate: "lb.example.com",
			combine:        false,
			expected: []*endpoint.Endpoint{
				{
					DNSName:          "static.example.com",
					Targets:          endpoint.Targets{"10.0.0.1"},
					RecordType:       endpoint.RecordTypeA,
					Labels:           endpoint.Labels{},
					ProviderSpecific: endpoint.ProviderSpecific{},
				},
			},
		},
		{
			title:              "fqdn-target-template without combine is ignored when virtual-host endpoints exist",
			proxyName:          "my-proxy",
			svc:                makeIPSvc("my-proxy", "10.0.0.1"),
			fqdnTargetTemplate: "{{.Name}}.example.com:lb.example.com",
			combine:            false,
			expected: []*endpoint.Endpoint{
				{
					DNSName:          "static.example.com",
					Targets:          endpoint.Targets{"10.0.0.1"},
					RecordType:       endpoint.RecordTypeA,
					Labels:           endpoint.Labels{},
					ProviderSpecific: endpoint.ProviderSpecific{},
				},
			},
		},
		{
			title:          "fqdn-template can reference .Kind",
			proxyName:      "my-proxy",
			svc:            makeEmptySvc("my-proxy"),
			fqdnTemplate:   "{{.Kind | toLower}}.{{.Name}}.example.com",
			targetTemplate: "lb.example.com",
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("proxy.my-proxy.example.com", endpoint.RecordTypeCNAME, "lb.example.com"),
			},
		},
		{
			title:              "fqdn-target-template can reference .APIVersion",
			proxyName:          "my-proxy",
			svc:                makeEmptySvc("my-proxy"),
			fqdnTargetTemplate: `{{.Name}}.{{replace "/" "." .APIVersion}}.example.com:1.2.3.4`,
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("my-proxy.gloo.solo.io.v1.example.com", endpoint.RecordTypeA, "1.2.3.4"),
			},
		},
		{
			title: "fqdn-template with combine alongside multiple virtual-host domains",
			customProxy: &proxy{
				TypeMeta: metav1.TypeMeta{
					APIVersion: proxyGVR.GroupVersion().String(),
					Kind:       "Proxy",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "multi-host-proxy",
					Namespace: ns,
				},
				Spec: proxySpec{
					Listeners: []proxySpecListener{
						{
							HTTPListener: proxySpecHTTPListener{
								VirtualHosts: []proxyVirtualHost{
									{Domains: []string{"host1.example.com", "host2.example.com"}},
									{Domains: []string{"host3.example.com"}},
								},
							},
						},
					},
				},
			},
			svc:            makeIPSvc("multi-host-proxy", "10.0.0.2"),
			fqdnTemplate:   "{{.Name}}.dns.example.com",
			targetTemplate: "lb.example.com",
			combine:        true,
			expected: []*endpoint.Endpoint{
				{
					DNSName:          "host1.example.com",
					Targets:          endpoint.Targets{"10.0.0.2"},
					RecordType:       endpoint.RecordTypeA,
					Labels:           endpoint.Labels{},
					ProviderSpecific: endpoint.ProviderSpecific{},
				},
				{
					DNSName:          "host2.example.com",
					Targets:          endpoint.Targets{"10.0.0.2"},
					RecordType:       endpoint.RecordTypeA,
					Labels:           endpoint.Labels{},
					ProviderSpecific: endpoint.ProviderSpecific{},
				},
				{
					DNSName:          "host3.example.com",
					Targets:          endpoint.Targets{"10.0.0.2"},
					RecordType:       endpoint.RecordTypeA,
					Labels:           endpoint.Labels{},
					ProviderSpecific: endpoint.ProviderSpecific{},
				},
				{
					DNSName:          "multi-host-proxy.dns.example.com",
					Targets:          endpoint.Targets{"lb.example.com"},
					RecordType:       endpoint.RecordTypeCNAME,
					Labels:           endpoint.Labels{},
					ProviderSpecific: endpoint.ProviderSpecific{},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			t.Parallel()

			p := makeProxy(tt.proxyName)
			if tt.customProxy != nil {
				p = *tt.customProxy
			}
			proxyJSON, err := json.Marshal(p)
			require.NoError(t, err)

			proxyObj := unstructured.Unstructured{}
			require.NoError(t, proxyObj.UnmarshalJSON(proxyJSON))

			fakeDynamicClient := newGlooDynamicClient(&proxyObj)
			fakeKubernetesClient := fakeKube.NewSimpleClientset(&tt.svc)

			src, err := NewGlooSource(t.Context(), fakeDynamicClient, fakeKubernetesClient, &Config{
				GlooNamespaces: []string{ns},
				TemplateEngine: templatetest.MustEngine(t, tt.fqdnTemplate, tt.targetTemplate, tt.fqdnTargetTemplate, tt.combine),
			})
			require.NoError(t, err)

			// Wait for the proxy to appear in the lister cache.
			count := &unstructured.UnstructuredList{}
			for len(count.Items) < 1 {
				count, _ = fakeDynamicClient.Resource(proxyGVR).Namespace(ns).List(t.Context(), metav1.ListOptions{})
			}

			endpoints, err := src.Endpoints(t.Context())
			assert.NoError(t, err)
			testutils.ValidateEndpoints(t, endpoints, tt.expected)
		})
	}
}
