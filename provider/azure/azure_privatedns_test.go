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

	azcoreruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	privatedns "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/privatedns/armprivatedns"
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
	pagingHandler azcoreruntime.PagingHandler[privatedns.PrivateZonesClientListByResourceGroupResponse]
}

func newMockPrivateZonesClient(zones []*privatedns.PrivateZone) mockPrivateZonesClient {
	pagingHandler := azcoreruntime.PagingHandler[privatedns.PrivateZonesClientListByResourceGroupResponse]{
		More: func(resp privatedns.PrivateZonesClientListByResourceGroupResponse) bool {
			return false
		},
		Fetcher: func(context.Context, *privatedns.PrivateZonesClientListByResourceGroupResponse) (privatedns.PrivateZonesClientListByResourceGroupResponse, error) {
			return privatedns.PrivateZonesClientListByResourceGroupResponse{
				PrivateZoneListResult: privatedns.PrivateZoneListResult{
					Value: zones,
				},
			}, nil
		},
	}
	return mockPrivateZonesClient{
		pagingHandler: pagingHandler,
	}
}

func (client *mockPrivateZonesClient) NewListByResourceGroupPager(resourceGroupName string, options *privatedns.PrivateZonesClientListByResourceGroupOptions) *azcoreruntime.Pager[privatedns.PrivateZonesClientListByResourceGroupResponse] {
	return azcoreruntime.NewPager(client.pagingHandler)
}

// mockPrivateRecordSetsClient implements the methods of the Azure Private DNS RecordSet Client which are used in the Azure Private DNS Provider
// and returns static results which are defined per test
type mockPrivateRecordSetsClient struct {
	pagingHandler    azcoreruntime.PagingHandler[privatedns.RecordSetsClientListResponse]
	deletedEndpoints []*endpoint.Endpoint
	updatedEndpoints []*endpoint.Endpoint
}

func newMockPrivateRecordSectsClient(recordSets []*privatedns.RecordSet) mockPrivateRecordSetsClient {
	pagingHandler := azcoreruntime.PagingHandler[privatedns.RecordSetsClientListResponse]{
		More: func(resp privatedns.RecordSetsClientListResponse) bool {
			return false
		},
		Fetcher: func(context.Context, *privatedns.RecordSetsClientListResponse) (privatedns.RecordSetsClientListResponse, error) {
			return privatedns.RecordSetsClientListResponse{
				RecordSetListResult: privatedns.RecordSetListResult{
					Value: recordSets,
				},
			}, nil
		},
	}
	return mockPrivateRecordSetsClient{
		pagingHandler: pagingHandler,
	}
}

func (client *mockPrivateRecordSetsClient) NewListPager(resourceGroupName string, privateZoneName string, options *privatedns.RecordSetsClientListOptions) *azcoreruntime.Pager[privatedns.RecordSetsClientListResponse] {
	return azcoreruntime.NewPager(client.pagingHandler)
}

func (client *mockPrivateRecordSetsClient) Delete(ctx context.Context, resourceGroupName string, privateZoneName string, recordType privatedns.RecordType, relativeRecordSetName string, options *privatedns.RecordSetsClientDeleteOptions) (privatedns.RecordSetsClientDeleteResponse, error) {
	client.deletedEndpoints = append(
		client.deletedEndpoints,
		endpoint.NewEndpoint(
			formatAzureDNSName(relativeRecordSetName, privateZoneName),
			string(recordType),
			"",
		),
	)
	return privatedns.RecordSetsClientDeleteResponse{}, nil
}

