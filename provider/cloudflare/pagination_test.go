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

package cloudflare

import (
	"errors"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockAutoPager[T any] struct {
	items    []T
	index    int
	err      error
	errIndex int
}

func (m *mockAutoPager[T]) Next() bool {
	m.index++
	return !m.hasError() && m.hasNext()
}

func (m *mockAutoPager[T]) Current() T {
	if m.hasNext() && !m.hasError() {
		return m.items[m.index-1]
	}
	var zero T
	return zero
}

func (m *mockAutoPager[T]) Err() error {
	return m.err
}

func (m *mockAutoPager[T]) hasError() bool {
	return m.err != nil && m.errIndex <= m.index
}

func (m *mockAutoPager[T]) hasNext() bool {
	return m.index > 0 && m.index <= len(m.items)
}

func TestAutoPagerIterator(t *testing.T) {
	t.Run("iterate empty", func(t *testing.T) {
		pager := &mockAutoPager[string]{}
		iterator := autoPagerIterator(pager)
		collected := slices.Collect(iterator)
		assert.Empty(t, collected)
	})

	t.Run("iterate all items", func(t *testing.T) {
		pager := &mockAutoPager[int]{items: []int{1, 2, 3, 4, 5}}
		iterator := autoPagerIterator(pager)
		collected := slices.Collect(iterator)
		assert.Equal(t, []int{1, 2, 3, 4, 5}, collected)
	})

	t.Run("iterate with early termination", func(t *testing.T) {
		pager := &mockAutoPager[int]{items: []int{1, 2, 3, 4, 5}}
		iterator := autoPagerIterator(pager)
		var collected []int
		for item := range iterator {
			collected = append(collected, item)
			if item == 3 {
				break
			}
		}
		assert.Equal(t, []int{1, 2, 3}, collected)
	})

	t.Run("iterate with error at index", func(t *testing.T) {
		expectedErr := errors.New("pager error")
		pager := &mockAutoPager[int]{items: []int{1, 2, 3, 4, 5}, err: expectedErr, errIndex: 3}
		iterator := autoPagerIterator(pager)
		collected := slices.Collect(iterator)
		assert.Equal(t, []int{1, 2}, collected)
	})
}
