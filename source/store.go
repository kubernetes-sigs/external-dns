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
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	cloudfoundry "github.com/cloudfoundry-community/go-cfclient"
	"github.com/linki/instrumented_http"
	openshift "github.com/openshift/client-go/route/clientset/versioned"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	istio "istio.io/client-go/pkg/clientset/versioned"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// ErrSourceNotFound is returned when a requested source doesn't exist.
var ErrSourceNotFound = errors.New("source not found")

// Config holds shared configuration options for all Sources.
type Config struct {
	Namespace                      string
	AnnotationFilter               string
	LabelFilter                    string
	FQDNTemplate                   string
	CombineFQDNAndAnnotation       bool
	IgnoreHostnameAnnotation       bool
	IgnoreIngressTLSSpec           bool
	IgnoreIngressRulesSpec         bool
	Compatibility                  string
	PublishInternal                bool
	PublishHostIP                  bool
	AlwaysPublishNotReadyAddresses bool
	ConnectorServer                string
	CRDSourceAPIVersion            string
	CRDSourceKind                  string
	KubeConfig                     string
	APIServerURL                   string
	ServiceTypeFilter              []string
	CFAPIEndpoint                  string
	CFUsername                     string
	CFPassword                     string
	ContourLoadBalancerService     string
	GlooNamespace                  string
	SkipperRouteGroupVersion       string
	RequestTimeout                 time.Duration
	DefaultTargets                 []string
}

// ClientGenerator provides clients.
type ClientGenerator interface {
	RESTConfig() (*rest.Config, error)
	KubeClient() (kubernetes.Interface, error)
	IstioClient() (istio.Interface, error)
	CloudFoundryClient(endpoint string, username string, password string) (*cloudfoundry.Client, error)
	DynamicKubernetesClient() (dynamic.Interface, error)
	OpenShiftClient() (openshift.Interface, error)
}

// SingletonClientGenerator stores provider clients and guarantees that only one instance of client
// will be generated
type SingletonClientGenerator struct {
	KubeConfig     string
	APIServerURL   string
	RequestTimeout time.Duration

	restConfig         onceWithError
	kubeClient         onceWithError
	istioClient        onceWithError
	cloudFoundryClient onceWithError
	dynKubeClient      onceWithError
	openshiftClient    onceWithError
}

type onceWithError struct {
	once sync.Once
	val  interface{}
	err  error
}

func (o *onceWithError) Do(fn func() (interface{}, error)) (interface{}, error) {
	o.once.Do(func() {
		o.val, o.err = fn()
	})
	return o.val, o.err
}

func (p *SingletonClientGenerator) RESTConfig() (*rest.Config, error) {
	iface, err := p.restConfig.Do(func() (interface{}, error) {
		log.Infof("Instantiating new Kubernetes REST config")
		cfg, err := restConfig(p.KubeConfig, p.APIServerURL, p.RequestTimeout)
		if err == nil {
			log.Infof("Created new Kubernetes REST config %s", cfg.Host)
		}
		return cfg, err
	})
	val, _ := iface.(*rest.Config)
	return val, err
}

func restConfig(kubeConfig, apiServerURL string, requestTimeout time.Duration) (*rest.Config, error) {
	if kubeConfig == "" {
		if _, err := os.Stat(clientcmd.RecommendedHomeFile); err == nil {
			kubeConfig = clientcmd.RecommendedHomeFile
		}
	}
	log.Debugf("apiServerURL: %s", apiServerURL)
	log.Debugf("kubeConfig: %s", kubeConfig)
	cfg, err := clientcmd.BuildConfigFromFlags(apiServerURL, kubeConfig)
	if err != nil {
		return nil, err
	}
	cfg.WrapTransport = func(rt http.RoundTripper) http.RoundTripper {
		return instrumented_http.NewTransport(rt, &instrumented_http.Callbacks{
			PathProcessor: func(path string) string {
				parts := strings.Split(path, "/")
				return parts[len(parts)-1]
			},
		})
	}
	cfg.Timeout = requestTimeout
	return cfg, nil
}

