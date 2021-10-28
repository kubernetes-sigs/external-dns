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

type tdIgnore struct {
	baseOKNil
}

// summary(Ignore): allows to ignore a comparison
// input(Ignore): all

// Ignore operator is always true, whatever data is. It is useful when
// comparing a slice with Slice and wanting to ignore some indexes,
// for example. Or comparing a struct with SStruct and wanting to
// ignore some fields:
//
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
//   td.Cmp(t, got, td.SStruct(
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
//   td.Cmp(t, td.SStruct(
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
//   td.Cmp(t, td.SStruct(
=======
//   td.Cmp(t, got, td.SStruct(
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
//   td.Cmp(t, td.SStruct(
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
//   td.Cmp(t, td.SStruct(
=======
//   td.Cmp(t, got, td.SStruct(
>>>>>>> 6b7ce455e (update vendored files)
//     Person{
//       Name: "John Doe",
//     },
//     td.StructFields{
//       Age:      td.Between(40, 45),
//       Children: td.Ignore(),
//     }),
//   )
func Ignore() TestDeep {
	return &tdIgnore{
		baseOKNil: newBaseOKNil(3),
	}
}

func (i *tdIgnore) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	return nil
}

func (i *tdIgnore) String() string {
	return "Ignore()"
}
