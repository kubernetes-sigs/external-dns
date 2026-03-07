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

	routev1 "github.com/openshift/api/route/v1"
	"github.com/openshift/client-go/route/clientset/versioned/fake"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/intstr"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/source/annotations"
)

func TestOpenShiftFqdnTemplatingExamples(t *testing.T) {
	for _, tt := range []struct {
		title        string
		ocpRoute     []*routev1.Route
		fqdnTemplate string
		combineFqdn  bool
		expected     []*endpoint.Endpoint
	}{
		{
			title:        "simple templating",
			fqdnTemplate: "{{.Name}}.tld.com",
			expected: []*endpoint.Endpoint{
				{DNSName: "example.org", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"router-default.example.com"}},
				{DNSName: "my-gateway.tld.com", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"router-default.example.com"}},
			},
			ocpRoute: []*routev1.Route{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-gateway",
						Namespace: "default",
					},
					Spec: routev1.RouteSpec{
						Host: "example.org",
						To: routev1.RouteTargetReference{
							Kind: "Service",
							Name: "my-service",
						},
						TLS: &routev1.TLSConfig{},
					},
					Status: routev1.RouteStatus{
						Ingress: []routev1.RouteIngress{
							{
								Host:                    "example.org",
								RouterCanonicalHostname: "router-default.example.com",
								Conditions: []routev1.RouteIngressCondition{
									{
										Type:   routev1.RouteAdmitted,
										Status: corev1.ConditionTrue,
									},
								},
							},
						},
					},
				},
			},
		},
		{
			title:        "templating with fqdn combine disabled",
			fqdnTemplate: "{{.Name}}.tld.com",
			expected: []*endpoint.Endpoint{
				{DNSName: "example.org", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"router-default.example.com"}},
			},
			combineFqdn: true,
			ocpRoute: []*routev1.Route{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-gateway",
						Namespace: "default",
					},
					Spec: routev1.RouteSpec{},
					Status: routev1.RouteStatus{
						Ingress: []routev1.RouteIngress{
							{
								Host:                    "example.org",
								RouterCanonicalHostname: "router-default.example.com",
								Conditions: []routev1.RouteIngressCondition{
									{
										Type:   routev1.RouteAdmitted,
										Status: corev1.ConditionTrue,
									},
								},
							},
						},
					},
				},
			},
		},
		{
			title:        "templating with namespace",
			fqdnTemplate: "{{.Name}}.{{.Namespace}}.tld.com",
			expected: []*endpoint.Endpoint{
				{DNSName: "my-gateway.kube-system.tld.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"10.1.1.0"}},
			},
			combineFqdn: true,
			ocpRoute: []*routev1.Route{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-gateway",
						Namespace: "kube-system",
						Annotations: map[string]string{
							annotations.TargetKey: "10.1.1.0",
						},
					},
				},
			},
		},
		{
			title:        "templating with complex fqdn template",
			fqdnTemplate: "{{ .Name }}.{{ .Namespace }}.tld.com,{{ if .Labels.env }}{{ .Labels.env }}.private{{ end }}",
			expected: []*endpoint.Endpoint{
				{DNSName: "no-labels-route-3.default.tld.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"10.1.1.3"}},
				{DNSName: "route-2.default.tld.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"10.1.1.3"}},
				{DNSName: "dev.private", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"10.1.1.3"}},
				{DNSName: "route-1.kube-system.tld.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"10.1.1.0"}},
				{DNSName: "prod.private", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"10.1.1.0"}},
			},
			combineFqdn: true,
			ocpRoute: []*routev1.Route{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "route-1",
						Namespace: "kube-system",
						Labels: map[string]string{
							"env": "prod",
						},
						Annotations: map[string]string{
							"env":                 "prod",
							annotations.TargetKey: "10.1.1.0",
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "route-2",
						Namespace: "default",
						Labels: map[string]string{
							"env": "dev",
						},
						Annotations: map[string]string{
							annotations.TargetKey: "10.1.1.3",
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "no-labels-route-3",
						Namespace: "default",
						Annotations: map[string]string{
							annotations.TargetKey: "10.1.1.3",
						},
					},
				},
			},
		},
		{
			title:        "template that skips when field is missing",
			fqdnTemplate: "{{ if and .Spec.Port .Spec.Port.TargetPort }}{{ .Name }}.{{ .Spec.Port.TargetPort }}.tld.com{{ end }}",
			expected: []*endpoint.Endpoint{
				{DNSName: "route-1.80.tld.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"10.1.1.0"}},
			},
			ocpRoute: []*routev1.Route{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "route-1",
						Namespace: "kube-system",
						Annotations: map[string]string{
							annotations.TargetKey: "10.1.1.0",
						},
					},
					Spec: routev1.RouteSpec{
						Port: &routev1.RoutePort{
							TargetPort: intstr.FromString("80"),
						},
					},
					Status: routev1.RouteStatus{},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "route-2",
						Namespace: "default",
						Annotations: map[string]string{
							annotations.TargetKey: "10.1.1.3",
						},
					},
				},
			},
		},
		{
			title:        "get canonical hostnames for admitted routes",
			fqdnTemplate: "{{ $name := .Name }}{{ range $ingress := .Status.Ingress }}{{ range $ingress.Conditions }}{{ if and (eq .Type \"Admitted\") (eq .Status \"True\") }}{{ $ingress.Host }},{{ end }}{{ end }}{{ end }}",
			expected: []*endpoint.Endpoint{
				{DNSName: "cluster.example.org", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"router-dmz.apps.dmz.example.com"}},
				{DNSName: "apps.example.org", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"router-dmz.apps.dmz.example.com"}},
			},
			ocpRoute: []*routev1.Route{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:        "my-route",
						Namespace:   "kube-system",
						Annotations: map[string]string{},
					},
					Status: routev1.RouteStatus{
						Ingress: []routev1.RouteIngress{
							{
								Host:                    "cluster.example.org",
								RouterCanonicalHostname: "router-dmz.apps.dmz.example.com",
								Conditions: []routev1.RouteIngressCondition{
									{
										Type:   routev1.RouteAdmitted,
										Status: corev1.ConditionTrue,
									},
								},
							},
							{
								Host:                    "apps.example.org",
								RouterCanonicalHostname: "router-internal.apps.internal.example.com",
								Conditions: []routev1.RouteIngressCondition{
									{
										Type:   routev1.RouteAdmitted,
										Status: corev1.ConditionTrue,
									},
								},
							},
							{
								Host:                    "wrong.example.org",
								RouterCanonicalHostname: "router-default.apps.cluster.example.com",
								Conditions: []routev1.RouteIngressCondition{
									{
										Type:   routev1.RouteAdmitted,
										Status: corev1.ConditionFalse,
									},
								},
							},
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "route-2",
						Namespace: "default",
						Annotations: map[string]string{
							annotations.TargetKey: "10.1.1.3",
						},
					},
				},
			},
		},
		{
			title:        "get canonical hostnames for admitted routes without prefix",
			fqdnTemplate: "{{ $name := .Name }}{{ range $ingress := .Status.Ingress }}{{ range $ingress.Conditions }}{{ if and (eq .Type \"Admitted\") (eq .Status \"True\") }}{{ with $ingress.RouterCanonicalHostname }}{{ $name }}.{{ trimPrefix . \"router-\" }},{{ end }}{{ end }}{{ end }}{{ end }}",
			expected: []*endpoint.Endpoint{
				{DNSName: "cluster.example.org", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"router-dmz.apps.dmz.example.com"}},
				{DNSName: "my-route.dmz.apps.dmz.example.com", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"router-dmz.apps.dmz.example.com"}},
				{DNSName: "my-route.internal.apps.internal.example.com", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"router-dmz.apps.dmz.example.com"}},
			},
			ocpRoute: []*routev1.Route{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:        "my-route",
						Namespace:   "kube-system",
						Annotations: map[string]string{},
					},
					Status: routev1.RouteStatus{
						Ingress: []routev1.RouteIngress{
							{
								Host:                    "cluster.example.org",
								RouterCanonicalHostname: "router-dmz.apps.dmz.example.com",
								Conditions: []routev1.RouteIngressCondition{
									{
										Type:   routev1.RouteAdmitted,
										Status: corev1.ConditionTrue,
									},
								},
							},
							{
								Host:                    "apps.example.org",
								RouterCanonicalHostname: "router-internal.apps.internal.example.com",
								Conditions: []routev1.RouteIngressCondition{
									{
										Type:   routev1.RouteAdmitted,
										Status: corev1.ConditionTrue,
									},
								},
							},
							{
								Host:                    "wrong.example.org",
								RouterCanonicalHostname: "router-default.apps.cluster.example.com",
								Conditions: []routev1.RouteIngressCondition{
									{
										Type:   routev1.RouteAdmitted,
										Status: corev1.ConditionFalse,
									},
								},
							},
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "route-2",
						Namespace: "default",
						Annotations: map[string]string{
							annotations.TargetKey: "10.1.1.3",
						},
					},
				},
			},
		},
	} {
		t.Run(tt.title, func(t *testing.T) {
			kubeClient := fake.NewClientset()
			for _, ocp := range tt.ocpRoute {
				_, err := kubeClient.RouteV1().Routes(ocp.Namespace).Create(context.Background(), ocp, metav1.CreateOptions{})
				require.NoError(t, err)
			}

			src, err := NewOcpRouteSource(
				t.Context(),
				kubeClient,
				"",
				"",
				tt.fqdnTemplate,
				!tt.combineFqdn,
				false,
				labels.Everything(),
				"",
			)
			require.NoError(t, err)

			endpoints, err := src.Endpoints(t.Context())
			require.NoError(t, err)

			validateEndpoints(t, endpoints, tt.expected)
		})
	}
}
