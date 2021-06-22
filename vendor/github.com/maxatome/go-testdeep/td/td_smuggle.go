// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/maxatome/go-testdeep/internal/ctxerr"
	"github.com/maxatome/go-testdeep/internal/types"
)

// SmuggledGot can be returned by a Smuggle function to name the
// transformed / returned value.
type SmuggledGot struct {
	Name string
	Got  interface{}
}

const smuggled = "<smuggled>"

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
}

var _ TestDeep = &tdSmuggle{}

type smuggleValue struct {
	Path  string
	Value reflect.Value
}

var smuggleValueType = reflect.TypeOf(smuggleValue{})

func nilFieldErr(path []string) error {
	return fmt.Errorf("field `%s' is nil", strings.Join(path, "."))
}

func buildStructFieldFn(path string) (func(interface{}) (smuggleValue, error), error) {
	parts := strings.Split(path, ".")

	// Check path parts validity
	for i, p := range parts {
		for j, r := range p {
			if !unicode.IsLetter(r) && (j == 0 || !unicode.IsNumber(r)) {
				return nil, fmt.Errorf("bad field name `%s' in FIELDS_PATH", parts[i])
			}
		}
	}

	return func(got interface{}) (smuggleValue, error) {
		vgot := reflect.ValueOf(got)

		for idxPart, fieldName := range parts {
			// Resolve only one interface{} dereference
			if vgot.Kind() == reflect.Interface {
				if vgot.IsNil() {
					return smuggleValue{}, nilFieldErr(parts[:idxPart])
				}
				vgot = vgot.Elem()
			}

			// Resolve multiple ptr dereferences
			for {
				if vgot.Kind() != reflect.Ptr {
					break
				}
				if vgot.IsNil() {
					return smuggleValue{}, nilFieldErr(parts[:idxPart])
				}
				vgot = vgot.Elem()
			}

			if vgot.Kind() != reflect.Struct {
				if idxPart == 0 {
					return smuggleValue{}, fmt.Errorf("it is not a struct and should be")
				}
				return smuggleValue{}, fmt.Errorf(
					"field `%s' is not a struct and should be",
					strings.Join(parts[:idxPart], "."))
			}

			vgot = vgot.FieldByName(fieldName)
			if !vgot.IsValid() {
				return smuggleValue{}, fmt.Errorf("field `%s' not found",
					strings.Join(parts[:idxPart+1], "."))
			}
		}
		return smuggleValue{
			Path:  path,
			Value: vgot,
		}, nil
	}, nil
}

// summary(Smuggle): changes data contents or mutates it into another
// type via a custom function or a struct fields-path before stepping
// down in favor of generic comparison process
// input(Smuggle): all

