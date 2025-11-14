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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetAnnotationPrefix(t *testing.T) {
	// Save original values
	originalPrefix := AnnotationKeyPrefix
	defer SetAnnotationPrefix(originalPrefix)

	// Test custom prefix
	customPrefix := "custom.io/"
	SetAnnotationPrefix(customPrefix)

	assert.Equal(t, customPrefix, AnnotationKeyPrefix)
	assert.Equal(t, "custom.io/hostname", HostnameKey)
	assert.Equal(t, "custom.io/internal-hostname", InternalHostnameKey)
	assert.Equal(t, "custom.io/ttl", TtlKey)
	assert.Equal(t, "custom.io/target", TargetKey)
	assert.Equal(t, "custom.io/controller", ControllerKey)
	assert.Equal(t, "custom.io/cloudflare-proxied", CloudflareProxiedKey)
	assert.Equal(t, "custom.io/cloudflare-custom-hostname", CloudflareCustomHostnameKey)
	assert.Equal(t, "custom.io/cloudflare-region-key", CloudflareRegionKey)
	assert.Equal(t, "custom.io/cloudflare-record-comment", CloudflareRecordCommentKey)
	assert.Equal(t, "custom.io/cloudflare-tags", CloudflareTagsKey)
	assert.Equal(t, "custom.io/aws-", AWSPrefix)
	assert.Equal(t, "custom.io/coredns-", CoreDNSPrefix)
	assert.Equal(t, "custom.io/scw-", SCWPrefix)
	assert.Equal(t, "custom.io/webhook-", WebhookPrefix)
	assert.Equal(t, "custom.io/cloudflare-", CloudflarePrefix)
	assert.Equal(t, "custom.io/set-identifier", SetIdentifierKey)
	assert.Equal(t, "custom.io/alias", AliasKey)
	assert.Equal(t, "custom.io/access", AccessKey)
	assert.Equal(t, "custom.io/endpoints-type", EndpointsTypeKey)
	assert.Equal(t, "custom.io/ingress", Ingress)
	assert.Equal(t, "custom.io/ingress-hostname-source", IngressHostnameSourceKey)

	// ControllerValue should remain constant
	assert.Equal(t, "dns-controller", ControllerValue)
}

func TestDefaultAnnotationPrefix(t *testing.T) {
	assert.Equal(t, "external-dns.alpha.kubernetes.io/", AnnotationKeyPrefix)
	assert.Equal(t, "external-dns.alpha.kubernetes.io/hostname", HostnameKey)
	assert.Equal(t, "external-dns.alpha.kubernetes.io/internal-hostname", InternalHostnameKey)
	assert.Equal(t, "external-dns.alpha.kubernetes.io/ttl", TtlKey)
	assert.Equal(t, "external-dns.alpha.kubernetes.io/controller", ControllerKey)
}

func TestSetAnnotationPrefixMultipleTimes(t *testing.T) {
	// Save original values
	originalPrefix := AnnotationKeyPrefix
	defer SetAnnotationPrefix(originalPrefix)

	// Set first custom prefix
	SetAnnotationPrefix("first.io/")
	assert.Equal(t, "first.io/", AnnotationKeyPrefix)
	assert.Equal(t, "first.io/hostname", HostnameKey)

	// Set second custom prefix
	SetAnnotationPrefix("second.io/")
	assert.Equal(t, "second.io/", AnnotationKeyPrefix)
	assert.Equal(t, "second.io/hostname", HostnameKey)

	// Restore to default
	SetAnnotationPrefix("external-dns.alpha.kubernetes.io/")
	assert.Equal(t, "external-dns.alpha.kubernetes.io/", AnnotationKeyPrefix)
	assert.Equal(t, "external-dns.alpha.kubernetes.io/hostname", HostnameKey)
}
