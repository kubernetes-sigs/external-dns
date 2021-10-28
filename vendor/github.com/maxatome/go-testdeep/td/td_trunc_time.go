// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

import (
	"fmt"
	"reflect"
	"time"

	"github.com/maxatome/go-testdeep/internal/ctxerr"
	"github.com/maxatome/go-testdeep/internal/types"
)

type tdTruncTime struct {
	tdExpectedType
	expectedTime time.Time
	trunc        time.Duration
}

var _ TestDeep = &tdTruncTime{}

// summary(TruncTime): compares time.Time (or assignable) values after
// truncating them
// input(TruncTime): struct(time.Time),ptr(todo)

// TruncTime operator compares time.Time (or assignable) values after
// truncating them to the optional "trunc" duration. See time.Truncate
// for details about the truncation.
//
// If "trunc" is missing, it defaults to 0.
//
// During comparison, location does not matter as time.Equal method is
// used behind the scenes: a time instant in two different locations
// is the same time instant.
//
// Whatever the "trunc" value is, the monotonic clock is stripped
// before the comparison against "expectedTime".
//
//   gotDate := time.Date(2018, time.March, 9, 1, 2, 3, 999999999, time.UTC).
//     In(time.FixedZone("UTC+2", 2))
//
//   expected := time.Date(2018, time.March, 9, 1, 2, 3, 0, time.UTC)
//
//   td.Cmp(t, gotDate, td.TruncTime(expected))              // fails, ns differ
//   td.Cmp(t, gotDate, td.TruncTime(expected, time.Second)) // succeeds
//
// TypeBehind method returns the reflect.Type of "expectedTime".
func TruncTime(expectedTime interface{}, trunc ...time.Duration) TestDeep {
	const usage = "(time.Time[, time.Duration])"

	t := tdTruncTime{
		tdExpectedType: tdExpectedType{
			base: newBase(3),
		},
	}

	if len(trunc) > 1 {
		t.err = ctxerr.OpTooManyParams("TruncTime", usage)
		return &t
	}

	if len(trunc) == 1 {
		t.trunc = trunc[0]
	}

	vval := reflect.ValueOf(expectedTime)

	t.expectedType = vval.Type()
	if t.expectedType == types.Time {
		t.expectedTime = expectedTime.(time.Time).Truncate(t.trunc)
		return &t
	}
	if !t.expectedType.ConvertibleTo(types.Time) {
		t.err = ctxerr.OpBad("TruncTime", "usage: TruncTime%s, 1st parameter must be time.Time or convertible to time.Time, but not %T",
			usage, expectedTime)
		return &t
	}

	t.expectedTime = vval.Convert(types.Time).
		Interface().(time.Time).Truncate(t.trunc)
	return &t
}

func (t *tdTruncTime) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	if t.err != nil {
		return ctx.CollectError(t.err)
	}

	err := t.checkType(ctx, got)
	if err != nil {
		return ctx.CollectError(err)
	}

	gotTime, err := getTime(ctx, got, got.Type() != types.Time)
	if err != nil {
		return ctx.CollectError(err)
	}
	gotTimeTrunc := gotTime.Truncate(t.trunc)

	if gotTimeTrunc.Equal(t.expectedTime) {
		return nil
	}

	// Fail
	if ctx.BooleanError {
		return ctxerr.BooleanError
	}

	var gotRawStr, gotTruncStr string
	if t.expectedType != types.Time &&
		t.expectedType.Implements(types.FmtStringer) {
		gotRawStr = got.Interface().(fmt.Stringer).String()
		gotTruncStr = reflect.ValueOf(gotTimeTrunc).Convert(t.expectedType).
			Interface().(fmt.Stringer).String()
	} else {
		gotRawStr = gotTime.String()
		gotTruncStr = gotTimeTrunc.String()
	}

	return ctx.CollectError(&ctxerr.Error{
		Message:  "values differ",
		Got:      types.RawString(gotRawStr + "\ntruncated to:\n" + gotTruncStr),
		Expected: t,
	})
}

func (t *tdTruncTime) String() string {
	if t.err != nil {
		return t.stringError()
	}

	if t.expectedType.Implements(types.FmtStringer) {
		return reflect.ValueOf(t.expectedTime).Convert(t.expectedType).
			Interface().(fmt.Stringer).String()
	}
	return t.expectedTime.String()
}
