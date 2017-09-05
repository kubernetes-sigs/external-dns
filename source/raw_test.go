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

	"github.com/kubernetes-incubator/external-dns/endpoint"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRawSource(t *testing.T) {
	t.Run("Interface", testRawSourceImplementsSource)
	t.Run("Endpoints", testRawEndpoints)
}

// testRawSourceImplementsSource tests that rawSource is a valid Source.
func testRawSourceImplementsSource(t *testing.T) {
	assert.Implements(t, (*Source)(nil), new(rawSource))
}

func testRawEndpoints(t *testing.T) {
	for _, ti := range []struct {
		title     string
		endpoints []*endpoint.Endpoint
		expected  []*endpoint.Endpoint
	}{
		{
			title: "test raw endpoints are the same",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName: "example.org",
					Target:  "8.8.8.8",
				},
				{
					DNSName: "new.org",
					Target:  "lb.com",
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName: "example.org",
					Target:  "8.8.8.8",
				},
				{
					DNSName: "new.org",
					Target:  "lb.com",
				},
			},
		},
	} {
		t.Run(ti.title, func(t *testing.T) {
			rawSource := NewRawSource(ti.endpoints)

			res, err := rawSource.Endpoints()
			require.NoError(t, err)

			validateEndpoints(t, res, ti.expected)
		})
	}
}
