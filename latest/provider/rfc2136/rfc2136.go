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

package rfc2136

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/bodgit/tsig"
	"github.com/bodgit/tsig/gss"
	"github.com/miekg/dns"

	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/tlsutils"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

const (
	// maximum time DNS client can be off from server for an update to succeed
	clockSkew = 300
)

// rfc2136 provider type
type rfc2136Provider struct {
	provider.BaseProvider
	nameservers     []string
	zoneNames       []string
	tsigKeyName     string
	tsigSecret      string
	tsigSecretAlg   string
	insecure        bool
	axfr            bool
	minTTL          time.Duration
	batchChangeSize int
	tlsConfig       TLSConfig
	createPTR       bool

	// options specific to rfc3645 gss-tsig support
	gssTsig      bool
	krb5Username string
	krb5Password string
	krb5Realm    string

	// only consider hosted zones managing domains ending in this suffix
	domainFilter *endpoint.DomainFilter
	dryRun       bool
	actions      rfc2136Actions

	// Counter for load balancing, and error handling
	counter int
	mu      sync.Mutex // Mutex for thread-safe counter

	// Load balancing strategy "round-robin", "random", or "disabled"
	loadBalancingStrategy string

	// Random number generator for random load balancing
	randGen *rand.Rand

	// Last error encountered
	lastErr error
}

// TLSConfig is comprised of the TLS-related fields necessary if we are using DNS over TLS
type TLSConfig struct {
	UseTLS                bool
	SkipTLSVerify         bool
	CAFilePath            string
	ClientCertFilePath    string
	ClientCertKeyFilePath string
}

// Map of supported TSIG algorithms
var tsigAlgs = map[string]string{
	"hmac-sha1":   dns.HmacSHA1,
	"hmac-sha224": dns.HmacSHA224,
	"hmac-sha256": dns.HmacSHA256,
	"hmac-sha384": dns.HmacSHA384,
	"hmac-sha512": dns.HmacSHA512,
}

type rfc2136Actions interface {
	SendMessage(msg *dns.Msg) error
	IncomeTransfer(m *dns.Msg, nameserver string) (env chan *dns.Envelope, err error)
}

// NewRfc2136Provider is a factory function for OpenStack rfc2136 providers
func NewRfc2136Provider(hosts []string, port int, zoneNames []string, insecure bool, keyName string, secret string, secretAlg string, axfr bool, domainFilter *endpoint.DomainFilter, dryRun bool, minTTL time.Duration, createPTR bool, gssTsig bool, krb5Username string, krb5Password string, krb5Realm string, batchChangeSize int, tlsConfig TLSConfig, loadBalancingStrategy string, actions rfc2136Actions) (provider.Provider, error) {
	secretAlgChecked, ok := tsigAlgs[secretAlg]
	if !ok && !insecure && !gssTsig {
		return nil, fmt.Errorf("%s is not supported TSIG algorithm", secretAlg)
	}

	// Set zone to root if no set
	if len(zoneNames) == 0 {
		zoneNames = append(zoneNames, ".")
	}

	// Sort zones
	sort.Slice(zoneNames, func(i, j int) bool {
		return len(strings.Split(zoneNames[i], ".")) > len(strings.Split(zoneNames[j], "."))
	})

	var nameservers []string
	for _, host := range hosts {
		host = net.JoinHostPort(host, strconv.Itoa(port))
		nameservers = append(nameservers, host)
	}

	r := &rfc2136Provider{
		nameservers:           nameservers,
		zoneNames:             zoneNames,
		insecure:              insecure,
		gssTsig:               gssTsig,
		createPTR:             createPTR,
		krb5Username:          krb5Username,
		krb5Password:          krb5Password,
		krb5Realm:             strings.ToUpper(krb5Realm),
		domainFilter:          domainFilter,
		dryRun:                dryRun,
		axfr:                  axfr,
		minTTL:                minTTL,
		batchChangeSize:       batchChangeSize,
		tlsConfig:             tlsConfig,
		loadBalancingStrategy: loadBalancingStrategy,
		randGen:               rand.New(rand.NewSource(time.Now().UnixNano())),
		counter:               0,
		lastErr:               nil,
	}
	if actions != nil {
		r.actions = actions
	} else {
		r.actions = r
	}

	if !insecure {
		r.tsigKeyName = dns.Fqdn(keyName)
		r.tsigSecret = secret
		r.tsigSecretAlg = secretAlgChecked
	}

	log.Infof("Configured RFC2136 with zones '%v' and nameservers '%v'", r.zoneNames, hosts)
	return r, nil
}

