// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/maxatome/go-testdeep/helpers/tdutil"
	"github.com/maxatome/go-testdeep/internal/ctxerr"
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

// MapEntries allows to pass map entries to check in function Map. It
// is a map whose each key is the expected entry key and the
// corresponding value the expected entry value (which can be a
// TestDeep operator as well as a zero value.)
type MapEntries map[interface{}]interface{}

func newMap(model interface{}, entries MapEntries, kind mapKind) *tdMap {
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
			m.populateExpectedEntries(entries, reflect.Value{})
			return &m
		}

		vmodel = vmodel.Elem()
		fallthrough

	case reflect.Map:
		m.expectedType = vmodel.Type()
		m.populateExpectedEntries(entries, vmodel)
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
	checkedEntries := make(map[interface{}]bool, len(entries))

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
				m.err = ctxerr.OpBad(
					m.GetLocation().Func,
					"expected key %s value cannot be nil as entries value type is %s",
					util.ToString(key),
					valueType)
				return
			}
		} else {
			entryInfo.expected = reflect.ValueOf(expectedValue)

			if _, ok := expectedValue.(TestDeep); !ok {
				if !entryInfo.expected.Type().AssignableTo(valueType) {
					m.err = ctxerr.OpBad(
						m.GetLocation().Func,
						"expected key %s value type mismatch: %s != model key type (%s)",
						util.ToString(key),
						entryInfo.expected.Type(),
						valueType)
					return
				}
			}
		}

		entryInfo.key = vkey
		m.expectedEntries = append(m.expectedEntries, entryInfo)
		checkedEntries[vkey.Interface()] = true
	}

	// Check entries in model
	if keysInModel == 0 {
		return
	}

	tdutil.MapEach(expectedModel, func(k, v reflect.Value) bool {
		entryInfo.expected = v

		if checkedEntries[k.Interface()] {
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
// values of "model" (if any) and the values of "expectedEntries".
//
// "model" must be the same type as compared data.
//
// "expectedEntries" can be nil, if no zero entries are expected and
// no TestDeep operator are involved.
//
// During a match, all expected entries must be found and all data
// entries must be expected to succeed.
//
//   got := map[string]string{
//     "foo": "test",
//     "bar": "wizz",
//     "zip": "buzz",
//   }
//   td.Cmp(t, got, td.Map(
//     map[string]string{
//       "foo": "test",
//       "bar": "wizz",
//     },
//     td.MapEntries{
//       "zip": td.HasSuffix("zz"),
//     }),
//   ) // succeeds
//
// TypeBehind method returns the reflect.Type of "model".
func Map(model interface{}, expectedEntries MapEntries) TestDeep {
	return newMap(model, expectedEntries, allMap)
}

// summary(SubMapOf): compares the contents of a map but with
// potentially some exclusions
// input(SubMapOf): map,ptr(ptr on map)

// SubMapOf operator compares the contents of a map against the non-zero
// values of "model" (if any) and the values of "expectedEntries".
//
// "model" must be the same type as compared data.
//
// "expectedEntries" can be nil, if no zero entries are expected and
// no TestDeep operator are involved.
//
// During a match, each map entry should be matched by an expected
// entry to succeed. But some expected entries can be missing from the
// compared map.
//
//   got := map[string]string{
//     "foo": "test",
//     "zip": "buzz",
//   }
//   td.Cmp(t, got, td.SubMapOf(
//     map[string]string{
//       "foo": "test",
//       "bar": "wizz",
//     },
//     td.MapEntries{
//       "zip": td.HasSuffix("zz"),
//     }),
//   ) // succeeds
//
//   td.Cmp(t, got, td.SubMapOf(
//     map[string]string{
//       "bar": "wizz",
//     },
//     td.MapEntries{
//       "zip": td.HasSuffix("zz"),
//     }),
//   ) // fails, extra {"foo": "test"} in got
//
// TypeBehind method returns the reflect.Type of "model".
func SubMapOf(model interface{}, expectedEntries MapEntries) TestDeep {
	return newMap(model, expectedEntries, subMap)
}

// summary(SuperMapOf): compares the contents of a map but with
// potentially some extra entries
// input(SuperMapOf): map,ptr(ptr on map)

// SuperMapOf operator compares the contents of a map against the non-zero
// values of "model" (if any) and the values of "expectedEntries".
//
// "model" must be the same type as compared data.
//
// "expectedEntries" can be nil, if no zero entries are expected and
// no TestDeep operator are involved.
//
// During a match, each expected entry should match in the compared
// map. But some entries in the compared map may not be expected.
//
//   got := map[string]string{
//     "foo": "test",
//     "bar": "wizz",
//     "zip": "buzz",
//   }
//   td.Cmp(t, got, td.SuperMapOf(
//     map[string]string{
//       "foo": "test",
//     },
//     td.MapEntries{
//       "zip": td.HasSuffix("zz"),
//     }),
//   ) // succeeds
//
//   td.Cmp(t, got, td.SuperMapOf(
//     map[string]string{
//       "foo": "test",
//     },
//     td.MapEntries{
//       "biz": td.HasSuffix("zz"),
//     }),
//   ) // fails, missing {"biz": …} in got
//
// TypeBehind method returns the reflect.Type of "model".
func SuperMapOf(model interface{}, expectedEntries MapEntries) TestDeep {
	return newMap(model, expectedEntries, superMap)
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
	foundKeys := map[interface{}]bool{}

	for _, entryInfo := range m.expectedEntries {
		gotValue := got.MapIndex(entryInfo.key)
		if !gotValue.IsValid() {
			notFoundKeys = append(notFoundKeys, entryInfo.key)
			continue
		}

		err = deepValueEqual(ctx.AddMapKey(entryInfo.key),
			got.MapIndex(entryInfo.key), entryInfo.expected)
		if err != nil {
			return err
		}
		foundKeys[entryInfo.key.Interface()] = true
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
		if !foundKeys[k.Interface()] {
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

	buf := &bytes.Buffer{}

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
			fmt.Fprintf(buf, "  %s: %s,\n", // nolint: errcheck
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
