/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package connrotation implements a connection dialer that tracks and can close
// all created connections.
//
// This is used for credential rotation of long-lived connections, when there's
// no way to re-authenticate on a live connection.
package connrotation

import (
	"context"
	"net"
	"sync"
)

// DialFunc is a shorthand for signature of net.DialContext.
type DialFunc func(ctx context.Context, network, address string) (net.Conn, error)

// Dialer opens connections through Dial and tracks them.
type Dialer struct {
	dial DialFunc
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	*ConnectionTracker
}

// NewDialer creates a new Dialer instance.
// Equivalent to NewDialerWithTracker(dial, nil).
func NewDialer(dial DialFunc) *Dialer {
	return NewDialerWithTracker(dial, nil)
}

// NewDialerWithTracker creates a new Dialer instance.
//
// If dial is not nil, it will be used to create new underlying connections.
// Otherwise net.DialContext is used.
// If tracker is not nil, it will be used to track new underlying connections.
// Otherwise NewConnectionTracker() is used.
func NewDialerWithTracker(dial DialFunc, tracker *ConnectionTracker) *Dialer {
	if tracker == nil {
		tracker = NewConnectionTracker()
	}
	return &Dialer{
		dial:              dial,
		ConnectionTracker: tracker,
	}
}

// ConnectionTracker keeps track of opened connections
type ConnectionTracker struct {
	mu    sync.Mutex
	conns map[*closableConn]struct{}
}

// NewConnectionTracker returns a connection tracker for use with NewDialerWithTracker
func NewConnectionTracker() *ConnectionTracker {
	return &ConnectionTracker{
		conns: make(map[*closableConn]struct{}),
	}
}

// CloseAll forcibly closes all tracked connections.
//
// Note: new connections may get created before CloseAll returns.
func (c *ConnectionTracker) CloseAll() {
	c.mu.Lock()
	conns := c.conns
	c.conns = make(map[*closableConn]struct{})
	c.mu.Unlock()

	for conn := range conns {
		conn.Close()
	}
}

// Track adds the connection to the list of tracked connections,
// and returns a wrapped copy of the connection that stops tracking the connection
// when it is closed.
func (c *ConnectionTracker) Track(conn net.Conn) net.Conn {
	closable := &closableConn{Conn: conn}

	// When the connection is closed, remove it from the map. This will
	// be no-op if the connection isn't in the map, e.g. if CloseAll()
	// is called.
	closable.onClose = func() {
		c.mu.Lock()
		delete(c.conns, closable)
		c.mu.Unlock()
	}

	// Start tracking the connection
	c.mu.Lock()
	c.conns[closable] = struct{}{}
	c.mu.Unlock()

	return closable
}

// Dial creates a new tracked connection.
func (d *Dialer) Dial(network, address string) (net.Conn, error) {
	return d.DialContext(context.Background(), network, address)
}

// DialContext creates a new tracked connection.
func (d *Dialer) DialContext(ctx context.Context, network, address string) (net.Conn, error) {
	conn, err := d.dial(ctx, network, address)
	if err != nil {
		return nil, err
	}
	return d.ConnectionTracker.Track(conn), nil
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 5ce8c7613 (update vendored files)
=======
	*ConnectionTracker
}
>>>>>>> 5ce8c7613 (update vendored files)

// NewDialer creates a new Dialer instance.
// Equivalent to NewDialerWithTracker(dial, nil).
func NewDialer(dial DialFunc) *Dialer {
	return NewDialerWithTracker(dial, nil)
}

// NewDialerWithTracker creates a new Dialer instance.
//
// If dial is not nil, it will be used to create new underlying connections.
// Otherwise net.DialContext is used.
// If tracker is not nil, it will be used to track new underlying connections.
// Otherwise NewConnectionTracker() is used.
func NewDialerWithTracker(dial DialFunc, tracker *ConnectionTracker) *Dialer {
	if tracker == nil {
		tracker = NewConnectionTracker()
	}
	return &Dialer{
		dial:              dial,
		ConnectionTracker: tracker,
	}
}

