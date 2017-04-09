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
	"github.com/kubernetes-incubator/external-dns/internal/testutils"
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

	rs, err := r.Records("random-zone")
	if err == nil || rs != nil {
		t.Error("incorrect zone should trigger error")
	}
}

func testTXTRegistryRecords(t *testing.T) {
	t.Run("With prefix", testTXTRegistryRecordsPrefixed)
	t.Run("No prefix", testTXTRegistryRecordsNoop)
}

func testTXTRegistryRecordsPrefixed(t *testing.T) {
	p := provider.NewInMemoryProvider()
	p.CreateZone(testZone)
	p.ApplyChanges(testZone, &plan.Changes{
		Create: []*endpoint.Endpoint{
			newEndpointWithType("foo.test-zone.example.org", "foo.loadbalancer.com", "CNAME"),
			newEndpointWithType("bar.test-zone.example.org", "my-domain.com", "CNAME"),
			newEndpointWithType("txt.bar.test-zone.example.org", "heritage=external-dns;record-owner-id=owner", "TXT"),
			newEndpointWithType("txt.bar.test-zone.example.org", "baz.test-zone.example.org", "ALIAS"),
			newEndpointWithType("qux.test-zone.example.org", "random", "TXT"),
			newEndpointWithType("tar.test-zone.example.org", "tar.loadbalancer.com", "ALIAS"),
			newEndpointWithType("txt.tar.test-zone.example.org", "heritage=external-dns;record-owner-id=owner-2", "TXT"),
			newEndpointWithType("foobar.test-zone.example.org", "foobar.loadbalancer.com", "ALIAS"),
			newEndpointWithType("foobar.test-zone.example.org", "heritage=external-dns;record-owner-id=owner", "TXT"),
		},
	})
	expectedRecords := []*endpoint.Endpoint{
		{
			DNSName:    "foo.test-zone.example.org",
			Target:     "foo.loadbalancer.com",
			RecordType: "CNAME",
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "",
			},
		},
		{
			DNSName:    "bar.test-zone.example.org",
			Target:     "my-domain.com",
			RecordType: "CNAME",
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "owner",
			},
		},
		{
			DNSName:    "txt.bar.test-zone.example.org",
			Target:     "baz.test-zone.example.org",
			RecordType: "ALIAS",
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "",
			},
		},
		{
			DNSName:    "qux.test-zone.example.org",
			Target:     "random",
			RecordType: "TXT",
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "",
			},
		},
		{
			DNSName:    "tar.test-zone.example.org",
			Target:     "tar.loadbalancer.com",
			RecordType: "ALIAS",
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "owner-2",
			},
		},
		{
			DNSName:    "foobar.test-zone.example.org",
			Target:     "foobar.loadbalancer.com",
			RecordType: "ALIAS",
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "",
			},
		},
	}

	r, _ := NewTXTRegistry(p, "txt", "owner")
	records, _ := r.Records(testZone)
	if !testutils.SameEndpoints(records, expectedRecords) {
		t.Error("incorrect result returned from txt registry")
	}
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
