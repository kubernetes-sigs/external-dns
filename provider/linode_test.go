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

package provider

import (
	"context"
	"os"
	"testing"

	"github.com/linode/linodego"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

type MockDomainClient struct {
	mock.Mock
}

func (m *MockDomainClient) ListDomainRecords(ctx context.Context, domainID int, opts *linodego.ListOptions) ([]*linodego.DomainRecord, error) {
	args := m.Called(ctx, domainID, opts)
	return args.Get(0).([]*linodego.DomainRecord), args.Error(1)
}

func (m *MockDomainClient) ListDomains(ctx context.Context, opts *linodego.ListOptions) ([]*linodego.Domain, error) {
	args := m.Called(ctx, opts)
	return args.Get(0).([]*linodego.Domain), args.Error(1)
}
func (m *MockDomainClient) CreateDomainRecord(ctx context.Context, domainID int, opts linodego.DomainRecordCreateOptions) (*linodego.DomainRecord, error) {
	args := m.Called(ctx, domainID, opts)
	return args.Get(0).(*linodego.DomainRecord), args.Error(1)
}
func (m *MockDomainClient) DeleteDomainRecord(ctx context.Context, domainID int, recordID int) error {
	args := m.Called(ctx, domainID, recordID)
	return args.Error(0)
}
func (m *MockDomainClient) UpdateDomainRecord(ctx context.Context, domainID int, recordID int, opts linodego.DomainRecordUpdateOptions) (*linodego.DomainRecord, error) {
	args := m.Called(ctx, domainID, recordID, opts)
	return args.Get(0).(*linodego.DomainRecord), args.Error(1)
}

func createZones() []*linodego.Domain {
	return []*linodego.Domain{
		{ID: 1, Domain: "foo.com"},
		{ID: 2, Domain: "bar.io"},
		{ID: 3, Domain: "baz.com"},
	}
}

func createFooRecords() []*linodego.DomainRecord {
	return []*linodego.DomainRecord{{
		ID:     11,
		Type:   linodego.RecordTypeA,
		Name:   "",
		Target: "targetFoo",
	}, {
		ID:     12,
		Type:   linodego.RecordTypeTXT,
		Name:   "",
		Target: "txt",
	}, {
		ID:     13,
		Type:   linodego.RecordTypeCAA,
		Name:   "foo.com",
		Target: "",
	}}
}

func createBarRecords() []*linodego.DomainRecord {
	return []*linodego.DomainRecord{}
}

func createBazRecords() []*linodego.DomainRecord {
	return []*linodego.DomainRecord{{
		ID:     31,
		Type:   linodego.RecordTypeA,
		Name:   "",
		Target: "targetBaz",
	}, {
		ID:     32,
		Type:   linodego.RecordTypeTXT,
		Name:   "",
		Target: "txt",
	}, {
		ID:     33,
		Type:   linodego.RecordTypeA,
		Name:   "api",
		Target: "targetBaz",
	}, {
		ID:     34,
		Type:   linodego.RecordTypeTXT,
		Name:   "api",
		Target: "txt",
	}}
}

func TestLinodeConvertRecordType(t *testing.T) {
	record, err := convertRecordType("A")
	require.NoError(t, err)
	assert.Equal(t, linodego.RecordTypeA, record)

	record, err = convertRecordType("AAAA")
	require.NoError(t, err)
	assert.Equal(t, linodego.RecordTypeAAAA, record)

	record, err = convertRecordType("CNAME")
	require.NoError(t, err)
	assert.Equal(t, linodego.RecordTypeCNAME, record)

	record, err = convertRecordType("TXT")
	require.NoError(t, err)
	assert.Equal(t, linodego.RecordTypeTXT, record)

	record, err = convertRecordType("SRV")
	require.NoError(t, err)
	assert.Equal(t, linodego.RecordTypeSRV, record)

	_, err = convertRecordType("INVALID")
	require.Error(t, err)
}

func TestNewLinodeProvider(t *testing.T) {
	_ = os.Setenv("LINODE_TOKEN", "xxxxxxxxxxxxxxxxx")
	_, err := NewLinodeProvider(endpoint.NewDomainFilter([]string{"ext-dns-test.zalando.to."}), true, "1.0")
	require.NoError(t, err)

	_ = os.Unsetenv("LINODE_TOKEN")
	_, err = NewLinodeProvider(endpoint.NewDomainFilter([]string{"ext-dns-test.zalando.to."}), true, "1.0")
	require.Error(t, err)
}

func TestLinodeStripRecordName(t *testing.T) {
	assert.Equal(t, "api", getStrippedRecordName(&linodego.Domain{
		Domain: "example.com",
	}, &endpoint.Endpoint{
		DNSName: "api.example.com",
	}))

	assert.Equal(t, "", getStrippedRecordName(&linodego.Domain{
		Domain: "example.com",
	}, &endpoint.Endpoint{
		DNSName: "example.com",
	}))
}

func TestLinodeFetchZonesNoFilters(t *testing.T) {
	mockDomainClient := MockDomainClient{}

	provider := &LinodeProvider{
		Client:       &mockDomainClient,
		domainFilter: endpoint.NewDomainFilter([]string{}),
		DryRun:       false,
	}

	mockDomainClient.On(
		"ListDomains",
		mock.Anything,
		mock.Anything,
	).Return(createZones(), nil).Once()

	expected := createZones()
	actual, err := provider.fetchZones(context.Background())
	require.NoError(t, err)

	mockDomainClient.AssertExpectations(t)
	assert.Equal(t, expected, actual)
}

func TestLinodeFetchZonesWithFilter(t *testing.T) {
	mockDomainClient := MockDomainClient{}

	provider := &LinodeProvider{
		Client:       &mockDomainClient,
		domainFilter: endpoint.NewDomainFilter([]string{".com"}),
		DryRun:       false,
	}

	mockDomainClient.On(
		"ListDomains",
		mock.Anything,
		mock.Anything,
	).Return(createZones(), nil).Once()

	expected := []*linodego.Domain{
		{ID: 1, Domain: "foo.com"},
		{ID: 3, Domain: "baz.com"},
	}
	actual, err := provider.fetchZones(context.Background())
	require.NoError(t, err)

	mockDomainClient.AssertExpectations(t)
	assert.Equal(t, expected, actual)
}

func TestLinodeGetStrippedRecordName(t *testing.T) {
	assert.Equal(t, "", getStrippedRecordName(&linodego.Domain{
		Domain: "foo.com",
	}, &endpoint.Endpoint{
		DNSName: "foo.com",
	}))

	assert.Equal(t, "api", getStrippedRecordName(&linodego.Domain{
		Domain: "foo.com",
	}, &endpoint.Endpoint{
		DNSName: "api.foo.com",
	}))
}

func TestLinodeRecords(t *testing.T) {
	mockDomainClient := MockDomainClient{}

	provider := &LinodeProvider{
		Client:       &mockDomainClient,
		domainFilter: endpoint.NewDomainFilter([]string{}),
		DryRun:       false,
	}

	mockDomainClient.On(
		"ListDomains",
		mock.Anything,
		mock.Anything,
	).Return(createZones(), nil).Once()

	mockDomainClient.On(
		"ListDomainRecords",
		mock.Anything,
		1,
		mock.Anything,
	).Return(createFooRecords(), nil).Once()
	mockDomainClient.On(
		"ListDomainRecords",
		mock.Anything,
		2,
		mock.Anything,
	).Return(createBarRecords(), nil).Once()
	mockDomainClient.On(
		"ListDomainRecords",
		mock.Anything,
		3,
		mock.Anything,
	).Return(createBazRecords(), nil).Once()

	actual, err := provider.Records(context.Background())
	require.NoError(t, err)

	expected := []*endpoint.Endpoint{
		{DNSName: "foo.com", Targets: []string{"targetFoo"}, RecordType: "A", RecordTTL: 0, Labels: endpoint.NewLabels()},
		{DNSName: "foo.com", Targets: []string{"txt"}, RecordType: "TXT", RecordTTL: 0, Labels: endpoint.NewLabels()},
		{DNSName: "baz.com", Targets: []string{"targetBaz"}, RecordType: "A", RecordTTL: 0, Labels: endpoint.NewLabels()},
		{DNSName: "baz.com", Targets: []string{"txt"}, RecordType: "TXT", RecordTTL: 0, Labels: endpoint.NewLabels()},
		{DNSName: "api.baz.com", Targets: []string{"targetBaz"}, RecordType: "A", RecordTTL: 0, Labels: endpoint.NewLabels()},
		{DNSName: "api.baz.com", Targets: []string{"txt"}, RecordType: "TXT", RecordTTL: 0, Labels: endpoint.NewLabels()},
	}

	mockDomainClient.AssertExpectations(t)
	assert.Equal(t, expected, actual)
}

func TestLinodeApplyChanges(t *testing.T) {
	mockDomainClient := MockDomainClient{}

	provider := &LinodeProvider{
		Client:       &mockDomainClient,
		domainFilter: endpoint.NewDomainFilter([]string{}),
		DryRun:       false,
	}

	// Dummy Data
	mockDomainClient.On(
		"ListDomains",
		mock.Anything,
		mock.Anything,
	).Return(createZones(), nil).Once()

	mockDomainClient.On(
		"ListDomainRecords",
		mock.Anything,
		1,
		mock.Anything,
	).Return(createFooRecords(), nil).Once()
	mockDomainClient.On(
		"ListDomainRecords",
		mock.Anything,
		2,
		mock.Anything,
	).Return(createBarRecords(), nil).Once()
	mockDomainClient.On(
		"ListDomainRecords",
		mock.Anything,
		3,
		mock.Anything,
	).Return(createBazRecords(), nil).Once()

	// Apply Actions
	mockDomainClient.On(
		"DeleteDomainRecord",
		mock.Anything,
		3,
		33,
	).Return(nil).Once()

	mockDomainClient.On(
		"DeleteDomainRecord",
		mock.Anything,
		3,
		34,
	).Return(nil).Once()

	mockDomainClient.On(
		"UpdateDomainRecord",
		mock.Anything,
		1,
		11,
		linodego.DomainRecordUpdateOptions{
			Type: "A", Name: "", Target: "targetFoo",
			Priority: getPriority(), Weight: getWeight(), Port: getPort(), TTLSec: 300,
		},
	).Return(&linodego.DomainRecord{}, nil).Once()

	mockDomainClient.On(
		"CreateDomainRecord",
		mock.Anything,
		2,
		linodego.DomainRecordCreateOptions{
			Type: "A", Name: "create", Target: "targetBar",
			Priority: getPriority(), Weight: getWeight(), Port: getPort(), TTLSec: 0,
		},
	).Return(&linodego.DomainRecord{}, nil).Once()

	mockDomainClient.On(
		"CreateDomainRecord",
		mock.Anything,
		2,
		linodego.DomainRecordCreateOptions{
			Type: "A", Name: "", Target: "targetBar",
			Priority: getPriority(), Weight: getWeight(), Port: getPort(), TTLSec: 0,
		},
	).Return(&linodego.DomainRecord{}, nil).Once()

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

func TestLinodeApplyChangesTargetAdded(t *testing.T) {
	mockDomainClient := MockDomainClient{}

	provider := &LinodeProvider{
		Client:       &mockDomainClient,
		domainFilter: endpoint.NewDomainFilter([]string{}),
		DryRun:       false,
	}

	// Dummy Data
	mockDomainClient.On(
		"ListDomains",
		mock.Anything,
		mock.Anything,
	).Return([]*linodego.Domain{{Domain: "example.com", ID: 1}}, nil).Once()

	mockDomainClient.On(
		"ListDomainRecords",
		mock.Anything,
		1,
		mock.Anything,
	).Return([]*linodego.DomainRecord{{ID: 11, Name: "", Type: "A", Target: "targetA"}}, nil).Once()

	// Apply Actions
	mockDomainClient.On(
		"UpdateDomainRecord",
		mock.Anything,
		1,
		11,
		linodego.DomainRecordUpdateOptions{
			Type: "A", Name: "", Target: "targetA",
			Priority: getPriority(), Weight: getWeight(), Port: getPort(),
		},
	).Return(&linodego.DomainRecord{}, nil).Once()

	mockDomainClient.On(
		"CreateDomainRecord",
		mock.Anything,
		1,
		linodego.DomainRecordCreateOptions{
			Type: "A", Name: "", Target: "targetB",
			Priority: getPriority(), Weight: getWeight(), Port: getPort(),
		},
	).Return(&linodego.DomainRecord{}, nil).Once()

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

func TestLinodeApplyChangesTargetRemoved(t *testing.T) {
	mockDomainClient := MockDomainClient{}

	provider := &LinodeProvider{
		Client:       &mockDomainClient,
		domainFilter: endpoint.NewDomainFilter([]string{}),
		DryRun:       false,
	}

	// Dummy Data
	mockDomainClient.On(
		"ListDomains",
		mock.Anything,
		mock.Anything,
	).Return([]*linodego.Domain{{Domain: "example.com", ID: 1}}, nil).Once()

	mockDomainClient.On(
		"ListDomainRecords",
		mock.Anything,
		1,
		mock.Anything,
	).Return([]*linodego.DomainRecord{{ID: 11, Name: "", Type: "A", Target: "targetA"}, {ID: 12, Type: "A", Name: "", Target: "targetB"}}, nil).Once()

	// Apply Actions
	mockDomainClient.On(
		"UpdateDomainRecord",
		mock.Anything,
		1,
		12,
		linodego.DomainRecordUpdateOptions{
			Type: "A", Name: "", Target: "targetB",
			Priority: getPriority(), Weight: getWeight(), Port: getPort(),
		},
	).Return(&linodego.DomainRecord{}, nil).Once()

	mockDomainClient.On(
		"DeleteDomainRecord",
		mock.Anything,
		1,
		11,
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

func TestLinodeApplyChangesNoChanges(t *testing.T) {
	mockDomainClient := MockDomainClient{}

	provider := &LinodeProvider{
		Client:       &mockDomainClient,
		domainFilter: endpoint.NewDomainFilter([]string{}),
		DryRun:       false,
	}

	// Dummy Data
	mockDomainClient.On(
		"ListDomains",
		mock.Anything,
		mock.Anything,
	).Return([]*linodego.Domain{{Domain: "example.com", ID: 1}}, nil).Once()

	mockDomainClient.On(
		"ListDomainRecords",
		mock.Anything,
		1,
		mock.Anything,
	).Return([]*linodego.DomainRecord{{ID: 11, Name: "", Type: "A", Target: "targetA"}}, nil).Once()

	err := provider.ApplyChanges(context.Background(), &plan.Changes{})
	require.NoError(t, err)

	mockDomainClient.AssertExpectations(t)
}
