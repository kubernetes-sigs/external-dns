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
  Id         int    `json:"id"`
  Name       string `json:"name"`
  Properties string `json:"properties"`
  Type       string `json:"type"`
}

type BluecatCNAMERecord struct {
  Id         int    `json:"id"`
  Name       string `json:"name"`
  Properties string `json:"properties"`
  Type       string `json:"type"`
}

//TODO verify model once implemented by Bluecat Gateway
type BluecatTXTRecord struct {
  Id   int `json:"id"`
  Name string `json:"name"`
  Text string `json:"text"`
}

type bluecatRecordSet struct {
  obj interface{}
  res interface{}
}

type bluecatCreateHostRecordRequest struct {
  Name       string `json:"absolute_name"`
  Ip4Address string `json:"ip4_address"`
}

type bluecatCreateCNAMERecordRequest struct {
  Name         string `json:"absolute_name"`
  LinkedRecord string `json:"linked_record"`
}

// TODO Verify request model once implemented by Bluecat Gateway
type bluecatCreateTXTRecordRequest struct {
  Name string `json:"name"`
  Text string `json:"text"`
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
      propMap := splitProperties(rec.Properties)
      ips := strings.Split(propMap["addresses"], ",")
      for _, ip := range ips {
        ep := endpoint.NewEndpoint(propMap["absoluteName"], endpoint.RecordTypeA, ip)
        endpoints = append(endpoints, ep)
      }
    }

    var resC []BluecatCNAMERecord
    err = p.gatewayClient.getCNAMERecords(zone, &resC)
    if err != nil {
      return nil, fmt.Errorf("could not fetch CNAME records for zone: '%s': %s", zone, err)
    }
    for _, rec := range resC {
      propMap := splitProperties(rec.Properties)
      endpoints = append(endpoints, endpoint.NewEndpoint(propMap["absoluteName"], endpoint.RecordTypeCNAME, propMap["linkedRecordName"]))
    }

    // TODO not yet implemented by Bluecat Gateway
    /*
    var resT []BluecatTXTRecord
    err = p.gatewayClient.getTXTRecords(zone, &resT)
    if err != nil {
      return nil, fmt.Errorf("could not fetch TXT records for zone: '#{zone}': #{err}")
    }
    for _, rec := range resT {
      endpoints = append(endpoints, endpoint.NewEndpoint(rec.Name, endpoint.RecordTypeTXT, rec.Text))
    }
    */
  }

  log.Debugf("fetched %d records from Bluecat", len(endpoints))
  return endpoints, nil
}

func (p *BluecatProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
  zones, err := p.zones()
  if err != nil {
    return err
  }
  log.Infof("zones is: %+v\n", zones)
  log.Infof("changes: %+v\n", changes)
  created, deleted := p.mapChanges(zones, changes)
  log.Infof("created: %+v\n", created)
  log.Infof("deleted: %+v\n", deleted)
  p.deleteRecords(deleted)
  p.createRecords(created)

  return nil
}

type bluecatChangeMap map[string][]*endpoint.Endpoint

func (p *BluecatProvider) mapChanges(zones []string, changes *plan.Changes) (bluecatChangeMap, bluecatChangeMap) {
  created := bluecatChangeMap{}
  deleted := bluecatChangeMap{}

  mapChange := func(changeMap bluecatChangeMap, change *endpoint.Endpoint) {
    zone := p.findZone(zones, change.DNSName)
    if zone == "" {
      log.Debugf("ignoring changes to '%s' because a suitable Bluecat DNS zone was not found", change.DNSName)
      return
    }
    changeMap[zone] = append(changeMap[zone], change)
  }

  for _, change := range changes.Delete {
    mapChange(deleted, change)
  }
  for _, change := range changes.UpdateOld {
    mapChange(deleted, change)
  }
  for _, change := range changes.Create {
    mapChange(created, change)
  }
  for _, change := range changes.UpdateNew {
    mapChange(created, change)
  }

  return created, deleted
}

