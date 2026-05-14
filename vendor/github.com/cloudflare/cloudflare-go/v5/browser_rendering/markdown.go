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

// MarkdownService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMarkdownService] method instead.
type MarkdownService struct {
	Options []option.RequestOption
}

// NewMarkdownService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMarkdownService(opts ...option.RequestOption) (r *MarkdownService) {
	r = &MarkdownService{}
	r.Options = opts
	return
}

// Gets markdown of a webpage from provided URL or HTML. Control page loading with
// `gotoOptions` and `waitFor*` options.
func (r *MarkdownService) New(ctx context.Context, params MarkdownNewParams, opts ...option.RequestOption) (res *string, err error) {
	var env MarkdownNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/browser-rendering/markdown", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type MarkdownNewParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// Cache TTL default is 5s. Set to 0 to disable.
	CacheTTL param.Field[float64] `query:"cacheTTL"`
	// The maximum duration allowed for the browser action to complete after the page
	// has loaded (such as taking screenshots, extracting content, or generating PDFs).
	// If this time limit is exceeded, the action stops and returns a timeout error.
	ActionTimeout param.Field[float64] `json:"actionTimeout"`
	// Adds a `<script>` tag into the page with the desired URL or content.
	AddScriptTag param.Field[[]MarkdownNewParamsAddScriptTag] `json:"addScriptTag"`
	// Adds a `<link rel="stylesheet">` tag into the page with the desired URL or a
	// `<style type="text/css">` tag with the content.
	AddStyleTag param.Field[[]MarkdownNewParamsAddStyleTag] `json:"addStyleTag"`
	// Only allow requests that match the provided regex patterns, eg. '/^.\*\.(css)'.
	AllowRequestPattern param.Field[[]string] `json:"allowRequestPattern"`
	// Only allow requests that match the provided resource types, eg. 'image' or
	// 'script'.
	AllowResourceTypes param.Field[[]MarkdownNewParamsAllowResourceType] `json:"allowResourceTypes"`
	// Provide credentials for HTTP authentication.
	Authenticate param.Field[MarkdownNewParamsAuthenticate] `json:"authenticate"`
	// Attempt to proceed when 'awaited' events fail or timeout.
	BestAttempt param.Field[bool] `json:"bestAttempt"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setcookie).
	Cookies          param.Field[[]MarkdownNewParamsCookie] `json:"cookies"`
	EmulateMediaType param.Field[string]                    `json:"emulateMediaType"`
	// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
	GotoOptions param.Field[MarkdownNewParamsGotoOptions] `json:"gotoOptions"`
	// Set the content of the page, eg: `<h1>Hello World!!</h1>`. Either `html` or
	// `url` must be set.
	HTML param.Field[string] `json:"html"`
	// Block undesired requests that match the provided regex patterns, eg.
	// '/^.\*\.(css)'.
	RejectRequestPattern param.Field[[]string] `json:"rejectRequestPattern"`
	// Block undesired requests that match the provided resource types, eg. 'image' or
	// 'script'.
	RejectResourceTypes  param.Field[[]MarkdownNewParamsRejectResourceType] `json:"rejectResourceTypes"`
	SetExtraHTTPHeaders  param.Field[map[string]string]                     `json:"setExtraHTTPHeaders"`
	SetJavaScriptEnabled param.Field[bool]                                  `json:"setJavaScriptEnabled"`
	// URL to navigate to, eg. `https://example.com`.
	URL       param.Field[string] `json:"url" format:"uri"`
	UserAgent param.Field[string] `json:"userAgent"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
	Viewport param.Field[MarkdownNewParamsViewport] `json:"viewport"`
	// Wait for the selector to appear in page. Check
	// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
	WaitForSelector param.Field[MarkdownNewParamsWaitForSelector] `json:"waitForSelector"`
	// Waits for a specified timeout before continuing.
	WaitForTimeout param.Field[float64] `json:"waitForTimeout"`
}

func (r MarkdownNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [MarkdownNewParams]'s query parameters as `url.Values`.
func (r MarkdownNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type MarkdownNewParamsAddScriptTag struct {
	ID      param.Field[string] `json:"id"`
	Content param.Field[string] `json:"content"`
	Type    param.Field[string] `json:"type"`
	URL     param.Field[string] `json:"url"`
}

func (r MarkdownNewParamsAddScriptTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MarkdownNewParamsAddStyleTag struct {
	Content param.Field[string] `json:"content"`
	URL     param.Field[string] `json:"url"`
}

func (r MarkdownNewParamsAddStyleTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MarkdownNewParamsAllowResourceType string

const (
	MarkdownNewParamsAllowResourceTypeDocument           MarkdownNewParamsAllowResourceType = "document"
	MarkdownNewParamsAllowResourceTypeStylesheet         MarkdownNewParamsAllowResourceType = "stylesheet"
	MarkdownNewParamsAllowResourceTypeImage              MarkdownNewParamsAllowResourceType = "image"
	MarkdownNewParamsAllowResourceTypeMedia              MarkdownNewParamsAllowResourceType = "media"
	MarkdownNewParamsAllowResourceTypeFont               MarkdownNewParamsAllowResourceType = "font"
	MarkdownNewParamsAllowResourceTypeScript             MarkdownNewParamsAllowResourceType = "script"
	MarkdownNewParamsAllowResourceTypeTexttrack          MarkdownNewParamsAllowResourceType = "texttrack"
	MarkdownNewParamsAllowResourceTypeXHR                MarkdownNewParamsAllowResourceType = "xhr"
	MarkdownNewParamsAllowResourceTypeFetch              MarkdownNewParamsAllowResourceType = "fetch"
	MarkdownNewParamsAllowResourceTypePrefetch           MarkdownNewParamsAllowResourceType = "prefetch"
	MarkdownNewParamsAllowResourceTypeEventsource        MarkdownNewParamsAllowResourceType = "eventsource"
	MarkdownNewParamsAllowResourceTypeWebsocket          MarkdownNewParamsAllowResourceType = "websocket"
	MarkdownNewParamsAllowResourceTypeManifest           MarkdownNewParamsAllowResourceType = "manifest"
	MarkdownNewParamsAllowResourceTypeSignedexchange     MarkdownNewParamsAllowResourceType = "signedexchange"
	MarkdownNewParamsAllowResourceTypePing               MarkdownNewParamsAllowResourceType = "ping"
	MarkdownNewParamsAllowResourceTypeCspviolationreport MarkdownNewParamsAllowResourceType = "cspviolationreport"
	MarkdownNewParamsAllowResourceTypePreflight          MarkdownNewParamsAllowResourceType = "preflight"
	MarkdownNewParamsAllowResourceTypeOther              MarkdownNewParamsAllowResourceType = "other"
)

func (r MarkdownNewParamsAllowResourceType) IsKnown() bool {
	switch r {
	case MarkdownNewParamsAllowResourceTypeDocument, MarkdownNewParamsAllowResourceTypeStylesheet, MarkdownNewParamsAllowResourceTypeImage, MarkdownNewParamsAllowResourceTypeMedia, MarkdownNewParamsAllowResourceTypeFont, MarkdownNewParamsAllowResourceTypeScript, MarkdownNewParamsAllowResourceTypeTexttrack, MarkdownNewParamsAllowResourceTypeXHR, MarkdownNewParamsAllowResourceTypeFetch, MarkdownNewParamsAllowResourceTypePrefetch, MarkdownNewParamsAllowResourceTypeEventsource, MarkdownNewParamsAllowResourceTypeWebsocket, MarkdownNewParamsAllowResourceTypeManifest, MarkdownNewParamsAllowResourceTypeSignedexchange, MarkdownNewParamsAllowResourceTypePing, MarkdownNewParamsAllowResourceTypeCspviolationreport, MarkdownNewParamsAllowResourceTypePreflight, MarkdownNewParamsAllowResourceTypeOther:
		return true
	}
	return false
}

// Provide credentials for HTTP authentication.
type MarkdownNewParamsAuthenticate struct {
	Password param.Field[string] `json:"password,required"`
	Username param.Field[string] `json:"username,required"`
}

func (r MarkdownNewParamsAuthenticate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MarkdownNewParamsCookie struct {
	Name         param.Field[string]                               `json:"name,required"`
	Value        param.Field[string]                               `json:"value,required"`
	Domain       param.Field[string]                               `json:"domain"`
	Expires      param.Field[float64]                              `json:"expires"`
	HTTPOnly     param.Field[bool]                                 `json:"httpOnly"`
	PartitionKey param.Field[string]                               `json:"partitionKey"`
	Path         param.Field[string]                               `json:"path"`
	Priority     param.Field[MarkdownNewParamsCookiesPriority]     `json:"priority"`
	SameParty    param.Field[bool]                                 `json:"sameParty"`
	SameSite     param.Field[MarkdownNewParamsCookiesSameSite]     `json:"sameSite"`
	Secure       param.Field[bool]                                 `json:"secure"`
	SourcePort   param.Field[float64]                              `json:"sourcePort"`
	SourceScheme param.Field[MarkdownNewParamsCookiesSourceScheme] `json:"sourceScheme"`
	URL          param.Field[string]                               `json:"url"`
}

func (r MarkdownNewParamsCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MarkdownNewParamsCookiesPriority string

const (
	MarkdownNewParamsCookiesPriorityLow    MarkdownNewParamsCookiesPriority = "Low"
	MarkdownNewParamsCookiesPriorityMedium MarkdownNewParamsCookiesPriority = "Medium"
	MarkdownNewParamsCookiesPriorityHigh   MarkdownNewParamsCookiesPriority = "High"
)

func (r MarkdownNewParamsCookiesPriority) IsKnown() bool {
	switch r {
	case MarkdownNewParamsCookiesPriorityLow, MarkdownNewParamsCookiesPriorityMedium, MarkdownNewParamsCookiesPriorityHigh:
		return true
	}
	return false
}

type MarkdownNewParamsCookiesSameSite string

const (
	MarkdownNewParamsCookiesSameSiteStrict MarkdownNewParamsCookiesSameSite = "Strict"
	MarkdownNewParamsCookiesSameSiteLax    MarkdownNewParamsCookiesSameSite = "Lax"
	MarkdownNewParamsCookiesSameSiteNone   MarkdownNewParamsCookiesSameSite = "None"
)

func (r MarkdownNewParamsCookiesSameSite) IsKnown() bool {
	switch r {
	case MarkdownNewParamsCookiesSameSiteStrict, MarkdownNewParamsCookiesSameSiteLax, MarkdownNewParamsCookiesSameSiteNone:
		return true
	}
	return false
}

type MarkdownNewParamsCookiesSourceScheme string

const (
	MarkdownNewParamsCookiesSourceSchemeUnset     MarkdownNewParamsCookiesSourceScheme = "Unset"
	MarkdownNewParamsCookiesSourceSchemeNonSecure MarkdownNewParamsCookiesSourceScheme = "NonSecure"
	MarkdownNewParamsCookiesSourceSchemeSecure    MarkdownNewParamsCookiesSourceScheme = "Secure"
)

func (r MarkdownNewParamsCookiesSourceScheme) IsKnown() bool {
	switch r {
	case MarkdownNewParamsCookiesSourceSchemeUnset, MarkdownNewParamsCookiesSourceSchemeNonSecure, MarkdownNewParamsCookiesSourceSchemeSecure:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
type MarkdownNewParamsGotoOptions struct {
	Referer        param.Field[string]                                     `json:"referer"`
	ReferrerPolicy param.Field[string]                                     `json:"referrerPolicy"`
	Timeout        param.Field[float64]                                    `json:"timeout"`
	WaitUntil      param.Field[MarkdownNewParamsGotoOptionsWaitUntilUnion] `json:"waitUntil"`
}

func (r MarkdownNewParamsGotoOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [browser_rendering.MarkdownNewParamsGotoOptionsWaitUntilString],
// [browser_rendering.MarkdownNewParamsGotoOptionsWaitUntilArray].
type MarkdownNewParamsGotoOptionsWaitUntilUnion interface {
	implementsMarkdownNewParamsGotoOptionsWaitUntilUnion()
}

type MarkdownNewParamsGotoOptionsWaitUntilString string

const (
	MarkdownNewParamsGotoOptionsWaitUntilStringLoad             MarkdownNewParamsGotoOptionsWaitUntilString = "load"
	MarkdownNewParamsGotoOptionsWaitUntilStringDomcontentloaded MarkdownNewParamsGotoOptionsWaitUntilString = "domcontentloaded"
	MarkdownNewParamsGotoOptionsWaitUntilStringNetworkidle0     MarkdownNewParamsGotoOptionsWaitUntilString = "networkidle0"
	MarkdownNewParamsGotoOptionsWaitUntilStringNetworkidle2     MarkdownNewParamsGotoOptionsWaitUntilString = "networkidle2"
)

func (r MarkdownNewParamsGotoOptionsWaitUntilString) IsKnown() bool {
	switch r {
	case MarkdownNewParamsGotoOptionsWaitUntilStringLoad, MarkdownNewParamsGotoOptionsWaitUntilStringDomcontentloaded, MarkdownNewParamsGotoOptionsWaitUntilStringNetworkidle0, MarkdownNewParamsGotoOptionsWaitUntilStringNetworkidle2:
		return true
	}
	return false
}

func (r MarkdownNewParamsGotoOptionsWaitUntilString) implementsMarkdownNewParamsGotoOptionsWaitUntilUnion() {
}

type MarkdownNewParamsGotoOptionsWaitUntilArray []MarkdownNewParamsGotoOptionsWaitUntilArrayItem

func (r MarkdownNewParamsGotoOptionsWaitUntilArray) implementsMarkdownNewParamsGotoOptionsWaitUntilUnion() {
}

type MarkdownNewParamsGotoOptionsWaitUntilArrayItem string

const (
	MarkdownNewParamsGotoOptionsWaitUntilArrayItemLoad             MarkdownNewParamsGotoOptionsWaitUntilArrayItem = "load"
	MarkdownNewParamsGotoOptionsWaitUntilArrayItemDomcontentloaded MarkdownNewParamsGotoOptionsWaitUntilArrayItem = "domcontentloaded"
	MarkdownNewParamsGotoOptionsWaitUntilArrayItemNetworkidle0     MarkdownNewParamsGotoOptionsWaitUntilArrayItem = "networkidle0"
	MarkdownNewParamsGotoOptionsWaitUntilArrayItemNetworkidle2     MarkdownNewParamsGotoOptionsWaitUntilArrayItem = "networkidle2"
)

func (r MarkdownNewParamsGotoOptionsWaitUntilArrayItem) IsKnown() bool {
	switch r {
	case MarkdownNewParamsGotoOptionsWaitUntilArrayItemLoad, MarkdownNewParamsGotoOptionsWaitUntilArrayItemDomcontentloaded, MarkdownNewParamsGotoOptionsWaitUntilArrayItemNetworkidle0, MarkdownNewParamsGotoOptionsWaitUntilArrayItemNetworkidle2:
		return true
	}
	return false
}

type MarkdownNewParamsRejectResourceType string

const (
	MarkdownNewParamsRejectResourceTypeDocument           MarkdownNewParamsRejectResourceType = "document"
	MarkdownNewParamsRejectResourceTypeStylesheet         MarkdownNewParamsRejectResourceType = "stylesheet"
	MarkdownNewParamsRejectResourceTypeImage              MarkdownNewParamsRejectResourceType = "image"
	MarkdownNewParamsRejectResourceTypeMedia              MarkdownNewParamsRejectResourceType = "media"
	MarkdownNewParamsRejectResourceTypeFont               MarkdownNewParamsRejectResourceType = "font"
	MarkdownNewParamsRejectResourceTypeScript             MarkdownNewParamsRejectResourceType = "script"
	MarkdownNewParamsRejectResourceTypeTexttrack          MarkdownNewParamsRejectResourceType = "texttrack"
	MarkdownNewParamsRejectResourceTypeXHR                MarkdownNewParamsRejectResourceType = "xhr"
	MarkdownNewParamsRejectResourceTypeFetch              MarkdownNewParamsRejectResourceType = "fetch"
	MarkdownNewParamsRejectResourceTypePrefetch           MarkdownNewParamsRejectResourceType = "prefetch"
	MarkdownNewParamsRejectResourceTypeEventsource        MarkdownNewParamsRejectResourceType = "eventsource"
	MarkdownNewParamsRejectResourceTypeWebsocket          MarkdownNewParamsRejectResourceType = "websocket"
	MarkdownNewParamsRejectResourceTypeManifest           MarkdownNewParamsRejectResourceType = "manifest"
	MarkdownNewParamsRejectResourceTypeSignedexchange     MarkdownNewParamsRejectResourceType = "signedexchange"
	MarkdownNewParamsRejectResourceTypePing               MarkdownNewParamsRejectResourceType = "ping"
	MarkdownNewParamsRejectResourceTypeCspviolationreport MarkdownNewParamsRejectResourceType = "cspviolationreport"
	MarkdownNewParamsRejectResourceTypePreflight          MarkdownNewParamsRejectResourceType = "preflight"
	MarkdownNewParamsRejectResourceTypeOther              MarkdownNewParamsRejectResourceType = "other"
)

func (r MarkdownNewParamsRejectResourceType) IsKnown() bool {
	switch r {
	case MarkdownNewParamsRejectResourceTypeDocument, MarkdownNewParamsRejectResourceTypeStylesheet, MarkdownNewParamsRejectResourceTypeImage, MarkdownNewParamsRejectResourceTypeMedia, MarkdownNewParamsRejectResourceTypeFont, MarkdownNewParamsRejectResourceTypeScript, MarkdownNewParamsRejectResourceTypeTexttrack, MarkdownNewParamsRejectResourceTypeXHR, MarkdownNewParamsRejectResourceTypeFetch, MarkdownNewParamsRejectResourceTypePrefetch, MarkdownNewParamsRejectResourceTypeEventsource, MarkdownNewParamsRejectResourceTypeWebsocket, MarkdownNewParamsRejectResourceTypeManifest, MarkdownNewParamsRejectResourceTypeSignedexchange, MarkdownNewParamsRejectResourceTypePing, MarkdownNewParamsRejectResourceTypeCspviolationreport, MarkdownNewParamsRejectResourceTypePreflight, MarkdownNewParamsRejectResourceTypeOther:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
type MarkdownNewParamsViewport struct {
	Height            param.Field[float64] `json:"height,required"`
	Width             param.Field[float64] `json:"width,required"`
	DeviceScaleFactor param.Field[float64] `json:"deviceScaleFactor"`
	HasTouch          param.Field[bool]    `json:"hasTouch"`
	IsLandscape       param.Field[bool]    `json:"isLandscape"`
	IsMobile          param.Field[bool]    `json:"isMobile"`
}

func (r MarkdownNewParamsViewport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Wait for the selector to appear in page. Check
// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
type MarkdownNewParamsWaitForSelector struct {
	Selector param.Field[string]                                  `json:"selector,required"`
	Hidden   param.Field[MarkdownNewParamsWaitForSelectorHidden]  `json:"hidden"`
	Timeout  param.Field[float64]                                 `json:"timeout"`
	Visible  param.Field[MarkdownNewParamsWaitForSelectorVisible] `json:"visible"`
}

func (r MarkdownNewParamsWaitForSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MarkdownNewParamsWaitForSelectorHidden bool

const (
	MarkdownNewParamsWaitForSelectorHiddenTrue MarkdownNewParamsWaitForSelectorHidden = true
)

func (r MarkdownNewParamsWaitForSelectorHidden) IsKnown() bool {
	switch r {
	case MarkdownNewParamsWaitForSelectorHiddenTrue:
		return true
	}
	return false
}

type MarkdownNewParamsWaitForSelectorVisible bool

const (
	MarkdownNewParamsWaitForSelectorVisibleTrue MarkdownNewParamsWaitForSelectorVisible = true
)

func (r MarkdownNewParamsWaitForSelectorVisible) IsKnown() bool {
	switch r {
	case MarkdownNewParamsWaitForSelectorVisibleTrue:
		return true
	}
	return false
}

type MarkdownNewResponseEnvelope struct {
	// Response status
	Status bool                                `json:"status,required"`
	Errors []MarkdownNewResponseEnvelopeErrors `json:"errors"`
	// Markdown
	Result string                          `json:"result"`
	JSON   markdownNewResponseEnvelopeJSON `json:"-"`
}

// markdownNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [MarkdownNewResponseEnvelope]
type markdownNewResponseEnvelopeJSON struct {
	Status      apijson.Field
	Errors      apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MarkdownNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r markdownNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MarkdownNewResponseEnvelopeErrors struct {
	// Error code
	Code float64 `json:"code,required"`
	// Error Message
	Message string                                `json:"message,required"`
	JSON    markdownNewResponseEnvelopeErrorsJSON `json:"-"`
}

// markdownNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [MarkdownNewResponseEnvelopeErrors]
type markdownNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MarkdownNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r markdownNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}
