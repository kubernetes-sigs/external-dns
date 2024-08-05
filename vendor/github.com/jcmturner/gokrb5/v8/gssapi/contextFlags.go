package gssapi

import "github.com/jcmturner/gofork/encoding/asn1"

// GSS-API context flags assigned numbers.
const (
	ContextFlagDeleg    = 1
	ContextFlagMutual   = 2
	ContextFlagReplay   = 4
	ContextFlagSequence = 8
	ContextFlagConf     = 16
	ContextFlagInteg    = 32
	ContextFlagAnon     = 64
)

// ContextFlags flags for GSSAPI
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
// DEPRECATED - do not use
type ContextFlags asn1.BitString

// NewContextFlags creates a new ContextFlags instance
// DEPRECATED - do not use
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 4d7e5ad26 (update vendored files)
=======
// DEPRECATED - do not use
>>>>>>> 4d7e5ad26 (update vendored files)
type ContextFlags asn1.BitString

<<<<<<< HEAD
// NewContextFlags creates a new ContextFlags instance.
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
// NewContextFlags creates a new ContextFlags instance.
=======
// NewContextFlags creates a new ContextFlags instance
// DEPRECATED - do not use
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
// DEPRECATED - do not use
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
type ContextFlags asn1.BitString

<<<<<<< HEAD
// NewContextFlags creates a new ContextFlags instance.
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
// NewContextFlags creates a new ContextFlags instance.
=======
// NewContextFlags creates a new ContextFlags instance
// DEPRECATED - do not use
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
func NewContextFlags() ContextFlags {
	var c ContextFlags
	c.BitLength = 32
	c.Bytes = make([]byte, 4)
	return c
}
