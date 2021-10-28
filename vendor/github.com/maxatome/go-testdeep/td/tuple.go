// Copyright (c) 2020, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

import (
	"reflect"

	"github.com/maxatome/go-testdeep/internal/flat"
)

var tupleType = reflect.TypeOf(tuple{})

// A Tuple is an immutable container. It is used to easily compare
// several values at once, typically when a function returns several
// values:
//
//   price := func(p float64) (float64, string, error) {
//     if p < 0 {
//       return 0, "", errors.New("negative price not supported")
//     }
//     return p * 1.2, "€", nil
//   }
//
//   td.Cmp(t,
//     td.TupleFrom(price(10)),
//     td.TupleFrom(float64(12), "€", nil),
//   )
//
//   td.Cmp(t,
//     td.TupleFrom(price(-10)),
//     td.TupleFrom(float64(0), "", td.Not(nil)),
//   )
//
// Once initialized with TupleFrom, a Tuple is immutable.
type Tuple interface {
	// Len returns t length, aka the number of items the tuple contains.
	Len() int
	// Index returns t's i'th element. It panics if i is out of range.
	Index(int) interface{}
}

// TupleFrom returns a new Tuple initialized to the values of "vals".
//
//   td.TupleFrom(float64(0), "", td.Not(nil))
//
// Flatten can be used to flatten non-[]interface{} slice/array into a
// new Tuple:
//
//   ints := []int64{1, 2, 3}
//   td.TupleFrom(td.Flatten(ints), "OK", nil)
//
// is the same as:
//
//   td.TupleFrom(int64(1), int64(2), int64(3), "OK", nil)
func TupleFrom(vals ...interface{}) Tuple {
	return tuple(flat.Interfaces(vals...))
}

type tuple []interface{}

func (t tuple) Len() int {
	return len(t)
}

func (t tuple) Index(i int) interface{} {
	return t[i]
}
