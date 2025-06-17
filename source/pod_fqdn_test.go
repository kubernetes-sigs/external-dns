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
	"k8s.io/client-go/kubernetes/fake"
	"sigs.k8s.io/external-dns/endpoint"
)

func TestNewPodSourceWithFqdn(t *testing.T) {
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
	} {
		t.Run(tt.title, func(t *testing.T) {
			_, err := NewPodSource(
				t.Context(),
				fake.NewClientset(),
				"",
				"",
				false,
				"",
				tt.fqdnTemplate,
				false)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestPodSourceFqdnTemplatingExamples(t *testing.T) {
	for _, tt := range []struct {
		title        string
		pods         []*v1.Pod
		nodes        []*v1.Node
		fqdnTemplate string
		expected     []*endpoint.Endpoint
		combineFQDN  bool
		sourceDomain string
	}{
		{
			title: "templating expansion with multiple domains",
			pods: []*v1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-pod-1",
						Namespace: "default",
					},
					Status: v1.PodStatus{
						PodIP: "100.67.94.101",
						PodIPs: []v1.PodIP{
							{IP: "100.67.94.101"},
						},
					},
				},
			},
			fqdnTemplate: "{{ .Name }}.domainA.com,{{ .Name }}.domainB.com",
			expected: []*endpoint.Endpoint{
				{DNSName: "my-pod-1.domainA.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"100.67.94.101"}},
				{DNSName: "my-pod-1.domainB.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"100.67.94.101"}},
			},
		},
		{
			title:        "templating expansion with multiple domains and fqdn combine and pod source domain",
			combineFQDN:  true,
			sourceDomain: "example.org",
			pods: []*v1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-pod-1",
						Namespace: "default",
					},
					Spec: v1.PodSpec{
						NodeName: "node-1.internal",
					},
					Status: v1.PodStatus{
						PodIP: "100.67.94.101",
						PodIPs: []v1.PodIP{
							{IP: "100.67.94.101"},
						},
					},
				},
			},
			nodes: []*v1.Node{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "node-1.internal",
					},
					Status: v1.NodeStatus{
						Addresses: []v1.NodeAddress{
							{Type: v1.NodeExternalIP, Address: "10.1.192.139"},
						},
					},
				},
			},
			fqdnTemplate: "{{ .Name }}.domainA.com,{{ .Name }}.domainB.com",
			expected: []*endpoint.Endpoint{
				{DNSName: "my-pod-1.domainA.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"100.67.94.101"}},
				{DNSName: "my-pod-1.domainB.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"100.67.94.101"}},
				{DNSName: "my-pod-1.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"100.67.94.101"}},
			},
		},
		{
			title: "templating with domain per namespace",
			pods: []*v1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "pod-1",
						Namespace: "default",
					},
					Status: v1.PodStatus{
						PodIP: "100.67.94.101",
						PodIPs: []v1.PodIP{
							{IP: "100.67.94.101"},
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "pod-2",
						Namespace: "kube-system",
					},
					Status: v1.PodStatus{
						PodIP: "100.67.94.102",
						PodIPs: []v1.PodIP{
							{IP: "100.67.94.102"},
						},
					},
				},
			},
			fqdnTemplate: "{{ .Name }}.{{ .Namespace }}.example.org",
			expected: []*endpoint.Endpoint{
				{DNSName: "pod-1.default.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"100.67.94.101"}},
				{DNSName: "pod-2.kube-system.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"100.67.94.102"}},
			},
		},
		{
			title: "templating with pod and multiple ips for types A and AAAA",
			pods: []*v1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "pod-1",
						Namespace: "default",
					},
					Status: v1.PodStatus{
						PodIP: "100.67.94.101",
						PodIPs: []v1.PodIP{
							{IP: "100.67.94.101"},
							{IP: "2041:0000:140F::875B:131B"},
						},
					},
				},
			},
			fqdnTemplate: "{{ .Name }}.example.org",
			expected: []*endpoint.Endpoint{
				{DNSName: "pod-1.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"100.67.94.101"}},
				{DNSName: "pod-1.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2041:0000:140F::875B:131B"}},
			},
		},
		{
			title: "templating with pod and target annotation that is currently not overriding target IPs",
			pods: []*v1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "pod-1",
						Namespace: "default",
						Annotations: map[string]string{
							"external-dns.alpha.kubernetes.io/target": "203.2.45.22",
						},
					},
					Status: v1.PodStatus{
						PodIP: "100.67.94.101",
						PodIPs: []v1.PodIP{
							{IP: "100.67.94.101"},
						},
					},
				},
			},
			fqdnTemplate: "{{ .Name }}.example.org",
			expected: []*endpoint.Endpoint{
				{DNSName: "pod-1.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"100.67.94.101"}},
			},
		},
		{
			title: "templating with pod and host annotation that is currently not overriding hostname",
			pods: []*v1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "pod-1",
						Namespace: "default",
						Annotations: map[string]string{
							"external-dns.alpha.kubernetes.io/hostname": "ip-10-1-176-1.internal.domain.com",
						},
					},
					Status: v1.PodStatus{
						PodIP: "100.67.94.101",
						PodIPs: []v1.PodIP{
							{IP: "100.67.94.101"},
						},
					},
				},
			},
			fqdnTemplate: "{{ .Name }}.example.org",
			expected: []*endpoint.Endpoint{
				{DNSName: "pod-1.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"100.67.94.101"}},
			},
		},
		{
			title: "templating with simple annotation expansion",
			pods: []*v1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "pod-1",
						Namespace: "kube-system",
						Annotations: map[string]string{
							"workload": "cluster-resources",
						},
					},
					Status: v1.PodStatus{
						PodIP: "100.67.94.101",
						PodIPs: []v1.PodIP{
							{IP: "100.67.94.101"},
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "pod-2",
						Namespace: "workloads",
						Annotations: map[string]string{
							"workload": "workloads",
						},
					},
					Status: v1.PodStatus{
						PodIP: "100.67.94.102",
						PodIPs: []v1.PodIP{
							{IP: "100.67.94.102"},
						},
					},
				},
			},
			fqdnTemplate: "{{ .Name }}.{{ .Annotations.workload }}.domain.tld",
			expected: []*endpoint.Endpoint{
				{DNSName: "pod-1.cluster-resources.domain.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"100.67.94.101"}},
				{DNSName: "pod-2.workloads.domain.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"100.67.94.102"}},
			},
		},
		{
			title: "templating with complex label expansion",
			pods: []*v1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "pod-1",
						Namespace: "kube-system",
						Labels: map[string]string{
							"topology.kubernetes.io/region": "eu-west-1a",
						},
					},
					Status: v1.PodStatus{
						PodIP: "100.67.94.101",
						PodIPs: []v1.PodIP{
							{IP: "100.67.94.101"},
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "pod-2",
						Namespace: "workloads",
						Labels: map[string]string{
							"topology.kubernetes.io/region": "eu-west-1b",
						},
					},
					Status: v1.PodStatus{
						PodIP: "100.67.94.102",
						PodIPs: []v1.PodIP{
							{IP: "100.67.94.102"},
						},
					},
				},
			},
			fqdnTemplate: "{{ .Name }}.{{ index .ObjectMeta.Labels \"topology.kubernetes.io/region\" }}.domain.tld",
			expected: []*endpoint.Endpoint{
				{DNSName: "pod-1.eu-west-1a.domain.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"100.67.94.101"}},
				{DNSName: "pod-2.eu-west-1b.domain.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"100.67.94.102"}},
			},
		},
		{
			title: "templating with shared all domain",
			pods: []*v1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "pod-1",
						Namespace: "kube-system",
					},
					Status: v1.PodStatus{
						PodIP: "100.67.94.101",
						PodIPs: []v1.PodIP{
							{IP: "100.67.94.101"},
							{IP: "100.67.94.102"},
							{IP: "100.67.94.103"},
							{IP: "2041:0000:140F::875B:131B"},
							{IP: "::11.22.33.44"},
						},
					},
				},
			},
			fqdnTemplate: "pods-all.domain.tld",
			expected: []*endpoint.Endpoint{
				{DNSName: "pods-all.domain.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"100.67.94.101", "100.67.94.102", "100.67.94.103"}},
				{DNSName: "pods-all.domain.tld", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2041:0000:140F::875B:131B", "::11.22.33.44"}},
			},
		},
		{
			title: "templating with fqdn template and IP not set as pod failed",
			pods: []*v1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "pod-1",
						Namespace: "kube-system",
					},
					Status: v1.PodStatus{
						Phase: v1.PodRunning,
						PodIP: "100.67.94.101",
						PodIPs: []v1.PodIP{
							{IP: "100.67.94.101"},
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "pod-2",
						Namespace: "kube-system",
					},
					Status: v1.PodStatus{
						Phase: v1.PodFailed,
					},
				},
			},
			fqdnTemplate: "{{ .Name }}.domain.tld",
			expected: []*endpoint.Endpoint{
				{DNSName: "pod-1.domain.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"100.67.94.101"}},
			},
		},
	} {
		t.Run(tt.title, func(t *testing.T) {
			kubeClient := fake.NewClientset()

			for _, node := range tt.nodes {
				_, err := kubeClient.CoreV1().Nodes().Create(t.Context(), node, metav1.CreateOptions{})
				require.NoError(t, err)
			}

			for _, pod := range tt.pods {
				_, err := kubeClient.CoreV1().Pods(pod.Namespace).Create(t.Context(), pod, metav1.CreateOptions{})
				require.NoError(t, err)
			}

			src, err := NewPodSource(
				t.Context(),
				kubeClient,
				"",
				"",
				false,
				tt.sourceDomain,
				tt.fqdnTemplate,
				tt.combineFQDN)
			require.NoError(t, err)

			endpoints, err := src.Endpoints(t.Context())
			require.NoError(t, err)

			validateEndpoints(t, endpoints, tt.expected)
		})
	}
}

