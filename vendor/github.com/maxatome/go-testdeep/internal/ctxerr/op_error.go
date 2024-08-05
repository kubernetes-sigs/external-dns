// Copyright (c) 2021 Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package ctxerr

import (
<<<<<<< HEAD
	"bytes"
	"fmt"
	"reflect"
)

// OpBadUsage returns a string to notice the user he passed a bad
// parameter to an operator constructor.
func OpBadUsage(op, usage string, param any, pos int, kind bool) *Error {
	var b bytes.Buffer
	fmt.Fprintf(&b, "usage: %s%s, but received ", op, usage)

	if param == nil {
		b.WriteString("nil")
	} else {
		t := reflect.TypeOf(param)
		if kind && t.String() != t.Kind().String() {
			fmt.Fprintf(&b, "%s (%s)", t, t.Kind())
		} else {
			b.WriteString(t.String())
		}
	}

	b.WriteString(" as ")
	switch pos {
	case 1:
		b.WriteString("1st")
	case 2:
		b.WriteString("2nd")
	case 3:
		b.WriteString("3rd")
	default:
		fmt.Fprintf(&b, "%dth", pos)
	}
	b.WriteString(" parameter")

	return &Error{
		Message: "bad usage of " + op + " operator",
		Summary: NewSummary(b.String()),
	}
}

// OpTooManyParams returns an [*Error] to notice the user he called a
// variadic operator constructor with too many parameters.
func OpTooManyParams(op, usage string) *Error {
	return &Error{
		Message: "bad usage of " + op + " operator",
		Summary: NewSummary("usage: " + op + usage + ", too many parameters"),
	}
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
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
	"fmt"
	"reflect"
	"strings"

	"github.com/maxatome/go-testdeep/internal/types"
)

// OpBadUsage returns a string to notice the user he passed a bad
// parameter to an operator constructor.
func OpBadUsage(op, usage string, param any, pos int, kind bool) *Error {
	var b strings.Builder
	fmt.Fprintf(&b, "usage: %s%s, but received ", op, usage)

	if param == nil {
		b.WriteString("nil")
	} else {
		t := reflect.TypeOf(param)
		if kind && t.String() != t.Kind().String() {
			fmt.Fprintf(&b, "%s (%s)", t, t.Kind())
		} else {
			b.WriteString(t.String())
		}
	}

	b.WriteString(" as ")
	switch pos {
	case 1:
		b.WriteString("1st")
	case 2:
		b.WriteString("2nd")
	case 3:
		b.WriteString("3rd")
	default:
		fmt.Fprintf(&b, "%dth", pos)
	}
	b.WriteString(" parameter")

	return &Error{
		Message: "bad usage of " + op + " operator",
		Summary: NewSummary(b.String()),
	}
}

// OpTooManyParams returns an [*Error] to notice the user he called a
// variadic operator constructor with too many parameters.
func OpTooManyParams(op, usage string) *Error {
	return &Error{
		Message: "bad usage of " + op + " operator",
		Summary: NewSummary("usage: " + op + usage + ", too many parameters"),
	}
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
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	}
}
