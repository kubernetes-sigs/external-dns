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

package crd

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"
	"unicode/utf8"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/client/interceptor"

	apiv1alpha1 "sigs.k8s.io/external-dns/apis/v1alpha1"
	logtest "sigs.k8s.io/external-dns/internal/testutils/log"
	"sigs.k8s.io/external-dns/pkg/events"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/source/types"
)

func newStatusWriter(t *testing.T, funcs interceptor.Funcs, objs ...client.Object) *StatusWriter {
	t.Helper()
	scheme := runtime.NewScheme()
	require.NoError(t, apiv1alpha1.AddToScheme(scheme))
	c := fake.NewClientBuilder().
		WithScheme(scheme).
		WithObjects(objs...).
		WithStatusSubresource(&apiv1alpha1.DNSEndpoint{}).
		WithInterceptorFuncs(funcs).
		Build()
	return NewStatusWriter(NewCRDClients(c, c), false)
}

// Every DNSEndpoint in this file lives in the same namespace.
const testNamespace = "foo"

func dnsEndpointRef(name string) *events.ObjectReference {
	return events.NewObjectReferenceFromParts("DNSEndpoint", "externaldns.k8s.io/v1alpha1", testNamespace, name, "", types.CRD)
}

// readDNSEndpoint fetches the stored copy, so assertions run against what the
// API server would hold rather than the object handed to the writer.
func readDNSEndpoint(t *testing.T, c client.Client, name string) *apiv1alpha1.DNSEndpoint {
	t.Helper()
	got := &apiv1alpha1.DNSEndpoint{}
	require.NoError(t, c.Get(t.Context(), client.ObjectKey{Namespace: testNamespace, Name: name}, got))
	return got
}

func TestStatusWriterReportStatus(t *testing.T) {
	for _, ti := range []struct {
		title           string
		planned         int
		unchanged       bool
		applyErr        error
		dryRun          bool
		wantStatus      metav1.ConditionStatus
		wantReason      string
		wantMessagePart string
	}{
		{
			title:           "provider applied the batch",
			planned:         1,
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
			// Its records were already in place, so the failed batch is not its own.
			title:           "provider rejected a batch the object was not in",
			planned:         1,
			unchanged:       true,
			applyErr:        errors.New("route53: throttled"),
			wantStatus:      metav1.ConditionTrue,
			wantReason:      apiv1alpha1.ProgrammedReason,
			wantMessagePart: "1 endpoint(s) applied to the DNS provider",
		},
		{
			// Nothing was offered to the provider, so Programmed would be a lie.
			title:           "every endpoint excluded by the filters",
			planned:         0,
			wantStatus:      metav1.ConditionFalse,
			wantReason:      apiv1alpha1.FilteredReason,
			wantMessagePart: "No endpoint reached the DNS provider",
		},
		{
			// --dry-run sends nothing, so Programmed would be a lie.
			title:           "dry run",
			planned:         1,
			dryRun:          true,
			wantStatus:      metav1.ConditionUnknown,
			wantReason:      apiv1alpha1.DryRunReason,
			wantMessagePart: "--dry-run kept them from the DNS provider",
		},
		{
			title:           "dry run with every endpoint filtered still reports the filter",
			planned:         0,
			dryRun:          true,
			wantStatus:      metav1.ConditionFalse,
			wantReason:      apiv1alpha1.FilteredReason,
			wantMessagePart: "No endpoint reached the DNS provider",
		},
	} {
		t.Run(ti.title, func(t *testing.T) {
			obj := &apiv1alpha1.DNSEndpoint{Namespace: testNamespace, Name: "test", Generation: 7}
			w := newStatusWriter(t, interceptor.Funcs{}, obj)
			w.dryRun = ti.dryRun

			w.ReportStatus(t.Context(), []plan.PlannedObject{{Ref: dnsEndpointRef("test"), Endpoints: ti.planned, Changed: !ti.unchanged}}, ti.applyErr)

			got := readDNSEndpoint(t, w.clients.Writer(), "test")
			cond := meta.FindStatusCondition(got.Status.Conditions, apiv1alpha1.ReadyCondition)
			require.NotNil(t, cond, "Ready condition must be set")
			assert.Equal(t, ti.wantStatus, cond.Status)
			assert.Equal(t, ti.wantReason, cond.Reason)
			assert.Equal(t, int64(7), cond.ObservedGeneration)
			assert.Contains(t, cond.Message, ti.wantMessagePart)
		})
	}
}

