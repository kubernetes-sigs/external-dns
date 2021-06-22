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
