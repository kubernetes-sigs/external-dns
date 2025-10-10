// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
//go:build !purego && !appengine
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
//go:build !purego && !appengine
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
// +build !purego,!appengine

||||||| parent of c5487e6d6 (NE-2142: UPSTREAM: 5739: Bump k8s and controller-runtime modules)
//go:build !purego && !appengine
// +build !purego,!appengine

=======
>>>>>>> c5487e6d6 (NE-2142: UPSTREAM: 5739: Bump k8s and controller-runtime modules)
package impl

// When using unsafe pointers, we can just treat enum values as int32s.

var (
	coderEnumNoZero      = coderInt32NoZero
	coderEnum            = coderInt32
	coderEnumPtr         = coderInt32Ptr
	coderEnumSlice       = coderInt32Slice
	coderEnumPackedSlice = coderInt32PackedSlice
)
