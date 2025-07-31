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
)

func TestUpdateDNSRecordParamV4_ARecord(t *testing.T) {
	proxied := true
	cfc := cloudFlareChange{
		ResourceRecord: cloudflare.DNSRecord{
			Name:    "test.example.com",
			Content: "192.168.1.1",
			Type:    "A",
			TTL:     300,
			Comment: "test comment",
			Proxied: &proxied,
		},
	}
	zoneID := "test-zone-id"

	params := updateDNSRecordParamV4(cfc, zoneID)

	assert.Equal(t, "test-zone-id", params.ZoneID.Value)

	// Extract the body as A record
	if aRecord, ok := params.Body.(dns.ARecordParam); ok {
		assert.Equal(t, "test.example.com", aRecord.Name.Value)
		assert.Equal(t, "192.168.1.1", aRecord.Content.Value)
		assert.InEpsilon(t, float64(dns.TTL(300)), float64(aRecord.TTL.Value), 0.01)
		assert.Equal(t, "test comment", aRecord.Comment.Value)
		assert.True(t, aRecord.Proxied.Value)
	} else {
		t.Fatalf("Expected A record type but got %T", params.Body)
	}
}

func TestUpdateDNSRecordParamV4_CNAMERecord(t *testing.T) {
	proxied := false
	cfc := cloudFlareChange{
		ResourceRecord: cloudflare.DNSRecord{
			Name:    "cname.example.com",
			Content: "target.example.com",
			Type:    "CNAME",
			TTL:     600,
			Proxied: &proxied,
		},
	}
	zoneID := "test-zone-id"

	params := updateDNSRecordParamV4(cfc, zoneID)

	assert.Equal(t, "test-zone-id", params.ZoneID.Value)

	// Extract the body as CNAME record
	if cnameRecord, ok := params.Body.(dns.CNAMERecordParam); ok {
		assert.Equal(t, "cname.example.com", cnameRecord.Name.Value)
		assert.Equal(t, "target.example.com", cnameRecord.Content.Value)
		assert.InEpsilon(t, float64(dns.TTL(600)), float64(cnameRecord.TTL.Value), 0.01)
		assert.False(t, cnameRecord.Proxied.Value)
	} else {
		t.Fatalf("Expected CNAME record type but got %T", params.Body)
	}
}

func TestUpdateDNSRecordParamV4_MXRecord(t *testing.T) {
	priority := uint16(10)
	cfc := cloudFlareChange{
		ResourceRecord: cloudflare.DNSRecord{
			Name:     "mx.example.com",
			Content:  "mail.example.com",
			Type:     "MX",
			TTL:      3600,
			Priority: &priority,
		},
	}
	zoneID := "test-zone-id"

	params := updateDNSRecordParamV4(cfc, zoneID)

	assert.Equal(t, "test-zone-id", params.ZoneID.Value)

	// Extract the body as MX record
	if mxRecord, ok := params.Body.(dns.MXRecordParam); ok {
		assert.Equal(t, "mx.example.com", mxRecord.Name.Value)
		assert.Equal(t, "mail.example.com", mxRecord.Content.Value)
		assert.InEpsilon(t, float64(dns.TTL(3600)), float64(mxRecord.TTL.Value), 0.01)
		assert.InEpsilon(t, float64(10), mxRecord.Priority.Value, 0.01)
	} else {
		t.Fatalf("Expected MX record type but got %T", params.Body)
	}
}

func TestGetCreateDNSRecordParamV4_ARecord(t *testing.T) {
	proxied := false
	cfc := cloudFlareChange{
		ResourceRecord: cloudflare.DNSRecord{
			Name:    "create.example.com",
			Content: "10.0.0.1",
			Type:    "A",
			TTL:     1200,
			Proxied: &proxied,
		},
	}
	zoneID := "create-zone-id"

	params := getCreateDNSRecordParamV4(cfc, zoneID)

	assert.Equal(t, "create-zone-id", params.ZoneID.Value)

	// Extract the body as A record
	if aRecord, ok := params.Body.(dns.ARecordParam); ok {
		assert.Equal(t, "create.example.com", aRecord.Name.Value)
		assert.Equal(t, "10.0.0.1", aRecord.Content.Value)
		assert.InEpsilon(t, float64(dns.TTL(1200)), float64(aRecord.TTL.Value), 0.01)
		assert.False(t, aRecord.Proxied.Value)
	} else {
		t.Fatalf("Expected A record type but got %T", params.Body)
	}
}

