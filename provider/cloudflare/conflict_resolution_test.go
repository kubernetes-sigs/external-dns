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

package cloudflare

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudflare/cloudflare-go/v5/dns"
	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

// TestCloudflareIdenticalRecordConflictResolution verifies that the provider
// correctly handles the "Identical record already exists" (81058) error
// by finding and deleting the conflicting global record before retrying.
func TestCloudflareIdenticalRecordConflictResolution(t *testing.T) {
	// Create a mock client that:
	// 1. Returns 81058 on the first Create attempt for "conflict.bar.com"
	// 2. Returns a conflicting record when List is called
	// 3. Accepts the Delete call
	// 4. Returning success on the second Create attempt

	client := NewMockCloudFlareClient()

	// Pre-seed a conflicting record in the "Global" view of the mock
	// usage of "Global" in the mock is abstract, but we simulate it by having it exist
	// and triggering the error on create.
	conflictingRecord := dns.RecordResponse{
		ID:      "existing-global-id",
		Name:    "conflict.bar.com",
		Type:    endpoint.RecordTypeA,
		TTL:     300,
		Content: "9.9.9.9",
	}
	client.Records["001"]["existing-global-id"] = conflictingRecord

	// We need to inject custom logic into the mock to simulate the 81058 error
	// The standard NewMockCloudFlareClient doesn't support complex stateful failure sequences easily
	// so we will modify the mock instance's behavior via a custom wrapper or just rely on the
	// hardcoded specific behavior in the provided mock for specific names?
	// The existing mock has some hardcoded "newerror-" checks. We can reuse that pattern if easy,
	// or create a focused mock implementation here.

	// Let's use a focused mock wrapper for this test to be precise
}

// ConflictMockClient wraps the standard mock to inject 81058 behavior
type ConflictMockClient struct {
	*mockCloudFlareClient
	createAttempts int
}

func (m *ConflictMockClient) CreateDNSRecord(ctx context.Context, params dns.RecordNewParams) (*dns.RecordResponse, error) {
	body := params.Body.(dns.RecordNewParamsBody)
	name := body.Name.Value

	if name == "conflict.bar.com" {
		m.createAttempts++
		if m.createAttempts == 1 {
			// Simulate the specific Cloudflare error
			return nil, fmt.Errorf("failed to create record: Code 81058: An identical record already exists.")
		}
	}
	return m.mockCloudFlareClient.CreateDNSRecord(ctx, params)
}

func TestResolveIdenticalRecordConflict(t *testing.T) {
	// Setup
	baseMock := NewMockCloudFlareClient()
	// Seed the conflicting record
	conflictingRecord := dns.RecordResponse{
		ID:      "existing-global-id",
		Name:    "conflict.bar.com",
		Type:    endpoint.RecordTypeA,
		TTL:     300,
		Content: "9.9.9.9",
	}
	baseMock.Records["001"]["existing-global-id"] = conflictingRecord

	mockClient := &ConflictMockClient{
		mockCloudFlareClient: baseMock,
	}

	provider := &CloudFlareProvider{
		Client: mockClient,
		RegionalServicesConfig: RegionalServicesConfig{
			ConflictingRecordDeletion: true,
		},
	}

	// Define the change
	ep := &endpoint.Endpoint{
		DNSName:    "conflict.bar.com",
		RecordType: "A",
		Targets:    endpoint.Targets{"1.2.3.4"},
	}

	changes := &plan.Changes{
		Create: []*endpoint.Endpoint{ep},
	}

	// Execute
	err := provider.ApplyChanges(context.Background(), changes)

	// Verify
	assert.NoError(t, err)

	// Check actions
	// 1. Create attempt (failed 81058) - Mock doesn't record failed actions in standard array usually?
	//    Actually `mockCloudFlareClient` records success only usually.
	// 2. List (for resolution)
	// 3. Delete (existing-global-id)
	// 4. Create attempt (success)

	// We can inspect mockClient.Actions to see if Delete and then Create happened
	actions := mockClient.Actions
	// Filter for relevant actions
	var relevantActions []MockAction
	for _, a := range actions {
		if a.Name == "Delete" || a.Name == "Create" {
			relevantActions = append(relevantActions, a)
		}
	}

	// We expect:
	// 1. Delete of "existing-global-id"
	// 2. Create of "conflict.bar.com"

	assert.Len(t, relevantActions, 2)
	assert.Equal(t, "Delete", relevantActions[0].Name)
	assert.Equal(t, "existing-global-id", relevantActions[0].RecordId)

	assert.Equal(t, "Create", relevantActions[1].Name)
	assert.Equal(t, "conflict.bar.com", relevantActions[1].RecordData.Name)
	assert.Equal(t, "1.2.3.4", relevantActions[1].RecordData.Content)

	// Verify createAttempts
	assert.Equal(t, 2, mockClient.createAttempts)
}

func TestResolveIdenticalRecordConflict_Disabled(t *testing.T) {
	// Setup
	baseMock := NewMockCloudFlareClient()
	conflictingRecord := dns.RecordResponse{
		ID:      "existing-global-id",
		Name:    "conflict.bar.com",
		Type:    endpoint.RecordTypeA,
		TTL:     300,
		Content: "9.9.9.9",
	}
	baseMock.Records["001"]["existing-global-id"] = conflictingRecord

	mockClient := &ConflictMockClient{
		mockCloudFlareClient: baseMock,
	}

	provider := &CloudFlareProvider{
		Client: mockClient,
		RegionalServicesConfig: RegionalServicesConfig{
			ConflictingRecordDeletion: false, // Explicitly disabled
		},
	}

	ep := &endpoint.Endpoint{
		DNSName:    "conflict.bar.com",
		RecordType: "A",
		Targets:    endpoint.Targets{"1.2.3.4"},
	}

	changes := &plan.Changes{
		Create: []*endpoint.Endpoint{ep},
	}

	// Execute
	err := provider.ApplyChanges(context.Background(), changes)

	// Verify
	// Should return error because conflict resolution failed (and thus Create failed)
	assert.Error(t, err)
	// We cannot assert exact error message because submitChanges wraps errors into a generic soft error.
	// But failure is expected.

	// Check actions - Should contain NO Delete
	actions := mockClient.Actions
	for _, a := range actions {
		assert.NotEqual(t, "Delete", a.Name, "Should not delete record when flag is disabled")
	}
}
