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

package endpoint

import (
	"cmp"
	"fmt"
	"net/netip"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/miekg/dns"
	log "github.com/sirupsen/logrus"
	"k8s.io/utils/set"

	"sigs.k8s.io/external-dns/pkg/events"
)

const (
	// RecordTypeA is a RecordType enum value
	RecordTypeA = "A"
	// RecordTypeAAAA is a RecordType enum value
	RecordTypeAAAA = "AAAA"
	// RecordTypeCNAME is a RecordType enum value
	RecordTypeCNAME = "CNAME"
	// RecordTypeTXT is a RecordType enum value
	RecordTypeTXT = "TXT"
	// RecordTypeSRV is a RecordType enum value
	RecordTypeSRV = "SRV"
	// RecordTypeNS is a RecordType enum value
	RecordTypeNS = "NS"
	// RecordTypePTR is a RecordType enum value
	RecordTypePTR = "PTR"
	// RecordTypeMX is a RecordType enum value
	RecordTypeMX = "MX"
	// RecordTypeNAPTR is a RecordType enum value
	RecordTypeNAPTR = "NAPTR"

	// ProviderSpecificAlias indicates whether a CNAME endpoint maps to a
	// provider-native alias record (e.g. AWS ALIAS).
	ProviderSpecificAlias = "alias"

	// ProviderSpecificRecordType is the provider-specific property name used to
	// request a particular DNS record type (e.g. "ptr") on an endpoint.
	ProviderSpecificRecordType = "record-type"
)

var (
	KnownRecordTypes = []string{
		RecordTypeA,
		RecordTypeAAAA,
		RecordTypeCNAME,
		RecordTypeTXT,
		RecordTypeSRV,
		RecordTypeNS,
		RecordTypePTR,
		RecordTypeMX,
		RecordTypeNAPTR,
	}
)

// TTL is a structure defining the TTL of a DNS record
type TTL int64

// IsConfigured returns true if TTL is configured, false otherwise
func (ttl TTL) IsConfigured() bool {
	return ttl > 0
}

// Targets is a representation of a list of targets for an endpoint.
type Targets []string

// MXTarget represents a single MX (Mail Exchange) record target, including its priority and host.
type MXTarget struct {
	priority uint16
	host     string
}

// NewTargets is a convenience method to create a new Targets object from a vararg of strings.
// Returns a new Targets slice with duplicates removed and elements sorted in order.
func NewTargets(target ...string) Targets {
	return set.New(target...).SortedList()
}

// String returns the targets joined by semicolons.
func (t Targets) String() string {
	return strings.Join(t, ";")
}

// Len returns the number of targets, satisfying sort.Interface.
func (t Targets) Len() int {
	return len(t)
}

// Less reports whether target i sorts before target j, using IP-aware comparison for valid addresses.
func (t Targets) Less(i, j int) bool {
	ipi, err := netip.ParseAddr(t[i])
	if err != nil {
		return t[i] < t[j]
	}

	ipj, err := netip.ParseAddr(t[j])
	if err != nil {
		return t[i] < t[j]
	}

	return ipi.String() < ipj.String()
}

// Swap exchanges targets at positions i and j, satisfying sort.Interface.
func (t Targets) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

// Same compares two Targets and returns true if they are identical (case-insensitive)
func (t Targets) Same(o Targets) bool {
	if len(t) != len(o) {
		return false
	}
	sort.Stable(t)
	sort.Stable(o)

	for i, e := range t {
		if !strings.EqualFold(e, o[i]) {
			// IPv6 can be shortened, so it should be parsed for equality checking
			ipA, err := netip.ParseAddr(e)
			if err != nil {
				log.WithFields(log.Fields{
					"targets":           t,
					"comparisonTargets": o,
				}).Debugf("Couldn't parse %s as an IP address: %v", e, err)
			}

			ipB, err := netip.ParseAddr(o[i])
			if err != nil {
				log.WithFields(log.Fields{
					"targets":           t,
					"comparisonTargets": o,
				}).Debugf("Couldn't parse %s as an IP address: %v", e, err)
			}

			// IPv6 Address Shortener == IPv6 Address Expander
			if ipA.IsValid() && ipB.IsValid() {
				return ipA.String() == ipB.String()
			}
			return false
		}
	}
	return true
}

