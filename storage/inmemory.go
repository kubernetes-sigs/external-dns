//drop and recreate the cache
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

	"github.com/kubernetes-incubator/external-dns/dnsprovider"
	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/prometheus/common/log"
)

var (
	resyncPeriod = 30 * time.Second
)

// InMemoryStorage implements storage interface - simple in-memory storage
// 1. Stores all information in memory -> non-persistent
// 2. To be used for testing purposes
// 3. State is not shared across different instances of external-dns
type InMemoryStorage struct {
	registry dnsprovider.DNSProvider
	zone     string
	owner    string //refers to the owner id of the current instance
	cache    map[string]*SharedEndpoint
	sync.Mutex
}

// NewInMemoryStorage returns new InMemoryStorage object
func NewInMemoryStorage(registry dnsprovider.DNSProvider, owner, zone string) (*InMemoryStorage, error) {
	if owner == "" || zone == "" {
		return nil, errors.New("owner and zone should not be empty strings")
	}
	return &InMemoryStorage{
		registry: registry,
		zone:     zone,
		owner:    owner,
		cache:    map[string]*SharedEndpoint{},
	}, nil
}

// Records returns the current records from the in-memory storage
func (im *InMemoryStorage) Records() []*SharedEndpoint {
	records := make([]*SharedEndpoint, len(im.cache))
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
	im.refreshCache() //with new rebuild cache records should already exists in the cache

	im.Lock()
	defer im.Unlock()
	for _, record := range records {
		if _, exist := im.cache[record.DNSName]; !exist {
			log.Errorf("record could not be assigned, record was not found in cache")
			continue
		}
		im.cache[record.DNSName].Owner = im.owner
	}

	return nil
}

// WaitForSync fetches the data from dns provider to build ownership information
func (im *InMemoryStorage) WaitForSync() error {
	//drop and recreate the cache
	im.refreshCache()
	return nil
}

func (im *InMemoryStorage) refreshCache() error {
	im.Lock()
	defer im.Unlock()

	records, err := im.registry.Records(im.zone)
	if err != nil {
		return err
	}

	curRecords := make([]*SharedEndpoint, len(im.cache))
	for _, record := range im.cache {
		curRecords = append(curRecords, record)
	}

	newCache := updatedCache(records, curRecords)

	im.cache = map[string]*SharedEndpoint{} //drop the current cache
	for _, newCacheRecord := range newCache {
		im.cache[newCacheRecord.DNSName] = newCacheRecord
	}

	return nil
}

// Poll periodically resyncs with the registry to update the cache
func (im *InMemoryStorage) Poll(stopChan <-chan struct{}) {
	for {
		select {
		case <-time.After(resyncPeriod):
			err := im.refreshCache()
			if err != nil {
				log.Errorf("failed to refresh cache: %v", err)
			}
		case <-stopChan:
			log.Infoln("terminating storage polling")
			return
		}
	}
}
