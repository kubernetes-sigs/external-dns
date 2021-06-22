<<<<<<< HEAD
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
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
// +build go1.9

package credentials

import "context"

// Context is an alias of the Go stdlib's context.Context interface.
// It can be used within the SDK's API operation "WithContext" methods.
//
// This type, aws.Context, and context.Context are equivalent.
//
// See https://golang.org/pkg/context on how to use contexts.
type Context = context.Context
