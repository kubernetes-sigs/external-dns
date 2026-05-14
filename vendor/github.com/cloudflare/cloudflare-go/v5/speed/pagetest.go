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

// PageTestService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPageTestService] method instead.
type PageTestService struct {
	Options []option.RequestOption
}

// NewPageTestService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPageTestService(opts ...option.RequestOption) (r *PageTestService) {
	r = &PageTestService{}
	r.Options = opts
	return
}

// Starts a test for a specific webpage, in a specific region.
func (r *PageTestService) New(ctx context.Context, url string, params PageTestNewParams, opts ...option.RequestOption) (res *Test, err error) {
	var env PageTestNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if url == "" {
		err = errors.New("missing required url parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/speed_api/pages/%s/tests", params.ZoneID, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Test history (list of tests) for a specific webpage.
func (r *PageTestService) List(ctx context.Context, url string, params PageTestListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[Test], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if url == "" {
		err = errors.New("missing required url parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/speed_api/pages/%s/tests", params.ZoneID, url)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// Test history (list of tests) for a specific webpage.
func (r *PageTestService) ListAutoPaging(ctx context.Context, url string, params PageTestListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[Test] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, url, params, opts...))
}

// Deletes all tests for a specific webpage from a specific region. Deleted tests
// are still counted as part of the quota.
func (r *PageTestService) Delete(ctx context.Context, url string, params PageTestDeleteParams, opts ...option.RequestOption) (res *PageTestDeleteResponse, err error) {
	var env PageTestDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if url == "" {
		err = errors.New("missing required url parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/speed_api/pages/%s/tests", params.ZoneID, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the result of a specific test.
func (r *PageTestService) Get(ctx context.Context, url string, testID string, query PageTestGetParams, opts ...option.RequestOption) (res *Test, err error) {
	var env PageTestGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if url == "" {
		err = errors.New("missing required url parameter")
		return
	}
	if testID == "" {
		err = errors.New("missing required test_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/speed_api/pages/%s/tests/%s", query.ZoneID, url, testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Test struct {
	// UUID.
	ID   string    `json:"id"`
	Date time.Time `json:"date" format:"date-time"`
	// The Lighthouse report.
	DesktopReport LighthouseReport `json:"desktopReport"`
	// The Lighthouse report.
	MobileReport LighthouseReport `json:"mobileReport"`
	// A test region with a label.
	Region LabeledRegion `json:"region"`
	// The frequency of the test.
	ScheduleFrequency TestScheduleFrequency `json:"scheduleFrequency"`
	// A URL.
	URL  string   `json:"url"`
	JSON testJSON `json:"-"`
}

// testJSON contains the JSON metadata for the struct [Test]
type testJSON struct {
	ID                apijson.Field
	Date              apijson.Field
	DesktopReport     apijson.Field
	MobileReport      apijson.Field
	Region            apijson.Field
	ScheduleFrequency apijson.Field
	URL               apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *Test) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testJSON) RawJSON() string {
	return r.raw
}

// The frequency of the test.
type TestScheduleFrequency string

const (
	TestScheduleFrequencyDaily  TestScheduleFrequency = "DAILY"
	TestScheduleFrequencyWeekly TestScheduleFrequency = "WEEKLY"
)

func (r TestScheduleFrequency) IsKnown() bool {
	switch r {
	case TestScheduleFrequencyDaily, TestScheduleFrequencyWeekly:
		return true
	}
	return false
}

type PageTestDeleteResponse struct {
	// Number of items affected.
	Count float64                    `json:"count"`
	JSON  pageTestDeleteResponseJSON `json:"-"`
}

// pageTestDeleteResponseJSON contains the JSON metadata for the struct
// [PageTestDeleteResponse]
type pageTestDeleteResponseJSON struct {
	Count       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageTestDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageTestDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type PageTestNewParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A test region.
	Region param.Field[PageTestNewParamsRegion] `json:"region"`
}

func (r PageTestNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A test region.
type PageTestNewParamsRegion string

const (
	PageTestNewParamsRegionAsiaEast1           PageTestNewParamsRegion = "asia-east1"
	PageTestNewParamsRegionAsiaNortheast1      PageTestNewParamsRegion = "asia-northeast1"
	PageTestNewParamsRegionAsiaNortheast2      PageTestNewParamsRegion = "asia-northeast2"
	PageTestNewParamsRegionAsiaSouth1          PageTestNewParamsRegion = "asia-south1"
	PageTestNewParamsRegionAsiaSoutheast1      PageTestNewParamsRegion = "asia-southeast1"
	PageTestNewParamsRegionAustraliaSoutheast1 PageTestNewParamsRegion = "australia-southeast1"
	PageTestNewParamsRegionEuropeNorth1        PageTestNewParamsRegion = "europe-north1"
	PageTestNewParamsRegionEuropeSouthwest1    PageTestNewParamsRegion = "europe-southwest1"
	PageTestNewParamsRegionEuropeWest1         PageTestNewParamsRegion = "europe-west1"
	PageTestNewParamsRegionEuropeWest2         PageTestNewParamsRegion = "europe-west2"
	PageTestNewParamsRegionEuropeWest3         PageTestNewParamsRegion = "europe-west3"
	PageTestNewParamsRegionEuropeWest4         PageTestNewParamsRegion = "europe-west4"
	PageTestNewParamsRegionEuropeWest8         PageTestNewParamsRegion = "europe-west8"
	PageTestNewParamsRegionEuropeWest9         PageTestNewParamsRegion = "europe-west9"
	PageTestNewParamsRegionMeWest1             PageTestNewParamsRegion = "me-west1"
	PageTestNewParamsRegionSouthamericaEast1   PageTestNewParamsRegion = "southamerica-east1"
	PageTestNewParamsRegionUsCentral1          PageTestNewParamsRegion = "us-central1"
	PageTestNewParamsRegionUsEast1             PageTestNewParamsRegion = "us-east1"
	PageTestNewParamsRegionUsEast4             PageTestNewParamsRegion = "us-east4"
	PageTestNewParamsRegionUsSouth1            PageTestNewParamsRegion = "us-south1"
	PageTestNewParamsRegionUsWest1             PageTestNewParamsRegion = "us-west1"
)

func (r PageTestNewParamsRegion) IsKnown() bool {
	switch r {
	case PageTestNewParamsRegionAsiaEast1, PageTestNewParamsRegionAsiaNortheast1, PageTestNewParamsRegionAsiaNortheast2, PageTestNewParamsRegionAsiaSouth1, PageTestNewParamsRegionAsiaSoutheast1, PageTestNewParamsRegionAustraliaSoutheast1, PageTestNewParamsRegionEuropeNorth1, PageTestNewParamsRegionEuropeSouthwest1, PageTestNewParamsRegionEuropeWest1, PageTestNewParamsRegionEuropeWest2, PageTestNewParamsRegionEuropeWest3, PageTestNewParamsRegionEuropeWest4, PageTestNewParamsRegionEuropeWest8, PageTestNewParamsRegionEuropeWest9, PageTestNewParamsRegionMeWest1, PageTestNewParamsRegionSouthamericaEast1, PageTestNewParamsRegionUsCentral1, PageTestNewParamsRegionUsEast1, PageTestNewParamsRegionUsEast4, PageTestNewParamsRegionUsSouth1, PageTestNewParamsRegionUsWest1:
		return true
	}
	return false
}

type PageTestNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful.
	Success bool                            `json:"success,required"`
	Result  Test                            `json:"result"`
	JSON    pageTestNewResponseEnvelopeJSON `json:"-"`
}

// pageTestNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [PageTestNewResponseEnvelope]
type pageTestNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageTestNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageTestNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type PageTestListParams struct {
	// Identifier.
	ZoneID  param.Field[string] `path:"zone_id,required"`
	Page    param.Field[int64]  `query:"page"`
	PerPage param.Field[int64]  `query:"per_page"`
	// A test region.
	Region param.Field[PageTestListParamsRegion] `query:"region"`
}

// URLQuery serializes [PageTestListParams]'s query parameters as `url.Values`.
func (r PageTestListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// A test region.
type PageTestListParamsRegion string

const (
	PageTestListParamsRegionAsiaEast1           PageTestListParamsRegion = "asia-east1"
	PageTestListParamsRegionAsiaNortheast1      PageTestListParamsRegion = "asia-northeast1"
	PageTestListParamsRegionAsiaNortheast2      PageTestListParamsRegion = "asia-northeast2"
	PageTestListParamsRegionAsiaSouth1          PageTestListParamsRegion = "asia-south1"
	PageTestListParamsRegionAsiaSoutheast1      PageTestListParamsRegion = "asia-southeast1"
	PageTestListParamsRegionAustraliaSoutheast1 PageTestListParamsRegion = "australia-southeast1"
	PageTestListParamsRegionEuropeNorth1        PageTestListParamsRegion = "europe-north1"
	PageTestListParamsRegionEuropeSouthwest1    PageTestListParamsRegion = "europe-southwest1"
	PageTestListParamsRegionEuropeWest1         PageTestListParamsRegion = "europe-west1"
	PageTestListParamsRegionEuropeWest2         PageTestListParamsRegion = "europe-west2"
	PageTestListParamsRegionEuropeWest3         PageTestListParamsRegion = "europe-west3"
	PageTestListParamsRegionEuropeWest4         PageTestListParamsRegion = "europe-west4"
	PageTestListParamsRegionEuropeWest8         PageTestListParamsRegion = "europe-west8"
	PageTestListParamsRegionEuropeWest9         PageTestListParamsRegion = "europe-west9"
	PageTestListParamsRegionMeWest1             PageTestListParamsRegion = "me-west1"
	PageTestListParamsRegionSouthamericaEast1   PageTestListParamsRegion = "southamerica-east1"
	PageTestListParamsRegionUsCentral1          PageTestListParamsRegion = "us-central1"
	PageTestListParamsRegionUsEast1             PageTestListParamsRegion = "us-east1"
	PageTestListParamsRegionUsEast4             PageTestListParamsRegion = "us-east4"
	PageTestListParamsRegionUsSouth1            PageTestListParamsRegion = "us-south1"
	PageTestListParamsRegionUsWest1             PageTestListParamsRegion = "us-west1"
)

func (r PageTestListParamsRegion) IsKnown() bool {
	switch r {
	case PageTestListParamsRegionAsiaEast1, PageTestListParamsRegionAsiaNortheast1, PageTestListParamsRegionAsiaNortheast2, PageTestListParamsRegionAsiaSouth1, PageTestListParamsRegionAsiaSoutheast1, PageTestListParamsRegionAustraliaSoutheast1, PageTestListParamsRegionEuropeNorth1, PageTestListParamsRegionEuropeSouthwest1, PageTestListParamsRegionEuropeWest1, PageTestListParamsRegionEuropeWest2, PageTestListParamsRegionEuropeWest3, PageTestListParamsRegionEuropeWest4, PageTestListParamsRegionEuropeWest8, PageTestListParamsRegionEuropeWest9, PageTestListParamsRegionMeWest1, PageTestListParamsRegionSouthamericaEast1, PageTestListParamsRegionUsCentral1, PageTestListParamsRegionUsEast1, PageTestListParamsRegionUsEast4, PageTestListParamsRegionUsSouth1, PageTestListParamsRegionUsWest1:
		return true
	}
	return false
}

type PageTestDeleteParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A test region.
	Region param.Field[PageTestDeleteParamsRegion] `query:"region"`
}

// URLQuery serializes [PageTestDeleteParams]'s query parameters as `url.Values`.
func (r PageTestDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// A test region.
type PageTestDeleteParamsRegion string

const (
	PageTestDeleteParamsRegionAsiaEast1           PageTestDeleteParamsRegion = "asia-east1"
	PageTestDeleteParamsRegionAsiaNortheast1      PageTestDeleteParamsRegion = "asia-northeast1"
	PageTestDeleteParamsRegionAsiaNortheast2      PageTestDeleteParamsRegion = "asia-northeast2"
	PageTestDeleteParamsRegionAsiaSouth1          PageTestDeleteParamsRegion = "asia-south1"
	PageTestDeleteParamsRegionAsiaSoutheast1      PageTestDeleteParamsRegion = "asia-southeast1"
	PageTestDeleteParamsRegionAustraliaSoutheast1 PageTestDeleteParamsRegion = "australia-southeast1"
	PageTestDeleteParamsRegionEuropeNorth1        PageTestDeleteParamsRegion = "europe-north1"
	PageTestDeleteParamsRegionEuropeSouthwest1    PageTestDeleteParamsRegion = "europe-southwest1"
	PageTestDeleteParamsRegionEuropeWest1         PageTestDeleteParamsRegion = "europe-west1"
	PageTestDeleteParamsRegionEuropeWest2         PageTestDeleteParamsRegion = "europe-west2"
	PageTestDeleteParamsRegionEuropeWest3         PageTestDeleteParamsRegion = "europe-west3"
	PageTestDeleteParamsRegionEuropeWest4         PageTestDeleteParamsRegion = "europe-west4"
	PageTestDeleteParamsRegionEuropeWest8         PageTestDeleteParamsRegion = "europe-west8"
	PageTestDeleteParamsRegionEuropeWest9         PageTestDeleteParamsRegion = "europe-west9"
	PageTestDeleteParamsRegionMeWest1             PageTestDeleteParamsRegion = "me-west1"
	PageTestDeleteParamsRegionSouthamericaEast1   PageTestDeleteParamsRegion = "southamerica-east1"
	PageTestDeleteParamsRegionUsCentral1          PageTestDeleteParamsRegion = "us-central1"
	PageTestDeleteParamsRegionUsEast1             PageTestDeleteParamsRegion = "us-east1"
	PageTestDeleteParamsRegionUsEast4             PageTestDeleteParamsRegion = "us-east4"
	PageTestDeleteParamsRegionUsSouth1            PageTestDeleteParamsRegion = "us-south1"
	PageTestDeleteParamsRegionUsWest1             PageTestDeleteParamsRegion = "us-west1"
)

func (r PageTestDeleteParamsRegion) IsKnown() bool {
	switch r {
	case PageTestDeleteParamsRegionAsiaEast1, PageTestDeleteParamsRegionAsiaNortheast1, PageTestDeleteParamsRegionAsiaNortheast2, PageTestDeleteParamsRegionAsiaSouth1, PageTestDeleteParamsRegionAsiaSoutheast1, PageTestDeleteParamsRegionAustraliaSoutheast1, PageTestDeleteParamsRegionEuropeNorth1, PageTestDeleteParamsRegionEuropeSouthwest1, PageTestDeleteParamsRegionEuropeWest1, PageTestDeleteParamsRegionEuropeWest2, PageTestDeleteParamsRegionEuropeWest3, PageTestDeleteParamsRegionEuropeWest4, PageTestDeleteParamsRegionEuropeWest8, PageTestDeleteParamsRegionEuropeWest9, PageTestDeleteParamsRegionMeWest1, PageTestDeleteParamsRegionSouthamericaEast1, PageTestDeleteParamsRegionUsCentral1, PageTestDeleteParamsRegionUsEast1, PageTestDeleteParamsRegionUsEast4, PageTestDeleteParamsRegionUsSouth1, PageTestDeleteParamsRegionUsWest1:
		return true
	}
	return false
}

type PageTestDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful.
	Success bool                               `json:"success,required"`
	Result  PageTestDeleteResponse             `json:"result"`
	JSON    pageTestDeleteResponseEnvelopeJSON `json:"-"`
}

// pageTestDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [PageTestDeleteResponseEnvelope]
type pageTestDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageTestDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageTestDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type PageTestGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type PageTestGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful.
	Success bool                            `json:"success,required"`
	Result  Test                            `json:"result"`
	JSON    pageTestGetResponseEnvelopeJSON `json:"-"`
}

// pageTestGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [PageTestGetResponseEnvelope]
type pageTestGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageTestGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageTestGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
