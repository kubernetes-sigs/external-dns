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
||||||| parent of 5ce8c7613 (update vendored files)
=======
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 6b7ce455e (update vendored files)
=======
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4d7e5ad26 (update vendored files)
=======
>>>>>>> 4d7e5ad26 (update vendored files)
//go:build !aix && !darwin && !dragonfly && !freebsd && !linux && !netbsd && !openbsd && !solaris && !windows && !zos
// +build !aix,!darwin,!dragonfly,!freebsd,!linux,!netbsd,!openbsd,!solaris,!windows,!zos

package socket

import "net"

const (
	sysAF_UNSPEC = 0x0
	sysAF_INET   = 0x2
	sysAF_INET6  = 0xa

	sysSOCK_RAW = 0x3

	sizeofSockaddrInet4 = 0x10
	sizeofSockaddrInet6 = 0x1c
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// +build !aix,!darwin,!dragonfly,!freebsd,!linux,!netbsd,!openbsd,!solaris,!windows,!zos

package socket

import "net"

const (
	sysAF_UNSPEC = 0x0
	sysAF_INET   = 0x2
	sysAF_INET6  = 0xa

	sysSOCK_RAW = 0x3
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
=======
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// +build !aix,!darwin,!dragonfly,!freebsd,!linux,!netbsd,!openbsd,!solaris,!windows,!zos

package socket

import "net"

const (
	sysAF_UNSPEC = 0x0
	sysAF_INET   = 0x2
	sysAF_INET6  = 0xa

	sysSOCK_RAW = 0x3
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
=======
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// +build !aix,!darwin,!dragonfly,!freebsd,!linux,!netbsd,!openbsd,!solaris,!windows,!zos

package socket

import "net"

const (
	sysAF_UNSPEC = 0x0
	sysAF_INET   = 0x2
	sysAF_INET6  = 0xa

	sysSOCK_RAW = 0x3
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
=======
>>>>>>> 4d7e5ad26 (update vendored files)
)

func marshalInetAddr(ip net.IP, port int, zone string) []byte {
	return nil
}

func parseInetAddr(b []byte, network string) (net.Addr, error) {
	return nil, errNotImplemented
}

func getsockopt(s uintptr, level, name int, b []byte) (int, error) {
	return 0, errNotImplemented
}

func setsockopt(s uintptr, level, name int, b []byte) error {
	return errNotImplemented
}

func recvmsg(s uintptr, buffers [][]byte, oob []byte, flags int, network string) (n, oobn int, recvflags int, from net.Addr, err error) {
	return 0, 0, 0, nil, errNotImplemented
}

func sendmsg(s uintptr, buffers [][]byte, oob []byte, to net.Addr, flags int) (int, error) {
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// +build !aix,!darwin,!dragonfly,!freebsd,!linux,!netbsd,!openbsd,!solaris,!windows,!zos

package socket

import "net"

const (
	sysAF_UNSPEC = 0x0
	sysAF_INET   = 0x2
	sysAF_INET6  = 0xa

	sysSOCK_RAW = 0x3
)

func marshalInetAddr(ip net.IP, port int, zone string) []byte {
	return nil
}

func parseInetAddr(b []byte, network string) (net.Addr, error) {
	return nil, errNotImplemented
}

func getsockopt(s uintptr, level, name int, b []byte) (int, error) {
	return 0, errNotImplemented
}

func setsockopt(s uintptr, level, name int, b []byte) error {
	return errNotImplemented
}

func recvmsg(s uintptr, h *msghdr, flags int) (int, error) {
	return 0, errNotImplemented
}

func sendmsg(s uintptr, h *msghdr, flags int) (int, error) {
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
	return 0, errNotImplemented
}

func recvmmsg(s uintptr, hs []mmsghdr, flags int) (int, error) {
	return 0, errNotImplemented
}

func sendmmsg(s uintptr, hs []mmsghdr, flags int) (int, error) {
	return 0, errNotImplemented
}
