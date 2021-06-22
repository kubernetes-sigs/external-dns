package dnsv2

import (
	"fmt"
	"github.com/akamai/AkamaiOPEN-edgegrid-golang/client-v1"
	edge "github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
)

/*
type ZoneQueryString struct {
	Contract string
	Group    string
}

type ZoneCreate struct {
	Zone                  string   `json:"zone"`
	Type                  string   `json:"type"`
	Masters               []string `json:"masters,omitempty"`
	Comment               string   `json:"comment,omitempty"`
	SignAndServe          bool     `json:"signAndServe"`
	SignAndServeAlgorithm string   `json:"signAndServeAlgorithm,omitempty"`
	TsigKey               *TSIGKey `json:"tsigKey,omitempty"`
	Target                string   `json:"target,omitempty"`
	EndCustomerId         string   `json:"endCustomerId,omitempty"`
	ContractId            string   `json:"contractId,omitempty"`
}
*/

type BulkZonesCreate struct {
	Zones []*ZoneCreate `json:"zones"`
}

type BulkZonesResponse struct {
	RequestId      string `json:"requestId"`
	ExpirationDate string `json:"expirationDate"`
}

type BulkStatusResponse struct {
	RequestId      string `json:"requestId"`
	ZonesSubmitted int    `json:"zonesSubmitted"`
	SuccessCount   int    `json:"successCount"`
	FailureCount   int    `json:"failureCount"`
	IsComplete     bool   `json:"isComplete"`
	ExpirationDate string `json:"expirationDate"`
}

type BulkFailedZone struct {
	Zone          string `json:"zone"`
	FailureReason string `json:"failureReason"`
}

type BulkCreateResultResponse struct {
	RequestId                string            `json:"requestId"`
	SuccessfullyCreatedZones []string          `json:"successfullyCreatedZones"`
	FailedZones              []*BulkFailedZone `JSON:"failedZones"`
}

type BulkDeleteResultResponse struct {
	RequestId                string            `json:"requestId"`
	SuccessfullyDeletedZones []string          `json:"successfullyDeletedZones"`
	FailedZones              []*BulkFailedZone `JSON:"failedZones"`
}

// Get Bulk Zone Create Status
func GetBulkZoneCreateStatus(requestid string) (*BulkStatusResponse, error) {

	bulkzonesurl := fmt.Sprintf("/config-dns/v2/zones/create-requests/%s", requestid)
	req, err := client.NewRequest(
		Config,
		"GET",
		bulkzonesurl,
		nil,
	)
	if err != nil {
		return nil, err
	}

	edge.PrintHttpRequest(req, true)

	res, err := client.Do(Config, req)

	// Network error
	if err != nil {
		return nil, &ZoneError{
			zoneName:         "bulk zone create status",
			httpErrorMessage: err.Error(),
			err:              err,
		}
	}

	edge.PrintHttpResponse(res, true)

	// API error
	if client.IsError(res) {
		err := client.NewAPIError(res)
		return nil, &ZoneError{zoneName: "bulk zone create status", apiErrorMessage: err.Detail, err: err}
	}

	bulkresponse := &BulkStatusResponse{}
	err = client.BodyJSON(res, bulkresponse)
	if err != nil {
		return nil, err
	}

	return bulkresponse, nil
}

// Get Bulk Zone Delete Status
func GetBulkZoneDeleteStatus(requestid string) (*BulkStatusResponse, error) {

	bulkzonesurl := fmt.Sprintf("/config-dns/v2/zones/delete-requests/%s", requestid)
	req, err := client.NewRequest(
		Config,
		"GET",
		bulkzonesurl,
		nil,
	)
	if err != nil {
		return nil, err
	}

	edge.PrintHttpRequest(req, true)

	res, err := client.Do(Config, req)

	// Network error
	if err != nil {
		return nil, &ZoneError{
			zoneName:         "bulk zone delete status",
			httpErrorMessage: err.Error(),
			err:              err,
		}
	}

	edge.PrintHttpResponse(res, true)

	// API error
	if client.IsError(res) {
		err := client.NewAPIError(res)
		return nil, &ZoneError{zoneName: "bulk zone delete status", apiErrorMessage: err.Detail, err: err}
	}

	bulkresponse := &BulkStatusResponse{}
	err = client.BodyJSON(res, bulkresponse)
	if err != nil {
		return nil, err
	}

	return bulkresponse, nil
}

