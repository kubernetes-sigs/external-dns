<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
=======
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
package packages

import (
	"cmp"
	"fmt"
	"iter"
	"os"
	"slices"
)

// Visit visits all the packages in the import graph whose roots are
// pkgs, calling the optional pre function the first time each package
// is encountered (preorder), and the optional post function after a
// package's dependencies have been visited (postorder).
// The boolean result of pre(pkg) determines whether
// the imports of package pkg are visited.
//
// Example:
//
//	pkgs, err := Load(...)
//	if err != nil { ... }
//	Visit(pkgs, nil, func(pkg *Package) {
//		log.Println(pkg)
//	})
//
// In most cases, it is more convenient to use [Postorder]:
//
//	for pkg := range Postorder(pkgs) {
//		log.Println(pkg)
//	}
func Visit(pkgs []*Package, pre func(*Package) bool, post func(*Package)) {
	seen := make(map[*Package]bool)
	var visit func(*Package)
	visit = func(pkg *Package) {
		if !seen[pkg] {
			seen[pkg] = true

			if pre == nil || pre(pkg) {
				for _, imp := range sorted(pkg.Imports) { // for determinism
					visit(imp)
				}
			}

			if post != nil {
				post(pkg)
			}
		}
	}
	for _, pkg := range pkgs {
		visit(pkg)
	}
}

// PrintErrors prints to os.Stderr the accumulated errors of all
// packages in the import graph rooted at pkgs, dependencies first.
// PrintErrors returns the number of errors printed.
func PrintErrors(pkgs []*Package) int {
	var n int
	errModules := make(map[*Module]bool)
	for pkg := range Postorder(pkgs) {
		for _, err := range pkg.Errors {
			fmt.Fprintln(os.Stderr, err)
			n++
		}

		// Print pkg.Module.Error once if present.
		mod := pkg.Module
		if mod != nil && mod.Error != nil && !errModules[mod] {
			errModules[mod] = true
			fmt.Fprintln(os.Stderr, mod.Error.Err)
			n++
		}
	}
	return n
}

// Postorder returns an iterator over the packages in
// the import graph whose roots are pkg.
// Packages are enumerated in dependencies-first order.
func Postorder(pkgs []*Package) iter.Seq[*Package] {
	return func(yield func(*Package) bool) {
		seen := make(map[*Package]bool)
		var visit func(*Package) bool
		visit = func(pkg *Package) bool {
			if !seen[pkg] {
				seen[pkg] = true
				for _, imp := range sorted(pkg.Imports) { // for determinism
					if !visit(imp) {
						return false
					}
				}
				if !yield(pkg) {
					return false
				}
			}
			return true
		}
		for _, pkg := range pkgs {
			if !visit(pkg) {
				break
			}
		}
	}
}

// -- copied from golang.org.x/tools/gopls/internal/util/moremaps --

// sorted returns an iterator over the entries of m in key order.
func sorted[M ~map[K]V, K cmp.Ordered, V any](m M) iter.Seq2[K, V] {
	// TODO(adonovan): use maps.Sorted if proposal #68598 is accepted.
	return func(yield func(K, V) bool) {
		keys := keySlice(m)
		slices.Sort(keys)
		for _, k := range keys {
			if !yield(k, m[k]) {
				break
			}
		}
	}
}

// KeySlice returns the keys of the map M, like slices.Collect(maps.Keys(m)).
func keySlice[M ~map[K]V, K comparable, V any](m M) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}
