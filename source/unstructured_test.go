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
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	discoveryfake "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/dynamic"
	dynamicfake "k8s.io/client-go/dynamic/fake"
	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/tools/cache"

	"sigs.k8s.io/external-dns/internal/testutils"
	"sigs.k8s.io/external-dns/source/types"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/source/annotations"
	templatetest "sigs.k8s.io/external-dns/source/template/testutil"
)

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
		combine          bool
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
		{
			title: "provider-specific annotation is not supported and is ignored",
			cfg: cfg{
				resources: []string{"machines.v1beta1.cluster.x-k8s.io"},
			},
			objects: []*unstructured.Unstructured{
				{
					Object: map[string]any{
						"apiVersion": "cluster.x-k8s.io/v1beta1",
						"kind":       "Machine",
						"metadata": map[string]any{
							"name":      "control-plane",
							"namespace": "default",
							"labels": map[string]any{
								"cluster.x-k8s.io/cluster-name":  "test-cluster",
								"cluster.x-k8s.io/control-plane": "",
							},
							"annotations": map[string]any{
								annotations.HostnameKey:      "control-plane.example.com",
								annotations.TargetKey:        "10.0.0.1",
								annotations.CloudflarePrefix: "cloudflare-specific-annotation",
							},
						},
						"spec": map[string]any{
							"clusterName": "test-cluster",
							"bootstrap": map[string]any{
								"dataSecretName": "control-plane-bootstrap",
							},
							"version": "v1.26.0",
						},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("control-plane.example.com", endpoint.RecordTypeA, "10.0.0.1").
					WithLabel(endpoint.ResourceLabelKey, "machine/default/control-plane"),
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
				&Config{
					AnnotationFilter:      tt.cfg.annotationFilter,
					LabelFilter:           labelSelector,
					UnstructuredResources: tt.cfg.resources,
					TemplateEngine:        templatetest.MustEngine(t, "", "", "", tt.cfg.combine),
				},
			)
			require.NoError(t, err)

			endpoints, err := src.Endpoints(t.Context())
			require.NoError(t, err)

			testutils.ValidateEndpoints(t, endpoints, tt.expected)

			for _, ep := range endpoints {
				require.Contains(t, ep.Labels, endpoint.ResourceLabelKey)
			}
		})
	}
}

func TestProcessEndpoint_Unstructured_RefObjectExist(t *testing.T) {
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
		&Config{
			LabelFilter:           labels.Everything(),
			UnstructuredResources: resources,
		},
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
			testutils.ValidateEndpoints(t, result, tc.expected)
		})
	}
}

// setupUnstructuredTestClients creates fake kube and dynamic clients with the given resources and objects.
func setupUnstructuredTestClients(t *testing.T, resources []string, objects []*unstructured.Unstructured) (
	kubernetes.Interface, dynamic.Interface,
) {
	t.Helper()

	// Parse all resource identifiers and build apiVersion → GVR map in one pass
	gvrs := make([]schema.GroupVersionResource, 0, len(resources))
	apiVersionToGVR := make(map[string]schema.GroupVersionResource, len(resources))
	for _, res := range resources {
		if strings.Count(res, ".") == 1 {
			res += "."
		}
		gvr, _ := schema.ParseResourceArg(res)
		require.NotNil(t, gvr, "invalid resource identifier: %s", res)
		gvrs = append(gvrs, *gvr)
		apiVersionToGVR[gvr.GroupVersion().String()] = *gvr
	}

	// Derive kind and list kind from objects
	gvrToKind := make(map[schema.GroupVersionResource]string, len(gvrs))
	gvrToListKind := make(map[schema.GroupVersionResource]string, len(gvrs))
	for _, obj := range objects {
		if gvr, ok := apiVersionToGVR[obj.GetAPIVersion()]; ok {
			gvrToKind[gvr] = obj.GetKind()
			gvrToListKind[gvr] = obj.GetKind() + "List"
		}
	}

	// Build discovery resource lists
	apiResourceLists := make([]*metav1.APIResourceList, 0, len(gvrs))
	for _, gvr := range gvrs {
		apiResourceLists = append(apiResourceLists, &metav1.APIResourceList{
			GroupVersion: gvr.GroupVersion().String(),
			APIResources: []metav1.APIResource{{
				Name:       gvr.Resource,
				Namespaced: true,
				Kind:       gvrToKind[gvr],
			}},
		})
	}

	kubeClient := fake.NewClientset()
	kubeClient.Discovery().(*discoveryfake.FakeDiscovery).Resources = apiResourceLists

	dynamicClient := dynamicfake.NewSimpleDynamicClientWithCustomListKinds(runtime.NewScheme(), gvrToListKind)

	for _, obj := range objects {
		gvr, ok := apiVersionToGVR[obj.GetAPIVersion()]
		require.True(t, ok, "no resource found for apiVersion %s", obj.GetAPIVersion())
		_, err := dynamicClient.Resource(gvr).Namespace(obj.GetNamespace()).Create(
			t.Context(), obj, metav1.CreateOptions{})
		require.NoError(t, err)
	}

	return kubeClient, dynamicClient
}

