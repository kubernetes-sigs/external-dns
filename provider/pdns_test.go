/*
Copyright 2018 The Kubernetes Authors.

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
	"context"
	"errors"
	"net/http"
	"strings"
	"testing"

	pgo "github.com/ffledgling/pdns-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"sigs.k8s.io/external-dns/endpoint"
)

// FIXME: What do we do about labels?

var (
	// Simple RRSets that contain 1 A record and 1 TXT record
	RRSetSimpleARecord = pgo.RrSet{
		Name:  "example.com.",
		Type_: "A",
		Ttl:   300,
		Records: []pgo.Record{
			{Content: "8.8.8.8", Disabled: false, SetPtr: false},
		},
	}
	RRSetSimpleTXTRecord = pgo.RrSet{
		Name:  "example.com.",
		Type_: "TXT",
		Ttl:   300,
		Records: []pgo.Record{
			{Content: "\"heritage=external-dns,external-dns/owner=tower-pdns\"", Disabled: false, SetPtr: false},
		},
	}
	RRSetLongARecord = pgo.RrSet{
		Name:  "a.very.long.domainname.example.com.",
		Type_: "A",
		Ttl:   300,
		Records: []pgo.Record{
			{Content: "8.8.8.8", Disabled: false, SetPtr: false},
		},
	}
	RRSetLongTXTRecord = pgo.RrSet{
		Name:  "a.very.long.domainname.example.com.",
		Type_: "TXT",
		Ttl:   300,
		Records: []pgo.Record{
			{Content: "\"heritage=external-dns,external-dns/owner=tower-pdns\"", Disabled: false, SetPtr: false},
		},
	}
	// RRSet with one record disabled
	RRSetDisabledRecord = pgo.RrSet{
		Name:  "example.com.",
		Type_: "A",
		Ttl:   300,
		Records: []pgo.Record{
			{Content: "8.8.8.8", Disabled: false, SetPtr: false},
			{Content: "8.8.4.4", Disabled: true, SetPtr: false},
		},
	}

	RRSetCNAMERecord = pgo.RrSet{
		Name:  "cname.example.com.",
		Type_: "CNAME",
		Ttl:   300,
		Records: []pgo.Record{
			{Content: "example.by.any.other.name.com", Disabled: false, SetPtr: false},
		},
	}
	RRSetTXTRecord = pgo.RrSet{
		Name:  "example.com.",
		Type_: "TXT",
		Ttl:   300,
		Records: []pgo.Record{
			{Content: "'would smell as sweet'", Disabled: false, SetPtr: false},
		},
	}

	// Multiple PDNS records in an RRSet of a single type
	RRSetMultipleRecords = pgo.RrSet{
		Name:  "example.com.",
		Type_: "A",
		Ttl:   300,
		Records: []pgo.Record{
			{Content: "8.8.8.8", Disabled: false, SetPtr: false},
			{Content: "8.8.4.4", Disabled: false, SetPtr: false},
			{Content: "4.4.4.4", Disabled: false, SetPtr: false},
		},
	}

	endpointsDisabledRecord = []*endpoint.Endpoint{
		endpoint.NewEndpointWithTTL("example.com", endpoint.RecordTypeA, endpoint.TTL(300), "8.8.8.8"),
	}

	endpointsSimpleRecord = []*endpoint.Endpoint{
		endpoint.NewEndpointWithTTL("example.com", endpoint.RecordTypeA, endpoint.TTL(300), "8.8.8.8"),
		endpoint.NewEndpointWithTTL("example.com", endpoint.RecordTypeTXT, endpoint.TTL(300), "\"heritage=external-dns,external-dns/owner=tower-pdns\""),
	}

	endpointsLongRecord = []*endpoint.Endpoint{
		endpoint.NewEndpointWithTTL("a.very.long.domainname.example.com", endpoint.RecordTypeA, endpoint.TTL(300), "8.8.8.8"),
		endpoint.NewEndpointWithTTL("a.very.long.domainname.example.com", endpoint.RecordTypeTXT, endpoint.TTL(300), "\"heritage=external-dns,external-dns/owner=tower-pdns\""),
	}

	endpointsNonexistantZone = []*endpoint.Endpoint{
		endpoint.NewEndpointWithTTL("does.not.exist.com", endpoint.RecordTypeA, endpoint.TTL(300), "8.8.8.8"),
		endpoint.NewEndpointWithTTL("does.not.exist.com", endpoint.RecordTypeTXT, endpoint.TTL(300), "\"heritage=external-dns,external-dns/owner=tower-pdns\""),
	}
	endpointsMultipleRecords = []*endpoint.Endpoint{
		endpoint.NewEndpointWithTTL("example.com", endpoint.RecordTypeA, endpoint.TTL(300), "8.8.8.8"),
		endpoint.NewEndpointWithTTL("example.com", endpoint.RecordTypeA, endpoint.TTL(300), "8.8.4.4"),
		endpoint.NewEndpointWithTTL("example.com", endpoint.RecordTypeA, endpoint.TTL(300), "4.4.4.4"),
	}

	endpointsMixedRecords = []*endpoint.Endpoint{
		endpoint.NewEndpointWithTTL("cname.example.com", endpoint.RecordTypeCNAME, endpoint.TTL(300), "example.by.any.other.name.com"),
		endpoint.NewEndpointWithTTL("example.com", endpoint.RecordTypeTXT, endpoint.TTL(300), "'would smell as sweet'"),
		endpoint.NewEndpointWithTTL("example.com", endpoint.RecordTypeA, endpoint.TTL(300), "8.8.8.8"),
		endpoint.NewEndpointWithTTL("example.com", endpoint.RecordTypeA, endpoint.TTL(300), "8.8.4.4"),
		endpoint.NewEndpointWithTTL("example.com", endpoint.RecordTypeA, endpoint.TTL(300), "4.4.4.4"),
	}

	endpointsMultipleZones = []*endpoint.Endpoint{
		endpoint.NewEndpointWithTTL("example.com", endpoint.RecordTypeA, endpoint.TTL(300), "8.8.8.8"),
		endpoint.NewEndpointWithTTL("example.com", endpoint.RecordTypeTXT, endpoint.TTL(300), "\"heritage=external-dns,external-dns/owner=tower-pdns\""),
		endpoint.NewEndpointWithTTL("mock.test", endpoint.RecordTypeA, endpoint.TTL(300), "9.9.9.9"),
		endpoint.NewEndpointWithTTL("mock.test", endpoint.RecordTypeTXT, endpoint.TTL(300), "\"heritage=external-dns,external-dns/owner=tower-pdns\""),
	}

	endpointsMultipleZones2 = []*endpoint.Endpoint{
		endpoint.NewEndpointWithTTL("example.com", endpoint.RecordTypeA, endpoint.TTL(300), "8.8.8.8"),
		endpoint.NewEndpointWithTTL("example.com", endpoint.RecordTypeTXT, endpoint.TTL(300), "\"heritage=external-dns,external-dns/owner=tower-pdns\""),
		endpoint.NewEndpointWithTTL("abcd.mock.test", endpoint.RecordTypeA, endpoint.TTL(300), "9.9.9.9"),
		endpoint.NewEndpointWithTTL("abcd.mock.test", endpoint.RecordTypeTXT, endpoint.TTL(300), "\"heritage=external-dns,external-dns/owner=tower-pdns\""),
	}

	endpointsMultipleZonesWithNoExist = []*endpoint.Endpoint{
		endpoint.NewEndpointWithTTL("example.com", endpoint.RecordTypeA, endpoint.TTL(300), "8.8.8.8"),
		endpoint.NewEndpointWithTTL("example.com", endpoint.RecordTypeTXT, endpoint.TTL(300), "\"heritage=external-dns,external-dns/owner=tower-pdns\""),
		endpoint.NewEndpointWithTTL("abcd.mock.noexist", endpoint.RecordTypeA, endpoint.TTL(300), "9.9.9.9"),
		endpoint.NewEndpointWithTTL("abcd.mock.noexist", endpoint.RecordTypeTXT, endpoint.TTL(300), "\"heritage=external-dns,external-dns/owner=tower-pdns\""),
	}
	endpointsMultipleZonesWithLongRecordNotInDomainFilter = []*endpoint.Endpoint{
		endpoint.NewEndpointWithTTL("example.com", endpoint.RecordTypeA, endpoint.TTL(300), "8.8.8.8"),
		endpoint.NewEndpointWithTTL("example.com", endpoint.RecordTypeTXT, endpoint.TTL(300), "\"heritage=external-dns,external-dns/owner=tower-pdns\""),
		endpoint.NewEndpointWithTTL("a.very.long.domainname.example.com", endpoint.RecordTypeA, endpoint.TTL(300), "9.9.9.9"),
		endpoint.NewEndpointWithTTL("a.very.long.domainname.example.com", endpoint.RecordTypeTXT, endpoint.TTL(300), "\"heritage=external-dns,external-dns/owner=tower-pdns\""),
	}
	endpointsMultipleZonesWithSimilarRecordNotInDomainFilter = []*endpoint.Endpoint{
		endpoint.NewEndpointWithTTL("example.com", endpoint.RecordTypeA, endpoint.TTL(300), "8.8.8.8"),
		endpoint.NewEndpointWithTTL("example.com", endpoint.RecordTypeTXT, endpoint.TTL(300), "\"heritage=external-dns,external-dns/owner=tower-pdns\""),
		endpoint.NewEndpointWithTTL("test.simexample.com", endpoint.RecordTypeA, endpoint.TTL(300), "9.9.9.9"),
		endpoint.NewEndpointWithTTL("test.simexample.com", endpoint.RecordTypeTXT, endpoint.TTL(300), "\"heritage=external-dns,external-dns/owner=tower-pdns\""),
	}

	ZoneEmpty = pgo.Zone{
		// Opaque zone id (string), assigned by the server, should not be interpreted by the application. Guaranteed to be safe for embedding in URLs.
		Id: "example.com.",
		// Name of the zone (e.g. “example.com.”) MUST have a trailing dot
		Name: "example.com.",
		// Set to “Zone”
		Type_: "Zone",
		// API endpoint for this zone
		Url: "/api/v1/servers/localhost/zones/example.com.",
		// Zone kind, one of “Native”, “Master”, “Slave”
		Kind: "Native",
		// RRSets in this zone
		Rrsets: []pgo.RrSet{},
	}

	ZoneEmptySimilar = pgo.Zone{
		Id:     "simexample.com.",
		Name:   "simexample.com.",
		Type_:  "Zone",
		Url:    "/api/v1/servers/localhost/zones/simexample.com.",
		Kind:   "Native",
		Rrsets: []pgo.RrSet{},
	}

	ZoneEmptyLong = pgo.Zone{
		Id:     "long.domainname.example.com.",
		Name:   "long.domainname.example.com.",
		Type_:  "Zone",
		Url:    "/api/v1/servers/localhost/zones/long.domainname.example.com.",
		Kind:   "Native",
		Rrsets: []pgo.RrSet{},
	}

	ZoneEmpty2 = pgo.Zone{
		Id:     "mock.test.",
		Name:   "mock.test.",
		Type_:  "Zone",
		Url:    "/api/v1/servers/localhost/zones/mock.test.",
		Kind:   "Native",
		Rrsets: []pgo.RrSet{},
	}

	ZoneMixed = pgo.Zone{
		Id:     "example.com.",
		Name:   "example.com.",
		Type_:  "Zone",
		Url:    "/api/v1/servers/localhost/zones/example.com.",
		Kind:   "Native",
		Rrsets: []pgo.RrSet{RRSetCNAMERecord, RRSetTXTRecord, RRSetMultipleRecords},
	}

	ZoneEmptyToSimplePatch = pgo.Zone{
		Id:    "example.com.",
		Name:  "example.com.",
		Type_: "Zone",
		Url:   "/api/v1/servers/localhost/zones/example.com.",
		Kind:  "Native",
		Rrsets: []pgo.RrSet{
			{
				Name:       "example.com.",
				Type_:      "A",
				Ttl:        300,
				Changetype: "REPLACE",
				Records: []pgo.Record{
					{
						Content:  "8.8.8.8",
						Disabled: false,
						SetPtr:   false,
					},
				},
				Comments: []pgo.Comment(nil),
			},
			{
				Name:       "example.com.",
				Type_:      "TXT",
				Ttl:        300,
				Changetype: "REPLACE",
				Records: []pgo.Record{
					{
						Content:  "\"heritage=external-dns,external-dns/owner=tower-pdns\"",
						Disabled: false,
						SetPtr:   false,
					},
				},
				Comments: []pgo.Comment(nil),
			},
		},
	}

	ZoneEmptyToSimplePatchLongRecordIgnoredInDomainFilter = pgo.Zone{
		Id:    "example.com.",
		Name:  "example.com.",
		Type_: "Zone",
		Url:   "/api/v1/servers/localhost/zones/example.com.",
		Kind:  "Native",
		Rrsets: []pgo.RrSet{
			{
				Name:       "a.very.long.domainname.example.com.",
				Type_:      "A",
				Ttl:        300,
				Changetype: "REPLACE",
				Records: []pgo.Record{
					{
						Content:  "9.9.9.9",
						Disabled: false,
						SetPtr:   false,
					},
				},
				Comments: []pgo.Comment(nil),
			},
			{
				Name:       "a.very.long.domainname.example.com.",
				Type_:      "TXT",
				Ttl:        300,
				Changetype: "REPLACE",
				Records: []pgo.Record{
					{
						Content:  "\"heritage=external-dns,external-dns/owner=tower-pdns\"",
						Disabled: false,
						SetPtr:   false,
					},
				},
				Comments: []pgo.Comment(nil),
			},
			{
				Name:       "example.com.",
				Type_:      "A",
				Ttl:        300,
				Changetype: "REPLACE",
				Records: []pgo.Record{
					{
						Content:  "8.8.8.8",
						Disabled: false,
						SetPtr:   false,
					},
				},
				Comments: []pgo.Comment(nil),
			},
			{
				Name:       "example.com.",
				Type_:      "TXT",
				Ttl:        300,
				Changetype: "REPLACE",
				Records: []pgo.Record{
					{
						Content:  "\"heritage=external-dns,external-dns/owner=tower-pdns\"",
						Disabled: false,
						SetPtr:   false,
					},
				},
				Comments: []pgo.Comment(nil),
			},
		},
	}

	ZoneEmptyToLongPatch = pgo.Zone{
		Id:    "long.domainname.example.com.",
		Name:  "long.domainname.example.com.",
		Type_: "Zone",
		Url:   "/api/v1/servers/localhost/zones/long.domainname.example.com.",
		Kind:  "Native",
		Rrsets: []pgo.RrSet{
			{
				Name:       "a.very.long.domainname.example.com.",
				Type_:      "A",
				Ttl:        300,
				Changetype: "REPLACE",
				Records: []pgo.Record{
					{
						Content:  "8.8.8.8",
						Disabled: false,
						SetPtr:   false,
					},
				},
				Comments: []pgo.Comment(nil),
			},
			{
				Name:       "a.very.long.domainname.example.com.",
				Type_:      "TXT",
				Ttl:        300,
				Changetype: "REPLACE",
				Records: []pgo.Record{
					{
						Content:  "\"heritage=external-dns,external-dns/owner=tower-pdns\"",
						Disabled: false,
						SetPtr:   false,
					},
				},
				Comments: []pgo.Comment(nil),
			},
		},
	}

	ZoneEmptyToSimplePatch2 = pgo.Zone{
		Id:    "mock.test.",
		Name:  "mock.test.",
		Type_: "Zone",
		Url:   "/api/v1/servers/localhost/zones/mock.test.",
		Kind:  "Native",
		Rrsets: []pgo.RrSet{
			{
				Name:       "mock.test.",
				Type_:      "A",
				Ttl:        300,
				Changetype: "REPLACE",
				Records: []pgo.Record{
					{
						Content:  "9.9.9.9",
						Disabled: false,
						SetPtr:   false,
					},
				},
				Comments: []pgo.Comment(nil),
			},
			{
				Name:       "mock.test.",
				Type_:      "TXT",
				Ttl:        300,
				Changetype: "REPLACE",
				Records: []pgo.Record{
					{
						Content:  "\"heritage=external-dns,external-dns/owner=tower-pdns\"",
						Disabled: false,
						SetPtr:   false,
					},
				},
				Comments: []pgo.Comment(nil),
			},
		},
	}

	ZoneEmptyToSimplePatch3 = pgo.Zone{
		Id:    "mock.test.",
		Name:  "mock.test.",
		Type_: "Zone",
		Url:   "/api/v1/servers/localhost/zones/mock.test.",
		Kind:  "Native",
		Rrsets: []pgo.RrSet{
			{
				Name:       "abcd.mock.test.",
				Type_:      "A",
				Ttl:        300,
				Changetype: "REPLACE",
				Records: []pgo.Record{
					{
						Content:  "9.9.9.9",
						Disabled: false,
						SetPtr:   false,
					},
				},
				Comments: []pgo.Comment(nil),
			},
			{
				Name:       "abcd.mock.test.",
				Type_:      "TXT",
				Ttl:        300,
				Changetype: "REPLACE",
				Records: []pgo.Record{
					{
						Content:  "\"heritage=external-dns,external-dns/owner=tower-pdns\"",
						Disabled: false,
						SetPtr:   false,
					},
				},
				Comments: []pgo.Comment(nil),
			},
		},
	}

	ZoneEmptyToSimpleDelete = pgo.Zone{
		Id:    "example.com.",
		Name:  "example.com.",
		Type_: "Zone",
		Url:   "/api/v1/servers/localhost/zones/example.com.",
		Kind:  "Native",
		Rrsets: []pgo.RrSet{
			{
				Name:       "example.com.",
				Type_:      "A",
				Changetype: "DELETE",
				Records: []pgo.Record{
					{
						Content:  "8.8.8.8",
						Disabled: false,
						SetPtr:   false,
					},
				},
				Comments: []pgo.Comment(nil),
			},
			{
				Name:       "example.com.",
				Type_:      "TXT",
				Changetype: "DELETE",
				Records: []pgo.Record{
					{
						Content:  "\"heritage=external-dns,external-dns/owner=tower-pdns\"",
						Disabled: false,
						SetPtr:   false,
					},
				},
				Comments: []pgo.Comment(nil),
			},
		},
	}

	DomainFilterListSingle = endpoint.DomainFilter{
		Filters: []string{
			"example.com",
		},
	}

	DomainFilterListMultiple = endpoint.DomainFilter{
		Filters: []string{
			"example.com",
			"mock.com",
		},
	}

	DomainFilterListEmpty = endpoint.DomainFilter{
		Filters: []string{},
	}

	DomainFilterEmptyClient = &PDNSAPIClient{
		dryRun:       false,
		authCtx:      context.WithValue(context.Background(), pgo.ContextAPIKey, pgo.APIKey{Key: "TEST-API-KEY"}),
		client:       pgo.NewAPIClient(pgo.NewConfiguration()),
		domainFilter: DomainFilterListEmpty,
	}

	DomainFilterSingleClient = &PDNSAPIClient{
		dryRun:       false,
		authCtx:      context.WithValue(context.Background(), pgo.ContextAPIKey, pgo.APIKey{Key: "TEST-API-KEY"}),
		client:       pgo.NewAPIClient(pgo.NewConfiguration()),
		domainFilter: DomainFilterListSingle,
	}

	DomainFilterMultipleClient = &PDNSAPIClient{
		dryRun:       false,
		authCtx:      context.WithValue(context.Background(), pgo.ContextAPIKey, pgo.APIKey{Key: "TEST-API-KEY"}),
		client:       pgo.NewAPIClient(pgo.NewConfiguration()),
		domainFilter: DomainFilterListMultiple,
	}
)

/******************************************************************************/
// API that returns a zone with multiple record types
type PDNSAPIClientStub struct {
}

