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

	cloudflare "github.com/cloudflare/cloudflare-go"
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

// RegionalHostnamesMap is a map of regional hostnames keyed by hostname.
type RegionalHostnamesMap map[string]cloudflare.RegionalHostname

type regionalHostnameChange struct {
	action changeAction
	cloudflare.RegionalHostname
}

func (z zoneService) ListDataLocalizationRegionalHostnames(ctx context.Context, rc *cloudflare.ResourceContainer, rp cloudflare.ListDataLocalizationRegionalHostnamesParams) ([]cloudflare.RegionalHostname, error) {
	return z.service.ListDataLocalizationRegionalHostnames(ctx, rc, rp)
}

func (z zoneService) CreateDataLocalizationRegionalHostname(ctx context.Context, rc *cloudflare.ResourceContainer, rp cloudflare.CreateDataLocalizationRegionalHostnameParams) error {
	_, err := z.service.CreateDataLocalizationRegionalHostname(ctx, rc, rp)
	return err
}

func (z zoneService) UpdateDataLocalizationRegionalHostname(ctx context.Context, rc *cloudflare.ResourceContainer, rp cloudflare.UpdateDataLocalizationRegionalHostnameParams) error {
	_, err := z.service.UpdateDataLocalizationRegionalHostname(ctx, rc, rp)
	return err
}

