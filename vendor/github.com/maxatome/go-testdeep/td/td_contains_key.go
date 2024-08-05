// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

import (
	"reflect"

	"github.com/maxatome/go-testdeep/helpers/tdutil"
	"github.com/maxatome/go-testdeep/internal/ctxerr"
	"github.com/maxatome/go-testdeep/internal/types"
	"github.com/maxatome/go-testdeep/internal/util"
)

type tdContainsKey struct {
	tdSmugglerBase
}

var _ TestDeep = &tdContainsKey{}

// summary(ContainsKey): checks that a map contains a key
// input(ContainsKey): map

// ContainsKey is a smuggler operator and works on maps only. It
<<<<<<< HEAD
<<<<<<< HEAD
// compares each key of map against expectedValue.
//
//	hash := map[string]int{"foo": 12, "bar": 34, "zip": 28}
//	td.Cmp(t, hash, td.ContainsKey("foo"))             // succeeds
//	td.Cmp(t, hash, td.ContainsKey(td.HasPrefix("z"))) // succeeds
//	td.Cmp(t, hash, td.ContainsKey(td.HasPrefix("x"))) // fails
//
//	hnum := map[int]string{1: "foo", 42: "bar"}
//	td.Cmp(t, hash, td.ContainsKey(42))                 // succeeds
//	td.Cmp(t, hash, td.ContainsKey(td.Between(40, 45))) // succeeds
//
// When ContainsKey(nil) is used, nil is automatically converted to a
// typed nil on the fly to avoid confusion (if the map key type allows
// it of course.) So all following [Cmp] calls are equivalent
// (except the (*byte)(nil) one):
//
//	num := 123
//	hnum := map[*int]bool{&num: true, nil: true}
//	td.Cmp(t, hnum, td.ContainsKey(nil))         // succeeds → (*int)(nil)
//	td.Cmp(t, hnum, td.ContainsKey((*int)(nil))) // succeeds
//	td.Cmp(t, hnum, td.ContainsKey(td.Nil()))    // succeeds
//	// But...
//	td.Cmp(t, hnum, td.ContainsKey((*byte)(nil))) // fails: (*byte)(nil) ≠ (*int)(nil)
//
// See also [Contains].
func ContainsKey(expectedValue any) TestDeep {
	c := tdContainsKey{
		tdSmugglerBase: newSmugglerBase(expectedValue),
	}

	if !c.isTestDeeper {
		c.expectedValue = reflect.ValueOf(expectedValue)
	}
	return &c
}

func (c *tdContainsKey) doesNotContainKey(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	if ctx.BooleanError {
		return ctxerr.BooleanError
	}
	return ctx.CollectError(&ctxerr.Error{
		Message: "does not contain key",
		Summary: ctxerr.ErrorSummaryItems{
			{
				Label: "expected key",
				Value: util.ToString(c.expectedValue),
			},
			{
				Label: "not in keys",
				Value: util.ToString(tdutil.MapSortedKeys(got)),
			},
		},
	})
}

// getExpectedValue returns the expected value handling the
// Contains(nil) case: in this case it returns a typed nil (same type
// as the keys of got).
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
// got is a map (it's the caller responsibility to check).
func (c *tdContainsKey) getExpectedValue(got reflect.Value) reflect.Value {
	// If the expectValue is non-typed nil
	if !c.expectedValue.IsValid() {
		// AND the kind of items in got is...
		switch got.Type().Key().Kind() {
		case reflect.Chan, reflect.Func, reflect.Interface,
			reflect.Map, reflect.Ptr, reflect.Slice:
			// returns a typed nil
			return reflect.Zero(got.Type().Key())
		}
	}
	return c.expectedValue
}

