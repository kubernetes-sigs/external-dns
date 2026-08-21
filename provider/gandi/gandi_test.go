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
	"fmt"
	"os"
	"testing"

	"github.com/go-gandi/go-gandi/domain"
	"github.com/go-gandi/go-gandi/livedns"
	"github.com/maxatome/go-testdeep/td"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
	"sigs.k8s.io/external-dns/plan"
)

type MockAction struct {
	Name   string
	FQDN   string
	Record livedns.DomainRecord
}

// mockGandiClient implements DomainClientAdapter and LiveDNSClientAdapter for provider tests.
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

func (m *mockGandiClient) GetDomainRecords(fqdn string) ([]livedns.DomainRecord, error) {
	m.Actions = append(m.Actions, MockAction{
		Name: "GetDomainRecords",
		FQDN: fqdn,
	})

	if m.FunctionToFail == "GetDomainRecords" {
		return nil, fmt.Errorf("injected error")
	}

	return m.RecordsToReturn, nil
}

func (m *mockGandiClient) CreateDomainRecord(fqdn, name, recordtype string, ttl int, values []string) (standardResponse, error) {
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

func (m *mockGandiClient) DeleteDomainRecord(fqdn, name, recordtype string) error {
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

func (m *mockGandiClient) UpdateDomainRecordByNameAndType(fqdn, name, recordtype string, ttl int, values []string) (standardResponse, error) {
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

func (m *mockGandiClient) ListDomains() ([]domain.ListResponse, error) {
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

func TestNew(t *testing.T) {
	domainFilter := endpoint.NewDomainFilter([]string{"example.com"})

	tests := []struct {
		name       string
		env        map[string]string
		dryRun     bool
		wantError  bool
		wantDryRun bool
	}{
		{
			name:       "succeeds with API key",
			env:        map[string]string{"GANDI_KEY": "myGandiKey"},
			dryRun:     true,
			wantDryRun: true,
		},
		{
			name:       "succeeds with PAT",
			env:        map[string]string{"GANDI_PAT": "myGandiPAT"},
			dryRun:     true,
			wantDryRun: true,
		},
		{
			name:       "succeeds with PAT and sharing ID",
			env:        map[string]string{"GANDI_PAT": "myGandiPAT", "GANDI_SHARING_ID": "aSharingId"},
			dryRun:     false,
			wantDryRun: false,
		},
		{
			name:      "errors without credentials",
			env:       map[string]string{},
			dryRun:    true,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("GANDI_KEY", "")
			t.Setenv("GANDI_PAT", "")
			t.Setenv("GANDI_SHARING_ID", "")
			_ = os.Unsetenv("GANDI_KEY")
			_ = os.Unsetenv("GANDI_PAT")
			_ = os.Unsetenv("GANDI_SHARING_ID")
			for key, value := range tt.env {
				t.Setenv(key, value)
			}

			prov, err := New(t.Context(), &externaldns.Config{DryRun: tt.dryRun}, domainFilter)
			if tt.wantError {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			gp, ok := prov.(*GandiProvider)
			require.True(t, ok)
			assert.Equal(t, tt.wantDryRun, gp.DryRun)
		})
	}
}

func TestGandiProvider_Records(t *testing.T) {
	tests := []struct {
		name         string
		records      []livedns.DomainRecord
		domainFilter *endpoint.DomainFilter
		want         []*endpoint.Endpoint
	}{
		{
			name: "returns correct endpoints",
			records: []livedns.DomainRecord{
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
			want: []*endpoint.Endpoint{
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
			},
		},
		{
			name: "filtered domains yield no endpoints",
			records: []livedns.DomainRecord{
				{
					RrsetType:   endpoint.RecordTypeCNAME,
					RrsetTTL:    600,
					RrsetName:   "@",
					RrsetHref:   exampleDotComUri + "/records/test/MX",
					RrsetValues: []string{"192.168.0.1"},
				},
			},
			domainFilter: endpoint.NewDomainFilterWithExclusions([]string{}, []string{"example.com"}),
			want:         []*endpoint.Endpoint{},
		},
		{
			name: "unsupported record types are skipped",
			records: []livedns.DomainRecord{
				{
					RrsetType:   "MX",
					RrsetTTL:    360,
					RrsetName:   "@",
					RrsetHref:   exampleDotComUri + "/records/%40/A",
					RrsetValues: []string{"smtp.example.com"},
				},
			},
			want: []*endpoint.Endpoint{},
		},
		{
			name: "multiple values per record",
			records: []livedns.DomainRecord{{
				RrsetType:   endpoint.RecordTypeA,
				RrsetTTL:    300,
				RrsetName:   "multi",
				RrsetValues: []string{"203.0.113.1", "203.0.113.2"},
			}},
			want: []*endpoint.Endpoint{
				{RecordType: endpoint.RecordTypeA, DNSName: "multi.example.com", Targets: endpoint.Targets{"203.0.113.1"}, RecordTTL: 300},
				{RecordType: endpoint.RecordTypeA, DNSName: "multi.example.com", Targets: endpoint.Targets{"203.0.113.2"}, RecordTTL: 300},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockGandiClient{RecordsToReturn: tt.records}
			provider := &GandiProvider{
				DomainClient:  mock,
				LiveDNSClient: mock,
				domainFilter:  tt.domainFilter,
			}

			endpoints, err := provider.Records(t.Context())
			require.NoError(t, err)

			assert.True(t, testutils.SameEndpoints(tt.want, endpoints), "expected %v, got %v", tt.want, endpoints)
		})
	}
}

func TestGandiProvider_ApplyChanges(t *testing.T) {
	tests := []struct {
		name         string
		dryRun       bool
		changes      *plan.Changes
		wantActions  []MockAction
		wantNoAction bool
	}{
		{
			name: "makes expected API calls",
			changes: &plan.Changes{
				Create: []*endpoint.Endpoint{{
					DNSName:    "test2.example.com",
					Targets:    endpoint.Targets{"192.168.0.1"},
					RecordType: "A",
					RecordTTL:  666,
				}},
				UpdateNew: []*endpoint.Endpoint{
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
				},
				Delete: []*endpoint.Endpoint{{
					DNSName:    "test4.example.com",
					Targets:    endpoint.Targets{"192.168.0.3"},
					RecordType: "A",
				}},
			},
			wantActions: []MockAction{
				{Name: "ListDomains"},
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
			},
		},
		{
			name:   "respects dry run",
			dryRun: true,
			changes: &plan.Changes{
				Create:    []*endpoint.Endpoint{{DNSName: "test2.example.com", Targets: endpoint.Targets{"192.168.0.1"}, RecordType: "A", RecordTTL: 666}},
				UpdateNew: []*endpoint.Endpoint{{DNSName: "test3.example.com", Targets: endpoint.Targets{"192.168.0.2"}, RecordType: "A", RecordTTL: 777}},
				Delete:    []*endpoint.Endpoint{{DNSName: "test4.example.com", Targets: endpoint.Targets{"192.168.0.3"}, RecordType: "A"}},
			},
			wantActions: []MockAction{{Name: "ListDomains"}},
		},
		{
			name:         "empty changes do nothing",
			changes:      &plan.Changes{},
			wantNoAction: true,
		},
		{
			name: "unknown domain does no update",
			changes: &plan.Changes{
				Create: []*endpoint.Endpoint{{
					DNSName:    "test.example.net",
					Targets:    endpoint.Targets{"192.168.0.1"},
					RecordType: "A",
					RecordTTL:  666,
				}},
			},
			wantActions: []MockAction{{Name: "ListDomains"}},
		},
		{
			name: "converts apex domain",
			changes: &plan.Changes{
				Create: []*endpoint.Endpoint{{
					DNSName:    "example.com",
					Targets:    endpoint.Targets{"192.168.0.1"},
					RecordType: "A",
					RecordTTL:  666,
				}},
			},
			wantActions: []MockAction{
				{Name: "ListDomains"},
				{
					Name: "CreateDomainRecord",
					FQDN: "example.com",
					Record: livedns.DomainRecord{
						RrsetType:   endpoint.RecordTypeA,
						RrsetName:   "@",
						RrsetValues: []string{"192.168.0.1"},
						RrsetTTL:    666,
					},
				},
			},
		},
		{
			name: "uses default TTL",
			changes: &plan.Changes{
				Create: []*endpoint.Endpoint{{
					DNSName:    "test.example.com",
					Targets:    endpoint.Targets{"192.168.0.1"},
					RecordType: endpoint.RecordTypeA,
				}},
			},
			wantActions: []MockAction{
				{Name: "ListDomains"},
				{
					Name: "CreateDomainRecord",
					FQDN: "example.com",
					Record: livedns.DomainRecord{
						RrsetType:   endpoint.RecordTypeA,
						RrsetName:   "test",
						RrsetValues: []string{"192.168.0.1"},
						RrsetTTL:    defaultTTL,
					},
				},
			},
		},
		{
			name: "CNAME keeps trailing dot",
			changes: &plan.Changes{
				UpdateNew: []*endpoint.Endpoint{{
					DNSName:    "www.example.com",
					Targets:    endpoint.Targets{"lb.example.net."},
					RecordType: endpoint.RecordTypeCNAME,
					RecordTTL:  300,
				}},
			},
			wantActions: []MockAction{
				{Name: "ListDomains"},
				{
					Name: "UpdateDomainRecordByNameAndType",
					FQDN: "example.com",
					Record: livedns.DomainRecord{
						RrsetType:   endpoint.RecordTypeCNAME,
						RrsetName:   "www",
						RrsetValues: []string{"lb.example.net."},
						RrsetTTL:    300,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockGandiClient{}
			provider := &GandiProvider{
				DryRun:        tt.dryRun,
				DomainClient:  mock,
				LiveDNSClient: mock,
			}

			err := provider.ApplyChanges(t.Context(), tt.changes)
			require.NoError(t, err)
			if tt.wantNoAction {
				assert.Empty(t, mock.Actions)
				return
			}
			td.Cmp(t, mock.Actions, tt.wantActions)
		})
	}
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

	_, err := mockedProvider.Records(t.Context())
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

	_, err = mockedProvider.Records(t.Context())
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

	err = mockedProvider.ApplyChanges(t.Context(), changes)
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

	err = mockedProvider.ApplyChanges(t.Context(), changes)
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

	err = mockedProvider.ApplyChanges(t.Context(), changes)
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

	err = mockedProvider.ApplyChanges(t.Context(), changes)
	if err == nil {
		t.Error("should have failed")
	}
}

func TestGandiProvider_Zones(t *testing.T) {
	mock := &mockGandiClient{}
	provider := &GandiProvider{
		DomainClient: mock,
		domainFilter: endpoint.NewDomainFilter([]string{"example.com"}),
	}

	zones, err := provider.Zones()
	require.NoError(t, err)
	assert.Equal(t, []string{"example.com"}, zones)

	provider.domainFilter = endpoint.NewDomainFilter([]string{"other.org"})
	zones, err = provider.Zones()
	require.NoError(t, err)
	assert.Empty(t, zones)

	provider.DomainClient = &mockGandiClient{FunctionToFail: "ListDomains"}
	_, err = provider.Zones()
	assert.Error(t, err)
}
