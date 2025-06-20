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
	"sigs.k8s.io/external-dns/internal/testutils"

	networkv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes/fake"
	"sigs.k8s.io/external-dns/endpoint"
)

func TestIngressSourceNewNodeSourceWithFqdn(t *testing.T) {
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
			title:        "complex template",
			expectError:  false,
			fqdnTemplate: "{{range .Status.Addresses}}{{if and (eq .Type \"ExternalIP\") (isIPv4 .Address)}}{{.Address | replace \".\" \"-\"}}{{break}}{{end}}{{end}}.ext-dns.test.com",
		},
		{
			title:        "valid template with multiple hosts",
			expectError:  false,
			fqdnTemplate: "{{.Name}}-{{.Namespace}}.ext-dns.test.com, {{.Name}}-{{.Namespace}}.ext-dna.test.com",
		},
	} {
		t.Run(tt.title, func(t *testing.T) {
			_, err := NewIngressSource(
				t.Context(),
				fake.NewClientset(),
				"",
				"",
				tt.fqdnTemplate,
				false,
				false,
				false,
				false,
				labels.Everything(),
				[]string{},
			)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestIngressSourceFqdnTemplatingExamples(t *testing.T) {

	for _, tt := range []struct {
		title        string
		ingresses    []*networkv1.Ingress
		fqdnTemplate string
		expected     []*endpoint.Endpoint
	}{
		{
			title: "templating resolve Ingress source hostnames to IP",
			ingresses: []*networkv1.Ingress{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-ingress",
						Namespace: "default",
					},
					Spec: networkv1.IngressSpec{
						IngressClassName: testutils.ToPtr("my-ingress"),
						Rules: []networkv1.IngressRule{
							{
								Host: "example.org",
								IngressRuleValue: networkv1.IngressRuleValue{
									HTTP: &networkv1.HTTPIngressRuleValue{
										Paths: []networkv1.HTTPIngressPath{
											{
												Backend: networkv1.IngressBackend{
													Service: &networkv1.IngressServiceBackend{
														Name: "my-service",
														Port: networkv1.ServiceBackendPort{
															Name: "http",
														},
													},
												},
												PathType: testutils.ToPtr(networkv1.PathTypePrefix),
												Path:     "/",
											},
										},
									},
								},
							},
						},
					},
					Status: networkv1.IngressStatus{
						LoadBalancer: networkv1.IngressLoadBalancerStatus{
							Ingress: []networkv1.IngressLoadBalancerIngress{
								{Hostname: "10.200.130.84.nip.io"},
							},
						},
					},
				},
			},
			fqdnTemplate: "{{.Name }}.nip.io",
			expected: []*endpoint.Endpoint{
				{DNSName: "example.org", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"10.200.130.84.nip.io"}},
				{DNSName: "my-ingress.nip.io", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"10.200.130.84.nip.io"}},
			},
		},
		{
			title: "templating resolve hostnames with nip.io",
			ingresses: []*networkv1.Ingress{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-ingress",
						Namespace: "default",
					},
					Spec: networkv1.IngressSpec{
						IngressClassName: testutils.ToPtr("my-ingress"),
						Rules: []networkv1.IngressRule{
							{Host: "example.org"},
						},
					},
					Status: networkv1.IngressStatus{
						LoadBalancer: networkv1.IngressLoadBalancerStatus{
							Ingress: []networkv1.IngressLoadBalancerIngress{
								{Hostname: "10.200.130.84.nip.io"},
							},
						},
					},
				},
			},
			fqdnTemplate: `{{ range .Status.LoadBalancer.Ingress }}{{ if contains .Hostname "nip.io" }}example.org{{end}}{{end}}`,
			expected: []*endpoint.Endpoint{
				{DNSName: "example.org", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"10.200.130.84.nip.io"}},
				{DNSName: "example.org", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"10.200.130.84.nip.io"}},
			},
		},
		{
			title: "templating resolve hostnames with nip.io and target annotation",
			ingresses: []*networkv1.Ingress{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-ingress",
						Namespace: "default",
						Annotations: map[string]string{
							"external-dns.alpha.kubernetes.io/target": "10.200.130.84",
						},
					},
					Spec: networkv1.IngressSpec{
						IngressClassName: testutils.ToPtr("my-ingress"),
						Rules: []networkv1.IngressRule{
							{Host: "example.org"},
						},
					},
					Status: networkv1.IngressStatus{
						LoadBalancer: networkv1.IngressLoadBalancerStatus{
							Ingress: []networkv1.IngressLoadBalancerIngress{
								{Hostname: "10.200.130.84.nip.io"},
							},
						},
					},
				},
			},
			fqdnTemplate: `{{ range .Status.LoadBalancer.Ingress }}{{ if contains .Hostname "nip.io" }}tld.org{{break}}{{end}}{{end}}`,
			expected: []*endpoint.Endpoint{
				{DNSName: "example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"10.200.130.84"}},
				{DNSName: "tld.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"10.200.130.84"}},
			},
		},
		{
			title: "templating resolve hostnames with nip.io and status IP",
			ingresses: []*networkv1.Ingress{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-ingress",
						Namespace: "default",
					},
					Spec: networkv1.IngressSpec{
						IngressClassName: testutils.ToPtr("my-ingress"),
						Rules: []networkv1.IngressRule{
							{
								Host: "example.org",
							},
						},
					},
					Status: networkv1.IngressStatus{
						LoadBalancer: networkv1.IngressLoadBalancerStatus{
							Ingress: []networkv1.IngressLoadBalancerIngress{
								{
									IP: "10.200.130.84",
								},
							},
						},
					},
				},
			},
			fqdnTemplate: "nip.io",
			expected: []*endpoint.Endpoint{
				{DNSName: "example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"10.200.130.84"}},
				{DNSName: "nip.io", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"10.200.130.84"}},
			},
		},
		{
			title: "templating resolve with different hostnames and rules",
			ingresses: []*networkv1.Ingress{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-ingress",
						Namespace: "default",
					},
					Spec: networkv1.IngressSpec{
						IngressClassName: testutils.ToPtr("ingress-with-override"),
						Rules: []networkv1.IngressRule{
							{Host: "foo.bar.com"},
							{Host: "bar.bar.com"},
							{Host: "bar.baz.com"},
						},
					},
					Status: networkv1.IngressStatus{
						LoadBalancer: networkv1.IngressLoadBalancerStatus{
							Ingress: []networkv1.IngressLoadBalancerIngress{
								{IP: "192.16.15.25"},
								{Hostname: "abc.org"},
							},
						},
					},
				},
			},
			fqdnTemplate: `{{ range .Spec.Rules }}{{ if contains .Host "bar.com" }}{{ .Host }}.internal{{break}}{{end}}{{end}}`,
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.bar.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.16.15.25"}},
				{DNSName: "foo.bar.com", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"abc.org"}},
				{DNSName: "bar.bar.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.16.15.25"}},
				{DNSName: "bar.bar.com", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"abc.org"}},
				{DNSName: "bar.baz.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.16.15.25"}},
				{DNSName: "bar.baz.com", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"abc.org"}},
				{DNSName: "foo.bar.com.internal", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.16.15.25"}},
				{DNSName: "foo.bar.com.internal", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"abc.org"}},
			},
		},
		{
			title: "templating resolve with rules and tls",
			ingresses: []*networkv1.Ingress{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-ingress",
						Namespace: "default",
					},
					Spec: networkv1.IngressSpec{
						IngressClassName: testutils.ToPtr("ingress-with-override"),
						Rules: []networkv1.IngressRule{
							{
								Host: "foo.bar.com",
							},
						},
						TLS: []networkv1.IngressTLS{
							{
								Hosts: []string{"https-example.foo.com", "https-example.bar.com"},
							},
						},
					},
					Status: networkv1.IngressStatus{
						LoadBalancer: networkv1.IngressLoadBalancerStatus{
							Ingress: []networkv1.IngressLoadBalancerIngress{
								{
									IP: "10.09.15.25",
								},
							},
						},
					},
				},
			},
			fqdnTemplate: `{{ .Name }}.test.org,{{ range .Spec.TLS }}{{ range $value := .Hosts }}{{ $value | replace "." "-" }}.internal{{break}}{{end}}{{end}}`,
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.bar.com", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"10.09.15.25"}},
				{DNSName: "https-example.foo.com", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"10.09.15.25"}},
				{DNSName: "https-example.bar.com", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"10.09.15.25"}},
				{DNSName: "my-ingress.test.org", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"10.09.15.25"}},
				{DNSName: "https-example-foo-com.internal", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"10.09.15.25"}},
			},
		},
	} {
		t.Run(tt.title, func(t *testing.T) {
			kubeClient := fake.NewClientset()

			for _, el := range tt.ingresses {
				_, err := kubeClient.NetworkingV1().Ingresses(el.Namespace).Create(t.Context(), el, metav1.CreateOptions{})
				require.NoError(t, err)
			}

			src, err := NewIngressSource(
				t.Context(),
				kubeClient,
				"",
				"",
				tt.fqdnTemplate,
				true,
				false,
				false,
				false,
				labels.Everything(),
				[]string{},
			)

			require.NoError(t, err)

			endpoints, err := src.Endpoints(t.Context())
			require.NoError(t, err)

			validateEndpoints(t, endpoints, tt.expected)
		})
	}
}
