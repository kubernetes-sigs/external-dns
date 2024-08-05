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
<<<<<<< HEAD
<<<<<<< HEAD
// comparing a slice with [Slice] and wanting to ignore some indexes,
// for example (if you don't want to use [SuperSliceOf]). Or comparing
// a struct with [SStruct] and wanting to ignore some fields:
//
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
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
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
//   td.Cmp(t, td.SStruct(
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
//   td.Cmp(t, td.SStruct(
=======
//   td.Cmp(t, got, td.SStruct(
>>>>>>> 4d7e5ad26 (update vendored files)
//     Person{
//       Name: "John Doe",
//     },
//     td.StructFields{
//       Age:      td.Between(40, 45),
//       Children: td.Ignore(),
//     }),
//   )
||||||| parent of e1cd8261c (UPSTREAM: <carry>: update vendored files v0.13.1)
//   td.Cmp(t, got, td.SStruct(
//     Person{
//       Name: "John Doe",
//     },
//     td.StructFields{
//       Age:      td.Between(40, 45),
//       Children: td.Ignore(),
//     }),
//   )
=======
//	td.Cmp(t, got, td.SStruct(
//	  Person{
//	    Name: "John Doe",
//	  },
//	  td.StructFields{
//	    Age:      td.Between(40, 45),
//	    Children: td.Ignore(),
//	  }),
//	)
>>>>>>> e1cd8261c (UPSTREAM: <carry>: update vendored files v0.13.1)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// comparing a slice with Slice and wanting to ignore some indexes,
// for example. Or comparing a struct with SStruct and wanting to
// ignore some fields:
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
// comparing a slice with Slice and wanting to ignore some indexes,
// for example. Or comparing a struct with SStruct and wanting to
// ignore some fields:
=======
// comparing a slice with [Slice] and wanting to ignore some indexes,
// for example (if you don't want to use [SuperSliceOf]). Or comparing
// a struct with [SStruct] and wanting to ignore some fields:
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
//
<<<<<<< HEAD
//   td.Cmp(t, td.SStruct(
//     Person{
//       Name: "John Doe",
//     },
//     td.StructFields{
//       Age:      td.Between(40, 45),
//       Children: td.Ignore(),
//     }),
//   )
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
//   td.Cmp(t, td.SStruct(
//     Person{
//       Name: "John Doe",
//     },
//     td.StructFields{
//       Age:      td.Between(40, 45),
//       Children: td.Ignore(),
//     }),
//   )
=======
//	td.Cmp(t, got, td.SStruct(
//	  Person{
//	    Name: "John Doe",
//	  },
//	  td.StructFields{
//	    Age:      td.Between(40, 45),
//	    Children: td.Ignore(),
//	  }),
//	)
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
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
