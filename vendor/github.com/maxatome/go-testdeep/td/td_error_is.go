// Copyright (c) 2022, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/maxatome/go-testdeep/internal/ctxerr"
	"github.com/maxatome/go-testdeep/internal/dark"
	"github.com/maxatome/go-testdeep/internal/types"
)

type tdErrorIs struct {
	tdSmugglerBase
	typeBehind reflect.Type
}

var _ TestDeep = &tdErrorIs{}

func errorToRawString(err error) types.RawString {
	if err == nil {
		return "nil"
	}
	return types.RawString(fmt.Sprintf("(%[1]T) %[1]q", err))
}

// summary(ErrorIs): checks the data is an error and matches a wrapped error
// input(ErrorIs): if(error)

// ErrorIs is a smuggler operator. It reports whether any error in an
// error's chain matches expectedError.
//
//	_, err := os.Open("/unknown/file")
//	td.Cmp(t, err, os.ErrNotExist)             // fails
//	td.Cmp(t, err, td.ErrorIs(os.ErrNotExist)) // succeeds
//
//	err1 := fmt.Errorf("failure1")
//	err2 := fmt.Errorf("failure2: %w", err1)
//	err3 := fmt.Errorf("failure3: %w", err2)
//	err := fmt.Errorf("failure4: %w", err3)
//	td.Cmp(t, err, td.ErrorIs(err))  // succeeds
//	td.Cmp(t, err, td.ErrorIs(err1)) // succeeds
//	td.Cmp(t, err1, td.ErrorIs(err)) // fails
//
//	var cerr myError
//	td.Cmp(t, err, td.ErrorIs(td.Catch(&cerr, td.String("my error..."))))
//
//	td.Cmp(t, err, td.ErrorIs(td.All(
//	  td.Isa(myError{}),
//	  td.String("my error..."),
//	)))
//
// Behind the scene it uses [errors.Is] function if expectedError is
// an [error] and [errors.As] function if expectedError is a
// [TestDeep] operator.
//
// Note that like [errors.Is], expectedError can be nil: in this case
// the comparison succeeds only when got is nil too.
//
// See also [CmpError] and [CmpNoError].
func ErrorIs(expectedError any) TestDeep {
	e := tdErrorIs{
		tdSmugglerBase: newSmugglerBase(expectedError),
	}

	switch expErr := expectedError.(type) {
	case nil:
	case error:
		e.expectedValue = reflect.ValueOf(expectedError)
	case TestDeep:
		e.typeBehind = expErr.TypeBehind()
		if e.typeBehind == nil {
			e.typeBehind = types.Interface
			break
		}
		if !e.typeBehind.Implements(types.Error) &&
			e.typeBehind.Kind() != reflect.Interface {
			e.err = ctxerr.OpBad("ErrorIs",
				"ErrorIs(%[1]s): type %[2]s behind %[1]s operator is not an interface or does not implement error",
				expErr.GetLocation().Func, e.typeBehind)
		}
	default:
		e.err = ctxerr.OpBadUsage("ErrorIs",
			"(error|TESTDEEP_OPERATOR)", expectedError, 1, false)
	}

	return &e
}

func (e *tdErrorIs) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	if e.err != nil {
		return ctx.CollectError(e.err)
	}

	// nil case
	if !got.IsValid() {
		// Special case
		if !e.expectedValue.IsValid() {
			return nil
		}
		if ctx.BooleanError {
			return ctxerr.BooleanError
		}
		return ctx.CollectError(&ctxerr.Error{
			Message:  "nil value",
			Got:      types.RawString("nil"),
			Expected: types.RawString("anything implementing error interface"),
		})
	}

	gotIf, ok := dark.GetInterface(got, true)
	if !ok {
		return ctx.CollectError(ctx.CannotCompareError())
	}

	gotErr, ok := gotIf.(error)
	if !ok {
		if ctx.BooleanError {
			return ctxerr.BooleanError
		}
		return ctx.CollectError(&ctxerr.Error{
			Message:  got.Type().String() + " does not implement error interface",
			Got:      gotIf,
			Expected: types.RawString("anything implementing error interface"),
		})
	}

	if e.isTestDeeper {
		target := reflect.New(e.typeBehind)

		if !errors.As(gotErr, target.Interface()) {
			if ctx.BooleanError {
				return ctxerr.BooleanError
			}
			return ctx.CollectError(&ctxerr.Error{
				Message:  "type is not found in err's tree",
				Got:      gotIf,
				Expected: types.RawString(e.typeBehind.String()),
			})
		}

		return deepValueEqual(ctx.AddCustomLevel(S(".ErrorIs(%s)", e.typeBehind)),
			target.Elem(), e.expectedValue)
	}

	var expErr error
	if e.expectedValue.IsValid() {
		expErr = e.expectedValue.Interface().(error)
		if errors.Is(gotErr, expErr) {
			return nil
		}
		if ctx.BooleanError {
			return ctxerr.BooleanError
		}
		return ctx.CollectError(&ctxerr.Error{
			Message:  "is not found in err's tree",
			Got:      errorToRawString(gotErr),
			Expected: errorToRawString(expErr),
		})
	}

	if ctx.BooleanError {
		return ctxerr.BooleanError
	}
	return ctx.CollectError(&ctxerr.Error{
		Message:  "is not nil",
		Got:      errorToRawString(gotErr),
		Expected: errorToRawString(expErr),
	})
}

func (e *tdErrorIs) String() string {
	if e.err != nil {
		return e.stringError()
	}
	if e.isTestDeeper {
		return "ErrorIs(" + e.expectedValue.Interface().(TestDeep).String() + ")"
	}
	if !e.expectedValue.IsValid() {
		return "ErrorIs(nil)"
	}
	return "ErrorIs(" + e.expectedValue.Interface().(error).Error() + ")"
}

func (e *tdErrorIs) HandleInvalid() bool {
	return true
}
