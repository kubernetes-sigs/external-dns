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
	"maps"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetAnnotationPrefixes(t *testing.T) {
	CleanupAnnotationPrefixes(t)

	// Test custom prefix
	customPrefix := "custom.io/"
	extraPrefixs := []string{"custom-extra-1.io/", "custom-extra-2.io/"}
	prefixes := append([]string{customPrefix}, extraPrefixs...)
	SetAnnotationPrefixes(prefixes...)

	assert.Equal(t, customPrefix, AnnotationKeyPrefix)
	assert.Equal(t, extraPrefixs, AnnotationKeyExtraPrefixes)
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

func TestSetAnnotationPrefixes_Panic(t *testing.T) {
	CleanupAnnotationPrefixes(t)
	assert.Panics(t, func() { SetAnnotationPrefixes() })
}

func TestDefaultAnnotationPrefix(t *testing.T) {
	CleanupAnnotationPrefixes(t)
	SetAnnotationPrefixes(DefaultAnnotationPrefix)
	assert.Equal(t, DefaultAnnotationPrefix, AnnotationKeyPrefix)
	assert.Equal(t, DefaultAnnotationPrefix+"hostname", HostnameKey)
	assert.Equal(t, DefaultAnnotationPrefix+"internal-hostname", InternalHostnameKey)
	assert.Equal(t, DefaultAnnotationPrefix+"ttl", TtlKey)
	assert.Equal(t, DefaultAnnotationPrefix+"controller", ControllerKey)
}

func TestSetAnnotationPrefixMultipleTimes(t *testing.T) {
	CleanupAnnotationPrefixes(t)

	// Set first custom prefix
	SetAnnotationPrefixes("first.io/")
	assert.Equal(t, "first.io/", AnnotationKeyPrefix)
	assert.Equal(t, "first.io/hostname", HostnameKey)

	// Set second custom prefix
	SetAnnotationPrefixes("second.io/")
	assert.Equal(t, "second.io/", AnnotationKeyPrefix)
	assert.Equal(t, "second.io/hostname", HostnameKey)

	// Restore to default
	SetAnnotationPrefixes(DefaultAnnotationPrefix)
	assert.Equal(t, DefaultAnnotationPrefix, AnnotationKeyPrefix)
	assert.Equal(t, DefaultAnnotationPrefix+"hostname", HostnameKey)
}

func TestResolveAnnotations(t *testing.T) {
	CleanupAnnotationPrefixes(t)
	SetAnnotationPrefixes("external-dns.kubernetes.io/", "external-dns.alpha.kubernetes.io/")

	tests := []struct {
		name        string
		annotations map[string]string
		want        map[string]string
	}{
		{
			name:        "nil annotations",
			annotations: nil,
			want:        nil,
		},
		{
			name:        "empty annotations",
			annotations: map[string]string{},
			want:        map[string]string{},
		},
		{
			name: "no old prefixes",
			annotations: map[string]string{
				"external-dns.kubernetes.io/hostname": "example.com",
			},
			want: map[string]string{
				"external-dns.kubernetes.io/hostname": "example.com",
			},
		},
		{
			name: "migrate single annotation",
			annotations: map[string]string{
				"external-dns.alpha.kubernetes.io/hostname": "example.com",
			},
			want: map[string]string{
				"external-dns.kubernetes.io/hostname": "example.com",
			},
		},
		{
			name: "migrate multiple annotations",
			annotations: map[string]string{
				"external-dns.alpha.kubernetes.io/hostname": "example.com",
				"external-dns.alpha.kubernetes.io/ttl":      "300",
				"external-dns.alpha.kubernetes.io/target":   "target.example.com",
				"external-dns.kubernetes.io/controller":     "dns-controller",
			},
			want: map[string]string{
				"external-dns.kubernetes.io/hostname":   "example.com",
				"external-dns.kubernetes.io/ttl":        "300",
				"external-dns.kubernetes.io/target":     "target.example.com",
				"external-dns.kubernetes.io/controller": "dns-controller",
			},
		},
		{
			name: "conflicting annotations",
			annotations: map[string]string{
				"external-dns.alpha.kubernetes.io/hostname": "example.com",
				"external-dns.kubernetes.io/hostname":       "conflict.com",
			},
			want: map[string]string{
				"external-dns.kubernetes.io/hostname": "conflict.com",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			annotations := maps.Clone(tt.annotations)
			ResolveAnnotations(annotations)
			assert.Equal(t, tt.want, annotations)
		})
	}
}

func CleanupAnnotationPrefixes(t *testing.T) {
	t.Helper()
	prefixes := append([]string{AnnotationKeyPrefix}, AnnotationKeyExtraPrefixes...)
	t.Cleanup(func() {
		SetAnnotationPrefixes(prefixes...)
	})
}
