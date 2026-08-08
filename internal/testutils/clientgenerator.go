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

package testutils

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

// StubClientGenerator is a ClientGenerator where all methods return errors.
// Use it for sources that don't require any Kubernetes client (e.g. the "fake" source).
type StubClientGenerator = stubClientGenerator

// MockClientGenerator is a full testify mock of source.ClientGenerator.
// Returned client values are stored in the *Value fields so tests can assert
// whether a specific client was actually requested.
type MockClientGenerator struct {
	mock.Mock
	KubeClientValue              kubernetes.Interface
	GatewayClientValue           gateway.Interface
	IstioClientValue             istioclient.Interface
	DynamicKubernetesClientValue dynamic.Interface
	OpenShiftClientValue         openshift.Interface
	RESTConfigValue              *rest.Config
}

// stubClientGenerator implements source.ClientGenerator where all methods
// return errors, except KubeClient when a clientset is provided.
type stubClientGenerator struct {
	kubeClient *fake.Clientset
}

// NewFakeClientGenerator returns a ClientGenerator whose KubeClient returns the
// provided fake clientset. All other methods return errors.
func NewFakeClientGenerator(client *fake.Clientset) stubClientGenerator {
	return stubClientGenerator{kubeClient: client}
}

func (m *MockClientGenerator) KubeClient() (kubernetes.Interface, error) {
	args := m.Called()
	if args.Error(1) == nil {
		m.KubeClientValue = args.Get(0).(kubernetes.Interface)
		return m.KubeClientValue, nil
	}
	return nil, args.Error(1)
}

func (m *MockClientGenerator) GatewayClient() (gateway.Interface, error) {
	args := m.Called()
	if args.Error(1) == nil {
		m.GatewayClientValue = args.Get(0).(gateway.Interface)
		return m.GatewayClientValue, nil
	}
	return nil, args.Error(1)
}

func (m *MockClientGenerator) IstioClient() (istioclient.Interface, error) {
	args := m.Called()
	if args.Error(1) == nil {
		m.IstioClientValue = args.Get(0).(istioclient.Interface)
		return m.IstioClientValue, nil
	}
	return nil, args.Error(1)
}

func (m *MockClientGenerator) DynamicKubernetesClient() (dynamic.Interface, error) {
	args := m.Called()
	if args.Error(1) == nil {
		m.DynamicKubernetesClientValue = args.Get(0).(dynamic.Interface)
		return m.DynamicKubernetesClientValue, nil
	}
	return nil, args.Error(1)
}

func (m *MockClientGenerator) OpenShiftClient() (openshift.Interface, error) {
	args := m.Called()
	if args.Error(1) == nil {
		m.OpenShiftClientValue = args.Get(0).(openshift.Interface)
		return m.OpenShiftClientValue, nil
	}
	return nil, args.Error(1)
}

func (m *MockClientGenerator) RESTConfig() (*rest.Config, error) {
	args := m.Called()
	if args.Error(1) == nil {
		m.RESTConfigValue = args.Get(0).(*rest.Config)
		return m.RESTConfigValue, nil
	}
	return nil, args.Error(1)
}

func (s stubClientGenerator) KubeClient() (kubernetes.Interface, error) {
	if s.kubeClient == nil {
		return nil, errNotAvailable("KubeClient")
	}
	return s.kubeClient, nil
}

func (stubClientGenerator) GatewayClient() (gateway.Interface, error) {
	return nil, errNotAvailable("GatewayClient")
}

func (stubClientGenerator) IstioClient() (istioclient.Interface, error) {
	return nil, errNotAvailable("IstioClient")
}

func (stubClientGenerator) DynamicKubernetesClient() (dynamic.Interface, error) {
	return nil, errNotAvailable("DynamicKubernetesClient")
}

func (stubClientGenerator) OpenShiftClient() (openshift.Interface, error) {
	return nil, errNotAvailable("OpenShiftClient")
}

func (stubClientGenerator) RESTConfig() (*rest.Config, error) {
	return nil, errNotAvailable("RESTConfig")
}

func errNotAvailable(method string) error {
	return fmt.Errorf("%s: not available in stub", method)
}
