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
	"fmt"
	"strings"

	"github.com/miekg/dns"
	log "github.com/sirupsen/logrus"
)

// NAPTRTarget is a parsed NAPTR target (RFC 2915): order preference "flags" "service" "regexp" replacement.
type NAPTRTarget struct {
	order       uint16
	preference  uint16
	flags       string
	service     string
	regexp      string
	replacement string
}

// NewNAPTRRecord parses a NAPTR target via miekg/dns because the quoted fields may contain spaces; a relative replacement becomes absolute.
func NewNAPTRRecord(target string) (*NAPTRTarget, error) {
	rr, err := dns.NewRR("x. 0 IN NAPTR " + strings.TrimSpace(target))
	if err != nil || rr == nil {
		return nil, fmt.Errorf("invalid NAPTR target %q: %w", target, err)
	}
	naptr, ok := rr.(*dns.NAPTR)
	if !ok {
		return nil, fmt.Errorf("invalid NAPTR target %q: parsed as %T", target, rr)
	}
	return &NAPTRTarget{
		order:       naptr.Order,
		preference:  naptr.Preference,
		flags:       naptr.Flags,
		service:     naptr.Service,
		regexp:      naptr.Regexp,
		replacement: naptr.Replacement,
	}, nil
}

// NewNAPTRTargetFromFields builds a NAPTRTarget from split fields; a replacement without a trailing dot becomes absolute, the root "." stays.
func NewNAPTRTargetFromFields(order, preference uint16, flags, service, regexp, replacement string) *NAPTRTarget {
	if replacement != "." && !strings.HasSuffix(replacement, ".") {
		replacement += "."
	}
	return &NAPTRTarget{
		order:       order,
		preference:  preference,
		flags:       flags,
		service:     service,
		regexp:      regexp,
		replacement: replacement,
	}
}

// GetOrder returns the NAPTR order field (RFC 2915 section 2).
func (n *NAPTRTarget) GetOrder() uint16 { return n.order }

// GetPreference returns the NAPTR preference field (RFC 2915 section 2).
func (n *NAPTRTarget) GetPreference() uint16 { return n.preference }

// GetFlags returns the NAPTR flags field (RFC 2915 section 2).
func (n *NAPTRTarget) GetFlags() string { return n.flags }

// GetService returns the NAPTR service field (RFC 2915 section 2).
func (n *NAPTRTarget) GetService() string { return n.service }

// GetRegexp returns the NAPTR regexp field (RFC 2915 section 2).
func (n *NAPTRTarget) GetRegexp() string { return n.regexp }

// GetReplacement returns the NAPTR replacement field as a fully-qualified name.
func (n *NAPTRTarget) GetReplacement() string { return n.replacement }

// String renders the canonical presentation form via miekg/dns so quoting round-trips with NewNAPTRRecord.
func (n *NAPTRTarget) String() string {
	rr := &dns.NAPTR{
		Order:       n.order,
		Preference:  n.preference,
		Flags:       n.flags,
		Service:     n.service,
		Regexp:      n.regexp,
		Replacement: n.replacement,
	}
	return strings.TrimPrefix(rr.String(), rr.Hdr.String())
}

// ValidateNAPTRRecord reports whether all targets are valid NAPTR record values.
func (t Targets) ValidateNAPTRRecord() bool {
	for _, target := range t {
		if _, err := NewNAPTRRecord(target); err != nil {
			log.Debugf("Invalid NAPTR record target: %s. %v", target, err)
			return false
		}
	}
	return true
}
