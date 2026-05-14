// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

import (
	"reflect"

	"github.com/maxatome/go-testdeep/internal/ctxerr"
	"github.com/maxatome/go-testdeep/internal/types"
)

type tdExpectedType struct {
	base
	expectedType reflect.Type
	isPtr        bool
}

func (t *tdExpectedType) errorTypeMismatch(gotType reflect.Type) *ctxerr.Error {
	expectedType := t.expectedType
	if t.isPtr {
		expectedType = reflect.PtrTo(expectedType)
	}
	return ctxerr.TypeMismatch(gotType, expectedType)
}

func (t *tdExpectedType) checkPtr(ctx ctxerr.Context, pGot *reflect.Value, nilAllowed bool) *ctxerr.Error {
	if t.isPtr {
		got := *pGot
		if got.Kind() != reflect.Ptr {
			if ctx.BooleanError {
				return ctxerr.BooleanError
			}
			return t.errorTypeMismatch(got.Type())
		}

		if !nilAllowed && got.IsNil() {
			if ctx.BooleanError {
				return ctxerr.BooleanError
			}
			return &ctxerr.Error{
				Message:  "values differ",
				Got:      got,
				Expected: types.RawString("non-nil"),
			}
		}

		*pGot = got.Elem()
	}
	return nil
}

func (t *tdExpectedType) checkType(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	if got.Type() != t.expectedType {
		if ctx.BeLax && t.expectedType.ConvertibleTo(got.Type()) {
			return nil
		}

		if ctx.BooleanError {
			return ctxerr.BooleanError
		}
		gt := got.Type()
		if t.isPtr {
			gt = reflect.PtrTo(gt)
		}
		return t.errorTypeMismatch(gt)
	}
	return nil
}

func (t *tdExpectedType) TypeBehind() reflect.Type {
	if t.err != nil {
		return nil
	}
	if t.isPtr {
		return reflect.New(t.expectedType).Type()
	}
	return t.expectedType
}

func (t *tdExpectedType) expectedTypeStr() string {
	if t.isPtr {
		return "*" + t.expectedType.String()
	}
	return t.expectedType.String()
}
