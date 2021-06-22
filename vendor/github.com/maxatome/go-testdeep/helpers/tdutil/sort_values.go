// Copyright (c) 2019, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package tdutil

import (
	"reflect"
	"sort"

	"github.com/maxatome/go-testdeep/internal/visited"
)

// SortableValues is used to allow the sorting of a []reflect.Value
// slice. It is used with the standard sort package:
//
//   vals := []reflect.Value{a, b, c, d}
//   sort.Sort(SortableValues(vals))
//   // vals contents now sorted
//
// Replace sort.Sort by sort.Stable for a stable sort. See sort documentation.
//
// Sorting rules are as follows:
//   - nil is always lower
//   - different types are sorted by their name
//   - false is lesser than true
//   - float and int numbers are sorted by their value
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
func SortableValues(s []reflect.Value) sort.Interface {
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
	return cmp(v.Visited, v.Slice[i], v.Slice[j]) < 0
}

func (v *rValues) Swap(i, j int) {
	v.Slice[i], v.Slice[j] = v.Slice[j], v.Slice[i]
}