func (c *PDNSAPIClientStub) ListZones() ([]pgo.Zone, *http.Response, error) {
	return []pgo.Zone{ZoneMixed}, nil, nil
}
func (c *PDNSAPIClientStub) PartitionZones(zones []pgo.Zone) ([]pgo.Zone, []pgo.Zone) {
	return zones, nil
}
func (c *PDNSAPIClientStub) ListZone(zoneID string) (pgo.Zone, *http.Response, error) {
	return ZoneMixed, nil, nil
}
func (c *PDNSAPIClientStub) PatchZone(zoneID string, zoneStruct pgo.Zone) (*http.Response, error) {
	return nil, nil
}

/******************************************************************************/
// API that returns a zones with no records
type PDNSAPIClientStubEmptyZones struct {
	// Keep track of all zones we receive via PatchZone
	patchedZones []pgo.Zone
}

func (c *PDNSAPIClientStubEmptyZones) ListZones() ([]pgo.Zone, *http.Response, error) {
	return []pgo.Zone{ZoneEmpty, ZoneEmptyLong, ZoneEmpty2}, nil, nil
}
func (c *PDNSAPIClientStubEmptyZones) PartitionZones(zones []pgo.Zone) ([]pgo.Zone, []pgo.Zone) {
	return zones, nil
}
func (c *PDNSAPIClientStubEmptyZones) ListZone(zoneID string) (pgo.Zone, *http.Response, error) {

	if strings.Contains(zoneID, "example.com") {
		return ZoneEmpty, nil, nil
	} else if strings.Contains(zoneID, "mock.test") {
		return ZoneEmpty2, nil, nil
	} else if strings.Contains(zoneID, "long.domainname.example.com") {
		return ZoneEmptyLong, nil, nil
	}
	return pgo.Zone{}, nil, nil

}
func (c *PDNSAPIClientStubEmptyZones) PatchZone(zoneID string, zoneStruct pgo.Zone) (*http.Response, error) {
	c.patchedZones = append(c.patchedZones, zoneStruct)
	return nil, nil
}

