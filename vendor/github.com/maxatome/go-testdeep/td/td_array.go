// Copyright (c) 2018-2021, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"

	"github.com/maxatome/go-testdeep/internal/ctxerr"
	"github.com/maxatome/go-testdeep/internal/types"
	"github.com/maxatome/go-testdeep/internal/util"
)

type tdArray struct {
	tdExpectedType
	expectedEntries []reflect.Value
	onlyIndexes     []int // only used by SuperSliceOf, nil otherwise
}

var _ TestDeep = &tdArray{}

// ArrayEntries allows to pass array or slice entries to check in
// functions Array, Slice and SuperSliceOf. It is a map whose each key
// is the item index and the corresponding value the expected item
// value (which can be a TestDeep operator as well as a zero value).
type ArrayEntries map[int]interface{}

const (
	arrayArray uint = iota
	arraySlice
	arraySuper
)

func newArray(kind uint, model interface{}, expectedEntries ArrayEntries) *tdArray {
	vmodel := reflect.ValueOf(model)

	a := tdArray{
		tdExpectedType: tdExpectedType{
			base: newBase(4),
		},
	}

	if kind == arraySuper {
		a.onlyIndexes = make([]int, 0, len(expectedEntries))
	}

	kindIsOK := func(k reflect.Kind) bool {
		switch kind {
		case arrayArray:
			return k == reflect.Array
		case arraySlice:
			return k == reflect.Slice
		default: // arraySuper
			return k == reflect.Slice || k == reflect.Array
		}
	}

	switch vk := vmodel.Kind(); {
	case vk == reflect.Ptr:
		if !kindIsOK(vmodel.Type().Elem().Kind()) {
			break
		}

		a.isPtr = true

		if vmodel.IsNil() {
			a.expectedType = vmodel.Type().Elem()
			a.populateExpectedEntries(expectedEntries, reflect.Value{})
			return &a
		}

		vmodel = vmodel.Elem()
		fallthrough

	case kindIsOK(vk):
		a.expectedType = vmodel.Type()
		a.populateExpectedEntries(expectedEntries, vmodel)
		return &a
	}

	switch kind {
	case arrayArray:
		a.err = ctxerr.OpBadUsage("Array",
			"(ARRAY|&ARRAY, EXPECTED_ENTRIES)", model, 1, true)
	case arraySlice:
		a.err = ctxerr.OpBadUsage("Slice",
			"(SLICE|&SLICE, EXPECTED_ENTRIES)", model, 1, true)
	default: // arraySuper
		a.err = ctxerr.OpBadUsage("SuperSliceOf",
			"(ARRAY|&ARRAY|SLICE|&SLICE, EXPECTED_ENTRIES)", model, 1, true)
	}
	return &a
}

// summary(Array): compares the contents of an array or a pointer on an array
// input(Array): array,ptr(ptr on array)

// Array operator compares the contents of an array or a pointer on an
// array against the values of "model" and the values of
// "expectedEntries". Entries with zero values of "model" are ignored
// if the same entry is present in "expectedEntries", otherwise they
// are taken into account. An entry cannot be present in both "model"
// and "expectedEntries", except if it is a zero-value in "model". At
// the end, all entries are checked. To check only some entries of an
// array, see SuperSliceOf operator.
//
// "model" must be the same type as compared data.
//
// "expectedEntries" can be nil, if no zero entries are expected and
// no TestDeep operator are involved.
//
//   got := [3]int{12, 14, 17}
//   td.Cmp(t, got, td.Array([3]int{0, 14}, td.ArrayEntries{0: 12, 2: 17})) // succeeds
//   td.Cmp(t, &got,
//     td.Array(&[3]int{0, 14}, td.ArrayEntries{0: td.Gt(10), 2: td.Gt(15)})) // succeeds
//
// TypeBehind method returns the reflect.Type of "model".
func Array(model interface{}, expectedEntries ArrayEntries) TestDeep {
	return newArray(arrayArray, model, expectedEntries)
}

// summary(Slice): compares the contents of a slice or a pointer on a slice
// input(Slice): slice,ptr(ptr on slice)

