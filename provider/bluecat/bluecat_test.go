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

package bluecat

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

type mockGatewayClient struct {
	mockBluecatZones  *[]BluecatZone
	mockBluecatHosts  *[]BluecatHostRecord
	mockBluecatCNAMEs *[]BluecatCNAMERecord
	mockBluecatTXTs   *[]BluecatTXTRecord
}

type Changes struct {
	// Records that need to be created
	Create []*endpoint.Endpoint
	// Records that need to be updated (current data)
	UpdateOld []*endpoint.Endpoint
	// Records that need to be updated (desired data)
	UpdateNew []*endpoint.Endpoint
	// Records that need to be deleted
	Delete []*endpoint.Endpoint
}

func (g mockGatewayClient) getBluecatZones(zoneName string) ([]BluecatZone, error) {
	return *g.mockBluecatZones, nil
}
func (g mockGatewayClient) getHostRecords(zone string, records *[]BluecatHostRecord) error {
	*records = *g.mockBluecatHosts
	return nil
}
func (g mockGatewayClient) getCNAMERecords(zone string, records *[]BluecatCNAMERecord) error {
	*records = *g.mockBluecatCNAMEs
	return nil
}
func (g mockGatewayClient) getHostRecord(name string, record *BluecatHostRecord) error {
	for _, currentRecord := range *g.mockBluecatHosts {
		if currentRecord.Name == strings.Split(name, ".")[0] {
			*record = currentRecord
			return nil
		}
	}
	return nil
}
func (g mockGatewayClient) getCNAMERecord(name string, record *BluecatCNAMERecord) error {
	for _, currentRecord := range *g.mockBluecatCNAMEs {
		if currentRecord.Name == strings.Split(name, ".")[0] {
			*record = currentRecord
			return nil
		}
	}
	return nil
}
func (g mockGatewayClient) createHostRecord(zone string, req *bluecatCreateHostRecordRequest) (res interface{}, err error) {
	return nil, nil
}
func (g mockGatewayClient) createCNAMERecord(zone string, req *bluecatCreateCNAMERecordRequest) (res interface{}, err error) {
	return nil, nil
}
func (g mockGatewayClient) deleteHostRecord(name string, zone string) (err error) {
	*g.mockBluecatHosts = nil
	return nil
}
func (g mockGatewayClient) deleteCNAMERecord(name string, zone string) (err error) {
	*g.mockBluecatCNAMEs = nil
	return nil
}
func (g mockGatewayClient) getTXTRecords(zone string, records *[]BluecatTXTRecord) error {
	*records = *g.mockBluecatTXTs
	return nil
}
func (g mockGatewayClient) getTXTRecord(name string, record *BluecatTXTRecord) error {
	for _, currentRecord := range *g.mockBluecatTXTs {
		if currentRecord.Name == name {
			*record = currentRecord
			return nil
		}
	}
	return nil
}
func (g mockGatewayClient) createTXTRecord(zone string, req *bluecatCreateTXTRecordRequest) (res interface{}, err error) {
	return nil, nil
}
func (g mockGatewayClient) deleteTXTRecord(name string, zone string) error {
	*g.mockBluecatTXTs = nil
	return nil
}

func (g mockGatewayClient) buildHTTPRequest(method, url string, body io.Reader) (*http.Request, error) {
	request, _ := http.NewRequest("GET", fmt.Sprintf("%s/users", "http://some.com/api/v1"), nil)
	return request, nil
}

func createMockBluecatZone(fqdn string) BluecatZone {
	props := "absoluteName=" + fqdn
	return BluecatZone{
		Properties: props,
		Name:       fqdn,
		ID:         3,
	}
}

func createMockBluecatHostRecord(fqdn, target string, ttl int) BluecatHostRecord {
	props := "absoluteName=" + fqdn + "|addresses=" + target + "|ttl=" + fmt.Sprint(ttl) + "|"
	nameParts := strings.Split(fqdn, ".")
	return BluecatHostRecord{
		Name:       nameParts[0],
		Properties: props,
		ID:         3,
	}
}

func createMockBluecatCNAME(alias, target string, ttl int) BluecatCNAMERecord {
	props := "absoluteName=" + alias + "|linkedRecordName=" + target + "|ttl=" + fmt.Sprint(ttl) + "|"
	nameParts := strings.Split(alias, ".")
	return BluecatCNAMERecord{
		Name:       nameParts[0],
		Properties: props,
	}
}

