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

	log "github.com/Sirupsen/logrus"
	"github.com/linki/instrumented_http"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// ErrSourceNotFound is returned when a requested source doesn't exist.
var ErrSourceNotFound = errors.New("source not found")

// Config holds shared configuration options for all Sources.
type Config struct {
	Namespace     string
	FQDNTemplate  string
	Compatibility string
}

// ClientGenerator provides clients
type ClientGenerator interface {
	KubeClient() (kubernetes.Interface, error)
}

// SingletonClientGenerator stores provider clients and guarantees that only one instance of client
// will be generated
type SingletonClientGenerator struct {
	KubeConfig string
	KubeMaster string
	client     kubernetes.Interface
	sync.Once
}

// KubeClient generates a kube client if it was not created before
func (p *SingletonClientGenerator) KubeClient() (kubernetes.Interface, error) {
	var err error
	p.Once.Do(func() {
		p.client, err = NewKubeClient(p.KubeConfig, p.KubeMaster)
	})
	return p.client, err
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
		return NewServiceSource(client, cfg.Namespace, cfg.FQDNTemplate, cfg.Compatibility)
	case "ingress":
		client, err := p.KubeClient()
		if err != nil {
			return nil, err
		}
		return NewIngressSource(client, cfg.Namespace, cfg.FQDNTemplate)
	case "fake":
		return NewFakeSource(cfg.FQDNTemplate)
	}
	return nil, ErrSourceNotFound
}

// NewKubeClient returns a new Kubernetes client object. It takes a Config and
// uses KubeMaster and KubeConfig attributes to connect to the cluster. If
// KubeConfig isn't provided it defaults to using the recommended default.
func NewKubeClient(kubeConfig, kubeMaster string) (*kubernetes.Clientset, error) {
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

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	log.Infof("Connected to cluster at %s", config.Host)

	return client, nil
}
