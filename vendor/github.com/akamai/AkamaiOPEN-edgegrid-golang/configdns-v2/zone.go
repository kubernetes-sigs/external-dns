package dnsv2

import (
	"bytes"
	"fmt"
	"github.com/akamai/AkamaiOPEN-edgegrid-golang/client-v1"
	edge "github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
	"io/ioutil"
	"log"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

var (
	zoneWriteLock sync.Mutex
)

// Zone represents a DNS zone
/*{
    "zone": "river.com",
    "type": "secondary",
    "masters": [
        "1.2.3.4",
        "1.2.3.5"
    ],
    "comment": "Adding bodies of water"
}

{
    "activationState": "ACTIVE",
    "contractId": "C-1FRYVV3",
    "lastActivationDate": "2018-03-20T06:49:30Z",
    "lastModifiedBy": "vwwuq65mjvsrbvcr",
    "lastModifiedDate": "2019-01-28T12:05:13Z",
    "signAndServe": false,
    "type": "PRIMARY",
    "versionId": "2e9aa959-5e99-405c-b233-360639449fa1",
    "zone": "akamaideveloper.net"
}

*/

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

var zoneStructMap map[string]string = map[string]string{
	"Zone":                  "zone",
	"Type":                  "type",
	"Masters":               "masters",
	"Comment":               "comment",
	"SignAndServe":          "signAndServe",
	"SignAndServeAlgorithm": "signAndServeAlgorithm",
	"TsigKey":               "tsigKey",
	"Target":                "target",
	"EndCustomerId":         "endCustomerId",
	"ContractId":            "contractId"}

type ZoneResponse struct {
	Zone                  string   `json:"zone,omitempty"`
	Type                  string   `json:"type,omitempty"`
	Masters               []string `json:"masters,omitempty"`
	Comment               string   `json:"comment,omitempty"`
	SignAndServe          bool     `json:"signAndServe"`
	SignAndServeAlgorithm string   `json:"signAndServeAlgorithm,omitempty"`
	TsigKey               *TSIGKey `json:"tsigKey,omitempty"`
	Target                string   `json:"target,omitempty"`
	EndCustomerId         string   `json:"endCustomerId,omitempty"`
	ContractId            string   `json:"contractId,omitempty"`
	AliasCount            int64    `json:"aliasCount,omitempty"`
	ActivationState       string   `json:"activationState,omitempty"`
	LastActivationDate    string   `json:"lastActivationDate,omitempty"`
	LastModifiedBy        string   `json:"lastModifiedBy,omitempty"`
	LastModifiedDate      string   `json:"lastModifiedDate,omitempty"`
	VersionId             string   `json:"versionId,omitempty"`
}

// Zone List Query args struct
type ZoneListQueryArgs struct {
	ContractIds string
	Page        int
	PageSize    int
	Search      string
	ShowAll     bool
	SortBy      string
	Types       string
}

type ListMetadata struct {
	ContractIds   []string `json:"contractIds"`
	Page          int      `json:"page"`
	PageSize      int      `json:"pageSize"`
	ShowAll       bool     `json:"showAll"`
	TotalElements int      `json:"totalElements"`
} //`json:"metadata"`

type ZoneListResponse struct {
	Metadata *ListMetadata   `json:"metadata,omitempty"`
	Zones    []*ZoneResponse `json:"zones,omitempty"`
}

type ChangeListResponse struct {
	Zone             string `json:"zone,omitempty"`
	ChangeTag        string `json:"changeTag,omitempty"`
	ZoneVersionId    string `json:"zoneVersionId,omitempty"`
	LastModifiedDate string `json:"lastModifiedDate,omitempty"`
	Stale            bool   `json:"stale,omitempty"`
}

// Zones List Response
type ZoneNameListResponse struct {
	Zones []string `json:"zones"`
}

/*
{
    "names": [
        "example.com",
        "bar.example.com"
    ]
}
*/
// returned list of Zone Names
type ZoneNamesResponse struct {
	Names []string `json:"names"`
}

/*
{
    "types": [
        "A",
        "MX"
    ]
}
*/
// Recordset Types for Zone|Name Response
type ZoneNameTypesResponse struct {
	Types []string `json:"types"`
}

// List Zones
func ListZones(queryArgs ...ZoneListQueryArgs) (*ZoneListResponse, error) {

	zoneListResp := &ZoneListResponse{}

	// construct GET url
	getURL := fmt.Sprintf("/config-dns/v2/zones")
	if len(queryArgs) > 1 {
		return nil, fmt.Errorf("ListZones QueryArgs invalid.")
	}

	req, err := client.NewRequest(
		Config,
		"GET",
		getURL,
		nil,
	)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	if len(queryArgs) > 0 {
		if queryArgs[0].Page > 0 {
			q.Add("page", strconv.Itoa(queryArgs[0].Page))
		}
		if queryArgs[0].PageSize > 0 {
			q.Add("pageSize", strconv.Itoa(queryArgs[0].PageSize))
		}
		if queryArgs[0].Search != "" {
			q.Add("search", queryArgs[0].Search)
		}
		q.Add("showAll", strconv.FormatBool(queryArgs[0].ShowAll))
		if queryArgs[0].SortBy != "" {
			q.Add("sortBy", queryArgs[0].SortBy)
		}
		if queryArgs[0].Types != "" {
			q.Add("types", queryArgs[0].Types)
		}
		if queryArgs[0].ContractIds != "" {
			q.Add("contractIds", queryArgs[0].ContractIds)
		}
		req.URL.RawQuery = q.Encode()
	}

	edge.PrintHttpRequest(req, true)

	res, err := client.Do(Config, req)
	if err != nil {
		return nil, err
	}

	edge.PrintHttpResponse(res, true)

	if client.IsError(res) {
		return nil, client.NewAPIError(res)
	} else {
		err = client.BodyJSON(res, zoneListResp)
		if err != nil {
			return nil, err
		}
		return zoneListResp, nil
	}
}

// NewZone creates a new Zone. Supports subset of fields
func NewZone(params ZoneCreate) *ZoneCreate {
	zone := &ZoneCreate{Zone: params.Zone,
		Type:                  params.Type,
		Masters:               params.Masters,
		TsigKey:               params.TsigKey,
		Target:                params.Target,
		EndCustomerId:         params.EndCustomerId,
		ContractId:            params.ContractId,
		Comment:               params.Comment,
		SignAndServe:          params.SignAndServe,
		SignAndServeAlgorithm: params.SignAndServeAlgorithm}
	return zone
}

func NewZoneResponse(zonename string) *ZoneResponse {
	zone := &ZoneResponse{Zone: zonename}
	return zone
}

func NewChangeListResponse(zone string) *ChangeListResponse {
	changelist := &ChangeListResponse{Zone: zone}
	return changelist
}

func NewZoneQueryString(Contract string, group string) *ZoneQueryString {
	zonequerystring := &ZoneQueryString{Contract: Contract, Group: group}
	return zonequerystring
}

// GetZone retrieves a DNS Zone for a given hostname
func GetZone(zonename string) (*ZoneResponse, error) {
	zone := NewZoneResponse(zonename)
	req, err := client.NewRequest(
		Config,
		"GET",
		//"/config-dns/v2/zones/"+zone.Zone,
		"/config-dns/v2/zones/"+zonename,
		nil,
	)
	if err != nil {
		return nil, err
	}

	edge.PrintHttpRequest(req, true)

	res, err := client.Do(Config, req)
	if err != nil {
		return nil, err
	}

	edge.PrintHttpResponse(res, true)

	if client.IsError(res) && res.StatusCode != 404 {
		return nil, client.NewAPIError(res)
	} else if res.StatusCode == 404 {
		return nil, &ZoneError{zoneName: zonename}
	} else {
		err = client.BodyJSON(res, zone)
		if err != nil {
			return nil, err
		}

		return zone, nil
	}
}

// GetZone retrieves a DNS Zone for a given hostname
func GetChangeList(zone string) (*ChangeListResponse, error) {
	changelist := NewChangeListResponse(zone)
	req, err := client.NewRequest(
		Config,
		"GET",
		"/config-dns/v2/changelists/"+zone,
		nil,
	)
	if err != nil {
		return nil, err
	}

	edge.PrintHttpRequest(req, true)

	res, err := client.Do(Config, req)
	if err != nil {
		return nil, err
	}

	edge.PrintHttpResponse(res, true)

	if client.IsError(res) && res.StatusCode != 404 {
		return nil, client.NewAPIError(res)
	} else if res.StatusCode == 404 {
		return nil, &ZoneError{zoneName: zone}
	} else {
		err = client.BodyJSON(res, changelist)
		if err != nil {
			return nil, err
		}

		return changelist, nil
	}
}

// GetZone retrieves a DNS Zone for a given hostname
func GetMasterZoneFile(zone string) (string, error) {

	req, err := client.NewRequest(
		Config,
		"GET",
		"/config-dns/v2/zones/"+zone+"/zone-file",
		nil,
	)
	if err != nil {
		return "", err
	}
	req.Header.Add("Accept", "text/dns")

	edge.PrintHttpRequest(req, true)

	res, err := client.Do(Config, req)
	if err != nil {
		log.Printf("[DEBUG] [Akamai LIB] ZM %v %v", res, err)
		return "", err
	}

	edge.PrintHttpResponse(res, true)

	if client.IsError(res) && res.StatusCode != 404 {
		return "", client.NewAPIError(res)
	} else if res.StatusCode == 404 {
		return "", &ZoneError{zoneName: zone}
	} else {

		bodyBytes, err2 := ioutil.ReadAll(res.Body)
		if err2 != nil {
			return "", err
		}
		masterZone := string(bodyBytes)
		return masterZone, nil
	}
}

// Update Master Zone file
func PostMasterZoneFile(zone string, filedata string) error {

	buf := bytes.NewReader([]byte(filedata))
	req, err := client.NewRequest(
		Config,
		"POST",
		fmt.Sprintf("/config-dns/v2/zones/%s/zone-file", zone),
		buf,
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "text/dns")

	edge.PrintHttpRequest(req, true)

	res, err := client.Do(Config, req)

	// Network error
	if err != nil {
		return &ZoneError{
			zoneName:         zone,
			httpErrorMessage: err.Error(),
			err:              err,
		}
	}

	edge.PrintHttpResponse(res, true)

	// API error
	if client.IsError(res) {
		err := client.NewAPIError(res)
		return &ZoneError{zoneName: zone, apiErrorMessage: err.Detail, err: err}
	}

	return nil
}

// Create a Zone
func (zone *ZoneCreate) Save(zonequerystring ZoneQueryString, clearConn ...bool) error {
	// This lock will restrict the concurrency of API calls
	// to 1 save request at a time. This is needed for the Soa.Serial value which
	// is required to be incremented for every subsequent update to a zone
	// so we have to save just one request at a time to ensure this is always
	// incremented properly
	zoneWriteLock.Lock()
	defer zoneWriteLock.Unlock()

	zoneMap := filterZoneCreate(zone)
	zoneurl := "/config-dns/v2/zones/?contractId=" + zonequerystring.Contract
	if len(zonequerystring.Group) > 0 {
		zoneurl += "&gid=" + zonequerystring.Group
	}
	req, err := client.NewJSONRequest(
		Config,
		"POST",
		zoneurl,
		zoneMap,
	)
	if err != nil {
		return err
	}

	edge.PrintHttpRequest(req, true)

	res, err := client.Do(Config, req)

	// Network error
	if err != nil {
		return &ZoneError{
			zoneName:         zone.Zone,
			httpErrorMessage: err.Error(),
			err:              err,
		}
	}

	edge.PrintHttpResponse(res, true)

	// API error
	if client.IsError(res) {
		err := client.NewAPIError(res)
		return &ZoneError{zoneName: zone.Zone, apiErrorMessage: err.Detail, err: err}
	}

	if strings.ToUpper(zone.Type) == "PRIMARY" {
		// Timing issue with Create immediately followed by SaveChangelist
		for _, clear := range clearConn {
			// should only be one entry
			if clear {
				edge.LogMultiline(edge.EdgegridLog.Traceln, "Clearing Idle Connections")
				client.Client.CloseIdleConnections()
			}
		}
	}

	return nil
}

// Create changelist for the Zone. Side effect is to create default NS SOA records
func (zone *ZoneCreate) SaveChangelist() error {
	// This lock will restrict the concurrency of API calls
	// to 1 save request at a time. This is needed for the Soa.Serial value which
	// is required to be incremented for every subsequent update to a zone
	// so we have to save just one request at a time to ensure this is always
	// incremented properly

	req, err := client.NewJSONRequest(
		Config,
		"POST",
		"/config-dns/v2/changelists/?zone="+zone.Zone,
		"",
	)
	if err != nil {
		return err
	}

	edge.PrintHttpRequest(req, true)

	res, err := client.Do(Config, req)

	// Network error
	if err != nil {
		return &ZoneError{
			zoneName:         zone.Zone,
			httpErrorMessage: err.Error(),
			err:              err,
		}
	}

	edge.PrintHttpResponse(res, true)

	// API error
	if client.IsError(res) {
		err := client.NewAPIError(res)
		return &ZoneError{zoneName: zone.Zone, apiErrorMessage: err.Detail, err: err}
	}

	return nil
}

// Save changelist for the Zone to create default NS SOA records
func (zone *ZoneCreate) SubmitChangelist() error {
	// This lock will restrict the concurrency of API calls
	// to 1 save request at a time. This is needed for the Soa.Serial value which
	// is required to be incremented for every subsequent update to a zone
	// so we have to save just one request at a time to ensure this is always
	// incremented properly

	req, err := client.NewJSONRequest(
		Config,
		"POST",
		"/config-dns/v2/changelists/"+zone.Zone+"/submit",
		"",
	)
	if err != nil {
		return err
	}

	edge.PrintHttpRequest(req, true)

	res, err := client.Do(Config, req)

	// Network error
	if err != nil {
		return &ZoneError{
			zoneName:         zone.Zone,
			httpErrorMessage: err.Error(),
			err:              err,
		}
	}

	edge.PrintHttpResponse(res, true)

	// API error
	if client.IsError(res) {
		err := client.NewAPIError(res)
		return &ZoneError{zoneName: zone.Zone, apiErrorMessage: err.Detail, err: err}
	}

	return nil
}

// Save updates the Zone
func (zone *ZoneCreate) Update(zonequerystring ZoneQueryString) error {
	// This lock will restrict the concurrency of API calls
	// to 1 save request at a time. This is needed for the Soa.Serial value which
	// is required to be incremented for every subsequent update to a zone
	// so we have to save just one request at a time to ensure this is always
	// incremented properly

	zoneMap := filterZoneCreate(zone)
	req, err := client.NewJSONRequest(
		Config,
		"PUT",
		"/config-dns/v2/zones/"+zone.Zone,
		zoneMap,
	)
	if err != nil {
		return err
	}

	edge.PrintHttpRequest(req, true)

	res, err := client.Do(Config, req)

	// Network error
	if err != nil {
		return &ZoneError{
			zoneName:         zone.Zone,
			httpErrorMessage: err.Error(),
			err:              err,
		}
	}

	edge.PrintHttpResponse(res, true)

	// API error
	if client.IsError(res) {
		err := client.NewAPIError(res)
		return &ZoneError{zoneName: zone.Zone, apiErrorMessage: err.Detail, err: err}
	}

	return nil
}

func (zone *ZoneCreate) Delete(zonequerystring ZoneQueryString) error {
	// remove all the records except for SOA
	// which is required and save the zone

	req, err := client.NewJSONRequest(
		Config,
		"DELETE",
		"/config-dns/v2/zones/"+zone.Zone,
		nil,
	)
	if err != nil {
		return err
	}

	edge.PrintHttpRequest(req, true)

	res, err := client.Do(Config, req)

	// Network error
	if err != nil {
		return &ZoneError{
			zoneName:         zone.Zone,
			httpErrorMessage: err.Error(),
			err:              err,
		}
	}

	edge.PrintHttpResponse(res, true)

	// API error
	if client.IsError(res) {
		if res.StatusCode != 404 {
			err := client.NewAPIError(res)
			return &ZoneError{zoneName: zone.Zone, apiErrorMessage: err.Detail, err: err}
		}
	}

	return nil

}

func filterZoneCreate(zone *ZoneCreate) map[string]interface{} {

	zoneType := strings.ToUpper(zone.Type)
	filteredZone := make(map[string]interface{})
	zoneElems := reflect.ValueOf(zone).Elem()
	for i := 0; i < zoneElems.NumField(); i++ {
		varName := zoneElems.Type().Field(i).Name
		varLower := zoneStructMap[varName]
		varValue := zoneElems.Field(i).Interface()
		switch varName {
		case "Target":
			if zoneType == "ALIAS" {
				filteredZone[varLower] = varValue
			}
		case "TsigKey":
			if zoneType == "SECONDARY" {
				filteredZone[varLower] = varValue
			}
		case "Masters":
			if zoneType == "SECONDARY" {
				filteredZone[varLower] = varValue
			}
		case "SignAndServe":
			if zoneType != "ALIAS" {
				filteredZone[varLower] = varValue
			}
		case "SignAndServeAlgorithm":
			if zoneType != "ALIAS" {
				filteredZone[varLower] = varValue
			}
		default:
			filteredZone[varLower] = varValue
		}
	}

	return filteredZone

}

// Validate ZoneCreate Object
func ValidateZone(zone *ZoneCreate) error {

	if len(zone.Zone) == 0 {
		return fmt.Errorf("Zone name is required")
	}
	ztype := strings.ToUpper(zone.Type)
	if ztype != "PRIMARY" && ztype != "SECONDARY" && ztype != "ALIAS" {
		return fmt.Errorf("Invalid zone type")
	}
	if ztype != "SECONDARY" && zone.TsigKey != nil {
		return fmt.Errorf("TsigKey is invalid for %s zone type", ztype)
	}
	if ztype == "ALIAS" {
		if len(zone.Target) == 0 {
			return fmt.Errorf("Target is required for Alias zone type")
		}
		if zone.Masters != nil && len(zone.Masters) > 0 {
			return fmt.Errorf("Masters is invalid for Alias zone type")
		}
		if zone.SignAndServe {
			return fmt.Errorf("SignAndServe is invalid for Alias zone type")
		}
		if len(zone.SignAndServeAlgorithm) > 0 {
			return fmt.Errorf("SignAndServeAlgorithm is invalid for Alias zone type")
		}
		return nil
	}
	// Primary or Secondary
	if len(zone.Target) > 0 {
		return fmt.Errorf("Target is invalid for %s zone type", ztype)
	}
	if zone.Masters != nil && len(zone.Masters) > 0 && ztype == "PRIMARY" {
		return fmt.Errorf("Masters is invalid for Primary zone type")
	}

	return nil

}

// Get Zone's Names
func GetZoneNames(zone string) (*ZoneNamesResponse, error) {

	zoneNameResponse := &ZoneNamesResponse{Names: make([]string, 0)}
	req, err := client.NewRequest(
		Config,
		"GET",
		fmt.Sprintf("/config-dns/v2/zones/%s/names", zone),
		nil,
	)
	if err != nil {
		return nil, err
	}

	edge.PrintHttpRequest(req, true)

	res, err := client.Do(Config, req)
	if err != nil {
		return nil, err
	}

	edge.PrintHttpResponse(res, true)

	if client.IsError(res) && res.StatusCode != 404 {
		return nil, client.NewAPIError(res)
	} else if res.StatusCode == 404 {
		return nil, &ZoneError{zoneName: zone}
	} else {
		err = client.BodyJSON(res, zoneNameResponse)
		if err != nil {
			return nil, err
		}

		return zoneNameResponse, nil
	}
}

// Get Zone Name's record types
func GetZoneNameTypes(zname string, zone string) (*ZoneNameTypesResponse, error) {

	zoneNameTypesResponse := &ZoneNameTypesResponse{Types: make([]string, 0)}
	req, err := client.NewRequest(
		Config,
		"GET",
		fmt.Sprintf("/config-dns/v2/zones/%s/names/%s/types", zone, zname),
		nil,
	)
	if err != nil {
		return nil, err
	}

	edge.PrintHttpRequest(req, true)

	res, err := client.Do(Config, req)
	if err != nil {
		return nil, err
	}

	edge.PrintHttpResponse(res, true)

	if client.IsError(res) && res.StatusCode != 404 {
		return nil, client.NewAPIError(res)
	} else if res.StatusCode == 404 {
		return nil, &ZoneError{zoneName: zone}
	} else {
		err = client.BodyJSON(res, zoneNameTypesResponse)
		if err != nil {
			return nil, err
		}

		return zoneNameTypesResponse, nil
	}
}
