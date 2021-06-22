//+build go1.9

package reflect2

import (
	"unsafe"
)

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
//go:linkname resolveTypeOff reflect.resolveTypeOff
func resolveTypeOff(rtype unsafe.Pointer, off int32) unsafe.Pointer

||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
=======
//go:linkname resolveTypeOff reflect.resolveTypeOff
func resolveTypeOff(rtype unsafe.Pointer, off int32) unsafe.Pointer

>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
//go:linkname makemap reflect.makemap
func makemap(rtype unsafe.Pointer, cap int) (m unsafe.Pointer)

func makeMapWithSize(rtype unsafe.Pointer, cap int) unsafe.Pointer {
	return makemap(rtype, cap)
}
