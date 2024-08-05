// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
//go:build go1.18
// +build go1.18

||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
package gcimporter

import "go/types"

const iexportVersion = iexportVersionGenerics

// additionalPredeclared returns additional predeclared types in go.1.18.
func additionalPredeclared() []types.Type {
	return []types.Type{
		// comparable
		types.Universe.Lookup("comparable").Type(),

		// any
		types.Universe.Lookup("any").Type(),
	}
}

// See cmd/compile/internal/types.SplitVargenSuffix.
func splitVargenSuffix(name string) (base, suffix string) {
	i := len(name)
	for i > 0 && name[i-1] >= '0' && name[i-1] <= '9' {
		i--
	}
	const dot = "·"
	if i >= len(dot) && name[i-len(dot):i] == dot {
		i -= len(dot)
		return name[:i], name[i:]
	}
	return name, ""
}
