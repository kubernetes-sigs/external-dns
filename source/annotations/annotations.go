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
	// AnnotationKeyPrefix is set on all annotations consumed by external-dns (outside of user templates)
	// to provide easy filtering.
	AnnotationKeyPrefix = "external-dns.alpha.kubernetes.io/"

	// CloudflareProxiedKey The annotation used for determining if traffic will go through Cloudflare
	CloudflareProxiedKey        = AnnotationKeyPrefix + "cloudflare-proxied"
	CloudflareCustomHostnameKey = AnnotationKeyPrefix + "cloudflare-custom-hostname"
	CloudflareRegionKey         = AnnotationKeyPrefix + "cloudflare-region-key"
	CloudflareRecordCommentKey  = AnnotationKeyPrefix + "cloudflare-record-comment"

	AWSPrefix        = AnnotationKeyPrefix + "aws-"
	SCWPrefix        = AnnotationKeyPrefix + "scw-"
	WebhookPrefix    = AnnotationKeyPrefix + "webhook-"
	CloudflarePrefix = AnnotationKeyPrefix + "cloudflare-"

	TtlKey     = AnnotationKeyPrefix + "ttl"
	ttlMinimum = 1
	ttlMaximum = math.MaxInt32

	SetIdentifierKey = AnnotationKeyPrefix + "set-identifier"
	AliasKey         = AnnotationKeyPrefix + "alias"
	TargetKey        = AnnotationKeyPrefix + "target"
	// The annotation used for figuring out which controller is responsible
	ControllerKey = AnnotationKeyPrefix + "controller"
	// The annotation used for defining the desired hostname
	HostnameKey = AnnotationKeyPrefix + "hostname"
	// The annotation used for specifying whether the public or private interface address is used
	AccessKey = AnnotationKeyPrefix + "access"
	// The annotation used for specifying the type of endpoints to use for headless services
	EndpointsTypeKey = AnnotationKeyPrefix + "endpoints-type"
	// The annotation used to determine the source of hostnames for ingresses.  This is an optional field - all
	// available hostname sources are used if not specified.
	IngressHostnameSourceKey = AnnotationKeyPrefix + "ingress-hostname-source"
	// The value of the controller annotation so that we feel responsible
	ControllerValue = "dns-controller"
	// The annotation used for defining the desired hostname
	InternalHostnameKey = AnnotationKeyPrefix + "internal-hostname"
)
