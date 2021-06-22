// Copyright 2017, The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
// license that can be found in the LICENSE file.
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// license that can be found in the LICENSE.md file.
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
// license that can be found in the LICENSE.md file.
=======
// license that can be found in the LICENSE file.
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// license that can be found in the LICENSE.md file.
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
// license that can be found in the LICENSE.md file.
=======
// license that can be found in the LICENSE file.
>>>>>>> 6b7ce455e (update vendored files)

//go:build !cmp_debug
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// license that can be found in the LICENSE.md file.
||||||| parent of 4d7e5ad26 (update vendored files)
// license that can be found in the LICENSE.md file.
=======
// license that can be found in the LICENSE file.
>>>>>>> 4d7e5ad26 (update vendored files)

<<<<<<< HEAD
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
=======
//go:build !cmp_debug
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// license that can be found in the LICENSE.md file.

>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
// +build !cmp_debug

package diff

var debug debugger

type debugger struct{}

func (debugger) Begin(_, _ int, f EqualFunc, _, _ *EditScript) EqualFunc {
	return f
}
func (debugger) Update() {}
func (debugger) Finish() {}
