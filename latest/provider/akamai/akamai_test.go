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
	"testing"

	log "github.com/sirupsen/logrus"

	dns "github.com/akamai/AkamaiOPEN-edgegrid-golang/configdns-v2"
	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

type edgednsStubData struct {
	objType       string // zone, record, recordsets
	output        []interface{}
	updateRecords []interface{}
	createRecords []interface{}
}

type edgednsStub struct {
	stubData map[string]edgednsStubData
}

func newStub() *edgednsStub {
	return &edgednsStub{
		stubData: make(map[string]edgednsStubData),
	}
}

func createAkamaiStubProvider(stub *edgednsStub, domfilter endpoint.DomainFilter, idfilter provider.ZoneIDFilter) (*AkamaiProvider, error) {
	akamaiConfig := AkamaiConfig{
		DomainFilter:          domfilter,
		ZoneIDFilter:          idfilter,
		ServiceConsumerDomain: "testzone.com",
		ClientToken:           "test_token",
		ClientSecret:          "test_client_secret",
		AccessToken:           "test_access_token",
	}

	prov, err := NewAkamaiProvider(akamaiConfig, stub)
	aprov := prov.(*AkamaiProvider)
	return aprov, err
}

func (r *edgednsStub) createStubDataEntry(objtype string) {
	log.Debugf("Creating stub data entry")
	if _, exists := r.stubData[objtype]; !exists {
		r.stubData[objtype] = edgednsStubData{objType: objtype}
	}

	return
}

func (r *edgednsStub) setOutput(objtype string, output []interface{}) {
	log.Debugf("Setting output to %v", output)
	r.createStubDataEntry(objtype)
	stubdata := r.stubData[objtype]
	stubdata.output = output
	r.stubData[objtype] = stubdata

	return
}

func (r *edgednsStub) setUpdateRecords(objtype string, records []interface{}) {
	log.Debugf("Setting updaterecords to %v", records)
	r.createStubDataEntry(objtype)
	stubdata := r.stubData[objtype]
	stubdata.updateRecords = records
	r.stubData[objtype] = stubdata

	return
}

func (r *edgednsStub) setCreateRecords(objtype string, records []interface{}) {
	log.Debugf("Setting createrecords to %v", records)
	r.createStubDataEntry(objtype)
	stubdata := r.stubData[objtype]
	stubdata.createRecords = records
	r.stubData[objtype] = stubdata

	return
}

func (r *edgednsStub) ListZones(queryArgs dns.ZoneListQueryArgs) (*dns.ZoneListResponse, error) {
	log.Debugf("Entering ListZones")
	// Ignore Metadata`
	resp := &dns.ZoneListResponse{}
	zones := make([]*dns.ZoneResponse, 0)
	for _, zname := range r.stubData["zone"].output {
		log.Debugf("Processing output: %v", zname)
		zn := &dns.ZoneResponse{Zone: zname.(string), ContractId: "contract"}
		log.Debugf("Created Zone Object: %v", zn)
		zones = append(zones, zn)
	}
	resp.Zones = zones
	return resp, nil
}

func (r *edgednsStub) GetRecordsets(zone string, queryArgs dns.RecordsetQueryArgs) (*dns.RecordSetResponse, error) {
	log.Debugf("Entering GetRecordsets")
	// Ignore Metadata`
	resp := &dns.RecordSetResponse{}
	sets := make([]dns.Recordset, 0)
	for _, rec := range r.stubData["recordset"].output {
		rset := rec.(dns.Recordset)
		sets = append(sets, rset)
	}
	resp.Recordsets = sets

	return resp, nil
}

func (r *edgednsStub) CreateRecordsets(recordsets *dns.Recordsets, zone string, reclock bool) error {
	return nil
}

func (r *edgednsStub) GetRecord(zone string, name string, record_type string) (*dns.RecordBody, error) {
	resp := &dns.RecordBody{}

	return resp, nil
}

func (r *edgednsStub) DeleteRecord(record *dns.RecordBody, zone string, recLock bool) error {
	return nil
}

