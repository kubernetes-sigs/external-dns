// Copyright (c) 2019, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

import (
	"reflect"

	"github.com/maxatome/go-testdeep/internal/ctxerr"
	"github.com/maxatome/go-testdeep/internal/util"
)

type tdLax struct {
	tdSmugglerBase
}

var _ TestDeep = &tdLax{}

// summary(Lax): temporarily enables [`BeLax` config flag]
// input(Lax): all

// Lax is a smuggler operator, it temporarily enables the BeLax config
// flag before letting the comparison process continue its course.
//
// It is more commonly used as CmpLax function than as an operator. It
// could be used when, for example, an operator is constructed once
// but applied to different, but compatible types as in:
//
//   bw := td.Between(20, 30)
//   intValue := 21
//   floatValue := 21.89
//   td.Cmp(t, intValue, bw)           // no need to be lax here: same int types
//   td.Cmp(t, floatValue, td.Lax(bw)) // be lax please, as float64 ≠ int
//
// Note that in the latter case, CmpLax() could be used as well:
//   td.CmpLax(t, floatValue, bw)
//
// TypeBehind method returns the greatest convertible or more common
// reflect.Type of "expectedValue" if it is a base type (bool, int*,
// uint*, float*, complex*, string), the reflect.Type of
// "expectedValue" otherwise, except if "expectedValue" is a TestDeep
// operator. In this case, it delegates TypeBehind() to the operator.
func Lax(expectedValue interface{}) TestDeep {
	c := tdLax{
		tdSmugglerBase: newSmugglerBase(expectedValue),
	}

	if !c.isTestDeeper {
		c.expectedValue = reflect.ValueOf(expectedValue)
	}
	return &c
}

func (l *tdLax) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	ctx.BeLax = true
	return deepValueEqual(ctx, got, l.expectedValue)
}

func (l *tdLax) HandleInvalid() bool {
	return true // Knows how to handle untyped nil values (aka invalid values)
}

func (l *tdLax) String() string {
	return "Lax(" + util.ToString(l.expectedValue) + ")"
}

func (l *tdLax) TypeBehind() reflect.Type {
	// If the expected value is a TestDeep operator, delegate TypeBehind to it
	if l.isTestDeeper {
		return l.expectedValue.Interface().(TestDeep).TypeBehind()
	}

	// For base types, returns the greatest convertible or more common one
	switch l.expectedValue.Kind() {
	case reflect.Invalid:
		return nil
	case reflect.Bool:
		return reflect.TypeOf(false)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return reflect.TypeOf(int64(0))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return reflect.TypeOf(uint64(0))
	case reflect.Float32, reflect.Float64:
		return reflect.TypeOf(float64(0))
	case reflect.Complex64, reflect.Complex128:
		return reflect.TypeOf(complex(128, -1))
	case reflect.String:
		return reflect.TypeOf("")
	default:
		return l.expectedValue.Type()
	}
}
