// Copyright (c) 2019, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package util

import (
	"errors"
	"unicode"
)

<<<<<<< HEAD
// ErrTagEmpty is the error returned by [CheckTag] for an empty tag.
var ErrTagEmpty = errors.New("A tag cannot be empty")

// ErrTagInvalid is the error returned by [CheckTag] for an invalid tag.
var ErrTagInvalid = errors.New("Invalid tag, should match (Letter|_)(Letter|_|Number)*")

// CheckTag checks that tag is a valid tag (see operator [Tag]) or not.
//
// [Tag]: https://go-testdeep.zetta.rocks/operators/tag/
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
var (
	// ErrTagEmpty is the error returned by CheckTag for an empty tag.
	ErrTagEmpty = errors.New("A tag cannot be empty")
	// ErrTagInvalid is the error returned by CheckTag for an invalid tag.
	ErrTagInvalid = errors.New("Invalid tag, should match (Letter|_)(Letter|_|Number)*")
)

// CheckTag checks that tag is a valid tag (see operator Tag) or not.
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
func CheckTag(tag string) error {
	if tag == "" {
		return ErrTagEmpty
	}

	for i, r := range tag {
		if !(unicode.IsLetter(r) || r == '_' || (i > 0 && unicode.IsNumber(r))) {
			return ErrTagInvalid
		}
	}

	return nil
}
