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
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/rest/fake"

	apiv1alpha1 "sigs.k8s.io/external-dns/apis/v1alpha1"
	"sigs.k8s.io/external-dns/endpoint"
)

const (
	testNamespace    = "test-ns"
	testEndpointName = "test-endpoint"
)

var (
	// Compile-time check that dnsEndpointClient implements DNSEndpointClient interface
	_ DNSEndpointClient = &dnsEndpointClient{}

	// created once
	testScheme = func() *runtime.Scheme {
		s := runtime.NewScheme()
		_ = apiv1alpha1.AddToScheme(s)
		return s
	}()
	testCodecFactory = serializer.WithoutConversionCodecFactory{
		CodecFactory: serializer.NewCodecFactory(testScheme),
	}
	testCodec = testCodecFactory.LegacyCodec(apiv1alpha1.GroupVersion)
	headers   = http.Header{"Content-Type": []string{runtime.ContentTypeJSON}}
)

func TestNewDNSEndpointClient(t *testing.T) {
	tests := []struct {
		kind             string
		expectedResource string
	}{
		{"DNSEndpoint", "dnsendpoints"},
		{"CustomEndpoint", "customendpoints"},
		{"endpoint", "endpoints"},
	}

	for _, tt := range tests {
		t.Run(tt.kind, func(t *testing.T) {
			client := newTestRESTClientWithKind(tt.kind, nil)

			require.NotNil(t, client)
			impl, ok := client.(*dnsEndpointClient)
			require.True(t, ok)
			assert.Equal(t, tt.expectedResource, impl.resource)
		})
	}
}

