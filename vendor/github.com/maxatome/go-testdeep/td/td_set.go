// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

// summary(Set): compares the contents of an array or a slice ignoring
// duplicates and without taking care of the order of items
// input(Set): array,slice,ptr(ptr on array/slice)

// Set operator compares the contents of an array or a slice (or a
// pointer on array/slice) ignoring duplicates and without taking care
// of the order of items.
//
// During a match, each expected item should match in the compared
// array/slice, and each array/slice item should be matched by an
// expected item to succeed.
//
//   td.Cmp(t, []int{1, 1, 2}, td.Set(1, 2))    // succeeds
//   td.Cmp(t, []int{1, 1, 2}, td.Set(2, 1))    // succeeds
//   td.Cmp(t, []int{1, 1, 2}, td.Set(1, 2, 3)) // fails, 3 is missing
func Set(expectedItems ...interface{}) TestDeep {
	set := newSetBase(allSet, true)
	set.Add(expectedItems...)
	return &set
}

// summary(SubSetOf): compares the contents of an array or a slice
// ignoring duplicates and without taking care of the order of items
// but with potentially some exclusions
// input(SubSetOf): array,slice,ptr(ptr on array/slice)

// SubSetOf operator compares the contents of an array or a slice (or a
// pointer on array/slice) ignoring duplicates and without taking care
// of the order of items.
//
// During a match, each array/slice item should be matched by an
// expected item to succeed. But some expected items can be missing
// from the compared array/slice.
//
//   td.Cmp(t, []int{1, 1}, td.SubSetOf(1, 2))    // succeeds
//   td.Cmp(t, []int{1, 1, 2}, td.SubSetOf(1, 3)) // fails, 2 is an extra item
func SubSetOf(expectedItems ...interface{}) TestDeep {
	set := newSetBase(subSet, true)
	set.Add(expectedItems...)
	return &set
}

// summary(SuperSetOf): compares the contents of an array or a slice
// ignoring duplicates and without taking care of the order of items
// but with potentially some extra items
// input(SuperSetOf): array,slice,ptr(ptr on array/slice)

// SuperSetOf operator compares the contents of an array or a slice (or
// a pointer on array/slice) ignoring duplicates and without taking
// care of the order of items.
//
// During a match, each expected item should match in the compared
// array/slice. But some items in the compared array/slice may not be
// expected.
//
//   td.Cmp(t, []int{1, 1, 2}, td.SuperSetOf(1))    // succeeds
//   td.Cmp(t, []int{1, 1, 2}, td.SuperSetOf(1, 3)) // fails, 3 is missing
func SuperSetOf(expectedItems ...interface{}) TestDeep {
	set := newSetBase(superSet, true)
	set.Add(expectedItems...)
	return &set
}

// summary(NotAny): compares the contents of an array or a slice, no
// values have to match
// input(NotAny): array,slice,ptr(ptr on array/slice)

// NotAny operator checks that the contents of an array or a slice (or
// a pointer on array/slice) does not contain any of "expectedItems".
//
//   td.Cmp(t, []int{1}, td.NotAny(1, 2, 3)) // fails
//   td.Cmp(t, []int{5}, td.NotAny(1, 2, 3)) // succeeds
//
// Beware that NotAny(…) is not equivalent to Not(Any(…)).
func NotAny(expectedItems ...interface{}) TestDeep {
	set := newSetBase(noneSet, true)
	set.Add(expectedItems...)
	return &set
}