func TestDiscoverResources_Errors(t *testing.T) {
	for _, tt := range []struct {
		title     string
		resources []string
		discovery []*metav1.APIResourceList
		wantErr   string
	}{
		{
			title:     "invalid resource identifier with no dots",
			resources: []string{"justname"},
			wantErr:   "invalid resource identifier",
		},
		{
			title:     "discovery fails for unknown group version",
			resources: []string{"virtualmachineinstances.v1.kubevirt.io"},
			discovery: []*metav1.APIResourceList{}, // empty: kubevirt.io/v1 not registered
			wantErr:   "failed to discover resources",
		},
		{
			title:     "resource name not found in group version",
			resources: []string{"nonexistent.v1.kubevirt.io"},
			discovery: []*metav1.APIResourceList{
				{
					GroupVersion: "kubevirt.io/v1",
					APIResources: []metav1.APIResource{{Name: "virtualmachineinstances"}},
				},
			},
			wantErr: "not found",
		},
	} {
		t.Run(tt.title, func(t *testing.T) {
			kubeClient := fake.NewClientset()
			kubeClient.Discovery().(*discoveryfake.FakeDiscovery).Resources = tt.discovery

			_, err := discoverResources(kubeClient, tt.resources)
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.wantErr)
		})
	}
}

func TestAddEventHandler_Unstructured(t *testing.T) {
	resources := []string{"virtualmachineinstances.v1.kubevirt.io"}
	objects := []*unstructured.Unstructured{
		{
			Object: map[string]any{
				"apiVersion": "kubevirt.io/v1",
				"kind":       "VirtualMachineInstance",
				"metadata": map[string]any{
					"name":      "my-vm",
					"namespace": "default",
				},
			},
		},
	}
	kubeClient, dynamicClient := setupUnstructuredTestClients(t, resources, objects)

	src, err := NewUnstructuredFQDNSource(
		t.Context(),
		dynamicClient,
		kubeClient,
		&Config{
			LabelFilter:           labels.Everything(),
			UnstructuredResources: resources,
		},
	)
	require.NoError(t, err)

	src.AddEventHandler(t.Context(), func() {})
}

func TestNewUnstructuredFQDNSource_Errors(t *testing.T) {
	for _, tt := range []struct {
		title         string
		resources     []string
		discovery     []*metav1.APIResourceList
		dynamicClient dynamic.Interface
		ctx           func() context.Context
		wantErr       string
	}{
		{
			title:         "discoverResources error propagates",
			resources:     []string{"virtualmachineinstances.v1.kubevirt.io"},
			discovery:     nil, // empty: group/version unknown → discovery fails
			dynamicClient: dynamicfake.NewSimpleDynamicClientWithCustomListKinds(runtime.NewScheme(), nil),
			ctx:           t.Context,
			wantErr:       "failed to discover",
		},
		{
			title:     "WaitForDynamicCacheSync error propagates",
			resources: []string{"virtualmachineinstances.v1.kubevirt.io"},
			discovery: []*metav1.APIResourceList{
				{
					GroupVersion: "kubevirt.io/v1",
					APIResources: []metav1.APIResource{{
						Name: "virtualmachineinstances", Namespaced: true, Kind: "VirtualMachineInstance",
					}},
				},
			},
			// Empty scheme: List always returns "no kind registered", so HasSynced() stays false.
			// Pre-cancelled context: WaitForCacheSync sees a closed stopCh and returns immediately.
			dynamicClient: dynamicfake.NewSimpleDynamicClientWithCustomListKinds(runtime.NewScheme(), nil),
			ctx: func() context.Context {
				ctx, cancel := context.WithCancel(t.Context())
				cancel()
				return ctx
			},
			wantErr: "failed to sync",
		},
	} {
		t.Run(tt.title, func(t *testing.T) {
			kubeClient := fake.NewClientset()
			kubeClient.Discovery().(*discoveryfake.FakeDiscovery).Resources = tt.discovery

			_, err := NewUnstructuredFQDNSource(tt.ctx(), tt.dynamicClient, kubeClient, &Config{
				LabelFilter:           labels.Everything(),
				UnstructuredResources: tt.resources,
			})
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.wantErr)
		})
	}
}

