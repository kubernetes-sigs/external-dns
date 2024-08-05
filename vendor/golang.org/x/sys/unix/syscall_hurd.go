// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build hurd
<<<<<<< HEAD
// +build hurd
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)

package unix

/*
#include <stdint.h>
int ioctl(int, unsigned long int, uintptr_t);
*/
import "C"

func ioctl(fd int, req uint, arg uintptr) (err error) {
	r0, er := C.ioctl(C.int(fd), C.ulong(req), C.uintptr_t(arg))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

func ioctlPtr(fd int, req uint, arg unsafe.Pointer) (err error) {
	r0, er := C.ioctl(C.int(fd), C.ulong(req), C.uintptr_t(uintptr(arg)))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}
