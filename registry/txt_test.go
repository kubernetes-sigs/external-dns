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
	"context"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/provider/inmemory"
)

const (
	testZone = "test-zone.example.org"
)

func TestTXTRegistry(t *testing.T) {
	t.Run("TestNewTXTRegistry", testTXTRegistryNew)
	t.Run("TestRecords", testTXTRegistryRecords)
	t.Run("TestApplyChanges", testTXTRegistryApplyChanges)
	t.Run("TestMissingRecords", testTXTRegistryMissingRecords)
}

func testTXTRegistryNew(t *testing.T) {
	p := inmemory.NewInMemoryProvider()
	_, err := NewTXTRegistry(p, "txt", "", "", time.Hour, "", []string{}, []string{}, false, nil)
	require.Error(t, err)

	_, err = NewTXTRegistry(p, "", "txt", "", time.Hour, "", []string{}, []string{}, false, nil)
	require.Error(t, err)

	r, err := NewTXTRegistry(p, "txt", "", "owner", time.Hour, "", []string{}, []string{}, false, nil)
	require.NoError(t, err)
	assert.Equal(t, p, r.provider)

	r, err = NewTXTRegistry(p, "", "txt", "owner", time.Hour, "", []string{}, []string{}, false, nil)
	require.NoError(t, err)

	_, err = NewTXTRegistry(p, "txt", "txt", "owner", time.Hour, "", []string{}, []string{}, false, nil)
	require.Error(t, err)

	_, ok := r.mapper.(affixNameMapper)
	require.True(t, ok)
	assert.Equal(t, "owner", r.ownerID)
	assert.Equal(t, p, r.provider)

	aesKey := []byte(";k&l)nUC/33:{?d{3)54+,AD?]SX%yh^")
	_, err = NewTXTRegistry(p, "", "", "owner", time.Hour, "", []string{}, []string{}, false, nil)
	require.NoError(t, err)

	_, err = NewTXTRegistry(p, "", "", "owner", time.Hour, "", []string{}, []string{}, false, aesKey)
	require.NoError(t, err)

	_, err = NewTXTRegistry(p, "", "", "owner", time.Hour, "", []string{}, []string{}, true, nil)
	require.Error(t, err)

	r, err = NewTXTRegistry(p, "", "", "owner", time.Hour, "", []string{}, []string{}, true, aesKey)
	require.NoError(t, err)

	_, ok = r.mapper.(affixNameMapper)
	assert.True(t, ok)
}

func testTXTRegistryRecords(t *testing.T) {
	t.Run("With prefix", testTXTRegistryRecordsPrefixed)
	t.Run("With suffix", testTXTRegistryRecordsSuffixed)
	t.Run("No prefix", testTXTRegistryRecordsNoPrefix)
	t.Run("With templated prefix", testTXTRegistryRecordsPrefixedTemplated)
	t.Run("With templated suffix", testTXTRegistryRecordsSuffixedTemplated)
}

func testTXTRegistryRecordsPrefixed(t *testing.T) {
	ctx := context.Background()
	p := inmemory.NewInMemoryProvider()
	p.CreateZone(testZone)
	p.ApplyChanges(ctx, &plan.Changes{
		Create: []*endpoint.Endpoint{
			newEndpointWithOwnerAndLabels("foo.test-zone.example.org", "foo.loadbalancer.com", endpoint.RecordTypeCNAME, "", endpoint.Labels{"foo": "somefoo"}),
			newEndpointWithOwnerAndLabels("bar.test-zone.example.org", "my-domain.com", endpoint.RecordTypeCNAME, "", endpoint.Labels{"bar": "somebar"}),
			newEndpointWithOwner("txt.bar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("txt.bar.test-zone.example.org", "baz.test-zone.example.org", endpoint.RecordTypeCNAME, ""),
			newEndpointWithOwner("qux.test-zone.example.org", "random", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwnerAndLabels("tar.test-zone.example.org", "tar.loadbalancer.com", endpoint.RecordTypeCNAME, "", endpoint.Labels{"tar": "sometar"}),
			newEndpointWithOwner("TxT.tar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner-2\"", endpoint.RecordTypeTXT, ""), // case-insensitive TXT prefix
			newEndpointWithOwner("foobar.test-zone.example.org", "foobar.loadbalancer.com", endpoint.RecordTypeCNAME, ""),
			newEndpointWithOwner("foobar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("multiple.test-zone.example.org", "lb1.loadbalancer.com", endpoint.RecordTypeCNAME, "").WithSetIdentifier("test-set-1"),
			newEndpointWithOwner("multiple.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, "").WithSetIdentifier("test-set-1"),
			newEndpointWithOwner("multiple.test-zone.example.org", "lb2.loadbalancer.com", endpoint.RecordTypeCNAME, "").WithSetIdentifier("test-set-2"),
			newEndpointWithOwner("multiple.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, "").WithSetIdentifier("test-set-2"),
			newEndpointWithOwner("*.wildcard.test-zone.example.org", "foo.loadbalancer.com", endpoint.RecordTypeCNAME, ""),
			newEndpointWithOwner("txt.wc.wildcard.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("dualstack.test-zone.example.org", "1.1.1.1", endpoint.RecordTypeA, ""),
			newEndpointWithOwner("txt.dualstack.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("dualstack.test-zone.example.org", "2001:DB8::1", endpoint.RecordTypeAAAA, ""),
			newEndpointWithOwner("txt.aaaa-dualstack.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner-2\"", endpoint.RecordTypeTXT, ""),
		},
	})
	expectedRecords := []*endpoint.Endpoint{
		{
			DNSName:    "foo.test-zone.example.org",
			Targets:    endpoint.Targets{"foo.loadbalancer.com"},
			RecordType: endpoint.RecordTypeCNAME,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "",
				"foo":                  "somefoo",
			},
		},
		{
			DNSName:    "bar.test-zone.example.org",
			Targets:    endpoint.Targets{"my-domain.com"},
			RecordType: endpoint.RecordTypeCNAME,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "owner",
				"bar":                  "somebar",
			},
		},
		{
			DNSName:    "txt.bar.test-zone.example.org",
			Targets:    endpoint.Targets{"baz.test-zone.example.org"},
			RecordType: endpoint.RecordTypeCNAME,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "",
			},
		},
		{
			DNSName:    "qux.test-zone.example.org",
			Targets:    endpoint.Targets{"random"},
			RecordType: endpoint.RecordTypeTXT,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "",
			},
		},
		{
			DNSName:    "tar.test-zone.example.org",
			Targets:    endpoint.Targets{"tar.loadbalancer.com"},
			RecordType: endpoint.RecordTypeCNAME,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "owner-2",
				"tar":                  "sometar",
			},
		},
		{
			DNSName:    "foobar.test-zone.example.org",
			Targets:    endpoint.Targets{"foobar.loadbalancer.com"},
			RecordType: endpoint.RecordTypeCNAME,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "",
			},
		},
		{
			DNSName:       "multiple.test-zone.example.org",
			Targets:       endpoint.Targets{"lb1.loadbalancer.com"},
			SetIdentifier: "test-set-1",
			RecordType:    endpoint.RecordTypeCNAME,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "",
			},
		},
		{
			DNSName:       "multiple.test-zone.example.org",
			Targets:       endpoint.Targets{"lb2.loadbalancer.com"},
			SetIdentifier: "test-set-2",
			RecordType:    endpoint.RecordTypeCNAME,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "",
			},
		},
		{
			DNSName:    "*.wildcard.test-zone.example.org",
			Targets:    endpoint.Targets{"foo.loadbalancer.com"},
			RecordType: endpoint.RecordTypeCNAME,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "owner",
			},
		},
		{
			DNSName:    "dualstack.test-zone.example.org",
			Targets:    endpoint.Targets{"1.1.1.1"},
			RecordType: endpoint.RecordTypeA,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "owner",
			},
		},
		{
			DNSName:    "dualstack.test-zone.example.org",
			Targets:    endpoint.Targets{"2001:DB8::1"},
			RecordType: endpoint.RecordTypeAAAA,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "owner-2",
			},
		},
	}

	r, _ := NewTXTRegistry(p, "txt.", "", "owner", time.Hour, "wc", []string{}, []string{}, false, nil)
	records, _ := r.Records(ctx)

	assert.True(t, testutils.SameEndpoints(records, expectedRecords))

	// Ensure prefix is case-insensitive
	r, _ = NewTXTRegistry(p, "TxT.", "", "owner", time.Hour, "wc", []string{}, []string{}, false, nil)
	records, _ = r.Records(ctx)

	assert.True(t, testutils.SameEndpoints(records, expectedRecords))
}

