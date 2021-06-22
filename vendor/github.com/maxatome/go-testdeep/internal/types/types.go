// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package types

import (
	"strconv"
)

// TestDeepStringer is a TestDeep specific interface for objects which
// know how to stringify themselves.
type TestDeepStringer interface {
	_TestDeep()
	String() string
}

// TestDeepStamp is a useful type providing the _TestDeep() method
// needed to implement TestDeepStringer interface.
type TestDeepStamp struct{}

func (t TestDeepStamp) _TestDeep() {}

// RawString implements TestDeepStringer interface.
type RawString string

func (s RawString) _TestDeep() {}

func (s RawString) String() string {
	return string(s)
}

// RawInt implements TestDeepStringer interface.
type RawInt int

func (i RawInt) _TestDeep() {}

func (i RawInt) String() string {
	return strconv.Itoa(int(i))
}
