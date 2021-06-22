// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package dark

import (
	"reflect"
	"unicode"
	"unicode/utf8"

	"github.com/maxatome/go-testdeep/helpers/tdutil"
)

// CopyValue does its best to copy val in a new reflect.Value instance.
func CopyValue(val reflect.Value) (reflect.Value, bool) {
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			newPtrVal := reflect.New(val.Type())
			return newPtrVal.Elem(), true
		}

		refVal, ok := CopyValue(val.Elem())
		if !ok {
			return reflect.Value{}, false
		}
		newPtrVal := reflect.New(refVal.Type())
		newPtrVal.Elem().Set(refVal)
		return newPtrVal, true
	}

	var newVal reflect.Value

	switch val.Kind() {
	case reflect.Bool:
		newPtrVal := reflect.New(val.Type())
		newVal = newPtrVal.Elem()
		newVal.SetBool(val.Bool())

	case reflect.Complex64, reflect.Complex128:
		newPtrVal := reflect.New(val.Type())
		newVal = newPtrVal.Elem()
		newVal.SetComplex(val.Complex())

	case reflect.Float32, reflect.Float64:
		newPtrVal := reflect.New(val.Type())
		newVal = newPtrVal.Elem()
		newVal.SetFloat(val.Float())

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		newPtrVal := reflect.New(val.Type())
		newVal = newPtrVal.Elem()
		newVal.SetInt(val.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		newPtrVal := reflect.New(val.Type())
		newVal = newPtrVal.Elem()
		newVal.SetUint(val.Uint())

	case reflect.Array:
		newPtrVal := reflect.New(val.Type())
		newVal = newPtrVal.Elem()
		var (
			item reflect.Value
			ok   bool
		)
		for i := val.Len() - 1; i >= 0; i-- {
			item, ok = CopyValue(val.Index(i))
			if !ok {
				return reflect.Value{}, false
			}
			newVal.Index(i).Set(item)
		}

	case reflect.Slice:
		if val.IsNil() {
			newPtrVal := reflect.New(val.Type())
			return newPtrVal.Elem(), true
		}

		newVal = reflect.MakeSlice(val.Type(), val.Len(), val.Cap())
		var (
			item reflect.Value
			ok   bool
		)
		for i := val.Len() - 1; i >= 0; i-- {
			item, ok = CopyValue(val.Index(i))
			if !ok {
				return reflect.Value{}, false
			}
			newVal.Index(i).Set(item)
		}

	case reflect.Map:
		if val.IsNil() {
			newPtrVal := reflect.New(val.Type())
			return newPtrVal.Elem(), true
		}

		newVal = reflect.MakeMapWithSize(val.Type(), val.Len())
		var (
			key, value reflect.Value
			ok         bool
		)
		if !tdutil.MapEach(val, func(k, v reflect.Value) bool {
			key, ok = CopyValue(k)
			if !ok {
				return false
			}
			value, ok = CopyValue(v)
			if !ok {
				return false
			}
			newVal.SetMapIndex(key, value)
			return true
		}) {
			return reflect.Value{}, false
		}

	case reflect.Interface:
		if val.IsNil() {
			newPtrVal := reflect.New(val.Type())
			return newPtrVal.Elem(), true
		}

		newPtrVal := reflect.New(val.Type())
		newVal = newPtrVal.Elem()
		refVal, ok := CopyValue(val.Elem())
		if !ok {
			return reflect.Value{}, false
		}
		newVal.Set(refVal)

	case reflect.String:
		newPtrVal := reflect.New(val.Type())
		newVal = newPtrVal.Elem()
		newVal.SetString(val.String())

	case reflect.Struct:
		// First, check if all fields are public
		sType := val.Type()
		for i, n := 0, val.NumField(); i < n; i++ {
			r, _ := utf8.DecodeRuneInString(sType.Field(i).Name)
			if !unicode.IsUpper(r) {
				return reflect.Value{}, false
			}
		}

		// OK all fields are public
		newPtrVal := reflect.New(sType)
		newVal = newPtrVal.Elem()

		var (
			fieldIdx []int
			fieldVal reflect.Value
			ok       bool
		)
		for i, n := 0, val.NumField(); i < n; i++ {
			fieldIdx = sType.Field(i).Index

			fieldVal, ok = CopyValue(val.FieldByIndex(fieldIdx))
			if !ok {
				return reflect.Value{}, false // Should not happen as already checked
			}
			newVal.FieldByIndex(fieldIdx).Set(fieldVal)
		}

		// Does not handle Chan, Func and UnsafePointer
	default:
		return reflect.Value{}, false
	}
	return newVal, true
}
