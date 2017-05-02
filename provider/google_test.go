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

package provider

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"golang.org/x/net/context"

	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"

	"google.golang.org/api/dns/v1"
	"google.golang.org/api/googleapi"
)

var (
	testZones   = map[string]*dns.ManagedZone{}
	testRecords = map[string]map[string]*dns.ResourceRecordSet{}
)

type mockManagedZonesCreateCall struct {
	project     string
	managedZone *dns.ManagedZone
}

func (m *mockManagedZonesCreateCall) Do(opts ...googleapi.CallOption) (*dns.ManagedZone, error) {
	zoneKey := zoneKey(m.project, m.managedZone.Name)

	if _, ok := testZones[zoneKey]; ok {
		return nil, &googleapi.Error{Code: http.StatusConflict}
	}

	testZones[zoneKey] = m.managedZone

	return m.managedZone, nil
}

type mockManagedZonesListCall struct {
	project string
}

func (m *mockManagedZonesListCall) Pages(ctx context.Context, f func(*dns.ManagedZonesListResponse) error) error {
	zones := []*dns.ManagedZone{}

	for k, v := range testZones {
		if strings.HasPrefix(k, m.project+"/") {
			zones = append(zones, v)
		}
	}

	return f(&dns.ManagedZonesListResponse{ManagedZones: zones})
}

type mockManagedZonesClient struct{}

func (m *mockManagedZonesClient) Create(project string, managedZone *dns.ManagedZone) managedZonesCreateCallInterface {
	return &mockManagedZonesCreateCall{project: project, managedZone: managedZone}
}

func (m *mockManagedZonesClient) List(project string) managedZonesListCallInterface {
	return &mockManagedZonesListCall{project: project}
}

type mockResourceRecordSetsListCall struct {
	project     string
	managedZone string
}

func (m *mockResourceRecordSetsListCall) Pages(ctx context.Context, f func(*dns.ResourceRecordSetsListResponse) error) error {
	zoneKey := zoneKey(m.project, m.managedZone)

	if _, ok := testZones[zoneKey]; !ok {
		return &googleapi.Error{Code: http.StatusNotFound}
	}

	resp := []*dns.ResourceRecordSet{}

	for _, v := range testRecords[zoneKey] {
		resp = append(resp, v)
	}

	return f(&dns.ResourceRecordSetsListResponse{Rrsets: resp})
}

type mockResourceRecordSetsClient struct{}

func (m *mockResourceRecordSetsClient) List(project string, managedZone string) resourceRecordSetsListCallInterface {
	return &mockResourceRecordSetsListCall{project: project, managedZone: managedZone}
}

type mockChangesCreateCall struct {
	project     string
	managedZone string
	change      *dns.Change
}

