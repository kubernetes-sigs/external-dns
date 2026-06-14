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
	"maps"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/validation"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	apiv1alpha1 "sigs.k8s.io/external-dns/apis/v1alpha1"
	"sigs.k8s.io/external-dns/endpoint"
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
	_, err := NewCRDRegistry(nil, "", "", "", "", time.Second)
	assert.Error(t, err, "expected an error when no ownerID is specified")

	_, err = NewCRDRegistry(nil, "/dev/null", "", "default", "ownerID", time.Second)
	assert.Error(t, err, "expected an error when the kubeconfig is invalid")
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
	records  []*endpoint.Endpoint
	addLabel bool
	applyErr error
}

func (m *mockProvider) Records(_ context.Context) ([]*endpoint.Endpoint, error) {
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
