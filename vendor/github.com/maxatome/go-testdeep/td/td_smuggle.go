// Copyright (c) 2018-2024, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"unicode"
	"unicode/utf8"

	"github.com/maxatome/go-testdeep/internal/ctxerr"
	"github.com/maxatome/go-testdeep/internal/types"
	"github.com/maxatome/go-testdeep/internal/util"
)

// SmuggledGot can be returned by a [Smuggle] function to name the
// transformed / returned value.
type SmuggledGot struct {
	Name string
	Got  any
}

const smuggled = "<smuggled>"

var (
	smuggleFnsMu sync.Mutex
	smuggleFns   = map[any]reflect.Value{}

	nilError = reflect.New(types.Error).Elem()
)

func (s SmuggledGot) contextAndGot(ctx ctxerr.Context) (ctxerr.Context, reflect.Value) {
	// If the Name starts with a Letter, prefix it by a "."
	var name string
	if s.Name != "" {
		first, _ := utf8.DecodeRuneInString(s.Name)
		if unicode.IsLetter(first) {
			name = "."
		}
		name += s.Name
	} else {
		name = smuggled
	}
	return ctx.AddCustomLevel(name), reflect.ValueOf(s.Got)
}

type tdSmuggle struct {
	tdSmugglerBase
	function reflect.Value
	argType  reflect.Type
	str      string
}

var _ TestDeep = &tdSmuggle{}

type smuggleValue struct {
	Path  string
	Value reflect.Value
}

var smuggleValueType = reflect.TypeOf(smuggleValue{})

type smuggleField struct {
	Name    string
	Indexed bool
	Method  bool
}

func joinFieldsPath(path []smuggleField) string {
	var buf strings.Builder
	for i, part := range path {
		if part.Indexed {
			fmt.Fprintf(&buf, "[%s]", part.Name)
		} else {
			if i > 0 {
				buf.WriteByte('.')
			}
			buf.WriteString(part.Name)
			if part.Method {
				buf.WriteString("()")
			}
		}
	}
	return buf.String()
}

func splitFieldsPath(origPath string) ([]smuggleField, error) {
	if origPath == "" {
		return nil, fmt.Errorf("FIELDS_PATH cannot be empty")
	}

	privateField := ""
	var res []smuggleField
	for path := origPath; len(path) > 0; {
		r, _ := utf8.DecodeRuneInString(path)
		switch r {
		case '[':
			path = path[1:]
			end := strings.IndexByte(path, ']')
			if end < 0 {
				return nil, fmt.Errorf("cannot find final ']' in FIELDS_PATH %q", origPath)
			}
			res = append(res, smuggleField{Name: path[:end], Indexed: true})
			path = path[end+1:]

		case '.':
			if len(res) == 0 {
				return nil, fmt.Errorf("'.' cannot be the first rune in FIELDS_PATH %q", origPath)
			}
			path = path[1:]
			if path == "" {
				return nil, fmt.Errorf("final '.' in FIELDS_PATH %q is not allowed", origPath)
			}
			r, _ = utf8.DecodeRuneInString(path)
			if r == '.' || r == '[' {
				return nil, fmt.Errorf("unexpected %q after '.' in FIELDS_PATH %q", r, origPath)
			}
			fallthrough

		default:
			var field string
			end := strings.IndexAny(path, ".[")
			if end < 0 {
				field, path = path, ""
			} else {
				field, path = path[:end], path[end:]
			}

			if strings.HasSuffix(field, "()") {
				if len(field) == 2 {
					return nil, fmt.Errorf("missing method name before () in FIELDS_PATH %q", origPath)
				}
				for j, r := range field[:len(field)-2] {
					if j == 0 && !unicode.IsUpper(r) {
						return nil, fmt.Errorf("method name %q is not public in FIELDS_PATH %q", field, origPath)
					}
					if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
						return nil, fmt.Errorf("unexpected %q in method name %q in FIELDS_PATH %q", r, field, origPath)
					}
				}
				if privateField != "" {
					return nil, fmt.Errorf("cannot call method %s as it is based on an unexported field %q in FIELDS_PATH %q", field, privateField, origPath)
				}
				res = append(res, smuggleField{Name: field[:len(field)-2], Method: true})
			} else {
				for j, r := range field {
					if privateField == "" && j == 0 && !unicode.IsUpper(r) {
						privateField = field
					}
					if !unicode.IsLetter(r) && (j == 0 || !unicode.IsNumber(r)) {
						return nil, fmt.Errorf("unexpected %q in field name %q in FIELDS_PATH %q", r, field, origPath)
					}
				}
				res = append(res, smuggleField{Name: field})
			}
		}
	}
	return res, nil
}

