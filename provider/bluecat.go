package provider

import (
  "bytes"
  "context"
  "crypto/tls"
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "strconv"
  "strings"

  log "github.com/sirupsen/logrus"

  yaml "gopkg.in/yaml.v2"

  "sigs.k8s.io/external-dns/endpoint"
  "sigs.k8s.io/external-dns/plan"
)

type bluecatConfig struct {
  GatewayHost          string `json:"gatewayHost" yaml:"gatewayHost"`
  GatewayUsername      string `json:"gatewayUsername" yaml:"gatewayUsername"`
  GatewayPassword      string `json:"gatewayPassword" yaml:"gatewayPassword"`
  DNSConfiguration     string `json:"dnsConfiguration" yaml:"dnsConfiguration"`
  View                 string `json:"dnsView" yaml:"dnsView"`
  RootZone             string `json:"rootZone" yaml:"rootZone"`
}

type BluecatProvider struct {
  domainFilter  DomainFilter
  zoneIdFilter  ZoneIDFilter
  dryRun        bool
  gatewayClient GatewayClient
}

type GatewayClient struct {
  Cookie           http.Cookie
  Token            string
  Host             string
  DNSConfiguration string
  View             string
  RootZone         string
}

type BluecatZone struct {
  Id         int    `json:"id"`
  Name       string `json:"name"`
  Properties string `json:"properties"`
  Type       string `json:"type"`
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

  token, cookie, err := getBluecatGatewayToken(cfg)
  if err != nil {
    return nil, fmt.Errorf("failed to get API token from Bluecat Gateway: %v", err)
  }
  log.Printf("Gateway API token is: %s", token)

  gatewayClient := NewGatewayClient(cookie, token, cfg.GatewayHost, cfg.DNSConfiguration, cfg.View, cfg.RootZone)

  provider := &BluecatProvider{
    domainFilter:  domainFilter,
    zoneIdFilter:  zoneIDFilter,
    dryRun:        dryRun,
    gatewayClient: gatewayClient,
  }
  return provider, nil
}

func NewGatewayClient(cookie http.Cookie, token, gatewayHost, dnsConfiguration, view, rootZone string) GatewayClient {
  // Right now the Bluecat gateway doesn't seem to have a way to get the root zone from the API. If the user
  // doesn't provide one via the config file we'll assume it's 'com'
  if rootZone == "" {
    rootZone = "com"
  }
  return GatewayClient{
    Cookie:           cookie,
    Token:            token,
    Host:             gatewayHost,
    DNSConfiguration: dnsConfiguration,
    View:             view,
    RootZone:         rootZone,
  }
}

func (p *BluecatProvider) Records(ctx context.Context) (endpoints []*endpoint.Endpoint, err error) {
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
        ep := endpoint.NewEndpoint(rec.AbsoluteName, endpoint.RecordTypeA, ip)
        endpoints = append(endpoints, ep)
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
  _, err := p.zones()
  if err != nil {
    return err
  }

  return nil
}

func (p *BluecatProvider) zones() ([]string, error) {
  log.Debugf("retrieving Bluecat zones for configuration: %s, view: %s", p.gatewayClient.DNSConfiguration, p.gatewayClient.View)
  var zones []string

  zonelist, err := p.gatewayClient.getBluecatZones(p.gatewayClient.RootZone)
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

    zoneProps := splitProperties(zone.Properties)

    zones = append(zones, zoneProps["absoluteName"])
  }
  log.Debugf("found %n zones", len(zones))
  return zones, nil
}

// getBluecatGatewayToken retrieves a Bluecat Gateway API token.
func getBluecatGatewayToken(cfg bluecatConfig) (string, http.Cookie, error) {
  body, err := json.Marshal(map[string]string{
    "username": cfg.GatewayUsername,
    "password": cfg.GatewayPassword,
  })
  if err != nil {
    return "", http.Cookie{}, fmt.Errorf("no credentials provided or could not unmarshal credentials for Bluecat Gateway")
  }

  transportCfg := &http.Transport{
    TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //ignore self-signed SSL cert check
  }
  c := &http.Client{Transport: transportCfg}

  resp, err := c.Post(cfg.GatewayHost + "/rest_login", "application/json", bytes.NewBuffer(body))
  if err != nil {
    return "", http.Cookie{}, fmt.Errorf("error obtaining API token from Bluecat Gateway: %s", err)
  }

  defer resp.Body.Close()

  res, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return "", http.Cookie{}, fmt.Errorf("error reading get_token response from Bluecat Gateway")
  }

  resJSON := map[string]string{}
  err = json.Unmarshal(res, &resJSON)
  if err != nil {
    return "", http.Cookie{}, fmt.Errorf("error unmarshaling json response (auth) from Bluecat Gateway")
  }

  // Example response: {"access_token": "BAMAuthToken: abc123"}
  // We only care about the actual token string - i.e. abc123
  // The gateway also creates a cookie as part of the response. This seems to be the actual auth mechanism, at least
  // for now.
  return strings.Split(resJSON["access_token"], " ")[1], *resp.Cookies()[0], nil
}

func (c *GatewayClient) getBluecatZones(zoneName string) ([]BluecatZone, error) {
  transportCfg := &http.Transport{
    TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //ignore self-signed SSL cert check
  }
  client := &http.Client{
    Transport: transportCfg,
  }

  zoneTree := expandZone(zoneName)
  url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration + "/views/" + c.View + "/" + zoneTree
  req, err := http.NewRequest("GET", url, nil)
  req.Header.Add("Accept", "application/json")
  req.Header.Add("Authorization", "Basic " + c.Token)
  req.AddCookie(&c.Cookie)

  resp, err := client.Do(req)
  if err != nil {
    return nil, fmt.Errorf("error retrieving zone(s) from gateway: %s, %s", zoneName, err)
  }

  defer resp.Body.Close()

  zones := []BluecatZone{}
  json.NewDecoder(resp.Body).Decode(&zones)

  // Bluecat Gateway only returns subzones one level deeper than the provided zone so this recursion is needed
  // to traverse subzones until none are returned
  for _, zone := range zones {
    zoneProps := splitProperties(zone.Properties)
    subZones, err := c.getBluecatZones(zoneProps["absoluteName"])
    if err != nil {
      return nil, fmt.Errorf("error retrieving subzones from gateway: %s, %s", zoneName, err)
    }
    zones = append(zones, subZones...)
  }

  return zones, nil
}


// TODO
func (c *GatewayClient) getHostRecords(zone string, records *[]BluecatHostRecord) error {
  return nil
}

// TODO
func (c *GatewayClient) getCNAMERecords(zone string, records *[]BluecatCNAMERecord) error {
  return nil
}

//splitProperties is a helper function to break a '|' separated string into key/value pairs
// i.e. "foo=bar|baz=mop"
func splitProperties(props string) map[string]string {
  propMap := make(map[string]string)

  // remove trailing | character before we split
  props = strings.TrimSuffix(props, "|")

  splits := strings.Split(props, "|")
  for _, pair := range splits {
    items := strings.Split(pair, "=")
    propMap[items[0]] = items[1]
  }

  return propMap
}

//expandZone takes an absolute domain name such as 'example.com' and returns a zone hierarchy used by Bluecat Gateway,
//such as '/zones/com/zones/example/'
func expandZone(zone string) string {
  ze := "zones/"
  parts := strings.Split(zone, ".")
  if len(parts) > 1 {
    last := len(parts) - 1
    for i := range parts {
      ze = ze + parts[last-i] + "/zones/"
    }
  } else {
    ze = ze + zone + "/zones/"
  }
  return ze
}
