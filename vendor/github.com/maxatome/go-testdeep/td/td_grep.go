// Copyright (c) 2022-2023, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

import (
	"reflect"

	"github.com/maxatome/go-testdeep/internal/ctxerr"
	"github.com/maxatome/go-testdeep/internal/types"
)

const grepUsage = "(FILTER_FUNC|FILTER_TESTDEEP_OPERATOR, TESTDEEP_OPERATOR|EXPECTED_VALUE)"

type tdGrepBase struct {
	tdSmugglerBase
	filter  reflect.Value // func (argType ≠ nil) OR TestDeep operator
	argType reflect.Type
}

func (g *tdGrepBase) initGrepBase(filter, expectedValue any) {
	g.tdSmugglerBase = newSmugglerBase(expectedValue, 1)

	if !g.isTestDeeper {
		g.expectedValue = reflect.ValueOf(expectedValue)
	}

	if op, ok := filter.(TestDeep); ok {
		g.filter = reflect.ValueOf(op)
		return
	}

	vfilter := reflect.ValueOf(filter)
	if vfilter.Kind() != reflect.Func {
		g.err = ctxerr.OpBad(g.GetLocation().Func,
			"usage: %s%s, FILTER_FUNC must be a function or FILTER_TESTDEEP_OPERATOR a TestDeep operator",
			g.GetLocation().Func, grepUsage)
		return
	}

	filterType := vfilter.Type()
	if filterType.IsVariadic() || filterType.NumIn() != 1 {
		g.err = ctxerr.OpBad(g.GetLocation().Func,
			"usage: %s%s, FILTER_FUNC must take only one non-variadic argument",
			g.GetLocation().Func, grepUsage)
		return
	}
	if filterType.NumOut() != 1 || filterType.Out(0) != types.Bool {
		g.err = ctxerr.OpBad(g.GetLocation().Func,
			"usage: %s%s, FILTER_FUNC must return bool",
			g.GetLocation().Func, grepUsage)
		return
	}

	g.argType = filterType.In(0)
	g.filter = vfilter
}

func (g *tdGrepBase) matchItem(ctx ctxerr.Context, idx int, item reflect.Value) (bool, *ctxerr.Error) {
	if g.argType == nil {
		// g.filter is a TestDeep operator
		return deepValueEqualFinalOK(ctx, item, g.filter), nil
	}

	// item is an interface, but the filter function does not expect an
	// interface, resolve it
	if item.Kind() == reflect.Interface && g.argType.Kind() != reflect.Interface {
		item = item.Elem()
	}

	if !item.Type().AssignableTo(g.argType) {
		if !types.IsConvertible(item, g.argType) {
			if ctx.BooleanError {
				return false, ctxerr.BooleanError
			}
			return false, ctx.AddArrayIndex(idx).CollectError(&ctxerr.Error{
				Message:  "incompatible parameter type",
				Got:      types.RawString(item.Type().String()),
				Expected: types.RawString(g.argType.String()),
			})
		}
		item = item.Convert(g.argType)
	}

	return g.filter.Call([]reflect.Value{item})[0].Bool(), nil
}

func (g *tdGrepBase) HandleInvalid() bool {
	return true // Knows how to handle untyped nil values (aka invalid values)
}

func (g *tdGrepBase) String() string {
	if g.err != nil {
		return g.stringError()
	}
	if g.argType == nil {
		return S("%s(%s)", g.GetLocation().Func, g.filter.Interface().(TestDeep))
	}
	return S("%s(%s)", g.GetLocation().Func, g.filter.Type())
}

func (g *tdGrepBase) TypeBehind() reflect.Type {
	if g.err != nil {
		return nil
	}
	return g.internalTypeBehind()
}

// sliceTypeBehind is used by First & Last TypeBehind method.
func (g *tdGrepBase) sliceTypeBehind() reflect.Type {
	typ := g.TypeBehind()
	if typ == nil {
		return nil
	}
	return reflect.SliceOf(typ)
}

func (g *tdGrepBase) notFound(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	if ctx.BooleanError {
		return ctxerr.BooleanError
	}
	return ctx.CollectError(&ctxerr.Error{
		Message:  "item not found",
		Got:      got,
		Expected: types.RawString(g.String()),
	})
}

