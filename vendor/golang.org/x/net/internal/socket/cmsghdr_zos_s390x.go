// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package socket

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
func (h *cmsghdr) set(l, lvl, typ int) {
	h.Len = int32(l)
	h.Level = int32(lvl)
	h.Type = int32(typ)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
import "syscall"

||||||| parent of 4d7e5ad26 (update vendored files)
import "syscall"

=======
>>>>>>> 4d7e5ad26 (update vendored files)
func (h *cmsghdr) set(l, lvl, typ int) {
	h.Len = int32(l)
	h.Level = int32(lvl)
	h.Type = int32(typ)
}
<<<<<<< HEAD

func controlHeaderLen() int {
	return syscall.CmsgLen(0)
}

func controlMessageLen(dataLen int) int {
	return syscall.CmsgLen(dataLen)
}

func controlMessageSpace(dataLen int) int {
	return syscall.CmsgSpace(dataLen)
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
}
||||||| parent of 4d7e5ad26 (update vendored files)

func controlHeaderLen() int {
	return syscall.CmsgLen(0)
}

func controlMessageLen(dataLen int) int {
	return syscall.CmsgLen(dataLen)
}

func controlMessageSpace(dataLen int) int {
	return syscall.CmsgSpace(dataLen)
}
=======
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
import "syscall"

||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
import "syscall"

=======
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
func (h *cmsghdr) set(l, lvl, typ int) {
	h.Len = int32(l)
	h.Level = int32(lvl)
	h.Type = int32(typ)
}
<<<<<<< HEAD

func controlHeaderLen() int {
	return syscall.CmsgLen(0)
}

func controlMessageLen(dataLen int) int {
	return syscall.CmsgLen(dataLen)
}

func controlMessageSpace(dataLen int) int {
	return syscall.CmsgSpace(dataLen)
}
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)

func controlHeaderLen() int {
	return syscall.CmsgLen(0)
}

func controlMessageLen(dataLen int) int {
	return syscall.CmsgLen(dataLen)
}

func controlMessageSpace(dataLen int) int {
	return syscall.CmsgSpace(dataLen)
}
=======
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
