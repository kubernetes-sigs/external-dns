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

package informers

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func fakeService() *corev1.Service {
	return &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "fake-service",
			Namespace: "ns",
			Labels:    map[string]string{"env": "prod", "team": "devops"},
			Annotations: map[string]string{
				"description":                               "some annotation",
				corev1.LastAppliedConfigAnnotation:          `{"apiVersion":"v1","kind":"Service"}`,
				"external-dns.alpha.kubernetes.io/hostname": "example.com",
			},
			UID: "1234",
			ManagedFields: []metav1.ManagedFieldsEntry{
				{Manager: "kubectl", Operation: metav1.ManagedFieldsOperationApply},
			},
		},
		Spec: corev1.ServiceSpec{
			Selector:    map[string]string{"app": "demo"},
			ExternalIPs: []string{"1.2.3.4"},
			Ports: []corev1.ServicePort{
				{
					Name:       "http",
					Port:       80,
					TargetPort: intstr.FromInt32(8080),
					Protocol:   corev1.ProtocolTCP,
				},
				{
					Name:       "https",
					Port:       443,
					TargetPort: intstr.FromInt32(8443),
					Protocol:   corev1.ProtocolTCP,
				},
			},
			Type: corev1.ServiceTypeLoadBalancer,
		},
		Status: corev1.ServiceStatus{
			LoadBalancer: corev1.LoadBalancerStatus{
				Ingress: []corev1.LoadBalancerIngress{
					{IP: "5.6.7.8", Hostname: "lb.example.com"},
				},
			},
			Conditions: []metav1.Condition{
				{
					Type:               "Available",
					Status:             metav1.ConditionTrue,
					Reason:             "MinimumReplicasAvailable",
					Message:            "Service is available",
					LastTransitionTime: metav1.Now(),
				},
			},
		},
	}
}

func fakeNode() *corev1.Node {
	return &corev1.Node{
		ObjectMeta: metav1.ObjectMeta{
			Name:   "fake-node",
			Labels: map[string]string{"region": "us-east-1"},
			Annotations: map[string]string{
				corev1.LastAppliedConfigAnnotation:          `{"apiVersion":"v1","kind":"Node"}`,
				"external-dns.alpha.kubernetes.io/hostname": "node.example.com",
			},
			UID: "9012",
			ManagedFields: []metav1.ManagedFieldsEntry{
				{Manager: "kubectl", Operation: metav1.ManagedFieldsOperationApply},
			},
		},
		Status: corev1.NodeStatus{
			Addresses: []corev1.NodeAddress{
				{Type: corev1.NodeExternalIP, Address: "1.2.3.4"},
			},
			Conditions: []corev1.NodeCondition{
				{Type: corev1.NodeReady, Status: corev1.ConditionTrue},
				{Type: corev1.NodeMemoryPressure, Status: corev1.ConditionFalse},
			},
		},
	}
}

func fakePod() *corev1.Pod {
	return &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "fake-pod",
			Namespace: "ns",
			Labels:    map[string]string{"app": "demo"},
			Annotations: map[string]string{
				corev1.LastAppliedConfigAnnotation:          `{"apiVersion":"v1","kind":"Pod"}`,
				"external-dns.alpha.kubernetes.io/hostname": "pod.example.com",
				"unrelated.io/annotation":                   "should-be-dropped",
			},
			UID: "5678",
			ManagedFields: []metav1.ManagedFieldsEntry{
				{Manager: "kubectl", Operation: metav1.ManagedFieldsOperationApply},
			},
		},
		Spec: corev1.PodSpec{
			NodeName: "node-1",
		},
		Status: corev1.PodStatus{
			PodIP: "10.0.0.1",
		},
	}
}
