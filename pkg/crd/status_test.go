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
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/client/interceptor"

	apiv1alpha1 "sigs.k8s.io/external-dns/apis/v1alpha1"
	"sigs.k8s.io/external-dns/endpoint"
	logtest "sigs.k8s.io/external-dns/internal/testutils/log"
	"sigs.k8s.io/external-dns/pkg/events"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/source/types"
)

func TestSyncStatus(t *testing.T) {
	seed := &apiv1alpha1.DNSEndpoint{
		Namespace: "default", Name: "example", Generation: 2,
		Status: apiv1alpha1.DNSEndpointStatus{ObservedGeneration: 1},
	}
	seed2 := &apiv1alpha1.DNSEndpoint{
		Namespace: "other-ns", Name: "other", Generation: 5,
		Status: apiv1alpha1.DNSEndpointStatus{ObservedGeneration: 4},
	}
	seedMultiA := &apiv1alpha1.DNSEndpoint{
		Namespace: "ns-a", Name: "a", Generation: 1,
		Status: apiv1alpha1.DNSEndpointStatus{ObservedGeneration: 1},
	}
	seedMultiB := &apiv1alpha1.DNSEndpoint{
		Namespace: "ns-b", Name: "b", Generation: 3,
		Status: apiv1alpha1.DNSEndpointStatus{ObservedGeneration: 2},
	}

	tests := []struct {
		title string
		// seeded are the DNSEndpoint objects pre-populated in the fake client, as if already in the cluster.
		seeded           []client.Object
		changes          *plan.Changes
		interceptorFuncs interceptor.Funcs
		wantLogContains  []string
		wantLogAbsent    []string
	}{
		{
			title:  "groups create/update per referenced CR",
			seeded: []client.Object{seed.DeepCopy()},
			changes: &plan.Changes{
				Create: []*endpoint.Endpoint{
					endpoint.NewEndpoint("a.example.com", endpoint.RecordTypeA, "1.2.3.4").
						WithRefObject(dnsEndpointRef("default", "example")),
				},
				UpdateNew: []*endpoint.Endpoint{
					endpoint.NewEndpoint("b.example.com", endpoint.RecordTypeA, "1.2.3.5").
						WithRefObject(dnsEndpointRef("default", "example")),
				},
			},
			wantLogContains: []string{"DNSEndpoint default/example: 1 created, 1 updated (generation=2, observedGeneration=1)"},
		},
		{
			title:  "ignores endpoints without a DNSEndpoint reference",
			seeded: []client.Object{seed.DeepCopy()},
			changes: &plan.Changes{
				Create: []*endpoint.Endpoint{endpoint.NewEndpoint("a.example.com", endpoint.RecordTypeA, "1.2.3.4")},
			},
			wantLogAbsent: []string{"DNSEndpoint default/example"},
		},
		{
			title:  "ignores refs with a non-DNSEndpoint kind",
			seeded: []client.Object{seed.DeepCopy()},
			changes: &plan.Changes{
				Create: []*endpoint.Endpoint{
					endpoint.NewEndpoint("a.example.com", endpoint.RecordTypeA, "1.2.3.4").
						WithRefObject(serviceRef("default", "example")),
				},
			},
			wantLogAbsent: []string{"DNSEndpoint"},
		},
		{
			title:  "endpoint with mixed refs: non-DNSEndpoint ref ignored, DNSEndpoint ref tallied",
			seeded: []client.Object{seed.DeepCopy()},
			changes: &plan.Changes{
				Create: []*endpoint.Endpoint{
					endpoint.NewEndpoint("a.example.com", endpoint.RecordTypeA, "1.2.3.4").
						WithRefObject(serviceRef("default", "example")).
						WithRefObject(dnsEndpointRef("default", "example")),
				},
			},
			wantLogContains: []string{"DNSEndpoint default/example: 1 created, 0 updated (generation=2, observedGeneration=1)"},
		},
		{
			title: "non-NotFound Get error logs a warning",
			changes: &plan.Changes{
				Create: []*endpoint.Endpoint{
					endpoint.NewEndpoint("a.example.com", endpoint.RecordTypeA, "1.2.3.4").
						WithRefObject(dnsEndpointRef("default", "example")),
				},
			},
			interceptorFuncs: interceptor.Funcs{
				Get: func(context.Context, client.WithWatch, client.ObjectKey, client.Object, ...client.GetOption) error {
					return assert.AnError
				},
			},
			wantLogContains: []string{"Could not get DNSEndpoint default/example"},
		},
		{
			title:  "two different CRs get separate counts",
			seeded: []client.Object{seed.DeepCopy(), seed2.DeepCopy()},
			changes: &plan.Changes{
				Create: []*endpoint.Endpoint{
					endpoint.NewEndpoint("a.example.com", endpoint.RecordTypeA, "1.2.3.4").
						WithRefObject(dnsEndpointRef("default", "example")),
				},
				UpdateNew: []*endpoint.Endpoint{
					endpoint.NewEndpoint("b.example.com", endpoint.RecordTypeA, "1.2.3.5").
						WithRefObject(dnsEndpointRef("other-ns", "other")),
				},
			},
			wantLogContains: []string{
				"DNSEndpoint default/example: 1 created, 0 updated (generation=2, observedGeneration=1)",
				"DNSEndpoint other-ns/other: 0 created, 1 updated (generation=5, observedGeneration=4)",
			},
		},
		{
			title:  "one missing CR does not affect a sibling CR's lookup",
			seeded: []client.Object{seed.DeepCopy()},
			changes: &plan.Changes{
				Create: []*endpoint.Endpoint{
					endpoint.NewEndpoint("a.example.com", endpoint.RecordTypeA, "1.2.3.4").
						WithRefObject(dnsEndpointRef("default", "example")),
					endpoint.NewEndpoint("c.example.com", endpoint.RecordTypeA, "1.2.3.6").
						WithRefObject(dnsEndpointRef("default", "missing")),
				},
			},
			wantLogContains: []string{"DNSEndpoint default/example: 1 created, 0 updated (generation=2, observedGeneration=1)"},
			wantLogAbsent:   []string{"DNSEndpoint default/missing"},
		},
		{
			title:  "counts accumulate beyond one",
			seeded: []client.Object{seed.DeepCopy()},
			changes: &plan.Changes{
				Create: []*endpoint.Endpoint{
					endpoint.NewEndpoint("a.example.com", endpoint.RecordTypeA, "1.2.3.4").
						WithRefObject(dnsEndpointRef("default", "example")),
					endpoint.NewEndpoint("b.example.com", endpoint.RecordTypeA, "1.2.3.5").
						WithRefObject(dnsEndpointRef("default", "example")),
				},
			},
			wantLogContains: []string{"DNSEndpoint default/example: 2 created, 0 updated (generation=2, observedGeneration=1)"},
		},
		{
			title:  "update-only changes are tallied (Create empty)",
			seeded: []client.Object{seed.DeepCopy()},
			changes: &plan.Changes{
				UpdateNew: []*endpoint.Endpoint{
					endpoint.NewEndpoint("a.example.com", endpoint.RecordTypeA, "1.2.3.4").
						WithRefObject(dnsEndpointRef("default", "example")),
				},
			},
			wantLogContains: []string{"DNSEndpoint default/example: 0 created, 1 updated (generation=2, observedGeneration=1)"},
		},
		{
			title:  "one endpoint referencing two different CRs attributes both",
			seeded: []client.Object{seedMultiA.DeepCopy(), seedMultiB.DeepCopy()},
			changes: &plan.Changes{
				Create: []*endpoint.Endpoint{
					endpoint.NewEndpoint("a.example.com", endpoint.RecordTypeA, "1.2.3.4").
						WithRefObject(dnsEndpointRef("ns-a", "a")).
						WithRefObject(dnsEndpointRef("ns-b", "b")),
				},
			},
			wantLogContains: []string{
				"DNSEndpoint ns-a/a: 1 created, 0 updated (generation=1, observedGeneration=1)",
				"DNSEndpoint ns-b/b: 1 created, 0 updated (generation=3, observedGeneration=2)",
			},
		},
		{
			title:  "create and delete for the same CR: delete does not affect its count",
			seeded: []client.Object{seed.DeepCopy()},
			changes: &plan.Changes{
				Create: []*endpoint.Endpoint{
					endpoint.NewEndpoint("a.example.com", endpoint.RecordTypeA, "1.2.3.4").
						WithRefObject(dnsEndpointRef("default", "example")),
				},
				Delete: []*endpoint.Endpoint{
					endpoint.NewEndpoint("old.example.com", endpoint.RecordTypeA, "1.2.3.9").
						WithRefObject(dnsEndpointRef("default", "example")),
				},
			},
			wantLogContains: []string{"DNSEndpoint default/example: 1 created, 0 updated (generation=2, observedGeneration=1)"},
		},
		{
			title:  "delete-only changes produce no lookups",
			seeded: []client.Object{seed.DeepCopy()},
			changes: &plan.Changes{
				Delete: []*endpoint.Endpoint{
					endpoint.NewEndpoint("a.example.com", endpoint.RecordTypeA, "1.2.3.4").
						WithRefObject(dnsEndpointRef("default", "example")),
				},
			},
			interceptorFuncs: forbidGet(t),
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			hook := logtest.LogsUnderTestWithLogLevel(log.DebugLevel, t)
			c := fake.NewClientBuilder().
				WithScheme(newTestScheme(t)).
				WithStatusSubresource(&apiv1alpha1.DNSEndpoint{}).
				WithObjects(tt.seeded...).
				Build()
			wrapped := interceptor.NewClient(c, tt.interceptorFuncs)

			SyncStatus(t.Context(), NewCRDClients(wrapped, nil), tt.changes)

			for _, want := range tt.wantLogContains {
				logtest.TestHelperLogContains(want, hook, t)
			}
			for _, absent := range tt.wantLogAbsent {
				logtest.TestHelperLogNotContains(absent, hook, t)
			}
		})
	}
}

