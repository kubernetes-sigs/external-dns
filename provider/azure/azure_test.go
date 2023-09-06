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
	dns "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dns/armdns"
	"github.com/stretchr/testify/assert"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

// mockZonesClient implements the methods of the Azure DNS Zones Client which are used in the Azure Provider
// and returns static results which are defined per test
type mockZonesClient struct {
	pagingHandler azcoreruntime.PagingHandler[dns.ZonesClientListByResourceGroupResponse]
}

func newMockZonesClient(zones []*dns.Zone) mockZonesClient {
	pagingHandler := azcoreruntime.PagingHandler[dns.ZonesClientListByResourceGroupResponse]{
		More: func(resp dns.ZonesClientListByResourceGroupResponse) bool {
			return false
		},
		Fetcher: func(context.Context, *dns.ZonesClientListByResourceGroupResponse) (dns.ZonesClientListByResourceGroupResponse, error) {
			return dns.ZonesClientListByResourceGroupResponse{
				ZoneListResult: dns.ZoneListResult{
					Value: zones,
				},
			}, nil
		},
	}
	return mockZonesClient{
		pagingHandler: pagingHandler,
	}
}

func (client *mockZonesClient) NewListByResourceGroupPager(resourceGroupName string, options *dns.ZonesClientListByResourceGroupOptions) *azcoreruntime.Pager[dns.ZonesClientListByResourceGroupResponse] {
	return azcoreruntime.NewPager(client.pagingHandler)
}

// mockZonesClient implements the methods of the Azure DNS RecordSet Client which are used in the Azure Provider
// and returns static results which are defined per test
type mockRecordSetsClient struct {
	pagingHandler    azcoreruntime.PagingHandler[dns.RecordSetsClientListAllByDNSZoneResponse]
	deletedEndpoints []*endpoint.Endpoint
	updatedEndpoints []*endpoint.Endpoint
}

func newMockRecordSetsClient(recordSets []*dns.RecordSet) mockRecordSetsClient {
	pagingHandler := azcoreruntime.PagingHandler[dns.RecordSetsClientListAllByDNSZoneResponse]{
		More: func(resp dns.RecordSetsClientListAllByDNSZoneResponse) bool {
			return false
		},
		Fetcher: func(context.Context, *dns.RecordSetsClientListAllByDNSZoneResponse) (dns.RecordSetsClientListAllByDNSZoneResponse, error) {
			return dns.RecordSetsClientListAllByDNSZoneResponse{
				RecordSetListResult: dns.RecordSetListResult{
					Value: recordSets,
				},
			}, nil
		},
	}
	return mockRecordSetsClient{
		pagingHandler: pagingHandler,
	}
}

func (client *mockRecordSetsClient) NewListAllByDNSZonePager(resourceGroupName string, zoneName string, options *dns.RecordSetsClientListAllByDNSZoneOptions) *azcoreruntime.Pager[dns.RecordSetsClientListAllByDNSZoneResponse] {
	return azcoreruntime.NewPager(client.pagingHandler)
}

func (client *mockRecordSetsClient) Delete(ctx context.Context, resourceGroupName string, zoneName string, relativeRecordSetName string, recordType dns.RecordType, options *dns.RecordSetsClientDeleteOptions) (dns.RecordSetsClientDeleteResponse, error) {
	client.deletedEndpoints = append(
		client.deletedEndpoints,
		endpoint.NewEndpoint(
			formatAzureDNSName(relativeRecordSetName, zoneName),
			string(recordType),
			"",
		),
	)
	return dns.RecordSetsClientDeleteResponse{}, nil
}

func (client *mockRecordSetsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, zoneName string, relativeRecordSetName string, recordType dns.RecordType, parameters dns.RecordSet, options *dns.RecordSetsClientCreateOrUpdateOptions) (dns.RecordSetsClientCreateOrUpdateResponse, error) {
	var ttl endpoint.TTL
	if parameters.Properties.TTL != nil {
		ttl = endpoint.TTL(*parameters.Properties.TTL)
	}
	client.updatedEndpoints = append(
		client.updatedEndpoints,
		endpoint.NewEndpointWithTTL(
			formatAzureDNSName(relativeRecordSetName, zoneName),
			string(recordType),
			ttl,
			extractAzureTargets(&parameters)...,
		),
	)
	return dns.RecordSetsClientCreateOrUpdateResponse{}, nil
}

func createMockZone(zone string, id string) *dns.Zone {
	return &dns.Zone{
		ID:   to.Ptr(id),
		Name: to.Ptr(zone),
	}
}

