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

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	require.Error(t, err)

	r, err := NewTXTRegistry(p, "txt", "owner")
	require.NoError(t, err)

	_, ok := r.mapper.(prefixNameMapper)
	require.True(t, ok)
	assert.Equal(t, "owner", r.ownerID)
	assert.Equal(t, p, r.provider)

	r, err = NewTXTRegistry(p, "", "owner")
	require.NoError(t, err)

	_, ok = r.mapper.(prefixNameMapper)
	assert.True(t, ok)
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
			newEndpointWithOwner("foo.test-zone.example.org", "foo.loadbalancer.com", endpoint.RecordTypeCNAME, ""),
			newEndpointWithOwner("bar.test-zone.example.org", "my-domain.com", endpoint.RecordTypeCNAME, ""),
			newEndpointWithOwner("txt.bar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("txt.bar.test-zone.example.org", "baz.test-zone.example.org", endpoint.RecordTypeCNAME, ""),
			newEndpointWithOwner("qux.test-zone.example.org", "random", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("tar.test-zone.example.org", "tar.loadbalancer.com", endpoint.RecordTypeCNAME, ""),
			newEndpointWithOwner("txt.tar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner-2\"", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("foobar.test-zone.example.org", "foobar.loadbalancer.com", endpoint.RecordTypeCNAME, ""),
			newEndpointWithOwner("foobar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
		},
	})
	expectedRecords := []*endpoint.Endpoint{
		{
			DNSName:    "foo.test-zone.example.org",
			Target:     "foo.loadbalancer.com",
			RecordType: endpoint.RecordTypeCNAME,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "",
			},
		},
		{
			DNSName:    "bar.test-zone.example.org",
			Target:     "my-domain.com",
			RecordType: endpoint.RecordTypeCNAME,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "owner",
			},
		},
		{
			DNSName:    "txt.bar.test-zone.example.org",
			Target:     "baz.test-zone.example.org",
			RecordType: endpoint.RecordTypeCNAME,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "",
			},
		},
		{
			DNSName:    "qux.test-zone.example.org",
			Target:     "random",
			RecordType: endpoint.RecordTypeTXT,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "",
			},
		},
		{
			DNSName:    "tar.test-zone.example.org",
			Target:     "tar.loadbalancer.com",
			RecordType: endpoint.RecordTypeCNAME,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "owner-2",
			},
		},
		{
			DNSName:    "foobar.test-zone.example.org",
			Target:     "foobar.loadbalancer.com",
			RecordType: endpoint.RecordTypeCNAME,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "",
			},
		},
	}

	r, _ := NewTXTRegistry(p, "txt.", "owner")
	records, _ := r.Records()
	assert.True(t, testutils.SameEndpoints(records, expectedRecords))
}

