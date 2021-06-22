// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
//go:build (arm || mips || mipsle || 386 || ppc) && linux
// +build arm mips mipsle 386 ppc
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// +build arm mips mipsle 386
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
// +build arm mips mipsle 386
=======
//go:build (arm || mips || mipsle || 386 || ppc) && linux
// +build arm mips mipsle 386 ppc
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// +build arm mips mipsle 386
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
// +build arm mips mipsle 386
=======
//go:build (arm || mips || mipsle || 386 || ppc) && linux
// +build arm mips mipsle 386 ppc
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// +build arm mips mipsle 386
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
// +build arm mips mipsle 386
=======
//go:build (arm || mips || mipsle || 386 || ppc) && linux
// +build arm mips mipsle 386 ppc
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// +build arm mips mipsle 386
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
// +build linux

package socket

import "unsafe"

func (h *msghdr) setIov(vs []iovec) {
	l := len(vs)
	if l == 0 {
		return
	}
	h.Iov = &vs[0]
	h.Iovlen = uint32(l)
}

func (h *msghdr) setControl(b []byte) {
	h.Control = (*byte)(unsafe.Pointer(&b[0]))
	h.Controllen = uint32(len(b))
}
