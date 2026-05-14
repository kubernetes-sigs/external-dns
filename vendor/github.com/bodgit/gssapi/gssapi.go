/*
Package gssapi implements a simplified wrapper around the
github.com/jcmturner/gokrb5 package.
*/
package gssapi

import "github.com/jcmturner/gokrb5/v8/gssapi"

const (
	supportedFlags = gssapi.ContextFlagMutual | gssapi.ContextFlagReplay |
		gssapi.ContextFlagSequence | gssapi.ContextFlagConf |
		gssapi.ContextFlagInteg
)
