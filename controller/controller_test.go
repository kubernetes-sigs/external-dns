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

package controller

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/registry"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// mockProvider returns mock endpoints and validates changes.
type mockProvider struct {
	provider.BaseProvider
	RecordsStore  []*endpoint.Endpoint
	ExpectChanges *plan.Changes
}

// Records returns the desired mock endpoints.
func (p *mockProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	return p.RecordsStore, nil
}

// ApplyChanges validates that the passed in changes satisfy the assumptions.
func (p *mockProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	if len(changes.Create) != len(p.ExpectChanges.Create) {
		return errors.New("number of created records is wrong")
	}

	for i := range changes.Create {
		if changes.Create[i].DNSName != p.ExpectChanges.Create[i].DNSName || !changes.Create[i].Targets.Same(p.ExpectChanges.Create[i].Targets) {
			return errors.New("created record is wrong")
		}
	}

	for i := range changes.UpdateNew {
		if changes.UpdateNew[i].DNSName != p.ExpectChanges.UpdateNew[i].DNSName || !changes.UpdateNew[i].Targets.Same(p.ExpectChanges.UpdateNew[i].Targets) {
			return errors.New("delete record is wrong")
		}
	}

	for i := range changes.UpdateOld {
		if changes.UpdateOld[i].DNSName != p.ExpectChanges.UpdateOld[i].DNSName || !changes.UpdateOld[i].Targets.Same(p.ExpectChanges.UpdateOld[i].Targets) {
			return errors.New("delete record is wrong")
		}
	}

	for i := range changes.Delete {
		if changes.Delete[i].DNSName != p.ExpectChanges.Delete[i].DNSName || !changes.Delete[i].Targets.Same(p.ExpectChanges.Delete[i].Targets) {
			return errors.New("delete record is wrong")
		}
	}

	if !reflect.DeepEqual(ctx.Value(provider.RecordsContextKey), p.RecordsStore) {
		return errors.New("context is wrong")
	}
	return nil
}

// newMockProvider creates a new mockProvider returning the given endpoints and validating the desired changes.
func newMockProvider(endpoints []*endpoint.Endpoint, changes *plan.Changes) provider.Provider {
	dnsProvider := &mockProvider{
		RecordsStore:  endpoints,
		ExpectChanges: changes,
	}

	return dnsProvider
}

// TestRunOnce tests that RunOnce correctly orchestrates the different components.
func TestRunOnce(t *testing.T) {
	// Fake some desired endpoints coming from our source.
	source := new(testutils.MockSource)
	source.On("Endpoints").Return([]*endpoint.Endpoint{
		{
			DNSName:    "create-record",
			RecordType: endpoint.RecordTypeA,
			Targets:    endpoint.Targets{"1.2.3.4"},
		},
		{
			DNSName:    "update-record",
			RecordType: endpoint.RecordTypeA,
			Targets:    endpoint.Targets{"8.8.4.4"},
		},
	}, nil)

	// Fake some existing records in our DNS provider and validate some desired changes.
	provider := newMockProvider(
		[]*endpoint.Endpoint{
			{
				DNSName:    "update-record",
				RecordType: endpoint.RecordTypeA,
				Targets:    endpoint.Targets{"8.8.8.8"},
			},
			{
				DNSName:    "delete-record",
				RecordType: endpoint.RecordTypeA,
				Targets:    endpoint.Targets{"4.3.2.1"},
			},
		},
		&plan.Changes{
			Create: []*endpoint.Endpoint{
				{DNSName: "create-record", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
			},
			UpdateNew: []*endpoint.Endpoint{
				{DNSName: "update-record", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"8.8.4.4"}},
			},
			UpdateOld: []*endpoint.Endpoint{
				{DNSName: "update-record", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"8.8.8.8"}},
			},
			Delete: []*endpoint.Endpoint{
				{DNSName: "delete-record", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"4.3.2.1"}},
			},
		},
	)

	r, err := registry.NewNoopRegistry(provider)
	require.NoError(t, err)

	// Run our controller once to trigger the validation.
	ctrl := &Controller{
		Source:   source,
		Registry: r,
		Policy:   &plan.SyncPolicy{},
	}

	assert.NoError(t, ctrl.RunOnce(context.Background()))

	// Validate that the mock source was called.
	source.AssertExpectations(t)
}

func TestShouldRunOnce(t *testing.T) {
	ctrl := &Controller{Interval: 10 * time.Minute, MinEventSyncInterval: 5 * time.Second}

	now := time.Now()

	// First run of Run loop should execute RunOnce
	assert.True(t, ctrl.ShouldRunOnce(now))

	// Second run should not
	assert.False(t, ctrl.ShouldRunOnce(now))

	now = now.Add(10 * time.Second)
	// Changes happen in ingresses or services
	ctrl.ScheduleRunOnce(now)
	ctrl.ScheduleRunOnce(now)

	// Because we batch changes, ShouldRunOnce returns False at first
	assert.False(t, ctrl.ShouldRunOnce(now))
	assert.False(t, ctrl.ShouldRunOnce(now.Add(100*time.Microsecond)))

	// But after MinInterval we should run reconciliation
	now = now.Add(5 * time.Second)
	assert.True(t, ctrl.ShouldRunOnce(now))

	// But just one time
	assert.False(t, ctrl.ShouldRunOnce(now))

	// We should wait maximum possible time after last reconciliation started
	now = now.Add(10*time.Minute - time.Second)
	assert.False(t, ctrl.ShouldRunOnce(now))

	// After exactly Interval it's OK again to reconcile
	now = now.Add(time.Second)
	assert.True(t, ctrl.ShouldRunOnce(now))

	// But not two times
	assert.False(t, ctrl.ShouldRunOnce(now))
}
