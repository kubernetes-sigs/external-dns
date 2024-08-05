<<<<<<< HEAD
// +build !go1.11 !aix,!darwin,!dragonfly,!freebsd,!linux,!netbsd,!openbsd

package dns

import "net"

const supportsReusePort = false

func listenTCP(network, addr string, reuseport bool) (net.Listener, error) {
	if reuseport {
		// TODO(tmthrgd): return an error?
	}

	return net.Listen(network, addr)
}

func listenUDP(network, addr string, reuseport bool) (net.PacketConn, error) {
	if reuseport {
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
//go:build !aix && !darwin && !dragonfly && !freebsd && !linux && !netbsd && !openbsd
// +build !aix,!darwin,!dragonfly,!freebsd,!linux,!netbsd,!openbsd

package dns

import "net"

const supportsReusePort = false

func listenTCP(network, addr string, reuseport, reuseaddr bool) (net.Listener, error) {
	if reuseport || reuseaddr {
		// TODO(tmthrgd): return an error?
	}

	return net.Listen(network, addr)
}

const supportsReuseAddr = false

func listenUDP(network, addr string, reuseport, reuseaddr bool) (net.PacketConn, error) {
	if reuseport || reuseaddr {
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
		// TODO(tmthrgd): return an error?
	}

	return net.ListenPacket(network, addr)
}
