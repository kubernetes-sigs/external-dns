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

package domeneshop

import (
	"context"
	"os"
	"sigs.k8s.io/external-dns/plan"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
)

type MockDomeneshopClient struct {
	mock.Mock
}

func (m *MockDomeneshopClient) ListDomains(ctx context.Context) ([]*Domain, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*Domain), args.Error(1)
}

func (m *MockDomeneshopClient) ListDNSRecords(ctx context.Context, domain *Domain, host, recordType string) ([]*DNSRecord, error) {
	args := m.Called(ctx, domain, host, recordType)
	return args.Get(0).([]*DNSRecord), args.Error(1)
}
func (m *MockDomeneshopClient) AddDNSRecord(ctx context.Context, domain *Domain, record *DNSRecord) error {
	args := m.Called(ctx, domain, record)
	return args.Error(0)
}

func (m *MockDomeneshopClient) UpdateDNSRecord(ctx context.Context, domain *Domain, record *DNSRecord) error {
	args := m.Called(ctx, domain, record)
	return args.Error(0)
}
func (m *MockDomeneshopClient) DeleteDNSRecord(ctx context.Context, domain *Domain, record *DNSRecord) error {
	args := m.Called(ctx, domain, record)
	return args.Error(0)
}

func createZones() []*Domain {
	return []*Domain{
		{Name: "foo.com"},
		{Name: "bar.io"},
		{Name: "baz.com"},
	}
}

func createFooRecords() []*DNSRecord {
	return []*DNSRecord{{
		ID:   11,
		Type: "A",
		Host: "@",
		Data: "targetFoo",
	}, {
		ID:   12,
		Type: "TXT",
		Host: "@",
		Data: "txt",
	}, {
		ID:   13,
		Type: "CAA",
		Host: "@",
		Data: "",
	}}
}

func createBarRecords() []*DNSRecord {
	return []*DNSRecord{}
}

func createBazRecords() []*DNSRecord {
	return []*DNSRecord{{
		ID:   31,
		Type: "A",
		Host: "@",
		Data: "targetBaz",
	}, {
		ID:   32,
		Type: "TXT",
		Host: "@",
		Data: "txt",
	}, {
		ID:   33,
		Type: "A",
		Host: "api",
		Data: "targetBaz",
	}, {
		ID:   34,
		Type: "TXT",
		Host: "api",
		Data: "txt",
	}}
}

func filteredRecords(records []*DNSRecord, host, recordType string) []*DNSRecord {
	var rs []*DNSRecord

	for _, record := range records {
		if host != "" && record.Host != host {
			continue
		}

		if recordType != "" && record.Type != recordType {
			continue
		}

		rs = append(rs, record)
	}

	return rs
}

var domainFilter endpoint.DomainFilter

func TestNewDomeneshopProvider(t *testing.T) {
	_ = os.Setenv("DOMENESHOP_API_TOKEN", "xxxxxxxxxxxxxxxxx")
	_ = os.Setenv("DOMENESHOP_API_SECRET", "xxxxxxxxxxxxxxxxx")
	_, err := NewDomeneshopProvider(context.Background(), endpoint.NewDomainFilter([]string{"ext-dns-test.foo.com."}), true, "1.0")
	require.NoError(t, err)

	_ = os.Unsetenv("DOMENESHOP_API_TOKEN")
	_ = os.Unsetenv("DOMENESHOP_API_SECRET")
	_, err = NewDomeneshopProvider(context.Background(), endpoint.NewDomainFilter([]string{"ext-dns-test.foo.com."}), true, "1.0")
	require.Error(t, err)
}

