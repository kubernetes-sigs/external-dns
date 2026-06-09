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
	"strings"
	"sync"
	"testing"
	"time"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	logtest "sigs.k8s.io/external-dns/internal/testutils/log"
	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
	"sigs.k8s.io/external-dns/pkg/events"
	"sigs.k8s.io/external-dns/pkg/events/fake"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/provider/fakes"
	"sigs.k8s.io/external-dns/registry"
	registryfactory "sigs.k8s.io/external-dns/registry/factory"
	"sigs.k8s.io/external-dns/registry/noop"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
	dto "github.com/prometheus/client_model/go"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

func (p *filteredMockProvider) GetDomainFilter() endpoint.DomainFilterInterface {
	return p.domainFilter
}

// Records returns the desired mock endpoints.
func (p *filteredMockProvider) Records(_ context.Context) ([]*endpoint.Endpoint, error) {
	p.RecordsCallCount++
	return p.RecordsStore, nil
}

// ApplyChanges stores all calls for later check
func (p *filteredMockProvider) ApplyChanges(_ context.Context, changes *plan.Changes) error {
	p.ApplyChangesCalls = append(p.ApplyChangesCalls, changes)
	return nil
}

// Records returns the desired mock endpoints.
func (p *mockProvider) Records(_ context.Context) ([]*endpoint.Endpoint, error) {
	return p.RecordsStore, nil
}

