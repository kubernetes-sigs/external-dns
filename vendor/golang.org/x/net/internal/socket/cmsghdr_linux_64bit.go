// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
//go:build (arm64 || amd64 || ppc64 || ppc64le || mips64 || mips64le || riscv64 || s390x) && linux
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
=======
//go:build (arm64 || amd64 || ppc64 || ppc64le || mips64 || mips64le || riscv64 || s390x) && linux
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
// +build arm64 amd64 ppc64 ppc64le mips64 mips64le riscv64 s390x
||||||| parent of 6b7ce455e (update vendored files)
// +build arm64 amd64 ppc64 ppc64le mips64 mips64le riscv64 s390x
=======
//go:build (arm64 || amd64 || loong64 || ppc64 || ppc64le || mips64 || mips64le || riscv64 || s390x) && linux
// +build arm64 amd64 loong64 ppc64 ppc64le mips64 mips64le riscv64 s390x
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// +build arm64 amd64 ppc64 ppc64le mips64 mips64le riscv64 s390x
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
// +build arm64 amd64 ppc64 ppc64le mips64 mips64le riscv64 s390x
=======
//go:build (arm64 || amd64 || loong64 || ppc64 || ppc64le || mips64 || mips64le || riscv64 || s390x) && linux
// +build arm64 amd64 loong64 ppc64 ppc64le mips64 mips64le riscv64 s390x
>>>>>>> 4d7e5ad26 (update vendored files)
// +build linux

package socket

func (h *cmsghdr) set(l, lvl, typ int) {
	h.Len = uint64(l)
	h.Level = int32(lvl)
	h.Type = int32(typ)
}