func testTXTRegistryRecordsSuffixed(t *testing.T) {
	ctx := context.Background()
	p := inmemory.NewInMemoryProvider()
	p.CreateZone(testZone)
	p.ApplyChanges(ctx, &plan.Changes{
		Create: []*endpoint.Endpoint{
			newEndpointWithOwnerAndLabels("foo.test-zone.example.org", "foo.loadbalancer.com", endpoint.RecordTypeCNAME, "", endpoint.Labels{"foo": "somefoo"}),
			newEndpointWithOwnerAndLabels("bar.test-zone.example.org", "my-domain.com", endpoint.RecordTypeCNAME, "", endpoint.Labels{"bar": "somebar"}),
			newEndpointWithOwner("bar-txt.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("bar-txt.test-zone.example.org", "baz.test-zone.example.org", endpoint.RecordTypeCNAME, ""),
			newEndpointWithOwner("qux.test-zone.example.org", "random", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwnerAndLabels("tar.test-zone.example.org", "tar.loadbalancer.com", endpoint.RecordTypeCNAME, "", endpoint.Labels{"tar": "sometar"}),
			newEndpointWithOwner("tar-TxT.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner-2\"", endpoint.RecordTypeTXT, ""), // case-insensitive TXT prefix
			newEndpointWithOwner("foobar.test-zone.example.org", "foobar.loadbalancer.com", endpoint.RecordTypeCNAME, ""),
			newEndpointWithOwner("foobar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("multiple.test-zone.example.org", "lb1.loadbalancer.com", endpoint.RecordTypeCNAME, "").WithSetIdentifier("test-set-1"),
			newEndpointWithOwner("multiple.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, "").WithSetIdentifier("test-set-1"),
			newEndpointWithOwner("multiple.test-zone.example.org", "lb2.loadbalancer.com", endpoint.RecordTypeCNAME, "").WithSetIdentifier("test-set-2"),
			newEndpointWithOwner("multiple.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, "").WithSetIdentifier("test-set-2"),
			newEndpointWithOwner("dualstack.test-zone.example.org", "1.1.1.1", endpoint.RecordTypeA, ""),
			newEndpointWithOwner("dualstack-txt.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("dualstack.test-zone.example.org", "2001:DB8::1", endpoint.RecordTypeAAAA, ""),
			newEndpointWithOwner("aaaa-dualstack-txt.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner-2\"", endpoint.RecordTypeTXT, ""),
		},
	})
	expectedRecords := []*endpoint.Endpoint{
		{
			DNSName:    "foo.test-zone.example.org",
			Targets:    endpoint.Targets{"foo.loadbalancer.com"},
			RecordType: endpoint.RecordTypeCNAME,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "",
				"foo":                  "somefoo",
			},
		},
		{
			DNSName:    "bar.test-zone.example.org",
			Targets:    endpoint.Targets{"my-domain.com"},
			RecordType: endpoint.RecordTypeCNAME,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "owner",
				"bar":                  "somebar",
			},
		},
		{
			DNSName:    "bar-txt.test-zone.example.org",
			Targets:    endpoint.Targets{"baz.test-zone.example.org"},
			RecordType: endpoint.RecordTypeCNAME,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "",
			},
		},
		{
			DNSName:    "qux.test-zone.example.org",
			Targets:    endpoint.Targets{"random"},
			RecordType: endpoint.RecordTypeTXT,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "",
			},
		},
		{
			DNSName:    "tar.test-zone.example.org",
			Targets:    endpoint.Targets{"tar.loadbalancer.com"},
			RecordType: endpoint.RecordTypeCNAME,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "owner-2",
				"tar":                  "sometar",
			},
		},
		{
			DNSName:    "foobar.test-zone.example.org",
			Targets:    endpoint.Targets{"foobar.loadbalancer.com"},
			RecordType: endpoint.RecordTypeCNAME,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "",
			},
		},
		{
			DNSName:       "multiple.test-zone.example.org",
			Targets:       endpoint.Targets{"lb1.loadbalancer.com"},
			SetIdentifier: "test-set-1",
			RecordType:    endpoint.RecordTypeCNAME,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "",
			},
		},
		{
			DNSName:       "multiple.test-zone.example.org",
			Targets:       endpoint.Targets{"lb2.loadbalancer.com"},
			SetIdentifier: "test-set-2",
			RecordType:    endpoint.RecordTypeCNAME,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "",
			},
		},
		{
			DNSName:    "dualstack.test-zone.example.org",
			Targets:    endpoint.Targets{"1.1.1.1"},
			RecordType: endpoint.RecordTypeA,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "owner",
			},
		},
		{
			DNSName:    "dualstack.test-zone.example.org",
			Targets:    endpoint.Targets{"2001:DB8::1"},
			RecordType: endpoint.RecordTypeAAAA,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "owner-2",
			},
		},
	}

	r, _ := NewTXTRegistry(p, "", "-txt", "owner", time.Hour, "", []string{}, []string{}, false, nil)
	records, _ := r.Records(ctx)

	assert.True(t, testutils.SameEndpoints(records, expectedRecords))

	// Ensure prefix is case-insensitive
	r, _ = NewTXTRegistry(p, "", "-TxT", "owner", time.Hour, "", []string{}, []string{}, false, nil)
	records, _ = r.Records(ctx)

	assert.True(t, testutils.SameEndpointLabels(records, expectedRecords))
}

func testTXTRegistryRecordsNoPrefix(t *testing.T) {
	p := inmemory.NewInMemoryProvider()
	ctx := context.Background()
	p.CreateZone(testZone)
	p.ApplyChanges(ctx, &plan.Changes{
		Create: []*endpoint.Endpoint{
			newEndpointWithOwner("foo.test-zone.example.org", "foo.loadbalancer.com", endpoint.RecordTypeCNAME, ""),
			newEndpointWithOwner("bar.test-zone.example.org", "my-domain.com", endpoint.RecordTypeCNAME, ""),
			newEndpointWithOwner("alias.test-zone.example.org", "my-domain.com", endpoint.RecordTypeA, "").WithProviderSpecific("alias", "true"),
			newEndpointWithOwner("cname-alias.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("txt.bar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner,external-dns/resource=ingress/default/my-ingress\"", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("txt.bar.test-zone.example.org", "baz.test-zone.example.org", endpoint.RecordTypeCNAME, ""),
			newEndpointWithOwner("qux.test-zone.example.org", "random", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("tar.test-zone.example.org", "tar.loadbalancer.com", endpoint.RecordTypeCNAME, ""),
			newEndpointWithOwner("txt.tar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner-2\"", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("foobar.test-zone.example.org", "foobar.loadbalancer.com", endpoint.RecordTypeCNAME, ""),
			newEndpointWithOwner("foobar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("dualstack.test-zone.example.org", "1.1.1.1", endpoint.RecordTypeA, ""),
			newEndpointWithOwner("dualstack.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("dualstack.test-zone.example.org", "2001:DB8::1", endpoint.RecordTypeAAAA, ""),
			newEndpointWithOwner("aaaa-dualstack.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner-2\"", endpoint.RecordTypeTXT, ""),
		},
	})
	expectedRecords := []*endpoint.Endpoint{
		{
			DNSName:    "foo.test-zone.example.org",
			Targets:    endpoint.Targets{"foo.loadbalancer.com"},
			RecordType: endpoint.RecordTypeCNAME,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "",
			},
		},
		{
			DNSName:    "bar.test-zone.example.org",
			Targets:    endpoint.Targets{"my-domain.com"},
			RecordType: endpoint.RecordTypeCNAME,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "",
			},
		},
		{
			DNSName:    "alias.test-zone.example.org",
			Targets:    endpoint.Targets{"my-domain.com"},
			RecordType: endpoint.RecordTypeA,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "owner",
			},
			ProviderSpecific: []endpoint.ProviderSpecificProperty{
				{
					Name:  "alias",
					Value: "true",
				},
			},
		},
		{
			DNSName:    "txt.bar.test-zone.example.org",
			Targets:    endpoint.Targets{"baz.test-zone.example.org"},
			RecordType: endpoint.RecordTypeCNAME,
			Labels: map[string]string{
				endpoint.OwnerLabelKey:    "owner",
				endpoint.ResourceLabelKey: "ingress/default/my-ingress",
			},
		},
		{
			DNSName:    "qux.test-zone.example.org",
			Targets:    endpoint.Targets{"random"},
			RecordType: endpoint.RecordTypeTXT,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "",
			},
		},
		{
			DNSName:    "tar.test-zone.example.org",
			Targets:    endpoint.Targets{"tar.loadbalancer.com"},
			RecordType: endpoint.RecordTypeCNAME,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "",
			},
		},
		{
			DNSName:    "foobar.test-zone.example.org",
			Targets:    endpoint.Targets{"foobar.loadbalancer.com"},
			RecordType: endpoint.RecordTypeCNAME,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "owner",
			},
		},
		{
			DNSName:    "dualstack.test-zone.example.org",
			Targets:    endpoint.Targets{"1.1.1.1"},
			RecordType: endpoint.RecordTypeA,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "owner",
			},
		},
		{
			DNSName:    "dualstack.test-zone.example.org",
			Targets:    endpoint.Targets{"2001:DB8::1"},
			RecordType: endpoint.RecordTypeAAAA,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "owner-2",
			},
		},
	}

	r, _ := NewTXTRegistry(p, "", "", "owner", time.Hour, "", []string{}, []string{}, false, nil)
	records, _ := r.Records(ctx)

	assert.True(t, testutils.SameEndpoints(records, expectedRecords))
}

func testTXTRegistryRecordsPrefixedTemplated(t *testing.T) {
	ctx := context.Background()
	p := inmemory.NewInMemoryProvider()
	p.CreateZone(testZone)
	p.ApplyChanges(ctx, &plan.Changes{
		Create: []*endpoint.Endpoint{
			newEndpointWithOwner("foo.test-zone.example.org", "1.1.1.1", endpoint.RecordTypeA, ""),
			newEndpointWithOwner("txt-a.foo.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
		},
	})
	expectedRecords := []*endpoint.Endpoint{
		{
			DNSName:    "foo.test-zone.example.org",
			Targets:    endpoint.Targets{"1.1.1.1"},
			RecordType: endpoint.RecordTypeA,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "owner",
			},
		},
	}

	r, _ := NewTXTRegistry(p, "txt-%{record_type}.", "", "owner", time.Hour, "wc", []string{}, []string{}, false, nil)
	records, _ := r.Records(ctx)

	assert.True(t, testutils.SameEndpoints(records, expectedRecords))

	r, _ = NewTXTRegistry(p, "TxT-%{record_type}.", "", "owner", time.Hour, "wc", []string{}, []string{}, false, nil)
	records, _ = r.Records(ctx)

	assert.True(t, testutils.SameEndpoints(records, expectedRecords))
}

func testTXTRegistryRecordsSuffixedTemplated(t *testing.T) {
	ctx := context.Background()
	p := inmemory.NewInMemoryProvider()
	p.CreateZone(testZone)
	p.ApplyChanges(ctx, &plan.Changes{
		Create: []*endpoint.Endpoint{
			newEndpointWithOwner("bar.test-zone.example.org", "8.8.8.8", endpoint.RecordTypeCNAME, ""),
			newEndpointWithOwner("bartxtcname.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
		},
	})
	expectedRecords := []*endpoint.Endpoint{
		{
			DNSName:    "bar.test-zone.example.org",
			Targets:    endpoint.Targets{"8.8.8.8"},
			RecordType: endpoint.RecordTypeCNAME,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "owner",
			},
		},
	}

	r, _ := NewTXTRegistry(p, "", "txt%{record_type}", "owner", time.Hour, "wc", []string{}, []string{}, false, nil)
	records, _ := r.Records(ctx)

	assert.True(t, testutils.SameEndpoints(records, expectedRecords))

	r, _ = NewTXTRegistry(p, "", "TxT%{record_type}", "owner", time.Hour, "wc", []string{}, []string{}, false, nil)
	records, _ = r.Records(ctx)

	assert.True(t, testutils.SameEndpoints(records, expectedRecords))
}

