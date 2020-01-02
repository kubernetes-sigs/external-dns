package provider

import (
  "context"
  "github.com/kubernetes-sigs/external-dns/endpoint"
  "github.com/kubernetes-sigs/external-dns/plan"
)

type BluecatConfig struct {

}

type BluecatProvider struct {

}

func NewBluecatProvider(bluecatConfig BluecatConfig) (*BluecatProvider, error) {
  provider := &BluecatProvider{}

  return provider, nil
}

func (p *BluecatProvider) Records() ([]*endpoint.Endpoint, error) {

}

func (p *BluecatProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {

}
