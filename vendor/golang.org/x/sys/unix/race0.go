// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
<<<<<<< HEAD
//go:build aix || (darwin && !race) || (linux && !race) || (freebsd && !race) || netbsd || openbsd || solaris || dragonfly || zos
// +build aix darwin,!race linux,!race freebsd,!race netbsd openbsd solaris dragonfly zos
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// +build aix darwin,!race linux,!race freebsd,!race netbsd openbsd solaris dragonfly
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
// +build aix darwin,!race linux,!race freebsd,!race netbsd openbsd solaris dragonfly
=======
//go:build aix || (darwin && !race) || (linux && !race) || (freebsd && !race) || netbsd || openbsd || solaris || dragonfly || zos
// +build aix darwin,!race linux,!race freebsd,!race netbsd openbsd solaris dragonfly zos
>>>>>>> 5ce8c7613 (update vendored files)

package unix

import (
	"unsafe"
)

const raceenabled = false

func raceAcquire(addr unsafe.Pointer) {
}

func raceReleaseMerge(addr unsafe.Pointer) {
}

func raceReadRange(addr unsafe.Pointer, len int) {
}

func raceWriteRange(addr unsafe.Pointer, len int) {
}
