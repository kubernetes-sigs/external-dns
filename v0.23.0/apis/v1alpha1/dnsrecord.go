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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"sigs.k8s.io/external-dns/endpoint"
)

const (
	// RecordOwnerLabel ties a DNSRecord to the external-dns instance that owns it.
	// The record identity (name, type, set identifier) is encoded in the object
	// name instead of labels, so it is not subject to the 63-character label-value
	// limit and records are looked up directly by name.
	RecordOwnerLabel string = "externaldns.k8s.io/owner"

	// ReadyCondition reports whether the endpoint is live in the DNS provider.
	// Its reason captures the lifecycle stage and is surfaced as the Status print
	// column: AcceptedReason while external-dns has taken the endpoint into its
	// plan but not yet programmed it (Ready=False), ProgrammedReason once the
	// provider has applied it (Ready=True), or FailedReason when the provider
	// rejected the batch it belonged to (Ready=False).
	ReadyCondition string = "Ready"

	// Reasons for the Ready condition. They double as the human-readable value of
	// the Status print column.
	AcceptedReason   string = "Accepted"
	ProgrammedReason string = "Programmed"
	FailedReason     string = "Failed"
)

// DNSRecordSpec defines the desired state of DNSRecord
// +kubebuilder:object:generate=true
type DNSRecordSpec struct {
	Endpoint endpoint.Endpoint `json:"endpoint"`
}

// MergeProviderLabels overwrites the record's endpoint labels with the labels
// the provider stored, preserving the external-dns owner and resource labels.
// Some providers (e.g. coredns) rewrite labels on apply, so the provider copy
// is authoritative for everything except ownership/resource identity.
func (r *DNSRecord) MergeProviderLabels(providerLabels endpoint.Labels, ownerID string) {
	resource := r.Spec.Endpoint.Labels[endpoint.ResourceLabelKey]
	r.Spec.Endpoint.Labels = providerLabels
	r.Spec.Endpoint.WithLabel(endpoint.OwnerLabelKey, ownerID)
	if resource != "" {
		r.Spec.Endpoint.WithLabel(endpoint.ResourceLabelKey, resource)
	}
}

// DNSRecordStatus defines the observed state of DNSRecord
// +kubebuilder:object:generate=true
type DNSRecordStatus struct {
	// Conditions represent the latest available observations of the DNSRecord state.
	// +optional
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DNSRecord is used to get all records managed by external-dns.
// It can be used as a registry with the status subresource.
// +k8s:openapi-gen=true
// +kubebuilder:resource:path=dnsrecords
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:metadata:annotations="api-approved.kubernetes.io=https://github.com/kubernetes-sigs/external-dns/pull/5372"
// +kubebuilder:printcolumn:name="DNS Name",type=string,JSONPath=`.spec.endpoint.dnsName`
// +kubebuilder:printcolumn:name="Type",type=string,JSONPath=`.spec.endpoint.recordType`
// +kubebuilder:printcolumn:name="Set ID",type=string,JSONPath=`.spec.endpoint.setIdentifier`
// +kubebuilder:printcolumn:name="Targets",type=string,JSONPath=`.spec.endpoint.targets`
// +kubebuilder:printcolumn:name="Status",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].reason`
// +versionName=v1alpha1

type DNSRecord struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DNSRecordSpec   `json:"spec,omitempty"`
	Status DNSRecordStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// DNSRecordList is a list of DNSRecord objects
type DNSRecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DNSRecord `json:"items"`
}
