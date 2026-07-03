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

	"github.com/stretchr/testify/require"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	"sigs.k8s.io/external-dns/source/annotations"
	templatetest "sigs.k8s.io/external-dns/source/template/testutil"
)

func TestPodFQDNTemplate(t *testing.T) {
	const (
		podName = "my-pod"
		podIP   = "100.67.94.101"
	)

	makePod := func(name, namespace, ip string, anns map[string]string) *v1.Pod {
		return &v1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:        name,
				Namespace:   namespace,
				Annotations: anns,
			},
			Status: v1.PodStatus{
				PodIP:  ip,
				PodIPs: []v1.PodIP{{IP: ip}},
			},
		}
	}

	// defaultPod is used for simple single-pod cases (fqdn-target-template cases).
	defaultPod := func(anns map[string]string) *v1.Pod {
		return makePod(podName, "default", podIP, anns)
	}

	for _, tt := range []struct {
		title              string
		pods               []*v1.Pod // nil = [defaultPod(nil)]
		nodes              []*v1.Node
		fqdnTemplate       string
		targetTemplate     string
		fqdnTargetTemplate string
		combine            bool
		sourceDomain       string
		expected           []*endpoint.Endpoint
	}{
		// ── fqdn-target-template cases ────────────────────────────────────────
		{
			title:              "fqdn-target-template generates A record when no annotation-derived endpoints",
			fqdnTargetTemplate: "{{.Name}}.example.com:1.2.3.4",
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint(podName+".example.com", endpoint.RecordTypeA, "1.2.3.4"),
			},
		},
		{
			title:              "fqdn-target-template generates CNAME for hostname target",
			fqdnTargetTemplate: "{{.Name}}.example.com:lb.example.com",
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint(podName+".example.com", endpoint.RecordTypeCNAME, "lb.example.com"),
			},
		},
		{
			title: "fqdn-target-template with combine adds endpoint alongside annotation-derived",
			pods: []*v1.Pod{defaultPod(map[string]string{
				annotations.InternalHostnameKey: "annotated.example.com",
			})},
			fqdnTargetTemplate: "{{.Name}}.tmpl.example.com:lb.example.com",
			combine:            true,
			expected: []*endpoint.Endpoint{
				{DNSName: "annotated.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{podIP}},
				{DNSName: podName + ".tmpl.example.com", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"lb.example.com"}},
			},
		},
		{
			title: "fqdn-target-template without combine is ignored when annotation-derived endpoints exist",
			pods: []*v1.Pod{defaultPod(map[string]string{
				annotations.InternalHostnameKey: "annotated.example.com",
			})},
			fqdnTargetTemplate: "{{.Name}}.tmpl.example.com:lb.example.com",
			combine:            false,
			expected: []*endpoint.Endpoint{
				{DNSName: "annotated.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{podIP}},
			},
		},
		{
			title:              "fqdn-target-template can reference .Kind",
			fqdnTargetTemplate: "{{.Kind | toLower}}.{{.Name}}.example.com:1.2.3.4",
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("pod."+podName+".example.com", endpoint.RecordTypeA, "1.2.3.4"),
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
				{DNSName: "pod." + podName + ".example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{podIP}},
			},
		},
		{
			title:              "fqdn-target-template can reference .APIVersion",
			fqdnTargetTemplate: "{{.Name}}.{{.APIVersion}}.example.com:1.2.3.4",
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint(podName+".v1.example.com", endpoint.RecordTypeA, "1.2.3.4"),
			},
		},
		// ── fqdn-template cases ───────────────────────────────────────────────
		{
			title:        "fqdn-template expansion with multiple domains",
			fqdnTemplate: "{{ .Name }}.domainA.com,{{ .Name }}.domainB.com",
			pods: []*v1.Pod{
				makePod("my-pod-1", "default", "100.67.94.101", nil),
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "my-pod-1.domainA.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"100.67.94.101"}},
				{DNSName: "my-pod-1.domainB.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"100.67.94.101"}},
			},
		},
		{
			title:        "fqdn-template with combine and pod source domain",
			fqdnTemplate: "{{ .Name }}.domainA.com,{{ .Name }}.domainB.com",
			combine:      true,
			sourceDomain: "example.org",
			pods: []*v1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{Name: "my-pod-1", Namespace: "default"},
					Spec:       v1.PodSpec{NodeName: "node-1.internal"},
					Status: v1.PodStatus{
						PodIP:  "100.67.94.101",
						PodIPs: []v1.PodIP{{IP: "100.67.94.101"}},
					},
				},
			},
			nodes: []*v1.Node{
				{
					ObjectMeta: metav1.ObjectMeta{Name: "node-1.internal"},
					Status: v1.NodeStatus{
						Addresses: []v1.NodeAddress{{Type: v1.NodeExternalIP, Address: "10.1.192.139"}},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "my-pod-1.domainA.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"100.67.94.101"}},
				{DNSName: "my-pod-1.domainB.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"100.67.94.101"}},
				{DNSName: "my-pod-1.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"100.67.94.101"}},
			},
		},
		{
			title:        "fqdn-template with namespace in domain",
			fqdnTemplate: "{{ .Name }}.{{ .Namespace }}.example.org",
			pods: []*v1.Pod{
				makePod("pod-1", "default", "100.67.94.101", nil),
				makePod("pod-2", "kube-system", "100.67.94.102", nil),
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "pod-1.default.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"100.67.94.101"}},
				{DNSName: "pod-2.kube-system.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"100.67.94.102"}},
			},
		},
		{
			title:        "fqdn-template with dual-stack pod produces A and AAAA records",
			fqdnTemplate: "{{ .Name }}.example.org",
			pods: []*v1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{Name: "pod-1", Namespace: "default"},
					Status: v1.PodStatus{
						PodIP:  "100.67.94.101",
						PodIPs: []v1.PodIP{{IP: "100.67.94.101"}, {IP: "2041:0000:140F::875B:131B"}},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "pod-1.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"100.67.94.101"}},
				{DNSName: "pod-1.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2041:0000:140F::875B:131B"}},
			},
		},
		{
			title:        "fqdn-template target annotation does not override pod IPs",
			fqdnTemplate: "{{ .Name }}.example.org",
			pods: []*v1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "pod-1", Namespace: "default",
						Annotations: map[string]string{"external-dns.kubernetes.io/target": "203.2.45.22"},
					},
					Status: v1.PodStatus{PodIP: "100.67.94.101", PodIPs: []v1.PodIP{{IP: "100.67.94.101"}}},
				},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "pod-1.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"100.67.94.101"}},
			},
		},
		{
			title:        "fqdn-template hostname annotation does not override template hostname",
			fqdnTemplate: "{{ .Name }}.example.org",
			pods: []*v1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "pod-1", Namespace: "default",
						Annotations: map[string]string{"external-dns.kubernetes.io/hostname": "ip-10-1-176-1.internal.domain.com"},
					},
					Status: v1.PodStatus{PodIP: "100.67.94.101", PodIPs: []v1.PodIP{{IP: "100.67.94.101"}}},
				},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "pod-1.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"100.67.94.101"}},
			},
		},
		{
			title:        "fqdn-template expands annotation value",
			fqdnTemplate: "{{ .Name }}.{{ .Annotations.workload }}.domain.tld",
			pods: []*v1.Pod{
				makePod("pod-1", "kube-system", "100.67.94.101", map[string]string{"workload": "cluster-resources"}),
				makePod("pod-2", "workloads", "100.67.94.102", map[string]string{"workload": "workloads"}),
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "pod-1.cluster-resources.domain.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"100.67.94.101"}},
				{DNSName: "pod-2.workloads.domain.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"100.67.94.102"}},
			},
		},
		{
			title:        "fqdn-template expands label value",
			fqdnTemplate: `{{ .Name }}.{{ index .ObjectMeta.Labels "topology.kubernetes.io/region" }}.domain.tld`,
			pods: []*v1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "pod-1", Namespace: "kube-system",
						Labels: map[string]string{"topology.kubernetes.io/region": "eu-west-1a"},
					},
					Status: v1.PodStatus{PodIP: "100.67.94.101", PodIPs: []v1.PodIP{{IP: "100.67.94.101"}}},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "pod-2", Namespace: "workloads",
						Labels: map[string]string{"topology.kubernetes.io/region": "eu-west-1b"},
					},
					Status: v1.PodStatus{PodIP: "100.67.94.102", PodIPs: []v1.PodIP{{IP: "100.67.94.102"}}},
				},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "pod-1.eu-west-1a.domain.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"100.67.94.101"}},
				{DNSName: "pod-2.eu-west-1b.domain.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"100.67.94.102"}},
			},
		},
		{
			title:        "fqdn-template shared domain merges targets from multiple pods",
			fqdnTemplate: "pods-all.domain.tld",
			pods: []*v1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{Name: "pod-1", Namespace: "kube-system"},
					Status: v1.PodStatus{
						PodIP: "100.67.94.101",
						PodIPs: []v1.PodIP{
							{IP: "100.67.94.101"}, {IP: "100.67.94.102"}, {IP: "100.67.94.103"},
							{IP: "2041:0000:140F::875B:131B"}, {IP: "::11.22.33.44"},
						},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "pods-all.domain.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"100.67.94.101", "100.67.94.102", "100.67.94.103"}},
				{DNSName: "pods-all.domain.tld", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2041:0000:140F::875B:131B", "::11.22.33.44"}},
			},
		},
		{
			title:        "fqdn-template skips pods with empty IP",
			fqdnTemplate: "{{ .Name }}.domain.tld",
			pods: []*v1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{Name: "pod-1", Namespace: "kube-system"},
					Status:     v1.PodStatus{Phase: v1.PodRunning, PodIP: "100.67.94.101", PodIPs: []v1.PodIP{{IP: "100.67.94.101"}}},
				},
				{
					ObjectMeta: metav1.ObjectMeta{Name: "pod-2", Namespace: "kube-system"},
					Status:     v1.PodStatus{Phase: v1.PodFailed},
				},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "pod-1.domain.tld", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"100.67.94.101"}},
			},
		},
		{
			title: "fqdn-template can reference .Kind and filter by label",
			fqdnTemplate: `{{ if eq .Kind "Pod" }}{{ range $k, $v := .Labels }}{{ if and (contains $k "app")
				(contains $v "my-service-") }}{{ $.Name }}.{{ $v }}.pod.tld.org{{ printf "," }}{{ end }}{{ end }}{{ end }}`,
			pods: []*v1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "pod-1", Namespace: "kube-system",
						Labels: map[string]string{"app1": "my-service-1"},
					},
					Status: v1.PodStatus{Phase: v1.PodRunning, PodIP: "100.67.94.101", PodIPs: []v1.PodIP{{IP: "100.67.94.101"}}},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "pod-2", Namespace: "kube-system",
						Labels: map[string]string{"app2": "my-service-2"},
					},
					Status: v1.PodStatus{Phase: v1.PodRunning, PodIPs: []v1.PodIP{{IP: "100.67.94.102"}}},
				},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "pod-1.my-service-1.pod.tld.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"100.67.94.101"}},
				{DNSName: "pod-2.my-service-2.pod.tld.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"100.67.94.102"}},
			},
		},
	} {
		t.Run(tt.title, func(t *testing.T) {
			kubeClient := fake.NewClientset()

			for _, node := range tt.nodes {
				_, err := kubeClient.CoreV1().Nodes().Create(t.Context(), node, metav1.CreateOptions{})
				require.NoError(t, err)
			}

			pods := tt.pods
			if pods == nil {
				pods = []*v1.Pod{defaultPod(nil)}
			}
			for _, pod := range pods {
				_, err := kubeClient.CoreV1().Pods(pod.Namespace).Create(t.Context(), pod, metav1.CreateOptions{})
				require.NoError(t, err)
			}

			src, err := NewPodSource(t.Context(), kubeClient, &Config{
				TemplateEngine:  templatetest.MustEngine(t, tt.fqdnTemplate, tt.targetTemplate, tt.fqdnTargetTemplate, tt.combine),
				PodSourceDomain: tt.sourceDomain,
			})
			require.NoError(t, err)

			endpoints, err := src.Endpoints(t.Context())
			require.NoError(t, err)
			testutils.ValidateEndpoints(t, endpoints, tt.expected)
		})
	}
}

func TestPodFQDNTemplate_Error(t *testing.T) {
	kubeClient := fake.NewClientset()
	_, err := kubeClient.CoreV1().Pods("kube-system").Create(t.Context(), &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{Name: "pod-1", Namespace: "kube-system"},
		Status:     v1.PodStatus{PodIP: "100.67.94.101", PodIPs: []v1.PodIP{{IP: "100.67.94.101"}}},
	}, metav1.CreateOptions{})
	require.NoError(t, err)

	src, err := NewPodSource(t.Context(), kubeClient, &Config{
		TemplateEngine: templatetest.MustEngine(t, "{{ .Name }}.{{ .ThisNotExist }}.domain.tld", "", "", false),
	})
	require.NoError(t, err)

	_, err = src.Endpoints(t.Context())
	require.Error(t, err)
}
