/*
 * Copyright (c) 2013-2016 Dave Collins <dave@davec.name>
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package spew

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"

	"github.com/maxatome/go-testdeep/internal/sort"
)

var (
	// uint8Type is a reflect.Type representing a uint8.  It is used to
	// convert cgo types to uint8 slices for hexdumping.
	uint8Type = reflect.TypeOf(uint8(0))

	// cCharSuffix is the suffix that matches a cgo char.
	// It is used to detect character arrays to hexdump them.
	cCharSuffix = `._Ctype_char`

	// cUnsignedCharSuffix is the suffix that matches a cgo unsigned
	// char.  It is used to detect unsigned character arrays to hexdump
	// them.
	cUnsignedCharSuffix = `._Ctype_unsignedchar`

	// cUint8tCharSuffix is the suffix that matches a cgo uint8_t.
	// It is used to detect uint8_t arrays to hexdump them.
	cUint8tCharSuffix = `._Ctype_uint8_t`
)

// dumpState contains information about the state of a dump operation.
type dumpState struct {
	w                io.Writer
	depth            int
	pointers         map[uintptr]int
	ignoreNextType   bool
	ignoreNextIndent bool
	cs               *ConfigState
}

// indent performs indentation according to the depth level and cs.Indent
// option.
func (d *dumpState) indent() {
	if d.ignoreNextIndent {
		d.ignoreNextIndent = false
		return
	}
	d.w.Write(bytes.Repeat([]byte(d.cs.Indent), d.depth)) //nolint: errcheck
}

// unpackValue returns values inside of non-nil interfaces when possible.
// This is useful for data types like structs, arrays, slices, and maps which
// can contain varying types packed inside an interface.
func (d *dumpState) unpackValue(v reflect.Value) reflect.Value {
	if v.Kind() == reflect.Interface && !v.IsNil() {
		v = v.Elem()
	}
	return v
}

// dumpPtr handles formatting of pointers by indirecting them as necessary.
// nolint: errcheck
func (d *dumpState) dumpPtr(v reflect.Value) {
	// Remove pointers at or below the current depth from map used to detect
	// circular refs.
	for k, depth := range d.pointers {
		if depth >= d.depth {
			delete(d.pointers, k)
		}
	}

	// Keep list of all dereferenced pointers to show later.
	pointerChain := make([]uintptr, 0)

	// Figure out how many levels of indirection there are by dereferencing
	// pointers and unpacking interfaces down the chain while detecting circular
	// references.
	nilFound := false
	cycleFound := false
	indirects := 0
	ve := v
	for ve.Kind() == reflect.Ptr {
		if ve.IsNil() {
			nilFound = true
			break
		}
		indirects++
		addr := ve.Pointer()
		pointerChain = append(pointerChain, addr)
		if pd, ok := d.pointers[addr]; ok && pd < d.depth {
			cycleFound = true
			indirects--
			break
		}
		d.pointers[addr] = d.depth

		ve = ve.Elem()
		if ve.Kind() == reflect.Interface {
			if ve.IsNil() {
				nilFound = true
				break
			}
			ve = ve.Elem()
		}
	}

	// Display type information.
	d.w.Write(openParenBytes)
	d.w.Write(bytes.Repeat(asteriskBytes, indirects))
	d.w.Write([]byte(ve.Type().String()))
	d.w.Write(closeParenBytes)

	// Display pointer information.
	if !d.cs.DisablePointerAddresses && len(pointerChain) > 0 {
		d.w.Write(openParenBytes)
		for i, addr := range pointerChain {
			if i > 0 {
				d.w.Write(pointerChainBytes)
			}
			printHexPtr(d.w, addr)
		}
		d.w.Write(closeParenBytes)
	}

	// Display dereferenced value.
	d.w.Write(openParenBytes)
	switch {
	case nilFound:
		d.w.Write(nilAngleBytes)

	case cycleFound:
		d.w.Write(circularBytes)

	default:
		d.ignoreNextType = true
		d.dump(ve)
	}
	d.w.Write(closeParenBytes)
}

// dumpSlice handles formatting of arrays and slices.  Byte (uint8 under
// reflection) arrays and slices are dumped in hexdump -C fashion.
// nolint: errcheck
func (d *dumpState) dumpSlice(v reflect.Value) {
	// Determine whether this type should be hex dumped or not.  Also,
	// for types which should be hexdumped, try to use the underlying data
	// first, then fall back to trying to convert them to a uint8 slice.
	var buf []uint8
	doHexDump := false
	numEntries := v.Len()
	if numEntries > 0 {
		vt := v.Index(0).Type()
		vts := vt.String()
		// C types that need to be converted.
		doConvert := strings.HasSuffix(vts, cCharSuffix) ||
			strings.HasSuffix(vts, cUnsignedCharSuffix) ||
			strings.HasSuffix(vts, cUint8tCharSuffix)

		// Try to use existing uint8 slices and fall back to converting
		// and copying if that fails.
		if !doConvert && vt.Kind() == reflect.Uint8 {
			// The underlying data needs to be converted if it can't
			// be type asserted to a uint8 slice.
			doConvert = true

			// We need an addressable interface to convert the type
			// to a byte slice.  However, the reflect package won't
			// give us an interface on certain things like
			// unexported struct fields in order to enforce
			// visibility rules.  We use unsafe, when available, to
			// bypass these restrictions since this package does not
			// mutate the values.
			vs := v
			if !vs.CanInterface() || !vs.CanAddr() {
				vs = UnsafeReflectValue(vs)
			}
			if !UnsafeDisabled {
				vs = vs.Slice(0, numEntries)

				// Use the existing uint8 slice if it can be type asserted.
				if slice, ok := vs.Interface().([]uint8); ok {
					buf = slice
					doHexDump = true
					doConvert = false
				}
			}
		}

		// Copy and convert the underlying type if needed.
		if doConvert && vt.ConvertibleTo(uint8Type) {
			// Convert and copy each element into a uint8 byte
			// slice.
			buf = make([]uint8, numEntries)
			for i := 0; i < numEntries; i++ {
				vv := v.Index(i)
				buf[i] = uint8(vv.Convert(uint8Type).Uint())
			}
			doHexDump = true
		}
	}

	// Hexdump the entire slice as needed.
	if doHexDump {
		indent := strings.Repeat(d.cs.Indent, d.depth)
		str := indent + hex.Dump(buf)
		str = strings.ReplaceAll(str, "\n", "\n"+indent)
		str = strings.TrimRight(str, d.cs.Indent)
		d.w.Write([]byte(str))
		return
	}

	// Recursively call dump for each item.
	for i := 0; i < numEntries; i++ {
		d.dump(d.unpackValue(v.Index(i)))
		if i < (numEntries - 1) {
			d.w.Write(commaNewlineBytes)
		} else {
			d.w.Write(newlineBytes)
		}
	}
}

// dump is the main workhorse for dumping a value.  It uses the passed reflect
// value to figure out what kind of object we are dealing with and formats it
// appropriately.  It is a recursive function, however circular data structures
// are detected and handled properly.
// nolint: errcheck
func (d *dumpState) dump(v reflect.Value) {
	// Handle invalid reflect values immediately.
	kind := v.Kind()
	if kind == reflect.Invalid {
		d.w.Write(invalidAngleBytes)
		return
	}

	// Handle pointers specially.
	if kind == reflect.Ptr {
		d.indent()
		d.dumpPtr(v)
		return
	}

	// Print type information unless already handled elsewhere.
	if !d.ignoreNextType {
		d.indent()
		d.w.Write(openParenBytes)
		d.w.Write([]byte(v.Type().String()))
		d.w.Write(closeParenBytes)
		d.w.Write(spaceBytes)
	}
	d.ignoreNextType = false

	// Display length and capacity if the built-in len and cap functions
	// work with the value's kind and the len/cap itself is non-zero.
	valueLen, valueCap := 0, 0
	switch v.Kind() {
	case reflect.Array, reflect.Slice, reflect.Chan:
		valueLen, valueCap = v.Len(), v.Cap()
	case reflect.Map, reflect.String:
		valueLen = v.Len()
	}
	if valueLen != 0 || d.cs.EnableCapacities && valueCap != 0 {
		d.w.Write(openParenBytes)
		if valueLen != 0 {
			d.w.Write(lenEqualsBytes)
			printInt(d.w, int64(valueLen), 10)
		}
		if d.cs.EnableCapacities && valueCap != 0 {
			if valueLen != 0 {
				d.w.Write(spaceBytes)
			}
			d.w.Write(capEqualsBytes)
			printInt(d.w, int64(valueCap), 10)
		}
		d.w.Write(closeParenBytes)
		d.w.Write(spaceBytes)
	}

	// Call fmt.Stringer/error interfaces if they exist and the handle
	// methods flag is enabled
	if !d.cs.DisableMethods &&
		kind != reflect.Invalid && kind != reflect.Interface {
		if handled := handleMethods(d.cs, d.w, v); handled {
			return
		}
	}

	switch kind {
	case reflect.Invalid, reflect.Ptr:
		// Do nothing.  We should never get here since invalid & pointers
		// have already been handled above.

	case reflect.Bool:
		printBool(d.w, v.Bool())

	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		printInt(d.w, v.Int(), 10)

	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		printUint(d.w, v.Uint(), 10)

	case reflect.Float32:
		printFloat(d.w, v.Float(), 32)

	case reflect.Float64:
		printFloat(d.w, v.Float(), 64)

	case reflect.Complex64:
		printComplex(d.w, v.Complex(), 32)

	case reflect.Complex128:
		printComplex(d.w, v.Complex(), 64)

	case reflect.Slice:
		if v.IsNil() {
			d.w.Write(nilAngleBytes)
			break
		}
		fallthrough

	case reflect.Array:
		d.w.Write(openBraceNewlineBytes)
		d.depth++
		if d.cs.MaxDepth != 0 && d.depth > d.cs.MaxDepth {
			d.indent()
			d.w.Write(maxNewlineBytes)
		} else {
			d.dumpSlice(v)
		}
		d.depth--
		d.indent()
		d.w.Write(closeBraceBytes)

	case reflect.String:
		d.w.Write([]byte(strconv.Quote(v.String())))

	case reflect.Interface:
		// The only time we should get here is for nil interfaces due to
		// unpackValue calls.
		if v.IsNil() {
			d.w.Write(nilAngleBytes)
		}

	case reflect.Map:
		// nil maps should be indicated as different than empty maps
		if v.IsNil() {
			d.w.Write(nilAngleBytes)
			break
		}

		d.w.Write(openBraceNewlineBytes)
		d.depth++
		if d.cs.MaxDepth != 0 && d.depth > d.cs.MaxDepth {
			d.indent()
			d.w.Write(maxNewlineBytes)
		} else {
			i := v.Len()
			sort.MapEach(v, func(key reflect.Value, val reflect.Value) bool {
				d.dump(d.unpackValue(key))
				d.w.Write(colonSpaceBytes)
				d.ignoreNextIndent = true
				d.dump(d.unpackValue(val))
				i--
				if i > 0 {
					d.w.Write(commaNewlineBytes)
				} else {
					d.w.Write(newlineBytes)
				}
				return true
			})
		}
		d.depth--
		d.indent()
		d.w.Write(closeBraceBytes)

	case reflect.Struct:
		d.w.Write(openBraceNewlineBytes)
		d.depth++
		if d.cs.MaxDepth != 0 && d.depth > d.cs.MaxDepth {
			d.indent()
			d.w.Write(maxNewlineBytes)
		} else {
			vt := v.Type()
			numFields := v.NumField()
			for i := 0; i < numFields; i++ {
				vf := v.Field(i)
				if vf.IsZero() {
					continue
				}
				d.indent()
				vtf := vt.Field(i)
				d.w.Write([]byte(vtf.Name))
				d.w.Write(colonSpaceBytes)
				d.ignoreNextIndent = true
				d.dump(d.unpackValue(vf))
				if i < (numFields - 1) {
					d.w.Write(commaNewlineBytes)
				} else {
					d.w.Write(newlineBytes)
				}
			}
		}
		d.depth--
		d.indent()
		d.w.Write(closeBraceBytes)

	case reflect.Uintptr:
		printHexPtr(d.w, uintptr(v.Uint()))

	case reflect.UnsafePointer, reflect.Chan, reflect.Func:
		printHexPtr(d.w, v.Pointer())

	// There were not any other types at the time this code was written, but
	// fall back to letting the default fmt package handle it in case any new
	// types are added.
	default:
		if v.CanInterface() {
			fmt.Fprintf(d.w, "%v", v.Interface())
		} else {
			fmt.Fprintf(d.w, "%v", v.String())
		}
	}
}

// fdump is a helper function to consolidate the logic from the various public
// methods which take varying writers and config states.
// nolint: errcheck
func fdump(cs *ConfigState, w io.Writer, arg any) {
	if arg == nil {
		w.Write(interfaceBytes)
		w.Write(spaceBytes)
		w.Write(nilAngleBytes)
		w.Write(newlineBytes)
		return
	}

	d := dumpState{w: w, cs: cs}
	d.pointers = make(map[uintptr]int)
	d.dump(reflect.ValueOf(arg))
	d.w.Write(newlineBytes)
}

// Sdump returns the string representation of a using config[0] if passed.
func Sdump(a any, config ...ConfigState) string {
	var buf strings.Builder
	cs := Config
	if len(config) > 0 {
		cs = config[0]
	}
	fdump(&cs, &buf, a)
	return strings.TrimRight(buf.String(), "\n")
}
