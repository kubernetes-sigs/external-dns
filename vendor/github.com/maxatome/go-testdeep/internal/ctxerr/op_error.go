// Copyright (c) 2021 Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package ctxerr

import (
	"fmt"
	"reflect"

	"github.com/maxatome/go-testdeep/internal/types"
	"github.com/maxatome/go-testdeep/internal/util"
)

// OpBadUsage returns an [*Error] to notice the user she/he passed a bad
// parameter to an operator constructor.
//
// If kind and param's kind name ≠ param's type name:
//
//	usage: {op}{usage}, but received {param type} ({param kind}) as {pos}th parameter
//
// else
//
//	usage: {op}{usage}, but received {param type} as {pos}th parameter
func OpBadUsage(op, usage string, param any, pos int, kind bool) *Error {
	return OpBad(op, "usage: %s%s, %s", op, usage, util.BadParam(param, pos, kind))
}

// OpTooManyParams returns an [*Error] to notice the user she/he called a
// variadic operator constructor with too many parameters.
//
//	usage: {op}{usage}, too many parameters
func OpTooManyParams(op, usage string) *Error {
	return OpBad(op, "usage: %s%s, too many parameters", op, usage)
}

// OpBad returns an [*Error] to notice the user a bad operator
// constructor usage. If len(args) is > 0, s and args are given to
// [fmt.Sprintf].
func OpBad(op, s string, args ...any) *Error {
	if len(args) > 0 {
		s = fmt.Sprintf(s, args...)
	}
	return &Error{
		Message: "bad usage of " + op + " operator",
		Summary: NewSummary(s),
		User:    true,
	}
}

// BadKind returns a “bad kind” [*Error], saying got kind does not
// match kind(s) listed in okKinds. It is the caller responsibility to
// check the kinds compatibility. got can be invalid, in this case it
// is displayed as nil.
func BadKind(got reflect.Value, okKinds string) *Error {
	return &Error{
		Message:  "bad kind",
		Got:      types.RawString(types.KindType(got)),
		Expected: types.RawString(okKinds),
	}
}

// NilPointer returns a “nil pointer” [*Error], saying got value is a
// nil pointer instead of what expected lists. It is the caller
// responsibility to check got contains a nil pointer. got should not
// be invalid.
func NilPointer(got reflect.Value, expected string) *Error {
	return &Error{
		Message:  "nil pointer",
		Got:      types.RawString("nil " + types.KindType(got)),
		Expected: types.RawString(expected),
	}
}
