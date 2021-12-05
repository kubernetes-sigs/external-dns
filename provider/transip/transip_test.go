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

package transip

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/transip/gotransip/v6/domain"
	"github.com/transip/gotransip/v6/rest"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/provider"
)

func newProvider() *TransIPProvider {
	return &TransIPProvider{
		zoneMap: provider.ZoneIDName{},
	}
}

func TestTransIPDnsEntriesAreEqual(t *testing.T) {
	// test with equal set
	a := []domain.DNSEntry{
		{
			Name:    "www.example.org",
			Type:    "CNAME",
			Expire:  3600,
			Content: "www.example.com",
		},
		{
			Name:    "www.example.com",
			Type:    "A",
			Expire:  3600,
			Content: "192.168.0.1",
		},
	}

	b := []domain.DNSEntry{
		{
			Name:    "www.example.com",
			Type:    "A",
			Expire:  3600,
			Content: "192.168.0.1",
		},
		{
			Name:    "www.example.org",
			Type:    "CNAME",
			Expire:  3600,
			Content: "www.example.com",
		},
	}

	assert.Equal(t, true, dnsEntriesAreEqual(a, b))

	// change type on one of b's records
	b[1].Type = "NS"
	assert.Equal(t, false, dnsEntriesAreEqual(a, b))
	b[1].Type = "CNAME"

	// change ttl on one of b's records
	b[1].Expire = 1800
	assert.Equal(t, false, dnsEntriesAreEqual(a, b))
	b[1].Expire = 3600

	// change name on one of b's records
	b[1].Name = "example.org"
	assert.Equal(t, false, dnsEntriesAreEqual(a, b))

	// remove last entry of b
	b = b[:1]
	assert.Equal(t, false, dnsEntriesAreEqual(a, b))
}

func TestTransIPGetMinimalValidTTL(t *testing.T) {
	// test with 'unconfigured' TTL
	ep := &endpoint.Endpoint{}
	assert.EqualValues(t, transipMinimalValidTTL, getMinimalValidTTL(ep))

	// test with lower than minimal ttl
	ep.RecordTTL = (transipMinimalValidTTL - 1)
	assert.EqualValues(t, transipMinimalValidTTL, getMinimalValidTTL(ep))

	// test with higher than minimal ttl
	ep.RecordTTL = (transipMinimalValidTTL + 1)
	assert.EqualValues(t, transipMinimalValidTTL+1, getMinimalValidTTL(ep))
}

func TestTransIPRecordNameForEndpoint(t *testing.T) {
	ep := &endpoint.Endpoint{
		DNSName: "example.org",
	}
	d := domain.Domain{
		Name: "example.org",
	}

	assert.Equal(t, "@", recordNameForEndpoint(ep, d.Name))

	ep.DNSName = "www.example.org"
	assert.Equal(t, "www", recordNameForEndpoint(ep, d.Name))
}

func TestTransIPEndpointNameForRecord(t *testing.T) {
	r := domain.DNSEntry{
		Name: "@",
	}
	d := domain.Domain{
		Name: "example.org",
	}

	assert.Equal(t, d.Name, endpointNameForRecord(r, d.Name))

	r.Name = "www"
	assert.Equal(t, "www.example.org", endpointNameForRecord(r, d.Name))
}

func TestTransIPAddEndpointToEntries(t *testing.T) {
	// prepare endpoint
	ep := &endpoint.Endpoint{
		DNSName:    "www.example.org",
		RecordType: "A",
		RecordTTL:  1800,
		Targets: []string{
			"192.168.0.1",
			"192.168.0.2",
		},
	}

	// prepare zone with DNS entry set
	zone := domain.Domain{
		Name: "example.org",
	}

	// add endpoint to zone's entries
	result := dnsEntriesForEndpoint(ep, zone.Name)

	if assert.Equal(t, 2, len(result)) {
		assert.Equal(t, "www", result[0].Name)
		assert.Equal(t, "A", result[0].Type)
		assert.Equal(t, "192.168.0.1", result[0].Content)
		assert.EqualValues(t, 1800, result[0].Expire)
		assert.Equal(t, "www", result[1].Name)
		assert.Equal(t, "A", result[1].Type)
		assert.Equal(t, "192.168.0.2", result[1].Content)
		assert.EqualValues(t, 1800, result[1].Expire)
	}

	// try again with CNAME
	ep.RecordType = "CNAME"
	ep.Targets = []string{"foo.bar"}
	result = dnsEntriesForEndpoint(ep, zone.Name)
	if assert.Equal(t, 1, len(result)) {
		assert.Equal(t, "CNAME", result[0].Type)
		assert.Equal(t, "foo.bar.", result[0].Content)
	}
}

func TestZoneNameForDNSName(t *testing.T) {
	p := newProvider()
	p.zoneMap.Add("example.com", "example.com")

	zoneName, err := p.zoneNameForDNSName("www.example.com")
	if assert.NoError(t, err) {
		assert.Equal(t, "example.com", zoneName)
	}

	_, err = p.zoneNameForDNSName("www.example.org")
	if assert.Error(t, err) {
		assert.Equal(t, "could not find zoneName for www.example.org", err.Error())
	}
}

