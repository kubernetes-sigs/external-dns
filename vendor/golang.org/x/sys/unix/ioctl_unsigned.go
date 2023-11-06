// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD:vendor/golang.org/x/sys/unix/ioctl.go
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
=======
//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
=======
//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
=======
//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris
>>>>>>> 4d7e5ad26 (update vendored files)
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris
||||||| parent of 5d0416aaf (UPSTREAM: 3984: CVE-2023-44487 - bump golang.org/x/net v0.17.0):vendor/golang.org/x/sys/unix/ioctl.go
//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris
=======
//go:build darwin || dragonfly || freebsd || hurd || linux || netbsd || openbsd
// +build darwin dragonfly freebsd hurd linux netbsd openbsd
>>>>>>> 5d0416aaf (UPSTREAM: 3984: CVE-2023-44487 - bump golang.org/x/net v0.17.0):vendor/golang.org/x/sys/unix/ioctl_unsigned.go

package unix

import (
	"unsafe"
)

// ioctl itself should not be exposed directly, but additional get/set
// functions for specific types are permissible.

// IoctlSetInt performs an ioctl operation which sets an integer value
// on fd, using the specified request number.
func IoctlSetInt(fd int, req uint, value int) error {
	return ioctl(fd, req, uintptr(value))
}

// IoctlSetPointerInt performs an ioctl operation which sets an
// integer value on fd, using the specified request number. The ioctl
// argument is called with a pointer to the integer value, rather than
// passing the integer value directly.
func IoctlSetPointerInt(fd int, req uint, value int) error {
	v := int32(value)
	return ioctlPtr(fd, req, unsafe.Pointer(&v))
}

// IoctlSetWinsize performs an ioctl on fd with a *Winsize argument.
//
// To change fd's window size, the req argument should be TIOCSWINSZ.
func IoctlSetWinsize(fd int, req uint, value *Winsize) error {
	// TODO: if we get the chance, remove the req parameter and
	// hardcode TIOCSWINSZ.
	return ioctlPtr(fd, req, unsafe.Pointer(value))
}

// IoctlSetTermios performs an ioctl on fd with a *Termios.
//
// The req value will usually be TCSETA or TIOCSETA.
func IoctlSetTermios(fd int, req uint, value *Termios) error {
	// TODO: if we get the chance, remove the req parameter.
	return ioctlPtr(fd, req, unsafe.Pointer(value))
}

// IoctlGetInt performs an ioctl operation which gets an integer value
// from fd, using the specified request number.
//
// A few ioctl requests use the return value as an output parameter;
// for those, IoctlRetInt should be used instead of this function.
func IoctlGetInt(fd int, req uint) (int, error) {
	var value int
	err := ioctlPtr(fd, req, unsafe.Pointer(&value))
	return value, err
}

func IoctlGetWinsize(fd int, req uint) (*Winsize, error) {
	var value Winsize
	err := ioctlPtr(fd, req, unsafe.Pointer(&value))
	return &value, err
}

func IoctlGetTermios(fd int, req uint) (*Termios, error) {
	var value Termios
	err := ioctlPtr(fd, req, unsafe.Pointer(&value))
	return &value, err
}
