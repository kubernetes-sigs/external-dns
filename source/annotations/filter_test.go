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

	"k8s.io/apimachinery/pkg/labels"

	logtest "sigs.k8s.io/external-dns/internal/testutils/log"
)

// Mock object implementing AnnotatedObject
type mockObj struct {
	annotations map[string]string
}

func (m mockObj) GetAnnotations() map[string]string {
	return m.annotations
}

func mustParseAnnotationFilter(s string) labels.Selector {
	sel, err := ParseFilter(s)
	if err != nil {
		panic(err)
	}
	return sel
}

func TestFilter(t *testing.T) {
	tests := []struct {
		name     string
		items    []mockObj
		filter   labels.Selector
		expected []mockObj
	}{
		{
			name: "nil filter returns all",
			items: []mockObj{
				{annotations: map[string]string{"foo": "bar"}},
				{annotations: map[string]string{"baz": "qux"}},
			},
			filter: nil,
			expected: []mockObj{
				{annotations: map[string]string{"foo": "bar"}},
				{annotations: map[string]string{"baz": "qux"}},
			},
		},
		{
			name: "empty selector returns all",
			items: []mockObj{
				{annotations: map[string]string{"foo": "bar"}},
				{annotations: map[string]string{"baz": "qux"}},
			},
			filter: labels.Everything(),
			expected: []mockObj{
				{annotations: map[string]string{"foo": "bar"}},
				{annotations: map[string]string{"baz": "qux"}},
			},
		},
		{
			name: "matching items",
			items: []mockObj{
				{annotations: map[string]string{"foo": "bar"}},
				{annotations: map[string]string{"foo": "baz"}},
			},
			filter: mustParseAnnotationFilter("foo=bar"),
			expected: []mockObj{
				{annotations: map[string]string{"foo": "bar"}},
			},
		},
		{
			name: "no matching items",
			items: []mockObj{
				{annotations: map[string]string{"foo": "baz"}},
			},
			filter:   mustParseAnnotationFilter("foo=bar"),
			expected: []mockObj{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Filter(tt.items, tt.filter)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFilter_LogOutput(t *testing.T) {
	hook := logtest.LogsUnderTestWithLogLevel(log.DebugLevel, t)

	items := []mockObj{
		{annotations: map[string]string{"foo": "bar"}},
		{annotations: map[string]string{"foo": "baz"}},
	}
	Filter(items, mustParseAnnotationFilter("foo=bar"))

	logtest.TestHelperLogContains("filtered '1' services out of '2' with annotation filter 'foo=bar'", hook, t)
}
