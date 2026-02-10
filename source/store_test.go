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
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
	"sigs.k8s.io/external-dns/internal/testutils"
	"sigs.k8s.io/external-dns/source/types"
	gateway "sigs.k8s.io/gateway-api/pkg/client/clientset/versioned"
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

	// When a source is the only one requested and fails, "no sources" error is returned
	for _, source := range []string{
		types.Node, types.Service, types.Ingress, types.Pod,
		types.IstioGateway, types.IstioVirtualService,
		types.AmbassadorHost, types.GlooProxy, types.TraefikProxy, types.CRD, types.KongTCPIngress,
		types.F5VirtualServer, types.F5TransportServer,
	} {
		_, err := ByNames(context.TODO(), &Config{
			sources: []string{source},
		}, mockClientGenerator)
		suite.Error(err, source+" should return an error when it is the only source and fails")
		suite.Contains(err.Error(), "no sources could be initialized")
	}
}

func (suite *ByNamesTestSuite) TestIstioClientFails() {
	mockClientGenerator := new(MockClientGenerator)
	mockClientGenerator.On("KubeClient").Return(fakeKube.NewSimpleClientset(), nil)
	mockClientGenerator.On("IstioClient").Return(nil, errors.New("foo"))
	mockClientGenerator.On("DynamicKubernetesClient").Return(nil, errors.New("foo"))

	// Istio sources are optional — when requested alone and failing, "no sources" error
	for _, source := range []string{types.IstioGateway, types.IstioVirtualService} {
		_, err := ByNames(context.TODO(), &Config{
			sources: []string{source},
		}, mockClientGenerator)
		suite.Error(err, source+" should return an error when it is the only source and fails")
		suite.Contains(err.Error(), "no sources could be initialized")
	}
}

func (suite *ByNamesTestSuite) TestDynamicKubernetesClientFails() {
	mockClientGenerator := new(MockClientGenerator)
	mockClientGenerator.On("KubeClient").Return(fakeKube.NewClientset(), nil)
	mockClientGenerator.On("IstioClient").Return(istiofake.NewSimpleClientset(), nil)
	mockClientGenerator.On("DynamicKubernetesClient").Return(nil, errors.New("foo"))

	// All dynamic-client-dependent sources are optional — "no sources" when alone
	for _, source := range []string{
		types.AmbassadorHost, types.ContourHTTPProxy, types.GlooProxy, types.TraefikProxy,
		types.KongTCPIngress, types.F5VirtualServer, types.F5TransportServer,
	} {
		_, err := ByNames(context.TODO(), &Config{
			sources: []string{source},
		}, mockClientGenerator)
		suite.Error(err, source+" should return an error when it is the only source and fails")
		suite.Contains(err.Error(), "no sources could be initialized")
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

func TestEmptySourcesList(t *testing.T) {
	m := new(MockClientGenerator)
	_, err := ByNames(context.TODO(), &Config{
		sources: []string{},
	}, m)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no sources could be initialized")
}

func TestOptionalSourceSkipped(t *testing.T) {
	m := new(MockClientGenerator)
	m.On("KubeClient").Return(fakeKube.NewSimpleClientset(), nil)
	// GatewayClient will fail — gateway-grpcroute should be skipped
	m.On("GatewayClient").Return(nil, errors.New("gateway CRD not installed"))

	sources, err := ByNames(context.TODO(), &Config{
		sources: []string{types.Service, types.GatewayGrpcRoute},
	}, m)
	assert.NoError(t, err)
	assert.Len(t, sources, 1, "service should succeed, gateway-grpcroute should be skipped")
}

func TestAllOptionalSourcesSkipped(t *testing.T) {
	m := new(MockClientGenerator)
	m.On("GatewayClient").Return(nil, errors.New("gateway CRD not installed"))

	sources, err := ByNames(context.TODO(), &Config{
		sources: []string{types.GatewayHttpRoute, types.GatewayGrpcRoute},
	}, m)
	assert.Error(t, err)
	assert.Nil(t, sources)
	assert.Contains(t, err.Error(), "no sources could be initialized")
}

func TestSingleSourceFailsReturnsError(t *testing.T) {
	m := new(MockClientGenerator)
	m.On("KubeClient").Return(nil, errors.New("kube client broken"))

	_, err := ByNames(context.TODO(), &Config{
		sources: []string{types.Service},
	}, m)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no sources could be initialized")
}

func TestMixedSourcesPartialInit(t *testing.T) {
	m := new(MockClientGenerator)
	m.On("KubeClient").Return(fakeKube.NewSimpleClientset(), nil)
	m.On("IstioClient").Return(nil, errors.New("istio not installed"))

	sources, err := ByNames(context.TODO(), &Config{
		sources: []string{types.Service, types.IstioGateway},
	}, m)
	assert.NoError(t, err)
	assert.Len(t, sources, 1, "service should succeed, istio-gateway should be skipped")
}

func TestOptionalSourceSkippedLogsWarning(t *testing.T) {
	hook := testutils.LogsUnderTestWithLogLevel(log.WarnLevel, t)

	m := new(MockClientGenerator)
	m.On("KubeClient").Return(fakeKube.NewSimpleClientset(), nil)
	m.On("GatewayClient").Return(nil, errors.New("gateway CRD not installed"))

	_, err := ByNames(context.TODO(), &Config{
		sources: []string{types.Service, types.GatewayGrpcRoute},
	}, m)
	assert.NoError(t, err)
	testutils.TestHelperLogContainsWithLogLevel("Skipping source", log.WarnLevel, hook, t)
	testutils.TestHelperLogContains("gateway-grpcroute", hook, t)
}

func TestAllSourceTypesHandled(t *testing.T) {
	// Every known source type must be handled by BuildWithConfig (not return
	// ErrSourceNotFound). If a new type constant is added to the types package
	// without a corresponding case in BuildWithConfig, add it here.
	allSources := []string{
		types.Node, types.Service, types.Ingress, types.Pod,
		types.Fake, types.Connector,
		types.GatewayHttpRoute, types.GatewayGrpcRoute, types.GatewayTlsRoute,
		types.GatewayTcpRoute, types.GatewayUdpRoute,
		types.IstioGateway, types.IstioVirtualService,
		types.AmbassadorHost, types.ContourHTTPProxy,
		types.GlooProxy, types.TraefikProxy, types.OpenShiftRoute,
		types.CRD, types.SkipperRouteGroup, types.KongTCPIngress,
		types.F5VirtualServer, types.F5TransportServer,
	}

	p := &minimalMockClientGenerator{}
	for _, s := range allSources {
		_, err := BuildWithConfig(context.Background(), s, p, &Config{LabelFilter: labels.NewSelector()})
		assert.NotErrorIs(t, err, ErrSourceNotFound,
			"%q should be handled by BuildWithConfig", s)
	}
}

func TestUnknownSourceStillFails(t *testing.T) {
	m := new(MockClientGenerator)
	m.On("KubeClient").Return(fakeKube.NewSimpleClientset(), nil)

	_, err := ByNames(context.TODO(), &Config{
		sources: []string{types.Service, "nonexistent-source"},
	}, m)
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrSourceNotFound)
}
