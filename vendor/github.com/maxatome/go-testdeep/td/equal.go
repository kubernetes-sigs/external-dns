// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.
//
// deepValueEqual function is heavily based on reflect.deepValueEqual function
// licensed under the BSD-style license found in the LICENSE file in the
// golang repository: https://github.com/golang/go/blob/master/LICENSE

package td

import (
	"fmt"
	"reflect"

	"github.com/maxatome/go-testdeep/helpers/tdutil"
	"github.com/maxatome/go-testdeep/internal/color"
	"github.com/maxatome/go-testdeep/internal/ctxerr"
	"github.com/maxatome/go-testdeep/internal/dark"
	"github.com/maxatome/go-testdeep/internal/types"
)

func isNilStr(isNil bool) types.RawString {
	if isNil {
		return "nil"
	}
	return "not nil"
}

func deepValueEqualFinal(ctx ctxerr.Context, got, expected reflect.Value) (err *ctxerr.Error) {
	err = deepValueEqual(ctx, got, expected)
	if err == nil {
		// Try to merge pending errors
		errMerge := ctx.MergeErrors()
		if errMerge != nil {
			return errMerge
		}
	}
	return
}

func deepValueEqualFinalOK(ctx ctxerr.Context, got, expected reflect.Value) bool {
	ctx = ctx.ResetErrors()
	ctx.BooleanError = true

	return deepValueEqualFinal(ctx, got, expected) == nil
}

// nilHandler is called when one of got or expected is nil (but never
// both, it is caller responsibility).
func nilHandler(ctx ctxerr.Context, got, expected reflect.Value) *ctxerr.Error {
	err := ctxerr.Error{}

	if expected.IsValid() { // here: !got.IsValid()
		if expected.Type().Implements(testDeeper) {
			curOperator := dark.MustGetInterface(expected).(TestDeep)
			ctx.CurOperator = curOperator
			if curOperator.HandleInvalid() {
				return curOperator.Match(ctx, got)
			}
			if ctx.BooleanError {
				return ctxerr.BooleanError
			}

			// Special case if "expected" is a TestDeep operator which
			// does not handle invalid values: the operator is not called,
			// but for the user the error comes from it
		} else if ctx.BooleanError {
			return ctxerr.BooleanError
		}

		err.Expected = expected
	} else { // here: !expected.IsValid() && got.IsValid()
		switch got.Kind() {
		// Special case: "got" is a nil interface, so consider as equal
		// to "expected" nil.
		case reflect.Interface:
			if got.IsNil() {
				return nil
			}
		case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Slice:
			// If BeLax, it is OK: we consider typed nil is equal to (untyped) nil
			if ctx.BeLax && got.IsNil() {
				return nil
			}
		}

		if ctx.BooleanError {
			return ctxerr.BooleanError
		}

		err.Got = got
	}

	err.Message = "values differ"
	return ctx.CollectError(&err)
}

func isCustomEqual(a, b reflect.Value) (bool, bool) {
	aType, bType := a.Type(), b.Type()

	equal, ok := aType.MethodByName("Equal")
	if ok {
		ft := equal.Type
		if !ft.IsVariadic() &&
			ft.NumIn() == 2 &&
			ft.NumOut() == 1 &&
			ft.In(0).AssignableTo(ft.In(1)) &&
			ft.Out(0) == types.Bool &&
			bType.AssignableTo(ft.In(1)) {
			return true, equal.Func.Call([]reflect.Value{a, b})[0].Bool()
		}
	}
	return false, false
}

