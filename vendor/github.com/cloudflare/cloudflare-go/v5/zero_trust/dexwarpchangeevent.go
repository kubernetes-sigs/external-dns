// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/tidwall/gjson"
)

// DEXWARPChangeEventService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDEXWARPChangeEventService] method instead.
type DEXWARPChangeEventService struct {
	Options []option.RequestOption
}

// NewDEXWARPChangeEventService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDEXWARPChangeEventService(opts ...option.RequestOption) (r *DEXWARPChangeEventService) {
	r = &DEXWARPChangeEventService{}
	r.Options = opts
	return
}

// List WARP configuration and enablement toggle change events by device.
func (r *DEXWARPChangeEventService) Get(ctx context.Context, params DEXWARPChangeEventGetParams, opts ...option.RequestOption) (res *[]DexwarpChangeEventGetResponse, err error) {
	var env DexwarpChangeEventGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dex/warp-change-events", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DexwarpChangeEventGetResponse struct {
	// The account name.
	AccountName string `json:"account_name"`
	// The public account identifier.
	AccountTag string `json:"account_tag"`
	// API Resource UUID tag.
	DeviceID string `json:"device_id"`
	// API Resource UUID tag.
	DeviceRegistration string `json:"device_registration"`
	// This field can have the runtime type of
	// [DexwarpChangeEventGetResponseDigitalExperienceMonitoringWARPConfigChangeEventFrom].
	From interface{} `json:"from"`
	// The hostname of the machine the event is from
	Hostname string `json:"hostname"`
	// The serial number of the machine the event is from
	SerialNumber string `json:"serial_number"`
	// Timestamp in ISO format
	Timestamp string `json:"timestamp"`
	// This field can have the runtime type of
	// [DexwarpChangeEventGetResponseDigitalExperienceMonitoringWARPConfigChangeEventTo].
	To interface{} `json:"to"`
	// The state of the WARP toggle.
	Toggle DexwarpChangeEventGetResponseToggle `json:"toggle"`
	// Email tied to the device
	UserEmail string                            `json:"user_email"`
	JSON      dexwarpChangeEventGetResponseJSON `json:"-"`
	union     DexwarpChangeEventGetResponseUnion
}

// dexwarpChangeEventGetResponseJSON contains the JSON metadata for the struct
// [DexwarpChangeEventGetResponse]
type dexwarpChangeEventGetResponseJSON struct {
	AccountName        apijson.Field
	AccountTag         apijson.Field
	DeviceID           apijson.Field
	DeviceRegistration apijson.Field
	From               apijson.Field
	Hostname           apijson.Field
	SerialNumber       apijson.Field
	Timestamp          apijson.Field
	To                 apijson.Field
	Toggle             apijson.Field
	UserEmail          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r dexwarpChangeEventGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *DexwarpChangeEventGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = DexwarpChangeEventGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DexwarpChangeEventGetResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DexwarpChangeEventGetResponseDigitalExperienceMonitoringWARPToggleChangeEvent],
// [DexwarpChangeEventGetResponseDigitalExperienceMonitoringWARPConfigChangeEvent].
func (r DexwarpChangeEventGetResponse) AsUnion() DexwarpChangeEventGetResponseUnion {
	return r.union
}

// Union satisfied by
// [DexwarpChangeEventGetResponseDigitalExperienceMonitoringWARPToggleChangeEvent]
// or
// [DexwarpChangeEventGetResponseDigitalExperienceMonitoringWARPConfigChangeEvent].
type DexwarpChangeEventGetResponseUnion interface {
	implementsDexwarpChangeEventGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DexwarpChangeEventGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DexwarpChangeEventGetResponseDigitalExperienceMonitoringWARPToggleChangeEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DexwarpChangeEventGetResponseDigitalExperienceMonitoringWARPConfigChangeEvent{}),
		},
	)
}

type DexwarpChangeEventGetResponseDigitalExperienceMonitoringWARPToggleChangeEvent struct {
	// The account name.
	AccountName string `json:"account_name"`
	// The public account identifier.
	AccountTag string `json:"account_tag"`
	// API Resource UUID tag.
	DeviceID string `json:"device_id"`
	// API Resource UUID tag.
	DeviceRegistration string `json:"device_registration"`
	// The hostname of the machine the event is from
	Hostname string `json:"hostname"`
	// The serial number of the machine the event is from
	SerialNumber string `json:"serial_number"`
	// Timestamp in ISO format
	Timestamp string `json:"timestamp"`
	// The state of the WARP toggle.
	Toggle DexwarpChangeEventGetResponseDigitalExperienceMonitoringWARPToggleChangeEventToggle `json:"toggle"`
	// Email tied to the device
	UserEmail string                                                                            `json:"user_email"`
	JSON      dexwarpChangeEventGetResponseDigitalExperienceMonitoringWARPToggleChangeEventJSON `json:"-"`
}

