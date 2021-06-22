// Copyright 2018, The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
<<<<<<< HEAD
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
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// license that can be found in the LICENSE.md file.

>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
// +build !purego

package value

import (
	"reflect"
	"unsafe"
)

// Pointer is an opaque typed pointer and is guaranteed to be comparable.
type Pointer struct {
	p unsafe.Pointer
	t reflect.Type
}

// PointerOf returns a Pointer from v, which must be a
// reflect.Ptr, reflect.Slice, or reflect.Map.
func PointerOf(v reflect.Value) Pointer {
	// The proper representation of a pointer is unsafe.Pointer,
	// which is necessary if the GC ever uses a moving collector.
	return Pointer{unsafe.Pointer(v.Pointer()), v.Type()}
}

// IsNil reports whether the pointer is nil.
func (p Pointer) IsNil() bool {
	return p.p == nil
}

// Uintptr returns the pointer as a uintptr.
func (p Pointer) Uintptr() uintptr {
	return uintptr(p.p)
}