// Get Bulk Zone Create Result
func GetBulkZoneCreateResult(requestid string) (*BulkCreateResultResponse, error) {

	bulkzonesurl := fmt.Sprintf("/config-dns/v2/zones/create-requests/%s/result", requestid)
	req, err := client.NewRequest(
		Config,
		"GET",
		bulkzonesurl,
		nil,
	)
	if err != nil {
		return nil, err
	}

	edge.PrintHttpRequest(req, true)

	res, err := client.Do(Config, req)

	// Network error
	if err != nil {
		return nil, &ZoneError{
			zoneName:         "bulk zone create result",
			httpErrorMessage: err.Error(),
			err:              err,
		}
	}

	edge.PrintHttpResponse(res, true)

	// API error
	if client.IsError(res) {
		err := client.NewAPIError(res)
		return nil, &ZoneError{zoneName: "bulk zone create result", apiErrorMessage: err.Detail, err: err}
	}

	bulkresponse := &BulkCreateResultResponse{}
	err = client.BodyJSON(res, bulkresponse)
	if err != nil {
		return nil, err
	}

	return bulkresponse, nil
}

// Get Bulk Zone Delete Result
func GetBulkZoneDeleteResult(requestid string) (*BulkDeleteResultResponse, error) {

	bulkzonesurl := fmt.Sprintf("/config-dns/v2/zones/delete-requests/%s/result", requestid)
	req, err := client.NewRequest(
		Config,
		"GET",
		bulkzonesurl,
		nil,
	)
	if err != nil {
		return nil, err
	}

	edge.PrintHttpRequest(req, true)

	res, err := client.Do(Config, req)

	// Network error
	if err != nil {
		return nil, &ZoneError{
			zoneName:         "bulk zone delete result",
			httpErrorMessage: err.Error(),
			err:              err,
		}
	}

	edge.PrintHttpResponse(res, true)

	// API error
	if client.IsError(res) {
		err := client.NewAPIError(res)
		return nil, &ZoneError{zoneName: "bulk zone delete result", apiErrorMessage: err.Detail, err: err}
	}

	bulkresponse := &BulkDeleteResultResponse{}
	err = client.BodyJSON(res, bulkresponse)
	if err != nil {
		return nil, err
	}

	return bulkresponse, nil
}

// Bulk Create Zones
func CreateBulkZones(bulkzones *BulkZonesCreate, zonequerystring ZoneQueryString) (*BulkZonesResponse, error) {

	bulkzonesurl := "/config-dns/v2/zones/create-requests?contractId=" + zonequerystring.Contract
	if len(zonequerystring.Group) > 0 {
		bulkzonesurl += "&gid=" + zonequerystring.Group
	}
	req, err := client.NewJSONRequest(
		Config,
		"POST",
		bulkzonesurl,
		bulkzones,
	)
	if err != nil {
		return nil, err
	}

	edge.PrintHttpRequest(req, true)

	res, err := client.Do(Config, req)

	// Network error
	if err != nil {
		return nil, &ZoneError{
			zoneName:         "bulk zone create",
			httpErrorMessage: err.Error(),
			err:              err,
		}
	}

	edge.PrintHttpResponse(res, true)

	// API error
	if client.IsError(res) {
		err := client.NewAPIError(res)
		return nil, &ZoneError{zoneName: "bulk zone create", apiErrorMessage: err.Detail, err: err}
	}

	bulkresponse := &BulkZonesResponse{}
	err = client.BodyJSON(res, bulkresponse)
	if err != nil {
		return nil, err
	}

	return bulkresponse, nil
}

// Bulk Delete Zones
func DeleteBulkZones(zoneslist *ZoneNameListResponse, bypassSafetyChecks ...bool) (*BulkZonesResponse, error) {

	bulkzonesurl := "/config-dns/v2/zones/delete-requests"
	if len(bypassSafetyChecks) > 0 {
		bulkzonesurl += fmt.Sprintf("?bypassSafetyChecks=%t", bypassSafetyChecks[0])
	}

	req, err := client.NewJSONRequest(
		Config,
		"POST",
		bulkzonesurl,
		zoneslist,
	)
	if err != nil {
		return nil, err
	}

	edge.PrintHttpRequest(req, true)

	res, err := client.Do(Config, req)

	// Network error
	if err != nil {
		return nil, &ZoneError{
			zoneName:         "bulk zone delete",
			httpErrorMessage: err.Error(),
			err:              err,
		}
	}

	edge.PrintHttpResponse(res, true)

	// API error
	if client.IsError(res) {
		err := client.NewAPIError(res)
		return nil, &ZoneError{zoneName: "bulk zone delete", apiErrorMessage: err.Detail, err: err}
	}

	bulkresponse := &BulkZonesResponse{}
	err = client.BodyJSON(res, bulkresponse)
	if err != nil {
		return nil, err
	}

	return bulkresponse, nil
}