func (r *edgednsStub) UpdateRecord(record *dns.RecordBody, zone string, recLock bool) error {
	return nil
}

// Test FetchZones
func TestFetchZonesZoneIDFilter(t *testing.T) {
	stub := newStub()
	domfilter := endpoint.DomainFilter{}
	idfilter := provider.NewZoneIDFilter([]string{"Test"})
	c, err := createAkamaiStubProvider(stub, domfilter, idfilter)
	assert.Nil(t, err)
	stub.setOutput("zone", []interface{}{"test1.testzone.com", "test2.testzone.com"})

	x, _ := c.fetchZones()
	y, _ := json.Marshal(x)
	if assert.NotNil(t, y) {
		assert.Equal(t, "{\"zones\":[{\"contractId\":\"contract\",\"zone\":\"test1.testzone.com\"},{\"contractId\":\"contract\",\"zone\":\"test2.testzone.com\"}]}", string(y))
	}
}

func TestFetchZonesEmpty(t *testing.T) {
	stub := newStub()
	domfilter := endpoint.NewDomainFilter([]string{"Nonexistent"})
	idfilter := provider.NewZoneIDFilter([]string{"Nonexistent"})
	c, err := createAkamaiStubProvider(stub, domfilter, idfilter)
	assert.Nil(t, err)
	stub.setOutput("zone", []interface{}{})

	x, _ := c.fetchZones()
	y, _ := json.Marshal(x)
	if assert.NotNil(t, y) {
		assert.Equal(t, "{\"zones\":[]}", string(y))
	}
}

// TestAkamaiRecords tests record endpoint
func TestAkamaiRecords(t *testing.T) {
	stub := newStub()
	domfilter := endpoint.DomainFilter{}
	idfilter := provider.ZoneIDFilter{}
	c, err := createAkamaiStubProvider(stub, domfilter, idfilter)
	assert.Nil(t, err)
	stub.setOutput("zone", []interface{}{"test1.testzone.com"})
	recordsets := make([]interface{}, 0)
	recordsets = append(recordsets, dns.Recordset{
		Name:  "www.example.com",
		Type:  endpoint.RecordTypeA,
		Rdata: []string{"10.0.0.2", "10.0.0.3"},
	})
	recordsets = append(recordsets, dns.Recordset{
		Name:  "www.example.com",
		Type:  endpoint.RecordTypeTXT,
		Rdata: []string{"heritage=external-dns,external-dns/owner=default"},
	})
	recordsets = append(recordsets, dns.Recordset{
		Name:  "www.exclude.me",
		Type:  endpoint.RecordTypeA,
		Rdata: []string{"192.168.0.1", "192.168.0.2"},
	})
	stub.setOutput("recordset", recordsets)
	endpoints := make([]*endpoint.Endpoint, 0)
	endpoints = append(endpoints, endpoint.NewEndpoint("www.example.com", endpoint.RecordTypeA, "10.0.0.2", "10.0.0.3"))
	endpoints = append(endpoints, endpoint.NewEndpoint("www.example.com", endpoint.RecordTypeTXT, "heritage=external-dns,external-dns/owner=default"))
	endpoints = append(endpoints, endpoint.NewEndpoint("www.exclude.me", endpoint.RecordTypeA, "192.168.0.1", "192.168.0.2"))

	x, _ := c.Records(context.Background())
	if assert.NotNil(t, x) {
		assert.Equal(t, endpoints, x)
	}
}

func TestAkamaiRecordsEmpty(t *testing.T) {
	stub := newStub()
	domfilter := endpoint.DomainFilter{}
	idfilter := provider.NewZoneIDFilter([]string{"Nonexistent"})
	c, err := createAkamaiStubProvider(stub, domfilter, idfilter)
	assert.Nil(t, err)
	stub.setOutput("zone", []interface{}{"test1.testzone.com"})
	recordsets := make([]interface{}, 0)
	stub.setOutput("recordset", recordsets)

	x, _ := c.Records(context.Background())
	assert.Nil(t, x)
}

