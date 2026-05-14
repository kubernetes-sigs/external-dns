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

type tdCode struct {
	base
	function reflect.Value
	argType  reflect.Type
	tParams  int
}

var _ TestDeep = &tdCode{}

// summary(Code): checks using a custom function
// input(Code): all

// Code operator allows to check data using a custom function. So
// fn is a function that must take one parameter whose type must be
// the same as the type of the compared value.
//
// fn can return a single bool kind value, telling that yes or no
// the custom test is successful:
//
//	td.Cmp(t, gotTime,
//	  td.Code(func(date time.Time) bool {
//	    return date.Year() == 2018
//	  }))
//
// or two values (bool, string) kinds. The bool value has the same
// meaning as above, and the string value is used to describe the
// test when it fails:
//
//	td.Cmp(t, gotTime,
//	  td.Code(func(date time.Time) (bool, string) {
//	    if date.Year() == 2018 {
//	      return true, ""
//	    }
//	    return false, "year must be 2018"
//	  }))
//
// or a single error value. If the returned error is nil, the test
// succeeded, else the error contains the reason of failure:
//
//	td.Cmp(t, gotJsonRawMesg,
//	  td.Code(func(b json.RawMessage) error {
//	    var c map[string]int
//	    err := json.Unmarshal(b, &c)
//	    if err != nil {
//	      return err
//	    }
//	    if c["test"] != 42 {
//	      return fmt.Errorf(`key "test" does not match 42`)
//	    }
//	    return nil
//	  }))
//
// This operator allows to handle any specific comparison not handled
// by standard operators.
//
// It is not recommended to call [Cmp] (or any other Cmp*
// functions or [*T] methods) inside the body of fn, because of
// confusion produced by output in case of failure. When the data
// needs to be transformed before being compared again, [Smuggle]
// operator should be used instead.
//
// But in some cases it can be better to handle yourself the
// comparison than to chain [TestDeep] operators. In this case, fn can
// be a function receiving one or two [*T] as first parameters and
// returning no values.
//
// When fn expects one [*T] parameter, it is directly derived from the
// [testing.TB] instance passed originally to [Cmp] (or its derivatives)
// using [NewT]:
//
//	td.Cmp(t, httpRequest, td.Code(func(t *td.T, r *http.Request) {
//	  token, err := DecodeToken(r.Header.Get("X-Token-1"))
//	  if t.CmpNoError(err) {
//	    t.True(token.OK())
//	  }
//	}))
//
// When fn expects two [*T] parameters, they are directly derived from
// the [testing.TB] instance passed originally to [Cmp] (or its derivatives)
// using [AssertRequire]:
//
//	td.Cmp(t, httpRequest, td.Code(func(assert, require *td.T, r *http.Request) {
//	  token, err := DecodeToken(r.Header.Get("X-Token-1"))
//	  require.CmpNoError(err)
//	  assert.True(token.OK())
//	}))
//
// Note that these forms do not work when there is no initial
// [testing.TB] instance, like when using [EqDeeplyError] or
// [EqDeeply] functions, or when the Code operator is called behind
// the following operators, as they just check if a match occurs
// without raising an error: [Any], [Bag], [Contains], [ContainsKey],
// [None], [Not], [NotAny], [Set], [SubBagOf], [SubSetOf],
// [SuperBagOf] and [SuperSetOf].
//
// RootName is inherited but not the current path, but it can be
// recovered if needed:
//
//	got := map[string]int{"foo": 123}
//	td.NewT(t).
//	  RootName("PIPO").
//	  Cmp(got, td.Map(map[string]int{}, td.MapEntries{
//	    "foo": td.Code(func(t *td.T, n int) {
//	      t.Cmp(n, 124)                                   // inherit only RootName
//	      t.RootName(t.Config.OriginalPath()).Cmp(n, 125) // recover current path
//	      t.RootName("").Cmp(n, 126)                      // undo RootName inheritance
//	    }),
//	  }))
//
// produces the following errors:
//
//	--- FAIL: TestCodeCustom (0.00s)
//	    td_code_test.go:339: Failed test
//	        PIPO: values differ             ← inherit only RootName
//	               got: 123
//	          expected: 124
//	    td_code_test.go:338: Failed test
//	        PIPO["foo"]: values differ      ← recover current path
//	               got: 123
//	          expected: 125
//	    td_code_test.go:342: Failed test
//	        DATA: values differ             ← undo RootName inheritance
//	               got: 123
//	          expected: 126
//
// TypeBehind method returns the [reflect.Type] of last parameter of fn.
func Code(fn any) TestDeep {
	vfn := reflect.ValueOf(fn)

	c := tdCode{
		base:     newBase(3),
		function: vfn,
	}

	if vfn.Kind() != reflect.Func {
		c.err = ctxerr.OpBadUsage("Code", "(FUNC)", fn, 1, true)
		return &c
	}
	if vfn.IsNil() {
		c.err = ctxerr.OpBad("Code", "Code(FUNC): FUNC cannot be a nil function")
		return &c
	}

	fnType := vfn.Type()
	in := fnType.NumIn()

	// We accept only:
	//   func (arg) bool
	//   func (arg) error
	//   func (arg) (bool, error)
	//   func (*td.T, arg) // with arg ≠ *td.T, as it is certainly an error
	//   func (assert, require *td.T, arg)
	if fnType.IsVariadic() || in == 0 || in > 3 ||
		(in > 1 && (fnType.In(0) != tType)) ||
		(in >= 2 && (in == 2) == (fnType.In(1) == tType)) {
		c.err = ctxerr.OpBad("Code",
			"Code(FUNC): FUNC must take only one non-variadic argument or (*td.T, arg) or (*td.T, *td.T, arg)")
		return &c
	}

	// func (arg) bool
	// func (arg) error
	// func (arg) (bool, error)
	if in == 1 {
		switch fnType.NumOut() {
		case 2: // (bool, *string*)
			if fnType.Out(1).Kind() != reflect.String {
				break
			}
			fallthrough

		case 1:
			// (*bool*) or (*bool*, string)
			if fnType.Out(0).Kind() == reflect.Bool ||
				// (*error*)
				(fnType.NumOut() == 1 && fnType.Out(0) == types.Error) {
				c.argType = fnType.In(0)
				return &c
			}
		}

		c.err = ctxerr.OpBad("Code",
			"Code(FUNC): FUNC must return bool or (bool, string) or error")
		return &c
	}

	// in == 2 || in == 3
	// func (*td.T, arg) (with arg ≠ *td.T)
	// func (assert, require *td.T, arg)

	if fnType.NumOut() != 0 {
		c.err = ctxerr.OpBad("Code", "Code(FUNC): FUNC must return nothing")
		return &c
	}

	c.tParams = in - 1
	c.argType = fnType.In(c.tParams)
	return &c
}

