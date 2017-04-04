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
	"time"

	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"
	"github.com/kubernetes-incubator/external-dns/provider"
)

var (
	resyncPeriod = 1 * time.Minute
)

// InMemoryRegistry implements storage interface - in-memory registry
// 1. Stores all information in memory -> non-persistent
// 2. To be used for testing purposes
// 3. State is not shared across different instances of external-dns
type InMemoryRegistry struct {
	provider provider.Provider
	zone     string
	ownerID  string //refers to the owner id of the current instance
}

var _ Registry = &InMemoryRegistry{}

// NewInMemoryRegistry returns new InMemoryStorage object
func NewInMemoryRegistry(dnsProvider provider.Provider, ownerID, zone string) (*InMemoryRegistry, error) {
	if ownerID == "" || zone == "" {
		return nil, errors.New("owner or zone is not provided")
	}
	return &InMemoryRegistry{
		provider: dnsProvider,
		zone:     zone,
		ownerID:  ownerID,
	}, nil
}

// Records returns the current records from the in-memory storage
func (im *InMemoryRegistry) Records() ([]*endpoint.Endpoint, error) {
	_, err := im.provider.Records(im.zone)
	if err != nil {
		return nil, err
	}
	// for _, ep := range endpoints {
	//do the conversion
	// }
	return nil, nil
}

func (im *InMemoryRegistry) ApplyChanges(zone string, changes *plan.Changes) error {
	return nil
}
