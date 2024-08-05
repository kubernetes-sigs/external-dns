package cloudflare

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/goccy/go-json"
	"github.com/google/go-querystring/query"
)

var (
	ErrMissingObservatoryUrl    = errors.New("missing required page url")
	ErrMissingObservatoryTestID = errors.New("missing required test id")
)

// ObservatoryPage describes all the tests for a web page.
type ObservatoryPage struct {
	URL               string                `json:"url"`
	Region            labeledRegion         `json:"region"`
	ScheduleFrequency string                `json:"scheduleFrequency"`
	Tests             []ObservatoryPageTest `json:"tests"`
}

// ObservatoryPageTest describes a single test for a web page.
type ObservatoryPageTest struct {
	ID                string                      `json:"id"`
	Date              *time.Time                  `json:"date"`
	URL               string                      `json:"url"`
	Region            labeledRegion               `json:"region"`
	ScheduleFrequency *string                     `json:"scheduleFrequency"`
	MobileReport      ObservatoryLighthouseReport `json:"mobileReport"`
	DesktopReport     ObservatoryLighthouseReport `json:"desktopReport"`
}

// labeledRegion describes a region the test was run in.
type labeledRegion struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

// ObservatorySchedule describe a test schedule.
type ObservatorySchedule struct {
	URL       string `json:"url"`
	Region    string `json:"region"`
	Frequency string `json:"frequency"`
}

// ObservatoryLighthouseReport describes the web vital metrics result.
type ObservatoryLighthouseReport struct {
	PerformanceScore int    `json:"performanceScore"`
	State            string `json:"state"`
	DeviceType       string `json:"deviceType"`
	// TTFB is time to first byte
	TTFB int `json:"ttfb"`
	// FCP is first contentful paint
	FCP int `json:"fcp"`
	// LCP is largest contentful pain
	LCP int `json:"lcp"`
	// TTI is time to interactive
	TTI int `json:"tti"`
	// TBT is total blocking time
	TBT int `json:"tbt"`
	// SI is speed index
	SI int `json:"si"`
	// CLS is cumulative layout shift
	CLS   float64          `json:"cls"`
	Error *lighthouseError `json:"error,omitempty"`
}

// lighthouseError describes the test error.
type lighthouseError struct {
	Code              string `json:"code"`
	Detail            string `json:"detail"`
	FinalDisplayedURL string `json:"finalDisplayedUrl"`
}

// ObservatoryPageTrend describes the web vital metrics trend.
type ObservatoryPageTrend struct {
	PerformanceScore []*int     `json:"performanceScore"`
	TTFB             []*int     `json:"ttfb"`
	FCP              []*int     `json:"fcp"`
	LCP              []*int     `json:"lcp"`
	TTI              []*int     `json:"tti"`
	TBT              []*int     `json:"tbt"`
	SI               []*int     `json:"si"`
	CLS              []*float64 `json:"cls"`
}

type ListObservatoryPagesParams struct {
}

// ObservatoryPagesResponse is the API response, containing a list of ObservatoryPage.
type ObservatoryPagesResponse struct {
	Response
	Result []ObservatoryPage `json:"result"`
}

