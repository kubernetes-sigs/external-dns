// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package browser_rendering

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

// ScrapeService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScrapeService] method instead.
type ScrapeService struct {
	Options []option.RequestOption
}

// NewScrapeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewScrapeService(opts ...option.RequestOption) (r *ScrapeService) {
	r = &ScrapeService{}
	r.Options = opts
	return
}

// Get meta attributes like height, width, text and others of selected elements.
func (r *ScrapeService) New(ctx context.Context, params ScrapeNewParams, opts ...option.RequestOption) (res *[]ScrapeNewResponse, err error) {
	var env ScrapeNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/browser-rendering/scrape", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ScrapeNewResponse struct {
	Results ScrapeNewResponseResults `json:"results,required"`
	// Selector
	Selector string                `json:"selector,required"`
	JSON     scrapeNewResponseJSON `json:"-"`
}

// scrapeNewResponseJSON contains the JSON metadata for the struct
// [ScrapeNewResponse]
type scrapeNewResponseJSON struct {
	Results     apijson.Field
	Selector    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScrapeNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scrapeNewResponseJSON) RawJSON() string {
	return r.raw
}

type ScrapeNewResponseResults struct {
	Attributes []ScrapeNewResponseResultsAttribute `json:"attributes,required"`
	// Element height
	Height float64 `json:"height,required"`
	// Html content
	HTML string `json:"html,required"`
	// Element left
	Left float64 `json:"left,required"`
	// Text content
	Text string `json:"text,required"`
	// Element top
	Top float64 `json:"top,required"`
	// Element width
	Width float64                      `json:"width,required"`
	JSON  scrapeNewResponseResultsJSON `json:"-"`
}

// scrapeNewResponseResultsJSON contains the JSON metadata for the struct
// [ScrapeNewResponseResults]
type scrapeNewResponseResultsJSON struct {
	Attributes  apijson.Field
	Height      apijson.Field
	HTML        apijson.Field
	Left        apijson.Field
	Text        apijson.Field
	Top         apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScrapeNewResponseResults) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scrapeNewResponseResultsJSON) RawJSON() string {
	return r.raw
}

type ScrapeNewResponseResultsAttribute struct {
	// Attribute name
	Name string `json:"name,required"`
	// Attribute value
	Value string                                `json:"value,required"`
	JSON  scrapeNewResponseResultsAttributeJSON `json:"-"`
}

// scrapeNewResponseResultsAttributeJSON contains the JSON metadata for the struct
// [ScrapeNewResponseResultsAttribute]
type scrapeNewResponseResultsAttributeJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScrapeNewResponseResultsAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scrapeNewResponseResultsAttributeJSON) RawJSON() string {
	return r.raw
}