func TestDomeneshopProvider_Records(t *testing.T) {
	mockDomainClient := MockDomeneshopClient{}

	provider := &DomeneshopProvider{
		Client:       &mockDomainClient,
		domainFilter: endpoint.NewDomainFilter([]string{}),
		DryRun:       false,
	}

	mockDomainClient.On(
		"ListDomains",
		mock.Anything,
	).Return(createZones(), nil).Once()

	mockDomainClient.On(
		"ListDNSRecords",
		mock.Anything,
		mock.MatchedBy(func(domain *Domain) bool { return domain.Name == "foo.com" }),
		"",
		"",
	).Return(createFooRecords(), nil).Once()

	mockDomainClient.On(
		"ListDNSRecords",
		mock.Anything,
		mock.MatchedBy(func(domain *Domain) bool { return domain.Name == "bar.io" }),
		"",
		"",
	).Return(createBarRecords(), nil).Once()

	mockDomainClient.On(
		"ListDNSRecords",
		mock.Anything,
		mock.MatchedBy(func(domain *Domain) bool { return domain.Name == "baz.com" }),
		"",
		"",
	).Return(createBazRecords(), nil).Once()

	actual, err := provider.Records(context.Background())
	require.NoError(t, err)

	expected := []*endpoint.Endpoint{
		{DNSName: "foo.com", Targets: []string{"targetFoo"}, RecordType: "A", RecordTTL: 0, Labels: endpoint.NewLabels()},
		{DNSName: "foo.com", Targets: []string{"txt"}, RecordType: "TXT", RecordTTL: 0, Labels: endpoint.NewLabels()},
		{DNSName: "foo.com", Targets: []string{""}, RecordType: "CAA", RecordTTL: 0, Labels: endpoint.NewLabels()},
		{DNSName: "baz.com", Targets: []string{"targetBaz"}, RecordType: "A", RecordTTL: 0, Labels: endpoint.NewLabels()},
		{DNSName: "baz.com", Targets: []string{"txt"}, RecordType: "TXT", RecordTTL: 0, Labels: endpoint.NewLabels()},
		{DNSName: "api.baz.com", Targets: []string{"targetBaz"}, RecordType: "A", RecordTTL: 0, Labels: endpoint.NewLabels()},
		{DNSName: "api.baz.com", Targets: []string{"txt"}, RecordType: "TXT", RecordTTL: 0, Labels: endpoint.NewLabels()},
	}

	mockDomainClient.AssertExpectations(t)
	assert.Equal(t, expected, actual)
}

func TestDomeneshopProvider_ApplyChanges(t *testing.T) {
	mockDomainClient := MockDomeneshopClient{}

	provider := &DomeneshopProvider{
		Client:       &mockDomainClient,
		domainFilter: endpoint.NewDomainFilter([]string{}),
		DryRun:       false,
	}

	mockDomainClient.On(
		"ListDomains",
		mock.Anything,
	).Return(createZones(), nil).Once()

	// Apply Actions
	mockDomainClient.On(
		"ListDNSRecords",
		mock.Anything,
		mock.MatchedBy(func(domain *Domain) bool { return domain.Name == "baz.com" }),
		"api",
		"A",
	).Return(filteredRecords(createBazRecords(), "api", "A"), nil).Once()

	mockDomainClient.On(
		"DeleteDNSRecord",
		mock.Anything,
		mock.MatchedBy(func(domain *Domain) bool { return domain.Name == "baz.com" }),
		mock.MatchedBy(func(record *DNSRecord) bool { return record.ID == 33 }),
	).Return(nil).Once()

	mockDomainClient.On(
		"ListDNSRecords",
		mock.Anything,
		mock.MatchedBy(func(domain *Domain) bool { return domain.Name == "baz.com" }),
		"api",
		"TXT",
	).Return(filteredRecords(createBazRecords(), "api", "TXT"), nil).Once()

	mockDomainClient.On(
		"DeleteDNSRecord",
		mock.Anything,
		mock.MatchedBy(func(domain *Domain) bool { return domain.Name == "baz.com" }),
		mock.MatchedBy(func(record *DNSRecord) bool { return record.ID == 34 }),
	).Return(nil).Once()

	mockDomainClient.On(
		"ListDNSRecords",
		mock.Anything,
		mock.MatchedBy(func(domain *Domain) bool { return domain.Name == "foo.com" }),
		"@",
		"A",
	).Return(filteredRecords(createFooRecords(), "@", "A"), nil).Once()

	mockDomainClient.On(
		"UpdateDNSRecord",
		mock.Anything,
		mock.MatchedBy(func(domain *Domain) bool { return domain.Name == "foo.com" }),
		mock.MatchedBy(func(record *DNSRecord) bool {
			return record.ID == 11 && record.Host == "@" && record.Type == "A" && record.Data == "targetFoo" && record.TTL == 300
		}),
	).Return(nil).Once()

	mockDomainClient.On(
		"AddDNSRecord",
		mock.Anything,
		mock.MatchedBy(func(domain *Domain) bool { return domain.Name == "bar.io" }),
		mock.MatchedBy(func(record *DNSRecord) bool {
			return record.Host == "create" && record.Type == "A" && record.Data == "targetBar" && record.TTL == 300
		}),
	).Return(nil).Once()

	mockDomainClient.On(
		"AddDNSRecord",
		mock.Anything,
		mock.MatchedBy(func(domain *Domain) bool { return domain.Name == "bar.io" }),
		mock.MatchedBy(func(record *DNSRecord) bool {
			return record.Host == "@" && record.Type == "A" && record.Data == "targetBar" && record.TTL == 300
		}),
	).Return(nil).Once()

	err := provider.ApplyChanges(context.Background(), &plan.Changes{
		Create: []*endpoint.Endpoint{{
			DNSName:    "create.bar.io",
			RecordType: "A",
			Targets:    []string{"targetBar"},
		}, {
			DNSName:    "bar.io",
			RecordType: "A",
			Targets:    []string{"targetBar"},
		}},
		Delete: []*endpoint.Endpoint{{
			DNSName:    "api.baz.com",
			RecordType: "A",
		}, {
			DNSName:    "api.baz.com",
			RecordType: "TXT",
		}},
		UpdateNew: []*endpoint.Endpoint{{
			DNSName:    "foo.com",
			RecordType: "A",
			RecordTTL:  300,
			Targets:    []string{"targetFoo"},
		}},
		UpdateOld: []*endpoint.Endpoint{},
	})
	require.NoError(t, err)

	mockDomainClient.AssertExpectations(t)
}

