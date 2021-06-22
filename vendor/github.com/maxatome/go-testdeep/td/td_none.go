// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

import (
	"fmt"
	"reflect"

	"github.com/maxatome/go-testdeep/internal/ctxerr"
)

type tdNone struct {
	tdList
}

var _ TestDeep = &tdNone{}

// summary(None): no values have to match
// input(None): all

// None operator compares data against several not expected
// values. During a match, none of them have to match to succeed.
//
<<<<<<< HEAD
<<<<<<< HEAD
//   td.Cmp(t, 12, td.None(8, 10, 14))     // succeeds
//   td.Cmp(t, 12, td.None(8, 10, 12, 14)) // fails
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
||||||| parent of e1cd8261c (UPSTREAM: <carry>: update vendored files v0.13.1)
//   td.Cmp(t, 12, td.None(8, 10, 14))     // succeeds
//   td.Cmp(t, 12, td.None(8, 10, 12, 14)) // fails
=======
//	td.Cmp(t, 12, td.None(8, 10, 14))     // succeeds
//	td.Cmp(t, 12, td.None(8, 10, 12, 14)) // fails
>>>>>>> e1cd8261c (UPSTREAM: <carry>: update vendored files v0.13.1)
//
// Note [Flatten] function can be used to group or reuse some values or
// operators and so avoid boring and inefficient copies:
//
//	prime := td.Flatten([]int{1, 2, 3, 5, 7, 11, 13})
//	even := td.Flatten([]int{2, 4, 6, 8, 10, 12, 14})
//	td.Cmp(t, 9, td.None(prime, even)) // succeeds
//
// See also [All], [Any] and [Not].
func None(notExpectedValues ...any) TestDeep {
	return &tdNone{
		tdList: newList(notExpectedValues...),
	}
}

// summary(Not): value must not match
// input(Not): all

// Not operator compares data against the not expected value. During a
// match, it must not match to succeed.
//
// Not is the same operator as [None] with only one argument. It is
// provided as a more readable function when only one argument is
// needed.
//
//	td.Cmp(t, 12, td.Not(10)) // succeeds
//	td.Cmp(t, 12, td.Not(12)) // fails
//
// See also [None].
func Not(notExpected any) TestDeep {
	return &tdNone{
		tdList: newList(notExpected),
	}
}

