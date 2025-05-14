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
	"testing"

	"github.com/go-gandi/go-gandi/domain"
	"github.com/go-gandi/go-gandi/livedns"
	"github.com/maxatome/go-testdeep/td"
	"github.com/stretchr/testify/assert"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	"sigs.k8s.io/external-dns/plan"
)

type MockAction struct {
	Name   string
	FQDN   string
	Record livedns.DomainRecord
}

type mockGandiClient struct {
	Actions         []MockAction
	FunctionToFail  string `default:""`
	RecordsToReturn []livedns.DomainRecord
}

const (
	domainUriPrefix  = "https://api.gandi.net/v5/domain/domains/"
	exampleDotComUri = domainUriPrefix + "example.com"
	exampleDotNetUri = domainUriPrefix + "example.net"
)

// Mock all methods

func (m *mockGandiClient) GetDomainRecords(fqdn string) (records []livedns.DomainRecord, err error) {
	m.Actions = append(m.Actions, MockAction{
		Name: "GetDomainRecords",
		FQDN: fqdn,
	})

	if m.FunctionToFail == "GetDomainRecords" {
		return nil, fmt.Errorf("injected error")
	}

	return m.RecordsToReturn, nil
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
		// Tests are using example.com
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
		// example.net returns "other" as NameServer, so it is ignored
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

	_ = os.Setenv("GANDI_PAT", "myGandiPAT")
	provider, err = NewGandiProvider(context.Background(), endpoint.NewDomainFilter([]string{"example.com"}), true)
	if err != nil {
		t.Errorf("failed : %s", err)
	}
	assert.Equal(t, true, provider.DryRun)

	_ = os.Unsetenv("GANDI_KEY")
	provider, err = NewGandiProvider(context.Background(), endpoint.NewDomainFilter([]string{"example.com"}), true)
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

	_ = os.Unsetenv("GANDI_PAT")
	_, err = NewGandiProvider(context.Background(), endpoint.NewDomainFilter([]string{"example.com"}), true)
	if err == nil {
		t.Errorf("expected to fail")
	}
}

func TestGandiProvider_RecordsReturnsCorrectEndpoints(t *testing.T) {
	mockedClient := &mockGandiClient{
		RecordsToReturn: []livedns.DomainRecord{
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
		},
	}

	mockedProvider := &GandiProvider{
		DomainClient:  mockedClient,
		LiveDNSClient: mockedClient,
	}

	actualEndpoints, err := mockedProvider.Records(context.Background())
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}

	expectedEndpoints := []*endpoint.Endpoint{
		{
			RecordType: endpoint.RecordTypeCNAME,
			DNSName:    "example.com",
			Targets:    endpoint.Targets{"192.168.0.1"},
			RecordTTL:  600,
		},
		{
			RecordType: endpoint.RecordTypeCNAME,
			DNSName:    "www.example.com",
			Targets:    endpoint.Targets{"lb.example.com"},
			RecordTTL:  600,
		},
		{
			RecordType: endpoint.RecordTypeA,
			DNSName:    "test.example.com",
			Targets:    endpoint.Targets{"192.168.0.2"},
			RecordTTL:  600,
		},
	}

	assert.Equal(t, len(expectedEndpoints), len(actualEndpoints))
	// we could use testutils.SameEndpoints (plural), but this makes it easier to identify which case is failing
	for i := range actualEndpoints {
		if !testutils.SameEndpoint(expectedEndpoints[i], actualEndpoints[i]) {
			t.Errorf("should be equal, expected:%v <> actual:%v", expectedEndpoints[i], actualEndpoints[i])

		}
	}
}

func TestGandiProvider_RecordsOnFilteredDomainsShouldYieldNoEndpoints(t *testing.T) {
	mockedClient := &mockGandiClient{
		RecordsToReturn: []livedns.DomainRecord{
			{
				RrsetType:   endpoint.RecordTypeCNAME,
				RrsetTTL:    600,
				RrsetName:   "@",
				RrsetHref:   exampleDotComUri + "/records/test/MX",
				RrsetValues: []string{"192.168.0.1"},
			},
		},
	}

	mockedProvider := &GandiProvider{
		DomainClient:  mockedClient,
		LiveDNSClient: mockedClient,
		domainFilter:  endpoint.NewDomainFilterWithExclusions([]string{}, []string{"example.com"}),
	}

	endpoints, _ := mockedProvider.Records(context.Background())
	assert.Empty(t, endpoints)
}