func (m *mockChangesCreateCall) Do(opts ...googleapi.CallOption) (*dns.Change, error) {
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

type mockChangesClient struct{}

func (m *mockChangesClient) Create(project string, managedZone string, change *dns.Change) changesCreateCallInterface {
	return &mockChangesCreateCall{project: project, managedZone: managedZone, change: change}
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
	case "CNAME":
		for _, rrd := range recordSet.Rrdatas {
			if !hasTrailingDot(rrd) {
				return false
			}
		}
	case "A", "TXT":
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

func hasTrailingDot(target string) bool {
	return strings.HasSuffix(target, ".")
}

func TestGoogleZones(t *testing.T) {
	provider := newGoogleProvider(t, "ext-dns-test-2.gcp.zalan.do.", false, []*endpoint.Endpoint{})

	zones, err := provider.Zones()
	if err != nil {
		t.Fatal(err)
	}

	validateZones(t, zones, map[string]*dns.ManagedZone{
		"zone-1-ext-dns-test-2-gcp-zalan-do": {Name: "zone-1-ext-dns-test-2-gcp-zalan-do", DnsName: "zone-1.ext-dns-test-2.gcp.zalan.do."},
		"zone-2-ext-dns-test-2-gcp-zalan-do": {Name: "zone-2-ext-dns-test-2-gcp-zalan-do", DnsName: "zone-2.ext-dns-test-2.gcp.zalan.do."},
		"zone-3-ext-dns-test-2-gcp-zalan-do": {Name: "zone-3-ext-dns-test-2-gcp-zalan-do", DnsName: "zone-3.ext-dns-test-2.gcp.zalan.do."},
	})
}

func TestGoogleRecords(t *testing.T) {
	originalEndpoints := []*endpoint.Endpoint{
		endpoint.NewEndpoint("list-test.zone-1.ext-dns-test-2.gcp.zalan.do", "1.2.3.4", "A"),
		endpoint.NewEndpoint("list-test.zone-2.ext-dns-test-2.gcp.zalan.do", "8.8.8.8", "A"),
		endpoint.NewEndpoint("list-test-alias.zone-1.ext-dns-test-2.gcp.zalan.do", "foo.elb.amazonaws.com", "CNAME"),
	}

	provider := newGoogleProvider(t, "ext-dns-test-2.gcp.zalan.do.", false, originalEndpoints)

	records, err := provider.Records()
	if err != nil {
		t.Fatal(err)
	}

	validateEndpoints(t, records, originalEndpoints)
}

func TestGoogleCreateRecords(t *testing.T) {
	provider := newGoogleProvider(t, "ext-dns-test-2.gcp.zalan.do.", false, []*endpoint.Endpoint{})

	records := []*endpoint.Endpoint{
		endpoint.NewEndpoint("create-test.zone-1.ext-dns-test-2.gcp.zalan.do", "1.2.3.4", ""),
		endpoint.NewEndpoint("create-test.zone-2.ext-dns-test-2.gcp.zalan.do", "8.8.8.8", ""),
		endpoint.NewEndpoint("create-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", "foo.elb.amazonaws.com", ""),
	}

	if err := provider.CreateRecords(records); err != nil {
		t.Fatal(err)
	}

	records, err := provider.Records()
	if err != nil {
		t.Fatal(err)
	}

	validateEndpoints(t, records, []*endpoint.Endpoint{
		endpoint.NewEndpoint("create-test.zone-1.ext-dns-test-2.gcp.zalan.do", "1.2.3.4", "A"),
		endpoint.NewEndpoint("create-test.zone-2.ext-dns-test-2.gcp.zalan.do", "8.8.8.8", "A"),
		endpoint.NewEndpoint("create-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", "foo.elb.amazonaws.com", "CNAME"),
	})
}

func TestGoogleUpdateRecords(t *testing.T) {
	provider := newGoogleProvider(t, "ext-dns-test-2.gcp.zalan.do.", false, []*endpoint.Endpoint{
		endpoint.NewEndpoint("update-test.zone-1.ext-dns-test-2.gcp.zalan.do", "8.8.8.8", "A"),
		endpoint.NewEndpoint("update-test.zone-2.ext-dns-test-2.gcp.zalan.do", "8.8.4.4", "A"),
		endpoint.NewEndpoint("update-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", "foo.elb.amazonaws.com", "CNAME"),
	})

	currentRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("update-test.zone-1.ext-dns-test-2.gcp.zalan.do", "8.8.8.8", "A"),
		endpoint.NewEndpoint("update-test.zone-2.ext-dns-test-2.gcp.zalan.do", "8.8.4.4", "A"),
		endpoint.NewEndpoint("update-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", "foo.elb.amazonaws.com", "CNAME"),
	}
	updatedRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("update-test.zone-1.ext-dns-test-2.gcp.zalan.do", "1.2.3.4", "A"),
		endpoint.NewEndpoint("update-test.zone-2.ext-dns-test-2.gcp.zalan.do", "4.3.2.1", "A"),
		endpoint.NewEndpoint("update-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", "bar.elb.amazonaws.com", "CNAME"),
	}

	if err := provider.UpdateRecords(updatedRecords, currentRecords); err != nil {
		t.Fatal(err)
	}

	records, err := provider.Records()
	if err != nil {
		t.Fatal(err)
	}

	validateEndpoints(t, records, []*endpoint.Endpoint{
		endpoint.NewEndpoint("update-test.zone-1.ext-dns-test-2.gcp.zalan.do", "1.2.3.4", "A"),
		endpoint.NewEndpoint("update-test.zone-2.ext-dns-test-2.gcp.zalan.do", "4.3.2.1", "A"),
		endpoint.NewEndpoint("update-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", "bar.elb.amazonaws.com", "CNAME"),
	})
}

func TestGoogleDeleteRecords(t *testing.T) {
	originalEndpoints := []*endpoint.Endpoint{
		endpoint.NewEndpoint("delete-test.zone-1.ext-dns-test-2.gcp.zalan.do", "1.2.3.4", "A"),
		endpoint.NewEndpoint("delete-test.zone-2.ext-dns-test-2.gcp.zalan.do", "8.8.8.8", "A"),
		endpoint.NewEndpoint("delete-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", "baz.elb.amazonaws.com", "CNAME"),
	}

	provider := newGoogleProvider(t, "ext-dns-test-2.gcp.zalan.do.", false, originalEndpoints)

	if err := provider.DeleteRecords(originalEndpoints); err != nil {
		t.Fatal(err)
	}

	records, err := provider.Records()
	if err != nil {
		t.Fatal(err)
	}

	validateEndpoints(t, records, []*endpoint.Endpoint{})
}