// dexwarpChangeEventGetResponseDigitalExperienceMonitoringWARPToggleChangeEventJSON
// contains the JSON metadata for the struct
// [DexwarpChangeEventGetResponseDigitalExperienceMonitoringWARPToggleChangeEvent]
type dexwarpChangeEventGetResponseDigitalExperienceMonitoringWARPToggleChangeEventJSON struct {
	AccountName        apijson.Field
	AccountTag         apijson.Field
	DeviceID           apijson.Field
	DeviceRegistration apijson.Field
	Hostname           apijson.Field
	SerialNumber       apijson.Field
	Timestamp          apijson.Field
	Toggle             apijson.Field
	UserEmail          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DexwarpChangeEventGetResponseDigitalExperienceMonitoringWARPToggleChangeEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexwarpChangeEventGetResponseDigitalExperienceMonitoringWARPToggleChangeEventJSON) RawJSON() string {
	return r.raw
}

func (r DexwarpChangeEventGetResponseDigitalExperienceMonitoringWARPToggleChangeEvent) implementsDexwarpChangeEventGetResponse() {
}

// The state of the WARP toggle.
type DexwarpChangeEventGetResponseDigitalExperienceMonitoringWARPToggleChangeEventToggle string

const (
	DexwarpChangeEventGetResponseDigitalExperienceMonitoringWARPToggleChangeEventToggleOn  DexwarpChangeEventGetResponseDigitalExperienceMonitoringWARPToggleChangeEventToggle = "on"
	DexwarpChangeEventGetResponseDigitalExperienceMonitoringWARPToggleChangeEventToggleOff DexwarpChangeEventGetResponseDigitalExperienceMonitoringWARPToggleChangeEventToggle = "off"
)

func (r DexwarpChangeEventGetResponseDigitalExperienceMonitoringWARPToggleChangeEventToggle) IsKnown() bool {
	switch r {
	case DexwarpChangeEventGetResponseDigitalExperienceMonitoringWARPToggleChangeEventToggleOn, DexwarpChangeEventGetResponseDigitalExperienceMonitoringWARPToggleChangeEventToggleOff:
		return true
	}
	return false
}

type DexwarpChangeEventGetResponseDigitalExperienceMonitoringWARPConfigChangeEvent struct {
	// API Resource UUID tag.
	DeviceID string `json:"device_id"`
	// API Resource UUID tag.
	DeviceRegistration string                                                                            `json:"device_registration"`
	From               DexwarpChangeEventGetResponseDigitalExperienceMonitoringWARPConfigChangeEventFrom `json:"from"`
	// The hostname of the machine the event is from
	Hostname string `json:"hostname"`
	// The serial number of the machine the event is from
	SerialNumber string `json:"serial_number"`
	// Timestamp in ISO format
	Timestamp string                                                                          `json:"timestamp"`
	To        DexwarpChangeEventGetResponseDigitalExperienceMonitoringWARPConfigChangeEventTo `json:"to"`
	// Email tied to the device
	UserEmail string                                                                            `json:"user_email"`
	JSON      dexwarpChangeEventGetResponseDigitalExperienceMonitoringWARPConfigChangeEventJSON `json:"-"`
}

// dexwarpChangeEventGetResponseDigitalExperienceMonitoringWARPConfigChangeEventJSON
// contains the JSON metadata for the struct
// [DexwarpChangeEventGetResponseDigitalExperienceMonitoringWARPConfigChangeEvent]
type dexwarpChangeEventGetResponseDigitalExperienceMonitoringWARPConfigChangeEventJSON struct {
	DeviceID           apijson.Field
	DeviceRegistration apijson.Field
	From               apijson.Field
	Hostname           apijson.Field
	SerialNumber       apijson.Field
	Timestamp          apijson.Field
	To                 apijson.Field
	UserEmail          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DexwarpChangeEventGetResponseDigitalExperienceMonitoringWARPConfigChangeEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexwarpChangeEventGetResponseDigitalExperienceMonitoringWARPConfigChangeEventJSON) RawJSON() string {
	return r.raw
}

