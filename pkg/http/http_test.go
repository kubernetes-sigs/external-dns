/*
Copyright 2025 The Kubernetes Authors.

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

package http

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type dummyTransport struct{}

func (d *dummyTransport) RoundTrip(_ *http.Request) (*http.Response, error) {
	return nil, fmt.Errorf("dummy error")
}

func TestNewInstrumentedTransport(t *testing.T) {
	dt := &dummyTransport{}
	rt := NewInstrumentedTransport(dt)
	crt, ok := rt.(*CustomRoundTripper)
	require.True(t, ok)
	require.Equal(t, dt, crt.next)

	// Should default to http.DefaultTransport if nil
	rt2 := NewInstrumentedTransport(nil)
	crt2, ok := rt2.(*CustomRoundTripper)
	require.True(t, ok)
	require.Equal(t, http.DefaultTransport, crt2.next)
}

func TestNewInstrumentedClient(t *testing.T) {
	client := &http.Client{Transport: &dummyTransport{}}
	result := NewInstrumentedClient(client)
	require.Equal(t, client, result)
	_, ok := result.Transport.(*CustomRoundTripper)
	require.True(t, ok)

	// Should default to http.DefaultClient if nil
	result2 := NewInstrumentedClient(nil)
	require.Equal(t, http.DefaultClient, result2)
	_, ok = result2.Transport.(*CustomRoundTripper)
	require.True(t, ok)
}

func TestCancelRequest(t *testing.T) {
	for _, tt := range []struct {
		title              string
		customRoundTripper CustomRoundTripper
		request            *http.Request
	}{
		{
			title:              "CancelRequest does nothing",
			customRoundTripper: CustomRoundTripper{},
			request:            &http.Request{},
		},
	} {
		t.Run(tt.title, func(_ *testing.T) {
			tt.customRoundTripper.CancelRequest(tt.request)
		})
	}
}

type mockRoundTripper struct {
	response *http.Response
	error    error
}

func (mrt mockRoundTripper) RoundTrip(*http.Request) (*http.Response, error) {
	return mrt.response, mrt.error
}

func TestRoundTrip(t *testing.T) {
	for _, tt := range []struct {
		title            string
		nextRoundTripper mockRoundTripper
		request          *http.Request
		method           string
		url              string
		body             io.Reader

		expectError      bool
		expectedResponse *http.Response
	}{
		{
			title:            "RoundTrip returns no error",
			nextRoundTripper: mockRoundTripper{},
			request: &http.Request{
				Method: http.MethodGet,
				URL: &url.URL{
					Scheme: "HTTPS",
					Host:   "test.local",
					Path:   "/path",
				},
				Body: nil,
			},
			expectError:      false,
			expectedResponse: nil,
		},
		{
			title: "RoundTrip extracts status from request",
			nextRoundTripper: mockRoundTripper{
				response: &http.Response{
					StatusCode: http.StatusOK,
				},
			},
			request: &http.Request{
				Method: http.MethodGet,
				URL: &url.URL{
					Scheme: "HTTPS",
					Host:   "test.local",
					Path:   "/path",
				},
				Body: nil,
			},
			expectError: false,
			expectedResponse: &http.Response{
				StatusCode: http.StatusOK,
			},
		},
	} {
		t.Run(tt.title, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, tt.url, tt.body)
			customRoundTripper := CustomRoundTripper{
				next: tt.nextRoundTripper,
			}

			resp, err := customRoundTripper.RoundTrip(req)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.expectedResponse, resp)
		})
	}
}
