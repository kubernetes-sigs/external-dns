package yandex

import (
	"context"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

type YandexConfig struct {
	DomainFilter   endpoint.DomainFilter
	ZoneNameFilter endpoint.DomainFilter
	DryRun         bool
}

type YandexProvider struct {
	provider.BaseProvider
}

func NewYandexProvider(cfg *YandexConfig) (*YandexProvider, error) {
	return &YandexProvider{}, nil
}

func (y *YandexProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	return nil, nil
}

func (y *YandexProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	return nil
}
