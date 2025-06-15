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
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes/fake"
	"sigs.k8s.io/external-dns/endpoint"
)

func TestNodeSourceNewNodeSourceWithFqdn(t *testing.T) {
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
	} {
		t.Run(tt.title, func(t *testing.T) {
			_, err := NewNodeSource(
				t.Context(),
				fake.NewClientset(),
				tt.annotationFilter,
				tt.fqdnTemplate,
				labels.Everything(),
				true,
				true,
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

func TestNodeSourceFqdnTemplatingExamples(t *testing.T) {
	for _, tt := range []struct {
		title        string
		nodes        []*v1.Node
		fqdnTemplate string
		expected     []*endpoint.Endpoint
		combineFQDN  bool
	}{
		{
			title: "templating expansion with multiple domains",
			nodes: []*v1.Node{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "ip-10-1-176-5.internal",
					},
					Status: v1.NodeStatus{
						Addresses: []v1.NodeAddress{
							{Type: v1.NodeInternalIP, Address: "10.1.176.1"},
							{Type: v1.NodeInternalIP, Address: "fc00:f853:ccd:e793::1"},
						},
					},
				},
			},
			fqdnTemplate: "{{.Name}}.domainA.com,{{.Name}}.domainB.com",
			expected: []*endpoint.Endpoint{
				{DNSName: "ip-10-1-176-5.internal.domainA.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"10.1.176.1"}},
				{DNSName: "ip-10-1-176-5.internal.domainA.com", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"fc00:f853:ccd:e793::1"}},
				{DNSName: "ip-10-1-176-5.internal.domainB.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"10.1.176.1"}},
				{DNSName: "ip-10-1-176-5.internal.domainB.com", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"fc00:f853:ccd:e793::1"}},
			},
		},
		{
			title: "templating contains namespace when node namespace is not a valid variable",
			nodes: []*v1.Node{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "node-name",
					},
					Status: v1.NodeStatus{
						Addresses: []v1.NodeAddress{
							{Type: v1.NodeInternalIP, Address: "10.1.176.1"},
						},
					},
				},
			},
			fqdnTemplate: "{{.Name}}.domainA.com,{{ .Name }}.{{ .Namespace }}.example.tld",
			expected: []*endpoint.Endpoint{
				{DNSName: "node-name.domainA.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"10.1.176.1"}},
				{DNSName: "node-name..example.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"10.1.176.1"}},
			},
		},
		{
			title: "templating with external IP and range of addresses",
			nodes: []*v1.Node{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "ip-10-1-176-1",
					},
					Status: v1.NodeStatus{
						Addresses: []v1.NodeAddress{
							{Type: v1.NodeExternalIP, Address: "243.186.136.160"},
							{Type: v1.NodeInternalIP, Address: "fc00:f853:ccd:e793::1"},
						},
					},
				},
			},
			fqdnTemplate: "{{ range .Status.Addresses }}{{if and (eq .Type \"ExternalIP\") (isIPv4 .Address)}}ip-{{ .Address | replace \".\" \"-\" }}{{ break }}{{ end }}{{ end }}.example.com",
			expected: []*endpoint.Endpoint{
				{DNSName: "ip-243-186-136-160.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"243.186.136.160"}},
				{DNSName: "ip-243-186-136-160.example.com", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"fc00:f853:ccd:e793::1"}},
			},
		},
		{
			title: "templating with name definition and ipv4 check",
			nodes: []*v1.Node{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "node-name-ip",
					},
					Status: v1.NodeStatus{
						Addresses: []v1.NodeAddress{
							{Type: v1.NodeExternalIP, Address: "243.186.136.160"},
							{Type: v1.NodeInternalIP, Address: "fc00:f853:ccd:e793::1"},
						},
					},
				},
			},
			fqdnTemplate: "{{ $name := .Name }}{{ range .Status.Addresses }}{{if (isIPv4 .Address)}}{{ $name }}.ipv4{{ break }}{{ end }}{{ end }}.example.com",
			expected: []*endpoint.Endpoint{
				{DNSName: "node-name-ip.ipv4.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"243.186.136.160"}},
				{DNSName: "node-name-ip.ipv4.example.com", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"fc00:f853:ccd:e793::1"}},
			},
		},
		{
			title: "templating with hostname annotation",
			nodes: []*v1.Node{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "ip-10-1-176-1",
						Annotations: map[string]string{
							"external-dns.alpha.kubernetes.io/hostname": "ip-10-1-176-1.internal.domain.com",
						},
					},
					Status: v1.NodeStatus{
						Addresses: []v1.NodeAddress{
							{Type: v1.NodeExternalIP, Address: "243.186.136.160"},
							{Type: v1.NodeInternalIP, Address: "fc00:f853:ccd:e793::1"},
						},
					},
				},
			},
			fqdnTemplate: "{{.Name}}.example.com",
			expected: []*endpoint.Endpoint{
				{DNSName: "ip-10-1-176-1.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"243.186.136.160"}},
				{DNSName: "ip-10-1-176-1.example.com", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"fc00:f853:ccd:e793::1"}},
			},
		},
		{
			title: "templating when target annotation and no external IP",
			nodes: []*v1.Node{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:   "node-name",
						Labels: nil,
						Annotations: map[string]string{
							"external-dns.alpha.kubernetes.io/target": "203.2.45.22",
						},
					},
					Status: v1.NodeStatus{
						Addresses: []v1.NodeAddress{
							{Type: v1.NodeExternalIP, Address: "243.186.136.160"},
							{Type: v1.NodeInternalIP, Address: "fc00:f853:ccd:e793::1"},
						},
					},
				},
			},
			fqdnTemplate: "{{.Name}}.example.com",
			expected: []*endpoint.Endpoint{
				{DNSName: "node-name.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"203.2.45.22"}},
			},
		},
		{
			title: "templating with simple annotation expansion",
			nodes: []*v1.Node{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "node-name",
						Annotations: map[string]string{
							"workload": "cluster-resources",
						},
					},
					Status: v1.NodeStatus{
						Addresses: []v1.NodeAddress{
							{Type: v1.NodeExternalIP, Address: "243.186.136.160"},
						},
					},
				},
			},
			fqdnTemplate: "{{ .Name }}.{{ .Annotations.workload }}.domain.tld",
			expected: []*endpoint.Endpoint{
				{DNSName: "node-name.cluster-resources.domain.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"243.186.136.160"}},
			},
		},
		{
			title: "templating with complex labels expansion",
			nodes: []*v1.Node{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "node-name",
						Labels: map[string]string{
							"topology.kubernetes.io/region": "eu-west-1",
						},
						Annotations: nil,
					},
					Spec: v1.NodeSpec{
						Unschedulable: false,
					},
					Status: v1.NodeStatus{
						Addresses: []v1.NodeAddress{
							{Type: v1.NodeExternalIP, Address: "243.186.136.160"},
						},
					},
				},
			},
			fqdnTemplate: "{{ .Name }}.{{ index .ObjectMeta.Labels \"topology.kubernetes.io/region\" }}.domain.tld",
			expected: []*endpoint.Endpoint{
				{DNSName: "node-name.eu-west-1.domain.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"243.186.136.160"}},
			},
		},
		{
			title: "templating with shared all domain",
			nodes: []*v1.Node{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "node-name-1",
					},
					Status: v1.NodeStatus{
						Addresses: []v1.NodeAddress{
							{Type: v1.NodeExternalIP, Address: "243.186.136.160"},
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "node-name-2",
					},
					Status: v1.NodeStatus{
						Addresses: []v1.NodeAddress{
							{Type: v1.NodeExternalIP, Address: "243.186.136.178"},
						},
					},
				},
			},
			fqdnTemplate: "{{ .Name }}.domain.tld,all.example.com",
			expected: []*endpoint.Endpoint{
				{DNSName: "all.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"243.186.136.160", "243.186.136.178"}},
				{DNSName: "node-name-1.domain.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"243.186.136.160"}},
				{DNSName: "node-name-2.domain.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"243.186.136.178"}},
			},
		},
		{
			title:       "templating with shared all domain and fqdn combination annotation",
			combineFQDN: true,
			nodes: []*v1.Node{
				{
					ObjectMeta: metav1.ObjectMeta{Name: "node-name-1"},
					Status: v1.NodeStatus{
						Addresses: []v1.NodeAddress{{Type: v1.NodeExternalIP, Address: "243.186.136.160"}},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{Name: "node-name-2"},
					Status: v1.NodeStatus{
						Addresses: []v1.NodeAddress{{Type: v1.NodeExternalIP, Address: "243.186.136.178"}},
					},
				},
			},
			fqdnTemplate: "{{ .Name }}.domain.tld,all.example.com",
			expected: []*endpoint.Endpoint{
				{DNSName: "all.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"243.186.136.160", "243.186.136.178"}},
				{DNSName: "node-name-1.domain.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"243.186.136.160"}},
				{DNSName: "node-name-2.domain.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"243.186.136.178"}},
				{DNSName: "node-name-1", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"243.186.136.160"}},
				{DNSName: "node-name-2", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"243.186.136.178"}},
			},
		},
	} {
		t.Run(tt.title, func(t *testing.T) {
			kubeClient := fake.NewClientset()

			for _, node := range tt.nodes {
				_, err := kubeClient.CoreV1().Nodes().Create(t.Context(), node, metav1.CreateOptions{})
				require.NoError(t, err)
			}

			src, err := NewNodeSource(
				t.Context(),
				kubeClient,
				"",
				tt.fqdnTemplate,
				labels.Everything(),
				true,
				true,
				tt.combineFQDN,
			)
			require.NoError(t, err)

			endpoints, err := src.Endpoints(t.Context())
			require.NoError(t, err)

			validateEndpoints(t, endpoints, tt.expected)
		})
	}
}
