// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/maxatome/go-testdeep/helpers/tdutil"
	"github.com/maxatome/go-testdeep/internal/ctxerr"
	"github.com/maxatome/go-testdeep/internal/dark"
	"github.com/maxatome/go-testdeep/internal/util"
)

type mapKind uint8

const (
	allMap mapKind = iota
	subMap
	superMap
)

type tdMap struct {
	tdExpectedType
	expectedEntries []mapEntryInfo
	kind            mapKind
}

var _ TestDeep = &tdMap{}

type mapEntryInfo struct {
	key      reflect.Value
	expected reflect.Value
}

// MapEntries allows to pass map entries to check in functions [Map],
// [SubMapOf] and [SuperMapOf]. It is a map whose each key is the
// expected entry key and the corresponding value the expected entry
// value (which can be a [TestDeep] operator as well as a zero value.)
type MapEntries map[any]any

func mergeMapEntries(mes ...MapEntries) MapEntries {
	switch len(mes) {
	case 0:
		return nil
	case 1:
		return mes[0]
	}

	ret := make(MapEntries, len(mes[0]))
	for _, me := range mes {
		for k, v := range me {
			ret[k] = v
		}
	}
	return ret
}

func newMap(model any, kind mapKind, mes ...MapEntries) *tdMap {
	vmodel := reflect.ValueOf(model)

	m := tdMap{
		tdExpectedType: tdExpectedType{
			base: newBase(4),
		},
		kind: kind,
	}

	switch vmodel.Kind() {
	case reflect.Ptr:
		if vmodel.Type().Elem().Kind() != reflect.Map {
			break
		}

		m.isPtr = true

		if vmodel.IsNil() {
			m.expectedType = vmodel.Type().Elem()
			m.populateExpectedEntries(mergeMapEntries(mes...), reflect.Value{})
			return &m
		}

		vmodel = vmodel.Elem()
		fallthrough

	case reflect.Map:
		m.expectedType = vmodel.Type()
		m.populateExpectedEntries(mergeMapEntries(mes...), vmodel)
		return &m
	}

	m.err = ctxerr.OpBadUsage(
		m.GetLocation().Func, "(MAP|&MAP, EXPECTED_ENTRIES)",
		model, 1, true)
	return &m
}

func (m *tdMap) populateExpectedEntries(entries MapEntries, expectedModel reflect.Value) {
	var keysInModel int
	if expectedModel.IsValid() {
		keysInModel = expectedModel.Len()
	}

	m.expectedEntries = make([]mapEntryInfo, 0, keysInModel+len(entries))
	checkedEntries := make(map[any]bool, len(entries))

	keyType := m.expectedType.Key()
	valueType := m.expectedType.Elem()

	var entryInfo mapEntryInfo

	for key, expectedValue := range entries {
		vkey := reflect.ValueOf(key)
		if !vkey.Type().AssignableTo(keyType) {
			m.err = ctxerr.OpBad(
				m.GetLocation().Func,
				"expected key %s type mismatch: %s != model key type (%s)",
				util.ToString(key),
				vkey.Type(),
				keyType)
			return
		}

		if expectedValue == nil {
			switch valueType.Kind() {
			case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map,
				reflect.Ptr, reflect.Slice:
				entryInfo.expected = reflect.Zero(valueType) // change to a typed nil
			default:
				entryInfo.expected = reflect.Value{}
				// Don't raise an error if map value cannot be nil as a
				// smuggle hook can change it at fly during the comparison
			}
		} else {
			entryInfo.expected = reflect.ValueOf(expectedValue)
			// Don't check vexpectedValue type against map value one as a
			// smuggle hook can change it at fly during the comparison
		}

		entryInfo.key = vkey
		m.expectedEntries = append(m.expectedEntries, entryInfo)
		checkedEntries[dark.MustGetInterface(vkey)] = true
	}

	// Check entries in model
	if keysInModel == 0 {
		return
	}

	tdutil.MapEach(expectedModel, func(k, v reflect.Value) bool {
		entryInfo.expected = v

		if checkedEntries[dark.MustGetInterface(k)] {
			m.err = ctxerr.OpBad(
				m.GetLocation().Func,
				"%s entry exists in both model & expectedEntries",
				util.ToString(k))
			return false
		}

		entryInfo.key = k
		m.expectedEntries = append(m.expectedEntries, entryInfo)
		return true
	})
}

