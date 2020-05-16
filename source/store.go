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
	"sync"
	"time"

	"github.com/cloudfoundry-community/go-cfclient"
	"github.com/linki/instrumented_http"
	openshift "github.com/openshift/client-go/route/clientset/versioned"
	log "github.com/sirupsen/logrus"
	istiocontroller "istio.io/istio/pilot/pkg/config/kube/crd/controller"
	istiomodel "istio.io/istio/pilot/pkg/model"
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
	FQDNTemplate                   string
	CombineFQDNAndAnnotation       bool
	IgnoreHostnameAnnotation       bool
	Compatibility                  string
	PublishInternal                bool
	PublishHostIP                  bool
	AlwaysPublishNotReadyAddresses bool
	ConnectorServer                string
	CRDSourceAPIVersion            string
	CRDSourceKind                  string
	KubeConfig                     string
	KubeMaster                     string
	ServiceTypeFilter              []string
	IstioIngressGatewayServices    []string
	CFAPIEndpoint                  string
	CFUsername                     string
	CFPassword                     string
	ContourLoadBalancerService     string
	SkipperRouteGroupVersion       string
	RequestTimeout                 time.Duration
}

// ClientGenerator provides clients
type ClientGenerator interface {
	KubeClient() (kubernetes.Interface, error)
	IstioClient() (istiomodel.ConfigStore, error)
	CloudFoundryClient(cfAPPEndpoint string, cfUsername string, cfPassword string) (*cfclient.Client, error)
	DynamicKubernetesClient() (dynamic.Interface, error)
	OpenShiftClient() (openshift.Interface, error)
}

// SingletonClientGenerator stores provider clients and guarantees that only one instance of client
// will be generated
type SingletonClientGenerator struct {
	KubeConfig      string
	KubeMaster      string
	RequestTimeout  time.Duration
	kubeClient      kubernetes.Interface
	istioClient     istiomodel.ConfigStore
	cfClient        *cfclient.Client
	contourClient   dynamic.Interface
	openshiftClient openshift.Interface
	kubeOnce        sync.Once
	istioOnce       sync.Once
	cfOnce          sync.Once
	contourOnce     sync.Once
	openshiftOnce   sync.Once
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

// DynamicKubernetesClient generates a contour client if it was not created before
func (p *SingletonClientGenerator) DynamicKubernetesClient() (dynamic.Interface, error) {
	var err error
	p.contourOnce.Do(func() {
		p.contourClient, err = NewDynamicKubernetesClient(p.KubeConfig, p.KubeMaster, p.RequestTimeout)
	})
	return p.contourClient, err
}

// OpenShiftClient generates an openshift client if it was not created before
func (p *SingletonClientGenerator) OpenShiftClient() (openshift.Interface, error) {
	var err error
	p.openshiftOnce.Do(func() {
		p.openshiftClient, err = NewOpenShiftClient(p.KubeConfig, p.KubeMaster, p.RequestTimeout)
	})
	return p.openshiftClient, err
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
		return NewIstioGatewaySource(kubernetesClient, istioClient, cfg.Namespace, cfg.AnnotationFilter, cfg.FQDNTemplate, cfg.CombineFQDNAndAnnotation, cfg.IgnoreHostnameAnnotation)
	case "cloudfoundry":
		cfClient, err := p.CloudFoundryClient(cfg.CFAPIEndpoint, cfg.CFUsername, cfg.CFPassword)
		if err != nil {
			return nil, err
		}
		return NewCloudFoundrySource(cfClient)
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
		crdClient, scheme, err := NewCRDClientForAPIVersionKind(client, cfg.KubeConfig, cfg.KubeMaster, cfg.CRDSourceAPIVersion, cfg.CRDSourceKind)
		if err != nil {
			return nil, err
		}
		return NewCRDSource(crdClient, cfg.Namespace, cfg.CRDSourceKind, cfg.AnnotationFilter, scheme)
	case "skipper-routegroup":
		master := cfg.KubeMaster
		tokenPath := ""
		token := ""
		restConfig, err := GetRestConfig(cfg.KubeConfig, cfg.KubeMaster)
		if err == nil {
			master = restConfig.Host
			tokenPath = restConfig.BearerTokenFile
			token = restConfig.BearerToken
		}
		return NewRouteGroupSource(cfg.RequestTimeout, token, tokenPath, master, cfg.Namespace, cfg.AnnotationFilter, cfg.FQDNTemplate, cfg.SkipperRouteGroupVersion, cfg.CombineFQDNAndAnnotation, cfg.IgnoreHostnameAnnotation)
	}
	return nil, ErrSourceNotFound
}

