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

package source

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/interceptor"

	apiv1alpha1 "sigs.k8s.io/external-dns/apis/v1alpha1"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/events"
	eventsfake "sigs.k8s.io/external-dns/pkg/events/fake"
	"sigs.k8s.io/external-dns/source/types"
)

var _ StatusReporter = &crdSource{}

// Every DNSEndpoint in this file uses the same name and namespace.
const (
	testDNSEndpointNamespace = "foo"
	testDNSEndpointName      = "test"
)

// readDNSEndpoint fetches the stored copy of obj, so assertions run against what
// was actually written to the API rather than the in-memory object.
func readDNSEndpoint(t *testing.T, c client.Client) *apiv1alpha1.DNSEndpoint {
	t.Helper()
	got := &apiv1alpha1.DNSEndpoint{}
	require.NoError(t, c.Get(t.Context(), client.ObjectKey{Namespace: testDNSEndpointNamespace, Name: testDNSEndpointName}, got))
	return got
}

func TestCRDSourceAcceptedCondition(t *testing.T) {
	for _, ti := range []struct {
		title           string
		endpoints       []*endpoint.Endpoint
		wantStatus      metav1.ConditionStatus
		wantReason      string
		wantEndpoints   int32
		wantMessagePart string
	}{
		{
			title: "all endpoints valid",
			endpoints: []*endpoint.Endpoint{
				{DNSName: "example.org", Targets: endpoint.Targets{"1.2.3.4"}, RecordType: endpoint.RecordTypeA},
				{DNSName: "www.example.org", Targets: endpoint.Targets{"example.org"}, RecordType: endpoint.RecordTypeCNAME},
			},
			wantStatus:      metav1.ConditionTrue,
			wantReason:      apiv1alpha1.AcceptedReason,
			wantEndpoints:   2,
			wantMessagePart: "2 endpoint(s) accepted",
		},
		{
			title: "one endpoint rejected reports the offending index and a fix",
			endpoints: []*endpoint.Endpoint{
				{DNSName: "example.org", Targets: endpoint.Targets{"1.2.3.4"}, RecordType: endpoint.RecordTypeA},
				{DNSName: "bad.example.org", Targets: endpoint.Targets{"1.2.3.4."}, RecordType: endpoint.RecordTypeA},
			},
			wantStatus:      metav1.ConditionFalse,
			wantReason:      apiv1alpha1.InvalidReason,
			wantEndpoints:   1,
			wantMessagePart: `spec.endpoints[1] (A bad.example.org): target "1.2.3.4." must not end with a dot`,
		},
		{
			title: "null entry in spec is reported rather than silently skipped",
			endpoints: []*endpoint.Endpoint{
				nil,
				{DNSName: "example.org", Targets: endpoint.Targets{"1.2.3.4"}, RecordType: endpoint.RecordTypeA},
			},
			wantStatus:      metav1.ConditionFalse,
			wantReason:      apiv1alpha1.InvalidReason,
			wantEndpoints:   1,
			wantMessagePart: "spec.endpoints[0]: entry is null",
		},
		{
			title:           "empty spec is accepted with zero endpoints",
			endpoints:       nil,
			wantStatus:      metav1.ConditionTrue,
			wantReason:      apiv1alpha1.AcceptedReason,
			wantEndpoints:   0,
			wantMessagePart: "0 endpoint(s) accepted",
		},
	} {
		t.Run(ti.title, func(t *testing.T) {
			obj := &apiv1alpha1.DNSEndpoint{
				ObjectMeta: metav1.ObjectMeta{Name: testDNSEndpointName, Namespace: testDNSEndpointNamespace, Generation: 3},
				Spec:       apiv1alpha1.DNSEndpointSpec{Endpoints: ti.endpoints},
			}

			fakeCache := newFakeCRDCache(t, nil, fakeCRDCacheFilter{}, obj)
			cs, err := newCrdSource(t.Context(), fakeCache, fakeCache.Client, "", nil, nil)
			require.NoError(t, err)

			_, err = cs.Endpoints(t.Context())
			require.NoError(t, err)

			got := readDNSEndpoint(t, fakeCache.Client)
			assert.Equal(t, int64(3), got.Status.ObservedGeneration)
			assert.Equal(t, ti.wantEndpoints, got.Status.Endpoints)

			cond := meta.FindStatusCondition(got.Status.Conditions, apiv1alpha1.AcceptedCondition)
			require.NotNil(t, cond, "Accepted condition must be set")
			assert.Equal(t, ti.wantStatus, cond.Status)
			assert.Equal(t, ti.wantReason, cond.Reason)
			assert.Equal(t, int64(3), cond.ObservedGeneration)
			assert.Contains(t, cond.Message, ti.wantMessagePart)
		})
	}
}

