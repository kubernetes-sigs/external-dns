//go:build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package crds

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSEntry) DeepCopyInto(out *DNSEntry) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSEntry.
func (in *DNSEntry) DeepCopy() *DNSEntry {
	if in == nil {
		return nil
	}
	out := new(DNSEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DNSEntry) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSEntryList) DeepCopyInto(out *DNSEntryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DNSEntry, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSEntryList.
func (in *DNSEntryList) DeepCopy() *DNSEntryList {
	if in == nil {
		return nil
	}
	out := new(DNSEntryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DNSEntryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSEntrySpec) DeepCopyInto(out *DNSEntrySpec) {
	*out = *in
	in.Endpoint.DeepCopyInto(&out.Endpoint)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSEntrySpec.
func (in *DNSEntrySpec) DeepCopy() *DNSEntrySpec {
	if in == nil {
		return nil
	}
	out := new(DNSEntrySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSEntryStatus) DeepCopyInto(out *DNSEntryStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSEntryStatus.
func (in *DNSEntryStatus) DeepCopy() *DNSEntryStatus {
	if in == nil {
		return nil
	}
	out := new(DNSEntryStatus)
	in.DeepCopyInto(out)
	return out
}
