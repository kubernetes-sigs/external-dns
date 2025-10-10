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
	"fmt"
	"os"
	"sort"
)

// Visit visits all the packages in the import graph whose roots are
// pkgs, calling the optional pre function the first time each package
// is encountered (preorder), and the optional post function after a
// package's dependencies have been visited (postorder).
// The boolean result of pre(pkg) determines whether
// the imports of package pkg are visited.
func Visit(pkgs []*Package, pre func(*Package) bool, post func(*Package)) {
	seen := make(map[*Package]bool)
	var visit func(*Package)
	visit = func(pkg *Package) {
		if !seen[pkg] {
			seen[pkg] = true

			if pre == nil || pre(pkg) {
				paths := make([]string, 0, len(pkg.Imports))
				for path := range pkg.Imports {
					paths = append(paths, path)
				}
				sort.Strings(paths) // Imports is a map, this makes visit stable
				for _, path := range paths {
					visit(pkg.Imports[path])
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
	Visit(pkgs, nil, func(pkg *Package) {
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
	})
	return n
}
