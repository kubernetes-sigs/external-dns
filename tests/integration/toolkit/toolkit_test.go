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
	"testing"

	"encoding/json"

	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadScenarios_Valid(t *testing.T) {
	yaml := []byte(`
scenarios:
  - name: my-scenario
    description: A test scenario
    config:
      sources: ["service"]
`)
	s, err := LoadScenarios(yaml)
	require.NoError(t, err)
	require.Len(t, s.Scenarios, 1)
	assert.Equal(t, "my-scenario", s.Scenarios[0].Name)
}

func TestLoadScenarios_InvalidYAML(t *testing.T) {
	_, err := LoadScenarios([]byte(":\tinvalid"))
	assert.Error(t, err)
}

func TestLoadScenarios_MissingName(t *testing.T) {
	yaml := []byte(`
scenarios:
  - description: no name here
    config:
      sources: ["service"]
`)
	_, err := LoadScenarios(yaml)
	assert.ErrorContains(t, err, "missing required field: name")
}

func TestLoadScenarios_MissingDescription(t *testing.T) {
	yaml := []byte(`
scenarios:
  - name: my-scenario
    config:
      sources: ["service"]
`)
	_, err := LoadScenarios(yaml)
	assert.ErrorContains(t, err, "missing required field: description")
}

func TestLoadScenarios_MissingSources(t *testing.T) {
	yaml := []byte(`
scenarios:
  - name: my-scenario
    description: A test scenario
    config: {}
`)
	_, err := LoadScenarios(yaml)
	assert.ErrorContains(t, err, "missing required field: config.sources")
}

func rawService() []byte {
	return []byte(`apiVersion: v1
kind: Service
metadata:
  name: svc
  namespace: default
spec:
  type: ClusterIP
  selector:
    app: myapp
`)
}

func rawIngress() []byte {
	return []byte(`apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: ing
  namespace: default
spec: {}
`)
}

func rawPod() []byte {
	return []byte(`apiVersion: v1
kind: Pod
metadata:
  name: pod-0
  namespace: default
spec: {}
`)
}

func rawEndpointSlice() []byte {
	return []byte(`apiVersion: discovery.k8s.io/v1
kind: EndpointSlice
metadata:
  name: eps
  namespace: default
addressType: IPv4
`)
}

func TestParseResources_Service(t *testing.T) {
	parsed, err := ParseResources([]ResourceWithDependencies{
		{Resource: runtime.RawExtension{Raw: rawService()}},
	})
	require.NoError(t, err)
	require.Len(t, parsed.Services, 1)
	assert.Equal(t, "svc", parsed.Services[0].Name)
}

func TestParseResources_ServiceWithPodDependencies(t *testing.T) {
	parsed, err := ParseResources([]ResourceWithDependencies{
		{
			Resource:     runtime.RawExtension{Raw: rawService()},
			Dependencies: &ResourceDependencies{Pods: &PodDependencies{Replicas: 2}},
		},
	})
	require.NoError(t, err)
	assert.Len(t, parsed.Services, 1)
	assert.Len(t, parsed.Pods, 2)
	assert.Len(t, parsed.EndpointSlices, 1)
}

func TestParseResources_Ingress(t *testing.T) {
	parsed, err := ParseResources([]ResourceWithDependencies{
		{Resource: runtime.RawExtension{Raw: rawIngress()}},
	})
	require.NoError(t, err)
	require.Len(t, parsed.Ingresses, 1)
	assert.Equal(t, "ing", parsed.Ingresses[0].Name)
}

func TestParseResources_Pod(t *testing.T) {
	parsed, err := ParseResources([]ResourceWithDependencies{
		{Resource: runtime.RawExtension{Raw: rawPod()}},
	})
	require.NoError(t, err)
	require.Len(t, parsed.Pods, 1)
	assert.Equal(t, "pod-0", parsed.Pods[0].Name)
}

func TestParseResources_EndpointSlice(t *testing.T) {
	parsed, err := ParseResources([]ResourceWithDependencies{
		{Resource: runtime.RawExtension{Raw: rawEndpointSlice()}},
	})
	require.NoError(t, err)
	require.Len(t, parsed.EndpointSlices, 1)
	assert.Equal(t, "eps", parsed.EndpointSlices[0].Name)
}

func TestParseResources_UnsupportedType(t *testing.T) {
	raw := []byte(`apiVersion: v1
kind: ConfigMap
metadata:
  name: cm
  namespace: default
`)
	_, err := ParseResources([]ResourceWithDependencies{
		{Resource: runtime.RawExtension{Raw: raw}},
	})
	assert.ErrorContains(t, err, "unsupported resource type")
}

func TestParseResources_InvalidRaw(t *testing.T) {
	_, err := ParseResources([]ResourceWithDependencies{
		{Resource: runtime.RawExtension{Raw: []byte("not yaml")}},
	})
	assert.Error(t, err)
}

