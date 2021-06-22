// Copyright 2019 nic.at GmbH. All rights reserved.
//
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rc0go

import (
	"encoding/json"
	"gopkg.in/resty.v1"
	"strconv"
)

type AccountStatsService service

//
type TopZone struct {
	ID     int    `json:"id, omitempty"`
	Domain string `json:"domain, omitempty"`
	Count  int    `json:"qc, omitempty"`
}

//
type TopQuery struct {
	Query
	ID     int    `json:"id, omitempty"`
	Domain string `json:"domain, omitempty"`
}

//
type TopNXDomain struct {
	NXDomain
	ID     int    `json:"id, omitempty"`
	Domain string `json:"domain, omitempty"`
}

//
type TopMagnitude struct {
	ID        int     `json:"id, omitempty"`
	Domain    string  `json:"domain, omitempty"`
	Magnitude float32 `json:"mag, omitempty"`
}

type QueryCount struct {
	Date  	string `json:"date, omitempty"`
	Count 	int    `json:"count, omitempty"`
	NXCount int	   `json:"nxcount, omitempty"`
}

type CountryQueryCount struct {
	CountryCode string `json:"cc, omitempty"`
	Country 	string `json:"country, omitempty"`
	Region 		string `json:"region, omitempty"`
	Subregion 	string `json:"subregion, omitempty"`
	QueryCount 	int    `json:"qc, omitempty"`
}

// Return the Top 1000 zones from your account with the highest number of queries in the given past period
//
// rcode0 API doc: https://my.rcodezero.at/api-doc/#api-account-statistics-top-zones-get
func (s *AccountStatsService) TopZones(days int) ([]*TopZone, error) {

	resp, err := accStatsRequest(s, days, RC0AccStatsTopZones)

	if err != nil {
		return nil, err
	}

	var topZones []*TopZone

	err = json.Unmarshal(resp.Body(), &topZones)

	if err != nil {
		return nil, err
	}

	return topZones, nil

}

// Returns the Top 1000 QNAMEs for zones in your account with the highest number of queries in the given past period
//
// rcode0 API doc: https://my.rcodezero.at/api-doc/#api-account-statistics-top-qnames-get
func (s *AccountStatsService) TopQNames(days int) ([]*TopQuery, error) {

	resp, err := accStatsRequest(s, days, RC0AccStatsTopQNames)

	if err != nil {
		return nil, err
	}

	var topQNames []*TopQuery

	err = json.Unmarshal(resp.Body(), &topQNames)

	if err != nil {
		return nil, err
	}

	return topQNames, nil

}

// Returns the Top 1000 QNAMEs with QTYPE answered with NXDOMAIN for zones in your account with the highest number of queries in the given past period
//
// rcode0 API doc: https://my.rcodezero.at/api-doc/#api-account-statistics-top-nxdomains-get
func (s *AccountStatsService) TopNXDomains(days int) ([]*TopNXDomain, error) {

	resp, err := accStatsRequest(s, days, RC0AccStatsTopNXDomains)

	if err != nil {
		return nil, err
	}

	var topNXDomains []*TopNXDomain

	err = json.Unmarshal(resp.Body(), &topNXDomains)

	if err != nil {
		return nil, err
	}

	return topNXDomains, nil

}

// Returns the Top 1000 zone in your account with the highest dns magnitude in the given past period
//
// rcode0 API doc: https://my.rcodezero.at/api-doc/#api-account-statistics-top-dns-magnitude-get
func (s *AccountStatsService) TopMagnitude(days int) ([]*TopMagnitude, error) {

	resp, err := accStatsRequest(s, days, RC0AccStatsTopDNSMagnitude)

	if err != nil {
		return nil, err
	}

	var topMagnitudes []*TopMagnitude

	err = json.Unmarshal(resp.Body(), &topMagnitudes)

	if err != nil {
		return nil, err
	}

	return topMagnitudes, nil

}

// Get the total number of queries and the number of queries answered with NXDOMAIN for all zones in your account for the given past period
//
// rcode0 API doc: https://my.rcodezero.at/api-doc/#api-account-statistics-queries-get
func (s *AccountStatsService) TotalQueryCount(days int) ([]*QueryCount, error) {

	resp, err := accStatsRequest(s, days, RC0AccStatsQueries)

	if err != nil {
		return nil, err
	}

	var queryCounts []*QueryCount

	err = json.Unmarshal(resp.Body(), &queryCounts)

	if err != nil {
		return nil, err
	}

	return queryCounts, nil

}

// Return the number of Queries grouped by the originating country/subregion and region for the given past period
//
// rcode0 API doc: https://my.rcodezero.at/api-doc/#api-account-statistics-countries-get
func (s *AccountStatsService) TotalQueryCountPerCountry(days int) ([]*CountryQueryCount, error) {

	resp, err := accStatsRequest(s, days, RC0AccStatsCountries)

	if err != nil {
		return nil, err
	}

	var queryCounts []*CountryQueryCount

	err = json.Unmarshal(resp.Body(), &queryCounts)

	if err != nil {
		return nil, err
	}

	return queryCounts, nil

}

// Helper method to avoid code duplication
func accStatsRequest(s *AccountStatsService, days int, operation string) (*resty.Response, error) {

	req := s.client.NewRequest()

	d := strconv.Itoa(days)

	if days > 0 {
		req.SetQueryParam(
			"days", d,
		)
	}

	resp, err := req.Get(
			s.client.BaseURL.String() +
				s.client.APIVersion +
				operation,
		)

	if err != nil {
		return nil, err
	}

	return resp, nil
}