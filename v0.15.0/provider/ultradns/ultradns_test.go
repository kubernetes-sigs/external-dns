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

package ultradns

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	_ "strings"
	"testing"

	"github.com/stretchr/testify/assert"
	udnssdk "github.com/ultradns/ultradns-sdk-go"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

type mockUltraDNSZone struct {
	client *udnssdk.Client
}

func (m *mockUltraDNSZone) SelectWithOffsetWithLimit(k *udnssdk.ZoneKey, offset int, limit int) (zones []udnssdk.Zone, ResultInfo udnssdk.ResultInfo, resp *http.Response, err error) {
	zones = []udnssdk.Zone{}
	zone := udnssdk.Zone{}
	zoneJson := `
                        {
                           "properties": {
                              "name":"test-ultradns-provider.com.",
                              "accountName":"teamrest",
                              "type":"PRIMARY",
                              "dnssecStatus":"UNSIGNED",
                              "status":"ACTIVE",
                              "owner":"teamrest",
                              "resourceRecordCount":7,
                              "lastModifiedDateTime":""
                           }
                        }`
	if err := json.Unmarshal([]byte(zoneJson), &zone); err != nil {
		log.Fatal(err)
	}

	zones = append(zones, zone)
	return zones, udnssdk.ResultInfo{}, nil, nil
}

type mockUltraDNSRecord struct {
	client *udnssdk.Client
}

func (m *mockUltraDNSRecord) Create(k udnssdk.RRSetKey, rrset udnssdk.RRSet) (*http.Response, error) {
	return nil, nil
}

func (m *mockUltraDNSRecord) Select(k udnssdk.RRSetKey) ([]udnssdk.RRSet, error) {
	return []udnssdk.RRSet{{
		OwnerName: "test-ultradns-provider.com.",
		RRType:    endpoint.RecordTypeA,
		RData:     []string{"1.1.1.1"},
		TTL:       86400,
	}}, nil
}

func (m *mockUltraDNSRecord) SelectWithOffset(k udnssdk.RRSetKey, offset int) ([]udnssdk.RRSet, udnssdk.ResultInfo, *http.Response, error) {
	return nil, udnssdk.ResultInfo{}, nil, nil
}

func (m *mockUltraDNSRecord) Update(udnssdk.RRSetKey, udnssdk.RRSet) (*http.Response, error) {
	return nil, nil
}

func (m *mockUltraDNSRecord) Delete(k udnssdk.RRSetKey) (*http.Response, error) {
	return nil, nil
}

func (m *mockUltraDNSRecord) SelectWithOffsetWithLimit(k udnssdk.RRSetKey, offset int, limit int) (rrsets []udnssdk.RRSet, ResultInfo udnssdk.ResultInfo, resp *http.Response, err error) {
	return []udnssdk.RRSet{{
		OwnerName: "test-ultradns-provider.com.",
		RRType:    endpoint.RecordTypeA,
		RData:     []string{"1.1.1.1"},
		TTL:       86400,
	}}, udnssdk.ResultInfo{}, nil, nil
}

// NewUltraDNSProvider Test scenario
func TestNewUltraDNSProvider(t *testing.T) {
	_ = os.Setenv("ULTRADNS_USERNAME", "")
	_ = os.Setenv("ULTRADNS_PASSWORD", "")
	_ = os.Setenv("ULTRADNS_BASEURL", "")
	_ = os.Setenv("ULTRADNS_ACCOUNTNAME", "")
	_, err := NewUltraDNSProvider(endpoint.NewDomainFilter([]string{"test-ultradns-provider.com"}), true)
	assert.Nil(t, err)

	_ = os.Unsetenv("ULTRADNS_PASSWORD")
	_ = os.Unsetenv("ULTRADNS_USERNAME")
	_ = os.Unsetenv("ULTRADNS_BASEURL")
	_ = os.Unsetenv("ULTRADNS_ACCOUNTNAME")
	_, err = NewUltraDNSProvider(endpoint.NewDomainFilter([]string{"test-ultradns-provider.com"}), true)
	assert.NotNilf(t, err, "Expected to fail %s", "formatted")
}

// zones function test scenario
func TestUltraDNSProvider_Zones(t *testing.T) {
	mocked := mockUltraDNSZone{}
	provider := &UltraDNSProvider{
		client: udnssdk.Client{
			Zone: &mocked,
		},
	}

	zoneKey := &udnssdk.ZoneKey{
		Zone:        "",
		AccountName: "teamrest",
	}

	expected, _, _, err := provider.client.Zone.SelectWithOffsetWithLimit(zoneKey, 0, 1000)
	assert.Nil(t, err)
	zones, err := provider.Zones(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, reflect.DeepEqual(expected, zones), true)
}

