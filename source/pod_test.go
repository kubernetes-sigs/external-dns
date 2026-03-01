/*
Copyright 2021 The Kubernetes Authors.

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
	"fmt"
	"math/rand"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1lister "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"

	"sigs.k8s.io/external-dns/endpoint"
	logtest "sigs.k8s.io/external-dns/internal/testutils/log"
	"sigs.k8s.io/external-dns/source/annotations"

	"k8s.io/client-go/kubernetes/fake"
)

// testPodSource tests that various services generate the correct endpoints.
func TestPodSource(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		title                    string
		targetNamespace          string
		compatibility            string
		ignoreNonHostNetworkPods bool
		PodSourceDomain          string
		expected                 []*endpoint.Endpoint
		expectError              bool
		nodes                    []*corev1.Node
		pods                     []*corev1.Pod
	}{
		{
			"create IPv4 records based on pod's external and internal IPs",
			"",
			"",
			true,
			"",
			[]*endpoint.Endpoint{
				{DNSName: "a.foo.example.org", Targets: endpoint.Targets{"54.10.11.1", "54.10.11.2"}, RecordType: endpoint.RecordTypeA},
				{DNSName: "internal.a.foo.example.org", Targets: endpoint.Targets{"10.0.1.1", "10.0.1.2"}, RecordType: endpoint.RecordTypeA},
			},
			false,
			nodesFixturesIPv4(),
			[]*corev1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-pod1",
						Namespace: "kube-system",
						Annotations: map[string]string{
							annotations.InternalHostnameKey: "internal.a.foo.example.org",
							annotations.HostnameKey:         "a.foo.example.org",
						},
					},
					Spec: corev1.PodSpec{
						HostNetwork: true,
						NodeName:    "my-node1",
					},
					Status: corev1.PodStatus{
						PodIP: "10.0.1.1",
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-pod2",
						Namespace: "kube-system",
						Annotations: map[string]string{
							annotations.InternalHostnameKey: "internal.a.foo.example.org",
							annotations.HostnameKey:         "a.foo.example.org",
						},
					},
					Spec: corev1.PodSpec{
						HostNetwork: true,
						NodeName:    "my-node2",
					},
					Status: corev1.PodStatus{
						PodIP: "10.0.1.2",
					},
				},
			},
		},
		{
			"create IPv4 records based on pod's external and internal IPs using DNS Controller annotations",
			"",
			"kops-dns-controller",
			true,
			"",
			[]*endpoint.Endpoint{
				{DNSName: "a.foo.example.org", Targets: endpoint.Targets{"54.10.11.1", "54.10.11.2"}, RecordType: endpoint.RecordTypeA},
				{DNSName: "internal.a.foo.example.org", Targets: endpoint.Targets{"10.0.1.1", "10.0.1.2"}, RecordType: endpoint.RecordTypeA},
			},
			false,
			nodesFixturesIPv4(),
			[]*corev1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-pod1",
						Namespace: "kube-system",
						Annotations: map[string]string{
							kopsDNSControllerInternalHostnameAnnotationKey: "internal.a.foo.example.org",
							kopsDNSControllerHostnameAnnotationKey:         "a.foo.example.org",
						},
					},
					Spec: corev1.PodSpec{
						HostNetwork: true,
						NodeName:    "my-node1",
					},
					Status: corev1.PodStatus{
						PodIP: "10.0.1.1",
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-pod2",
						Namespace: "kube-system",
						Annotations: map[string]string{
							kopsDNSControllerInternalHostnameAnnotationKey: "internal.a.foo.example.org",
							kopsDNSControllerHostnameAnnotationKey:         "a.foo.example.org",
						},
					},
					Spec: corev1.PodSpec{
						HostNetwork: true,
						NodeName:    "my-node2",
					},
					Status: corev1.PodStatus{
						PodIP: "10.0.1.2",
					},
				},
			},
		},
		{
			"create IPv6 records based on pod's external and internal IPs",
			"",
			"",
			true,
			"",
			[]*endpoint.Endpoint{
				{DNSName: "a.foo.example.org", Targets: endpoint.Targets{"2001:DB8::1", "2001:DB8::2"}, RecordType: endpoint.RecordTypeAAAA},
				{DNSName: "internal.a.foo.example.org", Targets: endpoint.Targets{"2001:DB8::1", "2001:DB8::2"}, RecordType: endpoint.RecordTypeAAAA},
			},
			false,
			nodesFixturesIPv6(),
			[]*corev1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-pod1",
						Namespace: "kube-system",
						Annotations: map[string]string{
							annotations.InternalHostnameKey: "internal.a.foo.example.org",
							annotations.HostnameKey:         "a.foo.example.org",
						},
					},
					Spec: corev1.PodSpec{
						HostNetwork: true,
						NodeName:    "my-node1",
					},
					Status: corev1.PodStatus{
						PodIP: "2001:DB8::1",
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-pod2",
						Namespace: "kube-system",
						Annotations: map[string]string{
							annotations.InternalHostnameKey: "internal.a.foo.example.org",
							annotations.HostnameKey:         "a.foo.example.org",
						},
					},
					Spec: corev1.PodSpec{
						HostNetwork: true,
						NodeName:    "my-node2",
					},
					Status: corev1.PodStatus{
						PodIP: "2001:DB8::2",
					},
				},
			},
		},
		{
			"create IPv6 records based on pod's external and internal IPs using DNS Controller annotations",
			"",
			"kops-dns-controller",
			true,
			"",
			[]*endpoint.Endpoint{
				{DNSName: "a.foo.example.org", Targets: endpoint.Targets{"2001:DB8::1", "2001:DB8::2"}, RecordType: endpoint.RecordTypeAAAA},
				{DNSName: "internal.a.foo.example.org", Targets: endpoint.Targets{"2001:DB8::1", "2001:DB8::2"}, RecordType: endpoint.RecordTypeAAAA},
			},
			false,
			nodesFixturesIPv6(),
			[]*corev1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-pod1",
						Namespace: "kube-system",
						Annotations: map[string]string{
							kopsDNSControllerInternalHostnameAnnotationKey: "internal.a.foo.example.org",
							kopsDNSControllerHostnameAnnotationKey:         "a.foo.example.org",
						},
					},
					Spec: corev1.PodSpec{
						HostNetwork: true,
						NodeName:    "my-node1",
					},
					Status: corev1.PodStatus{
						PodIP: "2001:DB8::1",
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-pod2",
						Namespace: "kube-system",
						Annotations: map[string]string{
							kopsDNSControllerInternalHostnameAnnotationKey: "internal.a.foo.example.org",
							kopsDNSControllerHostnameAnnotationKey:         "a.foo.example.org",
						},
					},
					Spec: corev1.PodSpec{
						HostNetwork: true,
						NodeName:    "my-node2",
					},
					Status: corev1.PodStatus{
						PodIP: "2001:DB8::2",
					},
				},
			},
		},
		{
			"create records based on pod's target annotation",
			"",
			"",
			true,
			"",
			[]*endpoint.Endpoint{
				{DNSName: "a.foo.example.org", Targets: endpoint.Targets{"208.1.2.1", "208.1.2.2"}, RecordType: endpoint.RecordTypeA},
				{DNSName: "internal.a.foo.example.org", Targets: endpoint.Targets{"208.1.2.1", "208.1.2.2"}, RecordType: endpoint.RecordTypeA},
			},
			false,
			nodesFixturesIPv4(),
			[]*corev1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-pod1",
						Namespace: "kube-system",
						Annotations: map[string]string{
							annotations.InternalHostnameKey: "internal.a.foo.example.org",
							annotations.HostnameKey:         "a.foo.example.org",
							annotations.TargetKey:           "208.1.2.1",
						},
					},
					Spec: corev1.PodSpec{
						HostNetwork: true,
						NodeName:    "my-node1",
					},
					Status: corev1.PodStatus{
						PodIP: "10.0.1.1",
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-pod2",
						Namespace: "kube-system",
						Annotations: map[string]string{
							annotations.InternalHostnameKey: "internal.a.foo.example.org",
							annotations.HostnameKey:         "a.foo.example.org",
							annotations.TargetKey:           "208.1.2.2",
						},
					},
					Spec: corev1.PodSpec{
						HostNetwork: true,
						NodeName:    "my-node2",
					},
					Status: corev1.PodStatus{
						PodIP: "10.0.1.2",
					},
				},
			},
		},
		{
			"create multiple records",
			"",
			"",
			true,
			"",
			[]*endpoint.Endpoint{
				{DNSName: "a.foo.example.org", Targets: endpoint.Targets{"54.10.11.1"}, RecordType: endpoint.RecordTypeA, RecordTTL: endpoint.TTL(5400)},
				{DNSName: "a.foo.example.org", Targets: endpoint.Targets{"2001:DB8::1"}, RecordType: endpoint.RecordTypeAAAA, RecordTTL: endpoint.TTL(5400)},
				{DNSName: "b.foo.example.org", Targets: endpoint.Targets{"54.10.11.2"}, RecordType: endpoint.RecordTypeA},
			},
			false,
			[]*corev1.Node{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "my-node1",
					},
					Status: corev1.NodeStatus{
						Addresses: []corev1.NodeAddress{
							{Type: corev1.NodeExternalIP, Address: "54.10.11.1"},
							{Type: corev1.NodeInternalIP, Address: "2001:DB8::1"},
							{Type: corev1.NodeInternalIP, Address: "10.0.1.1"},
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "my-node2",
					},
					Status: corev1.NodeStatus{
						Addresses: []corev1.NodeAddress{
							{Type: corev1.NodeExternalIP, Address: "54.10.11.2"},
							{Type: corev1.NodeInternalIP, Address: "10.0.1.2"},
						},
					},
				},
			},
			[]*corev1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-pod1",
						Namespace: "kube-system",
						Annotations: map[string]string{
							annotations.HostnameKey: "a.foo.example.org",
							annotations.TtlKey:      "1h30m",
						},
					},
					Spec: corev1.PodSpec{
						HostNetwork: true,
						NodeName:    "my-node1",
					},
					Status: corev1.PodStatus{
						PodIP: "10.0.1.1",
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-pod2",
						Namespace: "kube-system",
						Annotations: map[string]string{
							annotations.HostnameKey: "b.foo.example.org",
						},
					},
					Spec: corev1.PodSpec{
						HostNetwork: true,
						NodeName:    "my-node2",
					},
					Status: corev1.PodStatus{
						PodIP: "10.0.1.2",
					},
				},
			},
		},
		{
			"pods with hostNetwore=false should be ignored",
			"",
			"",
			true,
			"",
			[]*endpoint.Endpoint{
				{DNSName: "a.foo.example.org", Targets: endpoint.Targets{"54.10.11.1"}, RecordType: endpoint.RecordTypeA, RecordTTL: endpoint.TTL(1)},
				{DNSName: "internal.a.foo.example.org", Targets: endpoint.Targets{"10.0.1.1"}, RecordType: endpoint.RecordTypeA, RecordTTL: endpoint.TTL(1)},
			},
			false,
			nodesFixturesIPv4(),
			[]*corev1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-pod1",
						Namespace: "kube-system",
						Annotations: map[string]string{
							annotations.InternalHostnameKey: "internal.a.foo.example.org",
							annotations.HostnameKey:         "a.foo.example.org",
							annotations.TtlKey:              "1s",
						},
					},
					Spec: corev1.PodSpec{
						HostNetwork: true,
						NodeName:    "my-node1",
					},
					Status: corev1.PodStatus{
						PodIP: "10.0.1.1",
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-pod2",
						Namespace: "kube-system",
						Annotations: map[string]string{
							annotations.InternalHostnameKey: "internal.a.foo.example.org",
							annotations.HostnameKey:         "a.foo.example.org",
						},
					},
					Spec: corev1.PodSpec{
						HostNetwork: false,
						NodeName:    "my-node2",
					},
					Status: corev1.PodStatus{
						PodIP: "100.0.1.2",
					},
				},
			},
		},
		{
			"only watch a given namespace",
			"kube-system",
			"",
			true,
			"",
			[]*endpoint.Endpoint{
				{DNSName: "a.foo.example.org", Targets: endpoint.Targets{"54.10.11.1"}, RecordType: endpoint.RecordTypeA},
				{DNSName: "internal.a.foo.example.org", Targets: endpoint.Targets{"10.0.1.1"}, RecordType: endpoint.RecordTypeA},
			},
			false,
			nodesFixturesIPv4(),
			[]*corev1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-pod1",
						Namespace: "kube-system",
						Annotations: map[string]string{
							annotations.InternalHostnameKey: "internal.a.foo.example.org",
							annotations.HostnameKey:         "a.foo.example.org",
						},
					},
					Spec: corev1.PodSpec{
						HostNetwork: true,
						NodeName:    "my-node1",
					},
					Status: corev1.PodStatus{
						PodIP: "10.0.1.1",
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-pod2",
						Namespace: "default",
						Annotations: map[string]string{
							annotations.InternalHostnameKey: "internal.a.foo.example.org",
							annotations.HostnameKey:         "a.foo.example.org",
							annotations.TtlKey:              "1s",
						},
					},
					Spec: corev1.PodSpec{
						HostNetwork: true,
						NodeName:    "my-node2",
					},
					Status: corev1.PodStatus{
						PodIP: "100.0.1.2",
					},
				},
			},
		},
		{
			"split record for internal hostname annotation",
			"",
			"",
			true,
			"",
			[]*endpoint.Endpoint{
				{DNSName: "internal.a.foo.example.org", Targets: endpoint.Targets{"10.0.1.1"}, RecordType: endpoint.RecordTypeA},
				{DNSName: "internal.b.foo.example.org", Targets: endpoint.Targets{"10.0.1.1"}, RecordType: endpoint.RecordTypeA},
			},
			false,
			[]*corev1.Node{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "my-node1",
					},
					Status: corev1.NodeStatus{
						Addresses: []corev1.NodeAddress{
							{Type: corev1.NodeInternalIP, Address: "10.0.1.1"},
						},
					},
				},
			},
			[]*corev1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-pod1",
						Namespace: "kube-system",
						Annotations: map[string]string{
							annotations.InternalHostnameKey: "internal.a.foo.example.org,internal.b.foo.example.org",
						},
					},
					Spec: corev1.PodSpec{
						HostNetwork: true,
						NodeName:    "my-node1",
					},
					Status: corev1.PodStatus{
						PodIP: "10.0.1.1",
					},
				},
			},
		},
		{
			"create IPv4 records for non-host network pods",
			"",
			"",
			false,
			"example.org",
			[]*endpoint.Endpoint{
				{DNSName: "my-pod1.example.org", Targets: endpoint.Targets{"192.168.1.1"}, RecordType: endpoint.RecordTypeA, RecordTTL: endpoint.TTL(60)},
				{DNSName: "my-pod2.example.org", Targets: endpoint.Targets{"192.168.1.2"}, RecordType: endpoint.RecordTypeA},
			},
			false,
			nodesFixturesIPv4(),
			[]*corev1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-pod1",
						Namespace: "kube-system",
						Annotations: map[string]string{
							annotations.TtlKey: "1m",
						},
					},
					Spec: corev1.PodSpec{
						HostNetwork: false,
						NodeName:    "my-node1",
					},
					Status: corev1.PodStatus{
						PodIP: "192.168.1.1",
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:        "my-pod2",
						Namespace:   "kube-system",
						Annotations: map[string]string{},
					},
					Spec: corev1.PodSpec{
						HostNetwork: false,
						NodeName:    "my-node2",
					},
					Status: corev1.PodStatus{
						PodIP: "192.168.1.2",
					},
				},
			},
		},
		{
			"create records based on pod's target annotation with pod source domain",
			"",
			"",
			true,
			"example.org",
			[]*endpoint.Endpoint{
				{DNSName: "my-pod1.example.org", Targets: endpoint.Targets{"208.1.2.1"}, RecordType: endpoint.RecordTypeA},
				{DNSName: "my-pod2.example.org", Targets: endpoint.Targets{"208.1.2.2"}, RecordType: endpoint.RecordTypeA},
				{DNSName: "a.foo.example.org", Targets: endpoint.Targets{"208.1.2.1", "208.1.2.2"}, RecordType: endpoint.RecordTypeA},
				{DNSName: "internal.a.foo.example.org", Targets: endpoint.Targets{"208.1.2.1", "208.1.2.2"}, RecordType: endpoint.RecordTypeA},
			},
			false,
			nodesFixturesIPv4(),
			[]*corev1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-pod1",
						Namespace: "kube-system",
						Annotations: map[string]string{
							annotations.InternalHostnameKey: "internal.a.foo.example.org",
							annotations.HostnameKey:         "a.foo.example.org",
							annotations.TargetKey:           "208.1.2.1",
						},
					},
					Spec: corev1.PodSpec{
						HostNetwork: true,
						NodeName:    "my-node1",
					},
					Status: corev1.PodStatus{
						PodIP: "10.0.1.1",
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-pod2",
						Namespace: "kube-system",
						Annotations: map[string]string{
							annotations.InternalHostnameKey: "internal.a.foo.example.org",
							annotations.HostnameKey:         "a.foo.example.org",
							annotations.TargetKey:           "208.1.2.2",
						},
					},
					Spec: corev1.PodSpec{
						HostNetwork: true,
						NodeName:    "my-node2",
					},
					Status: corev1.PodStatus{
						PodIP: "10.0.1.2",
					},
				},
			},
		},
		{
			"host network pod on a missing node",
			"",
			"",
			true,
			"",
			[]*endpoint.Endpoint{},
			false,
			nodesFixturesIPv4(),
			[]*corev1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-pod1",
						Namespace: "kube-system",
						Annotations: map[string]string{
							annotations.HostnameKey: "a.foo.example.org",
						},
					},
					Spec: corev1.PodSpec{
						HostNetwork: true,
						NodeName:    "missing-node",
					},
					Status: corev1.PodStatus{
						PodIP: "10.0.1.1",
					},
				},
			},
		},
	} {
		t.Run(tc.title, func(t *testing.T) {
			kubernetes := fake.NewClientset()
			ctx := t.Context()

			// Create the nodes
			for _, node := range tc.nodes {
				if _, err := kubernetes.CoreV1().Nodes().Create(ctx, node, metav1.CreateOptions{}); err != nil {
					t.Fatal(err)
				}
			}

			// Create the pods
			for _, pod := range tc.pods {
				pods := kubernetes.CoreV1().Pods(pod.Namespace)

				if _, err := pods.Create(ctx, pod, metav1.CreateOptions{}); err != nil {
					t.Fatal(err)
				}
			}

			client, err := NewPodSource(ctx, kubernetes, tc.targetNamespace, tc.compatibility, tc.ignoreNonHostNetworkPods, tc.PodSourceDomain, "", false, "", nil)
			require.NoError(t, err)

			endpoints, err := client.Endpoints(ctx)
			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			// Validate returned endpoints against desired endpoints.
			validateEndpoints(t, endpoints, tc.expected)

			for _, ep := range endpoints {
				// TODO: source should always set the resource label key. currently not supported by the pod source.
				require.Empty(t, ep.Labels, "Labels should not be empty for endpoint %s", ep.DNSName)
				require.NotContains(t, ep.Labels, endpoint.ResourceLabelKey)
			}
		})
	}
}

func TestPodSourceLogs(t *testing.T) {
	t.Parallel()
	// Generate unique pod names to avoid log conflicts across parallel tests.
	// Since logs are globally shared, using the same pod names could cause
	// false positives in unexpectedDebugLogs assertions.
	suffix := fmt.Sprintf("%d", rand.Intn(100000))
	for _, tc := range []struct {
		title                    string
		ignoreNonHostNetworkPods bool
		pods                     []*corev1.Pod
		nodes                    []*corev1.Node
		expectedDebugLogs        []string
		unexpectedDebugLogs      []string
	}{
		{
			"pods with hostNetwore=false should be skipped logging",
			true,
			[]*corev1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      fmt.Sprintf("my-pod1-%s", suffix),
						Namespace: "kube-system",
						Annotations: map[string]string{
							annotations.InternalHostnameKey: "internal.a.foo.example.org",
							annotations.HostnameKey:         "a.foo.example.org",
						},
					},
					Spec: corev1.PodSpec{
						HostNetwork: false,
						NodeName:    "my-node1",
					},
					Status: corev1.PodStatus{
						PodIP: "100.0.1.1",
					},
				},
			},
			nodesFixturesIPv4(),
			[]string{fmt.Sprintf("skipping pod my-pod1-%s. hostNetwork=false", suffix)},
			nil,
		},
		{
			"host network pod on a missing node",
			true,
			[]*corev1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      fmt.Sprintf("missing-node-pod-%s", suffix),
						Namespace: "kube-system",
						Annotations: map[string]string{
							annotations.HostnameKey: "a.foo.example.org",
						},
					},
					Spec: corev1.PodSpec{
						HostNetwork: true,
						NodeName:    "missing-node",
					},
					Status: corev1.PodStatus{
						PodIP: "10.0.1.1",
					},
				},
			},
			nodesFixturesIPv4(),
			[]string{
				fmt.Sprintf(`Get node[missing-node] of pod[missing-node-pod-%s] error: node "missing-node" not found; ignoring`, suffix),
			},
			nil,
		},
		{
			"mixed valid and hostNetwork=false pods with missing node",
			true,
			[]*corev1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      fmt.Sprintf("valid-pod-%s", suffix),
						Namespace: "kube-system",
						Annotations: map[string]string{
							annotations.HostnameKey: "valid.foo.example.org",
						},
					},
					Spec: corev1.PodSpec{
						HostNetwork: true,
						NodeName:    "my-node1",
					},
					Status: corev1.PodStatus{
						PodIP: "10.0.1.1",
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      fmt.Sprintf("non-hostnet-pod-%s", suffix),
						Namespace: "kube-system",
						Annotations: map[string]string{
							annotations.HostnameKey: "nonhost.foo.example.org",
						},
					},
					Spec: corev1.PodSpec{
						HostNetwork: false,
						NodeName:    "my-node2",
					},
					Status: corev1.PodStatus{
						PodIP: "100.0.1.2",
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      fmt.Sprintf("missing-node-pod-%s", suffix),
						Namespace: "kube-system",
						Annotations: map[string]string{
							annotations.HostnameKey: "missing.foo.example.org",
						},
					},
					Spec: corev1.PodSpec{
						HostNetwork: true,
						NodeName:    "missing-node",
					},
					Status: corev1.PodStatus{
						PodIP: "10.0.1.3",
					},
				},
			},
			nodesFixturesIPv4(),
			[]string{
				fmt.Sprintf("skipping pod non-hostnet-pod-%s. hostNetwork=false", suffix),
				fmt.Sprintf(`Get node[missing-node] of pod[missing-node-pod-%s] error: node "missing-node" not found; ignoring`, suffix),
			},
			[]string{
				fmt.Sprintf("skipping pod valid-pod-%s. hostNetwork=false", suffix),
				fmt.Sprintf(`Get node[my-node1] of pod[valid-pod-%s] error: node "my-node1" not found; ignoring`, suffix),
			},
		},
		{
			"valid pods with hostNetwork=true should not generate logs",
			true,
			[]*corev1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      fmt.Sprintf("valid-pod-%s", suffix),
						Namespace: "kube-system",
						Annotations: map[string]string{
							annotations.HostnameKey: "valid.foo.example.org",
						},
					},
					Spec: corev1.PodSpec{
						HostNetwork: true,
						NodeName:    "my-node1",
					},
					Status: corev1.PodStatus{
						PodIP: "10.0.1.1",
					},
				},
			},
			nodesFixturesIPv4(),
			nil,
			[]string{
				fmt.Sprintf("skipping pod valid-pod-%s. hostNetwork=false", suffix),
				fmt.Sprintf(`Get node[my-node1] of pod[valid-pod-%s] error: node "my-node1" not found; ignoring`, suffix),
			},
		},
		{
			"when ignoreNonHostNetworkPods=false, no skip logs should be generated",
			false,
			[]*corev1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      fmt.Sprintf("my-pod1-%s", suffix),
						Namespace: "kube-system",
						Annotations: map[string]string{
							annotations.InternalHostnameKey: "internal.a.foo.example.org",
							annotations.HostnameKey:         "a.foo.example.org",
						},
					},
					Spec: corev1.PodSpec{
						HostNetwork: false,
						NodeName:    "my-node1",
					},
					Status: corev1.PodStatus{
						PodIP: "10.0.1.1",
					},
				},
			},
			nodesFixturesIPv4(),
			nil,
			[]string{
				fmt.Sprintf("skipping pod my-pod1-%s. hostNetwork=false", suffix),
			},
		},
	} {
		t.Run(tc.title, func(t *testing.T) {
			kubernetes := fake.NewClientset()
			ctx := context.Background()
			// Create the nodes
			for _, node := range tc.nodes {
				if _, err := kubernetes.CoreV1().Nodes().Create(ctx, node, metav1.CreateOptions{}); err != nil {
					t.Fatal(err)
				}
			}

			// Create the pods
			for _, pod := range tc.pods {
				pods := kubernetes.CoreV1().Pods(pod.Namespace)

				if _, err := pods.Create(ctx, pod, metav1.CreateOptions{}); err != nil {
					t.Fatal(err)
				}
			}

			client, err := NewPodSource(ctx, kubernetes, "", "", tc.ignoreNonHostNetworkPods, "", "", false, "", nil)
			require.NoError(t, err)

			hook := logtest.LogsUnderTestWithLogLevel(log.DebugLevel, t)

			_, err = client.Endpoints(ctx)
			require.NoError(t, err)

			// Check if all expected logs are present in actual logs.
			// We don't do an exact match because logs are globally shared,
			// making precise comparisons difficult
			for _, expectedLog := range tc.expectedDebugLogs {
				logtest.TestHelperLogContains(expectedLog, hook, t)
			}

			// Check that no unexpected logs are present.
			// This ensures that logs are not generated inappropriately.
			for _, unexpectedLog := range tc.unexpectedDebugLogs {
				logtest.TestHelperLogNotContains(unexpectedLog, hook, t)
			}
		})
	}
}

func TestPodSource_AddEventHandler(t *testing.T) {
	fakeInformer := new(fakePodInformer)
	inf := testInformer{}
	fakeInformer.On("Informer").Return(&inf)

	pSource := &podSource{
		podInformer: fakeInformer,
	}

	handlerCalled := false
	handler := func() { handlerCalled = true }

	pSource.AddEventHandler(t.Context(), handler)

	fakeInformer.AssertNumberOfCalls(t, "Informer", 1)
	assert.False(t, handlerCalled)
	assert.Equal(t, 1, inf.times)
}

type fakePodInformer struct {
	mock.Mock
}

func (f *fakePodInformer) Informer() cache.SharedIndexInformer {
	args := f.Called()
	return args.Get(0).(cache.SharedIndexInformer)
}

func (f *fakePodInformer) Lister() corev1lister.PodLister {
	return corev1lister.NewPodLister(f.Informer().GetIndexer())
}

func nodesFixturesIPv6() []*corev1.Node {
	return []*corev1.Node{
		{
			ObjectMeta: metav1.ObjectMeta{
				Name: "my-node1",
			},
			Status: corev1.NodeStatus{
				Addresses: []corev1.NodeAddress{
					{Type: corev1.NodeInternalIP, Address: "2001:DB8::1"},
				},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				Name: "my-node2",
			},
			Status: corev1.NodeStatus{
				Addresses: []corev1.NodeAddress{
					{Type: corev1.NodeInternalIP, Address: "2001:DB8::2"},
				},
			},
		},
	}
}

func nodesFixturesIPv4() []*corev1.Node {
	return []*corev1.Node{
		{
			ObjectMeta: metav1.ObjectMeta{
				Name: "my-node1",
			},
			Status: corev1.NodeStatus{
				Addresses: []corev1.NodeAddress{
					{Type: corev1.NodeExternalIP, Address: "54.10.11.1"},
					{Type: corev1.NodeInternalIP, Address: "10.0.1.1"},
				},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				Name: "my-node2",
			},
			Status: corev1.NodeStatus{
				Addresses: []corev1.NodeAddress{
					{Type: corev1.NodeExternalIP, Address: "54.10.11.2"},
					{Type: corev1.NodeInternalIP, Address: "10.0.1.2"},
				},
			},
		},
	}
}

func TestPodTransformerInPodSource(t *testing.T) {
	t.Run("transformer set", func(t *testing.T) {
		ctx := t.Context()
		fakeClient := fake.NewClientset()

		pod := &v1.Pod{
			Spec: v1.PodSpec{
				Containers: []v1.Container{{
					Name: "test",
				}},
				Hostname:    "test-hostname",
				NodeName:    "test-node",
				HostNetwork: true,
			},
			ObjectMeta: metav1.ObjectMeta{
				Namespace: "test-ns",
				Name:      "test-name",
				Labels: map[string]string{
					"label1": "value1",
					"label2": "value2",
					"label3": "value3",
				},
				Annotations: map[string]string{
					"user-annotation": "value",
					"external-dns.alpha.kubernetes.io/hostname": "test-hostname",
					"external-dns.alpha.kubernetes.io/random":   "value",
					"other/annotation":                          "value",
				},
				UID: "someuid",
			},
			Status: v1.PodStatus{
				PodIP:  "127.0.0.1",
				HostIP: "127.0.0.2",
				Conditions: []v1.PodCondition{{
					Type:   v1.PodReady,
					Status: v1.ConditionTrue,
				}, {
					Type:   v1.ContainersReady,
					Status: v1.ConditionFalse,
				}},
			},
		}

		_, err := fakeClient.CoreV1().Pods(pod.Namespace).Create(context.Background(), pod, metav1.CreateOptions{})
		require.NoError(t, err)

		// Should not error when creating the source
		src, err := NewPodSource(ctx, fakeClient, "", "", false, "", "", false, "", nil)
		require.NoError(t, err)
		ps, ok := src.(*podSource)
		require.True(t, ok)

		retrieved, err := ps.podInformer.Lister().Pods("test-ns").Get("test-name")
		require.NoError(t, err)

		// Metadata
		assert.Equal(t, "test-name", retrieved.Name)
		assert.Equal(t, "test-ns", retrieved.Namespace)
		assert.Empty(t, retrieved.UID)
		assert.Empty(t, retrieved.Labels)
		// Filtered
		assert.Equal(t, map[string]string{
			"user-annotation": "value",
			"external-dns.alpha.kubernetes.io/hostname": "test-hostname",
			"external-dns.alpha.kubernetes.io/random":   "value",
			"other/annotation":                          "value",
		}, retrieved.Annotations)

		// Spec
		assert.Empty(t, retrieved.Spec.Containers)
		assert.Empty(t, retrieved.Spec.Hostname)
		assert.Equal(t, "test-node", retrieved.Spec.NodeName)
		assert.True(t, retrieved.Spec.HostNetwork)

		// Status
		assert.Empty(t, retrieved.Status.ContainerStatuses)
		assert.Empty(t, retrieved.Status.InitContainerStatuses)
		assert.Empty(t, retrieved.Status.HostIP)
		assert.Equal(t, "127.0.0.1", retrieved.Status.PodIP)
		assert.Empty(t, retrieved.Status.Conditions)
	})

	t.Run("transformer is not used when fqdnTemplate is set", func(t *testing.T) {
		ctx := t.Context()
		fakeClient := fake.NewClientset()

		pod := &v1.Pod{
			Spec: v1.PodSpec{
				Containers: []v1.Container{{
					Name: "test",
				}},
				Hostname:    "test-hostname",
				NodeName:    "test-node",
				HostNetwork: true,
			},
			ObjectMeta: metav1.ObjectMeta{
				Namespace: "test-ns",
				Name:      "test-name",
				Labels: map[string]string{
					"label1": "value1",
					"label2": "value2",
					"label3": "value3",
				},
				Annotations: map[string]string{
					"user-annotation": "value",
					"external-dns.alpha.kubernetes.io/hostname": "test-hostname",
					"external-dns.alpha.kubernetes.io/random":   "value",
					"other/annotation":                          "value",
				},
				UID: "someuid",
			},
			Status: v1.PodStatus{
				PodIP:  "127.0.0.1",
				HostIP: "127.0.0.2",
				Conditions: []v1.PodCondition{{
					Type:   v1.PodReady,
					Status: v1.ConditionTrue,
				}, {
					Type:   v1.ContainersReady,
					Status: v1.ConditionFalse,
				}},
			},
		}

		_, err := fakeClient.CoreV1().Pods(pod.Namespace).Create(context.Background(), pod, metav1.CreateOptions{})
		require.NoError(t, err)

		// Should not error when creating the source
		src, err := NewPodSource(ctx, fakeClient, "", "", false, "", "template", false, "", nil)
		require.NoError(t, err)
		ps, ok := src.(*podSource)
		require.True(t, ok)

		retrieved, err := ps.podInformer.Lister().Pods("test-ns").Get("test-name")
		require.NoError(t, err)

		// Metadata
		assert.Equal(t, "test-name", retrieved.Name)
		assert.Equal(t, "test-ns", retrieved.Namespace)
		assert.NotEmpty(t, retrieved.UID)
		assert.NotEmpty(t, retrieved.Labels)
	})
}
