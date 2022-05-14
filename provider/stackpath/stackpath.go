/*
Copyright 2020 The Kubernetes Authors.

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

package stackpath

import (
	"context"
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"

	"github.com/wmarchesi123/stackpath-go/pkg/dns"
	"github.com/wmarchesi123/stackpath-go/pkg/oauth2"
)

type StackPathProvider struct {
	provider.BaseProvider
	client       *dns.APIClient
	context      context.Context
	domainFilter endpoint.DomainFilter
	zoneIdFilter provider.ZoneIDFilter
	stackId      string
	dryRun       bool
}

type StackPathConfig struct {
	Context      context.Context
	DomainFilter endpoint.DomainFilter
	ZoneIDFilter provider.ZoneIDFilter
	DryRun       bool
}

func NewStackPathProvider(config StackPathConfig) (*StackPathProvider, error) {
	clientId, ok := os.LookupEnv("STACKPATH_CLIENT_ID")
	if !ok {
		return nil, fmt.Errorf("STACKPATH_CLIENT_ID environment variable is not set")
	}

	clientSecret, ok := os.LookupEnv("STACKPATH_CLIENT_SECRET")
	if !ok {
		return nil, fmt.Errorf("STACKPATH_CLIENT_SECRET environment variable is not set")
	}

	stackId, ok := os.LookupEnv("STACKPATH_STACK_ID")
	if !ok {
		return nil, fmt.Errorf("STACKPATH_STACK_ID environment variable is not set")
	}

	oauthSource := oauth2.NewTokenSource(clientId, clientSecret, oauth2.HTTPClientOption(http.DefaultClient))
	_, err := oauthSource.Token()
	if err != nil {
		return nil, fmt.Errorf("STACKPATH_CLIENT_ID or STACKPATH_CLIENT_SECRET environment variable(s) are invalid")
	}

	authorizedContext := context.WithValue(config.Context, dns.ContextOAuth2, oauthSource)

	clientConfig := dns.NewConfiguration()

	client := dns.NewAPIClient(clientConfig)

	provider := &StackPathProvider{
		client:       client,
		context:      authorizedContext,
		domainFilter: config.DomainFilter,
		zoneIdFilter: config.ZoneIDFilter,
		stackId:      stackId,
		dryRun:       config.DryRun,
	}

	return provider, nil
}

//Base Provider Functions

func (p *StackPathProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {

	return nil, nil
}

func (p *StackPathProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	return nil
}

func (p *StackPathProvider) AdjustEndpoints(endpoints []*endpoint.Endpoint) []*endpoint.Endpoint {
	return endpoints
}

func (p *StackPathProvider) PropertyValuesEqual(name, previous, current string) bool {
	return previous == current
}

func (p *StackPathProvider) GetDomainFilter() endpoint.DomainFilterInterface {
	return endpoint.DomainFilter{}
}

//StackPath Helper Functions

func (p *StackPathProvider) filteredZones() ([]dns.ZoneZone, error) {
	zoneResponse, _, err := p.client.ZonesApi.GetZones(p.context, p.stackId).Execute()
	if err != nil {
		return nil, err
	}

	zones := zoneResponse.GetZones()

	filteredZones := []dns.ZoneZone{}

	for _, zone := range zones {
		if p.zoneIdFilter.Match(zone.GetId()) && p.domainFilter.Match(zone.GetDomain()) {
			filteredZones = append(filteredZones, zone)
			log.Debugf("Matched zone " + zone.GetId())
		} else {
			log.Debugf("Filtered zone " + zone.GetId())
		}
	}

	return filteredZones, nil
}
