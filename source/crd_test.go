/*
Copyright 2018 The Kubernetes Authors.

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
	"math/rand"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/cache"
	cachetesting "k8s.io/client-go/tools/cache/testing"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	log "github.com/sirupsen/logrus"

	apiv1alpha1 "sigs.k8s.io/external-dns/apis/v1alpha1"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/source/informers"
	"sigs.k8s.io/external-dns/internal/testutils"
)

var (
	_ Source = &crdSource{}
)

// Test-only fields for crdSource to allow tests to access the informer and watcher
type crdSourceTest struct {
	*crdSource
	testInformer cache.SharedIndexInformer
	testWatcher  *cachetesting.FakeControllerSource
}

// getInformer returns the test informer
func (cs *crdSourceTest) getInformer() cache.SharedIndexInformer {
	return cs.testInformer
}

// getWatcher returns the test watcher for manipulating the informer's cache
func (cs *crdSourceTest) getWatcher() *cachetesting.FakeControllerSource {
	return cs.testWatcher
}

type CRDSuite struct {
	suite.Suite
}

func newFakeClientBuilder(scheme *runtime.Scheme, objects ...runtime.Object) *fake.ClientBuilder {
	return fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(objects...)
}

func TestCRDSource(t *testing.T) {
	suite.Run(t, new(CRDSuite))
	t.Run("Endpoints", testCRDSourceEndpoints)
}

// testCRDSourceEndpoints tests various scenarios of using CRD source.
func testCRDSourceEndpoints(t *testing.T) {
	for _, ti := range []struct {
		title            string
		namespace        string
		objectNamespace  string
		endpoints        []*endpoint.Endpoint
		expectEndpoints  bool
		expectError      bool
		annotationFilter string
		labelFilter      string
		annotations      map[string]string
		labels           map[string]string
	}{
		{
			title:           "endpoints within a specific namespace",
			namespace:       "foo",
			objectNamespace: "foo",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "abc.example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  180,
				},
			},
			expectEndpoints: true,
		},
		{
			title:           "no endpoints within a specific namespace",
			namespace:       "foo",
			objectNamespace: "bar",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "abc.example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  180,
				},
			},
		},
		{
			title:           "valid crd with no targets (relies on default-targets)",
			namespace:       "foo",
			objectNamespace: "foo",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "no-targets.example.org",
					Targets:    endpoint.Targets{},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  180,
				},
			},
			expectEndpoints: true,
		},
		{
			title:           "valid crd gvk with single endpoint",
			namespace:       "foo",
			objectNamespace: "foo",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "abc.example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  180,
				},
			},
			expectEndpoints: true,
		},
		{
			title:           "valid crd gvk with multiple endpoints",
			namespace:       "foo",
			objectNamespace: "foo",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "abc.example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  180,
				},
				{
					DNSName:    "xyz.example.org",
					Targets:    endpoint.Targets{"abc.example.org"},
					RecordType: endpoint.RecordTypeCNAME,
					RecordTTL:  180,
				},
			},
			expectEndpoints: true,
		},
		{
			title:            "valid crd gvk with annotation and non matching annotation filter",
			namespace:        "foo",
			objectNamespace:  "foo",
			annotations:      map[string]string{"test": "that"},
			annotationFilter: "test=filter_something_else",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "abc.example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  180,
				},
			},
		},
		{
			title:            "valid crd gvk with annotation and matching annotation filter",
			namespace:        "foo",
			objectNamespace:  "foo",
			annotations:      map[string]string{"test": "that"},
			annotationFilter: "test=that",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "abc.example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  180,
				},
			},
			expectEndpoints: true,
		},
		{
			title:           "valid crd gvk with label and non matching label filter",
			namespace:       "foo",
			objectNamespace: "foo",
			labels:          map[string]string{"test": "that"},
			labelFilter:     "test=filter_something_else",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "abc.example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  180,
				},
			},
		},
		{
			title:           "valid crd gvk with label and matching label filter",
			namespace:       "foo",
			objectNamespace: "foo",
			labels:          map[string]string{"test": "that"},
			labelFilter:     "test=that",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "abc.example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  180,
				},
			},
			expectEndpoints: true,
		},
		{
			title:           "Create NS record",
			namespace:       "foo",
			objectNamespace: "foo",
			labels:          map[string]string{"test": "that"},
			labelFilter:     "test=that",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "abc.example.org",
					Targets:    endpoint.Targets{"ns1.k8s.io", "ns2.k8s.io"},
					RecordType: endpoint.RecordTypeNS,
					RecordTTL:  180,
				},
			},
			expectEndpoints: true,
		},
		{
			title:           "Create SRV record",
			namespace:       "foo",
			objectNamespace: "foo",
			labels:          map[string]string{"test": "that"},
			labelFilter:     "test=that",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "_svc._tcp.example.org",
					Targets:    endpoint.Targets{"0 0 80 abc.example.org", "0 0 80 def.example.org"},
					RecordType: endpoint.RecordTypeSRV,
					RecordTTL:  180,
				},
			},
			expectEndpoints: true,
		},
		{
			title:           "Create NAPTR record",
			namespace:       "foo",
			objectNamespace: "foo",
			labels:          map[string]string{"test": "that"},
			labelFilter:     "test=that",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "example.org",
					Targets:    endpoint.Targets{`100 10 "S" "SIP+D2U" "!^.*$!sip:customer-service@example.org!" _sip._udp.example.org.`, `102 10 "S" "SIP+D2T" "!^.*$!sip:customer-service@example.org!" _sip._tcp.example.org.`},
					RecordType: endpoint.RecordTypeNAPTR,
					RecordTTL:  180,
				},
			},
			expectEndpoints: true,
		},
		{
			title:           "illegal target CNAME",
			namespace:       "foo",
			objectNamespace: "foo",
			labels:          map[string]string{"test": "that"},
			labelFilter:     "test=that",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "example.org",
					Targets:    endpoint.Targets{"foo.example.org."},
					RecordType: endpoint.RecordTypeCNAME,
					RecordTTL:  180,
				},
			},
			expectEndpoints: true,
		},
		{
			title:                "CNAME target without trailing dot (relative name)",
			registeredAPIVersion: apiv1alpha1.GroupVersion.String(),
			apiVersion:           apiv1alpha1.GroupVersion.String(),
			registeredKind:       apiv1alpha1.DNSEndpointKind,
			kind:                 apiv1alpha1.DNSEndpointKind,
			namespace:            "foo",
			registeredNamespace:  "foo",
			labels:               map[string]string{"test": "that"},
			labelFilter:          "test=that",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "internal.example.com",
					Targets:    endpoint.Targets{"backend.cluster.local"},
					RecordType: endpoint.RecordTypeCNAME,
					RecordTTL:  300,
				},
			},
			expectEndpoints: true,
		},
		{
			title:           "illegal target NAPTR",
			namespace:       "foo",
			objectNamespace: "foo",
			labels:          map[string]string{"test": "that"},
			labelFilter:     "test=that",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "example.org",
					Targets:    endpoint.Targets{`100 10 "S" "SIP+D2U" "!^.*$!sip:customer-service@example.org!" _sip._udp.example.org`, `102 10 "S" "SIP+D2T" "!^.*$!sip:customer-service@example.org!" _sip._tcp.example.org`},
					RecordType: endpoint.RecordTypeNAPTR,
					RecordTTL:  180,
				},
			},
		},
		{
			title:           "valid target TXT",
			namespace:       "foo",
			objectNamespace: "foo",
			labels:          map[string]string{"test": "that"},
			labelFilter:     "test=that",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "example.org",
					Targets:    endpoint.Targets{"foo.example.org."},
					RecordType: endpoint.RecordTypeTXT,
					RecordTTL:  180,
				},
			},
			expectEndpoints: true,
		},
		{
			title:           "illegal target A",
			namespace:       "foo",
			objectNamespace: "foo",
			labels:          map[string]string{"test": "that"},
			labelFilter:     "test=that",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "example.org",
					Targets:    endpoint.Targets{"1.2.3.4."},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  180,
				},
			},
		},
		{
			title:           "MX Record allowing trailing dot in target",
			namespace:       "foo",
			objectNamespace: "foo",
			labels:          map[string]string{"test": "that"},
			labelFilter:     "test=that",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "example.org",
					Targets:    endpoint.Targets{"example.com."},
					RecordType: endpoint.RecordTypeMX,
					RecordTTL:  180,
				},
			},
			expectEndpoints: true,
			expectError:     false,
		},
		{
			title:           "MX Record without trailing dot in target",
			namespace:       "foo",
			objectNamespace: "foo",
			labels:          map[string]string{"test": "that"},
			labelFilter:     "test=that",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "example.org",
					Targets:    endpoint.Targets{"example.com"},
					RecordType: endpoint.RecordTypeMX,
					RecordTTL:  180,
				},
			},
			expectEndpoints: true,
			expectError:     false,
		},
	} {
		t.Run(ti.title, func(t *testing.T) {
			t.Parallel()

			scheme := runtime.NewScheme()
			err := apiv1alpha1.AddToScheme(scheme)
			require.NoError(t, err)

			labelSelector, err := labels.Parse(ti.labelFilter)
			require.NoError(t, err)

			dnsEndpoint := &apiv1alpha1.DNSEndpoint{
				TypeMeta: metav1.TypeMeta{
					APIVersion: apiv1alpha1.GroupVersion.String(),
					Kind:       apiv1alpha1.DNSEndpointKind,
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:        "test",
					Namespace:   ti.objectNamespace,
					Annotations: ti.annotations,
					Labels:      ti.labels,
					Generation:  1,
				},
				Spec: apiv1alpha1.DNSEndpointSpec{
					Endpoints: ti.endpoints,
				},
			}

			fakeClient := newFakeClientBuilder(scheme, dnsEndpoint).
				WithStatusSubresource(&apiv1alpha1.DNSEndpoint{}).
				Build()

			cs := helperCrdSourceBuilder(t, ti.namespace, ti.annotationFilter, labelSelector, fakeClient, *dnsEndpoint)

			receivedEndpoints, err := cs.Endpoints(t.Context())
			if ti.expectError {
				require.Errorf(t, err, "Received err %v", err)
			} else {
				require.NoErrorf(t, err, "Received err %v", err)
			}

			// Sync the updated object from fake client back to the informer
			// This simulates what happens in real usage where the informer watches the API server
			if err == nil && ti.expectEndpoints {
				var updated apiv1alpha1.DNSEndpoint
				if getErr := fakeClient.Get(t.Context(), client.ObjectKey{Name: dnsEndpoint.Name, Namespace: dnsEndpoint.Namespace}, &updated); getErr == nil {
					cs.getWatcher().Modify(&updated)
				}
			}

			if len(receivedEndpoints) == 0 && !ti.expectEndpoints {
				return
			}

			// Validate the ObservedGeneration was updated from the informer's cache
			if err == nil && ti.expectEndpoints {
				indexer := cs.getInformer().GetIndexer()
				for _, obj := range indexer.List() {
					dnsEndpoint, ok := obj.(*apiv1alpha1.DNSEndpoint)
					if !ok || dnsEndpoint.Namespace != ti.objectNamespace {
						continue
					}
					if dnsEndpoint.Status.ObservedGeneration != dnsEndpoint.Generation {
						require.Failf(t, "ObservedGeneration mismatch", "ObservedGeneration <%v> is not equal to Generation <%v>", dnsEndpoint.Status.ObservedGeneration, dnsEndpoint.Generation)
					}
				}
			}

			// Validate received endpoints against expected endpoints.
			validateEndpoints(t, receivedEndpoints, ti.endpoints)

			for _, e := range receivedEndpoints {
				// TODO: at the moment not all sources apply ResourceLabelKey
				require.GreaterOrEqual(t, len(e.Labels), 1, "endpoint must have at least one label")
				require.Contains(t, e.Labels, endpoint.ResourceLabelKey, "endpoint must include the ResourceLabelKey label")
			}
		})
	}
}

func TestCRDSource_Endpoints_SkipsKeyOnGetByKeyError(t *testing.T) {
	// Build an indexer whose IndexWithSelectors function accepts any object
	// regardless of type.  This lets us store an *unstructured.Unstructured
	// under the index, which will cause GetByKey[*apiv1alpha1.DNSEndpoint] to
	// fail with a type-assertion error.
	indexer := cache.NewIndexer(
		cache.MetaNamespaceKeyFunc,
		cache.Indexers{
			informers.IndexWithSelectors: func(obj any) ([]string, error) {
				key, err := cache.MetaNamespaceKeyFunc(obj)
				if err != nil {
					return nil, err
				}
				return []string{key}, nil
			},
		},
	)

	wrongType := &unstructured.Unstructured{}
	wrongType.SetName("test")
	wrongType.SetNamespace("default")
	require.NoError(t, indexer.Add(wrongType))

	cs := &crdSource{indexer: indexer}
	endpoints, err := cs.Endpoints(t.Context())
	require.NoError(t, err)
	require.Empty(t, endpoints)
}

func TestCRDSource_NoInformer(t *testing.T) {
	cs := &crdSource{informer: nil}
	called := false

	cs.AddEventHandler(context.Background(), func() { called = true })
	require.False(t, called, "handler must not be called when informer is nil")
}

func TestCRDSource_AddEventHandler_Add(t *testing.T) {
	ctx := t.Context()
	watcher, cs := helperCreateWatcherWithInformer(t)

	var counter atomic.Int32
	cs.AddEventHandler(ctx, func() {
		counter.Add(1)
	})

	obj := &unstructured.Unstructured{}
	obj.SetName("test")

	watcher.Add(obj)

	require.Eventually(t, func() bool {
		return counter.Load() == 1
	}, 2*time.Second, 10*time.Millisecond)
}

func TestCRDSource_AddEventHandler_Update(t *testing.T) {
	ctx := t.Context()
	watcher, cs := helperCreateWatcherWithInformer(t)

	var counter atomic.Int32
	cs.AddEventHandler(ctx, func() {
		counter.Add(1)
	})

	obj := unstructured.Unstructured{}
	obj.SetName("test")
	obj.SetNamespace("default")
	obj.SetUID("9be5b64e-3ee9-11f0-88ee-1eb95c6fd730")

	watcher.Add(&obj)

	require.Eventually(t, func() bool {
		return len(watcher.Items) == 1
	}, 2*time.Second, 10*time.Millisecond)

	modified := obj.DeepCopy()
	modified.SetLabels(map[string]string{"new-label": "this"})
	watcher.Modify(modified)

	require.Eventually(t, func() bool {
		return len(watcher.Items) == 1
	}, 2*time.Second, 10*time.Millisecond)

	require.Eventually(t, func() bool {
		return counter.Load() == 2
	}, 2*time.Second, 10*time.Millisecond)
}

func TestCRDSource_AddEventHandler_Delete(t *testing.T) {
	ctx := t.Context()
	watcher, cs := helperCreateWatcherWithInformer(t)

	var counter atomic.Int32
	cs.AddEventHandler(ctx, func() {
		counter.Add(1)
	})

	obj := &unstructured.Unstructured{}
	obj.SetName("test")

	watcher.Delete(obj)

	require.Eventually(t, func() bool {
		return counter.Load() == 1
	}, 2*time.Second, 10*time.Millisecond)
}

func TestDNSEndpointsWithSetResourceLabels(t *testing.T) {
	fixtures := buildDnsEndpointTestFixtures([]fixtureRecord{
		{recordType: endpoint.RecordTypeA, count: 3, namespace: "target-ns"},
		{recordType: endpoint.RecordTypeCNAME, count: 2, namespace: "target-ns"},
		{recordType: endpoint.RecordTypeNS, count: 7, namespace: "other-ns"},
		{recordType: endpoint.RecordTypeNAPTR, count: 1, namespace: "other-ns"},
	})

	cs := helperCrdSourceBuilder(t, "target-ns", "", labels.Everything(), newFakeClientFromFixtures(t, fixtures), fixtures...)

	res, err := cs.Endpoints(t.Context())
	require.NoError(t, err)

	// Only endpoints from target-ns (3 A + 2 CNAME = 5)
	require.Len(t, res, 5)
	for _, ep := range res {
		require.Contains(t, ep.Labels, endpoint.ResourceLabelKey)
	}
}

func TestDNSEndpointsWithLabelFilter(t *testing.T) {
	// 16 total: alpha(3A+2CNAME) beta(4A+1CNAME) gamma(2A+3NS) delta(1A)
	fixtures := buildDnsEndpointTestFixtures([]fixtureRecord{
		{recordType: endpoint.RecordTypeA, count: 3, namespace: "alpha",
			labels: map[string]string{"app": "web", "env": "prod", "team": "platform"}},
		{recordType: endpoint.RecordTypeCNAME, count: 2, namespace: "alpha",
			labels: map[string]string{"app": "api", "env": "prod", "team": "platform"}},
		{recordType: endpoint.RecordTypeA, count: 4, namespace: "beta",
			labels: map[string]string{"app": "web", "env": "staging", "team": "infra"}},
		{recordType: endpoint.RecordTypeCNAME, count: 1, namespace: "beta",
			labels: map[string]string{"app": "worker", "env": "staging", "team": "infra"}},
		{recordType: endpoint.RecordTypeA, count: 2, namespace: "gamma",
			labels: map[string]string{"app": "web", "env": "dev", "team": "platform", "product": "dns"}},
		{recordType: endpoint.RecordTypeNS, count: 3, namespace: "gamma",
			labels: map[string]string{"app": "api", "env": "dev", "team": "infra", "product": "cdn"}},
		{recordType: endpoint.RecordTypeA, count: 1, namespace: "delta",
			labels: map[string]string{"app": "worker", "env": "prod", "team": "platform", "product": "dns"}},
	})
	fakeClient := newFakeClientFromFixtures(t, fixtures)

	for _, ti := range []struct {
		title         string
		namespace     string
		labelFilter   string
		expectedCount int
	}{
		{"app=web all namespaces", "", "app=web", 9},
		{"env=prod all namespaces", "", "env=prod", 6},
		{"team=infra all namespaces", "", "team=infra", 8},
		{"product=dns all namespaces", "", "product=dns", 3},
		{"combined app=web,env=prod", "", "app=web,env=prod", 3},
		{"combined team=platform,env=dev", "", "team=platform,env=dev", 2},
		{"two full and one key-only", "", "team=platform,env=prod,product", 1},
		{"no match", "", "app=nonexistent", 0},
		{"app=web scoped to beta", "beta", "app=web", 4},
		{"env=prod scoped to alpha", "alpha", "env=prod", 5},
		{"team=platform scoped to gamma", "gamma", "team=platform", 2},
		{"env=staging scoped to beta", "beta", "env=staging", 5},
		{"app=worker scoped to delta", "delta", "app=worker", 1},
		{"no match scoped to delta", "delta", "app=api", 0},
	} {
		t.Run(ti.title, func(t *testing.T) {
			labelSelector, err := labels.Parse(ti.labelFilter)
			require.NoError(t, err)

			cs := helperCrdSourceBuilder(t, ti.namespace, "", labelSelector, fakeClient, fixtures...)

			res, err := cs.Endpoints(t.Context())
			require.NoError(t, err)
			require.Len(t, res, ti.expectedCount)
		})
	}
}

func TestDNSEndpointsWithAnnotationFilter(t *testing.T) {
	// 16 total: alpha(3A+2CNAME) beta(4A+1CNAME) gamma(2A+3NS) delta(1A)
	fixtures := buildDnsEndpointTestFixtures([]fixtureRecord{
		{recordType: endpoint.RecordTypeA, count: 3, namespace: "alpha",
			annotations: map[string]string{"external-dns/owner": "team-a"}},
		{recordType: endpoint.RecordTypeCNAME, count: 2, namespace: "alpha",
			annotations: map[string]string{"external-dns/owner": "team-a"}},
		{recordType: endpoint.RecordTypeA, count: 4, namespace: "beta",
			annotations: map[string]string{"external-dns/owner": "team-b"}},
		{recordType: endpoint.RecordTypeCNAME, count: 1, namespace: "beta",
			annotations: map[string]string{"external-dns/owner": "team-b", "external-dns/zone": "private"}},
		{recordType: endpoint.RecordTypeA, count: 2, namespace: "gamma",
			annotations: map[string]string{"external-dns/owner": "team-a", "external-dns/zone": "private"}},
		{recordType: endpoint.RecordTypeNS, count: 3, namespace: "gamma",
			annotations: map[string]string{"external-dns/owner": "team-b", "external-dns/zone": "public"}},
		{recordType: endpoint.RecordTypeA, count: 1, namespace: "delta",
			annotations: map[string]string{"external-dns/owner": "team-a", "external-dns/zone": "public"}},
	})
	fakeClient := newFakeClientFromFixtures(t, fixtures)

	for _, ti := range []struct {
		title            string
		namespace        string
		annotationFilter string
		expectedCount    int
	}{
		{"owner=team-a all namespaces", "", "external-dns/owner=team-a", 8},
		{"owner=team-b all namespaces", "", "external-dns/owner=team-b", 8},
		{"zone=private all namespaces", "", "external-dns/zone=private", 3},
		{"zone=public all namespaces", "", "external-dns/zone=public", 4},
		{"combined owner=team-b,zone=public", "", "external-dns/owner=team-b,external-dns/zone=public", 3},
		{"two full and one key-only", "", "external-dns/owner=team-a,external-dns/zone", 3},
		{"no match", "", "external-dns/owner=nobody", 0},
		{"owner=team-a scoped to alpha", "alpha", "external-dns/owner=team-a", 5},
		{"owner=team-b scoped to beta", "beta", "external-dns/owner=team-b", 5},
		{"zone=private scoped to gamma", "gamma", "external-dns/zone=private", 2},
		{"owner=team-a scoped to delta", "delta", "external-dns/owner=team-a", 1},
		{"no match scoped to gamma", "gamma", "external-dns/owner=nobody", 0},
	} {
		t.Run(ti.title, func(t *testing.T) {
			cs := helperCrdSourceBuilder(t, ti.namespace, ti.annotationFilter, labels.Everything(), fakeClient, fixtures...)

			res, err := cs.Endpoints(t.Context())
			require.NoError(t, err)
			require.Len(t, res, ti.expectedCount)
		})
	}
}

func helperCreateWatcherWithInformer(t *testing.T) (*cachetesting.FakeControllerSource, *crdSource) {
	t.Helper()
	ctx := t.Context()

	watcher := cachetesting.NewFakeControllerSource()

	informer := cache.NewSharedIndexInformer(watcher, &unstructured.Unstructured{}, 0, cache.Indexers{})

	go informer.RunWithContext(ctx)

	require.Eventually(t, func() bool {
		return cache.WaitForCacheSync(ctx.Done(), informer.HasSynced)
	}, 2*time.Second, 10*time.Millisecond)

	cs := &crdSource{
		informer: informer,
		indexer:  informer.GetIndexer(),
	}

	return watcher, cs
}

// helperCrdSourceBuilder creates a crdSource with a test informer populated with DNSEndpoint objects.
// Returns a crdSourceTest wrapper that provides getInformer() and getWatcher() methods for tests.
func helperCrdSourceBuilder(t *testing.T, namespace, annotationFilter string, labelSelector labels.Selector, fakeClient client.Client, dnsEndpoints ...apiv1alpha1.DNSEndpoint) *crdSourceTest {
	t.Helper()
	ctx := t.Context()

	watcher := cachetesting.NewFakeControllerSource()
	informer := cache.NewSharedIndexInformer(
		watcher,
		&apiv1alpha1.DNSEndpoint{},
		0,
		cache.Indexers{},
	)

	// Add objects before starting so they appear in the initial LIST.
	for i := range dnsEndpoints {
		watcher.Add(&dnsEndpoints[i])
	}

	cs, err := newCRDSource(informer, fakeClient, annotationFilter, labelSelector, namespace)
	require.NoError(t, err)

	go informer.RunWithContext(ctx)
	require.Eventually(t, func() bool {
		return cache.WaitForCacheSync(ctx.Done(), informer.HasSynced)
	}, 2*time.Second, 10*time.Millisecond)

	return &crdSourceTest{
		crdSource:    cs,
		testInformer: informer,
		testWatcher:  watcher,
	}
}

func newFakeClientFromFixtures(t *testing.T, fixtures []apiv1alpha1.DNSEndpoint) client.Client {
	t.Helper()
	scheme := runtime.NewScheme()
	require.NoError(t, apiv1alpha1.AddToScheme(scheme))
	objects := make([]runtime.Object, len(fixtures))
	for i := range fixtures {
		objects[i] = &fixtures[i]
	}
	return newFakeClientBuilder(scheme, objects...).
		WithStatusSubresource(&apiv1alpha1.DNSEndpoint{}).
		Build()
}

type fixtureRecord struct {
	recordType  string
	count       int
	namespace   string
	labels      map[string]string
	annotations map[string]string
}

// buildDnsEndpointTestFixtures generates DNSEndpoint CRDs from a list of fixture records.
// Each record specifies the type, count, namespace, labels and annotations.
func buildDnsEndpointTestFixtures(records []fixtureRecord) []apiv1alpha1.DNSEndpoint {
	var result []apiv1alpha1.DNSEndpoint
	idx := 0
	for _, r := range records {
		for range r.count {
			result = append(result, apiv1alpha1.DNSEndpoint{
				ObjectMeta: metav1.ObjectMeta{
					Name:        fmt.Sprintf("dnsendpoint-%d", idx),
					Namespace:   r.namespace,
					Labels:      r.labels,
					Annotations: r.annotations,
				},
				Spec: apiv1alpha1.DNSEndpointSpec{
					Endpoints: []*endpoint.Endpoint{
						{
							DNSName:    strings.ToLower(fmt.Sprintf("%s-%d.example.com", r.recordType, idx)),
							RecordType: r.recordType,
							Targets:    endpoint.Targets{fmt.Sprintf("192.0.2.%d", idx)},
							RecordTTL:  300,
						},
					},
				},
			})
			idx++
		}
	}
	rand.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})
	return result
}