func nilFieldErr(path []smuggleField) error {
	return fmt.Errorf("field %q is nil", joinFieldsPath(path))
}

func buildFieldsPathFn(path string) (func(any) (smuggleValue, error), error) {
	parts, err := splitFieldsPath(path)
	if err != nil {
		return nil, err
	}

	return func(got any) (smuggleValue, error) {
		vgot := reflect.ValueOf(got)

		for idxPart, field := range parts {
			if field.Method {
				var method reflect.Value
				for {
					method = vgot.MethodByName(field.Name)
					if !method.IsValid() {
						switch vgot.Kind() {
						case reflect.Interface, reflect.Ptr:
							if !vgot.IsNil() {
								vgot = vgot.Elem()
								continue
							}
							return smuggleValue{}, nilFieldErr(parts[:idxPart])
						}
						if idxPart > 0 {
							return smuggleValue{}, fmt.Errorf(
								"field %s (type %s) does not implement %s() method",
								joinFieldsPath(parts[:idxPart]),
								vgot.Type(),
								field.Name)
						}
						return smuggleValue{}, fmt.Errorf(
							"type %s has no method %s()", vgot.Type(), field.Name)
					}
					break
				}
				mt := method.Type()
				if mt.NumIn() != 0 ||
					(mt.NumOut() != 1 && (mt.NumOut() != 2 || mt.Out(1) != types.Error)) {
					return smuggleValue{}, fmt.Errorf(
						"cannot call %s, signature %s not handled, only func() A or func() (A, error) allowed",
						joinFieldsPath(parts[:idxPart+1]),
						method.Type())
				}
				var ret []reflect.Value
				var panicked any
				func() {
					defer func() { panicked = recover() }()
					ret = method.Call(nil)
				}()
				if panicked != nil {
					return smuggleValue{}, fmt.Errorf(
						"method %s panicked: %v",
						joinFieldsPath(parts[:idxPart+1]),
						panicked)
				}
				if len(ret) == 2 && !ret[1].IsNil() {
					return smuggleValue{}, fmt.Errorf(
						"method %s returned an error: %w",
						joinFieldsPath(parts[:idxPart+1]),
						ret[1].Interface().(error))
				}
				vgot = ret[0]
				continue
			}

			// Resolve all interface and pointer dereferences
			origKind := vgot.Kind()
			for {
				switch vgot.Kind() {
				case reflect.Interface, reflect.Ptr:
					if vgot.IsNil() {
						return smuggleValue{}, nilFieldErr(parts[:idxPart])
					}
					vgot = vgot.Elem()
					continue
				}
				break
			}

			if !field.Indexed {
				if vgot.Kind() == reflect.Struct {
					vgot = vgot.FieldByName(field.Name)
					if !vgot.IsValid() {
						return smuggleValue{}, fmt.Errorf(
							"field %q not found",
							joinFieldsPath(parts[:idxPart+1]))
					}
					continue
				}
				// Accept map but only map[string]…
				if vgot.Kind() != reflect.Map ||
					vgot.Type().Key().Kind() != reflect.String {
					deref := ""
					if origKind != vgot.Kind() {
						deref = " (after dereferencing)"
					}
					if idxPart == 0 {
						return smuggleValue{}, fmt.Errorf(
							"it is a %s%s and should be a struct or a map[string]…",
							vgot.Kind(), deref)
					}
					if parts[idxPart-1].Method {
						return smuggleValue{}, fmt.Errorf(
							"method %s returned a %s%s and should be a struct or a map[string]…",
							joinFieldsPath(parts[:idxPart]), vgot.Kind(), deref)
					}
					return smuggleValue{}, fmt.Errorf(
						"field %q is a %s%s and should be a struct or a map[string]…",
						joinFieldsPath(parts[:idxPart]), vgot.Kind(), deref)
				}
			}

			switch vgot.Kind() {
			case reflect.Map:
				tkey := vgot.Type().Key()
				var vkey reflect.Value
				switch tkey.Kind() {
				case reflect.String:
					vkey = reflect.ValueOf(field.Name)
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					i, err := strconv.ParseInt(field.Name, 10, 64)
					if err != nil {
						return smuggleValue{}, fmt.Errorf(
							"field %q, %q is not an integer and so cannot match %s map key type",
							joinFieldsPath(parts[:idxPart+1]), field.Name, tkey)
					}
					vkey = reflect.ValueOf(i).Convert(tkey)
				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
					i, err := strconv.ParseUint(field.Name, 10, 64)
					if err != nil {
						return smuggleValue{}, fmt.Errorf(
							"field %q, %q is not an unsigned integer and so cannot match %s map key type",
							joinFieldsPath(parts[:idxPart+1]), field.Name, tkey)
					}
					vkey = reflect.ValueOf(i).Convert(tkey)
				case reflect.Float32, reflect.Float64:
					f, err := strconv.ParseFloat(field.Name, 64)
					if err != nil {
						return smuggleValue{}, fmt.Errorf(
							"field %q, %q is not a float and so cannot match %s map key type",
							joinFieldsPath(parts[:idxPart+1]), field.Name, tkey)
					}
					vkey = reflect.ValueOf(f).Convert(tkey)
				case reflect.Complex64, reflect.Complex128:
					c, err := strconv.ParseComplex(field.Name, 128)
					if err != nil {
						return smuggleValue{}, fmt.Errorf(
							"field %q, %q is not a complex number and so cannot match %s map key type",
							joinFieldsPath(parts[:idxPart+1]), field.Name, tkey)
					}
					vkey = reflect.ValueOf(c).Convert(tkey)
				default:
					return smuggleValue{}, fmt.Errorf(
						"field %q, %q cannot match unsupported %s map key type",
						joinFieldsPath(parts[:idxPart+1]), field.Name, tkey)
				}
				vgot = vgot.MapIndex(vkey)
				if !vgot.IsValid() {
					return smuggleValue{}, fmt.Errorf("field %q, %q map key not found",
						joinFieldsPath(parts[:idxPart+1]), field.Name)
				}

			case reflect.Slice, reflect.Array:
				i, err := strconv.ParseInt(field.Name, 10, 64)
				if err != nil {
					return smuggleValue{}, fmt.Errorf(
						"field %q, %q is not a slice/array index",
						joinFieldsPath(parts[:idxPart+1]), field.Name)
				}
				if i < 0 {
					i = int64(vgot.Len()) + i
				}
				if i < 0 || i >= int64(vgot.Len()) {
					return smuggleValue{}, fmt.Errorf(
						"field %q, %d is out of slice/array range (len %d)",
						joinFieldsPath(parts[:idxPart+1]), i, vgot.Len())
				}
				vgot = vgot.Index(int(i))

			default:
				if idxPart == 0 {
					return smuggleValue{},
						fmt.Errorf("it is a %s, but a map, array or slice is expected",
							vgot.Kind())
				}
				return smuggleValue{}, fmt.Errorf(
					"field %q is a %s, but a map, array or slice is expected",
					joinFieldsPath(parts[:idxPart]), vgot.Kind())
			}
		}
		return smuggleValue{
			Path:  path,
			Value: vgot,
		}, nil
	}, nil
}