// findZone finds the most specific matching zone for a given record 'name' from a list of all zones
func (p *BluecatProvider) findZone(zones []string, name string) string {
  var result string

  for _, zone := range zones {
    if strings.HasSuffix(name, "."+zone) {
      if result == "" || len(zone) > len(result) {
        result = zone
      }
    } else if strings.EqualFold(name, zone) {
      if result == "" || len(zone) > len(result) {
        result = zone
      }
    }
  }

  return result
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
  log.Debugf("found %d zones", len(zones))
  return zones, nil
}

func (p *BluecatProvider) createRecords(created bluecatChangeMap) {
  for zone, endpoints := range created {
    for _, ep := range endpoints {
      if p.dryRun {
        log.Infof("would create %s record named '%s' to '%s' for Bluecat DNS zone '%s'.",
          ep.RecordType,
          ep.DNSName,
          ep.Targets,
          zone,
        )
        continue
      }

      log.Infof("creating %s record named '%s' to '%s' for Bluecat DNS zone '%s'.",
        ep.RecordType,
        ep.DNSName,
        ep.Targets,
        zone,
      )

      recordSet, err := p.recordSet(ep, false)
      if err != nil {
        log.Errorf(
          "Failed to retrieve %s record named '%s' to '%s' for Bluecat DNS zone '%s': %v",
          ep.RecordType,
          ep.DNSName,
          ep.Targets,
          zone,
          err,
        )
        continue
      }

      switch ep.RecordType {
      case endpoint.RecordTypeA:
        _, err = p.gatewayClient.createHostRecord(zone, recordSet.obj.(*bluecatCreateHostRecordRequest))
      case endpoint.RecordTypeCNAME:
        _, err = p.gatewayClient.createCNAMERecord(zone, recordSet.obj.(*bluecatCreateCNAMERecordRequest))
      // TODO TXT CRUD not yet implemented by Bluecat Gateway
      //case endpoint.RecordTypeTXT:
      //  _, err = p.gatewayClient.createTXTRecord(zone, recordSet.obj.(*bluecatCreateTXTRecordRequest))
      }
      if err != nil {
        log.Errorf(
          "Failed to create %s record named '%s' to '%s' for Bluecat DNS zone '%s': %v",
          ep.RecordType,
          ep.DNSName,
          ep.Targets,
          zone,
          err,
        )
      }
    }
  }
}

func (p *BluecatProvider) deleteRecords(deleted bluecatChangeMap) {
  // run deletions first
  for zone, endpoints := range deleted {
    for _, ep := range endpoints {
      if p.dryRun {
        log.Infof("would delete %s record named '%s' for Bluecat DNS zone '%s'.",
          ep.RecordType,
          ep.DNSName,
          zone,
        )
        continue
      } else {
        log.Infof("deleting %s record named '%s' for Bluecat DNS zone '%s'.",
          ep.RecordType,
          ep.DNSName,
          zone,
        )

        recordSet, err := p.recordSet(ep, true)
        if err != nil {
          log.Errorf(
            "Failed to retrieve %s record named '%s' to '%s' for Bluecat DNS zone '%s': %v",
            ep.RecordType,
            ep.DNSName,
            ep.Targets,
            zone,
            err,
          )
          continue
        }

        switch ep.RecordType {
        case endpoint.RecordTypeA:
          for _, record := range *recordSet.res.(*[]BluecatHostRecord) {
            err = p.gatewayClient.deleteHostRecord(record.Name)
          }
        case endpoint.RecordTypeCNAME:
          for _, record := range *recordSet.res.(*[]BluecatCNAMERecord) {
            err = p.gatewayClient.deleteCNAMERecord(record.Name)
          }
        // TODO Not yet implemented by Bluecat Gateway
        //case endpoint.RecordTypeTXT:
        //  for _, record := range *recordSet.res.(*[]BluecatTXTRecord) {
        //    err = p.gatewayClient.deleteTXTRecord(record.Name)
        //  }
        }
        if err != nil {
          log.Errorf("Failed to delete %s record named '%s' for Bluecat DNS zone '%s': %v",
            ep.RecordType,
            ep.DNSName,
            zone,
            err,)
        }
      }
    }
  }
}

