// Copyright (c) 2020, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

import (
	"reflect"

	"github.com/maxatome/go-testdeep/internal/color"
	"github.com/maxatome/go-testdeep/internal/flat"
)

// Flatten allows to flatten any slice, array or map in parameters of
// operators expecting ...interface{}.
//
// For example the Set operator is defined as:
//
//   func Set(expectedItems ...interface{}) TestDeep
//
// so when comparing to a []int slice, we usually do:
//
//   got := []int{42, 66, 22}
//   td.Cmp(t, got, td.Set(22, 42, 66))
//
// it works but if the expected items are already in a []int, we have
// to copy them in a []interface{} as it can not be flattened directly
// in Set parameters:
//
//   expected := []int{22, 42, 66}
//   expectedIf := make([]interface{}, len(expected))
//   for i, item := range expected {
//     expectedIf[i] = item
//   }
//   td.Cmp(t, got, td.Set(expectedIf...))
//
// but it is a bit boring and less efficient, as Set does not keep the
// []interface{} behind the scene.
//
// The same with Flatten follows:
//
//   expected := []int{22, 42, 66}
//   td.Cmp(t, got, td.Set(td.Flatten(expected)))
//
// Several Flatten calls can be passed, and even combined with normal
// parameters:
//
//   expectedPart1 := []int{11, 22, 33}
//   expectedPart2 := []int{55, 66, 77}
//   expectedPart3 := []int{99}
//   td.Cmp(t, got,
//     td.Set(
//       td.Flatten(expectedPart1),
//       44,
//       td.Flatten(expectedPart2),
//       88,
//       td.Flatten(expectedPart3),
//     ))
//
// is exactly the same as:
//
//   td.Cmp(t, got, td.Set(11, 22, 33, 44, 55, 66, 77, 88, 99))
//
// Note that Flatten calls can even be nested:
//
//   td.Cmp(t, got,
//     td.Set(
//       td.Flatten([]interface{}{
//         11,
//         td.Flatten([]int{22, 33}),
//         td.Flatten([]int{44, 55, 66}),
//       }),
//       77,
//     ))
//
// is exactly the same as:
//
//   td.Cmp(t, got, td.Set(11, 22, 33, 44, 55, 66, 77))
//
// Maps can be flattened too, keeping in mind there is no particular order:
//
//   td.Flatten(map[int]int{1: 2, 3: 4})
//
// is flattened as 1, 2, 3, 4 or 3, 4, 1, 2.
func Flatten(sliceOrMap interface{}) flat.Slice {
	switch reflect.ValueOf(sliceOrMap).Kind() {
	case reflect.Slice, reflect.Array, reflect.Map:
		return flat.Slice{Slice: sliceOrMap}
	default:
		panic(color.BadUsage("Flatten(SLICE|ARRAY|MAP)", sliceOrMap, 1, true))
	}
}