func (client *mockPrivateRecordSetsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, privateZoneName string, recordType privatedns.RecordType, relativeRecordSetName string, parameters privatedns.RecordSet, options *privatedns.RecordSetsClientCreateOrUpdateOptions) (privatedns.RecordSetsClientCreateOrUpdateResponse, error) {
	var ttl endpoint.TTL
	if parameters.Properties.TTL != nil {
		ttl = endpoint.TTL(*parameters.Properties.TTL)
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
	return privatedns.RecordSetsClientCreateOrUpdateResponse{}, nil
	//return parameters, nil
}

func createMockPrivateZone(zone string, id string) *privatedns.PrivateZone {
	return &privatedns.PrivateZone{
		ID:   to.Ptr(id),
		Name: to.Ptr(zone),
	}
}

func privateARecordSetPropertiesGetter(values []string, ttl int64) *privatedns.RecordSetProperties {
	aRecords := make([]*privatedns.ARecord, len(values))
	for i, value := range values {
		aRecords[i] = &privatedns.ARecord{
			IPv4Address: to.Ptr(value),
		}
	}
	return &privatedns.RecordSetProperties{
		TTL:      to.Ptr(ttl),
		ARecords: aRecords,
	}
}

func privateCNameRecordSetPropertiesGetter(values []string, ttl int64) *privatedns.RecordSetProperties {
	return &privatedns.RecordSetProperties{
		TTL: to.Ptr(ttl),
		CnameRecord: &privatedns.CnameRecord{
			Cname: to.Ptr(values[0]),
		},
	}
}

func privateMXRecordSetPropertiesGetter(values []string, ttl int64) *privatedns.RecordSetProperties {
	mxRecords := make([]*privatedns.MxRecord, len(values))
	for i, target := range values {
		mxRecord, _ := parseMxTarget[privatedns.MxRecord](target)
		mxRecords[i] = &mxRecord
	}
	return &privatedns.RecordSetProperties{
		TTL:       to.Ptr(ttl),
		MxRecords: mxRecords,
	}
}

func privateTxtRecordSetPropertiesGetter(values []string, ttl int64) *privatedns.RecordSetProperties {
	return &privatedns.RecordSetProperties{
		TTL: to.Ptr(ttl),
		TxtRecords: []*privatedns.TxtRecord{
			{
				Value: []*string{&values[0]},
			},
		},
	}
}

func privateOthersRecordSetPropertiesGetter(values []string, ttl int64) *privatedns.RecordSetProperties {
	return &privatedns.RecordSetProperties{
		TTL: to.Ptr(ttl),
	}
}

func createPrivateMockRecordSet(name, recordType string, values ...string) *privatedns.RecordSet {
	return createPrivateMockRecordSetMultiWithTTL(name, recordType, 0, values...)
}

func createPrivateMockRecordSetWithTTL(name, recordType, value string, ttl int64) *privatedns.RecordSet {
	return createPrivateMockRecordSetMultiWithTTL(name, recordType, ttl, value)
}

func createPrivateMockRecordSetMultiWithTTL(name, recordType string, ttl int64, values ...string) *privatedns.RecordSet {
	var getterFunc func(values []string, ttl int64) *privatedns.RecordSetProperties

	switch recordType {
	case endpoint.RecordTypeA:
		getterFunc = privateARecordSetPropertiesGetter
	case endpoint.RecordTypeCNAME:
		getterFunc = privateCNameRecordSetPropertiesGetter
	case endpoint.RecordTypeMX:
		getterFunc = privateMXRecordSetPropertiesGetter
	case endpoint.RecordTypeTXT:
		getterFunc = privateTxtRecordSetPropertiesGetter
	default:
		getterFunc = privateOthersRecordSetPropertiesGetter
	}
	return &privatedns.RecordSet{
		Name:       to.Ptr(name),
		Type:       to.Ptr("Microsoft.Network/privateDnsZones/" + recordType),
		Properties: getterFunc(values, ttl),
	}
}

// newMockedAzurePrivateDNSProvider creates an AzureProvider comprising the mocked clients for zones and recordsets
func newMockedAzurePrivateDNSProvider(domainFilter endpoint.DomainFilter, zoneIDFilter provider.ZoneIDFilter, dryRun bool, resourceGroup string, zones []*privatedns.PrivateZone, recordSets []*privatedns.RecordSet) (*AzurePrivateDNSProvider, error) {
	zonesClient := newMockPrivateZonesClient(zones)
	recordSetsClient := newMockPrivateRecordSectsClient(recordSets)
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
		[]*privatedns.PrivateZone{
			createMockPrivateZone("example.com", "/privateDnsZones/example.com"),
		},
		[]*privatedns.RecordSet{
			createPrivateMockRecordSet("@", "NS", "ns1-03.azure-dns.com."),
			createPrivateMockRecordSet("@", "SOA", "Email: azuredns-hostmaster.microsoft.com"),
			createPrivateMockRecordSet("@", endpoint.RecordTypeA, "123.123.123.122"),
			createPrivateMockRecordSet("@", endpoint.RecordTypeTXT, "heritage=external-dns,external-dns/owner=default"),
			createPrivateMockRecordSetWithTTL("nginx", endpoint.RecordTypeA, "123.123.123.123", 3600),
			createPrivateMockRecordSetWithTTL("nginx", endpoint.RecordTypeTXT, "heritage=external-dns,external-dns/owner=default", recordTTL),
			createPrivateMockRecordSetWithTTL("hack", endpoint.RecordTypeCNAME, "hack.azurewebsites.net", 10),
			createPrivateMockRecordSetWithTTL("mail", endpoint.RecordTypeMX, "10 example.com", 4000),
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
		endpoint.NewEndpointWithTTL("mail.example.com", endpoint.RecordTypeMX, 4000, "10 example.com"),
	}

	validateAzureEndpoints(t, actual, expected)
}

func TestAzurePrivateDNSMultiRecord(t *testing.T) {
	provider, err := newMockedAzurePrivateDNSProvider(endpoint.NewDomainFilter([]string{"example.com"}), provider.NewZoneIDFilter([]string{""}), true, "k8s",
		[]*privatedns.PrivateZone{
			createMockPrivateZone("example.com", "/privateDnsZones/example.com"),
		},
		[]*privatedns.RecordSet{
			createPrivateMockRecordSet("@", "NS", "ns1-03.azure-dns.com."),
			createPrivateMockRecordSet("@", "SOA", "Email: azuredns-hostmaster.microsoft.com"),
			createPrivateMockRecordSet("@", endpoint.RecordTypeA, "123.123.123.122", "234.234.234.233"),
			createPrivateMockRecordSet("@", endpoint.RecordTypeTXT, "heritage=external-dns,external-dns/owner=default"),
			createPrivateMockRecordSetMultiWithTTL("nginx", endpoint.RecordTypeA, 3600, "123.123.123.123", "234.234.234.234"),
			createPrivateMockRecordSetWithTTL("nginx", endpoint.RecordTypeTXT, "heritage=external-dns,external-dns/owner=default", recordTTL),
			createPrivateMockRecordSetWithTTL("hack", endpoint.RecordTypeCNAME, "hack.azurewebsites.net", 10),
			createPrivateMockRecordSetMultiWithTTL("mail", endpoint.RecordTypeMX, 4000, "10 example.com", "20 backup.example.com"),
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
		endpoint.NewEndpointWithTTL("mail.example.com", endpoint.RecordTypeMX, 4000, "10 example.com", "20 backup.example.com"),
	}

	validateAzureEndpoints(t, actual, expected)
}

func TestAzurePrivateDNSApplyChanges(t *testing.T) {
	recordsClient := mockPrivateRecordSetsClient{}

	testAzurePrivateDNSApplyChangesInternal(t, false, &recordsClient)

	validateAzureEndpoints(t, recordsClient.deletedEndpoints, []*endpoint.Endpoint{
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
		endpoint.NewEndpointWithTTL("newmail.example.com", endpoint.RecordTypeMX, 7200, "40 bar.other.com"),
		endpoint.NewEndpointWithTTL("mail.example.com", endpoint.RecordTypeMX, endpoint.TTL(recordTTL), "10 other.com"),
		endpoint.NewEndpointWithTTL("mail.example.com", endpoint.RecordTypeTXT, endpoint.TTL(recordTTL), "tag"),
	})
}

func TestAzurePrivateDNSApplyChangesDryRun(t *testing.T) {
	recordsClient := mockRecordSetsClient{}

	testAzureApplyChangesInternal(t, true, &recordsClient)

	validateAzureEndpoints(t, recordsClient.deletedEndpoints, []*endpoint.Endpoint{})

	validateAzureEndpoints(t, recordsClient.updatedEndpoints, []*endpoint.Endpoint{})
}

func testAzurePrivateDNSApplyChangesInternal(t *testing.T, dryRun bool, client PrivateRecordSetsClient) {
	zones := []*privatedns.PrivateZone{
		createMockPrivateZone("example.com", "/privateDnsZones/example.com"),
		createMockPrivateZone("other.com", "/privateDnsZones/other.com"),
	}
	zonesClient := newMockPrivateZonesClient(zones)

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
		endpoint.NewEndpoint("mail.example.com", endpoint.RecordTypeMX, "10 other.com"),
		endpoint.NewEndpoint("mail.example.com", endpoint.RecordTypeTXT, "tag"),
	}

	currentRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("old.example.com", endpoint.RecordTypeA, "121.212.121.212"),
		endpoint.NewEndpoint("oldcname.example.com", endpoint.RecordTypeCNAME, "other.com"),
		endpoint.NewEndpoint("old.nope.com", endpoint.RecordTypeA, "121.212.121.212"),
		endpoint.NewEndpoint("oldmail.example.com", endpoint.RecordTypeMX, "20 foo.other.com"),
	}
	updatedRecords := []*endpoint.Endpoint{
		endpoint.NewEndpointWithTTL("new.example.com", endpoint.RecordTypeA, 3600, "111.222.111.222"),
		endpoint.NewEndpointWithTTL("newcname.example.com", endpoint.RecordTypeCNAME, 10, "other.com"),
		endpoint.NewEndpoint("new.nope.com", endpoint.RecordTypeA, "222.111.222.111"),
		endpoint.NewEndpointWithTTL("newmail.example.com", endpoint.RecordTypeMX, 7200, "40 bar.other.com"),
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
