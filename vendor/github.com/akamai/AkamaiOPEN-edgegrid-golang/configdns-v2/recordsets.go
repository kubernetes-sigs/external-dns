package dnsv2

import (
	"errors"
	"fmt"
	"github.com/akamai/AkamaiOPEN-edgegrid-golang/client-v1"
	edge "github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
	"strconv"
	"sync"
)

var (
	zoneRecordsetsWriteLock sync.Mutex
)

// Recordset Query args struct
type RecordsetQueryArgs struct {
	Page     int
	PageSize int
	Search   string
	ShowAll  bool
	SortBy   string
	Types    string
}

// Recordsets Struct. Used for Create and Update Recordsets
type Recordsets struct {
	Recordsets []Recordset `json:"recordsets"`
}

type Recordset struct {
	Name  string   `json:"name"`
	Type  string   `json:"type"`
	TTL   int      `json:"ttl"`
	Rdata []string `json:"rdata"`
} //`json:"recordsets"`

type MetadataH struct {
	LastPage      int  `json:"lastPage"`
	Page          int  `json:"page"`
	PageSize      int  `json:"pageSize"`
	ShowAll       bool `json:"showAll"`
	TotalElements int  `json:"totalElements"`
} //`json:"metadata"`

type RecordSetResponse struct {
	Metadata   MetadataH   `json:"metadata"`
	Recordsets []Recordset `json:"recordsets"`
}

func NewRecordSetResponse(name string) *RecordSetResponse {
	recordset := &RecordSetResponse{}
	return recordset
}

// Get RecordSets with Query Args. No formatting of arg values!
func GetRecordsets(zone string, queryArgs ...RecordsetQueryArgs) (*RecordSetResponse, error) {

	recordsetResp := NewRecordSetResponse("")

	// construct GET url
	getURL := fmt.Sprintf("/config-dns/v2/zones/%s/recordsets", zone)
	if len(queryArgs) > 1 {
		return nil, errors.New("GetRecordsets QueryArgs invalid.")
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
		req.URL.RawQuery = q.Encode()
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
		err = client.BodyJSON(res, recordsetResp)
		if err != nil {
			return nil, err
		}
		return recordsetResp, nil
	}
}

// Create Recordstes
func (recordsets *Recordsets) Save(zone string, recLock ...bool) error {
	// This lock will restrict the concurrency of API calls
	// to 1 save request at a time. This is needed for the Soa.Serial value which
	// is required to be incremented for every subsequent update to a zone
	// so we have to save just one request at a time to ensure this is always
	// incremented properly

	if localLock(recLock) {
		zoneRecordsetsWriteLock.Lock()
		defer zoneRecordsetsWriteLock.Unlock()
	}

	req, err := client.NewJSONRequest(
		Config,
		"POST",
		"/config-dns/v2/zones/"+zone+"/recordsets",
		recordsets,
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

func (recordsets *Recordsets) Update(zone string, recLock ...bool) error {
	// This lock will restrict the concurrency of API calls
	// to 1 save request at a time. This is needed for the Soa.Serial value which
	// is required to be incremented for every subsequent update to a zone
	// so we have to save just one request at a time to ensure this is always
	// incremented properly

	if localLock(recLock) {
		zoneRecordsetsWriteLock.Lock()
		defer zoneRecordsetsWriteLock.Unlock()
	}

	req, err := client.NewJSONRequest(
		Config,
		"PUT",
		"/config-dns/v2/zones/"+zone+"/recordsets",
		recordsets,
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
