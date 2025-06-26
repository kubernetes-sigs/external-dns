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
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
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
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/rest/fake"
	"k8s.io/client-go/tools/cache"
	cachetesting "k8s.io/client-go/tools/cache/testing"
	apiv1alpha1 "sigs.k8s.io/external-dns/apis/v1alpha1"
	"sigs.k8s.io/external-dns/endpoint"
)

type CRDSuite struct {
	suite.Suite
}

func (suite *CRDSuite) SetupTest() {
}

func defaultHeader() http.Header {
	header := http.Header{}
	header.Set("Content-Type", runtime.ContentTypeJSON)
	return header
}

func objBody(codec runtime.Encoder, obj runtime.Object) io.ReadCloser {
	return io.NopCloser(bytes.NewReader([]byte(runtime.EncodeOrDie(codec, obj))))
}

func fakeRESTClient(endpoints []*endpoint.Endpoint, apiVersion, kind, namespace, name string, annotations map[string]string, labels map[string]string, _ *testing.T) rest.Interface {
	groupVersion, _ := schema.ParseGroupVersion(apiVersion)
	scheme := runtime.NewScheme()
	_ = addKnownTypes(scheme, groupVersion)

	dnsEndpointList := apiv1alpha1.DNSEndpointList{}
	dnsEndpoint := &apiv1alpha1.DNSEndpoint{
		TypeMeta: metav1.TypeMeta{
			APIVersion: apiVersion,
			Kind:       kind,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        name,
			Namespace:   namespace,
			Annotations: annotations,
			Labels:      labels,
			Generation:  1,
		},
		Spec: apiv1alpha1.DNSEndpointSpec{
			Endpoints: endpoints,
		},
	}

	codecFactory := serializer.WithoutConversionCodecFactory{
		CodecFactory: serializer.NewCodecFactory(scheme),
	}
	client := &fake.RESTClient{
		GroupVersion:         groupVersion,
		VersionedAPIPath:     "/apis/" + apiVersion,
		NegotiatedSerializer: codecFactory,
		Client: fake.CreateHTTPClient(func(req *http.Request) (*http.Response, error) {
			codec := codecFactory.LegacyCodec(groupVersion)
			switch p, m := req.URL.Path, req.Method; {
			case p == "/apis/"+apiVersion+"/"+strings.ToLower(kind)+"s" && m == http.MethodGet:
				fallthrough
			case p == "/apis/"+apiVersion+"/namespaces/"+namespace+"/"+strings.ToLower(kind)+"s" && m == http.MethodGet:
				dnsEndpointList.Items = dnsEndpointList.Items[:0]
				dnsEndpointList.Items = append(dnsEndpointList.Items, *dnsEndpoint)
				return &http.Response{StatusCode: http.StatusOK, Header: defaultHeader(), Body: objBody(codec, &dnsEndpointList)}, nil
			case strings.HasPrefix(p, "/apis/"+apiVersion+"/namespaces/") && strings.HasSuffix(p, strings.ToLower(kind)+"s") && m == http.MethodGet:
				return &http.Response{StatusCode: http.StatusOK, Header: defaultHeader(), Body: objBody(codec, &dnsEndpointList)}, nil
			case p == "/apis/"+apiVersion+"/namespaces/"+namespace+"/"+strings.ToLower(kind)+"s/"+name+"/status" && m == http.MethodPut:
				decoder := json.NewDecoder(req.Body)

				var body apiv1alpha1.DNSEndpoint
				err := decoder.Decode(&body)
				if err != nil {
					return nil, err
				}
				dnsEndpoint.Status.ObservedGeneration = body.Status.ObservedGeneration
				return &http.Response{StatusCode: http.StatusOK, Header: defaultHeader(), Body: objBody(codec, dnsEndpoint)}, nil
			default:
				return nil, fmt.Errorf("unexpected request: %#v\n%#v", req.URL, req)
			}
		}),
	}

	return client
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
func testCRDSourceEndpoints(t *testing.T) {
	for _, ti := range []struct {
		title                string
		registeredNamespace  string
		namespace            string
		registeredAPIVersion string
		apiVersion           string
		registeredKind       string
		kind                 string
		endpoints            []*endpoint.Endpoint
		expectEndpoints      bool
		expectError          bool
		annotationFilter     string
		labelFilter          string
		annotations          map[string]string
		labels               map[string]string
	}{
		{
			title:                "invalid crd api version",
			registeredAPIVersion: "test.k8s.io/v1alpha1",
			apiVersion:           "blah.k8s.io/v1alpha1",
			registeredKind:       "DNSEndpoint",
			kind:                 "DNSEndpoint",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "abc.example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  180,
				},
			},
			expectEndpoints: false,
			expectError:     true,
		},
		{
			title:                "invalid crd kind",
			registeredAPIVersion: "test.k8s.io/v1alpha1",
			apiVersion:           "test.k8s.io/v1alpha1",
			registeredKind:       "DNSEndpoint",
			kind:                 "JustEndpoint",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "abc.example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  180,
				},
			},
			expectEndpoints: false,
			expectError:     true,
		},
		{
			title:                "endpoints within a specific namespace",
			registeredAPIVersion: "test.k8s.io/v1alpha1",
			apiVersion:           "test.k8s.io/v1alpha1",
			registeredKind:       "DNSEndpoint",
			kind:                 "DNSEndpoint",
			namespace:            "foo",
			registeredNamespace:  "foo",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "abc.example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  180,
				},
			},
			expectEndpoints: true,
			expectError:     false,
		},
		{
			title:                "no endpoints within a specific namespace",
			registeredAPIVersion: "test.k8s.io/v1alpha1",
			apiVersion:           "test.k8s.io/v1alpha1",
			registeredKind:       "DNSEndpoint",
			kind:                 "DNSEndpoint",
			namespace:            "foo",
			registeredNamespace:  "bar",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "abc.example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  180,
				},
			},
			expectEndpoints: false,
			expectError:     false,
		},
		{
			title:                "valid crd with no targets (relies on default-targets)",
			registeredAPIVersion: "test.k8s.io/v1alpha1",
			apiVersion:           "test.k8s.io/v1alpha1",
			registeredKind:       "DNSEndpoint",
			kind:                 "DNSEndpoint",
			namespace:            "foo",
			registeredNamespace:  "foo",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "no-targets.example.org",
					Targets:    endpoint.Targets{},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  180,
				},
			},
			expectEndpoints: true,
			expectError:     false,
		},
		{
			title:                "valid crd gvk with single endpoint",
			registeredAPIVersion: "test.k8s.io/v1alpha1",
			apiVersion:           "test.k8s.io/v1alpha1",
			registeredKind:       "DNSEndpoint",
			kind:                 "DNSEndpoint",
			namespace:            "foo",
			registeredNamespace:  "foo",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "abc.example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  180,
				},
			},
			expectEndpoints: true,
			expectError:     false,
		},
		{
			title:                "valid crd gvk with multiple endpoints",
			registeredAPIVersion: "test.k8s.io/v1alpha1",
			apiVersion:           "test.k8s.io/v1alpha1",
			registeredKind:       "DNSEndpoint",
			kind:                 "DNSEndpoint",
			namespace:            "foo",
			registeredNamespace:  "foo",
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
			expectError:     false,
		},
		{
			title:                "valid crd gvk with annotation and non matching annotation filter",
			registeredAPIVersion: "test.k8s.io/v1alpha1",
			apiVersion:           "test.k8s.io/v1alpha1",
			registeredKind:       "DNSEndpoint",
			kind:                 "DNSEndpoint",
			namespace:            "foo",
			registeredNamespace:  "foo",
			annotations:          map[string]string{"test": "that"},
			annotationFilter:     "test=filter_something_else",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "abc.example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  180,
				},
			},
			expectEndpoints: false,
			expectError:     false,
		},
		{
			title:                "valid crd gvk with annotation and matching annotation filter",
			registeredAPIVersion: "test.k8s.io/v1alpha1",
			apiVersion:           "test.k8s.io/v1alpha1",
			registeredKind:       "DNSEndpoint",
			kind:                 "DNSEndpoint",
			namespace:            "foo",
			registeredNamespace:  "foo",
			annotations:          map[string]string{"test": "that"},
			annotationFilter:     "test=that",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "abc.example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  180,
				},
			},
			expectEndpoints: true,
			expectError:     false,
		},
		{
			title:                "valid crd gvk with label and non matching label filter",
			registeredAPIVersion: "test.k8s.io/v1alpha1",
			apiVersion:           "test.k8s.io/v1alpha1",
			registeredKind:       "DNSEndpoint",
			kind:                 "DNSEndpoint",
			namespace:            "foo",
			registeredNamespace:  "foo",
			labels:               map[string]string{"test": "that"},
			labelFilter:          "test=filter_something_else",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "abc.example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  180,
				},
			},
			expectEndpoints: false,
			expectError:     false,
		},
		{
			title:                "valid crd gvk with label and matching label filter",
			registeredAPIVersion: "test.k8s.io/v1alpha1",
			apiVersion:           "test.k8s.io/v1alpha1",
			registeredKind:       "DNSEndpoint",
			kind:                 "DNSEndpoint",
			namespace:            "foo",
			registeredNamespace:  "foo",
			labels:               map[string]string{"test": "that"},
			labelFilter:          "test=that",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "abc.example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  180,
				},
			},
			expectEndpoints: true,
			expectError:     false,
		},
		{
			title:                "Create NS record",
			registeredAPIVersion: "test.k8s.io/v1alpha1",
			apiVersion:           "test.k8s.io/v1alpha1",
			registeredKind:       "DNSEndpoint",
			kind:                 "DNSEndpoint",
			namespace:            "foo",
			registeredNamespace:  "foo",
			labels:               map[string]string{"test": "that"},
			labelFilter:          "test=that",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "abc.example.org",
					Targets:    endpoint.Targets{"ns1.k8s.io", "ns2.k8s.io"},
					RecordType: endpoint.RecordTypeNS,
					RecordTTL:  180,
				},
			},
			expectEndpoints: true,
			expectError:     false,
		},
		{
			title:                "Create SRV record",
			registeredAPIVersion: "test.k8s.io/v1alpha1",
			apiVersion:           "test.k8s.io/v1alpha1",
			registeredKind:       "DNSEndpoint",
			kind:                 "DNSEndpoint",
			namespace:            "foo",
			registeredNamespace:  "foo",
			labels:               map[string]string{"test": "that"},
			labelFilter:          "test=that",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "_svc._tcp.example.org",
					Targets:    endpoint.Targets{"0 0 80 abc.example.org", "0 0 80 def.example.org"},
					RecordType: endpoint.RecordTypeSRV,
					RecordTTL:  180,
				},
			},
			expectEndpoints: true,
			expectError:     false,
		},
		{
			title:                "Create NAPTR record",
			registeredAPIVersion: "test.k8s.io/v1alpha1",
			apiVersion:           "test.k8s.io/v1alpha1",
			registeredKind:       "DNSEndpoint",
			kind:                 "DNSEndpoint",
			namespace:            "foo",
			registeredNamespace:  "foo",
			labels:               map[string]string{"test": "that"},
			labelFilter:          "test=that",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "example.org",
					Targets:    endpoint.Targets{`100 10 "S" "SIP+D2U" "!^.*$!sip:customer-service@example.org!" _sip._udp.example.org.`, `102 10 "S" "SIP+D2T" "!^.*$!sip:customer-service@example.org!" _sip._tcp.example.org.`},
					RecordType: endpoint.RecordTypeNAPTR,
					RecordTTL:  180,
				},
			},
			expectEndpoints: true,
			expectError:     false,
		},
		{
			title:                "illegal target CNAME",
			registeredAPIVersion: "test.k8s.io/v1alpha1",
			apiVersion:           "test.k8s.io/v1alpha1",
			registeredKind:       "DNSEndpoint",
			kind:                 "DNSEndpoint",
			namespace:            "foo",
			registeredNamespace:  "foo",
			labels:               map[string]string{"test": "that"},
			labelFilter:          "test=that",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "example.org",
					Targets:    endpoint.Targets{"foo.example.org."},
					RecordType: endpoint.RecordTypeCNAME,
					RecordTTL:  180,
				},
			},
			expectEndpoints: false,
			expectError:     false,
		},
		{
			title:                "illegal target NAPTR",
			registeredAPIVersion: "test.k8s.io/v1alpha1",
			apiVersion:           "test.k8s.io/v1alpha1",
			registeredKind:       "DNSEndpoint",
			kind:                 "DNSEndpoint",
			namespace:            "foo",
			registeredNamespace:  "foo",
			labels:               map[string]string{"test": "that"},
			labelFilter:          "test=that",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "example.org",
					Targets:    endpoint.Targets{`100 10 "S" "SIP+D2U" "!^.*$!sip:customer-service@example.org!" _sip._udp.example.org`, `102 10 "S" "SIP+D2T" "!^.*$!sip:customer-service@example.org!" _sip._tcp.example.org`},
					RecordType: endpoint.RecordTypeNAPTR,
					RecordTTL:  180,
				},
			},
			expectEndpoints: false,
			expectError:     false,
		},
	} {
		t.Run(ti.title, func(t *testing.T) {
			t.Parallel()

			restClient := fakeRESTClient(ti.endpoints, ti.registeredAPIVersion, ti.registeredKind, ti.registeredNamespace, "test", ti.annotations, ti.labels, t)
			groupVersion, err := schema.ParseGroupVersion(ti.apiVersion)
			require.NoError(t, err)
			require.NotNil(t, groupVersion)

			scheme := runtime.NewScheme()
			err = addKnownTypes(scheme, groupVersion)
			require.NoError(t, err)

			labelSelector, err := labels.Parse(ti.labelFilter)
			require.NoError(t, err)

			// At present, client-go's fake.RESTClient (used by crd_test.go) is known to cause race conditions when used
			// with informers: https://github.com/kubernetes/kubernetes/issues/95372
			// So don't start the informer during testing.
			cs, err := NewCRDSource(restClient, ti.namespace, ti.kind, ti.annotationFilter, labelSelector, scheme, false)
			require.NoError(t, err)

			receivedEndpoints, err := cs.Endpoints(t.Context())
			if ti.expectError {
				require.Errorf(t, err, "Received err %v", err)
			} else {
				require.NoErrorf(t, err, "Received err %v", err)
			}

			if len(receivedEndpoints) == 0 && !ti.expectEndpoints {
				return
			}

			if err == nil {
				validateCRDResource(t, cs, ti.expectError)
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
	}, time.Second, 10*time.Millisecond)
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
	}, time.Second, 10*time.Millisecond)

	modified := obj.DeepCopy()
	modified.SetLabels(map[string]string{"new-label": "this"})
	watcher.Modify(modified)

	require.Eventually(t, func() bool {
		return len(watcher.Items) == 1
	}, time.Second, 10*time.Millisecond)

	require.Eventually(t, func() bool {
		return counter.Load() == 2
	}, time.Second, 10*time.Millisecond)
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
	}, time.Second, 10*time.Millisecond)
}

