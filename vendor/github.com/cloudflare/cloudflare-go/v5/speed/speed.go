// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package speed

import (
	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// SpeedService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSpeedService] method instead.
type SpeedService struct {
	Options        []option.RequestOption
	Schedule       *ScheduleService
	Availabilities *AvailabilityService
	Pages          *PageService
}

// NewSpeedService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSpeedService(opts ...option.RequestOption) (r *SpeedService) {
	r = &SpeedService{}
	r.Options = opts
	r.Schedule = NewScheduleService(opts...)
	r.Availabilities = NewAvailabilityService(opts...)
	r.Pages = NewPageService(opts...)
	return
}

// A test region with a label.
type LabeledRegion struct {
	Label string `json:"label"`
	// A test region.
	Value LabeledRegionValue `json:"value"`
	JSON  labeledRegionJSON  `json:"-"`
}

// labeledRegionJSON contains the JSON metadata for the struct [LabeledRegion]
type labeledRegionJSON struct {
	Label       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LabeledRegion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r labeledRegionJSON) RawJSON() string {
	return r.raw
}

// A test region.
type LabeledRegionValue string

const (
	LabeledRegionValueAsiaEast1           LabeledRegionValue = "asia-east1"
	LabeledRegionValueAsiaNortheast1      LabeledRegionValue = "asia-northeast1"
	LabeledRegionValueAsiaNortheast2      LabeledRegionValue = "asia-northeast2"
	LabeledRegionValueAsiaSouth1          LabeledRegionValue = "asia-south1"
	LabeledRegionValueAsiaSoutheast1      LabeledRegionValue = "asia-southeast1"
	LabeledRegionValueAustraliaSoutheast1 LabeledRegionValue = "australia-southeast1"
	LabeledRegionValueEuropeNorth1        LabeledRegionValue = "europe-north1"
	LabeledRegionValueEuropeSouthwest1    LabeledRegionValue = "europe-southwest1"
	LabeledRegionValueEuropeWest1         LabeledRegionValue = "europe-west1"
	LabeledRegionValueEuropeWest2         LabeledRegionValue = "europe-west2"
	LabeledRegionValueEuropeWest3         LabeledRegionValue = "europe-west3"
	LabeledRegionValueEuropeWest4         LabeledRegionValue = "europe-west4"
	LabeledRegionValueEuropeWest8         LabeledRegionValue = "europe-west8"
	LabeledRegionValueEuropeWest9         LabeledRegionValue = "europe-west9"
	LabeledRegionValueMeWest1             LabeledRegionValue = "me-west1"
	LabeledRegionValueSouthamericaEast1   LabeledRegionValue = "southamerica-east1"
	LabeledRegionValueUsCentral1          LabeledRegionValue = "us-central1"
	LabeledRegionValueUsEast1             LabeledRegionValue = "us-east1"
	LabeledRegionValueUsEast4             LabeledRegionValue = "us-east4"
	LabeledRegionValueUsSouth1            LabeledRegionValue = "us-south1"
	LabeledRegionValueUsWest1             LabeledRegionValue = "us-west1"
)

func (r LabeledRegionValue) IsKnown() bool {
	switch r {
	case LabeledRegionValueAsiaEast1, LabeledRegionValueAsiaNortheast1, LabeledRegionValueAsiaNortheast2, LabeledRegionValueAsiaSouth1, LabeledRegionValueAsiaSoutheast1, LabeledRegionValueAustraliaSoutheast1, LabeledRegionValueEuropeNorth1, LabeledRegionValueEuropeSouthwest1, LabeledRegionValueEuropeWest1, LabeledRegionValueEuropeWest2, LabeledRegionValueEuropeWest3, LabeledRegionValueEuropeWest4, LabeledRegionValueEuropeWest8, LabeledRegionValueEuropeWest9, LabeledRegionValueMeWest1, LabeledRegionValueSouthamericaEast1, LabeledRegionValueUsCentral1, LabeledRegionValueUsEast1, LabeledRegionValueUsEast4, LabeledRegionValueUsSouth1, LabeledRegionValueUsWest1:
		return true
	}
	return false
}

// The Lighthouse report.
type LighthouseReport struct {
	// Cumulative Layout Shift.
	CLS float64 `json:"cls"`
	// The type of device.
	DeviceType LighthouseReportDeviceType `json:"deviceType"`
	Error      LighthouseReportError      `json:"error"`
	// First Contentful Paint.
	FCP float64 `json:"fcp"`
	// The URL to the full Lighthouse JSON report.
	JsonReportURL string `json:"jsonReportUrl"`
	// Largest Contentful Paint.
	LCP float64 `json:"lcp"`
	// The Lighthouse performance score.
	PerformanceScore float64 `json:"performanceScore"`
	// Speed Index.
	Si float64 `json:"si"`
	// The state of the Lighthouse report.
	State LighthouseReportState `json:"state"`
	// Total Blocking Time.
	TBT float64 `json:"tbt"`
	// Time To First Byte.
	TTFB float64 `json:"ttfb"`
	// Time To Interactive.
	TTI  float64              `json:"tti"`
	JSON lighthouseReportJSON `json:"-"`
}

