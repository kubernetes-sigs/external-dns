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
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/source"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	"sigs.k8s.io/external-dns/pkg/events"
	"sigs.k8s.io/external-dns/source/types"
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
			endpoints, err := source.Endpoints(t.Context())
			require.NoError(t, err)

			// Validate returned endpoints against desired endpoints.
			testutils.ValidateEndpoints(t, endpoints, tc.expected)

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
	_, err := source.Endpoints(t.Context())
	require.EqualError(t, err, "some error")

	// Validate that the nested source was called.
	src.AssertExpectations(t)
}

func testMultiSourceEndpointsDefaultTargets(t *testing.T) {
	t.Run("Defaults applied when source targets are empty", func(t *testing.T) {
		defaultTargets := []string{"127.0.0.1", "127.0.0.2", "foo.example.org", "2001:db8::1"}
		labels := endpoint.Labels{"foo": "bar"}

		sourceEndpoints := []*endpoint.Endpoint{
			{DNSName: "foo", RecordType: "A", Targets: endpoint.Targets{}, Labels: labels},
			{DNSName: "bar", RecordType: "CNAME", Targets: endpoint.Targets{}, Labels: labels},
		}

		expectedEndpoints := []*endpoint.Endpoint{
			{DNSName: "foo", Targets: endpoint.Targets{"127.0.0.1", "127.0.0.2"}, RecordType: "A", Labels: labels},
			{DNSName: "bar", Targets: endpoint.Targets{"foo.example.org"}, RecordType: "CNAME", Labels: labels},
		}

		src := new(testutils.MockSource)
		src.On("Endpoints").Return(sourceEndpoints, nil)

		source := NewMultiSource([]source.Source{src}, defaultTargets, false)

		endpoints, err := source.Endpoints(t.Context())
		require.NoError(t, err)

		testutils.ValidateEndpoints(t, endpoints, expectedEndpoints)

		src.AssertExpectations(t)
	})

	t.Run("Defaults NOT applied when source targets exist", func(t *testing.T) {
		defaultTargets := []string{"127.0.0.1"}
		labels := endpoint.Labels{"foo": "bar"}

		sourceEndpoints := []*endpoint.Endpoint{
			{DNSName: "foo", RecordType: "A", Targets: endpoint.Targets{"8.8.8.8"}, Labels: labels},
			{DNSName: "bar", RecordType: "A", Targets: endpoint.Targets{"8.8.4.4"}, Labels: labels},
		}

		src := new(testutils.MockSource)
		src.On("Endpoints").Return(sourceEndpoints, nil)

		source := NewMultiSource([]source.Source{src}, defaultTargets, false)

		endpoints, err := source.Endpoints(t.Context())
		require.NoError(t, err)

		testutils.ValidateEndpoints(t, endpoints, sourceEndpoints)

		src.AssertExpectations(t)
	})
}

