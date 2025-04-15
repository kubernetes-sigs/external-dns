/*
Copyright 2017 The Kubernetes Authors.

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
	"errors"
	"fmt"
	"maps"
	"net/http"
	"slices"

	cloudflare "github.com/cloudflare/cloudflare-go"
	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/source"
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

type RegionalHostnamesMap map[string]cloudflare.RegionalHostname

type RegionalHostnameChange struct {
	action changeAction
	cloudflare.RegionalHostname
}

// createDataLocalizationRegionalHostnameParams is a function that returns the appropriate RegionalHostname Param based on the cloudFlareChange passed in
func createDataLocalizationRegionalHostnameParams(rhc RegionalHostnameChange) cloudflare.CreateDataLocalizationRegionalHostnameParams {
	return cloudflare.CreateDataLocalizationRegionalHostnameParams{
		Hostname:  rhc.Hostname,
		RegionKey: rhc.RegionKey,
	}
}

// updateDataLocalizationRegionalHostnameParams is a function that returns the appropriate RegionalHostname Param based on the cloudFlareChange passed in
func updateDataLocalizationRegionalHostnameParams(rhc RegionalHostnameChange) cloudflare.UpdateDataLocalizationRegionalHostnameParams {
	return cloudflare.UpdateDataLocalizationRegionalHostnameParams{
		Hostname:  rhc.Hostname,
		RegionKey: rhc.RegionKey,
	}
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

// submitRegionalHostnameChanges applies a set of regional hostname changes, returns false if it fails
func (p *CloudFlareProvider) submitRegionalHostnameChanges(ctx context.Context, rhChanges []RegionalHostnameChange, resourceContainer *cloudflare.ResourceContainer) bool {
	failedChange := false

	for _, rhChange := range rhChanges {
		logFields := log.Fields{
			"hostname":   rhChange.Hostname,
			"region_key": rhChange.RegionKey,
			"action":     rhChange.action,
			"zone":       resourceContainer.Identifier,
		}
		switch rhChange.action {
		case cloudFlareCreate:
			log.WithFields(logFields).Debug("Creating regional hostname")
			if p.DryRun {
				continue
			}
			regionalHostnameParam := createDataLocalizationRegionalHostnameParams(rhChange)
			err := p.Client.CreateDataLocalizationRegionalHostname(ctx, resourceContainer, regionalHostnameParam)
			if err != nil {
				failedChange = true
				log.WithFields(logFields).Errorf("failed to create regional hostname: %v", err)
			}
		case cloudFlareUpdate:
			log.WithFields(logFields).Debug("Updating regional hostname")
			if p.DryRun {
				continue
			}
			regionalHostnameParam := updateDataLocalizationRegionalHostnameParams(rhChange)
			err := p.Client.UpdateDataLocalizationRegionalHostname(ctx, resourceContainer, regionalHostnameParam)
			if err != nil {
				failedChange = true
				log.WithFields(logFields).Errorf("failed to update regional hostname: %v", err)
			}
		case cloudFlareDelete:
			log.WithFields(logFields).Debug("Deleting regional hostname")
			if p.DryRun {
				continue
			}
			err := p.Client.DeleteDataLocalizationRegionalHostname(ctx, resourceContainer, rhChange.Hostname)
			if err != nil {
				failedChange = true
				log.WithFields(logFields).Errorf("failed to delete regional hostname: %v", err)
			}
		}
	}

	return !failedChange
}

func (p *CloudFlareProvider) listDatalocalisationRegionalHostnames(ctx context.Context, resourceContainer *cloudflare.ResourceContainer) (RegionalHostnamesMap, error) {
	rhs, err := p.Client.ListDataLocalizationRegionalHostnames(ctx, resourceContainer, cloudflare.ListDataLocalizationRegionalHostnamesParams{})
	if err != nil {
		var apiErr *cloudflare.Error
		if errors.As(err, &apiErr) {
			if apiErr.ClientRateLimited() || apiErr.StatusCode >= http.StatusInternalServerError {
				// Handle rate limit error as a soft error
				return nil, provider.NewSoftError(err)
			}
		}
		return nil, err
	}
	rhsMap := make(RegionalHostnamesMap)
	for _, r := range rhs {
		rhsMap[r.Hostname] = r
	}
	return rhsMap, nil
}

func (p *CloudFlareProvider) regionalHostname(ep *endpoint.Endpoint) cloudflare.RegionalHostname {
	if !p.RegionalServicesConfig.Enabled || !recordTypeRegionalHostnameSupported[ep.RecordType] {
		return cloudflare.RegionalHostname{}
	}
	regionKey := p.RegionalServicesConfig.RegionKey
	for _, v := range ep.ProviderSpecific {
		if v.Name == source.CloudflareRegionKey {
			regionKey = v.Value
			break
		}
	}
	return cloudflare.RegionalHostname{
		Hostname:  ep.DNSName,
		RegionKey: regionKey,
	}
}

// addEnpointsProviderSpecificRegionKeyProperty fetch the regional hostnames on cloudflare and
// adds Cloudflare-specific region keys to the provided endpoints.
// Do nothing if the regional services feature is not enabled.
// Defaults to the region key configured in the provider config if not found in the regional hostnames.
func (p *CloudFlareProvider) addEnpointsProviderSpecificRegionKeyProperty(ctx context.Context, zoneID string, endpoints []*endpoint.Endpoint) error {
	if !p.RegionalServicesConfig.Enabled {
		return nil
	}

	// Filter unsupported record types
	endpoints = slices.Collect(func(yield func(*endpoint.Endpoint) bool) {
		for _, ep := range endpoints {
			if !recordTypeRegionalHostnameSupported[ep.RecordType] {
				continue
			}
			if !yield(ep) {
				return
			}
		}
	})
	if len(endpoints) == 0 {
		return nil
	}

	regionalHostnames, err := p.listDatalocalisationRegionalHostnames(ctx, cloudflare.ZoneIdentifier(zoneID))
	if err != nil {
		return err
	}

	for _, ep := range endpoints {
		regionKey := p.RegionalServicesConfig.RegionKey
		if rh, found := regionalHostnames[ep.DNSName]; found {
			regionKey = rh.RegionKey
		}
		ep.SetProviderSpecificProperty(source.CloudflareRegionKey, regionKey)
	}
	return nil
}

// desiredRegionalHostnames builds a list of desired regional hostnames from changes.
// Returns an error for empty or conflicting region keys.
func desiredRegionalHostnames(changes []*cloudFlareChange) ([]cloudflare.RegionalHostname, error) {
	rhs := make(map[string]cloudflare.RegionalHostname)
	for _, change := range changes {
		if change.RegionalHostname.Hostname == "" {
			continue
		}
		rh, found := rhs[change.RegionalHostname.Hostname]
		if !found {
			if change.Action == cloudFlareDelete {
				rhs[change.RegionalHostname.Hostname] = cloudflare.RegionalHostname{Hostname: change.RegionalHostname.Hostname}
				continue
			}
			rhs[change.RegionalHostname.Hostname] = change.RegionalHostname
			continue
		}
		if change.Action == cloudFlareDelete {
			continue
		}
		if rh.RegionKey == "" {
			rhs[change.RegionalHostname.Hostname] = change.RegionalHostname
			continue
		}
		if rh.RegionKey != change.RegionalHostname.RegionKey {
			return nil, fmt.Errorf("conflicting region keys for regional hostname %q: %q and %q", change.RegionalHostname.Hostname, rh.RegionKey, change.RegionalHostname.RegionKey)
		}
	}
	return slices.Collect(maps.Values(rhs)), nil
}

// regionalHostnamesChanges build a list of changes needed to synchronize regional hostnames with desired state.
// Returns slice of changes to apply.
func regionalHostnamesChanges(desired []cloudflare.RegionalHostname, regionalHostnames RegionalHostnamesMap) []RegionalHostnameChange {
	changes := make([]RegionalHostnameChange, 0)
	for _, rh := range desired {
		current, found := regionalHostnames[rh.Hostname]
		if rh.RegionKey == "" {
			if !found {
				continue
			}
			changes = append(changes, RegionalHostnameChange{
				action:           cloudFlareDelete,
				RegionalHostname: rh,
			})
			continue
		}
		if !found {
			changes = append(changes, RegionalHostnameChange{
				action:           cloudFlareCreate,
				RegionalHostname: rh,
			})
			continue
		}
		if rh.RegionKey != current.RegionKey {
			changes = append(changes, RegionalHostnameChange{
				action:           cloudFlareUpdate,
				RegionalHostname: rh,
			})
		}
	}
	return changes
}
