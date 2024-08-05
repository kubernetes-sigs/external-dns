// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build aix || windows || zos
<<<<<<< HEAD
// +build aix windows zos
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)

package socket

import (
	"syscall"
)

// ioComplete checks the flags and result of a syscall, to be used as return
// value in a syscall.RawConn.Read or Write callback.
func ioComplete(flags int, operr error) bool {
	if operr == syscall.EAGAIN || operr == syscall.EWOULDBLOCK {
		// No data available, block for I/O and try again.
		return false
	}
	return true
}
