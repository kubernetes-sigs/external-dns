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
	"time"

	"github.com/maxatome/go-testdeep/internal/ctxerr"
	"github.com/maxatome/go-testdeep/internal/types"
	"github.com/maxatome/go-testdeep/internal/util"
)

type boundCmp uint8

const (
	boundNone boundCmp = iota
	boundIn
	boundOut
)

type tdBetween struct {
	base
	expectedMin reflect.Value
	expectedMax reflect.Value

	minBound boundCmp
	maxBound boundCmp
}

var _ TestDeep = &tdBetween{}

// BoundsKind type qualifies the "Between" bounds.
type BoundsKind uint8

const (
	// BoundsInIn allows to match between "from" and "to" both included
	BoundsInIn BoundsKind = iota
	// BoundsInOut allows to match between "from" included and "to" excluded
	BoundsInOut
	// BoundsOutIn allows to match between "from" excluded and "to" included
	BoundsOutIn
	// BoundsOutOut allows to match between "from" and "to" both excluded
	BoundsOutOut
)

type tdBetweenTime struct {
	tdBetween
	expectedType reflect.Type
	mustConvert  bool
}

var _ TestDeep = &tdBetweenTime{}

// summary(Between): checks that a number, string or time.Time is
// between two bounds
// input(Between): str,int,float,cplx(todo),struct(time.Time)

// Between operator checks that data is between "from" and
// "to". "from" and "to" can be any numeric, string or time.Time (or
// assignable) value. "from" and "to" must be the same kind as the
// compared value if numeric, and the same type if string or time.Time (or
// assignable). "bounds" allows to specify whether bounds are included
// or not:
//   - BoundsInIn (default): between "from" and "to" both included
//   - BoundsInOut: between "from" included and "to" excluded
//   - BoundsOutIn: between "from" excluded and "to" included
//   - BoundsOutOut: between "from" and "to" both excluded
//
// If "bounds" is missing, it defaults to BoundsInIn.
//
//   tc.Cmp(t, 17, td.Between(17, 20))               // succeeds, BoundsInIn by default
//   tc.Cmp(t, 17, td.Between(10, 17, BoundsInOut))  // fails
//   tc.Cmp(t, 17, td.Between(10, 17, BoundsOutIn))  // succeeds
//   tc.Cmp(t, 17, td.Between(17, 20, BoundsOutOut)) // fails
//
// TypeBehind method returns the reflect.Type of "from" (same as the "to" one.)
func Between(from interface{}, to interface{}, bounds ...BoundsKind) TestDeep {
	b := tdBetween{
		expectedMin: reflect.ValueOf(from),
		expectedMax: reflect.ValueOf(to),
	}

	const usage = "usage: Between(NUM|STRING|TIME, NUM|STRING|TIME[, BOUNDS_KIND])"

	if len(bounds) > 0 {
		if len(bounds) > 1 {
			panic(usage)
		}

		if bounds[0] == BoundsInIn || bounds[0] == BoundsInOut {
			b.minBound = boundIn
		} else {
			b.minBound = boundOut
		}

		if bounds[0] == BoundsInIn || bounds[0] == BoundsOutIn {
			b.maxBound = boundIn
		} else {
			b.maxBound = boundOut
		}
	} else {
		b.minBound = boundIn
		b.maxBound = boundIn
	}

	if b.expectedMax.Type() != b.expectedMin.Type() {
		panic("from and to params must have the same type")
	}

	return b.initBetween(usage)
}

func (b *tdBetween) initBetween(usage string) TestDeep {
	b.base = newBase(4)

	if !b.expectedMax.IsValid() {
		b.expectedMax = b.expectedMin
	}

	switch b.expectedMin.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if b.expectedMin.Int() > b.expectedMax.Int() {
			b.expectedMin, b.expectedMax = b.expectedMax, b.expectedMin
		}
		return b

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if b.expectedMin.Uint() > b.expectedMax.Uint() {
			b.expectedMin, b.expectedMax = b.expectedMax, b.expectedMin
		}
		return b

	case reflect.Float32, reflect.Float64:
		if b.expectedMin.Float() > b.expectedMax.Float() {
			b.expectedMin, b.expectedMax = b.expectedMax, b.expectedMin
		}
		return b

	case reflect.String:
		if b.expectedMin.String() > b.expectedMax.String() {
			b.expectedMin, b.expectedMax = b.expectedMax, b.expectedMin
		}
		return b

	case reflect.Struct:
		var bt tdBetweenTime
		if b.expectedMin.Type() == timeType {
			bt = tdBetweenTime{
				tdBetween:    *b,
				expectedType: timeType,
			}
		} else if b.expectedMin.Type().ConvertibleTo(timeType) {
			bt = tdBetweenTime{
				tdBetween:    *b,
				expectedType: b.expectedMin.Type(),
				mustConvert:  true,
			}
			bt.expectedMin = b.expectedMin.Convert(timeType)
			bt.expectedMax = b.expectedMax.Convert(timeType)
		} else {
			break
		}

		if bt.expectedMin.Interface().(time.Time).After(
			bt.expectedMax.Interface().(time.Time)) {
			bt.expectedMin, bt.expectedMax = bt.expectedMax, bt.expectedMin
		}

		return &bt
	}
	panic(usage)
}

