// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

import (
	"reflect"
	"strings"

	"github.com/maxatome/go-testdeep/helpers/tdutil"
	"github.com/maxatome/go-testdeep/internal/ctxerr"
	"github.com/maxatome/go-testdeep/internal/types"
	"github.com/maxatome/go-testdeep/internal/util"
)

type tdMapEach struct {
	baseOKNil
	expected reflect.Value
}

var _ TestDeep = &tdMapEach{}

// summary(MapEach): compares each map entry
// input(MapEach): map,ptr(ptr on map)

// MapEach operator has to be applied on maps. It compares each value
// of data map against expected value. During a match, all values have
// to match to succeed.
//
//   got := map[string]string{"test": "foo", "buzz": "bar"}
//   td.Cmp(t, got, td.MapEach("bar"))     // fails, coz "foo" ≠ "bar"
//   td.Cmp(t, got, td.MapEach(td.Len(3))) // succeeds as values are 3 chars long
func MapEach(expectedValue interface{}) TestDeep {
	return &tdMapEach{
		baseOKNil: newBaseOKNil(3),
		expected:  reflect.ValueOf(expectedValue),
	}
}

func (m *tdMapEach) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	if !got.IsValid() {
		if ctx.BooleanError {
			return ctxerr.BooleanError
		}
		return ctx.CollectError(&ctxerr.Error{
			Message:  "nil value",
			Got:      types.RawString("nil"),
			Expected: types.RawString("Map OR *Map"),
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
				Expected: types.RawString("Map OR *Map"),
			})
		}

		if gotElem.Kind() != reflect.Map {
			break
		}
		got = gotElem
		fallthrough

	case reflect.Map:
		var err *ctxerr.Error
		tdutil.MapEach(got, func(k, v reflect.Value) bool {
			err = deepValueEqual(ctx.AddMapKey(k), v, m.expected)
			return err == nil
		})
		return err
	}

	if ctx.BooleanError {
		return ctxerr.BooleanError
	}
	return ctx.CollectError(&ctxerr.Error{
		Message:  "bad type",
		Got:      types.RawString(got.Type().String()),
		Expected: types.RawString("Map OR *Map"),
	})
}

func (m *tdMapEach) String() string {
	const prefix = "MapEach("

	content := util.ToString(m.expected)
	if strings.Contains(content, "\n") {
		return prefix + util.IndentString(content, "        ") + ")"
	}
	return prefix + content + ")"
}
