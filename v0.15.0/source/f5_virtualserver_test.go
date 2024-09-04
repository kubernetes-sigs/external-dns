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
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	fakeDynamic "k8s.io/client-go/dynamic/fake"
	fakeKube "k8s.io/client-go/kubernetes/fake"
	"sigs.k8s.io/external-dns/endpoint"

	f5 "github.com/F5Networks/k8s-bigip-ctlr/v2/config/apis/cis/v1"
)

const defaultF5VirtualServerNamespace = "virtualserver"

func TestF5VirtualServerEndpoints(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name             string
		annotationFilter string
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
					Annotations: map[string]string{
						targetAnnotationKey: "192.168.1.150",
					},
				},
				Spec: f5.VirtualServerSpec{
					Host:                 "www.example.com",
					VirtualServerAddress: "192.168.1.100",
				},
				Status: f5.VirtualServerStatus{
					VSAddress: "192.168.1.200",
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
				Status: f5.VirtualServerStatus{
					VSAddress: "192.168.1.200",
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
				Status: f5.VirtualServerStatus{
					VSAddress: "192.168.1.100",
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
				Status: f5.VirtualServerStatus{
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
						"external-dns.alpha.kubernetes.io/ttl": "600",
					},
				},
				Spec: f5.VirtualServerSpec{
					Host:                 "www.example.com",
					VirtualServerAddress: "192.168.1.100",
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
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			fakeKubernetesClient := fakeKube.NewSimpleClientset()
			scheme := runtime.NewScheme()
			scheme.AddKnownTypes(f5VirtualServerGVR.GroupVersion(), &f5.VirtualServer{}, &f5.VirtualServerList{})
			fakeDynamicClient := fakeDynamic.NewSimpleDynamicClient(scheme)

			virtualServer := unstructured.Unstructured{}

			virtualServerJSON, err := json.Marshal(tc.virtualServer)
			require.NoError(t, err)
			assert.NoError(t, virtualServer.UnmarshalJSON(virtualServerJSON))

			// Create VirtualServer resources
			_, err = fakeDynamicClient.Resource(f5VirtualServerGVR).Namespace(defaultF5VirtualServerNamespace).Create(context.Background(), &virtualServer, metav1.CreateOptions{})
			assert.NoError(t, err)

			source, err := NewF5VirtualServerSource(context.TODO(), fakeDynamicClient, fakeKubernetesClient, defaultF5VirtualServerNamespace, tc.annotationFilter)
			require.NoError(t, err)
			assert.NotNil(t, source)

			count := &unstructured.UnstructuredList{}
			for len(count.Items) < 1 {
				count, _ = fakeDynamicClient.Resource(f5VirtualServerGVR).Namespace(defaultF5VirtualServerNamespace).List(context.Background(), metav1.ListOptions{})
			}

			endpoints, err := source.Endpoints(context.Background())
			require.NoError(t, err)
			assert.Len(t, endpoints, len(tc.expected))
			assert.Equal(t, endpoints, tc.expected)
		})
	}
}