func TestPodSourceFqdnTemplatingExamples_Failed(t *testing.T) {
	for _, tt := range []struct {
		title        string
		pods         []*v1.Pod
		nodes        []*v1.Node
		fqdnTemplate string
		expected     []*endpoint.Endpoint
		combineFQDN  bool
		sourceDomain string
	}{
		{
			title: "templating with fqdn template correct but value does not exist",
			pods: []*v1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "pod-1",
						Namespace: "kube-system",
					},
					Status: v1.PodStatus{
						PodIP: "100.67.94.101",
						PodIPs: []v1.PodIP{
							{IP: "100.67.94.101"},
						},
					},
				},
			},
			fqdnTemplate: "{{ .Name }}.{{ .ThisNotExist }}.domain.tld",
			expected:     []*endpoint.Endpoint{},
		},
	} {
		t.Run(tt.title, func(t *testing.T) {
			kubeClient := fake.NewClientset()

			for _, node := range tt.nodes {
				_, err := kubeClient.CoreV1().Nodes().Create(t.Context(), node, metav1.CreateOptions{})
				require.NoError(t, err)
			}

			for _, pod := range tt.pods {
				_, err := kubeClient.CoreV1().Pods(pod.Namespace).Create(t.Context(), pod, metav1.CreateOptions{})
				require.NoError(t, err)
			}

			src, err := NewPodSource(
				t.Context(),
				kubeClient,
				"",
				"",
				false,
				tt.sourceDomain,
				tt.fqdnTemplate,
				tt.combineFQDN)
			require.NoError(t, err)

			_, err = src.Endpoints(t.Context())
			require.Error(t, err)
		})
	}
}
