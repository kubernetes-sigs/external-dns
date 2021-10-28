// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
<<<<<<< HEAD
//go:build !go1.10
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
=======
//go:build !go1.10
>>>>>>> 5ce8c7613 (update vendored files)
// +build !go1.10

package bidirule

func (t *Transformer) isFinal() bool {
	if !t.isRTL() {
		return true
	}
	return t.state == ruleLTRFinal || t.state == ruleRTLFinal || t.state == ruleInitial
}