// ApplyChanges validates that the passed in changes satisfy the assumptions.
func (p *mockProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	if err := verifyEndpoints(changes.Create, p.ExpectChanges.Create); err != nil {
		return err
	}

	if err := verifyEndpoints(changes.UpdateNew, p.ExpectChanges.UpdateNew); err != nil {
		return err
	}

	if err := verifyEndpoints(changes.UpdateOld, p.ExpectChanges.UpdateOld); err != nil {
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
	cfg.Registry = externaldns.RegistryNoop
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
			UpdateNew: []*endpoint.Endpoint{
				{DNSName: "update-aaaa-record", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:DB8::2"}},
				{DNSName: "update-record", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"8.8.4.4"}},
			},
			UpdateOld: []*endpoint.Endpoint{
				{DNSName: "update-aaaa-record", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:DB8::3"}},
				{DNSName: "update-record", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"8.8.8.8"}},
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

	emitter := fake.NewFakeEventEmitter()

	r, err := registryfactory.Select(cfg, provider)
	require.NoError(t, err)

	// Run our controller once to trigger the validation.
	ctrl := &Controller{
		Source:             source,
		Registry:           r,
		Policy:             &plan.SyncPolicy{},
		ManagedRecordTypes: cfg.ManagedDNSRecordTypes,
		EventEmitter:       emitter,
	}

	assert.NoError(t, ctrl.RunOnce(t.Context()))

	// Validate that the mock source was called.
	source.AssertExpectations(t)
	// check the verified records

	testutils.TestHelperVerifyMetricsGaugeVectorWithLabels(t, 1, verifiedRecords.Gauge, map[string]string{"record_type": "a"})
	testutils.TestHelperVerifyMetricsGaugeVectorWithLabels(t, 1, verifiedRecords.Gauge, map[string]string{"record_type": "aaaa"})

	emitter.AssertNumberOfCalls(t, "Add", 6)
}

// ownerIDRegistry wraps a Registry to override OwnerID — used by the
// suppressed-deletion metric test to exercise the owned=true/false branch
// without depending on TXT-registry config ceremony.
type ownerIDRegistry struct {
	registry.Registry
	ownerID string
}

func (r ownerIDRegistry) OwnerID() string { return r.ownerID }

// TestRunOnce_SuppressedDeletionMetric verifies the deletions_skipped_by_policy
// metric wiring for all three label branches:
//   - no OwnerID (noop registry) → owned="unknown"
//   - OwnerID set, mixed ownership → owned="true" + owned="false"
//   - no suppressions → series cleared
//
// Each subtest also asserts that unrelated label values are NOT emitted, so a
// future bug that mis-routes a series (e.g. emitting owned="true" alongside
// owned="unknown") would be caught rather than hidden by CollectAndCount.
func TestRunOnce_SuppressedDeletionMetric(t *testing.T) {
	keepRecord := &endpoint.Endpoint{DNSName: "keep-record", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.1.1.1"}}
	ownedRecord := &endpoint.Endpoint{
		DNSName: "owned-record", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"2.2.2.2"},
		Labels: endpoint.Labels{endpoint.OwnerLabelKey: "instance-a"},
	}
	foreignRecord := &endpoint.Endpoint{
		DNSName: "foreign-record", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"3.3.3.3"},
		Labels: endpoint.Labels{endpoint.OwnerLabelKey: "instance-b"},
	}

	type labelExpect struct {
		label string
		want  float64
	}

	for _, tc := range []struct {
		name        string
		ownerID     string // empty → use noop registry as-is
		storeExtra  []*endpoint.Endpoint
		wantCount   int
		wantLabels  []labelExpect
		absentLabel []string
	}{
		{
			name:        "noop registry uses owned=unknown",
			ownerID:     "",
			storeExtra:  []*endpoint.Endpoint{{DNSName: "stale-record", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"9.9.9.9"}}},
			wantCount:   1,
			wantLabels:  []labelExpect{{"unknown", 1}},
			absentLabel: []string{"true", "false"},
		},
		{
			name:        "owner id partitions owned vs foreign",
			ownerID:     "instance-a",
			storeExtra:  []*endpoint.Endpoint{ownedRecord, foreignRecord},
			wantCount:   2,
			wantLabels:  []labelExpect{{"true", 1}, {"false", 1}},
			absentLabel: []string{"unknown"},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			deletionsSkippedByPolicy.Reset()
			t.Cleanup(func() { deletionsSkippedByPolicy.Reset() })

			source := new(testutils.MockSource)
			source.On("Endpoints").Return([]*endpoint.Endpoint{keepRecord}, nil)

			p := &filteredMockProvider{
				RecordsStore: append([]*endpoint.Endpoint{keepRecord}, tc.storeExtra...),
			}

			cfg := getTestConfig()
			base, err := registryfactory.Select(cfg, p)
			require.NoError(t, err)
			var r registry.Registry = base
			if tc.ownerID != "" {
				r = ownerIDRegistry{Registry: base, ownerID: tc.ownerID}
			}

			ctrl := &Controller{
				Source:             source,
				Registry:           r,
				Policy:             &plan.UpsertOnlyPolicy{},
				ManagedRecordTypes: cfg.ManagedDNSRecordTypes,
				EventEmitter:       fake.NewFakeEventEmitter(),
			}

			require.NoError(t, ctrl.RunOnce(t.Context()))

			assert.Equal(t, tc.wantCount, testutil.CollectAndCount(deletionsSkippedByPolicy.Gauge),
				"unexpected number of label series published")
			for _, exp := range tc.wantLabels {
				assert.InDelta(t, exp.want,
					testutil.ToFloat64(deletionsSkippedByPolicy.Gauge.WithLabelValues(exp.label)),
					0, "owned=%q series value mismatch", exp.label)
			}
			for _, absent := range tc.absentLabel {
				// Collect BEFORE WithLabelValues so we don't inadvertently
				// create the child while asserting its absence.
				found := false
				for _, m := range collectLabeledValues(t, deletionsSkippedByPolicy.Gauge) {
					if m["owned"] == absent {
						found = true
						break
					}
				}
				assert.Falsef(t, found, "owned=%q series must not be published in this mode", absent)
			}

			// After clearing the extras, the series must be reset on the next reconcile.
			p.RecordsStore = []*endpoint.Endpoint{keepRecord}
			ctrl.Policy = &plan.SyncPolicy{}
			require.NoError(t, ctrl.RunOnce(t.Context()))
			assert.Equal(t, 0, testutil.CollectAndCount(deletionsSkippedByPolicy.Gauge))
		})
	}
}

