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

package rest

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/provider/mikrotik/common"
)

type object map[string]interface{}

type mikrotikREST struct {
	server        string
	username      string
	password      string
	ownerId       string
	domainFilter  endpoint.DomainFilter
	minimumTTL    endpoint.TTL
	dryRun        bool
	tlsSkipVerify bool
}

func NewMikrotikREST(server, username, password, ownerId string, domainFilter endpoint.DomainFilter, minimumTTL endpoint.TTL, dryRun, tlsSkipVerify bool) (common.MikrotikConnection, error) {
	r := &mikrotikREST{
		server:        server,
		username:      username,
		password:      password,
		ownerId:       ownerId,
		domainFilter:  domainFilter,
		minimumTTL:    minimumTTL,
		dryRun:        dryRun,
		tlsSkipVerify: tlsSkipVerify,
	}

	// run a quick command to test the connection
	_, err := r.ListRecords("")
	if err != nil {
		return nil, err
	}
	log.Info("Connected to Mikrotik REST server: ", server)
	return r, nil
}

func (di mikrotikREST) Disconnect() error {
	return nil
}

func (di mikrotikREST) queryArray(method, command string, args interface{}) ([]object, error) {
	result := make([]object, 0)
	body, err := di.query(method, command, args)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(body, &result)
	return result, err
}

func (di mikrotikREST) queryObject(method, command string, args interface{}) (object, error) {
	result := make(object, 0)
	body, err := di.query(method, command, args)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(body, &result)
	return result, err
}

func (di mikrotikREST) query(method, command string, args interface{}) ([]byte, error) {
	result := make([]byte, 0)

	// HTTP endpoint
	call := fmt.Sprintf("%s/rest/%s", di.server, command)
	// log.Infof("query: %s", call)

	var reader io.Reader = nil
	if args != nil {
		b, err := json.Marshal(args)
		if err != nil {
			return result, err
		}
		reader = bytes.NewReader(b)
	}

	// Create a HTTP post request
	r, err := http.NewRequest(method, call, reader)
	if err != nil {
		return result, err
	}

	r.SetBasicAuth(di.username, di.password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: di.tlsSkipVerify},
	}

	client := &http.Client{Transport: tr}
	res, err := client.Do(r)
	if err != nil {
		return result, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return result, err
	}
	if res.StatusCode >= 400 {
		return result, fmt.Errorf("command %s return code %d: %s\nArgs: %#v", command, res.StatusCode, string(data), args)
	}
	return data, nil
}

func (di mikrotikREST) objectToRecord(o object) *endpoint.Endpoint {
	result := endpoint.Endpoint{
		Targets: make(endpoint.Targets, 0),
		Labels:  make(endpoint.Labels),
	}

	if name, ok := o["name"]; ok {
		result.DNSName = name.(string)
	}

	if text, ok := o["text"]; ok {
		result.RecordType = endpoint.RecordTypeTXT
		result.Targets = append(result.Targets, text.(string))
	}

	if cname, ok := o["cname"]; ok {
		result.RecordType = endpoint.RecordTypeCNAME
		result.Targets = append(result.Targets, cname.(string))
	}

	if ttl, ok := o["ttl"]; ok {
		ttl := ttl.(string)
		ttlDuration, err := time.ParseDuration(ttl)
		if err != nil {
			log.Errorf("ttl conversion failed: %v", err)
		} else {
			result.RecordTTL = endpoint.TTL(ttlDuration.Seconds())
		}
	}

	if comment, ok := o["comment"]; ok {
		result.Labels[endpoint.OwnerLabelKey] = comment.(string)
	}

	if addr, ok := o["address"]; ok {
		_, ok := o["type"]
		if ok {
			result.RecordType = endpoint.RecordTypeAAAA
		} else {
			result.RecordType = endpoint.RecordTypeA
		}
		result.Targets = append(result.Targets, addr.(string))
	}
	return &result
}

func (di mikrotikREST) ListRecords(recordType string) ([]*endpoint.Endpoint, error) {
	result := []*endpoint.Endpoint{}

	terms := []string{}
	terms = append(terms, "comment="+di.ownerId)
	if recordType != "" {
		terms = append(terms, "type="+recordType)
		if recordType == endpoint.RecordTypeA {
			terms = append(terms, "-type")
			terms = append(terms, "#|&")
		} else {
			terms = append(terms, "#&")
		}
	}

	reply, err := di.findExistingRecord(terms)
	if err != nil {
		log.Errorf("listRecords: findExistingRecord: %v: %#v", err, reply)
		return result, err
	}

	for _, r := range reply {
		result = append(result, di.objectToRecord(r))
	}
	return result, nil
}

func (di mikrotikREST) CreateRecord(ep *endpoint.Endpoint) error {
	if di.dryRun {
		log.Infof("DRY RUN: create %s IN %s -> %s", ep.DNSName, ep.RecordType, ep.Targets[0])
		return nil
	}

	terms := []string{
		"name=" + ep.DNSName,
		"type=" + ep.RecordType,
		"#&+", // https://wiki.mikrotik.com/wiki/Manual:API#Queries
	}
	// record type "A" is the default and not part of the object!
	if ep.RecordType == endpoint.RecordTypeA {
		terms = []string{
			"name=" + ep.DNSName,
			"-type",
			"type=" + ep.RecordType,
			"#|&+",
		}
	}

	reply, err := di.findExistingRecord(terms)
	if err != nil {
		log.Errorf("update record: %v: %#v", err, reply)
		return err
	}

	typName := "address"
	if ep.RecordType == endpoint.RecordTypeA && ep.RecordType != endpoint.RecordTypeAAAA {
		typName = strings.ToLower(ep.RecordType)
	}

	ep.RecordTTL = endpoint.TTL(math.Max(float64(ep.RecordTTL), float64(di.minimumTTL)))

	// update record
	_, err = di.queryObject("PUT", "ip/dns/static",
		object{
			"name":    ep.DNSName,
			typName:   ep.Targets[0],
			"ttl":     ep.RecordTTL,
			"comment": di.ownerId,
		},
	)
	return err
}

func (di mikrotikREST) DeleteRecord(ep *endpoint.Endpoint) error {
	if di.dryRun {
		log.Infof("DRY RUN: delete %s IN %s -> %s", ep.DNSName, ep.RecordType, ep.Targets[0])
		return nil
	}

	terms := []string{}
	terms = append(terms, "comment="+di.ownerId)
	terms = append(terms, "type="+ep.RecordType)
	if ep.RecordType == endpoint.RecordTypeA {
		terms = append(terms, "-type")
		terms = append(terms, "#|&")
	} else {
		terms = append(terms, "#&")
	}

	reply, err := di.findExistingRecord(terms)
	if err != nil {
		log.Errorf("listRecords: findExistingRecord: %v: %#v", err, reply)
		return err
	}

	id := reply[0][".id"].(string)

	_, err = di.queryObject("DELETE", "ip/dns/static/"+id, nil)
	if err != nil {
		return err
	}
	return nil
}

func (di mikrotikREST) findExistingRecord(terms []string) ([]object, error) {
	result, err := di.queryArray("POST", "ip/dns/static/print",
		object{".query": terms},
	)
	if err != nil {
		return []object{}, err
	}

	// klog.Infof("Result: %#v", result)
	return result, nil
}
