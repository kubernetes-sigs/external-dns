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
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/client-go/tools/clientcmd"
)

func TestGetRestConfig_WithKubeConfig(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {}))
	defer svr.Close()

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
	err = os.WriteFile(mockKubeCfgPath, fmt.Appendf(nil, kubeCfgTemplate, svr.URL), 0644)
	require.NoError(t, err)

	config, err := GetRestConfig(mockKubeCfgPath, "")
	require.NoError(t, err)
	require.NotNil(t, config)
	assert.Equal(t, svr.URL, config.Host)
}

func TestInstrumentedRESTConfig_AddsMetrics(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {}))
	defer svr.Close()

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
	err = os.WriteFile(mockKubeCfgPath, fmt.Appendf(nil, kubeCfgTemplate, svr.URL), 0644)
	require.NoError(t, err)

	timeout := 30 * time.Second
	config, err := InstrumentedRESTConfig(mockKubeCfgPath, "", timeout)
	require.NoError(t, err)
	require.NotNil(t, config)

	assert.Equal(t, timeout, config.Timeout)
	assert.NotNil(t, config.WrapTransport, "WrapTransport should be set for metrics")
}

func TestGetRestConfig_RecommendedHomeFile(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {}))
	defer svr.Close()

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
`
	err = os.WriteFile(mockKubeCfgPath, fmt.Appendf(nil, kubeCfgTemplate, svr.URL), 0644)
	require.NoError(t, err)

	prevRecommendedHomeFile := clientcmd.RecommendedHomeFile
	t.Cleanup(func() {
		clientcmd.RecommendedHomeFile = prevRecommendedHomeFile
	})
	clientcmd.RecommendedHomeFile = mockKubeCfgPath

	config, err := GetRestConfig("", "")
	require.NoError(t, err)
	require.NotNil(t, config)
	assert.Equal(t, svr.URL, config.Host)
}
