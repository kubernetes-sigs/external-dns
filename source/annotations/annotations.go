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
	"strings"
)

const (
	// DefaultAnnotationPrefix is the default annotation prefix used by external-dns
	DefaultAnnotationPrefix = "external-dns.kubernetes.io/"

	// AlphaAnnotationPrefix is the legacy annotation prefix used by external-dns, still supported for backward compatibility
	AlphaAnnotationPrefix = "external-dns.alpha.kubernetes.io/"

	ttlMinimum = 1
	ttlMaximum = math.MaxInt32
)

var (
	// AnnotationKeyPrefix is set on all annotations consumed by external-dns (outside of user templates)
	// to provide easy filtering. Can be customized via SetAnnotationPrefix.
	AnnotationKeyPrefix        = DefaultAnnotationPrefix
	AnnotationKeyExtraPrefixes = []string{AlphaAnnotationPrefix}

	// CloudflareProxiedKey The annotation used for determining if traffic will go through Cloudflare
	CloudflareProxiedKey        = AnnotationKeyPrefix + "cloudflare-proxied"
	CloudflareCustomHostnameKey = AnnotationKeyPrefix + "cloudflare-custom-hostname"
	CloudflareRegionKey         = AnnotationKeyPrefix + "cloudflare-region-key"
	CloudflareRecordCommentKey  = AnnotationKeyPrefix + "cloudflare-record-comment"
	CloudflareTagsKey           = AnnotationKeyPrefix + "cloudflare-tags"

	// AzureTagsKey The annotation used for Azure DNS record tags
	AzureTagsKey = AnnotationKeyPrefix + "azure-tags"

	AWSPrefix        = AnnotationKeyPrefix + "aws-"
	CoreDNSPrefix    = AnnotationKeyPrefix + "coredns-"
	SCWPrefix        = AnnotationKeyPrefix + "scw-"
	WebhookPrefix    = AnnotationKeyPrefix + "webhook-"
	CloudflarePrefix = AnnotationKeyPrefix + "cloudflare-"

	TtlKey           = AnnotationKeyPrefix + "ttl"
	SetIdentifierKey = AnnotationKeyPrefix + "set-identifier"
	AliasKey         = AnnotationKeyPrefix + "alias"
	RecordTypeKey    = AnnotationKeyPrefix + "record-type"
	TargetKey        = AnnotationKeyPrefix + "target"
	// ControllerKey The annotation used for figuring out which controller is responsible
	ControllerKey = AnnotationKeyPrefix + "controller"
	// HostnameKey The annotation used for defining the desired hostname
	HostnameKey = AnnotationKeyPrefix + "hostname"
	// AccessKey The annotation used for specifying whether the public or private interface address is used
	AccessKey = AnnotationKeyPrefix + "access"
	// EndpointsTypeKey The annotation used for specifying the type of endpoints to use for headless services
	EndpointsTypeKey = AnnotationKeyPrefix + "endpoints-type"
	// Ingress the annotation used to determine if the gateway is implemented by an Ingress object
	Ingress = AnnotationKeyPrefix + "ingress"
	// IngressHostnameSourceKey The annotation used to determine the source of hostnames for ingresses.  This is an optional field - all
	// available hostname sources are used if not specified.
	IngressHostnameSourceKey = AnnotationKeyPrefix + "ingress-hostname-source"
	// ControllerValue The value of the controller annotation so that we feel responsible
	ControllerValue = "dns-controller"
	// InternalHostnameKey The annotation used for defining the desired hostname
	InternalHostnameKey = AnnotationKeyPrefix + "internal-hostname"
	// The annotation used for defining the desired hostname source for gateways
	GatewayHostnameSourceKey = AnnotationKeyPrefix + "gateway-hostname-source"
)

// SetAnnotationPrefixes sets custom annotation prefixes and rebuilds all annotation keys with the first one.
// This must be called before any sources are initialized.
// All prefixes must end with '/'.
func SetAnnotationPrefixes(prefixes ...string) {
	if len(prefixes) == 0 {
		panic("at least one annotation prefix must be provided")
	}
	AnnotationKeyPrefix = prefixes[0]
	AnnotationKeyExtraPrefixes = prefixes[1:]

	// Cloudflare annotations
	CloudflareProxiedKey = AnnotationKeyPrefix + "cloudflare-proxied"
	CloudflareCustomHostnameKey = AnnotationKeyPrefix + "cloudflare-custom-hostname"
	CloudflareRegionKey = AnnotationKeyPrefix + "cloudflare-region-key"
	CloudflareRecordCommentKey = AnnotationKeyPrefix + "cloudflare-record-comment"
	CloudflareTagsKey = AnnotationKeyPrefix + "cloudflare-tags"

	// Azure annotations
	AzureTagsKey = AnnotationKeyPrefix + "azure-tags"

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
	RecordTypeKey = AnnotationKeyPrefix + "record-type"
	TargetKey = AnnotationKeyPrefix + "target"
	ControllerKey = AnnotationKeyPrefix + "controller"
	HostnameKey = AnnotationKeyPrefix + "hostname"
	AccessKey = AnnotationKeyPrefix + "access"
	EndpointsTypeKey = AnnotationKeyPrefix + "endpoints-type"
	Ingress = AnnotationKeyPrefix + "ingress"
	IngressHostnameSourceKey = AnnotationKeyPrefix + "ingress-hostname-source"
	InternalHostnameKey = AnnotationKeyPrefix + "internal-hostname"
	GatewayHostnameSourceKey = AnnotationKeyPrefix + "gateway-hostname-source"
}

// ResolveAnnotations convert annotations with extra prefixes into their corresponding keys with the main prefix.
//
// This allows users to specify annotations with any of the supported prefixes and have them normalized
// to the main prefix, to be consumed uniformly by external-dns regardless of which prefix was used.
// For example, with the default configuration, an annotation like "external-dns.alpha.kubernetes.io/hostname"
// would be transformed to "external-dns.kubernetes.io/hostname".
//
// This function should be called on the annotations map of each relevant Kubernetes object before external-dns processes it.
// This is typically done in the informer's transform function.
//
// In case of conflicts, the value of the first occurring prefix takes precedence.
func ResolveAnnotations(annotations map[string]string) {
	if len(AnnotationKeyExtraPrefixes) == 0 || len(annotations) == 0 {
		return
	}
	for _, prefix := range AnnotationKeyExtraPrefixes {
		for k, v := range annotations {
			if after, ok := strings.CutPrefix(k, prefix); ok {
				delete(annotations, k)
				newKey := AnnotationKeyPrefix + after
				if _, ok := annotations[newKey]; !ok {
					annotations[newKey] = v
				}
			}
		}
	}
}