// ConnectionTracker keeps track of opened connections
type ConnectionTracker struct {
	mu    sync.Mutex
	conns map[*closableConn]struct{}
}

// NewConnectionTracker returns a connection tracker for use with NewDialerWithTracker
func NewConnectionTracker() *ConnectionTracker {
	return &ConnectionTracker{
		conns: make(map[*closableConn]struct{}),
	}
}

// CloseAll forcibly closes all tracked connections.
//
// Note: new connections may get created before CloseAll returns.
func (c *ConnectionTracker) CloseAll() {
	c.mu.Lock()
	conns := c.conns
	c.conns = make(map[*closableConn]struct{})
	c.mu.Unlock()

	for conn := range conns {
		conn.Close()
	}
}

// Track adds the connection to the list of tracked connections,
// and returns a wrapped copy of the connection that stops tracking the connection
// when it is closed.
func (c *ConnectionTracker) Track(conn net.Conn) net.Conn {
	closable := &closableConn{Conn: conn}

	// When the connection is closed, remove it from the map. This will
	// be no-op if the connection isn't in the map, e.g. if CloseAll()
	// is called.
	closable.onClose = func() {
		c.mu.Lock()
		delete(c.conns, closable)
		c.mu.Unlock()
	}

	// Start tracking the connection
	c.mu.Lock()
	c.conns[closable] = struct{}{}
	c.mu.Unlock()

	return closable
}

// Dial creates a new tracked connection.
func (d *Dialer) Dial(network, address string) (net.Conn, error) {
	return d.DialContext(context.Background(), network, address)
}

// DialContext creates a new tracked connection.
func (d *Dialer) DialContext(ctx context.Context, network, address string) (net.Conn, error) {
	conn, err := d.dial(ctx, network, address)
	if err != nil {
		return nil, err
	}
<<<<<<< HEAD

	closable := &closableConn{Conn: conn}

	// When the connection is closed, remove it from the map. This will
	// be no-op if the connection isn't in the map, e.g. if CloseAll()
	// is called.
	closable.onClose = func() {
		d.mu.Lock()
		delete(d.conns, closable)
		d.mu.Unlock()
	}

	// Start tracking the connection
	d.mu.Lock()
	d.conns[closable] = struct{}{}
	d.mu.Unlock()

	return closable, nil
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)

	closable := &closableConn{Conn: conn}

	// When the connection is closed, remove it from the map. This will
	// be no-op if the connection isn't in the map, e.g. if CloseAll()
	// is called.
	closable.onClose = func() {
		d.mu.Lock()
		delete(d.conns, closable)
		d.mu.Unlock()
	}

	// Start tracking the connection
	d.mu.Lock()
	d.conns[closable] = struct{}{}
	d.mu.Unlock()

	return closable, nil
=======
	return d.ConnectionTracker.Track(conn), nil
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 6b7ce455e (update vendored files)
=======
	*ConnectionTracker
}
>>>>>>> 6b7ce455e (update vendored files)

// NewDialer creates a new Dialer instance.
// Equivalent to NewDialerWithTracker(dial, nil).
func NewDialer(dial DialFunc) *Dialer {
	return NewDialerWithTracker(dial, nil)
}

// NewDialerWithTracker creates a new Dialer instance.
//
// If dial is not nil, it will be used to create new underlying connections.
// Otherwise net.DialContext is used.
// If tracker is not nil, it will be used to track new underlying connections.
// Otherwise NewConnectionTracker() is used.
func NewDialerWithTracker(dial DialFunc, tracker *ConnectionTracker) *Dialer {
	if tracker == nil {
		tracker = NewConnectionTracker()
	}
	return &Dialer{
		dial:              dial,
		ConnectionTracker: tracker,
	}
}

// ConnectionTracker keeps track of opened connections
type ConnectionTracker struct {
	mu    sync.Mutex
	conns map[*closableConn]struct{}
}

// NewConnectionTracker returns a connection tracker for use with NewDialerWithTracker
func NewConnectionTracker() *ConnectionTracker {
	return &ConnectionTracker{
		conns: make(map[*closableConn]struct{}),
	}
}