/******************************************************************************/
// API that returns error on PatchZone()
type PDNSAPIClientStubPatchZoneFailure struct {
	// Anonymous struct for composition
	PDNSAPIClientStubEmptyZones
}

// Just overwrite the PatchZone method to introduce a failure
func (c *PDNSAPIClientStubPatchZoneFailure) PatchZone(zoneID string, zoneStruct pgo.Zone) (*http.Response, error) {
	return nil, errors.New("Generic PDNS Error")
}

/******************************************************************************/
// API that returns error on ListZone()
type PDNSAPIClientStubListZoneFailure struct {
	// Anonymous struct for composition
	PDNSAPIClientStubEmptyZones
}

// Just overwrite the ListZone method to introduce a failure
func (c *PDNSAPIClientStubListZoneFailure) ListZone(zoneID string) (pgo.Zone, *http.Response, error) {
	return pgo.Zone{}, nil, errors.New("Generic PDNS Error")

}

/******************************************************************************/
// API that returns error on ListZones() (Zones - plural)
type PDNSAPIClientStubListZonesFailure struct {
	// Anonymous struct for composition
	PDNSAPIClientStubEmptyZones
}

// Just overwrite the ListZones method to introduce a failure
func (c *PDNSAPIClientStubListZonesFailure) ListZones() ([]pgo.Zone, *http.Response, error) {
	return []pgo.Zone{}, nil, errors.New("Generic PDNS Error")
}

