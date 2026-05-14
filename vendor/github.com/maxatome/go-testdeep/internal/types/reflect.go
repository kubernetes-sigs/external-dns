// Copyright (c) 2020-2022, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package types

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"
)

var (
	Bool            = reflect.TypeOf(false)
	Interface       = reflect.TypeOf((*any)(nil)).Elem()
	SliceInterface  = reflect.TypeOf(([]any)(nil))
	FmtStringer     = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	Error           = reflect.TypeOf((*error)(nil)).Elem()
	JsonUnmarshaler = reflect.TypeOf((*json.Unmarshaler)(nil)).Elem() //nolint: revive
	Time            = reflect.TypeOf(time.Time{})
	Int             = reflect.TypeOf(int(0))
	Uint8           = reflect.TypeOf(uint8(0))
	Rune            = reflect.TypeOf(rune(0))
	String          = reflect.TypeOf("")
)

// IsStruct returns true if t is a struct or a pointer on a struct
// (whatever the number of chained pointers), false otherwise.
func IsStruct(t reflect.Type) bool {
	for {
		switch t.Kind() {
		case reflect.Struct:
			return true
		case reflect.Ptr:
			t = t.Elem()
		default:
			return false
		}
	}
}

// IsTypeOrConvertible returns (true, false) if v type == target,
// (true, true) if v if convertible to target type, (false, false)
// otherwise.
//
// It handles go 1.17 slice to array pointer convertibility.
func IsTypeOrConvertible(v reflect.Value, target reflect.Type) (bool, bool) {
	if v.Type() == target {
		return true, false
	}
	if IsConvertible(v, target) {
		return true, true
	}
	return false, false
}

// IsConvertible returns true if v is convertible to target type,
// false otherwise.
//
// It handles go 1.17 slice to array pointer convertibility.
// It handles go 1.20 slice to array convertibility.
func IsConvertible(v reflect.Value, target reflect.Type) bool {
	if v.Type().ConvertibleTo(target) {
		tk := target.Kind()
		if v.Kind() != reflect.Slice ||
			(tk != reflect.Ptr && tk != reflect.Array) ||
			// Since go 1.17, a slice can be convertible to a pointer to an
			// array, but Convert() may still panic if the slice length is lesser
			// than array pointed one
			(tk == reflect.Ptr && (target.Elem().Kind() != reflect.Array ||
				v.Len() >= target.Elem().Len())) ||
			// Since go 1.20, a slice can also be convertible to an array, but
			// Convert() may still panic if the slice length is lesser than
			// array one
			(tk == reflect.Array && v.Len() >= target.Len()) {
			return true
		}
	}
	return false
}

// KindType returns the kind of val as a string. If the kind is
// [reflect.Ptr], a "*" is used as prefix of kind of
// val.Type().Elem(), and so on. If the final kind differs from
// val.Type(), the type is appended inside parenthesis.
func KindType(val reflect.Value) string {
	if !val.IsValid() {
		return "nil"
	}

	nptr := 0
	typ := val.Type()
	for typ.Kind() == reflect.Ptr {
		nptr++
		typ = typ.Elem()
	}
	kind := strings.Repeat("*", nptr) + typ.Kind().String()
	if typ := val.Type().String(); kind != typ {
		kind += " (" + typ + " type)"
	}
	return kind
}
