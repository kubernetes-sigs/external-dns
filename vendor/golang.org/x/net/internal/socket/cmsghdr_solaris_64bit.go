// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
//go:build amd64 && solaris
// +build amd64,solaris
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// +build amd64
// +build solaris
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
// +build amd64
// +build solaris
=======
//go:build amd64 && solaris
// +build amd64,solaris
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// +build amd64
// +build solaris
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)

package socket

func (h *cmsghdr) set(l, lvl, typ int) {
	h.Len = uint32(l)
	h.Level = int32(lvl)
	h.Type = int32(typ)
}