func TestStatusWriterSkipsWhatItCannotWrite(t *testing.T) {
	mine := &apiv1alpha1.DNSEndpoint{Namespace: testNamespace, Name: "mine", Generation: 3}

	t.Run("ignores foreign refs and deleted objects", func(t *testing.T) {
		hook := logtest.LogsUnderTestWithLogLevel(log.DebugLevel, t)
		w := newStatusWriter(t, interceptor.Funcs{}, mine)
		foreign := events.NewObjectReferenceFromParts("Ingress", "networking.k8s.io/v1", "foo", "mine", "", types.Ingress)

		w.ReportStatus(t.Context(), []plan.PlannedObject{
			{Ref: nil, Endpoints: 1},
			{Ref: foreign, Endpoints: 1},
			{Ref: dnsEndpointRef("gone"), Endpoints: 1},
		}, nil)

		got := readDNSEndpoint(t, w.clients.Writer(), "mine")
		assert.Empty(t, got.Status.Conditions, "no condition must be written for foreign or missing refs")
		logtest.TestHelperLogNotContains("foo/gone", hook, t)
	})

	t.Run("warns on read errors", func(t *testing.T) {
		hook := logtest.LogsUnderTestWithLogLevel(log.DebugLevel, t)
		w := newStatusWriter(t, interceptor.Funcs{
			Get: func(context.Context, client.WithWatch, client.ObjectKey, client.Object, ...client.GetOption) error {
				return errors.New("apiserver unavailable")
			},
		}, mine)

		w.ReportStatus(t.Context(), []plan.PlannedObject{{Ref: dnsEndpointRef("mine"), Endpoints: 1}}, nil)

		logtest.TestHelperLogContainsWithLogLevel("Could not get DNSEndpoint foo/mine: apiserver unavailable", log.WarnLevel, hook, t)
	})
}

// Accepted is written by the crd source from its listed (cache-backed) copy while
// Ready is written after the apply. The read path can lag the last write, so
// pushing the stale copy would drop the condition the other writer had just set.
// Its stale resourceVersion makes the API server reject it, and UpdateStatus
// then re-reads.
func TestUpdateStatusDoesNotClobberAStaleCondition(t *testing.T) {
	obj := &apiv1alpha1.DNSEndpoint{Namespace: testNamespace, Name: "test", Generation: 1}
	w := newStatusWriter(t, interceptor.Funcs{}, obj)
	c := w.clients.Writer()

	stale := readDNSEndpoint(t, c, "test")

	w.ReportStatus(t.Context(), []plan.PlannedObject{{Ref: dnsEndpointRef("test"), Endpoints: 1}}, nil)

	UpdateStatus(t.Context(), c, stale, func(status *apiv1alpha1.DNSEndpointStatus) {
		meta.SetStatusCondition(&status.Conditions, metav1.Condition{
			Type:    apiv1alpha1.AcceptedCondition,
			Status:  metav1.ConditionFalse,
			Reason:  apiv1alpha1.InvalidReason,
			Message: "spec.endpoints[0]: something changed",
		})
	})

	got := readDNSEndpoint(t, c, "test")
	accepted := meta.FindStatusCondition(got.Status.Conditions, apiv1alpha1.AcceptedCondition)
	require.NotNil(t, accepted)
	assert.Equal(t, apiv1alpha1.InvalidReason, accepted.Reason, "the new Accepted must be stored")

	ready := meta.FindStatusCondition(got.Status.Conditions, apiv1alpha1.ReadyCondition)
	require.NotNil(t, ready, "Ready must survive a write driven from a stale copy")
	assert.Equal(t, apiv1alpha1.ProgrammedReason, ready.Reason)
}

func TestTruncateConditionMessage(t *testing.T) {
	short := "all good"
	assert.Equal(t, short, TruncateConditionMessage(short))

	long := TruncateConditionMessage(fmt.Sprintf("%0*d", 40000, 0))
	assert.Len(t, long, 32768)
	assert.True(t, len(long) > 3 && long[len(long)-3:] == "...")

	// User-supplied names may be UTF-8; cutting bytes would split a rune.
	multibyte := TruncateConditionMessage(strings.Repeat("é", 40000))
	assert.Equal(t, 32768, utf8.RuneCountInString(multibyte))
	assert.True(t, utf8.ValidString(multibyte), "truncation must not split a rune")
}