func TestDomeneshopProvider_ApplyChanges_TargetAdded(t *testing.T) {
	mockDomainClient := MockDomeneshopClient{}

	provider := &DomeneshopProvider{
		Client:       &mockDomainClient,
		domainFilter: endpoint.NewDomainFilter([]string{}),
		DryRun:       false,
	}

	// Dummy Data
	mockDomainClient.On(
		"ListDomains",
		mock.Anything,
	).Return([]*Domain{{Name: "example.com"}}, nil).Once()

	// Apply Actions
	mockDomainClient.On(
		"ListDNSRecords",
		mock.Anything,
		mock.MatchedBy(func(domain *Domain) bool { return domain.Name == "example.com" }),
		"@",
		"A",
	).Return([]*DNSRecord{{ID: 11, Host: "@", Type: "A", Data: "targetA"}}, nil).Once()

	mockDomainClient.On(
		"UpdateDNSRecord",
		mock.Anything,
		mock.MatchedBy(func(domain *Domain) bool { return domain.Name == "example.com" }),
		mock.MatchedBy(func(record *DNSRecord) bool {
			return record.ID == 11 && record.Host == "@" && record.Type == "A" && record.Data == "targetA" && record.TTL == 300
		}),
	).Return(nil).Once()

	mockDomainClient.On(
		"AddDNSRecord",
		mock.Anything,
		mock.MatchedBy(func(domain *Domain) bool { return domain.Name == "example.com" }),
		mock.MatchedBy(func(record *DNSRecord) bool {
			return record.Host == "@" && record.Type == "A" && record.Data == "targetB" && record.TTL == 300
		}),
	).Return(nil).Once()

	err := provider.ApplyChanges(context.Background(), &plan.Changes{
		// From 1 target to 2
		UpdateNew: []*endpoint.Endpoint{{
			DNSName:    "example.com",
			RecordType: "A",
			Targets:    []string{"targetA", "targetB"},
		}},
		UpdateOld: []*endpoint.Endpoint{},
	})
	require.NoError(t, err)

	mockDomainClient.AssertExpectations(t)
}

