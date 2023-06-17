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

package api

import (
	"errors"
	"fmt"
	"math"
	"net"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/provider/mikrotik/common"
)

type Reply struct {
	Line  string
	Attrs map[string]string
}

type router struct {
	connection   net.Conn
	domainFilter endpoint.DomainFilter
	ownerId      string
	minimumTTL   endpoint.TTL
	dryRun       bool
}

// endpoint must be in the form "hostname_or_ip:port"
func NewMikrotikAPI(server, username, password, ownerId string, domainFilter endpoint.DomainFilter, minimumTTL endpoint.TTL, dryRun bool) (common.MikrotikConnection, error) {
	r := &router{
		domainFilter: domainFilter,
		ownerId:      ownerId,
		minimumTTL:   minimumTTL,
		dryRun:       dryRun,
	}
	var err error

	log.Info("Connecting to Mikrotik API server: ", username, "@", server)
	r.connection, err = net.Dial("tcp", server)
	if err != nil {
		return nil, err
	}
	log.Info("Connected")

	return r, r.login(username, password)
}

func (m *router) Disconnect() error {
	if m.connection != nil {
		err := m.connection.Close()
		m.connection = nil
		return err
	}
	return nil
}

func (m router) login(username, password string) error {
	err := m.send(
		"/login",
		"=name="+username,
		"=password="+password,
	)
	if err != nil {
		return err
	}

	_, err = m.receive()
	// log.Printf("%#v", reply)
	if err != nil {
		return err
	}

	return nil
}

func (m router) send(sentence ...string) error {
	if m.connection == nil {
		return errors.New("not connected")
	}

	err := m.writeSentence(sentence)
	if err != nil {
		return err
	}

	return nil
}

func (m router) receive() ([]Reply, error) {
	var result error
	reply := []Reply{}
	for {
		sentence, err := m.readSentence()
		if err != nil {
			return reply, err
		}
		if len(sentence) == 0 {
			continue
		}

		r := Reply{
			Line:  sentence[0],
			Attrs: map[string]string{},
		}

		for _, a := range sentence[1:] {
			i := strings.LastIndex(a, "=")
			if i == -1 {
				r.Attrs[a] = ""
			} else {
				r.Attrs[a[1:i]] = a[i+1:]
			}
		}

		reply = append(reply, r)
		if r.Line == "!done" {
			break
		} else if r.Line == "!trap" {
			result = errors.New(r.Attrs["message"])
		}
	}
	return reply, result
}

func (m router) ListRecords(recordType string) ([]*endpoint.Endpoint, error) {
	filter := ""
	if recordType != "" {
		filter = "?=type=" + recordType
		// A record is the default so 'type' is not defined.
		// catch this case and use a different filter
		if recordType == "A" {
			filter = "?-type"
		}
	}
	filter += "?=comment=" + m.ownerId
	// if filter contains a type filter too, AND them together with the comment
	// which is used as owner id
	if recordType != "" {
		filter += "?#&"
	}
	log.Infof("MikrotikProvider.listRecords: type=%s, filter=%s\n", recordType, filter)

	err := m.send("/ip/dns/static/print", filter)
	if err != nil {
		return nil, err
	}

	reply, err := m.receive()
	if err != nil {
		return nil, err
	}

	if len(reply) > 0 {
		if reply[0].Line == "!trap" {
			return nil, errors.New(reply[0].Attrs["message"])
		}
	}

	out := make([]*endpoint.Endpoint, 0)
	for _, r := range reply {
		if r.Line != "!done" {
			// fmt.Printf("%d: %#v\n", i, r.Attrs)
			name := r.Attrs["name"]
			target := r.Attrs["address"]
			ttl, err := strconv.ParseInt(r.Attrs["ttl"], 10, 64)
			if err != nil {
				ttl = 300
			}

			ttl = int64(math.Max(float64(ttl), float64(m.minimumTTL)))

			rtype, ok := r.Attrs["type"]
			if !ok {
				rtype = "A"
			}
			if rtype == "CNAME" {
				target = r.Attrs["cname"]
			}
			if rtype == "TXT" {
				target = r.Attrs["text"]
			}
			if m.domainFilter.IsConfigured() && !m.domainFilter.Match(name) {
				log.Debugf("Skipping %s that does not match domain filter", name)
				continue
			}
			log.Debugf("Considering %s record for %s: %#v", rtype, name, r.Attrs)
			ep := &endpoint.Endpoint{
				RecordType: rtype,
				DNSName:    name,
				Targets:    []string{target},
				RecordTTL:  endpoint.TTL(ttl),
				Labels:     endpoint.Labels{endpoint.OwnerLabelKey: m.ownerId},
			}
			out = append(out, ep)
		}
	}
	return out, nil
}

