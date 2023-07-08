/*
Copyright 2020 The Kubernetes Authors.

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
package civo

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/civo/civogo"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

func TestNewCivoProvider(t *testing.T) {
	_ = os.Setenv("CIVO_TOKEN", "xxxxxxxxxxxxxxx")
	_, err := NewCivoProvider(endpoint.NewDomainFilter([]string{"test.civo.com"}), true)
	require.NoError(t, err)

	_ = os.Unsetenv("CIVO_TOKEN")
}

func TestNewCivoProviderNoToken(t *testing.T) {
	_, err := NewCivoProvider(endpoint.NewDomainFilter([]string{"test.civo.com"}), true)
	assert.Error(t, err)

	assert.Equal(t, "no token found", err.Error())
}

func TestCivoProviderZones(t *testing.T) {
	client, server, _ := civogo.NewClientForTesting(map[string]string{
		"/v2/dns": `[
			{"id": "12345", "account_id": "1", "name": "example.com"},
			{"id": "12346", "account_id": "1", "name": "example.net"}
			]`,
	})
	defer server.Close()
	provider := &CivoProvider{
		Client: *client,
	}

	expected, err := client.ListDNSDomains()
	assert.NoError(t, err)

	zones, err := provider.Zones(context.Background())
	assert.NoError(t, err)

	// Check if the return is a DNSDomain type
	assert.Equal(t, reflect.TypeOf(zones), reflect.TypeOf(expected))
	assert.ElementsMatch(t, zones, expected)
}

func TestCivoProviderZonesWithError(t *testing.T) {
	client, server, _ := civogo.NewClientForTesting(map[string]string{
		"/v2/dns-error": `[]`,
	})
	defer server.Close()
	provider := &CivoProvider{
		Client: *client,
	}

	_, err := provider.Zones(context.Background())
	assert.Error(t, err)
}

func TestCivoProviderRecords(t *testing.T) {
	client, server, _ := civogo.NewAdvancedClientForTesting([]civogo.ConfigAdvanceClientForTesting{
		{
			Method: "GET",
			Value: []civogo.ValueAdvanceClientForTesting{
				{
					RequestBody: ``,
					URL:         "/v2/dns/12345/records",
					ResponseBody: `[
						{"id": "1", "domain_id":"12345", "account_id": "1", "name": "www", "type": "A", "value": "10.0.0.0", "ttl": 600},
						{"id": "2", "account_id": "1", "domain_id":"12345", "name": "mail", "type": "A", "value": "10.0.0.1", "ttl": 600}
						]`,
				},
				{
					RequestBody: ``,
					URL:         "/v2/dns",
					ResponseBody: `[
						{"id": "12345", "account_id": "1", "name": "example.com"},
						{"id": "12346", "account_id": "1", "name": "example.net"}
						]`,
				},
			},
		},
	})

	defer server.Close()
	provider := &CivoProvider{
		Client:       *client,
		domainFilter: endpoint.NewDomainFilter([]string{"example.com"}),
	}

	expected, err := client.ListDNSRecords("12345")
	assert.NoError(t, err)

	records, err := provider.Records(context.Background())
	assert.NoError(t, err)

	assert.Equal(t, strings.TrimSuffix(records[0].DNSName, ".example.com"), expected[0].Name)
	assert.Equal(t, records[0].RecordType, string(expected[0].Type))
	assert.Equal(t, int(records[0].RecordTTL), expected[0].TTL)

	assert.Equal(t, strings.TrimSuffix(records[1].DNSName, ".example.com"), expected[1].Name)
	assert.Equal(t, records[1].RecordType, string(expected[1].Type))
	assert.Equal(t, int(records[1].RecordTTL), expected[1].TTL)
}

func TestCivoProviderWithoutRecords(t *testing.T) {
	client, server, _ := civogo.NewClientForTesting(map[string]string{
		"/v2/dns/12345/records": `[]`,
		"/v2/dns": `[
			{"id": "12345", "account_id": "1", "name": "example.com"},
			{"id": "12346", "account_id": "1", "name": "example.net"}
			]`,
	})
	defer server.Close()
	provider := &CivoProvider{
		Client:       *client,
		domainFilter: endpoint.NewDomainFilter([]string{"example.com"}),
	}

	records, err := provider.Records(context.Background())
	assert.NoError(t, err)

	assert.Equal(t, len(records), 0)
}

func TestCivoProcessCreateActions(t *testing.T) {
	zoneByID := map[string]civogo.DNSDomain{
		"example.com": {
			ID:        "1",
			AccountID: "1",
			Name:      "example.com",
		},
	}

	recordsByZoneID := map[string][]civogo.DNSRecord{
		"example.com": {
			{
				ID:          "1",
				AccountID:   "1",
				DNSDomainID: "1",
				Name:        "txt",
				Value:       "12.12.12.1",
				Type:        "A",
				TTL:         600,
			},
		},
	}

	createsByZone := map[string][]*endpoint.Endpoint{
		"example.com": {
			endpoint.NewEndpoint("foo.example.com", endpoint.RecordTypeA, "1.2.3.4"),
			endpoint.NewEndpoint("txt.example.com", endpoint.RecordTypeCNAME, "foo.example.com"),
		},
	}

	var changes CivoChanges
	err := processCreateActions(zoneByID, recordsByZoneID, createsByZone, &changes)
	require.NoError(t, err)

	assert.Equal(t, 2, len(changes.Creates))
	assert.Equal(t, 0, len(changes.Updates))
	assert.Equal(t, 0, len(changes.Deletes))

	expectedCreates := []*CivoChangeCreate{
		{
			Domain: civogo.DNSDomain{
				ID:        "1",
				AccountID: "1",
				Name:      "example.com",
			},
			Options: &civogo.DNSRecordConfig{
				Type:  "A",
				Name:  "foo",
				Value: "1.2.3.4",
			},
		},
		{
			Domain: civogo.DNSDomain{
				ID:        "1",
				AccountID: "1",
				Name:      "example.com",
			},
			Options: &civogo.DNSRecordConfig{
				Type:  "CNAME",
				Name:  "txt",
				Value: "foo.example.com",
			},
		},
	}

	if !elementsMatch(t, expectedCreates, changes.Creates) {
		assert.Failf(t, "diff: %s", cmp.Diff(expectedCreates, changes.Creates))
	}
}

func TestCivoProcessCreateActionsWithError(t *testing.T) {
	zoneByID := map[string]civogo.DNSDomain{
		"example.com": {
			ID:        "1",
			AccountID: "1",
			Name:      "example.com",
		},
	}

	recordsByZoneID := map[string][]civogo.DNSRecord{
		"example.com": {
			{
				ID:          "1",
				AccountID:   "1",
				DNSDomainID: "1",
				Name:        "txt",
				Value:       "12.12.12.1",
				Type:        "A",
				TTL:         600,
			},
		},
	}

	createsByZone := map[string][]*endpoint.Endpoint{
		"example.com": {
			endpoint.NewEndpoint("foo.example.com", "AAAA", "1.2.3.4"),
			endpoint.NewEndpoint("txt.example.com", endpoint.RecordTypeCNAME, "foo.example.com"),
		},
	}

	var changes CivoChanges
	err := processCreateActions(zoneByID, recordsByZoneID, createsByZone, &changes)
	require.Error(t, err)
	assert.Equal(t, "invalid Record Type: AAAA", err.Error())
}

func TestCivoProcessUpdateActions(t *testing.T) {
	zoneByID := map[string]civogo.DNSDomain{
		"example.com": {
			ID:        "1",
			AccountID: "1",
			Name:      "example.com",
		},
	}

	recordsByZoneID := map[string][]civogo.DNSRecord{
		"example.com": {
			{
				ID:          "1",
				AccountID:   "1",
				DNSDomainID: "1",
				Name:        "txt",
				Value:       "1.2.3.4",
				Type:        "A",
				TTL:         600,
			},
			{
				ID:          "2",
				AccountID:   "1",
				DNSDomainID: "1",
				Name:        "foo",
				Value:       "foo.example.com",
				Type:        "CNAME",
				TTL:         600,
			},
			{
				ID:          "3",
				AccountID:   "1",
				DNSDomainID: "1",
				Name:        "bar",
				Value:       "10.10.10.1",
				Type:        "A",
				TTL:         600,
			},
		},
	}

	updatesByZone := map[string][]*endpoint.Endpoint{
		"example.com": {
			endpoint.NewEndpoint("txt.example.com", endpoint.RecordTypeA, "10.20.30.40"),
			endpoint.NewEndpoint("foo.example.com", endpoint.RecordTypeCNAME, "bar.example.com"),
		},
	}

	var changes CivoChanges
	err := processUpdateActions(zoneByID, recordsByZoneID, updatesByZone, &changes)
	require.NoError(t, err)

	assert.Equal(t, 2, len(changes.Creates))
	assert.Equal(t, 0, len(changes.Updates))
	assert.Equal(t, 2, len(changes.Deletes))

	expectedUpdate := []*CivoChangeCreate{
		{
			Domain: civogo.DNSDomain{
				ID:        "1",
				AccountID: "1",
				Name:      "example.com",
			},
			Options: &civogo.DNSRecordConfig{
				Type:  "A",
				Name:  "txt",
				Value: "10.20.30.40",
			},
		},
		{
			Domain: civogo.DNSDomain{
				ID:        "1",
				AccountID: "1",
				Name:      "example.com",
			},
			Options: &civogo.DNSRecordConfig{
				Type:  "CNAME",
				Name:  "foo",
				Value: "bar.example.com",
			},
		},
	}

	if !elementsMatch(t, expectedUpdate, changes.Creates) {
		assert.Failf(t, "diff: %s", cmp.Diff(expectedUpdate, changes.Creates))
	}

	expectedDelete := []*CivoChangeDelete{
		{
			Domain: civogo.DNSDomain{
				ID:        "1",
				AccountID: "1",
				Name:      "example.com",
			},
			DomainRecord: civogo.DNSRecord{
				ID:          "1",
				AccountID:   "1",
				DNSDomainID: "1",
				Name:        "txt",
				Value:       "1.2.3.4",
				Type:        "A",
				Priority:    0,
				TTL:         600,
			},
		},
		{
			Domain: civogo.DNSDomain{
				ID:        "1",
				AccountID: "1",
				Name:      "example.com",
			},
			DomainRecord: civogo.DNSRecord{
				ID:          "2",
				AccountID:   "1",
				DNSDomainID: "1",
				Name:        "foo",
				Value:       "foo.example.com",
				Type:        "CNAME",
				Priority:    0,
				TTL:         600,
			},
		},
	}

	if !elementsMatch(t, expectedDelete, changes.Deletes) {
		assert.Failf(t, "diff: %s", cmp.Diff(expectedDelete, changes.Deletes))
	}
}

func TestCivoProcessDeleteAction(t *testing.T) {
	zoneByID := map[string]civogo.DNSDomain{
		"example.com": {
			ID:        "1",
			AccountID: "1",
			Name:      "example.com",
		},
	}

	recordsByZoneID := map[string][]civogo.DNSRecord{
		"example.com": {
			{
				ID:          "1",
				AccountID:   "1",
				DNSDomainID: "1",
				Name:        "txt",
				Value:       "1.2.3.4",
				Type:        "A",
				TTL:         600,
			},
			{
				ID:          "2",
				AccountID:   "1",
				DNSDomainID: "1",
				Name:        "foo",
				Value:       "5.6.7.8",
				Type:        "A",
				TTL:         600,
			},
			{
				ID:          "3",
				AccountID:   "1",
				DNSDomainID: "1",
				Name:        "bar",
				Value:       "10.10.10.1",
				Type:        "A",
				TTL:         600,
			},
		},
	}

	deleteByDomain := map[string][]*endpoint.Endpoint{
		"example.com": {
			endpoint.NewEndpoint("txt.example.com", endpoint.RecordTypeA, "1.2.3.4"),
			endpoint.NewEndpoint("foo.example.com", endpoint.RecordTypeA, "5.6.7.8"),
		},
	}

	var changes CivoChanges
	err := processDeleteActions(zoneByID, recordsByZoneID, deleteByDomain, &changes)
	require.NoError(t, err)

	assert.Equal(t, 0, len(changes.Creates))
	assert.Equal(t, 0, len(changes.Updates))
	assert.Equal(t, 2, len(changes.Deletes))

	expectedDelete := []*CivoChangeDelete{
		{
			Domain: civogo.DNSDomain{
				ID:        "1",
				AccountID: "1",
				Name:      "example.com",
			},
			DomainRecord: civogo.DNSRecord{
				ID:          "1",
				AccountID:   "1",
				DNSDomainID: "1",
				Name:        "txt",
				Value:       "1.2.3.4",
				Type:        "A",
				TTL:         600,
			},
		},
		{
			Domain: civogo.DNSDomain{
				ID:        "1",
				AccountID: "1",
				Name:      "example.com",
			},
			DomainRecord: civogo.DNSRecord{
				ID:          "2",
				AccountID:   "1",
				DNSDomainID: "1",
				Type:        "A",
				Name:        "foo",
				Value:       "5.6.7.8",
				TTL:         600,
			},
		},
	}

	if !elementsMatch(t, expectedDelete, changes.Deletes) {
		assert.Failf(t, "diff: %s", cmp.Diff(expectedDelete, changes.Deletes))
	}
}

func TestCivoApplyChanges(t *testing.T) {
	client, server, _ := civogo.NewAdvancedClientForTesting([]civogo.ConfigAdvanceClientForTesting{
		{
			Method: "GET",
			Value: []civogo.ValueAdvanceClientForTesting{
				{
					RequestBody:  "",
					URL:          "/v2/dns",
					ResponseBody: `[{"id": "12345", "account_id": "1", "name": "example.com"}]`,
				},
				{
					RequestBody:  "",
					URL:          "/v2/dns/12345/records",
					ResponseBody: `[]`,
				},
			},
		},
	})
	defer server.Close()

	changes := &plan.Changes{}
	provider := &CivoProvider{
		Client: *client,
	}
	changes.Create = []*endpoint.Endpoint{
		{DNSName: "new.ext-dns-test.example.com", Targets: endpoint.Targets{"target"}, RecordType: endpoint.RecordTypeA},
		{DNSName: "new.ext-dns-test-with-ttl.example.com", Targets: endpoint.Targets{"target"}, RecordType: endpoint.RecordTypeA, RecordTTL: 100},
	}
	changes.Delete = []*endpoint.Endpoint{{DNSName: "foobar.ext-dns-test.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"target"}}}
	changes.UpdateOld = []*endpoint.Endpoint{{DNSName: "foobar.ext-dns-test.example.de", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"target-old"}}}
	changes.UpdateNew = []*endpoint.Endpoint{{DNSName: "foobar.ext-dns-test.foo.com", Targets: endpoint.Targets{"target-new"}, RecordType: endpoint.RecordTypeCNAME, RecordTTL: 100}}
	err := provider.ApplyChanges(context.Background(), changes)
	assert.NoError(t, err)
}

func TestCivoProviderFetchZones(t *testing.T) {
	client, server, _ := civogo.NewClientForTesting(map[string]string{
		"/v2/dns": `[
			{"id": "12345", "account_id": "1", "name": "example.com"},
			{"id": "12346", "account_id": "1", "name": "example.net"}
			]`,
	})
	defer server.Close()
	provider := &CivoProvider{
		Client: *client,
	}

	expected, err := client.ListDNSDomains()
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
	zones, err := provider.fetchZones(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	assert.ElementsMatch(t, zones, expected)
}
func TestCivoProviderFetchZonesWithFilter(t *testing.T) {
	client, server, _ := civogo.NewClientForTesting(map[string]string{
		"/v2/dns": `[
			{"id": "12345", "account_id": "1", "name": "example.com"},
			{"id": "12346", "account_id": "1", "name": "example.net"}
			]`,
	})
	defer server.Close()
	provider := &CivoProvider{
		Client:       *client,
		domainFilter: endpoint.NewDomainFilter([]string{".com"}),
	}

	expected := []civogo.DNSDomain{
		{ID: "12345", Name: "example.com", AccountID: "1"},
	}

	actual, err := provider.fetchZones(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	assert.ElementsMatch(t, expected, actual)
}

func TestCivoProviderFetchRecords(t *testing.T) {
	client, server, _ := civogo.NewClientForTesting(map[string]string{
		"/v2/dns/12345/records": `[
			{"id": "1", "domain_id":"12345", "account_id": "1", "name": "www", "type": "A", "value": "10.0.0.0", "ttl": 600},
			{"id": "2", "account_id": "1", "domain_id":"12345", "name": "mail", "type": "A", "value": "10.0.0.1", "ttl": 600}
			]`,
	})
	defer server.Close()
	provider := &CivoProvider{
		Client: *client,
	}

	expected, err := client.ListDNSRecords("12345")
	assert.NoError(t, err)

	actual, err := provider.fetchRecords(context.Background(), "12345")
	assert.NoError(t, err)

	assert.ElementsMatch(t, expected, actual)
}

func TestCivoProviderFetchRecordsWithError(t *testing.T) {
	client, server, _ := civogo.NewClientForTesting(map[string]string{
		"/v2/dns/12345/records": `[
			{"id": "1", "domain_id":"12345", "account_id": "1", "name": "www", "type": "A", "value": "10.0.0.0", "ttl": 600},
			{"id": "2", "account_id": "1", "domain_id":"12345", "name": "mail", "type": "A", "value": "10.0.0.1", "ttl": 600}
			]`,
	})
	defer server.Close()
	provider := &CivoProvider{
		Client: *client,
	}

	_, err := provider.fetchRecords(context.Background(), "235698")
	assert.Error(t, err)
}

func TestCivo_getStrippedRecordName(t *testing.T) {
	assert.Equal(t, "", getStrippedRecordName(civogo.DNSDomain{
		Name: "foo.com",
	}, endpoint.Endpoint{
		DNSName: "foo.com",
	}))

	assert.Equal(t, "api", getStrippedRecordName(civogo.DNSDomain{
		Name: "foo.com",
	}, endpoint.Endpoint{
		DNSName: "api.foo.com",
	}))
}

func TestCivo_convertRecordType(t *testing.T) {
	record, err := convertRecordType("A")
	recordA := civogo.DNSRecordType(civogo.DNSRecordTypeA)
	require.NoError(t, err)
	assert.Equal(t, recordA, record)

	record, err = convertRecordType("CNAME")
	recordCName := civogo.DNSRecordType(civogo.DNSRecordTypeCName)
	require.NoError(t, err)
	assert.Equal(t, recordCName, record)

	record, err = convertRecordType("TXT")
	recordTXT := civogo.DNSRecordType(civogo.DNSRecordTypeTXT)
	require.NoError(t, err)
	assert.Equal(t, recordTXT, record)

	record, err = convertRecordType("SRV")
	recordSRV := civogo.DNSRecordType(civogo.DNSRecordTypeSRV)
	require.NoError(t, err)
	assert.Equal(t, recordSRV, record)

	_, err = convertRecordType("INVALID")
	require.Error(t, err)

	assert.Equal(t, "invalid Record Type: INVALID", err.Error())
}

func TestCivoProviderGetRecordID(t *testing.T) {
	zone := civogo.DNSDomain{
		ID:   "12345",
		Name: "test.com",
	}

	record := []civogo.DNSRecord{{
		ID:          "1",
		Type:        "A",
		Name:        "www",
		Value:       "10.0.0.0",
		DNSDomainID: "12345",
		TTL:         600,
	}, {
		ID:          "2",
		Type:        "A",
		Name:        "api",
		Value:       "10.0.0.1",
		DNSDomainID: "12345",
		TTL:         600,
	}}

	endPoint := endpoint.Endpoint{DNSName: "www.test.com", Targets: endpoint.Targets{"10.0.0.0"}, RecordType: "A"}
	id := getRecordID(record, zone, endPoint)

	assert.Equal(t, id[0].ID, record[0].ID)
}

func TestCivo_submitChangesCreate(t *testing.T) {
	client, server, _ := civogo.NewAdvancedClientForTesting([]civogo.ConfigAdvanceClientForTesting{
		{
			Method: "POST",
			Value: []civogo.ValueAdvanceClientForTesting{
				{
					RequestBody: `{"type":"MX","name":"mail","value":"10.0.0.1","priority":10,"ttl":600}`,
					URL:         "/v2/dns/12345/records",
					ResponseBody: `{
						"id": "76cc107f-fbef-4e2b-b97f-f5d34f4075d3",
						"account_id": "1",
						"domain_id": "12345",
						"name": "mail",
						"value": "10.0.0.1",
						"type": "MX",
						"priority": 10,
						"ttl": 600
					}`,
				},
			},
		},
	})
	defer server.Close()

	provider := &CivoProvider{
		Client: *client,
		DryRun: false,
	}

	changes := CivoChanges{
		Creates: []*CivoChangeCreate{
			{
				Domain: civogo.DNSDomain{
					ID:        "12345",
					AccountID: "1",
					Name:      "example.com",
				},
				Options: &civogo.DNSRecordConfig{
					Type:     "MX",
					Name:     "mail",
					Value:    "10.0.0.1",
					Priority: 10,
					TTL:      600,
				},
			},
		},
	}

	err := provider.submitChanges(context.Background(), changes)
	assert.NoError(t, err)
}

func TestCivo_submitChangesUpdate(t *testing.T) {
	client, server, _ := civogo.NewAdvancedClientForTesting([]civogo.ConfigAdvanceClientForTesting{
		{
			Method: "PUT",
			Value: []civogo.ValueAdvanceClientForTesting{
				{
					RequestBody: `{"type":"MX","name":"mail","value":"10.0.0.2","priority":10,"ttl":600}`,
					URL:         "/v2/dns/12345/records/76cc107f-fbef-4e2b-b97f-f5d34f4075d3",
					ResponseBody: `{
						"id": "76cc107f-fbef-4e2b-b97f-f5d34f4075d3",
						"account_id": "1",
						"domain_id": "12345",
						"name": "mail",
						"value": "10.0.0.2",
						"type": "MX",
						"priority": 10,
						"ttl": 600
					}`,
				},
			},
		},
	})
	defer server.Close()

	provider := &CivoProvider{
		Client: *client,
		DryRun: false,
	}

	changes := CivoChanges{
		Updates: []*CivoChangeUpdate{
			{
				Domain: civogo.DNSDomain{ID: "12345", AccountID: "1", Name: "example.com"},
				DomainRecord: civogo.DNSRecord{
					ID:          "76cc107f-fbef-4e2b-b97f-f5d34f4075d3",
					AccountID:   "1",
					DNSDomainID: "12345",
					Name:        "mail",
					Value:       "10.0.0.1",
					Type:        "MX",
					Priority:    10,
					TTL:         600,
				},
				Options: civogo.DNSRecordConfig{
					Type:     "MX",
					Name:     "mail",
					Value:    "10.0.0.2",
					Priority: 10,
					TTL:      600,
				},
			},
		},
	}

	err := provider.submitChanges(context.Background(), changes)
	assert.NoError(t, err)
}

func TestCivo_submitChangesDelete(t *testing.T) {
	client, server, _ := civogo.NewClientForTesting(map[string]string{
		"/v2/dns/12345/records/76cc107f-fbef-4e2b-b97f-f5d34f4075d3": `{"result": "success"}`,
	})
	defer server.Close()

	provider := &CivoProvider{
		Client: *client,
		DryRun: false,
	}

	changes := CivoChanges{
		Deletes: []*CivoChangeDelete{
			{
				Domain: civogo.DNSDomain{ID: "12345", AccountID: "1", Name: "example.com"},
				DomainRecord: civogo.DNSRecord{
					ID:          "76cc107f-fbef-4e2b-b97f-f5d34f4075d3",
					AccountID:   "1",
					DNSDomainID: "12345",
					Name:        "mail",
					Value:       "10.0.0.2",
					Type:        "MX",
					Priority:    10,
					TTL:         600,
				},
			},
		},
	}

	err := provider.submitChanges(context.Background(), changes)
	assert.NoError(t, err)
}

func TestCivoChangesEmpty(t *testing.T) {
	// Test empty CivoChanges
	c := &CivoChanges{}
	assert.True(t, c.Empty())

	// Test CivoChanges with Creates
	c = &CivoChanges{
		Creates: []*CivoChangeCreate{
			{
				Domain: civogo.DNSDomain{
					ID:        "12345",
					AccountID: "1",
					Name:      "example.com",
				},
				Options: &civogo.DNSRecordConfig{
					Type:     civogo.DNSRecordTypeA,
					Name:     "www",
					Value:    "192.1.1.1",
					Priority: 0,
					TTL:      600,
				},
			},
		},
	}
	assert.False(t, c.Empty())

	// Test CivoChanges with Updates
	c = &CivoChanges{
		Updates: []*CivoChangeUpdate{
			{
				Domain: civogo.DNSDomain{
					ID:        "12345",
					AccountID: "1",
					Name:      "example.com",
				},
				DomainRecord: civogo.DNSRecord{
					ID:          "76cc107f-fbef-4e2b-b97f-f5d34f4075d3",
					AccountID:   "1",
					DNSDomainID: "12345",
					Name:        "mail",
					Value:       "192.168.1.2",
					Type:        "MX",
					Priority:    10,
					TTL:         600,
				},
				Options: civogo.DNSRecordConfig{
					Type:     "MX",
					Name:     "mail",
					Value:    "192.168.1.3",
					Priority: 10,
					TTL:      600,
				},
			},
		},
	}
	assert.False(t, c.Empty())

	// Test CivoChanges with Deletes
	c = &CivoChanges{
		Deletes: []*CivoChangeDelete{
			{
				Domain: civogo.DNSDomain{
					ID:        "12345",
					AccountID: "1",
					Name:      "example.com",
				},
				DomainRecord: civogo.DNSRecord{
					ID:          "76cc107f-fbef-4e2b-b97f-f5d34f4075d3",
					AccountID:   "1",
					DNSDomainID: "12345",
					Name:        "mail",
					Value:       "192.168.1.3",
					Type:        "MX",
					Priority:    10,
					TTL:         600,
				},
			},
		},
	}
	assert.False(t, c.Empty())

	// Test CivoChanges with Creates, Updates, and Deletes
	c = &CivoChanges{
		Creates: []*CivoChangeCreate{
			{
				Domain: civogo.DNSDomain{
					ID:        "12345",
					AccountID: "1",
					Name:      "example.com",
				},
				Options: &civogo.DNSRecordConfig{
					Type:     civogo.DNSRecordTypeA,
					Name:     "www",
					Value:    "192.1.1.1",
					Priority: 0,
					TTL:      600,
				},
			},
		},
		Updates: []*CivoChangeUpdate{
			{
				Domain: civogo.DNSDomain{
					ID:        "12345",
					AccountID: "1",
					Name:      "example.com",
				},
				DomainRecord: civogo.DNSRecord{
					ID:          "76cc107f-fbef-4e2b-b97f-f5d34f4075d3",
					AccountID:   "1",
					DNSDomainID: "12345",
					Name:        "mail",
					Value:       "192.168.1.2",
					Type:        "MX",
					Priority:    10,
					TTL:         600,
				},
				Options: civogo.DNSRecordConfig{
					Type:     "MX",
					Name:     "mail",
					Value:    "192.168.1.3",
					Priority: 10,
					TTL:      600,
				},
			},
		},
		Deletes: []*CivoChangeDelete{
			{
				Domain: civogo.DNSDomain{
					ID:        "12345",
					AccountID: "1",
					Name:      "example.com",
				},
				DomainRecord: civogo.DNSRecord{
					ID:          "76cc107f-fbef-4e2b-b97f-f5d34f4075d3",
					AccountID:   "1",
					DNSDomainID: "12345",
					Name:        "mail",
					Value:       "192.168.1.3",
					Type:        "MX",
					Priority:    10,
					TTL:         600,
				},
			},
		},
	}
	assert.False(t, c.Empty())
}

// This function is an adapted copy of the testify package's ElementsMatch function with the
// call to ObjectsAreEqual replaced with cmp.Equal which better handles struct's with pointers to
// other structs. It also ignores ordering when comparing unlike cmp.Equal.
func elementsMatch(t *testing.T, listA, listB interface{}, msgAndArgs ...interface{}) (ok bool) {
	if listA == nil && listB == nil {
		return true
	} else if listA == nil {
		return isEmpty(listB)
	} else if listB == nil {
		return isEmpty(listA)
	}

	aKind := reflect.TypeOf(listA).Kind()
	bKind := reflect.TypeOf(listB).Kind()

	if aKind != reflect.Array && aKind != reflect.Slice {
		return assert.Fail(t, fmt.Sprintf("%q has an unsupported type %s", listA, aKind), msgAndArgs...)
	}

	if bKind != reflect.Array && bKind != reflect.Slice {
		return assert.Fail(t, fmt.Sprintf("%q has an unsupported type %s", listB, bKind), msgAndArgs...)
	}

	aValue := reflect.ValueOf(listA)
	bValue := reflect.ValueOf(listB)

	aLen := aValue.Len()
	bLen := bValue.Len()

	if aLen != bLen {
		return assert.Fail(t, fmt.Sprintf("lengths don't match: %d != %d", aLen, bLen), msgAndArgs...)
	}

	// Mark indexes in bValue that we already used
	visited := make([]bool, bLen)
	for i := 0; i < aLen; i++ {
		element := aValue.Index(i).Interface()
		found := false
		for j := 0; j < bLen; j++ {
			if visited[j] {
				continue
			}
			if cmp.Equal(bValue.Index(j).Interface(), element) {
				visited[j] = true
				found = true
				break
			}
		}
		if !found {
			return assert.Fail(t, fmt.Sprintf("element %s appears more times in %s than in %s", element, aValue, bValue), msgAndArgs...)
		}
	}

	return true
}

func isEmpty(xs interface{}) bool {
	if xs != nil {
		objValue := reflect.ValueOf(xs)
		return objValue.Len() == 0
	}
	return true
}
