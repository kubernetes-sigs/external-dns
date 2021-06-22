// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
//go:build linux
// +build linux

package socket

import (
	"net"
	"os"
)

func (c *Conn) recvMsgs(ms []Message, flags int) (int, error) {
	for i := range ms {
		ms[i].raceWrite()
	}
	packer := defaultMmsghdrsPool.Get()
	defer defaultMmsghdrsPool.Put(packer)
	var parseFn func([]byte, string) (net.Addr, error)
	if c.network != "tcp" {
		parseFn = parseInetAddr
	}
	hs := packer.pack(ms, parseFn, nil)
	var operr error
	var n int
	fn := func(s uintptr) bool {
		n, operr = recvmmsg(s, hs, flags)
		return ioComplete(flags, operr)
	}
	if err := c.c.Read(fn); err != nil {
		return n, err
	}
	if operr != nil {
		return n, os.NewSyscallError("recvmmsg", operr)
	}
	if err := hs[:n].unpack(ms[:n], parseFn, c.network); err != nil {
		return n, err
	}
	return n, nil
}

func (c *Conn) sendMsgs(ms []Message, flags int) (int, error) {
	for i := range ms {
		ms[i].raceRead()
	}
	packer := defaultMmsghdrsPool.Get()
	defer defaultMmsghdrsPool.Put(packer)
	var marshalFn func(net.Addr, []byte) int
	if c.network != "tcp" {
		marshalFn = marshalInetAddr
	}
	hs := packer.pack(ms, nil, marshalFn)
	var operr error
	var n int
	fn := func(s uintptr) bool {
		n, operr = sendmmsg(s, hs, flags)
		return ioComplete(flags, operr)
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 5ce8c7613 (update vendored files)
=======
//go:build linux
>>>>>>> 5ce8c7613 (update vendored files)
// +build linux

package socket

import (
	"net"
	"os"
)

func (c *Conn) recvMsgs(ms []Message, flags int) (int, error) {
	for i := range ms {
		ms[i].raceWrite()
	}
	packer := defaultMmsghdrsPool.Get()
	defer defaultMmsghdrsPool.Put(packer)
	var parseFn func([]byte, string) (net.Addr, error)
	if c.network != "tcp" {
		parseFn = parseInetAddr
	}
	hs := packer.pack(ms, parseFn, nil)
	var operr error
	var n int
	fn := func(s uintptr) bool {
		n, operr = recvmmsg(s, hs, flags)
		return ioComplete(flags, operr)
	}
	if err := c.c.Read(fn); err != nil {
		return n, err
	}
	if operr != nil {
		return n, os.NewSyscallError("recvmmsg", operr)
	}
	if err := hs[:n].unpack(ms[:n], parseFn, c.network); err != nil {
		return n, err
	}
	return n, nil
}

func (c *Conn) sendMsgs(ms []Message, flags int) (int, error) {
	for i := range ms {
		ms[i].raceRead()
	}
	packer := defaultMmsghdrsPool.Get()
	defer defaultMmsghdrsPool.Put(packer)
	var marshalFn func(net.Addr, []byte) int
	if c.network != "tcp" {
		marshalFn = marshalInetAddr
	}
	hs := packer.pack(ms, nil, marshalFn)
	var operr error
	var n int
	fn := func(s uintptr) bool {
		n, operr = sendmmsg(s, hs, flags)
<<<<<<< HEAD
		if operr == syscall.EAGAIN {
			return false
		}
		return true
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
		if operr == syscall.EAGAIN {
			return false
		}
		return true
=======
		return ioComplete(flags, operr)
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// +build linux

package socket

import (
	"net"
	"os"
	"syscall"
)

func (c *Conn) recvMsgs(ms []Message, flags int) (int, error) {
	for i := range ms {
		ms[i].raceWrite()
	}
	hs := make(mmsghdrs, len(ms))
	var parseFn func([]byte, string) (net.Addr, error)
	if c.network != "tcp" {
		parseFn = parseInetAddr
	}
	if err := hs.pack(ms, parseFn, nil); err != nil {
		return 0, err
	}
	var operr error
	var n int
	fn := func(s uintptr) bool {
		n, operr = recvmmsg(s, hs, flags)
		if operr == syscall.EAGAIN {
			return false
		}
		return true
	}
	if err := c.c.Read(fn); err != nil {
		return n, err
	}
	if operr != nil {
		return n, os.NewSyscallError("recvmmsg", operr)
	}
	if err := hs[:n].unpack(ms[:n], parseFn, c.network); err != nil {
		return n, err
	}
	return n, nil
}

func (c *Conn) sendMsgs(ms []Message, flags int) (int, error) {
	for i := range ms {
		ms[i].raceRead()
	}
	hs := make(mmsghdrs, len(ms))
	var marshalFn func(net.Addr) []byte
	if c.network != "tcp" {
		marshalFn = marshalInetAddr
	}
	if err := hs.pack(ms, nil, marshalFn); err != nil {
		return 0, err
	}
	var operr error
	var n int
	fn := func(s uintptr) bool {
		n, operr = sendmmsg(s, hs, flags)
		if operr == syscall.EAGAIN {
			return false
		}
		return true
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
	}
	if err := c.c.Write(fn); err != nil {
		return n, err
	}
	if operr != nil {
		return n, os.NewSyscallError("sendmmsg", operr)
	}
	if err := hs[:n].unpack(ms[:n], nil, ""); err != nil {
		return n, err
	}
	return n, nil
}
