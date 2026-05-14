// Copyright (c) 2018-2025, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.
//
// DO NOT EDIT!!! AUTOMATICALLY GENERATED!!!

package td

import (
	"time"
)

// All is a shortcut for:
//
//	t.Cmp(got, td.All(expectedValues...), args...)
//
// See [All] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) All(got any, expectedValues []any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, All(expectedValues...), args...)
}

// Any is a shortcut for:
//
//	t.Cmp(got, td.Any(expectedValues...), args...)
//
// See [Any] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) Any(got any, expectedValues []any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, Any(expectedValues...), args...)
}

// Array is a shortcut for:
//
//	t.Cmp(got, td.Array(model, expectedEntries), args...)
//
// See [Array] for details.
//
// [Array] optional parameter expectedEntries is here mandatory.
// nil value should be passed to mimic its absence in
// original [Array] call.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) Array(got, model any, expectedEntries ArrayEntries, args ...any) bool {
	t.Helper()
	return t.Cmp(got, Array(model, expectedEntries), args...)
}

// ArrayEach is a shortcut for:
//
//	t.Cmp(got, td.ArrayEach(expectedValue), args...)
//
// See [ArrayEach] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) ArrayEach(got, expectedValue any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, ArrayEach(expectedValue), args...)
}

// Bag is a shortcut for:
//
//	t.Cmp(got, td.Bag(expectedItems...), args...)
//
// See [Bag] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) Bag(got any, expectedItems []any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, Bag(expectedItems...), args...)
}

// Between is a shortcut for:
//
//	t.Cmp(got, td.Between(from, to, bounds), args...)
//
// See [Between] for details.
//
// [Between] optional parameter bounds is here mandatory.
// [BoundsInIn] value should be passed to mimic its absence in
// original [Between] call.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) Between(got, from, to any, bounds BoundsKind, args ...any) bool {
	t.Helper()
	return t.Cmp(got, Between(from, to, bounds), args...)
}

// Cap is a shortcut for:
//
//	t.Cmp(got, td.Cap(expectedCap), args...)
//
// See [Cap] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) Cap(got, expectedCap any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, Cap(expectedCap), args...)
}

// Code is a shortcut for:
//
//	t.Cmp(got, td.Code(fn), args...)
//
// See [Code] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) Code(got, fn any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, Code(fn), args...)
}

// Contains is a shortcut for:
//
//	t.Cmp(got, td.Contains(expectedValue), args...)
//
// See [Contains] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) Contains(got, expectedValue any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, Contains(expectedValue), args...)
}

// ContainsKey is a shortcut for:
//
//	t.Cmp(got, td.ContainsKey(expectedValue), args...)
//
// See [ContainsKey] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) ContainsKey(got, expectedValue any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, ContainsKey(expectedValue), args...)
}

// Empty is a shortcut for:
//
//	t.Cmp(got, td.Empty(), args...)
//
// See [Empty] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) Empty(got any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, Empty(), args...)
}

// CmpErrorIs is a shortcut for:
//
//	t.Cmp(got, td.ErrorIs(expectedError), args...)
//
// See [ErrorIs] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) CmpErrorIs(got, expectedError any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, ErrorIs(expectedError), args...)
}

// First is a shortcut for:
//
//	t.Cmp(got, td.First(filter, expectedValue), args...)
//
// See [First] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) First(got, filter, expectedValue any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, First(filter, expectedValue), args...)
}

// Grep is a shortcut for:
//
//	t.Cmp(got, td.Grep(filter, expectedValue), args...)
//
// See [Grep] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) Grep(got, filter, expectedValue any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, Grep(filter, expectedValue), args...)
}

// Gt is a shortcut for:
//
//	t.Cmp(got, td.Gt(minExpectedValue), args...)
//
// See [Gt] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) Gt(got, minExpectedValue any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, Gt(minExpectedValue), args...)
}

// Gte is a shortcut for:
//
//	t.Cmp(got, td.Gte(minExpectedValue), args...)
//
// See [Gte] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) Gte(got, minExpectedValue any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, Gte(minExpectedValue), args...)
}

// HasPrefix is a shortcut for:
//
//	t.Cmp(got, td.HasPrefix(expected), args...)
//
// See [HasPrefix] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) HasPrefix(got any, expected string, args ...any) bool {
	t.Helper()
	return t.Cmp(got, HasPrefix(expected), args...)
}

