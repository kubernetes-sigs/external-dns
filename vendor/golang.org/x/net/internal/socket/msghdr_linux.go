// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package socket

import "unsafe"

func (h *msghdr) pack(vs []iovec, bs [][]byte, oob []byte, sa []byte) {
	for i := range vs {
		vs[i].set(bs[i])
	}
	h.setIov(vs)
	if len(oob) > 0 {
		h.setControl(oob)
	}
	if sa != nil {
		h.Name = (*byte)(unsafe.Pointer(&sa[0]))
		h.Namelen = uint32(len(sa))
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	} else {
		h.Name = nil
		h.Namelen = 0
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
=======
	} else {
		h.Name = nil
		h.Namelen = 0
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
=======
	} else {
		h.Name = nil
		h.Namelen = 0
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
=======
	} else {
		h.Name = nil
		h.Namelen = 0
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of e1cd8261c (UPSTREAM: <carry>: update vendored files v0.13.1)
	} else {
		h.Name = nil
		h.Namelen = 0
=======
>>>>>>> e1cd8261c (UPSTREAM: <carry>: update vendored files v0.13.1)
	}
}

func (h *msghdr) name() []byte {
	if h.Name != nil && h.Namelen > 0 {
		return (*[sizeofSockaddrInet6]byte)(unsafe.Pointer(h.Name))[:h.Namelen]
	}
	return nil
}

func (h *msghdr) controllen() int {
	return int(h.Controllen)
}

func (h *msghdr) flags() int {
	return int(h.Flags)
}
