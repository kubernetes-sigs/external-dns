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
	privatedns "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/privatedns/armprivatedns"

	"sigs.k8s.io/external-dns/provider/blueprint"

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
		More: func(_ privatedns.PrivateZonesClientListByResourceGroupResponse) bool {
			return false
		},
		Fetcher: func(context.Context, *privatedns.PrivateZonesClientListByResourceGroupResponse) (privatedns.PrivateZonesClientListByResourceGroupResponse, error) {
			return privatedns.PrivateZonesClientListByResourceGroupResponse{
				Value: zones,
			}, nil
		},
	}
	return mockPrivateZonesClient{
		pagingHandler: pagingHandler,
	}
}

func (client *mockPrivateZonesClient) NewListByResourceGroupPager(_ string, _ *privatedns.PrivateZonesClientListByResourceGroupOptions) *azcoreruntime.Pager[privatedns.PrivateZonesClientListByResourceGroupResponse] {
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
		More: func(_ privatedns.RecordSetsClientListResponse) bool {
			return false
		},
		Fetcher: func(context.Context, *privatedns.RecordSetsClientListResponse) (privatedns.RecordSetsClientListResponse, error) {
			return privatedns.RecordSetsClientListResponse{
				Value: recordSets,
			}, nil
		},
	}
	return mockPrivateRecordSetsClient{
		pagingHandler: pagingHandler,
	}
}

func (client *mockPrivateRecordSetsClient) NewListPager(_ string, _ string, _ *privatedns.RecordSetsClientListOptions) *azcoreruntime.Pager[privatedns.RecordSetsClientListResponse] {
	return azcoreruntime.NewPager(client.pagingHandler)
}

func (client *mockPrivateRecordSetsClient) Delete(_ context.Context, _ string, privateZoneName string, recordType privatedns.RecordType, relativeRecordSetName string, _ *privatedns.RecordSetsClientDeleteOptions) (privatedns.RecordSetsClientDeleteResponse, error) {
	ep, err := endpoint.NewEndpoint(
		formatAzureDNSName(relativeRecordSetName, privateZoneName),
		string(recordType),
		"",
	)
	if err != nil {
		return privatedns.RecordSetsClientDeleteResponse{}, err
	}
	client.deletedEndpoints = append(client.deletedEndpoints, ep)
	return privatedns.RecordSetsClientDeleteResponse{}, nil
}

func (client *mockPrivateRecordSetsClient) CreateOrUpdate(_ context.Context, _ string, privateZoneName string, recordType privatedns.RecordType, relativeRecordSetName string, parameters privatedns.RecordSet, _ *privatedns.RecordSetsClientCreateOrUpdateOptions) (privatedns.RecordSetsClientCreateOrUpdateResponse, error) {
	var ttl endpoint.TTL
	if parameters.Properties.TTL != nil {
		ttl = endpoint.TTL(*parameters.Properties.TTL)
	}
	ep, err := endpoint.NewEndpointWithTTL(
		formatAzureDNSName(relativeRecordSetName, privateZoneName),
		string(recordType),
		ttl,
		extractAzurePrivateDNSTargets(&parameters)...,
	)
	if err != nil {
		return privatedns.RecordSetsClientCreateOrUpdateResponse{}, err
	}
	client.updatedEndpoints = append(client.updatedEndpoints, ep)
	return privatedns.RecordSetsClientCreateOrUpdateResponse{}, nil
}

func createMockPrivateZone(zone string, id string) *privatedns.PrivateZone {
	return &privatedns.PrivateZone{
		ID:   new(id),
		Name: new(zone),
	}
}

func privateARecordSetPropertiesGetter(values []string, ttl int64) *privatedns.RecordSetProperties {
	aRecords := make([]*privatedns.ARecord, len(values))
	for i, value := range values {
		aRecords[i] = &privatedns.ARecord{
			IPv4Address: new(value),
		}
	}
	return &privatedns.RecordSetProperties{
		TTL:      new(ttl),
		ARecords: aRecords,
	}
}

