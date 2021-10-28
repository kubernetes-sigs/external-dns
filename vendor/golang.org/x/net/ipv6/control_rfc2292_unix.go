// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
//go:build darwin
// +build darwin

package ipv6

import (
	"unsafe"

	"golang.org/x/net/internal/iana"
	"golang.org/x/net/internal/socket"

	"golang.org/x/sys/unix"
)

func marshal2292HopLimit(b []byte, cm *ControlMessage) []byte {
	m := socket.ControlMessage(b)
	m.MarshalHeader(iana.ProtocolIPv6, unix.IPV6_2292HOPLIMIT, 4)
	if cm != nil {
		socket.NativeEndian.PutUint32(m.Data(4), uint32(cm.HopLimit))
	}
	return m.Next(4)
}

func marshal2292PacketInfo(b []byte, cm *ControlMessage) []byte {
	m := socket.ControlMessage(b)
	m.MarshalHeader(iana.ProtocolIPv6, unix.IPV6_2292PKTINFO, sizeofInet6Pktinfo)
	if cm != nil {
		pi := (*inet6Pktinfo)(unsafe.Pointer(&m.Data(sizeofInet6Pktinfo)[0]))
		if ip := cm.Src.To16(); ip != nil && ip.To4() == nil {
			copy(pi.Addr[:], ip)
		}
		if cm.IfIndex > 0 {
			pi.setIfindex(cm.IfIndex)
		}
	}
	return m.Next(sizeofInet6Pktinfo)
}

