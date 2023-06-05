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

package adguardhome

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

type mockAdGuardHomeAPI struct {
	listRecordsFunc  func(ctx context.Context) ([]*endpoint.Endpoint, error)
	createRecordFunc func(ctx context.Context, ep *endpoint.Endpoint) error
	updateRecordFunc func(ctx context.Context, old, new *endpoint.Endpoint) error
	deleteRecordFunc func(ctx context.Context, ep *endpoint.Endpoint) error
}

func (m *mockAdGuardHomeAPI) listRecords(ctx context.Context) ([]*endpoint.Endpoint, error) {
	if m.listRecordsFunc != nil {
		return m.listRecordsFunc(ctx)
	}
	return nil, nil
}

func (m *mockAdGuardHomeAPI) createRecord(ctx context.Context, ep *endpoint.Endpoint) error {
	if m.createRecordFunc != nil {
		return m.createRecordFunc(ctx, ep)
	}
	return nil
}

func (m *mockAdGuardHomeAPI) updateRecord(ctx context.Context, old, new *endpoint.Endpoint) error {
	if m.updateRecordFunc != nil {
		return m.updateRecordFunc(ctx, old, new)
	}
	return nil
}

func (m *mockAdGuardHomeAPI) deleteRecord(ctx context.Context, ep *endpoint.Endpoint) error {
	if m.deleteRecordFunc != nil {
		return m.deleteRecordFunc(ctx, ep)
	}
	return nil
}

func TestAdGuardHomeProvider_Records(t *testing.T) {
	ctx := context.TODO()

	// Create a mock AdGuardHomeProvider with a fake API
	provider := &AdGuardHomeProvider{
		api: &mockAdGuardHomeAPI{
			listRecordsFunc: func(ctx context.Context) ([]*endpoint.Endpoint, error) {
				return []*endpoint.Endpoint{
					{
						DNSName:    "example.com",
						Targets:    []string{"1.2.3.4"},
						RecordType: "A",
					},
					{
						DNSName:    "sub.example.com",
						Targets:    []string{"2.3.4.5"},
						RecordType: "A",
					},
					{
						DNSName:    "www.example.com",
						Targets:    []string{"cname.example.com"},
						RecordType: "CNAME",
					},
					{
						DNSName:    "ipv6.example.com",
						Targets:    []string{"2001:db8::1"},
						RecordType: "AAAA",
					},
				}, nil
			},
		},
	}

	// Call the Records function
	endpoints, err := provider.Records(ctx)

	// Verify the results
	assert.NoError(t, err)
	assert.Len(t, endpoints, 4)

	// A record
	assert.Equal(t, "example.com", endpoints[0].DNSName)
	assert.Equal(t, "1.2.3.4", endpoints[0].Targets[0])
	assert.Equal(t, "A", endpoints[0].RecordType)

	// A record
	assert.Equal(t, "sub.example.com", endpoints[1].DNSName)
	assert.Equal(t, "2.3.4.5", endpoints[1].Targets[0])
	assert.Equal(t, "A", endpoints[1].RecordType)

	// CNAME record
	assert.Equal(t, "www.example.com", endpoints[2].DNSName)
	assert.Equal(t, "cname.example.com", endpoints[2].Targets[0])
	assert.Equal(t, "CNAME", endpoints[2].RecordType)

	// AAAA record
	assert.Equal(t, "ipv6.example.com", endpoints[3].DNSName)
	assert.Equal(t, "2001:db8::1", endpoints[3].Targets[0])
	assert.Equal(t, "AAAA", endpoints[3].RecordType)
}

func TestAdGuardHomeProvider_ApplyChanges(t *testing.T) {
	ctx := context.TODO()

	// Create a mock AdGuardHomeProvider with a fake API
	provider := &AdGuardHomeProvider{
		api: &mockAdGuardHomeAPI{
			createRecordFunc: func(ctx context.Context, ep *endpoint.Endpoint) error {
				// Simulate creating a record
				return nil
			},
			updateRecordFunc: func(ctx context.Context, old, new *endpoint.Endpoint) error {
				// Simulate updating a record
				return nil
			},
			deleteRecordFunc: func(ctx context.Context, ep *endpoint.Endpoint) error {
				// Simulate deleting a record
				return nil
			},
		},
	}

	// Create a mock Changes object with sample endpoints
	changes := &plan.Changes{
		Create: []*endpoint.Endpoint{
			{
				DNSName:    "example.com",
				Targets:    []string{"1.2.3.4"},
				RecordType: "A",
			},
		},
		UpdateOld: []*endpoint.Endpoint{
			{
				DNSName:    "old.example.com",
				Targets:    []string{"2.3.4.5"},
				RecordType: "A",
			},
		},
		UpdateNew: []*endpoint.Endpoint{
			{
				DNSName:    "new.example.com",
				Targets:    []string{"3.4.5.6"},
				RecordType: "A",
			},
		},
		Delete: []*endpoint.Endpoint{
			{
				DNSName:    "deleted.example.com",
				Targets:    []string{"4.5.6.7"},
				RecordType: "A",
			},
		},
	}

	// Call the ApplyChanges function
	err := provider.ApplyChanges(ctx, changes)

	// Verify the results
	assert.NoError(t, err)
}