func (b *tdBetween) nInt(tolerance reflect.Value) {
	if diff := tolerance.Int(); diff != 0 {
		expectedBase := b.expectedMin.Int()

		max := expectedBase + diff
		if max < expectedBase {
			max = math.MaxInt64
		}

		min := expectedBase - diff
		if min > expectedBase {
			min = math.MinInt64
		}

		b.expectedMin = reflect.New(tolerance.Type()).Elem()
		b.expectedMin.SetInt(min)

		b.expectedMax = reflect.New(tolerance.Type()).Elem()
		b.expectedMax.SetInt(max)
	}
}

func (b *tdBetween) nUint(tolerance reflect.Value) {
	if diff := tolerance.Uint(); diff != 0 {
		base := b.expectedMin.Uint()

		max := base + diff
		if max < base {
			max = math.MaxUint64
		}

		min := base - diff
		if min > base {
			min = 0
		}

		b.expectedMin = reflect.New(tolerance.Type()).Elem()
		b.expectedMin.SetUint(min)

		b.expectedMax = reflect.New(tolerance.Type()).Elem()
		b.expectedMax.SetUint(max)
	}
}

func (b *tdBetween) nFloat(tolerance reflect.Value) {
	if diff := tolerance.Float(); diff != 0 {
		base := b.expectedMin.Float()

		b.expectedMin = reflect.New(tolerance.Type()).Elem()
		b.expectedMin.SetFloat(base - diff)

		b.expectedMax = reflect.New(tolerance.Type()).Elem()
		b.expectedMax.SetFloat(base + diff)
	}
}

// summary(N): compares a number with a tolerance value
// input(N): int,float,cplx(todo)

// N operator compares a numeric data against "num" ± "tolerance". If
// "tolerance" is missing, it defaults to 0. "num" and "tolerance"
// must be the same kind as the compared value.
//
//   td.Cmp(t, 12.2, td.N(12., 0.3)) // succeeds
//   td.Cmp(t, 12.2, td.N(12., 0.1)) // fails
//
// TypeBehind method returns the reflect.Type of "num".
func N(num interface{}, tolerance ...interface{}) TestDeep {
	n := tdBetween{
		base:        newBase(3),
		expectedMin: reflect.ValueOf(num),
		minBound:    boundIn,
		maxBound:    boundIn,
	}

	const usage = "usage: N({,U}INT{,8,16,32,64}|FLOAT{32,64}[, TOLERANCE])"

	switch n.expectedMin.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
	default:
		panic(usage)
	}

	n.expectedMax = n.expectedMin

	if len(tolerance) > 0 {
		if len(tolerance) > 1 {
			panic(usage)
		}

		tol := reflect.ValueOf(tolerance[0])
		if tol.Type() != n.expectedMin.Type() {
			panic("tolerance param must have the same type as num one")
		}

		switch tol.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			n.nInt(tol)

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
			reflect.Uint64:
			n.nUint(tol)

		default: // case reflect.Float32, reflect.Float64:
			n.nFloat(tol)
		}
	}

	return &n
}

// summary(Gt): checks that a number, string or time.Time is
// greater than a value
// input(Gt): str,int,float,cplx(todo),struct(time.Time)

// Gt operator checks that data is greater than
// "minExpectedValue". "minExpectedValue" can be any numeric or
// time.Time (or assignable) value. "minExpectedValue" must be the
// same kind as the compared value if numeric, and the same type if
// time.Time (or assignable).
//
//   td.Cmp(t, 17, td.Gt(15))
//   before := time.Now()
//   td.Cmp(t, time.Now(), td.Gt(before))
//
// TypeBehind method returns the reflect.Type of "minExpectedValue".
func Gt(minExpectedValue interface{}) TestDeep {
	b := &tdBetween{
		expectedMin: reflect.ValueOf(minExpectedValue),
		minBound:    boundOut,
	}
	return b.initBetween("usage: Gt(NUM|STRING|TIME)")
}

