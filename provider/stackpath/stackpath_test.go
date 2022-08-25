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

package stackpath

import (
	"context"
	"fmt"
	"os"
	"sort"
	"testing"

	"github.com/kubeslice/stackpath/pkg/dns"
	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

var testProvider = StackPathProvider{
	client:       &dns.APIClient{},
	context:      context.Background(),
	domainFilter: endpoint.DomainFilter{},
	zoneIDFilter: provider.ZoneIDFilter{},
	ownerID:      "test",
	stackID:      "TEST_STACK_ID",
	dryRun:       false,
	testing:      true,
}

func TestNewStackPathProvider(t *testing.T) {
	stackPathConfig := &StackPathConfig{
		Context:      context.Background(),
		DomainFilter: endpoint.NewDomainFilter(nil),
		ZoneIDFilter: provider.NewZoneIDFilter(nil),
		OwnerID:      "test",
		DryRun:       false,
		Testing:      true,
	}

	_, err := NewStackPathProvider(*stackPathConfig)
	assert.Equal(t, fmt.Errorf("STACKPATH_CLIENT_ID environment variable is not set"), err)
	_ = os.Setenv("STACKPATH_CLIENT_ID", "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

	_, err = NewStackPathProvider(*stackPathConfig)
	assert.Equal(t, fmt.Errorf("STACKPATH_CLIENT_SECRET environment variable is not set"), err)
	_ = os.Setenv("STACKPATH_CLIENT_SECRET", "IS_SET")

	_, err = NewStackPathProvider(*stackPathConfig)
	assert.Equal(t, fmt.Errorf("STACKPATH_STACK_ID environment variable is not set"), err)
	_ = os.Setenv("STACKPATH_STACK_ID", "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

	_, err = NewStackPathProvider(*stackPathConfig)
	assert.Equal(t, fmt.Errorf("Error obtaining oauth2 token: 404 Not Found"), err)
	_ = os.Setenv("STACKPATH_CLIENT_SECRET", "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

	_, err = NewStackPathProvider(*stackPathConfig)
	assert.Equal(t, nil, err)
}

func TestZones(t *testing.T) {
	testProvider.dryRun = false
	zones, err := testProvider.zones()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(zones))
	assert.Contains(t, zones[0].GetNameservers(), "ns1.example.com")
	assert.Equal(t, false, zones[0].GetDisabled())

	testProvider.dryRun = true
	_, err = testProvider.zones()
	assert.Equal(t, fmt.Errorf("testing"), err)
	testProvider.dryRun = false

	testProvider.zoneIDFilter = provider.NewZoneIDFilter([]string{"TEST_ZONE_ID"})
	_, err = testProvider.zones()
	testProvider.zoneIDFilter = provider.NewZoneIDFilter(nil)
	assert.NoError(t, err)
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

	testProvider.dryRun = true
	_, _, err = testProvider.getZoneRecords("")
	assert.Equal(t, fmt.Errorf("testing"), err)
	testProvider.dryRun = false
}

func TestRecords(t *testing.T) {
	endpoints, err := testProvider.Records(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, 2, len(endpoints))
	assert.GreaterOrEqual(t, len(endpoints[0].Targets), 1)
	assert.GreaterOrEqual(t, len(endpoints[1].Targets), 1)

	testProvider.dryRun = true
	_, err = testProvider.Records(context.Background())
	assert.Equal(t, fmt.Errorf("testing"), err)
	testProvider.dryRun = false
}

func TestStackPathStyleRecords(t *testing.T) {
	records, err := testProvider.StackPathStyleRecords()
	assert.NoError(t, err)
	assert.Equal(t, 3, len(records))

	testProvider.dryRun = true
	_, err = testProvider.StackPathStyleRecords()
	assert.Equal(t, fmt.Errorf("testing"), err)
	testProvider.dryRun = false
}

func TestApplyChanges(t *testing.T) {
	err := testProvider.ApplyChanges(context.Background(), &plan.Changes{})
	assert.NoError(t, err)

	testProvider.dryRun = true
	err = testProvider.ApplyChanges(context.Background(), &plan.Changes{})
	assert.Equal(t, nil, err)
	testProvider.dryRun = false

	err = testProvider.ApplyChanges(context.Background(), testChanges)
	assert.Equal(t, fmt.Errorf("record not found"), err)

	testProvider.dryRun = true
	err = testProvider.ApplyChanges(context.Background(), testChanges)
	assert.Equal(t, fmt.Errorf("testing"), err)
	testProvider.dryRun = false
}

func TestRecordFromTarget(t *testing.T) {
	endpoint := &endpoint.Endpoint{
		DNSName:          "www.one.com",
		Targets:          []string{"1.1.1.1"},
		RecordType:       "A",
		SetIdentifier:    "test",
		RecordTTL:        endpoint.TTL(60),
		Labels:           testZoneZoneRecordLabels,
		ProviderSpecific: endpoint.ProviderSpecific{},
	}

	record, err := recordFromTarget(endpoint, "1.1.1.1", &testGetZoneZoneRecords, "one.com", "test")
	assert.NoError(t, err)
	assert.Equal(t, "TEST_ZONE_ZONE_RECORD_ID1", record)

	record, err = recordFromTarget(endpoint, "2.2.2.2", &testGetZoneZoneRecords, "one.com", "test")
	assert.NoError(t, err)
	assert.Equal(t, "TEST_ZONE_ZONE_RECORD_ID2", record)

	endpoint.DNSName = ""
	_, err = recordFromTarget(endpoint, "3.3.3.3", &testGetZoneZoneRecords, "one.com", "test")
	assert.Equal(t, fmt.Errorf("record not found"), err)
}

func TestMergeEndpointsByNameType(t *testing.T) {
	endpoints := mergeEndpointsByNameType(threeTestEndpoints)

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

	assert.Equal(t, testMergedEndpoints, mergeEndpointsByNameType(testMergedEndpoints))
}

func TestEndpointsByZoneID(t *testing.T) {

	zoneIDNameMap := provider.ZoneIDName{}

	zoneResponse, _, err := testProvider.getZones()
	assert.NoError(t, err)
	zones := append(zoneResponse.GetZones(), badZone)

	for _, zone := range zones {
		zoneIDNameMap.Add(zone.GetId(), zone.GetDomain())
	}

	endpointByZoneIDMap := endpointsByZoneID(zoneIDNameMap, threeTestEndpoints)
	assert.Equal(t, 1, len(endpointByZoneIDMap))
}

func TestCreate(t *testing.T) {
	zoneIDNameMap := provider.ZoneIDName{}

	zoneResponse, _, err := testProvider.getZones()
	assert.NoError(t, err)
	zones := append(zoneResponse.GetZones(), badZone)

	for _, zone := range zones {
		zoneIDNameMap.Add(zone.GetId(), zone.GetDomain())
	}

	err = testProvider.create(allTestEndpoints[7:9], &zones, &zoneIDNameMap)
	assert.NoError(t, err)

	testProvider.dryRun = true
	err = testProvider.create(allTestEndpoints[7:9], &zones, &zoneIDNameMap)
	assert.NoError(t, err)
	testProvider.dryRun = false
}

func TestCreateTarget(t *testing.T) {
	err := testProvider.createTarget("TEST_ZONE_ID1", "one.com", allTestEndpoints[8], "8.8.8.8")
	assert.NoError(t, err)

	testProvider.dryRun = true
	err = testProvider.createTarget("TEST_ZONE_ID1", "one.com", allTestEndpoints[8], "8.8.8.8")
	assert.Equal(t, fmt.Errorf("testing"), err)
	testProvider.dryRun = false
}

//p.deleteTarget(endpoint, domain, target, zoneID, recordID)

func TestDeleteTarget(t *testing.T) {
	err := testProvider.deleteTarget(allTestEndpoints[8], "one.com", allTestEndpoints[8].Targets[0], "TEST_ZONE_ID1", "TEST_ZONE_ZONE_RECORD_ID1")
	assert.NoError(t, err)

	testProvider.dryRun = true
	err = testProvider.deleteTarget(allTestEndpoints[8], "one.com", allTestEndpoints[8].Targets[0], "TEST_ZONE_ID1", "TEST_ZONE_ZONE_RECORD_ID1")
	assert.Equal(t, fmt.Errorf("testing"), err)
	testProvider.dryRun = false
}
