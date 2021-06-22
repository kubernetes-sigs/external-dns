// Copyright 2019 nic.at GmbH. All rights reserved.
//
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rc0go

import (
	"encoding/json"
	"errors"
	"gopkg.in/resty.v1"
)

type ZoneStatsService service

//
type PerDay struct {
	Date      string `json:"date, omitempty"`
	Queries   int 	  `json:"qcount, omitempty"`
	NXDomains int    `json:"nxcount, omitempty"`
}

//
type Magnitude struct {
	Date      string `json:"date, omitempty"`
	Magnitude string `json:"mag, omitempty"`
}

//
type Query struct {
	Name  string `json:"qname, omitempty"`
	Type  string `json:"qtype, omitempty"`
	Count int    `json:"qc, omitempty"`
}

//
type NXDomain struct {
	Name  string `json:"qname, omitempty"`
	Type  string `json:"qtype, omitempty"`
	Count int `json:"qc, omitempty"`
}

// Get the total number of queries and the number of queries answered with NXDOMAIN for the given zone for the last 180 days (max.)
//
// rcode0 API doc: https://my.rcodezero.at/api-doc/#api-zone-statistics-queries-get
func (s *ZoneStatsService) Queries(zone string) ([]*PerDay, error) {

	resp, err := statsRequest(s, zone, RC0ZoneStatsQueries)

	if err != nil {
		return nil, err
	}

	var q []*PerDay

	err = json.Unmarshal(resp.Body(), &q)

	if err != nil {
		return nil, err
	}

	return q, nil

}

// Get the DNS magnitude for a given zone for the last 180 days
//
// rcode0 API doc: https://my.rcodezero.at/api-doc/#api-zone-statistics-dns-magnitude-get
func (s *ZoneStatsService) Magnitude(zone string) ([]*Magnitude, error) {

	resp, err := statsRequest(s, zone, RC0ZoneStatsMagnitude)

	if err != nil {
		return nil, err
	}

	var m []*Magnitude

	err = json.Unmarshal(resp.Body(), &m)

	if err != nil {
		return nil, err
	}

	return m, nil

}

// Returns yesterdays top 10 QNAMEs with QTYPE for the given domain
//
// rcode API doc: https://my.rcodezero.at/api-doc/#api-zone-statistics-qnames-get
func (s *ZoneStatsService) QNames(zone string) ([]*Query, error) {

	resp, err := statsRequest(s, zone, RC0ZoneStatsQNames)

	if err != nil {
		return nil, err
	}

	var q []*Query

	err = json.Unmarshal(resp.Body(), &q)

	if err != nil {
		return nil, err
	}

	return q, nil

}

// Returns yesterdays top 10 labels and QTYPE answered with NXDOMAIN
//
// rcode0 API doc: https://my.rcodezero.at/api-doc/#api-zone-statistics-nxdomains-get
func (s *ZoneStatsService) NXDomains(zone string) ([]*NXDomain, error) {

	resp, err := statsRequest(s, zone, RC0ZoneStatsNXDomains)

	if err != nil {
		return nil, err
	}

	var nxd []*NXDomain

	err = json.Unmarshal(resp.Body(), &nxd)

	if err != nil {
		return nil, err
	}

	return nxd, nil

}

func statsRequest(s *ZoneStatsService, zone string, operation string) (*resty.Response, error) {

	resp, err := s.client.NewRequest().
		SetPathParams(
			map[string]string{
				"zone": zone,
			}).
		Get(
			s.client.BaseURL.String() +
				s.client.APIVersion +
				operation,
		)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == 404 {
		status, err := s.client.ResponseToRC0StatusResponse(resp)

		if err != nil {
			return nil, err
		}

		return nil, errors.New(status.Message)
	}

	return resp, nil
}

