// Copyright (c) 2025, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package sort

import (
	"reflect"
	"sort"

	"github.com/maxatome/go-testdeep/internal/compare"
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
	return compare.Compare(s.v, s.s[i].key, s.s[j].key) < 0
}
func (s *kvSlice) Swap(i, j int) { s.s[i], s.s[j] = s.s[j], s.s[i] }

// MapEach calls fn for each key/value pair of map m using the order
// of keys. If fn returns false, it will not be called again.
// MapEach returns false if fn returned false.
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
