// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
//go:build !linux
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
=======
//go:build !linux
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
=======
//go:build !linux
>>>>>>> 6b7ce455e (update vendored files)
// +build !linux

package ipv4

const sizeofICMPFilter = 0x0

type icmpFilter struct {
}

func (f *icmpFilter) accept(typ ICMPType) {
}

func (f *icmpFilter) block(typ ICMPType) {
}

func (f *icmpFilter) setAll(block bool) {
}

func (f *icmpFilter) willBlock(typ ICMPType) bool {
	return false
}
