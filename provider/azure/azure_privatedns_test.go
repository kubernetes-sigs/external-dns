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

package azure

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/privatedns/mgmt/2018-09-01/privatedns"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

const (
	recordTTL = 300
)

// mockPrivateZonesClient implements the methods of the Azure Private DNS Zones Client which are used in the Azure Private DNS Provider
// and returns static results which are defined per test
type mockPrivateZonesClient struct {
	mockZonesClientIterator *privatedns.PrivateZoneListResultIterator
}

// mockPrivateRecordSetsClient implements the methods of the Azure Private DNS RecordSet Client which are used in the Azure Private DNS Provider
// and returns static results which are defined per test
type mockPrivateRecordSetsClient struct {
	mockRecordSetListIterator *privatedns.RecordSetListResultIterator
	deletedEndpoints          []*endpoint.Endpoint
	updatedEndpoints          []*endpoint.Endpoint
}

// mockPrivateZoneListResultPageIterator is used to paginate forward through a list of zones
type mockPrivateZoneListResultPageIterator struct {
	offset  int
	results []privatedns.PrivateZoneListResult
}

// getNextPage provides the next page based on the offset of the mockZoneListResultPageIterator
func (m *mockPrivateZoneListResultPageIterator) getNextPage(context.Context, privatedns.PrivateZoneListResult) (privatedns.PrivateZoneListResult, error) {
	// it assumed that instances of this kind of iterator are only skimmed through once per test
	// otherwise a real implementation is required, e.g. based on a linked list
	if m.offset < len(m.results) {
		m.offset++
		return m.results[m.offset-1], nil
	}

	// paged to last page or empty
	return privatedns.PrivateZoneListResult{}, nil
}

// mockPrivateRecordSetListResultPageIterator is used to paginate forward through a list of recordsets
type mockPrivateRecordSetListResultPageIterator struct {
	offset  int
	results []privatedns.RecordSetListResult
}

// getNextPage provides the next page based on the offset of the mockRecordSetListResultPageIterator
func (m *mockPrivateRecordSetListResultPageIterator) getNextPage(context.Context, privatedns.RecordSetListResult) (privatedns.RecordSetListResult, error) {
	// it assumed that instances of this kind of iterator are only skimmed through once per test
	// otherwise a real implementation is required, e.g. based on a linked list
	if m.offset < len(m.results) {
		m.offset++
		return m.results[m.offset-1], nil
	}

	// paged to last page or empty
	return privatedns.RecordSetListResult{}, nil
}

func createMockPrivateZone(zone string, id string) privatedns.PrivateZone {
	return privatedns.PrivateZone{
		ID:   to.StringPtr(id),
		Name: to.StringPtr(zone),
	}
}

func (client *mockPrivateZonesClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, top *int32) (result privatedns.PrivateZoneListResultIterator, err error) {
	// pre-iterate to first item to emulate behaviour of Azure SDK
	err = client.mockZonesClientIterator.NextWithContext(ctx)
	if err != nil {
		return *client.mockZonesClientIterator, err
	}

	return *client.mockZonesClientIterator, nil
}

func privateARecordSetPropertiesGetter(values []string, ttl int64) *privatedns.RecordSetProperties {
	aRecords := make([]privatedns.ARecord, len(values))
	for i, value := range values {
		aRecords[i] = privatedns.ARecord{
			Ipv4Address: to.StringPtr(value),
		}
	}
	return &privatedns.RecordSetProperties{
		TTL:      to.Int64Ptr(ttl),
		ARecords: &aRecords,
	}
}

func privateCNameRecordSetPropertiesGetter(values []string, ttl int64) *privatedns.RecordSetProperties {
	return &privatedns.RecordSetProperties{
		TTL: to.Int64Ptr(ttl),
		CnameRecord: &privatedns.CnameRecord{
			Cname: to.StringPtr(values[0]),
		},
	}
}

func privateTxtRecordSetPropertiesGetter(values []string, ttl int64) *privatedns.RecordSetProperties {
	return &privatedns.RecordSetProperties{
		TTL: to.Int64Ptr(ttl),
		TxtRecords: &[]privatedns.TxtRecord{
			{
				Value: &[]string{values[0]},
			},
		},
	}
}

