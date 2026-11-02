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
)

const (
	testDNSEndpointNamespace = "foo"
	testDNSEndpointName      = "test"
)

// readDNSEndpoint reads what was written to the API, not the in-memory object.
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
		forceDefaults   bool
		wantStatus      metav1.ConditionStatus
		wantReason      string
		wantMessagePart string
		// "" means PartiallyInvalid must be absent.
		wantPartial string
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
			title: "mixed spec stays accepted and lists the dropped endpoint",
			endpoints: []*endpoint.Endpoint{
				{DNSName: "example.org", Targets: endpoint.Targets{"1.2.3.4"}, RecordType: endpoint.RecordTypeA},
				{DNSName: "bad.example.org", Targets: endpoint.Targets{"1.2.3.4."}, RecordType: endpoint.RecordTypeA},
			},
			wantStatus:      metav1.ConditionTrue,
			wantReason:      apiv1alpha1.AcceptedReason,
			wantMessagePart: "1 endpoint(s) accepted",
			wantPartial:     `Dropped endpoint(s): spec.endpoints[1] (A bad.example.org): target "1.2.3.4." must not end with a dot`,
		},
		{
			title: "every endpoint rejected reports the offending index and a fix",
			endpoints: []*endpoint.Endpoint{
				{DNSName: "bad.example.org", Targets: endpoint.Targets{"1.2.3.4."}, RecordType: endpoint.RecordTypeA},
			},
			wantStatus:      metav1.ConditionFalse,
			wantReason:      apiv1alpha1.InvalidReason,
			wantMessagePart: `spec.endpoints[0] (A bad.example.org): target "1.2.3.4." must not end with a dot`,
		},
		{
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
			wantStatus:      metav1.ConditionTrue,
			wantReason:      apiv1alpha1.AcceptedReason,
			wantMessagePart: "1 endpoint(s) accepted",
			wantPartial:     "spec.endpoints[0]: entry is null",
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
			title: "CNAME without targets is reported without the --default-targets hint",
			endpoints: []*endpoint.Endpoint{
				{DNSName: "example.org", RecordType: endpoint.RecordTypeCNAME},
			},
			wantStatus:      metav1.ConditionFalse,
			wantReason:      apiv1alpha1.InvalidReason,
			wantMessagePart: "spec.endpoints[0] (CNAME example.org): no targets: a CNAME record needs one target",
		},
		{
			title: "CNAME without targets is reported even when --default-targets is set",
			endpoints: []*endpoint.Endpoint{
				{DNSName: "example.org", RecordType: endpoint.RecordTypeCNAME},
			},
			defaultTargets:  true,
			wantStatus:      metav1.ConditionFalse,
			wantReason:      apiv1alpha1.InvalidReason,
			wantMessagePart: "spec.endpoints[0] (CNAME example.org): no targets: a CNAME record needs one target",
		},
		{
			title: "targets are not checked when --force-default-targets replaces them",
			endpoints: []*endpoint.Endpoint{
				{DNSName: "example.org", Targets: endpoint.Targets{"foo.example.com"}, RecordType: endpoint.RecordTypeA},
			},
			defaultTargets:  true,
			forceDefaults:   true,
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
			cs, err := newCrdSource(t.Context(), fakeCache, fakeCache.Client, nil, nil, nil)
			require.NoError(t, err)
			cs.defaultTargets = ti.defaultTargets
			cs.forceDefaultTargets = ti.forceDefaults

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

			partial := meta.FindStatusCondition(got.Status.Conditions, apiv1alpha1.PartiallyInvalidCondition)
			if ti.wantPartial == "" {
				assert.Nil(t, partial)
				return
			}
			require.NotNil(t, partial)
			assert.Equal(t, metav1.ConditionTrue, partial.Status)
			assert.Equal(t, apiv1alpha1.InvalidReason, partial.Reason)
			assert.Contains(t, partial.Message, ti.wantPartial)
		})
	}
}

// An unchanged DNSEndpoint must not cost one API write per sync.
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

	cs, err := newCrdSource(t.Context(), fakeCache, countingWriter, nil, nil, nil)
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

// Re-reading before each write would double the API calls of the first sync after
// an upgrade, when every DNSEndpoint's status changes.
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

	cs, err := newCrdSource(t.Context(), fakeCache, counting, nil, nil, nil)
	require.NoError(t, err)

	_, err = cs.Endpoints(t.Context())
	require.NoError(t, err)

	assert.Equal(t, 1, writes)
	assert.Equal(t, 0, gets, "an uncontended status write costs one round trip")
}

func TestCRDSourcePartiallyInvalidIsRemovedOnceFixed(t *testing.T) {
	obj := &apiv1alpha1.DNSEndpoint{
		Name: testDNSEndpointName, Namespace: testDNSEndpointNamespace, Generation: 1,
		Spec: apiv1alpha1.DNSEndpointSpec{Endpoints: []*endpoint.Endpoint{
			{DNSName: "example.org", Targets: endpoint.Targets{"1.2.3.4"}, RecordType: endpoint.RecordTypeA},
			{DNSName: "bad.example.org", Targets: endpoint.Targets{"1.2.3.4."}, RecordType: endpoint.RecordTypeA},
		}},
	}

	fakeCache := newFakeCRDCache(t, nil, obj)
	cs, err := newCrdSource(t.Context(), fakeCache, fakeCache.Client, nil, nil, nil)
	require.NoError(t, err)

	_, err = cs.Endpoints(t.Context())
	require.NoError(t, err)
	stored := readDNSEndpoint(t, fakeCache.Client)
	require.NotNil(t, meta.FindStatusCondition(stored.Status.Conditions, apiv1alpha1.PartiallyInvalidCondition))

	stored.Spec.Endpoints[1].Targets = endpoint.Targets{"1.2.3.5"}
	require.NoError(t, fakeCache.Client.Update(t.Context(), stored))

	_, err = cs.Endpoints(t.Context())
	require.NoError(t, err)
	got := readDNSEndpoint(t, fakeCache.Client)
	assert.Nil(t, meta.FindStatusCondition(got.Status.Conditions, apiv1alpha1.PartiallyInvalidCondition))
	assert.True(t, meta.IsStatusConditionTrue(got.Status.Conditions, apiv1alpha1.AcceptedCondition))
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

// An outdated CRD prunes the conditions, so every sync would write every object
// again: fail instead, until the CRD is applied.
func TestCRDSourceFailsWhenCRDPrunesConditions(t *testing.T) {
	obj := &apiv1alpha1.DNSEndpoint{
		Name: testDNSEndpointName, Namespace: testDNSEndpointNamespace, Generation: 1,
		Spec: apiv1alpha1.DNSEndpointSpec{Endpoints: []*endpoint.Endpoint{
			{DNSName: "example.org", Targets: endpoint.Targets{"1.2.3.4"}, RecordType: endpoint.RecordTypeA},
		}},
	}
	fakeCache := newFakeCRDCache(t, nil, obj)

	pruning := interceptor.NewClient(fakeCache.Client.(client.WithWatch), interceptor.Funcs{
		SubResourceUpdate: func(
			ctx context.Context,
			c client.Client,
			subResource string,
			o client.Object,
			opts ...client.SubResourceUpdateOption) error {
			if subResource == "status" {
				o.(*apiv1alpha1.DNSEndpoint).Status.Conditions = nil
			}
			return c.Status().Update(ctx, o, opts...)
		},
	})

	cs, err := newCrdSource(t.Context(), fakeCache, pruning, nil, nil, nil)
	require.NoError(t, err)

	_, err = cs.Endpoints(t.Context())
	require.ErrorContains(t, err, "has no status.conditions")
}