func aRecordSetPropertiesGetter(values []string, ttl int64) *dns.RecordSetProperties {
	aRecords := make([]*dns.ARecord, len(values))
	for i, value := range values {
		aRecords[i] = &dns.ARecord{
			IPv4Address: to.Ptr(value),
		}
	}
	return &dns.RecordSetProperties{
		TTL:      to.Ptr(ttl),
		ARecords: aRecords,
	}
}

func cNameRecordSetPropertiesGetter(values []string, ttl int64) *dns.RecordSetProperties {
	return &dns.RecordSetProperties{
		TTL: to.Ptr(ttl),
		CnameRecord: &dns.CnameRecord{
			Cname: to.Ptr(values[0]),
		},
	}
}

func mxRecordSetPropertiesGetter(values []string, ttl int64) *dns.RecordSetProperties {
	mxRecords := make([]*dns.MxRecord, len(values))
	for i, target := range values {
		mxRecord, _ := parseMxTarget[dns.MxRecord](target)
		mxRecords[i] = &mxRecord
	}
	return &dns.RecordSetProperties{
		TTL:       to.Ptr(ttl),
		MxRecords: mxRecords,
	}
}

func txtRecordSetPropertiesGetter(values []string, ttl int64) *dns.RecordSetProperties {
	return &dns.RecordSetProperties{
		TTL: to.Ptr(ttl),
		TxtRecords: []*dns.TxtRecord{
			{
				Value: []*string{to.Ptr(values[0])},
			},
		},
	}
}

func othersRecordSetPropertiesGetter(values []string, ttl int64) *dns.RecordSetProperties {
	return &dns.RecordSetProperties{
		TTL: to.Ptr(ttl),
	}
}

func createMockRecordSet(name, recordType string, values ...string) *dns.RecordSet {
	return createMockRecordSetMultiWithTTL(name, recordType, 0, values...)
}

func createMockRecordSetWithTTL(name, recordType, value string, ttl int64) *dns.RecordSet {
	return createMockRecordSetMultiWithTTL(name, recordType, ttl, value)
}

func createMockRecordSetMultiWithTTL(name, recordType string, ttl int64, values ...string) *dns.RecordSet {
	var getterFunc func(values []string, ttl int64) *dns.RecordSetProperties

	switch recordType {
	case endpoint.RecordTypeA:
		getterFunc = aRecordSetPropertiesGetter
	case endpoint.RecordTypeCNAME:
		getterFunc = cNameRecordSetPropertiesGetter
	case endpoint.RecordTypeMX:
		getterFunc = mxRecordSetPropertiesGetter
	case endpoint.RecordTypeTXT:
		getterFunc = txtRecordSetPropertiesGetter
	default:
		getterFunc = othersRecordSetPropertiesGetter
	}
	return &dns.RecordSet{
		Name:       to.Ptr(name),
		Type:       to.Ptr("Microsoft.Network/dnszones/" + recordType),
		Properties: getterFunc(values, ttl),
	}
}

// newMockedAzureProvider creates an AzureProvider comprising the mocked clients for zones and recordsets
func newMockedAzureProvider(domainFilter endpoint.DomainFilter, zoneNameFilter endpoint.DomainFilter, zoneIDFilter provider.ZoneIDFilter, dryRun bool, resourceGroup string, userAssignedIdentityClientID string, zones []*dns.Zone, recordSets []*dns.RecordSet) (*AzureProvider, error) {
	zonesClient := newMockZonesClient(zones)
	recordSetsClient := newMockRecordSetsClient(recordSets)
	return newAzureProvider(domainFilter, zoneNameFilter, zoneIDFilter, dryRun, resourceGroup, userAssignedIdentityClientID, &zonesClient, &recordSetsClient), nil
}

func newAzureProvider(domainFilter endpoint.DomainFilter, zoneNameFilter endpoint.DomainFilter, zoneIDFilter provider.ZoneIDFilter, dryRun bool, resourceGroup string, userAssignedIdentityClientID string, zonesClient ZonesClient, recordsClient RecordSetsClient) *AzureProvider {
	return &AzureProvider{
		domainFilter:                 domainFilter,
		zoneNameFilter:               zoneNameFilter,
		zoneIDFilter:                 zoneIDFilter,
		dryRun:                       dryRun,
		resourceGroup:                resourceGroup,
		userAssignedIdentityClientID: userAssignedIdentityClientID,
		zonesClient:                  zonesClient,
		recordSetsClient:             recordsClient,
	}
}

func validateAzureEndpoints(t *testing.T, endpoints []*endpoint.Endpoint, expected []*endpoint.Endpoint) {
	assert.True(t, testutils.SameEndpoints(endpoints, expected), "expected and actual endpoints don't match. %s:%s", endpoints, expected)
}

