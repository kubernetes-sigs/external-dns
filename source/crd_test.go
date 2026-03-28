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
	"k8s.io/apimachinery/pkg/runtime/schema"
	k8stypes "k8s.io/apimachinery/pkg/types"
	toolscache "k8s.io/client-go/tools/cache"
	cachetesting "k8s.io/client-go/tools/cache/testing"
	crcache "sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	log "github.com/sirupsen/logrus"

	apiv1alpha1 "sigs.k8s.io/external-dns/apis/v1alpha1"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	logtest "sigs.k8s.io/external-dns/internal/testutils/log"
	"sigs.k8s.io/external-dns/source/types"
)

var (
	_ Source = &crdSource{}
)

type CRDSuite struct {
	suite.Suite
}

func (suite *CRDSuite) SetupTest() {
}

// dnsEndpointByObj extracts the single ByObject entry for DNSEndpoint from
// cache options. The map key is a pointer so we cannot look it up directly —
// we iterate instead.
func dnsEndpointByObj(t *testing.T, opts crcache.Options) crcache.ByObject {
	t.Helper()
	for obj, bo := range opts.ByObject {
		if _, ok := obj.(*apiv1alpha1.DNSEndpoint); ok {
			return bo
		}
	}
	t.Fatal("no ByObject entry for DNSEndpoint")
	panic("unreachable")
}

func TestBuildCacheOptions(t *testing.T) {
	t.Run("all namespaces when namespace is empty", func(t *testing.T) {
		opts, err := buildCacheOptions("", nil, nil)
		require.NoError(t, err)
		byObj := dnsEndpointByObj(t, opts)
		require.Contains(t, byObj.Namespaces, "", "empty string key means NamespaceAll")
		require.Nil(t, byObj.Label, "no label filter expected")
	})

	t.Run("single namespace", func(t *testing.T) {
		opts, err := buildCacheOptions("my-ns", nil, nil)
		require.NoError(t, err)
		byObj := dnsEndpointByObj(t, opts)
		require.Contains(t, byObj.Namespaces, "my-ns")
		require.NotContains(t, byObj.Namespaces, "")
	})

	t.Run("label filter applied", func(t *testing.T) {
		sel := labels.SelectorFromSet(labels.Set{"app": "foo"})
		opts, err := buildCacheOptions("", sel, nil)
		require.NoError(t, err)
		byObj := dnsEndpointByObj(t, opts)
		require.NotNil(t, byObj.Label)
		require.True(t, byObj.Label.Matches(labels.Set{"app": "foo"}))
		require.False(t, byObj.Label.Matches(labels.Set{"app": "bar"}))
	})

	t.Run("empty label selector not applied", func(t *testing.T) {
		opts, err := buildCacheOptions("", labels.Everything(), nil)
		require.NoError(t, err)
		byObj := dnsEndpointByObj(t, opts)
		require.Nil(t, byObj.Label)
	})

	t.Run("transform keeps object matching label filter", func(t *testing.T) {
		sel := labels.SelectorFromSet(labels.Set{"app": "foo"})
		opts, err := buildCacheOptions("", sel, nil)
		require.NoError(t, err)
		byObj := dnsEndpointByObj(t, opts)

		obj := &apiv1alpha1.DNSEndpoint{ObjectMeta: metav1.ObjectMeta{Labels: map[string]string{"app": "foo"}}}
		got, err := byObj.Transform(obj)
		require.NoError(t, err)
		require.NotNil(t, got)
	})

	t.Run("transform drops object not matching label filter", func(t *testing.T) {
		sel := labels.SelectorFromSet(labels.Set{"app": "foo"})
		opts, err := buildCacheOptions("", sel, nil)
		require.NoError(t, err)
		byObj := dnsEndpointByObj(t, opts)

		obj := &apiv1alpha1.DNSEndpoint{ObjectMeta: metav1.ObjectMeta{Labels: map[string]string{"app": "bar"}}}
		got, err := byObj.Transform(obj)
		require.NoError(t, err)
		require.Nil(t, got)
	})

	t.Run("transform keeps object matching annotation filter", func(t *testing.T) {
		opts, err := buildCacheOptions("", nil, labels.SelectorFromSet(labels.Set{"env": "prod"}))
		require.NoError(t, err)
		byObj := dnsEndpointByObj(t, opts)

		obj := &apiv1alpha1.DNSEndpoint{ObjectMeta: metav1.ObjectMeta{Annotations: map[string]string{"env": "prod"}}}
		got, err := byObj.Transform(obj)
		require.NoError(t, err)
		require.NotNil(t, got)
	})

	t.Run("transform drops object not matching annotation filter", func(t *testing.T) {
		opts, err := buildCacheOptions("", nil, labels.SelectorFromSet(labels.Set{"env": "prod"}))
		require.NoError(t, err)
		byObj := dnsEndpointByObj(t, opts)

		obj := &apiv1alpha1.DNSEndpoint{ObjectMeta: metav1.ObjectMeta{Annotations: map[string]string{"env": "staging"}}}
		got, err := byObj.Transform(obj)
		require.NoError(t, err)
		require.Nil(t, got)
	})
}

