// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
//go:build darwin || freebsd || linux
// +build darwin freebsd linux

package ipv4

import (
	"net"
	"unsafe"

	"golang.org/x/net/internal/socket"

	"golang.org/x/sys/unix"
)

func (so *sockOpt) getIPMreqn(c *socket.Conn) (*net.Interface, error) {
	b := make([]byte, so.Len)
	if _, err := so.Get(c, b); err != nil {
		return nil, err
	}
	mreqn := (*unix.IPMreqn)(unsafe.Pointer(&b[0]))
	if mreqn.Ifindex == 0 {
		return nil, nil
	}
	ifi, err := net.InterfaceByIndex(int(mreqn.Ifindex))
	if err != nil {
		return nil, err
	}
	return ifi, nil
}

func (so *sockOpt) setIPMreqn(c *socket.Conn, ifi *net.Interface, grp net.IP) error {
	var mreqn unix.IPMreqn
	if ifi != nil {
		mreqn.Ifindex = int32(ifi.Index)
	}
	if grp != nil {
		mreqn.Multiaddr = [4]byte{grp[0], grp[1], grp[2], grp[3]}
	}
	b := (*[unix.SizeofIPMreqn]byte)(unsafe.Pointer(&mreqn))[:unix.SizeofIPMreqn]
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 5ce8c7613 (update vendored files)
=======
//go:build darwin || freebsd || linux
>>>>>>> 5ce8c7613 (update vendored files)
// +build darwin freebsd linux

package ipv4

import (
	"net"
	"unsafe"

	"golang.org/x/net/internal/socket"

	"golang.org/x/sys/unix"
)

func (so *sockOpt) getIPMreqn(c *socket.Conn) (*net.Interface, error) {
	b := make([]byte, so.Len)
	if _, err := so.Get(c, b); err != nil {
		return nil, err
	}
	mreqn := (*unix.IPMreqn)(unsafe.Pointer(&b[0]))
	if mreqn.Ifindex == 0 {
		return nil, nil
	}
	ifi, err := net.InterfaceByIndex(int(mreqn.Ifindex))
	if err != nil {
		return nil, err
	}
	return ifi, nil
}

func (so *sockOpt) setIPMreqn(c *socket.Conn, ifi *net.Interface, grp net.IP) error {
	var mreqn unix.IPMreqn
	if ifi != nil {
		mreqn.Ifindex = int32(ifi.Index)
	}
	if grp != nil {
		mreqn.Multiaddr = [4]byte{grp[0], grp[1], grp[2], grp[3]}
	}
<<<<<<< HEAD
	b := (*[sizeofIPMreqn]byte)(unsafe.Pointer(&mreqn))[:sizeofIPMreqn]
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
	b := (*[sizeofIPMreqn]byte)(unsafe.Pointer(&mreqn))[:sizeofIPMreqn]
=======
	b := (*[unix.SizeofIPMreqn]byte)(unsafe.Pointer(&mreqn))[:unix.SizeofIPMreqn]
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 6b7ce455e (update vendored files)
=======
//go:build darwin || freebsd || linux
>>>>>>> 6b7ce455e (update vendored files)
// +build darwin freebsd linux

package ipv4

import (
	"net"
	"unsafe"

	"golang.org/x/net/internal/socket"

	"golang.org/x/sys/unix"
)

func (so *sockOpt) getIPMreqn(c *socket.Conn) (*net.Interface, error) {
	b := make([]byte, so.Len)
	if _, err := so.Get(c, b); err != nil {
		return nil, err
	}
	mreqn := (*unix.IPMreqn)(unsafe.Pointer(&b[0]))
	if mreqn.Ifindex == 0 {
		return nil, nil
	}
	ifi, err := net.InterfaceByIndex(int(mreqn.Ifindex))
	if err != nil {
		return nil, err
	}
	return ifi, nil
}

func (so *sockOpt) setIPMreqn(c *socket.Conn, ifi *net.Interface, grp net.IP) error {
	var mreqn unix.IPMreqn
	if ifi != nil {
		mreqn.Ifindex = int32(ifi.Index)
	}
	if grp != nil {
		mreqn.Multiaddr = [4]byte{grp[0], grp[1], grp[2], grp[3]}
	}
<<<<<<< HEAD
	b := (*[sizeofIPMreqn]byte)(unsafe.Pointer(&mreqn))[:sizeofIPMreqn]
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
	b := (*[sizeofIPMreqn]byte)(unsafe.Pointer(&mreqn))[:sizeofIPMreqn]
=======
	b := (*[unix.SizeofIPMreqn]byte)(unsafe.Pointer(&mreqn))[:unix.SizeofIPMreqn]
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 4d7e5ad26 (update vendored files)
=======
//go:build darwin || freebsd || linux
>>>>>>> 4d7e5ad26 (update vendored files)
// +build darwin freebsd linux

package ipv4

import (
	"net"
	"unsafe"

	"golang.org/x/net/internal/socket"

	"golang.org/x/sys/unix"
)

func (so *sockOpt) getIPMreqn(c *socket.Conn) (*net.Interface, error) {
	b := make([]byte, so.Len)
	if _, err := so.Get(c, b); err != nil {
		return nil, err
	}
	mreqn := (*unix.IPMreqn)(unsafe.Pointer(&b[0]))
	if mreqn.Ifindex == 0 {
		return nil, nil
	}
	ifi, err := net.InterfaceByIndex(int(mreqn.Ifindex))
	if err != nil {
		return nil, err
	}
	return ifi, nil
}

func (so *sockOpt) setIPMreqn(c *socket.Conn, ifi *net.Interface, grp net.IP) error {
	var mreqn unix.IPMreqn
	if ifi != nil {
		mreqn.Ifindex = int32(ifi.Index)
	}
	if grp != nil {
		mreqn.Multiaddr = [4]byte{grp[0], grp[1], grp[2], grp[3]}
	}
<<<<<<< HEAD
	b := (*[sizeofIPMreqn]byte)(unsafe.Pointer(&mreqn))[:sizeofIPMreqn]
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
	b := (*[sizeofIPMreqn]byte)(unsafe.Pointer(&mreqn))[:sizeofIPMreqn]
=======
	b := (*[unix.SizeofIPMreqn]byte)(unsafe.Pointer(&mreqn))[:unix.SizeofIPMreqn]
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// +build darwin freebsd linux
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
// +build darwin freebsd linux
=======
//go:build darwin || freebsd || linux
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)

package ipv4

import (
	"net"
	"unsafe"

	"golang.org/x/net/internal/socket"

	"golang.org/x/sys/unix"
)

func (so *sockOpt) getIPMreqn(c *socket.Conn) (*net.Interface, error) {
	b := make([]byte, so.Len)
	if _, err := so.Get(c, b); err != nil {
		return nil, err
	}
	mreqn := (*unix.IPMreqn)(unsafe.Pointer(&b[0]))
	if mreqn.Ifindex == 0 {
		return nil, nil
	}
	ifi, err := net.InterfaceByIndex(int(mreqn.Ifindex))
	if err != nil {
		return nil, err
	}
	return ifi, nil
}

func (so *sockOpt) setIPMreqn(c *socket.Conn, ifi *net.Interface, grp net.IP) error {
	var mreqn unix.IPMreqn
	if ifi != nil {
		mreqn.Ifindex = int32(ifi.Index)
	}
	if grp != nil {
		mreqn.Multiaddr = [4]byte{grp[0], grp[1], grp[2], grp[3]}
	}
<<<<<<< HEAD
	b := (*[sizeofIPMreqn]byte)(unsafe.Pointer(&mreqn))[:sizeofIPMreqn]
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	b := (*[sizeofIPMreqn]byte)(unsafe.Pointer(&mreqn))[:sizeofIPMreqn]
=======
	b := (*[unix.SizeofIPMreqn]byte)(unsafe.Pointer(&mreqn))[:unix.SizeofIPMreqn]
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	return so.Set(c, b)
}