// KubeClient generates a kube client if it was not created before
func (p *SingletonClientGenerator) KubeClient() (kubernetes.Interface, error) {
	iface, err := p.kubeClient.Do(func() (interface{}, error) {
		log.Infof("Instantiating new Kubernetes client")
		cfg, err := p.RESTConfig()
		if err != nil {
			return nil, err
		}
		return kubernetes.NewForConfig(cfg)
	})
	val, _ := iface.(kubernetes.Interface)
	return val, err
}

// IstioClient generates an istio go client if it was not created before
func (p *SingletonClientGenerator) IstioClient() (istio.Interface, error) {
	iface, err := p.istioClient.Do(func() (interface{}, error) {
		log.Infof("Instantiating new Istio client")
		cfg, err := p.RESTConfig()
		if err != nil {
			return nil, err
		}
		return istio.NewForConfig(cfg)
	})
	val, _ := iface.(istio.Interface)
	return val, err
}

// CloudFoundryClient generates a cf client if it was not created before
func (p *SingletonClientGenerator) CloudFoundryClient(endpoint string, username string, password string) (*cloudfoundry.Client, error) {
	iface, err := p.cloudFoundryClient.Do(func() (interface{}, error) {
		log.Infof("Instantiating new CloudFoundry client")
		c := &cloudfoundry.Config{
			ApiAddress: "https://" + endpoint,
			Username:   username,
			Password:   password,
		}
		client, err := cloudfoundry.NewClient(c)
		if err != nil {
			return nil, err
		}
		return client, nil
	})
	val, _ := iface.(*cloudfoundry.Client)
	return val, err
}

// DynamicKubernetesClient generates a dynamic client if it was not created before
func (p *SingletonClientGenerator) DynamicKubernetesClient() (dynamic.Interface, error) {
	iface, err := p.dynKubeClient.Do(func() (interface{}, error) {
		log.Infof("Instantiating new Dynamic Kubernetes client")
		cfg, err := p.RESTConfig()
		if err != nil {
			return nil, err
		}
		return dynamic.NewForConfig(cfg)
	})
	val, _ := iface.(dynamic.Interface)
	return val, err
}