func grepResolvePtr(ctx ctxerr.Context, got *reflect.Value) *ctxerr.Error {
	if got.Kind() == reflect.Ptr {
		gotElem := got.Elem()
		if !gotElem.IsValid() {
			if ctx.BooleanError {
				return ctxerr.BooleanError
			}
			return ctx.CollectError(ctxerr.NilPointer(*got, "non-nil *slice OR *array"))
		}
		switch gotElem.Kind() {
		case reflect.Slice, reflect.Array:
			*got = gotElem
		}
	}
	return nil
}

func grepBadKind(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	if ctx.BooleanError {
		return ctxerr.BooleanError
	}
	return ctx.CollectError(ctxerr.BadKind(got, "slice OR array OR *slice OR *array"))
}

type tdGrep struct {
	tdGrepBase
}

var _ TestDeep = &tdGrep{}

// summary(Grep): reduces a slice or an array before comparing its content
// input(Grep): array,slice,ptr(ptr on array/slice)

// Grep is a smuggler operator. It takes an array, a slice or a
// pointer on array/slice. For each item it applies filter, a
// [TestDeep] operator or a function returning a bool, and produces a
// slice consisting of those items for which the filter matched and
// compares it to expectedValue. The filter matches when it is a:
//   - [TestDeep] operator and it matches for the item;
//   - function receiving the item and it returns true.
//
// expectedValue can be a [TestDeep] operator or a slice (but never an
// array nor a pointer on a slice/array nor any other kind).
//
//	got := []int{-3, -2, -1, 0, 1, 2, 3}
//	td.Cmp(t, got, td.Grep(td.Gt(0), []int{1, 2, 3})) // succeeds
//	td.Cmp(t, got, td.Grep(
//	  func(x int) bool { return x%2 == 0 },
//	  []int{-2, 0, 2})) // succeeds
//	td.Cmp(t, got, td.Grep(
//	  func(x int) bool { return x%2 == 0 },
//	  td.Set(0, 2, -2))) // succeeds
//
// If Grep receives a nil slice or a pointer on a nil slice, it always
// returns a nil slice:
//
//	var got []int
//	td.Cmp(t, got, td.Grep(td.Gt(0), ([]int)(nil))) // succeeds
//	td.Cmp(t, got, td.Grep(td.Gt(0), td.Nil()))     // succeeds
//	td.Cmp(t, got, td.Grep(td.Gt(0), []int{}))      // fails
//
// See also [First], [Last] and [Flatten].
func Grep(filter, expectedValue any) TestDeep {
	g := tdGrep{}
	g.initGrepBase(filter, expectedValue)

	if g.err == nil && !g.isTestDeeper && g.expectedValue.Kind() != reflect.Slice {
		g.err = ctxerr.OpBad("Grep",
			"usage: Grep%s, EXPECTED_VALUE must be a slice not a %s",
			grepUsage, types.KindType(g.expectedValue))
	}
	return &g
}

func (g *tdGrep) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	if g.err != nil {
		return ctx.CollectError(g.err)
	}

	if rErr := grepResolvePtr(ctx, &got); rErr != nil {
		return rErr
	}

	switch got.Kind() {
	case reflect.Slice, reflect.Array:
		const grepped = "<grepped>"

		if got.Kind() == reflect.Slice && got.IsNil() {
			return deepValueEqual(
				ctx.AddCustomLevel(grepped),
				reflect.New(got.Type()).Elem(),
				g.expectedValue,
			)
		}

		l := got.Len()
		out := reflect.MakeSlice(reflect.SliceOf(got.Type().Elem()), 0, l)

		for idx := 0; idx < l; idx++ {
			item := got.Index(idx)
			ok, rErr := g.matchItem(ctx, idx, item)
			if rErr != nil {
				return rErr
			}
			if ok {
				out = reflect.Append(out, item)
			}
		}

		return deepValueEqual(ctx.AddCustomLevel(grepped), out, g.expectedValue)
	}

	return grepBadKind(ctx, got)
}

type tdFirst struct {
	tdGrepBase
}

var _ TestDeep = &tdFirst{}

// summary(First): find the first matching item of a slice or an array
// then compare its content
// input(First): array,slice,ptr(ptr on array/slice)

