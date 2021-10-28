<<<<<<< HEAD
<<<<<<< HEAD
//go:build gofuzz
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
=======
//go:build gofuzz
>>>>>>> 5ce8c7613 (update vendored files)
// +build gofuzz

package ini

import (
	"bytes"
)

func Fuzz(data []byte) int {
	b := bytes.NewReader(data)

	if _, err := Parse(b); err != nil {
		return 0
	}

	return 1
}