// Records function test case
func TestUltraDNSProvider_Records(t *testing.T) {
	mocked := mockUltraDNSRecord{}
	mockedDomain := mockUltraDNSZone{}

	provider := &UltraDNSProvider{
		client: udnssdk.Client{
			RRSets: &mocked,
			Zone:   &mockedDomain,
		},
	}
	rrsetKey := udnssdk.RRSetKey{}
	expected, _, _, err := provider.client.RRSets.SelectWithOffsetWithLimit(rrsetKey, 0, 1000)
	records, err := provider.Records(context.Background())
	assert.Nil(t, err)
	for _, v := range records {
		assert.Equal(t, fmt.Sprintf("%s.", v.DNSName), expected[0].OwnerName)
		assert.Equal(t, v.RecordType, expected[0].RRType)
		assert.Equal(t, int(v.RecordTTL), expected[0].TTL)
	}
}

// ApplyChanges function testcase
func TestUltraDNSProvider_ApplyChanges(t *testing.T) {
	changes := &plan.Changes{}
	mocked := mockUltraDNSRecord{nil}
	mockedDomain := mockUltraDNSZone{nil}

	provider := &UltraDNSProvider{
		client: udnssdk.Client{
			RRSets: &mocked,
			Zone:   &mockedDomain,
		},
	}

	changes.Create = []*endpoint.Endpoint{
		{DNSName: "test-ultradns-provider.com", Targets: endpoint.Targets{"1.1.1.1"}, RecordType: "A"},
		{DNSName: "ttl.test-ultradns-provider.com", Targets: endpoint.Targets{"1.1.1.1"}, RecordType: "A", RecordTTL: 100},
	}
	changes.Create = []*endpoint.Endpoint{{DNSName: "test-ultradns-provider.com", Targets: endpoint.Targets{"1.1.1.2"}, RecordType: "A"}}
	changes.UpdateNew = []*endpoint.Endpoint{{DNSName: "test-ultradns-provider.com", Targets: endpoint.Targets{"1.1.2.2"}, RecordType: "A", RecordTTL: 100}}
	changes.UpdateNew = []*endpoint.Endpoint{{DNSName: "test-ultradns-provider.com", Targets: endpoint.Targets{"1.1.2.2", "1.1.2.3", "1.1.2.4"}, RecordType: "A", RecordTTL: 100}}
	changes.Delete = []*endpoint.Endpoint{{DNSName: "test-ultradns-provider.com", Targets: endpoint.Targets{"1.1.2.2", "1.1.2.3", "1.1.2.4"}, RecordType: "A", RecordTTL: 100}}
	changes.Delete = []*endpoint.Endpoint{{DNSName: "ttl.test-ultradns-provider.com", Targets: endpoint.Targets{"1.1.1.1"}, RecordType: "A", RecordTTL: 100}}
	err := provider.ApplyChanges(context.Background(), changes)
	assert.Nilf(t, err, "Should not fail %s", "formatted")
}

// Testing function getSpecificRecord
func TestUltraDNSProvider_getSpecificRecord(t *testing.T) {
	mocked := mockUltraDNSRecord{nil}
	mockedDomain := mockUltraDNSZone{nil}

	provider := &UltraDNSProvider{
		client: udnssdk.Client{
			RRSets: &mocked,
			Zone:   &mockedDomain,
		},
	}

	recordSetKey := udnssdk.RRSetKey{
		Zone: "test-ultradns-provider.com.",
		Type: "A",
		Name: "teamrest",
	}
	err := provider.getSpecificRecord(context.Background(), recordSetKey)
	assert.Nil(t, err)
}

// Fail case scenario testing where CNAME and TXT Record name are same
func TestUltraDNSProvider_ApplyChangesCNAME(t *testing.T) {
	changes := &plan.Changes{}
	mocked := mockUltraDNSRecord{nil}
	mockedDomain := mockUltraDNSZone{nil}

	provider := &UltraDNSProvider{
		client: udnssdk.Client{
			RRSets: &mocked,
			Zone:   &mockedDomain,
		},
	}

	changes.Create = []*endpoint.Endpoint{
		{DNSName: "test-ultradns-provider.com", Targets: endpoint.Targets{"1.1.1.1"}, RecordType: "CNAME"},
		{DNSName: "test-ultradns-provider.com", Targets: endpoint.Targets{"1.1.1.1"}, RecordType: "TXT"},
	}

	err := provider.ApplyChanges(context.Background(), changes)
	assert.NotNil(t, err)
}

