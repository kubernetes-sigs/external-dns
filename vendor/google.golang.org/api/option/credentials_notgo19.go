// Copyright 2018 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
<<<<<<< HEAD
//go:build !go1.9
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
=======
//go:build !go1.9
>>>>>>> 5ce8c7613 (update vendored files)
// +build !go1.9

package option

import (
	"golang.org/x/oauth2/google"
	"google.golang.org/api/internal"
)

type withCreds google.DefaultCredentials

func (w *withCreds) Apply(o *internal.DialSettings) {
	o.Credentials = (*google.DefaultCredentials)(w)
}

func WithCredentials(creds *google.DefaultCredentials) ClientOption {
	return (*withCreds)(creds)
}
