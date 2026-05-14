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

// AnnotationService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAnnotationService] method instead.
type AnnotationService struct {
	Options []option.RequestOption
	Outages *AnnotationOutageService
}

// NewAnnotationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAnnotationService(opts ...option.RequestOption) (r *AnnotationService) {
	r = &AnnotationService{}
	r.Options = opts
	r.Outages = NewAnnotationOutageService(opts...)
	return
}

// Retrieves the latest annotations.
func (r *AnnotationService) List(ctx context.Context, query AnnotationListParams, opts ...option.RequestOption) (res *AnnotationListResponse, err error) {
	var env AnnotationListResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/annotations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AnnotationListResponse struct {
	Annotations []AnnotationListResponseAnnotation `json:"annotations,required"`
	JSON        annotationListResponseJSON         `json:"-"`
}

// annotationListResponseJSON contains the JSON metadata for the struct
// [AnnotationListResponse]
type annotationListResponseJSON struct {
	Annotations apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnnotationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r annotationListResponseJSON) RawJSON() string {
	return r.raw
}

type AnnotationListResponseAnnotation struct {
	ID               string                                             `json:"id,required"`
	ASNs             []int64                                            `json:"asns,required"`
	ASNsDetails      []AnnotationListResponseAnnotationsASNsDetail      `json:"asnsDetails,required"`
	DataSource       string                                             `json:"dataSource,required"`
	EventType        string                                             `json:"eventType,required"`
	Locations        []string                                           `json:"locations,required"`
	LocationsDetails []AnnotationListResponseAnnotationsLocationsDetail `json:"locationsDetails,required"`
	Outage           AnnotationListResponseAnnotationsOutage            `json:"outage,required"`
	StartDate        string                                             `json:"startDate,required"`
	Description      string                                             `json:"description"`
	EndDate          string                                             `json:"endDate"`
	LinkedURL        string                                             `json:"linkedUrl"`
	Scope            string                                             `json:"scope"`
	JSON             annotationListResponseAnnotationJSON               `json:"-"`
}

// annotationListResponseAnnotationJSON contains the JSON metadata for the struct
// [AnnotationListResponseAnnotation]
type annotationListResponseAnnotationJSON struct {
	ID               apijson.Field
	ASNs             apijson.Field
	ASNsDetails      apijson.Field
	DataSource       apijson.Field
	EventType        apijson.Field
	Locations        apijson.Field
	LocationsDetails apijson.Field
	Outage           apijson.Field
	StartDate        apijson.Field
	Description      apijson.Field
	EndDate          apijson.Field
	LinkedURL        apijson.Field
	Scope            apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AnnotationListResponseAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r annotationListResponseAnnotationJSON) RawJSON() string {
	return r.raw
}

type AnnotationListResponseAnnotationsASNsDetail struct {
	ASN       string                                                `json:"asn,required"`
	Name      string                                                `json:"name,required"`
	Locations AnnotationListResponseAnnotationsASNsDetailsLocations `json:"locations"`
	JSON      annotationListResponseAnnotationsASNsDetailJSON       `json:"-"`
}

// annotationListResponseAnnotationsASNsDetailJSON contains the JSON metadata for
// the struct [AnnotationListResponseAnnotationsASNsDetail]
type annotationListResponseAnnotationsASNsDetailJSON struct {
	ASN         apijson.Field
	Name        apijson.Field
	Locations   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnnotationListResponseAnnotationsASNsDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r annotationListResponseAnnotationsASNsDetailJSON) RawJSON() string {
	return r.raw
}

type AnnotationListResponseAnnotationsASNsDetailsLocations struct {
	Code string                                                    `json:"code,required"`
	Name string                                                    `json:"name,required"`
	JSON annotationListResponseAnnotationsASNsDetailsLocationsJSON `json:"-"`
}

// annotationListResponseAnnotationsASNsDetailsLocationsJSON contains the JSON
// metadata for the struct [AnnotationListResponseAnnotationsASNsDetailsLocations]
type annotationListResponseAnnotationsASNsDetailsLocationsJSON struct {
	Code        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnnotationListResponseAnnotationsASNsDetailsLocations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r annotationListResponseAnnotationsASNsDetailsLocationsJSON) RawJSON() string {
	return r.raw
}

type AnnotationListResponseAnnotationsLocationsDetail struct {
	Code string                                               `json:"code,required"`
	Name string                                               `json:"name,required"`
	JSON annotationListResponseAnnotationsLocationsDetailJSON `json:"-"`
}

// annotationListResponseAnnotationsLocationsDetailJSON contains the JSON metadata
// for the struct [AnnotationListResponseAnnotationsLocationsDetail]
type annotationListResponseAnnotationsLocationsDetailJSON struct {
	Code        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnnotationListResponseAnnotationsLocationsDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r annotationListResponseAnnotationsLocationsDetailJSON) RawJSON() string {
	return r.raw
}

type AnnotationListResponseAnnotationsOutage struct {
	OutageCause string                                      `json:"outageCause,required"`
	OutageType  string                                      `json:"outageType,required"`
	JSON        annotationListResponseAnnotationsOutageJSON `json:"-"`
}

// annotationListResponseAnnotationsOutageJSON contains the JSON metadata for the
// struct [AnnotationListResponseAnnotationsOutage]
type annotationListResponseAnnotationsOutageJSON struct {
	OutageCause apijson.Field
	OutageType  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnnotationListResponseAnnotationsOutage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r annotationListResponseAnnotationsOutageJSON) RawJSON() string {
	return r.raw
}

type AnnotationListParams struct {
	// Filters results by Autonomous System. Specify a single Autonomous System Number
	// (ASN) as integer.
	ASN param.Field[int64] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range.
	DateRange param.Field[string] `query:"dateRange"`
	// Start of the date range (inclusive).
	DateStart param.Field[time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[AnnotationListParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify an alpha-2 location code.
	Location param.Field[string] `query:"location"`
	// Skips the specified number of objects before fetching the results.
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [AnnotationListParams]'s query parameters as `url.Values`.
func (r AnnotationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type AnnotationListParamsFormat string

const (
	AnnotationListParamsFormatJson AnnotationListParamsFormat = "JSON"
	AnnotationListParamsFormatCsv  AnnotationListParamsFormat = "CSV"
)

func (r AnnotationListParamsFormat) IsKnown() bool {
	switch r {
	case AnnotationListParamsFormatJson, AnnotationListParamsFormatCsv:
		return true
	}
	return false
}

type AnnotationListResponseEnvelope struct {
	Result  AnnotationListResponse             `json:"result,required"`
	Success bool                               `json:"success,required"`
	JSON    annotationListResponseEnvelopeJSON `json:"-"`
}

// annotationListResponseEnvelopeJSON contains the JSON metadata for the struct
// [AnnotationListResponseEnvelope]
type annotationListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnnotationListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r annotationListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
