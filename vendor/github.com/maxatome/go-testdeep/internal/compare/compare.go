// Copyright (c) 2019-2025, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package compare

import (
	"math"
	"reflect"
	"sync"

	"github.com/maxatome/go-testdeep/internal/visited"
)

var intType = reflect.TypeOf(0) // avoid an import cycle if using internal/types

func cmpRet(less, gt bool) int {
	if less {
		return -1
	}
	if gt {
		return 1
	}
	return 0
}

func cmpFloat(a, b float64) int {
	if math.IsNaN(a) {
		return -1
	}
	if math.IsNaN(b) {
		return 1
	}
	return cmpRet(a < b, a > b)
}

var methodCache sync.Map

func methodCompare(typ reflect.Type) reflect.Value {
	if m, ok := methodCache.Load(typ); ok {
		return m.(reflect.Value)
	}
	m, ok := typ.MethodByName("Compare")
	if !ok ||
		m.Type.IsVariadic() ||
		m.Type.NumIn() != 2 ||
		m.Type.In(1) != typ ||
		m.Type.NumOut() != 1 ||
		m.Type.Out(0) != intType {
		m.Func = reflect.Value{}
	}
	methodCache.Store(typ, m.Func)
	return m.Func
}

// Compare returns -1 if a < b, 1 if a > b, 0 if a == b.
func Compare(v visited.Visited, a, b reflect.Value) int {
	if !a.IsValid() {
		if !b.IsValid() {
			return 0
		}
		return -1
	}
	if !b.IsValid() {
		return 1
	}

	at, bt := a.Type(), b.Type()
	if at != bt {
		sat, sbt := at.String(), bt.String()
		return cmpRet(sat < sbt, sat > sbt)
	}

	// Avoid looping forever on cyclic references
	if v.Record(a, b) {
		return 0
	}

	if a.Kind() != reflect.Interface {
		if cmpFn := methodCompare(at); cmpFn.IsValid() {
			ok, cmp := false, 0
			func() {
				defer recover() //nolint: errcheck
				cmp = int(cmpFn.Call([]reflect.Value{a, b})[0].Int())
				ok = true
			}()
			if ok {
				return cmp
			}
			// When a panic occurs, fallback on generic comparison
		}
	}

	switch a.Kind() {
	case reflect.Bool:
		if a.Bool() {
			if b.Bool() {
				return 0
			}
			return 1
		}
		if b.Bool() {
			return -1
		}
		return 0

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		na, nb := a.Int(), b.Int()
		return cmpRet(na < nb, na > nb)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Uint64, reflect.Uintptr:
		na, nb := a.Uint(), b.Uint()
		return cmpRet(na < nb, na > nb)

	case reflect.Float32, reflect.Float64:
		return cmpFloat(a.Float(), b.Float())

	case reflect.Complex64, reflect.Complex128:
		na, nb := a.Complex(), b.Complex()
		fa, fb := real(na), real(nb)
		if r := cmpFloat(fa, fb); r != 0 {
			return r
		}
		return cmpFloat(imag(na), imag(nb))

	case reflect.String:
		sa, sb := a.String(), b.String()
		return cmpRet(sa < sb, sa > sb)

	case reflect.Array:
		for i := 0; i < a.Len(); i++ {
			if r := Compare(v, a.Index(i), b.Index(i)); r != 0 {
				return r
			}
		}
		return 0

	case reflect.Slice:
		al, bl := a.Len(), b.Len()
		maxl := al
		if al > bl {
			maxl = bl
		}
		for i := 0; i < maxl; i++ {
			if r := Compare(v, a.Index(i), b.Index(i)); r != 0 {
				return r
			}
		}
		return cmpRet(al < bl, al > bl)

	case reflect.Interface:
		if a.IsNil() {
			if b.IsNil() {
				return 0
			}
			return -1
		}
		if b.IsNil() {
			return 1
		}
		return Compare(v, a.Elem(), b.Elem())

	case reflect.Struct:
		for i, m := 0, a.NumField(); i < m; i++ {
			if r := Compare(v, a.Field(i), b.Field(i)); r != 0 {
				return r
			}
		}
		return 0

	case reflect.Ptr:
		if a.Pointer() == b.Pointer() {
			return 0
		}
		if a.IsNil() {
			return -1
		}
		if b.IsNil() {
			return 1
		}
		return Compare(v, a.Elem(), b.Elem())

	case reflect.Map:
		// consider shorter maps are before longer ones
		al, bl := a.Len(), b.Len()
		if r := cmpRet(al < bl, al > bl); r != 0 {
			return r
		}
		// then fallback on pointers comparison. How to say a map is
		// before another one otherwise?
		fallthrough

	case reflect.Func, reflect.Chan, reflect.UnsafePointer:
		pa, pb := a.Pointer(), b.Pointer()
		return cmpRet(pa < pb, pa > pb)

	default:
		panic("don't know how to compare " + a.Kind().String())
	}
}
