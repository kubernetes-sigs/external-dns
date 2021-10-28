// Copyright (c) 2021-2022, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package types

import (
	"reflect"

	"github.com/maxatome/go-testdeep/internal/dark"
)

// NewOrder returns a function able to compare 2 non-nil values of type "t".
// It returns nil if the type "t" is not comparable.
func NewOrder(t reflect.Type) func(a, b reflect.Value) int {
	// Compare(T) int
	if m, ok := cmpMethod("Compare", t, Int); ok {
		return func(va, vb reflect.Value) int {
			// use dark.MustGetInterface() to bypass possible private fields
			ret := m.Call([]reflect.Value{
				reflect.ValueOf(dark.MustGetInterface(va)),
				reflect.ValueOf(dark.MustGetInterface(vb)),
			})
			return int(ret[0].Int())
		}
	}

	// Less(T) bool
	if m, ok := cmpMethod("Less", t, Bool); ok {
		return func(va, vb reflect.Value) int {
			// use dark.MustGetInterface() to bypass possible private fields
			va = reflect.ValueOf(dark.MustGetInterface(va))
			vb = reflect.ValueOf(dark.MustGetInterface(vb))
			ret := m.Call([]reflect.Value{va, vb})
			if ret[0].Bool() { // a < b
				return -1
			}
			ret = m.Call([]reflect.Value{vb, va})
			if ret[0].Bool() { // b < a
				return 1
			}
			return 0
		}
	}

	return nil
}

func cmpMethod(name string, in, out reflect.Type) (reflect.Value, bool) {
	if equal, ok := in.MethodByName(name); ok {
		ft := equal.Type
		if !ft.IsVariadic() &&
			ft.NumIn() == 2 &&
			ft.NumOut() == 1 &&
			ft.In(0) == in &&
			ft.In(1) == in &&
			ft.Out(0) == out {
			return equal.Func, true
		}
	}
	return reflect.Value{}, false
}
