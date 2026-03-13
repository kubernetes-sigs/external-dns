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

package txt

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"sigs.k8s.io/external-dns/registry/mapper"

	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	logtest "sigs.k8s.io/external-dns/internal/testutils/log"
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
	_, err := NewTXTRegistry(p, "txt", "", "", time.Hour, "", []string{}, []string{}, false, nil, "")
	require.Error(t, err)

	_, err = NewTXTRegistry(p, "", "txt", "", time.Hour, "", []string{}, []string{}, false, nil, "")
	require.Error(t, err)

	r, err := NewTXTRegistry(p, "txt", "", "owner", time.Hour, "", []string{}, []string{}, false, nil, "")
	require.NoError(t, err)
	assert.Equal(t, p, r.provider)

	r, err = NewTXTRegistry(p, "", "txt", "owner", time.Hour, "", []string{}, []string{}, false, nil, "")
	require.NoError(t, err)

	_, err = NewTXTRegistry(p, "txt", "txt", "owner", time.Hour, "", []string{}, []string{}, false, nil, "")
	require.Error(t, err)

	_, ok := r.mapper.(mapper.AffixNameMapper)
	require.True(t, ok)
	assert.Equal(t, "owner", r.ownerID)
	assert.Equal(t, p, r.provider)

	aesKey := []byte(";k&l)nUC/33:{?d{3)54+,AD?]SX%yh^")
	_, err = NewTXTRegistry(p, "", "", "owner", time.Hour, "", []string{}, []string{}, false, nil, "")
	require.NoError(t, err)

	_, err = NewTXTRegistry(p, "", "", "owner", time.Hour, "", []string{}, []string{}, false, aesKey, "")
	require.NoError(t, err)

	_, err = NewTXTRegistry(p, "", "", "owner", time.Hour, "", []string{}, []string{}, true, nil, "")
	require.Error(t, err)

	r, err = NewTXTRegistry(p, "", "", "owner", time.Hour, "", []string{}, []string{}, true, aesKey, "")
	require.NoError(t, err)

	_, ok = r.mapper.(mapper.AffixNameMapper)
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
	err := p.CreateZone(testZone)
	require.NoError(t, err)
	err = p.ApplyChanges(ctx, &plan.Changes{
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
			newEndpointWithOwner("mail.test-zone.example.org", "10 onemail.example.com", endpoint.RecordTypeMX, ""),
			newEndpointWithOwner("txt.mx-mail.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
			newMultiTargetEndpointWithOwner(
				"_sip._udp.sip1.test-zone.example.org",
				[]string{"1 50 5060 sip1-n1.test-zone.example.org", "1 50 5060 sip1-n2.test-zone.example.org"},
				endpoint.RecordTypeSRV,
				"",
			),
			newEndpointWithOwner("txt._sip._udp.sip1.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("sip1.test-zone.example.org", `10 "U" "SIP+DTU" "" _sip._udp.sip1.test-zone.example.org.`, endpoint.RecordTypeNAPTR, ""),
			newEndpointWithOwner("txt.sip1.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
		},
	})
	require.NoError(t, err)
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
		{
			DNSName:    "mail.test-zone.example.org",
			Targets:    endpoint.Targets{"10 onemail.example.com"},
			RecordType: endpoint.RecordTypeMX,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "owner",
			},
		},
		{
			DNSName: "_sip._udp.sip1.test-zone.example.org",
			Targets: endpoint.Targets{
				"1 50 5060 sip1-n1.test-zone.example.org",
				"1 50 5060 sip1-n2.test-zone.example.org",
			},
			RecordType: endpoint.RecordTypeSRV,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "owner",
			},
		},
		{
			DNSName:    "sip1.test-zone.example.org",
			Targets:    endpoint.Targets{`10 "U" "SIP+DTU" "" _sip._udp.sip1.test-zone.example.org.`},
			RecordType: endpoint.RecordTypeNAPTR,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "owner",
			},
		},
	}

	r, _ := NewTXTRegistry(p, "txt.", "", "owner", time.Hour, "wc", []string{}, []string{}, false, nil, "")
	records, _ := r.Records(ctx)

	assert.True(t, testutils.SameEndpoints(records, expectedRecords))

	// Ensure prefix is case-insensitive
	r, _ = NewTXTRegistry(p, "TxT.", "", "owner", time.Hour, "wc", []string{}, []string{}, false, nil, "")
	records, _ = r.Records(ctx)

	assert.True(t, testutils.SameEndpoints(records, expectedRecords))
}

func testTXTRegistryRecordsSuffixed(t *testing.T) {
	ctx := context.Background()
	p := inmemory.NewInMemoryProvider()
	err := p.CreateZone(testZone)
	require.NoError(t, err)
	err = p.ApplyChanges(ctx, &plan.Changes{
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
			newEndpointWithOwner("mail.test-zone.example.org", "10 onemail.example.com", endpoint.RecordTypeMX, ""),
			newEndpointWithOwner("mx-mail-txt.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
			newMultiTargetEndpointWithOwner(
				"_sip._udp.sip1.test-zone.example.org",
				[]string{"1 50 5060 sip1-n1.test-zone.example.org", "1 50 5060 sip1-n2.test-zone.example.org"},
				endpoint.RecordTypeSRV,
				"",
			),
			newEndpointWithOwner("_sip-txt._udp.sip1.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("sip1.test-zone.example.org", `10 "U" "SIP+DTU" "" _sip._udp.sip1.test-zone.example.org.`, endpoint.RecordTypeNAPTR, ""),
			newEndpointWithOwner("sip1-txt.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
		},
	})
	require.NoError(t, err)
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
		{
			DNSName:    "mail.test-zone.example.org",
			Targets:    endpoint.Targets{"10 onemail.example.com"},
			RecordType: endpoint.RecordTypeMX,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "owner",
			},
		},
		{
			DNSName: "_sip._udp.sip1.test-zone.example.org",
			Targets: endpoint.Targets{
				"1 50 5060 sip1-n1.test-zone.example.org",
				"1 50 5060 sip1-n2.test-zone.example.org",
			},
			RecordType: endpoint.RecordTypeSRV,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "owner",
			},
		},
		{
			DNSName:    "sip1.test-zone.example.org",
			Targets:    endpoint.Targets{`10 "U" "SIP+DTU" "" _sip._udp.sip1.test-zone.example.org.`},
			RecordType: endpoint.RecordTypeNAPTR,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "owner",
			},
		},
	}

	r, _ := NewTXTRegistry(p, "", "-txt", "owner", time.Hour, "", []string{}, []string{}, false, nil, "")
	records, _ := r.Records(ctx)

	assert.True(t, testutils.SameEndpoints(records, expectedRecords))

	// Ensure prefix is case-insensitive
	r, _ = NewTXTRegistry(p, "", "-TxT", "owner", time.Hour, "", []string{}, []string{}, false, nil, "")
	records, _ = r.Records(ctx)

	assert.True(t, testutils.SameEndpointLabels(records, expectedRecords))
}

func testTXTRegistryRecordsNoPrefix(t *testing.T) {
	p := inmemory.NewInMemoryProvider()
	ctx := context.Background()
	err := p.CreateZone(testZone)
	require.NoError(t, err)
	err = p.ApplyChanges(ctx, &plan.Changes{
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
			newEndpointWithOwner("mail.test-zone.example.org", "10 onemail.example.com", endpoint.RecordTypeMX, ""),
			newEndpointWithOwner("mx-mail.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
			newMultiTargetEndpointWithOwner(
				"_sip._udp.sip1.test-zone.example.org",
				[]string{"1 50 5060 sip1-n1.test-zone.example.org", "1 50 5060 sip1-n2.test-zone.example.org"},
				endpoint.RecordTypeSRV,
				"",
			),
			newEndpointWithOwner("_sip._udp.sip1.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("sip1.test-zone.example.org", `10 "U" "SIP+DTU" "" _sip._udp.sip1.test-zone.example.org.`, endpoint.RecordTypeNAPTR, ""),
			newEndpointWithOwner("sip1.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
		},
	})
	require.NoError(t, err)
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
		{
			DNSName:    "mail.test-zone.example.org",
			Targets:    endpoint.Targets{"10 onemail.example.com"},
			RecordType: endpoint.RecordTypeMX,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "owner",
			},
		},
		{
			DNSName: "_sip._udp.sip1.test-zone.example.org",
			Targets: endpoint.Targets{
				"1 50 5060 sip1-n1.test-zone.example.org",
				"1 50 5060 sip1-n2.test-zone.example.org",
			},
			RecordType: endpoint.RecordTypeSRV,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "owner",
			},
		},
		{
			DNSName:    "sip1.test-zone.example.org",
			Targets:    endpoint.Targets{`10 "U" "SIP+DTU" "" _sip._udp.sip1.test-zone.example.org.`},
			RecordType: endpoint.RecordTypeNAPTR,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "owner",
			},
		},
	}

	r, _ := NewTXTRegistry(p, "", "", "owner", time.Hour, "", []string{}, []string{}, false, nil, "")
	records, _ := r.Records(ctx)

	assert.True(t, testutils.SameEndpoints(records, expectedRecords))
}