/******************************************************************************/
// API that returns zone partitions given DomainFilter(s)
type PDNSAPIClientStubPartitionZones struct {
	// Anonymous struct for composition
	PDNSAPIClientStubEmptyZones
}

func (c *PDNSAPIClientStubPartitionZones) ListZones() ([]pgo.Zone, *http.Response, error) {
	return []pgo.Zone{ZoneEmpty, ZoneEmptyLong, ZoneEmpty2, ZoneEmptySimilar}, nil, nil
}

func (c *PDNSAPIClientStubPartitionZones) ListZone(zoneID string) (pgo.Zone, *http.Response, error) {

	if strings.Contains(zoneID, "example.com") {
		return ZoneEmpty, nil, nil
	} else if strings.Contains(zoneID, "mock.test") {
		return ZoneEmpty2, nil, nil
	} else if strings.Contains(zoneID, "long.domainname.example.com") {
		return ZoneEmptyLong, nil, nil
	} else if strings.Contains(zoneID, "simexample.com") {
		return ZoneEmptySimilar, nil, nil
	}
	return pgo.Zone{}, nil, nil
}

// Just overwrite the ListZones method to introduce a failure
func (c *PDNSAPIClientStubPartitionZones) PartitionZones(zones []pgo.Zone) ([]pgo.Zone, []pgo.Zone) {
	return []pgo.Zone{ZoneEmpty}, []pgo.Zone{ZoneEmptyLong, ZoneEmpty2}

}

