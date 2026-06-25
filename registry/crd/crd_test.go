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
	"maps"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/validation"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/client/interceptor"

	apiv1alpha1 "sigs.k8s.io/external-dns/apis/v1alpha1"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	logtest "sigs.k8s.io/external-dns/internal/testutils/log"
	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/provider/inmemory"
	"sigs.k8s.io/external-dns/registry"
)

// newTestScheme builds a scheme with the DNSRecord types registered.
func newTestScheme(t *testing.T) *runtime.Scheme {
	t.Helper()
	s := runtime.NewScheme()
	require.NoError(t, apiv1alpha1.AddToScheme(s))
	return s
}

// newTestRegistry wires a CRDRegistry on top of a controller-runtime fake client
// seeded with the given DNSRecords. The fake client serves both reads and writes.
func newTestRegistry(t *testing.T, p provider.Provider, ownerID string, objs ...client.Object) (*CRDRegistry, client.Client) {
	t.Helper()
	c := fake.NewClientBuilder().
		WithScheme(newTestScheme(t)).
		WithStatusSubresource(&apiv1alpha1.DNSRecord{}).
		WithObjects(objs...).
		Build()
	return &CRDRegistry{
		crReader:  c,
		crWriter:  c,
		namespace: "default",
		provider:  p,
		ownerID:   ownerID,
	}, c
}

// The endpoints needs to be part of the zone otherwise it will be filtered out.
func inMemoryProviderWithEntries(t *testing.T, ctx context.Context, zone string, endpoints ...*endpoint.Endpoint) *inmemory.InMemoryProvider {
	p := inmemory.NewInMemoryProvider(inmemory.InMemoryInitZones([]string{zone}))

	err := p.ApplyChanges(ctx, &plan.Changes{
		Create: endpoints,
	})
	if err != nil {
		t.Fatal("Could not create an in memory provider", err)
	}

	return p
}

func TestCRDRegistryImplementsRegistry(t *testing.T) {
	require.Implements(t, (*registry.Registry)(nil), new(CRDRegistry))
}

func TestCRDRegistryConstructor(t *testing.T) {
	tests := []struct {
		name            string
		kubeConfig      func(t *testing.T) string
		namespace       string
		ownerID         string
		timeout         time.Duration
		wantErrContains string
	}{
		{
			name:            "empty owner id",
			wantErrContains: "owner id cannot be empty",
		},
		{
			name:            "invalid kubeconfig",
			kubeConfig:      func(_ *testing.T) string { return "/dev/null" },
			namespace:       "default",
			ownerID:         "owner",
			timeout:         time.Second,
			wantErrContains: "unable to build rest config",
		},
		{
			name:            "bad TLS cert in kubeconfig",
			kubeConfig:      func(t *testing.T) string { return newFakeAPIServerKubeconfig(t, true) },
			namespace:       "default",
			ownerID:         "owner",
			timeout:         time.Second,
			wantErrContains: "unable to load root certificates",
		},
		{
			name:            "cache fails to sync",
			kubeConfig:      func(t *testing.T) string { return newFakeAPIServerKubeconfig(t, false) },
			namespace:       "default",
			ownerID:         "owner",
			timeout:         100 * time.Millisecond,
			wantErrContains: "cache failed to sync",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var kubeConfig string
			if tt.kubeConfig != nil {
				kubeConfig = tt.kubeConfig(t)
			}
			r, err := NewCRDRegistry(nil, kubeConfig, "", tt.namespace, tt.ownerID, tt.timeout)
			require.ErrorContains(t, err, tt.wantErrContains)
			require.Nil(t, r)
		})
	}
}

func TestNew(t *testing.T) {
	r, err := New(&externaldns.Config{}, nil)
	require.ErrorContains(t, err, "owner id cannot be empty")
	require.Nil(t, r)
}

func TestGetDomainFilter(t *testing.T) {
	reg, _ := newTestRegistry(t, &mockProvider{}, "owner")
	assert.Equal(t, &endpoint.DomainFilter{}, reg.GetDomainFilter())
}

func TestOwnerID(t *testing.T) {
	reg, _ := newTestRegistry(t, &mockProvider{}, "owner")
	assert.Equal(t, "owner", reg.OwnerID())
}