func getFieldsPathFn(fieldPath string) (reflect.Value, error) {
	smuggleFnsMu.Lock()
	defer smuggleFnsMu.Unlock()

	if vfn, ok := smuggleFns[fieldPath]; ok {
		return vfn, nil
	}

	fn, err := buildFieldsPathFn(fieldPath)
	if err != nil {
		return reflect.Value{}, err
	}

	vfn := reflect.ValueOf(fn)
	smuggleFns[fieldPath] = vfn
	return vfn, err
}

func getCaster(outType reflect.Type) reflect.Value {
	smuggleFnsMu.Lock()
	defer smuggleFnsMu.Unlock()

	if vfn, ok := smuggleFns[outType]; ok {
		return vfn
	}

	var fn reflect.Value

	switch outType.Kind() {
	case reflect.String:
		fn = buildCaster(outType, true)

	case reflect.Slice:
		if outType.Elem().Kind() == reflect.Uint8 {
			// Special case for slices of bytes: falls back on io.Reader if not []byte
			fn = buildCaster(outType, false)
			break
		}
		fallthrough

	default:
		// For all other types, take the received param and return
		// it. Smuggle already converted got to the type of param, so the
		// work is done.
		inOut := []reflect.Type{outType}
		fn = reflect.MakeFunc(
			reflect.FuncOf(inOut, inOut, false),
			func(args []reflect.Value) []reflect.Value { return args },
		)
	}

	smuggleFns[outType] = fn
	return fn
}