func TestMultiSourceDefaultTargetsRecordType(t *testing.T) {
	defaultTargets := []string{"192.0.2.1", "2001:db8::1", "lb.example.net"}

	for _, tt := range []struct {
		name     string
		force    bool
		source   *endpoint.Endpoint
		expected []*endpoint.Endpoint
	}{
		{
			name:     "A without targets gets only IPv4 defaults",
			source:   endpoint.NewEndpoint("a.example.com", endpoint.RecordTypeA),
			expected: []*endpoint.Endpoint{endpoint.NewEndpoint("a.example.com", endpoint.RecordTypeA, "192.0.2.1")},
		},
		{
			name:     "AAAA without targets gets only IPv6 defaults",
			source:   endpoint.NewEndpoint("aaaa.example.com", endpoint.RecordTypeAAAA),
			expected: []*endpoint.Endpoint{endpoint.NewEndpoint("aaaa.example.com", endpoint.RecordTypeAAAA, "2001:db8::1")},
		},
		{
			name:     "CNAME without targets gets only hostname defaults",
			source:   endpoint.NewEndpoint("cname.example.com", endpoint.RecordTypeCNAME),
			expected: []*endpoint.Endpoint{endpoint.NewEndpoint("cname.example.com", endpoint.RecordTypeCNAME, "lb.example.net")},
		},
		{
			name:     "TXT without targets is left alone",
			source:   endpoint.NewEndpoint("txt.example.com", endpoint.RecordTypeTXT),
			expected: []*endpoint.Endpoint{endpoint.NewEndpoint("txt.example.com", endpoint.RecordTypeTXT)},
		},
		{
			name:     "forced defaults keep the A type",
			force:    true,
			source:   endpoint.NewEndpoint("a.example.com", endpoint.RecordTypeA, "198.51.100.1"),
			expected: []*endpoint.Endpoint{endpoint.NewEndpoint("a.example.com", endpoint.RecordTypeA, "192.0.2.1")},
		},
		{
			name:     "forced defaults keep the CNAME type",
			force:    true,
			source:   endpoint.NewEndpoint("cname.example.com", endpoint.RecordTypeCNAME, "other.example.net"),
			expected: []*endpoint.Endpoint{endpoint.NewEndpoint("cname.example.com", endpoint.RecordTypeCNAME, "lb.example.net")},
		},
		{
			name:     "forced defaults leave TXT alone",
			force:    true,
			source:   endpoint.NewEndpoint("txt.example.com", endpoint.RecordTypeTXT, "hello"),
			expected: []*endpoint.Endpoint{endpoint.NewEndpoint("txt.example.com", endpoint.RecordTypeTXT, "hello")},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			src := new(testutils.MockSource)
			src.On("Endpoints").Return([]*endpoint.Endpoint{tt.source}, nil)

			endpoints, err := NewMultiSource([]source.Source{src}, defaultTargets, tt.force).Endpoints(t.Context())
			require.NoError(t, err)

			testutils.ValidateEndpoints(t, endpoints, tt.expected)
		})
	}

	t.Run("A without targets and no IPv4 default is skipped", func(t *testing.T) {
		src := new(testutils.MockSource)
		src.On("Endpoints").Return([]*endpoint.Endpoint{endpoint.NewEndpoint("a.example.com", endpoint.RecordTypeA)}, nil)

		endpoints, err := NewMultiSource([]source.Source{src}, []string{"lb.example.net"}, false).Endpoints(t.Context())
		require.NoError(t, err)

		assert.Empty(t, endpoints)
	})

	t.Run("forced defaults with no IPv4 default keep the A source targets", func(t *testing.T) {
		src := new(testutils.MockSource)
		src.On("Endpoints").Return([]*endpoint.Endpoint{endpoint.NewEndpoint("a.example.com", endpoint.RecordTypeA, "198.51.100.1")}, nil)

		endpoints, err := NewMultiSource([]source.Source{src}, []string{"lb.example.net"}, true).Endpoints(t.Context())
		require.NoError(t, err)

		testutils.ValidateEndpoints(t, endpoints, []*endpoint.Endpoint{endpoint.NewEndpoint("a.example.com", endpoint.RecordTypeA, "198.51.100.1")})
	})
}

func TestMultiSourceDefaultTargetsKeepRefObjects(t *testing.T) {
	ref := endpoint.NewObjectRef("DNSEndpoint", "externaldns.k8s.io/v1alpha1", "default", "example", "", types.CRD)
	src := new(testutils.MockSource)
	src.On("Endpoints").Return([]*endpoint.Endpoint{
		endpoint.NewEndpoint("example.org", endpoint.RecordTypeA).WithRefObject(ref),
	}, nil)

	endpoints, err := NewMultiSource([]source.Source{src}, []string{"192.0.2.1"}, false).Endpoints(t.Context())
	require.NoError(t, err)

	require.Len(t, endpoints, 1)
	for _, ep := range endpoints {
		assert.Equal(t, []*events.ObjectReference{ref}, ep.RefObjects(), "%s %s", ep.DNSName, ep.RecordType)
	}
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
