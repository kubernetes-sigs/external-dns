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
// TODO: add logging
// TODO: add timeouts
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
	CreateHostRecord(zone string, req *BluecatCreateHostRecordRequest) error
	CreateCNAMERecord(zone string, req *BluecatCreateCNAMERecordRequest) error
	DeleteHostRecord(name string, zone string) (err error)
	DeleteCNAMERecord(name string, zone string) (err error)
	GetTXTRecords(zone string, records *[]BluecatTXTRecord) error
	GetTXTRecord(name string, record *BluecatTXTRecord) error
	CreateTXTRecord(zone string, req *BluecatCreateTXTRecordRequest) error
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
	url := cfg.GatewayHost + "/rest_login"

	response, err := executeHTTPRequest(cfg.SkipTLSVerify, http.MethodPost, url, "", bytes.NewBuffer(body), http.Cookie{})
	if err != nil {
		return "", http.Cookie{}, errors.Wrap(err, "error obtaining API token from bluecat gateway")
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", http.Cookie{}, errors.Wrap(err, "failed to read login response from bluecat gateway")
	}

	if response.StatusCode != http.StatusOK {
		return "", http.Cookie{}, errors.Errorf("got HTTP response code %v, detailed message: %v", response.StatusCode, string(responseBody))
	}

	jsonResponse := map[string]string{}
	err = json.Unmarshal(responseBody, &jsonResponse)
	if err != nil {
		return "", http.Cookie{}, errors.Wrap(err, "error unmarshaling json response (auth) from bluecat gateway")
	}

	// Example response: {"access_token": "BAMAuthToken: abc123"}
	// We only care about the actual token string - i.e. abc123
	// The gateway also creates a cookie as part of the response. This seems to be the actual auth mechanism, at least
	// for now.
	return strings.Split(jsonResponse["access_token"], " ")[1], *response.Cookies()[0], nil
}

func (c GatewayClientConfig) GetBluecatZones(zoneName string) ([]BluecatZone, error) {
	zonePath := expandZone(zoneName)
	url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration + "/views/" + c.View + "/" + zonePath

	response, err := executeHTTPRequest(c.SkipTLSVerify, http.MethodGet, url, c.Token, nil, c.Cookie)
	if err != nil {
		return nil, errors.Wrapf(err, "error requesting zones from gateway: %v, %v", url, zoneName)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, errors.Errorf("received http %v requesting zones from gateway in zone %v", response.StatusCode, zoneName)
	}

	zones := []BluecatZone{}
	json.NewDecoder(response.Body).Decode(&zones)

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
	zonePath := expandZone(zone)
	// Remove the trailing 'zones/'
	zonePath = strings.TrimSuffix(zonePath, "zones/")

	url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration + "/views/" + c.View + "/" + zonePath + "host_records/"

	response, err := executeHTTPRequest(c.SkipTLSVerify, http.MethodGet, url, c.Token, nil, c.Cookie)
	if err != nil {
		return errors.Wrapf(err, "error requesting host records from gateway in zone %v", zone)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return errors.Errorf("received http %v requesting host records from gateway in zone %v", response.StatusCode, zone)
	}

	json.NewDecoder(response.Body).Decode(records)
	log.Debugf("Get Host Records Response: %v", records)

	return nil
}

func (c GatewayClientConfig) GetCNAMERecords(zone string, records *[]BluecatCNAMERecord) error {
	zonePath := expandZone(zone)
	// Remove the trailing 'zones/'
	zonePath = strings.TrimSuffix(zonePath, "zones/")

	url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration + "/views/" + c.View + "/" + zonePath + "cname_records/"

	response, err := executeHTTPRequest(c.SkipTLSVerify, http.MethodGet, url, c.Token, nil, c.Cookie)
	if err != nil {
		return errors.Wrapf(err, "error retrieving cname records from gateway in zone %v", zone)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return errors.Errorf("received http %v requesting cname records from gateway in zone %v", response.StatusCode, zone)
	}

	json.NewDecoder(response.Body).Decode(records)
	log.Debugf("Get CName Records Response: %v", records)

	return nil
}