// IsLess should fulfill the requirement to compare two targets and choose the 'lesser' one.
// In the past target was a simple string so simple string comparison could be used. Now we define 'less'
// as either being the shorter list of targets or where the first entry is less.
// FIXME We really need to define under which circumstances a list Targets is considered 'less'
// than another.
func (t Targets) IsLess(o Targets) bool {
	if len(t) < len(o) {
		return true
	}
	if len(t) > len(o) {
		return false
	}

	sort.Sort(t)
	sort.Sort(o)

	for i, e := range t {
		if e != o[i] {
			// Explicitly prefers IP addresses (e.g. A records) over FQDNs (e.g. CNAMEs).
			// This prevents behavior like `1-2-3-4.example.com` being "less" than `1.2.3.4` when doing lexicographical string comparison.
			ipA, err := netip.ParseAddr(e)
			if err != nil {
				// Ignoring parsing errors is fine due to the empty netip.Addr{} type being an invalid IP,
				// which is checked by IsValid() below. However, still log them in case a provider is experiencing
				// non-obvious issues with the records being created.
				log.WithFields(log.Fields{
					"targets":           t,
					"comparisonTargets": o,
				}).Debugf("Couldn't parse %s as an IP address: %v", e, err)
			}

			ipB, err := netip.ParseAddr(o[i])
			if err != nil {
				log.WithFields(log.Fields{
					"targets":           t,
					"comparisonTargets": o,
				}).Debugf("Couldn't parse %s as an IP address: %v", e, err)
			}

			// If both targets are valid IP addresses, use the built-in Less() function to do the comparison.
			// If one is a valid IP and the other is not, prefer the IP address (consider it "less").
			// If neither is a valid IP, use lexicographical string comparison to determine which string sorts first alphabetically.
			switch {
			case ipA.IsValid() && ipB.IsValid():
				return ipA.Less(ipB)
			case ipA.IsValid() && !ipB.IsValid():
				return true
			case !ipA.IsValid() && ipB.IsValid():
				return false
			default:
				return e < o[i]
			}
		}
	}
	return false
}

// ProviderSpecificProperty holds the name and value of a configuration which is specific to individual DNS providers
type ProviderSpecificProperty struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

// ProviderSpecific holds configuration which is specific to individual DNS providers
type ProviderSpecific []ProviderSpecificProperty

// EndpointKey is the type of a map key for separating endpoints or targets.
type EndpointKey struct {
	DNSName       string
	RecordType    string
	SetIdentifier string
	RecordTTL     TTL
	Target        string
}

func (ep EndpointKey) String() string {
	return fmt.Sprintf(`{%q %q %q "%d" %q}`, ep.DNSName, ep.RecordType, ep.SetIdentifier, ep.RecordTTL, ep.Target)
}

type ObjectRef = events.ObjectReference

// Endpoint is a high-level way of a connection between a service and an IP
// +kubebuilder:object:generate=true
type Endpoint struct {
	// The hostname of the DNS record
	DNSName string `json:"dnsName,omitempty"`
	// The targets the DNS record points to
	Targets Targets `json:"targets,omitempty"`
	// RecordType type of record, e.g. CNAME, A, AAAA, SRV, TXT etc
	RecordType string `json:"recordType,omitempty"`
	// Identifier to distinguish multiple records with the same name and type (e.g. Route53 records with routing policies other than 'simple')
	SetIdentifier string `json:"setIdentifier,omitempty"`
	// TTL for the record
	RecordTTL TTL `json:"recordTTL,omitempty"`
	// Labels stores labels defined for the Endpoint
	// +optional
	Labels Labels `json:"labels,omitempty"`
	// ProviderSpecific stores provider specific config
	// +optional
	ProviderSpecific ProviderSpecific `json:"providerSpecific,omitempty"`
	// refObject stores reference object
	// TODO: should be an array, as endpoints merged from multiple sources may have multiple ref objects
	// +optional
	refObject *ObjectRef `json:"-"`
}

// NewEndpoint initialization method to be used to create an endpoint
func NewEndpoint(dnsName, recordType string, targets ...string) *Endpoint {
	return NewEndpointWithTTL(dnsName, recordType, TTL(0), targets...)
}