type ScrapeNewParams struct {
	// Account ID.
	AccountID param.Field[string]                   `path:"account_id,required"`
	Elements  param.Field[[]ScrapeNewParamsElement] `json:"elements,required"`
	// Cache TTL default is 5s. Set to 0 to disable.
	CacheTTL param.Field[float64] `query:"cacheTTL"`
	// The maximum duration allowed for the browser action to complete after the page
	// has loaded (such as taking screenshots, extracting content, or generating PDFs).
	// If this time limit is exceeded, the action stops and returns a timeout error.
	ActionTimeout param.Field[float64] `json:"actionTimeout"`
	// Adds a `<script>` tag into the page with the desired URL or content.
	AddScriptTag param.Field[[]ScrapeNewParamsAddScriptTag] `json:"addScriptTag"`
	// Adds a `<link rel="stylesheet">` tag into the page with the desired URL or a
	// `<style type="text/css">` tag with the content.
	AddStyleTag param.Field[[]ScrapeNewParamsAddStyleTag] `json:"addStyleTag"`
	// Only allow requests that match the provided regex patterns, eg. '/^.\*\.(css)'.
	AllowRequestPattern param.Field[[]string] `json:"allowRequestPattern"`
	// Only allow requests that match the provided resource types, eg. 'image' or
	// 'script'.
	AllowResourceTypes param.Field[[]ScrapeNewParamsAllowResourceType] `json:"allowResourceTypes"`
	// Provide credentials for HTTP authentication.
	Authenticate param.Field[ScrapeNewParamsAuthenticate] `json:"authenticate"`
	// Attempt to proceed when 'awaited' events fail or timeout.
	BestAttempt param.Field[bool] `json:"bestAttempt"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setcookie).
	Cookies          param.Field[[]ScrapeNewParamsCookie] `json:"cookies"`
	EmulateMediaType param.Field[string]                  `json:"emulateMediaType"`
	// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
	GotoOptions param.Field[ScrapeNewParamsGotoOptions] `json:"gotoOptions"`
	// Set the content of the page, eg: `<h1>Hello World!!</h1>`. Either `html` or
	// `url` must be set.
	HTML param.Field[string] `json:"html"`
	// Block undesired requests that match the provided regex patterns, eg.
	// '/^.\*\.(css)'.
	RejectRequestPattern param.Field[[]string] `json:"rejectRequestPattern"`
	// Block undesired requests that match the provided resource types, eg. 'image' or
	// 'script'.
	RejectResourceTypes  param.Field[[]ScrapeNewParamsRejectResourceType] `json:"rejectResourceTypes"`
	SetExtraHTTPHeaders  param.Field[map[string]string]                   `json:"setExtraHTTPHeaders"`
	SetJavaScriptEnabled param.Field[bool]                                `json:"setJavaScriptEnabled"`
	// URL to navigate to, eg. `https://example.com`.
	URL       param.Field[string] `json:"url" format:"uri"`
	UserAgent param.Field[string] `json:"userAgent"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
	Viewport param.Field[ScrapeNewParamsViewport] `json:"viewport"`
	// Wait for the selector to appear in page. Check
	// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
	WaitForSelector param.Field[ScrapeNewParamsWaitForSelector] `json:"waitForSelector"`
	// Waits for a specified timeout before continuing.
	WaitForTimeout param.Field[float64] `json:"waitForTimeout"`
}

func (r ScrapeNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [ScrapeNewParams]'s query parameters as `url.Values`.
func (r ScrapeNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ScrapeNewParamsElement struct {
	Selector param.Field[string] `json:"selector,required"`
}

func (r ScrapeNewParamsElement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScrapeNewParamsAddScriptTag struct {
	ID      param.Field[string] `json:"id"`
	Content param.Field[string] `json:"content"`
	Type    param.Field[string] `json:"type"`
	URL     param.Field[string] `json:"url"`
}

func (r ScrapeNewParamsAddScriptTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScrapeNewParamsAddStyleTag struct {
	Content param.Field[string] `json:"content"`
	URL     param.Field[string] `json:"url"`
}

func (r ScrapeNewParamsAddStyleTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScrapeNewParamsAllowResourceType string

const (
	ScrapeNewParamsAllowResourceTypeDocument           ScrapeNewParamsAllowResourceType = "document"
	ScrapeNewParamsAllowResourceTypeStylesheet         ScrapeNewParamsAllowResourceType = "stylesheet"
	ScrapeNewParamsAllowResourceTypeImage              ScrapeNewParamsAllowResourceType = "image"
	ScrapeNewParamsAllowResourceTypeMedia              ScrapeNewParamsAllowResourceType = "media"
	ScrapeNewParamsAllowResourceTypeFont               ScrapeNewParamsAllowResourceType = "font"
	ScrapeNewParamsAllowResourceTypeScript             ScrapeNewParamsAllowResourceType = "script"
	ScrapeNewParamsAllowResourceTypeTexttrack          ScrapeNewParamsAllowResourceType = "texttrack"
	ScrapeNewParamsAllowResourceTypeXHR                ScrapeNewParamsAllowResourceType = "xhr"
	ScrapeNewParamsAllowResourceTypeFetch              ScrapeNewParamsAllowResourceType = "fetch"
	ScrapeNewParamsAllowResourceTypePrefetch           ScrapeNewParamsAllowResourceType = "prefetch"
	ScrapeNewParamsAllowResourceTypeEventsource        ScrapeNewParamsAllowResourceType = "eventsource"
	ScrapeNewParamsAllowResourceTypeWebsocket          ScrapeNewParamsAllowResourceType = "websocket"
	ScrapeNewParamsAllowResourceTypeManifest           ScrapeNewParamsAllowResourceType = "manifest"
	ScrapeNewParamsAllowResourceTypeSignedexchange     ScrapeNewParamsAllowResourceType = "signedexchange"
	ScrapeNewParamsAllowResourceTypePing               ScrapeNewParamsAllowResourceType = "ping"
	ScrapeNewParamsAllowResourceTypeCspviolationreport ScrapeNewParamsAllowResourceType = "cspviolationreport"
	ScrapeNewParamsAllowResourceTypePreflight          ScrapeNewParamsAllowResourceType = "preflight"
	ScrapeNewParamsAllowResourceTypeOther              ScrapeNewParamsAllowResourceType = "other"
)

func (r ScrapeNewParamsAllowResourceType) IsKnown() bool {
	switch r {
	case ScrapeNewParamsAllowResourceTypeDocument, ScrapeNewParamsAllowResourceTypeStylesheet, ScrapeNewParamsAllowResourceTypeImage, ScrapeNewParamsAllowResourceTypeMedia, ScrapeNewParamsAllowResourceTypeFont, ScrapeNewParamsAllowResourceTypeScript, ScrapeNewParamsAllowResourceTypeTexttrack, ScrapeNewParamsAllowResourceTypeXHR, ScrapeNewParamsAllowResourceTypeFetch, ScrapeNewParamsAllowResourceTypePrefetch, ScrapeNewParamsAllowResourceTypeEventsource, ScrapeNewParamsAllowResourceTypeWebsocket, ScrapeNewParamsAllowResourceTypeManifest, ScrapeNewParamsAllowResourceTypeSignedexchange, ScrapeNewParamsAllowResourceTypePing, ScrapeNewParamsAllowResourceTypeCspviolationreport, ScrapeNewParamsAllowResourceTypePreflight, ScrapeNewParamsAllowResourceTypeOther:
		return true
	}
	return false
}

// Provide credentials for HTTP authentication.
type ScrapeNewParamsAuthenticate struct {
	Password param.Field[string] `json:"password,required"`
	Username param.Field[string] `json:"username,required"`
}

func (r ScrapeNewParamsAuthenticate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScrapeNewParamsCookie struct {
	Name         param.Field[string]                             `json:"name,required"`
	Value        param.Field[string]                             `json:"value,required"`
	Domain       param.Field[string]                             `json:"domain"`
	Expires      param.Field[float64]                            `json:"expires"`
	HTTPOnly     param.Field[bool]                               `json:"httpOnly"`
	PartitionKey param.Field[string]                             `json:"partitionKey"`
	Path         param.Field[string]                             `json:"path"`
	Priority     param.Field[ScrapeNewParamsCookiesPriority]     `json:"priority"`
	SameParty    param.Field[bool]                               `json:"sameParty"`
	SameSite     param.Field[ScrapeNewParamsCookiesSameSite]     `json:"sameSite"`
	Secure       param.Field[bool]                               `json:"secure"`
	SourcePort   param.Field[float64]                            `json:"sourcePort"`
	SourceScheme param.Field[ScrapeNewParamsCookiesSourceScheme] `json:"sourceScheme"`
	URL          param.Field[string]                             `json:"url"`
}

func (r ScrapeNewParamsCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScrapeNewParamsCookiesPriority string

const (
	ScrapeNewParamsCookiesPriorityLow    ScrapeNewParamsCookiesPriority = "Low"
	ScrapeNewParamsCookiesPriorityMedium ScrapeNewParamsCookiesPriority = "Medium"
	ScrapeNewParamsCookiesPriorityHigh   ScrapeNewParamsCookiesPriority = "High"
)

func (r ScrapeNewParamsCookiesPriority) IsKnown() bool {
	switch r {
	case ScrapeNewParamsCookiesPriorityLow, ScrapeNewParamsCookiesPriorityMedium, ScrapeNewParamsCookiesPriorityHigh:
		return true
	}
	return false
}

type ScrapeNewParamsCookiesSameSite string

const (
	ScrapeNewParamsCookiesSameSiteStrict ScrapeNewParamsCookiesSameSite = "Strict"
	ScrapeNewParamsCookiesSameSiteLax    ScrapeNewParamsCookiesSameSite = "Lax"
	ScrapeNewParamsCookiesSameSiteNone   ScrapeNewParamsCookiesSameSite = "None"
)

func (r ScrapeNewParamsCookiesSameSite) IsKnown() bool {
	switch r {
	case ScrapeNewParamsCookiesSameSiteStrict, ScrapeNewParamsCookiesSameSiteLax, ScrapeNewParamsCookiesSameSiteNone:
		return true
	}
	return false
}

type ScrapeNewParamsCookiesSourceScheme string

const (
	ScrapeNewParamsCookiesSourceSchemeUnset     ScrapeNewParamsCookiesSourceScheme = "Unset"
	ScrapeNewParamsCookiesSourceSchemeNonSecure ScrapeNewParamsCookiesSourceScheme = "NonSecure"
	ScrapeNewParamsCookiesSourceSchemeSecure    ScrapeNewParamsCookiesSourceScheme = "Secure"
)

func (r ScrapeNewParamsCookiesSourceScheme) IsKnown() bool {
	switch r {
	case ScrapeNewParamsCookiesSourceSchemeUnset, ScrapeNewParamsCookiesSourceSchemeNonSecure, ScrapeNewParamsCookiesSourceSchemeSecure:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
type ScrapeNewParamsGotoOptions struct {
	Referer        param.Field[string]                                   `json:"referer"`
	ReferrerPolicy param.Field[string]                                   `json:"referrerPolicy"`
	Timeout        param.Field[float64]                                  `json:"timeout"`
	WaitUntil      param.Field[ScrapeNewParamsGotoOptionsWaitUntilUnion] `json:"waitUntil"`
}

func (r ScrapeNewParamsGotoOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [browser_rendering.ScrapeNewParamsGotoOptionsWaitUntilString],
// [browser_rendering.ScrapeNewParamsGotoOptionsWaitUntilArray].
type ScrapeNewParamsGotoOptionsWaitUntilUnion interface {
	implementsScrapeNewParamsGotoOptionsWaitUntilUnion()
}

type ScrapeNewParamsGotoOptionsWaitUntilString string

const (
	ScrapeNewParamsGotoOptionsWaitUntilStringLoad             ScrapeNewParamsGotoOptionsWaitUntilString = "load"
	ScrapeNewParamsGotoOptionsWaitUntilStringDomcontentloaded ScrapeNewParamsGotoOptionsWaitUntilString = "domcontentloaded"
	ScrapeNewParamsGotoOptionsWaitUntilStringNetworkidle0     ScrapeNewParamsGotoOptionsWaitUntilString = "networkidle0"
	ScrapeNewParamsGotoOptionsWaitUntilStringNetworkidle2     ScrapeNewParamsGotoOptionsWaitUntilString = "networkidle2"
)

func (r ScrapeNewParamsGotoOptionsWaitUntilString) IsKnown() bool {
	switch r {
	case ScrapeNewParamsGotoOptionsWaitUntilStringLoad, ScrapeNewParamsGotoOptionsWaitUntilStringDomcontentloaded, ScrapeNewParamsGotoOptionsWaitUntilStringNetworkidle0, ScrapeNewParamsGotoOptionsWaitUntilStringNetworkidle2:
		return true
	}
	return false
}

func (r ScrapeNewParamsGotoOptionsWaitUntilString) implementsScrapeNewParamsGotoOptionsWaitUntilUnion() {
}

type ScrapeNewParamsGotoOptionsWaitUntilArray []ScrapeNewParamsGotoOptionsWaitUntilArrayItem

func (r ScrapeNewParamsGotoOptionsWaitUntilArray) implementsScrapeNewParamsGotoOptionsWaitUntilUnion() {
}

type ScrapeNewParamsGotoOptionsWaitUntilArrayItem string

const (
	ScrapeNewParamsGotoOptionsWaitUntilArrayItemLoad             ScrapeNewParamsGotoOptionsWaitUntilArrayItem = "load"
	ScrapeNewParamsGotoOptionsWaitUntilArrayItemDomcontentloaded ScrapeNewParamsGotoOptionsWaitUntilArrayItem = "domcontentloaded"
	ScrapeNewParamsGotoOptionsWaitUntilArrayItemNetworkidle0     ScrapeNewParamsGotoOptionsWaitUntilArrayItem = "networkidle0"
	ScrapeNewParamsGotoOptionsWaitUntilArrayItemNetworkidle2     ScrapeNewParamsGotoOptionsWaitUntilArrayItem = "networkidle2"
)

func (r ScrapeNewParamsGotoOptionsWaitUntilArrayItem) IsKnown() bool {
	switch r {
	case ScrapeNewParamsGotoOptionsWaitUntilArrayItemLoad, ScrapeNewParamsGotoOptionsWaitUntilArrayItemDomcontentloaded, ScrapeNewParamsGotoOptionsWaitUntilArrayItemNetworkidle0, ScrapeNewParamsGotoOptionsWaitUntilArrayItemNetworkidle2:
		return true
	}
	return false
}

type ScrapeNewParamsRejectResourceType string

const (
	ScrapeNewParamsRejectResourceTypeDocument           ScrapeNewParamsRejectResourceType = "document"
	ScrapeNewParamsRejectResourceTypeStylesheet         ScrapeNewParamsRejectResourceType = "stylesheet"
	ScrapeNewParamsRejectResourceTypeImage              ScrapeNewParamsRejectResourceType = "image"
	ScrapeNewParamsRejectResourceTypeMedia              ScrapeNewParamsRejectResourceType = "media"
	ScrapeNewParamsRejectResourceTypeFont               ScrapeNewParamsRejectResourceType = "font"
	ScrapeNewParamsRejectResourceTypeScript             ScrapeNewParamsRejectResourceType = "script"
	ScrapeNewParamsRejectResourceTypeTexttrack          ScrapeNewParamsRejectResourceType = "texttrack"
	ScrapeNewParamsRejectResourceTypeXHR                ScrapeNewParamsRejectResourceType = "xhr"
	ScrapeNewParamsRejectResourceTypeFetch              ScrapeNewParamsRejectResourceType = "fetch"
	ScrapeNewParamsRejectResourceTypePrefetch           ScrapeNewParamsRejectResourceType = "prefetch"
	ScrapeNewParamsRejectResourceTypeEventsource        ScrapeNewParamsRejectResourceType = "eventsource"
	ScrapeNewParamsRejectResourceTypeWebsocket          ScrapeNewParamsRejectResourceType = "websocket"
	ScrapeNewParamsRejectResourceTypeManifest           ScrapeNewParamsRejectResourceType = "manifest"
	ScrapeNewParamsRejectResourceTypeSignedexchange     ScrapeNewParamsRejectResourceType = "signedexchange"
	ScrapeNewParamsRejectResourceTypePing               ScrapeNewParamsRejectResourceType = "ping"
	ScrapeNewParamsRejectResourceTypeCspviolationreport ScrapeNewParamsRejectResourceType = "cspviolationreport"
	ScrapeNewParamsRejectResourceTypePreflight          ScrapeNewParamsRejectResourceType = "preflight"
	ScrapeNewParamsRejectResourceTypeOther              ScrapeNewParamsRejectResourceType = "other"
)

func (r ScrapeNewParamsRejectResourceType) IsKnown() bool {
	switch r {
	case ScrapeNewParamsRejectResourceTypeDocument, ScrapeNewParamsRejectResourceTypeStylesheet, ScrapeNewParamsRejectResourceTypeImage, ScrapeNewParamsRejectResourceTypeMedia, ScrapeNewParamsRejectResourceTypeFont, ScrapeNewParamsRejectResourceTypeScript, ScrapeNewParamsRejectResourceTypeTexttrack, ScrapeNewParamsRejectResourceTypeXHR, ScrapeNewParamsRejectResourceTypeFetch, ScrapeNewParamsRejectResourceTypePrefetch, ScrapeNewParamsRejectResourceTypeEventsource, ScrapeNewParamsRejectResourceTypeWebsocket, ScrapeNewParamsRejectResourceTypeManifest, ScrapeNewParamsRejectResourceTypeSignedexchange, ScrapeNewParamsRejectResourceTypePing, ScrapeNewParamsRejectResourceTypeCspviolationreport, ScrapeNewParamsRejectResourceTypePreflight, ScrapeNewParamsRejectResourceTypeOther:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
type ScrapeNewParamsViewport struct {
	Height            param.Field[float64] `json:"height,required"`
	Width             param.Field[float64] `json:"width,required"`
	DeviceScaleFactor param.Field[float64] `json:"deviceScaleFactor"`
	HasTouch          param.Field[bool]    `json:"hasTouch"`
	IsLandscape       param.Field[bool]    `json:"isLandscape"`
	IsMobile          param.Field[bool]    `json:"isMobile"`
}

func (r ScrapeNewParamsViewport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Wait for the selector to appear in page. Check
// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
type ScrapeNewParamsWaitForSelector struct {
	Selector param.Field[string]                                `json:"selector,required"`
	Hidden   param.Field[ScrapeNewParamsWaitForSelectorHidden]  `json:"hidden"`
	Timeout  param.Field[float64]                               `json:"timeout"`
	Visible  param.Field[ScrapeNewParamsWaitForSelectorVisible] `json:"visible"`
}

func (r ScrapeNewParamsWaitForSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScrapeNewParamsWaitForSelectorHidden bool

const (
	ScrapeNewParamsWaitForSelectorHiddenTrue ScrapeNewParamsWaitForSelectorHidden = true
)

func (r ScrapeNewParamsWaitForSelectorHidden) IsKnown() bool {
	switch r {
	case ScrapeNewParamsWaitForSelectorHiddenTrue:
		return true
	}
	return false
}

type ScrapeNewParamsWaitForSelectorVisible bool

const (
	ScrapeNewParamsWaitForSelectorVisibleTrue ScrapeNewParamsWaitForSelectorVisible = true
)

func (r ScrapeNewParamsWaitForSelectorVisible) IsKnown() bool {
	switch r {
	case ScrapeNewParamsWaitForSelectorVisibleTrue:
		return true
	}
	return false
}

type ScrapeNewResponseEnvelope struct {
	Result []ScrapeNewResponse `json:"result,required"`
	// Response status
	Status bool                              `json:"status,required"`
	Errors []ScrapeNewResponseEnvelopeErrors `json:"errors"`
	JSON   scrapeNewResponseEnvelopeJSON     `json:"-"`
}

// scrapeNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScrapeNewResponseEnvelope]
type scrapeNewResponseEnvelopeJSON struct {
	Result      apijson.Field
	Status      apijson.Field
	Errors      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScrapeNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scrapeNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ScrapeNewResponseEnvelopeErrors struct {
	// Error code
	Code float64 `json:"code,required"`
	// Error Message
	Message string                              `json:"message,required"`
	JSON    scrapeNewResponseEnvelopeErrorsJSON `json:"-"`
}

// scrapeNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ScrapeNewResponseEnvelopeErrors]
type scrapeNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScrapeNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scrapeNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}