func TestCRDRegistryAdjustEndpoints(t *testing.T) {
	eps := []*endpoint.Endpoint{{DNSName: "foo.example.com", RecordType: "A"}}
	reg, _ := newTestRegistry(t, &mockProvider{}, "owner")
	got, err := reg.AdjustEndpoints(eps)
	require.NoError(t, err)
	assert.Equal(t, eps, got)
}

func TestCRDRegistryRecords(t *testing.T) {
	ctx := t.Context()

	prov := inMemoryProviderWithEntries(t, ctx, "mytestdomain.io", &endpoint.Endpoint{
		DNSName:       "sub.mytestdomain.io",
		RecordType:    "CNAME",
		SetIdentifier: "myid-1",
	})

	seeded := &apiv1alpha1.DNSRecord{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "sub-mytestdomain-io-cname",
			Namespace: "default",
			Labels: map[string]string{
				apiv1alpha1.RecordOwnerLabel: "test",
			},
		},
		Spec: apiv1alpha1.DNSRecordSpec{
			Endpoint: endpoint.Endpoint{
				DNSName:       "sub.mytestdomain.io",
				RecordType:    "CNAME",
				SetIdentifier: "myid-1",
				Labels:        map[string]string{endpoint.ResourceLabelKey: "resource"},
			},
		},
		Status: apiv1alpha1.DNSRecordStatus{
			Conditions: []metav1.Condition{{
				Type:   apiv1alpha1.ReadyCondition,
				Status: metav1.ConditionTrue,
				Reason: apiv1alpha1.ProgrammedReason,
			}},
		},
	}
	// A record owned by another instance must not be returned.
	other := &apiv1alpha1.DNSRecord{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "other-cname",
			Namespace: "default",
			Labels:    map[string]string{apiv1alpha1.RecordOwnerLabel: "someone-else"},
		},
	}

	reg, _ := newTestRegistry(t, prov, "test", seeded, other)

	endpoints, err := reg.Records(ctx)
	require.NoError(t, err)
	require.Len(t, endpoints, 1)
	assert.Equal(t, "resource", endpoints[0].Labels[endpoint.ResourceLabelKey])
}

func TestCRDRegistryRecordsListError(t *testing.T) {
	c := fake.NewClientBuilder().
		WithScheme(newTestScheme(t)).
		WithStatusSubresource(&apiv1alpha1.DNSRecord{}).
		Build()
	failReader := interceptor.NewClient(c, interceptor.Funcs{
		List: func(_ context.Context, _ client.WithWatch, _ client.ObjectList, _ ...client.ListOption) error {
			return assert.AnError
		},
	})
	reg := &CRDRegistry{
		crReader:  failReader,
		crWriter:  c,
		namespace: "default",
		provider:  &mockProvider{},
		ownerID:   "owner",
	}

	endpoints, err := reg.Records(t.Context())
	require.ErrorIs(t, err, assert.AnError)
	assert.Equal(t, []*endpoint.Endpoint{}, endpoints)
}

