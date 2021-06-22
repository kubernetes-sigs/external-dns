// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
//go:build dragonfly || freebsd || (linux && !s390x && !386) || netbsd || openbsd
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
=======
//go:build dragonfly || freebsd || (linux && !s390x && !386) || netbsd || openbsd
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
// +build dragonfly freebsd linux,!s390x,!386 netbsd openbsd
||||||| parent of 6b7ce455e (update vendored files)
// +build dragonfly freebsd linux,!s390x,!386 netbsd openbsd
=======
//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris
>>>>>>> 6b7ce455e (update vendored files)

package socket

import (
	"syscall"
	"unsafe"
)

//go:linkname syscall_getsockopt syscall.getsockopt
func syscall_getsockopt(s, level, name int, val unsafe.Pointer, vallen *uint32) error

//go:linkname syscall_setsockopt syscall.setsockopt
func syscall_setsockopt(s, level, name int, val unsafe.Pointer, vallen uintptr) error

//go:linkname syscall_recvmsg syscall.recvmsg
func syscall_recvmsg(s int, msg *syscall.Msghdr, flags int) (int, error)

//go:linkname syscall_sendmsg syscall.sendmsg
func syscall_sendmsg(s int, msg *syscall.Msghdr, flags int) (int, error)

func getsockopt(s uintptr, level, name int, b []byte) (int, error) {
	l := uint32(len(b))
	err := syscall_getsockopt(int(s), level, name, unsafe.Pointer(&b[0]), &l)
	return int(l), err
}

func setsockopt(s uintptr, level, name int, b []byte) error {
	return syscall_setsockopt(int(s), level, name, unsafe.Pointer(&b[0]), uintptr(len(b)))
}

func recvmsg(s uintptr, h *msghdr, flags int) (int, error) {
	return syscall_recvmsg(int(s), (*syscall.Msghdr)(unsafe.Pointer(h)), flags)
}

func sendmsg(s uintptr, h *msghdr, flags int) (int, error) {
	return syscall_sendmsg(int(s), (*syscall.Msghdr)(unsafe.Pointer(h)), flags)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// +build dragonfly freebsd linux,!s390x,!386 netbsd openbsd

package socket

import (
	"syscall"
	"unsafe"
)

func getsockopt(s uintptr, level, name int, b []byte) (int, error) {
	l := uint32(len(b))
	_, _, errno := syscall.Syscall6(syscall.SYS_GETSOCKOPT, s, uintptr(level), uintptr(name), uintptr(unsafe.Pointer(&b[0])), uintptr(unsafe.Pointer(&l)), 0)
	return int(l), errnoErr(errno)
}

func setsockopt(s uintptr, level, name int, b []byte) error {
	_, _, errno := syscall.Syscall6(syscall.SYS_SETSOCKOPT, s, uintptr(level), uintptr(name), uintptr(unsafe.Pointer(&b[0])), uintptr(len(b)), 0)
	return errnoErr(errno)
}

func recvmsg(s uintptr, h *msghdr, flags int) (int, error) {
	n, _, errno := syscall.Syscall(syscall.SYS_RECVMSG, s, uintptr(unsafe.Pointer(h)), uintptr(flags))
	return int(n), errnoErr(errno)
}

func sendmsg(s uintptr, h *msghdr, flags int) (int, error) {
	n, _, errno := syscall.Syscall(syscall.SYS_SENDMSG, s, uintptr(unsafe.Pointer(h)), uintptr(flags))
	return int(n), errnoErr(errno)
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
}
