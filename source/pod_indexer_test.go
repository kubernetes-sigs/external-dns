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
	"math/rand/v2"
	"net"
	"strconv"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	"sigs.k8s.io/external-dns/source/annotations"
)

type podSpec struct {
	namespace   string
	labels      map[string]string
	annotations map[string]string
	// with labels and annotations
	totalTarget int
	// without provided labels and annotations
	totalRandom int
}

func fixtureCreatePodsWithNodes(input []podSpec) []*corev1.Pod {
	var pods []*corev1.Pod

	var createPod = func(index int, spec podSpec) *corev1.Pod {
		return &corev1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:      fmt.Sprintf("pod-%d-%s", index, uuid.NewString()),
				Namespace: spec.namespace,
				Labels: func() map[string]string {
					if spec.totalTarget > index {
						return spec.labels
					}
					return map[string]string{
						"app":   fmt.Sprintf("my-app-%d", rand.IntN(10)),
						"index": strconv.Itoa(index),
					}
				}(),
				Annotations: func() map[string]string {
					if spec.totalTarget > index {
						return spec.annotations
					}
					return map[string]string{
						"key1": fmt.Sprintf("value-%d", rand.IntN(10)),
					}
				}(),
			},
			Spec: corev1.PodSpec{},
			Status: corev1.PodStatus{
				Phase: corev1.PodRunning,
				PodIPs: []corev1.PodIP{
					{IP: net.IPv4(192, byte(rand.IntN(250)), byte(rand.IntN(250)), byte(index)).String()},
				},
			},
		}
	}

	for _, el := range input {
		totalPods := el.totalTarget + el.totalRandom
		for i := 0; i < totalPods; i++ {
			pods = append(pods, createPod(i, el))
		}
	}

	for i := 0; i < 3; i++ {
		rand.Shuffle(len(pods), func(i, j int) {
			pods[i], pods[j] = pods[j], pods[i]
		})
	}
	// assign nodes to pods
	for i, pod := range pods {
		pod.Spec.NodeName = fmt.Sprintf("node-%d", i/5) // Assign 5 pods per node
	}
	return pods
}

func TestPodsWithAnnotationsAndLabels(t *testing.T) {
	// total target pods 700
	// total random pods 3950
	pods := fixtureCreatePodsWithNodes([]podSpec{
		{
			namespace:   "dev",
			labels:      map[string]string{"app": "nginx", "env": "dev", "agent": "enabled"},
			annotations: map[string]string{"arch": "amd64"},
			totalTarget: 300,
			totalRandom: 700,
		},
		{
			namespace:   "prod",
			labels:      map[string]string{"app": "nginx", "env": "prod", "agent": "enabled"},
			annotations: map[string]string{"arch": "amd64"},
			totalTarget: 150,
			totalRandom: 2700,
		},
		{
			namespace:   "default",
			labels:      map[string]string{"app": "nginx", "agent": "disabled"},
			annotations: map[string]string{"arch": "amd64"},
			totalTarget: 250,
			totalRandom: 450,
		},
		{
			namespace:   "kube-system",
			labels:      map[string]string{},
			annotations: map[string]string{},
			totalTarget: 0,
			totalRandom: 100,
		},
	})

	client := fake.NewClientset()

	nodes := map[string]bool{}

	for _, pod := range pods {
		if _, exists := nodes[pod.Spec.NodeName]; !exists {
			nodes[pod.Spec.NodeName] = true
			node := &corev1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Name: pod.Spec.NodeName,
				},
			}
			if _, err := client.CoreV1().Nodes().Create(t.Context(), node, metav1.CreateOptions{}); err != nil {
				assert.NoError(t, err)
			}
		}
		if _, err := client.CoreV1().Pods(pod.Namespace).Create(t.Context(), pod, metav1.CreateOptions{}); err != nil {
			assert.NoError(t, err)
		}
	}

	tests := []struct {
		name                  string
		namespace             string
		labelSelector         string
		annotationFilter      string
		expectedEndpointCount int
	}{
		{
			name:                  "prod namespace with labels",
			namespace:             "prod",
			labelSelector:         "app=nginx",
			expectedEndpointCount: 150,
		},
		{
			name:                  "prod namespace with annotations",
			namespace:             "prod",
			annotationFilter:      "arch=amd64",
			expectedEndpointCount: 150,
		},
		{
			name:                  "prod namespace with annotations and labels not exists",
			namespace:             "prod",
			labelSelector:         "app=not-exists",
			annotationFilter:      "arch=amd64",
			expectedEndpointCount: 0,
		},
		{
			name:                  "all namespaces with correct annotations and labels",
			namespace:             "",
			labelSelector:         "app=nginx,agent=enabled",
			annotationFilter:      "arch=amd64",
			expectedEndpointCount: 450, // 300 from dev + 150 from prod
		},
		{
			name:                  "all namespaces with loose annotations and labels",
			namespace:             "",
			labelSelector:         "app=nginx",
			annotationFilter:      "arch=amd64",
			expectedEndpointCount: 700, // 300 from dev + 150 from prod + 250 from default
		},
		{
			name:                  "all namespaces with loose annotations and labels",
			namespace:             "",
			labelSelector:         "agent",
			annotationFilter:      "arch",
			expectedEndpointCount: 700,
		},
		{
			name:                  "all namespaces without filters",
			namespace:             "",
			labelSelector:         "",
			annotationFilter:      "",
			expectedEndpointCount: 4650,
		},
		{
			name:                  "single namespace without filters",
			namespace:             "default",
			labelSelector:         "",
			annotationFilter:      "",
			expectedEndpointCount: 700,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			selector, _ := annotations.ParseFilter(tt.labelSelector)
			pSource, err := NewPodSource(
				t.Context(), client,
				tt.namespace, "",
				false, "",
				"{{ .Name }}.tld.org", false,
				tt.annotationFilter, selector)
			require.NoError(t, err)

			endpoints, err := pSource.Endpoints(t.Context())
			require.NoError(t, err)

			assert.Len(t, endpoints, tt.expectedEndpointCount)
		})
	}
}