func TestGoogleApplyChanges(t *testing.T) {
	provider := newGoogleProvider(t, "ext-dns-test-2.gcp.zalan.do.", false, []*endpoint.Endpoint{
		endpoint.NewEndpoint("update-test.zone-1.ext-dns-test-2.gcp.zalan.do", "8.8.8.8", "A"),
		endpoint.NewEndpoint("delete-test.zone-1.ext-dns-test-2.gcp.zalan.do", "8.8.8.8", "A"),
		endpoint.NewEndpoint("update-test.zone-2.ext-dns-test-2.gcp.zalan.do", "8.8.4.4", "A"),
		endpoint.NewEndpoint("delete-test.zone-2.ext-dns-test-2.gcp.zalan.do", "8.8.4.4", "A"),
		endpoint.NewEndpoint("update-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", "bar.elb.amazonaws.com", "CNAME"),
		endpoint.NewEndpoint("delete-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", "qux.elb.amazonaws.com", "CNAME"),
	})

	createRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("create-test.zone-1.ext-dns-test-2.gcp.zalan.do", "8.8.8.8", "A"),
		endpoint.NewEndpoint("create-test.zone-2.ext-dns-test-2.gcp.zalan.do", "8.8.4.4", "A"),
		endpoint.NewEndpoint("create-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", "foo.elb.amazonaws.com", "CNAME"),
	}

	currentRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("update-test.zone-1.ext-dns-test-2.gcp.zalan.do", "8.8.8.8", "A"),
		endpoint.NewEndpoint("update-test.zone-2.ext-dns-test-2.gcp.zalan.do", "8.8.4.4", "A"),
		endpoint.NewEndpoint("update-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", "bar.elb.amazonaws.com", "CNAME"),
	}
	updatedRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("update-test.zone-1.ext-dns-test-2.gcp.zalan.do", "1.2.3.4", "A"),
		endpoint.NewEndpoint("update-test.zone-2.ext-dns-test-2.gcp.zalan.do", "4.3.2.1", "A"),
		endpoint.NewEndpoint("update-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", "baz.elb.amazonaws.com", "CNAME"),
	}

	deleteRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("delete-test.zone-1.ext-dns-test-2.gcp.zalan.do", "8.8.8.8", "A"),
		endpoint.NewEndpoint("delete-test.zone-2.ext-dns-test-2.gcp.zalan.do", "8.8.4.4", "A"),
		endpoint.NewEndpoint("delete-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", "qux.elb.amazonaws.com", "CNAME"),
	}

	changes := &plan.Changes{
		Create:    createRecords,
		UpdateNew: updatedRecords,
		UpdateOld: currentRecords,
		Delete:    deleteRecords,
	}

	if err := provider.ApplyChanges(changes); err != nil {
		t.Fatal(err)
	}

	records, err := provider.Records()
	if err != nil {
		t.Fatal(err)
	}

	validateEndpoints(t, records, []*endpoint.Endpoint{
		endpoint.NewEndpoint("create-test.zone-1.ext-dns-test-2.gcp.zalan.do", "8.8.8.8", "A"),
		endpoint.NewEndpoint("update-test.zone-1.ext-dns-test-2.gcp.zalan.do", "1.2.3.4", "A"),
		endpoint.NewEndpoint("create-test.zone-2.ext-dns-test-2.gcp.zalan.do", "8.8.4.4", "A"),
		endpoint.NewEndpoint("update-test.zone-2.ext-dns-test-2.gcp.zalan.do", "4.3.2.1", "A"),
		endpoint.NewEndpoint("create-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", "foo.elb.amazonaws.com", "CNAME"),
		endpoint.NewEndpoint("update-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", "baz.elb.amazonaws.com", "CNAME"),
	})
}