func TestGetCreateDNSRecordParamV4_TXTRecord(t *testing.T) {
	cfc := cloudFlareChange{
		ResourceRecord: cloudflare.DNSRecord{
			Name:    "txt.example.com",
			Content: "v=spf1 include:_spf.google.com ~all",
			Type:    "TXT",
			TTL:     1800,
			Comment: "SPF record",
		},
	}
	zoneID := "txt-zone-id"

	params := getCreateDNSRecordParamV4(cfc, zoneID)

	assert.Equal(t, "txt-zone-id", params.ZoneID.Value)

	// Extract the body as TXT record
	if txtRecord, ok := params.Body.(dns.TXTRecordParam); ok {
		assert.Equal(t, "txt.example.com", txtRecord.Name.Value)
		assert.Equal(t, "v=spf1 include:_spf.google.com ~all", txtRecord.Content.Value)
		assert.InEpsilon(t, float64(dns.TTL(1800)), float64(txtRecord.TTL.Value), 0.01)
		assert.Equal(t, "SPF record", txtRecord.Comment.Value)
	} else {
		t.Fatalf("Expected TXT record type but got %T", params.Body)
	}
}

func TestUpdateDNSRecordParamV4_AAAARecord(t *testing.T) {
	proxied := true
	cfc := cloudFlareChange{
		ResourceRecord: cloudflare.DNSRecord{
			Name:    "ipv6.example.com",
			Content: "2001:db8::1",
			Type:    "AAAA",
			TTL:     7200,
			Comment: "IPv6 record",
			Proxied: &proxied,
		},
	}
	zoneID := "test-zone-id"

	params := updateDNSRecordParamV4(cfc, zoneID)

	assert.Equal(t, "test-zone-id", params.ZoneID.Value)

	// Extract the body as AAAA record
	if aaaaRecord, ok := params.Body.(dns.AAAARecordParam); ok {
		assert.Equal(t, "ipv6.example.com", aaaaRecord.Name.Value)
		assert.Equal(t, "2001:db8::1", aaaaRecord.Content.Value)
		assert.InEpsilon(t, float64(dns.TTL(7200)), float64(aaaaRecord.TTL.Value), 0.01)
		assert.Equal(t, "IPv6 record", aaaaRecord.Comment.Value)
		assert.True(t, aaaaRecord.Proxied.Value)
	} else {
		t.Fatalf("Expected AAAA record type but got %T", params.Body)
	}
}

func TestUpdateDNSRecordParamV4_TXTRecord(t *testing.T) {
	cfc := cloudFlareChange{
		ResourceRecord: cloudflare.DNSRecord{
			Name:    "txt.example.com",
			Content: "v=spf1 include:_spf.google.com ~all",
			Type:    "TXT",
			TTL:     1800,
			Comment: "SPF record",
		},
	}
	zoneID := "txt-zone-id"

	params := updateDNSRecordParamV4(cfc, zoneID)

	assert.Equal(t, "txt-zone-id", params.ZoneID.Value)

	// Extract the body as TXT record
	if txtRecord, ok := params.Body.(dns.TXTRecordParam); ok {
		assert.Equal(t, "txt.example.com", txtRecord.Name.Value)
		assert.Equal(t, "v=spf1 include:_spf.google.com ~all", txtRecord.Content.Value)
		assert.InEpsilon(t, float64(dns.TTL(1800)), float64(txtRecord.TTL.Value), 0.01)
		assert.Equal(t, "SPF record", txtRecord.Comment.Value)
	} else {
		t.Fatalf("Expected TXT record type but got %T", params.Body)
	}
}

func TestUpdateDNSRecordParamV4_NoProxied(t *testing.T) {
	cfc := cloudFlareChange{
		ResourceRecord: cloudflare.DNSRecord{
			Name:    "test.example.com",
			Content: "192.168.1.1",
			Type:    "A",
			TTL:     300,
			Comment: "no proxied field",
			// Proxied is nil
		},
	}
	zoneID := "test-zone-id"

	params := updateDNSRecordParamV4(cfc, zoneID)

	assert.Equal(t, "test-zone-id", params.ZoneID.Value)

	// Extract the body as A record
	if aRecord, ok := params.Body.(dns.ARecordParam); ok {
		assert.Equal(t, "test.example.com", aRecord.Name.Value)
		assert.Equal(t, "192.168.1.1", aRecord.Content.Value)
		assert.InEpsilon(t, float64(dns.TTL(300)), float64(aRecord.TTL.Value), 0.01)
		assert.Equal(t, "no proxied field", aRecord.Comment.Value)
		// Proxied should be false when not set
		assert.False(t, aRecord.Proxied.Value)
	} else {
		t.Fatalf("Expected A record type but got %T", params.Body)
	}
}

