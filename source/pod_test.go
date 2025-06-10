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
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"

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
		expectedDebugMsgs        []string
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
			nil,
			nodesFixturesIPv4(),
			[]*corev1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-pod1",
						Namespace: "kube-system",
						Annotations: map[string]string{
							internalHostnameAnnotationKey: "internal.a.foo.example.org",
							hostnameAnnotationKey:         "a.foo.example.org",
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
							internalHostnameAnnotationKey: "internal.a.foo.example.org",
							hostnameAnnotationKey:         "a.foo.example.org",
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
			nil,
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
			nil,
			nodesFixturesIPv6(),
			[]*corev1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-pod1",
						Namespace: "kube-system",
						Annotations: map[string]string{
							internalHostnameAnnotationKey: "internal.a.foo.example.org",
							hostnameAnnotationKey:         "a.foo.example.org",
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
							internalHostnameAnnotationKey: "internal.a.foo.example.org",
							hostnameAnnotationKey:         "a.foo.example.org",
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
			nil,
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
			nil,
			nodesFixturesIPv4(),
			[]*corev1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-pod1",
						Namespace: "kube-system",
						Annotations: map[string]string{
							internalHostnameAnnotationKey: "internal.a.foo.example.org",
							hostnameAnnotationKey:         "a.foo.example.org",
							targetAnnotationKey:           "208.1.2.1",
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
							internalHostnameAnnotationKey: "internal.a.foo.example.org",
							hostnameAnnotationKey:         "a.foo.example.org",
							targetAnnotationKey:           "208.1.2.2",
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
				{DNSName: "a.foo.example.org", Targets: endpoint.Targets{"54.10.11.1"}, RecordType: endpoint.RecordTypeA},
				{DNSName: "a.foo.example.org", Targets: endpoint.Targets{"2001:DB8::1"}, RecordType: endpoint.RecordTypeAAAA},
				{DNSName: "b.foo.example.org", Targets: endpoint.Targets{"54.10.11.2"}, RecordType: endpoint.RecordTypeA},
			},
			false,
			nil,
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
							hostnameAnnotationKey: "a.foo.example.org",
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
							hostnameAnnotationKey: "b.foo.example.org",
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
				{DNSName: "a.foo.example.org", Targets: endpoint.Targets{"54.10.11.1"}, RecordType: endpoint.RecordTypeA},
				{DNSName: "internal.a.foo.example.org", Targets: endpoint.Targets{"10.0.1.1"}, RecordType: endpoint.RecordTypeA},
			},
			false,
			[]string{"skipping pod my-pod2. hostNetwork=false"},
			nodesFixturesIPv4(),
			[]*corev1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-pod1",
						Namespace: "kube-system",
						Annotations: map[string]string{
							internalHostnameAnnotationKey: "internal.a.foo.example.org",
							hostnameAnnotationKey:         "a.foo.example.org",
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
							internalHostnameAnnotationKey: "internal.a.foo.example.org",
							hostnameAnnotationKey:         "a.foo.example.org",
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
			nil,
			nodesFixturesIPv4(),
			[]*corev1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-pod1",
						Namespace: "kube-system",
						Annotations: map[string]string{
							internalHostnameAnnotationKey: "internal.a.foo.example.org",
							hostnameAnnotationKey:         "a.foo.example.org",
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
							internalHostnameAnnotationKey: "internal.a.foo.example.org",
							hostnameAnnotationKey:         "a.foo.example.org",
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
			nil,
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
							internalHostnameAnnotationKey: "internal.a.foo.example.org,internal.b.foo.example.org",
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
				{DNSName: "my-pod1.example.org", Targets: endpoint.Targets{"192.168.1.1"}, RecordType: endpoint.RecordTypeA},
				{DNSName: "my-pod2.example.org", Targets: endpoint.Targets{"192.168.1.2"}, RecordType: endpoint.RecordTypeA},
			},
			false,
			nil,
			nodesFixturesIPv4(),
			[]*corev1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:        "my-pod1",
						Namespace:   "kube-system",
						Annotations: map[string]string{},
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
			nil,
			nodesFixturesIPv4(),
			[]*corev1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-pod1",
						Namespace: "kube-system",
						Annotations: map[string]string{
							internalHostnameAnnotationKey: "internal.a.foo.example.org",
							hostnameAnnotationKey:         "a.foo.example.org",
							targetAnnotationKey:           "208.1.2.1",
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
							internalHostnameAnnotationKey: "internal.a.foo.example.org",
							hostnameAnnotationKey:         "a.foo.example.org",
							targetAnnotationKey:           "208.1.2.2",
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
			[]string{`Get node[missing-node] of pod[my-pod1] error: node "missing-node" not found; ignoring`},
			nodesFixturesIPv4(),
			[]*corev1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "my-pod1",
						Namespace: "kube-system",
						Annotations: map[string]string{
							hostnameAnnotationKey: "a.foo.example.org",
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

			for _, pod := range tc.pods {
				pods := kubernetes.CoreV1().Pods(pod.Namespace)

				if _, err := pods.Create(ctx, pod, metav1.CreateOptions{}); err != nil {
					t.Fatal(err)
				}
			}

			client, err := NewPodSource(t.Context(), kubernetes, tc.targetNamespace, tc.compatibility, tc.ignoreNonHostNetworkPods, tc.PodSourceDomain, "", false)
			require.NoError(t, err)

			hook := testutils.LogsUnderTestWithLogLevel(log.DebugLevel, t)

			endpoints, err := client.Endpoints(ctx)
			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			if tc.expectedDebugMsgs != nil {
				for _, expectedMsg := range tc.expectedDebugMsgs {
					testutils.TestHelperLogContains(expectedMsg, hook, t)
				}
			} else {
				require.Empty(t, hook.AllEntries(), "Expected no debug messages")
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
