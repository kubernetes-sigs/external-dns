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
	"github.com/kubernetes-incubator/external-dns/endpoint"
)

// Storage is an interface which should enable external-dns track its state
// Record() returns ALL records registered with DNS provider (TODO: consider returning only specific hosted zone)
// each entry has a field `Owner` which is equal to the identifier passed to the external-dns
// Update([]*Entry) updates the information in the storage by replacing all data by whatever is passed to Update method
// WaitForSync waits until the cache is populated with data, this is required for ConfigMap implementation
// in cases of storage implementation directly fetching freshest data from DnsProvider this interface should just return nil
type Storage interface {
	Records() []*Entry
	Update([]*Entry) error
	WaitForSync() error
}

// Entry is a unit of data stored in the storage it should provide information such as
// 1. Owner - which external-dns instance is managing the records
// 2. DNSName and Target inherited from endpoint.Endpoint struct
type Entry struct {
	Owner string //refers to the Owner ID
	endpoint.Endpoint
}