func TestGoogleApplyChangesDryRun(t *testing.T) {
	originalEndpoints := []*endpoint.Endpoint{
		endpoint.NewEndpoint("update-test.zone-1.ext-dns-test-2.gcp.zalan.do", "8.8.8.8", "A"),
		endpoint.NewEndpoint("delete-test.zone-1.ext-dns-test-2.gcp.zalan.do", "8.8.8.8", "A"),
		endpoint.NewEndpoint("update-test.zone-2.ext-dns-test-2.gcp.zalan.do", "8.8.4.4", "A"),
		endpoint.NewEndpoint("delete-test.zone-2.ext-dns-test-2.gcp.zalan.do", "8.8.4.4", "A"),
		endpoint.NewEndpoint("update-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", "bar.elb.amazonaws.com", "CNAME"),
		endpoint.NewEndpoint("delete-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", "qux.elb.amazonaws.com", "CNAME"),
	}

	provider := newGoogleProvider(t, "ext-dns-test-2.gcp.zalan.do.", true, originalEndpoints)

	createRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("create-test.zone-1.ext-dns-test-2.gcp.zalan.do", "8.8.8.8", "A"),
		endpoint.NewEndpoint("create-test.zone-2.ext-dns-test-2.gcp.zalan.do", "8.8.4.4", "A"),
		endpoint.NewEndpoint("create-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", "foo.elb.amazonaws.com", "CNAME"),
	}

	currentRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("update-test.zone-1.ext-dns-test-2.gcp.zalan.do", "8.8.8.8", "A"),
		endpoint.NewEndpoint("update-test.zone-2.ext-dns-test-2.gcp.zalan.do", "8.8.4.4", "A"),
		endpoint.NewEndpoint("update-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", "bar.elb.amazonaws.com", "CNAME"),
	}
	updatedRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("update-test.zone-1.ext-dns-test-2.gcp.zalan.do", "1.2.3.4", "A"),
		endpoint.NewEndpoint("update-test.zone-2.ext-dns-test-2.gcp.zalan.do", "4.3.2.1", "A"),
		endpoint.NewEndpoint("update-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", "baz.elb.amazonaws.com", "CNAME"),
	}

	deleteRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("delete-test.zone-1.ext-dns-test-2.gcp.zalan.do", "8.8.8.8", "A"),
		endpoint.NewEndpoint("delete-test.zone-2.ext-dns-test-2.gcp.zalan.do", "8.8.4.4", "A"),
		endpoint.NewEndpoint("delete-test-cname.zone-1.ext-dns-test-2.gcp.zalan.do", "qux.elb.amazonaws.com", "CNAME"),
	}

	changes := &plan.Changes{
		Create:    createRecords,
		UpdateNew: updatedRecords,
		UpdateOld: currentRecords,
		Delete:    deleteRecords,
	}

	if err := provider.ApplyChanges(changes); err != nil {
		t.Fatal(err)
	}

	records, err := provider.Records()
	if err != nil {
		t.Fatal(err)
	}

	validateEndpoints(t, records, originalEndpoints)
}

func TestGoogleApplyChangesEmpty(t *testing.T) {
	provider := newGoogleProvider(t, "ext-dns-test-2.gcp.zalan.do.", false, []*endpoint.Endpoint{})

	if err := provider.ApplyChanges(&plan.Changes{}); err != nil {
		t.Error(err)
	}
}

func TestSeparateChanges(t *testing.T) {
	change := &dns.Change{
		Additions: []*dns.ResourceRecordSet{
			{Name: "qux.foo.example.org.", Ttl: 1},
			{Name: "qux.bar.example.org.", Ttl: 2},
		},
		Deletions: []*dns.ResourceRecordSet{
			{Name: "wambo.foo.example.org.", Ttl: 10},
			{Name: "wambo.bar.example.org.", Ttl: 20},
		},
	}

	zones := map[string]*dns.ManagedZone{
		"foo-example-org": {
			Name:    "foo-example-org",
			DnsName: "foo.example.org.",
		},
		"bar-example-org": {
			Name:    "bar-example-org",
			DnsName: "bar.example.org.",
		},
		"baz-example-org": {
			Name:    "baz-example-org",
			DnsName: "baz.example.org.",
		},
	}

	changes := separateChange(zones, change)

	if len(changes) != 2 {
		t.Fatalf("expected %d change(s), got %d", 2, len(changes))
	}

	validateChange(t, changes["foo-example-org"], &dns.Change{
		Additions: []*dns.ResourceRecordSet{
			{Name: "qux.foo.example.org.", Ttl: 1},
		},
		Deletions: []*dns.ResourceRecordSet{
			{Name: "wambo.foo.example.org.", Ttl: 10},
		},
	})

	validateChange(t, changes["bar-example-org"], &dns.Change{
		Additions: []*dns.ResourceRecordSet{
			{Name: "qux.bar.example.org.", Ttl: 2},
		},
		Deletions: []*dns.ResourceRecordSet{
			{Name: "wambo.bar.example.org.", Ttl: 20},
		},
	})
}