// HasSuffix is a shortcut for:
//
//	t.Cmp(got, td.HasSuffix(expected), args...)
//
// See [HasSuffix] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) HasSuffix(got any, expected string, args ...any) bool {
	t.Helper()
	return t.Cmp(got, HasSuffix(expected), args...)
}

// Isa is a shortcut for:
//
//	t.Cmp(got, td.Isa(model), args...)
//
// See [Isa] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) Isa(got, model any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, Isa(model), args...)
}

// JSON is a shortcut for:
//
//	t.Cmp(got, td.JSON(expectedJSON, params...), args...)
//
// See [JSON] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) JSON(got, expectedJSON any, params []any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, JSON(expectedJSON, params...), args...)
}

// JSONPointer is a shortcut for:
//
//	t.Cmp(got, td.JSONPointer(ptr, expectedValue), args...)
//
// See [JSONPointer] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) JSONPointer(got any, ptr string, expectedValue any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, JSONPointer(ptr, expectedValue), args...)
}

// Keys is a shortcut for:
//
//	t.Cmp(got, td.Keys(val), args...)
//
// See [Keys] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) Keys(got, val any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, Keys(val), args...)
}

// Last is a shortcut for:
//
//	t.Cmp(got, td.Last(filter, expectedValue), args...)
//
// See [Last] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) Last(got, filter, expectedValue any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, Last(filter, expectedValue), args...)
}

// CmpLax is a shortcut for:
//
//	t.Cmp(got, td.Lax(expectedValue), args...)
//
// See [Lax] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) CmpLax(got, expectedValue any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, Lax(expectedValue), args...)
}

// Len is a shortcut for:
//
//	t.Cmp(got, td.Len(expectedLen), args...)
//
// See [Len] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) Len(got, expectedLen any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, Len(expectedLen), args...)
}

// List is a shortcut for:
//
//	t.Cmp(got, td.List(expectedValues...), args...)
//
// See [List] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) List(got any, expectedValues []any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, List(expectedValues...), args...)
}

// Lt is a shortcut for:
//
//	t.Cmp(got, td.Lt(maxExpectedValue), args...)
//
// See [Lt] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) Lt(got, maxExpectedValue any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, Lt(maxExpectedValue), args...)
}

// Lte is a shortcut for:
//
//	t.Cmp(got, td.Lte(maxExpectedValue), args...)
//
// See [Lte] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) Lte(got, maxExpectedValue any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, Lte(maxExpectedValue), args...)
}

// Map is a shortcut for:
//
//	t.Cmp(got, td.Map(model, expectedEntries), args...)
//
// See [Map] for details.
//
// [Map] optional parameter expectedEntries is here mandatory.
// nil value should be passed to mimic its absence in
// original [Map] call.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) Map(got, model any, expectedEntries MapEntries, args ...any) bool {
	t.Helper()
	return t.Cmp(got, Map(model, expectedEntries), args...)
}

// MapEach is a shortcut for:
//
//	t.Cmp(got, td.MapEach(expectedValue), args...)
//
// See [MapEach] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) MapEach(got, expectedValue any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, MapEach(expectedValue), args...)
}

// N is a shortcut for:
//
//	t.Cmp(got, td.N(num, tolerance), args...)
//
// See [N] for details.
//
// [N] optional parameter tolerance is here mandatory.
// 0 value should be passed to mimic its absence in
// original [N] call.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) N(got, num, tolerance any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, N(num, tolerance), args...)
}

// NaN is a shortcut for:
//
//	t.Cmp(got, td.NaN(), args...)
//
// See [NaN] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) NaN(got any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, NaN(), args...)
}

// Nil is a shortcut for:
//
//	t.Cmp(got, td.Nil(), args...)
//
// See [Nil] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) Nil(got any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, Nil(), args...)
}

// None is a shortcut for:
//
//	t.Cmp(got, td.None(notExpectedValues...), args...)
//
// See [None] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) None(got any, notExpectedValues []any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, None(notExpectedValues...), args...)
}

// Not is a shortcut for:
//
//	t.Cmp(got, td.Not(notExpected), args...)
//
// See [Not] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) Not(got, notExpected any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, Not(notExpected), args...)
}

// NotAny is a shortcut for:
//
//	t.Cmp(got, td.NotAny(notExpectedItems...), args...)
//
// See [NotAny] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) NotAny(got any, notExpectedItems []any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, NotAny(notExpectedItems...), args...)
}

// NotEmpty is a shortcut for:
//
//	t.Cmp(got, td.NotEmpty(), args...)
//
// See [NotEmpty] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) NotEmpty(got any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, NotEmpty(), args...)
}

