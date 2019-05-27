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
	"net/http"
	"os"
	"strings"
	"time"

	"sync"

	cfclient "github.com/cloudfoundry-community/go-cfclient"
	"github.com/linki/instrumented_http"
	log "github.com/sirupsen/logrus"
	istiocrd "istio.io/istio/pilot/pkg/config/kube/crd"
	istiomodel "istio.io/istio/pilot/pkg/model"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// ErrSourceNotFound is returned when a requested source doesn't exist.
var ErrSourceNotFound = errors.New("source not found")

// Config holds shared configuration options for all Sources.
type Config struct {
	Namespace                   string
	AnnotationFilter            string
	FQDNTemplate                string
	CombineFQDNAndAnnotation    bool
	IgnoreHostnameAnnotation    bool
	Compatibility               string
	PublishInternal             bool
	PublishHostIP               bool
	ConnectorServer             string
	CRDSourceAPIVersion         string
	CRDSourceKind               string
	KubeConfig                  string
	KubeMaster                  string
	ServiceTypeFilter           []string
	IstioIngressGatewayServices []string
	CFAPIEndpoint               string
	CFUsername                  string
	CFPassword                  string
}

// ClientGenerator provides clients
type ClientGenerator interface {
	KubeClient() (kubernetes.Interface, error)
	IstioClient() (istiomodel.ConfigStore, error)
	CloudFoundryClient(cfAPPEndpoint string, cfUsername string, cfPassword string) (*cfclient.Client, error)
}

// SingletonClientGenerator stores provider clients and guarantees that only one instance of client
// will be generated
type SingletonClientGenerator struct {
	KubeConfig     string
	KubeMaster     string
	RequestTimeout time.Duration
	kubeClient     kubernetes.Interface
	istioClient    istiomodel.ConfigStore
	cfClient       *cfclient.Client
	kubeOnce       sync.Once
	istioOnce      sync.Once
	cfOnce         sync.Once
}

// KubeClient generates a kube client if it was not created before
func (p *SingletonClientGenerator) KubeClient() (kubernetes.Interface, error) {
	var err error
	p.kubeOnce.Do(func() {
		p.kubeClient, err = NewKubeClient(p.KubeConfig, p.KubeMaster, p.RequestTimeout)
	})
	return p.kubeClient, err
}

// IstioClient generates an istio client if it was not created before
func (p *SingletonClientGenerator) IstioClient() (istiomodel.ConfigStore, error) {
	var err error
	p.istioOnce.Do(func() {
		p.istioClient, err = NewIstioClient(p.KubeConfig)
	})
	return p.istioClient, err
}

// CloudFoundryClient generates a cf client if it was not created before
func (p *SingletonClientGenerator) CloudFoundryClient(cfAPIEndpoint string, cfUsername string, cfPassword string) (*cfclient.Client, error) {
	var err error
	p.cfOnce.Do(func() {
		p.cfClient, err = NewCFClient(cfAPIEndpoint, cfUsername, cfPassword)
	})
	return p.cfClient, err
}

