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

package provider

import (
	"net"
	"os"
	"fmt"
	"strconv"
	"time"

	"github.com/miekg/dns"
	log "github.com/sirupsen/logrus"

	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"
)

// rfc2136 provider type
type rfc2136Provider struct {
	nameserver  string
	zoneName    string
	tsigKeyName string
	tsigSecret  string
	insecure    bool

	// only consider hosted zones managing domains ending in this suffix
	domainFilter DomainFilter
	dryRun       bool
}

// NewRfc2136Provider is a factory function for OpenStack rfc2136 providers
func NewRfc2136Provider(domainFilter DomainFilter, dryRun bool) (Provider, error) {
	var host, port, keyName, secret, zoneName string
	var insecure bool
	var err error

	if host = os.Getenv("RFC2136_HOST"); len(host) == 0 {
		return nil,fmt.Errorf("RFC2136_HOST is not set")
	}

	if port = os.Getenv("RFC2136_PORT"); len(port) == 0 {
		return nil,fmt.Errorf("RFC2136_PORT is not set")
	}

	if insecureStr := os.Getenv("RFC2136_INSECURE"); len(insecureStr) == 0 {
		insecure = false
	} else {
		if insecure, err = strconv.ParseBool(insecureStr); err != nil {
			return nil,fmt.Errorf("RFC2136_INSECURE must be a boolean value")
		}
	}

	if !insecure {
		if keyName = os.Getenv("RFC2136_TSIG_KEYNAME"); len(keyName) == 0 {
			return nil,fmt.Errorf("RFC2136_TSIG_KEYNAME is not set")
		}

		if secret = os.Getenv("RFC2136_TSIG_SECRET"); len(secret) == 0 {
			return nil,fmt.Errorf("RFC2136_TSIG_SECRET is not set")
		}
	}

	if zoneName = os.Getenv("RFC2136_ZONE"); len(zoneName) == 0 {
		return nil,fmt.Errorf("RFC2136_ZONE is not set")
	}

	r := &rfc2136Provider{
		domainFilter: domainFilter,
		dryRun:       dryRun,
	}

	r.nameserver = net.JoinHostPort(host, port)
	r.zoneName = dns.Fqdn(zoneName)
	if !insecure {
		r.tsigKeyName = dns.Fqdn(keyName)
		r.tsigSecret = secret
	}

	r.insecure = insecure

	log.Infof("Configured RFC2136 with zone '%s' and nameserver '%s'", r.zoneName, r.nameserver)

	return r, nil
}

// Records returns the list of records.
func (r rfc2136Provider) Records() ([]*endpoint.Endpoint, error) {
	log.Debugf("Records")

	rrs, err := r.List()
	if err != nil {
		return nil, err
	}

	var eps []*endpoint.Endpoint

OuterLoop:
	for _, rr := range rrs {
		log.Debugf("Record=%s", rr)

		if rr.Header().Class != dns.ClassINET {
			continue
		}

		rrFqdn := rr.Header().Name
		rrTTL := endpoint.TTL(rr.Header().Ttl)
		var rrType string
		var rrValues []string
		switch rr.Header().Rrtype {
		case dns.TypeCNAME:
			rrValues = []string{rr.(*dns.CNAME).Target}
			rrType = "CNAME"
		case dns.TypeA:
			rrValues = []string{rr.(*dns.A).A.String()}
			rrType = "A"
		case dns.TypeAAAA:
			rrValues = []string{rr.(*dns.AAAA).AAAA.String()}
			rrType = "AAAA"
		case dns.TypeTXT:
			rrValues = (rr.(*dns.TXT).Txt)
			rrType = "TXT"
		default:
			continue // Unhandled record type
		}

		for idx, existingEndpoint := range eps {
			if existingEndpoint.DNSName == rrFqdn && existingEndpoint.RecordType == rrType {
				eps[idx].Targets = append(eps[idx].Targets, rrValues...)
				continue OuterLoop
			}
		}

		ep := endpoint.NewEndpointWithTTL(
			rrFqdn,
			rrType,
			rrTTL,
			rrValues...,
		)
		ep.Labels["originalText"] = "originalText"
		ep.Labels["prefix"] = "prefix"
		eps = append(eps, ep)
	}

	return eps, nil
}