func (c GatewayClientConfig) GetTXTRecords(zone string, records *[]BluecatTXTRecord) error {
	zonePath := expandZone(zone)
	// Remove the trailing 'zones/'
	zonePath = strings.TrimSuffix(zonePath, "zones/")

	url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration + "/views/" + c.View + "/" + zonePath + "text_records/"

	response, err := executeHTTPRequest(c.SkipTLSVerify, http.MethodGet, url, c.Token, nil, c.Cookie)
	if err != nil {
		return errors.Wrapf(err, "error retrieving txt records from gateway in zone %v", zone)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return errors.Errorf("received http %v requesting txt records from gateway in zone %v", response.StatusCode, zone)
	}

	log.Debugf("Get Txt Records response: %v", response)
	json.NewDecoder(response.Body).Decode(records)
	log.Debugf("Get TXT Records Body: %v", records)

	return nil
}

func (c GatewayClientConfig) GetHostRecord(name string, record *BluecatHostRecord) error {
	url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration +
		"/views/" + c.View + "/" +
		"host_records/" + name + "/"

	response, err := executeHTTPRequest(c.SkipTLSVerify, http.MethodGet, url, c.Token, nil, c.Cookie)
	if err != nil {
		return errors.Wrapf(err, "error retrieving host record %v from gateway", name)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return errors.Errorf("received http %v while retrieving host record %v from gateway", response.StatusCode, name)
	}

	json.NewDecoder(response.Body).Decode(record)
	log.Debugf("Get Host Record Response: %v", record)
	return nil
}

func (c GatewayClientConfig) GetCNAMERecord(name string, record *BluecatCNAMERecord) error {
	url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration +
		"/views/" + c.View + "/" +
		"cname_records/" + name + "/"

	response, err := executeHTTPRequest(c.SkipTLSVerify, http.MethodGet, url, c.Token, nil, c.Cookie)
	if err != nil {
		return errors.Wrapf(err, "error retrieving cname record %v from gateway", name)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return errors.Errorf("received http %v while retrieving cname record %v from gateway", response.StatusCode, name)
	}

	json.NewDecoder(response.Body).Decode(record)
	log.Debugf("Get CName Record Response: %v", record)
	return nil
}

func (c GatewayClientConfig) GetTXTRecord(name string, record *BluecatTXTRecord) error {
	url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration +
		"/views/" + c.View + "/" +
		"text_records/" + name + "/"

	response, err := executeHTTPRequest(c.SkipTLSVerify, http.MethodGet, url, c.Token, nil, c.Cookie)
	if err != nil {
		return errors.Wrapf(err, "error retrieving record %v from gateway", name)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return errors.Errorf("received http %v while retrieving txt record %v from gateway", response.StatusCode, name)
	}

	json.NewDecoder(response.Body).Decode(record)
	log.Debugf("Get TXT Record Response: %v", record)

	return nil
}

