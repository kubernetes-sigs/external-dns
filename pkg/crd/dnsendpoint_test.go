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
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/rest/fake"

	apiv1alpha1 "sigs.k8s.io/external-dns/apis/v1alpha1"
	"sigs.k8s.io/external-dns/endpoint"
)

func defaultHeader() http.Header {
	header := http.Header{}
	header.Set("Content-Type", runtime.ContentTypeJSON)
	return header
}

func objBody(codec runtime.Encoder, obj runtime.Object) io.ReadCloser {
	return io.NopCloser(bytes.NewReader([]byte(runtime.EncodeOrDie(codec, obj))))
}

func newTestScheme() *runtime.Scheme {
	scheme := runtime.NewScheme()
	_ = apiv1alpha1.AddToScheme(scheme)
	return scheme
}

func TestNewDNSEndpointClient(t *testing.T) {
	tests := []struct {
		name             string
		kind             string
		expectedResource string
	}{
		{
			name:             "standard DNSEndpoint kind",
			kind:             "DNSEndpoint",
			expectedResource: "dnsendpoints",
		},
		{
			name:             "custom kind",
			kind:             "CustomEndpoint",
			expectedResource: "customendpoints",
		},
		{
			name:             "already lowercase",
			kind:             "endpoint",
			expectedResource: "endpoints",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scheme := newTestScheme()
			codecFactory := serializer.NewCodecFactory(scheme)

			restClient := &fake.RESTClient{
				GroupVersion:         apiv1alpha1.GroupVersion,
				NegotiatedSerializer: codecFactory,
			}

			client := NewDNSEndpointClient(restClient, tt.kind)

			require.NotNil(t, client)
			impl, ok := client.(*dnsEndpointClient)
			require.True(t, ok)
			assert.Equal(t, tt.expectedResource, impl.resource)
		})
	}
}

func TestDNSEndpointClient_Get(t *testing.T) {
	ctx := context.Background()
	scheme := newTestScheme()
	codecFactory := serializer.WithoutConversionCodecFactory{
		CodecFactory: serializer.NewCodecFactory(scheme),
	}
	codec := codecFactory.LegacyCodec(apiv1alpha1.GroupVersion)

	expectedEndpoint := &apiv1alpha1.DNSEndpoint{
		TypeMeta: metav1.TypeMeta{
			APIVersion: apiv1alpha1.GroupVersion.String(),
			Kind:       apiv1alpha1.DNSEndpointKind,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-endpoint",
			Namespace: "test-ns",
		},
		Spec: apiv1alpha1.DNSEndpointSpec{
			Endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  300,
				},
			},
		},
	}

	restClient := &fake.RESTClient{
		GroupVersion:         apiv1alpha1.GroupVersion,
		NegotiatedSerializer: codecFactory,
		Client: fake.CreateHTTPClient(func(req *http.Request) (*http.Response, error) {
			if req.Method == http.MethodGet && req.URL.Path == "/namespaces/test-ns/dnsendpoints/test-endpoint" {
				return &http.Response{
					StatusCode: http.StatusOK,
					Header:     defaultHeader(),
					Body:       objBody(codec, expectedEndpoint),
				}, nil
			}
			return nil, fmt.Errorf("unexpected request: %s %s", req.Method, req.URL.Path)
		}),
	}

	client := NewDNSEndpointClient(restClient, apiv1alpha1.DNSEndpointKind)

	result, err := client.Get(ctx, "test-ns", "test-endpoint")

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, "test-endpoint", result.Name)
	assert.Equal(t, "test-ns", result.Namespace)
	require.Len(t, result.Spec.Endpoints, 1)
	assert.Equal(t, "example.org", result.Spec.Endpoints[0].DNSName)
}

func TestDNSEndpointClient_Get_NotFound(t *testing.T) {
	ctx := context.Background()
	scheme := newTestScheme()
	codecFactory := serializer.WithoutConversionCodecFactory{
		CodecFactory: serializer.NewCodecFactory(scheme),
	}

	restClient := &fake.RESTClient{
		GroupVersion:         apiv1alpha1.GroupVersion,
		NegotiatedSerializer: codecFactory,
		Client: fake.CreateHTTPClient(func(_ *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusNotFound,
				Header:     defaultHeader(),
				Body:       io.NopCloser(bytes.NewReader([]byte(`{"kind":"Status","apiVersion":"v1","status":"Failure","message":"not found","reason":"NotFound","code":404}`))),
			}, nil
		}),
	}

	client := NewDNSEndpointClient(restClient, apiv1alpha1.DNSEndpointKind)

	_, err := client.Get(ctx, "test-ns", "nonexistent")

	require.Error(t, err)
}