func privateOthersRecordSetPropertiesGetter(values []string, ttl int64) *privatedns.RecordSetProperties {
	return &privatedns.RecordSetProperties{
		TTL: to.Int64Ptr(ttl),
	}
}
func createPrivateMockRecordSet(name, recordType string, values ...string) privatedns.RecordSet {
	return createPrivateMockRecordSetMultiWithTTL(name, recordType, 0, values...)
}
func createPrivateMockRecordSetWithTTL(name, recordType, value string, ttl int64) privatedns.RecordSet {
	return createPrivateMockRecordSetMultiWithTTL(name, recordType, ttl, value)
}
func createPrivateMockRecordSetMultiWithTTL(name, recordType string, ttl int64, values ...string) privatedns.RecordSet {
	var getterFunc func(values []string, ttl int64) *privatedns.RecordSetProperties

	switch recordType {
	case endpoint.RecordTypeA:
		getterFunc = privateARecordSetPropertiesGetter
	case endpoint.RecordTypeCNAME:
		getterFunc = privateCNameRecordSetPropertiesGetter
	case endpoint.RecordTypeTXT:
		getterFunc = privateTxtRecordSetPropertiesGetter
	default:
		getterFunc = privateOthersRecordSetPropertiesGetter
	}
	return privatedns.RecordSet{
		Name:                to.StringPtr(name),
		Type:                to.StringPtr("Microsoft.Network/privateDnsZones/" + recordType),
		RecordSetProperties: getterFunc(values, ttl),
	}

}

func (client *mockPrivateRecordSetsClient) ListComplete(ctx context.Context, resourceGroupName string, zoneName string, top *int32, recordSetNameSuffix string) (result privatedns.RecordSetListResultIterator, err error) {
	// pre-iterate to first item to emulate behaviour of Azure SDK
	err = client.mockRecordSetListIterator.NextWithContext(ctx)
	if err != nil {
		return *client.mockRecordSetListIterator, err
	}

	return *client.mockRecordSetListIterator, nil
}

func (client *mockPrivateRecordSetsClient) Delete(ctx context.Context, resourceGroupName string, privateZoneName string, recordType privatedns.RecordType, relativeRecordSetName string, ifMatch string) (result autorest.Response, err error) {
	client.deletedEndpoints = append(
		client.deletedEndpoints,
		endpoint.NewEndpoint(
			formatAzureDNSName(relativeRecordSetName, privateZoneName),
			string(recordType),
			"",
		),
	)
	return autorest.Response{}, nil
}

func (client *mockPrivateRecordSetsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, privateZoneName string, recordType privatedns.RecordType, relativeRecordSetName string, parameters privatedns.RecordSet, ifMatch string, ifNoneMatch string) (result privatedns.RecordSet, err error) {
	var ttl endpoint.TTL
	if parameters.TTL != nil {
		ttl = endpoint.TTL(*parameters.TTL)
	}
	client.updatedEndpoints = append(
		client.updatedEndpoints,
		endpoint.NewEndpointWithTTL(
			formatAzureDNSName(relativeRecordSetName, privateZoneName),
			string(recordType),
			ttl,
			extractAzurePrivateDNSTargets(&parameters)...,
		),
	)
	return parameters, nil
}

// newMockedAzurePrivateDNSProvider creates an AzureProvider comprising the mocked clients for zones and recordsets
func newMockedAzurePrivateDNSProvider(domainFilter endpoint.DomainFilter, zoneIDFilter provider.ZoneIDFilter, dryRun bool, resourceGroup string, zones *[]privatedns.PrivateZone, recordSets *[]privatedns.RecordSet) (*AzurePrivateDNSProvider, error) {
	// init zone-related parts of the mock-client
	pageIterator := mockPrivateZoneListResultPageIterator{
		results: []privatedns.PrivateZoneListResult{
			{
				Value: zones,
			},
		},
	}

	mockZoneListResultPage := privatedns.NewPrivateZoneListResultPage(privatedns.PrivateZoneListResult{}, pageIterator.getNextPage)
	mockZoneClientIterator := privatedns.NewPrivateZoneListResultIterator(mockZoneListResultPage)
	zonesClient := mockPrivateZonesClient{
		mockZonesClientIterator: &mockZoneClientIterator,
	}

	// init record-related parts of the mock-client
	resultPageIterator := mockPrivateRecordSetListResultPageIterator{
		results: []privatedns.RecordSetListResult{
			{
				Value: recordSets,
			},
		},
	}

	mockRecordSetListResultPage := privatedns.NewRecordSetListResultPage(privatedns.RecordSetListResult{}, resultPageIterator.getNextPage)
	mockRecordSetListIterator := privatedns.NewRecordSetListResultIterator(mockRecordSetListResultPage)
	recordSetsClient := mockPrivateRecordSetsClient{
		mockRecordSetListIterator: &mockRecordSetListIterator,
	}

	return newAzurePrivateDNSProvider(domainFilter, zoneIDFilter, dryRun, resourceGroup, &zonesClient, &recordSetsClient), nil
}

