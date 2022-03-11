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

package api

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// TODO: Ensure DNS Deploy Type Defaults to no-deploy instead of ""
type BluecatConfig struct {
	GatewayHost      string `json:"gatewayHost"`
	GatewayUsername  string `json:"gatewayUsername,omitempty"`
	GatewayPassword  string `json:"gatewayPassword,omitempty"`
	DNSConfiguration string `json:"dnsConfiguration"`
	DNSServerName    string `json:"dnsServerName"`
	DNSDeployType    string `json:"dnsDeployType"`
	View             string `json:"dnsView"`
	RootZone         string `json:"rootZone"`
	SkipTLSVerify    bool   `json:"skipTLSVerify"`
}

type GatewayClient interface {
	GetBluecatZones(zoneName string) ([]BluecatZone, error)
	GetHostRecords(zone string, records *[]BluecatHostRecord) error
	GetCNAMERecords(zone string, records *[]BluecatCNAMERecord) error
	GetHostRecord(name string, record *BluecatHostRecord) error
	GetCNAMERecord(name string, record *BluecatCNAMERecord) error
	CreateHostRecord(zone string, req *BluecatCreateHostRecordRequest) (res interface{}, err error)
	CreateCNAMERecord(zone string, req *BluecatCreateCNAMERecordRequest) (res interface{}, err error)
	DeleteHostRecord(name string, zone string) (err error)
	DeleteCNAMERecord(name string, zone string) (err error)
	GetTXTRecords(zone string, records *[]BluecatTXTRecord) error
	GetTXTRecord(name string, record *BluecatTXTRecord) error
	CreateTXTRecord(zone string, req *BluecatCreateTXTRecordRequest) (res interface{}, err error)
	DeleteTXTRecord(name string, zone string) error
	ServerFullDeploy() error
}

// GatewayClientConfig defines the configuration for a Bluecat Gateway Client
type GatewayClientConfig struct {
	Cookie           http.Cookie
	Token            string
	Host             string
	DNSConfiguration string
	View             string
	RootZone         string
	DNSServerName    string
	SkipTLSVerify    bool
}

// BluecatZone defines a zone to hold records
type BluecatZone struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Properties string `json:"properties"`
	Type       string `json:"type"`
}

// BluecatHostRecord defines dns Host record
type BluecatHostRecord struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Properties string `json:"properties"`
	Type       string `json:"type"`
}

// BluecatCNAMERecord defines dns CNAME record
type BluecatCNAMERecord struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Properties string `json:"properties"`
	Type       string `json:"type"`
}

// BluecatTXTRecord defines dns TXT record
type BluecatTXTRecord struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Properties string `json:"properties"`
}

type BluecatCreateHostRecordRequest struct {
	AbsoluteName string `json:"absolute_name"`
	IP4Address   string `json:"ip4_address"`
	TTL          int    `json:"ttl"`
	Properties   string `json:"properties"`
}

type BluecatCreateCNAMERecordRequest struct {
	AbsoluteName string `json:"absolute_name"`
	LinkedRecord string `json:"linked_record"`
	TTL          int    `json:"ttl"`
	Properties   string `json:"properties"`
}

type BluecatCreateTXTRecordRequest struct {
	AbsoluteName string `json:"absolute_name"`
	Text         string `json:"txt"`
}

type BluecatServerFullDeployRequest struct {
	ServerName string `json:"server_name"`
}

// NewGatewayClient creates and returns a new Bluecat gateway client
func NewGatewayClientConfig(cookie http.Cookie, token, gatewayHost, dnsConfiguration, view, rootZone, dnsServerName string, skipTLSVerify bool) GatewayClientConfig {
	// TODO: do not handle defaulting here
	//
	// Right now the Bluecat gateway doesn't seem to have a way to get the root zone from the API. If the user
	// doesn't provide one via the config file we'll assume it's 'com'
	if rootZone == "" {
		rootZone = "com"
	}
	return GatewayClientConfig{
		Cookie:           cookie,
		Token:            token,
		Host:             gatewayHost,
		DNSConfiguration: dnsConfiguration,
		DNSServerName:    dnsServerName,
		View:             view,
		RootZone:         rootZone,
		SkipTLSVerify:    skipTLSVerify,
	}
}

