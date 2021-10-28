// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
<<<<<<< HEAD
//go:build arm && gc && linux
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
=======
//go:build arm && gc && linux
>>>>>>> 5ce8c7613 (update vendored files)
// +build arm,gc,linux

package unix

import "syscall"

// Underlying system call writes to newoffset via pointer.
// Implemented in assembly to avoid allocation.
func seek(fd int, offset int64, whence int) (newoffset int64, err syscall.Errno)