func (m router) CreateRecord(ep *endpoint.Endpoint) error {
	if m.dryRun {
		log.Infof("DRY RUN: create %s IN %s -> %s", ep.DNSName, ep.RecordType, ep.Targets[0])
		return nil
	}

	ep.RecordTTL = endpoint.TTL(math.Max(float64(ep.RecordTTL), float64(m.minimumTTL)))

	cmd := make([]string, 0)
	cmd = append(cmd, "/ip/dns/static/add")
	cmd = append(cmd, "=name="+ep.DNSName)
	cmd = append(cmd, "=ttl="+strconv.FormatInt(int64(ep.RecordTTL), 10))
	cmd = append(cmd, "=comment="+m.ownerId)
	switch ep.RecordType {
	case endpoint.RecordTypeA:
		cmd = append(cmd, "=address="+ep.Targets[0])
	case endpoint.RecordTypeAAAA:
		cmd = append(cmd, "=address="+ep.Targets[0])
	case endpoint.RecordTypeCNAME:
		cmd = append(cmd, "=cname="+ep.Targets[0])
	case endpoint.RecordTypeTXT:
		cmd = append(cmd, "=text="+ep.Targets[0])
	default:
		return fmt.Errorf("record type %s not implemented: %#v", ep.RecordType, ep)
	}

	err := m.send(cmd...)
	if err != nil {
		return err
	}

	reply, err := m.receive()
	if err != nil {
		return err
	}

	if len(reply) > 0 {
		if reply[0].Line == "!trap" {
			return errors.New(reply[0].Attrs["message"])
		}
	}
	return nil
}

func (m router) DeleteRecord(ep *endpoint.Endpoint) error {
	if m.dryRun {
		log.Infof("DRY RUN: delete %s IN %s -> %s", ep.DNSName, ep.RecordType, ep.Targets[0])
		return nil
	}

	id, err := m.findRecordId(ep)
	if err != nil {
		return err
	}
	log.Debugf("Delete record id %s", id)

	cmd := make([]string, 0)
	cmd = append(cmd, "/ip/dns/static/remove")
	cmd = append(cmd, "=.id="+id)

	err = m.send(cmd...)
	if err != nil {
		return err
	}

	reply, err := m.receive()
	if err != nil {
		return err
	}

	if len(reply) > 0 {
		if reply[0].Line == "!trap" {
			return errors.New(reply[0].Attrs["message"])
		}
	}

	return nil
}

func (m router) findRecordId(ep *endpoint.Endpoint) (string, error) {
	cmd := make([]string, 0)
	cmd = append(cmd, "/ip/dns/static/print")
	// A record is the default so 'type' is not defined.
	// catch this case and use a different filter
	if ep.RecordType == endpoint.RecordTypeA {
		cmd = append(cmd, "?-type")
	} else {
		cmd = append(cmd, "?=type="+ep.RecordType)
	}
	cmd = append(cmd, "?=name="+ep.DNSName)
	cmd = append(cmd, "?=comment="+m.ownerId)

	err := m.send(cmd...)
	if err != nil {
		return "", err
	}

	reply, err := m.receive()
	if err != nil {
		return "", err
	}

	if len(reply) > 0 {
		if reply[0].Line == "!trap" {
			return "", errors.New(reply[0].Attrs["message"])
		}
	}

	return reply[0].Attrs[".id"], nil
}
