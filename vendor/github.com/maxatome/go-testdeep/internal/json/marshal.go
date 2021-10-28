// Copyright (c) 2021, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package json

import (
	"bytes"
	ejson "encoding/json"
	"fmt"
	"math"
	"sort"
	"strconv"
	"unicode/utf8"
)

type marshaler struct {
	buf    *bytes.Buffer
	indent int
	tmp    []byte
}

// Marshal returns the JSON encoding of "v". It differs from
// encoding/json.Marshal() as it only handles map[string]interface{},
// []interface{}, bool, float64, string, nil and
// encoding/json.Marshaler values. It also accepts "invalid" JSON data
// returned by MarshalJSON method.
func Marshal(v interface{}, indent int) ([]byte, error) {
	m := marshaler{
		indent: indent,
		buf:    &bytes.Buffer{},
	}
	err := m.marshal(v)
	if err != nil {
		return nil, err
	}
	return m.buf.Bytes(), nil
}

// AppendMarshal does the same as Marshal but appends the JSON
// encoding to "buf".
func AppendMarshal(buf *bytes.Buffer, v interface{}, indent int) error {
	m := marshaler{
		indent: indent,
		buf:    buf,
	}
	return m.marshal(v)
}

func (m *marshaler) marshal(v interface{}) error {
	if v == nil {
		m.buf.WriteString("null")
		return nil
	}

	switch vt := v.(type) {
	case map[string]interface{}:
		if len(vt) == 0 {
			if vt == nil {
				m.buf.WriteString("null")
			} else {
				m.buf.WriteString("{}")
			}
			break
		}
		m.indent += 2
		keys := make([]string, 0, len(vt))
		for k := range vt {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		beg := "{\n"
		for _, k := range keys {
			m.buf.WriteString(beg)
			beg = ",\n"

			saveIndent, lenBefore := m.indent, m.buf.Len()
			fmt.Fprintf(m.buf, `%*s%q: `, m.indent, "", k)
			m.indent = utf8.RuneCount(m.buf.Bytes()[lenBefore:])
			if err := m.marshal(vt[k]); err != nil {
				return err
			}
			m.indent = saveIndent
		}
		m.indent -= 2
		fmt.Fprintf(m.buf, "\n%*s}", m.indent, "")

	case []interface{}:
		if len(vt) == 0 {
			if vt == nil {
				m.buf.WriteString("null")
			} else {
				m.buf.WriteString("[]")
			}
			break
		}
		m.indent += 2
		beg := "[\n"
		for _, v := range vt {
			m.buf.WriteString(beg)
			beg = ",\n"
			fmt.Fprintf(m.buf, "%*s", m.indent, "")
			if err := m.marshal(v); err != nil {
				return err
			}
		}
		m.indent -= 2
		fmt.Fprintf(m.buf, "\n%*s]", m.indent, "")

	case string:
		fmt.Fprintf(m.buf, `%q`, vt)

	case float64:
		m.marshalFloat64(vt)

	case bool:
		if vt {
			m.buf.WriteString("true")
		} else {
			m.buf.WriteString("false")
		}

	case ejson.Marshaler:
		b, err := vt.MarshalJSON()
		if err != nil {
			return err
		}
		repl := bytes.Repeat([]byte(" "), 1+m.indent)
		repl[0] = '\n'
		m.buf.Write(bytes.Replace(b, []byte("\n"), repl, -1)) //nolint: gocritic

	default:
		return fmt.Errorf("Cannot marshal %T", vt)
	}

	return nil
}

func (m *marshaler) marshalFloat64(f float64) {
	// Contrary to JSON standard we accept to marshal NaN and ±Inf
	if math.IsInf(f, 0) || math.IsNaN(f) {
		m.tmp = strconv.AppendFloat(m.tmp[:0], f, 'g', -1, 64)
		m.buf.Write(m.tmp)
		return
	}

	// Remainder based on encoding/json.floatEncoder.encode()

	// Convert as if by ES6 number to string conversion.
	// This matches most other JSON generators.
	// See golang.org/issue/6384 and golang.org/issue/14135.
	// Like fmt %g, but the exponent cutoffs are different
	// and exponents themselves are not padded to two digits.
	abs := math.Abs(f)
	fmt := byte('f')

	if abs != 0 {
		if abs < 1e-6 || abs >= 1e21 {
			fmt = 'e'
		}
	}

	m.tmp = strconv.AppendFloat(m.tmp[:0], f, fmt, -1, 64)
	if fmt == 'e' {
		// clean up e-09 to e-9
		n := len(m.tmp)
		if n >= 4 && m.tmp[n-4] == 'e' && m.tmp[n-3] == '-' && m.tmp[n-2] == '0' {
			m.tmp[n-2] = m.tmp[n-1]
			m.tmp = m.tmp[:n-1]
		}
	}

	m.buf.Write(m.tmp)
}