func TestAkamaiRecordsFilters(t *testing.T) {
	stub := newStub()
	domfilter := endpoint.NewDomainFilter([]string{"www.exclude.me"})
	idfilter := provider.ZoneIDFilter{}
	c, err := createAkamaiStubProvider(stub, domfilter, idfilter)
	assert.Nil(t, err)
	stub.setOutput("zone", []interface{}{"www.exclude.me"})
	recordsets := make([]interface{}, 0)
	recordsets = append(recordsets, dns.Recordset{
		Name:  "www.example.com",
		Type:  endpoint.RecordTypeA,
		Rdata: []string{"10.0.0.2", "10.0.0.3"},
	})
	recordsets = append(recordsets, dns.Recordset{
		Name:  "www.exclude.me",
		Type:  endpoint.RecordTypeA,
		Rdata: []string{"192.168.0.1", "192.168.0.2"},
	})
	stub.setOutput("recordset", recordsets)
	endpoints := make([]*endpoint.Endpoint, 0)
	endpoints = append(endpoints, endpoint.NewEndpoint("www.exclude.me", endpoint.RecordTypeA, "192.168.0.1", "192.168.0.2"))

	x, _ := c.Records(context.Background())
	if assert.NotNil(t, x) {
		assert.Equal(t, endpoints, x)
	}
}

// TestCreateRecords tests create function
// (p AkamaiProvider) createRecordsets(zoneNameIDMapper provider.ZoneIDName, endpoints []*endpoint.Endpoint) error
func TestCreateRecords(t *testing.T) {
	stub := newStub()
	domfilter := endpoint.DomainFilter{}
	idfilter := provider.ZoneIDFilter{}
	c, err := createAkamaiStubProvider(stub, domfilter, idfilter)
	assert.Nil(t, err)

	zoneNameIDMapper := provider.ZoneIDName{"example.com": "example.com"}
	endpoints := make([]*endpoint.Endpoint, 0)
	endpoints = append(endpoints, endpoint.NewEndpoint("www.example.com", endpoint.RecordTypeA, "10.0.0.2", "10.0.0.3"))
	endpoints = append(endpoints, endpoint.NewEndpoint("www.example.com", endpoint.RecordTypeTXT, "heritage=external-dns,external-dns/owner=default"))

	err = c.createRecordsets(zoneNameIDMapper, endpoints)
	assert.Nil(t, err)
}

func TestCreateRecordsDomainFilter(t *testing.T) {
	stub := newStub()
	domfilter := endpoint.DomainFilter{}
	idfilter := provider.ZoneIDFilter{}
	c, err := createAkamaiStubProvider(stub, domfilter, idfilter)
	assert.Nil(t, err)

	zoneNameIDMapper := provider.ZoneIDName{"example.com": "example.com"}
	endpoints := make([]*endpoint.Endpoint, 0)
	endpoints = append(endpoints, endpoint.NewEndpoint("www.example.com", endpoint.RecordTypeA, "10.0.0.2", "10.0.0.3"))
	endpoints = append(endpoints, endpoint.NewEndpoint("www.example.com", endpoint.RecordTypeTXT, "heritage=external-dns,external-dns/owner=default"))
	exclude := append(endpoints, endpoint.NewEndpoint("www.exclude.me", endpoint.RecordTypeA, "10.0.0.2", "10.0.0.3"))

	err = c.createRecordsets(zoneNameIDMapper, exclude)
	assert.Nil(t, err)
}

// TestDeleteRecords validate delete
func TestDeleteRecords(t *testing.T) {
	stub := newStub()
	domfilter := endpoint.DomainFilter{}
	idfilter := provider.ZoneIDFilter{}
	c, err := createAkamaiStubProvider(stub, domfilter, idfilter)
	assert.Nil(t, err)

	zoneNameIDMapper := provider.ZoneIDName{"example.com": "example.com"}
	endpoints := make([]*endpoint.Endpoint, 0)
	endpoints = append(endpoints, endpoint.NewEndpoint("www.example.com", endpoint.RecordTypeA, "10.0.0.2", "10.0.0.3"))
	endpoints = append(endpoints, endpoint.NewEndpoint("www.example.com", endpoint.RecordTypeTXT, "heritage=external-dns,external-dns/owner=default"))

	err = c.deleteRecordsets(zoneNameIDMapper, endpoints)
	assert.Nil(t, err)
}