// GetBluecatGatewayToken retrieves a Bluecat Gateway API token.
func GetBluecatGatewayToken(cfg BluecatConfig) (string, http.Cookie, error) {
	var username string
	if cfg.GatewayUsername != "" {
		username = cfg.GatewayUsername
	}
	if v, ok := os.LookupEnv("BLUECAT_USERNAME"); ok {
		username = v
	}

	var password string
	if cfg.GatewayPassword != "" {
		password = cfg.GatewayPassword
	}
	if v, ok := os.LookupEnv("BLUECAT_PASSWORD"); ok {
		password = v
	}

	body, err := json.Marshal(map[string]string{
		"username": username,
		"password": password,
	})
	if err != nil {
		return "", http.Cookie{}, errors.Wrap(err, "could not unmarshal credentials for bluecat gateway config")
	}

	c := newHTTPClient(cfg.SkipTLSVerify)

	resp, err := c.Post(cfg.GatewayHost+"/rest_login", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return "", http.Cookie{}, errors.Wrap(err, "error obtaining API token from bluecat gateway")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		details, _ := ioutil.ReadAll(resp.Body)
		return "", http.Cookie{}, errors.Errorf("got HTTP response code %v, detailed message: %v", resp.StatusCode, string(details))
	}

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", http.Cookie{}, errors.Wrap(err, "error reading get_token response from bluecat gateway")
	}

	resJSON := map[string]string{}
	err = json.Unmarshal(res, &resJSON)
	if err != nil {
		return "", http.Cookie{}, errors.Wrap(err, "error unmarshaling json response (auth) from bluecat gateway")
	}

	// Example response: {"access_token": "BAMAuthToken: abc123"}
	// We only care about the actual token string - i.e. abc123
	// The gateway also creates a cookie as part of the response. This seems to be the actual auth mechanism, at least
	// for now.
	return strings.Split(resJSON["access_token"], " ")[1], *resp.Cookies()[0], nil
}

func (c GatewayClientConfig) GetBluecatZones(zoneName string) ([]BluecatZone, error) {
	client := newHTTPClient(c.SkipTLSVerify)

	zonePath := expandZone(zoneName)
	url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration + "/views/" + c.View + "/" + zonePath
	req, err := c.buildHTTPRequest("GET", url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error building http request")
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "error requesting zones from gateway: %v, %v", url, zoneName)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("received http %v requesting zones from gateway in zone %v", resp.StatusCode, zoneName)
	}

	zones := []BluecatZone{}
	json.NewDecoder(resp.Body).Decode(&zones)

	// Bluecat Gateway only returns subzones one level deeper than the provided zone
	// so this recursion is needed to traverse subzones until none are returned
	for _, zone := range zones {
		zoneProps := SplitProperties(zone.Properties)
		subZones, err := c.GetBluecatZones(zoneProps["absoluteName"])
		if err != nil {
			return nil, errors.Wrapf(err, "error retrieving subzones from gateway: %v", zoneName)
		}
		zones = append(zones, subZones...)
	}

	return zones, nil
}

func (c GatewayClientConfig) GetHostRecords(zone string, records *[]BluecatHostRecord) error {
	client := newHTTPClient(c.SkipTLSVerify)

	zonePath := expandZone(zone)

	// Remove the trailing 'zones/'
	zonePath = strings.TrimSuffix(zonePath, "zones/")

	url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration + "/views/" + c.View + "/" + zonePath + "host_records/"
	req, err := c.buildHTTPRequest("GET", url, nil)
	if err != nil {
		return errors.Wrap(err, "error building http request")
	}

	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrapf(err, "error requesting host records from gateway in zone %v", zone)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.Errorf("received http %v requesting host records from gateway in zone %v", resp.StatusCode, zone)
	}

	json.NewDecoder(resp.Body).Decode(records)
	log.Debugf("Get Host Records Response: %v", records)

	return nil
}

func (c GatewayClientConfig) GetCNAMERecords(zone string, records *[]BluecatCNAMERecord) error {
	client := newHTTPClient(c.SkipTLSVerify)

	zonePath := expandZone(zone)

	// Remove the trailing 'zones/'
	zonePath = strings.TrimSuffix(zonePath, "zones/")

	url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration + "/views/" + c.View + "/" + zonePath + "cname_records/"
	req, err := c.buildHTTPRequest("GET", url, nil)
	if err != nil {
		return errors.Wrap(err, "error building http request")
	}

	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrapf(err, "error retrieving cname records from gateway in zone %v", zone)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.Errorf("received http %v requesting cname records from gateway in zone %v", resp.StatusCode, zone)
	}

	json.NewDecoder(resp.Body).Decode(records)
	log.Debugf("Get CName Records Response: %v", records)

	return nil
}