// collectLabeledValues returns the label maps of all child metrics currently
// registered on a GaugeVec, without side-effectfully instantiating missing
// children (unlike WithLabelValues). Used to assert the absence of a label
// value on deletions_skipped_by_policy.
func collectLabeledValues(t *testing.T, gv prometheus.GaugeVec) []map[string]string {
	t.Helper()
	ch := make(chan prometheus.Metric, 8)
	go func() {
		gv.Collect(ch)
		close(ch)
	}()
	var out []map[string]string
	for m := range ch {
		var dto dto.Metric
		require.NoError(t, m.Write(&dto))
		labels := make(map[string]string, len(dto.Label))
		for _, lp := range dto.Label {
			labels[lp.GetName()] = lp.GetValue()
		}
		out = append(out, labels)
	}
	return out
}

// TestRunOnce_SuppressedPolicyNoOpLog verifies that the reconcile-end
// log distinguishes a true no-op from a reconcile where policy held
// changes back. The trap case is create-only with only updates held
// back: there is no SuppressedDelete signal, so relying on that alone
// produced a misleading "All records are already up to date" info log
// that would mask a drifting source under --policy=create-only.
func TestRunOnce_SuppressedPolicyNoOpLog(t *testing.T) {
	const upToDateMsg = "All records are already up to date"
	const policyHoldMsg = "held back by policy"

	for _, tc := range []struct {
		name          string
		policy        plan.Policy
		sourceRecords []*endpoint.Endpoint
		providerStore []*endpoint.Endpoint
		ownerID       string // "" → noop registry as-is; set → wrap via ownerIDRegistry
		wantSubstr    string
		forbidSubstr  string
	}{
		{
			name:          "create-only with only updates suppressed logs updates",
			policy:        &plan.CreateOnlyPolicy{},
			sourceRecords: []*endpoint.Endpoint{{DNSName: "rec", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.1.1.9"}}},
			providerStore: []*endpoint.Endpoint{{DNSName: "rec", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.1.1.1"}}},
			wantSubstr:    "1 update held back by policy",
			forbidSubstr:  upToDateMsg,
		},
		{
			name:          "upsert-only with only deletions suppressed logs deletions",
			policy:        &plan.UpsertOnlyPolicy{},
			sourceRecords: []*endpoint.Endpoint{{DNSName: "keep", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.1.1.1"}}},
			providerStore: []*endpoint.Endpoint{
				{DNSName: "keep", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.1.1.1"}},
				{DNSName: "stale", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"2.2.2.2"}},
			},
			wantSubstr:   "1 deletion held back by policy",
			forbidSubstr: upToDateMsg,
		},
		{
			name:          "true no-op keeps the existing info log",
			policy:        &plan.SyncPolicy{},
			sourceRecords: []*endpoint.Endpoint{{DNSName: "rec", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.1.1.1"}}},
			providerStore: []*endpoint.Endpoint{{DNSName: "rec", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.1.1.1"}}},
			wantSubstr:    upToDateMsg,
			forbidSubstr:  policyHoldMsg,
		},
		{
			// Shared-zone foreign drift: create-only would nominally
			// suppress an update, but ownership filtering would have
			// discarded the change regardless. The log must NOT blame
			// policy, or operators will chase a false positive.
			name:   "create-only with only foreign-owned update does not blame policy",
			policy: &plan.CreateOnlyPolicy{},
			sourceRecords: []*endpoint.Endpoint{
				{DNSName: "foreign", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"9.9.9.9"},
					Labels: endpoint.Labels{endpoint.OwnerLabelKey: "other"}},
			},
			providerStore: []*endpoint.Endpoint{
				{DNSName: "foreign", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.1.1.1"},
					Labels: endpoint.Labels{endpoint.OwnerLabelKey: "other"}},
			},
			ownerID:      "instance-a",
			wantSubstr:   upToDateMsg,
			forbidSubstr: policyHoldMsg,
		},
		{
			// Same pattern on the deletion path: only a foreign record
			// would be suppressed, ownership filter would have dropped
			// it regardless, so the no-op log must stay quiet about
			// policy.
			name:   "upsert-only with only foreign-owned deletion does not blame policy",
			policy: &plan.UpsertOnlyPolicy{},
			sourceRecords: []*endpoint.Endpoint{
				{DNSName: "keep", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.1.1.1"},
					Labels: endpoint.Labels{endpoint.OwnerLabelKey: "instance-a"}},
			},
			providerStore: []*endpoint.Endpoint{
				{DNSName: "keep", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.1.1.1"},
					Labels: endpoint.Labels{endpoint.OwnerLabelKey: "instance-a"}},
				{DNSName: "foreign-stale", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"2.2.2.2"},
					Labels: endpoint.Labels{endpoint.OwnerLabelKey: "other"}},
			},
			ownerID:      "instance-a",
			wantSubstr:   upToDateMsg,
			forbidSubstr: policyHoldMsg,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			deletionsSkippedByPolicy.Reset()
			t.Cleanup(func() { deletionsSkippedByPolicy.Reset() })

			prevLevel := log.GetLevel()
			t.Cleanup(func() { log.SetLevel(prevLevel) })
			hook := logtest.LogsUnderTestWithLogLevel(log.InfoLevel, t)

			source := new(testutils.MockSource)
			source.On("Endpoints").Return(tc.sourceRecords, nil)
			p := &filteredMockProvider{RecordsStore: tc.providerStore}

			cfg := getTestConfig()
			base, err := registryfactory.Select(cfg, p)
			require.NoError(t, err)
			var r registry.Registry = base
			if tc.ownerID != "" {
				r = ownerIDRegistry{Registry: base, ownerID: tc.ownerID}
			}

			ctrl := &Controller{
				Source:             source,
				Registry:           r,
				Policy:             tc.policy,
				ManagedRecordTypes: cfg.ManagedDNSRecordTypes,
				EventEmitter:       fake.NewFakeEventEmitter(),
			}
			require.NoError(t, ctrl.RunOnce(t.Context()))

			var matched, forbidden bool
			for _, e := range hook.AllEntries() {
				if strings.Contains(e.Message, tc.wantSubstr) {
					matched = true
				}
				if tc.forbidSubstr != "" && strings.Contains(e.Message, tc.forbidSubstr) {
					forbidden = true
				}
			}
			assert.Truef(t, matched, "expected info log containing %q; got entries=%v", tc.wantSubstr, hook.AllEntries())
			assert.Falsef(t, forbidden, "unexpected info log containing %q", tc.forbidSubstr)
		})
	}
}

// TestRun tests that Run correctly starts and stops
func TestRun(t *testing.T) {
	source := getTestSource()
	cfg := getTestConfig()
	provider := getTestProvider()

	r, err := registryfactory.Select(cfg, provider)
	require.NoError(t, err)

	// Run our controller once to trigger the validation.
	ctrl := &Controller{
		Source:             source,
		Registry:           r,
		Policy:             &plan.SyncPolicy{},
		ManagedRecordTypes: cfg.ManagedDNSRecordTypes,
	}
	ctrl.nextRunAt = time.Now().Add(-time.Millisecond)
	ctx, cancel := context.WithCancel(t.Context())
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
	cfg.Registry = externaldns.RegistryNoop
	cfg.ManagedDNSRecordTypes = []string{endpoint.RecordTypeA, endpoint.RecordTypeAAAA, endpoint.RecordTypeCNAME}

	source := new(testutils.MockSource)
	source.On("Endpoints").Return(configuredEndpoints, nil)

	// Fake some existing records in our DNS provider and validate some desired changes.
	provider := &filteredMockProvider{
		RecordsStore: providerEndpoints,
	}
	r, err := registryfactory.Select(cfg, provider)
	require.NoError(t, err)

	ctrl := &Controller{
		Source:             source,
		Registry:           r,
		Policy:             &plan.SyncPolicy{},
		DomainFilter:       domainFilter,
		ManagedRecordTypes: cfg.ManagedDNSRecordTypes,
	}

	assert.NoError(t, ctrl.RunOnce(t.Context()))
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

func TestWhenNoFilterControllerConsidersAllDomains(t *testing.T) {
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
				UpdateOld: []*endpoint.Endpoint{
					{
						DNSName:    "some-record.used.tld",
						RecordType: endpoint.RecordTypeA,
						Targets:    endpoint.Targets{"8.8.8.8"},
						Labels:     endpoint.Labels{},
					},
				},
				UpdateNew: []*endpoint.Endpoint{
					{
						DNSName:    "some-record.used.tld",
						RecordType: endpoint.RecordTypeA,
						Targets:    endpoint.Targets{"1.1.1.1"},
						Labels: endpoint.Labels{
							"owner": "",
						},
					},
				},
			},
		},
	)
}

type toggleRegistry struct {
	noop.NoopRegistry
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

func (r *toggleRegistry) ApplyChanges(_ context.Context, _ *plan.Changes) error {
	return nil
}

func (r *toggleRegistry) AdjustEndpoints(endpoints []*endpoint.Endpoint) ([]*endpoint.Endpoint, error) {
	return endpoints, nil
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
	ctx, cancel := context.WithCancel(t.Context())
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

func TestRunOnce_EmitChangeEvent(t *testing.T) {
	tests := []struct {
		name           string
		applyErr       error
		expectedReason events.Reason
		expectErr      bool
	}{
		{
			name:           "emits RecordReady on success",
			expectedReason: events.RecordReady,
		},
		{
			name:           "emits RecordError on failure",
			applyErr:       errors.New("apply failed"),
			expectedReason: events.RecordError,
			expectErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			source := new(testutils.MockSource)
			source.On("Endpoints").Return([]*endpoint.Endpoint{
				endpoint.NewEndpoint("dot.com", endpoint.RecordTypeA, "1.2.3.4").
					WithRefObject(&events.ObjectReference{}),
			}, nil)

			r, err := registryfactory.Select(getTestConfig(), &fakes.MockProvider{ApplyChangesErr: tt.applyErr})
			require.NoError(t, err)

			emitter := fake.NewFakeEventEmitter()
			ctrl := &Controller{
				Source:             source,
				Registry:           r,
				Policy:             &plan.SyncPolicy{},
				ManagedRecordTypes: []string{endpoint.RecordTypeA},
				EventEmitter:       emitter,
			}

			err = ctrl.RunOnce(t.Context())
			assert.Equal(t, tt.expectErr, err != nil)

			emitter.AssertCalled(t, "Add", mock.MatchedBy(func(e events.Event) bool {
				return e.Reason() == tt.expectedReason
			}))
		})
	}
}

func TestRun_HardError(t *testing.T) {
	cfg := getTestConfig()
	r, err := registryfactory.Select(getTestConfig(), getTestProvider())
	require.NoError(t, err)

	source := new(testutils.MockSource)
	source.On("Endpoints").Return([]*endpoint.Endpoint(nil), errors.New("simulated hard error"))

	ctrl := &Controller{
		Source:             source,
		Registry:           r,
		Policy:             &plan.SyncPolicy{},
		ManagedRecordTypes: cfg.ManagedDNSRecordTypes,
		Interval:           5 * time.Second,
	}

	// Set nextRunAt to the past to trigger ShouldRunOnce immediately on the first tick
	ctrl.nextRunAt = time.Now().Add(-time.Millisecond)
	err = ctrl.Run(t.Context())

	require.Error(t, err)
	assert.ErrorContains(t, err, "failed to do run once")
	assert.ErrorContains(t, err, "simulated hard error")

	source.AssertExpectations(t)
}