func TestCRDSource_Watch(t *testing.T) {
	scheme := runtime.NewScheme()
	err := apiv1alpha1.AddToScheme(scheme)
	require.NoError(t, err)

	var watchCalled bool

	codecFactory := serializer.WithoutConversionCodecFactory{
		CodecFactory: serializer.NewCodecFactory(scheme),
	}

	versionApiPath := fmt.Sprintf("/apis/%s", apiv1alpha1.GroupVersion.String())

	client := &fake.RESTClient{
		GroupVersion:         apiv1alpha1.GroupVersion,
		VersionedAPIPath:     versionApiPath,
		NegotiatedSerializer: codecFactory,
		Client: fake.CreateHTTPClient(func(req *http.Request) (*http.Response, error) {
			if req.URL.Path == fmt.Sprintf("%s/namespaces/test-ns/dnsendpoints", versionApiPath) &&
				req.URL.Query().Get("watch") == "true" {
				watchCalled = true
				return &http.Response{
					StatusCode: http.StatusOK,
					Header:     make(http.Header),
				}, nil
			}
			t.Errorf("unexpected request: %v", req.URL)
			return nil, fmt.Errorf("unexpected request: %v", req.URL)
		}),
	}

	cs := &crdSource{
		crdClient:   client,
		namespace:   "test-ns",
		crdResource: "dnsendpoints",
		codec:       runtime.NewParameterCodec(scheme),
	}

	opts := &metav1.ListOptions{}

	_, err = cs.watch(t.Context(), opts)
	require.NoError(t, err)
	require.True(t, watchCalled)
	require.True(t, opts.Watch)
}