func TestDNSEndpointClient_List(t *testing.T) {
	ctx := context.Background()
	scheme := newTestScheme()
	codecFactory := serializer.WithoutConversionCodecFactory{
		CodecFactory: serializer.NewCodecFactory(scheme),
	}
	codec := codecFactory.LegacyCodec(apiv1alpha1.GroupVersion)

	expectedList := &apiv1alpha1.DNSEndpointList{
		TypeMeta: metav1.TypeMeta{
			APIVersion: apiv1alpha1.GroupVersion.String(),
			Kind:       "DNSEndpointList",
		},
		Items: []apiv1alpha1.DNSEndpoint{
			{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "endpoint-1",
					Namespace: "test-ns",
				},
				Spec: apiv1alpha1.DNSEndpointSpec{
					Endpoints: []*endpoint.Endpoint{
						{DNSName: "one.example.org", Targets: endpoint.Targets{"1.1.1.1"}},
					},
				},
			},
			{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "endpoint-2",
					Namespace: "test-ns",
				},
				Spec: apiv1alpha1.DNSEndpointSpec{
					Endpoints: []*endpoint.Endpoint{
						{DNSName: "two.example.org", Targets: endpoint.Targets{"2.2.2.2"}},
					},
				},
			},
		},
	}

	restClient := &fake.RESTClient{
		GroupVersion:         apiv1alpha1.GroupVersion,
		NegotiatedSerializer: codecFactory,
		Client: fake.CreateHTTPClient(func(req *http.Request) (*http.Response, error) {
			if req.Method == http.MethodGet && req.URL.Path == "/namespaces/test-ns/dnsendpoints" {
				return &http.Response{
					StatusCode: http.StatusOK,
					Header:     defaultHeader(),
					Body:       objBody(codec, expectedList),
				}, nil
			}
			return nil, fmt.Errorf("unexpected request: %s %s", req.Method, req.URL.Path)
		}),
	}

	client := NewDNSEndpointClient(restClient, apiv1alpha1.DNSEndpointKind)

	result, err := client.List(ctx, "test-ns", &metav1.ListOptions{})

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Len(t, result.Items, 2)
	assert.Equal(t, "endpoint-1", result.Items[0].Name)
	assert.Equal(t, "endpoint-2", result.Items[1].Name)
}

func TestDNSEndpointClient_List_AllNamespaces(t *testing.T) {
	ctx := context.Background()
	scheme := newTestScheme()
	codecFactory := serializer.WithoutConversionCodecFactory{
		CodecFactory: serializer.NewCodecFactory(scheme),
	}
	codec := codecFactory.LegacyCodec(apiv1alpha1.GroupVersion)

	expectedList := &apiv1alpha1.DNSEndpointList{
		Items: []apiv1alpha1.DNSEndpoint{
			{ObjectMeta: metav1.ObjectMeta{Name: "ep1", Namespace: "ns1"}},
			{ObjectMeta: metav1.ObjectMeta{Name: "ep2", Namespace: "ns2"}},
		},
	}

	restClient := &fake.RESTClient{
		GroupVersion:         apiv1alpha1.GroupVersion,
		NegotiatedSerializer: codecFactory,
		Client: fake.CreateHTTPClient(func(req *http.Request) (*http.Response, error) {
			// Empty namespace should query cluster-wide
			if req.Method == http.MethodGet && req.URL.Path == "/dnsendpoints" {
				return &http.Response{
					StatusCode: http.StatusOK,
					Header:     defaultHeader(),
					Body:       objBody(codec, expectedList),
				}, nil
			}
			return nil, fmt.Errorf("unexpected request: %s %s", req.Method, req.URL.Path)
		}),
	}

	client := NewDNSEndpointClient(restClient, apiv1alpha1.DNSEndpointKind)

	result, err := client.List(ctx, "", &metav1.ListOptions{})

	require.NoError(t, err)
	assert.Len(t, result.Items, 2)
}

func TestDNSEndpointClient_List_WithLabelSelector(t *testing.T) {
	ctx := context.Background()
	scheme := newTestScheme()
	codecFactory := serializer.WithoutConversionCodecFactory{
		CodecFactory: serializer.NewCodecFactory(scheme),
	}
	codec := codecFactory.LegacyCodec(apiv1alpha1.GroupVersion)

	expectedList := &apiv1alpha1.DNSEndpointList{
		Items: []apiv1alpha1.DNSEndpoint{
			{ObjectMeta: metav1.ObjectMeta{Name: "filtered-ep", Namespace: "test-ns"}},
		},
	}

	var capturedLabelSelector string
	restClient := &fake.RESTClient{
		GroupVersion:         apiv1alpha1.GroupVersion,
		NegotiatedSerializer: codecFactory,
		Client: fake.CreateHTTPClient(func(req *http.Request) (*http.Response, error) {
			capturedLabelSelector = req.URL.Query().Get("labelSelector")
			return &http.Response{
				StatusCode: http.StatusOK,
				Header:     defaultHeader(),
				Body:       objBody(codec, expectedList),
			}, nil
		}),
	}

	client := NewDNSEndpointClient(restClient, apiv1alpha1.DNSEndpointKind)

	_, err := client.List(ctx, "test-ns", &metav1.ListOptions{
		LabelSelector: "app=external-dns",
	})

	require.NoError(t, err)
	assert.Equal(t, "app=external-dns", capturedLabelSelector)
}

