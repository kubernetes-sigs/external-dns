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

// TLSContextSpec defines the desired state of TLSContext
type TLSContextSpec struct {
	AmbassadorID AmbassadorID `json:"ambassador_id,omitempty"`

	Hosts           []string `json:"hosts,omitempty"`
	Secret          string   `json:"secret,omitempty"`
	CertChainFile   string   `json:"cert_chain_file,omitempty"`
	PrivateKeyFile  string   `json:"private_key_file,omitempty"`
	CASecret        string   `json:"ca_secret,omitempty"`
	CACertChainFile string   `json:"cacert_chain_file,omitempty"`
	ALPNProtocols   string   `json:"alpn_protocols,omitempty"`
	CertRequired    bool     `json:"cert_required,omitempty"`
	// +kubebuilder:validation:Enum={"v1.0", "v1.1", "v1.2", "v1.3"}
	MinTLSVersion string `json:"min_tls_version,omitempty"`
	// +kubebuilder:validation:Enum={"v1.0", "v1.1", "v1.2", "v1.3"}
	MaxTLSVersion         string   `json:"max_tls_version,omitempty"`
	CipherSuites          []string `json:"cipher_suites,omitempty"`
	ECDHCurves            []string `json:"ecdh_curves,omitempty"`
	SecretNamespacing     bool     `json:"secret_namespacing,omitempty"`
	RedirectCleartextFrom int      `json:"redirect_cleartext_from,omitempty"`
	SNI                   string   `json:"sni,omitempty"`
}

// TLSContext is the Schema for the tlscontexts API
//
// +kubebuilder:object:root=true
type TLSContext struct {
	metav1.TypeMeta   `json:""`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec TLSContextSpec `json:"spec,omitempty"`
}

// TLSContextList contains a list of TLSContexts.
//
// +kubebuilder:object:root=true
type TLSContextList struct {
	metav1.TypeMeta `json:""`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TLSContext `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TLSContext{}, &TLSContextList{})
}