func TestAzureRecord(t *testing.T) {
	provider, err := newMockedAzureProvider(endpoint.NewDomainFilter([]string{"example.com"}), endpoint.NewDomainFilter([]string{}), provider.NewZoneIDFilter([]string{""}), true, "k8s", "",
		[]*dns.Zone{
			createMockZone("example.com", "/dnszones/example.com"),
		},
		[]*dns.RecordSet{
			createMockRecordSet("@", "NS", "ns1-03.azure-dns.com."),
			createMockRecordSet("@", "SOA", "Email: azuredns-hostmaster.microsoft.com"),
			createMockRecordSet("@", endpoint.RecordTypeA, "123.123.123.122"),
			createMockRecordSet("@", endpoint.RecordTypeTXT, "heritage=external-dns,external-dns/owner=default"),
			createMockRecordSetWithTTL("nginx", endpoint.RecordTypeA, "123.123.123.123", 3600),
			createMockRecordSetWithTTL("nginx", endpoint.RecordTypeTXT, "heritage=external-dns,external-dns/owner=default", recordTTL),
			createMockRecordSetWithTTL("hack", endpoint.RecordTypeCNAME, "hack.azurewebsites.net", 10),
			createMockRecordSetMultiWithTTL("mail", endpoint.RecordTypeMX, 4000, "10 example.com"),
		})
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	actual, err := provider.Records(ctx)
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

func TestAzureMultiRecord(t *testing.T) {
	provider, err := newMockedAzureProvider(endpoint.NewDomainFilter([]string{"example.com"}), endpoint.NewDomainFilter([]string{}), provider.NewZoneIDFilter([]string{""}), true, "k8s", "",
		[]*dns.Zone{
			createMockZone("example.com", "/dnszones/example.com"),
		},
		[]*dns.RecordSet{
			createMockRecordSet("@", "NS", "ns1-03.azure-dns.com."),
			createMockRecordSet("@", "SOA", "Email: azuredns-hostmaster.microsoft.com"),
			createMockRecordSet("@", endpoint.RecordTypeA, "123.123.123.122", "234.234.234.233"),
			createMockRecordSet("@", endpoint.RecordTypeTXT, "heritage=external-dns,external-dns/owner=default"),
			createMockRecordSetMultiWithTTL("nginx", endpoint.RecordTypeA, 3600, "123.123.123.123", "234.234.234.234"),
			createMockRecordSetWithTTL("nginx", endpoint.RecordTypeTXT, "heritage=external-dns,external-dns/owner=default", recordTTL),
			createMockRecordSetWithTTL("hack", endpoint.RecordTypeCNAME, "hack.azurewebsites.net", 10),
			createMockRecordSetMultiWithTTL("mail", endpoint.RecordTypeMX, 4000, "10 example.com", "20 backup.example.com"),
		})
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	actual, err := provider.Records(ctx)
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

func TestAzureApplyChanges(t *testing.T) {
	recordsClient := mockRecordSetsClient{}

	testAzureApplyChangesInternal(t, false, &recordsClient)

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

func TestAzureApplyChangesDryRun(t *testing.T) {
	recordsClient := mockRecordSetsClient{}

	testAzureApplyChangesInternal(t, true, &recordsClient)

	validateAzureEndpoints(t, recordsClient.deletedEndpoints, []*endpoint.Endpoint{})

	validateAzureEndpoints(t, recordsClient.updatedEndpoints, []*endpoint.Endpoint{})
}

func testAzureApplyChangesInternal(t *testing.T, dryRun bool, client RecordSetsClient) {
	zones := []*dns.Zone{
		createMockZone("example.com", "/dnszones/example.com"),
		createMockZone("other.com", "/dnszones/other.com"),
	}
	zonesClient := newMockZonesClient(zones)

	provider := newAzureProvider(
		endpoint.NewDomainFilter([]string{""}),
		endpoint.NewDomainFilter([]string{""}),
		provider.NewZoneIDFilter([]string{""}),
		dryRun,
		"group",
		"",
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

func TestAzureNameFilter(t *testing.T) {
	provider, err := newMockedAzureProvider(endpoint.NewDomainFilter([]string{"nginx.example.com"}), endpoint.NewDomainFilter([]string{"example.com"}), provider.NewZoneIDFilter([]string{""}), true, "k8s", "",
		[]*dns.Zone{
			createMockZone("example.com", "/dnszones/example.com"),
		},

		[]*dns.RecordSet{
			createMockRecordSet("@", "NS", "ns1-03.azure-dns.com."),
			createMockRecordSet("@", "SOA", "Email: azuredns-hostmaster.microsoft.com"),
			createMockRecordSet("@", endpoint.RecordTypeA, "123.123.123.122"),
			createMockRecordSet("@", endpoint.RecordTypeTXT, "heritage=external-dns,external-dns/owner=default"),
			createMockRecordSetWithTTL("test.nginx", endpoint.RecordTypeA, "123.123.123.123", 3600),
			createMockRecordSetWithTTL("nginx", endpoint.RecordTypeA, "123.123.123.123", 3600),
			createMockRecordSetWithTTL("nginx", endpoint.RecordTypeTXT, "heritage=external-dns,external-dns/owner=default", recordTTL),
			createMockRecordSetWithTTL("mail.nginx", endpoint.RecordTypeMX, "20 example.com", recordTTL),
			createMockRecordSetWithTTL("hack", endpoint.RecordTypeCNAME, "hack.azurewebsites.net", 10),
		})
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	actual, err := provider.Records(ctx)
	if err != nil {
		t.Fatal(err)
	}
	expected := []*endpoint.Endpoint{
		endpoint.NewEndpointWithTTL("test.nginx.example.com", endpoint.RecordTypeA, 3600, "123.123.123.123"),
		endpoint.NewEndpointWithTTL("nginx.example.com", endpoint.RecordTypeA, 3600, "123.123.123.123"),
		endpoint.NewEndpointWithTTL("nginx.example.com", endpoint.RecordTypeTXT, recordTTL, "heritage=external-dns,external-dns/owner=default"),
		endpoint.NewEndpointWithTTL("mail.nginx.example.com", endpoint.RecordTypeMX, recordTTL, "20 example.com"),
	}

	validateAzureEndpoints(t, actual, expected)
}

func TestAzureApplyChangesZoneName(t *testing.T) {
	recordsClient := mockRecordSetsClient{}

	testAzureApplyChangesInternalZoneName(t, false, &recordsClient)

	validateAzureEndpoints(t, recordsClient.deletedEndpoints, []*endpoint.Endpoint{
		endpoint.NewEndpoint("deleted.foo.example.com", endpoint.RecordTypeA, ""),
		endpoint.NewEndpoint("deletedcname.foo.example.com", endpoint.RecordTypeCNAME, ""),
	})

	validateAzureEndpoints(t, recordsClient.updatedEndpoints, []*endpoint.Endpoint{
		endpoint.NewEndpointWithTTL("foo.example.com", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "1.2.3.4", "1.2.3.5"),
		endpoint.NewEndpointWithTTL("foo.example.com", endpoint.RecordTypeTXT, endpoint.TTL(recordTTL), "tag"),
		endpoint.NewEndpointWithTTL("new.foo.example.com", endpoint.RecordTypeA, 3600, "111.222.111.222"),
		endpoint.NewEndpointWithTTL("newcname.foo.example.com", endpoint.RecordTypeCNAME, 10, "other.com"),
	})
}

func testAzureApplyChangesInternalZoneName(t *testing.T, dryRun bool, client RecordSetsClient) {
	zonesClient := newMockZonesClient([]*dns.Zone{createMockZone("example.com", "/dnszones/example.com")})

	provider := newAzureProvider(
		endpoint.NewDomainFilter([]string{"foo.example.com"}),
		endpoint.NewDomainFilter([]string{"example.com"}),
		provider.NewZoneIDFilter([]string{""}),
		dryRun,
		"group",
		"",
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
		endpoint.NewEndpoint("old.foo.example.com", endpoint.RecordTypeA, "121.212.121.212"),
		endpoint.NewEndpoint("oldcname.foo.example.com", endpoint.RecordTypeCNAME, "other.com"),
		endpoint.NewEndpoint("old.nope.example.com", endpoint.RecordTypeA, "121.212.121.212"),
	}
	updatedRecords := []*endpoint.Endpoint{
		endpoint.NewEndpointWithTTL("new.foo.example.com", endpoint.RecordTypeA, 3600, "111.222.111.222"),
		endpoint.NewEndpointWithTTL("newcname.foo.example.com", endpoint.RecordTypeCNAME, 10, "other.com"),
		endpoint.NewEndpoint("new.nope.example.com", endpoint.RecordTypeA, "222.111.222.111"),
	}

	deleteRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("deleted.foo.example.com", endpoint.RecordTypeA, "111.222.111.222"),
		endpoint.NewEndpoint("deletedcname.foo.example.com", endpoint.RecordTypeCNAME, "other.com"),
		endpoint.NewEndpoint("deleted.nope.example.com", endpoint.RecordTypeA, "222.111.222.111"),
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
