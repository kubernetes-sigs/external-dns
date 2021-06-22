// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package util

// TernRune returns a if cond is true, b otherwise.
func TernRune(cond bool, a, b rune) rune {
	if cond {
		return a
	}
	return b
}

// TernStr returns a if cond is true, b otherwise.
func TernStr(cond bool, a, b string) string {
	if cond {
		return a
	}
	return b
}
