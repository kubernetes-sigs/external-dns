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
	"context"
	"encoding/binary"
	"fmt"
	"math/rand/v2"
	"net"

	// "net"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubeinformers "k8s.io/client-go/informers"
	coreinformers "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/tools/cache"
	"sigs.k8s.io/external-dns/endpoint"
)

func TestEndpointsForHostname(t *testing.T) {
	tests := []struct {
		name             string
		hostname         string
		targets          endpoint.Targets
		ttl              endpoint.TTL
		providerSpecific endpoint.ProviderSpecific
		setIdentifier    string
		resource         string
		expected         []*endpoint.Endpoint
	}{
		{
			name:     "A record targets",
			hostname: "example.com",
			targets:  endpoint.Targets{"192.0.2.1", "192.0.2.2"},
			ttl:      endpoint.TTL(300),
			providerSpecific: endpoint.ProviderSpecific{
				{Name: "provider", Value: "value"},
			},
			setIdentifier: "identifier",
			resource:      "resource",
			expected: []*endpoint.Endpoint{
				{
					DNSName:          "example.com",
					Targets:          endpoint.Targets{"192.0.2.1", "192.0.2.2"},
					RecordType:       endpoint.RecordTypeA,
					RecordTTL:        endpoint.TTL(300),
					ProviderSpecific: endpoint.ProviderSpecific{{Name: "provider", Value: "value"}},
					SetIdentifier:    "identifier",
					Labels:           map[string]string{endpoint.ResourceLabelKey: "resource"},
				},
			},
		},
		{
			name:     "AAAA record targets",
			hostname: "example.com",
			targets:  endpoint.Targets{"2001:db8::1", "2001:db8::2"},
			ttl:      endpoint.TTL(300),
			providerSpecific: endpoint.ProviderSpecific{
				{Name: "provider", Value: "value"},
			},
			setIdentifier: "identifier",
			resource:      "resource",
			expected: []*endpoint.Endpoint{
				{
					DNSName:          "example.com",
					Targets:          endpoint.Targets{"2001:db8::1", "2001:db8::2"},
					RecordType:       endpoint.RecordTypeAAAA,
					RecordTTL:        endpoint.TTL(300),
					ProviderSpecific: endpoint.ProviderSpecific{{Name: "provider", Value: "value"}},
					SetIdentifier:    "identifier",
					Labels:           map[string]string{endpoint.ResourceLabelKey: "resource"},
				},
			},
		},
		{
			name:     "CNAME record targets",
			hostname: "example.com",
			targets:  endpoint.Targets{"cname.example.com"},
			ttl:      endpoint.TTL(300),
			providerSpecific: endpoint.ProviderSpecific{
				{Name: "provider", Value: "value"},
			},
			setIdentifier: "identifier",
			resource:      "resource",
			expected: []*endpoint.Endpoint{
				{
					DNSName:          "example.com",
					Targets:          endpoint.Targets{"cname.example.com"},
					RecordType:       endpoint.RecordTypeCNAME,
					RecordTTL:        endpoint.TTL(300),
					ProviderSpecific: endpoint.ProviderSpecific{{Name: "provider", Value: "value"}},
					SetIdentifier:    "identifier",
					Labels:           map[string]string{endpoint.ResourceLabelKey: "resource"},
				},
			},
		},
		{
			name:             "No targets",
			hostname:         "example.com",
			targets:          endpoint.Targets{},
			ttl:              endpoint.TTL(300),
			providerSpecific: endpoint.ProviderSpecific{},
			setIdentifier:    "",
			resource:         "",
			expected:         []*endpoint.Endpoint(nil),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := endpointsForHostname(tt.hostname, tt.targets, tt.ttl, tt.providerSpecific, tt.setIdentifier, tt.resource)
			assert.Equal(t, tt.expected, result)
		})
	}
}

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
			expected:  endpoint.Targets{"192.0.2.1", "158.123.32.23"},
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := fake.NewClientset()
			informerFactory := kubeinformers.NewSharedInformerFactoryWithOptions(client, 0,
				kubeinformers.WithNamespace(tt.namespace))
			serviceInformer := informerFactory.Core().V1().Services()

			for _, svc := range tt.services {
				_, err := client.CoreV1().Services(tt.namespace).Create(context.Background(), svc, metav1.CreateOptions{})
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

func TestEndpointTargetsFromServicesWithFixtures(b *testing.T) {
	svcInformer, err := svcInformerWithServices(2, 9)
	assert.NoError(b, err)

	sel := map[string]string{"app": "nginx", "env": "prod"}

	targets, err := EndpointTargetsFromServices(svcInformer, "default", sel)
	assert.NoError(b, err)
	assert.Equal(b, 2, targets.Len())
}

func BenchmarkEndpointTargetsFromServicesMedium(b *testing.B) {
	svcInformer, err := svcInformerWithServices(36, 1000)
	assert.NoError(b, err)

	sel := map[string]string{"app": "nginx", "env": "prod"}

	for b.Loop() {
		targets, _ := EndpointTargetsFromServices(svcInformer, "default", sel)
		assert.Equal(b, 36, targets.Len())
	}
}

func BenchmarkEndpointTargetsFromServicesHigh(b *testing.B) {
	svcInformer, err := svcInformerWithServices(36, 40000)
	assert.NoError(b, err)

	sel := map[string]string{"app": "nginx", "env": "prod"}

	for b.Loop() {
		targets, _ := EndpointTargetsFromServices(svcInformer, "default", sel)
		assert.Equal(b, 36, targets.Len())
	}
}

// helperToPopulateFakeClientWithServices populates a fake Kubernetes client with a specified services.
func svcInformerWithServices(toLookup, underTest int) (coreinformers.ServiceInformer, error) {
	client := fake.NewClientset()
	informerFactory := kubeinformers.NewSharedInformerFactoryWithOptions(client, 0, kubeinformers.WithNamespace("default"))
	svcInformer := informerFactory.Core().V1().Services()
	ctx := context.Background()

	_, err := svcInformer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
			},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to add event handler: %w", err)
	}

	services := fixturesSvcWithLabels(toLookup, underTest)
	for _, svc := range services {
		_, err := client.CoreV1().Services(svc.Namespace).Create(ctx, svc, metav1.CreateOptions{})
		if err != nil {
			return nil, fmt.Errorf("failed to create service %s: %w", svc.Name, err)
		}
	}

	stopCh := make(chan struct{})
	defer close(stopCh)
	go informerFactory.Start(stopCh)
	cache.WaitForCacheSync(stopCh, svcInformer.Informer().HasSynced)
	return svcInformer, nil
}

