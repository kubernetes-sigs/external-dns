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

	log "github.com/sirupsen/logrus"

	"github.com/kubernetes-incubator/external-dns/endpoint"
)

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

func FilterSources(cidrs []*net.IPNet, sources ...Source) []Source {
	result := make([]Source, len(sources))
	for i, s := range sources {
		result[i] = NewFilterSource(cidrs, s)
	}
	return result
}

// filterSource is a Source that removes duplicate endpoints from its wrapped source.
type filterSource struct {
	ignoreCIDRs []*net.IPNet
	source      Source
}

// NewFilterSource creates a new FilterSource wrapping the provided Source.
func NewFilterSource(cidrs []*net.IPNet, source Source) Source {
	return &filterSource{ignoreCIDRs: cidrs, source: source}
}

// Endpoints collects endpoints from its wrapped source and returns them without filtered IPs.
func (ms *filterSource) Endpoints() ([]*endpoint.Endpoint, error) {
	result := []*endpoint.Endpoint{}

	endpoints, err := ms.source.Endpoints()
	if err != nil {
		return nil, err
	}

EPS:
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

		result = append(result, ep)
	}

	return result, nil
}
