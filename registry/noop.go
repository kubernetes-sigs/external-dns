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

package registry

import (
	"context"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

// NoopRegistry implements registry interface without ownership directly propagating changes to dns provider
type NoopRegistry struct {
	provider provider.Provider
}

// NewNoopRegistry returns new NoopRegistry object
func NewNoopRegistry(provider provider.Provider) (*NoopRegistry, error) {
	return &NoopRegistry{
		provider: provider,
	}, nil
}

// Records returns the current records from the dns provider
func (im *NoopRegistry) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	return im.provider.Records(ctx)
}

// ApplyChanges propagates changes to the dns provider
func (im *NoopRegistry) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	return im.provider.ApplyChanges(ctx, changes)
}
