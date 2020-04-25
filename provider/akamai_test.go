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
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

type mockAkamaiClient struct {
	mock.Mock
}

func (m *mockAkamaiClient) NewRequest(config edgegrid.Config, met, p string, b io.Reader) (*http.Request, error) {
	switch {
	case met == "GET":
		switch {
		case strings.HasPrefix(p, "https:///config-dns/v2/zones?"):
			b = bytes.NewReader([]byte("{\"zones\":[{\"contractId\":\"Test\",\"zone\":\"example.com\"},{\"contractId\":\"Exclude-Me\",\"zone\":\"exclude.me\"}]}"))
		case strings.HasPrefix(p, "https:///config-dns/v2/zones/example.com/"):
			b = bytes.NewReader([]byte("{\"recordsets\":[{\"name\":\"www.example.com\",\"type\":\"A\",\"ttl\":300,\"rdata\":[\"10.0.0.2\",\"10.0.0.3\"]},{\"name\":\"www.example.com\",\"type\":\"TXT\",\"ttl\":300,\"rdata\":[\"heritage=external-dns,external-dns/owner=default\"]}]}"))
		case strings.HasPrefix(p, "https:///config-dns/v2/zones/exclude.me/"):
			b = bytes.NewReader([]byte("{\"recordsets\":[{\"name\":\"www.exclude.me\",\"type\":\"A\",\"ttl\":300,\"rdata\":[\"192.168.0.1\",\"192.168.0.2\"]}]}"))
		}
	case met == "DELETE":
		b = bytes.NewReader([]byte("{\"title\": \"Success\", \"status\": 200, \"detail\": \"Record deleted\", \"requestId\": \"4321\"}"))
	case met == "ERROR":
		b = bytes.NewReader([]byte("{\"status\": 404 }"))
	}
	req := httptest.NewRequest(met, p, b)
	return req, nil
}

func (m *mockAkamaiClient) Do(config edgegrid.Config, req *http.Request) (*http.Response, error) {
	handler := func(w http.ResponseWriter, r *http.Request) (isError bool) {
		b, _ := ioutil.ReadAll(r.Body)
		io.WriteString(w, string(b))
		return string(b) == "{\"status\": 404 }"
	}
	w := httptest.NewRecorder()
	err := handler(w, req)
	resp := w.Result()

	if err == true {
		resp.StatusCode = 400
	}

	return resp, nil
}

func TestRequestError(t *testing.T) {
	config := AkamaiConfig{}

	client := &mockAkamaiClient{}
	c := NewAkamaiProvider(config)
	c.client = client

	m := "ERROR"
	p := ""
	b := ""
	x, err := c.request(m, p, bytes.NewReader([]byte(b)))
	assert.Nil(t, x)
	assert.NotNil(t, err)
}

func TestFetchZonesZoneIDFilter(t *testing.T) {
	config := AkamaiConfig{
		ZoneIDFilter: NewZoneIDFilter([]string{"Test"}),
	}

	client := &mockAkamaiClient{}
	c := NewAkamaiProvider(config)
	c.client = client

	x, _ := c.fetchZones()
	y, _ := json.Marshal(x)
	if assert.NotNil(t, y) {
		assert.Equal(t, "{\"zones\":[{\"contractId\":\"Test\",\"zone\":\"example.com\"}]}", string(y))
	}
}

func TestFetchZonesEmpty(t *testing.T) {
	config := AkamaiConfig{
		DomainFilter: endpoint.NewDomainFilter([]string{"Nonexistent"}),
		ZoneIDFilter: NewZoneIDFilter([]string{"Nonexistent"}),
	}

	client := &mockAkamaiClient{}
	c := NewAkamaiProvider(config)
	c.client = client

	x, _ := c.fetchZones()
	y, _ := json.Marshal(x)
	if assert.NotNil(t, y) {
		assert.Equal(t, "{\"zones\":null}", string(y))
	}
}

