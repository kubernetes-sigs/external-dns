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

	"github.com/cloudfoundry-community/go-cfclient"
	openshift "github.com/openshift/client-go/route/clientset/versioned"
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
	"sigs.k8s.io/external-dns/source/types"
	gateway "sigs.k8s.io/gateway-api/pkg/client/clientset/versioned"
)

type MockClientGenerator struct {
	mock.Mock
	kubeClient              kubernetes.Interface
	gatewayClient           gateway.Interface
	istioClient             istioclient.Interface
	cloudFoundryClient      *cfclient.Client
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

func (m *MockClientGenerator) CloudFoundryClient(cfAPIEndpoint string, cfUsername string, cfPassword string) (*cfclient.Client, error) {
	args := m.Called()
	if args.Error(1) == nil {
		m.cloudFoundryClient = args.Get(0).(*cfclient.Client)
		return m.cloudFoundryClient, nil
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

	sources, err := ByNames(context.TODO(), mockClientGenerator, []string{
		types.Service, types.Ingress, types.IstioGateway, types.ContourHTTPProxy,
		types.KongTCPIngress, types.F5VirtualServer, types.F5TransportServer, types.TraefikProxy, types.Fake,
	}, &Config{})
	suite.NoError(err, "should not generate errors")
	suite.Len(sources, 9, "should generate all nine sources")
}

func (suite *ByNamesTestSuite) TestOnlyFake() {
	mockClientGenerator := new(MockClientGenerator)
	mockClientGenerator.On("KubeClient").Return(fakeKube.NewSimpleClientset(), nil)

	sources, err := ByNames(context.TODO(), mockClientGenerator, []string{types.Fake}, &Config{})
	suite.NoError(err, "should not generate errors")
	suite.Len(sources, 1, "should generate fake source")
	suite.Nil(mockClientGenerator.kubeClient, "client should not be created")
}

func (suite *ByNamesTestSuite) TestSourceNotFound() {
	mockClientGenerator := new(MockClientGenerator)
	mockClientGenerator.On("KubeClient").Return(fakeKube.NewSimpleClientset(), nil)

	sources, err := ByNames(context.TODO(), mockClientGenerator, []string{"foo"}, &Config{})
	suite.Equal(err, ErrSourceNotFound, "should return source not found")
	suite.Empty(sources, "should not returns any source")
}

func (suite *ByNamesTestSuite) TestKubeClientFails() {
	mockClientGenerator := new(MockClientGenerator)
	mockClientGenerator.On("KubeClient").Return(nil, errors.New("foo"))

	sourcesDependentOnKubeClient := []string{
		types.Node, types.Service, types.Ingress, types.Pod, types.IstioGateway, types.IstioVirtualService,
		types.AmbassadorHost, types.GlooProxy, types.TraefikProxy, types.CRD, types.KongTCPIngress,
		types.F5VirtualServer, types.F5TransportServer,
	}

	for _, source := range sourcesDependentOnKubeClient {
		_, err := ByNames(context.TODO(), mockClientGenerator, []string{source}, &Config{})
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
		_, err := ByNames(context.TODO(), mockClientGenerator, []string{source}, &Config{})
		suite.Error(err, source+" should return an error if istio client cannot be created")
	}
}

func (suite *ByNamesTestSuite) TestDynamicKubernetesClientFails() {
	mockClientGenerator := new(MockClientGenerator)
	mockClientGenerator.On("KubeClient").Return(fakeKube.NewSimpleClientset(), nil)
	mockClientGenerator.On("IstioClient").Return(istiofake.NewSimpleClientset(), nil)
	mockClientGenerator.On("DynamicKubernetesClient").Return(nil, errors.New("foo"))

	sourcesDependentOnDynamicKubernetesClient := []string{
		types.AmbassadorHost, types.ContourHTTPProxy, types.GlooProxy, types.TraefikProxy,
		types.KongTCPIngress, types.F5VirtualServer, types.F5TransportServer,
	}

	for _, source := range sourcesDependentOnDynamicKubernetesClient {
		_, err := ByNames(context.TODO(), mockClientGenerator, []string{source}, &Config{})
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
func (m *minimalMockClientGenerator) CloudFoundryClient(string, string, string) (*cfclient.Client, error) {
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