func TestDNSEndpointClient_UpdateStatus(t *testing.T) {
	ctx := context.Background()
	scheme := newTestScheme()
	codecFactory := serializer.WithoutConversionCodecFactory{
		CodecFactory: serializer.NewCodecFactory(scheme),
	}
	codec := codecFactory.LegacyCodec(apiv1alpha1.GroupVersion)

	inputEndpoint := &apiv1alpha1.DNSEndpoint{
		ObjectMeta: metav1.ObjectMeta{
			Name:       "test-endpoint",
			Namespace:  "test-ns",
			Generation: 2,
		},
		Status: apiv1alpha1.DNSEndpointStatus{
			ObservedGeneration: 2,
		},
	}

	var capturedBody *apiv1alpha1.DNSEndpoint
	restClient := &fake.RESTClient{
		GroupVersion:         apiv1alpha1.GroupVersion,
		NegotiatedSerializer: codecFactory,
		Client: fake.CreateHTTPClient(func(req *http.Request) (*http.Response, error) {
			if req.Method == http.MethodPut && req.URL.Path == "/namespaces/test-ns/dnsendpoints/test-endpoint/status" {
				decoder := json.NewDecoder(req.Body)
				capturedBody = &apiv1alpha1.DNSEndpoint{}
				if err := decoder.Decode(capturedBody); err != nil {
					return nil, err
				}
				return &http.Response{
					StatusCode: http.StatusOK,
					Header:     defaultHeader(),
					Body:       objBody(codec, inputEndpoint),
				}, nil
			}
			return nil, fmt.Errorf("unexpected request: %s %s", req.Method, req.URL.Path)
		}),
	}

	client := NewDNSEndpointClient(restClient, apiv1alpha1.DNSEndpointKind)

	result, err := client.UpdateStatus(ctx, inputEndpoint)

	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, capturedBody)
	assert.Equal(t, int64(2), capturedBody.Status.ObservedGeneration)
}

func TestDNSEndpointClient_Watch(t *testing.T) {
	ctx := context.Background()
	scheme := newTestScheme()
	codecFactory := serializer.WithoutConversionCodecFactory{
		CodecFactory: serializer.NewCodecFactory(scheme),
	}

	fakeWatcher := watch.NewFake()
	defer fakeWatcher.Stop()

	var capturedWatch bool
	restClient := &fake.RESTClient{
		GroupVersion:         apiv1alpha1.GroupVersion,
		NegotiatedSerializer: codecFactory,
		Client: fake.CreateHTTPClient(func(req *http.Request) (*http.Response, error) {
			capturedWatch = req.URL.Query().Get("watch") == "true"
			// Return a response that the fake client can use
			return &http.Response{
				StatusCode: http.StatusOK,
				Header:     defaultHeader(),
				Body:       io.NopCloser(bytes.NewReader(nil)),
			}, nil
		}),
	}

	client := NewDNSEndpointClient(restClient, apiv1alpha1.DNSEndpointKind)

	opts := &metav1.ListOptions{}
	_, _ = client.Watch(ctx, "test-ns", opts)

	// Verify watch=true was set in the options
	assert.True(t, opts.Watch)
	assert.True(t, capturedWatch)
}

func TestDNSEndpointClient_Watch_SetsWatchFlag(t *testing.T) {
	ctx := context.Background()
	scheme := newTestScheme()
	codecFactory := serializer.WithoutConversionCodecFactory{
		CodecFactory: serializer.NewCodecFactory(scheme),
	}

	restClient := &fake.RESTClient{
		GroupVersion:         apiv1alpha1.GroupVersion,
		NegotiatedSerializer: codecFactory,
		Client: fake.CreateHTTPClient(func(_ *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Header:     defaultHeader(),
				Body:       io.NopCloser(bytes.NewReader(nil)),
			}, nil
		}),
	}

	client := NewDNSEndpointClient(restClient, apiv1alpha1.DNSEndpointKind)

	opts := &metav1.ListOptions{Watch: false}
	_, _ = client.Watch(ctx, "test-ns", opts)

	// The Watch method should set opts.Watch = true
	assert.True(t, opts.Watch)
}

func TestDNSEndpointClientInterface(_ *testing.T) {
	// Verify that dnsEndpointClient implements DNSEndpointClient interface
	var _ DNSEndpointClient = (*dnsEndpointClient)(nil)
}

func TestDNSEndpointClient_ResourcePluralization(t *testing.T) {
	tests := []struct {
		kind             string
		expectedResource string
	}{
		{"DNSEndpoint", "dnsendpoints"},
		{"Endpoint", "endpoints"},
		{"Record", "records"},
		{"CNAME", "cnames"},
	}

	scheme := newTestScheme()
	codecFactory := serializer.NewCodecFactory(scheme)

	for _, tt := range tests {
		t.Run(tt.kind, func(t *testing.T) {
			restClient := &fake.RESTClient{
				GroupVersion:         schema.GroupVersion{Group: "test", Version: "v1"},
				NegotiatedSerializer: codecFactory,
			}

			client := NewDNSEndpointClient(restClient, tt.kind)
			impl := client.(*dnsEndpointClient)

			assert.Equal(t, tt.expectedResource, impl.resource)
		})
	}
}
