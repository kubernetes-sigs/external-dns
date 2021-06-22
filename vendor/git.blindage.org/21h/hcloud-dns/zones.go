package hclouddns

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// GetZone retrieve one single zone by ID.
// Accepts zone ID string.
// Returns HCloudAnswerGetZone with HCloudZone and error
func (d *HCloudClient) GetZone(ID string) (HCloudAnswerGetZone, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://dns.hetzner.com/api/v1/zones/%v", ID), nil)
	if err != nil {
		return HCloudAnswerGetZone{}, err
	}

	req.Header.Add("Auth-API-Token", d.Token)

	resp, err := client.Do(req)
	if err != nil {
		return HCloudAnswerGetZone{}, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return HCloudAnswerGetZone{}, err
	}

	answer := HCloudAnswerGetZone{}

	err = json.Unmarshal([]byte(respBody), &answer)
	if err != nil {
		return HCloudAnswerGetZone{}, err
	}

	// parse error
	errorResult := HCloudAnswerError{}
	err = json.Unmarshal([]byte(respBody), &errorResult)
	if err != nil {
		//ok, non-standard error, try another form
		errorResultString := HCloudAnswerErrorString{}
		err = json.Unmarshal([]byte(respBody), &errorResultString)
		if err != nil {
			return HCloudAnswerGetZone{}, err
		}
		errorResult.Error.Message = errorResultString.Error
		errorResult.Error.Code = resp.StatusCode
	}
	answer.Error = errorResult.Error

	return answer, nil
}

// GetZones retrieve all zones of user.
// Accepts exact name as string, search name with partial name.
// Returns HCloudAnswerGetZones with array of HCloudZone, Meta and error.
func (d *HCloudClient) GetZones(params HCloudGetZonesParams) (HCloudAnswerGetZones, error) {

	v := url.Values{}
	if params.Name != "" {
		v.Add("name", params.Name)
	}
	if params.SearchName != "" {
		v.Add("search_name", params.SearchName)
	}
	if params.Page != "" {
		v.Add("page", params.Page)
	}
	if params.PerPage != "" {
		v.Add("per_page", params.PerPage)
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://dns.hetzner.com/api/v1/zones?%v", v.Encode()), nil)
	if err != nil {
		return HCloudAnswerGetZones{}, err
	}

	req.Header.Add("Auth-API-Token", d.Token)

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
		return HCloudAnswerGetZones{}, parseFormErr
	}

	resp, err := client.Do(req)
	if err != nil {
		return HCloudAnswerGetZones{}, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return HCloudAnswerGetZones{}, err
	}

	answer := HCloudAnswerGetZones{}

	err = json.Unmarshal([]byte(respBody), &answer)
	if err != nil {
		return HCloudAnswerGetZones{}, err
	}

	// parse error
	errorResult := HCloudAnswerError{}
	err = json.Unmarshal([]byte(respBody), &errorResult)
	if err != nil {
		//ok, non-standard error, try another form
		errorResultString := HCloudAnswerErrorString{}
		err = json.Unmarshal([]byte(respBody), &errorResultString)
		if err != nil {
			return HCloudAnswerGetZones{}, err
		}
		errorResult.Error.Message = errorResultString.Error
		errorResult.Error.Code = resp.StatusCode
	}
	answer.Error = errorResult.Error

	return answer, nil
}

