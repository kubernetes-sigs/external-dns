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

import (
	yaml "gopkg.in/yaml.v3"
)

// Context contains state of the compiler as it traverses a document.
type Context struct {
	Parent            *Context
	Name              string
	Node              *yaml.Node
	ExtensionHandlers *[]ExtensionHandler
}

// NewContextWithExtensions returns a new object representing the compiler state
func NewContextWithExtensions(name string, node *yaml.Node, parent *Context, extensionHandlers *[]ExtensionHandler) *Context {
	return &Context{Name: name, Node: node, Parent: parent, ExtensionHandlers: extensionHandlers}
}

// NewContext returns a new object representing the compiler state
func NewContext(name string, node *yaml.Node, parent *Context) *Context {
	if parent != nil {
		return &Context{Name: name, Node: node, Parent: parent, ExtensionHandlers: parent.ExtensionHandlers}
	}
	return &Context{Name: name, Parent: parent, ExtensionHandlers: nil}
}

// Description returns a text description of the compiler state
func (context *Context) Description() string {
	name := context.Name
	if context.Parent != nil {
		name = context.Parent.Description() + "." + name
	}
	return name
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

import (
	yaml "gopkg.in/yaml.v3"
)

// Context contains state of the compiler as it traverses a document.
type Context struct {
	Parent            *Context
	Name              string
	Node              *yaml.Node
	ExtensionHandlers *[]ExtensionHandler
}

// NewContextWithExtensions returns a new object representing the compiler state
func NewContextWithExtensions(name string, node *yaml.Node, parent *Context, extensionHandlers *[]ExtensionHandler) *Context {
	return &Context{Name: name, Node: node, Parent: parent, ExtensionHandlers: extensionHandlers}
}

// NewContext returns a new object representing the compiler state
func NewContext(name string, node *yaml.Node, parent *Context) *Context {
	if parent != nil {
		return &Context{Name: name, Node: node, Parent: parent, ExtensionHandlers: parent.ExtensionHandlers}
	}
	return &Context{Name: name, Parent: parent, ExtensionHandlers: nil}
}

// Description returns a text description of the compiler state
func (context *Context) Description() string {
	name := context.Name
	if context.Parent != nil {
		name = context.Parent.Description() + "." + name
	}
<<<<<<< HEAD
	return context.Name
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
	return context.Name
=======
	return name
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

import (
	yaml "gopkg.in/yaml.v3"
)

// Context contains state of the compiler as it traverses a document.
type Context struct {
	Parent            *Context
	Name              string
	Node              *yaml.Node
	ExtensionHandlers *[]ExtensionHandler
}

// NewContextWithExtensions returns a new object representing the compiler state
func NewContextWithExtensions(name string, node *yaml.Node, parent *Context, extensionHandlers *[]ExtensionHandler) *Context {
	return &Context{Name: name, Node: node, Parent: parent, ExtensionHandlers: extensionHandlers}
}

// NewContext returns a new object representing the compiler state
func NewContext(name string, node *yaml.Node, parent *Context) *Context {
	if parent != nil {
		return &Context{Name: name, Node: node, Parent: parent, ExtensionHandlers: parent.ExtensionHandlers}
	}
	return &Context{Name: name, Parent: parent, ExtensionHandlers: nil}
}

// Description returns a text description of the compiler state
func (context *Context) Description() string {
	name := context.Name
	if context.Parent != nil {
		name = context.Parent.Description() + "." + name
	}
<<<<<<< HEAD
	return context.Name
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
	return context.Name
=======
	return name
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// Copyright 2017 Google Inc. All Rights Reserved.
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

// Context contains state of the compiler as it traverses a document.
type Context struct {
	Parent            *Context
	Name              string
	ExtensionHandlers *[]ExtensionHandler
}

// NewContextWithExtensions returns a new object representing the compiler state
func NewContextWithExtensions(name string, parent *Context, extensionHandlers *[]ExtensionHandler) *Context {
	return &Context{Name: name, Parent: parent, ExtensionHandlers: extensionHandlers}
}

// NewContext returns a new object representing the compiler state
func NewContext(name string, parent *Context) *Context {
	if parent != nil {
		return &Context{Name: name, Parent: parent, ExtensionHandlers: parent.ExtensionHandlers}
	}
	return &Context{Name: name, Parent: parent, ExtensionHandlers: nil}
}

// Description returns a text description of the compiler state
func (context *Context) Description() string {
	if context.Parent != nil {
		return context.Parent.Description() + "." + context.Name
	}
	return context.Name
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
}