func validateCRDResource(t *testing.T, src Source, expectError bool) {
	t.Helper()
	cs := src.(*crdSource)
	result, err := cs.List(context.Background(), &metav1.ListOptions{})
	if expectError {
		require.Errorf(t, err, "Received err %v", err)
	} else {
		require.NoErrorf(t, err, "Received err %v", err)
	}

	for _, dnsEndpoint := range result.Items {
		if dnsEndpoint.Status.ObservedGeneration != dnsEndpoint.Generation {
			require.Errorf(t, err, "Unexpected CRD resource result: ObservedGenerations <%v> is not equal to Generation<%v>", dnsEndpoint.Status.ObservedGeneration, dnsEndpoint.Generation)
		}
	}
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
			require.Empty(t, ep.Labels, "endpoint not have labels set")
			require.NotContains(t, ep.Labels, endpoint.ResourceLabelKey, "endpoint must not include the ResourceLabelKey label")
		}
	}

	scheme := runtime.NewScheme()
	err := apiv1alpha1.AddToScheme(scheme)
	require.NoError(t, err)

	codecFactory := serializer.WithoutConversionCodecFactory{
		CodecFactory: serializer.NewCodecFactory(scheme),
	}

	client := &fake.RESTClient{
		GroupVersion:         apiv1alpha1.GroupVersion,
		VersionedAPIPath:     fmt.Sprintf("/apis/%s", apiv1alpha1.GroupVersion.String()),
		NegotiatedSerializer: codecFactory,
		Client: fake.CreateHTTPClient(func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Header:     make(http.Header),
				Body:       objBody(codecFactory.LegacyCodec(apiv1alpha1.GroupVersion), &crds),
			}, nil
		}),
	}

	cs := &crdSource{
		crdClient:     client,
		namespace:     "test-ns",
		crdResource:   "dnsendpoints",
		codec:         runtime.NewParameterCodec(scheme),
		labelSelector: labels.Everything(),
	}

	res, err := cs.Endpoints(t.Context())
	require.NoError(t, err)

	for _, ep := range res {
		require.Contains(t, ep.Labels, endpoint.ResourceLabelKey)
	}
}

