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
)

type fakeStatusReporter struct {
	calls    int
	objects  []plan.PlannedObject
	applyErr error
}

func (f *fakeStatusReporter) ReportStatus(_ context.Context, objects []plan.PlannedObject, applyErr error) {
	f.calls++
	f.objects = objects
	f.applyErr = applyErr
}

// An in-sync object produces no changes but must still be reported.
func TestRunOnceReportsStatusWhenPlanHasNoChanges(t *testing.T) {
	ref := events.NewObjectReferenceFromParts("DNSEndpoint", "externaldns.k8s.io/v1alpha1", "ns", "steady", "", "crd")

	src := new(testutils.MockSource)
	src.On("Endpoints").Return([]*endpoint.Endpoint{
		endpoint.NewEndpoint("steady.example.com", endpoint.RecordTypeA, "1.2.3.4").WithRefObject(ref),
	}, nil)

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
		StatusReporter:     reporter,
	}

	require.NoError(t, ctrl.RunOnce(t.Context()))

	require.Equal(t, 1, reporter.calls)
	require.NoError(t, reporter.applyErr)
	require.Len(t, reporter.objects, 1)
	assert.Equal(t, 1, reporter.objects[0].Endpoints)
}

func TestReportSyncStatus(t *testing.T) {
	ref := events.NewObjectReferenceFromParts("DNSEndpoint", "externaldns.k8s.io/v1alpha1", "ns", "first", "", "crd")
	desired := []*endpoint.Endpoint{
		endpoint.NewEndpoint("a.example.com", endpoint.RecordTypeA, "10.0.0.1").WithRefObject(ref),
	}
	p := &plan.Plan{Desired: desired, Planned: desired}

	t.Run("forwards the apply error", func(t *testing.T) {
		reporter := &fakeStatusReporter{}
		applyErr := errors.New("provider unavailable")

		reportSyncStatus(t.Context(), reporter, p, applyErr)

		assert.Equal(t, 1, reporter.calls)
		assert.Equal(t, applyErr, reporter.applyErr)
	})

	t.Run("nil reporter is a no-op", func(t *testing.T) {
		assert.NotPanics(t, func() {
			reportSyncStatus(t.Context(), nil, p, nil)
		})
	})

	t.Run("endpoints without ref objects skip the reporter", func(t *testing.T) {
		reporter := &fakeStatusReporter{}
		refless := &plan.Plan{
			Desired: []*endpoint.Endpoint{endpoint.NewEndpoint("d.example.com", endpoint.RecordTypeA, "10.0.0.4")},
		}

		reportSyncStatus(t.Context(), reporter, refless, nil)

		assert.Equal(t, 0, reporter.calls)
	})
}
