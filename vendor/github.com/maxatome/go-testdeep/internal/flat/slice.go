// Copyright (c) 2020-2025, Maxime Soulé
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

// Slice allows to flatten any slice, array, map or int range.
type Slice struct {
	Slice any
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

	switch fv.Kind() {
	case reflect.Map:
		if f.isFlat() {
			return fv.Len() * 2
		}
		tdutil.MapEach(fv, func(k, v reflect.Value) bool {
			l += 1 + subLen(v)
			return true
		})

	case reflect.Array, reflect.Slice:
		fvLen := fv.Len()
		if f.isFlat() {
			return fvLen
		}
		for i := 0; i < fvLen; i++ {
			l += subLen(fv.Index(i))
		}

	default: // reflect.Int
		l = int(fv.Int())
		if l <= 0 {
			return 0
		}
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

	switch fv.Kind() {
	case reflect.Map:
		if f.isFlat() {
			tdutil.MapEach(fv, func(k, v reflect.Value) bool {
				sv = append(sv, k, v)
				return true
			})
			break
		}
		tdutil.MapEach(fv, func(k, v reflect.Value) bool {
			sv = append(sv, k)
			sv = subAppendValuesTo(sv, v)
			return true
		})

	case reflect.Array, reflect.Slice:
		fvLen := fv.Len()
		if f.isFlat() {
			for i := 0; i < fvLen; i++ {
				sv = append(sv, fv.Index(i))
			}
			break
		}
		for i := 0; i < fvLen; i++ {
			sv = subAppendValuesTo(sv, fv.Index(i))
		}

	default: // reflect.Int
		if l := int(fv.Int()); l > 0 {
			for i := 0; i < l; i++ {
				sv = append(sv, reflect.ValueOf(i))
			}
		}
	}

	return sv
}

func subAppendTo(si []any, v reflect.Value) []any {
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	i := v.Interface()
	if s, ok := i.(Slice); ok {
		return s.appendTo(si)
	}
	return append(si, i)
}

func (f Slice) appendTo(si []any) []any {
	fv := reflect.ValueOf(f.Slice)

	switch fv.Kind() {
	case reflect.Map:
		if f.isFlat() {
			tdutil.MapEach(fv, func(k, v reflect.Value) bool {
				si = append(si, k.Interface(), v.Interface())
				return true
			})
			break
		}
		tdutil.MapEach(fv, func(k, v reflect.Value) bool {
			si = append(si, k.Interface())
			si = subAppendTo(si, v)
			return true
		})

	case reflect.Array, reflect.Slice:
		fvLen := fv.Len()
		if f.isFlat() {
			for i := 0; i < fvLen; i++ {
				si = append(si, fv.Index(i).Interface())
			}
			break
		}
		for i := 0; i < fvLen; i++ {
			si = subAppendTo(si, fv.Index(i))
		}

	default: // reflect.Int
		if l := int(fv.Int()); l > 0 {
			for i := 0; i < l; i++ {
				si = append(si, i)
			}
		}
	}

	return si
}

// Len returns the number of items contained in items. Nested Slice
// items are counted as if they are flattened. It returns false if at
// least one [Slice] item is found, true otherwise meaning the slice
// is already flattened.
func Len(items []any) (int, bool) {
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

// Values returns the items values as a slice of
// [reflect.Value]. Nested [Slice] items are flattened.
func Values(items []any) []reflect.Value {
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
// any. Nested [Slice] items are flattened.
func Interfaces(items ...any) []any {
	l, flattened := Len(items)
	if flattened {
		return items
	}

	si := make([]any, 0, l)
	for _, item := range items {
		if f, ok := item.(Slice); ok {
			si = f.appendTo(si)
		} else {
			si = append(si, item)
		}
	}
	return si
}
