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

	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	fakeDynamic "k8s.io/client-go/dynamic/fake"
	fakeKube "k8s.io/client-go/kubernetes/fake"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	"sigs.k8s.io/external-dns/source/config"
	templatetest "sigs.k8s.io/external-dns/source/template/testutil"
)

func TestKongTCPIngressFQDNTemplate(t *testing.T) {
	for _, tt := range []struct {
		title          string
		targetTemplate string
		expected       []*endpoint.Endpoint
	}{
		{
			title: "fqdn-template uses the load balancer status",
			expected: []*endpoint.Endpoint{
				{DNSName: "tmpl-tcp.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"2.2.2.2"}},
			},
		},
		{
			title:          "target-template replaces the load balancer status",
			targetTemplate: "lb.example.net",
			expected: []*endpoint.Endpoint{
				{DNSName: "tmpl-tcp.example.com", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"lb.example.net"}},
			},
		},
	} {
		t.Run(tt.title, func(t *testing.T) {
			t.Parallel()

			scheme := runtime.NewScheme()
			scheme.AddKnownTypes(kongGroupdVersionResource.GroupVersion(), &TCPIngress{}, &TCPIngressList{})
			fakeDynamicClient := fakeDynamic.NewSimpleDynamicClient(scheme)

			// No hostname annotation or rule hosts, so only the fqdn template names the record.
			tcpIngressJSON, err := json.Marshal(TCPIngress{
				TypeMeta:   metav1.TypeMeta{APIVersion: kongGroupdVersionResource.GroupVersion().String(), Kind: "TCPIngress"},
				ObjectMeta: metav1.ObjectMeta{Name: "tmpl-tcp", Namespace: defaultKongNamespace},
				Status: tcpIngressStatus{
					LoadBalancer: corev1.LoadBalancerStatus{Ingress: []corev1.LoadBalancerIngress{{IP: "2.2.2.2"}}},
				},
			})
			require.NoError(t, err)
			tcpIngress := unstructured.Unstructured{}
			require.NoError(t, tcpIngress.UnmarshalJSON(tcpIngressJSON))
			_, err = fakeDynamicClient.Resource(kongGroupdVersionResource).Namespace(defaultKongNamespace).Create(t.Context(), &tcpIngress, metav1.CreateOptions{})
			require.NoError(t, err)

			src, err := NewKongTCPIngressSource(t.Context(), fakeDynamicClient, fakeKube.NewSimpleClientset(), &config.Config{
				Namespaces:     []string{defaultKongNamespace},
				LabelFilter:    labels.Everything(),
				TemplateEngine: templatetest.MustEngine(t, "{{.Name}}.example.com", tt.targetTemplate, "", false),
			})
			require.NoError(t, err)

			endpoints, err := src.Endpoints(t.Context())
			require.NoError(t, err)
			testutils.ValidateEndpoints(t, endpoints, tt.expected)
		})
	}
}
