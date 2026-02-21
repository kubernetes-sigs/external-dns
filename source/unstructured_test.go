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

package source

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"sigs.k8s.io/external-dns/internal/testutils"
	"sigs.k8s.io/external-dns/source/types"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/source/annotations"
)

func TestParseResourceArg(t *testing.T) {
	tests := []struct {
		name       string
		identifier string
		expected   schema.GroupVersionResource
		valid      bool
	}{
		{
			name:       "full resource with group",
			identifier: "virtualmachineinstances.v1.kubevirt.io",
			expected: schema.GroupVersionResource{
				Group:    "kubevirt.io",
				Version:  "v1",
				Resource: "virtualmachineinstances",
			},
			valid: true,
		},
		{
			name:       "resource with multi-part group",
			identifier: "ingresses.v1.networking.k8s.io",
			expected: schema.GroupVersionResource{
				Group:    "networking.k8s.io",
				Version:  "v1",
				Resource: "ingresses",
			},
			valid: true,
		},
		{
			name:       "core API resource (empty group with trailing dot)",
			identifier: "pods.v1.",
			expected: schema.GroupVersionResource{
				Group:    "",
				Version:  "v1",
				Resource: "pods",
			},
			valid: true,
		},
		{
			name:       "invalid - single part",
			identifier: "pods",
			valid:      false,
		},
		{
			name:       "invalid - empty string",
			identifier: "",
			valid:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gvr, _ := schema.ParseResourceArg(tt.identifier)
			if tt.valid {
				require.NotNil(t, gvr)
				assert.Equal(t, tt.expected, *gvr)
			} else {
				assert.Nil(t, gvr)
			}
		})
	}
}

func TestUnstructuredWrapperImplementsKubeObject(t *testing.T) {
	u := &unstructured.Unstructured{
		Object: map[string]any{
			"apiVersion": "kubevirt.io/v1",
			"kind":       "VirtualMachineInstance",
			"metadata": map[string]any{
				"name":      "test-vm",
				"namespace": "default",
				"labels": map[string]any{
					"app": "test",
				},
			},
		},
	}

	wrapped := newUnstructuredWrapper(u)
	assert.Equal(t, "test-vm", wrapped.Name)
	assert.Equal(t, "default", wrapped.Namespace)
	assert.Equal(t, "VirtualMachineInstance", wrapped.Kind)
	assert.Equal(t, "kubevirt.io/v1", wrapped.APIVersion)
	assert.Equal(t, map[string]string{"app": "test"}, wrapped.Labels)
	assert.Equal(t, "test-vm", wrapped.GetName())
	assert.Equal(t, "default", wrapped.GetNamespace())
	assert.Same(t, u, wrapped.Unstructured)
	// Verify it implements runtime.Object via embedding
	gvk := wrapped.GetObjectKind().GroupVersionKind()
	assert.Equal(t, "VirtualMachineInstance", gvk.Kind)
}

