// Copyright (c) 2020-2023, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

import (
	"reflect"
	"strings"

	"github.com/maxatome/go-testdeep/internal/color"
	"github.com/maxatome/go-testdeep/internal/flat"
	"github.com/maxatome/go-testdeep/internal/types"
)

func flattenFuncIsValid(typ reflect.Type) bool {
	return typ.Kind() == reflect.Func &&
		(typ.NumIn() == 1 && !typ.IsVariadic() ||
			typ.NumIn() == 2 && typ.IsVariadic()) &&
		(typ.NumOut() == 1 || typ.NumOut() == 2 && typ.Out(1) == types.Bool)
}

// Flatten allows to flatten any slice, array, map or int range in
// parameters of operators expecting ...any. fn parameter allows to
// filter and/or transform items before flattening and is described
// below.
//
// For example the [Set] operator is defined as:
//
//	func Set(expectedItems ...any) TestDeep
//
// so when comparing to a []int slice, we usually do:
//
//	got := []int{42, 66, 22}
//	td.Cmp(t, got, td.Set(22, 42, 66))
//
// it works but if the expected items are already in a []int, we have
// to copy them in a []any as it can not be flattened directly
// in [Set] parameters:
//
//	expected := []int{22, 42, 66}
//	expectedIf := make([]any, len(expected))
//	for i, item := range expected {
//	  expectedIf[i] = item
//	}
//	td.Cmp(t, got, td.Set(expectedIf...))
//
// but it is a bit boring and less efficient, as [Set] does not keep
// the []any behind the scene.
//
// The same with Flatten follows:
//
//	expected := []int{22, 42, 66}
//	td.Cmp(t, got, td.Set(td.Flatten(expected)))
//
// Several Flatten calls can be passed, and even combined with normal
// parameters:
//
//	expectedPart1 := []int{11, 22, 33}
//	expectedPart2 := []int{55, 66, 77}
//	expectedPart3 := []int{99}
//	td.Cmp(t, got,
//	  td.Set(
//	    td.Flatten(expectedPart1),
//	    44,
//	    td.Flatten(expectedPart2),
//	    88,
//	    td.Flatten(expectedPart3),
//	  ))
//
// is exactly the same as:
//
//	td.Cmp(t, got, td.Set(11, 22, 33, 44, 55, 66, 77, 88, 99))
//
// Note that Flatten calls can even be nested:
//
//	td.Cmp(t, got,
//	  td.Set(
//	    td.Flatten([]any{
//	      td.Flatten(5),
//	      11,
//	      td.Flatten([]int{22, 33}),
//	      td.Flatten([]int{44, 55, 66}),
//	    }),
//	    77,
//	  ))
//
// is exactly the same as:
//
//	td.Cmp(t, got, td.Set(0, 1, 2, 3, 4, 11, 22, 33, 44, 55, 66, 77))
//
// Maps can be flattened too, keeping in mind there is no particular order:
//
//	td.Flatten(map[int]int{1: 2, 3: 4})
//
// is flattened as 1, 2, 3, 4 or 3, 4, 1, 2.
//
// As seen in the example above, int range:
//
//	td.Flatten(5)
//
// is flattened as 0, 1, 2, 3, 4.
//
// Optional fn parameter can be used to filter and/or transform items
// before flattening. If passed, it has to be one element length and
// this single element can be:
//
//   - untyped nil: it is a no-op, as if it was not passed
//   - a function
//   - a string shortcut
//
// If it is a function, it must be a non-nil function with a signature like:
//
//	func(T) V
//	func(T) (V, bool)
//	func(T, X...) V
//	func(T, X...) (V, bool)
//
// T can be the same as V, but it is not mandatory. The (V, bool)
// returned cases allow to exclude some items when returning
// false. For the variadic cases, X does not matter as the function is
// always called without any variadic argument.
//
// If the function signature does not match these cases, Flatten panics.
//
// If the type of an item of sliceOrMapOrInt is not convertible to T,
// the item is dropped silently, as if fn returns false.
//
// This single element can also be a string among:
//
//	"Smuggle:FIELD"
//	"JSONPointer:/PATH"
//
// that are shortcuts for respectively:
//
//	func(in any) any { return td.Smuggle("FIELD", in) }
//	func(in any) any { return td.JSONPointer("/PATH", in) }
//
// See [Smuggle] and [JSONPointer] for a description of what "FIELD"
// and "/PATH" can really be.
//
// Flatten with an fn can be useful when testing some fields of
// structs in a slice with [Bag] or [Set] operators families as well
// as [List]. As an example, here we test only "Name" field for each
// item of a person slice:
//
//	type person struct {
//	  Name string `json:"name"`
//	  Age  int    `json:"age"`
//	}
//	got := []person{{"alice", 22}, {"bob", 18}, {"brian", 34}, {"britt", 32}}
//
//	td.Cmp(t, got,
//	  td.Bag(td.Flatten(
//	    []string{"alice", "britt", "brian", "bob"},
//	    func(name string) any { return td.Smuggle("Name", name) })))
//	// distributes td.Smuggle for each Name, so is equivalent of:
//	td.Cmp(t, got, td.Bag(
//	  td.Smuggle("Name", "alice"),
//	  td.Smuggle("Name", "britt"),
//	  td.Smuggle("Name", "brian"),
//	  td.Smuggle("Name", "bob"),
//	))
//
//	// Same here using Smuggle string shortcut
//	td.Cmp(t, got,
//	  td.Bag(td.Flatten(
//	    []string{"alice", "britt", "brian", "bob"}, "Smuggle:Name")))
//
//	// Same here, but using JSONPointer operator
//	td.Cmp(t, got,
//	  td.Bag(td.Flatten(
//	    []string{"alice", "britt", "brian", "bob"},
//	    func(name string) any { return td.JSONPointer("/name", name) })))
//
//	// Same here using JSONPointer string shortcut
//	td.Cmp(t, got,
//	  td.Bag(td.Flatten(
//	    []string{"alice", "britt", "brian", "bob"}, "JSONPointer:/name")))
//
//	// Same here, but using SuperJSONOf operator
//	td.Cmp(t, got,
//	  td.Bag(td.Flatten(
//	    []string{"alice", "britt", "brian", "bob"},
//	    func(name string) any { return td.SuperJSONOf(`{"name":$1}`, name) })))
//
//	// Same here, but using Struct operator
//	td.Cmp(t, got,
//	  td.Bag(td.Flatten(
//	    []string{"alice", "britt", "brian", "bob"},
//	    func(name string) any { return td.Struct(person{Name: name}) })))
//
// See also [Grep].
func Flatten(sliceOrMapOrInt any, fn ...any) flat.Slice {
	const (
		smugglePrefix     = "Smuggle:"
		jsonPointerPrefix = "JSONPointer:"
		usage             = "Flatten(SLICE|ARRAY|MAP|int[, FUNC])"
		usageFunc         = usage + `, FUNC should be non-nil func(T) V or func(T) (V, bool) or a string "` + smugglePrefix + `…" or "` + jsonPointerPrefix + `…"`
	)

	if _, isInt := sliceOrMapOrInt.(int); !isInt {
		switch reflect.ValueOf(sliceOrMapOrInt).Kind() {
		case reflect.Slice, reflect.Array, reflect.Map:
		default:
			panic(color.BadUsage(usage, sliceOrMapOrInt, 1, true))
		}
	}

	switch len(fn) {
	case 1:
		if fn[0] != nil {
			break
		}
		fallthrough
	case 0:
		return flat.Slice{Slice: sliceOrMapOrInt}
	default:
		panic(color.TooManyParams(usage))
	}

	f := fn[0]

	// Smuggle & JSONPointer specific shortcuts
	if s, ok := f.(string); ok {
		switch {
		case strings.HasPrefix(s, smugglePrefix):
			f = func(in any) any {
				return Smuggle(s[len(smugglePrefix):], in)
			}
		case strings.HasPrefix(s, jsonPointerPrefix):
			f = func(in any) any {
				return JSONPointer(s[len(jsonPointerPrefix):], in)
			}
		default:
			panic(color.Bad("usage: "+usageFunc+", but received %q as 2nd parameter", s))
		}
	}

	fnType := reflect.TypeOf(f)
	vfn := reflect.ValueOf(f)

	if !flattenFuncIsValid(fnType) {
		panic(color.BadUsage(usageFunc, f, 2, false))
	}
	if vfn.IsNil() {
		panic(color.Bad("usage: " + usageFunc))
	}

	inType := fnType.In(0)

	var final []any
	for _, v := range flat.Values([]any{flat.Slice{Slice: sliceOrMapOrInt}}) {
		if v.Type() != inType {
			if !v.Type().ConvertibleTo(inType) {
				continue
			}
			v = v.Convert(inType)
		}

		ret := vfn.Call([]reflect.Value{v})
		if len(ret) == 1 || ret[1].Bool() {
			final = append(final, ret[0].Interface())
		}
	}

	return flat.Slice{Slice: final}
}
