/*
Copyright 2021 The Kubernetes Authors.

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

package safedns

import (
  "context"
  "fmt"
  "os"

  ukf_client "github.com/ukfast/sdk-go/pkg/client"
  ukf_connection "github.com/ukfast/sdk-go/pkg/connection"
  "github.com/ukfast/sdk-go/pkg/service/safedns"

  "sigs.k8s.io/external-dns/provider"
  "sigs.k8s.io/external-dns/endpoint"
  "sigs.k8s.io/external-dns/plan"
)

// The SafeDNS interface is a subset of the SafeDNS service API that are actually used.
// Signatures must match exactly.
type SafeDNS interface {
  CreateZoneRecord(zoneName string, req safedns.CreateRecordRequest) (int, error)
  DeleteZoneRecord(zoneName string, recordID int) error
  GetZone(zoneName string) (safedns.Zone, error)
  GetZoneRecord(zoneName string, recordID int) (safedns.Record, error)
  GetZoneRecords(zoneName string, parameters ukf_connection.APIRequestParameters) ([]safedns.Record, error)
  GetZones(parameters ukf_connection.APIRequestParameters) ([]safedns.Zone, error)
  PatchZoneRecord(zoneName string, recordID int, patch safedns.PatchRecordRequest) (int, error)
  UpdateZoneRecord(zoneName string, record safedns.Record) (int, error)
}

type SafeDNSProvider struct {
  provider.BaseProvider
  Client SafeDNS
  // Only consider hosted zones managing domains ending in this suffix
  domainFilter      endpoint.DomainFilter
  DryRun            bool
  APIRequestParams  ukf_connection.APIRequestParameters
}

func NewSafeDNSProvider(domainFilter endpoint.DomainFilter, dryRun bool) (*SafeDNSProvider, error) {
  token, ok := os.LookupEnv("SAFEDNS_TOKEN")
  if !ok {
    return nil, fmt.Errorf("No SAFEDNS_TOKEN found in environment")
  }

  ukfAPIConnection := ukf_connection.NewAPIKeyCredentialsAPIConnection(token)
  ukfClient := ukf_client.NewClient(ukfAPIConnection)
  safeDNS := ukfClient.SafeDNSService()

  provider := &SafeDNSProvider{
    Client:           safeDNS,
    domainFilter:     domainFilter,
    DryRun:           dryRun,
    APIRequestParams: *ukf_connection.NewAPIRequestParameters(),
  }
  return provider, nil
}

// Zones returns the list of hosted zones in the SafeDNS account
func (p *SafeDNSProvider) Zones(ctx context.Context) ([]safedns.Zone, error) {
  var zones []safedns.Zone

  allZones, err := p.Client.GetZones(p.APIRequestParams)
  if err != nil {
    return nil, err
  }

  // Check each found zone to see whether they match the domain filter provided. If they do, append it to the array of
  // zones defined above. If not, continue to the next item in the loop.
  for _, zone := range allZones {
    if p.domainFilter.Match(zone.Name) {
      zones = append(zones, zone)
    } else {
      continue
    }
  }
  return zones, nil
}

// Records returns a list of Endpoint resources created from all records in supported zones.
func (p *SafeDNSProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
  zones, err := p.Zones(ctx)
  if err != nil {
    return nil, err
  }

  endpoints := []*endpoint.Endpoint{}
  for _, zone := range zones {
    // For each zone in the zonelist, get all records of an ExternalDNS supported type.
    records, err := p.Client.GetZoneRecords(zone.Name, p.APIRequestParams)
    if err != nil {
      return nil, err
    }
    for _, r := range records {
      if provider.SupportedRecordType(string(r.Type)) {
        endpoints = append(endpoints, endpoint.NewEndpointWithTTL(r.Name, string(r.Type), endpoint.TTL(r.TTL), r.Content))
      }
    }
  }
  return endpoints, nil
}

// ApplyChanges applies a given set of changes in a given zone.
func (p *SafeDNSProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
  // TODO: Implement this
  return nil
}