// NewEndpointWithTTL initialization method to be used to create an endpoint with a TTL struct
func NewEndpointWithTTL(dnsName, recordType string, ttl TTL, targets ...string) *Endpoint {
	cleanTargets := make([]string, len(targets))
	for idx, target := range targets {
		// Only trim trailing dots for domain name record types, not for TXT or NAPTR records
		// TXT records can contain arbitrary text including multiple dots
		// SRV can contain dots in their target part (RFC2782)
		switch recordType {
		case RecordTypeTXT, RecordTypeNAPTR, RecordTypeSRV:
			cleanTargets[idx] = target
		default:
			cleanTargets[idx] = strings.TrimSuffix(target, ".")
		}
	}

	for label := range strings.SplitSeq(dnsName, ".") {
		if len(label) > 63 {
			log.Errorf("label %s in %s is longer than 63 characters. Cannot create endpoint", label, dnsName)
			return nil
		}
	}

	return &Endpoint{
		DNSName:    strings.TrimSuffix(dnsName, "."),
		Targets:    cleanTargets,
		RecordType: recordType,
		Labels:     NewLabels(),
		RecordTTL:  ttl,
	}
}

// WithSetIdentifier applies the given set identifier to the endpoint.
func (e *Endpoint) WithSetIdentifier(setIdentifier string) *Endpoint {
	e.SetIdentifier = setIdentifier
	return e
}

// WithProviderSpecific attaches a key/value pair to the Endpoint and returns the Endpoint.
// This can be used to pass additional data through the stages of ExternalDNS's Endpoint processing.
// The assumption is that most of the time this will be provider specific metadata that doesn't
// warrant its own field on the Endpoint object itself. It differs from Labels in the fact that it's
// not persisted in the Registry but only kept in memory during a single record synchronization.
func (e *Endpoint) WithProviderSpecific(key, value string) *Endpoint {
	e.SetProviderSpecificProperty(key, value)
	return e
}

// GetProviderSpecificProperty returns the value of a ProviderSpecificProperty if the property exists.
func (e *Endpoint) GetProviderSpecificProperty(key string) (string, bool) {
	if len(e.ProviderSpecific) == 0 {
		return "", false
	}
	for _, providerSpecific := range e.ProviderSpecific {
		if providerSpecific.Name == key {
			return providerSpecific.Value, true
		}
	}
	return "", false
}

// GetBoolProviderSpecificProperty returns a boolean provider-specific property value.
func (e *Endpoint) GetBoolProviderSpecificProperty(key string) (bool, bool) {
	prop, ok := e.GetProviderSpecificProperty(key)
	if !ok {
		return false, false
	}
	switch prop {
	case "true":
		return true, true
	case "false":
		return false, true
	default:
		return false, true
	}
}

// SetProviderSpecificProperty sets the value of a ProviderSpecificProperty.
func (e *Endpoint) SetProviderSpecificProperty(key string, value string) {
	if len(e.ProviderSpecific) == 0 {
		e.ProviderSpecific = append(e.ProviderSpecific, ProviderSpecificProperty{
			Name:  key,
			Value: value,
		})
		return
	}
	for i, providerSpecific := range e.ProviderSpecific {
		if providerSpecific.Name == key {
			e.ProviderSpecific[i] = ProviderSpecificProperty{
				Name:  key,
				Value: value,
			}
			return
		}
	}

	e.ProviderSpecific = append(e.ProviderSpecific, ProviderSpecificProperty{Name: key, Value: value})
}

// DeleteProviderSpecificProperty deletes any ProviderSpecificProperty of the specified name.
func (e *Endpoint) DeleteProviderSpecificProperty(key string) {
	if len(e.ProviderSpecific) == 0 {
		return
	}
	for i, providerSpecific := range e.ProviderSpecific {
		if providerSpecific.Name == key {
			e.ProviderSpecific = append(e.ProviderSpecific[:i], e.ProviderSpecific[i+1:]...)
			return
		}
	}
}