// This will work if you would set the environment variables such as "ULTRADNS_INTEGRATION" and zone should be available "kubernetes-ultradns-provider-test.com"
func TestUltraDNSProvider_ApplyChanges_Integration(t *testing.T) {
	_, ok := os.LookupEnv("ULTRADNS_INTEGRATION")
	if !ok {
		log.Printf("Skipping test")
	} else {

		providerUltradns, err := NewUltraDNSProvider(endpoint.NewDomainFilter([]string{"kubernetes-ultradns-provider-test.com"}), false)
		changes := &plan.Changes{}
		changes.Create = []*endpoint.Endpoint{
			{DNSName: "kubernetes-ultradns-provider-test.com", Targets: endpoint.Targets{"1.1.1.1"}, RecordType: "A"},
			{DNSName: "ttl.kubernetes-ultradns-provider-test.com", Targets: endpoint.Targets{"2001:0db8:85a3:0000:0000:8a2e:0370:7334"}, RecordType: "AAAA", RecordTTL: 100},
		}

		err = providerUltradns.ApplyChanges(context.Background(), changes)
		assert.Nil(t, err)

		rrsetKey := udnssdk.RRSetKey{
			Zone: "kubernetes-ultradns-provider-test.com.",
			Name: "kubernetes-ultradns-provider-test.com.",
			Type: "A",
		}

		rrsets, _ := providerUltradns.client.RRSets.Select(rrsetKey)
		assert.Equal(t, rrsets[0].RData[0], "1.1.1.1")

		rrsetKey = udnssdk.RRSetKey{
			Zone: "kubernetes-ultradns-provider-test.com.",
			Name: "ttl.kubernetes-ultradns-provider-test.com.",
			Type: "AAAA",
		}

		rrsets, _ = providerUltradns.client.RRSets.Select(rrsetKey)
		assert.Equal(t, rrsets[0].RData[0], "2001:db8:85a3:0:0:8a2e:370:7334")

		changes = &plan.Changes{}
		changes.UpdateNew = []*endpoint.Endpoint{
			{DNSName: "kubernetes-ultradns-provider-test.com", Targets: endpoint.Targets{"1.1.2.2"}, RecordType: "A", RecordTTL: 100},
			{DNSName: "ttl.kubernetes-ultradns-provider-test.com", Targets: endpoint.Targets{"2001:0db8:85a3:0000:0000:8a2e:0370:7335"}, RecordType: "AAAA", RecordTTL: 100},
		}
		err = providerUltradns.ApplyChanges(context.Background(), changes)
		assert.Nil(t, err)

		rrsetKey = udnssdk.RRSetKey{
			Zone: "kubernetes-ultradns-provider-test.com.",
			Name: "kubernetes-ultradns-provider-test.com.",
			Type: "A",
		}

		rrsets, _ = providerUltradns.client.RRSets.Select(rrsetKey)
		assert.Equal(t, rrsets[0].RData[0], "1.1.2.2")

		rrsetKey = udnssdk.RRSetKey{
			Zone: "kubernetes-ultradns-provider-test.com.",
			Name: "ttl.kubernetes-ultradns-provider-test.com.",
			Type: "AAAA",
		}

		rrsets, _ = providerUltradns.client.RRSets.Select(rrsetKey)
		assert.Equal(t, rrsets[0].RData[0], "2001:db8:85a3:0:0:8a2e:370:7335")

		changes = &plan.Changes{}
		changes.Delete = []*endpoint.Endpoint{
			{DNSName: "ttl.kubernetes-ultradns-provider-test.com", Targets: endpoint.Targets{"2001:0db8:85a3:0000:0000:8a2e:0370:7335"}, RecordType: "AAAA", RecordTTL: 100},
			{DNSName: "kubernetes-ultradns-provider-test.com", Targets: endpoint.Targets{"1.1.2.2"}, RecordType: "A", RecordTTL: 100},
		}

		err = providerUltradns.ApplyChanges(context.Background(), changes)
		assert.Nil(t, err)

		resp, _ := providerUltradns.client.Do("GET", "zones/kubernetes-ultradns-provider-test.com./rrsets/AAAA/ttl.kubernetes-ultradns-provider-test.com.", nil, udnssdk.RRSetListDTO{})
		assert.Equal(t, resp.Status, "404 Not Found")

		resp, _ = providerUltradns.client.Do("GET", "zones/kubernetes-ultradns-provider-test.com./rrsets/A/kubernetes-ultradns-provider-test.com.", nil, udnssdk.RRSetListDTO{})
		assert.Equal(t, resp.Status, "404 Not Found")

	}
}

