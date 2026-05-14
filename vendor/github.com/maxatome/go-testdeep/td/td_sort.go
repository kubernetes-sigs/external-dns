// Copyright (c) 2024-2025, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/maxatome/go-testdeep/internal/compare"
	"github.com/maxatome/go-testdeep/internal/ctxerr"
	"github.com/maxatome/go-testdeep/internal/dark"
	"github.com/maxatome/go-testdeep/internal/types"
	"github.com/maxatome/go-testdeep/internal/util"
	"github.com/maxatome/go-testdeep/internal/visited"
)

type tdSortBase struct {
	mkSortFn func(reflect.Type) (reflect.Value, error)
}

func (sb *tdSortBase) initSortBase(how ...any) error {
	switch l := len(how); l {
	case 0:
		how = []any{1}
	case 1:
	default: // list of fields-paths used by Sorted only
		fieldsPaths := make([]string, l)
		for i, si := range how {
			s, ok := si.(string)
			if !ok {
				return errors.New("string... expected " + util.BadParam(si, i+1, true))
			}
			fieldsPaths[i] = s
		}
		how = []any{fieldsPaths}
	}

	switch v := how[0].(type) {
	case nil:
		sb.mkSortFn = mkSortAsc
	case int:
		sb.mkSortFn = mkSortAscDesc(v >= 0)
	case float64: // to be used in JSON, SubJSONOf & SuperJSONOf
		sb.mkSortFn = mkSortAscDesc(v >= 0)
	case string: // one fields-path
		sb.mkSortFn = func(typ reflect.Type) (reflect.Value, error) {
			return mkSortFieldsPaths(typ, []string{v})
		}
	case []any: // fields-paths list in JSON context
		ss := make([]string, len(v))
		var ok bool
		for i, s := range v {
			ss[i], ok = s.(string)
			if !ok {
				return fmt.Errorf(
					"slice of strings expected as how, %T encountered at pos %d", s, i)
			}
		}
		sb.mkSortFn = func(typ reflect.Type) (reflect.Value, error) {
			return mkSortFieldsPaths(typ, ss)
		}
	case []string: // fields-paths list
		sb.mkSortFn = func(typ reflect.Type) (reflect.Value, error) {
			return mkSortFieldsPaths(typ, v)
		}
	default:
		vv := reflect.ValueOf(v)
		if vv.Kind() != reflect.Func {
			return errors.New(util.BadParam(v, 1, true))
		}
		ft := vv.Type()
		if ft.IsVariadic() || ft.NumIn() != 2 || ft.In(0) != ft.In(1) ||
			ft.NumOut() != 1 || ft.Out(0) != types.Bool {
			return fmt.Errorf("SORT_FUNC must match func(T, T) bool signature, not %T", v)
		}
		sb.mkSortFn = func(typ reflect.Type) (reflect.Value, error) {
			if !typ.AssignableTo(ft.In(0)) {
				return reflect.Value{}, fmt.Errorf("%s is not assignable to %s", typ, ft.In(0))
			}
			return vv, nil
		}
	}
	return nil
}

func mkSortAscDesc(asc bool) func(reflect.Type) (reflect.Value, error) {
	if asc {
		return mkSortAsc
	}
	return mkSortDesc
}

func mkSortAsc(typ reflect.Type) (reflect.Value, error) {
	v := visited.NewVisited()
	return reflect.MakeFunc(
		reflect.FuncOf([]reflect.Type{typ, typ}, []reflect.Type{types.Bool}, false),
		func(args []reflect.Value) []reflect.Value {
			less := compare.Compare(v, args[0], args[1]) < 0
			return []reflect.Value{reflect.ValueOf(less)}
		}), nil
}

func mkSortDesc(typ reflect.Type) (reflect.Value, error) {
	v := visited.NewVisited()
	return reflect.MakeFunc(
		reflect.FuncOf([]reflect.Type{typ, typ}, []reflect.Type{types.Bool}, false),
		func(args []reflect.Value) []reflect.Value {
			less := compare.Compare(v, args[1], args[0]) < 0
			return []reflect.Value{reflect.ValueOf(less)}
		}), nil
}

