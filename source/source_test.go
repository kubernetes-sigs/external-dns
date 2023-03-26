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
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
)

func TestGetTTLFromAnnotations(t *testing.T) {
	for _, tc := range []struct {
		title       string
		annotations map[string]string
		expectedTTL endpoint.TTL
		expectedErr error
	}{
		{
			title:       "TTL annotation not present",
			annotations: map[string]string{"foo": "bar"},
			expectedTTL: endpoint.TTL(0),
			expectedErr: nil,
		},
		{
			title:       "TTL annotation value is not a number",
			annotations: map[string]string{ttlAnnotationKey: "foo"},
			expectedTTL: endpoint.TTL(0),
			expectedErr: fmt.Errorf("\"foo\" is not a valid TTL value"),
		},
		{
			title:       "TTL annotation value is empty",
			annotations: map[string]string{ttlAnnotationKey: ""},
			expectedTTL: endpoint.TTL(0),
			expectedErr: fmt.Errorf("\"\" is not a valid TTL value"),
		},
		{
			title:       "TTL annotation value is negative number",
			annotations: map[string]string{ttlAnnotationKey: "-1"},
			expectedTTL: endpoint.TTL(0),
			expectedErr: fmt.Errorf("TTL value must be between [%d, %d]", ttlMinimum, ttlMaximum),
		},
		{
			title:       "TTL annotation value is too high",
			annotations: map[string]string{ttlAnnotationKey: fmt.Sprintf("%d", 1<<32)},
			expectedTTL: endpoint.TTL(0),
			expectedErr: fmt.Errorf("TTL value must be between [%d, %d]", ttlMinimum, ttlMaximum),
		},
		{
			title:       "TTL annotation value is set correctly using integer",
			annotations: map[string]string{ttlAnnotationKey: "60"},
			expectedTTL: endpoint.TTL(60),
			expectedErr: nil,
		},
		{
			title:       "TTL annotation value is set correctly using duration (whole)",
			annotations: map[string]string{ttlAnnotationKey: "10m"},
			expectedTTL: endpoint.TTL(600),
			expectedErr: nil,
		},
		{
			title:       "TTL annotation value is set correctly using duration (fractional)",
			annotations: map[string]string{ttlAnnotationKey: "20.5s"},
			expectedTTL: endpoint.TTL(20),
			expectedErr: nil,
		},
	} {
		t.Run(tc.title, func(t *testing.T) {
			ttl, err := getTTLFromAnnotations(tc.annotations)
			assert.Equal(t, tc.expectedTTL, ttl)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

func TestSuitableType(t *testing.T) {
	for _, tc := range []struct {
		target, recordType, expected string
	}{
		{"8.8.8.8", "", "A"},
		{"foo.example.org", "", "CNAME"},
		{"bar.eu-central-1.elb.amazonaws.com", "", "CNAME"},
	} {

		recordType := suitableType(tc.target)

		if recordType != tc.expected {
			t.Errorf("expected %s, got %s", tc.expected, recordType)
		}
	}
}

func Test_splitByWhitespaceAndComma(t *testing.T) {
	// Run parallel with other top-level tests in this package.
	t.Parallel()

	testCases := []struct {
		name          string
		inputString   string
		expectedSlice []string
	}{
		// single hostname
		{
			name:          "single hostname",
			inputString:   "fu.bar",
			expectedSlice: []string{"fu.bar"},
		},

		// single hostname with double quotes
		{
			name:          "single hostname with double quotes",
			inputString:   `"fu.bar"`,
			expectedSlice: []string{"fu.bar"},
		},

		// single hostname with single quotes
		{
			name:          "single hostname with single quotes",
			inputString:   `'fu.bar'`,
			expectedSlice: []string{"fu.bar"},
		},

		// single hostname and whitespace
		{
			name:          "single hostname with leading whitespace",
			inputString:   "    fu.bar",
			expectedSlice: []string{"fu.bar"},
		},
		{
			name:          "single hostname with trailing whitespace",
			inputString:   "fu.bar   \t    ",
			expectedSlice: []string{"fu.bar"},
		},
		{
			name:          "single hostname with leading and trailing whitespace",
			inputString:   "  \tfu.bar     \n\r  ",
			expectedSlice: []string{"fu.bar"},
		},

		// single hostname and commas
		{
			name:          "single hostname with leading comma",
			inputString:   ",fu.bar",
			expectedSlice: []string{"fu.bar"},
		},
		{
			name:          "single hostname with trailing comma",
			inputString:   "fu.bar,",
			expectedSlice: []string{"fu.bar"},
		},
		{
			name:          "single hostname with leading and trailing comma",
			inputString:   ",fu.bar,",
			expectedSlice: []string{"fu.bar"},
		},
		{
			name:          "single hostname with multiple leading and trailing commas",
			inputString:   ",,fu.bar,,,",
			expectedSlice: []string{"fu.bar"},
		},

		// single hostname and whitespace and commas
		{
			name:          "single hostname with leading whitespace and trailing comma",
			inputString:   " fu.bar,",
			expectedSlice: []string{"fu.bar"},
		},
		{
			name:          "single hostname with leading comma and trailing whitespace",
			inputString:   ",fu.bar\n",
			expectedSlice: []string{"fu.bar"},
		},
		{
			name:          "single hostname with leading and trailing whitespace and comma",
			inputString:   ",\tfu.bar\r,",
			expectedSlice: []string{"fu.bar"},
		},
		{
			name:          "single hostname with multiple leading and trailing whitespace chars and commas",
			inputString:   ",,\n   \r,\t\t fu.bar,,\r,       \t\t\t\t \n   \r\t",
			expectedSlice: []string{"fu.bar"},
		},

		// two hostsnames and whitespace
		{
			name:          "two hostnames separated by a single whitespace char",
			inputString:   "fu.bar hello.world",
			expectedSlice: []string{"fu.bar", "hello.world"},
		},
		{
			name:          "two hostnames with leading whitespace separated by a single whitespace char",
			inputString:   "   fu.bar hello.world",
			expectedSlice: []string{"fu.bar", "hello.world"},
		},
		{
			name:          "two hostnames with trailing whitespace separated by a single whitespace char",
			inputString:   "fu.bar hello.world    ",
			expectedSlice: []string{"fu.bar", "hello.world"},
		},
		{
			name:          "two hostnames with leading and trailing whitespace separated by a single whitespace char",
			inputString:   "   fu.bar hello.world    ",
			expectedSlice: []string{"fu.bar", "hello.world"},
		},
		{
			name:          "two hostnames separated by multiple whitespace chars",
			inputString:   "fu.bar  \t\t    hello.world",
			expectedSlice: []string{"fu.bar", "hello.world"},
		},
		{
			name:          "two hostnames separated by multiple whitespace chars with leading and trailing whitespace separated by a single whitespace char",
			inputString:   "  \t\n    fu.bar\r      hello.world\r  ",
			expectedSlice: []string{"fu.bar", "hello.world"},
		},

		// the kitchen sink
		{
			name:          "the kitchen sink",
			inputString:   "\n\n,\n\n   \t \"fu.bar\",,,,,,,,,\t\r, ,hello.world,'bye.bye',,,,",
			expectedSlice: []string{"fu.bar", "hello.world", "bye.bye"},
		},
	}

	for i := range testCases {
		tc := testCases[i] // capture the range variable
		t.Run(tc.name, func(t *testing.T) {
			// All test cases should run in parallel.
			t.Parallel()
			actualSlice := splitByWhitespaceAndComma(tc.inputString)
			require.ElementsMatch(t, tc.expectedSlice, actualSlice)
		})
	}
}
