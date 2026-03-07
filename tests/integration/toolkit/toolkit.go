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
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"testing"

	corev1 "k8s.io/api/core/v1"
	discoveryv1 "k8s.io/api/discovery/v1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	"sigs.k8s.io/yaml"

	"sigs.k8s.io/external-dns/pkg/apis/externaldns"

	"sigs.k8s.io/external-dns/internal/testutils"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/source"
	"sigs.k8s.io/external-dns/source/wrappers"
)

// LoadScenarios loads test scenarios from the YAML file.
func LoadScenarios(dir string) (*TestScenarios, error) {
	filename := filepath.Join(dir, "scenarios", "tests.yaml")
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var scenarios TestScenarios
	if err := yaml.Unmarshal(data, &scenarios); err != nil {
		return nil, err
	}

	// Validate scenarios
	for i, s := range scenarios.Scenarios {
		if s.Description == "" {
			return nil, fmt.Errorf("scenario %d (%q) is missing required field: description", i, s.Name)
		}
	}

	return &scenarios, nil
}

// ParseResources parses the raw resources from a scenario into typed objects.
func ParseResources(resources []ResourceWithDependencies) (*ParsedResources, error) {
	parsed := &ParsedResources{}

	for _, item := range resources {
		raw := item.Resource

		// First unmarshal to get the kind
		var typeMeta metav1.TypeMeta
		if err := yaml.Unmarshal(raw.Raw, &typeMeta); err != nil {
			return nil, err
		}

		switch typeMeta.Kind {
		case "Ingress":
			var ingress networkingv1.Ingress
			if err := yaml.Unmarshal(raw.Raw, &ingress); err != nil {
				return nil, err
			}
			parsed.Ingresses = append(parsed.Ingresses, &ingress)
		case "Service":
			var svc corev1.Service
			if err := yaml.Unmarshal(raw.Raw, &svc); err != nil {
				return nil, err
			}
			parsed.Services = append(parsed.Services, &svc)

			// Auto-generate Pods and EndpointSlice if dependencies are specified
			if item.Dependencies != nil && item.Dependencies.Pods != nil {
				pods, endpointSlice := generatePodsAndEndpointSlice(&svc, item.Dependencies.Pods)
				parsed.Pods = append(parsed.Pods, pods...)
				parsed.EndpointSlices = append(parsed.EndpointSlices, endpointSlice)
			}
		case "EndpointSlice":
			var eps discoveryv1.EndpointSlice
			if err := yaml.Unmarshal(raw.Raw, &eps); err != nil {
				return nil, err
			}
			parsed.EndpointSlices = append(parsed.EndpointSlices, &eps)
		case "Pod":
			var pod corev1.Pod
			if err := yaml.Unmarshal(raw.Raw, &pod); err != nil {
				return nil, err
			}
			parsed.Pods = append(parsed.Pods, &pod)
		}
	}

	return parsed, nil
}

