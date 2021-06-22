// Copyright © 2020 VMware
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
package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	// GroupName is the group name for the Contour API
	GroupName = "projectcontour.io"
)

// SchemeGroupVersion is a compatibility name for the GroupVersion.
// New code should use GroupVersion.
var SchemeGroupVersion = GroupVersion

var HTTPProxyGVR = GroupVersion.WithResource("httpproxies")
var TLSCertificateDelegationGVR = GroupVersion.WithResource("tlscertificatedelegations")

// Resource gets an Contour GroupResource for a specified resource
func Resource(resource string) schema.GroupResource {
	return GroupVersion.WithResource(resource).GroupResource()
}

// AddKnownTypes is exported for backwards compatibility with third
// parties who depend on this symbol, but all new code should use
// AddToScheme.
func AddKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(
		GroupVersion,
		&HTTPProxy{},
		&HTTPProxyList{},
		&TLSCertificateDelegation{},
		&TLSCertificateDelegationList{},
	)
	metav1.AddToGroupVersion(scheme, GroupVersion)
	return nil
}

// The following declarations are kubebuilder-compatible and will be expected
// by third parties who import the Contour API types.

var (
	// GroupVersion is group version used to register these objects
	GroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = runtime.NewSchemeBuilder(AddKnownTypes)

	// AddToScheme adds the types in this group-version to the given scheme.
	AddToScheme = SchemeBuilder.AddToScheme
)
