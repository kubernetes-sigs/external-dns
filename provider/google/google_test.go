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
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	dns "google.golang.org/api/dns/v1"
	"google.golang.org/api/googleapi"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

var (
	testZones   = map[string]*dns.ManagedZone{}
	testRecords = map[string]map[string]*dns.ResourceRecordSet{}
)

type mockManagedZonesCreateCall struct {
	project     string
	managedZone *dns.ManagedZone
}

func (m *mockManagedZonesCreateCall) Do(_ ...googleapi.CallOption) (*dns.ManagedZone, error) {
	zoneKey := zoneKey(m.project, m.managedZone.Name)

	if _, ok := testZones[zoneKey]; ok {
		return nil, &googleapi.Error{Code: http.StatusConflict}
	}

	testZones[zoneKey] = m.managedZone

	return m.managedZone, nil
}

type mockManagedZonesListCall struct {
	project          string
	zonesListSoftErr error
}

func (m *mockManagedZonesListCall) Pages(_ context.Context, f func(*dns.ManagedZonesListResponse) error) error {
	zones := []*dns.ManagedZone{}

	for k, v := range testZones {
		if strings.HasPrefix(k, m.project+"/") {
			zones = append(zones, v)
		}
	}

	if m.zonesListSoftErr != nil {
		return m.zonesListSoftErr
	}

	return f(&dns.ManagedZonesListResponse{ManagedZones: zones})
}

type mockManagedZonesClient struct {
	project  string
	zonesErr error
}

func (m *mockManagedZonesClient) Create(managedZone *dns.ManagedZone) managedZonesCreateCallInterface {
	return &mockManagedZonesCreateCall{project: m.project, managedZone: managedZone}
}

func (m *mockManagedZonesClient) List() managedZonesListCallInterface {
	return &mockManagedZonesListCall{project: m.project, zonesListSoftErr: m.zonesErr}
}

type mockResourceRecordSetsGetCall struct {
	project     string
	managedZone string
	name        string
	type_       string
}

func (m *mockResourceRecordSetsGetCall) Do(_ ...googleapi.CallOption) (*dns.ResourceRecordSet, error) {
	zoneKey := zoneKey(m.project, m.managedZone)

	if records, ok := testRecords[zoneKey]; ok {
		if rrset, ok := records[recordKey(m.type_, m.name)]; ok {
			return rrset, nil
		}
	}

	return nil, &googleapi.Error{Code: http.StatusNotFound}
}

type mockResourceRecordSetsListCall struct {
	project            string
	managedZone        string
	recordsListSoftErr error
}

func (m *mockResourceRecordSetsListCall) Pages(_ context.Context, f func(*dns.ResourceRecordSetsListResponse) error) error {
	zoneKey := zoneKey(m.project, m.managedZone)

	if _, ok := testZones[zoneKey]; !ok {
		return &googleapi.Error{Code: http.StatusNotFound}
	}

	resp := []*dns.ResourceRecordSet{}

	for _, v := range testRecords[zoneKey] {
		resp = append(resp, v)
	}

	if m.recordsListSoftErr != nil {
		return m.recordsListSoftErr
	}

	return f(&dns.ResourceRecordSetsListResponse{Rrsets: resp})
}

type mockResourceRecordSetsClient struct {
	project    string
	recordsErr error
}

func (m *mockResourceRecordSetsClient) Get(managedZone string, name string, type_ string) resourceRecordSetsGetCallInterface {
	return &mockResourceRecordSetsGetCall{project: m.project, managedZone: managedZone, name: name, type_: type_}
}

func (m *mockResourceRecordSetsClient) List(managedZone string) resourceRecordSetsListCallInterface {
	return &mockResourceRecordSetsListCall{project: m.project, managedZone: managedZone, recordsListSoftErr: m.recordsErr}
}

type mockChangesCreateCall struct {
	project     string
	managedZone string
	change      *dns.Change
}

func (m *mockChangesCreateCall) Do(_ ...googleapi.CallOption) (*dns.Change, error) {
	zoneKey := zoneKey(m.project, m.managedZone)

	if _, ok := testZones[zoneKey]; !ok {
		return nil, &googleapi.Error{Code: http.StatusNotFound}
	}

	if _, ok := testRecords[zoneKey]; !ok {
		testRecords[zoneKey] = make(map[string]*dns.ResourceRecordSet)
	}

	for _, c := range append(m.change.Additions, m.change.Deletions...) {
		if !isValidRecordSet(c) {
			return nil, &googleapi.Error{
				Code:    http.StatusBadRequest,
				Message: fmt.Sprintf("invalid record: %v", c),
			}
		}
	}

	for _, del := range m.change.Deletions {
		recordKey := recordKey(del.Type, del.Name)
		delete(testRecords[zoneKey], recordKey)
	}

	for _, add := range m.change.Additions {
		recordKey := recordKey(add.Type, add.Name)
		testRecords[zoneKey][recordKey] = add
	}

	return m.change, nil
}

