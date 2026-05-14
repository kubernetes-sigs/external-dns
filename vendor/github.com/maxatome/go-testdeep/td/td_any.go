// Copyright (c) 2018-2025, Maxime Soulé
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
	tdListBase
}

var _ TestDeep = &tdAny{}

// summary(Any): at least one expected value have to match
// input(Any): all

// Any operator compares data against several expected values. During
// a match, at least one of them has to match to succeed. Consider it
// as a "OR" logical operator.
//
//	td.Cmp(t, "foo", td.Any("bar", "foo", "zip")) // succeeds
//	td.Cmp(t, "foo", td.Any(
//	  td.Len(4),
//	  td.HasPrefix("f"),
//	  td.HasSuffix("z"),
//	)) // succeeds coz "f" prefix
//
// Note [Flatten] function can be used to group or reuse some values or
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
		tdListBase: newListBase(expectedValues...),
	}
}

func (a *tdAny) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	for _, item := range a.items {
		ok, err := deepValueEqualFinalOK(ctx, got, item)
		if err != nil || ok {
			return err
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
}
