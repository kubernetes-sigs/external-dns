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
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

// ExternalDNSV1Alpha1Interface is an interface that wraps our clientset
type ExternalDNSV1Alpha1Interface interface {
	DNSZones(namespace string) DNSZoneInterface
	DNSRecords(namespace string) DNSRecordInterface
}

// ExternalDNSV1Alpha1Client is a client that can work with our CRDs
type ExternalDNSV1Alpha1Client struct {
	RestClient rest.Interface
}

// NewForConfig returns a new client from the provided configuration
func NewForConfig(c *rest.Config) (*ExternalDNSV1Alpha1Client, error) {

	AddToScheme(scheme.Scheme)

	config := *c
	config.ContentConfig.GroupVersion = &schema.GroupVersion{Group: GroupName, Version: GroupVersion}
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}
	config.UserAgent = rest.DefaultKubernetesUserAgent()

	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}

	return &ExternalDNSV1Alpha1Client{RestClient: client}, nil
}

// DNSZones implements ExternalDNSV1Alpha1Interface
func (c *ExternalDNSV1Alpha1Client) DNSZones(namespace string) DNSZoneInterface {
	return &DNSZoneClient{
		restClient: c.RestClient,
		ns:         namespace,
	}
}

// DNSRecords implements ExternalDNSV1Alpha1Interface
func (c *ExternalDNSV1Alpha1Client) DNSRecords(namespace string) DNSRecordInterface {
	return &DNSRecordClient{
		restClient: c.RestClient,
		ns:         namespace,
	}
}
