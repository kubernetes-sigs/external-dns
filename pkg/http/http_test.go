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
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

type dummyTransport struct{}

func (d *dummyTransport) RoundTrip(req *http.Request) (*http.Response, error) {
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

func TestPathProcessor(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"/foo/bar", "bar"},
		{"/foo/", ""},
		{"/", ""},
		{"", ""},
		{"/foo/bar/baz", "baz"},
		{"foo/bar", "bar"},
		{"foo", "foo"},
		{"foo/", ""},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			require.Equal(t, tt.expected, pathProcessor(tt.input))
		})
	}
}
