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

// ContentService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewContentService] method instead.
type ContentService struct {
	Options []option.RequestOption
}

// NewContentService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewContentService(opts ...option.RequestOption) (r *ContentService) {
	r = &ContentService{}
	r.Options = opts
	return
}

// Fetches rendered HTML content from provided URL or HTML. Check available options
// like `gotoOptions` and `waitFor*` to control page load behaviour.
func (r *ContentService) New(ctx context.Context, params ContentNewParams, opts ...option.RequestOption) (res *string, err error) {
	var env ContentNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/browser-rendering/content", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ContentNewParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// Cache TTL default is 5s. Set to 0 to disable.
	CacheTTL param.Field[float64] `query:"cacheTTL"`
	// The maximum duration allowed for the browser action to complete after the page
	// has loaded (such as taking screenshots, extracting content, or generating PDFs).
	// If this time limit is exceeded, the action stops and returns a timeout error.
	ActionTimeout param.Field[float64] `json:"actionTimeout"`
	// Adds a `<script>` tag into the page with the desired URL or content.
	AddScriptTag param.Field[[]ContentNewParamsAddScriptTag] `json:"addScriptTag"`
	// Adds a `<link rel="stylesheet">` tag into the page with the desired URL or a
	// `<style type="text/css">` tag with the content.
	AddStyleTag param.Field[[]ContentNewParamsAddStyleTag] `json:"addStyleTag"`
	// Only allow requests that match the provided regex patterns, eg. '/^.\*\.(css)'.
	AllowRequestPattern param.Field[[]string] `json:"allowRequestPattern"`
	// Only allow requests that match the provided resource types, eg. 'image' or
	// 'script'.
	AllowResourceTypes param.Field[[]ContentNewParamsAllowResourceType] `json:"allowResourceTypes"`
	// Provide credentials for HTTP authentication.
	Authenticate param.Field[ContentNewParamsAuthenticate] `json:"authenticate"`
	// Attempt to proceed when 'awaited' events fail or timeout.
	BestAttempt param.Field[bool] `json:"bestAttempt"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setcookie).
	Cookies          param.Field[[]ContentNewParamsCookie] `json:"cookies"`
	EmulateMediaType param.Field[string]                   `json:"emulateMediaType"`
	// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
	GotoOptions param.Field[ContentNewParamsGotoOptions] `json:"gotoOptions"`
	// Set the content of the page, eg: `<h1>Hello World!!</h1>`. Either `html` or
	// `url` must be set.
	HTML param.Field[string] `json:"html"`
	// Block undesired requests that match the provided regex patterns, eg.
	// '/^.\*\.(css)'.
	RejectRequestPattern param.Field[[]string] `json:"rejectRequestPattern"`
	// Block undesired requests that match the provided resource types, eg. 'image' or
	// 'script'.
	RejectResourceTypes  param.Field[[]ContentNewParamsRejectResourceType] `json:"rejectResourceTypes"`
	SetExtraHTTPHeaders  param.Field[map[string]string]                    `json:"setExtraHTTPHeaders"`
	SetJavaScriptEnabled param.Field[bool]                                 `json:"setJavaScriptEnabled"`
	// URL to navigate to, eg. `https://example.com`.
	URL       param.Field[string] `json:"url" format:"uri"`
	UserAgent param.Field[string] `json:"userAgent"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
	Viewport param.Field[ContentNewParamsViewport] `json:"viewport"`
	// Wait for the selector to appear in page. Check
	// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
	WaitForSelector param.Field[ContentNewParamsWaitForSelector] `json:"waitForSelector"`
	// Waits for a specified timeout before continuing.
	WaitForTimeout param.Field[float64] `json:"waitForTimeout"`
}

func (r ContentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [ContentNewParams]'s query parameters as `url.Values`.
func (r ContentNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ContentNewParamsAddScriptTag struct {
	ID      param.Field[string] `json:"id"`
	Content param.Field[string] `json:"content"`
	Type    param.Field[string] `json:"type"`
	URL     param.Field[string] `json:"url"`
}

func (r ContentNewParamsAddScriptTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContentNewParamsAddStyleTag struct {
	Content param.Field[string] `json:"content"`
	URL     param.Field[string] `json:"url"`
}

func (r ContentNewParamsAddStyleTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContentNewParamsAllowResourceType string

const (
	ContentNewParamsAllowResourceTypeDocument           ContentNewParamsAllowResourceType = "document"
	ContentNewParamsAllowResourceTypeStylesheet         ContentNewParamsAllowResourceType = "stylesheet"
	ContentNewParamsAllowResourceTypeImage              ContentNewParamsAllowResourceType = "image"
	ContentNewParamsAllowResourceTypeMedia              ContentNewParamsAllowResourceType = "media"
	ContentNewParamsAllowResourceTypeFont               ContentNewParamsAllowResourceType = "font"
	ContentNewParamsAllowResourceTypeScript             ContentNewParamsAllowResourceType = "script"
	ContentNewParamsAllowResourceTypeTexttrack          ContentNewParamsAllowResourceType = "texttrack"
	ContentNewParamsAllowResourceTypeXHR                ContentNewParamsAllowResourceType = "xhr"
	ContentNewParamsAllowResourceTypeFetch              ContentNewParamsAllowResourceType = "fetch"
	ContentNewParamsAllowResourceTypePrefetch           ContentNewParamsAllowResourceType = "prefetch"
	ContentNewParamsAllowResourceTypeEventsource        ContentNewParamsAllowResourceType = "eventsource"
	ContentNewParamsAllowResourceTypeWebsocket          ContentNewParamsAllowResourceType = "websocket"
	ContentNewParamsAllowResourceTypeManifest           ContentNewParamsAllowResourceType = "manifest"
	ContentNewParamsAllowResourceTypeSignedexchange     ContentNewParamsAllowResourceType = "signedexchange"
	ContentNewParamsAllowResourceTypePing               ContentNewParamsAllowResourceType = "ping"
	ContentNewParamsAllowResourceTypeCspviolationreport ContentNewParamsAllowResourceType = "cspviolationreport"
	ContentNewParamsAllowResourceTypePreflight          ContentNewParamsAllowResourceType = "preflight"
	ContentNewParamsAllowResourceTypeOther              ContentNewParamsAllowResourceType = "other"
)

func (r ContentNewParamsAllowResourceType) IsKnown() bool {
	switch r {
	case ContentNewParamsAllowResourceTypeDocument, ContentNewParamsAllowResourceTypeStylesheet, ContentNewParamsAllowResourceTypeImage, ContentNewParamsAllowResourceTypeMedia, ContentNewParamsAllowResourceTypeFont, ContentNewParamsAllowResourceTypeScript, ContentNewParamsAllowResourceTypeTexttrack, ContentNewParamsAllowResourceTypeXHR, ContentNewParamsAllowResourceTypeFetch, ContentNewParamsAllowResourceTypePrefetch, ContentNewParamsAllowResourceTypeEventsource, ContentNewParamsAllowResourceTypeWebsocket, ContentNewParamsAllowResourceTypeManifest, ContentNewParamsAllowResourceTypeSignedexchange, ContentNewParamsAllowResourceTypePing, ContentNewParamsAllowResourceTypeCspviolationreport, ContentNewParamsAllowResourceTypePreflight, ContentNewParamsAllowResourceTypeOther:
		return true
	}
	return false
}

// Provide credentials for HTTP authentication.
type ContentNewParamsAuthenticate struct {
	Password param.Field[string] `json:"password,required"`
	Username param.Field[string] `json:"username,required"`
}

func (r ContentNewParamsAuthenticate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContentNewParamsCookie struct {
	Name         param.Field[string]                              `json:"name,required"`
	Value        param.Field[string]                              `json:"value,required"`
	Domain       param.Field[string]                              `json:"domain"`
	Expires      param.Field[float64]                             `json:"expires"`
	HTTPOnly     param.Field[bool]                                `json:"httpOnly"`
	PartitionKey param.Field[string]                              `json:"partitionKey"`
	Path         param.Field[string]                              `json:"path"`
	Priority     param.Field[ContentNewParamsCookiesPriority]     `json:"priority"`
	SameParty    param.Field[bool]                                `json:"sameParty"`
	SameSite     param.Field[ContentNewParamsCookiesSameSite]     `json:"sameSite"`
	Secure       param.Field[bool]                                `json:"secure"`
	SourcePort   param.Field[float64]                             `json:"sourcePort"`
	SourceScheme param.Field[ContentNewParamsCookiesSourceScheme] `json:"sourceScheme"`
	URL          param.Field[string]                              `json:"url"`
}

func (r ContentNewParamsCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContentNewParamsCookiesPriority string

const (
	ContentNewParamsCookiesPriorityLow    ContentNewParamsCookiesPriority = "Low"
	ContentNewParamsCookiesPriorityMedium ContentNewParamsCookiesPriority = "Medium"
	ContentNewParamsCookiesPriorityHigh   ContentNewParamsCookiesPriority = "High"
)

func (r ContentNewParamsCookiesPriority) IsKnown() bool {
	switch r {
	case ContentNewParamsCookiesPriorityLow, ContentNewParamsCookiesPriorityMedium, ContentNewParamsCookiesPriorityHigh:
		return true
	}
	return false
}

type ContentNewParamsCookiesSameSite string

const (
	ContentNewParamsCookiesSameSiteStrict ContentNewParamsCookiesSameSite = "Strict"
	ContentNewParamsCookiesSameSiteLax    ContentNewParamsCookiesSameSite = "Lax"
	ContentNewParamsCookiesSameSiteNone   ContentNewParamsCookiesSameSite = "None"
)

func (r ContentNewParamsCookiesSameSite) IsKnown() bool {
	switch r {
	case ContentNewParamsCookiesSameSiteStrict, ContentNewParamsCookiesSameSiteLax, ContentNewParamsCookiesSameSiteNone:
		return true
	}
	return false
}

type ContentNewParamsCookiesSourceScheme string

const (
	ContentNewParamsCookiesSourceSchemeUnset     ContentNewParamsCookiesSourceScheme = "Unset"
	ContentNewParamsCookiesSourceSchemeNonSecure ContentNewParamsCookiesSourceScheme = "NonSecure"
	ContentNewParamsCookiesSourceSchemeSecure    ContentNewParamsCookiesSourceScheme = "Secure"
)

func (r ContentNewParamsCookiesSourceScheme) IsKnown() bool {
	switch r {
	case ContentNewParamsCookiesSourceSchemeUnset, ContentNewParamsCookiesSourceSchemeNonSecure, ContentNewParamsCookiesSourceSchemeSecure:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
type ContentNewParamsGotoOptions struct {
	Referer        param.Field[string]                                    `json:"referer"`
	ReferrerPolicy param.Field[string]                                    `json:"referrerPolicy"`
	Timeout        param.Field[float64]                                   `json:"timeout"`
	WaitUntil      param.Field[ContentNewParamsGotoOptionsWaitUntilUnion] `json:"waitUntil"`
}

func (r ContentNewParamsGotoOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [browser_rendering.ContentNewParamsGotoOptionsWaitUntilString],
// [browser_rendering.ContentNewParamsGotoOptionsWaitUntilArray].
type ContentNewParamsGotoOptionsWaitUntilUnion interface {
	implementsContentNewParamsGotoOptionsWaitUntilUnion()
}

type ContentNewParamsGotoOptionsWaitUntilString string

const (
	ContentNewParamsGotoOptionsWaitUntilStringLoad             ContentNewParamsGotoOptionsWaitUntilString = "load"
	ContentNewParamsGotoOptionsWaitUntilStringDomcontentloaded ContentNewParamsGotoOptionsWaitUntilString = "domcontentloaded"
	ContentNewParamsGotoOptionsWaitUntilStringNetworkidle0     ContentNewParamsGotoOptionsWaitUntilString = "networkidle0"
	ContentNewParamsGotoOptionsWaitUntilStringNetworkidle2     ContentNewParamsGotoOptionsWaitUntilString = "networkidle2"
)

func (r ContentNewParamsGotoOptionsWaitUntilString) IsKnown() bool {
	switch r {
	case ContentNewParamsGotoOptionsWaitUntilStringLoad, ContentNewParamsGotoOptionsWaitUntilStringDomcontentloaded, ContentNewParamsGotoOptionsWaitUntilStringNetworkidle0, ContentNewParamsGotoOptionsWaitUntilStringNetworkidle2:
		return true
	}
	return false
}

func (r ContentNewParamsGotoOptionsWaitUntilString) implementsContentNewParamsGotoOptionsWaitUntilUnion() {
}

type ContentNewParamsGotoOptionsWaitUntilArray []ContentNewParamsGotoOptionsWaitUntilArrayItem

func (r ContentNewParamsGotoOptionsWaitUntilArray) implementsContentNewParamsGotoOptionsWaitUntilUnion() {
}

type ContentNewParamsGotoOptionsWaitUntilArrayItem string

const (
	ContentNewParamsGotoOptionsWaitUntilArrayItemLoad             ContentNewParamsGotoOptionsWaitUntilArrayItem = "load"
	ContentNewParamsGotoOptionsWaitUntilArrayItemDomcontentloaded ContentNewParamsGotoOptionsWaitUntilArrayItem = "domcontentloaded"
	ContentNewParamsGotoOptionsWaitUntilArrayItemNetworkidle0     ContentNewParamsGotoOptionsWaitUntilArrayItem = "networkidle0"
	ContentNewParamsGotoOptionsWaitUntilArrayItemNetworkidle2     ContentNewParamsGotoOptionsWaitUntilArrayItem = "networkidle2"
)

func (r ContentNewParamsGotoOptionsWaitUntilArrayItem) IsKnown() bool {
	switch r {
	case ContentNewParamsGotoOptionsWaitUntilArrayItemLoad, ContentNewParamsGotoOptionsWaitUntilArrayItemDomcontentloaded, ContentNewParamsGotoOptionsWaitUntilArrayItemNetworkidle0, ContentNewParamsGotoOptionsWaitUntilArrayItemNetworkidle2:
		return true
	}
	return false
}

type ContentNewParamsRejectResourceType string

const (
	ContentNewParamsRejectResourceTypeDocument           ContentNewParamsRejectResourceType = "document"
	ContentNewParamsRejectResourceTypeStylesheet         ContentNewParamsRejectResourceType = "stylesheet"
	ContentNewParamsRejectResourceTypeImage              ContentNewParamsRejectResourceType = "image"
	ContentNewParamsRejectResourceTypeMedia              ContentNewParamsRejectResourceType = "media"
	ContentNewParamsRejectResourceTypeFont               ContentNewParamsRejectResourceType = "font"
	ContentNewParamsRejectResourceTypeScript             ContentNewParamsRejectResourceType = "script"
	ContentNewParamsRejectResourceTypeTexttrack          ContentNewParamsRejectResourceType = "texttrack"
	ContentNewParamsRejectResourceTypeXHR                ContentNewParamsRejectResourceType = "xhr"
	ContentNewParamsRejectResourceTypeFetch              ContentNewParamsRejectResourceType = "fetch"
	ContentNewParamsRejectResourceTypePrefetch           ContentNewParamsRejectResourceType = "prefetch"
	ContentNewParamsRejectResourceTypeEventsource        ContentNewParamsRejectResourceType = "eventsource"
	ContentNewParamsRejectResourceTypeWebsocket          ContentNewParamsRejectResourceType = "websocket"
	ContentNewParamsRejectResourceTypeManifest           ContentNewParamsRejectResourceType = "manifest"
	ContentNewParamsRejectResourceTypeSignedexchange     ContentNewParamsRejectResourceType = "signedexchange"
	ContentNewParamsRejectResourceTypePing               ContentNewParamsRejectResourceType = "ping"
	ContentNewParamsRejectResourceTypeCspviolationreport ContentNewParamsRejectResourceType = "cspviolationreport"
	ContentNewParamsRejectResourceTypePreflight          ContentNewParamsRejectResourceType = "preflight"
	ContentNewParamsRejectResourceTypeOther              ContentNewParamsRejectResourceType = "other"
)

func (r ContentNewParamsRejectResourceType) IsKnown() bool {
	switch r {
	case ContentNewParamsRejectResourceTypeDocument, ContentNewParamsRejectResourceTypeStylesheet, ContentNewParamsRejectResourceTypeImage, ContentNewParamsRejectResourceTypeMedia, ContentNewParamsRejectResourceTypeFont, ContentNewParamsRejectResourceTypeScript, ContentNewParamsRejectResourceTypeTexttrack, ContentNewParamsRejectResourceTypeXHR, ContentNewParamsRejectResourceTypeFetch, ContentNewParamsRejectResourceTypePrefetch, ContentNewParamsRejectResourceTypeEventsource, ContentNewParamsRejectResourceTypeWebsocket, ContentNewParamsRejectResourceTypeManifest, ContentNewParamsRejectResourceTypeSignedexchange, ContentNewParamsRejectResourceTypePing, ContentNewParamsRejectResourceTypeCspviolationreport, ContentNewParamsRejectResourceTypePreflight, ContentNewParamsRejectResourceTypeOther:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
type ContentNewParamsViewport struct {
	Height            param.Field[float64] `json:"height,required"`
	Width             param.Field[float64] `json:"width,required"`
	DeviceScaleFactor param.Field[float64] `json:"deviceScaleFactor"`
	HasTouch          param.Field[bool]    `json:"hasTouch"`
	IsLandscape       param.Field[bool]    `json:"isLandscape"`
	IsMobile          param.Field[bool]    `json:"isMobile"`
}

func (r ContentNewParamsViewport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Wait for the selector to appear in page. Check
// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
type ContentNewParamsWaitForSelector struct {
	Selector param.Field[string]                                 `json:"selector,required"`
	Hidden   param.Field[ContentNewParamsWaitForSelectorHidden]  `json:"hidden"`
	Timeout  param.Field[float64]                                `json:"timeout"`
	Visible  param.Field[ContentNewParamsWaitForSelectorVisible] `json:"visible"`
}

func (r ContentNewParamsWaitForSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContentNewParamsWaitForSelectorHidden bool

const (
	ContentNewParamsWaitForSelectorHiddenTrue ContentNewParamsWaitForSelectorHidden = true
)

func (r ContentNewParamsWaitForSelectorHidden) IsKnown() bool {
	switch r {
	case ContentNewParamsWaitForSelectorHiddenTrue:
		return true
	}
	return false
}

type ContentNewParamsWaitForSelectorVisible bool

const (
	ContentNewParamsWaitForSelectorVisibleTrue ContentNewParamsWaitForSelectorVisible = true
)

func (r ContentNewParamsWaitForSelectorVisible) IsKnown() bool {
	switch r {
	case ContentNewParamsWaitForSelectorVisibleTrue:
		return true
	}
	return false
}

type ContentNewResponseEnvelope struct {
	Meta ContentNewResponseEnvelopeMeta `json:"meta,required"`
	// Response status
	Status bool                               `json:"status,required"`
	Errors []ContentNewResponseEnvelopeErrors `json:"errors"`
	// HTML content
	Result string                         `json:"result"`
	JSON   contentNewResponseEnvelopeJSON `json:"-"`
}

// contentNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ContentNewResponseEnvelope]
type contentNewResponseEnvelopeJSON struct {
	Meta        apijson.Field
	Status      apijson.Field
	Errors      apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContentNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contentNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ContentNewResponseEnvelopeMeta struct {
	Status float64                            `json:"status,required"`
	Title  string                             `json:"title,required"`
	JSON   contentNewResponseEnvelopeMetaJSON `json:"-"`
}

// contentNewResponseEnvelopeMetaJSON contains the JSON metadata for the struct
// [ContentNewResponseEnvelopeMeta]
type contentNewResponseEnvelopeMetaJSON struct {
	Status      apijson.Field
	Title       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContentNewResponseEnvelopeMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contentNewResponseEnvelopeMetaJSON) RawJSON() string {
	return r.raw
}

type ContentNewResponseEnvelopeErrors struct {
	// Error code
	Code float64 `json:"code,required"`
	// Error Message
	Message string                               `json:"message,required"`
	JSON    contentNewResponseEnvelopeErrorsJSON `json:"-"`
}

// contentNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ContentNewResponseEnvelopeErrors]
type contentNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContentNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contentNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}
