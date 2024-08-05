// Copyright (c) 2023, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

//go:build go1.18
// +build go1.18

package td

// Anchor is a generic shortcut to [T.Anchor].
func Anchor[X any](t *T, operator TestDeep) X {
	var model X
	return t.Anchor(operator, model).(X)
}

// A is a generic shortcut to [T.A].
func A[X any](t *T, operator TestDeep) X {
	var model X
	return t.A(operator, model).(X)
}
