// Copyright 2012 The Go Authors. All rights reserved.
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
//go:build aix || darwin || dragonfly || freebsd || netbsd || openbsd
// +build aix darwin dragonfly freebsd netbsd openbsd

package ipv4

import (
	"net"
	"syscall"
	"unsafe"

	"golang.org/x/net/internal/iana"
	"golang.org/x/net/internal/socket"

	"golang.org/x/sys/unix"
)

func marshalDst(b []byte, cm *ControlMessage) []byte {
	m := socket.ControlMessage(b)
	m.MarshalHeader(iana.ProtocolIP, unix.IP_RECVDSTADDR, net.IPv4len)
	return m.Next(net.IPv4len)
}

func parseDst(cm *ControlMessage, b []byte) {
	if len(cm.Dst) < net.IPv4len {
		cm.Dst = make(net.IP, net.IPv4len)
	}
	copy(cm.Dst, b[:net.IPv4len])
}

func marshalInterface(b []byte, cm *ControlMessage) []byte {
	m := socket.ControlMessage(b)
	m.MarshalHeader(iana.ProtocolIP, sockoptReceiveInterface, syscall.SizeofSockaddrDatalink)
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 5ce8c7613 (update vendored files)
=======
//go:build aix || darwin || dragonfly || freebsd || netbsd || openbsd
>>>>>>> 5ce8c7613 (update vendored files)
// +build aix darwin dragonfly freebsd netbsd openbsd

package ipv4

import (
	"net"
	"syscall"
	"unsafe"

	"golang.org/x/net/internal/iana"
	"golang.org/x/net/internal/socket"

	"golang.org/x/sys/unix"
)

func marshalDst(b []byte, cm *ControlMessage) []byte {
	m := socket.ControlMessage(b)
	m.MarshalHeader(iana.ProtocolIP, unix.IP_RECVDSTADDR, net.IPv4len)
	return m.Next(net.IPv4len)
}

func parseDst(cm *ControlMessage, b []byte) {
	if len(cm.Dst) < net.IPv4len {
		cm.Dst = make(net.IP, net.IPv4len)
	}
	copy(cm.Dst, b[:net.IPv4len])
}

func marshalInterface(b []byte, cm *ControlMessage) []byte {
	m := socket.ControlMessage(b)
<<<<<<< HEAD
	m.MarshalHeader(iana.ProtocolIP, sysIP_RECVIF, syscall.SizeofSockaddrDatalink)
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
	m.MarshalHeader(iana.ProtocolIP, sysIP_RECVIF, syscall.SizeofSockaddrDatalink)
=======
	m.MarshalHeader(iana.ProtocolIP, sockoptReceiveInterface, syscall.SizeofSockaddrDatalink)
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 6b7ce455e (update vendored files)
=======
//go:build aix || darwin || dragonfly || freebsd || netbsd || openbsd
>>>>>>> 6b7ce455e (update vendored files)
// +build aix darwin dragonfly freebsd netbsd openbsd

package ipv4

import (
	"net"
	"syscall"
	"unsafe"

	"golang.org/x/net/internal/iana"
	"golang.org/x/net/internal/socket"

	"golang.org/x/sys/unix"
)

func marshalDst(b []byte, cm *ControlMessage) []byte {
	m := socket.ControlMessage(b)
	m.MarshalHeader(iana.ProtocolIP, unix.IP_RECVDSTADDR, net.IPv4len)
	return m.Next(net.IPv4len)
}

func parseDst(cm *ControlMessage, b []byte) {
	if len(cm.Dst) < net.IPv4len {
		cm.Dst = make(net.IP, net.IPv4len)
	}
	copy(cm.Dst, b[:net.IPv4len])
}

func marshalInterface(b []byte, cm *ControlMessage) []byte {
	m := socket.ControlMessage(b)
<<<<<<< HEAD
	m.MarshalHeader(iana.ProtocolIP, sysIP_RECVIF, syscall.SizeofSockaddrDatalink)
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
	m.MarshalHeader(iana.ProtocolIP, sysIP_RECVIF, syscall.SizeofSockaddrDatalink)
=======
	m.MarshalHeader(iana.ProtocolIP, sockoptReceiveInterface, syscall.SizeofSockaddrDatalink)
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 4d7e5ad26 (update vendored files)
=======
//go:build aix || darwin || dragonfly || freebsd || netbsd || openbsd
>>>>>>> 4d7e5ad26 (update vendored files)
// +build aix darwin dragonfly freebsd netbsd openbsd

package ipv4

import (
	"net"
	"syscall"
	"unsafe"

	"golang.org/x/net/internal/iana"
	"golang.org/x/net/internal/socket"

	"golang.org/x/sys/unix"
)

func marshalDst(b []byte, cm *ControlMessage) []byte {
	m := socket.ControlMessage(b)
	m.MarshalHeader(iana.ProtocolIP, unix.IP_RECVDSTADDR, net.IPv4len)
	return m.Next(net.IPv4len)
}

func parseDst(cm *ControlMessage, b []byte) {
	if len(cm.Dst) < net.IPv4len {
		cm.Dst = make(net.IP, net.IPv4len)
	}
	copy(cm.Dst, b[:net.IPv4len])
}

func marshalInterface(b []byte, cm *ControlMessage) []byte {
	m := socket.ControlMessage(b)
<<<<<<< HEAD
	m.MarshalHeader(iana.ProtocolIP, sysIP_RECVIF, syscall.SizeofSockaddrDatalink)
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
	m.MarshalHeader(iana.ProtocolIP, sysIP_RECVIF, syscall.SizeofSockaddrDatalink)
=======
	m.MarshalHeader(iana.ProtocolIP, sockoptReceiveInterface, syscall.SizeofSockaddrDatalink)
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// +build aix darwin dragonfly freebsd netbsd openbsd
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
// +build aix darwin dragonfly freebsd netbsd openbsd
=======
//go:build aix || darwin || dragonfly || freebsd || netbsd || openbsd
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)

package ipv4

import (
	"net"
	"syscall"
	"unsafe"

	"golang.org/x/net/internal/iana"
	"golang.org/x/net/internal/socket"

	"golang.org/x/sys/unix"
)

func marshalDst(b []byte, cm *ControlMessage) []byte {
	m := socket.ControlMessage(b)
	m.MarshalHeader(iana.ProtocolIP, unix.IP_RECVDSTADDR, net.IPv4len)
	return m.Next(net.IPv4len)
}

func parseDst(cm *ControlMessage, b []byte) {
	if len(cm.Dst) < net.IPv4len {
		cm.Dst = make(net.IP, net.IPv4len)
	}
	copy(cm.Dst, b[:net.IPv4len])
}

func marshalInterface(b []byte, cm *ControlMessage) []byte {
	m := socket.ControlMessage(b)
<<<<<<< HEAD
	m.MarshalHeader(iana.ProtocolIP, sysIP_RECVIF, syscall.SizeofSockaddrDatalink)
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	m.MarshalHeader(iana.ProtocolIP, sysIP_RECVIF, syscall.SizeofSockaddrDatalink)
=======
	m.MarshalHeader(iana.ProtocolIP, sockoptReceiveInterface, syscall.SizeofSockaddrDatalink)
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	return m.Next(syscall.SizeofSockaddrDatalink)
}

func parseInterface(cm *ControlMessage, b []byte) {
	var sadl syscall.SockaddrDatalink
	copy((*[unsafe.Sizeof(sadl)]byte)(unsafe.Pointer(&sadl))[:], b)
	cm.IfIndex = int(sadl.Index)
}
