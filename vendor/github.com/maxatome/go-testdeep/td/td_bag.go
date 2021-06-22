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
func Bag(expectedItems ...interface{}) TestDeep {
	bag := newSetBase(allSet, false)
	bag.Add(expectedItems...)
	return &bag
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
func SubBagOf(expectedItems ...interface{}) TestDeep {
	bag := newSetBase(subSet, false)
	bag.Add(expectedItems...)
	return &bag
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
func SuperBagOf(expectedItems ...interface{}) TestDeep {
	bag := newSetBase(superSet, false)
	bag.Add(expectedItems...)
	return &bag
}
