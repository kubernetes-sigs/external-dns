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
//
//   // works with slices/arrays of any type
//   td.Cmp(t, personSlice, td.Set(
//     Person{Name: "Bob", Age: 32},
//     Person{Name: "Alice", Age: 26},
//   ))
//
// To flatten a non-[]interface{} slice/array, use Flatten function
// and so avoid boring and inefficient copies:
//
//   expected := []int{2, 1}
//   td.Cmp(t, []int{1, 1, 2}, td.Set(td.Flatten(expected))) // succeeds
//   // = td.Cmp(t, []int{1, 1, 2}, td.Set(2, 1))
//
//   exp1 := []int{2, 1}
//   exp2 := []int{5, 8}
//   td.Cmp(t, []int{1, 5, 1, 2, 8, 3, 3},
//     td.Set(td.Flatten(exp1), 3, td.Flatten(exp2))) // succeeds
//   // = td.Cmp(t, []int{1, 5, 1, 2, 8, 3, 3}, td.Set(2, 1, 3, 5, 8))
//
// TypeBehind method can return a non-nil reflect.Type if all items
// known non-interface types are equal, or if only interface types
// are found (mostly issued from Isa()) and they are equal.
func Set(expectedItems ...interface{}) TestDeep {
	return newSetBase(allSet, true, expectedItems)
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
//
//   // works with slices/arrays of any type
//   td.Cmp(t, personSlice, td.SubSetOf(
//     Person{Name: "Bob", Age: 32},
//     Person{Name: "Alice", Age: 26},
//   ))
//
// To flatten a non-[]interface{} slice/array, use Flatten function
// and so avoid boring and inefficient copies:
//
//   expected := []int{2, 1}
//   td.Cmp(t, []int{1, 1}, td.SubSetOf(td.Flatten(expected))) // succeeds
//   // = td.Cmp(t, []int{1, 1}, td.SubSetOf(2, 1))
//
//   exp1 := []int{2, 1}
//   exp2 := []int{5, 8}
//   td.Cmp(t, []int{1, 5, 1, 3, 3},
//     td.SubSetOf(td.Flatten(exp1), 3, td.Flatten(exp2))) // succeeds
//   // = td.Cmp(t, []int{1, 5, 1, 3, 3}, td.SubSetOf(2, 1, 3, 5, 8))
//
// TypeBehind method can return a non-nil reflect.Type if all items
// known non-interface types are equal, or if only interface types
// are found (mostly issued from Isa()) and they are equal.
func SubSetOf(expectedItems ...interface{}) TestDeep {
	return newSetBase(subSet, true, expectedItems)
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
//
//   // works with slices/arrays of any type
//   td.Cmp(t, personSlice, td.SuperSetOf(
//     Person{Name: "Bob", Age: 32},
//     Person{Name: "Alice", Age: 26},
//   ))
//
// To flatten a non-[]interface{} slice/array, use Flatten function
// and so avoid boring and inefficient copies:
//
//   expected := []int{2, 1}
//   td.Cmp(t, []int{1, 1, 2, 8}, td.SuperSetOf(td.Flatten(expected))) // succeeds
//   // = td.Cmp(t, []int{1, 1, 2, 8}, td.SubSetOf(2, 1))
//
//   exp1 := []int{2, 1}
//   exp2 := []int{5, 8}
//   td.Cmp(t, []int{1, 5, 1, 8, 42, 3, 3},
//     td.SuperSetOf(td.Flatten(exp1), 3, td.Flatten(exp2))) // succeeds
//   // = td.Cmp(t, []int{1, 5, 1, 8, 42, 3, 3}, td.SuperSetOf(2, 1, 3, 5, 8))
//
// TypeBehind method can return a non-nil reflect.Type if all items
// known non-interface types are equal, or if only interface types
// are found (mostly issued from Isa()) and they are equal.
func SuperSetOf(expectedItems ...interface{}) TestDeep {
	return newSetBase(superSet, true, expectedItems)
}

// summary(NotAny): compares the contents of an array or a slice, no
// values have to match
// input(NotAny): array,slice,ptr(ptr on array/slice)

// NotAny operator checks that the contents of an array or a slice (or
// a pointer on array/slice) does not contain any of "notExpectedItems".
//
//   td.Cmp(t, []int{1}, td.NotAny(1, 2, 3)) // fails
//   td.Cmp(t, []int{5}, td.NotAny(1, 2, 3)) // succeeds
//
//   // works with slices/arrays of any type
//   td.Cmp(t, personSlice, td.NotAny(
//     Person{Name: "Bob", Age: 32},
//     Person{Name: "Alice", Age: 26},
//   ))
//
// To flatten a non-[]interface{} slice/array, use Flatten function
// and so avoid boring and inefficient copies:
//
//   notExpected := []int{2, 1}
//   td.Cmp(t, []int{4, 4, 3, 8}, td.NotAny(td.Flatten(notExpected))) // succeeds
//   // = td.Cmp(t, []int{4, 4, 3, 8}, td.NotAny(2, 1))
//
//   notExp1 := []int{2, 1}
//   notExp2 := []int{5, 8}
//   td.Cmp(t, []int{4, 4, 42, 8},
//     td.NotAny(td.Flatten(notExp1), 3, td.Flatten(notExp2))) // succeeds
//   // = td.Cmp(t, []int{4, 4, 42, 8}, td.NotAny(2, 1, 3, 5, 8))
//
// Beware that NotAny(…) is not equivalent to Not(Any(…)) but is like
// Not(SuperSet(…)).
//
// TypeBehind method can return a non-nil reflect.Type if all items
// known non-interface types are equal, or if only interface types
// are found (mostly issued from Isa()) and they are equal.
func NotAny(notExpectedItems ...interface{}) TestDeep {
	return newSetBase(noneSet, true, notExpectedItems)
}
