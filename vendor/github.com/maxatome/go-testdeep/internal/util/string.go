// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package util

import (
	"bytes"
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
func IndentString(str string, indent string) string {
	return strings.Replace(str, "\n", "\n"+indent, -1)
}

// IndentStringIn indents str lines (from 2nd one = 1st line is not
// indented) by indent and write it to w.
func IndentStringIn(w io.Writer, str string, indent string) {
	repl := strings.NewReplacer("\n", "\n"+indent)
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