// CloseAll forcibly closes all tracked connections.
//
// Note: new connections may get created before CloseAll returns.
func (c *ConnectionTracker) CloseAll() {
	c.mu.Lock()
	conns := c.conns
	c.conns = make(map[*closableConn]struct{})
	c.mu.Unlock()

	for conn := range conns {
		conn.Close()
	}
}

// Track adds the connection to the list of tracked connections,
// and returns a wrapped copy of the connection that stops tracking the connection
// when it is closed.
func (c *ConnectionTracker) Track(conn net.Conn) net.Conn {
	closable := &closableConn{Conn: conn}

	// When the connection is closed, remove it from the map. This will
	// be no-op if the connection isn't in the map, e.g. if CloseAll()
	// is called.
	closable.onClose = func() {
		c.mu.Lock()
		delete(c.conns, closable)
		c.mu.Unlock()
	}

	// Start tracking the connection
	c.mu.Lock()
	c.conns[closable] = struct{}{}
	c.mu.Unlock()

	return closable
}

// Dial creates a new tracked connection.
func (d *Dialer) Dial(network, address string) (net.Conn, error) {
	return d.DialContext(context.Background(), network, address)
}

// DialContext creates a new tracked connection.
func (d *Dialer) DialContext(ctx context.Context, network, address string) (net.Conn, error) {
	conn, err := d.dial(ctx, network, address)
	if err != nil {
		return nil, err
	}
<<<<<<< HEAD

	closable := &closableConn{Conn: conn}

	// When the connection is closed, remove it from the map. This will
	// be no-op if the connection isn't in the map, e.g. if CloseAll()
	// is called.
	closable.onClose = func() {
		d.mu.Lock()
		delete(d.conns, closable)
		d.mu.Unlock()
	}

	// Start tracking the connection
	d.mu.Lock()
	d.conns[closable] = struct{}{}
	d.mu.Unlock()

	return closable, nil
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)

	closable := &closableConn{Conn: conn}

	// When the connection is closed, remove it from the map. This will
	// be no-op if the connection isn't in the map, e.g. if CloseAll()
	// is called.
	closable.onClose = func() {
		d.mu.Lock()
		delete(d.conns, closable)
		d.mu.Unlock()
	}

	// Start tracking the connection
	d.mu.Lock()
	d.conns[closable] = struct{}{}
	d.mu.Unlock()

	return closable, nil
=======
	return d.ConnectionTracker.Track(conn), nil
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 4d7e5ad26 (update vendored files)
=======
	*ConnectionTracker
}
>>>>>>> 4d7e5ad26 (update vendored files)

// NewDialer creates a new Dialer instance.
// Equivalent to NewDialerWithTracker(dial, nil).
func NewDialer(dial DialFunc) *Dialer {
	return NewDialerWithTracker(dial, nil)
}

// NewDialerWithTracker creates a new Dialer instance.
//
// If dial is not nil, it will be used to create new underlying connections.
// Otherwise net.DialContext is used.
// If tracker is not nil, it will be used to track new underlying connections.
// Otherwise NewConnectionTracker() is used.
func NewDialerWithTracker(dial DialFunc, tracker *ConnectionTracker) *Dialer {
	if tracker == nil {
		tracker = NewConnectionTracker()
	}
	return &Dialer{
		dial:              dial,
		ConnectionTracker: tracker,
	}
}

// ConnectionTracker keeps track of opened connections
type ConnectionTracker struct {
	mu    sync.Mutex
	conns map[*closableConn]struct{}
}

// NewConnectionTracker returns a connection tracker for use with NewDialerWithTracker
func NewConnectionTracker() *ConnectionTracker {
	return &ConnectionTracker{
		conns: make(map[*closableConn]struct{}),
	}
}

// CloseAll forcibly closes all tracked connections.
//
// Note: new connections may get created before CloseAll returns.
func (c *ConnectionTracker) CloseAll() {
	c.mu.Lock()
	conns := c.conns
	c.conns = make(map[*closableConn]struct{})
	c.mu.Unlock()

	for conn := range conns {
		conn.Close()
	}
}

