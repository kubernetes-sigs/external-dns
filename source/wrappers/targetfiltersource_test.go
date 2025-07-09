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
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/context"
	"sigs.k8s.io/external-dns/source"

	"sigs.k8s.io/external-dns/endpoint"
)

type mockTargetNetFilter struct {
	targets map[string]bool
}

func NewMockTargetNetFilter(targets []string) endpoint.TargetFilterInterface {
	targetMap := make(map[string]bool)
	for _, target := range targets {
		targetMap[target] = true
	}
	return &mockTargetNetFilter{targets: targetMap}
}

func (m *mockTargetNetFilter) Match(target string) bool {
	return m.targets[target]
}

func (m *mockTargetNetFilter) IsEnabled() bool {
	return true
}

// echoSource is a Source that returns the endpoints passed in on creation.
type echoSource struct {
	mock.Mock
	endpoints []*endpoint.Endpoint
}

func (e *echoSource) AddEventHandler(ctx context.Context, handler func()) {
	e.Called(ctx)
}

// Endpoints returns all the endpoints passed in on creation
func (e *echoSource) Endpoints(ctx context.Context) ([]*endpoint.Endpoint, error) {
	return e.endpoints, nil
}

// NewEchoSource creates a new echoSource.
func NewEchoSource(endpoints []*endpoint.Endpoint) source.Source {
	return &echoSource{endpoints: endpoints}
}

func TestEchoSourceReturnGivenSources(t *testing.T) {
	startEndpoints := []*endpoint.Endpoint{{
		DNSName:    "foo.bar.com",
		RecordType: endpoint.RecordTypeA,
		Targets:    endpoint.Targets{"1.2.3.4"},
		RecordTTL:  endpoint.TTL(300),
		Labels:     endpoint.Labels{},
	}}
	e := NewEchoSource(startEndpoints)

	endpoints, err := e.Endpoints(context.Background())
	if err != nil {
		t.Errorf("Expected no error but got %s", err.Error())
	}

	for i, ep := range endpoints {
		if ep != startEndpoints[i] {
			t.Errorf("Expected %s but got %s", startEndpoints[i], ep)
		}
	}
}

func TestTargetFilterSource(t *testing.T) {
	t.Parallel()

	t.Run("Interface", TestTargetFilterSourceImplementsSource)
	t.Run("Endpoints", TestTargetFilterSourceEndpoints)
}

// TestTargetFilterSourceImplementsSource tests that targetFilterSource is a valid Source.
func TestTargetFilterSourceImplementsSource(t *testing.T) {
	var _ source.Source = &targetFilterSource{}
}

func TestTargetFilterSourceEndpoints(t *testing.T) {
	t.Parallel()

	tests := []struct {
		title     string
		filters   endpoint.TargetFilterInterface
		endpoints []*endpoint.Endpoint
		expected  []*endpoint.Endpoint
	}{
		{
			title:   "filter exclusion all",
			filters: NewMockTargetNetFilter([]string{}),
			endpoints: []*endpoint.Endpoint{
				endpoint.NewEndpoint("foo", endpoint.RecordTypeA, "1.2.3.4"),
				endpoint.NewEndpoint("foo", endpoint.RecordTypeA, "1.2.3.5"),
				endpoint.NewEndpoint("foo", endpoint.RecordTypeA, "1.2.3.6"),
				endpoint.NewEndpoint("foo", endpoint.RecordTypeA, "1.3.4.5"),
				endpoint.NewEndpoint("foo", endpoint.RecordTypeA, "1.4.4.5")},
			expected: []*endpoint.Endpoint{},
		},
		{
			title:   "filter exclude internal net",
			filters: NewMockTargetNetFilter([]string{"8.8.8.8"}),
			endpoints: []*endpoint.Endpoint{
				endpoint.NewEndpoint("foo", endpoint.RecordTypeA, "10.0.0.1"),
				endpoint.NewEndpoint("foo", endpoint.RecordTypeA, "8.8.8.8")},
			expected: []*endpoint.Endpoint{endpoint.NewEndpoint("foo", endpoint.RecordTypeA, "8.8.8.8")},
		},
		{
			title:   "filter only internal",
			filters: NewMockTargetNetFilter([]string{"10.0.0.1"}),
			endpoints: []*endpoint.Endpoint{
				endpoint.NewEndpoint("foo", endpoint.RecordTypeA, "10.0.0.1"),
				endpoint.NewEndpoint("foo", endpoint.RecordTypeA, "8.8.8.8")},
			expected: []*endpoint.Endpoint{endpoint.NewEndpoint("foo", endpoint.RecordTypeA, "10.0.0.1")},
		},
	}
	for _, tt := range tests {

		t.Run(tt.title, func(t *testing.T) {
			t.Parallel()

			echo := NewEchoSource(tt.endpoints)
			src := NewTargetFilterSource(echo, tt.filters)

			endpoints, err := src.Endpoints(context.Background())
			require.NoError(t, err, "failed to get Endpoints")
			validateEndpoints(t, endpoints, tt.expected)
		})
	}
}

