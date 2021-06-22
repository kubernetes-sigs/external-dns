// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
//go:build race
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
=======
//go:build race
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
=======
//go:build race
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
// +build race

package socket

import (
	"runtime"
	"unsafe"
)

// This package reads and writes the Message buffers using a
// direct system call, which the race detector can't see.
// These functions tell the race detector what is going on during the syscall.

func (m *Message) raceRead() {
	for _, b := range m.Buffers {
		if len(b) > 0 {
			runtime.RaceReadRange(unsafe.Pointer(&b[0]), len(b))
		}
	}
	if b := m.OOB; len(b) > 0 {
		runtime.RaceReadRange(unsafe.Pointer(&b[0]), len(b))
	}
}
func (m *Message) raceWrite() {
	for _, b := range m.Buffers {
		if len(b) > 0 {
			runtime.RaceWriteRange(unsafe.Pointer(&b[0]), len(b))
		}
	}
	if b := m.OOB; len(b) > 0 {
		runtime.RaceWriteRange(unsafe.Pointer(&b[0]), len(b))
	}
}