func TestLoadResources_Service(t *testing.T) {
	client, err := LoadResources(t.Context(), Scenario{
		Resources: []ResourceWithDependencies{
			{Resource: runtime.RawExtension{Raw: rawService()}},
		},
	})
	require.NoError(t, err)
	svcs, err := client.CoreV1().Services("default").List(t.Context(), metav1.ListOptions{})
	require.NoError(t, err)
	assert.Len(t, svcs.Items, 1)
}

func TestLoadResources_ServiceWithLoadBalancerStatus(t *testing.T) {
	svc := &corev1.Service{
		TypeMeta:   metav1.TypeMeta{APIVersion: "v1", Kind: "Service"},
		ObjectMeta: metav1.ObjectMeta{Name: "lb-svc", Namespace: "default"},
		Status: corev1.ServiceStatus{
			LoadBalancer: corev1.LoadBalancerStatus{
				Ingress: []corev1.LoadBalancerIngress{{IP: "1.2.3.4"}},
			},
		},
	}
	raw, err := encodeObject(svc)
	require.NoError(t, err)

	client, err := LoadResources(t.Context(), Scenario{
		Resources: []ResourceWithDependencies{
			{Resource: runtime.RawExtension{Raw: raw}},
		},
	})
	require.NoError(t, err)
	got, err := client.CoreV1().Services("default").Get(t.Context(), "lb-svc", metav1.GetOptions{})
	require.NoError(t, err)
	assert.Equal(t, "1.2.3.4", got.Status.LoadBalancer.Ingress[0].IP)
}

func TestLoadResources_IngressWithLoadBalancerStatus(t *testing.T) {
	ing := &networkingv1.Ingress{
		TypeMeta:   metav1.TypeMeta{APIVersion: "networking.k8s.io/v1", Kind: "Ingress"},
		ObjectMeta: metav1.ObjectMeta{Name: "lb-ing", Namespace: "default"},
		Status: networkingv1.IngressStatus{
			LoadBalancer: networkingv1.IngressLoadBalancerStatus{
				Ingress: []networkingv1.IngressLoadBalancerIngress{{IP: "5.6.7.8"}},
			},
		},
	}
	raw, err := encodeObject(ing)
	require.NoError(t, err)

	client, err := LoadResources(t.Context(), Scenario{
		Resources: []ResourceWithDependencies{
			{Resource: runtime.RawExtension{Raw: raw}},
		},
	})
	require.NoError(t, err)
	got, err := client.NetworkingV1().Ingresses("default").Get(t.Context(), "lb-ing", metav1.GetOptions{})
	require.NoError(t, err)
	assert.Equal(t, "5.6.7.8", got.Status.LoadBalancer.Ingress[0].IP)
}

func TestLoadResources_Pods(t *testing.T) {
	client, err := LoadResources(t.Context(), Scenario{
		Resources: []ResourceWithDependencies{
			{Resource: runtime.RawExtension{Raw: rawPod()}},
		},
	})
	require.NoError(t, err)
	pods, err := client.CoreV1().Pods("default").List(t.Context(), metav1.ListOptions{})
	require.NoError(t, err)
	assert.Len(t, pods.Items, 1)
}

func TestLoadResources_EndpointSlices(t *testing.T) {
	client, err := LoadResources(t.Context(), Scenario{
		Resources: []ResourceWithDependencies{
			{Resource: runtime.RawExtension{Raw: rawEndpointSlice()}},
		},
	})
	require.NoError(t, err)
	epsList, err := client.DiscoveryV1().EndpointSlices("default").List(t.Context(), metav1.ListOptions{})
	require.NoError(t, err)
	assert.Len(t, epsList.Items, 1)
}

func TestLoadResources_ParseError(t *testing.T) {
	_, err := LoadResources(t.Context(), Scenario{
		Resources: []ResourceWithDependencies{
			{Resource: runtime.RawExtension{Raw: []byte("not yaml")}},
		},
	})
	assert.Error(t, err)
}

func TestCreateWrappedSource(t *testing.T) {
	client, err := LoadResources(t.Context(), Scenario{
		Resources: []ResourceWithDependencies{
			{Resource: runtime.RawExtension{Raw: rawService()}},
		},
	})
	require.NoError(t, err)

	src, err := CreateWrappedSource(t.Context(), client, ScenarioConfig{
		Sources: []string{"service"},
	})
	require.NoError(t, err)
	assert.NotNil(t, src)
}

func TestScenarioToConfig(t *testing.T) {
	cfg, err := scenarioToConfig(ScenarioConfig{
		Sources:           []string{"service"},
		ServiceTypeFilter: []string{"LoadBalancer"},
		DefaultTargets:    []string{"1.2.3.4"},
		TargetNetFilter:   []string{"10.0.0.0/8"},
		Provider:          "inmemory",
	})
	require.NoError(t, err)
	assert.NotNil(t, cfg)
}

// encodeObject serializes a runtime.Object to JSON for use as RawExtension.Raw.
func encodeObject(obj runtime.Object) ([]byte, error) {
	return json.Marshal(obj)
}
