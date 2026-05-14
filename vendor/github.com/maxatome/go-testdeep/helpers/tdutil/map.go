// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package tdutil

import (
	"reflect"
	"sort"

	tdsort "github.com/maxatome/go-testdeep/internal/sort"
)

// MapSortedKeys returns a slice of all sorted keys of map m. It
// panics if m's [reflect.Kind] is not [reflect.Map].
func MapSortedKeys(m reflect.Value) []reflect.Value {
	ks := m.MapKeys()
	sort.Sort(tdsort.Values(ks))
	return ks
}

// MapEach calls fn for each key/value pair of map m. If fn
// returns false, it will not be called again.
// MapEach returns false if fn returned false.
func MapEach(m reflect.Value, fn func(k, v reflect.Value) bool) bool {
	return tdsort.MapEach(m, fn)
}

// MapEachValue calls fn for each value of map m. If fn returns
// false, it will not be called again.
// MapEachValue returns false if fn returned false.
func MapEachValue(m reflect.Value, fn func(k reflect.Value) bool) bool {
	iter := m.MapRange()
	for iter.Next() {
		if !fn(iter.Value()) {
			return false
		}
	}
	return true
}

// MapSortedValues returns a slice of all sorted values of map m. It
// panics if m's [reflect.Kind] is not [reflect.Map].
func MapSortedValues(m reflect.Value) []reflect.Value {
	vs := make([]reflect.Value, 0, m.Len())
	iter := m.MapRange()
	for iter.Next() {
		vs = append(vs, iter.Value())
	}
	sort.Sort(tdsort.Values(vs))
	return vs
}
