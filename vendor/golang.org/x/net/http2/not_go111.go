// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
//go:build !go1.11
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
=======
//go:build !go1.11
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
// +build !go1.11

package http2

import (
	"net/http/httptrace"
	"net/textproto"
)

func traceHasWroteHeaderField(trace *httptrace.ClientTrace) bool { return false }

func traceWroteHeaderField(trace *httptrace.ClientTrace, k, v string) {}

func traceGot1xxResponseFunc(trace *httptrace.ClientTrace) func(int, textproto.MIMEHeader) error {
	return nil
}
