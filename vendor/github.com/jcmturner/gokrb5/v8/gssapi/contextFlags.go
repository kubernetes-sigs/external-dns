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
// DEPRECATED - do not use
type ContextFlags asn1.BitString

// NewContextFlags creates a new ContextFlags instance
// DEPRECATED - do not use
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
type ContextFlags asn1.BitString

// NewContextFlags creates a new ContextFlags instance.
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
func NewContextFlags() ContextFlags {
	var c ContextFlags
	c.BitLength = 32
	c.Bytes = make([]byte, 4)
	return c
}