type mockChangesClient struct {
	project string
}

func (m *mockChangesClient) Create(managedZone string, change *dns.Change) changesCreateCallInterface {
	return &mockChangesCreateCall{project: m.project, managedZone: managedZone, change: change}
}

func zoneKey(project, zoneName string) string {
	return project + "/" + zoneName
}

func recordKey(recordType, recordName string) string {
	return recordType + "/" + recordName
}

func isValidRecordSet(recordSet *dns.ResourceRecordSet) bool {
	if !hasTrailingDot(recordSet.Name) {
		return false
	}

	switch recordSet.Type {
	case endpoint.RecordTypeCNAME:
		for _, rrd := range recordSet.Rrdatas {
			if !hasTrailingDot(rrd) {
				return false
			}
		}
	case endpoint.RecordTypeA, endpoint.RecordTypeTXT:
		if slices.ContainsFunc(recordSet.Rrdatas, hasTrailingDot) {
			return false
		}
	default:
		panic("unhandled record type")
	}

	return true
}

func hasTrailingDot(target string) bool {
	return strings.HasSuffix(target, ".")
}

func TestGoogleZonesIDFilter(t *testing.T) {
	provider := newGoogleProviderZoneOverlap(t, endpoint.NewDomainFilter([]string{"cluster.local."}), provider.NewZoneIDFilter([]string{"10002"}), provider.NewZoneTypeFilter(""), []*endpoint.Endpoint{})

	zones, err := provider.Zones(t.Context())
	require.NoError(t, err)

	validateZones(t, zones, map[string]*dns.ManagedZone{
		"internal-2": {Name: "internal-2", DnsName: "cluster.local.", Id: 10002, Visibility: "private"},
	})
}

func TestGoogleZonesNameFilter(t *testing.T) {
	provider := newGoogleProviderZoneOverlap(t, endpoint.NewDomainFilter([]string{"cluster.local."}), provider.NewZoneIDFilter([]string{"internal-2"}), provider.NewZoneTypeFilter(""), []*endpoint.Endpoint{})

	zones, err := provider.Zones(t.Context())
	require.NoError(t, err)

	validateZones(t, zones, map[string]*dns.ManagedZone{
		"internal-2": {Name: "internal-2", DnsName: "cluster.local.", Id: 10002, Visibility: "private"},
	})
}

func TestGoogleZonesVisibilityFilterPublic(t *testing.T) {
	provider := newGoogleProviderZoneOverlap(t, endpoint.NewDomainFilter([]string{"cluster.local."}), provider.NewZoneIDFilter([]string{"split-horizon-1"}), provider.NewZoneTypeFilter("public"), []*endpoint.Endpoint{})

	zones, err := provider.Zones(t.Context())
	require.NoError(t, err)

	validateZones(t, zones, map[string]*dns.ManagedZone{
		"split-horizon-1": {Name: "split-horizon-1", DnsName: "cluster.local.", Id: 10001, Visibility: "public"},
	})
}

func TestGoogleZonesVisibilityFilterPrivate(t *testing.T) {
	provider := newGoogleProviderZoneOverlap(t, endpoint.NewDomainFilter([]string{"cluster.local."}), provider.NewZoneIDFilter([]string{"split-horizon-1"}), provider.NewZoneTypeFilter("public"), []*endpoint.Endpoint{})

	zones, err := provider.Zones(t.Context())
	require.NoError(t, err)

	validateZones(t, zones, map[string]*dns.ManagedZone{
		"split-horizon-1": {Name: "split-horizon-1", DnsName: "cluster.local.", Id: 10001, Visibility: "public"},
	})
}

func TestGoogleZonesVisibilityFilterPrivatePeering(t *testing.T) {
	provider := newGoogleProviderZoneOverlap(t, endpoint.NewDomainFilter([]string{"svc.local."}), provider.NewZoneIDFilter([]string{""}), provider.NewZoneTypeFilter("private"), []*endpoint.Endpoint{})

	zones, err := provider.Zones(t.Context())
	require.NoError(t, err)

	validateZones(t, zones, map[string]*dns.ManagedZone{
		"svc-local": {Name: "svc-local", DnsName: "svc.local.", Id: 1005, Visibility: "private"},
	})
}

