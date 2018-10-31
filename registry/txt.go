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

package registry

import (
	"errors"
	"time"

	"strings"

	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"
	"github.com/kubernetes-incubator/external-dns/provider"
	log "github.com/sirupsen/logrus"
)

// TXTRegistry implements registry interface with ownership implemented via associated TXT records
type TXTRegistry struct {
	provider provider.Provider
	ownerID  string //refers to the owner id of the current instance
	mapper   nameMapper

	// cache the records in memory and update on an interval instead.
	recordsCache            []*endpoint.Endpoint
	recordsCacheRefreshTime time.Time
	cacheInterval           time.Duration
}

// NewTXTRegistry returns new TXTRegistry object
func NewTXTRegistry(provider provider.Provider, txtPrefix, ownerID string, cacheInterval time.Duration) (*TXTRegistry, error) {
	if ownerID == "" {
		return nil, errors.New("owner id cannot be empty")
	}

	mapper := newPrefixNameMapper(txtPrefix)

	return &TXTRegistry{
		provider:      provider,
		ownerID:       ownerID,
		mapper:        mapper,
		cacheInterval: cacheInterval,
	}, nil
}

type labelEntry struct {
	name   string
	labels endpoint.Labels
}

type labelList []labelEntry

func (l labelList) find(e *endpoint.Endpoint) endpoint.Labels {
	for _, entry := range l {
		if e.DNSName != entry.name {
			continue
		}
		if !e.HasNonUniqueRecords() {
			return entry.labels
		}
		target, ok := entry.labels[endpoint.TargetLabel]
		if ok && target == e.Targets[0] {
			return entry.labels
		}
	}
	return nil
}

// Records returns the current records from the registry excluding TXT Records
// If TXT records was created previously to indicate ownership its corresponding value
// will be added to the endpoints Labels map
func (im *TXTRegistry) Records() ([]*endpoint.Endpoint, error) {
	// If we have the zones cached AND we have refreshed the cache since the
	// last given interval, then just use the cached results.
	if im.recordsCache != nil && time.Since(im.recordsCacheRefreshTime) < im.cacheInterval {
		log.Debug("Using cached records.")
		return im.recordsCache, nil
	}

	records, err := im.provider.Records()
	if err != nil {
		return nil, err
	}

	endpoints := []*endpoint.Endpoint{}

	ll := labelList{}

	for _, record := range records {
		if record.RecordType != endpoint.RecordTypeTXT {
			endpoints = append(endpoints, record)
			continue
		}
		// We simply assume that TXT records for the registry will always have only one target.
		labels, err := endpoint.NewLabelsFromString(record.Targets[0])
		if err == endpoint.ErrInvalidHeritage {
			//if no heritage is found or it is invalid
			//case when value of txt record cannot be identified
			//record will not be removed as it will have empty owner
			endpoints = append(endpoints, record)
			continue
		}
		if err != nil {
			return nil, err
		}
		endpointDNSName := im.mapper.toEndpointName(record.DNSName)
		ll = append(ll, labelEntry{name: endpointDNSName, labels: labels})
	}

	for _, ep := range endpoints {
		ep.Labels = endpoint.NewLabels()
		if labels := ll.find(ep); labels != nil {
			for k, v := range labels {
				ep.Labels[k] = v
			}
		}
	}

	// Update the cache.
	if im.cacheInterval > 0 {
		im.recordsCache = endpoints
		im.recordsCacheRefreshTime = time.Now()
	}

	return endpoints, nil
}

// createEndpoint creates a TXT record endpoint for a resource
func (im *TXTRegistry) createEndpoint(e *endpoint.Endpoint) *endpoint.Endpoint {
	// Non uniqueu DNS names e.g. SRV records, need extra context to be added to the TXT
	// in order to match based on target as the name make be specified multiple times.
	if e.HasNonUniqueRecords() {
		e.Labels[endpoint.TargetLabel] = e.Targets[0]
	}
	return endpoint.NewEndpoint(im.mapper.toTXTName(e.DNSName), endpoint.RecordTypeTXT, e.Labels.Serialize(true))
}