// lighthouseReportJSON contains the JSON metadata for the struct
// [LighthouseReport]
type lighthouseReportJSON struct {
	CLS              apijson.Field
	DeviceType       apijson.Field
	Error            apijson.Field
	FCP              apijson.Field
	JsonReportURL    apijson.Field
	LCP              apijson.Field
	PerformanceScore apijson.Field
	Si               apijson.Field
	State            apijson.Field
	TBT              apijson.Field
	TTFB             apijson.Field
	TTI              apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LighthouseReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lighthouseReportJSON) RawJSON() string {
	return r.raw
}

// The type of device.
type LighthouseReportDeviceType string

const (
	LighthouseReportDeviceTypeDesktop LighthouseReportDeviceType = "DESKTOP"
	LighthouseReportDeviceTypeMobile  LighthouseReportDeviceType = "MOBILE"
)

func (r LighthouseReportDeviceType) IsKnown() bool {
	switch r {
	case LighthouseReportDeviceTypeDesktop, LighthouseReportDeviceTypeMobile:
		return true
	}
	return false
}

type LighthouseReportError struct {
	// The error code of the Lighthouse result.
	Code LighthouseReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                    `json:"finalDisplayedUrl"`
	JSON              lighthouseReportErrorJSON `json:"-"`
}

// lighthouseReportErrorJSON contains the JSON metadata for the struct
// [LighthouseReportError]
type lighthouseReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *LighthouseReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lighthouseReportErrorJSON) RawJSON() string {
	return r.raw
}

// The error code of the Lighthouse result.
type LighthouseReportErrorCode string

const (
	LighthouseReportErrorCodeNotReachable      LighthouseReportErrorCode = "NOT_REACHABLE"
	LighthouseReportErrorCodeDNSFailure        LighthouseReportErrorCode = "DNS_FAILURE"
	LighthouseReportErrorCodeNotHTML           LighthouseReportErrorCode = "NOT_HTML"
	LighthouseReportErrorCodeLighthouseTimeout LighthouseReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	LighthouseReportErrorCodeUnknown           LighthouseReportErrorCode = "UNKNOWN"
)

func (r LighthouseReportErrorCode) IsKnown() bool {
	switch r {
	case LighthouseReportErrorCodeNotReachable, LighthouseReportErrorCodeDNSFailure, LighthouseReportErrorCodeNotHTML, LighthouseReportErrorCodeLighthouseTimeout, LighthouseReportErrorCodeUnknown:
		return true
	}
	return false
}

// The state of the Lighthouse report.
type LighthouseReportState string

const (
	LighthouseReportStateRunning  LighthouseReportState = "RUNNING"
	LighthouseReportStateComplete LighthouseReportState = "COMPLETE"
	LighthouseReportStateFailed   LighthouseReportState = "FAILED"
)

func (r LighthouseReportState) IsKnown() bool {
	switch r {
	case LighthouseReportStateRunning, LighthouseReportStateComplete, LighthouseReportStateFailed:
		return true
	}
	return false
}

type Trend struct {
	// Cumulative Layout Shift trend.
	CLS []float64 `json:"cls"`
	// First Contentful Paint trend.
	FCP []float64 `json:"fcp"`
	// Largest Contentful Paint trend.
	LCP []float64 `json:"lcp"`
	// The Lighthouse score trend.
	PerformanceScore []float64 `json:"performanceScore"`
	// Speed Index trend.
	Si []float64 `json:"si"`
	// Total Blocking Time trend.
	TBT []float64 `json:"tbt"`
	// Time To First Byte trend.
	TTFB []float64 `json:"ttfb"`
	// Time To Interactive trend.
	TTI  []float64 `json:"tti"`
	JSON trendJSON `json:"-"`
}

// trendJSON contains the JSON metadata for the struct [Trend]
type trendJSON struct {
	CLS              apijson.Field
	FCP              apijson.Field
	LCP              apijson.Field
	PerformanceScore apijson.Field
	Si               apijson.Field
	TBT              apijson.Field
	TTFB             apijson.Field
	TTI              apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *Trend) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r trendJSON) RawJSON() string {
	return r.raw
}
