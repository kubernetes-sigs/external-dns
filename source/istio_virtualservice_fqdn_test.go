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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	istionetworking "istio.io/api/networking/v1beta1"
	networkingv1beta1 "istio.io/client-go/pkg/apis/networking/v1beta1"
	istiofake "istio.io/client-go/pkg/clientset/versioned/fake"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"

	"sigs.k8s.io/external-dns/source/annotations"

	"sigs.k8s.io/external-dns/endpoint"
)

func TestIstioVirtualServiceSourceNewSourceWithFqdn(t *testing.T) {
	for _, tt := range []struct {
		title            string
		annotationFilter string
		fqdnTemplate     string
		expectError      bool
	}{
		{
			title:        "invalid template",
			expectError:  true,
			fqdnTemplate: "{{.Name",
		},
		{
			title:       "valid empty template",
			expectError: false,
		},
		{
			title:        "valid template",
			expectError:  false,
			fqdnTemplate: "{{.Name}}-{{.Namespace}}.ext-dns.test.com",
		},
		{
			title:        "valid template with multiple hosts",
			expectError:  false,
			fqdnTemplate: "{{.Name}}-{{.Namespace}}.ext-dns.test.com, {{.Name}}-{{.Namespace}}.ext-dna.test.com",
		},
	} {
		t.Run(tt.title, func(t *testing.T) {
			_, err := NewIstioVirtualServiceSource(
				t.Context(),
				fake.NewClientset(),
				istiofake.NewSimpleClientset(),
				"",
				tt.annotationFilter,
				tt.fqdnTemplate,
				false,
				false,
			)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestIstioVirtualServiceSourceFqdnTemplatingExamples(t *testing.T) {
	annotations.SetAnnotationPrefix("external-dns.alpha.kubernetes.io/")
	for _, tt := range []struct {
		title           string
		virtualServices []*networkingv1beta1.VirtualService
		gateways        []*networkingv1beta1.Gateway
		services        []*v1.Service
		fqdnTemplate    string
		combineFqdn     bool
		expected        []*endpoint.Endpoint
	}{
		{
			title:        "simple templating with virtualservice name",
			fqdnTemplate: "{{.Name}}.test.com",
			expected: []*endpoint.Endpoint{
				{DNSName: "app.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
				{DNSName: "my-virtualservice.test.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
			},
			virtualServices: []*networkingv1beta1.VirtualService{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-virtualservice",
						Namespace: "default",
					},
					Spec: istionetworking.VirtualService{
						Hosts:    []string{"app.example.org"},
						Gateways: []string{"my-gateway"},
					},
				},
			},
			gateways: []*networkingv1beta1.Gateway{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-gateway",
						Namespace: "default",
					},
					Spec: istionetworking.Gateway{
						Selector: map[string]string{"istio": "ingressgateway"},
						Servers: []*istionetworking.Server{
							{Hosts: []string{"*"}},
						},
					},
				},
			},
			services: []*v1.Service{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "istio-ingressgateway",
						Namespace: "default",
						Labels:    map[string]string{"istio": "ingressgateway"},
					},
					Spec: v1.ServiceSpec{
						Type:     v1.ServiceTypeLoadBalancer,
						Selector: map[string]string{"istio": "ingressgateway"},
					},
					Status: v1.ServiceStatus{
						LoadBalancer: v1.LoadBalancerStatus{
							Ingress: []v1.LoadBalancerIngress{{IP: "1.2.3.4"}},
						},
					},
				},
			},
		},
		{
			title:        "templating with fqdn combine disabled",
			fqdnTemplate: "{{.Name}}.test.com",
			expected: []*endpoint.Endpoint{
				{DNSName: "app.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
			},
			combineFqdn: true,
			virtualServices: []*networkingv1beta1.VirtualService{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-virtualservice",
						Namespace: "default",
					},
					Spec: istionetworking.VirtualService{
						Hosts:    []string{"app.example.org"},
						Gateways: []string{"my-gateway"},
					},
				},
			},
			gateways: []*networkingv1beta1.Gateway{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-gateway",
						Namespace: "default",
					},
					Spec: istionetworking.Gateway{
						Selector: map[string]string{"istio": "ingressgateway"},
						Servers: []*istionetworking.Server{
							{Hosts: []string{"*"}},
						},
					},
				},
			},
			services: []*v1.Service{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "istio-ingressgateway",
						Namespace: "default",
						Labels:    map[string]string{"istio": "ingressgateway"},
					},
					Spec: v1.ServiceSpec{
						Type:     v1.ServiceTypeLoadBalancer,
						Selector: map[string]string{"istio": "ingressgateway"},
					},
					Status: v1.ServiceStatus{
						LoadBalancer: v1.LoadBalancerStatus{
							Ingress: []v1.LoadBalancerIngress{{IP: "1.2.3.4"}},
						},
					},
				},
			},
		},
		{
			title:        "templating with namespace",
			fqdnTemplate: "{{.Name}}.{{.Namespace}}.cluster.local",
			expected: []*endpoint.Endpoint{
				{DNSName: "api.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"5.6.7.8"}},
				{DNSName: "api-service.production.cluster.local", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"5.6.7.8"}},
				{DNSName: "web.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"::ffff:192.1.56.10"}},
				{DNSName: "web-service.staging.cluster.local", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"::ffff:192.1.56.10"}},
			},
			virtualServices: []*networkingv1beta1.VirtualService{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "api-service",
						Namespace: "production",
					},
					Spec: istionetworking.VirtualService{
						Hosts:    []string{"api.example.org"},
						Gateways: []string{"api-gateway"},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "web-service",
						Namespace: "staging",
					},
					Spec: istionetworking.VirtualService{
						Hosts:    []string{"web.example.org"},
						Gateways: []string{"web-gateway"},
					},
				},
			},
			gateways: []*networkingv1beta1.Gateway{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "api-gateway",
						Namespace: "production",
					},
					Spec: istionetworking.Gateway{
						Selector: map[string]string{"istio": "ingressgateway"},
						Servers: []*istionetworking.Server{
							{Hosts: []string{"*"}},
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "web-gateway",
						Namespace: "staging",
					},
					Spec: istionetworking.Gateway{
						Selector: map[string]string{"istio": "ingressgateway-staging"},
						Servers: []*istionetworking.Server{
							{Hosts: []string{"*"}},
						},
					},
				},
			},
			services: []*v1.Service{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "istio-ingressgateway",
						Namespace: "production",
						Labels:    map[string]string{"istio": "ingressgateway"},
					},
					Spec: v1.ServiceSpec{
						Type:     v1.ServiceTypeLoadBalancer,
						Selector: map[string]string{"istio": "ingressgateway"},
					},
					Status: v1.ServiceStatus{
						LoadBalancer: v1.LoadBalancerStatus{
							Ingress: []v1.LoadBalancerIngress{{IP: "5.6.7.8"}},
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "istio-ingressgateway-staging",
						Namespace: "staging",
						Labels:    map[string]string{"istio": "ingressgateway-staging"},
					},
					Spec: v1.ServiceSpec{
						Type:     v1.ServiceTypeLoadBalancer,
						Selector: map[string]string{"istio": "ingressgateway-staging"},
					},
					Status: v1.ServiceStatus{
						LoadBalancer: v1.LoadBalancerStatus{
							Ingress: []v1.LoadBalancerIngress{{IP: "::ffff:192.1.56.10"}},
						},
					},
				},
			},
		},
		{
			title:        "templating with multiple fqdn templates",
			fqdnTemplate: "{{.Name}}.example.com,{{.Name}}.example.org",
			expected: []*endpoint.Endpoint{
				{DNSName: "multi-host.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"10.0.0.1"}},
				{DNSName: "multi-host.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"10.0.0.1"}},
			},
			virtualServices: []*networkingv1beta1.VirtualService{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "multi-host",
						Namespace: "default",
					},
					Spec: istionetworking.VirtualService{
						Gateways: []string{"my-gateway"},
					},
				},
			},
			gateways: []*networkingv1beta1.Gateway{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-gateway",
						Namespace: "default",
					},
					Spec: istionetworking.Gateway{
						Selector: map[string]string{"istio": "ingressgateway"},
						Servers: []*istionetworking.Server{
							{Hosts: []string{"*"}},
						},
					},
				},
			},
			services: []*v1.Service{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "istio-ingressgateway",
						Namespace: "default",
						Labels:    map[string]string{"istio": "ingressgateway"},
					},
					Spec: v1.ServiceSpec{
						Type:     v1.ServiceTypeLoadBalancer,
						Selector: map[string]string{"istio": "ingressgateway"},
					},
					Status: v1.ServiceStatus{
						LoadBalancer: v1.LoadBalancerStatus{
							Ingress: []v1.LoadBalancerIngress{{IP: "10.0.0.1"}},
						},
					},
				},
			},
		},
		{
			title:        "combine FQDN annotation with template",
			fqdnTemplate: "{{.Name}}.internal.example.com",
			expected: []*endpoint.Endpoint{
				{DNSName: "app.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"172.16.0.1"}},
				{DNSName: "combined-vs.internal.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"172.16.0.1"}},
			},
			virtualServices: []*networkingv1beta1.VirtualService{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "combined-vs",
						Namespace: "default",
					},
					Spec: istionetworking.VirtualService{
						Hosts:    []string{"app.example.org"},
						Gateways: []string{"my-gateway"},
					},
				},
			},
			gateways: []*networkingv1beta1.Gateway{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-gateway",
						Namespace: "default",
					},
					Spec: istionetworking.Gateway{
						Selector: map[string]string{"istio": "ingressgateway"},
						Servers: []*istionetworking.Server{
							{Hosts: []string{"*"}},
						},
					},
				},
			},
			services: []*v1.Service{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "istio-ingressgateway",
						Namespace: "default",
						Labels:    map[string]string{"istio": "ingressgateway"},
					},
					Spec: v1.ServiceSpec{
						Type:     v1.ServiceTypeLoadBalancer,
						Selector: map[string]string{"istio": "ingressgateway"},
					},
					Status: v1.ServiceStatus{
						LoadBalancer: v1.LoadBalancerStatus{
							Ingress: []v1.LoadBalancerIngress{{IP: "172.16.0.1"}},
						},
					},
				},
			},
		},
		{
			title:        "complex templating with labels and hosts",
			fqdnTemplate: "{{ if .Labels.env }}{{.Name}}.{{.Labels.env}}.ex{{ end }}",
			expected: []*endpoint.Endpoint{
				{DNSName: "labeled-vs.dev.ex", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"172.16.0.1"}},
			},
			virtualServices: []*networkingv1beta1.VirtualService{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "labeled-vs",
						Namespace: "default",
						Labels: map[string]string{
							"env": "dev",
						},
					},
					Spec: istionetworking.VirtualService{
						Gateways: []string{"my-gateway"},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "no-labels",
						Namespace: "default",
					},
					Spec: istionetworking.VirtualService{
						Gateways: []string{"my-gateway"},
					},
				},
			},
			gateways: []*networkingv1beta1.Gateway{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-gateway",
						Namespace: "default",
					},
					Spec: istionetworking.Gateway{
						Selector: map[string]string{"istio": "ingressgateway"},
						Servers: []*istionetworking.Server{
							{Hosts: []string{"*"}},
						},
					},
				},
			},
			services: []*v1.Service{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "istio-ingressgateway",
						Namespace: "default",
					},
					Spec: v1.ServiceSpec{
						Type:     v1.ServiceTypeLoadBalancer,
						Selector: map[string]string{"istio": "ingressgateway"},
					},
					Status: v1.ServiceStatus{
						LoadBalancer: v1.LoadBalancerStatus{
							Ingress: []v1.LoadBalancerIngress{{IP: "172.16.0.1"}},
						},
					},
				},
			},
		},
		{
			title:        "templating with cross-namespace gateway reference",
			fqdnTemplate: "{{.Name}}.{{.Namespace}}.svc.cluster.local",
			expected: []*endpoint.Endpoint{
				{DNSName: "cross-ns.example.org", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"lb.example.com"}},
				{DNSName: "cross-ns-vs.app-namespace.svc.cluster.local", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"lb.example.com"}},
			},
			virtualServices: []*networkingv1beta1.VirtualService{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "cross-ns-vs",
						Namespace: "app-namespace",
					},
					Spec: istionetworking.VirtualService{
						Hosts:    []string{"cross-ns.example.org"},
						Gateways: []string{"istio-system/shared-gateway"},
					},
				},
			},
			gateways: []*networkingv1beta1.Gateway{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "shared-gateway",
						Namespace: "istio-system",
					},
					Spec: istionetworking.Gateway{
						Selector: map[string]string{"istio": "ingressgateway"},
						Servers: []*istionetworking.Server{
							{Hosts: []string{"*"}},
						},
					},
				},
			},
			services: []*v1.Service{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "istio-ingressgateway",
						Namespace: "istio-system",
						Labels:    map[string]string{"istio": "ingressgateway"},
					},
					Spec: v1.ServiceSpec{
						Type:     v1.ServiceTypeLoadBalancer,
						Selector: map[string]string{"istio": "ingressgateway"},
					},
					Status: v1.ServiceStatus{
						LoadBalancer: v1.LoadBalancerStatus{
							Ingress: []v1.LoadBalancerIngress{{Hostname: "lb.example.com"}},
						},
					},
				},
			},
		},
		{
			title:        "virtualservice with multiple hosts in spec",
			fqdnTemplate: "{{.Name}}.internal.local",
			expected: []*endpoint.Endpoint{
				{DNSName: "app1.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.1.100"}},
				{DNSName: "app2.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.1.100"}},
				{DNSName: "app3.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.1.100"}},
				{DNSName: "multi-host-vs.internal.local", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.1.100"}},
			},
			virtualServices: []*networkingv1beta1.VirtualService{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "multi-host-vs",
						Namespace: "default",
					},
					Spec: istionetworking.VirtualService{
						Hosts:    []string{"app1.example.org", "app2.example.org", "app3.example.org"},
						Gateways: []string{"my-gateway"},
					},
				},
			},
			gateways: []*networkingv1beta1.Gateway{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-gateway",
						Namespace: "default",
					},
					Spec: istionetworking.Gateway{
						Selector: map[string]string{"istio": "ingressgateway"},
						Servers: []*istionetworking.Server{
							{Hosts: []string{"*"}},
						},
					},
				},
			},
			services: []*v1.Service{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "istio-ingressgateway",
						Namespace: "default",
						Labels:    map[string]string{"istio": "ingressgateway"},
					},
					Spec: v1.ServiceSpec{
						Type:     v1.ServiceTypeLoadBalancer,
						Selector: map[string]string{"istio": "ingressgateway"},
					},
					Status: v1.ServiceStatus{
						LoadBalancer: v1.LoadBalancerStatus{
							Ingress: []v1.LoadBalancerIngress{{IP: "192.168.1.100"}},
						},
					},
				},
			},
		},
		{
			title:        "virtualservice with no matching gateway (no endpoints from spec)",
			fqdnTemplate: "{{.Name}}.fallback.local",
			expected: []*endpoint.Endpoint{
				{DNSName: "orphan.example.org", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"fallback.local"}},
			},
			virtualServices: []*networkingv1beta1.VirtualService{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "orphan-vs",
						Namespace: "default",
						Annotations: map[string]string{
							annotations.TargetKey: "fallback.local",
						},
					},
					Spec: istionetworking.VirtualService{
						Hosts:    []string{"orphan.example.org"},
						Gateways: []string{"non-existent-gateway"},
					},
				},
			},
		},
		{
			title:        "templating with annotations expansion",
			fqdnTemplate: `{{ index .ObjectMeta.Annotations "dns.company.com/subdomain" }}.company.local`,
			expected: []*endpoint.Endpoint{
				{DNSName: "api-v2.company.local", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.1.100"}},
			},
			virtualServices: []*networkingv1beta1.VirtualService{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "annotated-vs",
						Namespace: "default",
						Annotations: map[string]string{
							"dns.company.com/subdomain": "api-v2",
							annotations.TargetKey:       "10.20.30.40",
						},
					},
					Spec: istionetworking.VirtualService{
						Gateways: []string{"my-gateway"},
					},
				},
			},
			gateways: []*networkingv1beta1.Gateway{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-gateway",
						Namespace: "default",
					},
					Spec: istionetworking.Gateway{
						Selector: map[string]string{"istio": "ingressgateway"},
						Servers: []*istionetworking.Server{
							{Hosts: []string{"*"}},
						},
					},
				},
			},
			services: []*v1.Service{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "istio",
						Namespace: "default",
					},
					Spec: v1.ServiceSpec{
						Type:     v1.ServiceTypeLoadBalancer,
						Selector: map[string]string{"istio": "ingressgateway"},
					},
					Status: v1.ServiceStatus{
						LoadBalancer: v1.LoadBalancerStatus{
							Ingress: []v1.LoadBalancerIngress{{IP: "192.168.1.100"}},
						},
					},
				},
			},
		},
	} {
		t.Run(tt.title, func(t *testing.T) {
			kubeClient := fake.NewClientset()
			istioClient := istiofake.NewSimpleClientset()

			for _, svc := range tt.services {
				_, err := kubeClient.CoreV1().Services(svc.Namespace).Create(t.Context(), svc, metav1.CreateOptions{})
				require.NoError(t, err)
			}

			for _, gw := range tt.gateways {
				_, err := istioClient.NetworkingV1beta1().Gateways(gw.Namespace).Create(t.Context(), gw, metav1.CreateOptions{})
				require.NoError(t, err)
			}

			for _, vs := range tt.virtualServices {
				_, err := istioClient.NetworkingV1beta1().VirtualServices(vs.Namespace).Create(t.Context(), vs, metav1.CreateOptions{})
				require.NoError(t, err)
			}

			src, err := NewIstioVirtualServiceSource(
				t.Context(),
				kubeClient,
				istioClient,
				"",
				"",
				tt.fqdnTemplate,
				!tt.combineFqdn,
				false,
			)
			require.NoError(t, err)

			endpoints, err := src.Endpoints(t.Context())
			require.NoError(t, err)

			validateEndpoints(t, endpoints, tt.expected)
		})
	}
}