func TestCRDSource(t *testing.T) {
	suite.Run(t, new(CRDSuite))
	t.Run("Interface", testCRDSourceImplementsSource)
	t.Run("Endpoints", testCRDSourceEndpoints)
}

// testCRDSourceImplementsSource tests that crdSource is a valid Source.
func testCRDSourceImplementsSource(t *testing.T) {
	require.Implements(t, (*Source)(nil), new(crdSource))
}

// testCRDSourceEndpoints tests various scenarios of using CRD source.
//
// Namespace and label filtering are handled by the controller-runtime cache via
// ByObject at construction time — not inside Endpoints().  Tests mirror this by
// only adding objects to the fake cache that the real cache would deliver:
// objects whose namespace and labels match the source configuration.
// Annotation filtering and target validation are performed inside Endpoints()
// and are tested with objects already present in the fake cache.
func testCRDSourceEndpoints(t *testing.T) {
	for _, ti := range []struct {
		title              string
		namespaceFilter          string
		objectNamespace    string
		endpoints          []*endpoint.Endpoint
		expectEndpoints    bool
		expectError        bool
		annotationSelector labels.Selector
		labelSelector      labels.Selector
		annotations        map[string]string
		labels             map[string]string
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
			title:              "valid crd gvk with annotation and non matching annotation filter",
			namespace:          "foo",
			objectNamespace:    "foo",
			annotations:        map[string]string{"test": "that"},
			annotationSelector: labels.SelectorFromSet(labels.Set{"test": "filter_something_else"}),
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
			title:              "valid crd gvk with annotation and matching annotation filter",
			namespace:          "foo",
			objectNamespace:    "foo",
			annotations:        map[string]string{"test": "that"},
			annotationSelector: labels.SelectorFromSet(labels.Set{"test": "that"}),
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
			labelSelector:   labels.SelectorFromSet(labels.Set{"test": "filter_something_else"}),
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
			labelSelector:   labels.SelectorFromSet(labels.Set{"test": "that"}),
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
			labelSelector:   labels.SelectorFromSet(labels.Set{"test": "that"}),
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
			labelSelector:   labels.SelectorFromSet(labels.Set{"test": "that"}),
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
			labelSelector:   labels.SelectorFromSet(labels.Set{"test": "that"}),
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
			title:           "CNAME target with trailing dot (RFC 1035 §5.1 absolute FQDN) is valid",
			namespace:       "foo",
			objectNamespace: "foo",
			labels:          map[string]string{"test": "that"},
			labelSelector:   labels.SelectorFromSet(labels.Set{"test": "that"}),
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
			title:           "CNAME target without trailing dot (relative name)",
			namespace:       "foo",
			objectNamespace: "foo",
			labels:          map[string]string{"test": "that"},
			labelSelector:   labels.SelectorFromSet(labels.Set{"test": "that"}),
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
			labelSelector:   labels.SelectorFromSet(labels.Set{"test": "that"}),
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
			labelSelector:   labels.SelectorFromSet(labels.Set{"test": "that"}),
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
			labelSelector:   labels.SelectorFromSet(labels.Set{"test": "that"}),
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
			labelSelector:   labels.SelectorFromSet(labels.Set{"test": "that"}),
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "example.org",
					Targets:    endpoint.Targets{"example.com."},
					RecordType: endpoint.RecordTypeMX,
					RecordTTL:  180,
				},
			},
			expectEndpoints: true,
		},
		{
			title:           "MX Record without trailing dot in target",
			namespace:       "foo",
			objectNamespace: "foo",
			labels:          map[string]string{"test": "that"},
			labelSelector:   labels.SelectorFromSet(labels.Set{"test": "that"}),
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "example.org",
					Targets:    endpoint.Targets{"example.com"},
					RecordType: endpoint.RecordTypeMX,
					RecordTTL:  180,
				},
			},
			expectEndpoints: true,
		},
		{
			title:           "provider-specific properties are passed through from DNSEndpoint spec",
			namespace:       "foo",
			objectNamespace: "foo",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "subdomain.example.org",
					Targets:    endpoint.Targets{"other.example.org"},
					RecordType: endpoint.RecordTypeCNAME,
					RecordTTL:  180,
					ProviderSpecific: endpoint.ProviderSpecific{
						{Name: "aws/failover", Value: "PRIMARY"},
						{Name: "aws/health-check-id", Value: "asdf1234-as12-as12-as12-asdf12345678"},
						{Name: "aws/evaluate-target-health", Value: "true"},
					},
					SetIdentifier: "some-unique-id",
				},
			},
			expectEndpoints: true,
		},
	} {
		t.Run(ti.title, func(t *testing.T) {
			t.Parallel()

			obj := &apiv1alpha1.DNSEndpoint{
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

			fakeCache := newFakeCRDCache(t, nil, fakeCRDCacheFilter{ti.namespace, ti.labelSelector, ti.annotationSelector}, obj)
			cs, err := newCrdSource(t.Context(), fakeCache, fakeCache.Client)
			require.NoError(t, err)

			receivedEndpoints, err := cs.Endpoints(t.Context())
			if ti.expectError {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			if len(receivedEndpoints) == 0 && !ti.expectEndpoints {
				return
			}

			if ti.expectEndpoints {
				validateCRDResource(t, fakeCache.Client, ti.objectNamespace, "test")
			}

			testutils.ValidateEndpoints(t, receivedEndpoints, ti.endpoints)

			for _, e := range receivedEndpoints {
				require.GreaterOrEqual(t, len(e.Labels), 1, "endpoint must have at least one label")
				require.Contains(t, e.Labels, endpoint.ResourceLabelKey, "endpoint must include the ResourceLabelKey label")
			}
		})
	}
}

func TestCRDSourceIllegalTargetWarnings(t *testing.T) {
	for _, ti := range []struct {
		title       string
		endpoints   []*endpoint.Endpoint
		wantWarning string
	}{
		{
			title: "A record with trailing dot warns with fix suggestion",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "example.org",
					Targets:    endpoint.Targets{"1.2.3.4."},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  180,
				},
			},
			wantWarning: `illegal target "1.2.3.4." for A record — use "1.2.3.4" not "1.2.3.4."`,
		},
		{
			title: "NAPTR record without trailing dot warns with fix suggestion",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "example.org",
					Targets:    endpoint.Targets{"_sip._udp.example.org"},
					RecordType: endpoint.RecordTypeNAPTR,
					RecordTTL:  180,
				},
			},
			wantWarning: `illegal target "_sip._udp.example.org" for NAPTR record — use "_sip._udp.example.org." not "_sip._udp.example.org"`,
		},
		{
			title: "CNAME with empty targets produces no warning",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "example.org",
					Targets:    endpoint.Targets{},
					RecordType: endpoint.RecordTypeCNAME,
					RecordTTL:  180,
				},
			},
			wantWarning: ``,
		},
	} {
		t.Run(ti.title, func(t *testing.T) {
			hook := logtest.LogsUnderTestWithLogLevel(log.WarnLevel, t)

			obj := &apiv1alpha1.DNSEndpoint{
				ObjectMeta: metav1.ObjectMeta{
					Name:       "test",
					Namespace:  "foo",
					Generation: 1,
				},
				Spec: apiv1alpha1.DNSEndpointSpec{Endpoints: ti.endpoints},
			}

			fakeCache := newFakeCRDCache(t, nil, fakeCRDCacheFilter{}, obj)
			cs, err := newCrdSource(t.Context(), fakeCache, fakeCache.Client)
			require.NoError(t, err)

			_, err = cs.Endpoints(t.Context())
			require.NoError(t, err)

			if ti.wantWarning == "" {
				require.Empty(t, hook.Entries, "expected no warnings to be logged")
			} else {
				logtest.TestHelperLogContainsWithLogLevel(ti.wantWarning, log.WarnLevel, hook, t)
			}
		})
	}
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

