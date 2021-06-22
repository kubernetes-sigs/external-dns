// Copyright 2020 Datawire.  All rights reserved
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

///////////////////////////////////////////////////////////////////////////
// Important: Run "make update-yaml" to regenerate code after modifying
// this file.
///////////////////////////////////////////////////////////////////////////

// We use a bunch of magic comments called "+markers".  Mostly they
// annotate a type, or a field within a struct.  Just below here, we
// do the "global" package-level markers.
//
// The type markers of interest are:
//
//  - "+kubebuilder:object:generate=bool" whether to generate
//    `DeepCopy` and `DeepCopyInto` methods for this type; but we
//    don't actually set this on types, since we can set it to true
//    for all types at the package-level.
//
//  - "+kubebuilder:object:root=bool" whether to *also* generate a
//    `DeepCopyObject` method.  It upsets me that controller-gen
//    doesn't infer this based on the presence of metav1.TypeMeta
//    inside of the type.
//
//  - "+kubebuilder:subresource:status" whether to add "status" as a
//    subresource for that type.  It upsets me that controller-gen
//    doesn't infer this based on the presence of a `status` field
//    inside of the type.
//
// The field markers of interest are:
//
//  - The "+kubebuilder:validation:*" markers control the OpenAPI v3
//    validation schema that is generated for this field.  ":Optional"
//    or ":Required" may be applied at the package-level in order to
//    set the default for all fields.  Most of the others can also be
//    set at the type level.
//
// Package-level markers:
//
// The group name to use for the CRDs in the generated YAML:
// +groupName=getambassador.io
//
// By default, mark all types in this package to have DeepCopy methods
// generated (so we don't need to specify this for every type):
// +kubebuilder:object:generate=true
//
// By default, mark all fields as optional (so we don't need to
// specify this for every optional field, since most fields are
// optional; and also because controller-gen's "required-by-default"
// mode is broken and always makes everything optional, even if it's
// explicitly marked as required):
// +kubebuilder:validation:Optional

// Package v2 contains API Schema definitions for the getambassador.io v2 API group
package v2

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	// GroupVersion is group version used to register these objects
	GroupVersion = schema.GroupVersion{Group: "getambassador.io", Version: "v2"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: GroupVersion}

	// AddToScheme adds the types in this group-version to the given scheme.
	AddToScheme = SchemeBuilder.AddToScheme
)
