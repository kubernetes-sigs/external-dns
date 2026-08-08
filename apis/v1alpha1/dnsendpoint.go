/*
Copyright 2017 The Kubernetes Authors.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"sigs.k8s.io/external-dns/endpoint"
)

const (
	// AcceptedCondition reports whether every endpoint in spec passed source-level
	// validation. It is set before any provider call; ReadyCondition (shared with
	// DNSRecord, see dnsrecord.go) then reports whether the provider applied them.
	AcceptedCondition string = "Accepted"

	// InvalidReason marks a spec containing at least one endpoint external-dns
	// refused to plan, e.g. an SRV target without a trailing dot.
	InvalidReason string = "Invalid"

	// FilteredReason is a ReadyCondition reason: external-dns understood every
	// endpoint in spec but --domain-filter or the managed record types excluded
	// all of them, so none was ever offered to the provider.
	FilteredReason string = "Filtered"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DNSEndpoint is a contract that a user-specified CRD must implement to be used as a source for external-dns.
// The user-specified CRD should also have the status sub-resource.
// +k8s:openapi-gen=true
// +groupName=externaldns.k8s.io
// +kubebuilder:resource:path=dnsendpoints
// +kubebuilder:subresource:status
// +kubebuilder:metadata:annotations="api-approved.kubernetes.io=https://github.com/kubernetes-sigs/external-dns/pull/2007"
// +kubebuilder:printcolumn:name="Endpoints",type=integer,JSONPath=`.status.endpoints`
// +kubebuilder:printcolumn:name="Accepted",type=string,JSONPath=`.status.conditions[?(@.type=="Accepted")].status`
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].reason`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
// +versionName=v1alpha1
type DNSEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DNSEndpointSpec   `json:"spec,omitempty"`
	Status DNSEndpointStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DNSEndpointList is a list of DNSEndpoint objects
type DNSEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DNSEndpoint `json:"items"`
}

// DNSEndpointSpec defines the desired state of DNSEndpoint
type DNSEndpointSpec struct {
	// Endpoints is the list of DNS records this resource declares.
	// +optional
	// +kubebuilder:validation:MaxItems=1000
	Endpoints []*endpoint.Endpoint `json:"endpoints,omitempty"`
}

// DNSEndpointStatus defines the observed state of DNSEndpoint
type DNSEndpointStatus struct {
	// The generation observed by the external-dns controller.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// Endpoints is the number of endpoints from spec that external-dns took into
	// its plan on the last reconcile. Endpoints dropped by validation, or excluded
	// by --domain-filter or the managed record types, are not counted.
	// +optional
	Endpoints int32 `json:"endpoints"`

	// Conditions represent the latest available observations of the DNSEndpoint
	// state: Accepted (external-dns understood the spec) and Ready (the provider
	// applied it).
	// +optional
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