func TestCRDRegistryApplyChanges(t *testing.T) {
	ctx := t.Context()

	testCases := []struct {
		name       string
		changeType string // One of Create, Update, Delete
		ep         *endpoint.Endpoint
		oldEp      *endpoint.Endpoint // pre-change endpoint for Update; defaults to ep when nil
		seedRecord *apiv1alpha1.DNSRecord
		// assertFn receives the deterministic object name of tc.ep so cases do
		// not have to hard-code the hashed name.
		assertFn func(t *testing.T, c client.Client, name string)
	}{
		{
			name:       "Create",
			changeType: "Create",
			ep: &endpoint.Endpoint{
				DNSName:       "sub.mytestdomain.io",
				RecordType:    "CNAME",
				SetIdentifier: "myid-1",
				Targets:       endpoint.NewTargets("127.0.0.1"),
				Labels:        map[string]string{endpoint.OwnerLabelKey: "test"},
			},
			assertFn: func(t *testing.T, c client.Client, name string) {
				got := &apiv1alpha1.DNSRecord{}
				err := c.Get(ctx, types.NamespacedName{Namespace: "default", Name: name}, got)
				require.NoError(t, err)
				assert.Equal(t, "test", got.Labels[apiv1alpha1.RecordOwnerLabel])
			},
		},
		{
			name:       "Delete",
			changeType: "Delete",
			ep: &endpoint.Endpoint{
				DNSName:       "to.be.deleted.mytestdomain.io",
				RecordType:    "A",
				SetIdentifier: "myid-2",
				Targets:       endpoint.NewTargets("127.0.0.1"),
				Labels:        map[string]string{endpoint.OwnerLabelKey: "test"},
			},
			seedRecord: &apiv1alpha1.DNSRecord{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "default",
					Labels: map[string]string{
						apiv1alpha1.RecordOwnerLabel: "test",
					},
				},
			},
			assertFn: func(t *testing.T, c client.Client, name string) {
				got := &apiv1alpha1.DNSRecord{}
				err := c.Get(ctx, types.NamespacedName{Namespace: "default", Name: name}, got)
				assert.True(t, k8sErrors.IsNotFound(err), "expected DNSRecord to be deleted, got %v", err)
			},
		},
		{
			name:       "Update",
			changeType: "Update",
			ep: &endpoint.Endpoint{
				DNSName:       "to.be.updated.mytestdomain.io",
				RecordType:    "CNAME",
				SetIdentifier: "myid-3",
				Targets:       endpoint.NewTargets("127.0.0.2"),
				Labels:        map[string]string{endpoint.OwnerLabelKey: "test"},
			},
			// Same identity (name/type/set-id) as ep but the pre-change target.
			oldEp: &endpoint.Endpoint{
				DNSName:       "to.be.updated.mytestdomain.io",
				RecordType:    "CNAME",
				SetIdentifier: "myid-3",
				Targets:       endpoint.NewTargets("127.0.0.1"),
				Labels:        map[string]string{endpoint.OwnerLabelKey: "test"},
			},
			seedRecord: &apiv1alpha1.DNSRecord{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "default",
					Labels: map[string]string{
						apiv1alpha1.RecordOwnerLabel: "test",
					},
				},
				Spec: apiv1alpha1.DNSRecordSpec{Endpoint: endpoint.Endpoint{
					DNSName:       "to.be.updated.mytestdomain.io",
					RecordType:    "CNAME",
					SetIdentifier: "myid-3",
					Targets:       endpoint.NewTargets("127.0.0.1"),
				}},
			},
			assertFn: func(t *testing.T, c client.Client, name string) {
				got := &apiv1alpha1.DNSRecord{}
				err := c.Get(ctx, types.NamespacedName{Namespace: "default", Name: name}, got)
				require.NoError(t, err)
				assert.Equal(t, endpoint.NewTargets("127.0.0.2"), got.Spec.Endpoint.Targets)
			},
		},
		{
			name:       "Create refreshes spec when DNSRecord already exists",
			changeType: "Create",
			ep: &endpoint.Endpoint{
				DNSName:    "to.be.refreshed.mytestdomain.io",
				RecordType: "A",
				Targets:    endpoint.NewTargets("2.2.2.2"),
				Labels:     map[string]string{endpoint.OwnerLabelKey: "test"},
			},
			seedRecord: &apiv1alpha1.DNSRecord{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "default",
					Labels:    map[string]string{apiv1alpha1.RecordOwnerLabel: "test"},
				},
				Spec: apiv1alpha1.DNSRecordSpec{Endpoint: endpoint.Endpoint{
					DNSName:    "to.be.refreshed.mytestdomain.io",
					RecordType: "A",
					Targets:    endpoint.NewTargets("1.1.1.1"),
				}},
			},
			assertFn: func(t *testing.T, c client.Client, name string) {
				got := &apiv1alpha1.DNSRecord{}
				err := c.Get(ctx, types.NamespacedName{Namespace: "default", Name: name}, got)
				require.NoError(t, err)
				assert.Equal(t, endpoint.NewTargets("2.2.2.2"), got.Spec.Endpoint.Targets)
			},
		},
		{
			name:       "Update returns no error when DNSRecord is missing",
			changeType: "Update",
			ep: &endpoint.Endpoint{
				DNSName:    "noop.update.mytestdomain.io",
				RecordType: "A",
				Targets:    endpoint.NewTargets("127.0.0.1"),
				Labels:     map[string]string{endpoint.OwnerLabelKey: "test"},
			},
			assertFn: func(t *testing.T, c client.Client, name string) {
				got := &apiv1alpha1.DNSRecord{}
				err := c.Get(ctx, types.NamespacedName{Namespace: "default", Name: name}, got)
				assert.True(t, k8sErrors.IsNotFound(err))
			},
		},
		{
			name:       "Delete returns no error when DNSRecord is missing",
			changeType: "Delete",
			ep: &endpoint.Endpoint{
				DNSName:    "noop.delete.mytestdomain.io",
				RecordType: "A",
				Targets:    endpoint.NewTargets("127.0.0.1"),
				Labels:     map[string]string{endpoint.OwnerLabelKey: "test"},
			},
			assertFn: func(t *testing.T, c client.Client, name string) {
				got := &apiv1alpha1.DNSRecord{}
				err := c.Get(ctx, types.NamespacedName{Namespace: "default", Name: name}, got)
				assert.True(t, k8sErrors.IsNotFound(err))
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var seedEndpoints []*endpoint.Endpoint
			var changes plan.Changes
			switch tc.changeType {
			case "Create":
				changes.Create = []*endpoint.Endpoint{tc.ep}
			case "Delete":
				changes.Delete = []*endpoint.Endpoint{tc.ep}
				seedEndpoints = append(seedEndpoints, tc.ep)
			case "Update":
				oldEp := tc.oldEp
				if oldEp == nil {
					oldEp = tc.ep
				}
				changes.UpdateNew = []*endpoint.Endpoint{tc.ep}
				changes.UpdateOld = []*endpoint.Endpoint{oldEp}
				seedEndpoints = append(seedEndpoints, oldEp)
			}

			// The object name is derived deterministically from the endpoint
			// identity, so seeded records must use the same name the registry
			// will look up.
			name := recordObjectName(tc.ep)
			prov := inMemoryProviderWithEntries(t, ctx, "mytestdomain.io", seedEndpoints...)
			var objs []client.Object
			if tc.seedRecord != nil {
				tc.seedRecord.Name = name
				objs = append(objs, tc.seedRecord)
			}
			reg, c := newTestRegistry(t, prov, "test", objs...)

			require.NoError(t, reg.ApplyChanges(ctx, &changes))
			tc.assertFn(t, c, name)
		})
	}
}

