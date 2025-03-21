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

	SetIdentifierKey   = "external-dns.alpha.kubernetes.io/set-identifier"
	AliasAnnotationKey = "external-dns.alpha.kubernetes.io/alias"

	TargetAnnotationKey = "external-dns.alpha.kubernetes.io/target"

	TtlAnnotationKey = "external-dns.alpha.kubernetes.io/ttl"
	ttlMinimum       = 1
	ttlMaximum       = math.MaxInt32
)