// This will work if you would set the environment variables such as "ULTRADNS_INTEGRATION" and zone should be available "kubernetes-ultradns-provider-test.com" for multiple target
func TestUltraDNSProvider_ApplyChanges_MultipleTarget_integeration(t *testing.T) {
	_, ok := os.LookupEnv("ULTRADNS_INTEGRATION")
	if !ok {
		log.Printf("Skipping test")
	} else {

		provider, err := NewUltraDNSProvider(endpoint.NewDomainFilter([]string{"kubernetes-ultradns-provider-test.com"}), false)
		changes := &plan.Changes{}
		changes.Create = []*endpoint.Endpoint{
			{DNSName: "kubernetes-ultradns-provider-test.com", Targets: endpoint.Targets{"1.1.1.1", "1.1.2.2"}, RecordType: "A"},
		}

		err = provider.ApplyChanges(context.Background(), changes)
		assert.Nil(t, err)

		rrsetKey := udnssdk.RRSetKey{
			Zone: "kubernetes-ultradns-provider-test.com.",
			Name: "kubernetes-ultradns-provider-test.com.",
			Type: "A",
		}

		rrsets, _ := provider.client.RRSets.Select(rrsetKey)
		assert.Equal(t, rrsets[0].RData, []string{"1.1.1.1", "1.1.2.2"})

		changes = &plan.Changes{}
		changes.UpdateNew = []*endpoint.Endpoint{{DNSName: "kubernetes-ultradns-provider-test.com", Targets: endpoint.Targets{"1.1.2.2", "192.168.0.24", "1.2.3.4"}, RecordType: "A", RecordTTL: 100}}

		err = provider.ApplyChanges(context.Background(), changes)
		assert.Nil(t, err)

		rrsetKey = udnssdk.RRSetKey{
			Zone: "kubernetes-ultradns-provider-test.com.",
			Name: "kubernetes-ultradns-provider-test.com.",
			Type: "A",
		}

		rrsets, _ = provider.client.RRSets.Select(rrsetKey)
		assert.Equal(t, rrsets[0].RData, []string{"1.1.2.2", "192.168.0.24", "1.2.3.4"})

		changes = &plan.Changes{}
		changes.UpdateNew = []*endpoint.Endpoint{{DNSName: "kubernetes-ultradns-provider-test.com", Targets: endpoint.Targets{"1.1.2.2"}, RecordType: "A", RecordTTL: 100}}

		err = provider.ApplyChanges(context.Background(), changes)

		assert.Nil(t, err)

		rrsetKey = udnssdk.RRSetKey{
			Zone: "kubernetes-ultradns-provider-test.com.",
			Name: "kubernetes-ultradns-provider-test.com.",
			Type: "A",
		}

		rrsets, _ = provider.client.RRSets.Select(rrsetKey)
		assert.Equal(t, rrsets[0].RData, []string{"1.1.2.2"})

		changes = &plan.Changes{}
		changes.Delete = []*endpoint.Endpoint{{DNSName: "kubernetes-ultradns-provider-test.com", Targets: endpoint.Targets{"1.1.2.2", "192.168.0.24"}, RecordType: "A"}}

		err = provider.ApplyChanges(context.Background(), changes)

		assert.Nil(t, err)

		resp, _ := provider.client.Do("GET", "zones/kubernetes-ultradns-provider-test.com./rrsets/A/kubernetes-ultradns-provider-test.com.", nil, udnssdk.RRSetListDTO{})
		assert.Equal(t, resp.Status, "404 Not Found")

	}
}

// Test case to check sbpool creation
func TestUltraDNSProvider_newSBPoolObjectCreation(t *testing.T) {
	mocked := mockUltraDNSRecord{nil}
	mockedDomain := mockUltraDNSZone{nil}

	provider := &UltraDNSProvider{
		client: udnssdk.Client{
			RRSets: &mocked,
			Zone:   &mockedDomain,
		},
	}
	sbpoolRDataList := []udnssdk.SBRDataInfo{}
	changes := &plan.Changes{}
	changes.UpdateNew = []*endpoint.Endpoint{{DNSName: "kubernetes-ultradns-provider-test.com.", Targets: endpoint.Targets{"1.1.2.2", "192.168.0.24"}, RecordType: "A", RecordTTL: 100}}
	changesList := &UltraDNSChanges{
		Action: "UPDATE",
		ResourceRecordSetUltraDNS: udnssdk.RRSet{
			RRType:    "A",
			OwnerName: "kubernetes-ultradns-provider-test.com.",
			RData:     []string{"1.1.2.2", "192.168.0.24"},
			TTL:       100,
		},
	}

	for range changesList.ResourceRecordSetUltraDNS.RData {

		rrdataInfo := udnssdk.SBRDataInfo{
			RunProbes: true,
			Priority:  1,
			State:     "NORMAL",
			Threshold: 1,
			Weight:    nil,
		}
		sbpoolRDataList = append(sbpoolRDataList, rrdataInfo)
	}
	sbPoolObject := udnssdk.SBPoolProfile{
		Context:     udnssdk.SBPoolSchema,
		Order:       "ROUND_ROBIN",
		Description: "kubernetes-ultradns-provider-test.com.",
		MaxActive:   2,
		MaxServed:   2,
		RDataInfo:   sbpoolRDataList,
		RunProbes:   true,
		ActOnProbes: true,
	}

	actualSBPoolObject, _ := provider.newSBPoolObjectCreation(context.Background(), changesList)
	assert.Equal(t, sbPoolObject, actualSBPoolObject)
}

