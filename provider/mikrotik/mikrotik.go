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

package mikrotik

import (
	"context"
	"errors"
	"strings"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/provider/mikrotik/api"
	"sigs.k8s.io/external-dns/provider/mikrotik/common"
	"sigs.k8s.io/external-dns/provider/mikrotik/rest"
)

// ErrNoMikrotikServer is returned when there is no Mikrotik server configured
// in the environment.
var ErrNoMikrotikServer = errors.New("no Mikrotik server found in the environment or flags")

// ErrInvalidMikrotikServer is returned when the configured Mikrotik server format
// is invalid
var ErrInvalidMikrotikServer = errors.New("invalid Mikrotik server found in the environment or flags (must begin with 'http[s]://' or 'api://')")

// MikrotikConfig is used for configuring a MikrotikProvider.
type MikrotikConfig struct {
	// The root URL of the Mikrotik router.
	Server string
	// Username for the router.
	Username string
	// Password for the router.
	Password string
	// Skip TLS verification
	TLSInsecureSkipVerify bool
	// A filter to apply when looking up and applying records.
	DomainFilter endpoint.DomainFilter
	// The minimum TTL
	MinimumTTL endpoint.TTL
	// Do nothing and log what would have changed to stdout.
	DryRun bool
	// The owner Id
	OwnerId string
}

// MikrotikProvider is an implementation of Provider for Pi-hole Local DNS.
type MikrotikProvider struct {
	provider.BaseProvider
	connection common.MikrotikConnection
}

// NewMikrotikProvider initializes a new Pi-hole Local DNS based Provider.
func NewMikrotikProvider(cfg MikrotikConfig) (*MikrotikProvider, error) {
	if cfg.Server == "" {
		return nil, ErrNoMikrotikServer
	}

	split := strings.Split(cfg.Server, "://")
	if len(split) != 2 || (split[0] != "http" && split[0] != "https" && split[0] != "api") {
		return nil, ErrInvalidMikrotikServer
	}
	proto := split[0]
	server := split[1]

	// arbitrary limit because if it's zero, nothing works
	if cfg.MinimumTTL < 1 {
		cfg.MinimumTTL = 1
	}

	var conn common.MikrotikConnection
	var err error

	if proto == "api" {
		conn, err = api.NewMikrotikAPI(
			server,
			cfg.Username,
			cfg.Password,
			cfg.OwnerId,
			cfg.DomainFilter,
			cfg.MinimumTTL,
			cfg.DryRun,
		)
	} else {
		conn, err = rest.NewMikrotikREST(
			cfg.Server,
			cfg.Username,
			cfg.Password,
			cfg.OwnerId,
			cfg.DomainFilter,
			cfg.MinimumTTL,
			cfg.DryRun,
			cfg.TLSInsecureSkipVerify,
		)
	}

	if err != nil {
		return nil, err
	}
	return &MikrotikProvider{connection: conn}, nil
}

// Records implements Provider, populating a slice of endpoints from
// Pi-Hole local DNS.
func (p *MikrotikProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	aRecs, err := p.connection.ListRecords("")
	if err != nil {
		return nil, err
	}
	return aRecs, nil
}

// ApplyChanges implements Provider, syncing desired state with the Pi-hole server Local DNS.
func (p *MikrotikProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	// Handle deletions first - there are no endpoints for updating in place.
	for _, ep := range changes.Delete {
		if err := p.connection.DeleteRecord(ep); err != nil {
			return err
		}
	}
	for _, ep := range changes.UpdateOld {
		// log.Infof("MikrotikProvider.updateOld: %#v\n", ep)
		if err := p.connection.DeleteRecord(ep); err != nil {
			return err
		}
	}

	// Handle desired state
	for _, ep := range changes.Create {
		// log.Infof("MikrotikProvider.create: %#v\n", ep)
		if err := p.connection.CreateRecord(ep); err != nil {
			return err
		}
	}
	for _, ep := range changes.UpdateNew {
		// log.Infof("MikrotikProvider.updateNew: %#v\n", ep)
		if err := p.connection.CreateRecord(ep); err != nil {
			return err
		}
	}

	return nil
}
