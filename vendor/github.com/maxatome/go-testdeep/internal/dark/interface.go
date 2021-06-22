// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package dark

import (
	"reflect"
)

// GetInterface does its best to return the data behind val. If force
// is true, it tries to bypass golang protections using the unsafe
// package.
//
// It returns (nil, false) if the data behind val can not be retrieved
// as an interface{} (aka struct private + non-copyable field).
var GetInterface = func(val reflect.Value, force bool) (interface{}, bool) {
	if !val.IsValid() {
		return nil, true
	}

	if val.CanInterface() {
		return val.Interface(), true
	}

	if force {
		val = unsafeReflectValue(val)
		if val.CanInterface() {
			return val.Interface(), true
		}
	}

	// For some types, we can copy them in new visitable reflect.Value instances
	copyVal, ok := CopyValue(val)
	if ok && copyVal.CanInterface() {
		return copyVal.Interface(), true
	}

	// For others, in environments where "unsafe" package is not
	// available, we cannot go further
	return nil, false
}

// MustGetInterface does its best to return the data behind val. If it
// fails (struct private + non-copyable field), it panics.
func MustGetInterface(val reflect.Value) interface{} {
	ret, ok := GetInterface(val, true)
	if ok {
		return ret
	}
	panic("dark.GetInterface() does not handle private " +
		val.Kind().String() + " kind")
}
