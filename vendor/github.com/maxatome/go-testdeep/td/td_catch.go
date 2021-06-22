// Copyright (c) 2019, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

import (
	"reflect"

	"github.com/maxatome/go-testdeep/internal/ctxerr"
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	"github.com/maxatome/go-testdeep/internal/util"
)

type tdCatch struct {
	tdSmugglerBase
	target reflect.Value
}

var _ TestDeep = &tdCatch{}

// summary(Catch): catches data on the fly before comparing it
// input(Catch): all

// Catch is a smuggler operator. It allows to copy data in "target" on
// the fly before comparing it as usual against "expectedValue".
//
// "target" must be a non-nil pointer and data should be assignable to
// its pointed type. If BeLax config flag is true or called under Lax
// (and so JSON) operator, data should be convertible to its pointer
// type.
//
//   var id int64
//   if td.Cmp(t, CreateRecord("test"),
//     td.JSON(`{"id": $1, "name": "test"}`, td.Catch(&id, td.NotZero()))) {
//     t.Logf("Created record ID is %d", id)
//   }
//
// It is really useful when used with JSON operator and/or tdhttp helper.
//
//   var id int64
//   ta := tdhttp.NewTestAPI(t, api.Handler).
//     PostJSON("/item", `{"name":"foo"}`).
//     CmpStatus(http.StatusCreated).
//     CmpJSONBody(td.JSON(`{"id": $1, "name": "foo"}`, td.Catch(&id, td.Gt(0))))
//   if !ta.Failed() {
//     t.Logf("Created record ID is %d", id)
//   }
//
// If you need to only catch data without comparing it, use Ignore
// operator as "expectedValue" as in:
//
//   var id int64
//   if td.Cmp(t, CreateRecord("test"),
//     td.JSON(`{"id": $1, "name": "test"}`, td.Catch(&id, td.Ignore()))) {
//     t.Logf("Created record ID is %d", id)
//   }
//
// TypeBehind method returns the reflect.Type of "expectedValue",
// except if "expectedValue" is a TestDeep operator. In this case, it
// delegates TypeBehind() to the operator, but if nil is returned by
// this call, the dereferenced reflect.Type of "target" is returned.
func Catch(target, expectedValue interface{}) TestDeep {
	vt := reflect.ValueOf(target)
	c := tdCatch{
		tdSmugglerBase: newSmugglerBase(expectedValue),
		target:         vt,
	}

	if vt.Kind() != reflect.Ptr || vt.IsNil() || !vt.Elem().CanSet() {
		c.err = ctxerr.OpBadUsage("Catch", "(NON_NIL_PTR, EXPECTED_VALUE)", target, 1, true)
		return &c
	}

	if !c.isTestDeeper {
		c.expectedValue = reflect.ValueOf(expectedValue)
	}
	return &c
}

func (c *tdCatch) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	if c.err != nil {
		return ctx.CollectError(c.err)
	}

	if targetType := c.target.Elem().Type(); !got.Type().AssignableTo(targetType) {
		if !ctx.BeLax || !got.Type().ConvertibleTo(targetType) {
			if ctx.BooleanError {
				return ctxerr.BooleanError
			}
			return ctx.CollectError(ctxerr.TypeMismatch(got.Type(), c.target.Elem().Type()))
		}

		c.target.Elem().Set(got.Convert(targetType))
	} else {
		c.target.Elem().Set(got)
	}

	return deepValueEqual(ctx, got, c.expectedValue)
}

func (c *tdCatch) String() string {
	if c.err != nil {
		return c.stringError()
	}

	if c.isTestDeeper {
		return c.expectedValue.Interface().(TestDeep).String()
	}
	return util.ToString(c.expectedValue)
}

