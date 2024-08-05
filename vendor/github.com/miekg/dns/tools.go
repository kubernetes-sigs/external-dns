<<<<<<< HEAD
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
//go:build tools
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
// +build tools

// We include our tool dependencies for `go generate` here to ensure they're
// properly tracked by the go tool. See the Go Wiki for the rationale behind this:
// https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module.

package dns

import _ "golang.org/x/tools/go/packages"
