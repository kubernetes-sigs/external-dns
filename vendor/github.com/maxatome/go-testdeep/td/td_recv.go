// Copyright (c) 2022, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

import (
	"fmt"
	"reflect"
	"time"

	"github.com/maxatome/go-testdeep/internal/ctxerr"
	"github.com/maxatome/go-testdeep/internal/types"
)

// A RecvKind allows to match that nothing has been received on a
// channel or that a channel has been closed when using [Recv]
// operator.
type RecvKind = types.RecvKind

const (
	_           RecvKind = (iota & 1) == 0
	RecvNothing          // nothing read on channel
	RecvClosed           // channel closed
)

type tdRecv struct {
	tdSmugglerBase
	timeout time.Duration
}

var _ TestDeep = &tdRecv{}

// summary(Recv): checks the value read from a channel
// input(Recv): chan,ptr(ptr on chan)

// Recv is a smuggler operator. It reads from a channel or a pointer
// to a channel and compares the read value to expectedValue.
//
// expectedValue can be any value including a [TestDeep] operator. It
// can also be [RecvNothing] to test nothing can be read from the
// channel or [RecvClosed] to check the channel is closed.
//
// If timeout is passed it should be only one item. It means: try to
// read the channel during this duration to get a value before giving
// up. If timeout is missing or ≤ 0, it defaults to 0 meaning Recv
// does not wait for a value but gives up instantly if no value is
// available on the channel.
//
//	ch := make(chan int, 6)
//	td.Cmp(t, ch, td.Recv(td.RecvNothing)) // succeeds
//	td.Cmp(t, ch, td.Recv(42))             // fails, nothing to receive
//	// recv(DATA): values differ
//	//      got: nothing received on channel
//	// expected: 42
//
//	ch <- 42
//	td.Cmp(t, ch, td.Recv(td.RecvNothing)) // fails, 42 received instead
//	// recv(DATA): values differ
//	//      got: 42
//	// expected: nothing received on channel
//
//	td.Cmp(t, ch, td.Recv(42)) // fails, nothing to receive anymore
//	// recv(DATA): values differ
//	//      got: nothing received on channel
//	// expected: 42
//
//	ch <- 666
//	td.Cmp(t, ch, td.Recv(td.Between(600, 700))) // succeeds
//
//	close(ch)
//	td.Cmp(t, ch, td.Recv(td.RecvNothing)) // fails as channel is closed
//	// recv(DATA): values differ
//	//      got: channel is closed
//	// expected: nothing received on channel
//
//	td.Cmp(t, ch, td.Recv(td.RecvClosed)) // succeeds
//
// Note that for convenience Recv accepts pointer on channel:
//
//	ch := make(chan int, 6)
//	ch <- 42
//	td.Cmp(t, &ch, td.Recv(42)) // succeeds
//
// Each time Recv is called, it tries to consume one item from the
// channel, immediately or, if given, before timeout duration. To
// consume several items in a same [Cmp] call, one can use [All]
// operator as in:
//
//	ch := make(chan int, 6)
//	ch <- 1
//	ch <- 2
//	ch <- 3
//	close(ch)
//	td.Cmp(t, ch, td.All( // succeeds
//	  td.Recv(1),
//	  td.Recv(2),
//	  td.Recv(3),
//	  td.Recv(td.RecvClosed),
//	))
//
// To check nothing can be received during 100ms on channel ch (if
// something is received before, including a close, it fails):
//
//	td.Cmp(t, ch, td.Recv(td.RecvNothing, 100*time.Millisecond))
//
// note that in case of success, the above [Cmp] call always lasts 100ms.
//
// To check 42 can be received from channel ch during the next 100ms
// (if nothing is received during these 100ms or something different
// from 42, including a close, it fails):
//
//	td.Cmp(t, ch, td.Recv(42, 100*time.Millisecond))
//
// note that in case of success, the above [Cmp] call lasts less than 100ms.
//
// A nil channel is not handled specifically, so it “is never ready
// for communication” as specification says:
//
//	var ch chan int
//	td.Cmp(t, ch, td.Recv(td.RecvNothing)) // always succeeds
//	td.Cmp(t, ch, td.Recv(42))             // or any other value, always fails
//	td.Cmp(t, ch, td.Recv(td.RecvClosed))  // always fails
//
// so to check if a channel is not nil before reading from it, one can
// either do:
//
//	td.Cmp(t, ch, td.All(
//	  td.NotNil(),
//	  td.Recv(42),
//	))
//	// or
//	if td.Cmp(t, ch, td.NotNil()) {
//	  td.Cmp(t, ch, td.Recv(42))
//	}
//
// TypeBehind method returns the [reflect.Type] of expectedValue,
// except if expectedValue is a [TestDeep] operator. In this case, it
// delegates TypeBehind() to the operator.
//
// See also [Cap] and [Len].
func Recv(expectedValue any, timeout ...time.Duration) TestDeep {
	r := tdRecv{}
	r.tdSmugglerBase = newSmugglerBase(expectedValue, 0)

	if !r.isTestDeeper {
		r.expectedValue = reflect.ValueOf(expectedValue)
	}

	switch len(timeout) {
	case 0:
	case 1:
		r.timeout = timeout[0]
	default:
		r.err = ctxerr.OpTooManyParams(r.location.Func, "(EXPECTED[, TIMEOUT])")
	}
	return &r
}

