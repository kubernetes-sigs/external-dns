// Copyright 2019 nic.at GmbH. All rights reserved.
//
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rc0go

type DNSSECService service

// Starts DNSSEC signing of a zone
//
// rcode0 API doc: https://my.rcodezero.at/api-doc/#api-zone-management-dnssec-sign-zone-post
func (s *DNSSECService) Sign(zone string) (*StatusResponse, error) {

	return dnssecRequest(s, zone, RC0ZoneDNSSecSign)

}

// Stops DNSSEC signing of a zone, reverting the zone to unsigned
//
// rcode0 API doc: https://my.rcodezero.at/api-doc/#api-zone-management-dnssec-unsign-zone-post
func (s *DNSSECService) Unsign(zone string) (*StatusResponse, error) {

	return dnssecRequest(s, zone, RC0ZoneDNSSecUnsign)

}

// Starts a DNSSEC key rollover
//
// rcode0 API doc: https://my.rcodezero.at/api-doc/#api-zone-management-dnssec-key-rollover-post
func (s *DNSSECService) KeyRollover(zone string) (*StatusResponse, error) {

	return dnssecRequest(s, zone, RC0ZoneDNSSecKeyRollover)

}

// Acknowledges a DS update
//
// rcode0 API doc: https://my.rcodezero.at/api-doc/#api-zone-management-dnssec-acknowledge-ds-update-post
func (s *DNSSECService) DSUpdate(zone string) (*StatusResponse, error) {

	return dnssecRequest(s, zone, RC0ZoneDNSSecDSUpdate)

}

// Simulates that the DS records of all KSKs of a certain domain were seen in the parent zone.
// This allows to test key rollovers even if the DS of the currently active KSK was not seen in the parent zone.
// A DSSEEN event will be pushed ot the message queue.
// (available on test system only)
//
// rcode0 API doc: https://my.rcodezero.at/api-doc/#api-zone-management-simulate-dnssec-event-dsseen-post
func (s *DNSSECService) SimulateDSSEENEvent(zone string) (*StatusResponse, error) {

	return dnssecRequest(s, zone, RC0ZoneDNSSecDSSEEN)

}

// Simulates that the DS records of all KSKs of a certain domain were removed from the parent zone.
// This allows to subsequently “unsign” a domain.
// (available on test system only)
//
// rcode0 API doc: https://my.rcodezero.at/api-doc/#api-zone-management-simulate-dnssec-event-dsremoved-post
func (s *DNSSECService) SimulateDSREMOVEDEvent(zone string) (*StatusResponse, error) {

	return dnssecRequest(s, zone, RC0ZoneDNSSecDSREMOVED)

}

// Helper method to avoid code duplication
func dnssecRequest(s *DNSSECService, zone string, operation string) (*StatusResponse, error) {

	resp, err := s.client.NewRequest().
		SetPathParams(
			map[string]string{
				"zone": zone,
			}).
		Post(
			s.client.BaseURL.String() +
				s.client.APIVersion +
				operation,
		)

	if err != nil {
		return nil, err
	}

	return s.client.ResponseToRC0StatusResponse(resp)

}