func TestTargetFilterConcreteTargetFilter(t *testing.T) {
	tests := []struct {
		title     string
		filters   endpoint.TargetFilterInterface
		endpoints []*endpoint.Endpoint
		expected  []*endpoint.Endpoint
	}{
		{
			title:   "should skip filtering if no filters are set",
			filters: endpoint.NewTargetNetFilterWithExclusions([]string{}, []string{}),
			endpoints: []*endpoint.Endpoint{
				endpoint.NewEndpoint("foo", endpoint.RecordTypeA, "1.2.3.4"),
				endpoint.NewEndpoint("foo", endpoint.RecordTypeA, "1.2.3.5"),
				endpoint.NewEndpoint("foo", endpoint.RecordTypeA, "1.2.3.6"),
			},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("foo", endpoint.RecordTypeA, "1.2.3.4"),
				endpoint.NewEndpoint("foo", endpoint.RecordTypeA, "1.2.3.5"),
				endpoint.NewEndpoint("foo", endpoint.RecordTypeA, "1.2.3.6"),
			},
		},
		{
			title:   "should include all targets when filters are not correctly set",
			filters: endpoint.NewTargetNetFilterWithExclusions([]string{"8.8.8.8"}, []string{}),
			endpoints: []*endpoint.Endpoint{
				endpoint.NewEndpoint("foo", endpoint.RecordTypeA, "10.0.0.1"),
				endpoint.NewEndpoint("foo", endpoint.RecordTypeA, "8.8.8.8")},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("foo", endpoint.RecordTypeA, "10.0.0.1"),
				endpoint.NewEndpoint("foo", endpoint.RecordTypeA, "8.8.8.8"),
			},
		},
		{
			title:   "should include internal when include filter is set",
			filters: endpoint.NewTargetNetFilterWithExclusions([]string{"10.0.0.0/8"}, []string{}),
			endpoints: []*endpoint.Endpoint{
				endpoint.NewEndpoint("foo", endpoint.RecordTypeA, "10.0.0.1"),
				endpoint.NewEndpoint("foo", endpoint.RecordTypeA, "49.13.41.161")},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("foo", endpoint.RecordTypeA, "10.0.0.1"),
			},
		},
		{
			title:   "exclude internal keep public ips",
			filters: endpoint.NewTargetNetFilterWithExclusions([]string{}, []string{"10.0.0.0/8"}),
			endpoints: []*endpoint.Endpoint{
				endpoint.NewEndpoint("foo", endpoint.RecordTypeA, "10.0.178.43"),
				endpoint.NewEndpoint("foo", endpoint.RecordTypeA, "10.0.1.101"),
				endpoint.NewEndpoint("foo", endpoint.RecordTypeA, "49.13.41.161")},
			expected: []*endpoint.Endpoint{endpoint.NewEndpoint("foo", endpoint.RecordTypeA, "49.13.41.161")},
		},
		{
			title:   "should not exclude ipv6 when excluding ipv4",
			filters: endpoint.NewTargetNetFilterWithExclusions([]string{}, []string{"10.0.0.0/8"}),
			endpoints: []*endpoint.Endpoint{
				endpoint.NewEndpoint("foo", endpoint.RecordTypeA, "10.0.178.43"),
				endpoint.NewEndpoint("foo", endpoint.RecordTypeAAAA, "2a01:asdf:asdf:asdf::1"),
			},
			expected: []*endpoint.Endpoint{endpoint.NewEndpoint("foo", endpoint.RecordTypeAAAA, "2a01:asdf:asdf:asdf::1")},
		},
		{
			title:   "should not include ipv6 when including ipv4",
			filters: endpoint.NewTargetNetFilterWithExclusions([]string{"10.0.0.0/8"}, []string{}),
			endpoints: []*endpoint.Endpoint{
				endpoint.NewEndpoint("foo", endpoint.RecordTypeA, "10.0.178.43"),
				endpoint.NewEndpoint("foo", endpoint.RecordTypeAAAA, "2a01:asdf:asdf:asdf::1"),
			},
			expected: []*endpoint.Endpoint{endpoint.NewEndpoint("foo", endpoint.RecordTypeA, "10.0.178.43")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			echo := NewEchoSource(tt.endpoints)
			src := NewTargetFilterSource(echo, tt.filters)

			endpoints, err := src.Endpoints(context.Background())
			require.NoError(t, err, "failed to get Endpoints")

			validateEndpoints(t, endpoints, tt.expected)
		})
	}
}

func TestTargetFilterSource_AddEventHandler(t *testing.T) {
	tests := []struct {
		title   string
		filters endpoint.TargetFilterInterface
		times   int
	}{
		{
			title:   "should add event handler if target filter is enabled",
			filters: endpoint.NewTargetNetFilterWithExclusions([]string{"10.0.0.0/8"}, []string{}),
			times:   1,
		},
		{
			title:   "should not add event handler if target filter is disabled",
			filters: endpoint.NewTargetNetFilterWithExclusions([]string{}, []string{}),
			times:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			echo := NewEchoSource([]*endpoint.Endpoint{})

			m := echo.(*echoSource)
			m.On("AddEventHandler", t.Context()).Return()

			src := NewTargetFilterSource(echo, tt.filters)
			src.AddEventHandler(t.Context(), func() {})

			m.AssertNumberOfCalls(t, "AddEventHandler", tt.times)
		})
	}
}
