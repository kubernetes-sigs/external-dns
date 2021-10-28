// Copyright 2018, The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
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

//go:build purego
// +build purego

package value

import "reflect"

// Pointer is an opaque typed pointer and is guaranteed to be comparable.
type Pointer struct {
	p uintptr
	t reflect.Type
}

// PointerOf returns a Pointer from v, which must be a
// reflect.Ptr, reflect.Slice, or reflect.Map.
func PointerOf(v reflect.Value) Pointer {
	// NOTE: Storing a pointer as an uintptr is technically incorrect as it
	// assumes that the GC implementation does not use a moving collector.
	return Pointer{v.Pointer(), v.Type()}
}

// IsNil reports whether the pointer is nil.
func (p Pointer) IsNil() bool {
	return p.p == 0
}

// Uintptr returns the pointer as a uintptr.
func (p Pointer) Uintptr() uintptr {
	return p.p
}
