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

	"github.com/stretchr/testify/assert"
)

// Mock object implementing AnnotatedObject
type mockObj struct {
	annotations map[string]string
}

func (m mockObj) GetAnnotations() map[string]string {
	return m.annotations
}

func TestFilter_EmptyFilterReturnsAll(t *testing.T) {
	items := []mockObj{
		{annotations: map[string]string{"foo": "bar"}},
		{annotations: map[string]string{"baz": "qux"}},
	}
	result, err := Filter(items, "")
	assert.NoError(t, err)
	assert.Equal(t, items, result)
}

func TestFilter_MatchingItems(t *testing.T) {
	items := []mockObj{
		{annotations: map[string]string{"foo": "bar"}},
		{annotations: map[string]string{"foo": "baz"}},
	}
	result, err := Filter(items, "foo=bar")
	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "bar", result[0].annotations["foo"])
}

func TestFilter_NoMatchingItems(t *testing.T) {
	items := []mockObj{
		{annotations: map[string]string{"foo": "baz"}},
	}
	result, err := Filter(items, "foo=bar")
	assert.NoError(t, err)
	assert.Empty(t, result)
}
