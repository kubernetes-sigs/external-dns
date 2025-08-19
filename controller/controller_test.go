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
	"sort"
	"sync"
	"testing"
	"time"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
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

type filteredMockProvider struct {
	provider.BaseProvider
	domainFilter      *endpoint.DomainFilter
	RecordsStore      []*endpoint.Endpoint
	RecordsCallCount  int
	ApplyChangesCalls []*plan.Changes
}

type errorMockProvider struct {
	mockProvider
}

func (p *filteredMockProvider) GetDomainFilter() endpoint.DomainFilterInterface {
	return p.domainFilter
}

// Records returns the desired mock endpoints.
func (p *filteredMockProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	p.RecordsCallCount++
	return p.RecordsStore, nil
}

// ApplyChanges stores all calls for later check
func (p *filteredMockProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	p.ApplyChangesCalls = append(p.ApplyChangesCalls, changes)
	return nil
}

// Records returns the desired mock endpoints.
func (p *mockProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	return p.RecordsStore, nil
}

func (p *errorMockProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	return nil, errors.New("error for testing")
}

// ApplyChanges validates that the passed in changes satisfy the assumptions.
func (p *mockProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	if err := verifyEndpoints(changes.Create, p.ExpectChanges.Create); err != nil {
		return err
	}

	if err := verifyEndpoints(changes.UpdateNew(), p.ExpectChanges.UpdateNew()); err != nil {
		return err
	}

	if err := verifyEndpoints(changes.UpdateOld(), p.ExpectChanges.UpdateOld()); err != nil {
		return err
	}

	if err := verifyEndpoints(changes.Delete, p.ExpectChanges.Delete); err != nil {
		return err
	}

	if !reflect.DeepEqual(ctx.Value(provider.RecordsContextKey), p.RecordsStore) {
		return errors.New("context is wrong")
	}
	return nil
}

func verifyEndpoints(actual, expected []*endpoint.Endpoint) error {
	if len(actual) != len(expected) {
		return errors.New("number of records is wrong")
	}
	sort.Slice(actual, func(i, j int) bool {
		return actual[i].DNSName < actual[j].DNSName
	})
	for i := range actual {
		if actual[i].DNSName != expected[i].DNSName || !actual[i].Targets.Same(expected[i].Targets) {
			return errors.New("record is wrong")
		}
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

func getTestSource() *testutils.MockSource {
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
		{
			DNSName:    "create-aaaa-record",
			RecordType: endpoint.RecordTypeAAAA,
			Targets:    endpoint.Targets{"2001:DB8::1"},
		},
		{
			DNSName:    "update-aaaa-record",
			RecordType: endpoint.RecordTypeAAAA,
			Targets:    endpoint.Targets{"2001:DB8::2"},
		},
	}, nil)

	return source
}

func getTestConfig() *externaldns.Config {
	cfg := externaldns.NewConfig()
	cfg.ManagedDNSRecordTypes = []string{endpoint.RecordTypeA, endpoint.RecordTypeAAAA, endpoint.RecordTypeCNAME}
	return cfg
}

