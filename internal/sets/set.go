/*
Copyright 2026 The Kubernetes Authors.

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

// Package sets provides a generic set implementation using maps.
package sets

import (
	"cmp"
	"slices"
)

// Empty is a placeholder type for the value of the set map, since we only care about the keys.
type Empty struct{}

// Set is a generic set implementation using a map with empty struct values.
type Set[T comparable] map[T]Empty

// New creates a new set containing the provided items.
func New[T comparable](items ...T) Set[T] {
	s := make(Set[T], len(items))
	for _, item := range items {
		s.Insert(item)
	}
	return s
}

// NewFromMapKeys creates a new set from the keys of the provided map.
func NewFromMapKeys[T comparable, V any](m map[T]V) Set[T] {
	s := make(Set[T], len(m))
	for k := range m {
		s[k] = Empty{}
	}
	return s
}

// Insert adds item to the set.
func (s Set[T]) Insert(item T) {
	s[item] = Empty{}
}

// Delete removes item from the set.
func (s Set[T]) Delete(item T) {
	delete(s, item)
}

// Has returns true if the set contains the item.
func (s Set[T]) Has(item T) bool {
	_, exists := s[item]
	return exists
}

// List returns the items in the set as a slice.
// The order of the items is not guaranteed to be stable.
func (s Set[T]) List() []T {
	keys := make([]T, 0, len(s))
	for k := range s {
		keys = append(keys, k)
	}
	return keys
}

// Sorted returns the items in the set as a sorted slice.
func Sorted[T cmp.Ordered](s Set[T]) []T {
	l := s.List()
	slices.Sort(l)
	return l
}
