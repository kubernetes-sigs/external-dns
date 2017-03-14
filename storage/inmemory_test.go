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
	"testing"

	"github.com/kubernetes-incubator/external-dns/dnsprovider"
)

func TestInMemory(t *testing.T) {
	t.Run("TestNewInMemoryStorage", testNewInMemoryStorage)
}

func testNewInMemoryStorage(t *testing.T) {
	var owner, zone string
	var provider dnsprovider.InMemoryProvider
	if _, err := NewInMemoryStorage(&provider, owner, zone); err == nil {
		t.Errorf("should fail when owner/zone is empty")
	}
	owner = "test-1"
	zone = ""
	if _, err := NewInMemoryStorage(&provider, owner, zone); err == nil {
		t.Errorf("should fail when owner/zone is empty")
	}
	owner = ""
	zone = "zone-1"
	if _, err := NewInMemoryStorage(&provider, owner, zone); err == nil {
		t.Errorf("should fail when owner/zone is empty")
	}
	owner = "test-1"
	zone = "zone-1"
	storage, err := NewInMemoryStorage(&provider, owner, zone)
	if err != nil {
		t.Error(err)
	}
	if storage.zone != zone || storage.owner != owner || storage.cache == nil {
		t.Errorf("incorrectly initialized in memory storage provider")
	}
	if _, ok := storage.registry.(*dnsprovider.InMemoryProvider); !ok {
		t.Errorf("incorrect dns provider is used for registry")
	}
}
