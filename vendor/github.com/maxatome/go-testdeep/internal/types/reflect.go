// Copyright (c) 2020, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package types

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

var (
	Bool            = reflect.TypeOf(false)
	Interface       = reflect.TypeOf((*interface{})(nil)).Elem()
	SliceInterface  = reflect.TypeOf(([]interface{})(nil))
	FmtStringer     = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	Error           = reflect.TypeOf((*error)(nil)).Elem()
	JsonUnmarshaler = reflect.TypeOf((*json.Unmarshaler)(nil)).Elem() //nolint: revive
	Time            = reflect.TypeOf(time.Time{})
	Int             = reflect.TypeOf(int(0))
	Uint8           = reflect.TypeOf(uint8(0))
	Rune            = reflect.TypeOf(rune(0))
	String          = reflect.TypeOf("")
)

// IsStruct returns true if "t" is a struct or a pointer on a struct
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
<<<<<<< HEAD
||||||| parent of 6b7ce455e (update vendored files)
=======

// IsTypeOrConvertible returns (true, false) if "v" type == "target",
// (true, true) if "v" if convertible to "target" type, (false, false)
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

// IsConvertible returns true if "v" if convertible to "target" type,
// false otherwise.
//
// It handles go 1.17 slice to array pointer convertibility.
func IsConvertible(v reflect.Value, target reflect.Type) bool {
	if v.Type().ConvertibleTo(target) {
		// Since go 1.17, a slice can be convertible to a pointer of an
		// array, but Convert() may still panic if the slice length is lesser
		// than array pointed one
		if v.Kind() != reflect.Slice ||
			target.Kind() != reflect.Ptr ||
			target.Elem().Kind() != reflect.Array ||
			v.Len() >= target.Elem().Len() {
			return true
		}
	}
	return false
}
>>>>>>> 6b7ce455e (update vendored files)
