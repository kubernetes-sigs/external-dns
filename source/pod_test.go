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
	"testing"

	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"sigs.k8s.io/external-dns/endpoint"

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
	} {

		t.Run(tc.title, func(t *testing.T) {
			t.Parallel()

			kubernetes := fake.NewClientset()
			ctx := context.Background()

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

			client, err := NewPodSource(context.TODO(), kubernetes, tc.targetNamespace, tc.compatibility, tc.ignoreNonHostNetworkPods, tc.PodSourceDomain)
			require.NoError(t, err)

			endpoints, err := client.Endpoints(ctx)
			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			// Validate returned endpoints against desired endpoints.
			validateEndpoints(t, endpoints, tc.expected)
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