// UpdateZone makes update of single zone by ID.
// Accepts HCloudZone with fullfilled fields.
// Returns HCloudAnswerGetZone with HCloudZone and error.
func (d *HCloudClient) UpdateZone(zone HCloudZone) (HCloudAnswerGetZone, error) {

	jsonZoneString, err := json.Marshal(zone)
	if err != nil {
		return HCloudAnswerGetZone{}, err
	}
	body := bytes.NewBuffer(jsonZoneString)

	client := &http.Client{}
	req, err := http.NewRequest("PUT", fmt.Sprintf("https://dns.hetzner.com/api/v1/zones/%v", zone.ID), body)
	if err != nil {
		return HCloudAnswerGetZone{}, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Auth-API-Token", d.Token)

	resp, err := client.Do(req)
	if err != nil {
		return HCloudAnswerGetZone{}, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return HCloudAnswerGetZone{}, err
	}

	answer := HCloudAnswerGetZone{}

	err = json.Unmarshal([]byte(respBody), &answer)
	if err != nil {
		return HCloudAnswerGetZone{}, err
	}

	// parse error
	errorResult := HCloudAnswerError{}
	err = json.Unmarshal([]byte(respBody), &errorResult)
	if err != nil {
		//ok, non-standard error, try another form
		errorResultString := HCloudAnswerErrorString{}
		err = json.Unmarshal([]byte(respBody), &errorResultString)
		if err != nil {
			return HCloudAnswerGetZone{}, err
		}
		errorResult.Error.Message = errorResultString.Error
		errorResult.Error.Code = resp.StatusCode
	}
	answer.Error = errorResult.Error

	return answer, nil
}

// DeleteZone remove zone by ID.
// Accepts single ID string.
// Returns HCloudAnswerDeleteZone with error.
func (d *HCloudClient) DeleteZone(ID string) (HCloudAnswerDeleteZone, error) {

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", fmt.Sprintf("https://dns.hetzner.com/api/v1/zones/%v", ID), nil)
	if err != nil {
		return HCloudAnswerDeleteZone{}, err
	}

	req.Header.Add("Auth-API-Token", d.Token)

	resp, err := client.Do(req)
	if err != nil {
		return HCloudAnswerDeleteZone{}, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return HCloudAnswerDeleteZone{}, err
	}

	answer := HCloudAnswerDeleteZone{}

	// parse error
	errorResult := HCloudAnswerError{}
	err = json.Unmarshal([]byte(respBody), &errorResult)
	if err != nil {
		//ok, non-standard error, try another form
		errorResultString := HCloudAnswerErrorString{}
		err = json.Unmarshal([]byte(respBody), &errorResultString)
		if err != nil {
			return HCloudAnswerDeleteZone{}, err
		}
		errorResult.Error.Message = errorResultString.Error
		errorResult.Error.Code = resp.StatusCode
	}
	answer.Error = errorResult.Error

	return answer, nil
}

// CreateZone creates new single zone.
// Accepts HCloudZone with record to create, of cource no ID.
// Returns HCloudAnswerGetZone with HCloudZone and error.
func (d *HCloudClient) CreateZone(zone HCloudZone) (HCloudAnswerGetZone, error) {

	jsonZoneString, err := json.Marshal(zone)
	if err != nil {
		return HCloudAnswerGetZone{}, err
	}
	body := bytes.NewBuffer(jsonZoneString)

	client := &http.Client{}
	req, err := http.NewRequest("POST", fmt.Sprintf("https://dns.hetzner.com/api/v1/zones"), body)
	if err != nil {
		return HCloudAnswerGetZone{}, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Auth-API-Token", d.Token)

	resp, err := client.Do(req)
	if err != nil {
		return HCloudAnswerGetZone{}, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return HCloudAnswerGetZone{}, err
	}

	answer := HCloudAnswerGetZone{}

	err = json.Unmarshal([]byte(respBody), &answer)
	if err != nil {
		return HCloudAnswerGetZone{}, err
	}

	// parse error
	errorResult := HCloudAnswerError{}
	err = json.Unmarshal([]byte(respBody), &errorResult)
	if err != nil {
		//ok, non-standard error, try another form
		errorResultString := HCloudAnswerErrorString{}
		err = json.Unmarshal([]byte(respBody), &errorResultString)
		if err != nil {
			return HCloudAnswerGetZone{}, err
		}
		errorResult.Error.Message = errorResultString.Error
		errorResult.Error.Code = resp.StatusCode
	}
	answer.Error = errorResult.Error

	return answer, nil
}

// ImportZoneString imports single zone from imported text.
// Accepts ID and zonePlainText strings.
// Returns HCloudAnswerGetZone with HCloudZone and error.
func (d *HCloudClient) ImportZoneString(zoneID string, zonePlainText string) (HCloudAnswerGetZone, error) {

	body := strings.NewReader(zonePlainText)

	client := &http.Client{}
	req, err := http.NewRequest("POST", fmt.Sprintf("https://dns.hetzner.com/api/v1/zones/%v/import", zoneID), body)
	if err != nil {
		return HCloudAnswerGetZone{}, err
	}

	req.Header.Add("Content-Type", "text/plain")
	req.Header.Add("Auth-API-Token", d.Token)

	resp, err := client.Do(req)
	if err != nil {
		return HCloudAnswerGetZone{}, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return HCloudAnswerGetZone{}, err
	}

	answer := HCloudAnswerGetZone{}

	err = json.Unmarshal([]byte(respBody), &answer)
	if err != nil {
		return HCloudAnswerGetZone{}, err
	}

	// parse error
	errorResult := HCloudAnswerError{}
	err = json.Unmarshal([]byte(respBody), &errorResult)
	if err != nil {
		//ok, non-standard error, try another form
		errorResultString := HCloudAnswerErrorString{}
		err = json.Unmarshal([]byte(respBody), &errorResultString)
		if err != nil {
			return HCloudAnswerGetZone{}, err
		}
		errorResult.Error.Message = errorResultString.Error
		errorResult.Error.Code = resp.StatusCode
	}
	answer.Error = errorResult.Error

	return answer, nil
}

// ExportZoneToString exports single zone from imported text.
// Accepts ID and zonePlainText strings.
// Returns HCloudAnswerGetZonePlainText with HCloudZone and error.
func (d *HCloudClient) ExportZoneToString(zoneID string) (HCloudAnswerGetZonePlainText, error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://dns.hetzner.com/api/v1/zones/%v/export", zoneID), nil)
	if err != nil {
		return HCloudAnswerGetZonePlainText{}, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	req.Header.Add("Auth-API-Token", d.Token)

	resp, err := client.Do(req)
	if err != nil {
		return HCloudAnswerGetZonePlainText{}, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return HCloudAnswerGetZonePlainText{}, err
	}

	answer := HCloudAnswerGetZonePlainText{}
	answer.ZonePlainText = string(respBody)

	return answer, nil
}

// ValidateZoneString validate single zone from imported text.
// Accepts ID and zonePlainText strings.
// Returns HCloudAnswerZoneValidate with HCloudZone and error.
func (d *HCloudClient) ValidateZoneString(zonePlainText string) (HCloudAnswerZoneValidate, error) {

	body := strings.NewReader(zonePlainText)

	client := &http.Client{}
	req, err := http.NewRequest("POST", fmt.Sprintf("https://dns.hetzner.com/api/v1/zones/file/validate"), body)
	if err != nil {
		return HCloudAnswerZoneValidate{}, err
	}

	req.Header.Add("Content-Type", "text/plain")
	req.Header.Add("Auth-API-Token", d.Token)

	resp, err := client.Do(req)
	if err != nil {
		return HCloudAnswerZoneValidate{}, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return HCloudAnswerZoneValidate{}, err
	}

	answer := HCloudAnswerZoneValidate{}

	err = json.Unmarshal([]byte(respBody), &answer)
	if err != nil {
		return HCloudAnswerZoneValidate{}, err
	}

	// parse error
	errorResult := HCloudAnswerError{}
	err = json.Unmarshal([]byte(respBody), &errorResult)
	if err != nil {
		//ok, non-standard error, try another form
		errorResultString := HCloudAnswerErrorString{}
		err = json.Unmarshal([]byte(respBody), &errorResultString)
		if err != nil {
			return HCloudAnswerZoneValidate{}, err
		}
		errorResult.Error.Message = errorResultString.Error
		errorResult.Error.Code = resp.StatusCode
	}
	answer.Error = errorResult.Error

	return answer, nil
}
