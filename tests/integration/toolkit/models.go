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
	corev1 "k8s.io/api/core/v1"
	discoveryv1 "k8s.io/api/discovery/v1"
	networkingv1 "k8s.io/api/networking/v1"
	k8sruntime "k8s.io/apimachinery/pkg/runtime"

	"sigs.k8s.io/external-dns/endpoint"
)

// TestScenarios represents the root structure of the YAML file.
type TestScenarios struct {
	Scenarios []Scenario `json:"scenarios"`
}

// Scenario represents a single test scenario.
type Scenario struct {
	Name      string                     `json:"name"`
	Config    ScenarioConfig             `json:"config"`
	Resources []ResourceWithDependencies `json:"resources"`
	Expected  []*endpoint.Endpoint       `json:"expected"`
}

// ResourceWithDependencies wraps a K8s resource with optional dependencies.
type ResourceWithDependencies struct {
	Resource     k8sruntime.RawExtension `json:"resource"`
	Dependencies *ResourceDependencies   `json:"dependencies,omitempty"`
}

// ResourceDependencies defines auto-generated dependent resources.
type ResourceDependencies struct {
	Pods *PodDependencies `json:"pods,omitempty"`
}

// PodDependencies defines how to generate Pods and EndpointSlices for a Service.
type PodDependencies struct {
	Replicas int `json:"replicas"`
}

// ScenarioConfig holds the wrapper configuration for a scenario.
type ScenarioConfig struct {
	Sources             []string `json:"sources"`
	DefaultTargets      []string `json:"defaultTargets"`
	ForceDefaultTargets bool     `json:"forceDefaultTargets"`
	TargetNetFilter     []string `json:"targetNetFilter"`
	ServiceTypeFilter   []string `json:"serviceTypeFilter"`
}

// ParsedResources holds the parsed Kubernetes resources from a scenario.
type ParsedResources struct {
	Ingresses      []*networkingv1.Ingress
	Services       []*corev1.Service
	EndpointSlices []*discoveryv1.EndpointSlice
	Pods           []*corev1.Pod
}
