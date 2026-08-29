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

package mapper

import (
	"sort"
	"strings"

	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
)

const (
	recordTemplate = "%{record_type}"
)

var (
	supportedRecords = []string{
		endpoint.RecordTypeA,
		endpoint.RecordTypeAAAA,
		endpoint.RecordTypeCNAME,
		endpoint.RecordTypeNS,
		endpoint.RecordTypeMX,
		endpoint.RecordTypePTR,
		endpoint.RecordTypeSRV,
		endpoint.RecordTypeNAPTR,
		endpoint.RecordTypeTXT,
		endpoint.RecordTypeDNAME,
	}
)

// NameMapper is the interface for mapping between the endpoint for the source
// and the endpoint for the TXT record.
type NameMapper interface {
	ToEndpointName(string) (string, string)
	ToTXTName(string, string) string
}

// AffixNameMapper is a name mapper based on prefix/suffix affixes.
type AffixNameMapper struct {
	prefix              string
	suffix              string
	wildcardReplacement string
	// zones, longest first, split hostname from domain for names that
	// contain dots in the host label group (e.g. name-192.168.0.1.example.com).
	zones []string
}

// NewAffixNameMapper returns a new AffixNameMapper.
func NewAffixNameMapper(prefix, suffix, wildcardReplacement string) AffixNameMapper {
	return NewAffixNameMapperWithZones(prefix, suffix, wildcardReplacement, nil)
}

// NewAffixNameMapperWithZones returns a mapper that splits TXT names at a
// known zone boundary instead of the first dot. An empty zone list keeps the
// legacy first-dot split so existing records stay unchanged.
func NewAffixNameMapperWithZones(prefix, suffix, wildcardReplacement string, zones []string) AffixNameMapper {
	normalized := make([]string, 0, len(zones))
	for _, zone := range zones {
		z := strings.ToLower(strings.Trim(zone, "."))
		if z != "" {
			normalized = append(normalized, z)
		}
	}
	sort.Slice(normalized, func(i, j int) bool {
		return len(normalized[i]) > len(normalized[j])
	})
	return AffixNameMapper{
		prefix:              strings.ToLower(prefix),
		suffix:              strings.ToLower(suffix),
		wildcardReplacement: strings.ToLower(wildcardReplacement),
		zones:               normalized,
	}
}

// ZonesFromDomainFilter returns include filters for zone-aware TXT mapping.
func ZonesFromDomainFilter(df endpoint.DomainFilterInterface) []string {
	d, ok := df.(*endpoint.DomainFilter)
	if !ok || d == nil {
		return nil
	}
	return d.Filters
}

func (a AffixNameMapper) findZone(dns string) string {
	dns = strings.ToLower(strings.TrimSuffix(dns, "."))
	for _, zone := range a.zones {
		if strings.HasSuffix(dns, "."+zone) {
			return zone
		}
	}
	return ""
}

func (a AffixNameMapper) splitName(dns string) (hostname, domain string) {
	dns = strings.TrimSuffix(dns, ".")
	if zone := a.findZone(dns); zone != "" {
		return strings.TrimSuffix(dns, "."+zone), zone
	}
	parts := strings.SplitN(dns, ".", 2)
	hostname = parts[0]
	if len(parts) > 1 {
		domain = parts[1]
	}
	return hostname, domain
}

func (a AffixNameMapper) ToEndpointName(dns string) (string, string) {
	lowerDNSName := strings.ToLower(dns)

	// drop prefix
	if a.isPrefix() {
		return a.dropAffixExtractType(lowerDNSName)
	}

	// drop suffix
	if a.isSuffix() {
		if zone := a.findZone(lowerDNSName); zone != "" {
			hostPart := strings.TrimSuffix(lowerDNSName, "."+zone)
			r, rType := a.dropAffixExtractType(hostPart)
			if r == "" && rType == "" {
				return "", ""
			}
			return r + "." + zone, rType
		}
		dc := strings.Count(a.suffix, ".")
		parts := strings.SplitN(lowerDNSName, ".", 2+dc)
		if len(parts) <= dc {
			log.Debugf("skipping TXT record %q: too few labels for suffix %q", dns, a.suffix)
			return "", ""
		}
		r, rType := a.dropAffixExtractType(strings.Join(parts[:1+dc], "."))
		if len(parts) <= 1+dc {
			return r, rType
		}
		return r + "." + parts[1+dc], rType
	}
	return "", ""
}