// ListObservatoryPages returns a list of pages which have been tested.
//
// API reference: https://api.cloudflare.com/#speed-list-pages
func (api *API) ListObservatoryPages(ctx context.Context, rc *ResourceContainer, params ListObservatoryPagesParams) ([]ObservatoryPage, error) {
	uri := fmt.Sprintf("/zones/%s/speed_api/pages", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	var r ObservatoryPagesResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

type GetObservatoryPageTrendParams struct {
	URL        string     `url:"-"`
	Region     string     `url:"region"`
	DeviceType string     `url:"deviceType"`
	Start      *time.Time `url:"start"`
	End        *time.Time `url:"end,omitempty"`
	Timezone   string     `url:"tz"`
	Metrics    []string   `url:"metrics"`
}

type ObservatoryPageTrendResponse struct {
	Response
	Result ObservatoryPageTrend `json:"result"`
}

// GetObservatoryPageTrend returns a the trend of web vital metrics for a page in a specific region.
//
// API reference: https://api.cloudflare.com/#speed-list-page-trend
func (api *API) GetObservatoryPageTrend(ctx context.Context, rc *ResourceContainer, params GetObservatoryPageTrendParams) (*ObservatoryPageTrend, error) {
	if params.URL == "" {
		return nil, ErrMissingObservatoryUrl
	}
	// cannot use buildURI because params.URL contains "/" that should be encoded and buildURI will double encode %2F into %252F
	v, _ := query.Values(params)
	uri := fmt.Sprintf("/zones/%s/speed_api/pages/%s/trend?%s", rc.Identifier, url.PathEscape(params.URL), v.Encode())
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	var r ObservatoryPageTrendResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return &r.Result, nil
}

var listObservatoryPageTestDefaultPageSize = 20

type ListObservatoryPageTestParams struct {
	URL    string `url:"-"`
	Region string `url:"region"`
	ResultInfo
}

type ObservatoryPageTestsResponse struct {
	Response
	Result     []ObservatoryPageTest `json:"result"`
	ResultInfo `json:"result_info"`
}

// ListObservatoryPageTests returns a list of tests for a page in a specific region.
//
// API reference: https://api.cloudflare.com/#speed-list-test-history
func (api *API) ListObservatoryPageTests(ctx context.Context, rc *ResourceContainer, params ListObservatoryPageTestParams) ([]ObservatoryPageTest, *ResultInfo, error) {
	if params.URL == "" {
		return nil, nil, ErrMissingObservatoryUrl
	}
	autoPaginate := true
	if params.PerPage >= 1 || params.Page >= 1 {
		autoPaginate = false
	}
	if params.PerPage < 1 {
		params.PerPage = listObservatoryPageTestDefaultPageSize
	}
	if params.Page < 1 {
		params.Page = 1
	}
	var tests []ObservatoryPageTest
	var lastResultInfo ResultInfo
	for {
		// cannot use buildURI because params.URL contains "/" that should be encoded and buildURI will double encode %2F into %252F
		v, _ := query.Values(params)
		uri := fmt.Sprintf("/zones/%s/speed_api/pages/%s/tests?%s", rc.Identifier, url.PathEscape(params.URL), v.Encode())
		res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
		if err != nil {
			return nil, nil, err
		}
		var r ObservatoryPageTestsResponse
		err = json.Unmarshal(res, &r)
		if err != nil {
			return nil, nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
		}
		tests = append(tests, r.Result...)
		lastResultInfo = r.ResultInfo.Next()
		if params.ResultInfo.Done() || !autoPaginate {
			break
		}
	}
	return tests, &lastResultInfo, nil
}

type CreateObservatoryPageTestParams struct {
	URL      string
	Settings CreateObservatoryPageTestSettings
}
type CreateObservatoryPageTestSettings struct {
	Region string `json:"region"`
}

type ObservatoryPageTestResponse struct {
	Response
	Result ObservatoryPageTest `json:"result"`
}

// CreateObservatoryPageTest starts a test for a page in a specific region.
//
// API reference: https://api.cloudflare.com/#speed-create-test
func (api *API) CreateObservatoryPageTest(ctx context.Context, rc *ResourceContainer, params CreateObservatoryPageTestParams) (*ObservatoryPageTest, error) {
	if params.URL == "" {
		return nil, ErrMissingObservatoryUrl
	}
	uri := fmt.Sprintf("/zones/%s/speed_api/pages/%s/tests", rc.Identifier, url.PathEscape(params.URL))
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params.Settings)
	if err != nil {
		return nil, err
	}
	var r ObservatoryPageTestResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return &r.Result, nil
}

type DeleteObservatoryPageTestsParams struct {
	URL    string `url:"-"`
	Region string `url:"region"`
}

type ObservatoryCountResponse struct {
	Response
	Result struct {
		Count int `json:"count"`
	} `json:"result"`
}

// DeleteObservatoryPageTests deletes all tests for a page in a specific region.
//
// API reference: https://api.cloudflare.com/#speed-delete-tests
func (api *API) DeleteObservatoryPageTests(ctx context.Context, rc *ResourceContainer, params DeleteObservatoryPageTestsParams) (*int, error) {
	if params.URL == "" {
		return nil, ErrMissingObservatoryUrl
	}
	// cannot use buildURI because params.URL contains "/" that should be encoded and buildURI will double encode %2F into %252F
	v, _ := query.Values(params)
	uri := fmt.Sprintf("/zones/%s/speed_api/pages/%s/tests?%s", rc.Identifier, url.PathEscape(params.URL), v.Encode())
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return nil, err
	}
	var r ObservatoryCountResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return &r.Result.Count, nil
}

