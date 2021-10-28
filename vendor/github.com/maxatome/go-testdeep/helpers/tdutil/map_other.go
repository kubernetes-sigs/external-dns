// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

//go:build !go1.12
// +build !go1.12

package tdutil

import (
	"reflect"
	"sort"
)

// MapEach calls "fn" for each key/value pair of map "m". If "fn"
// returns false, it will not be called again.
func MapEach(m reflect.Value, fn func(k, v reflect.Value) bool) bool {
	ks := MapSortedKeys(m)
	for _, vkey := range ks {
		if !fn(vkey, m.MapIndex(vkey)) {
			return false
		}
	}
	return true
}

// MapEachValue calls "fn" for each value of map "m". If "fn" returns
// false, it will not be called again.
func MapEachValue(m reflect.Value, fn func(k reflect.Value) bool) bool {
	for _, vkey := range m.MapKeys() {
		if !fn(m.MapIndex(vkey)) {
			return false
		}
	}
	return true
}

// MapSortedValues returns a slice of all sorted values of map "m". It
// panics if "m"'s reflect.Kind is not reflect.Map.
func MapSortedValues(m reflect.Value) []reflect.Value {
	vs := make([]reflect.Value, 0, m.Len())
	for _, vkey := range m.MapKeys() {
		vs = append(vs, m.MapIndex(vkey))
	}
	sort.Sort(SortableValues(vs))
	return vs
}
