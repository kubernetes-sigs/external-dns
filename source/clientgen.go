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
	"fmt"
	"sync"
	"time"

	openshift "github.com/openshift/client-go/route/clientset/versioned"
	log "github.com/sirupsen/logrus"
	istioclient "istio.io/client-go/pkg/clientset/versioned"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	gateway "sigs.k8s.io/gateway-api/pkg/client/clientset/versioned"

	kubeclient "sigs.k8s.io/external-dns/pkg/client"
	"sigs.k8s.io/external-dns/source/config"
)

// ClientGenerator provides clients for various Kubernetes APIs and external services.
// This interface abstracts client creation and enables dependency injection for testing.
// It uses the singleton pattern to ensure only one instance of each client is created
// and reused across multiple source instances.
//
// Supported Client Types:
// - KubeClient: Standard Kubernetes API client
// - GatewayClient: Gateway API client for Gateway resources
// - IstioClient: Istio service mesh client
// - DynamicKubernetesClient: Dynamic client for custom resources
// - OpenShiftClient: OpenShift-specific client for Route resources
// - RESTConfig: Instrumented REST config for creating custom clients
//
// The singleton behavior is implemented in SingletonClientGenerator which uses
// sync.Once to guarantee single initialization of each client type.
type ClientGenerator interface {
	KubeClient() (kubernetes.Interface, error)
	GatewayClient() (gateway.Interface, error)
	IstioClient() (istioclient.Interface, error)
	DynamicKubernetesClient() (dynamic.Interface, error)
	OpenShiftClient() (openshift.Interface, error)
	RESTConfig() (*rest.Config, error)
}

// WithClientGenerator sets a custom ClientGenerator on a Config, overriding
// the default SingletonClientGenerator that ClientGeneratorFor would build.
// Intended for testing.
func WithClientGenerator(gen ClientGenerator) config.OverrideConfigOption {
	return func(cfg *config.Config) {
		cfg.SetClientGen(gen)
	}
}

// ClientGeneratorFor returns the ClientGenerator for cfg. If one was not
// provided via WithClientGenerator, a SingletonClientGenerator is lazily
// created from cfg's connection settings.
//
// The timeout behavior is special-cased: when UpdateEvents is true, the
// timeout is set to 0 (no timeout) to allow long-running watch operations for
// event-driven source updates.
func ClientGeneratorFor(cfg *config.Config) ClientGenerator {
	gen := cfg.EnsureClientGen(func() any {
		return &SingletonClientGenerator{
			KubeConfig:   cfg.KubeConfig,
			APIServerURL: cfg.APIServerURL,
			RequestTimeout: func() time.Duration {
				if cfg.UpdateEvents {
					return 0
				}
				return cfg.KubeAPIRequestTimeout
			}(),
			QPS:   cfg.KubeAPIQPS,
			Burst: cfg.KubeAPIBurst,
		}
	})
	return gen.(ClientGenerator)
}

// SingletonClientGenerator stores provider clients and guarantees that only one instance of each client
// will be generated throughout the application lifecycle.
//
// Thread Safety: Uses sync.Once for each client type to ensure thread-safe initialization.
// This is important because external-dns may create multiple sources concurrently.
//
// Memory Efficiency: Prevents creating multiple instances of expensive client objects
// that maintain their own connection pools and caches.
//
// Configuration: Clients are configured using KubeConfig, APIServerURL, RequestTimeout,
// QPS, and Burst which are set during SingletonClientGenerator initialization.
type SingletonClientGenerator struct {
	KubeConfig         string
	APIServerURL       string
	RequestTimeout     time.Duration
	QPS                int
	Burst              int
	restConfig         *rest.Config
	restConfigErr      error
	kubeClient         kubernetes.Interface
	kubeClientErr      error
	gatewayClient      gateway.Interface
	gatewayClientErr   error
	istioClient        *istioclient.Clientset
	istioClientErr     error
	dynKubeClient      dynamic.Interface
	dynKubeClientErr   error
	openshiftClient    openshift.Interface
	openshiftClientErr error
	restConfigOnce     sync.Once
	kubeOnce           sync.Once
	gatewayOnce        sync.Once
	istioOnce          sync.Once
	dynCliOnce         sync.Once
	openshiftOnce      sync.Once
}

