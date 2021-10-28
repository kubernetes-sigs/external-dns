// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

import (
	"reflect"
	"time"

	"github.com/maxatome/go-testdeep/internal/ctxerr"
	"github.com/maxatome/go-testdeep/internal/dark"
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	"github.com/maxatome/go-testdeep/internal/types"
)

// getTime returns the time.Time that is inside got or that can be
// converted from got contents.
func getTime(ctx ctxerr.Context, got reflect.Value, mustConvert bool) (time.Time, *ctxerr.Error) {
	var (
		gotIf interface{}
		ok    bool
	)
	if mustConvert {
		gotIf, ok = dark.GetInterface(got.Convert(types.Time), true)
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 5ce8c7613 (update vendored files)
=======
	"github.com/maxatome/go-testdeep/internal/types"
>>>>>>> 5ce8c7613 (update vendored files)
)

// getTime returns the time.Time that is inside got or that can be
// converted from got contents.
func getTime(ctx ctxerr.Context, got reflect.Value, mustConvert bool) (time.Time, *ctxerr.Error) {
	var (
		gotIf interface{}
		ok    bool
	)
	if mustConvert {
<<<<<<< HEAD
		gotIf, ok = dark.GetInterface(got.Convert(timeType), true)
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
		gotIf, ok = dark.GetInterface(got.Convert(timeType), true)
=======
		gotIf, ok = dark.GetInterface(got.Convert(types.Time), true)
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 6b7ce455e (update vendored files)
=======
	"github.com/maxatome/go-testdeep/internal/types"
>>>>>>> 6b7ce455e (update vendored files)
)

// getTime returns the time.Time that is inside got or that can be
// converted from got contents.
func getTime(ctx ctxerr.Context, got reflect.Value, mustConvert bool) (time.Time, *ctxerr.Error) {
	var (
		gotIf interface{}
		ok    bool
	)
	if mustConvert {
<<<<<<< HEAD
		gotIf, ok = dark.GetInterface(got.Convert(timeType), true)
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
		gotIf, ok = dark.GetInterface(got.Convert(timeType), true)
=======
		gotIf, ok = dark.GetInterface(got.Convert(types.Time), true)
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 4d7e5ad26 (update vendored files)
=======
	"github.com/maxatome/go-testdeep/internal/types"
>>>>>>> 4d7e5ad26 (update vendored files)
)

// getTime returns the time.Time that is inside got or that can be
// converted from got contents.
func getTime(ctx ctxerr.Context, got reflect.Value, mustConvert bool) (time.Time, *ctxerr.Error) {
	var (
		gotIf interface{}
		ok    bool
	)
	if mustConvert {
<<<<<<< HEAD
		gotIf, ok = dark.GetInterface(got.Convert(timeType), true)
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
		gotIf, ok = dark.GetInterface(got.Convert(timeType), true)
=======
		gotIf, ok = dark.GetInterface(got.Convert(types.Time), true)
>>>>>>> 4d7e5ad26 (update vendored files)
	} else {
		gotIf, ok = dark.GetInterface(got, true)
	}
	if !ok {
		return time.Time{}, ctx.CannotCompareError()
	}
	return gotIf.(time.Time), nil
}
