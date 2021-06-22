// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
//go:build aix && ppc
// +build aix,ppc
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// +build aix
// +build ppc
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
// +build aix
// +build ppc
=======
//go:build aix && ppc
// +build aix,ppc
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// +build aix
// +build ppc
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
// +build aix
// +build ppc
=======
//go:build aix && ppc
// +build aix,ppc
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// +build aix
// +build ppc
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
// +build aix
// +build ppc
=======
//go:build aix && ppc
// +build aix,ppc
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// +build aix
// +build ppc
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)

// Functions to access/create device major and minor numbers matching the
// encoding used by AIX.

package unix

// Major returns the major component of a Linux device number.
func Major(dev uint64) uint32 {
	return uint32((dev >> 16) & 0xffff)
}

// Minor returns the minor component of a Linux device number.
func Minor(dev uint64) uint32 {
	return uint32(dev & 0xffff)
}

// Mkdev returns a Linux device number generated from the given major and minor
// components.
func Mkdev(major, minor uint32) uint64 {
	return uint64(((major) << 16) | (minor))
}