func (c GatewayClientConfig) GetTXTRecords(zone string, records *[]BluecatTXTRecord) error {
	client := newHTTPClient(c.SkipTLSVerify)

	zonePath := expandZone(zone)

	// Remove the trailing 'zones/'
	zonePath = strings.TrimSuffix(zonePath, "zones/")

	url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration + "/views/" + c.View + "/" + zonePath + "text_records/"
	req, err := c.buildHTTPRequest("GET", url, nil)
	if err != nil {
		return errors.Wrap(err, "error building http request")
	}
	log.Debugf("Request: %v", req)

	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrapf(err, "error retrieving txt records from gateway in zone %v", zone)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.Errorf("received http %v requesting txt records from gateway in zone %v", resp.StatusCode, zone)
	}

	log.Debugf("Get Txt Records response: %v", resp)
	json.NewDecoder(resp.Body).Decode(records)
	log.Debugf("Get TXT Records Body: %v", records)

	return nil
}

func (c GatewayClientConfig) GetHostRecord(name string, record *BluecatHostRecord) error {
	client := newHTTPClient(c.SkipTLSVerify)

	url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration +
		"/views/" + c.View + "/" +
		"host_records/" + name + "/"
	req, err := c.buildHTTPRequest("GET", url, nil)
	if err != nil {
		return errors.Wrapf(err, "error building http request: %v", name)
	}

	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrapf(err, "error retrieving host record %v from gateway", name)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.Errorf("received http %v while retrieving host record %v from gateway", resp.StatusCode, name)
	}

	json.NewDecoder(resp.Body).Decode(record)
	log.Debugf("Get Host Record Response: %v", record)
	return nil
}

func (c GatewayClientConfig) GetCNAMERecord(name string, record *BluecatCNAMERecord) error {
	client := newHTTPClient(c.SkipTLSVerify)

	url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration +
		"/views/" + c.View + "/" +
		"cname_records/" + name + "/"
	req, err := c.buildHTTPRequest("GET", url, nil)
	if err != nil {
		return errors.Wrapf(err, "error building http request: %v", name)
	}

	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrapf(err, "error retrieving cname record %v from gateway", name)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.Errorf("received http %v while retrieving cname record %v from gateway", resp.StatusCode, name)
	}

	json.NewDecoder(resp.Body).Decode(record)
	log.Debugf("Get CName Record Response: %v", record)
	return nil
}

func (c GatewayClientConfig) GetTXTRecord(name string, record *BluecatTXTRecord) error {
	client := newHTTPClient(c.SkipTLSVerify)

	url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration +
		"/views/" + c.View + "/" +
		"text_records/" + name + "/"

	req, err := c.buildHTTPRequest("GET", url, nil)
	if err != nil {
		return errors.Wrap(err, "error building http request")
	}

	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrapf(err, "error retrieving record %v from gateway", name)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.Errorf("received http %v while retrieving txt record %v from gateway", resp.StatusCode, name)
	}

	json.NewDecoder(resp.Body).Decode(record)
	log.Debugf("Get TXT Record Response: %v", record)

	return nil
}

func (c GatewayClientConfig) CreateHostRecord(zone string, req *BluecatCreateHostRecordRequest) (res interface{}, err error) {
	client := newHTTPClient(c.SkipTLSVerify)

	zonePath := expandZone(zone)
	// Remove the trailing 'zones/'
	zonePath = strings.TrimSuffix(zonePath, "zones/")

	url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration + "/views/" + c.View + "/" + zonePath + "host_records/"
	body, _ := json.Marshal(req)
	hreq, err := c.buildHTTPRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, errors.Wrap(err, "error building http request")
	}
	hreq.Header.Add("Content-Type", "application/json")
	res, err = client.Do(hreq)

	return
}

func (c GatewayClientConfig) CreateCNAMERecord(zone string, req *BluecatCreateCNAMERecordRequest) (res interface{}, err error) {
	client := newHTTPClient(c.SkipTLSVerify)

	zonePath := expandZone(zone)
	// Remove the trailing 'zones/'
	zonePath = strings.TrimSuffix(zonePath, "zones/")

	url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration + "/views/" + c.View + "/" + zonePath + "cname_records/"
	body, _ := json.Marshal(req)

	hreq, err := c.buildHTTPRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, errors.Wrap(err, "error building http request")
	}

	hreq.Header.Add("Content-Type", "application/json")
	res, err = client.Do(hreq)

	return
}

