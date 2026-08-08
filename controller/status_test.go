/*
Copyright 2026 The Kubernetes Authors.

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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	"sigs.k8s.io/external-dns/pkg/events"
	"sigs.k8s.io/external-dns/plan"
	registryfactory "sigs.k8s.io/external-dns/registry/factory"
	"sigs.k8s.io/external-dns/source"
)

// fakeStatusReporter records what reportSyncStatus handed it.
type fakeStatusReporter struct {
	calls    int
	objects  []source.PlannedObject
	applyErr error
}

func (f *fakeStatusReporter) ReportStatus(_ context.Context, objects []source.PlannedObject, applyErr error) {
	f.calls++
	f.objects = objects
	f.applyErr = applyErr
}

// plannedFor returns the endpoint count reported for the object named name.
func plannedFor(objects []source.PlannedObject, name string) int {
	for _, obj := range objects {
		if obj.Ref.Name() == name {
			return obj.Endpoints
		}
	}
	return -1
}

// A resource whose records already exist at the provider produces no changes,
// but it still has to learn that it is programmed — otherwise its Ready
// condition stays empty forever.
func TestRunOnceReportsStatusWhenPlanHasNoChanges(t *testing.T) {
	ref := events.NewObjectReferenceFromParts("DNSEndpoint", "externaldns.k8s.io/v1alpha1", "ns", "steady", "", "crd")

	src := new(testutils.MockSource)
	src.On("Endpoints").Return([]*endpoint.Endpoint{
		endpoint.NewEndpoint("steady.example.com", endpoint.RecordTypeA, "1.2.3.4").WithRefObject(ref),
	}, nil)

	// The provider already holds exactly the desired record.
	prov := newMockProvider(
		[]*endpoint.Endpoint{{DNSName: "steady.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}}},
		&plan.Changes{},
	)

	cfg := getTestConfig()
	reg, err := registryfactory.Select(cfg, prov)
	require.NoError(t, err)

	reporter := &fakeStatusReporter{}
	ctrl := &Controller{
		Source:             src,
		Registry:           reg,
		Policy:             &plan.SyncPolicy{},
		ManagedRecordTypes: cfg.ManagedDNSRecordTypes,
		StatusReporters:    []source.StatusReporter{reporter},
	}

	require.NoError(t, ctrl.RunOnce(t.Context()))

	require.Equal(t, 1, reporter.calls)
	assert.NoError(t, reporter.applyErr)
	assert.Equal(t, 1, plannedFor(reporter.objects, "steady"))
}

func TestReportSyncStatus(t *testing.T) {
	// Two endpoints from the same object plus one from another: a reporter must
	// see each object once, not once per endpoint it produced.
	first := events.NewObjectReferenceFromParts("DNSEndpoint", "externaldns.k8s.io/v1alpha1", "ns", "first", "", "crd")
	second := events.NewObjectReferenceFromParts("DNSEndpoint", "externaldns.k8s.io/v1alpha1", "ns", "second", "", "crd")

	desired := []*endpoint.Endpoint{
		endpoint.NewEndpoint("a.example.com", endpoint.RecordTypeA, "10.0.0.1").WithRefObject(first),
		endpoint.NewEndpoint("b.example.com", endpoint.RecordTypeA, "10.0.0.2").WithRefObject(first),
		endpoint.NewEndpoint("c.example.com", endpoint.RecordTypeA, "10.0.0.3").WithRefObject(second),
	}
	p := &plan.Plan{Desired: desired, Planned: desired}

	t.Run("dedupes objects and forwards the apply error", func(t *testing.T) {
		reporter := &fakeStatusReporter{}
		applyErr := errors.New("provider unavailable")

		reportSyncStatus(t.Context(), []source.StatusReporter{reporter}, p, applyErr)

		assert.Equal(t, 1, reporter.calls)
		assert.Len(t, reporter.objects, 2)
		assert.Equal(t, applyErr, reporter.applyErr)
	})

	t.Run("counts the planned endpoints each object contributed", func(t *testing.T) {
		reporter := &fakeStatusReporter{}

		reportSyncStatus(t.Context(), []source.StatusReporter{reporter}, p, nil)

		assert.Equal(t, 2, plannedFor(reporter.objects, "first"))
		assert.Equal(t, 1, plannedFor(reporter.objects, "second"))
	})

	// An object whose endpoints were all dropped by --domain-filter or the
	// managed record types is still reported, at zero, so its source can say so
	// rather than claim the records were programmed.
	t.Run("reports filtered-out objects at zero", func(t *testing.T) {
		reporter := &fakeStatusReporter{}
		filtered := &plan.Plan{
			Desired: desired,
			Planned: []*endpoint.Endpoint{desired[0], desired[1]},
		}

		reportSyncStatus(t.Context(), []source.StatusReporter{reporter}, filtered, nil)

		assert.Len(t, reporter.objects, 2)
		assert.Equal(t, 2, plannedFor(reporter.objects, "first"))
		assert.Equal(t, 0, plannedFor(reporter.objects, "second"))
	})

	t.Run("every reporter is called", func(t *testing.T) {
		first, second := &fakeStatusReporter{}, &fakeStatusReporter{}

		reportSyncStatus(t.Context(), []source.StatusReporter{first, second}, p, nil)

		assert.Equal(t, 1, first.calls)
		assert.Equal(t, 1, second.calls)
	})

	t.Run("no reporters is a no-op", func(t *testing.T) {
		assert.NotPanics(t, func() {
			reportSyncStatus(t.Context(), nil, p, nil)
		})
	})

	t.Run("endpoints without ref objects skip the reporters", func(t *testing.T) {
		reporter := &fakeStatusReporter{}
		refless := &plan.Plan{
			Desired: []*endpoint.Endpoint{endpoint.NewEndpoint("d.example.com", endpoint.RecordTypeA, "10.0.0.4")},
		}

		reportSyncStatus(t.Context(), []source.StatusReporter{reporter}, refless, nil)

		assert.Equal(t, 0, reporter.calls)
	})
}
