package dnsv2

import (
	"fmt"
	"github.com/akamai/AkamaiOPEN-edgegrid-golang/client-v1"
	edge "github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
	"reflect"
	"strings"
	"sync"
)

var (
	tsigWriteLock sync.Mutex
)

// TODO: Add examples?

type TSIGQueryString struct {
	ContractIds []string `json:"contractIds,omitempty"`
	Search      string   `json:"search,omitempty"`
	SortBy      []string `json:"sortBy,omitempty"`
	Gid         int64    `json:"gid,omitempty"`
}

type TSIGKey struct {
	Name      string `json:"name"`
	Algorithm string `json:"algorithm,omitempty"`
	Secret    string `json:"secret,omitempty"`
}

type TSIGKeyResponse struct {
	TSIGKey
	ZoneCount int64 `json:"zoneCount,omitempty"`
}

type TSIGKeyBulkPost struct {
	Key   *TSIGKey `json:"key"`
	Zones []string `json:"zones"`
}

type TSIGZoneAliases struct {
	Aliases []string `json:"aliases"`
}

type TSIGReportMeta struct {
	TotalElements int64    `json:"totalElements"`
	Search        string   `json:"search,omitempty"`
	Contracts     []string `json:"contracts,omitempty"`
	Gid           int64    `json:"gid,omitempty"`
	SortBy        []string `json:"sortBy,omitempty"`
}

type TSIGReportResponse struct {
	Metadata *TSIGReportMeta    `json:"metadata"`
	Keys     []*TSIGKeyResponse `json:"keys,omitempty"`
}

// Return bare bones tsig key struct
func NewTSIGKey(name string) *TSIGKey {
	key := &TSIGKey{Name: name}
	return key
}

// Return empty query string struct. No elements required.
func NewTSIGQueryString() *TSIGQueryString {
	tsigquerystring := &TSIGQueryString{}
	return tsigquerystring
}

func constructTsigQueryString(tsigquerystring *TSIGQueryString) string {

	queryString := ""
	qsElems := reflect.ValueOf(tsigquerystring).Elem()
	for i := 0; i < qsElems.NumField(); i++ {
		varName := qsElems.Type().Field(i).Name
		varValue := qsElems.Field(i).Interface()
		keyVal := fmt.Sprint(varValue)
		switch varName {
		case "ContractIds":
			contractList := ""
			for j, id := range varValue.([]string) {
				contractList += id
				if j < len(varValue.([]string))-1 {
					contractList += "%2C"
				}
			}
			if len(varValue.([]string)) > 0 {
				queryString += "contractIds=" + contractList
			}
		case "SortBy":
			sortByList := ""
			for j, sb := range varValue.([]string) {
				sortByList += sb
				if j < len(varValue.([]string))-1 {
					sortByList += "%2C"
				}
			}
			if len(varValue.([]string)) > 0 {
				queryString += "sortBy=" + sortByList
			}
		case "Search":
			if keyVal != "" {
				queryString += "search=" + keyVal
			}
		case "Gid":
			if varValue.(int64) != 0 {
				queryString += "gid=" + keyVal
			}
		}
		if i < qsElems.NumField()-1 {
			queryString += "&"
		}
	}
	queryString = strings.TrimRight(queryString, "&")
	if len(queryString) > 0 {
		return "?" + queryString
	} else {
		return ""
	}
}

// List TSIG Keys
func ListTsigKeys(tsigquerystring *TSIGQueryString) (*TSIGReportResponse, error) {

	tsigList := &TSIGReportResponse{}
	req, err := client.NewRequest(
		Config,
		"GET",
		fmt.Sprintf("/config-dns/v2/keys%s", constructTsigQueryString(tsigquerystring)),
		nil,
	)
	if err != nil {
		return nil, err
	}

	edge.PrintHttpRequest(req, true)

	res, err := client.Do(Config, req)

	// Network error
	if err != nil {
		return nil, &TsigError{
			keyName:          "TsigKeyList",
			httpErrorMessage: err.Error(),
			err:              err,
		}
	}

	edge.PrintHttpResponse(res, true)

	// API error
	if client.IsError(res) {
		err := client.NewAPIError(res)
		return nil, &TsigError{keyName: "TsigKeyList", apiErrorMessage: err.Detail, err: err}
	}

	err = client.BodyJSON(res, tsigList)
	if err != nil {
		return nil, err
	}

	return tsigList, nil

}

