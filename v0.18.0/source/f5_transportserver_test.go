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

const defaultF5TransportServerNamespace = "transportserver"

func TestF5TransportServerEndpoints(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name             string
		annotationFilter string
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
					Annotations: map[string]string{
						targetAnnotationKey: "192.168.1.150",
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
				{
					DNSName:    "www.example.com",
					Targets:    []string{"192.168.1.150"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  0,
					Labels: endpoint.Labels{
						"resource": "f5-transportserver/transportserver/test-vs",
					},
				},
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
						"external-dns.alpha.kubernetes.io/ttl": "600",
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
						"external-dns.alpha.kubernetes.io/ttl": "600",
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
						"external-dns.alpha.kubernetes.io/ttl": "600",
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
			assert.NoError(t, transportServer.UnmarshalJSON(transportServerJSON))

			// Create TransportServer resources
			_, err = fakeDynamicClient.Resource(f5TransportServerGVR).Namespace(defaultF5TransportServerNamespace).Create(context.Background(), &transportServer, metav1.CreateOptions{})
			assert.NoError(t, err)

			source, err := NewF5TransportServerSource(context.TODO(), fakeDynamicClient, fakeKubernetesClient, defaultF5TransportServerNamespace, tc.annotationFilter)
			require.NoError(t, err)
			assert.NotNil(t, source)

			count := &unstructured.UnstructuredList{}
			for len(count.Items) < 1 {
				count, _ = fakeDynamicClient.Resource(f5TransportServerGVR).Namespace(defaultF5TransportServerNamespace).List(context.Background(), metav1.ListOptions{})
			}

			endpoints, err := source.Endpoints(context.Background())
			require.NoError(t, err)
			assert.Len(t, endpoints, len(tc.expected))
			assert.Equal(t, tc.expected, endpoints)
		})
	}
}