func TestSyncStatus_MissingCRSuppressesWarning(t *testing.T) {
	hook := logtest.LogsUnderTestWithLogLevel(log.DebugLevel, t)
	var getCalls int
	c := interceptor.NewClient(
		fake.NewClientBuilder().WithScheme(newTestScheme(t)).WithStatusSubresource(&apiv1alpha1.DNSEndpoint{}).Build(),
		interceptor.Funcs{
			Get: func(ctx context.Context, cli client.WithWatch, key client.ObjectKey, obj client.Object, opts ...client.GetOption) error {
				getCalls++
				return cli.Get(ctx, key, obj, opts...)
			},
		},
	)
	ep := endpoint.NewEndpoint("a.example.com", endpoint.RecordTypeA, "1.2.3.4").
		WithRefObject(dnsEndpointRef("default", "missing"))

	SyncStatus(t.Context(), NewCRDClients(c, nil), &plan.Changes{Create: []*endpoint.Endpoint{ep}})

	assert.Equal(t, 1, getCalls, "Get should have been attempted for the missing CR")
	for _, entry := range hook.AllEntries() {
		assert.NotEqual(t, log.WarnLevel, entry.Level, "unexpected warning logged: %s", entry.Message)
	}
}

func TestSyncStatus_NilGuards(t *testing.T) {
	t.Run("nil CRDClients", func(t *testing.T) {
		hook := logtest.LogsUnderTestWithLogLevel(log.DebugLevel, t)
		ep := endpoint.NewEndpoint("a.example.com", endpoint.RecordTypeA, "1.2.3.4").
			WithRefObject(dnsEndpointRef("default", "example"))
		require.NotPanics(t, func() {
			SyncStatus(t.Context(), nil, &plan.Changes{Create: []*endpoint.Endpoint{ep}})
		})
		assert.Empty(t, hook.AllEntries())
	})

	t.Run("non-nil CRDClients with nil Reader", func(t *testing.T) {
		hook := logtest.LogsUnderTestWithLogLevel(log.DebugLevel, t)
		ep := endpoint.NewEndpoint("a.example.com", endpoint.RecordTypeA, "1.2.3.4").
			WithRefObject(dnsEndpointRef("default", "example"))
		require.NotPanics(t, func() {
			SyncStatus(t.Context(), NewCRDClients(nil, nil), &plan.Changes{Create: []*endpoint.Endpoint{ep}})
		})
		assert.Empty(t, hook.AllEntries())
	})

	t.Run("nil changes", func(t *testing.T) {
		hook := logtest.LogsUnderTestWithLogLevel(log.DebugLevel, t)
		c := fake.NewClientBuilder().WithScheme(newTestScheme(t)).Build()
		require.NotPanics(t, func() {
			SyncStatus(t.Context(), NewCRDClients(c, nil), nil)
		})
		assert.Empty(t, hook.AllEntries())
	})

	t.Run("empty but non-nil changes", func(t *testing.T) {
		hook := logtest.LogsUnderTestWithLogLevel(log.DebugLevel, t)
		c := interceptor.NewClient(
			fake.NewClientBuilder().WithScheme(newTestScheme(t)).Build(),
			forbidGet(t),
		)
		require.NotPanics(t, func() {
			SyncStatus(t.Context(), NewCRDClients(c, nil), &plan.Changes{})
		})
		assert.Empty(t, hook.AllEntries())
	})
}

func newTestScheme(t *testing.T) *runtime.Scheme {
	t.Helper()
	s := runtime.NewScheme()
	require.NoError(t, apiv1alpha1.AddToScheme(s))
	return s
}

func dnsEndpointRef(namespace, name string) *events.ObjectReference {
	return events.NewObjectReferenceFromParts(
		dnsEndpointKind, apiv1alpha1.GroupVersion.String(), namespace, name, "", types.CRD)
}

func serviceRef(namespace, name string) *events.ObjectReference {
	return events.NewObjectReferenceFromParts("Service", "v1", namespace, name, "", "service")
}

// forbidGet fails the test if cl.Get is ever called, for asserting a code path
// that must not look anything up.
func forbidGet(t *testing.T) interceptor.Funcs {
	return interceptor.Funcs{
		Get: func(_ context.Context, _ client.WithWatch, key client.ObjectKey, _ client.Object, _ ...client.GetOption) error {
			t.Errorf("unexpected Get call for %v", key)
			return assert.AnError
		},
	}
}
