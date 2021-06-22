// Copyright © 2019 VMware
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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TLSCertificateDelegationSpec defines the spec of the CRD
type TLSCertificateDelegationSpec struct {
	Delegations []CertificateDelegation `json:"delegations"`
}

// CertificateDelegation maps the authority to reference a secret
// in the current namespace to a set of namespaces.
type CertificateDelegation struct {

	// required, the name of a secret in the current namespace.
	SecretName string `json:"secretName"`

	// required, the namespaces the authority to reference the
	// the secret will be delegated to.
	// If TargetNamespaces is nil or empty, the CertificateDelegation'
	// is ignored. If the TargetNamespace list contains the character, "*"
	// the secret will be delegated to all namespaces.
	TargetNamespaces []string `json:"targetNamespaces"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TLSCertificateDelegation is an TLS Certificate Delegation CRD specificiation.
// See design/tls-certificate-delegation.md for details.
// +k8s:openapi-gen=true
// +kubebuilder:resource:scope=Namespaced,path=tlscertificatedelegations,singular=tlscertificatedelegation
type TLSCertificateDelegation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec TLSCertificateDelegationSpec `json:"spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TLSCertificateDelegationList is a list of TLSCertificateDelegations.
type TLSCertificateDelegationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []TLSCertificateDelegation `json:"items"`
}
