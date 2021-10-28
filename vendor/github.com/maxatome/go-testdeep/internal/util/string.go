// Copyright (c) 2018, Maxime Soulé
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

// ToString does its best to stringify val.
func ToString(val interface{}) string {
	if val == nil {
		return "nil"
	}

	switch tval := val.(type) {
	case reflect.Value:
		newVal, ok := dark.GetInterface(tval, true)
		if ok {
			return ToString(newVal)
		}

	case []reflect.Value:
		var buf bytes.Buffer
		SliceToBuffer(&buf, tval)
		return buf.String()

		// no "(string) " prefix for printable strings
	case string:
		return tdutil.FormatString(tval)

		// no "(int) " prefix for ints
	case int:
		return strconv.Itoa(tval)

		// no "(bool)" prefix for booleans
	case bool:
		return TernStr(tval, "true", "false")

	case types.TestDeepStringer:
		return tval.String()
	}

	return tdutil.SpewString(val)
}

// IndentString indents str lines (from 2nd one = 1st line is not
// indented) by indent.
func IndentString(str, indent string) string {
	return strings.Replace(str, "\n", "\n"+indent, -1) //nolint: gocritic
}

// IndentStringIn indents str lines (from 2nd one = 1st line is not
// indented) by indent and write it to w. Before each end of line, colOff is inserted, and after each indent on new line, colOn is inserted.
func IndentStringIn(w io.Writer, str, indent, colOn, colOff string) {
	repl := strings.NewReplacer("\n", colOff+"\n"+indent+colOn)
	repl.WriteString(w, str) //nolint: errcheck
}

// SliceToBuffer stringifies items slice into buf then returns buf.
func SliceToBuffer(buf *bytes.Buffer, items []reflect.Value) *bytes.Buffer {
	buf.WriteByte('(')

	begLine := bytes.LastIndexByte(buf.Bytes(), '\n') + 1
	prefix := strings.Repeat(" ", buf.Len()-begLine)

	if len(items) < 2 {
		if len(items) > 0 {
			buf.WriteString(IndentString(ToString(items[0]), prefix))
		}
	} else {
		for idx, item := range items {
			if idx != 0 {
				buf.WriteString(prefix)
			}
			buf.WriteString(IndentString(ToString(item), prefix))
			buf.WriteString(",\n")
		}
		buf.Truncate(buf.Len() - 2)
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
