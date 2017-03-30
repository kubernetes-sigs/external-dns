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

import "github.com/kubernetes-incubator/external-dns/endpoint"

// Storage is an interface which should enable external-dns track its state
// Record() returns ALL records registered with DNS provider (TODO: consider returning only specific hosted zone)
// each entry has a field `Owner` which is equal to the identifier passed to the external-dns which created the record
// OwnRecords() returns list of records which are owned by current external-dns instance
// Assign([]*endpoint.Endpoint) assigns the owner to the provided list of endpoints and updates the storage
// WaitForSync() waits until the cache is populated with data, this should be called once to make sure that the storage is usable
// Poll(stopChan <- chan struct{]}) periodically resyncs and updates the cache from dnsprovider
type Storage interface {
	Records() []*endpoint.SharedEndpoint
	OwnRecords() []endpoint.Endpoint
	Assign([]endpoint.Endpoint) error
	Poll(stopChan <-chan struct{})
	WaitForSync() error
}

// updatedCache storage agnostic functionality to merge the existing cache records - including owner information
// with the freshest dnsprovider registered records
// make sure to include lock/unlock wrapper for the function call
func updatedCache(records []endpoint.Endpoint, cacheRecords []*endpoint.SharedEndpoint) []*endpoint.SharedEndpoint {
	newCache := make([]*endpoint.SharedEndpoint, len(records))
	for i, record := range records {
		newCache[i] = &endpoint.SharedEndpoint{Endpoint: record}
		for _, cache := range cacheRecords {
			if cache.DNSName == record.DNSName {
				newCache[i].Owner = cache.Owner
				break
			}
		}
	}
	return newCache
}
