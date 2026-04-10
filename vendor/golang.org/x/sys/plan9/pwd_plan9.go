// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
||||||| parent of 6b7ce455e (update vendored files)
=======
//go:build !go1.5
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4d7e5ad26 (update vendored files)
=======
//go:build !go1.5
>>>>>>> 4d7e5ad26 (update vendored files)
// +build !go1.5
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
//go:build !go1.5
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)

||||||| parent of 53ef3ded0 (UPSTREAM: 6362: OCPBUGS-79591: Bump deps to get google.golang.org/grpc v1.80.0)
//go:build !go1.5

=======
>>>>>>> 53ef3ded0 (UPSTREAM: 6362: OCPBUGS-79591: Bump deps to get google.golang.org/grpc v1.80.0)
package plan9

import "syscall"

func fixwd() {
	syscall.Fixwd()
}

func Getwd() (wd string, err error) {
	return syscall.Getwd()
}

func Chdir(path string) error {
	return syscall.Chdir(path)
}
