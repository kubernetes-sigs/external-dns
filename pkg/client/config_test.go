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
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/flowcontrol"
)

func TestGetRestConfig_WithKubeConfig(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {}))
	defer svr.Close()

	kubeCfgPath := writeKubeConfig(t, svr.URL)

	config, err := buildRestConfig(kubeCfgPath, "")
	require.NoError(t, err)
	require.NotNil(t, config)
	assert.Equal(t, svr.URL, config.Host)
}

func TestNewKubeClient(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {}))
	defer svr.Close()

	config, err := InstrumentedRESTConfig(writeKubeConfig(t, svr.URL), "", 30*time.Second, 0, 0)
	require.NoError(t, err)

	client, err := NewKubeClient(config)
	require.NoError(t, err)
	assert.NotNil(t, client)
}

func TestInstrumentedRESTConfig_AddsMetrics(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {}))
	defer svr.Close()

	timeout := 30 * time.Second
	config, err := InstrumentedRESTConfig(writeKubeConfig(t, svr.URL), "", timeout, 0, 0)
	require.NoError(t, err)
	require.NotNil(t, config)

	assert.Equal(t, timeout, config.Timeout)
	assert.NotNil(t, config.WrapTransport, "WrapTransport should be set for metrics")
	assert.NotNil(t, config.RateLimiter, "RateLimiter should always be set")
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

	config, err := buildRestConfig("", "")
	require.NoError(t, err)
	require.NotNil(t, config)
	assert.Equal(t, svr.URL, config.Host)
}

func TestInstrumentedRESTConfig_QPSAndBurstApplied(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {}))
	defer svr.Close()

	kubeCfgPath := writeKubeConfig(t, svr.URL)

	config, err := InstrumentedRESTConfig(kubeCfgPath, "", 30*time.Second, 20, 40)
	require.NoError(t, err)
	require.NotNil(t, config)

	assert.Equal(t, 20, int(config.QPS))
	assert.Equal(t, 40, config.Burst)
	assert.NotNil(t, config.RateLimiter)
	assert.Equal(t, 20, int(config.RateLimiter.QPS()))
}

func TestInstrumentedRESTConfig_ZeroQPSKeepsConfigDefaults(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {}))
	defer svr.Close()

	kubeCfgPath := writeKubeConfig(t, svr.URL)

	config, err := InstrumentedRESTConfig(kubeCfgPath, "", 30*time.Second, 0, 0)
	require.NoError(t, err)
	require.NotNil(t, config)

	// qps == 0: client-go defaults applied; rate limiter still installed
	assert.Equal(t, int(rest.DefaultQPS), int(config.QPS))
	assert.Equal(t, rest.DefaultBurst, config.Burst)
	assert.NotNil(t, config.RateLimiter)
	assert.Equal(t, int(rest.DefaultQPS), int(config.RateLimiter.QPS()))
}

func TestEnrichingRateLimiter(t *testing.T) {
	t.Run("delegates QPS and TryAccept", func(t *testing.T) {
		rl := &rateLimiter{delegate: flowcontrol.NewTokenBucketRateLimiter(5, 10)}
		assert.Equal(t, 5, int(rl.QPS()))
		assert.True(t, rl.TryAccept(), "first TryAccept should succeed with non-zero burst")
	})

	t.Run("Accept does not panic", func(t *testing.T) {
		rl := &rateLimiter{delegate: flowcontrol.NewTokenBucketRateLimiter(1000, 10)}
		assert.NotPanics(t, rl.Accept)
	})

	t.Run("Stop does not panic", func(t *testing.T) {
		rl := &rateLimiter{delegate: flowcontrol.NewTokenBucketRateLimiter(5, 10)}
		assert.NotPanics(t, rl.Stop)
	})

	t.Run("Wait enriches error with actionable hint", func(t *testing.T) {
		// burst=0: no tokens available; cancelled context triggers error immediately
		rl := &rateLimiter{delegate: flowcontrol.NewTokenBucketRateLimiter(0.0001, 0)}
		ctx, cancel := context.WithCancel(t.Context())
		cancel()
		err := rl.Wait(ctx)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "consider raising --kube-api-qps/--kube-api-burst")
	})

	t.Run("Wait returns nil when token is available", func(t *testing.T) {
		// burst=1: one token available immediately
		rl := &rateLimiter{delegate: flowcontrol.NewTokenBucketRateLimiter(100, 1)}
		assert.NoError(t, rl.Wait(t.Context()))
	})
}

// writeKubeConfig writes a minimal kubeconfig pointing at serverURL into a temp dir
// and returns the path.
func writeKubeConfig(t *testing.T, serverURL string) string {
	t.Helper()
	dir := filepath.Join(t.TempDir(), ".kube")
	path := filepath.Join(dir, "config")
	require.NoError(t, os.MkdirAll(dir, 0755))
	tmpl := `apiVersion: v1
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
	require.NoError(t, os.WriteFile(path, fmt.Appendf(nil, tmpl, serverURL), 0644))
	return path
}