// GetRestConfig returns the rest clients config to get automatically
// data if you run inside a cluster or by passing flags.
func GetRestConfig(kubeConfig, kubeMaster string) (*rest.Config, error) {
	if kubeConfig == "" {
		if _, err := os.Stat(clientcmd.RecommendedHomeFile); err == nil {
			kubeConfig = clientcmd.RecommendedHomeFile
		}
	}
	log.Debugf("kubeMaster: %s", kubeMaster)
	log.Debugf("kubeConfig: %s", kubeConfig)

	// evaluate whether to use kubeConfig-file or serviceaccount-token
	var (
		config *rest.Config
		err    error
	)
	if kubeConfig == "" {
		log.Infof("Using inCluster-config based on serviceaccount-token")
		config, err = rest.InClusterConfig()
	} else {
		log.Infof("Using kubeConfig")
		config, err = clientcmd.BuildConfigFromFlags(kubeMaster, kubeConfig)
	}
	if err != nil {
		return nil, err
	}

	return config, nil
}

// NewKubeClient returns a new Kubernetes client object. It takes a Config and
// uses KubeMaster and KubeConfig attributes to connect to the cluster. If
// KubeConfig isn't provided it defaults to using the recommended default.
func NewKubeClient(kubeConfig, kubeMaster string, requestTimeout time.Duration) (*kubernetes.Clientset, error) {
	log.Infof("Instantiating new Kubernetes client")
	config, err := GetRestConfig(kubeConfig, kubeMaster)
	if err != nil {
		return nil, err
	}

	config.Timeout = requestTimeout
	config.WrapTransport = func(rt http.RoundTripper) http.RoundTripper {
		return instrumented_http.NewTransport(rt, &instrumented_http.Callbacks{
			PathProcessor: func(path string) string {
				parts := strings.Split(path, "/")
				return parts[len(parts)-1]
			},
		})
	}

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
func NewIstioClient(kubeConfig string) (*istiocontroller.Client, error) {
	if kubeConfig == "" {
		if _, err := os.Stat(clientcmd.RecommendedHomeFile); err == nil {
			kubeConfig = clientcmd.RecommendedHomeFile
		}
	}

	client, err := istiocontroller.NewClient(
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

// NewDynamicKubernetesClient returns a new Dynamic Kubernetes client object. It takes a Config and
// uses KubeMaster and KubeConfig attributes to connect to the cluster. If
// KubeConfig isn't provided it defaults to using the recommended default.
func NewDynamicKubernetesClient(kubeConfig, kubeMaster string, requestTimeout time.Duration) (dynamic.Interface, error) {
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

	client, err := dynamic.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	log.Infof("Created Dynamic Kubernetes client %s", config.Host)

	return client, nil
}

// NewOpenShiftClient returns a new Openshift client object. It takes a Config and
// uses KubeMaster and KubeConfig attributes to connect to the cluster. If
// KubeConfig isn't provided it defaults to using the recommended default.
func NewOpenShiftClient(kubeConfig, kubeMaster string, requestTimeout time.Duration) (*openshift.Clientset, error) {
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

	client, err := openshift.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	log.Infof("Created OpenShift client %s", config.Host)

	return client, nil
}