func testTXTRegistryRecordsNoPrefix(t *testing.T) {
	p := provider.NewInMemoryProvider()
	p.CreateZone(testZone)
	p.ApplyChanges(&plan.Changes{
		Create: []*endpoint.Endpoint{
			newEndpointWithOwner("foo.test-zone.example.org", "foo.loadbalancer.com", endpoint.RecordTypeCNAME, ""),
			newEndpointWithOwner("bar.test-zone.example.org", "my-domain.com", endpoint.RecordTypeCNAME, ""),
			newEndpointWithOwner("txt.bar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("txt.bar.test-zone.example.org", "baz.test-zone.example.org", endpoint.RecordTypeCNAME, ""),
			newEndpointWithOwner("qux.test-zone.example.org", "random", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("tar.test-zone.example.org", "tar.loadbalancer.com", endpoint.RecordTypeCNAME, ""),
			newEndpointWithOwner("txt.tar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner-2\"", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("foobar.test-zone.example.org", "foobar.loadbalancer.com", endpoint.RecordTypeCNAME, ""),
			newEndpointWithOwner("foobar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
		},
	})
	expectedRecords := []*endpoint.Endpoint{
		{
			DNSName:    "foo.test-zone.example.org",
			Target:     "foo.loadbalancer.com",
			RecordType: endpoint.RecordTypeCNAME,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "",
			},
		},
		{
			DNSName:    "bar.test-zone.example.org",
			Target:     "my-domain.com",
			RecordType: endpoint.RecordTypeCNAME,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "",
			},
		},
		{
			DNSName:    "txt.bar.test-zone.example.org",
			Target:     "baz.test-zone.example.org",
			RecordType: endpoint.RecordTypeCNAME,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "owner",
			},
		},
		{
			DNSName:    "qux.test-zone.example.org",
			Target:     "random",
			RecordType: endpoint.RecordTypeTXT,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "",
			},
		},
		{
			DNSName:    "tar.test-zone.example.org",
			Target:     "tar.loadbalancer.com",
			RecordType: endpoint.RecordTypeCNAME,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "",
			},
		},
		{
			DNSName:    "foobar.test-zone.example.org",
			Target:     "foobar.loadbalancer.com",
			RecordType: endpoint.RecordTypeCNAME,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "owner",
			},
		},
	}

	r, _ := NewTXTRegistry(p, "", "owner")
	records, _ := r.Records()

	assert.True(t, testutils.SameEndpoints(records, expectedRecords))
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
			newEndpointWithOwner("foo.test-zone.example.org", "foo.loadbalancer.com", endpoint.RecordTypeCNAME, ""),
			newEndpointWithOwner("bar.test-zone.example.org", "my-domain.com", endpoint.RecordTypeCNAME, ""),
			newEndpointWithOwner("txt.bar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("txt.bar.test-zone.example.org", "baz.test-zone.example.org", endpoint.RecordTypeCNAME, ""),
			newEndpointWithOwner("qux.test-zone.example.org", "random", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("tar.test-zone.example.org", "tar.loadbalancer.com", endpoint.RecordTypeCNAME, ""),
			newEndpointWithOwner("txt.tar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("foobar.test-zone.example.org", "foobar.loadbalancer.com", endpoint.RecordTypeCNAME, ""),
			newEndpointWithOwner("txt.foobar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
		},
	})
	r, _ := NewTXTRegistry(p, "txt.", "owner")

	changes := &plan.Changes{
		Create: []*endpoint.Endpoint{
			newEndpointWithOwner("new-record-1.test-zone.example.org", "new-loadbalancer-1.lb.com", "", ""),
		},
		Delete: []*endpoint.Endpoint{
			newEndpointWithOwner("foobar.test-zone.example.org", "foobar.loadbalancer.com", endpoint.RecordTypeCNAME, "owner"),
		},
		UpdateNew: []*endpoint.Endpoint{
			newEndpointWithOwner("tar.test-zone.example.org", "new-tar.loadbalancer.com", endpoint.RecordTypeCNAME, "owner"),
		},
		UpdateOld: []*endpoint.Endpoint{
			newEndpointWithOwner("tar.test-zone.example.org", "tar.loadbalancer.com", endpoint.RecordTypeCNAME, "owner"),
		},
	}
	expected := &plan.Changes{
		Create: []*endpoint.Endpoint{
			newEndpointWithOwner("new-record-1.test-zone.example.org", "new-loadbalancer-1.lb.com", "", ""),
			newEndpointWithOwner("txt.new-record-1.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
		},
		Delete: []*endpoint.Endpoint{
			newEndpointWithOwner("foobar.test-zone.example.org", "foobar.loadbalancer.com", endpoint.RecordTypeCNAME, "owner"),
			newEndpointWithOwner("txt.foobar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
		},
		UpdateNew: []*endpoint.Endpoint{
			newEndpointWithOwner("tar.test-zone.example.org", "new-tar.loadbalancer.com", endpoint.RecordTypeCNAME, "owner"),
		},
		UpdateOld: []*endpoint.Endpoint{
			newEndpointWithOwner("tar.test-zone.example.org", "tar.loadbalancer.com", endpoint.RecordTypeCNAME, "owner"),
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
		assert.True(t, testutils.SamePlanChanges(mGot, mExpected))
	}
	err := r.ApplyChanges(changes)
	require.NoError(t, err)
}

func testTXTRegistryApplyChangesNoPrefix(t *testing.T) {
	p := provider.NewInMemoryProvider()
	p.CreateZone(testZone)
	p.ApplyChanges(&plan.Changes{
		Create: []*endpoint.Endpoint{
			newEndpointWithOwner("foo.test-zone.example.org", "foo.loadbalancer.com", endpoint.RecordTypeCNAME, ""),
			newEndpointWithOwner("bar.test-zone.example.org", "my-domain.com", endpoint.RecordTypeCNAME, ""),
			newEndpointWithOwner("txt.bar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("txt.bar.test-zone.example.org", "baz.test-zone.example.org", endpoint.RecordTypeCNAME, ""),
			newEndpointWithOwner("qux.test-zone.example.org", "random", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("tar.test-zone.example.org", "tar.loadbalancer.com", endpoint.RecordTypeCNAME, ""),
			newEndpointWithOwner("txt.tar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("foobar.test-zone.example.org", "foobar.loadbalancer.com", endpoint.RecordTypeCNAME, ""),
			newEndpointWithOwner("foobar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
		},
	})
	r, _ := NewTXTRegistry(p, "", "owner")

	changes := &plan.Changes{
		Create: []*endpoint.Endpoint{
			newEndpointWithOwner("new-record-1.test-zone.example.org", "new-loadbalancer-1.lb.com", "", ""),
		},
		Delete: []*endpoint.Endpoint{
			newEndpointWithOwner("foobar.test-zone.example.org", "foobar.loadbalancer.com", endpoint.RecordTypeCNAME, "owner"),
		},
		UpdateNew: []*endpoint.Endpoint{
			newEndpointWithOwner("tar.test-zone.example.org", "new-tar.loadbalancer.com", endpoint.RecordTypeCNAME, "owner-2"),
		},
		UpdateOld: []*endpoint.Endpoint{
			newEndpointWithOwner("tar.test-zone.example.org", "tar.loadbalancer.com", endpoint.RecordTypeCNAME, "owner-2"),
		},
	}
	expected := &plan.Changes{
		Create: []*endpoint.Endpoint{
			newEndpointWithOwner("new-record-1.test-zone.example.org", "new-loadbalancer-1.lb.com", "", ""),
			newEndpointWithOwner("new-record-1.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
		},
		Delete: []*endpoint.Endpoint{
			newEndpointWithOwner("foobar.test-zone.example.org", "foobar.loadbalancer.com", endpoint.RecordTypeCNAME, "owner"),
			newEndpointWithOwner("foobar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
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
		assert.True(t, testutils.SamePlanChanges(mGot, mExpected))
	}
	err := r.ApplyChanges(changes)
	require.NoError(t, err)
}

/**

helper methods

*/

func newEndpointWithOwner(dnsName, target, recordType, ownerID string) *endpoint.Endpoint {
	e := endpoint.NewEndpoint(dnsName, target, recordType)
	e.Labels[endpoint.OwnerLabelKey] = ownerID
	return e
}