// A DNSEndpoint that has not changed must not cost an API write on every
// reconcile — external-dns re-lists on a timer, so a write per pass would be a
// write per minute per object.
func TestCRDSourceStatusIsNotRewrittenWhenUnchanged(t *testing.T) {
	obj := &apiv1alpha1.DNSEndpoint{
		ObjectMeta: metav1.ObjectMeta{Name: testDNSEndpointName, Namespace: testDNSEndpointNamespace, Generation: 1},
		Spec: apiv1alpha1.DNSEndpointSpec{Endpoints: []*endpoint.Endpoint{
			{DNSName: "example.org", Targets: endpoint.Targets{"1.2.3.4"}, RecordType: endpoint.RecordTypeA},
		}},
	}

	fakeCache := newFakeCRDCache(t, nil, fakeCRDCacheFilter{}, obj)

	var writes int
	countingWriter := interceptor.NewClient(fakeCache.Client.(client.WithWatch), interceptor.Funcs{
		SubResourceUpdate: func(
			ctx context.Context,
			c client.Client,
			subResource string,
			o client.Object,
			opts ...client.SubResourceUpdateOption) error {
			if subResource == "status" {
				writes++
			}
			return c.Status().Update(ctx, o, opts...)
		},
	})

	cs, err := newCrdSource(t.Context(), fakeCache, countingWriter, "", nil, nil)
	require.NoError(t, err)

	_, err = cs.Endpoints(t.Context())
	require.NoError(t, err)
	require.Equal(t, 1, writes, "first reconcile must write the initial status")

	for range 3 {
		_, err = cs.Endpoints(t.Context())
		require.NoError(t, err)
	}
	assert.Equal(t, 1, writes, "unchanged status must not be rewritten")
}

func TestCRDSourceEmitsEventOnRejectedEndpoint(t *testing.T) {
	obj := &apiv1alpha1.DNSEndpoint{
		ObjectMeta: metav1.ObjectMeta{Name: testDNSEndpointName, Namespace: testDNSEndpointNamespace, Generation: 1},
		Spec: apiv1alpha1.DNSEndpointSpec{Endpoints: []*endpoint.Endpoint{
			{DNSName: "bad.example.org", Targets: endpoint.Targets{"1.2.3.4."}, RecordType: endpoint.RecordTypeA},
		}},
	}

	fakeCache := newFakeCRDCache(t, nil, fakeCRDCacheFilter{}, obj)
	emitter := eventsfake.NewFakeEventEmitter()

	cs, err := newCrdSource(t.Context(), fakeCache, fakeCache.Client, "", nil, emitter)
	require.NoError(t, err)

	_, err = cs.Endpoints(t.Context())
	require.NoError(t, err)

	emitter.AssertNumberOfCalls(t, "Add", 1)
	emitted, ok := emitter.Calls[0].Arguments.Get(0).(events.Event)
	require.True(t, ok)
	assert.Equal(t, events.RecordInvalid, emitted.Reason())
	assert.Equal(t, events.ActionRejected, emitted.Action())
	assert.Equal(t, events.EventTypeWarning, emitted.EventType())
}

func TestCRDSourceEmitsNoEventWhenAllEndpointsValid(t *testing.T) {
	obj := &apiv1alpha1.DNSEndpoint{
		ObjectMeta: metav1.ObjectMeta{Name: testDNSEndpointName, Namespace: testDNSEndpointNamespace, Generation: 1},
		Spec: apiv1alpha1.DNSEndpointSpec{Endpoints: []*endpoint.Endpoint{
			{DNSName: "example.org", Targets: endpoint.Targets{"1.2.3.4"}, RecordType: endpoint.RecordTypeA},
		}},
	}

	fakeCache := newFakeCRDCache(t, nil, fakeCRDCacheFilter{}, obj)
	emitter := eventsfake.NewFakeEventEmitter()

	cs, err := newCrdSource(t.Context(), fakeCache, fakeCache.Client, "", nil, emitter)
	require.NoError(t, err)

	_, err = cs.Endpoints(t.Context())
	require.NoError(t, err)

	emitter.AssertNumberOfCalls(t, "Add", 0)
}

func TestCRDSourceReportStatus(t *testing.T) {
	for _, ti := range []struct {
		title           string
		applyErr        error
		wantStatus      metav1.ConditionStatus
		wantReason      string
		wantMessagePart string
	}{
		{
			title:           "provider applied the batch",
			applyErr:        nil,
			wantStatus:      metav1.ConditionTrue,
			wantReason:      apiv1alpha1.ProgrammedReason,
			wantMessagePart: "applied to the DNS provider",
		},
		{
			title:           "provider rejected the batch",
			applyErr:        errors.New("route53: throttled"),
			wantStatus:      metav1.ConditionFalse,
			wantReason:      apiv1alpha1.FailedReason,
			wantMessagePart: "route53: throttled",
		},
	} {
		t.Run(ti.title, func(t *testing.T) {
			obj := &apiv1alpha1.DNSEndpoint{
				ObjectMeta: metav1.ObjectMeta{Name: testDNSEndpointName, Namespace: testDNSEndpointNamespace, Generation: 7},
				Spec: apiv1alpha1.DNSEndpointSpec{Endpoints: []*endpoint.Endpoint{
					{DNSName: "example.org", Targets: endpoint.Targets{"1.2.3.4"}, RecordType: endpoint.RecordTypeA},
				}},
			}

			fakeCache := newFakeCRDCache(t, nil, fakeCRDCacheFilter{}, obj)
			cs, err := newCrdSource(t.Context(), fakeCache, fakeCache.Client, "", nil, nil)
			require.NoError(t, err)

			ref := events.NewObjectReference(obj, types.CRD)
			cs.ReportStatus(t.Context(), []*events.ObjectReference{ref}, ti.applyErr)

			got := readDNSEndpoint(t, fakeCache.Client)
			cond := meta.FindStatusCondition(got.Status.Conditions, apiv1alpha1.ReadyCondition)
			require.NotNil(t, cond, "Ready condition must be set")
			assert.Equal(t, ti.wantStatus, cond.Status)
			assert.Equal(t, ti.wantReason, cond.Reason)
			assert.Equal(t, int64(7), cond.ObservedGeneration)
			assert.Contains(t, cond.Message, ti.wantMessagePart)
		})
	}
}