func (a AffixNameMapper) ToTXTName(dns, recordType string) string {
	hostname, domain := a.splitName(dns)
	recordType = strings.ToLower(recordType)
	recordT := recordType + "-"

	prefix := a.normalizeAffixTemplate(a.prefix, recordType)
	suffix := a.normalizeAffixTemplate(a.suffix, recordType)

	// If specified, replace a leading asterisk in the generated txt record name with some other string
	if a.wildcardReplacement != "" && hostname == "*" {
		hostname = a.wildcardReplacement
	}

	if !a.recordTypeInAffix() {
		hostname = recordT + hostname
	}

	if domain == "" {
		return prefix + hostname + suffix
	}

	return prefix + hostname + suffix + "." + domain
}

func (a AffixNameMapper) recordTypeInAffix() bool {
	if strings.Contains(a.prefix, recordTemplate) {
		return true
	}
	if strings.Contains(a.suffix, recordTemplate) {
		return true
	}
	return false
}

func (a AffixNameMapper) normalizeAffixTemplate(afix, recordType string) string {
	if strings.Contains(afix, recordTemplate) {
		return strings.ReplaceAll(afix, recordTemplate, recordType)
	}
	return afix
}

func (a AffixNameMapper) isPrefix() bool {
	return len(a.suffix) == 0
}

func (a AffixNameMapper) isSuffix() bool {
	return len(a.prefix) == 0 && len(a.suffix) > 0
}

func (a AffixNameMapper) dropAffixTemplate(name string) string {
	return strings.ReplaceAll(name, recordTemplate, "")
}

// dropAffixExtractType strips TXT record to find an endpoint name it manages.
// It also returns the record type.
func (a AffixNameMapper) dropAffixExtractType(name string) (string, string) {
	prefix := a.prefix
	suffix := a.suffix

	if a.recordTypeInAffix() {
		for _, t := range supportedRecords {
			tLower := strings.ToLower(t)
			iPrefix := strings.ReplaceAll(prefix, recordTemplate, tLower)
			iSuffix := strings.ReplaceAll(suffix, recordTemplate, tLower)

			if a.isPrefix() && strings.HasPrefix(name, iPrefix) {
				return strings.TrimPrefix(name, iPrefix), t
			}

			if a.isSuffix() && strings.HasSuffix(name, iSuffix) {
				return strings.TrimSuffix(name, iSuffix), t
			}
		}

		// handle old TXT records
		prefix = a.dropAffixTemplate(prefix)
		suffix = a.dropAffixTemplate(suffix)
	}

	if a.isPrefix() && strings.HasPrefix(name, prefix) {
		return extractRecordTypeDefaultPosition(strings.TrimPrefix(name, prefix))
	}

	if a.isSuffix() && strings.HasSuffix(name, suffix) {
		return extractRecordTypeDefaultPosition(strings.TrimSuffix(name, suffix))
	}

	return "", ""
}

// extractRecordTypeDefaultPosition extracts record type from the default position
// when not using '%{record_type}' in the prefix/suffix
func extractRecordTypeDefaultPosition(name string) (string, string) {
	nameS := strings.Split(name, "-")
	for _, t := range supportedRecords {
		if nameS[0] == strings.ToLower(t) {
			return strings.TrimPrefix(name, nameS[0]+"-"), t
		}
	}
	return name, ""
}