func TestRecordObjectName(t *testing.T) {
	ep := func(dns, recordType, setID string) *endpoint.Endpoint {
		return &endpoint.Endpoint{DNSName: dns, RecordType: recordType, SetIdentifier: setID}
	}

	t.Run("is deterministic", func(t *testing.T) {
		first := recordObjectName(ep("sub.example.io", "A", ""))
		second := recordObjectName(ep("sub.example.io", "A", ""))
		assert.Equal(t, first, second)
	})

	t.Run("produces RFC 1123 compliant names", func(t *testing.T) {
		cases := []*endpoint.Endpoint{
			ep("sub.example.io", "A", ""),
			ep("*.example.io", "A", ""),                         // wildcard
			ep("_dmarc.example.io", "TXT", ""),                  // underscore
			ep("example.io.", "A", ""),                          // trailing dot
			ep(strings.Repeat("a.", 200)+"example.io", "A", ""), // > 253 chars
		}
		for _, e := range cases {
			name := recordObjectName(e)
			assert.LessOrEqual(t, len(name), 253, "name %q exceeds 253 chars", name)
			assert.Empty(t, validation.IsDNS1123Subdomain(name), "name %q is not a valid RFC 1123 subdomain", name)
		}
	})

	t.Run("falls back to 'record' prefix when DNSName has no valid chars", func(t *testing.T) {
		assert.Equal(t, "record-6fbee3c6", recordObjectName(ep("...", "A", "")))
	})

	t.Run("distinct identities never collide", func(t *testing.T) {
		// Cases that collapsed to the same name under the old dns-with-dashes scheme.
		collisions := [][2]*endpoint.Endpoint{
			{ep("sub.example.io", "A", "eu"), ep("sub.example.io", "A", "us")}, // set identifier
			{ep("sub.example.io", "A", ""), ep("sub-example.io", "A", "")},     // dot vs dash
			{ep("sub.example.io", "A", ""), ep("sub.example.io", "AAAA", "")},  // record type
		}
		for _, c := range collisions {
			assert.NotEqual(t, recordObjectName(c[0]), recordObjectName(c[1]))
		}
	})
}

