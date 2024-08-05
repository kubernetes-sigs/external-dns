// Copyright 2019 The Go Authors. All rights reserved.
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
//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || zos
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris zos

package socket

import "golang.org/x/sys/unix"

const (
	sysAF_UNSPEC = unix.AF_UNSPEC
	sysAF_INET   = unix.AF_INET
	sysAF_INET6  = unix.AF_INET6

	sysSOCK_RAW = unix.SOCK_RAW

	sizeofSockaddrInet4 = unix.SizeofSockaddrInet4
	sizeofSockaddrInet6 = unix.SizeofSockaddrInet6
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris
||||||| parent of 5ce8c7613 (update vendored files)
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris
=======
//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || zos
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris zos
>>>>>>> 5ce8c7613 (update vendored files)

package socket

import "golang.org/x/sys/unix"

const (
	sysAF_UNSPEC = unix.AF_UNSPEC
	sysAF_INET   = unix.AF_INET
	sysAF_INET6  = unix.AF_INET6

	sysSOCK_RAW = unix.SOCK_RAW
<<<<<<< HEAD
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
=======

	sizeofSockaddrInet4 = unix.SizeofSockaddrInet4
	sizeofSockaddrInet6 = unix.SizeofSockaddrInet6
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris
||||||| parent of 6b7ce455e (update vendored files)
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris
=======
//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || zos
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris zos
>>>>>>> 6b7ce455e (update vendored files)

package socket

import "golang.org/x/sys/unix"

const (
	sysAF_UNSPEC = unix.AF_UNSPEC
	sysAF_INET   = unix.AF_INET
	sysAF_INET6  = unix.AF_INET6

	sysSOCK_RAW = unix.SOCK_RAW
<<<<<<< HEAD
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
=======

	sizeofSockaddrInet4 = unix.SizeofSockaddrInet4
	sizeofSockaddrInet6 = unix.SizeofSockaddrInet6
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris
||||||| parent of 4d7e5ad26 (update vendored files)
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris
=======
//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || zos
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris zos
>>>>>>> 4d7e5ad26 (update vendored files)

package socket

import "golang.org/x/sys/unix"

const (
	sysAF_UNSPEC = unix.AF_UNSPEC
	sysAF_INET   = unix.AF_INET
	sysAF_INET6  = unix.AF_INET6

	sysSOCK_RAW = unix.SOCK_RAW
<<<<<<< HEAD
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
=======

	sizeofSockaddrInet4 = unix.SizeofSockaddrInet4
	sizeofSockaddrInet6 = unix.SizeofSockaddrInet6
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris
=======
//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || zos
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)

package socket

import "golang.org/x/sys/unix"

const (
	sysAF_UNSPEC = unix.AF_UNSPEC
	sysAF_INET   = unix.AF_INET
	sysAF_INET6  = unix.AF_INET6

	sysSOCK_RAW = unix.SOCK_RAW
<<<<<<< HEAD
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======

	sizeofSockaddrInet4 = unix.SizeofSockaddrInet4
	sizeofSockaddrInet6 = unix.SizeofSockaddrInet6
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
)
