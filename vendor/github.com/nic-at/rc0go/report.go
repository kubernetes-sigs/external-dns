// Copyright 2019 nic.at GmbH. All rights reserved.
//
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rc0go

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
)

type ReportService service

type ProbZone struct {
	Domain    string 		`json:"domain, omitempty"`
	Type 	  string 		`json:"type, omitempty"`
	DNSSEC    string 		`json:"dnssec, omitempty"`
	Created   string 		`json:"created, omitempty"`
	LastCheck string 	 	`json:"last_check, omitempty"`
	Serial    int 			`json:"serial, omitempty"`
	Masters   []string 	 	`json:"masters, omitempty"`
}

// Returns the list of problematic zones (=zones with could not be checked or transferred successfully from the master server)
//
// rcode0 API doc: https://my.rcodezero.at/api-doc/#api-reports-reports-problematiczones-get
func (s *ReportService) ProblematicZones() ([]*ProbZone, *Page, error) {

	resp, err := s.client.NewRequest().
		//SetQueryParam("page_size", options.GetPageNumberAsString()). @todo: add this
		//SetQueryParam("page", options.GetPageNumberAsString()).
		Get(
			s.client.BaseURL.String() +
				s.client.APIVersion +
				RC0ReportsProblematiczones,
		)

	if err != nil {
		return nil, nil, err
	}

	var page *Page

	err = json.Unmarshal(resp.Body(), &page)

	if err != nil {
		return nil, nil, err
	}

	var zones []*ProbZone

	err = mapstructure.WeakDecode(page.Data, &zones)

	if err != nil {
		return nil, nil, err
	}

	return zones, page, nil
}

// @todo
//func (s *ReportService) NXDomains(day string) (?, error) {}
//func (s *ReportService) Accounting(month string) (?, error) {}
//func (s *ReportService) Queryrates(month string) (?, error) {}