type mockProvider struct {
	records    []*endpoint.Endpoint
	addLabel   bool
	applyErr   error
	recordsErr error
}

func (m *mockProvider) Records(_ context.Context) ([]*endpoint.Endpoint, error) {
	if m.recordsErr != nil {
		return nil, m.recordsErr
	}
	return m.records, nil
}

// This applychanges in mocked provider simulate when
// the provider change Labels of the records.
func (m *mockProvider) ApplyChanges(_ context.Context, changes *plan.Changes) error {
	if m.applyErr != nil {
		return m.applyErr
	}
	endpoints := changes.Create
	m.records = make([]*endpoint.Endpoint, 0, len(endpoints))
	for _, ep := range endpoints {
		newEp := endpoint.NewEndpointWithTTL(ep.DNSName, ep.RecordType, ep.RecordTTL, ep.Targets...).WithSetIdentifier(ep.SetIdentifier)
		newEp.Labels = endpoint.NewLabels()
		maps.Copy(newEp.Labels, ep.Labels)
		newEp.ProviderSpecific = append(endpoint.ProviderSpecific(nil), ep.ProviderSpecific...)
		// mocked specific change
		if m.addLabel {
			newEp.Labels["prefix"] = "random"
		}
		m.records = append(m.records, newEp)
	}
	return nil
}

func (m *mockProvider) AdjustEndpoints(eps []*endpoint.Endpoint) ([]*endpoint.Endpoint, error) {
	return eps, nil
}

func (m *mockProvider) GetDomainFilter() endpoint.DomainFilterInterface {
	return &endpoint.DomainFilter{}
}

// Ensure record is updated with expected Labels even when they are modified
// by the provider.ApplyChanges like, for instance, with coredns provider
func TestCRDApplyChangesMockedProvider(t *testing.T) {
	testCases := []struct {
		name        string
		provider    provider.Provider
		expectLabel bool
	}{
		{name: "provider adds a label", provider: &mockProvider{addLabel: true}, expectLabel: true},
		{name: "provider keeps labels", provider: &mockProvider{addLabel: false}, expectLabel: false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ep := endpoint.Endpoint{
				DNSName:       "sub.mytestdomain.io",
				RecordType:    "CNAME",
				SetIdentifier: "myid-1",
				Targets:       endpoint.NewTargets("127.0.0.1"),
				Labels:        map[string]string{endpoint.OwnerLabelKey: "owner"},
			}

			reg, c := newTestRegistry(t, tc.provider, "owner")
			require.NoError(t, reg.ApplyChanges(t.Context(), &plan.Changes{Create: []*endpoint.Endpoint{&ep}}))

			got := &apiv1alpha1.DNSRecord{}
			require.NoError(t, c.Get(t.Context(), types.NamespacedName{Namespace: "default", Name: recordObjectName(&ep)}, got))

			_, hasLabel := got.Spec.Endpoint.Labels["prefix"]
			assert.Equal(t, tc.expectLabel, hasLabel)
		})
	}
}

// A successful apply must mark the DNSRecord Ready with reason Programmed, so
// the stored record carries visible status.
func TestCRDApplyChangesSetsProgrammedStatus(t *testing.T) {
	ctx := t.Context()
	prov := inMemoryProviderWithEntries(t, ctx, "mytestdomain.io")

	ep := endpoint.Endpoint{
		DNSName:    "sub.mytestdomain.io",
		RecordType: "CNAME",
		Targets:    endpoint.NewTargets("127.0.0.1"),
		Labels:     map[string]string{endpoint.OwnerLabelKey: "test"},
	}

	reg, c := newTestRegistry(t, prov, "test")
	require.NoError(t, reg.ApplyChanges(ctx, &plan.Changes{Create: []*endpoint.Endpoint{&ep}}))

	got := &apiv1alpha1.DNSRecord{}
	require.NoError(t, c.Get(ctx, types.NamespacedName{Namespace: "default", Name: recordObjectName(&ep)}, got))

	cond := meta.FindStatusCondition(got.Status.Conditions, apiv1alpha1.ReadyCondition)
	require.NotNil(t, cond, "expected a Ready condition")
	assert.Equal(t, metav1.ConditionTrue, cond.Status)
	assert.Equal(t, apiv1alpha1.ProgrammedReason, cond.Reason)
}

