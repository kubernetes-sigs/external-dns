/*
Copyright 2017 The Kubernetes Authors.

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
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	istiofake "istio.io/client-go/pkg/clientset/versioned/fake"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	fakeDynamic "k8s.io/client-go/dynamic/fake"
	fakeKube "k8s.io/client-go/kubernetes/fake"

	"sigs.k8s.io/external-dns/internal/testutils"
	externaldns "sigs.k8s.io/external-dns/pkg/apis/externaldns"
	"sigs.k8s.io/external-dns/source/types"
)

type ByNamesTestSuite struct {
	suite.Suite
}

func (suite *ByNamesTestSuite) TestAllInitialized() {
	mockClientGenerator := new(testutils.MockClientGenerator)
	mockClientGenerator.On("KubeClient").Return(fakeKube.NewSimpleClientset(), nil)
	mockClientGenerator.On("IstioClient").Return(istiofake.NewSimpleClientset(), nil)
	mockClientGenerator.On("DynamicKubernetesClient").Return(fakeDynamic.NewSimpleDynamicClientWithCustomListKinds(runtime.NewScheme(),
		map[schema.GroupVersionResource]string{
			{
				Group:    "projectcontour.io",
				Version:  "v1",
				Resource: "httpproxies",
			}: "HTTPPRoxiesList",
			{
				Group:    "contour.heptio.com",
				Version:  "v1beta1",
				Resource: "tcpingresses",
			}: "TCPIngressesList",
			{
				Group:    "configuration.konghq.com",
				Version:  "v1beta1",
				Resource: "tcpingresses",
			}: "TCPIngressesList",
			{
				Group:    "cis.f5.com",
				Version:  "v1",
				Resource: "virtualservers",
			}: "VirtualServersList",
			{
				Group:    "cis.f5.com",
				Version:  "v1",
				Resource: "transportservers",
			}: "TransportServersList",
			{
				Group:    "traefik.containo.us",
				Version:  "v1alpha1",
				Resource: "ingressroutes",
			}: "IngressRouteList",
			{
				Group:    "traefik.containo.us",
				Version:  "v1alpha1",
				Resource: "ingressroutetcps",
			}: "IngressRouteTCPList",
			{
				Group:    "traefik.containo.us",
				Version:  "v1alpha1",
				Resource: "ingressrouteudps",
			}: "IngressRouteUDPList",
			{
				Group:    "traefik.io",
				Version:  "v1alpha1",
				Resource: "ingressroutes",
			}: "IngressRouteList",
			{
				Group:    "traefik.io",
				Version:  "v1alpha1",
				Resource: "ingressroutetcps",
			}: "IngressRouteTCPList",
			{
				Group:    "traefik.io",
				Version:  "v1alpha1",
				Resource: "ingressrouteudps",
			}: "IngressRouteUDPList",
		}), nil)

	ss := []string{
		types.Service, types.Ingress, types.IstioGateway, types.ContourHTTPProxy,
		types.KongTCPIngress, types.F5VirtualServer, types.F5TransportServer, types.TraefikProxy, types.Fake,
	}
	sources, err := ByNames(context.TODO(), &Config{
		sources: ss,
	}, mockClientGenerator)
	suite.NoError(err, "should not generate errors")
	suite.Len(sources, 9, "should generate all nine sources")
}

func (suite *ByNamesTestSuite) TestOnlyFake() {
	mockClientGenerator := new(testutils.MockClientGenerator)
	mockClientGenerator.On("KubeClient").Return(fakeKube.NewClientset(), nil)

	sources, err := ByNames(context.TODO(), &Config{
		sources: []string{types.Fake},
	}, mockClientGenerator)
	suite.NoError(err, "should not generate errors")
	suite.Len(sources, 1, "should generate fake source")
	suite.Nil(mockClientGenerator.KubeClientValue, "client should not be created")
}

func (suite *ByNamesTestSuite) TestSourceNotFound() {
	mockClientGenerator := new(testutils.MockClientGenerator)
	mockClientGenerator.On("KubeClient").Return(fakeKube.NewClientset(), nil)
	sources, err := ByNames(context.TODO(), &Config{
		sources: []string{"foo"},
	}, mockClientGenerator)
	suite.Equal(err, ErrSourceNotFound, "should return source not found")
	suite.Empty(sources, "should not returns any source")
}

func (suite *ByNamesTestSuite) TestKubeClientFails() {
	mockClientGenerator := new(testutils.MockClientGenerator)
	mockClientGenerator.On("KubeClient").Return(nil, errors.New("foo"))
	mockClientGenerator.On("RESTConfig").Return(nil, errors.New("foo"))

	sourceUnderTest := []string{
		types.Node, types.Service, types.Ingress, types.Pod, types.IstioGateway, types.IstioVirtualService,
		types.AmbassadorHost, types.GlooProxy, types.TraefikProxy, types.CRD, types.KongTCPIngress,
		types.F5VirtualServer, types.F5TransportServer,
	}

	for _, source := range sourceUnderTest {
		_, err := ByNames(context.TODO(), &Config{
			sources: []string{source},
		}, mockClientGenerator)
		suite.Error(err, source+" should return an error if kubernetes client cannot be created")
	}
}

func (suite *ByNamesTestSuite) TestIstioClientFails() {
	mockClientGenerator := new(testutils.MockClientGenerator)
	mockClientGenerator.On("KubeClient").Return(fakeKube.NewSimpleClientset(), nil)
	mockClientGenerator.On("IstioClient").Return(nil, errors.New("foo"))
	mockClientGenerator.On("DynamicKubernetesClient").Return(nil, errors.New("foo"))

	sourcesDependentOnIstioClient := []string{types.IstioGateway, types.IstioVirtualService}

	for _, source := range sourcesDependentOnIstioClient {
		_, err := ByNames(context.TODO(), &Config{
			sources: []string{source},
		}, mockClientGenerator)
		suite.Error(err, source+" should return an error if istio client cannot be created")
	}
}

func (suite *ByNamesTestSuite) TestDynamicKubernetesClientFails() {
	mockClientGenerator := new(testutils.MockClientGenerator)
	mockClientGenerator.On("KubeClient").Return(fakeKube.NewClientset(), nil)
	mockClientGenerator.On("IstioClient").Return(istiofake.NewSimpleClientset(), nil)
	mockClientGenerator.On("DynamicKubernetesClient").Return(nil, errors.New("foo"))

	sourcesDependentOnDynamicKubernetesClient := []string{
		types.AmbassadorHost, types.ContourHTTPProxy, types.GlooProxy, types.TraefikProxy,
		types.KongTCPIngress, types.F5VirtualServer, types.F5TransportServer,
	}

	for _, source := range sourcesDependentOnDynamicKubernetesClient {
		_, err := ByNames(context.TODO(), &Config{
			sources: []string{source},
		}, mockClientGenerator)
		suite.Error(err, source+" should return an error if dynamic kubernetes client cannot be created")
	}
}

func TestByNames(t *testing.T) {
	suite.Run(t, new(ByNamesTestSuite))
}

func TestBuildWithConfig_InvalidSource(t *testing.T) {
	ctx := t.Context()
	p := testutils.StubClientGenerator{}
	cfg := &Config{LabelFilter: labels.NewSelector()}

	src, err := BuildWithConfig(ctx, "not-a-source", p, cfg)
	if src != nil {
		t.Errorf("expected nil source for invalid type, got: %v", src)
	}
	if !errors.Is(err, ErrSourceNotFound) {
		t.Errorf("expected ErrSourceNotFound, got: %v", err)
	}
}

func TestConfig_ClientGenerator_Caching(t *testing.T) {
	cfg := &Config{
		KubeConfig:            "/path/to/kubeconfig",
		APIServerURL:          "https://api.example.com",
		KubeAPIRequestTimeout: 30 * time.Second,
	}

	gen1 := cfg.ClientGenerator()
	gen2 := cfg.ClientGenerator()

	assert.Same(t, gen1, gen2, "ClientGenerator should return the same cached instance")
}

// TestSingletonClientGenerator_RESTConfig_TimeoutPropagation verifies timeout configuration
func TestSingletonClientGenerator_RESTConfig_TimeoutPropagation(t *testing.T) {
	testCases := []struct {
		name           string
		requestTimeout time.Duration
	}{
		{
			name:           "30 second timeout",
			requestTimeout: 30 * time.Second,
		},
		{
			name:           "60 second timeout",
			requestTimeout: 60 * time.Second,
		},
		{
			name:           "zero timeout (for watches)",
			requestTimeout: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gen := &SingletonClientGenerator{
				KubeConfig:     "",
				APIServerURL:   "",
				RequestTimeout: tc.requestTimeout,
			}

			// Verify the generator was configured with correct timeout
			assert.Equal(t, tc.requestTimeout, gen.RequestTimeout,
				"SingletonClientGenerator should have the configured RequestTimeout")

			config, err := gen.RESTConfig()

			// Even if config creation failed, verify the timeout was set in generator
			assert.Equal(t, tc.requestTimeout, gen.RequestTimeout,
				"RequestTimeout should remain unchanged after RESTConfig() call")

			// If config was successfully created, verify timeout propagated correctly
			if err == nil {
				require.NotNil(t, config, "Config should not be nil when error is nil")
				assert.Equal(t, tc.requestTimeout, config.Timeout,
					"REST config should have timeout matching RequestTimeout field")
			}
		})
	}
}

// TestConfig_ClientGenerator_RESTConfig_Integration verifies Config → ClientGenerator → RESTConfig flow
func TestConfig_ClientGenerator_RESTConfig_Integration(t *testing.T) {
	t.Run("normal timeout is propagated", func(t *testing.T) {
		cfg := &Config{KubeAPIRequestTimeout: 45 * time.Second}

		config, err := cfg.ClientGenerator().RESTConfig()
		if err == nil {
			require.NotNil(t, config)
			assert.Equal(t, 45*time.Second, config.Timeout, "RESTConfig should propagate the timeout")
		}
	})

	t.Run("UpdateEvents sets timeout to zero", func(t *testing.T) {
		cfg := &Config{KubeAPIRequestTimeout: 45 * time.Second, UpdateEvents: true}

		config, err := cfg.ClientGenerator().RESTConfig()
		if err == nil {
			require.NotNil(t, config)
			assert.Equal(t, time.Duration(0), config.Timeout, "RESTConfig should have zero timeout for watch operations")
		}
	})
}

// TestSingletonClientGenerator_RESTConfig_SharedAcrossClients verifies singleton is shared
func TestSingletonClientGenerator_RESTConfig_SharedAcrossClients(t *testing.T) {
	gen := &SingletonClientGenerator{
		KubeConfig:     "/nonexistent/path/to/kubeconfig",
		APIServerURL:   "",
		RequestTimeout: 30 * time.Second,
	}

	// Get REST config multiple times
	restConfig1, err1 := gen.RESTConfig()
	restConfig2, err2 := gen.RESTConfig()
	restConfig3, err3 := gen.RESTConfig()

	// Verify singleton behavior - all should return same instance
	assert.Same(t, restConfig1, restConfig2, "RESTConfig should return same instance on second call")
	assert.Same(t, restConfig1, restConfig3, "RESTConfig should return same instance on third call")

	// Verify the internal field matches
	assert.Same(t, restConfig1, gen.restConfig,
		"Internal restConfig field should match returned value")

	// Verify first call had error (no valid kubeconfig)
	assert.Error(t, err1, "First call should return error when kubeconfig is invalid")

	// Due to sync.Once bug, subsequent calls won't return the error
	// This is documented in the TODO comment on SingletonClientGenerator
	require.NoError(t, err2, "Second call does not return error due to sync.Once bug")
	require.NoError(t, err3, "Third call does not return error due to sync.Once bug")
}

func TestNewSourceConfig(t *testing.T) {
	tests := []struct {
		name           string
		cfg            *externaldns.Config
		wantConfigured bool
		wantCombining  bool
		wantErr        bool
	}{
		{
			name: "no templates configured",
			cfg:  &externaldns.Config{},
		},
		{
			name: "fqdn template only",
			cfg: &externaldns.Config{
				FQDNTemplate: "{{.Name}}.example.com",
			},
			wantConfigured: true,
		},
		{
			name: "fqdn template with combine",
			cfg: &externaldns.Config{
				FQDNTemplate:             "{{.Name}}.example.com",
				CombineFQDNAndAnnotation: true,
			},
			wantConfigured: true,
			wantCombining:  true,
		},
		{
			name: "all three templates configured",
			cfg: &externaldns.Config{
				FQDNTemplate:             "{{.Name}}.example.com",
				TargetTemplate:           "{{.Name}}.targets.example.com",
				FQDNTargetTemplate:       "{{.Name}}.example.com:{{.Name}}.targets.example.com",
				CombineFQDNAndAnnotation: true,
			},
			wantConfigured: true,
			wantCombining:  true,
		},
		{
			name:    "invalid fqdn template",
			cfg:     &externaldns.Config{FQDNTemplate: "{{.Name"},
			wantErr: true,
		},
		{
			name:    "invalid target template",
			cfg:     &externaldns.Config{TargetTemplate: "{{.Status.LoadBalancer.Ingress"},
			wantErr: true,
		},
		{
			name:    "invalid fqdn-target template",
			cfg:     &externaldns.Config{FQDNTargetTemplate: "{{.Name}}.example.com:{{.Status"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSourceConfig(tt.cfg)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			tmpl := got.TemplateEngine
			assert.Equal(t, tt.wantConfigured, tmpl.IsConfigured(), "IsConfigured")
			assert.Equal(t, tt.wantCombining, tmpl.Combining(), "Combining")
		})
	}
}

func TestKubeAPIRateLimitPropagation(t *testing.T) {
	t.Run("NewSourceConfig propagates QPS and burst", func(t *testing.T) {
		cfg := &externaldns.Config{
			KubeAPIQPS:   20,
			KubeAPIBurst: 40,
		}
		got, err := NewSourceConfig(cfg)
		require.NoError(t, err)
		assert.Equal(t, 20, got.KubeAPIQPS)
		assert.Equal(t, 40, got.KubeAPIBurst)
	})

	t.Run("ClientGenerator wires QPS and burst into SingletonClientGenerator", func(t *testing.T) {
		cfg := &Config{
			KubeAPIQPS:   15,
			KubeAPIBurst: 30,
		}
		scg, ok := cfg.ClientGenerator().(*SingletonClientGenerator)
		require.True(t, ok)
		assert.Equal(t, 15, scg.QPS)
		assert.Equal(t, 30, scg.Burst)
	})
}