func validateCRDResource(t *testing.T, fakeClient client.Client, namespace, name string) {
	t.Helper()
	updated := &apiv1alpha1.DNSEndpoint{}
	err := fakeClient.Get(t.Context(), client.ObjectKey{Namespace: namespace, Name: name}, updated)
	require.NoError(t, err)
	require.Equal(t, updated.Generation, updated.Status.ObservedGeneration,
		"ObservedGeneration should be updated to match Generation after Endpoints() is called")
}

func TestDNSEndpointsWithSetResourceLabels(t *testing.T) {
	typeCounts := map[string]int{
		endpoint.RecordTypeA:     3,
		endpoint.RecordTypeCNAME: 2,
		endpoint.RecordTypeNS:    7,
		endpoint.RecordTypeNAPTR: 1,
	}

	crds := generateTestFixtureDNSEndpointsByType("test-ns", typeCounts)

	for _, crd := range crds.Items {
		for _, ep := range crd.Spec.Endpoints {
			require.Empty(t, ep.Labels, "endpoint should not have labels set")
			require.NotContains(t, ep.Labels, endpoint.ResourceLabelKey, "endpoint must not include the ResourceLabelKey label")
		}
	}

	fakeCache := newFakeCRDCache(t, nil, fakeCRDCacheFilter{}, dnsEndpointListToObjects(crds.Items)...)
	cs, err := newCrdSource(t.Context(), fakeCache, fakeCache.Client)
	require.NoError(t, err)

	res, err := cs.Endpoints(t.Context())
	require.NoError(t, err)

	for _, ep := range res {
		require.Contains(t, ep.Labels, endpoint.ResourceLabelKey)
	}
}

