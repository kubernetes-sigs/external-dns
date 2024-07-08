package crds

import (
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/external-dns/endpoint"
)

const (
	RegistryOwnerLabel      string = "externaldns.k8s.io/owner"
	RegistryRecordNameLabel string = "externaldns.k8s.io/record-name"
	RegistryRecordTypeLabel string = "externaldns.k8s.io/record-type"
	RegistryIdentifierLabel string = "externaldns.k8s.io/identifier"
	RegistryResourceLabel   string = "externaldns.k8s.io/resource"
)

// DNSEntrySpec defines the desired state of DNSEndpoint
// +kubebuilder:object:generate=true
type DNSEntrySpec struct {
	Endpoint endpoint.Endpoint `json:"endpoints,omitempty"`
}

// DNSEntryStatus defines the observed state of DNSENtry
// +kubebuilder:object:generate=true
type DNSEntryStatus struct {
	// The generation observed by the external-dns controller.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DNSEntry is a contract that a user-specified CRD must implement to be used as a source for external-dns.
// The user-specified CRD should also have the status sub-resource.
// +k8s:openapi-gen=true
// +groupName=externaldns.k8s.io
// +kubebuilder:resource:path=dnsentries
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +versionName=v1alpha1

type DNSEntry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DNSEntrySpec   `json:"spec,omitempty"`
	Status DNSEntryStatus `json:"status,omitempty"`
}

func (de DNSEntry) IsEndpoint(e *endpoint.Endpoint) bool {
	spec := de.Spec.Endpoint

	return spec.DNSName == strings.ToLower(e.DNSName) &&
		spec.RecordType == e.RecordType &&
		spec.SetIdentifier == e.SetIdentifier
}

func (de DNSEntry) EndpointLabels() endpoint.Labels {
	labels := endpoint.Labels{}

	labels[endpoint.OwnerLabelKey] = de.ObjectMeta.Labels[RegistryOwnerLabel]
	labels[endpoint.ResourceLabelKey] = de.ObjectMeta.Labels[RegistryResourceLabel]
	return labels
}

// +kubebuilder:object:root=true
// DNSEndpointList is a list of DNSEndpoint objects
type DNSEntryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DNSEntry `json:"items"`
}