func TestGandiProvider_RecordsWithUnsupportedTypesAreNotReturned(t *testing.T) {
	mockedClient := &mockGandiClient{
		RecordsToReturn: []livedns.DomainRecord{
			{
				RrsetType:   "MX",
				RrsetTTL:    360,
				RrsetName:   "@",
				RrsetHref:   exampleDotComUri + "/records/%40/A",
				RrsetValues: []string{"smtp.example.com"},
			},
		},
	}

	mockedProvider := &GandiProvider{
		DomainClient:  mockedClient,
		LiveDNSClient: mockedClient,
	}

	endpoints, _ := mockedProvider.Records(context.Background())
	assert.Empty(t, endpoints)
}

func TestGandiProvider_ApplyChangesMakesExpectedAPICalls(t *testing.T) {
	changes := &plan.Changes{}
	mockedClient := &mockGandiClient{}
	mockedProvider := &GandiProvider{
		DomainClient:  mockedClient,
		LiveDNSClient: mockedClient,
	}

	changes.Create = []*endpoint.Endpoint{
		{
			DNSName:    "test2.example.com",
			Targets:    endpoint.Targets{"192.168.0.1"},
			RecordType: "A",
			RecordTTL:  666,
		},
	}
	changes.UpdateNew = []*endpoint.Endpoint{
		{
			DNSName:    "test3.example.com",
			Targets:    endpoint.Targets{"192.168.0.2"},
			RecordType: "A",
			RecordTTL:  777,
		},
		{
			DNSName:    "example.com.example.com",
			Targets:    endpoint.Targets{"lb-2.example.net"},
			RecordType: "CNAME",
			RecordTTL:  777,
		},
	}
	changes.Delete = []*endpoint.Endpoint{
		{
			DNSName:    "test4.example.com",
			Targets:    endpoint.Targets{"192.168.0.3"},
			RecordType: "A",
		},
	}

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
				RrsetValues: []string{"192.168.0.1"},
				RrsetTTL:    666,
			},
		},
		{
			Name: "UpdateDomainRecordByNameAndType",
			FQDN: "example.com",
			Record: livedns.DomainRecord{
				RrsetType:   endpoint.RecordTypeA,
				RrsetName:   "test3",
				RrsetValues: []string{"192.168.0.2"},
				RrsetTTL:    777,
			},
		},
		{
			Name: "UpdateDomainRecordByNameAndType",
			FQDN: "example.com",
			Record: livedns.DomainRecord{
				RrsetType:   endpoint.RecordTypeCNAME,
				RrsetName:   "example.com",
				RrsetValues: []string{"lb-2.example.net."},
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

func TestGandiProvider_ApplyChangesRespectsDryRun(t *testing.T) {
	changes := &plan.Changes{}
	mockedClient := &mockGandiClient{}
	mockedProvider := &GandiProvider{
		DryRun:        true,
		DomainClient:  mockedClient,
		LiveDNSClient: mockedClient,
	}

	changes.Create = []*endpoint.Endpoint{{DNSName: "test2.example.com", Targets: endpoint.Targets{"192.168.0.1"}, RecordType: "A", RecordTTL: 666}}
	changes.UpdateNew = []*endpoint.Endpoint{{DNSName: "test3.example.com", Targets: endpoint.Targets{"192.168.0.2"}, RecordType: "A", RecordTTL: 777}}
	changes.Delete = []*endpoint.Endpoint{{DNSName: "test4.example.com", Targets: endpoint.Targets{"192.168.0.3"}, RecordType: "A"}}

	mockedProvider.ApplyChanges(context.Background(), changes)

	td.Cmp(t, mockedClient.Actions, []MockAction{
		{
			Name: "ListDomains",
		},
	})
}

func TestGandiProvider_ApplyChangesWithEmptyResultDoesNothing(t *testing.T) {
	changes := &plan.Changes{}
	mockedClient := &mockGandiClient{}
	mockedProvider := &GandiProvider{
		DomainClient:  mockedClient,
		LiveDNSClient: mockedClient,
	}

	mockedProvider.ApplyChanges(context.Background(), changes)

	assert.Empty(t, mockedClient.Actions)
}

func TestGandiProvider_ApplyChangesWithUnknownDomainDoesNoUpdate(t *testing.T) {
	changes := &plan.Changes{}
	mockedClient := &mockGandiClient{}
	mockedProvider := &GandiProvider{
		DomainClient:  mockedClient,
		LiveDNSClient: mockedClient,
	}

	changes.Create = []*endpoint.Endpoint{
		{
			DNSName:    "test.example.net",
			Targets:    endpoint.Targets{"192.168.0.1"},
			RecordType: "A",
			RecordTTL:  666,
		},
	}

	mockedProvider.ApplyChanges(context.Background(), changes)

	td.Cmp(t, mockedClient.Actions, []MockAction{
		{
			Name: "ListDomains",
		},
	})
}

func TestGandiProvider_FailingCases(t *testing.T) {
	changes := &plan.Changes{}
	changes.Create = []*endpoint.Endpoint{{DNSName: "test2.example.com", Targets: endpoint.Targets{"192.168.0.1"}, RecordType: "A", RecordTTL: 666}}
	changes.UpdateNew = []*endpoint.Endpoint{{DNSName: "test3.example.com", Targets: endpoint.Targets{"192.168.0.2"}, RecordType: "A", RecordTTL: 777}}
	changes.Delete = []*endpoint.Endpoint{{DNSName: "test4.example.com", Targets: endpoint.Targets{"192.168.0.3"}, RecordType: "A"}}

	// Failing ListDomains API call creates an error when calling Records
	mockedClient := &mockGandiClient{
		FunctionToFail: "ListDomains",
	}
	mockedProvider := &GandiProvider{
		DomainClient:  mockedClient,
		LiveDNSClient: mockedClient,
	}

	_, err := mockedProvider.Records(context.Background())
	if err == nil {
		t.Error("should have failed")
	}

	// Failing GetDomainRecords API call creates an error when calling Records
	mockedClient = &mockGandiClient{
		FunctionToFail: "GetDomainRecords",
	}
	mockedProvider = &GandiProvider{
		DomainClient:  mockedClient,
		LiveDNSClient: mockedClient,
	}

	_, err = mockedProvider.Records(context.Background())
	if err == nil {
		t.Error("should have failed")
	}

	// Failing ListDomains API call creates an error when calling ApplyChanges
	mockedClient = &mockGandiClient{
		FunctionToFail: "ListDomains",
	}
	mockedProvider = &GandiProvider{
		DomainClient:  mockedClient,
		LiveDNSClient: mockedClient,
	}

	err = mockedProvider.ApplyChanges(context.Background(), changes)
	if err == nil {
		t.Error("should have failed")
	}

	// Failing CreateDomainRecord API call creates an error when calling ApplyChanges
	mockedClient = &mockGandiClient{
		FunctionToFail: "CreateDomainRecord",
	}
	mockedProvider = &GandiProvider{
		DomainClient:  mockedClient,
		LiveDNSClient: mockedClient,
	}

	err = mockedProvider.ApplyChanges(context.Background(), changes)
	if err == nil {
		t.Error("should have failed")
	}

	// Failing DeleteDomainRecord API call creates an error when calling ApplyChanges
	mockedClient = &mockGandiClient{
		FunctionToFail: "DeleteDomainRecord",
	}
	mockedProvider = &GandiProvider{
		DomainClient:  mockedClient,
		LiveDNSClient: mockedClient,
	}

	err = mockedProvider.ApplyChanges(context.Background(), changes)
	if err == nil {
		t.Error("should have failed")
	}

	// Failing UpdateDomainRecordByNameAndType API call creates an error when calling ApplyChanges
	mockedClient = &mockGandiClient{
		FunctionToFail: "UpdateDomainRecordByNameAndType",
	}
	mockedProvider = &GandiProvider{
		DomainClient:  mockedClient,
		LiveDNSClient: mockedClient,
	}

	err = mockedProvider.ApplyChanges(context.Background(), changes)
	if err == nil {
		t.Error("should have failed")
	}
}
