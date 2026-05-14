// Copyright (c) 2018-2025, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/maxatome/go-testdeep/internal/ctxerr"
	"github.com/maxatome/go-testdeep/internal/flat"
	"github.com/maxatome/go-testdeep/internal/util"
)

type tdListBase struct {
	baseOKNil
	items []reflect.Value
}

func newListBase(items ...any) tdListBase {
	return tdListBase{
		baseOKNil: newBaseOKNil(4),
		items:     flat.Values(items),
	}
}

func (l *tdListBase) String() string {
	var b strings.Builder
	b.WriteString(l.GetLocation().Func)
	return util.SliceToString(&b, l.items).String()
}

type tdList struct {
	tdListBase
}

var _ TestDeep = &tdList{}

// summary(List): compares the contents of an array or a slice with taking
// care of the order of items
// input(List): array,slice,ptr(ptr on array/slice)

// List operator compares the contents of an array or a slice (or a
// pointer on array/slice) with taking care of the order of items.
//
// [Array] and [Slice] need to specify the type of array/slice being
// compared then to index all expected items. List does not. It acts
// as comparing a literal array/slice, but without having to specify
// the type and allowing to easily use TestDeep operators:
//
//	td.Cmp(t, []int{1, 9, 5}, td.List(1, 9, 5))                              // succeeds
//	td.Cmp(t, []int{1, 9, 5}, td.List(td.Gt(0), td.Between(8, 9), td.Lt(5))) // succeeds
//	td.Cmp(t, []int{1, 9, 5}, td.List(1, 9))                                 // fails, 5 is extra
//	td.Cmp(t, []int{1, 9, 5}, td.List(1, 9, 5, 4))                           // fails, 4 is missing
//
//	// works with slices/arrays of any type
//	td.Cmp(t, personSlice, td.List(
//	  Person{Name: "Bob", Age: 32},
//	  Person{Name: "Alice", Age: 26},
//	))
//
// To flatten a non-[]any slice/array, use [Flatten] function
// and so avoid boring and inefficient copies:
//
//	expected := []int{1, 2, 1}
//	td.Cmp(t, []int{1, 1, 2}, td.List(td.Flatten(expected))) // succeeds
//	// = td.Cmp(t, []int{1, 1, 2}, td.List(1, 2, 1))
//
//	// Compare only Name field of a slice of Person structs
//	td.Cmp(t, personSlice, td.List(td.Flatten([]string{"Bob", "Alice"}, "Smuggle:Name")))
//
// TypeBehind method can return a non-nil [reflect.Type] if all items
// known non-interface types are equal, or if only interface types
// are found (mostly issued from Isa()) and they are equal.
//
// See also [Bag], [Set] and [Sort].
func List(expectedValues ...any) TestDeep {
	return &tdList{
		tdListBase: newListBase(expectedValues...),
	}
}

func (l *tdList) Match(ctx ctxerr.Context, got reflect.Value) (err *ctxerr.Error) {
	switch got.Kind() {
	case reflect.Ptr:
		gotElem := got.Elem()
		if !gotElem.IsValid() {
			if ctx.BooleanError {
				return ctxerr.BooleanError
			}
			return ctx.CollectError(ctxerr.NilPointer(got, "non-nil *slice OR *array"))
		}

		if gotElem.Kind() != reflect.Array && gotElem.Kind() != reflect.Slice {
			break
		}
		got = gotElem
		fallthrough

	case reflect.Array, reflect.Slice:
		gotLen, expectedLen := got.Len(), len(l.items)
		if ctx.BooleanError && gotLen != expectedLen {
			return ctxerr.BooleanError // shortcut in boolean context
		}

		var maxLen int
		if gotLen >= expectedLen {
			maxLen = expectedLen
		} else {
			maxLen = gotLen
		}

		for i := 0; i < maxLen; i++ {
			err = deepValueEqual(ctx.AddArrayIndex(i), got.Index(i), l.items[i])
			if err != nil {
				return err
			}
		}
		if gotLen == expectedLen {
			return
		}

		res := tdSetResult{
			Kind: itemsSetResult,
			// do not sort Extra/Mising here
		}

		if gotLen > expectedLen {
			res.Extra = make([]reflect.Value, gotLen-expectedLen)
			for i := expectedLen; i < gotLen; i++ {
				res.Extra[i-expectedLen] = got.Index(i)
			}
		} else {
			res.Missing = l.items[gotLen:]
		}
		return ctx.CollectError(&ctxerr.Error{
			Message: fmt.Sprintf("comparing %s, from index #%d", got.Kind(), maxLen),
			Summary: res.Summary(),
		})
	}

	if ctx.BooleanError {
		return ctxerr.BooleanError
	}
	return ctx.CollectError(ctxerr.BadKind(got, "slice OR array OR *slice OR *array"))
}

func (l *tdList) TypeBehind() reflect.Type {
	typ := uniqTypeBehindSlice(l.items)
	if typ == nil {
		return nil
	}
	return reflect.SliceOf(typ)
}
