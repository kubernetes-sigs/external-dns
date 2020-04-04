/*
Copyright 2017 The Kubernetes Authors.

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
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
)

func TestMultiSource(t *testing.T) {
	t.Run("Interface", testMultiSourceImplementsSource)
	t.Run("Endpoints", testMultiSourceEndpoints)
	t.Run("EndpointsWithError", testMultiSourceEndpointsWithError)
}

// testMultiSourceImplementsSource tests that multiSource is a valid Source.
func testMultiSourceImplementsSource(t *testing.T) {
	assert.Implements(t, (*Source)(nil), new(multiSource))
}

// testMultiSourceEndpoints tests merged endpoints from children are returned.
func testMultiSourceEndpoints(t *testing.T) {
	foo := &endpoint.Endpoint{DNSName: "foo", Targets: endpoint.Targets{"8.8.8.8"}}
	bar := &endpoint.Endpoint{DNSName: "bar", Targets: endpoint.Targets{"8.8.4.4"}}

	for _, tc := range []struct {
		title           string
		nestedEndpoints [][]*endpoint.Endpoint
		expected        []*endpoint.Endpoint
	}{
		{
			"no child sources return no endpoints",
			nil,
			[]*endpoint.Endpoint{},
		},
		{
			"single empty child source returns no endpoints",
			[][]*endpoint.Endpoint{{}},
			[]*endpoint.Endpoint{},
		},
		{
			"single non-empty child source returns child's endpoints",
			[][]*endpoint.Endpoint{{foo}},
			[]*endpoint.Endpoint{foo},
		},
		{
			"multiple non-empty child sources returns merged children's endpoints",
			[][]*endpoint.Endpoint{{foo}, {bar}},
			[]*endpoint.Endpoint{foo, bar},
		},
	} {
		t.Run(tc.title, func(t *testing.T) {
			// Prepare the nested mock sources.
			sources := make([]Source, 0, len(tc.nestedEndpoints))

			// Populate the nested mock sources.
			for _, endpoints := range tc.nestedEndpoints {
				src := new(testutils.MockSource)
				src.On("Endpoints").Return(endpoints, nil)

				sources = append(sources, src)
			}

			// Create our object under test and get the endpoints.
			source := NewMultiSource(sources)

			// Get endpoints from the source.
			endpoints, err := source.Endpoints()
			require.NoError(t, err)

			// Validate returned endpoints against desired endpoints.
			validateEndpoints(t, endpoints, tc.expected)

			// Validate that the nested sources were called.
			for _, src := range sources {
				src.(*testutils.MockSource).AssertExpectations(t)
			}
		})
	}
}

// testMultiSourceEndpointsWithError tests that an error by a nested source is bubbled up.
func testMultiSourceEndpointsWithError(t *testing.T) {
	// Create the expected error.
	errSomeError := errors.New("some error")

	// Create a mocked source returning that error.
	src := new(testutils.MockSource)
	src.On("Endpoints").Return(nil, errSomeError)

	// Create our object under test and get the endpoints.
	source := NewMultiSource([]Source{src})

	// Get endpoints from our source.
	_, err := source.Endpoints()
	assert.EqualError(t, err, "some error")

	// Validate that the nested source was called.
	src.AssertExpectations(t)
}