// RetainProviderProperties retains only properties whose name is prefixed with
// "provider/" (e.g. "aws/evaluate-target-health" for provider "aws").
// Properties belonging to other providers are dropped.
// Properties with no provider prefix (e.g. "alias") are provider-agnostic and always retained.
// TODO: cloudflare does not follow the "provider/" prefix convention — its properties use the
// annotation form "external-dns.alpha.kubernetes.io/cloudflare-*", so filtering is skipped for
// cloudflare and all properties are retained (only sorted). This should be removed once cloudflare
// adopts the standard prefix convention.
func (e *Endpoint) RetainProviderProperties(provider string) {
	if len(e.ProviderSpecific) == 0 {
		return
	}
	if provider != "" && provider != "cloudflare" {
		prefix := provider + "/"
		e.ProviderSpecific = slices.DeleteFunc(e.ProviderSpecific, func(prop ProviderSpecificProperty) bool {
			return strings.Contains(prop.Name, "/") && !strings.HasPrefix(prop.Name, prefix)
		})
	}
	slices.SortFunc(e.ProviderSpecific, func(a, b ProviderSpecificProperty) int {
		return cmp.Compare(a.Name, b.Name)
	})
}

// WithLabel adds or updates a label for the Endpoint.
//
// Example usage:
//
//	ep.WithLabel("owner", "user123")
func (e *Endpoint) WithLabel(key, value string) *Endpoint {
	if e.Labels == nil {
		e.Labels = NewLabels()
	}
	e.Labels[key] = value
	return e
}

// WithRefObject sets the reference object for the Endpoint and returns the Endpoint.
// This can be used to associate the Endpoint with a specific Kubernetes object.
func (e *Endpoint) WithRefObject(obj *events.ObjectReference) *Endpoint {
	e.refObject = obj
	return e
}

// RefObject returns the Kubernetes object reference associated with this endpoint.
func (e *Endpoint) RefObject() *events.ObjectReference {
	return e.refObject
}

// Key returns the EndpointKey of the Endpoint.
func (e *Endpoint) Key() EndpointKey {
	return EndpointKey{
		DNSName:       e.DNSName,
		RecordType:    e.RecordType,
		SetIdentifier: e.SetIdentifier,
	}
}

// IsOwnedBy returns true if the endpoint owner label matches the given ownerID, false otherwise
func (e *Endpoint) IsOwnedBy(ownerID string) bool {
	endpointOwner, ok := e.Labels[OwnerLabelKey]
	return ok && endpointOwner == ownerID
}

// GetNakedDomain returns the parent domain of the DNS name (without the first label).
// For example, "www.example.com" returns "example.com".
// For apex/two-label names like "example.com", the full name is returned unchanged.
func (e *Endpoint) GetNakedDomain() string {
	if e.DNSName == "" {
		return ""
	}
	parts := strings.SplitN(e.DNSName, ".", 2)
	if len(parts) < 2 || !strings.Contains(parts[1], ".") {
		return e.DNSName
	}
	return parts[1]
}

// NewPTREndpoint creates a PTR endpoint from a forward IP target and one or more hostnames.
// It computes the reverse DNS name (in-addr.arpa / ip6.arpa) from the target IP.
func NewPTREndpoint(target string, ttl TTL, hostnames ...string) (*Endpoint, error) {
	revAddr, err := dns.ReverseAddr(target)
	if err != nil {
		return nil, fmt.Errorf("failed to compute reverse address for %s: %w", target, err)
	}
	ptrName := strings.TrimSuffix(revAddr, ".")
	return NewEndpointWithTTL(ptrName, RecordTypePTR, ttl, hostnames...), nil
}

// String returns a human-readable representation of the endpoint in zone-file style.
func (e *Endpoint) String() string {
	return fmt.Sprintf("%s %d IN %s %s %s %s", e.DNSName, e.RecordTTL, e.RecordType, e.SetIdentifier, e.Targets, e.ProviderSpecific)
}

// Describe returns a compact summary of the endpoint suitable for logging.
func (e *Endpoint) Describe() string {
	return fmt.Sprintf("record:%s, owner:%s, type:%s, targets:%s", e.DNSName, e.SetIdentifier, e.RecordType, strings.Join(e.Targets, ", "))
}

