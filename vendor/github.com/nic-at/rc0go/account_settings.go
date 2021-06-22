// Copyright 2019 nic.at GmbH. All rights reserved.
//
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rc0go

import "encoding/json"

type AccSettingsService service

type GlobalSetting struct {
	Secondaries []string `json:"secondaries"`
	TSIGOut 	string   `json:"tsigout"`
}

// Get global account settings. Value will be empty if an individual setting is not configured.
//
// rcode0 API doc: https://my.rcodezero.at/api-doc/#api-account-settings-settings-get
func (s *AccSettingsService) Get() (*GlobalSetting, error) {

	resp, err := s.client.NewRequest().
		Get(
			s.client.BaseURL.String() +
				s.client.APIVersion +
				RC0AccSettings,
		)

	if err != nil {
		return nil, err
	}

	var settings *GlobalSetting

	err = json.Unmarshal(resp.Body(), &settings)

	if err != nil {
		return nil, err
	}

	return settings, nil
}

// Configures the account setting “secondaries”.
// Those secondaries will receive notifies and may transfer out all zones under the management of the account
//
// rcode0 API doc: https://my.rcodezero.at/api-doc/#api-account-settings-set-secondaries-put
func (s *AccSettingsService) SetSecondaries(secondaries []string) (*StatusResponse, error) {

	resp, err := s.client.NewRequest().
		SetBody(
			map[string]interface{}{
				"secondaries": secondaries,
		}).
		Put(
			s.client.BaseURL.String() +
				s.client.APIVersion +
				RC0AccSecondaries,
		)

	if err != nil {
		return nil, err
	}

	return s.client.ResponseToRC0StatusResponse(resp)
}

// Removes the configured secondaries
//
// rcode0 API doc: https://my.rcodezero.at/api-doc/#api-account-settings-set-secondaries-delete
func (s *AccSettingsService) RemoveSecondaries() (*StatusResponse, error) {

	resp, err := s.client.NewRequest().
		Delete(
			s.client.BaseURL.String() +
				s.client.APIVersion +
				RC0AccSecondaries,
		)

	if err != nil {
		return nil, err
	}

	return s.client.ResponseToRC0StatusResponse(resp)
}

// Configures the TSIG key used for outbound zone transfers
//
// rcode0 API doc: https://my.rcodezero.at/api-doc/#api-account-settings-settings-tsigout-put
func (s *AccSettingsService) SetTSIG(tsigkey string) (*StatusResponse, error) {

	resp, err := s.client.NewRequest().
		SetBody(
			map[string]interface{}{
				"tsigkey": tsigkey,
			}).
		Put(
			s.client.BaseURL.String() +
				s.client.APIVersion +
				RC0AccTsigout,
		)

	if err != nil {
		return nil, err
	}

	return s.client.ResponseToRC0StatusResponse(resp)
}

// Removes the configured TSIG key
//
// rcode0 API doc: https://my.rcodezero.at/api-doc/#api-account-settings-settings-tsigout-delete
func (s *AccSettingsService) RemoveTSIG() (*StatusResponse, error) {

	resp, err := s.client.NewRequest().
		Delete(
			s.client.BaseURL.String() +
				s.client.APIVersion +
				RC0AccTsigout,
		)

	if err != nil {
		return nil, err
	}

	return s.client.ResponseToRC0StatusResponse(resp)
}