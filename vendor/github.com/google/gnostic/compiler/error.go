<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
// Copyright 2017 Google LLC. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package compiler

import "fmt"

// Error represents compiler errors and their location in the document.
type Error struct {
	Context *Context
	Message string
}

// NewError creates an Error.
func NewError(context *Context, message string) *Error {
	return &Error{Context: context, Message: message}
}

func (err *Error) locationDescription() string {
	if err.Context.Node != nil {
		return fmt.Sprintf("[%d,%d] %s", err.Context.Node.Line, err.Context.Node.Column, err.Context.Description())
	}
	return err.Context.Description()
}

// Error returns the string value of an Error.
func (err *Error) Error() string {
	if err.Context == nil {
		return err.Message
	}
	return err.locationDescription() + " " + err.Message
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// Copyright 2017 Google Inc. All Rights Reserved.
||||||| parent of 5ce8c7613 (update vendored files)
// Copyright 2017 Google Inc. All Rights Reserved.
=======
// Copyright 2017 Google LLC. All Rights Reserved.
>>>>>>> 5ce8c7613 (update vendored files)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package compiler

import "fmt"

// Error represents compiler errors and their location in the document.
type Error struct {
	Context *Context
	Message string
}

// NewError creates an Error.
func NewError(context *Context, message string) *Error {
	return &Error{Context: context, Message: message}
}

func (err *Error) locationDescription() string {
	if err.Context.Node != nil {
		return fmt.Sprintf("[%d,%d] %s", err.Context.Node.Line, err.Context.Node.Column, err.Context.Description())
	}
	return err.Context.Description()
}

// Error returns the string value of an Error.
func (err *Error) Error() string {
	if err.Context == nil {
		return err.Message
	}
<<<<<<< HEAD
	return "ERROR " + err.Context.Description() + " " + err.Message
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
	return "ERROR " + err.Context.Description() + " " + err.Message
=======
	return err.locationDescription() + " " + err.Message
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// Copyright 2017 Google Inc. All Rights Reserved.
||||||| parent of 6b7ce455e (update vendored files)
// Copyright 2017 Google Inc. All Rights Reserved.
=======
// Copyright 2017 Google LLC. All Rights Reserved.
>>>>>>> 6b7ce455e (update vendored files)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package compiler

import "fmt"

// Error represents compiler errors and their location in the document.
type Error struct {
	Context *Context
	Message string
}

// NewError creates an Error.
func NewError(context *Context, message string) *Error {
	return &Error{Context: context, Message: message}
}

func (err *Error) locationDescription() string {
	if err.Context.Node != nil {
		return fmt.Sprintf("[%d,%d] %s", err.Context.Node.Line, err.Context.Node.Column, err.Context.Description())
	}
	return err.Context.Description()
}

// Error returns the string value of an Error.
func (err *Error) Error() string {
	if err.Context == nil {
		return err.Message
	}
<<<<<<< HEAD
	return "ERROR " + err.Context.Description() + " " + err.Message
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
	return "ERROR " + err.Context.Description() + " " + err.Message
=======
	return err.locationDescription() + " " + err.Message
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// Copyright 2017 Google Inc. All Rights Reserved.
||||||| parent of 4d7e5ad26 (update vendored files)
// Copyright 2017 Google Inc. All Rights Reserved.
=======
// Copyright 2017 Google LLC. All Rights Reserved.
>>>>>>> 4d7e5ad26 (update vendored files)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package compiler

import "fmt"

// Error represents compiler errors and their location in the document.
type Error struct {
	Context *Context
	Message string
}

// NewError creates an Error.
func NewError(context *Context, message string) *Error {
	return &Error{Context: context, Message: message}
}

func (err *Error) locationDescription() string {
	if err.Context.Node != nil {
		return fmt.Sprintf("[%d,%d] %s", err.Context.Node.Line, err.Context.Node.Column, err.Context.Description())
	}
	return err.Context.Description()
}

// Error returns the string value of an Error.
func (err *Error) Error() string {
	if err.Context == nil {
		return err.Message
	}
<<<<<<< HEAD
	return "ERROR " + err.Context.Description() + " " + err.Message
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
	return "ERROR " + err.Context.Description() + " " + err.Message
=======
	return err.locationDescription() + " " + err.Message
>>>>>>> 4d7e5ad26 (update vendored files)
}

// ErrorGroup is a container for groups of Error values.
type ErrorGroup struct {
	Errors []error
}

// NewErrorGroupOrNil returns a new ErrorGroup for a slice of errors or nil if the slice is empty.
func NewErrorGroupOrNil(errors []error) error {
	if len(errors) == 0 {
		return nil
	} else if len(errors) == 1 {
		return errors[0]
	} else {
		return &ErrorGroup{Errors: errors}
	}
}

func (group *ErrorGroup) Error() string {
	result := ""
	for i, err := range group.Errors {
		if i > 0 {
			result += "\n"
		}
		result += err.Error()
	}
	return result
}