func testTXTRegistryRecordsPrefixedTemplated(t *testing.T) {
	ctx := context.Background()
	p := inmemory.NewInMemoryProvider()
	err := p.CreateZone(testZone)
	require.NoError(t, err)
	err = p.ApplyChanges(ctx, &plan.Changes{
		Create: []*endpoint.Endpoint{
			newEndpointWithOwner("foo.test-zone.example.org", "1.1.1.1", endpoint.RecordTypeA, ""),
			newEndpointWithOwner("txt-a.foo.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("mail.test-zone.example.org", "10 onemail.example.com", endpoint.RecordTypeMX, ""),
			newEndpointWithOwner("txt-mx.mail.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
		},
	})
	require.NoError(t, err)
	expectedRecords := []*endpoint.Endpoint{
		{
			DNSName:    "foo.test-zone.example.org",
			Targets:    endpoint.Targets{"1.1.1.1"},
			RecordType: endpoint.RecordTypeA,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "owner",
			},
		},
		{
			DNSName:    "mail.test-zone.example.org",
			Targets:    endpoint.Targets{"10 onemail.example.com"},
			RecordType: endpoint.RecordTypeMX,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "owner",
			},
		},
	}

	r, _ := NewTXTRegistry(p, "txt-%{record_type}.", "", "owner", time.Hour, "wc", []string{}, []string{}, false, nil, "")
	records, _ := r.Records(ctx)

	assert.True(t, testutils.SameEndpoints(records, expectedRecords))

	r, _ = NewTXTRegistry(p, "TxT-%{record_type}.", "", "owner", time.Hour, "wc", []string{}, []string{}, false, nil, "")
	records, _ = r.Records(ctx)

	assert.True(t, testutils.SameEndpoints(records, expectedRecords))
}