// When the provider rejects the changes, the DNSRecord is persisted with a
// Ready=False/Failed condition, so the failure is visible on the object while
// Records() still excludes it from current state (it is not Ready) and the plan
// re-applies it on the next reconcile.
func TestCRDApplyChangesProviderError(t *testing.T) {
	ctx := t.Context()
	ep := endpoint.Endpoint{
		DNSName:       "sub.mytestdomain.io",
		RecordType:    "CNAME",
		SetIdentifier: "myid-1",
		Targets:       endpoint.NewTargets("127.0.0.1"),
		Labels:        map[string]string{endpoint.OwnerLabelKey: "owner"},
	}

	reg, c := newTestRegistry(t, &mockProvider{applyErr: errors.New("boom")}, "owner")
	err := reg.ApplyChanges(ctx, &plan.Changes{Create: []*endpoint.Endpoint{&ep}})
	require.Error(t, err)

	got := &apiv1alpha1.DNSRecord{}
	require.NoError(t, c.Get(ctx, types.NamespacedName{Namespace: "default", Name: recordObjectName(&ep)}, got))

	cond := meta.FindStatusCondition(got.Status.Conditions, apiv1alpha1.ReadyCondition)
	require.NotNil(t, cond, "expected a Ready condition")
	assert.Equal(t, metav1.ConditionFalse, cond.Status)
	assert.Equal(t, apiv1alpha1.FailedReason, cond.Reason)

	// A non-Ready record must not be reported as current state.
	endpoints, err := reg.Records(ctx)
	require.NoError(t, err)
	assert.Empty(t, endpoints, "failed record must be excluded from current state")
}

