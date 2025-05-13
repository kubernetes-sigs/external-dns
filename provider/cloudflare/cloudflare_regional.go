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
	"github.com/cloudflare/cloudflare-go"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/source"
)

var recordTypeRegionalHostnameSupported = map[string]bool{
	"A":     true,
	"AAAA":  true,
	"CNAME": true,
}

type regionalHostnameChange struct {
	action changeAction
	cloudflare.RegionalHostname
}

// createDataLocalizationRegionalHostnameParams is a function that returns the appropriate RegionalHostname Param based on the cloudFlareChange passed in
func createDataLocalizationRegionalHostnameParams(rhc regionalHostnameChange) cloudflare.CreateDataLocalizationRegionalHostnameParams {
	return cloudflare.CreateDataLocalizationRegionalHostnameParams{
		Hostname:  rhc.Hostname,
		RegionKey: rhc.RegionKey,
	}
}

// updateDataLocalizationRegionalHostnameParams is a function that returns the appropriate RegionalHostname Param based on the cloudFlareChange passed in
func updateDataLocalizationRegionalHostnameParams(rhc regionalHostnameChange) cloudflare.UpdateDataLocalizationRegionalHostnameParams {
	return cloudflare.UpdateDataLocalizationRegionalHostnameParams{
		Hostname:  rhc.Hostname,
		RegionKey: rhc.RegionKey,
	}
}

func (p *CloudFlareProvider) regionalHostname(ep *endpoint.Endpoint) cloudflare.RegionalHostname {
	if p.RegionKey == "" || !recordTypeRegionalHostnameSupported[ep.RecordType] {
		return cloudflare.RegionalHostname{}
	}
	regionKey := p.RegionKey
	if epRegionKey, exists := ep.GetProviderSpecificProperty(source.CloudflareRegionKey); exists {
		regionKey = epRegionKey
	}
	return cloudflare.RegionalHostname{
		Hostname:  ep.DNSName,
		RegionKey: regionKey,
	}
}