// Slice operator compares the contents of a slice or a pointer on a
// slice against the values of "model" and the values of
// "expectedEntries". Entries with zero values of "model" are ignored
// if the same entry is present in "expectedEntries", otherwise they
// are taken into account. An entry cannot be present in both "model"
// and "expectedEntries", except if it is a zero-value in "model". At
// the end, all entries are checked. To check only some entries of a
// slice, see SuperSliceOf operator.
//
// "model" must be the same type as compared data.
//
// "expectedEntries" can be nil, if no zero entries are expected and
// no TestDeep operator are involved.
//
//   got := []int{12, 14, 17}
//   td.Cmp(t, got, td.Slice([]int{0, 14}, td.ArrayEntries{0: 12, 2: 17})) // succeeds
//   td.Cmp(t, &got,
//     td.Slice(&[]int{0, 14}, td.ArrayEntries{0: td.Gt(10), 2: td.Gt(15)})) // succeeds
//
// TypeBehind method returns the reflect.Type of "model".
func Slice(model interface{}, expectedEntries ArrayEntries) TestDeep {
	return newArray(arraySlice, model, expectedEntries)
}

// summary(SuperSliceOf): compares the contents of a slice, a pointer
// on a slice, an array or a pointer on an array but with potentially
// some extra entries
// input(SuperSliceOf): array,slice,ptr(ptr on array/slice)

// SuperSliceOf operator compares the contents of an array, a pointer
// on an array, a slice or a pointer on a slice against the non-zero
// values of "model" (if any) and the values of "expectedEntries". So
// entries with zero value of "model" are always ignored. If a zero
// value check is needed, this zero value has to be set in
// "expectedEntries". An entry cannot be present in both "model" and
// "expectedEntries", except if it is a zero-value in "model". At the
// end, only entries present in "expectedEntries" and non-zero ones
// present in "model" are checked. To check all entries of an array
// see Array operator. To check all entries of a slice see Slice
// operator.
//
// "model" must be the same type as compared data.
//
// "expectedEntries" can be nil, if no zero entries are expected and
// no TestDeep operator are involved.
//
// Works with slices:
//
//   got := []int{12, 14, 17}
//   td.Cmp(t, got, td.SuperSliceOf([]int{12}, nil))                                // succeeds
//   td.Cmp(t, got, td.SuperSliceOf([]int{12}, td.ArrayEntries{2: 17}))             // succeeds
//   td.Cmp(t, &got, td.SuperSliceOf(&[]int{0, 14}, td.ArrayEntries{2: td.Gt(16)})) // succeeds
//
// and arrays:
//
//   got := [5]int{12, 14, 17, 26, 56}
//   td.Cmp(t, got, td.SuperSliceOf([5]int{12}, nil))                                // succeeds
//   td.Cmp(t, got, td.SuperSliceOf([5]int{12}, td.ArrayEntries{2: 17}))             // succeeds
//   td.Cmp(t, &got, td.SuperSliceOf(&[5]int{0, 14}, td.ArrayEntries{2: td.Gt(16)})) // succeeds
func SuperSliceOf(model interface{}, expectedEntries ArrayEntries) TestDeep {
	return newArray(arraySuper, model, expectedEntries)
}

