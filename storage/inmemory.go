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

package storage

import (
	"errors"
	"sync"

	"time"

	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/provider"
)

var (
	resyncPeriod = 1 * time.Minute
)

// InMemoryStorage implements storage interface - simple in-memory storage
// 1. Stores all information in memory -> non-persistent
// 2. To be used for testing purposes
// 3. State is not shared across different instances of external-dns
type InMemoryStorage struct {
	registry provider.Provider
	zone     string
	owner    string //refers to the owner id of the current instance
	cache    map[string]*endpoint.SharedEndpoint
	sync.Mutex
}

var _ Storage = &InMemoryStorage{}

// NewInMemoryStorage returns new InMemoryStorage object
func NewInMemoryStorage(registry provider.Provider, owner, zone string) (*InMemoryStorage, error) {
	if owner == "" || zone == "" {
		return nil, errors.New("owner and zone should not be empty strings")
	}
	return &InMemoryStorage{
		registry: registry,
		zone:     zone,
		owner:    owner,
		cache:    map[string]*endpoint.SharedEndpoint{},
	}, nil
}

// Records returns the current records from the in-memory storage
func (im *InMemoryStorage) Records() []*endpoint.SharedEndpoint {
	im.Lock()
	defer im.Unlock()
	records := make([]*endpoint.SharedEndpoint, 0, len(im.cache))
	for _, record := range im.cache {
		records = append(records, record)
	}
	return records
}

// OwnRecords returns the list of records owned by the current instance of external-dns
func (im *InMemoryStorage) OwnRecords() []endpoint.Endpoint {
	ownRecords := []endpoint.Endpoint{}
	for _, record := range im.cache {
		if record.Owner == im.owner {
			ownRecords = append(ownRecords, record.Endpoint)
		}
	}
	return ownRecords
}

// Assign updates the cache after successful registry update call
func (im *InMemoryStorage) Assign(records []endpoint.Endpoint) error {
	im.Lock()
	defer im.Unlock()
	for _, record := range records {
		if _, exist := im.cache[record.DNSName]; !exist {
			im.cache[record.DNSName] = &endpoint.SharedEndpoint{
				Endpoint: endpoint.Endpoint{
					DNSName: record.DNSName,
					Target:  record.Target,
				},
			}
		}
		im.cache[record.DNSName].Owner = im.owner
	}

	return nil
}

// WaitForSync fetches the data from dns provider to build ownership information
func (im *InMemoryStorage) WaitForSync() error {
	im.Lock()
	defer im.Unlock()

	records, err := im.registry.Records(im.zone)
	if err != nil {
		return err
	}

	curRecords := make([]*endpoint.SharedEndpoint, 0, len(im.cache))
	for _, record := range im.cache {
		curRecords = append(curRecords, record)
	}

	newCache := updatedCache(records, curRecords)

	im.cache = map[string]*endpoint.SharedEndpoint{} //drop the current cache
	for _, newCacheRecord := range newCache {
		im.cache[newCacheRecord.DNSName] = newCacheRecord
	}

	return nil
}