// Testcase to check fail scenario for multiple AAAA targets
func TestUltraDNSProvider_MultipleTargetAAAA(t *testing.T) {
	_, ok := os.LookupEnv("ULTRADNS_INTEGRATION")
	if !ok {
		log.Printf("Skipping test")
	} else {
		_ = os.Setenv("ULTRADNS_POOL_TYPE", "sbpool")

		provider, _ := NewUltraDNSProvider(endpoint.NewDomainFilter([]string{"kubernetes-ultradns-provider-test.com"}), false)
		changes := &plan.Changes{}
		changes.Create = []*endpoint.Endpoint{
			{DNSName: "ttl.kubernetes-ultradns-provider-test.com", Targets: endpoint.Targets{"2001:0db8:85a3:0000:0000:8a2e:0370:7334", "2001:0db8:85a3:0000:0000:8a2e:0370:7335"}, RecordType: "AAAA", RecordTTL: 100},
		}
		err := provider.ApplyChanges(context.Background(), changes)
		assert.NotNilf(t, err, "We wanted it to fail since multiple AAAA targets are not allowed %s", "formatted")

		resp, _ := provider.client.Do("GET", "zones/kubernetes-ultradns-provider-test.com./rrsets/AAAA/ttl.kubernetes-ultradns-provider-test.com.", nil, udnssdk.RRSetListDTO{})
		assert.Equal(t, resp.Status, "404 Not Found")
		_ = os.Unsetenv("ULTRADNS_POOL_TYPE")
	}
}

// Testcase to check fail scenario for multiple AAAA targets
func TestUltraDNSProvider_MultipleTargetAAAARDPool(t *testing.T) {
	_, ok := os.LookupEnv("ULTRADNS_INTEGRATION")
	if !ok {
		log.Printf("Skipping test")
	} else {
		_ = os.Setenv("ULTRADNS_POOL_TYPE", "rdpool")
		provider, _ := NewUltraDNSProvider(endpoint.NewDomainFilter([]string{"kubernetes-ultradns-provider-test.com"}), false)
		changes := &plan.Changes{}
		changes.Create = []*endpoint.Endpoint{
			{DNSName: "ttl.kubernetes-ultradns-provider-test.com", Targets: endpoint.Targets{"2001:0db8:85a3:0000:0000:8a2e:0370:7334", "2001:0db8:85a3:0000:0000:8a2e:0370:7335"}, RecordType: "AAAA", RecordTTL: 100},
		}
		err := provider.ApplyChanges(context.Background(), changes)
		assert.Nilf(t, err, " multiple AAAA targets are allowed when pool is RDPool %s", "formatted")

		resp, _ := provider.client.Do("GET", "zones/kubernetes-ultradns-provider-test.com./rrsets/AAAA/ttl.kubernetes-ultradns-provider-test.com.", nil, udnssdk.RRSetListDTO{})
		assert.Equal(t, resp.Status, "200 OK")

		changes = &plan.Changes{}
		changes.Delete = []*endpoint.Endpoint{{DNSName: "ttl.kubernetes-ultradns-provider-test.com", Targets: endpoint.Targets{"2001:0db8:85a3:0000:0000:8a2e:0370:7334", "2001:0db8:85a3:0000:0000:8a2e:0370:7335"}, RecordType: "AAAA"}}

		err = provider.ApplyChanges(context.Background(), changes)

		assert.Nil(t, err)

		resp, _ = provider.client.Do("GET", "zones/kubernetes-ultradns-provider-test.com./rrsets/A/kubernetes-ultradns-provider-test.com.", nil, udnssdk.RRSetListDTO{})
		assert.Equal(t, resp.Status, "404 Not Found")

	}
}

// Test case to check multiple CNAME targets.
func TestUltraDNSProvider_MultipleTargetCNAME(t *testing.T) {
	_, ok := os.LookupEnv("ULTRADNS_INTEGRATION")
	if !ok {
		log.Printf("Skipping test")
	} else {
		provider, err := NewUltraDNSProvider(endpoint.NewDomainFilter([]string{"kubernetes-ultradns-provider-test.com"}), false)
		changes := &plan.Changes{}

		changes.Create = []*endpoint.Endpoint{
			{DNSName: "ttl.kubernetes-ultradns-provider-test.com", Targets: endpoint.Targets{"nginx.loadbalancer.com.", "nginx1.loadbalancer.com."}, RecordType: "CNAME", RecordTTL: 100},
		}
		err = provider.ApplyChanges(context.Background(), changes)

		assert.NotNilf(t, err, "We wanted it to fail since multiple CNAME targets are not allowed %s", "formatted")

		resp, _ := provider.client.Do("GET", "zones/kubernetes-ultradns-provider-test.com./rrsets/CNAME/kubernetes-ultradns-provider-test.com.", nil, udnssdk.RRSetListDTO{})
		assert.Equal(t, resp.Status, "404 Not Found")
	}
}