func (c *tdCatch) TypeBehind() reflect.Type {
	if c.err != nil {
		return nil
	}

	if c.isTestDeeper {
		if typ := c.expectedValue.Interface().(TestDeep).TypeBehind(); typ != nil {
			return typ
		}
		// Operator unknown type behind, fallback on target dereferenced type
		return c.target.Type().Elem()
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	"github.com/maxatome/go-testdeep/internal/types"
||||||| parent of 5ce8c7613 (update vendored files)
	"github.com/maxatome/go-testdeep/internal/types"
=======
>>>>>>> 5ce8c7613 (update vendored files)
	"github.com/maxatome/go-testdeep/internal/util"
)

type tdCatch struct {
	tdSmugglerBase
	target reflect.Value
}

var _ TestDeep = &tdCatch{}

// summary(Catch): catches data on the fly before comparing it
// input(Catch): all

// Catch is a smuggler operator. It allows to copy data in "target" on
// the fly before comparing it as usual against "expectedValue".
//
// "target" must be a non-nil pointer and data should be assignable to
// its pointed type. If BeLax config flag is true or called under Lax
// (and so JSON) operator, data should be convertible to its pointer
// type.
//
//   var id int64
//   if td.Cmp(t, CreateRecord("test"),
//     td.JSON(`{"id": $1, "name": "test"}`, td.Catch(&id, td.NotZero()))) {
//     t.Logf("Created record ID is %d", id)
//   }
//
// It is really useful when used with JSON operator and/or tdhttp helper.
//
//   var id int64
//   ta := tdhttp.NewTestAPI(t, api.Handler).
//     PostJSON("/item", `{"name":"foo"}`).
//     CmpStatus(http.StatusCreated).
//     CmpJSONBody(td.JSON(`{"id": $1, "name": "foo"}`, td.Catch(&id, td.Gt(0))))
//   if !ta.Failed() {
//     t.Logf("Created record ID is %d", id)
//   }
//
// If you need to only catch data without comparing it, use Ignore
// operator as "expectedValue" as in:
//
//   var id int64
//   if td.Cmp(t, CreateRecord("test"),
//     td.JSON(`{"id": $1, "name": "test"}`, td.Catch(&id, td.Ignore()))) {
//     t.Logf("Created record ID is %d", id)
//   }
//
// TypeBehind method returns the reflect.Type of "expectedValue",
// except if "expectedValue" is a TestDeep operator. In this case, it
// delegates TypeBehind() to the operator, but if nil is returned by
// this call, the dereferenced reflect.Type of "target" is returned.
func Catch(target, expectedValue interface{}) TestDeep {
	vt := reflect.ValueOf(target)
	c := tdCatch{
		tdSmugglerBase: newSmugglerBase(expectedValue),
		target:         vt,
	}

	if vt.Kind() != reflect.Ptr || vt.IsNil() || !vt.Elem().CanSet() {
		c.err = ctxerr.OpBadUsage("Catch", "(NON_NIL_PTR, EXPECTED_VALUE)", target, 1, true)
		return &c
	}

	if !c.isTestDeeper {
		c.expectedValue = reflect.ValueOf(expectedValue)
	}
	return &c
}

func (c *tdCatch) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	if c.err != nil {
		return ctx.CollectError(c.err)
	}

	if targetType := c.target.Elem().Type(); !got.Type().AssignableTo(targetType) {
		if !ctx.BeLax || !got.Type().ConvertibleTo(targetType) {
			if ctx.BooleanError {
				return ctxerr.BooleanError
			}
			return ctx.CollectError(ctxerr.TypeMismatch(got.Type(), c.target.Elem().Type()))
		}

		c.target.Elem().Set(got.Convert(targetType))
	} else {
		c.target.Elem().Set(got)
	}

	return deepValueEqual(ctx, got, c.expectedValue)
}

func (c *tdCatch) String() string {
	if c.err != nil {
		return c.stringError()
	}

	if c.isTestDeeper {
		return c.expectedValue.Interface().(TestDeep).String()
	}
	return util.ToString(c.expectedValue)
}

