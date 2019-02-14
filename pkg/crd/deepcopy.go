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
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto copies all properties of this object into another object of the
// same type that is provided as a pointer.
func (in *DNSZone) DeepCopyInto(out *DNSZone) {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	out.Spec.Name = in.Spec.Name

	if in.Spec.ProviderSpecificOptions != nil {
		out.Spec.ProviderSpecificOptions = make([]ProviderSpecific, len(in.Spec.ProviderSpecificOptions))
		for i := range in.Spec.ProviderSpecificOptions {
			out.Spec.ProviderSpecificOptions[i].Name = in.Spec.ProviderSpecificOptions[i].Name
			out.Spec.ProviderSpecificOptions[i].Value = in.Spec.ProviderSpecificOptions[i].Value
		}
	}

}

// DeepCopyObject returns a generically typed copy of an object
func (in *DNSZone) DeepCopyObject() runtime.Object {
	out := DNSZone{}
	in.DeepCopyInto(&out)

	return &out
}

// DeepCopyObject returns a generically typed copy of an object
func (in *DNSZoneList) DeepCopyObject() runtime.Object {
	out := DNSZoneList{}
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta

	if in.Items != nil {
		out.Items = make([]DNSZone, len(in.Items))
		for i := range in.Items {
			in.Items[i].DeepCopyInto(&out.Items[i])
		}
	}

	return &out
}

// DeepCopyInto copies all properties of this object into another object of the
// same type that is provided as a pointer.
func (in *DNSRecord) DeepCopyInto(out *DNSRecord) {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta

	out.Spec.Records = make([]Record, len(in.Spec.Records))
	for i := range in.Spec.Records {
		out.Spec.Records[i].Name = in.Spec.Records[i].Name
		out.Spec.Records[i].TTL = in.Spec.Records[i].TTL
		out.Spec.Records[i].Type = in.Spec.Records[i].Type
		out.Spec.Records[i].Targets = make([]string, len(in.Spec.Records[i].Targets))
		for j := range in.Spec.Records[i].Targets {
			out.Spec.Records[i].Targets[j] = in.Spec.Records[i].Targets[j]
		}
	}
}

// DeepCopyObject returns a generically typed copy of an object
func (in *DNSRecord) DeepCopyObject() runtime.Object {
	out := DNSRecord{}
	in.DeepCopyInto(&out)

	return &out
}

// DeepCopyObject returns a generically typed copy of an object
func (in *DNSRecordList) DeepCopyObject() runtime.Object {
	out := DNSRecordList{}
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta

	if in.Items != nil {
		out.Items = make([]DNSRecord, len(in.Items))
		for i := range in.Items {
			in.Items[i].DeepCopyInto(&out.Items[i])
		}
	}

	return &out
}
