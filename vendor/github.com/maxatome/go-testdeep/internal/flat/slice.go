// Copyright (c) 2020, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package flat

import (
	"reflect"

	"github.com/maxatome/go-testdeep/helpers/tdutil"
)

var sliceType = reflect.TypeOf(Slice{})

// Slice allows to flatten any slice, array or map.
type Slice struct {
	Slice interface{}
}

// isFlat returns true if no flat.Slice items can be contained in
// f.Slice, so this Slice is already flattened.
func (f Slice) isFlat() bool {
	t := reflect.TypeOf(f.Slice).Elem()
	return t != sliceType && t.Kind() != reflect.Interface
}

func subLen(v reflect.Value) int {
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	if s, ok := v.Interface().(Slice); ok {
		return s.len()
	}
	return 1
}

func (f Slice) len() int {
	fv := reflect.ValueOf(f.Slice)
	l := 0

	if fv.Kind() == reflect.Map {
		if f.isFlat() {
			return fv.Len() * 2
		}
		tdutil.MapEach(fv, func(k, v reflect.Value) bool {
			l += 1 + subLen(v)
			return true
		})
		return l
	}

	fvLen := fv.Len()
	if f.isFlat() {
		return fvLen
	}
	for i := 0; i < fvLen; i++ {
		l += subLen(fv.Index(i))
	}
	return l
}

func subAppendValuesTo(sv []reflect.Value, v reflect.Value) []reflect.Value {
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	if s, ok := v.Interface().(Slice); ok {
		return s.appendValuesTo(sv)
	}
	return append(sv, v)
}

func (f Slice) appendValuesTo(sv []reflect.Value) []reflect.Value {
	fv := reflect.ValueOf(f.Slice)

	if fv.Kind() == reflect.Map {
		if f.isFlat() {
			tdutil.MapEach(fv, func(k, v reflect.Value) bool {
				sv = append(sv, k, v)
				return true
			})
			return sv
		}

		tdutil.MapEach(fv, func(k, v reflect.Value) bool {
			sv = append(sv, k)
			sv = subAppendValuesTo(sv, v)
			return true
		})
		return sv
	}

	fvLen := fv.Len()

	if f.isFlat() {
		for i := 0; i < fvLen; i++ {
			sv = append(sv, fv.Index(i))
		}
		return sv
	}

	for i := 0; i < fvLen; i++ {
		sv = subAppendValuesTo(sv, fv.Index(i))
	}
	return sv
}

func subAppendTo(si []interface{}, v reflect.Value) []interface{} {
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	i := v.Interface()
	if s, ok := i.(Slice); ok {
		return s.appendTo(si)
	}
	return append(si, i)
}

func (f Slice) appendTo(si []interface{}) []interface{} {
	fv := reflect.ValueOf(f.Slice)

	if fv.Kind() == reflect.Map {
		if f.isFlat() {
			tdutil.MapEach(fv, func(k, v reflect.Value) bool {
				si = append(si, k.Interface(), v.Interface())
				return true
			})
			return si
		}

		tdutil.MapEach(fv, func(k, v reflect.Value) bool {
			si = append(si, k.Interface())
			si = subAppendTo(si, v)
			return true
		})
		return si
	}

	fvLen := fv.Len()

	if f.isFlat() {
		for i := 0; i < fvLen; i++ {
			si = append(si, fv.Index(i).Interface())
		}
		return si
	}

	for i := 0; i < fvLen; i++ {
		si = subAppendTo(si, fv.Index(i))
	}
	return si
}

// Len returns the number of items contained in items. Nested Slice
// items are counted as if they are flattened. It returns true if at
// least one Slice item is found, false otherwise.
func Len(items []interface{}) (int, bool) {
	l := len(items)
	flattened := true

	for _, item := range items {
		if subf, ok := item.(Slice); ok {
			l += subf.len() - 1
			flattened = false
		}
	}
	return l, flattened
}

// Values returns the items values as a slice of reflect.Value. Nested
// Slice items are flattened.
func Values(items []interface{}) []reflect.Value {
	l, flattened := Len(items)
	if flattened {
		sv := make([]reflect.Value, l)
		for i, item := range items {
			sv[i] = reflect.ValueOf(item)
		}
		return sv
	}

	sv := make([]reflect.Value, 0, l)
	for _, item := range items {
		if f, ok := item.(Slice); ok {
			sv = f.appendValuesTo(sv)
		} else {
			sv = append(sv, reflect.ValueOf(item))
		}
	}
	return sv
}

// Interfaces returns the items values as a slice of
// interface{}. Nested Slice items are flattened.
func Interfaces(items ...interface{}) []interface{} {
	l, flattened := Len(items)
	if flattened {
		return items
	}

	si := make([]interface{}, 0, l)
	for _, item := range items {
		if f, ok := item.(Slice); ok {
			si = f.appendTo(si)
		} else {
			si = append(si, item)
		}
	}
	return si
}
