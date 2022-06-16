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

package stackpath

import (
	"context"
	"os"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wmarchesi123/stackpath-go/pkg/dns"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

var testProvider = StackPathProvider{
	client:       &dns.APIClient{},
	context:      context.Background(),
	domainFilter: endpoint.DomainFilter{},
	zoneIDFilter: provider.ZoneIDFilter{},
	stackID:      "TEST_STACK_ID",
	dryRun:       false,
	testing:      true,
}

func TestNewStackPathProvider(t *testing.T) {
	stackPathConfig := &StackPathConfig{
		Context:      context.Background(),
		DomainFilter: endpoint.NewDomainFilter(nil),
		ZoneIDFilter: provider.NewZoneIDFilter(nil),
		DryRun:       false,
		Testing:      true,
	}

	_, err := NewStackPathProvider(*stackPathConfig)
	if err == nil {
		t.Fatalf("Expected to fail without a valid CLIENT_ID, CLIENT_SECRET, and STACK_ID")
	}

	_ = os.Setenv("STACKPATH_CLIENT_ID", "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	_ = os.Setenv("STACKPATH_CLIENT_SECRET", "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	_ = os.Setenv("STACKPATH_STACK_ID", "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

	_, err = NewStackPathProvider(*stackPathConfig)
	if err.Error() != "Error obtaining oauth2 token: 404 Not Found" {
		t.Fatalf("%v", err)
	}
}

func TestZones(t *testing.T) {
	zones, err := testProvider.zones()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(zones))
	assert.Contains(t, zones[0].GetNameservers(), "ns1.example.com")
	assert.Equal(t, false, zones[0].GetDisabled())
}

func TestGetZones(t *testing.T) {
	zoneGetZonesResponse, _, err := testProvider.getZones()

	assert.NoError(t, err)
	assert.Equal(t, true, zoneGetZonesResponse.HasZones())
	zones := zoneGetZonesResponse.GetZones()
	assert.Equal(t, 1, len(zones))
	assert.Contains(t, zones[0].GetNameservers(), "ns1.example.com")
	assert.Equal(t, false, zones[0].GetDisabled())
}

func TestGetZoneRecords(t *testing.T) {
	zoneGetZoneRecordsResponse, _, err := testProvider.getZoneRecords("")
	assert.NoError(t, err)
	assert.Equal(t, true, zoneGetZoneRecordsResponse.HasRecords())
	records := zoneGetZoneRecordsResponse.GetRecords()
	assert.Equal(t, 3, len(records))
	assert.Equal(t, "www", records[0].GetName())
	assert.Equal(t, "testing.com", records[2].GetData())
}

func TestRecords(t *testing.T) {
	endpoints, err := testProvider.Records(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, 2, len(endpoints))
	assert.GreaterOrEqual(t, len(endpoints[0].Targets), 1)
	assert.GreaterOrEqual(t, len(endpoints[1].Targets), 1)
}

func TestStackPathStyleRecords(t *testing.T) {
	records, err := testProvider.StackPathStyleRecords()
	assert.NoError(t, err)
	assert.Equal(t, 3, len(records))
}

func TestApplyChanges(t *testing.T) {
	err := testProvider.ApplyChanges(context.Background(), &plan.Changes{})
	assert.NoError(t, err)
}

func TestRecordFromTarget(t *testing.T) {
	endpoint := &endpoint.Endpoint{
		DNSName:          "www.one.com",
		Targets:          []string{"1.1.1.1"},
		RecordType:       "A",
		SetIdentifier:    "test",
		RecordTTL:        endpoint.TTL(60),
		Labels:           endpoint.Labels{},
		ProviderSpecific: endpoint.ProviderSpecific{},
	}

	record, err := recordFromTarget(endpoint, "1.1.1.1", &testGetZoneZoneRecords, "one.com")

	assert.NoError(t, err)
	assert.Equal(t, "TEST_ZONE_ZONE_RECORD_ID1", record)

	record, err = recordFromTarget(endpoint, "2.2.2.2", &testGetZoneZoneRecords, "one.com")

	assert.NoError(t, err)
	assert.Equal(t, "TEST_ZONE_ZONE_RECORD_ID2", record)
}

func TestMergeEndpointsByNameType(t *testing.T) {
	endpoints := mergeEndpointsByNameType(testEndpoints)

	sort.Slice(endpoints, func(i, j int) bool {
		if endpoints[i].DNSName < endpoints[j].DNSName {
			return true
		} else {
			return false
		}
	})

	assert.Equal(t, 2, len(endpoints))
	assert.Equal(t, testMergedEndpoints[0].DNSName, endpoints[0].DNSName)
	assert.Contains(t, endpoints[0].Targets, testMergedEndpoints[0].Targets[0])
	assert.Contains(t, endpoints[1].Targets, testMergedEndpoints[1].Targets[0])
	assert.Contains(t, endpoints[1].Targets, testMergedEndpoints[1].Targets[1])

	assert.Equal(t, testMergedEndpoints[0], endpoints[0])
}
