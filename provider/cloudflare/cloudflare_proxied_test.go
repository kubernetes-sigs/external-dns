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

package cloudflare

import (
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/cloudflare-go/v4/dns"
	"github.com/stretchr/testify/assert"

	"sigs.k8s.io/external-dns/internal/testutils"
)

func TestUpdateDNSRecordParamV4_ARecord(t *testing.T) {
	change := cloudFlareChange{
		ResourceRecord: cloudflare.DNSRecord{
			ZoneID:  "zone123",
			Name:    "test.example.com",
			Type:    "A",
			Content: "192.168.1.1",
			TTL:     300,
			Proxied: testutils.ToPtr(true),
			Comment: "Test record",
		},
	}

	result := updateDNSRecordParamV4(change)

	assert.Equal(t, "zone123", *result.ZoneID.Value)
	
	// Check that the body is an A record
	aRecord, ok := result.Body.(dns.ARecordParam)
	assert.True(t, ok, "Body should be an ARecordParam")
	assert.Equal(t, "test.example.com", *aRecord.Name.Value)
	assert.Equal(t, "192.168.1.1", *aRecord.Content.Value)
	assert.Equal(t, dns.TTL(300), aRecord.TTL)
	assert.Equal(t, true, *aRecord.Proxied.Value)
	assert.Equal(t, "Test record", *aRecord.Comment.Value)
}
