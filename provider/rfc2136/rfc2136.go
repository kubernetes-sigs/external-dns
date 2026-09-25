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
	"unicode/utf8"

	"github.com/bodgit/tsig"
	"github.com/bodgit/tsig/gss"
	"github.com/miekg/dns"

	"sigs.k8s.io/external-dns/pkg/apis/externaldns"

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
	axfrInsecure    bool
	axfr            bool
	minTTL          time.Duration
	batchChangeSize int
	tlsConfig       TLSConfig

	// options specific to rfc3645 gss-tsig support
	gssTsig      bool
	krb5Username string
	krb5Password string
	krb5Realm    string

	// only consider hosted zones managing domains ending in this suffix
	domainFilter *endpoint.DomainFilter
	dryRun       bool
	actions      rfc2136Actions

	// Counters and last-seen errors for load balancing. List (AXFR) and
	// Send (update) paths rotate and fail over independently, so each
	// tracks its own counter and last error instead of sharing state that
	// would let one operation skew or mask failover for the other.
	listCounter int
	sendCounter int
	listLastErr error
	sendLastErr error
	mu          sync.Mutex // Mutex for thread-safe counters and last errors

	// Load balancing strategy "round-robin", "random", or "disabled"
	loadBalancingStrategy string

	// Random number generator for random load balancing
	randGen *rand.Rand

	// txtWire holds the character-strings of each TXT RR seen by the last
	// Records(), keyed by owner name and canonical value. RFC 2136 deletes
	// match rdata byte for byte, so a record chunked differently from
	// chunkTXT (older external-dns, other tools) can only be removed with
	// the boundaries it has in the zone.
	txtWire   map[txtWireKey][]string
	txtWireMu sync.Mutex
}

type txtWireKey struct {
	fqdn  string
	value string
}

func newTXTWireKey(name, value string) txtWireKey {
	return txtWireKey{fqdn: strings.ToLower(dns.Fqdn(name)), value: value}
}

// nameserverOp identifies which rotation - list (AXFR) or send (update) -
// is requesting the next nameserver, so getNextNameserverFor can resolve
// the right counter and last-error state internally instead of the caller
// having to know which field to pass in.
type nameserverOp int

const (
	nameserverOpList nameserverOp = iota
	nameserverOpSend
)

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

// New creates an RFC2136 provider from the given configuration.
func New(_ context.Context, cfg *externaldns.Config, domainFilter *endpoint.DomainFilter) (provider.Provider, error) {
	tlsConfig := TLSConfig{
		UseTLS:                cfg.RFC2136UseTLS,
		SkipTLSVerify:         cfg.RFC2136SkipTLSVerify,
		CAFilePath:            cfg.TLSCA,
		ClientCertFilePath:    cfg.TLSClientCert,
		ClientCertKeyFilePath: cfg.TLSClientCertKey,
	}

	// Without AXFR records cannot be listed, so the plan never updates or deletes.
	if !cfg.RFC2136AXFR && cfg.Policy != "create-only" {
		log.Warnf("--rfc2136-axfr is not set: ExternalDNS cannot list existing records, so --policy=%s will never update or delete them", cfg.Policy)
	}

	return newProvider(cfg.RFC2136Host, cfg.RFC2136Port, cfg.RFC2136Zone, cfg.RFC2136Insecure, cfg.RFC2136AXFRInsecure, cfg.RFC2136TSIGKeyName, cfg.RFC2136TSIGSecret, cfg.RFC2136TSIGSecretAlg, cfg.RFC2136AXFR, domainFilter, cfg.DryRun, cfg.RFC2136MinTTL, cfg.RFC2136GSSTSIG, cfg.RFC2136KerberosUsername, cfg.RFC2136KerberosPassword, cfg.RFC2136KerberosRealm, cfg.RFC2136BatchChangeSize, tlsConfig, cfg.RFC2136LoadBalancingStrategy, nil)
}