// Testing creation of RD Pool
func TestUltraDNSProvider_newRDPoolObjectCreation(t *testing.T) {
	mocked := mockUltraDNSRecord{nil}
	mockedDomain := mockUltraDNSZone{nil}

	provider := &UltraDNSProvider{
		client: udnssdk.Client{
			RRSets: &mocked,
			Zone:   &mockedDomain,
		},
	}
	changes := &plan.Changes{}
	changes.UpdateNew = []*endpoint.Endpoint{{DNSName: "kubernetes-ultradns-provider-test.com.", Targets: endpoint.Targets{"1.1.2.2", "192.168.0.24"}, RecordType: "A", RecordTTL: 100}}
	changesList := &UltraDNSChanges{
		Action: "UPDATE",
		ResourceRecordSetUltraDNS: udnssdk.RRSet{
			RRType:    "A",
			OwnerName: "kubernetes-ultradns-provider-test.com.",
			RData:     []string{"1.1.2.2", "192.168.0.24"},
			TTL:       100,
		},
	}
	rdPoolObject := udnssdk.RDPoolProfile{
		Context:     udnssdk.RDPoolSchema,
		Order:       "ROUND_ROBIN",
		Description: "kubernetes-ultradns-provider-test.com.",
	}

	actualRDPoolObject, _ := provider.newRDPoolObjectCreation(context.Background(), changesList)
	assert.Equal(t, rdPoolObject, actualRDPoolObject)
}

// Testing Failure scenarios over NewUltraDNS Provider
func TestNewUltraDNSProvider_FailCases(t *testing.T) {
	_ = os.Setenv("ULTRADNS_USERNAME", "")
	_ = os.Setenv("ULTRADNS_PASSWORD", "")
	_ = os.Setenv("ULTRADNS_BASEURL", "")
	_ = os.Setenv("ULTRADNS_ACCOUNTNAME", "")
	_ = os.Setenv("ULTRADNS_POOL_TYPE", "xyz")
	_, err := NewUltraDNSProvider(endpoint.NewDomainFilter([]string{"test-ultradns-provider.com"}), true)
	assert.NotNilf(t, err, "Pool Type other than given type not working %s", "formatted")

	_ = os.Setenv("ULTRADNS_USERNAME", "")
	_ = os.Setenv("ULTRADNS_PASSWORD", "")
	_ = os.Setenv("ULTRADNS_BASEURL", "")
	_ = os.Setenv("ULTRADNS_ACCOUNTNAME", "")
	_ = os.Setenv("ULTRADNS_ENABLE_PROBING", "adefg")
	_, err = NewUltraDNSProvider(endpoint.NewDomainFilter([]string{"test-ultradns-provider.com"}), true)
	assert.NotNilf(t, err, "Probe value other than given values not working  %s", "formatted")

	_ = os.Setenv("ULTRADNS_USERNAME", "")
	_ = os.Setenv("ULTRADNS_PASSWORD", "")
	_ = os.Setenv("ULTRADNS_BASEURL", "")
	_ = os.Setenv("ULTRADNS_ACCOUNTNAME", "")
	_ = os.Setenv("ULTRADNS_ENABLE_ACTONPROBE", "adefg")
	_, err = NewUltraDNSProvider(endpoint.NewDomainFilter([]string{"test-ultradns-provider.com"}), true)
	assert.NotNilf(t, err, "ActOnProbe value other than given values not working %s", "formatted")

	_ = os.Setenv("ULTRADNS_USERNAME", "")
	_ = os.Setenv("ULTRADNS_BASEURL", "")
	_ = os.Unsetenv("ULTRADNS_PASSWORD")
	_ = os.Setenv("ULTRADNS_ACCOUNTNAME", "")
	_, err = NewUltraDNSProvider(endpoint.NewDomainFilter([]string{"test-ultradns-provider.com"}), true)
	assert.NotNilf(t, err, "Expected to give error if password is not set %s", "formatted")

	_ = os.Setenv("ULTRADNS_USERNAME", "")
	_ = os.Setenv("ULTRADNS_PASSWORD", "")
	_ = os.Unsetenv("ULTRADNS_BASEURL")
	_ = os.Setenv("ULTRADNS_ACCOUNTNAME", "")
	_, err = NewUltraDNSProvider(endpoint.NewDomainFilter([]string{"test-ultradns-provider.com"}), true)
	assert.NotNilf(t, err, "Expected to give error if baseurl is not set %s", "formatted")

	_ = os.Setenv("ULTRADNS_USERNAME", "")
	_ = os.Setenv("ULTRADNS_BASEURL", "")
	_ = os.Setenv("ULTRADNS_PASSWORD", "")
	_ = os.Unsetenv("ULTRADNS_ACCOUNTNAME")
	_ = os.Unsetenv("ULTRADNS_ENABLE_ACTONPROBE")
	_ = os.Unsetenv("ULTRADNS_ENABLE_PROBING")
	_ = os.Unsetenv("ULTRADNS_POOL_TYPE")
	_, accounterr := NewUltraDNSProvider(endpoint.NewDomainFilter([]string{"test-ultradns-provider.com"}), true)
	assert.Nil(t, accounterr)
}

