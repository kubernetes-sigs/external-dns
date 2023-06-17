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

package nextdns

import (
	"context"
	"fmt"
	"testing"

	api "github.com/amalucelli/nextdns-go/nextdns"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

type RewriteServiceStub struct {
	mock.Mock
}

func (_m *RewriteServiceStub) Create(ctx context.Context, req *api.CreateRewritesRequest) (string, error) {
	_m.Called(ctx, req)
	return "id", nil
}

func (_m *RewriteServiceStub) List(ctx context.Context, req *api.ListRewritesRequest) ([]*api.Rewrites, error) {
	_m.Called(ctx, req)
	return nil, nil
}

func (_m *RewriteServiceStub) Delete(ctx context.Context, req *api.DeleteRewritesRequest) error {
	_m.Called(ctx, req)
	return nil
}

type RewriteServiceCounter struct {
	wrapped api.RewritesService
	calls   map[string]int
}

func NewRewriteServiceCounter(w api.RewritesService) *RewriteServiceCounter {
	return &RewriteServiceCounter{
		wrapped: w,
		calls:   map[string]int{},
	}
}

func TestNewPiholeProvider(t *testing.T) {
	_, err := NewNextDNSProvider(NextDNSConfig{NextDNSAPIKey: "test"})
	if err == nil {
		t.Error("Expected error from invalid configuration")
	}

	_, err = NewNextDNSProvider(NextDNSConfig{NextDNSProfileId: "test"})
	if err == nil {
		t.Error("Expected error from invalid configuration")
	}

	// Test valid configuration
	_, err = NewNextDNSProvider(NextDNSConfig{NextDNSAPIKey: "test", NextDNSProfileId: "test"})
	if err != nil {
		t.Error("Expected no error from valid configuration, got:", err)
	}
}

func getRewrites() []*api.Rewrites {
	rewrites := []*api.Rewrites{}
	for i := 1; i <= 5; i++ {
		rewrites = append(rewrites, &api.Rewrites{
			ID:      fmt.Sprintf("test-%d", i),
			Name:    fmt.Sprintf("test-%d.example.com", i),
			Type:    "A",
			Content: fmt.Sprintf("192.168.0.%d", i),
		})
	}
	return rewrites
}

func TestRecords(t *testing.T) {
	mockRewrite := &RewriteServiceStub{}
	mockRewrite.On("List", mock.AnythingOfType("*context.emptyCtx"), mock.Anything).Return(getRewrites())

	provider := &NextDnsProvider{
		api:       mockRewrite,
		profileId: "test",
	}
	records, err := provider.Records(context.Background())
	if err != nil {
		t.Error("unexpected error", err)
	}

	for i, record := range records {
		assert.Equal(t, record.DNSName, fmt.Sprintf("test-%d.example.com", i+1))
		assert.Equal(t, record.Targets[0], fmt.Sprintf("192.168.0.%d", i+1))
		assert.Equal(t, record.RecordType, "A")
		id, _ := record.GetProviderSpecificProperty("id")
		assert.Equal(t, id, fmt.Sprintf("test-%d", i+1))
	}
}

func TestRun_ApplyChanges(t *testing.T) {
	mockRewrite := &RewriteServiceStub{}
	mockRewrite.On("Create", mock.Anything, mock.Anything).Return("test-id")
	mockRewrite.On("Delete", mock.Anything, mock.Anything).Return(nil)
	provider := &NextDnsProvider{
		api:       mockRewrite,
		profileId: "test",
	}
	changes := &plan.Changes{
		Create: []*endpoint.Endpoint{
			{
				DNSName:    "test-1.example.com",
				Targets:    endpoint.Targets{"192.168.0.1"},
				RecordType: "A",
			},
		},
		UpdateNew: []*endpoint.Endpoint{
			{
				DNSName:    "test-3.example.com",
				Targets:    endpoint.Targets{"192.168.0.3"},
				RecordType: "A",
			},
		},
		UpdateOld: []*endpoint.Endpoint{
			{
				DNSName:          "test-2.example.com",
				Targets:          endpoint.Targets{"192.168.0.2"},
				RecordType:       "A",
				ProviderSpecific: endpoint.ProviderSpecific{endpoint.ProviderSpecificProperty{Name: "id", Value: "test-id"}},
			},
			{
				DNSName:    "test-2.example.com",
				Targets:    endpoint.Targets{"192.168.0.2"},
				RecordType: "A",
			},
		},
		Delete: []*endpoint.Endpoint{
			{
				DNSName:          "test-4.example.com",
				Targets:          endpoint.Targets{"192.168.0.4"},
				RecordType:       "A",
				ProviderSpecific: endpoint.ProviderSpecific{endpoint.ProviderSpecificProperty{Name: "id", Value: "test-id"}},
			},
			{
				DNSName:    "test-4.example.com",
				Targets:    endpoint.Targets{"192.168.0.4"},
				RecordType: "A",
			},
		},
	}
	provider.ApplyChanges(context.Background(), changes)

	mockRewrite.AssertNumberOfCalls(t, "Create", 2)
	mockRewrite.AssertNumberOfCalls(t, "Delete", 2)
}

func TestDryRun_ApplyChanges(t *testing.T) {
	mockRewrite := &RewriteServiceStub{}
	mockRewrite.On("Create", mock.Anything, mock.Anything).Return("test-id")
	mockRewrite.On("Delete", mock.Anything, mock.Anything).Return(nil)
	provider := &NextDnsProvider{
		api:       mockRewrite,
		profileId: "test",
		dryRun:    true,
	}
	changes := &plan.Changes{
		Create: []*endpoint.Endpoint{
			{
				DNSName:    "test-1.example.com",
				Targets:    endpoint.Targets{"192.168.0.1"},
				RecordType: "A",
			},
		},
		UpdateNew: []*endpoint.Endpoint{
			{
				DNSName:    "test-3.example.com",
				Targets:    endpoint.Targets{"192.168.0.3"},
				RecordType: "A",
			},
		},
		UpdateOld: []*endpoint.Endpoint{
			{
				DNSName:          "test-2.example.com",
				Targets:          endpoint.Targets{"192.168.0.2"},
				RecordType:       "A",
				ProviderSpecific: endpoint.ProviderSpecific{endpoint.ProviderSpecificProperty{Name: "id", Value: "test-id"}},
			},
			{
				DNSName:    "test-2.example.com",
				Targets:    endpoint.Targets{"192.168.0.2"},
				RecordType: "A",
			},
		},
		Delete: []*endpoint.Endpoint{
			{
				DNSName:          "test-4.example.com",
				Targets:          endpoint.Targets{"192.168.0.4"},
				RecordType:       "A",
				ProviderSpecific: endpoint.ProviderSpecific{endpoint.ProviderSpecificProperty{Name: "id", Value: "test-id"}},
			},
			{
				DNSName:    "test-4.example.com",
				Targets:    endpoint.Targets{"192.168.0.4"},
				RecordType: "A",
			},
		},
	}
	provider.ApplyChanges(context.Background(), changes)

	mockRewrite.AssertNumberOfCalls(t, "Create", 0)
	mockRewrite.AssertNumberOfCalls(t, "Delete", 0)
}

func TestOldMatchesNew_ApplyChanges(t *testing.T) {
	mockRewrite := &RewriteServiceStub{}
	mockRewrite.On("Create", mock.Anything, mock.Anything).Return("test-id")
	mockRewrite.On("Delete", mock.Anything, mock.Anything).Return(nil)
	provider := &NextDnsProvider{
		api:       mockRewrite,
		profileId: "test",
	}
	
	changes := &plan.Changes{
		UpdateOld: []*endpoint.Endpoint{
			{
				DNSName:    "test1.example.com",
				Targets:    []string{"192.168.1.1"},
				RecordType: endpoint.RecordTypeA,
				ProviderSpecific: endpoint.ProviderSpecific{endpoint.ProviderSpecificProperty{Name: "id", Value: "test-id"}},
			},
			{
				DNSName:    "test2.example.com",
				Targets:    []string{"192.168.1.2"},
				RecordType: endpoint.RecordTypeA,
				ProviderSpecific: endpoint.ProviderSpecific{endpoint.ProviderSpecificProperty{Name: "id", Value: "test-id"}},
			},
		},
		UpdateNew: []*endpoint.Endpoint{
			{
				DNSName:    "test1.example.com",
				Targets:    []string{"192.168.1.1"},
				RecordType: endpoint.RecordTypeA,
				ProviderSpecific: endpoint.ProviderSpecific{endpoint.ProviderSpecificProperty{Name: "id", Value: "test-id"}},
			},
			{
				DNSName:    "test2.example.com",
				Targets:    []string{"10.0.0.1"},
				RecordType: endpoint.RecordTypeA,
				ProviderSpecific: endpoint.ProviderSpecific{endpoint.ProviderSpecificProperty{Name: "id", Value: "test-id"}},
			},
		},
	}
	provider.ApplyChanges(context.Background(), changes)

	mockRewrite.AssertNumberOfCalls(t, "Create", 1)
	mockRewrite.AssertNumberOfCalls(t, "Delete", 1)
}