func TestDomeneshopProvider_ApplyChanges_TargetRemoved(t *testing.T) {
	mockDomainClient := MockDomeneshopClient{}

	provider := &DomeneshopProvider{
		Client:       &mockDomainClient,
		domainFilter: endpoint.NewDomainFilter([]string{}),
		DryRun:       false,
	}

	// Dummy Data
	mockDomainClient.On(
		"ListDomains",
		mock.Anything,
	).Return([]*Domain{{Name: "example.com"}}, nil).Once()

	mockDomainClient.On(
		"ListDNSRecords",
		mock.Anything,
		mock.MatchedBy(func(domain *Domain) bool { return domain.Name == "example.com" }),
		"@",
		"A",
	).Return([]*DNSRecord{{ID: 11, Host: "@", Type: "A", Data: "targetA"}, {ID: 12, Host: "@", Type: "A", Data: "targetB"}}, nil).Once()

	// Apply Actions
	mockDomainClient.On(
		"UpdateDNSRecord",
		mock.Anything,
		mock.MatchedBy(func(domain *Domain) bool { return domain.Name == "example.com" }),
		mock.MatchedBy(func(record *DNSRecord) bool {
			return record.ID == 12 && record.Host == "@" && record.Type == "A" && record.Data == "targetB" && record.TTL == 300
		}),
	).Return(nil).Once()

	mockDomainClient.On(
		"DeleteDNSRecord",
		mock.Anything,
		mock.MatchedBy(func(domain *Domain) bool { return domain.Name == "example.com" }),
		mock.MatchedBy(func(record *DNSRecord) bool {
			return record.ID == 11 && record.Host == "@" && record.Type == "A" && record.Data == "targetA"
		}),
	).Return(nil).Once()

	err := provider.ApplyChanges(context.Background(), &plan.Changes{
		// From 2 targets to 1
		UpdateNew: []*endpoint.Endpoint{{
			DNSName:    "example.com",
			RecordType: "A",
			Targets:    []string{"targetB"},
		}},
		UpdateOld: []*endpoint.Endpoint{},
	})
	require.NoError(t, err)

	mockDomainClient.AssertExpectations(t)
}

func TestDomeneshopProvider_ApplyChanges_NoChanges(t *testing.T) {
	mockDomainClient := MockDomeneshopClient{}

	provider := &DomeneshopProvider{
		Client:       &mockDomainClient,
		domainFilter: endpoint.NewDomainFilter([]string{}),
		DryRun:       false,
	}

	mockDomainClient.On(
		"ListDomains",
		mock.Anything,
	).Return(createZones(), nil).Once()

	err := provider.ApplyChanges(context.Background(), &plan.Changes{})
	require.NoError(t, err)

	mockDomainClient.AssertExpectations(t)
}

func TestDomeneshopMergeEndpointsByNameType(t *testing.T) {
	xs := []*endpoint.Endpoint{
		endpoint.NewEndpoint("foo.example.com", "A", "1.2.3.4"),
		endpoint.NewEndpoint("bar.example.com", "A", "1.2.3.4"),
		endpoint.NewEndpoint("foo.example.com", "A", "5.6.7.8"),
		endpoint.NewEndpoint("foo.example.com", "CNAME", "somewhere.out.there.com"),
	}

	merged := mergeEndpointsByNameType(xs)

	assert.Equal(t, 3, len(merged))
	sort.SliceStable(merged, func(i, j int) bool {
		if merged[i].DNSName != merged[j].DNSName {
			return merged[i].DNSName < merged[j].DNSName
		}
		return merged[i].RecordType < merged[j].RecordType
	})
	assert.Equal(t, "bar.example.com", merged[0].DNSName)
	assert.Equal(t, "A", merged[0].RecordType)
	assert.Equal(t, 1, len(merged[0].Targets))
	assert.Equal(t, "1.2.3.4", merged[0].Targets[0])

	assert.Equal(t, "foo.example.com", merged[1].DNSName)
	assert.Equal(t, "A", merged[1].RecordType)
	assert.Equal(t, 2, len(merged[1].Targets))
	assert.ElementsMatch(t, []string{"1.2.3.4", "5.6.7.8"}, merged[1].Targets)

	assert.Equal(t, "foo.example.com", merged[2].DNSName)
	assert.Equal(t, "CNAME", merged[2].RecordType)
	assert.Equal(t, 1, len(merged[2].Targets))
	assert.Equal(t, "somewhere.out.there.com", merged[2].Targets[0])
}
