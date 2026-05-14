// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// DEXFleetStatusService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDEXFleetStatusService] method instead.
type DEXFleetStatusService struct {
	Options []option.RequestOption
	Devices *DEXFleetStatusDeviceService
}

// NewDEXFleetStatusService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDEXFleetStatusService(opts ...option.RequestOption) (r *DEXFleetStatusService) {
	r = &DEXFleetStatusService{}
	r.Options = opts
	r.Devices = NewDEXFleetStatusDeviceService(opts...)
	return
}

// List details for live (up to 60 minutes) devices using WARP
func (r *DEXFleetStatusService) Live(ctx context.Context, params DEXFleetStatusLiveParams, opts ...option.RequestOption) (res *DEXFleetStatusLiveResponse, err error) {
	var env DEXFleetStatusLiveResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dex/fleet-status/live", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List details for devices using WARP, up to 7 days
func (r *DEXFleetStatusService) OverTime(ctx context.Context, params DEXFleetStatusOverTimeParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dex/fleet-status/over-time", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, nil, opts...)
	return
}

type LiveStat struct {
	// Number of unique devices
	UniqueDevicesTotal float64      `json:"uniqueDevicesTotal"`
	Value              string       `json:"value"`
	JSON               liveStatJSON `json:"-"`
}

// liveStatJSON contains the JSON metadata for the struct [LiveStat]
type liveStatJSON struct {
	UniqueDevicesTotal apijson.Field
	Value              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *LiveStat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveStatJSON) RawJSON() string {
	return r.raw
}

type DEXFleetStatusLiveResponse struct {
	DeviceStats DEXFleetStatusLiveResponseDeviceStats `json:"deviceStats"`
	JSON        dexFleetStatusLiveResponseJSON        `json:"-"`
}

// dexFleetStatusLiveResponseJSON contains the JSON metadata for the struct
// [DEXFleetStatusLiveResponse]
type dexFleetStatusLiveResponseJSON struct {
	DeviceStats apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXFleetStatusLiveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexFleetStatusLiveResponseJSON) RawJSON() string {
	return r.raw
}

type DEXFleetStatusLiveResponseDeviceStats struct {
	ByColo     []LiveStat `json:"byColo,nullable"`
	ByMode     []LiveStat `json:"byMode,nullable"`
	ByPlatform []LiveStat `json:"byPlatform,nullable"`
	ByStatus   []LiveStat `json:"byStatus,nullable"`
	ByVersion  []LiveStat `json:"byVersion,nullable"`
	// Number of unique devices
	UniqueDevicesTotal float64                                   `json:"uniqueDevicesTotal"`
	JSON               dexFleetStatusLiveResponseDeviceStatsJSON `json:"-"`
}

// dexFleetStatusLiveResponseDeviceStatsJSON contains the JSON metadata for the
// struct [DEXFleetStatusLiveResponseDeviceStats]
type dexFleetStatusLiveResponseDeviceStatsJSON struct {
	ByColo             apijson.Field
	ByMode             apijson.Field
	ByPlatform         apijson.Field
	ByStatus           apijson.Field
	ByVersion          apijson.Field
	UniqueDevicesTotal apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DEXFleetStatusLiveResponseDeviceStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexFleetStatusLiveResponseDeviceStatsJSON) RawJSON() string {
	return r.raw
}

type DEXFleetStatusLiveParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Number of minutes before current time
	SinceMinutes param.Field[float64] `query:"since_minutes,required"`
}

// URLQuery serializes [DEXFleetStatusLiveParams]'s query parameters as
// `url.Values`.
func (r DEXFleetStatusLiveParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DEXFleetStatusLiveResponseEnvelope struct {
	Errors   []DEXFleetStatusLiveResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DEXFleetStatusLiveResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DEXFleetStatusLiveResponseEnvelopeSuccess `json:"success,required"`
	Result  DEXFleetStatusLiveResponse                `json:"result"`
	JSON    dexFleetStatusLiveResponseEnvelopeJSON    `json:"-"`
}

// dexFleetStatusLiveResponseEnvelopeJSON contains the JSON metadata for the struct
// [DEXFleetStatusLiveResponseEnvelope]
type dexFleetStatusLiveResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXFleetStatusLiveResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexFleetStatusLiveResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DEXFleetStatusLiveResponseEnvelopeErrors struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           DEXFleetStatusLiveResponseEnvelopeErrorsSource `json:"source"`
	JSON             dexFleetStatusLiveResponseEnvelopeErrorsJSON   `json:"-"`
}

// dexFleetStatusLiveResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DEXFleetStatusLiveResponseEnvelopeErrors]
type dexFleetStatusLiveResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DEXFleetStatusLiveResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexFleetStatusLiveResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DEXFleetStatusLiveResponseEnvelopeErrorsSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    dexFleetStatusLiveResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dexFleetStatusLiveResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [DEXFleetStatusLiveResponseEnvelopeErrorsSource]
type dexFleetStatusLiveResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXFleetStatusLiveResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexFleetStatusLiveResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DEXFleetStatusLiveResponseEnvelopeMessages struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           DEXFleetStatusLiveResponseEnvelopeMessagesSource `json:"source"`
	JSON             dexFleetStatusLiveResponseEnvelopeMessagesJSON   `json:"-"`
}

// dexFleetStatusLiveResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DEXFleetStatusLiveResponseEnvelopeMessages]
type dexFleetStatusLiveResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DEXFleetStatusLiveResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexFleetStatusLiveResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DEXFleetStatusLiveResponseEnvelopeMessagesSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    dexFleetStatusLiveResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dexFleetStatusLiveResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [DEXFleetStatusLiveResponseEnvelopeMessagesSource]
type dexFleetStatusLiveResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXFleetStatusLiveResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexFleetStatusLiveResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DEXFleetStatusLiveResponseEnvelopeSuccess bool

const (
	DEXFleetStatusLiveResponseEnvelopeSuccessTrue DEXFleetStatusLiveResponseEnvelopeSuccess = true
)

func (r DEXFleetStatusLiveResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DEXFleetStatusLiveResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DEXFleetStatusOverTimeParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Time range beginning in ISO format
	From param.Field[string] `query:"from,required"`
	// Time range end in ISO format
	To param.Field[string] `query:"to,required"`
	// Cloudflare colo
	Colo param.Field[string] `query:"colo"`
	// Device-specific ID, given as UUID v4
	DeviceID param.Field[string] `query:"device_id"`
}

// URLQuery serializes [DEXFleetStatusOverTimeParams]'s query parameters as
// `url.Values`.
func (r DEXFleetStatusOverTimeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