func (p *BluecatProvider) recordSet(ep *endpoint.Endpoint, getObject bool) (recordSet bluecatRecordSet, err error) {
  switch ep.RecordType {
  case endpoint.RecordTypeA:
    nameSplit := strings.Split(ep.DNSName, ".")
    var res []BluecatHostRecord
    obj := bluecatCreateHostRecordRequest{
      Name:       nameSplit[0],
      Ip4Address: ep.Targets[0],
    }
    if getObject {
      var record BluecatHostRecord
      err = p.gatewayClient.getHostRecord(ep.DNSName, &record)
      if err != nil {
        return
      }
      res = append(res, record)
    }
    recordSet = bluecatRecordSet{
      obj: obj,
      res: &res,
    }
  case endpoint.RecordTypeCNAME:
    var res []BluecatCNAMERecord
    obj := bluecatCreateCNAMERecordRequest{
      Name:         ep.DNSName,
      LinkedRecord: ep.Targets[0],
    }
    if getObject {
      var record BluecatCNAMERecord
      err = p.gatewayClient.getCNAMERecord(ep.DNSName, &record)
      if err != nil {
        return
      }
      res = append(res, record)
    }
    recordSet = bluecatRecordSet{
      obj: obj,
      res: &res,
    }
  //TODO Not yet implemented by Bluecat Gateway
  //case endpoint.RecordTypeTXT:
  //  var res []BluecatTXTRecord
  //  obj := bluecatCreateTXTRecordRequest{
  //    Name: ep.DNSName,
  //    Text: ep.Targets[0],
  //  }
  //  if getObject {
  //    var record BluecatTXTRecord
  //    err = p.gatewayClient.getTXTRecord(ep.DNSName, &record)
  //    if err != nil {
  //      return
  //    }
  //    res = append(res, record)
  //  }
  //  recordSet = bluecatRecordSet{
  //    obj: obj,
  //    res: &res,
  //  }
  }
  return
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

  zonePath := expandZone(zoneName)
  url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration + "/views/" + c.View + "/" + zonePath
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

  // Bluecat Gateway only returns subzones one level deeper than the provided zone
  // so this recursion is needed to traverse subzones until none are returned
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

func (c *GatewayClient) getHostRecords(zone string, records *[]BluecatHostRecord) error {
  transportCfg := &http.Transport{
    TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //ignore self-signed SSL cert check
  }
  client := &http.Client{
    Transport: transportCfg,
  }

  zonePath := expandZone(zone)

  // Remove the trailing 'zones/'
  zonePath = strings.TrimSuffix(zonePath, "zones/")

  url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration + "/views/" + c.View + "/" + zonePath + "host_records/"
  req, err := http.NewRequest("GET", url, nil)
  req.Header.Add("Accept", "application/json")
  req.Header.Add("Authorization", "Basic " + c.Token)
  req.AddCookie(&c.Cookie)

  resp, err := client.Do(req)
  if err != nil {
    return fmt.Errorf("error retrieving record(s) from gateway: %s, %s", zone, err)
  }

  defer resp.Body.Close()

  json.NewDecoder(resp.Body).Decode(records)
  return nil
}

func (c *GatewayClient) getCNAMERecords(zone string, records *[]BluecatCNAMERecord) error {
  transportCfg := &http.Transport{
    TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //ignore self-signed SSL cert check
  }
  client := &http.Client{
    Transport: transportCfg,
  }

  zonePath := expandZone(zone)

  // Remove the trailing 'zones/'
  zonePath = strings.TrimSuffix(zonePath, "zones/")

  url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration + "/views/" + c.View + "/" + zonePath + "cname_records/"
  req, err := http.NewRequest("GET", url, nil)
  req.Header.Add("Accept", "application/json")
  req.Header.Add("Authorization", "Basic " + c.Token)
  req.AddCookie(&c.Cookie)

  resp, err := client.Do(req)
  if err != nil {
    return fmt.Errorf("error retrieving record(s) from gateway: %s, %s", zone, err)
  }

  defer resp.Body.Close()

  json.NewDecoder(resp.Body).Decode(records)
  return nil
}

/* TODO not yet implemented by Bluecat Gateway
func (c *GatewayClient) getTXTRecords(zone string, records *[]BluecatTXTRecord) error {
  transportCfg := &http.Transport{
    TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //ignore self-signed SSL cert check
  }
  client := &http.Client{
    Transport: transportCfg,
  }

  zonePath := expandZone(zone)

  // Remove the trailing 'zones/'
  zonePath = strings.TrimSuffix(zonePath, "zones/")

  // TODO Verify correct route once implemented by Bluecat Gateway
  url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration + "/views/" + c.View + "/" + zonePath + "txt_records/"
  req, err := http.NewRequest("GET", url, nil)
  req.Header.Add("Accept", "application/json")
  req.Header.Add("Authorization", "Basic " + c.Token)
  req.AddCookie(&c.Cookie)

  resp, err := client.Do(req)
  if err != nil {
    return fmt.Errorf("error retrieving record(s) from gateway: %s, %s", zone, err)
  }

  defer resp.Body.Close()

  json.NewDecoder(resp.Body).Decode(records)
  return nil
}*/

func (c *GatewayClient) getHostRecord(name string, record *BluecatHostRecord) error {
  transportCfg := &http.Transport{
    TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //ignore self-signed SSL cert check
  }
  client := &http.Client{
    Transport: transportCfg,
  }

  url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration +
                  "/views/" + c.View + "/" +
                  "host_records/" + name + "/"
  req, err := http.NewRequest("GET", url, nil)
  req.Header.Add("Accept", "application/json")
  req.Header.Add("Authorization", "Basic " + c.Token)
  req.AddCookie(&c.Cookie)

  resp, err := client.Do(req)
  if err != nil {
    return fmt.Errorf("error retrieving record(s) from gateway: %s, %s", name, err)
  }

  defer resp.Body.Close()

  json.NewDecoder(resp.Body).Decode(record)
  return nil
}

func (c *GatewayClient) getCNAMERecord(name string, record *BluecatCNAMERecord) error {
  transportCfg := &http.Transport{
    TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //ignore self-signed SSL cert check
  }
  client := &http.Client{
    Transport: transportCfg,
  }

  url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration +
      "/views/" + c.View + "/" +
      "cname_records/" + name + "/"
  req, err := http.NewRequest("GET", url, nil)
  req.Header.Add("Accept", "application/json")
  req.Header.Add("Authorization", "Basic " + c.Token)
  req.AddCookie(&c.Cookie)

  resp, err := client.Do(req)
  if err != nil {
    return fmt.Errorf("error retrieving record(s) from gateway: %s, %s", name, err)
  }

  defer resp.Body.Close()

  json.NewDecoder(resp.Body).Decode(record)
  return nil
}

/* TODO Not yet implemented by Bluecat Gateway
func (c *GatewayClient) getTXTRecord(name string, record *BluecatTXTRecord) error {
  transportCfg := &http.Transport{
    TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //ignore self-signed SSL cert check
  }
  client := &http.Client{
    Transport: transportCfg,
  }

  // TODO verify correct route once implemented by Bluecat Gateway
  url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration +
      "/views/" + c.View + "/" +
      "txt_records/" + name + "/"
  req, err := http.NewRequest("GET", url, nil)
  req.Header.Add("Accept", "application/json")
  req.Header.Add("Authorization", "Basic " + c.Token)
  req.AddCookie(&c.Cookie)

  resp, err := client.Do(req)
  if err != nil {
    return fmt.Errorf("error retrieving record(s) from gateway: %s, %s", name, err)
  }

  defer resp.Body.Close()

  json.NewDecoder(resp.Body).Decode(record)
  return nil
}*/

func (c *GatewayClient) createHostRecord(zone string, req *bluecatCreateHostRecordRequest) (res interface{}, err error) {
  transportCfg := &http.Transport{
    TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //ignore self-signed SSL cert check
  }
  client := &http.Client{
    Transport: transportCfg,
  }

  zonePath := expandZone(zone)
  // Remove the trailing 'zones/'
  zonePath = strings.TrimSuffix(zonePath, "zones/")

  url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration + "/views/" + c.View + "/" + zonePath + "host_records/"
  body, _ := json.Marshal(req)

  hreq, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
  hreq.Header.Add("Accept", "application/json")
  hreq.Header.Add("Authorization", "Basic " + c.Token)
  hreq.AddCookie(&c.Cookie)

  res, err = client.Do(hreq)

  return
}

func (c *GatewayClient) createCNAMERecord(zone string, req *bluecatCreateCNAMERecordRequest) (res interface{}, err error) {
  transportCfg := &http.Transport{
    TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //ignore self-signed SSL cert check
  }
  client := &http.Client{
    Transport: transportCfg,
  }

  zonePath := expandZone(zone)
  // Remove the trailing 'zones/'
  zonePath = strings.TrimSuffix(zonePath, "zones/")

  url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration + "/views/" + c.View + "/" + zonePath + "cname_records/"
  body, _ := json.Marshal(req)

  hreq, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
  hreq.Header.Add("Accept", "application/json")
  hreq.Header.Add("Authorization", "Basic " + c.Token)
  hreq.AddCookie(&c.Cookie)

  res, err = client.Do(hreq)

  return
}

/* TODO Not yet implemented by Bluecat Gateway
func (c *GatewayClient) createTXTRecord(zone string, req *bluecatCreateTXTRecordRequest) (res interface{}, err error) {
  transportCfg := &http.Transport{
    TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //ignore self-signed SSL cert check
  }
  client := &http.Client{
    Transport: transportCfg,
  }

  zonePath := expandZone(zone)
  // Remove the trailing 'zones/'
  zonePath = strings.TrimSuffix(zonePath, "zones/")

  // TODO verify correct route once implemented by Bluecat Gateway
  url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration + "/views/" + c.View + "/" + zonePath + "txt_records/"
  body, _ := json.Marshal(req)

  hreq, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
  hreq.Header.Add("Accept", "application/json")
  hreq.Header.Add("Authorization", "Basic " + c.Token)
  hreq.AddCookie(&c.Cookie)

  res, err = client.Do(hreq)

  return
}*/

func (c *GatewayClient) deleteHostRecord(name string) error {
  transportCfg := &http.Transport{
    TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //ignore self-signed SSL cert check
  }
  client := &http.Client{
    Transport: transportCfg,
  }

  url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration +
      "/views/" + c.View + "/" +
      "host_records/" + name + "/"
  req, err := http.NewRequest("DELETE", url, nil)
  req.Header.Add("Accept", "application/json")
  req.Header.Add("Authorization", "Basic " + c.Token)
  req.AddCookie(&c.Cookie)

  _, err = client.Do(req)
  if err != nil {
    return fmt.Errorf("error deleting record(s) from gateway: %s, %s", name, err)
  }

  return nil
}

func (c *GatewayClient) deleteCNAMERecord(name string) error {
  transportCfg := &http.Transport{
    TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //ignore self-signed SSL cert check
  }
  client := &http.Client{
    Transport: transportCfg,
  }

  url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration +
      "/views/" + c.View + "/" +
      "cname_records/" + name + "/"
  req, err := http.NewRequest("DELETE", url, nil)
  req.Header.Add("Accept", "application/json")
  req.Header.Add("Authorization", "Basic " + c.Token)
  req.AddCookie(&c.Cookie)

  _, err = client.Do(req)
  if err != nil {
    return fmt.Errorf("error deleting record(s) from gateway: %s, %s", name, err)
  }

  return nil
}

/* TODO Not yet implemented by Bluecat Gateway
func (c *GatewayClient) deleteTXTRecord(name string) error {
  transportCfg := &http.Transport{
    TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //ignore self-signed SSL cert check
  }
  client := &http.Client{
    Transport: transportCfg,
  }

  // TODO Verify correct route once implemented by Bluecat Gateway
  url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration +
      "/views/" + c.View + "/" +
      "txt_records/" + name + "/"
  req, err := http.NewRequest("DELETE", url, nil)
  req.Header.Add("Accept", "application/json")
  req.Header.Add("Authorization", "Basic " + c.Token)
  req.AddCookie(&c.Cookie)

  _, err = client.Do(req)
  if err != nil {
    return fmt.Errorf("error deleting record(s) from gateway: %s, %s", name, err)
  }

  return nil
}*/

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
//such as '/zones/com/zones/example/zones/'
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
