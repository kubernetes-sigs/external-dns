/*
Copyright 2022 The Kubernetes Authors.
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
package desec

import (
	"context"
	"testing"

	nc "github.com/nrdcg/desec"
	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

func TestDesecProvider(t *testing.T) {
	t.Run("ExtractSubname", testExtractSubname)
	t.Run("EndpointZoneName", testEndpointZoneName)
	t.Run("ConvertToDesecRecord", testConvertToDesecRecord)
	t.Run("NewdesecProvider", testNewDesecProvider)
	t.Run("ApplyChanges", testApplyChanges)
	t.Run("Records", testRecords)

}

func testExtractSubname(t *testing.T) {
	// given
	fqdn1 := "foo.bar.org"
	fqdn2 := "foo.foo-bar.bar.org"

	// then
	subName1 := extractSubname(fqdn1)
	subName2 := extractSubname(fqdn2)

	// assert
	assert.Equal(t, "foo", subName1)
	assert.Equal(t, "foo.foo-bar", subName2)
}

func testEndpointZoneName(t *testing.T) {
	zoneList := []string{"bar.org", "baz.org"}

	// in zone list
	ep1 := endpoint.Endpoint{
		DNSName:    "foo.bar.org",
		Targets:    endpoint.Targets{"5.5.5.5"},
		RecordType: endpoint.RecordTypeA,
	}

	// not in zone list
	ep2 := endpoint.Endpoint{
		DNSName:    "foo.foo.org",
		Targets:    endpoint.Targets{"5.5.5.5"},
		RecordType: endpoint.RecordTypeA,
	}

	// matches zone exactly
	ep3 := endpoint.Endpoint{
		DNSName:    "baz.org",
		Targets:    endpoint.Targets{"5.5.5.5"},
		RecordType: endpoint.RecordTypeA,
	}

	assert.Equal(t, endpointZoneName(&ep1, zoneList), "bar.org")
	assert.Equal(t, endpointZoneName(&ep2, zoneList), "")
	assert.Equal(t, endpointZoneName(&ep3, zoneList), "baz.org")
}

func testConvertToDesecRecord(t *testing.T) {

	// given

	// in zone list
	ep1 := endpoint.Endpoint{
		DNSName:    "foo.bar.org",
		Targets:    endpoint.Targets{"5.5.5.5"},
		RecordType: endpoint.RecordTypeA,
	}

	// matches zone exactly
	ep2 := endpoint.Endpoint{
		DNSName:    "bar.org",
		Targets:    endpoint.Targets{"5.5.5.5"},
		RecordType: endpoint.RecordTypeA,
	}

	// txt type
	ep3 := endpoint.Endpoint{
		DNSName:    "foo.bar.org",
		Targets:    endpoint.Targets{"\"heritage=external-dns,external-dns/owner=default,external-dns/resource=service/default/nginx\""},
		RecordType: endpoint.RecordTypeTXT,
	}

	// cname ends with .
	ep4 := endpoint.Endpoint{
		DNSName:    "foofoo.bar.org",
		Targets:    endpoint.Targets{"foo.bar.org."},
		RecordType: endpoint.RecordTypeCNAME,
	}

	epList := []*endpoint.Endpoint{&ep1, &ep2, &ep3, &ep4}

	nc1 := nc.RRSet{
		Name:    "foo.bar.org.",
		Domain:  "bar.org",
		SubName: "foo",
		Type:    "A",
		Records: []string{"5.5.5.5"},
		TTL:     3600,
	}

	nc2 := nc.RRSet{
		Name:    "bar.org.",
		Domain:  "bar.org",
		SubName: "",
		Type:    "A",
		Records: []string{"5.5.5.5"},
		TTL:     3600,
	}

	nc3 := nc.RRSet{
		Name:    "foo.bar.org.",
		Domain:  "bar.org",
		SubName: "foo",
		Type:    "TXT",
		Records: []string{"\"heritage=external-dns,external-dns/owner=default,external-dns/resource=service/default/nginx\""},
		TTL:     3600,
	}

	nc4 := nc.RRSet{
		Name:    "foofoo.bar.org.",
		Domain:  "bar.org",
		SubName: "foofoo",
		Type:    "CNAME",
		Records: []string{"foo.bar.org."},
		TTL:     3600,
	}

	ncRecordList := []nc.RRSet{nc1, nc2, nc3, nc4}

	// then
	convertedDesecRecords := convertToDesecRecord(epList, "bar.org")

	// assert
	assert.Equal(t, &ncRecordList, convertedDesecRecords)
}

func testNewDesecProvider(t *testing.T) {
	p, err := NewDesecProvider(endpoint.NewDomainFilter([]string{"example.com"}), "KEY", true)
	assert.NotNil(t, p.client)
	assert.NoError(t, err)

	_, err = NewDesecProvider(endpoint.NewDomainFilter([]string{"example.com"}), "", true)
	assert.Error(t, err)

	_, err = NewDesecProvider(endpoint.NewDomainFilter([]string{}), "KEY", true)
	assert.Error(t, err)

}

func testApplyChanges(t *testing.T) {
	p, _ := NewDesecProvider(endpoint.NewDomainFilter([]string{"example.com"}), "KEY", true)
	changes1 := &plan.Changes{
		Create:    []*endpoint.Endpoint{},
		Delete:    []*endpoint.Endpoint{},
		UpdateNew: []*endpoint.Endpoint{},
		UpdateOld: []*endpoint.Endpoint{},
	}

	// No Changes
	err := p.ApplyChanges(context.TODO(), changes1)
	assert.NoError(t, err)

	// Changes
	changes2 := &plan.Changes{
		Create: []*endpoint.Endpoint{
			{
				DNSName:    "api.example.com",
				RecordType: "A",
			},
			{
				DNSName:    "api.baz.com",
				RecordType: "TXT",
			}},
		Delete: []*endpoint.Endpoint{
			{
				DNSName:    "api.example.com",
				RecordType: "A",
			},
			{
				DNSName:    "api.baz.com",
				RecordType: "TXT",
			}},
		UpdateNew: []*endpoint.Endpoint{
			{
				DNSName:    "api.example.com",
				RecordType: "A",
			},
			{
				DNSName:    "api.baz.com",
				RecordType: "TXT",
			}},
		UpdateOld: []*endpoint.Endpoint{
			{
				DNSName:    "api.example.com",
				RecordType: "A",
			},
			{
				DNSName:    "api.baz.com",
				RecordType: "TXT",
			}},
	}

	err = p.ApplyChanges(context.TODO(), changes2)
	assert.NoError(t, err)

}

func testRecords(t *testing.T) {
	p, _ := NewDesecProvider(endpoint.NewDomainFilter([]string{"example.com"}), "KEY", true)
	ep, err := p.Records(context.TODO())
	assert.Equal(t, []*endpoint.Endpoint{}, ep)
	assert.NoError(t, err)
}