// FilterEndpointsByOwnerID Apply filter to slice of endpoints and return new filtered slice that includes
// only endpoints that match.
func FilterEndpointsByOwnerID(ownerID string, eps []*Endpoint) []*Endpoint {
	filtered := []*Endpoint{}
	for _, ep := range eps {
		endpointOwner, ok := ep.Labels[OwnerLabelKey]
		switch {
		case !ok:
			log.Debugf(`Skipping endpoint %v because of missing owner label (required: "%s")`, ep, ownerID)
		case endpointOwner != ownerID:
			log.Debugf(`Skipping endpoint %v because owner id does not match (found: "%s", required: "%s")`, ep, endpointOwner, ownerID)
		default:
			filtered = append(filtered, ep)
		}
	}

	return filtered
}

// RemoveDuplicates returns a slice holding the unique endpoints.
// This function doesn't contemplate the Targets of an Endpoint
// as part of the primary Key
func RemoveDuplicates(endpoints []*Endpoint) []*Endpoint {
	visited := make(map[EndpointKey]struct{})
	result := []*Endpoint{}

	for _, ep := range endpoints {
		key := ep.Key()

		if _, found := visited[key]; !found {
			result = append(result, ep)
			visited[key] = struct{}{}
		} else {
			log.Debugf(`Skipping duplicated endpoint: %v`, ep)
		}
	}

	return result
}

// RequestedRecordType returns the value of the "record-type" provider-specific
// property, following the same pattern as the alias accessor.
func (e *Endpoint) RequestedRecordType() (string, bool) {
	return e.GetProviderSpecificProperty(ProviderSpecificRecordType)
}

// TODO: rename to Validate
// CheckEndpoint Check if endpoint is properly formatted according to RFC standards
func (e *Endpoint) CheckEndpoint() bool {
	if !e.supportsAlias() {
		if _, ok := e.GetBoolProviderSpecificProperty(ProviderSpecificAlias); ok {
			log.Warnf("Endpoint %s of type %s does not support alias records", e.DNSName, e.RecordType)
			return false
		}
	}

	switch recordType := e.RecordType; recordType {
	case RecordTypeA, RecordTypeAAAA:
		if !e.isAlias() {
			return e.Targets.ValidateIPRecord(recordType)
		}
	case RecordTypeMX:
		return e.Targets.ValidateMXRecord()
	case RecordTypeSRV:
		return e.Targets.ValidateSRVRecord()
	case RecordTypePTR:
		return e.ValidatePTRRecord()
	}
	return true
}

// isAlias returns true if the endpoint has the alias provider-specific property set to true.
func (e *Endpoint) isAlias() bool {
	val, ok := e.GetBoolProviderSpecificProperty(ProviderSpecificAlias)
	return ok && val
}

func (e *Endpoint) supportsAlias() bool {
	switch e.RecordType {
	case RecordTypeA, RecordTypeAAAA, RecordTypeCNAME:
		return true
	default:
		return false
	}
}

// WithMinTTL sets the endpoint's TTL to the given value if the current TTL is not configured.
func (e *Endpoint) WithMinTTL(ttl int64) {
	if !e.RecordTTL.IsConfigured() && ttl > 0 {
		log.Debugf("Overriding existing TTL %d with new value %d for endpoint %s", e.RecordTTL, ttl, e.DNSName)
		e.RecordTTL = TTL(ttl)
	}
}

// NewMXRecord parses a string representation of an MX record target (e.g., "10 mail.example.com")
// and returns an MXTarget struct. Returns an error if the input is invalid.
func NewMXRecord(target string) (*MXTarget, error) {
	parts := strings.Fields(strings.TrimSpace(target))
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid MX record target: %s. MX records must have a preference value and a host, e.g. '10 example.com'", target)
	}

	priority, err := strconv.ParseUint(parts[0], 10, 16)
	if err != nil {
		return nil, fmt.Errorf("invalid integer value in target: %s", target)
	}

	return &MXTarget{
		priority: uint16(priority),
		host:     parts[1],
	}, nil
}

// GetPriority returns the priority of the MX record target.
func (m *MXTarget) GetPriority() *uint16 {
	return &m.priority
}

// GetHost returns the host of the MX record target.
func (m *MXTarget) GetHost() *string {
	return &m.host
}