func newAzurePrivateDNSProvider(domainFilter endpoint.DomainFilter, zoneIDFilter provider.ZoneIDFilter, dryRun bool, resourceGroup string, privateZonesClient PrivateZonesClient, privateRecordsClient PrivateRecordSetsClient) *AzurePrivateDNSProvider {
	return &AzurePrivateDNSProvider{
		domainFilter:     domainFilter,
		zoneIDFilter:     zoneIDFilter,
		dryRun:           dryRun,
		resourceGroup:    resourceGroup,
		zonesClient:      privateZonesClient,
		recordSetsClient: privateRecordsClient,
	}
}

func TestAzurePrivateDNSRecord(t *testing.T) {
	provider, err := newMockedAzurePrivateDNSProvider(endpoint.NewDomainFilter([]string{"example.com"}), provider.NewZoneIDFilter([]string{""}), true, "k8s",
		&[]privatedns.PrivateZone{
			createMockPrivateZone("example.com", "/privateDnsZones/example.com"),
		},
		&[]privatedns.RecordSet{
			createPrivateMockRecordSet("@", "NS", "ns1-03.azure-dns.com."),
			createPrivateMockRecordSet("@", "SOA", "Email: azuredns-hostmaster.microsoft.com"),
			createPrivateMockRecordSet("@", endpoint.RecordTypeA, "123.123.123.122"),
			createPrivateMockRecordSet("@", endpoint.RecordTypeTXT, "heritage=external-dns,external-dns/owner=default"),
			createPrivateMockRecordSetWithTTL("nginx", endpoint.RecordTypeA, "123.123.123.123", 3600),
			createPrivateMockRecordSetWithTTL("nginx", endpoint.RecordTypeTXT, "heritage=external-dns,external-dns/owner=default", recordTTL),
			createPrivateMockRecordSetWithTTL("hack", endpoint.RecordTypeCNAME, "hack.azurewebsites.net", 10),
		})

	if err != nil {
		t.Fatal(err)
	}

	actual, err := provider.Records(context.Background())

	if err != nil {
		t.Fatal(err)
	}
	expected := []*endpoint.Endpoint{
		endpoint.NewEndpoint("example.com", endpoint.RecordTypeA, "123.123.123.122"),
		endpoint.NewEndpoint("example.com", endpoint.RecordTypeTXT, "heritage=external-dns,external-dns/owner=default"),
		endpoint.NewEndpointWithTTL("nginx.example.com", endpoint.RecordTypeA, 3600, "123.123.123.123"),
		endpoint.NewEndpointWithTTL("nginx.example.com", endpoint.RecordTypeTXT, recordTTL, "heritage=external-dns,external-dns/owner=default"),
		endpoint.NewEndpointWithTTL("hack.example.com", endpoint.RecordTypeCNAME, 10, "hack.azurewebsites.net"),
	}

	validateAzureEndpoints(t, actual, expected)

}