// ApplyChanges updates dns provider with the changes
// for each created/deleted record it will also take into account TXT records for creation/deletion
func (im *TXTRegistry) ApplyChanges(changes *plan.Changes) error {
	filteredChanges := &plan.Changes{
		Create:    changes.Create,
		UpdateNew: filterOwnedRecords(im.ownerID, changes.UpdateNew),
		UpdateOld: filterOwnedRecords(im.ownerID, changes.UpdateOld),
		Delete:    filterOwnedRecords(im.ownerID, changes.Delete),
	}
	for _, r := range filteredChanges.Create {
		if r.Labels == nil {
			r.Labels = make(map[string]string)
		}
		r.Labels[endpoint.OwnerLabelKey] = im.ownerID
		txt := im.createEndpoint(r)
		filteredChanges.Create = append(filteredChanges.Create, txt)

		if im.cacheInterval > 0 {
			im.addToCache(r)
		}
	}

	for _, r := range filteredChanges.Delete {
		txt := im.createEndpoint(r)

		// when we delete TXT records for which value has changed (due to new label) this would still work because
		// !!! TXT record value is uniquely generated from the Labels of the endpoint. Hence old TXT record can be uniquely reconstructed
		filteredChanges.Delete = append(filteredChanges.Delete, txt)

		if im.cacheInterval > 0 {
			im.removeFromCache(r)
		}
	}

	// make sure TXT records are consistently updated as well
	for _, r := range filteredChanges.UpdateOld {
		txt := im.createEndpoint(r)
		// when we updateOld TXT records for which value has changed (due to new label) this would still work because
		// !!! TXT record value is uniquely generated from the Labels of the endpoint. Hence old TXT record can be uniquely reconstructed
		filteredChanges.UpdateOld = append(filteredChanges.UpdateOld, txt)
		// remove old version of record from cache
		if im.cacheInterval > 0 {
			im.removeFromCache(r)
		}
	}

	// make sure TXT records are consistently updated as well
	for _, r := range filteredChanges.UpdateNew {
		txt := im.createEndpoint(r)
		filteredChanges.UpdateNew = append(filteredChanges.UpdateNew, txt)
		// add new version of record to cache
		if im.cacheInterval > 0 {
			im.addToCache(r)
		}
	}

	return im.provider.ApplyChanges(filteredChanges)
}

/**
  TXT registry specific private methods
*/

/**
  nameMapper defines interface which maps the dns name defined for the source
  to the dns name which TXT record will be created with
*/

type nameMapper interface {
	toEndpointName(string) string
	toTXTName(string) string
}

type prefixNameMapper struct {
	prefix string
}

var _ nameMapper = prefixNameMapper{}

func newPrefixNameMapper(prefix string) prefixNameMapper {
	return prefixNameMapper{prefix: prefix}
}

func (pr prefixNameMapper) toEndpointName(txtDNSName string) string {
	if strings.HasPrefix(txtDNSName, pr.prefix) {
		return strings.TrimPrefix(txtDNSName, pr.prefix)
	}
	return ""
}

func (pr prefixNameMapper) toTXTName(endpointDNSName string) string {
	return pr.prefix + endpointDNSName
}

func (im *TXTRegistry) addToCache(ep *endpoint.Endpoint) {
	if im.recordsCache != nil {
		im.recordsCache = append(im.recordsCache, ep)
	}
}

func (im *TXTRegistry) removeFromCache(ep *endpoint.Endpoint) {
	if im.recordsCache == nil || ep == nil {
		// return early.
		return
	}

	for i, e := range im.recordsCache {
		if e.DNSName == ep.DNSName && e.RecordType == ep.RecordType && e.Targets.Same(ep.Targets) {
			// We found a match delete the endpoint from the cache.
			im.recordsCache = append(im.recordsCache[:i], im.recordsCache[i+1:]...)
			return
		}
	}
}