func (c GatewayClientConfig) CreateHostRecord(zone string, req *BluecatCreateHostRecordRequest) error {
	zonePath := expandZone(zone)
	// Remove the trailing 'zones/'
	zonePath = strings.TrimSuffix(zonePath, "zones/")

	url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration + "/views/" + c.View + "/" + zonePath + "host_records/"
	body, err := json.Marshal(req)
	if err != nil {
		return errors.Wrap(err, "could not marshal body for create host record")
	}

	response, err := executeHTTPRequest(c.SkipTLSVerify, http.MethodPost, url, c.Token, bytes.NewBuffer(body), c.Cookie)
	if err != nil {
		return errors.Wrapf(err, "error creating host record %v in gateway", req.AbsoluteName)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusCreated {
		return errors.Errorf("received http %v while creating host record %v in gateway", response.StatusCode, req.AbsoluteName)
	}

	return nil
}

func (c GatewayClientConfig) CreateCNAMERecord(zone string, req *BluecatCreateCNAMERecordRequest) error {
	zonePath := expandZone(zone)
	// Remove the trailing 'zones/'
	zonePath = strings.TrimSuffix(zonePath, "zones/")

	url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration + "/views/" + c.View + "/" + zonePath + "cname_records/"
	body, err := json.Marshal(req)
	if err != nil {
		return errors.Wrap(err, "could not marshal body for create cname record")
	}

	response, err := executeHTTPRequest(c.SkipTLSVerify, http.MethodPost, url, c.Token, bytes.NewBuffer(body), c.Cookie)
	if err != nil {
		return errors.Wrapf(err, "error creating cname record %v in gateway", req.AbsoluteName)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusCreated {
		return errors.Errorf("received http %v while creating cname record %v to alias %v in gateway", response.StatusCode, req.AbsoluteName, req.LinkedRecord)
	}

	return nil
}

func (c GatewayClientConfig) CreateTXTRecord(zone string, req *BluecatCreateTXTRecordRequest) error {
	zonePath := expandZone(zone)
	// Remove the trailing 'zones/'
	zonePath = strings.TrimSuffix(zonePath, "zones/")

	url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration + "/views/" + c.View + "/" + zonePath + "text_records/"
	body, err := json.Marshal(req)
	if err != nil {
		return errors.Wrap(err, "could not marshal body for create txt record")
	}

	response, err := executeHTTPRequest(c.SkipTLSVerify, http.MethodPost, url, c.Token, bytes.NewBuffer(body), c.Cookie)
	if err != nil {
		return errors.Wrapf(err, "error creating txt record %v in gateway", req.AbsoluteName)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusCreated {
		return errors.Errorf("received http %v while creating txt record %v in gateway", response.StatusCode, req.AbsoluteName)
	}

	return nil
}

func (c GatewayClientConfig) DeleteHostRecord(name string, zone string) (err error) {
	url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration +
		"/views/" + c.View + "/" +
		"host_records/" + name + "." + zone + "/"

	response, err := executeHTTPRequest(c.SkipTLSVerify, http.MethodDelete, url, c.Token, nil, c.Cookie)
	if err != nil {
		return errors.Wrapf(err, "error deleting host record %v from gateway", name)
	}

	if response.StatusCode != http.StatusNoContent {
		return errors.Errorf("received http %v while deleting host record %v from gateway", response.StatusCode, name)
	}

	return nil
}

func (c GatewayClientConfig) DeleteCNAMERecord(name string, zone string) (err error) {
	url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration +
		"/views/" + c.View + "/" +
		"cname_records/" + name + "." + zone + "/"

	response, err := executeHTTPRequest(c.SkipTLSVerify, http.MethodDelete, url, c.Token, nil, c.Cookie)
	if err != nil {
		return errors.Wrapf(err, "error deleting cname record %v from gateway", name)
	}
	if response.StatusCode != http.StatusNoContent {
		return errors.Errorf("received http %v while deleting cname record %v from gateway", response.StatusCode, name)
	}

	return nil
}

func (c GatewayClientConfig) DeleteTXTRecord(name string, zone string) error {
	url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration +
		"/views/" + c.View + "/" +
		"text_records/" + name + "." + zone + "/"

	response, err := executeHTTPRequest(c.SkipTLSVerify, http.MethodDelete, url, c.Token, nil, c.Cookie)
	if err != nil {
		return errors.Wrapf(err, "error deleting txt record %v from gateway", name)
	}
	if response.StatusCode != http.StatusNoContent {
		return errors.Errorf("received http %v while deleting txt record %v from gateway", response.StatusCode, name)
	}

	return nil
}

func (c GatewayClientConfig) ServerFullDeploy() error {
	log.Infof("Executing full deploy on server %s", c.DNSServerName)
	url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration + "/server/full_deploy/"
	requestBody := BluecatServerFullDeployRequest{
		ServerName: c.DNSServerName,
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		return errors.Wrap(err, "could not marshal body for server full deploy")
	}

	response, err := executeHTTPRequest(c.SkipTLSVerify, http.MethodPost, url, c.Token, bytes.NewBuffer(body), c.Cookie)
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

func executeHTTPRequest(skipTLSVerify bool, method, url, token string, body io.Reader, cookie http.Cookie) (*http.Response, error) {
	httpClient := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: skipTLSVerify,
			},
		},
	}
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	if request.Method == http.MethodPost {
		request.Header.Add("Content-Type", "application/json")
	}
	request.Header.Add("Accept", "application/json")

	if token != "" {
		request.Header.Add("Authorization", "Basic "+token)
	}
	request.AddCookie(&cookie)

	return httpClient.Do(request)
}