func TestCRDApplyChangesErrorPaths(t *testing.T) {
	ctx := t.Context()

	makeEp := func(dns string) *endpoint.Endpoint {
		return &endpoint.Endpoint{
			DNSName:    dns,
			RecordType: "A",
			Targets:    endpoint.NewTargets("127.0.0.1"),
			Labels:     map[string]string{endpoint.OwnerLabelKey: "test"},
		}
	}
	seedFor := func(e *endpoint.Endpoint) *apiv1alpha1.DNSRecord {
		return &apiv1alpha1.DNSRecord{
			ObjectMeta: metav1.ObjectMeta{
				Name:      recordObjectName(e),
				Namespace: "default",
				Labels:    map[string]string{apiv1alpha1.RecordOwnerLabel: "test"},
			},
			Spec: apiv1alpha1.DNSRecordSpec{Endpoint: *e},
		}
	}

	epCreate := makeEp("create.example.io")
	epUpdate := makeEp("update.example.io")
	epDelete := makeEp("delete.example.io")
	epExist := makeEp("exists.example.io")

	alreadyExists := func(_ context.Context, _ client.WithWatch, _ client.Object, _ ...client.CreateOption) error {
		return k8sErrors.NewAlreadyExists(schema.GroupResource{Resource: "dnsrecords"}, "")
	}
	notFound := func(_ context.Context, _ client.WithWatch, _ client.ObjectKey, _ client.Object, _ ...client.GetOption) error {
		return k8sErrors.NewNotFound(schema.GroupResource{Resource: "dnsrecords"}, "")
	}

	tests := []struct {
		name            string
		changes         *plan.Changes
		seeded          []client.Object
		fns             interceptor.Funcs
		provider        provider.Provider
		wantErrContains string
		wantErrIs       error
		wantLogContains string
	}{
		{
			name:     "Create write fails",
			changes:  &plan.Changes{Create: []*endpoint.Endpoint{epCreate}},
			provider: &mockProvider{},
			fns: interceptor.Funcs{
				Create: func(_ context.Context, _ client.WithWatch, _ client.Object, _ ...client.CreateOption) error {
					return assert.AnError
				},
			},
			wantErrContains: "unable to create DNSRecord",
			wantErrIs:       assert.AnError,
		},
		{
			name: "Update get fails",
			changes: &plan.Changes{
				UpdateNew: []*endpoint.Endpoint{epUpdate},
				UpdateOld: []*endpoint.Endpoint{epUpdate},
			},
			seeded:   []client.Object{seedFor(epUpdate)},
			provider: &mockProvider{},
			fns: interceptor.Funcs{
				Get: func(_ context.Context, _ client.WithWatch, _ client.ObjectKey, _ client.Object, _ ...client.GetOption) error {
					return assert.AnError
				},
			},
			wantErrContains: "unable to get DNSRecord",
			wantErrIs:       assert.AnError,
		},
		{
			name: "Update write fails",
			changes: &plan.Changes{
				UpdateNew: []*endpoint.Endpoint{epUpdate},
				UpdateOld: []*endpoint.Endpoint{epUpdate},
			},
			seeded:   []client.Object{seedFor(epUpdate)},
			provider: &mockProvider{},
			fns: interceptor.Funcs{
				Update: func(_ context.Context, _ client.WithWatch, _ client.Object, _ ...client.UpdateOption) error {
					return assert.AnError
				},
			},
			wantErrContains: "unable to update DNSRecord",
			wantErrIs:       assert.AnError,
		},
		{
			name:     "Delete get fails",
			changes:  &plan.Changes{Delete: []*endpoint.Endpoint{epDelete}},
			seeded:   []client.Object{seedFor(epDelete)},
			provider: &mockProvider{},
			fns: interceptor.Funcs{
				Get: func(_ context.Context, _ client.WithWatch, _ client.ObjectKey, _ client.Object, _ ...client.GetOption) error {
					return assert.AnError
				},
			},
			wantErrContains: "unable to get DNSRecord",
			wantErrIs:       assert.AnError,
		},
		{
			name:     "Delete write fails",
			changes:  &plan.Changes{Delete: []*endpoint.Endpoint{epDelete}},
			seeded:   []client.Object{seedFor(epDelete)},
			provider: &mockProvider{},
			fns: interceptor.Funcs{
				Delete: func(_ context.Context, _ client.WithWatch, _ client.Object, _ ...client.DeleteOption) error {
					return assert.AnError
				},
			},
			wantErrContains: "unable to delete DNSRecord",
			wantErrIs:       assert.AnError,
		},
		{
			name:     "Delete ignores NotFound",
			changes:  &plan.Changes{Delete: []*endpoint.Endpoint{epDelete}},
			seeded:   []client.Object{seedFor(epDelete)},
			provider: &mockProvider{},
			fns: interceptor.Funcs{
				Delete: func(_ context.Context, _ client.WithWatch, _ client.Object, _ ...client.DeleteOption) error {
					return k8sErrors.NewNotFound(schema.GroupResource{Resource: "dnsrecords"}, "")
				},
			},
		},
		{
			name:     "Create AlreadyExists then Update fails",
			changes:  &plan.Changes{Create: []*endpoint.Endpoint{epExist}},
			seeded:   []client.Object{seedFor(epExist)},
			provider: &mockProvider{},
			fns: interceptor.Funcs{
				Update: func(_ context.Context, _ client.WithWatch, _ client.Object, _ ...client.UpdateOption) error {
					return assert.AnError
				},
			},
			wantErrContains: "unable to update DNSRecord",
			wantErrIs:       assert.AnError,
		},
		{
			name:     "Create AlreadyExists then Get fails",
			changes:  &plan.Changes{Create: []*endpoint.Endpoint{epExist}},
			seeded:   []client.Object{seedFor(epExist)},
			provider: &mockProvider{},
			fns: interceptor.Funcs{
				Get: func(_ context.Context, _ client.WithWatch, _ client.ObjectKey, _ client.Object, _ ...client.GetOption) error {
					return assert.AnError
				},
			},
			wantErrContains: "unable to get DNSRecord",
			wantErrIs:       assert.AnError,
		},
		{
			name:     "Create AlreadyExists but record vanished",
			changes:  &plan.Changes{Create: []*endpoint.Endpoint{epExist}},
			provider: &mockProvider{},
			fns: interceptor.Funcs{
				Create: alreadyExists,
				Get:    notFound,
			},
			wantErrContains: "reported as existing but was not found",
		},
		{
			name:            "Provider Records error during label adjust",
			changes:         &plan.Changes{Create: []*endpoint.Endpoint{epCreate}},
			provider:        &mockProvider{recordsErr: assert.AnError},
			wantErrContains: "unable to get records from provider",
			wantErrIs:       assert.AnError,
		},
		{
			name:     "Label merge Update fails",
			changes:  &plan.Changes{Create: []*endpoint.Endpoint{epCreate}},
			provider: &mockProvider{addLabel: true},
			fns: interceptor.Funcs{
				Update: func(_ context.Context, _ client.WithWatch, _ client.Object, _ ...client.UpdateOption) error {
					return assert.AnError
				},
			},
			wantErrContains: "unable to update DNSRecord",
			wantErrIs:       assert.AnError,
			wantLogContains: "update DNSRecord with modified labels from provider",
		},
		{
			name:     "Status update error is logged not returned",
			changes:  &plan.Changes{Create: []*endpoint.Endpoint{epCreate}},
			provider: &mockProvider{},
			fns: interceptor.Funcs{
				SubResourceUpdate: func(_ context.Context, _ client.Client, _ string, _ client.Object, _ ...client.SubResourceUpdateOption) error {
					return assert.AnError
				},
			},
			wantLogContains: "unable to update status of DNSRecord",
		},
		{
			// Uses Update instead of Create so the mock provider does not auto-populate
			// a matching record, forcing the registry to hit the no-match branch.
			name: "No provider record matches applied DNSRecord",
			changes: &plan.Changes{
				UpdateNew: []*endpoint.Endpoint{epUpdate},
				UpdateOld: []*endpoint.Endpoint{epUpdate},
			},
			seeded:          []client.Object{seedFor(epUpdate)},
			provider:        &mockProvider{},
			wantLogContains: "no provider record matched DNSRecord",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hook := logtest.LogsUnderTestWithLogLevel(log.DebugLevel, t)
			c := fake.NewClientBuilder().
				WithScheme(newTestScheme(t)).
				WithStatusSubresource(&apiv1alpha1.DNSRecord{}).
				WithObjects(tt.seeded...).
				Build()
			wrapped := interceptor.NewClient(c, tt.fns)
			reg := &CRDRegistry{
				crReader:  wrapped,
				crWriter:  wrapped,
				namespace: "default",
				provider:  tt.provider,
				ownerID:   "test",
			}
			err := reg.ApplyChanges(ctx, tt.changes)
			if tt.wantErrContains == "" {
				require.NoError(t, err)
			} else {
				require.ErrorContains(t, err, tt.wantErrContains)
				if tt.wantErrIs != nil {
					require.ErrorIs(t, err, tt.wantErrIs)
				}
			}
			if tt.wantLogContains != "" {
				logtest.TestHelperLogContains(tt.wantLogContains, hook, t)
			}
		})
	}
}

// newFakeAPIServerKubeconfig writes a minimal kubeconfig. When badTLS is true,
// invalid CA data forces transport construction to fail; otherwise it points at
// a fake discovery server with TLS verification disabled.
func newFakeAPIServerKubeconfig(t *testing.T, badTLS bool) string {
	t.Helper()
	apiResource := metav1.APIResource{Name: "dnsrecords", Namespaced: true, Kind: "DNSRecord", Verbs: metav1.Verbs{"list", "watch"}}
	server := "https://127.0.0.1:1"
	tlsLine := "certificate-authority-data: bm90LWEtcGVtLWNlcnQ="
	if !badTLS {
		server = testutils.NewFakeExternalDNSDiscoveryServer(t, apiResource).URL
		tlsLine = "insecure-skip-tls-verify: true"
	}
	path := filepath.Join(t.TempDir(), "kubeconfig")
	content := fmt.Sprintf(`apiVersion: v1
kind: Config
clusters:
- name: fake
  cluster:
    server: %s
    %s
contexts:
- name: fake
  context:
    cluster: fake
    user: fake
current-context: fake
users:
- name: fake
  user: {}
`, server, tlsLine)
	require.NoError(t, os.WriteFile(path, []byte(content), 0o600))
	return path
}
