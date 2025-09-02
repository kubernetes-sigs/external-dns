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
	"github.com/stretchr/testify/mock"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	corev1lister "k8s.io/client-go/listers/core/v1"
	discoveryv1lister "k8s.io/client-go/listers/discovery/v1"
	"k8s.io/client-go/tools/cache"
)

type FakeServiceInformer struct {
	mock.Mock
}

func (f *FakeServiceInformer) Informer() cache.SharedIndexInformer {
	args := f.Called()
	return args.Get(0).(cache.SharedIndexInformer)
}

func (f *FakeServiceInformer) Lister() corev1lister.ServiceLister {
	return corev1lister.NewServiceLister(f.Informer().GetIndexer())
}

type FakeEndpointSliceInformer struct {
	mock.Mock
}

func (f *FakeEndpointSliceInformer) Informer() cache.SharedIndexInformer {
	args := f.Called()
	return args.Get(0).(cache.SharedIndexInformer)
}

func (f *FakeEndpointSliceInformer) Lister() discoveryv1lister.EndpointSliceLister {
	return discoveryv1lister.NewEndpointSliceLister(f.Informer().GetIndexer())
}

type FakeNodeInformer struct {
	mock.Mock
}

func (f *FakeNodeInformer) Informer() cache.SharedIndexInformer {
	args := f.Called()
	return args.Get(0).(cache.SharedIndexInformer)
}

func (f *FakeNodeInformer) Lister() corev1lister.NodeLister {
	return corev1lister.NewNodeLister(f.Informer().GetIndexer())
}

func fakeService() *corev1.Service {
	return &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:        "fake-service",
			Namespace:   "ns",
			Labels:      map[string]string{"env": "prod", "team": "devops"},
			Annotations: map[string]string{"description": "Enriched service object"},
			UID:         "1234",
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