func mkSortFieldsPaths(typ reflect.Type, fieldsPaths []string) (reflect.Value, error) {
	type sortFP struct {
		fn  func(any) (smuggleValue, error)
		asc bool
	}
	fns := make([]sortFP, len(fieldsPaths))
	for i, fp := range fieldsPaths {
		var sfp sortFP
		if strings.HasPrefix(fp, "-") {
			fp = fp[1:]
		} else {
			sfp.asc = true
			fp = strings.TrimPrefix(fp, "+") // optional
		}
		fn, err := getFieldsPathFn(fp)
		if err != nil {
			return reflect.Value{}, err
		}
		sfp.fn = fn.Interface().(func(any) (smuggleValue, error))
		fns[i] = sfp
	}

	v := visited.NewVisited()
	return reflect.MakeFunc(
		reflect.FuncOf([]reflect.Type{typ, typ}, []reflect.Type{types.Bool}, false),
		func(args []reflect.Value) []reflect.Value {
			a, aOK := dark.GetInterface(args[0], true)
			b, bOK := dark.GetInterface(args[1], true)
			if aOK && bOK {
				for _, fn := range fns {
					va, aErr := fn.fn(a)
					vb, bErr := fn.fn(b)
					if aErr != nil || bErr != nil {
						if aErr == nil || bErr == nil {
							// nonexistent field is greater
							return []reflect.Value{reflect.ValueOf(aErr == nil)}
						}
						break // both nonexistent fields, use Compare
					}
					cmp := compare.Compare(v, va.Value, vb.Value)
					if cmp == 0 {
						continue
					}
					return []reflect.Value{reflect.ValueOf(cmp < 0 == fn.asc)}
				}
			}
			less := compare.Compare(v, args[0], args[1]) < 0
			return []reflect.Value{reflect.ValueOf(less)}
		}), nil
}

const sortUsage = "(SORT_FUNC|int|string|[]string, TESTDEEP_OPERATOR|EXPECTED_VALUE)"

type tdSort struct {
	tdSmugglerBase
	tdSortBase
	how any
}

var _ TestDeep = &tdSort{}

// summary(Sort): sorts a slice or an array before comparing its content
// input(Sort): array,slice,ptr(ptr on array/slice)

// Sort is a smuggler operator. It takes an array, a slice or a
// pointer on array/slice, it sorts it using how and compares the
// sorted result to expectedValue. It can be seen as an alternative to
// [Bag].
//
// how can be:
//   - nil or a float64/int >= 0 for a generic ascending order;
//   - a float64/int < 0 for a generic descending order;
//   - a string specifying a fields-path (optionally prefixed by "+"
//     or "-" for respectively an ascending or a descending order,
//     defaulting to ascending one);
//   - a []string containing a list of fields-paths (as above), second
//     and next fields-paths are checked when the previous ones are equal;
//   - a function matching func(a, b T) bool signature and returning
//     true if a is before b.
//
// A fields-path, also used by [Smuggle] and [Sorted] operators,
// allows to access nested structs fields and maps & slices items. See
// [Smuggle] for details on fields-path possibilities.
//
//	type A struct{ props map[string]int }
//	p12 := A{props: map[string]int{"priority": 12}}
//	p23 := A{props: map[string]int{"priority": 23}}
//	p34 := A{props: map[string]int{"priority": 34}}
//	got := []A{p23, p12, p34}
//	td.Cmp(t, got, td.Sort("-props[priority]", []A{p34, p23, p12})) // succeeds
//
// how can be a float64 to allow Sort to be used in expected JSON of
// [JSON], [SubJSONOf] & [SuperJSONOf] operators:
//
//	got := map[string][]string{"labels": {"c", "a", "b"}}
//	td.Cmp(t, got, td.JSON(`{ "labels": Sort(1, ["a", "b", "c"]) }`)) // succeeds
//
// or using fields-path feature:
//
//	type Person struct {
//	  Name string `json:"name"`
//	  Age  int    `json:"age"`
//	}
//	got := struct {
//	  People []Person `json:"people"`
//	}{
//	  People: []Person{
//	    {"Brian", 22},
//	    {"Bob", 19},
//	    {"Stephen", 19},
//	    {"Alice", 20},
//	    {"Marcel", 25},
//	  },
//	}
//	td.Cmp(t, got, td.JSON(`{
//	  "people": Sort("name", [ // sort by name ascending
//	    {"name": "Alice",   "age": 20},
//	    {"name": "Bob",     "age": 19},
//	    {"name": "Brian",   "age": 22},
//	    {"name": "Marcel",  "age": 25},
//	    {"name": "Stephen", "age": 19},
//	  ])
//	}`)) // succeeds
//	td.Cmp(t, got, td.JSON(`{
//	  "people": Sort([ "-age", "name" ], [ // sort by age desc, then by name asc
//	    {"name": "Marcel",  "age": 25},
//	    {"name": "Brian",   "age": 22},
//	    {"name": "Alice",   "age": 20},
//	    {"name": "Bob",     "age": 19},
//	    {"name": "Stephen", "age": 19},
//	  ])
//	}`)) // succeeds
//
// See also [Sorted], [Smuggle] and [Bag].
func Sort(how any, expectedValue any) TestDeep {
	s := tdSort{how: how}
	s.tdSmugglerBase = newSmugglerBase(expectedValue, 0)
	if !s.isTestDeeper {
		s.expectedValue = reflect.ValueOf(expectedValue)
	}

	err := s.initSortBase(how)
	if err != nil {
		s.err = ctxerr.OpBad("Sort", "usage: Sort%s, %s", sortUsage, err)
	} else if !s.isTestDeeper {
		switch s.expectedValue.Kind() {
		case reflect.Slice, reflect.Array:
		default:
			s.err = ctxerr.OpBad("Sort",
				"usage: Sort%s, EXPECTED_VALUE must be a slice or an array not a %s",
				sortUsage, types.KindType(s.expectedValue))
		}
	}
	return &s
}