func marshal2292NextHop(b []byte, cm *ControlMessage) []byte {
	m := socket.ControlMessage(b)
	m.MarshalHeader(iana.ProtocolIPv6, unix.IPV6_2292NEXTHOP, sizeofSockaddrInet6)
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 5ce8c7613 (update vendored files)
=======
//go:build darwin
>>>>>>> 5ce8c7613 (update vendored files)
// +build darwin

package ipv6

import (
	"unsafe"

	"golang.org/x/net/internal/iana"
	"golang.org/x/net/internal/socket"

	"golang.org/x/sys/unix"
)

func marshal2292HopLimit(b []byte, cm *ControlMessage) []byte {
	m := socket.ControlMessage(b)
	m.MarshalHeader(iana.ProtocolIPv6, unix.IPV6_2292HOPLIMIT, 4)
	if cm != nil {
		socket.NativeEndian.PutUint32(m.Data(4), uint32(cm.HopLimit))
	}
	return m.Next(4)
}

func marshal2292PacketInfo(b []byte, cm *ControlMessage) []byte {
	m := socket.ControlMessage(b)
	m.MarshalHeader(iana.ProtocolIPv6, unix.IPV6_2292PKTINFO, sizeofInet6Pktinfo)
	if cm != nil {
		pi := (*inet6Pktinfo)(unsafe.Pointer(&m.Data(sizeofInet6Pktinfo)[0]))
		if ip := cm.Src.To16(); ip != nil && ip.To4() == nil {
			copy(pi.Addr[:], ip)
		}
		if cm.IfIndex > 0 {
			pi.setIfindex(cm.IfIndex)
		}
	}
	return m.Next(sizeofInet6Pktinfo)
}

func marshal2292NextHop(b []byte, cm *ControlMessage) []byte {
	m := socket.ControlMessage(b)
<<<<<<< HEAD
	m.MarshalHeader(iana.ProtocolIPv6, sysIPV6_2292NEXTHOP, sizeofSockaddrInet6)
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
	m.MarshalHeader(iana.ProtocolIPv6, sysIPV6_2292NEXTHOP, sizeofSockaddrInet6)
=======
	m.MarshalHeader(iana.ProtocolIPv6, unix.IPV6_2292NEXTHOP, sizeofSockaddrInet6)
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 6b7ce455e (update vendored files)
=======
//go:build darwin
>>>>>>> 6b7ce455e (update vendored files)
// +build darwin

package ipv6

import (
	"unsafe"

	"golang.org/x/net/internal/iana"
	"golang.org/x/net/internal/socket"

	"golang.org/x/sys/unix"
)

func marshal2292HopLimit(b []byte, cm *ControlMessage) []byte {
	m := socket.ControlMessage(b)
	m.MarshalHeader(iana.ProtocolIPv6, unix.IPV6_2292HOPLIMIT, 4)
	if cm != nil {
		socket.NativeEndian.PutUint32(m.Data(4), uint32(cm.HopLimit))
	}
	return m.Next(4)
}

func marshal2292PacketInfo(b []byte, cm *ControlMessage) []byte {
	m := socket.ControlMessage(b)
	m.MarshalHeader(iana.ProtocolIPv6, unix.IPV6_2292PKTINFO, sizeofInet6Pktinfo)
	if cm != nil {
		pi := (*inet6Pktinfo)(unsafe.Pointer(&m.Data(sizeofInet6Pktinfo)[0]))
		if ip := cm.Src.To16(); ip != nil && ip.To4() == nil {
			copy(pi.Addr[:], ip)
		}
		if cm.IfIndex > 0 {
			pi.setIfindex(cm.IfIndex)
		}
	}
	return m.Next(sizeofInet6Pktinfo)
}

func marshal2292NextHop(b []byte, cm *ControlMessage) []byte {
	m := socket.ControlMessage(b)
<<<<<<< HEAD
	m.MarshalHeader(iana.ProtocolIPv6, sysIPV6_2292NEXTHOP, sizeofSockaddrInet6)
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
	m.MarshalHeader(iana.ProtocolIPv6, sysIPV6_2292NEXTHOP, sizeofSockaddrInet6)
=======
	m.MarshalHeader(iana.ProtocolIPv6, unix.IPV6_2292NEXTHOP, sizeofSockaddrInet6)
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 4d7e5ad26 (update vendored files)
=======
//go:build darwin
>>>>>>> 4d7e5ad26 (update vendored files)
// +build darwin

package ipv6

import (
	"unsafe"

	"golang.org/x/net/internal/iana"
	"golang.org/x/net/internal/socket"

	"golang.org/x/sys/unix"
)

func marshal2292HopLimit(b []byte, cm *ControlMessage) []byte {
	m := socket.ControlMessage(b)
	m.MarshalHeader(iana.ProtocolIPv6, unix.IPV6_2292HOPLIMIT, 4)
	if cm != nil {
		socket.NativeEndian.PutUint32(m.Data(4), uint32(cm.HopLimit))
	}
	return m.Next(4)
}

func marshal2292PacketInfo(b []byte, cm *ControlMessage) []byte {
	m := socket.ControlMessage(b)
	m.MarshalHeader(iana.ProtocolIPv6, unix.IPV6_2292PKTINFO, sizeofInet6Pktinfo)
	if cm != nil {
		pi := (*inet6Pktinfo)(unsafe.Pointer(&m.Data(sizeofInet6Pktinfo)[0]))
		if ip := cm.Src.To16(); ip != nil && ip.To4() == nil {
			copy(pi.Addr[:], ip)
		}
		if cm.IfIndex > 0 {
			pi.setIfindex(cm.IfIndex)
		}
	}
	return m.Next(sizeofInet6Pktinfo)
}

func marshal2292NextHop(b []byte, cm *ControlMessage) []byte {
	m := socket.ControlMessage(b)
<<<<<<< HEAD
	m.MarshalHeader(iana.ProtocolIPv6, sysIPV6_2292NEXTHOP, sizeofSockaddrInet6)
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
	m.MarshalHeader(iana.ProtocolIPv6, sysIPV6_2292NEXTHOP, sizeofSockaddrInet6)
=======
	m.MarshalHeader(iana.ProtocolIPv6, unix.IPV6_2292NEXTHOP, sizeofSockaddrInet6)
>>>>>>> 4d7e5ad26 (update vendored files)
	if cm != nil {
		sa := (*sockaddrInet6)(unsafe.Pointer(&m.Data(sizeofSockaddrInet6)[0]))
		sa.setSockaddr(cm.NextHop, cm.IfIndex)
	}
	return m.Next(sizeofSockaddrInet6)
}
