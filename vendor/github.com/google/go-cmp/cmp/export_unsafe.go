// Copyright 2017, The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
// license that can be found in the LICENSE file.
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// license that can be found in the LICENSE.md file.
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
// license that can be found in the LICENSE.md file.
=======
// license that can be found in the LICENSE file.
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// license that can be found in the LICENSE.md file.
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
// license that can be found in the LICENSE.md file.
=======
// license that can be found in the LICENSE file.
>>>>>>> 6b7ce455e (update vendored files)

//go:build !purego
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// license that can be found in the LICENSE.md file.
||||||| parent of 4d7e5ad26 (update vendored files)
// license that can be found in the LICENSE.md file.
=======
// license that can be found in the LICENSE file.
>>>>>>> 4d7e5ad26 (update vendored files)

<<<<<<< HEAD
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
=======
//go:build !purego
>>>>>>> 4d7e5ad26 (update vendored files)
// +build !purego

package cmp

import (
	"reflect"
	"unsafe"
)

const supportExporters = true

// retrieveUnexportedField uses unsafe to forcibly retrieve any field from
// a struct such that the value has read-write permissions.
//
// The parent struct, v, must be addressable, while f must be a StructField
// describing the field to retrieve. If addr is false,
// then the returned value will be shallowed copied to be non-addressable.
func retrieveUnexportedField(v reflect.Value, f reflect.StructField, addr bool) reflect.Value {
	ve := reflect.NewAt(f.Type, unsafe.Pointer(uintptr(unsafe.Pointer(v.UnsafeAddr()))+f.Offset)).Elem()
	if !addr {
		// A field is addressable if and only if the struct is addressable.
		// If the original parent value was not addressable, shallow copy the
		// value to make it non-addressable to avoid leaking an implementation
		// detail of how forcibly exporting a field works.
		if ve.Kind() == reflect.Interface && ve.IsNil() {
			return reflect.Zero(f.Type)
		}
		return reflect.ValueOf(ve.Interface()).Convert(f.Type)
	}
	return ve
}
