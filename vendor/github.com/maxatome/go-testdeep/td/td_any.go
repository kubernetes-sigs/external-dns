// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

import (
	"reflect"

	"github.com/maxatome/go-testdeep/internal/ctxerr"
)

type tdAny struct {
	tdList
}

var _ TestDeep = &tdAny{}

// summary(Any): at least one expected value have to match
// input(Any): all

// Any operator compares data against several expected values. During
// a match, at least one of them has to match to succeed. Consider it
// as a "OR" logical operator.
//
<<<<<<< HEAD
//	td.Cmp(t, "foo", td.Any("bar", "foo", "zip")) // succeeds
//	td.Cmp(t, "foo", td.Any(
//	  td.Len(4),
//	  td.HasPrefix("f"),
//	  td.HasSuffix("z"),
//	)) // succeeds coz "f" prefix
//
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
// Note Flatten function can be used to group or reuse some values or
||||||| parent of e1cd8261c (UPSTREAM: <carry>: update vendored files v0.13.1)
// Note Flatten function can be used to group or reuse some values or
=======
// Note [Flatten] function can be used to group or reuse some values or
>>>>>>> e1cd8261c (UPSTREAM: <carry>: update vendored files v0.13.1)
// operators and so avoid boring and inefficient copies:
//
//	stringOps := td.Flatten([]td.TestDeep{td.HasPrefix("f"), td.HasSuffix("z")})
//	td.Cmp(t, "foobar", td.All(
//	  td.Len(4),
//	  stringOps,
//	)) // succeeds coz "f" prefix
//
// TypeBehind method can return a non-nil [reflect.Type] if all items
// known non-interface types are equal, or if only interface types
// are found (mostly issued from Isa()) and they are equal.
//
// See also [All] and [None].
func Any(expectedValues ...any) TestDeep {
	return &tdAny{
		tdList: newList(expectedValues...),
	}
}

func (a *tdAny) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	for _, item := range a.items {
		if deepValueEqualFinalOK(ctx, got, item) {
			return nil
		}
	}

	if ctx.BooleanError {
		return ctxerr.BooleanError
	}
	return ctx.CollectError(&ctxerr.Error{
		Message:  "comparing with Any",
		Got:      got,
		Expected: a,
	})
}

func (a *tdAny) TypeBehind() reflect.Type {
	return uniqTypeBehindSlice(a.items)
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 5ce8c7613 (update vendored files)
=======
// Note Flatten function can be used to group or reuse some values or
// operators and so avoid boring and inefficient copies:
//
//   stringOps := td.Flatten([]td.TestDeep{td.HasPrefix("f"), td.HasSuffix("z")})
//   td.Cmp(t, "foobar", td.All(
//     td.Len(4),
//     stringOps,
//   )) // succeeds coz "f" prefix
//
>>>>>>> 5ce8c7613 (update vendored files)
// TypeBehind method can return a non-nil reflect.Type if all items
// known non-interface types are equal, or if only interface types
// are found (mostly issued from Isa()) and they are equal.
func Any(expectedValues ...interface{}) TestDeep {
	return &tdAny{
		tdList: newList(expectedValues...),
	}
}

func (a *tdAny) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	for _, item := range a.items {
		if deepValueEqualFinalOK(ctx, got, item) {
			return nil
		}
	}

	if ctx.BooleanError {
		return ctxerr.BooleanError
	}
	return ctx.CollectError(&ctxerr.Error{
		Message:  "comparing with Any",
		Got:      got,
		Expected: a,
	})
}

func (a *tdAny) TypeBehind() reflect.Type {
<<<<<<< HEAD
	return a.uniqTypeBehind()
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
	return a.uniqTypeBehind()
=======
	return uniqTypeBehindSlice(a.items)
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 6b7ce455e (update vendored files)
=======
// Note Flatten function can be used to group or reuse some values or
// operators and so avoid boring and inefficient copies:
//
//   stringOps := td.Flatten([]td.TestDeep{td.HasPrefix("f"), td.HasSuffix("z")})
//   td.Cmp(t, "foobar", td.All(
//     td.Len(4),
//     stringOps,
//   )) // succeeds coz "f" prefix
//
>>>>>>> 6b7ce455e (update vendored files)
// TypeBehind method can return a non-nil reflect.Type if all items
// known non-interface types are equal, or if only interface types
// are found (mostly issued from Isa()) and they are equal.
func Any(expectedValues ...interface{}) TestDeep {
	return &tdAny{
		tdList: newList(expectedValues...),
	}
}

func (a *tdAny) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	for _, item := range a.items {
		if deepValueEqualFinalOK(ctx, got, item) {
			return nil
		}
	}

	if ctx.BooleanError {
		return ctxerr.BooleanError
	}
	return ctx.CollectError(&ctxerr.Error{
		Message:  "comparing with Any",
		Got:      got,
		Expected: a,
	})
}

func (a *tdAny) TypeBehind() reflect.Type {
<<<<<<< HEAD
	return a.uniqTypeBehind()
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
	return a.uniqTypeBehind()
=======
	return uniqTypeBehindSlice(a.items)
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 4d7e5ad26 (update vendored files)
=======
// Note Flatten function can be used to group or reuse some values or
// operators and so avoid boring and inefficient copies:
//
//   stringOps := td.Flatten([]td.TestDeep{td.HasPrefix("f"), td.HasSuffix("z")})
//   td.Cmp(t, "foobar", td.All(
//     td.Len(4),
//     stringOps,
//   )) // succeeds coz "f" prefix
//
>>>>>>> 4d7e5ad26 (update vendored files)
// TypeBehind method can return a non-nil reflect.Type if all items
// known non-interface types are equal, or if only interface types
// are found (mostly issued from Isa()) and they are equal.
func Any(expectedValues ...interface{}) TestDeep {
	return &tdAny{
		tdList: newList(expectedValues...),
	}
}

func (a *tdAny) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	for _, item := range a.items {
		if deepValueEqualFinalOK(ctx, got, item) {
			return nil
		}
	}

	if ctx.BooleanError {
		return ctxerr.BooleanError
	}
	return ctx.CollectError(&ctxerr.Error{
		Message:  "comparing with Any",
		Got:      got,
		Expected: a,
	})
}

func (a *tdAny) TypeBehind() reflect.Type {
<<<<<<< HEAD
	return a.uniqTypeBehind()
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
	return a.uniqTypeBehind()
=======
	return uniqTypeBehindSlice(a.items)
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
//   td.Cmp(t, "foo", td.Any("bar", "foo", "zip")) // succeeds
//   td.Cmp(t, "foo", td.Any(
//     td.Len(4),
//     td.HasPrefix("f"),
//     td.HasSuffix("z"),
//   )) // succeeds coz "f" prefix
//
// TypeBehind method can return a non-nil reflect.Type if all items
// known non-interface types are equal, or if only interface types
// are found (mostly issued from Isa()) and they are equal.
func Any(expectedValues ...interface{}) TestDeep {
	return &tdAny{
		tdList: newList(expectedValues...),
	}
}

func (a *tdAny) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	for _, item := range a.items {
		if deepValueEqualOK(got, item) {
			return nil
		}
	}

	if ctx.BooleanError {
		return ctxerr.BooleanError
	}
	return ctx.CollectError(&ctxerr.Error{
		Message:  "comparing with Any",
		Got:      got,
		Expected: a,
	})
}

func (a *tdAny) TypeBehind() reflect.Type {
	return a.uniqTypeBehind()
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
}