func (n *tdNone) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	for idx, item := range n.items {
		if deepValueEqualFinalOK(ctx, got, item) {
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 5ce8c7613 (update vendored files)
=======
//
// Note Flatten function can be used to group or reuse some values or
// operators and so avoid boring and inefficient copies:
//
//   prime := td.Flatten([]int{1, 2, 3, 5, 7, 11, 13})
//   even := td.Flatten([]int{2, 4, 6, 8, 10, 12, 14})
//   td.Cmp(t, 9, td.None(prime, even)) // succeeds
>>>>>>> 5ce8c7613 (update vendored files)
func None(notExpectedValues ...interface{}) TestDeep {
	return &tdNone{
		tdList: newList(notExpectedValues...),
	}
}

// summary(Not): value must not match
// input(Not): all

// Not operator compares data against the not expected value. During a
// match, it must not match to succeed.
//
// Not is the same operator as None() with only one argument. It is
// provided as a more readable function when only one argument is
// needed.
//
//   td.Cmp(t, 12, td.Not(10)) // succeeds
//   td.Cmp(t, 12, td.Not(12)) // fails
func Not(notExpected interface{}) TestDeep {
	return &tdNone{
		tdList: newList(notExpected),
	}
}

func (n *tdNone) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	for idx, item := range n.items {
<<<<<<< HEAD
		if deepValueEqualOK(got, item) {
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
		if deepValueEqualOK(got, item) {
=======
		if deepValueEqualFinalOK(ctx, got, item) {
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 6b7ce455e (update vendored files)
=======
//
// Note Flatten function can be used to group or reuse some values or
// operators and so avoid boring and inefficient copies:
//
//   prime := td.Flatten([]int{1, 2, 3, 5, 7, 11, 13})
//   even := td.Flatten([]int{2, 4, 6, 8, 10, 12, 14})
//   td.Cmp(t, 9, td.None(prime, even)) // succeeds
>>>>>>> 6b7ce455e (update vendored files)
func None(notExpectedValues ...interface{}) TestDeep {
	return &tdNone{
		tdList: newList(notExpectedValues...),
	}
}

// summary(Not): value must not match
// input(Not): all

// Not operator compares data against the not expected value. During a
// match, it must not match to succeed.
//
// Not is the same operator as None() with only one argument. It is
// provided as a more readable function when only one argument is
// needed.
//
//   td.Cmp(t, 12, td.Not(10)) // succeeds
//   td.Cmp(t, 12, td.Not(12)) // fails
func Not(notExpected interface{}) TestDeep {
	return &tdNone{
		tdList: newList(notExpected),
	}
}

func (n *tdNone) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	for idx, item := range n.items {
<<<<<<< HEAD
		if deepValueEqualOK(got, item) {
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
		if deepValueEqualOK(got, item) {
=======
		if deepValueEqualFinalOK(ctx, got, item) {
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 4d7e5ad26 (update vendored files)
=======
//
// Note Flatten function can be used to group or reuse some values or
// operators and so avoid boring and inefficient copies:
//
//   prime := td.Flatten([]int{1, 2, 3, 5, 7, 11, 13})
//   even := td.Flatten([]int{2, 4, 6, 8, 10, 12, 14})
//   td.Cmp(t, 9, td.None(prime, even)) // succeeds
>>>>>>> 4d7e5ad26 (update vendored files)
func None(notExpectedValues ...interface{}) TestDeep {
	return &tdNone{
		tdList: newList(notExpectedValues...),
	}
}

// summary(Not): value must not match
// input(Not): all

// Not operator compares data against the not expected value. During a
// match, it must not match to succeed.
//
// Not is the same operator as None() with only one argument. It is
// provided as a more readable function when only one argument is
// needed.
//
//   td.Cmp(t, 12, td.Not(10)) // succeeds
//   td.Cmp(t, 12, td.Not(12)) // fails
func Not(notExpected interface{}) TestDeep {
	return &tdNone{
		tdList: newList(notExpected),
	}
}

func (n *tdNone) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	for idx, item := range n.items {
<<<<<<< HEAD
		if deepValueEqualOK(got, item) {
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
		if deepValueEqualOK(got, item) {
=======
		if deepValueEqualFinalOK(ctx, got, item) {
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
//   td.Cmp(t, 12, td.None(8, 10, 14))     // succeeds
//   td.Cmp(t, 12, td.None(8, 10, 12, 14)) // fails
func None(notExpectedValues ...interface{}) TestDeep {
	return &tdNone{
		tdList: newList(notExpectedValues...),
	}
}

// summary(Not): value must not match
// input(Not): all

// Not operator compares data against the not expected value. During a
// match, it must not match to succeed.
//
// Not is the same operator as None() with only one argument. It is
// provided as a more readable function when only one argument is
// needed.
//
//   td.Cmp(t, 12, td.Not(10)) // succeeds
//   td.Cmp(t, 12, td.Not(12)) // fails
func Not(notExpected interface{}) TestDeep {
	return &tdNone{
		tdList: newList(notExpected),
	}
}

func (n *tdNone) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	for idx, item := range n.items {
		if deepValueEqualOK(got, item) {
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
			if ctx.BooleanError {
				return ctxerr.BooleanError
			}

			var mesg string
			if n.GetLocation().Func == "Not" {
				mesg = "comparing with Not"
			} else {
				mesg = fmt.Sprintf("comparing with None (part %d of %d is OK)",
					idx+1, len(n.items))
			}
			return ctx.CollectError(&ctxerr.Error{
				Message:  mesg,
				Got:      got,
				Expected: n,
			})
		}
	}
	return nil
}
