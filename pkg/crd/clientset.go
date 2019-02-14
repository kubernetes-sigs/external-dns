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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

type DNSZoneInterface interface {
	List(opts metav1.ListOptions) (*DNSZoneList, error)
	Get(name string, options metav1.GetOptions) (*DNSZone, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
}

type DNSZoneClient struct {
	restClient rest.Interface
	ns         string
}

func (c *DNSZoneClient) List(opts metav1.ListOptions) (*DNSZoneList, error) {
	result := DNSZoneList{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource("dnszones").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(&result)

	return &result, err
}

func (c *DNSZoneClient) Get(name string, opts metav1.GetOptions) (*DNSZone, error) {
	result := DNSZone{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource("dnszones").
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(&result)

	return &result, err
}

func (c *DNSZoneClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.restClient.
		Get().
		Namespace(c.ns).
		Resource("dnszones").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

type DNSRecordInterface interface {
	List(opts metav1.ListOptions) (*DNSRecordList, error)
	Get(name string, options metav1.GetOptions) (*DNSRecord, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
}

type DNSRecordClient struct {
	restClient rest.Interface
	ns         string
}

func (c *DNSRecordClient) List(opts metav1.ListOptions) (*DNSRecordList, error) {
	result := DNSRecordList{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource("dnsrecords").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(&result)

	return &result, err
}

func (c *DNSRecordClient) Get(name string, opts metav1.GetOptions) (*DNSRecord, error) {
	result := DNSRecord{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource("dnsrecords").
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(&result)

	return &result, err
}

func (c *DNSRecordClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.restClient.
		Get().
		Namespace(c.ns).
		Resource("dnsrecords").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}
