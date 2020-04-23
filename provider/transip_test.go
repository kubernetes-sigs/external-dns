package provider

import (
	"testing"

	"github.com/stretchr/testify/assert"
	transip "github.com/transip/gotransip/v6/domain"

	"sigs.k8s.io/external-dns/endpoint"
)

func TestTransIPDnsEntriesAreEqual(t *testing.T) {
	p := TransIPProvider{}
	// test with equal set
	a := []transip.DNSEntry{
		{
			Name:    "www.example.org",
			Type:    "CNAME",
			Expire:  3600,
			Content: "www.example.com",
		},
		{
			Name:    "www.example.com",
			Type:    "A",
			Expire:  3600,
			Content: "192.168.0.1",
		},
	}

	b := []transip.DNSEntry{
		{
			Name:    "www.example.com",
			Type:    "A",
			Expire:  3600,
			Content: "192.168.0.1",
		},
		{
			Name:    "www.example.org",
			Type:    "CNAME",
			Expire:  3600,
			Content: "www.example.com",
		},
	}

	assert.Equal(t, true, p.dnsEntriesAreEqual(a, b))

	// change type on one of b's records
	b[1].Type = "NS"
	assert.Equal(t, false, p.dnsEntriesAreEqual(a, b))
	b[1].Type = "CNAME"

	// change ttl on one of b's records
	b[1].Expire = 1800
	assert.Equal(t, false, p.dnsEntriesAreEqual(a, b))
	b[1].Expire = 3600

	// change name on one of b's records
	b[1].Name = "example.org"
	assert.Equal(t, false, p.dnsEntriesAreEqual(a, b))

	// remove last entry of b
	b = b[:1]
	assert.Equal(t, false, p.dnsEntriesAreEqual(a, b))
}

func TestTransIPGetMinimalValidTTL(t *testing.T) {
	p := TransIPProvider{}
	// test with 'unconfigured' TTL
	ep := &endpoint.Endpoint{}
	assert.Equal(t, int64(transipMinimalValidTTL), p.getMinimalValidTTL(ep))

	// test with lower than minimal ttl
	ep.RecordTTL = transipMinimalValidTTL - 1
	assert.Equal(t, int64(transipMinimalValidTTL), p.getMinimalValidTTL(ep))

	// test with higher than minimal ttl
	ep.RecordTTL = transipMinimalValidTTL + 1
	assert.Equal(t, int64(transipMinimalValidTTL+1), p.getMinimalValidTTL(ep))
}

func TestTransIPRecordNameForEndpoint(t *testing.T) {
	p := TransIPProvider{}
	ep := &endpoint.Endpoint{
		DNSName: "example.org",
	}
	d := transip.Domain{
		Name: "example.org",
	}

	assert.Equal(t, "@", p.recordNameForEndpoint(ep, d))

	ep.DNSName = "www.example.org"
	assert.Equal(t, "www", p.recordNameForEndpoint(ep, d))
}

func TestTransIPEndpointNameForRecord(t *testing.T) {
	p := TransIPProvider{}
	r := transip.DNSEntry{
		Name: "@",
	}
	d := transip.Domain{
		Name: "example.org",
	}

	assert.Equal(t, d.Name, p.endpointNameForRecord(r, d))

	r.Name = "www"
	assert.Equal(t, "www.example.org", p.endpointNameForRecord(r, d))
}

func TestTransIPAddEndpointToEntries(t *testing.T) {
	p := TransIPProvider{}

	// prepare endpoint
	ep := &endpoint.Endpoint{
		DNSName:    "www.example.org",
		RecordType: "A",
		RecordTTL:  1800,
		Targets: []string{
			"192.168.0.1",
			"192.168.0.2",
		},
	}

	// prepare zone with DNS entry set
	zone := TransIPZone{
		domain: transip.Domain{
			Name: "example.org",
		},
		// 2 matching A records
		dnsEntries: []transip.DNSEntry{
			// 1 non-matching A record
			{
				Name:    "mail",
				Type:    "A",
				Content: "192.168.0.1",
				Expire:  3600,
			},
			// 1 non-matching MX record
			{
				Name:    "@",
				Type:    "MX",
				Content: "mail.example.org",
				Expire:  3600,
			},
		},
	}

	// add endpoint to zone's entries
	result := p.addEndpointToEntries(ep, zone, zone.dnsEntries)

	assert.Equal(t, 4, len(result))
	assert.Equal(t, "mail", result[0].Name)
	assert.Equal(t, "A", result[0].Type)
	assert.Equal(t, "@", result[1].Name)
	assert.Equal(t, "MX", result[1].Type)
	assert.Equal(t, "www", result[2].Name)
	assert.Equal(t, "A", result[2].Type)
	assert.Equal(t, "192.168.0.1", result[2].Content)
	assert.Equal(t, 1800, result[2].Expire)
	assert.Equal(t, "www", result[3].Name)
	assert.Equal(t, "A", result[3].Type)
	assert.Equal(t, "192.168.0.2", result[3].Content)
	assert.Equal(t, 1800, result[3].Expire)
}

func TestTransIPRemoveEndpointFromEntries(t *testing.T) {
	p := TransIPProvider{}

	// prepare endpoint
	ep := &endpoint.Endpoint{
		DNSName:    "www.example.org",
		RecordType: "A",
	}

	// prepare zone with DNS entry set
	zone := TransIPZone{
		domain: transip.Domain{
			Name: "example.org",
		},
		// 2 matching A records
		dnsEntries: []transip.DNSEntry{
			{
				Name:    "www",
				Type:    "A",
				Content: "192.168.0.1",
				Expire:  3600,
			},
			{
				Name:    "www",
				Type:    "A",
				Content: "192.168.0.2",
				Expire:  3600,
			},
			// 1 non-matching A record
			{
				Name:    "mail",
				Type:    "A",
				Content: "192.168.0.1",
				Expire:  3600,
			},
			// 1 non-matching MX record
			{
				Name:    "@",
				Type:    "MX",
				Content: "mail.example.org",
				Expire:  3600,
			},
		},
	}

	// remove endpoint from zone's entries
	result := p.removeEndpointFromEntries(ep, zone)

	assert.Equal(t, 2, len(result))
	assert.Equal(t, "mail", result[0].Name)
	assert.Equal(t, "A", result[0].Type)
	assert.Equal(t, "@", result[1].Name)
	assert.Equal(t, "MX", result[1].Type)
}