// buildCaster returns a function:
//
//	func(in any) (out outType, err error)
//
// dynamically checks…
//   - if useString is false, as outType is a slice of bytes:
//     1. in is a []byte or convertible to []byte
//     2. in implements io.Reader
//   - if useString is true, as outType is a string:
//     1. in is a []byte or convertible to string
//     2. in implements io.Reader
func buildCaster(outType reflect.Type, useString bool) reflect.Value {
	zeroRet := reflect.New(outType).Elem()

	return reflect.MakeFunc(
		reflect.FuncOf(
			[]reflect.Type{types.Interface},
			[]reflect.Type{outType, types.Error},
			false,
		),
		func(args []reflect.Value) []reflect.Value {
			if args[0].IsNil() {
				return []reflect.Value{
					zeroRet,
					reflect.ValueOf(&ctxerr.Error{
						Message:  "incompatible parameter type",
						Got:      types.RawString("nil"),
						Expected: types.RawString(outType.String() + " or convertible or io.Reader"),
					}),
				}
			}

			// 1st & only arg is always an interface
			args[0] = args[0].Elem()

			if ok, convertible := types.IsTypeOrConvertible(args[0], outType); ok {
				if convertible {
					return []reflect.Value{args[0].Convert(outType), nilError}
				}
				return []reflect.Value{args[0], nilError}
			}

			// Our caller encures Interface() can be called safely
			switch ta := args[0].Interface().(type) {
			case io.Reader:
				var b bytes.Buffer
				if _, err := b.ReadFrom(ta); err != nil {
					return []reflect.Value{
						zeroRet,
						reflect.ValueOf(&ctxerr.Error{
							Message: "an error occurred while reading from io.Reader",
							Summary: ctxerr.NewSummary(err.Error()),
						}),
					}
				}
				var buf any
				if useString {
					buf = b.String()
				} else {
					buf = b.Bytes()
				}
				return []reflect.Value{
					reflect.ValueOf(buf).Convert(outType),
					nilError,
				}

			default:
				return []reflect.Value{
					zeroRet,
					reflect.ValueOf(&ctxerr.Error{
						Message:  "incompatible parameter type",
						Got:      types.RawString(args[0].Type().String()),
						Expected: types.RawString(outType.String() + " or convertible or io.Reader"),
					}),
				}
			}
		})
}

// summary(Smuggle): changes data contents or mutates it into another
// type via a custom function or a struct fields-path before stepping
// down in favor of generic comparison process
// input(Smuggle): all

