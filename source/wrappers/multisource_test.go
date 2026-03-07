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

package wrappers

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/source"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
)

func TestMultiSource(t *testing.T) {
	t.Parallel()

	t.Run("Interface", testMultiSourceImplementsSource)
	t.Run("Endpoints", testMultiSourceEndpoints)
	t.Run("EndpointsWithError", testMultiSourceEndpointsWithError)
	t.Run("EndpointsDefaultTargets", testMultiSourceEndpointsDefaultTargets)
}

// testMultiSourceImplementsSource tests that multiSource is a valid Source.
func testMultiSourceImplementsSource(t *testing.T) {
	assert.Implements(t, (*source.Source)(nil), new(multiSource))
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
			[][]*endpoint.Endpoint{{foo.DeepCopy()}},
			[]*endpoint.Endpoint{foo.DeepCopy()},
		},
		{
			"multiple non-empty child sources returns merged children's endpoints",
			[][]*endpoint.Endpoint{{foo.DeepCopy()}, {bar.DeepCopy()}},
			[]*endpoint.Endpoint{foo.DeepCopy(), bar.DeepCopy()},
		},
	} {

		t.Run(tc.title, func(t *testing.T) {
			t.Parallel()

			// Prepare the nested mock sources.
			sources := make([]source.Source, 0, len(tc.nestedEndpoints))

			// Populate the nested mock sources.
			for _, endpoints := range tc.nestedEndpoints {
				src := new(testutils.MockSource)
				src.On("Endpoints").Return(endpoints, nil)

				sources = append(sources, src)
			}

			// Create our object under test and get the endpoints.
			source := NewMultiSource(sources, nil, false)

			// Get endpoints from the source.
			endpoints, err := source.Endpoints(context.Background())
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
	source := NewMultiSource([]source.Source{src}, nil, false)

	// Get endpoints from our source.
	_, err := source.Endpoints(context.Background())
	assert.EqualError(t, err, "some error")

	// Validate that the nested source was called.
	src.AssertExpectations(t)
}

func testMultiSourceEndpointsDefaultTargets(t *testing.T) {
	t.Run("Defaults applied when source targets are empty", func(t *testing.T) {
		defaultTargetsA := []string{"127.0.0.1", "127.0.0.2"}
		defaultTargetsAAAA := []string{"2001:db8::1"}
		defaultTargetsCName := []string{"foo.example.org"}
		defaultTargets := append(defaultTargetsA, defaultTargetsCName...) // nolint: gocritic // appendAssign
		defaultTargets = append(defaultTargets, defaultTargetsAAAA...)    // nolint: gocritic // appendAssign
		labels := endpoint.Labels{"foo": "bar"}

		// Endpoints FROM SOURCE has NO targets
		sourceEndpoints := []*endpoint.Endpoint{
			{DNSName: "foo", Targets: endpoint.Targets{}, Labels: labels},
			{DNSName: "bar", Targets: endpoint.Targets{}, Labels: labels},
		}

		// Expected endpoints SHOULD HAVE the default targets applied
		expectedEndpoints := []*endpoint.Endpoint{
			{DNSName: "foo", Targets: defaultTargetsA, RecordType: "A", Labels: labels},
			{DNSName: "bar", Targets: defaultTargetsA, RecordType: "A", Labels: labels},
			{DNSName: "foo", Targets: defaultTargetsAAAA, RecordType: "AAAA", Labels: labels},
			{DNSName: "bar", Targets: defaultTargetsAAAA, RecordType: "AAAA", Labels: labels},
			{DNSName: "foo", Targets: defaultTargetsCName, RecordType: "CNAME", Labels: labels},
			{DNSName: "bar", Targets: defaultTargetsCName, RecordType: "CNAME", Labels: labels},
		}

		src := new(testutils.MockSource)
		src.On("Endpoints").Return(sourceEndpoints, nil)

		// Test with forceDefaultTargets=false (default behavior)
		source := NewMultiSource([]source.Source{src}, defaultTargets, false)

		endpoints, err := source.Endpoints(context.Background())
		require.NoError(t, err)

		validateEndpoints(t, endpoints, expectedEndpoints)

		src.AssertExpectations(t)
	})

	t.Run("Defaults NOT applied when source targets exist", func(t *testing.T) {
		defaultTargets := []string{"127.0.0.1"} // Default target
		labels := endpoint.Labels{"foo": "bar"}

		// Endpoints FROM SOURCE HAS targets
		sourceEndpoints := []*endpoint.Endpoint{
			{DNSName: "foo", Targets: endpoint.Targets{"8.8.8.8"}, Labels: labels},
			{DNSName: "bar", Targets: endpoint.Targets{"8.8.4.4"}, Labels: labels},
		}

		// Expected endpoints SHOULD MATCH the source endpoints (defaults ignored)
		expectedEndpoints := []*endpoint.Endpoint{
			{DNSName: "foo", Targets: endpoint.Targets{"8.8.8.8"}, Labels: labels},
			{DNSName: "bar", Targets: endpoint.Targets{"8.8.4.4"}, Labels: labels},
		}

		src := new(testutils.MockSource)
		src.On("Endpoints").Return(sourceEndpoints, nil)

		// Test with forceDefaultTargets=false (default behavior)
		source := NewMultiSource([]source.Source{src}, defaultTargets, false)

		endpoints, err := source.Endpoints(context.Background())
		require.NoError(t, err)

		validateEndpoints(t, endpoints, expectedEndpoints)

		src.AssertExpectations(t)
	})

	t.Run("Defaults forced when source targets exist and flag is set", func(t *testing.T) {
		defaultTargetsA := []string{"127.0.0.1", "127.0.0.2"}
		defaultTargetsAAAA := []string{"2001:db8::1"}
		defaultTargetsCName := []string{"foo.example.org"}
		defaultTargets := append(defaultTargetsA, defaultTargetsCName...) // nolint: gocritic // appendAssign
		defaultTargets = append(defaultTargets, defaultTargetsAAAA...)    // nolint: gocritic // appendAssign
		labels := endpoint.Labels{"foo": "bar"}

		// Endpoints FROM SOURCE HAS targets
		sourceEndpoints := []*endpoint.Endpoint{
			{DNSName: "foo", Targets: endpoint.Targets{"8.8.8.8"}, Labels: labels},
			{DNSName: "bar", Targets: endpoint.Targets{"8.8.4.4"}, Labels: labels},
		}

		// Expected endpoints SHOULD HAVE the default targets applied (old behavior)
		expectedEndpoints := []*endpoint.Endpoint{
			{DNSName: "foo", Targets: defaultTargetsA, RecordType: "A", Labels: labels},
			{DNSName: "bar", Targets: defaultTargetsA, RecordType: "A", Labels: labels},
			{DNSName: "foo", Targets: defaultTargetsAAAA, RecordType: "AAAA", Labels: labels},
			{DNSName: "bar", Targets: defaultTargetsAAAA, RecordType: "AAAA", Labels: labels},
			{DNSName: "foo", Targets: defaultTargetsCName, RecordType: "CNAME", Labels: labels},
			{DNSName: "bar", Targets: defaultTargetsCName, RecordType: "CNAME", Labels: labels},
		}

		src := new(testutils.MockSource)
		src.On("Endpoints").Return(sourceEndpoints, nil)

		// Test with forceDefaultTargets=true (legacy behavior)
		source := NewMultiSource([]source.Source{src}, defaultTargets, true)

		endpoints, err := source.Endpoints(context.Background())
		require.NoError(t, err)

		validateEndpoints(t, endpoints, expectedEndpoints)

		src.AssertExpectations(t)
	})

	t.Run("Defaults applied when source targets are empty and flag is set", func(t *testing.T) {
		defaultTargetsA := []string{"127.0.0.1", "127.0.0.2"}
		defaultTargetsAAAA := []string{"2001:db8::1"}
		defaultTargetsCName := []string{"foo.example.org"}
		defaultTargets := append(defaultTargetsA, defaultTargetsAAAA...) // nolint: gocritic // appendAssign
		defaultTargets = append(defaultTargets, defaultTargetsCName...)  // nolint: gocritic // appendAssign

		labels := endpoint.Labels{"foo": "bar"}

		// Endpoints FROM SOURCE has NO targets
		sourceEndpoints := []*endpoint.Endpoint{
			{DNSName: "empty-target-test", Targets: endpoint.Targets{}, Labels: labels},
		}

		// Expected endpoints SHOULD HAVE the default targets applied
		expectedEndpoints := []*endpoint.Endpoint{
			{DNSName: "empty-target-test", Targets: defaultTargetsA, RecordType: "A", Labels: labels},
			{DNSName: "empty-target-test", Targets: defaultTargetsAAAA, RecordType: "AAAA", Labels: labels},
			{DNSName: "empty-target-test", Targets: defaultTargetsCName, RecordType: "CNAME", Labels: labels},
		}

		src := new(testutils.MockSource)
		src.On("Endpoints").Return(sourceEndpoints, nil)

		// Test with forceDefaultTargets=true
		source := NewMultiSource([]source.Source{src}, defaultTargets, true)

		endpoints, err := source.Endpoints(context.Background())
		require.NoError(t, err)

		validateEndpoints(t, endpoints, expectedEndpoints)

		src.AssertExpectations(t)
	})
}

func TestMultiSource_AddEventHandler(t *testing.T) {
	tests := []struct {
		title   string
		sources []source.Source
		times   int
	}{
		{
			title:   "should not add event handler when sources are empty",
			sources: []source.Source{},
			times:   0,
		},
		{
			title: "should add event handler when sources not empty",
			sources: []source.Source{
				testutils.NewMockSource(),
				testutils.NewMockSource(),
				testutils.NewMockSource(),
			},
			times: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			src := NewMultiSource(tt.sources, []string{}, true)
			src.AddEventHandler(t.Context(), func() {})

			count := 0

			for _, mockSource := range tt.sources {
				mSource := mockSource.(*testutils.MockSource)
				mSource.AssertNumberOfCalls(t, "AddEventHandler", 1)
				count += 1
			}

			assert.Equal(t, tt.times, count)
		})
	}
}
