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

package designate

import (
	"context"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"sort"
	"sync/atomic"
	"testing"

	"github.com/gophercloud/gophercloud/openstack/dns/v2/recordsets"
	"github.com/gophercloud/gophercloud/openstack/dns/v2/zones"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

var lastGeneratedDesignateID int32

func generateDesignateID() string {
	return fmt.Sprintf("id-%d", atomic.AddInt32(&lastGeneratedDesignateID, 1))
}

type fakeDesignateClient struct {
	managedZones map[string]*struct {
		zone       *zones.Zone
		recordSets map[string]*recordsets.RecordSet
	}
}

func (c fakeDesignateClient) AddZone(zone zones.Zone) string {
	if zone.ID == "" {
		zone.ID = zone.Name
	}
	c.managedZones[zone.ID] = &struct {
		zone       *zones.Zone
		recordSets map[string]*recordsets.RecordSet
	}{
		zone:       &zone,
		recordSets: make(map[string]*recordsets.RecordSet),
	}
	return zone.ID
}

func (c fakeDesignateClient) ForEachZone(handler func(zone *zones.Zone) error) error {
	for _, zone := range c.managedZones {
		if err := handler(zone.zone); err != nil {
			return err
		}
	}
	return nil
}

func (c fakeDesignateClient) ForEachRecordSet(zoneID string, handler func(recordSet *recordsets.RecordSet) error) error {
	zone := c.managedZones[zoneID]
	if zone == nil {
		return fmt.Errorf("unknown zone %s", zoneID)
	}
	for _, recordSet := range zone.recordSets {
		if err := handler(recordSet); err != nil {
			return err
		}
	}
	return nil
}

func (c fakeDesignateClient) CreateRecordSet(zoneID string, opts recordsets.CreateOpts) (string, error) {
	zone := c.managedZones[zoneID]
	if zone == nil {
		return "", fmt.Errorf("unknown zone %s", zoneID)
	}
	rs := &recordsets.RecordSet{
		ID:          generateDesignateID(),
		ZoneID:      zoneID,
		Name:        opts.Name,
		Description: opts.Description,
		Records:     opts.Records,
		TTL:         opts.TTL,
		Type:        opts.Type,
	}
	zone.recordSets[rs.ID] = rs
	return rs.ID, nil
}

func (c fakeDesignateClient) UpdateRecordSet(zoneID, recordSetID string, opts recordsets.UpdateOpts) error {
	zone := c.managedZones[zoneID]
	if zone == nil {
		return fmt.Errorf("unknown zone %s", zoneID)
	}
	rs := zone.recordSets[recordSetID]
	if rs == nil {
		return fmt.Errorf("unknown record-set %s", recordSetID)
	}
	if opts.Description != nil {
		rs.Description = *opts.Description
	}
	rs.TTL = *opts.TTL

	rs.Records = opts.Records
	return nil
}

func (c fakeDesignateClient) DeleteRecordSet(zoneID, recordSetID string) error {
	zone := c.managedZones[zoneID]
	if zone == nil {
		return fmt.Errorf("unknown zone %s", zoneID)
	}
	delete(zone.recordSets, recordSetID)
	return nil
}

func (c fakeDesignateClient) ToProvider() provider.Provider {
	return &designateProvider{client: c}
}

func newFakeDesignateClient() *fakeDesignateClient {
	return &fakeDesignateClient{
		make(map[string]*struct {
			zone       *zones.Zone
			recordSets map[string]*recordsets.RecordSet
		}),
	}
}

func TestNewDesignateProvider(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{
		  "token": {
		    "catalog": [
		      {
		        "id": "9615c2dfac3b4b19935226d4c9d4afce",
		        "name": "designate",
		        "type": "dns",
		        "endpoints": [
		          {
		            "id": "3d3cc3a273b54d0490ac43d6572e4c48",
		            "region": "RegionOne",
		            "region_id": "RegionOne",
		            "interface": "public",
		            "url": "https://example.com:9001"
		          }
		        ]
		      }
		    ]
		  }
		}`))
	}))
	defer ts.Close()

	block := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: ts.Certificate().Raw,
	}
	tmpfile, err := ioutil.TempFile("", "os-test.crt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	if err := pem.Encode(tmpfile, block); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	os.Setenv("OS_AUTH_URL", ts.URL+"/v3")
	os.Setenv("OS_USERNAME", "username")
	os.Setenv("OS_PASSWORD", "password")
	os.Setenv("OS_USER_DOMAIN_NAME", "Default")
	os.Setenv("OPENSTACK_CA_FILE", tmpfile.Name())

	if _, err := NewDesignateProvider(endpoint.DomainFilter{}, true); err != nil {
		t.Fatalf("Failed to initialize Designate provider: %s", err)
	}
}

func TestDesignateRecords(t *testing.T) {
	client := newFakeDesignateClient()

	zone1ID := client.AddZone(zones.Zone{
		Name:   "example.com.",
		Type:   "PRIMARY",
		Status: "ACTIVE",
	})
	rs11ID, _ := client.CreateRecordSet(zone1ID, recordsets.CreateOpts{
		Name:    "www.example.com.",
		Type:    endpoint.RecordTypeA,
		Records: []string{"10.1.1.1"},
	})
	rs12ID, _ := client.CreateRecordSet(zone1ID, recordsets.CreateOpts{
		Name:    "www.example.com.",
		Type:    endpoint.RecordTypeTXT,
		Records: []string{"text1"},
	})
	client.CreateRecordSet(zone1ID, recordsets.CreateOpts{
		Name:    "xxx.example.com.",
		Type:    "SRV",
		Records: []string{"http://test.com:1234"},
	})
	rs14ID, _ := client.CreateRecordSet(zone1ID, recordsets.CreateOpts{
		Name:    "ftp.example.com.",
		Type:    endpoint.RecordTypeA,
		Records: []string{"10.1.1.2"},
	})

	zone2ID := client.AddZone(zones.Zone{
		Name:   "test.net.",
		Type:   "PRIMARY",
		Status: "ACTIVE",
	})
	rs21ID, _ := client.CreateRecordSet(zone2ID, recordsets.CreateOpts{
		Name:    "srv.test.net.",
		Type:    endpoint.RecordTypeA,
		Records: []string{"10.2.1.1", "10.2.1.2"},
	})
	rs22ID, _ := client.CreateRecordSet(zone2ID, recordsets.CreateOpts{
		Name:    "db.test.net.",
		Type:    endpoint.RecordTypeCNAME,
		Records: []string{"sql.test.net."},
	})
	expected := []*endpoint.Endpoint{
		{
			DNSName:    "www.example.com",
			RecordType: endpoint.RecordTypeA,
			Targets:    endpoint.Targets{"10.1.1.1"},
			Labels: map[string]string{
				designateRecordSetID:     rs11ID,
				designateZoneID:          zone1ID,
				designateOriginalRecords: "10.1.1.1",
			},
		},
		{
			DNSName:    "www.example.com",
			RecordType: endpoint.RecordTypeTXT,
			Targets:    endpoint.Targets{"text1"},
			Labels: map[string]string{
				designateRecordSetID:     rs12ID,
				designateZoneID:          zone1ID,
				designateOriginalRecords: "text1",
			},
		},
		{
			DNSName:    "ftp.example.com",
			RecordType: endpoint.RecordTypeA,
			Targets:    endpoint.Targets{"10.1.1.2"},
			Labels: map[string]string{
				designateRecordSetID:     rs14ID,
				designateZoneID:          zone1ID,
				designateOriginalRecords: "10.1.1.2",
			},
		},
		{
			DNSName:    "srv.test.net",
			RecordType: endpoint.RecordTypeA,
			Targets:    endpoint.Targets{"10.2.1.1"},
			Labels: map[string]string{
				designateRecordSetID:     rs21ID,
				designateZoneID:          zone2ID,
				designateOriginalRecords: "10.2.1.1\00010.2.1.2",
			},
		},
		{
			DNSName:    "srv.test.net",
			RecordType: endpoint.RecordTypeA,
			Targets:    endpoint.Targets{"10.2.1.2"},
			Labels: map[string]string{
				designateRecordSetID:     rs21ID,
				designateZoneID:          zone2ID,
				designateOriginalRecords: "10.2.1.1\00010.2.1.2",
			},
		},
		{
			DNSName:    "db.test.net",
			RecordType: endpoint.RecordTypeCNAME,
			Targets:    endpoint.Targets{"sql.test.net"},
			Labels: map[string]string{
				designateRecordSetID:     rs22ID,
				designateZoneID:          zone2ID,
				designateOriginalRecords: "sql.test.net.",
			},
		},
	}

	endpoints, err := client.ToProvider().Records(context.Background())
	if err != nil {
		t.Fatal(err)
	}
out:
	for _, ep := range endpoints {
		for i, ex := range expected {
			if reflect.DeepEqual(ep, ex) {
				expected = append(expected[:i], expected[i+1:]...)
				continue out
			}
		}
		t.Errorf("unexpected endpoint %s/%s -> %s", ep.DNSName, ep.RecordType, ep.Targets)
	}
	if len(expected) != 0 {
		t.Errorf("not all expected endpoints were returned. Remained: %v", expected)
	}
}

func TestDesignateCreateRecords(t *testing.T) {
	client := newFakeDesignateClient()
	testDesignateCreateRecords(t, client)
}

func testDesignateCreateRecords(t *testing.T, client *fakeDesignateClient) []*recordsets.RecordSet {
	for i, zoneName := range []string{"example.com.", "test.net."} {
		client.AddZone(zones.Zone{
			ID:     fmt.Sprintf("zone-%d", i+1),
			Name:   zoneName,
			Type:   "PRIMARY",
			Status: "ACTIVE",
		})
	}
	endpoints := []*endpoint.Endpoint{
		{
			DNSName:    "www.example.com",
			RecordType: endpoint.RecordTypeA,
			Targets:    endpoint.Targets{"10.1.1.1"},
			Labels:     map[string]string{},
		},
		{
			DNSName:    "www.example.com",
			RecordType: endpoint.RecordTypeTXT,
			Targets:    endpoint.Targets{"text1"},
			Labels:     map[string]string{},
		},
		{
			DNSName:    "ftp.example.com",
			RecordType: endpoint.RecordTypeA,
			Targets:    endpoint.Targets{"10.1.1.2"},
			Labels:     map[string]string{},
		},
		{
			DNSName:    "srv.test.net",
			RecordType: endpoint.RecordTypeA,
			Targets:    endpoint.Targets{"10.2.1.1"},
			Labels:     map[string]string{},
		},
		{
			DNSName:    "srv.test.net",
			RecordType: endpoint.RecordTypeA,
			Targets:    endpoint.Targets{"10.2.1.2"},
			Labels:     map[string]string{},
		},
		{
			DNSName:    "db.test.net",
			RecordType: endpoint.RecordTypeCNAME,
			Targets:    endpoint.Targets{"sql.test.net"},
			Labels:     map[string]string{},
		},
	}
	expected := []*recordsets.RecordSet{
		{
			Name:    "www.example.com.",
			Type:    endpoint.RecordTypeA,
			Records: []string{"10.1.1.1"},
			ZoneID:  "zone-1",
		},
		{
			Name:    "www.example.com.",
			Type:    endpoint.RecordTypeTXT,
			Records: []string{"text1"},
			ZoneID:  "zone-1",
		},
		{
			Name:    "ftp.example.com.",
			Type:    endpoint.RecordTypeA,
			Records: []string{"10.1.1.2"},
			ZoneID:  "zone-1",
		},
		{
			Name:    "srv.test.net.",
			Type:    endpoint.RecordTypeA,
			Records: []string{"10.2.1.1", "10.2.1.2"},
			ZoneID:  "zone-2",
		},
		{
			Name:    "db.test.net.",
			Type:    endpoint.RecordTypeCNAME,
			Records: []string{"sql.test.net."},
			ZoneID:  "zone-2",
		},
	}
	expectedCopy := make([]*recordsets.RecordSet, len(expected))
	copy(expectedCopy, expected)

	err := client.ToProvider().ApplyChanges(context.Background(), &plan.Changes{Create: endpoints})
	if err != nil {
		t.Fatal(err)
	}

	client.ForEachZone(func(zone *zones.Zone) error {
		client.ForEachRecordSet(zone.ID, func(recordSet *recordsets.RecordSet) error {
			id := recordSet.ID
			recordSet.ID = ""
			for i, ex := range expected {
				sort.Strings(recordSet.Records)
				if reflect.DeepEqual(ex, recordSet) {
					ex.ID = id
					recordSet.ID = id
					expected = append(expected[:i], expected[i+1:]...)
					return nil
				}
			}
			t.Errorf("unexpected record-set %s/%s -> %v", recordSet.Name, recordSet.Type, recordSet.Records)
			return nil
		})
		return nil
	})

	if len(expected) != 0 {
		t.Errorf("not all expected record-sets were created. Remained: %v", expected)
	}
	return expectedCopy
}

func TestDesignateUpdateRecords(t *testing.T) {
	client := newFakeDesignateClient()
	testDesignateUpdateRecords(t, client)
}

func testDesignateUpdateRecords(t *testing.T, client *fakeDesignateClient) []*recordsets.RecordSet {
	expected := testDesignateCreateRecords(t, client)

	updatesOld := []*endpoint.Endpoint{
		{
			DNSName:    "ftp.example.com",
			RecordType: endpoint.RecordTypeA,
			Targets:    endpoint.Targets{"10.1.1.2"},
			Labels: map[string]string{
				designateZoneID:          "zone-1",
				designateRecordSetID:     expected[2].ID,
				designateOriginalRecords: "10.1.1.2",
			},
		},
		{
			DNSName:    "srv.test.net.",
			RecordType: endpoint.RecordTypeA,
			Targets:    endpoint.Targets{"10.2.1.2"},
			Labels: map[string]string{
				designateZoneID:          "zone-2",
				designateRecordSetID:     expected[3].ID,
				designateOriginalRecords: "10.2.1.1\00010.2.1.2",
			},
		},
	}
	updatesNew := []*endpoint.Endpoint{
		{
			DNSName:    "ftp.example.com",
			RecordType: endpoint.RecordTypeA,
			Targets:    endpoint.Targets{"10.3.3.1"},
			Labels: map[string]string{
				designateZoneID:          "zone-1",
				designateRecordSetID:     expected[2].ID,
				designateOriginalRecords: "10.1.1.2",
			},
		},
		{
			DNSName:    "srv.test.net.",
			RecordType: endpoint.RecordTypeA,
			Targets:    endpoint.Targets{"10.3.3.2"},
			Labels: map[string]string{
				designateZoneID:          "zone-2",
				designateRecordSetID:     expected[3].ID,
				designateOriginalRecords: "10.2.1.1\00010.2.1.2",
			},
		},
	}
	expectedCopy := make([]*recordsets.RecordSet, len(expected))
	copy(expectedCopy, expected)

	expected[2].Records = []string{"10.3.3.1"}
	expected[3].Records = []string{"10.2.1.1", "10.3.3.2"}

	err := client.ToProvider().ApplyChanges(context.Background(), &plan.Changes{UpdateOld: updatesOld, UpdateNew: updatesNew})
	if err != nil {
		t.Fatal(err)
	}

	client.ForEachZone(func(zone *zones.Zone) error {
		client.ForEachRecordSet(zone.ID, func(recordSet *recordsets.RecordSet) error {
			for i, ex := range expected {
				sort.Strings(recordSet.Records)
				if reflect.DeepEqual(ex, recordSet) {
					expected = append(expected[:i], expected[i+1:]...)
					return nil
				}
			}
			t.Errorf("unexpected record-set %s/%s -> %v", recordSet.Name, recordSet.Type, recordSet.Records)
			return nil
		})
		return nil
	})

	if len(expected) != 0 {
		t.Errorf("not all expected record-sets were updated. Remained: %v", expected)
	}
	return expectedCopy
}

func TestDesignateDeleteRecords(t *testing.T) {
	client := newFakeDesignateClient()
	testDesignateDeleteRecords(t, client)
}

func testDesignateDeleteRecords(t *testing.T, client *fakeDesignateClient) {
	expected := testDesignateUpdateRecords(t, client)
	deletes := []*endpoint.Endpoint{
		{
			DNSName:    "www.example.com.",
			RecordType: endpoint.RecordTypeA,
			Targets:    endpoint.Targets{"10.1.1.1"},
			Labels: map[string]string{
				designateZoneID:          "zone-1",
				designateRecordSetID:     expected[0].ID,
				designateOriginalRecords: "10.1.1.1",
			},
		},
		{
			DNSName:    "srv.test.net.",
			RecordType: endpoint.RecordTypeA,
			Targets:    endpoint.Targets{"10.2.1.1"},
			Labels: map[string]string{
				designateZoneID:          "zone-2",
				designateRecordSetID:     expected[3].ID,
				designateOriginalRecords: "10.2.1.1\00010.3.3.2",
			},
		},
	}
	expected[3].Records = []string{"10.3.3.2"}
	expected = expected[1:]

	err := client.ToProvider().ApplyChanges(context.Background(), &plan.Changes{Delete: deletes})
	if err != nil {
		t.Fatal(err)
	}

	client.ForEachZone(func(zone *zones.Zone) error {
		client.ForEachRecordSet(zone.ID, func(recordSet *recordsets.RecordSet) error {
			for i, ex := range expected {
				sort.Strings(recordSet.Records)
				if reflect.DeepEqual(ex, recordSet) {
					expected = append(expected[:i], expected[i+1:]...)
					return nil
				}
			}
			t.Errorf("unexpected record-set %s/%s -> %v", recordSet.Name, recordSet.Type, recordSet.Records)
			return nil
		})
		return nil
	})

	if len(expected) != 0 {
		t.Errorf("not all expected record-sets were deleted. Remained: %v", expected)
	}
}
