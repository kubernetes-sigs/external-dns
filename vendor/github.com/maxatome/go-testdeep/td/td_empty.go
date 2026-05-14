// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

import (
	"reflect"

	"github.com/maxatome/go-testdeep/internal/ctxerr"
	"github.com/maxatome/go-testdeep/internal/types"
)

const emptyBadKind = "array OR chan OR map OR slice OR string OR pointer(s) on them"

type tdEmpty struct {
	baseOKNil
}

var _ TestDeep = &tdEmpty{}

// summary(Empty): checks that an array, a channel, a map, a slice or
// a string is empty
// input(Empty): str,array,slice,map,ptr(ptr on array/slice/map/string),chan

// Empty operator checks that an array, a channel, a map, a slice or a
// string is empty. As a special case (non-typed) nil, as well as nil
// channel, map or slice are considered empty.
//
// Note that the compared data can be a pointer (of pointer of pointer
// etc.) on an array, a channel, a map, a slice or a string.
//
//	td.Cmp(t, "", td.Empty())                // succeeds
//	td.Cmp(t, map[string]bool{}, td.Empty()) // succeeds
//	td.Cmp(t, []string{"foo"}, td.Empty())   // fails
func Empty() TestDeep {
	return &tdEmpty{
		baseOKNil: newBaseOKNil(3),
	}
}

// isEmpty returns (isEmpty, kindError) boolean values with only 3
// possible cases:
//   - true, false  → "got" is empty
//   - false, false → "got" is not empty
//   - false, true  → "got" kind is not compatible with emptiness
func isEmpty(got reflect.Value) (bool, bool) {
	switch got.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		return got.Len() == 0, false

	case reflect.Ptr:
		switch got.Type().Elem().Kind() {
		case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice,
			reflect.String:
			if got.IsNil() {
				return true, false
			}
			fallthrough
		case reflect.Ptr:
			return isEmpty(got.Elem())
		default:
			return false, true // bad kind
		}

	default:
		// nil case
		if !got.IsValid() {
			return true, false
		}
		return false, true // bad kind
	}
}

func (e *tdEmpty) Match(ctx ctxerr.Context, got reflect.Value) (err *ctxerr.Error) {
	ok, badKind := isEmpty(got)
	if ok {
		return nil
	}

	if ctx.BooleanError {
		return ctxerr.BooleanError
	}

	if badKind {
		return ctx.CollectError(ctxerr.BadKind(got, emptyBadKind))
	}

	return ctx.CollectError(&ctxerr.Error{
		Message:  "not empty",
		Got:      got,
		Expected: types.RawString("empty"),
	})
}

func (e *tdEmpty) String() string {
	return "Empty()"
}

type tdNotEmpty struct {
	baseOKNil
}

var _ TestDeep = &tdNotEmpty{}

// summary(NotEmpty): checks that an array, a channel, a map, a slice
// or a string is not empty
// input(NotEmpty): str,array,slice,map,ptr(ptr on array/slice/map/string),chan

// NotEmpty operator checks that an array, a channel, a map, a slice
// or a string is not empty. As a special case (non-typed) nil, as
// well as nil channel, map or slice are considered empty.
//
// Note that the compared data can be a pointer (of pointer of pointer
// etc.) on an array, a channel, a map, a slice or a string.
//
//	td.Cmp(t, "", td.NotEmpty())                // fails
//	td.Cmp(t, map[string]bool{}, td.NotEmpty()) // fails
//	td.Cmp(t, []string{"foo"}, td.NotEmpty())   // succeeds
func NotEmpty() TestDeep {
	return &tdNotEmpty{
		baseOKNil: newBaseOKNil(3),
	}
}

func (e *tdNotEmpty) Match(ctx ctxerr.Context, got reflect.Value) (err *ctxerr.Error) {
	ok, badKind := isEmpty(got)
	if ok {
		if ctx.BooleanError {
			return ctxerr.BooleanError
		}
		return ctx.CollectError(&ctxerr.Error{
			Message:  "empty",
			Got:      got,
			Expected: types.RawString("not empty"),
		})
	}

	if badKind {
		if ctx.BooleanError {
			return ctxerr.BooleanError
		}
		return ctx.CollectError(ctxerr.BadKind(got, emptyBadKind))
	}
	return nil
}

func (e *tdNotEmpty) String() string {
	return "NotEmpty()"
}