func TestAzurePrivateDNSMultiRecord(t *testing.T) {
	provider, err := newMockedAzurePrivateDNSProvider(endpoint.NewDomainFilter([]string{"example.com"}), provider.NewZoneIDFilter([]string{""}), true, "k8s",
		&[]privatedns.PrivateZone{
			createMockPrivateZone("example.com", "/privateDnsZones/example.com"),
		},
		&[]privatedns.RecordSet{
			createPrivateMockRecordSet("@", "NS", "ns1-03.azure-dns.com."),
			createPrivateMockRecordSet("@", "SOA", "Email: azuredns-hostmaster.microsoft.com"),
			createPrivateMockRecordSet("@", endpoint.RecordTypeA, "123.123.123.122", "234.234.234.233"),
			createPrivateMockRecordSet("@", endpoint.RecordTypeTXT, "heritage=external-dns,external-dns/owner=default"),
			createPrivateMockRecordSetMultiWithTTL("nginx", endpoint.RecordTypeA, 3600, "123.123.123.123", "234.234.234.234"),
			createPrivateMockRecordSetWithTTL("nginx", endpoint.RecordTypeTXT, "heritage=external-dns,external-dns/owner=default", recordTTL),
			createPrivateMockRecordSetWithTTL("hack", endpoint.RecordTypeCNAME, "hack.azurewebsites.net", 10),
		})

	if err != nil {
		t.Fatal(err)
	}

	actual, err := provider.Records(context.Background())

	if err != nil {
		t.Fatal(err)
	}
	expected := []*endpoint.Endpoint{
		endpoint.NewEndpoint("example.com", endpoint.RecordTypeA, "123.123.123.122", "234.234.234.233"),
		endpoint.NewEndpoint("example.com", endpoint.RecordTypeTXT, "heritage=external-dns,external-dns/owner=default"),
		endpoint.NewEndpointWithTTL("nginx.example.com", endpoint.RecordTypeA, 3600, "123.123.123.123", "234.234.234.234"),
		endpoint.NewEndpointWithTTL("nginx.example.com", endpoint.RecordTypeTXT, recordTTL, "heritage=external-dns,external-dns/owner=default"),
		endpoint.NewEndpointWithTTL("hack.example.com", endpoint.RecordTypeCNAME, 10, "hack.azurewebsites.net"),
	}

	validateAzureEndpoints(t, actual, expected)

}

func TestAzurePrivateDNSApplyChanges(t *testing.T) {
	recordsClient := mockPrivateRecordSetsClient{}

	testAzurePrivateDNSApplyChangesInternal(t, false, &recordsClient)

	validateAzureEndpoints(t, recordsClient.deletedEndpoints, []*endpoint.Endpoint{
		endpoint.NewEndpoint("old.example.com", endpoint.RecordTypeA, ""),
		endpoint.NewEndpoint("oldcname.example.com", endpoint.RecordTypeCNAME, ""),
		endpoint.NewEndpoint("deleted.example.com", endpoint.RecordTypeA, ""),
		endpoint.NewEndpoint("deletedcname.example.com", endpoint.RecordTypeCNAME, ""),
	})

	validateAzureEndpoints(t, recordsClient.updatedEndpoints, []*endpoint.Endpoint{
		endpoint.NewEndpointWithTTL("example.com", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "1.2.3.4"),
		endpoint.NewEndpointWithTTL("example.com", endpoint.RecordTypeTXT, endpoint.TTL(recordTTL), "tag"),
		endpoint.NewEndpointWithTTL("foo.example.com", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "1.2.3.4", "1.2.3.5"),
		endpoint.NewEndpointWithTTL("foo.example.com", endpoint.RecordTypeTXT, endpoint.TTL(recordTTL), "tag"),
		endpoint.NewEndpointWithTTL("bar.example.com", endpoint.RecordTypeCNAME, endpoint.TTL(recordTTL), "other.com"),
		endpoint.NewEndpointWithTTL("bar.example.com", endpoint.RecordTypeTXT, endpoint.TTL(recordTTL), "tag"),
		endpoint.NewEndpointWithTTL("other.com", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "5.6.7.8"),
		endpoint.NewEndpointWithTTL("other.com", endpoint.RecordTypeTXT, endpoint.TTL(recordTTL), "tag"),
		endpoint.NewEndpointWithTTL("new.example.com", endpoint.RecordTypeA, 3600, "111.222.111.222"),
		endpoint.NewEndpointWithTTL("newcname.example.com", endpoint.RecordTypeCNAME, 10, "other.com"),
	})
}