func TestDeleteRecordsDomainFilter(t *testing.T) {
	stub := newStub()
	domfilter := endpoint.NewDomainFilter([]string{"example.com"})
	idfilter := provider.ZoneIDFilter{}
	c, err := createAkamaiStubProvider(stub, domfilter, idfilter)
	assert.Nil(t, err)

	zoneNameIDMapper := provider.ZoneIDName{"example.com": "example.com"}
	endpoints := make([]*endpoint.Endpoint, 0)
	endpoints = append(endpoints, endpoint.NewEndpoint("www.example.com", endpoint.RecordTypeA, "10.0.0.2", "10.0.0.3"))
	endpoints = append(endpoints, endpoint.NewEndpoint("www.example.com", endpoint.RecordTypeTXT, "heritage=external-dns,external-dns/owner=default"))
	exclude := append(endpoints, endpoint.NewEndpoint("www.exclude.me", endpoint.RecordTypeA, "10.0.0.2", "10.0.0.3"))

	err = c.deleteRecordsets(zoneNameIDMapper, exclude)
	assert.Nil(t, err)
}

// Test record update func
func TestUpdateRecords(t *testing.T) {
	stub := newStub()
	domfilter := endpoint.DomainFilter{}
	idfilter := provider.ZoneIDFilter{}
	c, err := createAkamaiStubProvider(stub, domfilter, idfilter)
	assert.Nil(t, err)

	zoneNameIDMapper := provider.ZoneIDName{"example.com": "example.com"}
	endpoints := make([]*endpoint.Endpoint, 0)
	endpoints = append(endpoints, endpoint.NewEndpoint("www.example.com", endpoint.RecordTypeA, "10.0.0.2", "10.0.0.3"))
	endpoints = append(endpoints, endpoint.NewEndpoint("www.example.com", endpoint.RecordTypeTXT, "heritage=external-dns,external-dns/owner=default"))

	err = c.updateNewRecordsets(zoneNameIDMapper, endpoints)
	assert.Nil(t, err)
}

func TestUpdateRecordsDomainFilter(t *testing.T) {
	stub := newStub()
	domfilter := endpoint.NewDomainFilter([]string{"example.com"})
	idfilter := provider.ZoneIDFilter{}
	c, err := createAkamaiStubProvider(stub, domfilter, idfilter)
	assert.Nil(t, err)

	zoneNameIDMapper := provider.ZoneIDName{"example.com": "example.com"}
	endpoints := make([]*endpoint.Endpoint, 0)
	endpoints = append(endpoints, endpoint.NewEndpoint("www.example.com", endpoint.RecordTypeA, "10.0.0.2", "10.0.0.3"))
	endpoints = append(endpoints, endpoint.NewEndpoint("www.example.com", endpoint.RecordTypeTXT, "heritage=external-dns,external-dns/owner=default"))
	exclude := append(endpoints, endpoint.NewEndpoint("www.exclude.me", endpoint.RecordTypeA, "10.0.0.2", "10.0.0.3"))

	err = c.updateNewRecordsets(zoneNameIDMapper, exclude)
	assert.Nil(t, err)
}

func TestAkamaiApplyChanges(t *testing.T) {
	stub := newStub()
	domfilter := endpoint.NewDomainFilter([]string{"example.com"})
	idfilter := provider.ZoneIDFilter{}
	c, err := createAkamaiStubProvider(stub, domfilter, idfilter)
	assert.Nil(t, err)

	stub.setOutput("zone", []interface{}{"example.com"})
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
	apply := c.ApplyChanges(context.Background(), changes)
	assert.Nil(t, apply)
}
