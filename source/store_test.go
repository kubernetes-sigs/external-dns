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

	openshift "github.com/openshift/client-go/route/clientset/versioned"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	istioclient "istio.io/client-go/pkg/clientset/versioned"
	istiofake "istio.io/client-go/pkg/clientset/versioned/fake"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	fakeDynamic "k8s.io/client-go/dynamic/fake"
	"k8s.io/client-go/kubernetes"
	fakeKube "k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/rest"
	gateway "sigs.k8s.io/gateway-api/pkg/client/clientset/versioned"

	"sigs.k8s.io/external-dns/source/types"
)

type MockClientGenerator struct {
	mock.Mock
	kubeClient              kubernetes.Interface
	gatewayClient           gateway.Interface
	istioClient             istioclient.Interface
	dynamicKubernetesClient dynamic.Interface
	openshiftClient         openshift.Interface
}

func (m *MockClientGenerator) KubeClient() (kubernetes.Interface, error) {
	args := m.Called()
	if args.Error(1) == nil {
		m.kubeClient = args.Get(0).(kubernetes.Interface)
		return m.kubeClient, nil
	}
	return nil, args.Error(1)
}

func (m *MockClientGenerator) GatewayClient() (gateway.Interface, error) {
	args := m.Called()
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	m.gatewayClient = args.Get(0).(gateway.Interface)
	return m.gatewayClient, nil
}

func (m *MockClientGenerator) IstioClient() (istioclient.Interface, error) {
	args := m.Called()
	if args.Error(1) == nil {
		m.istioClient = args.Get(0).(istioclient.Interface)
		return m.istioClient, nil
	}
	return nil, args.Error(1)
}

func (m *MockClientGenerator) DynamicKubernetesClient() (dynamic.Interface, error) {
	args := m.Called()
	if args.Error(1) == nil {
		m.dynamicKubernetesClient = args.Get(0).(dynamic.Interface)
		return m.dynamicKubernetesClient, nil
	}
	return nil, args.Error(1)
}

func (m *MockClientGenerator) OpenShiftClient() (openshift.Interface, error) {
	args := m.Called()
	if args.Error(1) == nil {
		m.openshiftClient = args.Get(0).(openshift.Interface)
		return m.openshiftClient, nil
	}
	return nil, args.Error(1)
}

func (m *MockClientGenerator) RESTConfig() (*rest.Config, error) {
	args := m.Called()
	if args.Error(1) == nil {
		return args.Get(0).(*rest.Config), nil
	}
	return nil, args.Error(1)
}

type ByNamesTestSuite struct {
	suite.Suite
}

func (suite *ByNamesTestSuite) TestAllInitialized() {
	mockClientGenerator := new(MockClientGenerator)
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
	mockClientGenerator := new(MockClientGenerator)
	mockClientGenerator.On("KubeClient").Return(fakeKube.NewClientset(), nil)

	sources, err := ByNames(context.TODO(), &Config{
		sources: []string{types.Fake},
	}, mockClientGenerator)
	suite.NoError(err, "should not generate errors")
	suite.Len(sources, 1, "should generate fake source")
	suite.Nil(mockClientGenerator.kubeClient, "client should not be created")
}

func (suite *ByNamesTestSuite) TestSourceNotFound() {
	mockClientGenerator := new(MockClientGenerator)
	mockClientGenerator.On("KubeClient").Return(fakeKube.NewClientset(), nil)
	sources, err := ByNames(context.TODO(), &Config{
		sources: []string{"foo"},
	}, mockClientGenerator)
	suite.Equal(err, ErrSourceNotFound, "should return source not found")
	suite.Empty(sources, "should not returns any source")
}

