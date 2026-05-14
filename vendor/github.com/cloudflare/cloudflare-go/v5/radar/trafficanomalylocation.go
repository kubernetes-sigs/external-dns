// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// TrafficAnomalyLocationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTrafficAnomalyLocationService] method instead.
type TrafficAnomalyLocationService struct {
	Options []option.RequestOption
}

// NewTrafficAnomalyLocationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTrafficAnomalyLocationService(opts ...option.RequestOption) (r *TrafficAnomalyLocationService) {
	r = &TrafficAnomalyLocationService{}
	r.Options = opts
	return
}

// Retrieves the sum of Internet traffic anomalies, grouped by location. These
// anomalies are signals that might indicate an outage, automatically detected by
// Radar and manually verified by our team.
func (r *TrafficAnomalyLocationService) Get(ctx context.Context, query TrafficAnomalyLocationGetParams, opts ...option.RequestOption) (res *TrafficAnomalyLocationGetResponse, err error) {
	var env TrafficAnomalyLocationGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/traffic_anomalies/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type TrafficAnomalyLocationGetResponse struct {
	TrafficAnomalies []TrafficAnomalyLocationGetResponseTrafficAnomaly `json:"trafficAnomalies,required"`
	JSON             trafficAnomalyLocationGetResponseJSON             `json:"-"`
}

// trafficAnomalyLocationGetResponseJSON contains the JSON metadata for the struct
// [TrafficAnomalyLocationGetResponse]
type trafficAnomalyLocationGetResponseJSON struct {
	TrafficAnomalies apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TrafficAnomalyLocationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r trafficAnomalyLocationGetResponseJSON) RawJSON() string {
	return r.raw
}

type TrafficAnomalyLocationGetResponseTrafficAnomaly struct {
	ClientCountryAlpha2 string `json:"clientCountryAlpha2,required"`
	ClientCountryName   string `json:"clientCountryName,required"`
	// A numeric string.
	Value string                                              `json:"value,required"`
	JSON  trafficAnomalyLocationGetResponseTrafficAnomalyJSON `json:"-"`
}

// trafficAnomalyLocationGetResponseTrafficAnomalyJSON contains the JSON metadata
// for the struct [TrafficAnomalyLocationGetResponseTrafficAnomaly]
type trafficAnomalyLocationGetResponseTrafficAnomalyJSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *TrafficAnomalyLocationGetResponseTrafficAnomaly) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r trafficAnomalyLocationGetResponseTrafficAnomalyJSON) RawJSON() string {
	return r.raw
}

type TrafficAnomalyLocationGetParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range.
	DateRange param.Field[string] `query:"dateRange"`
	// Start of the date range (inclusive).
	DateStart param.Field[time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[TrafficAnomalyLocationGetParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit  param.Field[int64]                                 `query:"limit"`
	Status param.Field[TrafficAnomalyLocationGetParamsStatus] `query:"status"`
}

// URLQuery serializes [TrafficAnomalyLocationGetParams]'s query parameters as
// `url.Values`.
func (r TrafficAnomalyLocationGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type TrafficAnomalyLocationGetParamsFormat string

const (
	TrafficAnomalyLocationGetParamsFormatJson TrafficAnomalyLocationGetParamsFormat = "JSON"
	TrafficAnomalyLocationGetParamsFormatCsv  TrafficAnomalyLocationGetParamsFormat = "CSV"
)

func (r TrafficAnomalyLocationGetParamsFormat) IsKnown() bool {
	switch r {
	case TrafficAnomalyLocationGetParamsFormatJson, TrafficAnomalyLocationGetParamsFormatCsv:
		return true
	}
	return false
}

type TrafficAnomalyLocationGetParamsStatus string

const (
	TrafficAnomalyLocationGetParamsStatusVerified   TrafficAnomalyLocationGetParamsStatus = "VERIFIED"
	TrafficAnomalyLocationGetParamsStatusUnverified TrafficAnomalyLocationGetParamsStatus = "UNVERIFIED"
)

func (r TrafficAnomalyLocationGetParamsStatus) IsKnown() bool {
	switch r {
	case TrafficAnomalyLocationGetParamsStatusVerified, TrafficAnomalyLocationGetParamsStatusUnverified:
		return true
	}
	return false
}

type TrafficAnomalyLocationGetResponseEnvelope struct {
	Result  TrafficAnomalyLocationGetResponse             `json:"result,required"`
	Success bool                                          `json:"success,required"`
	JSON    trafficAnomalyLocationGetResponseEnvelopeJSON `json:"-"`
}

// trafficAnomalyLocationGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [TrafficAnomalyLocationGetResponseEnvelope]
type trafficAnomalyLocationGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TrafficAnomalyLocationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r trafficAnomalyLocationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
