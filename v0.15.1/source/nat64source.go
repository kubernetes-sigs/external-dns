/*
Copyright 2024 The Kubernetes Authors.

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
	"context"
	"fmt"
	"net/netip"

	"sigs.k8s.io/external-dns/endpoint"
)

// nat64Source is a Source that adds A endpoints for AAAA records including an NAT64 address.
type nat64Source struct {
	source        Source
	nat64Prefixes []string
}

// NewNAT64Source creates a new nat64Source wrapping the provided Source.
func NewNAT64Source(source Source, nat64Prefixes []string) Source {
	return &nat64Source{source: source, nat64Prefixes: nat64Prefixes}
}

// Endpoints collects endpoints from its wrapped source and returns them without duplicates.
func (s *nat64Source) Endpoints(ctx context.Context) ([]*endpoint.Endpoint, error) {
	parsedNAT64Prefixes := make([]netip.Prefix, 0)
	for _, prefix := range s.nat64Prefixes {
		pPrefix, err := netip.ParsePrefix(prefix)
		if err != nil {
			return nil, err
		}

		if pPrefix.Bits() != 96 {
			return nil, fmt.Errorf("NAT64 prefixes need to be /96 prefixes.")
		}
		parsedNAT64Prefixes = append(parsedNAT64Prefixes, pPrefix)
	}

	additionalEndpoints := []*endpoint.Endpoint{}

	endpoints, err := s.source.Endpoints(ctx)
	if err != nil {
		return nil, err
	}

	for _, ep := range endpoints {
		if ep.RecordType != endpoint.RecordTypeAAAA {
			continue
		}

		v4Targets := make([]string, 0)

		for _, target := range ep.Targets {
			ip, err := netip.ParseAddr(target)
			if err != nil {
				return nil, err
			}

			var sPrefix *netip.Prefix

			for _, cPrefix := range parsedNAT64Prefixes {
				if cPrefix.Contains(ip) {
					sPrefix = &cPrefix
				}
			}

			// If we do not have a NAT64 prefix, we skip this record.
			if sPrefix == nil {
				continue
			}

			ipBytes := ip.As16()
			v4AddrBytes := ipBytes[12:16]

			v4Addr, isOk := netip.AddrFromSlice(v4AddrBytes)
			if !isOk {
				return nil, fmt.Errorf("Could not parse %v to IPv4 address", v4AddrBytes)
			}

			v4Targets = append(v4Targets, v4Addr.String())
		}

		if len(v4Targets) == 0 {
			continue
		}

		v4EP := ep.DeepCopy()
		v4EP.Targets = v4Targets
		v4EP.RecordType = endpoint.RecordTypeA

		additionalEndpoints = append(additionalEndpoints, v4EP)
	}
	return append(endpoints, additionalEndpoints...), nil
}

func (s *nat64Source) AddEventHandler(ctx context.Context, handler func()) {
	s.source.AddEventHandler(ctx, handler)
}
