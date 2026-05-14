// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// CtLogService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCtLogService] method instead.
type CtLogService struct {
	Options []option.RequestOption
}

// NewCtLogService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCtLogService(opts ...option.RequestOption) (r *CtLogService) {
	r = &CtLogService{}
	r.Options = opts
	return
}

// Retrieves a list of certificate logs.
func (r *CtLogService) List(ctx context.Context, query CtLogListParams, opts ...option.RequestOption) (res *CtLogListResponse, err error) {
	var env CtLogListResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/ct/logs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the requested certificate log information.
func (r *CtLogService) Get(ctx context.Context, logSlug string, query CtLogGetParams, opts ...option.RequestOption) (res *CtLogGetResponse, err error) {
	var env CtLogGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if logSlug == "" {
		err = errors.New("missing required log_slug parameter")
		return
	}
	path := fmt.Sprintf("radar/ct/logs/%s", logSlug)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type CtLogListResponse struct {
	CertificateLogs []CtLogListResponseCertificateLog `json:"certificateLogs,required"`
	JSON            ctLogListResponseJSON             `json:"-"`
}

// ctLogListResponseJSON contains the JSON metadata for the struct
// [CtLogListResponse]
type ctLogListResponseJSON struct {
	CertificateLogs apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *CtLogListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctLogListResponseJSON) RawJSON() string {
	return r.raw
}

type CtLogListResponseCertificateLog struct {
	// The API standard that the certificate log follows.
	API CtLogListResponseCertificateLogsAPI `json:"api,required"`
	// A brief description of the certificate log.
	Description string `json:"description,required"`
	// The end date and time for when the log will stop accepting certificates.
	EndExclusive time.Time `json:"endExclusive,required" format:"date-time"`
	// The organization responsible for operating the certificate log.
	Operator string `json:"operator,required"`
	// A URL-friendly, kebab-case identifier for the certificate log.
	Slug string `json:"slug,required"`
	// The start date and time for when the log starts accepting certificates.
	StartInclusive time.Time `json:"startInclusive,required" format:"date-time"`
	// The current state of the certificate log. More details about log states can be
	// found here:
	// https://googlechrome.github.io/CertificateTransparency/log_states.html
	State CtLogListResponseCertificateLogsState `json:"state,required"`
	// Timestamp of when the log state was last updated.
	StateTimestamp time.Time `json:"stateTimestamp,required" format:"date-time"`
	// The URL for the certificate log.
	URL  string                              `json:"url,required"`
	JSON ctLogListResponseCertificateLogJSON `json:"-"`
}

// ctLogListResponseCertificateLogJSON contains the JSON metadata for the struct
// [CtLogListResponseCertificateLog]
type ctLogListResponseCertificateLogJSON struct {
	API            apijson.Field
	Description    apijson.Field
	EndExclusive   apijson.Field
	Operator       apijson.Field
	Slug           apijson.Field
	StartInclusive apijson.Field
	State          apijson.Field
	StateTimestamp apijson.Field
	URL            apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CtLogListResponseCertificateLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctLogListResponseCertificateLogJSON) RawJSON() string {
	return r.raw
}

// The API standard that the certificate log follows.
type CtLogListResponseCertificateLogsAPI string

const (
	CtLogListResponseCertificateLogsAPIRfc6962 CtLogListResponseCertificateLogsAPI = "RFC6962"
	CtLogListResponseCertificateLogsAPIStatic  CtLogListResponseCertificateLogsAPI = "STATIC"
)

func (r CtLogListResponseCertificateLogsAPI) IsKnown() bool {
	switch r {
	case CtLogListResponseCertificateLogsAPIRfc6962, CtLogListResponseCertificateLogsAPIStatic:
		return true
	}
	return false
}

// The current state of the certificate log. More details about log states can be
// found here:
// https://googlechrome.github.io/CertificateTransparency/log_states.html
type CtLogListResponseCertificateLogsState string

const (
	CtLogListResponseCertificateLogsStateUsable    CtLogListResponseCertificateLogsState = "USABLE"
	CtLogListResponseCertificateLogsStatePending   CtLogListResponseCertificateLogsState = "PENDING"
	CtLogListResponseCertificateLogsStateQualified CtLogListResponseCertificateLogsState = "QUALIFIED"
	CtLogListResponseCertificateLogsStateReadOnly  CtLogListResponseCertificateLogsState = "READ_ONLY"
	CtLogListResponseCertificateLogsStateRetired   CtLogListResponseCertificateLogsState = "RETIRED"
	CtLogListResponseCertificateLogsStateRejected  CtLogListResponseCertificateLogsState = "REJECTED"
)

func (r CtLogListResponseCertificateLogsState) IsKnown() bool {
	switch r {
	case CtLogListResponseCertificateLogsStateUsable, CtLogListResponseCertificateLogsStatePending, CtLogListResponseCertificateLogsStateQualified, CtLogListResponseCertificateLogsStateReadOnly, CtLogListResponseCertificateLogsStateRetired, CtLogListResponseCertificateLogsStateRejected:
		return true
	}
	return false
}

type CtLogGetResponse struct {
	CertificateLog CtLogGetResponseCertificateLog `json:"certificateLog,required"`
	JSON           ctLogGetResponseJSON           `json:"-"`
}

// ctLogGetResponseJSON contains the JSON metadata for the struct
// [CtLogGetResponse]
type ctLogGetResponseJSON struct {
	CertificateLog apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CtLogGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctLogGetResponseJSON) RawJSON() string {
	return r.raw
}

type CtLogGetResponseCertificateLog struct {
	// The API standard that the certificate log follows.
	API CtLogGetResponseCertificateLogAPI `json:"api,required"`
	// A brief description of the certificate log.
	Description string `json:"description,required"`
	// The end date and time for when the log will stop accepting certificates.
	EndExclusive time.Time `json:"endExclusive,required" format:"date-time"`
	// The organization responsible for operating the certificate log.
	Operator string `json:"operator,required"`
	// Log performance metrics, including averages and per-endpoint details.
	Performance CtLogGetResponseCertificateLogPerformance `json:"performance,required,nullable"`
	// Logs from the same operator.
	Related []CtLogGetResponseCertificateLogRelated `json:"related,required"`
	// A URL-friendly, kebab-case identifier for the certificate log.
	Slug string `json:"slug,required"`
	// The start date and time for when the log starts accepting certificates.
	StartInclusive time.Time `json:"startInclusive,required" format:"date-time"`
	// The current state of the certificate log. More details about log states can be
	// found here:
	// https://googlechrome.github.io/CertificateTransparency/log_states.html
	State CtLogGetResponseCertificateLogState `json:"state,required"`
	// Timestamp of when the log state was last updated.
	StateTimestamp time.Time `json:"stateTimestamp,required" format:"date-time"`
	// The URL for the certificate log.
	URL  string                             `json:"url,required"`
	JSON ctLogGetResponseCertificateLogJSON `json:"-"`
}

// ctLogGetResponseCertificateLogJSON contains the JSON metadata for the struct
// [CtLogGetResponseCertificateLog]
type ctLogGetResponseCertificateLogJSON struct {
	API            apijson.Field
	Description    apijson.Field
	EndExclusive   apijson.Field
	Operator       apijson.Field
	Performance    apijson.Field
	Related        apijson.Field
	Slug           apijson.Field
	StartInclusive apijson.Field
	State          apijson.Field
	StateTimestamp apijson.Field
	URL            apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CtLogGetResponseCertificateLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctLogGetResponseCertificateLogJSON) RawJSON() string {
	return r.raw
}

// The API standard that the certificate log follows.
type CtLogGetResponseCertificateLogAPI string

const (
	CtLogGetResponseCertificateLogAPIRfc6962 CtLogGetResponseCertificateLogAPI = "RFC6962"
	CtLogGetResponseCertificateLogAPIStatic  CtLogGetResponseCertificateLogAPI = "STATIC"
)

func (r CtLogGetResponseCertificateLogAPI) IsKnown() bool {
	switch r {
	case CtLogGetResponseCertificateLogAPIRfc6962, CtLogGetResponseCertificateLogAPIStatic:
		return true
	}
	return false
}

// Log performance metrics, including averages and per-endpoint details.
type CtLogGetResponseCertificateLogPerformance struct {
	Endpoints    []CtLogGetResponseCertificateLogPerformanceEndpoint `json:"endpoints,required"`
	ResponseTime float64                                             `json:"responseTime,required"`
	Uptime       float64                                             `json:"uptime,required"`
	JSON         ctLogGetResponseCertificateLogPerformanceJSON       `json:"-"`
}

// ctLogGetResponseCertificateLogPerformanceJSON contains the JSON metadata for the
// struct [CtLogGetResponseCertificateLogPerformance]
type ctLogGetResponseCertificateLogPerformanceJSON struct {
	Endpoints    apijson.Field
	ResponseTime apijson.Field
	Uptime       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CtLogGetResponseCertificateLogPerformance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctLogGetResponseCertificateLogPerformanceJSON) RawJSON() string {
	return r.raw
}

type CtLogGetResponseCertificateLogPerformanceEndpoint struct {
	// The certificate log endpoint names used in performance metrics.
	Endpoint     CtLogGetResponseCertificateLogPerformanceEndpointsEndpoint `json:"endpoint,required"`
	ResponseTime float64                                                    `json:"responseTime,required"`
	Uptime       float64                                                    `json:"uptime,required"`
	JSON         ctLogGetResponseCertificateLogPerformanceEndpointJSON      `json:"-"`
}

// ctLogGetResponseCertificateLogPerformanceEndpointJSON contains the JSON metadata
// for the struct [CtLogGetResponseCertificateLogPerformanceEndpoint]
type ctLogGetResponseCertificateLogPerformanceEndpointJSON struct {
	Endpoint     apijson.Field
	ResponseTime apijson.Field
	Uptime       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CtLogGetResponseCertificateLogPerformanceEndpoint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctLogGetResponseCertificateLogPerformanceEndpointJSON) RawJSON() string {
	return r.raw
}

// The certificate log endpoint names used in performance metrics.
type CtLogGetResponseCertificateLogPerformanceEndpointsEndpoint string

const (
	CtLogGetResponseCertificateLogPerformanceEndpointsEndpointAddChainNew    CtLogGetResponseCertificateLogPerformanceEndpointsEndpoint = "add-chain (new)"
	CtLogGetResponseCertificateLogPerformanceEndpointsEndpointAddChainOld    CtLogGetResponseCertificateLogPerformanceEndpointsEndpoint = "add-chain (old)"
	CtLogGetResponseCertificateLogPerformanceEndpointsEndpointAddPreChainNew CtLogGetResponseCertificateLogPerformanceEndpointsEndpoint = "add-pre-chain (new)"
	CtLogGetResponseCertificateLogPerformanceEndpointsEndpointAddPreChainOld CtLogGetResponseCertificateLogPerformanceEndpointsEndpoint = "add-pre-chain (old)"
	CtLogGetResponseCertificateLogPerformanceEndpointsEndpointGetEntries     CtLogGetResponseCertificateLogPerformanceEndpointsEndpoint = "get-entries"
	CtLogGetResponseCertificateLogPerformanceEndpointsEndpointGetRoots       CtLogGetResponseCertificateLogPerformanceEndpointsEndpoint = "get-roots"
	CtLogGetResponseCertificateLogPerformanceEndpointsEndpointGetSth         CtLogGetResponseCertificateLogPerformanceEndpointsEndpoint = "get-sth"
)

func (r CtLogGetResponseCertificateLogPerformanceEndpointsEndpoint) IsKnown() bool {
	switch r {
	case CtLogGetResponseCertificateLogPerformanceEndpointsEndpointAddChainNew, CtLogGetResponseCertificateLogPerformanceEndpointsEndpointAddChainOld, CtLogGetResponseCertificateLogPerformanceEndpointsEndpointAddPreChainNew, CtLogGetResponseCertificateLogPerformanceEndpointsEndpointAddPreChainOld, CtLogGetResponseCertificateLogPerformanceEndpointsEndpointGetEntries, CtLogGetResponseCertificateLogPerformanceEndpointsEndpointGetRoots, CtLogGetResponseCertificateLogPerformanceEndpointsEndpointGetSth:
		return true
	}
	return false
}

type CtLogGetResponseCertificateLogRelated struct {
	// A brief description of the certificate log.
	Description string `json:"description,required"`
	// The end date and time for when the log will stop accepting certificates.
	EndExclusive time.Time `json:"endExclusive,required" format:"date-time"`
	// A URL-friendly, kebab-case identifier for the certificate log.
	Slug string `json:"slug,required"`
	// The start date and time for when the log starts accepting certificates.
	StartInclusive time.Time `json:"startInclusive,required" format:"date-time"`
	// The current state of the certificate log. More details about log states can be
	// found here:
	// https://googlechrome.github.io/CertificateTransparency/log_states.html
	State CtLogGetResponseCertificateLogRelatedState `json:"state,required"`
	JSON  ctLogGetResponseCertificateLogRelatedJSON  `json:"-"`
}

// ctLogGetResponseCertificateLogRelatedJSON contains the JSON metadata for the
// struct [CtLogGetResponseCertificateLogRelated]
type ctLogGetResponseCertificateLogRelatedJSON struct {
	Description    apijson.Field
	EndExclusive   apijson.Field
	Slug           apijson.Field
	StartInclusive apijson.Field
	State          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CtLogGetResponseCertificateLogRelated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctLogGetResponseCertificateLogRelatedJSON) RawJSON() string {
	return r.raw
}

// The current state of the certificate log. More details about log states can be
// found here:
// https://googlechrome.github.io/CertificateTransparency/log_states.html
type CtLogGetResponseCertificateLogRelatedState string

const (
	CtLogGetResponseCertificateLogRelatedStateUsable    CtLogGetResponseCertificateLogRelatedState = "USABLE"
	CtLogGetResponseCertificateLogRelatedStatePending   CtLogGetResponseCertificateLogRelatedState = "PENDING"
	CtLogGetResponseCertificateLogRelatedStateQualified CtLogGetResponseCertificateLogRelatedState = "QUALIFIED"
	CtLogGetResponseCertificateLogRelatedStateReadOnly  CtLogGetResponseCertificateLogRelatedState = "READ_ONLY"
	CtLogGetResponseCertificateLogRelatedStateRetired   CtLogGetResponseCertificateLogRelatedState = "RETIRED"
	CtLogGetResponseCertificateLogRelatedStateRejected  CtLogGetResponseCertificateLogRelatedState = "REJECTED"
)

func (r CtLogGetResponseCertificateLogRelatedState) IsKnown() bool {
	switch r {
	case CtLogGetResponseCertificateLogRelatedStateUsable, CtLogGetResponseCertificateLogRelatedStatePending, CtLogGetResponseCertificateLogRelatedStateQualified, CtLogGetResponseCertificateLogRelatedStateReadOnly, CtLogGetResponseCertificateLogRelatedStateRetired, CtLogGetResponseCertificateLogRelatedStateRejected:
		return true
	}
	return false
}

// The current state of the certificate log. More details about log states can be
// found here:
// https://googlechrome.github.io/CertificateTransparency/log_states.html
type CtLogGetResponseCertificateLogState string

const (
	CtLogGetResponseCertificateLogStateUsable    CtLogGetResponseCertificateLogState = "USABLE"
	CtLogGetResponseCertificateLogStatePending   CtLogGetResponseCertificateLogState = "PENDING"
	CtLogGetResponseCertificateLogStateQualified CtLogGetResponseCertificateLogState = "QUALIFIED"
	CtLogGetResponseCertificateLogStateReadOnly  CtLogGetResponseCertificateLogState = "READ_ONLY"
	CtLogGetResponseCertificateLogStateRetired   CtLogGetResponseCertificateLogState = "RETIRED"
	CtLogGetResponseCertificateLogStateRejected  CtLogGetResponseCertificateLogState = "REJECTED"
)

func (r CtLogGetResponseCertificateLogState) IsKnown() bool {
	switch r {
	case CtLogGetResponseCertificateLogStateUsable, CtLogGetResponseCertificateLogStatePending, CtLogGetResponseCertificateLogStateQualified, CtLogGetResponseCertificateLogStateReadOnly, CtLogGetResponseCertificateLogStateRetired, CtLogGetResponseCertificateLogStateRejected:
		return true
	}
	return false
}

type CtLogListParams struct {
	// Format in which results will be returned.
	Format param.Field[CtLogListParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Skips the specified number of objects before fetching the results.
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [CtLogListParams]'s query parameters as `url.Values`.
func (r CtLogListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type CtLogListParamsFormat string

const (
	CtLogListParamsFormatJson CtLogListParamsFormat = "JSON"
	CtLogListParamsFormatCsv  CtLogListParamsFormat = "CSV"
)

func (r CtLogListParamsFormat) IsKnown() bool {
	switch r {
	case CtLogListParamsFormatJson, CtLogListParamsFormatCsv:
		return true
	}
	return false
}

type CtLogListResponseEnvelope struct {
	Result  CtLogListResponse             `json:"result,required"`
	Success bool                          `json:"success,required"`
	JSON    ctLogListResponseEnvelopeJSON `json:"-"`
}

// ctLogListResponseEnvelopeJSON contains the JSON metadata for the struct
// [CtLogListResponseEnvelope]
type ctLogListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CtLogListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctLogListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CtLogGetParams struct {
	// Format in which results will be returned.
	Format param.Field[CtLogGetParamsFormat] `query:"format"`
}

// URLQuery serializes [CtLogGetParams]'s query parameters as `url.Values`.
func (r CtLogGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type CtLogGetParamsFormat string

const (
	CtLogGetParamsFormatJson CtLogGetParamsFormat = "JSON"
	CtLogGetParamsFormatCsv  CtLogGetParamsFormat = "CSV"
)

func (r CtLogGetParamsFormat) IsKnown() bool {
	switch r {
	case CtLogGetParamsFormatJson, CtLogGetParamsFormatCsv:
		return true
	}
	return false
}

type CtLogGetResponseEnvelope struct {
	Result  CtLogGetResponse             `json:"result,required"`
	Success bool                         `json:"success,required"`
	JSON    ctLogGetResponseEnvelopeJSON `json:"-"`
}

// ctLogGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [CtLogGetResponseEnvelope]
type ctLogGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CtLogGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctLogGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
