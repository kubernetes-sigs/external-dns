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
	KubeMaster    string
	KubeConfig    string
	Namespace     string
	FQDNTemplate  string
	Compatibility string
}

type clientProvider struct {
	kubeConfig string
	kubeMaster string
	client     kubernetes.Interface
	sync.Once
}

func (p *clientProvider) kubeClient() (kubernetes.Interface, error) {
	var err error
	p.Once.Do(func() {
		p.client, err = NewKubeClient(p.kubeConfig, p.kubeMaster)
	})
	return p.client, err
}

// ByNames returns multiple Sources given multiple names.
func ByNames(names []string, cfg *Config) ([]Source, error) {
	sources := []Source{}

	for _, name := range names {
		source, err := BuildWithConfig(name, cfg)
		if err != nil {
			return nil, err
		}
		sources = append(sources, source)
	}

	return sources, nil
}

func BuildWithConfig(source string, cfg *Config) (Source, error) {
	p := &clientProvider{
		kubeConfig: cfg.KubeConfig,
		kubeMaster: cfg.KubeMaster,
	}
	switch source {
	case "service":
		client, err := p.kubeClient()
		if err != nil {
			return nil, err
		}
		return NewServiceSource(client, cfg.FQDNTemplate, cfg.Namespace, cfg.Compatibility)
	case "ingress":
		client, err := p.kubeClient()
		if err != nil {
			return nil, err
		}
		return NewIngressSource(client, cfg.FQDNTemplate, cfg.Namespace)
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