// KeyData will return TKEY name and TSIG handle to use for followon actions with a secure connection
func (r *rfc2136Provider) KeyData(nameserver string) (keyName string, handle *gss.Client, err error) {
	handle, err = gss.NewClient(new(dns.Client))
	if err != nil {
		return keyName, handle, err
	}

	keyName, _, err = handle.NegotiateContextWithCredentials(nameserver, r.krb5Realm, r.krb5Username, r.krb5Password)
	if err != nil {
		return keyName, handle, err
	}

	return keyName, handle, nil
}

// Records returns the list of records.
func (r *rfc2136Provider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
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
		case dns.TypeNS:
			rrValues = []string{rr.(*dns.NS).Ns}
			rrType = "NS"
		case dns.TypePTR:
			rrValues = []string{rr.(*dns.PTR).Ptr}
			rrType = "PTR"
		default:
			continue // Unhandled record type
		}

		for idx, existingEndpoint := range eps {
			if existingEndpoint.DNSName == strings.TrimSuffix(rrFqdn, ".") && existingEndpoint.RecordType == rrType {
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

		eps = append(eps, ep)
	}

	return eps, nil
}

func (r *rfc2136Provider) IncomeTransfer(m *dns.Msg, nameserver string) (env chan *dns.Envelope, err error) {
	t := new(dns.Transfer)
	if !r.insecure && !r.gssTsig {
		t.TsigSecret = map[string]string{r.tsigKeyName: r.tsigSecret}
	}

	c, err := makeClient(r, nameserver)
	if err != nil {
		return nil, fmt.Errorf("error setting up TLS: %w", err)
	}
	conn, err := c.Dial(nameserver)
	if err != nil {
		return nil, fmt.Errorf("failed to connect for transfer: %w", err)
	}
	t.Conn = conn
	return t.In(m, nameserver)
}

func (r *rfc2136Provider) List() ([]dns.RR, error) {
	if !r.axfr {
		log.Debug("axfr is disabled")
		return make([]dns.RR, 0), nil
	}

	records := make([]dns.RR, 0)
	for _, zone := range r.zoneNames {
		log.Debugf("Fetching records for '%q'", zone)

		m := new(dns.Msg)
		m.SetAxfr(dns.Fqdn(zone))
		if !r.insecure && !r.gssTsig {
			m.SetTsig(r.tsigKeyName, r.tsigSecretAlg, clockSkew, time.Now().Unix())
		}

		var lastErr error
		for i := 0; i < len(r.nameservers); i++ {
			nameserver := r.getNextNameserver()
			log.Debugf("Fetching records from nameserver: %s", nameserver)

			env, err := r.actions.IncomeTransfer(m, nameserver)
			if err != nil {
				lastErr = fmt.Errorf("failed to fetch records via AXFR: %w", err)
				r.lastErr = lastErr
				continue
			}

			for e := range env {
				if e.Error != nil {
					if errors.Is(e.Error, dns.ErrSoa) {
						log.Error("AXFR error: unexpected response received from the server")
					} else {
						log.Errorf("AXFR error: %v", e.Error)
					}
					continue
				}
				records = append(records, e.RR...)
			}
			// If records were fetched successfully, break out of the loop
			if len(records) > 0 {
				break
			}
		}

		if lastErr != nil {
			r.lastErr = lastErr
			return nil, lastErr
		}
	}

	return records, nil
}