func TestDNSEndpointClient_Get(t *testing.T) {
	validEndpoint := &apiv1alpha1.DNSEndpoint{
		TypeMeta: metav1.TypeMeta{
			APIVersion: apiv1alpha1.GroupVersion.String(),
			Kind:       apiv1alpha1.DNSEndpointKind,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      testEndpointName,
			Namespace: testNamespace,
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

	tests := []struct {
		name    string
		epName  string
		handler func(*http.Request) (*http.Response, error)
		wantErr bool
		wantDNS string
	}{
		{
			name:   "success",
			epName: testEndpointName,
			handler: func(req *http.Request) (*http.Response, error) {
				if req.Method == http.MethodGet && req.URL.Path == "/namespaces/"+testNamespace+"/dnsendpoints/"+testEndpointName {
					return okResponse(validEndpoint), nil
				}
				return nil, fmt.Errorf("unexpected request: %s %s", req.Method, req.URL.Path)
			},
			wantDNS: "example.org",
		},
		{
			name:   "not found",
			epName: "nonexistent",
			handler: func(_ *http.Request) (*http.Response, error) {
				return notFoundResponse(), nil
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := newTestRESTClient(tt.handler)

			result, err := client.Get(t.Context(), testNamespace, tt.epName)

			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.NotNil(t, result)
			require.Len(t, result.Spec.Endpoints, 1)
			assert.Equal(t, tt.wantDNS, result.Spec.Endpoints[0].DNSName)
		})
	}
}

func TestDNSEndpointClient_List(t *testing.T) {
	tests := []struct {
		name          string
		namespace     string
		listOpts      *metav1.ListOptions
		expectedPath  string
		expectedLabel string
		responseItems []apiv1alpha1.DNSEndpoint
	}{
		{
			name:         "namespaced",
			namespace:    testNamespace,
			listOpts:     &metav1.ListOptions{},
			expectedPath: "/namespaces/" + testNamespace + "/dnsendpoints",
			responseItems: []apiv1alpha1.DNSEndpoint{
				{ObjectMeta: metav1.ObjectMeta{Name: "endpoint-1", Namespace: testNamespace}},
				{ObjectMeta: metav1.ObjectMeta{Name: "endpoint-2", Namespace: testNamespace}},
			},
		},
		{
			name:         "all namespaces",
			namespace:    "",
			listOpts:     &metav1.ListOptions{},
			expectedPath: "/dnsendpoints",
			responseItems: []apiv1alpha1.DNSEndpoint{
				{ObjectMeta: metav1.ObjectMeta{Name: "ep1", Namespace: "ns1"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "ep2", Namespace: "ns2"}},
			},
		},
		{
			name:          "with label selector",
			namespace:     testNamespace,
			listOpts:      &metav1.ListOptions{LabelSelector: "app=external-dns"},
			expectedPath:  "/namespaces/" + testNamespace + "/dnsendpoints",
			expectedLabel: "app=external-dns",
			responseItems: []apiv1alpha1.DNSEndpoint{
				{ObjectMeta: metav1.ObjectMeta{Name: "filtered-ep", Namespace: testNamespace}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var capturedLabel string
			client := newTestRESTClient(func(req *http.Request) (*http.Response, error) {
				capturedLabel = req.URL.Query().Get("labelSelector")
				if req.Method == http.MethodGet && req.URL.Path == tt.expectedPath {
					return okResponse(&apiv1alpha1.DNSEndpointList{Items: tt.responseItems}), nil
				}
				return nil, fmt.Errorf("unexpected request: %s %s", req.Method, req.URL.Path)
			})

			result, err := client.List(t.Context(), tt.namespace, tt.listOpts)

			require.NoError(t, err)
			assert.Len(t, result.Items, len(tt.responseItems))
			if tt.expectedLabel != "" {
				assert.Equal(t, tt.expectedLabel, capturedLabel)
			}
		})
	}
}

func TestDNSEndpointClient_UpdateStatus(t *testing.T) {
	inputEndpoint := &apiv1alpha1.DNSEndpoint{
		ObjectMeta: metav1.ObjectMeta{
			Name:       testEndpointName,
			Namespace:  testNamespace,
			Generation: 2,
		},
		Status: apiv1alpha1.DNSEndpointStatus{
			ObservedGeneration: 2,
		},
	}

	var capturedBody *apiv1alpha1.DNSEndpoint
	client := newTestRESTClient(func(req *http.Request) (*http.Response, error) {
		expectedPath := "/namespaces/" + testNamespace + "/dnsendpoints/" + testEndpointName + "/status"
		if req.Method == http.MethodPut && req.URL.Path == expectedPath {
			decoder := json.NewDecoder(req.Body)
			capturedBody = &apiv1alpha1.DNSEndpoint{}
			if err := decoder.Decode(capturedBody); err != nil {
				return nil, err
			}
			return okResponse(inputEndpoint), nil
		}
		return nil, fmt.Errorf("unexpected request: %s %s", req.Method, req.URL.Path)
	})

	result, err := client.UpdateStatus(t.Context(), inputEndpoint)

	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, capturedBody)
	assert.Equal(t, int64(2), capturedBody.Status.ObservedGeneration)
}

func TestDNSEndpointClient_Watch(t *testing.T) {
	var capturedWatch bool
	client := newTestRESTClient(func(req *http.Request) (*http.Response, error) {
		capturedWatch = req.URL.Query().Get("watch") == "true"
		return emptyResponse(), nil
	})

	// Test with Watch initially false - should be set to true
	opts := &metav1.ListOptions{Watch: false}
	_, err := client.Watch(t.Context(), testNamespace, opts)

	require.NoError(t, err)
	assert.True(t, opts.Watch, "Watch method should set opts.Watch = true")
	assert.True(t, capturedWatch, "watch=true should be sent in the request")
}

func newTestRESTClient(handler func(*http.Request) (*http.Response, error)) DNSEndpointClient {
	return newTestRESTClientWithKind(apiv1alpha1.DNSEndpointKind, handler)
}

func newTestRESTClientWithKind(kind string, handler func(*http.Request) (*http.Response, error)) DNSEndpointClient {
	restClient := &fake.RESTClient{
		GroupVersion:         apiv1alpha1.GroupVersion,
		NegotiatedSerializer: testCodecFactory,
		Client:               fake.CreateHTTPClient(handler),
	}
	return NewDNSEndpointClient(restClient, kind)
}

func okResponse(obj runtime.Object) *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Header:     headers,
		Body:       io.NopCloser(bytes.NewReader([]byte(runtime.EncodeOrDie(testCodec, obj)))),
	}
}

func emptyResponse() *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Header:     headers,
		Body:       io.NopCloser(bytes.NewReader(nil)),
	}
}

func notFoundResponse() *http.Response {
	return &http.Response{
		StatusCode: http.StatusNotFound,
		Header:     headers,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"kind":"Status","apiVersion":"v1","status":"Failure","code":404}`))),
	}
}