// newProvider is a factory function for OpenStack rfc2136 providers
func newProvider(hosts []string, port int, zoneNames []string, insecure bool, axfrInsecure bool, keyName string, secret string, secretAlg string, axfr bool, domainFilter *endpoint.DomainFilter, dryRun bool, minTTL time.Duration, gssTsig bool, krb5Username string, krb5Password string, krb5Realm string, batchChangeSize int, tlsConfig TLSConfig, loadBalancingStrategy string, actions rfc2136Actions) (provider.Provider, error) {
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
		axfrInsecure:          axfrInsecure,
		gssTsig:               gssTsig,
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
		listCounter:           0,
		sendCounter:           0,
		listLastErr:           nil,
		sendLastErr:           nil,
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

	if axfrInsecure && axfr {
		log.Warn("--rfc2136-axfr-insecure is set: zone transfers are unauthenticated")
	}

	log.Infof("Configured RFC2136 with zones '%v' and nameservers '%v'", r.zoneNames, hosts)
	return r, nil
}

// KeyData will return TKEY name and TSIG handle to use for followon actions with a secure connection
func (r *rfc2136Provider) KeyData(nameserver string) (string, *gss.Client, error) {
	handle, err := gss.NewClient(new(dns.Client))
	if err != nil {
		return "", handle, err
	}

	keyName, _, err := handle.NegotiateContextWithCredentials(nameserver, r.krb5Realm, r.krb5Username, r.krb5Password)
	if err != nil {
		return keyName, handle, err
	}

	return keyName, handle, nil
}

// Records returns the list of records.
func (r *rfc2136Provider) Records(_ context.Context) ([]*endpoint.Endpoint, error) {
	rrs, err := r.List()
	if err != nil {
		return nil, err
	}

	var eps []*endpoint.Endpoint
	txtWire := map[txtWireKey][]string{}

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
		case dns.TypeDNAME:
			rrValues = []string{rr.(*dns.DNAME).Target}
			rrType = "DNAME"
		case dns.TypeA:
			rrValues = []string{rr.(*dns.A).A.String()}
			rrType = "A"
		case dns.TypeAAAA:
			rrValues = []string{rr.(*dns.AAAA).AAAA.String()}
			rrType = "AAAA"
		case dns.TypeTXT:
			// Read as one value, the concatenation of the character-strings
			// (the SPF and DKIM convention, RFC 7208 §3.3 and RFC 6376
			// §3.6.2.2). Reading it exploded or presentation-escaped never
			// matches the canonical target produced by AdjustEndpoints, so
			// every reconcile loops on no-op updates (#1596).
			txt := rr.(*dns.TXT)
			value := txtValue(txt)
			txtWire[newTXTWireKey(rrFqdn, value)] = txt.Txt
			rrValues = []string{value}
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

	r.txtWireMu.Lock()
	r.txtWire = txtWire
	r.txtWireMu.Unlock()

	return eps, nil
}

// AdjustEndpoints canonicalizes the targets of TXT endpoints into the form
// Records() returns: the concatenation of the TXT character-strings (#1596).
// Without it, quoted or chunked targets (e.g. DKIM keys) never compare equal
// to the read-back records and every reconcile loops on no-op updates.
// Character-string boundaries given by the user are not preserved: the value
// is re-chunked at 255 bytes on write.
func (r *rfc2136Provider) AdjustEndpoints(endpoints []*endpoint.Endpoint) ([]*endpoint.Endpoint, error) {
	for _, ep := range endpoints {
		if ep.RecordType != endpoint.RecordTypeTXT {
			continue
		}
		for i, target := range ep.Targets {
			canonical, err := canonicalTXTTarget(target)
			if err != nil {
				// Any error here stops RunOnce before planning, freezing the
				// whole zone for one bad object; dropping the endpoint would
				// make policy=sync delete its record. Keep the target as is:
				// AddRecord fails to build it and only this record is skipped.
				log.Warnf("Keeping TXT target of %s as is: %v", ep.DNSName, err)
				continue
			}
			ep.Targets[i] = canonical
		}
	}
	return endpoints, nil
}

// isZoneQuoted reports whether target is in zone-file quoted form, as written
// by users for chunked values and by the TXT registry for heritage labels
// (Labels.Serialize(true)). Anything else is a canonical value, inner quotes
// included; a canonical value that both starts and ends with '"' is the one
// ambiguous case.
func isZoneQuoted(target string) bool {
	return len(target) >= 2 && target[0] == '"' && target[len(target)-1] == '"'
}

