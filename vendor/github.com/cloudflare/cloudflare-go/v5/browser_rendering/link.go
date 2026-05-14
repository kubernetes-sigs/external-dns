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

// LinkService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLinkService] method instead.
type LinkService struct {
	Options []option.RequestOption
}

// NewLinkService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewLinkService(opts ...option.RequestOption) (r *LinkService) {
	r = &LinkService{}
	r.Options = opts
	return
}

// Get links from a web page.
func (r *LinkService) New(ctx context.Context, params LinkNewParams, opts ...option.RequestOption) (res *[]string, err error) {
	var env LinkNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/browser-rendering/links", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type LinkNewParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// Cache TTL default is 5s. Set to 0 to disable.
	CacheTTL param.Field[float64] `query:"cacheTTL"`
	// The maximum duration allowed for the browser action to complete after the page
	// has loaded (such as taking screenshots, extracting content, or generating PDFs).
	// If this time limit is exceeded, the action stops and returns a timeout error.
	ActionTimeout param.Field[float64] `json:"actionTimeout"`
	// Adds a `<script>` tag into the page with the desired URL or content.
	AddScriptTag param.Field[[]LinkNewParamsAddScriptTag] `json:"addScriptTag"`
	// Adds a `<link rel="stylesheet">` tag into the page with the desired URL or a
	// `<style type="text/css">` tag with the content.
	AddStyleTag param.Field[[]LinkNewParamsAddStyleTag] `json:"addStyleTag"`
	// Only allow requests that match the provided regex patterns, eg. '/^.\*\.(css)'.
	AllowRequestPattern param.Field[[]string] `json:"allowRequestPattern"`
	// Only allow requests that match the provided resource types, eg. 'image' or
	// 'script'.
	AllowResourceTypes param.Field[[]LinkNewParamsAllowResourceType] `json:"allowResourceTypes"`
	// Provide credentials for HTTP authentication.
	Authenticate param.Field[LinkNewParamsAuthenticate] `json:"authenticate"`
	// Attempt to proceed when 'awaited' events fail or timeout.
	BestAttempt param.Field[bool] `json:"bestAttempt"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setcookie).
	Cookies          param.Field[[]LinkNewParamsCookie] `json:"cookies"`
	EmulateMediaType param.Field[string]                `json:"emulateMediaType"`
	// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
	GotoOptions param.Field[LinkNewParamsGotoOptions] `json:"gotoOptions"`
	// Set the content of the page, eg: `<h1>Hello World!!</h1>`. Either `html` or
	// `url` must be set.
	HTML param.Field[string] `json:"html"`
	// Block undesired requests that match the provided regex patterns, eg.
	// '/^.\*\.(css)'.
	RejectRequestPattern param.Field[[]string] `json:"rejectRequestPattern"`
	// Block undesired requests that match the provided resource types, eg. 'image' or
	// 'script'.
	RejectResourceTypes  param.Field[[]LinkNewParamsRejectResourceType] `json:"rejectResourceTypes"`
	SetExtraHTTPHeaders  param.Field[map[string]string]                 `json:"setExtraHTTPHeaders"`
	SetJavaScriptEnabled param.Field[bool]                              `json:"setJavaScriptEnabled"`
	// URL to navigate to, eg. `https://example.com`.
	URL       param.Field[string] `json:"url" format:"uri"`
	UserAgent param.Field[string] `json:"userAgent"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
	Viewport         param.Field[LinkNewParamsViewport] `json:"viewport"`
	VisibleLinksOnly param.Field[bool]                  `json:"visibleLinksOnly"`
	// Wait for the selector to appear in page. Check
	// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
	WaitForSelector param.Field[LinkNewParamsWaitForSelector] `json:"waitForSelector"`
	// Waits for a specified timeout before continuing.
	WaitForTimeout param.Field[float64] `json:"waitForTimeout"`
}

func (r LinkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [LinkNewParams]'s query parameters as `url.Values`.
func (r LinkNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LinkNewParamsAddScriptTag struct {
	ID      param.Field[string] `json:"id"`
	Content param.Field[string] `json:"content"`
	Type    param.Field[string] `json:"type"`
	URL     param.Field[string] `json:"url"`
}

func (r LinkNewParamsAddScriptTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LinkNewParamsAddStyleTag struct {
	Content param.Field[string] `json:"content"`
	URL     param.Field[string] `json:"url"`
}

func (r LinkNewParamsAddStyleTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LinkNewParamsAllowResourceType string

const (
	LinkNewParamsAllowResourceTypeDocument           LinkNewParamsAllowResourceType = "document"
	LinkNewParamsAllowResourceTypeStylesheet         LinkNewParamsAllowResourceType = "stylesheet"
	LinkNewParamsAllowResourceTypeImage              LinkNewParamsAllowResourceType = "image"
	LinkNewParamsAllowResourceTypeMedia              LinkNewParamsAllowResourceType = "media"
	LinkNewParamsAllowResourceTypeFont               LinkNewParamsAllowResourceType = "font"
	LinkNewParamsAllowResourceTypeScript             LinkNewParamsAllowResourceType = "script"
	LinkNewParamsAllowResourceTypeTexttrack          LinkNewParamsAllowResourceType = "texttrack"
	LinkNewParamsAllowResourceTypeXHR                LinkNewParamsAllowResourceType = "xhr"
	LinkNewParamsAllowResourceTypeFetch              LinkNewParamsAllowResourceType = "fetch"
	LinkNewParamsAllowResourceTypePrefetch           LinkNewParamsAllowResourceType = "prefetch"
	LinkNewParamsAllowResourceTypeEventsource        LinkNewParamsAllowResourceType = "eventsource"
	LinkNewParamsAllowResourceTypeWebsocket          LinkNewParamsAllowResourceType = "websocket"
	LinkNewParamsAllowResourceTypeManifest           LinkNewParamsAllowResourceType = "manifest"
	LinkNewParamsAllowResourceTypeSignedexchange     LinkNewParamsAllowResourceType = "signedexchange"
	LinkNewParamsAllowResourceTypePing               LinkNewParamsAllowResourceType = "ping"
	LinkNewParamsAllowResourceTypeCspviolationreport LinkNewParamsAllowResourceType = "cspviolationreport"
	LinkNewParamsAllowResourceTypePreflight          LinkNewParamsAllowResourceType = "preflight"
	LinkNewParamsAllowResourceTypeOther              LinkNewParamsAllowResourceType = "other"
)

func (r LinkNewParamsAllowResourceType) IsKnown() bool {
	switch r {
	case LinkNewParamsAllowResourceTypeDocument, LinkNewParamsAllowResourceTypeStylesheet, LinkNewParamsAllowResourceTypeImage, LinkNewParamsAllowResourceTypeMedia, LinkNewParamsAllowResourceTypeFont, LinkNewParamsAllowResourceTypeScript, LinkNewParamsAllowResourceTypeTexttrack, LinkNewParamsAllowResourceTypeXHR, LinkNewParamsAllowResourceTypeFetch, LinkNewParamsAllowResourceTypePrefetch, LinkNewParamsAllowResourceTypeEventsource, LinkNewParamsAllowResourceTypeWebsocket, LinkNewParamsAllowResourceTypeManifest, LinkNewParamsAllowResourceTypeSignedexchange, LinkNewParamsAllowResourceTypePing, LinkNewParamsAllowResourceTypeCspviolationreport, LinkNewParamsAllowResourceTypePreflight, LinkNewParamsAllowResourceTypeOther:
		return true
	}
	return false
}

// Provide credentials for HTTP authentication.
type LinkNewParamsAuthenticate struct {
	Password param.Field[string] `json:"password,required"`
	Username param.Field[string] `json:"username,required"`
}

func (r LinkNewParamsAuthenticate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LinkNewParamsCookie struct {
	Name         param.Field[string]                           `json:"name,required"`
	Value        param.Field[string]                           `json:"value,required"`
	Domain       param.Field[string]                           `json:"domain"`
	Expires      param.Field[float64]                          `json:"expires"`
	HTTPOnly     param.Field[bool]                             `json:"httpOnly"`
	PartitionKey param.Field[string]                           `json:"partitionKey"`
	Path         param.Field[string]                           `json:"path"`
	Priority     param.Field[LinkNewParamsCookiesPriority]     `json:"priority"`
	SameParty    param.Field[bool]                             `json:"sameParty"`
	SameSite     param.Field[LinkNewParamsCookiesSameSite]     `json:"sameSite"`
	Secure       param.Field[bool]                             `json:"secure"`
	SourcePort   param.Field[float64]                          `json:"sourcePort"`
	SourceScheme param.Field[LinkNewParamsCookiesSourceScheme] `json:"sourceScheme"`
	URL          param.Field[string]                           `json:"url"`
}

func (r LinkNewParamsCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LinkNewParamsCookiesPriority string

const (
	LinkNewParamsCookiesPriorityLow    LinkNewParamsCookiesPriority = "Low"
	LinkNewParamsCookiesPriorityMedium LinkNewParamsCookiesPriority = "Medium"
	LinkNewParamsCookiesPriorityHigh   LinkNewParamsCookiesPriority = "High"
)

func (r LinkNewParamsCookiesPriority) IsKnown() bool {
	switch r {
	case LinkNewParamsCookiesPriorityLow, LinkNewParamsCookiesPriorityMedium, LinkNewParamsCookiesPriorityHigh:
		return true
	}
	return false
}

type LinkNewParamsCookiesSameSite string

const (
	LinkNewParamsCookiesSameSiteStrict LinkNewParamsCookiesSameSite = "Strict"
	LinkNewParamsCookiesSameSiteLax    LinkNewParamsCookiesSameSite = "Lax"
	LinkNewParamsCookiesSameSiteNone   LinkNewParamsCookiesSameSite = "None"
)

func (r LinkNewParamsCookiesSameSite) IsKnown() bool {
	switch r {
	case LinkNewParamsCookiesSameSiteStrict, LinkNewParamsCookiesSameSiteLax, LinkNewParamsCookiesSameSiteNone:
		return true
	}
	return false
}

type LinkNewParamsCookiesSourceScheme string

const (
	LinkNewParamsCookiesSourceSchemeUnset     LinkNewParamsCookiesSourceScheme = "Unset"
	LinkNewParamsCookiesSourceSchemeNonSecure LinkNewParamsCookiesSourceScheme = "NonSecure"
	LinkNewParamsCookiesSourceSchemeSecure    LinkNewParamsCookiesSourceScheme = "Secure"
)

func (r LinkNewParamsCookiesSourceScheme) IsKnown() bool {
	switch r {
	case LinkNewParamsCookiesSourceSchemeUnset, LinkNewParamsCookiesSourceSchemeNonSecure, LinkNewParamsCookiesSourceSchemeSecure:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
type LinkNewParamsGotoOptions struct {
	Referer        param.Field[string]                                 `json:"referer"`
	ReferrerPolicy param.Field[string]                                 `json:"referrerPolicy"`
	Timeout        param.Field[float64]                                `json:"timeout"`
	WaitUntil      param.Field[LinkNewParamsGotoOptionsWaitUntilUnion] `json:"waitUntil"`
}

func (r LinkNewParamsGotoOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [browser_rendering.LinkNewParamsGotoOptionsWaitUntilString],
// [browser_rendering.LinkNewParamsGotoOptionsWaitUntilArray].
type LinkNewParamsGotoOptionsWaitUntilUnion interface {
	implementsLinkNewParamsGotoOptionsWaitUntilUnion()
}

type LinkNewParamsGotoOptionsWaitUntilString string

const (
	LinkNewParamsGotoOptionsWaitUntilStringLoad             LinkNewParamsGotoOptionsWaitUntilString = "load"
	LinkNewParamsGotoOptionsWaitUntilStringDomcontentloaded LinkNewParamsGotoOptionsWaitUntilString = "domcontentloaded"
	LinkNewParamsGotoOptionsWaitUntilStringNetworkidle0     LinkNewParamsGotoOptionsWaitUntilString = "networkidle0"
	LinkNewParamsGotoOptionsWaitUntilStringNetworkidle2     LinkNewParamsGotoOptionsWaitUntilString = "networkidle2"
)

func (r LinkNewParamsGotoOptionsWaitUntilString) IsKnown() bool {
	switch r {
	case LinkNewParamsGotoOptionsWaitUntilStringLoad, LinkNewParamsGotoOptionsWaitUntilStringDomcontentloaded, LinkNewParamsGotoOptionsWaitUntilStringNetworkidle0, LinkNewParamsGotoOptionsWaitUntilStringNetworkidle2:
		return true
	}
	return false
}

func (r LinkNewParamsGotoOptionsWaitUntilString) implementsLinkNewParamsGotoOptionsWaitUntilUnion() {}

type LinkNewParamsGotoOptionsWaitUntilArray []LinkNewParamsGotoOptionsWaitUntilArrayItem

func (r LinkNewParamsGotoOptionsWaitUntilArray) implementsLinkNewParamsGotoOptionsWaitUntilUnion() {}

type LinkNewParamsGotoOptionsWaitUntilArrayItem string

const (
	LinkNewParamsGotoOptionsWaitUntilArrayItemLoad             LinkNewParamsGotoOptionsWaitUntilArrayItem = "load"
	LinkNewParamsGotoOptionsWaitUntilArrayItemDomcontentloaded LinkNewParamsGotoOptionsWaitUntilArrayItem = "domcontentloaded"
	LinkNewParamsGotoOptionsWaitUntilArrayItemNetworkidle0     LinkNewParamsGotoOptionsWaitUntilArrayItem = "networkidle0"
	LinkNewParamsGotoOptionsWaitUntilArrayItemNetworkidle2     LinkNewParamsGotoOptionsWaitUntilArrayItem = "networkidle2"
)

func (r LinkNewParamsGotoOptionsWaitUntilArrayItem) IsKnown() bool {
	switch r {
	case LinkNewParamsGotoOptionsWaitUntilArrayItemLoad, LinkNewParamsGotoOptionsWaitUntilArrayItemDomcontentloaded, LinkNewParamsGotoOptionsWaitUntilArrayItemNetworkidle0, LinkNewParamsGotoOptionsWaitUntilArrayItemNetworkidle2:
		return true
	}
	return false
}

type LinkNewParamsRejectResourceType string

const (
	LinkNewParamsRejectResourceTypeDocument           LinkNewParamsRejectResourceType = "document"
	LinkNewParamsRejectResourceTypeStylesheet         LinkNewParamsRejectResourceType = "stylesheet"
	LinkNewParamsRejectResourceTypeImage              LinkNewParamsRejectResourceType = "image"
	LinkNewParamsRejectResourceTypeMedia              LinkNewParamsRejectResourceType = "media"
	LinkNewParamsRejectResourceTypeFont               LinkNewParamsRejectResourceType = "font"
	LinkNewParamsRejectResourceTypeScript             LinkNewParamsRejectResourceType = "script"
	LinkNewParamsRejectResourceTypeTexttrack          LinkNewParamsRejectResourceType = "texttrack"
	LinkNewParamsRejectResourceTypeXHR                LinkNewParamsRejectResourceType = "xhr"
	LinkNewParamsRejectResourceTypeFetch              LinkNewParamsRejectResourceType = "fetch"
	LinkNewParamsRejectResourceTypePrefetch           LinkNewParamsRejectResourceType = "prefetch"
	LinkNewParamsRejectResourceTypeEventsource        LinkNewParamsRejectResourceType = "eventsource"
	LinkNewParamsRejectResourceTypeWebsocket          LinkNewParamsRejectResourceType = "websocket"
	LinkNewParamsRejectResourceTypeManifest           LinkNewParamsRejectResourceType = "manifest"
	LinkNewParamsRejectResourceTypeSignedexchange     LinkNewParamsRejectResourceType = "signedexchange"
	LinkNewParamsRejectResourceTypePing               LinkNewParamsRejectResourceType = "ping"
	LinkNewParamsRejectResourceTypeCspviolationreport LinkNewParamsRejectResourceType = "cspviolationreport"
	LinkNewParamsRejectResourceTypePreflight          LinkNewParamsRejectResourceType = "preflight"
	LinkNewParamsRejectResourceTypeOther              LinkNewParamsRejectResourceType = "other"
)

func (r LinkNewParamsRejectResourceType) IsKnown() bool {
	switch r {
	case LinkNewParamsRejectResourceTypeDocument, LinkNewParamsRejectResourceTypeStylesheet, LinkNewParamsRejectResourceTypeImage, LinkNewParamsRejectResourceTypeMedia, LinkNewParamsRejectResourceTypeFont, LinkNewParamsRejectResourceTypeScript, LinkNewParamsRejectResourceTypeTexttrack, LinkNewParamsRejectResourceTypeXHR, LinkNewParamsRejectResourceTypeFetch, LinkNewParamsRejectResourceTypePrefetch, LinkNewParamsRejectResourceTypeEventsource, LinkNewParamsRejectResourceTypeWebsocket, LinkNewParamsRejectResourceTypeManifest, LinkNewParamsRejectResourceTypeSignedexchange, LinkNewParamsRejectResourceTypePing, LinkNewParamsRejectResourceTypeCspviolationreport, LinkNewParamsRejectResourceTypePreflight, LinkNewParamsRejectResourceTypeOther:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
type LinkNewParamsViewport struct {
	Height            param.Field[float64] `json:"height,required"`
	Width             param.Field[float64] `json:"width,required"`
	DeviceScaleFactor param.Field[float64] `json:"deviceScaleFactor"`
	HasTouch          param.Field[bool]    `json:"hasTouch"`
	IsLandscape       param.Field[bool]    `json:"isLandscape"`
	IsMobile          param.Field[bool]    `json:"isMobile"`
}

func (r LinkNewParamsViewport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Wait for the selector to appear in page. Check
// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
type LinkNewParamsWaitForSelector struct {
	Selector param.Field[string]                              `json:"selector,required"`
	Hidden   param.Field[LinkNewParamsWaitForSelectorHidden]  `json:"hidden"`
	Timeout  param.Field[float64]                             `json:"timeout"`
	Visible  param.Field[LinkNewParamsWaitForSelectorVisible] `json:"visible"`
}

func (r LinkNewParamsWaitForSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LinkNewParamsWaitForSelectorHidden bool

const (
	LinkNewParamsWaitForSelectorHiddenTrue LinkNewParamsWaitForSelectorHidden = true
)

func (r LinkNewParamsWaitForSelectorHidden) IsKnown() bool {
	switch r {
	case LinkNewParamsWaitForSelectorHiddenTrue:
		return true
	}
	return false
}

type LinkNewParamsWaitForSelectorVisible bool

const (
	LinkNewParamsWaitForSelectorVisibleTrue LinkNewParamsWaitForSelectorVisible = true
)

func (r LinkNewParamsWaitForSelectorVisible) IsKnown() bool {
	switch r {
	case LinkNewParamsWaitForSelectorVisibleTrue:
		return true
	}
	return false
}

type LinkNewResponseEnvelope struct {
	Result []string `json:"result,required"`
	// Response status
	Status bool                            `json:"status,required"`
	Errors []LinkNewResponseEnvelopeErrors `json:"errors"`
	JSON   linkNewResponseEnvelopeJSON     `json:"-"`
}

// linkNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [LinkNewResponseEnvelope]
type linkNewResponseEnvelopeJSON struct {
	Result      apijson.Field
	Status      apijson.Field
	Errors      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LinkNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r linkNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type LinkNewResponseEnvelopeErrors struct {
	// Error code
	Code float64 `json:"code,required"`
	// Error Message
	Message string                            `json:"message,required"`
	JSON    linkNewResponseEnvelopeErrorsJSON `json:"-"`
}

// linkNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [LinkNewResponseEnvelopeErrors]
type linkNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LinkNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r linkNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}