// Smuggle operator allows to change data contents or mutate it into
// another type before stepping down in favor of generic comparison
// process. Of course it is a smuggler operator. So "fn" is a function
// that must take one parameter whose type must be convertible to the
// type of the compared value (as a convenient shortcut, "fn" can be a
// string specifying a fields-path through structs, see below for
// details).
//
// "fn" must return at least one value. These value will be compared as is
// to "expectedValue", here integer 28:
//
//   td.Cmp(t, "0028",
//     td.Smuggle(func(value string) int {
//       num, _ := strconv.Atoi(value)
//       return num
//     }, 28),
//   )
//
// or using an other TestDeep operator, here Between(28, 30):
//
//   td.Cmp(t, "0029",
//     td.Smuggle(func(value string) int {
//       num, _ := strconv.Atoi(value)
//       return num
//     }, td.Between(28, 30)),
//   )
//
// "fn" can return a second boolean value, used to tell that a problem
// occurred and so stop the comparison:
//
//   td.Cmp(t, "0029",
//     td.Smuggle(func(value string) (int, bool) {
//       num, err := strconv.Atoi(value)
//       return num, err == nil
//     }, td.Between(28, 30)),
//   )
//
// "fn" can return a third string value which is used to describe the
// test when a problem occurred (false second boolean value):
//
//   td.Cmp(t, "0029",
//     td.Smuggle(func(value string) (int, bool, string) {
//       num, err := strconv.Atoi(value)
//       if err != nil {
//         return 0, false, "string must contain a number"
//       }
//       return num, true, ""
//     }, td.Between(28, 30)),
//   )
//
// Instead of returning (X, bool) or (X, bool, string), "fn" can
// return (X, error). When a problem occurs, the returned error is
// non-nil, as in:
//
//   td.Cmp(t, "0029",
//     td.Smuggle(func(value string) (int, error) {
//       num, err := strconv.Atoi(value)
//       return num, err
//     }, td.Between(28, 30)),
//   )
//
// Which can be simplified to:
//
//   td.Cmp(t, "0029", td.Smuggle(strconv.Atoi, td.Between(28, 30)))
//
// Imagine you want to compare that the Year of a date is between 2010
// and 2020:
//
//   td.Cmp(t, time.Date(2015, time.May, 1, 1, 2, 3, 0, time.UTC),
//     td.Smuggle(func(date time.Time) int { return date.Year() },
//       td.Between(2010, 2020)),
//   )
//
// In this case the data location forwarded to next test will be
// something like "DATA.MyTimeField<smuggled>", but you can act on it
// too by returning a SmuggledGot struct (by value or by address):
//
//   td.Cmp(t, time.Date(2015, time.May, 1, 1, 2, 3, 0, time.UTC),
//     td.Smuggle(func(date time.Time) SmuggledGot {
//       return SmuggledGot{
//         Name: "Year",
//         Got:  date.Year(),
//       }
//     }, td.Between(2010, 2020)),
//   )
//
// then the data location forwarded to next test will be something like
// "DATA.MyTimeField.Year". The "."  between the current path (here
// "DATA.MyTimeField") and the returned Name "Year" is automatically
// added when Name starts with a Letter.
//
// Note that SmuggledGot and *SmuggledGot returns are treated equally,
// and they are only used when "fn" has only one returned value or
// when the second boolean returned value is true.
//
// Of course, all cases can go together:
//
//   // Accepts a "YYYY/mm/DD HH:MM:SS" string to produce a time.Time and tests
//   // whether this date is contained between 2 hours before now and now.
//   td.Cmp(t, "2020-01-25 12:13:14",
//     td.Smuggle(func(date string) (*SmuggledGot, bool, string) {
//       date, err := time.Parse("2006/01/02 15:04:05", date)
//       if err != nil {
//         return nil, false, `date must conform to "YYYY/mm/DD HH:MM:SS" format`
//       }
//       return &SmuggledGot{
//         Name: "Date",
//         Got:  date,
//       }, true, ""
//     }, td.Between(time.Now().Add(-2*time.Hour), time.Now())),
//   )
//
// or:
//
//   // Accepts a "YYYY/mm/DD HH:MM:SS" string to produce a time.Time and tests
//   // whether this date is contained between 2 hours before now and now.
//   td.Cmp(t, "2020-01-25 12:13:14",
//     td.Smuggle(func(date string) (*SmuggledGot, error) {
//       date, err := time.Parse("2006/01/02 15:04:05", date)
//       if err != nil {
//         return nil, err
//       }
//       return &SmuggledGot{
//         Name: "Date",
//         Got:  date,
//       }, nil
//     }, td.Between(time.Now().Add(-2*time.Hour), time.Now())),
//   )
//
// Smuggle can also be used to access a struct field embedded in
// several struct layers.
//
//   type A struct{ Num int }
//   type B struct{ A *A }
//   type C struct{ B B }
//   got := C{B: B{A: &A{Num: 12}}}
//
//   // Tests that got.B.A.Num is 12
//   td.Cmp(t, got,
//     td.Smuggle(func(c C) int {
//       return c.B.A.Num
//     }, 12))
//
// As brought up above, a field-path can be passed as "fn" value
// instead of a function pointer. Using this feature, the Cmp
// call in the above example can be rewritten as follows:
//
//   // Tests that got.B.A.Num is 12
//   td.Cmp(t, got, td.Smuggle("B.A.Num", 12))
//
// Behind the scenes, a temporary function is automatically created to
// achieve the same goal, but add some checks against nil values and
// auto-dereference interfaces and pointers.
//
// The difference between Smuggle and Code operators is that Code is
// used to do a final comparison while Smuggle transforms the data and
// then steps down in favor of generic comparison process. Moreover,
// the type accepted as input for the function is more lax to
// facilitate the tests writing (e.g. the function can accept a float64
// and the got value be an int). See examples. On the other hand, the
// output type is strict and must match exactly the expected value
// type. The fields-path string "fn" shortcut is not available with
// Code operator.
//
// TypeBehind method returns the reflect.Type of only parameter of
// "fn". For the case where "fn" is a fields-path, it is always
// interface{}, as the type can not be known in advance.
func Smuggle(fn interface{}, expectedValue interface{}) TestDeep {
	vfn := reflect.ValueOf(fn)

	const usage = "Smuggle(FUNC|FIELDS_PATH, TESTDEEP_OPERATOR|EXPECTED_VALUE)"

	switch vfn.Kind() {
	case reflect.String:
		fn, err := buildStructFieldFn(vfn.String())
		if err != nil {
			panic(usage + ": " + err.Error())
		}
		vfn = reflect.ValueOf(fn)

	case reflect.Func:
		// nothing to do check

	default:
		panic("usage: " + usage)
	}

	fnType := vfn.Type()
	if fnType.IsVariadic() || fnType.NumIn() != 1 {
		panic(usage + ": FUNC must take only one argument")
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
				fnType.Out(1) != errorInterface) {
			break
		}
		fallthrough

	case 1: // (value)
		s := tdSmuggle{
			tdSmugglerBase: newSmugglerBase(expectedValue),
			function:       vfn,
			argType:        fnType.In(0),
		}
		if !s.isTestDeeper {
			s.expectedValue = reflect.ValueOf(expectedValue)
		}
		return &s
	}

	panic(usage +
		": FUNC must return value or (value, bool) or (value, bool, string) or (value, error)")
}

func (s *tdSmuggle) laxConvert(got reflect.Value) (reflect.Value, bool) {
	if got.IsValid() && got.Type().ConvertibleTo(s.argType) {
		return got.Convert(s.argType), true
	}
	return got, false
}

func (s *tdSmuggle) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
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
	return "Smuggle(" + s.function.Type().String() + ")"
}

func (s *tdSmuggle) TypeBehind() reflect.Type {
	return s.argType
}