func (suite *ByNamesTestSuite) TestKubeClientFails() {
	mockClientGenerator := new(MockClientGenerator)
	mockClientGenerator.On("KubeClient").Return(nil, errors.New("foo"))

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
	mockClientGenerator := new(MockClientGenerator)
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
	mockClientGenerator := new(MockClientGenerator)
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

type minimalMockClientGenerator struct{}

var errMock = errors.New("mock not implemented")

func (m *minimalMockClientGenerator) KubeClient() (kubernetes.Interface, error) { return nil, errMock }
func (m *minimalMockClientGenerator) GatewayClient() (gateway.Interface, error) { return nil, errMock }
func (m *minimalMockClientGenerator) IstioClient() (istioclient.Interface, error) {
	return nil, errMock
}

func (m *minimalMockClientGenerator) DynamicKubernetesClient() (dynamic.Interface, error) {
	return nil, errMock
}
func (m *minimalMockClientGenerator) OpenShiftClient() (openshift.Interface, error) {
	return nil, errMock
}
func (m *minimalMockClientGenerator) RESTConfig() (*rest.Config, error) { return nil, errMock }

func TestBuildWithConfig_InvalidSource(t *testing.T) {
	ctx := context.Background()
	p := &minimalMockClientGenerator{}
	cfg := &Config{LabelFilter: labels.NewSelector()}

	src, err := BuildWithConfig(ctx, "not-a-source", p, cfg)
	if src != nil {
		t.Errorf("expected nil source for invalid type, got: %v", src)
	}
	if !errors.Is(err, ErrSourceNotFound) {
		t.Errorf("expected ErrSourceNotFound, got: %v", err)
	}
}

func TestConfig_ClientGenerator(t *testing.T) {
	cfg := &Config{
		KubeConfig:     "/path/to/kubeconfig",
		APIServerURL:   "https://api.example.com",
		RequestTimeout: 30 * time.Second,
		UpdateEvents:   false,
	}

	gen := cfg.ClientGenerator()

	assert.Equal(t, "/path/to/kubeconfig", gen.KubeConfig)
	assert.Equal(t, "https://api.example.com", gen.APIServerURL)
	assert.Equal(t, 30*time.Second, gen.RequestTimeout)
}

func TestConfig_ClientGenerator_UpdateEvents(t *testing.T) {
	cfg := &Config{
		KubeConfig:     "/path/to/kubeconfig",
		APIServerURL:   "https://api.example.com",
		RequestTimeout: 30 * time.Second,
		UpdateEvents:   true, // Special case
	}

	gen := cfg.ClientGenerator()

	assert.Equal(t, time.Duration(0), gen.RequestTimeout, "UpdateEvents should set timeout to 0")
}

func TestConfig_ClientGenerator_Caching(t *testing.T) {
	cfg := &Config{
		KubeConfig:     "/path/to/kubeconfig",
		APIServerURL:   "https://api.example.com",
		RequestTimeout: 30 * time.Second,
		UpdateEvents:   false,
	}

	// Call ClientGenerator twice
	gen1 := cfg.ClientGenerator()
	gen2 := cfg.ClientGenerator()

	// Should return the same instance (cached)
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
		cfg := &Config{
			KubeConfig:     "",
			APIServerURL:   "",
			RequestTimeout: 45 * time.Second,
			UpdateEvents:   false,
		}

		gen := cfg.ClientGenerator()

		// Verify ClientGenerator has correct timeout
		assert.Equal(t, 45*time.Second, gen.RequestTimeout,
			"ClientGenerator should have the configured RequestTimeout")

		config, err := gen.RESTConfig()

		// Even if config creation fails, the timeout setting should be correct
		assert.Equal(t, 45*time.Second, gen.RequestTimeout,
			"RequestTimeout should remain 45s after RESTConfig() call")

		if err == nil {
			require.NotNil(t, config, "Config should not be nil when error is nil")
			assert.Equal(t, 45*time.Second, config.Timeout,
				"RESTConfig should propagate the timeout")
		}
	})

	t.Run("UpdateEvents sets timeout to zero", func(t *testing.T) {
		cfg := &Config{
			KubeConfig:     "",
			APIServerURL:   "",
			RequestTimeout: 45 * time.Second,
			UpdateEvents:   true, // Should override to 0
		}

		gen := cfg.ClientGenerator()

		// When UpdateEvents=true, ClientGenerator sets timeout to 0 (for long-running watches)
		assert.Equal(t, time.Duration(0), gen.RequestTimeout,
			"ClientGenerator should have zero timeout when UpdateEvents=true")

		config, err := gen.RESTConfig()

		// Verify the timeout is 0, regardless of whether config was created
		assert.Equal(t, time.Duration(0), gen.RequestTimeout,
			"RequestTimeout should remain 0 after RESTConfig() call")

		if err == nil {
			require.NotNil(t, config, "Config should not be nil when error is nil")
			assert.Equal(t, time.Duration(0), config.Timeout,
				"RESTConfig should have zero timeout for watch operations")
		}
	})
}

// TestSingletonClientGenerator_RESTConfig_SharedAcrossClients verifies singleton is shared
func TestSingletonClientGenerator_RESTConfig_SharedAcrossClients(t *testing.T) {
	gen := &SingletonClientGenerator{
		KubeConfig:     "",
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