// Track adds the connection to the list of tracked connections,
// and returns a wrapped copy of the connection that stops tracking the connection
// when it is closed.
func (c *ConnectionTracker) Track(conn net.Conn) net.Conn {
	closable := &closableConn{Conn: conn}

	// When the connection is closed, remove it from the map. This will
	// be no-op if the connection isn't in the map, e.g. if CloseAll()
	// is called.
	closable.onClose = func() {
		c.mu.Lock()
		delete(c.conns, closable)
		c.mu.Unlock()
	}

	// Start tracking the connection
	c.mu.Lock()
	c.conns[closable] = struct{}{}
	c.mu.Unlock()

	return closable
}

// Dial creates a new tracked connection.
func (d *Dialer) Dial(network, address string) (net.Conn, error) {
	return d.DialContext(context.Background(), network, address)
}

// DialContext creates a new tracked connection.
func (d *Dialer) DialContext(ctx context.Context, network, address string) (net.Conn, error) {
	conn, err := d.dial(ctx, network, address)
	if err != nil {
		return nil, err
	}
<<<<<<< HEAD

	closable := &closableConn{Conn: conn}

	// When the connection is closed, remove it from the map. This will
	// be no-op if the connection isn't in the map, e.g. if CloseAll()
	// is called.
	closable.onClose = func() {
		d.mu.Lock()
		delete(d.conns, closable)
		d.mu.Unlock()
	}

	// Start tracking the connection
	d.mu.Lock()
	d.conns[closable] = struct{}{}
	d.mu.Unlock()

	return closable, nil
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)

	closable := &closableConn{Conn: conn}

	// When the connection is closed, remove it from the map. This will
	// be no-op if the connection isn't in the map, e.g. if CloseAll()
	// is called.
	closable.onClose = func() {
		d.mu.Lock()
		delete(d.conns, closable)
		d.mu.Unlock()
	}

	// Start tracking the connection
	d.mu.Lock()
	d.conns[closable] = struct{}{}
	d.mu.Unlock()

	return closable, nil
=======
	return d.ConnectionTracker.Track(conn), nil
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======

	mu    sync.Mutex
	conns map[*closableConn]struct{}
}

// NewDialer creates a new Dialer instance.
//
// If dial is not nil, it will be used to create new underlying connections.
// Otherwise net.DialContext is used.
func NewDialer(dial DialFunc) *Dialer {
	return &Dialer{
		dial:  dial,
		conns: make(map[*closableConn]struct{}),
	}
}

// CloseAll forcibly closes all tracked connections.
//
// Note: new connections may get created before CloseAll returns.
func (d *Dialer) CloseAll() {
	d.mu.Lock()
	conns := d.conns
	d.conns = make(map[*closableConn]struct{})
	d.mu.Unlock()

	for conn := range conns {
		conn.Close()
	}
}

// Dial creates a new tracked connection.
func (d *Dialer) Dial(network, address string) (net.Conn, error) {
	return d.DialContext(context.Background(), network, address)
}

// DialContext creates a new tracked connection.
func (d *Dialer) DialContext(ctx context.Context, network, address string) (net.Conn, error) {
	conn, err := d.dial(ctx, network, address)
	if err != nil {
		return nil, err
	}

	closable := &closableConn{Conn: conn}

	// When the connection is closed, remove it from the map. This will
	// be no-op if the connection isn't in the map, e.g. if CloseAll()
	// is called.
	closable.onClose = func() {
		d.mu.Lock()
		delete(d.conns, closable)
		d.mu.Unlock()
	}

	// Start tracking the connection
	d.mu.Lock()
	d.conns[closable] = struct{}{}
	d.mu.Unlock()

	return closable, nil
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
}

type closableConn struct {
	onClose func()
	net.Conn
}

func (c *closableConn) Close() error {
	go c.onClose()
	return c.Conn.Close()
}
