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

package endpoint

import (
	"cmp"
	"iter"
	"slices"
)

// OrderedMap groups endpoints by key while remembering the order in which keys
// were first seen, so iteration is deterministic (unlike a plain map).
type OrderedMap[K comparable] struct {
	keys []K
	m    map[K][]*Endpoint
}

func NewOrderedMap[K comparable](capacity int) *OrderedMap[K] {
	return &OrderedMap[K]{
		keys: make([]K, 0, capacity),
		m:    make(map[K][]*Endpoint, capacity),
	}
}

func (o *OrderedMap[K]) Append(k K, v *Endpoint) {
	if _, ok := o.m[k]; !ok {
		o.keys = append(o.keys, k)
	}
	o.m[k] = append(o.m[k], v)
}

func (o *OrderedMap[K]) All() iter.Seq2[K, []*Endpoint] {
	return func(yield func(K, []*Endpoint) bool) {
		for _, k := range o.keys {
			if !yield(k, o.m[k]) {
				return
			}
		}
	}
}

func CompareByResource(a, b *Endpoint) int {
	return cmp.Compare(a.Labels[ResourceLabelKey], b.Labels[ResourceLabelKey])
}

func SortByResource(eps []*Endpoint) {
	slices.SortStableFunc(eps, CompareByResource)
}

func AbsorbRefObjects(dst *Endpoint, src ...*Endpoint) {
	for _, s := range src {
		for _, ref := range s.refObjects {
			dst.WithRefObject(ref)
		}
	}
}
