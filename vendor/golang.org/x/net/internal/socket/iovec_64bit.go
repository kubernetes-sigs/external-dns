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
<<<<<<< HEAD
//go:build (arm64 || amd64 || ppc64 || ppc64le || mips64 || mips64le || riscv64 || s390x) && (aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || zos)
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
=======
//go:build (arm64 || amd64 || ppc64 || ppc64le || mips64 || mips64le || riscv64 || s390x) && (aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || zos)
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
// +build arm64 amd64 ppc64 ppc64le mips64 mips64le riscv64 s390x
||||||| parent of 6b7ce455e (update vendored files)
// +build arm64 amd64 ppc64 ppc64le mips64 mips64le riscv64 s390x
=======
//go:build (arm64 || amd64 || loong64 || ppc64 || ppc64le || mips64 || mips64le || riscv64 || s390x) && (aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || zos)
// +build arm64 amd64 loong64 ppc64 ppc64le mips64 mips64le riscv64 s390x
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// +build arm64 amd64 ppc64 ppc64le mips64 mips64le riscv64 s390x
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
// +build arm64 amd64 ppc64 ppc64le mips64 mips64le riscv64 s390x
=======
//go:build (arm64 || amd64 || loong64 || ppc64 || ppc64le || mips64 || mips64le || riscv64 || s390x) && (aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || zos)
// +build arm64 amd64 loong64 ppc64 ppc64le mips64 mips64le riscv64 s390x
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// +build arm64 amd64 ppc64 ppc64le mips64 mips64le riscv64 s390x
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
// +build aix darwin dragonfly freebsd linux netbsd openbsd zos
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
// +build arm64 amd64 ppc64 ppc64le mips64 mips64le riscv64 s390x
// +build aix darwin dragonfly freebsd linux netbsd openbsd zos
=======
//go:build (arm64 || amd64 || loong64 || ppc64 || ppc64le || mips64 || mips64le || riscv64 || s390x) && (aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || zos)
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)

package socket

import "unsafe"

func (v *iovec) set(b []byte) {
	l := len(b)
	if l == 0 {
		return
	}
	v.Base = (*byte)(unsafe.Pointer(&b[0]))
	v.Len = uint64(l)
}