func TestFetchRecordset1(t *testing.T) {
	config := AkamaiConfig{}

	client := &mockAkamaiClient{}
	c := NewAkamaiProvider(config)
	c.client = client

	x, _ := c.fetchRecordSet("example.com")
	y, _ := json.Marshal(x)
	if assert.NotNil(t, y) {
		assert.Equal(t, "{\"recordsets\":[{\"name\":\"www.example.com\",\"type\":\"A\",\"ttl\":300,\"rdata\":[\"10.0.0.2\",\"10.0.0.3\"]},{\"name\":\"www.example.com\",\"type\":\"TXT\",\"ttl\":300,\"rdata\":[\"heritage=external-dns,external-dns/owner=default\"]}]}", string(y))
	}
}

func TestFetchRecordset2(t *testing.T) {
	config := AkamaiConfig{}

	client := &mockAkamaiClient{}
	c := NewAkamaiProvider(config)
	c.client = client

	x, _ := c.fetchRecordSet("exclude.me")
	y, _ := json.Marshal(x)
	if assert.NotNil(t, y) {
		assert.Equal(t, "{\"recordsets\":[{\"name\":\"www.exclude.me\",\"type\":\"A\",\"ttl\":300,\"rdata\":[\"192.168.0.1\",\"192.168.0.2\"]}]}", string(y))
	}
}

func TestAkamaiRecords(t *testing.T) {
	config := AkamaiConfig{}

	client := &mockAkamaiClient{}
	c := NewAkamaiProvider(config)
	c.client = client

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
	config := AkamaiConfig{
		ZoneIDFilter: NewZoneIDFilter([]string{"Nonexistent"}),
	}

	client := &mockAkamaiClient{}
	c := NewAkamaiProvider(config)
	c.client = client

	x, _ := c.Records(context.Background())
	assert.Nil(t, x)
}

func TestAkamaiRecordsFilters(t *testing.T) {
	config := AkamaiConfig{
		DomainFilter: endpoint.NewDomainFilter([]string{"www.exclude.me"}),
		ZoneIDFilter: NewZoneIDFilter([]string{"Exclude-Me"}),
	}

	client := &mockAkamaiClient{}
	c := NewAkamaiProvider(config)
	c.client = client

	endpoints := make([]*endpoint.Endpoint, 0)
	endpoints = append(endpoints, endpoint.NewEndpoint("www.exclude.me", endpoint.RecordTypeA, "192.168.0.1", "192.168.0.2"))

	x, _ := c.Records(context.Background())
	if assert.NotNil(t, x) {
		assert.Equal(t, endpoints, x)
	}
}

func TestCreateRecords(t *testing.T) {
	config := AkamaiConfig{}

	client := &mockAkamaiClient{}
	c := NewAkamaiProvider(config)
	c.client = client

	zoneNameIDMapper := zoneIDName{"example.com": "example.com"}
	endpoints := make([]*endpoint.Endpoint, 0)
	endpoints = append(endpoints, endpoint.NewEndpoint("www.example.com", endpoint.RecordTypeA, "10.0.0.2", "10.0.0.3"))
	endpoints = append(endpoints, endpoint.NewEndpoint("www.example.com", endpoint.RecordTypeTXT, "heritage=external-dns,external-dns/owner=default"))

	x, _ := c.createRecords(zoneNameIDMapper, endpoints)
	if assert.NotNil(t, x) {
		assert.Equal(t, endpoints, x)
	}
}

func TestCreateRecordsDomainFilter(t *testing.T) {
	config := AkamaiConfig{
		DomainFilter: endpoint.NewDomainFilter([]string{"example.com"}),
	}

	client := &mockAkamaiClient{}
	c := NewAkamaiProvider(config)
	c.client = client

	zoneNameIDMapper := zoneIDName{"example.com": "example.com"}
	endpoints := make([]*endpoint.Endpoint, 0)
	endpoints = append(endpoints, endpoint.NewEndpoint("www.example.com", endpoint.RecordTypeA, "10.0.0.2", "10.0.0.3"))
	endpoints = append(endpoints, endpoint.NewEndpoint("www.example.com", endpoint.RecordTypeTXT, "heritage=external-dns,external-dns/owner=default"))
	exclude := append(endpoints, endpoint.NewEndpoint("www.exclude.me", endpoint.RecordTypeA, "10.0.0.2", "10.0.0.3"))

	x, _ := c.createRecords(zoneNameIDMapper, exclude)
	if assert.NotNil(t, x) {
		assert.Equal(t, endpoints, x)
	}
}

