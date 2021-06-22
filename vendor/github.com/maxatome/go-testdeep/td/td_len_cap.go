// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

import (
	"fmt"
	"reflect"

	"github.com/maxatome/go-testdeep/internal/ctxerr"
	"github.com/maxatome/go-testdeep/internal/types"
)

type tdLenCapBase struct {
	tdSmugglerBase
}

func (b *tdLenCapBase) initLenCapBase(val interface{}) bool {
	vval := reflect.ValueOf(val)
	if vval.IsValid() {
		b.tdSmugglerBase = newSmugglerBase(val, 5)

		if b.isTestDeeper {
			return true
		}

		// A len or capacity is always an int
		if vval.Type() == intType {
			b.expectedValue = vval
			return true
		}
	}
	return false
}

func (b *tdLenCapBase) isEqual(ctx ctxerr.Context, got int) (bool, *ctxerr.Error) {
	if b.isTestDeeper {
		return true, deepValueEqual(ctx, reflect.ValueOf(got), b.expectedValue)
	}

	if int64(got) == b.expectedValue.Int() {
		return true, nil
	}

	return false, nil
}

type tdLen struct {
	tdLenCapBase
}

var _ TestDeep = &tdLen{}

// summary(Len): checks an array, slice, map, string or channel length
// input(Len): array,slice,map,chan

// Len is a smuggler operator. It takes data, applies len() function
// on it and compares its result to "expectedLen". Of course, the
// compared value must be an array, a channel, a map, a slice or a
// string.
//
// "expectedLen" can be an int value:
//
//   td.Cmp(t, gotSlice, td.Len(12))
//
// as well as an other operator:
//
//   td.Cmp(t, gotSlice, td.Len(td.Between(3, 4)))
func Len(expectedLen interface{}) TestDeep {
	l := tdLen{}
	if l.initLenCapBase(expectedLen) {
		return &l
	}
	panic("usage: Len(TESTDEEP_OPERATOR|INT)")
}

func (l *tdLen) String() string {
	if l.isTestDeeper {
		return "len: " + l.expectedValue.Interface().(TestDeep).String()
	}
	return fmt.Sprintf("len=%d", l.expectedValue.Int())
}

func (l *tdLen) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	switch got.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		ret, err := l.isEqual(ctx.AddFunctionCall("len"), got.Len())
		if ret {
			return err
		}
		if ctx.BooleanError {
			return ctxerr.BooleanError
		}
		return ctx.CollectError(&ctxerr.Error{
			Message:  "bad length",
			Got:      types.RawInt(got.Len()),
			Expected: types.RawInt(l.expectedValue.Int()),
		})

	default:
		if ctx.BooleanError {
			return ctxerr.BooleanError
		}
		return ctx.CollectError(&ctxerr.Error{
			Message:  "bad type",
			Got:      types.RawString(got.Type().String()),
			Expected: types.RawString("Array, Chan, Map, Slice or string"),
		})
	}
}

type tdCap struct {
	tdLenCapBase
}

var _ TestDeep = &tdCap{}

// summary(Cap): checks an array, slice or channel capacity
// input(Cap): array,slice,chan

// Cap is a smuggler operator. It takes data, applies cap() function
// on it and compares its result to "expectedCap". Of course, the
// compared value must be an array, a channel or a slice.
//
// "expectedCap" can be an int value:
//
//   td.Cmp(t, gotSlice, td.Cap(12))
//
// as well as an other operator:
//
//   td.Cmp(t, gotSlice, td.Cap(td.Between(3, 4)))
func Cap(expectedCap interface{}) TestDeep {
	c := tdCap{}
	if c.initLenCapBase(expectedCap) {
		return &c
	}
	panic("usage: Cap(TESTDEEP_OPERATOR|INT)")
}

func (c *tdCap) String() string {
	if c.isTestDeeper {
		return "cap: " + c.expectedValue.Interface().(TestDeep).String()
	}
	return fmt.Sprintf("cap=%d", c.expectedValue.Int())
}

func (c *tdCap) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	switch got.Kind() {
	case reflect.Array, reflect.Chan, reflect.Slice:
		ret, err := c.isEqual(ctx.AddFunctionCall("cap"), got.Cap())
		if ret {
			return err
		}
		if ctx.BooleanError {
			return ctxerr.BooleanError
		}
		return ctx.CollectError(&ctxerr.Error{
			Message:  "bad capacity",
			Got:      types.RawInt(got.Cap()),
			Expected: types.RawInt(c.expectedValue.Int()),
		})

	default:
		if ctx.BooleanError {
			return ctxerr.BooleanError
		}
		return ctx.CollectError(&ctxerr.Error{
			Message:  "bad type",
			Got:      types.RawString(got.Type().String()),
			Expected: types.RawString("Array, Chan or Slice"),
		})
	}
}
