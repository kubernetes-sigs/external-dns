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

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	logtest "sigs.k8s.io/external-dns/internal/testutils/log"
)

func TestSetAnnotationPrefix(t *testing.T) {
	t.Cleanup(func() { SetAnnotationPrefix(DefaultAnnotationPrefix) })

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
	t.Cleanup(func() { SetAnnotationPrefix(DefaultAnnotationPrefix) })
	SetAnnotationPrefix(DefaultAnnotationPrefix)
	assert.Equal(t, DefaultAnnotationPrefix, AnnotationKeyPrefix)
	assert.Equal(t, DefaultAnnotationPrefix+"hostname", HostnameKey)
	assert.Equal(t, DefaultAnnotationPrefix+"internal-hostname", InternalHostnameKey)
	assert.Equal(t, DefaultAnnotationPrefix+"ttl", TtlKey)
	assert.Equal(t, DefaultAnnotationPrefix+"controller", ControllerKey)
}

func TestSetAnnotationPrefixMultipleTimes(t *testing.T) {
	t.Cleanup(func() { SetAnnotationPrefix(DefaultAnnotationPrefix) })

	// Set first custom prefix
	SetAnnotationPrefix("first.io/")
	assert.Equal(t, "first.io/", AnnotationKeyPrefix)
	assert.Equal(t, "first.io/hostname", HostnameKey)

	// Set second custom prefix
	SetAnnotationPrefix("second.io/")
	assert.Equal(t, "second.io/", AnnotationKeyPrefix)
	assert.Equal(t, "second.io/hostname", HostnameKey)

	// Restore to default
	SetAnnotationPrefix(DefaultAnnotationPrefix)
	assert.Equal(t, DefaultAnnotationPrefix, AnnotationKeyPrefix)
	assert.Equal(t, DefaultAnnotationPrefix+"hostname", HostnameKey)
}

func TestSetLegacyAnnotationPrefix(t *testing.T) {
	t.Cleanup(func() { SetLegacyAnnotationPrefix("") })

	anns := map[string]string{LegacyAnnotationPrefix + "hostname": "example.org"}

	t.Run("disabled by default", func(t *testing.T) {
		assert.False(t, ResolveLegacyAnnotations("Service", "default", "svc", anns))
		assert.Equal(t, map[string]string{LegacyAnnotationPrefix + "hostname": "example.org"}, anns)
	})

	t.Run("a legacy prefix equal to the configured prefix disables resolution", func(t *testing.T) {
		SetLegacyAnnotationPrefix(AnnotationKeyPrefix)
		assert.False(t, ResolveLegacyAnnotations("Service", "default", "svc", anns))
		assert.Equal(t, map[string]string{LegacyAnnotationPrefix + "hostname": "example.org"}, anns)
	})
}

func TestResolveLegacyAnnotations(t *testing.T) {
	SetLegacyAnnotationPrefix(LegacyAnnotationPrefix)
	t.Cleanup(func() { SetLegacyAnnotationPrefix("") })

	for _, tc := range []struct {
		name        string
		annotations map[string]string
		expected    map[string]string
		resolved    bool
		warning     string
	}{
		{
			name: "nil annotations",
		},
		{
			name:        "no legacy annotations",
			annotations: map[string]string{HostnameKey: "example.org", "description": "kept"},
			expected:    map[string]string{HostnameKey: "example.org", "description": "kept"},
		},
		{
			name:        "legacy annotation is rewritten to the configured prefix",
			annotations: map[string]string{LegacyAnnotationPrefix + "hostname": "example.org"},
			expected:    map[string]string{HostnameKey: "example.org"},
			resolved:    true,
		},
		{
			name: "every legacy annotation is rewritten, other annotations are untouched",
			annotations: map[string]string{
				LegacyAnnotationPrefix + "hostname": "example.org",
				LegacyAnnotationPrefix + "ttl":      "60",
				TargetKey:                           "1.2.3.4",
				"description":                       "kept",
			},
			expected: map[string]string{
				HostnameKey:   "example.org",
				TtlKey:        "60",
				TargetKey:     "1.2.3.4",
				"description": "kept",
			},
			resolved: true,
		},
		{
			name: "same value under both prefixes is not a conflict",
			annotations: map[string]string{
				LegacyAnnotationPrefix + "hostname": "example.org",
				HostnameKey:                         "example.org",
			},
			expected: map[string]string{HostnameKey: "example.org"},
			resolved: true,
		},
		{
			name: "the configured prefix wins on conflict and the conflict is logged",
			annotations: map[string]string{
				LegacyAnnotationPrefix + "hostname": "legacy.example.org",
				HostnameKey:                         "example.org",
			},
			expected: map[string]string{HostnameKey: "example.org"},
			resolved: true,
			warning:  `Ingress default/web: ignoring annotation external-dns.alpha.kubernetes.io/hostname="legacy.example.org" because external-dns.kubernetes.io/hostname="example.org" is set`,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			hook := logtest.LogsUnderTestWithLogLevel(log.WarnLevel, t)

			assert.Equal(t, tc.resolved, ResolveLegacyAnnotations("Ingress", "default", "web", tc.annotations))
			assert.Equal(t, tc.expected, tc.annotations)
			if tc.warning != "" {
				logtest.TestHelperLogContains(tc.warning, hook, t)
			} else {
				assert.Empty(t, hook.AllEntries())
			}
		})
	}
}