func TestGetCreateDNSRecordParamV4_AAAARecord(t *testing.T) {
	proxied := false
	cfc := cloudFlareChange{
		ResourceRecord: cloudflare.DNSRecord{
			Name:    "ipv6.example.com",
			Content: "2001:db8::2",
			Type:    "AAAA",
			TTL:     3600,
			Comment: "IPv6 create",
			Proxied: &proxied,
		},
	}
	zoneID := "create-zone-id"

	params := getCreateDNSRecordParamV4(cfc, zoneID)

	assert.Equal(t, "create-zone-id", params.ZoneID.Value)

	// Extract the body as AAAA record
	if aaaaRecord, ok := params.Body.(dns.AAAARecordParam); ok {
		assert.Equal(t, "ipv6.example.com", aaaaRecord.Name.Value)
		assert.Equal(t, "2001:db8::2", aaaaRecord.Content.Value)
		assert.InEpsilon(t, float64(dns.TTL(3600)), float64(aaaaRecord.TTL.Value), 0.01)
		assert.Equal(t, "IPv6 create", aaaaRecord.Comment.Value)
		assert.False(t, aaaaRecord.Proxied.Value)
	} else {
		t.Fatalf("Expected AAAA record type but got %T", params.Body)
	}
}

func TestGetCreateDNSRecordParamV4_CNAMERecord(t *testing.T) {
	proxied := true
	cfc := cloudFlareChange{
		ResourceRecord: cloudflare.DNSRecord{
			Name:    "alias.example.com",
			Content: "target.example.com",
			Type:    "CNAME",
			TTL:     600,
			Comment: "alias record",
			Proxied: &proxied,
		},
	}
	zoneID := "cname-zone-id"

	params := getCreateDNSRecordParamV4(cfc, zoneID)

	assert.Equal(t, "cname-zone-id", params.ZoneID.Value)

	// Extract the body as CNAME record
	if cnameRecord, ok := params.Body.(dns.CNAMERecordParam); ok {
		assert.Equal(t, "alias.example.com", cnameRecord.Name.Value)
		assert.Equal(t, "target.example.com", cnameRecord.Content.Value)
		assert.InEpsilon(t, float64(dns.TTL(600)), float64(cnameRecord.TTL.Value), 0.01)
		assert.Equal(t, "alias record", cnameRecord.Comment.Value)
		assert.True(t, cnameRecord.Proxied.Value)
	} else {
		t.Fatalf("Expected CNAME record type but got %T", params.Body)
	}
}

func TestGetCreateDNSRecordParamV4_MXRecord(t *testing.T) {
	priority := uint16(20)
	cfc := cloudFlareChange{
		ResourceRecord: cloudflare.DNSRecord{
			Name:     "mail.example.com",
			Content:  "mailserver.example.com",
			Type:     "MX",
			TTL:      1800,
			Comment:  "mail exchange",
			Priority: &priority,
		},
	}
	zoneID := "mx-zone-id"

	params := getCreateDNSRecordParamV4(cfc, zoneID)

	assert.Equal(t, "mx-zone-id", params.ZoneID.Value)

	// Extract the body as MX record
	if mxRecord, ok := params.Body.(dns.MXRecordParam); ok {
		assert.Equal(t, "mail.example.com", mxRecord.Name.Value)
		assert.Equal(t, "mailserver.example.com", mxRecord.Content.Value)
		assert.InEpsilon(t, float64(dns.TTL(1800)), float64(mxRecord.TTL.Value), 0.01)
		assert.Equal(t, "mail exchange", mxRecord.Comment.Value)
		assert.InEpsilon(t, float64(20), mxRecord.Priority.Value, 0.01)
	} else {
		t.Fatalf("Expected MX record type but got %T", params.Body)
	}
}

func TestGetCreateDNSRecordParamV4_NoProxiedNoComment(t *testing.T) {
	cfc := cloudFlareChange{
		ResourceRecord: cloudflare.DNSRecord{
			Name:    "simple.example.com",
			Content: "10.0.0.2",
			Type:    "A",
			TTL:     300,
			// No Comment, no Proxied
		},
	}
	zoneID := "simple-zone-id"

	params := getCreateDNSRecordParamV4(cfc, zoneID)

	assert.Equal(t, "simple-zone-id", params.ZoneID.Value)

	// Extract the body as A record
	if aRecord, ok := params.Body.(dns.ARecordParam); ok {
		assert.Equal(t, "simple.example.com", aRecord.Name.Value)
		assert.Equal(t, "10.0.0.2", aRecord.Content.Value)
		assert.InEpsilon(t, float64(dns.TTL(300)), float64(aRecord.TTL.Value), 0.01)
		assert.Empty(t, aRecord.Comment.Value) // Empty comment
		assert.False(t, aRecord.Proxied.Value) // Should be false when not set
	} else {
		t.Fatalf("Expected A record type but got %T", params.Body)
	}
}
