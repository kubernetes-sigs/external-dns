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

package akamai

import (
	"context"
	"encoding/json"
	"sync"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"

	dns "github.com/akamai/AkamaiOPEN-edgegrid-golang/v13/pkg/dns"
	"github.com/stretchr/testify/assert"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

type edgednsStubData struct {
	objType string // zone, record, recordsets
	output  []any
}

// edgednsStub records how many calls each Edge DNS operation received. The provider is
// dominated by these round trips, so the call count is what the performance tests assert.
type edgednsStub struct {
	stubData map[string]edgednsStubData

	// Records lists the zones concurrently, so the counters below are written from
	// several goroutines at once.
	mu                sync.Mutex
	listZonesCalls    int
	getRecordSetCalls int
	createCalls       int
	updateCalls       int
	deleteCalls       int
}

func newStub() *edgednsStub {
	return &edgednsStub{
		stubData: make(map[string]edgednsStubData),
	}
}

func createAkamaiStubProvider(stub *edgednsStub, domfilter *endpoint.DomainFilter, idfilter provider.ZoneIDFilter) (*AkamaiProvider, error) {
	return createAkamaiStubProviderWithCache(stub, domfilter, idfilter, 0)
}

func createAkamaiStubProviderWithCache(stub *edgednsStub, domfilter *endpoint.DomainFilter, idfilter provider.ZoneIDFilter, cacheDuration time.Duration) (*AkamaiProvider, error) {
	akamaiConfig := AkamaiConfig{
		DomainFilter:          domfilter,
		ZoneIDFilter:          idfilter,
		ServiceConsumerDomain: "testzone.com",
		ClientToken:           "test_token",
		ClientSecret:          "test_client_secret",
		AccessToken:           "test_access_token",
		ZoneCacheDuration:     cacheDuration,
	}

	prov, err := newProvider(akamaiConfig, stub)
	aprov := prov.(*AkamaiProvider)
	return aprov, err
}

func (r *edgednsStub) createStubDataEntry(objtype string) {
	log.Debugf("Creating stub data entry")
	if _, exists := r.stubData[objtype]; !exists {
		r.stubData[objtype] = edgednsStubData{objType: objtype}
	}
}

func (r *edgednsStub) setOutput(objtype string, output []any) {
	log.Debugf("Setting output to %v", output)
	r.createStubDataEntry(objtype)
	stubdata := r.stubData[objtype]
	stubdata.output = output
	r.stubData[objtype] = stubdata
}

func (r *edgednsStub) ListZones(_ context.Context, _ dns.ListZonesRequest) (*dns.ZoneListResponse, error) {
	log.Debugf("Entering ListZones")
	r.mu.Lock()
	r.listZonesCalls++
	r.mu.Unlock()
	// Ignore Metadata
	resp := &dns.ZoneListResponse{}
	zones := make([]dns.ZoneResponse, 0)
	for _, zname := range r.stubData["zone"].output {
		log.Debugf("Processing output: %v", zname)
		zn := dns.ZoneResponse{Zone: zname.(string), ContractID: "contract"}
		log.Debugf("Created Zone Object: %v", zn)
		zones = append(zones, zn)
	}
	resp.Zones = zones
	return resp, nil
}

func (r *edgednsStub) GetRecordSets(_ context.Context, _ dns.GetRecordSetsRequest) (*dns.GetRecordSetsResponse, error) {
	log.Debugf("Entering GetRecordSets")
	r.mu.Lock()
	r.getRecordSetCalls++
	r.mu.Unlock()
	// Ignore Metadata
	resp := &dns.GetRecordSetsResponse{}
	sets := make([]dns.RecordSet, 0)
	for _, rec := range r.stubData["recordset"].output {
		rset := rec.(dns.RecordSet)
		sets = append(sets, rset)
	}
	resp.RecordSets = sets

	return resp, nil
}

func (r *edgednsStub) CreateRecordSets(_ context.Context, _ dns.CreateRecordSetsRequest) error {
	r.mu.Lock()
	r.createCalls++
	r.mu.Unlock()
	return nil
}

func (r *edgednsStub) UpdateRecord(_ context.Context, _ dns.UpdateRecordRequest) error {
	r.mu.Lock()
	r.updateCalls++
	r.mu.Unlock()
	return nil
}

func (r *edgednsStub) DeleteRecord(_ context.Context, _ dns.DeleteRecordRequest) error {
	r.mu.Lock()
	r.deleteCalls++
	r.mu.Unlock()
	return nil
}

// Test FetchZones
func TestFetchZonesZoneIDFilter(t *testing.T) {
	stub := newStub()
	domfilter := &endpoint.DomainFilter{}
	idfilter := provider.NewZoneIDFilter([]string{"Test"})
	c, err := createAkamaiStubProvider(stub, domfilter, idfilter)
	assert.NoError(t, err)
	stub.setOutput("zone", []any{"test1.testzone.com", "test2.testzone.com"})

	x, _ := c.fetchZones(t.Context())
	y, err := json.Marshal(x)
	require.NoError(t, err)
	if assert.NotNil(t, y) {
		assert.JSONEq(t, "{\"zones\":[{\"contractId\":\"contract\",\"zone\":\"test1.testzone.com\"},{\"contractId\":\"contract\",\"zone\":\"test2.testzone.com\"}]}", string(y))
	}
}

func TestFetchZonesEmpty(t *testing.T) {
	stub := newStub()
	domfilter := endpoint.NewDomainFilter([]string{"Nonexistent"})
	idfilter := provider.NewZoneIDFilter([]string{"Nonexistent"})
	c, err := createAkamaiStubProvider(stub, domfilter, idfilter)
	require.NoError(t, err)
	stub.setOutput("zone", []any{})

	x, _ := c.fetchZones(t.Context())
	y, err := json.Marshal(x)
	require.NoError(t, err)
	if assert.NotNil(t, y) {
		assert.JSONEq(t, "{\"zones\":[]}", string(y))
	}
}

// TestAkamaiRecords tests record endpoint
func TestAkamaiRecords(t *testing.T) {
	stub := newStub()
	domfilter := &endpoint.DomainFilter{}
	idfilter := provider.ZoneIDFilter{}
	c, err := createAkamaiStubProvider(stub, domfilter, idfilter)
	require.NoError(t, err)
	stub.setOutput("zone", []any{"test1.testzone.com"})
	recordsets := make([]any, 0)
	recordsets = append(recordsets, dns.RecordSet{
		Name:  "www.example.com",
		Type:  endpoint.RecordTypeA,
		Rdata: []string{"10.0.0.2", "10.0.0.3"},
	})
	recordsets = append(recordsets, dns.RecordSet{
		Name:  "www.example.com",
		Type:  endpoint.RecordTypeTXT,
		Rdata: []string{"heritage=external-dns,external-dns/owner=default"},
	})
	recordsets = append(recordsets, dns.RecordSet{
		Name:  "www.exclude.me",
		Type:  endpoint.RecordTypeA,
		Rdata: []string{"192.168.0.1", "192.168.0.2"},
	})
	stub.setOutput("recordset", recordsets)
	endpoints := make([]*endpoint.Endpoint, 0)
	endpoints = append(endpoints, endpoint.NewEndpoint("www.example.com", endpoint.RecordTypeA, "10.0.0.2", "10.0.0.3"))
	endpoints = append(endpoints, endpoint.NewEndpoint("www.example.com", endpoint.RecordTypeTXT, "heritage=external-dns,external-dns/owner=default"))
	endpoints = append(endpoints, endpoint.NewEndpoint("www.exclude.me", endpoint.RecordTypeA, "192.168.0.1", "192.168.0.2"))

	x, _ := c.Records(t.Context())
	if assert.NotNil(t, x) {
		assert.Equal(t, endpoints, x)
	}
}

func TestAkamaiRecordsEmpty(t *testing.T) {
	stub := newStub()
	domfilter := &endpoint.DomainFilter{}
	idfilter := provider.NewZoneIDFilter([]string{"Nonexistent"})
	c, err := createAkamaiStubProvider(stub, domfilter, idfilter)
	require.NoError(t, err)
	stub.setOutput("zone", []any{"test1.testzone.com"})
	recordsets := make([]any, 0)
	stub.setOutput("recordset", recordsets)

	x, _ := c.Records(t.Context())
	assert.Nil(t, x)
}

func TestAkamaiRecordsFilters(t *testing.T) {
	stub := newStub()
	domfilter := endpoint.NewDomainFilter([]string{"www.exclude.me"})
	idfilter := provider.ZoneIDFilter{}
	c, err := createAkamaiStubProvider(stub, domfilter, idfilter)
	assert.NoError(t, err)
	stub.setOutput("zone", []any{"www.exclude.me"})
	recordsets := make([]any, 0)
	recordsets = append(recordsets, dns.RecordSet{
		Name:  "www.example.com",
		Type:  endpoint.RecordTypeA,
		Rdata: []string{"10.0.0.2", "10.0.0.3"},
	})
	recordsets = append(recordsets, dns.RecordSet{
		Name:  "www.exclude.me",
		Type:  endpoint.RecordTypeA,
		Rdata: []string{"192.168.0.1", "192.168.0.2"},
	})
	stub.setOutput("recordset", recordsets)
	endpoints := make([]*endpoint.Endpoint, 0)
	endpoints = append(endpoints, endpoint.NewEndpoint("www.exclude.me", endpoint.RecordTypeA, "192.168.0.1", "192.168.0.2"))

	x, _ := c.Records(t.Context())
	if assert.NotNil(t, x) {
		assert.Equal(t, endpoints, x)
	}
}

// TestRecordsListsEachZoneOnce guards the concurrent zone listing: one listing per zone,
// no more, and the endpoints stay ordered by zone regardless of completion order.
func TestRecordsListsEachZoneOnce(t *testing.T) {
	stub := newStub()
	c, err := createAkamaiStubProvider(stub, &endpoint.DomainFilter{}, provider.ZoneIDFilter{})
	require.NoError(t, err)

	stub.setOutput("zone", []any{"a.testzone.com", "b.testzone.com", "c.testzone.com", "d.testzone.com", "e.testzone.com"})
	stub.setOutput("recordset", []any{dns.RecordSet{
		Name:  "www.example.com",
		Type:  endpoint.RecordTypeA,
		Rdata: []string{"10.0.0.2"},
	}})

	_, err = c.Records(t.Context())
	require.NoError(t, err)
	assert.Equal(t, 1, stub.listZonesCalls)
	assert.Equal(t, 5, stub.getRecordSetCalls)
}

// TestCreateRecords tests create function
func TestCreateRecords(t *testing.T) {
	stub := newStub()
	domfilter := &endpoint.DomainFilter{}
	idfilter := provider.ZoneIDFilter{}
	c, err := createAkamaiStubProvider(stub, domfilter, idfilter)
	assert.NoError(t, err)

	zoneNameIDMapper := provider.ZoneIDName{"example.com": "example.com"}
	endpoints := make([]*endpoint.Endpoint, 0)
	endpoints = append(endpoints, endpoint.NewEndpoint("www.example.com", endpoint.RecordTypeA, "10.0.0.2", "10.0.0.3"))
	endpoints = append(endpoints, endpoint.NewEndpoint("www.example.com", endpoint.RecordTypeTXT, "heritage=external-dns,external-dns/owner=default"))

	err = c.createRecordsets(t.Context(), zoneNameIDMapper, endpoints)
	assert.NoError(t, err)
}

func TestCreateRecordsDomainFilter(t *testing.T) {
	stub := newStub()
	domfilter := &endpoint.DomainFilter{}
	idfilter := provider.ZoneIDFilter{}
	c, err := createAkamaiStubProvider(stub, domfilter, idfilter)
	assert.NoError(t, err)

	zoneNameIDMapper := provider.ZoneIDName{"example.com": "example.com"}
	exclude := []*endpoint.Endpoint{
		endpoint.NewEndpoint("www.example.com", endpoint.RecordTypeA, "10.0.0.2", "10.0.0.3"),
		endpoint.NewEndpoint("www.example.com", endpoint.RecordTypeTXT, "heritage=external-dns,external-dns/owner=default"),
		endpoint.NewEndpoint("www.exclude.me", endpoint.RecordTypeA, "10.0.0.2", "10.0.0.3"),
	}

	err = c.createRecordsets(t.Context(), zoneNameIDMapper, exclude)
	assert.NoError(t, err)
}

// TestDeleteRecords validate delete
func TestDeleteRecords(t *testing.T) {
	stub := newStub()
	domfilter := &endpoint.DomainFilter{}
	idfilter := provider.ZoneIDFilter{}
	c, err := createAkamaiStubProvider(stub, domfilter, idfilter)
	assert.NoError(t, err)

	zoneNameIDMapper := provider.ZoneIDName{"example.com": "example.com"}
	endpoints := make([]*endpoint.Endpoint, 0)
	endpoints = append(endpoints, endpoint.NewEndpoint("www.example.com", endpoint.RecordTypeA, "10.0.0.2", "10.0.0.3"))
	endpoints = append(endpoints, endpoint.NewEndpoint("www.example.com", endpoint.RecordTypeTXT, "heritage=external-dns,external-dns/owner=default"))

	err = c.deleteRecordsets(t.Context(), zoneNameIDMapper, endpoints)
	assert.NoError(t, err)
	// One DELETE per endpoint and nothing else: the record is not read back first.
	assert.Equal(t, 2, stub.deleteCalls)
	assert.Equal(t, 0, stub.getRecordSetCalls)
}

func TestDeleteRecordsDomainFilter(t *testing.T) {
	stub := newStub()
	domfilter := endpoint.NewDomainFilter([]string{"example.com"})
	idfilter := provider.ZoneIDFilter{}
	c, err := createAkamaiStubProvider(stub, domfilter, idfilter)
	require.NoError(t, err)

	zoneNameIDMapper := provider.ZoneIDName{"example.com": "example.com"}
	exclude := []*endpoint.Endpoint{
		endpoint.NewEndpoint("www.example.com", endpoint.RecordTypeA, "10.0.0.2", "10.0.0.3"),
		endpoint.NewEndpoint("www.example.com", endpoint.RecordTypeTXT, "heritage=external-dns,external-dns/owner=default"),
		endpoint.NewEndpoint("www.exclude.me", endpoint.RecordTypeA, "10.0.0.2", "10.0.0.3"),
	}

	err = c.deleteRecordsets(t.Context(), zoneNameIDMapper, exclude)
	assert.NoError(t, err)
}

// Test record update func
func TestUpdateRecords(t *testing.T) {
	stub := newStub()
	domfilter := &endpoint.DomainFilter{}
	idfilter := provider.ZoneIDFilter{}
	c, err := createAkamaiStubProvider(stub, domfilter, idfilter)
	require.NoError(t, err)

	zoneNameIDMapper := provider.ZoneIDName{"example.com": "example.com"}
	endpoints := make([]*endpoint.Endpoint, 0)
	endpoints = append(endpoints, endpoint.NewEndpoint("www.example.com", endpoint.RecordTypeA, "10.0.0.2", "10.0.0.3"))
	endpoints = append(endpoints, endpoint.NewEndpoint("www.example.com", endpoint.RecordTypeTXT, "heritage=external-dns,external-dns/owner=default"))

	err = c.updateNewRecordsets(t.Context(), zoneNameIDMapper, endpoints)
	require.NoError(t, err)
	// One PUT per endpoint and nothing else: the record is not read back first.
	assert.Equal(t, 2, stub.updateCalls)
	assert.Equal(t, 0, stub.getRecordSetCalls)
}

func TestUpdateRecordsDomainFilter(t *testing.T) {
	stub := newStub()
	domfilter := endpoint.NewDomainFilter([]string{"example.com"})
	idfilter := provider.ZoneIDFilter{}
	c, err := createAkamaiStubProvider(stub, domfilter, idfilter)
	require.NoError(t, err)

	zoneNameIDMapper := provider.ZoneIDName{"example.com": "example.com"}
	exclude := []*endpoint.Endpoint{
		endpoint.NewEndpoint("www.example.com", endpoint.RecordTypeA, "10.0.0.2", "10.0.0.3"),
		endpoint.NewEndpoint("www.example.com", endpoint.RecordTypeTXT, "heritage=external-dns,external-dns/owner=default"),
		endpoint.NewEndpoint("www.exclude.me", endpoint.RecordTypeA, "10.0.0.2", "10.0.0.3"),
	}

	err = c.updateNewRecordsets(t.Context(), zoneNameIDMapper, exclude)
	require.NoError(t, err)
}

func TestAkamaiApplyChanges(t *testing.T) {
	stub := newStub()
	domfilter := endpoint.NewDomainFilter([]string{"example.com"})
	idfilter := provider.ZoneIDFilter{}
	c, err := createAkamaiStubProvider(stub, domfilter, idfilter)
	assert.NoError(t, err)

	stub.setOutput("zone", []any{"example.com"})
	changes := &plan.Changes{}
	changes.Create = []*endpoint.Endpoint{
		{DNSName: "www.example.com", RecordType: "A", Targets: endpoint.Targets{"target"}, RecordTTL: 300},
		{DNSName: "test.example.com", RecordType: "A", Targets: endpoint.Targets{"target"}, RecordTTL: 300},
		{DNSName: "test.this.example.com", RecordType: "A", Targets: endpoint.Targets{"127.0.0.1"}, RecordTTL: 300},
		{DNSName: "www.example.com", RecordType: "TXT", Targets: endpoint.Targets{"heritage=external-dns,external-dns/owner=default"}, RecordTTL: 300},
		{DNSName: "test.example.com", RecordType: "TXT", Targets: endpoint.Targets{"heritage=external-dns,external-dns/owner=default"}, RecordTTL: 300},
		{DNSName: "test.this.example.com", RecordType: "TXT", Targets: endpoint.Targets{"heritage=external-dns,external-dns/owner=default"}, RecordTTL: 300},
		{DNSName: "another.example.com", RecordType: "A", Targets: endpoint.Targets{"target"}},
	}
	changes.Delete = []*endpoint.Endpoint{{DNSName: "delete.example.com", RecordType: "A", Targets: endpoint.Targets{"target"}, RecordTTL: 300}}
	changes.UpdateOld = []*endpoint.Endpoint{{DNSName: "old.example.com", RecordType: "A", Targets: endpoint.Targets{"target-old"}, RecordTTL: 300}}
	changes.UpdateNew = []*endpoint.Endpoint{{DNSName: "update.example.com", Targets: endpoint.Targets{"target-new"}, RecordType: "CNAME", RecordTTL: 300}}
	apply := c.ApplyChanges(t.Context(), changes)
	assert.NoError(t, apply)
}

// TestApplyChangesEmptyIssuesNoCall pins the steady-state path: with nothing to apply the
// provider must not reach Edge DNS at all, not even to list zones.
func TestApplyChangesEmptyIssuesNoCall(t *testing.T) {
	stub := newStub()
	c, err := createAkamaiStubProvider(stub, endpoint.NewDomainFilter([]string{"example.com"}), provider.ZoneIDFilter{})
	require.NoError(t, err)
	stub.setOutput("zone", []any{"example.com"})

	require.NoError(t, c.ApplyChanges(t.Context(), &plan.Changes{}))
	assert.Equal(t, 0, stub.listZonesCalls)
	assert.Equal(t, 0, stub.createCalls)
	assert.Equal(t, 0, stub.updateCalls)
	assert.Equal(t, 0, stub.deleteCalls)
}

// TestZoneCacheAvoidsRefetch pins the zone cache: Records followed by ApplyChanges lists
// the zones once, not twice. With the cache disabled the second listing comes back.
func TestZoneCacheAvoidsRefetch(t *testing.T) {
	for _, tc := range []struct {
		name          string
		cacheDuration time.Duration
		wantListZones int
	}{
		{name: "cache enabled", cacheDuration: time.Hour, wantListZones: 1},
		{name: "cache disabled", cacheDuration: 0, wantListZones: 2},
	} {
		t.Run(tc.name, func(t *testing.T) {
			stub := newStub()
			c, err := createAkamaiStubProviderWithCache(stub, endpoint.NewDomainFilter([]string{"example.com"}), provider.ZoneIDFilter{}, tc.cacheDuration)
			require.NoError(t, err)
			stub.setOutput("zone", []any{"example.com"})
			stub.setOutput("recordset", []any{})

			_, err = c.Records(t.Context())
			require.NoError(t, err)
			changes := &plan.Changes{Create: []*endpoint.Endpoint{
				{DNSName: "www.example.com", RecordType: "A", Targets: endpoint.Targets{"10.0.0.1"}, RecordTTL: 300},
			}}
			require.NoError(t, c.ApplyChanges(t.Context(), changes))

			assert.Equal(t, tc.wantListZones, stub.listZonesCalls)
		})
	}
}
