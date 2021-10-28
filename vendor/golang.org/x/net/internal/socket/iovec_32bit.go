// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
//go:build (arm || mips || mipsle || 386 || ppc) && (darwin || dragonfly || freebsd || linux || netbsd || openbsd)
// +build arm mips mipsle 386 ppc
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// +build arm mips mipsle 386
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
// +build arm mips mipsle 386
=======
//go:build (arm || mips || mipsle || 386 || ppc) && (darwin || dragonfly || freebsd || linux || netbsd || openbsd)
// +build arm mips mipsle 386 ppc
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// +build arm mips mipsle 386
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
// +build arm mips mipsle 386
=======
//go:build (arm || mips || mipsle || 386 || ppc) && (darwin || dragonfly || freebsd || linux || netbsd || openbsd)
// +build arm mips mipsle 386 ppc
>>>>>>> 6b7ce455e (update vendored files)
// +build darwin dragonfly freebsd linux netbsd openbsd

package socket

import "unsafe"

func (v *iovec) set(b []byte) {
	l := len(b)
	if l == 0 {
		return
	}
	v.Base = (*byte)(unsafe.Pointer(&b[0]))
	v.Len = uint32(l)
}
