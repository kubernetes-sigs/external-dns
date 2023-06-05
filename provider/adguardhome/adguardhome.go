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

package adguardhome

import (
	"context"

	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

// AdGuardHomeConfig is used for configuring a AdGuardHomeProvider.
type AdGuardHomeConfig struct {
	// The root URL of the AdGuardHome server.
	Server string
	// An optional username if the server is protected.
	Username string
	// An optional password if the server is protected.
	Password string
	// Disable verification of TLS certificates.
	TLSInsecureSkipVerify bool
	// A filter to apply when looking up and applying records.
	DomainFilter endpoint.DomainFilter
	// Do nothing and log what would have changed to stdout.
	DryRun bool
}

type AdGuardHomeProvider struct {
	provider.BaseProvider
	api adGuardHomeAPI
}

func NewAdGuardHomeProvider(cfg AdGuardHomeConfig) (*AdGuardHomeProvider, error) {
	api, err := newAdGuardHomeClient(cfg)
	if err != nil {
		return nil, err
	}
	return &AdGuardHomeProvider{api: api}, nil
}

// Records implements Provider, populating a slice of endpoints from
// AdGuardHome local DNS.
func (p *AdGuardHomeProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	records, err := p.api.listRecords(ctx)
	log.Debugf("Records: %s", records)
	return records, err
}

// ApplyChanges implements Provider, syncing desired state with the AdGuardHome server Local DNS.
func (p *AdGuardHomeProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	log.Debugf("ApplyChanges (Create: %d, UpdateOld: %d, UpdateNew: %d, Delete: %d)", len(changes.Create), len(changes.UpdateOld), len(changes.UpdateNew), len(changes.Delete))

	for _, ep := range changes.Create {
		if err := p.api.createRecord(ctx, ep); err != nil {
			return err
		}
		log.Debugf("CREATE: %s", ep)
	}

	for _, ep := range changes.Delete {
		if err := p.api.deleteRecord(ctx, ep); err != nil {
			return err
		}
		log.Debugf("DELETE: %s", ep)
	}

	for i, new := range changes.UpdateNew {
		old := changes.UpdateOld[i]
		if err := p.api.updateRecord(ctx, old, new); err != nil {
			return err
		}
		log.Debugf("Old: %s -> New: %s", old, new)
	}
	return nil
}