// generatePodsAndEndpointSlice creates Pods and an EndpointSlice for a headless service.
func generatePodsAndEndpointSlice(svc *corev1.Service, deps *PodDependencies) ([]*corev1.Pod, *discoveryv1.EndpointSlice) {
	var pods []*corev1.Pod
	var endpoints []discoveryv1.Endpoint

	for i := 0; i < deps.Replicas; i++ {
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
	endpointSliceLabels := make(map[string]string)
	maps.Copy(endpointSliceLabels, svc.Spec.Selector)
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
func scenarioToConfig(scenarioCfg ScenarioConfig) *source.Config {
	return source.NewSourceConfig(&externaldns.Config{
		Sources:             scenarioCfg.Sources,
		ServiceTypeFilter:   scenarioCfg.ServiceTypeFilter,
		DefaultTargets:      scenarioCfg.DefaultTargets,
		ForceDefaultTargets: scenarioCfg.ForceDefaultTargets,
		TargetNetFilter:     scenarioCfg.TargetNetFilter,
	})
}

// CreateWrappedSource creates sources using source.BuildWithConfig and wraps them with wrappers.WrapSources.
// TODO: could we reuse the same source.BuildWithConfig() code as the controller instead of duplicating it here? It would require refactoring to allow passing in a custom client generator, but it would ensure we're testing the same code as the controller.
func CreateWrappedSource(
	ctx context.Context,
	client *fake.Clientset,
	scenarioCfg ScenarioConfig) (source.Source, error) {
	clientGen := newMockClientGenerator(client)
	cfg := scenarioToConfig(scenarioCfg)

	// TODO: copied from controller/execute.go#buildSources
	sources, err := source.ByNames(ctx, cfg, clientGen)
	if err != nil {
		return nil, err
	}
	opts := wrappers.NewConfig(
		wrappers.WithDefaultTargets(cfg.DefaultTargets),
		wrappers.WithForceDefaultTargets(cfg.ForceDefaultTargets),
		wrappers.WithNAT64Networks(cfg.NAT64Networks),
		wrappers.WithTargetNetFilter(cfg.TargetNetFilter),
		wrappers.WithExcludeTargetNets(cfg.ExcludeTargetNets),
		wrappers.WithMinTTL(cfg.MinTTL))

	return wrappers.WrapSources(sources, opts)
}

// TODO: copied from source/wrappers/source_test.go - unify in following PR
func ValidateEndpoints(t *testing.T, endpoints, expected []*endpoint.Endpoint) {
	t.Helper()

	if len(endpoints) != len(expected) {
		t.Fatalf("expected %d endpoints, got %d", len(expected), len(endpoints))
	}

	// Make sure endpoints are sorted - validateEndpoint() depends on it.
	sortEndpoints(endpoints)
	sortEndpoints(expected)

	for i := range endpoints {
		validateEndpoint(t, endpoints[i], expected[i])
	}
}

// TODO: copied from source/wrappers/source_test.go - unify in following PR
func validateEndpoint(t *testing.T, endpoint, expected *endpoint.Endpoint) {
	t.Helper()

	if endpoint.DNSName != expected.DNSName {
		t.Errorf("DNSName expected %q, got %q", expected.DNSName, endpoint.DNSName)
	}

	if !endpoint.Targets.Same(expected.Targets) {
		t.Errorf("Targets expected %q, got %q", expected.Targets, endpoint.Targets)
	}

	if endpoint.RecordTTL != expected.RecordTTL {
		t.Errorf("RecordTTL expected %v, got %v", expected.RecordTTL, endpoint.RecordTTL)
	}

	// if a non-empty record type is expected, check that it matches.
	if endpoint.RecordType != expected.RecordType {
		t.Errorf("RecordType expected %q, got %q", expected.RecordType, endpoint.RecordType)
	}

	// if non-empty labels are expected, check that they match.
	if expected.Labels != nil && !reflect.DeepEqual(endpoint.Labels, expected.Labels) {
		t.Errorf("Labels expected %s, got %s", expected.Labels, endpoint.Labels)
	}

	if (len(expected.ProviderSpecific) != 0 || len(endpoint.ProviderSpecific) != 0) &&
		!reflect.DeepEqual(endpoint.ProviderSpecific, expected.ProviderSpecific) {
		t.Errorf("ProviderSpecific expected %s, got %s", expected.ProviderSpecific, endpoint.ProviderSpecific)
	}

	if endpoint.SetIdentifier != expected.SetIdentifier {
		t.Errorf("SetIdentifier expected %q, got %q", expected.SetIdentifier, endpoint.SetIdentifier)
	}
}

// TODO: copied from source/wrappers/source_test.go - unify in following PR
func sortEndpoints(endpoints []*endpoint.Endpoint) {
	for _, ep := range endpoints {
		sort.Strings(ep.Targets)
	}
	sort.Slice(endpoints, func(i, k int) bool {
		// Sort by DNSName, RecordType, and Targets
		ei, ek := endpoints[i], endpoints[k]
		if ei.DNSName != ek.DNSName {
			return ei.DNSName < ek.DNSName
		}
		if ei.RecordType != ek.RecordType {
			return ei.RecordType < ek.RecordType
		}
		// Targets are sorted ahead of time.
		for j, ti := range ei.Targets {
			if j >= len(ek.Targets) {
				return true
			}
			if tk := ek.Targets[j]; ti != tk {
				return ti < tk
			}
		}
		return false
	})
}
