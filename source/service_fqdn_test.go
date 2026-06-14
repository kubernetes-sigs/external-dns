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
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes/fake"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	"sigs.k8s.io/external-dns/source/annotations"
	templatetest "sigs.k8s.io/external-dns/source/template/testutil"
)

func TestServiceFQDNTemplate(t *testing.T) {
	const (
		svcName   = "my-svc"
		clusterIP = "10.0.0.1"
	)

	makeSvc := func(anns map[string]string) *v1.Service {
		return &v1.Service{
			ObjectMeta: metav1.ObjectMeta{
				Name:        svcName,
				Namespace:   "default",
				Annotations: anns,
			},
			Spec: v1.ServiceSpec{
				Type:      v1.ServiceTypeClusterIP,
				ClusterIP: clusterIP,
			},
		}
	}

	for _, tt := range []struct {
		title              string
		services           []*v1.Service // nil = [makeSvc(nil)]
		endpointSlices     []*discoveryv1.EndpointSlice
		fqdnTemplate       string
		targetTemplate     string
		fqdnTargetTemplate string
		combine            bool
		publishHostIP      bool
		serviceTypesFilter []string
		expected           []*endpoint.Endpoint
	}{
		// ── fqdn-target-template cases ────────────────────────────────────────
		{
			title:              "fqdn-target-template generates A record when no annotation-derived endpoints",
			fqdnTargetTemplate: "{{.Name}}.example.com:1.2.3.4",
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint(svcName+".example.com", endpoint.RecordTypeA, "1.2.3.4"),
			},
		},
		{
			title:              "fqdn-target-template generates CNAME for hostname target",
			fqdnTargetTemplate: "{{.Name}}.example.com:lb.example.com",
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint(svcName+".example.com", endpoint.RecordTypeCNAME, "lb.example.com"),
			},
		},
		{
			title: "fqdn-target-template with combine adds endpoint alongside annotation-derived",
			services: []*v1.Service{makeSvc(map[string]string{
				annotations.HostnameKey: "annotated.example.com",
			})},
			fqdnTargetTemplate: "{{.Name}}.tmpl.example.com:lb.example.com",
			combine:            true,
			expected: []*endpoint.Endpoint{
				{DNSName: "annotated.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{clusterIP}},
				{DNSName: svcName + ".tmpl.example.com", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"lb.example.com"}},
			},
		},
		{
			title: "fqdn-target-template without combine is ignored when annotation-derived endpoints exist",
			services: []*v1.Service{makeSvc(map[string]string{
				annotations.HostnameKey: "annotated.example.com",
			})},
			fqdnTargetTemplate: "{{.Name}}.tmpl.example.com:lb.example.com",
			combine:            false,
			expected: []*endpoint.Endpoint{
				{DNSName: "annotated.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{clusterIP}},
			},
		},
		{
			title:              "fqdn-target-template can reference .Kind",
			fqdnTargetTemplate: "{{.Kind | toLower}}.{{.Name}}.example.com:1.2.3.4",
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("service."+svcName+".example.com", endpoint.RecordTypeA, "1.2.3.4"),
			},
		},
		{
			title:              "fqdn-target-template pair missing colon is skipped",
			fqdnTargetTemplate: "{{.Name}}.example.com",
		},
		{
			title:        "fqdn-template can reference .Kind",
			fqdnTemplate: "{{.Kind | toLower}}.{{.Name}}.example.com",
			expected: []*endpoint.Endpoint{
				{DNSName: "service." + svcName + ".example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{clusterIP}},
			},
		},
		{
			title:              "fqdn-target-template can reference .APIVersion",
			fqdnTargetTemplate: "{{.Name}}.{{.APIVersion}}.example.com:1.2.3.4",
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint(svcName+".v1.example.com", endpoint.RecordTypeA, "1.2.3.4"),
			},
		},
		// ── fqdn-template cases ───────────────────────────────────────────────
		{
			title:   "fqdn-template with combine across multiple services",
			combine: true,
			services: []*v1.Service{
				{
					ObjectMeta: metav1.ObjectMeta{Namespace: "default", Name: "service-1"},
					Spec:       v1.ServiceSpec{Type: v1.ServiceTypeClusterIP, ClusterIP: "170.19.58.167"},
				},
				{
					ObjectMeta: metav1.ObjectMeta{Namespace: "kube-system", Name: "service-2"},
					Spec:       v1.ServiceSpec{Type: v1.ServiceTypeClusterIP, ClusterIP: "127.20.24.218"},
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
			title:   "fqdn-template with combine alongside internal hostname annotation",
			combine: true,
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
							Ingress: []v1.LoadBalancerIngress{{Hostname: "service-one.example.tld"}},
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
			title: "fqdn-template filters by service type",
			services: []*v1.Service{
				{
					ObjectMeta: metav1.ObjectMeta{Namespace: "default", Name: "service-one"},
					Spec:       v1.ServiceSpec{Type: v1.ServiceTypeLoadBalancer},
					Status: v1.ServiceStatus{LoadBalancer: v1.LoadBalancerStatus{
						Ingress: []v1.LoadBalancerIngress{{Hostname: "service-one.example.tld"}},
					}},
				},
				{
					ObjectMeta: metav1.ObjectMeta{Namespace: "default", Name: "service-two"},
					Spec: v1.ServiceSpec{
						Type:         v1.ServiceTypeExternalName,
						ExternalName: "bucket-name.s3.us-east-1.amazonaws.com",
					},
				},
			},
			fqdnTemplate: `{{ if eq .Spec.Type "ExternalName" }}{{ .Name }}.external.example.tld{{ end}}`,
			expected: []*endpoint.Endpoint{
				{DNSName: "service-two.external.example.tld", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"bucket-name.s3.us-east-1.amazonaws.com"}},
			},
		},
		{
			title: "fqdn-template filters by selector value",
			services: []*v1.Service{
				{
					ObjectMeta: metav1.ObjectMeta{Namespace: "default", Name: "service-one"},
					Spec: v1.ServiceSpec{
						Type: v1.ServiceTypeExternalName, ExternalName: "api.example.tld",
						Selector: map[string]string{"app": "my-app"},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{Namespace: "default", Name: "service-two"},
					Spec: v1.ServiceSpec{
						Type: v1.ServiceTypeExternalName, ExternalName: "www.bucket-name.amazonaws.com",
						Selector: map[string]string{"app": "my-website"},
					},
				},
			},
			fqdnTemplate: `{{ if eq (index .Spec.Selector "app") "my-website" }}www.{{ .Name }}.website.example.tld{{ end}}`,
			expected: []*endpoint.Endpoint{
				{DNSName: "www.service-two.website.example.tld", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"www.bucket-name.amazonaws.com"}},
			},
		},
		{
			title:              "fqdn-template with NodeExternalIP endpoint type and loose service filter",
			serviceTypesFilter: []string{},
			services: []*v1.Service{
				{
					ObjectMeta: metav1.ObjectMeta{
						Namespace:   "svc-ns",
						Name:        "svc-one",
						Annotations: map[string]string{annotations.EndpointsTypeKey: EndpointsTypeNodeExternalIP},
					},
					Spec: v1.ServiceSpec{
						Type: v1.ServiceTypeClusterIP, ClusterIP: v1.ClusterIPNone,
						ClusterIPs: []string{v1.ClusterIPNone},
					},
				},
			},
			endpointSlices: []*discoveryv1.EndpointSlice{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "svc-one-xxxxx", Namespace: "svc-ns",
						Labels: map[string]string{
							discoveryv1.LabelServiceName: "svc-one",
							v1.IsHeadlessService:         "",
						},
					},
					AddressType: discoveryv1.AddressTypeIPv4,
					Endpoints: []discoveryv1.Endpoint{
						{
							Addresses: []string{"100.66.2.246"},
							Hostname:  new("ip-10-1-164-158.internal"),
							NodeName:  new("test-node"),
							TargetRef: &v1.ObjectReference{Kind: "Pod", Name: "pod-1", Namespace: "svc-ns"},
						},
						{
							Addresses: []string{"100.66.2.247"},
							Hostname:  new("ip-10-1-164-158.internal"),
							NodeName:  new("test-node"),
							TargetRef: &v1.ObjectReference{Kind: "Pod", Name: "pod-2", Namespace: "svc-ns"},
						},
					},
				},
			},
			fqdnTemplate: "{{.Name}}.{{.Namespace}}.cluster.com",
			expected: []*endpoint.Endpoint{
				{DNSName: "ip-10-1-164-158.internal.svc-one.svc-ns.cluster.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"203.0.113.10"}},
				{DNSName: "svc-one.svc-ns.cluster.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"203.0.113.10"}},
			},
		},
		{
			title:              "fqdn-template filtered out when service type filter excludes ClusterIP",
			serviceTypesFilter: []string{string(v1.ServiceTypeClusterIP)},
			services: []*v1.Service{
				{
					ObjectMeta: metav1.ObjectMeta{
						Namespace:   "svc-ns",
						Name:        "svc-one",
						Annotations: map[string]string{annotations.EndpointsTypeKey: EndpointsTypeNodeExternalIP},
					},
					Spec: v1.ServiceSpec{
						Type: v1.ServiceTypeClusterIP, ClusterIP: v1.ClusterIPNone,
						ClusterIPs: []string{v1.ClusterIPNone},
					},
				},
			},
			endpointSlices: []*discoveryv1.EndpointSlice{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "svc-one-xxxxx", Namespace: "svc-ns",
						Labels: map[string]string{
							discoveryv1.LabelServiceName: "svc-one",
							v1.IsHeadlessService:         "",
						},
					},
					AddressType: discoveryv1.AddressTypeIPv4,
					Endpoints: []discoveryv1.Endpoint{
						{
							Addresses: []string{"100.66.2.246"},
							Hostname:  new("ip-10-1-164-158.internal"),
							NodeName:  new("test-node"),
							TargetRef: &v1.ObjectReference{Kind: "Pod", Name: "pod-1", Namespace: "svc-ns"},
						},
						{
							Addresses: []string{"100.66.2.247"},
							Hostname:  new("ip-10-1-164-158.internal"),
							NodeName:  new("test-node"),
							TargetRef: &v1.ObjectReference{Kind: "Pod", Name: "pod-2", Namespace: "svc-ns"},
						},
					},
				},
			},
			fqdnTemplate: "{{.Name}}.{{.Namespace}}.cluster.com",
			expected:     []*endpoint.Endpoint{},
		},
		{
			title: "fqdn-template with TrafficDistribution zone annotation",
			services: []*v1.Service{
				{
					ObjectMeta: metav1.ObjectMeta{
						Namespace:   "default",
						Name:        "service-one",
						Annotations: map[string]string{"topology.kubernetes.io/zone": "us-west-1a"},
					},
					Spec: v1.ServiceSpec{
						Type: v1.ServiceTypeClusterIP, ClusterIP: "192.51.100.22",
						ExternalIPs:         []string{"198.51.100.30"},
						TrafficDistribution: new("PreferSameZone"),
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Namespace:   "default",
						Name:        "service-two",
						Annotations: map[string]string{"topology.kubernetes.io/zone": "us-west-1c"},
					},
					Spec: v1.ServiceSpec{
						Type: v1.ServiceTypeClusterIP, ClusterIP: "192.51.100.5",
						ExternalIPs:         []string{"198.51.100.32"},
						TrafficDistribution: new("PreferSameZone"),
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Namespace:   "default",
						Name:        "service-three",
						Annotations: map[string]string{"topology.kubernetes.io/zone": "us-west-1a"},
					},
					Spec: v1.ServiceSpec{
						Type: v1.ServiceTypeClusterIP, ClusterIP: "192.51.100.33",
						ExternalIPs:         []string{"198.51.100.70"},
						TrafficDistribution: new("PreferClose"),
					},
				},
			},
			fqdnTemplate: `{{ $annotations := .ObjectMeta.Annotations }}{{ .Name }}{{ if eq (.Spec.TrafficDistribution | printf) "PreferSameZone" }}.zone.{{ index $annotations "topology.kubernetes.io/zone" }}{{ else }}.close{{ end }}.example.tld`,
			expected: []*endpoint.Endpoint{
				{DNSName: "service-one.zone.us-west-1a.example.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.51.100.22"}},
				{DNSName: "service-two.zone.us-west-1c.example.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.51.100.5"}},
				{DNSName: "service-three.close.example.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.51.100.33"}},
			},
		},
		{
			title: "fqdn-template with specific port names",
			services: []*v1.Service{
				{
					ObjectMeta: metav1.ObjectMeta{Namespace: "default", Name: "service-one"},
					Spec: v1.ServiceSpec{
						Type: v1.ServiceTypeClusterIP, ClusterIP: "192.51.100.22",
						Ports: []v1.ServicePort{{Name: "http", Port: 8080}, {Name: "debug", Port: 8082}},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{Namespace: "default", Name: "service-two"},
					Spec: v1.ServiceSpec{
						Type: v1.ServiceTypeClusterIP, ClusterIP: "192.51.100.5",
						Ports: []v1.ServicePort{{Name: "http", Port: 8080}, {Name: "http2", Port: 8086}},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{Namespace: "default", Name: "service-three"},
					Spec: v1.ServiceSpec{
						Type: v1.ServiceTypeClusterIP, ClusterIP: "2041:0000:140F::875B:131B",
						Ports: []v1.ServicePort{{Name: "debug", Port: 8082}, {Name: "http2", Port: 8086}},
					},
				},
			},
			fqdnTemplate: `{{ $name := .Name }}{{ range .Spec.Ports -}}{{ $name }}{{ if eq .Name "http2" }}.http2{{ else if eq .Name "debug" }}.debug{{ end }}.example.tld.{{printf "," }}{{ end }}`,
			expected: []*endpoint.Endpoint{
				{DNSName: "service-one.debug.example.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.51.100.22"}},
				{DNSName: "service-one.example.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.51.100.22"}},
				{DNSName: "service-three.debug.example.tld", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2041:0000:140F::875B:131B"}},
				{DNSName: "service-three.http2.example.tld", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2041:0000:140F::875B:131B"}},
				{DNSName: "service-two.example.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.51.100.5"}},
				{DNSName: "service-two.http2.example.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.51.100.5"}},
			},
		},
		{
			title: "fqdn-template resolves headless services using pod IPs",
			services: []*v1.Service{
				{
					ObjectMeta: metav1.ObjectMeta{Namespace: "default", Name: "service-one"},
					Spec: v1.ServiceSpec{
						Type: v1.ServiceTypeClusterIP, ClusterIP: v1.ClusterIPNone,
						IPFamilies: []v1.IPFamily{v1.IPv4Protocol},
						Ports:      []v1.ServicePort{{Name: "http", Port: 8080}},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{Namespace: "default", Name: "service-two"},
					Spec: v1.ServiceSpec{
						Type: v1.ServiceTypeClusterIP, ClusterIP: v1.ClusterIPNone,
						IPFamilies: []v1.IPFamily{v1.IPv4Protocol, v1.IPv6Protocol},
						Ports:      []v1.ServicePort{{Name: "http", Port: 8080}},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{Namespace: "default", Name: "service-three"},
					Spec: v1.ServiceSpec{
						Type: v1.ServiceTypeClusterIP, ClusterIP: v1.ClusterIPNone,
						IPFamilies: []v1.IPFamily{v1.IPv4Protocol},
						Ports:      []v1.ServicePort{{Name: "debug", Port: 8082}},
					},
				},
			},
			endpointSlices: []*discoveryv1.EndpointSlice{
				{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "default", Name: "service-one-xxxxx",
						Labels: map[string]string{discoveryv1.LabelServiceName: "service-one"},
					},
					AddressType: discoveryv1.AddressTypeIPv4,
					Endpoints: []discoveryv1.Endpoint{
						{
							Addresses: []string{"100.66.2.241"}, Hostname: new("ip-10-1-164-158.internal"),
							TargetRef: &v1.ObjectReference{Kind: "Pod", Name: "pod-1", Namespace: "default"},
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "default", Name: "service-two-xxxxx",
						Labels: map[string]string{discoveryv1.LabelServiceName: "service-two"},
					},
					AddressType: discoveryv1.AddressTypeIPv4,
					Endpoints: []discoveryv1.Endpoint{
						{
							Addresses: []string{"100.66.2.244"}, Hostname: new("ip-10-1-164-152.internal"),
							TargetRef: &v1.ObjectReference{Kind: "Pod", Name: "pod-2", Namespace: "default"},
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "default", Name: "service-three-xxxxx",
						Labels: map[string]string{discoveryv1.LabelServiceName: "service-three"},
					},
					AddressType: discoveryv1.AddressTypeIPv4,
					Endpoints: []discoveryv1.Endpoint{
						{
							Addresses: []string{"100.66.2.246"}, Hostname: new("ip-10-1-164-158.internal"),
							TargetRef: &v1.ObjectReference{Kind: "Pod", Name: "pod-3", Namespace: "default"},
						},
						{
							Addresses: []string{"100.66.2.247"}, Hostname: new("ip-10-1-164-158.internal"),
							TargetRef: &v1.ObjectReference{Kind: "Pod", Name: "pod-4", Namespace: "default"},
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
			title:         "fqdn-template resolves headless services using host IP",
			publishHostIP: true,
			services: []*v1.Service{
				{
					ObjectMeta: metav1.ObjectMeta{Namespace: "default", Name: "service-one"},
					Spec: v1.ServiceSpec{
						Type: v1.ServiceTypeClusterIP, ClusterIP: v1.ClusterIPNone,
						IPFamilies: []v1.IPFamily{v1.IPv4Protocol},
						Ports:      []v1.ServicePort{{Name: "http", Port: 8080}},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{Namespace: "default", Name: "service-two"},
					Spec: v1.ServiceSpec{
						Type: v1.ServiceTypeClusterIP, ClusterIP: v1.ClusterIPNone,
						IPFamilies: []v1.IPFamily{v1.IPv4Protocol, v1.IPv6Protocol},
						Ports:      []v1.ServicePort{{Name: "http", Port: 8080}},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{Namespace: "default", Name: "service-three"},
					Spec: v1.ServiceSpec{
						Type: v1.ServiceTypeClusterIP, ClusterIP: v1.ClusterIPNone,
						IPFamilies: []v1.IPFamily{v1.IPv4Protocol},
						Ports:      []v1.ServicePort{{Name: "debug", Port: 8082}},
					},
				},
			},
			endpointSlices: []*discoveryv1.EndpointSlice{
				{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "default", Name: "service-one-xxxxx",
						Labels: map[string]string{discoveryv1.LabelServiceName: "service-one"},
					},
					AddressType: discoveryv1.AddressTypeIPv4,
					Endpoints: []discoveryv1.Endpoint{
						{
							Addresses: []string{"100.66.2.241"}, Hostname: new("ip-10-1-164-158.internal"),
							TargetRef: &v1.ObjectReference{Kind: "Pod", Name: "pod-1", Namespace: "default"},
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "default", Name: "service-two-xxxxx",
						Labels: map[string]string{discoveryv1.LabelServiceName: "service-two"},
					},
					AddressType: discoveryv1.AddressTypeIPv4,
					Endpoints: []discoveryv1.Endpoint{
						{
							Addresses: []string{"100.66.2.244"}, Hostname: new("ip-10-1-164-152.internal"),
							TargetRef: &v1.ObjectReference{Kind: "Pod", Name: "pod-2", Namespace: "default"},
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "default", Name: "service-three-xxxxx",
						Labels: map[string]string{discoveryv1.LabelServiceName: "service-three"},
					},
					AddressType: discoveryv1.AddressTypeIPv4,
					Endpoints: []discoveryv1.Endpoint{
						{
							Addresses: []string{"100.66.2.246"}, Hostname: new("ip-10-1-164-158.internal"),
							TargetRef: &v1.ObjectReference{Kind: "Pod", Name: "pod-3", Namespace: "default"},
						},
						{
							Addresses: []string{"100.66.2.247"}, Hostname: new("ip-10-1-164-158.internal"),
							TargetRef: &v1.ObjectReference{Kind: "Pod", Name: "pod-4", Namespace: "default"},
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
		{
			title: "fqdn-template for NodePort services produces SRV records",
			services: []*v1.Service{
				{
					ObjectMeta: metav1.ObjectMeta{Namespace: "default", Name: "service-one"},
					Spec: v1.ServiceSpec{
						Type: v1.ServiceTypeNodePort, ClusterIP: "10.96.41.131",
						Ports: []v1.ServicePort{
							{Name: "http", Port: 80, TargetPort: intstr.FromInt32(8080), NodePort: 30080},
							{Name: "debug", Port: 8082, TargetPort: intstr.FromInt32(8082), NodePort: 30082},
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{Namespace: "default", Name: "service-two"},
					Spec: v1.ServiceSpec{
						Type: v1.ServiceTypeClusterIP, ClusterIP: "10.96.41.132",
						Ports: []v1.ServicePort{{Name: "http", Port: 8080}, {Name: "http2", Port: 8086}},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{Namespace: "default", Name: "service-three"},
					Spec: v1.ServiceSpec{
						Type: v1.ServiceTypeNodePort, ClusterIP: "10.96.41.133",
						Ports: []v1.ServicePort{
							{Name: "debug", Port: 8082, TargetPort: intstr.FromInt32(8083), Protocol: v1.ProtocolUDP, NodePort: 30083},
							{Name: "minecraft", Port: 2525, TargetPort: intstr.FromInt32(25256), NodePort: 25565},
						},
					},
				},
			},
			fqdnTemplate: `{{ if eq .Spec.Type "NodePort" }}{{ range .Spec.Ports }}{{ .Name }}.host.tld{{printf "," }}{{end}}{{ end }}`,
			expected: []*endpoint.Endpoint{
				{DNSName: "_service-one._tcp.debug.host.tld", RecordType: endpoint.RecordTypeSRV, Targets: endpoint.Targets{"0 50 30080 debug.host.tld.", "0 50 30082 debug.host.tld."}},
				{DNSName: "_service-one._tcp.http.host.tld", RecordType: endpoint.RecordTypeSRV, Targets: endpoint.Targets{"0 50 30080 http.host.tld.", "0 50 30082 http.host.tld."}},
				{DNSName: "_service-three._tcp.debug.host.tld", RecordType: endpoint.RecordTypeSRV, Targets: endpoint.Targets{"0 50 25565 debug.host.tld."}},
				{DNSName: "_service-three._tcp.minecraft.host.tld", RecordType: endpoint.RecordTypeSRV, Targets: endpoint.Targets{"0 50 25565 minecraft.host.tld."}},
				{DNSName: "_service-three._udp.debug.host.tld", RecordType: endpoint.RecordTypeSRV, Targets: endpoint.Targets{"0 50 30083 debug.host.tld."}},
				{DNSName: "_service-three._udp.minecraft.host.tld", RecordType: endpoint.RecordTypeSRV, Targets: endpoint.Targets{"0 50 30083 minecraft.host.tld."}},
				{DNSName: "debug.host.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"203.0.113.10"}},
				{DNSName: "http.host.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"203.0.113.10"}},
				{DNSName: "minecraft.host.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"203.0.113.10"}},
			},
		},
		{
			title: "fqdn-template can reference .Kind and label contains",
			fqdnTemplate: `{{ if eq .Kind "Service" }}{{ range $key, $value := .Labels }}
				{{ if and (contains $key "app") (contains $value "my-service-") }}
				{{ $.Name }}.{{ $value }}.example.com,{{ end }}{{ end }}{{ end }}`,
			expected: []*endpoint.Endpoint{
				{DNSName: "service-one.my-service-123.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"100.66.2.241"}},
				{DNSName: "service-two.my-service-345.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"100.66.2.244"}},
			},
			services: []*v1.Service{
				{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "default", Name: "service-one",
						Labels: map[string]string{"app1": "my-service-123"},
					},
					Spec: v1.ServiceSpec{
						Type: v1.ServiceTypeClusterIP, ClusterIP: v1.ClusterIPNone,
						IPFamilies: []v1.IPFamily{v1.IPv4Protocol},
						Ports:      []v1.ServicePort{{Name: "http", Port: 8080}},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "default", Name: "service-two",
						Labels: map[string]string{"app2": "my-service-345"},
					},
					Spec: v1.ServiceSpec{
						Type: v1.ServiceTypeClusterIP, ClusterIP: v1.ClusterIPNone,
						IPFamilies: []v1.IPFamily{v1.IPv4Protocol},
						Ports:      []v1.ServicePort{{Name: "http", Port: 8080}},
					},
				},
			},
			endpointSlices: []*discoveryv1.EndpointSlice{
				{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "default", Name: "service-one-xxxxx",
						Labels: map[string]string{discoveryv1.LabelServiceName: "service-one"},
					},
					AddressType: discoveryv1.AddressTypeIPv4,
					Endpoints: []discoveryv1.Endpoint{
						{
							Addresses: []string{"100.66.2.241"},
							TargetRef: &v1.ObjectReference{Kind: "Pod", Name: "pod-1", Namespace: "default"},
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: "default", Name: "service-two-xxxxx",
						Labels: map[string]string{discoveryv1.LabelServiceName: "service-two"},
					},
					AddressType: discoveryv1.AddressTypeIPv4,
					Endpoints: []discoveryv1.Endpoint{
						{
							Addresses: []string{"100.66.2.244"},
							TargetRef: &v1.ObjectReference{Kind: "Pod", Name: "pod-2", Namespace: "default"},
						},
					},
				},
			},
		},
	} {
		t.Run(tt.title, func(t *testing.T) {
			kubeClient := fake.NewClientset()

			services := tt.services
			if services == nil {
				services = []*v1.Service{makeSvc(nil)}
			}
			for _, svc := range services {
				_, err := kubeClient.CoreV1().Services(svc.Namespace).Create(t.Context(), svc, metav1.CreateOptions{})
				require.NoError(t, err)
			}

			_, err := kubeClient.CoreV1().Nodes().Create(t.Context(), &v1.Node{
				ObjectMeta: metav1.ObjectMeta{Name: "test-node"},
				Status: v1.NodeStatus{
					Addresses: []v1.NodeAddress{
						{Type: v1.NodeExternalIP, Address: "203.0.113.10"},
						{Type: v1.NodeInternalIP, Address: "10.0.0.10"},
					},
				},
			}, metav1.CreateOptions{})
			require.NoError(t, err)

			for _, el := range tt.endpointSlices {
				_, err := kubeClient.DiscoveryV1().EndpointSlices(el.Namespace).Create(t.Context(), el, metav1.CreateOptions{})
				require.NoError(t, err)
				for i, ep := range el.Endpoints {
					hostname := ""
					if ep.Hostname != nil {
						hostname = *ep.Hostname
					}
					_, err = kubeClient.CoreV1().Pods(el.Namespace).Create(t.Context(), &v1.Pod{
						ObjectMeta: metav1.ObjectMeta{Name: ep.TargetRef.Name, Namespace: el.Namespace},
						Spec:       v1.PodSpec{Hostname: hostname, NodeName: "test-node"},
						Status:     v1.PodStatus{HostIP: fmt.Sprintf("10.1.20.4%d", i)},
					}, metav1.CreateOptions{})
					require.NoError(t, err)
				}
			}

			src, err := NewServiceSource(t.Context(), kubeClient, &Config{
				TemplateEngine:                 templatetest.MustEngine(t, tt.fqdnTemplate, tt.targetTemplate, tt.fqdnTargetTemplate, tt.combine),
				PublishHostIP:                  tt.publishHostIP,
				ServiceTypeFilter:              tt.serviceTypesFilter,
				PublishInternal:                true,
				AlwaysPublishNotReadyAddresses: true,
				ExposeInternalIPv6:             true,
				ExcludeUnschedulable:           true,
				LabelFilter:                    labels.Everything(),
			})
			require.NoError(t, err)

			endpoints, err := src.Endpoints(t.Context())
			require.NoError(t, err)

			testutils.ValidateEndpoints(t, endpoints, tt.expected)
		})
	}
}
