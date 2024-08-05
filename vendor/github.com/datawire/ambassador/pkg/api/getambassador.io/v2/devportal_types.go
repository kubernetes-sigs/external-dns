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

package v2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DevPortalContentSpec defines the content origin
type DevPortalContentSpec struct {
	URL    string `json:"url,omitempty"`
	Branch string `json:"branch,omitempty"`
	Dir    string `json:"dir,omitempty"`
}

// DevPortalSelectorSpec is the selector for filtering mappings
// that are used in this DevPortal. They can be filtered by things like
// namespace, labels, etc...
type DevPortalSelectorSpec struct {
	// MatchNamespaces is a list of namespaces that will be included in
	// this DevPortal.
	MatchNamespaces []string `json:"matchNamespaces,omitempty"`

	// MatchLabels specifies the list of labels that must be present
	// in Mappings for being present in this DevPortal.
	MatchLabels map[string]string `json:"matchLabels,omitempty"`
}

// DevPortalDocsSpec is a static documentation definition:
// instead of using a Selector for finding documentation for services,
// users can provide a static list of <service>:<URL> tuples. These services
// will be shown in the Dev Portal with the documentation obtained from
// this URL.
type DevPortalDocsSpec struct {
	// Service is the service being documented
	Service string `json:"service,omitempty"`

	// URL is the URL used for obtaining docs
	URL string `json:"url,omitempty"`
}

// DevPortalSpec defines the desired state of DevPortal
type DevPortalSpec struct {
	AmbassadorID AmbassadorID `json:"ambassador_id,omitempty"`

	// Default must be true when this is the default DevPortal
	Default *bool `json:"default,omitempty"`

	// Content specifies where the content shown in the DevPortal come from
	Content *DevPortalContentSpec `json:"content,omitempty"`

	// Docs is a static docs definition
	Docs []*DevPortalDocsSpec `json:"docs,omitempty"`

	// Selector is used for choosing what is shown in the DevPortal
	Selector *DevPortalSelectorSpec `json:"selector,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DevPortal is the Schema for the DevPortals API
//
// DevPortal resources specify the `what` and `how` is shown in a DevPortal:
//
// * `what` is in a DevPortal can be controlled with
//   - a `selector`, that can be used for filtering `Mappings`.
//   - a `docs` listing of (services, url)
// * `how` is a pointer to some `contents` (a checkout of a Git repository
//   with go-templates/markdown/css).
//
// Multiple `DevPortal`s can exist in the cluster, and the Dev Portal server
// will show them at different endpoints. A `DevPortal` resource with a special
// name, `ambassador`, will be used for configuring the default Dev Portal
// (served at `/docs/` by default).
//
// +kubebuilder:object:root=true
// +kubebuilder:resource:path=devportals,scope=Namespaced
type DevPortal struct {
	metav1.TypeMeta   `json:""`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec DevPortalSpec `json:"spec,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DevPortalList contains a list of DevPortals.
//
// +kubebuilder:object:root=true
type DevPortalList struct {
	metav1.TypeMeta `json:""`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DevPortal `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DevPortal{}, &DevPortalList{})
}
