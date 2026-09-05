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

	log "github.com/sirupsen/logrus"
)

const (
	// DefaultAnnotationPrefix is the default annotation prefix used by external-dns
	DefaultAnnotationPrefix = "external-dns.kubernetes.io/"

	// LegacyAnnotationPrefix is the annotation prefix used by releases before v0.22.0.
	// It is only honored when enabled with SetLegacyAnnotationPrefix, see ResolveLegacyAnnotations.
	LegacyAnnotationPrefix = "external-dns.alpha.kubernetes.io/"

	ttlMinimum = 1
	ttlMaximum = math.MaxInt32
)

var (
	// AnnotationKeyPrefix is set on all annotations consumed by external-dns (outside of user templates)
	// to provide easy filtering. Can be customized via SetAnnotationPrefix.
	AnnotationKeyPrefix = DefaultAnnotationPrefix

	// legacyAnnotationPrefix is an additional prefix accepted by ResolveLegacyAnnotations. Empty means disabled.
	legacyAnnotationPrefix string

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

// SetLegacyAnnotationPrefix makes ResolveLegacyAnnotations accept annotations carrying prefix in addition
// to AnnotationKeyPrefix. An empty prefix, or one equal to AnnotationKeyPrefix, disables the resolution.
// Like SetAnnotationPrefix, this must be called before any sources are initialized.
func SetLegacyAnnotationPrefix(prefix string) {
	if prefix == AnnotationKeyPrefix {
		prefix = ""
	}
	legacyAnnotationPrefix = prefix
}

// ResolveLegacyAnnotations rewrites, in place, every annotation carrying the legacy prefix to its
// AnnotationKeyPrefix equivalent so that sources only ever see the configured prefix. When both forms
// of a key are present the value under AnnotationKeyPrefix wins and the conflict is logged; kind,
// namespace and name identify the object in that log line. It returns true when anns was modified
// and is a no-op unless a legacy prefix was enabled with SetLegacyAnnotationPrefix.
func ResolveLegacyAnnotations(kind, namespace, name string, anns map[string]string) bool {
	if legacyAnnotationPrefix == "" {
		return false
	}
	var legacyKeys []string
	for key := range anns {
		if strings.HasPrefix(key, legacyAnnotationPrefix) {
			legacyKeys = append(legacyKeys, key)
		}
	}
	for _, key := range legacyKeys {
		value := anns[key]
		delete(anns, key)
		resolved := AnnotationKeyPrefix + strings.TrimPrefix(key, legacyAnnotationPrefix)
		if existing, ok := anns[resolved]; !ok {
			anns[resolved] = value
		} else if existing != value {
			log.Warnf("%s %s/%s: ignoring annotation %s=%q because %s=%q is set", kind, namespace, name, key, value, resolved, existing)
		}
	}
	return len(legacyKeys) > 0
}
