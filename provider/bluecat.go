package provider

import (
  "bytes"
  "context"
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "strings"

  log "github.com/sirupsen/logrus"

  yaml "gopkg.in/yaml.v2"

  "github.com/kubernetes-sigs/external-dns/endpoint"
  "github.com/kubernetes-sigs/external-dns/plan"
)

type bluecatConfig struct {
  GatewayHost     string
  GatewayUsername string
  GatewayPassword string
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
  Token string
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

  gatewayClient := NewGatewayClient(token)

  provider := &BluecatProvider{
    domainFilter:  domainFilter,
    zoneIdFilter:  zoneIDFilter,
    dryRun:        dryRun,
    gatewayClient: gatewayClient,
  }
  return provider, nil
}

func NewGatewayClient(token string) GatewayClient {
  return GatewayClient{Token: token}
}

func (p *BluecatProvider) Records() ([]*endpoint.Endpoint, error) {
  return []*endpoint.Endpoint{}, nil
}

func (p *BluecatProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
  return nil
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