// Smuggle operator allows to change data contents or mutate it into
// another type before stepping down in favor of generic comparison
// process. Of course it is a smuggler operator. So fn is a function
// that must take one parameter whose type must be convertible to the
// type of the compared value.
//
// As convenient shortcuts, fn can be a string specifying a
// fields-path through structs, maps & slices, or any other type, in
// this case a simple cast is done (see below for details).
//
// fn must return at least one value. These value will be compared as is
// to expectedValue, here integer 28:
//
//	td.Cmp(t, "0028",
//	  td.Smuggle(func(value string) int {
//	    num, _ := strconv.Atoi(value)
//	    return num
//	  }, 28),
//	)
//
// or using an other [TestDeep] operator, here [Between](28, 30):
//
//	td.Cmp(t, "0029",
//	  td.Smuggle(func(value string) int {
//	    num, _ := strconv.Atoi(value)
//	    return num
//	  }, td.Between(28, 30)),
//	)
//
// fn can return a second boolean value, used to tell that a problem
// occurred and so stop the comparison:
//
//	td.Cmp(t, "0029",
//	  td.Smuggle(func(value string) (int, bool) {
//	    num, err := strconv.Atoi(value)
//	    return num, err == nil
//	  }, td.Between(28, 30)),
//	)
//
// fn can return a third string value which is used to describe the
// test when a problem occurred (false second boolean value):
//
//	td.Cmp(t, "0029",
//	  td.Smuggle(func(value string) (int, bool, string) {
//	    num, err := strconv.Atoi(value)
//	    if err != nil {
//	      return 0, false, "string must contain a number"
//	    }
//	    return num, true, ""
//	  }, td.Between(28, 30)),
//	)
//
// Instead of returning (X, bool) or (X, bool, string), fn can
// return (X, error). When a problem occurs, the returned error is
// non-nil, as in:
//
//	td.Cmp(t, "0029",
//	  td.Smuggle(func(value string) (int, error) {
//	    num, err := strconv.Atoi(value)
//	    return num, err
//	  }, td.Between(28, 30)),
//	)
//
// Which can be simplified to:
//
//	td.Cmp(t, "0029", td.Smuggle(strconv.Atoi, td.Between(28, 30)))
//
// Imagine you want to compare that the Year of a date is between 2010
// and 2020:
//
//	td.Cmp(t, time.Date(2015, time.May, 1, 1, 2, 3, 0, time.UTC),
//	  td.Smuggle(func(date time.Time) int { return date.Year() },
//	    td.Between(2010, 2020)),
//	)
//
// In this case the data location forwarded to next test will be
// something like "DATA.MyTimeField<smuggled>", but you can act on it
// too by returning a [SmuggledGot] struct (by value or by address):
//
//	td.Cmp(t, time.Date(2015, time.May, 1, 1, 2, 3, 0, time.UTC),
//	  td.Smuggle(func(date time.Time) SmuggledGot {
//	    return SmuggledGot{
//	      Name: "Year",
//	      Got:  date.Year(),
//	    }
//	  }, td.Between(2010, 2020)),
//	)
//
// then the data location forwarded to next test will be something like
// "DATA.MyTimeField.Year". The "." between the current path (here
// "DATA.MyTimeField") and the returned Name "Year" is automatically
// added when Name starts with a Letter.
//
// Note that [SmuggledGot] and [*SmuggledGot] returns are treated
// equally, and they are only used when fn has only one returned value
// or when the second boolean returned value is true.
//
// Of course, all cases can go together:
//
//	// Accepts a "YYYY/mm/DD HH:MM:SS" string to produce a time.Time and tests
//	// whether this date is contained between 2 hours before now and now.
//	td.Cmp(t, "2020-01-25 12:13:14",
//	  td.Smuggle(func(date string) (*SmuggledGot, bool, string) {
//	    date, err := time.Parse("2006/01/02 15:04:05", date)
//	    if err != nil {
//	      return nil, false, `date must conform to "YYYY/mm/DD HH:MM:SS" format`
//	    }
//	    return &SmuggledGot{
//	      Name: "Date",
//	      Got:  date,
//	    }, true, ""
//	  }, td.Between(time.Now().Add(-2*time.Hour), time.Now())),
//	)
//
// or:
//
//	// Accepts a "YYYY/mm/DD HH:MM:SS" string to produce a time.Time and tests
//	// whether this date is contained between 2 hours before now and now.
//	td.Cmp(t, "2020-01-25 12:13:14",
//	  td.Smuggle(func(date string) (*SmuggledGot, error) {
//	    date, err := time.Parse("2006/01/02 15:04:05", date)
//	    if err != nil {
//	      return nil, err
//	    }
//	    return &SmuggledGot{
//	      Name: "Date",
//	      Got:  date,
//	    }, nil
//	  }, td.Between(time.Now().Add(-2*time.Hour), time.Now())),
//	)
//
// Smuggle can also be used to access a struct field embedded in
// several struct layers.
//
//	type A struct{ Num int }
//	// func (a *A) String() string { return fmt.Sprintf("Num is %d", a.Num) }
//	type B struct{ As map[string]*A }
//	type C struct{ B B }
//	got := C{B: B{As: map[string]*A{"foo": {Num: 12}}}}
//
//	// Tests that got.B.A.Num is 12
//	td.Cmp(t, got,
//	  td.Smuggle(func(c C) int {
//	    return c.B.As["foo"].Num
//	  }, 12))
//
// As brought up above, a fields-path can be passed as fn value
// instead of a function pointer. Using this feature, the [Cmp]
// call in the above example can be rewritten as follows:
//
//	// Tests that got.B.As["foo"].Num is 12
//	td.Cmp(t, got, td.Smuggle("B.As[foo].Num", 12))
//
// For convenience, if a map[string]… is addressed, it can be done
// like a struct field as "foo" key in:
//
//	td.Cmp(t, got, td.Smuggle("B.As.foo.Num", 12))
//
// In addition, simple public methods can also be called like in:
//
//	td.Cmp(t, got, td.Smuggle("B.As[foo].String()", "Num is 12"))
//
// Allowed methods must not take any parameter and must return one
// value or a value and an error. For the latter case, if the method
// returns a non-nil error, the comparison fails. The comparison also
// fails if a panic occurs or if a method cannot be called. No private
// fields should be traversed before calling the method. For fun,
// consider a more complex example involving [reflect] and chaining
// method calls:
//
//	got := reflect.Valueof(&C{B: B{As: map[string]*A{"foo": {Num: 12}}}})
//	td.Cmp(t, got, td.Smuggle("Elem().Interface().B.As[foo].String()", "Num is 12"))
//
// Contrary to [JSONPointer] operator, private fields can be followed
// and public methods on public fields can be called. Arrays, slices
// and maps work using the index/key inside square brackets (e.g. [12]
// or [foo]). Maps work only for simple key types (string or numbers),
// without "" when using strings (e.g. [foo]).
//
// Behind the scenes, a temporary function is automatically created to
// achieve the same goal, but adds some checks against nil values and
// auto-dereferences interfaces and pointers, even on several levels,
// like in:
//
//	type A struct{ N any }
//	num := 12
//	pnum := &num
//	td.Cmp(t, A{N: &pnum}, td.Smuggle("N", 12))
//
// Last but not least, a simple type can be passed as fn to operate
// a cast, handling specifically strings and slices of bytes:
//
//	td.Cmp(t, `{"foo":1}`, td.Smuggle(json.RawMessage{}, td.JSON(`{"foo":1}`)))
//	// or equally
//	td.Cmp(t, `{"foo":1}`, td.Smuggle(json.RawMessage(nil), td.JSON(`{"foo":1}`)))
//
// converts on the fly a string to a [json.RawMessage] so [JSON] operator
// can parse it as JSON. This is mostly a shortcut for:
//
//	td.Cmp(t, `{"foo":1}`, td.Smuggle(
//	  func(r json.RawMessage) json.RawMessage { return r },
//	  td.JSON(`{"foo":1}`)))
//
// except that for strings and slices of bytes (like here), it accepts
// [io.Reader] interface too:
//
//	var body io.Reader
//	// …
//	td.Cmp(t, body, td.Smuggle(json.RawMessage{}, td.JSON(`{"foo":1}`)))
//	// or equally
//	td.Cmp(t, body, td.Smuggle(json.RawMessage(nil), td.JSON(`{"foo":1}`)))
//
// This last example allows to easily inject body content into JSON
// operator.
//
// The difference between Smuggle and [Code] operators is that [Code]
// is used to do a final comparison while Smuggle transforms the data
// and then steps down in favor of generic comparison
// process. Moreover, the type accepted as input for the function is
// more lax to facilitate the writing of tests (e.g. the function can
// accept a float64 and the got value be an int). See examples. On the
// other hand, the output type is strict and must match exactly the
// expected value type. The fields-path string fn shortcut and the
// cast feature are not available with [Code] operator.
//
// TypeBehind method returns the [reflect.Type] of only parameter of
// fn. For the case where fn is a fields-path, it is always
// any, as the type can not be known in advance.
//
// See also [Code], [JSONPointer] and [Flatten].
//
// [json.RawMessage]: https://pkg.go.dev/encoding/json#RawMessage
func Smuggle(fn, expectedValue any) TestDeep {
	s := tdSmuggle{
		tdSmugglerBase: newSmugglerBase(expectedValue),
	}

	const usage = "(FUNC|FIELDS_PATH|ANY_TYPE, TESTDEEP_OPERATOR|EXPECTED_VALUE)"
	const fullUsage = "Smuggle" + usage

	var vfn reflect.Value

	switch rfn := fn.(type) {
	case reflect.Type:
		switch rfn.Kind() {
		case reflect.Func, reflect.Invalid, reflect.Interface:
			s.err = ctxerr.OpBad("Smuggle",
				"usage: Smuggle%s, ANY_TYPE reflect.Type cannot be Func nor Interface", usage)
			return &s

		default:
			vfn = getCaster(rfn)
			s.str = "type:" + rfn.String()
		}

	case string:
		if rfn == "" {
			vfn = getCaster(reflect.TypeOf(fn))
			s.str = "type:string"
			break
		}
		var err error
		vfn, err = getFieldsPathFn(rfn)
		if err != nil {
			s.err = ctxerr.OpBad("Smuggle", "Smuggle%s: %s", usage, err)
			return &s
		}
		s.str = strconv.Quote(rfn)

	default:
		vfn = reflect.ValueOf(fn)
		switch vfn.Kind() {
		case reflect.Func:
			s.str = vfn.Type().String()
			// nothing to check

		case reflect.Invalid, reflect.Interface:
			s.err = ctxerr.OpBad("Smuggle",
				"usage: Smuggle%s, ANY_TYPE cannot be nil nor Interface", usage)
			return &s

		default:
			typ := vfn.Type()
			vfn = getCaster(typ)
			s.str = "type:" + typ.String()
		}
	}

	fnType := vfn.Type()
	if fnType.IsVariadic() || fnType.NumIn() != 1 {
		s.err = ctxerr.OpBad("Smuggle", fullUsage+": FUNC must take only one non-variadic argument")
		return &s
	}

	switch fnType.NumOut() {
	case 3: // (value, bool, string)
		if fnType.Out(2).Kind() != reflect.String {
			break
		}
		fallthrough

	case 2:
		// (value, *bool*) or (value, *bool*, string)
		if fnType.Out(1).Kind() != reflect.Bool &&
			// (value, *error*)
			(fnType.NumOut() > 2 ||
				fnType.Out(1) != types.Error) {
			break
		}
		fallthrough

	case 1: // (value)
		if vfn.IsNil() {
			s.err = ctxerr.OpBad("Smuggle", "Smuggle(FUNC): FUNC cannot be a nil function")
			return &s
		}

		s.argType = fnType.In(0)
		s.function = vfn

		if !s.isTestDeeper {
			s.expectedValue = reflect.ValueOf(expectedValue)
		}
		return &s
	}

	s.err = ctxerr.OpBad("Smuggle",
		fullUsage+": FUNC must return value or (value, bool) or (value, bool, string) or (value, error)")
	return &s
}

