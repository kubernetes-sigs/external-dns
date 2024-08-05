// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package socket

import (
	"syscall"
	"unsafe"
)

const (
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	sysRECVMMSG = 0x13
	sysSENDMMSG = 0x14
)

func socketcall(call, a0, a1, a2, a3, a4, a5 uintptr) (uintptr, syscall.Errno)
func rawsocketcall(call, a0, a1, a2, a3, a4, a5 uintptr) (uintptr, syscall.Errno)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	sysSETSOCKOPT = 0xe
	sysGETSOCKOPT = 0xf
	sysSENDMSG    = 0x10
	sysRECVMSG    = 0x11
	sysRECVMMSG   = 0x13
	sysSENDMMSG   = 0x14
||||||| parent of 4d7e5ad26 (update vendored files)
	sysSETSOCKOPT = 0xe
	sysGETSOCKOPT = 0xf
	sysSENDMSG    = 0x10
	sysRECVMSG    = 0x11
	sysRECVMMSG   = 0x13
	sysSENDMMSG   = 0x14
=======
	sysRECVMMSG = 0x13
	sysSENDMMSG = 0x14
>>>>>>> 4d7e5ad26 (update vendored files)
)

func socketcall(call, a0, a1, a2, a3, a4, a5 uintptr) (uintptr, syscall.Errno)
func rawsocketcall(call, a0, a1, a2, a3, a4, a5 uintptr) (uintptr, syscall.Errno)

<<<<<<< HEAD
func getsockopt(s uintptr, level, name int, b []byte) (int, error) {
	l := uint32(len(b))
	_, errno := socketcall(sysGETSOCKOPT, s, uintptr(level), uintptr(name), uintptr(unsafe.Pointer(&b[0])), uintptr(unsafe.Pointer(&l)), 0)
	return int(l), errnoErr(errno)
}

func setsockopt(s uintptr, level, name int, b []byte) error {
	_, errno := socketcall(sysSETSOCKOPT, s, uintptr(level), uintptr(name), uintptr(unsafe.Pointer(&b[0])), uintptr(len(b)), 0)
	return errnoErr(errno)
}

func recvmsg(s uintptr, h *msghdr, flags int) (int, error) {
	n, errno := socketcall(sysRECVMSG, s, uintptr(unsafe.Pointer(h)), uintptr(flags), 0, 0, 0)
	return int(n), errnoErr(errno)
}

func sendmsg(s uintptr, h *msghdr, flags int) (int, error) {
	n, errno := socketcall(sysSENDMSG, s, uintptr(unsafe.Pointer(h)), uintptr(flags), 0, 0, 0)
	return int(n), errnoErr(errno)
}
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)

||||||| parent of 4d7e5ad26 (update vendored files)
func getsockopt(s uintptr, level, name int, b []byte) (int, error) {
	l := uint32(len(b))
	_, errno := socketcall(sysGETSOCKOPT, s, uintptr(level), uintptr(name), uintptr(unsafe.Pointer(&b[0])), uintptr(unsafe.Pointer(&l)), 0)
	return int(l), errnoErr(errno)
}

func setsockopt(s uintptr, level, name int, b []byte) error {
	_, errno := socketcall(sysSETSOCKOPT, s, uintptr(level), uintptr(name), uintptr(unsafe.Pointer(&b[0])), uintptr(len(b)), 0)
	return errnoErr(errno)
}

func recvmsg(s uintptr, h *msghdr, flags int) (int, error) {
	n, errno := socketcall(sysRECVMSG, s, uintptr(unsafe.Pointer(h)), uintptr(flags), 0, 0, 0)
	return int(n), errnoErr(errno)
}

func sendmsg(s uintptr, h *msghdr, flags int) (int, error) {
	n, errno := socketcall(sysSENDMSG, s, uintptr(unsafe.Pointer(h)), uintptr(flags), 0, 0, 0)
	return int(n), errnoErr(errno)
}

