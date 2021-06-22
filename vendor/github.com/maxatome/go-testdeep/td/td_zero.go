// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

import (
	"reflect"

	"github.com/maxatome/go-testdeep/internal/ctxerr"
)

type tdZero struct {
	baseOKNil
}

var _ TestDeep = &tdZero{}

// summary(Zero): checks data against its zero'ed conterpart
// input(Zero): all

// Zero operator checks that data is zero regarding its type.
//
//   - nil is the zero value of pointers, maps, slices, channels and functions;
//   - 0 is the zero value of numbers;
//   - "" is the 0 value of strings;
//   - false is the zero value of booleans;
//   - zero value of structs is the struct with no fields initialized.
//
// Beware that:
//
//   td.Cmp(t, AnyStruct{}, td.Zero())          // is true
//   td.Cmp(t, &AnyStruct{}, td.Zero())         // is false, coz pointer ≠ nil
//   td.Cmp(t, &AnyStruct{}, td.Ptr(td.Zero())) // is true
func Zero() TestDeep {
	return &tdZero{
		baseOKNil: newBaseOKNil(3),
	}
}

func (z *tdZero) Match(ctx ctxerr.Context, got reflect.Value) (err *ctxerr.Error) {
	// nil case
	if !got.IsValid() {
		return nil
	}
	return deepValueEqual(ctx, got, reflect.New(got.Type()).Elem())
}

func (z *tdZero) String() string {
	return "Zero()"
}

type tdNotZero struct {
	baseOKNil
}

var _ TestDeep = &tdNotZero{}

// summary(NotZero): checks that data is not zero regarding its type
// input(NotZero): all

// NotZero operator checks that data is not zero regarding its type.
//
//   - nil is the zero value of pointers, maps, slices, channels and functions;
//   - 0 is the zero value of numbers;
//   - "" is the 0 value of strings;
//   - false is the zero value of booleans;
//   - zero value of structs is the struct with no fields initialized.
//
// Beware that:
//
//   td.Cmp(t, AnyStruct{}, td.NotZero())          // is false
//   td.Cmp(t, &AnyStruct{}, td.NotZero())         // is true, coz pointer ≠ nil
//   td.Cmp(t, &AnyStruct{}, td.Ptr(td.NotZero())) // is false
func NotZero() TestDeep {
	return &tdNotZero{
		baseOKNil: newBaseOKNil(3),
	}
}

func (z *tdNotZero) Match(ctx ctxerr.Context, got reflect.Value) (err *ctxerr.Error) {
	if got.IsValid() && !deepValueEqualOK(got, reflect.New(got.Type()).Elem()) {
		return nil
	}
	if ctx.BooleanError {
		return ctxerr.BooleanError
	}
	return ctx.CollectError(&ctxerr.Error{
		Message:  "zero value",
		Got:      got,
		Expected: z,
	})
}

func (z *tdNotZero) String() string {
	return "NotZero()"
}
