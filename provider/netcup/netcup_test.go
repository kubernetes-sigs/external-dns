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
package netcup

import (
	"context"
	"testing"

	nc "github.com/aellwein/netcup-dns-api/pkg/v1"

	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

func TestNetcupProvider(t *testing.T) {
	t.Run("EndpointZoneName", testEndpointZoneName)
	t.Run("GetIDforRecord", testGetIDforRecord)
	t.Run("ConvertToNetcupRecord", testConvertToNetcupRecord)
	t.Run("NewNetcupProvider", testNewNetcupProvider)
	t.Run("ApplyChanges", testApplyChanges)
	t.Run("Records", testRecords)
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

func testGetIDforRecord(t *testing.T) {

	recordName := "foo.example.com"
	target1 := "heritage=external-dns,external-dns/owner=default,external-dns/resource=service/default/nginx"
	target2 := "5.5.5.5"
	recordType := "TXT"

	nc1 := nc.DnsRecord{
		Hostname:     "foo.example.com",
		Type:         "TXT",
		Destination:  "heritage=external-dns,external-dns/owner=default,external-dns/resource=service/default/nginx",
		Id:           "10",
		DeleteRecord: false,
	}
	nc2 := nc.DnsRecord{
		Hostname:     "foo.foo.org",
		Type:         "A",
		Destination:  "5.5.5.5",
		Id:           "10",
		DeleteRecord: false,
	}

	nc3 := nc.DnsRecord{
		Id:           "",
		Hostname:     "baz.org",
		Type:         "A",
		Destination:  "5.5.5.5",
		DeleteRecord: false,
	}

	ncRecordList := []nc.DnsRecord{nc1, nc2, nc3}

	assert.Equal(t, "10", getIDforRecord(recordName, target1, recordType, &ncRecordList))
	assert.Equal(t, "", getIDforRecord(recordName, target2, recordType, &ncRecordList))

}

func testConvertToNetcupRecord(t *testing.T) {
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
		DNSName:    "bar.org",
		Targets:    endpoint.Targets{"5.5.5.5"},
		RecordType: endpoint.RecordTypeA,
	}

	ep4 := endpoint.Endpoint{
		DNSName:    "foo.baz.org",
		Targets:    endpoint.Targets{"\"heritage=external-dns,external-dns/owner=default,external-dns/resource=service/default/nginx\""},
		RecordType: endpoint.RecordTypeTXT,
	}

	epList := []*endpoint.Endpoint{&ep1, &ep2, &ep3, &ep4}

	nc1 := nc.DnsRecord{
		Hostname:     "foo",
		Type:         "A",
		Destination:  "5.5.5.5",
		Id:           "10",
		DeleteRecord: false,
	}
	nc2 := nc.DnsRecord{
		Hostname:     "foo.foo.org",
		Type:         "A",
		Destination:  "5.5.5.5",
		Id:           "15",
		DeleteRecord: false,
	}

	nc3 := nc.DnsRecord{
		Id:           "",
		Hostname:     "@",
		Type:         "A",
		Destination:  "5.5.5.5",
		DeleteRecord: false,
	}

	nc4 := nc.DnsRecord{
		Id:           "",
		Hostname:     "foo.baz.org",
		Type:         "TXT",
		Destination:  "heritage=external-dns,external-dns/owner=default,external-dns/resource=service/default/nginx",
		DeleteRecord: false,
	}

	ncRecordList := []nc.DnsRecord{nc1, nc2, nc3, nc4}

	// No deletion
	assert.Equal(t, convertToNetcupRecord(&ncRecordList, epList, "bar.org", false), &ncRecordList)
	// Deletion active

	nc1.DeleteRecord = true
	nc2.DeleteRecord = true
	nc3.DeleteRecord = true
	nc4.DeleteRecord = true
	ncRecordList2 := []nc.DnsRecord{nc1, nc2, nc3, nc4}
	assert.Equal(t, convertToNetcupRecord(&ncRecordList2, epList, "bar.org", true), &ncRecordList2)

}

func testNewNetcupProvider(t *testing.T) {
	p, err := NewNetcupProvider(endpoint.NewDomainFilter([]string{"example.com"}), 10, "KEY", "PASSWORD", true)
	assert.NotNil(t, p.client)
	assert.NoError(t, err)

	_, err = NewNetcupProvider(endpoint.NewDomainFilter([]string{"example.com"}), 0, "KEY", "PASSWORD", true)
	assert.Error(t, err)

	_, err = NewNetcupProvider(endpoint.NewDomainFilter([]string{"example.com"}), 10, "", "PASSWORD", true)
	assert.Error(t, err)

	_, err = NewNetcupProvider(endpoint.NewDomainFilter([]string{"example.com"}), 10, "KEY", "", true)
	assert.Error(t, err)

	_, err = NewNetcupProvider(endpoint.NewDomainFilter([]string{}), 10, "KEY", "PASSWORD", true)
	assert.Error(t, err)

}

func testApplyChanges(t *testing.T) {
	p, _ := NewNetcupProvider(endpoint.NewDomainFilter([]string{"example.com"}), 10, "KEY", "PASSWORD", true)
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
	p, _ := NewNetcupProvider(endpoint.NewDomainFilter([]string{"example.com"}), 10, "KEY", "PASSWORD", true)
	ep, err := p.Records(context.TODO())
	assert.Equal(t, []*endpoint.Endpoint{}, ep)
	assert.NoError(t, err)
}