func (s *tdSort) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	if s.err != nil {
		return ctx.CollectError(s.err)
	}

	if rErr := grepResolvePtr(ctx, &got); rErr != nil {
		return ctx.CollectError(rErr)
	}

	switch got.Kind() {
	case reflect.Slice, reflect.Array:
	default:
		return grepBadKind(ctx, got)
	}

	const sorted = "<sorted>"

	itemType := got.Type().Elem()
	fn, err := s.mkSortFn(itemType)
	if err != nil {
		if ctx.BooleanError {
			return ctxerr.BooleanError
		}
		return ctx.CollectError(&ctxerr.Error{
			Message: "cannot sort items",
			Summary: ctxerr.NewSummary(err.Error()),
		})
	}

	l := got.Len()
	if l <= 1 {
		return deepValueEqual(ctx.AddCustomLevel(sorted), got, s.expectedValue)
	}

	var out reflect.Value
	if got.Kind() == reflect.Slice {
		out = reflect.MakeSlice(reflect.SliceOf(itemType), l, l)
	} else {
		out = reflect.New(got.Type()).Elem()
	}
	reflect.Copy(out, got)

	sort.SliceStable(out.Slice(0, out.Len()).Interface(), func(i, j int) bool {
		return fn.Call([]reflect.Value{out.Index(i), out.Index(j)})[0].Bool()
	})
	return deepValueEqual(ctx.AddCustomLevel(sorted), out, s.expectedValue)
}

func (s *tdSort) String() string {
	if s.err != nil {
		return s.stringError()
	}
	how, typ := s.how, reflect.TypeOf(s.how)
	if typ != nil && typ.Kind() == reflect.Func {
		how = typ.String()
	}
	return S("Sort(%v, %s)", how, util.ToString(s.expectedValue))
}

type tdSorted struct {
	baseOKNil
	tdSortBase
	how []any
}

var _ TestDeep = &tdSorted{}

const sortedUsage = "(SORT_FUNC|int|[]string|string...)"

// summary(Sorted): checks a slice or an array is sorted
// input(Sorted): array,slice,ptr(ptr on array/slice)