func getTestProvider() provider.Provider {
	// Fake some existing records in our DNS provider and validate some desired changes.
	return newMockProvider(
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
			{
				DNSName:    "update-aaaa-record",
				RecordType: endpoint.RecordTypeAAAA,
				Targets:    endpoint.Targets{"2001:DB8::3"},
			},
			{
				DNSName:    "delete-aaaa-record",
				RecordType: endpoint.RecordTypeAAAA,
				Targets:    endpoint.Targets{"2001:DB8::4"},
			},
		},
		&plan.Changes{
			Create: []*endpoint.Endpoint{
				{DNSName: "create-aaaa-record", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:DB8::1"}},
				{DNSName: "create-record", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
			},
			Update: []*plan.Update{
				{
					New: &endpoint.Endpoint{DNSName: "update-aaaa-record", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:DB8::2"}},
					Old: &endpoint.Endpoint{DNSName: "update-aaaa-record", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:DB8::3"}},
				},
				{
					New: &endpoint.Endpoint{DNSName: "update-record", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"8.8.4.4"}},
					Old: &endpoint.Endpoint{DNSName: "update-record", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"8.8.8.8"}},
				},
			},
			Delete: []*endpoint.Endpoint{
				{DNSName: "delete-aaaa-record", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:DB8::4"}},
				{DNSName: "delete-record", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"4.3.2.1"}},
			},
		},
	)
}

// TestRunOnce tests that RunOnce correctly orchestrates the different components.
func TestRunOnce(t *testing.T) {
	source := getTestSource()
	cfg := getTestConfig()
	provider := getTestProvider()

	r, err := registry.NewNoopRegistry(provider)
	require.NoError(t, err)

	// Run our controller once to trigger the validation.
	ctrl := &Controller{
		Source:             source,
		Registry:           r,
		Policy:             &plan.SyncPolicy{},
		ManagedRecordTypes: cfg.ManagedDNSRecordTypes,
	}

	assert.NoError(t, ctrl.RunOnce(context.Background()))

	// Validate that the mock source was called.
	source.AssertExpectations(t)
	// check the verified records

	testutils.TestHelperVerifyMetricsGaugeVectorWithLabels(t, 1, verifiedRecords.Gauge, map[string]string{"record_type": "a"})
	testutils.TestHelperVerifyMetricsGaugeVectorWithLabels(t, 1, verifiedRecords.Gauge, map[string]string{"record_type": "aaaa"})
}

// TestRun tests that Run correctly starts and stops
func TestRun(t *testing.T) {
	source := getTestSource()
	cfg := getTestConfig()
	provider := getTestProvider()

	r, err := registry.NewNoopRegistry(provider)
	require.NoError(t, err)

	// Run our controller once to trigger the validation.
	ctrl := &Controller{
		Source:             source,
		Registry:           r,
		Policy:             &plan.SyncPolicy{},
		ManagedRecordTypes: cfg.ManagedDNSRecordTypes,
	}
	ctrl.nextRunAt = time.Now().Add(-time.Millisecond)
	ctx, cancel := context.WithCancel(context.Background())
	stopped := make(chan struct{})
	go func() {
		ctrl.Run(ctx)
		close(stopped)
	}()
	time.Sleep(1500 * time.Millisecond)
	cancel() // start shutdown
	<-stopped

	// Validate that the mock source was called.
	source.AssertExpectations(t)

	testutils.TestHelperVerifyMetricsGaugeVectorWithLabels(t, 1, verifiedRecords.Gauge, map[string]string{"record_type": "a"})
	testutils.TestHelperVerifyMetricsGaugeVectorWithLabels(t, 1, verifiedRecords.Gauge, map[string]string{"record_type": "aaaa"})
}

func TestShouldRunOnce(t *testing.T) {
	ctrl := &Controller{Interval: 10 * time.Minute, MinEventSyncInterval: 15 * time.Second}

	now := time.Now()

	// First run of Run loop should execute RunOnce
	assert.True(t, ctrl.ShouldRunOnce(now))
	assert.Equal(t, now.Add(10*time.Minute), ctrl.nextRunAt)

	// Second run should not
	assert.False(t, ctrl.ShouldRunOnce(now))
	ctrl.lastRunAt = now

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

	// Multiple ingresses or services changes, closer than MinInterval from each other
	ctrl.lastRunAt = now
	firstChangeTime := now
	secondChangeTime := firstChangeTime.Add(time.Second)
	// First change
	ctrl.ScheduleRunOnce(firstChangeTime)
	// Second change
	ctrl.ScheduleRunOnce(secondChangeTime)

	// Executions should be spaced by at least MinEventSyncInterval
	assert.False(t, ctrl.ShouldRunOnce(now.Add(5*time.Second)))

	// Should not postpone the reconciliation further than firstChangeTime + MinInterval
	now = now.Add(ctrl.MinEventSyncInterval)
	assert.True(t, ctrl.ShouldRunOnce(now))
}

func testControllerFiltersDomains(t *testing.T, configuredEndpoints []*endpoint.Endpoint, domainFilter *endpoint.DomainFilter, providerEndpoints []*endpoint.Endpoint, expectedChanges []*plan.Changes) {
	t.Helper()
	cfg := externaldns.NewConfig()
	cfg.ManagedDNSRecordTypes = []string{endpoint.RecordTypeA, endpoint.RecordTypeAAAA, endpoint.RecordTypeCNAME}

	source := new(testutils.MockSource)
	source.On("Endpoints").Return(configuredEndpoints, nil)

	// Fake some existing records in our DNS provider and validate some desired changes.
	provider := &filteredMockProvider{
		RecordsStore: providerEndpoints,
	}
	r, err := registry.NewNoopRegistry(provider)

	require.NoError(t, err)

	ctrl := &Controller{
		Source:             source,
		Registry:           r,
		Policy:             &plan.SyncPolicy{},
		DomainFilter:       domainFilter,
		ManagedRecordTypes: cfg.ManagedDNSRecordTypes,
	}

	assert.NoError(t, ctrl.RunOnce(context.Background()))
	assert.Equal(t, 1, provider.RecordsCallCount)
	require.Len(t, provider.ApplyChangesCalls, len(expectedChanges))
	for i, change := range expectedChanges {
		assert.Equal(t, *change, *provider.ApplyChangesCalls[i])
	}
}

func TestControllerSkipsEmptyChanges(t *testing.T) {
	testControllerFiltersDomains(
		t,
		[]*endpoint.Endpoint{
			{
				DNSName:    "create-record.other.tld",
				RecordType: endpoint.RecordTypeA,
				Targets:    endpoint.Targets{"1.2.3.4"},
			},
			{
				DNSName:    "some-record.used.tld",
				RecordType: endpoint.RecordTypeA,
				Targets:    endpoint.Targets{"8.8.8.8"},
			},
		},
		endpoint.NewDomainFilter([]string{"used.tld"}),
		[]*endpoint.Endpoint{
			{
				DNSName:    "some-record.used.tld",
				RecordType: endpoint.RecordTypeA,
				Targets:    endpoint.Targets{"8.8.8.8"},
			},
		},
		[]*plan.Changes{},
	)
}

func TestWhenNoFilterControllerConsidersAllComain(t *testing.T) {
	testControllerFiltersDomains(
		t,
		[]*endpoint.Endpoint{
			{
				DNSName:    "create-record.other.tld",
				RecordType: endpoint.RecordTypeA,
				Targets:    endpoint.Targets{"1.2.3.4"},
			},
			{
				DNSName:    "some-record.used.tld",
				RecordType: endpoint.RecordTypeA,
				Targets:    endpoint.Targets{"8.8.8.8"},
			},
		},
		&endpoint.DomainFilter{},
		[]*endpoint.Endpoint{
			{
				DNSName:    "some-record.used.tld",
				RecordType: endpoint.RecordTypeA,
				Targets:    endpoint.Targets{"8.8.8.8"},
			},
		},
		[]*plan.Changes{
			{
				Create: []*endpoint.Endpoint{
					{
						DNSName:    "create-record.other.tld",
						RecordType: endpoint.RecordTypeA,
						Targets:    endpoint.Targets{"1.2.3.4"},
					},
				},
			},
		},
	)
}

func TestWhenMultipleControllerConsidersAllFilteredComain(t *testing.T) {
	testControllerFiltersDomains(
		t,
		[]*endpoint.Endpoint{
			{
				DNSName:    "create-record.other.tld",
				RecordType: endpoint.RecordTypeA,
				Targets:    endpoint.Targets{"1.2.3.4"},
			},
			{
				DNSName:    "some-record.used.tld",
				RecordType: endpoint.RecordTypeA,
				Targets:    endpoint.Targets{"1.1.1.1"},
			},
			{
				DNSName:    "create-record.unused.tld",
				RecordType: endpoint.RecordTypeA,
				Targets:    endpoint.Targets{"1.2.3.4"},
			},
		},
		endpoint.NewDomainFilter([]string{"used.tld", "other.tld"}),
		[]*endpoint.Endpoint{
			{
				DNSName:    "some-record.used.tld",
				RecordType: endpoint.RecordTypeA,
				Targets:    endpoint.Targets{"8.8.8.8"},
			},
		},
		[]*plan.Changes{
			{
				Create: []*endpoint.Endpoint{
					{
						DNSName:    "create-record.other.tld",
						RecordType: endpoint.RecordTypeA,
						Targets:    endpoint.Targets{"1.2.3.4"},
					},
				},
				Update: []*plan.Update{
					{
						Old: &endpoint.Endpoint{

							DNSName:    "some-record.used.tld",
							RecordType: endpoint.RecordTypeA,
							Targets:    endpoint.Targets{"8.8.8.8"},
							Labels:     endpoint.Labels{},
						},
						New: &endpoint.Endpoint{

							DNSName:    "some-record.used.tld",
							RecordType: endpoint.RecordTypeA,
							Targets:    endpoint.Targets{"1.1.1.1"},
							Labels:     endpoint.Labels{},
						},
					},
				},
			},
		},
	)
}

type toggleRegistry struct {
	registry.NoopRegistry
	failCount   int
	failCountMu sync.Mutex // protects failCount
}

const toggleRegistryFailureCount = 3

func (r *toggleRegistry) Records(_ context.Context) ([]*endpoint.Endpoint, error) {
	r.failCountMu.Lock()
	defer r.failCountMu.Unlock()
	if r.failCount < toggleRegistryFailureCount {
		r.failCount++
		return nil, provider.SoftError
	}
	return []*endpoint.Endpoint{}, nil
}

func (r *toggleRegistry) ApplyChanges(_ context.Context, changes *plan.Changes) error {
	return nil
}

func TestToggleRegistry(t *testing.T) {
	source := getTestSource()
	cfg := getTestConfig()
	r := &toggleRegistry{}

	interval := 10 * time.Millisecond
	ctrl := &Controller{
		Source:             source,
		Registry:           r,
		Policy:             &plan.SyncPolicy{},
		ManagedRecordTypes: cfg.ManagedDNSRecordTypes,
		Interval:           interval,
	}
	ctrl.nextRunAt = time.Now().Add(-time.Millisecond)
	ctx, cancel := context.WithCancel(context.Background())
	stopped := make(chan struct{})
	go func() {
		ctrl.Run(ctx)
		close(stopped)
	}()

	// Wait up to 1 minute for failCount to reach at least 3
	// The timeout serves as a safety net against infinite loops while being
	// sufficiently large to accommodate slow CI environments
	deadline := time.Now().Add(15 * time.Second)
	for {
		r.failCountMu.Lock()
		count := r.failCount
		r.failCountMu.Unlock()
		if count >= toggleRegistryFailureCount {
			break
		}
		if time.Now().After(deadline) {
			break
		}
		// Sleep for the controller interval to avoid busy waiting
		// since the controller won't run again until the interval passes
		time.Sleep(interval)
	}
	cancel()
	<-stopped

	r.failCountMu.Lock()
	finalCount := r.failCount
	r.failCountMu.Unlock()
	assert.Equal(t, toggleRegistryFailureCount, finalCount, "failCount should be at least %d", toggleRegistryFailureCount)
}
