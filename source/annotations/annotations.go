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
	// CloudflareProxiedKey The annotation used for determining if traffic will go through Cloudflare
	CloudflareProxiedKey        = "external-dns.alpha.kubernetes.io/cloudflare-proxied"
	CloudflareCustomHostnameKey = "external-dns.alpha.kubernetes.io/cloudflare-custom-hostname"
	CloudflareRegionKey         = "external-dns.alpha.kubernetes.io/cloudflare-region-key"
	CloudflareRecordCommentKey  = "external-dns.alpha.kubernetes.io/cloudflare-record-comment"

	AWSPrefix        = "external-dns.alpha.kubernetes.io/aws-"
	SCWPrefix        = "external-dns.alpha.kubernetes.io/scw-"
	WebhookPrefix    = "external-dns.alpha.kubernetes.io/webhook-"
	CloudflarePrefix = "external-dns.alpha.kubernetes.io/cloudflare-"

	TtlKey     = "external-dns.alpha.kubernetes.io/ttl"
	ttlMinimum = 1
	ttlMaximum = math.MaxInt32

	SetIdentifierKey = "external-dns.alpha.kubernetes.io/set-identifier"
	AliasKey         = "external-dns.alpha.kubernetes.io/alias"
	TargetKey        = "external-dns.alpha.kubernetes.io/target"
	// The annotation used for figuring out which controller is responsible
	ControllerKey = "external-dns.alpha.kubernetes.io/controller"
	// The annotation used for defining the desired hostname
	HostnameKey = "external-dns.alpha.kubernetes.io/hostname"
	// The annotation used for specifying whether the public or private interface address is used
	AccessKey = "external-dns.alpha.kubernetes.io/access"
	// The annotation used for specifying the type of endpoints to use for headless services
	EndpointsTypeKey = "external-dns.alpha.kubernetes.io/endpoints-type"
	// The annotation used to determine the source of hostnames for ingresses.  This is an optional field - all
	// available hostname sources are used if not specified.
	IngressHostnameSourceKey = "external-dns.alpha.kubernetes.io/ingress-hostname-source"
	// The value of the controller annotation so that we feel responsible
	ControllerValue = "dns-controller"
	// The annotation used for defining the desired hostname
	InternalHostnameKey = "external-dns.alpha.kubernetes.io/internal-hostname"
)