// ValidateIPRecord reports whether all targets are valid IP addresses of the given record type (A or AAAA).
func (t Targets) ValidateIPRecord(recordType string) bool {
	for _, target := range t {
		addr, err := netip.ParseAddr(target)
		if err != nil {
			log.Debugf("Invalid %s record target: %s is not a valid IP address", recordType, target)
			return false
		}
		if recordType == RecordTypeA && addr.Is6() {
			log.Debugf("Invalid A record target: %s is an IPv6 address", target)
			return false
		}
		if recordType == RecordTypeAAAA && addr.Is4() {
			log.Debugf("Invalid AAAA record target: %s is an IPv4 address", target)
			return false
		}
	}
	return true
}

// ValidateMXRecord reports whether all targets are valid MX record values (priority + host).
func (t Targets) ValidateMXRecord() bool {
	for _, target := range t {
		_, err := NewMXRecord(target)
		if err != nil {
			log.Debugf("Invalid MX record target: %s. %v", target, err)
			return false
		}
	}

	return true
}

// ValidateSRVRecord reports whether all targets are valid SRV record values (priority weight port host).
func (t Targets) ValidateSRVRecord() bool {
	for _, target := range t {
		// SRV records must have a priority, weight, a port value and a target e.g. "10 5 5060 example.com."
		// as per https://www.rfc-editor.org/rfc/rfc2782.txt the target host has to end with a dot.
		targetParts := strings.Fields(strings.TrimSpace(target))
		if len(targetParts) != 4 {
			log.Debugf("Invalid SRV record target: %s. SRV records must have a priority, weight, a port value and a target host, e.g. '10 5 5060 example.com.'", target)
			return false
		}
		if !strings.HasSuffix(targetParts[3], ".") {
			log.Debugf("Invalid SRV record target: %s. Target host does not end with a dot.'", target)
			return false
		}

		for _, part := range targetParts[:3] {
			_, err := strconv.ParseUint(part, 10, 16)
			if err != nil {
				log.Debugf("Invalid SRV record target: %s. Invalid integer value in target.", target)
				return false
			}
		}
	}
	return true
}

// ValidatePTRRecord checks that a PTR endpoint has a valid reverse DNS name
// (ending in .in-addr.arpa or .ip6.arpa) and that targets are non-empty hostnames.
func (e *Endpoint) ValidatePTRRecord() bool {
	name := strings.ToLower(e.DNSName)
	if !isReverseDNSName(name) {
		log.Debugf("Invalid PTR record: DNSName %q must be a valid reverse DNS name under .in-addr.arpa or .ip6.arpa", e.DNSName)
		return false
	}
	if len(e.Targets) == 0 {
		log.Debugf("Invalid PTR record: at least one target is required for %s", e.DNSName)
		return false
	}
	for _, target := range e.Targets {
		if strings.TrimSpace(target) == "" {
			log.Debugf("Invalid PTR record: target must not be empty for %s", e.DNSName)
			return false
		}
		if _, err := netip.ParseAddr(target); err == nil {
			log.Debugf("Invalid PTR record: target %q for %s must be a hostname, not an IP address", target, e.DNSName)
			return false
		}
	}
	return true
}

// isReverseDNSName checks that name ends with .in-addr.arpa or .ip6.arpa
// and has at least one label before the suffix.
func isReverseDNSName(name string) bool {
	for _, suffix := range []string{".in-addr.arpa", ".ip6.arpa"} {
		if prefix, ok := strings.CutSuffix(name, suffix); ok {
			return len(prefix) > 0 && prefix[0] != '.'
		}
	}
	return false
}

// GetDNSName returns the DNS name of the endpoint.
func (e *Endpoint) GetDNSName() string {
	return e.DNSName
}

// GetRecordType returns the record type of the endpoint.
func (e *Endpoint) GetRecordType() string {
	return e.RecordType
}

// GetRecordTTL returns the TTL of the endpoint as int64.
func (e *Endpoint) GetRecordTTL() int64 {
	return int64(e.RecordTTL)
}

// GetTargets returns the targets of the endpoint.
func (e *Endpoint) GetTargets() []string {
	return e.Targets
}

// GetOwner returns the owner of the endpoint from labels or set identifier.
func (e *Endpoint) GetOwner() string {
	if val, ok := e.Labels[OwnerLabelKey]; ok {
		return val
	}
	return e.SetIdentifier
}