func (c *tdCatch) TypeBehind() reflect.Type {
	if c.err != nil {
		return nil
	}

	if c.isTestDeeper {
<<<<<<< HEAD
		return c.expectedValue.Interface().(TestDeep).TypeBehind()
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
		return c.expectedValue.Interface().(TestDeep).TypeBehind()
=======
		if typ := c.expectedValue.Interface().(TestDeep).TypeBehind(); typ != nil {
			return typ
		}
		// Operator unknown type behind, fallback on target dereferenced type
		return c.target.Type().Elem()
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	"github.com/maxatome/go-testdeep/internal/types"
	"github.com/maxatome/go-testdeep/internal/util"
)

type tdCatch struct {
	tdSmugglerBase
	target reflect.Value
}

var _ TestDeep = &tdCatch{}

// summary(Catch): catches data on the fly before comparing it
// input(Catch): all

// Catch is a smuggler operator. It allows to copy data in "target" on
// the fly before comparing it as usual against "expectedValue".
//
// "target" must be a non-nil pointer and data should be assignable to
// its pointed type. If BeLax config flag is true or called under Lax
// (and so JSON) operator, data should be convertible to its pointer
// type.
//
//   var id int64
//   if td.Cmp(t, CreateRecord("test"),
//     td.JSON(`{"id": $1, "name": "test"}`, td.Catch(&id, td.NotZero()))) {
//     t.Logf("Created record ID is %d", id)
//   }
//
// It is really useful when used with JSON operator and/or tdhttp helper.
//
//   var id int64
//   ta := tdhttp.NewTestAPI(t, api.Handler).
//     PostJSON("/item", `{"name":"foo"}`).
//     CmpStatus(http.StatusCreated).
//     CmpJSONBody(td.JSON(`{"id": $1, "name": "foo"}`, td.Catch(&id, td.Gt(0))))
//   if !ta.Failed() {
//     t.Logf("Created record ID is %d", id)
//   }
//
// If you need to only catch data without comparing it, use Ignore
// operator as "expectedValue" as in:
//
//   var id int64
//   if td.Cmp(t, CreateRecord("test"),
//     td.JSON(`{"id": $1, "name": "test"}`, td.Catch(&id, td.Ignore()))) {
//     t.Logf("Created record ID is %d", id)
//   }
//
// TypeBehind method returns the reflect.Type of "expectedValue",
// except if "expectedValue" is a TestDeep operator. In this case, it
// delegates TypeBehind() to the operator, but if nil is returned by
// this call, the dereferenced reflect.Type of "target" is returned.
func Catch(target, expectedValue interface{}) TestDeep {
	vt := reflect.ValueOf(target)
	c := tdCatch{
		tdSmugglerBase: newSmugglerBase(expectedValue),
		target:         vt,
	}

	if vt.Kind() != reflect.Ptr || vt.IsNil() || !vt.Elem().CanSet() {
		c.err = ctxerr.OpBadUsage("Catch", "(NON_NIL_PTR, EXPECTED_VALUE)", target, 1, true)
		return &c
	}

	if !c.isTestDeeper {
		c.expectedValue = reflect.ValueOf(expectedValue)
	}
	return &c
}

func (c *tdCatch) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	if c.err != nil {
		return ctx.CollectError(c.err)
	}

	if targetType := c.target.Elem().Type(); !got.Type().AssignableTo(targetType) {
		if !ctx.BeLax || !types.IsConvertible(got, targetType) {
			if ctx.BooleanError {
				return ctxerr.BooleanError
			}
			return ctx.CollectError(ctxerr.TypeMismatch(got.Type(), c.target.Elem().Type()))
		}

		c.target.Elem().Set(got.Convert(targetType))
	} else {
		c.target.Elem().Set(got)
	}

	return deepValueEqual(ctx, got, c.expectedValue)
}

func (c *tdCatch) String() string {
	if c.err != nil {
		return c.stringError()
	}

	if c.isTestDeeper {
		return c.expectedValue.Interface().(TestDeep).String()
	}
	return util.ToString(c.expectedValue)
}

