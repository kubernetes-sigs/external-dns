package dnsv2

import (
	"github.com/akamai/AkamaiOPEN-edgegrid-golang/client-v1"
	edge "github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"

	"sync"
)

// The record types implemented and their fields are as defined here
// https://developer.akamai.com/api/luna/config-dns/data.html

type RecordBody struct {
	Name       string `json:"name,omitempty"`
	RecordType string `json:"type,omitempty"`
	TTL        int    `json:"ttl,omitempty"`
	// Active field no longer used in v2
	Active bool     `json:"active,omitempty"`
	Target []string `json:"rdata,omitempty"`
	// Remaining Fields are not used in the v2 API
	Subtype             int    `json:"subtype,omitempty"`                //AfsdbRecord
	Flags               int    `json:"flags,omitempty"`                  //DnskeyRecord Nsec3paramRecord
	Protocol            int    `json:"protocol,omitempty"`               //DnskeyRecord
	Algorithm           int    `json:"algorithm,omitempty"`              //DnskeyRecord DsRecord Nsec3paramRecord RrsigRecord SshfpRecord
	Key                 string `json:"key,omitempty"`                    //DnskeyRecord
	Keytag              int    `json:"keytag,omitempty"`                 //DsRecord RrsigRecord
	DigestType          int    `json:"digest_type,omitempty"`            //DsRecord
	Digest              string `json:"digest,omitempty"`                 //DsRecord
	Hardware            string `json:"hardware,omitempty"`               //HinfoRecord
	Software            string `json:"software,omitempty"`               //HinfoRecord
	Priority            int    `json:"priority,omitempty"`               //MxRecord SrvRecord
	Order               uint16 `json:"order,omitempty"`                  //NaptrRecord
	Preference          uint16 `json:"preference,omitempty"`             //NaptrRecord
	FlagsNaptr          string `json:"flags,omitempty"`                  //NaptrRecord
	Service             string `json:"service,omitempty"`                //NaptrRecord
	Regexp              string `json:"regexp,omitempty"`                 //NaptrRecord
	Replacement         string `json:"replacement,omitempty"`            //NaptrRecord
	Iterations          int    `json:"iterations,omitempty"`             //Nsec3Record Nsec3paramRecord
	Salt                string `json:"salt,omitempty"`                   //Nsec3Record Nsec3paramRecord
	NextHashedOwnerName string `json:"next_hashed_owner_name,omitempty"` //Nsec3Record
	TypeBitmaps         string `json:"type_bitmaps,omitempty"`           //Nsec3Record
	Mailbox             string `json:"mailbox,omitempty"`                //RpRecord
	Txt                 string `json:"txt,omitempty"`                    //RpRecord
	TypeCovered         string `json:"type_covered,omitempty"`           //RrsigRecord
	OriginalTTL         int    `json:"original_ttl,omitempty"`           //RrsigRecord
	Expiration          string `json:"expiration,omitempty"`             //RrsigRecord
	Inception           string `json:"inception,omitempty"`              //RrsigRecord
	Signer              string `json:"signer,omitempty"`                 //RrsigRecord
	Signature           string `json:"signature,omitempty"`              //RrsigRecord
	Labels              int    `json:"labels,omitempty"`                 //RrsigRecord
	Weight              uint16 `json:"weight,omitempty"`                 //SrvRecord
	Port                uint16 `json:"port,omitempty"`                   //SrvRecord
	FingerprintType     int    `json:"fingerprint_type,omitempty"`       //SshfpRecord
	Fingerprint         string `json:"fingerprint,omitempty"`            //SshfpRecord
	PriorityIncrement   int    `json:"priority_increment,omitempty"`     //MX priority Increment
}

var (
	zoneRecordWriteLock sync.Mutex
)

func (record *RecordBody) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"name":   record.Name,
		"ttl":    record.TTL,
		"active": record.Active,
		"target": record.Target,
	}
}

