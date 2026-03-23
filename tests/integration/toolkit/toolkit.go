/*
Copyright 2026 The Kubernetes Authors.

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

package toolkit

import (
	"context"
	"fmt"
	"maps"

	corev1 "k8s.io/api/core/v1"
	discoveryv1 "k8s.io/api/discovery/v1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/kubernetes/fake"
	"sigs.k8s.io/yaml"

	"sigs.k8s.io/external-dns/pkg/apis/externaldns"

	"sigs.k8s.io/external-dns/internal/testutils"

	"sigs.k8s.io/external-dns/source"
	"sigs.k8s.io/external-dns/source/wrappers"
)

// Initialized at package load; safe for concurrent use after that.
var (
	scheme = func() *runtime.Scheme {
		s := runtime.NewScheme()
		utilruntime.Must(corev1.AddToScheme(s))
		utilruntime.Must(discoveryv1.AddToScheme(s))
		utilruntime.Must(networkingv1.AddToScheme(s))
		return s
	}()
	decoder = serializer.NewCodecFactory(scheme).UniversalDeserializer()
)

// LoadScenarios loads test scenarios from the embedded YAML data.
func LoadScenarios(data []byte) (*TestScenarios, error) {
	var scenarios TestScenarios
	if err := yaml.Unmarshal(data, &scenarios); err != nil {
		return nil, err
	}

	// Validate scenarios
	for i, s := range scenarios.Scenarios {
		if s.Name == "" {
			return nil, fmt.Errorf("scenario %d is missing required field: name", i)
		}
		if s.Description == "" {
			return nil, fmt.Errorf("scenario %d (%q) is missing required field: description", i, s.Name)
		}
		if len(s.Config.Sources) == 0 {
			return nil, fmt.Errorf("scenario %d (%q) is missing required field: config.sources", i, s.Name)
		}
	}

	return &scenarios, nil
}

// ParseResources parses the raw resources from a scenario into typed objects.
func ParseResources(resources []ResourceWithDependencies) (*ParsedResources, error) {
	parsed := &ParsedResources{}

	for _, item := range resources {
		obj, _, err := decoder.Decode(item.Resource.Raw, nil, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to decode resource: %w", err)
		}

		switch res := obj.(type) {
		case *corev1.Pod:
			parsed.Pods = append(parsed.Pods, res)
		case *corev1.Service:
			parsed.Services = append(parsed.Services, res)
			// Auto-generate Pods and EndpointSlice if dependencies are specified
			if item.Dependencies != nil && item.Dependencies.Pods != nil {
				pods, endpointSlice := generatePodsAndEndpointSlice(res, item.Dependencies.Pods)
				parsed.Pods = append(parsed.Pods, pods...)
				parsed.EndpointSlices = append(parsed.EndpointSlices, endpointSlice)
			}
		case *networkingv1.Ingress:
			parsed.Ingresses = append(parsed.Ingresses, res)
		case *discoveryv1.EndpointSlice:
			parsed.EndpointSlices = append(parsed.EndpointSlices, res)
		default:
			return nil, fmt.Errorf("unsupported resource type %T", obj)
		}
	}

	return parsed, nil
}

// generatePodsAndEndpointSlice creates Pods and an EndpointSlice for a headless service.
func generatePodsAndEndpointSlice(svc *corev1.Service, deps *PodDependencies) ([]*corev1.Pod, *discoveryv1.EndpointSlice) {
	var pods []*corev1.Pod
	var endpoints []discoveryv1.Endpoint

	for i := range deps.Replicas {
		podName := fmt.Sprintf("%s-%d", svc.Name, i)
		podIP := fmt.Sprintf("10.0.0.%d", i+1)

		pod := &corev1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:      podName,
				Namespace: svc.Namespace,
				Labels:    svc.Spec.Selector,
			},
			Spec: corev1.PodSpec{
				Hostname: podName,
			},
			Status: corev1.PodStatus{
				PodIP: podIP,
			},
		}
		pods = append(pods, pod)

		endpoints = append(endpoints, discoveryv1.Endpoint{
			Addresses: []string{podIP},
			TargetRef: &corev1.ObjectReference{
				Kind: "Pod",
				Name: podName,
			},
			Conditions: discoveryv1.EndpointConditions{
				Ready: testutils.ToPtr(true),
			},
		})
	}

	// Create EndpointSlice with the service name label
	endpointSliceLabels := maps.Clone(svc.Spec.Selector)
	endpointSliceLabels[discoveryv1.LabelServiceName] = svc.Name

	endpointSlice := &discoveryv1.EndpointSlice{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("%s-slice", svc.Name),
			Namespace: svc.Namespace,
			Labels:    endpointSliceLabels,
		},
		AddressType: discoveryv1.AddressTypeIPv4,
		Endpoints:   endpoints,
	}

	return pods, endpointSlice
}

func createIngressWithOptionalStatus(ctx context.Context, client *fake.Clientset, ing *networkingv1.Ingress) error {
	created, err := client.NetworkingV1().Ingresses(ing.Namespace).Create(ctx, ing, metav1.CreateOptions{})
	if err != nil {
		return err
	}

	// Update status separately since Create doesn't set status in the fake client.
	if len(ing.Status.LoadBalancer.Ingress) > 0 {
		created.Status = ing.Status
		_, err = client.NetworkingV1().Ingresses(ing.Namespace).UpdateStatus(ctx, created, metav1.UpdateOptions{})
		if err != nil {
			return err
		}
	}

	return nil
}

func createServiceWithOptionalStatus(ctx context.Context, client *fake.Clientset, svc *corev1.Service) error {
	created, err := client.CoreV1().Services(svc.Namespace).Create(ctx, svc, metav1.CreateOptions{})
	if err != nil {
		return err
	}

	// Update status separately since Create doesn't set status in the fake client.
	if len(svc.Status.LoadBalancer.Ingress) > 0 {
		created.Status = svc.Status
		_, err = client.CoreV1().Services(svc.Namespace).UpdateStatus(ctx, created, metav1.UpdateOptions{})
		if err != nil {
			return err
		}
	}

	return nil
}

// LoadResources creates the resources in the fake client using the API.
// This must be called BEFORE creating sources so the informers can see the resources.
func LoadResources(ctx context.Context, scenario Scenario) (*fake.Clientset, error) {
	client := fake.NewClientset()

	// Parse resources from scenario
	resources, err := ParseResources(scenario.Resources)
	if err != nil {
		return nil, err
	}

	for _, ing := range resources.Ingresses {
		if err := createIngressWithOptionalStatus(ctx, client, ing); err != nil {
			return nil, err
		}
	}
	for _, svc := range resources.Services {
		if err := createServiceWithOptionalStatus(ctx, client, svc); err != nil {
			return nil, err
		}
	}
	for _, pod := range resources.Pods {
		_, err := client.CoreV1().Pods(pod.Namespace).Create(ctx, pod, metav1.CreateOptions{})
		if err != nil {
			return nil, err
		}
	}
	for _, eps := range resources.EndpointSlices {
		_, err := client.DiscoveryV1().EndpointSlices(eps.Namespace).Create(ctx, eps, metav1.CreateOptions{})
		if err != nil {
			return nil, err
		}
	}
	return client, nil
}

// scenarioToConfig creates a source.Config for testing with the scenario config.
func scenarioToConfig(scenarioCfg ScenarioConfig, opts ...source.OverrideConfigOption) *source.Config {
	return source.NewSourceConfig(&externaldns.Config{
		Sources:             scenarioCfg.Sources,
		ServiceTypeFilter:   scenarioCfg.ServiceTypeFilter,
		DefaultTargets:      scenarioCfg.DefaultTargets,
		ForceDefaultTargets: scenarioCfg.ForceDefaultTargets,
		TargetNetFilter:     scenarioCfg.TargetNetFilter,
		ExcludeTargetNets:   scenarioCfg.ExcludeTargetNets,
		NAT64Networks:       scenarioCfg.NAT64Networks,
		Provider:            scenarioCfg.Provider,
		PreferAlias:         scenarioCfg.PreferAlias,
	}, opts...)
}

// CreateWrappedSource builds all named sources using a mock client and wraps
// them with the same pipeline used by the controller.
func CreateWrappedSource(
	ctx context.Context,
	client *fake.Clientset,
	scenarioCfg ScenarioConfig) (source.Source, error) {
	cfg := scenarioToConfig(scenarioCfg, source.WithClientGenerator(newMockClientGenerator(client)))
	return wrappers.Build(ctx, cfg)
}
