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
	"context"
	"fmt"
	"maps"
	"slices"

	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/addressing"

	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/source/annotations"
)

type RegionalServicesConfig struct {
	Enabled   bool
	RegionKey string
}

var recordTypeRegionalHostnameSupported = map[string]bool{
	"A":     true,
	"AAAA":  true,
	"CNAME": true,
}

type regionalHostname struct {
	hostname  string
	regionKey string
}

// regionalHostnamesMap is a map of regional hostnames keyed by hostname.
type regionalHostnamesMap map[string]regionalHostname

type regionalHostnameChange struct {
	action changeAction
	regionalHostname
}

func (z zoneService) ListDataLocalizationRegionalHostnames(ctx context.Context, params addressing.RegionalHostnameListParams) autoPager[addressing.RegionalHostnameListResponse] {
	return z.service.Addressing.RegionalHostnames.ListAutoPaging(ctx, params)
}

func (z zoneService) CreateDataLocalizationRegionalHostname(ctx context.Context, params addressing.RegionalHostnameNewParams) error {
	_, err := z.service.Addressing.RegionalHostnames.New(ctx, params)
	return err
}

func (z zoneService) UpdateDataLocalizationRegionalHostname(ctx context.Context, hostname string, params addressing.RegionalHostnameEditParams) error {
	_, err := z.service.Addressing.RegionalHostnames.Edit(ctx, hostname, params)
	return err
}

func (z zoneService) DeleteDataLocalizationRegionalHostname(ctx context.Context, hostname string, params addressing.RegionalHostnameDeleteParams) error {
	_, err := z.service.Addressing.RegionalHostnames.Delete(ctx, hostname, params)
	return err
}

// listDataLocalizationRegionalHostnamesParams is a function that returns the appropriate RegionalHostname List Param based on the zoneID
func listDataLocalizationRegionalHostnamesParams(zoneID string) addressing.RegionalHostnameListParams {
	return addressing.RegionalHostnameListParams{
		ZoneID: cloudflare.F(zoneID),
	}
}

// createDataLocalizationRegionalHostnameParams is a function that returns the appropriate RegionalHostname Param based on the cloudFlareChange passed in
func createDataLocalizationRegionalHostnameParams(zoneID string, rhc regionalHostnameChange) addressing.RegionalHostnameNewParams {
	return addressing.RegionalHostnameNewParams{
		ZoneID:    cloudflare.F(zoneID),
		Hostname:  cloudflare.F(rhc.hostname),
		RegionKey: cloudflare.F(rhc.regionKey),
	}
}

// updateDataLocalizationRegionalHostnameParams is a function that returns the appropriate RegionalHostname Param based on the cloudFlareChange passed in
func updateDataLocalizationRegionalHostnameParams(zoneID string, rhc regionalHostnameChange) addressing.RegionalHostnameEditParams {
	return addressing.RegionalHostnameEditParams{
		ZoneID:    cloudflare.F(zoneID),
		RegionKey: cloudflare.F(rhc.regionKey),
	}
}

// deleteDataLocalizationRegionalHostnameParams is a function that returns the appropriate RegionalHostname Param based on the cloudFlareChange passed in
func deleteDataLocalizationRegionalHostnameParams(zoneID string, rhc regionalHostnameChange) addressing.RegionalHostnameDeleteParams {
	return addressing.RegionalHostnameDeleteParams{
		ZoneID: cloudflare.F(zoneID),
	}
}

// submitRegionalHostnameChanges applies a set of regional hostname changes, returns false if at least one fails
func (p *CloudFlareProvider) submitRegionalHostnameChanges(ctx context.Context, zoneID string, rhChanges []regionalHostnameChange) bool {
	failedChange := false

	for _, rhChange := range rhChanges {
		if !p.submitRegionalHostnameChange(ctx, zoneID, rhChange) {
			failedChange = true
		}
	}

	return !failedChange
}

