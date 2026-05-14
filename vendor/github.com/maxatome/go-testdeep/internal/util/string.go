// Copyright (c) 2018-2022, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package util

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"

	"github.com/maxatome/go-testdeep/helpers/tdutil"
	"github.com/maxatome/go-testdeep/internal/dark"
	"github.com/maxatome/go-testdeep/internal/types"
)

// ToString does its best to stringify val. inReflectValue is used
// internally to avoid treating specifically reflect.Value type.
func ToString(val any, inReflectValue ...bool) string {
	if val == nil {
		return "nil"
	}

	switch tval := val.(type) {
	case reflect.Value:
		if len(inReflectValue) > 0 && inReflectValue[0] {
			break
		}
		newVal, ok := dark.GetInterface(tval, true)
		if ok {
			return ToString(newVal, true)
		}

	case []reflect.Value:
		if len(inReflectValue) > 0 && inReflectValue[0] {
			break
		}
		var buf strings.Builder
		SliceToString(&buf, tval)
		return buf.String()

		// no "(string) " prefix for printable strings
	case string:
		return tdutil.FormatString(tval)

		// no "(int) " prefix for ints
	case int:
		return strconv.Itoa(tval)

		// no "(float64) " prefix for float64s
	case float64:
		s := strconv.FormatFloat(tval, 'g', -1, 64)
		if strings.ContainsAny(s, "e.IN") { // I for Inf, N for NaN
			return s
		}
		return s + ".0" // to distinguish from ints

		// no "(bool) " prefix for booleans
	case bool:
		if tval {
			return "true"
		}
		return "false"

	case types.TestDeepStringer:
		return tval.String()
	}

	return tdutil.SpewString(val)
}

// IndentString indents str lines (from 2nd one = 1st line is not
// indented) by indent.
func IndentString(str, indent string) string {
	return strings.ReplaceAll(str, "\n", "\n"+indent)
}

// IndentStringIn indents str lines (from 2nd one = 1st line is not
// indented) by indent and write it to w.
func IndentStringIn(w io.Writer, str, indent string) {
	repl := strings.NewReplacer("\n", "\n"+indent)
	repl.WriteString(w, str) //nolint: errcheck
}

// IndentColorizeStringIn indents str lines (from 2nd one = 1st line
// is not indented) by indent and write it to w. Before each end of
// line, colOff is inserted, and after each indent on new line, colOn
// is inserted.
func IndentColorizeStringIn(w io.Writer, str, indent, colOn, colOff string) {
	if str != "" {
		if colOn == "" && colOff == "" {
			IndentStringIn(w, str, indent)
			return
		}
		repl := strings.NewReplacer("\n", colOff+"\n"+indent+colOn)
		io.WriteString(w, colOn)  //nolint: errcheck
		repl.WriteString(w, str)  //nolint: errcheck
		io.WriteString(w, colOff) //nolint: errcheck
	}
}

// SliceToString stringifies items slice into buf then returns buf.
func SliceToString(buf *strings.Builder, items []reflect.Value) *strings.Builder {
	buf.WriteByte('(')

	begLine := strings.LastIndexByte(buf.String(), '\n') + 1
	prefix := strings.Repeat(" ", buf.Len()-begLine)

	if len(items) < 2 {
		if len(items) > 0 {
			buf.WriteString(IndentString(ToString(items[0]), prefix))
		}
	} else {
		buf.WriteString(IndentString(ToString(items[0]), prefix))
		for _, item := range items[1:] {
			buf.WriteString(",\n")
			buf.WriteString(prefix)
			buf.WriteString(IndentString(ToString(item), prefix))
		}
	}
	buf.WriteByte(')')

	return buf
}

// TypeFullName returns the t type name with packages fully visible
// instead of the last package part in t.String().
func TypeFullName(t reflect.Type) string {
	var b bytes.Buffer
	typeFullName(&b, t)
	return b.String()
}

func typeFullName(b *bytes.Buffer, t reflect.Type) {
	if t.Name() != "" {
		if pkg := t.PkgPath(); pkg != "" {
			fmt.Fprintf(b, "%s.", pkg)
		}
		b.WriteString(t.Name())
		return
	}

	switch t.Kind() {
	case reflect.Ptr:
		b.WriteByte('*')
		typeFullName(b, t.Elem())

	case reflect.Slice:
		b.WriteString("[]")
		typeFullName(b, t.Elem())

	case reflect.Array:
		fmt.Fprintf(b, "[%d]", t.Len())
		typeFullName(b, t.Elem())

	case reflect.Map:
		b.WriteString("map[")
		typeFullName(b, t.Key())
		b.WriteByte(']')
		typeFullName(b, t.Elem())

	case reflect.Struct:
		b.WriteString("struct {")
		if num := t.NumField(); num > 0 {
			for i := 0; i < num; i++ {
				sf := t.Field(i)
				if !sf.Anonymous {
					b.WriteByte(' ')
					b.WriteString(sf.Name)
				}
				b.WriteByte(' ')
				typeFullName(b, sf.Type)
				b.WriteByte(';')
			}
			b.Truncate(b.Len() - 1)
			b.WriteByte(' ')
		}
		b.WriteByte('}')

	case reflect.Func:
		b.WriteString("func(")
		if num := t.NumIn(); num > 0 {
			for i := 0; i < num; i++ {
				if i == num-1 && t.IsVariadic() {
					b.WriteString("...")
					typeFullName(b, t.In(i).Elem())
				} else {
					typeFullName(b, t.In(i))
				}
				b.WriteString(", ")
			}
			b.Truncate(b.Len() - 2)
		}
		b.WriteByte(')')

		if num := t.NumOut(); num > 0 {
			if num == 1 {
				b.WriteByte(' ')
			} else {
				b.WriteString(" (")
			}
			for i := 0; i < num; i++ {
				typeFullName(b, t.Out(i))
				b.WriteString(", ")
			}
			b.Truncate(b.Len() - 2)
			if num > 1 {
				b.WriteByte(')')
			}
		}

	case reflect.Chan:
		switch t.ChanDir() {
		case reflect.RecvDir:
			b.WriteString("<-chan ")
		case reflect.SendDir:
			b.WriteString("chan<- ")
		case reflect.BothDir:
			b.WriteString("chan ")
		}
		typeFullName(b, t.Elem())

	default:
		// Fallback to default implementation
		b.WriteString(t.String())
	}
}