// fixturesSvcWithLabels creates a list of Services for testing purposes.
// It generates a specified number of services with static labels and random labels.
// The first `toLookup` services have specific labels, while the next `underTest` services have random labels.
func fixturesSvcWithLabels(toLookup, underTest int) []*corev1.Service {
	var services []*corev1.Service

	var randomLabels = func(input int) map[string]string {
		if input%3 == 0 {
			// every third service has no labels
			return map[string]string{}
		}
		return map[string]string{
			"app":                                fmt.Sprintf("service-%d", rand.IntN(100)),
			fmt.Sprintf("key%d", rand.IntN(100)): fmt.Sprintf("value%d", rand.IntN(100)),
		}
	}

	var randomIPs = func() []string {
		ip := rand.Uint32()
		buf := make([]byte, 4)
		binary.LittleEndian.PutUint32(buf, ip)
		return []string{net.IP(buf).String()}
	}

	var createService = func(name string, namespace string, selector map[string]string) *corev1.Service {
		return &corev1.Service{
			ObjectMeta: metav1.ObjectMeta{
				Name:      name,
				Namespace: namespace,
			},
			Spec: corev1.ServiceSpec{
				Selector:    selector,
				ExternalIPs: randomIPs(),
			},
		}
	}

	// services with specific labels
	for i := 0; i < toLookup; i++ {
		svc := createService("nginx-svc-"+strconv.Itoa(i), "default", map[string]string{"app": "nginx", "env": "prod"})
		services = append(services, svc)
	}

	// services with random labels
	for i := 0; i < underTest; i++ {
		svc := createService("random-svc-"+strconv.Itoa(i), "default", randomLabels(i))
		services = append(services, svc)
	}

	// Shuffle the services to ensure randomness
	for i := 0; i < 3; i++ {
		rand.Shuffle(len(services), func(i, j int) {
			services[i], services[j] = services[j], services[i]
		})
	}

	return services
}
