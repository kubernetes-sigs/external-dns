<<<<<<< HEAD
<<<<<<< HEAD
//go:build go1.9
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
=======
//go:build go1.9
>>>>>>> 5ce8c7613 (update vendored files)
// +build go1.9

package aws

import "context"

// Context is an alias of the Go stdlib's context.Context interface.
// It can be used within the SDK's API operation "WithContext" methods.
//
// See https://golang.org/pkg/context on how to use contexts.
type Context = context.Context