func (c GatewayClientConfig) CreateTXTRecord(zone string, req *BluecatCreateTXTRecordRequest) (interface{}, error) {
	client := newHTTPClient(c.SkipTLSVerify)

	zonePath := expandZone(zone)
	// Remove the trailing 'zones/'
	zonePath = strings.TrimSuffix(zonePath, "zones/")

	url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration + "/views/" + c.View + "/" + zonePath + "text_records/"
	body, _ := json.Marshal(req)
	hreq, err := c.buildHTTPRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	hreq.Header.Add("Content-Type", "application/json")
	res, err := client.Do(hreq)

	return res, err
}

func (c GatewayClientConfig) DeleteHostRecord(name string, zone string) (err error) {
	client := newHTTPClient(c.SkipTLSVerify)

	url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration +
		"/views/" + c.View + "/" +
		"host_records/" + name + "." + zone + "/"
	req, err := c.buildHTTPRequest("DELETE", url, nil)
	if err != nil {
		return errors.Wrapf(err, "error building http request: %v", name)
	}

	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrapf(err, "error deleting host record %v from gateway", name)
	}

	if resp.StatusCode != http.StatusOK {
		return errors.Errorf("received http %v while deleting host record %v from gateway", resp.StatusCode, name)
	}

	return nil
}

func (c GatewayClientConfig) DeleteCNAMERecord(name string, zone string) (err error) {
	client := newHTTPClient(c.SkipTLSVerify)

	url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration +
		"/views/" + c.View + "/" +
		"cname_records/" + name + "." + zone + "/"
	req, err := c.buildHTTPRequest("DELETE", url, nil)
	if err != nil {
		return errors.Wrapf(err, "error building http request: %v", name)
	}

	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrapf(err, "error deleting cname record %v from gateway", name)
	}

	if resp.StatusCode != http.StatusOK {
		return errors.Errorf("received http %v while deleting cname record %v from gateway", resp.StatusCode, name)
	}

	return nil
}

func (c GatewayClientConfig) DeleteTXTRecord(name string, zone string) error {
	client := newHTTPClient(c.SkipTLSVerify)

	url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration +
		"/views/" + c.View + "/" +
		"text_records/" + name + "." + zone + "/"

	req, err := c.buildHTTPRequest("DELETE", url, nil)
	if err != nil {
		return errors.Wrap(err, "error building http request")
	}

	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrapf(err, "error deleting txt record %v from gateway", name)
	}

	if resp.StatusCode != http.StatusOK {
		return errors.Errorf("received http %v while deleting txt record %v from gateway", resp.StatusCode, name)
	}

	return nil
}

func (c GatewayClientConfig) ServerFullDeploy() error {
	log.Infof("Executing full deploy on server %s", c.DNSServerName)
	httpClient := newHTTPClient(c.SkipTLSVerify)
	url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration + "/server/full_deploy/"
	requestBody := BluecatServerFullDeployRequest{
		ServerName: c.DNSServerName,
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		return errors.Wrap(err, "could not marshal body for server full deploy")
	}

	request, err := c.buildHTTPRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return errors.Wrap(err, "error building http request")
	}

	request.Header.Add("Content-Type", "application/json")
	response, err := httpClient.Do(request)
	if err != nil {
		return errors.Wrap(err, "error executing full deploy")
	}

	if response.StatusCode != http.StatusCreated {
		responseBody, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return errors.Wrap(err, "failed to read full deploy response body")
		}
		return errors.Errorf("got HTTP response code %v, detailed message: %v", response.StatusCode, string(responseBody))
	}

	return nil
}

// buildHTTPRequest builds a standard http Request and adds authentication headers required by Bluecat Gateway
func (c GatewayClientConfig) buildHTTPRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Basic "+c.Token)
	req.AddCookie(&c.Cookie)
	return req, err
}

// SplitProperties is a helper function to break a '|' separated string into key/value pairs
// i.e. "foo=bar|baz=mop"
func SplitProperties(props string) map[string]string {
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

// IsValidDNSDeployType validates the deployment type provided by a users configuration is supported by the Bluecat Provider.
func IsValidDNSDeployType(deployType string) bool {
	validDNSDeployTypes := []string{"no-deploy", "full-deploy"}
	for _, t := range validDNSDeployTypes {
		if t == deployType {
			return true
		}
	}
	return false
}

// expandZone takes an absolute domain name such as 'example.com' and returns a zone hierarchy used by Bluecat Gateway,
// such as '/zones/com/zones/example/zones/'
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

// newHTTPClient returns an instance of http client
func newHTTPClient(skipTLSVerify bool) *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: skipTLSVerify,
			},
		},
	}
}