// Testing success scenarios for newly introduced environment variables
func TestNewUltraDNSProvider_NewEnvVariableSuccessCases(t *testing.T) {
	_ = os.Setenv("ULTRADNS_USERNAME", "")
	_ = os.Setenv("ULTRADNS_PASSWORD", "")
	_ = os.Setenv("ULTRADNS_BASEURL", "")
	_ = os.Setenv("ULTRADNS_ACCOUNTNAME", "")
	_ = os.Setenv("ULTRADNS_POOL_TYPE", "rdpool")
	_, err := NewUltraDNSProvider(endpoint.NewDomainFilter([]string{"test-ultradns-provider.com"}), true)
	assert.Nilf(t, err, "Pool Type not working in proper scenario %s", "formatted")

	_ = os.Setenv("ULTRADNS_USERNAME", "")
	_ = os.Setenv("ULTRADNS_PASSWORD", "")
	_ = os.Setenv("ULTRADNS_BASEURL", "")
	_ = os.Setenv("ULTRADNS_ACCOUNTNAME", "")
	_ = os.Setenv("ULTRADNS_ENABLE_PROBING", "false")
	_, err1 := NewUltraDNSProvider(endpoint.NewDomainFilter([]string{"test-ultradns-provider.com"}), true)
	assert.Nilf(t, err1, "Probe given value is  not working %s", "formatted")

	_ = os.Setenv("ULTRADNS_USERNAME", "")
	_ = os.Setenv("ULTRADNS_PASSWORD", "")
	_ = os.Setenv("ULTRADNS_BASEURL", "")
	_ = os.Setenv("ULTRADNS_ACCOUNTNAME", "")
	_ = os.Setenv("ULTRADNS_ENABLE_ACTONPROBE", "true")
	_, err2 := NewUltraDNSProvider(endpoint.NewDomainFilter([]string{"test-ultradns-provider.com"}), true)
	assert.Nilf(t, err2, "ActOnProbe given value is not working %s", "formatted")
}

// Base64 Bad string decoding scenario
func TestNewUltraDNSProvider_Base64DecodeFailcase(t *testing.T) {
	_ = os.Setenv("ULTRADNS_USERNAME", "")
	_ = os.Setenv("ULTRADNS_PASSWORD", "12345")
	_ = os.Setenv("ULTRADNS_BASEURL", "")
	_ = os.Setenv("ULTRADNS_ACCOUNTNAME", "")
	_ = os.Setenv("ULTRADNS_ENABLE_ACTONPROBE", "true")
	_, err := NewUltraDNSProvider(endpoint.NewDomainFilter([]string{"test-ultradns-provider.com"}), true)
	assert.NotNilf(t, err, "Base64 decode should fail in this case %s", "formatted")
}

func TestUltraDNSProvider_PoolConversionCase(t *testing.T) {
	_, ok := os.LookupEnv("ULTRADNS_INTEGRATION")
	if !ok {
		log.Printf("Skipping test")
	} else {
		// Creating SBPool Record
		_ = os.Setenv("ULTRADNS_POOL_TYPE", "sbpool")
		provider, _ := NewUltraDNSProvider(endpoint.NewDomainFilter([]string{"kubernetes-ultradns-provider-test.com"}), false)
		changes := &plan.Changes{}
		changes.Create = []*endpoint.Endpoint{{DNSName: "ttl.kubernetes-ultradns-provider-test.com", Targets: endpoint.Targets{"1.1.1.1", "1.2.3.4"}, RecordType: "A", RecordTTL: 100}}
		err := provider.ApplyChanges(context.Background(), changes)
		assert.Nilf(t, err, " multiple A record creation with SBPool %s", "formatted")

		resp, _ := provider.client.Do("GET", "zones/kubernetes-ultradns-provider-test.com./rrsets/A/ttl.kubernetes-ultradns-provider-test.com.", nil, udnssdk.RRSetListDTO{})
		assert.Equal(t, resp.Status, "200 OK")

		// Converting to RD Pool
		_ = os.Setenv("ULTRADNS_POOL_TYPE", "rdpool")
		provider, _ = NewUltraDNSProvider(endpoint.NewDomainFilter([]string{"kubernetes-ultradns-provider-test.com"}), false)
		changes = &plan.Changes{}
		changes.UpdateNew = []*endpoint.Endpoint{{DNSName: "ttl.kubernetes-ultradns-provider-test.com", Targets: endpoint.Targets{"1.1.1.1", "1.2.3.5"}, RecordType: "A"}}
		err = provider.ApplyChanges(context.Background(), changes)
		assert.Nil(t, err)
		resp, _ = provider.client.Do("GET", "zones/kubernetes-ultradns-provider-test.com./rrsets/A/ttl.kubernetes-ultradns-provider-test.com.", nil, udnssdk.RRSetListDTO{})
		assert.Equal(t, resp.Status, "200 OK")

		// Converting back to SB Pool
		_ = os.Setenv("ULTRADNS_POOL_TYPE", "sbpool")
		provider, _ = NewUltraDNSProvider(endpoint.NewDomainFilter([]string{"kubernetes-ultradns-provider-test.com"}), false)
		changes = &plan.Changes{}
		changes.UpdateNew = []*endpoint.Endpoint{{DNSName: "ttl.kubernetes-ultradns-provider-test.com", Targets: endpoint.Targets{"1.1.1.1", "1.2.3.4"}, RecordType: "A"}}
		err = provider.ApplyChanges(context.Background(), changes)
		assert.Nil(t, err)
		resp, _ = provider.client.Do("GET", "zones/kubernetes-ultradns-provider-test.com./rrsets/A/ttl.kubernetes-ultradns-provider-test.com.", nil, udnssdk.RRSetListDTO{})
		assert.Equal(t, resp.Status, "200 OK")

		// Deleting Record
		changes = &plan.Changes{}
		changes.Delete = []*endpoint.Endpoint{{DNSName: "ttl.kubernetes-ultradns-provider-test.com", Targets: endpoint.Targets{"1.1.1.1", "1.2.3.4"}, RecordType: "A"}}
		err = provider.ApplyChanges(context.Background(), changes)
		assert.Nil(t, err)
		resp, _ = provider.client.Do("GET", "zones/kubernetes-ultradns-provider-test.com./rrsets/A/kubernetes-ultradns-provider-test.com.", nil, udnssdk.RRSetListDTO{})
		assert.Equal(t, resp.Status, "404 Not Found")
	}
}