func TestProcessEndpoint_CRD_RefObjectExist(t *testing.T) {
	typeCounts := map[string]int{
		endpoint.RecordTypeA:    2,
		endpoint.RecordTypeAAAA: 3,
	}

	elements := generateTestFixtureDNSEndpointsByType("test-ns", typeCounts)

	fakeCache := newFakeCRDCache(t, nil, fakeCRDCacheFilter{}, dnsEndpointListToObjects(elements.Items)...)
	cs, err := newCrdSource(t.Context(), fakeCache, fakeCache.Client)
	require.NoError(t, err)

	endpoints, err := cs.Endpoints(t.Context())
	require.NoError(t, err)
	testutils.AssertEndpointsHaveRefObject(t, endpoints, types.CRD, len(elements.Items))
}

// helperCreateWatcherWithInformer creates a FakeControllerSource-backed informer,
// wires it into a fakeCRDCache, and returns a crdSource so tests can inject
// events and verify handler invocations.
//
// toolscache.NewSharedIndexInformer is used (rather than NewSharedInformer) because
// SharedIndexInformer satisfies the controller-runtime cache.Informer interface,
// which additionally requires AddIndexers.
func helperCreateWatcherWithInformer(t *testing.T) (*cachetesting.FakeControllerSource, *crdSource) {
	t.Helper()
	ctx := t.Context()

	watcher := cachetesting.NewFakeControllerSource()
	informer := toolscache.NewSharedIndexInformer(watcher, &unstructured.Unstructured{}, 0, toolscache.Indexers{})

	go informer.RunWithContext(ctx)
	require.Eventually(t, func() bool {
		return toolscache.WaitForCacheSync(ctx.Done(), informer.HasSynced)
	}, 2*time.Second, 10*time.Millisecond)

	fakeCache := newFakeCRDCache(t, informer, fakeCRDCacheFilter{})
	cs, err := newCrdSource(ctx, fakeCache, fakeCache.Client)
	require.NoError(t, err)

	return watcher, cs
}

