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

package annotations

import (
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	logtest "sigs.k8s.io/external-dns/internal/testutils/log"
)

// Mock object implementing AnnotatedObject
type mockObj struct {
	annotations map[string]string
}

func (m mockObj) GetAnnotations() map[string]string {
	return m.annotations
}

func TestFilter(t *testing.T) {
	tests := []struct {
		name        string
		items       []mockObj
		filter      string
		expected    []mockObj
		expectError bool
	}{
		{
			name: "Empty filter returns all",
			items: []mockObj{
				{annotations: map[string]string{"foo": "bar"}},
				{annotations: map[string]string{"baz": "qux"}},
			},
			filter: "",
			expected: []mockObj{
				{annotations: map[string]string{"foo": "bar"}},
				{annotations: map[string]string{"baz": "qux"}},
			},
		},
		{
			name: "Matching items",
			items: []mockObj{
				{annotations: map[string]string{"foo": "bar"}},
				{annotations: map[string]string{"foo": "baz"}},
			},
			filter: "foo=bar",
			expected: []mockObj{
				{annotations: map[string]string{"foo": "bar"}},
			},
		},
		{
			name: "No matching items",
			items: []mockObj{
				{annotations: map[string]string{"foo": "baz"}},
			},
			filter:   "foo=bar",
			expected: []mockObj{},
		},
		{
			name: "Whitespace filter returns all",
			items: []mockObj{
				{annotations: map[string]string{"foo": "bar"}},
				{annotations: map[string]string{"baz": "qux"}},
			},
			filter: "   ",
			expected: []mockObj{
				{annotations: map[string]string{"foo": "bar"}},
				{annotations: map[string]string{"baz": "qux"}},
			},
		},
		{
			name: "empty filter returns all",
			items: []mockObj{
				{annotations: map[string]string{"foo": "bar"}},
				{annotations: map[string]string{"baz": "qux"}},
			},
			filter: "",
			expected: []mockObj{
				{annotations: map[string]string{"foo": "bar"}},
				{annotations: map[string]string{"baz": "qux"}},
			},
		},
		{
			name: "invalid filter returns error",
			items: []mockObj{
				{annotations: map[string]string{"foo": "bar"}},
				{annotations: map[string]string{"baz": "qux"}},
			},
			filter: "=invalid",
			expected: []mockObj{
				{annotations: map[string]string{"foo": "bar"}},
				{annotations: map[string]string{"baz": "qux"}},
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Filter(tt.items, tt.filter)
			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestFilter_LogOutput(t *testing.T) {
	hook := logtest.LogsUnderTestWithLogLevel(log.DebugLevel, t)

	items := []mockObj{
		{annotations: map[string]string{"foo": "bar"}},
		{annotations: map[string]string{"foo": "baz"}},
	}
	filter := "foo=bar"
	_, _ = Filter(items, filter)

	logtest.TestHelperLogContains("filtered '1' services out of '2' with annotation filter 'foo=bar'", hook, t)
}
