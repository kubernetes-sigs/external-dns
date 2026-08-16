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
	"strings"
	"testing"
	"unicode/utf8"

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
			wantMessagePart: `spec.endpoints[1] (A bad.example.org): target "1.2.3.4." must not end with a dot`,
		},
		{
			// dedupSource drops what CheckEndpoint rejects with only a log line,
			// so the source has to catch the grammar for anyone to hear about it.
			title: "SRV target with a relative host is reported",
			endpoints: []*endpoint.Endpoint{
				{DNSName: "_svc._tcp.example.org", Targets: endpoint.Targets{"0 0 80 abc.example.org"}, RecordType: endpoint.RecordTypeSRV},
			},
			wantStatus:      metav1.ConditionFalse,
			wantReason:      apiv1alpha1.InvalidReason,
			wantMessagePart: `spec.endpoints[0] (SRV _svc._tcp.example.org): SRV targets must be "<priority> <weight> <port> <host>"`,
		},
		{
			title: "MX target without a preference is reported",
			endpoints: []*endpoint.Endpoint{
				{DNSName: "example.org", Targets: endpoint.Targets{"example.com."}, RecordType: endpoint.RecordTypeMX},
			},
			wantStatus:      metav1.ConditionFalse,
			wantReason:      apiv1alpha1.InvalidReason,
			wantMessagePart: `MX targets must be "<preference> <host>"`,
		},
		{
			title: "AAAA record carrying an IPv4 target is reported",
			endpoints: []*endpoint.Endpoint{
				{DNSName: "example.org", Targets: endpoint.Targets{"1.2.3.4"}, RecordType: endpoint.RecordTypeAAAA},
			},
			wantStatus:      metav1.ConditionFalse,
			wantReason:      apiv1alpha1.InvalidReason,
			wantMessagePart: "must be IPv6 addresses",
		},
		{
			title: "null entry in spec is reported rather than silently skipped",
			endpoints: []*endpoint.Endpoint{
				nil,
				{DNSName: "example.org", Targets: endpoint.Targets{"1.2.3.4"}, RecordType: endpoint.RecordTypeA},
			},
			wantStatus:      metav1.ConditionFalse,
			wantReason:      apiv1alpha1.InvalidReason,
			wantMessagePart: "spec.endpoints[0]: entry is null",
		},
		{
			title:           "empty spec is accepted with zero endpoints",
			endpoints:       nil,
			wantStatus:      metav1.ConditionTrue,
			wantReason:      apiv1alpha1.AcceptedReason,
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

// The status write goes out on the copy the source already holds. Re-reading
// first would double the round trips of the reconcile after an upgrade, where
// every DNSEndpoint's status changes at once.
func TestCRDSourceStatusWriteDoesNotReReadWithoutAConflict(t *testing.T) {
	obj := &apiv1alpha1.DNSEndpoint{
		ObjectMeta: metav1.ObjectMeta{Name: testDNSEndpointName, Namespace: testDNSEndpointNamespace, Generation: 1},
		Spec: apiv1alpha1.DNSEndpointSpec{Endpoints: []*endpoint.Endpoint{
			{DNSName: "example.org", Targets: endpoint.Targets{"1.2.3.4"}, RecordType: endpoint.RecordTypeA},
		}},
	}

	fakeCache := newFakeCRDCache(t, nil, fakeCRDCacheFilter{}, obj)

	var gets, writes int
	counting := interceptor.NewClient(fakeCache.Client.(client.WithWatch), interceptor.Funcs{
		Get: func(ctx context.Context, c client.WithWatch, key client.ObjectKey, o client.Object, opts ...client.GetOption) error {
			gets++
			return c.Get(ctx, key, o, opts...)
		},
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

	cs, err := newCrdSource(t.Context(), fakeCache, counting, "", nil, nil)
	require.NoError(t, err)

	_, err = cs.Endpoints(t.Context())
	require.NoError(t, err)

	assert.Equal(t, 1, writes)
	assert.Equal(t, 0, gets, "an uncontended status write costs one round trip")
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

// Events carry a timestamped name, so nothing collapses them into a series:
// emitting per reconcile would create one Event a minute for an untouched spec.
func TestCRDSourceEmitsRejectionEventOnlyWhenTheVerdictChanges(t *testing.T) {
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

	for range 4 {
		_, err = cs.Endpoints(t.Context())
		require.NoError(t, err)
	}

	emitter.AssertNumberOfCalls(t, "Add", 1)

	// A different rejection is news again.
	stored := readDNSEndpoint(t, fakeCache.Client)
	stored.Spec.Endpoints[0].Targets = endpoint.Targets{"5.6.7.8."}
	require.NoError(t, fakeCache.Client.Update(t.Context(), stored))

	_, err = cs.Endpoints(t.Context())
	require.NoError(t, err)

	emitter.AssertNumberOfCalls(t, "Add", 2)
}

// An object whose endpoints were all rejected contributes nothing to the plan, so
// ReportStatus never sees it and the previous verdict would linger.
func TestCRDSourceClearsReadyWhenEveryEndpointIsRejected(t *testing.T) {
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
	cs.ReportStatus(t.Context(), []PlannedObject{
		{Ref: events.NewObjectReference(obj, types.CRD), Endpoints: 1},
	}, nil)

	programmed := readDNSEndpoint(t, fakeCache.Client)
	require.Equal(t, int32(1), programmed.Status.Endpoints)
	require.Equal(t, apiv1alpha1.ProgrammedReason,
		meta.FindStatusCondition(programmed.Status.Conditions, apiv1alpha1.ReadyCondition).Reason)

	// The user breaks the only endpoint.
	programmed.Spec.Endpoints[0].Targets = endpoint.Targets{"1.2.3.4."}
	require.NoError(t, fakeCache.Client.Update(t.Context(), programmed))

	_, err = cs.Endpoints(t.Context())
	require.NoError(t, err)

	got := readDNSEndpoint(t, fakeCache.Client)
	assert.Equal(t, apiv1alpha1.InvalidReason,
		meta.FindStatusCondition(got.Status.Conditions, apiv1alpha1.AcceptedCondition).Reason)

	ready := meta.FindStatusCondition(got.Status.Conditions, apiv1alpha1.ReadyCondition)
	require.NotNil(t, ready)
	assert.Equal(t, metav1.ConditionFalse, ready.Status)
	assert.Equal(t, apiv1alpha1.InvalidReason, ready.Reason, "Ready must not still claim Programmed")
	assert.Zero(t, got.Status.Endpoints, "the endpoint count must not survive the rejection")
}

// An emptied spec has nothing to be ready about, so the condition goes away.
func TestCRDSourceRemovesReadyWhenSpecBecomesEmpty(t *testing.T) {
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
	cs.ReportStatus(t.Context(), []PlannedObject{
		{Ref: events.NewObjectReference(obj, types.CRD), Endpoints: 1},
	}, nil)

	emptied := readDNSEndpoint(t, fakeCache.Client)
	emptied.Spec.Endpoints = nil
	require.NoError(t, fakeCache.Client.Update(t.Context(), emptied))

	_, err = cs.Endpoints(t.Context())
	require.NoError(t, err)

	got := readDNSEndpoint(t, fakeCache.Client)
	assert.Nil(t, meta.FindStatusCondition(got.Status.Conditions, apiv1alpha1.ReadyCondition))
	assert.Zero(t, got.Status.Endpoints)
	assert.Equal(t, apiv1alpha1.AcceptedReason,
		meta.FindStatusCondition(got.Status.Conditions, apiv1alpha1.AcceptedCondition).Reason)
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
		planned         int
		applyErr        error
		wantStatus      metav1.ConditionStatus
		wantReason      string
		wantMessagePart string
	}{
		{
			title:           "provider applied the batch",
			planned:         1,
			applyErr:        nil,
			wantStatus:      metav1.ConditionTrue,
			wantReason:      apiv1alpha1.ProgrammedReason,
			wantMessagePart: "1 endpoint(s) applied to the DNS provider",
		},
		{
			title:           "provider rejected the batch",
			planned:         1,
			applyErr:        errors.New("route53: throttled"),
			wantStatus:      metav1.ConditionFalse,
			wantReason:      apiv1alpha1.FailedReason,
			wantMessagePart: "route53: throttled",
		},
		{
			// Nothing was offered to the provider, so Programmed would be a lie.
			title:           "every endpoint excluded by the filters",
			planned:         0,
			applyErr:        nil,
			wantStatus:      metav1.ConditionFalse,
			wantReason:      apiv1alpha1.FilteredReason,
			wantMessagePart: "No endpoint reached the DNS provider",
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

			planned := PlannedObject{Ref: events.NewObjectReference(obj, types.CRD), Endpoints: ti.planned}
			cs.ReportStatus(t.Context(), []PlannedObject{planned}, ti.applyErr)

			got := readDNSEndpoint(t, fakeCache.Client)
			assert.Equal(t, int32(ti.planned), got.Status.Endpoints) // #nosec G115 -- test data

			cond := meta.FindStatusCondition(got.Status.Conditions, apiv1alpha1.ReadyCondition)
			require.NotNil(t, cond, "Ready condition must be set")
			assert.Equal(t, ti.wantStatus, cond.Status)
			assert.Equal(t, ti.wantReason, cond.Reason)
			assert.Equal(t, int64(7), cond.ObservedGeneration)
			assert.Contains(t, cond.Message, ti.wantMessagePart)
		})
	}
}

// The controller hands every source the full object set, so a crd source must
// not touch references produced elsewhere, nor blow up on a deleted object.
func TestCRDSourceReportStatusIgnoresForeignAndMissingRefs(t *testing.T) {
	obj := &apiv1alpha1.DNSEndpoint{
		ObjectMeta: metav1.ObjectMeta{Name: testDNSEndpointName, Namespace: testDNSEndpointNamespace, Generation: 1},
	}

	fakeCache := newFakeCRDCache(t, nil, fakeCRDCacheFilter{}, obj)
	cs, err := newCrdSource(t.Context(), fakeCache, fakeCache.Client, "", nil, nil)
	require.NoError(t, err)

	foreign := events.NewObjectReferenceFromParts("Ingress", "networking.k8s.io/v1", "foo", "test", "", types.Ingress)
	deleted := events.NewObjectReferenceFromParts("DNSEndpoint", "externaldns.k8s.io/v1alpha1", "foo", "gone", "", types.CRD)

	cs.ReportStatus(t.Context(), []PlannedObject{
		{Ref: nil, Endpoints: 1},
		{Ref: foreign, Endpoints: 1},
		{Ref: deleted, Endpoints: 1},
	}, nil)

	got := readDNSEndpoint(t, fakeCache.Client)
	assert.Empty(t, got.Status.Conditions, "no condition must be written for foreign or missing refs")
}

// Accepted is written from the listed (cache-backed) copy while Ready is written
// after the apply. The read path can lag the last write, so pushing the stale
// copy would drop the condition the other writer had just set. Its stale
// resourceVersion makes the API server reject it, and updateStatus then re-reads.
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

	cs.ReportStatus(t.Context(), []PlannedObject{
		{Ref: events.NewObjectReference(obj, types.CRD), Endpoints: 1},
	}, nil)

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

	// User-supplied names may be UTF-8; cutting bytes would split a rune.
	multibyte := truncateConditionMessage(strings.Repeat("é", 40000))
	assert.Equal(t, 32768, utf8.RuneCountInString(multibyte))
	assert.True(t, utf8.ValidString(multibyte), "truncation must not split a rune")
}
