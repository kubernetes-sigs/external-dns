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
	cloudflarev4 "github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/dns"
)

// updateDNSRecordParamV4 converts cloudFlareChange to v4 SDK UpdateDNSRecordParams
func updateDNSRecordParamV4(cfc cloudFlareChange, zoneID string) dns.RecordUpdateParams {
	params := dns.RecordUpdateParams{
		ZoneID: cloudflarev4.F(zoneID),
	}

	switch cfc.ResourceRecord.Type {
	case "A":
		aRecord := dns.ARecordParam{
			Name:    cloudflarev4.F(cfc.ResourceRecord.Name),
			Content: cloudflarev4.F(cfc.ResourceRecord.Content),
			TTL:     cloudflarev4.F(dns.TTL(cfc.ResourceRecord.TTL)),
			Comment: cloudflarev4.F(cfc.ResourceRecord.Comment),
		}
		if cfc.ResourceRecord.Proxied != nil {
			aRecord.Proxied = cloudflarev4.F(*cfc.ResourceRecord.Proxied)
		}
		params.Body = aRecord
	case "AAAA":
		aaaaRecord := dns.AAAARecordParam{
			Name:    cloudflarev4.F(cfc.ResourceRecord.Name),
			Content: cloudflarev4.F(cfc.ResourceRecord.Content),
			TTL:     cloudflarev4.F(dns.TTL(cfc.ResourceRecord.TTL)),
			Comment: cloudflarev4.F(cfc.ResourceRecord.Comment),
		}
		if cfc.ResourceRecord.Proxied != nil {
			aaaaRecord.Proxied = cloudflarev4.F(*cfc.ResourceRecord.Proxied)
		}
		params.Body = aaaaRecord
	case "CNAME":
		cnameRecord := dns.CNAMERecordParam{
			Name:    cloudflarev4.F(cfc.ResourceRecord.Name),
			Content: cloudflarev4.F(cfc.ResourceRecord.Content),
			TTL:     cloudflarev4.F(dns.TTL(cfc.ResourceRecord.TTL)),
			Comment: cloudflarev4.F(cfc.ResourceRecord.Comment),
		}
		if cfc.ResourceRecord.Proxied != nil {
			cnameRecord.Proxied = cloudflarev4.F(*cfc.ResourceRecord.Proxied)
		}
		params.Body = cnameRecord
	case "MX":
		mxRecord := dns.MXRecordParam{
			Name:    cloudflarev4.F(cfc.ResourceRecord.Name),
			Content: cloudflarev4.F(cfc.ResourceRecord.Content),
			TTL:     cloudflarev4.F(dns.TTL(cfc.ResourceRecord.TTL)),
			Comment: cloudflarev4.F(cfc.ResourceRecord.Comment),
		}
		if cfc.ResourceRecord.Priority != nil {
			mxRecord.Priority = cloudflarev4.F(float64(*cfc.ResourceRecord.Priority))
		}
		params.Body = mxRecord
	case "TXT":
		params.Body = dns.TXTRecordParam{
			Name:    cloudflarev4.F(cfc.ResourceRecord.Name),
			Content: cloudflarev4.F(cfc.ResourceRecord.Content),
			TTL:     cloudflarev4.F(dns.TTL(cfc.ResourceRecord.TTL)),
			Comment: cloudflarev4.F(cfc.ResourceRecord.Comment),
		}
	}

	return params
}

// getCreateDNSRecordParamV4 converts cloudFlareChange to v4 SDK CreateDNSRecordParams
func getCreateDNSRecordParamV4(cfc cloudFlareChange, zoneID string) dns.RecordNewParams {
	params := dns.RecordNewParams{
		ZoneID: cloudflarev4.F(zoneID),
	}

	switch cfc.ResourceRecord.Type {
	case "A":
		aRecord := dns.ARecordParam{
			Name:    cloudflarev4.F(cfc.ResourceRecord.Name),
			Content: cloudflarev4.F(cfc.ResourceRecord.Content),
			TTL:     cloudflarev4.F(dns.TTL(cfc.ResourceRecord.TTL)),
			Comment: cloudflarev4.F(cfc.ResourceRecord.Comment),
		}
		if cfc.ResourceRecord.Proxied != nil {
			aRecord.Proxied = cloudflarev4.F(*cfc.ResourceRecord.Proxied)
		}
		params.Body = aRecord
	case "AAAA":
		aaaaRecord := dns.AAAARecordParam{
			Name:    cloudflarev4.F(cfc.ResourceRecord.Name),
			Content: cloudflarev4.F(cfc.ResourceRecord.Content),
			TTL:     cloudflarev4.F(dns.TTL(cfc.ResourceRecord.TTL)),
			Comment: cloudflarev4.F(cfc.ResourceRecord.Comment),
		}
		if cfc.ResourceRecord.Proxied != nil {
			aaaaRecord.Proxied = cloudflarev4.F(*cfc.ResourceRecord.Proxied)
		}
		params.Body = aaaaRecord
	case "CNAME":
		cnameRecord := dns.CNAMERecordParam{
			Name:    cloudflarev4.F(cfc.ResourceRecord.Name),
			Content: cloudflarev4.F(cfc.ResourceRecord.Content),
			TTL:     cloudflarev4.F(dns.TTL(cfc.ResourceRecord.TTL)),
			Comment: cloudflarev4.F(cfc.ResourceRecord.Comment),
		}
		if cfc.ResourceRecord.Proxied != nil {
			cnameRecord.Proxied = cloudflarev4.F(*cfc.ResourceRecord.Proxied)
		}
		params.Body = cnameRecord
	case "MX":
		mxRecord := dns.MXRecordParam{
			Name:    cloudflarev4.F(cfc.ResourceRecord.Name),
			Content: cloudflarev4.F(cfc.ResourceRecord.Content),
			TTL:     cloudflarev4.F(dns.TTL(cfc.ResourceRecord.TTL)),
			Comment: cloudflarev4.F(cfc.ResourceRecord.Comment),
		}
		if cfc.ResourceRecord.Priority != nil {
			mxRecord.Priority = cloudflarev4.F(float64(*cfc.ResourceRecord.Priority))
		}
		params.Body = mxRecord
	case "TXT":
		params.Body = dns.TXTRecordParam{
			Name:    cloudflarev4.F(cfc.ResourceRecord.Name),
			Content: cloudflarev4.F(cfc.ResourceRecord.Content),
			TTL:     cloudflarev4.F(dns.TTL(cfc.ResourceRecord.TTL)),
			Comment: cloudflarev4.F(cfc.ResourceRecord.Comment),
		}
	}

	return params
}