// The controller hands every source the full ref set, so a crd source must not
// touch references produced elsewhere, nor blow up on a deleted object.
func TestCRDSourceReportStatusIgnoresForeignAndMissingRefs(t *testing.T) {
	obj := &apiv1alpha1.DNSEndpoint{
		ObjectMeta: metav1.ObjectMeta{Name: testDNSEndpointName, Namespace: testDNSEndpointNamespace, Generation: 1},
	}

	fakeCache := newFakeCRDCache(t, nil, fakeCRDCacheFilter{}, obj)
	cs, err := newCrdSource(t.Context(), fakeCache, fakeCache.Client, "", nil, nil)
	require.NoError(t, err)

	foreign := events.NewObjectReferenceFromParts("Ingress", "networking.k8s.io/v1", "foo", "test", "", types.Ingress)
	deleted := events.NewObjectReferenceFromParts("DNSEndpoint", "externaldns.k8s.io/v1alpha1", "foo", "gone", "", types.CRD)

	cs.ReportStatus(t.Context(), []*events.ObjectReference{nil, foreign, deleted}, nil)

	got := readDNSEndpoint(t, fakeCache.Client)
	assert.Empty(t, got.Status.Conditions, "no condition must be written for foreign or missing refs")
}

// Accepted is written from the listed (cache-backed) copy while Ready is written
// after the apply. The read path can lag the last write, so updateStatus must
// re-read before writing: mutating and pushing the stale copy would drop the
// condition the other writer had just set.
func TestCRDSourceStatusUpdateDoesNotClobberAStaleCondition(t *testing.T) {
	obj := &apiv1alpha1.DNSEndpoint{
		ObjectMeta: metav1.ObjectMeta{Name: testDNSEndpointName, Namespace: testDNSEndpointNamespace, Generation: 1},
		Spec: apiv1alpha1.DNSEndpointSpec{Endpoints: []*endpoint.Endpoint{
			{DNSName: "example.org", Targets: endpoint.Targets{"1.2.3.4"}, RecordType: endpoint.RecordTypeA},
		}},
	}

	fakeCache := newFakeCRDCache(t, nil, fakeCRDCacheFilter{}, obj)
	cs, err := newCrdSource(t.Context(), fakeCache, fakeCache.Client, "", nil, nil)
	require.NoError(t, err)

	_, err = cs.Endpoints(t.Context())
	require.NoError(t, err)

	// Snapshot the object as a lagging cache would hand it back: Accepted is set,
	// Ready is not.
	stale := readDNSEndpoint(t, fakeCache.Client)
	require.Nil(t, meta.FindStatusCondition(stale.Status.Conditions, apiv1alpha1.ReadyCondition))

	cs.ReportStatus(t.Context(), []*events.ObjectReference{events.NewObjectReference(obj, types.CRD)}, nil)

	// Now write Accepted again from the stale copy, changing it so a write happens.
	cs.updateStatus(t.Context(), stale, func(status *apiv1alpha1.DNSEndpointStatus) {
		meta.SetStatusCondition(&status.Conditions, metav1.Condition{
			Type:    apiv1alpha1.AcceptedCondition,
			Status:  metav1.ConditionFalse,
			Reason:  apiv1alpha1.InvalidReason,
			Message: "spec.endpoints[0]: something changed",
		})
	})

	got := readDNSEndpoint(t, fakeCache.Client)
	accepted := meta.FindStatusCondition(got.Status.Conditions, apiv1alpha1.AcceptedCondition)
	require.NotNil(t, accepted)
	assert.Equal(t, apiv1alpha1.InvalidReason, accepted.Reason, "the new Accepted must be stored")

	ready := meta.FindStatusCondition(got.Status.Conditions, apiv1alpha1.ReadyCondition)
	require.NotNil(t, ready, "Ready must survive a write driven from a stale copy")
	assert.Equal(t, apiv1alpha1.ProgrammedReason, ready.Reason)
}

func TestTruncateConditionMessage(t *testing.T) {
	short := "all good"
	assert.Equal(t, short, truncateConditionMessage(short))

	long := truncateConditionMessage(fmt.Sprintf("%0*d", 40000, 0))
	assert.Len(t, long, 32768)
	assert.True(t, len(long) > 3 && long[len(long)-3:] == "...")
}
