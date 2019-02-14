/*
Copyright 2018 The Kubernetes Authors.

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

package crd

import (
	"github.com/kubernetes-incubator/external-dns/endpoint"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DNSZone is the internal represenation of the DNSZone CRD
type DNSZone struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec DNSZoneSpec `json:"spec,omitempty"`
}

// DNSZoneSpec represents the zone options that can be configured
type DNSZoneSpec struct {
	Name                    string             `json:"name"`
	ProviderSpecificOptions []ProviderSpecific `json:"providerSpecific,omitempty"`
}

//DNSRecord is the internal representation of the DNSRecord CRD
type DNSRecord struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec DNSRecordSpec `json:"spec,omitempty"`
}

// DNSRecordSpec represents the record options that can be configured
type DNSRecordSpec struct {
	Records []Record `json:"records,omitempty"`
}

// DNSZoneList is a list of DNSZone objects
type DNSZoneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []DNSZone `json:"items,omitempty"`
}

// DNSRecordList is a list of DNSRecord objects
type DNSRecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []DNSRecord `json:"items,omitempty"`
}

// ProviderSpecific holds key/value pairs of options to be passed to different DNS providers
type ProviderSpecific struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

// Record is an internal representation of a DNS record
type Record struct {
	Name                    string                              `json:"name,omitempty"`
	TTL                     endpoint.TTL                        `json:"ttl,omitempty"`
	Type                    string                              `json:"type,omitempty"`
	Targets                 []string                            `json:"targets,omitempty"`
	ProviderSpecificOptions []endpoint.ProviderSpecificProperty `json:"providerSpecific,omitempty"`
}
