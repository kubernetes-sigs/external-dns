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
)

// Validates that mockSource is a Source
var _ Source = &mockSource{}

func TestMockSource(t *testing.T) {
	t.Run("Endpoints", testMockSourceEndpoints)
}

// testMockSourceEndpoints tests that endpoints are returned as given.
func testMockSourceEndpoints(t *testing.T) {
	for _, tc := range []struct {
		title            string
		givenAndExpected []*endpoint.Endpoint
	}{
		{
			"no endpoints given return no endpoints",
			[]*endpoint.Endpoint{},
		},
		{
			"single endpoint given returns single endpoint",
			[]*endpoint.Endpoint{
				{DNSName: "foo", Target: "8.8.8.8"},
			},
		},
	} {
		t.Run(tc.title, func(t *testing.T) {
			// Create our object under test and get the endpoints.
			source := NewMockSource(tc.givenAndExpected)

			endpoints, err := source.Endpoints()
			if err != nil {
				t.Fatal(err)
			}

			// Validate returned endpoints against desired endpoints.
			validateEndpoints(t, endpoints, tc.givenAndExpected)
		})
	}
}
