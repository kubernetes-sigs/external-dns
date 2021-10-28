<<<<<<< HEAD
<<<<<<< HEAD
//go:build !go1.7
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
=======
//go:build !go1.7
>>>>>>> 5ce8c7613 (update vendored files)
// +build !go1.7

package sdkio

// Copy of Go 1.7 io package's Seeker constants.
const (
	SeekStart   = 0 // seek relative to the origin of the file
	SeekCurrent = 1 // seek relative to the current offset
	SeekEnd     = 2 // seek relative to the end
)