=======
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	sysSETSOCKOPT = 0xe
	sysGETSOCKOPT = 0xf
	sysSENDMSG    = 0x10
	sysRECVMSG    = 0x11
	sysRECVMMSG   = 0x13
	sysSENDMMSG   = 0x14
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	sysSETSOCKOPT = 0xe
	sysGETSOCKOPT = 0xf
	sysSENDMSG    = 0x10
	sysRECVMSG    = 0x11
	sysRECVMMSG   = 0x13
	sysSENDMMSG   = 0x14
=======
	sysRECVMMSG = 0x13
	sysSENDMMSG = 0x14
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
)

func socketcall(call, a0, a1, a2, a3, a4, a5 uintptr) (uintptr, syscall.Errno)
func rawsocketcall(call, a0, a1, a2, a3, a4, a5 uintptr) (uintptr, syscall.Errno)

<<<<<<< HEAD
func getsockopt(s uintptr, level, name int, b []byte) (int, error) {
	l := uint32(len(b))
	_, errno := socketcall(sysGETSOCKOPT, s, uintptr(level), uintptr(name), uintptr(unsafe.Pointer(&b[0])), uintptr(unsafe.Pointer(&l)), 0)
	return int(l), errnoErr(errno)
}

func setsockopt(s uintptr, level, name int, b []byte) error {
	_, errno := socketcall(sysSETSOCKOPT, s, uintptr(level), uintptr(name), uintptr(unsafe.Pointer(&b[0])), uintptr(len(b)), 0)
	return errnoErr(errno)
}

func recvmsg(s uintptr, h *msghdr, flags int) (int, error) {
	n, errno := socketcall(sysRECVMSG, s, uintptr(unsafe.Pointer(h)), uintptr(flags), 0, 0, 0)
	return int(n), errnoErr(errno)
}

func sendmsg(s uintptr, h *msghdr, flags int) (int, error) {
	n, errno := socketcall(sysSENDMSG, s, uintptr(unsafe.Pointer(h)), uintptr(flags), 0, 0, 0)
	return int(n), errnoErr(errno)
}

>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
func getsockopt(s uintptr, level, name int, b []byte) (int, error) {
	l := uint32(len(b))
	_, errno := socketcall(sysGETSOCKOPT, s, uintptr(level), uintptr(name), uintptr(unsafe.Pointer(&b[0])), uintptr(unsafe.Pointer(&l)), 0)
	return int(l), errnoErr(errno)
}

func setsockopt(s uintptr, level, name int, b []byte) error {
	_, errno := socketcall(sysSETSOCKOPT, s, uintptr(level), uintptr(name), uintptr(unsafe.Pointer(&b[0])), uintptr(len(b)), 0)
	return errnoErr(errno)
}

func recvmsg(s uintptr, h *msghdr, flags int) (int, error) {
	n, errno := socketcall(sysRECVMSG, s, uintptr(unsafe.Pointer(h)), uintptr(flags), 0, 0, 0)
	return int(n), errnoErr(errno)
}

func sendmsg(s uintptr, h *msghdr, flags int) (int, error) {
	n, errno := socketcall(sysSENDMSG, s, uintptr(unsafe.Pointer(h)), uintptr(flags), 0, 0, 0)
	return int(n), errnoErr(errno)
}

=======
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
func recvmmsg(s uintptr, hs []mmsghdr, flags int) (int, error) {
	n, errno := socketcall(sysRECVMMSG, s, uintptr(unsafe.Pointer(&hs[0])), uintptr(len(hs)), uintptr(flags), 0, 0)
	return int(n), errnoErr(errno)
}

func sendmmsg(s uintptr, hs []mmsghdr, flags int) (int, error) {
	n, errno := socketcall(sysSENDMMSG, s, uintptr(unsafe.Pointer(&hs[0])), uintptr(len(hs)), uintptr(flags), 0, 0)
	return int(n), errnoErr(errno)
}
