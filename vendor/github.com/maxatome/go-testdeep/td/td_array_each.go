// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

import (
	"reflect"
	"strings"

	"github.com/maxatome/go-testdeep/internal/ctxerr"
	"github.com/maxatome/go-testdeep/internal/types"
	"github.com/maxatome/go-testdeep/internal/util"
)

type tdArrayEach struct {
	baseOKNil
	expected reflect.Value
}

var _ TestDeep = &tdArrayEach{}

// summary(ArrayEach): compares each array or slice item
// input(ArrayEach): array,slice,ptr(ptr on array/slice)

// ArrayEach operator has to be applied on arrays or slices or on
// pointers on array/slice. It compares each item of data array/slice
// against "expectedValue". During a match, all items have to match to
// succeed.
//
//   got := [3]string{"foo", "bar", "biz"}
//   td.Cmp(t, got, td.ArrayEach(td.Len(3)))         // succeeds
//   td.Cmp(t, got, td.ArrayEach(td.HasPrefix("b"))) // fails coz "foo"
//
// Works on slices as well:
//
//   got := []Person{
//     {Name: "Bob", Age: 42},
//     {Name: "Alice", Age: 24},
//   }
//   td.Cmp(t, got, td.ArrayEach(
//     td.Struct(Person{}, td.StructFields{
//       Age: td.Between(20, 45),
//     })),
//   ) // succeeds, each Person has Age field between 20 and 45
func ArrayEach(expectedValue interface{}) TestDeep {
	return &tdArrayEach{
		baseOKNil: newBaseOKNil(3),
		expected:  reflect.ValueOf(expectedValue),
	}
}

func (a *tdArrayEach) Match(ctx ctxerr.Context, got reflect.Value) (err *ctxerr.Error) {
	if !got.IsValid() {
		if ctx.BooleanError {
			return ctxerr.BooleanError
		}
		return ctx.CollectError(&ctxerr.Error{
			Message:  "nil value",
			Got:      types.RawString("nil"),
			Expected: types.RawString("Slice OR Array OR *Slice OR *Array"),
		})
	}

	switch got.Kind() {
	case reflect.Ptr:
		gotElem := got.Elem()
		if !gotElem.IsValid() {
			if ctx.BooleanError {
				return ctxerr.BooleanError
			}
			return ctx.CollectError(&ctxerr.Error{
				Message:  "nil pointer",
				Got:      types.RawString("nil " + got.Type().String()),
				Expected: types.RawString("Slice OR Array OR *Slice OR *Array"),
			})
		}

		if gotElem.Kind() != reflect.Array && gotElem.Kind() != reflect.Slice {
			break
		}
		got = gotElem
		fallthrough

	case reflect.Array, reflect.Slice:
		gotLen := got.Len()

		var err *ctxerr.Error
		for idx := 0; idx < gotLen; idx++ {
			err = deepValueEqual(ctx.AddArrayIndex(idx), got.Index(idx), a.expected)
			if err != nil {
				return err
			}
		}
		return nil
	}

	if ctx.BooleanError {
		return ctxerr.BooleanError
	}
	return ctx.CollectError(&ctxerr.Error{
		Message:  "bad type",
		Got:      types.RawString(got.Type().String()),
		Expected: types.RawString("Slice OR Array OR *Slice OR *Array"),
	})
}

func (a *tdArrayEach) String() string {
	const prefix = "ArrayEach("

	content := util.ToString(a.expected)
	if strings.Contains(content, "\n") {
		return prefix + util.IndentString(content, "          ") + ")"
	}
	return prefix + content + ")"
}
