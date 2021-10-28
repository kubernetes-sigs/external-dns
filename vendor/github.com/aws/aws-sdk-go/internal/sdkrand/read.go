<<<<<<< HEAD
<<<<<<< HEAD
//go:build go1.6
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
=======
//go:build go1.6
>>>>>>> 5ce8c7613 (update vendored files)
// +build go1.6

package sdkrand

import "math/rand"

// Read provides the stub for math.Rand.Read method support for go version's
// 1.6 and greater.
func Read(r *rand.Rand, p []byte) (int, error) {
	return r.Read(p)
}