func (c *tdContainsKey) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	if got.Kind() == reflect.Map {
		expectedValue := c.getExpectedValue(got)

		// If expected value is a TestDeep operator OR BeLax, check each key
		if c.isTestDeeper || ctx.BeLax {
			for _, k := range got.MapKeys() {
				if deepValueEqualFinalOK(ctx, k, expectedValue) {
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// got is a map (it's the caller responsibility to check)
||||||| parent of 5ce8c7613 (update vendored files)
// got is a map (it's the caller responsibility to check)
=======
// got is a map (it's the caller responsibility to check).
>>>>>>> 5ce8c7613 (update vendored files)
func (c *tdContainsKey) getExpectedValue(got reflect.Value) reflect.Value {
	// If the expectValue is non-typed nil
	if !c.expectedValue.IsValid() {
		// AND the kind of items in got is...
		switch got.Type().Key().Kind() {
		case reflect.Chan, reflect.Func, reflect.Interface,
			reflect.Map, reflect.Ptr, reflect.Slice:
			// returns a typed nil
			return reflect.Zero(got.Type().Key())
		}
	}
	return c.expectedValue
}

func (c *tdContainsKey) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	if got.Kind() == reflect.Map {
		expectedValue := c.getExpectedValue(got)

		// If expected value is a TestDeep operator OR BeLax, check each key
		if c.isTestDeeper || ctx.BeLax {
			for _, k := range got.MapKeys() {
<<<<<<< HEAD
				if deepValueEqualOK(k, expectedValue) {
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
				if deepValueEqualOK(k, expectedValue) {
=======
				if deepValueEqualFinalOK(ctx, k, expectedValue) {
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// got is a map (it's the caller responsibility to check)
||||||| parent of 6b7ce455e (update vendored files)
// got is a map (it's the caller responsibility to check)
=======
// got is a map (it's the caller responsibility to check).
>>>>>>> 6b7ce455e (update vendored files)
func (c *tdContainsKey) getExpectedValue(got reflect.Value) reflect.Value {
	// If the expectValue is non-typed nil
	if !c.expectedValue.IsValid() {
		// AND the kind of items in got is...
		switch got.Type().Key().Kind() {
		case reflect.Chan, reflect.Func, reflect.Interface,
			reflect.Map, reflect.Ptr, reflect.Slice:
			// returns a typed nil
			return reflect.Zero(got.Type().Key())
		}
	}
	return c.expectedValue
}

func (c *tdContainsKey) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	if got.Kind() == reflect.Map {
		expectedValue := c.getExpectedValue(got)

		// If expected value is a TestDeep operator OR BeLax, check each key
		if c.isTestDeeper || ctx.BeLax {
			for _, k := range got.MapKeys() {
<<<<<<< HEAD
				if deepValueEqualOK(k, expectedValue) {
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
				if deepValueEqualOK(k, expectedValue) {
=======
				if deepValueEqualFinalOK(ctx, k, expectedValue) {
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// got is a map (it's the caller responsibility to check)
||||||| parent of 4d7e5ad26 (update vendored files)
// got is a map (it's the caller responsibility to check)
=======
// got is a map (it's the caller responsibility to check).
>>>>>>> 4d7e5ad26 (update vendored files)
func (c *tdContainsKey) getExpectedValue(got reflect.Value) reflect.Value {
	// If the expectValue is non-typed nil
	if !c.expectedValue.IsValid() {
		// AND the kind of items in got is...
		switch got.Type().Key().Kind() {
		case reflect.Chan, reflect.Func, reflect.Interface,
			reflect.Map, reflect.Ptr, reflect.Slice:
			// returns a typed nil
			return reflect.Zero(got.Type().Key())
		}
	}
	return c.expectedValue
}

func (c *tdContainsKey) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	if got.Kind() == reflect.Map {
		expectedValue := c.getExpectedValue(got)

		// If expected value is a TestDeep operator OR BeLax, check each key
		if c.isTestDeeper || ctx.BeLax {
			for _, k := range got.MapKeys() {
<<<<<<< HEAD
				if deepValueEqualOK(k, expectedValue) {
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
				if deepValueEqualOK(k, expectedValue) {
=======
				if deepValueEqualFinalOK(ctx, k, expectedValue) {
>>>>>>> 4d7e5ad26 (update vendored files)
					return nil
				}
			}
		} else if expectedValue.IsValid() &&
			got.Type().Key() == expectedValue.Type() &&
			got.MapIndex(expectedValue).IsValid() {
			return nil
		}
		return c.doesNotContainKey(ctx, got)
	}

	if ctx.BooleanError {
		return ctxerr.BooleanError
	}
	var expectedType any
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// compares each key of map against "expectedValue".
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
// compares each key of map against "expectedValue".
=======
// compares each key of map against expectedValue.
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
//
//	hash := map[string]int{"foo": 12, "bar": 34, "zip": 28}
//	td.Cmp(t, hash, td.ContainsKey("foo"))             // succeeds
//	td.Cmp(t, hash, td.ContainsKey(td.HasPrefix("z"))) // succeeds
//	td.Cmp(t, hash, td.ContainsKey(td.HasPrefix("x"))) // fails
//
//	hnum := map[int]string{1: "foo", 42: "bar"}
//	td.Cmp(t, hash, td.ContainsKey(42))                 // succeeds
//	td.Cmp(t, hash, td.ContainsKey(td.Between(40, 45))) // succeeds
//
// When ContainsKey(nil) is used, nil is automatically converted to a
// typed nil on the fly to avoid confusion (if the map key type allows
// it of course.) So all following [Cmp] calls are equivalent
// (except the (*byte)(nil) one):
//
//	num := 123
//	hnum := map[*int]bool{&num: true, nil: true}
//	td.Cmp(t, hnum, td.ContainsKey(nil))         // succeeds → (*int)(nil)
//	td.Cmp(t, hnum, td.ContainsKey((*int)(nil))) // succeeds
//	td.Cmp(t, hnum, td.ContainsKey(td.Nil()))    // succeeds
//	// But...
//	td.Cmp(t, hnum, td.ContainsKey((*byte)(nil))) // fails: (*byte)(nil) ≠ (*int)(nil)
//
// See also [Contains].
func ContainsKey(expectedValue any) TestDeep {
	c := tdContainsKey{
		tdSmugglerBase: newSmugglerBase(expectedValue),
	}

	if !c.isTestDeeper {
		c.expectedValue = reflect.ValueOf(expectedValue)
	}
	return &c
}

func (c *tdContainsKey) doesNotContainKey(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	if ctx.BooleanError {
		return ctxerr.BooleanError
	}
	return ctx.CollectError(&ctxerr.Error{
		Message: "does not contain key",
		Summary: ctxerr.ErrorSummaryItems{
			{
				Label: "expected key",
				Value: util.ToString(c.expectedValue),
			},
			{
				Label: "not in keys",
				Value: util.ToString(tdutil.MapSortedKeys(got)),
			},
		},
	})
}

// getExpectedValue returns the expected value handling the
// Contains(nil) case: in this case it returns a typed nil (same type
// as the keys of got).
// got is a map (it's the caller responsibility to check).
func (c *tdContainsKey) getExpectedValue(got reflect.Value) reflect.Value {
	// If the expectValue is non-typed nil
	if !c.expectedValue.IsValid() {
		// AND the kind of items in got is...
		switch got.Type().Key().Kind() {
		case reflect.Chan, reflect.Func, reflect.Interface,
			reflect.Map, reflect.Ptr, reflect.Slice:
			// returns a typed nil
			return reflect.Zero(got.Type().Key())
		}
	}
	return c.expectedValue
}

func (c *tdContainsKey) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	if got.Kind() == reflect.Map {
		expectedValue := c.getExpectedValue(got)

		// If expected value is a TestDeep operator OR BeLax, check each key
		if c.isTestDeeper || ctx.BeLax {
			for _, k := range got.MapKeys() {
				if deepValueEqualFinalOK(ctx, k, expectedValue) {
					return nil
				}
			}
		} else if expectedValue.IsValid() &&
			got.Type().Key() == expectedValue.Type() &&
			got.MapIndex(expectedValue).IsValid() {
			return nil
		}
		return c.doesNotContainKey(ctx, got)
	}

	if ctx.BooleanError {
		return ctxerr.BooleanError
	}
<<<<<<< HEAD
	var expectedType interface{}
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	var expectedType interface{}
=======
	var expectedType any
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	if c.expectedValue.IsValid() {
		expectedType = types.RawString(c.expectedValue.Type().String())
	} else {
		expectedType = c
	}

	return ctx.CollectError(&ctxerr.Error{
		Message:  "cannot check contains key",
		Got:      types.RawString(got.Type().String()),
		Expected: expectedType,
	})
}

func (c *tdContainsKey) String() string {
	return "ContainsKey(" + util.ToString(c.expectedValue) + ")"
}