func createMockBluecatTXT(fqdn, txt string) BluecatTXTRecord {
	return BluecatTXTRecord{
		Name:       fqdn,
		Properties: txt,
	}
}

func newBluecatProvider(domainFilter endpoint.DomainFilter, zoneIDFilter provider.ZoneIDFilter, dryRun bool, client GatewayClient) *BluecatProvider {
	return &BluecatProvider{
		domainFilter:  domainFilter,
		zoneIDFilter:  zoneIDFilter,
		dryRun:        dryRun,
		gatewayClient: client,
	}
}

type bluecatTestData []struct {
	TestDescription string
	Endpoints       []*endpoint.Endpoint
}

var tests = bluecatTestData{
	{
		"first test case", // TODO: better test description
		[]*endpoint.Endpoint{
			{
				DNSName:    "example.com",
				RecordType: endpoint.RecordTypeA,
				Targets:    endpoint.Targets{"123.123.123.122"},
				RecordTTL:  endpoint.TTL(30),
			},
			{
				DNSName:    "nginx.example.com",
				RecordType: endpoint.RecordTypeA,
				Targets:    endpoint.Targets{"123.123.123.123"},
				RecordTTL:  endpoint.TTL(30),
			},
			{
				DNSName:    "whitespace.example.com",
				RecordType: endpoint.RecordTypeA,
				Targets:    endpoint.Targets{"123.123.123.124"},
				RecordTTL:  endpoint.TTL(30),
			},
			{
				DNSName:    "hack.example.com",
				RecordType: endpoint.RecordTypeCNAME,
				Targets:    endpoint.Targets{"bluecatnetworks.com"},
				RecordTTL:  endpoint.TTL(30),
			},
			{
				DNSName:    "wack.example.com",
				RecordType: endpoint.RecordTypeTXT,
				Targets:    endpoint.Targets{"hello"},
				Labels:     endpoint.Labels{"owner": ""},
			},
			{
				DNSName:    "sack.example.com",
				RecordType: endpoint.RecordTypeTXT,
				Targets:    endpoint.Targets{""},
				Labels:     endpoint.Labels{"owner": ""},
			},
			{
				DNSName:    "kdb.example.com",
				RecordType: endpoint.RecordTypeTXT,
				Targets:    endpoint.Targets{"heritage=external-dns,external-dns/owner=default,external-dns/resource=service/openshift-ingress/router-default"},
				Labels:     endpoint.Labels{"owner": "default"},
			},
		},
	},
}

