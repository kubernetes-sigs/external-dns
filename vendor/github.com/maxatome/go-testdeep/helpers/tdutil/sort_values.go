// Copyright (c) 2019-2025, Maxime Soulé
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

// SortableValues is used to allow the sorting of a [][reflect.Value]
// slice. It is used with the standard sort package:
//
//	vals := []reflect.Value{a, b, c, d}
//	sort.Sort(SortableValues(vals))
//	// vals contents now sorted
//
// Replace [sort.Sort] by [sort.Stable] for a stable sort. See [sort]
// documentation.
//
// Sorting rules are as follows:
//   - invalid value is always lower
//   - nil is always lower
//   - different types are sorted by their name
//   - if method TYPE.Compare(TYPE) int exits, calls it
//   - false is lesser than true
//   - float and int numbers are sorted by their value, NaN is always lower
//   - complex numbers are sorted by their real, then by their imaginary parts
//   - strings are sorted by their value
//   - map: shorter length is lesser, then sorted by address
//   - functions, channels and unsafe pointer are sorted by their address
//   - struct: comparison is spread to each field
//   - pointer: comparison is spread to the pointed value
//   - arrays: comparison is spread to each item
//   - slice: comparison is spread to each item, then shorter length is lesser
//   - interface: comparison is spread to the value
//
// Cyclic references are correctly handled.
//
// See also [CmpValuesFunc].
func SortableValues(s []reflect.Value) sort.Interface {
	return tdsort.Values(s)
}

// CmpValuesFunc returns a function able to compare 2 [reflect.Value] values.
//
// The sorting rules are listed in [SortableValues] documentation.
//
//	values := s := []reflect.Value{
//	  reflect.ValueOf(4),
//	  reflect.ValueOf(3),
//	  reflect.ValueOf(1),
//	}
//	slices.SortFunc(s, tdutil.CmpValuesFunc())
//
// Cyclic references are correctly handled.
//
// See also [SortableValues].
func CmpValuesFunc() func(a, b reflect.Value) int {
	return tdsort.Func()
}
