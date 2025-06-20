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
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	v1 "k8s.io/api/core/v1"
	discoveryv1 "k8s.io/api/discovery/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes/fake"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	"sigs.k8s.io/external-dns/source/annotations"
)

func TestServiceSourceFqdnTemplatingExamples(t *testing.T) {

	for _, tt := range []struct {
		title          string
		services       []*v1.Service
		endpointSlices []*discoveryv1.EndpointSlice
		fqdnTemplate   string
		combineFQDN    bool
		publishHostIp  bool
		expected       []*endpoint.Endpoint
	}{
		{
			title:       "templating with multiple services",
			combineFQDN: true,
			services: []*v1.Service{
				{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "default",
						Name:      "service-1",
					},
					Spec: v1.ServiceSpec{
						Type:      v1.ServiceTypeClusterIP,
						ClusterIP: "170.19.58.167",
					},
					Status: v1.ServiceStatus{},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "kube-system",
						Name:      "service-2",
					},
					Spec: v1.ServiceSpec{
						Type:      v1.ServiceTypeClusterIP,
						ClusterIP: "127.20.24.218",
					},
					Status: v1.ServiceStatus{},
				},
			},
			fqdnTemplate: "{{ .Name }}.{{ .Namespace }}.example.tld, all.example.org",
			expected: []*endpoint.Endpoint{
				{DNSName: "all.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"127.20.24.218", "170.19.58.167"}},
				{DNSName: "service-1.default.example.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"170.19.58.167"}},
				{DNSName: "service-2.kube-system.example.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"127.20.24.218"}},
			},
		},
		{
			title:       "templating resolve service source with internal hostnames",
			combineFQDN: true,
			services: []*v1.Service{
				{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "default",
						Name:      "service-one",
						Annotations: map[string]string{
							annotations.InternalHostnameKey: "service-one.internal.tld,service-one.internal.example.tld",
						},
					},
					Spec: v1.ServiceSpec{
						Type:       v1.ServiceTypeLoadBalancer,
						ClusterIP:  "192.240.240.3",
						ClusterIPs: []string{"192.240.240.3", "192.240.240.4"},
					},
					Status: v1.ServiceStatus{
						LoadBalancer: v1.LoadBalancerStatus{
							Ingress: []v1.LoadBalancerIngress{
								{Hostname: "service-one.example.tld"},
							},
						},
					},
				},
			},
			fqdnTemplate: "{{.Name }}.example.tld",
			expected: []*endpoint.Endpoint{
				{DNSName: "service-one.example.tld", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"service-one.example.tld"}},
				{DNSName: "service-one.internal.example.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.240.240.3"}},
				{DNSName: "service-one.internal.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.240.240.3"}},
			},
		},
		{
			title: "templating resolve service by service type",
			services: []*v1.Service{
				{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "default",
						Name:      "service-one",
					},
					Spec: v1.ServiceSpec{
						Type: v1.ServiceTypeLoadBalancer,
					},
					Status: v1.ServiceStatus{
						LoadBalancer: v1.LoadBalancerStatus{
							Ingress: []v1.LoadBalancerIngress{
								{Hostname: "service-one.example.tld"},
							},
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "default",
						Name:      "service-two",
					},
					Spec: v1.ServiceSpec{
						Type:         v1.ServiceTypeExternalName,
						ExternalName: "bucket-name.s3.us-east-1.amazonaws.com",
					},
				},
			},
			fqdnTemplate: `{{ if eq .Spec.Type "ExternalName" }}{{ .Name }}.external.example.tld{{ end}}`,
			expected: []*endpoint.Endpoint{
				// TODO: This test shows that there is a bug that needs to be fixed in the external-dns logic.
				{DNSName: "", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"service-one.example.tld"}},
				{DNSName: "service-two.external.example.tld", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"bucket-name.s3.us-east-1.amazonaws.com"}},
			},
		},
		{
			title:       "templating resolve service with selector",
			combineFQDN: false,
			services: []*v1.Service{
				{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "default",
						Name:      "service-one",
					},
					Spec: v1.ServiceSpec{
						Type:         v1.ServiceTypeExternalName,
						ExternalName: "api.example.tld",
						Selector: map[string]string{
							"app": "my-app",
						},
					},
					Status: v1.ServiceStatus{},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "default",
						Name:      "service-two",
					},
					Spec: v1.ServiceSpec{
						Type:         v1.ServiceTypeExternalName,
						ExternalName: "www.bucket-name.amazonaws.com",
						Selector: map[string]string{
							"app": "my-website",
						},
					},
				},
			},
			fqdnTemplate: `{{ if eq (index .Spec.Selector "app") "my-website" }}www.{{ .Name }}.website.example.tld{{ end}}`,
			expected: []*endpoint.Endpoint{
				// TODO: This test shows that there is a bug that needs to be fixed in the external-dns logic.
				{DNSName: "", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"api.example.tld"}},
				{DNSName: "www.service-two.website.example.tld", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"www.bucket-name.amazonaws.com"}},
			},
		},
		{
			title: "templating resolve service with zone PreferSameTrafficDistribution and topology.kubernetes.io/zone annotation",
			services: []*v1.Service{
				{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "default",
						Name:      "service-one",
						Annotations: map[string]string{
							"topology.kubernetes.io/zone": "us-west-1a",
						},
					},
					Spec: v1.ServiceSpec{
						Type:        v1.ServiceTypeClusterIP,
						ClusterIP:   "192.51.100.22",
						ExternalIPs: []string{"198.51.100.30"},
						// https://kubernetes.io/docs/reference/networking/virtual-ips/#traffic-distribution
						TrafficDistribution: testutils.ToPtr("PreferSameZone"),
					},
					Status: v1.ServiceStatus{},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "default",
						Name:      "service-two",
						Annotations: map[string]string{
							"topology.kubernetes.io/zone": "us-west-1c",
						},
					},
					Spec: v1.ServiceSpec{
						Type:                v1.ServiceTypeClusterIP,
						ClusterIP:           "192.51.100.5",
						ExternalIPs:         []string{"198.51.100.32"},
						TrafficDistribution: testutils.ToPtr("PreferSameZone"),
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "default",
						Name:      "service-three",
						Annotations: map[string]string{
							"topology.kubernetes.io/zone": "us-west-1a",
						},
					},
					Spec: v1.ServiceSpec{
						Type:                v1.ServiceTypeClusterIP,
						ClusterIP:           "192.51.100.33",
						ExternalIPs:         []string{"198.51.100.70"},
						TrafficDistribution: testutils.ToPtr("PreferClose"),
					},
				},
			},
			// printf is used to ensure the template is evaluated as a string, as the TrafficDistribution field is a pointer.
			fqdnTemplate: `{{ $annotations := .ObjectMeta.Annotations }}{{ .Name }}{{ if eq (.Spec.TrafficDistribution | printf) "PreferSameZone" }}.zone.{{ index $annotations "topology.kubernetes.io/zone"  }}{{ else }}.close{{ end }}.example.tld`,
			expected: []*endpoint.Endpoint{
				{DNSName: "service-one.zone.us-west-1a.example.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.51.100.22"}},
				{DNSName: "service-two.zone.us-west-1c.example.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.51.100.5"}},
				{DNSName: "service-three.close.example.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.51.100.33"}},
			},
		},
		{
			title: "templating resolve services with specific port names",
			services: []*v1.Service{
				{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "default",
						Name:      "service-one",
					},
					Spec: v1.ServiceSpec{
						Type:      v1.ServiceTypeClusterIP,
						ClusterIP: "192.51.100.22",
						Ports: []v1.ServicePort{
							{Name: "http", Port: 8080},
							{Name: "debug", Port: 8082},
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "default",
						Name:      "service-two",
					},
					Spec: v1.ServiceSpec{
						Type:      v1.ServiceTypeClusterIP,
						ClusterIP: "192.51.100.5",
						Ports: []v1.ServicePort{
							{Name: "http", Port: 8080},
							{Name: "http2", Port: 8086},
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "default",
						Name:      "service-three",
					},
					Spec: v1.ServiceSpec{
						Type:      v1.ServiceTypeClusterIP,
						ClusterIP: "2041:0000:140F::875B:131B",
						Ports: []v1.ServicePort{
							{Name: "debug", Port: 8082},
							{Name: "http2", Port: 8086},
						},
					},
				},
			},
			fqdnTemplate: `{{ $name := .Name }}{{ range .Spec.Ports -}}{{ $name }}{{ if eq .Name "http2" }}.http2{{ else if eq .Name "debug" }}.debug{{ end }}.example.tld{{printf "," }}{{ end }}`,
			expected: []*endpoint.Endpoint{
				// TODO: This test shows that there is a bug that needs to be fixed in the external-dns logic.
				{DNSName: "", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.51.100.22", "192.51.100.5"}},
				{DNSName: "", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2041:0000:140F::875B:131B"}},
				{DNSName: "service-one.debug.example.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.51.100.22"}},
				{DNSName: "service-one.example.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.51.100.22"}},
				{DNSName: "service-three.debug.example.tld", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2041:0000:140F::875B:131B"}},
				{DNSName: "service-three.http2.example.tld", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2041:0000:140F::875B:131B"}},
				{DNSName: "service-two.example.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.51.100.5"}},
				{DNSName: "service-two.http2.example.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.51.100.5"}},
			},
		},
		{
			title:         "templating resolves headless services",
			publishHostIp: false,
			services: []*v1.Service{
				{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "default",
						Name:      "service-one",
					},
					Spec: v1.ServiceSpec{
						Type:       v1.ServiceTypeClusterIP,
						ClusterIP:  v1.ClusterIPNone,
						IPFamilies: []v1.IPFamily{v1.IPv4Protocol},
						Ports: []v1.ServicePort{
							{Name: "http", Port: 8080},
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "default",
						Name:      "service-two",
					},
					Spec: v1.ServiceSpec{
						Type:       v1.ServiceTypeClusterIP,
						ClusterIP:  v1.ClusterIPNone,
						IPFamilies: []v1.IPFamily{v1.IPv4Protocol, v1.IPv6Protocol},
						Ports: []v1.ServicePort{
							{Name: "http", Port: 8080},
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "default",
						Name:      "service-three",
					},
					Spec: v1.ServiceSpec{
						Type:       v1.ServiceTypeClusterIP,
						ClusterIP:  v1.ClusterIPNone,
						IPFamilies: []v1.IPFamily{v1.IPv4Protocol},
						Ports: []v1.ServicePort{
							{Name: "debug", Port: 8082},
						},
					},
				},
			},
			endpointSlices: []*discoveryv1.EndpointSlice{
				{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "default",
						Name:      "service-one-xxxxx",
						Labels: map[string]string{
							discoveryv1.LabelServiceName: "service-one",
						},
					},
					AddressType: discoveryv1.AddressTypeIPv4,
					Endpoints: []discoveryv1.Endpoint{
						{
							Addresses: []string{"100.66.2.241"},
							Hostname:  testutils.ToPtr("ip-10-1-164-158.internal"),
							TargetRef: &v1.ObjectReference{
								Kind:      "Pod",
								Name:      "pod-1",
								Namespace: "default",
							},
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "default",
						Name:      "service-two-xxxxx",
						Labels: map[string]string{
							discoveryv1.LabelServiceName: "service-two",
						},
					},
					AddressType: discoveryv1.AddressTypeIPv4,
					Endpoints: []discoveryv1.Endpoint{
						{
							Addresses: []string{"100.66.2.244"},
							Hostname:  testutils.ToPtr("ip-10-1-164-152.internal"),
							TargetRef: &v1.ObjectReference{
								Kind:      "Pod",
								Name:      "pod-2",
								Namespace: "default",
							},
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "default",
						Name:      "service-three-xxxxx",
						Labels: map[string]string{
							discoveryv1.LabelServiceName: "service-three",
						},
					},
					AddressType: discoveryv1.AddressTypeIPv4,
					Endpoints: []discoveryv1.Endpoint{
						{
							Addresses: []string{"100.66.2.246"},
							Hostname:  testutils.ToPtr("ip-10-1-164-158.internal"),
							TargetRef: &v1.ObjectReference{
								Kind:      "Pod",
								Name:      "pod-3",
								Namespace: "default",
							},
						},
						{
							Addresses: []string{"100.66.2.247"},
							Hostname:  testutils.ToPtr("ip-10-1-164-158.internal"),
							TargetRef: &v1.ObjectReference{
								Kind:      "Pod",
								Name:      "pod-4",
								Namespace: "default",
							},
						},
					},
				},
			},
			fqdnTemplate: `{{ .Name }}.org.tld`,
			expected: []*endpoint.Endpoint{
				{DNSName: "ip-10-1-164-152.internal.service-two.org.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"100.66.2.244"}},
				{DNSName: "ip-10-1-164-158.internal.service-one.org.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"100.66.2.241"}},
				{DNSName: "ip-10-1-164-158.internal.service-three.org.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"100.66.2.246", "100.66.2.247"}},
				{DNSName: "service-one.org.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"100.66.2.241"}},
				{DNSName: "service-three.org.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"100.66.2.246", "100.66.2.247"}},
				{DNSName: "service-two.org.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"100.66.2.244"}},
			},
		},
		{
			title:         "templating resolves headless services with publishHostIp set to true",
			publishHostIp: true,
			services: []*v1.Service{
				{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "default",
						Name:      "service-one",
					},
					Spec: v1.ServiceSpec{
						Type:       v1.ServiceTypeClusterIP,
						ClusterIP:  v1.ClusterIPNone,
						IPFamilies: []v1.IPFamily{v1.IPv4Protocol},
						Ports: []v1.ServicePort{
							{Name: "http", Port: 8080},
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "default",
						Name:      "service-two",
					},
					Spec: v1.ServiceSpec{
						Type:       v1.ServiceTypeClusterIP,
						ClusterIP:  v1.ClusterIPNone,
						IPFamilies: []v1.IPFamily{v1.IPv4Protocol, v1.IPv6Protocol},
						Ports: []v1.ServicePort{
							{Name: "http", Port: 8080},
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "default",
						Name:      "service-three",
					},
					Spec: v1.ServiceSpec{
						Type:       v1.ServiceTypeClusterIP,
						ClusterIP:  v1.ClusterIPNone,
						IPFamilies: []v1.IPFamily{v1.IPv4Protocol},
						Ports: []v1.ServicePort{
							{Name: "debug", Port: 8082},
						},
					},
				},
			},
			endpointSlices: []*discoveryv1.EndpointSlice{
				{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "default",
						Name:      "service-one-xxxxx",
						Labels: map[string]string{
							discoveryv1.LabelServiceName: "service-one",
						},
					},
					AddressType: discoveryv1.AddressTypeIPv4,
					Endpoints: []discoveryv1.Endpoint{
						{
							Addresses: []string{"100.66.2.241"},
							Hostname:  testutils.ToPtr("ip-10-1-164-158.internal"),
							TargetRef: &v1.ObjectReference{
								Kind:      "Pod",
								Name:      "pod-1",
								Namespace: "default",
							},
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "default",
						Name:      "service-two-xxxxx",
						Labels: map[string]string{
							discoveryv1.LabelServiceName: "service-two",
						},
					},
					AddressType: discoveryv1.AddressTypeIPv4,
					Endpoints: []discoveryv1.Endpoint{
						{
							Addresses: []string{"100.66.2.244"},
							Hostname:  testutils.ToPtr("ip-10-1-164-152.internal"),
							TargetRef: &v1.ObjectReference{
								Kind:      "Pod",
								Name:      "pod-2",
								Namespace: "default",
							},
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "default",
						Name:      "service-three-xxxxx",
						Labels: map[string]string{
							discoveryv1.LabelServiceName: "service-three",
						},
					},
					AddressType: discoveryv1.AddressTypeIPv4,
					Endpoints: []discoveryv1.Endpoint{
						{
							Addresses: []string{"100.66.2.246"},
							Hostname:  testutils.ToPtr("ip-10-1-164-158.internal"),
							TargetRef: &v1.ObjectReference{
								Kind:      "Pod",
								Name:      "pod-3",
								Namespace: "default",
							},
						},
						{
							Addresses: []string{"100.66.2.247"},
							Hostname:  testutils.ToPtr("ip-10-1-164-158.internal"),
							TargetRef: &v1.ObjectReference{
								Kind:      "Pod",
								Name:      "pod-4",
								Namespace: "default",
							},
						},
					},
				},
			},
			fqdnTemplate: `{{ .Name }}.org.tld`,
			expected: []*endpoint.Endpoint{
				{DNSName: "ip-10-1-164-152.internal.service-two.org.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"10.1.20.40"}},
				{DNSName: "ip-10-1-164-158.internal.service-one.org.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"10.1.20.40"}},
				{DNSName: "ip-10-1-164-158.internal.service-three.org.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"10.1.20.40", "10.1.20.41"}},
				{DNSName: "service-one.org.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"10.1.20.40"}},
				{DNSName: "service-three.org.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"10.1.20.40", "10.1.20.41"}},
				{DNSName: "service-two.org.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"10.1.20.40"}},
			},
		},
	} {
		t.Run(tt.title, func(t *testing.T) {
			kubeClient := fake.NewClientset()

			for _, el := range tt.services {
				_, err := kubeClient.CoreV1().Services(el.Namespace).Create(t.Context(), el, metav1.CreateOptions{})
				require.NoError(t, err)
			}

			// Create endpoints and pods for the services
			for _, el := range tt.endpointSlices {
				_, err := kubeClient.DiscoveryV1().EndpointSlices(el.Namespace).Create(t.Context(), el, metav1.CreateOptions{})
				require.NoError(t, err)
				for i, ep := range el.Endpoints {
					_, err = kubeClient.CoreV1().Pods(el.Namespace).Create(t.Context(), &v1.Pod{
						ObjectMeta: metav1.ObjectMeta{
							Name:      ep.TargetRef.Name,
							Namespace: el.Namespace,
						},
						Spec: v1.PodSpec{
							Hostname: *ep.Hostname,
						},
						Status: v1.PodStatus{
							HostIP: fmt.Sprintf("10.1.20.4%d", i),
						},
					}, metav1.CreateOptions{})
					require.NoError(t, err)
				}
			}

			src, err := NewServiceSource(
				t.Context(),
				kubeClient,
				"",
				"",
				tt.fqdnTemplate,
				tt.combineFQDN,
				"",
				true,
				tt.publishHostIp,
				true,
				[]string{},
				false,
				labels.Everything(),
				false,
				false,
				true,
			)
			require.NoError(t, err)

			endpoints, err := src.Endpoints(t.Context())
			require.NoError(t, err)

			validateEndpoints(t, endpoints, tt.expected)

			// TODO; when all resources have the resource label, we could add this check to the validateEndpoints function.
			for _, ep := range endpoints {
				require.Contains(t, ep.Labels, endpoint.ResourceLabelKey)
			}
		})
	}
}
