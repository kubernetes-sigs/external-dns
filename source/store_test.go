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

	cfclient "github.com/cloudfoundry-community/go-cfclient"
	openshift "github.com/openshift/client-go/route/clientset/versioned"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	istioclient "istio.io/client-go/pkg/clientset/versioned"
	istiofake "istio.io/client-go/pkg/clientset/versioned/fake"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	fakeDynamic "k8s.io/client-go/dynamic/fake"
	"k8s.io/client-go/kubernetes"
	fakeKube "k8s.io/client-go/kubernetes/fake"
	gateway "sigs.k8s.io/gateway-api/pkg/client/clientset/gateway/versioned"
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
		}), nil)

	sources, err := ByNames(context.TODO(), mockClientGenerator, []string{"service", "ingress", "istio-gateway", "contour-httpproxy", "kong-tcpingress", "fake"}, minimalConfig)
	suite.NoError(err, "should not generate errors")
	suite.Len(sources, 6, "should generate all six sources")
}

func (suite *ByNamesTestSuite) TestOnlyFake() {
	mockClientGenerator := new(MockClientGenerator)
	mockClientGenerator.On("KubeClient").Return(fakeKube.NewSimpleClientset(), nil)

	sources, err := ByNames(context.TODO(), mockClientGenerator, []string{"fake"}, minimalConfig)
	suite.NoError(err, "should not generate errors")
	suite.Len(sources, 1, "should generate fake source")
	suite.Nil(mockClientGenerator.kubeClient, "client should not be created")
}

func (suite *ByNamesTestSuite) TestSourceNotFound() {
	mockClientGenerator := new(MockClientGenerator)
	mockClientGenerator.On("KubeClient").Return(fakeKube.NewSimpleClientset(), nil)

	sources, err := ByNames(context.TODO(), mockClientGenerator, []string{"foo"}, minimalConfig)
	suite.Equal(err, ErrSourceNotFound, "should return source not found")
	suite.Len(sources, 0, "should not returns any source")
}

func (suite *ByNamesTestSuite) TestKubeClientFails() {
	mockClientGenerator := new(MockClientGenerator)
	mockClientGenerator.On("KubeClient").Return(nil, errors.New("foo"))

	_, err := ByNames(context.TODO(), mockClientGenerator, []string{"service"}, minimalConfig)
	suite.Error(err, "should return an error if kubernetes client cannot be created")

	_, err = ByNames(context.TODO(), mockClientGenerator, []string{"ingress"}, minimalConfig)
	suite.Error(err, "should return an error if kubernetes client cannot be created")

	_, err = ByNames(context.TODO(), mockClientGenerator, []string{"istio-gateway"}, minimalConfig)
	suite.Error(err, "should return an error if kubernetes client cannot be created")

	_, err = ByNames(context.TODO(), mockClientGenerator, []string{"kong-tcpingress"}, minimalConfig)
	suite.Error(err, "should return an error if kubernetes client cannot be created")
}

func (suite *ByNamesTestSuite) TestIstioClientFails() {
	mockClientGenerator := new(MockClientGenerator)
	mockClientGenerator.On("KubeClient").Return(fakeKube.NewSimpleClientset(), nil)
	mockClientGenerator.On("IstioClient").Return(nil, errors.New("foo"))
	mockClientGenerator.On("DynamicKubernetesClient").Return(nil, errors.New("foo"))

	_, err := ByNames(context.TODO(), mockClientGenerator, []string{"istio-gateway"}, minimalConfig)
	suite.Error(err, "should return an error if istio client cannot be created")

	_, err = ByNames(context.TODO(), mockClientGenerator, []string{"contour-httpproxy"}, minimalConfig)
	suite.Error(err, "should return an error if contour client cannot be created")
}

func TestByNames(t *testing.T) {
	suite.Run(t, new(ByNamesTestSuite))
}

var minimalConfig = &Config{
	ContourLoadBalancerService: "heptio-contour/contour",
}