func (c *tdCode) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	if c.err != nil {
		return ctx.CollectError(c.err)
	}

	if !got.Type().AssignableTo(c.argType) {
		if !ctx.BeLax || !types.IsConvertible(got, c.argType) {
			if ctx.BooleanError {
				return ctxerr.BooleanError
			}
			return ctx.CollectError(&ctxerr.Error{
				Message:  "incompatible parameter type",
				Got:      types.RawString(got.Type().String()),
				Expected: types.RawString(c.argType.String()),
			})
		}
		got = got.Convert(c.argType)
	}

	// Refuse to override unexported fields access in this case. It is a
	// choice, as we think it is better to use Code() on surrounding
	// struct instead.
	if !got.CanInterface() {
		if ctx.BooleanError {
			return ctxerr.BooleanError
		}
		return ctx.CollectError(&ctxerr.Error{
			Message: "cannot compare unexported field",
			Summary: ctxerr.NewSummary("use Code() on surrounding struct instead"),
		})
	}

	if c.tParams == 0 {
		ret := c.function.Call([]reflect.Value{got})
		if ret[0].Kind() == reflect.Bool {
			if ret[0].Bool() {
				return nil
			}
		} else if ret[0].IsNil() { // reflect.Interface
			return nil
		}

		if ctx.BooleanError {
			return ctxerr.BooleanError
		}

		var reason string
		if len(ret) > 1 { // (bool, string)
			reason = ret[1].String()
		} else if ret[0].Kind() == reflect.Interface { // (error)
			// For internal use only
			if cErr, ok := ret[0].Interface().(*ctxerr.Error); ok {
				return ctx.CollectError(cErr)
			}
			reason = ret[0].Interface().(error).Error()
		}
		// else (bool) so no reason to report

		return ctx.CollectError(&ctxerr.Error{
			Message: "ran code with %% as argument",
			Summary: ctxerr.NewSummaryReason(got, reason),
		})
	}

	if ctx.OriginalTB == nil {
		return ctx.CollectError(&ctxerr.Error{
			Message: "cannot build *td.T instance",
			Summary: ctxerr.NewSummary("original testing.TB instance is missing"),
		})
	}

	t := NewT(ctx.OriginalTB)
	t.Config.forkedFromCtx = &ctx

	// func(*td.T, arg)
	if c.tParams == 1 {
		c.function.Call([]reflect.Value{
			reflect.ValueOf(t),
			got,
		})
		return nil
	}

	// func(assert, require *td.T, arg)
	assert, require := AssertRequire(t)
	c.function.Call([]reflect.Value{
		reflect.ValueOf(assert),
		reflect.ValueOf(require),
		got,
	})
	return nil
}

func (c *tdCode) String() string {
	if c.err != nil {
		return c.stringError()
	}
	return "Code(" + c.function.Type().String() + ")"
}

func (c *tdCode) TypeBehind() reflect.Type {
	if c.err != nil {
		return nil
	}
	return c.argType
}
