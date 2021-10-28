// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

import (
	"bytes"
	"reflect"
	"strings"

	"github.com/maxatome/go-testdeep/helpers/tdutil"
	"github.com/maxatome/go-testdeep/internal/ctxerr"
	"github.com/maxatome/go-testdeep/internal/types"
	"github.com/maxatome/go-testdeep/internal/util"
)

type tdContains struct {
	tdSmugglerBase
}

var _ TestDeep = &tdContains{}

// summary(Contains): checks that a string, []byte, error or
// fmt.Stringer interfaces contain a rune, byte or a sub-string; or a
// slice contains a single value or a sub-slice; or an array or map
// contain a single value
// input(Contains): str,array,slice,map,if(✓ + fmt.Stringer/error)

// Contains is a smuggler operator to check if something is contained
// in another thing. Contains has to be applied on arrays, slices, maps or
// strings. It tries to be as smarter as possible.
//
// If "expectedValue" is a TestDeep operator, each item of data
// array/slice/map/string (rune for strings) is compared to it. The
// use of a TestDeep operator as "expectedValue" works only in this
// way: item per item.
//
// If data is a slice, and "expectedValue" has the same type, then
// "expectedValue" is searched as a sub-slice, otherwise
// "expectedValue" is compared to each slice value.
//
//   list := []int{12, 34, 28}
//   td.Cmp(t, list, td.Contains(34))                 // succeeds
//   td.Cmp(t, list, td.Contains(td.Between(30, 35))) // succeeds too
//   td.Cmp(t, list, td.Contains(35))                 // fails
//   td.Cmp(t, list, td.Contains([]int{34, 28}))      // succeeds
//
// If data is an array or a map, each value is compared to
// "expectedValue". Map keys are not checked: see ContainsKey to check
// map keys existence.
//
//   hash := map[string]int{"foo": 12, "bar": 34, "zip": 28}
//   td.Cmp(t, hash, td.Contains(34))                 // succeeds
//   td.Cmp(t, hash, td.Contains(td.Between(30, 35))) // succeeds too
//   td.Cmp(t, hash, td.Contains(35))                 // fails
//
//   array := [...]int{12, 34, 28}
//   td.Cmp(t, array, td.Contains(34))                 // succeeds
//   td.Cmp(t, array, td.Contains(td.Between(30, 35))) // succeeds too
//   td.Cmp(t, array, td.Contains(35))                 // fails
//
// If data is a string (or convertible), []byte (or convertible),
// error or fmt.Stringer interface (error interface is tested before
// fmt.Stringer), "expectedValue" can be a string, a []byte, a rune or
// a byte. In this case, it tests if the got string contains this
// expected string, []byte, rune or byte.
//
//   got := "foo bar"
//   td.Cmp(t, got, td.Contains('o'))                  // succeeds
//   td.Cmp(t, got, td.Contains(rune('o')))            // succeeds
//   td.Cmp(t, got, td.Contains(td.Between('n', 'p'))) // succeeds
//   td.Cmp(t, got, td.Contains("bar"))                // succeeds
//   td.Cmp(t, got, td.Contains([]byte("bar")))        // succeeds
//
//   td.Cmp(t, []byte("foobar"), td.Contains("ooba")) // succeeds
//
//   type Foobar string
//   td.Cmp(t, Foobar("foobar"), td.Contains("ooba")) // succeeds
//
//   err := errors.New("error!")
//   td.Cmp(t, err, td.Contains("ror")) // succeeds
//
//   bstr := bytes.NewBufferString("fmt.Stringer!")
//   td.Cmp(t, bstr, td.Contains("String")) // succeeds
//
// Pitfall: if you want to check if 2 words are contained in got, don't do:
//
//   td.Cmp(t, "foobar", td.Contains(td.All("foo", "bar"))) // Bad!
//
// as TestDeep operator All in Contains operates on each rune, so it
// does not work as expected, but do::
//
//   td.Cmp(t, "foobar", td.All(td.Contains("foo"), td.Contains("bar")))
//
// When Contains(nil) is used, nil is automatically converted to a
// typed nil on the fly to avoid confusion (if the array/slice/map
// item type allows it of course.) So all following Cmp calls
// are equivalent (except the (*byte)(nil) one):
//
//   num := 123
//   list := []*int{&num, nil}
//   td.Cmp(t, list, td.Contains(nil))         // succeeds → (*int)(nil)
//   td.Cmp(t, list, td.Contains((*int)(nil))) // succeeds
//   td.Cmp(t, list, td.Contains(td.Nil()))    // succeeds
//   // But...
//   td.Cmp(t, list, td.Contains((*byte)(nil))) // fails: (*byte)(nil) ≠ (*int)(nil)
//
// As well as these ones:
//
//   hash := map[string]*int{"foo": nil, "bar": &num}
//   td.Cmp(t, hash, td.Contains(nil))         // succeeds → (*int)(nil)
//   td.Cmp(t, hash, td.Contains((*int)(nil))) // succeeds
//   td.Cmp(t, hash, td.Contains(td.Nil()))    // succeeds
func Contains(expectedValue interface{}) TestDeep {
	c := tdContains{
		tdSmugglerBase: newSmugglerBase(expectedValue),
	}

	if !c.isTestDeeper {
		c.expectedValue = reflect.ValueOf(expectedValue)
	}
	return &c
}

func (c *tdContains) doesNotContainErr(ctx ctxerr.Context, got interface{}) *ctxerr.Error {
	if ctx.BooleanError {
		return ctxerr.BooleanError
	}
	return ctx.CollectError(&ctxerr.Error{
		Message:  "does not contain",
		Got:      got,
		Expected: c,
	})
}