// NotNaN is a shortcut for:
//
//	t.Cmp(got, td.NotNaN(), args...)
//
// See [NotNaN] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) NotNaN(got any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, NotNaN(), args...)
}

// NotNil is a shortcut for:
//
//	t.Cmp(got, td.NotNil(), args...)
//
// See [NotNil] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) NotNil(got any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, NotNil(), args...)
}

// NotZero is a shortcut for:
//
//	t.Cmp(got, td.NotZero(), args...)
//
// See [NotZero] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) NotZero(got any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, NotZero(), args...)
}

// PPtr is a shortcut for:
//
//	t.Cmp(got, td.PPtr(val), args...)
//
// See [PPtr] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) PPtr(got, val any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, PPtr(val), args...)
}

// Ptr is a shortcut for:
//
//	t.Cmp(got, td.Ptr(val), args...)
//
// See [Ptr] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) Ptr(got, val any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, Ptr(val), args...)
}

// Re is a shortcut for:
//
//	t.Cmp(got, td.Re(reg, capture), args...)
//
// See [Re] for details.
//
// [Re] optional parameter capture is here mandatory.
// nil value should be passed to mimic its absence in
// original [Re] call.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) Re(got, reg, capture any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, Re(reg, capture), args...)
}

// ReAll is a shortcut for:
//
//	t.Cmp(got, td.ReAll(reg, capture), args...)
//
// See [ReAll] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) ReAll(got, reg, capture any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, ReAll(reg, capture), args...)
}

// Recv is a shortcut for:
//
//	t.Cmp(got, td.Recv(expectedValue, timeout), args...)
//
// See [Recv] for details.
//
// [Recv] optional parameter timeout is here mandatory.
// 0 value should be passed to mimic its absence in
// original [Recv] call.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) Recv(got, expectedValue any, timeout time.Duration, args ...any) bool {
	t.Helper()
	return t.Cmp(got, Recv(expectedValue, timeout), args...)
}

// Set is a shortcut for:
//
//	t.Cmp(got, td.Set(expectedItems...), args...)
//
// See [Set] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) Set(got any, expectedItems []any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, Set(expectedItems...), args...)
}

// Shallow is a shortcut for:
//
//	t.Cmp(got, td.Shallow(expectedPtr), args...)
//
// See [Shallow] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) Shallow(got, expectedPtr any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, Shallow(expectedPtr), args...)
}

// Slice is a shortcut for:
//
//	t.Cmp(got, td.Slice(model, expectedEntries), args...)
//
// See [Slice] for details.
//
// [Slice] optional parameter expectedEntries is here mandatory.
// nil value should be passed to mimic its absence in
// original [Slice] call.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) Slice(got, model any, expectedEntries ArrayEntries, args ...any) bool {
	t.Helper()
	return t.Cmp(got, Slice(model, expectedEntries), args...)
}

// Smuggle is a shortcut for:
//
//	t.Cmp(got, td.Smuggle(fn, expectedValue), args...)
//
// See [Smuggle] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) Smuggle(got, fn, expectedValue any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, Smuggle(fn, expectedValue), args...)
}

// Sort is a shortcut for:
//
//	t.Cmp(got, td.Sort(how, expectedValue), args...)
//
// See [Sort] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) Sort(got, how, expectedValue any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, Sort(how, expectedValue), args...)
}

// Sorted is a shortcut for:
//
//	t.Cmp(got, td.Sorted(how), args...)
//
// See [Sorted] for details.
//
// [Sorted] optional parameter how is here mandatory.
// nil value should be passed to mimic its absence in
// original [Sorted] call.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) Sorted(got, how any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, Sorted(how), args...)
}

// SStruct is a shortcut for:
//
//	t.Cmp(got, td.SStruct(model, expectedFields), args...)
//
// See [SStruct] for details.
//
// [SStruct] optional parameter expectedFields is here mandatory.
// nil value should be passed to mimic its absence in
// original [SStruct] call.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) SStruct(got, model any, expectedFields StructFields, args ...any) bool {
	t.Helper()
	return t.Cmp(got, SStruct(model, expectedFields), args...)
}

// String is a shortcut for:
//
//	t.Cmp(got, td.String(expected), args...)
//
// See [String] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) String(got any, expected string, args ...any) bool {
	t.Helper()
	return t.Cmp(got, String(expected), args...)
}

