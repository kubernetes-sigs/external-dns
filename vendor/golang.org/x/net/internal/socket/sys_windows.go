// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package socket

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

func probeProtocolStack() int {
	var p uintptr
	return int(unsafe.Sizeof(p))
}

const (
	sysAF_UNSPEC = windows.AF_UNSPEC
	sysAF_INET   = windows.AF_INET
	sysAF_INET6  = windows.AF_INET6

	sysSOCK_RAW = windows.SOCK_RAW
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD

	sizeofSockaddrInet4 = 0x10
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
)
||||||| parent of 5ce8c7613 (update vendored files)
)
=======
>>>>>>> 5ce8c7613 (update vendored files)

<<<<<<< HEAD
type sockaddrInet struct {
	Family uint16
	Port   uint16
	Addr   [4]byte /* in_addr */
	Zero   [8]uint8
}

type sockaddrInet6 struct {
	Family   uint16
	Port     uint16
	Flowinfo uint32
	Addr     [16]byte /* in6_addr */
	Scope_id uint32
}

const (
	sizeofSockaddrInet  = 0x10
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
type sockaddrInet struct {
	Family uint16
	Port   uint16
	Addr   [4]byte /* in_addr */
	Zero   [8]uint8
}

type sockaddrInet6 struct {
	Family   uint16
	Port     uint16
	Flowinfo uint32
	Addr     [16]byte /* in6_addr */
	Scope_id uint32
}

const (
	sizeofSockaddrInet  = 0x10
=======
	sizeofSockaddrInet4 = 0x10
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
)

type sockaddrInet struct {
	Family uint16
	Port   uint16
	Addr   [4]byte /* in_addr */
	Zero   [8]uint8
}

type sockaddrInet6 struct {
	Family   uint16
	Port     uint16
	Flowinfo uint32
	Addr     [16]byte /* in6_addr */
	Scope_id uint32
}

const (
	sizeofSockaddrInet  = 0x10
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
	sizeofSockaddrInet6 = 0x1c
)

func getsockopt(s uintptr, level, name int, b []byte) (int, error) {
	l := uint32(len(b))
	err := syscall.Getsockopt(syscall.Handle(s), int32(level), int32(name), (*byte)(unsafe.Pointer(&b[0])), (*int32)(unsafe.Pointer(&l)))
	return int(l), err
}

func setsockopt(s uintptr, level, name int, b []byte) error {
	return syscall.Setsockopt(syscall.Handle(s), int32(level), int32(name), (*byte)(unsafe.Pointer(&b[0])), int32(len(b)))
}

func recvmsg(s uintptr, h *msghdr, flags int) (int, error) {
	return 0, errNotImplemented
}

func sendmsg(s uintptr, h *msghdr, flags int) (int, error) {
	return 0, errNotImplemented
}

func recvmmsg(s uintptr, hs []mmsghdr, flags int) (int, error) {
	return 0, errNotImplemented
}

func sendmmsg(s uintptr, hs []mmsghdr, flags int) (int, error) {
	return 0, errNotImplemented
}
