// Copyright (c) 2019, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

import (
	"reflect"

	"github.com/maxatome/go-testdeep/helpers/tdutil"
	"github.com/maxatome/go-testdeep/internal/ctxerr"
	"github.com/maxatome/go-testdeep/internal/types"
	"github.com/maxatome/go-testdeep/internal/util"
)

type tdKVBase struct {
	tdSmugglerBase
}

func (b *tdKVBase) initKVBase(val interface{}) bool {
	b.tdSmugglerBase = newSmugglerBase(val, 1)

	if vval := reflect.ValueOf(val); vval.IsValid() {
		if b.isTestDeeper {
			return true
		}

		if vval.Kind() == reflect.Slice {
			b.expectedValue = vval
			return true
		}
	}
	return false
}

type tdKeys struct {
	tdKVBase
}

var _ TestDeep = &tdKeys{}

// summary(Keys): checks keys of a map
// input(Keys): map

// Keys is a smuggler operator. It takes a map and compares its
// ordered keys to "val".
//
// "val" can be a slice of items of the same type as the map keys:
//
//   got := map[string]bool{"c": true, "a": false, "b": true}
//   td.Cmp(t, got, td.Keys([]string{"a", "b", "c"})) // succeeds, keys sorted
//   td.Cmp(t, got, td.Keys([]string{"c", "a", "b"})) // fails as not sorted
//
// as well as an other operator as Bag, for example, to test keys in
// an unsorted manner:
//
//   got := map[string]bool{"c": true, "a": false, "b": true}
//   td.Cmp(t, got, td.Keys(td.Bag("c", "a", "b"))) // succeeds
func Keys(val interface{}) TestDeep {
	k := tdKeys{}
	if !k.initKVBase(val) {
		k.err = ctxerr.OpBadUsage("Keys", "(TESTDEEP_OPERATOR|SLICE)", val, 1, true)
	}
	return &k
}

func (k *tdKeys) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	if k.err != nil {
		return ctx.CollectError(k.err)
	}

	if got.Kind() != reflect.Map {
		if ctx.BooleanError {
			return ctxerr.BooleanError
		}
		return ctx.CollectError(&ctxerr.Error{
			Message:  "bad kind",
			Got:      types.RawString(got.Kind().String()),
			Expected: types.RawString(reflect.Map.String()),
		})
	}

	// Build a sorted slice of keys
	l := got.Len()
	keys := reflect.MakeSlice(reflect.SliceOf(got.Type().Key()), l, l)
	for i, k := range tdutil.MapSortedKeys(got) {
		keys.Index(i).Set(k)
	}
	return deepValueEqual(ctx.AddFunctionCall("keys"), keys, k.expectedValue)
}

func (k *tdKeys) String() string {
	if k.err != nil {
		return k.stringError()
	}
	if k.isTestDeeper {
		return "keys: " + k.expectedValue.Interface().(TestDeep).String()
	}
	return "keys=" + util.ToString(k.expectedValue.Interface())
}

type tdValues struct {
	tdKVBase
}

var _ TestDeep = &tdValues{}

// summary(Values): checks values of a map
// input(Values): map

// Values is a smuggler operator. It takes a map and compares its
// ordered values to "val".
//
// "val" can be a slice of items of the same type as the map values:
//
//   got := map[int]string{3: "c", 1: "a", 2: "b"}
//   td.Cmp(t, got, td.Values([]string{"a", "b", "c"})) // succeeds, values sorted
//   td.Cmp(t, got, td.Values([]string{"c", "a", "b"})) // fails as not sorted
//
// as well as an other operator as Bag, for example, to test values in
// an unsorted manner:
//
//   got := map[int]string{3: "c", 1: "a", 2: "b"}
//   td.Cmp(t, got, td.Values(td.Bag("c", "a", "b"))) // succeeds
func Values(val interface{}) TestDeep {
	v := tdValues{}
	if !v.initKVBase(val) {
		v.err = ctxerr.OpBadUsage("Values", "(TESTDEEP_OPERATOR|SLICE)", val, 1, true)
	}
	return &v
}

func (v *tdValues) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	if v.err != nil {
		return ctx.CollectError(v.err)
	}

	if got.Kind() != reflect.Map {
		if ctx.BooleanError {
			return ctxerr.BooleanError
		}
		return ctx.CollectError(&ctxerr.Error{
			Message:  "bad kind",
			Got:      types.RawString(got.Kind().String()),
			Expected: types.RawString(reflect.Map.String()),
		})
	}

	// Build a sorted slice of values
	l := got.Len()
	values := reflect.MakeSlice(reflect.SliceOf(got.Type().Elem()), l, l)
	for i, v := range tdutil.MapSortedValues(got) {
		values.Index(i).Set(v)
	}
	return deepValueEqual(ctx.AddFunctionCall("values"), values, v.expectedValue)
}

func (v *tdValues) String() string {
	if v.err != nil {
		return v.stringError()
	}
	if v.isTestDeeper {
		return "values: " + v.expectedValue.Interface().(TestDeep).String()
	}
	return "values=" + util.ToString(v.expectedValue.Interface())
}
