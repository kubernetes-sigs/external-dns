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

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	istionetworking "istio.io/api/networking/v1beta1"
	networkingv1beta1 "istio.io/client-go/pkg/apis/networking/v1beta1"
	istiofake "istio.io/client-go/pkg/clientset/versioned/fake"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes/fake"

	"sigs.k8s.io/external-dns/source/annotations"

	"sigs.k8s.io/external-dns/endpoint"
)

func TestIstioGatewaySourceNewSourceWithFqdn(t *testing.T) {
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
			_, err := NewIstioGatewaySource(
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

func TestIstioGatewaySourceFqdnTemplatingExamples(t *testing.T) {
	for _, tt := range []struct {
		title        string
		gateways     []*networkingv1beta1.Gateway
		services     []*v1.Service
		fqdnTemplate string
		combineFqdn  bool
		expected     []*endpoint.Endpoint
	}{
		{
			title:        "simple templating with gateway name",
			fqdnTemplate: "{{.Name}}.test.com",
			expected: []*endpoint.Endpoint{
				{DNSName: "example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
				{DNSName: "my-gateway.test.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
			},
			gateways: []*networkingv1beta1.Gateway{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-gateway",
						Namespace: "istio-system",
					},
					Spec: istionetworking.Gateway{
						Selector: map[string]string{"istio": "ingressgateway"},
						Servers: []*istionetworking.Server{
							{Hosts: []string{"example.org"}},
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
				{DNSName: "example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
			},
			combineFqdn: true,
			gateways: []*networkingv1beta1.Gateway{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-gateway",
						Namespace: "istio-system",
					},
					Spec: istionetworking.Gateway{
						Selector: map[string]string{"istio": "ingressgateway"},
						Servers: []*istionetworking.Server{
							{Hosts: []string{"example.org"}},
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
				{DNSName: "api-gateway.kube-system.cluster.local", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"::ffff:192.1.56.10"}},
				{DNSName: "api-gateway.production.cluster.local", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"5.6.7.8"}},
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
							{Hosts: []string{"api.example.org"}},
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "api-gateway",
						Namespace: "kube-system",
					},
					Spec: istionetworking.Gateway{
						Selector: map[string]string{"istio": "ingressgateway-extra"},
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
						Name:      "kube-metrics-server",
						Namespace: "kube-system",
						Labels:    map[string]string{"istio": "ingressgateway-extra"},
					},
					Spec: v1.ServiceSpec{
						Type:     v1.ServiceTypeLoadBalancer,
						Selector: map[string]string{"istio": "ingressgateway-extra"},
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
			title:        "templating with complex fqdn template",
			fqdnTemplate: "{{.Name}}.example.com,{{.Name}}.example.org",
			expected: []*endpoint.Endpoint{
				{DNSName: "multi-gateway.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"10.0.0.1"}},
				{DNSName: "multi-gateway.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"10.0.0.1"}},
			},
			gateways: []*networkingv1beta1.Gateway{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "multi-gateway",
						Namespace: "default",
					},
					Spec: istionetworking.Gateway{
						Selector: map[string]string{"istio": "ingressgateway"},
						Servers:  []*istionetworking.Server{},
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
				{DNSName: "combined-gateway.internal.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"172.16.0.1"}},
			},
			gateways: []*networkingv1beta1.Gateway{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "combined-gateway",
						Namespace: "default",
					},
					Spec: istionetworking.Gateway{
						Selector: map[string]string{"istio": "ingressgateway"},
						Servers: []*istionetworking.Server{
							{Hosts: []string{"app.example.org"}},
						},
					},
				},
			},
			services: []*v1.Service{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "istio-ingressgateway",
						Namespace: "default",
						Labels: map[string]string{
							"istio": "ingressgateway",
						},
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
			title: "templating with labels",
			gateways: []*networkingv1beta1.Gateway{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "labeled-gateway",
						Namespace: "default",
						Labels: map[string]string{
							"environment": "staging",
						},
						Annotations: map[string]string{
							annotations.TargetKey: "203.0.113.1",
						},
					},
					Spec: istionetworking.Gateway{
						Selector: map[string]string{"istio": "ingressgateway"},
						Servers:  []*istionetworking.Server{},
					},
				},
			},
			fqdnTemplate: "{{.Name}}.{{.Labels.environment}}.example.com",
			expected: []*endpoint.Endpoint{
				{DNSName: "labeled-gateway.staging.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"203.0.113.1"}},
			},
		},
		{
			title:        "srv record with node port and cluster ip services without external ips",
			fqdnTemplate: "{{.Name}}.example.com",
			expected:     []*endpoint.Endpoint{},
			gateways: []*networkingv1beta1.Gateway{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "labeled-gateway",
						Namespace: "default",
					},
					Spec: istionetworking.Gateway{
						Selector: map[string]string{"istio": "ingressgateway"},
						Servers:  []*istionetworking.Server{},
					},
				},
			},
			services: []*v1.Service{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-node-port",
						Namespace: "default",
						Labels: map[string]string{
							"istio": "ingressgateway",
						},
					},
					Spec: v1.ServiceSpec{
						Type:      v1.ServiceTypeNodePort,
						Selector:  map[string]string{"istio": "ingressgateway"},
						ClusterIP: "10.96.41.133",
						Ports: []v1.ServicePort{
							{Name: "dns", Port: 8082, TargetPort: intstr.FromInt32(8083), Protocol: v1.ProtocolUDP, NodePort: 30083},
							{Name: "dns-tcp", Port: 2525, TargetPort: intstr.FromInt32(25256), NodePort: 25565},
						},
					},
					Status: v1.ServiceStatus{
						LoadBalancer: v1.LoadBalancerStatus{},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-cluster-ip",
						Namespace: "default",
						Labels:    map[string]string{"istio": "ingressgateway"},
					},
					Spec: v1.ServiceSpec{
						Type:      v1.ServiceTypeClusterIP,
						Selector:  map[string]string{"istio": "ingressgateway"},
						ClusterIP: "10.96.41.133",
						Ports: []v1.ServicePort{
							{Name: "dns", Port: 53, TargetPort: intstr.FromInt32(30053), Protocol: v1.ProtocolUDP},
							{Name: "dns-tcp", Port: 53, TargetPort: intstr.FromInt32(30054), NodePort: 25565},
						},
					},
				},
			},
		},
		{
			title:        "srv record with node port and cluster ip services with external ips",
			fqdnTemplate: "{{.Name}}.tld.org",
			expected: []*endpoint.Endpoint{
				{DNSName: "nodeport-external.tld.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.132.253"}},
			},
			gateways: []*networkingv1beta1.Gateway{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "nodeport-external",
						Namespace: "default",
					},
					Spec: istionetworking.Gateway{
						Selector: map[string]string{"istio": "ingressgateway"},
					},
				},
			},
			services: []*v1.Service{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-node-port",
						Namespace: "default",
						Labels: map[string]string{
							"istio": "ingressgateway",
						},
					},
					Spec: v1.ServiceSpec{
						Type:        v1.ServiceTypeNodePort,
						Selector:    map[string]string{"istio": "ingressgateway"},
						ClusterIP:   "10.96.41.133",
						ExternalIPs: []string{"192.168.132.253"},
						Ports: []v1.ServicePort{
							{Name: "dns", Port: 8082, TargetPort: intstr.FromInt32(8083), Protocol: v1.ProtocolUDP, NodePort: 30083},
							{Name: "dns-tcp", Port: 2525, TargetPort: intstr.FromInt32(25256), NodePort: 25565},
						},
					},
				},
			},
		},
		{
			title:        "with host as subdomain in reversed order",
			fqdnTemplate: "{{ range $server := .Spec.Servers }}{{ range $host := $server.Hosts }}{{ $host }}.{{ $server.Port.Name }}.{{ $server.Port.Number }}.tld.com,{{ end }}{{ end }}",
			expected: []*endpoint.Endpoint{
				{DNSName: "www.bookinfo", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.132.253"}},
				{DNSName: "bookinfo", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.132.253"}},
				{DNSName: "www.bookinfo.http.443.tld.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.132.253"}},
				{DNSName: "bookinfo.dns.8080.tld.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.132.253"}},
			},
			gateways: []*networkingv1beta1.Gateway{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "nodeport-external",
						Namespace: "default",
					},
					Spec: istionetworking.Gateway{
						Selector: map[string]string{"istio": "ingressgateway"},
						Servers: []*istionetworking.Server{
							{
								Hosts: []string{"www.bookinfo"},
								Name:  "main",
								Port:  &istionetworking.Port{Number: 443, Name: "http", Protocol: "HTTPS"},
							},
							{
								Hosts: []string{"bookinfo"},
								Name:  "debug",
								Port:  &istionetworking.Port{Number: 8080, Name: "dns", Protocol: "UDP"},
							},
						},
					},
				},
			},
			services: []*v1.Service{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-node-port",
						Namespace: "default",
						Labels: map[string]string{
							"istio": "ingressgateway",
						},
					},
					Spec: v1.ServiceSpec{
						Type:        v1.ServiceTypeNodePort,
						Selector:    map[string]string{"istio": "ingressgateway"},
						ClusterIP:   "10.96.41.133",
						ExternalIPs: []string{"192.168.132.253"},
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

			src, err := NewIstioGatewaySource(
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
