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
<<<<<<< HEAD
<<<<<<< HEAD
// are passed as is to [fmt.Sprintf], else they are passed to [fmt.Sprint].
func BuildTestName(args ...any) string {
	if len(args) == 0 {
		return ""
	}

	var b bytes.Buffer
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
	if ok && len(args) > 1 && strings.ContainsRune(str, '%') {
<<<<<<< HEAD
<<<<<<< HEAD
		fmt.Fprintf(w, str, args[1:]...) //nolint: errcheck
	} else {
		// create a new slice to fool govet and avoid "call has possible
		// formatting directive" errors
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
		fmt.Fprint(w, args[:]...) // nolint: errcheck,gocritic
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
		fmt.Fprint(w, args[:]...) // nolint: errcheck
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
		fmt.Fprint(w, args[:]...) // nolint: errcheck
=======
		fmt.Fprint(w, args[:]...) // nolint: errcheck,gocritic
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
		fmt.Fprint(w, args[:]...) // nolint: errcheck
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
		fmt.Fprint(w, args[:]...) // nolint: errcheck
=======
		fmt.Fprint(w, args[:]...) //nolint: errcheck,gocritic
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
		fmt.Fprintf(w, str, args[1:]...) // nolint: errcheck
||||||| parent of 4d7e5ad26 (update vendored files)
		fmt.Fprintf(w, str, args[1:]...) // nolint: errcheck
=======
		fmt.Fprintf(w, str, args[1:]...) //nolint: errcheck
>>>>>>> 4d7e5ad26 (update vendored files)
	} else {
		// create a new slice to fool govet and avoid "call has possible
		// formatting directive" errors
<<<<<<< HEAD
		fmt.Fprint(w, args[:]...) // nolint: errcheck
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
		fmt.Fprint(w, args[:]...) // nolint: errcheck
=======
		fmt.Fprint(w, args[:]...) //nolint: errcheck,gocritic
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// are passed as is to fmt.Sprintf, else they are passed to fmt.Sprint.
func BuildTestName(args ...interface{}) string {
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
// are passed as is to fmt.Sprintf, else they are passed to fmt.Sprint.
func BuildTestName(args ...interface{}) string {
=======
// are passed as is to [fmt.Fprintf], else they are passed to [fmt.Fprint].
func BuildTestName(args ...any) string {
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
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
<<<<<<< HEAD
	if ok && len(args) > 1 && strings.ContainsRune(str, '%') {
		fmt.Fprintf(w, str, args[1:]...) // nolint: errcheck
	} else {
		// create a new slice to fool govet and avoid "call has possible
		// formatting directive" errors
		fmt.Fprint(w, args[:]...) // nolint: errcheck
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	if ok && len(args) > 1 && strings.ContainsRune(str, '%') {
		fmt.Fprintf(w, str, args[1:]...) // nolint: errcheck
	} else {
		// create a new slice to fool govet and avoid "call has possible
		// formatting directive" errors
		fmt.Fprint(w, args[:]...) // nolint: errcheck
=======
	if ok && len(args) > 1 {
		if pos := strings.IndexRune(str, '%'); pos >= 0 && pos < len(str)-1 {
			fmt.Fprintf(w, str, args[1:]...) //nolint: errcheck
			return
		}
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	}

	// create a new slice to fool govet and avoid "call has possible
	// formatting directive" errors
	fmt.Fprint(w, args[:]...) //nolint: errcheck,gocritic
}