/******************************************************************************/

type NewPDNSProviderTestSuite struct {
	suite.Suite
}

func (suite *NewPDNSProviderTestSuite) TestPDNSProviderCreate() {

	_, err := NewPDNSProvider(
		context.Background(),
		PDNSConfig{
			Server:       "http://localhost:8081",
			DomainFilter: endpoint.NewDomainFilter([]string{""}),
		})
	assert.Error(suite.T(), err, "--pdns-api-key should be specified")

	_, err = NewPDNSProvider(
		context.Background(),
		PDNSConfig{
			Server:       "http://localhost:8081",
			APIKey:       "foo",
			DomainFilter: endpoint.NewDomainFilter([]string{"example.com", "example.org"}),
		})
	assert.Nil(suite.T(), err, "--domain-filter should raise no error")

	_, err = NewPDNSProvider(
		context.Background(),
		PDNSConfig{
			Server:       "http://localhost:8081",
			APIKey:       "foo",
			DomainFilter: endpoint.NewDomainFilter([]string{""}),
			DryRun:       true,
		})
	assert.Error(suite.T(), err, "--dry-run should raise an error")

	// This is our "regular" code path, no error should be thrown
	_, err = NewPDNSProvider(
		context.Background(),
		PDNSConfig{
			Server:       "http://localhost:8081",
			APIKey:       "foo",
			DomainFilter: endpoint.NewDomainFilter([]string{""}),
		})
	assert.Nil(suite.T(), err, "Regular case should raise no error")
}

