// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spectrum

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

// AnalyticsAggregateCurrentService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAnalyticsAggregateCurrentService] method instead.
type AnalyticsAggregateCurrentService struct {
	Options []option.RequestOption
}

// NewAnalyticsAggregateCurrentService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAnalyticsAggregateCurrentService(opts ...option.RequestOption) (r *AnalyticsAggregateCurrentService) {
	r = &AnalyticsAggregateCurrentService{}
	r.Options = opts
	return
}

// Retrieves analytics aggregated from the last minute of usage on Spectrum
// applications underneath a given zone.
func (r *AnalyticsAggregateCurrentService) Get(ctx context.Context, params AnalyticsAggregateCurrentGetParams, opts ...option.RequestOption) (res *[]AnalyticsAggregateCurrentGetResponse, err error) {
	var env AnalyticsAggregateCurrentGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/spectrum/analytics/aggregate/current", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AnalyticsAggregateCurrentGetResponse struct {
	// Application identifier.
	AppID string `json:"appID,required"`
	// Number of bytes sent
	BytesEgress float64 `json:"bytesEgress,required"`
	// Number of bytes received
	BytesIngress float64 `json:"bytesIngress,required"`
	// Number of connections
	Connections float64 `json:"connections,required"`
	// Average duration of connections
	DurationAvg float64                                  `json:"durationAvg,required"`
	JSON        analyticsAggregateCurrentGetResponseJSON `json:"-"`
}

// analyticsAggregateCurrentGetResponseJSON contains the JSON metadata for the
// struct [AnalyticsAggregateCurrentGetResponse]
type analyticsAggregateCurrentGetResponseJSON struct {
	AppID        apijson.Field
	BytesEgress  apijson.Field
	BytesIngress apijson.Field
	Connections  apijson.Field
	DurationAvg  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AnalyticsAggregateCurrentGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsAggregateCurrentGetResponseJSON) RawJSON() string {
	return r.raw
}

type AnalyticsAggregateCurrentGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Comma-delimited list of Spectrum Application Id(s). If provided, the response
	// will be limited to Spectrum Application Id(s) that match.
	AppID param.Field[string] `query:"appID"`
	// Co-location identifier.
	ColoName param.Field[string] `query:"colo_name"`
}

// URLQuery serializes [AnalyticsAggregateCurrentGetParams]'s query parameters as
// `url.Values`.
func (r AnalyticsAggregateCurrentGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AnalyticsAggregateCurrentGetResponseEnvelope struct {
	Errors   []AnalyticsAggregateCurrentGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AnalyticsAggregateCurrentGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AnalyticsAggregateCurrentGetResponseEnvelopeSuccess `json:"success,required"`
	Result  []AnalyticsAggregateCurrentGetResponse              `json:"result"`
	JSON    analyticsAggregateCurrentGetResponseEnvelopeJSON    `json:"-"`
}

// analyticsAggregateCurrentGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [AnalyticsAggregateCurrentGetResponseEnvelope]
type analyticsAggregateCurrentGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsAggregateCurrentGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsAggregateCurrentGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AnalyticsAggregateCurrentGetResponseEnvelopeErrors struct {
	Code             int64                                                    `json:"code,required"`
	Message          string                                                   `json:"message,required"`
	DocumentationURL string                                                   `json:"documentation_url"`
	Source           AnalyticsAggregateCurrentGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             analyticsAggregateCurrentGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// analyticsAggregateCurrentGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AnalyticsAggregateCurrentGetResponseEnvelopeErrors]
type analyticsAggregateCurrentGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AnalyticsAggregateCurrentGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsAggregateCurrentGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AnalyticsAggregateCurrentGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                       `json:"pointer"`
	JSON    analyticsAggregateCurrentGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// analyticsAggregateCurrentGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [AnalyticsAggregateCurrentGetResponseEnvelopeErrorsSource]
type analyticsAggregateCurrentGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsAggregateCurrentGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsAggregateCurrentGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AnalyticsAggregateCurrentGetResponseEnvelopeMessages struct {
	Code             int64                                                      `json:"code,required"`
	Message          string                                                     `json:"message,required"`
	DocumentationURL string                                                     `json:"documentation_url"`
	Source           AnalyticsAggregateCurrentGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             analyticsAggregateCurrentGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// analyticsAggregateCurrentGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AnalyticsAggregateCurrentGetResponseEnvelopeMessages]
type analyticsAggregateCurrentGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AnalyticsAggregateCurrentGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsAggregateCurrentGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AnalyticsAggregateCurrentGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                         `json:"pointer"`
	JSON    analyticsAggregateCurrentGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// analyticsAggregateCurrentGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [AnalyticsAggregateCurrentGetResponseEnvelopeMessagesSource]
type analyticsAggregateCurrentGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsAggregateCurrentGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsAggregateCurrentGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AnalyticsAggregateCurrentGetResponseEnvelopeSuccess bool

const (
	AnalyticsAggregateCurrentGetResponseEnvelopeSuccessTrue AnalyticsAggregateCurrentGetResponseEnvelopeSuccess = true
)

func (r AnalyticsAggregateCurrentGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AnalyticsAggregateCurrentGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