// GetZones retrieves DNS Zones using tsig key
func (tsigKey *TSIGKey) GetZones() (*ZoneNameListResponse, error) {

	zonesList := &ZoneNameListResponse{}
	req, err := client.NewJSONRequest(
		Config,
		"POST",
		"/config-dns/v2/keys/used-by",
		tsigKey,
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
		return nil, &TsigError{keyName: tsigKey.Name}
	} else {
		err = client.BodyJSON(res, zonesList)
		if err != nil {
			return nil, err
		}

		return zonesList, nil
	}
}

// GetZoneKeyAliases retrieves a DNS Zone's aliases
//func GetZoneKeyAliases(zone string) (*TSIGZoneAliases, error) {
//
// There is a discrepency between the technical doc and API operation. API currently returns a zone name list.
// TODO: Reconcile
//
func GetZoneKeyAliases(zone string) (*ZoneNameListResponse, error) {

	zonesList := &ZoneNameListResponse{}
	//zoneAliases :=&TSIGZoneAliases{}
	req, err := client.NewRequest(
		Config,
		"GET",
		fmt.Sprintf("/config-dns/v2/zones/%s/key/used-by", zone),
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
		//err = client.BodyJSON(res, zoneAliases)
		err = client.BodyJSON(res, zonesList)
		if err != nil {
			return nil, err
		}

		//return zoneAliases, nil
		return zonesList, nil
	}
}

// Bulk Zones tsig key update
func (tsigBulk *TSIGKeyBulkPost) BulkUpdate() error {

	req, err := client.NewJSONRequest(
		Config,
		"POST",
		"/config-dns/v2/keys/bulk-update",
		tsigBulk,
	)
	if err != nil {
		return err
	}

	edge.PrintHttpRequest(req, true)

	res, err := client.Do(Config, req)

	edge.PrintHttpResponse(res, true)

	// Network error
	if err != nil {
		return &TsigError{
			keyName:          tsigBulk.Key.Name,
			httpErrorMessage: err.Error(),
			err:              err,
		}
	}

	// API error
	if client.IsError(res) {
		err := client.NewAPIError(res)
		return &TsigError{keyName: tsigBulk.Key.Name, apiErrorMessage: err.Detail, err: err}
	}

	return nil
}

// GetZoneKey retrieves a DNS Zone's key
func GetZoneKey(zone string) (*TSIGKeyResponse, error) {

	zonekey := &TSIGKeyResponse{}
	req, err := client.NewRequest(
		Config,
		"GET",
		fmt.Sprintf("/config-dns/v2/zones/%s/key", zone),
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
		err = client.BodyJSON(res, zonekey)
		if err != nil {
			return nil, err
		}

		return zonekey, nil
	}
}

// Delete tsig key for zone
func DeleteZoneKey(zone string) error {

	req, err := client.NewRequest(
		Config,
		"DELETE",
		fmt.Sprintf("/config-dns/v2/zones/%s/key", zone),
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

// Update tsig key for zone
func (tsigKey *TSIGKey) Update(zone string) error {

	req, err := client.NewJSONRequest(
		Config,
		"PUT",
		fmt.Sprintf("/config-dns/v2/zones/%s/key", zone),
		tsigKey,
	)
	if err != nil {
		return err
	}

	edge.PrintHttpRequest(req, true)

	res, err := client.Do(Config, req)
	// Network error
	if err != nil {
		return &TsigError{
			keyName:          tsigKey.Name,
			httpErrorMessage: err.Error(),
			err:              err,
		}
	}

	edge.PrintHttpResponse(res, true)

	// API error
	if client.IsError(res) {
		err := client.NewAPIError(res)
		return &TsigError{keyName: tsigKey.Name, apiErrorMessage: err.Detail, err: err}
	}

	return nil

}