func (suite *NewPDNSProviderTestSuite) TestPDNSProviderCreateTLS() {

	_, err := NewPDNSProvider(
		context.Background(),
		PDNSConfig{
			Server:       "http://localhost:8081",
			APIKey:       "foo",
			DomainFilter: endpoint.NewDomainFilter([]string{""}),
		})
	assert.Nil(suite.T(), err, "Omitted TLS Config case should raise no error")

	_, err = NewPDNSProvider(
		context.Background(),
		PDNSConfig{
			Server:       "http://localhost:8081",
			APIKey:       "foo",
			DomainFilter: endpoint.NewDomainFilter([]string{""}),
			TLSConfig: TLSConfig{
				TLSEnabled: false,
			},
		})
	assert.Nil(suite.T(), err, "Disabled TLS Config should raise no error")

	_, err = NewPDNSProvider(
		context.Background(),
		PDNSConfig{
			Server:       "http://localhost:8081",
			APIKey:       "foo",
			DomainFilter: endpoint.NewDomainFilter([]string{""}),
			TLSConfig: TLSConfig{
				TLSEnabled:            false,
				CAFilePath:            "/path/to/ca.crt",
				ClientCertFilePath:    "/path/to/cert.pem",
				ClientCertKeyFilePath: "/path/to/cert-key.pem",
			},
		})
	assert.Nil(suite.T(), err, "Disabled TLS Config with additional flags should raise no error")

	_, err = NewPDNSProvider(
		context.Background(),
		PDNSConfig{
			Server:       "http://localhost:8081",
			APIKey:       "foo",
			DomainFilter: endpoint.NewDomainFilter([]string{""}),
			TLSConfig: TLSConfig{
				TLSEnabled: true,
			},
		})
	assert.Error(suite.T(), err, "Enabled TLS Config without --tls-ca should raise an error")

	_, err = NewPDNSProvider(
		context.Background(),
		PDNSConfig{
			Server:       "http://localhost:8081",
			APIKey:       "foo",
			DomainFilter: endpoint.NewDomainFilter([]string{""}),
			TLSConfig: TLSConfig{
				TLSEnabled: true,
				CAFilePath: "../internal/testresources/ca.pem",
			},
		})
	assert.Nil(suite.T(), err, "Enabled TLS Config with --tls-ca should raise no error")

	_, err = NewPDNSProvider(
		context.Background(),
		PDNSConfig{
			Server:       "http://localhost:8081",
			APIKey:       "foo",
			DomainFilter: endpoint.NewDomainFilter([]string{""}),
			TLSConfig: TLSConfig{
				TLSEnabled:         true,
				CAFilePath:         "../internal/testresources/ca.pem",
				ClientCertFilePath: "../internal/testresources/client-cert.pem",
			},
		})
	assert.Error(suite.T(), err, "Enabled TLS Config with --tls-client-cert only should raise an error")

	_, err = NewPDNSProvider(
		context.Background(),
		PDNSConfig{
			Server:       "http://localhost:8081",
			APIKey:       "foo",
			DomainFilter: endpoint.NewDomainFilter([]string{""}),
			TLSConfig: TLSConfig{
				TLSEnabled:            true,
				CAFilePath:            "../internal/testresources/ca.pem",
				ClientCertKeyFilePath: "../internal/testresources/client-cert-key.pem",
			},
		})
	assert.Error(suite.T(), err, "Enabled TLS Config with --tls-client-cert-key only should raise an error")

	_, err = NewPDNSProvider(
		context.Background(),
		PDNSConfig{
			Server:       "http://localhost:8081",
			APIKey:       "foo",
			DomainFilter: endpoint.NewDomainFilter([]string{""}),
			TLSConfig: TLSConfig{
				TLSEnabled:            true,
				CAFilePath:            "../internal/testresources/ca.pem",
				ClientCertFilePath:    "../internal/testresources/client-cert.pem",
				ClientCertKeyFilePath: "../internal/testresources/client-cert-key.pem",
			},
		})
	assert.Nil(suite.T(), err, "Enabled TLS Config with all flags should raise no error")
}

