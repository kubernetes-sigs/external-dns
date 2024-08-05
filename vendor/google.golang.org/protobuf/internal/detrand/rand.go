// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package detrand provides deterministically random functionality.
//
// The pseudo-randomness of these functions is seeded by the program binary
// itself and guarantees that the output does not change within a program,
// while ensuring that the output is unstable across different builds.
package detrand

import (
	"encoding/binary"
	"hash/fnv"
	"os"
)

// Disable disables detrand such that all functions returns the zero value.
// This function is not concurrent-safe and must be called during program init.
func Disable() {
	randSeed = 0
}

// Bool returns a deterministically random boolean.
func Bool() bool {
	return randSeed%2 == 1
}

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
// Intn returns a deterministically random integer between 0 and n-1, inclusive.
func Intn(n int) int {
	if n <= 0 {
		panic("must be positive")
	}
	return int(randSeed % uint64(n))
}

||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
=======
// Intn returns a deterministically random integer between 0 and n-1, inclusive.
func Intn(n int) int {
	if n <= 0 {
		panic("must be positive")
	}
	return int(randSeed % uint64(n))
}

>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
=======
// Intn returns a deterministically random integer between 0 and n-1, inclusive.
func Intn(n int) int {
	if n <= 0 {
		panic("must be positive")
	}
	return int(randSeed % uint64(n))
}

>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
=======
// Intn returns a deterministically random integer between 0 and n-1, inclusive.
func Intn(n int) int {
	if n <= 0 {
		panic("must be positive")
	}
	return int(randSeed % uint64(n))
}

>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
// Intn returns a deterministically random integer between 0 and n-1, inclusive.
func Intn(n int) int {
	if n <= 0 {
		panic("must be positive")
	}
	return int(randSeed % uint64(n))
}

>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
// randSeed is a best-effort at an approximate hash of the Go binary.
var randSeed = binaryHash()

func binaryHash() uint64 {
	// Open the Go binary.
	s, err := os.Executable()
	if err != nil {
		return 0
	}
	f, err := os.Open(s)
	if err != nil {
		return 0
	}
	defer f.Close()

	// Hash the size and several samples of the Go binary.
	const numSamples = 8
	var buf [64]byte
	h := fnv.New64()
	fi, err := f.Stat()
	if err != nil {
		return 0
	}
	binary.LittleEndian.PutUint64(buf[:8], uint64(fi.Size()))
	h.Write(buf[:8])
	for i := int64(0); i < numSamples; i++ {
		if _, err := f.ReadAt(buf[:], i*fi.Size()/numSamples); err != nil {
			return 0
		}
		h.Write(buf[:])
	}
	return h.Sum64()
}
