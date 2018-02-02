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

package source

import (
	"net"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/kubernetes-incubator/external-dns/endpoint"
)

// CIDRs convert an array of strings into an array of CIDR objects.
func CIDRs(raw []string) []*net.IPNet {
	cidrs := make([]*net.IPNet, len(raw))
	for i, f := range raw {
		_, cidr, err := net.ParseCIDR(f)
		if err != nil {
			log.Fatalf("failed to parse cidr '%s': %v", f, err)
		}
		cidrs[i] = cidr
	}
	return cidrs
}

// filterSource is a Source that removes duplicate endpoints from its wrapped source.
type filterSource struct {
	ignoreCIDRs []*net.IPNet
	ignoreDNS   []string
	source      Source
}

// NewFilterSource creates a new FilterSource wrapping the provided Source.
func NewFilterSource(cidrs []*net.IPNet, dns []string, source Source) Source {
	if cidrs == nil {
		cidrs = []*net.IPNet{}
	}
	if dns == nil {
		dns = []string{}
	}
	return &filterSource{ignoreCIDRs: cidrs, ignoreDNS: dns, source: source}
}

// Endpoints collects endpoints from its wrapped source and returns them without filtered IPs.
func (ms *filterSource) Endpoints() ([]*endpoint.Endpoint, error) {
	result := []*endpoint.Endpoint{}

	endpoints, err := ms.source.Endpoints()
	if err != nil {
		return nil, err
	}

EPS:
	// filter IPs
	for _, ep := range endpoints {
		if ep.RecordType == endpoint.RecordTypeA {
			for _, t := range ep.Targets {
				ip := net.ParseIP(t)
				if ip != nil {
					for _, cidr := range ms.ignoreCIDRs {
						if cidr.Contains(ip) {
							ep.Targets = ep.Targets.Remove(t)
							if len(ep.Targets) == 0 {
								log.Debugf("Removing endpoint %s omitted because of CIDR %s", ep, cidr)
								continue EPS
							} else {
								log.Debugf("Removing target %s from endpoint %s because of CIDR %s", t, ep, cidr)
							}
						}
					}
				}
			}
		}

		// filter DNS entries
		for _, dns := range ms.ignoreDNS {
			if ms.matchDNS(ep.DNSName, dns) {
				log.Debugf("Removing endpoint %s omitted because of DNS %s", ep, dns)
				continue EPS
			}
		}

		result = append(result, ep)
	}

	return result, nil
}

func (ms *filterSource) matchDNS(name, pattern string) bool {
	if strings.HasPrefix(pattern, "*.") {
		if strings.HasPrefix(name, "*.") {
			return false
		}
		i := strings.Index(name, ".")
		if i >= 0 {
			if name[i+1:] == pattern[2:] {
				return true
			}
		}
	}
	return name == pattern
}
