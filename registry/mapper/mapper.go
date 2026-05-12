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

	// maxDNSLabelLen is the maximum length of a single DNS label per RFC 1035.
	maxDNSLabelLen = 63
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
		// In the separate-label fallback form, the suffix-bearing portion
		// consists of only the record-type token; the parent's first label
		// sits in the next position. Avoid a stray leading dot in that case.
		if r == "" {
			return DNSName[1+dc], rType
		}
		return r + "." + DNSName[1+dc], rType
	}
	return "", ""
}

// ToTXTName projects a parent record name onto the TXT registry name.
//
// The default "inline" form prepends "<recordType>-" to the parent's first label:
//
//	foo.example.com (CNAME) -> cname-foo.example.com
//
// When the inline form would push any label past RFC 1035's 63-char limit, the
// mapper falls back to the "separate-label" form, where the record-type token
// occupies its own DNS label and the parent's first label is preserved unmodified:
//
//	<60-char-label>.example.com (CNAME) -> cname.<60-char-label>.example.com
//
// Both forms are recognised by ToEndpointName, so this is transparent to readers
// and backward-compatible with existing TXT records.
//
// The fallback applies only when neither the prefix nor the suffix already
// templates the record type — templated affixes are left untouched to preserve
// the user's explicitly-configured layout.
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

	rest := ""
	if len(DNSName) >= 2 {
		rest = "." + DNSName[1]
	}

	if a.recordTypeInAffix() {
		return prefix + DNSName[0] + suffix + rest
	}

	inlineHead := prefix + recordT + DNSName[0] + suffix
	if maxLabelLen(inlineHead) <= maxDNSLabelLen {
		return inlineHead + rest
	}

	// Separate-label fallback: <prefix><recordType><suffix>.<firstLabel>.<rest>
	return prefix + recordType + suffix + "." + DNSName[0] + rest
}

// maxLabelLen returns the length of the longest dot-separated label in s.
func maxLabelLen(s string) int {
	max := 0
	for _, label := range strings.Split(s, ".") {
		if len(label) > max {
			max = len(label)
		}
	}
	return max
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
// when not using '%{record_type}' in the prefix/suffix.
//
// Recognises both the inline form "<recordType>-<rest>" and the separate-label
// fallback form "<recordType>.<rest>" emitted by ToTXTName when the inline form
// would overflow RFC 1035's per-label limit. The bare "<recordType>" case is
// encountered in suffix mode after the suffix has been trimmed.
func extractRecordTypeDefaultPosition(name string) (string, string) {
	if dashIdx := strings.IndexByte(name, '-'); dashIdx > 0 {
		head := name[:dashIdx]
		for _, t := range supportedRecords {
			if head == strings.ToLower(t) {
				return name[dashIdx+1:], t
			}
		}
	}
	if dotIdx := strings.IndexByte(name, '.'); dotIdx > 0 {
		head := name[:dotIdx]
		for _, t := range supportedRecords {
			if head == strings.ToLower(t) {
				return name[dotIdx+1:], t
			}
		}
	}
	for _, t := range supportedRecords {
		if name == strings.ToLower(t) {
			return "", t
		}
	}
	return name, ""
}