// summary(Gte): checks that a number, string or time.Time is
// greater or equal than a value
// input(Gte): str,int,float,cplx(todo),struct(time.Time)

// Gte operator checks that data is greater or equal than
// "minExpectedValue". "minExpectedValue" can be any numeric or
// time.Time (or assignable) value. "minExpectedValue" must be the
// same kind as the compared value if numeric, and the same type if
// time.Time (or assignable).
//
//   td.Cmp(t, 17, td.Gte(17))
//   before := time.Now()
//   td.Cmp(t, time.Now(), td.Gte(before))
//
// TypeBehind method returns the reflect.Type of "minExpectedValue".
func Gte(minExpectedValue interface{}) TestDeep {
	b := &tdBetween{
		expectedMin: reflect.ValueOf(minExpectedValue),
		minBound:    boundIn,
	}
	return b.initBetween("usage: Gte(NUM|STRING|TIME)")
}

// summary(Lt): checks that a number, string or time.Time is
// lesser than a value
// input(Lt): str,int,float,cplx(todo),struct(time.Time)

// Lt operator checks that data is lesser than
// "maxExpectedValue". "maxExpectedValue" can be any numeric or
// time.Time (or assignable) value. "maxExpectedValue" must be the
// same kind as the compared value if numeric, and the same type if
// time.Time (or assignable).
//
//   td.Cmp(t, 17, td.Lt(19))
//   before := time.Now()
//   td.Cmp(t, before, td.Lt(time.Now()))
//
// TypeBehind method returns the reflect.Type of "maxExpectedValue".
func Lt(maxExpectedValue interface{}) TestDeep {
	b := &tdBetween{
		expectedMin: reflect.ValueOf(maxExpectedValue),
		maxBound:    boundOut,
	}
	return b.initBetween("usage: Lt(NUM|STRING|TIME)")
}

// summary(Lte): checks that a number, string or time.Time is
// lesser or equal than a value
// input(Lte): str,int,float,cplx(todo),struct(time.Time)

// Lte operator checks that data is lesser or equal than
// "maxExpectedValue". "maxExpectedValue" can be any numeric or
// time.Time (or assignable) value. "maxExpectedValue" must be the
// same kind as the compared value if numeric, and the same type if
// time.Time (or assignable).
//
//   td.Cmp(t, 17, td.Lte(17))
//   before := time.Now()
//   td.Cmp(t, before, td.Lt(time.Now()))
//
// TypeBehind method returns the reflect.Type of "maxExpectedValue".
func Lte(maxExpectedValue interface{}) TestDeep {
	b := &tdBetween{
		expectedMin: reflect.ValueOf(maxExpectedValue),
		maxBound:    boundIn,
	}
	return b.initBetween("usage: Lte(NUM|STRING|TIME)")
}

func (b *tdBetween) matchInt(got reflect.Value) (ok bool) {
	switch b.minBound {
	case boundIn:
		ok = got.Int() >= b.expectedMin.Int()
	case boundOut:
		ok = got.Int() > b.expectedMin.Int()
	default:
		ok = true
	}
	if ok {
		switch b.maxBound {
		case boundIn:
			ok = got.Int() <= b.expectedMax.Int()
		case boundOut:
			ok = got.Int() < b.expectedMax.Int()
		default:
			ok = true
		}
	}
	return
}

func (b *tdBetween) matchUint(got reflect.Value) (ok bool) {
	switch b.minBound {
	case boundIn:
		ok = got.Uint() >= b.expectedMin.Uint()
	case boundOut:
		ok = got.Uint() > b.expectedMin.Uint()
	default:
		ok = true
	}
	if ok {
		switch b.maxBound {
		case boundIn:
			ok = got.Uint() <= b.expectedMax.Uint()
		case boundOut:
			ok = got.Uint() < b.expectedMax.Uint()
		default:
			ok = true
		}
	}
	return
}

func (b *tdBetween) matchFloat(got reflect.Value) (ok bool) {
	switch b.minBound {
	case boundIn:
		ok = got.Float() >= b.expectedMin.Float()
	case boundOut:
		ok = got.Float() > b.expectedMin.Float()
	default:
		ok = true
	}
	if ok {
		switch b.maxBound {
		case boundIn:
			ok = got.Float() <= b.expectedMax.Float()
		case boundOut:
			ok = got.Float() < b.expectedMax.Float()
		default:
			ok = true
		}
	}
	return
}

func (b *tdBetween) matchString(got reflect.Value) (ok bool) {
	switch b.minBound {
	case boundIn:
		ok = got.String() >= b.expectedMin.String()
	case boundOut:
		ok = got.String() > b.expectedMin.String()
	default:
		ok = true
	}
	if ok {
		switch b.maxBound {
		case boundIn:
			ok = got.String() <= b.expectedMax.String()
		case boundOut:
			ok = got.String() < b.expectedMax.String()
		default:
			ok = true
		}
	}
	return
}

