/*
Copyright 2021 The Kubernetes Authors.
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

package gandi

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/go-gandi/go-gandi/domain"
	"github.com/go-gandi/go-gandi/livedns"
	"github.com/maxatome/go-testdeep/td"

	"github.com/stretchr/testify/assert"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

type MockAction struct {
	Name   string
	FQDN   string
	Record livedns.DomainRecord
}

type mockGandiClient struct {
	Actions         []MockAction
	FunctionToFail  string
	RecordsToReturn []livedns.DomainRecord
}

func mockGandiClientNew() *mockGandiClient {
	return &mockGandiClient{
		RecordsToReturn: testRecords(),
	}
}

func mockGandiClientNewWithRecords(recordsToReturn []livedns.DomainRecord) *mockGandiClient {
	return &mockGandiClient{
		RecordsToReturn: recordsToReturn,
	}
}

func mockGandiClientNewWithFailure(functionToFail string) *mockGandiClient {
	return &mockGandiClient{
		FunctionToFail:  functionToFail,
		RecordsToReturn: testRecords(),
	}
}

const (
	domainUriPrefix  = "https://api.gandi.net/v5/domain/domains/"
	exampleDotComUri = domainUriPrefix + "example.com"
	exampleDotNetUri = domainUriPrefix + "example.net"
)

func testRecords() []livedns.DomainRecord {
	return []livedns.DomainRecord{
		{
			RrsetType:   endpoint.RecordTypeCNAME,
			RrsetTTL:    600,
			RrsetName:   "@",
			RrsetHref:   exampleDotComUri + "/records/%40/A",
			RrsetValues: []string{"192.168.0.1"},
		},
		{
			RrsetType:   endpoint.RecordTypeCNAME,
			RrsetTTL:    600,
			RrsetName:   "www",
			RrsetHref:   exampleDotComUri + "/records/www/CNAME",
			RrsetValues: []string{"lb.example.com"},
		},
		{
			RrsetType:   endpoint.RecordTypeA,
			RrsetTTL:    600,
			RrsetName:   "test",
			RrsetHref:   exampleDotComUri + "/records/test/A",
			RrsetValues: []string{"192.168.0.2"},
		},
	}
}

// Mock all methods

func (m *mockGandiClient) GetDomainRecords(fqdn string) (records []livedns.DomainRecord, err error) {
	m.Actions = append(m.Actions, MockAction{
		Name: "GetDomainRecords",
		FQDN: fqdn,
	})
	if m.FunctionToFail == "GetDomainRecords" {
		return nil, fmt.Errorf("injected error")
	}

	return m.RecordsToReturn, err
}

func (m *mockGandiClient) CreateDomainRecord(fqdn, name, recordtype string, ttl int, values []string) (response standardResponse, err error) {
	m.Actions = append(m.Actions, MockAction{
		Name: "CreateDomainRecord",
		FQDN: fqdn,
		Record: livedns.DomainRecord{
			RrsetType:   recordtype,
			RrsetTTL:    ttl,
			RrsetName:   name,
			RrsetValues: values,
		},
	})
	if m.FunctionToFail == "CreateDomainRecord" {
		return standardResponse{}, fmt.Errorf("injected error")
	}

	return standardResponse{}, nil
}

func (m *mockGandiClient) DeleteDomainRecord(fqdn, name, recordtype string) (err error) {
	m.Actions = append(m.Actions, MockAction{
		Name: "DeleteDomainRecord",
		FQDN: fqdn,
		Record: livedns.DomainRecord{
			RrsetType: recordtype,
			RrsetName: name,
		},
	})

	if m.FunctionToFail == "DeleteDomainRecord" {
		return fmt.Errorf("injected error")
	}

	return nil
}

func (m *mockGandiClient) UpdateDomainRecordByNameAndType(fqdn, name, recordtype string, ttl int, values []string) (response standardResponse, err error) {
	m.Actions = append(m.Actions, MockAction{
		Name: "UpdateDomainRecordByNameAndType",
		FQDN: fqdn,
		Record: livedns.DomainRecord{
			RrsetType:   recordtype,
			RrsetTTL:    ttl,
			RrsetName:   name,
			RrsetValues: values,
		},
	})

	if m.FunctionToFail == "UpdateDomainRecordByNameAndType" {
		return standardResponse{}, fmt.Errorf("injected error")
	}

	return standardResponse{}, nil
}

func (m *mockGandiClient) ListDomains() (domains []domain.ListResponse, err error) {
	m.Actions = append(m.Actions, MockAction{
		Name: "ListDomains",
	})

	if m.FunctionToFail == "ListDomains" {
		return []domain.ListResponse{}, fmt.Errorf("injected error")
	}

	return []domain.ListResponse{
		{
			FQDN:        "example.com",
			FQDNUnicode: "example.com",
			Href:        exampleDotComUri,
			ID:          "b3e9c271-1c29-4441-97d9-bc021a7ac7c3",
			NameServer: &domain.NameServerConfig{
				Current: gandiLiveDNSProvider,
			},
			TLD: "com",
		},
		{
			FQDN:        "example.net",
			FQDNUnicode: "example.net",
			Href:        exampleDotNetUri,
			ID:          "dc78c1d8-6143-4edb-93bc-3a20d8bc3570",
			NameServer: &domain.NameServerConfig{
				Current: "other",
			},
			TLD: "net",
		},
	}, nil
}

// Tests

func TestNewGandiProvider(t *testing.T) {
	_ = os.Setenv("GANDI_KEY", "myGandiKey")
	provider, err := NewGandiProvider(context.Background(), endpoint.NewDomainFilter([]string{"example.com"}), true)
	if err != nil {
		t.Errorf("failed : %s", err)
	}
	assert.Equal(t, true, provider.DryRun)

	_ = os.Setenv("GANDI_SHARING_ID", "aSharingId")
	provider, err = NewGandiProvider(context.Background(), endpoint.NewDomainFilter([]string{"example.com"}), false)
	if err != nil {
		t.Errorf("failed : %s", err)
	}
	assert.Equal(t, false, provider.DryRun)

	_ = os.Unsetenv("GANDI_KEY")
	_, err = NewGandiProvider(context.Background(), endpoint.NewDomainFilter([]string{"example.com"}), true)
	if err == nil {
		t.Errorf("expected to fail")
	}
}

func TestGandiProvider_TestData(t *testing.T) {
	mockedClient := mockGandiClientNew()

	// Check test zone data is ok
	expectedZonesAnswer := []domain.ListResponse{
		{
			FQDN:        "example.com",
			FQDNUnicode: "example.com",
			Href:        exampleDotComUri,
			ID:          "b3e9c271-1c29-4441-97d9-bc021a7ac7c3",
			NameServer: &domain.NameServerConfig{
				Current: gandiLiveDNSProvider,
			},
			TLD: "com",
		},
		{
			FQDN:        "example.net",
			FQDNUnicode: "example.net",
			Href:        exampleDotNetUri,
			ID:          "dc78c1d8-6143-4edb-93bc-3a20d8bc3570",
			NameServer: &domain.NameServerConfig{
				Current: "other",
			},
			TLD: "net",
		},
	}

	testingZonesAnswer, err := mockedClient.ListDomains()
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}

	if !reflect.DeepEqual(expectedZonesAnswer, testingZonesAnswer) {
		t.Errorf("should be equal, %s", err)
	}

	// Check test record data is ok
	expectedRecordsAnswer := []livedns.DomainRecord{
		{
			RrsetType:   endpoint.RecordTypeCNAME,
			RrsetTTL:    600,
			RrsetName:   "@",
			RrsetHref:   exampleDotComUri + "/records/%40/A",
			RrsetValues: []string{"192.168.0.1"},
		},
		{
			RrsetType:   endpoint.RecordTypeCNAME,
			RrsetTTL:    600,
			RrsetName:   "www",
			RrsetHref:   exampleDotComUri + "/records/www/CNAME",
			RrsetValues: []string{"lb.example.com"},
		},
		{
			RrsetType:   endpoint.RecordTypeA,
			RrsetTTL:    600,
			RrsetName:   "test",
			RrsetHref:   exampleDotComUri + "/records/test/A",
			RrsetValues: []string{"192.168.0.2"},
		},
	}

	testingRecordsAnswer, err := mockedClient.GetDomainRecords("example.com")
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}

	if !reflect.DeepEqual(expectedRecordsAnswer, testingRecordsAnswer) {
		t.Errorf("should be equal, %s", err)
	}
}

func TestGandiProvider_Records(t *testing.T) {
	mockedClient := mockGandiClientNew()

	mockedProvider := &GandiProvider{
		DomainClient:  mockedClient,
		LiveDNSClient: mockedClient,
	}

	expectedActions := []MockAction{
		{
			Name: "ListDomains",
		},
		{
			Name: "GetDomainRecords",
			FQDN: "example.com",
		},
	}

	endpoints, err := mockedProvider.Records(context.Background())
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
	assert.Equal(t, 3, len(endpoints))
	fmt.Printf("%+v\n", endpoints[0].DNSName)
	assert.Equal(t, "example.com", endpoints[0].DNSName)
	assert.Equal(t, endpoint.RecordTypeCNAME, endpoints[0].RecordType)
	td.Cmp(t, expectedActions, mockedClient.Actions)
}

func TestGandiProvider_RecordsAppliesDomainFilter(t *testing.T) {
	mockedClient := mockGandiClientNew()

	mockedProvider := &GandiProvider{
		DomainClient:  mockedClient,
		LiveDNSClient: mockedClient,
		domainFilter:  endpoint.NewDomainFilterWithExclusions([]string{}, []string{"example.com"}),
	}

	expectedActions := []MockAction{
		{
			Name: "ListDomains",
		},
	}

	endpoints, err := mockedProvider.Records(context.Background())
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
	assert.Equal(t, 0, len(endpoints))
	td.Cmp(t, expectedActions, mockedClient.Actions)
}

func TestGandiProvider_RecordsErrorOnMultipleValues(t *testing.T) {
	mockedClient := mockGandiClientNewWithRecords([]livedns.DomainRecord{
		{
			RrsetValues: []string{"foo", "bar"},
			RrsetType:   endpoint.RecordTypeCNAME,
		},
	})

	mockedProvider := &GandiProvider{
		DomainClient:  mockedClient,
		LiveDNSClient: mockedClient,
	}

	expectedActions := []MockAction{
		{
			Name: "ListDomains",
		},
		{
			Name: "GetDomainRecords",
			FQDN: "example.com",
		},
	}

	endpoints, err := mockedProvider.Records(context.Background())
	if err == nil {
		t.Errorf("expected to fail")
	}
	assert.Equal(t, 0, len(endpoints))
	assert.True(t, strings.HasPrefix(err.Error(), "can't handle multiple values for rrset"))
	td.Cmp(t, expectedActions, mockedClient.Actions)
}

func TestGandiProvider_ApplyChangesEmpty(t *testing.T) {
	changes := &plan.Changes{}
	mockedClient := mockGandiClientNew()
	mockedProvider := &GandiProvider{
		DomainClient:  mockedClient,
		LiveDNSClient: mockedClient,
	}

	err := mockedProvider.ApplyChanges(context.Background(), changes)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}

	if mockedClient.Actions != nil {
		t.Error("expected no changes")
	}
}

func TestGandiProvider_ApplyChanges(t *testing.T) {
	changes := &plan.Changes{}
	mockedClient := mockGandiClientNew()
	mockedProvider := &GandiProvider{
		DomainClient:  mockedClient,
		LiveDNSClient: mockedClient,
	}

	changes.Create = []*endpoint.Endpoint{{DNSName: "test2.example.com", Targets: endpoint.Targets{"target"}, RecordType: "A", RecordTTL: 666}}
	changes.UpdateNew = []*endpoint.Endpoint{{DNSName: "test3.example.com", Targets: endpoint.Targets{"target-new"}, RecordType: "A", RecordTTL: 777}}
	changes.Delete = []*endpoint.Endpoint{{DNSName: "test4.example.com", Targets: endpoint.Targets{"target-other"}, RecordType: "A"}}

	err := mockedProvider.ApplyChanges(context.Background(), changes)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}

	td.Cmp(t, mockedClient.Actions, []MockAction{
		{
			Name: "ListDomains",
		},
		{
			Name: "CreateDomainRecord",
			FQDN: "example.com",
			Record: livedns.DomainRecord{
				RrsetType:   endpoint.RecordTypeA,
				RrsetName:   "test2",
				RrsetValues: []string{"target"},
				RrsetTTL:    666,
			},
		},
		{
			Name: "UpdateDomainRecordByNameAndType",
			FQDN: "example.com",
			Record: livedns.DomainRecord{
				RrsetType:   endpoint.RecordTypeA,
				RrsetName:   "test3",
				RrsetValues: []string{"target-new"},
				RrsetTTL:    777,
			},
		},
		{
			Name: "DeleteDomainRecord",
			FQDN: "example.com",
			Record: livedns.DomainRecord{
				RrsetType: endpoint.RecordTypeA,
				RrsetName: "test4",
			},
		},
	})
}

func TestGandiProvider_ApplyChangesSkipsNonManaged(t *testing.T) {
	changes := &plan.Changes{}
	mockedClient := mockGandiClientNew()
	mockedProvider := &GandiProvider{
		DomainClient:  mockedClient,
		LiveDNSClient: mockedClient,
	}

	changes.Create = []*endpoint.Endpoint{{DNSName: "example.net", Targets: endpoint.Targets{"target"}}}
	changes.UpdateNew = []*endpoint.Endpoint{{DNSName: "test.example.net", Targets: endpoint.Targets{"target-new"}, RecordType: "A", RecordTTL: 777}}
	changes.Delete = []*endpoint.Endpoint{{DNSName: "test2.example.net", Targets: endpoint.Targets{"target"}, RecordType: "A"}}

	err := mockedProvider.ApplyChanges(context.Background(), changes)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}

	td.Cmp(t, mockedClient.Actions, []MockAction{
		{
			Name: "ListDomains",
		},
	})
}

func TestGandiProvider_ApplyChangesCreateUpdateCname(t *testing.T) {
	changes := &plan.Changes{}
	mockedClient := mockGandiClientNew()
	mockedProvider := &GandiProvider{
		DomainClient:  mockedClient,
		LiveDNSClient: mockedClient,
	}

	changes.Create = []*endpoint.Endpoint{
		{DNSName: "test-cname.example.com", Targets: endpoint.Targets{"target"}, RecordTTL: 666, RecordType: "CNAME"},
	}
	changes.UpdateNew = []*endpoint.Endpoint{{DNSName: "test-cname2.example.com", Targets: endpoint.Targets{"target-new"}, RecordType: "CNAME", RecordTTL: 777}}

	err := mockedProvider.ApplyChanges(context.Background(), changes)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}

	td.Cmp(t, mockedClient.Actions, []MockAction{
		{
			Name: "ListDomains",
		},
		{
			Name: "CreateDomainRecord",
			FQDN: "example.com",
			Record: livedns.DomainRecord{
				RrsetType:   endpoint.RecordTypeCNAME,
				RrsetName:   "test-cname",
				RrsetValues: []string{"target."},
				RrsetTTL:    666,
			},
		},
		{
			Name: "UpdateDomainRecordByNameAndType",
			FQDN: "example.com",
			Record: livedns.DomainRecord{
				RrsetType:   endpoint.RecordTypeCNAME,
				RrsetName:   "test-cname2",
				RrsetValues: []string{"target-new."},
				RrsetTTL:    777,
			},
		},
	})
}

func TestGandiProvider_ApplyChangesCreateEmpty(t *testing.T) {
	changes := &plan.Changes{}
	mockedClient := mockGandiClientNew()
	mockedProvider := &GandiProvider{
		DomainClient:  mockedClient,
		LiveDNSClient: mockedClient,
	}

	changes.Create = []*endpoint.Endpoint{
		{DNSName: "example.com", Targets: endpoint.Targets{"target"}, RecordTTL: 666, RecordType: "A"},
	}
	changes.UpdateNew = []*endpoint.Endpoint{}

	err := mockedProvider.ApplyChanges(context.Background(), changes)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}

	td.Cmp(t, mockedClient.Actions, []MockAction{
		{
			Name: "ListDomains",
		},
		{
			Name: "CreateDomainRecord",
			FQDN: "example.com",
			Record: livedns.DomainRecord{
				RrsetType:   endpoint.RecordTypeA,
				RrsetName:   "@",
				RrsetValues: []string{"target"},
				RrsetTTL:    666,
			},
		},
	})
}

func TestGandiProvider_ApplyChangesRespectsDryRun(t *testing.T) {
	changes := &plan.Changes{}
	mockedClient := mockGandiClientNew()
	mockedProvider := &GandiProvider{
		DomainClient:  mockedClient,
		LiveDNSClient: mockedClient,
		DryRun:        true,
	}

	changes.Create = []*endpoint.Endpoint{
		{DNSName: "foo.example.com", Targets: endpoint.Targets{"target"}, RecordTTL: 666, RecordType: "A"},
	}
	changes.UpdateNew = []*endpoint.Endpoint{
		{DNSName: "bar.example.com", Targets: endpoint.Targets{"target"}, RecordTTL: 666, RecordType: "A"},
	}
	changes.Delete = []*endpoint.Endpoint{
		{DNSName: "baz.example.com", Targets: endpoint.Targets{"target"}, RecordTTL: 666, RecordType: "A"},
	}

	err := mockedProvider.ApplyChanges(context.Background(), changes)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}

	td.Cmp(t, mockedClient.Actions, []MockAction{
		{
			Name: "ListDomains",
		},
	})
}

func TestGandiProvider_ApplyChangesErrorListDomains(t *testing.T) {
	changes := &plan.Changes{}
	mockedClient := mockGandiClientNewWithFailure("ListDomains")
	mockedProvider := &GandiProvider{
		DomainClient:  mockedClient,
		LiveDNSClient: mockedClient,
	}

	changes.Create = []*endpoint.Endpoint{
		{DNSName: "foo.example.com", Targets: endpoint.Targets{"target"}, RecordTTL: 666, RecordType: "A"},
	}
	changes.UpdateNew = []*endpoint.Endpoint{
		{DNSName: "bar.example.com", Targets: endpoint.Targets{"target"}, RecordTTL: 666, RecordType: "A"},
	}
	changes.Delete = []*endpoint.Endpoint{
		{DNSName: "baz.example.com", Targets: endpoint.Targets{"target"}, RecordTTL: 666, RecordType: "A"},
	}

	err := mockedProvider.ApplyChanges(context.Background(), changes)
	if err == nil {
		t.Error("should have failed")
	}

	td.Cmp(t, mockedClient.Actions, []MockAction{
		{
			Name: "ListDomains",
		},
	})
}

func TestGandiProvider_ApplyChangesErrorCreate(t *testing.T) {
	changes := &plan.Changes{}
	mockedClient := mockGandiClientNewWithFailure("CreateDomainRecord")
	mockedProvider := &GandiProvider{
		DomainClient:  mockedClient,
		LiveDNSClient: mockedClient,
	}

	changes.Create = []*endpoint.Endpoint{
		{DNSName: "foo.example.com", Targets: endpoint.Targets{"target"}, RecordTTL: 666, RecordType: "A"},
	}
	changes.UpdateNew = []*endpoint.Endpoint{
		{DNSName: "bar.example.com", Targets: endpoint.Targets{"target"}, RecordTTL: 666, RecordType: "A"},
	}
	changes.Delete = []*endpoint.Endpoint{
		{DNSName: "baz.example.com", Targets: endpoint.Targets{"target"}, RecordTTL: 666, RecordType: "A"},
	}

	err := mockedProvider.ApplyChanges(context.Background(), changes)
	if err == nil {
		t.Error("should have failed")
	}

	td.Cmp(t, mockedClient.Actions, []MockAction{
		{
			Name: "ListDomains",
		},
		{
			Name: "CreateDomainRecord",
			FQDN: "example.com",
			Record: livedns.DomainRecord{
				RrsetType:   endpoint.RecordTypeA,
				RrsetName:   "foo",
				RrsetValues: []string{"target"},
				RrsetTTL:    666,
			},
		},
	})
}

func TestGandiProvider_ApplyChangesErrorUpdate(t *testing.T) {
	changes := &plan.Changes{}
	mockedClient := mockGandiClientNewWithFailure("UpdateDomainRecordByNameAndType")
	mockedProvider := &GandiProvider{
		DomainClient:  mockedClient,
		LiveDNSClient: mockedClient,
	}

	changes.Create = []*endpoint.Endpoint{
		{DNSName: "foo.example.com", Targets: endpoint.Targets{"target"}, RecordTTL: 666, RecordType: "A"},
	}
	changes.UpdateNew = []*endpoint.Endpoint{
		{DNSName: "bar.example.com", Targets: endpoint.Targets{"target"}, RecordTTL: 666, RecordType: "A"},
	}
	changes.Delete = []*endpoint.Endpoint{
		{DNSName: "baz.example.com", Targets: endpoint.Targets{"target"}, RecordTTL: 666, RecordType: "A"},
	}

	err := mockedProvider.ApplyChanges(context.Background(), changes)
	if err == nil {
		t.Error("should have failed")
	}

	td.Cmp(t, mockedClient.Actions, []MockAction{
		{
			Name: "ListDomains",
		},
		{
			Name: "CreateDomainRecord",
			FQDN: "example.com",
			Record: livedns.DomainRecord{
				RrsetType:   endpoint.RecordTypeA,
				RrsetName:   "foo",
				RrsetValues: []string{"target"},
				RrsetTTL:    666,
			},
		},
		{
			Name: "UpdateDomainRecordByNameAndType",
			FQDN: "example.com",
			Record: livedns.DomainRecord{
				RrsetType:   endpoint.RecordTypeA,
				RrsetName:   "bar",
				RrsetValues: []string{"target"},
				RrsetTTL:    666,
			},
		},
	})
}

func TestGandiProvider_ApplyChangesErrorDelete(t *testing.T) {
	changes := &plan.Changes{}
	mockedClient := mockGandiClientNewWithFailure("DeleteDomainRecord")
	mockedProvider := &GandiProvider{
		DomainClient:  mockedClient,
		LiveDNSClient: mockedClient,
	}

	changes.Create = []*endpoint.Endpoint{
		{DNSName: "foo.example.com", Targets: endpoint.Targets{"target"}, RecordTTL: 666, RecordType: "A"},
	}
	changes.UpdateNew = []*endpoint.Endpoint{
		{DNSName: "bar.example.com", Targets: endpoint.Targets{"target"}, RecordTTL: 666, RecordType: "A"},
	}
	changes.Delete = []*endpoint.Endpoint{
		{DNSName: "baz.example.com", Targets: endpoint.Targets{"target"}, RecordTTL: 666, RecordType: "A"},
	}

	err := mockedProvider.ApplyChanges(context.Background(), changes)
	if err == nil {
		t.Error("should have failed")
	}

	td.Cmp(t, mockedClient.Actions, []MockAction{
		{
			Name: "ListDomains",
		},
		{
			Name: "CreateDomainRecord",
			FQDN: "example.com",
			Record: livedns.DomainRecord{
				RrsetType:   endpoint.RecordTypeA,
				RrsetName:   "foo",
				RrsetValues: []string{"target"},
				RrsetTTL:    666,
			},
		},
		{
			Name: "UpdateDomainRecordByNameAndType",
			FQDN: "example.com",
			Record: livedns.DomainRecord{
				RrsetType:   endpoint.RecordTypeA,
				RrsetName:   "bar",
				RrsetValues: []string{"target"},
				RrsetTTL:    666,
			},
		},
		{
			Name: "DeleteDomainRecord",
			FQDN: "example.com",
			Record: livedns.DomainRecord{
				RrsetType: endpoint.RecordTypeA,
				RrsetName: "baz",
			},
		},
	})
}
