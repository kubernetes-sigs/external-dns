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
	"context"
	"errors"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

const (
	recordTemplate              = "%{record_type}"
	providerSpecificForceUpdate = "txt/force-update"
)

// TXTRegistry implements registry interface with ownership implemented via associated TXT records
type TXTRegistry struct {
	provider provider.Provider
	ownerID  string // refers to the owner id of the current instance
	mapper   nameMapper

	// cache the records in memory and update on an interval instead.
	recordsCache            []*endpoint.Endpoint
	recordsCacheRefreshTime time.Time
	cacheInterval           time.Duration

	// optional string to use to replace the asterisk in wildcard entries - without using this,
	// registry TXT records corresponding to wildcard records will be invalid (and rejected by most providers), due to
	// having a '*' appear (not as the first character) - see https://tools.ietf.org/html/rfc1034#section-4.3.3
	wildcardReplacement string

	managedRecordTypes []string

	// encrypt text records
	txtEncryptEnabled bool
	txtEncryptAESKey  []byte
}

// NewTXTRegistry returns new TXTRegistry object
func NewTXTRegistry(provider provider.Provider, txtPrefix, txtSuffix, ownerID string, cacheInterval time.Duration, txtWildcardReplacement string, managedRecordTypes []string, txtEncryptEnabled bool, txtEncryptAESKey []byte) (*TXTRegistry, error) {
	if ownerID == "" {
		return nil, errors.New("owner id cannot be empty")
	}
	if len(txtEncryptAESKey) == 0 {
		txtEncryptAESKey = nil
	} else if len(txtEncryptAESKey) != 32 {
		return nil, errors.New("the AES Encryption key must have a length of 32 bytes")
	}
	if txtEncryptEnabled && txtEncryptAESKey == nil {
		return nil, errors.New("the AES Encryption key must be set when TXT record encryption is enabled")
	}

	if len(txtPrefix) > 0 && len(txtSuffix) > 0 {
		return nil, errors.New("txt-prefix and txt-suffix are mutual exclusive")
	}

	mapper := newaffixNameMapper(txtPrefix, txtSuffix, txtWildcardReplacement)

	return &TXTRegistry{
		provider:            provider,
		ownerID:             ownerID,
		mapper:              mapper,
		cacheInterval:       cacheInterval,
		wildcardReplacement: txtWildcardReplacement,
		managedRecordTypes:  managedRecordTypes,
		txtEncryptEnabled:   txtEncryptEnabled,
		txtEncryptAESKey:    txtEncryptAESKey,
	}, nil
}

func getSupportedTypes() []string {
	return []string{endpoint.RecordTypeA, endpoint.RecordTypeAAAA, endpoint.RecordTypeCNAME, endpoint.RecordTypeNS}
}

func (im *TXTRegistry) GetDomainFilter() endpoint.DomainFilter {
	return im.provider.GetDomainFilter()
}

