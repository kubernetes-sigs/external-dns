// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

// summary(Bag): compares the contents of an array or a slice without taking
// care of the order of items
// input(Bag): array,slice,ptr(ptr on array/slice)

// Bag operator compares the contents of an array or a slice (or a
// pointer on array/slice) without taking care of the order of items.
//
// During a match, each expected item should match in the compared
// array/slice, and each array/slice item should be matched by an
// expected item to succeed.
//
//   td.Cmp(t, []int{1, 1, 2}, td.Bag(1, 1, 2))    // succeeds
//   td.Cmp(t, []int{1, 1, 2}, td.Bag(1, 2, 1))    // succeeds
//   td.Cmp(t, []int{1, 1, 2}, td.Bag(2, 1, 1))    // succeeds
//   td.Cmp(t, []int{1, 1, 2}, td.Bag(1, 2))       // fails, one 1 is missing
//   td.Cmp(t, []int{1, 1, 2}, td.Bag(1, 2, 1, 3)) // fails, 3 is missing
//
//   // works with slices/arrays of any type
//   td.Cmp(t, personSlice, td.Bag(
//     Person{Name: "Bob", Age: 32},
//     Person{Name: "Alice", Age: 26},
//   ))
//
// To flatten a non-[]interface{} slice/array, use Flatten function
// and so avoid boring and inefficient copies:
//
//   expected := []int{1, 2, 1}
//   td.Cmp(t, []int{1, 1, 2}, td.Bag(td.Flatten(expected))) // succeeds
//   // = td.Cmp(t, []int{1, 1, 2}, td.Bag(1, 2, 1))
//
//   exp1 := []int{5, 1, 1}
//   exp2 := []int{8, 42, 3}
//   td.Cmp(t, []int{1, 5, 1, 8, 42, 3, 3},
//     td.Bag(td.Flatten(exp1), 3, td.Flatten(exp2))) // succeeds
//   // = td.Cmp(t, []int{1, 5, 1, 8, 42, 3, 3}, td.Bag(5, 1, 1, 3, 8, 42, 3))
//
// TypeBehind method can return a non-nil reflect.Type if all items
// known non-interface types are equal, or if only interface types
// are found (mostly issued from Isa()) and they are equal.
func Bag(expectedItems ...interface{}) TestDeep {
	return newSetBase(allSet, false, expectedItems)
}

// summary(SubBagOf): compares the contents of an array or a slice
// without taking care of the order of items but with potentially some
// exclusions
// input(SubBagOf): array,slice,ptr(ptr on array/slice)

// SubBagOf operator compares the contents of an array or a slice (or a
// pointer on array/slice) without taking care of the order of items.
//
// During a match, each array/slice item should be matched by an
// expected item to succeed. But some expected items can be missing
// from the compared array/slice.
//
//   td.Cmp(t, []int{1}, td.SubBagOf(1, 1, 2))       // succeeds
//   td.Cmp(t, []int{1, 1, 1}, td.SubBagOf(1, 1, 2)) // fails, one 1 is an extra item
//
//   // works with slices/arrays of any type
//   td.Cmp(t, personSlice, td.SubBagOf(
//     Person{Name: "Bob", Age: 32},
//     Person{Name: "Alice", Age: 26},
//   ))
//
// To flatten a non-[]interface{} slice/array, use Flatten function
// and so avoid boring and inefficient copies:
//
//   expected := []int{1, 2, 1}
//   td.Cmp(t, []int{1}, td.SubBagOf(td.Flatten(expected))) // succeeds
//   // = td.Cmp(t, []int{1}, td.SubBagOf(1, 2, 1))
//
//   exp1 := []int{5, 1, 1}
//   exp2 := []int{8, 42, 3}
//   td.Cmp(t, []int{1, 42, 3},
//     td.SubBagOf(td.Flatten(exp1), 3, td.Flatten(exp2))) // succeeds
//   // = td.Cmp(t, []int{1, 42, 3}, td.SubBagOf(5, 1, 1, 3, 8, 42, 3))
//
// TypeBehind method can return a non-nil reflect.Type if all items
// known non-interface types are equal, or if only interface types
// are found (mostly issued from Isa()) and they are equal.
func SubBagOf(expectedItems ...interface{}) TestDeep {
	return newSetBase(subSet, false, expectedItems)
}

// summary(SuperBagOf): compares the contents of an array or a slice
// without taking care of the order of items but with potentially some
// extra items
// input(SuperBagOf): array,slice,ptr(ptr on array/slice)

// SuperBagOf operator compares the contents of an array or a slice (or a
// pointer on array/slice) without taking care of the order of items.
//
// During a match, each expected item should match in the compared
// array/slice. But some items in the compared array/slice may not be
// expected.
//
//   td.Cmp(t, []int{1, 1, 2}, td.SuperBagOf(1))       // succeeds
//   td.Cmp(t, []int{1, 1, 2}, td.SuperBagOf(1, 1, 1)) // fails, one 1 is missing
//
//   // works with slices/arrays of any type
//   td.Cmp(t, personSlice, td.SuperBagOf(
//     Person{Name: "Bob", Age: 32},
//     Person{Name: "Alice", Age: 26},
//   ))
//
// To flatten a non-[]interface{} slice/array, use Flatten function
// and so avoid boring and inefficient copies:
//
//   expected := []int{1, 2, 1}
//   td.Cmp(t, []int{1}, td.SuperBagOf(td.Flatten(expected))) // succeeds
//   // = td.Cmp(t, []int{1}, td.SuperBagOf(1, 2, 1))
//
//   exp1 := []int{5, 1, 1}
//   exp2 := []int{8, 42}
//   td.Cmp(t, []int{1, 5, 1, 8, 42, 3, 3, 6},
//     td.SuperBagOf(td.Flatten(exp1), 3, td.Flatten(exp2))) // succeeds
//   // = td.Cmp(t, []int{1, 5, 1, 8, 42, 3, 3, 6}, td.SuperBagOf(5, 1, 1, 3, 8, 42))
//
// TypeBehind method can return a non-nil reflect.Type if all items
// known non-interface types are equal, or if only interface types
// are found (mostly issued from Isa()) and they are equal.
func SuperBagOf(expectedItems ...interface{}) TestDeep {
	return newSetBase(superSet, false, expectedItems)
}
