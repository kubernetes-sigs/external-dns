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

type AuthServiceIncludeBody struct {
	// +kubebuilder:validation:Required
	MaxBytes int `json:"max_bytes,omitempty"`

	// +kubebuilder:validation:Required
	AllowPartial bool `json:"allow_partial,omitempty"`
}

// Why isn't this just an int??
type AuthServiceStatusOnError struct {
	Code int `json:"code,omitempty"`
}

// AuthServiceSpec defines the desired state of AuthService
type AuthServiceSpec struct {
	AmbassadorID AmbassadorID `json:"ambassador_id,omitempty"`

	AuthService string       `json:"auth_service,omitempty"`
	PathPrefix  string       `json:"path_prefix,omitempty"`
	TLS         BoolOrString `json:"tls,omitempty"`
	// +kubebuilder:validation:Enum={"http","grpc"}
	Proto                       string                    `json:"proto,omitempty"`
	TimeoutMs                   int                       `json:"timeout_ms,omitempty"`
	AllowedRequestHeaders       []string                  `json:"allowed_request_headers,omitempty"`
	AllowedAuthorizationHeaders []string                  `json:"allowed_authorization_headers,omitempty"`
	AddAuthHeaders              map[string]BoolOrString   `json:"add_auth_headers,omitempty"`
	AllowRequestBody            bool                      `json:"allow_request_body,omitempty"`
	AddLinkerdHeaders           bool                      `json:"add_linkerd_headers,omitempty"`
	FailureModeAllow            bool                      `json:"failure_mode_allow,omitempty"`
	IncludeBody                 *AuthServiceIncludeBody   `json:"include_body,omitempty"`
	StatusOnError               *AuthServiceStatusOnError `json:"status_on_error,omitempty"`
}

// AuthService is the Schema for the authservices API
//
// +kubebuilder:object:root=true
type AuthService struct {
	metav1.TypeMeta   `json:""`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec AuthServiceSpec `json:"spec,omitempty"`
}

// AuthServiceList contains a list of AuthServices.
//
// +kubebuilder:object:root=true
type AuthServiceList struct {
	metav1.TypeMeta `json:""`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AuthService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AuthService{}, &AuthServiceList{})
}
