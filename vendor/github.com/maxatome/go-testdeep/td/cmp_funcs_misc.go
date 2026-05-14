// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

import (
	"runtime"

	"github.com/maxatome/go-testdeep/internal/ctxerr"
	"github.com/maxatome/go-testdeep/internal/types"
	"github.com/maxatome/go-testdeep/internal/util"
)

// CmpTrue is a shortcut for:
//
//	td.Cmp(t, got, true, args...)
//
// Returns true if the test is OK, false if it fails.
//
//	td.CmpTrue(t, IsAvailable(x), "x should be available")
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
//
// See also [CmpFalse].
func CmpTrue(t TestingT, got bool, args ...any) bool {
	t.Helper()
	return Cmp(t, got, true, args...)
}

// CmpFalse is a shortcut for:
//
//	td.Cmp(t, got, false, args...)
//
// Returns true if the test is OK, false if it fails.
//
//	td.CmpFalse(t, IsAvailable(x), "x should not be available")
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
//
// See also [CmpTrue].
func CmpFalse(t TestingT, got bool, args ...any) bool {
	t.Helper()
	return Cmp(t, got, false, args...)
}

func cmpError(ctx ctxerr.Context, t TestingT, got error, args ...any) bool {
	if got != nil {
		return true
	}

	t.Helper()

	formatError(t,
		ctx.FailureIsFatal,
		&ctxerr.Error{
			Context:  ctx,
			Message:  "should be an error",
			Got:      types.RawString("nil"),
			Expected: types.RawString("non-nil error"),
		},
		args...)

	return false
}

func cmpNoError(ctx ctxerr.Context, t TestingT, got error, args ...any) bool {
	if got == nil {
		return true
	}

	t.Helper()

	formatError(t,
		ctx.FailureIsFatal,
		&ctxerr.Error{
			Context:  ctx,
			Message:  "should NOT be an error",
			Got:      got,
			Expected: types.RawString("nil"),
		},
		args...)

	return false
}

// CmpError checks that got is non-nil error.
//
//	_, err := MyFunction(1, 2, 3)
//	td.CmpError(t, err, "MyFunction(1, 2, 3) should return an error")
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
//
// See also [CmpNoError].
func CmpError(t TestingT, got error, args ...any) bool {
	t.Helper()
	return cmpError(newContext(t), t, got, args...)
}

// CmpNoError checks that got is nil error.
//
//	value, err := MyFunction(1, 2, 3)
//	if td.CmpNoError(t, err) {
//	  // one can now check value...
//	}
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
//
// See also [CmpError].
func CmpNoError(t TestingT, got error, args ...any) bool {
	t.Helper()
	return cmpNoError(newContext(t), t, got, args...)
}

func cmpPanic(ctx ctxerr.Context, t TestingT, fn func(), expected any, args ...any) bool {
	t.Helper()

	if ctx.Path.Len() == 1 && ctx.Path.String() == contextDefaultRootName {
		ctx.Path = ctxerr.NewPath(contextPanicRootName)
	}

	var (
		panicked   bool
		panicParam any
	)

	func() {
		defer func() { panicParam = recover() }()
		panicked = true
		fn()
		panicked = false
	}()

	if !panicked {
		formatError(t,
			ctx.FailureIsFatal,
			&ctxerr.Error{
				Context: ctx,
				Message: "should have panicked",
				Summary: ctxerr.NewSummary("did not panic"),
			},
			args...)
		return false
	}

	return cmpDeeply(ctx.AddCustomLevel("→panic()"), t, panicParam, expected, args...)
}

func cmpNotPanic(ctx ctxerr.Context, t TestingT, fn func(), args ...any) bool {
	var (
		panicked   bool
		stackTrace types.RawString
	)

	func() {
		defer func() {
			panicParam := recover()
			if panicked {
				buf := make([]byte, 8192)
				n := runtime.Stack(buf, false)
				for ; n > 0; n-- {
					if buf[n-1] != '\n' {
						break
					}
				}
				stackTrace = types.RawString("panic: " + util.ToString(panicParam) + "\n\n" +
					string(buf[:n]))
			}
		}()
		panicked = true
		fn()
		panicked = false
	}()

	if !panicked {
		return true
	}

	t.Helper()

	if ctx.Path.Len() == 1 && ctx.Path.String() == contextDefaultRootName {
		ctx.Path = ctxerr.NewPath(contextPanicRootName)
	}

	formatError(t,
		ctx.FailureIsFatal,
		&ctxerr.Error{
			Context:  ctx,
			Message:  "should NOT have panicked",
			Got:      stackTrace,
			Expected: types.RawString("not panicking at all"),
		},
		args...)
	return false
}

// CmpPanic calls fn and checks a panic() occurred with the
// expectedPanic parameter. It returns true only if both conditions
// are fulfilled.
//
// Note that calling panic(nil) in fn body is always detected as a
// panic. [runtime] package says: before Go 1.21, programs that called
// panic(nil) observed recover returning nil. Starting in Go 1.21,
// programs that call panic(nil) observe recover returning a
// [*runtime.PanicNilError]. Programs can change back to the old
// behavior by setting GODEBUG=panicnil=1.
//
//	td.CmpPanic(t,
//	  func() { panic("I am panicking!") },
//	  "I am panicking!",
//	  "The function should panic with the right string") // succeeds
//
//	td.CmpPanic(t,
//	  func() { panic("I am panicking!") },
//	  Contains("panicking!"),
//	  "The function should panic with a string containing `panicking!`") // succeeds
//
//	td.CmpPanic(t, func() { panic(nil) }, nil, "Checks for panic(nil)") // succeeds
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
//
// See also [CmpNotPanic].
func CmpPanic(t TestingT, fn func(), expectedPanic any, args ...any) bool {
	t.Helper()
	return cmpPanic(newContext(t), t, fn, expectedPanic, args...)
}

// CmpNotPanic calls fn and checks no panic() occurred. If a panic()
// occurred false is returned then the panic() parameter and the stack
// trace appear in the test report.
//
// Note that calling panic(nil) in fn body is always detected as a
// panic. [runtime] package says: before Go 1.21, programs that called
// panic(nil) observed recover returning nil. Starting in Go 1.21,
// programs that call panic(nil) observe recover returning a
// [*runtime.PanicNilError]. Programs can change back to the old
// behavior by setting GODEBUG=panicnil=1.
//
//	td.CmpNotPanic(t, func() {}) // succeeds as function does not panic
//
//	td.CmpNotPanic(t, func() { panic("I am panicking!") }) // fails
//	td.CmpNotPanic(t, func() { panic(nil) })               // fails too
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
//
// See also [CmpPanic].
func CmpNotPanic(t TestingT, fn func(), args ...any) bool {
	t.Helper()
	return cmpNotPanic(newContext(t), t, fn, args...)
}
