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

package noop

import (
	"context"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/registry"
)

// NoopRegistry implements registry interface without ownership directly propagating changes to dns provider
type NoopRegistry struct {
	provider provider.Provider
}

// New creates a NoopRegistry from the given configuration.
func New(_ *externaldns.Config, p provider.Provider) (registry.Registry, error) {
	return newRegistry(p), nil
}

// newRegistry returns new NoopRegistry object
func newRegistry(provider provider.Provider) *NoopRegistry {
	return &NoopRegistry{
		provider: provider,
	}
}

// GetDomainFilter returns the domain filter from the underlying provider.
func (im *NoopRegistry) GetDomainFilter() endpoint.DomainFilterInterface {
	return im.provider.GetDomainFilter()
}

// OwnerID returns an empty string as the noop registry does not track ownership.
func (im *NoopRegistry) OwnerID() string {
	return ""
}

// Records returns the current records from the dns provider
func (im *NoopRegistry) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	return im.provider.Records(ctx)
}

// ApplyChanges propagates changes to the dns provider
func (im *NoopRegistry) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	return im.provider.ApplyChanges(ctx, changes)
}

// AdjustEndpoints modifies the endpoints as needed by the specific provider
func (im *NoopRegistry) AdjustEndpoints(endpoints []*endpoint.Endpoint) ([]*endpoint.Endpoint, error) {
	return im.provider.AdjustEndpoints(endpoints)
}