// getExpectedValue returns the expected value handling the
// Contains(nil) case: in this case it returns a typed nil (same type
// as the items of got).
// got is an array, a slice or a map (it's the caller responsibility to check).
func (c *tdContains) getExpectedValue(got reflect.Value) reflect.Value {
	// If the expectValue is non-typed nil
	if !c.expectedValue.IsValid() {
		// AND the kind of items in got is...
		switch got.Type().Elem().Kind() {
		case reflect.Chan, reflect.Func, reflect.Interface,
			reflect.Map, reflect.Ptr, reflect.Slice:
			// returns a typed nil
			return reflect.Zero(got.Type().Elem())
		}
	}
	return c.expectedValue
}

func (c *tdContains) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	switch got.Kind() {
	case reflect.Slice:
		if !c.isTestDeeper && c.expectedValue.IsValid() {
			// Special case for []byte & expected []byte or string
			if got.Type().Elem() == types.Uint8 {
				switch c.expectedValue.Kind() {
				case reflect.String:
					if bytes.Contains(got.Bytes(), []byte(c.expectedValue.String())) {
						return nil
					}
					return c.doesNotContainErr(ctx, got)

				case reflect.Slice:
					if c.expectedValue.Type().Elem() == types.Uint8 {
						if bytes.Contains(got.Bytes(), c.expectedValue.Bytes()) {
							return nil
						}
						return c.doesNotContainErr(ctx, got)
					}

				case reflect.Int32: // rune
					if bytes.ContainsRune(got.Bytes(), rune(c.expectedValue.Int())) {
						return nil
					}
					return c.doesNotContainErr(ctx, got)

				case reflect.Uint8: // byte
					if bytes.ContainsRune(got.Bytes(), rune(c.expectedValue.Uint())) {
						return nil
					}
					return c.doesNotContainErr(ctx, got)
				}

				// fall back on string conversion
				break
			}

			// Search slice in slice
			if got.Type() == c.expectedValue.Type() {
				gotLen, expectedLen := got.Len(), c.expectedValue.Len()
				if expectedLen == 0 {
					return nil
				}
				if expectedLen > gotLen {
					return c.doesNotContainErr(ctx, got)
				}
				if expectedLen == gotLen {
					if deepValueEqualOK(got, c.expectedValue) {
						return nil
					}
					return c.doesNotContainErr(ctx, got)
				}

				for i := 0; i < gotLen-expectedLen; i++ {
					if deepValueEqualOK(got.Slice(i, i+expectedLen), c.expectedValue) {
						return nil
					}
				}
			}
		}
		fallthrough

	case reflect.Array:
		expectedValue := c.getExpectedValue(got)
		for index := got.Len() - 1; index >= 0; index-- {
			if deepValueEqualFinalOK(ctx, got.Index(index), expectedValue) {
				return nil
			}
		}
		return c.doesNotContainErr(ctx, got)

	case reflect.Map:
		expectedValue := c.getExpectedValue(got)
		if !tdutil.MapEachValue(got, func(v reflect.Value) bool {
			return !deepValueEqualFinalOK(ctx, v, expectedValue)
		}) {
			return nil
		}
		return c.doesNotContainErr(ctx, got)
	}

	str, err := getString(ctx, got)
	if err != nil {
		return err
	}

	// If a TestDeep operator is expected, applies this operator on
	// each character of the string
	if c.isTestDeeper {
		// If the type behind the operator is known *and* is not rune,
		// then no need to go further, but return an explicit error to
		// help our user to fix his probably bogus code
		op := c.expectedValue.Interface().(TestDeep)
		if typeBehind := op.TypeBehind(); typeBehind != nil && typeBehind != types.Rune && !ctx.BeLax {
			if ctx.BooleanError {
				return ctxerr.BooleanError
			}
			return ctx.CollectError(&ctxerr.Error{
				Message:  op.GetLocation().Func + " operator has to match rune in string, but it does not",
				Got:      types.RawString(typeBehind.String()),
				Expected: types.RawString("rune"),
			})
		}

		for _, chr := range str {
			if deepValueEqualFinalOK(ctx, reflect.ValueOf(chr), c.expectedValue) {
				return nil
			}
		}
		return c.doesNotContainErr(ctx, got)
	}

	// If expectedValue is a []byte, a string, a rune or a byte, we
	// check whether it is contained in the string or not
	var contains bool
	switch expectedKind := c.expectedValue.Kind(); expectedKind {
	case reflect.String:
		contains = strings.Contains(str, c.expectedValue.String())

	case reflect.Int32: // rune
		contains = strings.ContainsRune(str, rune(c.expectedValue.Int()))

	case reflect.Uint8: // byte
		contains = strings.ContainsRune(str, rune(c.expectedValue.Uint()))

	case reflect.Slice:
		// Only []byte
		if c.expectedValue.Type().Elem() == types.Uint8 {
			contains = strings.Contains(str, string(c.expectedValue.Bytes()))
			break
		}
		fallthrough

	default:
		if ctx.BooleanError {
			return ctxerr.BooleanError
		}
		var expectedType interface{}
		if c.expectedValue.IsValid() {
			expectedType = types.RawString(c.expectedValue.Type().String())
		} else {
			expectedType = c
		}

		return ctx.CollectError(&ctxerr.Error{
			Message:  "cannot check contains",
			Got:      types.RawString(got.Type().String()),
			Expected: expectedType,
		})
	}

	if contains {
		return nil
	}
	return c.doesNotContainErr(ctx, str)
}

func (c *tdContains) String() string {
	return "Contains(" + util.ToString(c.expectedValue) + ")"
}
