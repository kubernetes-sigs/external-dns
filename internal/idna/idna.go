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

package idna

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"golang.org/x/net/idna"
)

var (
	Profile = idna.New(
		idna.MapForLookup(),
		idna.Transitional(true),
		idna.StrictDomainName(false),
	)
)

// ToASCII converts a value to ASCII using the default IDNA profile.
// It returns the original value on conversion errors.
func ToASCII(value string) string {
	ascii, err := Profile.ToASCII(value)
	if err != nil {
		log.Debugf("Failed to convert %q to ASCII: %v", value, err)
		return value
	}
	return ascii
}

// ToUnicode converts a value to Unicode using the default IDNA profile.
// It returns the original value on conversion errors.
func ToUnicode(value string) string {
	unicode, err := Profile.ToUnicode(value)
	if err != nil {
		log.Debugf("Failed to convert %q to Unicode: %v", value, err)
		return value
	}
	return unicode
}

// NormalizeDNSName converts a DNS name to a canonical form, so that we can use string equality
// it: removes space, get ASCII version of dnsName complient with Section 5 of RFC 5891, ensures there is a trailing dot
func NormalizeDNSName(dnsName string) string {
	s, err := Profile.ToASCII(strings.TrimSpace(dnsName))
	if err != nil {
		log.Warnf(`Got error while parsing DNSName %s: %v`, dnsName, err)
	}
	if !strings.HasSuffix(s, ".") {
		s += "."
	}
	return s
}
