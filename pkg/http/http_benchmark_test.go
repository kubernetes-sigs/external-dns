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
	"bytes"
	"io"
	"net/http"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type roundTripFunc func(req *http.Request) *http.Response

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

// newTestClient returns *http.client with Transport replaced to avoid making real calls
func newTestClient(fn roundTripFunc) *http.Client {
	return &http.Client{
		Transport: NewInstrumentedTransport(fn),
	}
}

type apiUnderTest struct {
	client  *http.Client
	baseURL string
}

func (api *apiUnderTest) doStuff() ([]byte, error) {
	resp, err := api.client.Get(api.baseURL + "/some/path")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func BenchmarkRoundTripper(b *testing.B) {
	client := newTestClient(func(_ *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(`OK`)),
			Header:     make(http.Header),
		}
	})

	for b.Loop() {
		api := apiUnderTest{client, "http://example.com"}
		body, err := api.doStuff()
		require.NoError(b, err)
		assert.Equal(b, []byte("OK"), body)
	}
}

func TestRoundTripper_Concurrent(t *testing.T) {
	client := newTestClient(func(_ *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(`OK`)),
			Header:     make(http.Header),
		}
	})
	api := &apiUnderTest{client: client, baseURL: "http://example.com"}

	const numGoroutines = 100
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for range numGoroutines {
		go func() {
			defer wg.Done()
			body, err := api.doStuff()
			assert.NoError(t, err)
			assert.Equal(t, []byte("OK"), body)
		}()
	}
	wg.Wait()
}