func deepValueEqual(ctx ctxerr.Context, got, expected reflect.Value) (err *ctxerr.Error) {
	// "got" must not implement testDeeper
	if got.IsValid() && got.Type().Implements(testDeeper) {
		panic(color.Bad("Found a TestDeep operator in got param, " +
			"can only use it in expected one!"))
	}

	if !got.IsValid() || !expected.IsValid() {
		if got.IsValid() == expected.IsValid() {
			return
		}
		return nilHandler(ctx, got, expected)
	}

	// Check if a Smuggle hook matches got type
	if handled, e := ctx.Hooks.Smuggle(&got); handled {
		if e != nil {
			// ctx.BooleanError is always false here as hooks cannot be set globally
			return ctx.CollectError(&ctxerr.Error{
				Message:  e.Error(),
				Got:      got,
				Expected: expected,
			})
		}
	}

	// Check if a Cmp hook matches got & expected types
	if handled, e := ctx.Hooks.Cmp(got, expected); handled {
		if e == nil {
			return
		}
		// ctx.BooleanError is always false here as hooks cannot be set globally
		return ctx.CollectError(&ctxerr.Error{
			Message:  e.Error(),
			Got:      got,
			Expected: expected,
		})
	}

	// Look for an Equal() method
	if ctx.UseEqual || ctx.Hooks.UseEqual(got.Type()) {
		hasEqual, isEqual := isCustomEqual(got, expected)
		if hasEqual {
			if isEqual {
				return
			}
			if ctx.BooleanError {
				return ctxerr.BooleanError
			}
			return ctx.CollectError(&ctxerr.Error{
				Message:  "got.Equal(expected) failed",
				Got:      got,
				Expected: expected,
			})
		}
	}

	if got.Type() != expected.Type() {
		if expected.Type().Implements(testDeeper) {
			curOperator := dark.MustGetInterface(expected).(TestDeep)

			// Resolve interface
			if got.Kind() == reflect.Interface {
				got = got.Elem()

				if !got.IsValid() {
					return nilHandler(ctx, got, expected)
				}
			}

			ctx.CurOperator = curOperator
			return curOperator.Match(ctx, got)
		}

		// "expected" is not a TestDeep operator

		if ctx.BeLax && expected.Type().ConvertibleTo(got.Type()) {
			return deepValueEqual(ctx, got, expected.Convert(got.Type()))
		}

		// If "got" is an interface, try to see what is behind before failing
		// Used by Set/Bag Match method in such cases:
		//     []interface{}{123, "foo"}  →  Bag("foo", 123)
		//    Interface kind -^-----^   but String-^ and ^- Int kinds
		if got.Kind() == reflect.Interface {
			return deepValueEqual(ctx, got.Elem(), expected)
		}

		if ctx.BooleanError {
			return ctxerr.BooleanError
		}
		return ctx.CollectError(ctxerr.TypeMismatch(got.Type(), expected.Type()))
	}

	// if ctx.Depth > 10 { panic("deepValueEqual") } // for debugging

	// Avoid looping forever on cyclic references
	if ctx.Visited.Record(got, expected) {
		return
	}

	// Try to see if a TestDeep operator is anchored in expected
	if op, ok := ctx.Anchors.ResolveAnchor(expected); ok {
		return deepValueEqual(ctx, got, op)
	}

	switch got.Kind() {
	case reflect.Array:
		for i, l := 0, got.Len(); i < l; i++ {
			err = deepValueEqual(ctx.AddArrayIndex(i),
				got.Index(i), expected.Index(i))
			if err != nil {
				return
			}
		}
		return

	case reflect.Slice:
		if got.IsNil() != expected.IsNil() {
			if ctx.BooleanError {
				return ctxerr.BooleanError
			}
			return ctx.CollectError(&ctxerr.Error{
				Message:  "nil slice",
				Got:      isNilStr(got.IsNil()),
				Expected: isNilStr(expected.IsNil()),
			})
		}

		var (
			gotLen      = got.Len()
			expectedLen = expected.Len()
		)

		if gotLen != expectedLen {
			// Shortcut in boolean context
			if ctx.BooleanError {
				return ctxerr.BooleanError
			}
		} else {
			if got.Pointer() == expected.Pointer() {
				return
			}
		}

		var maxLen int
		if gotLen >= expectedLen {
			maxLen = expectedLen
		} else {
			maxLen = gotLen
		}

		// Special case for internal tuple type: it is clearer to read
		// TUPLE instead of DATA when an error occurs when using this type
		if got.Type() == tupleType &&
			ctx.Path.Len() == 1 && ctx.Path.String() == contextDefaultRootName {
			ctx = ctx.ResetPath("TUPLE")
		}

		for i := 0; i < maxLen; i++ {
			err = deepValueEqual(ctx.AddArrayIndex(i),
				got.Index(i), expected.Index(i))
			if err != nil {
				return
			}
		}

		if gotLen != expectedLen {
			res := tdSetResult{
				Kind: itemsSetResult,
				// do not sort Extra/Mising here
			}

			if gotLen > expectedLen {
				res.Extra = make([]reflect.Value, gotLen-expectedLen)
				for i := expectedLen; i < gotLen; i++ {
					res.Extra[i-expectedLen] = got.Index(i)
				}
			} else {
				res.Missing = make([]reflect.Value, expectedLen-gotLen)
				for i := gotLen; i < expectedLen; i++ {
					res.Missing[i-gotLen] = expected.Index(i)
				}
			}

			return ctx.CollectError(&ctxerr.Error{
				Message: fmt.Sprintf("comparing slices, from index #%d", maxLen),
				Summary: res.Summary(),
			})
		}
		return

	case reflect.Interface:
		return deepValueEqual(ctx, got.Elem(), expected.Elem())

	case reflect.Ptr:
		if got.Pointer() == expected.Pointer() {
			return
		}
		return deepValueEqual(ctx.AddPtr(1), got.Elem(), expected.Elem())

	case reflect.Struct:
		sType := got.Type()
		ignoreUnexported := ctx.IgnoreUnexported || ctx.Hooks.IgnoreUnexported(sType)
		for i, n := 0, got.NumField(); i < n; i++ {
			field := sType.Field(i)
			if ignoreUnexported && field.PkgPath != "" {
				continue
			}
			err = deepValueEqual(ctx.AddField(field.Name),
				got.Field(i), expected.Field(i))
			if err != nil {
				return
			}
		}
		return

	case reflect.Map:
		if got.IsNil() != expected.IsNil() {
			if ctx.BooleanError {
				return ctxerr.BooleanError
			}
			return ctx.CollectError(&ctxerr.Error{
				Message:  "nil map",
				Got:      isNilStr(got.IsNil()),
				Expected: isNilStr(expected.IsNil()),
			})
		}

		// Shortcut in boolean context
		if ctx.BooleanError && got.Len() != expected.Len() {
			return ctxerr.BooleanError
		}

		if got.Pointer() == expected.Pointer() {
			return
		}

		var notFoundKeys []reflect.Value
		foundKeys := map[interface{}]bool{}

		for _, vkey := range tdutil.MapSortedKeys(expected) {
			gotValue := got.MapIndex(vkey)
			if !gotValue.IsValid() {
				notFoundKeys = append(notFoundKeys, vkey)
				continue
			}

			err = deepValueEqual(ctx.AddMapKey(vkey),
				gotValue, expected.MapIndex(vkey))
			if err != nil {
				return
			}
			foundKeys[dark.MustGetInterface(vkey)] = true
		}

		if got.Len() == len(foundKeys) {
			if len(notFoundKeys) == 0 {
				return
			}
			return ctx.CollectError(&ctxerr.Error{
				Message: "comparing map",
				Summary: (tdSetResult{
					Kind:    keysSetResult,
					Missing: notFoundKeys,
					Sort:    true,
				}).Summary(),
			})
		}

		if ctx.BooleanError {
			return ctxerr.BooleanError
		}

		// Retrieve extra keys
		res := tdSetResult{
			Kind:    keysSetResult,
			Missing: notFoundKeys,
			Extra:   make([]reflect.Value, 0, got.Len()-len(foundKeys)),
			Sort:    true,
		}

		for _, k := range tdutil.MapSortedKeys(got) {
			if !foundKeys[k.Interface()] {
				res.Extra = append(res.Extra, k)
			}
		}

		return ctx.CollectError(&ctxerr.Error{
			Message: "comparing map",
			Summary: res.Summary(),
		})

	case reflect.Func:
		if got.IsNil() && expected.IsNil() {
			return
		}
		if ctx.BooleanError {
			return ctxerr.BooleanError
		}
		// Can't do better than this:
		return ctx.CollectError(&ctxerr.Error{
			Message: "functions mismatch",
			Summary: ctxerr.NewSummary("<can not be compared>"),
		})

	default:
		// Normal equality suffices
		if dark.MustGetInterface(got) == dark.MustGetInterface(expected) {
			return
		}
		if ctx.BooleanError {
			return ctxerr.BooleanError
		}
		return ctx.CollectError(&ctxerr.Error{
			Message:  "values differ",
			Got:      got,
			Expected: expected,
		})
	}
}

