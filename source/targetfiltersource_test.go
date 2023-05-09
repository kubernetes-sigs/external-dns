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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/context"

	"sigs.k8s.io/external-dns/endpoint"
)

func TestTargetFilterSource(t *testing.T) {
	t.Parallel()

	t.Run("Interface", TestTargetFilterSourceImplementsSource)
	t.Run("Endpoints", TestTargetFilterSourceEndpoints)
}

// TestTargetFilterSourceImplementsSource tests that targetFilterSource is a valid Source.
func TestTargetFilterSourceImplementsSource(t *testing.T) {
	assert.Implements(t, (*Source)(nil), new(targetFilterSource))
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
			title:     "no filters",
			filters:   endpoint.NewTargetNetFilter(nil),
			endpoints: []*endpoint.Endpoint{endpoint.NewEndpoint("foo", "A", "1.2.3.4")},
			expected:  []*endpoint.Endpoint{endpoint.NewEndpoint("foo", "A", "1.2.3.4")},
		},
		{
			title:     "filter matched specific",
			filters:   endpoint.NewTargetNetFilter([]string{"1.2.3.4"}),
			endpoints: []*endpoint.Endpoint{endpoint.NewEndpoint("foo", "A", "1.2.3.4")},
			expected:  []*endpoint.Endpoint{endpoint.NewEndpoint("foo", "A", "1.2.3.4")},
		},
		{
			title:     "filter matched net",
			filters:   endpoint.NewTargetNetFilter([]string{"1.2.3.0/24"}),
			endpoints: []*endpoint.Endpoint{endpoint.NewEndpoint("foo", "A", "1.2.3.4")},
			expected:  []*endpoint.Endpoint{endpoint.NewEndpoint("foo", "A", "1.2.3.4")},
		},
		{
			title:     "filter not matched specific",
			filters:   endpoint.NewTargetNetFilter([]string{"1.2.3.5/32"}),
			endpoints: []*endpoint.Endpoint{endpoint.NewEndpoint("foo", "A", "1.2.3.4")},
			expected:  []*endpoint.Endpoint{},
		},
		{
			title:     "filter not matched net",
			filters:   endpoint.NewTargetNetFilter([]string{"1.2.4.0/24"}),
			endpoints: []*endpoint.Endpoint{endpoint.NewEndpoint("foo", "A", "1.2.3.4")},
			expected:  []*endpoint.Endpoint{},
		},
		{
			title:     "filter not matched CNAME endpoint",
			filters:   endpoint.NewTargetNetFilter([]string{"1.2.4.0/24"}),
			endpoints: []*endpoint.Endpoint{endpoint.NewEndpoint("foo", "CNAME", "foo.bar")},
			expected:  []*endpoint.Endpoint{},
		},
		{
			title:   "filter matched one of two",
			filters: endpoint.NewTargetNetFilter([]string{"1.2.4.0/24"}),
			endpoints: []*endpoint.Endpoint{
				endpoint.NewEndpoint("foo", "CNAME", "foo.bar"),
				endpoint.NewEndpoint("foo", "A", "1.2.4.1")},
			expected: []*endpoint.Endpoint{endpoint.NewEndpoint("foo", "A", "1.2.4.1")},
		},
		{
			title: "filter exclusion all",
			filters: endpoint.NewTargetNetFilterWithExclusions(
				[]string{""},
				[]string{"1.0.0.0/8"}),
			endpoints: []*endpoint.Endpoint{
				endpoint.NewEndpoint("foo", "A", "1.2.3.4"),
				endpoint.NewEndpoint("foo", "A", "1.2.3.5"),
				endpoint.NewEndpoint("foo", "A", "1.2.3.6"),
				endpoint.NewEndpoint("foo", "A", "1.3.4.5"),
				endpoint.NewEndpoint("foo", "A", "1.4.4.5")},
			expected: []*endpoint.Endpoint{},
		},
		{
			title: "filter exclude internal net",
			filters: endpoint.NewTargetNetFilterWithExclusions(
				[]string{""},
				[]string{"10.0.0.0/8"}),
			endpoints: []*endpoint.Endpoint{
				endpoint.NewEndpoint("foo", "A", "10.0.0.1"),
				endpoint.NewEndpoint("foo", "A", "8.8.8.8")},
			expected: []*endpoint.Endpoint{endpoint.NewEndpoint("foo", "A", "8.8.8.8")},
		},
		{
			title: "filter only internal",
			filters: endpoint.NewTargetNetFilterWithExclusions(
				[]string{"10.0.0.0/8"}, []string{}),
			endpoints: []*endpoint.Endpoint{
				endpoint.NewEndpoint("foo", "A", "10.0.0.1"),
				endpoint.NewEndpoint("foo", "A", "8.8.8.8")},
			expected: []*endpoint.Endpoint{endpoint.NewEndpoint("foo", "A", "10.0.0.1")},
		},
	}
	for _, tt := range tests {
		tt := tt
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