func privateAAAARecordSetPropertiesGetter(values []string, ttl int64) *privatedns.RecordSetProperties {
	aaaaRecords := make([]*privatedns.AaaaRecord, len(values))
	for i, value := range values {
		aaaaRecords[i] = &privatedns.AaaaRecord{
			IPv6Address: new(value),
		}
	}
	return &privatedns.RecordSetProperties{
		TTL:         new(ttl),
		AaaaRecords: aaaaRecords,
	}
}

func privateCNameRecordSetPropertiesGetter(values []string, ttl int64) *privatedns.RecordSetProperties {
	return &privatedns.RecordSetProperties{
		TTL: new(ttl),
		CnameRecord: &privatedns.CnameRecord{
			Cname: new(values[0]),
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
		TTL:       new(ttl),
		MxRecords: mxRecords,
	}
}

func privateTxtRecordSetPropertiesGetter(values []string, ttl int64) *privatedns.RecordSetProperties {
	return &privatedns.RecordSetProperties{
		TTL: new(ttl),
		TxtRecords: []*privatedns.TxtRecord{
			{
				Value: []*string{&values[0]},
			},
		},
	}
}

func privateOthersRecordSetPropertiesGetter(_ []string, ttl int64) *privatedns.RecordSetProperties {
	return &privatedns.RecordSetProperties{
		TTL: new(ttl),
	}
}

func createPrivateMockRecordSet(recordType string, values ...string) *privatedns.RecordSet {
	return createPrivateMockRecordSetMultiWithTTL("@", recordType, 0, values...)
}

func createPrivateMockRecordSetWithNameAndTTL(name, recordType, value string, ttl int64) *privatedns.RecordSet {
	return createPrivateMockRecordSetMultiWithTTL(name, recordType, ttl, value)
}

func createPrivateMockRecordSetMultiWithTTL(name, recordType string, ttl int64, values ...string) *privatedns.RecordSet {
	var getterFunc func(values []string, ttl int64) *privatedns.RecordSetProperties

	switch recordType {
	case endpoint.RecordTypeA:
		getterFunc = privateARecordSetPropertiesGetter
	case endpoint.RecordTypeAAAA:
		getterFunc = privateAAAARecordSetPropertiesGetter
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
		Name:       new(name),
		Type:       new("Microsoft.Network/privateDnsZones/" + recordType),
		Properties: getterFunc(values, ttl),
	}
}

// newMockedAzurePrivateDNSProvider creates an AzureProvider comprising the mocked clients for zones and recordsets
func newMockedAzurePrivateDNSProvider(domainFilter *endpoint.DomainFilter, zoneNameFilter *endpoint.DomainFilter, zoneIDFilter provider.ZoneIDFilter, dryRun bool, resourceGroup string, zones []*privatedns.PrivateZone, recordSets []*privatedns.RecordSet, maxRetriesCount int) *AzurePrivateDNSProvider {
	zonesClient := newMockPrivateZonesClient(zones)
	recordSetsClient := newMockPrivateRecordSectsClient(recordSets)
	return newAzurePrivateDNSProvider(domainFilter, zoneNameFilter, zoneIDFilter, dryRun, resourceGroup, &zonesClient, &recordSetsClient, maxRetriesCount)
}

func newAzurePrivateDNSProvider(domainFilter *endpoint.DomainFilter, zoneNameFilter *endpoint.DomainFilter, zoneIDFilter provider.ZoneIDFilter, dryRun bool, resourceGroup string, privateZonesClient PrivateZonesClient, privateRecordsClient PrivateRecordSetsClient, maxRetriesCount int) *AzurePrivateDNSProvider {
	return &AzurePrivateDNSProvider{
		domainFilter:     domainFilter,
		zoneNameFilter:   zoneNameFilter,
		zoneIDFilter:     zoneIDFilter,
		dryRun:           dryRun,
		resourceGroup:    resourceGroup,
		zonesClient:      privateZonesClient,
		zonesCache:       blueprint.NewZoneCache[[]privatedns.PrivateZone](0),
		recordSetsClient: privateRecordsClient,
		maxRetriesCount:  maxRetriesCount,
	}
}

func TestAzurePrivateDNSRecord(t *testing.T) {
	provider := newMockedAzurePrivateDNSProvider(endpoint.NewDomainFilter([]string{"example.com"}), endpoint.NewDomainFilter([]string{}), provider.NewZoneIDFilter([]string{""}), true, "k8s",
		[]*privatedns.PrivateZone{
			createMockPrivateZone("example.com", "/privateDnsZones/example.com"),
		},
		[]*privatedns.RecordSet{
			createPrivateMockRecordSet("NS", "ns1-03.azure-dns.com."),
			createPrivateMockRecordSet("SOA", "Email: azuredns-hostmaster.microsoft.com"),
			createPrivateMockRecordSet(endpoint.RecordTypeA, "123.123.123.122"),
			createPrivateMockRecordSet(endpoint.RecordTypeAAAA, "2001::123:123:123:122"),
			createPrivateMockRecordSet(endpoint.RecordTypeTXT, "heritage=external-dns,external-dns/owner=default"),
			createPrivateMockRecordSetWithNameAndTTL("nginx", endpoint.RecordTypeA, "123.123.123.123", 3600),
			createPrivateMockRecordSetWithNameAndTTL("nginx", endpoint.RecordTypeAAAA, "2001::123:123:123:123", 3600),
			createPrivateMockRecordSetWithNameAndTTL("nginx", endpoint.RecordTypeTXT, "heritage=external-dns,external-dns/owner=default", recordTTL),
			createPrivateMockRecordSetWithNameAndTTL("hack", endpoint.RecordTypeCNAME, "hack.azurewebsites.net", 10),
			createPrivateMockRecordSetWithNameAndTTL("mail", endpoint.RecordTypeMX, "10 example.com", 4000),
		}, 3)

	actual, err := provider.Records(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	expected := []*endpoint.Endpoint{
		endpoint.MustNewEndpoint("example.com", endpoint.RecordTypeA, "123.123.123.122"),
		endpoint.MustNewEndpoint("example.com", endpoint.RecordTypeAAAA, "2001::123:123:123:122"),
		endpoint.MustNewEndpoint("example.com", endpoint.RecordTypeTXT, "heritage=external-dns,external-dns/owner=default"),
		endpoint.MustNewEndpointWithTTL("nginx.example.com", endpoint.RecordTypeA, 3600, "123.123.123.123"),
		endpoint.MustNewEndpointWithTTL("nginx.example.com", endpoint.RecordTypeAAAA, 3600, "2001::123:123:123:123"),
		endpoint.MustNewEndpointWithTTL("nginx.example.com", endpoint.RecordTypeTXT, recordTTL, "heritage=external-dns,external-dns/owner=default"),
		endpoint.MustNewEndpointWithTTL("hack.example.com", endpoint.RecordTypeCNAME, 10, "hack.azurewebsites.net"),
		endpoint.MustNewEndpointWithTTL("mail.example.com", endpoint.RecordTypeMX, 4000, "10 example.com"),
	}

	validateAzureEndpoints(t, actual, expected)
}

func TestAzurePrivateDNSMultiRecord(t *testing.T) {
	provider := newMockedAzurePrivateDNSProvider(endpoint.NewDomainFilter([]string{"example.com"}), endpoint.NewDomainFilter([]string{}), provider.NewZoneIDFilter([]string{""}), true, "k8s",
		[]*privatedns.PrivateZone{
			createMockPrivateZone("example.com", "/privateDnsZones/example.com"),
		},
		[]*privatedns.RecordSet{
			createPrivateMockRecordSet("NS", "ns1-03.azure-dns.com."),
			createPrivateMockRecordSet("SOA", "Email: azuredns-hostmaster.microsoft.com"),
			createPrivateMockRecordSet(endpoint.RecordTypeA, "123.123.123.122", "234.234.234.233"),
			createPrivateMockRecordSet(endpoint.RecordTypeAAAA, "2001::123:123:123:122", "2001::234:234:234:233"),
			createPrivateMockRecordSet(endpoint.RecordTypeTXT, "heritage=external-dns,external-dns/owner=default"),
			createPrivateMockRecordSetMultiWithTTL("nginx", endpoint.RecordTypeA, 3600, "123.123.123.123", "234.234.234.234"),
			createPrivateMockRecordSetMultiWithTTL("nginx", endpoint.RecordTypeAAAA, 3600, "2001::123:123:123:123", "2001::234:234:234:234"),
			createPrivateMockRecordSetWithNameAndTTL("nginx", endpoint.RecordTypeTXT, "heritage=external-dns,external-dns/owner=default", recordTTL),
			createPrivateMockRecordSetWithNameAndTTL("hack", endpoint.RecordTypeCNAME, "hack.azurewebsites.net", 10),
			createPrivateMockRecordSetMultiWithTTL("mail", endpoint.RecordTypeMX, 4000, "10 example.com", "20 backup.example.com"),
		}, 3)

	actual, err := provider.Records(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	expected := []*endpoint.Endpoint{
		endpoint.MustNewEndpoint("example.com", endpoint.RecordTypeA, "123.123.123.122", "234.234.234.233"),
		endpoint.MustNewEndpoint("example.com", endpoint.RecordTypeAAAA, "2001::123:123:123:122", "2001::234:234:234:233"),
		endpoint.MustNewEndpoint("example.com", endpoint.RecordTypeTXT, "heritage=external-dns,external-dns/owner=default"),
		endpoint.MustNewEndpointWithTTL("nginx.example.com", endpoint.RecordTypeA, 3600, "123.123.123.123", "234.234.234.234"),
		endpoint.MustNewEndpointWithTTL("nginx.example.com", endpoint.RecordTypeAAAA, 3600, "2001::123:123:123:123", "2001::234:234:234:234"),
		endpoint.MustNewEndpointWithTTL("nginx.example.com", endpoint.RecordTypeTXT, recordTTL, "heritage=external-dns,external-dns/owner=default"),
		endpoint.MustNewEndpointWithTTL("hack.example.com", endpoint.RecordTypeCNAME, 10, "hack.azurewebsites.net"),
		endpoint.MustNewEndpointWithTTL("mail.example.com", endpoint.RecordTypeMX, 4000, "10 example.com", "20 backup.example.com"),
	}

	validateAzureEndpoints(t, actual, expected)
}

func TestAzurePrivateDNSApplyChanges(t *testing.T) {
	recordsClient := mockPrivateRecordSetsClient{}

	testAzurePrivateDNSApplyChangesInternal(t, false, &recordsClient)

	validateAzureEndpoints(t, recordsClient.deletedEndpoints, []*endpoint.Endpoint{
		endpoint.MustNewEndpoint("deleted.example.com", endpoint.RecordTypeA, ""),
		endpoint.MustNewEndpoint("deletedaaaa.example.com", endpoint.RecordTypeAAAA, ""),
		endpoint.MustNewEndpoint("deletedcname.example.com", endpoint.RecordTypeCNAME, ""),
	})

	validateAzureEndpoints(t, recordsClient.updatedEndpoints, []*endpoint.Endpoint{
		endpoint.MustNewEndpointWithTTL("example.com", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "1.2.3.4"),
		endpoint.MustNewEndpointWithTTL("example.com", endpoint.RecordTypeAAAA, endpoint.TTL(recordTTL), "2001::1:2:3:4"),
		endpoint.MustNewEndpointWithTTL("example.com", endpoint.RecordTypeTXT, endpoint.TTL(recordTTL), "tag"),
		endpoint.MustNewEndpointWithTTL("foo.example.com", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "1.2.3.4", "1.2.3.5"),
		endpoint.MustNewEndpointWithTTL("foo.example.com", endpoint.RecordTypeAAAA, endpoint.TTL(recordTTL), "2001::1:2:3:4", "2001::1:2:3:5"),
		endpoint.MustNewEndpointWithTTL("foo.example.com", endpoint.RecordTypeTXT, endpoint.TTL(recordTTL), "tag"),
		endpoint.MustNewEndpointWithTTL("bar.example.com", endpoint.RecordTypeCNAME, endpoint.TTL(recordTTL), "other.com"),
		endpoint.MustNewEndpointWithTTL("bar.example.com", endpoint.RecordTypeTXT, endpoint.TTL(recordTTL), "tag"),
		endpoint.MustNewEndpointWithTTL("other.com", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "5.6.7.8"),
		endpoint.MustNewEndpointWithTTL("other.com", endpoint.RecordTypeAAAA, endpoint.TTL(recordTTL), "2001::5:6:7:8"),
		endpoint.MustNewEndpointWithTTL("other.com", endpoint.RecordTypeTXT, endpoint.TTL(recordTTL), "tag"),
		endpoint.MustNewEndpointWithTTL("new.example.com", endpoint.RecordTypeA, 3600, "111.222.111.222"),
		endpoint.MustNewEndpointWithTTL("new.example.com", endpoint.RecordTypeAAAA, 3600, "2001::111:222:111:222"),
		endpoint.MustNewEndpointWithTTL("newcname.example.com", endpoint.RecordTypeCNAME, 10, "other.com"),
		endpoint.MustNewEndpointWithTTL("newmail.example.com", endpoint.RecordTypeMX, 7200, "40 bar.other.com"),
		endpoint.MustNewEndpointWithTTL("mail.example.com", endpoint.RecordTypeMX, endpoint.TTL(recordTTL), "10 other.com"),
		endpoint.MustNewEndpointWithTTL("mail.example.com", endpoint.RecordTypeTXT, endpoint.TTL(recordTTL), "tag"),
	})
}

func TestAzurePrivateDNSApplyChangesDryRun(t *testing.T) {
	recordsClient := mockPrivateRecordSetsClient{}

	testAzurePrivateDNSApplyChangesInternal(t, true, &recordsClient)

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
		endpoint.NewDomainFilter([]string{""}),
		provider.NewZoneIDFilter([]string{""}),
		dryRun,
		"group",
		&zonesClient,
		client,
		3,
	)

	createRecords := []*endpoint.Endpoint{
		endpoint.MustNewEndpoint("example.com", endpoint.RecordTypeA, "1.2.3.4"),
		endpoint.MustNewEndpoint("example.com", endpoint.RecordTypeAAAA, "2001::1:2:3:4"),
		endpoint.MustNewEndpoint("example.com", endpoint.RecordTypeTXT, "tag"),
		endpoint.MustNewEndpoint("foo.example.com", endpoint.RecordTypeA, "1.2.3.5", "1.2.3.4"),
		endpoint.MustNewEndpoint("foo.example.com", endpoint.RecordTypeAAAA, "2001::1:2:3:5", "2001::1:2:3:4"),
		endpoint.MustNewEndpoint("foo.example.com", endpoint.RecordTypeTXT, "tag"),
		endpoint.MustNewEndpoint("bar.example.com", endpoint.RecordTypeCNAME, "other.com"),
		endpoint.MustNewEndpoint("bar.example.com", endpoint.RecordTypeTXT, "tag"),
		endpoint.MustNewEndpoint("other.com", endpoint.RecordTypeA, "5.6.7.8"),
		endpoint.MustNewEndpoint("other.com", endpoint.RecordTypeAAAA, "2001::5:6:7:8"),
		endpoint.MustNewEndpoint("other.com", endpoint.RecordTypeTXT, "tag"),
		endpoint.MustNewEndpoint("nope.com", endpoint.RecordTypeA, "4.4.4.4"),
		endpoint.MustNewEndpoint("nope.com", endpoint.RecordTypeAAAA, "2001::4:4:4:4"),
		endpoint.MustNewEndpoint("nope.com", endpoint.RecordTypeTXT, "tag"),
		endpoint.MustNewEndpoint("mail.example.com", endpoint.RecordTypeMX, "10 other.com"),
		endpoint.MustNewEndpoint("mail.example.com", endpoint.RecordTypeTXT, "tag"),
	}

	currentRecords := []*endpoint.Endpoint{
		endpoint.MustNewEndpoint("old.example.com", endpoint.RecordTypeA, "121.212.121.212"),
		endpoint.MustNewEndpoint("oldcname.example.com", endpoint.RecordTypeCNAME, "other.com"),
		endpoint.MustNewEndpoint("old.nope.com", endpoint.RecordTypeA, "121.212.121.212"),
		endpoint.MustNewEndpoint("oldmail.example.com", endpoint.RecordTypeMX, "20 foo.other.com"),
	}
	updatedRecords := []*endpoint.Endpoint{
		endpoint.MustNewEndpointWithTTL("new.example.com", endpoint.RecordTypeA, 3600, "111.222.111.222"),
		endpoint.MustNewEndpointWithTTL("new.example.com", endpoint.RecordTypeAAAA, 3600, "2001::111:222:111:222"),
		endpoint.MustNewEndpointWithTTL("newcname.example.com", endpoint.RecordTypeCNAME, 10, "other.com"),
		endpoint.MustNewEndpoint("new.nope.com", endpoint.RecordTypeA, "222.111.222.111"),
		endpoint.MustNewEndpoint("new.nope.com", endpoint.RecordTypeAAAA, "2001::222:111:222:111"),
		endpoint.MustNewEndpointWithTTL("newmail.example.com", endpoint.RecordTypeMX, 7200, "40 bar.other.com"),
	}

	deleteRecords := []*endpoint.Endpoint{
		endpoint.MustNewEndpoint("deleted.example.com", endpoint.RecordTypeA, "111.222.111.222"),
		endpoint.MustNewEndpoint("deletedaaaa.example.com", endpoint.RecordTypeAAAA, "2001::111:222:111:222"),
		endpoint.MustNewEndpoint("deletedcname.example.com", endpoint.RecordTypeCNAME, "other.com"),
		endpoint.MustNewEndpoint("deleted.nope.com", endpoint.RecordTypeA, "222.111.222.111"),
		endpoint.MustNewEndpoint("deleted.nope.com", endpoint.RecordTypeAAAA, "2001::222:111:222:111"),
	}

	changes := &plan.Changes{
		Create:    createRecords,
		UpdateNew: updatedRecords,
		UpdateOld: currentRecords,
		Delete:    deleteRecords,
	}

	if err := provider.ApplyChanges(t.Context(), changes); err != nil {
		t.Fatal(err)
	}
}

func TestAzurePrivateDNSNameFilter(t *testing.T) {
	provider := newMockedAzurePrivateDNSProvider(endpoint.NewDomainFilter([]string{"nginx.example.com"}), endpoint.NewDomainFilter([]string{"example.com"}), provider.NewZoneIDFilter([]string{""}), true, "k8s",
		[]*privatedns.PrivateZone{
			createMockPrivateZone("example.com", "/privateDnsZones/example.com"),
		},

		[]*privatedns.RecordSet{
			createPrivateMockRecordSet("NS", "ns1-03.azure-dns.com."),
			createPrivateMockRecordSet("SOA", "Email: azuredns-hostmaster.microsoft.com"),
			createPrivateMockRecordSet(endpoint.RecordTypeA, "123.123.123.122"),
			createPrivateMockRecordSet(endpoint.RecordTypeTXT, "heritage=external-dns,external-dns/owner=default"),
			createPrivateMockRecordSetWithNameAndTTL("test.nginx", endpoint.RecordTypeA, "123.123.123.123", 3600),
			createPrivateMockRecordSetWithNameAndTTL("nginx", endpoint.RecordTypeA, "123.123.123.123", 3600),
			createPrivateMockRecordSetWithNameAndTTL("nginx", endpoint.RecordTypeTXT, "heritage=external-dns,external-dns/owner=default", recordTTL),
			createPrivateMockRecordSetWithNameAndTTL("mail.nginx", endpoint.RecordTypeMX, "20 example.com", recordTTL),
			createPrivateMockRecordSetWithNameAndTTL("hack", endpoint.RecordTypeCNAME, "hack.azurewebsites.net", 10),
		}, 3)

	ctx := t.Context()
	actual, err := provider.Records(ctx)
	if err != nil {
		t.Fatal(err)
	}
	expected := []*endpoint.Endpoint{
		endpoint.MustNewEndpointWithTTL("test.nginx.example.com", endpoint.RecordTypeA, 3600, "123.123.123.123"),
		endpoint.MustNewEndpointWithTTL("nginx.example.com", endpoint.RecordTypeA, 3600, "123.123.123.123"),
		endpoint.MustNewEndpointWithTTL("nginx.example.com", endpoint.RecordTypeTXT, recordTTL, "heritage=external-dns,external-dns/owner=default"),
		endpoint.MustNewEndpointWithTTL("mail.nginx.example.com", endpoint.RecordTypeMX, recordTTL, "20 example.com"),
	}

	validateAzureEndpoints(t, actual, expected)
}

func TestAzurePrivateDNSApplyChangesZoneName(t *testing.T) {
	recordsClient := mockPrivateRecordSetsClient{}

	testAzurePrivateDNSApplyChangesInternalZoneName(t, false, &recordsClient)

	validateAzureEndpoints(t, recordsClient.deletedEndpoints, []*endpoint.Endpoint{
		endpoint.MustNewEndpoint("deleted.foo.example.com", endpoint.RecordTypeA, ""),
		endpoint.MustNewEndpoint("deletedaaaa.foo.example.com", endpoint.RecordTypeAAAA, ""),
		endpoint.MustNewEndpoint("deletedcname.foo.example.com", endpoint.RecordTypeCNAME, ""),
	})

	validateAzureEndpoints(t, recordsClient.updatedEndpoints, []*endpoint.Endpoint{
		endpoint.MustNewEndpointWithTTL("foo.example.com", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "1.2.3.4", "1.2.3.5"),
		endpoint.MustNewEndpointWithTTL("foo.example.com", endpoint.RecordTypeAAAA, endpoint.TTL(recordTTL), "2001::1:2:3:4", "2001::1:2:3:5"),
		endpoint.MustNewEndpointWithTTL("foo.example.com", endpoint.RecordTypeTXT, endpoint.TTL(recordTTL), "tag"),
		endpoint.MustNewEndpointWithTTL("new.foo.example.com", endpoint.RecordTypeA, 3600, "111.222.111.222"),
		endpoint.MustNewEndpointWithTTL("new.foo.example.com", endpoint.RecordTypeAAAA, 3600, "2001::111:222:111:222"),
		endpoint.MustNewEndpointWithTTL("newcname.foo.example.com", endpoint.RecordTypeCNAME, 10, "other.com"),
	})
}

func testAzurePrivateDNSApplyChangesInternalZoneName(t *testing.T, dryRun bool, client PrivateRecordSetsClient) {
	zones := []*privatedns.PrivateZone{
		createMockPrivateZone("example.com", "/privateDnsZones/example.com"),
	}
	zonesClient := newMockPrivateZonesClient(zones)

	provider := newAzurePrivateDNSProvider(
		endpoint.NewDomainFilter([]string{"foo.example.com"}),
		endpoint.NewDomainFilter([]string{"example.com"}),
		provider.NewZoneIDFilter([]string{""}),
		dryRun,
		"group",
		&zonesClient,
		client,
		3,
	)

	createRecords := []*endpoint.Endpoint{
		endpoint.MustNewEndpoint("example.com", endpoint.RecordTypeA, "1.2.3.4"),
		endpoint.MustNewEndpoint("example.com", endpoint.RecordTypeAAAA, "2001::1:2:3:4"),
		endpoint.MustNewEndpoint("example.com", endpoint.RecordTypeTXT, "tag"),
		endpoint.MustNewEndpoint("foo.example.com", endpoint.RecordTypeA, "1.2.3.5", "1.2.3.4"),
		endpoint.MustNewEndpoint("foo.example.com", endpoint.RecordTypeAAAA, "2001::1:2:3:5", "2001::1:2:3:4"),
		endpoint.MustNewEndpoint("foo.example.com", endpoint.RecordTypeTXT, "tag"),
		endpoint.MustNewEndpoint("bar.example.com", endpoint.RecordTypeCNAME, "other.com"),
		endpoint.MustNewEndpoint("bar.example.com", endpoint.RecordTypeTXT, "tag"),
		endpoint.MustNewEndpoint("other.com", endpoint.RecordTypeA, "5.6.7.8"),
		endpoint.MustNewEndpoint("other.com", endpoint.RecordTypeTXT, "tag"),
		endpoint.MustNewEndpoint("nope.com", endpoint.RecordTypeA, "4.4.4.4"),
		endpoint.MustNewEndpoint("nope.com", endpoint.RecordTypeTXT, "tag"),
	}

	currentRecords := []*endpoint.Endpoint{
		endpoint.MustNewEndpoint("old.foo.example.com", endpoint.RecordTypeA, "121.212.121.212"),
		endpoint.MustNewEndpoint("oldcname.foo.example.com", endpoint.RecordTypeCNAME, "other.com"),
		endpoint.MustNewEndpoint("old.nope.example.com", endpoint.RecordTypeA, "121.212.121.212"),
	}
	updatedRecords := []*endpoint.Endpoint{
		endpoint.MustNewEndpointWithTTL("new.foo.example.com", endpoint.RecordTypeA, 3600, "111.222.111.222"),
		endpoint.MustNewEndpointWithTTL("new.foo.example.com", endpoint.RecordTypeAAAA, 3600, "2001::111:222:111:222"),
		endpoint.MustNewEndpointWithTTL("newcname.foo.example.com", endpoint.RecordTypeCNAME, 10, "other.com"),
		endpoint.MustNewEndpoint("new.nope.example.com", endpoint.RecordTypeA, "222.111.222.111"),
		endpoint.MustNewEndpoint("new.nope.example.com", endpoint.RecordTypeAAAA, "2001::222:111:222:111"),
	}

	deleteRecords := []*endpoint.Endpoint{
		endpoint.MustNewEndpoint("deleted.foo.example.com", endpoint.RecordTypeA, "111.222.111.222"),
		endpoint.MustNewEndpoint("deletedaaaa.foo.example.com", endpoint.RecordTypeAAAA, "2001::111:222:111:222"),
		endpoint.MustNewEndpoint("deletedcname.foo.example.com", endpoint.RecordTypeCNAME, "other.com"),
		endpoint.MustNewEndpoint("deleted.nope.example.com", endpoint.RecordTypeA, "222.111.222.111"),
	}

	changes := &plan.Changes{
		Create:    createRecords,
		UpdateNew: updatedRecords,
		UpdateOld: currentRecords,
		Delete:    deleteRecords,
	}

	if err := provider.ApplyChanges(t.Context(), changes); err != nil {
		t.Fatal(err)
	}
}
