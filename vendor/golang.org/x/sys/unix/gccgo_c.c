// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
<<<<<<< HEAD
//go:build gccgo && !aix && !hurd
// +build gccgo,!aix,!hurd
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// +build gccgo
// +build !aix
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
// +build gccgo
// +build !aix
=======
//go:build gccgo && !aix && !hurd
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)

#include <errno.h>
#include <stdint.h>
#include <unistd.h>

#define _STRINGIFY2_(x) #x
#define _STRINGIFY_(x) _STRINGIFY2_(x)
#define GOSYM_PREFIX _STRINGIFY_(__USER_LABEL_PREFIX__)

// Call syscall from C code because the gccgo support for calling from
// Go to C does not support varargs functions.

struct ret {
	uintptr_t r;
	uintptr_t err;
};

struct ret gccgoRealSyscall(uintptr_t trap, uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6, uintptr_t a7, uintptr_t a8, uintptr_t a9)
  __asm__(GOSYM_PREFIX GOPKGPATH ".realSyscall");

struct ret
gccgoRealSyscall(uintptr_t trap, uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6, uintptr_t a7, uintptr_t a8, uintptr_t a9)
{
	struct ret r;

	errno = 0;
	r.r = syscall(trap, a1, a2, a3, a4, a5, a6, a7, a8, a9);
	r.err = errno;
	return r;
}

uintptr_t gccgoRealSyscallNoError(uintptr_t trap, uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6, uintptr_t a7, uintptr_t a8, uintptr_t a9)
  __asm__(GOSYM_PREFIX GOPKGPATH ".realSyscallNoError");

uintptr_t
gccgoRealSyscallNoError(uintptr_t trap, uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6, uintptr_t a7, uintptr_t a8, uintptr_t a9)
{
	return syscall(trap, a1, a2, a3, a4, a5, a6, a7, a8, a9);
}
