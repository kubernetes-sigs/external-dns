// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package speed

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
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// PageService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPageService] method instead.
type PageService struct {
	Options []option.RequestOption
	Tests   *PageTestService
}

// NewPageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPageService(opts ...option.RequestOption) (r *PageService) {
	r = &PageService{}
	r.Options = opts
	r.Tests = NewPageTestService(opts...)
	return
}

// Lists all webpages which have been tested.
func (r *PageService) List(ctx context.Context, query PageListParams, opts ...option.RequestOption) (res *pagination.SinglePage[PageListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/speed_api/pages", query.ZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Lists all webpages which have been tested.
func (r *PageService) ListAutoPaging(ctx context.Context, query PageListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[PageListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Lists the core web vital metrics trend over time for a specific page.
func (r *PageService) Trend(ctx context.Context, url string, params PageTrendParams, opts ...option.RequestOption) (res *Trend, err error) {
	var env PageTrendResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if url == "" {
		err = errors.New("missing required url parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/speed_api/pages/%s/trend", params.ZoneID, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type PageListResponse struct {
	// A test region with a label.
	Region LabeledRegion `json:"region"`
	// The frequency of the test.
	ScheduleFrequency PageListResponseScheduleFrequency `json:"scheduleFrequency"`
	Tests             []Test                            `json:"tests"`
	// A URL.
	URL  string               `json:"url"`
	JSON pageListResponseJSON `json:"-"`
}

// pageListResponseJSON contains the JSON metadata for the struct
// [PageListResponse]
type pageListResponseJSON struct {
	Region            apijson.Field
	ScheduleFrequency apijson.Field
	Tests             apijson.Field
	URL               apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PageListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageListResponseJSON) RawJSON() string {
	return r.raw
}

// The frequency of the test.
type PageListResponseScheduleFrequency string

const (
	PageListResponseScheduleFrequencyDaily  PageListResponseScheduleFrequency = "DAILY"
	PageListResponseScheduleFrequencyWeekly PageListResponseScheduleFrequency = "WEEKLY"
)

func (r PageListResponseScheduleFrequency) IsKnown() bool {
	switch r {
	case PageListResponseScheduleFrequencyDaily, PageListResponseScheduleFrequencyWeekly:
		return true
	}
	return false
}

type PageListParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type PageTrendParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The type of device.
	DeviceType param.Field[PageTrendParamsDeviceType] `query:"deviceType,required"`
	// A comma-separated list of metrics to include in the results.
	Metrics param.Field[string] `query:"metrics,required"`
	// A test region.
	Region param.Field[PageTrendParamsRegion] `query:"region,required"`
	Start  param.Field[time.Time]             `query:"start,required" format:"date-time"`
	// The timezone of the start and end timestamps.
	Tz  param.Field[string]    `query:"tz,required"`
	End param.Field[time.Time] `query:"end" format:"date-time"`
}

// URLQuery serializes [PageTrendParams]'s query parameters as `url.Values`.
func (r PageTrendParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The type of device.
type PageTrendParamsDeviceType string

const (
	PageTrendParamsDeviceTypeDesktop PageTrendParamsDeviceType = "DESKTOP"
	PageTrendParamsDeviceTypeMobile  PageTrendParamsDeviceType = "MOBILE"
)

func (r PageTrendParamsDeviceType) IsKnown() bool {
	switch r {
	case PageTrendParamsDeviceTypeDesktop, PageTrendParamsDeviceTypeMobile:
		return true
	}
	return false
}

// A test region.
type PageTrendParamsRegion string

const (
	PageTrendParamsRegionAsiaEast1           PageTrendParamsRegion = "asia-east1"
	PageTrendParamsRegionAsiaNortheast1      PageTrendParamsRegion = "asia-northeast1"
	PageTrendParamsRegionAsiaNortheast2      PageTrendParamsRegion = "asia-northeast2"
	PageTrendParamsRegionAsiaSouth1          PageTrendParamsRegion = "asia-south1"
	PageTrendParamsRegionAsiaSoutheast1      PageTrendParamsRegion = "asia-southeast1"
	PageTrendParamsRegionAustraliaSoutheast1 PageTrendParamsRegion = "australia-southeast1"
	PageTrendParamsRegionEuropeNorth1        PageTrendParamsRegion = "europe-north1"
	PageTrendParamsRegionEuropeSouthwest1    PageTrendParamsRegion = "europe-southwest1"
	PageTrendParamsRegionEuropeWest1         PageTrendParamsRegion = "europe-west1"
	PageTrendParamsRegionEuropeWest2         PageTrendParamsRegion = "europe-west2"
	PageTrendParamsRegionEuropeWest3         PageTrendParamsRegion = "europe-west3"
	PageTrendParamsRegionEuropeWest4         PageTrendParamsRegion = "europe-west4"
	PageTrendParamsRegionEuropeWest8         PageTrendParamsRegion = "europe-west8"
	PageTrendParamsRegionEuropeWest9         PageTrendParamsRegion = "europe-west9"
	PageTrendParamsRegionMeWest1             PageTrendParamsRegion = "me-west1"
	PageTrendParamsRegionSouthamericaEast1   PageTrendParamsRegion = "southamerica-east1"
	PageTrendParamsRegionUsCentral1          PageTrendParamsRegion = "us-central1"
	PageTrendParamsRegionUsEast1             PageTrendParamsRegion = "us-east1"
	PageTrendParamsRegionUsEast4             PageTrendParamsRegion = "us-east4"
	PageTrendParamsRegionUsSouth1            PageTrendParamsRegion = "us-south1"
	PageTrendParamsRegionUsWest1             PageTrendParamsRegion = "us-west1"
)

func (r PageTrendParamsRegion) IsKnown() bool {
	switch r {
	case PageTrendParamsRegionAsiaEast1, PageTrendParamsRegionAsiaNortheast1, PageTrendParamsRegionAsiaNortheast2, PageTrendParamsRegionAsiaSouth1, PageTrendParamsRegionAsiaSoutheast1, PageTrendParamsRegionAustraliaSoutheast1, PageTrendParamsRegionEuropeNorth1, PageTrendParamsRegionEuropeSouthwest1, PageTrendParamsRegionEuropeWest1, PageTrendParamsRegionEuropeWest2, PageTrendParamsRegionEuropeWest3, PageTrendParamsRegionEuropeWest4, PageTrendParamsRegionEuropeWest8, PageTrendParamsRegionEuropeWest9, PageTrendParamsRegionMeWest1, PageTrendParamsRegionSouthamericaEast1, PageTrendParamsRegionUsCentral1, PageTrendParamsRegionUsEast1, PageTrendParamsRegionUsEast4, PageTrendParamsRegionUsSouth1, PageTrendParamsRegionUsWest1:
		return true
	}
	return false
}

type PageTrendResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful.
	Success bool                          `json:"success,required"`
	Result  Trend                         `json:"result"`
	JSON    pageTrendResponseEnvelopeJSON `json:"-"`
}

// pageTrendResponseEnvelopeJSON contains the JSON metadata for the struct
// [PageTrendResponseEnvelope]
type pageTrendResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageTrendResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageTrendResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