// fakeTestIndexer is a minimal cache.Indexer that returns a wrong-type object,
// causing GetByKey[*unstructured.Unstructured] to fail the type assertion.
type fakeTestIndexer struct {
	cache.Indexer // nil embed; only overridden methods are called in tests
}

func (f *fakeTestIndexer) ListIndexFuncValues(_ string) []string {
	return []string{"default/my-vm"}
}

func (f *fakeTestIndexer) GetByKey(_ string) (any, bool, error) {
	return "not-an-unstructured-object", true, nil
}

// fakeTestSharedIndexInformer wraps a cache.SharedIndexInformer with a custom indexer.
type fakeTestSharedIndexInformer struct {
	cache.SharedIndexInformer // nil embed
	indexer                   cache.Indexer
}

func (f *fakeTestSharedIndexInformer) GetIndexer() cache.Indexer { return f.indexer }

// fakeTestGenericInformer implements kubeinformers.GenericInformer with a custom shared informer.
type fakeTestGenericInformer struct {
	inf cache.SharedIndexInformer
}

func (f *fakeTestGenericInformer) Informer() cache.SharedIndexInformer { return f.inf }
func (f *fakeTestGenericInformer) Lister() cache.GenericLister         { panic("not needed") }

// TestEndpointsFromInformer_GetByKeyError verifies that a GetByKey type-assertion failure
// causes the object to be skipped (continue) rather than returning an error.
func TestEndpointsFromInformer_GetByKeyError(t *testing.T) {
	idx := &fakeTestIndexer{}
	inf := &fakeTestSharedIndexInformer{indexer: idx}

	src := &unstructuredSource{
		informers: []kubeinformers.GenericInformer{
			&fakeTestGenericInformer{inf: inf},
		},
	}

	endpoints, err := src.Endpoints(t.Context())
	require.NoError(t, err)
	assert.Empty(t, endpoints)
}

// TestUnstructuredSource_TemplateErrors verifies that template execution errors propagate
// correctly through endpointsFromTemplate/endpointsFromFQDNTargetTemplate → endpointsFromInformer → Endpoints.
// {{index . 0}} reliably fails at runtime for *unstructuredWrapper (can't index a struct).
func TestUnstructuredSource_TemplateErrors(t *testing.T) {
	resources := []string{"virtualmachineinstances.v1.kubevirt.io"}
	objects := []*unstructured.Unstructured{
		{
			Object: map[string]any{
				"apiVersion": "kubevirt.io/v1",
				"kind":       "VirtualMachineInstance",
				"metadata": map[string]any{
					"name":      "my-vm",
					"namespace": "default",
				},
			},
		},
	}

	for _, tt := range []struct {
		title              string
		fqdnTemplate       string
		targetTemplate     string
		fqdnTargetTemplate string
	}{
		{
			title:        "ExecFQDN runtime error propagates to Endpoints",
			fqdnTemplate: "{{index . 0}}",
		},
		{
			title:          "ExecTarget runtime error propagates to Endpoints",
			fqdnTemplate:   "{{.Name}}.example.com",
			targetTemplate: "{{index . 0}}",
		},
		{
			title:              "ExecFQDNTarget runtime error propagates to Endpoints",
			fqdnTargetTemplate: "{{index . 0}}",
		},
	} {
		t.Run(tt.title, func(t *testing.T) {
			kubeClient, dynamicClient := setupUnstructuredTestClients(t, resources, objects)

			src, err := NewUnstructuredFQDNSource(
				t.Context(),
				dynamicClient,
				kubeClient,
				&Config{
					LabelFilter:           labels.Everything(),
					UnstructuredResources: resources,
					TemplateEngine:        templatetest.MustEngine(t, tt.fqdnTemplate, tt.targetTemplate, tt.fqdnTargetTemplate, false),
				},
			)
			require.NoError(t, err)

			_, err = src.Endpoints(t.Context())
			require.Error(t, err)
		})
	}
}