// summary(Map): compares the contents of a map
// input(Map): map,ptr(ptr on map)

// Map operator compares the contents of a map against the non-zero
// values of model (if any) and the values of expectedEntries.
//
// model must be the same type as compared data.
//
// expectedEntries can be omitted, if no [TestDeep] operators are
// involved. If expectedEntries contains more than one item, all items
// are merged before their use, from left to right.
//
// During a match, all expected entries must be found and all data
// entries must be expected to succeed.
//
//	got := map[string]string{
//	  "foo": "test",
//	  "bar": "wizz",
//	  "zip": "buzz",
//	}
//	td.Cmp(t, got, td.Map(
//	  map[string]string{
//	    "foo": "test",
//	    "bar": "wizz",
//	  },
//	  td.MapEntries{
//	    "zip": td.HasSuffix("zz"),
//	  }),
//	) // succeeds
//
// TypeBehind method returns the [reflect.Type] of model.
//
// See also [SubMapOf] and [SuperMapOf].
func Map(model any, expectedEntries ...MapEntries) TestDeep {
	return newMap(model, allMap, expectedEntries...)
}

// summary(SubMapOf): compares the contents of a map but with
// potentially some exclusions
// input(SubMapOf): map,ptr(ptr on map)

// SubMapOf operator compares the contents of a map against the non-zero
// values of model (if any) and the values of expectedEntries.
//
// model must be the same type as compared data.
//
// expectedEntries can be omitted, if no [TestDeep] operators are
// involved. If expectedEntries contains more than one item, all items
// are merged before their use, from left to right.
//
// During a match, each map entry should be matched by an expected
// entry to succeed. But some expected entries can be missing from the
// compared map.
//
//	got := map[string]string{
//	  "foo": "test",
//	  "zip": "buzz",
//	}
//	td.Cmp(t, got, td.SubMapOf(
//	  map[string]string{
//	    "foo": "test",
//	    "bar": "wizz",
//	  },
//	  td.MapEntries{
//	    "zip": td.HasSuffix("zz"),
//	  }),
//	) // succeeds
//
//	td.Cmp(t, got, td.SubMapOf(
//	  map[string]string{
//	    "bar": "wizz",
//	  },
//	  td.MapEntries{
//	    "zip": td.HasSuffix("zz"),
//	  }),
//	) // fails, extra {"foo": "test"} in got
//
// TypeBehind method returns the [reflect.Type] of model.
//
// See also [Map] and [SuperMapOf].
func SubMapOf(model any, expectedEntries ...MapEntries) TestDeep {
	return newMap(model, subMap, expectedEntries...)
}

// summary(SuperMapOf): compares the contents of a map but with
// potentially some extra entries
// input(SuperMapOf): map,ptr(ptr on map)

// SuperMapOf operator compares the contents of a map against the non-zero
// values of model (if any) and the values of expectedEntries.
//
// model must be the same type as compared data.
//
// expectedEntries can be omitted, if no [TestDeep] operators are
// involved. If expectedEntries contains more than one item, all items
// are merged before their use, from left to right.
//
// During a match, each expected entry should match in the compared
// map. But some entries in the compared map may not be expected.
//
//	got := map[string]string{
//	  "foo": "test",
//	  "bar": "wizz",
//	  "zip": "buzz",
//	}
//	td.Cmp(t, got, td.SuperMapOf(
//	  map[string]string{
//	    "foo": "test",
//	  },
//	  td.MapEntries{
//	    "zip": td.HasSuffix("zz"),
//	  }),
//	) // succeeds
//
//	td.Cmp(t, got, td.SuperMapOf(
//	  map[string]string{
//	    "foo": "test",
//	  },
//	  td.MapEntries{
//	    "biz": td.HasSuffix("zz"),
//	  }),
//	) // fails, missing {"biz": …} in got
//
// TypeBehind method returns the [reflect.Type] of model.
//
// See also [SuperMapOf] and [SubMapOf].
func SuperMapOf(model any, expectedEntries ...MapEntries) TestDeep {
	return newMap(model, superMap, expectedEntries...)
}

