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

package kubeclient

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/tools/clientcmd"

	apiv1alpha1 "sigs.k8s.io/external-dns/apis/v1alpha1"
)

func createTestKubeConfig(t *testing.T, serverURL string) string {
	t.Helper()

	mockKubeCfgDir := filepath.Join(t.TempDir(), ".kube")
	mockKubeCfgPath := filepath.Join(mockKubeCfgDir, "config")
	err := os.MkdirAll(mockKubeCfgDir, 0755)
	require.NoError(t, err)

	kubeCfgTemplate := `apiVersion: v1
kind: Config
clusters:
- cluster:
    server: %s
  name: test-cluster
contexts:
- context:
    cluster: test-cluster
    user: test-user
  name: test-context
current-context: test-context
users:
- name: test-user
  user:
    token: fake-token
`
	err = os.WriteFile(mockKubeCfgPath, fmt.Appendf(nil, kubeCfgTemplate, serverURL), 0644)
	require.NoError(t, err)

	return mockKubeCfgPath
}

func TestNewCRDClientForAPIVersionKind_Success(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {
	}))
	defer svr.Close()

	kubeConfigPath := createTestKubeConfig(t, svr.URL)

	fakeClient := fake.NewClientset()
	fakeClient.Resources = []*metav1.APIResourceList{
		{
			GroupVersion: apiv1alpha1.GroupVersion.String(),
			APIResources: []metav1.APIResource{
				{
					Name:       "dnsendpoints",
					Kind:       apiv1alpha1.DNSEndpointKind,
					Namespaced: true,
				},
			},
		},
	}

	restClient, err := NewCRDClientForAPIVersionKind(
		fakeClient,
		kubeConfigPath,
		"",
		apiv1alpha1.GroupVersion.String(),
		apiv1alpha1.DNSEndpointKind,
	)

	require.NoError(t, err)
	require.NotNil(t, restClient)
}

func TestNewCRDClientForAPIVersionKind_InvalidAPIVersion(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {}))
	defer svr.Close()

	kubeConfigPath := createTestKubeConfig(t, svr.URL)
	fakeClient := fake.NewClientset()

	_, err := NewCRDClientForAPIVersionKind(
		fakeClient,
		kubeConfigPath,
		"",
		"invalid-api-version-format",
		"DNSEndpoint",
	)

	require.Error(t, err)
}

func TestNewCRDClientForAPIVersionKind_GroupVersionNotFound(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {}))
	defer svr.Close()

	kubeConfigPath := createTestKubeConfig(t, svr.URL)

	fakeClient := fake.NewClientset()
	// Empty resources - the requested GroupVersion won't be found
	fakeClient.Resources = []*metav1.APIResourceList{}

	_, err := NewCRDClientForAPIVersionKind(
		fakeClient,
		kubeConfigPath,
		"",
		apiv1alpha1.GroupVersion.String(),
		apiv1alpha1.DNSEndpointKind,
	)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "error listing resources in GroupVersion")
}

func TestNewCRDClientForAPIVersionKind_KindNotFound(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {}))
	defer svr.Close()

	kubeConfigPath := createTestKubeConfig(t, svr.URL)

	fakeClient := fake.NewClientset()
	fakeClient.Resources = []*metav1.APIResourceList{
		{
			GroupVersion: apiv1alpha1.GroupVersion.String(),
			APIResources: []metav1.APIResource{
				{
					Name:       "otherendpoints",
					Kind:       "OtherEndpoint",
					Namespaced: true,
				},
			},
		},
	}

	_, err := NewCRDClientForAPIVersionKind(
		fakeClient,
		kubeConfigPath,
		"",
		apiv1alpha1.GroupVersion.String(),
		apiv1alpha1.DNSEndpointKind,
	)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "unable to find Resource Kind")
	assert.Contains(t, err.Error(), apiv1alpha1.DNSEndpointKind)
}

func TestNewCRDClientForAPIVersionKind_UsesRecommendedHomeFile(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {}))
	defer svr.Close()

	kubeConfigPath := createTestKubeConfig(t, svr.URL)

	prevRecommendedHomeFile := clientcmd.RecommendedHomeFile
	t.Cleanup(func() {
		clientcmd.RecommendedHomeFile = prevRecommendedHomeFile
	})
	clientcmd.RecommendedHomeFile = kubeConfigPath

	fakeClient := fake.NewClientset()
	fakeClient.Resources = []*metav1.APIResourceList{
		{
			GroupVersion: apiv1alpha1.GroupVersion.String(),
			APIResources: []metav1.APIResource{
				{
					Name:       "dnsendpoints",
					Kind:       apiv1alpha1.DNSEndpointKind,
					Namespaced: true,
				},
			},
		},
	}

	// Pass empty kubeConfig to trigger RecommendedHomeFile usage
	restClient, err := NewCRDClientForAPIVersionKind(
		fakeClient,
		"",
		"",
		apiv1alpha1.GroupVersion.String(),
		apiv1alpha1.DNSEndpointKind,
	)

	require.NoError(t, err)
	require.NotNil(t, restClient)
}