func (a *tdArray) populateExpectedEntries(expectedEntries ArrayEntries, expectedModel reflect.Value) {
	// Compute highest expected index
	maxExpectedIdx := -1
	for index := range expectedEntries {
		if index > maxExpectedIdx {
			maxExpectedIdx = index
		}
	}

	var numEntries int
	array := a.expectedType.Kind() == reflect.Array
	if array {
		numEntries = a.expectedType.Len()

		if numEntries <= maxExpectedIdx {
			a.err = ctxerr.OpBad(
				a.GetLocation().Func,
				"array length is %d, so cannot have #%d expected index",
				numEntries,
				maxExpectedIdx)
			return
		}
	} else {
		numEntries = maxExpectedIdx + 1

		// If slice is non-nil
		if expectedModel.IsValid() {
			if numEntries < expectedModel.Len() {
				numEntries = expectedModel.Len()
			}
		}
	}

	a.expectedEntries = make([]reflect.Value, numEntries)

	elemType := a.expectedType.Elem()
	var vexpectedValue reflect.Value
	for index, expectedValue := range expectedEntries {
		if expectedValue == nil {
			switch elemType.Kind() {
			case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map,
				reflect.Ptr, reflect.Slice:
				vexpectedValue = reflect.Zero(elemType) // change to a typed nil
			default:
				a.err = ctxerr.OpBad(
					a.GetLocation().Func,
					"expected value of #%d cannot be nil as items type is %s",
					index,
					elemType)
				return
			}
		} else {
			vexpectedValue = reflect.ValueOf(expectedValue)

			if _, ok := expectedValue.(TestDeep); !ok {
				if !vexpectedValue.Type().AssignableTo(elemType) {
					a.err = ctxerr.OpBad(
						a.GetLocation().Func,
						"type %s of #%d expected value differs from %s contents (%s)",
						vexpectedValue.Type(),
						index,
						util.TernStr(array, "array", "slice"),
						elemType)
					return
				}
			}
		}

		a.expectedEntries[index] = vexpectedValue

		// SuperSliceOf
		if a.onlyIndexes != nil {
			a.onlyIndexes = append(a.onlyIndexes, index)
		}
	}

	vzero := reflect.Zero(elemType)
	// Check initialized entries in model
	if expectedModel.IsValid() {
		zero := vzero.Interface()
		for index := expectedModel.Len() - 1; index >= 0; index-- {
			ventry := expectedModel.Index(index)

			modelIsZero := reflect.DeepEqual(zero, ventry.Interface())

			// Entry already expected
			if _, ok := expectedEntries[index]; ok {
				// If non-zero entry, consider it as an error (= 2 expected
				// values for the same item)
				if !modelIsZero {
					a.err = ctxerr.OpBad(
						a.GetLocation().Func,
						"non zero #%d entry in model already exists in expectedEntries",
						index)
					return
				}
				continue
			}

			// Expect this entry except if not SuperSliceOf || not zero entry
			if a.onlyIndexes == nil || !modelIsZero {
				a.expectedEntries[index] = ventry

				// SuperSliceOf
				if a.onlyIndexes != nil {
					a.onlyIndexes = append(a.onlyIndexes, index)
				}
			}
		}
	} else if a.expectedType.Kind() == reflect.Slice {
		sort.Ints(a.onlyIndexes)
		// nil slice
		return
	}

	// For SuperSliceOf, we don't want to initialize missing entries
	if a.onlyIndexes != nil {
		sort.Ints(a.onlyIndexes)
		return
	}

	var index int

	// Array case, all is OK
	if array {
		// Non-nil array => a.expectedEntries already fully initialized
		if expectedModel.IsValid() {
			return
		}
		// nil array => a.expectedEntries must be initialized from index=0
		// to numEntries - 1 below
	} else {
		// Non-nil slice => a.expectedEntries must be initialized from
		// index=len(slice) to last entry index of expectedEntries
		index = expectedModel.Len()
	}

	// Slice case, initialize missing expected items to zero
	for ; index < numEntries; index++ {
		if _, ok := expectedEntries[index]; !ok {
			a.expectedEntries[index] = vzero
		}
	}
}

func (a *tdArray) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	if a.err != nil {
		return ctx.CollectError(a.err)
	}

	err := a.checkPtr(ctx, &got, true)
	if err != nil {
		return ctx.CollectError(err)
	}

	err = a.checkType(ctx, got)
	if err != nil {
		return ctx.CollectError(err)
	}

	gotLen := got.Len()

	check := func(index int, expectedValue reflect.Value) *ctxerr.Error {
		curCtx := ctx.AddArrayIndex(index)

		if index >= gotLen {
			if curCtx.BooleanError {
				return ctxerr.BooleanError
			}
			return curCtx.CollectError(&ctxerr.Error{
				Message:  "expected value out of range",
				Got:      types.RawString("<non-existent value>"),
				Expected: expectedValue,
			})
		}

		return deepValueEqual(curCtx, got.Index(index), expectedValue)
	}

	// SuperSliceOf, only check some indexes
	if a.onlyIndexes != nil {
		for _, index := range a.onlyIndexes {
			err = check(index, a.expectedEntries[index])
			if err != nil {
				return err
			}
		}
		return nil
	}

	// Array or Slice
	for index, expectedValue := range a.expectedEntries {
		err = check(index, expectedValue)
		if err != nil {
			return err
		}
	}

	if gotLen > len(a.expectedEntries) {
		if ctx.BooleanError {
			return ctxerr.BooleanError
		}
		return ctx.AddArrayIndex(len(a.expectedEntries)).
			CollectError(&ctxerr.Error{
				Message:  "got value out of range",
				Got:      got.Index(len(a.expectedEntries)),
				Expected: types.RawString("<non-existent value>"),
			})
	}

	return nil
}

func (a *tdArray) String() string {
	if a.err != nil {
		return a.stringError()
	}

	buf := bytes.NewBufferString(a.GetLocation().Func)
	buf.WriteByte('(')
	buf.WriteString(a.expectedTypeStr())

	if len(a.expectedEntries) == 0 {
		buf.WriteString("{})")
	} else {
		buf.WriteString("{\n")

		for index, expectedValue := range a.expectedEntries {
			fmt.Fprintf(buf, "  %d: %s\n", // nolint: errcheck
				index, util.ToString(expectedValue))
		}

		buf.WriteString("})")
	}
	return buf.String()
}
