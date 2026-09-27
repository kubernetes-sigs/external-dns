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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/interceptor"

	apiv1alpha1 "sigs.k8s.io/external-dns/apis/v1alpha1"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/crd"
	"sigs.k8s.io/external-dns/pkg/events"
	eventsfake "sigs.k8s.io/external-dns/pkg/events/fake"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/source/types"
)

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
		defaultTargets  bool
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
			title: "endpoint without targets is reported when --default-targets is unset",
			endpoints: []*endpoint.Endpoint{
				{DNSName: "example.org", RecordType: endpoint.RecordTypeA},
			},
			wantStatus:      metav1.ConditionFalse,
			wantReason:      apiv1alpha1.InvalidReason,
			wantMessagePart: "spec.endpoints[0] (A example.org): no targets",
		},
		{
			title: "endpoint without targets is accepted when --default-targets fills it",
			endpoints: []*endpoint.Endpoint{
				{DNSName: "example.org", RecordType: endpoint.RecordTypeA},
			},
			defaultTargets:  true,
			wantStatus:      metav1.ConditionTrue,
			wantReason:      apiv1alpha1.AcceptedReason,
			wantMessagePart: "1 endpoint(s) accepted",
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
				Name: testDNSEndpointName, Namespace: testDNSEndpointNamespace, Generation: 3,
				Spec: apiv1alpha1.DNSEndpointSpec{Endpoints: ti.endpoints},
			}

			fakeCache := newFakeCRDCache(t, nil, obj)
			cs, err := newCrdSource(t.Context(), fakeCache, fakeCache.Client, "", nil, nil, nil)
			require.NoError(t, err)
			cs.defaultTargets = ti.defaultTargets

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
		Name: testDNSEndpointName, Namespace: testDNSEndpointNamespace, Generation: 1,
		Spec: apiv1alpha1.DNSEndpointSpec{Endpoints: []*endpoint.Endpoint{
			{DNSName: "example.org", Targets: endpoint.Targets{"1.2.3.4"}, RecordType: endpoint.RecordTypeA},
		}},
	}

	fakeCache := newFakeCRDCache(t, nil, obj)

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

	cs, err := newCrdSource(t.Context(), fakeCache, countingWriter, "", nil, nil, nil)
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
		Name: testDNSEndpointName, Namespace: testDNSEndpointNamespace, Generation: 1,
		Spec: apiv1alpha1.DNSEndpointSpec{Endpoints: []*endpoint.Endpoint{
			{DNSName: "example.org", Targets: endpoint.Targets{"1.2.3.4"}, RecordType: endpoint.RecordTypeA},
		}},
	}

	fakeCache := newFakeCRDCache(t, nil, obj)

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

	cs, err := newCrdSource(t.Context(), fakeCache, counting, "", nil, nil, nil)
	require.NoError(t, err)

	_, err = cs.Endpoints(t.Context())
	require.NoError(t, err)

	assert.Equal(t, 1, writes)
	assert.Equal(t, 0, gets, "an uncontended status write costs one round trip")
}

func TestCRDSourceEmitsEventOnRejectedEndpoint(t *testing.T) {
	obj := &apiv1alpha1.DNSEndpoint{
		Name: testDNSEndpointName, Namespace: testDNSEndpointNamespace, Generation: 1,
		Spec: apiv1alpha1.DNSEndpointSpec{Endpoints: []*endpoint.Endpoint{
			{DNSName: "bad.example.org", Targets: endpoint.Targets{"1.2.3.4."}, RecordType: endpoint.RecordTypeA},
		}},
	}

	fakeCache := newFakeCRDCache(t, nil, obj)
	emitter := eventsfake.NewFakeEventEmitter()

	cs, err := newCrdSource(t.Context(), fakeCache, fakeCache.Client, "", nil, nil, emitter)
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
		Name: testDNSEndpointName, Namespace: testDNSEndpointNamespace, Generation: 1,
		Spec: apiv1alpha1.DNSEndpointSpec{Endpoints: []*endpoint.Endpoint{
			{DNSName: "bad.example.org", Targets: endpoint.Targets{"1.2.3.4."}, RecordType: endpoint.RecordTypeA},
		}},
	}

	fakeCache := newFakeCRDCache(t, nil, obj)
	emitter := eventsfake.NewFakeEventEmitter()

	cs, err := newCrdSource(t.Context(), fakeCache, fakeCache.Client, "", nil, nil, emitter)
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
// the status writer never sees it and the previous verdict would linger.
func TestCRDSourceClearsReadyWhenEveryEndpointIsRejected(t *testing.T) {
	obj := &apiv1alpha1.DNSEndpoint{
		Name: testDNSEndpointName, Namespace: testDNSEndpointNamespace, Generation: 1,
		Spec: apiv1alpha1.DNSEndpointSpec{Endpoints: []*endpoint.Endpoint{
			{DNSName: "example.org", Targets: endpoint.Targets{"1.2.3.4"}, RecordType: endpoint.RecordTypeA},
		}},
	}

	fakeCache := newFakeCRDCache(t, nil, obj)
	cs, err := newCrdSource(t.Context(), fakeCache, fakeCache.Client, "", nil, nil, nil)
	require.NoError(t, err)

	_, err = cs.Endpoints(t.Context())
	require.NoError(t, err)
	reportProgrammed(t, fakeCache, obj)

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
		Name: testDNSEndpointName, Namespace: testDNSEndpointNamespace, Generation: 1,
		Spec: apiv1alpha1.DNSEndpointSpec{Endpoints: []*endpoint.Endpoint{
			{DNSName: "example.org", Targets: endpoint.Targets{"1.2.3.4"}, RecordType: endpoint.RecordTypeA},
		}},
	}

	fakeCache := newFakeCRDCache(t, nil, obj)
	cs, err := newCrdSource(t.Context(), fakeCache, fakeCache.Client, "", nil, nil, nil)
	require.NoError(t, err)

	_, err = cs.Endpoints(t.Context())
	require.NoError(t, err)
	reportProgrammed(t, fakeCache, obj)

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
		Name: testDNSEndpointName, Namespace: testDNSEndpointNamespace, Generation: 1,
		Spec: apiv1alpha1.DNSEndpointSpec{Endpoints: []*endpoint.Endpoint{
			{DNSName: "example.org", Targets: endpoint.Targets{"1.2.3.4"}, RecordType: endpoint.RecordTypeA},
		}},
	}

	fakeCache := newFakeCRDCache(t, nil, obj)
	emitter := eventsfake.NewFakeEventEmitter()

	cs, err := newCrdSource(t.Context(), fakeCache, fakeCache.Client, "", nil, nil, emitter)
	require.NoError(t, err)

	_, err = cs.Endpoints(t.Context())
	require.NoError(t, err)

	emitter.AssertNumberOfCalls(t, "Add", 0)
}

// reportProgrammed stands in for the controller's status writer after a
// successful sync of obj's endpoints.
func reportProgrammed(t *testing.T, fakeCache *fakeCRDCache, obj *apiv1alpha1.DNSEndpoint) {
	t.Helper()
	w := crd.NewStatusWriter(crd.NewCRDClients(fakeCache, fakeCache.Client), false)
	w.ReportStatus(t.Context(), []plan.PlannedObject{
		{Ref: events.NewObjectReference(obj, types.CRD), Endpoints: 1},
	}, nil)
}