func (z zoneService) DeleteDataLocalizationRegionalHostname(ctx context.Context, rc *cloudflare.ResourceContainer, hostname string) error {
	return z.service.DeleteDataLocalizationRegionalHostname(ctx, rc, hostname)
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

// submitRegionalHostnameChanges applies a set of regional hostname changes, returns false if at least one fails
func (p *CloudFlareProvider) submitRegionalHostnameChanges(ctx context.Context, rhChanges []regionalHostnameChange, resourceContainer *cloudflare.ResourceContainer) bool {
	failedChange := false

	for _, rhChange := range rhChanges {
		if !p.submitRegionalHostnameChange(ctx, rhChange, resourceContainer) {
			failedChange = true
		}
	}

	return !failedChange
}

// submitRegionalHostnameChange applies a single regional hostname change, returns false if it fails
func (p *CloudFlareProvider) submitRegionalHostnameChange(ctx context.Context, rhChange regionalHostnameChange, resourceContainer *cloudflare.ResourceContainer) bool {
	changeLog := log.WithFields(log.Fields{
		"hostname":   rhChange.Hostname,
		"region_key": rhChange.RegionKey,
		"action":     rhChange.action,
		"zone":       resourceContainer.Identifier,
	})
	if p.DryRun {
		changeLog.Debug("Dry run: skipping regional hostname change", rhChange.action)
		return true
	}
	switch rhChange.action {
	case cloudFlareCreate:
		changeLog.Debug("Creating regional hostname")
		regionalHostnameParam := createDataLocalizationRegionalHostnameParams(rhChange)
		if err := p.Client.CreateDataLocalizationRegionalHostname(ctx, resourceContainer, regionalHostnameParam); err != nil {
			changeLog.Errorf("failed to create regional hostname: %v", err)
			return false
		}
	case cloudFlareUpdate:
		changeLog.Debug("Updating regional hostname")
		regionalHostnameParam := updateDataLocalizationRegionalHostnameParams(rhChange)
		if err := p.Client.UpdateDataLocalizationRegionalHostname(ctx, resourceContainer, regionalHostnameParam); err != nil {
			changeLog.Errorf("failed to update regional hostname: %v", err)
			return false
		}
	case cloudFlareDelete:
		changeLog.Debug("Deleting regional hostname")
		if err := p.Client.DeleteDataLocalizationRegionalHostname(ctx, resourceContainer, rhChange.Hostname); err != nil {
			changeLog.Errorf("failed to delete regional hostname: %v", err)
			return false
		}
	}
	return true
}

func (p *CloudFlareProvider) listDataLocalisationRegionalHostnames(ctx context.Context, resourceContainer *cloudflare.ResourceContainer) (RegionalHostnamesMap, error) {
	rhs, err := p.Client.ListDataLocalizationRegionalHostnames(ctx, resourceContainer, cloudflare.ListDataLocalizationRegionalHostnamesParams{})
	if err != nil {
		return nil, convertCloudflareError(err)
	}
	rhsMap := make(RegionalHostnamesMap)
	for _, r := range rhs {
		rhsMap[r.Hostname] = r
	}
	return rhsMap, nil
}

// regionalHostname returns a RegionalHostname for the given endpoint.
//
// If the regional services feature is not enabled or the record type does not support regional hostnames,
// it returns an empty RegionalHostname.
// If the endpoint has a specific region key set, it uses that; otherwise, it defaults to the region key configured in the provider.
func (p *CloudFlareProvider) regionalHostname(ep *endpoint.Endpoint) cloudflare.RegionalHostname {
	if !p.RegionalServicesConfig.Enabled || !recordTypeRegionalHostnameSupported[ep.RecordType] {
		return cloudflare.RegionalHostname{}
	}
	regionKey := p.RegionalServicesConfig.RegionKey
	if epRegionKey, exists := ep.GetProviderSpecificProperty(annotations.CloudflareRegionKey); exists {
		regionKey = epRegionKey
	}
	return cloudflare.RegionalHostname{
		Hostname:  ep.DNSName,
		RegionKey: regionKey,
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

	regionalHostnames, err := p.listDataLocalisationRegionalHostnames(ctx, cloudflare.ZoneIdentifier(zoneID))
	if err != nil {
		return err
	}

	for _, ep := range supportedEndpoints {
		if rh, found := regionalHostnames[ep.DNSName]; found {
			ep.SetProviderSpecificProperty(annotations.CloudflareRegionKey, rh.RegionKey)
		}
	}
	return nil
}

// desiredRegionalHostnames builds a list of desired regional hostnames from changes.
//
// If there is a delete and a create or update action for the same hostname,
// The create or update takes precedence.
// Returns an error for conflicting region keys.
func desiredRegionalHostnames(changes []*cloudFlareChange) ([]cloudflare.RegionalHostname, error) {
	rhs := make(map[string]cloudflare.RegionalHostname)
	for _, change := range changes {
		if change.RegionalHostname.Hostname == "" {
			continue
		}
		rh, found := rhs[change.RegionalHostname.Hostname]
		if !found {
			if change.Action == cloudFlareDelete {
				rhs[change.RegionalHostname.Hostname] = cloudflare.RegionalHostname{
					Hostname:  change.RegionalHostname.Hostname,
					RegionKey: "", // Indicate that this regional hostname should not exists
				}
				continue
			}
			rhs[change.RegionalHostname.Hostname] = change.RegionalHostname
			continue
		}
		if change.Action == cloudFlareDelete {
			// A previous regional hostname exists so we can skip this delete action
			continue
		}
		if rh.RegionKey == "" {
			// If the existing regional hostname has no region key, we can overwrite it
			rhs[change.RegionalHostname.Hostname] = change.RegionalHostname
			continue
		}
		if rh.RegionKey != change.RegionalHostname.RegionKey {
			return nil, fmt.Errorf("conflicting region keys for regional hostname %q: %q and %q", change.RegionalHostname.Hostname, rh.RegionKey, change.RegionalHostname.RegionKey)
		}
	}
	return slices.Collect(maps.Values(rhs)), nil
}

// regionalHostnamesChanges build a list of changes needed to synchronize the current regional hostnames state with the desired state.
func regionalHostnamesChanges(desired []cloudflare.RegionalHostname, regionalHostnames RegionalHostnamesMap) []regionalHostnameChange {
	changes := make([]regionalHostnameChange, 0)
	for _, rh := range desired {
		current, found := regionalHostnames[rh.Hostname]
		if rh.RegionKey == "" {
			// If the region key is empty, we don't want a regional hostname
			if !found {
				continue
			}
			changes = append(changes, regionalHostnameChange{
				action:           cloudFlareDelete,
				RegionalHostname: rh,
			})
			continue
		}
		if !found {
			changes = append(changes, regionalHostnameChange{
				action:           cloudFlareCreate,
				RegionalHostname: rh,
			})
			continue
		}
		if rh.RegionKey != current.RegionKey {
			changes = append(changes, regionalHostnameChange{
				action:           cloudFlareUpdate,
				RegionalHostname: rh,
			})
		}
	}
	return changes
}