// generateTestFixtureDNSEndpointsByType generates DNSEndpoint CRDs according to the provided counts per RecordType.
func generateTestFixtureDNSEndpointsByType(namespace string, typeCounts map[string]int) apiv1alpha1.DNSEndpointList {
	var result []apiv1alpha1.DNSEndpoint
	idx := 0
	for rt, count := range typeCounts {
		for range count {
			result = append(result, apiv1alpha1.DNSEndpoint{
				ObjectMeta: metav1.ObjectMeta{
					Name:      fmt.Sprintf("dnsendpoint-%s-%d", rt, idx),
					Namespace: namespace,
					UID:       k8stypes.UID(fmt.Sprintf("uid-%d", idx)),
				},
				Spec: apiv1alpha1.DNSEndpointSpec{
					Endpoints: []*endpoint.Endpoint{
						{
							DNSName:    strings.ToLower(fmt.Sprintf("%s-%d.example.com", rt, idx)),
							RecordType: rt,
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

	return apiv1alpha1.DNSEndpointList{
		Items: result,
	}
}

// fakeCRDCache implements crCache.Cache for unit tests.
// Reads are backed by a fake.Client so test fixtures are populated with
// WithObjects; cache lifecycle methods are no-ops that return immediately.
type fakeCRDCache struct {
	client.Client // provides Get + List from fake.Client
	informer      toolscache.SharedIndexInformer
}

func (f *fakeCRDCache) GetInformer(_ context.Context, _ client.Object, _ ...crcache.InformerGetOption) (crcache.Informer, error) {
	return f.informer, nil
}

func (f *fakeCRDCache) GetInformerForKind(_ context.Context, _ schema.GroupVersionKind, _ ...crcache.InformerGetOption) (crcache.Informer, error) {
	return f.informer, nil
}

func (*fakeCRDCache) RemoveInformer(_ context.Context, _ client.Object) error { return nil }
func (*fakeCRDCache) Start(_ context.Context) error                           { return nil }
func (*fakeCRDCache) WaitForCacheSync(_ context.Context) bool                 { return true }
func (*fakeCRDCache) IndexField(_ context.Context, _ client.Object, _ string, _ client.IndexerFunc) error {
	return nil
}

// fakeCRDCacheFilter holds the admission criteria applied by the real controller-runtime
// cache (namespace, label selector, annotation selector). Zero value means no filtering.
type fakeCRDCacheFilter struct {
	namespace          string
	labelSelector      labels.Selector
	annotationSelector labels.Selector
}

// newFakeCRDCache builds a test cache backed by the given objects.
// It applies the same admission filtering as the real cache built by buildCacheOptions:
// namespace scope, label selector, and annotation filter are all applied via the transform.
// When informer is nil a real SharedIndexInformer backed by a FakeControllerSource
// is created to satisfy newCrdSource's GetInformer call; it is not started.
func newFakeCRDCache(t *testing.T, informer toolscache.SharedIndexInformer, filter fakeCRDCacheFilter, objs ...client.Object) *fakeCRDCache {
	t.Helper()
	if informer == nil {
		informer = toolscache.NewSharedIndexInformer(
			cachetesting.NewFakeControllerSource(),
			&apiv1alpha1.DNSEndpoint{},
			0,
			toolscache.Indexers{},
		)
	}
	if len(objs) > 0 {
		cacheOpts, err := buildCacheOptions(filter.namespace, filter.labelSelector, filter.annotationSelector)
		require.NoError(t, err)
		byObj := dnsEndpointByObj(t, cacheOpts)
		var admitted []client.Object
		for _, obj := range objs {
			got, err := byObj.Transform(obj)
			require.NoError(t, err)
			if got != nil {
				admitted = append(admitted, obj)
			}
		}
		objs = admitted
	}
	fc := fake.NewClientBuilder().
		WithScheme(newCRDTestScheme(t)).
		WithStatusSubresource(&apiv1alpha1.DNSEndpoint{}).
		WithObjects(objs...).
		Build()
	return &fakeCRDCache{Client: fc, informer: informer}
}

// newCRDTestScheme returns a scheme with the external-dns API types registered.
func newCRDTestScheme(t *testing.T) *runtime.Scheme {
	t.Helper()
	scheme := runtime.NewScheme()
	require.NoError(t, apiv1alpha1.AddToScheme(scheme))
	return scheme
}

// dnsEndpointListToObjects converts a DNSEndpointList slice into a []client.Object slice.
func dnsEndpointListToObjects(items []apiv1alpha1.DNSEndpoint) []client.Object {
	objs := make([]client.Object, len(items))
	for i := range items {
		objs[i] = &items[i]
	}
	return objs
}
