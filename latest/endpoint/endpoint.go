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
	"fmt"
	"net/netip"
	"sort"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
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
)

var (
	KnownRecordTypes = []string{
		RecordTypeA,
		RecordTypeAAAA,
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

// NewTargets is a convenience method to create a new Targets object from a vararg of strings
func NewTargets(target ...string) Targets {
	t := make(Targets, 0, len(target))
	t = append(t, target...)
	return t
}

func (t Targets) String() string {
	return strings.Join(t, ";")
}

func (t Targets) Len() int {
	return len(t)
}

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

func (t Targets) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

// Same compares to Targets and returns true if they are identical (case-insensitive)
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
}

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
}

// NewEndpoint initialization method to be used to create an endpoint
func NewEndpoint(dnsName, recordType string, targets ...string) *Endpoint {
	return NewEndpointWithTTL(dnsName, recordType, TTL(0), targets...)
}

// NewEndpointWithTTL initialization method to be used to create an endpoint with a TTL struct
func NewEndpointWithTTL(dnsName, recordType string, ttl TTL, targets ...string) *Endpoint {
	cleanTargets := make([]string, len(targets))
	for idx, target := range targets {
		cleanTargets[idx] = strings.TrimSuffix(target, ".")
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
	for _, providerSpecific := range e.ProviderSpecific {
		if providerSpecific.Name == key {
			return providerSpecific.Value, true
		}
	}
	return "", false
}

// SetProviderSpecificProperty sets the value of a ProviderSpecificProperty.
func (e *Endpoint) SetProviderSpecificProperty(key string, value string) {
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
	for i, providerSpecific := range e.ProviderSpecific {
		if providerSpecific.Name == key {
			e.ProviderSpecific = append(e.ProviderSpecific[:i], e.ProviderSpecific[i+1:]...)
			return
		}
	}
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

func (e *Endpoint) String() string {
	return fmt.Sprintf("%s %d IN %s %s %s %s", e.DNSName, e.RecordTTL, e.RecordType, e.SetIdentifier, e.Targets, e.ProviderSpecific)
}

// Apply filter to slice of endpoints and return new filtered slice that includes
// only endpoints that match.
func FilterEndpointsByOwnerID(ownerID string, eps []*Endpoint) []*Endpoint {
	filtered := []*Endpoint{}
	for _, ep := range eps {
		if endpointOwner, ok := ep.Labels[OwnerLabelKey]; !ok || endpointOwner != ownerID {
			log.Debugf(`Skipping endpoint %v because owner id does not match, found: "%s", required: "%s"`, ep, endpointOwner, ownerID)
		} else {
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

// CheckEndpoint Check if endpoint is properly formatted according to RFC standards
func (e *Endpoint) CheckEndpoint() bool {
	switch recordType := e.RecordType; recordType {
	case RecordTypeMX:
		return e.Targets.ValidateMXRecord()
	case RecordTypeSRV:
		return e.Targets.ValidateSRVRecord()
	}
	return true
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

func (t Targets) ValidateSRVRecord() bool {
	for _, target := range t {
		// SRV records must have a priority, weight, and port value, e.g. "10 5 5060 example.com"
		// as per https://www.rfc-editor.org/rfc/rfc2782.txt
		targetParts := strings.Fields(strings.TrimSpace(target))
		if len(targetParts) != 4 {
			log.Debugf("Invalid SRV record target: %s. SRV records must have a priority, weight, and port value, e.g. '10 5 5060 example.com'", target)
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
