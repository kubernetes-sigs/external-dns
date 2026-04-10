// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http2

import "math"

// NewRandomWriteScheduler constructs a WriteScheduler that ignores HTTP/2
// priorities. Control frames like SETTINGS and PING are written before DATA
// frames, but if no control frames are queued and multiple streams have queued
// HEADERS or DATA frames, Pop selects a ready stream arbitrarily.
//
// Deprecated: User-provided write schedulers are deprecated.
func NewRandomWriteScheduler() WriteScheduler {
	return &randomWriteScheduler{sq: make(map[uint32]*writeQueue)}
}

type randomWriteScheduler struct {
	// zero are frames not associated with a specific stream.
	zero writeQueue

	// sq contains the stream-specific queues, keyed by stream ID.
	// When a stream is idle, closed, or emptied, it's deleted
	// from the map.
	sq map[uint32]*writeQueue

	// pool of empty queues for reuse.
	queuePool writeQueuePool
}

func (ws *randomWriteScheduler) OpenStream(streamID uint32, options OpenStreamOptions) {
	// no-op: idle streams are not tracked
}

func (ws *randomWriteScheduler) CloseStream(streamID uint32) {
	q, ok := ws.sq[streamID]
	if !ok {
		return
	}
	delete(ws.sq, streamID)
	ws.queuePool.put(q)
}

func (ws *randomWriteScheduler) AdjustStream(streamID uint32, priority PriorityParam) {
	// no-op: priorities are ignored
}

func (ws *randomWriteScheduler) Push(wr FrameWriteRequest) {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	if wr.isControl() {
		ws.zero.push(wr)
		return
	}
	id := wr.StreamID()
	q, ok := ws.sq[id]
	if !ok {
		q = ws.queuePool.get()
		ws.sq[id] = q
	}
	q.push(wr)
}

func (ws *randomWriteScheduler) Pop() (FrameWriteRequest, bool) {
	// Control and RST_STREAM frames first.
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	id := wr.StreamID()
	if id == 0 {
||||||| parent of 4d7e5ad26 (update vendored files)
	id := wr.StreamID()
	if id == 0 {
=======
	if wr.isControl() {
>>>>>>> 4d7e5ad26 (update vendored files)
		ws.zero.push(wr)
		return
	}
	id := wr.StreamID()
	q, ok := ws.sq[id]
	if !ok {
		q = ws.queuePool.get()
		ws.sq[id] = q
	}
	q.push(wr)
}

func (ws *randomWriteScheduler) Pop() (FrameWriteRequest, bool) {
<<<<<<< HEAD
	// Control frames first.
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
	// Control frames first.
=======
	// Control and RST_STREAM frames first.
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	id := wr.StreamID()
	if id == 0 {
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	id := wr.StreamID()
	if id == 0 {
=======
	if wr.isControl() {
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
		ws.zero.push(wr)
		return
	}
	id := wr.StreamID()
	q, ok := ws.sq[id]
	if !ok {
		q = ws.queuePool.get()
		ws.sq[id] = q
	}
	q.push(wr)
}

func (ws *randomWriteScheduler) Pop() (FrameWriteRequest, bool) {
<<<<<<< HEAD
	// Control frames first.
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	// Control frames first.
=======
	// Control and RST_STREAM frames first.
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	if !ws.zero.empty() {
		return ws.zero.shift(), true
	}
	// Iterate over all non-idle streams until finding one that can be consumed.
	for streamID, q := range ws.sq {
		if wr, ok := q.consume(math.MaxInt32); ok {
			if q.empty() {
				delete(ws.sq, streamID)
				ws.queuePool.put(q)
			}
			return wr, true
		}
	}
	return FrameWriteRequest{}, false
}