func (r *rfc2136Provider) AddReverseRecord(ip string, hostname string) error {
	changes := r.GenerateReverseRecord(ip, hostname)
	return r.ApplyChanges(context.Background(), &plan.Changes{Create: changes})
}

func (r *rfc2136Provider) RemoveReverseRecord(ip string, hostname string) error {
	changes := r.GenerateReverseRecord(ip, hostname)
	return r.ApplyChanges(context.Background(), &plan.Changes{Delete: changes})
}

func (r *rfc2136Provider) GenerateReverseRecord(ip string, hostname string) []*endpoint.Endpoint {
	// Generate PTR notation record starting from the IP address
	var records []*endpoint.Endpoint

	log.Debugf("Reverse zone is: %s %s", ip, dns.Fqdn(ip))
	reverseAddress, _ := dns.ReverseAddr(ip)

	// PTR
	records = append(records, &endpoint.Endpoint{
		DNSName:    reverseAddress[:len(reverseAddress)-1],
		RecordType: "PTR",
		Targets:    endpoint.Targets{hostname},
	})

	return records
}

// ApplyChanges applies a given set of changes in a given zone.
func (r *rfc2136Provider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	log.Debugf("ApplyChanges (Create: %d, UpdateOld: %d, UpdateNew: %d, Delete: %d)", len(changes.Create), len(changes.UpdateOld), len(changes.UpdateNew), len(changes.Delete))

	var errs []error

	for c, chunk := range chunkBy(changes.Create, r.batchChangeSize) {
		log.Debugf("Processing batch %d of create changes", c)

		m := make(map[string]*dns.Msg)
		m["."] = new(dns.Msg) // Add the root zone
		for _, z := range r.zoneNames {
			z = dns.Fqdn(z)
			m[z] = new(dns.Msg)
		}
		for _, ep := range chunk {
			if !r.domainFilter.Match(ep.DNSName) {
				log.Debugf("Skipping record %s because it was filtered out by the specified --domain-filter", ep.DNSName)
				continue
			}

			zone := findMsgZone(ep, r.zoneNames)
			m[zone].SetUpdate(zone)

			r.AddRecord(m[zone], ep)

			if r.createPTR && (ep.RecordType == "A" || ep.RecordType == "AAAA") {
				r.AddReverseRecord(ep.Targets[0], ep.DNSName)
			}
		}

		// only send if there are records available
		for _, z := range m {
			if len(z.Ns) > 0 {
				if err := r.actions.SendMessage(z); err != nil {
					log.Errorf("RFC2136 create record failed: %v", err)
					errs = append(errs, err)
					continue
				}
			}
		}
	}

	for c, chunk := range chunkBy(changes.UpdateNew, r.batchChangeSize) {
		log.Debugf("Processing batch %d of update changes", c)

		m := make(map[string]*dns.Msg)
		m["."] = new(dns.Msg) // Add the root zone
		for _, z := range r.zoneNames {
			z = dns.Fqdn(z)
			m[z] = new(dns.Msg)
		}

		for i, ep := range chunk {
			if !r.domainFilter.Match(ep.DNSName) {
				log.Debugf("Skipping record %s because it was filtered out by the specified --domain-filter", ep.DNSName)
				continue
			}

			zone := findMsgZone(ep, r.zoneNames)
			m[zone].SetUpdate(zone)

			r.UpdateRecord(m[zone], changes.UpdateOld[i], ep)
			if r.createPTR && (ep.RecordType == "A" || ep.RecordType == "AAAA") {
				r.RemoveReverseRecord(changes.UpdateOld[i].Targets[0], ep.DNSName)
				r.AddReverseRecord(ep.Targets[0], ep.DNSName)
			}
		}

		// only send if there are records available
		for _, z := range m {
			if len(z.Ns) > 0 {
				if err := r.actions.SendMessage(z); err != nil {
					log.Errorf("RFC2136 update record failed: %v", err)
					errs = append(errs, err)
					continue
				}
			}
		}
	}

	for c, chunk := range chunkBy(changes.Delete, r.batchChangeSize) {
		log.Debugf("Processing batch %d of delete changes", c)

		m := make(map[string]*dns.Msg)
		m["."] = new(dns.Msg) // Add the root zone
		for _, z := range r.zoneNames {
			z = dns.Fqdn(z)
			m[z] = new(dns.Msg)
		}
		for _, ep := range chunk {
			if !r.domainFilter.Match(ep.DNSName) {
				log.Debugf("Skipping record %s because it was filtered out by the specified --domain-filter", ep.DNSName)
				continue
			}

			zone := findMsgZone(ep, r.zoneNames)
			m[zone].SetUpdate(zone)

			r.RemoveRecord(m[zone], ep)
			if r.createPTR && (ep.RecordType == "A" || ep.RecordType == "AAAA") {
				r.RemoveReverseRecord(ep.Targets[0], ep.DNSName)
			}
		}

		// only send if there are records available
		for _, z := range m {
			if len(z.Ns) > 0 {
				if err := r.actions.SendMessage(z); err != nil {
					log.Errorf("RFC2136 delete record failed: %v", err)
					errs = append(errs, err)
					continue
				}
			}
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("RFC2136 had errors in one or more of its batches: %v", errs)
	}

	return nil
}

func (r *rfc2136Provider) UpdateRecord(m *dns.Msg, oldEp *endpoint.Endpoint, newEp *endpoint.Endpoint) error {
	err := r.RemoveRecord(m, oldEp)
	if err != nil {
		return err
	}

	return r.AddRecord(m, newEp)
}

func (r *rfc2136Provider) AddRecord(m *dns.Msg, ep *endpoint.Endpoint) error {
	log.Debugf("AddRecord.ep=%s", ep)

	ttl := int64(r.minTTL.Seconds())
	if ep.RecordTTL.IsConfigured() && int64(ep.RecordTTL) > ttl {
		ttl = int64(ep.RecordTTL)
	}

	for _, target := range ep.Targets {
		newRR := fmt.Sprintf("%s %d %s %s", ep.DNSName, ttl, ep.RecordType, target)
		log.Infof("Adding RR: %s", newRR)

		rr, err := dns.NewRR(newRR)
		if err != nil {
			return fmt.Errorf("failed to build RR: %w", err)
		}

		m.Insert([]dns.RR{rr})
	}

	return nil
}

func (r *rfc2136Provider) RemoveRecord(m *dns.Msg, ep *endpoint.Endpoint) error {
	log.Debugf("RemoveRecord.ep=%s", ep)
	for _, target := range ep.Targets {
		newRR := fmt.Sprintf("%s %d %s %s", ep.DNSName, ep.RecordTTL, ep.RecordType, target)
		log.Infof("Removing RR: %s", newRR)

		rr, err := dns.NewRR(newRR)
		if err != nil {
			return fmt.Errorf("failed to build RR: %w", err)
		}

		m.Remove([]dns.RR{rr})
	}

	return nil
}

func (r *rfc2136Provider) getNextNameserver() string {
	if len(r.nameservers) == 1 {
		return r.nameservers[0]
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if r.lastErr != nil {
		log.Warnf("Last operation failed for nameserver %s", r.nameservers[r.counter])
		log.Warnf("Last operation error message: %v", r.lastErr)
	}

	var nameserver string
	switch r.loadBalancingStrategy {
	case "random":
		for {
			nameserver = r.nameservers[r.randGen.Intn(len(r.nameservers))]
			// Ensure that we don't get the same nameserver as the last one
			if nameserver != r.nameservers[r.counter] {
				break
			}
		}
	case "round-robin":
		nameserver = r.nameservers[r.counter]
		r.counter = (r.counter + 1) % len(r.nameservers)
	default:
		if r.lastErr != nil {
			r.counter = (r.counter + 1) % len(r.nameservers)
			nameserver = r.nameservers[r.counter]
		} else {
			nameserver = r.nameservers[r.counter]
		}
	}

	// Last error has been logged, reset it for the next operation
	r.lastErr = nil
	return nameserver
}

func (r *rfc2136Provider) SendMessage(msg *dns.Msg) error {
	if r.dryRun {
		log.Debugf("SendMessage.skipped")
		return nil
	}
	log.Debugf("SendMessage")

	var lastErr error
	for i := 0; i < len(r.nameservers); i++ {
		nameserver := r.getNextNameserver()
		log.Debugf("Sending message to nameserver: %s", nameserver)

		c, err := makeClient(r, nameserver)
		if err != nil {
			lastErr = fmt.Errorf("error setting up TLS: %w", err)
			r.lastErr = lastErr
			continue
		}

		if !r.insecure {
			if r.gssTsig {
				keyName, handle, err := r.KeyData(nameserver)
				if err != nil {
					lastErr = err
					r.lastErr = lastErr
					continue
				}
				defer handle.Close()
				defer handle.DeleteContext(keyName)

				c.TsigProvider = handle

				msg.SetTsig(keyName, tsig.GSS, clockSkew, time.Now().Unix())
			} else {
				c.TsigProvider = tsig.HMAC{r.tsigKeyName: r.tsigSecret}
				msg.SetTsig(r.tsigKeyName, r.tsigSecretAlg, clockSkew, time.Now().Unix())
			}
		}

		resp, _, err := c.Exchange(msg, nameserver)
		if err != nil {
			if resp != nil && resp.Rcode != dns.RcodeSuccess {
				log.Infof("error in dns.Client.Exchange: %s", err)
				lastErr = err
				r.lastErr = lastErr
				continue
			}
			log.Warnf("warn in dns.Client.Exchange: %s", err)
			lastErr = err
			r.lastErr = lastErr
			continue
		}
		if resp != nil && resp.Rcode != dns.RcodeSuccess {
			log.Infof("Bad dns.Client.Exchange response: %s", resp)
			lastErr = fmt.Errorf("bad return code: %s", dns.RcodeToString[resp.Rcode])
			r.lastErr = lastErr
			continue
		}

		log.Debugf("SendMessage.success")
		return nil
	}

	r.lastErr = lastErr
	return lastErr
}

func chunkBy(slice []*endpoint.Endpoint, chunkSize int) [][]*endpoint.Endpoint {
	var chunks [][]*endpoint.Endpoint

	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}

	return chunks
}

func findMsgZone(ep *endpoint.Endpoint, zoneNames []string) string {
	for _, zone := range zoneNames {
		if strings.HasSuffix(ep.DNSName, zone) {
			return dns.Fqdn(zone)
		}
	}

	log.Warnf("No available zone found for %s, set it to 'root'", ep.DNSName)
	return dns.Fqdn(".")
}

func makeClient(r *rfc2136Provider, nameserver string) (*dns.Client, error) {
	c := new(dns.Client)

	// Remove port from nameserver
	nameserver = strings.Split(nameserver, ":")[0]

	if r.tlsConfig.UseTLS {
		log.Debug("RFC2136 Connecting via TLS")
		c.Net = "tcp-tls"
		tlsConfig, err := tlsutils.NewTLSConfig(
			r.tlsConfig.ClientCertFilePath,
			r.tlsConfig.ClientCertKeyFilePath,
			r.tlsConfig.CAFilePath,
			nameserver, // Use the current nameserver
			r.tlsConfig.SkipTLSVerify,
			// Per RFC9103
			tls.VersionTLS13,
		)
		if err != nil {
			return nil, err
		}
		if tlsConfig.NextProtos == nil {
			// Per RFC9103
			tlsConfig.NextProtos = []string{"dot"}
		}
		c.TLSConfig = tlsConfig
	} else {
		c.Net = "tcp"
	}

	return c, nil
}
