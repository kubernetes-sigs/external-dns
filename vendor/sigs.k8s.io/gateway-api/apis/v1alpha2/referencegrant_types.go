/*
Copyright 2021 The Kubernetes Authors.

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

package v1alpha2

<<<<<<< HEAD
import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:resource:categories=gateway-api,shortName=refgrant
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`

// ReferenceGrant identifies kinds of resources in other namespaces that are
// trusted to reference the specified kinds of resources in the same namespace
// as the policy.
//
// Each ReferenceGrant can be used to represent a unique trust relationship.
// Additional Reference Grants can be used to add to the set of trusted
// sources of inbound references for the namespace they are defined within.
//
// All cross-namespace references in Gateway API (with the exception of cross-namespace
// Gateway-route attachment) require a ReferenceGrant.
//
// Support: Core
//
type ReferenceGrant struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the desired state of ReferenceGrant.
	Spec ReferenceGrantSpec `json:"spec,omitempty"`

	// Note that `Status` sub-resource has been excluded at the
	// moment as it was difficult to work out the design.
	// `Status` sub-resource may be added in future.
}

// +kubebuilder:object:root=true
// ReferenceGrantList contains a list of ReferenceGrant.
type ReferenceGrantList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ReferenceGrant `json:"items"`
}

// ReferenceGrantSpec identifies a cross namespace relationship that is trusted
// for Gateway API.
type ReferenceGrantSpec struct {
	// From describes the trusted namespaces and kinds that can reference the
	// resources described in "To". Each entry in this list must be considered
	// to be an additional place that references can be valid from, or to put
	// this another way, entries must be combined using OR.
	//
	// Support: Core
	//
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:MaxItems=16
	From []ReferenceGrantFrom `json:"from"`

	// To describes the resources that may be referenced by the resources
	// described in "From". Each entry in this list must be considered to be an
	// additional place that references can be valid to, or to put this another
	// way, entries must be combined using OR.
	//
	// Support: Core
	//
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:MaxItems=16
	To []ReferenceGrantTo `json:"to"`
}

// ReferenceGrantFrom describes trusted namespaces and kinds.
type ReferenceGrantFrom struct {
	// Group is the group of the referent.
	// When empty, the Kubernetes core API group is inferred.
	//
	// Support: Core
	Group Group `json:"group"`

	// Kind is the kind of the referent. Although implementations may support
	// additional resources, the following types are part of the "Core"
	// support level for this field.
	//
	// When used to permit a SecretObjectReference:
	//
	// * Gateway
	//
	// When used to permit a BackendObjectReference:
	//
	// * HTTPRoute
	// * TCPRoute
	// * TLSRoute
	// * UDPRoute
	Kind Kind `json:"kind"`

	// Namespace is the namespace of the referent.
	//
	// Support: Core
	Namespace Namespace `json:"namespace"`
}

// ReferenceGrantTo describes what Kinds are allowed as targets of the
// references.
type ReferenceGrantTo struct {
	// Group is the group of the referent.
	// When empty, the Kubernetes core API group is inferred.
	//
	// Support: Core
	Group Group `json:"group"`

	// Kind is the kind of the referent. Although implementations may support
	// additional resources, the following types are part of the "Core"
	// support level for this field:
	//
	// * Secret when used to permit a SecretObjectReference
	// * Service when used to permit a BackendObjectReference
	Kind Kind `json:"kind"`

	// Name is the name of the referent. When unspecified, this policy
	// refers to all resources of the specified Group and Kind in the local
	// namespace.
	//
	// +optional
	Name *ObjectName `json:"name,omitempty"`
}
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"sigs.k8s.io/gateway-api/apis/v1beta1"
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:resource:categories=gateway-api,shortName=refgrant
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
// +kubebuilder:deprecatedversion:warning="The v1alpha2 version of ReferenceGrant has been deprecated and will be removed in a future release of the API. Please upgrade to v1beta1."

// ReferenceGrant identifies kinds of resources in other namespaces that are
// trusted to reference the specified kinds of resources in the same namespace
// as the policy.
//
// Each ReferenceGrant can be used to represent a unique trust relationship.
// Additional Reference Grants can be used to add to the set of trusted
// sources of inbound references for the namespace they are defined within.
//
// A ReferenceGrant is required for all cross-namespace references in Gateway API
// (with the exception of cross-namespace Route-Gateway attachment, which is
// governed by the AllowedRoutes configuration on the Gateway, and cross-namespace
// Service ParentRefs on a "consumer" mesh Route, which defines routing rules
// applicable only to workloads in the Route namespace). ReferenceGrants allowing
// a reference from a Route to a Service are only applicable to BackendRefs.
//
// ReferenceGrant is a form of runtime verification allowing users to assert
// which cross-namespace object references are permitted. Implementations that
// support ReferenceGrant MUST NOT permit cross-namespace references which have
// no grant, and MUST respond to the removal of a grant by revoking the access
// that the grant allowed.
type ReferenceGrant v1beta1.ReferenceGrant

// +kubebuilder:object:root=true
// ReferenceGrantList contains a list of ReferenceGrant.
type ReferenceGrantList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ReferenceGrant `json:"items"`
}

// ReferenceGrantSpec identifies a cross namespace relationship that is trusted
// for Gateway API.
// +k8s:deepcopy-gen=false
type ReferenceGrantSpec = v1beta1.ReferenceGrantSpec

// ReferenceGrantFrom describes trusted namespaces and kinds.
// +k8s:deepcopy-gen=false
type ReferenceGrantFrom = v1beta1.ReferenceGrantFrom

// ReferenceGrantTo describes what Kinds are allowed as targets of the
// references.
// +k8s:deepcopy-gen=false
type ReferenceGrantTo = v1beta1.ReferenceGrantTo
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
