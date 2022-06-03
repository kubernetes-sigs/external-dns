// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || windows || zos
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris windows zos

package socket

import (
	"net"
	"os"
)

func (c *Conn) recvMsg(m *Message, flags int) error {
	m.raceWrite()
	var (
		operr     error
		n         int
		oobn      int
		recvflags int
		from      net.Addr
	)
	fn := func(s uintptr) bool {
		n, oobn, recvflags, from, operr = recvmsg(s, m.Buffers, m.OOB, flags, c.network)
		return ioComplete(flags, operr)
	}
	if err := c.c.Read(fn); err != nil {
		return err
	}
	if operr != nil {
		return os.NewSyscallError("recvmsg", operr)
	}
	m.Addr = from
	m.N = n
	m.NN = oobn
	m.Flags = recvflags
	return nil
}

func (c *Conn) sendMsg(m *Message, flags int) error {
	m.raceRead()
	var (
		operr error
		n     int
	)
	fn := func(s uintptr) bool {
		n, operr = sendmsg(s, m.Buffers, m.OOB, m.Addr, flags)
		return ioComplete(flags, operr)
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 5ce8c7613 (update vendored files)
=======
//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || windows || zos
>>>>>>> 5ce8c7613 (update vendored files)
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris windows zos

package socket

import (
	"os"
)

func (c *Conn) recvMsg(m *Message, flags int) error {
	m.raceWrite()
	var h msghdr
	vs := make([]iovec, len(m.Buffers))
	var sa []byte
	if c.network != "tcp" {
		sa = make([]byte, sizeofSockaddrInet6)
	}
	h.pack(vs, m.Buffers, m.OOB, sa)
	var operr error
	var n int
	fn := func(s uintptr) bool {
		n, operr = recvmsg(s, &h, flags)
		return ioComplete(flags, operr)
	}
	if err := c.c.Read(fn); err != nil {
		return err
	}
	if operr != nil {
		return os.NewSyscallError("recvmsg", operr)
	}
	if c.network != "tcp" {
		var err error
		m.Addr, err = parseInetAddr(sa[:], c.network)
		if err != nil {
			return err
		}
	}
	m.N = n
	m.NN = h.controllen()
	m.Flags = h.flags()
	return nil
}

func (c *Conn) sendMsg(m *Message, flags int) error {
	m.raceRead()
	var h msghdr
	vs := make([]iovec, len(m.Buffers))
	var sa []byte
	if m.Addr != nil {
		var a [sizeofSockaddrInet6]byte
		n := marshalInetAddr(m.Addr, a[:])
		sa = a[:n]
	}
	h.pack(vs, m.Buffers, m.OOB, sa)
	var operr error
	var n int
	fn := func(s uintptr) bool {
		n, operr = sendmsg(s, &h, flags)
<<<<<<< HEAD
		if operr == syscall.EAGAIN || (runtime.GOOS == "zos" && operr == syscall.EWOULDBLOCK) {
			return false
		}
		return true
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
		if operr == syscall.EAGAIN || (runtime.GOOS == "zos" && operr == syscall.EWOULDBLOCK) {
			return false
		}
		return true
=======
		return ioComplete(flags, operr)
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 6b7ce455e (update vendored files)
=======
//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || windows || zos
>>>>>>> 6b7ce455e (update vendored files)
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris windows zos

package socket

import (
	"os"
)

func (c *Conn) recvMsg(m *Message, flags int) error {
	m.raceWrite()
	var h msghdr
	vs := make([]iovec, len(m.Buffers))
	var sa []byte
	if c.network != "tcp" {
		sa = make([]byte, sizeofSockaddrInet6)
	}
	h.pack(vs, m.Buffers, m.OOB, sa)
	var operr error
	var n int
	fn := func(s uintptr) bool {
		n, operr = recvmsg(s, &h, flags)
		return ioComplete(flags, operr)
	}
	if err := c.c.Read(fn); err != nil {
		return err
	}
	if operr != nil {
		return os.NewSyscallError("recvmsg", operr)
	}
	if c.network != "tcp" {
		var err error
		m.Addr, err = parseInetAddr(sa[:], c.network)
		if err != nil {
			return err
		}
	}
	m.N = n
	m.NN = h.controllen()
	m.Flags = h.flags()
	return nil
}

func (c *Conn) sendMsg(m *Message, flags int) error {
	m.raceRead()
	var h msghdr
	vs := make([]iovec, len(m.Buffers))
	var sa []byte
	if m.Addr != nil {
		var a [sizeofSockaddrInet6]byte
		n := marshalInetAddr(m.Addr, a[:])
		sa = a[:n]
	}
	h.pack(vs, m.Buffers, m.OOB, sa)
	var operr error
	var n int
	fn := func(s uintptr) bool {
		n, operr = sendmsg(s, &h, flags)
<<<<<<< HEAD
		if operr == syscall.EAGAIN || (runtime.GOOS == "zos" && operr == syscall.EWOULDBLOCK) {
			return false
		}
		return true
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
		if operr == syscall.EAGAIN || (runtime.GOOS == "zos" && operr == syscall.EWOULDBLOCK) {
			return false
		}
		return true
=======
		return ioComplete(flags, operr)
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 4d7e5ad26 (update vendored files)
=======
//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || windows || zos
>>>>>>> 4d7e5ad26 (update vendored files)
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris windows zos

package socket

import (
	"os"
)

func (c *Conn) recvMsg(m *Message, flags int) error {
	m.raceWrite()
	var h msghdr
	vs := make([]iovec, len(m.Buffers))
	var sa []byte
	if c.network != "tcp" {
		sa = make([]byte, sizeofSockaddrInet6)
	}
	h.pack(vs, m.Buffers, m.OOB, sa)
	var operr error
	var n int
	fn := func(s uintptr) bool {
		n, operr = recvmsg(s, &h, flags)
		return ioComplete(flags, operr)
	}
	if err := c.c.Read(fn); err != nil {
		return err
	}
	if operr != nil {
		return os.NewSyscallError("recvmsg", operr)
	}
	if c.network != "tcp" {
		var err error
		m.Addr, err = parseInetAddr(sa[:], c.network)
		if err != nil {
			return err
		}
	}
	m.N = n
	m.NN = h.controllen()
	m.Flags = h.flags()
	return nil
}

func (c *Conn) sendMsg(m *Message, flags int) error {
	m.raceRead()
	var h msghdr
	vs := make([]iovec, len(m.Buffers))
	var sa []byte
	if m.Addr != nil {
		var a [sizeofSockaddrInet6]byte
		n := marshalInetAddr(m.Addr, a[:])
		sa = a[:n]
	}
	h.pack(vs, m.Buffers, m.OOB, sa)
	var operr error
	var n int
	fn := func(s uintptr) bool {
		n, operr = sendmsg(s, &h, flags)
<<<<<<< HEAD
		if operr == syscall.EAGAIN || (runtime.GOOS == "zos" && operr == syscall.EWOULDBLOCK) {
			return false
		}
		return true
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
		if operr == syscall.EAGAIN || (runtime.GOOS == "zos" && operr == syscall.EWOULDBLOCK) {
			return false
		}
		return true
=======
		return ioComplete(flags, operr)
>>>>>>> 4d7e5ad26 (update vendored files)
	}
	if err := c.c.Write(fn); err != nil {
		return err
	}
	if operr != nil {
		return os.NewSyscallError("sendmsg", operr)
	}
	m.N = n
	m.NN = len(m.OOB)
	return nil
}