func TestUltraDNSProvider_DomainFilter(t *testing.T) {
	_, ok := os.LookupEnv("ULTRADNS_INTEGRATION")
	if !ok {
		log.Printf("Skipping test")
	} else {
		provider, _ := NewUltraDNSProvider(endpoint.NewDomainFilter([]string{"kubernetes-ultradns-provider-test.com", "kubernetes-ultradns-provider-test.com"}), true)
		zones, err := provider.Zones(context.Background())
		assert.Equal(t, zones[0].Properties.Name, "kubernetes-ultradns-provider-test.com.")
		assert.Equal(t, zones[1].Properties.Name, "kubernetes-ultradns-provider-test.com.")
		assert.Nilf(t, err, " Multiple domain filter failed %s", "formatted")

		provider, _ = NewUltraDNSProvider(endpoint.NewDomainFilter([]string{}), true)
		zones, err = provider.Zones(context.Background())
		assert.Nilf(t, err, " Multiple domain filter failed %s", "formatted")

	}
}

func TestUltraDNSProvider_DomainFiltersZonesFailCase(t *testing.T) {
	_, ok := os.LookupEnv("ULTRADNS_INTEGRATION")
	if !ok {
		log.Printf("Skipping test")
	} else {
		provider, _ := NewUltraDNSProvider(endpoint.NewDomainFilter([]string{"kubernetes-ultradns-provider-test.com", "kubernetes-uldsvdsvadvvdsvadvstradns-provider-test.com"}), true)
		_, err := provider.Zones(context.Background())
		assert.NotNilf(t, err, " Multiple domain filter failed %s", "formatted")
	}
}

// zones function with domain filter test scenario
func TestUltraDNSProvider_DomainFilterZonesMocked(t *testing.T) {
	mocked := mockUltraDNSZone{}
	provider := &UltraDNSProvider{
		client: udnssdk.Client{
			Zone: &mocked,
		},
		domainFilter: endpoint.NewDomainFilter([]string{"test-ultradns-provider.com."}),
	}

	zoneKey := &udnssdk.ZoneKey{
		Zone:        "test-ultradns-provider.com.",
		AccountName: "",
	}

	// When AccountName not given
	expected, _, _, err := provider.client.Zone.SelectWithOffsetWithLimit(zoneKey, 0, 1000)
	assert.Nil(t, err)
	zones, err := provider.Zones(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, reflect.DeepEqual(expected, zones), true)
	accountName = "teamrest"
	// When AccountName is set
	provider = &UltraDNSProvider{
		client: udnssdk.Client{
			Zone: &mocked,
		},
		domainFilter: endpoint.NewDomainFilter([]string{"test-ultradns-provider.com."}),
	}

	zoneKey = &udnssdk.ZoneKey{
		Zone:        "test-ultradns-provider.com.",
		AccountName: "teamrest",
	}

	expected, _, _, err = provider.client.Zone.SelectWithOffsetWithLimit(zoneKey, 0, 1000)
	assert.Nil(t, err)
	zones, err = provider.Zones(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, reflect.DeepEqual(expected, zones), true)

	// When zone is not given but account is provided
	provider = &UltraDNSProvider{
		client: udnssdk.Client{
			Zone: &mocked,
		},
	}

	zoneKey = &udnssdk.ZoneKey{
		AccountName: "teamrest",
	}

	expected, _, _, err = provider.client.Zone.SelectWithOffsetWithLimit(zoneKey, 0, 1000)
	assert.Nil(t, err)
	zones, err = provider.Zones(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, reflect.DeepEqual(expected, zones), true)
}