type GetObservatoryPageTestParams struct {
	URL    string
	TestID string
}

// GetObservatoryPageTest returns a specific test for a page.
//
// API reference: https://api.cloudflare.com/#speed-get-test
func (api *API) GetObservatoryPageTest(ctx context.Context, rc *ResourceContainer, params GetObservatoryPageTestParams) (*ObservatoryPageTest, error) {
	if params.URL == "" {
		return nil, ErrMissingObservatoryUrl
	}
	if params.TestID == "" {
		return nil, ErrMissingObservatoryTestID
	}
	uri := fmt.Sprintf("/zones/%s/speed_api/pages/%s/tests/%s", rc.Identifier, url.PathEscape(params.URL), params.TestID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	var r ObservatoryPageTestResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return &r.Result, nil
}

type CreateObservatoryScheduledPageTestParams struct {
	URL       string `url:"-" json:"-"`
	Region    string `url:"region" json:"-"`
	Frequency string `url:"frequency" json:"-"`
}

type ObservatoryScheduledPageTest struct {
	Schedule ObservatorySchedule `json:"schedule"`
	Test     ObservatoryPageTest `json:"test"`
}

type CreateObservatoryScheduledPageTestResponse struct {
	Response
	Result ObservatoryScheduledPageTest `json:"result"`
}

// CreateObservatoryScheduledPageTest creates a scheduled test for a page in a specific region.
//
// API reference: https://api.cloudflare.com/#speed-create-scheduled-test
func (api *API) CreateObservatoryScheduledPageTest(ctx context.Context, rc *ResourceContainer, params CreateObservatoryScheduledPageTestParams) (*ObservatoryScheduledPageTest, error) {
	if params.URL == "" {
		return nil, ErrMissingObservatoryUrl
	}
	// cannot use buildURI because params.URL contains "/" that should be encoded and buildURI will double encode %2F into %252F
	v, _ := query.Values(params)
	uri := fmt.Sprintf("/zones/%s/speed_api/schedule/%s?%s", rc.Identifier, url.PathEscape(params.URL), v.Encode())
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return nil, err
	}
	var r CreateObservatoryScheduledPageTestResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return &r.Result, nil
}

type GetObservatoryScheduledPageTestParams struct {
	URL    string `url:"-"`
	Region string `url:"region"`
}

type ObservatoryScheduleResponse struct {
	Response
	Result ObservatorySchedule `json:"result"`
}

// GetObservatoryScheduledPageTest returns the test schedule for a page in a specific region.
//
// API reference: https://api.cloudflare.com/#speed-get-scheduled-test
func (api *API) GetObservatoryScheduledPageTest(ctx context.Context, rc *ResourceContainer, params GetObservatoryScheduledPageTestParams) (*ObservatorySchedule, error) {
	if params.URL == "" {
		return nil, ErrMissingObservatoryUrl
	}
	// cannot use buildURI because params.URL contains "/" that should be encoded and buildURI will double encode %2F into %252F
	v, _ := query.Values(params)
	uri := fmt.Sprintf("/zones/%s/speed_api/schedule/%s?%s", rc.Identifier, url.PathEscape(params.URL), v.Encode())
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	var r ObservatoryScheduleResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return &r.Result, nil
}

type DeleteObservatoryScheduledPageTestParams struct {
	URL    string `url:"-"`
	Region string `url:"region"`
}

// DeleteObservatoryScheduledPageTest deletes the test schedule for a page in a specific region.
//
// API reference: https://api.cloudflare.com/#speed-delete-scheduled-test
func (api *API) DeleteObservatoryScheduledPageTest(ctx context.Context, rc *ResourceContainer, params DeleteObservatoryScheduledPageTestParams) (*int, error) {
	if params.URL == "" {
		return nil, ErrMissingObservatoryUrl
	}
	// cannot use buildURI because params.URL contains "/" that should be encoded and buildURI will double encode %2F into %252F
	v, _ := query.Values(params)
	uri := fmt.Sprintf("/zones/%s/speed_api/schedule/%s?%s", rc.Identifier, url.PathEscape(params.URL), v.Encode())
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return nil, err
	}
	var r ObservatoryCountResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return &r.Result.Count, nil
}
