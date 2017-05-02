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
	"testing"

	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/internal/testutils"
	"github.com/kubernetes-incubator/external-dns/plan"
	"github.com/kubernetes-incubator/external-dns/provider"
)

const (
	testZone = "test-zone.example.org"
)

func TestTXTRegistry(t *testing.T) {
	t.Run("TestNewTXTRegistry", testTXTRegistryNew)
	t.Run("TestRecords", testTXTRegistryRecords)
	t.Run("TestApplyChanges", testTXTRegistryApplyChanges)
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
	if _, ok := r.mapper.(prefixNameMapper); !ok {
		t.Error("incorrectly initialized txt registry instance")
	}
	if r.ownerID != "owner" || r.provider != p {
		t.Error("incorrectly initialized txt registry instance")
	}

	r, err = NewTXTRegistry(p, "", "owner")
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := r.mapper.(prefixNameMapper); !ok {
		t.Error("Incorrect type of prefix name mapper")
	}
}

func testTXTRegistryRecords(t *testing.T) {
	t.Run("With prefix", testTXTRegistryRecordsPrefixed)
	t.Run("No prefix", testTXTRegistryRecordsNoPrefix)
}

func testTXTRegistryRecordsPrefixed(t *testing.T) {
	p := provider.NewInMemoryProvider()
	p.CreateZone(testZone)
	p.ApplyChanges(&plan.Changes{
		Create: []*endpoint.Endpoint{
			newEndpointWithOwner("foo.test-zone.example.org", "foo.loadbalancer.com", "CNAME", ""),
			newEndpointWithOwner("bar.test-zone.example.org", "my-domain.com", "CNAME", ""),
			newEndpointWithOwner("txt.bar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", "TXT", ""),
			newEndpointWithOwner("txt.bar.test-zone.example.org", "baz.test-zone.example.org", "ALIAS", ""),
			newEndpointWithOwner("qux.test-zone.example.org", "random", "TXT", ""),
			newEndpointWithOwner("tar.test-zone.example.org", "tar.loadbalancer.com", "ALIAS", ""),
			newEndpointWithOwner("txt.tar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner-2\"", "TXT", ""),
			newEndpointWithOwner("foobar.test-zone.example.org", "foobar.loadbalancer.com", "ALIAS", ""),
			newEndpointWithOwner("foobar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", "TXT", ""),
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

	r, _ := NewTXTRegistry(p, "txt.", "owner")
	records, _ := r.Records()
	if !testutils.SameEndpoints(records, expectedRecords) {
		t.Error("incorrect result returned from txt registry")
	}
}

func testTXTRegistryRecordsNoPrefix(t *testing.T) {
	p := provider.NewInMemoryProvider()
	p.CreateZone(testZone)
	p.ApplyChanges(&plan.Changes{
		Create: []*endpoint.Endpoint{
			newEndpointWithOwner("foo.test-zone.example.org", "foo.loadbalancer.com", "CNAME", ""),
			newEndpointWithOwner("bar.test-zone.example.org", "my-domain.com", "CNAME", ""),
			newEndpointWithOwner("txt.bar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", "TXT", ""),
			newEndpointWithOwner("txt.bar.test-zone.example.org", "baz.test-zone.example.org", "ALIAS", ""),
			newEndpointWithOwner("qux.test-zone.example.org", "random", "TXT", ""),
			newEndpointWithOwner("tar.test-zone.example.org", "tar.loadbalancer.com", "ALIAS", ""),
			newEndpointWithOwner("txt.tar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner-2\"", "TXT", ""),
			newEndpointWithOwner("foobar.test-zone.example.org", "foobar.loadbalancer.com", "ALIAS", ""),
			newEndpointWithOwner("foobar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", "TXT", ""),
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
				endpoint.OwnerLabelKey: "",
			},
		},
		{
			DNSName:    "txt.bar.test-zone.example.org",
			Target:     "baz.test-zone.example.org",
			RecordType: "ALIAS",
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "owner",
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
				endpoint.OwnerLabelKey: "",
			},
		},
		{
			DNSName:    "foobar.test-zone.example.org",
			Target:     "foobar.loadbalancer.com",
			RecordType: "ALIAS",
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "owner",
			},
		},
	}

	r, _ := NewTXTRegistry(p, "", "owner")
	records, _ := r.Records()

	if !testutils.SameEndpoints(records, expectedRecords) {
		t.Error("incorrect result returned from txt registry")
	}
}

func testTXTRegistryApplyChanges(t *testing.T) {
	t.Run("With Prefix", testTXTRegistryApplyChangesWithPrefix)
	t.Run("No prefix", testTXTRegistryApplyChangesNoPrefix)
}

func testTXTRegistryApplyChangesWithPrefix(t *testing.T) {
	p := provider.NewInMemoryProvider()
	p.CreateZone(testZone)
	p.ApplyChanges(&plan.Changes{
		Create: []*endpoint.Endpoint{
			newEndpointWithOwner("foo.test-zone.example.org", "foo.loadbalancer.com", "CNAME", ""),
			newEndpointWithOwner("bar.test-zone.example.org", "my-domain.com", "CNAME", ""),
			newEndpointWithOwner("txt.bar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", "TXT", ""),
			newEndpointWithOwner("txt.bar.test-zone.example.org", "baz.test-zone.example.org", "ALIAS", ""),
			newEndpointWithOwner("qux.test-zone.example.org", "random", "TXT", ""),
			newEndpointWithOwner("tar.test-zone.example.org", "tar.loadbalancer.com", "ALIAS", ""),
			newEndpointWithOwner("txt.tar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", "TXT", ""),
			newEndpointWithOwner("foobar.test-zone.example.org", "foobar.loadbalancer.com", "ALIAS", ""),
			newEndpointWithOwner("txt.foobar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", "TXT", ""),
		},
	})
	r, _ := NewTXTRegistry(p, "txt.", "owner")

	changes := &plan.Changes{
		Create: []*endpoint.Endpoint{
			newEndpointWithOwner("new-record-1.test-zone.example.org", "new-loadbalancer-1.lb.com", "", ""),
		},
		Delete: []*endpoint.Endpoint{
			newEndpointWithOwner("foobar.test-zone.example.org", "foobar.loadbalancer.com", "ALIAS", "owner"),
		},
		UpdateNew: []*endpoint.Endpoint{
			newEndpointWithOwner("tar.test-zone.example.org", "new-tar.loadbalancer.com", "ALIAS", "owner"),
		},
		UpdateOld: []*endpoint.Endpoint{
			newEndpointWithOwner("tar.test-zone.example.org", "tar.loadbalancer.com", "ALIAS", "owner"),
		},
	}
	expected := &plan.Changes{
		Create: []*endpoint.Endpoint{
			newEndpointWithOwner("new-record-1.test-zone.example.org", "new-loadbalancer-1.lb.com", "", ""),
			newEndpointWithOwner("txt.new-record-1.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", "TXT", ""),
		},
		Delete: []*endpoint.Endpoint{
			newEndpointWithOwner("foobar.test-zone.example.org", "foobar.loadbalancer.com", "ALIAS", "owner"),
			newEndpointWithOwner("txt.foobar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", "TXT", ""),
		},
		UpdateNew: []*endpoint.Endpoint{
			newEndpointWithOwner("tar.test-zone.example.org", "new-tar.loadbalancer.com", "ALIAS", "owner"),
		},
		UpdateOld: []*endpoint.Endpoint{
			newEndpointWithOwner("tar.test-zone.example.org", "tar.loadbalancer.com", "ALIAS", "owner"),
		},
	}
	p.OnApplyChanges = func(got *plan.Changes) {
		mExpected := map[string][]*endpoint.Endpoint{
			"Create":    expected.Create,
			"UpdateNew": expected.UpdateNew,
			"UpdateOld": expected.UpdateOld,
			"Delete":    expected.Delete,
		}
		mGot := map[string][]*endpoint.Endpoint{
			"Create":    got.Create,
			"UpdateNew": got.UpdateNew,
			"UpdateOld": got.UpdateOld,
			"Delete":    got.Delete,
		}
		if !testutils.SamePlanChanges(mGot, mExpected) {
			t.Error("incorrect plan changes are passed to provider")
		}
	}
	err := r.ApplyChanges(changes)
	if err != nil {
		t.Fatal(err)
	}
}

func testTXTRegistryApplyChangesNoPrefix(t *testing.T) {
	p := provider.NewInMemoryProvider()
	p.CreateZone(testZone)
	p.ApplyChanges(&plan.Changes{
		Create: []*endpoint.Endpoint{
			newEndpointWithOwner("foo.test-zone.example.org", "foo.loadbalancer.com", "CNAME", ""),
			newEndpointWithOwner("bar.test-zone.example.org", "my-domain.com", "CNAME", ""),
			newEndpointWithOwner("txt.bar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", "TXT", ""),
			newEndpointWithOwner("txt.bar.test-zone.example.org", "baz.test-zone.example.org", "ALIAS", ""),
			newEndpointWithOwner("qux.test-zone.example.org", "random", "TXT", ""),
			newEndpointWithOwner("tar.test-zone.example.org", "tar.loadbalancer.com", "ALIAS", ""),
			newEndpointWithOwner("txt.tar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", "TXT", ""),
			newEndpointWithOwner("foobar.test-zone.example.org", "foobar.loadbalancer.com", "ALIAS", ""),
			newEndpointWithOwner("foobar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", "TXT", ""),
		},
	})
	r, _ := NewTXTRegistry(p, "", "owner")

	changes := &plan.Changes{
		Create: []*endpoint.Endpoint{
			newEndpointWithOwner("new-record-1.test-zone.example.org", "new-loadbalancer-1.lb.com", "", ""),
		},
		Delete: []*endpoint.Endpoint{
			newEndpointWithOwner("foobar.test-zone.example.org", "foobar.loadbalancer.com", "ALIAS", "owner"),
		},
		UpdateNew: []*endpoint.Endpoint{
			newEndpointWithOwner("tar.test-zone.example.org", "new-tar.loadbalancer.com", "ALIAS", "owner-2"),
		},
		UpdateOld: []*endpoint.Endpoint{
			newEndpointWithOwner("tar.test-zone.example.org", "tar.loadbalancer.com", "ALIAS", "owner-2"),
		},
	}
	expected := &plan.Changes{
		Create: []*endpoint.Endpoint{
			newEndpointWithOwner("new-record-1.test-zone.example.org", "new-loadbalancer-1.lb.com", "", ""),
			newEndpointWithOwner("new-record-1.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", "TXT", ""),
		},
		Delete: []*endpoint.Endpoint{
			newEndpointWithOwner("foobar.test-zone.example.org", "foobar.loadbalancer.com", "ALIAS", "owner"),
			newEndpointWithOwner("foobar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", "TXT", ""),
		},
		UpdateNew: []*endpoint.Endpoint{},
		UpdateOld: []*endpoint.Endpoint{},
	}
	p.OnApplyChanges = func(got *plan.Changes) {
		mExpected := map[string][]*endpoint.Endpoint{
			"Create":    expected.Create,
			"UpdateNew": expected.UpdateNew,
			"UpdateOld": expected.UpdateOld,
			"Delete":    expected.Delete,
		}
		mGot := map[string][]*endpoint.Endpoint{
			"Create":    got.Create,
			"UpdateNew": got.UpdateNew,
			"UpdateOld": got.UpdateOld,
			"Delete":    got.Delete,
		}
		if !testutils.SamePlanChanges(mGot, mExpected) {
			t.Error("incorrect plan changes are passed to provider")
		}
	}
	err := r.ApplyChanges(changes)
	if err != nil {
		t.Fatal(err)
	}
}

/**

helper methods

*/

func newEndpointWithOwner(dnsName, target, recordType, ownerID string) *endpoint.Endpoint {
	e := endpoint.NewEndpoint(dnsName, target, recordType)
	e.Labels[endpoint.OwnerLabelKey] = ownerID
	return e
}