func (s *tdSmuggle) laxConvert(got reflect.Value) (reflect.Value, bool) {
	if got.IsValid() {
		if types.IsConvertible(got, s.argType) {
			return got.Convert(s.argType), true
		}
	} else if s.argType == types.Interface {
		// nil only accepted if any expected
		return reflect.New(types.Interface).Elem(), true
	}
	return got, false
}

func (s *tdSmuggle) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	if s.err != nil {
		return ctx.CollectError(s.err)
	}

	got, ok := s.laxConvert(got)
	if !ok {
		if ctx.BooleanError {
			return ctxerr.BooleanError
		}
		err := ctxerr.Error{
			Message:  "incompatible parameter type",
			Expected: types.RawString(s.argType.String()),
		}
		if got.IsValid() {
			err.Got = types.RawString(got.Type().String())
		} else {
			err.Got = types.RawString("nil")
		}
		return ctx.CollectError(&err)
	}

	// Refuse to override unexported fields access in this case. It is a
	// choice, as we think it is better to work on surrounding struct
	// instead.
	if !got.CanInterface() {
		if ctx.BooleanError {
			return ctxerr.BooleanError
		}
		return ctx.CollectError(&ctxerr.Error{
			Message: "cannot smuggle unexported field",
			Summary: ctxerr.NewSummary("work on surrounding struct instead"),
		})
	}

	ret := s.function.Call([]reflect.Value{got})
	if len(ret) == 1 ||
		(ret[1].Kind() == reflect.Bool && ret[1].Bool()) ||
		(ret[1].Kind() == reflect.Interface && ret[1].IsNil()) {
		newGot := ret[0]

		var newCtx ctxerr.Context
		if newGot.IsValid() {
			switch newGot.Type() {
			case smuggledGotType:
				newCtx, newGot = newGot.Interface().(SmuggledGot).contextAndGot(ctx)

			case smuggledGotPtrType:
				if smGot := newGot.Interface().(*SmuggledGot); smGot == nil {
					newCtx, newGot = ctx, reflect.ValueOf(nil)
				} else {
					newCtx, newGot = smGot.contextAndGot(ctx)
				}

			case smuggleValueType:
				smv := newGot.Interface().(smuggleValue)
				newCtx, newGot = ctx.AddCustomLevel("."+smv.Path), smv.Value

			default:
				newCtx = ctx.AddCustomLevel(smuggled)
			}
		}
		return deepValueEqual(newCtx, newGot, s.expectedValue)
	}

	if ctx.BooleanError {
		return ctxerr.BooleanError
	}

	var reason string
	switch len(ret) {
	case 3: // (value, false, string)
		reason = ret[2].String()
	case 2:
		// (value, error)
		if ret[1].Kind() == reflect.Interface {
			// For internal use only
			if cErr, ok := ret[1].Interface().(*ctxerr.Error); ok {
				return ctx.CollectError(cErr)
			}
			reason = ret[1].Interface().(error).Error()
		}
		// (value, false)
	}
	return ctx.CollectError(&ctxerr.Error{
		Message: "ran smuggle code with %% as argument",
		Summary: ctxerr.NewSummaryReason(got, reason),
	})
}

func (s *tdSmuggle) HandleInvalid() bool {
	return true // Knows how to handle untyped nil values (aka invalid values)
}

func (s *tdSmuggle) String() string {
	if s.err != nil {
		return s.stringError()
	}
	return "Smuggle(" + s.str + ", " + util.ToString(s.expectedValue) + ")"
}

func (s *tdSmuggle) TypeBehind() reflect.Type {
	if s.err != nil {
		return nil
	}
	return s.argType
}
