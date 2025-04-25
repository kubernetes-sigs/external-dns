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
	"errors"
	"fmt"
	"maps"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go"
	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/source/annotations"
)

var recordTypeRegionalHostnameSupported = map[string]bool{
	"A":     true,
	"AAAA":  true,
	"CNAME": true,
}

type regionalHostnameChange struct {
	action string
	cloudflare.RegionalHostname
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

// submitDataLocalizationRegionalHostnameChanges applies a set of data localization regional hostname changes, returns false if it fails
func (p *CloudFlareProvider) submitDataLocalizationRegionalHostnameChanges(ctx context.Context, rhChanges []regionalHostnameChange, resourceContainer *cloudflare.ResourceContainer) bool {
	failedChange := false

	for _, rhChange := range rhChanges {
		logFields := log.Fields{
			"hostname":   rhChange.Hostname,
			"region_key": rhChange.RegionKey,
			"action":     rhChange.action,
			"zone":       resourceContainer.Identifier,
		}
		log.WithFields(logFields).Info("Changing regional hostname")
		switch rhChange.action {
		case cloudFlareCreate:
			log.WithFields(logFields).Debug("Creating regional hostname")
			if p.DryRun {
				continue
			}
			regionalHostnameParam := createDataLocalizationRegionalHostnameParams(rhChange)
			err := p.Client.CreateDataLocalizationRegionalHostname(ctx, resourceContainer, regionalHostnameParam)
			if err != nil {
				var apiErr *cloudflare.Error
				if errors.As(err, &apiErr) && apiErr.StatusCode == http.StatusConflict {
					log.WithFields(logFields).Debug("Regional hostname already exists, updating instead")
					params := updateDataLocalizationRegionalHostnameParams(rhChange)
					err := p.Client.UpdateDataLocalizationRegionalHostname(ctx, resourceContainer, params)
					if err != nil {
						failedChange = true
						log.WithFields(logFields).Errorf("failed to update regional hostname: %v", err)
					}
					continue
				}
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
				var apiErr *cloudflare.Error
				if errors.As(err, &apiErr) && apiErr.StatusCode == http.StatusNotFound {
					log.WithFields(logFields).Debug("Regional hostname not does not exists, creating instead")
					params := createDataLocalizationRegionalHostnameParams(rhChange)
					err := p.Client.CreateDataLocalizationRegionalHostname(ctx, resourceContainer, params)
					if err != nil {
						failedChange = true
						log.WithFields(logFields).Errorf("failed to create regional hostname: %v", err)
					}
					continue
				}
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
				var apiErr *cloudflare.Error
				if errors.As(err, &apiErr) && apiErr.StatusCode == http.StatusNotFound {
					log.WithFields(logFields).Debug("Regional hostname does not exists, nothing to do")
					continue
				}
				failedChange = true
				log.WithFields(logFields).Errorf("failed to delete regional hostname: %v", err)
			}
		}
	}

	return !failedChange
}

func (p *CloudFlareProvider) regionalHostname(ep *endpoint.Endpoint) cloudflare.RegionalHostname {
	if p.RegionKey == "" || !recordTypeRegionalHostnameSupported[ep.RecordType] {
		return cloudflare.RegionalHostname{}
	}
	regionKey := p.RegionKey
	if epRegionKey, exists := ep.GetProviderSpecificProperty(annotations.CloudflareRegionKey); exists {
		regionKey = epRegionKey
	}
	return cloudflare.RegionalHostname{
		Hostname:  ep.DNSName,
		RegionKey: regionKey,
	}
}

// dataLocalizationRegionalHostnamesChanges processes a slice of cloudFlare changes and consolidates them
// into a list of data localization regional hostname changes.
// returns nil if no changes are needed
func dataLocalizationRegionalHostnamesChanges(changes []*cloudFlareChange) ([]regionalHostnameChange, error) {
	regionalHostnameChanges := make(map[string]regionalHostnameChange)
	for _, change := range changes {
		if change.RegionalHostname.Hostname == "" {
			continue
		}
		if change.RegionalHostname.RegionKey == "" {
			return nil, fmt.Errorf("region key is empty for regional hostname %q", change.RegionalHostname.Hostname)
		}
		regionalHostname, ok := regionalHostnameChanges[change.RegionalHostname.Hostname]
		switch change.Action {
		case cloudFlareCreate, cloudFlareUpdate:
			if !ok {
				regionalHostnameChanges[change.RegionalHostname.Hostname] = regionalHostnameChange{
					action:           change.Action,
					RegionalHostname: change.RegionalHostname,
				}
				continue
			}
			if regionalHostname.RegionKey != change.RegionalHostname.RegionKey {
				return nil, fmt.Errorf("conflicting region keys for regional hostname %q: %q and %q", change.RegionalHostname.Hostname, regionalHostname.RegionKey, change.RegionalHostname.RegionKey)
			}
			if (change.Action == cloudFlareUpdate && regionalHostname.action != cloudFlareUpdate) ||
				regionalHostname.action == cloudFlareDelete {
				regionalHostnameChanges[change.RegionalHostname.Hostname] = regionalHostnameChange{
					action:           cloudFlareUpdate,
					RegionalHostname: change.RegionalHostname,
				}
			}
		case cloudFlareDelete:
			if !ok {
				regionalHostnameChanges[change.RegionalHostname.Hostname] = regionalHostnameChange{
					action:           cloudFlareDelete,
					RegionalHostname: change.RegionalHostname,
				}
				continue
			}
		}
	}
	return slices.Collect(maps.Values(regionalHostnameChanges)), nil
}
