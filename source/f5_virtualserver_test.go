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

	"sigs.k8s.io/external-dns/internal/testutils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	fakeDynamic "k8s.io/client-go/dynamic/fake"
	fakeKube "k8s.io/client-go/kubernetes/fake"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/source/annotations"
	"sigs.k8s.io/external-dns/source/types"

	f5 "github.com/F5Networks/k8s-bigip-ctlr/v2/config/apis/cis/v1"
)

const defaultF5VirtualServerNamespace = "virtualserver"

func TestF5VirtualServerEndpoints(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name             string
		annotationFilter string
		labelFilter      labels.Selector
		virtualServer    f5.VirtualServer
		expected         []*endpoint.Endpoint
	}{
		{
			name:             "F5 VirtualServer with target annotation",
			annotationFilter: "",
			virtualServer: f5.VirtualServer{
				TypeMeta: metav1.TypeMeta{
					APIVersion: f5VirtualServerGVR.GroupVersion().String(),
					Kind:       "VirtualServer",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-vs",
					Namespace: defaultF5VirtualServerNamespace,
					UID:       "f5-vs-uid-1234",
					Annotations: map[string]string{
						annotations.TargetKey: "192.168.1.150",
					},
				},
				Spec: f5.VirtualServerSpec{
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
						"resource": "f5-virtualserver/virtualserver/test-vs",
					},
				}).WithRefObject(testutils.RefSource(string(types.F5VirtualServer))),
			},
		},
		{
			name:             "F5 VirtualServer with host and virtualServerAddress set",
			annotationFilter: "",
			virtualServer: f5.VirtualServer{
				TypeMeta: metav1.TypeMeta{
					APIVersion: f5VirtualServerGVR.GroupVersion().String(),
					Kind:       "VirtualServer",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-vs",
					Namespace: defaultF5VirtualServerNamespace,
				},
				Spec: f5.VirtualServerSpec{
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
						"resource": "f5-virtualserver/virtualserver/test-vs",
					},
				},
			},
		},
		{
			name:             "F5 VirtualServer with host set and IP address from the status field",
			annotationFilter: "",
			virtualServer: f5.VirtualServer{
				TypeMeta: metav1.TypeMeta{
					APIVersion: f5VirtualServerGVR.GroupVersion().String(),
					Kind:       "VirtualServer",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-vs",
					Namespace: defaultF5VirtualServerNamespace,
				},
				Spec: f5.VirtualServerSpec{
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
						"resource": "f5-virtualserver/virtualserver/test-vs",
					},
				},
			},
		},
		{
			name:             "F5 VirtualServer with no IP address set",
			annotationFilter: "",
			virtualServer: f5.VirtualServer{
				TypeMeta: metav1.TypeMeta{
					APIVersion: f5VirtualServerGVR.GroupVersion().String(),
					Kind:       "VirtualServer",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-vs",
					Namespace: defaultF5VirtualServerNamespace,
				},
				Spec: f5.VirtualServerSpec{
					Host: "www.example.com",
				},
				Status: f5.CustomResourceStatus{
					VSAddress: "",
				},
			},
			expected: nil,
		},
		{
			name:             "F5 VirtualServer with matching annotation filter",
			annotationFilter: "foo=bar",
			virtualServer: f5.VirtualServer{
				TypeMeta: metav1.TypeMeta{
					APIVersion: f5VirtualServerGVR.GroupVersion().String(),
					Kind:       "VirtualServer",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-vs",
					Namespace: defaultF5VirtualServerNamespace,
					Annotations: map[string]string{
						"foo": "bar",
					},
				},
				Spec: f5.VirtualServerSpec{
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
						"resource": "f5-virtualserver/virtualserver/test-vs",
					},
				},
			},
		},
		{
			name:             "F5 VirtualServer with non-matching annotation filter",
			annotationFilter: "foo=bar",
			virtualServer: f5.VirtualServer{
				TypeMeta: metav1.TypeMeta{
					APIVersion: f5VirtualServerGVR.GroupVersion().String(),
					Kind:       "VirtualServer",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-vs",
					Namespace: defaultF5VirtualServerNamespace,
					Annotations: map[string]string{
						"bar": "foo",
					},
				},
				Spec: f5.VirtualServerSpec{
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
			name: "F5 VirtualServer TTL annotation",
			virtualServer: f5.VirtualServer{
				TypeMeta: metav1.TypeMeta{
					APIVersion: f5VirtualServerGVR.GroupVersion().String(),
					Kind:       "VirtualServer",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-vs",
					Namespace: defaultF5VirtualServerNamespace,
					Annotations: map[string]string{
						"external-dns.kubernetes.io/ttl": "600",
					},
				},
				Spec: f5.VirtualServerSpec{
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
						"resource": "f5-virtualserver/virtualserver/test-vs",
					},
				},
			},
		},
		{
			name: "F5 VirtualServer with error status but valid IP",
			virtualServer: f5.VirtualServer{
				TypeMeta: metav1.TypeMeta{
					APIVersion: f5VirtualServerGVR.GroupVersion().String(),
					Kind:       "VirtualServer",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-vs",
					Namespace: defaultF5VirtualServerNamespace,
					Annotations: map[string]string{
						"external-dns.kubernetes.io/ttl": "600",
					},
				},
				Spec: f5.VirtualServerSpec{
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
						"resource": "f5-virtualserver/virtualserver/test-vs",
					},
				},
			},
		},
		{
			name: "F5 VirtualServer with missing IP address and OK status",
			virtualServer: f5.VirtualServer{
				TypeMeta: metav1.TypeMeta{
					APIVersion: f5VirtualServerGVR.GroupVersion().String(),
					Kind:       "VirtualServer",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-vs",
					Namespace: defaultF5VirtualServerNamespace,
					Annotations: map[string]string{
						"external-dns.kubernetes.io/ttl": "600",
					},
				},
				Spec: f5.VirtualServerSpec{
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
			name: "F5 VirtualServer with hostAliases",
			virtualServer: f5.VirtualServer{
				TypeMeta: metav1.TypeMeta{
					APIVersion: f5VirtualServerGVR.GroupVersion().String(),
					Kind:       "VirtualServer",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-vs",
					Namespace: defaultF5VirtualServerNamespace,
				},
				Spec: f5.VirtualServerSpec{
					Host:                 "www.example.com",
					VirtualServerAddress: "192.168.1.100",
					HostAliases:          []string{"alias1.example.com", "alias2.example.com"},
				},
				Status: f5.CustomResourceStatus{
					VSAddress: "192.168.1.100",
					Status:    "OK",
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "alias1.example.com",
					Targets:    []string{"192.168.1.100"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  0,
					Labels: endpoint.Labels{
						"resource": "f5-virtualserver/virtualserver/test-vs",
					},
				},
				{
					DNSName:    "alias2.example.com",
					Targets:    []string{"192.168.1.100"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  0,
					Labels: endpoint.Labels{
						"resource": "f5-virtualserver/virtualserver/test-vs",
					},
				},
				{
					DNSName:    "www.example.com",
					Targets:    []string{"192.168.1.100"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  0,
					Labels: endpoint.Labels{
						"resource": "f5-virtualserver/virtualserver/test-vs",
					},
				},
			},
		},
		{
			name: "F5 VirtualServer with hostAliases and target annotation",
			virtualServer: f5.VirtualServer{
				TypeMeta: metav1.TypeMeta{
					APIVersion: f5VirtualServerGVR.GroupVersion().String(),
					Kind:       "VirtualServer",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-vs",
					Namespace: defaultF5VirtualServerNamespace,
					Annotations: map[string]string{
						annotations.TargetKey: "192.168.1.150",
					},
				},
				Spec: f5.VirtualServerSpec{
					Host:                 "www.example.com",
					VirtualServerAddress: "192.168.1.100",
					HostAliases:          []string{"alias1.example.com", "alias2.example.com"},
				},
				Status: f5.CustomResourceStatus{
					VSAddress: "192.168.1.100",
					Status:    "OK",
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "www.example.com",
					Targets:    []string{"192.168.1.150"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  0,
					Labels: endpoint.Labels{
						"resource": "f5-virtualserver/virtualserver/test-vs",
					},
				},
				{
					DNSName:    "alias1.example.com",
					Targets:    []string{"192.168.1.150"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  0,
					Labels: endpoint.Labels{
						"resource": "f5-virtualserver/virtualserver/test-vs",
					},
				},
				{
					DNSName:    "alias2.example.com",
					Targets:    []string{"192.168.1.150"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  0,
					Labels: endpoint.Labels{
						"resource": "f5-virtualserver/virtualserver/test-vs",
					},
				},
			},
		},
		{
			name: "F5 VirtualServer with hostAliases and TTL annotation",
			virtualServer: f5.VirtualServer{
				TypeMeta: metav1.TypeMeta{
					APIVersion: f5VirtualServerGVR.GroupVersion().String(),
					Kind:       "VirtualServer",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-vs",
					Namespace: defaultF5VirtualServerNamespace,
					Annotations: map[string]string{
						"external-dns.kubernetes.io/ttl": "300",
					},
				},
				Spec: f5.VirtualServerSpec{
					Host:                 "www.example.com",
					VirtualServerAddress: "192.168.1.100",
					HostAliases:          []string{"alias1.example.com", "alias2.example.com"},
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
					RecordTTL:  300,
					Labels: endpoint.Labels{
						"resource": "f5-virtualserver/virtualserver/test-vs",
					},
				},
				{
					DNSName:    "alias1.example.com",
					Targets:    []string{"192.168.1.100"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  300,
					Labels: endpoint.Labels{
						"resource": "f5-virtualserver/virtualserver/test-vs",
					},
				},
				{
					DNSName:    "alias2.example.com",
					Targets:    []string{"192.168.1.100"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  300,
					Labels: endpoint.Labels{
						"resource": "f5-virtualserver/virtualserver/test-vs",
					},
				},
			},
		},
		{
			name: "F5 VirtualServer with empty hostAliases",
			virtualServer: f5.VirtualServer{
				TypeMeta: metav1.TypeMeta{
					APIVersion: f5VirtualServerGVR.GroupVersion().String(),
					Kind:       "VirtualServer",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-vs",
					Namespace: defaultF5VirtualServerNamespace,
				},
				Spec: f5.VirtualServerSpec{
					Host:                 "www.example.com",
					VirtualServerAddress: "192.168.1.100",
					HostAliases:          []string{},
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
						"resource": "f5-virtualserver/virtualserver/test-vs",
					},
				},
			},
		},
		{
			name: "F5 VirtualServer with hostAliases containing empty strings",
			virtualServer: f5.VirtualServer{
				TypeMeta: metav1.TypeMeta{
					APIVersion: f5VirtualServerGVR.GroupVersion().String(),
					Kind:       "VirtualServer",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-vs",
					Namespace: defaultF5VirtualServerNamespace,
				},
				Spec: f5.VirtualServerSpec{
					Host:                 "www.example.com",
					VirtualServerAddress: "192.168.1.100",
					HostAliases:          []string{"alias1.example.com", "", "alias2.example.com"},
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
						"resource": "f5-virtualserver/virtualserver/test-vs",
					},
				},
				{
					DNSName:    "alias1.example.com",
					Targets:    []string{"192.168.1.100"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  0,
					Labels: endpoint.Labels{
						"resource": "f5-virtualserver/virtualserver/test-vs",
					},
				},
				{
					DNSName:    "alias2.example.com",
					Targets:    []string{"192.168.1.100"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  0,
					Labels: endpoint.Labels{
						"resource": "f5-virtualserver/virtualserver/test-vs",
					},
				},
			},
		},
		{
			name:        "F5 VirtualServer with matching label filter",
			labelFilter: labels.SelectorFromSet(labels.Set{"app": "test"}),
			virtualServer: f5.VirtualServer{
				TypeMeta: metav1.TypeMeta{
					APIVersion: f5VirtualServerGVR.GroupVersion().String(),
					Kind:       "VirtualServer",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-vs",
					Namespace: defaultF5VirtualServerNamespace,
					Labels:    map[string]string{"app": "test"},
				},
				Spec: f5.VirtualServerSpec{
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
					Labels:     endpoint.Labels{"resource": "f5-virtualserver/virtualserver/test-vs"},
				},
			},
		},
		{
			name:        "F5 VirtualServer with non-matching label filter",
			labelFilter: labels.SelectorFromSet(labels.Set{"app": "test"}),
			virtualServer: f5.VirtualServer{
				TypeMeta: metav1.TypeMeta{
					APIVersion: f5VirtualServerGVR.GroupVersion().String(),
					Kind:       "VirtualServer",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-vs",
					Namespace: defaultF5VirtualServerNamespace,
					Labels:    map[string]string{"app": "other"},
				},
				Spec: f5.VirtualServerSpec{
					Host:                 "www.example.com",
					VirtualServerAddress: "192.168.1.100",
				},
				Status: f5.CustomResourceStatus{VSAddress: "192.168.1.100"},
			},
			expected: nil,
		},
		{
			name: "F5 VirtualServer does not support provider-specific annotations",
			virtualServer: f5.VirtualServer{
				TypeMeta: metav1.TypeMeta{
					APIVersion: f5VirtualServerGVR.GroupVersion().String(),
					Kind:       "VirtualServer",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-vs",
					Namespace: defaultF5VirtualServerNamespace,
					Annotations: map[string]string{
						annotations.AWSPrefix + "weight": "10",
					},
				},
				Spec: f5.VirtualServerSpec{
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
						"resource": "f5-virtualserver/virtualserver/test-vs",
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			fakeKubernetesClient := fakeKube.NewClientset()
			scheme := runtime.NewScheme()
			scheme.AddKnownTypes(f5VirtualServerGVR.GroupVersion(), &f5.VirtualServer{}, &f5.VirtualServerList{})
			fakeDynamicClient := fakeDynamic.NewSimpleDynamicClient(scheme)

			virtualServer := unstructured.Unstructured{}

			virtualServerJSON, err := json.Marshal(tc.virtualServer)
			require.NoError(t, err)
			assert.NoError(t, virtualServer.UnmarshalJSON(virtualServerJSON))

			// Create VirtualServer resources
			_, err = fakeDynamicClient.Resource(f5VirtualServerGVR).Namespace(defaultF5VirtualServerNamespace).Create(t.Context(), &virtualServer, metav1.CreateOptions{})
			assert.NoError(t, err)

			labelFilter := tc.labelFilter
			if labelFilter == nil {
				labelFilter = labels.Everything()
			}
			source, err := NewF5VirtualServerSource(t.Context(), fakeDynamicClient, fakeKubernetesClient,
				&Config{
					Namespace:        defaultF5VirtualServerNamespace,
					AnnotationFilter: parseAnnotationFilterOrNil(tc.annotationFilter),
					LabelFilter:      labelFilter,
				})
			require.NoError(t, err)
			assert.NotNil(t, source)

			count := &unstructured.UnstructuredList{}
			for len(count.Items) < 1 {
				count, _ = fakeDynamicClient.Resource(f5VirtualServerGVR).Namespace(defaultF5VirtualServerNamespace).List(t.Context(), metav1.ListOptions{})
			}

			endpoints, err := source.Endpoints(t.Context())
			require.NoError(t, err)
			assert.Len(t, endpoints, len(tc.expected))
			testutils.ValidateEndpoints(t, endpoints, tc.expected)
		})
	}
}

func TestF5VirtualServerSource_InformerTransform(t *testing.T) {
	t.Parallel()

	uc, err := newVSUnstructuredConverter()
	require.NoError(t, err)

	fakeClient := fakeKube.NewClientset()
	fakeDynamicClient := fakeDynamic.NewSimpleDynamicClient(uc.scheme)

	source, err := NewF5VirtualServerSource(t.Context(), fakeDynamicClient, fakeClient, &Config{LabelFilter: labels.Everything()})
	require.NoError(t, err)
	require.IsType(t, &f5VirtualServerSource{}, source)

	testDynamicInformerTransformHelper(t,
		f5VirtualServerGVR,
		fakeDynamicClient,
		source.(*f5VirtualServerSource).virtualServerInformer,
		withRemovedLastAppliedConfigAnnotation(),
		withRemovedManagedFields(),
	)
}

// TestF5VirtualServerIndexer verifies that the VirtualServer indexer correctly filters resources
// by annotation filter and label selector at index time, so that only matching resources are
// returned by Endpoints().
func TestF5VirtualServerIndexer(t *testing.T) {
	t.Parallel()

	makeEntity := func(name, vsAddress string, ann, lbls map[string]string) *f5.VirtualServer {
		if ann == nil {
			ann = map[string]string{}
		}
		if lbls == nil {
			lbls = map[string]string{}
		}
		return &f5.VirtualServer{
			TypeMeta: metav1.TypeMeta{
				APIVersion: f5VirtualServerGVR.GroupVersion().String(),
				Kind:       "VirtualServer",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:        name,
				Namespace:   defaultF5VirtualServerNamespace,
				Annotations: ann,
				Labels:      lbls,
			},
			Spec: f5.VirtualServerSpec{
				Host:                 name + ".example.org",
				VirtualServerAddress: vsAddress,
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
		servers          []*f5.VirtualServer
		expectedCount    int
	}{
		{
			name:          "no filters returns all servers",
			expectedCount: 3,
			servers: []*f5.VirtualServer{
				makeEntity("vs1", "1.2.3.1", nil, nil),
				makeEntity("vs2", "1.2.3.2", nil, nil),
				makeEntity("vs3", "1.2.3.3", nil, nil),
			},
		},
		{
			name:             "annotation filter includes matching servers",
			annotationFilter: "tier=frontend",
			expectedCount:    2,
			servers: []*f5.VirtualServer{
				makeEntity("vs1", "1.2.3.1", map[string]string{"tier": "frontend"}, nil),
				makeEntity("vs2", "1.2.3.2", map[string]string{"tier": "frontend"}, nil),
				makeEntity("vs3", "1.2.3.3", map[string]string{"tier": "backend"}, nil),
			},
		},
		{
			name:          "label filter includes matching servers",
			labelFilter:   "env=prod",
			expectedCount: 1,
			servers: []*f5.VirtualServer{
				makeEntity("vs1", "1.2.3.1", nil, map[string]string{"env": "prod"}),
				makeEntity("vs2", "1.2.3.2", nil, map[string]string{"env": "staging"}),
				makeEntity("vs3", "1.2.3.3", nil, nil),
			},
		},
		{
			name:             "annotation and label filter combined",
			annotationFilter: "tier=frontend",
			labelFilter:      "env=prod",
			expectedCount:    1,
			servers: []*f5.VirtualServer{
				makeEntity("vs1", "1.2.3.1", map[string]string{"tier": "frontend"}, map[string]string{"env": "prod"}),
				makeEntity("vs2", "1.2.3.2", map[string]string{"tier": "frontend"}, map[string]string{"env": "staging"}),
				makeEntity("vs3", "1.2.3.3", map[string]string{"tier": "backend"}, map[string]string{"env": "prod"}),
			},
		},
		{
			name:             "no matches returns empty",
			annotationFilter: "tier=missing",
			expectedCount:    0,
			servers: []*f5.VirtualServer{
				makeEntity("vs1", "1.2.3.1", map[string]string{"tier": "frontend"}, nil),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			uc, err := newVSUnstructuredConverter()
			require.NoError(t, err)

			fakeKubernetesClient := fakeKube.NewSimpleClientset()
			fakeDynamicClient := fakeDynamic.NewSimpleDynamicClient(uc.scheme)

			for _, srv := range tt.servers {
				data, err := json.Marshal(srv)
				require.NoError(t, err)
				obj := unstructured.Unstructured{}
				require.NoError(t, obj.UnmarshalJSON(data))
				_, err = fakeDynamicClient.Resource(f5VirtualServerGVR).Namespace(defaultF5VirtualServerNamespace).Create(t.Context(), &obj, metav1.CreateOptions{})
				require.NoError(t, err)
			}

			labelSel := labels.Everything()
			if tt.labelFilter != "" {
				labelSel, err = labels.Parse(tt.labelFilter)
				require.NoError(t, err)
			}

			src, err := NewF5VirtualServerSource(t.Context(), fakeDynamicClient, fakeKubernetesClient, &Config{
				Namespace:        defaultF5VirtualServerNamespace,
				AnnotationFilter: parseAnnotationFilterOrNil(tt.annotationFilter),
				LabelFilter:      labelSel,
			})
			require.NoError(t, err)

			endpoints, err := src.Endpoints(t.Context())
			require.NoError(t, err)
			assert.Len(t, endpoints, tt.expectedCount)
		})
	}
}