func (r rfc2136Provider) List() ([]dns.RR, error) {
	log.Debugf("Fetching records for '%s'", r.zoneName)

	t := new(dns.Transfer)
	if !r.insecure {
		t.TsigSecret = map[string]string{r.tsigKeyName: r.tsigSecret}
	}

	m := new(dns.Msg)
	m.SetAxfr(r.zoneName)
	if !r.insecure {
		m.SetTsig(r.tsigKeyName, dns.HmacMD5, 300, time.Now().Unix())
	}

	env, err := t.In(m, r.nameserver)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch records via AXFR: %v", err)
	}

	records := make([]dns.RR, 0)
	for e := range env {
		if e.Error != nil {
			if e.Error == dns.ErrSoa {
				log.Error("AXFR error: unexpected response received from the server")
			} else {
				log.Errorf("AXFR error: %v", e.Error)
			}
			continue
		}
		records = append(records, e.RR...)
	}

	return records, nil
}

// ApplyChanges applies a given set of changes in a given zone.
func (r rfc2136Provider) ApplyChanges(changes *plan.Changes) error {
	log.Debugf("ApplyChanges")

	for _, ep := range changes.Create {

		if !r.domainFilter.Match(ep.DNSName) {
			log.Debugf("Skipping record %s because it was filtered out by the specified --domain-filter", ep.DNSName)
			continue
		}

		r.AddRecord(ep)
	}
	for _, ep := range changes.UpdateNew {

		if !r.domainFilter.Match(ep.DNSName) {
			log.Debugf("Skipping record %s because it was filtered out by the specified --domain-filter", ep.DNSName)
			continue
		}

		r.UpdateRecord(ep)
	}
	for _, ep := range changes.Delete {

		if !r.domainFilter.Match(ep.DNSName) {
			log.Debugf("Skipping record %s because it was filtered out by the specified --domain-filter", ep.DNSName)
			continue
		}

		r.RemoveRecord(ep)
	}
	
	return nil
}

func (r rfc2136Provider) UpdateRecord(ep *endpoint.Endpoint) error {
	err := r.RemoveRecord(ep)
	if err != nil {
		return err
	}

	return r.AddRecord(ep)
}

func (r rfc2136Provider) AddRecord(ep *endpoint.Endpoint) error {
	log.Debugf("AddRecord.ep=%s", ep)

	newRR := fmt.Sprintf("%s %d %s %s", ep.DNSName, ep.RecordTTL, ep.RecordType, ep.Targets)
	log.Debugf("Adding RR: %s", newRR)

	rr, err := dns.NewRR(newRR)
	if err != nil {
		return fmt.Errorf("Failed to build RR: %v", err)
	}

	rrs := make([]dns.RR, 1)
	rrs[0] = rr

	m := new(dns.Msg)
	m.SetUpdate(r.zoneName)
	m.Insert(rrs)

	err = r.SendMessage(m)
	if err != nil {
		return fmt.Errorf("RFC2136 query failed: %v", err)
	}

	return nil
}

func (r rfc2136Provider) RemoveRecord(ep *endpoint.Endpoint) error {
	log.Debugf("RemoveRecord.ep=%s", ep)

	newRR := fmt.Sprintf("%s 0 %s 0.0.0.0", ep.DNSName, ep.RecordType)
	log.Debugf("Adding RR: %s", newRR)

	rr, err := dns.NewRR(newRR)
	if err != nil {
		return fmt.Errorf("Failed to build RR: %v", err)
	}

	rrs := make([]dns.RR, 1)
	rrs[0] = rr

	m := new(dns.Msg)
	m.SetUpdate(r.zoneName)
	m.RemoveRRset(rrs)

	err = r.SendMessage(m)
	if err != nil {
		return fmt.Errorf("RFC2136 query failed: %v", err)
	}

	return nil
}

func (r rfc2136Provider) SendMessage(msg *dns.Msg) error {
	if !r.dryRun {
		log.Debugf("SendMessage")
	} else {
		log.Debugf("SendMessage.skipped")
		return nil
	}

	c := new(dns.Client)
	c.SingleInflight = true

	if !r.insecure {
		c.TsigSecret = map[string]string{r.tsigKeyName: r.tsigSecret}
		msg.SetTsig(r.tsigKeyName, dns.HmacMD5, 300, time.Now().Unix())
	}

	resp, _, err := c.Exchange(msg, r.nameserver)
	if err != nil {
		log.Infof("error in dns.Client.Exchange: %s", err)
		return err
	}
	if resp != nil && resp.Rcode != dns.RcodeSuccess {
		log.Infof("Bad dns.Client.Exchange response: %s", resp)
		return fmt.Errorf("Bad return code: %s", dns.RcodeToString[resp.Rcode])
	}

	log.Debugf("SendMessage.success")
	return nil
}
