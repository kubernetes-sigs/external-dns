// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Windows environment variables.

package windows

import (
	"syscall"
	"unsafe"
)

func Getenv(key string) (value string, found bool) {
	return syscall.Getenv(key)
}

func Setenv(key, value string) error {
	return syscall.Setenv(key, value)
}

func Clearenv() {
	syscall.Clearenv()
}

func Environ() []string {
	return syscall.Environ()
}

// Returns a default environment associated with the token, rather than the current
// process. If inheritExisting is true, then this environment also inherits the
// environment of the current process.
func (token Token) Environ(inheritExisting bool) (env []string, err error) {
	var block *uint16
	err = CreateEnvironmentBlock(&block, token, inheritExisting)
	if err != nil {
		return nil, err
	}
	defer DestroyEnvironmentBlock(block)
<<<<<<< HEAD
<<<<<<< HEAD
	blockp := unsafe.Pointer(block)
	for {
		entry := UTF16PtrToString((*uint16)(blockp))
		if len(entry) == 0 {
			break
		}
		env = append(env, entry)
		blockp = unsafe.Add(blockp, 2*(len(entry)+1))
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	blockp := uintptr(unsafe.Pointer(block))
	for {
		entry := UTF16PtrToString((*uint16)(unsafe.Pointer(blockp)))
		if len(entry) == 0 {
			break
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	blockp := uintptr(unsafe.Pointer(block))
	for {
		entry := UTF16PtrToString((*uint16)(unsafe.Pointer(blockp)))
		if len(entry) == 0 {
			break
=======
	size := unsafe.Sizeof(*block)
	for *block != 0 {
		// find NUL terminator
		end := unsafe.Pointer(block)
		for *(*uint16)(end) != 0 {
			end = unsafe.Add(end, size)
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
		}
<<<<<<< HEAD
		env = append(env, entry)
		blockp += 2 * (uintptr(len(entry)) + 1)
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
		env = append(env, entry)
		blockp += 2 * (uintptr(len(entry)) + 1)
=======

		entry := unsafe.Slice(block, (uintptr(end)-uintptr(unsafe.Pointer(block)))/size)
		env = append(env, UTF16ToString(entry))
		block = (*uint16)(unsafe.Add(end, size))
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	}
	return env, nil
}

func Unsetenv(key string) error {
	return syscall.Unsetenv(key)
}
