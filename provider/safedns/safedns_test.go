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

package safedns

import (
	"context"
	"os"
	"testing"

	ansConnection "github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/safedns"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

// Create an implementation of the SafeDNS interface for Mocking
type MockSafeDNSService struct {
	mock.Mock
}

func (m *MockSafeDNSService) CreateZoneRecord(zoneName string, req safedns.CreateRecordRequest) (int, error) {
	args := m.Called(zoneName, req)
	return args.Int(0), args.Error(1)
}

func (m *MockSafeDNSService) DeleteZoneRecord(zoneName string, recordID int) error {
	args := m.Called(zoneName, recordID)
	return args.Error(0)
}

func (m *MockSafeDNSService) GetZone(zoneName string) (safedns.Zone, error) {
	args := m.Called(zoneName)
	return args.Get(0).(safedns.Zone), args.Error(1)
}

func (m *MockSafeDNSService) GetZoneRecord(zoneName string, recordID int) (safedns.Record, error) {
	args := m.Called(zoneName, recordID)
	return args.Get(0).(safedns.Record), args.Error(1)
}

func (m *MockSafeDNSService) GetZoneRecords(zoneName string, parameters ansConnection.APIRequestParameters) ([]safedns.Record, error) {
	args := m.Called(zoneName, parameters)
	return args.Get(0).([]safedns.Record), args.Error(1)
}

func (m *MockSafeDNSService) GetZones(parameters ansConnection.APIRequestParameters) ([]safedns.Zone, error) {
	args := m.Called(parameters)
	return args.Get(0).([]safedns.Zone), args.Error(1)
}

func (m *MockSafeDNSService) PatchZoneRecord(zoneName string, recordID int, patch safedns.PatchRecordRequest) (int, error) {
	args := m.Called(zoneName, recordID, patch)
	return args.Int(0), args.Error(1)
}

func (m *MockSafeDNSService) UpdateZoneRecord(zoneName string, record safedns.Record) (int, error) {
	args := m.Called(zoneName, record)
	return args.Int(0), args.Error(1)
}

// Utility functions
func createZones() []safedns.Zone {
	return []safedns.Zone{
		{Name: "foo.com", Description: "Foo dot com"},
		{Name: "bar.io", Description: ""},
		{Name: "baz.org", Description: "Org"},
	}
}

func createFooRecords() []safedns.Record {
	return []safedns.Record{
		{
			ID:      11,
			Type:    safedns.RecordTypeA,
			Name:    "foo.com",
			Content: "targetFoo",
			TTL:     safedns.RecordTTL(3600),
		},
		{
			ID:      12,
			Type:    safedns.RecordTypeTXT,
			Name:    "foo.com",
			Content: "text",
			TTL:     safedns.RecordTTL(3600),
		},
		{
			ID:      13,
			Type:    safedns.RecordTypeCAA,
			Name:    "foo.com",
			Content: "",
			TTL:     safedns.RecordTTL(3600),
		},
	}
}

func createBarRecords() []safedns.Record {
	return []safedns.Record{}
}

func createBazRecords() []safedns.Record {
	return []safedns.Record{
		{
			ID:      31,
			Type:    safedns.RecordTypeA,
			Name:    "baz.org",
			Content: "targetBaz",
			TTL:     safedns.RecordTTL(3600),
		},
		{
			ID:      32,
			Type:    safedns.RecordTypeTXT,
			Name:    "baz.org",
			Content: "text",
			TTL:     safedns.RecordTTL(3600),
		},
		{
			ID:      33,
			Type:    safedns.RecordTypeA,
			Name:    "api.baz.org",
			Content: "targetBazAPI",
			TTL:     safedns.RecordTTL(3600),
		},
		{
			ID:      34,
			Type:    safedns.RecordTypeTXT,
			Name:    "api.baz.org",
			Content: "text",
			TTL:     safedns.RecordTTL(3600),
		},
	}
}

// Actual tests
func TestNewSafeDNSProvider(t *testing.T) {
	_ = os.Setenv("SAFEDNS_TOKEN", "DUMMYVALUE")
	_, err := NewSafeDNSProvider(endpoint.NewDomainFilter([]string{"ext-dns-test.zalando.to."}), true)
	require.NoError(t, err)

	_ = os.Unsetenv("SAFEDNS_TOKEN")
	_, err = NewSafeDNSProvider(endpoint.NewDomainFilter([]string{"ext-dns-test.zalando.to."}), true)
	require.Error(t, err)
}

func TestRecords(t *testing.T) {
	mockSafeDNSService := MockSafeDNSService{}

	provider := &SafeDNSProvider{
		Client:       &mockSafeDNSService,
		domainFilter: endpoint.NewDomainFilter([]string{}),
		DryRun:       false,
	}

	mockSafeDNSService.On(
		"GetZones",
		mock.Anything,
	).Return(createZones(), nil).Once()

	mockSafeDNSService.On(
		"GetZoneRecords",
		"foo.com",
		mock.Anything,
	).Return(createFooRecords(), nil).Once()

	mockSafeDNSService.On(
		"GetZoneRecords",
		"bar.io",
		mock.Anything,
	).Return(createBarRecords(), nil).Once()

	mockSafeDNSService.On(
		"GetZoneRecords",
		"baz.org",
		mock.Anything,
	).Return(createBazRecords(), nil).Once()

	actual, err := provider.Records(context.Background())
	require.NoError(t, err)

	expected := []*endpoint.Endpoint{
		{
			DNSName:    "foo.com",
			Targets:    []string{"targetFoo"},
			RecordType: "A",
			RecordTTL:  3600,
			Labels:     endpoint.NewLabels(),
		},
		{
			DNSName:    "foo.com",
			Targets:    []string{"text"},
			RecordType: "TXT",
			RecordTTL:  3600,
			Labels:     endpoint.NewLabels(),
		},
		{
			DNSName:    "baz.org",
			Targets:    []string{"targetBaz"},
			RecordType: "A",
			RecordTTL:  3600,
			Labels:     endpoint.NewLabels(),
		},
		{
			DNSName:    "baz.org",
			Targets:    []string{"text"},
			RecordType: "TXT",
			RecordTTL:  3600,
			Labels:     endpoint.NewLabels(),
		},
		{
			DNSName:    "api.baz.org",
			Targets:    []string{"targetBazAPI"},
			RecordType: "A",
			RecordTTL:  3600,
			Labels:     endpoint.NewLabels(),
		},
		{
			DNSName:    "api.baz.org",
			Targets:    []string{"text"},
			RecordType: "TXT",
			RecordTTL:  3600,
			Labels:     endpoint.NewLabels(),
		},
	}

	mockSafeDNSService.AssertExpectations(t)
	assert.Equal(t, expected, actual)
}

func TestSafeDNSApplyChanges(t *testing.T) {
	mockSafeDNSService := MockSafeDNSService{}

	provider := &SafeDNSProvider{
		Client:       &mockSafeDNSService,
		domainFilter: endpoint.NewDomainFilter([]string{}),
		DryRun:       false,
	}

	// Dummy data
	mockSafeDNSService.On(
		"GetZones",
		mock.Anything,
	).Return(createZones(), nil).Once()
	mockSafeDNSService.On(
		"GetZones",
		mock.Anything,
	).Return(createZones(), nil).Once()

	mockSafeDNSService.On(
		"GetZoneRecords",
		"foo.com",
		mock.Anything,
	).Return(createFooRecords(), nil).Once()

	mockSafeDNSService.On(
		"GetZoneRecords",
		"bar.io",
		mock.Anything,
	).Return(createBarRecords(), nil).Once()

	mockSafeDNSService.On(
		"GetZoneRecords",
		"baz.org",
		mock.Anything,
	).Return(createBazRecords(), nil).Once()

	// Apply actions
	mockSafeDNSService.On(
		"DeleteZoneRecord",
		"baz.org",
		33,
	).Return(nil).Once()
	mockSafeDNSService.On(
		"DeleteZoneRecord",
		"baz.org",
		34,
	).Return(nil).Once()

	TTL300 := safedns.RecordTTL(300)
	mockSafeDNSService.On(
		"PatchZoneRecord",
		"foo.com",
		11,
		safedns.PatchRecordRequest{
			Type:    "A",
			Name:    "foo.com",
			Content: "targetFoo",
			TTL:     &TTL300,
		},
	).Return(123, nil).Once()

	mockSafeDNSService.On(
		"CreateZoneRecord",
		"bar.io",
		safedns.CreateRecordRequest{
			Type:    "A",
			Name:    "create.bar.io",
			Content: "targetBar",
		},
	).Return(246, nil).Once()

	mockSafeDNSService.On(
		"CreateZoneRecord",
		"bar.io",
		safedns.CreateRecordRequest{
			Type:    "A",
			Name:    "bar.io",
			Content: "targetBar",
		},
	).Return(369, nil).Once()

	err := provider.ApplyChanges(context.Background(), &plan.Changes{
		Create: []*endpoint.Endpoint{
			{
				DNSName:    "create.bar.io",
				RecordType: "A",
				Targets:    []string{"targetBar"},
				RecordTTL:  3600,
			},
			{
				DNSName:    "bar.io",
				RecordType: "A",
				Targets:    []string{"targetBar"},
				RecordTTL:  3600,
			},
		},
		Delete: []*endpoint.Endpoint{
			{
				DNSName:    "api.baz.org",
				RecordType: "A",
			},
			{
				DNSName:    "api.baz.org",
				RecordType: "TXT",
			},
		},
		UpdateNew: []*endpoint.Endpoint{
			{
				DNSName:    "foo.com",
				RecordType: "A",
				RecordTTL:  300,
				Targets:    []string{"targetFoo"},
			},
		},
		UpdateOld: []*endpoint.Endpoint{},
	})
	require.NoError(t, err)

	mockSafeDNSService.AssertExpectations(t)
}
