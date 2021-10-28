// Copyright (c) 2020, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

// +build !go1.14

package td

import (
	"runtime"
	"testing"
)

func cleanupTB(tb testing.TB, finalize func()) {
	runtime.SetFinalizer(tb, func(_ testing.TB) { finalize() })
}
