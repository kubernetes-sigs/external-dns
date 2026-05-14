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

// Values is used to allow the sorting of a [][reflect.Value]
// slice. It is used with the standard sort package:
//
//	vals := []reflect.Value{a, b, c, d}
//	sort.Sort(Values(vals))
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
// See also [Func].
func Values(s []reflect.Value) sort.Interface {
	r := &rValues{
		Slice: s,
	}
	if len(s) > 1 {
		r.Visited = visited.NewVisited()
	}
	return r
}

type rValues struct {
	Visited visited.Visited
	Slice   []reflect.Value
}

func (v *rValues) Len() int {
	return len(v.Slice)
}

func (v *rValues) Less(i, j int) bool {
	return compare.Compare(v.Visited, v.Slice[i], v.Slice[j]) < 0
}

func (v *rValues) Swap(i, j int) {
	v.Slice[i], v.Slice[j] = v.Slice[j], v.Slice[i]
}

// Func returns a function able to compare 2 [reflect.Value] values.
//
// The sorting rules are listed in [Values] documentation.
//
// Cyclic references are correctly handled.
func Func() func(a, b reflect.Value) int {
	var v visited.Visited
	return func(a, b reflect.Value) int {
		return compare.Compare(v, a, b)
	}
}
