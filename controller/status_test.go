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

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/events"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/source"
)

// fakeStatusReporter records what reportSyncStatus handed it.
type fakeStatusReporter struct {
	calls    int
	refs     []*events.ObjectReference
	applyErr error
}

func (f *fakeStatusReporter) ReportStatus(_ context.Context, refs []*events.ObjectReference, applyErr error) {
	f.calls++
	f.refs = refs
	f.applyErr = applyErr
}

func TestReportSyncStatus(t *testing.T) {
	// Two endpoints from the same object plus one from another: a reporter must
	// see each object once, not once per endpoint it produced.
	first := events.NewObjectReferenceFromParts("DNSEndpoint", "externaldns.k8s.io/v1alpha1", "ns", "first", "", "crd")
	second := events.NewObjectReferenceFromParts("DNSEndpoint", "externaldns.k8s.io/v1alpha1", "ns", "second", "", "crd")

	changes := &plan.Changes{
		Create: []*endpoint.Endpoint{
			endpoint.NewEndpoint("a.example.com", endpoint.RecordTypeA, "10.0.0.1").WithRefObject(first),
			endpoint.NewEndpoint("b.example.com", endpoint.RecordTypeA, "10.0.0.2").WithRefObject(first),
		},
		UpdateNew: []*endpoint.Endpoint{
			endpoint.NewEndpoint("c.example.com", endpoint.RecordTypeA, "10.0.0.3").WithRefObject(second),
		},
	}

	t.Run("dedupes refs and forwards the apply error", func(t *testing.T) {
		reporter := &fakeStatusReporter{}
		applyErr := errors.New("provider unavailable")

		reportSyncStatus(t.Context(), []source.StatusReporter{reporter}, changes, applyErr)

		assert.Equal(t, 1, reporter.calls)
		assert.Len(t, reporter.refs, 2)
		assert.Equal(t, applyErr, reporter.applyErr)
	})

	t.Run("every reporter is called", func(t *testing.T) {
		first, second := &fakeStatusReporter{}, &fakeStatusReporter{}

		reportSyncStatus(t.Context(), []source.StatusReporter{first, second}, changes, nil)

		assert.Equal(t, 1, first.calls)
		assert.Equal(t, 1, second.calls)
	})

	t.Run("no reporters is a no-op", func(t *testing.T) {
		assert.NotPanics(t, func() {
			reportSyncStatus(t.Context(), nil, changes, nil)
		})
	})

	t.Run("changes without ref objects skip the reporters", func(t *testing.T) {
		reporter := &fakeStatusReporter{}
		refless := &plan.Changes{
			Delete: []*endpoint.Endpoint{endpoint.NewEndpoint("d.example.com", endpoint.RecordTypeA, "10.0.0.4")},
		}

		reportSyncStatus(t.Context(), []source.StatusReporter{reporter}, refless, nil)

		assert.Equal(t, 0, reporter.calls)
	})
}
