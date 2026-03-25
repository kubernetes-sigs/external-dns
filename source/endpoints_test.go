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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes/fake"

	"sigs.k8s.io/external-dns/endpoint"
)

func TestEndpointTargetsFromServices(t *testing.T) {
	tests := []struct {
		name      string
		services  []*corev1.Service
		namespace string
		selector  map[string]string
		expected  endpoint.Targets
		wantErr   bool
	}{
		{
			name:      "no services",
			services:  []*corev1.Service{},
			namespace: "default",
			selector:  map[string]string{"app": "nginx"},
			expected:  endpoint.Targets{},
		},
		{
			name: "matching service with external IPs",
			services: []*corev1.Service{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "svc1",
						Namespace: "default",
					},
					Spec: corev1.ServiceSpec{
						Selector:    map[string]string{"app": "nginx"},
						ExternalIPs: []string{"192.0.2.1", "158.123.32.23"},
					},
				},
			},
			namespace: "default",
			selector:  map[string]string{"app": "nginx"},
			expected:  endpoint.Targets{"158.123.32.23", "192.0.2.1"},
		},
		{
			name: "matching service with duplicate external IPs",
			services: []*corev1.Service{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "svc1",
						Namespace: "default",
					},
					Spec: corev1.ServiceSpec{
						Selector:    map[string]string{"app": "nginx"},
						ExternalIPs: []string{"192.0.2.1", "192.0.2.1", "158.123.32.23"},
					},
				},
			},
			namespace: "default",
			selector:  map[string]string{"app": "nginx"},
			expected:  endpoint.Targets{"158.123.32.23", "192.0.2.1"},
		},
		{
			name: "no matching service as service without selector",
			services: []*corev1.Service{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "svc1",
						Namespace: "default",
					},
					Spec: corev1.ServiceSpec{
						ExternalIPs: []string{"192.0.2.1"},
					},
				},
			},
			namespace: "default",
			selector:  map[string]string{"app": "nginx"},
			expected:  endpoint.Targets{},
		},
		{
			name: "matching service with load balancer IP",
			services: []*corev1.Service{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "svc2",
						Namespace: "default",
					},
					Spec: corev1.ServiceSpec{
						Selector: map[string]string{"app": "nginx"},
					},
					Status: corev1.ServiceStatus{
						LoadBalancer: corev1.LoadBalancerStatus{
							Ingress: []corev1.LoadBalancerIngress{
								{IP: "192.0.2.2"},
							},
						},
					},
				},
			},
			namespace: "default",
			selector:  map[string]string{"app": "nginx"},
			expected:  endpoint.Targets{"192.0.2.2"},
		},
		{
			name: "matching service with load balancer hostname",
			services: []*corev1.Service{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "svc3",
						Namespace: "default",
					},
					Spec: corev1.ServiceSpec{
						Selector: map[string]string{"app": "nginx"},
					},
					Status: corev1.ServiceStatus{
						LoadBalancer: corev1.LoadBalancerStatus{
							Ingress: []corev1.LoadBalancerIngress{
								{Hostname: "lb.example.com"},
							},
						},
					},
				},
			},
			namespace: "default",
			selector:  map[string]string{"app": "nginx"},
			expected:  endpoint.Targets{"lb.example.com"},
		},
		{
			name: "no matching services",
			services: []*corev1.Service{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "svc4",
						Namespace: "default",
					},
					Spec: corev1.ServiceSpec{
						Selector: map[string]string{"app": "apache"},
					},
				},
			},
			namespace: "default",
			selector:  map[string]string{"app": "nginx"},
			expected:  endpoint.Targets{},
		},
		{
			name: "multiple selectors",
			services: []*corev1.Service{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "fake",
						Namespace: "default",
					},
					Spec: corev1.ServiceSpec{
						Selector:    map[string]string{"app": "apache", "version": "v1"},
						ExternalIPs: []string{"158.123.32.23"},
					},
				},
			},
			namespace: "default",
			selector:  map[string]string{"version": "v1"},
			expected:  endpoint.Targets{"158.123.32.23"},
		},
		{
			name: "complex selectors",
			services: []*corev1.Service{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "fake",
						Namespace: "default",
					},
					Spec: corev1.ServiceSpec{
						Selector: map[string]string{
							"app":     "demo",
							"env":     "prod",
							"team":    "devops",
							"version": "v1",
							"release": "stable",
							"track":   "daily",
							"tier":    "backend",
						},
						ExternalIPs: []string{"158.123.32.23"},
					},
				},
			},
			namespace: "default",
			selector: map[string]string{
				"version": "v1",
				"release": "stable",
				"tier":    "backend",
				"app":     "demo",
			},
			expected: endpoint.Targets{"158.123.32.23"},
		},
		{
			// Gateway selector is a SUPERSET of the service selector: the service is
			// missing a label the gateway requires. The index returns the service as a
			// candidate (it has the first queried k=v), but the label selector must
			// reject it because the remaining required label is absent.
			name: "gateway selector is superset of service selector — no match",
			services: []*corev1.Service{
				{
					ObjectMeta: metav1.ObjectMeta{Name: "igw", Namespace: "default"},
					Spec: corev1.ServiceSpec{
						Selector:    map[string]string{"istio": "ingressgateway"},
						ExternalIPs: []string{"10.0.0.1"},
					},
				},
			},
			namespace: "default",
			selector:  map[string]string{"istio": "ingressgateway", "app": "required"},
			expected:  endpoint.Targets{},
		},
		{
			// Reproduces the bug from PR #5708: the gateway selector is a strict subset of
			// the service's spec.selector. A hash-of-full-selector index would miss this.
			name: "gateway selector is subset of service selector",
			services: []*corev1.Service{
				{
					ObjectMeta: metav1.ObjectMeta{Name: "igw", Namespace: "default"},
					Spec: corev1.ServiceSpec{
						Selector:    map[string]string{"istio": "ingressgateway", "release": "istio"},
						ExternalIPs: []string{"10.0.0.1"},
					},
				},
			},
			namespace: "default",
			selector:  map[string]string{"istio": "ingressgateway"},
			expected:  endpoint.Targets{"10.0.0.1"},
		},
		{
			// Two services share the same first index entry ("istio=ingressgateway") but
			// only one satisfies the full gateway selector. Validates that the label selector
			// correctly eliminates the false positive returned by the index.
			name: "index returns multiple candidates, post-filter eliminates false positives",
			services: []*corev1.Service{
				{
					ObjectMeta: metav1.ObjectMeta{Name: "igw-a", Namespace: "default"},
					Spec: corev1.ServiceSpec{
						Selector:    map[string]string{"istio": "ingressgateway", "app": "foo"},
						ExternalIPs: []string{"10.0.0.1"},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{Name: "igw-b", Namespace: "default"},
					Spec: corev1.ServiceSpec{
						Selector:    map[string]string{"istio": "ingressgateway", "app": "bar"},
						ExternalIPs: []string{"10.0.0.2"},
					},
				},
			},
			namespace: "default",
			selector:  map[string]string{"istio": "ingressgateway", "app": "foo"},
			expected:  endpoint.Targets{"10.0.0.1"},
		},
		{
			// Empty gateway selector takes the lister path (no index key to query) and
			// returns all services in the namespace — same behaviour as an empty label selector matching everything.
			name: "empty selector returns all services",
			services: []*corev1.Service{
				{
					ObjectMeta: metav1.ObjectMeta{Name: "svc-a", Namespace: "default"},
					Spec: corev1.ServiceSpec{
						Selector:    map[string]string{"app": "foo"},
						ExternalIPs: []string{"10.0.0.1"},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{Name: "svc-b", Namespace: "default"},
					Spec: corev1.ServiceSpec{
						Selector:    map[string]string{"app": "bar"},
						ExternalIPs: []string{"10.0.0.2"},
					},
				},
			},
			namespace: "default",
			selector:  map[string]string{},
			expected:  endpoint.Targets{"10.0.0.1", "10.0.0.2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := fake.NewClientset()
			informerFactory := kubeinformers.NewSharedInformerFactoryWithOptions(client, 0,
				kubeinformers.WithNamespace(tt.namespace))
			serviceInformer := informerFactory.Core().V1().Services()

			for _, svc := range tt.services {
				_, err := client.CoreV1().Services(tt.namespace).Create(t.Context(), svc, metav1.CreateOptions{})
				assert.NoError(t, err)

				err = serviceInformer.Informer().GetIndexer().Add(svc)
				assert.NoError(t, err)
			}

			result, err := EndpointTargetsFromServices(serviceInformer, tt.namespace, tt.selector)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestEndpointTargetsFromServicesWithFixtures(t *testing.T) {
	svcInformer, err := svcInformerWithServices(2, 9)
	assert.NoError(t, err)

	sel := map[string]string{"app": "nginx", "env": "prod"}

	targets, err := EndpointTargetsFromServices(svcInformer, "default", sel)
	assert.NoError(t, err)
	assert.Equal(t, 2, targets.Len())
}