func testTXTRegistryRecordsSuffixedTemplated(t *testing.T) {
	ctx := context.Background()
	p := inmemory.NewInMemoryProvider()
	err := p.CreateZone(testZone)
	require.NoError(t, err)
	err = p.ApplyChanges(ctx, &plan.Changes{
		Create: []*endpoint.Endpoint{
			newEndpointWithOwner("bar.test-zone.example.org", "8.8.8.8", endpoint.RecordTypeCNAME, ""),
			newEndpointWithOwner("bartxtcname.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("mail.test-zone.example.org", "10 onemail.example.com", endpoint.RecordTypeMX, ""),
			newEndpointWithOwner("mailtxt.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", endpoint.RecordTypeTXT, ""),
		},
	})
	require.NoError(t, err)
	expectedRecords := []*endpoint.Endpoint{
		{
			DNSName:    "bar.test-zone.example.org",
			Targets:    endpoint.Targets{"8.8.8.8"},
			RecordType: endpoint.RecordTypeCNAME,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "owner",
			},
		},
		{
			DNSName:    "mail.test-zone.example.org",
			Targets:    endpoint.Targets{"10 onemail.example.com"},
			RecordType: endpoint.RecordTypeMX,
			Labels: map[string]string{
				endpoint.OwnerLabelKey: "owner",
			},
		},
	}

	r, _ := NewTXTRegistry(p, "", "txt%{record_type}", "owner", time.Hour, "wc", []string{}, []string{}, false, nil, "")
	records, _ := r.Records(ctx)

	assert.True(t, testutils.SameEndpoints(records, expectedRecords))

	r, _ = NewTXTRegistry(p, "", "TxT%{record_type}", "owner", time.Hour, "wc", []string{}, []string{}, false, nil, "")
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
	_ = p.CreateZone(testZone)
	var ctxEndpoints []*endpoint.Endpoint
	ctx := context.WithValue(context.Background(), provider.RecordsContextKey, ctxEndpoints)
	p.OnApplyChanges = func(ctx context.Context, _ *plan.Changes) {
		assert.Equal(t, ctxEndpoints, ctx.Value(provider.RecordsContextKey))
	}
	_ = p.ApplyChanges(ctx, &plan.Changes{
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
	r, _ := NewTXTRegistry(p, "txt.", "", "owner", time.Hour, "", []string{}, []string{}, false, nil, "")

	changes := &plan.Changes{
		Create: []*endpoint.Endpoint{
			newCNAMEEndpointWithOwnerResource("new-record-1.test-zone.example.org", "new-loadbalancer-1.lb.com", "owner", "ingress/default/my-ingress"),
			newCNAMEEndpointWithOwnerResource("multiple.test-zone.example.org", "lb3.loadbalancer.com", "owner", "ingress/default/my-ingress").WithSetIdentifier("test-set-3"),
			newCNAMEEndpointWithOwnerResource("example", "new-loadbalancer-1.lb.com", "owner", "ingress/default/my-ingress"),
		},
		Delete: []*endpoint.Endpoint{
			newEndpointWithOwner("foobar.test-zone.example.org", "foobar.loadbalancer.com", endpoint.RecordTypeCNAME, "owner"),
			newEndpointWithOwner("multiple.test-zone.example.org", "lb1.loadbalancer.com", endpoint.RecordTypeCNAME, "owner").WithSetIdentifier("test-set-1"),
		},
		UpdateNew: []*endpoint.Endpoint{
			newCNAMEEndpointWithOwnerResource("tar.test-zone.example.org", "new-tar.loadbalancer.com", "owner", "ingress/default/my-ingress-2"),
			newCNAMEEndpointWithOwnerResource("multiple.test-zone.example.org", "new.loadbalancer.com", "owner", "ingress/default/my-ingress-2").WithSetIdentifier("test-set-2"),
		},
		UpdateOld: []*endpoint.Endpoint{
			newEndpointWithOwner("tar.test-zone.example.org", "tar.loadbalancer.com", endpoint.RecordTypeCNAME, "owner"),
			newEndpointWithOwner("multiple.test-zone.example.org", "lb2.loadbalancer.com", endpoint.RecordTypeCNAME, "owner").WithSetIdentifier("test-set-2"),
		},
	}
	expected := &plan.Changes{
		Create: []*endpoint.Endpoint{
			newCNAMEEndpointWithOwnerResource("new-record-1.test-zone.example.org", "new-loadbalancer-1.lb.com", "owner", "ingress/default/my-ingress"),
			newTXTEndpointWithOwnedRecord("txt.cname-new-record-1.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner,external-dns/resource=ingress/default/my-ingress\"", "new-record-1.test-zone.example.org"),
			newCNAMEEndpointWithOwnerResource("multiple.test-zone.example.org", "lb3.loadbalancer.com", "owner", "ingress/default/my-ingress").WithSetIdentifier("test-set-3"),
			newTXTEndpointWithOwnedRecord("txt.cname-multiple.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner,external-dns/resource=ingress/default/my-ingress\"", "multiple.test-zone.example.org").WithSetIdentifier("test-set-3"),
			newCNAMEEndpointWithOwnerResource("example", "new-loadbalancer-1.lb.com", "owner", "ingress/default/my-ingress"),
			newTXTEndpointWithOwnedRecord("txt.cname-example", "\"heritage=external-dns,external-dns/owner=owner,external-dns/resource=ingress/default/my-ingress\"", "example"),
		},
		Delete: []*endpoint.Endpoint{
			newEndpointWithOwner("foobar.test-zone.example.org", "foobar.loadbalancer.com", endpoint.RecordTypeCNAME, "owner"),
			newTXTEndpointWithOwnedRecord("txt.cname-foobar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", "foobar.test-zone.example.org"),
			newEndpointWithOwner("multiple.test-zone.example.org", "lb1.loadbalancer.com", endpoint.RecordTypeCNAME, "owner").WithSetIdentifier("test-set-1"),
			newTXTEndpointWithOwnedRecord("txt.cname-multiple.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", "multiple.test-zone.example.org").WithSetIdentifier("test-set-1"),
		},
		UpdateNew: []*endpoint.Endpoint{
			newCNAMEEndpointWithOwnerResource("tar.test-zone.example.org", "new-tar.loadbalancer.com", "owner", "ingress/default/my-ingress-2"),
			newTXTEndpointWithOwnedRecord("txt.cname-tar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner,external-dns/resource=ingress/default/my-ingress-2\"", "tar.test-zone.example.org"),
			newCNAMEEndpointWithOwnerResource("multiple.test-zone.example.org", "new.loadbalancer.com", "owner", "ingress/default/my-ingress-2").WithSetIdentifier("test-set-2"),
			newTXTEndpointWithOwnedRecord("txt.cname-multiple.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner,external-dns/resource=ingress/default/my-ingress-2\"", "multiple.test-zone.example.org").WithSetIdentifier("test-set-2"),
		},
		UpdateOld: []*endpoint.Endpoint{
			newEndpointWithOwner("tar.test-zone.example.org", "tar.loadbalancer.com", endpoint.RecordTypeCNAME, "owner"),
			newTXTEndpointWithOwnedRecord("txt.cname-tar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", "tar.test-zone.example.org"),
			newEndpointWithOwner("multiple.test-zone.example.org", "lb2.loadbalancer.com", endpoint.RecordTypeCNAME, "owner").WithSetIdentifier("test-set-2"),
			newTXTEndpointWithOwnedRecord("txt.cname-multiple.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", "multiple.test-zone.example.org").WithSetIdentifier("test-set-2"),
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
		assert.Nil(t, ctx.Value(provider.RecordsContextKey))
	}
	err := r.ApplyChanges(ctx, changes)
	require.NoError(t, err)
}

func testTXTRegistryApplyChangesWithTemplatedPrefix(t *testing.T) {
	p := inmemory.NewInMemoryProvider()
	err := p.CreateZone(testZone)
	require.NoError(t, err)
	var ctxEndpoints []*endpoint.Endpoint
	ctx := context.WithValue(context.Background(), provider.RecordsContextKey, ctxEndpoints)
	p.OnApplyChanges = func(ctx context.Context, _ *plan.Changes) {
		assert.Equal(t, ctxEndpoints, ctx.Value(provider.RecordsContextKey))
	}
	_ = p.ApplyChanges(ctx, &plan.Changes{
		Create: []*endpoint.Endpoint{},
	})
	r, _ := NewTXTRegistry(p, "prefix%{record_type}.", "", "owner-1", time.Hour, "", []string{}, []string{}, false, nil, "")
	changes := &plan.Changes{
		Create: []*endpoint.Endpoint{
			newCNAMEEndpointWithOwnerResource("new-record-1.test-zone.example.org", "new-loadbalancer-1.lb.com", "owner-1", "ingress/default/my-ingress"),
		},
		Delete:    []*endpoint.Endpoint{},
		UpdateOld: []*endpoint.Endpoint{},
		UpdateNew: []*endpoint.Endpoint{},
	}
	expected := &plan.Changes{
		Create: []*endpoint.Endpoint{
			newCNAMEEndpointWithOwnerResource("new-record-1.test-zone.example.org", "new-loadbalancer-1.lb.com", "owner-1", "ingress/default/my-ingress"),
			newTXTEndpointWithOwnedRecord("prefixcname.new-record-1.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner-1,external-dns/resource=ingress/default/my-ingress\"", "new-record-1.test-zone.example.org"),
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
		assert.Nil(t, ctx.Value(provider.RecordsContextKey))
	}
	err = r.ApplyChanges(ctx, changes)
	require.NoError(t, err)
}

func testTXTRegistryApplyChangesWithTemplatedSuffix(t *testing.T) {
	p := inmemory.NewInMemoryProvider()
	_ = p.CreateZone(testZone)
	var ctxEndpoints []*endpoint.Endpoint
	ctx := context.WithValue(context.Background(), provider.RecordsContextKey, ctxEndpoints)
	p.OnApplyChanges = func(ctx context.Context, _ *plan.Changes) {
		assert.Equal(t, ctxEndpoints, ctx.Value(provider.RecordsContextKey))
	}
	r, _ := NewTXTRegistry(p, "", "-%{record_type}suffix", "owner-2", time.Hour, "", []string{}, []string{}, false, nil, "")
	changes := &plan.Changes{
		Create: []*endpoint.Endpoint{
			newCNAMEEndpointWithOwnerResource("new-record-1.test-zone.example.org", "new-loadbalancer-1.lb.com", "owner-2", "ingress/default/my-ingress"),
		},
		Delete:    []*endpoint.Endpoint{},
		UpdateOld: []*endpoint.Endpoint{},
		UpdateNew: []*endpoint.Endpoint{},
	}
	expected := &plan.Changes{
		Create: []*endpoint.Endpoint{
			newCNAMEEndpointWithOwnerResource("new-record-1.test-zone.example.org", "new-loadbalancer-1.lb.com", "owner-2", "ingress/default/my-ingress"),
			newTXTEndpointWithOwnedRecord("new-record-1-cnamesuffix.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner-2,external-dns/resource=ingress/default/my-ingress\"", "new-record-1.test-zone.example.org"),
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
		assert.Nil(t, ctx.Value(provider.RecordsContextKey))
	}
	err := r.ApplyChanges(ctx, changes)
	require.NoError(t, err)
}

func testTXTRegistryApplyChangesWithSuffix(t *testing.T) {
	p := inmemory.NewInMemoryProvider()
	err := p.CreateZone(testZone)
	require.NoError(t, err)
	var ctxEndpoints []*endpoint.Endpoint
	ctx := context.WithValue(context.Background(), provider.RecordsContextKey, ctxEndpoints)
	p.OnApplyChanges = func(ctx context.Context, _ *plan.Changes) {
		assert.Equal(t, ctxEndpoints, ctx.Value(provider.RecordsContextKey))
	}
	err = p.ApplyChanges(ctx, &plan.Changes{
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
	require.NoError(t, err)
	r, _ := NewTXTRegistry(p, "", "-txt", "owner", time.Hour, "wildcard", []string{}, []string{}, false, nil, "")

	changes := &plan.Changes{
		Create: []*endpoint.Endpoint{
			newCNAMEEndpointWithOwnerResource("new-record-1.test-zone.example.org", "new-loadbalancer-1.lb.com", "owner", "ingress/default/my-ingress"),
			newCNAMEEndpointWithOwnerResource("multiple.test-zone.example.org", "lb3.loadbalancer.com", "owner", "ingress/default/my-ingress").WithSetIdentifier("test-set-3"),
			newCNAMEEndpointWithOwnerResource("example", "new-loadbalancer-1.lb.com", "owner", "ingress/default/my-ingress"),
			newCNAMEEndpointWithOwnerResource("*.wildcard.test-zone.example.org", "new-loadbalancer-1.lb.com", "owner", "ingress/default/my-ingress"),
		},
		Delete: []*endpoint.Endpoint{
			newEndpointWithOwner("foobar.test-zone.example.org", "foobar.loadbalancer.com", endpoint.RecordTypeCNAME, "owner"),
			newEndpointWithOwner("multiple.test-zone.example.org", "lb1.loadbalancer.com", endpoint.RecordTypeCNAME, "owner").WithSetIdentifier("test-set-1"),
		},
		UpdateNew: []*endpoint.Endpoint{
			newCNAMEEndpointWithOwnerResource("tar.test-zone.example.org", "new-tar.loadbalancer.com", "owner", "ingress/default/my-ingress-2"),
			newCNAMEEndpointWithOwnerResource("multiple.test-zone.example.org", "new.loadbalancer.com", "owner", "ingress/default/my-ingress-2").WithSetIdentifier("test-set-2"),
		},
		UpdateOld: []*endpoint.Endpoint{
			newEndpointWithOwner("tar.test-zone.example.org", "tar.loadbalancer.com", endpoint.RecordTypeCNAME, "owner"),
			newEndpointWithOwner("multiple.test-zone.example.org", "lb2.loadbalancer.com", endpoint.RecordTypeCNAME, "owner").WithSetIdentifier("test-set-2"),
		},
	}
	expected := &plan.Changes{
		Create: []*endpoint.Endpoint{
			newCNAMEEndpointWithOwnerResource("new-record-1.test-zone.example.org", "new-loadbalancer-1.lb.com", "owner", "ingress/default/my-ingress"),
			newTXTEndpointWithOwnedRecord("cname-new-record-1-txt.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner,external-dns/resource=ingress/default/my-ingress\"", "new-record-1.test-zone.example.org"),
			newCNAMEEndpointWithOwnerResource("multiple.test-zone.example.org", "lb3.loadbalancer.com", "owner", "ingress/default/my-ingress").WithSetIdentifier("test-set-3"),
			newTXTEndpointWithOwnedRecord("cname-multiple-txt.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner,external-dns/resource=ingress/default/my-ingress\"", "multiple.test-zone.example.org").WithSetIdentifier("test-set-3"),
			newCNAMEEndpointWithOwnerResource("example", "new-loadbalancer-1.lb.com", "owner", "ingress/default/my-ingress"),
			newTXTEndpointWithOwnedRecord("cname-example-txt", "\"heritage=external-dns,external-dns/owner=owner,external-dns/resource=ingress/default/my-ingress\"", "example"),
			newCNAMEEndpointWithOwnerResource("*.wildcard.test-zone.example.org", "new-loadbalancer-1.lb.com", "owner", "ingress/default/my-ingress"),
			newTXTEndpointWithOwnedRecord("cname-wildcard-txt.wildcard.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner,external-dns/resource=ingress/default/my-ingress\"", "*.wildcard.test-zone.example.org"),
		},
		Delete: []*endpoint.Endpoint{
			newEndpointWithOwner("foobar.test-zone.example.org", "foobar.loadbalancer.com", endpoint.RecordTypeCNAME, "owner"),
			newTXTEndpointWithOwnedRecord("cname-foobar-txt.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", "foobar.test-zone.example.org"),
			newEndpointWithOwner("multiple.test-zone.example.org", "lb1.loadbalancer.com", endpoint.RecordTypeCNAME, "owner").WithSetIdentifier("test-set-1"),
			newTXTEndpointWithOwnedRecord("cname-multiple-txt.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", "multiple.test-zone.example.org").WithSetIdentifier("test-set-1"),
		},
		UpdateNew: []*endpoint.Endpoint{
			newCNAMEEndpointWithOwnerResource("tar.test-zone.example.org", "new-tar.loadbalancer.com", "owner", "ingress/default/my-ingress-2"),
			newTXTEndpointWithOwnedRecord("cname-tar-txt.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner,external-dns/resource=ingress/default/my-ingress-2\"", "tar.test-zone.example.org"),
			newCNAMEEndpointWithOwnerResource("multiple.test-zone.example.org", "new.loadbalancer.com", "owner", "ingress/default/my-ingress-2").WithSetIdentifier("test-set-2"),
			newTXTEndpointWithOwnedRecord("cname-multiple-txt.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner,external-dns/resource=ingress/default/my-ingress-2\"", "multiple.test-zone.example.org").WithSetIdentifier("test-set-2"),
		},
		UpdateOld: []*endpoint.Endpoint{
			newEndpointWithOwner("tar.test-zone.example.org", "tar.loadbalancer.com", endpoint.RecordTypeCNAME, "owner"),
			newTXTEndpointWithOwnedRecord("cname-tar-txt.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", "tar.test-zone.example.org"),
			newEndpointWithOwner("multiple.test-zone.example.org", "lb2.loadbalancer.com", endpoint.RecordTypeCNAME, "owner").WithSetIdentifier("test-set-2"),
			newTXTEndpointWithOwnedRecord("cname-multiple-txt.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", "multiple.test-zone.example.org").WithSetIdentifier("test-set-2"),
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
		assert.Nil(t, ctx.Value(provider.RecordsContextKey))
	}
	err = r.ApplyChanges(ctx, changes)
	require.NoError(t, err)
}

func testTXTRegistryApplyChangesNoPrefix(t *testing.T) {
	p := inmemory.NewInMemoryProvider()
	err := p.CreateZone(testZone)
	require.NoError(t, err)
	var ctxEndpoints []*endpoint.Endpoint
	ctx := context.WithValue(context.Background(), provider.RecordsContextKey, ctxEndpoints)
	p.OnApplyChanges = func(ctx context.Context, _ *plan.Changes) {
		assert.Equal(t, ctxEndpoints, ctx.Value(provider.RecordsContextKey))
	}
	err = p.ApplyChanges(ctx, &plan.Changes{
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
	require.NoError(t, err)
	r, _ := NewTXTRegistry(p, "", "", "owner", time.Hour, "", []string{}, []string{}, false, nil, "")

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
			newTXTEndpointWithOwnedRecord("cname-new-record-1.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", "new-record-1.test-zone.example.org"),
			newEndpointWithOwner("example", "new-loadbalancer-1.lb.com", endpoint.RecordTypeCNAME, "owner"),
			newTXTEndpointWithOwnedRecord("cname-example", "\"heritage=external-dns,external-dns/owner=owner\"", "example"),
			newEndpointWithOwner("new-alias.test-zone.example.org", "my-domain.com", endpoint.RecordTypeA, "owner").WithProviderSpecific("alias", "true"),
			newTXTEndpointWithOwnedRecord("cname-new-alias.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", "new-alias.test-zone.example.org").WithProviderSpecific("alias", "true"),
		},
		Delete: []*endpoint.Endpoint{
			newEndpointWithOwner("foobar.test-zone.example.org", "foobar.loadbalancer.com", endpoint.RecordTypeCNAME, "owner"),
			newTXTEndpointWithOwnedRecord("cname-foobar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", "foobar.test-zone.example.org"),
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
		assert.Nil(t, ctx.Value(provider.RecordsContextKey))
	}
	err = r.ApplyChanges(ctx, changes)
	require.NoError(t, err)
}

func testTXTRegistryMissingRecords(t *testing.T) {
	t.Run("No prefix", testTXTRegistryMissingRecordsNoPrefix)
	t.Run("With Prefix", testTXTRegistryMissingRecordsWithPrefix)
}

func testTXTRegistryMissingRecordsNoPrefix(t *testing.T) {
	ctx := context.Background()
	p := inmemory.NewInMemoryProvider()
	err := p.CreateZone(testZone)
	require.NoError(t, err)
	err = p.ApplyChanges(ctx, &plan.Changes{
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
	require.NoError(t, err)
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

	r, _ := NewTXTRegistry(p, "", "", "owner", time.Hour, "wc", []string{endpoint.RecordTypeCNAME, endpoint.RecordTypeA, endpoint.RecordTypeNS}, []string{}, false, nil, "")
	records, _ := r.Records(ctx)

	assert.True(t, testutils.SameEndpoints(records, expectedRecords))
}

func testTXTRegistryMissingRecordsWithPrefix(t *testing.T) {
	ctx := context.Background()
	p := inmemory.NewInMemoryProvider()
	err := p.CreateZone(testZone)
	require.NoError(t, err)
	err = p.ApplyChanges(ctx, &plan.Changes{
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
	require.NoError(t, err)
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

	r, _ := NewTXTRegistry(p, "txt.", "", "owner", time.Hour, "wc", []string{endpoint.RecordTypeCNAME, endpoint.RecordTypeA, endpoint.RecordTypeNS, endpoint.RecordTypeTXT}, []string{}, false, nil, "")
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

func TestNewTXTScheme(t *testing.T) {
	p := inmemory.NewInMemoryProvider()
	err := p.CreateZone(testZone)
	require.NoError(t, err)
	var ctxEndpoints []*endpoint.Endpoint
	ctx := context.WithValue(context.Background(), provider.RecordsContextKey, ctxEndpoints)
	p.OnApplyChanges = func(ctx context.Context, _ *plan.Changes) {
		assert.Equal(t, ctxEndpoints, ctx.Value(provider.RecordsContextKey))
	}
	err = p.ApplyChanges(ctx, &plan.Changes{
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
	require.NoError(t, err)
	r, err := NewTXTRegistry(p, "", "", "owner", time.Hour, "", []string{}, []string{}, false, nil, "")
	require.NoError(t, err)

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
			newTXTEndpointWithOwnedRecord("cname-new-record-1.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", "new-record-1.test-zone.example.org"),
			newEndpointWithOwner("example", "new-loadbalancer-1.lb.com", endpoint.RecordTypeCNAME, "owner"),
			newTXTEndpointWithOwnedRecord("cname-example", "\"heritage=external-dns,external-dns/owner=owner\"", "example"),
		},
		Delete: []*endpoint.Endpoint{
			newEndpointWithOwner("foobar.test-zone.example.org", "foobar.loadbalancer.com", endpoint.RecordTypeCNAME, "owner"),
			newTXTEndpointWithOwnedRecord("cname-foobar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=owner\"", "foobar.test-zone.example.org"),
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
		assert.Nil(t, ctx.Value(provider.RecordsContextKey))
	}
	err = r.ApplyChanges(ctx, changes)
	require.NoError(t, err)
}

func TestGenerateTXT(t *testing.T) {
	record := newEndpointWithOwner("foo.test-zone.example.org", "new-foo.loadbalancer.com", endpoint.RecordTypeCNAME, "owner")
	expectedTXT := []*endpoint.Endpoint{
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
	err := p.CreateZone(testZone)
	require.NoError(t, err)
	r, _ := NewTXTRegistry(p, "", "", "owner", time.Hour, "", []string{}, []string{}, false, nil, "")
	gotTXT := r.generateTXTRecord(record)
	assert.Equal(t, expectedTXT, gotTXT)
}

func TestGenerateTXTWithMigration(t *testing.T) {
	record := newEndpointWithOwner("foo.test-zone.example.org", "1.2.3.4", endpoint.RecordTypeA, "owner")
	expectedTXTBeforeMigration := []*endpoint.Endpoint{
		{
			DNSName:    "a-foo.test-zone.example.org",
			Targets:    endpoint.Targets{"\"heritage=external-dns,external-dns/owner=owner\""},
			RecordType: endpoint.RecordTypeTXT,
			Labels: map[string]string{
				endpoint.OwnedRecordLabelKey: "foo.test-zone.example.org",
			},
		},
	}
	p := inmemory.NewInMemoryProvider()
	err := p.CreateZone(testZone)
	require.NoError(t, err)
	r, _ := NewTXTRegistry(p, "", "", "owner", time.Hour, "", []string{}, []string{}, false, nil, "")
	gotTXTBeforeMigration := r.generateTXTRecord(record)
	assert.Equal(t, expectedTXTBeforeMigration, gotTXTBeforeMigration)

	expectedTXTAfterMigration := []*endpoint.Endpoint{
		{
			DNSName:    "a-foo.test-zone.example.org",
			Targets:    endpoint.Targets{"\"heritage=external-dns,external-dns/owner=foobar\""},
			RecordType: endpoint.RecordTypeTXT,
			Labels: map[string]string{
				endpoint.OwnedRecordLabelKey: "foo.test-zone.example.org",
			},
		},
	}

	rMigrated, _ := NewTXTRegistry(p, "", "", "foobar", time.Hour, "", []string{}, []string{}, false, nil, "owner")
	gotTXTAfterMigration := rMigrated.generateTXTRecord(record)
	assert.Equal(t, expectedTXTAfterMigration, gotTXTAfterMigration)

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
	err := p.CreateZone(testZone)
	require.NoError(t, err)
	r, _ := NewTXTRegistry(p, "", "", "owner", time.Hour, "", []string{}, []string{}, false, nil, "")
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
	expectedTXT := make([]*endpoint.Endpoint, 0)
	p := inmemory.NewInMemoryProvider()
	err := p.CreateZone(testZone)
	require.NoError(t, err)
	r, _ := NewTXTRegistry(p, "", "", "owner", time.Hour, "", []string{}, []string{}, false, nil, "")
	gotTXT := r.generateTXTRecord(cnameRecord)
	assert.Equal(t, expectedTXT, gotTXT)
}

func TestTXTRegistryApplyChangesEncrypt(t *testing.T) {
	p := inmemory.NewInMemoryProvider()
	err := p.CreateZone(testZone)
	require.NoError(t, err)
	var ctxEndpoints []*endpoint.Endpoint
	ctx := context.WithValue(context.Background(), provider.RecordsContextKey, ctxEndpoints)

	err = p.ApplyChanges(ctx, &plan.Changes{
		Create: []*endpoint.Endpoint{
			newEndpointWithOwner("foobar.test-zone.example.org", "foobar.loadbalancer.com", endpoint.RecordTypeCNAME, ""),
			newTXTEndpointWithOwnedRecord("txt.cname-foobar.test-zone.example.org", "\"h8UQ6jelUFUsEIn7SbFktc2MYXPx/q8lySqI4VwfVtVaIbb2nkHWV/88KKbuLtu7fJNzMir8ELVeVnRSY01KdiIuj7ledqZe5ailEjQaU5Z6uEKd5pgs6sH8\"", "foobar.test-zone.example.org"),
		},
	})
	require.NoError(t, err)

	r, _ := NewTXTRegistry(p, "txt.", "", "owner", time.Hour, "", []string{}, []string{}, true, []byte("12345678901234567890123456789012"), "")
	records, _ := r.Records(ctx)
	changes := &plan.Changes{
		Delete: records,
	}

	// ensure that encryption nonce gets reused when deleting records
	expected := &plan.Changes{
		Delete: []*endpoint.Endpoint{
			newEndpointWithOwner("foobar.test-zone.example.org", "foobar.loadbalancer.com", endpoint.RecordTypeCNAME, "owner"),
			newTXTEndpointWithOwnedRecord("txt.cname-foobar.test-zone.example.org", "\"h8UQ6jelUFUsEIn7SbFktc2MYXPx/q8lySqI4VwfVtVaIbb2nkHWV/88KKbuLtu7fJNzMir8ELVeVnRSY01KdiIuj7ledqZe5ailEjQaU5Z6uEKd5pgs6sH8\"", "foobar.test-zone.example.org"),
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
		assert.Nil(t, ctx.Value(provider.RecordsContextKey))
	}
	err = r.ApplyChanges(ctx, changes)
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
	err := p.CreateZone(testZone)
	require.NoError(t, err)
	err = p.ApplyChanges(ctx, &plan.Changes{
		Create: []*endpoint.Endpoint{
			// records on cluster using A record for ingress address
			newEndpointWithOwner("bar.test-zone.example.org", "\"heritage=external-dns,external-dns/owner=cat,external-dns/resource=ingress/default/foo\"", endpoint.RecordTypeTXT, ""),
			newEndpointWithOwner("bar.test-zone.example.org", "1.2.3.4", endpoint.RecordTypeA, ""),
		},
	})
	require.NoError(t, err)

	r, _ := NewTXTRegistry(p, "_owner.", "", "bar", time.Hour, "", []string{}, []string{}, false, nil, "")
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
	p.OnApplyChanges = func(_ context.Context, changes *plan.Changes) {
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

	err = r.ApplyChanges(ctx, changes.Changes)
	require.NoError(t, err)
}

func TestGenerateTXTRecordWithNewFormatOnly(t *testing.T) {
	p := inmemory.NewInMemoryProvider()
	testCases := []struct {
		name            string
		endpoint        *endpoint.Endpoint
		expectedRecords int
		expectedPrefix  string
		description     string
	}{
		{
			name:            "legacy format enabled - standard record",
			endpoint:        newEndpointWithOwner("foo.test-zone.example.org", "1.2.3.4", endpoint.RecordTypeA, "owner"),
			expectedRecords: 1,
			expectedPrefix:  "a-",
			description:     "Should generate only new format TXT records",
		},
		{
			name:            "new format only - standard record",
			endpoint:        newEndpointWithOwner("foo.test-zone.example.org", "1.2.3.4", endpoint.RecordTypeA, "owner"),
			expectedRecords: 1,
			expectedPrefix:  "a-",
			description:     "Should only generate new format TXT record",
		},
		{
			name:            "legacy format enabled - AAAA record",
			endpoint:        newEndpointWithOwner("foo.test-zone.example.org", "2001:db8::1", endpoint.RecordTypeAAAA, "owner"),
			expectedRecords: 1,
			expectedPrefix:  "aaaa-",
			description:     "Should only generate new format for AAAA records regardless of setting",
		},
		{
			name:            "new format only - AAAA record",
			endpoint:        newEndpointWithOwner("foo.test-zone.example.org", "2001:db8::1", endpoint.RecordTypeAAAA, "owner"),
			expectedRecords: 1,
			expectedPrefix:  "aaaa-",
			description:     "Should only generate new format for AAAA records",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r, _ := NewTXTRegistry(p, "", "", "owner", time.Hour, "", []string{}, []string{}, false, nil, "")
			records := r.generateTXTRecord(tc.endpoint)

			assert.Len(t, records, tc.expectedRecords, tc.description)

			for _, record := range records {
				assert.Equal(t, endpoint.RecordTypeTXT, record.RecordType)
			}

			if tc.endpoint.RecordType == endpoint.RecordTypeAAAA {
				hasNewFormat := false
				for _, record := range records {
					if strings.HasPrefix(record.DNSName, tc.expectedPrefix) {
						hasNewFormat = true
						break
					}
				}
				assert.True(t, hasNewFormat,
					"Should have at least one record with prefix %s when using new format", tc.expectedPrefix)
			}
		})
	}
}

func TestApplyChangesWithNewFormatOnly(t *testing.T) {
	p := inmemory.NewInMemoryProvider()
	err := p.CreateZone(testZone)
	require.NoError(t, err)
	ctx := context.Background()

	r, _ := NewTXTRegistry(p, "", "", "owner", time.Hour, "", []string{}, []string{}, false, nil, "")

	changes := &plan.Changes{
		Create: []*endpoint.Endpoint{
			newEndpointWithOwner("new-record.test-zone.example.org", "1.2.3.4", endpoint.RecordTypeA, "owner"),
		},
	}

	err = r.ApplyChanges(ctx, changes)
	require.NoError(t, err)

	records, err := p.Records(ctx)
	require.NoError(t, err)

	var txtRecords []*endpoint.Endpoint
	for _, record := range records {
		if record.RecordType == endpoint.RecordTypeTXT {
			txtRecords = append(txtRecords, record)
		}
	}

	assert.Len(t, txtRecords, 1, "Should only create one TXT record in new format")

	if len(txtRecords) > 0 {
		assert.True(t, strings.HasPrefix(txtRecords[0].DNSName, "a-"),
			"TXT record should have 'a-' prefix when using new format only")
	}
}

func TestTXTRegistryRecordsWithEmptyTargets(t *testing.T) {
	ctx := context.Background()
	p := inmemory.NewInMemoryProvider()
	err := p.CreateZone(testZone)
	require.NoError(t, err)
	err = p.ApplyChanges(ctx, &plan.Changes{
		Create: []*endpoint.Endpoint{
			{
				DNSName:    "empty-targets.test-zone.example.org",
				RecordType: endpoint.RecordTypeTXT,
				Targets:    endpoint.Targets{},
			},
			{
				DNSName:    "valid-targets.test-zone.example.org",
				RecordType: endpoint.RecordTypeTXT,
				Targets:    endpoint.Targets{"target1"},
			},
		},
	})
	require.NoError(t, err)

	r, _ := NewTXTRegistry(p, "", "", "owner", time.Hour, "", []string{}, []string{}, false, nil, "")
	hook := logtest.LogsUnderTestWithLogLevel(log.ErrorLevel, t)
	records, err := r.Records(ctx)
	require.NoError(t, err)

	expectedRecords := []*endpoint.Endpoint{
		{
			DNSName:    "valid-targets.test-zone.example.org",
			Targets:    endpoint.Targets{"target1"},
			RecordType: endpoint.RecordTypeTXT,
			Labels:     map[string]string{},
		},
	}

	assert.True(t, testutils.SameEndpoints(records, expectedRecords))

	logtest.TestHelperLogContains("TXT record has no targets empty-targets.test-zone.example.org", hook, t)
}

// TestTXTRegistryRecreatesMissingRecords reproduces issue #4914.
// It verifies that External‑DNS recreates A/CNAME records that were accidentally deleted while their corresponding TXT records remain.
// An InMemoryProvider is used because, like Route53, it throws an error when attempting to create a duplicate record.
func TestTXTRegistryRecreatesMissingRecords(t *testing.T) {
	ownerId := "owner"
	tests := []struct {
		name           string
		desired        []*endpoint.Endpoint
		existing       []*endpoint.Endpoint
		expectedCreate []*endpoint.Endpoint
	}{
		{
			name: "Recreate missing A record when TXT exists",
			desired: []*endpoint.Endpoint{
				newEndpointWithOwner("new-record-1.test-zone.example.org", "1.1.1.1", endpoint.RecordTypeA, ""),
			},
			existing: []*endpoint.Endpoint{
				newEndpointWithOwner("new-record-1.test-zone.example.org", "\"heritage=external-dns,external-dns/owner="+ownerId+"\"", endpoint.RecordTypeTXT, ownerId),
				newEndpointWithOwner("a-new-record-1.test-zone.example.org", "\"heritage=external-dns,external-dns/owner="+ownerId+"\"", endpoint.RecordTypeTXT, ownerId),
			},
			expectedCreate: []*endpoint.Endpoint{
				newEndpointWithOwner("new-record-1.test-zone.example.org", "1.1.1.1", endpoint.RecordTypeA, ownerId),
			},
		},
		{
			name: "Recreate missing AAAA record when TXT exists",
			desired: []*endpoint.Endpoint{
				newEndpointWithOwner("new-record-1.test-zone.example.org", "2001:db8::1", endpoint.RecordTypeAAAA, ""),
			},
			existing: []*endpoint.Endpoint{
				newEndpointWithOwner("new-record-1.test-zone.example.org", "\"heritage=external-dns,external-dns/owner="+ownerId+"\"", endpoint.RecordTypeTXT, ownerId),
				newEndpointWithOwner("aaaa-new-record-1.test-zone.example.org", "\"heritage=external-dns,external-dns/owner="+ownerId+"\"", endpoint.RecordTypeTXT, ownerId),
			},
			expectedCreate: []*endpoint.Endpoint{
				newEndpointWithOwner("new-record-1.test-zone.example.org", "2001:db8::1", endpoint.RecordTypeAAAA, ownerId),
			},
		},
		{
			name: "Recreate missing CNAME record when TXT exists",
			desired: []*endpoint.Endpoint{
				newEndpointWithOwner("new-record-1.test-zone.example.org", "new-loadbalancer-1.lb.com", endpoint.RecordTypeCNAME, ""),
			},
			existing: []*endpoint.Endpoint{
				newEndpointWithOwner("new-record-1.test-zone.example.org", "\"heritage=external-dns,external-dns/owner="+ownerId+"\"", endpoint.RecordTypeTXT, ownerId),
				newEndpointWithOwner("cname-new-record-1.test-zone.example.org", "\"heritage=external-dns,external-dns/owner="+ownerId+"\"", endpoint.RecordTypeTXT, ownerId),
			},
			expectedCreate: []*endpoint.Endpoint{
				newEndpointWithOwner("new-record-1.test-zone.example.org", "new-loadbalancer-1.lb.com", endpoint.RecordTypeCNAME, ownerId)},
		},
		{
			name: "Recreate missing A and CNAME records when TXT exists",
			desired: []*endpoint.Endpoint{
				newEndpointWithOwner("new-record-1.test-zone.example.org", "1.1.1.1", endpoint.RecordTypeA, ""),
				newEndpointWithOwner("new-record-2.test-zone.example.org", "new-loadbalancer-1.lb.com", endpoint.RecordTypeCNAME, ""),
			},
			existing: []*endpoint.Endpoint{
				newEndpointWithOwner("new-record-1.test-zone.example.org", "\"heritage=external-dns,external-dns/owner="+ownerId+"\"", endpoint.RecordTypeTXT, ownerId),
				newEndpointWithOwner("new-record-2.test-zone.example.org", "\"heritage=external-dns,external-dns/owner="+ownerId+"\"", endpoint.RecordTypeTXT, ownerId),
				newEndpointWithOwner("a-new-record-1.test-zone.example.org", "\"heritage=external-dns,external-dns/owner="+ownerId+"\"", endpoint.RecordTypeTXT, ownerId),
				newEndpointWithOwner("cname-new-record-2.test-zone.example.org", "\"heritage=external-dns,external-dns/owner="+ownerId+"\"", endpoint.RecordTypeTXT, ownerId),
			},
			expectedCreate: []*endpoint.Endpoint{
				newEndpointWithOwner("new-record-1.test-zone.example.org", "1.1.1.1", endpoint.RecordTypeA, ownerId),
				newEndpointWithOwner("new-record-2.test-zone.example.org", "new-loadbalancer-1.lb.com", endpoint.RecordTypeCNAME, ownerId),
			},
		},
		{
			name: "Recreate missing A records when TXT and CNAME exists",
			desired: []*endpoint.Endpoint{
				newEndpointWithOwner("new-record-1.test-zone.example.org", "1.1.1.1", endpoint.RecordTypeA, ""),
				newEndpointWithOwner("new-record-2.test-zone.example.org", "new-loadbalancer-1.lb.com", endpoint.RecordTypeCNAME, ""),
			},
			existing: []*endpoint.Endpoint{
				newEndpointWithOwner("new-record-2.test-zone.example.org", "new-loadbalancer-1.lb.com", endpoint.RecordTypeCNAME, ownerId),
				newEndpointWithOwner("new-record-1.test-zone.example.org", "\"heritage=external-dns,external-dns/owner="+ownerId+"\"", endpoint.RecordTypeTXT, ownerId),
				newEndpointWithOwner("new-record-2.test-zone.example.org", "\"heritage=external-dns,external-dns/owner="+ownerId+"\"", endpoint.RecordTypeTXT, ownerId),
				newEndpointWithOwner("a-new-record-1.test-zone.example.org", "\"heritage=external-dns,external-dns/owner="+ownerId+"\"", endpoint.RecordTypeTXT, ownerId),
				newEndpointWithOwner("cname-new-record-2.test-zone.example.org", "\"heritage=external-dns,external-dns/owner="+ownerId+"\"", endpoint.RecordTypeTXT, ownerId),
			},
			expectedCreate: []*endpoint.Endpoint{
				newEndpointWithOwner("new-record-1.test-zone.example.org", "1.1.1.1", endpoint.RecordTypeA, ownerId),
			},
		},
		{
			name: "Only one A record is missing among several existing records",
			desired: []*endpoint.Endpoint{
				newEndpointWithOwner("record-1.test-zone.example.org", "1.1.1.1", endpoint.RecordTypeA, ""),
				newEndpointWithOwner("record-2.test-zone.example.org", "1.1.1.2", endpoint.RecordTypeA, ""),
				newEndpointWithOwner("record-3.test-zone.example.org", "1.1.1.3", endpoint.RecordTypeA, ""),
				newEndpointWithOwner("record-4.test-zone.example.org", "2001:db8::4", endpoint.RecordTypeAAAA, ""),
				newEndpointWithOwner("record-5.test-zone.example.org", "cluster-b", endpoint.RecordTypeCNAME, ""),
			},
			existing: []*endpoint.Endpoint{
				newEndpointWithOwner("record-1.test-zone.example.org", "\"heritage=external-dns,external-dns/owner="+ownerId+"\"", endpoint.RecordTypeTXT, ownerId),
				newEndpointWithOwner("a-record-1.test-zone.example.org", "\"heritage=external-dns,external-dns/owner="+ownerId+"\"", endpoint.RecordTypeTXT, ownerId),

				newEndpointWithOwner("record-2.test-zone.example.org", "1.1.1.2", endpoint.RecordTypeA, ownerId),
				newEndpointWithOwner("record-2.test-zone.example.org", "\"heritage=external-dns,external-dns/owner="+ownerId+"\"", endpoint.RecordTypeTXT, ownerId),
				newEndpointWithOwner("a-record-2.test-zone.example.org", "\"heritage=external-dns,external-dns/owner="+ownerId+"\"", endpoint.RecordTypeTXT, ownerId),

				newEndpointWithOwner("record-3.test-zone.example.org", "1.1.1.3", endpoint.RecordTypeA, ownerId),
				newEndpointWithOwner("record-3.test-zone.example.org", "\"heritage=external-dns,external-dns/owner="+ownerId+"\"", endpoint.RecordTypeTXT, ownerId),
				newEndpointWithOwner("a-record-3.test-zone.example.org", "\"heritage=external-dns,external-dns/owner="+ownerId+"\"", endpoint.RecordTypeTXT, ownerId),

				newEndpointWithOwner("record-4.test-zone.example.org", "2001:db8::4", endpoint.RecordTypeAAAA, ownerId),
				newEndpointWithOwner("record-4.test-zone.example.org", "\"heritage=external-dns,external-dns/owner="+ownerId+"\"", endpoint.RecordTypeTXT, ownerId),
				newEndpointWithOwner("aaaa-record-4.test-zone.example.org", "\"heritage=external-dns,external-dns/owner="+ownerId+"\"", endpoint.RecordTypeTXT, ownerId),

				newEndpointWithOwner("record-5.test-zone.example.org", "cluster-b", endpoint.RecordTypeCNAME, ownerId),
				newEndpointWithOwner("record-5.test-zone.example.org", "\"heritage=external-dns,external-dns/owner="+ownerId+"\"", endpoint.RecordTypeTXT, ownerId),
				newEndpointWithOwner("cname-record-5.test-zone.example.org", "\"heritage=external-dns,external-dns/owner="+ownerId+"\"", endpoint.RecordTypeTXT, ownerId),
			},
			expectedCreate: []*endpoint.Endpoint{
				newEndpointWithOwner("record-1.test-zone.example.org", "1.1.1.1", endpoint.RecordTypeA, ownerId),
			},
		},
		{
			name: "Should not recreate TXT records for existing A records without owner",
			desired: []*endpoint.Endpoint{
				newEndpointWithOwner("record-1.test-zone.example.org", "1.1.1.1", endpoint.RecordTypeA, ""),
			},
			existing: []*endpoint.Endpoint{
				newEndpointWithOwner("record-1.test-zone.example.org", "1.1.1.1", endpoint.RecordTypeA, ownerId),
				// Missing TXT record for the existing A record
			},
			expectedCreate: []*endpoint.Endpoint{},
		},
		{
			name: "Should not recreate TXT records for existing A records with another owner",
			desired: []*endpoint.Endpoint{
				newEndpointWithOwner("record-1.test-zone.example.org", "1.1.1.1", endpoint.RecordTypeA, ""),
			},
			existing: []*endpoint.Endpoint{
				// This test uses the `ownerId` variable, and "another-owner" simulates a different owner.
				// In this case, TXT records should not be recreated.
				newEndpointWithOwner("record-1.test-zone.example.org", "1.1.1.1", endpoint.RecordTypeA, "another-owner"),
				newEndpointWithOwner("a-record-1.test-zone.example.org", "\"heritage=external-dns,external-dns/owner="+"another-owner"+"\"", endpoint.RecordTypeTXT, "another-owner"),
			},
			expectedCreate: []*endpoint.Endpoint{},
		},
	}
	for _, tt := range tests {
		for _, setIdentifier := range []string{"", "set-identifier"} {
			for pName, policy := range plan.Policies {
				// Clone inputs per policy to avoid data races when using t.Parallel.
				desired := cloneEndpointsWithOpts(tt.desired, func(e *endpoint.Endpoint) {
					e.WithSetIdentifier(setIdentifier)
				})
				existing := cloneEndpointsWithOpts(tt.existing, func(e *endpoint.Endpoint) {
					e.WithSetIdentifier(setIdentifier)
				})
				expectedCreate := cloneEndpointsWithOpts(tt.expectedCreate, func(e *endpoint.Endpoint) {
					e.WithSetIdentifier(setIdentifier)
				})

				t.Run(fmt.Sprintf("%s with %s policy and setIdentifier=%s", tt.name, pName, setIdentifier), func(t *testing.T) {
					t.Parallel()
					ctx := context.Background()
					p := inmemory.NewInMemoryProvider()

					// Given: Register existing records
					err := p.CreateZone(testZone)
					require.NoError(t, err)
					err = p.ApplyChanges(ctx, &plan.Changes{Create: existing})
					assert.NoError(t, err)

					// The first ApplyChanges call should create the expected records.
					// Subsequent calls are expected to be no-ops (i.e., no additional creates).
					isCalled := false
					p.OnApplyChanges = func(_ context.Context, changes *plan.Changes) {
						if isCalled {
							assert.Empty(t, changes.Create, "ApplyChanges should not be called multiple times with new changes")
						} else {
							assert.True(t,
								testutils.SameEndpoints(changes.Create, expectedCreate),
								"Expected create changes: %v, but got: %v", expectedCreate, changes.Create,
							)
						}
						assert.Empty(t, changes.UpdateNew, "UpdateNew should be empty")
						assert.Empty(t, changes.UpdateOld, "UpdateOld should be empty")
						assert.Empty(t, changes.Delete, "Delete should be empty")
						isCalled = true
					}

					// When: Apply changes to recreate missing A records
					managedRecords := []string{endpoint.RecordTypeA, endpoint.RecordTypeCNAME, endpoint.RecordTypeAAAA, endpoint.RecordTypeTXT}
					registry, err := NewTXTRegistry(p, "", "", ownerId, time.Hour, "", managedRecords, nil, false, nil, "")
					assert.NoError(t, err)

					expectedRecords := append(existing, expectedCreate...) // nolint:gocritic

					// Simulate the reconciliation loop by executing multiple times
					reconciliationLoops := 3
					for i := range reconciliationLoops {
						records, err := registry.Records(ctx)
						assert.NoError(t, err)
						pl := &plan.Plan{
							Policies:       []plan.Policy{policy},
							Current:        records,
							Desired:        desired,
							ManagedRecords: managedRecords,
							OwnerID:        ownerId,
						}
						pln := pl.Calculate()
						err = registry.ApplyChanges(ctx, pln.Changes)
						assert.NoError(t, err)

						// Then: Verify that the missing records are recreated or the existing records are not modified
						records, err = p.Records(ctx)
						assert.NoError(t, err)
						assert.True(t, testutils.SameEndpoints(records, expectedRecords),
							"Expected records after reconciliation loop #%d: %v, but got: %v",
							i, expectedRecords, records,
						)
					}
				})
			}
		}
	}
}

func TestTXTRecordMigration(t *testing.T) {
	ctx := context.Background()
	p := inmemory.NewInMemoryProvider()
	err := p.CreateZone(testZone)
	require.NoError(t, err)

	r, _ := NewTXTRegistry(p, "%{record_type}-", "", "foo", time.Hour, "", []string{}, []string{}, false, nil, "")

	err = r.ApplyChanges(ctx, &plan.Changes{
		Create: []*endpoint.Endpoint{
			// records on cluster using A record for ingress address
			newEndpointWithOwnerAndLabels("bar.test-zone.example.org", "1.2.3.4", endpoint.RecordTypeA, "foo", endpoint.Labels{endpoint.OwnerLabelKey: "owner"}),
		},
	})
	require.NoError(t, err)

	createdRecords, _ := r.Records(ctx)

	newTXTRecord := r.generateTXTRecord(createdRecords[0])

	expectedTXTRecords := []*endpoint.Endpoint{
		{
			DNSName:    "a-bar.test-zone.example.org",
			Targets:    endpoint.Targets{"\"heritage=external-dns,external-dns/owner=foo\""},
			RecordType: endpoint.RecordTypeTXT,
		},
	}

	assert.Equal(t, expectedTXTRecords[0].Targets, newTXTRecord[0].Targets)

	r, _ = NewTXTRegistry(p, "%{record_type}-", "", "foobar", time.Hour, "", []string{}, []string{}, false, nil, "foo")

	updatedRecords, _ := r.Records(ctx)

	updatedTXTRecord := r.generateTXTRecord(updatedRecords[0])

	expectedFinalTXT := []*endpoint.Endpoint{
		{
			DNSName:    "a-bar.test-zone.example.org",
			Targets:    endpoint.Targets{"\"heritage=external-dns,external-dns/owner=foobar\""},
			RecordType: endpoint.RecordTypeTXT,
		},
	}

	assert.Equal(t, updatedTXTRecord[0].Targets, expectedFinalTXT[0].Targets)
}

// TestRecreateRecordAfterDeletion ensures that when A and TXT records are deleted,
// both are correctly recreated in subsequent reconciliation loops.
// This prevents regression of the issue where stale TXT record state
// caused ExternalDNS to skip recreating TXT records after deletion.
func TestRecreateRecordAfterDeletion(t *testing.T) {
	ownerID := "foo"
	ctx := context.Background()
	p := inmemory.NewInMemoryProvider()
	err := p.CreateZone(testZone)
	require.NoError(t, err)

	r, _ := NewTXTRegistry(p, "%{record_type}-", "", "foo", 0, "", []string{endpoint.RecordTypeA}, []string{}, false, nil, "")

	createdRecords := newEndpointWithOwnerAndLabels("bar.test-zone.example.org", "1.2.3.4", endpoint.RecordTypeA, ownerID, nil)
	txtRecord := r.generateTXTRecord(createdRecords)

	// 1. Create initial A and TXT records.
	creates := append([]*endpoint.Endpoint{createdRecords}, txtRecord...)
	err = p.ApplyChanges(ctx, &plan.Changes{
		Create: creates,
	})
	assert.NoError(t, err)

	// 2. Simulate a "no change" reconciliation (ApplyChanges won't be called).
	desired := []*endpoint.Endpoint{
		{
			DNSName: "bar.test-zone.example.org",
			Targets: endpoint.Targets{
				"1.2.3.4",
			},
			RecordType: endpoint.RecordTypeA,
		},
	}

	records, err := r.Records(ctx)
	assert.NoError(t, err)

	calculated := &plan.Plan{
		Policies:       []plan.Policy{&plan.SyncPolicy{}},
		ManagedRecords: []string{endpoint.RecordTypeA},
		Current:        records,
		Desired:        desired,
		OwnerID:        ownerID,
	}
	calculated = calculated.Calculate()
	// ApplyChanges is not called to simulate no changes.
	assert.False(t, calculated.Changes.HasChanges(), "There should be no changes")

	// 3. Delete both A and TXT records (simulate manual deletion)
	deletes := append([]*endpoint.Endpoint{createdRecords}, txtRecord...)
	err = p.ApplyChanges(ctx, &plan.Changes{
		Delete: deletes,
	})
	assert.NoError(t, err)

	// 4. Run reconciliation again — both A and TXT should be recreated.
	records, err = r.Records(ctx)
	assert.NoError(t, err)

	calculated = &plan.Plan{
		Policies:       []plan.Policy{&plan.SyncPolicy{}},
		ManagedRecords: []string{endpoint.RecordTypeA},
		Current:        records,
		Desired:        desired,
		OwnerID:        ownerID,
	}
	calculated = calculated.Calculate()
	if !calculated.Changes.HasChanges() {
		assert.Fail(t, "There should be changes")
	}

	err = r.ApplyChanges(ctx, calculated.Changes)
	assert.NoError(t, err)

	// 5. Verify that both A and TXT records are recreated successfully.
	records, err = p.Records(ctx)
	assert.NoError(t, err)
	assert.True(t, testutils.SameEndpoints(records, append(desired, txtRecord...)), "Expected records after reconciliation: %v, but got: %v", append(desired, txtRecord...), records)
}
