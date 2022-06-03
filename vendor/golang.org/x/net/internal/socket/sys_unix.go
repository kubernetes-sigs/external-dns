// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
//go:build dragonfly || freebsd || (linux && !s390x && !386) || netbsd || openbsd
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
=======
//go:build dragonfly || freebsd || (linux && !s390x && !386) || netbsd || openbsd
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
// +build dragonfly freebsd linux,!s390x,!386 netbsd openbsd
||||||| parent of 6b7ce455e (update vendored files)
// +build dragonfly freebsd linux,!s390x,!386 netbsd openbsd
=======
//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris
>>>>>>> 6b7ce455e (update vendored files)

package socket

import (
	"net"
	"unsafe"

	"golang.org/x/sys/unix"
)

//go:linkname syscall_getsockopt syscall.getsockopt
func syscall_getsockopt(s, level, name int, val unsafe.Pointer, vallen *uint32) error

//go:linkname syscall_setsockopt syscall.setsockopt
func syscall_setsockopt(s, level, name int, val unsafe.Pointer, vallen uintptr) error

func getsockopt(s uintptr, level, name int, b []byte) (int, error) {
	l := uint32(len(b))
	err := syscall_getsockopt(int(s), level, name, unsafe.Pointer(&b[0]), &l)
	return int(l), err
}

func setsockopt(s uintptr, level, name int, b []byte) error {
	return syscall_setsockopt(int(s), level, name, unsafe.Pointer(&b[0]), uintptr(len(b)))
}

func recvmsg(s uintptr, buffers [][]byte, oob []byte, flags int, network string) (n, oobn int, recvflags int, from net.Addr, err error) {
	var unixFrom unix.Sockaddr
	n, oobn, recvflags, unixFrom, err = unix.RecvmsgBuffers(int(s), buffers, oob, flags)
	if unixFrom != nil {
		from = sockaddrToAddr(unixFrom, network)
	}
	return
}

<<<<<<< HEAD
func sendmsg(s uintptr, h *msghdr, flags int) (int, error) {
	return syscall_sendmsg(int(s), (*syscall.Msghdr)(unsafe.Pointer(h)), flags)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// +build dragonfly freebsd linux,!s390x,!386 netbsd openbsd
||||||| parent of 4d7e5ad26 (update vendored files)
// +build dragonfly freebsd linux,!s390x,!386 netbsd openbsd
=======
//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris
>>>>>>> 4d7e5ad26 (update vendored files)

package socket

import (
	"syscall"
	"unsafe"
)

//go:linkname syscall_getsockopt syscall.getsockopt
func syscall_getsockopt(s, level, name int, val unsafe.Pointer, vallen *uint32) error

//go:linkname syscall_setsockopt syscall.setsockopt
func syscall_setsockopt(s, level, name int, val unsafe.Pointer, vallen uintptr) error

//go:linkname syscall_recvmsg syscall.recvmsg
func syscall_recvmsg(s int, msg *syscall.Msghdr, flags int) (int, error)

//go:linkname syscall_sendmsg syscall.sendmsg
func syscall_sendmsg(s int, msg *syscall.Msghdr, flags int) (int, error)

func getsockopt(s uintptr, level, name int, b []byte) (int, error) {
	l := uint32(len(b))
	err := syscall_getsockopt(int(s), level, name, unsafe.Pointer(&b[0]), &l)
	return int(l), err
}

func setsockopt(s uintptr, level, name int, b []byte) error {
	return syscall_setsockopt(int(s), level, name, unsafe.Pointer(&b[0]), uintptr(len(b)))
}

func recvmsg(s uintptr, h *msghdr, flags int) (int, error) {
	return syscall_recvmsg(int(s), (*syscall.Msghdr)(unsafe.Pointer(h)), flags)
}

func sendmsg(s uintptr, h *msghdr, flags int) (int, error) {
<<<<<<< HEAD
	n, _, errno := syscall.Syscall(syscall.SYS_SENDMSG, s, uintptr(unsafe.Pointer(h)), uintptr(flags))
	return int(n), errnoErr(errno)
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
	n, _, errno := syscall.Syscall(syscall.SYS_SENDMSG, s, uintptr(unsafe.Pointer(h)), uintptr(flags))
	return int(n), errnoErr(errno)
=======
	return syscall_sendmsg(int(s), (*syscall.Msghdr)(unsafe.Pointer(h)), flags)
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of e1cd8261c (UPSTREAM: <carry>: update vendored files v0.13.1)
func sendmsg(s uintptr, h *msghdr, flags int) (int, error) {
	return syscall_sendmsg(int(s), (*syscall.Msghdr)(unsafe.Pointer(h)), flags)
=======
func sendmsg(s uintptr, buffers [][]byte, oob []byte, to net.Addr, flags int) (int, error) {
	var unixTo unix.Sockaddr
	if to != nil {
		unixTo = addrToSockaddr(to)
	}
	return unix.SendmsgBuffers(int(s), buffers, oob, unixTo, flags)
}

// addrToSockaddr converts a net.Addr to a unix.Sockaddr.
func addrToSockaddr(a net.Addr) unix.Sockaddr {
	var (
		ip   net.IP
		port int
		zone string
	)
	switch a := a.(type) {
	case *net.TCPAddr:
		ip = a.IP
		port = a.Port
		zone = a.Zone
	case *net.UDPAddr:
		ip = a.IP
		port = a.Port
		zone = a.Zone
	case *net.IPAddr:
		ip = a.IP
		zone = a.Zone
	default:
		return nil
	}

	if ip4 := ip.To4(); ip4 != nil {
		sa := unix.SockaddrInet4{Port: port}
		copy(sa.Addr[:], ip4)
		return &sa
	}

	if ip6 := ip.To16(); ip6 != nil && ip.To4() == nil {
		sa := unix.SockaddrInet6{Port: port}
		copy(sa.Addr[:], ip6)
		if zone != "" {
			sa.ZoneId = uint32(zoneCache.index(zone))
		}
		return &sa
	}

	return nil
}

// sockaddrToAddr converts a unix.Sockaddr to a net.Addr.
func sockaddrToAddr(sa unix.Sockaddr, network string) net.Addr {
	var (
		ip   net.IP
		port int
		zone string
	)
	switch sa := sa.(type) {
	case *unix.SockaddrInet4:
		ip = make(net.IP, net.IPv4len)
		copy(ip, sa.Addr[:])
		port = sa.Port
	case *unix.SockaddrInet6:
		ip = make(net.IP, net.IPv6len)
		copy(ip, sa.Addr[:])
		port = sa.Port
		if sa.ZoneId > 0 {
			zone = zoneCache.name(int(sa.ZoneId))
		}
	default:
		return nil
	}

	switch network {
	case "tcp", "tcp4", "tcp6":
		return &net.TCPAddr{IP: ip, Port: port, Zone: zone}
	case "udp", "udp4", "udp6":
		return &net.UDPAddr{IP: ip, Port: port, Zone: zone}
	default:
		return &net.IPAddr{IP: ip, Zone: zone}
	}
>>>>>>> e1cd8261c (UPSTREAM: <carry>: update vendored files v0.13.1)
}
