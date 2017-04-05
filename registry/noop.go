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
	"errors"

	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"
	"github.com/kubernetes-incubator/external-dns/provider"
)

// NoopRegistry implements storage interface without ownership directly propagating changes to dns provider
type NoopRegistry struct {
	provider provider.Provider
	ownerID  string //refers to the owner id of the current instance
}

var _ Registry = &NoopRegistry{}

// NewNoopRegistry returns new InMemoryStorage object
func NewNoopRegistry(provider provider.Provider, ownerID string) (*NoopRegistry, error) {
	if ownerID == "" {
		return nil, errors.New("owner is not provided")
	}
	return &NoopRegistry{
		provider: provider,
		ownerID:  ownerID,
	}, nil
}

//TODO(ideahitme): Do the conversion endpoint <-> DNSRecord

// Records returns the current records from the in-memory storage
func (im *NoopRegistry) Records(zone string) ([]*endpoint.Endpoint, error) {
	eps, err := im.provider.Records(zone)
	if err != nil {
		return nil, err
	}
	return eps, err
}

// ApplyChanges updates in memory dns provider including ownership information
func (im *NoopRegistry) ApplyChanges(zone string, changes *plan.Changes) error {
	return im.provider.ApplyChanges(zone, changes)
}
