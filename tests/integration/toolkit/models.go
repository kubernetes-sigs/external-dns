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
	"k8s.io/client-go/kubernetes/fake"

	apiv1alpha1 "sigs.k8s.io/external-dns/apis/v1alpha1"

	"sigs.k8s.io/external-dns/endpoint"
)

// TestScenarios represents the root structure of the YAML file.
type TestScenarios struct {
	Scenarios []Scenario `json:"scenarios"`
}

// ExpectedRefObject describes an expected Kubernetes object reference on an endpoint.
// Key is matched against ObjectReference.Key() ("source/namespace/name").
type ExpectedRefObject struct {
	Key string `json:"key"`
}

// ExpectedEndpoint mirrors the JSON-serialisable fields of endpoint.Endpoint and
// adds an optional RefObjects list for asserting event-source attribution.
type ExpectedEndpoint struct {
	DNSName          string                    `json:"dnsName,omitempty"`
	Targets          endpoint.Targets          `json:"targets,omitempty"`
	RecordType       string                    `json:"recordType,omitempty"`
	SetIdentifier    string                    `json:"setIdentifier,omitempty"`
	RecordTTL        endpoint.TTL              `json:"recordTTL,omitempty"`
	Labels           endpoint.Labels           `json:"labels,omitempty"`
	ProviderSpecific endpoint.ProviderSpecific `json:"providerSpecific,omitempty"`
	// RefObjects is optional. When non-empty, the test asserts that the actual
	// endpoint has exactly this many ref objects and that each one matches the
	// corresponding ExpectedRefObject (partial match: only non-empty fields checked).
	RefObjects []ExpectedRefObject `json:"refObjects,omitempty"`
}

// ToEndpoint converts an ExpectedEndpoint to a plain *endpoint.Endpoint for use
// with the standard field validation helpers.
func (e *ExpectedEndpoint) ToEndpoint() *endpoint.Endpoint {
	ep := &endpoint.Endpoint{
		DNSName:          e.DNSName,
		Targets:          e.Targets,
		RecordType:       e.RecordType,
		SetIdentifier:    e.SetIdentifier,
		RecordTTL:        e.RecordTTL,
		Labels:           e.Labels,
		ProviderSpecific: e.ProviderSpecific,
	}
	return ep
}

// Scenario represents a single test scenario.
type Scenario struct {
	Name        string                     `json:"name"`
	Description string                     `json:"description"`
	Config      ScenarioConfig             `json:"config"`
	Resources   []ResourceWithDependencies `json:"resources"`
	Expected    []*ExpectedEndpoint        `json:"expected"`
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
	ExcludeTargetNets   []string `json:"excludeTargetNets"`
	NAT64Networks       []string `json:"nat64Networks"`
	ServiceTypeFilter   []string `json:"serviceTypeFilter"`
	Provider            string   `json:"provider"`
	PreferAlias         bool     `json:"preferAlias"`
}

// ParsedResources holds the parsed Kubernetes resources from a scenario.
type ParsedResources struct {
	Ingresses      []*networkingv1.Ingress
	Services       []*corev1.Service
	EndpointSlices []*discoveryv1.EndpointSlice
	Pods           []*corev1.Pod
	Nodes          []*corev1.Node
	DNSEndpoints   []*apiv1alpha1.DNSEndpoint
}

// LoadedResources holds the clients.
type LoadedResources struct {
	// K8sClient is the fake Kubernetes clientset for core/networking/discovery resources.
	K8sClient *fake.Clientset
	// DNSEndpoints are the parsed DNSEndpoint CRD objects ready to be injected into the CRD source fake cache.
	DNSEndpoints []*apiv1alpha1.DNSEndpoint
}
