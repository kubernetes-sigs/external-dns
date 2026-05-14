// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package url_scanner

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

// ScanService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScanService] method instead.
type ScanService struct {
	Options []option.RequestOption
}

// NewScanService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewScanService(opts ...option.RequestOption) (r *ScanService) {
	r = &ScanService{}
	r.Options = opts
	return
}

// Submit a URL to scan. Check limits at
// https://developers.cloudflare.com/security-center/investigate/scan-limits/.
func (r *ScanService) New(ctx context.Context, params ScanNewParams, opts ...option.RequestOption) (res *ScanNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/urlscanner/v2/scan", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Use a subset of ElasticSearch Query syntax to filter scans. Some example
// queries:<br/> <br/>- 'path:"/bundles/jquery.js"': Searches for scans who
// requested resources with the given path.<br/>- 'page.asn:AS24940 AND hash:xxx':
// Websites hosted in AS24940 where a resource with the given hash was
// downloaded.<br/>- 'page.domain:microsoft\* AND verdicts.malicious:true AND NOT
// page.domain:microsoft.com': malicious scans whose hostname starts with
// "microsoft".<br/>- 'apikey:me AND date:[2025-01 TO 2025-02]': my scans from 2025
// January to 2025 February.
func (r *ScanService) List(ctx context.Context, params ScanListParams, opts ...option.RequestOption) (res *ScanListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/urlscanner/v2/search", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Submit URLs to scan. Check limits at
// https://developers.cloudflare.com/security-center/investigate/scan-limits/ and
// take into account scans submitted in bulk have lower priority and may take
// longer to finish.
func (r *ScanService) BulkNew(ctx context.Context, params ScanBulkNewParams, opts ...option.RequestOption) (res *[]ScanBulkNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/urlscanner/v2/bulk", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Returns a plain text response, with the scan's DOM content as rendered by
// Chrome.
func (r *ScanService) DOM(ctx context.Context, scanID string, query ScanDOMParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scanID == "" {
		err = errors.New("missing required scan_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/urlscanner/v2/dom/%s", query.AccountID, scanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get URL scan by uuid
func (r *ScanService) Get(ctx context.Context, scanID string, query ScanGetParams, opts ...option.RequestOption) (res *ScanGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scanID == "" {
		err = errors.New("missing required scan_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/urlscanner/v2/result/%s", query.AccountID, scanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get a URL scan's HAR file. See HAR spec at
// http://www.softwareishard.com/blog/har-12-spec/.
func (r *ScanService) HAR(ctx context.Context, scanID string, query ScanHARParams, opts ...option.RequestOption) (res *ScanHARResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scanID == "" {
		err = errors.New("missing required scan_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/urlscanner/v2/har/%s", query.AccountID, scanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get scan's screenshot by resolution (desktop/mobile/tablet).
func (r *ScanService) Screenshot(ctx context.Context, scanID string, params ScanScreenshotParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "image/png")}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scanID == "" {
		err = errors.New("missing required scan_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/urlscanner/v2/screenshots/%s.png", params.AccountID, scanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

type ScanNewResponse struct {
	// URL to api report.
	API     string `json:"api,required"`
	Message string `json:"message,required"`
	// Public URL to report.
	Result string `json:"result,required"`
	// Canonical form of submitted URL. Use this if you want to later search by URL.
	URL string `json:"url,required"`
	// Scan ID.
	UUID string `json:"uuid,required" format:"uuid"`
	// Submitted visibility status.
	Visibility ScanNewResponseVisibility `json:"visibility,required"`
	Options    ScanNewResponseOptions    `json:"options"`
	JSON       scanNewResponseJSON       `json:"-"`
}

// scanNewResponseJSON contains the JSON metadata for the struct [ScanNewResponse]
type scanNewResponseJSON struct {
	API         apijson.Field
	Message     apijson.Field
	Result      apijson.Field
	URL         apijson.Field
	UUID        apijson.Field
	Visibility  apijson.Field
	Options     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanNewResponseJSON) RawJSON() string {
	return r.raw
}

// Submitted visibility status.
type ScanNewResponseVisibility string

const (
	ScanNewResponseVisibilityPublic   ScanNewResponseVisibility = "public"
	ScanNewResponseVisibilityUnlisted ScanNewResponseVisibility = "unlisted"
)

func (r ScanNewResponseVisibility) IsKnown() bool {
	switch r {
	case ScanNewResponseVisibilityPublic, ScanNewResponseVisibilityUnlisted:
		return true
	}
	return false
}

type ScanNewResponseOptions struct {
	Useragent string                     `json:"useragent"`
	JSON      scanNewResponseOptionsJSON `json:"-"`
}

// scanNewResponseOptionsJSON contains the JSON metadata for the struct
// [ScanNewResponseOptions]
type scanNewResponseOptionsJSON struct {
	Useragent   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanNewResponseOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanNewResponseOptionsJSON) RawJSON() string {
	return r.raw
}

type ScanListResponse struct {
	Results []ScanListResponseResult `json:"results,required"`
	JSON    scanListResponseJSON     `json:"-"`
}

// scanListResponseJSON contains the JSON metadata for the struct
// [ScanListResponse]
type scanListResponseJSON struct {
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanListResponseJSON) RawJSON() string {
	return r.raw
}

type ScanListResponseResult struct {
	ID       string                          `json:"_id,required"`
	Page     ScanListResponseResultsPage     `json:"page,required"`
	Result   string                          `json:"result,required"`
	Stats    ScanListResponseResultsStats    `json:"stats,required"`
	Task     ScanListResponseResultsTask     `json:"task,required"`
	Verdicts ScanListResponseResultsVerdicts `json:"verdicts,required"`
	JSON     scanListResponseResultJSON      `json:"-"`
}

// scanListResponseResultJSON contains the JSON metadata for the struct
// [ScanListResponseResult]
type scanListResponseResultJSON struct {
	ID          apijson.Field
	Page        apijson.Field
	Result      apijson.Field
	Stats       apijson.Field
	Task        apijson.Field
	Verdicts    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanListResponseResultJSON) RawJSON() string {
	return r.raw
}

type ScanListResponseResultsPage struct {
	ASN     string                          `json:"asn,required"`
	Country string                          `json:"country,required"`
	IP      string                          `json:"ip,required"`
	URL     string                          `json:"url,required"`
	JSON    scanListResponseResultsPageJSON `json:"-"`
}

// scanListResponseResultsPageJSON contains the JSON metadata for the struct
// [ScanListResponseResultsPage]
type scanListResponseResultsPageJSON struct {
	ASN         apijson.Field
	Country     apijson.Field
	IP          apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanListResponseResultsPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanListResponseResultsPageJSON) RawJSON() string {
	return r.raw
}

type ScanListResponseResultsStats struct {
	DataLength    float64                          `json:"dataLength,required"`
	Requests      float64                          `json:"requests,required"`
	UniqCountries float64                          `json:"uniqCountries,required"`
	UniqIPs       float64                          `json:"uniqIPs,required"`
	JSON          scanListResponseResultsStatsJSON `json:"-"`
}

// scanListResponseResultsStatsJSON contains the JSON metadata for the struct
// [ScanListResponseResultsStats]
type scanListResponseResultsStatsJSON struct {
	DataLength    apijson.Field
	Requests      apijson.Field
	UniqCountries apijson.Field
	UniqIPs       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ScanListResponseResultsStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanListResponseResultsStatsJSON) RawJSON() string {
	return r.raw
}

type ScanListResponseResultsTask struct {
	Time       string                          `json:"time,required"`
	URL        string                          `json:"url,required"`
	UUID       string                          `json:"uuid,required"`
	Visibility string                          `json:"visibility,required"`
	JSON       scanListResponseResultsTaskJSON `json:"-"`
}

// scanListResponseResultsTaskJSON contains the JSON metadata for the struct
// [ScanListResponseResultsTask]
type scanListResponseResultsTaskJSON struct {
	Time        apijson.Field
	URL         apijson.Field
	UUID        apijson.Field
	Visibility  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanListResponseResultsTask) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanListResponseResultsTaskJSON) RawJSON() string {
	return r.raw
}

type ScanListResponseResultsVerdicts struct {
	Malicious bool                                `json:"malicious,required"`
	JSON      scanListResponseResultsVerdictsJSON `json:"-"`
}

// scanListResponseResultsVerdictsJSON contains the JSON metadata for the struct
// [ScanListResponseResultsVerdicts]
type scanListResponseResultsVerdictsJSON struct {
	Malicious   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanListResponseResultsVerdicts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanListResponseResultsVerdictsJSON) RawJSON() string {
	return r.raw
}

type ScanBulkNewResponse struct {
	// URL to api report.
	API string `json:"api,required"`
	// URL to report.
	Result string `json:"result,required"`
	// Submitted URL
	URL string `json:"url,required"`
	// Scan ID.
	UUID string `json:"uuid,required" format:"uuid"`
	// Submitted visibility status.
	Visibility ScanBulkNewResponseVisibility `json:"visibility,required"`
	Options    ScanBulkNewResponseOptions    `json:"options"`
	JSON       scanBulkNewResponseJSON       `json:"-"`
}

// scanBulkNewResponseJSON contains the JSON metadata for the struct
// [ScanBulkNewResponse]
type scanBulkNewResponseJSON struct {
	API         apijson.Field
	Result      apijson.Field
	URL         apijson.Field
	UUID        apijson.Field
	Visibility  apijson.Field
	Options     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanBulkNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanBulkNewResponseJSON) RawJSON() string {
	return r.raw
}

// Submitted visibility status.
type ScanBulkNewResponseVisibility string

const (
	ScanBulkNewResponseVisibilityPublic   ScanBulkNewResponseVisibility = "public"
	ScanBulkNewResponseVisibilityUnlisted ScanBulkNewResponseVisibility = "unlisted"
)

func (r ScanBulkNewResponseVisibility) IsKnown() bool {
	switch r {
	case ScanBulkNewResponseVisibilityPublic, ScanBulkNewResponseVisibilityUnlisted:
		return true
	}
	return false
}

type ScanBulkNewResponseOptions struct {
	Useragent string                         `json:"useragent"`
	JSON      scanBulkNewResponseOptionsJSON `json:"-"`
}

// scanBulkNewResponseOptionsJSON contains the JSON metadata for the struct
// [ScanBulkNewResponseOptions]
type scanBulkNewResponseOptionsJSON struct {
	Useragent   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanBulkNewResponseOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanBulkNewResponseOptionsJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponse struct {
	Data     ScanGetResponseData     `json:"data,required"`
	Lists    ScanGetResponseLists    `json:"lists,required"`
	Meta     ScanGetResponseMeta     `json:"meta,required"`
	Page     ScanGetResponsePage     `json:"page,required"`
	Scanner  ScanGetResponseScanner  `json:"scanner,required"`
	Stats    ScanGetResponseStats    `json:"stats,required"`
	Task     ScanGetResponseTask     `json:"task,required"`
	Verdicts ScanGetResponseVerdicts `json:"verdicts,required"`
	JSON     scanGetResponseJSON     `json:"-"`
}

// scanGetResponseJSON contains the JSON metadata for the struct [ScanGetResponse]
type scanGetResponseJSON struct {
	Data        apijson.Field
	Lists       apijson.Field
	Meta        apijson.Field
	Page        apijson.Field
	Scanner     apijson.Field
	Stats       apijson.Field
	Task        apijson.Field
	Verdicts    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseData struct {
	Console     []ScanGetResponseDataConsole     `json:"console,required"`
	Cookies     []ScanGetResponseDataCookie      `json:"cookies,required"`
	Globals     []ScanGetResponseDataGlobal      `json:"globals,required"`
	Links       []ScanGetResponseDataLink        `json:"links,required"`
	Performance []ScanGetResponseDataPerformance `json:"performance,required"`
	Requests    []ScanGetResponseDataRequest     `json:"requests,required"`
	JSON        scanGetResponseDataJSON          `json:"-"`
}

// scanGetResponseDataJSON contains the JSON metadata for the struct
// [ScanGetResponseData]
type scanGetResponseDataJSON struct {
	Console     apijson.Field
	Cookies     apijson.Field
	Globals     apijson.Field
	Links       apijson.Field
	Performance apijson.Field
	Requests    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseDataConsole struct {
	Message ScanGetResponseDataConsoleMessage `json:"message,required"`
	JSON    scanGetResponseDataConsoleJSON    `json:"-"`
}

// scanGetResponseDataConsoleJSON contains the JSON metadata for the struct
// [ScanGetResponseDataConsole]
type scanGetResponseDataConsoleJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseDataConsole) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseDataConsoleJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseDataConsoleMessage struct {
	Level  string                                `json:"level,required"`
	Source string                                `json:"source,required"`
	Text   string                                `json:"text,required"`
	URL    string                                `json:"url,required"`
	JSON   scanGetResponseDataConsoleMessageJSON `json:"-"`
}

// scanGetResponseDataConsoleMessageJSON contains the JSON metadata for the struct
// [ScanGetResponseDataConsoleMessage]
type scanGetResponseDataConsoleMessageJSON struct {
	Level       apijson.Field
	Source      apijson.Field
	Text        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseDataConsoleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseDataConsoleMessageJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseDataCookie struct {
	Domain       string                        `json:"domain,required"`
	Expires      float64                       `json:"expires,required"`
	HTTPOnly     bool                          `json:"httpOnly,required"`
	Name         string                        `json:"name,required"`
	Path         string                        `json:"path,required"`
	Priority     string                        `json:"priority,required"`
	SameParty    bool                          `json:"sameParty,required"`
	Secure       bool                          `json:"secure,required"`
	Session      bool                          `json:"session,required"`
	Size         float64                       `json:"size,required"`
	SourcePort   float64                       `json:"sourcePort,required"`
	SourceScheme string                        `json:"sourceScheme,required"`
	Value        string                        `json:"value,required"`
	JSON         scanGetResponseDataCookieJSON `json:"-"`
}

// scanGetResponseDataCookieJSON contains the JSON metadata for the struct
// [ScanGetResponseDataCookie]
type scanGetResponseDataCookieJSON struct {
	Domain       apijson.Field
	Expires      apijson.Field
	HTTPOnly     apijson.Field
	Name         apijson.Field
	Path         apijson.Field
	Priority     apijson.Field
	SameParty    apijson.Field
	Secure       apijson.Field
	Session      apijson.Field
	Size         apijson.Field
	SourcePort   apijson.Field
	SourceScheme apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ScanGetResponseDataCookie) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseDataCookieJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseDataGlobal struct {
	Prop string                        `json:"prop,required"`
	Type string                        `json:"type,required"`
	JSON scanGetResponseDataGlobalJSON `json:"-"`
}

// scanGetResponseDataGlobalJSON contains the JSON metadata for the struct
// [ScanGetResponseDataGlobal]
type scanGetResponseDataGlobalJSON struct {
	Prop        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseDataGlobal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseDataGlobalJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseDataLink struct {
	Href string                      `json:"href,required"`
	Text string                      `json:"text,required"`
	JSON scanGetResponseDataLinkJSON `json:"-"`
}

// scanGetResponseDataLinkJSON contains the JSON metadata for the struct
// [ScanGetResponseDataLink]
type scanGetResponseDataLinkJSON struct {
	Href        apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseDataLink) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseDataLinkJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseDataPerformance struct {
	Duration  float64                            `json:"duration,required"`
	EntryType string                             `json:"entryType,required"`
	Name      string                             `json:"name,required"`
	StartTime float64                            `json:"startTime,required"`
	JSON      scanGetResponseDataPerformanceJSON `json:"-"`
}

// scanGetResponseDataPerformanceJSON contains the JSON metadata for the struct
// [ScanGetResponseDataPerformance]
type scanGetResponseDataPerformanceJSON struct {
	Duration    apijson.Field
	EntryType   apijson.Field
	Name        apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseDataPerformance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseDataPerformanceJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseDataRequest struct {
	Request  ScanGetResponseDataRequestsRequest   `json:"request,required"`
	Response ScanGetResponseDataRequestsResponse  `json:"response,required"`
	Requests []ScanGetResponseDataRequestsRequest `json:"requests"`
	JSON     scanGetResponseDataRequestJSON       `json:"-"`
}

// scanGetResponseDataRequestJSON contains the JSON metadata for the struct
// [ScanGetResponseDataRequest]
type scanGetResponseDataRequestJSON struct {
	Request     apijson.Field
	Response    apijson.Field
	Requests    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseDataRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseDataRequestJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseDataRequestsRequest struct {
	DocumentURL          string                                             `json:"documentURL,required"`
	HasUserGesture       bool                                               `json:"hasUserGesture,required"`
	Initiator            ScanGetResponseDataRequestsRequestInitiator        `json:"initiator,required"`
	RedirectHasExtraInfo bool                                               `json:"redirectHasExtraInfo,required"`
	Request              ScanGetResponseDataRequestsRequestRequest          `json:"request,required"`
	RequestID            string                                             `json:"requestId,required"`
	Type                 string                                             `json:"type,required"`
	WallTime             float64                                            `json:"wallTime,required"`
	FrameID              string                                             `json:"frameId"`
	LoaderID             string                                             `json:"loaderId"`
	PrimaryRequest       bool                                               `json:"primaryRequest"`
	RedirectResponse     ScanGetResponseDataRequestsRequestRedirectResponse `json:"redirectResponse"`
	JSON                 scanGetResponseDataRequestsRequestJSON             `json:"-"`
}

// scanGetResponseDataRequestsRequestJSON contains the JSON metadata for the struct
// [ScanGetResponseDataRequestsRequest]
type scanGetResponseDataRequestsRequestJSON struct {
	DocumentURL          apijson.Field
	HasUserGesture       apijson.Field
	Initiator            apijson.Field
	RedirectHasExtraInfo apijson.Field
	Request              apijson.Field
	RequestID            apijson.Field
	Type                 apijson.Field
	WallTime             apijson.Field
	FrameID              apijson.Field
	LoaderID             apijson.Field
	PrimaryRequest       apijson.Field
	RedirectResponse     apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ScanGetResponseDataRequestsRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseDataRequestsRequestJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseDataRequestsRequestInitiator struct {
	Host string                                          `json:"host,required"`
	Type string                                          `json:"type,required"`
	URL  string                                          `json:"url,required"`
	JSON scanGetResponseDataRequestsRequestInitiatorJSON `json:"-"`
}

// scanGetResponseDataRequestsRequestInitiatorJSON contains the JSON metadata for
// the struct [ScanGetResponseDataRequestsRequestInitiator]
type scanGetResponseDataRequestsRequestInitiatorJSON struct {
	Host        apijson.Field
	Type        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseDataRequestsRequestInitiator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseDataRequestsRequestInitiatorJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseDataRequestsRequestRequest struct {
	InitialPriority  string                                        `json:"initialPriority,required"`
	IsSameSite       bool                                          `json:"isSameSite,required"`
	Method           string                                        `json:"method,required"`
	MixedContentType string                                        `json:"mixedContentType,required"`
	ReferrerPolicy   string                                        `json:"referrerPolicy,required"`
	URL              string                                        `json:"url,required"`
	Headers          interface{}                                   `json:"headers"`
	JSON             scanGetResponseDataRequestsRequestRequestJSON `json:"-"`
}

// scanGetResponseDataRequestsRequestRequestJSON contains the JSON metadata for the
// struct [ScanGetResponseDataRequestsRequestRequest]
type scanGetResponseDataRequestsRequestRequestJSON struct {
	InitialPriority  apijson.Field
	IsSameSite       apijson.Field
	Method           apijson.Field
	MixedContentType apijson.Field
	ReferrerPolicy   apijson.Field
	URL              apijson.Field
	Headers          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScanGetResponseDataRequestsRequestRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseDataRequestsRequestRequestJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseDataRequestsRequestRedirectResponse struct {
	Charset         string                                                             `json:"charset,required"`
	MimeType        string                                                             `json:"mimeType,required"`
	Protocol        string                                                             `json:"protocol,required"`
	RemoteIPAddress string                                                             `json:"remoteIPAddress,required"`
	RemotePort      float64                                                            `json:"remotePort,required"`
	SecurityHeaders []ScanGetResponseDataRequestsRequestRedirectResponseSecurityHeader `json:"securityHeaders,required"`
	SecurityState   string                                                             `json:"securityState,required"`
	Status          float64                                                            `json:"status,required"`
	StatusText      string                                                             `json:"statusText,required"`
	URL             string                                                             `json:"url,required"`
	Headers         interface{}                                                        `json:"headers"`
	JSON            scanGetResponseDataRequestsRequestRedirectResponseJSON             `json:"-"`
}

// scanGetResponseDataRequestsRequestRedirectResponseJSON contains the JSON
// metadata for the struct [ScanGetResponseDataRequestsRequestRedirectResponse]
type scanGetResponseDataRequestsRequestRedirectResponseJSON struct {
	Charset         apijson.Field
	MimeType        apijson.Field
	Protocol        apijson.Field
	RemoteIPAddress apijson.Field
	RemotePort      apijson.Field
	SecurityHeaders apijson.Field
	SecurityState   apijson.Field
	Status          apijson.Field
	StatusText      apijson.Field
	URL             apijson.Field
	Headers         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ScanGetResponseDataRequestsRequestRedirectResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseDataRequestsRequestRedirectResponseJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseDataRequestsRequestRedirectResponseSecurityHeader struct {
	Name  string                                                               `json:"name,required"`
	Value string                                                               `json:"value,required"`
	JSON  scanGetResponseDataRequestsRequestRedirectResponseSecurityHeaderJSON `json:"-"`
}

// scanGetResponseDataRequestsRequestRedirectResponseSecurityHeaderJSON contains
// the JSON metadata for the struct
// [ScanGetResponseDataRequestsRequestRedirectResponseSecurityHeader]
type scanGetResponseDataRequestsRequestRedirectResponseSecurityHeaderJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseDataRequestsRequestRedirectResponseSecurityHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseDataRequestsRequestRedirectResponseSecurityHeaderJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseDataRequestsResponse struct {
	ASN               ScanGetResponseDataRequestsResponseASN      `json:"asn,required"`
	DataLength        float64                                     `json:"dataLength,required"`
	EncodedDataLength float64                                     `json:"encodedDataLength,required"`
	Geoip             ScanGetResponseDataRequestsResponseGeoip    `json:"geoip,required"`
	HasExtraInfo      bool                                        `json:"hasExtraInfo,required"`
	RequestID         string                                      `json:"requestId,required"`
	Response          ScanGetResponseDataRequestsResponseResponse `json:"response,required"`
	Size              float64                                     `json:"size,required"`
	Type              string                                      `json:"type,required"`
	ContentAvailable  bool                                        `json:"contentAvailable"`
	Hash              string                                      `json:"hash"`
	JSON              scanGetResponseDataRequestsResponseJSON     `json:"-"`
}

// scanGetResponseDataRequestsResponseJSON contains the JSON metadata for the
// struct [ScanGetResponseDataRequestsResponse]
type scanGetResponseDataRequestsResponseJSON struct {
	ASN               apijson.Field
	DataLength        apijson.Field
	EncodedDataLength apijson.Field
	Geoip             apijson.Field
	HasExtraInfo      apijson.Field
	RequestID         apijson.Field
	Response          apijson.Field
	Size              apijson.Field
	Type              apijson.Field
	ContentAvailable  apijson.Field
	Hash              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ScanGetResponseDataRequestsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseDataRequestsResponseJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseDataRequestsResponseASN struct {
	ASN         string                                     `json:"asn,required"`
	Country     string                                     `json:"country,required"`
	Description string                                     `json:"description,required"`
	IP          string                                     `json:"ip,required"`
	Name        string                                     `json:"name,required"`
	Org         string                                     `json:"org,required"`
	JSON        scanGetResponseDataRequestsResponseASNJSON `json:"-"`
}

// scanGetResponseDataRequestsResponseASNJSON contains the JSON metadata for the
// struct [ScanGetResponseDataRequestsResponseASN]
type scanGetResponseDataRequestsResponseASNJSON struct {
	ASN         apijson.Field
	Country     apijson.Field
	Description apijson.Field
	IP          apijson.Field
	Name        apijson.Field
	Org         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseDataRequestsResponseASN) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseDataRequestsResponseASNJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseDataRequestsResponseGeoip struct {
	City        string                                       `json:"city,required"`
	Country     string                                       `json:"country,required"`
	CountryName string                                       `json:"country_name,required"`
	GeonameID   string                                       `json:"geonameId,required"`
	Ll          []float64                                    `json:"ll,required"`
	Region      string                                       `json:"region,required"`
	JSON        scanGetResponseDataRequestsResponseGeoipJSON `json:"-"`
}

// scanGetResponseDataRequestsResponseGeoipJSON contains the JSON metadata for the
// struct [ScanGetResponseDataRequestsResponseGeoip]
type scanGetResponseDataRequestsResponseGeoipJSON struct {
	City        apijson.Field
	Country     apijson.Field
	CountryName apijson.Field
	GeonameID   apijson.Field
	Ll          apijson.Field
	Region      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseDataRequestsResponseGeoip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseDataRequestsResponseGeoipJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseDataRequestsResponseResponse struct {
	Charset         string                                                      `json:"charset,required"`
	MimeType        string                                                      `json:"mimeType,required"`
	Protocol        string                                                      `json:"protocol,required"`
	RemoteIPAddress string                                                      `json:"remoteIPAddress,required"`
	RemotePort      float64                                                     `json:"remotePort,required"`
	SecurityDetails ScanGetResponseDataRequestsResponseResponseSecurityDetails  `json:"securityDetails,required"`
	SecurityHeaders []ScanGetResponseDataRequestsResponseResponseSecurityHeader `json:"securityHeaders,required"`
	SecurityState   string                                                      `json:"securityState,required"`
	Status          float64                                                     `json:"status,required"`
	StatusText      string                                                      `json:"statusText,required"`
	URL             string                                                      `json:"url,required"`
	Headers         interface{}                                                 `json:"headers"`
	JSON            scanGetResponseDataRequestsResponseResponseJSON             `json:"-"`
}

// scanGetResponseDataRequestsResponseResponseJSON contains the JSON metadata for
// the struct [ScanGetResponseDataRequestsResponseResponse]
type scanGetResponseDataRequestsResponseResponseJSON struct {
	Charset         apijson.Field
	MimeType        apijson.Field
	Protocol        apijson.Field
	RemoteIPAddress apijson.Field
	RemotePort      apijson.Field
	SecurityDetails apijson.Field
	SecurityHeaders apijson.Field
	SecurityState   apijson.Field
	Status          apijson.Field
	StatusText      apijson.Field
	URL             apijson.Field
	Headers         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ScanGetResponseDataRequestsResponseResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseDataRequestsResponseResponseJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseDataRequestsResponseResponseSecurityDetails struct {
	CertificateID                     float64                                                        `json:"certificateId,required"`
	CertificateTransparencyCompliance string                                                         `json:"certificateTransparencyCompliance,required"`
	Cipher                            string                                                         `json:"cipher,required"`
	EncryptedClientHello              bool                                                           `json:"encryptedClientHello,required"`
	Issuer                            string                                                         `json:"issuer,required"`
	KeyExchange                       string                                                         `json:"keyExchange,required"`
	KeyExchangeGroup                  string                                                         `json:"keyExchangeGroup,required"`
	Protocol                          string                                                         `json:"protocol,required"`
	SanList                           []string                                                       `json:"sanList,required"`
	ServerSignatureAlgorithm          float64                                                        `json:"serverSignatureAlgorithm,required"`
	SubjectName                       string                                                         `json:"subjectName,required"`
	ValidFrom                         float64                                                        `json:"validFrom,required"`
	ValidTo                           float64                                                        `json:"validTo,required"`
	JSON                              scanGetResponseDataRequestsResponseResponseSecurityDetailsJSON `json:"-"`
}

// scanGetResponseDataRequestsResponseResponseSecurityDetailsJSON contains the JSON
// metadata for the struct
// [ScanGetResponseDataRequestsResponseResponseSecurityDetails]
type scanGetResponseDataRequestsResponseResponseSecurityDetailsJSON struct {
	CertificateID                     apijson.Field
	CertificateTransparencyCompliance apijson.Field
	Cipher                            apijson.Field
	EncryptedClientHello              apijson.Field
	Issuer                            apijson.Field
	KeyExchange                       apijson.Field
	KeyExchangeGroup                  apijson.Field
	Protocol                          apijson.Field
	SanList                           apijson.Field
	ServerSignatureAlgorithm          apijson.Field
	SubjectName                       apijson.Field
	ValidFrom                         apijson.Field
	ValidTo                           apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *ScanGetResponseDataRequestsResponseResponseSecurityDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseDataRequestsResponseResponseSecurityDetailsJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseDataRequestsResponseResponseSecurityHeader struct {
	Name  string                                                        `json:"name,required"`
	Value string                                                        `json:"value,required"`
	JSON  scanGetResponseDataRequestsResponseResponseSecurityHeaderJSON `json:"-"`
}

// scanGetResponseDataRequestsResponseResponseSecurityHeaderJSON contains the JSON
// metadata for the struct
// [ScanGetResponseDataRequestsResponseResponseSecurityHeader]
type scanGetResponseDataRequestsResponseResponseSecurityHeaderJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseDataRequestsResponseResponseSecurityHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseDataRequestsResponseResponseSecurityHeaderJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseLists struct {
	ASNs         []string                          `json:"asns,required"`
	Certificates []ScanGetResponseListsCertificate `json:"certificates,required"`
	Continents   []string                          `json:"continents,required"`
	Countries    []string                          `json:"countries,required"`
	Domains      []string                          `json:"domains,required"`
	Hashes       []string                          `json:"hashes,required"`
	IPs          []string                          `json:"ips,required"`
	LinkDomains  []string                          `json:"linkDomains,required"`
	Servers      []string                          `json:"servers,required"`
	URLs         []string                          `json:"urls,required"`
	JSON         scanGetResponseListsJSON          `json:"-"`
}

// scanGetResponseListsJSON contains the JSON metadata for the struct
// [ScanGetResponseLists]
type scanGetResponseListsJSON struct {
	ASNs         apijson.Field
	Certificates apijson.Field
	Continents   apijson.Field
	Countries    apijson.Field
	Domains      apijson.Field
	Hashes       apijson.Field
	IPs          apijson.Field
	LinkDomains  apijson.Field
	Servers      apijson.Field
	URLs         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ScanGetResponseLists) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseListsJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseListsCertificate struct {
	Issuer      string                              `json:"issuer,required"`
	SubjectName string                              `json:"subjectName,required"`
	ValidFrom   float64                             `json:"validFrom,required"`
	ValidTo     float64                             `json:"validTo,required"`
	JSON        scanGetResponseListsCertificateJSON `json:"-"`
}

// scanGetResponseListsCertificateJSON contains the JSON metadata for the struct
// [ScanGetResponseListsCertificate]
type scanGetResponseListsCertificateJSON struct {
	Issuer      apijson.Field
	SubjectName apijson.Field
	ValidFrom   apijson.Field
	ValidTo     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseListsCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseListsCertificateJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseMeta struct {
	Processors ScanGetResponseMetaProcessors `json:"processors,required"`
	JSON       scanGetResponseMetaJSON       `json:"-"`
}

// scanGetResponseMetaJSON contains the JSON metadata for the struct
// [ScanGetResponseMeta]
type scanGetResponseMetaJSON struct {
	Processors  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseMetaProcessors struct {
	ASN              ScanGetResponseMetaProcessorsASN              `json:"asn,required"`
	DNS              ScanGetResponseMetaProcessorsDNS              `json:"dns,required"`
	DomainCategories ScanGetResponseMetaProcessorsDomainCategories `json:"domainCategories,required"`
	Geoip            ScanGetResponseMetaProcessorsGeoip            `json:"geoip,required"`
	Phishing         ScanGetResponseMetaProcessorsPhishing         `json:"phishing,required"`
	RadarRank        ScanGetResponseMetaProcessorsRadarRank        `json:"radarRank,required"`
	Wappa            ScanGetResponseMetaProcessorsWappa            `json:"wappa,required"`
	URLCategories    ScanGetResponseMetaProcessorsURLCategories    `json:"urlCategories"`
	JSON             scanGetResponseMetaProcessorsJSON             `json:"-"`
}

// scanGetResponseMetaProcessorsJSON contains the JSON metadata for the struct
// [ScanGetResponseMetaProcessors]
type scanGetResponseMetaProcessorsJSON struct {
	ASN              apijson.Field
	DNS              apijson.Field
	DomainCategories apijson.Field
	Geoip            apijson.Field
	Phishing         apijson.Field
	RadarRank        apijson.Field
	Wappa            apijson.Field
	URLCategories    apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScanGetResponseMetaProcessors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseMetaProcessorsJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseMetaProcessorsASN struct {
	Data []ScanGetResponseMetaProcessorsASNData `json:"data,required"`
	JSON scanGetResponseMetaProcessorsASNJSON   `json:"-"`
}

// scanGetResponseMetaProcessorsASNJSON contains the JSON metadata for the struct
// [ScanGetResponseMetaProcessorsASN]
type scanGetResponseMetaProcessorsASNJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseMetaProcessorsASN) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseMetaProcessorsASNJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseMetaProcessorsASNData struct {
	ASN         string                                   `json:"asn,required"`
	Country     string                                   `json:"country,required"`
	Description string                                   `json:"description,required"`
	IP          string                                   `json:"ip,required"`
	Name        string                                   `json:"name,required"`
	JSON        scanGetResponseMetaProcessorsASNDataJSON `json:"-"`
}

// scanGetResponseMetaProcessorsASNDataJSON contains the JSON metadata for the
// struct [ScanGetResponseMetaProcessorsASNData]
type scanGetResponseMetaProcessorsASNDataJSON struct {
	ASN         apijson.Field
	Country     apijson.Field
	Description apijson.Field
	IP          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseMetaProcessorsASNData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseMetaProcessorsASNDataJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseMetaProcessorsDNS struct {
	Data []ScanGetResponseMetaProcessorsDNSData `json:"data,required"`
	JSON scanGetResponseMetaProcessorsDNSJSON   `json:"-"`
}

// scanGetResponseMetaProcessorsDNSJSON contains the JSON metadata for the struct
// [ScanGetResponseMetaProcessorsDNS]
type scanGetResponseMetaProcessorsDNSJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseMetaProcessorsDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseMetaProcessorsDNSJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseMetaProcessorsDNSData struct {
	Address     string                                   `json:"address,required"`
	DNSSECValid bool                                     `json:"dnssec_valid,required"`
	Name        string                                   `json:"name,required"`
	Type        string                                   `json:"type,required"`
	JSON        scanGetResponseMetaProcessorsDNSDataJSON `json:"-"`
}

// scanGetResponseMetaProcessorsDNSDataJSON contains the JSON metadata for the
// struct [ScanGetResponseMetaProcessorsDNSData]
type scanGetResponseMetaProcessorsDNSDataJSON struct {
	Address     apijson.Field
	DNSSECValid apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseMetaProcessorsDNSData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseMetaProcessorsDNSDataJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseMetaProcessorsDomainCategories struct {
	Data []ScanGetResponseMetaProcessorsDomainCategoriesData `json:"data,required"`
	JSON scanGetResponseMetaProcessorsDomainCategoriesJSON   `json:"-"`
}

// scanGetResponseMetaProcessorsDomainCategoriesJSON contains the JSON metadata for
// the struct [ScanGetResponseMetaProcessorsDomainCategories]
type scanGetResponseMetaProcessorsDomainCategoriesJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseMetaProcessorsDomainCategories) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseMetaProcessorsDomainCategoriesJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseMetaProcessorsDomainCategoriesData struct {
	Inherited interface{}                                           `json:"inherited,required"`
	IsPrimary bool                                                  `json:"isPrimary,required"`
	Name      string                                                `json:"name,required"`
	JSON      scanGetResponseMetaProcessorsDomainCategoriesDataJSON `json:"-"`
}

// scanGetResponseMetaProcessorsDomainCategoriesDataJSON contains the JSON metadata
// for the struct [ScanGetResponseMetaProcessorsDomainCategoriesData]
type scanGetResponseMetaProcessorsDomainCategoriesDataJSON struct {
	Inherited   apijson.Field
	IsPrimary   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseMetaProcessorsDomainCategoriesData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseMetaProcessorsDomainCategoriesDataJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseMetaProcessorsGeoip struct {
	Data []ScanGetResponseMetaProcessorsGeoipData `json:"data,required"`
	JSON scanGetResponseMetaProcessorsGeoipJSON   `json:"-"`
}

// scanGetResponseMetaProcessorsGeoipJSON contains the JSON metadata for the struct
// [ScanGetResponseMetaProcessorsGeoip]
type scanGetResponseMetaProcessorsGeoipJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseMetaProcessorsGeoip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseMetaProcessorsGeoipJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseMetaProcessorsGeoipData struct {
	Geoip ScanGetResponseMetaProcessorsGeoipDataGeoip `json:"geoip,required"`
	IP    string                                      `json:"ip,required"`
	JSON  scanGetResponseMetaProcessorsGeoipDataJSON  `json:"-"`
}

// scanGetResponseMetaProcessorsGeoipDataJSON contains the JSON metadata for the
// struct [ScanGetResponseMetaProcessorsGeoipData]
type scanGetResponseMetaProcessorsGeoipDataJSON struct {
	Geoip       apijson.Field
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseMetaProcessorsGeoipData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseMetaProcessorsGeoipDataJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseMetaProcessorsGeoipDataGeoip struct {
	City        string                                          `json:"city,required"`
	Country     string                                          `json:"country,required"`
	CountryName string                                          `json:"country_name,required"`
	Ll          []float64                                       `json:"ll,required"`
	Region      string                                          `json:"region,required"`
	JSON        scanGetResponseMetaProcessorsGeoipDataGeoipJSON `json:"-"`
}

// scanGetResponseMetaProcessorsGeoipDataGeoipJSON contains the JSON metadata for
// the struct [ScanGetResponseMetaProcessorsGeoipDataGeoip]
type scanGetResponseMetaProcessorsGeoipDataGeoipJSON struct {
	City        apijson.Field
	Country     apijson.Field
	CountryName apijson.Field
	Ll          apijson.Field
	Region      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseMetaProcessorsGeoipDataGeoip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseMetaProcessorsGeoipDataGeoipJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseMetaProcessorsPhishing struct {
	Data []string                                  `json:"data,required"`
	JSON scanGetResponseMetaProcessorsPhishingJSON `json:"-"`
}

// scanGetResponseMetaProcessorsPhishingJSON contains the JSON metadata for the
// struct [ScanGetResponseMetaProcessorsPhishing]
type scanGetResponseMetaProcessorsPhishingJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseMetaProcessorsPhishing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseMetaProcessorsPhishingJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseMetaProcessorsRadarRank struct {
	Data []ScanGetResponseMetaProcessorsRadarRankData `json:"data,required"`
	JSON scanGetResponseMetaProcessorsRadarRankJSON   `json:"-"`
}

// scanGetResponseMetaProcessorsRadarRankJSON contains the JSON metadata for the
// struct [ScanGetResponseMetaProcessorsRadarRank]
type scanGetResponseMetaProcessorsRadarRankJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseMetaProcessorsRadarRank) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseMetaProcessorsRadarRankJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseMetaProcessorsRadarRankData struct {
	Bucket   string                                         `json:"bucket,required"`
	Hostname string                                         `json:"hostname,required"`
	Rank     float64                                        `json:"rank"`
	JSON     scanGetResponseMetaProcessorsRadarRankDataJSON `json:"-"`
}

// scanGetResponseMetaProcessorsRadarRankDataJSON contains the JSON metadata for
// the struct [ScanGetResponseMetaProcessorsRadarRankData]
type scanGetResponseMetaProcessorsRadarRankDataJSON struct {
	Bucket      apijson.Field
	Hostname    apijson.Field
	Rank        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseMetaProcessorsRadarRankData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseMetaProcessorsRadarRankDataJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseMetaProcessorsWappa struct {
	Data []ScanGetResponseMetaProcessorsWappaData `json:"data,required"`
	JSON scanGetResponseMetaProcessorsWappaJSON   `json:"-"`
}

// scanGetResponseMetaProcessorsWappaJSON contains the JSON metadata for the struct
// [ScanGetResponseMetaProcessorsWappa]
type scanGetResponseMetaProcessorsWappaJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseMetaProcessorsWappa) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseMetaProcessorsWappaJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseMetaProcessorsWappaData struct {
	App             string                                             `json:"app,required"`
	Categories      []ScanGetResponseMetaProcessorsWappaDataCategory   `json:"categories,required"`
	Confidence      []ScanGetResponseMetaProcessorsWappaDataConfidence `json:"confidence,required"`
	ConfidenceTotal float64                                            `json:"confidenceTotal,required"`
	Icon            string                                             `json:"icon,required"`
	Website         string                                             `json:"website,required"`
	JSON            scanGetResponseMetaProcessorsWappaDataJSON         `json:"-"`
}

// scanGetResponseMetaProcessorsWappaDataJSON contains the JSON metadata for the
// struct [ScanGetResponseMetaProcessorsWappaData]
type scanGetResponseMetaProcessorsWappaDataJSON struct {
	App             apijson.Field
	Categories      apijson.Field
	Confidence      apijson.Field
	ConfidenceTotal apijson.Field
	Icon            apijson.Field
	Website         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ScanGetResponseMetaProcessorsWappaData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseMetaProcessorsWappaDataJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseMetaProcessorsWappaDataCategory struct {
	Name     string                                             `json:"name,required"`
	Priority float64                                            `json:"priority,required"`
	JSON     scanGetResponseMetaProcessorsWappaDataCategoryJSON `json:"-"`
}

// scanGetResponseMetaProcessorsWappaDataCategoryJSON contains the JSON metadata
// for the struct [ScanGetResponseMetaProcessorsWappaDataCategory]
type scanGetResponseMetaProcessorsWappaDataCategoryJSON struct {
	Name        apijson.Field
	Priority    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseMetaProcessorsWappaDataCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseMetaProcessorsWappaDataCategoryJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseMetaProcessorsWappaDataConfidence struct {
	Confidence  float64                                              `json:"confidence,required"`
	Name        string                                               `json:"name,required"`
	Pattern     string                                               `json:"pattern,required"`
	PatternType string                                               `json:"patternType,required"`
	JSON        scanGetResponseMetaProcessorsWappaDataConfidenceJSON `json:"-"`
}

// scanGetResponseMetaProcessorsWappaDataConfidenceJSON contains the JSON metadata
// for the struct [ScanGetResponseMetaProcessorsWappaDataConfidence]
type scanGetResponseMetaProcessorsWappaDataConfidenceJSON struct {
	Confidence  apijson.Field
	Name        apijson.Field
	Pattern     apijson.Field
	PatternType apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseMetaProcessorsWappaDataConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseMetaProcessorsWappaDataConfidenceJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseMetaProcessorsURLCategories struct {
	Data []ScanGetResponseMetaProcessorsURLCategoriesData `json:"data,required"`
	JSON scanGetResponseMetaProcessorsURLCategoriesJSON   `json:"-"`
}

// scanGetResponseMetaProcessorsURLCategoriesJSON contains the JSON metadata for
// the struct [ScanGetResponseMetaProcessorsURLCategories]
type scanGetResponseMetaProcessorsURLCategoriesJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseMetaProcessorsURLCategories) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseMetaProcessorsURLCategoriesJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseMetaProcessorsURLCategoriesData struct {
	Content   []ScanGetResponseMetaProcessorsURLCategoriesDataContent `json:"content,required"`
	Inherited ScanGetResponseMetaProcessorsURLCategoriesDataInherited `json:"inherited,required"`
	Name      string                                                  `json:"name,required"`
	Risks     []ScanGetResponseMetaProcessorsURLCategoriesDataRisk    `json:"risks,required"`
	JSON      scanGetResponseMetaProcessorsURLCategoriesDataJSON      `json:"-"`
}

// scanGetResponseMetaProcessorsURLCategoriesDataJSON contains the JSON metadata
// for the struct [ScanGetResponseMetaProcessorsURLCategoriesData]
type scanGetResponseMetaProcessorsURLCategoriesDataJSON struct {
	Content     apijson.Field
	Inherited   apijson.Field
	Name        apijson.Field
	Risks       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseMetaProcessorsURLCategoriesData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseMetaProcessorsURLCategoriesDataJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseMetaProcessorsURLCategoriesDataContent struct {
	ID              float64                                                   `json:"id,required"`
	Name            string                                                    `json:"name,required"`
	SuperCategoryID float64                                                   `json:"super_category_id,required"`
	JSON            scanGetResponseMetaProcessorsURLCategoriesDataContentJSON `json:"-"`
}

// scanGetResponseMetaProcessorsURLCategoriesDataContentJSON contains the JSON
// metadata for the struct [ScanGetResponseMetaProcessorsURLCategoriesDataContent]
type scanGetResponseMetaProcessorsURLCategoriesDataContentJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ScanGetResponseMetaProcessorsURLCategoriesDataContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseMetaProcessorsURLCategoriesDataContentJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseMetaProcessorsURLCategoriesDataInherited struct {
	Content []ScanGetResponseMetaProcessorsURLCategoriesDataInheritedContent `json:"content,required"`
	From    string                                                           `json:"from,required"`
	Risks   []ScanGetResponseMetaProcessorsURLCategoriesDataInheritedRisk    `json:"risks,required"`
	JSON    scanGetResponseMetaProcessorsURLCategoriesDataInheritedJSON      `json:"-"`
}

// scanGetResponseMetaProcessorsURLCategoriesDataInheritedJSON contains the JSON
// metadata for the struct
// [ScanGetResponseMetaProcessorsURLCategoriesDataInherited]
type scanGetResponseMetaProcessorsURLCategoriesDataInheritedJSON struct {
	Content     apijson.Field
	From        apijson.Field
	Risks       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseMetaProcessorsURLCategoriesDataInherited) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseMetaProcessorsURLCategoriesDataInheritedJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseMetaProcessorsURLCategoriesDataInheritedContent struct {
	ID              float64                                                            `json:"id,required"`
	Name            string                                                             `json:"name,required"`
	SuperCategoryID float64                                                            `json:"super_category_id,required"`
	JSON            scanGetResponseMetaProcessorsURLCategoriesDataInheritedContentJSON `json:"-"`
}

// scanGetResponseMetaProcessorsURLCategoriesDataInheritedContentJSON contains the
// JSON metadata for the struct
// [ScanGetResponseMetaProcessorsURLCategoriesDataInheritedContent]
type scanGetResponseMetaProcessorsURLCategoriesDataInheritedContentJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ScanGetResponseMetaProcessorsURLCategoriesDataInheritedContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseMetaProcessorsURLCategoriesDataInheritedContentJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseMetaProcessorsURLCategoriesDataInheritedRisk struct {
	ID              float64                                                         `json:"id,required"`
	Name            string                                                          `json:"name,required"`
	SuperCategoryID float64                                                         `json:"super_category_id,required"`
	JSON            scanGetResponseMetaProcessorsURLCategoriesDataInheritedRiskJSON `json:"-"`
}

// scanGetResponseMetaProcessorsURLCategoriesDataInheritedRiskJSON contains the
// JSON metadata for the struct
// [ScanGetResponseMetaProcessorsURLCategoriesDataInheritedRisk]
type scanGetResponseMetaProcessorsURLCategoriesDataInheritedRiskJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ScanGetResponseMetaProcessorsURLCategoriesDataInheritedRisk) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseMetaProcessorsURLCategoriesDataInheritedRiskJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseMetaProcessorsURLCategoriesDataRisk struct {
	ID              float64                                                `json:"id,required"`
	Name            string                                                 `json:"name,required"`
	SuperCategoryID float64                                                `json:"super_category_id,required"`
	JSON            scanGetResponseMetaProcessorsURLCategoriesDataRiskJSON `json:"-"`
}

// scanGetResponseMetaProcessorsURLCategoriesDataRiskJSON contains the JSON
// metadata for the struct [ScanGetResponseMetaProcessorsURLCategoriesDataRisk]
type scanGetResponseMetaProcessorsURLCategoriesDataRiskJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ScanGetResponseMetaProcessorsURLCategoriesDataRisk) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseMetaProcessorsURLCategoriesDataRiskJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponsePage struct {
	ApexDomain   string                        `json:"apexDomain,required"`
	ASN          string                        `json:"asn,required"`
	Asnname      string                        `json:"asnname,required"`
	City         string                        `json:"city,required"`
	Country      string                        `json:"country,required"`
	Domain       string                        `json:"domain,required"`
	IP           string                        `json:"ip,required"`
	MimeType     string                        `json:"mimeType,required"`
	Server       string                        `json:"server,required"`
	Status       string                        `json:"status,required"`
	Title        string                        `json:"title,required"`
	TLSAgeDays   float64                       `json:"tlsAgeDays,required"`
	TLSIssuer    string                        `json:"tlsIssuer,required"`
	TLSValidDays float64                       `json:"tlsValidDays,required"`
	TLSValidFrom string                        `json:"tlsValidFrom,required"`
	URL          string                        `json:"url,required"`
	Screenshot   ScanGetResponsePageScreenshot `json:"screenshot"`
	JSON         scanGetResponsePageJSON       `json:"-"`
}

// scanGetResponsePageJSON contains the JSON metadata for the struct
// [ScanGetResponsePage]
type scanGetResponsePageJSON struct {
	ApexDomain   apijson.Field
	ASN          apijson.Field
	Asnname      apijson.Field
	City         apijson.Field
	Country      apijson.Field
	Domain       apijson.Field
	IP           apijson.Field
	MimeType     apijson.Field
	Server       apijson.Field
	Status       apijson.Field
	Title        apijson.Field
	TLSAgeDays   apijson.Field
	TLSIssuer    apijson.Field
	TLSValidDays apijson.Field
	TLSValidFrom apijson.Field
	URL          apijson.Field
	Screenshot   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ScanGetResponsePage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponsePageJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponsePageScreenshot struct {
	Dhash   string                            `json:"dhash,required"`
	Mm3Hash float64                           `json:"mm3Hash,required"`
	Name    string                            `json:"name,required"`
	Phash   string                            `json:"phash,required"`
	JSON    scanGetResponsePageScreenshotJSON `json:"-"`
}

// scanGetResponsePageScreenshotJSON contains the JSON metadata for the struct
// [ScanGetResponsePageScreenshot]
type scanGetResponsePageScreenshotJSON struct {
	Dhash       apijson.Field
	Mm3Hash     apijson.Field
	Name        apijson.Field
	Phash       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponsePageScreenshot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponsePageScreenshotJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseScanner struct {
	Colo    string                     `json:"colo,required"`
	Country string                     `json:"country,required"`
	JSON    scanGetResponseScannerJSON `json:"-"`
}

// scanGetResponseScannerJSON contains the JSON metadata for the struct
// [ScanGetResponseScanner]
type scanGetResponseScannerJSON struct {
	Colo        apijson.Field
	Country     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseScanner) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseScannerJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseStats struct {
	DomainStats      []ScanGetResponseStatsDomainStat   `json:"domainStats,required"`
	IPStats          []ScanGetResponseStatsIPStat       `json:"ipStats,required"`
	IPv6Percentage   float64                            `json:"IPv6Percentage,required"`
	Malicious        float64                            `json:"malicious,required"`
	ProtocolStats    []ScanGetResponseStatsProtocolStat `json:"protocolStats,required"`
	ResourceStats    []ScanGetResponseStatsResourceStat `json:"resourceStats,required"`
	SecurePercentage float64                            `json:"securePercentage,required"`
	SecureRequests   float64                            `json:"secureRequests,required"`
	ServerStats      []ScanGetResponseStatsServerStat   `json:"serverStats,required"`
	TLSStats         []ScanGetResponseStatsTLSStat      `json:"tlsStats,required"`
	TotalLinks       float64                            `json:"totalLinks,required"`
	UniqASNs         float64                            `json:"uniqASNs,required"`
	UniqCountries    float64                            `json:"uniqCountries,required"`
	JSON             scanGetResponseStatsJSON           `json:"-"`
}

// scanGetResponseStatsJSON contains the JSON metadata for the struct
// [ScanGetResponseStats]
type scanGetResponseStatsJSON struct {
	DomainStats      apijson.Field
	IPStats          apijson.Field
	IPv6Percentage   apijson.Field
	Malicious        apijson.Field
	ProtocolStats    apijson.Field
	ResourceStats    apijson.Field
	SecurePercentage apijson.Field
	SecureRequests   apijson.Field
	ServerStats      apijson.Field
	TLSStats         apijson.Field
	TotalLinks       apijson.Field
	UniqASNs         apijson.Field
	UniqCountries    apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScanGetResponseStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseStatsJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseStatsDomainStat struct {
	Count       float64                            `json:"count,required"`
	Countries   []string                           `json:"countries,required"`
	Domain      string                             `json:"domain,required"`
	EncodedSize float64                            `json:"encodedSize,required"`
	Index       float64                            `json:"index,required"`
	Initiators  []string                           `json:"initiators,required"`
	IPs         []string                           `json:"ips,required"`
	Redirects   float64                            `json:"redirects,required"`
	Size        float64                            `json:"size,required"`
	JSON        scanGetResponseStatsDomainStatJSON `json:"-"`
}

// scanGetResponseStatsDomainStatJSON contains the JSON metadata for the struct
// [ScanGetResponseStatsDomainStat]
type scanGetResponseStatsDomainStatJSON struct {
	Count       apijson.Field
	Countries   apijson.Field
	Domain      apijson.Field
	EncodedSize apijson.Field
	Index       apijson.Field
	Initiators  apijson.Field
	IPs         apijson.Field
	Redirects   apijson.Field
	Size        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseStatsDomainStat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseStatsDomainStatJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseStatsIPStat struct {
	ASN         ScanGetResponseStatsIPStatsASN   `json:"asn,required"`
	Countries   []string                         `json:"countries,required"`
	Domains     []string                         `json:"domains,required"`
	EncodedSize float64                          `json:"encodedSize,required"`
	Geoip       ScanGetResponseStatsIPStatsGeoip `json:"geoip,required"`
	Index       float64                          `json:"index,required"`
	IP          string                           `json:"ip,required"`
	IPV6        bool                             `json:"ipv6,required"`
	Redirects   float64                          `json:"redirects,required"`
	Requests    float64                          `json:"requests,required"`
	Size        float64                          `json:"size,required"`
	Count       float64                          `json:"count"`
	JSON        scanGetResponseStatsIPStatJSON   `json:"-"`
}

// scanGetResponseStatsIPStatJSON contains the JSON metadata for the struct
// [ScanGetResponseStatsIPStat]
type scanGetResponseStatsIPStatJSON struct {
	ASN         apijson.Field
	Countries   apijson.Field
	Domains     apijson.Field
	EncodedSize apijson.Field
	Geoip       apijson.Field
	Index       apijson.Field
	IP          apijson.Field
	IPV6        apijson.Field
	Redirects   apijson.Field
	Requests    apijson.Field
	Size        apijson.Field
	Count       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseStatsIPStat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseStatsIPStatJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseStatsIPStatsASN struct {
	ASN         string                             `json:"asn,required"`
	Country     string                             `json:"country,required"`
	Description string                             `json:"description,required"`
	IP          string                             `json:"ip,required"`
	Name        string                             `json:"name,required"`
	Org         string                             `json:"org,required"`
	JSON        scanGetResponseStatsIPStatsASNJSON `json:"-"`
}

// scanGetResponseStatsIPStatsASNJSON contains the JSON metadata for the struct
// [ScanGetResponseStatsIPStatsASN]
type scanGetResponseStatsIPStatsASNJSON struct {
	ASN         apijson.Field
	Country     apijson.Field
	Description apijson.Field
	IP          apijson.Field
	Name        apijson.Field
	Org         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseStatsIPStatsASN) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseStatsIPStatsASNJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseStatsIPStatsGeoip struct {
	City        string                               `json:"city,required"`
	Country     string                               `json:"country,required"`
	CountryName string                               `json:"country_name,required"`
	Ll          []float64                            `json:"ll,required"`
	Region      string                               `json:"region,required"`
	JSON        scanGetResponseStatsIPStatsGeoipJSON `json:"-"`
}

// scanGetResponseStatsIPStatsGeoipJSON contains the JSON metadata for the struct
// [ScanGetResponseStatsIPStatsGeoip]
type scanGetResponseStatsIPStatsGeoipJSON struct {
	City        apijson.Field
	Country     apijson.Field
	CountryName apijson.Field
	Ll          apijson.Field
	Region      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseStatsIPStatsGeoip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseStatsIPStatsGeoipJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseStatsProtocolStat struct {
	Count       float64                              `json:"count,required"`
	Countries   []string                             `json:"countries,required"`
	EncodedSize float64                              `json:"encodedSize,required"`
	IPs         []string                             `json:"ips,required"`
	Protocol    string                               `json:"protocol,required"`
	Size        float64                              `json:"size,required"`
	JSON        scanGetResponseStatsProtocolStatJSON `json:"-"`
}

// scanGetResponseStatsProtocolStatJSON contains the JSON metadata for the struct
// [ScanGetResponseStatsProtocolStat]
type scanGetResponseStatsProtocolStatJSON struct {
	Count       apijson.Field
	Countries   apijson.Field
	EncodedSize apijson.Field
	IPs         apijson.Field
	Protocol    apijson.Field
	Size        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseStatsProtocolStat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseStatsProtocolStatJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseStatsResourceStat struct {
	Compression float64                              `json:"compression,required"`
	Count       float64                              `json:"count,required"`
	Countries   []string                             `json:"countries,required"`
	EncodedSize float64                              `json:"encodedSize,required"`
	IPs         []string                             `json:"ips,required"`
	Percentage  float64                              `json:"percentage,required"`
	Size        float64                              `json:"size,required"`
	Type        string                               `json:"type,required"`
	JSON        scanGetResponseStatsResourceStatJSON `json:"-"`
}

// scanGetResponseStatsResourceStatJSON contains the JSON metadata for the struct
// [ScanGetResponseStatsResourceStat]
type scanGetResponseStatsResourceStatJSON struct {
	Compression apijson.Field
	Count       apijson.Field
	Countries   apijson.Field
	EncodedSize apijson.Field
	IPs         apijson.Field
	Percentage  apijson.Field
	Size        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseStatsResourceStat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseStatsResourceStatJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseStatsServerStat struct {
	Count       float64                            `json:"count,required"`
	Countries   []string                           `json:"countries,required"`
	EncodedSize float64                            `json:"encodedSize,required"`
	IPs         []string                           `json:"ips,required"`
	Server      string                             `json:"server,required"`
	Size        float64                            `json:"size,required"`
	JSON        scanGetResponseStatsServerStatJSON `json:"-"`
}

// scanGetResponseStatsServerStatJSON contains the JSON metadata for the struct
// [ScanGetResponseStatsServerStat]
type scanGetResponseStatsServerStatJSON struct {
	Count       apijson.Field
	Countries   apijson.Field
	EncodedSize apijson.Field
	IPs         apijson.Field
	Server      apijson.Field
	Size        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseStatsServerStat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseStatsServerStatJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseStatsTLSStat struct {
	Count         float64                               `json:"count,required"`
	Countries     []string                              `json:"countries,required"`
	EncodedSize   float64                               `json:"encodedSize,required"`
	IPs           []string                              `json:"ips,required"`
	Protocols     ScanGetResponseStatsTLSStatsProtocols `json:"protocols,required"`
	SecurityState string                                `json:"securityState,required"`
	Size          float64                               `json:"size,required"`
	JSON          scanGetResponseStatsTLSStatJSON       `json:"-"`
}

// scanGetResponseStatsTLSStatJSON contains the JSON metadata for the struct
// [ScanGetResponseStatsTLSStat]
type scanGetResponseStatsTLSStatJSON struct {
	Count         apijson.Field
	Countries     apijson.Field
	EncodedSize   apijson.Field
	IPs           apijson.Field
	Protocols     apijson.Field
	SecurityState apijson.Field
	Size          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ScanGetResponseStatsTLSStat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseStatsTLSStatJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseStatsTLSStatsProtocols struct {
	TLS1_3Aes128Gcm float64                                   `json:"TLS 1.3 / AES_128_GCM,required"`
	JSON            scanGetResponseStatsTLSStatsProtocolsJSON `json:"-"`
}

// scanGetResponseStatsTLSStatsProtocolsJSON contains the JSON metadata for the
// struct [ScanGetResponseStatsTLSStatsProtocols]
type scanGetResponseStatsTLSStatsProtocolsJSON struct {
	TLS1_3Aes128Gcm apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ScanGetResponseStatsTLSStatsProtocols) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseStatsTLSStatsProtocolsJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseTask struct {
	ApexDomain    string                     `json:"apexDomain,required"`
	Domain        string                     `json:"domain,required"`
	DOMURL        string                     `json:"domURL,required"`
	Method        string                     `json:"method,required"`
	Options       ScanGetResponseTaskOptions `json:"options,required"`
	ReportURL     string                     `json:"reportURL,required"`
	ScreenshotURL string                     `json:"screenshotURL,required"`
	Source        string                     `json:"source,required"`
	Success       bool                       `json:"success,required"`
	Time          string                     `json:"time,required"`
	URL           string                     `json:"url,required"`
	UUID          string                     `json:"uuid,required"`
	Visibility    string                     `json:"visibility,required"`
	JSON          scanGetResponseTaskJSON    `json:"-"`
}

// scanGetResponseTaskJSON contains the JSON metadata for the struct
// [ScanGetResponseTask]
type scanGetResponseTaskJSON struct {
	ApexDomain    apijson.Field
	Domain        apijson.Field
	DOMURL        apijson.Field
	Method        apijson.Field
	Options       apijson.Field
	ReportURL     apijson.Field
	ScreenshotURL apijson.Field
	Source        apijson.Field
	Success       apijson.Field
	Time          apijson.Field
	URL           apijson.Field
	UUID          apijson.Field
	Visibility    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ScanGetResponseTask) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseTaskJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseTaskOptions struct {
	// Custom headers set.
	CustomHeaders          interface{}                    `json:"customHeaders"`
	ScreenshotsResolutions []string                       `json:"screenshotsResolutions"`
	JSON                   scanGetResponseTaskOptionsJSON `json:"-"`
}

// scanGetResponseTaskOptionsJSON contains the JSON metadata for the struct
// [ScanGetResponseTaskOptions]
type scanGetResponseTaskOptionsJSON struct {
	CustomHeaders          apijson.Field
	ScreenshotsResolutions apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ScanGetResponseTaskOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseTaskOptionsJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseVerdicts struct {
	Overall ScanGetResponseVerdictsOverall `json:"overall,required"`
	JSON    scanGetResponseVerdictsJSON    `json:"-"`
}

// scanGetResponseVerdictsJSON contains the JSON metadata for the struct
// [ScanGetResponseVerdicts]
type scanGetResponseVerdictsJSON struct {
	Overall     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseVerdicts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseVerdictsJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseVerdictsOverall struct {
	Categories  []string                           `json:"categories,required"`
	HasVerdicts bool                               `json:"hasVerdicts,required"`
	Malicious   bool                               `json:"malicious,required"`
	Tags        []string                           `json:"tags,required"`
	JSON        scanGetResponseVerdictsOverallJSON `json:"-"`
}

// scanGetResponseVerdictsOverallJSON contains the JSON metadata for the struct
// [ScanGetResponseVerdictsOverall]
type scanGetResponseVerdictsOverallJSON struct {
	Categories  apijson.Field
	HasVerdicts apijson.Field
	Malicious   apijson.Field
	Tags        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseVerdictsOverall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseVerdictsOverallJSON) RawJSON() string {
	return r.raw
}

type ScanHARResponse struct {
	Log  ScanHARResponseLog  `json:"log,required"`
	JSON scanHARResponseJSON `json:"-"`
}

// scanHARResponseJSON contains the JSON metadata for the struct [ScanHARResponse]
type scanHARResponseJSON struct {
	Log         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanHARResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanHARResponseJSON) RawJSON() string {
	return r.raw
}

type ScanHARResponseLog struct {
	Creator ScanHARResponseLogCreator `json:"creator,required"`
	Entries []ScanHARResponseLogEntry `json:"entries,required"`
	Pages   []ScanHARResponseLogPage  `json:"pages,required"`
	Version string                    `json:"version,required"`
	JSON    scanHARResponseLogJSON    `json:"-"`
}

// scanHARResponseLogJSON contains the JSON metadata for the struct
// [ScanHARResponseLog]
type scanHARResponseLogJSON struct {
	Creator     apijson.Field
	Entries     apijson.Field
	Pages       apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanHARResponseLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanHARResponseLogJSON) RawJSON() string {
	return r.raw
}

type ScanHARResponseLogCreator struct {
	Comment string                        `json:"comment,required"`
	Name    string                        `json:"name,required"`
	Version string                        `json:"version,required"`
	JSON    scanHARResponseLogCreatorJSON `json:"-"`
}

// scanHARResponseLogCreatorJSON contains the JSON metadata for the struct
// [ScanHARResponseLogCreator]
type scanHARResponseLogCreatorJSON struct {
	Comment     apijson.Field
	Name        apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanHARResponseLogCreator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanHARResponseLogCreatorJSON) RawJSON() string {
	return r.raw
}

type ScanHARResponseLogEntry struct {
	InitialPriority string                            `json:"_initialPriority,required"`
	InitiatorType   string                            `json:"_initiator_type,required"`
	Priority        string                            `json:"_priority,required"`
	RequestID       string                            `json:"_requestId,required"`
	RequestTime     float64                           `json:"_requestTime,required"`
	ResourceType    string                            `json:"_resourceType,required"`
	Cache           interface{}                       `json:"cache,required"`
	Connection      string                            `json:"connection,required"`
	Pageref         string                            `json:"pageref,required"`
	Request         ScanHARResponseLogEntriesRequest  `json:"request,required"`
	Response        ScanHARResponseLogEntriesResponse `json:"response,required"`
	ServerIPAddress string                            `json:"serverIPAddress,required"`
	StartedDateTime string                            `json:"startedDateTime,required"`
	Time            float64                           `json:"time,required"`
	JSON            scanHARResponseLogEntryJSON       `json:"-"`
}

// scanHARResponseLogEntryJSON contains the JSON metadata for the struct
// [ScanHARResponseLogEntry]
type scanHARResponseLogEntryJSON struct {
	InitialPriority apijson.Field
	InitiatorType   apijson.Field
	Priority        apijson.Field
	RequestID       apijson.Field
	RequestTime     apijson.Field
	ResourceType    apijson.Field
	Cache           apijson.Field
	Connection      apijson.Field
	Pageref         apijson.Field
	Request         apijson.Field
	Response        apijson.Field
	ServerIPAddress apijson.Field
	StartedDateTime apijson.Field
	Time            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ScanHARResponseLogEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanHARResponseLogEntryJSON) RawJSON() string {
	return r.raw
}

type ScanHARResponseLogEntriesRequest struct {
	BodySize    float64                                  `json:"bodySize,required"`
	Headers     []ScanHARResponseLogEntriesRequestHeader `json:"headers,required"`
	HeadersSize float64                                  `json:"headersSize,required"`
	HTTPVersion string                                   `json:"httpVersion,required"`
	Method      string                                   `json:"method,required"`
	URL         string                                   `json:"url,required"`
	JSON        scanHARResponseLogEntriesRequestJSON     `json:"-"`
}

// scanHARResponseLogEntriesRequestJSON contains the JSON metadata for the struct
// [ScanHARResponseLogEntriesRequest]
type scanHARResponseLogEntriesRequestJSON struct {
	BodySize    apijson.Field
	Headers     apijson.Field
	HeadersSize apijson.Field
	HTTPVersion apijson.Field
	Method      apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanHARResponseLogEntriesRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanHARResponseLogEntriesRequestJSON) RawJSON() string {
	return r.raw
}

type ScanHARResponseLogEntriesRequestHeader struct {
	Name  string                                     `json:"name,required"`
	Value string                                     `json:"value,required"`
	JSON  scanHARResponseLogEntriesRequestHeaderJSON `json:"-"`
}

// scanHARResponseLogEntriesRequestHeaderJSON contains the JSON metadata for the
// struct [ScanHARResponseLogEntriesRequestHeader]
type scanHARResponseLogEntriesRequestHeaderJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanHARResponseLogEntriesRequestHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanHARResponseLogEntriesRequestHeaderJSON) RawJSON() string {
	return r.raw
}

type ScanHARResponseLogEntriesResponse struct {
	TransferSize float64                                   `json:"_transferSize,required"`
	BodySize     float64                                   `json:"bodySize,required"`
	Content      ScanHARResponseLogEntriesResponseContent  `json:"content,required"`
	Headers      []ScanHARResponseLogEntriesResponseHeader `json:"headers,required"`
	HeadersSize  float64                                   `json:"headersSize,required"`
	HTTPVersion  string                                    `json:"httpVersion,required"`
	RedirectURL  string                                    `json:"redirectURL,required"`
	Status       float64                                   `json:"status,required"`
	StatusText   string                                    `json:"statusText,required"`
	JSON         scanHARResponseLogEntriesResponseJSON     `json:"-"`
}

// scanHARResponseLogEntriesResponseJSON contains the JSON metadata for the struct
// [ScanHARResponseLogEntriesResponse]
type scanHARResponseLogEntriesResponseJSON struct {
	TransferSize apijson.Field
	BodySize     apijson.Field
	Content      apijson.Field
	Headers      apijson.Field
	HeadersSize  apijson.Field
	HTTPVersion  apijson.Field
	RedirectURL  apijson.Field
	Status       apijson.Field
	StatusText   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ScanHARResponseLogEntriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanHARResponseLogEntriesResponseJSON) RawJSON() string {
	return r.raw
}

type ScanHARResponseLogEntriesResponseContent struct {
	MimeType    string                                       `json:"mimeType,required"`
	Size        float64                                      `json:"size,required"`
	Compression int64                                        `json:"compression"`
	JSON        scanHARResponseLogEntriesResponseContentJSON `json:"-"`
}

// scanHARResponseLogEntriesResponseContentJSON contains the JSON metadata for the
// struct [ScanHARResponseLogEntriesResponseContent]
type scanHARResponseLogEntriesResponseContentJSON struct {
	MimeType    apijson.Field
	Size        apijson.Field
	Compression apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanHARResponseLogEntriesResponseContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanHARResponseLogEntriesResponseContentJSON) RawJSON() string {
	return r.raw
}

type ScanHARResponseLogEntriesResponseHeader struct {
	Name  string                                      `json:"name,required"`
	Value string                                      `json:"value,required"`
	JSON  scanHARResponseLogEntriesResponseHeaderJSON `json:"-"`
}

// scanHARResponseLogEntriesResponseHeaderJSON contains the JSON metadata for the
// struct [ScanHARResponseLogEntriesResponseHeader]
type scanHARResponseLogEntriesResponseHeaderJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanHARResponseLogEntriesResponseHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanHARResponseLogEntriesResponseHeaderJSON) RawJSON() string {
	return r.raw
}

type ScanHARResponseLogPage struct {
	ID              string                             `json:"id,required"`
	PageTimings     ScanHARResponseLogPagesPageTimings `json:"pageTimings,required"`
	StartedDateTime string                             `json:"startedDateTime,required"`
	Title           string                             `json:"title,required"`
	JSON            scanHARResponseLogPageJSON         `json:"-"`
}

// scanHARResponseLogPageJSON contains the JSON metadata for the struct
// [ScanHARResponseLogPage]
type scanHARResponseLogPageJSON struct {
	ID              apijson.Field
	PageTimings     apijson.Field
	StartedDateTime apijson.Field
	Title           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ScanHARResponseLogPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanHARResponseLogPageJSON) RawJSON() string {
	return r.raw
}

type ScanHARResponseLogPagesPageTimings struct {
	OnContentLoad float64                                `json:"onContentLoad,required"`
	OnLoad        float64                                `json:"onLoad,required"`
	JSON          scanHARResponseLogPagesPageTimingsJSON `json:"-"`
}

// scanHARResponseLogPagesPageTimingsJSON contains the JSON metadata for the struct
// [ScanHARResponseLogPagesPageTimings]
type scanHARResponseLogPagesPageTimingsJSON struct {
	OnContentLoad apijson.Field
	OnLoad        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ScanHARResponseLogPagesPageTimings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanHARResponseLogPagesPageTimingsJSON) RawJSON() string {
	return r.raw
}

type ScanNewParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	URL       param.Field[string] `json:"url,required"`
	// Country to geo egress from
	Country     param.Field[ScanNewParamsCountry] `json:"country"`
	Customagent param.Field[string]               `json:"customagent"`
	// Set custom headers.
	CustomHeaders param.Field[map[string]string] `json:"customHeaders"`
	Referer       param.Field[string]            `json:"referer"`
	// Take multiple screenshots targeting different device types.
	ScreenshotsResolutions param.Field[[]ScanNewParamsScreenshotsResolution] `json:"screenshotsResolutions"`
	// The option `Public` means it will be included in listings like recent scans and
	// search results. `Unlisted` means it will not be included in the aforementioned
	// listings, users will need to have the scan's ID to access it. A a scan will be
	// automatically marked as unlisted if it fails, if it contains potential PII or
	// other sensitive material.
	Visibility param.Field[ScanNewParamsVisibility] `json:"visibility"`
}

func (r ScanNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Country to geo egress from
type ScanNewParamsCountry string

const (
	ScanNewParamsCountryAf ScanNewParamsCountry = "AF"
	ScanNewParamsCountryAl ScanNewParamsCountry = "AL"
	ScanNewParamsCountryDz ScanNewParamsCountry = "DZ"
	ScanNewParamsCountryAD ScanNewParamsCountry = "AD"
	ScanNewParamsCountryAo ScanNewParamsCountry = "AO"
	ScanNewParamsCountryAg ScanNewParamsCountry = "AG"
	ScanNewParamsCountryAr ScanNewParamsCountry = "AR"
	ScanNewParamsCountryAm ScanNewParamsCountry = "AM"
	ScanNewParamsCountryAu ScanNewParamsCountry = "AU"
	ScanNewParamsCountryAt ScanNewParamsCountry = "AT"
	ScanNewParamsCountryAz ScanNewParamsCountry = "AZ"
	ScanNewParamsCountryBh ScanNewParamsCountry = "BH"
	ScanNewParamsCountryBd ScanNewParamsCountry = "BD"
	ScanNewParamsCountryBb ScanNewParamsCountry = "BB"
	ScanNewParamsCountryBy ScanNewParamsCountry = "BY"
	ScanNewParamsCountryBe ScanNewParamsCountry = "BE"
	ScanNewParamsCountryBz ScanNewParamsCountry = "BZ"
	ScanNewParamsCountryBj ScanNewParamsCountry = "BJ"
	ScanNewParamsCountryBm ScanNewParamsCountry = "BM"
	ScanNewParamsCountryBt ScanNewParamsCountry = "BT"
	ScanNewParamsCountryBo ScanNewParamsCountry = "BO"
	ScanNewParamsCountryBa ScanNewParamsCountry = "BA"
	ScanNewParamsCountryBw ScanNewParamsCountry = "BW"
	ScanNewParamsCountryBr ScanNewParamsCountry = "BR"
	ScanNewParamsCountryBn ScanNewParamsCountry = "BN"
	ScanNewParamsCountryBg ScanNewParamsCountry = "BG"
	ScanNewParamsCountryBf ScanNewParamsCountry = "BF"
	ScanNewParamsCountryBi ScanNewParamsCountry = "BI"
	ScanNewParamsCountryKh ScanNewParamsCountry = "KH"
	ScanNewParamsCountryCm ScanNewParamsCountry = "CM"
	ScanNewParamsCountryCA ScanNewParamsCountry = "CA"
	ScanNewParamsCountryCv ScanNewParamsCountry = "CV"
	ScanNewParamsCountryKy ScanNewParamsCountry = "KY"
	ScanNewParamsCountryCf ScanNewParamsCountry = "CF"
	ScanNewParamsCountryTd ScanNewParamsCountry = "TD"
	ScanNewParamsCountryCl ScanNewParamsCountry = "CL"
	ScanNewParamsCountryCn ScanNewParamsCountry = "CN"
	ScanNewParamsCountryCo ScanNewParamsCountry = "CO"
	ScanNewParamsCountryKm ScanNewParamsCountry = "KM"
	ScanNewParamsCountryCg ScanNewParamsCountry = "CG"
	ScanNewParamsCountryCr ScanNewParamsCountry = "CR"
	ScanNewParamsCountryCi ScanNewParamsCountry = "CI"
	ScanNewParamsCountryHr ScanNewParamsCountry = "HR"
	ScanNewParamsCountryCu ScanNewParamsCountry = "CU"
	ScanNewParamsCountryCy ScanNewParamsCountry = "CY"
	ScanNewParamsCountryCz ScanNewParamsCountry = "CZ"
	ScanNewParamsCountryCd ScanNewParamsCountry = "CD"
	ScanNewParamsCountryDK ScanNewParamsCountry = "DK"
	ScanNewParamsCountryDj ScanNewParamsCountry = "DJ"
	ScanNewParamsCountryDm ScanNewParamsCountry = "DM"
	ScanNewParamsCountryDo ScanNewParamsCountry = "DO"
	ScanNewParamsCountryEc ScanNewParamsCountry = "EC"
	ScanNewParamsCountryEg ScanNewParamsCountry = "EG"
	ScanNewParamsCountrySv ScanNewParamsCountry = "SV"
	ScanNewParamsCountryGq ScanNewParamsCountry = "GQ"
	ScanNewParamsCountryEr ScanNewParamsCountry = "ER"
	ScanNewParamsCountryEe ScanNewParamsCountry = "EE"
	ScanNewParamsCountrySz ScanNewParamsCountry = "SZ"
	ScanNewParamsCountryEt ScanNewParamsCountry = "ET"
	ScanNewParamsCountryFj ScanNewParamsCountry = "FJ"
	ScanNewParamsCountryFi ScanNewParamsCountry = "FI"
	ScanNewParamsCountryFr ScanNewParamsCountry = "FR"
	ScanNewParamsCountryGa ScanNewParamsCountry = "GA"
	ScanNewParamsCountryGe ScanNewParamsCountry = "GE"
	ScanNewParamsCountryDe ScanNewParamsCountry = "DE"
	ScanNewParamsCountryGh ScanNewParamsCountry = "GH"
	ScanNewParamsCountryGr ScanNewParamsCountry = "GR"
	ScanNewParamsCountryGl ScanNewParamsCountry = "GL"
	ScanNewParamsCountryGd ScanNewParamsCountry = "GD"
	ScanNewParamsCountryGt ScanNewParamsCountry = "GT"
	ScanNewParamsCountryGn ScanNewParamsCountry = "GN"
	ScanNewParamsCountryGw ScanNewParamsCountry = "GW"
	ScanNewParamsCountryGy ScanNewParamsCountry = "GY"
	ScanNewParamsCountryHt ScanNewParamsCountry = "HT"
	ScanNewParamsCountryHn ScanNewParamsCountry = "HN"
	ScanNewParamsCountryHu ScanNewParamsCountry = "HU"
	ScanNewParamsCountryIs ScanNewParamsCountry = "IS"
	ScanNewParamsCountryIn ScanNewParamsCountry = "IN"
	ScanNewParamsCountryID ScanNewParamsCountry = "ID"
	ScanNewParamsCountryIr ScanNewParamsCountry = "IR"
	ScanNewParamsCountryIq ScanNewParamsCountry = "IQ"
	ScanNewParamsCountryIe ScanNewParamsCountry = "IE"
	ScanNewParamsCountryIl ScanNewParamsCountry = "IL"
	ScanNewParamsCountryIt ScanNewParamsCountry = "IT"
	ScanNewParamsCountryJm ScanNewParamsCountry = "JM"
	ScanNewParamsCountryJp ScanNewParamsCountry = "JP"
	ScanNewParamsCountryJo ScanNewParamsCountry = "JO"
	ScanNewParamsCountryKz ScanNewParamsCountry = "KZ"
	ScanNewParamsCountryKe ScanNewParamsCountry = "KE"
	ScanNewParamsCountryKi ScanNewParamsCountry = "KI"
	ScanNewParamsCountryKw ScanNewParamsCountry = "KW"
	ScanNewParamsCountryKg ScanNewParamsCountry = "KG"
	ScanNewParamsCountryLa ScanNewParamsCountry = "LA"
	ScanNewParamsCountryLv ScanNewParamsCountry = "LV"
	ScanNewParamsCountryLB ScanNewParamsCountry = "LB"
	ScanNewParamsCountryLs ScanNewParamsCountry = "LS"
	ScanNewParamsCountryLr ScanNewParamsCountry = "LR"
	ScanNewParamsCountryLy ScanNewParamsCountry = "LY"
	ScanNewParamsCountryLi ScanNewParamsCountry = "LI"
	ScanNewParamsCountryLt ScanNewParamsCountry = "LT"
	ScanNewParamsCountryLu ScanNewParamsCountry = "LU"
	ScanNewParamsCountryMo ScanNewParamsCountry = "MO"
	ScanNewParamsCountryMg ScanNewParamsCountry = "MG"
	ScanNewParamsCountryMw ScanNewParamsCountry = "MW"
	ScanNewParamsCountryMy ScanNewParamsCountry = "MY"
	ScanNewParamsCountryMv ScanNewParamsCountry = "MV"
	ScanNewParamsCountryMl ScanNewParamsCountry = "ML"
	ScanNewParamsCountryMr ScanNewParamsCountry = "MR"
	ScanNewParamsCountryMu ScanNewParamsCountry = "MU"
	ScanNewParamsCountryMX ScanNewParamsCountry = "MX"
	ScanNewParamsCountryFm ScanNewParamsCountry = "FM"
	ScanNewParamsCountryMd ScanNewParamsCountry = "MD"
	ScanNewParamsCountryMc ScanNewParamsCountry = "MC"
	ScanNewParamsCountryMn ScanNewParamsCountry = "MN"
	ScanNewParamsCountryMs ScanNewParamsCountry = "MS"
	ScanNewParamsCountryMa ScanNewParamsCountry = "MA"
	ScanNewParamsCountryMz ScanNewParamsCountry = "MZ"
	ScanNewParamsCountryMm ScanNewParamsCountry = "MM"
	ScanNewParamsCountryNa ScanNewParamsCountry = "NA"
	ScanNewParamsCountryNr ScanNewParamsCountry = "NR"
	ScanNewParamsCountryNp ScanNewParamsCountry = "NP"
	ScanNewParamsCountryNl ScanNewParamsCountry = "NL"
	ScanNewParamsCountryNz ScanNewParamsCountry = "NZ"
	ScanNewParamsCountryNi ScanNewParamsCountry = "NI"
	ScanNewParamsCountryNe ScanNewParamsCountry = "NE"
	ScanNewParamsCountryNg ScanNewParamsCountry = "NG"
	ScanNewParamsCountryKp ScanNewParamsCountry = "KP"
	ScanNewParamsCountryMk ScanNewParamsCountry = "MK"
	ScanNewParamsCountryNo ScanNewParamsCountry = "NO"
	ScanNewParamsCountryOm ScanNewParamsCountry = "OM"
	ScanNewParamsCountryPk ScanNewParamsCountry = "PK"
	ScanNewParamsCountryPs ScanNewParamsCountry = "PS"
	ScanNewParamsCountryPa ScanNewParamsCountry = "PA"
	ScanNewParamsCountryPg ScanNewParamsCountry = "PG"
	ScanNewParamsCountryPy ScanNewParamsCountry = "PY"
	ScanNewParamsCountryPe ScanNewParamsCountry = "PE"
	ScanNewParamsCountryPh ScanNewParamsCountry = "PH"
	ScanNewParamsCountryPl ScanNewParamsCountry = "PL"
	ScanNewParamsCountryPt ScanNewParamsCountry = "PT"
	ScanNewParamsCountryQa ScanNewParamsCountry = "QA"
	ScanNewParamsCountryRo ScanNewParamsCountry = "RO"
	ScanNewParamsCountryRu ScanNewParamsCountry = "RU"
	ScanNewParamsCountryRw ScanNewParamsCountry = "RW"
	ScanNewParamsCountrySh ScanNewParamsCountry = "SH"
	ScanNewParamsCountryKn ScanNewParamsCountry = "KN"
	ScanNewParamsCountryLc ScanNewParamsCountry = "LC"
	ScanNewParamsCountryVc ScanNewParamsCountry = "VC"
	ScanNewParamsCountryWs ScanNewParamsCountry = "WS"
	ScanNewParamsCountrySm ScanNewParamsCountry = "SM"
	ScanNewParamsCountrySt ScanNewParamsCountry = "ST"
	ScanNewParamsCountrySa ScanNewParamsCountry = "SA"
	ScanNewParamsCountrySn ScanNewParamsCountry = "SN"
	ScanNewParamsCountryRs ScanNewParamsCountry = "RS"
	ScanNewParamsCountrySc ScanNewParamsCountry = "SC"
	ScanNewParamsCountrySl ScanNewParamsCountry = "SL"
	ScanNewParamsCountrySk ScanNewParamsCountry = "SK"
	ScanNewParamsCountrySi ScanNewParamsCountry = "SI"
	ScanNewParamsCountrySb ScanNewParamsCountry = "SB"
	ScanNewParamsCountrySo ScanNewParamsCountry = "SO"
	ScanNewParamsCountryZa ScanNewParamsCountry = "ZA"
	ScanNewParamsCountryKr ScanNewParamsCountry = "KR"
	ScanNewParamsCountrySS ScanNewParamsCountry = "SS"
	ScanNewParamsCountryEs ScanNewParamsCountry = "ES"
	ScanNewParamsCountryLk ScanNewParamsCountry = "LK"
	ScanNewParamsCountrySd ScanNewParamsCountry = "SD"
	ScanNewParamsCountrySr ScanNewParamsCountry = "SR"
	ScanNewParamsCountrySe ScanNewParamsCountry = "SE"
	ScanNewParamsCountryCh ScanNewParamsCountry = "CH"
	ScanNewParamsCountrySy ScanNewParamsCountry = "SY"
	ScanNewParamsCountryTw ScanNewParamsCountry = "TW"
	ScanNewParamsCountryTj ScanNewParamsCountry = "TJ"
	ScanNewParamsCountryTz ScanNewParamsCountry = "TZ"
	ScanNewParamsCountryTh ScanNewParamsCountry = "TH"
	ScanNewParamsCountryBs ScanNewParamsCountry = "BS"
	ScanNewParamsCountryGm ScanNewParamsCountry = "GM"
	ScanNewParamsCountryTl ScanNewParamsCountry = "TL"
	ScanNewParamsCountryTg ScanNewParamsCountry = "TG"
	ScanNewParamsCountryTo ScanNewParamsCountry = "TO"
	ScanNewParamsCountryTt ScanNewParamsCountry = "TT"
	ScanNewParamsCountryTn ScanNewParamsCountry = "TN"
	ScanNewParamsCountryTr ScanNewParamsCountry = "TR"
	ScanNewParamsCountryTm ScanNewParamsCountry = "TM"
	ScanNewParamsCountryUg ScanNewParamsCountry = "UG"
	ScanNewParamsCountryUA ScanNewParamsCountry = "UA"
	ScanNewParamsCountryAe ScanNewParamsCountry = "AE"
	ScanNewParamsCountryGB ScanNewParamsCountry = "GB"
	ScanNewParamsCountryUs ScanNewParamsCountry = "US"
	ScanNewParamsCountryUy ScanNewParamsCountry = "UY"
	ScanNewParamsCountryUz ScanNewParamsCountry = "UZ"
	ScanNewParamsCountryVu ScanNewParamsCountry = "VU"
	ScanNewParamsCountryVe ScanNewParamsCountry = "VE"
	ScanNewParamsCountryVn ScanNewParamsCountry = "VN"
	ScanNewParamsCountryYe ScanNewParamsCountry = "YE"
	ScanNewParamsCountryZm ScanNewParamsCountry = "ZM"
	ScanNewParamsCountryZw ScanNewParamsCountry = "ZW"
)

func (r ScanNewParamsCountry) IsKnown() bool {
	switch r {
	case ScanNewParamsCountryAf, ScanNewParamsCountryAl, ScanNewParamsCountryDz, ScanNewParamsCountryAD, ScanNewParamsCountryAo, ScanNewParamsCountryAg, ScanNewParamsCountryAr, ScanNewParamsCountryAm, ScanNewParamsCountryAu, ScanNewParamsCountryAt, ScanNewParamsCountryAz, ScanNewParamsCountryBh, ScanNewParamsCountryBd, ScanNewParamsCountryBb, ScanNewParamsCountryBy, ScanNewParamsCountryBe, ScanNewParamsCountryBz, ScanNewParamsCountryBj, ScanNewParamsCountryBm, ScanNewParamsCountryBt, ScanNewParamsCountryBo, ScanNewParamsCountryBa, ScanNewParamsCountryBw, ScanNewParamsCountryBr, ScanNewParamsCountryBn, ScanNewParamsCountryBg, ScanNewParamsCountryBf, ScanNewParamsCountryBi, ScanNewParamsCountryKh, ScanNewParamsCountryCm, ScanNewParamsCountryCA, ScanNewParamsCountryCv, ScanNewParamsCountryKy, ScanNewParamsCountryCf, ScanNewParamsCountryTd, ScanNewParamsCountryCl, ScanNewParamsCountryCn, ScanNewParamsCountryCo, ScanNewParamsCountryKm, ScanNewParamsCountryCg, ScanNewParamsCountryCr, ScanNewParamsCountryCi, ScanNewParamsCountryHr, ScanNewParamsCountryCu, ScanNewParamsCountryCy, ScanNewParamsCountryCz, ScanNewParamsCountryCd, ScanNewParamsCountryDK, ScanNewParamsCountryDj, ScanNewParamsCountryDm, ScanNewParamsCountryDo, ScanNewParamsCountryEc, ScanNewParamsCountryEg, ScanNewParamsCountrySv, ScanNewParamsCountryGq, ScanNewParamsCountryEr, ScanNewParamsCountryEe, ScanNewParamsCountrySz, ScanNewParamsCountryEt, ScanNewParamsCountryFj, ScanNewParamsCountryFi, ScanNewParamsCountryFr, ScanNewParamsCountryGa, ScanNewParamsCountryGe, ScanNewParamsCountryDe, ScanNewParamsCountryGh, ScanNewParamsCountryGr, ScanNewParamsCountryGl, ScanNewParamsCountryGd, ScanNewParamsCountryGt, ScanNewParamsCountryGn, ScanNewParamsCountryGw, ScanNewParamsCountryGy, ScanNewParamsCountryHt, ScanNewParamsCountryHn, ScanNewParamsCountryHu, ScanNewParamsCountryIs, ScanNewParamsCountryIn, ScanNewParamsCountryID, ScanNewParamsCountryIr, ScanNewParamsCountryIq, ScanNewParamsCountryIe, ScanNewParamsCountryIl, ScanNewParamsCountryIt, ScanNewParamsCountryJm, ScanNewParamsCountryJp, ScanNewParamsCountryJo, ScanNewParamsCountryKz, ScanNewParamsCountryKe, ScanNewParamsCountryKi, ScanNewParamsCountryKw, ScanNewParamsCountryKg, ScanNewParamsCountryLa, ScanNewParamsCountryLv, ScanNewParamsCountryLB, ScanNewParamsCountryLs, ScanNewParamsCountryLr, ScanNewParamsCountryLy, ScanNewParamsCountryLi, ScanNewParamsCountryLt, ScanNewParamsCountryLu, ScanNewParamsCountryMo, ScanNewParamsCountryMg, ScanNewParamsCountryMw, ScanNewParamsCountryMy, ScanNewParamsCountryMv, ScanNewParamsCountryMl, ScanNewParamsCountryMr, ScanNewParamsCountryMu, ScanNewParamsCountryMX, ScanNewParamsCountryFm, ScanNewParamsCountryMd, ScanNewParamsCountryMc, ScanNewParamsCountryMn, ScanNewParamsCountryMs, ScanNewParamsCountryMa, ScanNewParamsCountryMz, ScanNewParamsCountryMm, ScanNewParamsCountryNa, ScanNewParamsCountryNr, ScanNewParamsCountryNp, ScanNewParamsCountryNl, ScanNewParamsCountryNz, ScanNewParamsCountryNi, ScanNewParamsCountryNe, ScanNewParamsCountryNg, ScanNewParamsCountryKp, ScanNewParamsCountryMk, ScanNewParamsCountryNo, ScanNewParamsCountryOm, ScanNewParamsCountryPk, ScanNewParamsCountryPs, ScanNewParamsCountryPa, ScanNewParamsCountryPg, ScanNewParamsCountryPy, ScanNewParamsCountryPe, ScanNewParamsCountryPh, ScanNewParamsCountryPl, ScanNewParamsCountryPt, ScanNewParamsCountryQa, ScanNewParamsCountryRo, ScanNewParamsCountryRu, ScanNewParamsCountryRw, ScanNewParamsCountrySh, ScanNewParamsCountryKn, ScanNewParamsCountryLc, ScanNewParamsCountryVc, ScanNewParamsCountryWs, ScanNewParamsCountrySm, ScanNewParamsCountrySt, ScanNewParamsCountrySa, ScanNewParamsCountrySn, ScanNewParamsCountryRs, ScanNewParamsCountrySc, ScanNewParamsCountrySl, ScanNewParamsCountrySk, ScanNewParamsCountrySi, ScanNewParamsCountrySb, ScanNewParamsCountrySo, ScanNewParamsCountryZa, ScanNewParamsCountryKr, ScanNewParamsCountrySS, ScanNewParamsCountryEs, ScanNewParamsCountryLk, ScanNewParamsCountrySd, ScanNewParamsCountrySr, ScanNewParamsCountrySe, ScanNewParamsCountryCh, ScanNewParamsCountrySy, ScanNewParamsCountryTw, ScanNewParamsCountryTj, ScanNewParamsCountryTz, ScanNewParamsCountryTh, ScanNewParamsCountryBs, ScanNewParamsCountryGm, ScanNewParamsCountryTl, ScanNewParamsCountryTg, ScanNewParamsCountryTo, ScanNewParamsCountryTt, ScanNewParamsCountryTn, ScanNewParamsCountryTr, ScanNewParamsCountryTm, ScanNewParamsCountryUg, ScanNewParamsCountryUA, ScanNewParamsCountryAe, ScanNewParamsCountryGB, ScanNewParamsCountryUs, ScanNewParamsCountryUy, ScanNewParamsCountryUz, ScanNewParamsCountryVu, ScanNewParamsCountryVe, ScanNewParamsCountryVn, ScanNewParamsCountryYe, ScanNewParamsCountryZm, ScanNewParamsCountryZw:
		return true
	}
	return false
}

// Device resolutions.
type ScanNewParamsScreenshotsResolution string

const (
	ScanNewParamsScreenshotsResolutionDesktop ScanNewParamsScreenshotsResolution = "desktop"
	ScanNewParamsScreenshotsResolutionMobile  ScanNewParamsScreenshotsResolution = "mobile"
	ScanNewParamsScreenshotsResolutionTablet  ScanNewParamsScreenshotsResolution = "tablet"
)

func (r ScanNewParamsScreenshotsResolution) IsKnown() bool {
	switch r {
	case ScanNewParamsScreenshotsResolutionDesktop, ScanNewParamsScreenshotsResolutionMobile, ScanNewParamsScreenshotsResolutionTablet:
		return true
	}
	return false
}

// The option `Public` means it will be included in listings like recent scans and
// search results. `Unlisted` means it will not be included in the aforementioned
// listings, users will need to have the scan's ID to access it. A a scan will be
// automatically marked as unlisted if it fails, if it contains potential PII or
// other sensitive material.
type ScanNewParamsVisibility string

const (
	ScanNewParamsVisibilityPublic   ScanNewParamsVisibility = "Public"
	ScanNewParamsVisibilityUnlisted ScanNewParamsVisibility = "Unlisted"
)

func (r ScanNewParamsVisibility) IsKnown() bool {
	switch r {
	case ScanNewParamsVisibilityPublic, ScanNewParamsVisibilityUnlisted:
		return true
	}
	return false
}

type ScanListParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// Filter scans
	Q param.Field[string] `query:"q"`
	// Limit the number of objects in the response.
	Size param.Field[int64] `query:"size"`
}

// URLQuery serializes [ScanListParams]'s query parameters as `url.Values`.
func (r ScanListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ScanBulkNewParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// List of urls to scan (up to a 100).
	Body []ScanBulkNewParamsBody `json:"body"`
}

func (r ScanBulkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ScanBulkNewParamsBody struct {
	URL         param.Field[string] `json:"url,required"`
	Customagent param.Field[string] `json:"customagent"`
	// Set custom headers.
	CustomHeaders param.Field[map[string]string] `json:"customHeaders"`
	Referer       param.Field[string]            `json:"referer"`
	// Take multiple screenshots targeting different device types.
	ScreenshotsResolutions param.Field[[]ScanBulkNewParamsBodyScreenshotsResolution] `json:"screenshotsResolutions"`
	// The option `Public` means it will be included in listings like recent scans and
	// search results. `Unlisted` means it will not be included in the aforementioned
	// listings, users will need to have the scan's ID to access it. A a scan will be
	// automatically marked as unlisted if it fails, if it contains potential PII or
	// other sensitive material.
	Visibility param.Field[ScanBulkNewParamsBodyVisibility] `json:"visibility"`
}

func (r ScanBulkNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Device resolutions.
type ScanBulkNewParamsBodyScreenshotsResolution string

const (
	ScanBulkNewParamsBodyScreenshotsResolutionDesktop ScanBulkNewParamsBodyScreenshotsResolution = "desktop"
	ScanBulkNewParamsBodyScreenshotsResolutionMobile  ScanBulkNewParamsBodyScreenshotsResolution = "mobile"
	ScanBulkNewParamsBodyScreenshotsResolutionTablet  ScanBulkNewParamsBodyScreenshotsResolution = "tablet"
)

func (r ScanBulkNewParamsBodyScreenshotsResolution) IsKnown() bool {
	switch r {
	case ScanBulkNewParamsBodyScreenshotsResolutionDesktop, ScanBulkNewParamsBodyScreenshotsResolutionMobile, ScanBulkNewParamsBodyScreenshotsResolutionTablet:
		return true
	}
	return false
}

// The option `Public` means it will be included in listings like recent scans and
// search results. `Unlisted` means it will not be included in the aforementioned
// listings, users will need to have the scan's ID to access it. A a scan will be
// automatically marked as unlisted if it fails, if it contains potential PII or
// other sensitive material.
type ScanBulkNewParamsBodyVisibility string

const (
	ScanBulkNewParamsBodyVisibilityPublic   ScanBulkNewParamsBodyVisibility = "Public"
	ScanBulkNewParamsBodyVisibilityUnlisted ScanBulkNewParamsBodyVisibility = "Unlisted"
)

func (r ScanBulkNewParamsBodyVisibility) IsKnown() bool {
	switch r {
	case ScanBulkNewParamsBodyVisibilityPublic, ScanBulkNewParamsBodyVisibilityUnlisted:
		return true
	}
	return false
}

type ScanDOMParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ScanGetParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ScanHARParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ScanScreenshotParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// Target device type.
	Resolution param.Field[ScanScreenshotParamsResolution] `query:"resolution"`
}

// URLQuery serializes [ScanScreenshotParams]'s query parameters as `url.Values`.
func (r ScanScreenshotParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Target device type.
type ScanScreenshotParamsResolution string

const (
	ScanScreenshotParamsResolutionDesktop ScanScreenshotParamsResolution = "desktop"
	ScanScreenshotParamsResolutionMobile  ScanScreenshotParamsResolution = "mobile"
	ScanScreenshotParamsResolutionTablet  ScanScreenshotParamsResolution = "tablet"
)

func (r ScanScreenshotParamsResolution) IsKnown() bool {
	switch r {
	case ScanScreenshotParamsResolutionDesktop, ScanScreenshotParamsResolutionMobile, ScanScreenshotParamsResolutionTablet:
		return true
	}
	return false
}