// Struct is a shortcut for:
//
//	t.Cmp(got, td.Struct(model, expectedFields), args...)
//
// See [Struct] for details.
//
// [Struct] optional parameter expectedFields is here mandatory.
// nil value should be passed to mimic its absence in
// original [Struct] call.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) Struct(got, model any, expectedFields StructFields, args ...any) bool {
	t.Helper()
	return t.Cmp(got, Struct(model, expectedFields), args...)
}

// SubBagOf is a shortcut for:
//
//	t.Cmp(got, td.SubBagOf(expectedItems...), args...)
//
// See [SubBagOf] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) SubBagOf(got any, expectedItems []any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, SubBagOf(expectedItems...), args...)
}

// SubJSONOf is a shortcut for:
//
//	t.Cmp(got, td.SubJSONOf(expectedJSON, params...), args...)
//
// See [SubJSONOf] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) SubJSONOf(got, expectedJSON any, params []any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, SubJSONOf(expectedJSON, params...), args...)
}

// SubMapOf is a shortcut for:
//
//	t.Cmp(got, td.SubMapOf(model, expectedEntries), args...)
//
// See [SubMapOf] for details.
//
// [SubMapOf] optional parameter expectedEntries is here mandatory.
// nil value should be passed to mimic its absence in
// original [SubMapOf] call.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) SubMapOf(got, model any, expectedEntries MapEntries, args ...any) bool {
	t.Helper()
	return t.Cmp(got, SubMapOf(model, expectedEntries), args...)
}

// SubSetOf is a shortcut for:
//
//	t.Cmp(got, td.SubSetOf(expectedItems...), args...)
//
// See [SubSetOf] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) SubSetOf(got any, expectedItems []any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, SubSetOf(expectedItems...), args...)
}

// SuperBagOf is a shortcut for:
//
//	t.Cmp(got, td.SuperBagOf(expectedItems...), args...)
//
// See [SuperBagOf] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) SuperBagOf(got any, expectedItems []any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, SuperBagOf(expectedItems...), args...)
}

// SuperJSONOf is a shortcut for:
//
//	t.Cmp(got, td.SuperJSONOf(expectedJSON, params...), args...)
//
// See [SuperJSONOf] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) SuperJSONOf(got, expectedJSON any, params []any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, SuperJSONOf(expectedJSON, params...), args...)
}

// SuperMapOf is a shortcut for:
//
//	t.Cmp(got, td.SuperMapOf(model, expectedEntries), args...)
//
// See [SuperMapOf] for details.
//
// [SuperMapOf] optional parameter expectedEntries is here mandatory.
// nil value should be passed to mimic its absence in
// original [SuperMapOf] call.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) SuperMapOf(got, model any, expectedEntries MapEntries, args ...any) bool {
	t.Helper()
	return t.Cmp(got, SuperMapOf(model, expectedEntries), args...)
}

// SuperSetOf is a shortcut for:
//
//	t.Cmp(got, td.SuperSetOf(expectedItems...), args...)
//
// See [SuperSetOf] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) SuperSetOf(got any, expectedItems []any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, SuperSetOf(expectedItems...), args...)
}

// SuperSliceOf is a shortcut for:
//
//	t.Cmp(got, td.SuperSliceOf(model, expectedEntries), args...)
//
// See [SuperSliceOf] for details.
//
// [SuperSliceOf] optional parameter expectedEntries is here mandatory.
// nil value should be passed to mimic its absence in
// original [SuperSliceOf] call.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) SuperSliceOf(got, model any, expectedEntries ArrayEntries, args ...any) bool {
	t.Helper()
	return t.Cmp(got, SuperSliceOf(model, expectedEntries), args...)
}

// TruncTime is a shortcut for:
//
//	t.Cmp(got, td.TruncTime(expectedTime, trunc), args...)
//
// See [TruncTime] for details.
//
// [TruncTime] optional parameter trunc is here mandatory.
// 0 value should be passed to mimic its absence in
// original [TruncTime] call.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) TruncTime(got, expectedTime any, trunc time.Duration, args ...any) bool {
	t.Helper()
	return t.Cmp(got, TruncTime(expectedTime, trunc), args...)
}

// Values is a shortcut for:
//
//	t.Cmp(got, td.Values(val), args...)
//
// See [Values] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) Values(got, val any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, Values(val), args...)
}

// Zero is a shortcut for:
//
//	t.Cmp(got, td.Zero(), args...)
//
// See [Zero] for details.
//
// Returns true if the test is OK, false if it fails.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) Zero(got any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, Zero(), args...)
}
