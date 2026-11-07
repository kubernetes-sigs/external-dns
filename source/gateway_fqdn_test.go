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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubefake "k8s.io/client-go/kubernetes/fake"
	v1 "sigs.k8s.io/gateway-api/apis/v1"
	gatewayfake "sigs.k8s.io/gateway-api/pkg/client/clientset/versioned/fake"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	"sigs.k8s.io/external-dns/source/config"
	templatetest "sigs.k8s.io/external-dns/source/template/testutil"
)

type gatewayRouteTemplateCase struct {
	title              string
	hostnames          []v1.Hostname
	fqdnTemplate       string
	targetTemplate     string
	fqdnTargetTemplate string
	combine            bool
	expected           []*endpoint.Endpoint
}

func TestGatewayRouteFQDNTargetTemplate(t *testing.T) {
	runGatewayRouteTemplateCases(t, []gatewayRouteTemplateCase{
		{
			title:              "pairs are published without listener matching",
			fqdnTargetTemplate: "{{.Name}}.other.org:5.6.7.8",
			expected: []*endpoint.Endpoint{
				newTestEndpoint("templated-route.other.org", "5.6.7.8"),
			},
		},
		{
			title:              "spec hostnames suppress pairs without combine",
			hostnames:          []v1.Hostname{"spec.example.internal"},
			fqdnTargetTemplate: "{{.Name}}.other.org:5.6.7.8",
			expected: []*endpoint.Endpoint{
				newTestEndpoint("spec.example.internal", "1.2.3.4"),
			},
		},
	})
}

func TestGatewayRouteTargetTemplate(t *testing.T) {
	runGatewayRouteTemplateCases(t, []gatewayRouteTemplateCase{
		{
			title:          "target template replaces gateway addresses for template hostnames",
			fqdnTemplate:   "{{.Name}}.example.internal",
			targetTemplate: "lb.example.net",
			expected: []*endpoint.Endpoint{
				newTestEndpointWithTTL("templated-route.example.internal", endpoint.RecordTypeCNAME, 0, "lb.example.net"),
			},
		},
		{
			title:          "spec hostnames keep gateway addresses when combined",
			hostnames:      []v1.Hostname{"spec.example.internal"},
			fqdnTemplate:   "{{.Name}}.example.internal",
			targetTemplate: "lb.example.net",
			combine:        true,
			expected: []*endpoint.Endpoint{
				newTestEndpoint("spec.example.internal", "1.2.3.4"),
				newTestEndpointWithTTL("templated-route.example.internal", endpoint.RecordTypeCNAME, 0, "lb.example.net"),
			},
		},
	})
}

func runGatewayRouteTemplateCases(t *testing.T, cases []gatewayRouteTemplateCase) {
	t.Helper()
	for _, tt := range cases {
		t.Run(tt.title, func(t *testing.T) {
			t.Parallel()

			gwClient := gatewayfake.NewSimpleClientset()
			_, err := gwClient.GatewayV1().Gateways("default").Create(t.Context(), &v1.Gateway{
				Namespace: "default", Name: "test",
				Spec: v1.GatewaySpec{
					Listeners: []v1.Listener{{
						Protocol: v1.HTTPProtocolType,
						Hostname: new(v1.Hostname("*.example.internal")),
					}},
				},
				Status: gatewayStatus("1.2.3.4"),
			}, metav1.CreateOptions{})
			require.NoError(t, err)

			_, err = gwClient.GatewayV1().HTTPRoutes("default").Create(t.Context(), &v1.HTTPRoute{
				Namespace: "default", Name: "templated-route",
				Spec: v1.HTTPRouteSpec{
					CommonRouteSpec: v1.CommonRouteSpec{
						ParentRefs: []v1.ParentReference{gwParentRef("default", "test")},
					},
					Hostnames: tt.hostnames,
				},
				Status: httpRouteStatus(gwParentRef("default", "test")),
			}, metav1.CreateOptions{})
			require.NoError(t, err)

			kubeClient := kubefake.NewClientset()
			_, err = kubeClient.CoreV1().Namespaces().Create(t.Context(), &corev1.Namespace{
				Name: "default",
			}, metav1.CreateOptions{})
			require.NoError(t, err)

			clients := new(testutils.MockClientGenerator)
			clients.On("GatewayClient").Return(gwClient, nil)
			clients.On("KubeClient").Return(kubeClient, nil)

			src, err := NewGatewayHTTPRouteSource(t.Context(), clients, &config.Config{
				TemplateEngine: templatetest.MustEngine(t, tt.fqdnTemplate, tt.targetTemplate, tt.fqdnTargetTemplate, tt.combine),
			})
			require.NoError(t, err)

			endpoints, err := src.Endpoints(t.Context())
			require.NoError(t, err)

			testutils.ValidateEndpoints(t, endpoints, tt.expected)
		})
	}
}