func (suite *NewPDNSProviderTestSuite) TestPDNSRRSetToEndpoints() {
	// Function definition: convertRRSetToEndpoints(rr pgo.RrSet) (endpoints []*endpoint.Endpoint, _ error)

	// Create a new provider to run tests against
	p := &PDNSProvider{
		client: &PDNSAPIClientStub{},
	}

	/* given an RRSet with three records, we test:
	   - We correctly create corresponding endpoints
	*/
	eps, err := p.convertRRSetToEndpoints(RRSetMultipleRecords)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), endpointsMultipleRecords, eps)

	/* Given an RRSet with two records, one of which is disabled, we test:
	   - We can correctly convert the RRSet into a list of valid endpoints
	   - We correctly discard/ignore the disabled record.
	*/
	eps, err = p.convertRRSetToEndpoints(RRSetDisabledRecord)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), endpointsDisabledRecord, eps)

}

func (suite *NewPDNSProviderTestSuite) TestPDNSRecords() {
	// Function definition: Records() (endpoints []*endpoint.Endpoint, _ error)

	// Create a new provider to run tests against
	p := &PDNSProvider{
		client: &PDNSAPIClientStub{},
	}

	ctx := context.Background()

	/* We test that endpoints are returned correctly for a Zone when Records() is called
	 */
	eps, err := p.Records(ctx)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), endpointsMixedRecords, eps)

	// Test failures are handled correctly
	// Create a new provider to run tests against
	p = &PDNSProvider{
		client: &PDNSAPIClientStubListZoneFailure{},
	}
	_, err = p.Records(ctx)
	assert.NotNil(suite.T(), err)

	p = &PDNSProvider{
		client: &PDNSAPIClientStubListZonesFailure{},
	}
	_, err = p.Records(ctx)
	assert.NotNil(suite.T(), err)

}

func (suite *NewPDNSProviderTestSuite) TestPDNSConvertEndpointsToZones() {
	// Function definition: ConvertEndpointsToZones(endpoints []*endpoint.Endpoint, changetype pdnsChangeType) (zonelist []pgo.Zone, _ error)

	// Create a new provider to run tests against
	p := &PDNSProvider{
		client: &PDNSAPIClientStubEmptyZones{},
	}

	// Check inserting endpoints from a single zone
	zlist, err := p.ConvertEndpointsToZones(endpointsSimpleRecord, PdnsReplace)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), []pgo.Zone{ZoneEmptyToSimplePatch}, zlist)

	// Check deleting endpoints from a single zone
	zlist, err = p.ConvertEndpointsToZones(endpointsSimpleRecord, PdnsDelete)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), []pgo.Zone{ZoneEmptyToSimpleDelete}, zlist)

	// Check endpoints from multiple zones #1
	zlist, err = p.ConvertEndpointsToZones(endpointsMultipleZones, PdnsReplace)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), []pgo.Zone{ZoneEmptyToSimplePatch, ZoneEmptyToSimplePatch2}, zlist)

	// Check endpoints from multiple zones #2
	zlist, err = p.ConvertEndpointsToZones(endpointsMultipleZones2, PdnsReplace)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), []pgo.Zone{ZoneEmptyToSimplePatch, ZoneEmptyToSimplePatch3}, zlist)

	// Check endpoints from multiple zones where some endpoints which don't exist
	zlist, err = p.ConvertEndpointsToZones(endpointsMultipleZonesWithNoExist, PdnsReplace)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), []pgo.Zone{ZoneEmptyToSimplePatch}, zlist)

	// Check endpoints from a zone that does not exist
	zlist, err = p.ConvertEndpointsToZones(endpointsNonexistantZone, PdnsReplace)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), []pgo.Zone{}, zlist)

	// Check endpoints that match multiple zones (one longer than other), is assigned to the right zone
	zlist, err = p.ConvertEndpointsToZones(endpointsLongRecord, PdnsReplace)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), []pgo.Zone{ZoneEmptyToLongPatch}, zlist)

	// Check endpoints of type CNAME always have their target records end with a dot.
	zlist, err = p.ConvertEndpointsToZones(endpointsMixedRecords, PdnsReplace)
	assert.Nil(suite.T(), err)

	for _, z := range zlist {
		for _, rs := range z.Rrsets {
			if "CNAME" == rs.Type_ {
				for _, r := range rs.Records {
					assert.Equal(suite.T(), uint8(0x2e), r.Content[len(r.Content)-1])
				}
			}
		}
	}
}