// Sorted operator checks that data is an array, a slice or a pointer
// on array/slice, and it is well sorted as how tells it should be.
//
// how... can be:
//   - empty to check a generic ascending order;
//   - nil or a float64/int >= 0 to check a generic ascending order;
//   - a float64/int < 0 to check a generic descending order;
//   - strings specifying fields-paths (each optionally prefixed by "+"
//     or "-" for respectively checking an ascending or a descending order,
//     defaulting to ascending one);
//   - a function matching func(a, b T) bool signature and returning
//     true if a is before b.
//
// A fields-path, also used by [Smuggle] and [Sort] operators,
// allows to access nested structs fields and maps & slices items. See
// [Smuggle] for details on fields-path possibilities.
//
//	type A struct{ props map[string]int }
//	p12 := A{props: map[string]int{"priority": 12}}
//	p23 := A{props: map[string]int{"priority": 23}}
//	p34 := A{props: map[string]int{"priority": 34}}
//	got := []A{p34, p23, p12}
//	td.Cmp(t, got, td.Sorted("-props[priority]")) // succeeds
//
// how can be a float64 to allow Sort to be used in expected JSON of
// [JSON], [SubJSONOf] & [SuperJSONOf] operators:
//
//	got := map[string][]string{"labels": {"a", "b", "c"}}
//	td.Cmp(t, got, td.JSON(`{ "labels": Sorted }`)) // succeeds
//
// or using fields-path feature:
//
//	type Person struct {
//	  Name string `json:"name"`
//	  Age  int    `json:"age"`
//	}
//	got := struct {
//	  People []Person `json:"people"`
//	}{
//	  People: []Person{
//	    {"Marcel", 25},
//	    {"Brian", 22},
//	    {"Alice", 20},
//	    {"Bob", 19},
//	    {"Stephen", 19},
//	  },
//	}
//	// sorted by age desc, then by name asc
//	td.Cmp(t, got, td.JSON(`{ "people": Sorted("-age", "name") }`)) // succeeds
//
// See also [Sort].
func Sorted(how ...any) TestDeep {
	s := tdSorted{
		baseOKNil: newBaseOKNil(3),
		how:       how,
	}

	err := s.initSortBase(how...)
	if err != nil {
		s.err = ctxerr.OpBad("Sorted", "usage: Sorted%s, %s", sortedUsage, err)
	}
	return &s
}

func (s *tdSorted) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	if s.err != nil {
		return ctx.CollectError(s.err)
	}

	if rErr := grepResolvePtr(ctx, &got); rErr != nil {
		return ctx.CollectError(rErr)
	}

	switch got.Kind() {
	case reflect.Slice, reflect.Array:
	default:
		return grepBadKind(ctx, got)
	}

	itemType := got.Type().Elem()
	fn, err := s.mkSortFn(itemType)
	if err != nil {
		if ctx.BooleanError {
			return ctxerr.BooleanError
		}
		return ctx.CollectError(&ctxerr.Error{
			Message: "cannot sort items",
			Summary: ctxerr.NewSummary(err.Error()),
		})
	}

	for i, l := 1, got.Len(); i < l; i++ {
		if fn.Call([]reflect.Value{got.Index(i), got.Index(i - 1)})[0].Bool() {
			return ctx.CollectError(&ctxerr.Error{
				Message: fmt.Sprintf("not sorted, item #%d value is before #%d one while it should not", i, i-1),
				Summary: ctxerr.ErrorSummaryItems{
					{
						Label: fmt.Sprintf("item #%d", i-1),
						Value: util.ToString(got.Index(i - 1)),
					},
					{
						Label: fmt.Sprintf("item #%d", i),
						Value: util.ToString(got.Index(i)),
					},
				},
			})
		}
	}

	return nil
}

func (s *tdSorted) String() string {
	if s.err != nil {
		return s.stringError()
	}
	var b strings.Builder
	b.WriteString("Sorted(")
	for i, cur := range s.how {
		how, typ := cur, reflect.TypeOf(cur)
		if typ != nil && typ.Kind() == reflect.Func {
			how = typ.String()
		}
		if i > 0 {
			b.WriteString(", ")
		}
		fmt.Fprintf(&b, "%v", how)
	}
	b.WriteByte(')')
	return b.String()
}