// submitRegionalHostnameChange applies a single regional hostname change, returns false if it fails
func (p *CloudFlareProvider) submitRegionalHostnameChange(ctx context.Context, zoneID string, rhChange regionalHostnameChange) bool {
	changeLog := log.WithFields(log.Fields{
		"hostname":   rhChange.hostname,
		"region_key": rhChange.regionKey,
		"action":     rhChange.action.String(),
		"zone":       zoneID,
	})
	if p.DryRun {
		changeLog.Debug("Dry run: skipping regional hostname change", rhChange.action)
		return true
	}
	switch rhChange.action {
	case cloudFlareCreate:
		changeLog.Debug("Creating regional hostname")
		params := createDataLocalizationRegionalHostnameParams(zoneID, rhChange)
		if err := p.Client.CreateDataLocalizationRegionalHostname(ctx, params); err != nil {
			changeLog.Errorf("failed to create regional hostname: %v", err)
			return false
		}
	case cloudFlareUpdate:
		changeLog.Debug("Updating regional hostname")
		params := updateDataLocalizationRegionalHostnameParams(zoneID, rhChange)
		if err := p.Client.UpdateDataLocalizationRegionalHostname(ctx, rhChange.hostname, params); err != nil {
			changeLog.Errorf("failed to update regional hostname: %v", err)
			return false
		}
	case cloudFlareDelete:
		changeLog.Debug("Deleting regional hostname")
		params := deleteDataLocalizationRegionalHostnameParams(zoneID, rhChange)
		if err := p.Client.DeleteDataLocalizationRegionalHostname(ctx, rhChange.hostname, params); err != nil {
			changeLog.Errorf("failed to delete regional hostname: %v", err)
			return false
		}
	}
	return true
}

// listDataLocalisationRegionalHostnames fetches the current regional hostnames for the given zone ID.
//
// It returns a map of hostnames to regional hostnames, or an error if the request fails.
func (p *CloudFlareProvider) listDataLocalisationRegionalHostnames(ctx context.Context, zoneID string) (regionalHostnamesMap, error) {
	params := listDataLocalizationRegionalHostnamesParams(zoneID)
	iter := p.Client.ListDataLocalizationRegionalHostnames(ctx, params)
	rhsMap := make(regionalHostnamesMap)
	for rh := range autoPagerIterator(iter) {
		rhsMap[rh.Hostname] = regionalHostname{
			hostname:  rh.Hostname,
			regionKey: rh.RegionKey,
		}
	}
	if iter.Err() != nil {
		return nil, convertCloudflareError(iter.Err())
	}
	return rhsMap, nil
}

// regionalHostname returns a regionalHostname for the given endpoint.
//
// If the regional services feature is not enabled or the record type does not support regional hostnames,
// it returns an empty regionalHostname.
// If the endpoint has a specific region key set, it uses that; otherwise, it defaults to the region key configured in the provider.
func (p *CloudFlareProvider) regionalHostname(ep *endpoint.Endpoint) regionalHostname {
	if !p.RegionalServicesConfig.Enabled || !recordTypeRegionalHostnameSupported[ep.RecordType] {
		return regionalHostname{}
	}
	regionKey := p.RegionalServicesConfig.RegionKey
	if epRegionKey, exists := ep.GetProviderSpecificProperty(annotations.CloudflareRegionKey); exists {
		regionKey = epRegionKey
	}
	return regionalHostname{
		hostname:  ep.DNSName,
		regionKey: regionKey,
	}
}

// addEnpointsProviderSpecificRegionKeyProperty fetch the regional hostnames on cloudflare and
// adds Cloudflare-specific region keys to the provided endpoints.
//
// Do nothing if the regional services feature is not enabled.
// Defaults to the region key configured in the provider config if not found in the regional hostnames.
func (p *CloudFlareProvider) addEnpointsProviderSpecificRegionKeyProperty(ctx context.Context, zoneID string, endpoints []*endpoint.Endpoint) error {
	if !p.RegionalServicesConfig.Enabled {
		return nil
	}

	// Filter endpoints to only those that support regional hostnames
	// so we can skip regional hostname lookups if not needed.
	var supportedEndpoints []*endpoint.Endpoint
	for _, ep := range endpoints {
		if recordTypeRegionalHostnameSupported[ep.RecordType] {
			supportedEndpoints = append(supportedEndpoints, ep)
		}
	}
	if len(supportedEndpoints) == 0 {
		return nil
	}

	regionalHostnames, err := p.listDataLocalisationRegionalHostnames(ctx, zoneID)
	if err != nil {
		return err
	}

	for _, ep := range supportedEndpoints {
		var regionKey string
		if rh, found := regionalHostnames[ep.DNSName]; found {
			regionKey = rh.regionKey
		}
		ep.SetProviderSpecificProperty(annotations.CloudflareRegionKey, regionKey)
	}
	return nil
}