func testTXTRegistryApplyChanges(t *testing.T) {
	t.Run("With Prefix", testTXTRegistryApplyChangesWithPrefix)
	t.Run("With Templated Prefix", testTXTRegistryApplyChangesWithTemplatedPrefix)
	t.Run("With Templated Suffix", testTXTRegistryApplyChangesWithTemplatedSuffix)
	t.Run("With Suffix", testTXTRegistryApplyChangesWithSuffix)
	t.Run("No prefix", testTXTRegistryApplyChangesNoPrefix)
}

func testTXTRegistryApplyChangesWithPrefix(t *testing.T) {
	p := inmemory.NewInMemoryProvider()
	p.CreateZone(testZone)
	ctxEndpoints := []*endpoint.Endpoint{}
	ctx := context.WithValue(context.Background(), provider.RecordsContextKey, ctxEndpoints)
	p.OnApplyChanges = func(ctx context.Context, got *plan.Changes) {
		assert.Equal(t, ctxEndpoints, ctx.Value(provider.RecordsContextKey))
	}
	p.ApplyChanges(ctx, &plan.Changes{
		Create: []*endpoint.Endpoint{
			newEndpointWithOwner("foo.test-zone.example.org", "foo.loadbalancer.com", endpoint.RecordTypeCNAME, ""),
			newEndpointWithOwner("bar.test-zone.example.org", "my-domain.com", endpoint.RecordTypeCNAME, ""),
			newEndpointWithOwner("txt.bar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("txt.bar.test-zone.example.org", "baz.test-zone.example.org", endpoint.RecordTypeCNAME, ""),
			newEndpointWithOwner("qux.test-zone.example.org", "random", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("tar.test-zone.example.org", "tar.loadbalancer.com", endpoint.RecordTypeCNAME, ""),
			newEndpointWithOwner("txt.tar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("txt.cname-tar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("foobar.test-zone.example.org", "foobar.loadbalancer.com", endpoint.RecordTypeCNAME, ""),
			newEndpointWithOwner("txt.foobar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("txt.cname-foobar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("multiple.test-zone.example.org", "lb1.loadbalancer.com", endpoint.RecordTypeCNAME, "").WithSetIdentifier("test-set-1"),
			newEndpointWithOwner("txt.multiple.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, "").WithSetIdentifier("test-set-1"),
			newEndpointWithOwner("txt.cname-multiple.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, "").WithSetIdentifier("test-set-1"),
			newEndpointWithOwner("multiple.test-zone.example.org", "lb2.loadbalancer.com", endpoint.RecordTypeCNAME, "").WithSetIdentifier("test-set-2"),
			newEndpointWithOwner("txt.multiple.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, "").WithSetIdentifier("test-set-2"),
			newEndpointWithOwner("txt.cname-multiple.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, "").WithSetIdentifier("test-set-2"),
		},
	})
	r, _ := NewTXTRegistry(p, "txt.", "", "owner", time.Hour, "", []string{}, []string{}, false, nil)

	changes := &plan.Changes{
		Create: []*endpoint.Endpoint{
			newEndpointWithOwnerResource("new-record-1.test-zone.example.org", "new-loadbalancer-1.lb.com", endpoint.RecordTypeCNAME, "", "ingress/default/my-ingress"),
			newEndpointWithOwnerResource("multiple.test-zone.example.org", "lb3.loadbalancer.com", endpoint.RecordTypeCNAME, "", "ingress/default/my-ingress").WithSetIdentifier("test-set-3"),
			newEndpointWithOwnerResource("example", "new-loadbalancer-1.lb.com", endpoint.RecordTypeCNAME, "", "ingress/default/my-ingress"),
		},
		Delete: []*endpoint.Endpoint{
			newEndpointWithOwner("foobar.test-zone.example.org", "foobar.loadbalancer.com", endpoint.RecordTypeCNAME, "owner"),
			newEndpointWithOwner("multiple.test-zone.example.org", "lb1.loadbalancer.com", endpoint.RecordTypeCNAME, "owner").WithSetIdentifier("test-set-1"),
		},
		UpdateNew: []*endpoint.Endpoint{
			newEndpointWithOwnerResource("tar.test-zone.example.org", "new-tar.loadbalancer.com", endpoint.RecordTypeCNAME, "owner", "ingress/default/my-ingress-2"),
			newEndpointWithOwnerResource("multiple.test-zone.example.org", "new.loadbalancer.com", endpoint.RecordTypeCNAME, "owner", "ingress/default/my-ingress-2").WithSetIdentifier("test-set-2"),
		},
		UpdateOld: []*endpoint.Endpoint{
			newEndpointWithOwner("tar.test-zone.example.org", "tar.loadbalancer.com", endpoint.RecordTypeCNAME, "owner"),
			newEndpointWithOwner("multiple.test-zone.example.org", "lb2.loadbalancer.com", endpoint.RecordTypeCNAME, "owner").WithSetIdentifier("test-set-2"),
		},
	}
	expected := &plan.Changes{
		Create: []*endpoint.Endpoint{
			newEndpointWithOwnerResource("new-record-1.test-zone.example.org", "new-loadbalancer-1.lb.com", endpoint.RecordTypeCNAME, "owner", "ingress/default/my-ingress"),
			newEndpointWithOwnerAndOwnedRecord("txt.new-record-1.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner,external-dns/resource=ingress/default/my-ingress\"", endpoint.RecordTypeTXT, "", "new-record-1.test-zone.example.org"),
			newEndpointWithOwnerAndOwnedRecord("txt.cname-new-record-1.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner,external-dns/resource=ingress/default/my-ingress\"", endpoint.RecordTypeTXT, "", "new-record-1.test-zone.example.org"),
			newEndpointWithOwnerResource("multiple.test-zone.example.org", "lb3.loadbalancer.com", endpoint.RecordTypeCNAME, "owner", "ingress/default/my-ingress").WithSetIdentifier("test-set-3"),
			newEndpointWithOwnerAndOwnedRecord("txt.multiple.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner,external-dns/resource=ingress/default/my-ingress\"", endpoint.RecordTypeTXT, "", "multiple.test-zone.example.org").WithSetIdentifier("test-set-3"),
			newEndpointWithOwnerAndOwnedRecord("txt.cname-multiple.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner,external-dns/resource=ingress/default/my-ingress\"", endpoint.RecordTypeTXT, "", "multiple.test-zone.example.org").WithSetIdentifier("test-set-3"),
			newEndpointWithOwnerResource("example", "new-loadbalancer-1.lb.com", endpoint.RecordTypeCNAME, "owner", "ingress/default/my-ingress"),
			newEndpointWithOwnerAndOwnedRecord("txt.example", "\"heritage=external-dns,external-dns/owner=owner,external-dns/resource=ingress/default/my-ingress\"", endpoint.RecordTypeTXT, "", "example"),
			newEndpointWithOwnerAndOwnedRecord("txt.cname-example", "\"heritage=external-dns,external-dns/owner=owner,external-dns/resource=ingress/default/my-ingress\"", endpoint.RecordTypeTXT, "", "example"),
		},
		Delete: []*endpoint.Endpoint{
			newEndpointWithOwner("foobar.test-zone.example.org", "foobar.loadbalancer.com", endpoint.RecordTypeCNAME, "owner"),
			newEndpointWithOwnerAndOwnedRecord("txt.foobar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, "", "foobar.test-zone.example.org"),
			newEndpointWithOwnerAndOwnedRecord("txt.cname-foobar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, "", "foobar.test-zone.example.org"),
			newEndpointWithOwner("multiple.test-zone.example.org", "lb1.loadbalancer.com", endpoint.RecordTypeCNAME, "owner").WithSetIdentifier("test-set-1"),
			newEndpointWithOwnerAndOwnedRecord("txt.multiple.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, "", "multiple.test-zone.example.org").WithSetIdentifier("test-set-1"),
			newEndpointWithOwnerAndOwnedRecord("txt.cname-multiple.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, "", "multiple.test-zone.example.org").WithSetIdentifier("test-set-1"),
		},
		UpdateNew: []*endpoint.Endpoint{
			newEndpointWithOwnerResource("tar.test-zone.example.org", "new-tar.loadbalancer.com", endpoint.RecordTypeCNAME, "owner", "ingress/default/my-ingress-2"),
			newEndpointWithOwnerAndOwnedRecord("txt.tar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner,external-dns/resource=ingress/default/my-ingress-2\"", endpoint.RecordTypeTXT, "", "tar.test-zone.example.org"),
			newEndpointWithOwnerAndOwnedRecord("txt.cname-tar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner,external-dns/resource=ingress/default/my-ingress-2\"", endpoint.RecordTypeTXT, "", "tar.test-zone.example.org"),
			newEndpointWithOwnerResource("multiple.test-zone.example.org", "new.loadbalancer.com", endpoint.RecordTypeCNAME, "owner", "ingress/default/my-ingress-2").WithSetIdentifier("test-set-2"),
			newEndpointWithOwnerAndOwnedRecord("txt.multiple.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner,external-dns/resource=ingress/default/my-ingress-2\"", endpoint.RecordTypeTXT, "", "multiple.test-zone.example.org").WithSetIdentifier("test-set-2"),
			newEndpointWithOwnerAndOwnedRecord("txt.cname-multiple.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner,external-dns/resource=ingress/default/my-ingress-2\"", endpoint.RecordTypeTXT, "", "multiple.test-zone.example.org").WithSetIdentifier("test-set-2"),
		},
		UpdateOld: []*endpoint.Endpoint{
			newEndpointWithOwner("tar.test-zone.example.org", "tar.loadbalancer.com", endpoint.RecordTypeCNAME, "owner"),
			newEndpointWithOwnerAndOwnedRecord("txt.tar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, "", "tar.test-zone.example.org"),
			newEndpointWithOwnerAndOwnedRecord("txt.cname-tar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, "", "tar.test-zone.example.org"),
			newEndpointWithOwner("multiple.test-zone.example.org", "lb2.loadbalancer.com", endpoint.RecordTypeCNAME, "owner").WithSetIdentifier("test-set-2"),
			newEndpointWithOwnerAndOwnedRecord("txt.multiple.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, "", "multiple.test-zone.example.org").WithSetIdentifier("test-set-2"),
			newEndpointWithOwnerAndOwnedRecord("txt.cname-multiple.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, "", "multiple.test-zone.example.org").WithSetIdentifier("test-set-2"),
		},
	}
	p.OnApplyChanges = func(ctx context.Context, got *plan.Changes) {
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
		assert.Equal(t, nil, ctx.Value(provider.RecordsContextKey))
	}
	err := r.ApplyChanges(ctx, changes)
	require.NoError(t, err)
}

func testTXTRegistryApplyChangesWithTemplatedPrefix(t *testing.T) {
	p := inmemory.NewInMemoryProvider()
	p.CreateZone(testZone)
	ctxEndpoints := []*endpoint.Endpoint{}
	ctx := context.WithValue(context.Background(), provider.RecordsContextKey, ctxEndpoints)
	p.OnApplyChanges = func(ctx context.Context, got *plan.Changes) {
		assert.Equal(t, ctxEndpoints, ctx.Value(provider.RecordsContextKey))
	}
	p.ApplyChanges(ctx, &plan.Changes{
		Create: []*endpoint.Endpoint{},
	})
	r, _ := NewTXTRegistry(p, "prefix%{record_type}.", "", "owner", time.Hour, "", []string{}, []string{}, false, nil)
	changes := &plan.Changes{
		Create: []*endpoint.Endpoint{
			newEndpointWithOwnerResource("new-record-1.test-zone.example.org", "new-loadbalancer-1.lb.com", endpoint.RecordTypeCNAME, "", "ingress/default/my-ingress"),
		},
		Delete:    []*endpoint.Endpoint{},
		UpdateOld: []*endpoint.Endpoint{},
		UpdateNew: []*endpoint.Endpoint{},
	}
	expected := &plan.Changes{
		Create: []*endpoint.Endpoint{
			newEndpointWithOwnerResource("new-record-1.test-zone.example.org", "new-loadbalancer-1.lb.com", endpoint.RecordTypeCNAME, "owner", "ingress/default/my-ingress"),
			newEndpointWithOwnerAndOwnedRecord("prefixcname.new-record-1.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner,external-dns/resource=ingress/default/my-ingress\"", endpoint.RecordTypeTXT, "", "new-record-1.test-zone.example.org"),
		},
	}
	p.OnApplyChanges = func(ctx context.Context, got *plan.Changes) {
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
		assert.Equal(t, nil, ctx.Value(provider.RecordsContextKey))
	}
	err := r.ApplyChanges(ctx, changes)
	require.NoError(t, err)
}

func testTXTRegistryApplyChangesWithTemplatedSuffix(t *testing.T) {
	p := inmemory.NewInMemoryProvider()
	p.CreateZone(testZone)
	ctxEndpoints := []*endpoint.Endpoint{}
	ctx := context.WithValue(context.Background(), provider.RecordsContextKey, ctxEndpoints)
	p.OnApplyChanges = func(ctx context.Context, got *plan.Changes) {
		assert.Equal(t, ctxEndpoints, ctx.Value(provider.RecordsContextKey))
	}
	r, _ := NewTXTRegistry(p, "", "-%{record_type}suffix", "owner", time.Hour, "", []string{}, []string{}, false, nil)
	changes := &plan.Changes{
		Create: []*endpoint.Endpoint{
			newEndpointWithOwnerResource("new-record-1.test-zone.example.org", "new-loadbalancer-1.lb.com", endpoint.RecordTypeCNAME, "", "ingress/default/my-ingress"),
		},
		Delete:    []*endpoint.Endpoint{},
		UpdateOld: []*endpoint.Endpoint{},
		UpdateNew: []*endpoint.Endpoint{},
	}
	expected := &plan.Changes{
		Create: []*endpoint.Endpoint{
			newEndpointWithOwnerResource("new-record-1.test-zone.example.org", "new-loadbalancer-1.lb.com", endpoint.RecordTypeCNAME, "owner", "ingress/default/my-ingress"),
			newEndpointWithOwnerAndOwnedRecord("new-record-1-cnamesuffix.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner,external-dns/resource=ingress/default/my-ingress\"", endpoint.RecordTypeTXT, "", "new-record-1.test-zone.example.org"),
		},
	}
	p.OnApplyChanges = func(ctx context.Context, got *plan.Changes) {
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
		assert.Equal(t, nil, ctx.Value(provider.RecordsContextKey))
	}
	err := r.ApplyChanges(ctx, changes)
	require.NoError(t, err)
}

func testTXTRegistryApplyChangesWithSuffix(t *testing.T) {
	p := inmemory.NewInMemoryProvider()
	p.CreateZone(testZone)
	ctxEndpoints := []*endpoint.Endpoint{}
	ctx := context.WithValue(context.Background(), provider.RecordsContextKey, ctxEndpoints)
	p.OnApplyChanges = func(ctx context.Context, got *plan.Changes) {
		assert.Equal(t, ctxEndpoints, ctx.Value(provider.RecordsContextKey))
	}
	p.ApplyChanges(ctx, &plan.Changes{
		Create: []*endpoint.Endpoint{
			newEndpointWithOwner("foo.test-zone.example.org", "foo.loadbalancer.com", endpoint.RecordTypeCNAME, ""),
			newEndpointWithOwner("bar.test-zone.example.org", "my-domain.com", endpoint.RecordTypeCNAME, ""),
			newEndpointWithOwner("bar-txt.test-zone.example.org", "baz.test-zone.example.org", endpoint.RecordTypeCNAME, ""),
			newEndpointWithOwner("bar-txt.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("cname-bar-txt.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("qux.test-zone.example.org", "random", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("tar.test-zone.example.org", "tar.loadbalancer.com", endpoint.RecordTypeCNAME, ""),
			newEndpointWithOwner("tar-txt.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("cname-tar-txt.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("foobar.test-zone.example.org", "foobar.loadbalancer.com", endpoint.RecordTypeCNAME, ""),
			newEndpointWithOwner("foobar-txt.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("cname-foobar-txt.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("multiple.test-zone.example.org", "lb1.loadbalancer.com", endpoint.RecordTypeCNAME, "").WithSetIdentifier("test-set-1"),
			newEndpointWithOwner("multiple-txt.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, "").WithSetIdentifier("test-set-1"),
			newEndpointWithOwner("cname-multiple-txt.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, "").WithSetIdentifier("test-set-1"),
			newEndpointWithOwner("multiple.test-zone.example.org", "lb2.loadbalancer.com", endpoint.RecordTypeCNAME, "").WithSetIdentifier("test-set-2"),
			newEndpointWithOwner("multiple-txt.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, "").WithSetIdentifier("test-set-2"),
			newEndpointWithOwner("cname-multiple-txt.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, "").WithSetIdentifier("test-set-2"),
		},
	})
	r, _ := NewTXTRegistry(p, "", "-txt", "owner", time.Hour, "wildcard", []string{}, []string{}, false, nil)

	changes := &plan.Changes{
		Create: []*endpoint.Endpoint{
			newEndpointWithOwnerResource("new-record-1.test-zone.example.org", "new-loadbalancer-1.lb.com", endpoint.RecordTypeCNAME, "", "ingress/default/my-ingress"),
			newEndpointWithOwnerResource("multiple.test-zone.example.org", "lb3.loadbalancer.com", endpoint.RecordTypeCNAME, "", "ingress/default/my-ingress").WithSetIdentifier("test-set-3"),
			newEndpointWithOwnerResource("example", "new-loadbalancer-1.lb.com", endpoint.RecordTypeCNAME, "", "ingress/default/my-ingress"),
			newEndpointWithOwnerResource("*.wildcard.test-zone.example.org", "new-loadbalancer-1.lb.com", endpoint.RecordTypeCNAME, "", "ingress/default/my-ingress"),
		},
		Delete: []*endpoint.Endpoint{
			newEndpointWithOwner("foobar.test-zone.example.org", "foobar.loadbalancer.com", endpoint.RecordTypeCNAME, "owner"),
			newEndpointWithOwner("multiple.test-zone.example.org", "lb1.loadbalancer.com", endpoint.RecordTypeCNAME, "owner").WithSetIdentifier("test-set-1"),
		},
		UpdateNew: []*endpoint.Endpoint{
			newEndpointWithOwnerResource("tar.test-zone.example.org", "new-tar.loadbalancer.com", endpoint.RecordTypeCNAME, "owner", "ingress/default/my-ingress-2"),
			newEndpointWithOwnerResource("multiple.test-zone.example.org", "new.loadbalancer.com", endpoint.RecordTypeCNAME, "owner", "ingress/default/my-ingress-2").WithSetIdentifier("test-set-2"),
		},
		UpdateOld: []*endpoint.Endpoint{
			newEndpointWithOwner("tar.test-zone.example.org", "tar.loadbalancer.com", endpoint.RecordTypeCNAME, "owner"),
			newEndpointWithOwner("multiple.test-zone.example.org", "lb2.loadbalancer.com", endpoint.RecordTypeCNAME, "owner").WithSetIdentifier("test-set-2"),
		},
	}
	expected := &plan.Changes{
		Create: []*endpoint.Endpoint{
			newEndpointWithOwnerResource("new-record-1.test-zone.example.org", "new-loadbalancer-1.lb.com", endpoint.RecordTypeCNAME, "owner", "ingress/default/my-ingress"),
			newEndpointWithOwnerAndOwnedRecord("new-record-1-txt.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner,external-dns/resource=ingress/default/my-ingress\"", endpoint.RecordTypeTXT, "", "new-record-1.test-zone.example.org"),
			newEndpointWithOwnerAndOwnedRecord("cname-new-record-1-txt.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner,external-dns/resource=ingress/default/my-ingress\"", endpoint.RecordTypeTXT, "", "new-record-1.test-zone.example.org"),
			newEndpointWithOwnerResource("multiple.test-zone.example.org", "lb3.loadbalancer.com", endpoint.RecordTypeCNAME, "owner", "ingress/default/my-ingress").WithSetIdentifier("test-set-3"),
			newEndpointWithOwnerAndOwnedRecord("multiple-txt.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner,external-dns/resource=ingress/default/my-ingress\"", endpoint.RecordTypeTXT, "", "multiple.test-zone.example.org").WithSetIdentifier("test-set-3"),
			newEndpointWithOwnerAndOwnedRecord("cname-multiple-txt.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner,external-dns/resource=ingress/default/my-ingress\"", endpoint.RecordTypeTXT, "", "multiple.test-zone.example.org").WithSetIdentifier("test-set-3"),
			newEndpointWithOwnerResource("example", "new-loadbalancer-1.lb.com", endpoint.RecordTypeCNAME, "owner", "ingress/default/my-ingress"),
			newEndpointWithOwnerAndOwnedRecord("example-txt", "\"heritage=external-dns,external-dns/owner=owner,external-dns/resource=ingress/default/my-ingress\"", endpoint.RecordTypeTXT, "", "example"),
			newEndpointWithOwnerAndOwnedRecord("cname-example-txt", "\"heritage=external-dns,external-dns/owner=owner,external-dns/resource=ingress/default/my-ingress\"", endpoint.RecordTypeTXT, "", "example"),
			newEndpointWithOwnerResource("*.wildcard.test-zone.example.org", "new-loadbalancer-1.lb.com", endpoint.RecordTypeCNAME, "owner", "ingress/default/my-ingress"),
			newEndpointWithOwnerAndOwnedRecord("wildcard-txt.wildcard.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner,external-dns/resource=ingress/default/my-ingress\"", endpoint.RecordTypeTXT, "", "*.wildcard.test-zone.example.org"),
			newEndpointWithOwnerAndOwnedRecord("cname-wildcard-txt.wildcard.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner,external-dns/resource=ingress/default/my-ingress\"", endpoint.RecordTypeTXT, "", "*.wildcard.test-zone.example.org"),
		},
		Delete: []*endpoint.Endpoint{
			newEndpointWithOwner("foobar.test-zone.example.org", "foobar.loadbalancer.com", endpoint.RecordTypeCNAME, "owner"),
			newEndpointWithOwnerAndOwnedRecord("foobar-txt.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, "", "foobar.test-zone.example.org"),
			newEndpointWithOwnerAndOwnedRecord("cname-foobar-txt.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, "", "foobar.test-zone.example.org"),
			newEndpointWithOwner("multiple.test-zone.example.org", "lb1.loadbalancer.com", endpoint.RecordTypeCNAME, "owner").WithSetIdentifier("test-set-1"),
			newEndpointWithOwnerAndOwnedRecord("multiple-txt.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, "", "multiple.test-zone.example.org").WithSetIdentifier("test-set-1"),
			newEndpointWithOwnerAndOwnedRecord("cname-multiple-txt.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, "", "multiple.test-zone.example.org").WithSetIdentifier("test-set-1"),
		},
		UpdateNew: []*endpoint.Endpoint{
			newEndpointWithOwnerResource("tar.test-zone.example.org", "new-tar.loadbalancer.com", endpoint.RecordTypeCNAME, "owner", "ingress/default/my-ingress-2"),
			newEndpointWithOwnerAndOwnedRecord("tar-txt.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner,external-dns/resource=ingress/default/my-ingress-2\"", endpoint.RecordTypeTXT, "", "tar.test-zone.example.org"),
			newEndpointWithOwnerAndOwnedRecord("cname-tar-txt.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner,external-dns/resource=ingress/default/my-ingress-2\"", endpoint.RecordTypeTXT, "", "tar.test-zone.example.org"),
			newEndpointWithOwnerResource("multiple.test-zone.example.org", "new.loadbalancer.com", endpoint.RecordTypeCNAME, "owner", "ingress/default/my-ingress-2").WithSetIdentifier("test-set-2"),
			newEndpointWithOwnerAndOwnedRecord("multiple-txt.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner,external-dns/resource=ingress/default/my-ingress-2\"", endpoint.RecordTypeTXT, "", "multiple.test-zone.example.org").WithSetIdentifier("test-set-2"),
			newEndpointWithOwnerAndOwnedRecord("cname-multiple-txt.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner,external-dns/resource=ingress/default/my-ingress-2\"", endpoint.RecordTypeTXT, "", "multiple.test-zone.example.org").WithSetIdentifier("test-set-2"),
		},
		UpdateOld: []*endpoint.Endpoint{
			newEndpointWithOwner("tar.test-zone.example.org", "tar.loadbalancer.com", endpoint.RecordTypeCNAME, "owner"),
			newEndpointWithOwnerAndOwnedRecord("tar-txt.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, "", "tar.test-zone.example.org"),
			newEndpointWithOwnerAndOwnedRecord("cname-tar-txt.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, "", "tar.test-zone.example.org"),
			newEndpointWithOwner("multiple.test-zone.example.org", "lb2.loadbalancer.com", endpoint.RecordTypeCNAME, "owner").WithSetIdentifier("test-set-2"),
			newEndpointWithOwnerAndOwnedRecord("multiple-txt.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, "", "multiple.test-zone.example.org").WithSetIdentifier("test-set-2"),
			newEndpointWithOwnerAndOwnedRecord("cname-multiple-txt.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, "", "multiple.test-zone.example.org").WithSetIdentifier("test-set-2"),
		},
	}
	p.OnApplyChanges = func(ctx context.Context, got *plan.Changes) {
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
		assert.Equal(t, nil, ctx.Value(provider.RecordsContextKey))
	}
	err := r.ApplyChanges(ctx, changes)
	require.NoError(t, err)
}

func testTXTRegistryApplyChangesNoPrefix(t *testing.T) {
	p := inmemory.NewInMemoryProvider()
	p.CreateZone(testZone)
	ctxEndpoints := []*endpoint.Endpoint{}
	ctx := context.WithValue(context.Background(), provider.RecordsContextKey, ctxEndpoints)
	p.OnApplyChanges = func(ctx context.Context, got *plan.Changes) {
		assert.Equal(t, ctxEndpoints, ctx.Value(provider.RecordsContextKey))
	}
	p.ApplyChanges(ctx, &plan.Changes{
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
			newEndpointWithOwner("cname-foobar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
		},
	})
	r, _ := NewTXTRegistry(p, "", "", "owner", time.Hour, "", []string{}, []string{}, false, nil)

	changes := &plan.Changes{
		Create: []*endpoint.Endpoint{
			newEndpointWithOwner("new-record-1.test-zone.example.org", "new-loadbalancer-1.lb.com", endpoint.RecordTypeCNAME, ""),
			newEndpointWithOwner("example", "new-loadbalancer-1.lb.com", endpoint.RecordTypeCNAME, ""),
			newEndpointWithOwner("new-alias.test-zone.example.org", "my-domain.com", endpoint.RecordTypeA, "").WithProviderSpecific("alias", "true"),
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
			newEndpointWithOwner("new-record-1.test-zone.example.org", "new-loadbalancer-1.lb.com", endpoint.RecordTypeCNAME, "owner"),
			newEndpointWithOwnerAndOwnedRecord("new-record-1.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, "", "new-record-1.test-zone.example.org"),
			newEndpointWithOwnerAndOwnedRecord("cname-new-record-1.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, "", "new-record-1.test-zone.example.org"),
			newEndpointWithOwner("example", "new-loadbalancer-1.lb.com", endpoint.RecordTypeCNAME, "owner"),
			newEndpointWithOwnerAndOwnedRecord("example", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, "", "example"),
			newEndpointWithOwnerAndOwnedRecord("cname-example", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, "", "example"),
			newEndpointWithOwner("new-alias.test-zone.example.org", "my-domain.com", endpoint.RecordTypeA, "owner").WithProviderSpecific("alias", "true"),
			// TODO: It's not clear why the TXT registry copies ProviderSpecificProperties to ownership records; that doesn't seem correct.
			newEndpointWithOwnerAndOwnedRecord("new-alias.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, "", "new-alias.test-zone.example.org").WithProviderSpecific("alias", "true"),
			newEndpointWithOwnerAndOwnedRecord("cname-new-alias.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, "", "new-alias.test-zone.example.org").WithProviderSpecific("alias", "true"),
		},
		Delete: []*endpoint.Endpoint{
			newEndpointWithOwner("foobar.test-zone.example.org", "foobar.loadbalancer.com", endpoint.RecordTypeCNAME, "owner"),
			newEndpointWithOwnerAndOwnedRecord("foobar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, "", "foobar.test-zone.example.org"),
			newEndpointWithOwnerAndOwnedRecord("cname-foobar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, "", "foobar.test-zone.example.org"),
		},
		UpdateNew: []*endpoint.Endpoint{},
		UpdateOld: []*endpoint.Endpoint{},
	}
	p.OnApplyChanges = func(ctx context.Context, got *plan.Changes) {
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
		assert.Equal(t, nil, ctx.Value(provider.RecordsContextKey))
	}
	err := r.ApplyChanges(ctx, changes)
	require.NoError(t, err)
}

func testTXTRegistryMissingRecords(t *testing.T) {
	t.Run("No prefix", testTXTRegistryMissingRecordsNoPrefix)
	t.Run("With Prefix", testTXTRegistryMissingRecordsWithPrefix)
}

func testTXTRegistryMissingRecordsNoPrefix(t *testing.T) {
	ctx := context.Background()
	p := inmemory.NewInMemoryProvider()
	p.CreateZone(testZone)
	p.ApplyChanges(ctx, &plan.Changes{
		Create: []*endpoint.Endpoint{
			newEndpointWithOwner("oldformat.test-zone.example.org", "foo.loadbalancer.com", endpoint.RecordTypeCNAME, ""),
			newEndpointWithOwner("oldformat.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("oldformat2.test-zone.example.org", "bar.loadbalancer.com", endpoint.RecordTypeA, ""),
			newEndpointWithOwner("oldformat2.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("newformat.test-zone.example.org", "foobar.nameserver.com", endpoint.RecordTypeNS, ""),
			newEndpointWithOwner("ns-newformat.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("newformat.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("noheritage.test-zone.example.org", "random", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("oldformat-otherowner.test-zone.example.org", "bar.loadbalancer.com", endpoint.RecordTypeA, ""),
			newEndpointWithOwner("oldformat-otherowner.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=otherowner\"", endpoint.RecordTypeTXT, ""),
			endpoint.NewEndpoint("unmanaged1.test-zone.example.org", endpoint.RecordTypeA, "unmanaged1.loadbalancer.com"),
			endpoint.NewEndpoint("unmanaged2.test-zone.example.org", endpoint.RecordTypeCNAME, "unmanaged2.loadbalancer.com"),
			newEndpointWithOwner("this-is-a-63-characters-long-label-that-we-do-expect-will-work.test-zone.example.org", "foo.loadbalancer.com", endpoint.RecordTypeCNAME, ""),
			newEndpointWithOwner("this-is-a-63-characters-long-label-that-we-do-expect-will-work.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
		},
	})
	expectedRecords := []*endpoint.Endpoint{
		{
			DNSName:    "oldformat.test-zone.example.org",
			Targets:    endpoint.Targets{"foo.loadbalancer.com"},
			RecordType: endpoint.RecordTypeCNAME,
			Labels: map[string]string{
				// owner was added from the TXT record's target
				endpoint.OwnerLabelKey: "owner",
			},
			ProviderSpecific: []endpoint.ProviderSpecificProperty{
				{
					Name:  "txt/force-update",
					Value: "true",
				},
			},
		},
		{
			DNSName:    "oldformat2.test-zone.example.org",
			Targets:    endpoint.Targets{"bar.loadbalancer.com"},
			RecordType: endpoint.RecordTypeA,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "owner",
			},
			ProviderSpecific: []endpoint.ProviderSpecificProperty{
				{
					Name:  "txt/force-update",
					Value: "true",
				},
			},
		},
		{
			DNSName:    "newformat.test-zone.example.org",
			Targets:    endpoint.Targets{"foobar.nameserver.com"},
			RecordType: endpoint.RecordTypeNS,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "owner",
			},
		},
		// Only TXT records with the wrong heritage are returned by Records()
		{
			DNSName:    "noheritage.test-zone.example.org",
			Targets:    endpoint.Targets{"random"},
			RecordType: endpoint.RecordTypeTXT,
			Labels: map[string]string{
				// No owner because it's not external-dns heritage
				endpoint.OwnerLabelKey: "",
			},
		},
		{
			DNSName:    "oldformat-otherowner.test-zone.example.org",
			Targets:    endpoint.Targets{"bar.loadbalancer.com"},
			RecordType: endpoint.RecordTypeA,
			Labels: map[string]string{
				// Records() retrieves all the records of the zone, no matter the owner
				endpoint.OwnerLabelKey: "otherowner",
			},
		},
		{
			DNSName:    "unmanaged1.test-zone.example.org",
			Targets:    endpoint.Targets{"unmanaged1.loadbalancer.com"},
			RecordType: endpoint.RecordTypeA,
		},
		{
			DNSName:    "unmanaged2.test-zone.example.org",
			Targets:    endpoint.Targets{"unmanaged2.loadbalancer.com"},
			RecordType: endpoint.RecordTypeCNAME,
		},
		{
			DNSName:    "this-is-a-63-characters-long-label-that-we-do-expect-will-work.test-zone.example.org",
			Targets:    endpoint.Targets{"foo.loadbalancer.com"},
			RecordType: endpoint.RecordTypeCNAME,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "owner",
			},
		},
	}

	r, _ := NewTXTRegistry(p, "", "", "owner", time.Hour, "wc", []string{endpoint.RecordTypeCNAME, endpoint.RecordTypeA, endpoint.RecordTypeNS}, []string{}, false, nil)
	records, _ := r.Records(ctx)

	assert.True(t, testutils.SameEndpoints(records, expectedRecords))
}

func testTXTRegistryMissingRecordsWithPrefix(t *testing.T) {
	ctx := context.Background()
	p := inmemory.NewInMemoryProvider()
	p.CreateZone(testZone)
	p.ApplyChanges(ctx, &plan.Changes{
		Create: []*endpoint.Endpoint{
			newEndpointWithOwner("oldformat.test-zone.example.org", "foo.loadbalancer.com", endpoint.RecordTypeCNAME, ""),
			newEndpointWithOwner("txt.oldformat.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("oldformat2.test-zone.example.org", "bar.loadbalancer.com", endpoint.RecordTypeA, ""),
			newEndpointWithOwner("txt.oldformat2.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("newformat.test-zone.example.org", "foobar.nameserver.com", endpoint.RecordTypeNS, ""),
			newEndpointWithOwner("txt.ns-newformat.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("oldformat3.test-zone.example.org", "random", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("txt.oldformat3.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("txt.newformat.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("noheritage.test-zone.example.org", "random", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("oldformat-otherowner.test-zone.example.org", "bar.loadbalancer.com", endpoint.RecordTypeA, ""),
			newEndpointWithOwner("txt.oldformat-otherowner.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=otherowner\"", endpoint.RecordTypeTXT, ""),
			endpoint.NewEndpoint("unmanaged1.test-zone.example.org", endpoint.RecordTypeA, "unmanaged1.loadbalancer.com"),
			endpoint.NewEndpoint("unmanaged2.test-zone.example.org", endpoint.RecordTypeCNAME, "unmanaged2.loadbalancer.com"),
		},
	})
	expectedRecords := []*endpoint.Endpoint{
		{
			DNSName:    "oldformat.test-zone.example.org",
			Targets:    endpoint.Targets{"foo.loadbalancer.com"},
			RecordType: endpoint.RecordTypeCNAME,
			Labels: map[string]string{
				// owner was added from the TXT record's target
				endpoint.OwnerLabelKey: "owner",
			},
			ProviderSpecific: []endpoint.ProviderSpecificProperty{
				{
					Name:  "txt/force-update",
					Value: "true",
				},
			},
		},
		{
			DNSName:    "oldformat2.test-zone.example.org",
			Targets:    endpoint.Targets{"bar.loadbalancer.com"},
			RecordType: endpoint.RecordTypeA,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "owner",
			},
			ProviderSpecific: []endpoint.ProviderSpecificProperty{
				{
					Name:  "txt/force-update",
					Value: "true",
				},
			},
		},
		{
			DNSName:    "oldformat3.test-zone.example.org",
			Targets:    endpoint.Targets{"random"},
			RecordType: endpoint.RecordTypeTXT,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "owner",
			},
			ProviderSpecific: []endpoint.ProviderSpecificProperty{
				{
					Name:  "txt/force-update",
					Value: "true",
				},
			},
		},
		{
			DNSName:    "newformat.test-zone.example.org",
			Targets:    endpoint.Targets{"foobar.nameserver.com"},
			RecordType: endpoint.RecordTypeNS,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "owner",
			},
		},
		{
			DNSName:    "noheritage.test-zone.example.org",
			Targets:    endpoint.Targets{"random"},
			RecordType: endpoint.RecordTypeTXT,
			Labels: map[string]string{
				// No owner because it's not external-dns heritage
				endpoint.OwnerLabelKey: "",
			},
		},
		{
			DNSName:    "oldformat-otherowner.test-zone.example.org",
			Targets:    endpoint.Targets{"bar.loadbalancer.com"},
			RecordType: endpoint.RecordTypeA,
			Labels: map[string]string{
				// All the records of the zone are retrieved, no matter the owner
				endpoint.OwnerLabelKey: "otherowner",
			},
		},
		{
			DNSName:    "unmanaged1.test-zone.example.org",
			Targets:    endpoint.Targets{"unmanaged1.loadbalancer.com"},
			RecordType: endpoint.RecordTypeA,
		},
		{
			DNSName:    "unmanaged2.test-zone.example.org",
			Targets:    endpoint.Targets{"unmanaged2.loadbalancer.com"},
			RecordType: endpoint.RecordTypeCNAME,
		},
	}

	r, _ := NewTXTRegistry(p, "txt.", "", "owner", time.Hour, "wc", []string{endpoint.RecordTypeCNAME, endpoint.RecordTypeA, endpoint.RecordTypeNS, endpoint.RecordTypeTXT}, []string{}, false, nil)
	records, _ := r.Records(ctx)

	assert.True(t, testutils.SameEndpoints(records, expectedRecords))
}

func TestCacheMethods(t *testing.T) {
	cache := []*endpoint.Endpoint{
		newEndpointWithOwner("thing.com", "1.2.3.4", "A", "owner"),
		newEndpointWithOwner("thing1.com", "1.2.3.6", "A", "owner"),
		newEndpointWithOwner("thing2.com", "1.2.3.4", "CNAME", "owner"),
		newEndpointWithOwner("thing3.com", "1.2.3.4", "A", "owner"),
		newEndpointWithOwner("thing4.com", "1.2.3.4", "A", "owner"),
	}
	registry := &TXTRegistry{
		recordsCache:  cache,
		cacheInterval: time.Hour,
	}

	expectedCacheAfterAdd := []*endpoint.Endpoint{
		newEndpointWithOwner("thing.com", "1.2.3.4", "A", "owner"),
		newEndpointWithOwner("thing1.com", "1.2.3.6", "A", "owner"),
		newEndpointWithOwner("thing2.com", "1.2.3.4", "CNAME", "owner"),
		newEndpointWithOwner("thing3.com", "1.2.3.4", "A", "owner"),
		newEndpointWithOwner("thing4.com", "1.2.3.4", "A", "owner"),
		newEndpointWithOwner("thing4.com", "2001:DB8::1", "AAAA", "owner"),
		newEndpointWithOwner("thing5.com", "1.2.3.5", "A", "owner"),
	}

	expectedCacheAfterUpdate := []*endpoint.Endpoint{
		newEndpointWithOwner("thing1.com", "1.2.3.6", "A", "owner"),
		newEndpointWithOwner("thing2.com", "1.2.3.4", "CNAME", "owner"),
		newEndpointWithOwner("thing3.com", "1.2.3.4", "A", "owner"),
		newEndpointWithOwner("thing4.com", "1.2.3.4", "A", "owner"),
		newEndpointWithOwner("thing5.com", "1.2.3.5", "A", "owner"),
		newEndpointWithOwner("thing.com", "1.2.3.6", "A", "owner2"),
		newEndpointWithOwner("thing4.com", "2001:DB8::2", "AAAA", "owner"),
	}

	expectedCacheAfterDelete := []*endpoint.Endpoint{
		newEndpointWithOwner("thing1.com", "1.2.3.6", "A", "owner"),
		newEndpointWithOwner("thing2.com", "1.2.3.4", "CNAME", "owner"),
		newEndpointWithOwner("thing3.com", "1.2.3.4", "A", "owner"),
		newEndpointWithOwner("thing4.com", "1.2.3.4", "A", "owner"),
		newEndpointWithOwner("thing5.com", "1.2.3.5", "A", "owner"),
	}
	// test add cache
	registry.addToCache(newEndpointWithOwner("thing4.com", "2001:DB8::1", "AAAA", "owner"))
	registry.addToCache(newEndpointWithOwner("thing5.com", "1.2.3.5", "A", "owner"))

	if !reflect.DeepEqual(expectedCacheAfterAdd, registry.recordsCache) {
		t.Fatalf("expected endpoints should match endpoints from cache: expected %v, but got %v", expectedCacheAfterAdd, registry.recordsCache)
	}

	// test update cache
	registry.removeFromCache(newEndpointWithOwner("thing.com", "1.2.3.4", "A", "owner"))
	registry.addToCache(newEndpointWithOwner("thing.com", "1.2.3.6", "A", "owner2"))
	registry.removeFromCache(newEndpointWithOwner("thing4.com", "2001:DB8::1", "AAAA", "owner"))
	registry.addToCache(newEndpointWithOwner("thing4.com", "2001:DB8::2", "AAAA", "owner"))
	// ensure it was updated
	if !reflect.DeepEqual(expectedCacheAfterUpdate, registry.recordsCache) {
		t.Fatalf("expected endpoints should match endpoints from cache: expected %v, but got %v", expectedCacheAfterUpdate, registry.recordsCache)
	}

	// test deleting a record
	registry.removeFromCache(newEndpointWithOwner("thing.com", "1.2.3.6", "A", "owner2"))
	registry.removeFromCache(newEndpointWithOwner("thing4.com", "2001:DB8::2", "AAAA", "owner"))
	// ensure it was deleted
	if !reflect.DeepEqual(expectedCacheAfterDelete, registry.recordsCache) {
		t.Fatalf("expected endpoints should match endpoints from cache: expected %v, but got %v", expectedCacheAfterDelete, registry.recordsCache)
	}
}

func TestDropPrefix(t *testing.T) {
	mapper := newaffixNameMapper("foo-%{record_type}-", "", "")
	expectedOutput := "test.example.com"

	tests := []string{
		"foo-cname-test.example.com",
		"foo-a-test.example.com",
		"foo--test.example.com",
	}

	for _, tc := range tests {
		t.Run(tc, func(t *testing.T) {
			actualOutput, _ := mapper.dropAffixExtractType(tc)
			assert.Equal(t, expectedOutput, actualOutput)
		})
	}
}

func TestDropSuffix(t *testing.T) {
	mapper := newaffixNameMapper("", "-%{record_type}-foo", "")
	expectedOutput := "test.example.com"

	tests := []string{
		"test-a-foo.example.com",
		"test--foo.example.com",
	}

	for _, tc := range tests {
		t.Run(tc, func(t *testing.T) {
			r := strings.SplitN(tc, ".", 2)
			rClean, _ := mapper.dropAffixExtractType(r[0])
			actualOutput := rClean + "." + r[1]
			assert.Equal(t, expectedOutput, actualOutput)
		})
	}
}

func TestExtractRecordTypeDefaultPosition(t *testing.T) {
	tests := []struct {
		input        string
		expectedName string
		expectedType string
	}{
		{
			input:        "ns-zone.example.com",
			expectedName: "zone.example.com",
			expectedType: "NS",
		},
		{
			input:        "aaaa-zone.example.com",
			expectedName: "zone.example.com",
			expectedType: "AAAA",
		},
		{
			input:        "ptr-zone.example.com",
			expectedName: "ptr-zone.example.com",
			expectedType: "",
		},
		{
			input:        "zone.example.com",
			expectedName: "zone.example.com",
			expectedType: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			actualName, actualType := extractRecordTypeDefaultPosition(tc.input)
			assert.Equal(t, tc.expectedName, actualName)
			assert.Equal(t, tc.expectedType, actualType)
		})
	}
}

func TestToEndpointNameNewTXT(t *testing.T) {
	tests := []struct {
		name       string
		mapper     affixNameMapper
		domain     string
		txtDomain  string
		recordType string
	}{
		{
			name:       "prefix",
			mapper:     newaffixNameMapper("foo", "", ""),
			domain:     "example.com",
			recordType: "A",
			txtDomain:  "fooa-example.com",
		},
		{
			name:       "suffix",
			mapper:     newaffixNameMapper("", "foo", ""),
			domain:     "example.com",
			recordType: "AAAA",
			txtDomain:  "aaaa-examplefoo.com",
		},
		{
			name:       "prefix with dash",
			mapper:     newaffixNameMapper("foo-", "", ""),
			domain:     "example.com",
			recordType: "A",
			txtDomain:  "foo-a-example.com",
		},
		{
			name:       "suffix with dash",
			mapper:     newaffixNameMapper("", "-foo", ""),
			domain:     "example.com",
			recordType: "CNAME",
			txtDomain:  "cname-example-foo.com",
		},
		{
			name:       "prefix with dot",
			mapper:     newaffixNameMapper("foo.", "", ""),
			domain:     "example.com",
			recordType: "CNAME",
			txtDomain:  "foo.cname-example.com",
		},
		{
			name:       "suffix with dot",
			mapper:     newaffixNameMapper("", ".foo", ""),
			domain:     "example.com",
			recordType: "CNAME",
			txtDomain:  "cname-example.foo.com",
		},
		{
			name:       "prefix with multiple dots",
			mapper:     newaffixNameMapper("foo.bar.", "", ""),
			domain:     "example.com",
			recordType: "CNAME",
			txtDomain:  "foo.bar.cname-example.com",
		},
		{
			name:       "suffix with multiple dots",
			mapper:     newaffixNameMapper("", ".foo.bar.test", ""),
			domain:     "example.com",
			recordType: "CNAME",
			txtDomain:  "cname-example.foo.bar.test.com",
		},
		{
			name:       "templated prefix",
			mapper:     newaffixNameMapper("%{record_type}-foo", "", ""),
			domain:     "example.com",
			recordType: "A",
			txtDomain:  "a-fooexample.com",
		},
		{
			name:       "templated suffix",
			mapper:     newaffixNameMapper("", "foo-%{record_type}", ""),
			domain:     "example.com",
			recordType: "A",
			txtDomain:  "examplefoo-a.com",
		},
		{
			name:       "templated prefix with dot",
			mapper:     newaffixNameMapper("%{record_type}foo.", "", ""),
			domain:     "example.com",
			recordType: "CNAME",
			txtDomain:  "cnamefoo.example.com",
		},
		{
			name:       "templated suffix with dot",
			mapper:     newaffixNameMapper("", ".foo%{record_type}", ""),
			domain:     "example.com",
			recordType: "A",
			txtDomain:  "example.fooa.com",
		},
		{
			name:       "templated prefix with multiple dots",
			mapper:     newaffixNameMapper("bar.%{record_type}.foo.", "", ""),
			domain:     "example.com",
			recordType: "CNAME",
			txtDomain:  "bar.cname.foo.example.com",
		},
		{
			name:       "templated suffix with multiple dots",
			mapper:     newaffixNameMapper("", ".foo%{record_type}.bar", ""),
			domain:     "example.com",
			recordType: "A",
			txtDomain:  "example.fooa.bar.com",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			txtDomain := tc.mapper.toNewTXTName(tc.domain, tc.recordType)
			assert.Equal(t, tc.txtDomain, txtDomain)

			domain, _ := tc.mapper.toEndpointName(txtDomain)
			assert.Equal(t, tc.domain, domain)
		})
	}
}

func TestNewTXTScheme(t *testing.T) {
	p := inmemory.NewInMemoryProvider()
	p.CreateZone(testZone)
	ctxEndpoints := []*endpoint.Endpoint{}
	ctx := context.WithValue(context.Background(), provider.RecordsContextKey, ctxEndpoints)
	p.OnApplyChanges = func(ctx context.Context, got *plan.Changes) {
		assert.Equal(t, ctxEndpoints, ctx.Value(provider.RecordsContextKey))
	}
	p.ApplyChanges(ctx, &plan.Changes{
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
			newEndpointWithOwner("cname-foobar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
		},
	})
	r, _ := NewTXTRegistry(p, "", "", "owner", time.Hour, "", []string{}, []string{}, false, nil)

	changes := &plan.Changes{
		Create: []*endpoint.Endpoint{
			newEndpointWithOwner("new-record-1.test-zone.example.org", "new-loadbalancer-1.lb.com", endpoint.RecordTypeCNAME, ""),
			newEndpointWithOwner("example", "new-loadbalancer-1.lb.com", endpoint.RecordTypeCNAME, ""),
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
			newEndpointWithOwner("new-record-1.test-zone.example.org", "new-loadbalancer-1.lb.com", endpoint.RecordTypeCNAME, "owner"),
			newEndpointWithOwnerAndOwnedRecord("new-record-1.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, "", "new-record-1.test-zone.example.org"),
			newEndpointWithOwnerAndOwnedRecord("cname-new-record-1.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, "", "new-record-1.test-zone.example.org"),
			newEndpointWithOwner("example", "new-loadbalancer-1.lb.com", endpoint.RecordTypeCNAME, "owner"),
			newEndpointWithOwnerAndOwnedRecord("example", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, "", "example"),
			newEndpointWithOwnerAndOwnedRecord("cname-example", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, "", "example"),
		},
		Delete: []*endpoint.Endpoint{
			newEndpointWithOwner("foobar.test-zone.example.org", "foobar.loadbalancer.com", endpoint.RecordTypeCNAME, "owner"),
			newEndpointWithOwnerAndOwnedRecord("foobar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, "", "foobar.test-zone.example.org"),
			newEndpointWithOwnerAndOwnedRecord("cname-foobar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, "", "foobar.test-zone.example.org"),
		},
		UpdateNew: []*endpoint.Endpoint{},
		UpdateOld: []*endpoint.Endpoint{},
	}
	p.OnApplyChanges = func(ctx context.Context, got *plan.Changes) {
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
		assert.Equal(t, nil, ctx.Value(provider.RecordsContextKey))
	}
	err := r.ApplyChanges(ctx, changes)
	require.NoError(t, err)
}

func TestGenerateTXT(t *testing.T) {
	record := newEndpointWithOwner("foo.test-zone.example.org", "new-foo.loadbalancer.com", endpoint.RecordTypeCNAME, "owner")
	expectedTXT := []*endpoint.Endpoint{
		{
			DNSName:    "foo.test-zone.example.org",
			Targets:    endpoint.Targets{"\"heritage=external-dns,external-dns/owner=owner\""},
			RecordType: endpoint.RecordTypeTXT,
			Labels: map[string]string{
				endpoint.OwnedRecordLabelKey: "foo.test-zone.example.org",
			},
		},
		{
			DNSName:    "cname-foo.test-zone.example.org",
			Targets:    endpoint.Targets{"\"heritage=external-dns,external-dns/owner=owner\""},
			RecordType: endpoint.RecordTypeTXT,
			Labels: map[string]string{
				endpoint.OwnedRecordLabelKey: "foo.test-zone.example.org",
			},
		},
	}
	p := inmemory.NewInMemoryProvider()
	p.CreateZone(testZone)
	r, _ := NewTXTRegistry(p, "", "", "owner", time.Hour, "", []string{}, []string{}, false, nil)
	gotTXT := r.generateTXTRecord(record)
	assert.Equal(t, expectedTXT, gotTXT)
}

func TestGenerateTXTForAAAA(t *testing.T) {
	record := newEndpointWithOwner("foo.test-zone.example.org", "2001:DB8::1", endpoint.RecordTypeAAAA, "owner")
	expectedTXT := []*endpoint.Endpoint{
		{
			DNSName:    "aaaa-foo.test-zone.example.org",
			Targets:    endpoint.Targets{"\"heritage=external-dns,external-dns/owner=owner\""},
			RecordType: endpoint.RecordTypeTXT,
			Labels: map[string]string{
				endpoint.OwnedRecordLabelKey: "foo.test-zone.example.org",
			},
		},
	}
	p := inmemory.NewInMemoryProvider()
	p.CreateZone(testZone)
	r, _ := NewTXTRegistry(p, "", "", "owner", time.Hour, "", []string{}, []string{}, false, nil)
	gotTXT := r.generateTXTRecord(record)
	assert.Equal(t, expectedTXT, gotTXT)
}

func TestFailGenerateTXT(t *testing.T) {

	cnameRecord := &endpoint.Endpoint{
		DNSName:    "foo-some-really-big-name-not-supported-and-will-fail-000000000000000000.test-zone.example.org",
		Targets:    endpoint.Targets{"new-foo.loadbalancer.com"},
		RecordType: endpoint.RecordTypeCNAME,
		Labels:     map[string]string{},
	}
	// A bad DNS name returns empty expected TXT
	expectedTXT := []*endpoint.Endpoint{}
	p := inmemory.NewInMemoryProvider()
	p.CreateZone(testZone)
	r, _ := NewTXTRegistry(p, "", "", "owner", time.Hour, "", []string{}, []string{}, false, nil)
	gotTXT := r.generateTXTRecord(cnameRecord)
	assert.Equal(t, expectedTXT, gotTXT)
}

func TestTXTRegistryApplyChangesEncrypt(t *testing.T) {
	p := inmemory.NewInMemoryProvider()
	p.CreateZone(testZone)
	ctxEndpoints := []*endpoint.Endpoint{}
	ctx := context.WithValue(context.Background(), provider.RecordsContextKey, ctxEndpoints)

	p.ApplyChanges(ctx, &plan.Changes{
		Create: []*endpoint.Endpoint{
			newEndpointWithOwner("foobar.test-zone.example.org", "foobar.loadbalancer.com", endpoint.RecordTypeCNAME, ""),
			newEndpointWithOwnerAndOwnedRecord("txt.cname-foobar.test-zone.example.org", "\"h8UQ6jelUFUsEIn7SbFktc2MYXPx/q8lySqI4VwfVtVaIbb2nkHWV/88KKbuLtu7fJNzMir8ELVeVnRSY01KdiIuj7ledqZe5ailEjQaU5Z6uEKd5pgs6sH8\"", endpoint.RecordTypeTXT, "", "foobar.test-zone.example.org"),
		},
	})

	r, _ := NewTXTRegistry(p, "txt.", "", "owner", time.Hour, "", []string{}, []string{}, true, []byte("12345678901234567890123456789012"))
	records, _ := r.Records(ctx)
	changes := &plan.Changes{
		Delete: records,
	}

	// ensure that encryption nonce gets reused when deleting records
	expected := &plan.Changes{
		Delete: []*endpoint.Endpoint{
			newEndpointWithOwner("foobar.test-zone.example.org", "foobar.loadbalancer.com", endpoint.RecordTypeCNAME, "owner"),
			newEndpointWithOwnerAndOwnedRecord("txt.cname-foobar.test-zone.example.org", "\"h8UQ6jelUFUsEIn7SbFktc2MYXPx/q8lySqI4VwfVtVaIbb2nkHWV/88KKbuLtu7fJNzMir8ELVeVnRSY01KdiIuj7ledqZe5ailEjQaU5Z6uEKd5pgs6sH8\"", endpoint.RecordTypeTXT, "", "foobar.test-zone.example.org"),
		},
	}

	p.OnApplyChanges = func(ctx context.Context, got *plan.Changes) {
		mExpected := map[string][]*endpoint.Endpoint{
			"Delete": expected.Delete,
		}
		mGot := map[string][]*endpoint.Endpoint{
			"Delete": got.Delete,
		}
		assert.True(t, testutils.SamePlanChanges(mGot, mExpected))
		assert.Equal(t, nil, ctx.Value(provider.RecordsContextKey))
	}
	err := r.ApplyChanges(ctx, changes)
	require.NoError(t, err)
}

// TestMultiClusterDifferentRecordTypeOwnership validates the registry handles environments where the same zone is managed by
// external-dns in different clusters and the ingress record type is different. For example one uses A records and the other
// uses CNAME. In this environment the first cluster that establishes the owner record should maintain ownership even
// if the same ingress host is deployed to the other. With the introduction of Dual Record support each record type
// was treated independently and would cause each cluster to fight over ownership. This tests ensure that the default
// Dual Stack record support only treats AAAA records independently and while keeping A and CNAME record ownership intact.
func TestMultiClusterDifferentRecordTypeOwnership(t *testing.T) {
	ctx := context.Background()
	p := inmemory.NewInMemoryProvider()
	p.CreateZone(testZone)
	p.ApplyChanges(ctx, &plan.Changes{
		Create: []*endpoint.Endpoint{
			// records on cluster using A record for ingress address
			newEndpointWithOwner("bar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=cat,external-dns/resource=ingress/default/foo\"", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("bar.test-zone.example.org", "1.2.3.4", endpoint.RecordTypeA, ""),
		},
	})

	r, _ := NewTXTRegistry(p, "_owner.", "", "bar", time.Hour, "", []string{}, []string{}, false, nil)
	records, _ := r.Records(ctx)

	// new cluster has same ingress host as other cluster and uses CNAME ingress address
	cname := &endpoint.Endpoint{
		DNSName:    "bar.test-zone.example.org",
		Targets:    endpoint.Targets{"cluster-b"},
		RecordType: "CNAME",
		Labels: map[string]string{
			endpoint.ResourceLabelKey: "ingress/default/foo-127",
		},
	}
	desired := []*endpoint.Endpoint{cname}

	pl := &plan.Plan{
		Policies:       []plan.Policy{&plan.SyncPolicy{}},
		Current:        records,
		Desired:        desired,
		ManagedRecords: []string{endpoint.RecordTypeA, endpoint.RecordTypeAAAA, endpoint.RecordTypeCNAME},
	}

	changes := pl.Calculate()
	p.OnApplyChanges = func(ctx context.Context, changes *plan.Changes) {
		got := map[string][]*endpoint.Endpoint{
			"Create":    changes.Create,
			"UpdateNew": changes.UpdateNew,
			"UpdateOld": changes.UpdateOld,
			"Delete":    changes.Delete,
		}
		expected := map[string][]*endpoint.Endpoint{
			"Create":    {},
			"UpdateNew": {},
			"UpdateOld": {},
			"Delete":    {},
		}
		testutils.SamePlanChanges(got, expected)
	}

	err := r.ApplyChanges(ctx, changes.Changes)
	if err != nil {
		t.Error(err)
	}
}

/**

helper methods

*/

func newEndpointWithOwner(dnsName, target, recordType, ownerID string) *endpoint.Endpoint {
	return newEndpointWithOwnerAndLabels(dnsName, target, recordType, ownerID, nil)
}

func newEndpointWithOwnerAndOwnedRecord(dnsName, target, recordType, ownerID, ownedRecord string) *endpoint.Endpoint {
	return newEndpointWithOwnerAndLabels(dnsName, target, recordType, ownerID, endpoint.Labels{endpoint.OwnedRecordLabelKey: ownedRecord})
}

func newEndpointWithOwnerAndLabels(dnsName, target, recordType, ownerID string, labels endpoint.Labels) *endpoint.Endpoint {
	e := endpoint.NewEndpoint(dnsName, recordType, target)
	e.Labels[endpoint.OwnerLabelKey] = ownerID
	for k, v := range labels {
		e.Labels[k] = v
	}
	return e
}

func newEndpointWithOwnerResource(dnsName, target, recordType, ownerID, resource string) *endpoint.Endpoint {
	e := endpoint.NewEndpoint(dnsName, recordType, target)
	e.Labels[endpoint.OwnerLabelKey] = ownerID
	e.Labels[endpoint.ResourceLabelKey] = resource
	return e
}
