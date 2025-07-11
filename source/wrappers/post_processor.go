package wrappers

import (
	"context"
	"time"

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
		cTTL := int64(ttl.Seconds())
		if cTTL > 0 {
			cfg.isConfigured = true
			cfg.ttl = cTTL
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
		ep.WithMinTTL(pp.cfg.ttl)
		// Additional post-processing can be added here.
	}

	return endpoints, nil
}

func (pp *postProcessor) AddEventHandler(_ context.Context, handler func()) {

}
