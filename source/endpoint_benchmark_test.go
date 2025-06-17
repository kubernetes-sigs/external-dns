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
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubeinformers "k8s.io/client-go/informers"
	coreinformers "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes/fake"

	v1alpha3 "istio.io/api/networking/v1alpha3"
	istiov1a "istio.io/client-go/pkg/apis/networking/v1"

	"k8s.io/client-go/tools/cache"
)

func BenchmarkEndpointTargetsFromServicesMedium(b *testing.B) {
	svcInformer, err := svcInformerWithServices(36, 1000)
	assert.NoError(b, err)

	sel := map[string]string{"app": "nginx", "env": "prod"}

	for b.Loop() {
		targets, _ := EndpointTargetsFromServices(svcInformer, "default", sel)
		assert.Equal(b, 36, targets.Len())
	}
}

func BenchmarkEndpointTargetsFromServicesMediumIterateOverGateways(b *testing.B) {
	svcInformer, err := svcInformerWithServices(36, 500)
	assert.NoError(b, err)

	gateways := fixturesIstioGatewaySvcWithLabels(15, 70)

	for b.Loop() {
		for _, gateway := range gateways {
			_, _ = EndpointTargetsFromServices(svcInformer, gateway.Namespace, gateway.Spec.Selector)
		}
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

// This benchmark tests the performance of EndpointTargetsFromServices with a high number of services and gateways.
func BenchmarkEndpointTargetsFromServicesHighIterateOverGateways(b *testing.B) {
	svcInformer, err := svcInformerWithServices(36, 40000)
	assert.NoError(b, err)

	gateways := fixturesIstioGatewaySvcWithLabels(50, 1000)

	for b.Loop() {
		for _, gateway := range gateways {
			_, _ = EndpointTargetsFromServices(svcInformer, gateway.Namespace, gateway.Spec.Selector)
		}
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
	informerFactory.Start(stopCh)
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

// fixturesIstioGatewaySvcWithLabels creates a list of Services for testing purposes.
// It generates a specified number of gateways with static labels and random labels.
// The first `toLookup` services have specific labels, while the next `underTest` services have random labels.
func fixturesIstioGatewaySvcWithLabels(toLookup, underTest int) []*istiov1a.Gateway {
	var result []*istiov1a.Gateway

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

	var createGateway = func(name string, namespace string, selector map[string]string) *istiov1a.Gateway {
		return &istiov1a.Gateway{
			ObjectMeta: metav1.ObjectMeta{
				Name:      name,
				Namespace: namespace,
			},
			Spec: v1alpha3.Gateway{
				Selector: selector,
				Servers: []*v1alpha3.Server{
					{
						Port:  &v1alpha3.Port{},
						Hosts: []string{"*"},
					},
				},
			},
		}
	}
	// services with specific labels
	for i := 0; i < toLookup; i++ {
		svc := createGateway("istio-gw-"+strconv.Itoa(i), "default", map[string]string{"app": "nginx", "env": "prod"})
		result = append(result, svc)
	}

	// services with random labels
	for i := 0; i < underTest; i++ {
		svc := createGateway("istio-random-svc-"+strconv.Itoa(i), "default", randomLabels(i))
		result = append(result, svc)
	}

	// Shuffle the services to ensure randomness
	for i := 0; i < 3; i++ {
		rand.Shuffle(len(result), func(i, j int) {
			result[i], result[j] = result[j], result[i]
		})
	}

	return result
}
