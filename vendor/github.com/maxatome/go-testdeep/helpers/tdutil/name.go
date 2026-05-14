// Copyright (c) 2019, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package tdutil

import (
	"fmt"
	"io"
	"strings"
)

// BuildTestName builds a string from given args.
//
// If optional first args is a string containing at least one %, args
// are passed as is to [fmt.Fprintf], else they are passed to [fmt.Fprint].
func BuildTestName(args ...any) string {
	if len(args) == 0 {
		return ""
	}

	var b strings.Builder
	FbuildTestName(&b, args...)
	return b.String()
}

// FbuildTestName builds a string from given args.
//
// If optional first args is a string containing at least one %, args
// are passed as is to [fmt.Fprintf], else they are passed to [fmt.Fprint].
func FbuildTestName(w io.Writer, args ...any) {
	if len(args) == 0 {
		return
	}

	str, ok := args[0].(string)
	if ok && len(args) > 1 {
		if pos := strings.IndexRune(str, '%'); pos >= 0 && pos < len(str)-1 {
			fmt.Fprintf(w, str, args[1:]...) //nolint: errcheck
			return
		}
	}

	// create a new slice to fool govet and avoid "call has possible
	// formatting directive" errors
	fmt.Fprint(w, args[:]...) //nolint: errcheck,gocritic
}
