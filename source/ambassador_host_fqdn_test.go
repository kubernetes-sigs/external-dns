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

	ambassador "github.com/emissary-ingress/emissary/v3/pkg/api/getambassador.io/v3alpha1"
	"github.com/stretchr/testify/require"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	fakeDynamic "k8s.io/client-go/dynamic/fake"
	fakeKube "k8s.io/client-go/kubernetes/fake"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	"sigs.k8s.io/external-dns/source/config"
	templatetest "sigs.k8s.io/external-dns/source/template/testutil"
)

func TestAmbassadorHostFQDNTemplate(t *testing.T) {
	for _, tt := range []struct {
		title          string
		targetTemplate string
		expected       []*endpoint.Endpoint
	}{
		{
			title: "fqdn-template uses the ambassador service load balancer",
			expected: []*endpoint.Endpoint{
				{DNSName: "tmpl-host.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.1.1.1"}},
			},
		},
		{
			title:          "target-template replaces the ambassador service load balancer",
			targetTemplate: "lb.example.net",
			expected: []*endpoint.Endpoint{
				{DNSName: "tmpl-host.example.org", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"lb.example.net"}},
			},
		},
	} {
		t.Run(tt.title, func(t *testing.T) {
			t.Parallel()

			fakeKubernetesClient := fakeKube.NewSimpleClientset()
			_, err := fakeKubernetesClient.CoreV1().Services(defaultAmbassadorNamespace).Create(t.Context(), &v1.Service{
				Name: defaultAmbassadorServiceName,
				Status: v1.ServiceStatus{
					LoadBalancer: v1.LoadBalancerStatus{Ingress: []v1.LoadBalancerIngress{{IP: "1.1.1.1"}}},
				},
			}, metav1.CreateOptions{})
			require.NoError(t, err)

			ambassadorScheme := runtime.NewScheme()
			require.NoError(t, ambassador.AddToScheme(ambassadorScheme))
			fakeDynamicClient := fakeDynamic.NewSimpleDynamicClient(ambassadorScheme)

			// No spec hostname, so only the fqdn template names the record.
			host, err := createAmbassadorHost(&ambassador.Host{
				Name: "tmpl-host",
				Annotations: map[string]string{
					ambHostAnnotation: defaultAmbassadorNamespace + "/" + defaultAmbassadorServiceName,
				},
				Spec: &ambassador.HostSpec{},
			})
			require.NoError(t, err)
			_, err = fakeDynamicClient.Resource(ambHostGVR).Namespace(v1.NamespaceDefault).Create(t.Context(), host, metav1.CreateOptions{})
			require.NoError(t, err)

			src, err := NewAmbassadorHostSource(t.Context(), fakeDynamicClient, fakeKubernetesClient, &config.Config{
				Namespaces:     []string{v1.NamespaceDefault},
				LabelFilter:    labels.Everything(),
				TemplateEngine: templatetest.MustEngine(t, "{{.Name}}.example.org", tt.targetTemplate, "", false),
			})
			require.NoError(t, err)

			endpoints, err := src.Endpoints(t.Context())
			require.NoError(t, err)
			testutils.ValidateEndpoints(t, endpoints, tt.expected)
		})
	}
}