func TestDeleteRecords(t *testing.T) {
	config := AkamaiConfig{}

	client := &mockAkamaiClient{}
	c := NewAkamaiProvider(config)
	c.client = client

	zoneNameIDMapper := zoneIDName{"example.com": "example.com"}
	endpoints := make([]*endpoint.Endpoint, 0)
	endpoints = append(endpoints, endpoint.NewEndpoint("www.example.com", endpoint.RecordTypeA, "10.0.0.2", "10.0.0.3"))
	endpoints = append(endpoints, endpoint.NewEndpoint("www.example.com", endpoint.RecordTypeTXT, "heritage=external-dns,external-dns/owner=default"))

	x, _ := c.deleteRecords(zoneNameIDMapper, endpoints)
	if assert.NotNil(t, x) {
		assert.Equal(t, endpoints, x)
	}
}

func TestDeleteRecordsDomainFilter(t *testing.T) {
	config := AkamaiConfig{
		DomainFilter: endpoint.NewDomainFilter([]string{"example.com"}),
	}

	client := &mockAkamaiClient{}
	c := NewAkamaiProvider(config)
	c.client = client

	zoneNameIDMapper := zoneIDName{"example.com": "example.com"}
	endpoints := make([]*endpoint.Endpoint, 0)
	endpoints = append(endpoints, endpoint.NewEndpoint("www.example.com", endpoint.RecordTypeA, "10.0.0.2", "10.0.0.3"))
	endpoints = append(endpoints, endpoint.NewEndpoint("www.example.com", endpoint.RecordTypeTXT, "heritage=external-dns,external-dns/owner=default"))
	exclude := append(endpoints, endpoint.NewEndpoint("www.exclude.me", endpoint.RecordTypeA, "10.0.0.2", "10.0.0.3"))

	x, _ := c.deleteRecords(zoneNameIDMapper, exclude)
	if assert.NotNil(t, x) {
		assert.Equal(t, endpoints, x)
	}
}

func TestUpdateRecords(t *testing.T) {
	config := AkamaiConfig{}

	client := &mockAkamaiClient{}
	c := NewAkamaiProvider(config)
	c.client = client

	zoneNameIDMapper := zoneIDName{"example.com": "example.com"}
	endpoints := make([]*endpoint.Endpoint, 0)
	endpoints = append(endpoints, endpoint.NewEndpoint("www.example.com", endpoint.RecordTypeA, "10.0.0.2", "10.0.0.3"))
	endpoints = append(endpoints, endpoint.NewEndpoint("www.example.com", endpoint.RecordTypeTXT, "heritage=external-dns,external-dns/owner=default"))

	x, _ := c.updateNewRecords(zoneNameIDMapper, endpoints)
	if assert.NotNil(t, x) {
		assert.Equal(t, endpoints, x)
	}
}

func TestUpdateRecordsDomainFilter(t *testing.T) {
	config := AkamaiConfig{
		DomainFilter: endpoint.NewDomainFilter([]string{"example.com"}),
	}

	client := &mockAkamaiClient{}
	c := NewAkamaiProvider(config)
	c.client = client

	zoneNameIDMapper := zoneIDName{"example.com": "example.com"}
	endpoints := make([]*endpoint.Endpoint, 0)
	endpoints = append(endpoints, endpoint.NewEndpoint("www.example.com", endpoint.RecordTypeA, "10.0.0.2", "10.0.0.3"))
	endpoints = append(endpoints, endpoint.NewEndpoint("www.example.com", endpoint.RecordTypeTXT, "heritage=external-dns,external-dns/owner=default"))
	exclude := append(endpoints, endpoint.NewEndpoint("www.exclude.me", endpoint.RecordTypeA, "10.0.0.2", "10.0.0.3"))

	x, _ := c.updateNewRecords(zoneNameIDMapper, exclude)
	if assert.NotNil(t, x) {
		assert.Equal(t, endpoints, x)
	}
}

func TestAkamaiApplyChanges(t *testing.T) {
	config := AkamaiConfig{}

	client := &mockAkamaiClient{}
	c := NewAkamaiProvider(config)
	c.client = client

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
