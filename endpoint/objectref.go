/*
Copyright 2026 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package endpoint

import (
	"k8s.io/apimachinery/pkg/types"
)

// ObjectRef identifies the Kubernetes object an Endpoint was derived from.
//
// It lives here rather than in pkg/events so that this package, which every
// webhook provider imports, does not depend on client-go. Keep its imports
// free of the Kubernetes client stack.
type ObjectRef struct {
	kind       string
	apiVersion string
	namespace  string
	name       string
	uid        types.UID
	source     string
}

// NewObjectRef constructs an ObjectRef from its components.
func NewObjectRef(kind, apiVersion, namespace, name string, uid types.UID, source string) *ObjectRef {
	return &ObjectRef{
		kind:       kind,
		apiVersion: apiVersion,
		namespace:  namespace,
		name:       name,
		uid:        uid,
		source:     source,
	}
}

// Key returns a stable string that uniquely identifies this object reference
// in the form "source/namespace/name".
func (r *ObjectRef) Key() string {
	return r.source + "/" + r.namespace + "/" + r.name
}

// Kind returns the Kubernetes kind of the referenced object (e.g. "Service", "Ingress").
func (r *ObjectRef) Kind() string {
	return r.kind
}

// APIVersion returns the API version of the referenced Kubernetes object.
func (r *ObjectRef) APIVersion() string {
	return r.apiVersion
}

// Namespace returns the namespace of the referenced Kubernetes object.
func (r *ObjectRef) Namespace() string {
	return r.namespace
}

// Name returns the name of the referenced Kubernetes object.
func (r *ObjectRef) Name() string {
	return r.name
}

// Source returns the source identifier of the ObjectRef (e.g. "ingress", "service").
func (r *ObjectRef) Source() string {
	return r.source
}

// UID returns the UID of the referenced Kubernetes object.
func (r *ObjectRef) UID() types.UID {
	return r.uid
}