// canonicalTXTTarget converts a TXT target to its canonical form: the
// concatenation of its raw character-string bytes. A target without
// zone-file quoting is already canonical and returned unchanged; a quoted
// target is parsed as it would be on the wire, surfacing the same syntax
// errors the write path used to. The parser yields the presentation form
// (miekg/dns keeps \DDD and \X escapes in dns.TXT.Txt), so they are undone.
func canonicalTXTTarget(target string) (string, error) {
	if !isZoneQuoted(target) {
		return target, nil
	}
	rr, err := dns.NewRR(fmt.Sprintf("external-dns.invalid. 60 IN TXT %s", target))
	if err != nil {
		return "", fmt.Errorf("failed to parse TXT target %q: %w", target, err)
	}
	return unescapeTXT(strings.Join(rr.(*dns.TXT).Txt, "")), nil
}

// chunkTXT splits a canonical TXT target into character-strings of at most
// 255 bytes, the DNS wire limit, never breaking a multi-byte UTF-8 rune.
func chunkTXT(s string) []string {
	chunks := []string{}
	for len(s) > 255 {
		end := 255
		for end > 0 && !utf8.RuneStart(s[end]) {
			end--
		}
		if end == 0 {
			// No rune boundary in range: invalid UTF-8, split at the byte limit.
			end = 255
		}
		chunks = append(chunks, s[:end])
		s = s[end:]
	}
	return append(chunks, s)
}

// dns.TXT.Txt holds the zone-presentation form: packTxtString decodes \DDD
// and \X escapes when packing to the wire, and unpackString encodes ", \ and
// bytes outside 0x20-0x7E when unpacking. escapeTXT and unescapeTXT convert
// between that form and the canonical raw bytes the provider compares.

// escapeTXT renders raw canonical bytes in the presentation form, so that
// packing puts the value on the wire verbatim.
func escapeTXT(s string) string {
	needsEscape := false
	for i := 0; i < len(s); i++ {
		if c := s[i]; c == '"' || c == '\\' || c < ' ' || c > '~' {
			needsEscape = true
			break
		}
	}
	if !needsEscape {
		return s
	}

	var b strings.Builder
	b.Grow(len(s))
	for i := 0; i < len(s); i++ {
		switch c := s[i]; {
		case c == '"' || c == '\\':
			b.WriteByte('\\')
			b.WriteByte(c)
		case c < ' ' || c > '~':
			b.WriteByte('\\')
			b.WriteByte('0' + c/100)
			b.WriteByte('0' + (c/10)%10)
			b.WriteByte('0' + c%10)
		default:
			b.WriteByte(c)
		}
	}
	return b.String()
}

