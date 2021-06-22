// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
//go:build (linux && 386) || (linux && arm) || (linux && mips) || (linux && mipsle) || (linux && ppc)
// +build linux,386 linux,arm linux,mips linux,mipsle linux,ppc
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// +build linux,386 linux,arm linux,mips linux,mipsle
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
// +build linux,386 linux,arm linux,mips linux,mipsle
=======
//go:build (linux && 386) || (linux && arm) || (linux && mips) || (linux && mipsle) || (linux && ppc)
// +build linux,386 linux,arm linux,mips linux,mipsle linux,ppc
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// +build linux,386 linux,arm linux,mips linux,mipsle
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
// +build linux,386 linux,arm linux,mips linux,mipsle
=======
//go:build (linux && 386) || (linux && arm) || (linux && mips) || (linux && mipsle) || (linux && ppc)
// +build linux,386 linux,arm linux,mips linux,mipsle linux,ppc
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// +build linux,386 linux,arm linux,mips linux,mipsle
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
// +build linux,386 linux,arm linux,mips linux,mipsle
=======
//go:build (linux && 386) || (linux && arm) || (linux && mips) || (linux && mipsle) || (linux && ppc)
// +build linux,386 linux,arm linux,mips linux,mipsle linux,ppc
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// +build linux,386 linux,arm linux,mips linux,mipsle
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)

package unix

func init() {
	// On 32-bit Linux systems, the fcntl syscall that matches Go's
	// Flock_t type is SYS_FCNTL64, not SYS_FCNTL.
	fcntl64Syscall = SYS_FCNTL64
}
