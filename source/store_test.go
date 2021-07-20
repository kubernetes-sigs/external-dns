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
	"errors"
	"testing"

	cloudfoundry "github.com/cloudfoundry-community/go-cfclient"
	openshift "github.com/openshift/client-go/route/clientset/versioned"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	istio "istio.io/client-go/pkg/clientset/versioned"
	istiofake "istio.io/client-go/pkg/clientset/versioned/fake"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	fakeKube "k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/rest"
)

type MockClientGenerator struct {
	mock.Mock
}

func (m *MockClientGenerator) RESTConfig() (*rest.Config, error) {
	args := m.Called()
	cfg, _ := args.Get(0).(*rest.Config)
	return cfg, args.Error(1)
}

func (m *MockClientGenerator) KubeClient() (kubernetes.Interface, error) {
	args := m.Called()
	client, _ := args.Get(0).(kubernetes.Interface)
	return client, args.Error(1)
}

func (m *MockClientGenerator) IstioClient() (istio.Interface, error) {
	args := m.Called()
	client, _ := args.Get(0).(istio.Interface)
	return client, args.Error(1)
}

func (m *MockClientGenerator) CloudFoundryClient(cfAPIEndpoint string, cfUsername string, cfPassword string) (*cloudfoundry.Client, error) {
	args := m.Called()
	client, _ := args.Get(0).(*cloudfoundry.Client)
	return client, args.Error(1)
}

func (m *MockClientGenerator) DynamicKubernetesClient() (dynamic.Interface, error) {
	args := m.Called()
	client, _ := args.Get(0).(dynamic.Interface)
	return client, args.Error(1)
}

func (m *MockClientGenerator) OpenShiftClient() (openshift.Interface, error) {
	args := m.Called()
	client, _ := args.Get(0).(openshift.Interface)
	return client, args.Error(1)
}

type ByNamesTestSuite struct {
	suite.Suite
}

func (suite *ByNamesTestSuite) TestAllInitialized() {
	fakeDynamic, _ := newDynamicKubernetesClient()

	mockClientGenerator := new(MockClientGenerator)
	mockClientGenerator.On("KubeClient").Return(fakeKube.NewSimpleClientset(), nil)
	mockClientGenerator.On("IstioClient").Return(istiofake.NewSimpleClientset(), nil)
	mockClientGenerator.On("DynamicKubernetesClient").Return(fakeDynamic, nil)

	sources, err := ByNames(mockClientGenerator, []string{"service", "ingress", "istio-gateway", "contour-ingressroute", "contour-httpproxy", "kong-tcpingress", "fake"}, minimalConfig)
	suite.NoError(err, "should not generate errors")
	suite.Len(sources, 7, "should generate all six sources")
}

func (suite *ByNamesTestSuite) TestOnlyFake() {
	mockClientGenerator := new(MockClientGenerator)
	mockClientGenerator.On("KubeClient").Return(fakeKube.NewSimpleClientset(), nil)
	mockClientGenerator.On("KubeClient").RunFn = func(mock.Arguments) {
		suite.Fail("KubeClient should not be created")
	}

	sources, err := ByNames(mockClientGenerator, []string{"fake"}, minimalConfig)
	suite.NoError(err, "should not generate errors")
	suite.Len(sources, 1, "should generate fake source")
}

func (suite *ByNamesTestSuite) TestSourceNotFound() {
	mockClientGenerator := new(MockClientGenerator)
	mockClientGenerator.On("KubeClient").Return(fakeKube.NewSimpleClientset(), nil)

	sources, err := ByNames(mockClientGenerator, []string{"foo"}, minimalConfig)
	suite.Equal(err, ErrSourceNotFound, "should return source not found")
	suite.Len(sources, 0, "should not returns any source")
}

func (suite *ByNamesTestSuite) TestKubeClientFails() {
	mockClientGenerator := new(MockClientGenerator)
	mockClientGenerator.On("KubeClient").Return(nil, errors.New("foo"))

	_, err := ByNames(mockClientGenerator, []string{"service"}, minimalConfig)
	suite.Error(err, "should return an error if kubernetes client cannot be created")

	_, err = ByNames(mockClientGenerator, []string{"ingress"}, minimalConfig)
	suite.Error(err, "should return an error if kubernetes client cannot be created")

	_, err = ByNames(mockClientGenerator, []string{"istio-gateway"}, minimalConfig)
	suite.Error(err, "should return an error if kubernetes client cannot be created")

	_, err = ByNames(mockClientGenerator, []string{"contour-ingressroute"}, minimalConfig)
	suite.Error(err, "should return an error if kubernetes client cannot be created")

	_, err = ByNames(mockClientGenerator, []string{"kong-tcpingress"}, minimalConfig)
	suite.Error(err, "should return an error if kubernetes client cannot be created")
}

func (suite *ByNamesTestSuite) TestIstioClientFails() {
	mockClientGenerator := new(MockClientGenerator)
	mockClientGenerator.On("KubeClient").Return(fakeKube.NewSimpleClientset(), nil)
	mockClientGenerator.On("IstioClient").Return(nil, errors.New("foo"))
	mockClientGenerator.On("DynamicKubernetesClient").Return(nil, errors.New("foo"))

	_, err := ByNames(mockClientGenerator, []string{"istio-gateway"}, minimalConfig)
	suite.Error(err, "should return an error if istio client cannot be created")

	_, err = ByNames(mockClientGenerator, []string{"contour-ingressroute"}, minimalConfig)
	suite.Error(err, "should return an error if contour client cannot be created")
	_, err = ByNames(mockClientGenerator, []string{"contour-httpproxy"}, minimalConfig)
	suite.Error(err, "should return an error if contour client cannot be created")
}

func TestByNames(t *testing.T) {
	suite.Run(t, new(ByNamesTestSuite))
}

var minimalConfig = &Config{
	ContourLoadBalancerService: "heptio-contour/contour",
}
