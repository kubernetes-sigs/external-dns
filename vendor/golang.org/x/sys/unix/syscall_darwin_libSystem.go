// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
//go:build darwin && go1.12
// +build darwin,go1.12

package unix

import _ "unsafe"

// Implemented in the runtime package (runtime/sys_darwin.go)
func syscall_syscall(fn, a1, a2, a3 uintptr) (r1, r2 uintptr, err Errno)
func syscall_syscall6(fn, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno)
func syscall_syscall6X(fn, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno)
func syscall_syscall9(fn, a1, a2, a3, a4, a5, a6, a7, a8, a9 uintptr) (r1, r2 uintptr, err Errno) // 32-bit only
func syscall_rawSyscall(fn, a1, a2, a3 uintptr) (r1, r2 uintptr, err Errno)
func syscall_rawSyscall6(fn, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno)
func syscall_syscallPtr(fn, a1, a2, a3 uintptr) (r1, r2 uintptr, err Errno)

//go:linkname syscall_syscall syscall.syscall
//go:linkname syscall_syscall6 syscall.syscall6
//go:linkname syscall_syscall6X syscall.syscall6X
//go:linkname syscall_syscall9 syscall.syscall9
//go:linkname syscall_rawSyscall syscall.rawSyscall
//go:linkname syscall_rawSyscall6 syscall.rawSyscall6
//go:linkname syscall_syscallPtr syscall.syscallPtr
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 5ce8c7613 (update vendored files)
=======
//go:build darwin && go1.12
>>>>>>> 5ce8c7613 (update vendored files)
// +build darwin,go1.12

package unix

import _ "unsafe"

// Implemented in the runtime package (runtime/sys_darwin.go)
func syscall_syscall(fn, a1, a2, a3 uintptr) (r1, r2 uintptr, err Errno)
func syscall_syscall6(fn, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno)
func syscall_syscall6X(fn, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno)
func syscall_syscall9(fn, a1, a2, a3, a4, a5, a6, a7, a8, a9 uintptr) (r1, r2 uintptr, err Errno) // 32-bit only
func syscall_rawSyscall(fn, a1, a2, a3 uintptr) (r1, r2 uintptr, err Errno)
func syscall_rawSyscall6(fn, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno)
func syscall_syscallPtr(fn, a1, a2, a3 uintptr) (r1, r2 uintptr, err Errno)

//go:linkname syscall_syscall syscall.syscall
//go:linkname syscall_syscall6 syscall.syscall6
//go:linkname syscall_syscall6X syscall.syscall6X
//go:linkname syscall_syscall9 syscall.syscall9
//go:linkname syscall_rawSyscall syscall.rawSyscall
//go:linkname syscall_rawSyscall6 syscall.rawSyscall6
//go:linkname syscall_syscallPtr syscall.syscallPtr
<<<<<<< HEAD

// Find the entry point for f. See comments in runtime/proc.go for the
// function of the same name.
//go:nosplit
func funcPC(f func()) uintptr {
	return **(**uintptr)(unsafe.Pointer(&f))
}
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)

// Find the entry point for f. See comments in runtime/proc.go for the
// function of the same name.
//go:nosplit
func funcPC(f func()) uintptr {
	return **(**uintptr)(unsafe.Pointer(&f))
}
=======
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 6b7ce455e (update vendored files)
=======
//go:build darwin && go1.12
>>>>>>> 6b7ce455e (update vendored files)
// +build darwin,go1.12

package unix

import _ "unsafe"

// Implemented in the runtime package (runtime/sys_darwin.go)
func syscall_syscall(fn, a1, a2, a3 uintptr) (r1, r2 uintptr, err Errno)
func syscall_syscall6(fn, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno)
func syscall_syscall6X(fn, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno)
func syscall_syscall9(fn, a1, a2, a3, a4, a5, a6, a7, a8, a9 uintptr) (r1, r2 uintptr, err Errno) // 32-bit only
func syscall_rawSyscall(fn, a1, a2, a3 uintptr) (r1, r2 uintptr, err Errno)
func syscall_rawSyscall6(fn, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno)
func syscall_syscallPtr(fn, a1, a2, a3 uintptr) (r1, r2 uintptr, err Errno)

//go:linkname syscall_syscall syscall.syscall
//go:linkname syscall_syscall6 syscall.syscall6
//go:linkname syscall_syscall6X syscall.syscall6X
//go:linkname syscall_syscall9 syscall.syscall9
//go:linkname syscall_rawSyscall syscall.rawSyscall
//go:linkname syscall_rawSyscall6 syscall.rawSyscall6
//go:linkname syscall_syscallPtr syscall.syscallPtr
<<<<<<< HEAD

// Find the entry point for f. See comments in runtime/proc.go for the
// function of the same name.
//go:nosplit
func funcPC(f func()) uintptr {
	return **(**uintptr)(unsafe.Pointer(&f))
}
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)

// Find the entry point for f. See comments in runtime/proc.go for the
// function of the same name.
//go:nosplit
func funcPC(f func()) uintptr {
	return **(**uintptr)(unsafe.Pointer(&f))
}
=======
>>>>>>> 6b7ce455e (update vendored files)
