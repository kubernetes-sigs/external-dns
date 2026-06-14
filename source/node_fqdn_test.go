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

	"sigs.k8s.io/external-dns/internal/testutils"

	"github.com/stretchr/testify/require"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes/fake"

	"sigs.k8s.io/external-dns/endpoint"
	templatetest "sigs.k8s.io/external-dns/source/template/testutil"
)

func TestNodeFQDNTemplate(t *testing.T) {
	const (
		nodeName  = "my-node"
		nodeExtIP = "10.0.0.1"
	)

	makeNode := func() *v1.Node {
		return &v1.Node{
			ObjectMeta: metav1.ObjectMeta{Name: nodeName},
			Status: v1.NodeStatus{
				Addresses: []v1.NodeAddress{
					{Type: v1.NodeExternalIP, Address: nodeExtIP},
				},
			},
		}
	}

	for _, tt := range []struct {
		title              string
		nodes              []*v1.Node // nil = use makeNode()
		fqdnTemplate       string
		targetTemplate     string
		fqdnTargetTemplate string
		combine            bool
		expected           []*endpoint.Endpoint
	}{
		{
			title:              "fqdn-target-template generates A record when no other endpoints",
			fqdnTargetTemplate: "{{.Name}}.example.com:1.2.3.4",
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint(nodeName+".example.com", endpoint.RecordTypeA, "1.2.3.4"),
			},
		},
		{
			title:              "fqdn-target-template generates CNAME for hostname target",
			fqdnTargetTemplate: "{{.Name}}.example.com:lb.example.com",
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint(nodeName+".example.com", endpoint.RecordTypeCNAME, "lb.example.com"),
			},
		},
		{
			title:              "fqdn-target-template without combine replaces default node-name endpoint",
			fqdnTargetTemplate: "{{.Name}}.tmpl.example.com:lb.example.com",
			combine:            false,
			expected: []*endpoint.Endpoint{
				{DNSName: nodeName + ".tmpl.example.com", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"lb.example.com"}},
			},
		},
		{
			title:              "fqdn-target-template with combine adds endpoint alongside default node-name endpoint",
			fqdnTargetTemplate: "{{.Name}}.tmpl.example.com:lb.example.com",
			combine:            true,
			expected: []*endpoint.Endpoint{
				{DNSName: nodeName, RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{nodeExtIP}},
				{DNSName: nodeName + ".tmpl.example.com", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"lb.example.com"}},
			},
		},
		{
			title:              "fqdn-target-template can reference .Kind",
			fqdnTargetTemplate: "{{.Kind | toLower}}.{{.Name}}.example.com:1.2.3.4",
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("node."+nodeName+".example.com", endpoint.RecordTypeA, "1.2.3.4"),
			},
		},
		{
			title:              "fqdn-target-template can reference .APIVersion",
			fqdnTargetTemplate: "{{.Name}}.{{.APIVersion}}.example.com:1.2.3.4",
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint(nodeName+".v1.example.com", endpoint.RecordTypeA, "1.2.3.4"),
			},
		},
		{
			title:        "fqdn-template can reference .Kind",
			fqdnTemplate: "{{.Kind | toLower}}.{{.Name}}.example.com",
			expected: []*endpoint.Endpoint{
				{DNSName: "node." + nodeName + ".example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{nodeExtIP}},
			},
		},
		{
			title:          "target-template alone still generates default node-name endpoints",
			targetTemplate: "lb.example.com",
			expected: []*endpoint.Endpoint{
				{DNSName: nodeName, RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{nodeExtIP}},
			},
		},
		{
			title:        "fqdn-template expansion with multiple domains",
			fqdnTemplate: "{{.Name}}.domainA.com,{{.Name}}.domainB.com",
			nodes: []*v1.Node{
				{
					ObjectMeta: metav1.ObjectMeta{Name: "ip-10-1-176-5.internal"},
					Status: v1.NodeStatus{
						Addresses: []v1.NodeAddress{
							{Type: v1.NodeInternalIP, Address: "10.1.176.1"},
							{Type: v1.NodeInternalIP, Address: "fc00:f853:ccd:e793::1"},
						},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "ip-10-1-176-5.internal.domainA.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"10.1.176.1"}},
				{DNSName: "ip-10-1-176-5.internal.domainA.com", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"fc00:f853:ccd:e793::1"}},
				{DNSName: "ip-10-1-176-5.internal.domainB.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"10.1.176.1"}},
				{DNSName: "ip-10-1-176-5.internal.domainB.com", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"fc00:f853:ccd:e793::1"}},
			},
		},
		{
			title:        "fqdn-template with empty namespace produces double-dot",
			fqdnTemplate: "{{.Name}}.domainA.com,{{ .Name }}.{{ .Namespace }}.example.tld",
			nodes: []*v1.Node{
				{
					ObjectMeta: metav1.ObjectMeta{Name: "node-name"},
					Status: v1.NodeStatus{
						Addresses: []v1.NodeAddress{
							{Type: v1.NodeInternalIP, Address: "10.1.176.1"},
						},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "node-name.domainA.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"10.1.176.1"}},
				{DNSName: "node-name..example.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"10.1.176.1"}},
			},
		},
		{
			title:        "fqdn-template derived from external IP range",
			fqdnTemplate: "{{ range .Status.Addresses }}{{if and (eq .Type \"ExternalIP\") (isIPv4 .Address)}}ip-{{ .Address | replace \".\" \"-\" }}{{ break }}{{ end }}{{ end }}.example.com",
			nodes: []*v1.Node{
				{
					ObjectMeta: metav1.ObjectMeta{Name: "ip-10-1-176-1"},
					Status: v1.NodeStatus{
						Addresses: []v1.NodeAddress{
							{Type: v1.NodeExternalIP, Address: "243.186.136.160"},
							{Type: v1.NodeInternalIP, Address: "fc00:f853:ccd:e793::1"},
						},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "ip-243-186-136-160.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"243.186.136.160"}},
				{DNSName: "ip-243-186-136-160.example.com", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"fc00:f853:ccd:e793::1"}},
			},
		},
		{
			title:        "fqdn-template with IPv4 address check",
			fqdnTemplate: "{{ $name := .Name }}{{ range .Status.Addresses }}{{if (isIPv4 .Address)}}{{ $name }}.ipv4{{ break }}{{ end }}{{ end }}.example.com",
			nodes: []*v1.Node{
				{
					ObjectMeta: metav1.ObjectMeta{Name: "node-name-ip"},
					Status: v1.NodeStatus{
						Addresses: []v1.NodeAddress{
							{Type: v1.NodeExternalIP, Address: "243.186.136.160"},
							{Type: v1.NodeInternalIP, Address: "fc00:f853:ccd:e793::1"},
						},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "node-name-ip.ipv4.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"243.186.136.160"}},
				{DNSName: "node-name-ip.ipv4.example.com", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"fc00:f853:ccd:e793::1"}},
			},
		},
		{
			title:        "fqdn-template overrides hostname annotation",
			fqdnTemplate: "{{.Name}}.example.com",
			nodes: []*v1.Node{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "ip-10-1-176-1",
						Annotations: map[string]string{
							"external-dns.kubernetes.io/hostname": "ip-10-1-176-1.internal.domain.com",
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
			expected: []*endpoint.Endpoint{
				{DNSName: "ip-10-1-176-1.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"243.186.136.160"}},
				{DNSName: "ip-10-1-176-1.example.com", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"fc00:f853:ccd:e793::1"}},
			},
		},
		{
			title:        "fqdn-template uses target annotation instead of node addresses",
			fqdnTemplate: "{{.Name}}.example.com",
			nodes: []*v1.Node{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "node-name",
						Annotations: map[string]string{
							"external-dns.kubernetes.io/target": "203.2.45.22",
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
			expected: []*endpoint.Endpoint{
				{DNSName: "node-name.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"203.2.45.22"}},
			},
		},
		{
			title:        "fqdn-template expands annotation value",
			fqdnTemplate: "{{ .Name }}.{{ .Annotations.workload }}.domain.tld",
			nodes: []*v1.Node{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:        "node-name",
						Annotations: map[string]string{"workload": "cluster-resources"},
					},
					Status: v1.NodeStatus{
						Addresses: []v1.NodeAddress{
							{Type: v1.NodeExternalIP, Address: "243.186.136.160"},
						},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "node-name.cluster-resources.domain.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"243.186.136.160"}},
			},
		},
		{
			title:        "fqdn-template expands label value",
			fqdnTemplate: `{{ .Name }}.{{ index .ObjectMeta.Labels "topology.kubernetes.io/region" }}.domain.tld`,
			nodes: []*v1.Node{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:   "node-name",
						Labels: map[string]string{"topology.kubernetes.io/region": "eu-west-1"},
					},
					Status: v1.NodeStatus{
						Addresses: []v1.NodeAddress{
							{Type: v1.NodeExternalIP, Address: "243.186.136.160"},
						},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "node-name.eu-west-1.domain.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"243.186.136.160"}},
			},
		},
		{
			title:        "fqdn-template shared domain across multiple nodes",
			fqdnTemplate: "{{ .Name }}.domain.tld,all.example.com",
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
			expected: []*endpoint.Endpoint{
				{DNSName: "all.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"243.186.136.160", "243.186.136.178"}},
				{DNSName: "node-name-1.domain.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"243.186.136.160"}},
				{DNSName: "node-name-2.domain.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"243.186.136.178"}},
			},
		},
		{
			title:        "fqdn-template with combine adds template endpoints alongside default node-name endpoints",
			fqdnTemplate: "{{ .Name }}.domain.tld,all.example.com",
			combine:      true,
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
			expected: []*endpoint.Endpoint{
				{DNSName: "all.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"243.186.136.160", "243.186.136.178"}},
				{DNSName: "node-name-1.domain.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"243.186.136.160"}},
				{DNSName: "node-name-2.domain.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"243.186.136.178"}},
				{DNSName: "node-name-1", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"243.186.136.160"}},
				{DNSName: "node-name-2", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"243.186.136.178"}},
			},
		},
		{
			title:        "fqdn-template can reference .Kind and .Status fields",
			fqdnTemplate: `{{ if eq .Kind "Node" }}{{.Name}}.{{.Status.NodeInfo.Architecture}}.node.tld{{ end }}`,
			nodes: []*v1.Node{
				{
					ObjectMeta: metav1.ObjectMeta{Name: "node-name-1"},
					Status: v1.NodeStatus{
						Addresses: []v1.NodeAddress{{Type: v1.NodeExternalIP, Address: "10.0.0.1"}},
						NodeInfo:  v1.NodeSystemInfo{Architecture: "arm64"},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{Name: "node-name-2"},
					Status: v1.NodeStatus{
						Addresses: []v1.NodeAddress{{Type: v1.NodeExternalIP, Address: "10.0.0.2"}},
						NodeInfo:  v1.NodeSystemInfo{Architecture: "x86_64"},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "node-name-1.arm64.node.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"10.0.0.1"}},
				{DNSName: "node-name-2.x86_64.node.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"10.0.0.2"}},
			},
		},
	} {
		t.Run(tt.title, func(t *testing.T) {
			kubeClient := fake.NewClientset()

			nodes := tt.nodes
			if nodes == nil {
				nodes = []*v1.Node{makeNode()}
			}
			for _, node := range nodes {
				_, err := kubeClient.CoreV1().Nodes().Create(t.Context(), node, metav1.CreateOptions{})
				require.NoError(t, err)
			}

			src, err := NewNodeSource(t.Context(), kubeClient, &Config{
				TemplateEngine:       templatetest.MustEngine(t, tt.fqdnTemplate, tt.targetTemplate, tt.fqdnTargetTemplate, tt.combine),
				ExcludeUnschedulable: true,
				ExposeInternalIPv6:   true,
				LabelFilter:          labels.Everything(),
			})
			require.NoError(t, err)

			endpoints, err := src.Endpoints(t.Context())
			require.NoError(t, err)
			testutils.ValidateEndpoints(t, endpoints, tt.expected)
		})
	}
}