func TestNewCRDClientForAPIVersionKind_WithAPIServerURL(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {}))
	defer svr.Close()

	kubeConfigPath := createTestKubeConfig(t, svr.URL)

	fakeClient := fake.NewClientset()
	fakeClient.Resources = []*metav1.APIResourceList{
		{
			GroupVersion: apiv1alpha1.GroupVersion.String(),
			APIResources: []metav1.APIResource{
				{
					Name:       "dnsendpoints",
					Kind:       apiv1alpha1.DNSEndpointKind,
					Namespaced: true,
				},
			},
		},
	}

	restClient, err := NewCRDClientForAPIVersionKind(
		fakeClient,
		kubeConfigPath,
		svr.URL,
		apiv1alpha1.GroupVersion.String(),
		apiv1alpha1.DNSEndpointKind,
	)

	require.NoError(t, err)
	require.NotNil(t, restClient)
}

func TestNewCRDClientForAPIVersionKind_InvalidKubeConfig(t *testing.T) {
	fakeClient := fake.NewClientset()

	_, err := NewCRDClientForAPIVersionKind(
		fakeClient,
		"/nonexistent/path/to/kubeconfig",
		"",
		apiv1alpha1.GroupVersion.String(),
		apiv1alpha1.DNSEndpointKind,
	)

	require.Error(t, err)
}

func TestNewCRDClientForAPIVersionKind_MultipleResourcesFindsCorrectKind(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {}))
	defer svr.Close()

	kubeConfigPath := createTestKubeConfig(t, svr.URL)

	fakeClient := fake.NewClientset()
	fakeClient.Resources = []*metav1.APIResourceList{
		{
			GroupVersion: apiv1alpha1.GroupVersion.String(),
			APIResources: []metav1.APIResource{
				{
					Name:       "otherendpoints",
					Kind:       "OtherEndpoint",
					Namespaced: true,
				},
				{
					Name:       "dnsendpoints",
					Kind:       apiv1alpha1.DNSEndpointKind,
					Namespaced: true,
				},
				{
					Name:       "moreendpoints",
					Kind:       "MoreEndpoint",
					Namespaced: true,
				},
			},
		},
	}

	restClient, err := NewCRDClientForAPIVersionKind(
		fakeClient,
		kubeConfigPath,
		"",
		apiv1alpha1.GroupVersion.String(),
		apiv1alpha1.DNSEndpointKind,
	)

	require.NoError(t, err)
	require.NotNil(t, restClient)
}

func TestNewCRDClientForAPIVersionKind_DifferentGroupVersions(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {}))
	defer svr.Close()

	kubeConfigPath := createTestKubeConfig(t, svr.URL)

	tests := []struct {
		name       string
		apiVersion string
		kind       string
		resources  []*metav1.APIResourceList
		expectErr  bool
	}{
		{
			name:       "v1alpha1 version",
			apiVersion: "externaldns.k8s.io/v1alpha1",
			kind:       "DNSEndpoint",
			resources: []*metav1.APIResourceList{
				{
					GroupVersion: "externaldns.k8s.io/v1alpha1",
					APIResources: []metav1.APIResource{
						{Name: "dnsendpoints", Kind: "DNSEndpoint", Namespaced: true},
					},
				},
			},
			expectErr: false,
		},
		{
			name:       "custom group",
			apiVersion: "custom.example.com/v1",
			kind:       "CustomEndpoint",
			resources: []*metav1.APIResourceList{
				{
					GroupVersion: "custom.example.com/v1",
					APIResources: []metav1.APIResource{
						{Name: "customendpoints", Kind: "CustomEndpoint", Namespaced: true},
					},
				},
			},
			expectErr: false,
		},
		{
			name:       "wrong version requested",
			apiVersion: "externaldns.k8s.io/v1beta1",
			kind:       "DNSEndpoint",
			resources: []*metav1.APIResourceList{
				{
					GroupVersion: "externaldns.k8s.io/v1alpha1",
					APIResources: []metav1.APIResource{
						{Name: "dnsendpoints", Kind: "DNSEndpoint", Namespaced: true},
					},
				},
			},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fakeClient := fake.NewClientset()
			fakeClient.Resources = tt.resources

			restClient, err := NewCRDClientForAPIVersionKind(
				fakeClient,
				kubeConfigPath,
				"",
				tt.apiVersion,
				tt.kind,
			)

			if tt.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.NotNil(t, restClient)
			}
		})
	}
}
