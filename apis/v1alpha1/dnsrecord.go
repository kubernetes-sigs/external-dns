/*
Copyright 2025 The Kubernetes Authors.

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
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"sigs.k8s.io/external-dns/endpoint"
)

const (
	RecordOwnerLabel    string = "externaldns.k8s.io/owner"
	RecordNameLabel     string = "externaldns.k8s.io/record-name"
	RecordTypeLabel     string = "externaldns.k8s.io/record-type"
	RecordKeyLabel      string = "externaldns.k8s.io/key"
	RecordResourceLabel string = "externaldns.k8s.io/resource"
)

// DNSRecordSpec defines the desired state of DNSEndpoint
// +kubebuilder:object:generate=true
type DNSRecordSpec struct {
	Endpoint endpoint.Endpoint `json:"endpoints,omitempty"`
}

// DNSRecordStatus defines the observed state of DNSRecord
// +kubebuilder:object:generate=true
type DNSRecordStatus struct {
	// The generation observed by the external-dns controller.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DNSRecord is used to get all records managed by external-dns.
// It can be used as a registry with the status subresource.
// +k8s:openapi-gen=true
// +groupName=externaldns.k8s.io
// +kubebuilder:resource:path=dnsrecords
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:metadata:annotations="api-approved.kubernetes.io=https://github.com/kubernetes-sigs/external-dns/pull/5372"
// +versionName=v1alpha1

type DNSRecord struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DNSRecordSpec   `json:"spec,omitempty"`
	Status DNSRecordStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// DNSEndpointList is a list of DNSEndpoint objects
type DNSRecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DNSRecord `json:"items"`
}

func (dr DNSRecord) IsEndpoint(e *endpoint.Endpoint) bool {
	spec := dr.Spec.Endpoint

	return spec.DNSName == strings.ToLower(e.DNSName) &&
		spec.RecordType == e.RecordType &&
		spec.SetIdentifier == e.SetIdentifier
}

func (dr DNSRecord) EndpointLabels() endpoint.Labels {
	labels := endpoint.Labels{}

	labels[endpoint.OwnerLabelKey] = dr.Labels[RecordOwnerLabel]
	labels[endpoint.ResourceLabelKey] = dr.Labels[RecordResourceLabel]
	return labels
}
