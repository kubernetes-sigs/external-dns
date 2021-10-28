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

type tdIsa struct {
	tdExpectedType
	checkImplement bool
}

var _ TestDeep = &tdIsa{}

// summary(Isa): checks the data type or whether data implements an
// interface or not
// input(Isa): bool,str,int,float,cplx,array,slice,map,struct,ptr,chan,func

// Isa operator checks the data type or whether data implements an
// interface or not.
//
// Typical type checks:
//
//   td.Cmp(t, time.Now(), td.Isa(time.Time{}))  // succeeds
//   td.Cmp(t, time.Now(), td.Isa(&time.Time{})) // fails, as not a *time.Time
//   td.Cmp(t, got, td.Isa(map[string]time.Time{}))
//
// For interfaces, it is a bit more complicated, as:
//
//   fmt.Stringer(nil)
//
// is not an interface, but just nil… To bypass this golang
// limitation, Isa accepts pointers on interfaces. So checking that
// data implements fmt.Stringer interface should be written as:
//
//   td.Cmp(t, bytes.Buffer{}, td.Isa((*fmt.Stringer)(nil))) // succeeds
//
// Of course, in the latter case, if checked data type is
// *fmt.Stringer, Isa will match too (in fact before checking whether
// it implements fmt.Stringer or not).
//
// TypeBehind method returns the reflect.Type of "model".
func Isa(model interface{}) TestDeep {
	modelType := reflect.TypeOf(model)
	i := tdIsa{
		tdExpectedType: tdExpectedType{
			base:         newBase(3),
			expectedType: modelType,
		},
	}

	if modelType == nil {
		i.err = ctxerr.OpBad("Isa", "Isa(nil) is not allowed. To check an interface, try Isa((*fmt.Stringer)(nil)), for fmt.Stringer for example")
		return &i
	}

	i.checkImplement = modelType.Kind() == reflect.Ptr &&
		modelType.Elem().Kind() == reflect.Interface
	return &i
}

func (i *tdIsa) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	if i.err != nil {
		return ctx.CollectError(i.err)
	}

	gotType := got.Type()

	if gotType == i.expectedType {
		return nil
	}

	if i.checkImplement {
		if gotType.Implements(i.expectedType.Elem()) {
			return nil
		}
	}

	if ctx.BooleanError {
		return ctxerr.BooleanError
	}
	return ctx.CollectError(i.errorTypeMismatch(gotType))
}

func (i *tdIsa) String() string {
	if i.err != nil {
		return i.stringError()
	}
	return i.expectedType.String()
}
