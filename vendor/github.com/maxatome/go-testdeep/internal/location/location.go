<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
// Copyright (c) 2018-2021, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package location

import (
	"fmt"
	"runtime"
	"strings"
)

// Location records a place in a source file.
type Location struct {
	File      string // File name
	Func      string // Function name
	Line      int    // Line number inside file
	Inside    string // Inside is used when Location is inside something else
	BehindCmp bool   // BehindCmp is true when operator is behind a Cmp* function
}

// GetLocationer is the interface that wraps the basic GetLocation method.
type GetLocationer interface {
	GetLocation() Location
}

// New returns a new [Location]. callDepth is the number of
// stack frames to ascend to get the calling function (Func field),
// added to 1 to get the File & Line fields.
//
// If the location can not be determined, ok is false and location is
// not valid.
func New(callDepth int) (loc Location, ok bool) {
	_, loc.File, loc.Line, ok = runtime.Caller(callDepth + 1)
	if !ok {
		return
	}

	if index := strings.LastIndexAny(loc.File, `/\`); index >= 0 {
		loc.File = loc.File[index+1:]
	}

	pc, _, _, ok := runtime.Caller(callDepth)
	if !ok {
		return
	}

	loc.Func = runtime.FuncForPC(pc).Name()
	return
}

// IsInitialized returns true if l is initialized
// (e.g. [NewLocation] called without an error), false otherwise.
func (l Location) IsInitialized() bool {
	return l.File != ""
}

// Implements [fmt.Stringer].
func (l Location) String() string {
	return fmt.Sprintf("%s %sat %s:%d", l.Func, l.Inside, l.File, l.Line)
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// Copyright (c) 2018, Maxime Soulé
||||||| parent of 5ce8c7613 (update vendored files)
// Copyright (c) 2018, Maxime Soulé
=======
// Copyright (c) 2018-2021, Maxime Soulé
>>>>>>> 5ce8c7613 (update vendored files)
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package location

import (
	"fmt"
	"runtime"
	"strings"
)

// Location records a place in a source file.
type Location struct {
	File      string // File name
	Func      string // Function name
	Line      int    // Line number inside file
	Inside    string // Inside is used when Location is inside something else
	BehindCmp bool   // BehindCmp is true when operator is behind a Cmp* function
}

// GetLocationer is the interface that wraps the basic GetLocation method.
type GetLocationer interface {
	GetLocation() Location
}

// New returns a new Location. "callDepth" is the number of
// stack frames to ascend to get the calling function (Func field),
// added to 1 to get the File & Line fields.
//
// If the location can not be determined, false is returned and
// the Location is not valid.
func New(callDepth int) (loc Location, ok bool) {
	_, loc.File, loc.Line, ok = runtime.Caller(callDepth + 1)
	if !ok {
		return
	}

	if index := strings.LastIndexAny(loc.File, `/\`); index >= 0 {
		loc.File = loc.File[index+1:]
	}

	pc, _, _, ok := runtime.Caller(callDepth)
	if !ok {
		return
	}

	loc.Func = runtime.FuncForPC(pc).Name()
	return
}

// IsInitialized returns true if the Location is initialized
// (e.g. NewLocation() called without an error), false otherwise.
func (l Location) IsInitialized() bool {
	return l.File != ""
}

// Implements fmt.Stringer.
func (l Location) String() string {
<<<<<<< HEAD
	return fmt.Sprintf("%s at %s:%d", l.Func, l.File, l.Line)
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
	return fmt.Sprintf("%s at %s:%d", l.Func, l.File, l.Line)
=======
	return fmt.Sprintf("%s %sat %s:%d", l.Func, l.Inside, l.File, l.Line)
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// Copyright (c) 2018, Maxime Soulé
||||||| parent of 6b7ce455e (update vendored files)
// Copyright (c) 2018, Maxime Soulé
=======
// Copyright (c) 2018-2021, Maxime Soulé
>>>>>>> 6b7ce455e (update vendored files)
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package location

import (
	"fmt"
	"runtime"
	"strings"
)

// Location records a place in a source file.
type Location struct {
	File      string // File name
	Func      string // Function name
	Line      int    // Line number inside file
	Inside    string // Inside is used when Location is inside something else
	BehindCmp bool   // BehindCmp is true when operator is behind a Cmp* function
}

// GetLocationer is the interface that wraps the basic GetLocation method.
type GetLocationer interface {
	GetLocation() Location
}

// New returns a new Location. "callDepth" is the number of
// stack frames to ascend to get the calling function (Func field),
// added to 1 to get the File & Line fields.
//
// If the location can not be determined, false is returned and
// the Location is not valid.
func New(callDepth int) (loc Location, ok bool) {
	_, loc.File, loc.Line, ok = runtime.Caller(callDepth + 1)
	if !ok {
		return
	}

	if index := strings.LastIndexAny(loc.File, `/\`); index >= 0 {
		loc.File = loc.File[index+1:]
	}

	pc, _, _, ok := runtime.Caller(callDepth)
	if !ok {
		return
	}

	loc.Func = runtime.FuncForPC(pc).Name()
	return
}

// IsInitialized returns true if the Location is initialized
// (e.g. NewLocation() called without an error), false otherwise.
func (l Location) IsInitialized() bool {
	return l.File != ""
}

// Implements fmt.Stringer.
func (l Location) String() string {
<<<<<<< HEAD
	return fmt.Sprintf("%s at %s:%d", l.Func, l.File, l.Line)
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
	return fmt.Sprintf("%s at %s:%d", l.Func, l.File, l.Line)
=======
	return fmt.Sprintf("%s %sat %s:%d", l.Func, l.Inside, l.File, l.Line)
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// Copyright (c) 2018, Maxime Soulé
||||||| parent of 4d7e5ad26 (update vendored files)
// Copyright (c) 2018, Maxime Soulé
=======
// Copyright (c) 2018-2021, Maxime Soulé
>>>>>>> 4d7e5ad26 (update vendored files)
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package location

import (
	"fmt"
	"runtime"
	"strings"
)

// Location records a place in a source file.
type Location struct {
	File      string // File name
	Func      string // Function name
	Line      int    // Line number inside file
	Inside    string // Inside is used when Location is inside something else
	BehindCmp bool   // BehindCmp is true when operator is behind a Cmp* function
}

// GetLocationer is the interface that wraps the basic GetLocation method.
type GetLocationer interface {
	GetLocation() Location
}

// New returns a new Location. "callDepth" is the number of
// stack frames to ascend to get the calling function (Func field),
// added to 1 to get the File & Line fields.
//
// If the location can not be determined, false is returned and
// the Location is not valid.
func New(callDepth int) (loc Location, ok bool) {
	_, loc.File, loc.Line, ok = runtime.Caller(callDepth + 1)
	if !ok {
		return
	}

	if index := strings.LastIndexAny(loc.File, `/\`); index >= 0 {
		loc.File = loc.File[index+1:]
	}

	pc, _, _, ok := runtime.Caller(callDepth)
	if !ok {
		return
	}

	loc.Func = runtime.FuncForPC(pc).Name()
	return
}

// IsInitialized returns true if the Location is initialized
// (e.g. NewLocation() called without an error), false otherwise.
func (l Location) IsInitialized() bool {
	return l.File != ""
}

// Implements fmt.Stringer.
func (l Location) String() string {
<<<<<<< HEAD
	return fmt.Sprintf("%s at %s:%d", l.Func, l.File, l.Line)
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
	return fmt.Sprintf("%s at %s:%d", l.Func, l.File, l.Line)
=======
	return fmt.Sprintf("%s %sat %s:%d", l.Func, l.Inside, l.File, l.Line)
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package location

import (
	"fmt"
	"runtime"
	"strings"
)

// Location record a place in a source file.
type Location struct {
	File      string // File name
	Func      string // Function name
	Line      int    // Line number inside file
	BehindCmp bool   // BehindCmp is true when operator is behind a Cmp* function
}

// GetLocationer is the interface that wraps the basic GetLocation method.
type GetLocationer interface {
	GetLocation() Location
}

// New returns a new Location. "callDepth" is the number of
// stack frames to ascend to get the calling function (Func field),
// added to 1 to get the File & Line fields.
//
// If the location can not be determined, false is returned and
// the Location is not valid.
func New(callDepth int) (loc Location, ok bool) {
	_, loc.File, loc.Line, ok = runtime.Caller(callDepth + 1)
	if !ok {
		return
	}

	if index := strings.LastIndexAny(loc.File, `/\`); index >= 0 {
		loc.File = loc.File[index+1:]
	}

	pc, _, _, ok := runtime.Caller(callDepth)
	if !ok {
		return
	}

	loc.Func = runtime.FuncForPC(pc).Name()
	return
}

// IsInitialized returns true if the Location is initialized
// (e.g. NewLocation() called without an error), false otherwise.
func (l Location) IsInitialized() bool {
	return l.File != ""
}

// Implements fmt.Stringer.
func (l Location) String() string {
	return fmt.Sprintf("%s at %s:%d", l.Func, l.File, l.Line)
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
}
