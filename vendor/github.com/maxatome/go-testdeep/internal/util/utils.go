// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package util

import (
	"fmt"
	"reflect"
	"strings"
)

// BadParam returns a string noticing a misuse of a function parameter.
//
// If kind and param's kind name ≠ param's type name:
//
//	but received {param type} ({param kind}) as {pos}th parameter
//
// else
//
//	but received {param type} as {pos}th parameter
func BadParam(param any, pos int, kind bool) string {
	var b strings.Builder
	b.WriteString("but received ")

	if param == nil {
		b.WriteString("nil")
	} else {
		t := reflect.TypeOf(param)
		if kind && t.String() != t.Kind().String() {
			fmt.Fprintf(&b, "%s (%s)", t, t.Kind())
		} else {
			b.WriteString(t.String())
		}
	}

	b.WriteString(" as ")
	switch pos {
	case 1:
		b.WriteString("1st")
	case 2:
		b.WriteString("2nd")
	case 3:
		b.WriteString("3rd")
	default:
		fmt.Fprintf(&b, "%dth", pos)
	}
	b.WriteString(" parameter")
	return b.String()
}

// TernRune returns a if cond is true, b otherwise.
func TernRune(cond bool, a, b rune) rune {
	if cond {
		return a
	}
	return b
}
