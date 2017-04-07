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
	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"
	"github.com/kubernetes-incubator/external-dns/provider"
	"testing"
)

const (
	testZone = "test-zone.example.com."
)

func TestTXTRegistry(t *testing.T) {
	t.Run("TestNewTXTRegistry", testTXTRegistryNew)
	t.Run("TestTXTRegistryRecords", testTXTRegistryRecords)
}

func testTXTRegistryNew(t *testing.T) {
	p := provider.NewInMemoryProvider()
	_, err := NewTXTRegistry(p, "txt", "")
	if err == nil {
		t.Fatal("owner should be specified")
	}

	r, err := NewTXTRegistry(p, "txt", "owner")
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := r.mapper.(*prefixNameMapper); !ok {
		t.Error("incorrectly initialized txt registry instance")
	}
	if r.ownerID != "owner" || r.provider != p {
		t.Error("incorrectly initialized txt registry instance")
	}

	r, err = NewTXTRegistry(p, "", "owner")
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := r.mapper.(*noopNameMapper); !ok {
		t.Error("Incorrect type of prefix name mapper")
	}
}

func testTXTRegistryRecords(t *testing.T) {
	t.Run("With prefix", testTXTRegistryRecordsPrefixed)
	t.Run("No prefix", testTXTRegistryRecordsNoop)
}

func testTXTRegistryRecordsPrefixed(t *testing.T) {
	r, _ := NewTXTRegistry(getInitializedProvider(), "txt", "owner")
	r.Records(testZone)
}

func testTXTRegistryRecordsNoop(t *testing.T) {
	p := provider.NewInMemoryProvider()
	r, err := NewTXTRegistry(p, "", "owner")
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := r.mapper.(*noopNameMapper); !ok {
		t.Error("Incorrect type of prefix name mapper")
	}
}

/**

helper methods

*/

func newEndpointWithType(dnsName, target, recordType string) *endpoint.Endpoint {
	e := endpoint.NewEndpoint(dnsName, target)
	e.RecordType = recordType
	return e
}

func getInitializedProvider() provider.Provider {
	p := provider.NewInMemoryProvider()
	p.CreateZone(testZone)
	p.ApplyChanges(testZone, &plan.Changes{
		Create: []*endpoint.Endpoint{
			newEndpointWithType("", "", ""),
			newEndpointWithType("", "", ""),
			newEndpointWithType("", "", ""),
			newEndpointWithType("", "", ""),
		},
	})
	return p
}
