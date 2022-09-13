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
	"sort"
	"sync/atomic"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/gophercloud/gophercloud/openstack/dns/v2/recordsets"
	"github.com/gophercloud/gophercloud/openstack/dns/v2/zones"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
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

func (c fakeDesignateClient) ToProvider() *designateProvider {
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
	client.CreateRecordSet(zone1ID, recordsets.CreateOpts{
		Name:    "www.example.com.",
		Type:    endpoint.RecordTypeA,
		Records: []string{"10.1.1.1"},
	})
	client.CreateRecordSet(zone1ID, recordsets.CreateOpts{
		Name:    "www.example.com.",
		Type:    endpoint.RecordTypeTXT,
		Records: []string{"text1"},
	})
	client.CreateRecordSet(zone1ID, recordsets.CreateOpts{
		Name:    "xxx.example.com.",
		Type:    "SRV",
		Records: []string{"http://test.com:1234"},
	})
	client.CreateRecordSet(zone1ID, recordsets.CreateOpts{
		Name:    "ftp.example.com.",
		Type:    endpoint.RecordTypeA,
		Records: []string{"10.1.1.2"},
	})

	zone2ID := client.AddZone(zones.Zone{
		Name:   "test.net.",
		Type:   "PRIMARY",
		Status: "ACTIVE",
	})
	client.CreateRecordSet(zone2ID, recordsets.CreateOpts{
		Name:    "srv.test.net.",
		Type:    endpoint.RecordTypeA,
		Records: []string{"10.2.1.1", "10.2.1.2"},
	})
	client.CreateRecordSet(zone2ID, recordsets.CreateOpts{
		Name:    "db.test.net.",
		Type:    endpoint.RecordTypeCNAME,
		Records: []string{"sql.test.net."},
	})
	expected := []*endpoint.Endpoint{
		{
			DNSName:    "www.example.com",
			RecordType: endpoint.RecordTypeA,
			Targets:    endpoint.Targets{"10.1.1.1"},
			Labels:     endpoint.Labels{},
		},
		{
			DNSName:    "www.example.com",
			RecordType: endpoint.RecordTypeTXT,
			Targets:    endpoint.Targets{"text1"},
			Labels:     endpoint.Labels{},
		},
		{
			DNSName:    "ftp.example.com",
			RecordType: endpoint.RecordTypeA,
			Targets:    endpoint.Targets{"10.1.1.2"},
			Labels:     endpoint.Labels{},
		},
		{
			DNSName:    "srv.test.net",
			RecordType: endpoint.RecordTypeA,
			Targets:    endpoint.Targets{"10.2.1.1", "10.2.1.2"},
			Labels:     endpoint.Labels{},
		},
		{
			DNSName:    "db.test.net",
			RecordType: endpoint.RecordTypeCNAME,
			Targets:    endpoint.Targets{"sql.test.net"},
			Labels:     endpoint.Labels{},
		},
	}

	provider := client.ToProvider()
	endpoints, err := provider.Records(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	sortEndpoints := func(endpoints []*endpoint.Endpoint) {
		sort.Slice(endpoints,
			func(i, j int) bool {
				key1 := fmt.Sprintf("%s/%s", endpoints[i].DNSName, endpoints[i].RecordType)
				key2 := fmt.Sprintf("%s/%s", endpoints[j].DNSName, endpoints[j].RecordType)
				return key1 < key2
			})
	}

	sortEndpoints(endpoints)
	sortEndpoints(expected)

	if diff := cmp.Diff(expected, endpoints); diff != "" {
		t.Fatalf("unexpected endpoints:\n%s", diff)
	}
}

func TestDesignateCreateRecords(t *testing.T) {
	client := newFakeDesignateClient()
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
		},
		{
			DNSName:    "www.example.com",
			RecordType: endpoint.RecordTypeTXT,
			Targets:    endpoint.Targets{"text1"},
		},
		{
			DNSName:    "ftp.example.com",
			RecordType: endpoint.RecordTypeA,
			Targets:    endpoint.Targets{"10.1.1.2"},
		},
		{
			DNSName:    "srv.test.net",
			RecordType: endpoint.RecordTypeA,
			Targets:    endpoint.Targets{"10.2.1.1", "10.2.1.2"},
		},
		{
			DNSName:    "db.test.net",
			RecordType: endpoint.RecordTypeCNAME,
			Targets:    endpoint.Targets{"sql.test.net"},
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

	verifyDesignateEntries(t, client, expected)
}

func verifyDesignateEntries(t *testing.T, client *fakeDesignateClient, expected []*recordsets.RecordSet) {
	var results []*recordsets.RecordSet
	client.ForEachZone(func(zone *zones.Zone) error {
		client.ForEachRecordSet(zone.ID, func(recordSet *recordsets.RecordSet) error {
			recordSet.ID = "" // we don't know the id, so ignoring it
			results = append(results, recordSet)
			return nil
		})
		return nil
	})

	sortRecordSets := func(rs []*recordsets.RecordSet) {
		sort.Slice(rs, func(i, j int) bool {
			key1 := fmt.Sprintf("%s/%s", rs[i].Name, rs[i].Type)
			key2 := fmt.Sprintf("%s/%s", rs[j].Name, rs[j].Type)
			return key1 < key2
		})
	}
	sortRecordSets(expected)
	sortRecordSets(results)

	if diff := cmp.Diff(expected, results); diff != "" {
		t.Fatalf("unexpected results:\n%s", diff)
	}
}

func setupPoplulatedDesignate(client *fakeDesignateClient) {
	zone1ID := client.AddZone(zones.Zone{
		ID:     "zone-1",
		Name:   "example.com.",
		Type:   "PRIMARY",
		Status: "ACTIVE",
	})
	zone2ID := client.AddZone(zones.Zone{
		ID:     "zone-2",
		Name:   "test.net.",
		Type:   "PRIMARY",
		Status: "ACTIVE",
	})
	client.CreateRecordSet(zone1ID, recordsets.CreateOpts{
		Name:    "ftp.example.com.",
		Type:    endpoint.RecordTypeA,
		Records: []string{"10.1.1.2"},
	})
	client.CreateRecordSet(zone2ID, recordsets.CreateOpts{
		Name:    "srv.test.net.",
		Type:    endpoint.RecordTypeTXT,
		Records: []string{"hello world"},
	})
}

func TestDesignateUpdateRecords(t *testing.T) {
	client := newFakeDesignateClient()
	setupPoplulatedDesignate(client)

	// NOTE: it is rather important to test TXT records here, since they
	// are handled very differently in the registry stack
	updatesNew := []*endpoint.Endpoint{
		{
			DNSName:    "ftp.example.com",
			RecordType: endpoint.RecordTypeA,
			Targets:    endpoint.Targets{"10.3.3.1"},
		},
		{
			DNSName:    "srv.test.net.",
			RecordType: endpoint.RecordTypeTXT,
			Targets:    endpoint.Targets{"10.3.3.2"},
		},
	}
	expected := []*recordsets.RecordSet{
		{
			Name:    "ftp.example.com.",
			Type:    endpoint.RecordTypeA,
			Records: []string{"10.3.3.1"},
			ZoneID:  "zone-1",
		},
		{
			Name:    "srv.test.net.",
			Type:    endpoint.RecordTypeTXT,
			Records: []string{"10.3.3.2"},
			ZoneID:  "zone-2",
		},
	}

	err := client.ToProvider().ApplyChanges(context.Background(), &plan.Changes{UpdateNew: updatesNew})
	if err != nil {
		t.Fatal(err)
	}

	verifyDesignateEntries(t, client, expected)
}

func TestDesignateDeleteRecords(t *testing.T) {
	client := newFakeDesignateClient()
	setupPoplulatedDesignate(client)

	deletes := []*endpoint.Endpoint{
		{
			DNSName:    "srv.test.net.",
			RecordType: endpoint.RecordTypeTXT,
			Targets:    endpoint.Targets{"10.2.1.1"},
		},
	}
	expected := []*recordsets.RecordSet{
		{
			Name:    "ftp.example.com.",
			Type:    endpoint.RecordTypeA,
			Records: []string{"10.1.1.2"},
			ZoneID:  "zone-1",
		},
	}
	err := client.ToProvider().ApplyChanges(context.Background(), &plan.Changes{Delete: deletes})
	if err != nil {
		t.Fatal(err)
	}

	verifyDesignateEntries(t, client, expected)
}

func TestRecordsCache(t *testing.T) {
	client := newFakeDesignateClient()
	provider := client.ToProvider()

	provider.rsCache = map[string]*recordsets.RecordSet{
		"something": {
			Name: "www.example.com.",
			Type: "CNAME",
		},
	}
	provider.cacheRefresh = time.Now()
	provider.cacheTimeout = 1 * time.Hour

	managedZones := map[string]string{
		"ZONEID": "example.com.",
	}

	expect := map[string]*recordsets.RecordSet{
		"something": {
			Name: "www.example.com.",
			Type: "CNAME",
		},
	}

	result, err := provider.getRecordSets(context.Background(), managedZones)
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(expect, result); diff != "" {
		t.Fatalf("unexpected recordSets:\n%s", diff)
	}
}

func TestZoneCache(t *testing.T) {
	client := newFakeDesignateClient()
	provider := client.ToProvider()
	provider.zoneCache = map[string]string{
		"ZONEID": "example.com.",
	}
	provider.cacheRefresh = time.Now()
	provider.cacheTimeout = 1 * time.Hour

	expect := map[string]string{
		"ZONEID": "example.com.",
	}

	result, err := provider.getZones()
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(expect, result); diff != "" {
		t.Fatalf("unexpceted zones:\n%s", diff)
	}
}

func TestDryRun(t *testing.T) {
	client := newFakeDesignateClient()
	provider := client.ToProvider()
	provider.dryRun = true

	setupPoplulatedDesignate(client)

	deletes := []*endpoint.Endpoint{
		{
			DNSName:    "srv.test.net.",
			RecordType: endpoint.RecordTypeTXT,
			Targets:    endpoint.Targets{"hello world"},
		},
	}
	updates := []*endpoint.Endpoint{
		{
			DNSName:    "srv.test.net.",
			RecordType: endpoint.RecordTypeTXT,
			Targets:    endpoint.Targets{"nothing special"},
		},
	}
	expected := []*recordsets.RecordSet{
		{
			Name:    "ftp.example.com.",
			Type:    endpoint.RecordTypeA,
			Records: []string{"10.1.1.2"},
			ZoneID:  "zone-1",
		},
		{
			Name:    "srv.test.net.",
			Type:    endpoint.RecordTypeTXT,
			Records: []string{"hello world"},
			ZoneID:  "zone-2",
		},
	}
	err := provider.ApplyChanges(context.Background(), &plan.Changes{Delete: deletes, UpdateNew: updates})
	if err != nil {
		t.Fatal(err)
	}

	verifyDesignateEntries(t, client, expected)
}