// unescapeTXT reverses the unpack escaping: \" and \\ collapse to the byte
// they quote, \DDD decodes a byte by value. Input is unpackString's output,
// which never emits other escapes; a malformed sequence is kept verbatim.
func unescapeTXT(s string) string {
	if !strings.Contains(s, `\`) {
		return s
	}

	var b strings.Builder
	b.Grow(len(s))
	for i := 0; i < len(s); i++ {
		if s[i] != '\\' || i+1 >= len(s) {
			b.WriteByte(s[i])
			continue
		}
		if i+3 < len(s) && isDigitByte(s[i+1]) && isDigitByte(s[i+2]) && isDigitByte(s[i+3]) {
			v := int(s[i+1]-'0')*100 + int(s[i+2]-'0')*10 + int(s[i+3]-'0')
			if v <= 255 {
				b.WriteByte(byte(v))
				i += 3
				continue
			}
			b.WriteByte(s[i]) // out of range: not a valid escape, keep verbatim
			continue
		}
		b.WriteByte(s[i+1])
		i++
	}
	return b.String()
}

func isDigitByte(b byte) bool {
	return b >= '0' && b <= '9'
}

// txtValue returns the canonical target form of a received TXT RR: the
// concatenation of its character-strings, with the presentation escaping
// undone.
func txtValue(rr *dns.TXT) string {
	return unescapeTXT(strings.Join(rr.Txt, ""))
}

// shouldSignAXFR reports whether TSIG should be attached to zone transfers.
func (r *rfc2136Provider) shouldSignAXFR() bool {
	return !r.insecure && !r.gssTsig && !r.axfrInsecure
}

func (r *rfc2136Provider) IncomeTransfer(m *dns.Msg, nameserver string) (chan *dns.Envelope, error) {
	t := new(dns.Transfer)

	if r.shouldSignAXFR() {
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

		var lastErr error
		for i := 0; i < len(r.nameservers); i++ {
			nameserver := r.getNextNameserverFor(nameserverOpList)
			log.Debugf("Fetching records from nameserver: %s", nameserver)

			// Signing strips the TSIG RR, so a reused message goes out unsigned.
			m := new(dns.Msg)
			m.SetAxfr(dns.Fqdn(zone))
			if r.shouldSignAXFR() {
				m.SetTsig(r.tsigKeyName, r.tsigSecretAlg, clockSkew, time.Now().Unix())
			}

			env, err := r.actions.IncomeTransfer(m, nameserver)
			if err != nil {
				lastErr = fmt.Errorf("failed to fetch records via AXFR for zone %q from %s: %w", zone, nameserver, err)
				r.listLastErr = lastErr
				continue
			}

			var attempt []dns.RR
			var attemptErr error
			for e := range env {
				if e.Error != nil {
					if errors.Is(e.Error, dns.ErrSoa) {
						log.Error("AXFR error: unexpected response received from the server")
					} else {
						log.Errorf("AXFR error: %v", e.Error)
					}
					attemptErr = e.Error
					// Producer closes the channel after a single error envelope.
					break
				}
				// Error envelopes can carry RRs; those must not be accumulated.
				attempt = append(attempt, e.RR...)
			}
			if attemptErr != nil {
				lastErr = fmt.Errorf("failed to read AXFR response for zone %q from %s: %w", zone, nameserver, attemptErr)
				r.listLastErr = lastErr
				continue
			}
			// Clear an earlier attempt's error so the post-loop guard does not report
			// a failure that was already retried away. r.listLastErr is left alone:
			// getNextNameserverFor reads and resets it to drive the "disabled" strategy.
			lastErr = nil
			records = append(records, attempt...)
			// If records were fetched successfully, break out of the loop
			if len(records) > 0 {
				break
			}
		}

		if lastErr != nil {
			r.listLastErr = lastErr
			return nil, provider.NewSoftError(lastErr)
		}
	}

	return records, nil
}

// ApplyChanges applies a given set of changes in a given zone.
func (r *rfc2136Provider) ApplyChanges(_ context.Context, changes *plan.Changes) error {
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

			// calculate corresponding index in the unsplitted UpdateOld for current endpoint ep in chunk
			j := (c * r.batchChangeSize) + i
			r.UpdateRecord(m[zone], changes.UpdateOld[j], ep)
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
		return provider.NewSoftErrorf("RFC2136 had errors in one or more of its batches: %v", errs)
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
		rr, err := r.buildRR(ep, uint32(ttl), target)
		if err != nil {
			return err
		}
		log.Infof("Adding RR: %s", rr)

		m.Insert([]dns.RR{rr})
	}

	return nil
}

func (r *rfc2136Provider) RemoveRecord(m *dns.Msg, ep *endpoint.Endpoint) error {
	log.Debugf("RemoveRecord.ep=%s", ep)
	for _, target := range ep.Targets {
		rr, err := r.buildRR(ep, uint32(ep.RecordTTL), target)
		if err != nil {
			return err
		}
		log.Infof("Removing RR: %s", rr)

		m.Remove([]dns.RR{rr})
	}

	return nil
}

// buildRR renders one target of ep as an RR. Three regimes coexist for TXT
// (#1596): a value read by Records() reuses its character-strings from the
// zone, so deletes match records chunked by someone else; user targets arrive
// canonicalized by AdjustEndpoints (no quoting) and are written verbatim,
// re-chunked to 255-byte character-strings, so ';' is not mistaken for a
// zone-file comment; but the registry serializes heritage labels with
// syntactic quotes (Labels.Serialize(true)) downstream of AdjustEndpoints,
// and that quoting must be interpreted like the legacy dns.NewRR path did, or
// literal quotes land on the wire and the read-back heritage no longer
// parses. Other record types keep the string path.
func (r *rfc2136Provider) buildRR(ep *endpoint.Endpoint, ttl uint32, target string) (dns.RR, error) {
	if ep.RecordType == endpoint.RecordTypeTXT {
		if chunks, ok := r.txtChunks(ep.DNSName, target); ok {
			return &dns.TXT{
				Hdr: dns.RR_Header{
					Name:   dns.Fqdn(ep.DNSName),
					Rrtype: dns.TypeTXT,
					Class:  dns.ClassINET,
					Ttl:    ttl,
				},
				Txt: chunks,
			}, nil
		}
	}

	rr, err := dns.NewRR(fmt.Sprintf("%s %d %s %s", ep.DNSName, ttl, ep.RecordType, target))
	if err != nil {
		return nil, fmt.Errorf("failed to build RR: %w", err)
	}
	return rr, nil
}

// txtChunks returns the character-strings to write for a TXT target, in the
// zone-presentation form dns.TXT.Txt holds, or false when the target is
// zone-quoted and must go through the zone parser.
func (r *rfc2136Provider) txtChunks(name, target string) ([]string, bool) {
	r.txtWireMu.Lock()
	chunks, seen := r.txtWire[newTXTWireKey(name, target)]
	r.txtWireMu.Unlock()
	if seen {
		return chunks, true
	}
	if isZoneQuoted(target) {
		return nil, false
	}
	// Packing decodes the presentation form: chunk the raw canonical value,
	// then escape each chunk so the wire carries the target byte for byte.
	chunks = chunkTXT(target)
	for i := range chunks {
		chunks[i] = escapeTXT(chunks[i])
	}
	return chunks, true
}

// getNextNameserverFor picks the next nameserver to use for the given
// operation according to the configured load balancing strategy. List
// (AXFR) and Send (update) operations rotate and fail over independently,
// so the caller states its intent via op rather than reaching into the
// provider's internal counter/error fields itself.
func (r *rfc2136Provider) getNextNameserverFor(op nameserverOp) string {
	if len(r.nameservers) == 1 {
		return r.nameservers[0]
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	counter, lastErr := r.rotationStateFor(op)

	if *lastErr != nil {
		log.Warnf("Last operation failed for nameserver %s", r.nameservers[*counter])
		log.Warnf("Last operation error message: %v", *lastErr)
	}

	var nameserver string
	switch r.loadBalancingStrategy {
	case "random":
		for {
			nameserver = r.nameservers[r.randGen.Intn(len(r.nameservers))]
			// Ensure that we don't get the same nameserver as the last one
			if nameserver != r.nameservers[*counter] {
				break
			}
		}
	case "round-robin":
		nameserver = r.nameservers[*counter]
		*counter = (*counter + 1) % len(r.nameservers)
	default:
		if *lastErr != nil {
			*counter = (*counter + 1) % len(r.nameservers)
			nameserver = r.nameservers[*counter]
		} else {
			nameserver = r.nameservers[*counter]
		}
	}

	// Last error has been logged, reset it for the next operation
	*lastErr = nil
	return nameserver
}

// rotationStateFor returns the counter and last-error slots that belong to
// the given operation, keeping the per-operation field bookkeeping in one
// place instead of spread across List() and SendMessage().
func (r *rfc2136Provider) rotationStateFor(op nameserverOp) (*int, *error) {
	if op == nameserverOpList {
		return &r.listCounter, &r.listLastErr
	}
	return &r.sendCounter, &r.sendLastErr
}

func (r *rfc2136Provider) SendMessage(msg *dns.Msg) error {
	if r.dryRun {
		log.Debugf("SendMessage.skipped")
		return nil
	}
	log.Debugf("SendMessage")

	var lastErr error
	for i := 0; i < len(r.nameservers); i++ {
		nameserver := r.getNextNameserverFor(nameserverOpSend)
		log.Debugf("Sending message to nameserver: %s", nameserver)

		c, err := makeClient(r, nameserver)
		if err != nil {
			lastErr = fmt.Errorf("error setting up TLS: %w", err)
			r.sendLastErr = lastErr
			continue
		}

		if !r.insecure {
			if r.gssTsig {
				keyName, handle, err := r.KeyData(nameserver)
				if err != nil {
					lastErr = err
					r.sendLastErr = lastErr
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
				r.sendLastErr = lastErr
				continue
			}
			log.Warnf("warn in dns.Client.Exchange: %s", err)
			lastErr = err
			r.sendLastErr = lastErr
			continue
		}
		if resp != nil && resp.Rcode != dns.RcodeSuccess {
			log.Infof("Bad dns.Client.Exchange response: %s", resp)
			lastErr = fmt.Errorf("bad return code: %s", dns.RcodeToString[resp.Rcode])
			r.sendLastErr = lastErr
			continue
		}

		log.Debugf("SendMessage.success")
		return nil
	}

	r.sendLastErr = lastErr
	return provider.NewSoftError(lastErr)
}

func chunkBy(slice []*endpoint.Endpoint, chunkSize int) [][]*endpoint.Endpoint {
	var chunks [][]*endpoint.Endpoint

	for i := 0; i < len(slice); i += chunkSize {
		end := min(i+chunkSize, len(slice))

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
