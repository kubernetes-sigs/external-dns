/*
Copyright 2025 The Kubernetes Authors.

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

package wrappers

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/source"
)

type postProcessor struct {
	source source.Source
	cfg    PostProcessorConfig
}

type PostProcessorConfig struct {
	ttl          int64
	isConfigured bool
}

type PostProcessorOption func(*PostProcessorConfig)

func WithTTL(ttl time.Duration) PostProcessorOption {
	return func(cfg *PostProcessorConfig) {
		if int64(ttl.Seconds()) > 0 {
			cfg.isConfigured = true
			cfg.ttl = int64(ttl.Seconds())
		}
	}
}

func NewPostProcessor(source source.Source, opts ...PostProcessorOption) source.Source {
	cfg := PostProcessorConfig{}
	for _, opt := range opts {
		opt(&cfg)
	}
	return &postProcessor{source: source, cfg: cfg}
}

func (pp *postProcessor) Endpoints(ctx context.Context) ([]*endpoint.Endpoint, error) {
	endpoints, err := pp.source.Endpoints(ctx)
	if err != nil {
		return nil, err
	}

	if !pp.cfg.isConfigured {
		return endpoints, nil
	}

	for _, ep := range endpoints {
		if ep == nil {
			continue
		}
		ep.WithMinTTL(pp.cfg.ttl)
		// Additional post-processing can be added here.
	}

	return endpoints, nil
}

func (pp *postProcessor) AddEventHandler(ctx context.Context, handler func()) {
	log.Debug("postProcessor: adding event handler")
	pp.source.AddEventHandler(ctx, handler)
}