func TestAzurePrivateDNSApplyChangesDryRun(t *testing.T) {
	recordsClient := mockRecordSetsClient{}

	testAzureApplyChangesInternal(t, true, &recordsClient)

	validateAzureEndpoints(t, recordsClient.deletedEndpoints, []*endpoint.Endpoint{})

	validateAzureEndpoints(t, recordsClient.updatedEndpoints, []*endpoint.Endpoint{})
}

func testAzurePrivateDNSApplyChangesInternal(t *testing.T, dryRun bool, client PrivateRecordSetsClient) {
	zlr := privatedns.PrivateZoneListResult{
		Value: &[]privatedns.PrivateZone{
			createMockPrivateZone("example.com", "/privateDnsZones/example.com"),
			createMockPrivateZone("other.com", "/privateDnsZones/other.com"),
		},
	}

	results := []privatedns.PrivateZoneListResult{
		zlr,
	}

	mockZoneListResultPage := privatedns.NewPrivateZoneListResultPage(privatedns.PrivateZoneListResult{}, func(ctxParam context.Context, zlrParam privatedns.PrivateZoneListResult) (privatedns.PrivateZoneListResult, error) {
		if len(results) > 0 {
			result := results[0]
			results = nil
			return result, nil
		}
		return privatedns.PrivateZoneListResult{}, nil
	})
	mockZoneClientIterator := privatedns.NewPrivateZoneListResultIterator(mockZoneListResultPage)

	zonesClient := mockPrivateZonesClient{
		mockZonesClientIterator: &mockZoneClientIterator,
	}

	provider := newAzurePrivateDNSProvider(
		endpoint.NewDomainFilter([]string{""}),
		provider.NewZoneIDFilter([]string{""}),
		dryRun,
		"group",
		&zonesClient,
		client,
	)

	createRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("example.com", endpoint.RecordTypeA, "1.2.3.4"),
		endpoint.NewEndpoint("example.com", endpoint.RecordTypeTXT, "tag"),
		endpoint.NewEndpoint("foo.example.com", endpoint.RecordTypeA, "1.2.3.5", "1.2.3.4"),
		endpoint.NewEndpoint("foo.example.com", endpoint.RecordTypeTXT, "tag"),
		endpoint.NewEndpoint("bar.example.com", endpoint.RecordTypeCNAME, "other.com"),
		endpoint.NewEndpoint("bar.example.com", endpoint.RecordTypeTXT, "tag"),
		endpoint.NewEndpoint("other.com", endpoint.RecordTypeA, "5.6.7.8"),
		endpoint.NewEndpoint("other.com", endpoint.RecordTypeTXT, "tag"),
		endpoint.NewEndpoint("nope.com", endpoint.RecordTypeA, "4.4.4.4"),
		endpoint.NewEndpoint("nope.com", endpoint.RecordTypeTXT, "tag"),
	}

	currentRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("old.example.com", endpoint.RecordTypeA, "121.212.121.212"),
		endpoint.NewEndpoint("oldcname.example.com", endpoint.RecordTypeCNAME, "other.com"),
		endpoint.NewEndpoint("old.nope.com", endpoint.RecordTypeA, "121.212.121.212"),
	}
	updatedRecords := []*endpoint.Endpoint{
		endpoint.NewEndpointWithTTL("new.example.com", endpoint.RecordTypeA, 3600, "111.222.111.222"),
		endpoint.NewEndpointWithTTL("newcname.example.com", endpoint.RecordTypeCNAME, 10, "other.com"),
		endpoint.NewEndpoint("new.nope.com", endpoint.RecordTypeA, "222.111.222.111"),
	}

	deleteRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("deleted.example.com", endpoint.RecordTypeA, "111.222.111.222"),
		endpoint.NewEndpoint("deletedcname.example.com", endpoint.RecordTypeCNAME, "other.com"),
		endpoint.NewEndpoint("deleted.nope.com", endpoint.RecordTypeA, "222.111.222.111"),
	}

	changes := &plan.Changes{
		Create:    createRecords,
		UpdateNew: updatedRecords,
		UpdateOld: currentRecords,
		Delete:    deleteRecords,
	}

	if err := provider.ApplyChanges(context.Background(), changes); err != nil {
		t.Fatal(err)
	}
}
