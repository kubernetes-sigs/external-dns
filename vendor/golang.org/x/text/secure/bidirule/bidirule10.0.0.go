// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
//go:build go1.10
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
// +build go1.10

package bidirule

func (t *Transformer) isFinal() bool {
	return t.state == ruleLTRFinal || t.state == ruleRTLFinal || t.state == ruleInitial
}
