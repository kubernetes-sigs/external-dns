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
	"strings"

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
		endpoint.RecordTypeSRV,
		endpoint.RecordTypeNAPTR,
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
}

// NewAffixNameMapper returns a new AffixNameMapper.
func NewAffixNameMapper(prefix, suffix, wildcardReplacement string) AffixNameMapper {
	return AffixNameMapper{
		prefix:              strings.ToLower(prefix),
		suffix:              strings.ToLower(suffix),
		wildcardReplacement: strings.ToLower(wildcardReplacement),
	}
}

func (a AffixNameMapper) ToEndpointName(dns string) (string, string) {
	lowerDNSName := strings.ToLower(dns)

	// drop prefix
	if a.isPrefix() {
		return a.dropAffixExtractType(lowerDNSName)
	}

	// drop suffix
	if a.isSuffix() {
		dc := strings.Count(a.suffix, ".")
		DNSName := strings.SplitN(lowerDNSName, ".", 2+dc)
		domainWithSuffix := strings.Join(DNSName[:1+dc], ".")

		r, rType := a.dropAffixExtractType(domainWithSuffix)
		if !strings.Contains(lowerDNSName, ".") {
			return r, rType
		}
		return r + "." + DNSName[1+dc], rType
	}
	return "", ""
}

func (a AffixNameMapper) ToTXTName(dns, recordType string) string {
	DNSName := strings.SplitN(dns, ".", 2)
	recordType = strings.ToLower(recordType)
	recordT := recordType + "-"

	prefix := a.normalizeAffixTemplate(a.prefix, recordType)
	suffix := a.normalizeAffixTemplate(a.suffix, recordType)

	// If specified, replace a leading asterisk in the generated txt record name with some other string
	if a.wildcardReplacement != "" && DNSName[0] == "*" {
		DNSName[0] = a.wildcardReplacement
	}

	if !a.recordTypeInAffix() {
		DNSName[0] = recordT + DNSName[0]
	}

	if len(DNSName) < 2 {
		return prefix + DNSName[0] + suffix
	}

	return prefix + DNSName[0] + suffix + "." + DNSName[1]
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
