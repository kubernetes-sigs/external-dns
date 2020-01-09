package provider

import (
  "bytes"
  "context"
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "strconv"
  "strings"

  log "github.com/sirupsen/logrus"

  yaml "gopkg.in/yaml.v2"

  "github.com/kubernetes-sigs/external-dns/endpoint"
  "github.com/kubernetes-sigs/external-dns/plan"
)

type bluecatConfig struct {
  GatewayHost          string `json:"gatewayHost" yaml:"gatewayHost"`
  GatewayUsername      string `json:"gatewayUsername" yaml:"gatewayUsername"`
  GatewayPassword      string `json:"gatewayPassword" yaml:"gatewayPassword"`
  BluecatConfiguration string `json:"bluecatConfiguration" yaml:"bluecatConfiguration"`
  View                 string `json:"dnsView" yaml:"dnsView"`
}

type BluecatProvider struct {
  domainFilter  DomainFilter
  zoneIdFilter  ZoneIDFilter
  dryRun        bool
  gatewayClient GatewayClient
}

type IGatewayClient interface {

}

type GatewayClient struct {
  Token                string
  BluecatConfiguration string
  View                 string
}

type BluecatZone struct {
  Id           int
  Name         string
  AbsoluteName string
  Deployable   bool
}

type BluecatHostRecord struct {
  Id            int
  Name          string
  AbsoluteName  string
  Addresses     []string
  ReverseRecord bool
  Type          string
}

type BluecatCNAMERecord struct {
  Id               int
  Name             string
  AbsoluteName     string
  LinkedRecordName string
  Type             string
}

func NewBluecatProvider(configFile string, domainFilter DomainFilter, zoneIDFilter ZoneIDFilter, dryRun bool) (*BluecatProvider, error) {
  contents, err := ioutil.ReadFile(configFile)
  if err != nil {
    return nil, fmt.Errorf("failed to read Bluecat config file '%s': %v", configFile, err)
  }

  cfg := bluecatConfig{}
  err = yaml.Unmarshal(contents, &cfg)
  if err != nil {
    return nil, fmt.Errorf("failed to read Bluecat config file '%s': %v", configFile, err)
  }

  token, err := getBluecatGatewayToken(cfg)
  if err != nil {
    return nil, fmt.Errorf("failed to get API token from Bluecat Gateway: %v", err)
  }
  log.Printf("Gateway API token is: %s", token)

  gatewayClient := NewGatewayClient(token, cfg.BluecatConfiguration, cfg.View)

  provider := &BluecatProvider{
    domainFilter:  domainFilter,
    zoneIdFilter:  zoneIDFilter,
    dryRun:        dryRun,
    gatewayClient: gatewayClient,
  }
  return provider, nil
}

func NewGatewayClient(token, bluecatConfiguration, view string) GatewayClient {
  return GatewayClient{
    Token:                token,
    BluecatConfiguration: bluecatConfiguration,
    View:                 view,
  }
}

func (p *BluecatProvider) Records() (endpoints []*endpoint.Endpoint, err error) {
  zones, err := p.zones()
  if err != nil {
    return nil, fmt.Errorf("could not fetch zones: %s", err)
  }

  for _, zone := range zones {
    log.Debugf("fetching records from zone '%s'", zone)
    var resH []BluecatHostRecord
    err = p.gatewayClient.getHostRecords(zone, &resH)
    if err != nil {
      return nil, fmt.Errorf("could not fetch host records for zone: '%s': %s", zone, err)
    }
    for _, rec := range resH {
      for _, ip := range rec.Addresses {
        endpoints = append(endpoints, endpoint.NewEndpoint(rec.AbsoluteName, endpoint.RecordTypeA, ip))
      }
    }

    var resC []BluecatCNAMERecord
    err = p.gatewayClient.getCNAMERecords(zone, &resC)
    if err != nil {
      return nil, fmt.Errorf("could not fetch CNAME records for zone: '%s': %s", zone, err)
    }
    for _, rec := range resC {
      endpoints = append(endpoints, endpoint.NewEndpoint(rec.AbsoluteName, endpoint.RecordTypeCNAME, rec.LinkedRecordName))
    }
  }

  log.Debugf("fetched %d records from Bluecat", len(endpoints))
  return endpoints, nil
}

func (p *BluecatProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
  return nil
}

func (p *BluecatProvider) zones() ([]string, error) {
  log.Debugf("retrieving Bluecat zones for configuration: %s, view: %s", p.gatewayClient.BluecatConfiguration, p.gatewayClient.View)
  var zones []string

  zonelist, err := p.gatewayClient.getBluecatZones()
  if err != nil {
    return nil, err
  }

  for  _, zone := range zonelist {
    if !p.domainFilter.Match(zone.Name) {
      continue
    }

    if !p.zoneIdFilter.Match(strconv.Itoa(zone.Id)) {
      continue
    }

    zones = append(zones, zone.Name)
  }
  return zones, nil
}

// getBluecatGatewayToken retrieves a Bluecat Gateway API token.
func getBluecatGatewayToken(cfg bluecatConfig) (string, error) {
  body, err := json.Marshal(map[string]string{
    "username": cfg.GatewayUsername,
    "password": cfg.GatewayPassword,
  })
  if err != nil {
    return "", fmt.Errorf("no credentials provided or could not unmarshal credentials for Bluecat Gateway")
  }

  resp, err := http.Post(fmt.Sprintf("%s/%s", cfg.GatewayHost, "/rest_login"), "application/json", bytes.NewBuffer(body))
  if err != nil {
    return "", fmt.Errorf("error obtaining API token from Bluecat Gateway")
  }

  defer resp.Body.Close()

  res, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return "", fmt.Errorf("error reading get_token response from Bluecat Gateway")
  }

  resJSON := map[string]string{}
  err = json.Unmarshal(res, &resJSON)
  if err != nil {
    return "", fmt.Errorf("error unmarshaling json response (auth) from Bluecat Gateway")
  }

  // Example response: {"access_token": "BAMAuthToken: abc123"}
  // We only care about the actual token string - i.e. abc123
  return strings.Split(resJSON["access_token"], " ")[1], nil
}

// TODO
func (c *GatewayClient) getBluecatZones() ([]BluecatZone, error) {
  return []BluecatZone{}, nil
}

// TODO
func (c *GatewayClient) getHostRecords(zone string, records *[]BluecatHostRecord) error {
  return nil
}

// TODO
func (c *GatewayClient) getCNAMERecords(zone string, records *[]BluecatCNAMERecord) error {
  return nil
}