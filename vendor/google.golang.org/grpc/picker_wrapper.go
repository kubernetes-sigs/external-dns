/*
 *
 * Copyright 2017 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package grpc

import (
	"context"
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	"io"
	"sync"

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/internal/channelz"
	"google.golang.org/grpc/internal/transport"
	"google.golang.org/grpc/status"
)

// pickerWrapper is a wrapper of balancer.Picker. It blocks on certain pick
// actions and unblock when there's a picker update.
type pickerWrapper struct {
	mu         sync.Mutex
	done       bool
	blockingCh chan struct{}
	picker     balancer.Picker
}

func newPickerWrapper() *pickerWrapper {
	return &pickerWrapper{blockingCh: make(chan struct{})}
}

// updatePicker is called by UpdateBalancerState. It unblocks all blocked pick.
func (pw *pickerWrapper) updatePicker(p balancer.Picker) {
	pw.mu.Lock()
	if pw.done {
		pw.mu.Unlock()
		return
	}
	pw.picker = p
	// pw.blockingCh should never be nil.
	close(pw.blockingCh)
	pw.blockingCh = make(chan struct{})
	pw.mu.Unlock()
}

func doneChannelzWrapper(acw *acBalancerWrapper, done func(balancer.DoneInfo)) func(balancer.DoneInfo) {
	acw.mu.Lock()
	ac := acw.ac
	acw.mu.Unlock()
	ac.incrCallsStarted()
	return func(b balancer.DoneInfo) {
		if b.Err != nil && b.Err != io.EOF {
			ac.incrCallsFailed()
		} else {
			ac.incrCallsSucceeded()
		}
		if done != nil {
			done(b)
		}
	}
}

// pick returns the transport that will be used for the RPC.
// It may block in the following cases:
// - there's no picker
// - the current picker returns ErrNoSubConnAvailable
// - the current picker returns other errors and failfast is false.
// - the subConn returned by the current picker is not READY
// When one of these situations happens, pick blocks until the picker gets updated.
func (pw *pickerWrapper) pick(ctx context.Context, failfast bool, info balancer.PickInfo) (transport.ClientTransport, func(balancer.DoneInfo), error) {
	var ch chan struct{}

	var lastPickErr error
	for {
		pw.mu.Lock()
		if pw.done {
			pw.mu.Unlock()
			return nil, nil, ErrClientConnClosing
		}

		if pw.picker == nil {
			ch = pw.blockingCh
		}
		if ch == pw.blockingCh {
			// This could happen when either:
			// - pw.picker is nil (the previous if condition), or
			// - has called pick on the current picker.
			pw.mu.Unlock()
			select {
			case <-ctx.Done():
				var errStr string
				if lastPickErr != nil {
					errStr = "latest balancer error: " + lastPickErr.Error()
				} else {
					errStr = ctx.Err().Error()
				}
				switch ctx.Err() {
				case context.DeadlineExceeded:
					return nil, nil, status.Error(codes.DeadlineExceeded, errStr)
				case context.Canceled:
					return nil, nil, status.Error(codes.Canceled, errStr)
				}
			case <-ch:
			}
			continue
		}

		ch = pw.blockingCh
		p := pw.picker
		pw.mu.Unlock()

		pickResult, err := p.Pick(info)

		if err != nil {
			if err == balancer.ErrNoSubConnAvailable {
				continue
			}
			if _, ok := status.FromError(err); ok {
				// Status error: end the RPC unconditionally with this status.
				return nil, nil, err
			}
			// For all other errors, wait for ready RPCs should block and other
			// RPCs should fail with unavailable.
			if !failfast {
				lastPickErr = err
				continue
			}
			return nil, nil, status.Error(codes.Unavailable, err.Error())
		}

		acw, ok := pickResult.SubConn.(*acBalancerWrapper)
		if !ok {
			logger.Error("subconn returned from pick is not *acBalancerWrapper")
			continue
		}
		if t := acw.getAddrConn().getReadyTransport(); t != nil {
			if channelz.IsOn() {
				return t, doneChannelzWrapper(acw, pickResult.Done), nil
			}
			return t, pickResult.Done, nil
		}
		if pickResult.Done != nil {
			// Calling done with nil error, no bytes sent and no bytes received.
			// DoneInfo with default value works.
			pickResult.Done(balancer.DoneInfo{})
		}
		logger.Infof("blockingPicker: the picked transport is not ready, loop back to repick")
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	"fmt"
||||||| parent of 5ce8c7613 (update vendored files)
	"fmt"
=======
>>>>>>> 5ce8c7613 (update vendored files)
	"io"
	"sync"

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/internal/channelz"
	"google.golang.org/grpc/internal/transport"
	"google.golang.org/grpc/status"
)

// pickerWrapper is a wrapper of balancer.Picker. It blocks on certain pick
// actions and unblock when there's a picker update.
type pickerWrapper struct {
	mu         sync.Mutex
	done       bool
	blockingCh chan struct{}
	picker     balancer.Picker
}

func newPickerWrapper() *pickerWrapper {
	return &pickerWrapper{blockingCh: make(chan struct{})}
}

// updatePicker is called by UpdateBalancerState. It unblocks all blocked pick.
func (pw *pickerWrapper) updatePicker(p balancer.Picker) {
	pw.mu.Lock()
	if pw.done {
		pw.mu.Unlock()
		return
	}
	pw.picker = p
	// pw.blockingCh should never be nil.
	close(pw.blockingCh)
	pw.blockingCh = make(chan struct{})
	pw.mu.Unlock()
}

func doneChannelzWrapper(acw *acBalancerWrapper, done func(balancer.DoneInfo)) func(balancer.DoneInfo) {
	acw.mu.Lock()
	ac := acw.ac
	acw.mu.Unlock()
	ac.incrCallsStarted()
	return func(b balancer.DoneInfo) {
		if b.Err != nil && b.Err != io.EOF {
			ac.incrCallsFailed()
		} else {
			ac.incrCallsSucceeded()
		}
		if done != nil {
			done(b)
		}
	}
}

// pick returns the transport that will be used for the RPC.
// It may block in the following cases:
// - there's no picker
// - the current picker returns ErrNoSubConnAvailable
// - the current picker returns other errors and failfast is false.
// - the subConn returned by the current picker is not READY
// When one of these situations happens, pick blocks until the picker gets updated.
func (pw *pickerWrapper) pick(ctx context.Context, failfast bool, info balancer.PickInfo) (transport.ClientTransport, func(balancer.DoneInfo), error) {
	var ch chan struct{}

	var lastPickErr error
	for {
		pw.mu.Lock()
		if pw.done {
			pw.mu.Unlock()
			return nil, nil, ErrClientConnClosing
		}

		if pw.picker == nil {
			ch = pw.blockingCh
		}
		if ch == pw.blockingCh {
			// This could happen when either:
			// - pw.picker is nil (the previous if condition), or
			// - has called pick on the current picker.
			pw.mu.Unlock()
			select {
			case <-ctx.Done():
				var errStr string
				if lastPickErr != nil {
					errStr = "latest balancer error: " + lastPickErr.Error()
				} else {
					errStr = ctx.Err().Error()
				}
				switch ctx.Err() {
				case context.DeadlineExceeded:
					return nil, nil, status.Error(codes.DeadlineExceeded, errStr)
				case context.Canceled:
					return nil, nil, status.Error(codes.Canceled, errStr)
				}
			case <-ch:
			}
			continue
		}

		ch = pw.blockingCh
		p := pw.picker
		pw.mu.Unlock()

		pickResult, err := p.Pick(info)

		if err != nil {
			if err == balancer.ErrNoSubConnAvailable {
				continue
			}
			if _, ok := status.FromError(err); ok {
				// Status error: end the RPC unconditionally with this status.
				return nil, nil, err
			}
			// For all other errors, wait for ready RPCs should block and other
			// RPCs should fail with unavailable.
			if !failfast {
				lastPickErr = err
				continue
			}
			return nil, nil, status.Error(codes.Unavailable, err.Error())
		}

		acw, ok := pickResult.SubConn.(*acBalancerWrapper)
		if !ok {
			logger.Error("subconn returned from pick is not *acBalancerWrapper")
			continue
		}
		if t := acw.getAddrConn().getReadyTransport(); t != nil {
			if channelz.IsOn() {
				return t, doneChannelzWrapper(acw, pickResult.Done), nil
			}
			return t, pickResult.Done, nil
		}
		if pickResult.Done != nil {
			// Calling done with nil error, no bytes sent and no bytes received.
			// DoneInfo with default value works.
			pickResult.Done(balancer.DoneInfo{})
		}
<<<<<<< HEAD
		grpclog.Infof("blockingPicker: the picked transport is not ready, loop back to repick")
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
		grpclog.Infof("blockingPicker: the picked transport is not ready, loop back to repick")
=======
		logger.Infof("blockingPicker: the picked transport is not ready, loop back to repick")
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	"fmt"
||||||| parent of 6b7ce455e (update vendored files)
	"fmt"
=======
>>>>>>> 6b7ce455e (update vendored files)
	"io"
	"sync"

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/internal/channelz"
	istatus "google.golang.org/grpc/internal/status"
	"google.golang.org/grpc/internal/transport"
	"google.golang.org/grpc/status"
)

// pickerWrapper is a wrapper of balancer.Picker. It blocks on certain pick
// actions and unblock when there's a picker update.
type pickerWrapper struct {
	mu         sync.Mutex
	done       bool
	idle       bool
	blockingCh chan struct{}
	picker     balancer.Picker
}

func newPickerWrapper() *pickerWrapper {
	return &pickerWrapper{blockingCh: make(chan struct{})}
}

// updatePicker is called by UpdateBalancerState. It unblocks all blocked pick.
func (pw *pickerWrapper) updatePicker(p balancer.Picker) {
	pw.mu.Lock()
	if pw.done || pw.idle {
		// There is a small window where a picker update from the LB policy can
		// race with the channel going to idle mode. If the picker is idle here,
		// it is because the channel asked it to do so, and therefore it is sage
		// to ignore the update from the LB policy.
		pw.mu.Unlock()
		return
	}
	pw.picker = p
	// pw.blockingCh should never be nil.
	close(pw.blockingCh)
	pw.blockingCh = make(chan struct{})
	pw.mu.Unlock()
}

// doneChannelzWrapper performs the following:
//   - increments the calls started channelz counter
//   - wraps the done function in the passed in result to increment the calls
//     failed or calls succeeded channelz counter before invoking the actual
//     done function.
func doneChannelzWrapper(acbw *acBalancerWrapper, result *balancer.PickResult) {
	ac := acbw.ac
	ac.incrCallsStarted()
	done := result.Done
	result.Done = func(b balancer.DoneInfo) {
		if b.Err != nil && b.Err != io.EOF {
			ac.incrCallsFailed()
		} else {
			ac.incrCallsSucceeded()
		}
		if done != nil {
			done(b)
		}
	}
}

// pick returns the transport that will be used for the RPC.
// It may block in the following cases:
// - there's no picker
// - the current picker returns ErrNoSubConnAvailable
// - the current picker returns other errors and failfast is false.
// - the subConn returned by the current picker is not READY
// When one of these situations happens, pick blocks until the picker gets updated.
func (pw *pickerWrapper) pick(ctx context.Context, failfast bool, info balancer.PickInfo) (transport.ClientTransport, balancer.PickResult, error) {
	var ch chan struct{}

	var lastPickErr error
	for {
		pw.mu.Lock()
		if pw.done {
			pw.mu.Unlock()
			return nil, balancer.PickResult{}, ErrClientConnClosing
		}

		if pw.picker == nil {
			ch = pw.blockingCh
		}
		if ch == pw.blockingCh {
			// This could happen when either:
			// - pw.picker is nil (the previous if condition), or
			// - has called pick on the current picker.
			pw.mu.Unlock()
			select {
			case <-ctx.Done():
				var errStr string
				if lastPickErr != nil {
					errStr = "latest balancer error: " + lastPickErr.Error()
				} else {
					errStr = ctx.Err().Error()
				}
				switch ctx.Err() {
				case context.DeadlineExceeded:
					return nil, balancer.PickResult{}, status.Error(codes.DeadlineExceeded, errStr)
				case context.Canceled:
					return nil, balancer.PickResult{}, status.Error(codes.Canceled, errStr)
				}
			case <-ch:
			}
			continue
		}

		ch = pw.blockingCh
		p := pw.picker
		pw.mu.Unlock()

		pickResult, err := p.Pick(info)
		if err != nil {
			if err == balancer.ErrNoSubConnAvailable {
				continue
			}
			if st, ok := status.FromError(err); ok {
				// Status error: end the RPC unconditionally with this status.
				// First restrict the code to the list allowed by gRFC A54.
				if istatus.IsRestrictedControlPlaneCode(st) {
					err = status.Errorf(codes.Internal, "received picker error with illegal status: %v", err)
				}
				return nil, balancer.PickResult{}, dropError{error: err}
			}
			// For all other errors, wait for ready RPCs should block and other
			// RPCs should fail with unavailable.
			if !failfast {
				lastPickErr = err
				continue
			}
			return nil, balancer.PickResult{}, status.Error(codes.Unavailable, err.Error())
		}

		acbw, ok := pickResult.SubConn.(*acBalancerWrapper)
		if !ok {
			logger.Errorf("subconn returned from pick is type %T, not *acBalancerWrapper", pickResult.SubConn)
			continue
		}
		if t := acbw.ac.getReadyTransport(); t != nil {
			if channelz.IsOn() {
				doneChannelzWrapper(acbw, &pickResult)
				return t, pickResult, nil
			}
			return t, pickResult, nil
		}
		if pickResult.Done != nil {
			// Calling done with nil error, no bytes sent and no bytes received.
			// DoneInfo with default value works.
			pickResult.Done(balancer.DoneInfo{})
		}
<<<<<<< HEAD
		grpclog.Infof("blockingPicker: the picked transport is not ready, loop back to repick")
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
		grpclog.Infof("blockingPicker: the picked transport is not ready, loop back to repick")
=======
		logger.Infof("blockingPicker: the picked transport is not ready, loop back to repick")
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	"fmt"
||||||| parent of 4d7e5ad26 (update vendored files)
	"fmt"
=======
>>>>>>> 4d7e5ad26 (update vendored files)
	"io"
	"sync"

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/internal/channelz"
	"google.golang.org/grpc/internal/transport"
	"google.golang.org/grpc/status"
)

// pickerWrapper is a wrapper of balancer.Picker. It blocks on certain pick
// actions and unblock when there's a picker update.
type pickerWrapper struct {
	mu         sync.Mutex
	done       bool
	blockingCh chan struct{}
	picker     balancer.Picker
}

func newPickerWrapper() *pickerWrapper {
	return &pickerWrapper{blockingCh: make(chan struct{})}
}

// updatePicker is called by UpdateBalancerState. It unblocks all blocked pick.
func (pw *pickerWrapper) updatePicker(p balancer.Picker) {
	pw.mu.Lock()
	if pw.done {
		pw.mu.Unlock()
		return
	}
	pw.picker = p
	// pw.blockingCh should never be nil.
	close(pw.blockingCh)
	pw.blockingCh = make(chan struct{})
	pw.mu.Unlock()
}

func doneChannelzWrapper(acw *acBalancerWrapper, done func(balancer.DoneInfo)) func(balancer.DoneInfo) {
	acw.mu.Lock()
	ac := acw.ac
	acw.mu.Unlock()
	ac.incrCallsStarted()
	return func(b balancer.DoneInfo) {
		if b.Err != nil && b.Err != io.EOF {
			ac.incrCallsFailed()
		} else {
			ac.incrCallsSucceeded()
		}
		if done != nil {
			done(b)
		}
	}
}

// pick returns the transport that will be used for the RPC.
// It may block in the following cases:
// - there's no picker
// - the current picker returns ErrNoSubConnAvailable
// - the current picker returns other errors and failfast is false.
// - the subConn returned by the current picker is not READY
// When one of these situations happens, pick blocks until the picker gets updated.
func (pw *pickerWrapper) pick(ctx context.Context, failfast bool, info balancer.PickInfo) (transport.ClientTransport, func(balancer.DoneInfo), error) {
	var ch chan struct{}

	var lastPickErr error
	for {
		pw.mu.Lock()
		if pw.done {
			pw.mu.Unlock()
			return nil, nil, ErrClientConnClosing
		}

		if pw.picker == nil {
			ch = pw.blockingCh
		}
		if ch == pw.blockingCh {
			// This could happen when either:
			// - pw.picker is nil (the previous if condition), or
			// - has called pick on the current picker.
			pw.mu.Unlock()
			select {
			case <-ctx.Done():
				var errStr string
				if lastPickErr != nil {
					errStr = "latest balancer error: " + lastPickErr.Error()
				} else {
					errStr = ctx.Err().Error()
				}
				switch ctx.Err() {
				case context.DeadlineExceeded:
					return nil, nil, status.Error(codes.DeadlineExceeded, errStr)
				case context.Canceled:
					return nil, nil, status.Error(codes.Canceled, errStr)
				}
			case <-ch:
			}
			continue
		}

		ch = pw.blockingCh
		p := pw.picker
		pw.mu.Unlock()

		pickResult, err := p.Pick(info)

		if err != nil {
			if err == balancer.ErrNoSubConnAvailable {
				continue
			}
			if _, ok := status.FromError(err); ok {
				// Status error: end the RPC unconditionally with this status.
				return nil, nil, err
			}
			// For all other errors, wait for ready RPCs should block and other
			// RPCs should fail with unavailable.
			if !failfast {
				lastPickErr = err
				continue
			}
			return nil, nil, status.Error(codes.Unavailable, err.Error())
		}

		acw, ok := pickResult.SubConn.(*acBalancerWrapper)
		if !ok {
			logger.Errorf("subconn returned from pick is type %T, not *acBalancerWrapper", pickResult.SubConn)
			continue
		}
		if t := acw.getAddrConn().getReadyTransport(); t != nil {
			if channelz.IsOn() {
				return t, doneChannelzWrapper(acw, pickResult.Done), nil
			}
			return t, pickResult.Done, nil
		}
		if pickResult.Done != nil {
			// Calling done with nil error, no bytes sent and no bytes received.
			// DoneInfo with default value works.
			pickResult.Done(balancer.DoneInfo{})
		}
<<<<<<< HEAD
		grpclog.Infof("blockingPicker: the picked transport is not ready, loop back to repick")
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
		grpclog.Infof("blockingPicker: the picked transport is not ready, loop back to repick")
=======
		logger.Infof("blockingPicker: the picked transport is not ready, loop back to repick")
>>>>>>> 4d7e5ad26 (update vendored files)
		// If ok == false, ac.state is not READY.
		// A valid picker always returns READY subConn. This means the state of ac
		// just changed, and picker will be updated shortly.
		// continue back to the beginning of the for loop to repick.
	}
}

func (pw *pickerWrapper) close() {
	pw.mu.Lock()
	defer pw.mu.Unlock()
	if pw.done {
		return
	}
	pw.done = true
	close(pw.blockingCh)
}

func (pw *pickerWrapper) enterIdleMode() {
	pw.mu.Lock()
	defer pw.mu.Unlock()
	if pw.done {
		return
	}
	pw.idle = true
}

func (pw *pickerWrapper) exitIdleMode() {
	pw.mu.Lock()
	defer pw.mu.Unlock()
	if pw.done {
		return
	}
	pw.blockingCh = make(chan struct{})
	pw.idle = false
}

// dropError is a wrapper error that indicates the LB policy wishes to drop the
// RPC and not retry it.
type dropError struct {
	error
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	"fmt"
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	"fmt"
=======
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
||||||| parent of c5487e6d6 (NE-2142: UPSTREAM: 5739: Bump k8s and controller-runtime modules)
=======
	"fmt"
>>>>>>> c5487e6d6 (NE-2142: UPSTREAM: 5739: Bump k8s and controller-runtime modules)
	"io"
	"sync/atomic"

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/internal/channelz"
	istatus "google.golang.org/grpc/internal/status"
	"google.golang.org/grpc/internal/transport"
	"google.golang.org/grpc/status"
)

// pickerGeneration stores a picker and a channel used to signal that a picker
// newer than this one is available.
type pickerGeneration struct {
	// picker is the picker produced by the LB policy.  May be nil if a picker
	// has never been produced.
	picker balancer.Picker
	// blockingCh is closed when the picker has been invalidated because there
	// is a new one available.
	blockingCh chan struct{}
}

// pickerWrapper is a wrapper of balancer.Picker. It blocks on certain pick
// actions and unblock when there's a picker update.
type pickerWrapper struct {
	// If pickerGen holds a nil pointer, the pickerWrapper is closed.
	pickerGen atomic.Pointer[pickerGeneration]
}

func newPickerWrapper() *pickerWrapper {
	pw := &pickerWrapper{}
	pw.pickerGen.Store(&pickerGeneration{
		blockingCh: make(chan struct{}),
	})
	return pw
}

// updatePicker is called by UpdateState calls from the LB policy. It
// unblocks all blocked pick.
func (pw *pickerWrapper) updatePicker(p balancer.Picker) {
	old := pw.pickerGen.Swap(&pickerGeneration{
		picker:     p,
		blockingCh: make(chan struct{}),
	})
	close(old.blockingCh)
}

// doneChannelzWrapper performs the following:
//   - increments the calls started channelz counter
//   - wraps the done function in the passed in result to increment the calls
//     failed or calls succeeded channelz counter before invoking the actual
//     done function.
func doneChannelzWrapper(acbw *acBalancerWrapper, result *balancer.PickResult) {
	ac := acbw.ac
	ac.incrCallsStarted()
	done := result.Done
	result.Done = func(b balancer.DoneInfo) {
		if b.Err != nil && b.Err != io.EOF {
			ac.incrCallsFailed()
		} else {
			ac.incrCallsSucceeded()
		}
		if done != nil {
			done(b)
		}
	}
}

type pick struct {
	transport transport.ClientTransport // the selected transport
	result    balancer.PickResult       // the contents of the pick from the LB policy
	blocked   bool                      // set if a picker call queued for a new picker
}

// pick returns the transport that will be used for the RPC.
// It may block in the following cases:
// - there's no picker
// - the current picker returns ErrNoSubConnAvailable
// - the current picker returns other errors and failfast is false.
// - the subConn returned by the current picker is not READY
// When one of these situations happens, pick blocks until the picker gets updated.
func (pw *pickerWrapper) pick(ctx context.Context, failfast bool, info balancer.PickInfo) (pick, error) {
	var ch chan struct{}

	var lastPickErr error
	pickBlocked := false

	for {
		pg := pw.pickerGen.Load()
		if pg == nil {
			return pick{}, ErrClientConnClosing
		}
		if pg.picker == nil {
			ch = pg.blockingCh
		}
		if ch == pg.blockingCh {
			// This could happen when either:
			// - pw.picker is nil (the previous if condition), or
			// - we have already called pick on the current picker.
			select {
			case <-ctx.Done():
				var errStr string
				if lastPickErr != nil {
					errStr = "latest balancer error: " + lastPickErr.Error()
				} else {
					errStr = fmt.Sprintf("%v while waiting for connections to become ready", ctx.Err())
				}
				switch ctx.Err() {
				case context.DeadlineExceeded:
					return pick{}, status.Error(codes.DeadlineExceeded, errStr)
				case context.Canceled:
					return pick{}, status.Error(codes.Canceled, errStr)
				}
			case <-ch:
			}
			continue
		}

		// If the channel is set, it means that the pick call had to wait for a
		// new picker at some point. Either it's the first iteration and this
		// function received the first picker, or a picker errored with
		// ErrNoSubConnAvailable or errored with failfast set to false, which
		// will trigger a continue to the next iteration. In the first case this
		// conditional will hit if this call had to block (the channel is set).
		// In the second case, the only way it will get to this conditional is
		// if there is a new picker.
		if ch != nil {
			pickBlocked = true
		}

		ch = pg.blockingCh
		p := pg.picker

		pickResult, err := p.Pick(info)
		if err != nil {
			if err == balancer.ErrNoSubConnAvailable {
				continue
			}
			if st, ok := status.FromError(err); ok {
				// Status error: end the RPC unconditionally with this status.
				// First restrict the code to the list allowed by gRFC A54.
				if istatus.IsRestrictedControlPlaneCode(st) {
					err = status.Errorf(codes.Internal, "received picker error with illegal status: %v", err)
				}
				return pick{}, dropError{error: err}
			}
			// For all other errors, wait for ready RPCs should block and other
			// RPCs should fail with unavailable.
			if !failfast {
				lastPickErr = err
				continue
			}
			return pick{}, status.Error(codes.Unavailable, err.Error())
		}

		acbw, ok := pickResult.SubConn.(*acBalancerWrapper)
		if !ok {
			logger.Errorf("subconn returned from pick is type %T, not *acBalancerWrapper", pickResult.SubConn)
			continue
		}
		if t := acbw.ac.getReadyTransport(); t != nil {
			if channelz.IsOn() {
				doneChannelzWrapper(acbw, &pickResult)
			}
			return pick{transport: t, result: pickResult, blocked: pickBlocked}, nil
		}
		if pickResult.Done != nil {
			// Calling done with nil error, no bytes sent and no bytes received.
			// DoneInfo with default value works.
			pickResult.Done(balancer.DoneInfo{})
		}
		if logger.V(2) {
			logger.Infof("blockingPicker: the picked transport is not ready, loop back to repick")
		}
		// If ok == false, ac.state is not READY.
		// A valid picker always returns READY subConn. This means the state of ac
		// just changed, and picker will be updated shortly.
		// continue back to the beginning of the for loop to repick.
	}
}

func (pw *pickerWrapper) close() {
<<<<<<< HEAD
	pw.mu.Lock()
	defer pw.mu.Unlock()
	if pw.done {
		return
	}
	pw.done = true
	close(pw.blockingCh)
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of c5487e6d6 (NE-2142: UPSTREAM: 5739: Bump k8s and controller-runtime modules)
	pw.mu.Lock()
	defer pw.mu.Unlock()
	if pw.done {
		return
	}
	pw.done = true
	close(pw.blockingCh)
=======
	old := pw.pickerGen.Swap(nil)
	close(old.blockingCh)
>>>>>>> c5487e6d6 (NE-2142: UPSTREAM: 5739: Bump k8s and controller-runtime modules)
}

// reset clears the pickerWrapper and prepares it for being used again when idle
// mode is exited.
func (pw *pickerWrapper) reset() {
	old := pw.pickerGen.Swap(&pickerGeneration{blockingCh: make(chan struct{})})
	close(old.blockingCh)
}

// dropError is a wrapper error that indicates the LB policy wishes to drop the
// RPC and not retry it.
type dropError struct {
	error
}
