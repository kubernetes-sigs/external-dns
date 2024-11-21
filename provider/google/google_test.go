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

package google

import (
	"fmt"
	"maps"
	"net/http"
	"reflect"
	"slices"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/context"
	dns "google.golang.org/api/dns/v1"
	"google.golang.org/api/googleapi"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

func TestGoogleZonesFilters(t *testing.T) {
	p := newGoogleProvider().WithMockClients(newMockClients(t).WithRecords(
		&map[*dns.ManagedZone][]*dns.ResourceRecordSet{
			&dns.ManagedZone{
				Name: "domain",
				DnsName: "domain.",
				Id: 10000,
				Visibility: "private",
			}: []*dns.ResourceRecordSet{},
			&dns.ManagedZone{
				Name: "internal-1",
				DnsName: "cluster.local.",
				Id: 10001,
				Visibility: "private",
			}: []*dns.ResourceRecordSet{},
			&dns.ManagedZone{
				Name: "internal-2",
				DnsName: "cluster.local.",
				Id: 10002,
				Visibility: "private",
			}: []*dns.ResourceRecordSet{},
			&dns.ManagedZone{
				Name: "internal-3",
				DnsName: "cluster.local.",
				Id: 10003,
				Visibility: "private",
			}: []*dns.ResourceRecordSet{},
			&dns.ManagedZone{
				Name: "split-horizon-1",
				DnsName: "cluster.local.",
				Id: 10004,
				Visibility: "public",
			}: []*dns.ResourceRecordSet{},
			&dns.ManagedZone{
				Name: "split-horizon-2",
				DnsName: "cluster.local.",
				Id: 10004,
				Visibility: "private",
			}: []*dns.ResourceRecordSet{},
			&dns.ManagedZone{
				Name: "svc-local",
				DnsName: "svc.local.",
				Id: 10005,
				Visibility: "private",
			}: []*dns.ResourceRecordSet{},
			&dns.ManagedZone{
				Name: "svc-local-peer",
				DnsName: "svc.local.",
				Id: 10006,
				Visibility: "private",
				PeeringConfig: &dns.ManagedZonePeeringConfig{TargetNetwork: nil},
			}: []*dns.ResourceRecordSet{},
		},
	))
	for _, test := range []struct{
		name string
		domainFilter endpoint.DomainFilter
		zoneIDFilter provider.ZoneIDFilter
		zoneTypeFilter provider.ZoneTypeFilter
		filterdZones map[string]*dns.ManagedZone
	}{
		{
			name: "Domain",
			domainFilter: endpoint.NewDomainFilter([]string{"domain."}),
			zoneIDFilter: provider.NewZoneIDFilter([]string{}),
			zoneTypeFilter: provider.NewZoneTypeFilter(""),
			filterdZones: map[string]*dns.ManagedZone{
				"domain": {Name: "domain", DnsName: "domain.", Id: 10000, Visibility: "private"},
			},
		},
		{
			name: "ID",
			domainFilter: endpoint.NewDomainFilter([]string{}),
			zoneIDFilter: provider.NewZoneIDFilter([]string{"10002"}),
			zoneTypeFilter: provider.NewZoneTypeFilter(""),
			filterdZones: map[string]*dns.ManagedZone{
				"internal-2": {Name: "internal-2", DnsName: "cluster.local.", Id: 10002, Visibility: "private"},
			},
		},
		{
			name: "Name",
			domainFilter: endpoint.NewDomainFilter([]string{}),
			zoneIDFilter: provider.NewZoneIDFilter([]string{"internal-2"}),
			zoneTypeFilter: provider.NewZoneTypeFilter(""),
			filterdZones: map[string]*dns.ManagedZone{
				"internal-2": {Name: "internal-2", DnsName: "cluster.local.", Id: 10002, Visibility: "private"},
			},
		},
		{
			name: "Public",
			domainFilter: endpoint.NewDomainFilter([]string{}),
			zoneIDFilter: provider.NewZoneIDFilter([]string{"10004"}),
			zoneTypeFilter: provider.NewZoneTypeFilter("public"),
			filterdZones: map[string]*dns.ManagedZone{
				"split-horizon-1": {Name: "split-horizon-1", DnsName: "cluster.local.", Id: 10004, Visibility: "public"},
			},
		},
		{
			name: "Private",
			domainFilter: endpoint.NewDomainFilter([]string{}),
			zoneIDFilter: provider.NewZoneIDFilter([]string{"10004"}),
			zoneTypeFilter: provider.NewZoneTypeFilter("private"),
			filterdZones: map[string]*dns.ManagedZone{
				"split-horizon-2": {Name: "split-horizon-2", DnsName: "cluster.local.", Id: 10004, Visibility: "private"},
			},
		},
		{
			name: "Peering",
			domainFilter: endpoint.NewDomainFilter([]string{"svc.local."}),
			zoneIDFilter: provider.NewZoneIDFilter([]string{""}),
			zoneTypeFilter: provider.NewZoneTypeFilter(""),
			filterdZones: map[string]*dns.ManagedZone{
				"svc-local": {Name: "svc-local", DnsName: "svc.local.", Id: 10005, Visibility: "private"},
			},
		},
	} {
		p.domainFilter = test.domainFilter
		p.zoneTypeFilter = test.zoneTypeFilter
		p.zoneIDFilter = test.zoneIDFilter
		t.Run(test.name, func (t *testing.T) {
			zones, err := p.Zones(context.Background())
			require.NoError(t, err)
			assert.Equal(t, test.filterdZones, zones)
		})
	}
}

func TestGoogleRecords(t *testing.T) {
	p := newGoogleProvider().WithMockClients(newMockClients(t).WithRecords(
		&map[*dns.ManagedZone][]*dns.ResourceRecordSet{
			&dns.ManagedZone{
				Name:    "zone-1",
				DnsName: "zone-1.local.",
			}: []*dns.ResourceRecordSet{
				&dns.ResourceRecordSet{
					Name: "a.zone-1.local.",
					Type: "A",
					Ttl: 1,
					Rrdatas: []string{"1.0.0.0"},
				},
				&dns.ResourceRecordSet{
					Name: "cname.zone-1.local.",
					Type: "CNAME",
					Ttl: 3,
					Rrdatas: []string{"cname."},
				},
			},
			&dns.ManagedZone{
				Name:    "zone-2",
				DnsName: "zone-2.local.",
			}: []*dns.ResourceRecordSet{
				&dns.ResourceRecordSet{
					Name: "a.zone-2.local.",
					Type: "A",
					Ttl: 2,
					Rrdatas: []string{"2.0.0.0"},
				},
			},
			&dns.ManagedZone{
				Name:    "zone-3",
				DnsName: "zone-3.local.",
			}: []*dns.ResourceRecordSet{},
		},
	))
	records, err := p.Records(context.Background())
	require.NoError(t, err)
	validateEndpoints(t, records, []*endpoint.Endpoint{
		endpoint.NewEndpointWithTTL("a.zone-1.local", endpoint.RecordTypeA, endpoint.TTL(1), "1.0.0.0"),
		endpoint.NewEndpointWithTTL("cname.zone-1.local", endpoint.RecordTypeCNAME, endpoint.TTL(3), "cname"),
		endpoint.NewEndpointWithTTL("a.zone-2.local", endpoint.RecordTypeA, endpoint.TTL(2), "2.0.0.0"),
	})
}

func TestGoogleRecordsFilter(t *testing.T) {
	p := newGoogleProvider().WithMockClients(newMockClients(t).WithRecords(
		&map[*dns.ManagedZone][]*dns.ResourceRecordSet{
			&dns.ManagedZone{
				Name:    "zone-1",
				DnsName: "zone-1.local.",
			}: []*dns.ResourceRecordSet{
				&dns.ResourceRecordSet{
					Name: "a.zone-1.local.",
					Type: "A",
					Ttl: googleRecordTTL,
					Rrdatas: []string{"1.0.0.0"},
				},
			},
			&dns.ManagedZone{
				Name:    "zone-2",
				DnsName: "zone-2.local.",
			}: []*dns.ResourceRecordSet{
				&dns.ResourceRecordSet{
					Name: "a.zone-2.local.",
					Type: "A",
					Ttl: googleRecordTTL,
					Rrdatas: []string{"2.0.0.0"},
				},
			},
			&dns.ManagedZone{
				Name:    "zone-3",
				DnsName: "zone-3.local.",
			}: []*dns.ResourceRecordSet{
				&dns.ResourceRecordSet{
					Name: "a.zone-3.local.",
					Type: "A",
					Ttl: googleRecordTTL,
					Rrdatas: []string{"3.0.0.0"},
				},
			},
		},
	))
	p.domainFilter = endpoint.NewDomainFilter([]string{
		// our two valid zones
		"zone-1.local.",
		"zone-2.local.",
		// we filter for a zone that doesn't exist, should have no effect.
		"zone-0.local.",
		// there exists a third zone "zone-3" that we want to exclude from being managed.
	})
	records, err := p.Records(context.Background())
	require.NoError(t, err)
	validateEndpoints(t, records, []*endpoint.Endpoint{
		endpoint.NewEndpointWithTTL("a.zone-1.local", endpoint.RecordTypeA, googleRecordTTL, "1.0.0.0"),
		endpoint.NewEndpointWithTTL("a.zone-2.local", endpoint.RecordTypeA, googleRecordTTL, "2.0.0.0"),
	})
}

func TestGoogleApplyChanges(t *testing.T) {
	records := map[*dns.ManagedZone][]*dns.ResourceRecordSet{
		&dns.ManagedZone{
			Name:    "zone-1",
			DnsName: "zone-1.local.",
		}: []*dns.ResourceRecordSet{
			&dns.ResourceRecordSet{
				Name: "update-test.zone-1.local.",
				Type: "A",
				Ttl: googleRecordTTL,
				Rrdatas: []string{"8.8.8.8"},
			},
			&dns.ResourceRecordSet{
				Name: "delete-test.zone-1.local.",
				Type: "A",
				Ttl: googleRecordTTL,
				Rrdatas: []string{"8.8.8.8"},
			},
			&dns.ResourceRecordSet{
				Name: "update-test-cname.zone-1.local.",
				Type: "CNAME",
				Ttl: googleRecordTTL,
				Rrdatas: []string{"update-test-cname."},
			},
			&dns.ResourceRecordSet{
				Name: "delete-test-cname.zone-1.local.",
				Type: "CNAME",
				Ttl: googleRecordTTL,
				Rrdatas: []string{"delete-test-cname."},
			},
		},
		&dns.ManagedZone{
			Name:    "zone-2",
			DnsName: "zone-2.local.",
		}: []*dns.ResourceRecordSet{
			&dns.ResourceRecordSet{
				Name: "update-test-ttl.zone-2.local.",
				Type: "A",
				Ttl: googleRecordTTL,
				Rrdatas: []string{"8.8.4.4"},
			},
			&dns.ResourceRecordSet{
				Name: "delete-test.zone-2.local.",
				Type: "A",
				Ttl: googleRecordTTL,
				Rrdatas: []string{"8.8.4.4"},
			},
		},
	}
	mockClients := newMockClients(t).WithRecords(&records)
	p := newGoogleProvider().WithMockClients(mockClients)
	require.NoError(t, p.ApplyChanges(context.Background(), &plan.Changes{
		Create: []*endpoint.Endpoint{
			endpoint.NewEndpoint("create-test.zone-1.local", endpoint.RecordTypeA, "8.8.8.8"),
			endpoint.NewEndpointWithTTL("create-test-ttl.zone-2.local", endpoint.RecordTypeA, endpoint.TTL(15), "8.8.4.4"),
			endpoint.NewEndpoint("create-test-cname.zone-1.local", endpoint.RecordTypeCNAME, "create-test-cname"),
			endpoint.NewEndpoint("create-test-ns.zone-1.local", endpoint.RecordTypeNS, "create-test-ns"),
			endpoint.NewEndpoint("filter-create-test.zone-3.local", endpoint.RecordTypeA, "4.2.2.2"),
			endpoint.NewEndpoint("nomatch-create-test.zone-0.local", endpoint.RecordTypeA, "4.2.2.1"),
		},
		UpdateOld: []*endpoint.Endpoint{
			endpoint.NewEndpoint("update-test.zone-1.local", endpoint.RecordTypeA, "8.8.8.8"),
			endpoint.NewEndpoint("update-test-ttl.zone-2.local", endpoint.RecordTypeA, "8.8.4.4"),
			endpoint.NewEndpoint("update-test-cname.zone-1.local", endpoint.RecordTypeCNAME, "update-test-cname"),
			endpoint.NewEndpoint("filter-update-test.zone-3.local", endpoint.RecordTypeA, "4.2.2.2"),
		},
		UpdateNew: []*endpoint.Endpoint{
			endpoint.NewEndpoint("update-test.zone-1.local", endpoint.RecordTypeA, "1.2.3.4"),
			endpoint.NewEndpointWithTTL("update-test-ttl.zone-2.local", endpoint.RecordTypeA, endpoint.TTL(25), "4.3.2.1"),
			endpoint.NewEndpoint("update-test-cname.zone-1.local", endpoint.RecordTypeCNAME, "updated-test-cname"),
			endpoint.NewEndpoint("filter-update-test.zone-3.local", endpoint.RecordTypeA, "5.6.7.8"),
			endpoint.NewEndpoint("nomatch-update-test.zone-0.local", endpoint.RecordTypeA, "8.7.6.5"),
		},
		Delete: []*endpoint.Endpoint{
			endpoint.NewEndpoint("delete-test.zone-1.local", endpoint.RecordTypeA, "8.8.8.8"),
			endpoint.NewEndpoint("delete-test.zone-2.local", endpoint.RecordTypeA, "8.8.4.4"),
			endpoint.NewEndpoint("delete-test-cname.zone-1.local", endpoint.RecordTypeCNAME, "delete-test-cname"),
			endpoint.NewEndpoint("filter-delete-test.zone-3.local", endpoint.RecordTypeA, "4.2.2.2"),
			endpoint.NewEndpoint("nomatch-delete-test.zone-0.local", endpoint.RecordTypeA, "4.2.2.1"),
		},
	}))
	mockClients.EqualRecords(&map[*dns.ManagedZone][]*dns.ResourceRecordSet{
		&dns.ManagedZone{
			Name:    "zone-1",
			DnsName: "zone-1.local.",
		}: []*dns.ResourceRecordSet{
			&dns.ResourceRecordSet{
				Name: "create-test.zone-1.local.",
				Type: "A",
				Ttl: googleRecordTTL,
				Rrdatas: []string{"8.8.8.8"},
			},
			&dns.ResourceRecordSet{
				Name: "update-test.zone-1.local.",
				Type: "A",
				Ttl: googleRecordTTL,
				Rrdatas: []string{"1.2.3.4"},
			},
			&dns.ResourceRecordSet{
				Name: "create-test-cname.zone-1.local.",
				Type: "CNAME",
				Ttl: googleRecordTTL,
				Rrdatas: []string{"create-test-cname."},
			},
			&dns.ResourceRecordSet{
				Name: "update-test-cname.zone-1.local.",
				Type: "CNAME",
				Ttl: googleRecordTTL,
				Rrdatas: []string{"updated-test-cname."},
			},
			&dns.ResourceRecordSet{
				Name: "create-test-ns.zone-1.local.",
				Type: "NS",
				Ttl: googleRecordTTL,
				Rrdatas: []string{"create-test-ns."},
			},
		},
		&dns.ManagedZone{
			Name:    "zone-2",
			DnsName: "zone-2.local.",
		}: []*dns.ResourceRecordSet{
			&dns.ResourceRecordSet{
				Name: "create-test-ttl.zone-2.local.",
				Type: "A",
				Ttl: 15,
				Rrdatas: []string{"8.8.4.4"},
			},
			&dns.ResourceRecordSet{
				Name: "update-test-ttl.zone-2.local.",
				Type: "A",
				Ttl: 25,
				Rrdatas: []string{"4.3.2.1"},
			},
		},
	})
}

func TestGoogleApplyChangesDryRun(t *testing.T) {
	records := map[*dns.ManagedZone][]*dns.ResourceRecordSet{
		&dns.ManagedZone{
			Name:    "zone-1",
			DnsName: "zone-1.local.",
		}: []*dns.ResourceRecordSet{
			&dns.ResourceRecordSet{
				Name: "update-test.zone-1.local.",
				Type: "A",
				Ttl: googleRecordTTL,
				Rrdatas: []string{"8.8.8.8"},
			},
			&dns.ResourceRecordSet{
				Name: "delete-test.zone-1.local.",
				Type: "A",
				Ttl: googleRecordTTL,
				Rrdatas: []string{"8.8.8.8"},
			},
			&dns.ResourceRecordSet{
				Name: "update-test-cname.zone-1.local.",
				Type: "CNAME",
				Ttl: googleRecordTTL,
				Rrdatas: []string{"update-test-cname."},
			},
			&dns.ResourceRecordSet{
				Name: "delete-test-cname.zone-1.local.",
				Type: "CNAME",
				Ttl: googleRecordTTL,
				Rrdatas: []string{"delete-test-cname."},
			},
		},
		&dns.ManagedZone{
			Name:    "zone-2",
			DnsName: "zone-2.local.",
		}: []*dns.ResourceRecordSet{
			&dns.ResourceRecordSet{
				Name: "update-test.zone-2.local.",
				Type: "A",
				Ttl: googleRecordTTL,
				Rrdatas: []string{"8.8.4.4"},
			},
			&dns.ResourceRecordSet{
				Name: "delete-test.zone-2.local.",
				Type: "A",
				Ttl: googleRecordTTL,
				Rrdatas: []string{"8.8.4.4"},
			},
		},
	}
	mockClients := newMockClients(t).WithRecords(&records)
	p := newGoogleProvider().WithMockClients(mockClients)
	p.dryRun = true
	require.NoError(t, p.ApplyChanges(context.Background(), &plan.Changes{
		Create: []*endpoint.Endpoint{
			endpoint.NewEndpoint("create-test.zone-1.local", endpoint.RecordTypeA, "8.8.8.8"),
			endpoint.NewEndpoint("create-test.zone-2.local", endpoint.RecordTypeA, "8.8.4.4"),
			endpoint.NewEndpoint("create-test-cname.zone-1.local", endpoint.RecordTypeCNAME, "create-test-cname"),
		},
		UpdateOld: []*endpoint.Endpoint{
			endpoint.NewEndpoint("update-test.zone-1.local", endpoint.RecordTypeA, "8.8.8.8"),
			endpoint.NewEndpoint("update-test.zone-2.local", endpoint.RecordTypeA, "8.8.4.4"),
			endpoint.NewEndpoint("update-test-cname.zone-1.local", endpoint.RecordTypeCNAME, "update-test-cname"),
		},
		UpdateNew: []*endpoint.Endpoint{
			endpoint.NewEndpoint("update-test.zone-1.local", endpoint.RecordTypeA, "1.2.3.4"),
			endpoint.NewEndpoint("update-test.zone-2.local", endpoint.RecordTypeA, "4.3.2.1"),
			endpoint.NewEndpoint("update-test-cname.zone-1.local", endpoint.RecordTypeCNAME, "updated-test-cname"),
		},
		Delete: []*endpoint.Endpoint{
			endpoint.NewEndpoint("delete-test.zone-1.local", endpoint.RecordTypeA, "8.8.8.8"),
			endpoint.NewEndpoint("delete-test.zone-2.local", endpoint.RecordTypeA, "8.8.4.4"),
			endpoint.NewEndpoint("delete-test-cname.zone-1.local", endpoint.RecordTypeCNAME, "delete-test-cname"),
		},
	}))
	mockClients.EqualRecords(&records)
}

func TestGoogleApplyChangesEmpty(t *testing.T) {
	p := newGoogleProvider().WithMockClients(newMockClients(t))
	assert.NoError(t, p.ApplyChanges(context.Background(), &plan.Changes{}))
}

func TestNewChange(t *testing.T) {
	p := newGoogleProvider().WithMockClients(newMockClients(t))

	records := []*dns.ResourceRecordSet{}
	for _, ep := range []*endpoint.Endpoint{
		endpoint.NewEndpointWithTTL("update-test.zone-2.local", endpoint.RecordTypeA, 1, "8.8.4.4"),
		endpoint.NewEndpointWithTTL("delete-test.zone-2.local", endpoint.RecordTypeA, 120, "8.8.4.4"),
		endpoint.NewEndpointWithTTL("update-test-cname.zone-1.local", endpoint.RecordTypeCNAME, 4000, "update-test-cname"),
		// test fallback to Ttl:300 when Ttl==0 :
		endpoint.NewEndpointWithTTL("update-test.zone-1.local", endpoint.RecordTypeA, 0, "8.8.8.8"),
		endpoint.NewEndpointWithTTL("update-test-mx.zone-1.local", endpoint.RecordTypeMX, 6000, "10 mail"),
		endpoint.NewEndpoint("delete-test.zone-1.local", endpoint.RecordTypeA, "8.8.8.8"),
		endpoint.NewEndpoint("delete-test-cname.zone-1.local", endpoint.RecordTypeCNAME, "delete-test-cname"),
	} {
		records = append(records, p.newChange(&plan.RRSetChange{
			Name: plan.RRName(provider.EnsureTrailingDot(ep.DNSName)),
			Type: plan.RRType(ep.RecordType),
			Create: []*endpoint.Endpoint{ep},
		}).Additions...)
	}

	validateChangeRecords(t, records, []*dns.ResourceRecordSet{
		{Name: "update-test.zone-2.local.", Rrdatas: []string{"8.8.4.4"}, Type: "A", Ttl: 1},
		{Name: "delete-test.zone-2.local.", Rrdatas: []string{"8.8.4.4"}, Type: "A", Ttl: 120},
		{Name: "update-test-cname.zone-1.local.", Rrdatas: []string{"update-test-cname."}, Type: "CNAME", Ttl: 4000},
		{Name: "update-test.zone-1.local.", Rrdatas: []string{"8.8.8.8"}, Type: "A", Ttl: 300},
		{Name: "update-test-mx.zone-1.local.", Rrdatas: []string{"10 mail."}, Type: "MX", Ttl: 6000},
		{Name: "delete-test.zone-1.local.", Rrdatas: []string{"8.8.8.8"}, Type: "A", Ttl: 300},
		{Name: "delete-test-cname.zone-1.local.", Rrdatas: []string{"delete-test-cname."}, Type: "CNAME", Ttl: 300},
	})
}

func TestGoogleBatchChangeSet(t *testing.T) {
	mockClients := newMockClients(t)
	mockClients.changesClient.maxChangeSize = 1
	p := newGoogleProvider().WithMockClients(mockClients.WithRecords(
		&map[*dns.ManagedZone][]*dns.ResourceRecordSet{
			&dns.ManagedZone{
				Name:       "zone",
				DnsName:    "zone.local.",
			}: []*dns.ResourceRecordSet{},
		},
	))
	p.batchChangeSize = mockClients.changesClient.maxChangeSize
	changes := plan.Changes{
		Create: make([]*endpoint.Endpoint, mockClients.changesClient.maxChangeSize + 1),
	}
	for i := 0; i < len(changes.Create); i += 1 {
		changes.Create[i] = endpoint.NewEndpointWithTTL(fmt.Sprintf("record%d.zone.local.", i), endpoint.RecordTypeA, 1, "1.2.3.4")
	}
	ctx := context.Background()
	require.NoError(t, p.ApplyChanges(ctx, &changes))
	records, err := p.Records(ctx)
	require.NoError(t, err)
	validateEndpoints(t, records, changes.Create)
}

func TestSoftErrListZones(t *testing.T) {
	mockClients := newMockClients(t)
	mockClients.managedZonesClient.zonesErr = provider.NewSoftError(fmt.Errorf("failed to list zones"))
	p := newGoogleProvider().WithMockClients(mockClients)
	zones, err := p.Zones(context.Background())
	require.Error(t, err)
	require.ErrorIs(t, err, provider.SoftError)
	require.Empty(t, zones)
}

func TestSoftErrListRecords(t *testing.T) {
	mockClients := newMockClients(t)
	mockClients.resourceRecordSetsClient.recordsErr = provider.NewSoftError(fmt.Errorf("failed to list records in zone"))
	p := newGoogleProvider().WithMockClients(mockClients.WithRecords(
		&map[*dns.ManagedZone][]*dns.ResourceRecordSet{
			&dns.ManagedZone{
				Name:       "zone",
				DnsName:    "zone.local.",
			}: []*dns.ResourceRecordSet{},
		},
	))
	records, err := p.Records(context.Background())
	require.Error(t, err)
	require.ErrorIs(t, err, provider.SoftError)
	require.Empty(t, records)
}

func sortChangesByName(cs *dns.Change) {
	sort.SliceStable(cs.Additions, func(i, j int) bool {
		return cs.Additions[i].Name < cs.Additions[j].Name
	})

	sort.SliceStable(cs.Deletions, func(i, j int) bool {
		return cs.Deletions[i].Name < cs.Deletions[j].Name
	})
}

func validateChangeRecords(t *testing.T, records []*dns.ResourceRecordSet, expected []*dns.ResourceRecordSet) {
	require.Len(t, records, len(expected))

	for i := range records {
		validateChangeRecord(t, records[i], expected[i])
	}
}

func validateChangeRecord(t *testing.T, record *dns.ResourceRecordSet, expected *dns.ResourceRecordSet) {
	assert.Equal(t, expected.Name, record.Name)
	assert.Equal(t, expected.Rrdatas, record.Rrdatas)
	assert.Equal(t, expected.Ttl, record.Ttl)
	assert.Equal(t, expected.Type, record.Type)
}

func validateEndpoints(t *testing.T, endpoints []*endpoint.Endpoint, expected []*endpoint.Endpoint) {
	assert.True(t, testutils.SameEndpoints(endpoints, expected), "actual and expected endpoints don't match. %s:%s", endpoints, expected)
}

type mockChangesClient struct{
	records *map[*dns.ManagedZone][]*dns.ResourceRecordSet
	maxChangeSize int
}

type mockChangesCreateCall struct {
	client *mockChangesClient
	managedZone string
	change *dns.Change
	maxChangeSize int
}

func (m *mockChangesCreateCall) Do(opts ...googleapi.CallOption) (*dns.Change, error) {
	var managedZone *dns.ManagedZone
	for zone := range maps.Keys(*m.client.records) {
		if zone.Name == m.managedZone {
			managedZone = zone
			break
		}
	}
	if managedZone == nil {
		return nil, &googleapi.Error{Code: http.StatusNotFound}
	}
	size := len(m.change.Additions) + len(m.change.Deletions)
	if m.maxChangeSize > 0 && size > m.maxChangeSize {
		return nil, &googleapi.Error{
			Code: http.StatusBadRequest,
			Message: fmt.Sprintf("change size %d exceeds maximum %d", size, m.maxChangeSize),
		}
	}
	for _, records := range [][]*dns.ResourceRecordSet{m.change.Additions, m.change.Deletions} {
		for _, record := range records {
			if !strings.HasSuffix(record.Name, managedZone.DnsName) {
				return nil, &googleapi.Error{
					Code: http.StatusBadRequest,
					Message: fmt.Sprintf("invalid name %s for zone %s", record.Name, managedZone.DnsName),
				}
			}
			if !isValidRecordSet(record) {
				return nil, &googleapi.Error{
					Code: http.StatusBadRequest,
					Message: fmt.Sprintf("invalid record: %v", record),
				}
			}
		}
	}
	for _, deletion := range m.change.Deletions {
		index := -1
		for i, record := range (*m.client.records)[managedZone] {
			if reflect.DeepEqual(record, deletion) {
				index = i
				break
			}
		}
		if index == -1 {
			return nil, &googleapi.Error{
				Code: http.StatusBadRequest,
				Message: fmt.Sprintf("record not found: %v", deletion),
			}
		}
		(*m.client.records)[managedZone] = append(
			(*m.client.records)[managedZone][:index],
			(*m.client.records)[managedZone][index+1:]...
		)
	}
	for _, addition := range m.change.Additions {
		for _, record := range (*m.client.records)[managedZone] {
			if record.Name == addition.Name && record.Type == addition.Type {
				return nil, &googleapi.Error{
					Code: http.StatusBadRequest,
					Message: fmt.Sprintf("record exists: %v", addition),
				}
			}
		}
		(*m.client.records)[managedZone] = append((*m.client.records)[managedZone], addition)
	}
	return m.change, nil
}

func (m *mockChangesClient) Create(managedZone string, change *dns.Change) changesCreateCallInterface {
	return &mockChangesCreateCall{client: m, managedZone: managedZone, change: change, maxChangeSize: m.maxChangeSize}
}

type mockResourceRecordSetsClient struct {
	records *map[*dns.ManagedZone][]*dns.ResourceRecordSet
	recordsErr error
}

type mockResourceRecordSetsListCall struct {
	client *mockResourceRecordSetsClient
	managedZone string
}

func (m *mockResourceRecordSetsListCall) Pages(ctx context.Context, f func(*dns.ResourceRecordSetsListResponse) error) error {
	if m.client.recordsErr != nil {
		return m.client.recordsErr
	}
	var managedZone *dns.ManagedZone
	for zone := range maps.Keys(*m.client.records) {
		if zone.Name == m.managedZone {
			managedZone = zone
			break
		}
	}
	if managedZone == nil {
		return &googleapi.Error{Code: http.StatusNotFound}
	}
	return f(&dns.ResourceRecordSetsListResponse{Rrsets: (*m.client.records)[managedZone]})
}

func (m *mockResourceRecordSetsClient) List(managedZone string) resourceRecordSetsListCallInterface {
	return &mockResourceRecordSetsListCall{client: m, managedZone: managedZone}
}

type mockManagedZonesClient struct {
	records *map[*dns.ManagedZone][]*dns.ResourceRecordSet
	zonesErr error
}

type mockManagedZonesCreateCall struct {
	client *mockManagedZonesClient
	managedZone *dns.ManagedZone
}

func (m *mockManagedZonesCreateCall) Do(opts ...googleapi.CallOption) (*dns.ManagedZone, error) {
	for zone := range maps.Keys(*m.client.records) {
		if zone.Name == m.managedZone.Name {
			return nil, &googleapi.Error{Code: http.StatusConflict}
		}
	}
	(*m.client.records)[m.managedZone] = []*dns.ResourceRecordSet{}
	return m.managedZone, nil
}

type mockManagedZonesListCall struct {
	client *mockManagedZonesClient
}

func (m *mockManagedZonesListCall) Pages(ctx context.Context, f func(*dns.ManagedZonesListResponse) error) error {
	if m.client.zonesErr != nil {
		return m.client.zonesErr
	}
	return f(&dns.ManagedZonesListResponse{ManagedZones: slices.Collect(maps.Keys(*m.client.records))})
}

func (m *mockManagedZonesClient) Create(managedZone *dns.ManagedZone) managedZonesCreateCallInterface {
	return &mockManagedZonesCreateCall{client: m, managedZone: managedZone}
}

func (m *mockManagedZonesClient) List() managedZonesListCallInterface {
	return &mockManagedZonesListCall{client: m}
}

func hasTrailingDot(target string) bool {
	return strings.HasSuffix(target, ".")
}

func isValidRecordSet(recordSet *dns.ResourceRecordSet) bool {
	if !hasTrailingDot(recordSet.Name) {
		return false
	}

	switch recordSet.Type {
	case endpoint.RecordTypeCNAME, endpoint.RecordTypeMX, endpoint.RecordTypeNS, endpoint.RecordTypeSRV:
		for _, rrd := range recordSet.Rrdatas {
			if !hasTrailingDot(rrd) {
				return false
			}
		}
	case endpoint.RecordTypeA, endpoint.RecordTypeTXT:
		for _, rrd := range recordSet.Rrdatas {
			if hasTrailingDot(rrd) {
				return false
			}
		}
	default:
		panic("unhandled record type")
	}

	return true
}

type mockClients struct {
	t *testing.T
	records *map[*dns.ManagedZone][]*dns.ResourceRecordSet
	changesClient *mockChangesClient
	managedZonesClient *mockManagedZonesClient
	resourceRecordSetsClient *mockResourceRecordSetsClient
}

func newMockClients(t *testing.T) *mockClients {
	records := map[*dns.ManagedZone][]*dns.ResourceRecordSet{}
	return &mockClients{
		t: t,
		records: &records,
		changesClient: &mockChangesClient{
			records: &records,
		},
		managedZonesClient: &mockManagedZonesClient{
			records: &records,
		},
		resourceRecordSetsClient: &mockResourceRecordSetsClient{
			records: &records,
		},
	}
}

func (c *mockClients) WithRecords(records *map[*dns.ManagedZone][]*dns.ResourceRecordSet) *mockClients {
	for managedZone, resourceRecordSets := range *records {
		if _, err := c.managedZonesClient.Create(managedZone).Do(); err != nil {
			if err, ok := err.(*googleapi.Error); !ok || err.Code != http.StatusConflict {
				require.NoError(c.t, err)
			}
		}
		change := &dns.Change{
			Additions: resourceRecordSets,
		}
		c.resourceRecordSetsClient.List(managedZone.Name).Pages(context.Background(), func(resp *dns.ResourceRecordSetsListResponse) error {
			change.Deletions = append(change.Deletions, resp.Rrsets...)
			return nil
		})
		if len(change.Additions) + len(change.Deletions) > 0 {
			_, err := c.changesClient.Create(managedZone.Name, change).Do()
			require.NoError(c.t, err)
		}
	}
	return c
}

func (c *mockClients) EqualRecords(records *map[*dns.ManagedZone][]*dns.ResourceRecordSet) {
	c.t.Helper()
	assert.Equal(c.t, len(*records), len(*c.records))
	for expectedZone, expectedRecords := range *records {
		var actualZone *dns.ManagedZone
		var actualRecords []*dns.ResourceRecordSet
		for searchZone, searchRecords := range *c.records {
			if reflect.DeepEqual(searchZone, expectedZone) {
				actualZone = searchZone
				actualRecords = searchRecords
				break
			}
		}
		if actualZone == nil {
			assert.ElementsMatch(c.t, slices.Collect(maps.Keys(*records)), slices.Collect(maps.Keys(*c.records)))
		}
		assert.ElementsMatch(c.t, expectedRecords, actualRecords)
	}
}

func newGoogleProvider() *GoogleProvider {
	return &GoogleProvider{}
}

func (p *GoogleProvider) WithMockClients(clients *mockClients) *GoogleProvider {
	p.resourceRecordSetsClient = clients.resourceRecordSetsClient
	p.managedZonesClient = clients.managedZonesClient
	p.changesClient = clients.changesClient
	return p
}