func TestGoogleSuitableZone(t *testing.T) {
	zones := map[string]*dns.ManagedZone{
		"example-org":     {Name: "example-org", DnsName: "example.org."},
		"bar-example-org": {Name: "bar-example-org", DnsName: "bar.example.org."},
	}

	for _, tc := range []struct {
		hostname string
		expected *dns.ManagedZone
	}{
		{"foo.bar.example.org.", zones["bar-example-org"]},
		{"foo.example.org.", zones["example-org"]},
		{"foo.kubernetes.io.", nil},
	} {
		suitableZone := suitableManagedZone(tc.hostname, zones)

		if suitableZone != tc.expected {
			t.Errorf("expected %v, got %v", tc.expected, suitableZone)
		}
	}
}

func validateZones(t *testing.T, zones map[string]*dns.ManagedZone, expected map[string]*dns.ManagedZone) {
	if len(zones) != len(expected) {
		t.Fatalf("expected %d zone(s), got %d", len(expected), len(zones))
	}

	for i, zone := range zones {
		validateZone(t, zone, expected[i])
	}
}

func validateZone(t *testing.T, zone *dns.ManagedZone, expected *dns.ManagedZone) {
	if zone.Name != expected.Name {
		t.Errorf("expected %s, got %s", expected.Name, zone.Name)
	}

	if zone.DnsName != expected.DnsName {
		t.Errorf("expected %s, got %s", expected.DnsName, zone.DnsName)
	}
}

func validateChange(t *testing.T, change *dns.Change, expected *dns.Change) {
	validateChangeRecords(t, change.Additions, expected.Additions)
	validateChangeRecords(t, change.Deletions, expected.Deletions)
}

func validateChangeRecords(t *testing.T, records []*dns.ResourceRecordSet, expected []*dns.ResourceRecordSet) {
	if len(records) != len(expected) {
		t.Fatalf("expected %d change(s), got %d", len(expected), len(records))
	}

	for i := range records {
		validateChangeRecord(t, records[i], expected[i])
	}
}

func validateChangeRecord(t *testing.T, record *dns.ResourceRecordSet, expected *dns.ResourceRecordSet) {
	if record.Name != expected.Name {
		t.Errorf("expected %s, got %s", expected.Name, record.Name)
	}

	if record.Ttl != expected.Ttl {
		t.Errorf("expected %d, got %d", expected.Ttl, record.Ttl)
	}
}

func newGoogleProvider(t *testing.T, domainFilter string, dryRun bool, records []*endpoint.Endpoint) *googleProvider {
	provider := &googleProvider{
		project:      "zalando-external-dns-test",
		domainFilter: domainFilter,
		dryRun:       false,
		resourceRecordSetsClient: &mockResourceRecordSetsClient{},
		managedZonesClient:       &mockManagedZonesClient{},
		changesClient:            &mockChangesClient{},
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

	setupGoogleRecords(t, provider, records)

	provider.dryRun = dryRun

	return provider
}

func createZone(t *testing.T, provider *googleProvider, zone *dns.ManagedZone) {
	zone.Description = "Testing zone for kubernetes.io/external-dns"

	if _, err := provider.managedZonesClient.Create("zalando-external-dns-test", zone).Do(); err != nil {
		if err, ok := err.(*googleapi.Error); !ok || err.Code != http.StatusConflict {
			t.Fatal(err)
		}
	}
}

func setupGoogleRecords(t *testing.T, provider *googleProvider, endpoints []*endpoint.Endpoint) {
	clearGoogleRecords(t, provider, "zone-1-ext-dns-test-2-gcp-zalan-do")
	clearGoogleRecords(t, provider, "zone-2-ext-dns-test-2-gcp-zalan-do")

	records, err := provider.Records()
	if err != nil {
		t.Fatal(err)
	}

	validateEndpoints(t, records, []*endpoint.Endpoint{})

	if err = provider.CreateRecords(endpoints); err != nil {
		t.Fatal(err)
	}

	records, err = provider.Records()
	if err != nil {
		t.Fatal(err)
	}

	validateEndpoints(t, records, endpoints)
}

func clearGoogleRecords(t *testing.T, provider *googleProvider, zone string) {
	recordSets := []*dns.ResourceRecordSet{}
	if err := provider.resourceRecordSetsClient.List(provider.project, zone).Pages(context.TODO(), func(resp *dns.ResourceRecordSetsListResponse) error {
		for _, r := range resp.Rrsets {
			switch r.Type {
			case "A", "CNAME":
				recordSets = append(recordSets, r)
			}
		}
		return nil
	}); err != nil {
		t.Fatal(err)
	}

	if len(recordSets) != 0 {
		if _, err := provider.changesClient.Create(provider.project, zone, &dns.Change{
			Deletions: recordSets,
		}).Do(); err != nil {
			t.Fatal(err)
		}
	}
}
