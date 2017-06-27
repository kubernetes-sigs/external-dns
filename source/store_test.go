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

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/fake"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockClientProvider struct {
	mock.Mock
	client kubernetes.Interface
}

func (m *MockClientProvider) KubeClient() (kubernetes.Interface, error) {
	args := m.Called()
	if args.Error(1) == nil {
		m.client = args.Get(0).(kubernetes.Interface)
		return m.client, nil
	}
	return nil, args.Error(1)
}

type ByNamesTestSuite struct {
	suite.Suite
}

func (suite *ByNamesTestSuite) TestAllInitialized() {
	mockClientProvider := new(MockClientProvider)
	mockClientProvider.On("KubeClient").Return(fake.NewSimpleClientset(), nil)

	sources, err := ByNames(mockClientProvider, []string{"service", "ingress", "fake"}, &Config{})
	suite.NoError(err, "should not generate errors")
	suite.Len(sources, 3, "should generate all three sources")
}

func (suite *ByNamesTestSuite) TestOnlyFake() {
	mockClientProvider := new(MockClientProvider)
	mockClientProvider.On("KubeClient").Return(fake.NewSimpleClientset(), nil)

	sources, err := ByNames(mockClientProvider, []string{"fake"}, &Config{})
	suite.NoError(err, "should not generate errors")
	suite.Len(sources, 1, "should generate all three sources")
	suite.Nil(mockClientProvider.client, "client should not be created")
}

func (suite *ByNamesTestSuite) TestSourceNotFound() {
	mockClientProvider := new(MockClientProvider)
	mockClientProvider.On("KubeClient").Return(fake.NewSimpleClientset(), nil)

	sources, err := ByNames(mockClientProvider, []string{"foo"}, &Config{})
	suite.Equal(err, ErrSourceNotFound, "should return sourcen not found")
	suite.Len(sources, 0, "should not returns any source")
}

func (suite *ByNamesTestSuite) TestKubeClientFails() {
	mockClientProvider := new(MockClientProvider)
	mockClientProvider.On("KubeClient").Return(nil, errors.New("foo"))

	_, err := ByNames(mockClientProvider, []string{"service"}, &Config{})
	suite.Error(err, "should return an error if client cannot be created")

	_, err = ByNames(mockClientProvider, []string{"ingress"}, &Config{})
	suite.Error(err, "should return an error if client cannot be created")
}

func TestByNames(t *testing.T) {
	suite.Run(t, new(ByNamesTestSuite))
}
