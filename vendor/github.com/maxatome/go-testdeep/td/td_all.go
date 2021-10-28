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

type tdAll struct {
	tdList
}

var _ TestDeep = &tdAll{}

// summary(All): all expected values have to match
// input(All): all

// All operator compares data against several expected values. During
// a match, all of them have to match to succeed. Consider it
// as a "AND" logical operator.
//
//   td.Cmp(t, "foobar", td.All(
//     td.Len(6),
//     td.HasPrefix("fo"),
//     td.HasSuffix("ar"),
//   )) // succeeds
//
// Note Flatten function can be used to group or reuse some values or
// operators and so avoid boring and inefficient copies:
//
//   stringOps := td.Flatten([]td.TestDeep{td.HasPrefix("fo"), td.HasSuffix("ar")})
//   td.Cmp(t, "foobar", td.All(
//     td.Len(6),
//     stringOps,
//   )) // succeeds
//
// One can do the same with All operator itself:
//
//   stringOps := td.All(td.HasPrefix("fo"), td.HasSuffix("ar"))
//   td.Cmp(t, "foobar", td.All(
//     td.Len(6),
//     stringOps,
//   )) // succeeds
//
// but if an error occurs in the nested All, the report is a bit more
// complex to read due to the nested level. Flatten does not create a
// new level, its slice is just flattened in the All parameters.
//
// TypeBehind method can return a non-nil reflect.Type if all items
// known non-interface types are equal, or if only interface types
// are found (mostly issued from Isa()) and they are equal.
func All(expectedValues ...interface{}) TestDeep {
	return &tdAll{
		tdList: newList(expectedValues...),
	}
}

func (a *tdAll) Match(ctx ctxerr.Context, got reflect.Value) (err *ctxerr.Error) {
	var origErr *ctxerr.Error
	for idx, item := range a.items {
		// Use deepValueEqualFinal here instead of deepValueEqual as we
		// want to know whether an error occurred or not, we do not want
		// to accumulate it silently
		origErr = deepValueEqualFinal(
			ctx.ResetErrors().
				AddCustomLevel(fmt.Sprintf("<All#%d/%d>", idx+1, len(a.items))),
			got, item)
		if origErr != nil {
			if ctx.BooleanError {
				return ctxerr.BooleanError
			}
			err := &ctxerr.Error{
				Message:  fmt.Sprintf("compared (part %d of %d)", idx+1, len(a.items)),
				Got:      got,
				Expected: item,
			}
			if item.IsValid() && item.Type().Implements(testDeeper) {
				err.Origin = origErr
			}
			return ctx.CollectError(err)
		}
	}
	return nil
}

func (a *tdAll) TypeBehind() reflect.Type {
	return uniqTypeBehindSlice(a.items)
}
