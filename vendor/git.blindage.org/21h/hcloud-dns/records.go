package hclouddns

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// GetRecord retrieve one single record by ID.
// Accepts single ID of record.
// Returns HCloudAnswerGetRecord with HCloudRecord and error.
func (d *HCloudClient) GetRecord(ID string) (HCloudAnswerGetRecord, error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://dns.hetzner.com/api/v1/records/%v", ID), nil)
	if err != nil {
		return HCloudAnswerGetRecord{}, err
	}

	req.Header.Add("Auth-API-Token", d.Token)

	resp, err := client.Do(req)
	if err != nil {
		return HCloudAnswerGetRecord{}, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return HCloudAnswerGetRecord{}, err
	}

	answer := HCloudAnswerGetRecord{}

	err = json.Unmarshal([]byte(respBody), &answer)
	if err != nil {
		return HCloudAnswerGetRecord{}, err
	}

	// parse error
	errorResult := HCloudAnswerError{}
	err = json.Unmarshal([]byte(respBody), &errorResult)
	if err != nil {
		//ok, non-standard error, try another form
		errorResultString := HCloudAnswerErrorString{}
		err = json.Unmarshal([]byte(respBody), &errorResultString)
		if err != nil {
			return HCloudAnswerGetRecord{}, err
		}
		errorResult.Error.Message = errorResultString.Error
		errorResult.Error.Code = resp.StatusCode
	}
	answer.Error = errorResult.Error

	return answer, nil
}

// GetRecords retrieve all records of user.
// Accepts HCloudGetRecordsParams struct
// Returns HCloudAnswerGetRecords with array of HCloudRecord, Meta and error.
func (d *HCloudClient) GetRecords(params HCloudGetRecordsParams) (HCloudAnswerGetRecords, error) {

	v := url.Values{}
	if params.Page != "" {
		v.Add("page", params.Page)
	}
	if params.PerPage != "" {
		v.Add("per_page", params.PerPage)
	}
	if params.ZoneID != "" {
		v.Add("zone_id", params.ZoneID)
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://dns.hetzner.com/api/v1/records?%v", v.Encode()), nil)
	if err != nil {
		return HCloudAnswerGetRecords{}, err
	}

	req.Header.Add("Auth-API-Token", d.Token)

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
		return HCloudAnswerGetRecords{}, parseFormErr
	}

	resp, err := client.Do(req)
	if err != nil {
		return HCloudAnswerGetRecords{}, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return HCloudAnswerGetRecords{}, err
	}

	answer := HCloudAnswerGetRecords{}

	err = json.Unmarshal([]byte(respBody), &answer)
	if err != nil {
		return HCloudAnswerGetRecords{}, err
	}

	// parse error
	errorResult := HCloudAnswerError{}
	err = json.Unmarshal([]byte(respBody), &errorResult)
	if err != nil {
		//ok, non-standard error, try another form
		errorResultString := HCloudAnswerErrorString{}
		err = json.Unmarshal([]byte(respBody), &errorResultString)
		if err != nil {
			return HCloudAnswerGetRecords{}, err
		}
		errorResult.Error.Message = errorResultString.Error
		errorResult.Error.Code = resp.StatusCode
	}
	answer.Error = errorResult.Error

	return answer, nil
}