// NewCFClient return a new CF client object.
func NewCFClient(cfAPIEndpoint string, cfUsername string, cfPassword string) (*cfclient.Client, error) {
	c := &cfclient.Config{
		ApiAddress: "https://" + cfAPIEndpoint,
		Username:   cfUsername,
		Password:   cfPassword,
	}
	client, err := cfclient.NewClient(c)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// ByNames returns multiple Sources given multiple names.
func ByNames(p ClientGenerator, names []string, cfg *Config) ([]Source, error) {
	sources := []Source{}
	for _, name := range names {
		source, err := BuildWithConfig(name, p, cfg)
		if err != nil {
			return nil, err
		}
		sources = append(sources, source)
	}

	return sources, nil
}

// BuildWithConfig allows to generate a Source implementation from the shared config
func BuildWithConfig(source string, p ClientGenerator, cfg *Config) (Source, error) {
	switch source {
	case "service":
		client, err := p.KubeClient()
		if err != nil {
			return nil, err
		}
		return NewServiceSource(client, cfg.Namespace, cfg.AnnotationFilter, cfg.FQDNTemplate, cfg.CombineFQDNAndAnnotation, cfg.Compatibility, cfg.PublishInternal, cfg.PublishHostIP, cfg.ServiceTypeFilter, cfg.IgnoreHostnameAnnotation)
	case "ingress":
		client, err := p.KubeClient()
		if err != nil {
			return nil, err
		}
		return NewIngressSource(client, cfg.Namespace, cfg.AnnotationFilter, cfg.FQDNTemplate, cfg.CombineFQDNAndAnnotation, cfg.IgnoreHostnameAnnotation)
	case "istio-gateway":
		kubernetesClient, err := p.KubeClient()
		if err != nil {
			return nil, err
		}
		istioClient, err := p.IstioClient()
		if err != nil {
			return nil, err
		}
		return NewIstioGatewaySource(kubernetesClient, istioClient, cfg.IstioIngressGatewayServices, cfg.Namespace, cfg.AnnotationFilter, cfg.FQDNTemplate, cfg.CombineFQDNAndAnnotation, cfg.IgnoreHostnameAnnotation)

	case "istio-virtual-service":
		kubernetesClient, err := p.KubeClient()
		if err != nil {
			return nil, err
		}
		istioClient, err := p.IstioClient()
		if err != nil {
			return nil, err
		}
		return NewIstioVirtualServiceSource(kubernetesClient, istioClient, cfg.IstioIngressGatewayServices, cfg.Namespace, cfg.AnnotationFilter, cfg.FQDNTemplate, cfg.CombineFQDNAndAnnotation, cfg.IgnoreHostnameAnnotation)
	case "cloudfoundry":
		cfClient, err := p.CloudFoundryClient(cfg.CFAPIEndpoint, cfg.CFUsername, cfg.CFPassword)
		if err != nil {
			return nil, err
		}
		return NewCloudFoundrySource(cfClient)
	case "fake":
		return NewFakeSource(cfg.FQDNTemplate)
	case "connector":
		return NewConnectorSource(cfg.ConnectorServer)
	case "crd":
		client, err := p.KubeClient()
		if err != nil {
			return nil, err
		}
		crdClient, scheme, err := NewCRDClientForAPIVersionKind(client, cfg.KubeConfig, cfg.KubeMaster, cfg.CRDSourceAPIVersion, cfg.CRDSourceKind)
		if err != nil {
			return nil, err
		}
		return NewCRDSource(crdClient, cfg.Namespace, cfg.CRDSourceKind, scheme)
	}
	return nil, ErrSourceNotFound
}

// NewKubeClient returns a new Kubernetes client object. It takes a Config and
// uses KubeMaster and KubeConfig attributes to connect to the cluster. If
// KubeConfig isn't provided it defaults to using the recommended default.
func NewKubeClient(kubeConfig, kubeMaster string, requestTimeout time.Duration) (*kubernetes.Clientset, error) {
	if kubeConfig == "" {
		if _, err := os.Stat(clientcmd.RecommendedHomeFile); err == nil {
			kubeConfig = clientcmd.RecommendedHomeFile
		}
	}

	config, err := clientcmd.BuildConfigFromFlags(kubeMaster, kubeConfig)
	if err != nil {
		return nil, err
	}

	config.WrapTransport = func(rt http.RoundTripper) http.RoundTripper {
		return instrumented_http.NewTransport(rt, &instrumented_http.Callbacks{
			PathProcessor: func(path string) string {
				parts := strings.Split(path, "/")
				return parts[len(parts)-1]
			},
		})
	}

	config.Timeout = requestTimeout

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	log.Infof("Created Kubernetes client %s", config.Host)

	return client, nil
}

// NewIstioClient returns a new Istio client object. It uses the configured
// KubeConfig attribute to connect to the cluster. If KubeConfig isn't provided
// it defaults to using the recommended default.
// NB: Istio controls the creation of the underlying Kubernetes client, so we
// have no ability to tack on transport wrappers (e.g., Prometheus request
// wrappers) to the client's config at this level. Furthermore, the Istio client
// constructor does not expose the ability to override the Kubernetes master,
// so the Master config attribute has no effect.
func NewIstioClient(kubeConfig string) (*istiocrd.Client, error) {
	if kubeConfig == "" {
		if _, err := os.Stat(clientcmd.RecommendedHomeFile); err == nil {
			kubeConfig = clientcmd.RecommendedHomeFile
		}
	}

	client, err := istiocrd.NewClient(
		kubeConfig,
		"",
		istiomodel.ConfigDescriptor{istiomodel.Gateway},
		"",
	)
	if err != nil {
		return nil, err
	}

	log.Info("Created Istio client")

	return client, nil
}

func NewIstioVirtualServiceClient(kubeConfig string) (*istiocrd.Client, error) {
	if kubeConfig == "" {
		if _, err := os.Stat(clientcmd.RecommendedHomeFile); err == nil {
			kubeConfig = clientcmd.RecommendedHomeFile
		}
	}

	client, err := istiocrd.NewClient(
		kubeConfig,
		"",
		istiomodel.ConfigDescriptor{istiomodel.VirtualService},
		"",
	)
	if err != nil {
		return nil, err
	}

	log.Info("Created Istio client")

	return client, nil
}
