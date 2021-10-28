// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

import (
	"fmt"
	"math"
	"reflect"

	"github.com/maxatome/go-testdeep/internal/ctxerr"
	"github.com/maxatome/go-testdeep/internal/types"
)

type tdLenCapBase struct {
	tdSmugglerBase
}

func (b *tdLenCapBase) initLenCapBase(val interface{}) {
	b.tdSmugglerBase = newSmugglerBase(val, 1)

	// math.MaxInt appeared in go1.17
	const (
		maxUint = ^uint(0)
		maxInt  = int(maxUint >> 1)
		minInt  = -maxInt - 1
		usage   = "(TESTDEEP_OPERATOR|INT)"
	)

	if val == nil {
		b.err = ctxerr.OpBadUsage(b.GetLocation().Func, usage, val, 1, true)
		return
	}

	if b.isTestDeeper {
		return
	}

	vval := reflect.ValueOf(val)

	// A len or capacity is always an int, but accept any MinInt ≤ num ≤ MaxInt,
	// so it can be used in JSON, SubJSONOf and SuperJSONOf as float64
	switch vval.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		num := vval.Int()
		if num >= int64(minInt) && num <= int64(maxInt) {
			b.expectedValue = reflect.ValueOf(int(num))
			return
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		num := vval.Uint()
		if num <= uint64(maxInt) {
			b.expectedValue = reflect.ValueOf(int(num))
			return
		}
	case reflect.Float32, reflect.Float64:
		num := vval.Float()
		if num == math.Trunc(num) && num >= float64(minInt) && num <= float64(maxInt) {
			b.expectedValue = reflect.ValueOf(int(num))
			return
		}
	default:
		b.err = ctxerr.OpBadUsage(b.GetLocation().Func, usage, val, 1, true)
		return
	}

	op := b.GetLocation().Func
	b.err = ctxerr.OpBad(op, "usage: "+op+usage+
		", but received an out of bounds or not integer 1st parameter (%v), should be in int range", val)
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
	l.initLenCapBase(expectedLen)
	return &l
}

func (l *tdLen) String() string {
	if l.err != nil {
		return l.stringError()
	}
	if l.isTestDeeper {
		return "len: " + l.expectedValue.Interface().(TestDeep).String()
	}
	return fmt.Sprintf("len=%d", l.expectedValue.Int())
}

func (l *tdLen) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	if l.err != nil {
		return ctx.CollectError(l.err)
	}

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
	c.initLenCapBase(expectedCap)
	return &c
}

func (c *tdCap) String() string {
	if c.err != nil {
		return c.stringError()
	}
	if c.isTestDeeper {
		return "cap: " + c.expectedValue.Interface().(TestDeep).String()
	}
	return fmt.Sprintf("cap=%d", c.expectedValue.Int())
}

func (c *tdCap) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	if c.err != nil {
		return ctx.CollectError(c.err)
	}

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
