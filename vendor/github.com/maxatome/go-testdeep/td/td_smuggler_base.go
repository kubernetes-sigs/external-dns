// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/maxatome/go-testdeep/internal/ctxerr"
)

// tdSmugglerBase is the base class of all smuggler TestDeep operators.
type tdSmugglerBase struct {
	base
	expectedValue reflect.Value
	isTestDeeper  bool
}

func newSmugglerBase(val interface{}, depth ...int) (ret tdSmugglerBase) {
	callDepth := 4
	if len(depth) > 0 {
		callDepth += depth[0]
	}
	ret.base = newBase(callDepth)

	// Initializes only if TestDeep operator. Other cases are specific.
	if _, ok := val.(TestDeep); ok {
		ret.expectedValue = reflect.ValueOf(val)
		ret.isTestDeeper = true
	}
	return
}

// internalTypeBehind returns the type behind expectedValue or nil if
// it cannot be determined.
func (s *tdSmugglerBase) internalTypeBehind() reflect.Type {
	if s.isTestDeeper {
		return s.expectedValue.Interface().(TestDeep).TypeBehind()
	}
	if s.expectedValue.IsValid() {
		return s.expectedValue.Type()
	}
	return nil
}

// jsonValueEqual compares "got" to expectedValue, trying to do it
// using a JSON point of view. It is the caller responsibility to
// ensure that "got" value is either a bool, float64, string,
// []interface{}, a map[string]interface{} or simply nil.
//
// If the type behind expectedValue can be determined and is different
// from "got" type, "got" value is JSON marshaled, then unmarshaled
// in a new value of this type. This new value is then compared to
// expectedValue.
//
// Otherwise, "got" value is compared as-is to expectedValue.
func (s *tdSmugglerBase) jsonValueEqual(ctx ctxerr.Context, got interface{}) *ctxerr.Error {
	expectedType := s.internalTypeBehind()

	// Unknown expected type (operator with nil TypeBehind() result or
	// untyped nil), lets deepValueEqual() handles the comparison using
	// BeLax flag
	if expectedType == nil {
		return deepValueEqual(ctx, reflect.ValueOf(got), s.expectedValue)
	}

	// Same type for got & expected type, no need to Marshal/Unmarshal
	if got != nil && expectedType == reflect.TypeOf(got) {
		return deepValueEqual(ctx, reflect.ValueOf(got), s.expectedValue)
	}

	// Unmarshal got into the expectedType
	b, _ := json.Marshal(got) // No error can occur here

	finalGot := reflect.New(expectedType)
	if err := json.Unmarshal(b, finalGot.Interface()); err != nil {
		if ctx.BooleanError {
			return ctxerr.BooleanError
		}
		return ctx.CollectError(&ctxerr.Error{
			Message: fmt.Sprintf(
				"an error occurred while unmarshalling JSON into %s", expectedType),
			Summary: ctxerr.NewSummary(err.Error()),
		})
	}

	return deepValueEqual(ctx, finalGot.Elem(), s.expectedValue)
}
