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

// Package kubeclient provides shared utilities for creating Kubernetes REST configurations
// and clients with standardized metrics instrumentation.
package kubeclient

import (
	"context"
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/flowcontrol"

	"sigs.k8s.io/external-dns/pkg/apis/externaldns"

	extdnshttp "sigs.k8s.io/external-dns/pkg/http"
)

// InstrumentedRESTConfig builds a REST config with Prometheus transport metrics, request timeout,
// and a token-bucket rate limiter. When qps > 0, it overrides the client-go defaults (5 QPS / 10 burst).
func InstrumentedRESTConfig(
	kubeConfig, apiServerURL string,
	requestTimeout time.Duration,
	qps int, burst int) (*rest.Config, error) {
	config, err := buildRestConfig(kubeConfig, apiServerURL)
	if err != nil {
		return nil, err
	}

	config.UserAgent = externaldns.UserAgent()
	config.WrapTransport = extdnshttp.NewInstrumentedTransport
	config.Timeout = requestTimeout

	config.QPS = rest.DefaultQPS
	config.Burst = rest.DefaultBurst
	if qps > 0 {
		config.QPS = float32(qps)
	}
	if burst > 0 {
		config.Burst = burst
	}
	log.Debugf("kube client qps: %f, burst %d", config.QPS, config.Burst)
	config.RateLimiter = &rateLimiter{
		delegate: flowcontrol.NewTokenBucketRateLimiter(config.QPS, config.Burst),
	}
	return config, nil
}

// NewKubeClient creates a Kubernetes client from the given REST config.
func NewKubeClient(config *rest.Config) (kubernetes.Interface, error) {
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	log.Infof("Created Kubernetes client %s", config.Host)
	return client, nil
}

// buildRestConfig returns the REST client configuration for Kubernetes API access.
// Supports both in-cluster and external cluster configurations.
//
// Configuration Priority:
// 1. KubeConfig file if specified
// 2. Recommended home file (~/.kube/config)
// 3. In-cluster config
// TODO: consider clientcmd.NewDefaultClientConfigLoadingRules() with clientcmd.NewNonInteractiveDeferredLoadingClientConfig
func buildRestConfig(kubeConfig, apiServerURL string) (*rest.Config, error) {
	if kubeConfig == "" {
		if _, err := os.Stat(clientcmd.RecommendedHomeFile); err == nil {
			kubeConfig = clientcmd.RecommendedHomeFile
		}
	}
	log.Debugf("apiServerURL: %s", apiServerURL)
	log.Debugf("kubeConfig: %s", kubeConfig)

	// evaluate whether to use kubeConfig-file or serviceaccount-token
	var (
		config *rest.Config
		err    error
	)
	if kubeConfig == "" {
		log.Debug("Using inCluster-config based on serviceaccount-token")
		config, err = rest.InClusterConfig()
	} else {
		log.Debug("Using kubeConfig")
		config, err = clientcmd.BuildConfigFromFlags(apiServerURL, kubeConfig)
	}
	if err != nil {
		return nil, err
	}

	return config, nil
}

// rateLimiter wraps a RateLimiter and enriches Wait errors with an
// actionable hint so callers get a clear message without needing to inspect
// the error string themselves.
type rateLimiter struct {
	delegate flowcontrol.RateLimiter
}

func (r *rateLimiter) TryAccept() bool { return r.delegate.TryAccept() }
func (r *rateLimiter) Accept()         { r.delegate.Accept() }
func (r *rateLimiter) Stop()           { r.delegate.Stop() }
func (r *rateLimiter) QPS() float32    { return r.delegate.QPS() }

// Wait blocks until a token is available or the context is done.
// Any error from Wait is a rate limit timeout; it is enriched with an actionable hint.
func (r *rateLimiter) Wait(ctx context.Context) error {
	if err := r.delegate.Wait(ctx); err != nil {
		return fmt.Errorf("consider raising --kube-api-qps/--kube-api-burst: %w", err)
	}
	return nil
}
