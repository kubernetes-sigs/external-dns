/*
Copyright 2022 The Kubernetes Authors.

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
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	fakeDynamic "k8s.io/client-go/dynamic/fake"
	fakeKube "k8s.io/client-go/kubernetes/fake"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	"sigs.k8s.io/external-dns/source/annotations"
	"sigs.k8s.io/external-dns/source/types"

	f5 "github.com/F5Networks/k8s-bigip-ctlr/v2/config/apis/cis/v1"
)

const defaultF5TransportServerNamespace = "transportserver"

func TestF5TransportServerEndpoints(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name             string
		annotationFilter string
		labelFilter      labels.Selector
		transportServer  f5.TransportServer
		expected         []*endpoint.Endpoint
	}{
		{
			name:             "F5 TransportServer with target annotation",
			annotationFilter: "",
			transportServer: f5.TransportServer{
				TypeMeta: metav1.TypeMeta{
					APIVersion: f5TransportServerGVR.GroupVersion().String(),
					Kind:       "TransportServer",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-vs",
					Namespace: defaultF5TransportServerNamespace,
					UID:       "f5-ts-uid-1234",
					Annotations: map[string]string{
						annotations.TargetKey: "192.168.1.150",
					},
				},
				Spec: f5.TransportServerSpec{
					Host:                 "www.example.com",
					VirtualServerAddress: "192.168.1.100",
				},
				Status: f5.CustomResourceStatus{
					VSAddress: "192.168.1.200",
					Status:    "OK",
				},
			},
			expected: []*endpoint.Endpoint{
				(&endpoint.Endpoint{
					DNSName:    "www.example.com",
					Targets:    []string{"192.168.1.150"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  0,
					Labels: endpoint.Labels{
						"resource": "f5-transportserver/transportserver/test-vs",
					},
				}).WithRefObject(testutils.RefSource(string(types.F5TransportServer))),
			},
		},
		{
			name:             "F5 TransportServer with host and VirtualServerAddress set",
			annotationFilter: "",
			transportServer: f5.TransportServer{
				TypeMeta: metav1.TypeMeta{
					APIVersion: f5TransportServerGVR.GroupVersion().String(),
					Kind:       "TransportServer",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-vs",
					Namespace: defaultF5TransportServerNamespace,
				},
				Spec: f5.TransportServerSpec{
					Host:                 "www.example.com",
					VirtualServerAddress: "192.168.1.100",
				},
				Status: f5.CustomResourceStatus{
					VSAddress: "192.168.1.200",
					Status:    "OK",
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "www.example.com",
					Targets:    []string{"192.168.1.100"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  0,
					Labels: endpoint.Labels{
						"resource": "f5-transportserver/transportserver/test-vs",
					},
				},
			},
		},
		{
			name:             "F5 TransportServer with host set and IP address from the status field",
			annotationFilter: "",
			transportServer: f5.TransportServer{
				TypeMeta: metav1.TypeMeta{
					APIVersion: f5TransportServerGVR.GroupVersion().String(),
					Kind:       "TransportServer",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-vs",
					Namespace: defaultF5TransportServerNamespace,
				},
				Spec: f5.TransportServerSpec{
					Host: "www.example.com",
				},
				Status: f5.CustomResourceStatus{
					VSAddress: "192.168.1.100",
					Status:    "OK",
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "www.example.com",
					Targets:    []string{"192.168.1.100"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  0,
					Labels: endpoint.Labels{
						"resource": "f5-transportserver/transportserver/test-vs",
					},
				},
			},
		},
		{
			name:             "F5 TransportServer with no IP address set",
			annotationFilter: "",
			transportServer: f5.TransportServer{
				TypeMeta: metav1.TypeMeta{
					APIVersion: f5TransportServerGVR.GroupVersion().String(),
					Kind:       "TransportServer",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-vs",
					Namespace: defaultF5TransportServerNamespace,
				},
				Spec: f5.TransportServerSpec{
					Host: "www.example.com",
				},
				Status: f5.CustomResourceStatus{
					VSAddress: "",
				},
			},
			expected: nil,
		},
		{
			name:             "F5 TransportServer with matching annotation filter",
			annotationFilter: "foo=bar",
			transportServer: f5.TransportServer{
				TypeMeta: metav1.TypeMeta{
					APIVersion: f5TransportServerGVR.GroupVersion().String(),
					Kind:       "TransportServer",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-vs",
					Namespace: defaultF5TransportServerNamespace,
					Annotations: map[string]string{
						"foo": "bar",
					},
				},
				Spec: f5.TransportServerSpec{
					Host:                 "www.example.com",
					VirtualServerAddress: "192.168.1.100",
				},
				Status: f5.CustomResourceStatus{
					VSAddress: "192.168.1.100",
					Status:    "OK",
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "www.example.com",
					Targets:    []string{"192.168.1.100"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  0,
					Labels: endpoint.Labels{
						"resource": "f5-transportserver/transportserver/test-vs",
					},
				},
			},
		},
		{
			name:             "F5 TransportServer with non-matching annotation filter",
			annotationFilter: "foo=bar",
			transportServer: f5.TransportServer{
				TypeMeta: metav1.TypeMeta{
					APIVersion: f5TransportServerGVR.GroupVersion().String(),
					Kind:       "TransportServer",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-vs",
					Namespace: defaultF5TransportServerNamespace,
					Annotations: map[string]string{
						"bar": "foo",
					},
				},
				Spec: f5.TransportServerSpec{
					Host:                 "www.example.com",
					VirtualServerAddress: "192.168.1.100",
				},
				Status: f5.CustomResourceStatus{
					VSAddress: "192.168.1.100",
					Status:    "OK",
				},
			},
			expected: nil,
		},
		{
			name: "F5 TransportServer TTL annotation",
			transportServer: f5.TransportServer{
				TypeMeta: metav1.TypeMeta{
					APIVersion: f5TransportServerGVR.GroupVersion().String(),
					Kind:       "TransportServer",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-vs",
					Namespace: defaultF5TransportServerNamespace,
					Annotations: map[string]string{
						"external-dns.kubernetes.io/ttl": "600",
					},
				},
				Spec: f5.TransportServerSpec{
					Host:                 "www.example.com",
					VirtualServerAddress: "192.168.1.100",
				},
				Status: f5.CustomResourceStatus{
					VSAddress: "192.168.1.100",
					Status:    "OK",
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "www.example.com",
					Targets:    []string{"192.168.1.100"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  600,
					Labels: endpoint.Labels{
						"resource": "f5-transportserver/transportserver/test-vs",
					},
				},
			},
		},
		{
			name: "F5 TransportServer with error status but valid IP",
			transportServer: f5.TransportServer{
				TypeMeta: metav1.TypeMeta{
					APIVersion: f5TransportServerGVR.GroupVersion().String(),
					Kind:       "TransportServer",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-ts",
					Namespace: defaultF5TransportServerNamespace,
					Annotations: map[string]string{
						"external-dns.kubernetes.io/ttl": "600",
					},
				},
				Spec: f5.TransportServerSpec{
					Host:                 "www.example.com",
					VirtualServerAddress: "192.168.1.100",
				},
				Status: f5.CustomResourceStatus{
					VSAddress: "192.168.1.100",
					Status:    "ERROR",
					Error:     "Some error status message",
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "www.example.com",
					Targets:    []string{"192.168.1.100"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  600,
					Labels: endpoint.Labels{
						"resource": "f5-transportserver/transportserver/test-ts",
					},
				},
			},
		},
		{
			name: "F5 TransportServer with missing IP address and OK status",
			transportServer: f5.TransportServer{
				TypeMeta: metav1.TypeMeta{
					APIVersion: f5TransportServerGVR.GroupVersion().String(),
					Kind:       "TransportServer",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-ts",
					Namespace: defaultF5TransportServerNamespace,
					Annotations: map[string]string{
						"external-dns.kubernetes.io/ttl": "600",
					},
				},
				Spec: f5.TransportServerSpec{
					Host:      "www.example.com",
					IPAMLabel: "test",
				},
				Status: f5.CustomResourceStatus{
					VSAddress: "None",
					Status:    "OK",
				},
			},
			expected: nil,
		},
		{
			name:        "F5 TransportServer with matching label filter",
			labelFilter: labels.SelectorFromSet(labels.Set{"app": "test"}),
			transportServer: f5.TransportServer{
				TypeMeta: metav1.TypeMeta{
					APIVersion: f5TransportServerGVR.GroupVersion().String(),
					Kind:       "TransportServer",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-vs",
					Namespace: defaultF5TransportServerNamespace,
					Labels:    map[string]string{"app": "test"},
				},
				Spec: f5.TransportServerSpec{
					Host:                 "www.example.com",
					VirtualServerAddress: "192.168.1.100",
				},
				Status: f5.CustomResourceStatus{VSAddress: "192.168.1.100"},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "www.example.com",
					Targets:    []string{"192.168.1.100"},
					RecordType: endpoint.RecordTypeA,
					Labels:     endpoint.Labels{"resource": "f5-transportserver/transportserver/test-vs"},
				},
			},
		},
		{
			name:        "F5 TransportServer with non-matching label filter",
			labelFilter: labels.SelectorFromSet(labels.Set{"app": "test"}),
			transportServer: f5.TransportServer{
				TypeMeta: metav1.TypeMeta{
					APIVersion: f5TransportServerGVR.GroupVersion().String(),
					Kind:       "TransportServer",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-vs",
					Namespace: defaultF5TransportServerNamespace,
					Labels:    map[string]string{"app": "other"},
				},
				Spec: f5.TransportServerSpec{
					Host:                 "www.example.com",
					VirtualServerAddress: "192.168.1.100",
				},
				Status: f5.CustomResourceStatus{VSAddress: "192.168.1.100"},
			},
			expected: nil,
		},
		{
			name: "F5 TransportServer does not support provider-specific annotations",
			transportServer: f5.TransportServer{
				TypeMeta: metav1.TypeMeta{
					APIVersion: f5TransportServerGVR.GroupVersion().String(),
					Kind:       "TransportServer",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-vs",
					Namespace: defaultF5TransportServerNamespace,
					Annotations: map[string]string{
						annotations.AWSPrefix + "weight": "10",
					},
				},
				Spec: f5.TransportServerSpec{
					Host:                 "www.example.com",
					VirtualServerAddress: "192.168.1.100",
				},
				Status: f5.CustomResourceStatus{
					VSAddress: "192.168.1.100",
					Status:    "OK",
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "www.example.com",
					Targets:    []string{"192.168.1.100"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  0,
					Labels: endpoint.Labels{
						"resource": "f5-transportserver/transportserver/test-vs",
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			fakeKubernetesClient := fakeKube.NewSimpleClientset()
			scheme := runtime.NewScheme()
			scheme.AddKnownTypes(f5TransportServerGVR.GroupVersion(), &f5.TransportServer{}, &f5.TransportServerList{})
			fakeDynamicClient := fakeDynamic.NewSimpleDynamicClient(scheme)

			transportServer := unstructured.Unstructured{}

			transportServerJSON, err := json.Marshal(tc.transportServer)
			require.NoError(t, err)
			require.NoError(t, transportServer.UnmarshalJSON(transportServerJSON))

			// Create TransportServer resources
			_, err = fakeDynamicClient.Resource(f5TransportServerGVR).Namespace(defaultF5TransportServerNamespace).Create(t.Context(), &transportServer, metav1.CreateOptions{})
			require.NoError(t, err)

			labelFilter := tc.labelFilter
			if labelFilter == nil {
				labelFilter = labels.Everything()
			}
			source, err := NewF5TransportServerSource(t.Context(), fakeDynamicClient, fakeKubernetesClient,
				&Config{
					Namespace:        defaultF5TransportServerNamespace,
					AnnotationFilter: parseAnnotationFilterOrNil(tc.annotationFilter),
					LabelFilter:      labelFilter,
				})
			require.NoError(t, err)
			require.NotNil(t, source)

			count := &unstructured.UnstructuredList{}
			for len(count.Items) < 1 {
				count, _ = fakeDynamicClient.Resource(f5TransportServerGVR).Namespace(defaultF5TransportServerNamespace).List(t.Context(), metav1.ListOptions{})
			}

			endpoints, err := source.Endpoints(t.Context())
			require.NoError(t, err)
			assert.Len(t, endpoints, len(tc.expected))
			testutils.ValidateEndpoints(t, endpoints, tc.expected)
		})
	}
}

func TestF5TransportServerSource_InformerTransform(t *testing.T) {
	t.Parallel()

	uc, err := newTSUnstructuredConverter()
	require.NoError(t, err)

	fakeKubernetesClient := fakeKube.NewSimpleClientset()
	fakeDynamicClient := fakeDynamic.NewSimpleDynamicClient(uc.scheme)

	source, err := NewF5TransportServerSource(t.Context(), fakeDynamicClient, fakeKubernetesClient, &Config{LabelFilter: labels.Everything()})
	require.NoError(t, err)
	require.IsType(t, &f5TransportServerSource{}, source)

	testDynamicInformerTransformHelper(t,
		f5TransportServerGVR,
		fakeDynamicClient,
		source.(*f5TransportServerSource).transportServerInformer,
		withRemovedLastAppliedConfigAnnotation(),
		withRemovedManagedFields(),
	)
}

// TestF5TransportServerIndexer verifies that the TransportServer indexer correctly filters resources
// by annotation filter and label selector at index time, so that only matching resources are
// returned by Endpoints().
func TestF5TransportServerIndexer(t *testing.T) {
	t.Parallel()

	makeEntity := func(name, vsAddress string, ann, lbls map[string]string) *f5.TransportServer {
		if ann == nil {
			ann = map[string]string{}
		}
		if lbls == nil {
			lbls = map[string]string{}
		}
		return &f5.TransportServer{
			TypeMeta: metav1.TypeMeta{
				APIVersion: f5TransportServerGVR.GroupVersion().String(),
				Kind:       "TransportServer",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:        name,
				Namespace:   defaultF5TransportServerNamespace,
				Annotations: ann,
				Labels:      lbls,
			},
			Spec: f5.TransportServerSpec{
				Host: name + ".example.org",
			},
			Status: f5.CustomResourceStatus{
				VSAddress: vsAddress,
			},
		}
	}

	tests := []struct {
		name             string
		annotationFilter string
		labelFilter      string
		servers          []*f5.TransportServer
		expectedCount    int
	}{
		{
			name:          "no filters returns all servers",
			expectedCount: 3,
			servers: []*f5.TransportServer{
				makeEntity("ts1", "1.2.3.1", nil, nil),
				makeEntity("ts2", "1.2.3.2", nil, nil),
				makeEntity("ts3", "1.2.3.3", nil, nil),
			},
		},
		{
			name:             "annotation filter includes matching servers",
			annotationFilter: "tier=frontend",
			expectedCount:    2,
			servers: []*f5.TransportServer{
				makeEntity("ts1", "1.2.3.1", map[string]string{"tier": "frontend"}, nil),
				makeEntity("ts2", "1.2.3.2", map[string]string{"tier": "frontend"}, nil),
				makeEntity("ts3", "1.2.3.3", map[string]string{"tier": "backend"}, nil),
			},
		},
		{
			name:          "label filter includes matching servers",
			labelFilter:   "env=prod",
			expectedCount: 1,
			servers: []*f5.TransportServer{
				makeEntity("ts1", "1.2.3.1", nil, map[string]string{"env": "prod"}),
				makeEntity("ts2", "1.2.3.2", nil, map[string]string{"env": "staging"}),
				makeEntity("ts3", "1.2.3.3", nil, nil),
			},
		},
		{
			name:             "annotation and label filter combined",
			annotationFilter: "tier=frontend",
			labelFilter:      "env=prod",
			expectedCount:    1,
			servers: []*f5.TransportServer{
				makeEntity("ts1", "1.2.3.1", map[string]string{"tier": "frontend"}, map[string]string{"env": "prod"}),
				makeEntity("ts2", "1.2.3.2", map[string]string{"tier": "frontend"}, map[string]string{"env": "staging"}),
				makeEntity("ts3", "1.2.3.3", map[string]string{"tier": "backend"}, map[string]string{"env": "prod"}),
			},
		},
		{
			name:             "no matches returns empty",
			annotationFilter: "tier=missing",
			expectedCount:    0,
			servers: []*f5.TransportServer{
				makeEntity("ts1", "1.2.3.1", map[string]string{"tier": "frontend"}, nil),
			},
		},
		{
			name:          "controller mismatch is excluded",
			expectedCount: 0,
			servers: []*f5.TransportServer{
				makeEntity("ts1", "1.2.3.1", map[string]string{annotations.ControllerKey: "other-controller"}, nil),
			},
		},
	}

	uc, err := newTSUnstructuredConverter()
	require.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			fakeKubernetesClient := fakeKube.NewSimpleClientset()
			fakeDynamicClient := fakeDynamic.NewSimpleDynamicClient(uc.scheme)

			for _, srv := range tt.servers {
				data, err := json.Marshal(srv)
				require.NoError(t, err)
				obj := unstructured.Unstructured{}
				require.NoError(t, obj.UnmarshalJSON(data))
				_, err = fakeDynamicClient.Resource(f5TransportServerGVR).Namespace(defaultF5TransportServerNamespace).Create(t.Context(), &obj, metav1.CreateOptions{})
				require.NoError(t, err)
			}

			src, err := NewF5TransportServerSource(t.Context(), fakeDynamicClient, fakeKubernetesClient, &Config{
				Namespace:        defaultF5TransportServerNamespace,
				AnnotationFilter: parseAnnotationFilterOrNil(tt.annotationFilter),
				LabelFilter:      parseLabelSelectorOrEverything(t, tt.labelFilter),
			})
			require.NoError(t, err)

			endpoints, err := src.Endpoints(t.Context())
			require.NoError(t, err)
			assert.Len(t, endpoints, tt.expectedCount)
		})
	}
}
