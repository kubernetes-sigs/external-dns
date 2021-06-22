// Copyright 2019 nic.at GmbH. All rights reserved.
//
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rc0go

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
)

// ZoneManagementService handles communication with the zone related
// methods of the rcode0 API.
//
// rcode0 API docs: https://my.rcodezero.at/api-doc/#api-zone-management
type ZoneManagementService service

type ZoneManagementServiceInterface interface {
	List(options *ListOptions) (zones []*Zone, page *Page, err error)
	Get(zone string) (*Zone, error)
	Create(zoneCreate *ZoneCreate) (*StatusResponse, error)
	Edit(zone string, zoneEdit *ZoneEdit) (*StatusResponse,  error)
	Delete(zone string) (*StatusResponse, error)
	Transfer(zone string) (*StatusResponse, error)
}

// Zone struct
type Zone struct {
	ID                    int      		`json:"id, omitempty"`
	Domain                string   		`json:"domain, omitempty"`
	Type                  string   		`json:"type, omitempty"`
	Masters               []string 		`json:"masters, omitempty"`
	Serial                int  			`json:"serial, omitempty"`
	LastCheck             string   		`json:"last_check, omitempty"`
	DNSSECStatus          string   		`json:"dnssec_status, omitempty"`
	DNSSECStatusDetail    string   		`json:"dnssec_status_detail, omitempty"`
	DNSSECKSKStatus       string   		`json:"dnssec_ksk_status, omitempty"`
	DNSSECKSKStatusDetail string   		`json:"dnssec_ksk_status_detail, omitempty"`
	DNSSECDS              string   		`json:"dnssec_ds, omitempty"`
	DNSSECDNSKey          string   		`json:"dnssec_dns_key, omitempty"`
	DNSSECSafeToUnsign    string   		`json:"dnssec_sage_to_unsign, omitempty"`
}

// ZoneCreate is used for adding a new zone to rc0
type ZoneCreate struct {
	Domain 	string   `json:"domain, omitempty"`
	Type 	string   `json:"type, omitempty"`
	Masters []string `json:"masters, omitempty"`
}

// ZoneEdit is used to change the type (slave/master) of the zone on rc0
type ZoneEdit struct {
	Type 	string   `json:"type, omitempty"`
	Masters []string `json:"masters, omitempty"`
}

// List all zones
//
// rcode0 API docs: https://my.rcodezero.at/api-doc/#api-zone-management-zones-get
func (s *ZoneManagementService) List(options *ListOptions) (zones []*Zone, page *Page, err error) {

	resp, err := s.client.NewRequest().
		SetQueryParam("page_size", 	options.PageSizeAsString()).
		SetQueryParam("page", 		options.PageNumberAsString()).
		Get(
			s.client.BaseURL.String() +
				s.client.APIVersion +
				RC0Zones,
		)

	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(resp.Body(), &page)

	if err != nil {
		return nil, nil, err
	}

	err = mapstructure.WeakDecode(page.Data, &zones)
	if err != nil {
		return nil, nil, err
	}

	return zones, page, nil
}

// Get a single zone
//
// rcode0 API docs: https://my.rcodezero.at/api-doc/#api-zone-management-zone-details-get
func (s *ZoneManagementService) Get(zone string) (*Zone, error) {

	resp, err := s.client.NewRequest().
		SetPathParams(
			map[string]string{
				"zone": zone,
		}).
		Get(
			s.client.BaseURL.String() +
				s.client.APIVersion +
				RC0Zone,
		)

	if err != nil {
		return nil, err
	}

	var z *Zone

	err = json.Unmarshal(resp.Body(), &z)

	if err != nil {
		return nil, err
	}

	return z, nil

}

// Add a new zone (master or slave) to the anycast network.
//
// rcode0 API docs: https://my.rcodezero.at/api-doc/#api-zone-management-zones-post
func (s *ZoneManagementService) Create(zoneCreate *ZoneCreate) (*StatusResponse, error) {

	resp, err := s.client.NewRequest().
		SetBody(
			map[string]interface{}{
				"domain": zoneCreate.Domain,
				"type": zoneCreate.Type,
				"masters": zoneCreate.Masters,
			}).
		Post(
			s.client.BaseURL.String() +
				s.client.APIVersion +
				RC0Zones,
		)

	if err != nil {
		return nil, err
	}

	return s.client.ResponseToRC0StatusResponse(resp)

}

// Update a zone
//
// rcode0 API docs: https://my.rcodezero.at/api-doc/#api-zone-management-zone-details-put
func (s *ZoneManagementService) Edit(zone string, zoneEdit *ZoneEdit) (*StatusResponse,  error) {

	body := make(map[string]interface{})

	body["type"] = zoneEdit.Type
	body["masters"] = zoneEdit.Masters

	resp, err := s.client.NewRequest().
		SetPathParams(
			map[string]string{
				"zone": zone,
			}).
		SetBody(body).
		Put(
			s.client.BaseURL.String() +
				s.client.APIVersion +
				RC0Zone,
		)

	if err != nil {
		return nil, err
	}

	return s.client.ResponseToRC0StatusResponse(resp)
}

// Removes a zone
//
// rcode0 API docs: https://my.rcodezero.at/api-doc/#api-zone-management-zone-details-delete
func (s *ZoneManagementService) Delete(zone string) (*StatusResponse, error) {

	resp, err := s.client.NewRequest().
		SetPathParams(
			map[string]string{
				"zone": zone,
			}).
		Delete(
			s.client.BaseURL.String() +
				s.client.APIVersion +
				RC0Zone,
		)

	if err != nil {
		return nil, err
	}

	return s.client.ResponseToRC0StatusResponse(resp)

}

// Queues a zone transfer dnssecRequest for the given zone
//
// rcode0 API docs: https://my.rcodezero.at/api-doc/#api-zone-management-zone-transfer-post
func (s *ZoneManagementService) Transfer(zone string) (*StatusResponse, error) {

	resp, err := s.client.NewRequest().
		SetPathParams(
			map[string]string{
				"zone": zone,
			}).
		Post(
			s.client.BaseURL.String() +
				s.client.APIVersion +
				RC0ZoneTransfer,
		)

	if err != nil {
		return nil, err
	}

	return s.client.ResponseToRC0StatusResponse(resp)

}