func (r DexwarpChangeEventGetResponseDigitalExperienceMonitoringWARPConfigChangeEvent) implementsDexwarpChangeEventGetResponse() {
}

type DexwarpChangeEventGetResponseDigitalExperienceMonitoringWARPConfigChangeEventFrom struct {
	// The account name.
	AccountName string `json:"account_name"`
	// API Resource UUID tag.
	AccountTag string `json:"account_tag"`
	// The name of the WARP configuration.
	ConfigName string                                                                                `json:"config_name"`
	JSON       dexwarpChangeEventGetResponseDigitalExperienceMonitoringWARPConfigChangeEventFromJSON `json:"-"`
}

// dexwarpChangeEventGetResponseDigitalExperienceMonitoringWARPConfigChangeEventFromJSON
// contains the JSON metadata for the struct
// [DexwarpChangeEventGetResponseDigitalExperienceMonitoringWARPConfigChangeEventFrom]
type dexwarpChangeEventGetResponseDigitalExperienceMonitoringWARPConfigChangeEventFromJSON struct {
	AccountName apijson.Field
	AccountTag  apijson.Field
	ConfigName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexwarpChangeEventGetResponseDigitalExperienceMonitoringWARPConfigChangeEventFrom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexwarpChangeEventGetResponseDigitalExperienceMonitoringWARPConfigChangeEventFromJSON) RawJSON() string {
	return r.raw
}

type DexwarpChangeEventGetResponseDigitalExperienceMonitoringWARPConfigChangeEventTo struct {
	// The account name.
	AccountName string `json:"account_name"`
	// API Resource UUID tag.
	AccountTag string `json:"account_tag"`
	// The name of the WARP configuration.
	ConfigName string                                                                              `json:"config_name"`
	JSON       dexwarpChangeEventGetResponseDigitalExperienceMonitoringWARPConfigChangeEventToJSON `json:"-"`
}

// dexwarpChangeEventGetResponseDigitalExperienceMonitoringWARPConfigChangeEventToJSON
// contains the JSON metadata for the struct
// [DexwarpChangeEventGetResponseDigitalExperienceMonitoringWARPConfigChangeEventTo]
type dexwarpChangeEventGetResponseDigitalExperienceMonitoringWARPConfigChangeEventToJSON struct {
	AccountName apijson.Field
	AccountTag  apijson.Field
	ConfigName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexwarpChangeEventGetResponseDigitalExperienceMonitoringWARPConfigChangeEventTo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexwarpChangeEventGetResponseDigitalExperienceMonitoringWARPConfigChangeEventToJSON) RawJSON() string {
	return r.raw
}

// The state of the WARP toggle.
type DexwarpChangeEventGetResponseToggle string

const (
	DexwarpChangeEventGetResponseToggleOn  DexwarpChangeEventGetResponseToggle = "on"
	DexwarpChangeEventGetResponseToggleOff DexwarpChangeEventGetResponseToggle = "off"
)

func (r DexwarpChangeEventGetResponseToggle) IsKnown() bool {
	switch r {
	case DexwarpChangeEventGetResponseToggleOn, DexwarpChangeEventGetResponseToggleOff:
		return true
	}
	return false
}

type DEXWARPChangeEventGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Start time for the query in ISO (RFC3339 - ISO 8601) format
	From param.Field[string] `query:"from,required"`
	// Page number of paginated results
	Page param.Field[float64] `query:"page,required"`
	// Number of items per page
	PerPage param.Field[float64] `query:"per_page,required"`
	// End time for the query in ISO (RFC3339 - ISO 8601) format
	To param.Field[string] `query:"to,required"`
	// Filter events by account name.
	AccountName param.Field[string] `query:"account_name"`
	// Filter events by WARP configuration name changed from or to. Applicable to
	// type='config' events only.
	ConfigName param.Field[string] `query:"config_name"`
	// Sort response by event timestamp.
	SortOrder param.Field[DexwarpChangeEventGetParamsSortOrder] `query:"sort_order"`
	// Filter events by type toggle value. Applicable to type='toggle' events only.
	Toggle param.Field[DexwarpChangeEventGetParamsToggle] `query:"toggle"`
	// Filter events by type 'config' or 'toggle'
	Type param.Field[DexwarpChangeEventGetParamsType] `query:"type"`
}