func helperCreateWatcherWithInformer(t *testing.T) (*cachetesting.FakeControllerSource, crdSource) {
	t.Helper()
	ctx := t.Context()

	watcher := cachetesting.NewFakeControllerSource()

	informer := cache.NewSharedInformer(watcher, &unstructured.Unstructured{}, 0)

	go informer.RunWithContext(ctx)

	require.Eventually(t, func() bool {
		return cache.WaitForCacheSync(ctx.Done(), informer.HasSynced)
	}, time.Second, 10*time.Millisecond)

	cs := &crdSource{
		informer: &informer,
	}

	return watcher, *cs
}

// generateTestFixtureDNSEndpointsByType generates DNSEndpoint CRDs according to the provided counts per RecordType.
func generateTestFixtureDNSEndpointsByType(namespace string, typeCounts map[string]int) apiv1alpha1.DNSEndpointList {
	var result []apiv1alpha1.DNSEndpoint
	idx := 0
	for rt, count := range typeCounts {
		for i := 0; i < count; i++ {
			result = append(result, apiv1alpha1.DNSEndpoint{
				ObjectMeta: metav1.ObjectMeta{
					Name:      fmt.Sprintf("dnsendpoint-%s-%d", rt, idx),
					Namespace: namespace,
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
	// Shuffle the result to ensure randomness in the order.
	rand.New(rand.NewSource(time.Now().UnixNano()))
	rand.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})

	return apiv1alpha1.DNSEndpointList{
		Items: result,
	}
}
