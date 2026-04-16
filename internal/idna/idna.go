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