// URLQuery serializes [DEXWARPChangeEventGetParams]'s query parameters as
// `url.Values`.
func (r DEXWARPChangeEventGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Sort response by event timestamp.
type DexwarpChangeEventGetParamsSortOrder string

const (
	DexwarpChangeEventGetParamsSortOrderAsc  DexwarpChangeEventGetParamsSortOrder = "ASC"
	DexwarpChangeEventGetParamsSortOrderDesc DexwarpChangeEventGetParamsSortOrder = "DESC"
)

func (r DexwarpChangeEventGetParamsSortOrder) IsKnown() bool {
	switch r {
	case DexwarpChangeEventGetParamsSortOrderAsc, DexwarpChangeEventGetParamsSortOrderDesc:
		return true
	}
	return false
}

// Filter events by type toggle value. Applicable to type='toggle' events only.
type DexwarpChangeEventGetParamsToggle string

const (
	DexwarpChangeEventGetParamsToggleOn  DexwarpChangeEventGetParamsToggle = "on"
	DexwarpChangeEventGetParamsToggleOff DexwarpChangeEventGetParamsToggle = "off"
)

func (r DexwarpChangeEventGetParamsToggle) IsKnown() bool {
	switch r {
	case DexwarpChangeEventGetParamsToggleOn, DexwarpChangeEventGetParamsToggleOff:
		return true
	}
	return false
}

// Filter events by type 'config' or 'toggle'
type DexwarpChangeEventGetParamsType string

const (
	DexwarpChangeEventGetParamsTypeConfig DexwarpChangeEventGetParamsType = "config"
	DexwarpChangeEventGetParamsTypeToggle DexwarpChangeEventGetParamsType = "toggle"
)

func (r DexwarpChangeEventGetParamsType) IsKnown() bool {
	switch r {
	case DexwarpChangeEventGetParamsTypeConfig, DexwarpChangeEventGetParamsTypeToggle:
		return true
	}
	return false
}

type DexwarpChangeEventGetResponseEnvelope struct {
	Errors   []DexwarpChangeEventGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DexwarpChangeEventGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success    DexwarpChangeEventGetResponseEnvelopeSuccess    `json:"success,required"`
	Result     []DexwarpChangeEventGetResponse                 `json:"result"`
	ResultInfo DexwarpChangeEventGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       dexwarpChangeEventGetResponseEnvelopeJSON       `json:"-"`
}

// dexwarpChangeEventGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [DexwarpChangeEventGetResponseEnvelope]
type dexwarpChangeEventGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexwarpChangeEventGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexwarpChangeEventGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DexwarpChangeEventGetResponseEnvelopeErrors struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           DexwarpChangeEventGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             dexwarpChangeEventGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// dexwarpChangeEventGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DexwarpChangeEventGetResponseEnvelopeErrors]
type dexwarpChangeEventGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DexwarpChangeEventGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexwarpChangeEventGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DexwarpChangeEventGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    dexwarpChangeEventGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dexwarpChangeEventGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [DexwarpChangeEventGetResponseEnvelopeErrorsSource]
type dexwarpChangeEventGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexwarpChangeEventGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexwarpChangeEventGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DexwarpChangeEventGetResponseEnvelopeMessages struct {
	Code             int64                                               `json:"code,required"`
	Message          string                                              `json:"message,required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           DexwarpChangeEventGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             dexwarpChangeEventGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// dexwarpChangeEventGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DexwarpChangeEventGetResponseEnvelopeMessages]
type dexwarpChangeEventGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DexwarpChangeEventGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexwarpChangeEventGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DexwarpChangeEventGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    dexwarpChangeEventGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dexwarpChangeEventGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [DexwarpChangeEventGetResponseEnvelopeMessagesSource]
type dexwarpChangeEventGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexwarpChangeEventGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexwarpChangeEventGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DexwarpChangeEventGetResponseEnvelopeSuccess bool

const (
	DexwarpChangeEventGetResponseEnvelopeSuccessTrue DexwarpChangeEventGetResponseEnvelopeSuccess = true
)

func (r DexwarpChangeEventGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DexwarpChangeEventGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DexwarpChangeEventGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                             `json:"total_count"`
	JSON       dexwarpChangeEventGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// dexwarpChangeEventGetResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [DexwarpChangeEventGetResponseEnvelopeResultInfo]
type dexwarpChangeEventGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexwarpChangeEventGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexwarpChangeEventGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
