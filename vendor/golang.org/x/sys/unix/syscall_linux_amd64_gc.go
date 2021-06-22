// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
//go:build amd64 && linux && gc
// +build amd64,linux,gc
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// +build amd64,linux
// +build gc
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
// +build amd64,linux
// +build gc
=======
//go:build amd64 && linux && gc
// +build amd64,linux,gc
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// +build amd64,linux
// +build gc
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)

package unix

import "syscall"

//go:noescape
func gettimeofday(tv *Timeval) (err syscall.Errno)