func TestBluecatRecords(t *testing.T) {
	client := mockGatewayClient{
		mockBluecatZones: &[]BluecatZone{
			createMockBluecatZone("example.com"),
		},
		mockBluecatTXTs: &[]BluecatTXTRecord{
			createMockBluecatTXT("kdb.example.com", "heritage=external-dns,external-dns/owner=default,external-dns/resource=service/openshift-ingress/router-default"),
			createMockBluecatTXT("wack.example.com", "hello"),
			createMockBluecatTXT("sack.example.com", ""),
		},
		mockBluecatHosts: &[]BluecatHostRecord{
			createMockBluecatHostRecord("example.com", "123.123.123.122", 30),
			createMockBluecatHostRecord("nginx.example.com", "123.123.123.123", 30),
			createMockBluecatHostRecord("whitespace.example.com", "123.123.123.124", 30),
		},
		mockBluecatCNAMEs: &[]BluecatCNAMERecord{
			createMockBluecatCNAME("hack.example.com", "bluecatnetworks.com", 30),
		},
	}

	provider := newBluecatProvider(
		endpoint.NewDomainFilter([]string{"example.com"}),
		provider.NewZoneIDFilter([]string{""}), false, client)

	for _, ti := range tests {
		actual, err := provider.Records(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		validateEndpoints(t, actual, ti.Endpoints)
	}
}

func TestBluecatApplyChangesCreate(t *testing.T) {
	client := mockGatewayClient{
		mockBluecatZones: &[]BluecatZone{
			createMockBluecatZone("example.com"),
		},
		mockBluecatHosts:  &[]BluecatHostRecord{},
		mockBluecatCNAMEs: &[]BluecatCNAMERecord{},
		mockBluecatTXTs:   &[]BluecatTXTRecord{},
	}

	provider := newBluecatProvider(
		endpoint.NewDomainFilter([]string{"example.com"}),
		provider.NewZoneIDFilter([]string{""}), false, client)

	for _, ti := range tests {
		err := provider.ApplyChanges(context.Background(), &plan.Changes{Create: ti.Endpoints})
		if err != nil {
			t.Fatal(err)
		}

		actual, err := provider.Records(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		validateEndpoints(t, actual, []*endpoint.Endpoint{})
	}
}
func TestBluecatApplyChangesDelete(t *testing.T) {
	client := mockGatewayClient{
		mockBluecatZones: &[]BluecatZone{
			createMockBluecatZone("example.com"),
		},
		mockBluecatHosts: &[]BluecatHostRecord{
			createMockBluecatHostRecord("example.com", "123.123.123.122", 30),
			createMockBluecatHostRecord("nginx.example.com", "123.123.123.123", 30),
			createMockBluecatHostRecord("whitespace.example.com", "123.123.123.124", 30),
		},
		mockBluecatCNAMEs: &[]BluecatCNAMERecord{
			createMockBluecatCNAME("hack.example.com", "bluecatnetworks.com", 30),
		},
		mockBluecatTXTs: &[]BluecatTXTRecord{
			createMockBluecatTXT("kdb.example.com", "heritage=external-dns,external-dns/owner=default,external-dns/resource=service/openshift-ingress/router-default"),
			createMockBluecatTXT("wack.example.com", "hello"),
			createMockBluecatTXT("sack.example.com", ""),
		},
	}

	provider := newBluecatProvider(
		endpoint.NewDomainFilter([]string{"example.com"}),
		provider.NewZoneIDFilter([]string{""}), false, client)

	for _, ti := range tests {
		err := provider.ApplyChanges(context.Background(), &plan.Changes{Delete: ti.Endpoints})
		if err != nil {
			t.Fatal(err)
		}

		actual, err := provider.Records(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		validateEndpoints(t, actual, []*endpoint.Endpoint{})
	}
}

func TestBluecatApplyChangesDeleteWithOwner(t *testing.T) {
	client := mockGatewayClient{
		mockBluecatZones: &[]BluecatZone{
			createMockBluecatZone("example.com"),
		},
		mockBluecatHosts: &[]BluecatHostRecord{
			createMockBluecatHostRecord("example.com", "123.123.123.122", 30),
			createMockBluecatHostRecord("nginx.example.com", "123.123.123.123", 30),
			createMockBluecatHostRecord("whitespace.example.com", "123.123.123.124", 30),
		},
		mockBluecatCNAMEs: &[]BluecatCNAMERecord{
			createMockBluecatCNAME("hack.example.com", "bluecatnetworks.com", 30),
		},
		mockBluecatTXTs: &[]BluecatTXTRecord{
			createMockBluecatTXT("kdb.example.com", "heritage=external-dns,external-dns/owner=default,external-dns/resource=service/openshift-ingress/router-default"),
			createMockBluecatTXT("wack.example.com", "hello"),
			createMockBluecatTXT("sack.example.com", ""),
		},
	}

	provider := newBluecatProvider(
		endpoint.NewDomainFilter([]string{"example.com"}),
		provider.NewZoneIDFilter([]string{""}), false, client)

	for _, ti := range tests {
		for _, ep := range ti.Endpoints {
			if strings.Contains(ep.Targets.String(), "external-dns") {
				owner, err := extractOwnerfromTXTRecord(ep.Targets.String())
				if err != nil {
					continue
				}
				t.Logf("Owner %s %s", owner, err)
			}
		}
		err := provider.ApplyChanges(context.Background(), &plan.Changes{Delete: ti.Endpoints})
		if err != nil {
			t.Fatal(err)
		}
		actual, err := provider.Records(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		validateEndpoints(t, actual, []*endpoint.Endpoint{})
	}

}

func TestExpandZones(t *testing.T) {
	mockZones := []string{"example.com", "nginx.example.com", "hack.example.com"}
	expected := []string{"zones/com/zones/example/zones/", "zones/com/zones/example/zones/nginx/zones/", "zones/com/zones/example/zones/hack/zones/"}
	for i := range mockZones {
		if expandZone(mockZones[i]) != expected[i] {
			t.Fatalf("%s", expected[i])
		}
	}
}

func TestBluecatNewGatewayClient(t *testing.T) {
	testCookie := http.Cookie{Name: "testCookie", Value: "exampleCookie"}
	testToken := "exampleToken"
	testgateWayHost := "exampleHost"
	testDNSConfiguration := "exampleDNSConfiguration"
	testView := "testView"
	testZone := "example.com"
	testVerify := true

	client := NewGatewayClient(testCookie, testToken, testgateWayHost, testDNSConfiguration, testView, testZone, testVerify)

	if client.Cookie.Value != testCookie.Value || client.Cookie.Name != testCookie.Name || client.Token != testToken || client.Host != testgateWayHost || client.DNSConfiguration != testDNSConfiguration || client.View != testView || client.RootZone != testZone || client.SkipTLSVerify != testVerify {
		t.Fatal("Client values dont match")
	}
}

// TODO: ensure findZone method is tested
// TODO: ensure zones method is tested
// TODO: ensure createRecords method is tested
// TODO: ensure deleteRecords method is tested
// TODO: ensure recordSet method is tested

// TODO: Figure out why recordSet.res is not being set properly
func TestBluecatRecordset(t *testing.T) {
	client := mockGatewayClient{
		mockBluecatZones: &[]BluecatZone{
			createMockBluecatZone("example.com"),
		},
		mockBluecatHosts: &[]BluecatHostRecord{
			createMockBluecatHostRecord("example.com", "123.123.123.122", 30),
			createMockBluecatHostRecord("nginx.example.com", "123.123.123.123", 30),
			createMockBluecatHostRecord("whitespace.example.com", "123.123.123.124", 30),
		},
		mockBluecatCNAMEs: &[]BluecatCNAMERecord{
			createMockBluecatCNAME("hack.example.com", "bluecatnetworks.com", 30),
		},
		mockBluecatTXTs: &[]BluecatTXTRecord{
			createMockBluecatTXT("abc.example.com", "hello"),
		},
	}

	provider := newBluecatProvider(
		endpoint.NewDomainFilter([]string{"example.com"}),
		provider.NewZoneIDFilter([]string{""}), false, client)

	// Test txt records for recordSet function
	testTxtEndpoint := endpoint.NewEndpoint("abc.example.com", endpoint.RecordTypeTXT, "hello")
	txtObj := bluecatCreateTXTRecordRequest{
		AbsoluteName: testTxtEndpoint.DNSName,
		Text:         testTxtEndpoint.Targets[0],
	}
	txtRecords := []BluecatTXTRecord{
		createMockBluecatTXT("abc.example.com", "hello"),
	}
	expected := bluecatRecordSet{
		obj: &txtObj,
		res: &txtRecords,
	}
	actual, err := provider.recordSet(testTxtEndpoint, true)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, actual.obj, expected.obj)
	assert.Equal(t, actual.res, expected.res)

	// Test a records for recordSet function
	testHostEndpoint := endpoint.NewEndpoint("whitespace.example.com", endpoint.RecordTypeA, "123.123.123.124")
	hostObj := bluecatCreateHostRecordRequest{
		AbsoluteName: testHostEndpoint.DNSName,
		IP4Address:   testHostEndpoint.Targets[0],
	}
	hostRecords := []BluecatHostRecord{
		createMockBluecatHostRecord("whitespace.example.com", "123.123.123.124", 30),
	}
	hostExpected := bluecatRecordSet{
		obj: &hostObj,
		res: &hostRecords,
	}
	hostActual, err := provider.recordSet(testHostEndpoint, true)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, hostActual.obj, hostExpected.obj)
	assert.Equal(t, hostActual.res, hostExpected.res)

	// Test CName records for recordSet function
	testCnameEndpoint := endpoint.NewEndpoint("hack.example.com", endpoint.RecordTypeCNAME, "bluecatnetworks.com")
	cnameObj := bluecatCreateCNAMERecordRequest{
		AbsoluteName: testCnameEndpoint.DNSName,
		LinkedRecord: testCnameEndpoint.Targets[0],
	}
	cnameRecords := []BluecatCNAMERecord{
		createMockBluecatCNAME("hack.example.com", "bluecatnetworks.com", 30),
	}
	cnameExpected := bluecatRecordSet{
		obj: &cnameObj,
		res: &cnameRecords,
	}
	cnameActual, err := provider.recordSet(testCnameEndpoint, true)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, cnameActual.obj, cnameExpected.obj)
	assert.Equal(t, cnameActual.res, cnameExpected.res)
}

func validateEndpoints(t *testing.T, actual, expected []*endpoint.Endpoint) {
	assert.True(t, testutils.SameEndpoints(actual, expected), "actual and expected endpoints don't match. %s:%s", actual, expected)
}
