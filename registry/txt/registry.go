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

package txt

import (
	"context"
	"errors"
	"maps"

	"strings"
	"time"

	b64 "encoding/base64"

	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/registry/mapper"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

const (
	providerSpecificForceUpdate = "txt/force-update"
)

// TXTRegistry implements registry interface with ownership implemented via associated TXT records
type TXTRegistry struct {
	provider provider.Provider
	ownerID  string // refers to the owner id of the current instance
	mapper   mapper.NameMapper

	// cache the records in memory and update on an interval instead.
	recordsCache            []*endpoint.Endpoint
	recordsCacheRefreshTime time.Time
	cacheInterval           time.Duration

	// optional string to use to replace the asterisk in wildcard entries - without using this,
	// registry TXT records corresponding to wildcard records will be invalid (and rejected by most providers), due to
	// having a '*' appear (not as the first character) - see https://tools.ietf.org/html/rfc1034#section-4.3.3
	wildcardReplacement string

	managedRecordTypes []string
	excludeRecordTypes []string

	// encrypt text records
	txtEncryptEnabled bool
	txtEncryptAESKey  []byte

	// Handle Owner ID migration
	oldOwnerID string

	// existingTXTs is the TXT records that already exist in the zone so that
	// ApplyChanges() can skip re-creating them. See the struct below for details.
	existingTXTs *existingTXTs
}

// existingTXTs stores pre‑existing TXT records to avoid duplicate creation.
// It relies on the fact that Records() is always called **before** ApplyChanges()
// within a single reconciliation cycle.
type existingTXTs struct {
	entries map[recordKey]struct{}
}

type recordKey struct {
	dnsName       string
	setIdentifier string
}

func newExistingTXTs() *existingTXTs {
	return &existingTXTs{
		entries: make(map[recordKey]struct{}),
	}
}

func (im *existingTXTs) add(r *endpoint.Endpoint) {
	key := recordKey{
		dnsName:       r.DNSName,
		setIdentifier: r.SetIdentifier,
	}
	im.entries[key] = struct{}{}
}

// isAbsent returns true when there is no entry for the given name in the store.
// This is intended for the "if absent -> create" pattern.
func (im *existingTXTs) isAbsent(ep *endpoint.Endpoint) bool {
	key := recordKey{
		dnsName:       ep.DNSName,
		setIdentifier: ep.SetIdentifier,
	}
	_, ok := im.entries[key]
	return !ok
}

func (im *existingTXTs) reset() {
	// Reset the existing TXT records for the next reconciliation loop.
	// This is necessary because the existing TXT records are only relevant for the current reconciliation cycle.
	im.entries = make(map[recordKey]struct{})
}

// NewTXTRegistry returns a new TXTRegistry object. When newFormatOnly is true, it will only
// generate new format TXT records, otherwise it generates both old and new formats for
// backwards compatibility.
func NewTXTRegistry(provider provider.Provider, txtPrefix, txtSuffix, ownerID string,
	cacheInterval time.Duration, txtWildcardReplacement string,
	managedRecordTypes, excludeRecordTypes []string,
	txtEncryptEnabled bool, txtEncryptAESKey []byte,
	oldOwnerID string) (*TXTRegistry, error) {
	if ownerID == "" {
		return nil, errors.New("owner id cannot be empty")
	}

	// TODO: encryption logic duplicated in DynamoDB registry; refactor into common utility function.
	if len(txtEncryptAESKey) == 0 {
		txtEncryptAESKey = nil
	} else if len(txtEncryptAESKey) != 32 {
		var err error
		if txtEncryptAESKey, err = b64.StdEncoding.DecodeString(string(txtEncryptAESKey)); err != nil || len(txtEncryptAESKey) != 32 {
			return nil, errors.New("the AES Encryption key must be 32 bytes long, in either plain text or base64-encoded format")
		}
	}

	if txtEncryptEnabled && txtEncryptAESKey == nil {
		return nil, errors.New("the AES Encryption key must be set when TXT record encryption is enabled")
	}

	if len(txtPrefix) > 0 && len(txtSuffix) > 0 {
		return nil, errors.New("txt-prefix and txt-suffix are mutual exclusive")
	}

	return &TXTRegistry{
		provider:            provider,
		ownerID:             ownerID,
		mapper:              mapper.NewAffixNameMapper(txtPrefix, txtSuffix, txtWildcardReplacement),
		cacheInterval:       cacheInterval,
		wildcardReplacement: txtWildcardReplacement,
		managedRecordTypes:  managedRecordTypes,
		excludeRecordTypes:  excludeRecordTypes,
		txtEncryptEnabled:   txtEncryptEnabled,
		txtEncryptAESKey:    txtEncryptAESKey,
		oldOwnerID:          oldOwnerID,
		existingTXTs:        newExistingTXTs(),
	}, nil
}

func (im *TXTRegistry) GetDomainFilter() endpoint.DomainFilterInterface {
	return im.provider.GetDomainFilter()
}

func (im *TXTRegistry) OwnerID() string {
	return im.ownerID
}

// Records returns the current records from the registry excluding TXT Records
// If TXT records was created previously to indicate ownership its corresponding value
// will be added to the endpoints Labels map
func (im *TXTRegistry) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	// existingTXTs must always hold the latest TXT records, so it needs to be reset every time.
	// Previously, it was reset with a defer after ApplyChanges, but ApplyChanges is not called
	// when plan.HasChanges() is false (i.e., when there are no changes to apply).
	// In that case, stale TXT record information could remain, so we reset it here instead.
	im.existingTXTs.reset()

	// If we have the zones cached AND we have refreshed the cache since the
	// last given interval, then just use the cached results.
	if im.recordsCache != nil && time.Since(im.recordsCacheRefreshTime) < im.cacheInterval {
		log.Debug("Using cached records.")
		return im.recordsCache, nil
	}

	records, err := im.provider.Records(ctx)
	if err != nil {
		return nil, err
	}

	endpoints := []*endpoint.Endpoint{}

	labelMap := map[endpoint.EndpointKey]endpoint.Labels{}
	txtRecordsMap := map[string]struct{}{}

	for _, record := range records {
		if record.RecordType != endpoint.RecordTypeTXT {
			endpoints = append(endpoints, record)
			continue
		}
		// We simply assume that TXT records for the registry will always have only one target.
		// If there are no targets (e.g for routing policy based records in google), direct targets will be empty
		if len(record.Targets) == 0 {
			log.Errorf("TXT record has no targets %s", record.DNSName)
			continue
		}
		labels, err := endpoint.NewLabelsFromString(record.Targets[0], im.txtEncryptAESKey)
		if errors.Is(err, endpoint.ErrInvalidHeritage) {
			// if no heritage is found or it is invalid
			// case when value of txt record cannot be identified
			// record will not be removed as it will have empty owner
			endpoints = append(endpoints, record)
			continue
		}
		if err != nil {
			return nil, err
		}

		endpointName, recordType := im.mapper.ToEndpointName(record.DNSName)
		key := endpoint.EndpointKey{
			DNSName:       endpointName,
			RecordType:    recordType,
			SetIdentifier: record.SetIdentifier,
		}
		labelMap[key] = labels
		txtRecordsMap[record.DNSName] = struct{}{}
		im.existingTXTs.add(record)
	}

	for _, ep := range endpoints {
		if ep.Labels == nil {
			ep.Labels = endpoint.NewLabels()
		}
		dnsNameSplit := strings.Split(ep.DNSName, ".")
		// If specified, replace a leading asterisk in the generated txt record name with some other string
		if im.wildcardReplacement != "" && dnsNameSplit[0] == "*" {
			dnsNameSplit[0] = im.wildcardReplacement
		}
		dnsName := strings.Join(dnsNameSplit, ".")
		key := endpoint.EndpointKey{
			DNSName:       dnsName,
			RecordType:    ep.RecordType,
			SetIdentifier: ep.SetIdentifier,
		}

		// AWS Alias records have "new" format encoded as type "cname"
		if isAlias, found := ep.GetBoolProviderSpecificProperty("alias"); found && isAlias && ep.RecordType == endpoint.RecordTypeA {
			key.RecordType = endpoint.RecordTypeCNAME
		}

		// Handle both new and old registry format with the preference for the new one
		labels, labelsExist := labelMap[key]
		if !labelsExist && ep.RecordType != endpoint.RecordTypeAAAA {
			key.RecordType = ""
			labels, labelsExist = labelMap[key]
		}
		if labelsExist {
			maps.Copy(ep.Labels, labels)
		}

		if im.oldOwnerID != "" && ep.Labels[endpoint.OwnerLabelKey] == im.oldOwnerID {
			ep.Labels[endpoint.OwnerLabelKey] = im.ownerID
		}

		// TODO: remove this migration logic in some future release
		// Handle the migration of TXT records created before the new format (introduced in v0.12.0).
		// The migration is done for the TXT records owned by this instance only.
		if len(txtRecordsMap) > 0 && ep.Labels[endpoint.OwnerLabelKey] == im.ownerID {
			if plan.IsManagedRecord(ep.RecordType, im.managedRecordTypes, im.excludeRecordTypes) {
				// Get desired TXT records and detect the missing ones
				desiredTXTs := im.generateTXTRecord(ep)
				for _, desiredTXT := range desiredTXTs {
					if _, exists := txtRecordsMap[desiredTXT.DNSName]; !exists {
						ep.WithProviderSpecific(providerSpecificForceUpdate, "true")
					}
				}
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

// generateTXTRecord generates TXT records in either both formats (old and new) or new format only,
// depending on the newFormatOnly configuration. The old format is maintained for backwards
// compatibility but can be disabled to reduce the number of DNS records.
func (im *TXTRegistry) generateTXTRecord(r *endpoint.Endpoint) []*endpoint.Endpoint {
	return im.generateTXTRecordWithFilter(r, func(_ *endpoint.Endpoint) bool { return true })
}

func (im *TXTRegistry) generateTXTRecordWithFilter(r *endpoint.Endpoint, filter func(*endpoint.Endpoint) bool) []*endpoint.Endpoint {
	endpoints := make([]*endpoint.Endpoint, 0)

	// Always create new format record
	recordType := r.RecordType
	// AWS Alias records are encoded as type "cname"
	if isAlias, found := r.GetBoolProviderSpecificProperty("alias"); found && isAlias && recordType == endpoint.RecordTypeA {
		recordType = endpoint.RecordTypeCNAME
	}

	if im.oldOwnerID != "" && r.Labels[endpoint.OwnerLabelKey] == im.oldOwnerID {
		r.Labels[endpoint.OwnerLabelKey] = im.ownerID
	}

	txtNew := endpoint.NewEndpoint(im.mapper.ToTXTName(r.DNSName, recordType), endpoint.RecordTypeTXT, r.Labels.Serialize(true, im.txtEncryptEnabled, im.txtEncryptAESKey))
	if txtNew != nil {
		txtNew.WithSetIdentifier(r.SetIdentifier)
		txtNew.Labels[endpoint.OwnedRecordLabelKey] = r.DNSName
		txtNew.ProviderSpecific = r.ProviderSpecific
		if filter(txtNew) {
			endpoints = append(endpoints, txtNew)
		}
	}
	return endpoints
}

// ApplyChanges updates dns provider with the changes
// for each created/deleted record it will also take into account TXT records for creation/deletion
func (im *TXTRegistry) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	filteredChanges := &plan.Changes{
		Create:    changes.Create,
		UpdateNew: endpoint.FilterEndpointsByOwnerID(im.ownerID, changes.UpdateNew),
		UpdateOld: endpoint.FilterEndpointsByOwnerID(im.ownerID, changes.UpdateOld),
		Delete:    endpoint.FilterEndpointsByOwnerID(im.ownerID, changes.Delete),
	}

	for _, r := range filteredChanges.Create {
		if r.Labels == nil {
			r.Labels = make(map[string]string)
		}
		r.Labels[endpoint.OwnerLabelKey] = im.ownerID

		filteredChanges.Create = append(filteredChanges.Create, im.generateTXTRecordWithFilter(r, im.existingTXTs.isAbsent)...)

		if im.cacheInterval > 0 {
			im.addToCache(r)
		}
	}

	for _, r := range filteredChanges.Delete {
		// when we delete TXT records for which value has changed (due to new label) this would still work because
		// !!! TXT record value is uniquely generated from the Labels of the endpoint. Hence old TXT record can be uniquely reconstructed
		// !!! After migration to the new TXT registry format we can drop records in old format here!!!
		filteredChanges.Delete = append(filteredChanges.Delete, im.generateTXTRecord(r)...)

		if im.cacheInterval > 0 {
			im.removeFromCache(r)
		}
	}

	// make sure TXT records are consistently updated as well
	for _, r := range filteredChanges.UpdateOld {
		// when we updateOld TXT records for which value has changed (due to new label) this would still work because
		// !!! TXT record value is uniquely generated from the Labels of the endpoint. Hence old TXT record can be uniquely reconstructed
		filteredChanges.UpdateOld = append(filteredChanges.UpdateOld, im.generateTXTRecord(r)...)
		// remove old version of record from cache
		if im.cacheInterval > 0 {
			im.removeFromCache(r)
		}
	}

	// make sure TXT records are consistently updated as well
	for _, r := range filteredChanges.UpdateNew {
		filteredChanges.UpdateNew = append(filteredChanges.UpdateNew, im.generateTXTRecord(r)...)
		// add new version of record to cache
		if im.cacheInterval > 0 {
			im.addToCache(r)
		}
	}

	// when caching is enabled, disable the provider from using the cache
	if im.cacheInterval > 0 {
		ctx = context.WithValue(ctx, provider.RecordsContextKey, nil)
	}
	return im.provider.ApplyChanges(ctx, filteredChanges)
}

// AdjustEndpoints modifies the endpoints as needed by the specific provider
func (im *TXTRegistry) AdjustEndpoints(endpoints []*endpoint.Endpoint) ([]*endpoint.Endpoint, error) {
	return im.provider.AdjustEndpoints(endpoints)
}

func (im *TXTRegistry) addToCache(ep *endpoint.Endpoint) {
	if im.recordsCache != nil {
		im.recordsCache = append(im.recordsCache, ep)
	}
}

func (im *TXTRegistry) removeFromCache(ep *endpoint.Endpoint) {
	if im.recordsCache == nil || ep == nil {
		return
	}

	for i, e := range im.recordsCache {
		if e.DNSName == ep.DNSName && e.RecordType == ep.RecordType && e.SetIdentifier == ep.SetIdentifier && e.Targets.Same(ep.Targets) {
			// We found a match delete the endpoint from the cache.
			im.recordsCache = append(im.recordsCache[:i], im.recordsCache[i+1:]...)
			return
		}
	}
}