func (r *tdRecv) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	if r.err != nil {
		return ctx.CollectError(r.err)
	}

	switch got.Kind() {
	case reflect.Ptr:
		gotElem := got.Elem()
		if !gotElem.IsValid() {
			if ctx.BooleanError {
				return ctxerr.BooleanError
			}
			return ctx.CollectError(ctxerr.NilPointer(got, "non-nil *chan"))
		}
		if gotElem.Kind() != reflect.Chan {
			break
		}
		got = gotElem
		fallthrough

	case reflect.Chan:
		cases := [2]reflect.SelectCase{
			{
				Dir:  reflect.SelectRecv,
				Chan: got,
			},
		}
		var timer *time.Timer
		if r.timeout > 0 {
			timer = time.NewTimer(r.timeout)
			cases[1] = reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(timer.C),
			}
		} else {
			cases[1] = reflect.SelectCase{
				Dir: reflect.SelectDefault,
			}
		}

		chosen, recv, recvOK := reflect.Select(cases[:])
		if chosen == 1 && timer != nil {
			// check quickly both timeout & expected case didn't occur
			// concurrently and timeout masked the expected case
			cases[1] = reflect.SelectCase{
				Dir: reflect.SelectDefault,
			}
			chosen, recv, recvOK = reflect.Select(cases[:])
		}
		if chosen == 0 {
			if !recvOK {
				recv = reflect.ValueOf(RecvClosed)
			}
			if timer != nil {
				timer.Stop()
			}
		} else {
			recv = reflect.ValueOf(RecvNothing)
		}

		return deepValueEqual(ctx.AddFunctionCall("recv"), recv, r.expectedValue)
	}

	if ctx.BooleanError {
		return ctxerr.BooleanError
	}
	return ctx.CollectError(ctxerr.BadKind(got, "chan OR *chan"))
}

func (r *tdRecv) HandleInvalid() bool {
	return true // Knows how to handle untyped nil values (aka invalid values)
}

func (r *tdRecv) String() string {
	if r.err != nil {
		return r.stringError()
	}
	if r.isTestDeeper {
		return "recv: " + r.expectedValue.Interface().(TestDeep).String()
	}
	return fmt.Sprintf("recv=%d", r.expectedValue.Int())
}

func (r *tdRecv) TypeBehind() reflect.Type {
	if r.err != nil {
		return nil
	}
	return r.internalTypeBehind()
}
