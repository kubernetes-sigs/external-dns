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

package registry

import (
	"context"
	"maps"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
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
	c := fake.NewClientBuilder().WithScheme(newTestScheme(t)).WithObjects(objs...).Build()
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
				apiv1alpha1.RecordOwnerLabel:         "test",
				apiv1alpha1.RecordNameLabel:          "sub.mytestdomain.io",
				apiv1alpha1.RecordTypeLabel:          "CNAME",
				apiv1alpha1.RecordSetIdentifierLabel: "myid-1",
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
		seedRecord *apiv1alpha1.DNSRecord
		assertFn   func(t *testing.T, c client.Client)
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
			assertFn: func(t *testing.T, c client.Client) {
				got := &apiv1alpha1.DNSRecord{}
				err := c.Get(ctx, types.NamespacedName{Namespace: "default", Name: "sub-mytestdomain-io-cname"}, got)
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
					Name:      "test-myid-2",
					Namespace: "default",
					Labels: map[string]string{
						apiv1alpha1.RecordOwnerLabel:         "test",
						apiv1alpha1.RecordNameLabel:          "to.be.deleted.mytestdomain.io",
						apiv1alpha1.RecordTypeLabel:          "A",
						apiv1alpha1.RecordSetIdentifierLabel: "myid-2",
					},
				},
			},
			assertFn: func(t *testing.T, c client.Client) {
				got := &apiv1alpha1.DNSRecord{}
				err := c.Get(ctx, types.NamespacedName{Namespace: "default", Name: "test-myid-2"}, got)
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
			seedRecord: &apiv1alpha1.DNSRecord{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-myid-3",
					Namespace: "default",
					Labels: map[string]string{
						apiv1alpha1.RecordOwnerLabel:         "test",
						apiv1alpha1.RecordNameLabel:          "to.be.updated.mytestdomain.io",
						apiv1alpha1.RecordTypeLabel:          "CNAME",
						apiv1alpha1.RecordSetIdentifierLabel: "myid-3",
					},
				},
				Spec: apiv1alpha1.DNSRecordSpec{Endpoint: endpoint.Endpoint{
					DNSName:    "to.be.updated.mytestdomain.io",
					RecordType: "CNAME",
					Targets:    endpoint.NewTargets("127.0.0.1"),
				}},
			},
			assertFn: func(t *testing.T, c client.Client) {
				got := &apiv1alpha1.DNSRecord{}
				err := c.Get(ctx, types.NamespacedName{Namespace: "default", Name: "test-myid-3"}, got)
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
				changes.UpdateNew = []*endpoint.Endpoint{tc.ep}
				changes.UpdateOld = []*endpoint.Endpoint{tc.ep}
				seedEndpoints = append(seedEndpoints, tc.ep)
			}

			prov := inMemoryProviderWithEntries(t, ctx, "mytestdomain.io", seedEndpoints...)
			var objs []client.Object
			if tc.seedRecord != nil {
				objs = append(objs, tc.seedRecord)
			}
			reg, c := newTestRegistry(t, prov, "test", objs...)

			require.NoError(t, reg.ApplyChanges(ctx, &changes))
			tc.assertFn(t, c)
		})
	}
}

type mockProvider struct {
	records  []*endpoint.Endpoint
	addLabel bool
}

func (m *mockProvider) Records(_ context.Context) ([]*endpoint.Endpoint, error) {
	return m.records, nil
}

// This applychanges in mocked provider simulate when
// the provider change Labels of the records.
func (m *mockProvider) ApplyChanges(_ context.Context, changes *plan.Changes) error {
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
			require.NoError(t, c.Get(t.Context(), types.NamespacedName{Namespace: "default", Name: "sub-mytestdomain-io-cname"}, got))

			_, hasLabel := got.Spec.Endpoint.Labels["prefix"]
			assert.Equal(t, tc.expectLabel, hasLabel)
		})
	}
}