func (b *tdBetween) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	if got.Type() != b.expectedMin.Type() {
		if ctx.BeLax && b.expectedMin.Type().ConvertibleTo(got.Type()) {
			nb := *b
			nb.expectedMin = b.expectedMin.Convert(got.Type())
			nb.expectedMax = b.expectedMax.Convert(got.Type())
			b = &nb
		} else {
			if ctx.BooleanError {
				return ctxerr.BooleanError
			}
			return ctx.CollectError(&ctxerr.Error{
				Message:  "type mismatch",
				Got:      types.RawString(got.Type().String()),
				Expected: types.RawString(b.expectedMin.Type().String()),
			})
		}
	}

	var ok bool

	switch got.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		ok = b.matchInt(got)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		ok = b.matchUint(got)

	case reflect.Float32, reflect.Float64:
		ok = b.matchFloat(got)

	case reflect.String:
		ok = b.matchString(got)
	}

	if ok {
		return nil
	}

	if ctx.BooleanError {
		return ctxerr.BooleanError
	}

	var gotStr types.RawString
	if got.Kind() == reflect.String {
		gotStr = types.RawString(util.ToString(got))
	} else {
		gotStr = types.RawString(fmt.Sprintf("%v", got))
	}

	return ctx.CollectError(&ctxerr.Error{
		Message:  "values differ",
		Got:      gotStr,
		Expected: types.RawString(b.String()),
	})
}

func (b *tdBetween) String() string {
	var (
		min, max       interface{}
		minStr, maxStr string
	)

	if b.minBound != boundNone {
		min = b.expectedMin.Interface()

		// We want strings be double-quoted
		if b.expectedMin.Kind() == reflect.String {
			minStr = util.ToString(min)
		} else {
			minStr = fmt.Sprintf("%v", min)
		}
	}
	if b.maxBound != boundNone {
		max = b.expectedMax.Interface()

		// We want strings be double-quoted
		if b.expectedMax.Kind() == reflect.String {
			maxStr = util.ToString(max)
		} else {
			maxStr = fmt.Sprintf("%v", max)
		}
	}

	if min == max {
		return minStr
	}

	if min != nil {
		if max != nil {
			return fmt.Sprintf("%v %c got %c %v",
				minStr,
				util.TernRune(b.minBound == boundIn, '≤', '<'),
				util.TernRune(b.maxBound == boundIn, '≤', '<'),
				maxStr)
		}

		return fmt.Sprintf("%c %v",
			util.TernRune(b.minBound == boundIn, '≥', '>'), minStr)
	}

	return fmt.Sprintf("%c %v",
		util.TernRune(b.maxBound == boundIn, '≤', '<'), maxStr)
}

func (b *tdBetween) TypeBehind() reflect.Type {
	return b.expectedMin.Type()
}

var _ TestDeep = &tdBetweenTime{}

func (b *tdBetweenTime) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	if got.Type() != b.expectedType {
		if ctx.BeLax && got.Type().ConvertibleTo(b.expectedType) {
			got = got.Convert(b.expectedType)
		} else {
			if ctx.BooleanError {
				return ctxerr.BooleanError
			}
			return ctx.CollectError(&ctxerr.Error{
				Message:  "type mismatch",
				Got:      types.RawString(got.Type().String()),
				Expected: types.RawString(b.expectedType.String()),
			})
		}
	}

	cmpGot, err := getTime(ctx, got, b.mustConvert)
	if err != nil {
		return ctx.CollectError(err)
	}

	var ok bool
	if b.minBound != boundNone {
		min := b.expectedMin.Interface().(time.Time)

		if b.minBound == boundIn {
			ok = !min.After(cmpGot)
		} else {
			ok = cmpGot.After(min)
		}
	} else {
		ok = true
	}

	if ok && b.maxBound != boundNone {
		max := b.expectedMax.Interface().(time.Time)

		if b.maxBound == boundIn {
			ok = !max.Before(cmpGot)
		} else {
			ok = cmpGot.Before(max)
		}
	}

	if ok {
		return nil
	}

	if ctx.BooleanError {
		return ctxerr.BooleanError
	}
	return ctx.CollectError(&ctxerr.Error{
		Message:  "values differ",
		Got:      types.RawString(fmt.Sprintf("%v", got.Interface())),
		Expected: types.RawString(b.String()),
	})
}

func (b *tdBetweenTime) TypeBehind() reflect.Type {
	return b.expectedType
}