func TestUnstructured_DifferentScenarios(t *testing.T) {
	type cfg struct {
		resources        []string
		labelSelector    string
		annotationFilter string
	}

	for _, tt := range []struct {
		title    string
		cfg      cfg
		objects  []*unstructured.Unstructured
		expected []*endpoint.Endpoint
	}{
		{
			title: "read from annotations with IPv6 target",
			cfg: cfg{
				resources: []string{"virtualmachineinstances.v1.kubevirt.io"},
			},
			objects: []*unstructured.Unstructured{
				{
					Object: map[string]any{
						"apiVersion": "kubevirt.io/v1",
						"kind":       "VirtualMachineInstance",
						"metadata": map[string]any{
							"name":      "my-vm",
							"namespace": "default",
							"annotations": map[string]any{
								annotations.HostnameKey: "my-vm.example.com",
								annotations.TargetKey:   "::1234:5678",
							},
						},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("my-vm.example.com", endpoint.RecordTypeAAAA, "::1234:5678").
					WithLabel(endpoint.ResourceLabelKey, "virtualmachineinstance/default/my-vm"),
			},
		},
		{
			title: "rancher node with ttl",
			cfg: cfg{
				resources: []string{"nodes.v3.management.cattle.io"},
			},
			objects: []*unstructured.Unstructured{
				{
					Object: map[string]any{
						"apiVersion": "management.cattle.io/v3",
						"kind":       "Node",
						"metadata": map[string]any{
							"name":      "my-node-1",
							"namespace": "cattle-system",
							"labels": map[string]any{
								"cattle.io/creator":                    "norman",
								"node-role.kubernetes.io/controlplane": "true",
							},
							"annotations": map[string]any{
								annotations.HostnameKey: "my-node-1.nodes.example.com",
								annotations.TargetKey:   "203.0.113.10",
								annotations.TtlKey:      "300",
							},
						},
						"spec": map[string]any{
							"clusterName": "c-abcde",
							"hostname":    "my-node-1",
						},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpointWithTTL("my-node-1.nodes.example.com", endpoint.RecordTypeA, 300, "203.0.113.10").
					WithLabel(endpoint.ResourceLabelKey, "node/cattle-system/my-node-1"),
			},
		},
		{
			title: "with controller annotations match",
			cfg: cfg{
				resources: []string{"replicationgroups.v1.elasticache.upbound.io"},
			},
			objects: []*unstructured.Unstructured{
				{
					Object: map[string]any{
						"apiVersion": "elasticache.upbound.io/v1",
						"kind":       "ReplicationGroup",
						"metadata": map[string]any{
							"name":      "cache",
							"namespace": "default",
							"annotations": map[string]any{
								annotations.HostnameKey:   "my-vm.redis.tld",
								annotations.TargetKey:     "1.1.1.0",
								annotations.ControllerKey: annotations.ControllerValue,
							},
						},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("my-vm.redis.tld", endpoint.RecordTypeA, "1.1.1.0").
					WithLabel(endpoint.ResourceLabelKey, "replicationgroup/default/cache"),
			},
		},
		{
			title: "with controller annotations do not match",
			cfg: cfg{
				resources: []string{"replicationgroups.v1.elasticache.upbound.io"},
			},
			objects: []*unstructured.Unstructured{
				{
					Object: map[string]any{
						"apiVersion": "elasticache.upbound.io/v1",
						"kind":       "ReplicationGroup",
						"metadata": map[string]any{
							"name":      "my-vm",
							"namespace": "default",
							"annotations": map[string]any{
								annotations.HostnameKey:   "my-vm.redis.tld",
								annotations.TargetKey:     "10.10.10.0",
								annotations.ControllerKey: "custom-controller",
							},
						},
					},
				},
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			title: "labelSelector matches",
			cfg: cfg{
				resources:     []string{"virtualmachineinstances.v1.kubevirt.io"},
				labelSelector: "env=prod",
			},
			objects: []*unstructured.Unstructured{
				{
					Object: map[string]any{
						"apiVersion": "kubevirt.io/v1",
						"kind":       "VirtualMachineInstance",
						"metadata": map[string]any{
							"name":      "prod-vm",
							"namespace": "default",
							"labels": map[string]any{
								"env": "prod",
							},
							"annotations": map[string]any{
								annotations.HostnameKey: "prod-vm.example.com",
								annotations.TargetKey:   "10.0.0.1",
							},
						},
					},
				},
				{
					Object: map[string]any{
						"apiVersion": "kubevirt.io/v1",
						"kind":       "VirtualMachineInstance",
						"metadata": map[string]any{
							"name":      "dev-vm",
							"namespace": "default",
							"labels": map[string]any{
								"env": "dev",
							},
							"annotations": map[string]any{
								annotations.HostnameKey: "dev-vm.example.com",
								annotations.TargetKey:   "10.0.0.2",
							},
						},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("prod-vm.example.com", endpoint.RecordTypeA, "10.0.0.1").
					WithLabel(endpoint.ResourceLabelKey, "virtualmachineinstance/default/prod-vm"),
			},
		},
		{
			title: "labelSelector no match",
			cfg: cfg{
				resources:     []string{"virtualmachineinstances.v1.kubevirt.io"},
				labelSelector: "env=staging",
			},
			objects: []*unstructured.Unstructured{
				{
					Object: map[string]any{
						"apiVersion": "kubevirt.io/v1",
						"kind":       "VirtualMachineInstance",
						"metadata": map[string]any{
							"name":      "prod-vm",
							"namespace": "default",
							"labels": map[string]any{
								"env": "prod",
							},
							"annotations": map[string]any{
								annotations.HostnameKey: "prod-vm.example.com",
								annotations.TargetKey:   "10.0.0.1",
							},
						},
					},
				},
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			title: "annotationFilter matches",
			cfg: cfg{
				resources:        []string{"virtualmachineinstances.v1.kubevirt.io"},
				annotationFilter: "team=platform",
			},
			objects: []*unstructured.Unstructured{
				{
					Object: map[string]any{
						"apiVersion": "kubevirt.io/v1",
						"kind":       "VirtualMachineInstance",
						"metadata": map[string]any{
							"name":      "platform-vm",
							"namespace": "default",
							"annotations": map[string]any{
								"team":                  "platform",
								annotations.HostnameKey: "platform-vm.example.com",
								annotations.TargetKey:   "10.0.0.1",
							},
						},
					},
				},
				{
					Object: map[string]any{
						"apiVersion": "kubevirt.io/v1",
						"kind":       "VirtualMachineInstance",
						"metadata": map[string]any{
							"name":      "backend-vm",
							"namespace": "default",
							"annotations": map[string]any{
								"team":                  "backend",
								annotations.HostnameKey: "backend-vm.example.com",
								annotations.TargetKey:   "10.0.0.2",
							},
						},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("platform-vm.example.com", endpoint.RecordTypeA, "10.0.0.1").
					WithLabel(endpoint.ResourceLabelKey, "virtualmachineinstance/default/platform-vm"),
			},
		},
		{
			title: "annotationFilter no match",
			cfg: cfg{
				resources:        []string{"virtualmachineinstances.v1.kubevirt.io"},
				annotationFilter: "team=security",
			},
			objects: []*unstructured.Unstructured{
				{
					Object: map[string]any{
						"apiVersion": "kubevirt.io/v1",
						"kind":       "VirtualMachineInstance",
						"metadata": map[string]any{
							"name":      "platform-vm",
							"namespace": "default",
							"annotations": map[string]any{
								"team":                  "platform",
								annotations.HostnameKey: "platform-vm.example.com",
								annotations.TargetKey:   "10.0.0.1",
							},
						},
					},
				},
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			title: "labelSelector and annotationFilter combined",
			cfg: cfg{
				resources:        []string{"virtualmachineinstances.v1.kubevirt.io"},
				labelSelector:    "env=prod",
				annotationFilter: "team=platform",
			},
			objects: []*unstructured.Unstructured{
				{
					Object: map[string]any{
						"apiVersion": "kubevirt.io/v1",
						"kind":       "VirtualMachineInstance",
						"metadata": map[string]any{
							"name":      "prod-platform-vm",
							"namespace": "default",
							"labels": map[string]any{
								"env": "prod",
							},
							"annotations": map[string]any{
								"team":                  "platform",
								annotations.HostnameKey: "prod-platform-vm.example.com",
								annotations.TargetKey:   "10.0.0.1",
							},
						},
					},
				},
				{
					Object: map[string]any{
						"apiVersion": "kubevirt.io/v1",
						"kind":       "VirtualMachineInstance",
						"metadata": map[string]any{
							"name":      "prod-backend-vm",
							"namespace": "default",
							"labels": map[string]any{
								"env": "prod",
							},
							"annotations": map[string]any{
								"team":                  "backend",
								annotations.HostnameKey: "prod-backend-vm.example.com",
								annotations.TargetKey:   "10.0.0.2",
							},
						},
					},
				},
				{
					Object: map[string]any{
						"apiVersion": "kubevirt.io/v1",
						"kind":       "VirtualMachineInstance",
						"metadata": map[string]any{
							"name":      "dev-platform-vm",
							"namespace": "default",
							"labels": map[string]any{
								"env": "dev",
							},
							"annotations": map[string]any{
								"team":                  "platform",
								annotations.HostnameKey: "dev-platform-vm.example.com",
								annotations.TargetKey:   "10.0.0.3",
							},
						},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("prod-platform-vm.example.com", endpoint.RecordTypeA, "10.0.0.1").
					WithLabel(endpoint.ResourceLabelKey, "virtualmachineinstance/default/prod-platform-vm"),
			},
		},
	} {
		t.Run(tt.title, func(t *testing.T) {
			kubeClient, dynamicClient := setupUnstructuredTestClients(t, tt.cfg.resources, tt.objects)

			labelSelector := labels.Everything()
			if tt.cfg.labelSelector != "" {
				var err error
				labelSelector, err = labels.Parse(tt.cfg.labelSelector)
				require.NoError(t, err)
			}

			src, err := NewUnstructuredFQDNSource(
				t.Context(),
				dynamicClient,
				kubeClient,
				"",
				tt.cfg.annotationFilter,
				labelSelector,
				tt.cfg.resources,
				"",
				"",
				"",
				false,
			)
			require.NoError(t, err)

			endpoints, err := src.Endpoints(t.Context())
			require.NoError(t, err)

			validateEndpoints(t, endpoints, tt.expected)

			for _, ep := range endpoints {
				require.Contains(t, ep.Labels, endpoint.ResourceLabelKey)
			}
		})
	}
}

func TestProcessEndpoint_Service_RefObjectExist(t *testing.T) {
	resources := []string{"virtualmachineinstances.v1.kubevirt.io"}
	objects := []*unstructured.Unstructured{
		{
			Object: map[string]any{
				"apiVersion": "kubevirt.io/v1",
				"kind":       "VirtualMachineInstance",
				"metadata": map[string]any{
					"name":      "prod-platform-vm",
					"namespace": "default",
					"labels": map[string]any{
						"env": "prod",
					},
					"annotations": map[string]any{
						"team":                  "platform",
						annotations.HostnameKey: "prod-platform-vm.example.com",
						annotations.TargetKey:   "10.0.0.1",
					},
					"uid": "12345",
				},
			},
		},
	}

	kubeClient, dynamicClient := setupUnstructuredTestClients(t, resources, objects)

	src, err := NewUnstructuredFQDNSource(
		t.Context(),
		dynamicClient,
		kubeClient,
		"",
		"",
		labels.Everything(),
		resources,
		"",
		"",
		"",
		false,
	)
	require.NoError(t, err)

	endpoints, err := src.Endpoints(t.Context())
	require.NoError(t, err)
	testutils.AssertEndpointsHaveRefObject(t, endpoints, types.Unstructured, len(objects))
}

func TestEndpointsForHostsAndTargets(t *testing.T) {
	tests := []struct {
		name      string
		hostnames []string
		targets   []string
		expected  []*endpoint.Endpoint
	}{
		{
			name:      "empty hostnames returns nil",
			hostnames: []string{},
			targets:   []string{"192.168.1.1"},
			expected:  nil,
		},
		{
			name:      "empty targets returns nil",
			hostnames: []string{"example.com"},
			targets:   []string{},
			expected:  nil,
		},
		{
			name:      "duplicate hostname with IPv4 and IPv6 targets",
			hostnames: []string{"example.com", "example.com"},
			targets:   []string{"192.168.1.1", "192.168.1.1", "2001:db8::1"},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("example.com", endpoint.RecordTypeA, "192.168.1.1"),
				endpoint.NewEndpoint("example.com", endpoint.RecordTypeAAAA, "2001:db8::1"),
			},
		},
		{
			name:      "multiple hostnames with single target",
			hostnames: []string{"example.com", "www.example.com"},
			targets:   []string{"192.168.1.1"},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("example.com", endpoint.RecordTypeA, "192.168.1.1"),
				endpoint.NewEndpoint("www.example.com", endpoint.RecordTypeA, "192.168.1.1"),
			},
		},
		{
			name:      "multiple of each type maintains grouping",
			hostnames: []string{"example.com"},
			targets:   []string{"192.168.1.1", "192.168.1.2", "2001:db8::1", "2001:db8::2", "a.example.com", "b.example.com"},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("example.com", endpoint.RecordTypeA, "192.168.1.1", "192.168.1.2"),
				endpoint.NewEndpoint("example.com", endpoint.RecordTypeAAAA, "2001:db8::1", "2001:db8::2"),
				endpoint.NewEndpoint("example.com", endpoint.RecordTypeCNAME, "a.example.com", "b.example.com"),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := EndpointsForHostsAndTargets(tc.hostnames, tc.targets)
			if tc.expected == nil {
				assert.Nil(t, result)
				return
			}
			assert.Len(t, result, len(tc.expected))
			for i, ep := range result {
				assert.Equal(t, tc.expected[i].DNSName, ep.DNSName, "DNSName mismatch at index %d", i)
				assert.Equal(t, tc.expected[i].RecordType, ep.RecordType, "RecordType mismatch at index %d", i)
				assert.ElementsMatch(t, tc.expected[i].Targets, ep.Targets, "Targets mismatch at index %d", i)
			}
		})
	}
}