// fakeClient mocks the REST API client
type fakeClient struct {
	getFunc func(rest.Request, interface{}) error
}

func (f *fakeClient) Get(request rest.Request, dest interface{}) error {
	if f.getFunc == nil {
		return errors.New("GET not defined")
	}

	return f.getFunc(request, dest)
}

func (f fakeClient) Put(request rest.Request) error {
	return errors.New("PUT not implemented")
}

func (f fakeClient) Post(request rest.Request) error {
	return errors.New("POST not implemented")
}

func (f fakeClient) Delete(request rest.Request) error {
	return errors.New("DELETE not implemented")
}

func (f fakeClient) Patch(request rest.Request) error {
	return errors.New("PATCH not implemented")
}

func TestProviderRecords(t *testing.T) {
	// set up the fake REST client
	client := &fakeClient{}
	client.getFunc = func(req rest.Request, dest interface{}) error {
		var data []byte
		switch {
		case req.Endpoint == "/domains":
			// return list of some domain names
			// names only, other fields are not used
			data = []byte(`{"domains":[{"name":"example.org"}, {"name":"example.com"}]}`)
		case strings.HasSuffix(req.Endpoint, "/dns"):
			// return list of DNS entries
			// also some unsupported types
			data = []byte(`{"dnsEntries":[{"name":"www", "expire":1234, "type":"CNAME", "content":"@"},{"type":"MX"},{"type":"AAAA"}]}`)
		}

		// unmarshal the prepared return data into the given destination type
		return json.Unmarshal(data, &dest)
	}

	// set up provider
	p := newProvider()
	p.domainRepo = domain.Repository{Client: client}

	endpoints, err := p.Records(context.TODO())
	if assert.NoError(t, err) {
		if assert.Equal(t, 4, len(endpoints)) {
			assert.Equal(t, "www.example.org", endpoints[0].DNSName)
			assert.EqualValues(t, "@", endpoints[0].Targets[0])
			assert.Equal(t, "CNAME", endpoints[0].RecordType)
			assert.Equal(t, 0, len(endpoints[0].Labels))
			assert.EqualValues(t, 1234, endpoints[0].RecordTTL)
		}
	}
}

func TestProviderEntriesForEndpoint(t *testing.T) {
	// set up fake REST client
	client := &fakeClient{}

	// set up provider
	p := newProvider()
	p.domainRepo = domain.Repository{Client: client}
	p.zoneMap.Add("example.com", "example.com")

	// get entries for endpoint with unknown zone
	_, _, err := p.entriesForEndpoint(&endpoint.Endpoint{
		DNSName: "www.example.org",
	})
	if assert.Error(t, err) {
		assert.Equal(t, "could not find zoneName for www.example.org", err.Error())
	}

	// get entries for endpoint with known zone but client returns error
	// we leave GET functions undefined so we know which error to expect
	zoneName, _, err := p.entriesForEndpoint(&endpoint.Endpoint{
		DNSName: "www.example.com",
	})
	if assert.Error(t, err) {
		assert.Equal(t, "GET not defined", err.Error())
	}
	assert.Equal(t, "example.com", zoneName)

	// to be able to return a valid set of DNS entries through the API, we define
	// some first, then JSON encode them and have the fake API client's Get function
	// return that
	// in this set are some entries that do and others that don't match the given
	// endpoint
	dnsEntries := []domain.DNSEntry{
		{
			Name:    "www",
			Type:    "A",
			Expire:  3600,
			Content: "1.2.3.4",
		},
		{
			Name:    "ftp",
			Type:    "A",
			Expire:  86400,
			Content: "3.4.5.6",
		},
		{
			Name:    "www",
			Type:    "A",
			Expire:  3600,
			Content: "2.3.4.5",
		},
		{
			Name:    "www",
			Type:    "CNAME",
			Expire:  3600,
			Content: "@",
		},
	}
	var v struct {
		DNSEntries []domain.DNSEntry `json:"dnsEntries"`
	}
	v.DNSEntries = dnsEntries
	returnData, err := json.Marshal(&v)
	require.NoError(t, err)

	// define GET function
	client.getFunc = func(unused rest.Request, dest interface{}) error {
		// unmarshal the prepared return data into the given dnsEntriesWrapper
		return json.Unmarshal(returnData, &dest)
	}
	_, entries, err := p.entriesForEndpoint(&endpoint.Endpoint{
		DNSName:    "www.example.com",
		RecordType: "A",
	})
	if assert.NoError(t, err) {
		if assert.Equal(t, 2, len(entries)) {
			// only first and third entry should be returned
			assert.Equal(t, dnsEntries[0], entries[0])
			assert.Equal(t, dnsEntries[2], entries[1])
		}
	}
}