// KubeClient generates a kube client if it was not created before
func (p *SingletonClientGenerator) KubeClient() (kubernetes.Interface, error) {
	p.kubeOnce.Do(func() {
		var config *rest.Config
		config, p.kubeClientErr = p.RESTConfig()
		if p.kubeClientErr != nil {
			return
		}
		p.kubeClient, p.kubeClientErr = kubeclient.NewKubeClient(config)
	})
	return p.kubeClient, p.kubeClientErr
}

// RESTConfig generates an instrumented REST config if it was not created before.
// The config includes request timeout handling and metrics instrumentation.
// This is useful for sources that need to create custom clients (e.g., controller-runtime clients).
func (p *SingletonClientGenerator) RESTConfig() (*rest.Config, error) {
	p.restConfigOnce.Do(func() {
		p.restConfig, p.restConfigErr = kubeclient.InstrumentedRESTConfig(p.KubeConfig, p.APIServerURL, p.RequestTimeout, p.QPS, p.Burst)
	})
	return p.restConfig, p.restConfigErr
}

// GatewayClient generates a gateway client if it was not created before
func (p *SingletonClientGenerator) GatewayClient() (gateway.Interface, error) {
	p.gatewayOnce.Do(func() {
		var config *rest.Config
		config, p.gatewayClientErr = p.RESTConfig()
		if p.gatewayClientErr != nil {
			return
		}
		p.gatewayClient, p.gatewayClientErr = gateway.NewForConfig(config)
		if p.gatewayClientErr != nil {
			return
		}
		log.Infof("Created GatewayAPI client %s", config.Host)
	})
	return p.gatewayClient, p.gatewayClientErr
}

// IstioClient generates an istio go client if it was not created before
func (p *SingletonClientGenerator) IstioClient() (istioclient.Interface, error) {
	p.istioOnce.Do(func() {
		var config *rest.Config
		config, p.istioClientErr = p.RESTConfig()
		if p.istioClientErr != nil {
			return
		}
		p.istioClient, p.istioClientErr = NewIstioClient(config)
	})
	return p.istioClient, p.istioClientErr
}

// DynamicKubernetesClient generates a dynamic client if it was not created before
func (p *SingletonClientGenerator) DynamicKubernetesClient() (dynamic.Interface, error) {
	p.dynCliOnce.Do(func() {
		var config *rest.Config
		config, p.dynKubeClientErr = p.RESTConfig()
		if p.dynKubeClientErr != nil {
			return
		}
		p.dynKubeClient, p.dynKubeClientErr = dynamic.NewForConfig(config)
		if p.dynKubeClientErr != nil {
			return
		}
		log.Infof("Created Dynamic Kubernetes client %s", config.Host)
	})
	return p.dynKubeClient, p.dynKubeClientErr
}

// OpenShiftClient generates an openshift client if it was not created before
func (p *SingletonClientGenerator) OpenShiftClient() (openshift.Interface, error) {
	p.openshiftOnce.Do(func() {
		var config *rest.Config
		config, p.openshiftClientErr = p.RESTConfig()
		if p.openshiftClientErr != nil {
			return
		}
		p.openshiftClient, p.openshiftClientErr = openshift.NewForConfig(config)
		if p.openshiftClientErr != nil {
			return
		}
		log.Infof("Created OpenShift client %s", config.Host)
	})
	return p.openshiftClient, p.openshiftClientErr
}

// NewIstioClient returns a new Istio client object from the given REST config.
func NewIstioClient(restCfg *rest.Config) (*istioclient.Clientset, error) {
	ic, err := istioclient.NewForConfig(restCfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create istio client: %w", err)
	}

	return ic, nil
}
