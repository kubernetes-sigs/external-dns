<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
//go:build !go1.6
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
=======
//go:build !go1.6
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
=======
//go:build !go1.6
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
=======
//go:build !go1.6
>>>>>>> 4d7e5ad26 (update vendored files)
// +build !go1.6

package sdkrand

import "math/rand"

// Read backfills Go 1.6's math.Rand.Reader for Go 1.5
func Read(r *rand.Rand, p []byte) (n int, err error) {
	// Copy of Go standard libraries math package's read function not added to
	// standard library until Go 1.6.
	var pos int8
	var val int64
	for n = 0; n < len(p); n++ {
		if pos == 0 {
			val = r.Int63()
			pos = 7
		}
		p[n] = byte(val)
		val >>= 8
		pos--
	}

	return n, err
}
