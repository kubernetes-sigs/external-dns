// Copyright (c) 2019, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package visited

import (
	"reflect"
)

// visitKey is used by Context and its Visited map to handle cyclic references.
type visitedKey struct {
	a1  uintptr
	a2  uintptr
	typ reflect.Type
}

// Visited allows to remember couples of same type pointers, typically
// to not do the same action twice if the couple has already been seen.
type Visited map[visitedKey]bool

// NewVisited returns a new Visited instance.
func NewVisited() Visited {
	return Visited{}
}

// Record checks and, if needed, records a new entry for (got,
// expected) couple. It returns true if got & expected are pointers
// and have already been seen together. It returns false otherwise.
// It is the caller responsibility to check that got and expected
// types are the same.
func (v Visited) Record(got, expected reflect.Value) bool {
	var addr1, addr2 uintptr
	switch got.Kind() {
	// Pointer() can not be used for interfaces and for slices the
	// returned address is the array behind the slice, use UnsafeAddr()
	// instead
	case reflect.Slice, reflect.Interface:
		if got.IsNil() || expected.IsNil() ||
			!got.CanAddr() || !expected.CanAddr() {
			return false
		}
		addr1 = got.UnsafeAddr()
		addr2 = expected.UnsafeAddr()

		// For maps and pointers use Pointer() to automatically handle
		// indirect pointers
	case reflect.Map, reflect.Ptr:
		if got.IsNil() || expected.IsNil() {
			return false
		}
		addr1 = got.Pointer()
		addr2 = expected.Pointer()

	default:
		return false
	}

	if addr1 > addr2 {
		// Canonicalize order to reduce number of entries in v.
		// Assumes non-moving garbage collector.
		addr1, addr2 = addr2, addr1
	}

	k := visitedKey{
		a1:  addr1,
		a2:  addr2,
		typ: got.Type(),
	}
	if v[k] {
		return true // references already seen
	}

	// Remember for later.
	v[k] = true
	return false
}
