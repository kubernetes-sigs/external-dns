// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
<<<<<<< HEAD
//go:build (aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || zos) && go1.9
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris zos
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris
=======
//go:build (aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || zos) && go1.9
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris zos
>>>>>>> 5ce8c7613 (update vendored files)
// +build go1.9

package unix

import "syscall"

type Signal = syscall.Signal
type Errno = syscall.Errno
type SysProcAttr = syscall.SysProcAttr
