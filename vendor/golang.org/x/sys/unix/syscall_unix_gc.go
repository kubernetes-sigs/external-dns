// Copyright 2016 The Go Authors. All rights reserved.
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
<<<<<<< HEAD
//go:build (darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris) && gc && !ppc64le && !ppc64
// +build darwin dragonfly freebsd linux netbsd openbsd solaris
||||||| parent of e1cd8261c (UPSTREAM: <carry>: update vendored files v0.13.1)
//go:build (darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris) && gc && !ppc64le && !ppc64
// +build darwin dragonfly freebsd linux netbsd openbsd solaris
=======
//go:build (darwin || dragonfly || freebsd || (linux && !ppc64 && !ppc64le) || netbsd || openbsd || solaris) && gc
// +build darwin dragonfly freebsd linux,!ppc64,!ppc64le netbsd openbsd solaris
>>>>>>> e1cd8261c (UPSTREAM: <carry>: update vendored files v0.13.1)
// +build gc
<<<<<<< HEAD
// +build !ppc64le
// +build !ppc64
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 5ce8c7613 (update vendored files)
=======
//go:build (darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris) && gc && !ppc64le && !ppc64
>>>>>>> 5ce8c7613 (update vendored files)
// +build darwin dragonfly freebsd linux netbsd openbsd solaris
<<<<<<< HEAD
// +build gc,!ppc64le,!ppc64
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
// +build gc,!ppc64le,!ppc64
=======
// +build gc
// +build !ppc64le
// +build !ppc64
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 6b7ce455e (update vendored files)
=======
//go:build (darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris) && gc && !ppc64le && !ppc64
>>>>>>> 6b7ce455e (update vendored files)
// +build darwin dragonfly freebsd linux netbsd openbsd solaris
<<<<<<< HEAD
// +build gc,!ppc64le,!ppc64
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
// +build gc,!ppc64le,!ppc64
=======
// +build gc
// +build !ppc64le
// +build !ppc64
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 4d7e5ad26 (update vendored files)
=======
//go:build (darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris) && gc && !ppc64le && !ppc64
>>>>>>> 4d7e5ad26 (update vendored files)
// +build darwin dragonfly freebsd linux netbsd openbsd solaris
<<<<<<< HEAD
// +build gc,!ppc64le,!ppc64
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
// +build gc,!ppc64le,!ppc64
=======
// +build gc
// +build !ppc64le
// +build !ppc64
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of e1cd8261c (UPSTREAM: <carry>: update vendored files v0.13.1)
// +build !ppc64le
// +build !ppc64
=======
>>>>>>> e1cd8261c (UPSTREAM: <carry>: update vendored files v0.13.1)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// +build darwin dragonfly freebsd linux netbsd openbsd solaris
// +build gc,!ppc64le,!ppc64
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
// +build darwin dragonfly freebsd linux netbsd openbsd solaris
// +build gc,!ppc64le,!ppc64
=======
//go:build (darwin || dragonfly || freebsd || (linux && !ppc64 && !ppc64le) || netbsd || openbsd || solaris) && gc
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)

package unix

import "syscall"

func Syscall(trap, a1, a2, a3 uintptr) (r1, r2 uintptr, err syscall.Errno)
func Syscall6(trap, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err syscall.Errno)
func RawSyscall(trap, a1, a2, a3 uintptr) (r1, r2 uintptr, err syscall.Errno)
func RawSyscall6(trap, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err syscall.Errno)
