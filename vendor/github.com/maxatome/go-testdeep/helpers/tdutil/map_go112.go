// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

//go:build go1.12
// +build go1.12

package tdutil

import (
	"reflect"
	"sort"

	"github.com/maxatome/go-testdeep/internal/visited"
)

type kv struct {
	key   reflect.Value
	value reflect.Value
}

type kvSlice struct {
	v visited.Visited
	s []kv
}

func newKvSlice(l int) *kvSlice {
	s := kvSlice{}
	if l > 0 {
		s.s = make([]kv, 0, l)
		if l > 1 {
			s.v = visited.NewVisited()
		}
	}
	return &s
}

func (s *kvSlice) Len() int { return len(s.s) }
func (s *kvSlice) Less(i, j int) bool {
	return cmp(s.v, s.s[i].key, s.s[j].key) < 0
}
func (s *kvSlice) Swap(i, j int) { s.s[i], s.s[j] = s.s[j], s.s[i] }

// MapEach calls "fn" for each key/value pair of map "m". If "fn"
// returns false, it will not be called again.
func MapEach(m reflect.Value, fn func(k, v reflect.Value) bool) bool {
	kvs := newKvSlice(m.Len())
	iter := m.MapRange()
	for iter.Next() {
		kvs.s = append(kvs.s, kv{key: iter.Key(), value: iter.Value()})
	}
	sort.Sort(kvs)

	for _, kv := range kvs.s {
		if !fn(kv.key, kv.value) {
			return false
		}
	}
	return true
}

// MapEachValue calls "fn" for each value of map "m". If "fn" returns
// false, it will not be called again.
func MapEachValue(m reflect.Value, fn func(k reflect.Value) bool) bool {
	iter := m.MapRange()
	for iter.Next() {
		if !fn(iter.Value()) {
			return false
		}
	}
	return true
}

// MapSortedValues returns a slice of all sorted values of map "m". It
// panics if "m"'s reflect.Kind is not reflect.Map.
func MapSortedValues(m reflect.Value) []reflect.Value {
	vs := make([]reflect.Value, 0, m.Len())
	iter := m.MapRange()
	for iter.Next() {
		vs = append(vs, iter.Value())
	}
	sort.Sort(SortableValues(vs))
	return vs
}
