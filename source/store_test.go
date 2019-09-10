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

	cfclient "github.com/cloudfoundry-community/go-cfclient"
	contour "github.com/heptio/contour/apis/generated/clientset/versioned"
	fakeContour "github.com/heptio/contour/apis/generated/clientset/versioned/fake"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	istiomodel "istio.io/istio/pilot/pkg/model"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/fake"
)

type MockClientGenerator struct {
	mock.Mock
	kubeClient         kubernetes.Interface
	istioClient        istiomodel.ConfigStore
	cloudFoundryClient *cfclient.Client
	contourClient      contour.Interface
}

func (m *MockClientGenerator) KubeClient() (kubernetes.Interface, error) {
	args := m.Called()
	if args.Error(1) == nil {
		m.kubeClient = args.Get(0).(kubernetes.Interface)
		return m.kubeClient, nil
	}
	return nil, args.Error(1)
}

func (m *MockClientGenerator) IstioClient() (istiomodel.ConfigStore, error) {
	args := m.Called()
	if args.Error(1) == nil {
		m.istioClient = args.Get(0).(istiomodel.ConfigStore)
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

func (m *MockClientGenerator) ContourClient() (contour.Interface, error) {
	args := m.Called()
	if args.Error(1) == nil {
		m.contourClient = args.Get(0).(contour.Interface)
		return m.contourClient, nil
	}
	return nil, args.Error(1)
}

type ByNamesTestSuite struct {
	suite.Suite
}

func (suite *ByNamesTestSuite) TestAllInitialized() {
	mockClientGenerator := new(MockClientGenerator)
	mockClientGenerator.On("KubeClient").Return(fake.NewSimpleClientset(), nil)
	mockClientGenerator.On("IstioClient").Return(NewFakeConfigStore(), nil)
	mockClientGenerator.On("ContourClient").Return(fakeContour.NewSimpleClientset(), nil)

	sources, err := ByNames(mockClientGenerator, []string{"service", "ingress", "istio-gateway", "contour-ingressroute", "fake"}, minimalConfig)
	suite.NoError(err, "should not generate errors")
	suite.Len(sources, 5, "should generate all five sources")
}

func (suite *ByNamesTestSuite) TestOnlyFake() {
	mockClientGenerator := new(MockClientGenerator)
	mockClientGenerator.On("KubeClient").Return(fake.NewSimpleClientset(), nil)

	sources, err := ByNames(mockClientGenerator, []string{"fake"}, minimalConfig)
	suite.NoError(err, "should not generate errors")
	suite.Len(sources, 1, "should generate fake source")
	suite.Nil(mockClientGenerator.kubeClient, "client should not be created")
}

func (suite *ByNamesTestSuite) TestSourceNotFound() {
	mockClientGenerator := new(MockClientGenerator)
	mockClientGenerator.On("KubeClient").Return(fake.NewSimpleClientset(), nil)

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
}

func (suite *ByNamesTestSuite) TestIstioClientFails() {
	mockClientGenerator := new(MockClientGenerator)
	mockClientGenerator.On("KubeClient").Return(fake.NewSimpleClientset(), nil)
	mockClientGenerator.On("IstioClient").Return(nil, errors.New("foo"))
	mockClientGenerator.On("ContourClient").Return(nil, errors.New("foo"))

	_, err := ByNames(mockClientGenerator, []string{"istio-gateway"}, minimalConfig)
	suite.Error(err, "should return an error if istio client cannot be created")

	_, err = ByNames(mockClientGenerator, []string{"contour-ingressroute"}, minimalConfig)
	suite.Error(err, "should return an error if contour client cannot be created")
}

func TestByNames(t *testing.T) {
	suite.Run(t, new(ByNamesTestSuite))
}

var minimalConfig = &Config{
	IstioIngressGatewayServices: []string{"istio-system/istio-ingressgateway"},
	ContourLoadBalancerService:  "heptio-contour/contour",
}
