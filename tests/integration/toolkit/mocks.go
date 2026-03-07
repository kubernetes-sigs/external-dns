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

package toolkit

import (
	"fmt"

	openshift "github.com/openshift/client-go/route/clientset/versioned"
	"github.com/stretchr/testify/mock"
	istioclient "istio.io/client-go/pkg/clientset/versioned"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/rest"
	gateway "sigs.k8s.io/gateway-api/pkg/client/clientset/versioned"
)

// MockClientGenerator implements source.ClientGenerator for testing.
type MockClientGenerator struct {
	mock.Mock
}

func (m *MockClientGenerator) RESTConfig() (*rest.Config, error) {
	return nil, fmt.Errorf("RESTConfig: not implemented")
}

func (m *MockClientGenerator) KubeClient() (kubernetes.Interface, error) {
	args := m.Called()
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(kubernetes.Interface), nil
}

func (m *MockClientGenerator) GatewayClient() (gateway.Interface, error) {
	args := m.Called()
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(gateway.Interface), nil
}

func (m *MockClientGenerator) IstioClient() (istioclient.Interface, error) {
	args := m.Called()
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(istioclient.Interface), nil
}

func (m *MockClientGenerator) DynamicKubernetesClient() (dynamic.Interface, error) {
	args := m.Called()
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(dynamic.Interface), nil
}

func (m *MockClientGenerator) OpenShiftClient() (openshift.Interface, error) {
	args := m.Called()
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(openshift.Interface), nil
}

// newMockClientGenerator creates a MockClientGenerator that returns the provided fake client.
func newMockClientGenerator(client *fake.Clientset) *MockClientGenerator {
	m := new(MockClientGenerator)
	m.On("KubeClient").Return(client, nil)
	return m
}
