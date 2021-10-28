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
