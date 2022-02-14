package source

import (
	"context"
	log "github.com/sirupsen/logrus"
	"net"
	"sigs.k8s.io/external-dns/endpoint"
)

type suppressedSource struct {
	unfiltered Source
}

func NewSuppressedSource(original Source) Source {
	return &suppressedSource{
		unfiltered: original,
	}
}

func getIp4Targets(targets endpoint.Targets) endpoint.Targets {
	result := []string{}
	for _, target := range targets {
		ip := net.ParseIP(target)
		if ip != nil && ip.To4() != nil {
			// This is an IPv4
			result = append(result, target)
		} else {
			log.Debugf("Suppressed %s, not IPv4 address", target)
		}
	}
	return result
}

func (s *suppressedSource) Endpoints(ctx context.Context) ([]*endpoint.Endpoint, error) {
	endpoints, err := s.unfiltered.Endpoints(ctx)
	if err != nil {
		return endpoints, err
	}
	results := []*endpoint.Endpoint{}
	for _, endpoint := range endpoints {
		targets := getIp4Targets(endpoint.Targets)
		if len(targets) > 0 {
			endpointCopy := *endpoint
			endpointCopy.Targets = targets

			results = append(results, &endpointCopy)
		} else {
			log.Debugf("Suppressed %s. No IPv4 targets", endpoint.DNSName)
		}
	}

	return results, nil
}

func (s *suppressedSource) AddEventHandler(ctx context.Context, f func()) {
	s.unfiltered.AddEventHandler(ctx, f)
}