func (c *tdCatch) TypeBehind() reflect.Type {
	if c.err != nil {
		return nil
	}

	if c.isTestDeeper {
<<<<<<< HEAD
		return c.expectedValue.Interface().(TestDeep).TypeBehind()
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
		return c.expectedValue.Interface().(TestDeep).TypeBehind()
=======
		if typ := c.expectedValue.Interface().(TestDeep).TypeBehind(); typ != nil {
			return typ
		}
		// Operator unknown type behind, fallback on target dereferenced type
		return c.target.Type().Elem()
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	"github.com/maxatome/go-testdeep/internal/types"
	"github.com/maxatome/go-testdeep/internal/util"
)

type tdCatch struct {
	tdSmugglerBase
	target reflect.Value
}

var _ TestDeep = &tdCatch{}

// summary(Catch): catches data on the fly before comparing it
// input(Catch): all

// Catch is a smuggler operator. It allows to copy data in "target" on
// the fly before comparing it as usual against "expectedValue".
//
// "target" must be a non-nil pointer and data should be assignable to
// its pointed type. If BeLax config flag is true or called under Lax
// (and so JSON) operator, data should be convertible to its pointer
// type.
//
//   var id int64
//   if td.Cmp(t, CreateRecord("test"),
//     td.JSON(`{"id": $1, "name": "test"}`, td.Catch(&id, td.NotZero()))) {
//     t.Logf("Created record ID is %d", id)
//   }
//
// It is really useful when used with JSON operator and/or tdhttp helper.
//
//   var id int64
//   if tdhttp.CmpJSONResponse(t,
//     tdhttp.Post("/item", `{"name":"foo"}`),
//     api.Handler,
//     tdhttp.Response{
//       Status: http.StatusCreated,
//       Body: td.JSON(`{"id": $id, "name": "foo"}`,
//         td.Tag("id", td.Catch(&id, td.Gt(0)))),
//     }) {
//     t.Logf("Created record ID is %d", id)
//   }
//
// If you need to only catch data without comparing it, use Ignore
// operator as "expectedValue" as in:
//
//   var id int64
//   if td.Cmp(t, CreateRecord("test"),
//     td.JSON(`{"id": $1, "name": "test"}`, td.Catch(&id, td.Ignore()))) {
//     t.Logf("Created record ID is %d", id)
//   }
func Catch(target interface{}, expectedValue interface{}) TestDeep {
	vt := reflect.ValueOf(target)
	if vt.Kind() != reflect.Ptr || vt.IsNil() || !vt.Elem().CanSet() {
		panic("usage: Catch(NON_NIL_PTR, EXPECTED_VALUE)")
	}

	c := tdCatch{
		tdSmugglerBase: newSmugglerBase(expectedValue),
		target:         vt,
	}
	if !c.isTestDeeper {
		c.expectedValue = reflect.ValueOf(expectedValue)
	}
	return &c
}

func (c *tdCatch) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	if targetType := c.target.Elem().Type(); !got.Type().AssignableTo(targetType) {
		if !ctx.BeLax || !got.Type().ConvertibleTo(targetType) {
			if ctx.BooleanError {
				return ctxerr.BooleanError
			}
			return ctx.CollectError(&ctxerr.Error{
				Message:  "type mismatch",
				Got:      types.RawString(got.Type().String()),
				Expected: types.RawString(c.target.Elem().Type().String()),
			})
		}

		c.target.Elem().Set(got.Convert(targetType))
	} else {
		c.target.Elem().Set(got)
	}

	return deepValueEqual(ctx, got, c.expectedValue)
}

func (c *tdCatch) String() string {
	if c.isTestDeeper {
		return c.expectedValue.Interface().(TestDeep).String()
	}
	return util.ToString(c.expectedValue)
}

func (c *tdCatch) TypeBehind() reflect.Type {
	if c.isTestDeeper {
		return c.expectedValue.Interface().(TestDeep).TypeBehind()
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
	}
	if c.expectedValue.IsValid() {
		return c.expectedValue.Type()
	}
	return nil
}