func deepValueEqualOK(got, expected reflect.Value) bool {
	return deepValueEqualFinal(newBooleanContext(), got, expected) == nil
}

// EqDeeply returns true if "got" matches "expected". "expected" can
// be the same type as "got" is, or contains some TestDeep operators.
//
//   got := "foobar"
//   td.EqDeeply(got, "foobar")            // returns true
//   td.EqDeeply(got, td.HasPrefix("foo")) // returns true
func EqDeeply(got, expected interface{}) bool {
	return deepValueEqualOK(reflect.ValueOf(got), reflect.ValueOf(expected))
}

// EqDeeplyError returns nil if "got" matches "expected". "expected"
// can be the same type as got is, or contains some TestDeep
// operators. If "got" does not match "expected", the returned *ctxerr.Error
// contains the reason of the first mismatch detected.
//
//   got := "foobar"
//   if err := td.EqDeeplyError(got, "foobar"); err != nil {
//     // …
//   }
//   if err := td.EqDeeplyError(got, td.HasPrefix("foo")); err != nil {
//     // …
//   }
func EqDeeplyError(got, expected interface{}) error {
	err := deepValueEqualFinal(newContext(),
		reflect.ValueOf(got), reflect.ValueOf(expected))
	if err == nil {
		return nil
	}
	return err
}