// Records returns the current records from the registry excluding TXT Records
// If TXT records was created previously to indicate ownership its corresponding value
// will be added to the endpoints Labels map
func (im *TXTRegistry) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
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
		labels, err := endpoint.NewLabelsFromString(record.Targets[0], im.txtEncryptAESKey)
		if err == endpoint.ErrInvalidHeritage {
			// if no heritage is found or it is invalid
			// case when value of txt record cannot be identified
			// record will not be removed as it will have empty owner
			endpoints = append(endpoints, record)
			continue
		}
		if err != nil {
			return nil, err
		}

		endpointName, recordType := im.mapper.toEndpointName(record.DNSName)
		key := endpoint.EndpointKey{
			DNSName:       endpointName,
			RecordType:    recordType,
			SetIdentifier: record.SetIdentifier,
		}
		labelMap[key] = labels
		txtRecordsMap[record.DNSName] = struct{}{}
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

		// Handle both new and old registry format with the preference for the new one
		labels, labelsExist := labelMap[key]
		if !labelsExist && ep.RecordType != endpoint.RecordTypeAAAA {
			key.RecordType = ""
			labels, labelsExist = labelMap[key]
		}
		if labelsExist {
			for k, v := range labels {
				ep.Labels[k] = v
			}
		}

		// Handle the migration of TXT records created before the new format (introduced in v0.12.0).
		// The migration is done for the TXT records owned by this instance only.
		if len(txtRecordsMap) > 0 && ep.Labels[endpoint.OwnerLabelKey] == im.ownerID {
			if plan.IsManagedRecord(ep.RecordType, im.managedRecordTypes) {
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

// generateTXTRecord generates both "old" and "new" TXT records.
// Once we decide to drop old format we need to drop toTXTName() and rename toNewTXTName
func (im *TXTRegistry) generateTXTRecord(r *endpoint.Endpoint) []*endpoint.Endpoint {
	// Missing TXT records are added to the set of changes.
	// Obviously, we don't need any other TXT record for them.
	if r.RecordType == endpoint.RecordTypeTXT {
		return nil
	}

	endpoints := make([]*endpoint.Endpoint, 0)

	if !im.mapper.recordTypeInAffix() && r.RecordType != endpoint.RecordTypeAAAA {
		// old TXT record format
		txt := endpoint.NewEndpoint(im.mapper.toTXTName(r.DNSName), endpoint.RecordTypeTXT, r.Labels.Serialize(true, im.txtEncryptEnabled, im.txtEncryptAESKey))
		if txt != nil {
			txt.WithSetIdentifier(r.SetIdentifier)
			txt.Labels[endpoint.OwnedRecordLabelKey] = r.DNSName
			txt.ProviderSpecific = r.ProviderSpecific
			endpoints = append(endpoints, txt)
		}
	}
	// new TXT record format (containing record type)
	txtNew := endpoint.NewEndpoint(im.mapper.toNewTXTName(r.DNSName, r.RecordType), endpoint.RecordTypeTXT, r.Labels.Serialize(true, im.txtEncryptEnabled, im.txtEncryptAESKey))
	if txtNew != nil {
		txtNew.WithSetIdentifier(r.SetIdentifier)
		txtNew.Labels[endpoint.OwnedRecordLabelKey] = r.DNSName
		txtNew.ProviderSpecific = r.ProviderSpecific
		endpoints = append(endpoints, txtNew)
	}

	return endpoints
}

// ApplyChanges updates dns provider with the changes
// for each created/deleted record it will also take into account TXT records for creation/deletion
func (im *TXTRegistry) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
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

		filteredChanges.Create = append(filteredChanges.Create, im.generateTXTRecord(r)...)

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
func (im *TXTRegistry) AdjustEndpoints(endpoints []*endpoint.Endpoint) []*endpoint.Endpoint {
	return im.provider.AdjustEndpoints(endpoints)
}

/**
  nameMapper is the interface for mapping between the endpoint for the source
  and the endpoint for the TXT record.
*/

type nameMapper interface {
	toEndpointName(string) (endpointName string, recordType string)
	toTXTName(string) string
	toNewTXTName(string, string) string
	recordTypeInAffix() bool
}

type affixNameMapper struct {
	prefix              string
	suffix              string
	wildcardReplacement string
}

var _ nameMapper = affixNameMapper{}

func newaffixNameMapper(prefix, suffix, wildcardReplacement string) affixNameMapper {
	return affixNameMapper{prefix: strings.ToLower(prefix), suffix: strings.ToLower(suffix), wildcardReplacement: strings.ToLower(wildcardReplacement)}
}

// extractRecordTypeDefaultPosition extracts record type from the default position
// when not using '%{record_type}' in the prefix/suffix
func extractRecordTypeDefaultPosition(name string) (baseName, recordType string) {
	nameS := strings.Split(name, "-")
	for _, t := range getSupportedTypes() {
		if nameS[0] == strings.ToLower(t) {
			return strings.TrimPrefix(name, nameS[0]+"-"), t
		}
	}
	return name, ""
}

// dropAffixExtractType strips TXT record to find an endpoint name it manages
// it also returns the record type
func (pr affixNameMapper) dropAffixExtractType(name string) (baseName, recordType string) {
	prefix := pr.prefix
	suffix := pr.suffix

	if pr.recordTypeInAffix() {
		for _, t := range getSupportedTypes() {
			t = strings.ToLower(t)
			iPrefix := strings.ReplaceAll(prefix, recordTemplate, t)
			iSuffix := strings.ReplaceAll(suffix, recordTemplate, t)

			if pr.isPrefix() && strings.HasPrefix(name, iPrefix) {
				return strings.TrimPrefix(name, iPrefix), t
			}

			if pr.isSuffix() && strings.HasSuffix(name, iSuffix) {
				return strings.TrimSuffix(name, iSuffix), t
			}
		}

		// handle old TXT records
		prefix = pr.dropAffixTemplate(prefix)
		suffix = pr.dropAffixTemplate(suffix)
	}

	if pr.isPrefix() && strings.HasPrefix(name, prefix) {
		return extractRecordTypeDefaultPosition(strings.TrimPrefix(name, prefix))
	}

	if pr.isSuffix() && strings.HasSuffix(name, suffix) {
		return extractRecordTypeDefaultPosition(strings.TrimSuffix(name, suffix))
	}

	return "", ""
}

func (pr affixNameMapper) dropAffixTemplate(name string) string {
	return strings.ReplaceAll(name, recordTemplate, "")
}

func (pr affixNameMapper) isPrefix() bool {
	return len(pr.suffix) == 0
}

func (pr affixNameMapper) isSuffix() bool {
	return len(pr.prefix) == 0 && len(pr.suffix) > 0
}

func (pr affixNameMapper) toEndpointName(txtDNSName string) (endpointName string, recordType string) {
	lowerDNSName := strings.ToLower(txtDNSName)

	// drop prefix
	if pr.isPrefix() {
		return pr.dropAffixExtractType(lowerDNSName)
	}

	// drop suffix
	if pr.isSuffix() {
		dc := strings.Count(pr.suffix, ".")
		DNSName := strings.SplitN(lowerDNSName, ".", 2+dc)
		domainWithSuffix := strings.Join(DNSName[:1+dc], ".")

		r, rType := pr.dropAffixExtractType(domainWithSuffix)
		return r + "." + DNSName[1+dc], rType
	}
	return "", ""
}

func (pr affixNameMapper) toTXTName(endpointDNSName string) string {
	DNSName := strings.SplitN(endpointDNSName, ".", 2)

	prefix := pr.dropAffixTemplate(pr.prefix)
	suffix := pr.dropAffixTemplate(pr.suffix)
	// If specified, replace a leading asterisk in the generated txt record name with some other string
	if pr.wildcardReplacement != "" && DNSName[0] == "*" {
		DNSName[0] = pr.wildcardReplacement
	}

	if len(DNSName) < 2 {
		return prefix + DNSName[0] + suffix
	}
	return prefix + DNSName[0] + suffix + "." + DNSName[1]
}

func (pr affixNameMapper) recordTypeInAffix() bool {
	if strings.Contains(pr.prefix, recordTemplate) {
		return true
	}
	if strings.Contains(pr.suffix, recordTemplate) {
		return true
	}
	return false
}

func (pr affixNameMapper) normalizeAffixTemplate(afix, recordType string) string {
	if strings.Contains(afix, recordTemplate) {
		return strings.ReplaceAll(afix, recordTemplate, recordType)
	}
	return afix
}

func (pr affixNameMapper) toNewTXTName(endpointDNSName, recordType string) string {
	DNSName := strings.SplitN(endpointDNSName, ".", 2)
	recordType = strings.ToLower(recordType)
	recordT := recordType + "-"

	prefix := pr.normalizeAffixTemplate(pr.prefix, recordType)
	suffix := pr.normalizeAffixTemplate(pr.suffix, recordType)

	// If specified, replace a leading asterisk in the generated txt record name with some other string
	if pr.wildcardReplacement != "" && DNSName[0] == "*" {
		DNSName[0] = pr.wildcardReplacement
	}

	if !pr.recordTypeInAffix() {
		DNSName[0] = recordT + DNSName[0]
	}

	if len(DNSName) < 2 {
		return prefix + DNSName[0] + suffix
	}

	return prefix + DNSName[0] + suffix + "." + DNSName[1]
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