func NewRecordBody(params RecordBody) *RecordBody {
	recordbody := &RecordBody{Name: params.Name}
	return recordbody
}

// Eval option lock arg passed into writable endpoints. Default is true, e.g. lock
func localLock(lockArg []bool) bool {

	for _, lock := range lockArg {
		// should only be one entry
		return lock
	}

	return true

}

func (record *RecordBody) Save(zone string, recLock ...bool) error {
	// This lock will restrict the concurrency of API calls
	// to 1 save request at a time. This is needed for the Soa.Serial value which
	// is required to be incremented for every subsequent update to a zone
	// so we have to save just one request at a time to ensure this is always
	// incremented properly

	if localLock(recLock) {
		zoneRecordWriteLock.Lock()
		defer zoneRecordWriteLock.Unlock()
	}

	req, err := client.NewJSONRequest(
		Config,
		"POST",
		"/config-dns/v2/zones/"+zone+"/names/"+record.Name+"/types/"+record.RecordType,
		record,
	)
	if err != nil {
		return err
	}

	edge.PrintHttpRequest(req, true)

	res, err := client.Do(Config, req)

	// Network error
	if err != nil {
		return &RecordError{
			fieldName:        record.Name,
			httpErrorMessage: err.Error(),
			err:              err,
		}
	}

	edge.PrintHttpResponse(res, true)

	// API error
	if client.IsError(res) {
		err := client.NewAPIError(res)
		return &RecordError{fieldName: record.Name, apiErrorMessage: err.Detail, err: err}
	}

	return nil
}

func (record *RecordBody) Update(zone string, recLock ...bool) error {
	// This lock will restrict the concurrency of API calls
	// to 1 save request at a time. This is needed for the Soa.Serial value which
	// is required to be incremented for every subsequent update to a zone
	// so we have to save just one request at a time to ensure this is always
	// incremented properly

	if localLock(recLock) {
		zoneRecordWriteLock.Lock()
		defer zoneRecordWriteLock.Unlock()
	}

	req, err := client.NewJSONRequest(
		Config,
		"PUT",
		"/config-dns/v2/zones/"+zone+"/names/"+record.Name+"/types/"+record.RecordType,
		record,
	)
	if err != nil {
		return err
	}

	edge.PrintHttpRequest(req, true)

	res, err := client.Do(Config, req)

	// Network error
	if err != nil {
		return &RecordError{
			fieldName:        record.Name,
			httpErrorMessage: err.Error(),
			err:              err,
		}
	}

	edge.PrintHttpResponse(res, true)

	// API error
	if client.IsError(res) {
		err := client.NewAPIError(res)
		return &RecordError{fieldName: record.Name, apiErrorMessage: err.Detail, err: err}
	}

	return nil
}

func (record *RecordBody) Delete(zone string, recLock ...bool) error {
	// This lock will restrict the concurrency of API calls
	// to 1 save request at a time. This is needed for the Soa.Serial value which
	// is required to be incremented for every subsequent update to a zone
	// so we have to save just one request at a time to ensure this is always
	// incremented properly

	if localLock(recLock) {
		zoneRecordWriteLock.Lock()
		defer zoneRecordWriteLock.Unlock()
	}

	req, err := client.NewJSONRequest(
		Config,
		"DELETE",
		"/config-dns/v2/zones/"+zone+"/names/"+record.Name+"/types/"+record.RecordType,
		nil,
	)
	if err != nil {
		return err
	}

	edge.PrintHttpRequest(req, true)

	res, err := client.Do(Config, req)

	// Network error
	if err != nil {
		return &RecordError{
			fieldName:        record.Name,
			httpErrorMessage: err.Error(),
			err:              err,
		}
	}

	edge.PrintHttpResponse(res, true)

	// API error
	if client.IsError(res) {
		if res.StatusCode != 404 {
			err := client.NewAPIError(res)
			return &RecordError{fieldName: record.Name, apiErrorMessage: err.Detail, err: err}
		}
	}

	return nil
}