func (m *tdMap) Match(ctx ctxerr.Context, got reflect.Value) (err *ctxerr.Error) {
	if m.err != nil {
		return ctx.CollectError(m.err)
	}

	err = m.checkPtr(ctx, &got, true)
	if err != nil {
		return ctx.CollectError(err)
	}

	return m.match(ctx, got)
}

func (m *tdMap) match(ctx ctxerr.Context, got reflect.Value) (err *ctxerr.Error) {
	err = m.checkType(ctx, got)
	if err != nil {
		return ctx.CollectError(err)
	}

	var notFoundKeys []reflect.Value
	foundKeys := map[any]bool{}

	for _, entryInfo := range m.expectedEntries {
		gotValue := got.MapIndex(entryInfo.key)
		if !gotValue.IsValid() {
			notFoundKeys = append(notFoundKeys, entryInfo.key)
			continue
		}

		err = deepValueEqual(ctx.AddMapKey(entryInfo.key),
			gotValue, entryInfo.expected)
		if err != nil {
			return err
		}
		foundKeys[dark.MustGetInterface(entryInfo.key)] = true
	}

	const errorMessage = "comparing hash keys of %%"

	// For SuperMapOf we don't care about extra keys
	if m.kind == superMap {
		if len(notFoundKeys) == 0 {
			return nil
		}

		if ctx.BooleanError {
			return ctxerr.BooleanError
		}
		return ctx.CollectError(&ctxerr.Error{
			Message: errorMessage,
			Summary: (tdSetResult{
				Kind:    keysSetResult,
				Missing: notFoundKeys,
				Sort:    true,
			}).Summary(),
		})
	}

	// No extra key to search, all got keys have been found
	if got.Len() == len(foundKeys) {
		if m.kind == subMap {
			return nil
		}
		// allMap

		if len(notFoundKeys) == 0 {
			return nil
		}

		if ctx.BooleanError {
			return ctxerr.BooleanError
		}
		return ctx.CollectError(&ctxerr.Error{
			Message: errorMessage,
			Summary: (tdSetResult{
				Kind:    keysSetResult,
				Missing: notFoundKeys,
				Sort:    true,
			}).Summary(),
		})
	}

	if ctx.BooleanError {
		return ctxerr.BooleanError
	}

	// Retrieve extra keys
	res := tdSetResult{
		Kind:    keysSetResult,
		Missing: notFoundKeys,
		Extra:   make([]reflect.Value, 0, got.Len()-len(foundKeys)),
		Sort:    true,
	}

	for _, k := range tdutil.MapSortedKeys(got) {
		if !foundKeys[dark.MustGetInterface(k)] {
			res.Extra = append(res.Extra, k)
		}
	}

	return ctx.CollectError(&ctxerr.Error{
		Message: errorMessage,
		Summary: res.Summary(),
	})
}

func (m *tdMap) String() string {
	if m.err != nil {
		return m.stringError()
	}

	var buf strings.Builder

	if m.kind != allMap {
		buf.WriteString(m.GetLocation().Func)
		buf.WriteByte('(')
	}

	buf.WriteString(m.expectedTypeStr())

	if len(m.expectedEntries) == 0 {
		buf.WriteString("{}")
	} else {
		buf.WriteString("{\n")

		for _, entryInfo := range m.expectedEntries {
			fmt.Fprintf(&buf, "  %s: %s,\n", //nolint: errcheck
				util.ToString(entryInfo.key),
				util.ToString(entryInfo.expected))
		}

		buf.WriteByte('}')
	}

	if m.kind != allMap {
		buf.WriteByte(')')
	}

	return buf.String()
}