// OpenShiftClient generates an openshift client if it was not created before
func (p *SingletonClientGenerator) OpenShiftClient() (openshift.Interface, error) {
	iface, err := p.openshiftClient.Do(func() (interface{}, error) {
		log.Infof("Instantiating new Openshift client")
		cfg, err := p.RESTConfig()
		if err != nil {
			return nil, err
		}
		return openshift.NewForConfig(cfg)
	})
	val, _ := iface.(openshift.Interface)
	return val, err
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
	case "node":
		client, err := p.KubeClient()
		if err != nil {
			return nil, err
		}
		return NewNodeSource(client, cfg.AnnotationFilter, cfg.FQDNTemplate)
	case "service":
		client, err := p.KubeClient()
		if err != nil {
			return nil, err
		}
		return NewServiceSource(client, cfg.Namespace, cfg.AnnotationFilter, cfg.FQDNTemplate, cfg.CombineFQDNAndAnnotation, cfg.Compatibility, cfg.PublishInternal, cfg.PublishHostIP, cfg.AlwaysPublishNotReadyAddresses, cfg.ServiceTypeFilter, cfg.IgnoreHostnameAnnotation)
	case "ingress":
		client, err := p.KubeClient()
		if err != nil {
			return nil, err
		}
		return NewIngressSource(client, cfg.Namespace, cfg.AnnotationFilter, cfg.FQDNTemplate, cfg.CombineFQDNAndAnnotation, cfg.IgnoreHostnameAnnotation, cfg.IgnoreIngressTLSSpec, cfg.IgnoreIngressRulesSpec)
	case "pod":
		client, err := p.KubeClient()
		if err != nil {
			return nil, err
		}
		return NewPodSource(client, cfg.Namespace, cfg.Compatibility)
	case "istio-gateway":
		kubernetesClient, err := p.KubeClient()
		if err != nil {
			return nil, err
		}
		istioClient, err := p.IstioClient()
		if err != nil {
			return nil, err
		}
		return NewIstioGatewaySource(kubernetesClient, istioClient, cfg.Namespace, cfg.AnnotationFilter, cfg.FQDNTemplate, cfg.CombineFQDNAndAnnotation, cfg.IgnoreHostnameAnnotation)
	case "istio-virtualservice":
		kubernetesClient, err := p.KubeClient()
		if err != nil {
			return nil, err
		}
		istioClient, err := p.IstioClient()
		if err != nil {
			return nil, err
		}
		return NewIstioVirtualServiceSource(kubernetesClient, istioClient, cfg.Namespace, cfg.AnnotationFilter, cfg.FQDNTemplate, cfg.CombineFQDNAndAnnotation, cfg.IgnoreHostnameAnnotation)
	case "cloudfoundry":
		cloudFoundryClient, err := p.CloudFoundryClient(cfg.CFAPIEndpoint, cfg.CFUsername, cfg.CFPassword)
		if err != nil {
			return nil, err
		}
		return NewCloudFoundrySource(cloudFoundryClient)
	case "ambassador-host":
		kubernetesClient, err := p.KubeClient()
		if err != nil {
			return nil, err
		}
		dynamicClient, err := p.DynamicKubernetesClient()
		if err != nil {
			return nil, err
		}
		return NewAmbassadorHostSource(dynamicClient, kubernetesClient, cfg.Namespace)
	case "contour-ingressroute":
		kubernetesClient, err := p.KubeClient()
		if err != nil {
			return nil, err
		}
		dynamicClient, err := p.DynamicKubernetesClient()
		if err != nil {
			return nil, err
		}
		return NewContourIngressRouteSource(dynamicClient, kubernetesClient, cfg.ContourLoadBalancerService, cfg.Namespace, cfg.AnnotationFilter, cfg.FQDNTemplate, cfg.CombineFQDNAndAnnotation, cfg.IgnoreHostnameAnnotation)
	case "contour-httpproxy":
		dynamicClient, err := p.DynamicKubernetesClient()
		if err != nil {
			return nil, err
		}
		return NewContourHTTPProxySource(dynamicClient, cfg.Namespace, cfg.AnnotationFilter, cfg.FQDNTemplate, cfg.CombineFQDNAndAnnotation, cfg.IgnoreHostnameAnnotation)
	case "gloo-proxy":
		kubernetesClient, err := p.KubeClient()
		if err != nil {
			return nil, err
		}
		dynamicClient, err := p.DynamicKubernetesClient()
		if err != nil {
			return nil, err
		}
		return NewGlooSource(dynamicClient, kubernetesClient, cfg.GlooNamespace)
	case "openshift-route":
		ocpClient, err := p.OpenShiftClient()
		if err != nil {
			return nil, err
		}
		return NewOcpRouteSource(ocpClient, cfg.Namespace, cfg.AnnotationFilter, cfg.FQDNTemplate, cfg.CombineFQDNAndAnnotation, cfg.IgnoreHostnameAnnotation)
	case "fake":
		return NewFakeSource(cfg.FQDNTemplate)
	case "connector":
		return NewConnectorSource(cfg.ConnectorServer)
	case "crd":
		client, err := p.KubeClient()
		if err != nil {
			return nil, err
		}
		crdClient, scheme, err := NewCRDClientForAPIVersionKind(client, cfg.KubeConfig, cfg.APIServerURL, cfg.CRDSourceAPIVersion, cfg.CRDSourceKind)
		if err != nil {
			return nil, err
		}
		return NewCRDSource(crdClient, cfg.Namespace, cfg.CRDSourceKind, cfg.AnnotationFilter, cfg.LabelFilter, scheme)
	case "skipper-routegroup":
		apiServerURL := cfg.APIServerURL
		tokenPath := ""
		token := ""
		if c, err := p.RESTConfig(); err == nil {
			apiServerURL = c.Host
			tokenPath = c.BearerTokenFile
			token = c.BearerToken
		}
		return NewRouteGroupSource(cfg.RequestTimeout, token, tokenPath, apiServerURL, cfg.Namespace, cfg.AnnotationFilter, cfg.FQDNTemplate, cfg.SkipperRouteGroupVersion, cfg.CombineFQDNAndAnnotation, cfg.IgnoreHostnameAnnotation)
	case "kong-tcpingress":
		kubernetesClient, err := p.KubeClient()
		if err != nil {
			return nil, err
		}
		dynamicClient, err := p.DynamicKubernetesClient()
		if err != nil {
			return nil, err
		}
		return NewKongTCPIngressSource(dynamicClient, kubernetesClient, cfg.Namespace, cfg.AnnotationFilter)
	}
	return nil, ErrSourceNotFound
}