func (suite *NewPDNSProviderTestSuite) TestPDNSConvertEndpointsToZonesPartitionZones() {
	// Test DomainFilters
	p := &PDNSProvider{
		client: &PDNSAPIClientStubPartitionZones{},
	}

	// Check inserting endpoints from a single zone which is specified in DomainFilter
	zlist, err := p.ConvertEndpointsToZones(endpointsSimpleRecord, PdnsReplace)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), []pgo.Zone{ZoneEmptyToSimplePatch}, zlist)

	// Check deleting endpoints from a single zone which is specified in DomainFilter
	zlist, err = p.ConvertEndpointsToZones(endpointsSimpleRecord, PdnsDelete)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), []pgo.Zone{ZoneEmptyToSimpleDelete}, zlist)

	// Check endpoints from multiple zones # which one is specified in DomainFilter and one is not
	zlist, err = p.ConvertEndpointsToZones(endpointsMultipleZones, PdnsReplace)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), []pgo.Zone{ZoneEmptyToSimplePatch}, zlist)

	// Check endpoints from multiple zones where some endpoints which don't exist and one that does
	// and is part of DomainFilter
	zlist, err = p.ConvertEndpointsToZones(endpointsMultipleZonesWithNoExist, PdnsReplace)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), []pgo.Zone{ZoneEmptyToSimplePatch}, zlist)

	// Check endpoints from a zone that does not exist
	zlist, err = p.ConvertEndpointsToZones(endpointsNonexistantZone, PdnsReplace)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), []pgo.Zone{}, zlist)

	// Check endpoints that match multiple zones (one longer than other), is assigned to the right zone when the longer
	// zone is not part of the DomainFilter
	zlist, err = p.ConvertEndpointsToZones(endpointsMultipleZonesWithLongRecordNotInDomainFilter, PdnsReplace)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), []pgo.Zone{ZoneEmptyToSimplePatchLongRecordIgnoredInDomainFilter}, zlist)

	// Check endpoints that match multiple zones (one longer than other and one is very similar)
	// is assigned to the right zone when the similar zone is not part of the DomainFilter
	zlist, err = p.ConvertEndpointsToZones(endpointsMultipleZonesWithSimilarRecordNotInDomainFilter, PdnsReplace)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), []pgo.Zone{ZoneEmptyToSimplePatch}, zlist)
}

func (suite *NewPDNSProviderTestSuite) TestPDNSmutateRecords() {
	// Function definition: mutateRecords(endpoints []*endpoint.Endpoint, changetype pdnsChangeType) error

	// Create a new provider to run tests against
	c := &PDNSAPIClientStubEmptyZones{}
	p := &PDNSProvider{
		client: c,
	}

	// Check inserting endpoints from a single zone
	err := p.mutateRecords(endpointsSimpleRecord, pdnsChangeType("REPLACE"))
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), []pgo.Zone{ZoneEmptyToSimplePatch}, c.patchedZones)

	// Reset the "patchedZones"
	c.patchedZones = []pgo.Zone{}

	// Check deleting endpoints from a single zone
	err = p.mutateRecords(endpointsSimpleRecord, pdnsChangeType("DELETE"))
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), []pgo.Zone{ZoneEmptyToSimpleDelete}, c.patchedZones)

	// Check we fail correctly when patching fails for whatever reason
	p = &PDNSProvider{
		client: &PDNSAPIClientStubPatchZoneFailure{},
	}
	// Check inserting endpoints from a single zone
	err = p.mutateRecords(endpointsSimpleRecord, pdnsChangeType("REPLACE"))
	assert.NotNil(suite.T(), err)

}

func (suite *NewPDNSProviderTestSuite) TestPDNSClientPartitionZones() {
	zoneList := []pgo.Zone{
		ZoneEmpty,
		ZoneEmpty2,
	}

	partitionResultFilteredEmptyFilter := []pgo.Zone{
		ZoneEmpty,
		ZoneEmpty2,
	}

	partitionResultResidualEmptyFilter := ([]pgo.Zone)(nil)

	partitionResultFilteredSingleFilter := []pgo.Zone{
		ZoneEmpty,
	}

	partitionResultResidualSingleFilter := []pgo.Zone{
		ZoneEmpty2,
	}

	partitionResultFilteredMultipleFilter := []pgo.Zone{
		ZoneEmpty,
	}

	partitionResultResidualMultipleFilter := []pgo.Zone{
		ZoneEmpty2,
	}

	// Check filtered, residual zones when no domain filter specified
	filteredZones, residualZones := DomainFilterEmptyClient.PartitionZones(zoneList)
	assert.Equal(suite.T(), partitionResultFilteredEmptyFilter, filteredZones)
	assert.Equal(suite.T(), partitionResultResidualEmptyFilter, residualZones)

	// Check filtered, residual zones when a single domain filter specified
	filteredZones, residualZones = DomainFilterSingleClient.PartitionZones(zoneList)
	assert.Equal(suite.T(), partitionResultFilteredSingleFilter, filteredZones)
	assert.Equal(suite.T(), partitionResultResidualSingleFilter, residualZones)

	// Check filtered, residual zones when a multiple domain filter specified
	filteredZones, residualZones = DomainFilterMultipleClient.PartitionZones(zoneList)
	assert.Equal(suite.T(), partitionResultFilteredMultipleFilter, filteredZones)
	assert.Equal(suite.T(), partitionResultResidualMultipleFilter, residualZones)
}

func TestNewPDNSProviderTestSuite(t *testing.T) {
	suite.Run(t, new(NewPDNSProviderTestSuite))
}