// First is a smuggler operator. It takes an array, a slice or a
// pointer on array/slice. For each item it applies filter, a
// [TestDeep] operator or a function returning a bool. It takes the
// first item for which the filter matched and compares it to
// expectedValue. The filter matches when it is a:
//   - [TestDeep] operator and it matches for the item;
//   - function receiving the item and it returns true.
//
// expectedValue can of course be a [TestDeep] operator.
//
//	got := []int{-3, -2, -1, 0, 1, 2, 3}
//	td.Cmp(t, got, td.First(td.Gt(0), 1))                                    // succeeds
//	td.Cmp(t, got, td.First(func(x int) bool { return x%2 == 0 }, -2))       // succeeds
//	td.Cmp(t, got, td.First(func(x int) bool { return x%2 == 0 }, td.Lt(0))) // succeeds
//
// If the input is empty (and/or nil for a slice), an "item not found"
// error is raised before comparing to expectedValue.
//
//	var got []int
//	td.Cmp(t, got, td.First(td.Gt(0), td.Gt(0)))      // fails
//	td.Cmp(t, []int{}, td.First(td.Gt(0), td.Gt(0)))  // fails
//	td.Cmp(t, [0]int{}, td.First(td.Gt(0), td.Gt(0))) // fails
//
// See also [Last] and [Grep].
func First(filter, expectedValue any) TestDeep {
	g := tdFirst{}
	g.initGrepBase(filter, expectedValue)
	return &g
}

func (g *tdFirst) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	if g.err != nil {
		return ctx.CollectError(g.err)
	}

	if rErr := grepResolvePtr(ctx, &got); rErr != nil {
		return rErr
	}

	switch got.Kind() {
	case reflect.Slice, reflect.Array:
		for idx, l := 0, got.Len(); idx < l; idx++ {
			item := got.Index(idx)
			ok, rErr := g.matchItem(ctx, idx, item)
			if rErr != nil {
				return rErr
			}
			if ok {
				return deepValueEqual(
					ctx.AddCustomLevel(S("<first#%d>", idx)),
					item,
					g.expectedValue,
				)
			}
		}
		return g.notFound(ctx, got)
	}

	return grepBadKind(ctx, got)
}

func (g *tdFirst) TypeBehind() reflect.Type {
	return g.sliceTypeBehind()
}

type tdLast struct {
	tdGrepBase
}

var _ TestDeep = &tdLast{}

// summary(Last): find the last matching item of a slice or an array
// then compare its content
// input(Last): array,slice,ptr(ptr on array/slice)

// Last is a smuggler operator. It takes an array, a slice or a
// pointer on array/slice. For each item it applies filter, a
// [TestDeep] operator or a function returning a bool. It takes the
// last item for which the filter matched and compares it to
// expectedValue. The filter matches when it is a:
//   - [TestDeep] operator and it matches for the item;
//   - function receiving the item and it returns true.
//
// expectedValue can of course be a [TestDeep] operator.
//
//	got := []int{-3, -2, -1, 0, 1, 2, 3}
//	td.Cmp(t, got, td.Last(td.Lt(0), -1))                                   // succeeds
//	td.Cmp(t, got, td.Last(func(x int) bool { return x%2 == 0 }, 2))        // succeeds
//	td.Cmp(t, got, td.Last(func(x int) bool { return x%2 == 0 }, td.Gt(0))) // succeeds
//
// If the input is empty (and/or nil for a slice), an "item not found"
// error is raised before comparing to expectedValue.
//
//	var got []int
//	td.Cmp(t, got, td.Last(td.Gt(0), td.Gt(0)))      // fails
//	td.Cmp(t, []int{}, td.Last(td.Gt(0), td.Gt(0)))  // fails
//	td.Cmp(t, [0]int{}, td.Last(td.Gt(0), td.Gt(0))) // fails
//
// See also [First] and [Grep].
func Last(filter, expectedValue any) TestDeep {
	g := tdLast{}
	g.initGrepBase(filter, expectedValue)
	return &g
}

func (g *tdLast) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	if g.err != nil {
		return ctx.CollectError(g.err)
	}

	if rErr := grepResolvePtr(ctx, &got); rErr != nil {
		return rErr
	}

	switch got.Kind() {
	case reflect.Slice, reflect.Array:
		for idx := got.Len() - 1; idx >= 0; idx-- {
			item := got.Index(idx)
			ok, rErr := g.matchItem(ctx, idx, item)
			if rErr != nil {
				return rErr
			}
			if ok {
				return deepValueEqual(
					ctx.AddCustomLevel(S("<last#%d>", idx)),
					item,
					g.expectedValue,
				)
			}
		}
		return g.notFound(ctx, got)
	}

	return grepBadKind(ctx, got)
}

func (g *tdLast) TypeBehind() reflect.Type {
	return g.sliceTypeBehind()
}