// adjustEnpointProviderSpecificRegionKeyProperty updates the given endpoint's provider-specific
// Cloudflare region key based on the provider's RegionalServicesConfig.
//   - If regional services are disabled or the endpoint's record type does not
//     support regional hostnames, the Cloudflare region key is removed.
//   - If enabled and supported, and the key is not already set, it is initialized
//     to the provider's default RegionKey.
//
// The endpoint is modified in place and any explicitly set region key is left unchanged.
func (p *CloudFlareProvider) adjustEnpointProviderSpecificRegionKeyProperty(ep *endpoint.Endpoint) {
	if !p.RegionalServicesConfig.Enabled || !recordTypeRegionalHostnameSupported[ep.RecordType] {
		ep.DeleteProviderSpecificProperty(annotations.CloudflareRegionKey)
		return
	}
	// Add default region key if not set
	if _, ok := ep.GetProviderSpecificProperty(annotations.CloudflareRegionKey); !ok {
		ep.SetProviderSpecificProperty(annotations.CloudflareRegionKey, p.RegionalServicesConfig.RegionKey)
	}
}

// desiredRegionalHostnames builds a list of desired regional hostnames from changes.
//
// If there is a delete and a create or update action for the same hostname,
// The create or update takes precedence.
// Returns an error for conflicting region keys.
func desiredRegionalHostnames(changes []*cloudFlareChange) ([]regionalHostname, error) {
	rhs := make(map[string]regionalHostname)
	for _, change := range changes {
		if change.RegionalHostname.hostname == "" {
			continue
		}
		rh, found := rhs[change.RegionalHostname.hostname]
		if !found {
			if change.Action == cloudFlareDelete {
				rhs[change.RegionalHostname.hostname] = regionalHostname{
					hostname:  change.RegionalHostname.hostname,
					regionKey: "", // Indicate that this regional hostname should not exists
				}
				continue
			}
			rhs[change.RegionalHostname.hostname] = change.RegionalHostname
			continue
		}
		if change.Action == cloudFlareDelete {
			// A previous regional hostname exists so we can skip this delete action
			continue
		}
		if rh.regionKey == "" {
			// If the existing regional hostname has no region key, we can overwrite it
			rhs[change.RegionalHostname.hostname] = change.RegionalHostname
			continue
		}
		if rh.regionKey != change.RegionalHostname.regionKey {
			return nil, fmt.Errorf("conflicting region keys for regional hostname %q: %q and %q", change.RegionalHostname.hostname, rh.regionKey, change.RegionalHostname.regionKey)
		}
	}
	return slices.Collect(maps.Values(rhs)), nil
}

// regionalHostnamesChanges build a list of changes needed to synchronize the current regional hostnames state with the desired state.
func regionalHostnamesChanges(desired []regionalHostname, regionalHostnames regionalHostnamesMap) []regionalHostnameChange {
	changes := make([]regionalHostnameChange, 0)
	for _, rh := range desired {
		current, found := regionalHostnames[rh.hostname]
		if rh.regionKey == "" {
			// If the region key is empty, we don't want a regional hostname
			if !found {
				continue
			}
			changes = append(changes, regionalHostnameChange{
				action:           cloudFlareDelete,
				regionalHostname: rh,
			})
			continue
		}
		if !found {
			changes = append(changes, regionalHostnameChange{
				action:           cloudFlareCreate,
				regionalHostname: rh,
			})
			continue
		}
		if rh.regionKey != current.regionKey {
			changes = append(changes, regionalHostnameChange{
				action:           cloudFlareUpdate,
				regionalHostname: rh,
			})
		}
	}
	return changes
}