func TestGoogleRecords(t *testing.T) {
	originalEndpoints := []*endpoint.Endpoint{
		endpoint.NewEndpointWithTTL("list-test.zone-1.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, endpoint.TTL(1), "1.2.3.4"),
		endpoint.NewEndpointWithTTL("list-test.zone-2.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, endpoint.TTL(2), "8.8.8.8"),
		endpoint.NewEndpointWithTTL("list-test-alias.zone-1.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeCNAME, endpoint.TTL(3), "foo.elb.amazonaws.com"),
	}

	provider := newGoogleProvider(t, endpoint.NewDomainFilter([]string{"ext-dns-test-2.gcp.zalan.do."}), provider.NewZoneIDFilter([]string{""}), false, originalEndpoints, nil, nil)

	records, err := provider.Records(t.Context())
	require.NoError(t, err)

	validateEndpoints(t, records, originalEndpoints)
}

func TestGoogleRecordsFilter(t *testing.T) {
	originalEndpoints := []*endpoint.Endpoint{
		endpoint.NewEndpointWithTTL("update-test.zone-1.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, defaultTTL, "8.8.8.8"),
		endpoint.NewEndpointWithTTL("delete-test.zone-1.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, defaultTTL, "8.8.8.8"),
		endpoint.NewEndpointWithTTL("update-test.zone-2.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, defaultTTL, "8.8.4.4"),
		endpoint.NewEndpointWithTTL("delete-test.zone-2.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, defaultTTL, "8.8.4.4"),
		endpoint.NewEndpointWithTTL("update-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeCNAME, defaultTTL, "bar.elb.amazonaws.com"),
		endpoint.NewEndpointWithTTL("delete-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeCNAME, defaultTTL, "qux.elb.amazonaws.com"),
	}

	provider := newGoogleProvider(
		t,
		endpoint.NewDomainFilter([]string{
			// our two valid zones
			"zone-1.ext-dns-test-2.gcp.zalan.do.",
			"zone-2.ext-dns-test-2.gcp.zalan.do.",
			// we filter for a zone that doesn't exist, should have no effect.
			"zone-0.ext-dns-test-2.gcp.zalan.do.",
			// there exists a third zone "zone-3" that we want to exclude from being managed.
		}),
		provider.NewZoneIDFilter([]string{""}),
		false,
		originalEndpoints,
		nil,
		nil,
	)

	// these records should be filtered out since they don't match a hosted zone or domain filter.
	ignoredEndpoints := []*endpoint.Endpoint{
		endpoint.NewEndpoint("filter-create-test.zone-0.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, "4.2.2.2"),
		endpoint.NewEndpoint("filter-update-test.zone-0.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, "4.2.2.2"),
		endpoint.NewEndpoint("filter-delete-test.zone-0.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, "4.2.2.2"),
		endpoint.NewEndpoint("filter-create-test.zone-3.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, "4.2.2.2"),
		endpoint.NewEndpoint("filter-update-test.zone-3.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, "4.2.2.2"),
		endpoint.NewEndpoint("filter-delete-test.zone-3.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, "4.2.2.2"),
	}

	require.NoError(t, provider.ApplyChanges(t.Context(), &plan.Changes{
		Create: ignoredEndpoints,
	}))

	records, err := provider.Records(t.Context())
	require.NoError(t, err)

	// assert that due to filtering no changes were made.
	validateEndpoints(t, records, originalEndpoints)
}

func TestGoogleApplyChanges(t *testing.T) {
	provider := newGoogleProvider(
		t,
		endpoint.NewDomainFilter([]string{
			// our two valid zones
			"zone-1.ext-dns-test-2.gcp.zalan.do.",
			"zone-2.ext-dns-test-2.gcp.zalan.do.",
			// we filter for a zone that doesn't exist, should have no effect.
			"zone-0.ext-dns-test-2.gcp.zalan.do.",
			// there exists a third zone "zone-3" that we want to exclude from being managed.
		}),
		provider.NewZoneIDFilter([]string{""}),
		false,
		[]*endpoint.Endpoint{
			endpoint.NewEndpointWithTTL("update-test.zone-1.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, defaultTTL, "8.8.8.8"),
			endpoint.NewEndpointWithTTL("delete-test.zone-1.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, defaultTTL, "8.8.8.8"),
			endpoint.NewEndpointWithTTL("update-test-ttl.zone-2.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, endpoint.TTL(10), "8.8.4.4"),
			endpoint.NewEndpointWithTTL("delete-test.zone-2.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, defaultTTL, "8.8.4.4"),
			endpoint.NewEndpointWithTTL("update-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeCNAME, defaultTTL, "bar.elb.amazonaws.com"),
			endpoint.NewEndpointWithTTL("delete-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeCNAME, defaultTTL, "qux.elb.amazonaws.com"),
		},
		nil,
		nil,
	)

	createRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("create-test.zone-1.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, "8.8.8.8"),
		endpoint.NewEndpointWithTTL("create-test-ttl.zone-2.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, endpoint.TTL(15), "8.8.4.4"),
		endpoint.NewEndpoint("create-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeCNAME, "foo.elb.amazonaws.com"),
		endpoint.NewEndpoint("filter-create-test.zone-3.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, "4.2.2.2"),
		endpoint.NewEndpoint("nomatch-create-test.zone-0.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, "4.2.2.1"),
	}

	currentRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("update-test.zone-1.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, "8.8.8.8"),
		endpoint.NewEndpoint("update-test.zone-2.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, "8.8.4.4"),
		endpoint.NewEndpoint("update-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeCNAME, "bar.elb.amazonaws.com"),
		endpoint.NewEndpoint("filter-update-test.zone-3.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, "4.2.2.2"),
	}
	updatedRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("update-test.zone-1.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, "1.2.3.4"),
		endpoint.NewEndpointWithTTL("update-test-ttl.zone-2.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, endpoint.TTL(25), "4.3.2.1"),
		endpoint.NewEndpoint("update-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeCNAME, "baz.elb.amazonaws.com"),
		endpoint.NewEndpoint("filter-update-test.zone-3.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, "5.6.7.8"),
		endpoint.NewEndpoint("nomatch-update-test.zone-0.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, "8.7.6.5"),
	}

	deleteRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("delete-test.zone-1.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, "8.8.8.8"),
		endpoint.NewEndpoint("delete-test.zone-2.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, "8.8.4.4"),
		endpoint.NewEndpoint("delete-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeCNAME, "qux.elb.amazonaws.com"),
		endpoint.NewEndpoint("filter-delete-test.zone-3.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, "4.2.2.2"),
		endpoint.NewEndpoint("nomatch-delete-test.zone-0.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, "4.2.2.1"),
	}

	changes := &plan.Changes{
		Create:    createRecords,
		UpdateNew: updatedRecords,
		UpdateOld: currentRecords,
		Delete:    deleteRecords,
	}

	require.NoError(t, provider.ApplyChanges(t.Context(), changes))

	records, err := provider.Records(t.Context())
	require.NoError(t, err)

	validateEndpoints(t, records, []*endpoint.Endpoint{
		endpoint.NewEndpointWithTTL("create-test.zone-1.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, defaultTTL, "8.8.8.8"),
		endpoint.NewEndpointWithTTL("update-test.zone-1.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, defaultTTL, "1.2.3.4"),
		endpoint.NewEndpointWithTTL("create-test-ttl.zone-2.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, endpoint.TTL(15), "8.8.4.4"),
		endpoint.NewEndpointWithTTL("update-test-ttl.zone-2.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, endpoint.TTL(25), "4.3.2.1"),
		endpoint.NewEndpointWithTTL("create-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeCNAME, defaultTTL, "foo.elb.amazonaws.com"),
		endpoint.NewEndpointWithTTL("update-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeCNAME, defaultTTL, "baz.elb.amazonaws.com"),
	})
}

func TestGoogleApplyChangesDryRun(t *testing.T) {
	originalEndpoints := []*endpoint.Endpoint{
		endpoint.NewEndpointWithTTL("update-test.zone-1.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, defaultTTL, "8.8.8.8"),
		endpoint.NewEndpointWithTTL("delete-test.zone-1.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, defaultTTL, "8.8.8.8"),
		endpoint.NewEndpointWithTTL("update-test.zone-2.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, defaultTTL, "8.8.4.4"),
		endpoint.NewEndpointWithTTL("delete-test.zone-2.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, defaultTTL, "8.8.4.4"),
		endpoint.NewEndpointWithTTL("update-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeCNAME, defaultTTL, "bar.elb.amazonaws.com"),
		endpoint.NewEndpointWithTTL("delete-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeCNAME, defaultTTL, "qux.elb.amazonaws.com"),
	}

	provider := newGoogleProvider(t, endpoint.NewDomainFilter([]string{"ext-dns-test-2.gcp.zalan.do."}), provider.NewZoneIDFilter([]string{""}), true, originalEndpoints, nil, nil)

	createRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("create-test.zone-1.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, "8.8.8.8"),
		endpoint.NewEndpoint("create-test.zone-2.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, "8.8.4.4"),
		endpoint.NewEndpoint("create-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeCNAME, "foo.elb.amazonaws.com"),
	}

	currentRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("update-test.zone-1.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, "8.8.8.8"),
		endpoint.NewEndpoint("update-test.zone-2.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, "8.8.4.4"),
		endpoint.NewEndpoint("update-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeCNAME, "bar.elb.amazonaws.com"),
	}
	updatedRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("update-test.zone-1.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, "1.2.3.4"),
		endpoint.NewEndpoint("update-test.zone-2.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, "4.3.2.1"),
		endpoint.NewEndpoint("update-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeCNAME, "baz.elb.amazonaws.com"),
	}

	deleteRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("delete-test.zone-1.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, "8.8.8.8"),
		endpoint.NewEndpoint("delete-test.zone-2.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, "8.8.4.4"),
		endpoint.NewEndpoint("delete-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeCNAME, "qux.elb.amazonaws.com"),
	}

	changes := &plan.Changes{
		Create:    createRecords,
		UpdateNew: updatedRecords,
		UpdateOld: currentRecords,
		Delete:    deleteRecords,
	}

	ctx := t.Context()
	require.NoError(t, provider.ApplyChanges(ctx, changes))

	records, err := provider.Records(ctx)
	require.NoError(t, err)

	validateEndpoints(t, records, originalEndpoints)
}

func TestGoogleApplyChangesEmpty(t *testing.T) {
	provider := newGoogleProvider(t, endpoint.NewDomainFilter([]string{"ext-dns-test-2.gcp.zalan.do."}), provider.NewZoneIDFilter([]string{""}), false, []*endpoint.Endpoint{}, nil, nil)
	require.NoError(t, provider.ApplyChanges(t.Context(), &plan.Changes{}))
}

func TestSoftErrListZonesConflict(t *testing.T) {
	p := newGoogleProvider(t, endpoint.NewDomainFilter([]string{"ext-dns-test-2.gcp.zalan.do."}), provider.NewZoneIDFilter([]string{}), false, []*endpoint.Endpoint{}, provider.NewSoftErrorf("failed to list zones"), nil)

	zones, err := p.Zones(t.Context())
	require.Error(t, err)
	require.ErrorIs(t, err, provider.SoftError)

	require.Empty(t, zones)
}

func TestSoftErrListRecordsConflict(t *testing.T) {
	p := newGoogleProvider(t, endpoint.NewDomainFilter([]string{"ext-dns-test-2.gcp.zalan.do."}), provider.NewZoneIDFilter([]string{}), false, []*endpoint.Endpoint{}, nil, provider.NewSoftErrorf("failed to list records in zone"))

	records, err := p.Records(t.Context())
	require.Error(t, err)
	require.ErrorIs(t, err, provider.SoftError)

	require.Empty(t, records)
}

func TestGoogleApplyChangesGeoRoutingPolicyCreate(t *testing.T) {
	p := newGoogleProvider(t, endpoint.NewDomainFilter([]string{"ext-dns-test-2.gcp.zalan.do."}), provider.NewZoneIDFilter([]string{""}), false, []*endpoint.Endpoint{}, nil, nil)

	geoUS := endpoint.NewEndpointWithTTL("geo-test.zone-1.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, defaultTTL, "1.1.1.1").
		WithProviderSpecific(providerSpecificRoutingPolicy, providerSpecificRoutingPolicyGeo).
		WithProviderSpecific(providerSpecificLocation, "us-east1").
		WithSetIdentifier("us-east1")
	geoEU := endpoint.NewEndpointWithTTL("geo-test.zone-1.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, defaultTTL, "2.2.2.2").
		WithProviderSpecific(providerSpecificRoutingPolicy, providerSpecificRoutingPolicyGeo).
		WithProviderSpecific(providerSpecificLocation, "europe-west1").
		WithSetIdentifier("europe-west1")

	require.NoError(t, p.ApplyChanges(t.Context(), &plan.Changes{
		Create: []*endpoint.Endpoint{geoUS, geoEU},
	}))

	records, err := p.Records(t.Context())
	require.NoError(t, err)
	validateEndpoints(t, records, []*endpoint.Endpoint{geoUS, geoEU})
}

func TestGoogleApplyChangesGeoRoutingPolicyUpdateOneLocation(t *testing.T) {
	p := newGoogleProvider(t, endpoint.NewDomainFilter([]string{"ext-dns-test-2.gcp.zalan.do."}), provider.NewZoneIDFilter([]string{""}), false, []*endpoint.Endpoint{}, nil, nil)

	geoUSOld := endpoint.NewEndpointWithTTL("geo-update-test.zone-1.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, defaultTTL, "1.1.1.1").
		WithProviderSpecific(providerSpecificRoutingPolicy, providerSpecificRoutingPolicyGeo).
		WithProviderSpecific(providerSpecificLocation, "us-east1").
		WithSetIdentifier("us-east1")
	geoEU := endpoint.NewEndpointWithTTL("geo-update-test.zone-1.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, defaultTTL, "2.2.2.2").
		WithProviderSpecific(providerSpecificRoutingPolicy, providerSpecificRoutingPolicyGeo).
		WithProviderSpecific(providerSpecificLocation, "europe-west1").
		WithSetIdentifier("europe-west1")

	require.NoError(t, p.ApplyChanges(t.Context(), &plan.Changes{
		Create: []*endpoint.Endpoint{geoUSOld, geoEU},
	}))

	// Change only the us-east1 item's target; europe-west1 must survive untouched
	// since Google's RRSet is a single object covering every location.
	geoUSNew := endpoint.NewEndpointWithTTL("geo-update-test.zone-1.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, defaultTTL, "9.9.9.9").
		WithProviderSpecific(providerSpecificRoutingPolicy, providerSpecificRoutingPolicyGeo).
		WithProviderSpecific(providerSpecificLocation, "us-east1").
		WithSetIdentifier("us-east1")

	require.NoError(t, p.ApplyChanges(t.Context(), &plan.Changes{
		UpdateOld: []*endpoint.Endpoint{geoUSOld},
		UpdateNew: []*endpoint.Endpoint{geoUSNew},
	}))

	records, err := p.Records(t.Context())
	require.NoError(t, err)
	validateEndpoints(t, records, []*endpoint.Endpoint{geoUSNew, geoEU})
}

func TestGoogleApplyChangesGeoRoutingPolicyDeleteLastLocationRemovesRRSet(t *testing.T) {
	p := newGoogleProvider(t, endpoint.NewDomainFilter([]string{"ext-dns-test-2.gcp.zalan.do."}), provider.NewZoneIDFilter([]string{""}), false, []*endpoint.Endpoint{}, nil, nil)

	geoUS := endpoint.NewEndpointWithTTL("geo-delete-test.zone-1.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, defaultTTL, "1.1.1.1").
		WithProviderSpecific(providerSpecificRoutingPolicy, providerSpecificRoutingPolicyGeo).
		WithProviderSpecific(providerSpecificLocation, "us-east1").
		WithSetIdentifier("us-east1")

	require.NoError(t, p.ApplyChanges(t.Context(), &plan.Changes{
		Create: []*endpoint.Endpoint{geoUS},
	}))

	require.NoError(t, p.ApplyChanges(t.Context(), &plan.Changes{
		Delete: []*endpoint.Endpoint{geoUS},
	}))

	records, err := p.Records(t.Context())
	require.NoError(t, err)
	for _, r := range records {
		assert.NotEqual(t, "geo-delete-test.zone-1.ext-dns-test-2.gcp.zalan.do", r.DNSName)
	}
}

func TestGoogleApplyChangesWrrRoutingPolicyCreate(t *testing.T) {
	p := newGoogleProvider(t, endpoint.NewDomainFilter([]string{"ext-dns-test-2.gcp.zalan.do."}), provider.NewZoneIDFilter([]string{""}), false, []*endpoint.Endpoint{}, nil, nil)

	wrr0 := endpoint.NewEndpointWithTTL("wrr-test.zone-1.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, defaultTTL, "1.1.1.1").
		WithProviderSpecific(providerSpecificRoutingPolicy, providerSpecificRoutingPolicyWrr).
		WithProviderSpecific(providerSpecificWeight, "30").
		WithSetIdentifier("0")
	wrr1 := endpoint.NewEndpointWithTTL("wrr-test.zone-1.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, defaultTTL, "2.2.2.2").
		WithProviderSpecific(providerSpecificRoutingPolicy, providerSpecificRoutingPolicyWrr).
		WithProviderSpecific(providerSpecificWeight, "70").
		WithSetIdentifier("1")

	require.NoError(t, p.ApplyChanges(t.Context(), &plan.Changes{
		Create: []*endpoint.Endpoint{wrr0, wrr1},
	}))

	records, err := p.Records(t.Context())
	require.NoError(t, err)
	validateEndpoints(t, records, []*endpoint.Endpoint{wrr0, wrr1})
}

func TestGoogleApplyChangesWrrRoutingPolicyShrinkCompactsItems(t *testing.T) {
	p := newGoogleProvider(t, endpoint.NewDomainFilter([]string{"ext-dns-test-2.gcp.zalan.do."}), provider.NewZoneIDFilter([]string{""}), false, []*endpoint.Endpoint{}, nil, nil)

	wrr0 := endpoint.NewEndpointWithTTL("wrr-shrink-test.zone-1.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, defaultTTL, "1.1.1.1").
		WithProviderSpecific(providerSpecificRoutingPolicy, providerSpecificRoutingPolicyWrr).
		WithProviderSpecific(providerSpecificWeight, "33").
		WithSetIdentifier("0")
	wrr1 := endpoint.NewEndpointWithTTL("wrr-shrink-test.zone-1.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, defaultTTL, "2.2.2.2").
		WithProviderSpecific(providerSpecificRoutingPolicy, providerSpecificRoutingPolicyWrr).
		WithProviderSpecific(providerSpecificWeight, "33").
		WithSetIdentifier("1")
	wrr2 := endpoint.NewEndpointWithTTL("wrr-shrink-test.zone-1.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, defaultTTL, "3.3.3.3").
		WithProviderSpecific(providerSpecificRoutingPolicy, providerSpecificRoutingPolicyWrr).
		WithProviderSpecific(providerSpecificWeight, "34").
		WithSetIdentifier("2")

	require.NoError(t, p.ApplyChanges(t.Context(), &plan.Changes{
		Create: []*endpoint.Endpoint{wrr0, wrr1, wrr2},
	}))

	// Drop the last item (index 2). The remaining items must still decode cleanly,
	// exercising rrSetRoutingPolicyWrrItemPlaceholder's index-compaction path.
	require.NoError(t, p.ApplyChanges(t.Context(), &plan.Changes{
		Delete: []*endpoint.Endpoint{wrr2},
	}))

	records, err := p.Records(t.Context())
	require.NoError(t, err)
	validateEndpoints(t, records, []*endpoint.Endpoint{wrr0, wrr1})
}

func TestGoogleAdjustEndpointsWrrDefaultsWeight(t *testing.T) {
	p := newGoogleProvider(t, endpoint.NewDomainFilter([]string{"ext-dns-test-2.gcp.zalan.do."}), provider.NewZoneIDFilter([]string{""}), false, []*endpoint.Endpoint{}, nil, nil)

	ep := endpoint.NewEndpointWithTTL("wrr-adjust-test.zone-1.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, defaultTTL, "1.1.1.1").
		WithProviderSpecific(providerSpecificRoutingPolicy, providerSpecificRoutingPolicyWrr).
		WithSetIdentifier("0")

	adjusted, err := p.AdjustEndpoints([]*endpoint.Endpoint{ep})
	require.NoError(t, err)
	require.Len(t, adjusted, 1)

	weight, ok := adjusted[0].GetProviderSpecificProperty(providerSpecificWeight)
	require.True(t, ok)
	assert.Equal(t, "100", weight)
}

func TestGoogleAdjustEndpointsGeoDefaultsLocationFromProvider(t *testing.T) {
	p := newGoogleProvider(t, endpoint.NewDomainFilter([]string{"ext-dns-test-2.gcp.zalan.do."}), provider.NewZoneIDFilter([]string{""}), false, []*endpoint.Endpoint{}, nil, nil)
	p.location = "us-central1"

	ep := endpoint.NewEndpointWithTTL("geo-adjust-test.zone-1.ext-dns-test-2.gcp.zalan.do", endpoint.RecordTypeA, defaultTTL, "1.1.1.1").
		WithProviderSpecific(providerSpecificRoutingPolicy, providerSpecificRoutingPolicyGeo)

	adjusted, err := p.AdjustEndpoints([]*endpoint.Endpoint{ep})
	require.NoError(t, err)
	require.Len(t, adjusted, 1)

	location, ok := adjusted[0].GetProviderSpecificProperty(providerSpecificLocation)
	require.True(t, ok)
	assert.Equal(t, "us-central1", location)
	assert.Equal(t, "us-central1", adjusted[0].SetIdentifier)
}

func validateZones(t *testing.T, zones map[string]*dns.ManagedZone, expected map[string]*dns.ManagedZone) {
	require.Len(t, zones, len(expected))

	for i, zone := range zones {
		validateZone(t, zone, expected[i])
	}
}

func validateZone(t *testing.T, zone *dns.ManagedZone, expected *dns.ManagedZone) {
	assert.Equal(t, expected.Name, zone.Name)
	assert.Equal(t, expected.DnsName, zone.DnsName)
	assert.Equal(t, expected.Visibility, zone.Visibility)
}

func newGoogleProviderZoneOverlap(t *testing.T, domainFilter *endpoint.DomainFilter, zoneIDFilter provider.ZoneIDFilter, zoneTypeFilter provider.ZoneTypeFilter, _ []*endpoint.Endpoint) *GoogleProvider {
	provider := &GoogleProvider{
		dryRun:                   false,
		domainFilter:             domainFilter,
		zoneIDFilter:             zoneIDFilter,
		zoneTypeFilter:           zoneTypeFilter,
		resourceRecordSetsClient: &mockResourceRecordSetsClient{project: "zalando-external-dns-test"},
		managedZonesClient:       &mockManagedZonesClient{project: "zalando-external-dns-test"},
		changesClient:            &mockChangesClient{project: "zalando-external-dns-test"},
	}

	createZone(t, provider, &dns.ManagedZone{
		Name:       "internal-1",
		DnsName:    "cluster.local.",
		Id:         10001,
		Visibility: "private",
	})

	createZone(t, provider, &dns.ManagedZone{
		Name:       "internal-2",
		DnsName:    "cluster.local.",
		Id:         10002,
		Visibility: "private",
	})

	createZone(t, provider, &dns.ManagedZone{
		Name:       "internal-3",
		DnsName:    "cluster.local.",
		Id:         10003,
		Visibility: "private",
	})

	createZone(t, provider, &dns.ManagedZone{
		Name:       "split-horizon-1",
		DnsName:    "cluster.local.",
		Id:         10004,
		Visibility: "public",
	})

	createZone(t, provider, &dns.ManagedZone{
		Name:       "split-horizon-1",
		DnsName:    "cluster.local.",
		Id:         10004,
		Visibility: "private",
	})

	createZone(t, provider, &dns.ManagedZone{
		Name:       "svc-local",
		DnsName:    "svc.local.",
		Id:         10005,
		Visibility: "private",
	})

	createZone(t, provider, &dns.ManagedZone{
		Name:          "svc-local-peer",
		DnsName:       "svc.local.",
		Id:            10006,
		Visibility:    "private",
		PeeringConfig: &dns.ManagedZonePeeringConfig{TargetNetwork: nil},
	})

	return provider
}

func newGoogleProvider(t *testing.T, domainFilter *endpoint.DomainFilter, zoneIDFilter provider.ZoneIDFilter, dryRun bool, records []*endpoint.Endpoint, zonesErr, recordsErr error) *GoogleProvider {
	provider := &GoogleProvider{
		dryRun:       false,
		domainFilter: domainFilter,
		zoneIDFilter: zoneIDFilter,
		resourceRecordSetsClient: &mockResourceRecordSetsClient{
			project:    "zalando-external-dns-test",
			recordsErr: recordsErr,
		},
		managedZonesClient: &mockManagedZonesClient{
			project:  "zalando-external-dns-test",
			zonesErr: zonesErr,
		},
		changesClient: &mockChangesClient{project: "zalando-external-dns-test"},
	}

	createZone(t, provider, &dns.ManagedZone{
		Name:    "zone-1-ext-dns-test-2-gcp-zalan-do",
		DnsName: "zone-1.ext-dns-test-2.gcp.zalan.do.",
	})

	createZone(t, provider, &dns.ManagedZone{
		Name:    "zone-2-ext-dns-test-2-gcp-zalan-do",
		DnsName: "zone-2.ext-dns-test-2.gcp.zalan.do.",
	})

	createZone(t, provider, &dns.ManagedZone{
		Name:    "zone-3-ext-dns-test-2-gcp-zalan-do",
		DnsName: "zone-3.ext-dns-test-2.gcp.zalan.do.",
	})

	// filtered out by domain filter
	createZone(t, provider, &dns.ManagedZone{
		Name:    "zone-4-ext-dns-test-3-gcp-zalan-do",
		DnsName: "zone-4.ext-dns-test-3.gcp.zalan.do.",
	})

	setupGoogleRecords(t, provider, records)

	provider.dryRun = dryRun

	return provider
}

func createZone(t *testing.T, p *GoogleProvider, zone *dns.ManagedZone) {
	zone.Description = "Testing zone for kubernetes.io/external-dns"

	if _, err := p.managedZonesClient.Create(zone).Do(); err != nil {
		var errs *googleapi.Error
		if !errors.As(err, &errs) || errs.Code != http.StatusConflict {
			require.NoError(t, err)
		}
	}
}

func setupGoogleRecords(t *testing.T, provider *GoogleProvider, endpoints []*endpoint.Endpoint) {
	clearGoogleRecords(t, provider, "zone-1-ext-dns-test-2-gcp-zalan-do")
	clearGoogleRecords(t, provider, "zone-2-ext-dns-test-2-gcp-zalan-do")
	clearGoogleRecords(t, provider, "zone-3-ext-dns-test-2-gcp-zalan-do")

	ctx := t.Context()
	records, _ := provider.Records(ctx)

	validateEndpoints(t, records, []*endpoint.Endpoint{})

	require.NoError(t, provider.ApplyChanges(t.Context(), &plan.Changes{
		Create: endpoints,
	}))

	records, _ = provider.Records(ctx)

	validateEndpoints(t, records, endpoints)
}

func clearGoogleRecords(t *testing.T, provider *GoogleProvider, zone string) {
	recordSets := []*dns.ResourceRecordSet{}

	provider.resourceRecordSetsClient.List(zone).Pages(t.Context(), func(resp *dns.ResourceRecordSetsListResponse) error {
		for _, r := range resp.Rrsets {
			switch r.Type {
			case endpoint.RecordTypeA, endpoint.RecordTypeCNAME:
				recordSets = append(recordSets, r)
			}
		}
		return nil
	})

	if len(recordSets) != 0 {
		_, err := provider.changesClient.Create(zone, &dns.Change{
			Deletions: recordSets,
		}).Do()
		require.NoError(t, err)
	}
}

func validateEndpoints(t *testing.T, endpoints []*endpoint.Endpoint, expected []*endpoint.Endpoint) {
	assert.True(t, testutils.SameEndpoints(endpoints, expected), "actual and expected endpoints don't match. %s:%s", endpoints, expected)
}
