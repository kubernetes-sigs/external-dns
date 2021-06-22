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

// KubernetesServiceResolver tells Ambassador to use Kubernetes Service
// resources to resolve services. It actually has no spec other than the
// AmbassadorID.
type KubernetesServiceResolverSpec struct {
	AmbassadorID AmbassadorID `json:"ambassador_id,omitempty"`
}

// KubernetesServiceResolver is the Schema for the kubernetesserviceresolver API
//
// +kubebuilder:object:root=true
type KubernetesServiceResolver struct {
	metav1.TypeMeta   `json:""`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec KubernetesServiceResolverSpec `json:"spec,omitempty"`
}

// KubernetesServiceResolverList contains a list of KubernetesServiceResolvers.
//
// +kubebuilder:object:root=true
type KubernetesServiceResolverList struct {
	metav1.TypeMeta `json:""`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KubernetesServiceResolver `json:"items"`
}

// KubernetesEndpointResolver tells Ambassador to use Kubernetes Endpoints
// resources to resolve services. It actually has no spec other than the
// AmbassadorID.
type KubernetesEndpointResolverSpec struct {
	AmbassadorID AmbassadorID `json:"ambassador_id,omitempty"`
}

// KubernetesEndpointResolver is the Schema for the kubernetesendpointresolver API
//
// +kubebuilder:object:root=true
type KubernetesEndpointResolver struct {
	metav1.TypeMeta   `json:""`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec KubernetesEndpointResolverSpec `json:"spec,omitempty"`
}

// KubernetesEndpointResolverList contains a list of KubernetesEndpointResolvers.
//
// +kubebuilder:object:root=true
type KubernetesEndpointResolverList struct {
	metav1.TypeMeta `json:""`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KubernetesEndpointResolver `json:"items"`
}

// ConsulResolver tells Ambassador to use Consul to resolve services. In addition
// to the AmbassadorID, it needs information about which Consul server and DC to
// use.
type ConsulResolverSpec struct {
	AmbassadorID AmbassadorID `json:"ambassador_id,omitempty"`

	Address    string `json:"address,omitempty"`
	Datacenter string `json:"datacenter,omitempty"`
}

// ConsulResolver is the Schema for the ConsulResolver API
//
// +kubebuilder:object:root=true
type ConsulResolver struct {
	metav1.TypeMeta   `json:""`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ConsulResolverSpec `json:"spec,omitempty"`
}

// ConsulResolverList contains a list of ConsulResolvers.
//
// +kubebuilder:object:root=true
type ConsulResolverList struct {
	metav1.TypeMeta `json:""`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConsulResolver `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KubernetesServiceResolver{}, &KubernetesServiceResolverList{})
	SchemeBuilder.Register(&KubernetesEndpointResolver{}, &KubernetesEndpointResolverList{})
	SchemeBuilder.Register(&ConsulResolver{}, &ConsulResolverList{})
}
