// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
//go:build aix || darwin || freebsd || linux || netbsd || openbsd || solaris || zos
// +build aix darwin freebsd linux netbsd openbsd solaris zos

package unix

import (
	"runtime"
)

// Round the length of a raw sockaddr up to align it properly.
func cmsgAlignOf(salen int) int {
	salign := SizeofPtr

	// dragonfly needs to check ABI version at runtime, see cmsgAlignOf in
	// sockcmsg_dragonfly.go
	switch runtime.GOOS {
	case "aix":
		// There is no alignment on AIX.
		salign = 1
	case "darwin", "ios", "illumos", "solaris":
		// NOTE: It seems like 64-bit Darwin, Illumos and Solaris
		// kernels still require 32-bit aligned access to network
		// subsystem.
		if SizeofPtr == 8 {
			salign = 4
		}
	case "netbsd", "openbsd":
		// NetBSD and OpenBSD armv7 require 64-bit alignment.
		if runtime.GOARCH == "arm" {
			salign = 8
		}
		// NetBSD aarch64 requires 128-bit alignment.
		if runtime.GOOS == "netbsd" && runtime.GOARCH == "arm64" {
			salign = 16
		}
	case "zos":
		// z/OS socket macros use [32-bit] sizeof(int) alignment,
		// not pointer width.
		salign = SizeofInt
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// +build aix darwin freebsd linux netbsd openbsd solaris
||||||| parent of 5ce8c7613 (update vendored files)
// +build aix darwin freebsd linux netbsd openbsd solaris
=======
//go:build aix || darwin || freebsd || linux || netbsd || openbsd || solaris || zos
// +build aix darwin freebsd linux netbsd openbsd solaris zos
>>>>>>> 5ce8c7613 (update vendored files)

package unix

import (
	"runtime"
)

// Round the length of a raw sockaddr up to align it properly.
func cmsgAlignOf(salen int) int {
	salign := SizeofPtr

	// dragonfly needs to check ABI version at runtime, see cmsgAlignOf in
	// sockcmsg_dragonfly.go
	switch runtime.GOOS {
	case "aix":
		// There is no alignment on AIX.
		salign = 1
	case "darwin", "ios", "illumos", "solaris":
		// NOTE: It seems like 64-bit Darwin, Illumos and Solaris
		// kernels still require 32-bit aligned access to network
		// subsystem.
		if SizeofPtr == 8 {
			salign = 4
		}
	case "netbsd", "openbsd":
		// NetBSD and OpenBSD armv7 require 64-bit alignment.
		if runtime.GOARCH == "arm" {
			salign = 8
		}
		// NetBSD aarch64 requires 128-bit alignment.
		if runtime.GOOS == "netbsd" && runtime.GOARCH == "arm64" {
			salign = 16
		}
<<<<<<< HEAD
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
=======
	case "zos":
		// z/OS socket macros use [32-bit] sizeof(int) alignment,
		// not pointer width.
		salign = SizeofInt
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// +build aix darwin freebsd linux netbsd openbsd solaris
||||||| parent of 6b7ce455e (update vendored files)
// +build aix darwin freebsd linux netbsd openbsd solaris
=======
//go:build aix || darwin || freebsd || linux || netbsd || openbsd || solaris || zos
// +build aix darwin freebsd linux netbsd openbsd solaris zos
>>>>>>> 6b7ce455e (update vendored files)

package unix

import (
	"runtime"
)

// Round the length of a raw sockaddr up to align it properly.
func cmsgAlignOf(salen int) int {
	salign := SizeofPtr

	// dragonfly needs to check ABI version at runtime, see cmsgAlignOf in
	// sockcmsg_dragonfly.go
	switch runtime.GOOS {
	case "aix":
		// There is no alignment on AIX.
		salign = 1
	case "darwin", "ios", "illumos", "solaris":
		// NOTE: It seems like 64-bit Darwin, Illumos and Solaris
		// kernels still require 32-bit aligned access to network
		// subsystem.
		if SizeofPtr == 8 {
			salign = 4
		}
	case "netbsd", "openbsd":
		// NetBSD and OpenBSD armv7 require 64-bit alignment.
		if runtime.GOARCH == "arm" {
			salign = 8
		}
		// NetBSD aarch64 requires 128-bit alignment.
		if runtime.GOOS == "netbsd" && runtime.GOARCH == "arm64" {
			salign = 16
		}
<<<<<<< HEAD
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
=======
	case "zos":
		// z/OS socket macros use [32-bit] sizeof(int) alignment,
		// not pointer width.
		salign = SizeofInt
>>>>>>> 6b7ce455e (update vendored files)
	}

	return (salen + salign - 1) & ^(salign - 1)
}
