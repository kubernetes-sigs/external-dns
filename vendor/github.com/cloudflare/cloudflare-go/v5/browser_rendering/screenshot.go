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

// ScreenshotService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScreenshotService] method instead.
type ScreenshotService struct {
	Options []option.RequestOption
}

// NewScreenshotService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewScreenshotService(opts ...option.RequestOption) (r *ScreenshotService) {
	r = &ScreenshotService{}
	r.Options = opts
	return
}

// Takes a screenshot of a webpage from provided URL or HTML. Control page loading
// with `gotoOptions` and `waitFor*` options. Customize screenshots with
// `viewport`, `fullPage`, `clip` and others.
func (r *ScreenshotService) New(ctx context.Context, params ScreenshotNewParams, opts ...option.RequestOption) (res *ScreenshotNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/browser-rendering/screenshot", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type ScreenshotNewResponse struct {
	// Response status
	Status bool                         `json:"status,required"`
	Errors []ScreenshotNewResponseError `json:"errors"`
	JSON   screenshotNewResponseJSON    `json:"-"`
}

// screenshotNewResponseJSON contains the JSON metadata for the struct
// [ScreenshotNewResponse]
type screenshotNewResponseJSON struct {
	Status      apijson.Field
	Errors      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScreenshotNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r screenshotNewResponseJSON) RawJSON() string {
	return r.raw
}

type ScreenshotNewResponseError struct {
	// Error code
	Code float64 `json:"code,required"`
	// Error Message
	Message string                         `json:"message,required"`
	JSON    screenshotNewResponseErrorJSON `json:"-"`
}

// screenshotNewResponseErrorJSON contains the JSON metadata for the struct
// [ScreenshotNewResponseError]
type screenshotNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScreenshotNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r screenshotNewResponseErrorJSON) RawJSON() string {
	return r.raw
}

type ScreenshotNewParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// Cache TTL default is 5s. Set to 0 to disable.
	CacheTTL param.Field[float64] `query:"cacheTTL"`
	// The maximum duration allowed for the browser action to complete after the page
	// has loaded (such as taking screenshots, extracting content, or generating PDFs).
	// If this time limit is exceeded, the action stops and returns a timeout error.
	ActionTimeout param.Field[float64] `json:"actionTimeout"`
	// Adds a `<script>` tag into the page with the desired URL or content.
	AddScriptTag param.Field[[]ScreenshotNewParamsAddScriptTag] `json:"addScriptTag"`
	// Adds a `<link rel="stylesheet">` tag into the page with the desired URL or a
	// `<style type="text/css">` tag with the content.
	AddStyleTag param.Field[[]ScreenshotNewParamsAddStyleTag] `json:"addStyleTag"`
	// Only allow requests that match the provided regex patterns, eg. '/^.\*\.(css)'.
	AllowRequestPattern param.Field[[]string] `json:"allowRequestPattern"`
	// Only allow requests that match the provided resource types, eg. 'image' or
	// 'script'.
	AllowResourceTypes param.Field[[]ScreenshotNewParamsAllowResourceType] `json:"allowResourceTypes"`
	// Provide credentials for HTTP authentication.
	Authenticate param.Field[ScreenshotNewParamsAuthenticate] `json:"authenticate"`
	// Attempt to proceed when 'awaited' events fail or timeout.
	BestAttempt param.Field[bool] `json:"bestAttempt"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setcookie).
	Cookies          param.Field[[]ScreenshotNewParamsCookie] `json:"cookies"`
	EmulateMediaType param.Field[string]                      `json:"emulateMediaType"`
	// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
	GotoOptions param.Field[ScreenshotNewParamsGotoOptions] `json:"gotoOptions"`
	// Set the content of the page, eg: `<h1>Hello World!!</h1>`. Either `html` or
	// `url` must be set.
	HTML param.Field[string] `json:"html"`
	// Block undesired requests that match the provided regex patterns, eg.
	// '/^.\*\.(css)'.
	RejectRequestPattern param.Field[[]string] `json:"rejectRequestPattern"`
	// Block undesired requests that match the provided resource types, eg. 'image' or
	// 'script'.
	RejectResourceTypes param.Field[[]ScreenshotNewParamsRejectResourceType] `json:"rejectResourceTypes"`
	// Check [options](https://pptr.dev/api/puppeteer.screenshotoptions).
	ScreenshotOptions    param.Field[ScreenshotNewParamsScreenshotOptions] `json:"screenshotOptions"`
	ScrollPage           param.Field[bool]                                 `json:"scrollPage"`
	Selector             param.Field[string]                               `json:"selector"`
	SetExtraHTTPHeaders  param.Field[map[string]string]                    `json:"setExtraHTTPHeaders"`
	SetJavaScriptEnabled param.Field[bool]                                 `json:"setJavaScriptEnabled"`
	// URL to navigate to, eg. `https://example.com`.
	URL       param.Field[string] `json:"url" format:"uri"`
	UserAgent param.Field[string] `json:"userAgent"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
	Viewport param.Field[ScreenshotNewParamsViewport] `json:"viewport"`
	// Wait for the selector to appear in page. Check
	// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
	WaitForSelector param.Field[ScreenshotNewParamsWaitForSelector] `json:"waitForSelector"`
	// Waits for a specified timeout before continuing.
	WaitForTimeout param.Field[float64] `json:"waitForTimeout"`
}

func (r ScreenshotNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [ScreenshotNewParams]'s query parameters as `url.Values`.
func (r ScreenshotNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ScreenshotNewParamsAddScriptTag struct {
	ID      param.Field[string] `json:"id"`
	Content param.Field[string] `json:"content"`
	Type    param.Field[string] `json:"type"`
	URL     param.Field[string] `json:"url"`
}

func (r ScreenshotNewParamsAddScriptTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScreenshotNewParamsAddStyleTag struct {
	Content param.Field[string] `json:"content"`
	URL     param.Field[string] `json:"url"`
}

func (r ScreenshotNewParamsAddStyleTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScreenshotNewParamsAllowResourceType string

const (
	ScreenshotNewParamsAllowResourceTypeDocument           ScreenshotNewParamsAllowResourceType = "document"
	ScreenshotNewParamsAllowResourceTypeStylesheet         ScreenshotNewParamsAllowResourceType = "stylesheet"
	ScreenshotNewParamsAllowResourceTypeImage              ScreenshotNewParamsAllowResourceType = "image"
	ScreenshotNewParamsAllowResourceTypeMedia              ScreenshotNewParamsAllowResourceType = "media"
	ScreenshotNewParamsAllowResourceTypeFont               ScreenshotNewParamsAllowResourceType = "font"
	ScreenshotNewParamsAllowResourceTypeScript             ScreenshotNewParamsAllowResourceType = "script"
	ScreenshotNewParamsAllowResourceTypeTexttrack          ScreenshotNewParamsAllowResourceType = "texttrack"
	ScreenshotNewParamsAllowResourceTypeXHR                ScreenshotNewParamsAllowResourceType = "xhr"
	ScreenshotNewParamsAllowResourceTypeFetch              ScreenshotNewParamsAllowResourceType = "fetch"
	ScreenshotNewParamsAllowResourceTypePrefetch           ScreenshotNewParamsAllowResourceType = "prefetch"
	ScreenshotNewParamsAllowResourceTypeEventsource        ScreenshotNewParamsAllowResourceType = "eventsource"
	ScreenshotNewParamsAllowResourceTypeWebsocket          ScreenshotNewParamsAllowResourceType = "websocket"
	ScreenshotNewParamsAllowResourceTypeManifest           ScreenshotNewParamsAllowResourceType = "manifest"
	ScreenshotNewParamsAllowResourceTypeSignedexchange     ScreenshotNewParamsAllowResourceType = "signedexchange"
	ScreenshotNewParamsAllowResourceTypePing               ScreenshotNewParamsAllowResourceType = "ping"
	ScreenshotNewParamsAllowResourceTypeCspviolationreport ScreenshotNewParamsAllowResourceType = "cspviolationreport"
	ScreenshotNewParamsAllowResourceTypePreflight          ScreenshotNewParamsAllowResourceType = "preflight"
	ScreenshotNewParamsAllowResourceTypeOther              ScreenshotNewParamsAllowResourceType = "other"
)

func (r ScreenshotNewParamsAllowResourceType) IsKnown() bool {
	switch r {
	case ScreenshotNewParamsAllowResourceTypeDocument, ScreenshotNewParamsAllowResourceTypeStylesheet, ScreenshotNewParamsAllowResourceTypeImage, ScreenshotNewParamsAllowResourceTypeMedia, ScreenshotNewParamsAllowResourceTypeFont, ScreenshotNewParamsAllowResourceTypeScript, ScreenshotNewParamsAllowResourceTypeTexttrack, ScreenshotNewParamsAllowResourceTypeXHR, ScreenshotNewParamsAllowResourceTypeFetch, ScreenshotNewParamsAllowResourceTypePrefetch, ScreenshotNewParamsAllowResourceTypeEventsource, ScreenshotNewParamsAllowResourceTypeWebsocket, ScreenshotNewParamsAllowResourceTypeManifest, ScreenshotNewParamsAllowResourceTypeSignedexchange, ScreenshotNewParamsAllowResourceTypePing, ScreenshotNewParamsAllowResourceTypeCspviolationreport, ScreenshotNewParamsAllowResourceTypePreflight, ScreenshotNewParamsAllowResourceTypeOther:
		return true
	}
	return false
}

// Provide credentials for HTTP authentication.
type ScreenshotNewParamsAuthenticate struct {
	Password param.Field[string] `json:"password,required"`
	Username param.Field[string] `json:"username,required"`
}

func (r ScreenshotNewParamsAuthenticate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScreenshotNewParamsCookie struct {
	Name         param.Field[string]                                 `json:"name,required"`
	Value        param.Field[string]                                 `json:"value,required"`
	Domain       param.Field[string]                                 `json:"domain"`
	Expires      param.Field[float64]                                `json:"expires"`
	HTTPOnly     param.Field[bool]                                   `json:"httpOnly"`
	PartitionKey param.Field[string]                                 `json:"partitionKey"`
	Path         param.Field[string]                                 `json:"path"`
	Priority     param.Field[ScreenshotNewParamsCookiesPriority]     `json:"priority"`
	SameParty    param.Field[bool]                                   `json:"sameParty"`
	SameSite     param.Field[ScreenshotNewParamsCookiesSameSite]     `json:"sameSite"`
	Secure       param.Field[bool]                                   `json:"secure"`
	SourcePort   param.Field[float64]                                `json:"sourcePort"`
	SourceScheme param.Field[ScreenshotNewParamsCookiesSourceScheme] `json:"sourceScheme"`
	URL          param.Field[string]                                 `json:"url"`
}

func (r ScreenshotNewParamsCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScreenshotNewParamsCookiesPriority string

const (
	ScreenshotNewParamsCookiesPriorityLow    ScreenshotNewParamsCookiesPriority = "Low"
	ScreenshotNewParamsCookiesPriorityMedium ScreenshotNewParamsCookiesPriority = "Medium"
	ScreenshotNewParamsCookiesPriorityHigh   ScreenshotNewParamsCookiesPriority = "High"
)

func (r ScreenshotNewParamsCookiesPriority) IsKnown() bool {
	switch r {
	case ScreenshotNewParamsCookiesPriorityLow, ScreenshotNewParamsCookiesPriorityMedium, ScreenshotNewParamsCookiesPriorityHigh:
		return true
	}
	return false
}

type ScreenshotNewParamsCookiesSameSite string

const (
	ScreenshotNewParamsCookiesSameSiteStrict ScreenshotNewParamsCookiesSameSite = "Strict"
	ScreenshotNewParamsCookiesSameSiteLax    ScreenshotNewParamsCookiesSameSite = "Lax"
	ScreenshotNewParamsCookiesSameSiteNone   ScreenshotNewParamsCookiesSameSite = "None"
)

func (r ScreenshotNewParamsCookiesSameSite) IsKnown() bool {
	switch r {
	case ScreenshotNewParamsCookiesSameSiteStrict, ScreenshotNewParamsCookiesSameSiteLax, ScreenshotNewParamsCookiesSameSiteNone:
		return true
	}
	return false
}

type ScreenshotNewParamsCookiesSourceScheme string

const (
	ScreenshotNewParamsCookiesSourceSchemeUnset     ScreenshotNewParamsCookiesSourceScheme = "Unset"
	ScreenshotNewParamsCookiesSourceSchemeNonSecure ScreenshotNewParamsCookiesSourceScheme = "NonSecure"
	ScreenshotNewParamsCookiesSourceSchemeSecure    ScreenshotNewParamsCookiesSourceScheme = "Secure"
)

func (r ScreenshotNewParamsCookiesSourceScheme) IsKnown() bool {
	switch r {
	case ScreenshotNewParamsCookiesSourceSchemeUnset, ScreenshotNewParamsCookiesSourceSchemeNonSecure, ScreenshotNewParamsCookiesSourceSchemeSecure:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
type ScreenshotNewParamsGotoOptions struct {
	Referer        param.Field[string]                                       `json:"referer"`
	ReferrerPolicy param.Field[string]                                       `json:"referrerPolicy"`
	Timeout        param.Field[float64]                                      `json:"timeout"`
	WaitUntil      param.Field[ScreenshotNewParamsGotoOptionsWaitUntilUnion] `json:"waitUntil"`
}

func (r ScreenshotNewParamsGotoOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [browser_rendering.ScreenshotNewParamsGotoOptionsWaitUntilString],
// [browser_rendering.ScreenshotNewParamsGotoOptionsWaitUntilArray].
type ScreenshotNewParamsGotoOptionsWaitUntilUnion interface {
	implementsScreenshotNewParamsGotoOptionsWaitUntilUnion()
}

type ScreenshotNewParamsGotoOptionsWaitUntilString string

const (
	ScreenshotNewParamsGotoOptionsWaitUntilStringLoad             ScreenshotNewParamsGotoOptionsWaitUntilString = "load"
	ScreenshotNewParamsGotoOptionsWaitUntilStringDomcontentloaded ScreenshotNewParamsGotoOptionsWaitUntilString = "domcontentloaded"
	ScreenshotNewParamsGotoOptionsWaitUntilStringNetworkidle0     ScreenshotNewParamsGotoOptionsWaitUntilString = "networkidle0"
	ScreenshotNewParamsGotoOptionsWaitUntilStringNetworkidle2     ScreenshotNewParamsGotoOptionsWaitUntilString = "networkidle2"
)

func (r ScreenshotNewParamsGotoOptionsWaitUntilString) IsKnown() bool {
	switch r {
	case ScreenshotNewParamsGotoOptionsWaitUntilStringLoad, ScreenshotNewParamsGotoOptionsWaitUntilStringDomcontentloaded, ScreenshotNewParamsGotoOptionsWaitUntilStringNetworkidle0, ScreenshotNewParamsGotoOptionsWaitUntilStringNetworkidle2:
		return true
	}
	return false
}

func (r ScreenshotNewParamsGotoOptionsWaitUntilString) implementsScreenshotNewParamsGotoOptionsWaitUntilUnion() {
}

type ScreenshotNewParamsGotoOptionsWaitUntilArray []ScreenshotNewParamsGotoOptionsWaitUntilArrayItem

func (r ScreenshotNewParamsGotoOptionsWaitUntilArray) implementsScreenshotNewParamsGotoOptionsWaitUntilUnion() {
}

type ScreenshotNewParamsGotoOptionsWaitUntilArrayItem string

const (
	ScreenshotNewParamsGotoOptionsWaitUntilArrayItemLoad             ScreenshotNewParamsGotoOptionsWaitUntilArrayItem = "load"
	ScreenshotNewParamsGotoOptionsWaitUntilArrayItemDomcontentloaded ScreenshotNewParamsGotoOptionsWaitUntilArrayItem = "domcontentloaded"
	ScreenshotNewParamsGotoOptionsWaitUntilArrayItemNetworkidle0     ScreenshotNewParamsGotoOptionsWaitUntilArrayItem = "networkidle0"
	ScreenshotNewParamsGotoOptionsWaitUntilArrayItemNetworkidle2     ScreenshotNewParamsGotoOptionsWaitUntilArrayItem = "networkidle2"
)

func (r ScreenshotNewParamsGotoOptionsWaitUntilArrayItem) IsKnown() bool {
	switch r {
	case ScreenshotNewParamsGotoOptionsWaitUntilArrayItemLoad, ScreenshotNewParamsGotoOptionsWaitUntilArrayItemDomcontentloaded, ScreenshotNewParamsGotoOptionsWaitUntilArrayItemNetworkidle0, ScreenshotNewParamsGotoOptionsWaitUntilArrayItemNetworkidle2:
		return true
	}
	return false
}

type ScreenshotNewParamsRejectResourceType string

const (
	ScreenshotNewParamsRejectResourceTypeDocument           ScreenshotNewParamsRejectResourceType = "document"
	ScreenshotNewParamsRejectResourceTypeStylesheet         ScreenshotNewParamsRejectResourceType = "stylesheet"
	ScreenshotNewParamsRejectResourceTypeImage              ScreenshotNewParamsRejectResourceType = "image"
	ScreenshotNewParamsRejectResourceTypeMedia              ScreenshotNewParamsRejectResourceType = "media"
	ScreenshotNewParamsRejectResourceTypeFont               ScreenshotNewParamsRejectResourceType = "font"
	ScreenshotNewParamsRejectResourceTypeScript             ScreenshotNewParamsRejectResourceType = "script"
	ScreenshotNewParamsRejectResourceTypeTexttrack          ScreenshotNewParamsRejectResourceType = "texttrack"
	ScreenshotNewParamsRejectResourceTypeXHR                ScreenshotNewParamsRejectResourceType = "xhr"
	ScreenshotNewParamsRejectResourceTypeFetch              ScreenshotNewParamsRejectResourceType = "fetch"
	ScreenshotNewParamsRejectResourceTypePrefetch           ScreenshotNewParamsRejectResourceType = "prefetch"
	ScreenshotNewParamsRejectResourceTypeEventsource        ScreenshotNewParamsRejectResourceType = "eventsource"
	ScreenshotNewParamsRejectResourceTypeWebsocket          ScreenshotNewParamsRejectResourceType = "websocket"
	ScreenshotNewParamsRejectResourceTypeManifest           ScreenshotNewParamsRejectResourceType = "manifest"
	ScreenshotNewParamsRejectResourceTypeSignedexchange     ScreenshotNewParamsRejectResourceType = "signedexchange"
	ScreenshotNewParamsRejectResourceTypePing               ScreenshotNewParamsRejectResourceType = "ping"
	ScreenshotNewParamsRejectResourceTypeCspviolationreport ScreenshotNewParamsRejectResourceType = "cspviolationreport"
	ScreenshotNewParamsRejectResourceTypePreflight          ScreenshotNewParamsRejectResourceType = "preflight"
	ScreenshotNewParamsRejectResourceTypeOther              ScreenshotNewParamsRejectResourceType = "other"
)

func (r ScreenshotNewParamsRejectResourceType) IsKnown() bool {
	switch r {
	case ScreenshotNewParamsRejectResourceTypeDocument, ScreenshotNewParamsRejectResourceTypeStylesheet, ScreenshotNewParamsRejectResourceTypeImage, ScreenshotNewParamsRejectResourceTypeMedia, ScreenshotNewParamsRejectResourceTypeFont, ScreenshotNewParamsRejectResourceTypeScript, ScreenshotNewParamsRejectResourceTypeTexttrack, ScreenshotNewParamsRejectResourceTypeXHR, ScreenshotNewParamsRejectResourceTypeFetch, ScreenshotNewParamsRejectResourceTypePrefetch, ScreenshotNewParamsRejectResourceTypeEventsource, ScreenshotNewParamsRejectResourceTypeWebsocket, ScreenshotNewParamsRejectResourceTypeManifest, ScreenshotNewParamsRejectResourceTypeSignedexchange, ScreenshotNewParamsRejectResourceTypePing, ScreenshotNewParamsRejectResourceTypeCspviolationreport, ScreenshotNewParamsRejectResourceTypePreflight, ScreenshotNewParamsRejectResourceTypeOther:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.screenshotoptions).
type ScreenshotNewParamsScreenshotOptions struct {
	CaptureBeyondViewport param.Field[bool]                                         `json:"captureBeyondViewport"`
	Clip                  param.Field[ScreenshotNewParamsScreenshotOptionsClip]     `json:"clip"`
	Encoding              param.Field[ScreenshotNewParamsScreenshotOptionsEncoding] `json:"encoding"`
	FromSurface           param.Field[bool]                                         `json:"fromSurface"`
	FullPage              param.Field[bool]                                         `json:"fullPage"`
	OmitBackground        param.Field[bool]                                         `json:"omitBackground"`
	OptimizeForSpeed      param.Field[bool]                                         `json:"optimizeForSpeed"`
	Quality               param.Field[float64]                                      `json:"quality"`
	Type                  param.Field[ScreenshotNewParamsScreenshotOptionsType]     `json:"type"`
}

func (r ScreenshotNewParamsScreenshotOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScreenshotNewParamsScreenshotOptionsClip struct {
	Height param.Field[float64] `json:"height,required"`
	Width  param.Field[float64] `json:"width,required"`
	X      param.Field[float64] `json:"x,required"`
	Y      param.Field[float64] `json:"y,required"`
	Scale  param.Field[float64] `json:"scale"`
}

func (r ScreenshotNewParamsScreenshotOptionsClip) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScreenshotNewParamsScreenshotOptionsEncoding string

const (
	ScreenshotNewParamsScreenshotOptionsEncodingBinary ScreenshotNewParamsScreenshotOptionsEncoding = "binary"
	ScreenshotNewParamsScreenshotOptionsEncodingBase64 ScreenshotNewParamsScreenshotOptionsEncoding = "base64"
)

func (r ScreenshotNewParamsScreenshotOptionsEncoding) IsKnown() bool {
	switch r {
	case ScreenshotNewParamsScreenshotOptionsEncodingBinary, ScreenshotNewParamsScreenshotOptionsEncodingBase64:
		return true
	}
	return false
}

type ScreenshotNewParamsScreenshotOptionsType string

const (
	ScreenshotNewParamsScreenshotOptionsTypePNG  ScreenshotNewParamsScreenshotOptionsType = "png"
	ScreenshotNewParamsScreenshotOptionsTypeJPEG ScreenshotNewParamsScreenshotOptionsType = "jpeg"
	ScreenshotNewParamsScreenshotOptionsTypeWebP ScreenshotNewParamsScreenshotOptionsType = "webp"
)

func (r ScreenshotNewParamsScreenshotOptionsType) IsKnown() bool {
	switch r {
	case ScreenshotNewParamsScreenshotOptionsTypePNG, ScreenshotNewParamsScreenshotOptionsTypeJPEG, ScreenshotNewParamsScreenshotOptionsTypeWebP:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
type ScreenshotNewParamsViewport struct {
	Height            param.Field[float64] `json:"height,required"`
	Width             param.Field[float64] `json:"width,required"`
	DeviceScaleFactor param.Field[float64] `json:"deviceScaleFactor"`
	HasTouch          param.Field[bool]    `json:"hasTouch"`
	IsLandscape       param.Field[bool]    `json:"isLandscape"`
	IsMobile          param.Field[bool]    `json:"isMobile"`
}

func (r ScreenshotNewParamsViewport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Wait for the selector to appear in page. Check
// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
type ScreenshotNewParamsWaitForSelector struct {
	Selector param.Field[string]                                    `json:"selector,required"`
	Hidden   param.Field[ScreenshotNewParamsWaitForSelectorHidden]  `json:"hidden"`
	Timeout  param.Field[float64]                                   `json:"timeout"`
	Visible  param.Field[ScreenshotNewParamsWaitForSelectorVisible] `json:"visible"`
}

func (r ScreenshotNewParamsWaitForSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScreenshotNewParamsWaitForSelectorHidden bool

const (
	ScreenshotNewParamsWaitForSelectorHiddenTrue ScreenshotNewParamsWaitForSelectorHidden = true
)

func (r ScreenshotNewParamsWaitForSelectorHidden) IsKnown() bool {
	switch r {
	case ScreenshotNewParamsWaitForSelectorHiddenTrue:
		return true
	}
	return false
}

type ScreenshotNewParamsWaitForSelectorVisible bool

const (
	ScreenshotNewParamsWaitForSelectorVisibleTrue ScreenshotNewParamsWaitForSelectorVisible = true
)

func (r ScreenshotNewParamsWaitForSelectorVisible) IsKnown() bool {
	switch r {
	case ScreenshotNewParamsWaitForSelectorVisibleTrue:
		return true
	}
	return false
}
