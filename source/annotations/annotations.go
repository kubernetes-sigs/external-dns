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

package annotations

import (
	"math"
)

const (
	ttlMinimum = 1
	ttlMaximum = math.MaxInt32
)

var (
	// AnnotationKeyPrefix is set on all annotations consumed by external-dns (outside of user templates)
	// to provide easy filtering. Can be customized via SetAnnotationPrefix.
	AnnotationKeyPrefix = "external-dns.alpha.kubernetes.io/"

	// CloudflareProxiedKey The annotation used for determining if traffic will go through Cloudflare
	CloudflareProxiedKey        string
	CloudflareCustomHostnameKey string
	CloudflareRegionKey         string
	CloudflareRecordCommentKey  string
	CloudflareTagsKey           string

	AWSPrefix        string
	CoreDNSPrefix    string
	SCWPrefix        string
	WebhookPrefix    string
	CloudflarePrefix string

	TtlKey           string
	SetIdentifierKey string
	AliasKey         string
	TargetKey        string
	// ControllerKey The annotation used for figuring out which controller is responsible
	ControllerKey string
	// HostnameKey The annotation used for defining the desired hostname
	HostnameKey string
	// AccessKey The annotation used for specifying whether the public or private interface address is used
	AccessKey string
	// EndpointsTypeKey The annotation used for specifying the type of endpoints to use for headless services
	EndpointsTypeKey string
	// Ingress the annotation used to determine if the gateway is implemented by an Ingress object
	Ingress string
	// IngressHostnameSourceKey The annotation used to determine the source of hostnames for ingresses.  This is an optional field - all
	// available hostname sources are used if not specified.
	IngressHostnameSourceKey string
	// ControllerValue The value of the controller annotation so that we feel responsible
	ControllerValue = "dns-controller"
	// InternalHostnameKey The annotation used for defining the desired hostname
	InternalHostnameKey string
)

// SetAnnotationPrefix sets a custom annotation prefix and rebuilds all annotation keys.
// This must be called before any sources are initialized.
// The prefix must end with '/'.
func SetAnnotationPrefix(prefix string) {
	AnnotationKeyPrefix = prefix

	// Cloudflare annotations
	CloudflareProxiedKey = AnnotationKeyPrefix + "cloudflare-proxied"
	CloudflareCustomHostnameKey = AnnotationKeyPrefix + "cloudflare-custom-hostname"
	CloudflareRegionKey = AnnotationKeyPrefix + "cloudflare-region-key"
	CloudflareRecordCommentKey = AnnotationKeyPrefix + "cloudflare-record-comment"
	CloudflareTagsKey = AnnotationKeyPrefix + "cloudflare-tags"

	// Provider prefixes
	AWSPrefix = AnnotationKeyPrefix + "aws-"
	CoreDNSPrefix = AnnotationKeyPrefix + "coredns-"
	SCWPrefix = AnnotationKeyPrefix + "scw-"
	WebhookPrefix = AnnotationKeyPrefix + "webhook-"
	CloudflarePrefix = AnnotationKeyPrefix + "cloudflare-"

	// Core annotations
	TtlKey = AnnotationKeyPrefix + "ttl"
	SetIdentifierKey = AnnotationKeyPrefix + "set-identifier"
	AliasKey = AnnotationKeyPrefix + "alias"
	TargetKey = AnnotationKeyPrefix + "target"
	ControllerKey = AnnotationKeyPrefix + "controller"
	HostnameKey = AnnotationKeyPrefix + "hostname"
	AccessKey = AnnotationKeyPrefix + "access"
	EndpointsTypeKey = AnnotationKeyPrefix + "endpoints-type"
	Ingress = AnnotationKeyPrefix + "ingress"
	IngressHostnameSourceKey = AnnotationKeyPrefix + "ingress-hostname-source"
	InternalHostnameKey = AnnotationKeyPrefix + "internal-hostname"
}