// UpdateRecord makes update of single record by ID.
// Accepts HCloudRecord with fullfilled fields.
// Returns HCloudAnswerGetRecord with HCloudRecord and error.
func (d *HCloudClient) UpdateRecord(record HCloudRecord) (HCloudAnswerGetRecord, error) {

	jsonRecordString, err := json.Marshal(record)
	if err != nil {
		return HCloudAnswerGetRecord{}, err
	}
	body := bytes.NewBuffer(jsonRecordString)

	client := &http.Client{}
	req, err := http.NewRequest("PUT", fmt.Sprintf("https://dns.hetzner.com/api/v1/records/%v", record.ID), body)
	if err != nil {
		return HCloudAnswerGetRecord{}, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Auth-API-Token", d.Token)

	resp, err := client.Do(req)
	if err != nil {
		return HCloudAnswerGetRecord{}, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return HCloudAnswerGetRecord{}, err
	}

	answer := HCloudAnswerGetRecord{}

	err = json.Unmarshal([]byte(respBody), &answer)
	if err != nil {
		return HCloudAnswerGetRecord{}, err
	}

	// parse error
	errorResult := HCloudAnswerError{}
	err = json.Unmarshal([]byte(respBody), &errorResult)
	if err != nil {
		//ok, non-standard error, try another form
		errorResultString := HCloudAnswerErrorString{}
		err = json.Unmarshal([]byte(respBody), &errorResultString)
		if err != nil {
			return HCloudAnswerGetRecord{}, err
		}
		errorResult.Error.Message = errorResultString.Error
		errorResult.Error.Code = resp.StatusCode
	}
	answer.Error = errorResult.Error

	return answer, nil
}

// DeleteRecord remove record by ID.
// Accepts single ID string.
// Returns HCloudAnswerDeleteRecord and error.
func (d *HCloudClient) DeleteRecord(ID string) (HCloudAnswerDeleteRecord, error) {

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", fmt.Sprintf("https://dns.hetzner.com/api/v1/records/%v", ID), nil)
	if err != nil {
		return HCloudAnswerDeleteRecord{}, err
	}

	req.Header.Add("Auth-API-Token", d.Token)

	resp, err := client.Do(req)
	if err != nil {
		return HCloudAnswerDeleteRecord{}, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return HCloudAnswerDeleteRecord{}, err
	}

	answer := HCloudAnswerDeleteRecord{}

	// parse error
	errorResult := HCloudAnswerError{}
	err = json.Unmarshal([]byte(respBody), &errorResult)
	if err != nil {
		//ok, non-standard error, try another form
		errorResultString := HCloudAnswerErrorString{}
		err = json.Unmarshal([]byte(respBody), &errorResultString)
		if err != nil {
			return HCloudAnswerDeleteRecord{}, err
		}
		errorResult.Error.Message = errorResultString.Error
		errorResult.Error.Code = resp.StatusCode
	}
	answer.Error = errorResult.Error

	return answer, nil
}

// CreateRecord creates new single record.
// Accepts HCloudRecord with record to create, of cource no ID.
// Returns HCloudAnswerGetRecord with HCloudRecord and error.
func (d *HCloudClient) CreateRecord(record HCloudRecord) (HCloudAnswerGetRecord, error) {

	jsonRecordString, err := json.Marshal(record)
	if err != nil {
		return HCloudAnswerGetRecord{}, err
	}
	body := bytes.NewBuffer(jsonRecordString)

	client := &http.Client{}
	req, err := http.NewRequest("POST", fmt.Sprintf("https://dns.hetzner.com/api/v1/records"), body)
	if err != nil {
		return HCloudAnswerGetRecord{}, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Auth-API-Token", d.Token)

	resp, err := client.Do(req)
	if err != nil {
		return HCloudAnswerGetRecord{}, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return HCloudAnswerGetRecord{}, err
	}

	answer := HCloudAnswerGetRecord{}

	err = json.Unmarshal([]byte(respBody), &answer)
	if err != nil {
		return HCloudAnswerGetRecord{}, err
	}

	// parse error
	errorResult := HCloudAnswerError{}
	err = json.Unmarshal([]byte(respBody), &errorResult)
	if err != nil {
		//ok, non-standard error, try another form
		errorResultString := HCloudAnswerErrorString{}
		err = json.Unmarshal([]byte(respBody), &errorResultString)
		if err != nil {
			return HCloudAnswerGetRecord{}, err
		}
		errorResult.Error.Message = errorResultString.Error
		errorResult.Error.Code = resp.StatusCode
	}
	answer.Error = errorResult.Error

	return answer, nil
}

// CreateRecordBulk creates many records at once.
// Accepts array of HCloudRecord, converts it to json and makes POST to Hetzner.
// Returns HCloudAnswerCreateRecords with arrays of HCloudRecord with whole list, valid and invalid, error.
func (d *HCloudClient) CreateRecordBulk(record []HCloudRecord) (HCloudAnswerCreateRecords, error) {

	jsonRecordString, err := json.Marshal(record)
	if err != nil {
		return HCloudAnswerCreateRecords{}, err
	}
	body := bytes.NewBuffer(jsonRecordString)

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://dns.hetzner.com/api/v1/api/v1/records/bulk", body)
	if err != nil {
		return HCloudAnswerCreateRecords{}, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Auth-API-Token", d.Token)

	resp, err := client.Do(req)
	if err != nil {
		return HCloudAnswerCreateRecords{}, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return HCloudAnswerCreateRecords{}, err
	}

	answer := HCloudAnswerCreateRecords{}

	err = json.Unmarshal([]byte(respBody), &answer)
	if err != nil {
		return HCloudAnswerCreateRecords{}, err
	}

	// parse error
	errorResult := HCloudAnswerError{}
	err = json.Unmarshal([]byte(respBody), &errorResult)
	if err != nil {
		//ok, non-standard error, try another form
		errorResultString := HCloudAnswerErrorString{}
		err = json.Unmarshal([]byte(respBody), &errorResultString)
		if err != nil {
			return HCloudAnswerCreateRecords{}, err
		}
		errorResult.Error.Message = errorResultString.Error
		errorResult.Error.Code = resp.StatusCode
	}
	answer.Error = errorResult.Error

	return answer, nil
}

// UpdateRecordBulk updates many records at once.
// Accepts array of HCloudRecord, converting to json and makes PUT to Hetzner.
// Returns HCloudAnswerUpdateRecords with arrays of HCloudRecord with updated and failed, error.
func (d *HCloudClient) UpdateRecordBulk(record []HCloudRecord) (HCloudAnswerUpdateRecords, error) {

	jsonRecordString, err := json.Marshal(record)
	if err != nil {
		return HCloudAnswerUpdateRecords{}, err
	}
	body := bytes.NewBuffer(jsonRecordString)

	client := &http.Client{}
	req, err := http.NewRequest("PUT", "https://dns.hetzner.com/api/v1/api/v1/records/bulk", body)
	if err != nil {
		return HCloudAnswerUpdateRecords{}, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Auth-API-Token", d.Token)

	resp, err := client.Do(req)
	if err != nil {
		return HCloudAnswerUpdateRecords{}, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return HCloudAnswerUpdateRecords{}, err
	}

	answer := HCloudAnswerUpdateRecords{}

	err = json.Unmarshal([]byte(respBody), &answer)
	if err != nil {
		return HCloudAnswerUpdateRecords{}, err
	}

	// parse error
	errorResult := HCloudAnswerError{}
	err = json.Unmarshal([]byte(respBody), &errorResult)
	if err != nil {
		//ok, non-standard error, try another form
		errorResultString := HCloudAnswerErrorString{}
		err = json.Unmarshal([]byte(respBody), &errorResultString)
		if err != nil {
			return HCloudAnswerUpdateRecords{}, err
		}
		errorResult.Error.Message = errorResultString.Error
		errorResult.Error.Code = resp.StatusCode
	}
	answer.Error = errorResult.Error

	return answer, nil
}
