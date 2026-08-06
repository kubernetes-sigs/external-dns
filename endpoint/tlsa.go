/*
Copyright 2026 The Kubernetes Authors.

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

package endpoint

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

// TLSATarget represents a single TLSA (DANE) record target in RFC 6698
// presentation format: "<usage> <selector> <matchingType> <certificate>".
//
// Providers whose APIs model TLSA as discrete fields rather than a single
// content string use this to decompose a target, and to re-render one
// canonically so that a record read back from the provider compares equal to
// the record that was written.
type TLSATarget struct {
	usage        uint8
	selector     uint8
	matchingType uint8
	certificate  string
}

// NewTLSARecord parses a TLSA target in presentation format.
//
// The certificate association data is normalised to lowercase hex without
// separators, per RFC 6698 section 2.2. Providers differ in how they render it,
// so normalising on both read and write keeps record identity stable and
// prevents a spurious update on every reconciliation.
func NewTLSARecord(target string) (*TLSATarget, error) {
	fields := strings.Fields(target)
	if len(fields) != 4 {
		return nil, fmt.Errorf("invalid TLSA target %q: expected 4 fields (usage selector matchingType certificate), got %d", target, len(fields))
	}

	usage, err := parseTLSAField(fields[0], "usage", 3)
	if err != nil {
		return nil, err
	}
	selector, err := parseTLSAField(fields[1], "selector", 1)
	if err != nil {
		return nil, err
	}
	matchingType, err := parseTLSAField(fields[2], "matchingType", 2)
	if err != nil {
		return nil, err
	}

	certificate := strings.ToLower(strings.ReplaceAll(fields[3], ":", ""))
	if certificate == "" {
		return nil, fmt.Errorf("invalid TLSA target %q: certificate association data is empty", target)
	}
	if _, err := hex.DecodeString(certificate); err != nil {
		return nil, fmt.Errorf("invalid TLSA target %q: certificate association data is not valid hex: %w", target, err)
	}
	// A digest is fixed width; matching type 0 carries the full certificate or
	// SPKI and has no expected length.
	switch matchingType {
	case 1:
		if len(certificate) != 64 {
			return nil, fmt.Errorf("invalid TLSA target %q: matchingType 1 (SHA-256) needs 64 hex characters, got %d", target, len(certificate))
		}
	case 2:
		if len(certificate) != 128 {
			return nil, fmt.Errorf("invalid TLSA target %q: matchingType 2 (SHA-512) needs 128 hex characters, got %d", target, len(certificate))
		}
	}

	return &TLSATarget{
		usage:        usage,
		selector:     selector,
		matchingType: matchingType,
		certificate:  certificate,
	}, nil
}

func parseTLSAField(value, name string, maxValue uint8) (uint8, error) {
	parsed, err := strconv.ParseUint(value, 10, 8)
	if err != nil {
		return 0, fmt.Errorf("invalid TLSA %s %q: not an integer", name, value)
	}
	if parsed > uint64(maxValue) {
		return 0, fmt.Errorf("invalid TLSA %s %d: must be between 0 and %d", name, parsed, maxValue)
	}
	return uint8(parsed), nil
}

// GetUsage returns the TLSA certificate usage field (RFC 6698 section 2.1.1).
func (t *TLSATarget) GetUsage() uint8 { return t.usage }

// GetSelector returns the TLSA selector field (RFC 6698 section 2.1.2).
func (t *TLSATarget) GetSelector() uint8 { return t.selector }

// GetMatchingType returns the TLSA matching type field (RFC 6698 section 2.1.3).
func (t *TLSATarget) GetMatchingType() uint8 { return t.matchingType }

// GetCertificate returns the certificate association data as lowercase hex.
func (t *TLSATarget) GetCertificate() string { return t.certificate }

// String renders the target in canonical presentation format.
func (t *TLSATarget) String() string {
	return fmt.Sprintf("%d %d %d %s", t.usage, t.selector, t.matchingType, t.certificate)
}
